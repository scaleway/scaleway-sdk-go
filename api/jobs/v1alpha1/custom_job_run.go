package jobs

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultRetryInterval = 15 * time.Second
	defaultTimeout       = 15 * time.Minute
)

type WaitForJobRunRequest struct {
	JobRunID      string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForJobRun waits for the job run to be in a "terminal state" before returning.
// This function can be used to wait for a job run to fail for example.
func (s *API) WaitForJobRun(req *WaitForJobRunRequest, opts ...scw.RequestOption) (*JobRun, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[JobRunState]struct{}{
		JobRunStateSucceeded: {},
		JobRunStateFailed:    {},
		JobRunStateCanceled:  {},
	}

	jobRun, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetJobRun(&GetJobRunRequest{
				JobRunID: req.JobRunID,
				Region:   req.Region,
			}, opts...)
			if err != nil {
				return nil, false, err
			}
			_, isTerminal := terminalStatus[res.State]

			return res, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for job run failed")
	}

	return jobRun.(*JobRun), nil
}

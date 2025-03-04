package edge_services //nolint:revive

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultRetryInterval = 5 * time.Second
	defaultTimeout       = 5 * time.Minute
)

// WaitForPipelineRequest is used by WaitForPipeline method.
type WaitForPipelineRequest struct {
	PipelineID    string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForPipeline wait for a pipeline to be in a "terminal state" before returning.
func (s *API) WaitForPipeline(req *WaitForPipelineRequest, opts ...scw.RequestOption) (*Pipeline, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[PipelineStatus]struct{}{
		PipelineStatusReady:         {},
		PipelineStatusError:         {},
		PipelineStatusUnknownStatus: {},
		PipelineStatusWarning:       {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			pipeline, err := s.GetPipeline(&GetPipelineRequest{
				PipelineID: req.PipelineID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[pipeline.Status]
			return pipeline, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for pipeline failed")
	}

	return res.(*Pipeline), nil
}

// WaitForPurgeRequestRequest is used by WaitForPurgeRequest method.
type WaitForPurgeRequestRequest struct {
	PurgeRequestID string
	Timeout        *time.Duration
	RetryInterval  *time.Duration
}

// WaitForPurgeRequest wait for a purge request to be in a "terminal state" before returning.
func (s *API) WaitForPurgeRequest(req *WaitForPurgeRequestRequest, opts ...scw.RequestOption) (*PurgeRequest, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[PurgeRequestStatus]struct{}{
		PurgeRequestStatusDone:          {},
		PurgeRequestStatusError:         {},
		PurgeRequestStatusUnknownStatus: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			purgeRequest, err := s.GetPurgeRequest(&GetPurgeRequestRequest{
				PurgeRequestID: req.PurgeRequestID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[purgeRequest.Status]
			return purgeRequest, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for purge request failed")
	}

	return res.(*PurgeRequest), nil
}

package cockpit

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultRetryInterval = 3 * time.Second
	defaultTimeout       = 5 * time.Minute
)

// WaitForCockpitRequest is used by WaitForCockpit method.
type WaitForCockpitRequest struct {
	ProjectID     string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

func (s *API) WaitForCockpit(
	req *WaitForCockpitRequest,
	opts ...scw.RequestOption,
) (*Cockpit, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	transientStatus := map[CockpitStatus]struct{}{
		CockpitStatusCreating: {},
		CockpitStatusUpdating: {},
		CockpitStatusDeleting: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			namespace, err := s.GetCockpit(&GetCockpitRequest{
				ProjectID: req.ProjectID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTransient := transientStatus[namespace.Status]
			return namespace, !isTransient, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})

	if err != nil {
		return nil, errors.Wrap(err, "waiting for Cockpit failed")
	}

	return res.(*Cockpit), nil
}

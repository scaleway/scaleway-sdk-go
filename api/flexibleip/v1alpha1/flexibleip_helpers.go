package flexibleip

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	waitForFlexibleIPDefaultTimeout = 15 * time.Minute
	defaultRetryInterval            = 5 * time.Second
)

// WaitForFlexibleIPRequest is used by WaitForFlexibleIP method.
type WaitForFlexibleIPRequest struct {
	FipID         string
	Zone          scw.Zone
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForFlexibleIP waits for the FlexibleIP to be in a ready state before returning.
func (s *API) WaitForFlexibleIP(req *WaitForFlexibleIPRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	timeout := waitForFlexibleIPDefaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[FlexibleIPStatus]struct{}{
		FlexibleIPStatusError:    {},
		FlexibleIPStatusReady:    {},
		FlexibleIPStatusAttached: {},
		FlexibleIPStatusLocked:   {},
	}

	fip, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			fip, err := s.GetFlexibleIP(&GetFlexibleIPRequest{
				FipID: req.FipID,
				Zone:  req.Zone,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[fip.Status]
			return fip, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for FlexibleIP failed")
	}

	return fip.(*FlexibleIP), nil
}

package iot

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	waitForHubDefaultTimeout = 15 * time.Minute
	defaultRetryInterval     = 5 * time.Second
)

// WaitForHubRequest is used by WaitForHub method.
type WaitForHubRequest struct {
	HubID         string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForHub waits for the hub to be in a ready state before returning.
func (s *API) WaitForHub(req *WaitForHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	timeout := waitForHubDefaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[HubStatus]struct{}{
		HubStatusError:    {},
		HubStatusReady:    {},
		HubStatusDisabled: {},
	}

	hub, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			hub, err := s.GetHub(&GetHubRequest{
				HubID:  req.HubID,
				Region: req.Region,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[hub.Status]
			return hub, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for hub failed")
	}

	return hub.(*Hub), nil
}

package webhosting

import (
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/async"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/scw"
)

const (
	defaultRetryInterval = 5 * time.Second
	defaultTimeout       = 5 * time.Minute
)

// WaitForHostingRequest is used by WaitForHosting method.
type WaitForHostingRequest struct {
	HostingID     string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForHosting wait for a hosting to be in a "terminal state" before returning.
func (s *API) WaitForHosting(req *WaitForHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[HostingStatus]struct{}{
		HostingStatusReady:         {},
		HostingStatusError:         {},
		HostingStatusUnknownStatus: {},
		HostingStatusLocked:        {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			hosting, err := s.GetHosting(&GetHostingRequest{
				HostingID: req.HostingID,
				Region:    req.Region,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[hosting.Status]
			return hosting, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for hosting failed")
	}

	return res.(*Hosting), nil
}

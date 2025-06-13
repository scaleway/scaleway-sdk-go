package webhosting

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

// WaitForHostingRequest is used by WaitForHosting method.
type WaitForHostingRequest struct {
	HostingID     string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForHosting waits for a hosting to be in a "terminal" state before returning.
// Terminal states are defined as: HostingStatusReady, HostingStatusError, HostingStatusUnknownStatus, and HostingStatusLocked.
func (s *HostingAPI) WaitForHosting(req *WaitForHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
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
		Get: func() (any, bool, error) {
			hosting, err := s.GetHosting(&HostingAPIGetHostingRequest{
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

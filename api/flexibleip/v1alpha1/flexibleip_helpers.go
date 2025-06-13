package flexibleip

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
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

	fipTerminalStatus := map[FlexibleIPStatus]struct{}{
		FlexibleIPStatusError:    {},
		FlexibleIPStatusReady:    {},
		FlexibleIPStatusAttached: {},
		FlexibleIPStatusLocked:   {},
	}

	macAddressTerminalStatus := map[MACAddressStatus]struct{}{
		MACAddressStatusUnknown: {},
		MACAddressStatusReady:   {},
		MACAddressStatusUsed:    {},
		MACAddressStatusError:   {},
	}

	fip, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			fip, err := s.GetFlexibleIP(&GetFlexibleIPRequest{
				FipID: req.FipID,
				Zone:  req.Zone,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			// Check if the MACAddress is in a terminal state
			isMacAddressTerminal := true
			if fip.MacAddress != nil {
				_, isMacAddressTerminal = macAddressTerminalStatus[fip.MacAddress.Status]
			}

			_, isTerminal := fipTerminalStatus[fip.Status]
			return fip, isTerminal && isMacAddressTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for FlexibleIP failed")
	}

	return fip.(*FlexibleIP), nil
}

package baremetal

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultRetryInterval = 15 * time.Second
	defaultTimeout       = 2 * time.Hour
)

// WaitForServerPrivateNetworksRequest is used by WaitForServerPrivateNetworks method.
type WaitForServerPrivateNetworksRequest struct {
	ServerID      string
	Zone          scw.Zone
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForServerPrivateNetworks wait for all server private networks to be in a "terminal state" before returning.
// This function can be used to wait for all server private networks to be set.
func (s *PrivateNetworkAPI) WaitForServerPrivateNetworks(req *WaitForServerPrivateNetworksRequest, opts ...scw.RequestOption) ([]*ServerPrivateNetwork, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[ServerPrivateNetworkStatus]struct{}{
		ServerPrivateNetworkStatusAttached:      {},
		ServerPrivateNetworkStatusError:         {},
		ServerPrivateNetworkStatusUnknownStatus: {},
		ServerPrivateNetworkStatusLocked:        {},
	}

	serverPrivateNetwork, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.ListServerPrivateNetworks(&PrivateNetworkAPIListServerPrivateNetworksRequest{
				ServerID: &req.ServerID,
				Zone:     req.Zone,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			for i := range res.ServerPrivateNetworks {
				_, isTerminal := terminalStatus[res.ServerPrivateNetworks[i].Status]
				if !isTerminal {
					return res.ServerPrivateNetworks, isTerminal, nil
				}
			}
			return res.ServerPrivateNetworks, true, err
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for server private networks failed")
	}

	return serverPrivateNetwork.([]*ServerPrivateNetwork), nil
}

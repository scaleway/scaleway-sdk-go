package applesilicon

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultRetryInterval = 15 * time.Second
	defaultTimeout       = 60 * time.Minute
)

// WaitForInstanceRequest is used by WaitForServer method.
type WaitForServerRequest struct {
	ServerID      string
	Zone          scw.Zone
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForServer waits for the instance to be in a "terminal state" before returning.
// This function can be used to wait for an instance to be ready for example.
func (s *API) WaitForServer(req *WaitForServerRequest, opts ...scw.RequestOption) (*Server, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[ServerStatus]struct{}{
		ServerStatusReady: {},
		ServerStatusError: {},
	}

	server, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := s.GetServer(&GetServerRequest{
				ServerID: req.ServerID,
				Zone:     req.Zone,
			}, opts...)
			if err != nil {
				return nil, false, err
			}
			_, isTerminal := terminalStatus[res.Status]

			return res, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for server failed")
	}
	return server.(*Server), nil
}

func (s *API) WaitForPossibleDeletion(req *WaitForServerRequest, opts ...scw.RequestOption) (*Server, error) {
	server, err := s.WaitForServer(&WaitForServerRequest{
		ServerID:      req.ServerID,
		Zone:          req.Zone,
		Timeout:       scw.TimeDurationPtr(defaultTimeout),
		RetryInterval: scw.TimeDurationPtr(defaultRetryInterval),
	}, opts...,
	)
	if err != nil {
		return nil, errors.Wrap(err, "waiting for server failed")
	}
	timeToDelete := *server.DeletableAt
	time.Sleep(time.Until(timeToDelete))
	return server, nil
}

func (s *PrivateNetworkAPI) WaitForServerPrivateNetworks(req *WaitForServerRequest, opts ...scw.RequestOption) (*[]ServerPrivateNetwork, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[ServerPrivateNetworkServerStatus]struct{}{
		ServerPrivateNetworkServerStatusAttached:      {},
		ServerPrivateNetworkServerStatusError:         {},
		ServerPrivateNetworkServerStatusUnknownStatus: {},
		ServerPrivateNetworkServerStatusLocked:        {},
	}

	serverPrivateNetworks, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := s.ListServerPrivateNetworks(&PrivateNetworkAPIListServerPrivateNetworksRequest{
				Zone:     req.Zone,
				ServerID: &req.ServerID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}
			for _, serverPrivateNetwork := range res.ServerPrivateNetworks {
				_, isTerminal := terminalStatus[serverPrivateNetwork.Status]
				if !isTerminal {
					return res.ServerPrivateNetworks, isTerminal, err
				}
			}
			return res.ServerPrivateNetworks, true, err
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for server private network failed")
	}

	return serverPrivateNetworks.(*[]ServerPrivateNetwork), nil
}

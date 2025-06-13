package dedibox

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultRetryInterval = time.Second * 15
	defaultTimeout       = time.Minute * 30
)

type WaitForServiceRequest struct {
	ServiceID     uint64
	Zone          scw.Zone
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

func (s *API) WaitForService(req *WaitForServiceRequest, opts ...scw.RequestOption) (*Service, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	terminalStatus := map[ServiceProvisioningStatus]struct{}{
		ServiceProvisioningStatusReady:   {},
		ServiceProvisioningStatusError:   {},
		ServiceProvisioningStatusExpired: {},
	}
	service, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			service, err := s.GetService(&GetServiceRequest{
				Zone:      req.Zone,
				ServiceID: req.ServiceID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}
			_, isTerminal := terminalStatus[service.ProvisioningStatus]
			return service, isTerminal, nil
		},
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
		Timeout:          timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for service failed")
	}
	return service.(*Service), nil
}

type WaitForServerRequest struct {
	ServerID      uint64
	Zone          scw.Zone
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

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
		ServerStatusReady:   {},
		ServerStatusError:   {},
		ServerStatusLocked:  {},
		ServerStatusStopped: {},
		ServerStatusBusy:    {},
		ServerStatusRescue:  {},
	}
	server, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			server, err := s.GetServer(&GetServerRequest{
				Zone:     req.Zone,
				ServerID: req.ServerID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}
			_, isTerminal := terminalStatus[server.Status]
			return server, isTerminal, nil
		},
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
		Timeout:          timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for server failed")
	}
	return server.(*Server), nil
}

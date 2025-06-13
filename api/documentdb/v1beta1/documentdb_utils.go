package documentdb

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultRetryInterval = 15 * time.Second
	defaultTimeout       = 15 * time.Minute
)

// WaitForInstanceRequest is used by WaitForInstance method.
type WaitForInstanceRequest struct {
	InstanceID    string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForInstance waits for the instance to be in a "terminal state" before returning.
// This function can be used to wait for an instance to be ready for example.
func (s *API) WaitForInstance(req *WaitForInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[InstanceStatus]struct{}{
		InstanceStatusReady:    {},
		InstanceStatusDiskFull: {},
		InstanceStatusError:    {},
	}

	instance, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetInstance(&GetInstanceRequest{
				InstanceID: req.InstanceID,
				Region:     req.Region,
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
		return nil, errors.Wrap(err, "waiting for instance failed")
	}
	return instance.(*Instance), nil
}

type WaitForInstanceLogRequest struct {
	InstanceLogID string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForInstanceLog waits for the instance logs to be in a "terminal state" before returning.
// This function can be used to wait for an instance logs to be ready for example.
func (s *API) WaitForInstanceLog(req *WaitForInstanceLogRequest, opts ...scw.RequestOption) (*InstanceLog, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[InstanceLogStatus]struct{}{
		InstanceLogStatusReady: {},
		InstanceLogStatusError: {},
	}

	logs, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetInstanceLog(&GetInstanceLogRequest{
				Region: req.Region,
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
		return nil, errors.Wrap(err, "waiting for instance logs failed")
	}
	return logs.(*InstanceLog), nil
}

// WaitForReadReplicaRequest is used by WaitForReadReplica method.
type WaitForReadReplicaRequest struct {
	ReadReplicaID string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForReadReplica waits for the read replica to be in a "terminal state" before returning.
// This function can be used to wait for a read replica to be ready for example.
func (s *API) WaitForReadReplica(req *WaitForReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[ReadReplicaStatus]struct{}{
		ReadReplicaStatusReady: {},
		ReadReplicaStatusError: {},
	}

	readReplica, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetReadReplica(&GetReadReplicaRequest{
				ReadReplicaID: req.ReadReplicaID,
				Region:        req.Region,
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
		return nil, errors.Wrap(err, "waiting for read replica failed")
	}
	return readReplica.(*ReadReplica), nil
}

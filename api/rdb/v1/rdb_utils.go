package rdb

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
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
func (s *API) WaitForInstance(req *WaitForInstanceRequest) (*Instance, error) {
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
		Get: func() (interface{}, bool, error) {
			res, err := s.GetInstance(&GetInstanceRequest{
				InstanceID: req.InstanceID,
				Region:     req.Region,
			})

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

type WaitForDatabaseBackupRequest struct {
	DatabaseBackupID string
	Region           scw.Region
	Timeout          *time.Duration
	RetryInterval    *time.Duration
}

// WaitForDatabaseBackup waits for the backup to be in a "terminal state" before returning.
// This function can be used to wait for a backup to be ready for example.
func (s *API) WaitForDatabaseBackup(req *WaitForDatabaseBackupRequest) (*DatabaseBackup, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[DatabaseBackupStatus]struct{}{
		DatabaseBackupStatusReady: {},
		DatabaseBackupStatusError: {},
	}

	backup, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := s.GetDatabaseBackup(&GetDatabaseBackupRequest{
				DatabaseBackupID: req.DatabaseBackupID,
				Region:           req.Region,
			})

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
		return nil, errors.Wrap(err, "waiting for database backup failed")
	}
	return backup.(*DatabaseBackup), nil
}

type WaitForInstanceLogRequest struct {
	InstanceLogID string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForInstanceLog waits for the instance logs to be in a "terminal state" before returning.
// This function can be used to wait for an instance logs to be ready for example.
func (s *API) WaitForInstanceLog(req *WaitForInstanceLogRequest) (*InstanceLog, error) {
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
		Get: func() (interface{}, bool, error) {
			res, err := s.GetInstanceLog(&GetInstanceLogRequest{
				Region: req.Region,
			})

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

package serverless_sqldb

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

// WaitForDatabaseRequest is used by WaitForDatabase method.
type WaitForDatabaseRequest struct {
	DatabaseID    string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForDatabase waits for the database to be in a "terminal state" before returning.
// This function can be used to wait for a database to be ready for example.
func (s *API) WaitForDatabase(req *WaitForDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[DatabaseStatus]struct{}{
		DatabaseStatusReady:  {},
		DatabaseStatusError:  {},
		DatabaseStatusLocked: {},
	}

	database, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := s.GetDatabase(&GetDatabaseRequest{
				DatabaseID: req.DatabaseID,
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
		return nil, errors.Wrap(err, "waiting for database failed")
	}
	return database.(*Database), nil
}

// WaitForDatabaseBackupRequest is used by WaitForDatabase method.
type WaitForDatabaseBackupRequest struct {
	BackupID      string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForDatabaseBackup waits for the backup to be in a "terminal state" before returning.
// This function can be used to wait for a backup to be ready for example.
func (s *API) WaitForDatabaseBackup(req *WaitForDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[DatabaseBackupStatus]struct{}{
		DatabaseBackupStatusReady:  {},
		DatabaseBackupStatusError:  {},
		DatabaseBackupStatusLocked: {},
	}

	backup, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := s.GetDatabaseBackup(&GetDatabaseBackupRequest{
				BackupID: req.BackupID,
				Region:   req.Region,
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
		return nil, errors.Wrap(err, "waiting for database backup failed")
	}
	return backup.(*DatabaseBackup), nil
}

package file

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultRetryInterval = 15 * time.Second
	defaultTimeout       = 10 * time.Minute
)

type WaitForFileSystemRequest struct {
	FileSystemID  string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

func (s *API) WaitForFileSystem(req *WaitForFileSystemRequest, opts ...scw.RequestOption) (*FileSystem, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[FileSystemStatus]struct{}{
		FileSystemStatusAvailable: {},
		FileSystemStatusError:     {},
	}

	fileSystem, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetFileSystem(&GetFileSystemRequest{
				Region:       req.Region,
				FilesystemID: req.FileSystemID,
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
		return nil, errors.Wrap(err, "waiting for fileSystem failed")
	}

	return fileSystem.(*FileSystem), nil
}

package instance

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// WaitForImageRequest is used by WaitForImage method.
type WaitForSnapshotRequest struct {
	SnapshotID    string
	Zone          scw.Zone
	Timeout       *time.Duration
	RetryInterval time.Duration
}

// WaitForSnapshot wait for the snapshot to be in a "terminal state" before returning.
func (s *API) WaitForSnapshot(req *WaitForSnapshotRequest) (*Snapshot, error) {
	timeout := *req.Timeout
	if timeout == 0 {
		timeout = defaultTimeout
	}
	retryInterval := req.RetryInterval
	if retryInterval == 0 {
		retryInterval = defaultRetryInterval
	}

	terminalStatus := map[SnapshotState]struct{}{
		SnapshotStateAvailable: {},
		SnapshotStateError:     {},
	}

	snapshot, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := s.GetSnapshot(&GetSnapshotRequest{
				SnapshotID: req.SnapshotID,
				Zone:       req.Zone,
			})

			if err != nil {
				return nil, false, err
			}
			_, isTerminal := terminalStatus[res.Snapshot.State]

			return res.Snapshot, isTerminal, err
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for snapshot failed")
	}
	return snapshot.(*Snapshot), nil
}

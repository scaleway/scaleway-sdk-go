package mongodb

import (
	"context"
	"fmt"
	"sort"
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultInstanceRetryInterval = 15 * time.Second
	defaultInstanceTimeout       = 15 * time.Minute
)

type WaitForSnapshotRequest struct {
	InstanceID    string
	SnapshotID    string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForSnapshot waits for the snapshot to reach a "terminal state" before returning.
func (s *API) WaitForSnapshot(req *WaitForSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
	timeout := defaultInstanceTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultInstanceRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[SnapshotStatus]struct{}{
		SnapshotStatusReady:  {},
		SnapshotStatusError:  {},
		SnapshotStatusLocked: {},
	}

	snapshot, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			getSnapshotResponse, err := s.GetSnapshot(&GetSnapshotRequest{
				Region:     req.Region,
				SnapshotID: req.SnapshotID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[getSnapshotResponse.Status]
			return getSnapshotResponse, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for snapshot failed")
	}
	return snapshot.(*Snapshot), nil
}

func (s *API) FetchLatestEngineVersion() (*Version, error) {
	resp, err := s.ListVersions(&ListVersionsRequest{}, scw.WithContext(context.Background()))
	if err != nil {
		return nil, fmt.Errorf("error fetching MongoDB versions: %w", err)
	}

	if len(resp.Versions) == 0 {
		return nil, errors.New("no MongoDB versions found")
	}

	sort.Slice(resp.Versions, func(i, j int) bool {
		return resp.Versions[i].Version > resp.Versions[j].Version
	})

	return resp.Versions[0], nil
}

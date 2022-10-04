package redis

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

// WaitForClusterRequest is used by WaitForCluster method.
type WaitForClusterRequest struct {
	ClusterID     string
	Zone          scw.Zone
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForCluster waits for the cluster to be in a "terminal state" before returning.
// This function can be used to wait for a cluster to be ready for example.
func (s *API) WaitForCluster(req *WaitForClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[ClusterStatus]struct{}{
		ClusterStatusReady:     {},
		ClusterStatusLocked:    {},
		ClusterStatusError:     {},
		ClusterStatusSuspended: {},
	}

	cluster, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := s.GetCluster(&GetClusterRequest{
				Zone:      req.Zone,
				ClusterID: req.ClusterID,
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
		return nil, errors.Wrap(err, "waiting for cluster failed")
	}
	return cluster.(*Cluster), nil
}

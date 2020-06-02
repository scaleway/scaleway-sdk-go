package k8s

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	waitForClusterDefaultTimeout = 15 * time.Minute
	defaultRetryInterval         = 5 * time.Second
)

// WaitForClusterRequest is used by WaitForCluster method.
type WaitForClusterRequest struct {
	ClusterID     string
	Region        scw.Region
	Status        ClusterStatus
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForCluster waits for the cluster to be in a "terminal state" before returning.
func (s *API) WaitForCluster(req *WaitForClusterRequest) (*Cluster, error) {
	timeout := waitForClusterDefaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[ClusterStatus]struct{}{
		ClusterStatusReady:   {},
		ClusterStatusError:   {},
		ClusterStatusWarning: {},
		ClusterStatusLocked:  {},
		ClusterStatusDeleted: {},
	}

	cluster, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			cluster, err := s.GetCluster(&GetClusterRequest{
				ClusterID: req.ClusterID,
				Region:    req.Region,
			})
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[cluster.Status]
			return cluster, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for cluster failed")
	}
	return cluster.(*Cluster), nil
}

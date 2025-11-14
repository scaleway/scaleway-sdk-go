package k8s

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	waitForClusterDefaultTimeout = 15 * time.Minute
	defaultRetryInterval         = 5 * time.Second
)

// WaitForClusterPoolRequest is used by WaitForClusterPool method.
type WaitForClusterPoolRequest struct {
	ClusterID     string
	Region        scw.Region
	Status        ClusterStatus
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForClusterPool waits for the pool associated with a cluster to be in a "terminal state" before returning.
func (s *API) WaitForClusterPool(req *WaitForClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	timeout := waitForClusterDefaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalClusterStatus := map[ClusterStatus]struct{}{
		ClusterStatusPoolRequired: {},
		ClusterStatusReady:        {},
	}

	terminalPoolStatus := map[PoolStatus]struct{}{
		PoolStatusReady:   {},
		PoolStatusWarning: {},
	}

	optsWithAllPages := append(opts, scw.WithAllPages()) //nolint:gocritic

	cluster, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			cluster, err := s.GetCluster(&GetClusterRequest{
				ClusterID: req.ClusterID,
				Region:    req.Region,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalClusterStatus[cluster.Status]
			if !isTerminal {
				return cluster, false, nil
			}

			pools, err := s.ListPools(&ListPoolsRequest{
				Region:    req.Region,
				ClusterID: req.ClusterID,
			}, optsWithAllPages...)
			if err != nil {
				return nil, false, err
			}

			for _, pool := range pools.Pools {
				_, isTerminal = terminalPoolStatus[pool.Status]
				if !isTerminal {
					return cluster, false, nil
				}
			}

			return cluster, true, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for cluster failed")
	}

	return cluster.(*Cluster), nil
}

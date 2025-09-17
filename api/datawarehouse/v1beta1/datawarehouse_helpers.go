package datawarehouse

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	waitForDeploymentDefaultTimeout = 30 * time.Minute
	defaultRetryIntervalDW          = 5 * time.Second
)

type WaitForDeploymentRequest struct {
	DeploymentID   string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForDeployment waits until deployment reaches a terminal state.
func (s *API) WaitForDeployment(req *WaitForDeploymentRequest, opts ...scw.RequestOption) (*Deployment, error) {
	timeout := waitForDeploymentDefaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryIntervalDW
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[DeploymentStatus]struct{}{
		DeploymentStatusReady:   {},
		DeploymentStatusError:   {},
		DeploymentStatusStopped: {},
	}

	deployment, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := s.GetDeployment(&GetDeploymentRequest{
				DeploymentID: req.DeploymentID,
				Region:      req.Region,
			}, opts...)

			if err != nil {
				// Don't consider 404 as an error, deployment might not be created yet
				if errors.Is(err, &errors.ResponseError{StatusCode: 404}) {
					return nil, false, nil
				}
				return nil, false, err
			}

			_, isTerminal := terminalStatus[res.Status]
			return res, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, err
	}
	return deployment.(*Deployment), nil
}
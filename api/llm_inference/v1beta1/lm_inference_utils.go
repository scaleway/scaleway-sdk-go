package llm_inference

import (
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"time"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultRetryInterval = 2 * time.Second
	defaultTimeout       = 5 * time.Minute
)

type WaitForDeploymentRequest struct {
	DeploymentId  string
	Region        scw.Region
	Status        DeploymentStatus
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

func (s *API) WaitForDeployment(req *WaitForDeploymentRequest, opts ...scw.RequestOption) (*Deployment, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[DeploymentStatus]struct{}{
		DeploymentStatusReady:  {},
		DeploymentStatusError:  {},
		DeploymentStatusLocked: {},
	}

	deployment, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			deployment, err := s.GetDeployment(&GetDeploymentRequest{
				Region:       req.Region,
				DeploymentID: req.DeploymentId,
			}, opts...)
			if err != nil {
				return nil, false, err
			}
			_, isTerminal := terminalStatus[deployment.Status]
			return deployment, isTerminal, nil
		},
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
		Timeout:          timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for deployment failed")
	}
	return deployment.(*Deployment), nil
}

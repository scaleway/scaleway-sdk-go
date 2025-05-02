package inference

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultRetryInterval = 15 * time.Second
	defaultTimeout       = 30 * time.Minute
)

type WaitForDeploymentRequest struct {
	DeploymentID  string
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
				DeploymentID: req.DeploymentID,
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

type WaitForModelRequest struct {
	ModelID       string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

func (s *API) WaitForModel(req *WaitForModelRequest, opts ...scw.RequestOption) (*Model, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[ModelStatus]struct{}{
		ModelStatusReady: {},
		ModelStatusError: {},
	}

	model, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			model, err := s.GetModel(&GetModelRequest{
				Region:  req.Region,
				ModelID: req.ModelID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}
			_, isTerminal := terminalStatus[model.Status]
			return model, isTerminal, nil
		},
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
		Timeout:          timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for model failed")
	}

	return model.(*Model), nil
}

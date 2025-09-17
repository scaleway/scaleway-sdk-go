package datawarehouse

import (
"time"

"github.com/scaleway/scaleway-sdk-go/errors"
"github.com/scaleway/scaleway-sdk-go/internal/async"
"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
waitForDeploymentDefaultTimeout = 30 * time.Minute
defaultRetryIntervalDW           = 5 * time.Second
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
DeploymentStatusError: {},
DeploymentStatusReady: {},
DeploymentStatusLocked: {},
}

deployment, err := async.WaitSync(&async.WaitSyncConfig{
Get: func() (any, bool, error) {
deployment, err := s.GetDeployment(&GetDeploymentRequest{
DeploymentID: req.DeploymentID,
Region:       req.Region,
}, opts...)
if err != nil {
return nil, false, err
}

_, isTerminal := terminalStatus[deployment.Status]
return deployment, isTerminal, nil
},
Timeout:          timeout,
IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
})
if err != nil {
return nil, errors.Wrap(err, "waiting for Deployment failed")
}

return deployment.(*Deployment), nil
}

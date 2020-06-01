package registry

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultTimeout       = 5 * time.Minute
	defaultRetryInterval = 15 * time.Second
)

// WaitForNamespaceRequest is used by WaitForNamespace method
type WaitForNamespaceRequest struct {
	NamespaceID   string
	Region        scw.Region
	Timeout       time.Duration
	RetryInterval time.Duration
}

// WaitForNamespace wait for the namespace to be in a "terminal state" before returning.
// This function can be used to wait for a namespace to be ready for example.
func (s *API) WaitForNamespace(req *WaitForNamespaceRequest) (*Namespace, error) {
	if req.Timeout == 0 {
		req.Timeout = defaultTimeout
	}
	if req.RetryInterval == 0 {
		req.RetryInterval = defaultRetryInterval
	}

	terminalStatus := map[NamespaceStatus]struct{}{
		NamespaceStatusReady:   {},
		NamespaceStatusLocked:  {},
		NamespaceStatusError:   {},
		NamespaceStatusUnknown: {},
	}

	namespace, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			ns, err := s.GetNamespace(&GetNamespaceRequest{
				Region:      req.Region,
				NamespaceID: req.NamespaceID,
			})
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[ns.Status]

			return ns, isTerminal, err
		},
		Timeout:          req.Timeout,
		IntervalStrategy: async.LinearIntervalStrategy(req.RetryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for namespace failed")
	}
	return namespace.(*Namespace), nil
}

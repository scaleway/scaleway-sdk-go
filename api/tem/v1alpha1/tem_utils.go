package tem

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

// WaitForDomainRequest is used by WaitForDomain method
type WaitForDomainRequest struct {
	DomainID      string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForDomain wait for the domain to be in a "terminal state" before returning.
// This function can be used to wait for a domain to be checked for example.
func (s *API) WaitForDomain(req *WaitForDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[DomainStatus]struct{}{
		DomainStatusChecked:   {},
		DomainStatusUnchecked: {},
		DomainStatusInvalid:   {},
		DomainStatusLocked:    {},
		DomainStatusRevoked:   {},
		DomainStatusUnknown:   {},
	}

	domain, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			img, err := s.GetDomain(&GetDomainRequest{
				Region:   req.Region,
				DomainID: req.DomainID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[img.Status]

			return img, isTerminal, err
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for domain failed")
	}
	return domain.(*Domain), nil
}

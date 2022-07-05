package function

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	waitForNamespaceDefaultTimeout = 15 * time.Minute
	waitForCronDefaultTimeout      = 15 * time.Minute
	waitForDomainDefaultTimeout    = 15 * time.Minute
	waitForFunctionDefaultTimeout  = 15 * time.Minute
	defaultRetryInterval           = 5 * time.Second
)

// WaitForNamespaceRequest is used by WaitForNamespace method.
type WaitForNamespaceRequest struct {
	NamespaceID   string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForNamespace waits for the Namespace to be in a ready state before returning.
func (s *API) WaitForNamespace(req *WaitForNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	timeout := waitForNamespaceDefaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[NamespaceStatus]struct{}{
		NamespaceStatusError:  {},
		NamespaceStatusReady:  {},
		NamespaceStatusLocked: {},
	}

	namespace, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			namespace, err := s.GetNamespace(&GetNamespaceRequest{
				NamespaceID: req.NamespaceID,
				Region:      req.Region,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[namespace.Status]
			return namespace, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for Namespace failed")
	}

	return namespace.(*Namespace), nil
}

// WaitForFunctionRequest is used by WaitForNamespace method.
type WaitForFunctionRequest struct {
	FunctionID    string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForFunction waits for the Function to be in a ready state before returning.
func (s *API) WaitForFunction(req *WaitForFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	timeout := waitForFunctionDefaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[FunctionStatus]struct{}{
		FunctionStatusCreated: {},
		FunctionStatusError:   {},
		FunctionStatusLocked:  {},
		FunctionStatusReady:   {},
	}

	function, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			function, err := s.GetFunction(&GetFunctionRequest{
				FunctionID: req.FunctionID,
				Region:     req.Region,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[function.Status]
			return function, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for function failed")
	}

	return function.(*Function), nil
}

// WaitForCronRequest is used by WaitForNamespace method.
type WaitForCronRequest struct {
	CronID        string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForCron waits for the Cron to be in a ready state before returning.
func (s *API) WaitForCron(req *WaitForCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	timeout := waitForCronDefaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[CronStatus]struct{}{
		CronStatusError:  {},
		CronStatusReady:  {},
		CronStatusLocked: {},
	}

	cron, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			cron, err := s.GetCron(&GetCronRequest{
				CronID: req.CronID,
				Region: req.Region,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[cron.Status]
			return cron, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for Cron failed")
	}

	return cron.(*Cron), nil
}

// WaitForDomainRequest waits for the Domain to be in a ready state before returning.
type WaitForDomainRequest struct {
	DomainID      string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForDomain waits for the Domain to be in a ready state before returning.
func (s *API) WaitForDomain(req *WaitForDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	timeout := waitForDomainDefaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[DomainStatus]struct{}{
		DomainStatusError: {},
		DomainStatusReady: {},
	}

	domain, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			domain, err := s.GetDomain(&GetDomainRequest{
				DomainID: req.DomainID,
				Region:   req.Region,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[domain.Status]
			return domain, isTerminal, nil
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for Domain failed")
	}

	return domain.(*Domain), nil
}

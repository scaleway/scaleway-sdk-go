package lb

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultRetryInterval = 2 * time.Second
	defaultTimeout       = 5 * time.Minute
)

// WaitForLBRequest is used by WaitForLb method.
type WaitForLBRequest struct {
	LBID          string
	Region        scw.Region
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForLb waits for the lb to be in a "terminal state" before returning.
// This function can be used to wait for a lb to be ready for example.
func (s *API) WaitForLb(req *WaitForLBRequest, opts ...scw.RequestOption) (*LB, error) {
	return waitForLb(req.Timeout, req.RetryInterval, func() (*LB, error) {
		return s.GetLB(&GetLBRequest{
			Region: req.Region,
			LBID:   req.LBID,
		})
	})
}

// ZonedAPIWaitForLBRequest is used by WaitForLb method.
type ZonedAPIWaitForLBRequest struct {
	LBID          string
	Zone          scw.Zone
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForLb waits for the lb to be in a "terminal state" before returning.
// This function can be used to wait for a lb to be ready for example.
func (s *ZonedAPI) WaitForLb(req *ZonedAPIWaitForLBRequest, opts ...scw.RequestOption) (*LB, error) {
	return waitForLb(req.Timeout, req.RetryInterval, func() (*LB, error) {
		return s.GetLB(&ZonedAPIGetLBRequest{
			Zone: req.Zone,
			LBID: req.LBID,
		})
	})
}

func waitForLb(timeout *time.Duration, retryInterval *time.Duration, getLB func() (*LB, error)) (*LB, error) {
	t := defaultTimeout
	if timeout != nil {
		t = *timeout
	}
	r := defaultRetryInterval
	if retryInterval != nil {
		r = *retryInterval
	}

	terminalStatus := map[LBStatus]struct{}{
		LBStatusReady:   {},
		LBStatusStopped: {},
		LBStatusError:   {},
		LBStatusLocked:  {},
	}

	lb, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := getLB()

			if err != nil {
				return nil, false, err
			}
			_, isTerminal := terminalStatus[res.Status]

			return res, isTerminal, nil
		},
		Timeout:          t,
		IntervalStrategy: async.LinearIntervalStrategy(r),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for lb failed")
	}
	return lb.(*LB), nil
}

package registry

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// WaitForTagRequest is used by WaitForTag method
type WaitForTagRequest struct {
	TagID         string
	Region        scw.Region
	Timeout       time.Duration
	RetryInterval time.Duration
}

// WaitForTag wait for the tag to be in a "terminal state" before returning.
// This function can be used to wait for a tag to be ready for example.
func (s *API) WaitForTag(req *WaitForTagRequest) (*Tag, error) {
	if req.Timeout == 0 {
		req.Timeout = defaultTimeout
	}
	if req.RetryInterval == 0 {
		req.RetryInterval = defaultRetryInterval
	}

	terminalStatus := map[TagStatus]struct{}{
		TagStatusReady:   {},
		TagStatusLocked:  {},
		TagStatusError:   {},
		TagStatusUnknown: {},
	}

	tag, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			t, err := s.GetTag(&GetTagRequest{
				Region: req.Region,
				TagID:  req.TagID,
			})
			if err != nil {
				return nil, false, err
			}

			_, isTerminal := terminalStatus[t.Status]

			return t, isTerminal, err
		},
		Timeout:          req.Timeout,
		IntervalStrategy: async.LinearIntervalStrategy(req.RetryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for tag failed")
	}
	return tag.(*Tag), nil
}

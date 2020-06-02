package instance

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// WaitForImageRequest is used by WaitForImage method.
type WaitForImageRequest struct {
	ImageID       string
	Zone          scw.Zone
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForImage wait for the image to be in a "terminal state" before returning.
func (s *API) WaitForImage(req *WaitForImageRequest) (*Image, error) {
	timeout := *req.Timeout
	if timeout == 0 {
		timeout = defaultTimeout
	}
	retryInterval := defaultRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}

	terminalStatus := map[ImageState]struct{}{
		ImageStateAvailable: {},
		ImageStateError:     {},
	}

	image, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := s.GetImage(&GetImageRequest{
				ImageID: req.ImageID,
				Zone:    req.Zone,
			})

			if err != nil {
				return nil, false, err
			}
			_, isTerminal := terminalStatus[res.Image.State]

			return res.Image, isTerminal, err
		},
		Timeout:          timeout,
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for image failed")
	}
	return image.(*Image), nil
}

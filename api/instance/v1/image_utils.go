package instance

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// WaitForServerRequest is used by WaitForServer method.
type WaitForImageRequest struct {
	ImageID string
	Zone    scw.Zone
	Timeout time.Duration
}

// WaitForServer wait for the server to be in a "terminal state" before returning.
// This function can be used to wait for a server to be started for example.
func (s *API) WaitForImage(req *WaitForImageRequest) (*Image, error) {
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
		Timeout:          req.Timeout,
		IntervalStrategy: async.LinearIntervalStrategy(RetryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for image failed")
	}
	return image.(*Image), nil
}

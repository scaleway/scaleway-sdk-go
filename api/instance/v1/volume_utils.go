package instance

import (
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

// WaitForImageRequest is used by WaitForImage method.
type WaitForVolumeRequest struct {
	VolumeID string
	Zone     scw.Zone
	Timeout  time.Duration
}

// WaitForSnapshot wait for the snapshot to be in a "terminal state" before returning.
func (s *API) WaitForVolume(req *WaitForVolumeRequest) (*Volume, error) {
	if req.Timeout == 0 {
		req.Timeout = defaultTimeout
	}

	terminalStatus := map[VolumeState]struct{}{
		VolumeStateAvailable: {},
		VolumeStateError:     {},
	}

	volume, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := s.GetVolume(&GetVolumeRequest{
				VolumeID: req.VolumeID,
				Zone:     req.Zone,
			})

			if err != nil {
				return nil, false, err
			}
			_, isTerminal := terminalStatus[res.Volume.State]

			return res.Volume, isTerminal, err
		},
		Timeout:          req.Timeout,
		IntervalStrategy: async.LinearIntervalStrategy(RetryInterval),
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for volume failed")
	}
	return volume.(*Volume), nil
}

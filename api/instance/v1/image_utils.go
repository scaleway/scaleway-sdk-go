package instance

import (
	"fmt"
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
func (s *API) WaitForImage(req *WaitForImageRequest, opts ...scw.RequestOption) (*Image, error) {
	timeout := defaultTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
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
			}, opts...)

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

type UpdateImageRequest struct {
	Zone                scw.Zone
	ImageID             string
	Name                *string   `json:"name,omitempty"`
	RootVolumeID        string    `json:"root_volume_id,omitempty"`
	Arch                Arch      `json:"architecture,omitempty"`
	DefaultBootscriptID string    `json:"default_bootscript_id,omitempty"`
	ExtraVolumesIDs     []string  `json:"additional_volume_ids,omitempty"`
	Public              bool      `json:"public,omitempty"`
	Tags                *[]string `json:"tags,omitempty"`
}

type UpdateImageResponse struct {
	Image *Image
}

func (s *API) UpdateImage(req *UpdateImageRequest, opts ...scw.RequestOption) (*UpdateImageResponse, error) {
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}
	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	getImageResponse, err := s.GetImage(&GetImageRequest{
		Zone:    req.Zone,
		ImageID: req.ImageID,
	}, opts...)
	if err != nil {
		return nil, err
	}

	setRequest := &SetImageRequest{
		Zone:              getImageResponse.Image.Zone,
		ID:                getImageResponse.Image.ID,
		Name:              getImageResponse.Image.Name,
		Arch:              getImageResponse.Image.Arch,
		CreationDate:      getImageResponse.Image.CreationDate,
		ModificationDate:  getImageResponse.Image.ModificationDate,
		DefaultBootscript: getImageResponse.Image.DefaultBootscript,
		ExtraVolumes:      getImageResponse.Image.ExtraVolumes,
		FromServer:        getImageResponse.Image.FromServer,
		Organization:      getImageResponse.Image.Organization,
		Public:            getImageResponse.Image.Public,
		RootVolume:        getImageResponse.Image.RootVolume,
		State:             getImageResponse.Image.State,
		Project:           getImageResponse.Image.Project,
		Tags:              &getImageResponse.Image.Tags,
	}

	// Override the values that need to be updated
	if req.Name != nil {
		setRequest.Name = *req.Name
	}
	if req.Tags != nil {
		setRequest.Tags = req.Tags
	}
	if req.RootVolumeID != "" {
		setRequest.RootVolume = req.RootVolumeID
	}
	if req.Arch != "" {
		setRequest.Arch = req.Arch
	}
	if req.DefaultBootscriptID != "" {
		setRequest.DefaultBootscript = req.DefaultBootscriptID
	}
	if req.ExtraVolumesIDs != nil {
		setRequest.ExtraVolumes = req.ExtraVolumesIDs
	}
	if req.Public != setRequest.Public {
		setRequest.Public = req.Public
	}

	setRes, err := s.setImage(setRequest, opts...)
	if err != nil {
		return nil, err
	}

	return &UpdateImageResponse{
		Image: setRes.Image,
	}, nil
}

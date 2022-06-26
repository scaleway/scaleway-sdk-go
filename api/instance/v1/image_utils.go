package instance

import (
	"fmt"
	"net/http"
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
	Zone              scw.Zone           `json:"zone"`
	ImageID           string             `json:"id"`
	Name              *string            `json:"name,omitempty"`
	Arch              Arch               `json:"arch,omitempty"`
	CreationDate      *time.Time         `json:"creation_date"`
	ModificationDate  *time.Time         `json:"modification_date"`
	DefaultBootscript *Bootscript        `json:"default_bootscript,omitempty"`
	ExtraVolumes      map[string]*Volume `json:"extra_volumes"`
	FromServer        string             `json:"from_server,omitempty"`
	Organization      string             `json:"organization"`
	Public            bool               `json:"public"`
	RootVolume        *VolumeSummary     `json:"root_volume,omitempty"`
	State             ImageState         `json:"state"`
	Project           string             `json:"project"`
	Tags              *[]string          `json:"tags,omitempty"`
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

	// Fill in computed fields
	req.FromServer = getImageResponse.Image.FromServer
	req.Organization = getImageResponse.Image.Organization
	req.State = getImageResponse.Image.State
	req.Project = getImageResponse.Image.Project
	req.CreationDate = getImageResponse.Image.CreationDate
	req.ModificationDate = getImageResponse.Image.ModificationDate

	// Ensure that no field is empty in request
	if req.Name == nil {
		req.Name = &getImageResponse.Image.Name
	}
	if req.Tags == nil {
		req.Tags = &getImageResponse.Image.Tags
	}
	if req.RootVolume == nil {
		req.RootVolume = getImageResponse.Image.RootVolume
	}
	if req.Arch == "" {
		req.Arch = getImageResponse.Image.Arch
	}
	if req.DefaultBootscript == nil {
		req.DefaultBootscript = getImageResponse.Image.DefaultBootscript
	}
	if req.ExtraVolumes == nil {
		req.ExtraVolumes = getImageResponse.Image.ExtraVolumes
	}

	// Here starts the equivalent of the private setImage function that is not usable because the json tags and types
	// are not compatible with what the instance API expects

	if req.Project == "" {
		defaultProject, _ := s.client.GetDefaultProjectID()
		req.Project = defaultProject
	}
	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}
	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}
	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.ImageID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateImageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}

	return &resp, nil
}

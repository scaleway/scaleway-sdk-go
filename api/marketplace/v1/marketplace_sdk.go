// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package marketplace

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/utils"
)

// always import dependencies
var (
	_ fmt.Stringer
	_ json.Unmarshaler
	_ url.URL
	_ net.IP
	_ http.Header
	_ bytes.Reader
	_ time.Time

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ utils.File
	_ = parameter.AddToQuery
)

// Api: marketplace API
type Api struct {
	client *scw.Client
}

// NewApi returns a Api object from a Scaleway client.
func NewApi(client *scw.Client) *Api {
	return &Api{
		client: client,
	}
}

type GetImageResponse struct {
	Image *Image `json:"image,omitempty"`
}

type GetServiceInfoResponse struct {
	Api string `json:"api,omitempty"`

	Description string `json:"description,omitempty"`

	Version string `json:"version,omitempty"`
}

type GetVersionResponse struct {
	Version *Version `json:"version,omitempty"`
}

type Image struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Logo string `json:"logo,omitempty"`

	Categories []string `json:"categories,omitempty"`

	Organization *Organization `json:"organization,omitempty"`

	ValidUntil time.Time `json:"valid_until,omitempty"`

	CreationDate time.Time `json:"creation_date,omitempty"`

	ModificationDate time.Time `json:"modification_date,omitempty"`

	Versions []*Version `json:"versions,omitempty"`

	CurrentPublicVersion string `json:"current_public_version,omitempty"`
}

type ListImagesResponse struct {
	Images []*Image `json:"images,omitempty"`
}

type ListVersionsResponse struct {
	Versions []*Version `json:"versions,omitempty"`
}

type LocalImage struct {
	Id string `json:"id,omitempty"`

	Arch string `json:"arch,omitempty"`

	Zone utils.Zone `json:"zone,omitempty"`

	CompatibleCommercialTypes []string `json:"compatible_commercial_types,omitempty"`
}

type Organization struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

type Version struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	CreationDate time.Time `json:"creation_date,omitempty"`

	ModificationDate time.Time `json:"modification_date,omitempty"`

	LocalImages []*LocalImage `json:"local_images,omitempty"`
}

// Service Api

type GetServiceInfoRequest struct {
}

func (s *Api) GetServiceInfo(req *GetServiceInfoRequest) (*GetServiceInfoResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/marketplace/v1",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetServiceInfoResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListImagesRequest struct {
	PerPage *int32 `json:"-"`

	Page *int32 `json:"-"`
}

func (s *Api) ListImages(req *ListImagesRequest) (*ListImagesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/marketplace/v1/images",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListImagesResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetImageRequest struct {
	ImageId string `json:"-"`
}

func (s *Api) GetImage(req *GetImageRequest) (*GetImageResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/marketplace/v1/images/" + fmt.Sprint(req.ImageId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetImageResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListVersionsRequest struct {
	ImageId string `json:"-"`
}

func (s *Api) ListVersions(req *ListVersionsRequest) (*ListVersionsResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/marketplace/v1/images/" + fmt.Sprint(req.ImageId) + "/versions",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListVersionsResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetVersionRequest struct {
	ImageId string `json:"-"`

	VersionId string `json:"-"`
}

func (s *Api) GetVersion(req *GetVersionRequest) (*GetVersionResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/marketplace/v1/images/" + fmt.Sprint(req.ImageId) + "/versions/" + fmt.Sprint(req.VersionId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetVersionResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

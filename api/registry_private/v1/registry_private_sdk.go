// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package registry_private provides methods and message types of the registry_private v1 API.
package registry_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/scw"
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
	_ = strings.Join

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ scw.File
	_ = parameter.AddToQuery
	_ = namegenerator.GetRandomName
)

// GetImageStatsResponse: get image stats response.
type GetImageStatsResponse struct {
	ID string `json:"id"`

	Size uint64 `json:"size"`

	PullVolume uint64 `json:"pull_volume"`

	Tags uint32 `json:"tags"`
}

// GetNamespaceStatsResponse: get namespace stats response.
type GetNamespaceStatsResponse struct {
	ID string `json:"id"`

	Size uint64 `json:"size"`

	PullVolume uint64 `json:"pull_volume"`

	Images uint32 `json:"images"`
}

// GetPricingResponse: get pricing response.
type GetPricingResponse struct {
	MonthlyStorage *scw.Money `json:"monthly_storage"`

	Push *scw.Money `json:"push"`

	Pull *scw.Money `json:"pull"`
}

// GetStatsResponse: get stats response.
type GetStatsResponse struct {
	Namespaces uint32 `json:"namespaces"`

	Images uint32 `json:"images"`

	Size uint64 `json:"size"`

	PullVolume uint64 `json:"pull_volume"`
}

// StatsAPIGetImageStatsRequest: stats api get image stats request.
type StatsAPIGetImageStatsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ImageID string `json:"-"`

	StartDate *time.Time `json:"-"`

	EndDate *time.Time `json:"-"`
}

// StatsAPIGetNamespaceStatsRequest: stats api get namespace stats request.
type StatsAPIGetNamespaceStatsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	NamespaceID string `json:"-"`

	StartDate *time.Time `json:"-"`

	EndDate *time.Time `json:"-"`
}

// StatsAPIGetPricingRequest: stats api get pricing request.
type StatsAPIGetPricingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// StatsAPIGetStatsRequest: stats api get stats request.
type StatsAPIGetStatsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	OrganizationID string `json:"-"`

	StartDate *time.Time `json:"-"`

	EndDate *time.Time `json:"-"`

	ProjectID *string `json:"-"`
}

type StatsAPI struct {
	client *scw.Client
}

// NewStatsAPI returns a StatsAPI object from a Scaleway client.
func NewStatsAPI(client *scw.Client) *StatsAPI {
	return &StatsAPI{
		client: client,
	}
}

// GetStats:
func (s *StatsAPI) GetStats(req *StatsAPIGetStatsRequest, opts ...scw.RequestOption) (*GetStatsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/registry-private/v1/regions/" + fmt.Sprint(req.Region) + "/stats",
		Query:  query,
	}

	var resp GetStatsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPricing:
func (s *StatsAPI) GetPricing(req *StatsAPIGetPricingRequest, opts ...scw.RequestOption) (*GetPricingResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/registry-private/v1/regions/" + fmt.Sprint(req.Region) + "/pricing",
	}

	var resp GetPricingResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNamespaceStats:
func (s *StatsAPI) GetNamespaceStats(req *StatsAPIGetNamespaceStatsRequest, opts ...scw.RequestOption) (*GetNamespaceStatsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/registry-private/v1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "/stats",
		Query:  query,
	}

	var resp GetNamespaceStatsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImageStats:
func (s *StatsAPI) GetImageStats(req *StatsAPIGetImageStatsRequest, opts ...scw.RequestOption) (*GetImageStatsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/registry-private/v1/regions/" + fmt.Sprint(req.Region) + "/images/" + fmt.Sprint(req.ImageID) + "/stats",
		Query:  query,
	}

	var resp GetImageStatsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

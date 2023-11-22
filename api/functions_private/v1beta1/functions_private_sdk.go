// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package functions_private provides methods and message types of the functions_private v1beta1 API.
package functions_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/functions_private/v1beta1/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/functions_private/v1beta1/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/functions_private/v1beta1/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/functions_private/v1beta1/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/functions_private/v1beta1/scw"
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

// Resource: resource.
type Resource struct {
	MemoryLimit uint32 `json:"memory_limit"`

	CPULimit uint32 `json:"cpu_limit"`
}

// ApplicationMetricsResponse: application metrics response.
type ApplicationMetricsResponse struct {
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

// ConsoleAPIGetApplicationMetricsRequest: console api get application metrics request.
type ConsoleAPIGetApplicationMetricsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Precisely one of ApplicationID must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	StartDate *time.Time `json:"-"`

	EndDate *time.Time `json:"-"`

	MetricName *string `json:"-"`
}

// ConsoleAPIGetContainerLimitsRequest: console api get container limits request.
type ConsoleAPIGetContainerLimitsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// ConsoleAPIGetFunctionLimitsRequest: console api get function limits request.
type ConsoleAPIGetFunctionLimitsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// ConsoleAPIGetServiceInfoRequest: console api get service info request.
type ConsoleAPIGetServiceInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// ConsoleAPIListContainerResourcesRequest: console api list container resources request.
type ConsoleAPIListContainerResourcesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// ConsoleAPIListFunctionResourcesRequest: console api list function resources request.
type ConsoleAPIListFunctionResourcesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// GetContainerLimitsResponse: get container limits response.
type GetContainerLimitsResponse struct {
	MaxImageArchive scw.Size `json:"max_image_archive"`

	MaxImageSize scw.Size `json:"max_image_size"`

	MaxEnvVarsSize uint32 `json:"max_env_vars_size"`

	MaxEnvVarsNumber uint32 `json:"max_env_vars_number"`

	MaxRequestPayload scw.Size `json:"max_request_payload"`

	// Deprecated
	MaxScale *uint32 `json:"max_scale,omitempty"`

	EphemeralDiskSize uint32 `json:"ephemeral_disk_size"`

	MaxConcurrency uint32 `json:"max_concurrency"`

	MaxInvocationRate uint32 `json:"max_invocation_rate"`

	ScaleDownTime *scw.Duration `json:"scale_down_time"`

	ScaleZeroTime *scw.Duration `json:"scale_zero_time"`

	MaxRequestTimeout *scw.Duration `json:"max_request_timeout"`

	ForbiddenPorts []uint32 `json:"forbidden_ports"`

	DefaultMaxScale uint32 `json:"default_max_scale"`

	MinMinScale uint32 `json:"min_min_scale"`

	MaxMinScale uint32 `json:"max_min_scale"`

	MinMaxScale uint32 `json:"min_max_scale"`

	MaxMaxScale uint32 `json:"max_max_scale"`
}

// GetFunctionLimitsResponse: get function limits response.
type GetFunctionLimitsResponse struct {
	MaxZipSize scw.Size `json:"max_zip_size"`

	MaxCodeSize scw.Size `json:"max_code_size"`

	MaxEnvVarsSize uint32 `json:"max_env_vars_size"`

	MaxEnvVarsNumber uint32 `json:"max_env_vars_number"`

	MaxRequestPayload scw.Size `json:"max_request_payload"`

	// Deprecated
	MaxScale *uint32 `json:"max_scale,omitempty"`

	EphemeralDiskSize uint32 `json:"ephemeral_disk_size"`

	MaxConcurrency uint32 `json:"max_concurrency"`

	MaxInvocationRate uint32 `json:"max_invocation_rate"`

	ScaleDownTime *scw.Duration `json:"scale_down_time"`

	ScaleZeroTime *scw.Duration `json:"scale_zero_time"`

	MaxRequestTimeout *scw.Duration `json:"max_request_timeout"`

	DefaultMaxScale uint32 `json:"default_max_scale"`

	MinMinScale uint32 `json:"min_min_scale"`

	MaxMinScale uint32 `json:"max_min_scale"`

	MinMaxScale uint32 `json:"min_max_scale"`

	MaxMaxScale uint32 `json:"max_max_scale"`
}

// ListContainerResourcesResponse: list container resources response.
type ListContainerResourcesResponse struct {
	Resources []*Resource `json:"resources"`
}

// ListFunctionResourcesResponse: list function resources response.
type ListFunctionResourcesResponse struct {
	Resources []*Resource `json:"resources"`
}

type ConsoleAPI struct {
	client *scw.Client
}

// NewConsoleAPI returns a ConsoleAPI object from a Scaleway client.
func NewConsoleAPI(client *scw.Client) *ConsoleAPI {
	return &ConsoleAPI{
		client: client,
	}
}

// GetServiceInfo:
func (s *ConsoleAPI) GetServiceInfo(req *ConsoleAPIGetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
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
		Path:   "/functions-private/v1beta1/regions/" + fmt.Sprint(req.Region) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFunctionResources:
func (s *ConsoleAPI) ListFunctionResources(req *ConsoleAPIListFunctionResourcesRequest, opts ...scw.RequestOption) (*ListFunctionResourcesResponse, error) {
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
		Path:   "/functions-private/v1beta1/regions/" + fmt.Sprint(req.Region) + "/function-resources",
	}

	var resp ListFunctionResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListContainerResources:
func (s *ConsoleAPI) ListContainerResources(req *ConsoleAPIListContainerResourcesRequest, opts ...scw.RequestOption) (*ListContainerResourcesResponse, error) {
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
		Path:   "/functions-private/v1beta1/regions/" + fmt.Sprint(req.Region) + "/container-resources",
	}

	var resp ListContainerResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetApplicationMetrics:
func (s *ConsoleAPI) GetApplicationMetrics(req *ConsoleAPIGetApplicationMetricsRequest, opts ...scw.RequestOption) (*ApplicationMetricsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)
	parameter.AddToQuery(query, "metric_name", req.MetricName)
	parameter.AddToQuery(query, "application_id", req.ApplicationID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions-private/v1beta1/regions/" + fmt.Sprint(req.Region) + "/metrics",
		Query:  query,
	}

	var resp ApplicationMetricsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFunctionLimits:
func (s *ConsoleAPI) GetFunctionLimits(req *ConsoleAPIGetFunctionLimitsRequest, opts ...scw.RequestOption) (*GetFunctionLimitsResponse, error) {
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
		Path:   "/functions-private/v1beta1/regions/" + fmt.Sprint(req.Region) + "/function-limits",
	}

	var resp GetFunctionLimitsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetContainerLimits:
func (s *ConsoleAPI) GetContainerLimits(req *ConsoleAPIGetContainerLimitsRequest, opts ...scw.RequestOption) (*GetContainerLimitsResponse, error) {
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
		Path:   "/functions-private/v1beta1/regions/" + fmt.Sprint(req.Region) + "/container-limits",
	}

	var resp GetContainerLimitsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

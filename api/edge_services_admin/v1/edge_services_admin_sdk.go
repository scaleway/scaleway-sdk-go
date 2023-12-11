// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package edge_services_admin provides methods and message types of the edge_services_admin v1 API.
package edge_services_admin

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

// NetworkFiltersRule: network filters rule.
type NetworkFiltersRule struct {
	Filter string `json:"filter"`

	Exclude []string `json:"exclude"`
}

// NetworkFilters: network filters.
type NetworkFilters struct {
	Rules []*NetworkFiltersRule `json:"rules"`
}

// GetServiceInfoRequest: get service info request.
type GetServiceInfoRequest struct {
}

// SetNetworkFiltersRequest: set network filters request.
type SetNetworkFiltersRequest struct {
	Filters *NetworkFilters `json:"filters,omitempty"`
}

type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// GetServiceInfo:
func (s *API) GetServiceInfo(req *GetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services-admin/v1",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNetworkFilters:
func (s *API) GetNetworkFilters(opts ...scw.RequestOption) (*NetworkFilters, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services-admin/v1/network-filters",
	}

	var resp NetworkFilters

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetNetworkFilters:
func (s *API) SetNetworkFilters(req *SetNetworkFiltersRequest, opts ...scw.RequestOption) (*NetworkFilters, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/edge-services-admin/v1/network-filters",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp NetworkFilters

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package vpc_gw_private provides methods and message types of the vpc_gw_private v1 API.
package vpc_gw_private

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
	vpc_v1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/vpc/v1"
	vpcgw_v1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/vpcgw/v1"
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

// IPWithGateway: ip with gateway.
type IPWithGateway struct {
	IP *vpcgw_v1.IP `json:"ip"`

	Gateway *vpcgw_v1.Gateway `json:"gateway"`
}

// ConsoleAPIGetGatewayRequest: console api get gateway request.
type ConsoleAPIGetGatewayRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GatewayID: ID of the gateway to fetch.
	GatewayID string `json:"-"`
}

// ConsoleAPIListGatewaysRequest: console api list gateways request.
type ConsoleAPIListGatewaysRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy vpcgw_v1.ListGatewaysRequestOrderBy `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: gateways per page.
	PageSize *uint32 `json:"-"`

	// OrganizationID: include only gateways in this Organization.
	OrganizationID *string `json:"-"`

	// ProjectID: include only gateways in this Project.
	ProjectID *string `json:"-"`

	// Name: filter for gateways which have this search term in their name.
	Name *string `json:"-"`

	// Tags: filter for gateways with these tags.
	Tags []string `json:"-"`

	// Type: filter for gateways of this type.
	Type *string `json:"-"`

	// Status: filter for gateways with this current status. Use `unknown` to include all statuses.
	// Default value: unknown
	Status vpcgw_v1.GatewayStatus `json:"-"`

	// PrivateNetworkID: filter for gateways attached to this Private nNetwork.
	PrivateNetworkID *string `json:"-"`
}

// ConsoleAPIListIPsRequest: console api list i ps request.
type ConsoleAPIListIPsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy vpcgw_v1.ListIPsRequestOrderBy `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: IP addresses per page.
	PageSize *uint32 `json:"-"`

	// OrganizationID: filter for IP addresses in this Organization.
	OrganizationID *string `json:"-"`

	// ProjectID: filter for IP addresses in this Project.
	ProjectID *string `json:"-"`

	// Tags: filter for IP addresses with these tags.
	Tags []string `json:"-"`

	// Reverse: filter for IP addresses that have a reverse containing this string.
	Reverse *string `json:"-"`

	// IsFree: filter based on whether the IP is attached to a gateway or not.
	IsFree *bool `json:"-"`
}

// GatewayWithPrivateNetworks: gateway with private networks.
type GatewayWithPrivateNetworks struct {
	Gateway *vpcgw_v1.Gateway `json:"gateway"`

	PrivateNetworks []*vpc_v1.PrivateNetwork `json:"private_networks"`

	InstanceID string `json:"instance_id"`
}

// ListGatewaysResponse: list gateways response.
type ListGatewaysResponse struct {
	Gateways []*vpcgw_v1.Gateway `json:"gateways"`

	TotalCount uint32 `json:"total_count"`

	PrivateNetworks []*vpc_v1.PrivateNetwork `json:"private_networks"`
}

// ListIPsResponse: list i ps response.
type ListIPsResponse struct {
	IPs []*IPWithGateway `json:"ips"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.IPs = append(r.IPs, results.IPs...)
	r.TotalCount += uint32(len(results.IPs))
	return uint32(len(results.IPs)), nil
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

// ListGateways:
func (s *ConsoleAPI) ListGateways(req *ConsoleAPIListGatewaysRequest, opts ...scw.RequestOption) (*ListGatewaysResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "private_network_id", req.PrivateNetworkID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-gw-private/v1/zones/" + fmt.Sprint(req.Zone) + "/gateways",
		Query:  query,
	}

	var resp ListGatewaysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGateway:
func (s *ConsoleAPI) GetGateway(req *ConsoleAPIGetGatewayRequest, opts ...scw.RequestOption) (*GatewayWithPrivateNetworks, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayID) == "" {
		return nil, errors.New("field GatewayID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-gw-private/v1/zones/" + fmt.Sprint(req.Zone) + "/gateways/" + fmt.Sprint(req.GatewayID) + "",
	}

	var resp GatewayWithPrivateNetworks

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListIPs:
func (s *ConsoleAPI) ListIPs(req *ConsoleAPIListIPsRequest, opts ...scw.RequestOption) (*ListIPsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "reverse", req.Reverse)
	parameter.AddToQuery(query, "is_free", req.IsFree)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-gw-private/v1/zones/" + fmt.Sprint(req.Zone) + "/ips",
		Query:  query,
	}

	var resp ListIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package vpc_internal provides methods and message types of the vpc_internal v1 API.
package vpc_internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	vpc_v1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/vpc/v1"
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

type ListPrivateNetworksRequestOrderBy string

const (
	ListPrivateNetworksRequestOrderByCreatedAtDesc = ListPrivateNetworksRequestOrderBy("created_at_desc")
	ListPrivateNetworksRequestOrderByCreatedAtAsc  = ListPrivateNetworksRequestOrderBy("created_at_asc")
	ListPrivateNetworksRequestOrderByUpdatedAtDesc = ListPrivateNetworksRequestOrderBy("updated_at_desc")
	ListPrivateNetworksRequestOrderByUpdatedAtAsc  = ListPrivateNetworksRequestOrderBy("updated_at_asc")
)

func (enum ListPrivateNetworksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_desc"
	}
	return string(enum)
}

func (enum ListPrivateNetworksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPrivateNetworksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPrivateNetworksRequestOrderBy(ListPrivateNetworksRequestOrderBy(tmp).String())
	return nil
}

type ListSubnetsRequestOrderBy string

const (
	ListSubnetsRequestOrderByCreatedAtAsc  = ListSubnetsRequestOrderBy("created_at_asc")
	ListSubnetsRequestOrderByCreatedAtDesc = ListSubnetsRequestOrderBy("created_at_desc")
	ListSubnetsRequestOrderBySubnetAsc     = ListSubnetsRequestOrderBy("subnet_asc")
	ListSubnetsRequestOrderBySubnetDesc    = ListSubnetsRequestOrderBy("subnet_desc")
	ListSubnetsRequestOrderBySizeAsc       = ListSubnetsRequestOrderBy("size_asc")
	ListSubnetsRequestOrderBySizeDesc      = ListSubnetsRequestOrderBy("size_desc")
)

func (enum ListSubnetsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSubnetsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSubnetsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSubnetsRequestOrderBy(ListSubnetsRequestOrderBy(tmp).String())
	return nil
}

// PrivateNetwork: private network.
type PrivateNetwork struct {
	PrivateNetwork *vpc_v1.PrivateNetwork `json:"private_network"`

	Vni uint32 `json:"vni"`

	VpcID string `json:"vpc_id"`

	DHCPEnabled bool `json:"dhcp_enabled"`

	DNSServers []net.IP `json:"dns_servers"`

	DNSSearchOverride *string `json:"dns_search_override"`
}

// Subnet: subnet.
type Subnet struct {
	ID string `json:"id"`

	ProjectID string `json:"project_id"`

	PrivateNetworkID string `json:"private_network_id"`

	CreatedAt *time.Time `json:"created_at"`

	Subnet scw.IPNet `json:"subnet"`
}

// GetPrivateNetworkRequest: get private network request.
type GetPrivateNetworkRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PrivateNetworkID string `json:"-"`
}

// GetPrivateNetworksRequest: get private networks request.
type GetPrivateNetworksRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PrivateNetworksIDs []string `json:"-"`
}

// GetPrivateNetworksResponse: get private networks response.
type GetPrivateNetworksResponse struct {
	PrivateNetworks []*PrivateNetwork `json:"private_networks"`
}

// GetServiceInfoRequest: get service info request.
type GetServiceInfoRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// GetSubnetRequest: get subnet request.
type GetSubnetRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SubnetID string `json:"-"`
}

// GetSubnetsRequest: get subnets request.
type GetSubnetsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SubnetIDs []string `json:"-"`
}

// GetSubnetsResponse: get subnets response.
type GetSubnetsResponse struct {
	Subnets []*Subnet `json:"subnets"`
}

// ListPrivateNetworksRequest: list private networks request.
type ListPrivateNetworksRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: default value: created_at_desc
	OrderBy ListPrivateNetworksRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	ProjectID *string `json:"-"`

	Vni *uint32 `json:"-"`
}

// ListPrivateNetworksResponse: list private networks response.
type ListPrivateNetworksResponse struct {
	PrivateNetworks []*PrivateNetwork `json:"private_networks"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPrivateNetworksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPrivateNetworksResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPrivateNetworksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PrivateNetworks = append(r.PrivateNetworks, results.PrivateNetworks...)
	r.TotalCount += uint32(len(results.PrivateNetworks))
	return uint32(len(results.PrivateNetworks)), nil
}

// ListSubnetsRequest: list subnets request.
type ListSubnetsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListSubnetsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	ProjectID *string `json:"-"`

	PrivateNetworkID *string `json:"-"`

	ContainedBy *scw.IPNet `json:"-"`

	Containing *scw.IPNet `json:"-"`

	IsIPv6 *bool `json:"-"`
}

// ListSubnetsResponse: list subnets response.
type ListSubnetsResponse struct {
	Subnets []*Subnet `json:"subnets"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSubnetsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSubnetsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSubnetsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Subnets = append(r.Subnets, results.Subnets...)
	r.TotalCount += uint32(len(results.Subnets))
	return uint32(len(results.Subnets)), nil
}

// SetPrivateNicMetadataRequest: set private nic metadata request.
type SetPrivateNicMetadataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PrivateNicID string `json:"-"`

	ProductType string `json:"product_type"`

	ResourceType string `json:"resource_type"`

	ResourceID string `json:"resource_id"`
}

// UnsetPrivateNicMetadataRequest: unset private nic metadata request.
type UnsetPrivateNicMetadataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PrivateNicID string `json:"-"`
}

// UpdatePrivateNetworkRequest: update private network request.
type UpdatePrivateNetworkRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PrivateNetworkID string `json:"-"`

	DHCPEnabled *bool `json:"dhcp_enabled,omitempty"`

	DNSServers *[]string `json:"dns_servers,omitempty"`

	DNSSearchOverride *string `json:"dns_search_override,omitempty"`
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
func (s *API) Zones() []scw.Zone {
	return []scw.Zone{}
}

// GetServiceInfo:
func (s *API) GetServiceInfo(req *GetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-internal/v1/zones/" + fmt.Sprint(req.Zone) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrivateNetwork:
func (s *API) GetPrivateNetwork(req *GetPrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrivateNetworks:
func (s *API) GetPrivateNetworks(req *GetPrivateNetworksRequest, opts ...scw.RequestOption) (*GetPrivateNetworksResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworksIDs) == "" {
		return nil, errors.New("field PrivateNetworksIDs cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/private-networks/multiple/" + fmt.Sprint(req.PrivateNetworksIDs) + "",
	}

	var resp GetPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPrivateNetworks:
func (s *API) ListPrivateNetworks(req *ListPrivateNetworksRequest, opts ...scw.RequestOption) (*ListPrivateNetworksResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "vni", req.Vni)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/private-networks",
		Query:  query,
	}

	var resp ListPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePrivateNetwork:
func (s *API) UpdatePrivateNetwork(req *UpdatePrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/vpc-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetPrivateNicMetadata:
func (s *API) SetPrivateNicMetadata(req *SetPrivateNicMetadataRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNicID) == "" {
		return errors.New("field PrivateNicID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/vpc-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/nics/" + fmt.Sprint(req.PrivateNicID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UnsetPrivateNicMetadata:
func (s *API) UnsetPrivateNicMetadata(req *UnsetPrivateNicMetadataRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNicID) == "" {
		return errors.New("field PrivateNicID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/nics/" + fmt.Sprint(req.PrivateNicID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListSubnets:
func (s *API) ListSubnets(req *ListSubnetsRequest, opts ...scw.RequestOption) (*ListSubnetsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "private_network_id", req.PrivateNetworkID)
	parameter.AddToQuery(query, "contained_by", req.ContainedBy)
	parameter.AddToQuery(query, "containing", req.Containing)
	parameter.AddToQuery(query, "is_ipv6", req.IsIPv6)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/subnets",
		Query:  query,
	}

	var resp ListSubnetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSubnet:
func (s *API) GetSubnet(req *GetSubnetRequest, opts ...scw.RequestOption) (*Subnet, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SubnetID) == "" {
		return nil, errors.New("field SubnetID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/subnets/" + fmt.Sprint(req.SubnetID) + "",
	}

	var resp Subnet

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSubnets:
func (s *API) GetSubnets(req *GetSubnetsRequest, opts ...scw.RequestOption) (*GetSubnetsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SubnetIDs) == "" {
		return nil, errors.New("field SubnetIDs cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/subnets/multiple/" + fmt.Sprint(req.SubnetIDs) + "",
	}

	var resp GetSubnetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

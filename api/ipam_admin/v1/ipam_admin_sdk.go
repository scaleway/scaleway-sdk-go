// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package ipam_admin provides methods and message types of the ipam_admin v1 API.
package ipam_admin

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

type ListIPAllocationsRequestOrderBy string

const (
	ListIPAllocationsRequestOrderByCreatedAtDesc = ListIPAllocationsRequestOrderBy("created_at_desc")
	ListIPAllocationsRequestOrderByCreatedAtAsc  = ListIPAllocationsRequestOrderBy("created_at_asc")
	ListIPAllocationsRequestOrderByUpdatedAtDesc = ListIPAllocationsRequestOrderBy("updated_at_desc")
	ListIPAllocationsRequestOrderByUpdatedAtAsc  = ListIPAllocationsRequestOrderBy("updated_at_asc")
)

func (enum ListIPAllocationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_desc"
	}
	return string(enum)
}

func (enum ListIPAllocationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListIPAllocationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListIPAllocationsRequestOrderBy(ListIPAllocationsRequestOrderBy(tmp).String())
	return nil
}

type ListIPsRequestOrderBy string

const (
	ListIPsRequestOrderByCreatedAtDesc  = ListIPsRequestOrderBy("created_at_desc")
	ListIPsRequestOrderByCreatedAtAsc   = ListIPsRequestOrderBy("created_at_asc")
	ListIPsRequestOrderByUpdatedAtDesc  = ListIPsRequestOrderBy("updated_at_desc")
	ListIPsRequestOrderByUpdatedAtAsc   = ListIPsRequestOrderBy("updated_at_asc")
	ListIPsRequestOrderByAttachedAtDesc = ListIPsRequestOrderBy("attached_at_desc")
	ListIPsRequestOrderByAttachedAtAsc  = ListIPsRequestOrderBy("attached_at_asc")
)

func (enum ListIPsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_desc"
	}
	return string(enum)
}

func (enum ListIPsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListIPsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListIPsRequestOrderBy(ListIPsRequestOrderBy(tmp).String())
	return nil
}

type ResourceType string

const (
	ResourceTypeUnknownType         = ResourceType("unknown_type")
	ResourceTypeInstanceServer      = ResourceType("instance_server")
	ResourceTypeInstanceIP          = ResourceType("instance_ip")
	ResourceTypeInstancePrivateNic  = ResourceType("instance_private_nic")
	ResourceTypeLBServer            = ResourceType("lb_server")
	ResourceTypeFipIP               = ResourceType("fip_ip")
	ResourceTypeVpcGateway          = ResourceType("vpc_gateway")
	ResourceTypeVpcGatewayNetwork   = ResourceType("vpc_gateway_network")
	ResourceTypeK8sNode             = ResourceType("k8s_node")
	ResourceTypeK8sCluster          = ResourceType("k8s_cluster")
	ResourceTypeRdbInstance         = ResourceType("rdb_instance")
	ResourceTypeRedisCluster        = ResourceType("redis_cluster")
	ResourceTypeBaremetalServer     = ResourceType("baremetal_server")
	ResourceTypeBaremetalPrivateNic = ResourceType("baremetal_private_nic")
)

func (enum ResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum ResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ResourceType(ResourceType(tmp).String())
	return nil
}

// Resource: resource.
type Resource struct {
	// Type: type of resource the IP is attached to.
	// Default value: unknown_type
	Type ResourceType `json:"type"`

	// ID: ID of the resource the IP is attached to.
	ID string `json:"id"`

	// MacAddress: mAC of the resource the IP is attached to.
	MacAddress *string `json:"mac_address"`

	// Name: when the IP is in a Private Network, then a DNS record is available to resolve the resource name to this IP.
	Name *string `json:"name"`
}

// Reverse: reverse.
type Reverse struct {
	Address *scw.IPNet `json:"address"`

	Hostname string `json:"hostname"`
}

// Source: source.
type Source struct {
	// Zonal: this source is global.
	// Precisely one of Zonal, ZonalNat, Regional, PrivateNetworkID, SubnetID must be set.
	Zonal *string `json:"zonal,omitempty"`

	// Precisely one of Zonal, ZonalNat, Regional, PrivateNetworkID, SubnetID must be set.
	ZonalNat *string `json:"zonal_nat,omitempty"`

	// Precisely one of Zonal, ZonalNat, Regional, PrivateNetworkID, SubnetID must be set.
	Regional *bool `json:"regional,omitempty"`

	// PrivateNetworkID: this source is specific.
	// Precisely one of Zonal, ZonalNat, Regional, PrivateNetworkID, SubnetID must be set.
	PrivateNetworkID *string `json:"private_network_id,omitempty"`

	// SubnetID: this source is specific.
	// Precisely one of Zonal, ZonalNat, Regional, PrivateNetworkID, SubnetID must be set.
	SubnetID *string `json:"subnet_id,omitempty"`
}

// BookIPSetRequestIP: book ip set request ip.
type BookIPSetRequestIP struct {
	IsIPv6 bool `json:"is_ipv6"`

	Tags []string `json:"tags"`
}

// IP: ip.
type IP struct {
	// ID: IP ID.
	ID string `json:"id"`

	// Address: iPv4 or IPv6 address in CIDR notation.
	Address scw.IPNet `json:"address"`

	// ProjectID: scaleway Project the IP belongs to.
	ProjectID string `json:"project_id"`

	// IsIPv6: defines whether the IP is an IPv6 (false = IPv4).
	IsIPv6 bool `json:"is_ipv6"`

	// CreatedAt: date the IP was booked.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the IP was last modified.
	UpdatedAt *time.Time `json:"updated_at"`

	// Precisely one of Regional, Zonal, ZonalNat, SubnetID must be set.
	Regional *bool `json:"regional,omitempty"`

	// Precisely one of Regional, Zonal, ZonalNat, SubnetID must be set.
	Zonal *string `json:"zonal,omitempty"`

	// Precisely one of Regional, Zonal, ZonalNat, SubnetID must be set.
	ZonalNat *string `json:"zonal_nat,omitempty"`

	// Precisely one of Regional, Zonal, ZonalNat, SubnetID must be set.
	SubnetID *string `json:"subnet_id,omitempty"`

	// Resource: resource which the IP is attached to.
	Resource *Resource `json:"resource"`

	// Tags: tags for the IP.
	Tags []string `json:"tags"`

	Reverses []*Reverse `json:"reverses"`

	// Region: region of the IP.
	Region scw.Region `json:"region"`

	// Zone: zone of the IP, if zonal.
	Zone *scw.Zone `json:"zone"`
}

// IPAllocation: ip allocation.
type IPAllocation struct {
	AllocatedAt *time.Time `json:"allocated_at"`

	IPID string `json:"ip_id"`

	ProjectID string `json:"project_id"`

	ReleasedAt *time.Time `json:"released_at"`
}

// AttachIPRequest: attach ip request.
type AttachIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPID string `json:"-"`

	Resource *Resource `json:"resource,omitempty"`
}

// AttachIPSetRequest: attach ip set request.
type AttachIPSetRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPIDs []string `json:"ip_ids"`

	Resource *Resource `json:"resource,omitempty"`
}

// BookIPRequest: book ip request.
type BookIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: when creating an IP in a Private Network, the Project must match the Private Network's Project.
	ProjectID string `json:"project_id"`

	// Source: source in which to book the IP. Not all sources are available for booking.
	Source *Source `json:"source,omitempty"`

	// IsIPv6: request an IPv6 instead of an IPv4.
	IsIPv6 bool `json:"is_ipv6"`

	Resource *Resource `json:"resource,omitempty"`

	// Tags: tags for the IP.
	Tags []string `json:"tags"`
}

// BookIPSetRequest: book ip set request.
type BookIPSetRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPs []*BookIPSetRequestIP `json:"ips"`

	ProjectID string `json:"project_id"`

	Resource *Resource `json:"resource,omitempty"`

	Source *Source `json:"source,omitempty"`
}

// CheckIPSetConsistencyRequest: check ip set consistency request.
type CheckIPSetConsistencyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	AuthID string `json:"auth_id"`

	IPIDs []string `json:"ip_ids"`

	ResourceProjectID string `json:"resource_project_id"`

	Source *Source `json:"source,omitempty"`
}

// DetachIPRequest: detach ip request.
type DetachIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPID string `json:"-"`

	Resource *Resource `json:"resource,omitempty"`
}

// DetachIPSetRequest: detach ip set request.
type DetachIPSetRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPIDs []string `json:"ip_ids"`

	Resource *Resource `json:"resource,omitempty"`
}

// GetIPRequest: get ip request.
type GetIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPID: IP ID.
	IPID string `json:"-"`
}

// GetIPSetRequest: get ip set request.
type GetIPSetRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPIDs []string `json:"ip_ids"`
}

// GetServiceInfoRequest: get service info request.
type GetServiceInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// IPSet: ip set.
type IPSet struct {
	IPs []*IP `json:"ips"`
}

// ImportIPRequest: import ip request.
type ImportIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPV4 net.IP `json:"ip_v4"`

	ProjectID string `json:"project_id"`

	Resource *Resource `json:"resource,omitempty"`

	Tags []string `json:"tags"`
}

// ListIPAllocationsRequest: list ip allocations request.
type ListIPAllocationsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPID string `json:"-"`

	AllocatedAfter *time.Time `json:"-"`

	AllocatedBefore *time.Time `json:"-"`

	// OrderBy: default value: created_at_desc
	OrderBy ListIPAllocationsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListIPAllocationsResponse: list ip allocations response.
type ListIPAllocationsResponse struct {
	TotalCount uint64 `json:"total_count"`

	Address net.IP `json:"address"`

	IPAllocations []*IPAllocation `json:"ip_allocations"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIPAllocationsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIPAllocationsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListIPAllocationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.IPAllocations = append(r.IPAllocations, results.IPAllocations...)
	r.TotalCount += uint64(len(results.IPAllocations))
	return uint64(len(results.IPAllocations)), nil
}

// ListIPsRequest: list i ps request.
type ListIPsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: sort order of the returned IPs.
	// Default value: created_at_desc
	OrderBy ListIPsRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: maximum number of IPs to return per page.
	PageSize *uint32 `json:"-"`

	// ProjectID: project ID to filter for. Only IPs belonging to this Project will be returned.
	ProjectID *string `json:"-"`

	// Zonal: zone to filter for. Only IPs that are zonal, and in this zone, will be returned.
	// Precisely one of Zonal, PrivateNetworkID, ZonalNat, Regional, SubnetID must be set.
	Zonal *string `json:"zonal,omitempty"`

	// PrivateNetworkID: only IPs that are private, and in this Private Network, will be returned.
	// Precisely one of Zonal, PrivateNetworkID, ZonalNat, Regional, SubnetID must be set.
	PrivateNetworkID *string `json:"private_network_id,omitempty"`

	// Attached: defines whether to filter only for IPs which are attached to a resource.
	Attached *bool `json:"-"`

	// ResourceID: resource ID to filter for. Only IPs attached to this resource will be returned.
	ResourceID *string `json:"-"`

	// ResourceType: resource type to filter for. Only IPs attached to this type of resource will be returned.
	// Default value: unknown_type
	ResourceType ResourceType `json:"-"`

	// MacAddress: mAC address to filter for. Only IPs attached to a resource with this MAC address will be returned.
	MacAddress *string `json:"-"`

	// Tags: tags to filter for, only IPs with one or more matching tags will be returned.
	Tags []string `json:"-"`

	// OrganizationID: organization ID to filter for. Only IPs belonging to this Organization will be returned.
	OrganizationID *string `json:"-"`

	// IsIPv6: defines whether to filter only for IPv4s or IPv6s.
	IsIPv6 *bool `json:"-"`

	// ResourceName: attached resource name to filter for, only IPs attached to a resource with this string within their name will be returned.
	ResourceName *string `json:"-"`

	// Deprecated
	Source *Source `json:"-"`

	// Precisely one of Zonal, PrivateNetworkID, ZonalNat, Regional, SubnetID must be set.
	ZonalNat *string `json:"zonal_nat,omitempty"`

	// Precisely one of Zonal, PrivateNetworkID, ZonalNat, Regional, SubnetID must be set.
	Regional *bool `json:"regional,omitempty"`

	// Precisely one of Zonal, PrivateNetworkID, ZonalNat, Regional, SubnetID must be set.
	SubnetID *string `json:"subnet_id,omitempty"`

	// Deprecated
	TagValues *[]string `json:"-"`

	ResourceIDs []string `json:"-"`
}

// ListIPsResponse: list i ps response.
type ListIPsResponse struct {
	TotalCount uint64 `json:"total_count"`

	IPs []*IP `json:"ips"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.IPs = append(r.IPs, results.IPs...)
	r.TotalCount += uint64(len(results.IPs))
	return uint64(len(results.IPs)), nil
}

// MoveIPRequest: move ip request.
type MoveIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPID string `json:"-"`

	FromResource *Resource `json:"from_resource,omitempty"`

	ToResource *Resource `json:"to_resource,omitempty"`
}

// MoveIPSetRequest: move ip set request.
type MoveIPSetRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	FromResource *Resource `json:"from_resource,omitempty"`

	IPIDs []string `json:"ip_ids"`

	ToResource *Resource `json:"to_resource,omitempty"`
}

// PurgeIPSetRequest: purge ip set request.
type PurgeIPSetRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Addresses []scw.IPNet `json:"addresses"`
}

// ReleaseIPRequest: release ip request.
type ReleaseIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPID: IP ID.
	IPID string `json:"-"`

	Resource *Resource `json:"resource,omitempty"`
}

// ReleaseIPSetRequest: release ip set request.
type ReleaseIPSetRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPIDs []string `json:"ip_ids"`

	Resource *Resource `json:"resource,omitempty"`
}

// SetIPMACAddressRequest: set ipmac address request.
type SetIPMACAddressRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPID string `json:"-"`

	MacAddress string `json:"mac_address"`
}

// SetIPResourceNameRequest: set ip resource name request.
type SetIPResourceNameRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPID string `json:"-"`

	OnlyIfUnset bool `json:"only_if_unset"`

	ResourceName string `json:"resource_name"`
}

// SetIPSetMACAddressRequest: set ip set mac address request.
type SetIPSetMACAddressRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IPIDs []string `json:"ip_ids"`

	MacAddress string `json:"mac_address"`
}

// UpdateIPRequest: update ip request.
type UpdateIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPID: IP ID.
	IPID string `json:"-"`

	// Tags: tags for the IP.
	Tags *[]string `json:"tags,omitempty"`

	Reverses []*Reverse `json:"reverses"`
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
func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// GetServiceInfo:
func (s *API) GetServiceInfo(req *GetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
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
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BookIP: Book a new IP from the specified source. Currently IPs can only be booked from a Private Network.
func (s *API) BookIP(req *BookIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ips",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BookIPSet:
func (s *API) BookIPSet(req *BookIPSetRequest, opts ...scw.RequestOption) (*IPSet, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ip-sets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IPSet

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ImportIP:
func (s *API) ImportIP(req *ImportIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/import-ip",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReleaseIP: Release an IP not currently attached to a resource, and returns it to the available IP pool.
func (s *API) ReleaseIP(req *ReleaseIPRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "/release",
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

// ReleaseIPSet:
func (s *API) ReleaseIPSet(req *ReleaseIPSetRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ip-sets/release",
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

// PurgeIPSet:
func (s *API) PurgeIPSet(req *PurgeIPSetRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ip-sets/purge",
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

// GetIP: Retrieve details of an existing IP, specified by its IP ID.
func (s *API) GetIP(req *GetIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIPSet:
func (s *API) GetIPSet(req *GetIPSetRequest, opts ...scw.RequestOption) (*IPSet, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "ip_ids", req.IPIDs)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ip-sets",
		Query:  query,
	}

	var resp IPSet

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateIP: Update parameters including tags of the specified IP.
func (s *API) UpdateIP(req *UpdateIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListIPs: List existing IPs in the specified region using various filters. For example, you can filter for IPs within a specified Private Network, or for public IPs within a specified Project. By default, the IPs returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.
func (s *API) ListIPs(req *ListIPsRequest, opts ...scw.RequestOption) (*ListIPsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
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
	parameter.AddToQuery(query, "attached", req.Attached)
	parameter.AddToQuery(query, "resource_id", req.ResourceID)
	parameter.AddToQuery(query, "resource_type", req.ResourceType)
	parameter.AddToQuery(query, "mac_address", req.MacAddress)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "is_ipv6", req.IsIPv6)
	parameter.AddToQuery(query, "resource_name", req.ResourceName)
	parameter.AddToQuery(query, "source", req.Source)
	parameter.AddToQuery(query, "tag_values", req.TagValues)
	parameter.AddToQuery(query, "resource_ids", req.ResourceIDs)
	parameter.AddToQuery(query, "zonal", req.Zonal)
	parameter.AddToQuery(query, "private_network_id", req.PrivateNetworkID)
	parameter.AddToQuery(query, "zonal_nat", req.ZonalNat)
	parameter.AddToQuery(query, "regional", req.Regional)
	parameter.AddToQuery(query, "subnet_id", req.SubnetID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ips",
		Query:  query,
	}

	var resp ListIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachIP:
func (s *API) AttachIP(req *AttachIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "/attach",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachIPSet:
func (s *API) AttachIPSet(req *AttachIPSetRequest, opts ...scw.RequestOption) (*IPSet, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ip-sets/attach",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IPSet

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachIP:
func (s *API) DetachIP(req *DetachIPRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "/detach",
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

// DetachIPSet:
func (s *API) DetachIPSet(req *DetachIPSetRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ip-sets/detach",
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

// MoveIP:
func (s *API) MoveIP(req *MoveIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "/move",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MoveIPSet:
func (s *API) MoveIPSet(req *MoveIPSetRequest, opts ...scw.RequestOption) (*IPSet, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ip-sets/move",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IPSet

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIPMACAddress:
func (s *API) SetIPMACAddress(req *SetIPMACAddressRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "/mac-address",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIPSetMACAddress:
func (s *API) SetIPSetMACAddress(req *SetIPSetMACAddressRequest, opts ...scw.RequestOption) (*IPSet, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ip-sets/mac-address",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IPSet

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetIPResourceName:
func (s *API) SetIPResourceName(req *SetIPResourceNameRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "/resource-name",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListIPAllocations:
func (s *API) ListIPAllocations(req *ListIPAllocationsRequest, opts ...scw.RequestOption) (*ListIPAllocationsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "allocated_after", req.AllocatedAfter)
	parameter.AddToQuery(query, "allocated_before", req.AllocatedBefore)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "/allocations",
		Query:  query,
	}

	var resp ListIPAllocationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckIPSetConsistency:
func (s *API) CheckIPSetConsistency(req *CheckIPSetConsistencyRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipam-admin/v1/regions/" + fmt.Sprint(req.Region) + "/ip-sets/check-consistency",
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

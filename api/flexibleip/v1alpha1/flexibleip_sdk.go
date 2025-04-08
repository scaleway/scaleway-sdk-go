// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package flexibleip provides methods and message types of the flexibleip v1alpha1 API.
package flexibleip

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
	"github.com/scaleway/scaleway-sdk-go/scw"
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

type FlexibleIPStatus string

const (
	FlexibleIPStatusUnknown   = FlexibleIPStatus("unknown")
	FlexibleIPStatusReady     = FlexibleIPStatus("ready")
	FlexibleIPStatusUpdating  = FlexibleIPStatus("updating")
	FlexibleIPStatusAttached  = FlexibleIPStatus("attached")
	FlexibleIPStatusError     = FlexibleIPStatus("error")
	FlexibleIPStatusDetaching = FlexibleIPStatus("detaching")
	FlexibleIPStatusLocked    = FlexibleIPStatus("locked")
)

func (enum FlexibleIPStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(FlexibleIPStatusUnknown)
	}
	return string(enum)
}

func (enum FlexibleIPStatus) Values() []FlexibleIPStatus {
	return []FlexibleIPStatus{
		"unknown",
		"ready",
		"updating",
		"attached",
		"error",
		"detaching",
		"locked",
	}
}

func (enum FlexibleIPStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FlexibleIPStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FlexibleIPStatus(FlexibleIPStatus(tmp).String())
	return nil
}

type ListFlexibleIPsRequestOrderBy string

const (
	ListFlexibleIPsRequestOrderByCreatedAtAsc  = ListFlexibleIPsRequestOrderBy("created_at_asc")
	ListFlexibleIPsRequestOrderByCreatedAtDesc = ListFlexibleIPsRequestOrderBy("created_at_desc")
)

func (enum ListFlexibleIPsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListFlexibleIPsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListFlexibleIPsRequestOrderBy) Values() []ListFlexibleIPsRequestOrderBy {
	return []ListFlexibleIPsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListFlexibleIPsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFlexibleIPsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFlexibleIPsRequestOrderBy(ListFlexibleIPsRequestOrderBy(tmp).String())
	return nil
}

type MACAddressStatus string

const (
	MACAddressStatusUnknown  = MACAddressStatus("unknown")
	MACAddressStatusReady    = MACAddressStatus("ready")
	MACAddressStatusUpdating = MACAddressStatus("updating")
	MACAddressStatusUsed     = MACAddressStatus("used")
	MACAddressStatusError    = MACAddressStatus("error")
	MACAddressStatusDeleting = MACAddressStatus("deleting")
)

func (enum MACAddressStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(MACAddressStatusUnknown)
	}
	return string(enum)
}

func (enum MACAddressStatus) Values() []MACAddressStatus {
	return []MACAddressStatus{
		"unknown",
		"ready",
		"updating",
		"used",
		"error",
		"deleting",
	}
}

func (enum MACAddressStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MACAddressStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MACAddressStatus(MACAddressStatus(tmp).String())
	return nil
}

type MACAddressType string

const (
	MACAddressTypeUnknownType = MACAddressType("unknown_type")
	MACAddressTypeVmware      = MACAddressType("vmware")
	MACAddressTypeXen         = MACAddressType("xen")
	MACAddressTypeKvm         = MACAddressType("kvm")
)

func (enum MACAddressType) String() string {
	if enum == "" {
		// return default value if empty
		return string(MACAddressTypeUnknownType)
	}
	return string(enum)
}

func (enum MACAddressType) Values() []MACAddressType {
	return []MACAddressType{
		"unknown_type",
		"vmware",
		"xen",
		"kvm",
	}
}

func (enum MACAddressType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MACAddressType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MACAddressType(MACAddressType(tmp).String())
	return nil
}

// MACAddress: mac address.
type MACAddress struct {
	// ID: ID of the flexible IP.
	ID string `json:"id"`

	// MacAddress: mAC address of the Virtual MAC.
	MacAddress string `json:"mac_address"`

	// MacType: type of virtual MAC.
	// Default value: unknown_type
	MacType MACAddressType `json:"mac_type"`

	// Status: status of virtual MAC.
	// Default value: unknown
	Status MACAddressStatus `json:"status"`

	// UpdatedAt: date on which the virtual MAC was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// CreatedAt: date on which the virtual MAC was created.
	CreatedAt *time.Time `json:"created_at"`

	// Zone: mAC address IP Availability Zone.
	Zone scw.Zone `json:"zone"`
}

// FlexibleIP: flexible ip.
type FlexibleIP struct {
	// ID: ID of the flexible IP.
	ID string `json:"id"`

	// OrganizationID: ID of the Organization the flexible IP is attached to.
	OrganizationID string `json:"organization_id"`

	// ProjectID: ID of the Project the flexible IP is attached to.
	ProjectID string `json:"project_id"`

	// Description: flexible IP description.
	Description string `json:"description"`

	// Tags: flexible IP tags.
	Tags []string `json:"tags"`

	// UpdatedAt: date on which the flexible IP was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// CreatedAt: date on which the flexible IP was created.
	CreatedAt *time.Time `json:"created_at"`

	// Status: - ready : flexible IP is created and ready to be attached to a server or to be associated with a virtual MAC.
	// - updating: flexible IP is being attached to a server or a virtual MAC operation is ongoing
	// - attached: flexible IP is attached to a server
	// - error: a flexible IP operation resulted in an error
	// - detaching: flexible IP is being detached from a server
	// - locked: the resource of the flexible IP is locked.
	// Default value: unknown
	Status FlexibleIPStatus `json:"status"`

	// IPAddress: IP of the flexible IP.
	IPAddress scw.IPNet `json:"ip_address"`

	// MacAddress: mAC address of the flexible IP.
	MacAddress *MACAddress `json:"mac_address"`

	// ServerID: ID of the server linked to the flexible IP.
	ServerID *string `json:"server_id"`

	// Reverse: reverse DNS value.
	Reverse string `json:"reverse"`

	// Zone: availability Zone of the flexible IP.
	Zone scw.Zone `json:"zone"`
}

// AttachFlexibleIPRequest: attach flexible ip request.
type AttachFlexibleIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// FipsIDs: multiple IDs can be provided, but note that flexible IPs must belong to the same MAC group (see details about MAC groups).
	FipsIDs []string `json:"fips_ids"`

	// ServerID: ID of the server on which to attach the flexible IPs.
	ServerID string `json:"server_id"`
}

// AttachFlexibleIPsResponse: attach flexible i ps response.
type AttachFlexibleIPsResponse struct {
	// TotalCount: total count of flexible IPs that are being updated.
	TotalCount uint32 `json:"total_count"`

	// FlexibleIPs: list of flexible IPs in an updating state.
	FlexibleIPs []*FlexibleIP `json:"flexible_ips"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *AttachFlexibleIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *AttachFlexibleIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*AttachFlexibleIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.FlexibleIPs = append(r.FlexibleIPs, results.FlexibleIPs...)
	r.TotalCount += uint32(len(results.FlexibleIPs))
	return uint32(len(results.FlexibleIPs)), nil
}

// CreateFlexibleIPRequest: create flexible ip request.
type CreateFlexibleIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ProjectID: ID of the project to associate with the Flexible IP.
	ProjectID string `json:"project_id"`

	// Description: flexible IP description (max. of 255 characters).
	Description string `json:"description"`

	// Tags: tags to associate to the flexible IP.
	Tags []string `json:"tags"`

	// ServerID: ID of the server to which the newly created flexible IP will be attached.
	ServerID *string `json:"server_id,omitempty"`

	// Reverse: value of the reverse DNS.
	Reverse *string `json:"reverse,omitempty"`

	// IsIPv6: defines whether the flexible IP has an IPv6 address.
	IsIPv6 bool `json:"is_ipv6"`
}

// DeleteFlexibleIPRequest: delete flexible ip request.
type DeleteFlexibleIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// FipID: ID of the flexible IP to delete.
	FipID string `json:"-"`
}

// DeleteMACAddrRequest: delete mac addr request.
type DeleteMACAddrRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// FipID: if the flexible IP belongs to a MAC group, the MAC will be removed from both the MAC group and flexible IP.
	FipID string `json:"-"`
}

// DetachFlexibleIPRequest: detach flexible ip request.
type DetachFlexibleIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// FipsIDs: list of flexible IP IDs to detach from a server. Multiple IDs can be provided. Note that flexible IPs must belong to the same MAC group.
	FipsIDs []string `json:"fips_ids"`
}

// DetachFlexibleIPsResponse: detach flexible i ps response.
type DetachFlexibleIPsResponse struct {
	// TotalCount: total count of flexible IPs that are being detached.
	TotalCount uint32 `json:"total_count"`

	// FlexibleIPs: list of flexible IPs in a detaching state.
	FlexibleIPs []*FlexibleIP `json:"flexible_ips"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *DetachFlexibleIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *DetachFlexibleIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*DetachFlexibleIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.FlexibleIPs = append(r.FlexibleIPs, results.FlexibleIPs...)
	r.TotalCount += uint32(len(results.FlexibleIPs))
	return uint32(len(results.FlexibleIPs)), nil
}

// DuplicateMACAddrRequest: duplicate mac addr request.
type DuplicateMACAddrRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// FipID: note that the flexible IPs need to be attached to the same server.
	FipID string `json:"-"`

	// DuplicateFromFipID: note that flexible IPs need to be attached to the same server.
	DuplicateFromFipID string `json:"duplicate_from_fip_id"`
}

// GenerateMACAddrRequest: generate mac addr request.
type GenerateMACAddrRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// FipID: ID of the flexible IP for which to generate a virtual MAC.
	FipID string `json:"-"`

	// MacType: tODO.
	// Default value: unknown_type
	MacType MACAddressType `json:"mac_type"`
}

// GetFlexibleIPRequest: get flexible ip request.
type GetFlexibleIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// FipID: ID of the flexible IP.
	FipID string `json:"-"`
}

// ListFlexibleIPsRequest: list flexible i ps request.
type ListFlexibleIPsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: sort order of the returned flexible IPs.
	// Default value: created_at_asc
	OrderBy ListFlexibleIPsRequestOrderBy `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of flexible IPs per page.
	PageSize *uint32 `json:"-"`

	// Tags: filter by tag, only flexible IPs with one or more matching tags will be returned.
	Tags []string `json:"-"`

	// Status: filter by status, only flexible IPs with this status will be returned.
	Status []FlexibleIPStatus `json:"-"`

	// ServerIDs: filter by server IDs, only flexible IPs with these server IDs will be returned.
	ServerIDs []string `json:"-"`

	// OrganizationID: filter by Organization ID, only flexible IPs from this Organization will be returned.
	OrganizationID *string `json:"-"`

	// ProjectID: filter by Project ID, only flexible IPs from this Project will be returned.
	ProjectID *string `json:"-"`
}

// ListFlexibleIPsResponse: list flexible i ps response.
type ListFlexibleIPsResponse struct {
	// TotalCount: total count of matching flexible IPs.
	TotalCount uint32 `json:"total_count"`

	// FlexibleIPs: list of all flexible IPs.
	FlexibleIPs []*FlexibleIP `json:"flexible_ips"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFlexibleIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFlexibleIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListFlexibleIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.FlexibleIPs = append(r.FlexibleIPs, results.FlexibleIPs...)
	r.TotalCount += uint32(len(results.FlexibleIPs))
	return uint32(len(results.FlexibleIPs)), nil
}

// MoveMACAddrRequest: move mac addr request.
type MoveMACAddrRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	FipID string `json:"-"`

	DstFipID string `json:"dst_fip_id"`
}

// UpdateFlexibleIPRequest: update flexible ip request.
type UpdateFlexibleIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// FipID: ID of the flexible IP to update.
	FipID string `json:"-"`

	// Description: flexible IP description (max. 255 characters).
	Description *string `json:"description,omitempty"`

	// Tags: tags associated with the flexible IP.
	Tags *[]string `json:"tags,omitempty"`

	// Reverse: value of the reverse DNS.
	Reverse *string `json:"reverse,omitempty"`
}

// This API allows you to manage your Elastic Metal servers' flexible public IP addresses.
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
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3}
}

// CreateFlexibleIP: Generate a new flexible IP within a given zone, specifying its configuration including Project ID and description.
func (s *API) CreateFlexibleIP(req *CreateFlexibleIPRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFlexibleIP: Retrieve information about an existing flexible IP, specified by its ID and zone. Its full details, including Project ID, description and status, are returned in the response object.
func (s *API) GetFlexibleIP(req *GetFlexibleIPRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return nil, errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "",
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFlexibleIPs: List all flexible IPs within a given zone.
func (s *API) ListFlexibleIPs(req *ListFlexibleIPsRequest, opts ...scw.RequestOption) (*ListFlexibleIPsResponse, error) {
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
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "server_ids", req.ServerIDs)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips",
		Query:  query,
	}

	var resp ListFlexibleIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateFlexibleIP: Update the parameters of an existing flexible IP, specified by its ID and zone. These parameters include tags and description.
func (s *API) UpdateFlexibleIP(req *UpdateFlexibleIPRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return nil, errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFlexibleIP: Delete an existing flexible IP, specified by its ID and zone. Note that deleting a flexible IP is permanent and cannot be undone.
func (s *API) DeleteFlexibleIP(req *DeleteFlexibleIPRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// AttachFlexibleIP: Attach an existing flexible IP to a specified Elastic Metal server.
func (s *API) AttachFlexibleIP(req *AttachFlexibleIPRequest, opts ...scw.RequestOption) (*AttachFlexibleIPsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/attach",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AttachFlexibleIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachFlexibleIP: Detach an existing flexible IP from a specified Elastic Metal server.
func (s *API) DetachFlexibleIP(req *DetachFlexibleIPRequest, opts ...scw.RequestOption) (*DetachFlexibleIPsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/detach",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DetachFlexibleIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateMACAddr: Generate a virtual MAC (Media Access Control) address on an existing flexible IP.
func (s *API) GenerateMACAddr(req *GenerateMACAddrRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return nil, errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/mac",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DuplicateMACAddr: Duplicate a virtual MAC address from a given flexible IP to another flexible IP attached to the same server.
func (s *API) DuplicateMACAddr(req *DuplicateMACAddrRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return nil, errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/mac/duplicate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MoveMACAddr: Relocate a virtual MAC (Media Access Control) address from an existing flexible IP to a different flexible IP.
func (s *API) MoveMACAddr(req *MoveMACAddrRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return nil, errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/mac/move",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteMACAddr: Detach a given MAC (Media Access Control) address from an existing flexible IP.
func (s *API) DeleteMACAddr(req *DeleteMACAddrRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/mac",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

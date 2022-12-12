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

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
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

// API: flexible IP API
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type FlexibleIPStatus string

const (
	// FlexibleIPStatusUnknown is [insert doc].
	FlexibleIPStatusUnknown = FlexibleIPStatus("unknown")
	// FlexibleIPStatusReady is [insert doc].
	FlexibleIPStatusReady = FlexibleIPStatus("ready")
	// FlexibleIPStatusUpdating is [insert doc].
	FlexibleIPStatusUpdating = FlexibleIPStatus("updating")
	// FlexibleIPStatusAttached is [insert doc].
	FlexibleIPStatusAttached = FlexibleIPStatus("attached")
	// FlexibleIPStatusError is [insert doc].
	FlexibleIPStatusError = FlexibleIPStatus("error")
	// FlexibleIPStatusDetaching is [insert doc].
	FlexibleIPStatusDetaching = FlexibleIPStatus("detaching")
	// FlexibleIPStatusLocked is [insert doc].
	FlexibleIPStatusLocked = FlexibleIPStatus("locked")
)

func (enum FlexibleIPStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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
	// ListFlexibleIPsRequestOrderByCreatedAtAsc is [insert doc].
	ListFlexibleIPsRequestOrderByCreatedAtAsc = ListFlexibleIPsRequestOrderBy("created_at_asc")
	// ListFlexibleIPsRequestOrderByCreatedAtDesc is [insert doc].
	ListFlexibleIPsRequestOrderByCreatedAtDesc = ListFlexibleIPsRequestOrderBy("created_at_desc")
)

func (enum ListFlexibleIPsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
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
	// MACAddressStatusUnknown is [insert doc].
	MACAddressStatusUnknown = MACAddressStatus("unknown")
	// MACAddressStatusReady is [insert doc].
	MACAddressStatusReady = MACAddressStatus("ready")
	// MACAddressStatusUpdating is [insert doc].
	MACAddressStatusUpdating = MACAddressStatus("updating")
	// MACAddressStatusUsed is [insert doc].
	MACAddressStatusUsed = MACAddressStatus("used")
	// MACAddressStatusError is [insert doc].
	MACAddressStatusError = MACAddressStatus("error")
	// MACAddressStatusDeleting is [insert doc].
	MACAddressStatusDeleting = MACAddressStatus("deleting")
)

func (enum MACAddressStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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
	// MACAddressTypeUnknownType is [insert doc].
	MACAddressTypeUnknownType = MACAddressType("unknown_type")
	// MACAddressTypeVmware is [insert doc].
	MACAddressTypeVmware = MACAddressType("vmware")
	// MACAddressTypeXen is [insert doc].
	MACAddressTypeXen = MACAddressType("xen")
	// MACAddressTypeKvm is [insert doc].
	MACAddressTypeKvm = MACAddressType("kvm")
)

func (enum MACAddressType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
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

// AttachFlexibleIPsResponse: attach flexible i ps response
type AttachFlexibleIPsResponse struct {
	// TotalCount: total count of Flexible IPs being updated
	TotalCount uint32 `json:"total_count"`
	// FlexibleIPs: listing of Flexible IPs in updating state
	FlexibleIPs []*FlexibleIP `json:"flexible_ips"`
}

// DetachFlexibleIPsResponse: detach flexible i ps response
type DetachFlexibleIPsResponse struct {
	// TotalCount: total count of Flexible IPs being detached
	TotalCount uint32 `json:"total_count"`
	// FlexibleIPs: listing of Flexible IPs in detaching state
	FlexibleIPs []*FlexibleIP `json:"flexible_ips"`
}

// FlexibleIP: flexible ip
type FlexibleIP struct {
	// ID: ID of the Flexible IP
	ID string `json:"id"`
	// OrganizationID: organization ID the Flexible IP is attached to
	OrganizationID string `json:"organization_id"`
	// ProjectID: project ID the Flexible IP is attached to
	ProjectID string `json:"project_id"`
	// Description: description of the Flexible IP
	Description string `json:"description"`
	// Tags: tags associated with the Flexible IP
	Tags []string `json:"tags"`
	// UpdatedAt: date of last update of the Flexible IP
	UpdatedAt *time.Time `json:"updated_at"`
	// CreatedAt: date of creation of the Flexible IP
	CreatedAt *time.Time `json:"created_at"`
	// Status: status of the Flexible IP
	//
	// - ready : Flexible IP is created and ready to be attached to a server or to have a virtual MAC generated.
	// - updating: Flexible IP is being attached to a server or a virtual MAC operation is ongoing
	// - attached: Flexible IP is attached to a server
	// - error: a Flexible IP operation resulted in an error
	// - detaching: Flexible IP is being detached from a server
	// - locked: Flexible IP resource is locked
	//
	// Default value: unknown
	Status FlexibleIPStatus `json:"status"`
	// IPAddress: IP of the Flexible IP
	IPAddress scw.IPNet `json:"ip_address"`
	// MacAddress: mAC address of the Flexible IP
	MacAddress *MACAddress `json:"mac_address"`
	// ServerID: ID of the server linked to the Flexible IP
	ServerID *string `json:"server_id"`
	// Reverse: reverse DNS value
	Reverse string `json:"reverse"`
	// Zone: flexible IP Availability Zone
	Zone scw.Zone `json:"zone"`
}

// ListFlexibleIPsResponse: list flexible i ps response
type ListFlexibleIPsResponse struct {
	// TotalCount: total count of matching Flexible IPs
	TotalCount uint32 `json:"total_count"`
	// FlexibleIPs: listing of Flexible IPs
	FlexibleIPs []*FlexibleIP `json:"flexible_ips"`
}

// MACAddress: mac address
type MACAddress struct {
	// ID: ID of the Flexible IP
	ID string `json:"id"`
	// MacAddress: mAC address of the Virtual MAC
	MacAddress string `json:"mac_address"`
	// MacType: virtual MAC type
	//
	// Default value: unknown_type
	MacType MACAddressType `json:"mac_type"`
	// Status: virtual MAC status
	//
	// Default value: unknown
	Status MACAddressStatus `json:"status"`
	// UpdatedAt: date of last update of the Virtual MAC
	UpdatedAt *time.Time `json:"updated_at"`
	// CreatedAt: date of creation of the Virtual MAC
	CreatedAt *time.Time `json:"created_at"`
	// Zone: mAC Addr IP Availability Zone
	Zone scw.Zone `json:"zone"`
}

// Service API

// Zones list localities the api is available in
func (s *API) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1}
}

type CreateFlexibleIPRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ProjectID: ID of the project to associate with the Flexible IP
	ProjectID string `json:"project_id"`
	// Description: description to associate with the Flexible IP, max 255 characters
	Description string `json:"description"`
	// Tags: tags to associate to the Flexible IP
	Tags []string `json:"tags"`
	// ServerID: server ID on which to attach the created Flexible IP
	ServerID *string `json:"server_id"`
	// Reverse: reverse DNS value
	Reverse *string `json:"reverse"`
	// IsIPv6: if true, creates a Flexible IP with an ipv6 address
	IsIPv6 bool `json:"is_ipv6"`
}

// CreateFlexibleIP: create a Flexible IP
func (s *API) CreateFlexibleIP(req *CreateFlexibleIPRequest, opts ...scw.RequestOption) (*FlexibleIP, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips",
		Headers: http.Header{},
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

type GetFlexibleIPRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// FipID: flexible IP ID
	FipID string `json:"-"`
}

// GetFlexibleIP: get a Flexible IP
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
		Method:  "GET",
		Path:    "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "",
		Headers: http.Header{},
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListFlexibleIPsRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// OrderBy: the sort order of the returned Flexible IPs
	//
	// Default value: created_at_asc
	OrderBy ListFlexibleIPsRequestOrderBy `json:"-"`
	// Page: the page number for the returned Flexible IPs
	Page *int32 `json:"-"`
	// PageSize: the maximum number of Flexible IPs per page
	PageSize *uint32 `json:"-"`
	// Tags: filter Flexible IPs with one or more matching tags
	Tags []string `json:"-"`
	// Status: filter Flexible IPs by status
	Status []FlexibleIPStatus `json:"-"`
	// ServerIDs: filter Flexible IPs by server IDs
	ServerIDs []string `json:"-"`
	// OrganizationID: filter Flexible IPs by organization ID
	OrganizationID *string `json:"-"`
	// ProjectID: filter Flexible IPs by project ID
	ProjectID *string `json:"-"`
}

// ListFlexibleIPs: list Flexible IPs
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
		Method:  "GET",
		Path:    "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListFlexibleIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateFlexibleIPRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// FipID: ID of the Flexible IP to update
	FipID string `json:"-"`
	// Description: description to associate with the Flexible IP, max 255 characters
	Description *string `json:"description"`
	// Tags: tags to associate with the Flexible IP
	Tags *[]string `json:"tags"`
	// Reverse: reverse DNS value
	Reverse *string `json:"reverse"`
}

// UpdateFlexibleIP: update a Flexible IP
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
		Method:  "PATCH",
		Path:    "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "",
		Headers: http.Header{},
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

type DeleteFlexibleIPRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// FipID: ID of the Flexible IP to delete
	FipID string `json:"-"`
}

// DeleteFlexibleIP: delete a Flexible IP
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
		Method:  "DELETE",
		Path:    "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type AttachFlexibleIPRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// FipsIDs: a list of Flexible IP IDs to attach
	//
	// Multiple IDs can be provided as long as Flexible IPs belong to the same MAC groups (see details about MAC groups).
	FipsIDs []string `json:"fips_ids"`
	// ServerID: a server ID on which to attach the Flexible IPs
	ServerID string `json:"server_id"`
}

// AttachFlexibleIP: attach a Flexible IP to a server
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
		Method:  "POST",
		Path:    "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/attach",
		Headers: http.Header{},
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

type DetachFlexibleIPRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// FipsIDs: a list of Flexible IP IDs to detach
	//
	// Multiple IDs can be provided as long as Flexible IPs belong to the same MAC groups (see details about MAC groups).
	FipsIDs []string `json:"fips_ids"`
}

// DetachFlexibleIP: detach a Flexible IP from a server
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
		Method:  "POST",
		Path:    "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/detach",
		Headers: http.Header{},
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

type GenerateMACAddrRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// FipID: flexible IP ID on which to generate a Virtual MAC
	FipID string `json:"-"`
	// MacType: tODO
	//
	// Default value: unknown_type
	MacType MACAddressType `json:"mac_type"`
}

// GenerateMACAddr: generate a virtual MAC on a given Flexible IP
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
		Method:  "POST",
		Path:    "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/mac",
		Headers: http.Header{},
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

type DuplicateMACAddrRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// FipID: flexible IP ID on which to duplicate the Virtual MAC
	//
	// Flexible IPs need to be attached to the same server.
	FipID string `json:"-"`
	// DuplicateFromFipID: flexible IP ID to duplicate the Virtual MAC from
	//
	// Flexible IPs need to be attached to the same server.
	DuplicateFromFipID string `json:"duplicate_from_fip_id"`
}

// DuplicateMACAddr: duplicate a Virtual MAC
//
// Duplicate a Virtual MAC from a given Flexible IP onto another attached on the same server.
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
		Method:  "POST",
		Path:    "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/mac/duplicate",
		Headers: http.Header{},
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

type MoveMACAddrRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`

	FipID string `json:"-"`

	DstFipID string `json:"dst_fip_id"`
}

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
		Method:  "POST",
		Path:    "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/mac/move",
		Headers: http.Header{},
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

type DeleteMACAddrRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// FipID: flexible IP ID from which to delete the Virtual MAC
	//
	// If the Flexible IP belongs to a MAC group, the MAC will be removed from the MAC group and from the Flexible IP.
	FipID string `json:"-"`
}

// DeleteMACAddr: remove a virtual MAC from a Flexible IP
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
		Method:  "DELETE",
		Path:    "/flexible-ip/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/mac",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
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

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package flexibleip_admin provides methods and message types of the flexibleip_admin v1alpha1 API.
package flexibleip_admin

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
	flexibleip_v1alpha1 "github.com/scaleway/scaleway-sdk-go/api/flexibleip/v1alpha1"
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

type ListEventsRequestOrderBy string

const (
	ListEventsRequestOrderByUnknown       = ListEventsRequestOrderBy("unknown")
	ListEventsRequestOrderByCreatedAtAsc  = ListEventsRequestOrderBy("created_at_asc")
	ListEventsRequestOrderByCreatedAtDesc = ListEventsRequestOrderBy("created_at_desc")
)

func (enum ListEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ListEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListEventsRequestOrderBy(ListEventsRequestOrderBy(tmp).String())
	return nil
}

// Event: event.
type Event struct {
	ID string `json:"id"`

	FlexibleIPID string `json:"flexible_ip_id"`

	// OldStatus: default value: unknown
	OldStatus flexibleip_v1alpha1.FlexibleIPStatus `json:"old_status"`

	// NewStatus: default value: unknown
	NewStatus flexibleip_v1alpha1.FlexibleIPStatus `json:"new_status"`

	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`

	ServerID string `json:"server_id"`

	NetboxID int32 `json:"netbox_id"`

	UpdatedAt *time.Time `json:"updated_at"`

	CreatedAt *time.Time `json:"created_at"`
}

// FlexibleIP: flexible ip.
type FlexibleIP struct {
	FlexibleIP *flexibleip_v1alpha1.FlexibleIP `json:"flexible_ip"`

	NetboxID int32 `json:"netbox_id"`
}

// GetFlexibleIPRequest: get flexible ip request.
type GetFlexibleIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// FipID: ID of the flexible IP.
	FipID string `json:"-"`
}

// GetServiceInfoRequest: get service info request.
type GetServiceInfoRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// ListEventsRequest: list events request.
type ListEventsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	FipID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: unknown
	OrderBy ListEventsRequestOrderBy `json:"-"`

	OldStatus []flexibleip_v1alpha1.FlexibleIPStatus `json:"-"`

	NewStatus []flexibleip_v1alpha1.FlexibleIPStatus `json:"-"`

	ServerIDs []string `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// ListEventsResponse: list events response.
type ListEventsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Events []*Event `json:"events"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListEventsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListEventsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListEventsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Events = append(r.Events, results.Events...)
	r.TotalCount += uint32(len(results.Events))
	return uint32(len(results.Events)), nil
}

// ListFlexibleIPsRequest: list flexible i ps request.
type ListFlexibleIPsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: sort order of the returned flexible IPs.
	// Default value: created_at_asc
	OrderBy flexibleip_v1alpha1.ListFlexibleIPsRequestOrderBy `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of flexible IPs per page.
	PageSize *uint32 `json:"-"`

	// Tags: filter by tag, only flexible IPs with one or more matching tags will be returned.
	Tags []string `json:"-"`

	// Status: filter by status, only flexible IPs with this status will be returned.
	Status []flexibleip_v1alpha1.FlexibleIPStatus `json:"-"`

	// ServerIDs: filter by server IDs, only flexible IPs with these server IDs will be returned.
	ServerIDs []string `json:"-"`

	// OrganizationID: filter by Organization ID, only flexible IPs from this Organization will be returned.
	OrganizationID *string `json:"-"`

	// ProjectID: filter by Project ID, only flexible IPs from this Project will be returned.
	ProjectID *string `json:"-"`
}

// ListFlexibleIPsResponse: list flexible i ps response.
type ListFlexibleIPsResponse struct {
	TotalCount uint32 `json:"total_count"`

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
	return []scw.Region{}
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
		Path:   "/flexible-ip-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFlexibleIPs:
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
		Path:   "/flexible-ip-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips",
		Query:  query,
	}

	var resp ListFlexibleIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFlexibleIP:
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
		Path:   "/flexible-ip-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "",
	}

	var resp FlexibleIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListEvents:
func (s *API) ListEvents(req *ListEventsRequest, opts ...scw.RequestOption) (*ListEventsResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "old_status", req.OldStatus)
	parameter.AddToQuery(query, "new_status", req.NewStatus)
	parameter.AddToQuery(query, "server_ids", req.ServerIDs)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.FipID) == "" {
		return nil, errors.New("field FipID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/flexible-ip-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/fips/" + fmt.Sprint(req.FipID) + "/events",
		Query:  query,
	}

	var resp ListEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

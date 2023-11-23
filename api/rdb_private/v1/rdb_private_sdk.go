// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package rdb_private provides methods and message types of the rdb_private v1 API.
package rdb_private

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

type InstanceEventResourceType string

const (
	InstanceEventResourceTypeNone         = InstanceEventResourceType("none")
	InstanceEventResourceTypeBackups      = InstanceEventResourceType("backups")
	InstanceEventResourceTypeInstances    = InstanceEventResourceType("instances")
	InstanceEventResourceTypeSnapshots    = InstanceEventResourceType("snapshots")
	InstanceEventResourceTypeReadReplicas = InstanceEventResourceType("read_replicas")
)

func (enum InstanceEventResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "none"
	}
	return string(enum)
}

func (enum InstanceEventResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceEventResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceEventResourceType(InstanceEventResourceType(tmp).String())
	return nil
}

type InstanceEventStatus string

const (
	InstanceEventStatusInfo    = InstanceEventStatus("info")
	InstanceEventStatusSuccess = InstanceEventStatus("success")
	InstanceEventStatusError   = InstanceEventStatus("error")
)

func (enum InstanceEventStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "info"
	}
	return string(enum)
}

func (enum InstanceEventStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceEventStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceEventStatus(InstanceEventStatus(tmp).String())
	return nil
}

type ListInstanceEventsRequestOrderBy string

const (
	ListInstanceEventsRequestOrderByCreatedAtAsc  = ListInstanceEventsRequestOrderBy("created_at_asc")
	ListInstanceEventsRequestOrderByCreatedAtDesc = ListInstanceEventsRequestOrderBy("created_at_desc")
)

func (enum ListInstanceEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListInstanceEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInstanceEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInstanceEventsRequestOrderBy(ListInstanceEventsRequestOrderBy(tmp).String())
	return nil
}

// InstanceEvent: instance event.
type InstanceEvent struct {
	ID string `json:"id"`

	// Status: default value: info
	Status InstanceEventStatus `json:"status"`

	Name string `json:"name"`

	Details *string `json:"details"`

	CreatedAt *time.Time `json:"created_at"`

	ResourceID *string `json:"resource_id"`

	// ResourceType: default value: none
	ResourceType InstanceEventResourceType `json:"resource_type"`
}

// ConsoleAPIGetDashboardRequest: console api get dashboard request.
type ConsoleAPIGetDashboardRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// ConsoleAPIGetInstanceDashboardRequest: console api get instance dashboard request.
type ConsoleAPIGetInstanceDashboardRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`
}

// ConsoleAPIGetServiceInfoRequest: console api get service info request.
type ConsoleAPIGetServiceInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// ConsoleAPIListInstanceEventsRequest: console api list instance events request.
type ConsoleAPIListInstanceEventsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListInstanceEventsRequestOrderBy `json:"-"`
}

// Dashboard: dashboard.
type Dashboard struct {
	InstanceCount int32 `json:"instance_count"`

	BackupCount int32 `json:"backup_count"`

	SnapshotCount int32 `json:"snapshot_count"`
}

// InstanceDashboard: instance dashboard.
type InstanceDashboard struct {
	UserCount int32 `json:"user_count"`

	BackupCount int32 `json:"backup_count"`

	AllowedIPCount int32 `json:"allowed_ip_count"`

	AdvancedSettingsCount int32 `json:"advanced_settings_count"`

	DatabaseCount int32 `json:"database_count"`

	SnapshotCount int32 `json:"snapshot_count"`
}

// ListInstanceEventsResponse: list instance events response.
type ListInstanceEventsResponse struct {
	TotalCount uint32 `json:"total_count"`

	InstanceEvents []*InstanceEvent `json:"instance_events"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInstanceEventsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInstanceEventsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListInstanceEventsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.InstanceEvents = append(r.InstanceEvents, results.InstanceEvents...)
	r.TotalCount += uint32(len(results.InstanceEvents))
	return uint32(len(results.InstanceEvents)), nil
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
func (s *ConsoleAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar}
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
		Path:   "/rdb-private/v1/regions/" + fmt.Sprint(req.Region) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDashboard:
func (s *ConsoleAPI) GetDashboard(req *ConsoleAPIGetDashboardRequest, opts ...scw.RequestOption) (*Dashboard, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb-private/v1/regions/" + fmt.Sprint(req.Region) + "/instances/dashboard",
		Query:  query,
	}

	var resp Dashboard

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInstanceDashboard:
func (s *ConsoleAPI) GetInstanceDashboard(req *ConsoleAPIGetInstanceDashboardRequest, opts ...scw.RequestOption) (*InstanceDashboard, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb-private/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/dashboard",
	}

	var resp InstanceDashboard

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInstanceEvents:
func (s *ConsoleAPI) ListInstanceEvents(req *ConsoleAPIListInstanceEventsRequest, opts ...scw.RequestOption) (*ListInstanceEventsResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb-private/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/events",
		Query:  query,
	}

	var resp ListInstanceEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

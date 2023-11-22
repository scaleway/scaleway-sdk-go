// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package applesilicon_admin provides methods and message types of the applesilicon_admin v1alpha1 API.
package applesilicon_admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/applesilicon_admin/v1alpha1/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/applesilicon_admin/v1alpha1/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/applesilicon_admin/v1alpha1/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/applesilicon_admin/v1alpha1/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/applesilicon_admin/v1alpha1/scw"
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

type ListMacsRequestOrderBy string

const (
	ListMacsRequestOrderByNameAsc = ListMacsRequestOrderBy("name_asc")
)

func (enum ListMacsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListMacsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListMacsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListMacsRequestOrderBy(ListMacsRequestOrderBy(tmp).String())
	return nil
}

type ListServersRequestOrderBy string

const (
	ListServersRequestOrderByCreatedAtAsc  = ListServersRequestOrderBy("created_at_asc")
	ListServersRequestOrderByCreatedAtDesc = ListServersRequestOrderBy("created_at_desc")
)

func (enum ListServersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListServersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServersRequestOrderBy(ListServersRequestOrderBy(tmp).String())
	return nil
}

type ListWorkflowsRequestOrderBy string

const (
	ListWorkflowsRequestOrderByCreatedAtAsc  = ListWorkflowsRequestOrderBy("created_at_asc")
	ListWorkflowsRequestOrderByCreatedAtDesc = ListWorkflowsRequestOrderBy("created_at_desc")
)

func (enum ListWorkflowsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListWorkflowsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListWorkflowsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListWorkflowsRequestOrderBy(ListWorkflowsRequestOrderBy(tmp).String())
	return nil
}

type MacStatus string

const (
	MacStatusUnknownStatus = MacStatus("unknown_status")
	MacStatusStarting      = MacStatus("starting")
	MacStatusReady         = MacStatus("ready")
	MacStatusError         = MacStatus("error")
	MacStatusInstalling    = MacStatus("installing")
	MacStatusInstalled     = MacStatus("installed")
	MacStatusRebooting     = MacStatus("rebooting")
	MacStatusUpdatingAgent = MacStatus("updating_agent")
	MacStatusLocking       = MacStatus("locking")
	MacStatusLocked        = MacStatus("locked")
	MacStatusUnlocking     = MacStatus("unlocking")
	MacStatusConfiguring   = MacStatus("configuring")
	MacStatusReinstalling  = MacStatus("reinstalling")
)

func (enum MacStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum MacStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MacStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MacStatus(MacStatus(tmp).String())
	return nil
}

type WorkflowTaskStatus string

const (
	WorkflowTaskStatusUnknownTaskStatus = WorkflowTaskStatus("unknown_task_status")
	WorkflowTaskStatusPending           = WorkflowTaskStatus("pending")
	WorkflowTaskStatusProcessing        = WorkflowTaskStatus("processing")
	WorkflowTaskStatusError             = WorkflowTaskStatus("error")
	WorkflowTaskStatusCompleted         = WorkflowTaskStatus("completed")
)

func (enum WorkflowTaskStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_task_status"
	}
	return string(enum)
}

func (enum WorkflowTaskStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *WorkflowTaskStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = WorkflowTaskStatus(WorkflowTaskStatus(tmp).String())
	return nil
}

// RelayBoard: relay board.
type RelayBoard struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Mac string `json:"mac"`

	Port uint32 `json:"port"`
}

// Mac: mac.
type Mac struct {
	ID string `json:"id"`

	Name string `json:"name"`

	// Status: default value: unknown_status
	Status MacStatus `json:"status"`

	IP net.IP `json:"ip"`

	MacAddress string `json:"mac_address"`

	SerialNumber string `json:"serial_number"`

	ControllerIP string `json:"controller_ip"`

	ControllerMac string `json:"controller_mac"`

	RelayBoard *RelayBoard `json:"relay_board"`

	RelayBoardRelay uint32 `json:"relay_board_relay"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	LastHeartbeatAt *time.Time `json:"last_heartbeat_at"`

	Maintenance bool `json:"maintenance"`
}

// Server: server.
type Server struct {
	ID string `json:"id"`

	Mac *Mac `json:"mac"`

	Name string `json:"name"`

	ProjectID string `json:"project_id"`

	OrganizationID string `json:"organization_id"`

	StartedAt *time.Time `json:"started_at"`

	BillingStartedAt *time.Time `json:"billing_started_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	StoppedAt *time.Time `json:"stopped_at"`

	DeletableAt *time.Time `json:"deletable_at"`
}

// Workflow: workflow.
type Workflow struct {
	ID string `json:"id"`

	Type string `json:"type"`

	MacID string `json:"mac_id"`

	TaskIndex int32 `json:"task_index"`

	TaskTotalCount int32 `json:"task_total_count"`

	// TaskStatus: default value: unknown_task_status
	TaskStatus WorkflowTaskStatus `json:"task_status"`

	Retry uint32 `json:"retry"`

	ErrorMessage string `json:"error_message"`

	CreatedAt *time.Time `json:"created_at"`

	TaskAcquiredAt *time.Time `json:"task_acquired_at"`

	TaskPlannedAt *time.Time `json:"task_planned_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// CreateServerRequest: create server request.
type CreateServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Name string `json:"name"`

	ProjectID string `json:"project_id"`

	Type string `json:"type"`

	MacID string `json:"mac_id"`
}

// DeleteServerRequest: delete server request.
type DeleteServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

// GetMacRequest: get mac request.
type GetMacRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	MacID string `json:"-"`
}

// GetServerRequest: get server request.
type GetServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

// GetServiceInfoRequest: get service info request.
type GetServiceInfoRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// InstallMacRequest: install mac request.
type InstallMacRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	MacID string `json:"-"`
}

// ListMacsRequest: list macs request.
type ListMacsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListMacsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListMacsResponse: list macs response.
type ListMacsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Macs []*Mac `json:"macs"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListMacsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListMacsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListMacsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Macs = append(r.Macs, results.Macs...)
	r.TotalCount += uint32(len(results.Macs))
	return uint32(len(results.Macs)), nil
}

// ListServersRequest: list servers request.
type ListServersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListServersRequestOrderBy `json:"-"`

	ProjectID *string `json:"-"`

	OrganizationID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	MacID *string `json:"-"`

	CreatedBefore *time.Time `json:"-"`

	DeletedAfter *time.Time `json:"-"`

	IncludeDeleted *bool `json:"-"`
}

// ListServersResponse: list servers response.
type ListServersResponse struct {
	TotalCount uint32 `json:"total_count"`

	Servers []*Server `json:"servers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint32(len(results.Servers))
	return uint32(len(results.Servers)), nil
}

// ListWorkflowsRequest: list workflows request.
type ListWorkflowsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListWorkflowsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	MacID *string `json:"-"`

	ServerID *string `json:"-"`
}

// ListWorkflowsResponse: list workflows response.
type ListWorkflowsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Workflows []*Workflow `json:"workflows"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListWorkflowsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListWorkflowsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListWorkflowsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Workflows = append(r.Workflows, results.Workflows...)
	r.TotalCount += uint32(len(results.Workflows))
	return uint32(len(results.Workflows)), nil
}

// TriggerAgentUpdateRequest: trigger agent update request.
type TriggerAgentUpdateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Precisely one of AllRunningServers, ServerIDs must be set.
	AllRunningServers *bool `json:"all_running_servers,omitempty"`

	// Precisely one of AllRunningServers, ServerIDs must be set.
	ServerIDs *[]string `json:"server_ids,omitempty"`

	AgentURL string `json:"agent_url"`
}

// TriggerAgentUpdateResponse: trigger agent update response.
type TriggerAgentUpdateResponse struct {
	TaskIDs []string `json:"task_ids"`
}

// UpdateMacRequest: update mac request.
type UpdateMacRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	MacID string `json:"-"`

	Maintenance *bool `json:"maintenance,omitempty"`
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

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TriggerAgentUpdate:
func (s *API) TriggerAgentUpdate(req *TriggerAgentUpdateRequest, opts ...scw.RequestOption) (*TriggerAgentUpdateResponse, error) {
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
		Path:   "/apple-silicon-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/trigger-agent-update",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TriggerAgentUpdateResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateServer:
func (s *API) CreateServer(req *CreateServerRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/apple-silicon-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServers:
func (s *API) ListServers(req *ListServersRequest, opts ...scw.RequestOption) (*ListServersResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "mac_id", req.MacID)
	parameter.AddToQuery(query, "created_before", req.CreatedBefore)
	parameter.AddToQuery(query, "deleted_after", req.DeletedAfter)
	parameter.AddToQuery(query, "include_deleted", req.IncludeDeleted)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:  query,
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServer:
func (s *API) GetServer(req *GetServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteServer:
func (s *API) DeleteServer(req *DeleteServerRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/apple-silicon-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListMacs:
func (s *API) ListMacs(req *ListMacsRequest, opts ...scw.RequestOption) (*ListMacsResponse, error) {
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

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/macs",
		Query:  query,
	}

	var resp ListMacsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMac:
func (s *API) GetMac(req *GetMacRequest, opts ...scw.RequestOption) (*Mac, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.MacID) == "" {
		return nil, errors.New("field MacID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/macs/" + fmt.Sprint(req.MacID) + "",
	}

	var resp Mac

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateMac:
func (s *API) UpdateMac(req *UpdateMacRequest, opts ...scw.RequestOption) (*Mac, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.MacID) == "" {
		return nil, errors.New("field MacID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/apple-silicon-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/macs/" + fmt.Sprint(req.MacID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Mac

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InstallMac:
func (s *API) InstallMac(req *InstallMacRequest, opts ...scw.RequestOption) (*Mac, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.MacID) == "" {
		return nil, errors.New("field MacID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/apple-silicon-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/macs/" + fmt.Sprint(req.MacID) + "/install",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Mac

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListWorkflows:
func (s *API) ListWorkflows(req *ListWorkflowsRequest, opts ...scw.RequestOption) (*ListWorkflowsResponse, error) {
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
	parameter.AddToQuery(query, "mac_id", req.MacID)
	parameter.AddToQuery(query, "server_id", req.ServerID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/workflows",
		Query:  query,
	}

	var resp ListWorkflowsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

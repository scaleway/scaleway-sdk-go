// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package smart_labeling provides methods and message types of the smart_labeling v1alpha1 API.
package smart_labeling

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/smart_labeling/v1alpha1/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/smart_labeling/v1alpha1/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/smart_labeling/v1alpha1/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/smart_labeling/v1alpha1/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/smart_labeling/v1alpha1/scw"
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

type ListWorkspacesRequestOrderBy string

const (
	ListWorkspacesRequestOrderByCreatedAtDesc = ListWorkspacesRequestOrderBy("created_at_desc")
	ListWorkspacesRequestOrderByCreatedAtAsc  = ListWorkspacesRequestOrderBy("created_at_asc")
	ListWorkspacesRequestOrderByNameAsc       = ListWorkspacesRequestOrderBy("name_asc")
	ListWorkspacesRequestOrderByNameDesc      = ListWorkspacesRequestOrderBy("name_desc")
)

func (enum ListWorkspacesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_desc"
	}
	return string(enum)
}

func (enum ListWorkspacesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListWorkspacesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListWorkspacesRequestOrderBy(ListWorkspacesRequestOrderBy(tmp).String())
	return nil
}

type WorkspaceStatus string

const (
	WorkspaceStatusUnknownStatus = WorkspaceStatus("unknown_status")
	WorkspaceStatusCreating      = WorkspaceStatus("creating")
	WorkspaceStatusReady         = WorkspaceStatus("ready")
	WorkspaceStatusDeleting      = WorkspaceStatus("deleting")
	WorkspaceStatusError         = WorkspaceStatus("error")
)

func (enum WorkspaceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum WorkspaceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *WorkspaceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = WorkspaceStatus(WorkspaceStatus(tmp).String())
	return nil
}

// Bucket: bucket.
type Bucket struct {
	ID string `json:"id"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// Workspace: workspace.
type Workspace struct {
	ID string `json:"id"`

	ProjectID string `json:"project_id"`

	OrganizationID string `json:"organization_id"`

	// Status: default value: unknown_status
	Status WorkspaceStatus `json:"status"`

	StatusMessage *string `json:"status_message"`

	Name string `json:"name"`

	Description *string `json:"description"`

	Endpoint *string `json:"endpoint"`

	Buckets map[string]*Bucket `json:"buckets"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// CreateWorkspaceRequest: create workspace request.
type CreateWorkspaceRequest struct {
	ProjectID string `json:"project_id"`

	Name string `json:"name"`

	Description *string `json:"description,omitempty"`

	Buckets []*Bucket `json:"buckets"`

	CvatUsername string `json:"cvat_username"`

	CvatEmail string `json:"cvat_email"`

	CvatPassword string `json:"cvat_password"`
}

// DeleteWorkspaceRequest: delete workspace request.
type DeleteWorkspaceRequest struct {
	WorkspaceID string `json:"-"`
}

// GetWorkspaceRequest: get workspace request.
type GetWorkspaceRequest struct {
	WorkspaceID string `json:"-"`
}

// ListWorkspacesRequest: list workspaces request.
type ListWorkspacesRequest struct {
	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_desc
	OrderBy ListWorkspacesRequestOrderBy `json:"-"`
}

// ListWorkspacesResponse: list workspaces response.
type ListWorkspacesResponse struct {
	Workspaces []*Workspace `json:"workspaces"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListWorkspacesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListWorkspacesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListWorkspacesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Workspaces = append(r.Workspaces, results.Workspaces...)
	r.TotalCount += uint32(len(results.Workspaces))
	return uint32(len(results.Workspaces)), nil
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
func (s *API) GetServiceInfo(opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/smart-labeling/v1alpha1",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateWorkspace:
func (s *API) CreateWorkspace(req *CreateWorkspaceRequest, opts ...scw.RequestOption) (*Workspace, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/smart-labeling/v1alpha1/workspaces",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Workspace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWorkspace:
func (s *API) GetWorkspace(req *GetWorkspaceRequest, opts ...scw.RequestOption) (*Workspace, error) {
	var err error

	if fmt.Sprint(req.WorkspaceID) == "" {
		return nil, errors.New("field WorkspaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/smart-labeling/v1alpha1/workspaces/" + fmt.Sprint(req.WorkspaceID) + "",
	}

	var resp Workspace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListWorkspaces:
func (s *API) ListWorkspaces(req *ListWorkspacesRequest, opts ...scw.RequestOption) (*ListWorkspacesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/smart-labeling/v1alpha1/workspaces",
		Query:  query,
	}

	var resp ListWorkspacesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteWorkspace:
func (s *API) DeleteWorkspace(req *DeleteWorkspaceRequest, opts ...scw.RequestOption) (*Workspace, error) {
	var err error

	if fmt.Sprint(req.WorkspaceID) == "" {
		return nil, errors.New("field WorkspaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/smart-labeling/v1alpha1/workspaces/" + fmt.Sprint(req.WorkspaceID) + "",
	}

	var resp Workspace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

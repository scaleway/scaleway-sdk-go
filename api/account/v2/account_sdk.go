// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package account provides methods and message types of the account v2 API.
package account

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

type ListProjectsRequestOrderBy string

const (
	// Creation date ascending.
	ListProjectsRequestOrderByCreatedAtAsc = ListProjectsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListProjectsRequestOrderByCreatedAtDesc = ListProjectsRequestOrderBy("created_at_desc")
	// Name ascending.
	ListProjectsRequestOrderByNameAsc = ListProjectsRequestOrderBy("name_asc")
	// Name descending.
	ListProjectsRequestOrderByNameDesc = ListProjectsRequestOrderBy("name_desc")
)

func (enum ListProjectsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListProjectsRequestOrderBy) Values() []ListProjectsRequestOrderBy {
	return []ListProjectsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListProjectsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListProjectsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListProjectsRequestOrderBy(ListProjectsRequestOrderBy(tmp).String())
	return nil
}

// Project: project.
type Project struct {
	// ID: ID of the Project.
	ID string `json:"id"`

	// Name: name of the Project.
	Name string `json:"name"`

	// OrganizationID: organization ID of the Project.
	OrganizationID string `json:"organization_id"`

	// CreatedAt: creation date of the Project.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: update date of the Project.
	UpdatedAt *time.Time `json:"updated_at"`

	// Description: description of the Project.
	Description string `json:"description"`
}

// CreateProjectRequest: create project request.
type CreateProjectRequest struct {
	// Name: name of the Project.
	Name string `json:"name"`

	// OrganizationID: organization ID of the Project.
	OrganizationID string `json:"organization_id"`

	// Description: description of the Project.
	Description *string `json:"description,omitempty"`
}

// DeleteProjectRequest: delete project request.
type DeleteProjectRequest struct {
	// ProjectID: project ID of the Project.
	ProjectID string `json:"-"`
}

// GetProjectRequest: get project request.
type GetProjectRequest struct {
	// ProjectID: project ID of the Project.
	ProjectID string `json:"-"`
}

// ListProjectsRequest: list projects request.
type ListProjectsRequest struct {
	// OrganizationID: organization ID of the Project.
	OrganizationID string `json:"-"`

	// Name: name of the Project.
	Name *string `json:"-"`

	// Page: page number for the returned Projects.
	Page *int32 `json:"-"`

	// PageSize: maximum number of Project per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned Projects.
	// Default value: created_at_asc
	OrderBy ListProjectsRequestOrderBy `json:"-"`

	// ProjectIDs: project IDs to filter for. The results will be limited to any Projects with an ID in this array.
	ProjectIDs []string `json:"-"`
}

// ListProjectsResponse: list projects response.
type ListProjectsResponse struct {
	// TotalCount: total number of Projects.
	TotalCount uint32 `json:"total_count"`

	// Projects: paginated returned Projects.
	Projects []*Project `json:"projects"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListProjectsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListProjectsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListProjectsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Projects = append(r.Projects, results.Projects...)
	r.TotalCount += uint32(len(results.Projects))
	return uint32(len(results.Projects)), nil
}

// UpdateProjectRequest: update project request.
type UpdateProjectRequest struct {
	// ProjectID: project ID of the Project.
	ProjectID string `json:"-"`

	// Name: name of the Project.
	Name *string `json:"name,omitempty"`

	// Description: description of the Project.
	Description *string `json:"description,omitempty"`
}

// This API allows you to manage your Scaleway Projects.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// Deprecated: CreateProject: Deprecated in favor of Account API v3.
// Generate a new Project for an Organization, specifying its configuration including name and description.
func (s *API) CreateProject(req *CreateProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("proj")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/projects",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: ListProjects: Deprecated in favor of Account API v3.
// List all Projects of an Organization. The response will include the total number of Projects as well as their associated Organizations, names and IDs. Other information include the creation and update date of the Project.
func (s *API) ListProjects(req *ListProjectsRequest, opts ...scw.RequestOption) (*ListProjectsResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_ids", req.ProjectIDs)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v2/projects",
		Query:  query,
	}

	var resp ListProjectsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: GetProject: Deprecated in favor of Account API v3.
// Retrieve information about an existing Project, specified by its Project ID. Its full details, including ID, name and description, are returned in the response object.
func (s *API) GetProject(req *GetProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return nil, errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v2/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: DeleteProject: Deprecated in favor of Account API v3.
// Delete an existing Project, specified by its Project ID. The Project needs to be empty (meaning there are no resources left in it) to be deleted effectively. Note that deleting a Project is permanent, and cannot be undone.
func (s *API) DeleteProject(req *DeleteProjectRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account/v2/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: UpdateProject: Deprecated in favor of Account API v3.
// Update the parameters of an existing Project, specified by its Project ID. These parameters include the name and description.
func (s *API) UpdateProject(req *UpdateProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return nil, errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account/v2/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

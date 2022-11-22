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

// API: user related data
//
// This API allows you to manage projects.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type ListProjectsRequestOrderBy string

const (
	// ListProjectsRequestOrderByCreatedAtAsc is [insert doc].
	ListProjectsRequestOrderByCreatedAtAsc = ListProjectsRequestOrderBy("created_at_asc")
	// ListProjectsRequestOrderByCreatedAtDesc is [insert doc].
	ListProjectsRequestOrderByCreatedAtDesc = ListProjectsRequestOrderBy("created_at_desc")
	// ListProjectsRequestOrderByNameAsc is [insert doc].
	ListProjectsRequestOrderByNameAsc = ListProjectsRequestOrderBy("name_asc")
	// ListProjectsRequestOrderByNameDesc is [insert doc].
	ListProjectsRequestOrderByNameDesc = ListProjectsRequestOrderBy("name_desc")
)

func (enum ListProjectsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
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

// ListProjectsResponse: list projects response
type ListProjectsResponse struct {
	// TotalCount: the total number of projects
	TotalCount uint32 `json:"total_count"`
	// Projects: the paginated returned projects
	Projects []*Project `json:"projects"`
}

// Project: project
type Project struct {
	// ID: the ID of the project
	ID string `json:"id"`
	// Name: the name of the project
	Name string `json:"name"`
	// OrganizationID: the organization ID of the project
	OrganizationID string `json:"organization_id"`
	// CreatedAt: the creation date of the project
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: the update date of the project
	UpdatedAt *time.Time `json:"updated_at"`
	// Description: the description of the project
	Description string `json:"description"`
}

// Service API

type CreateProjectRequest struct {
	// Name: the name of the project
	Name string `json:"name"`
	// OrganizationID: the organization ID of the project
	OrganizationID string `json:"organization_id"`
	// Description: the description of the project
	Description *string `json:"description"`
}

// CreateProject: create project
func (s *API) CreateProject(req *CreateProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/account/v2/projects",
		Headers: http.Header{},
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

type ListProjectsRequest struct {
	// OrganizationID: the organization ID of the project
	OrganizationID string `json:"-"`
	// Name: the name of the project
	Name *string `json:"-"`
	// Page: the page number for the returned projects
	Page *int32 `json:"-"`
	// PageSize: the maximum number of project per page
	PageSize *uint32 `json:"-"`
	// OrderBy: the sort order of the returned projects
	//
	// Default value: created_at_asc
	OrderBy ListProjectsRequestOrderBy `json:"-"`
	// ProjectIDs: filter out by a list of project ID
	ProjectIDs []string `json:"-"`
}

// ListProjects: list projects
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
		Method:  "GET",
		Path:    "/account/v2/projects",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListProjectsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetProjectRequest struct {
	// ProjectID: the project ID of the project
	ProjectID string `json:"-"`
}

// GetProject: get project
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
		Method:  "GET",
		Path:    "/account/v2/projects/" + fmt.Sprint(req.ProjectID) + "",
		Headers: http.Header{},
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteProjectRequest struct {
	// ProjectID: the project ID of the project
	ProjectID string `json:"-"`
}

// DeleteProject: delete project
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
		Method:  "DELETE",
		Path:    "/account/v2/projects/" + fmt.Sprint(req.ProjectID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type UpdateProjectRequest struct {
	// ProjectID: the project ID of the project
	ProjectID string `json:"-"`
	// Name: the name of the project
	Name *string `json:"name"`
	// Description: the description of the project
	Description *string `json:"description"`
}

// UpdateProject: update project
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
		Method:  "PATCH",
		Path:    "/account/v2/projects/" + fmt.Sprint(req.ProjectID) + "",
		Headers: http.Header{},
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

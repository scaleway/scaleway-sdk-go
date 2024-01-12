// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package serverless_jobs_admin provides methods and message types of the serverless_jobs_admin v1 API.
package serverless_jobs_admin

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

type JobDefinitionStatus string

const (
	JobDefinitionStatusUnknownStatus = JobDefinitionStatus("unknown_status")
	JobDefinitionStatusUnlocked      = JobDefinitionStatus("unlocked")
	JobDefinitionStatusLocked        = JobDefinitionStatus("locked")
)

func (enum JobDefinitionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum JobDefinitionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *JobDefinitionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = JobDefinitionStatus(JobDefinitionStatus(tmp).String())
	return nil
}

type JobRunStatus string

const (
	JobRunStatusUnknownStatus = JobRunStatus("unknown_status")
	JobRunStatusQueued        = JobRunStatus("queued")
	JobRunStatusScheduled     = JobRunStatus("scheduled")
	JobRunStatusRunning       = JobRunStatus("running")
	JobRunStatusSucceeded     = JobRunStatus("succeeded")
	JobRunStatusFailed        = JobRunStatus("failed")
	JobRunStatusCanceled      = JobRunStatus("canceled")
)

func (enum JobRunStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum JobRunStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *JobRunStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = JobRunStatus(JobRunStatus(tmp).String())
	return nil
}

type ListJobDefinitionsRequestOrderBy string

const (
	ListJobDefinitionsRequestOrderByCreatedAtAsc  = ListJobDefinitionsRequestOrderBy("created_at_asc")
	ListJobDefinitionsRequestOrderByCreatedAtDesc = ListJobDefinitionsRequestOrderBy("created_at_desc")
)

func (enum ListJobDefinitionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListJobDefinitionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListJobDefinitionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListJobDefinitionsRequestOrderBy(ListJobDefinitionsRequestOrderBy(tmp).String())
	return nil
}

type ListJobRunsRequestOrderBy string

const (
	ListJobRunsRequestOrderByCreatedAtAsc  = ListJobRunsRequestOrderBy("created_at_asc")
	ListJobRunsRequestOrderByCreatedAtDesc = ListJobRunsRequestOrderBy("created_at_desc")
)

func (enum ListJobRunsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListJobRunsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListJobRunsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListJobRunsRequestOrderBy(ListJobRunsRequestOrderBy(tmp).String())
	return nil
}

// JobDefinition: job definition.
type JobDefinition struct {
	ID string `json:"id"`

	Name string `json:"name"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	MvcpuLimit uint32 `json:"mvcpu_limit"`

	MemoryLimitMegabytes uint32 `json:"memory_limit_megabytes"`

	ImageURI string `json:"image_uri"`

	Command string `json:"command"`

	ProjectID string `json:"project_id"`

	EnvironmentVariablesKeys []string `json:"environment_variables_keys"`

	Description string `json:"description"`

	JobTimeout *scw.Duration `json:"job_timeout"`

	OrganizationID string `json:"organization_id"`

	// Status: default value: unknown_status
	Status JobDefinitionStatus `json:"status"`

	LockReason string `json:"lock_reason"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// JobRun: job run.
type JobRun struct {
	ID string `json:"id"`

	JobDefinitionID string `json:"job_definition_id"`

	// Status: default value: unknown_status
	Status JobRunStatus `json:"status"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	TerminatedAt *time.Time `json:"terminated_at"`

	ExitCode *int32 `json:"exit_code"`

	RunDuration *scw.Duration `json:"run_duration"`

	ErrorMessage string `json:"error_message"`

	K8sNamespace string `json:"k8s_namespace"`

	MvcpuLimit uint32 `json:"mvcpu_limit"`

	MemoryLimitMegabytes uint32 `json:"memory_limit_megabytes"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// GetJobDefinitionRequest: get job definition request.
type GetJobDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	JobDefinitionID string `json:"-"`
}

// GetJobRunRequest: get job run request.
type GetJobRunRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	JobRunID string `json:"-"`
}

// GetServiceInfoRequest: get service info request.
type GetServiceInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// ListJobDefinitionsRequest: list job definitions request.
type ListJobDefinitionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListJobDefinitionsRequestOrderBy `json:"-"`

	ProjectID *string `json:"-"`

	OrganizationID *string `json:"-"`

	// Status: default value: unknown_status
	Status JobDefinitionStatus `json:"-"`
}

// ListJobDefinitionsResponse: list job definitions response.
type ListJobDefinitionsResponse struct {
	Definitions []*JobDefinition `json:"definitions"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListJobDefinitionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListJobDefinitionsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListJobDefinitionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Definitions = append(r.Definitions, results.Definitions...)
	r.TotalCount += uint64(len(results.Definitions))
	return uint64(len(results.Definitions)), nil
}

// ListJobRunsRequest: list job runs request.
type ListJobRunsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListJobRunsRequestOrderBy `json:"-"`

	JobDefinitionID *string `json:"-"`

	ProjectID *string `json:"-"`

	OrganizationID *string `json:"-"`

	// Status: default value: unknown_status
	Status JobRunStatus `json:"-"`
}

// ListJobRunsResponse: list job runs response.
type ListJobRunsResponse struct {
	Runs []*JobRun `json:"runs"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListJobRunsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListJobRunsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListJobRunsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Runs = append(r.Runs, results.Runs...)
	r.TotalCount += uint64(len(results.Runs))
	return uint64(len(results.Runs)), nil
}

// This API is for the admin console of Serverless Jobs.
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

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-jobs-admin/v1/regions/" + fmt.Sprint(req.Region) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetJobDefinition:
func (s *API) GetJobDefinition(req *GetJobDefinitionRequest, opts ...scw.RequestOption) (*JobDefinition, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.JobDefinitionID) == "" {
		return nil, errors.New("field JobDefinitionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-jobs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "",
	}

	var resp JobDefinition

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListJobDefinitions:
func (s *API) ListJobDefinitions(req *ListJobDefinitionsRequest, opts ...scw.RequestOption) (*ListJobDefinitionsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-jobs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/job-definitions",
		Query:  query,
	}

	var resp ListJobDefinitionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetJobRun:
func (s *API) GetJobRun(req *GetJobRunRequest, opts ...scw.RequestOption) (*JobRun, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.JobRunID) == "" {
		return nil, errors.New("field JobRunID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-jobs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/job-runs/" + fmt.Sprint(req.JobRunID) + "",
	}

	var resp JobRun

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListJobRuns:
func (s *API) ListJobRuns(req *ListJobRunsRequest, opts ...scw.RequestOption) (*ListJobRunsResponse, error) {
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
	parameter.AddToQuery(query, "job_definition_id", req.JobDefinitionID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-jobs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/job-runs",
		Query:  query,
	}

	var resp ListJobRunsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

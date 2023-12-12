// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package jobs provides methods and message types of the jobs v1alpha1 API.
package jobs

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

type JobRunState string

const (
	JobRunStateUnknownState = JobRunState("unknown_state")
	JobRunStateQueued       = JobRunState("queued")
	JobRunStateScheduled    = JobRunState("scheduled")
	JobRunStateRunning      = JobRunState("running")
	JobRunStateSucceeded    = JobRunState("succeeded")
	JobRunStateFailed       = JobRunState("failed")
	JobRunStateCanceled     = JobRunState("canceled")
)

func (enum JobRunState) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_state"
	}
	return string(enum)
}

func (enum JobRunState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *JobRunState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = JobRunState(JobRunState(tmp).String())
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

	CPULimit uint32 `json:"cpu_limit"`

	MemoryLimit uint32 `json:"memory_limit"`

	ImageURI string `json:"image_uri"`

	Command string `json:"command"`

	ProjectID string `json:"project_id"`

	EnvironmentVariables map[string]string `json:"environment_variables"`

	Description string `json:"description"`

	JobTimeout *scw.Duration `json:"job_timeout"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// JobRun: job run.
type JobRun struct {
	ID string `json:"id"`

	JobDefinitionID string `json:"job_definition_id"`

	// State: default value: unknown_state
	State JobRunState `json:"state"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	TerminatedAt *time.Time `json:"terminated_at"`

	ExitCode *int32 `json:"exit_code"`

	RunDuration *scw.Duration `json:"run_duration"`

	ErrorMessage string `json:"error_message"`

	CPULimit uint32 `json:"cpu_limit"`

	MemoryLimit uint32 `json:"memory_limit"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// CreateJobDefinitionRequest: create job definition request.
type CreateJobDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Name string `json:"name"`

	CPULimit uint32 `json:"cpu_limit"`

	MemoryLimit uint32 `json:"memory_limit"`

	ImageURI string `json:"image_uri"`

	Command string `json:"command"`

	ProjectID string `json:"project_id"`

	EnvironmentVariables map[string]string `json:"environment_variables"`

	Description string `json:"description"`

	JobTimeout *scw.Duration `json:"job_timeout,omitempty"`
}

// DeleteJobDefinitionRequest: delete job definition request.
type DeleteJobDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	JobDefinitionID string `json:"-"`
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

// ListJobDefinitionsRequest: list job definitions request.
type ListJobDefinitionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListJobDefinitionsRequestOrderBy `json:"-"`

	ProjectID *string `json:"-"`
}

// ListJobDefinitionsResponse: list job definitions response.
type ListJobDefinitionsResponse struct {
	JobDefinitions []*JobDefinition `json:"job_definitions"`

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

	r.JobDefinitions = append(r.JobDefinitions, results.JobDefinitions...)
	r.TotalCount += uint64(len(results.JobDefinitions))
	return uint64(len(results.JobDefinitions)), nil
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
}

// ListJobRunsResponse: list job runs response.
type ListJobRunsResponse struct {
	JobRuns []*JobRun `json:"job_runs"`

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

	r.JobRuns = append(r.JobRuns, results.JobRuns...)
	r.TotalCount += uint64(len(results.JobRuns))
	return uint64(len(results.JobRuns)), nil
}

// StartJobDefinitionRequest: start job definition request.
type StartJobDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	JobDefinitionID string `json:"-"`
}

// StopJobRunRequest: stop job run request.
type StopJobRunRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	JobRunID string `json:"-"`
}

// UpdateJobDefinitionRequest: update job definition request.
type UpdateJobDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	JobDefinitionID string `json:"-"`

	Name *string `json:"name,omitempty"`

	CPULimit *uint32 `json:"cpu_limit,omitempty"`

	MemoryLimit *uint32 `json:"memory_limit,omitempty"`

	ImageURI *string `json:"image_uri,omitempty"`

	Command *string `json:"command,omitempty"`

	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	Description *string `json:"description,omitempty"`

	JobTimeout *scw.Duration `json:"job_timeout,omitempty"`
}

// Serverless Jobs API.
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

// CreateJobDefinition:
func (s *API) CreateJobDefinition(req *CreateJobDefinitionRequest, opts ...scw.RequestOption) (*JobDefinition, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("job")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-definitions",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp JobDefinition

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
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "",
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-definitions",
		Query:  query,
	}

	var resp ListJobDefinitionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateJobDefinition:
func (s *API) UpdateJobDefinition(req *UpdateJobDefinitionRequest, opts ...scw.RequestOption) (*JobDefinition, error) {
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
		Method: "PATCH",
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp JobDefinition

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteJobDefinition:
func (s *API) DeleteJobDefinition(req *DeleteJobDefinitionRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.JobDefinitionID) == "" {
		return errors.New("field JobDefinitionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// StartJobDefinition:
func (s *API) StartJobDefinition(req *StartJobDefinitionRequest, opts ...scw.RequestOption) (*JobRun, error) {
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
		Method: "POST",
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "/start",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp JobRun

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
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-runs/" + fmt.Sprint(req.JobRunID) + "",
	}

	var resp JobRun

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopJobRun:
func (s *API) StopJobRun(req *StopJobRunRequest, opts ...scw.RequestOption) (*JobRun, error) {
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
		Method: "POST",
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-runs/" + fmt.Sprint(req.JobRunID) + "/stop",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-runs",
		Query:  query,
	}

	var resp ListJobRunsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

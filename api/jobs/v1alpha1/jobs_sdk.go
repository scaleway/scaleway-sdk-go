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

type JobRunState string

const (
	JobRunStateUnknownState  = JobRunState("unknown_state")
	JobRunStateQueued        = JobRunState("queued")
	JobRunStateScheduled     = JobRunState("scheduled")
	JobRunStateRunning       = JobRunState("running")
	JobRunStateSucceeded     = JobRunState("succeeded")
	JobRunStateFailed        = JobRunState("failed")
	JobRunStateCanceled      = JobRunState("canceled")
	JobRunStateInternalError = JobRunState("internal_error")
)

func (enum JobRunState) String() string {
	if enum == "" {
		// return default value if empty
		return string(JobRunStateUnknownState)
	}
	return string(enum)
}

func (enum JobRunState) Values() []JobRunState {
	return []JobRunState{
		"unknown_state",
		"queued",
		"scheduled",
		"running",
		"succeeded",
		"failed",
		"canceled",
		"internal_error",
	}
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
		return string(ListJobDefinitionsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListJobDefinitionsRequestOrderBy) Values() []ListJobDefinitionsRequestOrderBy {
	return []ListJobDefinitionsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
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
		return string(ListJobRunsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListJobRunsRequestOrderBy) Values() []ListJobRunsRequestOrderBy {
	return []ListJobRunsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
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

// SecretEnvVar: secret env var.
type SecretEnvVar struct {
	Name string `json:"name"`
}

// SecretFile: secret file.
type SecretFile struct {
	Path string `json:"path"`
}

// CronSchedule: cron schedule.
type CronSchedule struct {
	// Schedule: uNIX cron schedule to run job (e.g., '* * * * *').
	Schedule string `json:"schedule"`

	// Timezone: timezone for the cron schedule, in tz database format (e.g., 'Europe/Paris').
	Timezone string `json:"timezone"`
}

// CreateJobDefinitionRequestCronScheduleConfig: create job definition request cron schedule config.
type CreateJobDefinitionRequestCronScheduleConfig struct {
	Schedule string `json:"schedule"`

	Timezone string `json:"timezone"`
}

// CreateJobDefinitionSecretsRequestSecretConfig: create job definition secrets request secret config.
type CreateJobDefinitionSecretsRequestSecretConfig struct {
	SecretManagerID string `json:"secret_manager_id"`

	SecretManagerVersion string `json:"secret_manager_version"`

	// Precisely one of Path, EnvVarName must be set.
	Path *string `json:"path,omitempty"`

	// Precisely one of Path, EnvVarName must be set.
	EnvVarName *string `json:"env_var_name,omitempty"`
}

// Secret: secret.
type Secret struct {
	// SecretID: UUID of the secret reference within the job.
	SecretID string `json:"secret_id"`

	// SecretManagerID: UUID of the secret in Secret Manager.
	SecretManagerID string `json:"secret_manager_id"`

	// SecretManagerVersion: version of the secret in Secret Manager.
	SecretManagerVersion string `json:"secret_manager_version"`

	// File: file secret mounted inside the job.
	// Precisely one of File, EnvVar must be set.
	File *SecretFile `json:"file,omitempty"`

	// EnvVar: environment variable used to expose the secret.
	// Precisely one of File, EnvVar must be set.
	EnvVar *SecretEnvVar `json:"env_var,omitempty"`
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

	CronSchedule *CronSchedule `json:"cron_schedule"`

	LocalStorageCapacity uint32 `json:"local_storage_capacity"`

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

	Command string `json:"command"`

	EnvironmentVariables map[string]string `json:"environment_variables"`

	LocalStorageCapacity uint32 `json:"local_storage_capacity"`

	StartedAt *time.Time `json:"started_at"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// Resource: resource.
type Resource struct {
	CPULimit uint32 `json:"cpu_limit"`

	MemoryLimit uint32 `json:"memory_limit"`
}

// UpdateJobDefinitionRequestCronScheduleConfig: update job definition request cron schedule config.
type UpdateJobDefinitionRequestCronScheduleConfig struct {
	Schedule *string `json:"schedule"`

	Timezone *string `json:"timezone"`
}

// CreateJobDefinitionRequest: create job definition request.
type CreateJobDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: name of the job definition.
	Name string `json:"name"`

	// CPULimit: CPU limit of the job.
	CPULimit uint32 `json:"cpu_limit"`

	// MemoryLimit: memory limit of the job (in MiB).
	MemoryLimit uint32 `json:"memory_limit"`

	// LocalStorageCapacity: local storage capacity of the job (in MiB).
	LocalStorageCapacity *uint32 `json:"local_storage_capacity,omitempty"`

	// ImageURI: image to use for the job.
	ImageURI string `json:"image_uri"`

	// Command: startup command. If empty or not defined, the image's default command is used.
	Command string `json:"command"`

	// ProjectID: UUID of the Scaleway Project containing the job.
	ProjectID string `json:"project_id"`

	// EnvironmentVariables: environment variables of the job.
	EnvironmentVariables map[string]string `json:"environment_variables"`

	// Description: description of the job.
	Description string `json:"description"`

	// JobTimeout: timeout of the job in seconds.
	JobTimeout *scw.Duration `json:"job_timeout,omitempty"`

	// CronSchedule: configure a cron for the job.
	CronSchedule *CreateJobDefinitionRequestCronScheduleConfig `json:"cron_schedule,omitempty"`
}

// CreateJobDefinitionSecretsRequest: create job definition secrets request.
type CreateJobDefinitionSecretsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobDefinitionID: UUID of the job definition.
	JobDefinitionID string `json:"-"`

	// Secrets: list of secrets to inject into the job.
	Secrets []*CreateJobDefinitionSecretsRequestSecretConfig `json:"secrets"`
}

// CreateJobDefinitionSecretsResponse: create job definition secrets response.
type CreateJobDefinitionSecretsResponse struct {
	// Secrets: list of secrets created.
	Secrets []*Secret `json:"secrets"`
}

// DeleteJobDefinitionRequest: delete job definition request.
type DeleteJobDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobDefinitionID: UUID of the job definition to delete.
	JobDefinitionID string `json:"-"`
}

// DeleteJobDefinitionSecretRequest: delete job definition secret request.
type DeleteJobDefinitionSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobDefinitionID: UUID of the job definition.
	JobDefinitionID string `json:"-"`

	// SecretID: UUID of the secret reference within the job.
	SecretID string `json:"-"`
}

// GetJobDefinitionRequest: get job definition request.
type GetJobDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobDefinitionID: UUID of the job definition to get.
	JobDefinitionID string `json:"-"`
}

// GetJobDefinitionSecretRequest: get job definition secret request.
type GetJobDefinitionSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobDefinitionID: UUID of the job definition.
	JobDefinitionID string `json:"-"`

	// SecretID: UUID of the secret reference within the job.
	SecretID string `json:"-"`
}

// GetJobRunRequest: get job run request.
type GetJobRunRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobRunID: UUID of the job run to get.
	JobRunID string `json:"-"`
}

// GetJobsLimitsRequest: get jobs limits request.
type GetJobsLimitsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// JobsLimits: jobs limits.
type JobsLimits struct {
	SecretsPerJobDefinition uint32 `json:"secrets_per_job_definition"`
}

// ListJobDefinitionSecretsRequest: list job definition secrets request.
type ListJobDefinitionSecretsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobDefinitionID: UUID of the job definition.
	JobDefinitionID string `json:"-"`
}

// ListJobDefinitionSecretsResponse: list job definition secrets response.
type ListJobDefinitionSecretsResponse struct {
	// Secrets: list of secret references within a job definition.
	Secrets []*Secret `json:"secrets"`

	// TotalCount: total count of secret references within a job definition.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListJobDefinitionSecretsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListJobDefinitionSecretsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListJobDefinitionSecretsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Secrets = append(r.Secrets, results.Secrets...)
	r.TotalCount += uint64(len(results.Secrets))
	return uint64(len(results.Secrets)), nil
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
func (r *ListJobDefinitionsResponse) UnsafeAppend(res any) (uint64, error) {
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

	OrganizationID *string `json:"-"`

	// State: default value: unknown_state
	State JobRunState `json:"-"`

	States []JobRunState `json:"-"`
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
func (r *ListJobRunsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListJobRunsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.JobRuns = append(r.JobRuns, results.JobRuns...)
	r.TotalCount += uint64(len(results.JobRuns))
	return uint64(len(results.JobRuns)), nil
}

// ListJobsResourcesRequest: list jobs resources request.
type ListJobsResourcesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// ListJobsResourcesResponse: list jobs resources response.
type ListJobsResourcesResponse struct {
	Resources []*Resource `json:"resources"`
}

// StartJobDefinitionRequest: start job definition request.
type StartJobDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobDefinitionID: UUID of the job definition to start.
	JobDefinitionID string `json:"-"`

	// Command: contextual startup command for this specific job run.
	Command *string `json:"command,omitempty"`

	// EnvironmentVariables: contextual environment variables for this specific job run.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// Replicas: number of jobs to run.
	Replicas *uint32 `json:"replicas,omitempty"`
}

// StartJobDefinitionResponse: start job definition response.
type StartJobDefinitionResponse struct {
	JobRuns []*JobRun `json:"job_runs"`
}

// StopJobRunRequest: stop job run request.
type StopJobRunRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobRunID: UUID of the job run to stop.
	JobRunID string `json:"-"`
}

// UpdateJobDefinitionRequest: update job definition request.
type UpdateJobDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobDefinitionID: UUID of the job definition to update.
	JobDefinitionID string `json:"-"`

	// Name: name of the job definition.
	Name *string `json:"name,omitempty"`

	// CPULimit: CPU limit of the job.
	CPULimit *uint32 `json:"cpu_limit,omitempty"`

	// MemoryLimit: memory limit of the job (in MiB).
	MemoryLimit *uint32 `json:"memory_limit,omitempty"`

	// LocalStorageCapacity: local storage capacity of the job (in MiB).
	LocalStorageCapacity *uint32 `json:"local_storage_capacity,omitempty"`

	// ImageURI: image to use for the job.
	ImageURI *string `json:"image_uri,omitempty"`

	// Command: startup command.
	Command *string `json:"command,omitempty"`

	// EnvironmentVariables: environment variables of the job.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// Description: description of the job.
	Description *string `json:"description,omitempty"`

	// JobTimeout: timeout of the job in seconds.
	JobTimeout *scw.Duration `json:"job_timeout,omitempty"`

	CronSchedule *UpdateJobDefinitionRequestCronScheduleConfig `json:"cron_schedule,omitempty"`
}

// UpdateJobDefinitionSecretRequest: update job definition secret request.
type UpdateJobDefinitionSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobDefinitionID: UUID of the job definition.
	JobDefinitionID string `json:"-"`

	// SecretID: UUID of the secret reference within the job.
	SecretID string `json:"-"`

	// SecretManagerVersion: version of the secret in Secret Manager.
	SecretManagerVersion *string `json:"secret_manager_version,omitempty"`

	// Path: path of the secret to mount inside the job (either `path` or `env_var_name` must be set).
	// Precisely one of Path, EnvVarName must be set.
	Path *string `json:"path,omitempty"`

	// EnvVarName: environment variable name used to expose the secret inside the job (either `path` or `env_var_name` must be set).
	// Precisely one of Path, EnvVarName must be set.
	EnvVarName *string `json:"env_var_name,omitempty"`
}

// This API allows you to manage your Serverless Jobs.
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

// CreateJobDefinition: Create a new job definition in a specified Project.
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

// GetJobDefinition: Get a job definition by its unique identifier.
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

// ListJobDefinitions: List all your job definitions with filters.
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

// UpdateJobDefinition: Update an existing job definition associated with the specified unique identifier.
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

// DeleteJobDefinition: Delete an exsisting job definition by its unique identifier.
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

// StartJobDefinition: Run an existing job definition by its unique identifier. This will create a new job run.
func (s *API) StartJobDefinition(req *StartJobDefinitionRequest, opts ...scw.RequestOption) (*StartJobDefinitionResponse, error) {
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

	var resp StartJobDefinitionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateJobDefinitionSecrets: Create a secret reference within a job definition.
func (s *API) CreateJobDefinitionSecrets(req *CreateJobDefinitionSecretsRequest, opts ...scw.RequestOption) (*CreateJobDefinitionSecretsResponse, error) {
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
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "/secrets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateJobDefinitionSecretsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetJobDefinitionSecret: Get a secret references within a job definition.
func (s *API) GetJobDefinitionSecret(req *GetJobDefinitionSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
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

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListJobDefinitionSecrets: List secrets references within a job definition.
func (s *API) ListJobDefinitionSecrets(req *ListJobDefinitionSecretsRequest, opts ...scw.RequestOption) (*ListJobDefinitionSecretsResponse, error) {
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
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "/secrets",
	}

	var resp ListJobDefinitionSecretsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateJobDefinitionSecret: Update a secret reference within a job definition.
func (s *API) UpdateJobDefinitionSecret(req *UpdateJobDefinitionSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
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

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteJobDefinitionSecret: Delete a secret reference within a job definition.
func (s *API) DeleteJobDefinitionSecret(req *DeleteJobDefinitionSecretRequest, opts ...scw.RequestOption) error {
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

	if fmt.Sprint(req.SecretID) == "" {
		return errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetJobRun: Get a job run by its unique identifier.
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

// StopJobRun: Stop a job run by its unique identifier.
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

// ListJobRuns: List all job runs with filters.
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
	parameter.AddToQuery(query, "state", req.State)
	parameter.AddToQuery(query, "states", req.States)

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

// ListJobsResources: List jobs resources for the console.
func (s *API) ListJobsResources(req *ListJobsResourcesRequest, opts ...scw.RequestOption) (*ListJobsResourcesResponse, error) {
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
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/jobs-resources",
	}

	var resp ListJobsResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetJobsLimits: Get jobs limits for the console.
func (s *API) GetJobsLimits(req *GetJobsLimitsRequest, opts ...scw.RequestOption) (*JobsLimits, error) {
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
		Path:   "/serverless-jobs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/jobs-limits",
	}

	var resp JobsLimits

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

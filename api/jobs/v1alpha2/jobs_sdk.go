// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package jobs provides methods and message types of the jobs v1alpha2 API.
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

type JobRunReason string

const (
	JobRunReasonUnknownReason        = JobRunReason("unknown_reason")
	JobRunReasonInvalidRequest       = JobRunReason("invalid_request")
	JobRunReasonTimeout              = JobRunReason("timeout")
	JobRunReasonCancellation         = JobRunReason("cancellation")
	JobRunReasonTechnicalError       = JobRunReason("technical_error")
	JobRunReasonImageNotFound        = JobRunReason("image_not_found")
	JobRunReasonInvalidImage         = JobRunReason("invalid_image")
	JobRunReasonMemoryUsageExceeded  = JobRunReason("memory_usage_exceeded")
	JobRunReasonStorageUsageExceeded = JobRunReason("storage_usage_exceeded")
	JobRunReasonExitedWithError      = JobRunReason("exited_with_error")
	JobRunReasonSecretDisabled       = JobRunReason("secret_disabled")
	JobRunReasonSecretNotFound       = JobRunReason("secret_not_found")
	JobRunReasonQuotaExceeded        = JobRunReason("quota_exceeded")
)

func (enum JobRunReason) String() string {
	if enum == "" {
		// return default value if empty
		return string(JobRunReasonUnknownReason)
	}
	return string(enum)
}

func (enum JobRunReason) Values() []JobRunReason {
	return []JobRunReason{
		"unknown_reason",
		"invalid_request",
		"timeout",
		"cancellation",
		"technical_error",
		"image_not_found",
		"invalid_image",
		"memory_usage_exceeded",
		"storage_usage_exceeded",
		"exited_with_error",
		"secret_disabled",
		"secret_not_found",
		"quota_exceeded",
	}
}

func (enum JobRunReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *JobRunReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = JobRunReason(JobRunReason(tmp).String())
	return nil
}

type JobRunState string

const (
	JobRunStateUnknownState = JobRunState("unknown_state")
	JobRunStateInitialized  = JobRunState("initialized")
	JobRunStateValidated    = JobRunState("validated")
	JobRunStateQueued       = JobRunState("queued")
	JobRunStateRunning      = JobRunState("running")
	JobRunStateSucceeded    = JobRunState("succeeded")
	JobRunStateFailed       = JobRunState("failed")
	JobRunStateInterrupting = JobRunState("interrupting")
	JobRunStateInterrupted  = JobRunState("interrupted")
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
		"initialized",
		"validated",
		"queued",
		"running",
		"succeeded",
		"failed",
		"interrupting",
		"interrupted",
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

// CreateSecretsRequestSecretConfig: create secrets request secret config.
type CreateSecretsRequestSecretConfig struct {
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

	// JobDefinitionID: UUID of the job definition.
	JobDefinitionID string `json:"job_definition_id"`

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

	ProjectID string `json:"project_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	CPULimit uint32 `json:"cpu_limit"`

	MemoryLimit uint32 `json:"memory_limit"`

	LocalStorageCapacity uint32 `json:"local_storage_capacity"`

	ImageURI string `json:"image_uri"`

	// Deprecated
	Command *string `json:"command,omitempty"`

	EnvironmentVariables map[string]string `json:"environment_variables"`

	JobTimeout *scw.Duration `json:"job_timeout"`

	Description string `json:"description"`

	CronSchedule *CronSchedule `json:"cron_schedule"`

	StartupCommand []string `json:"startup_command"`

	Args []string `json:"args"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// Resource: resource.
type Resource struct {
	ComputeLimitMvcpu uint32 `json:"compute_limit_mvcpu"`

	MemoryLimitBytes scw.Size `json:"memory_limit_bytes"`
}

// JobRun: job run.
type JobRun struct {
	ID string `json:"id"`

	JobDefinitionID string `json:"job_definition_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	StartedAt *time.Time `json:"started_at"`

	TerminatedAt *time.Time `json:"terminated_at"`

	RunDuration *scw.Duration `json:"run_duration"`

	// State: default value: unknown_state
	State JobRunState `json:"state"`

	// Reason: default value: unknown_reason
	Reason *JobRunReason `json:"reason"`

	ExitCode *int32 `json:"exit_code"`

	ErrorMessage *string `json:"error_message"`

	CPULimit uint32 `json:"cpu_limit"`

	MemoryLimit uint32 `json:"memory_limit"`

	LocalStorageCapacity uint32 `json:"local_storage_capacity"`

	// Deprecated
	Command *string `json:"command,omitempty"`

	EnvironmentVariables map[string]string `json:"environment_variables"`

	StartupCommand []string `json:"startup_command"`

	Args []string `json:"args"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
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

	// CPULimit: CPU limit of the job (in mvCPU).
	CPULimit uint32 `json:"cpu_limit"`

	// MemoryLimit: memory limit of the job (in MiB).
	MemoryLimit uint32 `json:"memory_limit"`

	// LocalStorageCapacity: local storage capacity of the job (in MiB).
	LocalStorageCapacity uint32 `json:"local_storage_capacity"`

	// ImageURI: image to use for the job.
	ImageURI string `json:"image_uri"`

	// Deprecated: Command: deprecated: please use startup_command instead.
	Command *string `json:"command,omitempty"`

	// StartupCommand: the main executable or entrypoint script to run.
	// If both command and startup_command are provided, only startup_command will be used.
	StartupCommand []string `json:"startup_command"`

	// Args: passed to the startup command at runtime.
	// Environment variables and secrets can be included, and will be expanded before the arguments are used.
	Args []string `json:"args"`

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

// CreateSecretsRequest: create secrets request.
type CreateSecretsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobDefinitionID: UUID of the job definition.
	JobDefinitionID string `json:"job_definition_id"`

	// Secrets: list of secrets to inject into the job.
	Secrets []*CreateSecretsRequestSecretConfig `json:"secrets"`
}

// CreateSecretsResponse: create secrets response.
type CreateSecretsResponse struct {
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

// DeleteSecretRequest: delete secret request.
type DeleteSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

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

// GetJobLimitsRequest: get job limits request.
type GetJobLimitsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// GetJobRunRequest: get job run request.
type GetJobRunRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobRunID: UUID of the job run to get.
	JobRunID string `json:"-"`
}

// GetSecretRequest: get secret request.
type GetSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: UUID of the secret reference within the job.
	SecretID string `json:"-"`
}

// JobLimits: job limits.
type JobLimits struct {
	SecretsPerJobDefinition uint32 `json:"secrets_per_job_definition"`
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

// ListJobResourcesRequest: list job resources request.
type ListJobResourcesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// ListJobResourcesResponse: list job resources response.
type ListJobResourcesResponse struct {
	Resources []*Resource `json:"resources"`
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

	Reasons []JobRunReason `json:"-"`
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

// ListSecretsRequest: list secrets request.
type ListSecretsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobDefinitionID: UUID of the job definition.
	JobDefinitionID string `json:"job_definition_id"`
}

// ListSecretsResponse: list secrets response.
type ListSecretsResponse struct {
	// Secrets: list of secret references within a job definition.
	Secrets []*Secret `json:"secrets"`

	// TotalCount: total count of secret references within a job definition.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecretsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecretsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListSecretsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Secrets = append(r.Secrets, results.Secrets...)
	r.TotalCount += uint64(len(results.Secrets))
	return uint64(len(results.Secrets)), nil
}

// StartJobDefinitionRequest: start job definition request.
type StartJobDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// JobDefinitionID: UUID of the job definition to start.
	JobDefinitionID string `json:"-"`

	// Deprecated: Command: if empty or not defined, the image's default command is used.
	// Deprecated: please use startup_command instead.
	Command *string `json:"command,omitempty"`

	// StartupCommand: overrides the default defined in the job image.
	// The main executable or entrypoint script to run.
	// If both command and startup_command are provided, only startup_command will be used.
	StartupCommand *[]string `json:"startup_command,omitempty"`

	// Args: overrides the default arguments defined in the job image.
	// Passed to the contextual startup command at runtime.
	// Environment variables and secrets can be included, and will be expanded before the arguments are used.
	Args *[]string `json:"args,omitempty"`

	// EnvironmentVariables: contextual environment variables for this specific job run.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// Replicas: number of jobs to run.
	Replicas *uint32 `json:"replicas,omitempty"`
}

// StartJobDefinitionResponse: start job definition response.
type StartJobDefinitionResponse struct {
	// JobRuns: list of started job runs.
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

	// CPULimit: CPU limit of the job (in mvCPU).
	CPULimit *uint32 `json:"cpu_limit,omitempty"`

	// MemoryLimit: memory limit of the job (in MiB).
	MemoryLimit *uint32 `json:"memory_limit,omitempty"`

	// LocalStorageCapacity: local storage capacity of the job (in MiB).
	LocalStorageCapacity *uint32 `json:"local_storage_capacity,omitempty"`

	// ImageURI: image to use for the job.
	ImageURI *string `json:"image_uri,omitempty"`

	// Deprecated: Command: deprecated: please use startup_command instead.
	Command *string `json:"command,omitempty"`

	// StartupCommand: the main executable or entrypoint script to run.
	// If both command and startup_command are provided, only startup_command will be used.
	StartupCommand *[]string `json:"startup_command,omitempty"`

	// Args: passed to the startup command at runtime.
	// Environment variables and secrets can be included, and will be expanded before the arguments are used.
	Args *[]string `json:"args,omitempty"`

	// EnvironmentVariables: environment variables of the job.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// Description: description of the job.
	Description *string `json:"description,omitempty"`

	// JobTimeout: timeout of the job in seconds.
	JobTimeout *scw.Duration `json:"job_timeout,omitempty"`

	CronSchedule *UpdateJobDefinitionRequestCronScheduleConfig `json:"cron_schedule,omitempty"`
}

// UpdateSecretRequest: update secret request.
type UpdateSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

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
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/job-definitions",
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
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "",
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
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/job-definitions",
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
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "",
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

// DeleteJobDefinition: Delete an existing job definition by its unique identifier.
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
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// StartJobDefinition: Run an existing job definition using its unique identifier and create a new job run.
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
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/job-definitions/" + fmt.Sprint(req.JobDefinitionID) + "/start",
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
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/job-runs/" + fmt.Sprint(req.JobRunID) + "",
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
	parameter.AddToQuery(query, "reasons", req.Reasons)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/job-runs",
		Query:  query,
	}

	var resp ListJobRunsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopJobRun: Stop a job run using its unique identifier.
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
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/job-runs/" + fmt.Sprint(req.JobRunID) + "/stop",
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

// CreateSecrets: Create a secret reference within a job definition.
func (s *API) CreateSecrets(req *CreateSecretsRequest, opts ...scw.RequestOption) (*CreateSecretsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/secrets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateSecretsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSecret: Get a secret references within a job definition.
func (s *API) GetSecret(req *GetSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSecrets: List secrets references within a job definition.
func (s *API) ListSecrets(req *ListSecretsRequest, opts ...scw.RequestOption) (*ListSecretsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "job_definition_id", req.JobDefinitionID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/secrets",
		Query:  query,
	}

	var resp ListSecretsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSecret: Update a secret reference within a job definition.
func (s *API) UpdateSecret(req *UpdateSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
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

// DeleteSecret: Delete a secret reference within a job definition.
func (s *API) DeleteSecret(req *DeleteSecretRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListJobResources: List job resources for the console.
func (s *API) ListJobResources(req *ListJobResourcesRequest, opts ...scw.RequestOption) (*ListJobResourcesResponse, error) {
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
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/job-resources",
	}

	var resp ListJobResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetJobLimits: Get job limits for the console.
func (s *API) GetJobLimits(req *GetJobLimitsRequest, opts ...scw.RequestOption) (*JobLimits, error) {
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
		Path:   "/serverless-jobs/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/job-limits",
	}

	var resp JobLimits

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

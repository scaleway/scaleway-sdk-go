// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package dataorchestrator provides methods and message types of the dataorchestrator v1alpha2 API.
package dataorchestrator

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
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultDataorchestratorRetryInterval = 15 * time.Second
	defaultDataorchestratorTimeout       = 5 * time.Minute
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

type DefaultTaskDefaultRunReason string

const (
	DefaultTaskDefaultRunReasonUnknownDefaultRunReason = DefaultTaskDefaultRunReason("unknown_default_run_reason")
	DefaultTaskDefaultRunReasonTimeout                 = DefaultTaskDefaultRunReason("timeout")
	DefaultTaskDefaultRunReasonCancellation            = DefaultTaskDefaultRunReason("cancellation")
	DefaultTaskDefaultRunReasonTechnicalError          = DefaultTaskDefaultRunReason("technical_error")
	DefaultTaskDefaultRunReasonFunctionalError         = DefaultTaskDefaultRunReason("functional_error")
)

func (enum DefaultTaskDefaultRunReason) String() string {
	if enum == "" {
		// return default value if empty
		return string(DefaultTaskDefaultRunReasonUnknownDefaultRunReason)
	}
	return string(enum)
}

func (enum DefaultTaskDefaultRunReason) Values() []DefaultTaskDefaultRunReason {
	return []DefaultTaskDefaultRunReason{
		"unknown_default_run_reason",
		"timeout",
		"cancellation",
		"technical_error",
		"functional_error",
	}
}

func (enum DefaultTaskDefaultRunReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DefaultTaskDefaultRunReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DefaultTaskDefaultRunReason(DefaultTaskDefaultRunReason(tmp).String())
	return nil
}

type ListWorkflowDefinitionsRequestOrderBy string

const (
	ListWorkflowDefinitionsRequestOrderByCreatedAtAsc  = ListWorkflowDefinitionsRequestOrderBy("created_at_asc")
	ListWorkflowDefinitionsRequestOrderByCreatedAtDesc = ListWorkflowDefinitionsRequestOrderBy("created_at_desc")
	ListWorkflowDefinitionsRequestOrderByNameAsc       = ListWorkflowDefinitionsRequestOrderBy("name_asc")
	ListWorkflowDefinitionsRequestOrderByNameDesc      = ListWorkflowDefinitionsRequestOrderBy("name_desc")
	ListWorkflowDefinitionsRequestOrderByUpdatedAtAsc  = ListWorkflowDefinitionsRequestOrderBy("updated_at_asc")
	ListWorkflowDefinitionsRequestOrderByUpdatedAtDesc = ListWorkflowDefinitionsRequestOrderBy("updated_at_desc")
)

func (enum ListWorkflowDefinitionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListWorkflowDefinitionsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListWorkflowDefinitionsRequestOrderBy) Values() []ListWorkflowDefinitionsRequestOrderBy {
	return []ListWorkflowDefinitionsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
		"updated_at_asc",
		"updated_at_desc",
	}
}

func (enum ListWorkflowDefinitionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListWorkflowDefinitionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListWorkflowDefinitionsRequestOrderBy(ListWorkflowDefinitionsRequestOrderBy(tmp).String())
	return nil
}

type ListWorkflowRunsRequestOrderBy string

const (
	ListWorkflowRunsRequestOrderByCreatedAtAsc               = ListWorkflowRunsRequestOrderBy("created_at_asc")
	ListWorkflowRunsRequestOrderByCreatedAtDesc              = ListWorkflowRunsRequestOrderBy("created_at_desc")
	ListWorkflowRunsRequestOrderByTerminatedAtAsc            = ListWorkflowRunsRequestOrderBy("terminated_at_asc")
	ListWorkflowRunsRequestOrderByTerminatedAtDesc           = ListWorkflowRunsRequestOrderBy("terminated_at_desc")
	ListWorkflowRunsRequestOrderByDurationAsc                = ListWorkflowRunsRequestOrderBy("duration_asc")
	ListWorkflowRunsRequestOrderByDurationDesc               = ListWorkflowRunsRequestOrderBy("duration_desc")
	ListWorkflowRunsRequestOrderByWorkflowDefinitionNameAsc  = ListWorkflowRunsRequestOrderBy("workflow_definition_name_asc")
	ListWorkflowRunsRequestOrderByWorkflowDefinitionNameDesc = ListWorkflowRunsRequestOrderBy("workflow_definition_name_desc")
)

func (enum ListWorkflowRunsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListWorkflowRunsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListWorkflowRunsRequestOrderBy) Values() []ListWorkflowRunsRequestOrderBy {
	return []ListWorkflowRunsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"terminated_at_asc",
		"terminated_at_desc",
		"duration_asc",
		"duration_desc",
		"workflow_definition_name_asc",
		"workflow_definition_name_desc",
	}
}

func (enum ListWorkflowRunsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListWorkflowRunsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListWorkflowRunsRequestOrderBy(ListWorkflowRunsRequestOrderBy(tmp).String())
	return nil
}

type RunServerlessContainersTaskAction string

const (
	RunServerlessContainersTaskActionUnknownAction = RunServerlessContainersTaskAction("unknown_action")
	RunServerlessContainersTaskActionDeploy        = RunServerlessContainersTaskAction("deploy")
	RunServerlessContainersTaskActionDelete        = RunServerlessContainersTaskAction("delete")
	RunServerlessContainersTaskActionGet           = RunServerlessContainersTaskAction("get")
)

func (enum RunServerlessContainersTaskAction) String() string {
	if enum == "" {
		// return default value if empty
		return string(RunServerlessContainersTaskActionUnknownAction)
	}
	return string(enum)
}

func (enum RunServerlessContainersTaskAction) Values() []RunServerlessContainersTaskAction {
	return []RunServerlessContainersTaskAction{
		"unknown_action",
		"deploy",
		"delete",
		"get",
	}
}

func (enum RunServerlessContainersTaskAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RunServerlessContainersTaskAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RunServerlessContainersTaskAction(RunServerlessContainersTaskAction(tmp).String())
	return nil
}

type RunServerlessContainersTaskReason string

const (
	RunServerlessContainersTaskReasonUnknownReason   = RunServerlessContainersTaskReason("unknown_reason")
	RunServerlessContainersTaskReasonTimeout         = RunServerlessContainersTaskReason("timeout")
	RunServerlessContainersTaskReasonCancellation    = RunServerlessContainersTaskReason("cancellation")
	RunServerlessContainersTaskReasonTechnicalError  = RunServerlessContainersTaskReason("technical_error")
	RunServerlessContainersTaskReasonFunctionalError = RunServerlessContainersTaskReason("functional_error")
)

func (enum RunServerlessContainersTaskReason) String() string {
	if enum == "" {
		// return default value if empty
		return string(RunServerlessContainersTaskReasonUnknownReason)
	}
	return string(enum)
}

func (enum RunServerlessContainersTaskReason) Values() []RunServerlessContainersTaskReason {
	return []RunServerlessContainersTaskReason{
		"unknown_reason",
		"timeout",
		"cancellation",
		"technical_error",
		"functional_error",
	}
}

func (enum RunServerlessContainersTaskReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RunServerlessContainersTaskReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RunServerlessContainersTaskReason(RunServerlessContainersTaskReason(tmp).String())
	return nil
}

type RunServerlessContainersTaskStatus string

const (
	RunServerlessContainersTaskStatusUnknownStatus = RunServerlessContainersTaskStatus("unknown_status")
	RunServerlessContainersTaskStatusReady         = RunServerlessContainersTaskStatus("ready")
	RunServerlessContainersTaskStatusDeleting      = RunServerlessContainersTaskStatus("deleting")
	RunServerlessContainersTaskStatusError         = RunServerlessContainersTaskStatus("error")
	RunServerlessContainersTaskStatusUpdating      = RunServerlessContainersTaskStatus("updating")
	RunServerlessContainersTaskStatusCreating      = RunServerlessContainersTaskStatus("creating")
	RunServerlessContainersTaskStatusLocked        = RunServerlessContainersTaskStatus("locked")
	RunServerlessContainersTaskStatusUpgrading     = RunServerlessContainersTaskStatus("upgrading")
	RunServerlessContainersTaskStatusRebalancing   = RunServerlessContainersTaskStatus("rebalancing")
	RunServerlessContainersTaskStatusTriggering    = RunServerlessContainersTaskStatus("triggering")
)

func (enum RunServerlessContainersTaskStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(RunServerlessContainersTaskStatusUnknownStatus)
	}
	return string(enum)
}

func (enum RunServerlessContainersTaskStatus) Values() []RunServerlessContainersTaskStatus {
	return []RunServerlessContainersTaskStatus{
		"unknown_status",
		"ready",
		"deleting",
		"error",
		"updating",
		"creating",
		"locked",
		"upgrading",
		"rebalancing",
		"triggering",
	}
}

func (enum RunServerlessContainersTaskStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RunServerlessContainersTaskStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RunServerlessContainersTaskStatus(RunServerlessContainersTaskStatus(tmp).String())
	return nil
}

type RunStatus string

const (
	RunStatusUnknownRunStatus = RunStatus("unknown_run_status")
	RunStatusInitialized      = RunStatus("initialized")
	RunStatusValidated        = RunStatus("validated")
	RunStatusQueued           = RunStatus("queued")
	RunStatusRunning          = RunStatus("running")
	RunStatusSucceeded        = RunStatus("succeeded")
	RunStatusFailed           = RunStatus("failed")
	RunStatusInterrupting     = RunStatus("interrupting")
	RunStatusInterrupted      = RunStatus("interrupted")
	RunStatusPausing          = RunStatus("pausing")
	RunStatusPaused           = RunStatus("paused")
)

func (enum RunStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(RunStatusUnknownRunStatus)
	}
	return string(enum)
}

func (enum RunStatus) Values() []RunStatus {
	return []RunStatus{
		"unknown_run_status",
		"initialized",
		"validated",
		"queued",
		"running",
		"succeeded",
		"failed",
		"interrupting",
		"interrupted",
		"pausing",
		"paused",
	}
}

func (enum RunStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RunStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RunStatus(RunStatus(tmp).String())
	return nil
}

type S3TaskReason string

const (
	S3TaskReasonUnknownDefaultRunReason = S3TaskReason("unknown_default_run_reason")
	S3TaskReasonTimeout                 = S3TaskReason("timeout")
	S3TaskReasonCancellation            = S3TaskReason("cancellation")
	S3TaskReasonTechnicalError          = S3TaskReason("technical_error")
	S3TaskReasonFunctionalError         = S3TaskReason("functional_error")
)

func (enum S3TaskReason) String() string {
	if enum == "" {
		// return default value if empty
		return string(S3TaskReasonUnknownDefaultRunReason)
	}
	return string(enum)
}

func (enum S3TaskReason) Values() []S3TaskReason {
	return []S3TaskReason{
		"unknown_default_run_reason",
		"timeout",
		"cancellation",
		"technical_error",
		"functional_error",
	}
}

func (enum S3TaskReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *S3TaskReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = S3TaskReason(S3TaskReason(tmp).String())
	return nil
}

type ScheduleOverlapPolicy string

const (
	ScheduleOverlapPolicyUnknownScheduleOverlapPolicy = ScheduleOverlapPolicy("unknown_schedule_overlap_policy")
	ScheduleOverlapPolicyForbid                       = ScheduleOverlapPolicy("forbid")
	ScheduleOverlapPolicyBufferOne                    = ScheduleOverlapPolicy("buffer_one")
	ScheduleOverlapPolicyBufferAll                    = ScheduleOverlapPolicy("buffer_all")
	ScheduleOverlapPolicyAllowAll                     = ScheduleOverlapPolicy("allow_all")
)

func (enum ScheduleOverlapPolicy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ScheduleOverlapPolicyUnknownScheduleOverlapPolicy)
	}
	return string(enum)
}

func (enum ScheduleOverlapPolicy) Values() []ScheduleOverlapPolicy {
	return []ScheduleOverlapPolicy{
		"unknown_schedule_overlap_policy",
		"forbid",
		"buffer_one",
		"buffer_all",
		"allow_all",
	}
}

func (enum ScheduleOverlapPolicy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ScheduleOverlapPolicy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ScheduleOverlapPolicy(ScheduleOverlapPolicy(tmp).String())
	return nil
}

type ServerlessJobsTaskServerlessJobRunReason string

const (
	ServerlessJobsTaskServerlessJobRunReasonUnknownServerlessJobRunReason = ServerlessJobsTaskServerlessJobRunReason("unknown_serverless_job_run_reason")
	ServerlessJobsTaskServerlessJobRunReasonInvalidRequest                = ServerlessJobsTaskServerlessJobRunReason("invalid_request")
	ServerlessJobsTaskServerlessJobRunReasonTimeout                       = ServerlessJobsTaskServerlessJobRunReason("timeout")
	ServerlessJobsTaskServerlessJobRunReasonCancellation                  = ServerlessJobsTaskServerlessJobRunReason("cancellation")
	ServerlessJobsTaskServerlessJobRunReasonTechnicalError                = ServerlessJobsTaskServerlessJobRunReason("technical_error")
	ServerlessJobsTaskServerlessJobRunReasonImageNotFound                 = ServerlessJobsTaskServerlessJobRunReason("image_not_found")
	ServerlessJobsTaskServerlessJobRunReasonInvalidImage                  = ServerlessJobsTaskServerlessJobRunReason("invalid_image")
	ServerlessJobsTaskServerlessJobRunReasonMemoryUsageExceeded           = ServerlessJobsTaskServerlessJobRunReason("memory_usage_exceeded")
	ServerlessJobsTaskServerlessJobRunReasonStorageUsageExceeded          = ServerlessJobsTaskServerlessJobRunReason("storage_usage_exceeded")
	ServerlessJobsTaskServerlessJobRunReasonExitedWithError               = ServerlessJobsTaskServerlessJobRunReason("exited_with_error")
	ServerlessJobsTaskServerlessJobRunReasonSecretDisabled                = ServerlessJobsTaskServerlessJobRunReason("secret_disabled")
	ServerlessJobsTaskServerlessJobRunReasonSecretNotFound                = ServerlessJobsTaskServerlessJobRunReason("secret_not_found")
	ServerlessJobsTaskServerlessJobRunReasonQuotaExceeded                 = ServerlessJobsTaskServerlessJobRunReason("quota_exceeded")
)

func (enum ServerlessJobsTaskServerlessJobRunReason) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerlessJobsTaskServerlessJobRunReasonUnknownServerlessJobRunReason)
	}
	return string(enum)
}

func (enum ServerlessJobsTaskServerlessJobRunReason) Values() []ServerlessJobsTaskServerlessJobRunReason {
	return []ServerlessJobsTaskServerlessJobRunReason{
		"unknown_serverless_job_run_reason",
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

func (enum ServerlessJobsTaskServerlessJobRunReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerlessJobsTaskServerlessJobRunReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerlessJobsTaskServerlessJobRunReason(ServerlessJobsTaskServerlessJobRunReason(tmp).String())
	return nil
}

type TriggerType string

const (
	TriggerTypeUnknownTriggerType = TriggerType("unknown_trigger_type")
	TriggerTypeManual             = TriggerType("manual")
	TriggerTypeScheduled          = TriggerType("scheduled")
)

func (enum TriggerType) String() string {
	if enum == "" {
		// return default value if empty
		return string(TriggerTypeUnknownTriggerType)
	}
	return string(enum)
}

func (enum TriggerType) Values() []TriggerType {
	return []TriggerType{
		"unknown_trigger_type",
		"manual",
		"scheduled",
	}
}

func (enum TriggerType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TriggerType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TriggerType(TriggerType(tmp).String())
	return nil
}

type WorkflowRunReason string

const (
	WorkflowRunReasonUnknownWorkflowRunReason = WorkflowRunReason("unknown_workflow_run_reason")
	WorkflowRunReasonTimeout                  = WorkflowRunReason("timeout")
	WorkflowRunReasonCancellation             = WorkflowRunReason("cancellation")
	WorkflowRunReasonTechnicalError           = WorkflowRunReason("technical_error")
	WorkflowRunReasonFunctionalError          = WorkflowRunReason("functional_error")
	WorkflowRunReasonNotFound                 = WorkflowRunReason("not_found")
)

func (enum WorkflowRunReason) String() string {
	if enum == "" {
		// return default value if empty
		return string(WorkflowRunReasonUnknownWorkflowRunReason)
	}
	return string(enum)
}

func (enum WorkflowRunReason) Values() []WorkflowRunReason {
	return []WorkflowRunReason{
		"unknown_workflow_run_reason",
		"timeout",
		"cancellation",
		"technical_error",
		"functional_error",
		"not_found",
	}
}

func (enum WorkflowRunReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *WorkflowRunReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = WorkflowRunReason(WorkflowRunReason(tmp).String())
	return nil
}

// DefaultTask: default task.
type DefaultTask struct {
	// Reason: current reason of the task run.
	// Default value: unknown_default_run_reason
	Reason *DefaultTaskDefaultRunReason `json:"reason"`

	// ErrorMessage: error message of the task run.
	ErrorMessage *string `json:"error_message"`
}

// RunServerlessContainersTask: run serverless containers task.
type RunServerlessContainersTask struct {
	// ContainerID: unique identifier of the serverless-container deployment.
	ContainerID string `json:"container_id"`

	// ImageURL: image deployed with the serverless-container.
	ImageURL string `json:"image_url"`

	// Status: status of the serverless-container deployment.
	// Default value: unknown_status
	Status RunServerlessContainersTaskStatus `json:"status"`

	// Action: action used on the serverless-container deployment.
	// Default value: unknown_action
	Action RunServerlessContainersTaskAction `json:"action"`

	// Reason: current reason of the task run.
	// Default value: unknown_reason
	Reason *RunServerlessContainersTaskReason `json:"reason"`

	// ErrorMessage: error message of the task run.
	ErrorMessage *string `json:"error_message"`
}

// S3Task: s3 task.
type S3Task struct {
	Action string `json:"action"`

	// Reason: default value: unknown_default_run_reason
	Reason *S3TaskReason `json:"reason"`

	ErrorMessage *string `json:"error_message"`
}

// ServerlessJobsTask: serverless jobs task.
type ServerlessJobsTask struct {
	// JobDefinitionID: unique identifier of the serverless-job definition.
	JobDefinitionID string `json:"job_definition_id"`

	// JobRunID: unique identifier of the serverless-job run.
	JobRunID string `json:"job_run_id"`

	// Reason: terminated reason of the serverless-job run.
	// Default value: unknown_serverless_job_run_reason
	Reason *ServerlessJobsTaskServerlessJobRunReason `json:"reason"`

	// ExitCode: exit code of the serverless-job run.
	ExitCode *int32 `json:"exit_code"`

	// ErrorMessage: error message of the serverless-job run.
	ErrorMessage *string `json:"error_message"`
}

// CronSchedule: cron schedule.
type CronSchedule struct {
	// Schedule: uNIX cron schedule to run workflow (e.g., '* * * * *').
	Schedule string `json:"schedule"`

	// TimezoneIDentifier: timezone identifier for the cron schedule, in tz database format (e.g., 'Europe/Paris').
	TimezoneIDentifier string `json:"timezone_identifier"`

	// Overlap: controls what happens when an WorkflowDefinition would be started by a Schedule.
	// Default value: unknown_schedule_overlap_policy
	Overlap ScheduleOverlapPolicy `json:"overlap"`
}

// WorkflowRunInfo: workflow run info.
type WorkflowRunInfo struct {
	// WorkflowRunID: workflow run ID.
	WorkflowRunID string `json:"workflow_run_id"`

	// Status: workflow run status.
	// Default value: unknown_run_status
	Status RunStatus `json:"status"`

	// Reason: workflow run reason.
	// Default value: unknown_workflow_run_reason
	Reason *WorkflowRunReason `json:"reason"`
}

// WorkflowVersion: workflow version.
type WorkflowVersion struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// Name: name of the workflow version.
	Name string `json:"name"`

	// Description: description of the workflow version.
	Description string `json:"description"`

	// CreatedAt: creation date of the workflow version.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the workflow version.
	UpdatedAt *time.Time `json:"updated_at"`

	// DeletedAt: deleted date of the workflow version.
	DeletedAt *time.Time `json:"deleted_at"`

	// WorkflowDsl: dSL of workflow in YAML format based from https://open-workflow-specification.org/schemas/1.0.3/workflow.yaml.
	WorkflowDsl string `json:"workflow_dsl"`

	// IsCurrent: indicate if the version is the current selected by the workflow to run.
	IsCurrent bool `json:"is_current"`
}

// TaskRun: task run.
type TaskRun struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// TaskName: unique identifier of the task.
	TaskName string `json:"task_name"`

	// CreatedAt: creation date of the task run.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the task run.
	UpdatedAt *time.Time `json:"updated_at"`

	// TerminatedAt: terminated date of the task run.
	TerminatedAt *time.Time `json:"terminated_at"`

	// Duration: run duration of the task run.
	Duration *scw.Duration `json:"duration"`

	// Status: status of the task run.
	// Default value: unknown_run_status
	Status RunStatus `json:"status"`

	// Output: output of the task run.
	Output *string `json:"output"`

	// ServerlessJobTaskRunDetails: task run details for a serverless-job task.
	// Precisely one of ServerlessJobTaskRunDetails, RaiseTaskRunDetails, SetTaskRunDetails, CallHTTPTaskRunDetails, WaitTaskRunDetails, ServerlessFunctionTaskRunDetails, RunServerlessContainerTaskRunDetails, CallServerlessContainerTaskRunDetails, S3TaskRunDetails, GenerativeAPITaskRunDetails must be set.
	ServerlessJobTaskRunDetails *ServerlessJobsTask `json:"serverless_job_task_run_details,omitempty"`

	// RaiseTaskRunDetails: task run details for a raise task.
	// Precisely one of ServerlessJobTaskRunDetails, RaiseTaskRunDetails, SetTaskRunDetails, CallHTTPTaskRunDetails, WaitTaskRunDetails, ServerlessFunctionTaskRunDetails, RunServerlessContainerTaskRunDetails, CallServerlessContainerTaskRunDetails, S3TaskRunDetails, GenerativeAPITaskRunDetails must be set.
	RaiseTaskRunDetails *DefaultTask `json:"raise_task_run_details,omitempty"`

	// SetTaskRunDetails: task run details for a set task.
	// Precisely one of ServerlessJobTaskRunDetails, RaiseTaskRunDetails, SetTaskRunDetails, CallHTTPTaskRunDetails, WaitTaskRunDetails, ServerlessFunctionTaskRunDetails, RunServerlessContainerTaskRunDetails, CallServerlessContainerTaskRunDetails, S3TaskRunDetails, GenerativeAPITaskRunDetails must be set.
	SetTaskRunDetails *DefaultTask `json:"set_task_run_details,omitempty"`

	// CallHTTPTaskRunDetails: task run details for a call http task.
	// Precisely one of ServerlessJobTaskRunDetails, RaiseTaskRunDetails, SetTaskRunDetails, CallHTTPTaskRunDetails, WaitTaskRunDetails, ServerlessFunctionTaskRunDetails, RunServerlessContainerTaskRunDetails, CallServerlessContainerTaskRunDetails, S3TaskRunDetails, GenerativeAPITaskRunDetails must be set.
	CallHTTPTaskRunDetails *DefaultTask `json:"call_http_task_run_details,omitempty"`

	// WaitTaskRunDetails: task run details for a wait task.
	// Precisely one of ServerlessJobTaskRunDetails, RaiseTaskRunDetails, SetTaskRunDetails, CallHTTPTaskRunDetails, WaitTaskRunDetails, ServerlessFunctionTaskRunDetails, RunServerlessContainerTaskRunDetails, CallServerlessContainerTaskRunDetails, S3TaskRunDetails, GenerativeAPITaskRunDetails must be set.
	WaitTaskRunDetails *DefaultTask `json:"wait_task_run_details,omitempty"`

	// ServerlessFunctionTaskRunDetails: task run details for a serverless-function task.
	// Precisely one of ServerlessJobTaskRunDetails, RaiseTaskRunDetails, SetTaskRunDetails, CallHTTPTaskRunDetails, WaitTaskRunDetails, ServerlessFunctionTaskRunDetails, RunServerlessContainerTaskRunDetails, CallServerlessContainerTaskRunDetails, S3TaskRunDetails, GenerativeAPITaskRunDetails must be set.
	ServerlessFunctionTaskRunDetails *DefaultTask `json:"serverless_function_task_run_details,omitempty"`

	// RunServerlessContainerTaskRunDetails: task run details for a run-serverless-container task.
	// Precisely one of ServerlessJobTaskRunDetails, RaiseTaskRunDetails, SetTaskRunDetails, CallHTTPTaskRunDetails, WaitTaskRunDetails, ServerlessFunctionTaskRunDetails, RunServerlessContainerTaskRunDetails, CallServerlessContainerTaskRunDetails, S3TaskRunDetails, GenerativeAPITaskRunDetails must be set.
	RunServerlessContainerTaskRunDetails *RunServerlessContainersTask `json:"run_serverless_container_task_run_details,omitempty"`

	// CallServerlessContainerTaskRunDetails: task run details for a call-serverless-container task.
	// Precisely one of ServerlessJobTaskRunDetails, RaiseTaskRunDetails, SetTaskRunDetails, CallHTTPTaskRunDetails, WaitTaskRunDetails, ServerlessFunctionTaskRunDetails, RunServerlessContainerTaskRunDetails, CallServerlessContainerTaskRunDetails, S3TaskRunDetails, GenerativeAPITaskRunDetails must be set.
	CallServerlessContainerTaskRunDetails *DefaultTask `json:"call_serverless_container_task_run_details,omitempty"`

	// S3TaskRunDetails: task run details for a S3 task.
	// Precisely one of ServerlessJobTaskRunDetails, RaiseTaskRunDetails, SetTaskRunDetails, CallHTTPTaskRunDetails, WaitTaskRunDetails, ServerlessFunctionTaskRunDetails, RunServerlessContainerTaskRunDetails, CallServerlessContainerTaskRunDetails, S3TaskRunDetails, GenerativeAPITaskRunDetails must be set.
	S3TaskRunDetails *S3Task `json:"s3_task_run_details,omitempty"`

	// GenerativeAPITaskRunDetails: task run details for a generative-api task.
	// Precisely one of ServerlessJobTaskRunDetails, RaiseTaskRunDetails, SetTaskRunDetails, CallHTTPTaskRunDetails, WaitTaskRunDetails, ServerlessFunctionTaskRunDetails, RunServerlessContainerTaskRunDetails, CallServerlessContainerTaskRunDetails, S3TaskRunDetails, GenerativeAPITaskRunDetails must be set.
	GenerativeAPITaskRunDetails *DefaultTask `json:"generative_api_task_run_details,omitempty"`

	// RunIteration: run iteration of the task (in loop or retry).
	RunIteration *int32 `json:"run_iteration"`
}

// CreateWorkflowDefinitionRequestCronScheduleConfig: create workflow definition request cron schedule config.
type CreateWorkflowDefinitionRequestCronScheduleConfig struct {
	Schedule string `json:"schedule"`

	TimezoneIDentifier string `json:"timezone_identifier"`

	// Overlap: default value: unknown_schedule_overlap_policy
	Overlap *ScheduleOverlapPolicy `json:"overlap"`
}

// CreateWorkflowDefinitionRequestWorkflowVersionConfig: create workflow definition request workflow version config.
type CreateWorkflowDefinitionRequestWorkflowVersionConfig struct {
	Name string `json:"name"`

	Description string `json:"description"`

	WorkflowDsl *string `json:"workflow_dsl"`
}

// WorkflowDefinition: workflow definition.
type WorkflowDefinition struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// Name: name of the workflow definition.
	Name string `json:"name"`

	// Description: description of the workflow definition.
	Description string `json:"description"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// CreatedAt: creation date of the workflow definition.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the workflow definition.
	UpdatedAt *time.Time `json:"updated_at"`

	// Versions: versions of workflow.
	Versions []*WorkflowVersion `json:"versions"`

	// CronSchedule: configure a cron for the workflow.
	CronSchedule *CronSchedule `json:"cron_schedule"`

	// EnvironmentVariables: environment variables of the workflow.
	EnvironmentVariables map[string]string `json:"environment_variables"`

	// WorkflowTimeout: timeout of the workflow in seconds.
	WorkflowTimeout *scw.Duration `json:"workflow_timeout"`

	// LastRun: last workflow run information.
	LastRun *WorkflowRunInfo `json:"last_run"`

	// Region: region of the workflow.
	Region scw.Region `json:"region"`

	// Srn: the SRN of the workflow definition.
	Srn string `json:"srn"`
}

//nolint:unused,nolintlint // generated, no check on whether it's used or not
func (m *WorkflowDefinition) setSRN(platform string) {
	if m.Srn != "" {
		// if the field is set server-side, trust the server
		return
	}

	// We do not check that *m.XYZ != "", as there are currently no use cases for an
	// optional value in an SRN where the value set to the empty string makes sense.

	if fmt.Sprint(m.Region) != "" && fmt.Sprint(m.ID) != "" {
		m.Srn = fmt.Sprintf("srn://data-orchestrator.%s/regions/%s/workflow-definitions/%s", platform, fmt.Sprint(m.Region), fmt.Sprint(m.ID))
		return
	}
}

// WorkflowRun: workflow run.
type WorkflowRun struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// WorkflowDefinitionID: unique identifier of the workflow definition.
	WorkflowDefinitionID string `json:"workflow_definition_id"`

	// WorkflowDefinitionName: name of the workflow definition.
	WorkflowDefinitionName string `json:"workflow_definition_name"`

	// WorkflowVersionID: unique identifier of the workflow version.
	WorkflowVersionID string `json:"workflow_version_id"`

	// ProjectID: unique identifier of the project.
	ProjectID string `json:"project_id"`

	// TriggerType: type of trigger that started the workflow run.
	// Default value: unknown_trigger_type
	TriggerType TriggerType `json:"trigger_type"`

	// CreatedAt: creation date of the workflow run.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the workflow run.
	UpdatedAt *time.Time `json:"updated_at"`

	// TerminatedAt: terminated date of the workflow run.
	TerminatedAt *time.Time `json:"terminated_at"`

	// Duration: run duration of the workflow run.
	Duration *scw.Duration `json:"duration"`

	// Status: current status of the workflow run.
	// Default value: unknown_run_status
	Status RunStatus `json:"status"`

	// Reason: current reason of the workflow run.
	// Default value: unknown_workflow_run_reason
	Reason *WorkflowRunReason `json:"reason"`

	// TaskRuns: task runs of the workflow run.
	TaskRuns []*TaskRun `json:"task_runs"`

	// EnvironmentVariables: environment variables of the workflow run.
	EnvironmentVariables map[string]string `json:"environment_variables"`

	// WorkflowDsl: dSL of workflow in YAML format based from https://open-workflow-specification.org/schemas/1.0.3/workflow.yaml.
	WorkflowDsl string `json:"workflow_dsl"`

	// Output: output of the workflow run.
	Output *string `json:"output"`

	// ErrorMessage: error message of the workflow run.
	ErrorMessage *string `json:"error_message"`

	// Region: region of the workflow run.
	Region scw.Region `json:"region"`

	// Srn: the SRN of the workflow run.
	Srn string `json:"srn"`
}

//nolint:unused,nolintlint // generated, no check on whether it's used or not
func (m *WorkflowRun) setSRN(platform string) {
	if m.Srn != "" {
		// if the field is set server-side, trust the server
		return
	}

	// We do not check that *m.XYZ != "", as there are currently no use cases for an
	// optional value in an SRN where the value set to the empty string makes sense.

	if fmt.Sprint(m.Region) != "" && fmt.Sprint(m.ID) != "" {
		m.Srn = fmt.Sprintf("srn://data-orchestrator.%s/regions/%s/workflow-runs/%s", platform, fmt.Sprint(m.Region), fmt.Sprint(m.ID))
		return
	}
}

// UpdateWorkflowDefinitionRequestCronScheduleConfig: update workflow definition request cron schedule config.
type UpdateWorkflowDefinitionRequestCronScheduleConfig struct {
	Schedule *string `json:"schedule"`

	TimezoneIDentifier *string `json:"timezone_identifier"`

	// Overlap: default value: unknown_schedule_overlap_policy
	Overlap *ScheduleOverlapPolicy `json:"overlap"`
}

// CreateWorkflowDefinitionRequest: create workflow definition request.
type CreateWorkflowDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: name of the workflow definition.
	Name string `json:"name"`

	// Description: description of the workflow.
	Description string `json:"description"`

	// ProjectID: UUID of the Scaleway Project containing the workflow.
	ProjectID string `json:"project_id"`

	// EnvironmentVariables: environment variables of the workflow.
	EnvironmentVariables map[string]string `json:"environment_variables"`

	// WorkflowTimeout: timeout of the workflow in seconds.
	WorkflowTimeout *scw.Duration `json:"workflow_timeout,omitempty"`

	// CronSchedule: configure a cron for the workflow.
	CronSchedule *CreateWorkflowDefinitionRequestCronScheduleConfig `json:"cron_schedule,omitempty"`

	// WorkflowVersion: version of the workflow definition.
	WorkflowVersion *CreateWorkflowDefinitionRequestWorkflowVersionConfig `json:"workflow_version"`
}

// DeleteWorkflowDefinitionRequest: delete workflow definition request.
type DeleteWorkflowDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WorkflowDefinitionID: UUID of the workflow definition to delete.
	WorkflowDefinitionID string `json:"-"`
}

// DeleteWorkflowRunRequest: delete workflow run request.
type DeleteWorkflowRunRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WorkflowRunID: UUID of the workflow run to delete.
	WorkflowRunID string `json:"-"`
}

// GetWorkflowDefinitionRequest: get workflow definition request.
type GetWorkflowDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WorkflowDefinitionID: UUID of the workflow definition to get.
	WorkflowDefinitionID string `json:"-"`
}

// GetWorkflowRunRequest: get workflow run request.
type GetWorkflowRunRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WorkflowRunID: UUID of the workflow run to get.
	WorkflowRunID string `json:"-"`
}

// ListWorkflowDefinitionsRequest: list workflow definitions request.
type ListWorkflowDefinitionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListWorkflowDefinitionsRequestOrderBy `json:"-"`

	Name *string `json:"-"`

	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	Statuses []RunStatus `json:"-"`

	CreatedAfter *time.Time `json:"-"`

	CreatedBefore *time.Time `json:"-"`

	TriggerTypes []TriggerType `json:"-"`
}

// ListWorkflowDefinitionsResponse: list workflow definitions response.
type ListWorkflowDefinitionsResponse struct {
	WorkflowDefinitions []*WorkflowDefinition `json:"workflow_definitions"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListWorkflowDefinitionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListWorkflowDefinitionsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListWorkflowDefinitionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.WorkflowDefinitions = append(r.WorkflowDefinitions, results.WorkflowDefinitions...)
	r.TotalCount += uint64(len(results.WorkflowDefinitions))
	return uint64(len(results.WorkflowDefinitions)), nil
}

// ListWorkflowRunsRequest: list workflow runs request.
type ListWorkflowRunsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListWorkflowRunsRequestOrderBy `json:"-"`

	WorkflowDefinitionID *string `json:"-"`

	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	Statuses []RunStatus `json:"-"`

	CreatedAfter *time.Time `json:"-"`

	CreatedBefore *time.Time `json:"-"`

	TriggerTypes []TriggerType `json:"-"`

	WorkflowDefinitionName *string `json:"-"`
}

// ListWorkflowRunsResponse: list workflow runs response.
type ListWorkflowRunsResponse struct {
	// WorkflowRuns: list workflow runs.
	WorkflowRuns []*WorkflowRun `json:"workflow_runs"`

	// TotalCount: total count of workflow runs.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListWorkflowRunsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListWorkflowRunsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListWorkflowRunsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.WorkflowRuns = append(r.WorkflowRuns, results.WorkflowRuns...)
	r.TotalCount += uint64(len(results.WorkflowRuns))
	return uint64(len(results.WorkflowRuns)), nil
}

// PauseWorkflowRunRequest: pause workflow run request.
type PauseWorkflowRunRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WorkflowRunID: UUID of the workflow run to pause.
	WorkflowRunID string `json:"-"`
}

// ResumeWorkflowRunRequest: resume workflow run request.
type ResumeWorkflowRunRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WorkflowRunID: UUID of the workflow run to resume.
	WorkflowRunID string `json:"-"`
}

// StartWorkflowDefinitionRequest: start workflow definition request.
type StartWorkflowDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WorkflowDefinitionID: UUID of the workflow definition to start.
	WorkflowDefinitionID string `json:"-"`

	// EnvironmentVariables: contextual environment variables for this specific workflow run.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`
}

// TerminateWorkflowRunRequest: terminate workflow run request.
type TerminateWorkflowRunRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WorkflowRunID: UUID of the workflow run to terminate.
	WorkflowRunID string `json:"-"`
}

// UpdateWorkflowDefinitionRequest: update workflow definition request.
type UpdateWorkflowDefinitionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WorkflowDefinitionID: UUID of the workflow definition to update.
	WorkflowDefinitionID string `json:"-"`

	// Name: name of the workflow definition.
	Name *string `json:"name,omitempty"`

	// Description: description of the workflow.
	Description *string `json:"description,omitempty"`

	// EnvironmentVariables: environment variables of the workflow.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// WorkflowTimeout: timeout of the workflow in seconds.
	WorkflowTimeout *scw.Duration `json:"workflow_timeout,omitempty"`

	// CronSchedule: configure a cron for the workflow.
	CronSchedule *UpdateWorkflowDefinitionRequestCronScheduleConfig `json:"cron_schedule,omitempty"`
}

// UpdateWorkflowVersionRequest: update workflow version request.
type UpdateWorkflowVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WorkflowVersionID: UUID of the workflow version to update.
	WorkflowVersionID string `json:"-"`

	// Name: name of the workflow version.
	Name *string `json:"name,omitempty"`

	// Description: description of the workflow version.
	Description *string `json:"description,omitempty"`

	// WorkflowDsl: dSL of workflow in YAML format based from https://open-workflow-specification.org/schemas/1.0.3/workflow.yaml.
	WorkflowDsl *string `json:"workflow_dsl,omitempty"`
}

// This API is a boilerplate for products.
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
	return []scw.Region{scw.RegionFrPar}
}

// CreateWorkflowDefinition: Create a new workflow definition in a specified Project.
func (s *API) CreateWorkflowDefinition(req *CreateWorkflowDefinitionRequest, opts ...scw.RequestOption) (*WorkflowDefinition, error) {
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
		req.Name = namegenerator.GetRandomName("workflow")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-definitions",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp WorkflowDefinition

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// GetWorkflowDefinition: Get a workflow definition by its unique identifier.
func (s *API) GetWorkflowDefinition(req *GetWorkflowDefinitionRequest, opts ...scw.RequestOption) (*WorkflowDefinition, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WorkflowDefinitionID) == "" {
		return nil, errors.New("field WorkflowDefinitionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-definitions/" + fmt.Sprint(req.WorkflowDefinitionID) + "",
	}

	var resp WorkflowDefinition

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// ListWorkflowDefinitions: List all your workflow definitions with filters.
func (s *API) ListWorkflowDefinitions(req *ListWorkflowDefinitionsRequest, opts ...scw.RequestOption) (*ListWorkflowDefinitionsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "created_after", req.CreatedAfter)
	parameter.AddToQuery(query, "created_before", req.CreatedBefore)
	parameter.AddToQuery(query, "trigger_types", req.TriggerTypes)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-definitions",
		Query:  query,
	}

	var resp ListWorkflowDefinitionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		for _, el := range resp.WorkflowDefinitions {
			el.setSRN(apiMetadata.Domain)
		}
	}
	return &resp, nil
}

// UpdateWorkflowDefinition: Update an existing workflow definition by its unique identifier.
func (s *API) UpdateWorkflowDefinition(req *UpdateWorkflowDefinitionRequest, opts ...scw.RequestOption) (*WorkflowDefinition, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WorkflowDefinitionID) == "" {
		return nil, errors.New("field WorkflowDefinitionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-definitions/" + fmt.Sprint(req.WorkflowDefinitionID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp WorkflowDefinition

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// DeleteWorkflowDefinition: Delete an existing workflow definition by its unique identifier.
func (s *API) DeleteWorkflowDefinition(req *DeleteWorkflowDefinitionRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WorkflowDefinitionID) == "" {
		return errors.New("field WorkflowDefinitionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-definitions/" + fmt.Sprint(req.WorkflowDefinitionID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateWorkflowVersion: Update an existing workflow version by its unique identifier.
func (s *API) UpdateWorkflowVersion(req *UpdateWorkflowVersionRequest, opts ...scw.RequestOption) (*WorkflowVersion, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WorkflowVersionID) == "" {
		return nil, errors.New("field WorkflowVersionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-versions/" + fmt.Sprint(req.WorkflowVersionID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp WorkflowVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartWorkflowDefinition: Run an existing workflow definition using its unique identifier and create a new workflow run.
func (s *API) StartWorkflowDefinition(req *StartWorkflowDefinitionRequest, opts ...scw.RequestOption) (*WorkflowRun, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WorkflowDefinitionID) == "" {
		return nil, errors.New("field WorkflowDefinitionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-definitions/" + fmt.Sprint(req.WorkflowDefinitionID) + "/start",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp WorkflowRun

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// GetWorkflowRun: Get a workflow run by its unique identifier.
func (s *API) GetWorkflowRun(req *GetWorkflowRunRequest, opts ...scw.RequestOption) (*WorkflowRun, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WorkflowRunID) == "" {
		return nil, errors.New("field WorkflowRunID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-runs/" + fmt.Sprint(req.WorkflowRunID) + "",
	}

	var resp WorkflowRun

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// WaitForWorkflowRunRequest is used by WaitForWorkflowRun method.
type WaitForWorkflowRunRequest struct {
	Region        scw.Region
	WorkflowRunID string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForWorkflowRun waits for the WorkflowRun to reach a terminal state.
func (s *API) WaitForWorkflowRun(req *WaitForWorkflowRunRequest, opts ...scw.RequestOption) (*WorkflowRun, error) {
	timeout := defaultDataorchestratorTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultDataorchestratorRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[RunStatus]struct{}{
		RunStatusInitialized:  {},
		RunStatusValidated:    {},
		RunStatusQueued:       {},
		RunStatusRunning:      {},
		RunStatusInterrupting: {},
		RunStatusPausing:      {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetWorkflowRun(&GetWorkflowRunRequest{
				Region:        req.Region,
				WorkflowRunID: req.WorkflowRunID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTransient := transientStatuses[res.Status]

			return res, !isTransient, nil
		},
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
		Timeout:          timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for WorkflowRun failed")
	}

	return res.(*WorkflowRun), nil
}

// ListWorkflowRuns: List all your workflow runs with filters.
func (s *API) ListWorkflowRuns(req *ListWorkflowRunsRequest, opts ...scw.RequestOption) (*ListWorkflowRunsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "workflow_definition_id", req.WorkflowDefinitionID)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "created_after", req.CreatedAfter)
	parameter.AddToQuery(query, "created_before", req.CreatedBefore)
	parameter.AddToQuery(query, "trigger_types", req.TriggerTypes)
	parameter.AddToQuery(query, "workflow_definition_name", req.WorkflowDefinitionName)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-runs",
		Query:  query,
	}

	var resp ListWorkflowRunsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		for _, el := range resp.WorkflowRuns {
			el.setSRN(apiMetadata.Domain)
		}
	}
	return &resp, nil
}

// TerminateWorkflowRun: Terminate a workflow run using its unique identifier.
func (s *API) TerminateWorkflowRun(req *TerminateWorkflowRunRequest, opts ...scw.RequestOption) (*WorkflowRun, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WorkflowRunID) == "" {
		return nil, errors.New("field WorkflowRunID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-runs/" + fmt.Sprint(req.WorkflowRunID) + "/terminate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp WorkflowRun

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// PauseWorkflowRun: Pause a workflow run using its unique identifier.
func (s *API) PauseWorkflowRun(req *PauseWorkflowRunRequest, opts ...scw.RequestOption) (*WorkflowRun, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WorkflowRunID) == "" {
		return nil, errors.New("field WorkflowRunID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-runs/" + fmt.Sprint(req.WorkflowRunID) + "/pause",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp WorkflowRun

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// ResumeWorkflowRun: Resume a workflow run using its unique identifier.
func (s *API) ResumeWorkflowRun(req *ResumeWorkflowRunRequest, opts ...scw.RequestOption) (*WorkflowRun, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WorkflowRunID) == "" {
		return nil, errors.New("field WorkflowRunID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-runs/" + fmt.Sprint(req.WorkflowRunID) + "/resume",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp WorkflowRun

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// DeleteWorkflowRun: Delete an existing workflow run by its unique identifier.
func (s *API) DeleteWorkflowRun(req *DeleteWorkflowRunRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WorkflowRunID) == "" {
		return errors.New("field WorkflowRunID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/data-orchestrator/v1alpha2/regions/" + fmt.Sprint(req.Region) + "/workflow-runs/" + fmt.Sprint(req.WorkflowRunID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

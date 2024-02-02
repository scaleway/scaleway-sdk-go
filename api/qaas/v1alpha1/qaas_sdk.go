// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package qaas provides methods and message types of the qaas v1alpha1 API.
package qaas

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

type JobStatus string

const (
	JobStatusUnknownStatus = JobStatus("unknown_status")
	JobStatusWaiting       = JobStatus("waiting")
	JobStatusError         = JobStatus("error")
	JobStatusRunning       = JobStatus("running")
	JobStatusCompleted     = JobStatus("completed")
	JobStatusCancelling    = JobStatus("cancelling")
	JobStatusCancelled     = JobStatus("cancelled")
)

func (enum JobStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum JobStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *JobStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = JobStatus(JobStatus(tmp).String())
	return nil
}

type ListJobsRequestOrderBy string

const (
	ListJobsRequestOrderByCreatedAtDesc    = ListJobsRequestOrderBy("created_at_desc")
	ListJobsRequestOrderByCreatedAtAsc     = ListJobsRequestOrderBy("created_at_asc")
	ListJobsRequestOrderByStatusAsc        = ListJobsRequestOrderBy("status_asc")
	ListJobsRequestOrderByStatusDesc       = ListJobsRequestOrderBy("status_desc")
	ListJobsRequestOrderByPlatformNameAsc  = ListJobsRequestOrderBy("platform_name_asc")
	ListJobsRequestOrderByPlatformNameDesc = ListJobsRequestOrderBy("platform_name_desc")
	ListJobsRequestOrderByNameAsc          = ListJobsRequestOrderBy("name_asc")
	ListJobsRequestOrderByNameDesc         = ListJobsRequestOrderBy("name_desc")
	ListJobsRequestOrderBySessionNameAsc   = ListJobsRequestOrderBy("session_name_asc")
	ListJobsRequestOrderBySessionNameDesc  = ListJobsRequestOrderBy("session_name_desc")
)

func (enum ListJobsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_desc"
	}
	return string(enum)
}

func (enum ListJobsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListJobsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListJobsRequestOrderBy(ListJobsRequestOrderBy(tmp).String())
	return nil
}

type ListPlatformsRequestOrderBy string

const (
	ListPlatformsRequestOrderByNameAsc          = ListPlatformsRequestOrderBy("name_asc")
	ListPlatformsRequestOrderByNameDesc         = ListPlatformsRequestOrderBy("name_desc")
	ListPlatformsRequestOrderByProviderNameAsc  = ListPlatformsRequestOrderBy("provider_name_asc")
	ListPlatformsRequestOrderByProviderNameDesc = ListPlatformsRequestOrderBy("provider_name_desc")
	ListPlatformsRequestOrderByTypeAsc          = ListPlatformsRequestOrderBy("type_asc")
	ListPlatformsRequestOrderByTypeDesc         = ListPlatformsRequestOrderBy("type_desc")
	ListPlatformsRequestOrderByTechnologyAsc    = ListPlatformsRequestOrderBy("technology_asc")
	ListPlatformsRequestOrderByTechnologyDesc   = ListPlatformsRequestOrderBy("technology_desc")
)

func (enum ListPlatformsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListPlatformsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPlatformsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPlatformsRequestOrderBy(ListPlatformsRequestOrderBy(tmp).String())
	return nil
}

type ListSessionsRequestOrderBy string

const (
	ListSessionsRequestOrderByNameAsc       = ListSessionsRequestOrderBy("name_asc")
	ListSessionsRequestOrderByNameDesc      = ListSessionsRequestOrderBy("name_desc")
	ListSessionsRequestOrderByStartedAtAsc  = ListSessionsRequestOrderBy("started_at_asc")
	ListSessionsRequestOrderByStartedAtDesc = ListSessionsRequestOrderBy("started_at_desc")
	ListSessionsRequestOrderByStatusAsc     = ListSessionsRequestOrderBy("status_asc")
	ListSessionsRequestOrderByStatusDesc    = ListSessionsRequestOrderBy("status_desc")
	ListSessionsRequestOrderByCreatedAtDesc = ListSessionsRequestOrderBy("created_at_desc")
	ListSessionsRequestOrderByCreatedAtAsc  = ListSessionsRequestOrderBy("created_at_asc")
)

func (enum ListSessionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListSessionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSessionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSessionsRequestOrderBy(ListSessionsRequestOrderBy(tmp).String())
	return nil
}

type PlatformTechnology string

const (
	PlatformTechnologyUnknownTechnology = PlatformTechnology("unknown_technology")
	PlatformTechnologyPhotonic          = PlatformTechnology("photonic")
)

func (enum PlatformTechnology) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_technology"
	}
	return string(enum)
}

func (enum PlatformTechnology) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlatformTechnology) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlatformTechnology(PlatformTechnology(tmp).String())
	return nil
}

type PlatformType string

const (
	PlatformTypeUnknownType = PlatformType("unknown_type")
	PlatformTypeSimulator   = PlatformType("simulator")
	PlatformTypeQpu         = PlatformType("qpu")
)

func (enum PlatformType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum PlatformType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlatformType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlatformType(PlatformType(tmp).String())
	return nil
}

type SessionStatus string

const (
	SessionStatusUnknownStatus = SessionStatus("unknown_status")
	SessionStatusRunning       = SessionStatus("running")
	SessionStatusStopped       = SessionStatus("stopped")
	SessionStatusStarting      = SessionStatus("starting")
	SessionStatusStopping      = SessionStatus("stopping")
)

func (enum SessionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum SessionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SessionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SessionStatus(SessionStatus(tmp).String())
	return nil
}

// JobCircuit: job circuit.
type JobCircuit struct {
	// PercevalCircuit: circuit generated by Perceval that should be executed.
	// Precisely one of PercevalCircuit must be set.
	PercevalCircuit *string `json:"perceval_circuit,omitempty"`
}

// Job: job.
type Job struct {
	// ID: unique ID of the job.
	ID string `json:"id"`

	// Name: job name.
	Name string `json:"name"`

	// Tags: tags of the job.
	Tags *[]string `json:"tags"`

	// SessionID: session ID in which the job is executed.
	SessionID string `json:"session_id"`

	// CreatedAt: time at which the job was created.
	CreatedAt *time.Time `json:"created_at"`

	// StartedAt: time at which the job was started.
	StartedAt *time.Time `json:"started_at"`

	// UpdatedAt: time at which the job was updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// Status: job status.
	// Default value: unknown_status
	Status JobStatus `json:"status"`

	// ProgressMessage: last progress message, if the job has started.
	ProgressMessage *string `json:"progress_message"`

	// JobDuration: duration of the job, if the job is finished.
	JobDuration *scw.Duration `json:"job_duration"`

	// ResultDistribution: result of the job, if the job is finished.
	ResultDistribution *string `json:"result_distribution"`
}

// Platform: platform.
type Platform struct {
	// ID: unique ID of the platform.
	ID string `json:"id"`

	// Version: verison of the platform.
	Version string `json:"version"`

	// Name: name of the platform.
	Name string `json:"name"`

	// ProviderName: provider name of the platform.
	ProviderName string `json:"provider_name"`

	// Type: type of the platform.
	// Default value: unknown_type
	Type PlatformType `json:"type"`

	// Technology: technology used by the platform.
	// Default value: unknown_technology
	Technology PlatformTechnology `json:"technology"`

	// MaxQubitCount: maximum number of qubits supported by the platform.
	MaxQubitCount uint32 `json:"max_qubit_count"`

	// Metadata: metadata provided by the platform.
	Metadata string `json:"metadata"`
}

// Session: session.
type Session struct {
	// ID: unique ID of the session.
	ID string `json:"id"`

	// Name: name of the session.
	Name string `json:"name"`

	// PlatformID: platform ID for which the session has been created.
	PlatformID string `json:"platform_id"`

	// CreatedAt: the time at which the session was created.
	CreatedAt *time.Time `json:"created_at"`

	// StartedAt: the time at which the session started.
	StartedAt *time.Time `json:"started_at"`

	// UpdatedAt: the time at which the session was updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// TerminatedAt: the time at which the session was terminated.
	TerminatedAt *time.Time `json:"terminated_at"`

	// MaxIDleDuration: maximum idle time before the session ends.
	MaxIDleDuration *scw.Duration `json:"max_idle_duration"`

	// MaxDuration: maximum duration before the session ends.
	MaxDuration *scw.Duration `json:"max_duration"`

	// WaitingJobCount: number of waiting jobs linked to the session.
	WaitingJobCount uint64 `json:"waiting_job_count"`

	// FinishedJobCount: number of finished jobs linked to the session.
	FinishedJobCount uint64 `json:"finished_job_count"`

	// Status: status of the session.
	// Default value: unknown_status
	Status SessionStatus `json:"status"`

	// ProjectID: project ID in which the session has been created.
	ProjectID string `json:"project_id"`

	// Tags: tags of the session.
	Tags *[]string `json:"tags"`

	// DeduplicationID: deduplication ID of the session.
	DeduplicationID string `json:"deduplication_id"`
}

// CancelJobRequest: cancel job request.
type CancelJobRequest struct {
	// JobID: unique ID of the job.
	JobID string `json:"-"`
}

// CreateJobRequest: create job request.
type CreateJobRequest struct {
	// Name: name of the job.
	Name string `json:"name"`

	// Tags: tags of the job.
	Tags *[]string `json:"tags,omitempty"`

	// SessionID: session in which the job is executed.
	SessionID string `json:"session_id"`

	// Circuit: quantum circuit that should be executed.
	Circuit *JobCircuit `json:"circuit"`

	// MaxDuration: maximum duration of the job.
	MaxDuration *scw.Duration `json:"max_duration,omitempty"`
}

// CreateSessionRequest: create session request.
type CreateSessionRequest struct {
	// ProjectID: ID of the Project in which the session was created.
	ProjectID string `json:"project_id"`

	// PlatformID: ID of the Platform for which the session was created.
	PlatformID string `json:"platform_id"`

	// Name: name of the session.
	Name *string `json:"name,omitempty"`

	// MaxIDleDuration: maximum idle duration before the session ends.
	MaxIDleDuration *scw.Duration `json:"max_idle_duration,omitempty"`

	// MaxDuration: maximum duration before the session ends.
	MaxDuration *scw.Duration `json:"max_duration,omitempty"`

	// Tags: tags of the session.
	Tags *[]string `json:"tags,omitempty"`

	// DeduplicationID: deduplication ID of the session.
	DeduplicationID *string `json:"deduplication_id,omitempty"`
}

// DeleteJobRequest: delete job request.
type DeleteJobRequest struct {
	// JobID: unique ID of the job.
	JobID string `json:"-"`
}

// DeleteSessionRequest: delete session request.
type DeleteSessionRequest struct {
	// SessionID: unique ID of the session.
	SessionID string `json:"-"`
}

// GetJobCircuitRequest: get job circuit request.
type GetJobCircuitRequest struct {
	// JobID: unique ID of the job.
	JobID string `json:"-"`
}

// GetJobRequest: get job request.
type GetJobRequest struct {
	// JobID: unique ID of the job you want to get.
	JobID string `json:"-"`
}

// GetPlatformRequest: get platform request.
type GetPlatformRequest struct {
	// PlatformID: unique ID of the platform.
	PlatformID string `json:"-"`
}

// GetSessionRequest: get session request.
type GetSessionRequest struct {
	// SessionID: unique ID of the session.
	SessionID string `json:"-"`
}

// ListJobsRequest: list jobs request.
type ListJobsRequest struct {
	// SessionID: list jobs with this session ID.
	// Precisely one of SessionID, ProjectID must be set.
	SessionID *string `json:"session_id,omitempty"`

	// ProjectID: list jobs with this project ID.
	// Precisely one of SessionID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// Tags: list jobs with these tags.
	Tags *[]string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of jobs to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned jobs.
	// Default value: created_at_desc
	OrderBy ListJobsRequestOrderBy `json:"-"`
}

// ListJobsResponse: list jobs response.
type ListJobsResponse struct {
	// TotalCount: total number of jobs.
	TotalCount uint64 `json:"total_count"`

	// Jobs: list of jobs.
	Jobs []*Job `json:"jobs"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListJobsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListJobsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListJobsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Jobs = append(r.Jobs, results.Jobs...)
	r.TotalCount += uint64(len(results.Jobs))
	return uint64(len(results.Jobs)), nil
}

// ListPlatformsRequest: list platforms request.
type ListPlatformsRequest struct {
	// ProviderName: list platforms with this provider name.
	ProviderName *string `json:"-"`

	// Name: list platforms with this name.
	Name *string `json:"-"`

	// PlatformType: list platforms with this type.
	// Default value: unknown_type
	PlatformType PlatformType `json:"-"`

	// PlatformTechnology: list platforms with this technology.
	// Default value: unknown_technology
	PlatformTechnology PlatformTechnology `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of platforms to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned platforms.
	// Default value: name_asc
	OrderBy ListPlatformsRequestOrderBy `json:"-"`
}

// ListPlatformsResponse: list platforms response.
type ListPlatformsResponse struct {
	// TotalCount: total number of platforms.
	TotalCount uint64 `json:"total_count"`

	// Platforms: list of platforms.
	Platforms []*Platform `json:"platforms"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPlatformsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPlatformsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPlatformsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Platforms = append(r.Platforms, results.Platforms...)
	r.TotalCount += uint64(len(results.Platforms))
	return uint64(len(results.Platforms)), nil
}

// ListSessionsRequest: list sessions request.
type ListSessionsRequest struct {
	// PlatformID: list sessions that have been created for this platform.
	PlatformID *string `json:"-"`

	// Tags: list sessions with these tags.
	Tags *[]string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of sessions to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned sessions.
	// Default value: name_asc
	OrderBy ListSessionsRequestOrderBy `json:"-"`

	// ProjectID: list sessions belonging to this project ID.
	ProjectID string `json:"-"`
}

// ListSessionsResponse: list sessions response.
type ListSessionsResponse struct {
	// TotalCount: total number of sessions.
	TotalCount uint64 `json:"total_count"`

	// Sessions: list of sessions.
	Sessions []*Session `json:"sessions"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSessionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSessionsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListSessionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Sessions = append(r.Sessions, results.Sessions...)
	r.TotalCount += uint64(len(results.Sessions))
	return uint64(len(results.Sessions)), nil
}

// TerminateSessionRequest: terminate session request.
type TerminateSessionRequest struct {
	// SessionID: unique ID of the session.
	SessionID string `json:"-"`
}

// UpdateJobRequest: update job request.
type UpdateJobRequest struct {
	// JobID: unique ID of the job.
	JobID string `json:"-"`

	// Name: name of the job.
	Name *string `json:"name,omitempty"`

	// Tags: tags of the job.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateSessionRequest: update session request.
type UpdateSessionRequest struct {
	// SessionID: unique ID of the session.
	SessionID string `json:"-"`

	// Name: name of the session.
	Name *string `json:"name,omitempty"`

	// MaxIDleDuration: maximum idle duration before the session ends.
	MaxIDleDuration *scw.Duration `json:"max_idle_duration,omitempty"`

	// MaxDuration: maximum time before the session ends.
	MaxDuration *scw.Duration `json:"max_duration,omitempty"`

	// Tags: tags of the session.
	Tags *[]string `json:"tags,omitempty"`
}

// This API allows you to manage Scaleway Quantum as a Service.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// GetJob: Retrieve information about the provided **job ID**, such as status, payload, and result.
func (s *API) GetJob(req *GetJobRequest, opts ...scw.RequestOption) (*Job, error) {
	var err error

	if fmt.Sprint(req.JobID) == "" {
		return nil, errors.New("field JobID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/jobs/" + fmt.Sprint(req.JobID) + "",
	}

	var resp Job

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListJobs: Retrieve information about all jobs within a given project or session.
func (s *API) ListJobs(req *ListJobsRequest, opts ...scw.RequestOption) (*ListJobsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.SessionID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "session_id", req.SessionID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/jobs",
		Query:  query,
	}

	var resp ListJobsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateJob: Create a job to be executed inside a session.
func (s *API) CreateJob(req *CreateJobRequest, opts ...scw.RequestOption) (*Job, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/qaas/v1alpha1/jobs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Job

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateJob: Update job information about the provided **job ID**.
func (s *API) UpdateJob(req *UpdateJobRequest, opts ...scw.RequestOption) (*Job, error) {
	var err error

	if fmt.Sprint(req.JobID) == "" {
		return nil, errors.New("field JobID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/qaas/v1alpha1/jobs/" + fmt.Sprint(req.JobID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Job

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelJob: Cancel the job corresponding to the provided **job ID**.
func (s *API) CancelJob(req *CancelJobRequest, opts ...scw.RequestOption) (*Job, error) {
	var err error

	if fmt.Sprint(req.JobID) == "" {
		return nil, errors.New("field JobID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/qaas/v1alpha1/jobs/" + fmt.Sprint(req.JobID) + "/cancel",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Job

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteJob: Delete the job corresponding to the provided **job ID**.
func (s *API) DeleteJob(req *DeleteJobRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.JobID) == "" {
		return errors.New("field JobID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/qaas/v1alpha1/jobs/" + fmt.Sprint(req.JobID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetJobCircuit: Retrieve the circuit of the provided **job ID**.
func (s *API) GetJobCircuit(req *GetJobCircuitRequest, opts ...scw.RequestOption) (*JobCircuit, error) {
	var err error

	if fmt.Sprint(req.JobID) == "" {
		return nil, errors.New("field JobID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/jobs/" + fmt.Sprint(req.JobID) + "/circuit",
	}

	var resp JobCircuit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPlatform: Retrieve information about the provided **platform ID**, such as provider name, technology, and type.
func (s *API) GetPlatform(req *GetPlatformRequest, opts ...scw.RequestOption) (*Platform, error) {
	var err error

	if fmt.Sprint(req.PlatformID) == "" {
		return nil, errors.New("field PlatformID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/platforms/" + fmt.Sprint(req.PlatformID) + "",
	}

	var resp Platform

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPlatforms: Retrieve information about all platforms.
func (s *API) ListPlatforms(req *ListPlatformsRequest, opts ...scw.RequestOption) (*ListPlatformsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "provider_name", req.ProviderName)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "platform_type", req.PlatformType)
	parameter.AddToQuery(query, "platform_technology", req.PlatformTechnology)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/platforms",
		Query:  query,
	}

	var resp ListPlatformsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSession: Retrieve information about the provided **session ID**, such as name, status, and number of executed jobs.
func (s *API) GetSession(req *GetSessionRequest, opts ...scw.RequestOption) (*Session, error) {
	var err error

	if fmt.Sprint(req.SessionID) == "" {
		return nil, errors.New("field SessionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/sessions/" + fmt.Sprint(req.SessionID) + "",
	}

	var resp Session

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSessions: Retrieve information about all sessions.
func (s *API) ListSessions(req *ListSessionsRequest, opts ...scw.RequestOption) (*ListSessionsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "platform_id", req.PlatformID)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/sessions",
		Query:  query,
	}

	var resp ListSessionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSession: Create a dedicated session for the specified platform.
func (s *API) CreateSession(req *CreateSessionRequest, opts ...scw.RequestOption) (*Session, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/qaas/v1alpha1/sessions",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Session

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSession: Update session information of the provided **session ID**.
func (s *API) UpdateSession(req *UpdateSessionRequest, opts ...scw.RequestOption) (*Session, error) {
	var err error

	if fmt.Sprint(req.SessionID) == "" {
		return nil, errors.New("field SessionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/qaas/v1alpha1/sessions/" + fmt.Sprint(req.SessionID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Session

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TerminateSession: Terminate a session by its unique ID and cancel all its attached jobs.
func (s *API) TerminateSession(req *TerminateSessionRequest, opts ...scw.RequestOption) (*Session, error) {
	var err error

	if fmt.Sprint(req.SessionID) == "" {
		return nil, errors.New("field SessionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/qaas/v1alpha1/sessions/" + fmt.Sprint(req.SessionID) + "/terminate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Session

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSession: Delete a session by its unique ID and delete all its attached jobs.
func (s *API) DeleteSession(req *DeleteSessionRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.SessionID) == "" {
		return errors.New("field SessionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/qaas/v1alpha1/sessions/" + fmt.Sprint(req.SessionID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

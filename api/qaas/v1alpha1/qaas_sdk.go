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

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultQaasRetryInterval = 15 * time.Second
	defaultQaasTimeout       = 5 * time.Minute
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

type ApplicationType string

const (
	ApplicationTypeUnknownType = ApplicationType("unknown_type")
	// Variational Quantum Eigensolver is a type hybrid algorithm to find the ground state of a given physical system.
	ApplicationTypeVqe = ApplicationType("vqe")
)

func (enum ApplicationType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ApplicationTypeUnknownType)
	}
	return string(enum)
}

func (enum ApplicationType) Values() []ApplicationType {
	return []ApplicationType{
		"unknown_type",
		"vqe",
	}
}

func (enum ApplicationType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ApplicationType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ApplicationType(ApplicationType(tmp).String())
	return nil
}

type BookingStatus string

const (
	BookingStatusUnknownStatus = BookingStatus("unknown_status")
	BookingStatusWaiting       = BookingStatus("waiting")
	BookingStatusValidating    = BookingStatus("validating")
	BookingStatusValidated     = BookingStatus("validated")
	BookingStatusCancelling    = BookingStatus("cancelling")
	BookingStatusCancelled     = BookingStatus("cancelled")
	BookingStatusError         = BookingStatus("error")
)

func (enum BookingStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(BookingStatusUnknownStatus)
	}
	return string(enum)
}

func (enum BookingStatus) Values() []BookingStatus {
	return []BookingStatus{
		"unknown_status",
		"waiting",
		"validating",
		"validated",
		"cancelling",
		"cancelled",
		"error",
	}
}

func (enum BookingStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BookingStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BookingStatus(BookingStatus(tmp).String())
	return nil
}

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
		return string(JobStatusUnknownStatus)
	}
	return string(enum)
}

func (enum JobStatus) Values() []JobStatus {
	return []JobStatus{
		"unknown_status",
		"waiting",
		"error",
		"running",
		"completed",
		"cancelling",
		"cancelled",
	}
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

type ListApplicationsRequestOrderBy string

const (
	ListApplicationsRequestOrderByNameAsc  = ListApplicationsRequestOrderBy("name_asc")
	ListApplicationsRequestOrderByNameDesc = ListApplicationsRequestOrderBy("name_desc")
	ListApplicationsRequestOrderByTypeAsc  = ListApplicationsRequestOrderBy("type_asc")
	ListApplicationsRequestOrderByTypeDesc = ListApplicationsRequestOrderBy("type_desc")
)

func (enum ListApplicationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListApplicationsRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListApplicationsRequestOrderBy) Values() []ListApplicationsRequestOrderBy {
	return []ListApplicationsRequestOrderBy{
		"name_asc",
		"name_desc",
		"type_asc",
		"type_desc",
	}
}

func (enum ListApplicationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListApplicationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListApplicationsRequestOrderBy(ListApplicationsRequestOrderBy(tmp).String())
	return nil
}

type ListBookingsRequestOrderBy string

const (
	ListBookingsRequestOrderByCreatedAtDesc = ListBookingsRequestOrderBy("created_at_desc")
	ListBookingsRequestOrderByCreatedAtAsc  = ListBookingsRequestOrderBy("created_at_asc")
	ListBookingsRequestOrderByStartedAtDesc = ListBookingsRequestOrderBy("started_at_desc")
	ListBookingsRequestOrderByStartedAtAsc  = ListBookingsRequestOrderBy("started_at_asc")
)

func (enum ListBookingsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListBookingsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListBookingsRequestOrderBy) Values() []ListBookingsRequestOrderBy {
	return []ListBookingsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"started_at_desc",
		"started_at_asc",
	}
}

func (enum ListBookingsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListBookingsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListBookingsRequestOrderBy(ListBookingsRequestOrderBy(tmp).String())
	return nil
}

type ListJobResultsRequestOrderBy string

const (
	ListJobResultsRequestOrderByCreatedAtDesc = ListJobResultsRequestOrderBy("created_at_desc")
	ListJobResultsRequestOrderByCreatedAtAsc  = ListJobResultsRequestOrderBy("created_at_asc")
)

func (enum ListJobResultsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListJobResultsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListJobResultsRequestOrderBy) Values() []ListJobResultsRequestOrderBy {
	return []ListJobResultsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
	}
}

func (enum ListJobResultsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListJobResultsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListJobResultsRequestOrderBy(ListJobResultsRequestOrderBy(tmp).String())
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
		return string(ListJobsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListJobsRequestOrderBy) Values() []ListJobsRequestOrderBy {
	return []ListJobsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"status_asc",
		"status_desc",
		"platform_name_asc",
		"platform_name_desc",
		"name_asc",
		"name_desc",
		"session_name_asc",
		"session_name_desc",
	}
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

type ListModelsRequestOrderBy string

const (
	ListModelsRequestOrderByCreatedAtDesc = ListModelsRequestOrderBy("created_at_desc")
	ListModelsRequestOrderByCreatedAtAsc  = ListModelsRequestOrderBy("created_at_asc")
)

func (enum ListModelsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListModelsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListModelsRequestOrderBy) Values() []ListModelsRequestOrderBy {
	return []ListModelsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
	}
}

func (enum ListModelsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListModelsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListModelsRequestOrderBy(ListModelsRequestOrderBy(tmp).String())
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
	ListPlatformsRequestOrderByBackendNameAsc   = ListPlatformsRequestOrderBy("backend_name_asc")
	ListPlatformsRequestOrderByBackendNameDesc  = ListPlatformsRequestOrderBy("backend_name_desc")
)

func (enum ListPlatformsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPlatformsRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListPlatformsRequestOrderBy) Values() []ListPlatformsRequestOrderBy {
	return []ListPlatformsRequestOrderBy{
		"name_asc",
		"name_desc",
		"provider_name_asc",
		"provider_name_desc",
		"type_asc",
		"type_desc",
		"technology_asc",
		"technology_desc",
		"backend_name_asc",
		"backend_name_desc",
	}
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

type ListProcessResultsRequestOrderBy string

const (
	ListProcessResultsRequestOrderByCreatedAtDesc = ListProcessResultsRequestOrderBy("created_at_desc")
	ListProcessResultsRequestOrderByCreatedAtAsc  = ListProcessResultsRequestOrderBy("created_at_asc")
)

func (enum ListProcessResultsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListProcessResultsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListProcessResultsRequestOrderBy) Values() []ListProcessResultsRequestOrderBy {
	return []ListProcessResultsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
	}
}

func (enum ListProcessResultsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListProcessResultsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListProcessResultsRequestOrderBy(ListProcessResultsRequestOrderBy(tmp).String())
	return nil
}

type ListProcessesRequestOrderBy string

const (
	ListProcessesRequestOrderByCreatedAtDesc = ListProcessesRequestOrderBy("created_at_desc")
	ListProcessesRequestOrderByCreatedAtAsc  = ListProcessesRequestOrderBy("created_at_asc")
	ListProcessesRequestOrderByNameAsc       = ListProcessesRequestOrderBy("name_asc")
	ListProcessesRequestOrderByNameDesc      = ListProcessesRequestOrderBy("name_desc")
	ListProcessesRequestOrderByStartedAtAsc  = ListProcessesRequestOrderBy("started_at_asc")
	ListProcessesRequestOrderByStartedAtDesc = ListProcessesRequestOrderBy("started_at_desc")
	ListProcessesRequestOrderByStatusAsc     = ListProcessesRequestOrderBy("status_asc")
	ListProcessesRequestOrderByStatusDesc    = ListProcessesRequestOrderBy("status_desc")
)

func (enum ListProcessesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListProcessesRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListProcessesRequestOrderBy) Values() []ListProcessesRequestOrderBy {
	return []ListProcessesRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"name_asc",
		"name_desc",
		"started_at_asc",
		"started_at_desc",
		"status_asc",
		"status_desc",
	}
}

func (enum ListProcessesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListProcessesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListProcessesRequestOrderBy(ListProcessesRequestOrderBy(tmp).String())
	return nil
}

type ListSessionACLsRequestOrderBy string

const (
	ListSessionACLsRequestOrderByAccessAsc  = ListSessionACLsRequestOrderBy("access_asc")
	ListSessionACLsRequestOrderByAccessDesc = ListSessionACLsRequestOrderBy("access_desc")
)

func (enum ListSessionACLsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListSessionACLsRequestOrderByAccessAsc)
	}
	return string(enum)
}

func (enum ListSessionACLsRequestOrderBy) Values() []ListSessionACLsRequestOrderBy {
	return []ListSessionACLsRequestOrderBy{
		"access_asc",
		"access_desc",
	}
}

func (enum ListSessionACLsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSessionACLsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSessionACLsRequestOrderBy(ListSessionACLsRequestOrderBy(tmp).String())
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
		return string(ListSessionsRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListSessionsRequestOrderBy) Values() []ListSessionsRequestOrderBy {
	return []ListSessionsRequestOrderBy{
		"name_asc",
		"name_desc",
		"started_at_asc",
		"started_at_desc",
		"status_asc",
		"status_desc",
		"created_at_desc",
		"created_at_asc",
	}
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

type PlatformAvailability string

const (
	PlatformAvailabilityUnknownAvailability = PlatformAvailability("unknown_availability")
	PlatformAvailabilityAvailable           = PlatformAvailability("available")
	PlatformAvailabilityShortage            = PlatformAvailability("shortage")
	PlatformAvailabilityScarce              = PlatformAvailability("scarce")
	PlatformAvailabilityMaintenance         = PlatformAvailability("maintenance")
)

func (enum PlatformAvailability) String() string {
	if enum == "" {
		// return default value if empty
		return string(PlatformAvailabilityUnknownAvailability)
	}
	return string(enum)
}

func (enum PlatformAvailability) Values() []PlatformAvailability {
	return []PlatformAvailability{
		"unknown_availability",
		"available",
		"shortage",
		"scarce",
		"maintenance",
	}
}

func (enum PlatformAvailability) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlatformAvailability) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlatformAvailability(PlatformAvailability(tmp).String())
	return nil
}

type PlatformTechnology string

const (
	PlatformTechnologyUnknownTechnology = PlatformTechnology("unknown_technology")
	PlatformTechnologyPhotonic          = PlatformTechnology("photonic")
	PlatformTechnologyGeneralPurpose    = PlatformTechnology("general_purpose")
	PlatformTechnologyTrappedIon        = PlatformTechnology("trapped_ion")
	PlatformTechnologySuperconducting   = PlatformTechnology("superconducting")
	PlatformTechnologyNeutralAtom       = PlatformTechnology("neutral_atom")
)

func (enum PlatformTechnology) String() string {
	if enum == "" {
		// return default value if empty
		return string(PlatformTechnologyUnknownTechnology)
	}
	return string(enum)
}

func (enum PlatformTechnology) Values() []PlatformTechnology {
	return []PlatformTechnology{
		"unknown_technology",
		"photonic",
		"general_purpose",
		"trapped_ion",
		"superconducting",
		"neutral_atom",
	}
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
		return string(PlatformTypeUnknownType)
	}
	return string(enum)
}

func (enum PlatformType) Values() []PlatformType {
	return []PlatformType{
		"unknown_type",
		"simulator",
		"qpu",
	}
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

type ProcessStatus string

const (
	ProcessStatusUnknownStatus = ProcessStatus("unknown_status")
	ProcessStatusError         = ProcessStatus("error")
	ProcessStatusStarting      = ProcessStatus("starting")
	ProcessStatusRunning       = ProcessStatus("running")
	ProcessStatusCompleted     = ProcessStatus("completed")
	ProcessStatusCancelling    = ProcessStatus("cancelling")
	ProcessStatusCancelled     = ProcessStatus("cancelled")
)

func (enum ProcessStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ProcessStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ProcessStatus) Values() []ProcessStatus {
	return []ProcessStatus{
		"unknown_status",
		"error",
		"starting",
		"running",
		"completed",
		"cancelling",
		"cancelled",
	}
}

func (enum ProcessStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ProcessStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ProcessStatus(ProcessStatus(tmp).String())
	return nil
}

type SessionAccess string

const (
	SessionAccessUnknownAccess    = SessionAccess("unknown_access")
	SessionAccessFull             = SessionAccess("full")
	SessionAccessReadSession      = SessionAccess("read_session")
	SessionAccessReadWriteSession = SessionAccess("read_write_session")
	SessionAccessReadJobResult    = SessionAccess("read_job_result")
	SessionAccessReadJobCircuit   = SessionAccess("read_job_circuit")
	SessionAccessReadJob          = SessionAccess("read_job")
	SessionAccessReadWriteJob     = SessionAccess("read_write_job")
)

func (enum SessionAccess) String() string {
	if enum == "" {
		// return default value if empty
		return string(SessionAccessUnknownAccess)
	}
	return string(enum)
}

func (enum SessionAccess) Values() []SessionAccess {
	return []SessionAccess{
		"unknown_access",
		"full",
		"read_session",
		"read_write_session",
		"read_job_result",
		"read_job_circuit",
		"read_job",
		"read_write_job",
	}
}

func (enum SessionAccess) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SessionAccess) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SessionAccess(SessionAccess(tmp).String())
	return nil
}

type SessionOriginType string

const (
	SessionOriginTypeUnknownOriginType = SessionOriginType("unknown_origin_type")
	SessionOriginTypeCustomer          = SessionOriginType("customer")
	SessionOriginTypeProcess           = SessionOriginType("process")
)

func (enum SessionOriginType) String() string {
	if enum == "" {
		// return default value if empty
		return string(SessionOriginTypeUnknownOriginType)
	}
	return string(enum)
}

func (enum SessionOriginType) Values() []SessionOriginType {
	return []SessionOriginType{
		"unknown_origin_type",
		"customer",
		"process",
	}
}

func (enum SessionOriginType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SessionOriginType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SessionOriginType(SessionOriginType(tmp).String())
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
		return string(SessionStatusUnknownStatus)
	}
	return string(enum)
}

func (enum SessionStatus) Values() []SessionStatus {
	return []SessionStatus{
		"unknown_status",
		"running",
		"stopped",
		"starting",
		"stopping",
	}
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

// PlatformBookingRequirement: platform booking requirement.
type PlatformBookingRequirement struct {
	// MinDuration: minimal duration of any booking based on this platform.
	MinDuration *scw.Duration `json:"min_duration"`

	// MaxDuration: maximal duration of any bookings based on this platform.
	MaxDuration *scw.Duration `json:"max_duration"`

	// MaxCancellationDuration: allowed time to cancel a booking attached to this platform before the beginning of the session.
	MaxCancellationDuration *scw.Duration `json:"max_cancellation_duration"`

	// MaxPlanificationDuration: allowed planification time from now where the platform can be booked in the future.
	MaxPlanificationDuration *scw.Duration `json:"max_planification_duration"`

	// MinPlanificationDuration: minimum planification time before a platform can be booked.
	MinPlanificationDuration *scw.Duration `json:"min_planification_duration"`

	// MaxBookingPerWeek: maximum amount of booking allowed for one organization per week.
	MaxBookingPerWeek uint32 `json:"max_booking_per_week"`

	// MaxBookingPerDay: maximum amount of booking allowed for one organization per day.
	MaxBookingPerDay uint32 `json:"max_booking_per_day"`
}

// PlatformHardware: platform hardware.
type PlatformHardware struct {
	// Name: product name of the hardware.
	Name string `json:"name"`

	// Vcpus: number of vCPUs available.
	Vcpus uint32 `json:"vcpus"`

	// Gpus: number of GPUs available (0 if no GPU).
	Gpus uint32 `json:"gpus"`

	// GpusNetwork: network topology of GPUs (PCIe, NVLink...).
	GpusNetwork string `json:"gpus_network"`

	// RAM: amount of RAM available in byte.
	RAM uint64 `json:"ram"`

	// Vram: amount of VRAM available in byte (0 if no GPU).
	Vram uint64 `json:"vram"`
}

// JobCircuit: job circuit.
type JobCircuit struct {
	// PercevalCircuit: circuit generated by Perceval that should be executed.
	// Precisely one of PercevalCircuit, QiskitCircuit must be set.
	PercevalCircuit *string `json:"perceval_circuit,omitempty"`

	// QiskitCircuit: circuit generated by Qiskit that should be executed.
	// Precisely one of PercevalCircuit, QiskitCircuit must be set.
	QiskitCircuit *string `json:"qiskit_circuit,omitempty"`
}

// CreateSessionRequestBookingDemand: create session request booking demand.
type CreateSessionRequestBookingDemand struct {
	StartedAt *time.Time `json:"started_at"`

	FinishedAt *time.Time `json:"finished_at"`

	Description *string `json:"description"`

	TimeZone *string `json:"time_zone"`
}

// Application: application.
type Application struct {
	// ID: unique ID of the application.
	ID string `json:"id"`

	// Name: name of the application.
	Name string `json:"name"`

	// Type: type of the application.
	// Default value: unknown_type
	Type ApplicationType `json:"type"`

	// CompatiblePlatformIDs: list of compatible platform (by IDs) able to run this application.
	CompatiblePlatformIDs []string `json:"compatible_platform_ids"`

	// InputTemplate: JSON format describing the expected input.
	InputTemplate string `json:"input_template"`
}

// Booking: booking.
type Booking struct {
	// ID: unique ID of the booking.
	ID string `json:"id"`

	// CreatedAt: time at which the booking was created.
	CreatedAt *time.Time `json:"created_at"`

	// StartedAt: time at which the booking starts.
	StartedAt *time.Time `json:"started_at"`

	// UpdatedAt: time at which the booking was updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// FinishedAt: time at which the booking finishes.
	FinishedAt *time.Time `json:"finished_at"`

	// Status: status of the booking.
	// Default value: unknown_status
	Status BookingStatus `json:"status"`

	// Description: description of the booking slot.
	Description string `json:"description"`

	// ProgressMessage: any progress message of the booking.
	ProgressMessage string `json:"progress_message"`

	// TimeZone: time zone for the booking schedule, in tz database format (e.g. 'Europe/Paris').
	TimeZone *string `json:"time_zone"`
}

// JobResult: job result.
type JobResult struct {
	// JobID: ID of the parent job.
	JobID string `json:"job_id"`

	// Result: result in JSON format.
	Result *string `json:"result"`

	// URL: URL to download a large result (optional).
	URL *string `json:"url"`

	// CreatedAt: creation time of the result.
	CreatedAt *time.Time `json:"created_at"`
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

	// ModelID: computation model ID executed by the job.
	ModelID *string `json:"model_id"`

	// Parameters: execution parameters for this job.
	Parameters *string `json:"parameters"`
}

// Model: model.
type Model struct {
	// ID: unique ID of the model.
	ID string `json:"id"`

	// CreatedAt: time at which the model was created.
	CreatedAt *time.Time `json:"created_at"`

	// URL: storage URL of the model.
	URL *string `json:"url"`

	// ProjectID: project ID in which the model has been created.
	ProjectID string `json:"project_id"`
}

// Platform: platform.
type Platform struct {
	// ID: unique ID of the platform.
	ID string `json:"id"`

	// Version: version of the platform.
	Version string `json:"version"`

	// Name: name of the platform.
	Name string `json:"name"`

	// ProviderName: name of the technological provider of the platform in lowercase (quandela, pasqal...).
	ProviderName string `json:"provider_name"`

	// BackendName: name of the running emulation backend or QPU model of the platform in lowercase (mosaiq, qsim, aer...).
	BackendName string `json:"backend_name"`

	// Type: type of the platform (emulator or qpu).
	// Default value: unknown_type
	Type PlatformType `json:"type"`

	// Technology: quantum technology used by the platform (trapped-ion, photonic, superconducting qubits...).
	// Default value: unknown_technology
	Technology PlatformTechnology `json:"technology"`

	// MaxQubitCount: maximum number of qubits supported by the platform (estimated for emulator).
	MaxQubitCount uint32 `json:"max_qubit_count"`

	// MaxShotCount: maximum number of shots during a circuit execution.
	MaxShotCount uint32 `json:"max_shot_count"`

	// MaxCircuitCount: maximum number of circuit that can be executed in one call.
	MaxCircuitCount uint32 `json:"max_circuit_count"`

	// Availability: availability of the platform.
	// Default value: unknown_availability
	Availability PlatformAvailability `json:"availability"`

	// Metadata: metadata provided by the platform.
	Metadata string `json:"metadata"`

	// PricePerHour: price to be paid per hour (excluding free tiers).
	PricePerHour *scw.Money `json:"price_per_hour"`

	// PricePerShot: price to be paid per shot (excluding free tiers).
	PricePerShot *scw.Money `json:"price_per_shot"`

	// PricePerCircuit: price to be paid per circuit setup before its execution (excluding free tiers).
	PricePerCircuit *scw.Money `json:"price_per_circuit"`

	// Hardware: specifications of the underlying hardware.
	Hardware *PlatformHardware `json:"hardware"`

	// BookingRequirement: booking constraints to fit if the platform is bookable.
	BookingRequirement *PlatformBookingRequirement `json:"booking_requirement"`

	// Description: english description of the platform.
	Description string `json:"description"`

	// DocumentationURL: documentation link to external documentation to learn more on this platform.
	DocumentationURL string `json:"documentation_url"`

	// IsBookable: specify if the platform is bookable.
	IsBookable bool `json:"is_bookable"`
}

// ProcessResult: process result.
type ProcessResult struct {
	// ProcessID: ID of the parent process.
	ProcessID string `json:"process_id"`

	// Result: result in JSON format.
	Result string `json:"result"`

	// CreatedAt: creation time of the result.
	CreatedAt *time.Time `json:"created_at"`
}

// Process: process.
type Process struct {
	// ID: unique ID of the process.
	ID string `json:"id"`

	// Name: name of the process.
	Name string `json:"name"`

	// ApplicationID: application ID for which the process has been created.
	ApplicationID *string `json:"application_id"`

	// PlatformID: platform ID for which the process has been created.
	PlatformID *string `json:"platform_id"`

	// AttachedSessionIDs: list of sessions generated by the process.
	AttachedSessionIDs []string `json:"attached_session_ids"`

	// CreatedAt: time at which the process was created.
	CreatedAt *time.Time `json:"created_at"`

	// StartedAt: time at which the process started.
	StartedAt *time.Time `json:"started_at"`

	// UpdatedAt: time at which the process was updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// FinishedAt: time at which the process was finished.
	FinishedAt *time.Time `json:"finished_at"`

	// Status: status of the process.
	// Default value: unknown_status
	Status ProcessStatus `json:"status"`

	// ProjectID: project ID in which the process has been created.
	ProjectID string `json:"project_id"`

	// Tags: tags of the process.
	Tags []string `json:"tags"`

	// Progress: progress of the process, from 0 to 1.
	Progress *uint32 `json:"progress"`

	// ProgressMessage: any progress of the process.
	ProgressMessage *string `json:"progress_message"`

	// Input: input payload of the process as JSON string.
	Input *string `json:"input"`
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

	// OriginType: resource type that creates the session.
	// Default value: unknown_origin_type
	OriginType SessionOriginType `json:"origin_type"`

	// OriginID: unique ID of the session's origin resource (if exists).
	OriginID *string `json:"origin_id"`

	// ProgressMessage: any progress of the session.
	ProgressMessage *string `json:"progress_message"`

	// BookingID: an optional booking unique ID of an attached booking.
	BookingID *string `json:"booking_id"`

	// ModelID: default computation model ID to be executed by job assigned to this session.
	ModelID *string `json:"model_id"`

	// Parameters: platform configuration parameters applied to this session.
	Parameters *string `json:"parameters"`
}

// CancelJobRequest: cancel job request.
type CancelJobRequest struct {
	// JobID: unique ID of the job.
	JobID string `json:"-"`
}

// CancelProcessRequest: cancel process request.
type CancelProcessRequest struct {
	// ProcessID: unique ID of the process.
	ProcessID string `json:"-"`
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

	// ModelID: computation model ID to be executed by the job.
	ModelID *string `json:"model_id,omitempty"`

	// Parameters: execution parameters for this job.
	Parameters *string `json:"parameters,omitempty"`
}

// CreateModelRequest: create model request.
type CreateModelRequest struct {
	// ProjectID: project ID to attach this model.
	ProjectID string `json:"project_id"`

	// Payload: the serialized model data.
	Payload *string `json:"payload,omitempty"`
}

// CreateProcessRequest: create process request.
type CreateProcessRequest struct {
	// ProjectID: ID of the project in which the process was created.
	ProjectID string `json:"project_id"`

	// PlatformID: ID of the platform for which the process was created.
	PlatformID *string `json:"platform_id,omitempty"`

	// ApplicationID: ID of the application for which the process was created.
	ApplicationID *string `json:"application_id,omitempty"`

	// Name: name of the process.
	Name string `json:"name"`

	// Input: process parameters in JSON format.
	Input *string `json:"input,omitempty"`

	// Tags: tags of the process.
	Tags []string `json:"tags"`
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

	// BookingDemand: a booking demand to schedule the session, only applicable if the platform is bookable.
	BookingDemand *CreateSessionRequestBookingDemand `json:"booking_demand,omitempty"`

	// ModelID: default computation model ID to be executed by job assigned to this session.
	ModelID *string `json:"model_id,omitempty"`

	// Parameters: optional platform configuration parameters applied to this session.
	Parameters *string `json:"parameters,omitempty"`
}

// DeleteJobRequest: delete job request.
type DeleteJobRequest struct {
	// JobID: unique ID of the job.
	JobID string `json:"-"`
}

// DeleteProcessRequest: delete process request.
type DeleteProcessRequest struct {
	// ProcessID: unique ID of the process.
	ProcessID string `json:"-"`
}

// DeleteSessionRequest: delete session request.
type DeleteSessionRequest struct {
	// SessionID: unique ID of the session.
	SessionID string `json:"-"`
}

// GetApplicationRequest: get application request.
type GetApplicationRequest struct {
	// ApplicationID: unique ID of the application.
	ApplicationID string `json:"-"`
}

// GetBookingRequest: get booking request.
type GetBookingRequest struct {
	// BookingID: unique ID of the booking.
	BookingID string `json:"-"`
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

// GetModelRequest: get model request.
type GetModelRequest struct {
	// ModelID: unique ID of the model.
	ModelID string `json:"-"`
}

// GetPlatformRequest: get platform request.
type GetPlatformRequest struct {
	// PlatformID: unique ID of the platform.
	PlatformID string `json:"-"`
}

// GetProcessRequest: get process request.
type GetProcessRequest struct {
	// ProcessID: unique ID of the process.
	ProcessID string `json:"-"`
}

// GetSessionRequest: get session request.
type GetSessionRequest struct {
	// SessionID: unique ID of the session.
	SessionID string `json:"-"`
}

// ListApplicationsRequest: list applications request.
type ListApplicationsRequest struct {
	// Name: list applications with this name.
	Name *string `json:"-"`

	// ApplicationType: list applications with this type.
	// Default value: unknown_type
	ApplicationType ApplicationType `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of applications a to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned applications.
	// Default value: name_asc
	OrderBy ListApplicationsRequestOrderBy `json:"-"`
}

// ListApplicationsResponse: list applications response.
type ListApplicationsResponse struct {
	// TotalCount: total number of applications.
	TotalCount uint64 `json:"total_count"`

	// Applications: list of applications.
	Applications []*Application `json:"applications"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListApplicationsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListApplicationsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListApplicationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Applications = append(r.Applications, results.Applications...)
	r.TotalCount += uint64(len(results.Applications))
	return uint64(len(results.Applications)), nil
}

// ListBookingsRequest: list bookings request.
type ListBookingsRequest struct {
	// ProjectID: list bookings belonging to this project ID.
	ProjectID *string `json:"-"`

	// PlatformID: list bookings attached to this platform ID.
	PlatformID *string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of results to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned results.
	// Default value: created_at_desc
	OrderBy ListBookingsRequestOrderBy `json:"-"`
}

// ListBookingsResponse: list bookings response.
type ListBookingsResponse struct {
	// TotalCount: total number of bookings.
	TotalCount uint64 `json:"total_count"`

	// Bookings: list of bookings.
	Bookings []*Booking `json:"bookings"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBookingsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBookingsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListBookingsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Bookings = append(r.Bookings, results.Bookings...)
	r.TotalCount += uint64(len(results.Bookings))
	return uint64(len(results.Bookings)), nil
}

// ListJobResultsRequest: list job results request.
type ListJobResultsRequest struct {
	// JobID: ID of the job.
	JobID string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of results to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned results.
	// Default value: created_at_desc
	OrderBy ListJobResultsRequestOrderBy `json:"-"`
}

// ListJobResultsResponse: list job results response.
type ListJobResultsResponse struct {
	// TotalCount: total number of results.
	TotalCount uint64 `json:"total_count"`

	// JobResults: list of results.
	JobResults []*JobResult `json:"job_results"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListJobResultsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListJobResultsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListJobResultsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.JobResults = append(r.JobResults, results.JobResults...)
	r.TotalCount += uint64(len(results.JobResults))
	return uint64(len(results.JobResults)), nil
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
	Tags []string `json:"-"`

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
func (r *ListJobsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListJobsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Jobs = append(r.Jobs, results.Jobs...)
	r.TotalCount += uint64(len(results.Jobs))
	return uint64(len(results.Jobs)), nil
}

// ListModelsRequest: list models request.
type ListModelsRequest struct {
	// ProjectID: list models belonging to this project ID.
	ProjectID string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of results to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned results.
	// Default value: created_at_desc
	OrderBy ListModelsRequestOrderBy `json:"-"`
}

// ListModelsResponse: list models response.
type ListModelsResponse struct {
	// TotalCount: total number of models.
	TotalCount uint64 `json:"total_count"`

	// Models: list of models.
	Models []*Model `json:"models"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListModelsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListModelsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListModelsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Models = append(r.Models, results.Models...)
	r.TotalCount += uint64(len(results.Models))
	return uint64(len(results.Models)), nil
}

// ListPlatformsRequest: list platforms request.
type ListPlatformsRequest struct {
	// ProviderName: list platforms with this provider name.
	ProviderName *string `json:"-"`

	// BackendName: list platforms with this backend name.
	BackendName *string `json:"-"`

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
func (r *ListPlatformsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListPlatformsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Platforms = append(r.Platforms, results.Platforms...)
	r.TotalCount += uint64(len(results.Platforms))
	return uint64(len(results.Platforms)), nil
}

// ListProcessResultsRequest: list process results request.
type ListProcessResultsRequest struct {
	// ProcessID: ID of the process.
	ProcessID string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of results to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned results.
	// Default value: created_at_desc
	OrderBy ListProcessResultsRequestOrderBy `json:"-"`
}

// ListProcessResultsResponse: list process results response.
type ListProcessResultsResponse struct {
	// TotalCount: total number of results.
	TotalCount uint64 `json:"total_count"`

	// ProcessResults: list of results.
	ProcessResults []*ProcessResult `json:"process_results"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListProcessResultsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListProcessResultsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListProcessResultsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ProcessResults = append(r.ProcessResults, results.ProcessResults...)
	r.TotalCount += uint64(len(results.ProcessResults))
	return uint64(len(results.ProcessResults)), nil
}

// ListProcessesRequest: list processes request.
type ListProcessesRequest struct {
	// ApplicationID: list processes that have been created for this application.
	ApplicationID *string `json:"-"`

	// Tags: list processes with these tags.
	Tags []string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of processes to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned processes.
	// Default value: created_at_desc
	OrderBy ListProcessesRequestOrderBy `json:"-"`

	// ProjectID: list processes belonging to this project ID.
	ProjectID string `json:"-"`
}

// ListProcessesResponse: list processes response.
type ListProcessesResponse struct {
	// TotalCount: total number of processes.
	TotalCount uint64 `json:"total_count"`

	// Processes: list of processes.
	Processes []*Process `json:"processes"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListProcessesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListProcessesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListProcessesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Processes = append(r.Processes, results.Processes...)
	r.TotalCount += uint64(len(results.Processes))
	return uint64(len(results.Processes)), nil
}

// ListSessionACLsRequest: list session ac ls request.
type ListSessionACLsRequest struct {
	SessionID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: access_asc
	OrderBy ListSessionACLsRequestOrderBy `json:"-"`
}

// ListSessionACLsResponse: list session ac ls response.
type ListSessionACLsResponse struct {
	TotalCount uint64 `json:"total_count"`

	ACLs []SessionAccess `json:"acls"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSessionACLsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSessionACLsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListSessionACLsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ACLs = append(r.ACLs, results.ACLs...)
	r.TotalCount += uint64(len(results.ACLs))
	return uint64(len(results.ACLs)), nil
}

// ListSessionsRequest: list sessions request.
type ListSessionsRequest struct {
	// PlatformID: list sessions that have been created for this platform.
	PlatformID *string `json:"-"`

	// Tags: list sessions with these tags.
	Tags []string `json:"-"`

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
func (r *ListSessionsResponse) UnsafeAppend(res any) (uint64, error) {
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

// UpdateBookingRequest: update booking request.
type UpdateBookingRequest struct {
	// BookingID: unique ID of the booking.
	BookingID string `json:"-"`

	// Description: description of the booking slot.
	Description *string `json:"description,omitempty"`
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

// UpdateProcessRequest: update process request.
type UpdateProcessRequest struct {
	// ProcessID: unique ID of the process.
	ProcessID string `json:"-"`

	// Name: name of the process.
	Name *string `json:"name,omitempty"`

	// Tags: tags of the process.
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

// This API allows you to allocate and program Quantum Processing Units (QPUs) to run quantum algorithms.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// GetJob: Retrieve information about the provided **job ID**, mainly used to get the current status.
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

// WaitForJobRequest is used by WaitForJob method.
type WaitForJobRequest struct {
	JobID         string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForJob waits for the Job to reach a terminal state.
func (s *API) WaitForJob(req *WaitForJobRequest, opts ...scw.RequestOption) (*Job, error) {
	timeout := defaultQaasTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultQaasRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[JobStatus]struct{}{
		JobStatusWaiting:    {},
		JobStatusRunning:    {},
		JobStatusCancelling: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetJob(&GetJobRequest{
				JobID: req.JobID,
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
		return nil, errors.Wrap(err, "waiting for Job failed")
	}

	return res.(*Job), nil
}

// ListJobs: Retrieve information about all jobs within a given session.
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

// ListJobResults: Retrieve all intermediate and final results of a job.
func (s *API) ListJobResults(req *ListJobResultsRequest, opts ...scw.RequestOption) (*ListJobResultsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.JobID) == "" {
		return nil, errors.New("field JobID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/jobs/" + fmt.Sprint(req.JobID) + "/results",
		Query:  query,
	}

	var resp ListJobResultsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateJob: Create a job to be executed inside a QPU session.
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
	parameter.AddToQuery(query, "backend_name", req.BackendName)
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

// GetSession: Retrieve information about the provided **session ID**, such as name and status.
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

// WaitForSessionRequest is used by WaitForSession method.
type WaitForSessionRequest struct {
	SessionID     string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForSession waits for the Session to reach a terminal state.
func (s *API) WaitForSession(req *WaitForSessionRequest, opts ...scw.RequestOption) (*Session, error) {
	timeout := defaultQaasTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultQaasRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[SessionStatus]struct{}{
		SessionStatusStarting: {},
		SessionStatusStopping: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetSession(&GetSessionRequest{
				SessionID: req.SessionID,
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
		return nil, errors.Wrap(err, "waiting for Session failed")
	}

	return res.(*Session), nil
}

// ListSessions: Retrieve information about all QPU sessions.
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

// CreateSession: Create a new QPU session for the specified platform. Once ready, jobs can be sent to this session.
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

// TerminateSession: Terminate a session by its unique ID and cancel all its attached jobs and bookings.
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

// DeleteSession: Delete a session by its unique ID and delete all its attached jobs and bookings.
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

// ListSessionACLs:
func (s *API) ListSessionACLs(req *ListSessionACLsRequest, opts ...scw.RequestOption) (*ListSessionACLsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.SessionID) == "" {
		return nil, errors.New("field SessionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/sessions/" + fmt.Sprint(req.SessionID) + "/acls",
		Query:  query,
	}

	var resp ListSessionACLsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateProcess: Create a new process for the specified application on a specified platform.
func (s *API) CreateProcess(req *CreateProcessRequest, opts ...scw.RequestOption) (*Process, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/qaas/v1alpha1/processes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Process

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetProcess: Retrieve information about the provided **process ID**, such as name, status and progress.
func (s *API) GetProcess(req *GetProcessRequest, opts ...scw.RequestOption) (*Process, error) {
	var err error

	if fmt.Sprint(req.ProcessID) == "" {
		return nil, errors.New("field ProcessID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/processes/" + fmt.Sprint(req.ProcessID) + "",
	}

	var resp Process

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForProcessRequest is used by WaitForProcess method.
type WaitForProcessRequest struct {
	ProcessID     string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForProcess waits for the Process to reach a terminal state.
func (s *API) WaitForProcess(req *WaitForProcessRequest, opts ...scw.RequestOption) (*Process, error) {
	timeout := defaultQaasTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultQaasRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[ProcessStatus]struct{}{
		ProcessStatusStarting:   {},
		ProcessStatusRunning:    {},
		ProcessStatusCancelling: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetProcess(&GetProcessRequest{
				ProcessID: req.ProcessID,
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
		return nil, errors.Wrap(err, "waiting for Process failed")
	}

	return res.(*Process), nil
}

// ListProcesses: Retrieve information about all processes.
func (s *API) ListProcesses(req *ListProcessesRequest, opts ...scw.RequestOption) (*ListProcessesResponse, error) {
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
	parameter.AddToQuery(query, "application_id", req.ApplicationID)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/processes",
		Query:  query,
	}

	var resp ListProcessesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateProcess: Update process information of the provided **process ID**.
func (s *API) UpdateProcess(req *UpdateProcessRequest, opts ...scw.RequestOption) (*Process, error) {
	var err error

	if fmt.Sprint(req.ProcessID) == "" {
		return nil, errors.New("field ProcessID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/qaas/v1alpha1/processes/" + fmt.Sprint(req.ProcessID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Process

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelProcess: Cancel a process by its unique ID. Intermediate results are still available.
func (s *API) CancelProcess(req *CancelProcessRequest, opts ...scw.RequestOption) (*Process, error) {
	var err error

	if fmt.Sprint(req.ProcessID) == "" {
		return nil, errors.New("field ProcessID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/qaas/v1alpha1/processes/" + fmt.Sprint(req.ProcessID) + "/cancel",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Process

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteProcess: Delete a process by its unique ID and delete all its data.
func (s *API) DeleteProcess(req *DeleteProcessRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.ProcessID) == "" {
		return errors.New("field ProcessID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/qaas/v1alpha1/processes/" + fmt.Sprint(req.ProcessID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListProcessResults: Retrieve all intermediate and final result of a process.
func (s *API) ListProcessResults(req *ListProcessResultsRequest, opts ...scw.RequestOption) (*ListProcessResultsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.ProcessID) == "" {
		return nil, errors.New("field ProcessID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/processes/" + fmt.Sprint(req.ProcessID) + "/results",
		Query:  query,
	}

	var resp ListProcessResultsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetApplication: Retrieve information about the provided **application ID**, such as name, type and compatible platforms.
func (s *API) GetApplication(req *GetApplicationRequest, opts ...scw.RequestOption) (*Application, error) {
	var err error

	if fmt.Sprint(req.ApplicationID) == "" {
		return nil, errors.New("field ApplicationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/applications/" + fmt.Sprint(req.ApplicationID) + "",
	}

	var resp Application

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListApplications: Retrieve information about all applications.
func (s *API) ListApplications(req *ListApplicationsRequest, opts ...scw.RequestOption) (*ListApplicationsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "application_type", req.ApplicationType)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/applications",
		Query:  query,
	}

	var resp ListApplicationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBooking: Retrieve information about the provided **booking ID**, such as description, status and progress message.
func (s *API) GetBooking(req *GetBookingRequest, opts ...scw.RequestOption) (*Booking, error) {
	var err error

	if fmt.Sprint(req.BookingID) == "" {
		return nil, errors.New("field BookingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/bookings/" + fmt.Sprint(req.BookingID) + "",
	}

	var resp Booking

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForBookingRequest is used by WaitForBooking method.
type WaitForBookingRequest struct {
	BookingID     string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForBooking waits for the Booking to reach a terminal state.
func (s *API) WaitForBooking(req *WaitForBookingRequest, opts ...scw.RequestOption) (*Booking, error) {
	timeout := defaultQaasTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultQaasRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[BookingStatus]struct{}{
		BookingStatusWaiting:    {},
		BookingStatusValidating: {},
		BookingStatusCancelling: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetBooking(&GetBookingRequest{
				BookingID: req.BookingID,
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
		return nil, errors.Wrap(err, "waiting for Booking failed")
	}

	return res.(*Booking), nil
}

// ListBookings: Retrieve information about all bookings of the provided **project ID** or ** platform ID**.
func (s *API) ListBookings(req *ListBookingsRequest, opts ...scw.RequestOption) (*ListBookingsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "platform_id", req.PlatformID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/bookings",
		Query:  query,
	}

	var resp ListBookingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBooking: Update booking information of the provided **booking ID**.
func (s *API) UpdateBooking(req *UpdateBookingRequest, opts ...scw.RequestOption) (*Booking, error) {
	var err error

	if fmt.Sprint(req.BookingID) == "" {
		return nil, errors.New("field BookingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/qaas/v1alpha1/bookings/" + fmt.Sprint(req.BookingID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Booking

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateModel: Create and register a new model that can be executed through next jobs. A model can also be assigned to a Session.
func (s *API) CreateModel(req *CreateModelRequest, opts ...scw.RequestOption) (*Model, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/qaas/v1alpha1/models",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Model

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetModel: Retrieve information about of the provided **model ID**.
func (s *API) GetModel(req *GetModelRequest, opts ...scw.RequestOption) (*Model, error) {
	var err error

	if fmt.Sprint(req.ModelID) == "" {
		return nil, errors.New("field ModelID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/models/" + fmt.Sprint(req.ModelID) + "",
	}

	var resp Model

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListModels: Retrieve information about all models of the provided **project ID**.
func (s *API) ListModels(req *ListModelsRequest, opts ...scw.RequestOption) (*ListModelsResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/qaas/v1alpha1/models",
		Query:  query,
	}

	var resp ListModelsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package edge_services provides methods and message types of the edge_services v1beta1 API.
package edge_services

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
	defaultEdgeServicesRetryInterval = 15 * time.Second
	defaultEdgeServicesTimeout       = 5 * time.Minute
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

type DNSStageType string

const (
	// DNS stage type unknown (default).
	DNSStageTypeUnknownType = DNSStageType("unknown_type")
	// The DNS stage is configured with the default FQDN.
	DNSStageTypeAuto = DNSStageType("auto")
	// Custom FQDN has been provided, managed by Scaleway Domains and DNS.
	DNSStageTypeManaged = DNSStageType("managed")
	// Custom FQDN has been provided, not managed by Scaleway Domains and DNS.
	DNSStageTypeCustom = DNSStageType("custom")
)

func (enum DNSStageType) String() string {
	if enum == "" {
		// return default value if empty
		return string(DNSStageTypeUnknownType)
	}
	return string(enum)
}

func (enum DNSStageType) Values() []DNSStageType {
	return []DNSStageType{
		"unknown_type",
		"auto",
		"managed",
		"custom",
	}
}

func (enum DNSStageType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DNSStageType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DNSStageType(DNSStageType(tmp).String())
	return nil
}

type LBOriginError string

const (
	LBOriginErrorUnknown           = LBOriginError("unknown")
	LBOriginErrorTimeout           = LBOriginError("timeout")
	LBOriginErrorConnectionRefused = LBOriginError("connection_refused")
	LBOriginErrorTLSError          = LBOriginError("tls_error")
)

func (enum LBOriginError) String() string {
	if enum == "" {
		// return default value if empty
		return string(LBOriginErrorUnknown)
	}
	return string(enum)
}

func (enum LBOriginError) Values() []LBOriginError {
	return []LBOriginError{
		"unknown",
		"timeout",
		"connection_refused",
		"tls_error",
	}
}

func (enum LBOriginError) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LBOriginError) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LBOriginError(LBOriginError(tmp).String())
	return nil
}

type ListBackendStagesRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListBackendStagesRequestOrderByCreatedAtAsc = ListBackendStagesRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListBackendStagesRequestOrderByCreatedAtDesc = ListBackendStagesRequestOrderBy("created_at_desc")
)

func (enum ListBackendStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListBackendStagesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListBackendStagesRequestOrderBy) Values() []ListBackendStagesRequestOrderBy {
	return []ListBackendStagesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListBackendStagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListBackendStagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListBackendStagesRequestOrderBy(ListBackendStagesRequestOrderBy(tmp).String())
	return nil
}

type ListCacheStagesRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListCacheStagesRequestOrderByCreatedAtAsc = ListCacheStagesRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListCacheStagesRequestOrderByCreatedAtDesc = ListCacheStagesRequestOrderBy("created_at_desc")
)

func (enum ListCacheStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListCacheStagesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListCacheStagesRequestOrderBy) Values() []ListCacheStagesRequestOrderBy {
	return []ListCacheStagesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListCacheStagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCacheStagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCacheStagesRequestOrderBy(ListCacheStagesRequestOrderBy(tmp).String())
	return nil
}

type ListDNSStagesRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListDNSStagesRequestOrderByCreatedAtAsc = ListDNSStagesRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListDNSStagesRequestOrderByCreatedAtDesc = ListDNSStagesRequestOrderBy("created_at_desc")
)

func (enum ListDNSStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListDNSStagesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListDNSStagesRequestOrderBy) Values() []ListDNSStagesRequestOrderBy {
	return []ListDNSStagesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListDNSStagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDNSStagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDNSStagesRequestOrderBy(ListDNSStagesRequestOrderBy(tmp).String())
	return nil
}

type ListPipelinesRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListPipelinesRequestOrderByCreatedAtAsc = ListPipelinesRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListPipelinesRequestOrderByCreatedAtDesc = ListPipelinesRequestOrderBy("created_at_desc")
	// Order by name (ascending alphabetical order).
	ListPipelinesRequestOrderByNameAsc = ListPipelinesRequestOrderBy("name_asc")
	// Order by name (descending alphabetical order).
	ListPipelinesRequestOrderByNameDesc = ListPipelinesRequestOrderBy("name_desc")
)

func (enum ListPipelinesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPipelinesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListPipelinesRequestOrderBy) Values() []ListPipelinesRequestOrderBy {
	return []ListPipelinesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListPipelinesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPipelinesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPipelinesRequestOrderBy(ListPipelinesRequestOrderBy(tmp).String())
	return nil
}

type ListPipelinesWithStagesRequestOrderBy string

const (
	ListPipelinesWithStagesRequestOrderByCreatedAtAsc  = ListPipelinesWithStagesRequestOrderBy("created_at_asc")
	ListPipelinesWithStagesRequestOrderByCreatedAtDesc = ListPipelinesWithStagesRequestOrderBy("created_at_desc")
	ListPipelinesWithStagesRequestOrderByNameAsc       = ListPipelinesWithStagesRequestOrderBy("name_asc")
	ListPipelinesWithStagesRequestOrderByNameDesc      = ListPipelinesWithStagesRequestOrderBy("name_desc")
)

func (enum ListPipelinesWithStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPipelinesWithStagesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListPipelinesWithStagesRequestOrderBy) Values() []ListPipelinesWithStagesRequestOrderBy {
	return []ListPipelinesWithStagesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListPipelinesWithStagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPipelinesWithStagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPipelinesWithStagesRequestOrderBy(ListPipelinesWithStagesRequestOrderBy(tmp).String())
	return nil
}

type ListPurgeRequestsRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListPurgeRequestsRequestOrderByCreatedAtAsc = ListPurgeRequestsRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListPurgeRequestsRequestOrderByCreatedAtDesc = ListPurgeRequestsRequestOrderBy("created_at_desc")
)

func (enum ListPurgeRequestsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPurgeRequestsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListPurgeRequestsRequestOrderBy) Values() []ListPurgeRequestsRequestOrderBy {
	return []ListPurgeRequestsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListPurgeRequestsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPurgeRequestsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPurgeRequestsRequestOrderBy(ListPurgeRequestsRequestOrderBy(tmp).String())
	return nil
}

type ListRouteStagesRequestOrderBy string

const (
	ListRouteStagesRequestOrderByCreatedAtAsc  = ListRouteStagesRequestOrderBy("created_at_asc")
	ListRouteStagesRequestOrderByCreatedAtDesc = ListRouteStagesRequestOrderBy("created_at_desc")
)

func (enum ListRouteStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRouteStagesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRouteStagesRequestOrderBy) Values() []ListRouteStagesRequestOrderBy {
	return []ListRouteStagesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRouteStagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRouteStagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRouteStagesRequestOrderBy(ListRouteStagesRequestOrderBy(tmp).String())
	return nil
}

type ListTLSStagesRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListTLSStagesRequestOrderByCreatedAtAsc = ListTLSStagesRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListTLSStagesRequestOrderByCreatedAtDesc = ListTLSStagesRequestOrderBy("created_at_desc")
)

func (enum ListTLSStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListTLSStagesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListTLSStagesRequestOrderBy) Values() []ListTLSStagesRequestOrderBy {
	return []ListTLSStagesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListTLSStagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTLSStagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTLSStagesRequestOrderBy(ListTLSStagesRequestOrderBy(tmp).String())
	return nil
}

type ListWafStagesRequestOrderBy string

const (
	ListWafStagesRequestOrderByCreatedAtAsc  = ListWafStagesRequestOrderBy("created_at_asc")
	ListWafStagesRequestOrderByCreatedAtDesc = ListWafStagesRequestOrderBy("created_at_desc")
)

func (enum ListWafStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListWafStagesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListWafStagesRequestOrderBy) Values() []ListWafStagesRequestOrderBy {
	return []ListWafStagesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListWafStagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListWafStagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListWafStagesRequestOrderBy(ListWafStagesRequestOrderBy(tmp).String())
	return nil
}

type PipelineErrorCode string

const (
	PipelineErrorCodeUnknownCode               = PipelineErrorCode("unknown_code")
	PipelineErrorCodeDNSInvalidFormat          = PipelineErrorCode("dns_invalid_format")
	PipelineErrorCodeDNSInvalidTld             = PipelineErrorCode("dns_invalid_tld")
	PipelineErrorCodeDNSForbiddenRootDomain    = PipelineErrorCode("dns_forbidden_root_domain")
	PipelineErrorCodeDNSForbiddenScwCloud      = PipelineErrorCode("dns_forbidden_scw_cloud")
	PipelineErrorCodeDNSDomainDontExist        = PipelineErrorCode("dns_domain_dont_exist")
	PipelineErrorCodeDNSCnameDontExist         = PipelineErrorCode("dns_cname_dont_exist")
	PipelineErrorCodeDNSCnameResolve           = PipelineErrorCode("dns_cname_resolve")
	PipelineErrorCodeDNSFqdnAlreadyExists      = PipelineErrorCode("dns_fqdn_already_exists")
	PipelineErrorCodeDNSFqdnAlreadyInUse       = PipelineErrorCode("dns_fqdn_already_in_use")
	PipelineErrorCodeTLSCertDeleted            = PipelineErrorCode("tls_cert_deleted")
	PipelineErrorCodeTLSCertDisabled           = PipelineErrorCode("tls_cert_disabled")
	PipelineErrorCodeTLSCertExpired            = PipelineErrorCode("tls_cert_expired")
	PipelineErrorCodeTLSCertInvalidFormat      = PipelineErrorCode("tls_cert_invalid_format")
	PipelineErrorCodeTLSCertMissing            = PipelineErrorCode("tls_cert_missing")
	PipelineErrorCodeTLSChainOrder             = PipelineErrorCode("tls_chain_order")
	PipelineErrorCodeTLSKeyInvalidFormat       = PipelineErrorCode("tls_key_invalid_format")
	PipelineErrorCodeTLSKeyMissing             = PipelineErrorCode("tls_key_missing")
	PipelineErrorCodeTLSKeyTooMany             = PipelineErrorCode("tls_key_too_many")
	PipelineErrorCodeTLSManagedDomainRateLimit = PipelineErrorCode("tls_managed_domain_rate_limit")
	PipelineErrorCodeTLSManagedInternal        = PipelineErrorCode("tls_managed_internal")
	PipelineErrorCodeTLSPairMismatch           = PipelineErrorCode("tls_pair_mismatch")
	PipelineErrorCodeTLSRootInconsistent       = PipelineErrorCode("tls_root_inconsistent")
	PipelineErrorCodeTLSRootIncorrect          = PipelineErrorCode("tls_root_incorrect")
	PipelineErrorCodeTLSRootMissing            = PipelineErrorCode("tls_root_missing")
	PipelineErrorCodeTLSSanMismatch            = PipelineErrorCode("tls_san_mismatch")
	PipelineErrorCodeTLSSelfSigned             = PipelineErrorCode("tls_self_signed")
	PipelineErrorCodeTLSCaaMalfunction         = PipelineErrorCode("tls_caa_malfunction")
	PipelineErrorCodePipelineInvalidWorkflow   = PipelineErrorCode("pipeline_invalid_workflow")
	PipelineErrorCodePipelineMissingHeadStage  = PipelineErrorCode("pipeline_missing_head_stage")
)

func (enum PipelineErrorCode) String() string {
	if enum == "" {
		// return default value if empty
		return string(PipelineErrorCodeUnknownCode)
	}
	return string(enum)
}

func (enum PipelineErrorCode) Values() []PipelineErrorCode {
	return []PipelineErrorCode{
		"unknown_code",
		"dns_invalid_format",
		"dns_invalid_tld",
		"dns_forbidden_root_domain",
		"dns_forbidden_scw_cloud",
		"dns_domain_dont_exist",
		"dns_cname_dont_exist",
		"dns_cname_resolve",
		"dns_fqdn_already_exists",
		"dns_fqdn_already_in_use",
		"tls_cert_deleted",
		"tls_cert_disabled",
		"tls_cert_expired",
		"tls_cert_invalid_format",
		"tls_cert_missing",
		"tls_chain_order",
		"tls_key_invalid_format",
		"tls_key_missing",
		"tls_key_too_many",
		"tls_managed_domain_rate_limit",
		"tls_managed_internal",
		"tls_pair_mismatch",
		"tls_root_inconsistent",
		"tls_root_incorrect",
		"tls_root_missing",
		"tls_san_mismatch",
		"tls_self_signed",
		"tls_caa_malfunction",
		"pipeline_invalid_workflow",
		"pipeline_missing_head_stage",
	}
}

func (enum PipelineErrorCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PipelineErrorCode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PipelineErrorCode(PipelineErrorCode(tmp).String())
	return nil
}

type PipelineErrorSeverity string

const (
	PipelineErrorSeverityUnknownSeverity = PipelineErrorSeverity("unknown_severity")
	PipelineErrorSeverityWarning         = PipelineErrorSeverity("warning")
	PipelineErrorSeverityCritical        = PipelineErrorSeverity("critical")
)

func (enum PipelineErrorSeverity) String() string {
	if enum == "" {
		// return default value if empty
		return string(PipelineErrorSeverityUnknownSeverity)
	}
	return string(enum)
}

func (enum PipelineErrorSeverity) Values() []PipelineErrorSeverity {
	return []PipelineErrorSeverity{
		"unknown_severity",
		"warning",
		"critical",
	}
}

func (enum PipelineErrorSeverity) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PipelineErrorSeverity) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PipelineErrorSeverity(PipelineErrorSeverity(tmp).String())
	return nil
}

type PipelineErrorStage string

const (
	PipelineErrorStageUnknownStage = PipelineErrorStage("unknown_stage")
	PipelineErrorStageDNS          = PipelineErrorStage("dns")
	PipelineErrorStageTLS          = PipelineErrorStage("tls")
	PipelineErrorStageCache        = PipelineErrorStage("cache")
	PipelineErrorStageBackend      = PipelineErrorStage("backend")
)

func (enum PipelineErrorStage) String() string {
	if enum == "" {
		// return default value if empty
		return string(PipelineErrorStageUnknownStage)
	}
	return string(enum)
}

func (enum PipelineErrorStage) Values() []PipelineErrorStage {
	return []PipelineErrorStage{
		"unknown_stage",
		"dns",
		"tls",
		"cache",
		"backend",
	}
}

func (enum PipelineErrorStage) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PipelineErrorStage) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PipelineErrorStage(PipelineErrorStage(tmp).String())
	return nil
}

type PipelineErrorType string

const (
	PipelineErrorTypeUnknownType = PipelineErrorType("unknown_type")
	PipelineErrorTypeRuntime     = PipelineErrorType("runtime")
	PipelineErrorTypeConfig      = PipelineErrorType("config")
)

func (enum PipelineErrorType) String() string {
	if enum == "" {
		// return default value if empty
		return string(PipelineErrorTypeUnknownType)
	}
	return string(enum)
}

func (enum PipelineErrorType) Values() []PipelineErrorType {
	return []PipelineErrorType{
		"unknown_type",
		"runtime",
		"config",
	}
}

func (enum PipelineErrorType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PipelineErrorType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PipelineErrorType(PipelineErrorType(tmp).String())
	return nil
}

type PipelineStatus string

const (
	// Pipeline status unknown (default).
	PipelineStatusUnknownStatus = PipelineStatus("unknown_status")
	// The pipeline has been created and processed.
	PipelineStatusReady = PipelineStatus("ready")
	// An error occurred during the pipeline creation.
	PipelineStatusError = PipelineStatus("error")
	// The pipeline status is pending.
	PipelineStatusPending = PipelineStatus("pending")
	// An event occurred and the pipeline may not work.
	PipelineStatusWarning = PipelineStatus("warning")
	PipelineStatusLocked  = PipelineStatus("locked")
)

func (enum PipelineStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(PipelineStatusUnknownStatus)
	}
	return string(enum)
}

func (enum PipelineStatus) Values() []PipelineStatus {
	return []PipelineStatus{
		"unknown_status",
		"ready",
		"error",
		"pending",
		"warning",
		"locked",
	}
}

func (enum PipelineStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PipelineStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PipelineStatus(PipelineStatus(tmp).String())
	return nil
}

type PlanName string

const (
	PlanNameUnknownName  = PlanName("unknown_name")
	PlanNameStarter      = PlanName("starter")
	PlanNameProfessional = PlanName("professional")
	PlanNameAdvanced     = PlanName("advanced")
)

func (enum PlanName) String() string {
	if enum == "" {
		// return default value if empty
		return string(PlanNameUnknownName)
	}
	return string(enum)
}

func (enum PlanName) Values() []PlanName {
	return []PlanName{
		"unknown_name",
		"starter",
		"professional",
		"advanced",
	}
}

func (enum PlanName) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlanName) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlanName(PlanName(tmp).String())
	return nil
}

type PurgeRequestStatus string

const (
	// Purge request status unknown (default).
	PurgeRequestStatusUnknownStatus = PurgeRequestStatus("unknown_status")
	// The purge request has been processed successfully.
	PurgeRequestStatusDone = PurgeRequestStatus("done")
	// An error occurred during the purge request.
	PurgeRequestStatusError = PurgeRequestStatus("error")
	// The purge request status is pending.
	PurgeRequestStatusPending = PurgeRequestStatus("pending")
)

func (enum PurgeRequestStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(PurgeRequestStatusUnknownStatus)
	}
	return string(enum)
}

func (enum PurgeRequestStatus) Values() []PurgeRequestStatus {
	return []PurgeRequestStatus{
		"unknown_status",
		"done",
		"error",
		"pending",
	}
}

func (enum PurgeRequestStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PurgeRequestStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PurgeRequestStatus(PurgeRequestStatus(tmp).String())
	return nil
}

type RuleHTTPMatchMethodFilter string

const (
	// Method filter unknown (default).
	RuleHTTPMatchMethodFilterUnknownMethodFilter = RuleHTTPMatchMethodFilter("unknown_method_filter")
	// GET HTTP method.
	RuleHTTPMatchMethodFilterGet = RuleHTTPMatchMethodFilter("get")
	// POST HTTP method.
	RuleHTTPMatchMethodFilterPost = RuleHTTPMatchMethodFilter("post")
	// PUT HTTP method.
	RuleHTTPMatchMethodFilterPut = RuleHTTPMatchMethodFilter("put")
	// PATCH HTTP method.
	RuleHTTPMatchMethodFilterPatch = RuleHTTPMatchMethodFilter("patch")
	// DELETE HTTP method.
	RuleHTTPMatchMethodFilterDelete = RuleHTTPMatchMethodFilter("delete")
	// HEAD HTTP method.
	RuleHTTPMatchMethodFilterHead = RuleHTTPMatchMethodFilter("head")
	// OPTIONS HTTP method.
	RuleHTTPMatchMethodFilterOptions = RuleHTTPMatchMethodFilter("options")
)

func (enum RuleHTTPMatchMethodFilter) String() string {
	if enum == "" {
		// return default value if empty
		return string(RuleHTTPMatchMethodFilterUnknownMethodFilter)
	}
	return string(enum)
}

func (enum RuleHTTPMatchMethodFilter) Values() []RuleHTTPMatchMethodFilter {
	return []RuleHTTPMatchMethodFilter{
		"unknown_method_filter",
		"get",
		"post",
		"put",
		"patch",
		"delete",
		"head",
		"options",
	}
}

func (enum RuleHTTPMatchMethodFilter) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RuleHTTPMatchMethodFilter) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RuleHTTPMatchMethodFilter(RuleHTTPMatchMethodFilter(tmp).String())
	return nil
}

type RuleHTTPMatchPathFilterPathFilterType string

const (
	// Path filter type unknown (default).
	RuleHTTPMatchPathFilterPathFilterTypeUnknownPathFilter = RuleHTTPMatchPathFilterPathFilterType("unknown_path_filter")
	// Value of path to filter for is evaluated as regex.
	RuleHTTPMatchPathFilterPathFilterTypeRegex = RuleHTTPMatchPathFilterPathFilterType("regex")
)

func (enum RuleHTTPMatchPathFilterPathFilterType) String() string {
	if enum == "" {
		// return default value if empty
		return string(RuleHTTPMatchPathFilterPathFilterTypeUnknownPathFilter)
	}
	return string(enum)
}

func (enum RuleHTTPMatchPathFilterPathFilterType) Values() []RuleHTTPMatchPathFilterPathFilterType {
	return []RuleHTTPMatchPathFilterPathFilterType{
		"unknown_path_filter",
		"regex",
	}
}

func (enum RuleHTTPMatchPathFilterPathFilterType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RuleHTTPMatchPathFilterPathFilterType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RuleHTTPMatchPathFilterPathFilterType(RuleHTTPMatchPathFilterPathFilterType(tmp).String())
	return nil
}

type SearchBackendStagesRequestOrderBy string

const (
	SearchBackendStagesRequestOrderByCreatedAtAsc  = SearchBackendStagesRequestOrderBy("created_at_asc")
	SearchBackendStagesRequestOrderByCreatedAtDesc = SearchBackendStagesRequestOrderBy("created_at_desc")
)

func (enum SearchBackendStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(SearchBackendStagesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum SearchBackendStagesRequestOrderBy) Values() []SearchBackendStagesRequestOrderBy {
	return []SearchBackendStagesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum SearchBackendStagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SearchBackendStagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SearchBackendStagesRequestOrderBy(SearchBackendStagesRequestOrderBy(tmp).String())
	return nil
}

type SearchRouteRulesRequestOrderBy string

const (
	SearchRouteRulesRequestOrderByCreatedAtAsc  = SearchRouteRulesRequestOrderBy("created_at_asc")
	SearchRouteRulesRequestOrderByCreatedAtDesc = SearchRouteRulesRequestOrderBy("created_at_desc")
)

func (enum SearchRouteRulesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(SearchRouteRulesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum SearchRouteRulesRequestOrderBy) Values() []SearchRouteRulesRequestOrderBy {
	return []SearchRouteRulesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum SearchRouteRulesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SearchRouteRulesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SearchRouteRulesRequestOrderBy(SearchRouteRulesRequestOrderBy(tmp).String())
	return nil
}

type SearchWafStagesRequestOrderBy string

const (
	SearchWafStagesRequestOrderByCreatedAtAsc  = SearchWafStagesRequestOrderBy("created_at_asc")
	SearchWafStagesRequestOrderByCreatedAtDesc = SearchWafStagesRequestOrderBy("created_at_desc")
)

func (enum SearchWafStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(SearchWafStagesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum SearchWafStagesRequestOrderBy) Values() []SearchWafStagesRequestOrderBy {
	return []SearchWafStagesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum SearchWafStagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SearchWafStagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SearchWafStagesRequestOrderBy(SearchWafStagesRequestOrderBy(tmp).String())
	return nil
}

type WafStageMode string

const (
	// WAF mode unknown (default).
	WafStageModeUnknownMode = WafStageMode("unknown_mode")
	// WAF engine is disabled and does not evaluate requests.
	WafStageModeDisable = WafStageMode("disable")
	// WAF engine is enabled and evaluates requests but does not block them. HTTP requests which are identified as malicious, and would have been blocked in `enable` mode are logged in Scaleway Cockpit.
	WafStageModeLogOnly = WafStageMode("log_only")
	// WAF engine is enabled and evaluates requests. HTTP requests classed as malicious according to the defined `paranoia_level` are blocked and logged in Scaleway Cockpit.
	WafStageModeEnable = WafStageMode("enable")
)

func (enum WafStageMode) String() string {
	if enum == "" {
		// return default value if empty
		return string(WafStageModeUnknownMode)
	}
	return string(enum)
}

func (enum WafStageMode) Values() []WafStageMode {
	return []WafStageMode{
		"unknown_mode",
		"disable",
		"log_only",
		"enable",
	}
}

func (enum WafStageMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *WafStageMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = WafStageMode(WafStageMode(tmp).String())
	return nil
}

// ScalewayLB: scaleway lb.
type ScalewayLB struct {
	// ID: ID of the Load Balancer.
	ID string `json:"id"`

	// Zone: zone of the Load Balancer.
	Zone scw.Zone `json:"zone"`

	// FrontendID: ID of the frontend linked to the Load Balancer.
	FrontendID string `json:"frontend_id"`

	// IsSsl: defines whether the Load Balancer's frontend handles SSL connections.
	IsSsl *bool `json:"is_ssl"`

	// DomainName: fully Qualified Domain Name (in the format subdomain.example.com) to use in HTTP requests sent towards your Load Balancer.
	DomainName *string `json:"domain_name"`

	// HasWebsocket: defines whether to forward websocket requests to the load balancer.
	HasWebsocket *bool `json:"has_websocket"`
}

// RuleHTTPMatchPathFilter: rule http match path filter.
type RuleHTTPMatchPathFilter struct {
	// PathFilterType: type of filter to match for the HTTP URL path. For now, all path filters must be written in regex and use the `regex` type.
	// Default value: unknown_path_filter
	PathFilterType RuleHTTPMatchPathFilterPathFilterType `json:"path_filter_type"`

	// Value: value to be matched for the HTTP URL path.
	Value string `json:"value"`
}

// ScalewayLBBackendConfig: scaleway lb backend config.
type ScalewayLBBackendConfig struct {
	// LBs: load Balancer information.
	LBs []*ScalewayLB `json:"lbs"`
}

// ScalewayS3BackendConfig: scaleway s3 backend config.
type ScalewayS3BackendConfig struct {
	// BucketName: name of the Bucket.
	BucketName *string `json:"bucket_name"`

	// BucketRegion: region of the Bucket.
	BucketRegion *string `json:"bucket_region"`

	// IsWebsite: defines whether the bucket website feature is enabled.
	IsWebsite *bool `json:"is_website"`
}

// ScalewayServerlessContainerBackendConfig: scaleway serverless container backend config.
type ScalewayServerlessContainerBackendConfig struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`

	ContainerID string `json:"container_id"`
}

// ScalewayServerlessFunctionBackendConfig: scaleway serverless function backend config.
type ScalewayServerlessFunctionBackendConfig struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`

	FunctionID string `json:"function_id"`
}

// PipelineError: pipeline error.
type PipelineError struct {
	// Stage: default value: unknown_stage
	Stage PipelineErrorStage `json:"stage"`

	// Code: default value: unknown_code
	Code PipelineErrorCode `json:"code"`

	// Severity: default value: unknown_severity
	Severity PipelineErrorSeverity `json:"severity"`

	Message string `json:"message"`

	// Type: default value: unknown_type
	Type PipelineErrorType `json:"type"`
}

// TLSSecret: tls secret.
type TLSSecret struct {
	// SecretID: ID of the Secret.
	SecretID string `json:"secret_id"`

	// Region: region of the Secret.
	Region scw.Region `json:"region"`
}

// RuleHTTPMatch: rule http match.
type RuleHTTPMatch struct {
	// MethodFilters: HTTP methods to filter for. A request using any of these methods will be considered to match the rule. Possible values are `get`, `post`, `put`, `patch`, `delete`, `head`, `options`. All methods will match if none is provided.
	MethodFilters []RuleHTTPMatchMethodFilter `json:"method_filters"`

	// PathFilter: HTTP URL path to filter for. A request whose path matches the given filter will be considered to match the rule. All paths will match if none is provided.
	PathFilter *RuleHTTPMatchPathFilter `json:"path_filter"`
}

// BackendStage: backend stage.
type BackendStage struct {
	// ID: ID of the backend stage.
	ID string `json:"id"`

	// PipelineID: pipeline ID the backend stage belongs to.
	PipelineID string `json:"pipeline_id"`

	// CreatedAt: date the backend stage was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the backend stage was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// ScalewayS3: scaleway Object Storage origin bucket (S3) linked to the backend stage.
	// Precisely one of ScalewayS3, ScalewayLB, ScalewayServerlessContainer, ScalewayServerlessFunction must be set.
	ScalewayS3 *ScalewayS3BackendConfig `json:"scaleway_s3,omitempty"`

	// ScalewayLB: scaleway Load Balancer origin linked to the backend stage.
	// Precisely one of ScalewayS3, ScalewayLB, ScalewayServerlessContainer, ScalewayServerlessFunction must be set.
	ScalewayLB *ScalewayLBBackendConfig `json:"scaleway_lb,omitempty"`

	// Precisely one of ScalewayS3, ScalewayLB, ScalewayServerlessContainer, ScalewayServerlessFunction must be set.
	ScalewayServerlessContainer *ScalewayServerlessContainerBackendConfig `json:"scaleway_serverless_container,omitempty"`

	// Precisely one of ScalewayS3, ScalewayLB, ScalewayServerlessContainer, ScalewayServerlessFunction must be set.
	ScalewayServerlessFunction *ScalewayServerlessFunctionBackendConfig `json:"scaleway_serverless_function,omitempty"`
}

// CacheStage: cache stage.
type CacheStage struct {
	// ID: ID of the cache stage.
	ID string `json:"id"`

	// PipelineID: pipeline ID the cache stage belongs to.
	PipelineID string `json:"pipeline_id"`

	// FallbackTTL: time To Live (TTL) in seconds. Defines how long content is cached.
	FallbackTTL *scw.Duration `json:"fallback_ttl"`

	// IncludeCookies: defines whether responses to requests with cookies must be stored in the cache.
	IncludeCookies bool `json:"include_cookies"`

	// CreatedAt: date the cache stage was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the cache stage was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// BackendStageID: backend stage ID the cache stage is linked to.
	// Precisely one of BackendStageID, WafStageID, RouteStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`

	// Precisely one of BackendStageID, WafStageID, RouteStageID must be set.
	WafStageID *string `json:"waf_stage_id,omitempty"`

	// Precisely one of BackendStageID, WafStageID, RouteStageID must be set.
	RouteStageID *string `json:"route_stage_id,omitempty"`
}

// DNSStage: dns stage.
type DNSStage struct {
	// ID: ID of the DNS stage.
	ID string `json:"id"`

	// DefaultFqdn: default Fully Qualified Domain Name attached to the stage.
	DefaultFqdn string `json:"default_fqdn"`

	// Fqdns: list of additional (custom) Fully Qualified Domain Names attached to the stage.
	Fqdns []string `json:"fqdns"`

	// Type: type of the stage.
	// Default value: unknown_type
	Type DNSStageType `json:"type"`

	// PipelineID: pipeline ID the DNS stage belongs to.
	PipelineID string `json:"pipeline_id"`

	// CreatedAt: date the DNS stage was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the DNS stage was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// TLSStageID: TLS stage ID the DNS stage is linked to.
	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	TLSStageID *string `json:"tls_stage_id,omitempty"`

	// CacheStageID: cache stage ID the DNS stage is linked to.
	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// BackendStageID: backend stage ID the DNS stage is linked to.
	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// Pipeline: pipeline.
type Pipeline struct {
	// ID: ID of the pipeline.
	ID string `json:"id"`

	// Name: name of the pipeline.
	Name string `json:"name"`

	// Description: description of the pipeline.
	Description string `json:"description"`

	// Status: status of the pipeline.
	// Default value: unknown_status
	Status PipelineStatus `json:"status"`

	// Errors: errors of the pipeline.
	Errors []*PipelineError `json:"errors"`

	// ProjectID: project ID of the pipeline.
	ProjectID string `json:"project_id"`

	// OrganizationID: organization ID of the pipeline.
	OrganizationID string `json:"organization_id"`

	// CreatedAt: date the pipeline was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the pipeline was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
}

// RouteStage: route stage.
type RouteStage struct {
	// ID: ID of the route stage.
	ID string `json:"id"`

	// PipelineID: pipeline ID the route stage belongs to.
	PipelineID string `json:"pipeline_id"`

	// WafStageID: ID of the WAF stage HTTP requests should be forwarded to when no rules are matched.
	// Precisely one of WafStageID must be set.
	WafStageID *string `json:"waf_stage_id,omitempty"`

	// CreatedAt: date the route stage was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the route stage was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
}

// TLSStage: tls stage.
type TLSStage struct {
	// ID: ID of the TLS stage.
	ID string `json:"id"`

	// Secrets: secret (from Scaleway Secret Manager) containing your custom certificate.
	Secrets []*TLSSecret `json:"secrets"`

	// ManagedCertificate: true when Scaleway generates and manages a Let's Encrypt certificate for the TLS stage/custom endpoint.
	ManagedCertificate bool `json:"managed_certificate"`

	// PipelineID: pipeline ID the TLS stage belongs to.
	PipelineID string `json:"pipeline_id"`

	// CertificateExpiresAt: expiration date of the certificate.
	CertificateExpiresAt *time.Time `json:"certificate_expires_at"`

	// CreatedAt: date the TLS stage was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the TLS stage was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// CacheStageID: cache stage ID the TLS stage is linked to.
	// Precisely one of CacheStageID, BackendStageID, WafStageID, RouteStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// BackendStageID: backend stage ID the TLS stage is linked to.
	// Precisely one of CacheStageID, BackendStageID, WafStageID, RouteStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`

	// Precisely one of CacheStageID, BackendStageID, WafStageID, RouteStageID must be set.
	WafStageID *string `json:"waf_stage_id,omitempty"`

	// Precisely one of CacheStageID, BackendStageID, WafStageID, RouteStageID must be set.
	RouteStageID *string `json:"route_stage_id,omitempty"`
}

// WafStage: waf stage.
type WafStage struct {
	// ID: ID of the WAF stage.
	ID string `json:"id"`

	// PipelineID: pipeline ID the WAF stage belongs to.
	PipelineID string `json:"pipeline_id"`

	// Mode: mode defining WAF behavior (`disable`/`log_only`/`enable`).
	// Default value: unknown_mode
	Mode WafStageMode `json:"mode"`

	// ParanoiaLevel: sensitivity level (`1`,`2`,`3`,`4`) to use when classifying requests as malicious. With a high level, requests are more likely to be classed as malicious, and false positives are expected. With a lower level, requests are more likely to be classed as benign.
	ParanoiaLevel uint32 `json:"paranoia_level"`

	// CreatedAt: date the WAF stage was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the WAF stage was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// BackendStageID: ID of the backend stage to forward requests to after the WAF stage.
	// Precisely one of BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// SetRouteRulesRequestRouteRule: set route rules request route rule.
type SetRouteRulesRequestRouteRule struct {
	// RuleHTTPMatch: rule condition to be matched. Requests matching the condition defined here will be directly forwarded to the backend specified by the `backend_stage_id` field. Requests that do not match will be checked by the next rule's condition.
	// Precisely one of RuleHTTPMatch must be set.
	RuleHTTPMatch *RuleHTTPMatch `json:"rule_http_match,omitempty"`

	// BackendStageID: ID of the backend stage that requests matching the rule should be forwarded to.
	// Precisely one of BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// RouteRule: route rule.
type RouteRule struct {
	// RuleHTTPMatch: rule condition to be matched. Requests matching the condition defined here will be directly forwarded to the backend specified by the `backend_stage_id` field. Requests that do not match will be checked by the next rule's condition.
	// Precisely one of RuleHTTPMatch must be set.
	RuleHTTPMatch *RuleHTTPMatch `json:"rule_http_match,omitempty"`

	// BackendStageID: ID of the backend stage that requests matching the rule should be forwarded to.
	// Precisely one of BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`

	// Position: position of the rule which determines the order of processing within the route stage.
	Position int32 `json:"position"`

	// RouteStageID: route stage ID the route rule belongs to.
	RouteStageID string `json:"route_stage_id"`
}

// CheckPEMChainRequestSecretChain: check pem chain request secret chain.
type CheckPEMChainRequestSecretChain struct {
	SecretID string `json:"secret_id"`

	SecretRegion string `json:"secret_region"`
}

// PlanDetails: plan details.
type PlanDetails struct {
	// PlanName: subscription plan name.
	// Default value: unknown_name
	PlanName PlanName `json:"plan_name"`

	// PackageGb: amount of egress data from cache included in subscription plan.
	PackageGb uint64 `json:"package_gb"`

	// PipelineLimit: number of pipelines included in subscription plan.
	PipelineLimit uint32 `json:"pipeline_limit"`

	// WafRequests: number of WAF requests included in subscription plan.
	WafRequests uint64 `json:"waf_requests"`
}

// PlanUsageDetails: plan usage details.
type PlanUsageDetails struct {
	// PlanCost: cost to date (this month) for the corresponding Edge Services subscription plan.
	PlanCost *scw.Money `json:"plan_cost"`
}

// HeadStageResponseHeadStage: head stage response head stage.
type HeadStageResponseHeadStage struct {
	// Precisely one of DNSStageID must be set.
	DNSStageID *string `json:"dns_stage_id,omitempty"`
}

// ListHeadStagesResponseHeadStage: list head stages response head stage.
type ListHeadStagesResponseHeadStage struct {
	// Precisely one of DNSStageID must be set.
	DNSStageID *string `json:"dns_stage_id,omitempty"`
}

// PipelineStages: pipeline stages.
type PipelineStages struct {
	Pipeline *Pipeline `json:"pipeline"`

	DNSStages []*DNSStage `json:"dns_stages"`

	TLSStages []*TLSStage `json:"tls_stages"`

	CacheStages []*CacheStage `json:"cache_stages"`

	BackendStages []*BackendStage `json:"backend_stages"`

	WafStages []*WafStage `json:"waf_stages"`

	RouteStages []*RouteStage `json:"route_stages"`
}

// PurgeRequest: purge request.
type PurgeRequest struct {
	// ID: ID of the purge request.
	ID string `json:"id"`

	// PipelineID: pipeline ID the purge request belongs to.
	PipelineID string `json:"pipeline_id"`

	// Status: status of the purge request.
	// Default value: unknown_status
	Status PurgeRequestStatus `json:"status"`

	// Assets: list of asserts to purge.
	// Precisely one of Assets, All must be set.
	Assets *[]string `json:"assets,omitempty"`

	// All: defines whether to purge all content.
	// Precisely one of Assets, All must be set.
	All *bool `json:"all,omitempty"`

	// CreatedAt: date the purge request was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the purge request was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
}

// SetHeadStageRequestAddNewHeadStage: set head stage request add new head stage.
type SetHeadStageRequestAddNewHeadStage struct {
	NewStageID string `json:"new_stage_id"`
}

// SetHeadStageRequestRemoveHeadStage: set head stage request remove head stage.
type SetHeadStageRequestRemoveHeadStage struct {
	RemoveStageID string `json:"remove_stage_id"`
}

// SetHeadStageRequestSwapHeadStage: set head stage request swap head stage.
type SetHeadStageRequestSwapHeadStage struct {
	NewStageID string `json:"new_stage_id"`

	CurrentStageID string `json:"current_stage_id"`
}

// TLSSecretsConfig: tls secrets config.
type TLSSecretsConfig struct {
	// TLSSecrets: secret information (from Secret Manager).
	TLSSecrets []*TLSSecret `json:"tls_secrets"`
}

// AddRouteRulesRequest: add route rules request.
type AddRouteRulesRequest struct {
	// RouteStageID: ID of the route stage to update.
	RouteStageID string `json:"-"`

	// RouteRules: list of rules to be checked against every HTTP request. The first matching rule will forward the request to its specified backend stage. If no rules are matched, the request is forwarded to the WAF stage defined by `waf_stage_id`.
	RouteRules []*SetRouteRulesRequestRouteRule `json:"route_rules"`

	// AfterPosition: add rules after the given position.
	// Precisely one of AfterPosition, BeforePosition must be set.
	AfterPosition *uint64 `json:"after_position,omitempty"`

	// BeforePosition: add rules before the given position.
	// Precisely one of AfterPosition, BeforePosition must be set.
	BeforePosition *uint64 `json:"before_position,omitempty"`
}

// AddRouteRulesResponse: add route rules response.
type AddRouteRulesResponse struct {
	// RouteRules: list of rules to be checked against every HTTP request. The first matching rule will forward the request to its specified backend stage. If no rules are matched, the request is forwarded to the WAF stage defined by `waf_stage_id`.
	RouteRules []*RouteRule `json:"route_rules"`
}

// CheckDomainRequest: check domain request.
type CheckDomainRequest struct {
	ProjectID string `json:"project_id"`

	Fqdn string `json:"fqdn"`

	Cname string `json:"cname"`
}

// CheckDomainResponse: check domain response.
type CheckDomainResponse struct {
	IsValid bool `json:"is_valid"`
}

// CheckLBOriginRequest: check lb origin request.
type CheckLBOriginRequest struct {
	LB *ScalewayLB `json:"lb,omitempty"`
}

// CheckLBOriginResponse: check lb origin response.
type CheckLBOriginResponse struct {
	IsValid bool `json:"is_valid"`

	// ErrorType: default value: unknown
	ErrorType LBOriginError `json:"error_type"`
}

// CheckPEMChainRequest: check pem chain request.
type CheckPEMChainRequest struct {
	ProjectID string `json:"project_id"`

	Fqdn string `json:"fqdn"`

	// Precisely one of Secret, Raw must be set.
	Secret *CheckPEMChainRequestSecretChain `json:"secret,omitempty"`

	// Precisely one of Secret, Raw must be set.
	Raw *string `json:"raw,omitempty"`
}

// CheckPEMChainResponse: check pem chain response.
type CheckPEMChainResponse struct {
	IsValid bool `json:"is_valid"`
}

// CreateBackendStageRequest: create backend stage request.
type CreateBackendStageRequest struct {
	// PipelineID: pipeline ID the Backend stage belongs to.
	PipelineID string `json:"-"`

	// ScalewayS3: scaleway Object Storage origin bucket (S3) linked to the backend stage.
	// Precisely one of ScalewayS3, ScalewayLB, ScalewayServerlessContainer, ScalewayServerlessFunction must be set.
	ScalewayS3 *ScalewayS3BackendConfig `json:"scaleway_s3,omitempty"`

	// ScalewayLB: scaleway Load Balancer origin linked to the backend stage.
	// Precisely one of ScalewayS3, ScalewayLB, ScalewayServerlessContainer, ScalewayServerlessFunction must be set.
	ScalewayLB *ScalewayLBBackendConfig `json:"scaleway_lb,omitempty"`

	// Precisely one of ScalewayS3, ScalewayLB, ScalewayServerlessContainer, ScalewayServerlessFunction must be set.
	ScalewayServerlessContainer *ScalewayServerlessContainerBackendConfig `json:"scaleway_serverless_container,omitempty"`

	// Precisely one of ScalewayS3, ScalewayLB, ScalewayServerlessContainer, ScalewayServerlessFunction must be set.
	ScalewayServerlessFunction *ScalewayServerlessFunctionBackendConfig `json:"scaleway_serverless_function,omitempty"`
}

// CreateCacheStageRequest: create cache stage request.
type CreateCacheStageRequest struct {
	// PipelineID: pipeline ID the Cache stage belongs to.
	PipelineID string `json:"-"`

	// FallbackTTL: time To Live (TTL) in seconds. Defines how long content is cached.
	FallbackTTL *scw.Duration `json:"fallback_ttl,omitempty"`

	// IncludeCookies: defines whether responses to requests with cookies must be stored in the cache.
	IncludeCookies *bool `json:"include_cookies,omitempty"`

	// BackendStageID: backend stage ID the cache stage will be linked to.
	// Precisely one of BackendStageID, WafStageID, RouteStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`

	// Precisely one of BackendStageID, WafStageID, RouteStageID must be set.
	WafStageID *string `json:"waf_stage_id,omitempty"`

	// Precisely one of BackendStageID, WafStageID, RouteStageID must be set.
	RouteStageID *string `json:"route_stage_id,omitempty"`
}

// CreateDNSStageRequest: create dns stage request.
type CreateDNSStageRequest struct {
	// PipelineID: pipeline ID the DNS stage belongs to.
	PipelineID string `json:"-"`

	// Fqdns: fully Qualified Domain Name (in the format subdomain.example.com) to attach to the stage.
	Fqdns *[]string `json:"fqdns,omitempty"`

	// TLSStageID: TLS stage ID the DNS stage will be linked to.
	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	TLSStageID *string `json:"tls_stage_id,omitempty"`

	// CacheStageID: cache stage ID the DNS stage will be linked to.
	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// BackendStageID: backend stage ID the DNS stage will be linked to.
	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// CreatePipelineRequest: create pipeline request.
type CreatePipelineRequest struct {
	// ProjectID: project ID in which the pipeline will be created.
	ProjectID string `json:"project_id"`

	// Name: name of the pipeline.
	Name string `json:"name"`

	// Description: description of the pipeline.
	Description string `json:"description"`
}

// CreatePurgeRequestRequest: create purge request request.
type CreatePurgeRequestRequest struct {
	// PipelineID: pipeline ID in which the purge request will be created.
	PipelineID string `json:"pipeline_id"`

	// Assets: list of asserts to purge.
	// Precisely one of Assets, All must be set.
	Assets *[]string `json:"assets,omitempty"`

	// All: defines whether to purge all content.
	// Precisely one of Assets, All must be set.
	All *bool `json:"all,omitempty"`
}

// CreateRouteStageRequest: create route stage request.
type CreateRouteStageRequest struct {
	// PipelineID: pipeline ID the route stage belongs to.
	PipelineID string `json:"-"`

	// WafStageID: ID of the WAF stage HTTP requests should be forwarded to when no rules are matched.
	// Precisely one of WafStageID must be set.
	WafStageID *string `json:"waf_stage_id,omitempty"`
}

// CreateTLSStageRequest: create tls stage request.
type CreateTLSStageRequest struct {
	// PipelineID: pipeline ID the TLS stage belongs to.
	PipelineID string `json:"-"`

	// Secrets: secret (from Scaleway Secret Manager) containing your custom certificate.
	Secrets []*TLSSecret `json:"secrets"`

	// ManagedCertificate: true when Scaleway generates and manages a Let's Encrypt certificate for the TLS stage/custom endpoint.
	ManagedCertificate *bool `json:"managed_certificate,omitempty"`

	// CacheStageID: cache stage ID the TLS stage will be linked to.
	// Precisely one of CacheStageID, BackendStageID, RouteStageID, WafStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// BackendStageID: backend stage ID the TLS stage will be linked to.
	// Precisely one of CacheStageID, BackendStageID, RouteStageID, WafStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`

	// Precisely one of CacheStageID, BackendStageID, RouteStageID, WafStageID must be set.
	RouteStageID *string `json:"route_stage_id,omitempty"`

	// Precisely one of CacheStageID, BackendStageID, RouteStageID, WafStageID must be set.
	WafStageID *string `json:"waf_stage_id,omitempty"`
}

// CreateWafStageRequest: create waf stage request.
type CreateWafStageRequest struct {
	// PipelineID: pipeline ID the WAF stage belongs to.
	PipelineID string `json:"-"`

	// Mode: mode defining WAF behavior (`disable`/`log_only`/`enable`).
	// Default value: unknown_mode
	Mode WafStageMode `json:"mode"`

	// ParanoiaLevel: sensitivity level (`1`,`2`,`3`,`4`) to use when classifying requests as malicious. With a high level, requests are more likely to be classed as malicious, and false positives are expected. With a lower level, requests are more likely to be classed as benign.
	ParanoiaLevel uint32 `json:"paranoia_level"`

	// BackendStageID: ID of the backend stage to forward requests to after the WAF stage.
	// Precisely one of BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// DeleteBackendStageRequest: delete backend stage request.
type DeleteBackendStageRequest struct {
	// BackendStageID: ID of the backend stage to delete.
	BackendStageID string `json:"-"`
}

// DeleteCacheStageRequest: delete cache stage request.
type DeleteCacheStageRequest struct {
	// CacheStageID: ID of the cache stage to delete.
	CacheStageID string `json:"-"`
}

// DeleteCurrentPlanRequest: delete current plan request.
type DeleteCurrentPlanRequest struct {
	ProjectID string `json:"-"`
}

// DeleteDNSStageRequest: delete dns stage request.
type DeleteDNSStageRequest struct {
	// DNSStageID: ID of the DNS stage to delete.
	DNSStageID string `json:"-"`
}

// DeletePipelineRequest: delete pipeline request.
type DeletePipelineRequest struct {
	// PipelineID: ID of the pipeline to delete.
	PipelineID string `json:"-"`
}

// DeleteRouteStageRequest: delete route stage request.
type DeleteRouteStageRequest struct {
	// RouteStageID: ID of the route stage to delete.
	RouteStageID string `json:"-"`
}

// DeleteTLSStageRequest: delete tls stage request.
type DeleteTLSStageRequest struct {
	// TLSStageID: ID of the TLS stage to delete.
	TLSStageID string `json:"-"`
}

// DeleteWafStageRequest: delete waf stage request.
type DeleteWafStageRequest struct {
	// WafStageID: ID of the WAF stage to delete.
	WafStageID string `json:"-"`
}

// GetBackendStageRequest: get backend stage request.
type GetBackendStageRequest struct {
	// BackendStageID: ID of the requested backend stage.
	BackendStageID string `json:"-"`
}

// GetBillingRequest: get billing request.
type GetBillingRequest struct {
	ProjectID string `json:"-"`
}

// GetBillingResponse: get billing response.
type GetBillingResponse struct {
	// CurrentPlan: information on the currently-selected, active Edge Services subscription plan.
	CurrentPlan *PlanDetails `json:"current_plan"`

	// PlanCost: cost to date (this month) for Edge Service subscription plans. This comprises the pro-rata cost of the current subscription plan, and any previous subscription plans that were active earlier in the month.
	PlanCost *scw.Money `json:"plan_cost"`

	// PipelineNumber: total number of pipelines currently configured.
	PipelineNumber uint32 `json:"pipeline_number"`

	// ExtraPipelinesCost: cost to date (this month) of pipelines not included in the subscription plans.
	ExtraPipelinesCost *scw.Money `json:"extra_pipelines_cost"`

	// CurrentPlanCacheUsage: total amount of data egressed from the cache in gigabytes from the beginning of the month, for the active subscription plan.
	CurrentPlanCacheUsage uint64 `json:"current_plan_cache_usage"`

	// ExtraCacheUsage: total amount of extra data egressed from cache in gigabytes from the beginning of the month, not included in the subscription plans.
	ExtraCacheUsage uint64 `json:"extra_cache_usage"`

	// ExtraCacheCost: cost to date (this month) of the data egressed from the cache that is not included in the subscription plans.
	ExtraCacheCost *scw.Money `json:"extra_cache_cost"`

	// CurrentPlanWafUsage: total number of requests processed by the WAF since the beginning of the current month, for the active subscription plan.
	CurrentPlanWafUsage uint64 `json:"current_plan_waf_usage"`

	// ExtraWafUsage: total number of extra requests processed by the WAF from the beginning of the month, not included in the subscription plans.
	ExtraWafUsage uint64 `json:"extra_waf_usage"`

	// ExtraWafCost: cost to date (this month) of the extra requests processed by the WAF that were not included in the subscription plans.
	ExtraWafCost *scw.Money `json:"extra_waf_cost"`

	// WafAddOn: cost of activating WAF add-on (where subscription plan does not include WAF).
	WafAddOn *scw.Money `json:"waf_add_on"`

	// PlansUsageDetails: detailed costs and usage for all Edge Services subscription plans that were activated during the month.
	PlansUsageDetails map[string]*PlanUsageDetails `json:"plans_usage_details"`

	// TotalCost: total cost to date (this month) of all Edge Services resources including active subscription plan, previously active plans, extra pipelines and extra egress cache data.
	TotalCost *scw.Money `json:"total_cost"`
}

// GetCacheStageRequest: get cache stage request.
type GetCacheStageRequest struct {
	// CacheStageID: ID of the requested cache stage.
	CacheStageID string `json:"-"`
}

// GetCurrentPlanRequest: get current plan request.
type GetCurrentPlanRequest struct {
	ProjectID string `json:"-"`
}

// GetDNSStageRequest: get dns stage request.
type GetDNSStageRequest struct {
	// DNSStageID: ID of the requested DNS stage.
	DNSStageID string `json:"-"`
}

// GetPipelineRequest: get pipeline request.
type GetPipelineRequest struct {
	// PipelineID: ID of the requested pipeline.
	PipelineID string `json:"-"`
}

// GetPurgeRequestRequest: get purge request request.
type GetPurgeRequestRequest struct {
	// PurgeRequestID: ID of the requested purge request.
	PurgeRequestID string `json:"-"`
}

// GetRouteStageRequest: get route stage request.
type GetRouteStageRequest struct {
	// RouteStageID: ID of the requested route stage.
	RouteStageID string `json:"-"`
}

// GetTLSStageRequest: get tls stage request.
type GetTLSStageRequest struct {
	// TLSStageID: ID of the requested TLS stage.
	TLSStageID string `json:"-"`
}

// GetWafStageRequest: get waf stage request.
type GetWafStageRequest struct {
	// WafStageID: ID of the requested WAF stage.
	WafStageID string `json:"-"`
}

// HeadStageResponse: head stage response.
type HeadStageResponse struct {
	// HeadStage: modified or created head stage.
	HeadStage *HeadStageResponseHeadStage `json:"head_stage"`
}

// ListBackendStagesRequest: list backend stages request.
type ListBackendStagesRequest struct {
	// PipelineID: pipeline ID to filter for. Only backend stages from this pipeline will be returned.
	PipelineID string `json:"-"`

	// OrderBy: sort order of backend stages in the response.
	// Default value: created_at_asc
	OrderBy ListBackendStagesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of backend stages to return per page.
	PageSize *uint32 `json:"-"`

	// BucketName: bucket name to filter for. Only backend stages from this Bucket will be returned.
	BucketName *string `json:"-"`

	// BucketRegion: bucket region to filter for. Only backend stages with buckets in this region will be returned.
	BucketRegion *string `json:"-"`

	// LBID: load Balancer ID to filter for. Only backend stages with this Load Balancer will be returned.
	LBID *string `json:"-"`
}

// ListBackendStagesResponse: list backend stages response.
type ListBackendStagesResponse struct {
	// Stages: paginated list of backend stages.
	Stages []*BackendStage `json:"stages"`

	// TotalCount: count of all backend stages matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBackendStagesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBackendStagesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListBackendStagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Stages = append(r.Stages, results.Stages...)
	r.TotalCount += uint64(len(results.Stages))
	return uint64(len(results.Stages)), nil
}

// ListCacheStagesRequest: list cache stages request.
type ListCacheStagesRequest struct {
	// PipelineID: pipeline ID to filter for. Only cache stages from this pipeline will be returned.
	PipelineID string `json:"-"`

	// OrderBy: sort order of cache stages in the response.
	// Default value: created_at_asc
	OrderBy ListCacheStagesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of cache stages to return per page.
	PageSize *uint32 `json:"-"`
}

// ListCacheStagesResponse: list cache stages response.
type ListCacheStagesResponse struct {
	// Stages: paginated list of cache stages.
	Stages []*CacheStage `json:"stages"`

	// TotalCount: count of all cache stages matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCacheStagesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCacheStagesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListCacheStagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Stages = append(r.Stages, results.Stages...)
	r.TotalCount += uint64(len(results.Stages))
	return uint64(len(results.Stages)), nil
}

// ListDNSStagesRequest: list dns stages request.
type ListDNSStagesRequest struct {
	// PipelineID: pipeline ID to filter for. Only DNS stages from this pipeline will be returned.
	PipelineID string `json:"-"`

	// OrderBy: sort order of DNS stages in the response.
	// Default value: created_at_asc
	OrderBy ListDNSStagesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of DNS stages to return per page.
	PageSize *uint32 `json:"-"`

	// Fqdn: fully Qualified Domain Name to filter for (in the format subdomain.example.com). Only DNS stages with this FQDN will be returned.
	Fqdn *string `json:"-"`
}

// ListDNSStagesResponse: list dns stages response.
type ListDNSStagesResponse struct {
	// Stages: paginated list of DNS stages.
	Stages []*DNSStage `json:"stages"`

	// TotalCount: count of all DNS stages matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDNSStagesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDNSStagesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListDNSStagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Stages = append(r.Stages, results.Stages...)
	r.TotalCount += uint64(len(results.Stages))
	return uint64(len(results.Stages)), nil
}

// ListHeadStagesRequest: list head stages request.
type ListHeadStagesRequest struct {
	// PipelineID: ID of the pipeline to update.
	PipelineID string `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of head stages to return per page.
	PageSize *uint32 `json:"-"`
}

// ListHeadStagesResponse: list head stages response.
type ListHeadStagesResponse struct {
	// HeadStages: number of head stages to return per page.
	HeadStages []*ListHeadStagesResponseHeadStage `json:"head_stages"`

	// TotalCount: count of all head stages matching the requested pipeline_id.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListHeadStagesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListHeadStagesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListHeadStagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.HeadStages = append(r.HeadStages, results.HeadStages...)
	r.TotalCount += uint64(len(results.HeadStages))
	return uint64(len(results.HeadStages)), nil
}

// ListPipelinesRequest: list pipelines request.
type ListPipelinesRequest struct {
	// OrderBy: sort order of pipelines in the response.
	// Default value: created_at_asc
	OrderBy ListPipelinesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of pipelines to return per page.
	PageSize *uint32 `json:"-"`

	// Name: pipeline name to filter for. Only pipelines with this string within their name will be returned.
	Name *string `json:"-"`

	// OrganizationID: organization ID to filter for. Only pipelines from this Organization will be returned.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID to filter for. Only pipelines from this Project will be returned.
	ProjectID *string `json:"-"`

	// HasBackendStageLB: filter on backend stage. Only pipelines with a Load Balancer origin will be returned.
	HasBackendStageLB *bool `json:"-"`
}

// ListPipelinesResponse: list pipelines response.
type ListPipelinesResponse struct {
	// Pipelines: paginated list of pipelines.
	Pipelines []*Pipeline `json:"pipelines"`

	// TotalCount: count of all pipelines matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPipelinesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPipelinesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListPipelinesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Pipelines = append(r.Pipelines, results.Pipelines...)
	r.TotalCount += uint64(len(results.Pipelines))
	return uint64(len(results.Pipelines)), nil
}

// ListPipelinesWithStagesRequest: list pipelines with stages request.
type ListPipelinesWithStagesRequest struct {
	// OrderBy: default value: created_at_asc
	OrderBy ListPipelinesWithStagesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// ListPipelinesWithStagesResponse: list pipelines with stages response.
type ListPipelinesWithStagesResponse struct {
	Pipelines []*PipelineStages `json:"pipelines"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPipelinesWithStagesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPipelinesWithStagesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListPipelinesWithStagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Pipelines = append(r.Pipelines, results.Pipelines...)
	r.TotalCount += uint64(len(results.Pipelines))
	return uint64(len(results.Pipelines)), nil
}

// ListPlansResponse: list plans response.
type ListPlansResponse struct {
	TotalCount uint64 `json:"total_count"`

	Plans []*PlanDetails `json:"plans"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPlansResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPlansResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListPlansResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Plans = append(r.Plans, results.Plans...)
	r.TotalCount += uint64(len(results.Plans))
	return uint64(len(results.Plans)), nil
}

// ListPurgeRequestsRequest: list purge requests request.
type ListPurgeRequestsRequest struct {
	// OrderBy: sort order of purge requests in the response.
	// Default value: created_at_asc
	OrderBy ListPurgeRequestsRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of purge requests to return per page.
	PageSize *uint32 `json:"-"`

	// OrganizationID: organization ID to filter for. Only purge requests from this Project will be returned.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID to filter for. Only purge requests from this Project will be returned.
	ProjectID *string `json:"-"`

	// PipelineID: pipeline ID to filter for. Only purge requests from this pipeline will be returned.
	PipelineID *string `json:"-"`
}

// ListPurgeRequestsResponse: list purge requests response.
type ListPurgeRequestsResponse struct {
	// PurgeRequests: paginated list of purge requests.
	PurgeRequests []*PurgeRequest `json:"purge_requests"`

	// TotalCount: count of all purge requests matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPurgeRequestsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPurgeRequestsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListPurgeRequestsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PurgeRequests = append(r.PurgeRequests, results.PurgeRequests...)
	r.TotalCount += uint64(len(results.PurgeRequests))
	return uint64(len(results.PurgeRequests)), nil
}

// ListRouteRulesRequest: list route rules request.
type ListRouteRulesRequest struct {
	// RouteStageID: route stage ID to filter for. Only route rules from this route stage will be returned.
	RouteStageID string `json:"-"`
}

// ListRouteRulesResponse: list route rules response.
type ListRouteRulesResponse struct {
	// RouteRules: list of rules to be checked against every HTTP request. The first matching rule will forward the request to its specified backend stage. If no rules are matched, the request is forwarded to the WAF stage defined by `waf_stage_id`.
	RouteRules []*RouteRule `json:"route_rules"`

	// TotalCount: count of all route rules matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRouteRulesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRouteRulesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListRouteRulesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.RouteRules = append(r.RouteRules, results.RouteRules...)
	r.TotalCount += uint64(len(results.RouteRules))
	return uint64(len(results.RouteRules)), nil
}

// ListRouteStagesRequest: list route stages request.
type ListRouteStagesRequest struct {
	// PipelineID: pipeline ID to filter for. Only route stages from this pipeline will be returned.
	PipelineID string `json:"-"`

	// OrderBy: sort order of route stages in the response.
	// Default value: created_at_asc
	OrderBy ListRouteStagesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of route stages to return per page.
	PageSize *uint32 `json:"-"`
}

// ListRouteStagesResponse: list route stages response.
type ListRouteStagesResponse struct {
	// Stages: paginated list of summarized route stages.
	Stages []*RouteStage `json:"stages"`

	// TotalCount: count of all route stages matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRouteStagesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRouteStagesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListRouteStagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Stages = append(r.Stages, results.Stages...)
	r.TotalCount += uint64(len(results.Stages))
	return uint64(len(results.Stages)), nil
}

// ListTLSStagesRequest: list tls stages request.
type ListTLSStagesRequest struct {
	// PipelineID: pipeline ID to filter for. Only TLS stages from this pipeline will be returned.
	PipelineID string `json:"-"`

	// OrderBy: sort order of TLS stages in the response.
	// Default value: created_at_asc
	OrderBy ListTLSStagesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of TLS stages to return per page.
	PageSize *uint32 `json:"-"`

	// SecretID: secret ID to filter for. Only TLS stages with this Secret ID will be returned.
	SecretID *string `json:"-"`

	// SecretRegion: secret region to filter for. Only TLS stages with a Secret in this region will be returned.
	SecretRegion *string `json:"-"`
}

// ListTLSStagesResponse: list tls stages response.
type ListTLSStagesResponse struct {
	// Stages: paginated list of TLS stages.
	Stages []*TLSStage `json:"stages"`

	// TotalCount: count of all TLS stages matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTLSStagesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTLSStagesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListTLSStagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Stages = append(r.Stages, results.Stages...)
	r.TotalCount += uint64(len(results.Stages))
	return uint64(len(results.Stages)), nil
}

// ListWafStagesRequest: list waf stages request.
type ListWafStagesRequest struct {
	// PipelineID: pipeline ID to filter for. Only WAF stages from this pipeline will be returned.
	PipelineID string `json:"-"`

	// OrderBy: sort order of WAF stages in the response.
	// Default value: created_at_asc
	OrderBy ListWafStagesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of WAF stages to return per page.
	PageSize *uint32 `json:"-"`
}

// ListWafStagesResponse: list waf stages response.
type ListWafStagesResponse struct {
	// Stages: paginated list of WAF stages.
	Stages []*WafStage `json:"stages"`

	// TotalCount: count of all WAF stages matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListWafStagesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListWafStagesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListWafStagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Stages = append(r.Stages, results.Stages...)
	r.TotalCount += uint64(len(results.Stages))
	return uint64(len(results.Stages)), nil
}

// Plan: plan.
type Plan struct {
	// PlanName: default value: unknown_name
	PlanName PlanName `json:"plan_name"`
}

// SearchBackendStagesRequest: search backend stages request.
type SearchBackendStagesRequest struct {
	// OrderBy: default value: created_at_asc
	OrderBy SearchBackendStagesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	ProjectID string `json:"-"`

	BucketName *string `json:"-"`

	BucketRegion *string `json:"-"`

	LBID *string `json:"-"`
}

// SearchRouteRulesRequest: search route rules request.
type SearchRouteRulesRequest struct {
	// OrderBy: default value: created_at_asc
	OrderBy SearchRouteRulesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// SearchWafStagesRequest: search waf stages request.
type SearchWafStagesRequest struct {
	// OrderBy: default value: created_at_asc
	OrderBy SearchWafStagesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	ProjectID string `json:"-"`
}

// SelectPlanRequest: select plan request.
type SelectPlanRequest struct {
	ProjectID string `json:"project_id"`

	// PlanName: default value: unknown_name
	PlanName PlanName `json:"plan_name"`
}

// SetHeadStageRequest: set head stage request.
type SetHeadStageRequest struct {
	// PipelineID: ID of the pipeline to update.
	PipelineID string `json:"-"`

	// AddNewHeadStage: add a new head stage.
	// Precisely one of AddNewHeadStage, RemoveHeadStage, SwapHeadStage must be set.
	AddNewHeadStage *SetHeadStageRequestAddNewHeadStage `json:"add_new_head_stage,omitempty"`

	// RemoveHeadStage: remove a head stage.
	// Precisely one of AddNewHeadStage, RemoveHeadStage, SwapHeadStage must be set.
	RemoveHeadStage *SetHeadStageRequestRemoveHeadStage `json:"remove_head_stage,omitempty"`

	// SwapHeadStage: replace a head stage with a new one.
	// Precisely one of AddNewHeadStage, RemoveHeadStage, SwapHeadStage must be set.
	SwapHeadStage *SetHeadStageRequestSwapHeadStage `json:"swap_head_stage,omitempty"`
}

// SetRouteRulesRequest: set route rules request.
type SetRouteRulesRequest struct {
	// RouteStageID: ID of the route stage to update.
	RouteStageID string `json:"-"`

	// RouteRules: list of rules to be checked against every HTTP request. The first matching rule will forward the request to its specified backend stage. If no rules are matched, the request is forwarded to the WAF stage defined by `waf_stage_id`.
	RouteRules []*SetRouteRulesRequestRouteRule `json:"route_rules"`
}

// SetRouteRulesResponse: set route rules response.
type SetRouteRulesResponse struct {
	// RouteRules: list of rules to be checked against every HTTP request. The first matching rule will forward the request to its specified backend stage. If no rules are matched, the request is forwarded to the WAF stage defined by `waf_stage_id`.
	RouteRules []*RouteRule `json:"route_rules"`
}

// UpdateBackendStageRequest: update backend stage request.
type UpdateBackendStageRequest struct {
	// BackendStageID: ID of the backend stage to update.
	BackendStageID string `json:"-"`

	// ScalewayS3: scaleway Object Storage origin bucket (S3) linked to the backend stage.
	// Precisely one of ScalewayS3, ScalewayLB, ScalewayServerlessContainer, ScalewayServerlessFunction must be set.
	ScalewayS3 *ScalewayS3BackendConfig `json:"scaleway_s3,omitempty"`

	// ScalewayLB: scaleway Load Balancer origin linked to the backend stage.
	// Precisely one of ScalewayS3, ScalewayLB, ScalewayServerlessContainer, ScalewayServerlessFunction must be set.
	ScalewayLB *ScalewayLBBackendConfig `json:"scaleway_lb,omitempty"`

	// Precisely one of ScalewayS3, ScalewayLB, ScalewayServerlessContainer, ScalewayServerlessFunction must be set.
	ScalewayServerlessContainer *ScalewayServerlessContainerBackendConfig `json:"scaleway_serverless_container,omitempty"`

	// Precisely one of ScalewayS3, ScalewayLB, ScalewayServerlessContainer, ScalewayServerlessFunction must be set.
	ScalewayServerlessFunction *ScalewayServerlessFunctionBackendConfig `json:"scaleway_serverless_function,omitempty"`

	// PipelineID: pipeline ID the Backend stage belongs to.
	PipelineID string `json:"pipeline_id"`
}

// UpdateCacheStageRequest: update cache stage request.
type UpdateCacheStageRequest struct {
	// CacheStageID: ID of the cache stage to update.
	CacheStageID string `json:"-"`

	// FallbackTTL: time To Live (TTL) in seconds. Defines how long content is cached.
	FallbackTTL *scw.Duration `json:"fallback_ttl,omitempty"`

	// IncludeCookies: defines whether responses to requests with cookies must be stored in the cache.
	IncludeCookies *bool `json:"include_cookies,omitempty"`

	// BackendStageID: backend stage ID the cache stage will be linked to.
	// Precisely one of BackendStageID, WafStageID, RouteStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`

	// Precisely one of BackendStageID, WafStageID, RouteStageID must be set.
	WafStageID *string `json:"waf_stage_id,omitempty"`

	// Precisely one of BackendStageID, WafStageID, RouteStageID must be set.
	RouteStageID *string `json:"route_stage_id,omitempty"`
}

// UpdateDNSStageRequest: update dns stage request.
type UpdateDNSStageRequest struct {
	// DNSStageID: ID of the DNS stage to update.
	DNSStageID string `json:"-"`

	// Fqdns: fully Qualified Domain Name (in the format subdomain.example.com) attached to the stage.
	Fqdns *[]string `json:"fqdns,omitempty"`

	// TLSStageID: TLS stage ID the DNS stage will be linked to.
	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	TLSStageID *string `json:"tls_stage_id,omitempty"`

	// CacheStageID: cache stage ID the DNS stage will be linked to.
	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// BackendStageID: backend stage ID the DNS stage will be linked to.
	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// UpdatePipelineRequest: update pipeline request.
type UpdatePipelineRequest struct {
	// PipelineID: ID of the pipeline to update.
	PipelineID string `json:"-"`

	// Name: name of the pipeline.
	Name *string `json:"name,omitempty"`

	// Description: description of the pipeline.
	Description *string `json:"description,omitempty"`
}

// UpdateRouteStageRequest: update route stage request.
type UpdateRouteStageRequest struct {
	// RouteStageID: ID of the route stage to update.
	RouteStageID string `json:"-"`

	// WafStageID: ID of the WAF stage HTTP requests should be forwarded to when no rules are matched.
	// Precisely one of WafStageID must be set.
	WafStageID *string `json:"waf_stage_id,omitempty"`
}

// UpdateTLSStageRequest: update tls stage request.
type UpdateTLSStageRequest struct {
	// TLSStageID: ID of the TLS stage to update.
	TLSStageID string `json:"-"`

	// TLSSecretsConfig: secret (from Scaleway Secret-Manager) containing your custom certificate.
	TLSSecretsConfig *TLSSecretsConfig `json:"tls_secrets_config,omitempty"`

	// ManagedCertificate: true when Scaleway generates and manages a Let's Encrypt certificate for the TLS stage/custom endpoint.
	ManagedCertificate *bool `json:"managed_certificate,omitempty"`

	// CacheStageID: cache stage ID the TLS stage will be linked to.
	// Precisely one of CacheStageID, BackendStageID, RouteStageID, WafStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// BackendStageID: backend stage ID the TLS stage will be linked to.
	// Precisely one of CacheStageID, BackendStageID, RouteStageID, WafStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`

	// Precisely one of CacheStageID, BackendStageID, RouteStageID, WafStageID must be set.
	RouteStageID *string `json:"route_stage_id,omitempty"`

	// Precisely one of CacheStageID, BackendStageID, RouteStageID, WafStageID must be set.
	WafStageID *string `json:"waf_stage_id,omitempty"`
}

// UpdateWafStageRequest: update waf stage request.
type UpdateWafStageRequest struct {
	// WafStageID: ID of the WAF stage to update.
	WafStageID string `json:"-"`

	// Mode: mode defining WAF behavior (`disable`/`log_only`/`enable`).
	// Default value: unknown_mode
	Mode WafStageMode `json:"mode"`

	// ParanoiaLevel: sensitivity level (`1`,`2`,`3`,`4`) to use when classifying requests as malicious. With a high level, requests are more likely to be classed as malicious, and false positives are expected. With a lower level, requests are more likely to be classed as benign.
	ParanoiaLevel *uint32 `json:"paranoia_level,omitempty"`

	// BackendStageID: ID of the backend stage to forward requests to after the WAF stage.
	// Precisely one of BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
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

// ListPipelines: List all pipelines, for a Scaleway Organization or Scaleway Project. By default, the pipelines returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListPipelines(req *ListPipelinesRequest, opts ...scw.RequestOption) (*ListPipelinesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "has_backend_stage_lb", req.HasBackendStageLB)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/pipelines",
		Query:  query,
	}

	var resp ListPipelinesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePipeline: Create a new pipeline. You must specify a `dns_stage_id` to form a stage-chain that goes all the way to the backend stage (origin), so the HTTP request will be processed according to the stages you created.
func (s *API) CreatePipeline(req *CreatePipelineRequest, opts ...scw.RequestOption) (*Pipeline, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/pipelines",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Pipeline

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPipeline: Retrieve information about an existing pipeline, specified by its `pipeline_id`. Its full details, including errors, are returned in the response object.
func (s *API) GetPipeline(req *GetPipelineRequest, opts ...scw.RequestOption) (*Pipeline, error) {
	var err error

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "",
	}

	var resp Pipeline

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForPipelineRequest is used by WaitForPipeline method.
type WaitForPipelineRequest struct {
	PipelineID    string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForPipeline waits for the Pipeline to reach a terminal state.
func (s *API) WaitForPipeline(req *WaitForPipelineRequest, opts ...scw.RequestOption) (*Pipeline, error) {
	timeout := defaultEdgeServicesTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultEdgeServicesRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[PipelineStatus]struct{}{
		PipelineStatusPending: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetPipeline(&GetPipelineRequest{
				PipelineID: req.PipelineID,
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
		return nil, errors.Wrap(err, "waiting for Pipeline failed")
	}

	return res.(*Pipeline), nil
}

// ListPipelinesWithStages:
func (s *API) ListPipelinesWithStages(req *ListPipelinesWithStagesRequest, opts ...scw.RequestOption) (*ListPipelinesWithStagesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/pipelines-stages",
		Query:  query,
	}

	var resp ListPipelinesWithStagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePipeline: Update the parameters of an existing pipeline, specified by its `pipeline_id`. Parameters which can be updated include the `name`, `description` and `dns_stage_id`.
func (s *API) UpdatePipeline(req *UpdatePipelineRequest, opts ...scw.RequestOption) (*Pipeline, error) {
	var err error

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Pipeline

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePipeline: Delete an existing pipeline, specified by its `pipeline_id`. Deleting a pipeline is permanent, and cannot be undone. Note that all stages linked to the pipeline are also deleted.
func (s *API) DeletePipeline(req *DeletePipelineRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.PipelineID) == "" {
		return errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListHeadStages: List Head stage for your pipeline.
func (s *API) ListHeadStages(req *ListHeadStagesRequest, opts ...scw.RequestOption) (*ListHeadStagesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/head-stages",
		Query:  query,
	}

	var resp ListHeadStagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetHeadStage: You must specify either a `add_new_head_stage` (to add a new head stage), `remove_head_stage` (to remove a head stage) or `swap_head_stage` (to replace a head stage).
func (s *API) SetHeadStage(req *SetHeadStageRequest, opts ...scw.RequestOption) (*HeadStageResponse, error) {
	var err error

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/set-head-stage",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp HeadStageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDNSStages: List all DNS stages, for a Scaleway Organization or Scaleway Project. By default, the DNS stages returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListDNSStages(req *ListDNSStagesRequest, opts ...scw.RequestOption) (*ListDNSStagesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "fqdn", req.Fqdn)

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/dns-stages",
		Query:  query,
	}

	var resp ListDNSStagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDNSStage: Create a new DNS stage. You must specify the `fqdns` field to customize the domain endpoint, using a domain you already own.
func (s *API) CreateDNSStage(req *CreateDNSStageRequest, opts ...scw.RequestOption) (*DNSStage, error) {
	var err error

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/dns-stages",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DNSStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDNSStage: Retrieve information about an existing DNS stage, specified by its `dns_stage_id`. Its full details, including FQDNs, are returned in the response object.
func (s *API) GetDNSStage(req *GetDNSStageRequest, opts ...scw.RequestOption) (*DNSStage, error) {
	var err error

	if fmt.Sprint(req.DNSStageID) == "" {
		return nil, errors.New("field DNSStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/dns-stages/" + fmt.Sprint(req.DNSStageID) + "",
	}

	var resp DNSStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDNSStage: Update the parameters of an existing DNS stage, specified by its `dns_stage_id`.
func (s *API) UpdateDNSStage(req *UpdateDNSStageRequest, opts ...scw.RequestOption) (*DNSStage, error) {
	var err error

	if fmt.Sprint(req.DNSStageID) == "" {
		return nil, errors.New("field DNSStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/edge-services/v1beta1/dns-stages/" + fmt.Sprint(req.DNSStageID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DNSStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDNSStage: Delete an existing DNS stage, specified by its `dns_stage_id`. Deleting a DNS stage is permanent, and cannot be undone.
func (s *API) DeleteDNSStage(req *DeleteDNSStageRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.DNSStageID) == "" {
		return errors.New("field DNSStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/edge-services/v1beta1/dns-stages/" + fmt.Sprint(req.DNSStageID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListTLSStages: List all TLS stages, for a Scaleway Organization or Scaleway Project. By default, the TLS stages returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListTLSStages(req *ListTLSStagesRequest, opts ...scw.RequestOption) (*ListTLSStagesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "secret_id", req.SecretID)
	parameter.AddToQuery(query, "secret_region", req.SecretRegion)

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/tls-stages",
		Query:  query,
	}

	var resp ListTLSStagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTLSStage: Create a new TLS stage. You must specify either the `secrets` or `managed_certificate` fields to customize the SSL/TLS certificate of your endpoint. Choose `secrets` if you are using a pre-existing certificate held in Scaleway Secret Manager, or `managed_certificate` to let Scaleway generate and manage a Let's Encrypt certificate for your customized endpoint.
func (s *API) CreateTLSStage(req *CreateTLSStageRequest, opts ...scw.RequestOption) (*TLSStage, error) {
	var err error

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/tls-stages",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TLSStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTLSStage: Retrieve information about an existing TLS stage, specified by its `tls_stage_id`. Its full details, including secrets and certificate expiration date are returned in the response object.
func (s *API) GetTLSStage(req *GetTLSStageRequest, opts ...scw.RequestOption) (*TLSStage, error) {
	var err error

	if fmt.Sprint(req.TLSStageID) == "" {
		return nil, errors.New("field TLSStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/tls-stages/" + fmt.Sprint(req.TLSStageID) + "",
	}

	var resp TLSStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateTLSStage: Update the parameters of an existing TLS stage, specified by its `tls_stage_id`. Both `tls_secrets_config` and `managed_certificate` parameters can be updated.
func (s *API) UpdateTLSStage(req *UpdateTLSStageRequest, opts ...scw.RequestOption) (*TLSStage, error) {
	var err error

	if fmt.Sprint(req.TLSStageID) == "" {
		return nil, errors.New("field TLSStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/edge-services/v1beta1/tls-stages/" + fmt.Sprint(req.TLSStageID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TLSStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteTLSStage: Delete an existing TLS stage, specified by its `tls_stage_id`. Deleting a TLS stage is permanent, and cannot be undone.
func (s *API) DeleteTLSStage(req *DeleteTLSStageRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.TLSStageID) == "" {
		return errors.New("field TLSStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/edge-services/v1beta1/tls-stages/" + fmt.Sprint(req.TLSStageID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListCacheStages: List all cache stages, for a Scaleway Organization or Scaleway Project. By default, the cache stages returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListCacheStages(req *ListCacheStagesRequest, opts ...scw.RequestOption) (*ListCacheStagesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/cache-stages",
		Query:  query,
	}

	var resp ListCacheStagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCacheStage: Create a new cache stage. You must specify the `fallback_ttl` field to customize the TTL of the cache.
func (s *API) CreateCacheStage(req *CreateCacheStageRequest, opts ...scw.RequestOption) (*CacheStage, error) {
	var err error

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/cache-stages",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CacheStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCacheStage: Retrieve information about an existing cache stage, specified by its `cache_stage_id`. Its full details, including Time To Live (TTL), are returned in the response object.
func (s *API) GetCacheStage(req *GetCacheStageRequest, opts ...scw.RequestOption) (*CacheStage, error) {
	var err error

	if fmt.Sprint(req.CacheStageID) == "" {
		return nil, errors.New("field CacheStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/cache-stages/" + fmt.Sprint(req.CacheStageID) + "",
	}

	var resp CacheStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCacheStage: Update the parameters of an existing cache stage, specified by its `cache_stage_id`. Parameters which can be updated include the `fallback_ttl`, `include_cookies` and `backend_stage_id`.
func (s *API) UpdateCacheStage(req *UpdateCacheStageRequest, opts ...scw.RequestOption) (*CacheStage, error) {
	var err error

	if fmt.Sprint(req.CacheStageID) == "" {
		return nil, errors.New("field CacheStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/edge-services/v1beta1/cache-stages/" + fmt.Sprint(req.CacheStageID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CacheStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCacheStage: Delete an existing cache stage, specified by its `cache_stage_id`. Deleting a cache stage is permanent, and cannot be undone.
func (s *API) DeleteCacheStage(req *DeleteCacheStageRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.CacheStageID) == "" {
		return errors.New("field CacheStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/edge-services/v1beta1/cache-stages/" + fmt.Sprint(req.CacheStageID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListBackendStages: List all backend stages, for a Scaleway Organization or Scaleway Project. By default, the backend stages returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListBackendStages(req *ListBackendStagesRequest, opts ...scw.RequestOption) (*ListBackendStagesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "bucket_name", req.BucketName)
	parameter.AddToQuery(query, "bucket_region", req.BucketRegion)
	parameter.AddToQuery(query, "lb_id", req.LBID)

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/backend-stages",
		Query:  query,
	}

	var resp ListBackendStagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBackendStage: Create a new backend stage. You must specify either a `scaleway_s3` (for a Scaleway Object Storage bucket) or `scaleway_lb` (for a Scaleway Load Balancer) field to configure the origin.
func (s *API) CreateBackendStage(req *CreateBackendStageRequest, opts ...scw.RequestOption) (*BackendStage, error) {
	var err error

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/backend-stages",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp BackendStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBackendStage: Retrieve information about an existing backend stage, specified by its `backend_stage_id`. Its full details, including `scaleway_s3` or `scaleway_lb`, are returned in the response object.
func (s *API) GetBackendStage(req *GetBackendStageRequest, opts ...scw.RequestOption) (*BackendStage, error) {
	var err error

	if fmt.Sprint(req.BackendStageID) == "" {
		return nil, errors.New("field BackendStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/backend-stages/" + fmt.Sprint(req.BackendStageID) + "",
	}

	var resp BackendStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBackendStage: Update the parameters of an existing backend stage, specified by its `backend_stage_id`.
func (s *API) UpdateBackendStage(req *UpdateBackendStageRequest, opts ...scw.RequestOption) (*BackendStage, error) {
	var err error

	if fmt.Sprint(req.BackendStageID) == "" {
		return nil, errors.New("field BackendStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/edge-services/v1beta1/backend-stages/" + fmt.Sprint(req.BackendStageID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp BackendStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBackendStage: Delete an existing backend stage, specified by its `backend_stage_id`. Deleting a backend stage is permanent, and cannot be undone.
func (s *API) DeleteBackendStage(req *DeleteBackendStageRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.BackendStageID) == "" {
		return errors.New("field BackendStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/edge-services/v1beta1/backend-stages/" + fmt.Sprint(req.BackendStageID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// SearchBackendStages:
func (s *API) SearchBackendStages(req *SearchBackendStagesRequest, opts ...scw.RequestOption) (*ListBackendStagesResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "bucket_name", req.BucketName)
	parameter.AddToQuery(query, "bucket_region", req.BucketRegion)
	parameter.AddToQuery(query, "lb_id", req.LBID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/search-backend-stages",
		Query:  query,
	}

	var resp ListBackendStagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListWafStages: List all WAF stages, for a Scaleway Organization or Scaleway Project. By default, the WAF stages returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListWafStages(req *ListWafStagesRequest, opts ...scw.RequestOption) (*ListWafStagesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/waf-stages",
		Query:  query,
	}

	var resp ListWafStagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateWafStage: Create a new WAF stage. You must specify the `mode` and `paranoia_level` fields to customize the WAF.
func (s *API) CreateWafStage(req *CreateWafStageRequest, opts ...scw.RequestOption) (*WafStage, error) {
	var err error

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/waf-stages",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp WafStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWafStage: Retrieve information about an existing WAF stage, specified by its `waf_stage_id`. Its full details are returned in the response object.
func (s *API) GetWafStage(req *GetWafStageRequest, opts ...scw.RequestOption) (*WafStage, error) {
	var err error

	if fmt.Sprint(req.WafStageID) == "" {
		return nil, errors.New("field WafStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/waf-stages/" + fmt.Sprint(req.WafStageID) + "",
	}

	var resp WafStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateWafStage: Update the parameters of an existing WAF stage, specified by its `waf_stage_id`. Both `mode` and `paranoia_level` parameters can be updated.
func (s *API) UpdateWafStage(req *UpdateWafStageRequest, opts ...scw.RequestOption) (*WafStage, error) {
	var err error

	if fmt.Sprint(req.WafStageID) == "" {
		return nil, errors.New("field WafStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/edge-services/v1beta1/waf-stages/" + fmt.Sprint(req.WafStageID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp WafStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteWafStage: Delete an existing WAF stage, specified by its `waf_stage_id`. Deleting a WAF stage is permanent, and cannot be undone.
func (s *API) DeleteWafStage(req *DeleteWafStageRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.WafStageID) == "" {
		return errors.New("field WafStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/edge-services/v1beta1/waf-stages/" + fmt.Sprint(req.WafStageID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// SearchWafStages:
func (s *API) SearchWafStages(req *SearchWafStagesRequest, opts ...scw.RequestOption) (*ListWafStagesResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/search-waf-stages",
		Query:  query,
	}

	var resp ListWafStagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRouteStages: List all route stages, for a given pipeline. By default, the route stages returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListRouteStages(req *ListRouteStagesRequest, opts ...scw.RequestOption) (*ListRouteStagesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/route-stages",
		Query:  query,
	}

	var resp ListRouteStagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRouteStage: Create a new route stage. You must specify the `waf_stage_id` field to customize the route.
func (s *API) CreateRouteStage(req *CreateRouteStageRequest, opts ...scw.RequestOption) (*RouteStage, error) {
	var err error

	if fmt.Sprint(req.PipelineID) == "" {
		return nil, errors.New("field PipelineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/pipelines/" + fmt.Sprint(req.PipelineID) + "/route-stages",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RouteStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRouteStage: Retrieve information about an existing route stage, specified by its `route_stage_id`. The summary of the route stage (without route rules) is returned in the response object.
func (s *API) GetRouteStage(req *GetRouteStageRequest, opts ...scw.RequestOption) (*RouteStage, error) {
	var err error

	if fmt.Sprint(req.RouteStageID) == "" {
		return nil, errors.New("field RouteStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/route-stages/" + fmt.Sprint(req.RouteStageID) + "",
	}

	var resp RouteStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateRouteStage: Update the parameters of an existing route stage, specified by its `route_stage_id`.
func (s *API) UpdateRouteStage(req *UpdateRouteStageRequest, opts ...scw.RequestOption) (*RouteStage, error) {
	var err error

	if fmt.Sprint(req.RouteStageID) == "" {
		return nil, errors.New("field RouteStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/edge-services/v1beta1/route-stages/" + fmt.Sprint(req.RouteStageID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RouteStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteRouteStage: Delete an existing route stage, specified by its `route_stage_id`. Deleting a route stage is permanent, and cannot be undone.
func (s *API) DeleteRouteStage(req *DeleteRouteStageRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.RouteStageID) == "" {
		return errors.New("field RouteStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/edge-services/v1beta1/route-stages/" + fmt.Sprint(req.RouteStageID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListRouteRules: List all route rules of an existing route stage, specified by its `route_stage_id`.
func (s *API) ListRouteRules(req *ListRouteRulesRequest, opts ...scw.RequestOption) (*ListRouteRulesResponse, error) {
	var err error

	if fmt.Sprint(req.RouteStageID) == "" {
		return nil, errors.New("field RouteStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/route-stages/" + fmt.Sprint(req.RouteStageID) + "/route-rules",
	}

	var resp ListRouteRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetRouteRules: Set the rules of an existing route stage, specified by its `route_stage_id`.
func (s *API) SetRouteRules(req *SetRouteRulesRequest, opts ...scw.RequestOption) (*SetRouteRulesResponse, error) {
	var err error

	if fmt.Sprint(req.RouteStageID) == "" {
		return nil, errors.New("field RouteStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/edge-services/v1beta1/route-stages/" + fmt.Sprint(req.RouteStageID) + "/route-rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetRouteRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddRouteRules: Add route rules to an existing route stage, specified by its `route_stage_id`.
func (s *API) AddRouteRules(req *AddRouteRulesRequest, opts ...scw.RequestOption) (*AddRouteRulesResponse, error) {
	var err error

	if fmt.Sprint(req.RouteStageID) == "" {
		return nil, errors.New("field RouteStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/route-stages/" + fmt.Sprint(req.RouteStageID) + "/route-rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddRouteRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SearchRouteRules: List all route rules of an organization or project.
func (s *API) SearchRouteRules(req *SearchRouteRulesRequest, opts ...scw.RequestOption) (*ListRouteRulesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/search-route-rules",
		Query:  query,
	}

	var resp ListRouteRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckDomain:
func (s *API) CheckDomain(req *CheckDomainRequest, opts ...scw.RequestOption) (*CheckDomainResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/check-domain",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CheckDomainResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckPEMChain:
func (s *API) CheckPEMChain(req *CheckPEMChainRequest, opts ...scw.RequestOption) (*CheckPEMChainResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/check-pem-chain",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CheckPEMChainResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPurgeRequests: List all purge requests, for a Scaleway Organization or Scaleway Project. This enables you to retrieve a history of all previously-made purge requests. By default, the purge requests returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListPurgeRequests(req *ListPurgeRequestsRequest, opts ...scw.RequestOption) (*ListPurgeRequestsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "pipeline_id", req.PipelineID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/purge-requests",
		Query:  query,
	}

	var resp ListPurgeRequestsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePurgeRequest: Create a new purge request. You must specify either the `all` field (to purge all content) or a list of `assets` (to define the precise assets to purge).
func (s *API) CreatePurgeRequest(req *CreatePurgeRequestRequest, opts ...scw.RequestOption) (*PurgeRequest, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/purge-requests",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PurgeRequest

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPurgeRequest: Retrieve information about a purge request, specified by its `purge_request_id`. Its full details, including `status` and `target`, are returned in the response object.
func (s *API) GetPurgeRequest(req *GetPurgeRequestRequest, opts ...scw.RequestOption) (*PurgeRequest, error) {
	var err error

	if fmt.Sprint(req.PurgeRequestID) == "" {
		return nil, errors.New("field PurgeRequestID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/purge-requests/" + fmt.Sprint(req.PurgeRequestID) + "",
	}

	var resp PurgeRequest

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForPurgeRequestRequest is used by WaitForPurgeRequest method.
type WaitForPurgeRequestRequest struct {
	PurgeRequestID string
	Timeout        *time.Duration
	RetryInterval  *time.Duration
}

// WaitForPurgeRequest waits for the PurgeRequest to reach a terminal state.
func (s *API) WaitForPurgeRequest(req *WaitForPurgeRequestRequest, opts ...scw.RequestOption) (*PurgeRequest, error) {
	timeout := defaultEdgeServicesTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultEdgeServicesRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[PurgeRequestStatus]struct{}{
		PurgeRequestStatusPending: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetPurgeRequest(&GetPurgeRequestRequest{
				PurgeRequestID: req.PurgeRequestID,
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
		return nil, errors.Wrap(err, "waiting for PurgeRequest failed")
	}

	return res.(*PurgeRequest), nil
}

// CheckLBOrigin:
func (s *API) CheckLBOrigin(req *CheckLBOriginRequest, opts ...scw.RequestOption) (*CheckLBOriginResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1beta1/check-lb-origin",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CheckLBOriginResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPlans:
func (s *API) ListPlans(opts ...scw.RequestOption) (*ListPlansResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1beta1/plans",
	}

	var resp ListPlansResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SelectPlan:
func (s *API) SelectPlan(req *SelectPlanRequest, opts ...scw.RequestOption) (*Plan, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/edge-services/v1beta1/current-plan",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Plan

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCurrentPlan:
func (s *API) GetCurrentPlan(req *GetCurrentPlanRequest, opts ...scw.RequestOption) (*Plan, error) {
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
		Path:   "/edge-services/v1beta1/current-plan/" + fmt.Sprint(req.ProjectID) + "",
	}

	var resp Plan

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCurrentPlan:
func (s *API) DeleteCurrentPlan(req *DeleteCurrentPlanRequest, opts ...scw.RequestOption) error {
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
		Path:   "/edge-services/v1beta1/current-plan/" + fmt.Sprint(req.ProjectID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetBilling: Gives information on the currently selected Edge Services subscription plan, resource usage and associated billing information for this calendar month (including whether consumption falls within or exceeds the currently selected subscription plan.).
func (s *API) GetBilling(req *GetBillingRequest, opts ...scw.RequestOption) (*GetBillingResponse, error) {
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
		Path:   "/edge-services/v1beta1/billing/" + fmt.Sprint(req.ProjectID) + "",
	}

	var resp GetBillingResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

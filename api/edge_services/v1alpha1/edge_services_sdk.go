// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package edge_services provides methods and message types of the edge_services v1alpha1 API.
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

type DNSStageType string

const (
	// The DNS stage type is unknown by default.
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
		return "unknown_type"
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
		return "unknown"
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
		return "created_at_asc"
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
		return "created_at_asc"
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
		return "created_at_asc"
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
	ListPipelinesRequestOrderByCreatedAtAsc  = ListPipelinesRequestOrderBy("created_at_asc")
	ListPipelinesRequestOrderByCreatedAtDesc = ListPipelinesRequestOrderBy("created_at_desc")
	ListPipelinesRequestOrderByNameAsc       = ListPipelinesRequestOrderBy("name_asc")
	ListPipelinesRequestOrderByNameDesc      = ListPipelinesRequestOrderBy("name_desc")
)

func (enum ListPipelinesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
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
		return "created_at_asc"
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
		return "created_at_asc"
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
		return "created_at_asc"
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
)

func (enum PipelineErrorCode) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_code"
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
		return "unknown_severity"
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
		return "unknown_stage"
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
		return "unknown_type"
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
	PipelineStatusUnknownStatus = PipelineStatus("unknown_status")
	PipelineStatusReady         = PipelineStatus("ready")
	PipelineStatusError         = PipelineStatus("error")
	PipelineStatusPending       = PipelineStatus("pending")
	PipelineStatusWarning       = PipelineStatus("warning")
)

func (enum PipelineStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
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
		return "unknown_name"
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
	// The purge request status is unknown by default.
	PurgeRequestStatusUnknownStatus = PurgeRequestStatus("unknown_status")
	// The purge request has been processed sucessfully.
	PurgeRequestStatusDone = PurgeRequestStatus("done")
	// An error occurred during the purge request.
	PurgeRequestStatusError = PurgeRequestStatus("error")
	// The purge request status is pending.
	PurgeRequestStatusPending = PurgeRequestStatus("pending")
)

func (enum PurgeRequestStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
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

// BackendStage: backend stage.
type BackendStage struct {
	// ID: ID of the backend stage.
	ID string `json:"id"`

	// PipelineID: pipeline ID the backend stage belongs to.
	PipelineID *string `json:"pipeline_id"`

	// ProjectID: project ID of the backend stage.
	ProjectID string `json:"project_id"`

	// CreatedAt: date the backend stage was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the backend stage was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// ScalewayS3: scaleway Object Storage origin bucket (S3) linked to the backend stage.
	// Precisely one of ScalewayS3, ScalewayLB must be set.
	ScalewayS3 *ScalewayS3BackendConfig `json:"scaleway_s3,omitempty"`

	// ScalewayLB: scaleway Load Balancer origin linked to the backend stage.
	// Precisely one of ScalewayS3, ScalewayLB must be set.
	ScalewayLB *ScalewayLBBackendConfig `json:"scaleway_lb,omitempty"`
}

// CacheStage: cache stage.
type CacheStage struct {
	// ID: ID of the cache stage.
	ID string `json:"id"`

	// PipelineID: pipeline ID the cache stage belongs to.
	PipelineID *string `json:"pipeline_id"`

	// ProjectID: project ID of the cache stage.
	ProjectID string `json:"project_id"`

	// FallbackTTL: time To Live (TTL) in seconds. Defines how long content is cached.
	FallbackTTL *scw.Duration `json:"fallback_ttl"`

	// CreatedAt: date the cache stage was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the cache stage was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// BackendStageID: backend stage ID the cache stage is linked to.
	// Precisely one of BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// DNSStage: dns stage.
type DNSStage struct {
	// ID: ID of the DNS stage.
	ID string `json:"id"`

	// Fqdns: list of Fully Qualified Domain Names attached to the stage.
	Fqdns []string `json:"fqdns"`

	// Type: type of the stage.
	// Default value: unknown_type
	Type DNSStageType `json:"type"`

	// PipelineID: pipeline ID the DNS stage belongs to.
	PipelineID *string `json:"pipeline_id"`

	// ProjectID: project ID of the DNS stage.
	ProjectID string `json:"project_id"`

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

	// DNSStageID: DNS stage ID the pipeline is attached to.
	// Precisely one of DNSStageID must be set.
	DNSStageID *string `json:"dns_stage_id,omitempty"`
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
	PipelineID *string `json:"pipeline_id"`

	// ProjectID: project ID of the TLS stage.
	ProjectID string `json:"project_id"`

	// CertificateExpiresAt: expiration date of the certificate.
	CertificateExpiresAt *time.Time `json:"certificate_expires_at"`

	// CreatedAt: date the TLS stage was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the TLS stage was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// CacheStageID: cache stage ID the TLS stage is linked to.
	// Precisely one of CacheStageID, BackendStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// BackendStageID: backend stage ID the TLS stage is linked to.
	// Precisely one of CacheStageID, BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
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
}

// PipelineStages: pipeline stages.
type PipelineStages struct {
	Pipeline *Pipeline `json:"pipeline"`

	DNSStages []*DNSStage `json:"dns_stages"`

	TLSStages []*TLSStage `json:"tls_stages"`

	CacheStages []*CacheStage `json:"cache_stages"`

	BackendStages []*BackendStage `json:"backend_stages"`
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

// TLSSecretsConfig: tls secrets config.
type TLSSecretsConfig struct {
	// TLSSecrets: secret information (from Secret Manager).
	TLSSecrets []*TLSSecret `json:"tls_secrets"`
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
	// ProjectID: project ID in which the backend stage will be created.
	ProjectID string `json:"project_id"`

	// ScalewayS3: scaleway Object Storage origin bucket (S3) linked to the backend stage.
	// Precisely one of ScalewayS3, ScalewayLB must be set.
	ScalewayS3 *ScalewayS3BackendConfig `json:"scaleway_s3,omitempty"`

	// ScalewayLB: scaleway Load Balancer origin linked to the backend stage.
	// Precisely one of ScalewayS3, ScalewayLB must be set.
	ScalewayLB *ScalewayLBBackendConfig `json:"scaleway_lb,omitempty"`
}

// CreateCacheStageRequest: create cache stage request.
type CreateCacheStageRequest struct {
	// ProjectID: project ID in which the cache stage will be created.
	ProjectID string `json:"project_id"`

	// FallbackTTL: time To Live (TTL) in seconds. Defines how long content is cached.
	FallbackTTL *scw.Duration `json:"fallback_ttl,omitempty"`

	// BackendStageID: backend stage ID the cache stage will be linked to.
	// Precisely one of BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// CreateDNSStageRequest: create dns stage request.
type CreateDNSStageRequest struct {
	// ProjectID: project ID in which the DNS stage will be created.
	ProjectID string `json:"project_id"`

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

	// DNSStageID: DNS stage ID the pipeline will be attached to.
	// Precisely one of DNSStageID must be set.
	DNSStageID *string `json:"dns_stage_id,omitempty"`
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

// CreateTLSStageRequest: create tls stage request.
type CreateTLSStageRequest struct {
	// ProjectID: project ID in which the TLS stage will be created.
	ProjectID string `json:"project_id"`

	// Secrets: secret (from Scaleway Secret Manager) containing your custom certificate.
	Secrets []*TLSSecret `json:"secrets"`

	// ManagedCertificate: true when Scaleway generates and manages a Let's Encrypt certificate for the TLS stage/custom endpoint.
	ManagedCertificate *bool `json:"managed_certificate,omitempty"`

	// CacheStageID: cache stage ID the TLS stage will be linked to.
	// Precisely one of CacheStageID, BackendStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// BackendStageID: backend stage ID the TLS stage will be linked to.
	// Precisely one of CacheStageID, BackendStageID must be set.
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

// DeleteTLSStageRequest: delete tls stage request.
type DeleteTLSStageRequest struct {
	// TLSStageID: ID of the TLS stage to delete.
	TLSStageID string `json:"-"`
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

// GetTLSStageRequest: get tls stage request.
type GetTLSStageRequest struct {
	// TLSStageID: ID of the requested TLS stage.
	TLSStageID string `json:"-"`
}

// ListBackendStagesRequest: list backend stages request.
type ListBackendStagesRequest struct {
	// OrderBy: sort order of backend stages in the response.
	// Default value: created_at_asc
	OrderBy ListBackendStagesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of backend stages to return per page.
	PageSize *uint32 `json:"-"`

	// PipelineID: pipeline ID to filter for, only backend stages from this pipeline will be returned.
	PipelineID *string `json:"-"`

	// ProjectID: project ID to filter for, only backend stages from this Project will be returned.
	ProjectID *string `json:"-"`

	// BucketName: bucket name to filter for, only backend stages from this Bucket will be returned.
	BucketName *string `json:"-"`

	// BucketRegion: bucket region to filter for, only backend stages with buckets in this region will be returned.
	BucketRegion *string `json:"-"`

	// LBID: load Balancer ID to filter for, only backend stages with this Load Balancer will be returned.
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
func (r *ListBackendStagesResponse) UnsafeAppend(res interface{}) (uint64, error) {
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
	// OrderBy: sort order of cache stages in the response.
	// Default value: created_at_asc
	OrderBy ListCacheStagesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of cache stages to return per page.
	PageSize *uint32 `json:"-"`

	// PipelineID: pipeline ID to filter for, only cache stages from this pipeline will be returned.
	PipelineID *string `json:"-"`

	// ProjectID: project ID to filter for, only cache stages from this Project will be returned.
	ProjectID *string `json:"-"`
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
func (r *ListCacheStagesResponse) UnsafeAppend(res interface{}) (uint64, error) {
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
	// OrderBy: sort order of DNS stages in the response.
	// Default value: created_at_asc
	OrderBy ListDNSStagesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of DNS stages to return per page.
	PageSize *uint32 `json:"-"`

	// PipelineID: pipeline ID to filter for, only DNS stages from this pipeline will be returned.
	PipelineID *string `json:"-"`

	// ProjectID: project ID to filter for, only DNS stages from this Project will be returned.
	ProjectID *string `json:"-"`

	// Fqdn: fully Qualified Domain Name to filter for (in the format subdomain.example.com), only DNS stages with this FQDN will be returned.
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
func (r *ListDNSStagesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListDNSStagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Stages = append(r.Stages, results.Stages...)
	r.TotalCount += uint64(len(results.Stages))
	return uint64(len(results.Stages)), nil
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

	// Name: pipeline name to filter for, only pipelines with this string within their name will be returned.
	Name *string `json:"-"`

	// OrganizationID: organization ID to filter for, only pipelines from this Organization will be returned.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID to filter for, only pipelines from this Project will be returned.
	ProjectID *string `json:"-"`

	// HasBackendStageLB: filter on backend stage, only pipelines with a Load Balancer origin will be returned.
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
func (r *ListPipelinesResponse) UnsafeAppend(res interface{}) (uint64, error) {
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
func (r *ListPipelinesWithStagesResponse) UnsafeAppend(res interface{}) (uint64, error) {
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
func (r *ListPlansResponse) UnsafeAppend(res interface{}) (uint64, error) {
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

	// OrganizationID: organization ID to filter for, only purge requests from this Project will be returned.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID to filter for, only purge requests from this Project will be returned.
	ProjectID *string `json:"-"`

	// PipelineID: pipeline ID to filter for, only purge requests from this pipeline will be returned.
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
func (r *ListPurgeRequestsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPurgeRequestsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PurgeRequests = append(r.PurgeRequests, results.PurgeRequests...)
	r.TotalCount += uint64(len(results.PurgeRequests))
	return uint64(len(results.PurgeRequests)), nil
}

// ListTLSStagesRequest: list tls stages request.
type ListTLSStagesRequest struct {
	// OrderBy: sort order of TLS stages in the response.
	// Default value: created_at_asc
	OrderBy ListTLSStagesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of TLS stages to return per page.
	PageSize *uint32 `json:"-"`

	// PipelineID: pipeline ID to filter for, only TLS stages from this pipeline will be returned.
	PipelineID *string `json:"-"`

	// ProjectID: project ID to filter for, only TLS stages from this Project will be returned.
	ProjectID *string `json:"-"`

	// SecretID: secret ID to filter for, only TLS stages with this Secret ID will be returned.
	SecretID *string `json:"-"`

	// SecretRegion: secret region to filter for, only TLS stages with a Secret in this region will be returned.
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
func (r *ListTLSStagesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListTLSStagesResponse)
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

// SelectPlanRequest: select plan request.
type SelectPlanRequest struct {
	ProjectID string `json:"project_id"`

	// PlanName: default value: unknown_name
	PlanName PlanName `json:"plan_name"`
}

// UpdateBackendStageRequest: update backend stage request.
type UpdateBackendStageRequest struct {
	// BackendStageID: ID of the backend stage to update.
	BackendStageID string `json:"-"`

	// ScalewayS3: scaleway Object Storage origin bucket (S3) linked to the backend stage.
	// Precisely one of ScalewayS3, ScalewayLB must be set.
	ScalewayS3 *ScalewayS3BackendConfig `json:"scaleway_s3,omitempty"`

	// ScalewayLB: scaleway Load Balancer origin linked to the backend stage.
	// Precisely one of ScalewayS3, ScalewayLB must be set.
	ScalewayLB *ScalewayLBBackendConfig `json:"scaleway_lb,omitempty"`
}

// UpdateCacheStageRequest: update cache stage request.
type UpdateCacheStageRequest struct {
	// CacheStageID: ID of the cache stage to update.
	CacheStageID string `json:"-"`

	// FallbackTTL: time To Live (TTL) in seconds. Defines how long content is cached.
	FallbackTTL *scw.Duration `json:"fallback_ttl,omitempty"`

	// BackendStageID: backend stage ID the cache stage will be linked to.
	// Precisely one of BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
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

	// DNSStageID: DNS stage ID the pipeline will be attached to.
	// Precisely one of DNSStageID must be set.
	DNSStageID *string `json:"dns_stage_id,omitempty"`
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
	// Precisely one of CacheStageID, BackendStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// BackendStageID: backend stage ID the TLS stage will be linked to.
	// Precisely one of CacheStageID, BackendStageID must be set.
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
		Path:   "/edge-services/v1alpha1/pipelines",
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
		Path:   "/edge-services/v1alpha1/pipelines",
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
		Path:   "/edge-services/v1alpha1/pipelines/" + fmt.Sprint(req.PipelineID) + "",
	}

	var resp Pipeline

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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
		Path:   "/edge-services/v1alpha1/pipelines-stages",
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
		Path:   "/edge-services/v1alpha1/pipelines/" + fmt.Sprint(req.PipelineID) + "",
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
		Path:   "/edge-services/v1alpha1/pipelines/" + fmt.Sprint(req.PipelineID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
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
	parameter.AddToQuery(query, "pipeline_id", req.PipelineID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "fqdn", req.Fqdn)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1alpha1/dns-stages",
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

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1alpha1/dns-stages",
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
		Path:   "/edge-services/v1alpha1/dns-stages/" + fmt.Sprint(req.DNSStageID) + "",
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
		Path:   "/edge-services/v1alpha1/dns-stages/" + fmt.Sprint(req.DNSStageID) + "",
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
		Path:   "/edge-services/v1alpha1/dns-stages/" + fmt.Sprint(req.DNSStageID) + "",
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
	parameter.AddToQuery(query, "pipeline_id", req.PipelineID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "secret_id", req.SecretID)
	parameter.AddToQuery(query, "secret_region", req.SecretRegion)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1alpha1/tls-stages",
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

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1alpha1/tls-stages",
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
		Path:   "/edge-services/v1alpha1/tls-stages/" + fmt.Sprint(req.TLSStageID) + "",
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
		Path:   "/edge-services/v1alpha1/tls-stages/" + fmt.Sprint(req.TLSStageID) + "",
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
		Path:   "/edge-services/v1alpha1/tls-stages/" + fmt.Sprint(req.TLSStageID) + "",
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
	parameter.AddToQuery(query, "pipeline_id", req.PipelineID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1alpha1/cache-stages",
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

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1alpha1/cache-stages",
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
		Path:   "/edge-services/v1alpha1/cache-stages/" + fmt.Sprint(req.CacheStageID) + "",
	}

	var resp CacheStage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCacheStage: Update the parameters of an existing cache stage, specified by its `cache_stage_id`. Parameters which can be updated include the `fallback_ttl` and `backend_stage_id`.
func (s *API) UpdateCacheStage(req *UpdateCacheStageRequest, opts ...scw.RequestOption) (*CacheStage, error) {
	var err error

	if fmt.Sprint(req.CacheStageID) == "" {
		return nil, errors.New("field CacheStageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/edge-services/v1alpha1/cache-stages/" + fmt.Sprint(req.CacheStageID) + "",
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
		Path:   "/edge-services/v1alpha1/cache-stages/" + fmt.Sprint(req.CacheStageID) + "",
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
	parameter.AddToQuery(query, "pipeline_id", req.PipelineID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "bucket_name", req.BucketName)
	parameter.AddToQuery(query, "bucket_region", req.BucketRegion)
	parameter.AddToQuery(query, "lb_id", req.LBID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/edge-services/v1alpha1/backend-stages",
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

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1alpha1/backend-stages",
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
		Path:   "/edge-services/v1alpha1/backend-stages/" + fmt.Sprint(req.BackendStageID) + "",
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
		Path:   "/edge-services/v1alpha1/backend-stages/" + fmt.Sprint(req.BackendStageID) + "",
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
		Path:   "/edge-services/v1alpha1/backend-stages/" + fmt.Sprint(req.BackendStageID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
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
		Path:   "/edge-services/v1alpha1/check-domain",
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
		Path:   "/edge-services/v1alpha1/check-pem-chain",
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
		Path:   "/edge-services/v1alpha1/purge-requests",
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
		Path:   "/edge-services/v1alpha1/purge-requests",
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
		Path:   "/edge-services/v1alpha1/purge-requests/" + fmt.Sprint(req.PurgeRequestID) + "",
	}

	var resp PurgeRequest

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckLBOrigin:
func (s *API) CheckLBOrigin(req *CheckLBOriginRequest, opts ...scw.RequestOption) (*CheckLBOriginResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/edge-services/v1alpha1/check-lb-origin",
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
		Path:   "/edge-services/v1alpha1/plans",
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
		Path:   "/edge-services/v1alpha1/current-plan",
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
		Path:   "/edge-services/v1alpha1/current-plan/" + fmt.Sprint(req.ProjectID) + "",
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
		Path:   "/edge-services/v1alpha1/current-plan/" + fmt.Sprint(req.ProjectID) + "",
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
		Path:   "/edge-services/v1alpha1/billing/" + fmt.Sprint(req.ProjectID) + "",
	}

	var resp GetBillingResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

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

type DNSStageType string

const (
	DNSStageTypeUnknownType = DNSStageType("unknown_type")
	DNSStageTypeAuto        = DNSStageType("auto")
	DNSStageTypeManaged     = DNSStageType("managed")
	DNSStageTypeCustom      = DNSStageType("custom")
)

func (enum DNSStageType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
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

type ListBackendStagesRequestOrderBy string

const (
	ListBackendStagesRequestOrderByCreatedAtAsc  = ListBackendStagesRequestOrderBy("created_at_asc")
	ListBackendStagesRequestOrderByCreatedAtDesc = ListBackendStagesRequestOrderBy("created_at_desc")
	ListBackendStagesRequestOrderByNameAsc       = ListBackendStagesRequestOrderBy("name_asc")
	ListBackendStagesRequestOrderByNameDesc      = ListBackendStagesRequestOrderBy("name_desc")
)

func (enum ListBackendStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
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
	ListCacheStagesRequestOrderByCreatedAtAsc  = ListCacheStagesRequestOrderBy("created_at_asc")
	ListCacheStagesRequestOrderByCreatedAtDesc = ListCacheStagesRequestOrderBy("created_at_desc")
	ListCacheStagesRequestOrderByNameAsc       = ListCacheStagesRequestOrderBy("name_asc")
	ListCacheStagesRequestOrderByNameDesc      = ListCacheStagesRequestOrderBy("name_desc")
)

func (enum ListCacheStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
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
	ListDNSStagesRequestOrderByCreatedAtAsc  = ListDNSStagesRequestOrderBy("created_at_asc")
	ListDNSStagesRequestOrderByCreatedAtDesc = ListDNSStagesRequestOrderBy("created_at_desc")
	ListDNSStagesRequestOrderByNameAsc       = ListDNSStagesRequestOrderBy("name_asc")
	ListDNSStagesRequestOrderByNameDesc      = ListDNSStagesRequestOrderBy("name_desc")
)

func (enum ListDNSStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
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

type ListPurgeRequestsRequestOrderBy string

const (
	ListPurgeRequestsRequestOrderByCreatedAtAsc  = ListPurgeRequestsRequestOrderBy("created_at_asc")
	ListPurgeRequestsRequestOrderByCreatedAtDesc = ListPurgeRequestsRequestOrderBy("created_at_desc")
	ListPurgeRequestsRequestOrderByNameAsc       = ListPurgeRequestsRequestOrderBy("name_asc")
	ListPurgeRequestsRequestOrderByNameDesc      = ListPurgeRequestsRequestOrderBy("name_desc")
)

func (enum ListPurgeRequestsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
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
	ListTLSStagesRequestOrderByCreatedAtAsc  = ListTLSStagesRequestOrderBy("created_at_asc")
	ListTLSStagesRequestOrderByCreatedAtDesc = ListTLSStagesRequestOrderBy("created_at_desc")
	ListTLSStagesRequestOrderByNameAsc       = ListTLSStagesRequestOrderBy("name_asc")
	ListTLSStagesRequestOrderByNameDesc      = ListTLSStagesRequestOrderBy("name_desc")
)

func (enum ListTLSStagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
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
	PipelineErrorCodeUnknownCode          = PipelineErrorCode("unknown_code")
	PipelineErrorCodeDNSInvalidFormat     = PipelineErrorCode("dns_invalid_format")
	PipelineErrorCodeDNSInvalidTld        = PipelineErrorCode("dns_invalid_tld")
	PipelineErrorCodeDNSForbiddenScwCloud = PipelineErrorCode("dns_forbidden_scw_cloud")
	PipelineErrorCodeDNSDomainDontExist   = PipelineErrorCode("dns_domain_dont_exist")
	PipelineErrorCodeDNSCnameDontExist    = PipelineErrorCode("dns_cname_dont_exist")
	PipelineErrorCodeDNSCnameResolve      = PipelineErrorCode("dns_cname_resolve")
	PipelineErrorCodeDNSFqdnAlreadyInUse  = PipelineErrorCode("dns_fqdn_already_in_use")
	PipelineErrorCodeTLSCertDeleted       = PipelineErrorCode("tls_cert_deleted")
	PipelineErrorCodeTLSCertExpired       = PipelineErrorCode("tls_cert_expired")
	PipelineErrorCodeTLSCertInvalidFormat = PipelineErrorCode("tls_cert_invalid_format")
	PipelineErrorCodeTLSCertMissing       = PipelineErrorCode("tls_cert_missing")
	PipelineErrorCodeTLSChainOrder        = PipelineErrorCode("tls_chain_order")
	PipelineErrorCodeTLSKeyInvalidFormat  = PipelineErrorCode("tls_key_invalid_format")
	PipelineErrorCodeTLSKeyMissing        = PipelineErrorCode("tls_key_missing")
	PipelineErrorCodeTLSKeyTooMany        = PipelineErrorCode("tls_key_too_many")
	PipelineErrorCodeTLSPairMismatch      = PipelineErrorCode("tls_pair_mismatch")
	PipelineErrorCodeTLSRootInconsistent  = PipelineErrorCode("tls_root_inconsistent")
	PipelineErrorCodeTLSRootIncorrect     = PipelineErrorCode("tls_root_incorrect")
	PipelineErrorCodeTLSRootMissing       = PipelineErrorCode("tls_root_missing")
	PipelineErrorCodeTLSSanMismatch       = PipelineErrorCode("tls_san_mismatch")
	PipelineErrorCodeTLSSelfSigned        = PipelineErrorCode("tls_self_signed")
)

func (enum PipelineErrorCode) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_code"
	}
	return string(enum)
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

type PurgeRequestStatus string

const (
	PurgeRequestStatusUnknownStatus = PurgeRequestStatus("unknown_status")
	PurgeRequestStatusDone          = PurgeRequestStatus("done")
	PurgeRequestStatusError         = PurgeRequestStatus("error")
	PurgeRequestStatusPending       = PurgeRequestStatus("pending")
)

func (enum PurgeRequestStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
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

// ScalewayS3BackendConfig: scaleway s3 backend config.
type ScalewayS3BackendConfig struct {
	BucketName *string `json:"bucket_name"`

	BucketRegion *string `json:"bucket_region"`

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
	SecretID string `json:"secret_id"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// CheckPEMChainRequestSecretChain: check pem chain request secret chain.
type CheckPEMChainRequestSecretChain struct {
	SecretID string `json:"secret_id"`

	SecretRegion string `json:"secret_region"`
}

// BackendStage: backend stage.
type BackendStage struct {
	ID string `json:"id"`

	PipelineID *string `json:"pipeline_id"`

	ProjectID string `json:"project_id"`

	// Precisely one of ScalewayS3Config must be set.
	ScalewayS3Config *ScalewayS3BackendConfig `json:"scaleway_s3_config,omitempty"`
}

// CacheStage: cache stage.
type CacheStage struct {
	ID string `json:"id"`

	PipelineID *string `json:"pipeline_id"`

	ProjectID string `json:"project_id"`

	FallbackTTL *scw.Duration `json:"fallback_ttl"`

	// Precisely one of BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// DNSStage: dns stage.
type DNSStage struct {
	ID string `json:"id"`

	Fqdns []string `json:"fqdns"`

	// Type: default value: unknown_type
	Type DNSStageType `json:"type"`

	PipelineID *string `json:"pipeline_id"`

	ProjectID string `json:"project_id"`

	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	TLSStageID *string `json:"tls_stage_id,omitempty"`

	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// Pipeline: pipeline.
type Pipeline struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	// Status: default value: unknown_status
	Status PipelineStatus `json:"status"`

	Version uint32 `json:"version"`

	Errors []*PipelineError `json:"errors"`

	ProjectID string `json:"project_id"`

	// Precisely one of DNSStageID must be set.
	DNSStageID *string `json:"dns_stage_id,omitempty"`
}

// PurgeRequest: purge request.
type PurgeRequest struct {
	ID string `json:"id"`

	PipelineID string `json:"pipeline_id"`

	// Status: default value: unknown_status
	Status PurgeRequestStatus `json:"status"`

	// Precisely one of Assets, All must be set.
	Assets *[]string `json:"assets,omitempty"`

	// Precisely one of Assets, All must be set.
	All *bool `json:"all,omitempty"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// TLSStage: tls stage.
type TLSStage struct {
	ID string `json:"id"`

	Secrets []*TLSSecret `json:"secrets"`

	PipelineID *string `json:"pipeline_id"`

	ProjectID string `json:"project_id"`

	// Precisely one of CacheStageID, BackendStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// Precisely one of CacheStageID, BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// TLSSecretsConfig: tls secrets config.
type TLSSecretsConfig struct {
	TLSSecrets []*TLSSecret `json:"tls_secrets"`
}

// CheckDomainRequest: check domain request.
type CheckDomainRequest struct {
	Fqdn string `json:"fqdn"`

	Cname string `json:"cname"`
}

// CheckDomainResponse: check domain response.
type CheckDomainResponse struct {
	IsValid bool `json:"is_valid"`
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
	ProjectID string `json:"project_id"`

	// Precisely one of ScalewayS3 must be set.
	ScalewayS3 *ScalewayS3BackendConfig `json:"scaleway_s3,omitempty"`
}

// CreateCacheStageRequest: create cache stage request.
type CreateCacheStageRequest struct {
	ProjectID string `json:"project_id"`

	FallbackTTL *scw.Duration `json:"fallback_ttl,omitempty"`

	// Precisely one of BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// CreateDNSStageRequest: create dns stage request.
type CreateDNSStageRequest struct {
	ProjectID string `json:"project_id"`

	Fqdns *[]string `json:"fqdns,omitempty"`

	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	TLSStageID *string `json:"tls_stage_id,omitempty"`

	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// CreatePipelineRequest: create pipeline request.
type CreatePipelineRequest struct {
	ProjectID string `json:"project_id"`

	Name string `json:"name"`

	Description string `json:"description"`

	// Precisely one of DNSStageID must be set.
	DNSStageID *string `json:"dns_stage_id,omitempty"`
}

// CreatePurgeRequestRequest: create purge request request.
type CreatePurgeRequestRequest struct {
	PipelineID string `json:"pipeline_id"`

	// Precisely one of Assets, All must be set.
	Assets *[]string `json:"assets,omitempty"`

	// Precisely one of Assets, All must be set.
	All *bool `json:"all,omitempty"`
}

// CreateTLSStageRequest: create tls stage request.
type CreateTLSStageRequest struct {
	ProjectID string `json:"project_id"`

	Secrets []*TLSSecret `json:"secrets"`

	// Precisely one of CacheStageID, BackendStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// Precisely one of CacheStageID, BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// DeleteBackendStageRequest: delete backend stage request.
type DeleteBackendStageRequest struct {
	BackendStageID string `json:"-"`
}

// DeleteCacheStageRequest: delete cache stage request.
type DeleteCacheStageRequest struct {
	CacheStageID string `json:"-"`
}

// DeleteDNSStageRequest: delete dns stage request.
type DeleteDNSStageRequest struct {
	DNSStageID string `json:"-"`
}

// DeletePipelineRequest: delete pipeline request.
type DeletePipelineRequest struct {
	PipelineID string `json:"-"`
}

// DeleteTLSStageRequest: delete tls stage request.
type DeleteTLSStageRequest struct {
	TLSStageID string `json:"-"`
}

// GetBackendStageRequest: get backend stage request.
type GetBackendStageRequest struct {
	BackendStageID string `json:"-"`
}

// GetCacheStageRequest: get cache stage request.
type GetCacheStageRequest struct {
	CacheStageID string `json:"-"`
}

// GetDNSStageRequest: get dns stage request.
type GetDNSStageRequest struct {
	DNSStageID string `json:"-"`
}

// GetPipelineRequest: get pipeline request.
type GetPipelineRequest struct {
	PipelineID string `json:"-"`
}

// GetPurgeRequestRequest: get purge request request.
type GetPurgeRequestRequest struct {
	PurgeRequestID string `json:"-"`
}

// GetTLSStageRequest: get tls stage request.
type GetTLSStageRequest struct {
	TLSStageID string `json:"-"`
}

// ListBackendStagesRequest: list backend stages request.
type ListBackendStagesRequest struct {
	// OrderBy: default value: created_at_asc
	OrderBy ListBackendStagesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Name *string `json:"-"`

	PipelineID *string `json:"-"`

	ProjectID *string `json:"-"`

	BucketName *string `json:"-"`

	BucketRegion *string `json:"-"`
}

// ListBackendStagesResponse: list backend stages response.
type ListBackendStagesResponse struct {
	Stages []*BackendStage `json:"stages"`

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
	// OrderBy: default value: created_at_asc
	OrderBy ListCacheStagesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Name *string `json:"-"`

	PipelineID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// ListCacheStagesResponse: list cache stages response.
type ListCacheStagesResponse struct {
	Stages []*CacheStage `json:"stages"`

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
	// OrderBy: default value: created_at_asc
	OrderBy ListDNSStagesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Name *string `json:"-"`

	PipelineID *string `json:"-"`

	ProjectID *string `json:"-"`

	Fqdn *string `json:"-"`
}

// ListDNSStagesResponse: list dns stages response.
type ListDNSStagesResponse struct {
	Stages []*DNSStage `json:"stages"`

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
	// OrderBy: default value: created_at_asc
	OrderBy ListPipelinesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// ListPipelinesResponse: list pipelines response.
type ListPipelinesResponse struct {
	Pipelines []*Pipeline `json:"pipelines"`

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

// ListPurgeRequestsRequest: list purge requests request.
type ListPurgeRequestsRequest struct {
	// OrderBy: default value: created_at_asc
	OrderBy ListPurgeRequestsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`

	PipelineID *string `json:"-"`
}

// ListPurgeRequestsResponse: list purge requests response.
type ListPurgeRequestsResponse struct {
	PurgeRequests []*PurgeRequest `json:"purge_requests"`

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
	// OrderBy: default value: created_at_asc
	OrderBy ListTLSStagesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Name *string `json:"-"`

	PipelineID *string `json:"-"`

	ProjectID *string `json:"-"`

	SecretID *string `json:"-"`

	SecretRegion *string `json:"-"`
}

// ListTLSStagesResponse: list tls stages response.
type ListTLSStagesResponse struct {
	Stages []*TLSStage `json:"stages"`

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

// UpdateBackendStageRequest: update backend stage request.
type UpdateBackendStageRequest struct {
	BackendStageID string `json:"-"`

	// Precisely one of ScalewayS3 must be set.
	ScalewayS3 *ScalewayS3BackendConfig `json:"scaleway_s3,omitempty"`
}

// UpdateCacheStageRequest: update cache stage request.
type UpdateCacheStageRequest struct {
	CacheStageID string `json:"-"`

	FallbackTTL *scw.Duration `json:"fallback_ttl,omitempty"`

	// Precisely one of BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// UpdateDNSStageRequest: update dns stage request.
type UpdateDNSStageRequest struct {
	DNSStageID string `json:"-"`

	Fqdns *[]string `json:"fqdns,omitempty"`

	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	TLSStageID *string `json:"tls_stage_id,omitempty"`

	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

	// Precisely one of TLSStageID, CacheStageID, BackendStageID must be set.
	BackendStageID *string `json:"backend_stage_id,omitempty"`
}

// UpdatePipelineRequest: update pipeline request.
type UpdatePipelineRequest struct {
	PipelineID string `json:"-"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	// Precisely one of DNSStageID must be set.
	DNSStageID *string `json:"dns_stage_id,omitempty"`
}

// UpdateTLSStageRequest: update tls stage request.
type UpdateTLSStageRequest struct {
	TLSStageID string `json:"-"`

	TLSSecretsConfig *TLSSecretsConfig `json:"tls_secrets_config,omitempty"`

	// Precisely one of CacheStageID, BackendStageID must be set.
	CacheStageID *string `json:"cache_stage_id,omitempty"`

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

// ListPipelines:
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

// CreatePipeline:
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

// GetPipeline:
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

// UpdatePipeline:
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

// DeletePipeline:
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

// ListDNSStages:
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
	parameter.AddToQuery(query, "name", req.Name)
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

// CreateDNSStage:
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

// GetDNSStage:
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

// UpdateDNSStage:
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

// DeleteDNSStage:
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

// ListTLSStages:
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
	parameter.AddToQuery(query, "name", req.Name)
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

// CreateTLSStage:
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

// GetTLSStage:
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

// UpdateTLSStage:
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

// DeleteTLSStage:
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

// ListCacheStages:
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
	parameter.AddToQuery(query, "name", req.Name)
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

// CreateCacheStage:
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

// GetCacheStage:
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

// UpdateCacheStage:
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

// DeleteCacheStage:
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

// ListBackendStages:
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "pipeline_id", req.PipelineID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "bucket_name", req.BucketName)
	parameter.AddToQuery(query, "bucket_region", req.BucketRegion)

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

// CreateBackendStage:
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

// GetBackendStage:
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

// UpdateBackendStage:
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

// DeleteBackendStage:
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

// ListPurgeRequests:
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

// CreatePurgeRequest:
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

// GetPurgeRequest:
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

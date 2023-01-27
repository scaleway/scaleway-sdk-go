// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package function provides methods and message types of the function v1beta1 API.
package function

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

// API: serverless functions API
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type CronStatus string

const (
	// CronStatusUnknown is [insert doc].
	CronStatusUnknown = CronStatus("unknown")
	// CronStatusReady is [insert doc].
	CronStatusReady = CronStatus("ready")
	// CronStatusDeleting is [insert doc].
	CronStatusDeleting = CronStatus("deleting")
	// CronStatusError is [insert doc].
	CronStatusError = CronStatus("error")
	// CronStatusLocked is [insert doc].
	CronStatusLocked = CronStatus("locked")
	// CronStatusCreating is [insert doc].
	CronStatusCreating = CronStatus("creating")
	// CronStatusPending is [insert doc].
	CronStatusPending = CronStatus("pending")
)

func (enum CronStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum CronStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CronStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CronStatus(CronStatus(tmp).String())
	return nil
}

type DomainStatus string

const (
	// DomainStatusUnknown is [insert doc].
	DomainStatusUnknown = DomainStatus("unknown")
	// DomainStatusReady is [insert doc].
	DomainStatusReady = DomainStatus("ready")
	// DomainStatusDeleting is [insert doc].
	DomainStatusDeleting = DomainStatus("deleting")
	// DomainStatusError is [insert doc].
	DomainStatusError = DomainStatus("error")
	// DomainStatusCreating is [insert doc].
	DomainStatusCreating = DomainStatus("creating")
	// DomainStatusPending is [insert doc].
	DomainStatusPending = DomainStatus("pending")
)

func (enum DomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DomainStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainStatus(DomainStatus(tmp).String())
	return nil
}

type FunctionHTTPOption string

const (
	// FunctionHTTPOptionUnknownHTTPOption is [insert doc].
	FunctionHTTPOptionUnknownHTTPOption = FunctionHTTPOption("unknown_http_option")
	// FunctionHTTPOptionEnabled is [insert doc].
	FunctionHTTPOptionEnabled = FunctionHTTPOption("enabled")
	// FunctionHTTPOptionRedirected is [insert doc].
	FunctionHTTPOptionRedirected = FunctionHTTPOption("redirected")
)

func (enum FunctionHTTPOption) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_http_option"
	}
	return string(enum)
}

func (enum FunctionHTTPOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionHTTPOption) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionHTTPOption(FunctionHTTPOption(tmp).String())
	return nil
}

type FunctionPrivacy string

const (
	// FunctionPrivacyUnknownPrivacy is [insert doc].
	FunctionPrivacyUnknownPrivacy = FunctionPrivacy("unknown_privacy")
	// FunctionPrivacyPublic is [insert doc].
	FunctionPrivacyPublic = FunctionPrivacy("public")
	// FunctionPrivacyPrivate is [insert doc].
	FunctionPrivacyPrivate = FunctionPrivacy("private")
)

func (enum FunctionPrivacy) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_privacy"
	}
	return string(enum)
}

func (enum FunctionPrivacy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionPrivacy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionPrivacy(FunctionPrivacy(tmp).String())
	return nil
}

type FunctionRuntime string

const (
	// FunctionRuntimeUnknownRuntime is [insert doc].
	FunctionRuntimeUnknownRuntime = FunctionRuntime("unknown_runtime")
	// FunctionRuntimeGolang is [insert doc].
	FunctionRuntimeGolang = FunctionRuntime("golang")
	// FunctionRuntimePython is [insert doc].
	FunctionRuntimePython = FunctionRuntime("python")
	// FunctionRuntimePython3 is [insert doc].
	FunctionRuntimePython3 = FunctionRuntime("python3")
	// FunctionRuntimeNode8 is [insert doc].
	FunctionRuntimeNode8 = FunctionRuntime("node8")
	// FunctionRuntimeNode10 is [insert doc].
	FunctionRuntimeNode10 = FunctionRuntime("node10")
	// FunctionRuntimeNode14 is [insert doc].
	FunctionRuntimeNode14 = FunctionRuntime("node14")
	// FunctionRuntimeNode16 is [insert doc].
	FunctionRuntimeNode16 = FunctionRuntime("node16")
	// FunctionRuntimeNode17 is [insert doc].
	FunctionRuntimeNode17 = FunctionRuntime("node17")
	// FunctionRuntimePython37 is [insert doc].
	FunctionRuntimePython37 = FunctionRuntime("python37")
	// FunctionRuntimePython38 is [insert doc].
	FunctionRuntimePython38 = FunctionRuntime("python38")
	// FunctionRuntimePython39 is [insert doc].
	FunctionRuntimePython39 = FunctionRuntime("python39")
	// FunctionRuntimePython310 is [insert doc].
	FunctionRuntimePython310 = FunctionRuntime("python310")
	// FunctionRuntimeGo113 is [insert doc].
	FunctionRuntimeGo113 = FunctionRuntime("go113")
	// FunctionRuntimeGo117 is [insert doc].
	FunctionRuntimeGo117 = FunctionRuntime("go117")
	// FunctionRuntimeGo118 is [insert doc].
	FunctionRuntimeGo118 = FunctionRuntime("go118")
	// FunctionRuntimeNode18 is [insert doc].
	FunctionRuntimeNode18 = FunctionRuntime("node18")
	// FunctionRuntimeRust165 is [insert doc].
	FunctionRuntimeRust165 = FunctionRuntime("rust165")
	// FunctionRuntimeGo119 is [insert doc].
	FunctionRuntimeGo119 = FunctionRuntime("go119")
	// FunctionRuntimePython311 is [insert doc].
	FunctionRuntimePython311 = FunctionRuntime("python311")
	// FunctionRuntimePhp82 is [insert doc].
	FunctionRuntimePhp82 = FunctionRuntime("php82")
)

func (enum FunctionRuntime) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_runtime"
	}
	return string(enum)
}

func (enum FunctionRuntime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionRuntime) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionRuntime(FunctionRuntime(tmp).String())
	return nil
}

type FunctionStatus string

const (
	// FunctionStatusUnknown is [insert doc].
	FunctionStatusUnknown = FunctionStatus("unknown")
	// FunctionStatusReady is [insert doc].
	FunctionStatusReady = FunctionStatus("ready")
	// FunctionStatusDeleting is [insert doc].
	FunctionStatusDeleting = FunctionStatus("deleting")
	// FunctionStatusError is [insert doc].
	FunctionStatusError = FunctionStatus("error")
	// FunctionStatusLocked is [insert doc].
	FunctionStatusLocked = FunctionStatus("locked")
	// FunctionStatusCreating is [insert doc].
	FunctionStatusCreating = FunctionStatus("creating")
	// FunctionStatusPending is [insert doc].
	FunctionStatusPending = FunctionStatus("pending")
	// FunctionStatusCreated is [insert doc].
	FunctionStatusCreated = FunctionStatus("created")
)

func (enum FunctionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum FunctionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionStatus(FunctionStatus(tmp).String())
	return nil
}

type ListCronsRequestOrderBy string

const (
	// ListCronsRequestOrderByCreatedAtAsc is [insert doc].
	ListCronsRequestOrderByCreatedAtAsc = ListCronsRequestOrderBy("created_at_asc")
	// ListCronsRequestOrderByCreatedAtDesc is [insert doc].
	ListCronsRequestOrderByCreatedAtDesc = ListCronsRequestOrderBy("created_at_desc")
)

func (enum ListCronsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListCronsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCronsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCronsRequestOrderBy(ListCronsRequestOrderBy(tmp).String())
	return nil
}

type ListDomainsRequestOrderBy string

const (
	// ListDomainsRequestOrderByCreatedAtAsc is [insert doc].
	ListDomainsRequestOrderByCreatedAtAsc = ListDomainsRequestOrderBy("created_at_asc")
	// ListDomainsRequestOrderByCreatedAtDesc is [insert doc].
	ListDomainsRequestOrderByCreatedAtDesc = ListDomainsRequestOrderBy("created_at_desc")
	// ListDomainsRequestOrderByHostnameAsc is [insert doc].
	ListDomainsRequestOrderByHostnameAsc = ListDomainsRequestOrderBy("hostname_asc")
	// ListDomainsRequestOrderByHostnameDesc is [insert doc].
	ListDomainsRequestOrderByHostnameDesc = ListDomainsRequestOrderBy("hostname_desc")
)

func (enum ListDomainsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListDomainsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDomainsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDomainsRequestOrderBy(ListDomainsRequestOrderBy(tmp).String())
	return nil
}

type ListFunctionsRequestOrderBy string

const (
	// ListFunctionsRequestOrderByCreatedAtAsc is [insert doc].
	ListFunctionsRequestOrderByCreatedAtAsc = ListFunctionsRequestOrderBy("created_at_asc")
	// ListFunctionsRequestOrderByCreatedAtDesc is [insert doc].
	ListFunctionsRequestOrderByCreatedAtDesc = ListFunctionsRequestOrderBy("created_at_desc")
	// ListFunctionsRequestOrderByNameAsc is [insert doc].
	ListFunctionsRequestOrderByNameAsc = ListFunctionsRequestOrderBy("name_asc")
	// ListFunctionsRequestOrderByNameDesc is [insert doc].
	ListFunctionsRequestOrderByNameDesc = ListFunctionsRequestOrderBy("name_desc")
)

func (enum ListFunctionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListFunctionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFunctionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFunctionsRequestOrderBy(ListFunctionsRequestOrderBy(tmp).String())
	return nil
}

type ListLogsRequestOrderBy string

const (
	// ListLogsRequestOrderByTimestampDesc is [insert doc].
	ListLogsRequestOrderByTimestampDesc = ListLogsRequestOrderBy("timestamp_desc")
	// ListLogsRequestOrderByTimestampAsc is [insert doc].
	ListLogsRequestOrderByTimestampAsc = ListLogsRequestOrderBy("timestamp_asc")
)

func (enum ListLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "timestamp_desc"
	}
	return string(enum)
}

func (enum ListLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListLogsRequestOrderBy(ListLogsRequestOrderBy(tmp).String())
	return nil
}

type ListNamespacesRequestOrderBy string

const (
	// ListNamespacesRequestOrderByCreatedAtAsc is [insert doc].
	ListNamespacesRequestOrderByCreatedAtAsc = ListNamespacesRequestOrderBy("created_at_asc")
	// ListNamespacesRequestOrderByCreatedAtDesc is [insert doc].
	ListNamespacesRequestOrderByCreatedAtDesc = ListNamespacesRequestOrderBy("created_at_desc")
	// ListNamespacesRequestOrderByNameAsc is [insert doc].
	ListNamespacesRequestOrderByNameAsc = ListNamespacesRequestOrderBy("name_asc")
	// ListNamespacesRequestOrderByNameDesc is [insert doc].
	ListNamespacesRequestOrderByNameDesc = ListNamespacesRequestOrderBy("name_desc")
)

func (enum ListNamespacesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNamespacesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNamespacesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNamespacesRequestOrderBy(ListNamespacesRequestOrderBy(tmp).String())
	return nil
}

type ListTokensRequestOrderBy string

const (
	// ListTokensRequestOrderByCreatedAtAsc is [insert doc].
	ListTokensRequestOrderByCreatedAtAsc = ListTokensRequestOrderBy("created_at_asc")
	// ListTokensRequestOrderByCreatedAtDesc is [insert doc].
	ListTokensRequestOrderByCreatedAtDesc = ListTokensRequestOrderBy("created_at_desc")
)

func (enum ListTokensRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListTokensRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTokensRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTokensRequestOrderBy(ListTokensRequestOrderBy(tmp).String())
	return nil
}

type ListTriggerInputsRequestOrderBy string

const (
	// ListTriggerInputsRequestOrderByCreatedAtAsc is [insert doc].
	ListTriggerInputsRequestOrderByCreatedAtAsc = ListTriggerInputsRequestOrderBy("created_at_asc")
	// ListTriggerInputsRequestOrderByCreatedAtDesc is [insert doc].
	ListTriggerInputsRequestOrderByCreatedAtDesc = ListTriggerInputsRequestOrderBy("created_at_desc")
)

func (enum ListTriggerInputsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListTriggerInputsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTriggerInputsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTriggerInputsRequestOrderBy(ListTriggerInputsRequestOrderBy(tmp).String())
	return nil
}

type ListTriggersRequestOrderBy string

const (
	// ListTriggersRequestOrderByCreatedAtAsc is [insert doc].
	ListTriggersRequestOrderByCreatedAtAsc = ListTriggersRequestOrderBy("created_at_asc")
	// ListTriggersRequestOrderByCreatedAtDesc is [insert doc].
	ListTriggersRequestOrderByCreatedAtDesc = ListTriggersRequestOrderBy("created_at_desc")
)

func (enum ListTriggersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListTriggersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTriggersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTriggersRequestOrderBy(ListTriggersRequestOrderBy(tmp).String())
	return nil
}

type LogStream string

const (
	// LogStreamUnknown is [insert doc].
	LogStreamUnknown = LogStream("unknown")
	// LogStreamStdout is [insert doc].
	LogStreamStdout = LogStream("stdout")
	// LogStreamStderr is [insert doc].
	LogStreamStderr = LogStream("stderr")
)

func (enum LogStream) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum LogStream) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LogStream) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LogStream(LogStream(tmp).String())
	return nil
}

type NamespaceStatus string

const (
	// NamespaceStatusUnknown is [insert doc].
	NamespaceStatusUnknown = NamespaceStatus("unknown")
	// NamespaceStatusReady is [insert doc].
	NamespaceStatusReady = NamespaceStatus("ready")
	// NamespaceStatusDeleting is [insert doc].
	NamespaceStatusDeleting = NamespaceStatus("deleting")
	// NamespaceStatusError is [insert doc].
	NamespaceStatusError = NamespaceStatus("error")
	// NamespaceStatusLocked is [insert doc].
	NamespaceStatusLocked = NamespaceStatus("locked")
	// NamespaceStatusCreating is [insert doc].
	NamespaceStatusCreating = NamespaceStatus("creating")
	// NamespaceStatusPending is [insert doc].
	NamespaceStatusPending = NamespaceStatus("pending")
)

func (enum NamespaceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NamespaceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NamespaceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NamespaceStatus(NamespaceStatus(tmp).String())
	return nil
}

type NullValue string

const (
	// NullValueNULLVALUE is [insert doc].
	NullValueNULLVALUE = NullValue("NULL_VALUE")
)

func (enum NullValue) String() string {
	if enum == "" {
		// return default value if empty
		return "NULL_VALUE"
	}
	return string(enum)
}

func (enum NullValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NullValue) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NullValue(NullValue(tmp).String())
	return nil
}

type RuntimeStatus string

const (
	// RuntimeStatusUnknownStatus is [insert doc].
	RuntimeStatusUnknownStatus = RuntimeStatus("unknown_status")
	// RuntimeStatusBeta is [insert doc].
	RuntimeStatusBeta = RuntimeStatus("beta")
	// RuntimeStatusAvailable is [insert doc].
	RuntimeStatusAvailable = RuntimeStatus("available")
	// RuntimeStatusDeprecated is [insert doc].
	RuntimeStatusDeprecated = RuntimeStatus("deprecated")
	// RuntimeStatusEndOfSupport is [insert doc].
	RuntimeStatusEndOfSupport = RuntimeStatus("end_of_support")
	// RuntimeStatusEndOfLife is [insert doc].
	RuntimeStatusEndOfLife = RuntimeStatus("end_of_life")
)

func (enum RuntimeStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum RuntimeStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RuntimeStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RuntimeStatus(RuntimeStatus(tmp).String())
	return nil
}

type TokenStatus string

const (
	// TokenStatusUnknown is [insert doc].
	TokenStatusUnknown = TokenStatus("unknown")
	// TokenStatusReady is [insert doc].
	TokenStatusReady = TokenStatus("ready")
	// TokenStatusDeleting is [insert doc].
	TokenStatusDeleting = TokenStatus("deleting")
	// TokenStatusError is [insert doc].
	TokenStatusError = TokenStatus("error")
	// TokenStatusCreating is [insert doc].
	TokenStatusCreating = TokenStatus("creating")
)

func (enum TokenStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum TokenStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TokenStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TokenStatus(TokenStatus(tmp).String())
	return nil
}

type TriggerInputStatus string

const (
	// TriggerInputStatusUnknown is [insert doc].
	TriggerInputStatusUnknown = TriggerInputStatus("unknown")
	// TriggerInputStatusReady is [insert doc].
	TriggerInputStatusReady = TriggerInputStatus("ready")
	// TriggerInputStatusDeleting is [insert doc].
	TriggerInputStatusDeleting = TriggerInputStatus("deleting")
	// TriggerInputStatusError is [insert doc].
	TriggerInputStatusError = TriggerInputStatus("error")
	// TriggerInputStatusCreating is [insert doc].
	TriggerInputStatusCreating = TriggerInputStatus("creating")
	// TriggerInputStatusPending is [insert doc].
	TriggerInputStatusPending = TriggerInputStatus("pending")
)

func (enum TriggerInputStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum TriggerInputStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TriggerInputStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TriggerInputStatus(TriggerInputStatus(tmp).String())
	return nil
}

type TriggerStatus string

const (
	// TriggerStatusUnknownStatus is [insert doc].
	TriggerStatusUnknownStatus = TriggerStatus("unknown_status")
	// TriggerStatusReady is [insert doc].
	TriggerStatusReady = TriggerStatus("ready")
	// TriggerStatusDeleting is [insert doc].
	TriggerStatusDeleting = TriggerStatus("deleting")
	// TriggerStatusError is [insert doc].
	TriggerStatusError = TriggerStatus("error")
	// TriggerStatusCreating is [insert doc].
	TriggerStatusCreating = TriggerStatus("creating")
	// TriggerStatusPending is [insert doc].
	TriggerStatusPending = TriggerStatus("pending")
)

func (enum TriggerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum TriggerStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TriggerStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TriggerStatus(TriggerStatus(tmp).String())
	return nil
}

type TriggerType string

const (
	// TriggerTypeUnknownTriggerType is [insert doc].
	TriggerTypeUnknownTriggerType = TriggerType("unknown_trigger_type")
	// TriggerTypeNats is [insert doc].
	TriggerTypeNats = TriggerType("nats")
	// TriggerTypeSqs is [insert doc].
	TriggerTypeSqs = TriggerType("sqs")
)

func (enum TriggerType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_trigger_type"
	}
	return string(enum)
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

type CreateTriggerInputRequestNatsClientConfigSpec struct {
	Subject string `json:"subject"`
}

type CreateTriggerInputRequestSqsClientConfigSpec struct {
	Queue string `json:"queue"`
}

type CreateTriggerRequestNatsFailureHandlingPolicy struct {
	RetryPolicy *CreateTriggerRequestNatsFailureHandlingPolicyRetryPolicy `json:"retry_policy"`

	// Precisely one of NatsDeadLetter, SqsDeadLetter must be set.
	NatsDeadLetter *CreateTriggerRequestNatsFailureHandlingPolicyNatsDeadLetter `json:"nats_dead_letter,omitempty"`

	// Precisely one of NatsDeadLetter, SqsDeadLetter must be set.
	SqsDeadLetter *CreateTriggerRequestNatsFailureHandlingPolicySqsDeadLetter `json:"sqs_dead_letter,omitempty"`
}

type CreateTriggerRequestNatsFailureHandlingPolicyNatsDeadLetter struct {
	MnqNamespaceID *string `json:"mnq_namespace_id"`

	Subject *string `json:"subject"`
}

type CreateTriggerRequestNatsFailureHandlingPolicyRetryPolicy struct {
	MaxRetries *uint32 `json:"max_retries"`

	RetryPeriod *scw.Duration `json:"retry_period"`
}

type CreateTriggerRequestNatsFailureHandlingPolicySqsDeadLetter struct {
	MnqNamespaceID *string `json:"mnq_namespace_id"`

	Queue *string `json:"queue"`
}

type CreateTriggerRequestSqsFailureHandlingPolicy struct {
}

// Cron: cron
type Cron struct {
	ID string `json:"id"`

	FunctionID string `json:"function_id"`

	Schedule string `json:"schedule"`

	Args *scw.JSONObject `json:"args"`
	// Status:
	//
	// Default value: unknown
	Status CronStatus `json:"status"`

	Name string `json:"name"`
}

// Domain: domain
type Domain struct {
	ID string `json:"id"`

	Hostname string `json:"hostname"`

	FunctionID string `json:"function_id"`

	URL string `json:"url"`
	// Status:
	//
	// Default value: unknown
	Status DomainStatus `json:"status"`

	ErrorMessage *string `json:"error_message"`
}

type DownloadURL struct {
	URL string `json:"url"`

	Headers map[string]*[]string `json:"headers"`
}

// Function: function
type Function struct {
	ID string `json:"id"`

	Name string `json:"name"`

	NamespaceID string `json:"namespace_id"`
	// Status:
	//
	// Default value: unknown
	Status FunctionStatus `json:"status"`

	EnvironmentVariables map[string]string `json:"environment_variables"`

	MinScale uint32 `json:"min_scale"`

	MaxScale uint32 `json:"max_scale"`
	// Runtime:
	//
	// Default value: unknown_runtime
	Runtime FunctionRuntime `json:"runtime"`

	MemoryLimit uint32 `json:"memory_limit"`

	CPULimit uint32 `json:"cpu_limit"`

	Timeout *scw.Duration `json:"timeout"`

	Handler string `json:"handler"`

	ErrorMessage *string `json:"error_message"`
	// Privacy:
	//
	// Default value: unknown_privacy
	Privacy FunctionPrivacy `json:"privacy"`

	Description *string `json:"description"`

	DomainName string `json:"domain_name"`

	SecretEnvironmentVariables []*SecretHashedValue `json:"secret_environment_variables"`

	Region scw.Region `json:"region"`
	// HTTPOption: configure how HTTP and HTTPS requests are handled
	//
	// possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	//
	// Default value: enabled
	HTTPOption FunctionHTTPOption `json:"http_option"`

	RuntimeMessage string `json:"runtime_message"`
}

// ListCronsResponse: list crons response
type ListCronsResponse struct {
	Crons []*Cron `json:"crons"`

	TotalCount uint32 `json:"total_count"`
}

// ListDomainsResponse: list domains response
type ListDomainsResponse struct {
	Domains []*Domain `json:"domains"`

	TotalCount uint32 `json:"total_count"`
}

// ListFunctionRuntimesResponse: list function runtimes response
type ListFunctionRuntimesResponse struct {
	Runtimes []*Runtime `json:"runtimes"`

	TotalCount uint32 `json:"total_count"`
}

// ListFunctionsResponse: list functions response
type ListFunctionsResponse struct {
	Functions []*Function `json:"functions"`

	TotalCount uint32 `json:"total_count"`
}

// ListLogsResponse: list logs response
type ListLogsResponse struct {
	Logs []*Log `json:"logs"`

	TotalCount uint32 `json:"total_count"`
}

// ListNamespacesResponse: list namespaces response
type ListNamespacesResponse struct {
	Namespaces []*Namespace `json:"namespaces"`

	TotalCount uint32 `json:"total_count"`
}

type ListTokensResponse struct {
	Tokens []*Token `json:"tokens"`

	TotalCount uint32 `json:"total_count"`
}

type ListTriggerInputsResponse struct {
	Inputs []*TriggerInput `json:"inputs"`

	TotalCount uint32 `json:"total_count"`
}

type ListTriggersResponse struct {
	Triggers []*Trigger `json:"triggers"`

	TotalCount uint32 `json:"total_count"`
}

// Log: log
type Log struct {
	Message string `json:"message"`

	Timestamp *time.Time `json:"timestamp"`

	ID string `json:"id"`
	// Level: contains the severity of the log (info, debug, error, ...)
	Level string `json:"level"`
	// Source: source of the log (core runtime or user code)
	Source string `json:"source"`
	// Stream: can be stdout or stderr
	//
	// Default value: unknown
	Stream LogStream `json:"stream"`
}

// Namespace: namespace
type Namespace struct {
	ID string `json:"id"`

	Name string `json:"name"`

	EnvironmentVariables map[string]string `json:"environment_variables"`

	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`
	// Status:
	//
	// Default value: unknown
	Status NamespaceStatus `json:"status"`

	RegistryNamespaceID string `json:"registry_namespace_id"`

	ErrorMessage *string `json:"error_message"`

	RegistryEndpoint string `json:"registry_endpoint"`

	Description *string `json:"description"`

	SecretEnvironmentVariables []*SecretHashedValue `json:"secret_environment_variables"`

	Region scw.Region `json:"region"`
}

type Runtime struct {
	Name string `json:"name"`

	Language string `json:"language"`

	Version string `json:"version"`

	DefaultHandler string `json:"default_handler"`

	CodeSample string `json:"code_sample"`
	// Status:
	//
	// Default value: unknown_status
	Status RuntimeStatus `json:"status"`

	StatusMessage string `json:"status_message"`

	Extension string `json:"extension"`

	Implementation string `json:"implementation"`
}

type Secret struct {
	Key string `json:"key"`

	Value *string `json:"value"`
}

type SecretHashedValue struct {
	Key string `json:"key"`

	HashedValue string `json:"hashed_value"`
}

type SetTriggerInputsRequestNatsConfigs struct {
	Configs []*CreateTriggerInputRequestNatsClientConfigSpec `json:"configs"`
}

type SetTriggerInputsRequestSqsConfigs struct {
	Configs []*CreateTriggerInputRequestSqsClientConfigSpec `json:"configs"`
}

type SetTriggerInputsResponse struct {
	TriggerInputs []*TriggerInput `json:"trigger_inputs"`
}

// Token: token
type Token struct {
	ID string `json:"id"`

	Token string `json:"token"`

	// Precisely one of FunctionID, NamespaceID must be set.
	FunctionID *string `json:"function_id,omitempty"`

	// Precisely one of FunctionID, NamespaceID must be set.
	NamespaceID *string `json:"namespace_id,omitempty"`
	// Deprecated
	PublicKey *string `json:"public_key,omitempty"`
	// Status:
	//
	// Default value: unknown
	Status TokenStatus `json:"status"`

	Description *string `json:"description"`

	ExpiresAt *time.Time `json:"expires_at"`
}

type Trigger struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`
	// Type:
	//
	// Default value: unknown_trigger_type
	Type TriggerType `json:"type"`
	// Status:
	//
	// Default value: unknown_status
	Status TriggerStatus `json:"status"`

	ErrorMessage *string `json:"error_message"`

	FunctionID string `json:"function_id"`

	// Precisely one of NatsFailureHandlingPolicy, SqsFailureHandlingPolicy must be set.
	NatsFailureHandlingPolicy *TriggerNatsFailureHandlingPolicy `json:"nats_failure_handling_policy,omitempty"`

	// Precisely one of NatsFailureHandlingPolicy, SqsFailureHandlingPolicy must be set.
	SqsFailureHandlingPolicy *TriggerSqsFailureHandlingPolicy `json:"sqs_failure_handling_policy,omitempty"`
}

type TriggerInput struct {
	ID string `json:"id"`

	MnqNamespaceID *string `json:"mnq_namespace_id"`
	// Status:
	//
	// Default value: unknown
	Status TriggerInputStatus `json:"status"`

	ErrorMessage *string `json:"error_message"`

	// Precisely one of NatsConfig, SqsConfig must be set.
	NatsConfig *TriggerInputNatsClientConfig `json:"nats_config,omitempty"`

	// Precisely one of NatsConfig, SqsConfig must be set.
	SqsConfig *TriggerInputSqsClientConfig `json:"sqs_config,omitempty"`
}

type TriggerInputNatsClientConfig struct {
	Subject string `json:"subject"`
}

type TriggerInputSqsClientConfig struct {
	Queue string `json:"queue"`
}

type TriggerNatsDeadLetter struct {
	MnqNamespaceID string `json:"mnq_namespace_id"`

	Subject string `json:"subject"`
}

type TriggerNatsFailureHandlingPolicy struct {
	RetryPolicy *TriggerRetryPolicy `json:"retry_policy"`

	// Precisely one of NatsDeadLetter, SqsDeadLetter must be set.
	NatsDeadLetter *TriggerNatsDeadLetter `json:"nats_dead_letter,omitempty"`

	// Precisely one of NatsDeadLetter, SqsDeadLetter must be set.
	SqsDeadLetter *TriggerSqsDeadLetter `json:"sqs_dead_letter,omitempty"`
}

type TriggerRetryPolicy struct {
	MaxRetries uint32 `json:"max_retries"`

	RetryPeriod *scw.Duration `json:"retry_period"`
}

type TriggerSqsDeadLetter struct {
	MnqNamespaceID string `json:"mnq_namespace_id"`

	Queue string `json:"queue"`
}

type TriggerSqsFailureHandlingPolicy struct {
}

type UpdateTriggerInputRequestNatsClientConfigSpec struct {
	Subject *string `json:"subject"`
}

type UpdateTriggerInputRequestSqsClientConfigSpec struct {
	Queue *string `json:"queue"`
}

// UploadURL: upload url
type UploadURL struct {
	URL string `json:"url"`

	Headers map[string]*[]string `json:"headers"`
}

// Service API

// Regions list localities the api is available in
func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

type ListNamespacesRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListNamespacesRequestOrderBy `json:"-"`

	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// ListNamespaces: list all your namespaces
func (s *API) ListNamespaces(req *ListNamespacesRequest, opts ...scw.RequestOption) (*ListNamespacesResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListNamespacesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetNamespaceRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	NamespaceID string `json:"-"`
}

// GetNamespace: get a namespace
//
// Get the namespace associated with the given id.
func (s *API) GetNamespace(req *GetNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
		Headers: http.Header{},
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateNamespaceRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	Name string `json:"name"`

	EnvironmentVariables *map[string]string `json:"environment_variables"`

	ProjectID string `json:"project_id"`

	Description *string `json:"description"`

	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
}

// CreateNamespace: create a new namespace
func (s *API) CreateNamespace(req *CreateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("ns")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateNamespaceRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	NamespaceID string `json:"-"`

	EnvironmentVariables *map[string]string `json:"environment_variables"`

	Description *string `json:"description"`

	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
}

// UpdateNamespace: update an existing namespace
//
// Update the space associated with the given id.
func (s *API) UpdateNamespace(req *UpdateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteNamespaceRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	NamespaceID string `json:"-"`
}

// DeleteNamespace: delete an existing namespace
//
// Delete the namespace associated with the given id.
func (s *API) DeleteNamespace(req *DeleteNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
		Headers: http.Header{},
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListFunctionsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListFunctionsRequestOrderBy `json:"-"`

	NamespaceID string `json:"-"`

	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// ListFunctions: list all your functions
func (s *API) ListFunctions(req *ListFunctionsRequest, opts ...scw.RequestOption) (*ListFunctionsResponse, error) {
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
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListFunctionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetFunctionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`
}

// GetFunction: get a function
//
// Get the function associated with the given id.
func (s *API) GetFunction(req *GetFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
		Headers: http.Header{},
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateFunctionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	Name string `json:"name"`

	NamespaceID string `json:"namespace_id"`

	EnvironmentVariables *map[string]string `json:"environment_variables"`

	MinScale *uint32 `json:"min_scale"`

	MaxScale *uint32 `json:"max_scale"`
	// Runtime:
	//
	// Default value: unknown_runtime
	Runtime FunctionRuntime `json:"runtime"`

	MemoryLimit *uint32 `json:"memory_limit"`

	Timeout *scw.Duration `json:"timeout"`

	Handler *string `json:"handler"`
	// Privacy:
	//
	// Default value: unknown_privacy
	Privacy FunctionPrivacy `json:"privacy"`

	Description *string `json:"description"`

	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
	// HTTPOption: configure how HTTP and HTTPS requests are handled
	//
	// possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	//
	// Default value: enabled
	HTTPOption FunctionHTTPOption `json:"http_option"`
}

// CreateFunction: create a new function
func (s *API) CreateFunction(req *CreateFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("fn")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateFunctionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`

	EnvironmentVariables *map[string]string `json:"environment_variables"`

	MinScale *uint32 `json:"min_scale"`

	MaxScale *uint32 `json:"max_scale"`
	// Runtime:
	//
	// Default value: unknown_runtime
	Runtime FunctionRuntime `json:"runtime"`

	MemoryLimit *uint32 `json:"memory_limit"`

	Timeout *scw.Duration `json:"timeout"`

	Redeploy *bool `json:"redeploy"`

	Handler *string `json:"handler"`
	// Privacy:
	//
	// Default value: unknown_privacy
	Privacy FunctionPrivacy `json:"privacy"`

	Description *string `json:"description"`

	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`
	// HTTPOption: configure how HTTP and HTTPS requests are handled
	//
	// possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	//
	// Default value: enabled
	HTTPOption FunctionHTTPOption `json:"http_option"`
}

// UpdateFunction: update an existing function
//
// Update the function associated with the given id.
func (s *API) UpdateFunction(req *UpdateFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteFunctionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`
}

// DeleteFunction: delete a function
//
// Delete the function associated with the given id.
func (s *API) DeleteFunction(req *DeleteFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
		Headers: http.Header{},
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeployFunctionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`
}

// DeployFunction: deploy a function
//
// Deploy a function associated with the given id.
func (s *API) DeployFunction(req *DeployFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/deploy",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListFunctionRuntimesRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
}

// ListFunctionRuntimes: list function runtimes
//
// List available function runtimes.
func (s *API) ListFunctionRuntimes(req *ListFunctionRuntimesRequest, opts ...scw.RequestOption) (*ListFunctionRuntimesResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/runtimes",
		Headers: http.Header{},
	}

	var resp ListFunctionRuntimesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetFunctionUploadURLRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`

	ContentLength uint64 `json:"-"`
}

// GetFunctionUploadURL: get an upload URL of a function
//
// Get an upload URL of a function associated with the given id.
func (s *API) GetFunctionUploadURL(req *GetFunctionUploadURLRequest, opts ...scw.RequestOption) (*UploadURL, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "content_length", req.ContentLength)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/upload-url",
		Query:   query,
		Headers: http.Header{},
	}

	var resp UploadURL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetFunctionDownloadURLRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`
}

// GetFunctionDownloadURL: get a download URL of a function
//
// Get a download URL for a function associated with the given id.
func (s *API) GetFunctionDownloadURL(req *GetFunctionDownloadURLRequest, opts ...scw.RequestOption) (*DownloadURL, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/download-url",
		Headers: http.Header{},
	}

	var resp DownloadURL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListCronsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListCronsRequestOrderBy `json:"-"`

	FunctionID string `json:"-"`
}

// ListCrons: list all your crons
func (s *API) ListCrons(req *ListCronsRequest, opts ...scw.RequestOption) (*ListCronsResponse, error) {
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
	parameter.AddToQuery(query, "function_id", req.FunctionID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListCronsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetCronRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	CronID string `json:"-"`
}

// GetCron: get a cron
//
// Get the cron associated with the given id.
func (s *API) GetCron(req *GetCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CronID) == "" {
		return nil, errors.New("field CronID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
		Headers: http.Header{},
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateCronRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	FunctionID string `json:"function_id"`

	Schedule string `json:"schedule"`

	Args *scw.JSONObject `json:"args"`

	Name *string `json:"name"`
}

// CreateCron: create a new cron
func (s *API) CreateCron(req *CreateCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateCronRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	CronID string `json:"-"`

	FunctionID *string `json:"function_id"`

	Schedule *string `json:"schedule"`

	Args *scw.JSONObject `json:"args"`

	Name *string `json:"name"`
}

// UpdateCron: update an existing cron
//
// Update the cron associated with the given id.
func (s *API) UpdateCron(req *UpdateCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CronID) == "" {
		return nil, errors.New("field CronID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteCronRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	CronID string `json:"-"`
}

// DeleteCron: delete an existing cron
//
// Delete the cron associated with the given id.
func (s *API) DeleteCron(req *DeleteCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CronID) == "" {
		return nil, errors.New("field CronID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
		Headers: http.Header{},
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListLogsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: timestamp_desc
	OrderBy ListLogsRequestOrderBy `json:"-"`
}

// ListLogs: list your application logs
func (s *API) ListLogs(req *ListLogsRequest, opts ...scw.RequestOption) (*ListLogsResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/logs",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListDomainsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListDomainsRequestOrderBy `json:"-"`

	FunctionID string `json:"-"`
}

// ListDomains: list all domain name bindings
func (s *API) ListDomains(req *ListDomainsRequest, opts ...scw.RequestOption) (*ListDomainsResponse, error) {
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
	parameter.AddToQuery(query, "function_id", req.FunctionID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDomainRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	DomainID string `json:"-"`
}

// GetDomain: get a domain name binding
func (s *API) GetDomain(req *GetDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
		Headers: http.Header{},
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateDomainRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	Hostname string `json:"hostname"`

	FunctionID string `json:"function_id"`
}

// CreateDomain: create a domain name binding
func (s *API) CreateDomain(req *CreateDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteDomainRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	DomainID string `json:"-"`
}

// DeleteDomain: delete a domain name binding
func (s *API) DeleteDomain(req *DeleteDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
		Headers: http.Header{},
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type IssueJWTRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	FunctionID *string `json:"-"`

	NamespaceID *string `json:"-"`

	ExpiresAt *time.Time `json:"-"`
}

// Deprecated
func (s *API) IssueJWT(req *IssueJWTRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "function_id", req.FunctionID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "expires_at", req.ExpiresAt)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/issue-jwt",
		Query:   query,
		Headers: http.Header{},
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateTokenRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	// Precisely one of FunctionID, NamespaceID must be set.
	FunctionID *string `json:"function_id,omitempty"`

	// Precisely one of FunctionID, NamespaceID must be set.
	NamespaceID *string `json:"namespace_id,omitempty"`

	Description *string `json:"description"`

	ExpiresAt *time.Time `json:"expires_at"`
}

// CreateToken: create a new revocable token
func (s *API) CreateToken(req *CreateTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetTokenRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	TokenID string `json:"-"`
}

// GetToken: get a token
func (s *API) GetToken(req *GetTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TokenID) == "" {
		return nil, errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens/" + fmt.Sprint(req.TokenID) + "",
		Headers: http.Header{},
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListTokensRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListTokensRequestOrderBy `json:"-"`

	FunctionID *string `json:"-"`

	NamespaceID *string `json:"-"`
}

// ListTokens: list all tokens
func (s *API) ListTokens(req *ListTokensRequest, opts ...scw.RequestOption) (*ListTokensResponse, error) {
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
	parameter.AddToQuery(query, "function_id", req.FunctionID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListTokensResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteTokenRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	TokenID string `json:"-"`
}

// DeleteToken: delete a token
func (s *API) DeleteToken(req *DeleteTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TokenID) == "" {
		return nil, errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens/" + fmt.Sprint(req.TokenID) + "",
		Headers: http.Header{},
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateTriggerRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	Name string `json:"name"`

	Description string `json:"description"`

	FunctionID string `json:"function_id"`
	// Type:
	//
	// Default value: unknown_trigger_type
	Type TriggerType `json:"type"`

	// Precisely one of NatsFailureHandlingPolicy, SqsFailureHandlingPolicy must be set.
	NatsFailureHandlingPolicy *CreateTriggerRequestNatsFailureHandlingPolicy `json:"nats_failure_handling_policy,omitempty"`

	// Precisely one of NatsFailureHandlingPolicy, SqsFailureHandlingPolicy must be set.
	SqsFailureHandlingPolicy *CreateTriggerRequestSqsFailureHandlingPolicy `json:"sqs_failure_handling_policy,omitempty"`
}

func (s *API) CreateTrigger(req *CreateTriggerRequest, opts ...scw.RequestOption) (*Trigger, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetTriggerRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	TriggerID string `json:"-"`
}

func (s *API) GetTrigger(req *GetTriggerRequest, opts ...scw.RequestOption) (*Trigger, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TriggerID) == "" {
		return nil, errors.New("field TriggerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
		Headers: http.Header{},
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListTriggersRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListTriggersRequestOrderBy `json:"-"`

	FunctionID string `json:"-"`
}

func (s *API) ListTriggers(req *ListTriggersRequest, opts ...scw.RequestOption) (*ListTriggersResponse, error) {
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
	parameter.AddToQuery(query, "function_id", req.FunctionID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListTriggersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateTriggerRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	TriggerID string `json:"-"`

	Name *string `json:"name"`

	Description *string `json:"description"`

	// Precisely one of NatsConfig, SqsConfig must be set.
	NatsConfig *CreateTriggerRequestNatsFailureHandlingPolicy `json:"nats_config,omitempty"`

	// Precisely one of NatsConfig, SqsConfig must be set.
	SqsConfig *CreateTriggerRequestSqsFailureHandlingPolicy `json:"sqs_config,omitempty"`
}

func (s *API) UpdateTrigger(req *UpdateTriggerRequest, opts ...scw.RequestOption) (*Trigger, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TriggerID) == "" {
		return nil, errors.New("field TriggerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteTriggerRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	TriggerID string `json:"-"`
}

func (s *API) DeleteTrigger(req *DeleteTriggerRequest, opts ...scw.RequestOption) (*Trigger, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TriggerID) == "" {
		return nil, errors.New("field TriggerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
		Headers: http.Header{},
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateTriggerInputRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	TriggerID string `json:"trigger_id"`

	MnqNamespaceID *string `json:"mnq_namespace_id"`

	// Precisely one of NatsConfig, SqsConfig must be set.
	NatsConfig *CreateTriggerInputRequestNatsClientConfigSpec `json:"nats_config,omitempty"`

	// Precisely one of NatsConfig, SqsConfig must be set.
	SqsConfig *CreateTriggerInputRequestSqsClientConfigSpec `json:"sqs_config,omitempty"`
}

func (s *API) CreateTriggerInput(req *CreateTriggerInputRequest, opts ...scw.RequestOption) (*TriggerInput, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/trigger-inputs",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TriggerInput

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetTriggerInputRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	TriggerInputID string `json:"-"`
}

func (s *API) GetTriggerInput(req *GetTriggerInputRequest, opts ...scw.RequestOption) (*TriggerInput, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TriggerInputID) == "" {
		return nil, errors.New("field TriggerInputID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/trigger-inputs/" + fmt.Sprint(req.TriggerInputID) + "",
		Headers: http.Header{},
	}

	var resp TriggerInput

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListTriggerInputsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListTriggerInputsRequestOrderBy `json:"-"`

	TriggerID string `json:"-"`
}

func (s *API) ListTriggerInputs(req *ListTriggerInputsRequest, opts ...scw.RequestOption) (*ListTriggerInputsResponse, error) {
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
	parameter.AddToQuery(query, "trigger_id", req.TriggerID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/trigger-inputs",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListTriggerInputsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetTriggerInputsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	TriggerInputID string `json:"trigger_input_id"`

	// Precisely one of Nats, Sqs must be set.
	Sqs *SetTriggerInputsRequestSqsConfigs `json:"sqs,omitempty"`

	// Precisely one of Nats, Sqs must be set.
	Nats *SetTriggerInputsRequestNatsConfigs `json:"nats,omitempty"`
}

func (s *API) SetTriggerInputs(req *SetTriggerInputsRequest, opts ...scw.RequestOption) (*SetTriggerInputsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/trigger-inputs",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetTriggerInputsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateTriggerInputRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	TriggerInputID string `json:"-"`

	// Precisely one of NatsConfig, SqsConfig must be set.
	NatsConfig *UpdateTriggerInputRequestNatsClientConfigSpec `json:"nats_config,omitempty"`

	// Precisely one of NatsConfig, SqsConfig must be set.
	SqsConfig *UpdateTriggerInputRequestSqsClientConfigSpec `json:"sqs_config,omitempty"`
}

func (s *API) UpdateTriggerInput(req *UpdateTriggerInputRequest, opts ...scw.RequestOption) (*TriggerInput, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TriggerInputID) == "" {
		return nil, errors.New("field TriggerInputID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/trigger-inputs/" + fmt.Sprint(req.TriggerInputID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TriggerInput

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteTriggerInputRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	TriggerInputID string `json:"-"`
}

func (s *API) DeleteTriggerInput(req *DeleteTriggerInputRequest, opts ...scw.RequestOption) (*TriggerInput, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TriggerInputID) == "" {
		return nil, errors.New("field TriggerInputID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/trigger-inputs/" + fmt.Sprint(req.TriggerInputID) + "",
		Headers: http.Header{},
	}

	var resp TriggerInput

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNamespacesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Namespaces = append(r.Namespaces, results.Namespaces...)
	r.TotalCount += uint32(len(results.Namespaces))
	return uint32(len(results.Namespaces)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFunctionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFunctionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListFunctionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Functions = append(r.Functions, results.Functions...)
	r.TotalCount += uint32(len(results.Functions))
	return uint32(len(results.Functions)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCronsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCronsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCronsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Crons = append(r.Crons, results.Crons...)
	r.TotalCount += uint32(len(results.Crons))
	return uint32(len(results.Crons)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListLogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Logs = append(r.Logs, results.Logs...)
	r.TotalCount += uint32(len(results.Logs))
	return uint32(len(results.Logs)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDomainsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Domains = append(r.Domains, results.Domains...)
	r.TotalCount += uint32(len(results.Domains))
	return uint32(len(results.Domains)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTokensResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTokensResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTokensResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tokens = append(r.Tokens, results.Tokens...)
	r.TotalCount += uint32(len(results.Tokens))
	return uint32(len(results.Tokens)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTriggersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTriggersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTriggersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Triggers = append(r.Triggers, results.Triggers...)
	r.TotalCount += uint32(len(results.Triggers))
	return uint32(len(results.Triggers)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTriggerInputsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTriggerInputsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTriggerInputsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Inputs = append(r.Inputs, results.Inputs...)
	r.TotalCount += uint32(len(results.Inputs))
	return uint32(len(results.Inputs)), nil
}

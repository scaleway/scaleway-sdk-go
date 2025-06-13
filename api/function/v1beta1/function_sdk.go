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

type CronStatus string

const (
	CronStatusUnknown  = CronStatus("unknown")
	CronStatusReady    = CronStatus("ready")
	CronStatusDeleting = CronStatus("deleting")
	CronStatusError    = CronStatus("error")
	CronStatusLocked   = CronStatus("locked")
	CronStatusCreating = CronStatus("creating")
	CronStatusPending  = CronStatus("pending")
)

func (enum CronStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(CronStatusUnknown)
	}
	return string(enum)
}

func (enum CronStatus) Values() []CronStatus {
	return []CronStatus{
		"unknown",
		"ready",
		"deleting",
		"error",
		"locked",
		"creating",
		"pending",
	}
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
	DomainStatusUnknown  = DomainStatus("unknown")
	DomainStatusReady    = DomainStatus("ready")
	DomainStatusDeleting = DomainStatus("deleting")
	DomainStatusError    = DomainStatus("error")
	DomainStatusCreating = DomainStatus("creating")
	DomainStatusPending  = DomainStatus("pending")
)

func (enum DomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(DomainStatusUnknown)
	}
	return string(enum)
}

func (enum DomainStatus) Values() []DomainStatus {
	return []DomainStatus{
		"unknown",
		"ready",
		"deleting",
		"error",
		"creating",
		"pending",
	}
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
	FunctionHTTPOptionUnknownHTTPOption = FunctionHTTPOption("unknown_http_option")
	FunctionHTTPOptionEnabled           = FunctionHTTPOption("enabled")
	FunctionHTTPOptionRedirected        = FunctionHTTPOption("redirected")
)

func (enum FunctionHTTPOption) String() string {
	if enum == "" {
		// return default value if empty
		return string(FunctionHTTPOptionUnknownHTTPOption)
	}
	return string(enum)
}

func (enum FunctionHTTPOption) Values() []FunctionHTTPOption {
	return []FunctionHTTPOption{
		"unknown_http_option",
		"enabled",
		"redirected",
	}
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
	FunctionPrivacyUnknownPrivacy = FunctionPrivacy("unknown_privacy")
	FunctionPrivacyPublic         = FunctionPrivacy("public")
	FunctionPrivacyPrivate        = FunctionPrivacy("private")
)

func (enum FunctionPrivacy) String() string {
	if enum == "" {
		// return default value if empty
		return string(FunctionPrivacyUnknownPrivacy)
	}
	return string(enum)
}

func (enum FunctionPrivacy) Values() []FunctionPrivacy {
	return []FunctionPrivacy{
		"unknown_privacy",
		"public",
		"private",
	}
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
	FunctionRuntimeUnknownRuntime = FunctionRuntime("unknown_runtime")
	FunctionRuntimeGolang         = FunctionRuntime("golang")
	FunctionRuntimePython         = FunctionRuntime("python")
	FunctionRuntimePython3        = FunctionRuntime("python3")
	FunctionRuntimeNode8          = FunctionRuntime("node8")
	FunctionRuntimeNode10         = FunctionRuntime("node10")
	FunctionRuntimeNode14         = FunctionRuntime("node14")
	FunctionRuntimeNode16         = FunctionRuntime("node16")
	FunctionRuntimeNode17         = FunctionRuntime("node17")
	FunctionRuntimePython37       = FunctionRuntime("python37")
	FunctionRuntimePython38       = FunctionRuntime("python38")
	FunctionRuntimePython39       = FunctionRuntime("python39")
	FunctionRuntimePython310      = FunctionRuntime("python310")
	FunctionRuntimeGo113          = FunctionRuntime("go113")
	FunctionRuntimeGo117          = FunctionRuntime("go117")
	FunctionRuntimeGo118          = FunctionRuntime("go118")
	FunctionRuntimeNode18         = FunctionRuntime("node18")
	FunctionRuntimeRust165        = FunctionRuntime("rust165")
	FunctionRuntimeGo119          = FunctionRuntime("go119")
	FunctionRuntimePython311      = FunctionRuntime("python311")
	FunctionRuntimePhp82          = FunctionRuntime("php82")
	FunctionRuntimeNode19         = FunctionRuntime("node19")
	FunctionRuntimeGo120          = FunctionRuntime("go120")
	FunctionRuntimeNode20         = FunctionRuntime("node20")
	FunctionRuntimeGo121          = FunctionRuntime("go121")
	FunctionRuntimeNode22         = FunctionRuntime("node22")
	FunctionRuntimePython312      = FunctionRuntime("python312")
	FunctionRuntimePhp83          = FunctionRuntime("php83")
	FunctionRuntimeGo122          = FunctionRuntime("go122")
	FunctionRuntimeRust179        = FunctionRuntime("rust179")
	FunctionRuntimeGo123          = FunctionRuntime("go123")
	FunctionRuntimeGo124          = FunctionRuntime("go124")
	FunctionRuntimePython313      = FunctionRuntime("python313")
	FunctionRuntimeRust185        = FunctionRuntime("rust185")
	FunctionRuntimePhp84          = FunctionRuntime("php84")
)

func (enum FunctionRuntime) String() string {
	if enum == "" {
		// return default value if empty
		return string(FunctionRuntimeUnknownRuntime)
	}
	return string(enum)
}

func (enum FunctionRuntime) Values() []FunctionRuntime {
	return []FunctionRuntime{
		"unknown_runtime",
		"golang",
		"python",
		"python3",
		"node8",
		"node10",
		"node14",
		"node16",
		"node17",
		"python37",
		"python38",
		"python39",
		"python310",
		"go113",
		"go117",
		"go118",
		"node18",
		"rust165",
		"go119",
		"python311",
		"php82",
		"node19",
		"go120",
		"node20",
		"go121",
		"node22",
		"python312",
		"php83",
		"go122",
		"rust179",
		"go123",
		"go124",
		"python313",
		"rust185",
		"php84",
	}
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

type FunctionSandbox string

const (
	// Unknown sandbox.
	FunctionSandboxUnknownSandbox = FunctionSandbox("unknown_sandbox")
	// Legacy sandboxing with slower cold starts. Fully supports the Linux system call interface.
	FunctionSandboxV1 = FunctionSandbox("v1")
	// Recommended sandboxing with faster cold starts.
	FunctionSandboxV2 = FunctionSandbox("v2")
)

func (enum FunctionSandbox) String() string {
	if enum == "" {
		// return default value if empty
		return string(FunctionSandboxUnknownSandbox)
	}
	return string(enum)
}

func (enum FunctionSandbox) Values() []FunctionSandbox {
	return []FunctionSandbox{
		"unknown_sandbox",
		"v1",
		"v2",
	}
}

func (enum FunctionSandbox) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionSandbox) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionSandbox(FunctionSandbox(tmp).String())
	return nil
}

type FunctionStatus string

const (
	FunctionStatusUnknown  = FunctionStatus("unknown")
	FunctionStatusReady    = FunctionStatus("ready")
	FunctionStatusDeleting = FunctionStatus("deleting")
	FunctionStatusError    = FunctionStatus("error")
	FunctionStatusLocked   = FunctionStatus("locked")
	FunctionStatusCreating = FunctionStatus("creating")
	FunctionStatusPending  = FunctionStatus("pending")
	FunctionStatusCreated  = FunctionStatus("created")
)

func (enum FunctionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(FunctionStatusUnknown)
	}
	return string(enum)
}

func (enum FunctionStatus) Values() []FunctionStatus {
	return []FunctionStatus{
		"unknown",
		"ready",
		"deleting",
		"error",
		"locked",
		"creating",
		"pending",
		"created",
	}
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
	ListCronsRequestOrderByCreatedAtAsc  = ListCronsRequestOrderBy("created_at_asc")
	ListCronsRequestOrderByCreatedAtDesc = ListCronsRequestOrderBy("created_at_desc")
)

func (enum ListCronsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListCronsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListCronsRequestOrderBy) Values() []ListCronsRequestOrderBy {
	return []ListCronsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
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
	ListDomainsRequestOrderByCreatedAtAsc  = ListDomainsRequestOrderBy("created_at_asc")
	ListDomainsRequestOrderByCreatedAtDesc = ListDomainsRequestOrderBy("created_at_desc")
	ListDomainsRequestOrderByHostnameAsc   = ListDomainsRequestOrderBy("hostname_asc")
	ListDomainsRequestOrderByHostnameDesc  = ListDomainsRequestOrderBy("hostname_desc")
)

func (enum ListDomainsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListDomainsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListDomainsRequestOrderBy) Values() []ListDomainsRequestOrderBy {
	return []ListDomainsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"hostname_asc",
		"hostname_desc",
	}
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
	ListFunctionsRequestOrderByCreatedAtAsc  = ListFunctionsRequestOrderBy("created_at_asc")
	ListFunctionsRequestOrderByCreatedAtDesc = ListFunctionsRequestOrderBy("created_at_desc")
	ListFunctionsRequestOrderByNameAsc       = ListFunctionsRequestOrderBy("name_asc")
	ListFunctionsRequestOrderByNameDesc      = ListFunctionsRequestOrderBy("name_desc")
)

func (enum ListFunctionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListFunctionsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListFunctionsRequestOrderBy) Values() []ListFunctionsRequestOrderBy {
	return []ListFunctionsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
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

type ListNamespacesRequestOrderBy string

const (
	ListNamespacesRequestOrderByCreatedAtAsc  = ListNamespacesRequestOrderBy("created_at_asc")
	ListNamespacesRequestOrderByCreatedAtDesc = ListNamespacesRequestOrderBy("created_at_desc")
	ListNamespacesRequestOrderByNameAsc       = ListNamespacesRequestOrderBy("name_asc")
	ListNamespacesRequestOrderByNameDesc      = ListNamespacesRequestOrderBy("name_desc")
)

func (enum ListNamespacesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListNamespacesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListNamespacesRequestOrderBy) Values() []ListNamespacesRequestOrderBy {
	return []ListNamespacesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
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
	ListTokensRequestOrderByCreatedAtAsc  = ListTokensRequestOrderBy("created_at_asc")
	ListTokensRequestOrderByCreatedAtDesc = ListTokensRequestOrderBy("created_at_desc")
)

func (enum ListTokensRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListTokensRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListTokensRequestOrderBy) Values() []ListTokensRequestOrderBy {
	return []ListTokensRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
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

type ListTriggersRequestOrderBy string

const (
	// Order by creation date ascending.
	ListTriggersRequestOrderByCreatedAtAsc = ListTriggersRequestOrderBy("created_at_asc")
	// Order by creation date descending.
	ListTriggersRequestOrderByCreatedAtDesc = ListTriggersRequestOrderBy("created_at_desc")
)

func (enum ListTriggersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListTriggersRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListTriggersRequestOrderBy) Values() []ListTriggersRequestOrderBy {
	return []ListTriggersRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
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

type NamespaceStatus string

const (
	NamespaceStatusUnknown  = NamespaceStatus("unknown")
	NamespaceStatusReady    = NamespaceStatus("ready")
	NamespaceStatusDeleting = NamespaceStatus("deleting")
	NamespaceStatusError    = NamespaceStatus("error")
	NamespaceStatusLocked   = NamespaceStatus("locked")
	NamespaceStatusCreating = NamespaceStatus("creating")
	NamespaceStatusPending  = NamespaceStatus("pending")
)

func (enum NamespaceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(NamespaceStatusUnknown)
	}
	return string(enum)
}

func (enum NamespaceStatus) Values() []NamespaceStatus {
	return []NamespaceStatus{
		"unknown",
		"ready",
		"deleting",
		"error",
		"locked",
		"creating",
		"pending",
	}
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

type RuntimeStatus string

const (
	RuntimeStatusUnknownStatus = RuntimeStatus("unknown_status")
	RuntimeStatusBeta          = RuntimeStatus("beta")
	RuntimeStatusAvailable     = RuntimeStatus("available")
	RuntimeStatusDeprecated    = RuntimeStatus("deprecated")
	RuntimeStatusEndOfSupport  = RuntimeStatus("end_of_support")
	RuntimeStatusEndOfLife     = RuntimeStatus("end_of_life")
)

func (enum RuntimeStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(RuntimeStatusUnknownStatus)
	}
	return string(enum)
}

func (enum RuntimeStatus) Values() []RuntimeStatus {
	return []RuntimeStatus{
		"unknown_status",
		"beta",
		"available",
		"deprecated",
		"end_of_support",
		"end_of_life",
	}
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
	TokenStatusUnknown  = TokenStatus("unknown")
	TokenStatusReady    = TokenStatus("ready")
	TokenStatusDeleting = TokenStatus("deleting")
	TokenStatusError    = TokenStatus("error")
	TokenStatusCreating = TokenStatus("creating")
)

func (enum TokenStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(TokenStatusUnknown)
	}
	return string(enum)
}

func (enum TokenStatus) Values() []TokenStatus {
	return []TokenStatus{
		"unknown",
		"ready",
		"deleting",
		"error",
		"creating",
	}
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

type TriggerInputType string

const (
	// Unknown input type.
	TriggerInputTypeUnknownInputType = TriggerInputType("unknown_input_type")
	TriggerInputTypeSqs              = TriggerInputType("sqs")
	// Scaleway Messaging and Queuing SQS queue.
	TriggerInputTypeScwSqs = TriggerInputType("scw_sqs")
	TriggerInputTypeNats   = TriggerInputType("nats")
	// Scaleway Messaging and Queuing NATS subject.
	TriggerInputTypeScwNats = TriggerInputType("scw_nats")
)

func (enum TriggerInputType) String() string {
	if enum == "" {
		// return default value if empty
		return string(TriggerInputTypeUnknownInputType)
	}
	return string(enum)
}

func (enum TriggerInputType) Values() []TriggerInputType {
	return []TriggerInputType{
		"unknown_input_type",
		"sqs",
		"scw_sqs",
		"nats",
		"scw_nats",
	}
}

func (enum TriggerInputType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TriggerInputType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TriggerInputType(TriggerInputType(tmp).String())
	return nil
}

type TriggerStatus string

const (
	// Unknown status.
	TriggerStatusUnknownStatus = TriggerStatus("unknown_status")
	// Ready status.
	TriggerStatusReady = TriggerStatus("ready")
	// Deleting status.
	TriggerStatusDeleting = TriggerStatus("deleting")
	// Error status.
	TriggerStatusError = TriggerStatus("error")
	// Creating status.
	TriggerStatusCreating = TriggerStatus("creating")
	// Pending status.
	TriggerStatusPending = TriggerStatus("pending")
)

func (enum TriggerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(TriggerStatusUnknownStatus)
	}
	return string(enum)
}

func (enum TriggerStatus) Values() []TriggerStatus {
	return []TriggerStatus{
		"unknown_status",
		"ready",
		"deleting",
		"error",
		"creating",
		"pending",
	}
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

// SecretHashedValue: secret hashed value.
type SecretHashedValue struct {
	Key string `json:"key"`

	HashedValue string `json:"hashed_value"`
}

// TriggerMnqNatsClientConfig: trigger mnq nats client config.
type TriggerMnqNatsClientConfig struct {
	// Subject: name of the NATS subject the trigger listens to.
	Subject string `json:"subject"`

	// MnqNatsAccountID: ID of the Messaging and Queuing NATS account.
	MnqNatsAccountID string `json:"mnq_nats_account_id"`

	// MnqProjectID: ID of the Messaging and Queuing project.
	MnqProjectID string `json:"mnq_project_id"`

	// MnqRegion: currently, only the `fr-par` and `nl-ams` regions are available.
	MnqRegion string `json:"mnq_region"`

	// MnqCredentialID: ID of the Messaging and Queuing credentials used to subscribe to the NATS subject.
	MnqCredentialID *string `json:"mnq_credential_id"`
}

// TriggerMnqSqsClientConfig: trigger mnq sqs client config.
type TriggerMnqSqsClientConfig struct {
	// Queue: name of the SQS queue the trigger listens to.
	Queue string `json:"queue"`

	// MnqProjectID: ID of the Messaging and Queuing project.
	MnqProjectID string `json:"mnq_project_id"`

	// MnqRegion: currently, only the `fr-par` and `nl-ams` regions are available.
	MnqRegion string `json:"mnq_region"`

	// MnqCredentialID: ID of the Messaging and Queuing credentials used to read from the SQS queue.
	MnqCredentialID *string `json:"mnq_credential_id"`
}

// TriggerSqsClientConfig: trigger sqs client config.
type TriggerSqsClientConfig struct {
	Endpoint string `json:"endpoint"`

	QueueURL string `json:"queue_url"`

	AccessKey string `json:"access_key"`

	SecretKey string `json:"secret_key"`
}

// Secret: secret.
type Secret struct {
	Key string `json:"key"`

	Value *string `json:"value"`
}

// CreateTriggerRequestMnqNatsClientConfig: create trigger request mnq nats client config.
type CreateTriggerRequestMnqNatsClientConfig struct {
	// Subject: name of the NATS subject the trigger should listen to.
	Subject string `json:"subject"`

	// MnqNatsAccountID: ID of the Messaging and Queuing NATS account.
	MnqNatsAccountID string `json:"mnq_nats_account_id"`

	// MnqProjectID: ID of the Messaging and Queuing project.
	MnqProjectID string `json:"mnq_project_id"`

	// MnqRegion: currently, only the `fr-par` and `nl-ams` regions are available.
	MnqRegion string `json:"mnq_region"`
}

// CreateTriggerRequestMnqSqsClientConfig: create trigger request mnq sqs client config.
type CreateTriggerRequestMnqSqsClientConfig struct {
	// Queue: name of the SQS queue the trigger should listen to.
	Queue string `json:"queue"`

	// MnqProjectID: you must have activated SQS on this project.
	MnqProjectID string `json:"mnq_project_id"`

	// MnqRegion: currently, only the `fr-par` and `nl-ams` regions are available.
	MnqRegion string `json:"mnq_region"`
}

// CreateTriggerRequestSqsClientConfig: create trigger request sqs client config.
type CreateTriggerRequestSqsClientConfig struct {
	Endpoint string `json:"endpoint"`

	QueueURL string `json:"queue_url"`

	AccessKey string `json:"access_key"`

	SecretKey string `json:"secret_key"`
}

// Cron: cron.
type Cron struct {
	// ID: UUID of the cron.
	ID string `json:"id"`

	// FunctionID: UUID of the function the cron applies to.
	FunctionID string `json:"function_id"`

	// Schedule: schedule of the cron.
	Schedule string `json:"schedule"`

	// Args: arguments to pass with the cron.
	Args *scw.JSONObject `json:"args"`

	// Status: status of the cron.
	// Default value: unknown
	Status CronStatus `json:"status"`

	// Name: name of the cron.
	Name string `json:"name"`
}

// Domain: domain.
type Domain struct {
	// ID: UUID of the domain.
	ID string `json:"id"`

	// Hostname: hostname associated with the function.
	Hostname string `json:"hostname"`

	// FunctionID: UUID of the function the domain is associated with.
	FunctionID string `json:"function_id"`

	// URL: URL of the function.
	URL string `json:"url"`

	// Status: state of the doamin.
	// Default value: unknown
	Status DomainStatus `json:"status"`

	// ErrorMessage: error message if the domain is in "error" state.
	ErrorMessage *string `json:"error_message"`
}

// Runtime: runtime.
type Runtime struct {
	Name string `json:"name"`

	Language string `json:"language"`

	Version string `json:"version"`

	DefaultHandler string `json:"default_handler"`

	CodeSample string `json:"code_sample"`

	// Status: default value: unknown_status
	Status RuntimeStatus `json:"status"`

	StatusMessage string `json:"status_message"`

	Extension string `json:"extension"`

	Implementation string `json:"implementation"`

	LogoURL string `json:"logo_url"`
}

// Function: function.
type Function struct {
	// ID: UUID of the function.
	ID string `json:"id"`

	// Name: name of the function.
	Name string `json:"name"`

	// NamespaceID: UUID of the namespace the function belongs to.
	NamespaceID string `json:"namespace_id"`

	// Status: status of the function.
	// Default value: unknown
	Status FunctionStatus `json:"status"`

	// EnvironmentVariables: environment variables of the function.
	EnvironmentVariables map[string]string `json:"environment_variables"`

	// MinScale: minimum number of instances to scale the function to.
	MinScale uint32 `json:"min_scale"`

	// MaxScale: maximum number of instances to scale the function to.
	MaxScale uint32 `json:"max_scale"`

	// Runtime: runtime of the function.
	// Default value: unknown_runtime
	Runtime FunctionRuntime `json:"runtime"`

	// MemoryLimit: memory limit of the function in MB.
	MemoryLimit uint32 `json:"memory_limit"`

	// CPULimit: CPU limit of the function.
	CPULimit uint32 `json:"cpu_limit"`

	// Timeout: request processing time limit for the function.
	Timeout *scw.Duration `json:"timeout"`

	// Handler: handler to use for the function.
	Handler string `json:"handler"`

	// ErrorMessage: error message if the function is in "error" state.
	ErrorMessage *string `json:"error_message"`

	// BuildMessage: description of the current build step.
	BuildMessage *string `json:"build_message"`

	// Privacy: privacy setting of the function.
	// Default value: unknown_privacy
	Privacy FunctionPrivacy `json:"privacy"`

	// Description: description of the function.
	Description *string `json:"description"`

	// DomainName: domain name associated with the function.
	DomainName string `json:"domain_name"`

	// SecretEnvironmentVariables: secret environment variables of the function.
	SecretEnvironmentVariables []*SecretHashedValue `json:"secret_environment_variables"`

	// Region: region in which the function is deployed.
	Region scw.Region `json:"region"`

	// HTTPOption: possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	// Default value: unknown_http_option
	HTTPOption FunctionHTTPOption `json:"http_option"`

	RuntimeMessage string `json:"runtime_message"`

	// Sandbox: execution environment of the function.
	// Default value: unknown_sandbox
	Sandbox FunctionSandbox `json:"sandbox"`

	// CreatedAt: creation date of the function.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date of the function.
	UpdatedAt *time.Time `json:"updated_at"`

	// ReadyAt: last date when the function was successfully deployed and set to ready.
	ReadyAt *time.Time `json:"ready_at"`

	// Tags: list of tags applied to the Serverless Function.
	Tags []string `json:"tags"`

	// PrivateNetworkID: when connected to a Private Network, the function can access other Scaleway resources in this Private Network.
	PrivateNetworkID *string `json:"private_network_id"`
}

// Namespace: namespace.
type Namespace struct {
	// ID: UUID of the namespace.
	ID string `json:"id"`

	// Name: name of the namespace.
	Name string `json:"name"`

	// EnvironmentVariables: environment variables of the namespace.
	EnvironmentVariables map[string]string `json:"environment_variables"`

	// OrganizationID: UUID of the Organization the namespace belongs to.
	OrganizationID string `json:"organization_id"`

	// ProjectID: UUID of the Project the namespace belongs to.
	ProjectID string `json:"project_id"`

	// Status: status of the namespace.
	// Default value: unknown
	Status NamespaceStatus `json:"status"`

	// RegistryNamespaceID: UUID of the registry namespace.
	RegistryNamespaceID string `json:"registry_namespace_id"`

	// ErrorMessage: error message if the namespace is in "error" state.
	ErrorMessage *string `json:"error_message"`

	// RegistryEndpoint: registry endpoint of the namespace.
	RegistryEndpoint string `json:"registry_endpoint"`

	// Description: description of the namespace.
	Description *string `json:"description"`

	// SecretEnvironmentVariables: secret environment variables of the namespace.
	SecretEnvironmentVariables []*SecretHashedValue `json:"secret_environment_variables"`

	// Region: region in which the namespace is located.
	Region scw.Region `json:"region"`

	// Tags: list of tags applied to the Serverless Function Namespace.
	Tags []string `json:"tags"`

	// CreatedAt: creation date of the namespace.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date of the namespace.
	UpdatedAt *time.Time `json:"updated_at"`

	// Deprecated: VpcIntegrationActivated: when activated, functions in the namespace can be connected to a Private Network.
	// Note that activating the VPC integration can only be done when creating a new namespace.
	VpcIntegrationActivated *bool `json:"vpc_integration_activated,omitempty"`
}

// Token: token.
type Token struct {
	// ID: UUID of the token.
	ID string `json:"id"`

	// Token: string of the token.
	Token string `json:"token"`

	// FunctionID: UUID of the function the token is associated with.
	// Precisely one of FunctionID, NamespaceID must be set.
	FunctionID *string `json:"function_id,omitempty"`

	// NamespaceID: UUID of the namespace the token is associated with.
	// Precisely one of FunctionID, NamespaceID must be set.
	NamespaceID *string `json:"namespace_id,omitempty"`

	// Deprecated: PublicKey: public key of the token.
	PublicKey *string `json:"public_key,omitempty"`

	// Status: status of the token.
	// Default value: unknown
	Status TokenStatus `json:"status"`

	// Description: description of the token.
	Description *string `json:"description"`

	// ExpiresAt: date on which the token expires.
	ExpiresAt *time.Time `json:"expires_at"`
}

// Trigger: trigger.
type Trigger struct {
	// ID: ID of the trigger.
	ID string `json:"id"`

	// Name: name of the trigger.
	Name string `json:"name"`

	// Description: description of the trigger.
	Description string `json:"description"`

	// FunctionID: ID of the function to trigger.
	FunctionID string `json:"function_id"`

	// InputType: type of the input.
	// Default value: unknown_input_type
	InputType TriggerInputType `json:"input_type"`

	// Status: status of the trigger.
	// Default value: unknown_status
	Status TriggerStatus `json:"status"`

	// ErrorMessage: error message of the trigger.
	ErrorMessage *string `json:"error_message"`

	// ScwSqsConfig: configuration for a Scaleway Messaging and Queuing SQS queue.
	// Precisely one of ScwSqsConfig, ScwNatsConfig, SqsConfig must be set.
	ScwSqsConfig *TriggerMnqSqsClientConfig `json:"scw_sqs_config,omitempty"`

	// ScwNatsConfig: configuration for a Scaleway Messaging and Queuing NATS subject.
	// Precisely one of ScwSqsConfig, ScwNatsConfig, SqsConfig must be set.
	ScwNatsConfig *TriggerMnqNatsClientConfig `json:"scw_nats_config,omitempty"`

	// SqsConfig: configuration for an AWS SQS queue.
	// Precisely one of ScwSqsConfig, ScwNatsConfig, SqsConfig must be set.
	SqsConfig *TriggerSqsClientConfig `json:"sqs_config,omitempty"`
}

// UpdateTriggerRequestSqsClientConfig: update trigger request sqs client config.
type UpdateTriggerRequestSqsClientConfig struct {
	AccessKey *string `json:"access_key"`

	SecretKey *string `json:"secret_key"`
}

// CreateCronRequest: create cron request.
type CreateCronRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FunctionID: UUID of the function to use the cron with.
	FunctionID string `json:"function_id"`

	// Schedule: schedule of the cron in UNIX cron format.
	Schedule string `json:"schedule"`

	// Args: arguments to use with the cron.
	Args *scw.JSONObject `json:"args,omitempty"`

	// Name: name of the cron.
	Name *string `json:"name,omitempty"`
}

// CreateDomainRequest: create domain request.
type CreateDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Hostname: hostame to create.
	Hostname string `json:"hostname"`

	// FunctionID: UUID of the function to associate the domain with.
	FunctionID string `json:"function_id"`
}

// CreateFunctionRequest: create function request.
type CreateFunctionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: name of the function to create.
	Name string `json:"name"`

	// NamespaceID: UUID of the namespace the function will be created in.
	NamespaceID string `json:"namespace_id"`

	// EnvironmentVariables: environment variables of the function.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// MinScale: minimum number of instances to scale the function to.
	MinScale *uint32 `json:"min_scale,omitempty"`

	// MaxScale: maximum number of instances to scale the function to.
	MaxScale *uint32 `json:"max_scale,omitempty"`

	// Runtime: runtime to use with the function.
	// Default value: unknown_runtime
	Runtime FunctionRuntime `json:"runtime"`

	// MemoryLimit: memory limit of the function in MB.
	MemoryLimit *uint32 `json:"memory_limit,omitempty"`

	// Timeout: request processing time limit for the function.
	Timeout *scw.Duration `json:"timeout,omitempty"`

	// Handler: handler to use with the function.
	Handler *string `json:"handler,omitempty"`

	// Privacy: privacy setting of the function.
	// Default value: unknown_privacy
	Privacy FunctionPrivacy `json:"privacy"`

	// Description: description of the function.
	Description *string `json:"description,omitempty"`

	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`

	// HTTPOption: possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	// Default value: unknown_http_option
	HTTPOption FunctionHTTPOption `json:"http_option"`

	// Sandbox: execution environment of the function.
	// Default value: unknown_sandbox
	Sandbox FunctionSandbox `json:"sandbox"`

	// Tags: tags of the Serverless Function.
	Tags []string `json:"tags"`

	// PrivateNetworkID: when connected to a Private Network, the function can access other Scaleway resources in this Private Network.
	//
	// Note: this feature is currently in beta and requires a namespace with VPC integration activated, using the `activate_vpc_integration` flag.
	PrivateNetworkID *string `json:"private_network_id,omitempty"`
}

// CreateNamespaceRequest: create namespace request.
type CreateNamespaceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Name string `json:"name"`

	// EnvironmentVariables: environment variables of the namespace.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// ProjectID: UUID of the project in which the namespace will be created.
	ProjectID string `json:"project_id"`

	// Description: description of the namespace.
	Description *string `json:"description,omitempty"`

	// SecretEnvironmentVariables: secret environment variables of the namespace.
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`

	// Tags: tags of the Serverless Function Namespace.
	Tags []string `json:"tags"`

	// ActivateVpcIntegration: when activated, functions in the namespace can be connected to a Private Network.
	ActivateVpcIntegration bool `json:"activate_vpc_integration"`
}

// CreateTokenRequest: create token request.
type CreateTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FunctionID: UUID of the function to associate the token with.
	// Precisely one of FunctionID, NamespaceID must be set.
	FunctionID *string `json:"function_id,omitempty"`

	// NamespaceID: UUID of the namespace to associate the token with.
	// Precisely one of FunctionID, NamespaceID must be set.
	NamespaceID *string `json:"namespace_id,omitempty"`

	// Description: description of the token.
	Description *string `json:"description,omitempty"`

	// ExpiresAt: date on which the token expires.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// CreateTriggerRequest: create trigger request.
type CreateTriggerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: name of the trigger.
	Name string `json:"name"`

	// FunctionID: ID of the function to trigger.
	FunctionID string `json:"function_id"`

	// Description: description of the trigger.
	Description *string `json:"description,omitempty"`

	// ScwSqsConfig: configuration for a Scaleway Messaging and Queuing SQS queue.
	// Precisely one of ScwSqsConfig, ScwNatsConfig, SqsConfig must be set.
	ScwSqsConfig *CreateTriggerRequestMnqSqsClientConfig `json:"scw_sqs_config,omitempty"`

	// ScwNatsConfig: configuration for a Scaleway Messaging and Queuing NATS subject.
	// Precisely one of ScwSqsConfig, ScwNatsConfig, SqsConfig must be set.
	ScwNatsConfig *CreateTriggerRequestMnqNatsClientConfig `json:"scw_nats_config,omitempty"`

	// SqsConfig: configuration for an AWS SQS queue.
	// Precisely one of ScwSqsConfig, ScwNatsConfig, SqsConfig must be set.
	SqsConfig *CreateTriggerRequestSqsClientConfig `json:"sqs_config,omitempty"`
}

// DeleteCronRequest: delete cron request.
type DeleteCronRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// CronID: UUID of the cron to delete.
	CronID string `json:"-"`
}

// DeleteDomainRequest: delete domain request.
type DeleteDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainID: UUID of the domain to delete.
	DomainID string `json:"-"`
}

// DeleteFunctionRequest: delete function request.
type DeleteFunctionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FunctionID: UUID of the function to delete.
	FunctionID string `json:"-"`
}

// DeleteNamespaceRequest: delete namespace request.
type DeleteNamespaceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NamespaceID: UUID of the namespace.
	NamespaceID string `json:"-"`
}

// DeleteTokenRequest: delete token request.
type DeleteTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TokenID: UUID of the token to delete.
	TokenID string `json:"-"`
}

// DeleteTriggerRequest: delete trigger request.
type DeleteTriggerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TriggerID: ID of the trigger to delete.
	TriggerID string `json:"-"`
}

// DeployFunctionRequest: deploy function request.
type DeployFunctionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FunctionID: UUID of the function to deploy.
	FunctionID string `json:"-"`
}

// DownloadURL: download url.
type DownloadURL struct {
	URL string `json:"url"`

	Headers map[string]*[]string `json:"headers"`
}

// GetCronRequest: get cron request.
type GetCronRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// CronID: UUID of the cron to get.
	CronID string `json:"-"`
}

// GetDomainRequest: get domain request.
type GetDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainID: UUID of the domain to get.
	DomainID string `json:"-"`
}

// GetFunctionDownloadURLRequest: get function download url request.
type GetFunctionDownloadURLRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FunctionID: UUID of the function to get the download URL for.
	FunctionID string `json:"-"`
}

// GetFunctionRequest: get function request.
type GetFunctionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FunctionID: UUID of the function.
	FunctionID string `json:"-"`
}

// GetFunctionUploadURLRequest: get function upload url request.
type GetFunctionUploadURLRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FunctionID: UUID of the function to get the upload URL for.
	FunctionID string `json:"-"`

	// ContentLength: size of the archive to upload in bytes.
	ContentLength uint64 `json:"content_length"`
}

// GetNamespaceRequest: get namespace request.
type GetNamespaceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NamespaceID: UUID of the namespace.
	NamespaceID string `json:"-"`
}

// GetTokenRequest: get token request.
type GetTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TokenID: UUID of the token to get.
	TokenID string `json:"-"`
}

// GetTriggerRequest: get trigger request.
type GetTriggerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TriggerID: ID of the trigger to get.
	TriggerID string `json:"-"`
}

// ListCronsRequest: list crons request.
type ListCronsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: number of crons per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the crons.
	// Default value: created_at_asc
	OrderBy ListCronsRequestOrderBy `json:"-"`

	// FunctionID: UUID of the function.
	FunctionID string `json:"-"`
}

// ListCronsResponse: list crons response.
type ListCronsResponse struct {
	// Crons: array of crons.
	Crons []*Cron `json:"crons"`

	// TotalCount: total number of crons.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCronsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCronsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListCronsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Crons = append(r.Crons, results.Crons...)
	r.TotalCount += uint32(len(results.Crons))
	return uint32(len(results.Crons)), nil
}

// ListDomainsRequest: list domains request.
type ListDomainsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: number of domains per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the domains.
	// Default value: created_at_asc
	OrderBy ListDomainsRequestOrderBy `json:"-"`

	// FunctionID: UUID of the function the domain is associated with.
	FunctionID string `json:"-"`
}

// ListDomainsResponse: list domains response.
type ListDomainsResponse struct {
	// Domains: array of domains.
	Domains []*Domain `json:"domains"`

	// TotalCount: total number of domains.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListDomainsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Domains = append(r.Domains, results.Domains...)
	r.TotalCount += uint32(len(results.Domains))
	return uint32(len(results.Domains)), nil
}

// ListFunctionRuntimesRequest: list function runtimes request.
type ListFunctionRuntimesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// ListFunctionRuntimesResponse: list function runtimes response.
type ListFunctionRuntimesResponse struct {
	// Runtimes: array of runtimes available.
	Runtimes []*Runtime `json:"runtimes"`

	// TotalCount: total number of runtimes.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFunctionRuntimesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFunctionRuntimesResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListFunctionRuntimesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Runtimes = append(r.Runtimes, results.Runtimes...)
	r.TotalCount += uint32(len(results.Runtimes))
	return uint32(len(results.Runtimes)), nil
}

// ListFunctionsRequest: list functions request.
type ListFunctionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: number of functions per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the functions.
	// Default value: created_at_asc
	OrderBy ListFunctionsRequestOrderBy `json:"-"`

	// NamespaceID: UUID of the namespace the function belongs to.
	NamespaceID string `json:"-"`

	// Name: name of the function.
	Name *string `json:"-"`

	// OrganizationID: UUID of the Organization the function belongs to.
	OrganizationID *string `json:"-"`

	// ProjectID: UUID of the Project the function belongs to.
	ProjectID *string `json:"-"`
}

// ListFunctionsResponse: list functions response.
type ListFunctionsResponse struct {
	// Functions: array of functions.
	Functions []*Function `json:"functions"`

	// TotalCount: total number of functions.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFunctionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFunctionsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListFunctionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Functions = append(r.Functions, results.Functions...)
	r.TotalCount += uint32(len(results.Functions))
	return uint32(len(results.Functions)), nil
}

// ListNamespacesRequest: list namespaces request.
type ListNamespacesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: number of namespaces per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the namespaces.
	// Default value: created_at_asc
	OrderBy ListNamespacesRequestOrderBy `json:"-"`

	// Name: name of the namespace.
	Name *string `json:"-"`

	// OrganizationID: UUID of the Organization the namespace belongs to.
	OrganizationID *string `json:"-"`

	// ProjectID: UUID of the Project the namespace belongs to.
	ProjectID *string `json:"-"`
}

// ListNamespacesResponse: list namespaces response.
type ListNamespacesResponse struct {
	Namespaces []*Namespace `json:"namespaces"`

	// TotalCount: total number of namespaces.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListNamespacesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Namespaces = append(r.Namespaces, results.Namespaces...)
	r.TotalCount += uint32(len(results.Namespaces))
	return uint32(len(results.Namespaces)), nil
}

// ListTokensRequest: list tokens request.
type ListTokensRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: number of tokens per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order for the tokens.
	// Default value: created_at_asc
	OrderBy ListTokensRequestOrderBy `json:"-"`

	// FunctionID: UUID of the function the token is associated with.
	FunctionID *string `json:"-"`

	// NamespaceID: UUID of the namespace the token is associated with.
	NamespaceID *string `json:"-"`
}

// ListTokensResponse: list tokens response.
type ListTokensResponse struct {
	Tokens []*Token `json:"tokens"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTokensResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTokensResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListTokensResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tokens = append(r.Tokens, results.Tokens...)
	r.TotalCount += uint32(len(results.Tokens))
	return uint32(len(results.Tokens)), nil
}

// ListTriggersRequest: list triggers request.
type ListTriggersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of triggers to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListTriggersRequestOrderBy `json:"-"`

	// FunctionID: ID of the function the triggers belongs to.
	// Precisely one of FunctionID, NamespaceID, ProjectID must be set.
	FunctionID *string `json:"function_id,omitempty"`

	// NamespaceID: ID of the namespace the triggers belongs to.
	// Precisely one of FunctionID, NamespaceID, ProjectID must be set.
	NamespaceID *string `json:"namespace_id,omitempty"`

	// ProjectID: ID of the project the triggers belongs to.
	// Precisely one of FunctionID, NamespaceID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`
}

// ListTriggersResponse: list triggers response.
type ListTriggersResponse struct {
	// TotalCount: total count of existing triggers (matching any filters specified).
	TotalCount uint32 `json:"total_count"`

	// Triggers: triggers on this page.
	Triggers []*Trigger `json:"triggers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTriggersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTriggersResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListTriggersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Triggers = append(r.Triggers, results.Triggers...)
	r.TotalCount += uint32(len(results.Triggers))
	return uint32(len(results.Triggers)), nil
}

// UpdateCronRequest: update cron request.
type UpdateCronRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// CronID: UUID of the cron to update.
	CronID string `json:"-"`

	// FunctionID: UUID of the function to use the cron with.
	FunctionID *string `json:"function_id,omitempty"`

	// Schedule: schedule of the cron in UNIX cron format.
	Schedule *string `json:"schedule,omitempty"`

	// Args: arguments to use with the cron.
	Args *scw.JSONObject `json:"args,omitempty"`

	// Name: name of the cron.
	Name *string `json:"name,omitempty"`
}

// UpdateFunctionRequest: update function request.
type UpdateFunctionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FunctionID: UUID of the function to update.
	FunctionID string `json:"-"`

	// EnvironmentVariables: environment variables of the function to update.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// MinScale: minimum number of instances to scale the function to.
	MinScale *uint32 `json:"min_scale,omitempty"`

	// MaxScale: maximum number of instances to scale the function to.
	MaxScale *uint32 `json:"max_scale,omitempty"`

	// Runtime: runtime to use with the function.
	// Default value: unknown_runtime
	Runtime FunctionRuntime `json:"runtime"`

	// MemoryLimit: memory limit of the function in MB.
	MemoryLimit *uint32 `json:"memory_limit,omitempty"`

	// Timeout: processing time limit for the function.
	Timeout *scw.Duration `json:"timeout,omitempty"`

	// Redeploy: redeploy failed function.
	Redeploy *bool `json:"redeploy,omitempty"`

	// Handler: handler to use with the function.
	Handler *string `json:"handler,omitempty"`

	// Privacy: privacy setting of the function.
	// Default value: unknown_privacy
	Privacy FunctionPrivacy `json:"privacy"`

	// Description: description of the function.
	Description *string `json:"description,omitempty"`

	// SecretEnvironmentVariables: secret environment variables of the function.
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`

	// HTTPOption: possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	// Default value: unknown_http_option
	HTTPOption FunctionHTTPOption `json:"http_option"`

	// Sandbox: execution environment of the function.
	// Default value: unknown_sandbox
	Sandbox FunctionSandbox `json:"sandbox"`

	// Tags: tags of the Serverless Function.
	Tags *[]string `json:"tags,omitempty"`

	// PrivateNetworkID: when connected to a Private Network, the function can access other Scaleway resources in this Private Network.
	//
	// Note: this feature is currently in beta and requires a namespace with VPC integration activated, using the `activate_vpc_integration` flag.
	PrivateNetworkID *string `json:"private_network_id,omitempty"`
}

// UpdateNamespaceRequest: update namespace request.
type UpdateNamespaceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NamespaceID: UUID of the namespapce.
	NamespaceID string `json:"-"`

	// EnvironmentVariables: environment variables of the namespace.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// Description: description of the namespace.
	Description *string `json:"description,omitempty"`

	// SecretEnvironmentVariables: secret environment variables of the namespace.
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`

	// Tags: tags of the Serverless Function Namespace.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateTriggerRequest: update trigger request.
type UpdateTriggerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TriggerID: ID of the trigger to update.
	TriggerID string `json:"-"`

	// Name: name of the trigger.
	Name *string `json:"name,omitempty"`

	// Description: description of the trigger.
	Description *string `json:"description,omitempty"`

	// SqsConfig: configuration for an AWS SQS queue.
	// Precisely one of SqsConfig must be set.
	SqsConfig *UpdateTriggerRequestSqsClientConfig `json:"sqs_config,omitempty"`
}

// UploadURL: upload url.
type UploadURL struct {
	// URL: upload URL to upload the function to.
	URL string `json:"url"`

	// Headers: HTTP headers.
	Headers map[string]*[]string `json:"headers"`
}

// This API allows you to manage your Serverless Functions.
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

// ListNamespaces: List all existing namespaces in the specified region.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Query:  query,
	}

	var resp ListNamespacesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNamespace: Get the namespace associated with the specified ID.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateNamespace: Create a new namespace in a specified Organization or Project.
func (s *API) CreateNamespace(req *CreateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
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
		req.Name = namegenerator.GetRandomName("ns")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
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

// UpdateNamespace: Update the namespace associated with the specified ID.
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
		Method: "PATCH",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
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

// DeleteNamespace: Delete the namespace associated with the specified ID.
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
		Method: "DELETE",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFunctions: List all your functions.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions",
		Query:  query,
	}

	var resp ListFunctionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFunction: Get the function associated with the specified ID.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFunction: Create a new function in the specified region for a specified Organization or Project.
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
		Method: "POST",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions",
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

// UpdateFunction: Update the function associated with the specified ID.
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
		Method: "PATCH",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
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

// DeleteFunction: Delete the function associated with the specified ID.
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
		Method: "DELETE",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeployFunction: Deploy a function associated with the specified ID.
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
		Method: "POST",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/deploy",
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

// ListFunctionRuntimes: List available function runtimes.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/runtimes",
	}

	var resp ListFunctionRuntimesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFunctionUploadURL: Get an upload URL of a function associated with the specified ID.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/upload-url",
		Query:  query,
	}

	var resp UploadURL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFunctionDownloadURL: Get a download URL for a function associated with the specified ID.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/download-url",
	}

	var resp DownloadURL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCrons: List all the cronjobs in a specified region.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
		Query:  query,
	}

	var resp ListCronsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCron: Get the cron associated with the specified ID.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCron: Create a new cronjob for a function with the specified ID.
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
		Method: "POST",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
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

// UpdateCron: Update the cron associated with the specified ID.
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
		Method: "PATCH",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
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

// DeleteCron: Delete the cron associated with the specified ID.
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
		Method: "DELETE",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDomains: List all domain name bindings in a specified region.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Query:  query,
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomain: Get a domain name binding for the function with the specified ID.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDomain: Create a domain name binding for the function with the specified ID.
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
		Method: "POST",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains",
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

// DeleteDomain: Delete a domain name binding for the function with the specified ID.
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
		Method: "DELETE",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateToken: Create a new revocable token.
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
		Method: "POST",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens",
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

// GetToken: Get a token.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTokens: List all tokens.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens",
		Query:  query,
	}

	var resp ListTokensResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteToken: Delete a token.
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
		Method: "DELETE",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTrigger: Create a new trigger for a specified function.
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
		Method: "POST",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers",
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

// GetTrigger: Get a trigger with a specified ID.
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
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTriggers: List all triggers belonging to a specified Organization or Project.
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

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.FunctionID == nil && req.NamespaceID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "function_id", req.FunctionID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers",
		Query:  query,
	}

	var resp ListTriggersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateTrigger: Update a trigger with a specified ID.
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
		Method: "PATCH",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
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

// DeleteTrigger: Delete a trigger with a specified ID.
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
		Method: "DELETE",
		Path:   "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

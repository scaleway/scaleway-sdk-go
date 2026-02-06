// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package container provides methods and message types of the container v1beta1 API.
package container

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
	defaultContainerRetryInterval = 15 * time.Second
	defaultContainerTimeout       = 15 * time.Minute
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

type ContainerHTTPOption string

const (
	ContainerHTTPOptionUnknownHTTPOption = ContainerHTTPOption("unknown_http_option")
	ContainerHTTPOptionEnabled           = ContainerHTTPOption("enabled")
	ContainerHTTPOptionRedirected        = ContainerHTTPOption("redirected")
)

func (enum ContainerHTTPOption) String() string {
	if enum == "" {
		// return default value if empty
		return string(ContainerHTTPOptionUnknownHTTPOption)
	}
	return string(enum)
}

func (enum ContainerHTTPOption) Values() []ContainerHTTPOption {
	return []ContainerHTTPOption{
		"unknown_http_option",
		"enabled",
		"redirected",
	}
}

func (enum ContainerHTTPOption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContainerHTTPOption) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContainerHTTPOption(ContainerHTTPOption(tmp).String())
	return nil
}

type ContainerPrivacy string

const (
	ContainerPrivacyUnknownPrivacy = ContainerPrivacy("unknown_privacy")
	ContainerPrivacyPublic         = ContainerPrivacy("public")
	ContainerPrivacyPrivate        = ContainerPrivacy("private")
)

func (enum ContainerPrivacy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ContainerPrivacyUnknownPrivacy)
	}
	return string(enum)
}

func (enum ContainerPrivacy) Values() []ContainerPrivacy {
	return []ContainerPrivacy{
		"unknown_privacy",
		"public",
		"private",
	}
}

func (enum ContainerPrivacy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContainerPrivacy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContainerPrivacy(ContainerPrivacy(tmp).String())
	return nil
}

type ContainerProtocol string

const (
	ContainerProtocolUnknownProtocol = ContainerProtocol("unknown_protocol")
	ContainerProtocolHTTP1           = ContainerProtocol("http1")
	ContainerProtocolH2c             = ContainerProtocol("h2c")
)

func (enum ContainerProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return string(ContainerProtocolUnknownProtocol)
	}
	return string(enum)
}

func (enum ContainerProtocol) Values() []ContainerProtocol {
	return []ContainerProtocol{
		"unknown_protocol",
		"http1",
		"h2c",
	}
}

func (enum ContainerProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContainerProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContainerProtocol(ContainerProtocol(tmp).String())
	return nil
}

type ContainerSandbox string

const (
	// Unknown sandbox.
	ContainerSandboxUnknownSandbox = ContainerSandbox("unknown_sandbox")
	// Legacy sandboxing with slower cold starts. Fully supports the Linux system call interface.
	ContainerSandboxV1 = ContainerSandbox("v1")
	// Recommended sandboxing with faster cold starts.
	ContainerSandboxV2 = ContainerSandbox("v2")
)

func (enum ContainerSandbox) String() string {
	if enum == "" {
		// return default value if empty
		return string(ContainerSandboxUnknownSandbox)
	}
	return string(enum)
}

func (enum ContainerSandbox) Values() []ContainerSandbox {
	return []ContainerSandbox{
		"unknown_sandbox",
		"v1",
		"v2",
	}
}

func (enum ContainerSandbox) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContainerSandbox) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContainerSandbox(ContainerSandbox(tmp).String())
	return nil
}

type ContainerStatus string

const (
	ContainerStatusUnknown = ContainerStatus("unknown")
	// Ready status.
	ContainerStatusReady = ContainerStatus("ready")
	// Deleting status.
	ContainerStatusDeleting = ContainerStatus("deleting")
	// Error status.
	ContainerStatusError = ContainerStatus("error")
	// Locked status. Resource cannot be modified.
	ContainerStatusLocked = ContainerStatus("locked")
	// Creating status. Resource is being created.
	ContainerStatusCreating = ContainerStatus("creating")
	// Pending status. Resource is being deployed after its creation or an update.
	ContainerStatusPending = ContainerStatus("pending")
	// Created status. Resource has been created, but is waiting to be deployed. Call DeployContainer to trigger a deployment.
	ContainerStatusCreated = ContainerStatus("created")
	// Locking status.
	ContainerStatusLocking = ContainerStatus("locking")
	// Upgrading status. Resource is being upgraded as part of a planned maintenance. No downtime is expected.
	ContainerStatusUpgrading = ContainerStatus("upgrading")
)

func (enum ContainerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ContainerStatusUnknown)
	}
	return string(enum)
}

func (enum ContainerStatus) Values() []ContainerStatus {
	return []ContainerStatus{
		"unknown",
		"ready",
		"deleting",
		"error",
		"locked",
		"creating",
		"pending",
		"created",
		"locking",
		"upgrading",
	}
}

func (enum ContainerStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContainerStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContainerStatus(ContainerStatus(tmp).String())
	return nil
}

type CronStatus string

const (
	CronStatusUnknown     = CronStatus("unknown")
	CronStatusReady       = CronStatus("ready")
	CronStatusDeleting    = CronStatus("deleting")
	CronStatusError       = CronStatus("error")
	CronStatusLocked      = CronStatus("locked")
	CronStatusCreating    = CronStatus("creating")
	CronStatusPending     = CronStatus("pending")
	CronStatusLocking     = CronStatus("locking")
	CronStatusUpgrading   = CronStatus("upgrading")
	CronStatusRebalancing = CronStatus("rebalancing")
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
		"locking",
		"upgrading",
		"rebalancing",
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
	DomainStatusUnknown = DomainStatus("unknown")
	// Ready status.
	DomainStatusReady = DomainStatus("ready")
	// Deleting status.
	DomainStatusDeleting = DomainStatus("deleting")
	// Error status.
	DomainStatusError = DomainStatus("error")
	// Creating status. Resource is being created.
	DomainStatusCreating = DomainStatus("creating")
	// Pending status. Resource is being deployed after its creation or an update.
	DomainStatusPending = DomainStatus("pending")
	// Locked status. Resource cannot be modified.
	DomainStatusLocked = DomainStatus("locked")
	// Locking status.
	DomainStatusLocking = DomainStatus("locking")
	// Upgrading status. Resource is being upgraded as part of a planned maintenance. No downtime is expected.
	DomainStatusUpgrading = DomainStatus("upgrading")
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
		"locked",
		"locking",
		"upgrading",
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

type ListContainersRequestOrderBy string

const (
	ListContainersRequestOrderByCreatedAtAsc  = ListContainersRequestOrderBy("created_at_asc")
	ListContainersRequestOrderByCreatedAtDesc = ListContainersRequestOrderBy("created_at_desc")
	ListContainersRequestOrderByNameAsc       = ListContainersRequestOrderBy("name_asc")
	ListContainersRequestOrderByNameDesc      = ListContainersRequestOrderBy("name_desc")
)

func (enum ListContainersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListContainersRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListContainersRequestOrderBy) Values() []ListContainersRequestOrderBy {
	return []ListContainersRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListContainersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListContainersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListContainersRequestOrderBy(ListContainersRequestOrderBy(tmp).String())
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
	NamespaceStatusUnknown = NamespaceStatus("unknown")
	// Ready status.
	NamespaceStatusReady = NamespaceStatus("ready")
	// Deleting status.
	NamespaceStatusDeleting = NamespaceStatus("deleting")
	// Error status.
	NamespaceStatusError = NamespaceStatus("error")
	// Locked status. Resource cannot be modified.
	NamespaceStatusLocked = NamespaceStatus("locked")
	// Creating status. Resource is being created.
	NamespaceStatusCreating = NamespaceStatus("creating")
	// Pending status. Resource is being deployed after its creation or an update.
	NamespaceStatusPending = NamespaceStatus("pending")
	// Locking status.
	NamespaceStatusLocking = NamespaceStatus("locking")
	// Upgrading status. Resource is being upgraded as part of a planned maintenance. No downtime is expected.
	NamespaceStatusUpgrading = NamespaceStatus("upgrading")
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
		"locking",
		"upgrading",
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
	// Creating status. Resource is being created.
	TriggerStatusCreating = TriggerStatus("creating")
	// Pending status. Resource is being deployed after its creation or an update.
	TriggerStatusPending = TriggerStatus("pending")
	// Locked status. Resource cannot be modified.
	TriggerStatusLocked = TriggerStatus("locked")
	// Locking status.
	TriggerStatusLocking = TriggerStatus("locking")
	// Upgrading status. Resource is being upgraded as part of a planned maintenance. No downtime is expected.
	TriggerStatusUpgrading = TriggerStatus("upgrading")
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
		"locked",
		"locking",
		"upgrading",
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

// ContainerHealthCheckSpecHTTPProbe: container health check spec http probe.
type ContainerHealthCheckSpecHTTPProbe struct {
	// Path: path to use for the HTTP health check.
	Path string `json:"path"`
}

// ContainerHealthCheckSpecTCPProbe: container health check spec tcp probe.
type ContainerHealthCheckSpecTCPProbe struct{}

// ContainerHealthCheckSpec: container health check spec.
type ContainerHealthCheckSpec struct {
	// HTTP: HTTP health check configuration.
	// Precisely one of HTTP, TCP must be set.
	HTTP *ContainerHealthCheckSpecHTTPProbe `json:"http,omitempty"`

	// TCP: TCP health check configuration.
	// Precisely one of HTTP, TCP must be set.
	TCP *ContainerHealthCheckSpecTCPProbe `json:"tcp,omitempty"`

	// FailureThreshold: during a deployment, if a newly created container fails to pass the health check, the deployment is aborted.
	// As a result, lowering this value can help to reduce the time it takes to detect a failed deployment.
	FailureThreshold uint32 `json:"failure_threshold"`

	// Interval: period between health checks.
	Interval *scw.Duration `json:"interval"`
}

// ContainerScalingOption: container scaling option.
type ContainerScalingOption struct {
	// Precisely one of ConcurrentRequestsThreshold, CPUUsageThreshold, MemoryUsageThreshold must be set.
	ConcurrentRequestsThreshold *uint32 `json:"concurrent_requests_threshold,omitempty"`

	// Precisely one of ConcurrentRequestsThreshold, CPUUsageThreshold, MemoryUsageThreshold must be set.
	CPUUsageThreshold *uint32 `json:"cpu_usage_threshold,omitempty"`

	// Precisely one of ConcurrentRequestsThreshold, CPUUsageThreshold, MemoryUsageThreshold must be set.
	MemoryUsageThreshold *uint32 `json:"memory_usage_threshold,omitempty"`
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

// Container: container.
type Container struct {
	// ID: UUID of the container.
	ID string `json:"id"`

	// Name: name of the container.
	Name string `json:"name"`

	// NamespaceID: UUID of the namespace the container belongs to.
	NamespaceID string `json:"namespace_id"`

	// Status: status of the container.
	// Default value: unknown
	Status ContainerStatus `json:"status"`

	// EnvironmentVariables: environment variables of the container.
	EnvironmentVariables map[string]string `json:"environment_variables"`

	// MinScale: minimum number of instances to scale the container to.
	MinScale uint32 `json:"min_scale"`

	// MaxScale: maximum number of instances to scale the container to.
	MaxScale uint32 `json:"max_scale"`

	// MemoryLimit: memory limit of the container in MB.
	MemoryLimit uint32 `json:"memory_limit"`

	// CPULimit: CPU limit of the container in mvCPU.
	CPULimit uint32 `json:"cpu_limit"`

	// Timeout: processing time limit for the container.
	Timeout *scw.Duration `json:"timeout"`

	// ErrorMessage: last error message of the container.
	ErrorMessage *string `json:"error_message"`

	// Privacy: privacy setting of the container.
	// Default value: unknown_privacy
	Privacy ContainerPrivacy `json:"privacy"`

	// Description: description of the container.
	Description *string `json:"description"`

	// RegistryImage: name of the registry image (e.g. "rg.fr-par.scw.cloud/something/image:tag").
	RegistryImage string `json:"registry_image"`

	// MaxConcurrency: number of maximum concurrent executions of the container.
	MaxConcurrency uint32 `json:"max_concurrency"`

	// DomainName: domain name attributed to the contaioner.
	DomainName string `json:"domain_name"`

	// Protocol: protocol the container uses.
	// Default value: unknown_protocol
	Protocol ContainerProtocol `json:"protocol"`

	// Port: port the container listens on.
	Port uint32 `json:"port"`

	// SecretEnvironmentVariables: secret environment variables of the container.
	SecretEnvironmentVariables []*SecretHashedValue `json:"secret_environment_variables"`

	// HTTPOption: possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	// Default value: unknown_http_option
	HTTPOption ContainerHTTPOption `json:"http_option"`

	// Sandbox: execution environment of the container.
	// Default value: unknown_sandbox
	Sandbox ContainerSandbox `json:"sandbox"`

	// LocalStorageLimit: local storage limit of the container (in MB).
	LocalStorageLimit uint32 `json:"local_storage_limit"`

	// ScalingOption: possible values:
	// - concurrent_requests_threshold: Scale depending on the number of concurrent requests being processed per container instance.
	// - cpu_usage_threshold: Scale depending on the CPU usage of a container instance.
	// - memory_usage_threshold: Scale depending on the memory usage of a container instance.
	ScalingOption *ContainerScalingOption `json:"scaling_option"`

	// HealthCheck: health check configuration of the container.
	HealthCheck *ContainerHealthCheckSpec `json:"health_check"`

	// CreatedAt: creation date of the container.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date of the container.
	UpdatedAt *time.Time `json:"updated_at"`

	// ReadyAt: last date when the container was successfully deployed and set to ready.
	ReadyAt *time.Time `json:"ready_at"`

	// Region: region in which the container will be deployed.
	Region scw.Region `json:"region"`

	// Tags: list of tags applied to the Serverless Container.
	Tags []string `json:"tags"`

	// PrivateNetworkID: when connected to a Private Network, the container can access other Scaleway resources in this Private Network.
	PrivateNetworkID *string `json:"private_network_id"`

	// Command: command executed when the container starts. This overrides the default command defined in the container image. This is usually the main executable, or entry point script to run.
	Command []string `json:"command"`

	// Args: arguments passed to the command specified in the "command" field. These override the default arguments from the container image, and behave like command-line parameters.
	Args []string `json:"args"`
}

// Cron: cron.
type Cron struct {
	// ID: UUID of the cron.
	ID string `json:"id"`

	// ContainerID: UUID of the container invoked by this cron.
	ContainerID string `json:"container_id"`

	// Schedule: uNIX cron schedule.
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

	// Hostname: domain assigned to the container.
	Hostname string `json:"hostname"`

	// ContainerID: UUID of the container.
	ContainerID string `json:"container_id"`

	// URL: URL (TBD).
	URL string `json:"url"`

	// Status: status of the domain.
	// Default value: unknown
	Status DomainStatus `json:"status"`

	// ErrorMessage: last error message of the domain.
	ErrorMessage *string `json:"error_message"`
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

	// ErrorMessage: last error message of the namesace.
	ErrorMessage *string `json:"error_message"`

	// RegistryEndpoint: registry endpoint of the namespace.
	RegistryEndpoint string `json:"registry_endpoint"`

	// Description: description of the endpoint.
	Description *string `json:"description"`

	// SecretEnvironmentVariables: secret environment variables of the namespace.
	SecretEnvironmentVariables []*SecretHashedValue `json:"secret_environment_variables"`

	// Region: region in which the namespace will be created.
	Region scw.Region `json:"region"`

	// Tags: list of tags applied to the Serverless Container Namespace.
	Tags []string `json:"tags"`

	// CreatedAt: creation date of the namespace.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date of the namespace.
	UpdatedAt *time.Time `json:"updated_at"`

	// Deprecated: VpcIntegrationActivated: the value of this field doesn't matter anymore, and will be removed in a near future.
	VpcIntegrationActivated *bool `json:"vpc_integration_activated,omitempty"`
}

// Token: token.
type Token struct {
	// ID: UUID of the token.
	ID string `json:"id"`

	// Token: identifier of the token.
	Token string `json:"token"`

	// ContainerID: UUID of the container the token belongs to.
	// Precisely one of ContainerID, NamespaceID must be set.
	ContainerID *string `json:"container_id,omitempty"`

	// NamespaceID: UUID of the namespace the token belongs to.
	// Precisely one of ContainerID, NamespaceID must be set.
	NamespaceID *string `json:"namespace_id,omitempty"`

	// Deprecated: PublicKey: public key of the token.
	PublicKey *string `json:"public_key,omitempty"`

	// Status: status of the token.
	// Default value: unknown
	Status TokenStatus `json:"status"`

	// Description: description of the token.
	Description *string `json:"description"`

	// ExpiresAt: expiry date of the token.
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

	// ContainerID: ID of the container to trigger.
	ContainerID string `json:"container_id"`

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

// CreateContainerRequest: create container request.
type CreateContainerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NamespaceID: UUID of the namespace the container belongs to.
	NamespaceID string `json:"namespace_id"`

	// Name: name of the container.
	Name string `json:"name"`

	// EnvironmentVariables: environment variables of the container.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// MinScale: minimum number of instances to scale the container to.
	MinScale *uint32 `json:"min_scale,omitempty"`

	// MaxScale: maximum number of instances to scale the container to.
	MaxScale *uint32 `json:"max_scale,omitempty"`

	// MemoryLimit: memory limit of the container in MB.
	MemoryLimit *uint32 `json:"memory_limit,omitempty"`

	// CPULimit: CPU limit of the container in mvCPU.
	CPULimit *uint32 `json:"cpu_limit,omitempty"`

	// Timeout: processing time limit for the container.
	Timeout *scw.Duration `json:"timeout,omitempty"`

	// Privacy: privacy setting of the container.
	// Default value: unknown_privacy
	Privacy ContainerPrivacy `json:"privacy"`

	// Description: description of the container.
	Description *string `json:"description,omitempty"`

	// RegistryImage: name of the registry image (e.g. "rg.fr-par.scw.cloud/something/image:tag").
	RegistryImage *string `json:"registry_image,omitempty"`

	// Deprecated: MaxConcurrency: number of maximum concurrent executions of the container.
	MaxConcurrency *uint32 `json:"max_concurrency,omitempty"`

	// Protocol: protocol the container uses.
	// Default value: unknown_protocol
	Protocol ContainerProtocol `json:"protocol"`

	// Port: port the container listens on.
	Port *uint32 `json:"port,omitempty"`

	// SecretEnvironmentVariables: secret environment variables of the container.
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`

	// HTTPOption: possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	// Default value: unknown_http_option
	HTTPOption ContainerHTTPOption `json:"http_option"`

	// Sandbox: execution environment of the container.
	// Default value: unknown_sandbox
	Sandbox ContainerSandbox `json:"sandbox"`

	// LocalStorageLimit: local storage limit of the container (in MB).
	LocalStorageLimit *uint32 `json:"local_storage_limit,omitempty"`

	// ScalingOption: possible values:
	// - concurrent_requests_threshold: Scale depending on the number of concurrent requests being processed per container instance.
	// - cpu_usage_threshold: Scale depending on the CPU usage of a container instance.
	// - memory_usage_threshold: Scale depending on the memory usage of a container instance.
	ScalingOption *ContainerScalingOption `json:"scaling_option,omitempty"`

	// HealthCheck: health check configuration of the container.
	HealthCheck *ContainerHealthCheckSpec `json:"health_check,omitempty"`

	// Tags: tags of the Serverless Container.
	Tags []string `json:"tags"`

	// PrivateNetworkID: when connected to a Private Network, the container can access other Scaleway resources in this Private Network.
	PrivateNetworkID *string `json:"private_network_id,omitempty"`

	// Command: command executed when the container starts. This overrides the default command defined in the container image. This is usually the main executable, or entry point script to run.
	Command []string `json:"command"`

	// Args: arguments passed to the command specified in the "command" field. These override the default arguments from the container image, and behave like command-line parameters.
	Args []string `json:"args"`
}

// CreateCronRequest: create cron request.
type CreateCronRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ContainerID: UUID of the container to invoke by the cron.
	ContainerID string `json:"container_id"`

	// Schedule: uNIX cron schedule.
	Schedule string `json:"schedule"`

	// Args: arguments to pass with the cron.
	Args *scw.JSONObject `json:"args,omitempty"`

	// Name: name of the cron to create.
	Name *string `json:"name,omitempty"`
}

// CreateDomainRequest: create domain request.
type CreateDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Hostname: domain to assign.
	Hostname string `json:"hostname"`

	// ContainerID: UUID of the container to assign the domain to.
	ContainerID string `json:"container_id"`
}

// CreateNamespaceRequest: create namespace request.
type CreateNamespaceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: name of the namespace to create.
	Name string `json:"name"`

	// EnvironmentVariables: environment variables of the namespace to create.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// ProjectID: UUID of the Project in which the namespace will be created.
	ProjectID string `json:"project_id"`

	// Description: description of the namespace to create.
	Description *string `json:"description,omitempty"`

	// SecretEnvironmentVariables: secret environment variables of the namespace to create.
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`

	// Tags: tags of the Serverless Container Namespace.
	Tags []string `json:"tags"`

	// Deprecated: ActivateVpcIntegration: setting this field to true doesn't matter anymore. It will be removed in a near future.
	ActivateVpcIntegration *bool `json:"activate_vpc_integration,omitempty"`
}

// CreateTokenRequest: create token request.
type CreateTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ContainerID: UUID of the container to create the token for.
	// Precisely one of ContainerID, NamespaceID must be set.
	ContainerID *string `json:"container_id,omitempty"`

	// NamespaceID: UUID of the namespace to create the token for.
	// Precisely one of ContainerID, NamespaceID must be set.
	NamespaceID *string `json:"namespace_id,omitempty"`

	// Description: description of the token.
	Description *string `json:"description,omitempty"`

	// ExpiresAt: expiry date of the token.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// CreateTriggerRequest: create trigger request.
type CreateTriggerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: name of the trigger.
	Name string `json:"name"`

	// ContainerID: ID of the container to trigger.
	ContainerID string `json:"container_id"`

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

// DeleteContainerRequest: delete container request.
type DeleteContainerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ContainerID: UUID of the container to delete.
	ContainerID string `json:"-"`
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

// DeleteNamespaceRequest: delete namespace request.
type DeleteNamespaceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NamespaceID: UUID of the namespace to delete.
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

// DeployContainerRequest: deploy container request.
type DeployContainerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ContainerID: UUID of the container to deploy.
	ContainerID string `json:"-"`
}

// GetContainerRequest: get container request.
type GetContainerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ContainerID: UUID of the container to get.
	ContainerID string `json:"-"`
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

// GetNamespaceRequest: get namespace request.
type GetNamespaceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NamespaceID: UUID of the namespace to get.
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

// ListContainersRequest: list containers request.
type ListContainersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: number of containers per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the containers.
	// Default value: created_at_asc
	OrderBy ListContainersRequestOrderBy `json:"-"`

	// NamespaceID: UUID of the namespace the container belongs to.
	NamespaceID string `json:"-"`

	// Name: name of the container.
	Name *string `json:"-"`

	// OrganizationID: UUID of the Organization the container belongs to.
	OrganizationID *string `json:"-"`

	// ProjectID: UUID of the Project the container belongs to.
	ProjectID *string `json:"-"`
}

// ListContainersResponse: list containers response.
type ListContainersResponse struct {
	// Containers: array of containers.
	Containers []*Container `json:"containers"`

	// TotalCount: total number of containers.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListContainersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListContainersResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListContainersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Containers = append(r.Containers, results.Containers...)
	r.TotalCount += uint32(len(results.Containers))
	return uint32(len(results.Containers)), nil
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

	// ContainerID: UUID of the container invoked by the cron.
	ContainerID string `json:"-"`
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

	// ContainerID: UUID of the container the domain belongs to.
	ContainerID string `json:"-"`
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

	// Name: name of the namespaces.
	Name *string `json:"-"`

	// OrganizationID: UUID of the Organization the namespace belongs to.
	OrganizationID *string `json:"-"`

	// ProjectID: UUID of the Project the namespace belongs to.
	ProjectID *string `json:"-"`
}

// ListNamespacesResponse: list namespaces response.
type ListNamespacesResponse struct {
	// Namespaces: array of the namespaces.
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

	// OrderBy: order of the tokens.
	// Default value: created_at_asc
	OrderBy ListTokensRequestOrderBy `json:"-"`

	// ContainerID: UUID of the container the token belongs to.
	ContainerID *string `json:"-"`

	// NamespaceID: UUID of the namespace the token belongs to.
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

	// ContainerID: ID of the container the triggers belongs to.
	// Precisely one of ContainerID, NamespaceID, ProjectID must be set.
	ContainerID *string `json:"container_id,omitempty"`

	// NamespaceID: ID of the namespace the triggers belongs to.
	// Precisely one of ContainerID, NamespaceID, ProjectID must be set.
	NamespaceID *string `json:"namespace_id,omitempty"`

	// ProjectID: ID of the project the triggers belongs to.
	// Precisely one of ContainerID, NamespaceID, ProjectID must be set.
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

// UpdateContainerRequest: update container request.
type UpdateContainerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ContainerID: UUID of the container to update.
	ContainerID string `json:"-"`

	// EnvironmentVariables: environment variables of the container.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// MinScale: minimum number of instances to scale the container to.
	MinScale *uint32 `json:"min_scale,omitempty"`

	// MaxScale: maximum number of instances to scale the container to.
	MaxScale *uint32 `json:"max_scale,omitempty"`

	// MemoryLimit: memory limit of the container in MB.
	MemoryLimit *uint32 `json:"memory_limit,omitempty"`

	// CPULimit: CPU limit of the container in mvCPU.
	CPULimit *uint32 `json:"cpu_limit,omitempty"`

	// Timeout: processing time limit for the container.
	Timeout *scw.Duration `json:"timeout,omitempty"`

	// Deprecated: Redeploy: deprecated: future versions of this API will systematically redeploy containers when needed. As such,
	// passing `redeploy: false` will be ignored. Relying on this field is discouraged.
	//
	// To force the redeployment of a container, even if no configuration has changed, use the `DeployContainer` method instead.
	Redeploy *bool `json:"redeploy,omitempty"`

	// Privacy: privacy settings of the container.
	// Default value: unknown_privacy
	Privacy ContainerPrivacy `json:"privacy"`

	// Description: description of the container.
	Description *string `json:"description,omitempty"`

	// RegistryImage: name of the registry image (e.g. "rg.fr-par.scw.cloud/something/image:tag").
	RegistryImage *string `json:"registry_image,omitempty"`

	// Deprecated: MaxConcurrency: number of maximum concurrent executions of the container.
	MaxConcurrency *uint32 `json:"max_concurrency,omitempty"`

	// Protocol: protocol the container uses.
	// Default value: unknown_protocol
	Protocol ContainerProtocol `json:"protocol"`

	// Port: port the container listens on.
	Port *uint32 `json:"port,omitempty"`

	// SecretEnvironmentVariables: during an update, secret environment variables that are not specified in this field will be kept unchanged.
	//
	// In order to delete a specific secret environment variable, you must reference its key, but not provide any value for it.
	// For example, the following payload will delete the `TO_DELETE` secret environment variable:
	//
	// ```json
	// {
	//   "secret_environment_variables":[
	//     {"key":"TO_DELETE"}
	//   ]
	// }
	// ```.
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`

	// HTTPOption: possible values:
	//  - redirected: Responds to HTTP request with a 301 redirect to ask the clients to use HTTPS.
	//  - enabled: Serve both HTTP and HTTPS traffic.
	// Default value: unknown_http_option
	HTTPOption ContainerHTTPOption `json:"http_option"`

	// Sandbox: execution environment of the container.
	// Default value: unknown_sandbox
	Sandbox ContainerSandbox `json:"sandbox"`

	// LocalStorageLimit: local storage limit of the container (in MB).
	LocalStorageLimit *uint32 `json:"local_storage_limit,omitempty"`

	// ScalingOption: possible values:
	// - concurrent_requests_threshold: Scale depending on the number of concurrent requests being processed per container instance.
	// - cpu_usage_threshold: Scale depending on the CPU usage of a container instance.
	// - memory_usage_threshold: Scale depending on the memory usage of a container instance.
	ScalingOption *ContainerScalingOption `json:"scaling_option,omitempty"`

	// HealthCheck: health check configuration of the container.
	HealthCheck *ContainerHealthCheckSpec `json:"health_check,omitempty"`

	// Tags: tags of the Serverless Container.
	Tags *[]string `json:"tags,omitempty"`

	// PrivateNetworkID: when connected to a Private Network, the container can access other Scaleway resources in this Private Network.
	PrivateNetworkID *string `json:"private_network_id,omitempty"`

	// Command: command executed when the container starts. This overrides the default command defined in the container image. This is usually the main executable, or entry point script to run.
	Command *[]string `json:"command,omitempty"`

	// Args: arguments passed to the command specified in the "command" field. These override the default arguments from the container image, and behave like command-line parameters.
	Args *[]string `json:"args,omitempty"`
}

// UpdateCronRequest: update cron request.
type UpdateCronRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// CronID: UUID of the cron to update.
	CronID string `json:"-"`

	// ContainerID: UUID of the container invoked by the cron.
	ContainerID *string `json:"container_id,omitempty"`

	// Schedule: uNIX cron schedule.
	Schedule *string `json:"schedule,omitempty"`

	// Args: arguments to pass with the cron.
	Args *scw.JSONObject `json:"args,omitempty"`

	// Name: name of the cron.
	Name *string `json:"name,omitempty"`
}

// UpdateNamespaceRequest: update namespace request.
type UpdateNamespaceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NamespaceID: UUID of the namespace to update.
	NamespaceID string `json:"-"`

	// EnvironmentVariables: environment variables of the namespace to update.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// Description: description of the namespace to update.
	Description *string `json:"description,omitempty"`

	// SecretEnvironmentVariables: secret environment variables of the namespace to update.
	SecretEnvironmentVariables []*Secret `json:"secret_environment_variables"`

	// Tags: tags of the Serverless Container Namespace.
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

// This API allows you to manage your Serverless Containers.
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

// ListNamespaces: List all namespaces in a specified region.
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForNamespaceRequest is used by WaitForNamespace method.
type WaitForNamespaceRequest struct {
	Region        scw.Region
	NamespaceID   string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForNamespace waits for the Namespace to reach a terminal state.
func (s *API) WaitForNamespace(req *WaitForNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	timeout := defaultContainerTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultContainerRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[NamespaceStatus]struct{}{
		NamespaceStatusDeleting:  {},
		NamespaceStatusCreating:  {},
		NamespaceStatusPending:   {},
		NamespaceStatusLocking:   {},
		NamespaceStatusUpgrading: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetNamespace(&GetNamespaceRequest{
				Region:      req.Region,
				NamespaceID: req.NamespaceID,
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
		return nil, errors.Wrap(err, "waiting for Namespace failed")
	}

	return res.(*Namespace), nil
}

// CreateNamespace: Create a new namespace in a specified region.
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
		req.Name = namegenerator.GetRandomName("cns")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
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

// UpdateNamespace: Update the space associated with the specified ID.
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListContainers: List all containers for a specified region.
func (s *API) ListContainers(req *ListContainersRequest, opts ...scw.RequestOption) (*ListContainersResponse, error) {
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers",
		Query:  query,
	}

	var resp ListContainersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetContainer: Get the container associated with the specified ID.
func (s *API) GetContainer(req *GetContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "",
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForContainerRequest is used by WaitForContainer method.
type WaitForContainerRequest struct {
	Region        scw.Region
	ContainerID   string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForContainer waits for the Container to reach a terminal state.
func (s *API) WaitForContainer(req *WaitForContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	timeout := defaultContainerTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultContainerRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[ContainerStatus]struct{}{
		ContainerStatusDeleting:  {},
		ContainerStatusCreating:  {},
		ContainerStatusPending:   {},
		ContainerStatusLocking:   {},
		ContainerStatusUpgrading: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetContainer(&GetContainerRequest{
				Region:      req.Region,
				ContainerID: req.ContainerID,
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
		return nil, errors.Wrap(err, "waiting for Container failed")
	}

	return res.(*Container), nil
}

// CreateContainer: Create a new container in the specified region.
func (s *API) CreateContainer(req *CreateContainerRequest, opts ...scw.RequestOption) (*Container, error) {
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateContainer: Update the container associated with the specified ID.
//
// When updating a container, the container is automatically redeployed to apply the changes.
// This behavior can be changed by setting the `redeploy` field to `false` in the request.
func (s *API) UpdateContainer(req *UpdateContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteContainer: Delete the container associated with the specified ID.
func (s *API) DeleteContainer(req *DeleteContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "",
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeployContainer: Deploy a container associated with the specified ID.
func (s *API) DeployContainer(req *DeployContainerRequest, opts ...scw.RequestOption) (*Container, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ContainerID) == "" {
		return nil, errors.New("field ContainerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "/deploy",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCrons: List all your crons.
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
	parameter.AddToQuery(query, "container_id", req.ContainerID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForCronRequest is used by WaitForCron method.
type WaitForCronRequest struct {
	Region        scw.Region
	CronID        string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForCron waits for the Cron to reach a terminal state.
func (s *API) WaitForCron(req *WaitForCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	timeout := defaultContainerTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultContainerRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[CronStatus]struct{}{
		CronStatusDeleting:    {},
		CronStatusCreating:    {},
		CronStatusPending:     {},
		CronStatusLocking:     {},
		CronStatusUpgrading:   {},
		CronStatusRebalancing: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetCron(&GetCronRequest{
				Region: req.Region,
				CronID: req.CronID,
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
		return nil, errors.Wrap(err, "waiting for Cron failed")
	}

	return res.(*Cron), nil
}

// CreateCron: Create a new cron.
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDomains: List all custom domains in a specified region.
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
	parameter.AddToQuery(query, "container_id", req.ContainerID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Query:  query,
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomain: Get a custom domain for the container with the specified ID.
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForDomainRequest is used by WaitForDomain method.
type WaitForDomainRequest struct {
	Region        scw.Region
	DomainID      string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForDomain waits for the Domain to reach a terminal state.
func (s *API) WaitForDomain(req *WaitForDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	timeout := defaultContainerTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultContainerRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[DomainStatus]struct{}{
		DomainStatusDeleting:  {},
		DomainStatusCreating:  {},
		DomainStatusPending:   {},
		DomainStatusLocking:   {},
		DomainStatusUpgrading: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetDomain(&GetDomainRequest{
				Region:   req.Region,
				DomainID: req.DomainID,
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
		return nil, errors.Wrap(err, "waiting for Domain failed")
	}

	return res.(*Domain), nil
}

// CreateDomain: Create a custom domain for the container with the specified ID.
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains",
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

// DeleteDomain: Delete the custom domain with the specific ID.
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: CreateToken: Deprecated in favor of IAM authentication.
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens",
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

// GetToken: Get a token with a specified ID.
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForTokenRequest is used by WaitForToken method.
type WaitForTokenRequest struct {
	Region        scw.Region
	TokenID       string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForToken waits for the Token to reach a terminal state.
func (s *API) WaitForToken(req *WaitForTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	timeout := defaultContainerTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultContainerRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[TokenStatus]struct{}{
		TokenStatusDeleting: {},
		TokenStatusCreating: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetToken(&GetTokenRequest{
				Region:  req.Region,
				TokenID: req.TokenID,
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
		return nil, errors.Wrap(err, "waiting for Token failed")
	}

	return res.(*Token), nil
}

// ListTokens: List all tokens belonging to a specified Organization or Project.
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
	parameter.AddToQuery(query, "container_id", req.ContainerID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens",
		Query:  query,
	}

	var resp ListTokensResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteToken: Delete a token with a specified ID.
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTrigger: Create a new trigger for a specified container.
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers",
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForTriggerRequest is used by WaitForTrigger method.
type WaitForTriggerRequest struct {
	Region        scw.Region
	TriggerID     string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForTrigger waits for the Trigger to reach a terminal state.
func (s *API) WaitForTrigger(req *WaitForTriggerRequest, opts ...scw.RequestOption) (*Trigger, error) {
	timeout := defaultContainerTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultContainerRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[TriggerStatus]struct{}{
		TriggerStatusDeleting:  {},
		TriggerStatusCreating:  {},
		TriggerStatusPending:   {},
		TriggerStatusLocking:   {},
		TriggerStatusUpgrading: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetTrigger(&GetTriggerRequest{
				Region:    req.Region,
				TriggerID: req.TriggerID,
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
		return nil, errors.Wrap(err, "waiting for Trigger failed")
	}

	return res.(*Trigger), nil
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
	if exist && req.ContainerID == nil && req.NamespaceID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "container_id", req.ContainerID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers",
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
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
		Path:   "/containers/v1beta1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

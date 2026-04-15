// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package container provides methods and message types of the container v1 API.
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
	defaultContainerTimeout       = 5 * time.Minute
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

type ContainerPrivacy string

const (
	// Unknown privacy.
	ContainerPrivacyUnknownPrivacy = ContainerPrivacy("unknown_privacy")
	// Container is public, and can be accessed without authentication.
	ContainerPrivacyPublic = ContainerPrivacy("public")
	// Container is private, and can only be accessed by providing an authentication token (IAM or legacy JWTs).
	ContainerPrivacyPrivate = ContainerPrivacy("private")
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
	// Unknown protocol.
	ContainerProtocolUnknownProtocol = ContainerProtocol("unknown_protocol")
	// HTTP/1 protocol. Recommended for most use cases.
	ContainerProtocolHTTP1 = ContainerProtocol("http1")
	// HTTP/2 protocol. Useful for gRPC applications. Requires that the container explicitly supports HTTP/2 (e.g. Nginx's `listen 80 http2` directive).
	ContainerProtocolH2c = ContainerProtocol("h2c")
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
	// Legacy sandbox with slower cold starts. Fully supports the Linux system call interface.
	ContainerSandboxV1 = ContainerSandbox("v1")
	// Recommended sandbox with faster cold starts.
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
	// Unknown status.
	ContainerStatusUnknownStatus = ContainerStatus("unknown_status")
	// Updating status. Resource is being redeployed to match the desired configuration.
	ContainerStatusUpdating = ContainerStatus("updating")
	// Deleting status.
	ContainerStatusDeleting = ContainerStatus("deleting")
	// Locking status.
	ContainerStatusLocking = ContainerStatus("locking")
	// Ready status.
	ContainerStatusReady = ContainerStatus("ready")
	// Error status.
	ContainerStatusError = ContainerStatus("error")
	// Locked status. Resource cannot be modified.
	ContainerStatusLocked = ContainerStatus("locked")
	// Creating status. Resource is being created.
	ContainerStatusCreating = ContainerStatus("creating")
	// Upgrading status. Resource is being upgraded as part of a planned maintenance. No downtime is expected.
	ContainerStatusUpgrading = ContainerStatus("upgrading")
)

func (enum ContainerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ContainerStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ContainerStatus) Values() []ContainerStatus {
	return []ContainerStatus{
		"unknown_status",
		"updating",
		"deleting",
		"locking",
		"ready",
		"error",
		"locked",
		"creating",
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

type CreateTriggerRequestDestinationConfigHTTPMethod string

const (
	CreateTriggerRequestDestinationConfigHTTPMethodUnknownHTTPMethod = CreateTriggerRequestDestinationConfigHTTPMethod("unknown_http_method")
	CreateTriggerRequestDestinationConfigHTTPMethodGet               = CreateTriggerRequestDestinationConfigHTTPMethod("get")
	CreateTriggerRequestDestinationConfigHTTPMethodPost              = CreateTriggerRequestDestinationConfigHTTPMethod("post")
	CreateTriggerRequestDestinationConfigHTTPMethodPut               = CreateTriggerRequestDestinationConfigHTTPMethod("put")
	CreateTriggerRequestDestinationConfigHTTPMethodPatch             = CreateTriggerRequestDestinationConfigHTTPMethod("patch")
	CreateTriggerRequestDestinationConfigHTTPMethodDelete            = CreateTriggerRequestDestinationConfigHTTPMethod("delete")
)

func (enum CreateTriggerRequestDestinationConfigHTTPMethod) String() string {
	if enum == "" {
		// return default value if empty
		return string(CreateTriggerRequestDestinationConfigHTTPMethodUnknownHTTPMethod)
	}
	return string(enum)
}

func (enum CreateTriggerRequestDestinationConfigHTTPMethod) Values() []CreateTriggerRequestDestinationConfigHTTPMethod {
	return []CreateTriggerRequestDestinationConfigHTTPMethod{
		"unknown_http_method",
		"get",
		"post",
		"put",
		"patch",
		"delete",
	}
}

func (enum CreateTriggerRequestDestinationConfigHTTPMethod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CreateTriggerRequestDestinationConfigHTTPMethod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CreateTriggerRequestDestinationConfigHTTPMethod(CreateTriggerRequestDestinationConfigHTTPMethod(tmp).String())
	return nil
}

type DomainStatus string

const (
	// Unknown status.
	DomainStatusUnknownStatus = DomainStatus("unknown_status")
	// Creating status. Resource is being created.
	DomainStatusCreating = DomainStatus("creating")
	// Updating status. Resource is being redeployed to match the desired configuration.
	DomainStatusUpdating = DomainStatus("updating")
	// Deleting status.
	DomainStatusDeleting = DomainStatus("deleting")
	// Ready status.
	DomainStatusReady = DomainStatus("ready")
	// Error status.
	DomainStatusError = DomainStatus("error")
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
		return string(DomainStatusUnknownStatus)
	}
	return string(enum)
}

func (enum DomainStatus) Values() []DomainStatus {
	return []DomainStatus{
		"unknown_status",
		"creating",
		"updating",
		"deleting",
		"ready",
		"error",
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

type ListTriggersRequestOrderBy string

const (
	ListTriggersRequestOrderByCreatedAtAsc  = ListTriggersRequestOrderBy("created_at_asc")
	ListTriggersRequestOrderByCreatedAtDesc = ListTriggersRequestOrderBy("created_at_desc")
	ListTriggersRequestOrderByNameAsc       = ListTriggersRequestOrderBy("name_asc")
	ListTriggersRequestOrderByNameDesc      = ListTriggersRequestOrderBy("name_desc")
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
		"name_asc",
		"name_desc",
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
	// Unknown status.
	NamespaceStatusUnknownStatus = NamespaceStatus("unknown_status")
	// Updating status. Resource is being redeployed to match the desired configuration.
	NamespaceStatusUpdating = NamespaceStatus("updating")
	// Deleting status.
	NamespaceStatusDeleting = NamespaceStatus("deleting")
	// Locking status.
	NamespaceStatusLocking = NamespaceStatus("locking")
	// Ready status.
	NamespaceStatusReady = NamespaceStatus("ready")
	// Error status.
	NamespaceStatusError = NamespaceStatus("error")
	// Locked status. Resource cannot be modified.
	NamespaceStatusLocked = NamespaceStatus("locked")
	// Creating status. Resource is being created.
	NamespaceStatusCreating = NamespaceStatus("creating")
	// Upgrading status. Resource is being upgraded as part of a planned maintenance. No downtime is expected.
	NamespaceStatusUpgrading = NamespaceStatus("upgrading")
)

func (enum NamespaceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(NamespaceStatusUnknownStatus)
	}
	return string(enum)
}

func (enum NamespaceStatus) Values() []NamespaceStatus {
	return []NamespaceStatus{
		"unknown_status",
		"updating",
		"deleting",
		"locking",
		"ready",
		"error",
		"locked",
		"creating",
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

type TriggerDestinationConfigHTTPMethod string

const (
	TriggerDestinationConfigHTTPMethodUnknownHTTPMethod = TriggerDestinationConfigHTTPMethod("unknown_http_method")
	TriggerDestinationConfigHTTPMethodGet               = TriggerDestinationConfigHTTPMethod("get")
	TriggerDestinationConfigHTTPMethodPost              = TriggerDestinationConfigHTTPMethod("post")
	TriggerDestinationConfigHTTPMethodPut               = TriggerDestinationConfigHTTPMethod("put")
	TriggerDestinationConfigHTTPMethodPatch             = TriggerDestinationConfigHTTPMethod("patch")
	TriggerDestinationConfigHTTPMethodDelete            = TriggerDestinationConfigHTTPMethod("delete")
)

func (enum TriggerDestinationConfigHTTPMethod) String() string {
	if enum == "" {
		// return default value if empty
		return string(TriggerDestinationConfigHTTPMethodUnknownHTTPMethod)
	}
	return string(enum)
}

func (enum TriggerDestinationConfigHTTPMethod) Values() []TriggerDestinationConfigHTTPMethod {
	return []TriggerDestinationConfigHTTPMethod{
		"unknown_http_method",
		"get",
		"post",
		"put",
		"patch",
		"delete",
	}
}

func (enum TriggerDestinationConfigHTTPMethod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TriggerDestinationConfigHTTPMethod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TriggerDestinationConfigHTTPMethod(TriggerDestinationConfigHTTPMethod(tmp).String())
	return nil
}

type TriggerSourceType string

const (
	// Unknown source type.
	TriggerSourceTypeUnknownSourceType = TriggerSourceType("unknown_source_type")
	// Cron source type.
	TriggerSourceTypeCron = TriggerSourceType("cron")
	// SQS source type.
	TriggerSourceTypeSqs = TriggerSourceType("sqs")
	// NATS source type.
	TriggerSourceTypeNats = TriggerSourceType("nats")
)

func (enum TriggerSourceType) String() string {
	if enum == "" {
		// return default value if empty
		return string(TriggerSourceTypeUnknownSourceType)
	}
	return string(enum)
}

func (enum TriggerSourceType) Values() []TriggerSourceType {
	return []TriggerSourceType{
		"unknown_source_type",
		"cron",
		"sqs",
		"nats",
	}
}

func (enum TriggerSourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TriggerSourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TriggerSourceType(TriggerSourceType(tmp).String())
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
	// Updating status. Resource is being redeployed to match the desired configuration.
	TriggerStatusUpdating = TriggerStatus("updating")
	// Creating status. Resource is being created.
	TriggerStatusCreating = TriggerStatus("creating")
	// Locking status.
	TriggerStatusLocking = TriggerStatus("locking")
	// Locked status. Resource cannot be modified.
	TriggerStatusLocked = TriggerStatus("locked")
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
		"updating",
		"creating",
		"locking",
		"locked",
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

type UpdateTriggerRequestDestinationConfigHTTPMethod string

const (
	UpdateTriggerRequestDestinationConfigHTTPMethodUnknownHTTPMethod = UpdateTriggerRequestDestinationConfigHTTPMethod("unknown_http_method")
	UpdateTriggerRequestDestinationConfigHTTPMethodGet               = UpdateTriggerRequestDestinationConfigHTTPMethod("get")
	UpdateTriggerRequestDestinationConfigHTTPMethodPost              = UpdateTriggerRequestDestinationConfigHTTPMethod("post")
	UpdateTriggerRequestDestinationConfigHTTPMethodPut               = UpdateTriggerRequestDestinationConfigHTTPMethod("put")
	UpdateTriggerRequestDestinationConfigHTTPMethodPatch             = UpdateTriggerRequestDestinationConfigHTTPMethod("patch")
	UpdateTriggerRequestDestinationConfigHTTPMethodDelete            = UpdateTriggerRequestDestinationConfigHTTPMethod("delete")
)

func (enum UpdateTriggerRequestDestinationConfigHTTPMethod) String() string {
	if enum == "" {
		// return default value if empty
		return string(UpdateTriggerRequestDestinationConfigHTTPMethodUnknownHTTPMethod)
	}
	return string(enum)
}

func (enum UpdateTriggerRequestDestinationConfigHTTPMethod) Values() []UpdateTriggerRequestDestinationConfigHTTPMethod {
	return []UpdateTriggerRequestDestinationConfigHTTPMethod{
		"unknown_http_method",
		"get",
		"post",
		"put",
		"patch",
		"delete",
	}
}

func (enum UpdateTriggerRequestDestinationConfigHTTPMethod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UpdateTriggerRequestDestinationConfigHTTPMethod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UpdateTriggerRequestDestinationConfigHTTPMethod(UpdateTriggerRequestDestinationConfigHTTPMethod(tmp).String())
	return nil
}

// ContainerProbeHTTPProbe: container probe http probe.
type ContainerProbeHTTPProbe struct {
	// Path: HTTP path to perform the check on.
	Path string `json:"path"`
}

// ContainerProbeTCPProbe: container probe tcp probe.
type ContainerProbeTCPProbe struct{}

// ContainerProbe: container probe.
type ContainerProbe struct {
	// FailureThreshold: unhealthy containers do not receive traffic from incoming requests.
	FailureThreshold uint32 `json:"failure_threshold"`

	// Interval: time interval between checks.
	Interval *scw.Duration `json:"interval"`

	// Timeout: duration before the check times out.
	Timeout *scw.Duration `json:"timeout"`

	// TCP: the check is successful if a TCP connection can be established with the container within the specified timeout.
	// Precisely one of TCP, HTTP must be set.
	TCP *ContainerProbeTCPProbe `json:"tcp,omitempty"`

	// HTTP: the check is successful if an HTTP request to the specified path returns a successful status code (e.g. 2XX or 3XX) within the specified timeout.
	// Precisely one of TCP, HTTP must be set.
	HTTP *ContainerProbeHTTPProbe `json:"http,omitempty"`
}

// ContainerScalingOption: container scaling option.
type ContainerScalingOption struct {
	// ConcurrentRequestsThreshold: scale depending on the number of concurrent requests being processed per container instance. The threshold value is the number of concurrent requests above which the container will be scaled up.
	// Precisely one of ConcurrentRequestsThreshold, CPUUsageThreshold, MemoryUsageThreshold must be set.
	ConcurrentRequestsThreshold *uint32 `json:"concurrent_requests_threshold,omitempty"`

	// CPUUsageThreshold: scale depending on the CPU usage of a container instance. The threshold value is the percentage of CPU usage above which the container will be scaled up.
	// Precisely one of ConcurrentRequestsThreshold, CPUUsageThreshold, MemoryUsageThreshold must be set.
	CPUUsageThreshold *uint32 `json:"cpu_usage_threshold,omitempty"`

	// MemoryUsageThreshold: scale depending on the memory usage of a container instance. The threshold value is the percentage of memory usage above which the container will be scaled up.
	// Precisely one of ConcurrentRequestsThreshold, CPUUsageThreshold, MemoryUsageThreshold must be set.
	MemoryUsageThreshold *uint32 `json:"memory_usage_threshold,omitempty"`
}

// TriggerCronConfig: trigger cron config.
type TriggerCronConfig struct {
	// Schedule: uNIX cron schedule to run job (e.g., "* * * * *").
	Schedule string `json:"schedule"`

	// Timezone: timezone for the cron schedule, in tz database format (e.g., "Europe/Paris").
	Timezone string `json:"timezone"`

	// Body: body to send to the container when the trigger is invoked.
	Body string `json:"body"`

	// Headers: additional headers to send to the container when the trigger is invoked.
	Headers map[string]string `json:"headers"`
}

// TriggerDestinationConfig: trigger destination config.
type TriggerDestinationConfig struct {
	// HTTPPath: the HTTP path to send the request to (e.g., "/my-webhook-endpoint").
	HTTPPath string `json:"http_path"`

	// HTTPMethod: the HTTP method to use when sending the request (e.g., get, post, put, patch, delete). Must be specified as lowercase.
	// Default value: unknown_http_method
	HTTPMethod TriggerDestinationConfigHTTPMethod `json:"http_method"`
}

// TriggerNATSConfig: trigger nats config.
type TriggerNATSConfig struct {
	// ServerURLs: the URLs of the NATS servers (e.g., "nats://nats.mnq.fr-par.scaleway.com:4222").
	ServerURLs []string `json:"server_urls"`

	// Subject: nATS subject to subscribe to (e.g., "my-subject").
	Subject string `json:"subject"`
}

// TriggerSQSConfig: trigger sqs config.
type TriggerSQSConfig struct {
	// Region: the region where the SQS queue is hosted (e.g., "fr-par", "nl-ams").
	Region scw.Region `json:"region"`

	// Endpoint: endpoint URL to use to access SQS (e.g., "https://sqs.mnq.fr-par.scaleway.com").
	Endpoint string `json:"endpoint"`

	// AccessKeyID: the access key for accessing the SQS queue.
	AccessKeyID string `json:"access_key_id"`

	// QueueURL: the URL of the SQS queue to monitor for messages.
	QueueURL string `json:"queue_url"`
}

// UpdateContainerRequestProbeHTTPProbe: update container request probe http probe.
type UpdateContainerRequestProbeHTTPProbe struct {
	Path *string `json:"path"`
}

// UpdateContainerRequestProbeTCPProbe: update container request probe tcp probe.
type UpdateContainerRequestProbeTCPProbe struct{}

// CreateTriggerRequestCronConfig: create trigger request cron config.
type CreateTriggerRequestCronConfig struct {
	// Schedule: uNIX cron schedule to run job (e.g., "* * * * *").
	Schedule string `json:"schedule"`

	// Timezone: timezone for the cron schedule, in tz database format (e.g., "Europe/Paris").
	Timezone string `json:"timezone"`

	// Body: body to send to the container when the trigger is invoked.
	Body string `json:"body"`

	// Headers: additional headers to send to the container when the trigger is invoked.
	Headers map[string]string `json:"headers"`
}

// CreateTriggerRequestDestinationConfig: create trigger request destination config.
type CreateTriggerRequestDestinationConfig struct {
	// HTTPPath: the HTTP path to send the request to (e.g., "/my-webhook-endpoint").
	HTTPPath string `json:"http_path"`

	// HTTPMethod: the HTTP method to use when sending the request (e.g., get, post, put, patch, delete). Must be specified as lowercase.
	// Default value: unknown_http_method
	HTTPMethod CreateTriggerRequestDestinationConfigHTTPMethod `json:"http_method"`
}

// CreateTriggerRequestNATSConfig: create trigger request nats config.
type CreateTriggerRequestNATSConfig struct {
	// ServerURLs: the URLs of the NATS server (e.g., "nats://nats.mnq.fr-par.scaleway.com:4222").
	ServerURLs []string `json:"server_urls"`

	// Subject: nATS subject to subscribe to (e.g., "my-subject").
	Subject string `json:"subject"`

	// CredentialsFileContent: the credentials from this file will be used to authenticate with the NATS server and subscribe to the specified subject.
	CredentialsFileContent string `json:"credentials_file_content"`
}

// CreateTriggerRequestSQSConfig: create trigger request sqs config.
type CreateTriggerRequestSQSConfig struct {
	// Region: the region where the SQS queue is hosted (e.g., "fr-par", "nl-ams").
	Region scw.Region `json:"region"`

	// Endpoint: endpoint URL to use to access SQS (e.g., "https://sqs.mnq.fr-par.scaleway.com").
	Endpoint string `json:"endpoint"`

	// AccessKeyID: the access key for accessing the SQS queue.
	AccessKeyID string `json:"access_key_id"`

	// SecretAccessKey: the secret key for accessing the SQS queue.
	SecretAccessKey string `json:"secret_access_key"`

	// QueueURL: the URL of the SQS queue to monitor for messages.
	QueueURL string `json:"queue_url"`
}

// Container: container.
type Container struct {
	// ID: container unique ID.
	ID string `json:"id"`

	// Name: container name.
	Name string `json:"name"`

	// NamespaceID: unique ID of the namespace the container belongs to.
	NamespaceID string `json:"namespace_id"`

	// Description: container description.
	Description string `json:"description"`

	// Status: container status.
	// Default value: unknown_status
	Status ContainerStatus `json:"status"`

	// ErrorMessage: container last error message.
	ErrorMessage *string `json:"error_message"`

	// CreatedAt: container creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: container last update date.
	UpdatedAt *time.Time `json:"updated_at"`

	// EnvironmentVariables: environment variables of the container.
	EnvironmentVariables map[string]string `json:"environment_variables"`

	// SecretEnvironmentVariables: secret environment variables of the container.
	SecretEnvironmentVariables map[string]string `json:"secret_environment_variables"`

	// MinScale: minimum number of instances to scale the container to.
	MinScale uint32 `json:"min_scale"`

	// MaxScale: maximum number of instances to scale the container to.
	MaxScale uint32 `json:"max_scale"`

	// MemoryLimitBytes: memory limit of the container in bytes.
	MemoryLimitBytes scw.Size `json:"memory_limit_bytes"`

	// MvcpuLimit: CPU limit of the container in mvCPU.
	MvcpuLimit uint32 `json:"mvcpu_limit"`

	// LocalStorageLimitBytes: local storage limit of the container (in bytes).
	LocalStorageLimitBytes scw.Size `json:"local_storage_limit_bytes"`

	// Timeout: processing time limit for the container.
	Timeout *scw.Duration `json:"timeout"`

	// Privacy: privacy policy of the container.
	// Default value: unknown_privacy
	Privacy ContainerPrivacy `json:"privacy"`

	// Image: image reference (e.g. "rg.fr-par.scw.cloud/my-registry-namespace/image:tag").
	Image string `json:"image"`

	// Protocol: protocol the container uses.
	// Default value: unknown_protocol
	Protocol ContainerProtocol `json:"protocol"`

	// Port: port the container listens on.
	Port uint32 `json:"port"`

	// HTTPSConnectionsOnly: if true, it will allow only HTTPS connections to access your container to prevent it from being triggered by insecure connections (HTTP).
	HTTPSConnectionsOnly bool `json:"https_connections_only"`

	// Sandbox: execution environment of the container.
	// Default value: unknown_sandbox
	Sandbox ContainerSandbox `json:"sandbox"`

	// ScalingOption: parameter used to decide when to scale up or down.
	ScalingOption *ContainerScalingOption `json:"scaling_option"`

	// LivenessProbe: if the liveness probe fails, the container will be restarted.
	// It is performed periodically during the container's lifetime. The liveness probe is not executed until the startup probe (if defined) is successful.
	//
	// Possible check types:
	// - http: Perform HTTP check on the container with the specified path.
	// - tcp: Perform TCP check on the container.
	LivenessProbe *ContainerProbe `json:"liveness_probe"`

	// StartupProbe: if the startup probe fails, the container will be restarted.
	// This check is useful for applications that might take a long time to start. It is only performed when the container is starting.
	//
	// Possible check types:
	// - http: Perform HTTP check on the container with the specified path.
	// - tcp: Perform TCP check on the container.
	StartupProbe *ContainerProbe `json:"startup_probe"`

	// Tags: tags of the Serverless Container.
	Tags []string `json:"tags"`

	// PrivateNetworkID: when connected to a Private Network, the container can access other Scaleway resources in this Private Network.
	PrivateNetworkID *string `json:"private_network_id"`

	// Command: command executed when the container starts. This overrides the default command defined in the container image. This is usually the main executable, or ENTRYPOINT script to run.
	Command []string `json:"command"`

	// Args: arguments passed to the command specified in the "command" field. These override the default arguments from the container image, and behave like command-line parameters.
	Args []string `json:"args"`

	// PublicEndpoint: this is the default endpoint generated by Scaleway to access the container from the Internet.
	PublicEndpoint string `json:"public_endpoint"`

	// Region: region in which the container exists.
	Region scw.Region `json:"region"`
}

// Domain: domain.
type Domain struct {
	// ID: domain unique ID.
	ID string `json:"id"`

	// ContainerID: unique ID of the container the domain is assigned to.
	ContainerID string `json:"container_id"`

	// Hostname: domain assigned to the container.
	Hostname string `json:"hostname"`

	// Status: domain status.
	// Default value: unknown_status
	Status DomainStatus `json:"status"`

	// ErrorMessage: domain last error message.
	ErrorMessage *string `json:"error_message"`

	// CreatedAt: domain creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: domain last update date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Tags: a list of arbitrary tags associated with the domain.
	Tags []string `json:"tags"`
}

// Namespace: namespace.
type Namespace struct {
	// ID: namespace unique ID.
	ID string `json:"id"`

	// Name: namespace name.
	Name string `json:"name"`

	// OrganizationID: unique ID of the Organization the namespace belongs to.
	OrganizationID string `json:"organization_id"`

	// ProjectID: unique ID of the Project the namespace belongs to.
	ProjectID string `json:"project_id"`

	// Description: namespace description.
	Description string `json:"description"`

	// Status: namespace status.
	// Default value: unknown_status
	Status NamespaceStatus `json:"status"`

	// ErrorMessage: namespace last error message.
	ErrorMessage *string `json:"error_message"`

	// EnvironmentVariables: namespace environment variables.
	EnvironmentVariables map[string]string `json:"environment_variables"`

	// SecretEnvironmentVariables: namespace secret environment variables.
	SecretEnvironmentVariables map[string]string `json:"secret_environment_variables"`

	// Tags: a list of arbitrary tags associated with the namespace.
	Tags []string `json:"tags"`

	// CreatedAt: namespace creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: namespace last update date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Region: region in which the namespace will be created.
	Region scw.Region `json:"region"`
}

// Trigger: trigger.
type Trigger struct {
	// ID: trigger unique ID.
	ID string `json:"id"`

	// Name: name of the trigger.
	Name string `json:"name"`

	// Description: description of the trigger.
	Description string `json:"description"`

	// Tags: tags of the trigger.
	Tags []string `json:"tags"`

	// Status: trigger status.
	// Default value: unknown_status
	Status TriggerStatus `json:"status"`

	// ErrorMessage: trigger last error message.
	ErrorMessage *string `json:"error_message"`

	// ContainerID: ID of the container to trigger.
	ContainerID string `json:"container_id"`

	// DestinationConfig: configuration of the destination to trigger.
	DestinationConfig *TriggerDestinationConfig `json:"destination_config"`

	// SourceType: type of source that will trigger the container.
	// Default value: unknown_source_type
	SourceType TriggerSourceType `json:"source_type"`

	// CronConfig: configuration for a cron source.
	// Precisely one of CronConfig, SqsConfig, NatsConfig must be set.
	CronConfig *TriggerCronConfig `json:"cron_config,omitempty"`

	// SqsConfig: configuration for an SQS queue source.
	// Precisely one of CronConfig, SqsConfig, NatsConfig must be set.
	SqsConfig *TriggerSQSConfig `json:"sqs_config,omitempty"`

	// NatsConfig: configuration for a NATS source.
	// Precisely one of CronConfig, SqsConfig, NatsConfig must be set.
	NatsConfig *TriggerNATSConfig `json:"nats_config,omitempty"`

	// CreatedAt: trigger creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: trigger last update date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// UpdateContainerRequestProbe: update container request probe.
type UpdateContainerRequestProbe struct {
	FailureThreshold *uint32 `json:"failure_threshold"`

	Interval *scw.Duration `json:"interval"`

	Timeout *scw.Duration `json:"timeout"`

	// Precisely one of HTTP, TCP must be set.
	HTTP *UpdateContainerRequestProbeHTTPProbe `json:"http,omitempty"`

	// Precisely one of HTTP, TCP must be set.
	TCP *UpdateContainerRequestProbeTCPProbe `json:"tcp,omitempty"`
}

// UpdateTriggerRequestCronConfig: update trigger request cron config.
type UpdateTriggerRequestCronConfig struct {
	// Schedule: uNIX cron schedule to run job (e.g., "* * * * *").
	Schedule *string `json:"schedule"`

	// Timezone: timezone for the cron schedule, in tz database format (e.g., "Europe/Paris").
	Timezone *string `json:"timezone"`

	// Body: body to send to the container when the trigger is invoked.
	Body *string `json:"body"`

	// Headers: additional headers to send to the container when the trigger is invoked.
	Headers *map[string]string `json:"headers"`
}

// UpdateTriggerRequestDestinationConfig: update trigger request destination config.
type UpdateTriggerRequestDestinationConfig struct {
	// HTTPPath: the HTTP path to send the request to (e.g., "/my-webhook-endpoint").
	HTTPPath *string `json:"http_path"`

	// HTTPMethod: the HTTP method to use when sending the request (e.g., get, post, put, patch, delete). Must be specified as lowercase.
	// Default value: unknown_http_method
	HTTPMethod *UpdateTriggerRequestDestinationConfigHTTPMethod `json:"http_method"`
}

// UpdateTriggerRequestNATSConfig: update trigger request nats config.
type UpdateTriggerRequestNATSConfig struct {
	// ServerURLs: the URLs of the NATS server (e.g., "nats://nats.mnq.fr-par.scaleway.com:4222").
	ServerURLs *[]string `json:"server_urls"`

	// Subject: nATS subject to subscribe to (e.g., "my-subject").
	Subject *string `json:"subject"`

	// CredentialsFileContent: the credentials from this file will be used to authenticate with the NATS server and subscribe to the specified subject.
	CredentialsFileContent *string `json:"credentials_file_content"`
}

// UpdateTriggerRequestSQSConfig: update trigger request sqs config.
type UpdateTriggerRequestSQSConfig struct {
	// Region: the region where the SQS queue is hosted (e.g., "fr-par", "nl-ams").
	Region *scw.Region `json:"region"`

	// Endpoint: endpoint URL to use to access SQS (e.g., "https://sqs.mnq.fr-par.scaleway.com").
	Endpoint *string `json:"endpoint"`

	// AccessKeyID: the access key for accessing the SQS queue.
	AccessKeyID *string `json:"access_key_id"`

	// SecretAccessKey: the secret key for accessing the SQS queue.
	SecretAccessKey *string `json:"secret_access_key"`

	// QueueURL: the URL of the SQS queue to monitor for messages.
	QueueURL *string `json:"queue_url"`
}

// CreateContainerRequest: create container request.
type CreateContainerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NamespaceID: unique ID of the namespace the container belongs to.
	NamespaceID string `json:"namespace_id"`

	// Name: container name.
	Name string `json:"name"`

	// EnvironmentVariables: environment variables of the container.
	EnvironmentVariables map[string]string `json:"environment_variables"`

	// SecretEnvironmentVariables: secret environment variables of the container.
	SecretEnvironmentVariables map[string]string `json:"secret_environment_variables"`

	// MinScale: minimum number of instances to scale the container to.
	MinScale *uint32 `json:"min_scale,omitempty"`

	// MaxScale: maximum number of instances to scale the container to.
	MaxScale *uint32 `json:"max_scale,omitempty"`

	// MemoryLimitBytes: memory limit of the container in bytes.
	MemoryLimitBytes *scw.Size `json:"memory_limit_bytes,omitempty"`

	// MvcpuLimit: CPU limit of the container in mvCPU.
	MvcpuLimit *uint32 `json:"mvcpu_limit,omitempty"`

	// Timeout: processing time limit for the container.
	Timeout *scw.Duration `json:"timeout,omitempty"`

	// Privacy: privacy policy of the container.
	// Default value: unknown_privacy
	Privacy ContainerPrivacy `json:"privacy"`

	// Description: container description.
	Description *string `json:"description,omitempty"`

	// Image: image reference (e.g. "rg.fr-par.scw.cloud/my-registry-namespace/image:tag").
	Image string `json:"image"`

	// Protocol: protocol the container uses.
	// Default value: unknown_protocol
	Protocol ContainerProtocol `json:"protocol"`

	// Port: port the container listens on.
	Port *uint32 `json:"port,omitempty"`

	// HTTPSConnectionsOnly: if true, it will allow only HTTPS connections to access your container to prevent it from being triggered by insecure connections (HTTP).
	HTTPSConnectionsOnly *bool `json:"https_connections_only,omitempty"`

	// Sandbox: execution environment of the container.
	// Default value: unknown_sandbox
	Sandbox ContainerSandbox `json:"sandbox"`

	// LocalStorageLimitBytes: local storage limit of the container (in bytes).
	LocalStorageLimitBytes *scw.Size `json:"local_storage_limit_bytes,omitempty"`

	// ScalingOption: parameter used to decide when to scale up or down.
	ScalingOption *ContainerScalingOption `json:"scaling_option,omitempty"`

	// LivenessProbe: if the liveness probe fails, the container will be restarted.
	// It is performed periodically during the container's lifetime. The liveness probe is not executed until the startup probe (if defined) is successful.
	//
	// Possible check types:
	// - http: Perform HTTP check on the container with the specified path.
	// - tcp: Perform TCP check on the container.
	LivenessProbe *ContainerProbe `json:"liveness_probe,omitempty"`

	// StartupProbe: if the startup probe fails, the container will be restarted.
	// This check is useful for applications that might take a long time to start. It is only performed when the container is starting.
	//
	// Possible check types:
	// - http: Perform HTTP check on the container with the specified path.
	// - tcp: Perform TCP check on the container.
	StartupProbe *ContainerProbe `json:"startup_probe,omitempty"`

	// Tags: tags of the Serverless Container.
	Tags []string `json:"tags"`

	// PrivateNetworkID: when connected to a Private Network, the container can access other Scaleway resources in this Private Network.
	PrivateNetworkID *string `json:"private_network_id,omitempty"`

	// Command: command executed when the container starts. This overrides the default command defined in the container image. This is usually the main executable, or ENTRYPOINT script to run.
	Command []string `json:"command"`

	// Args: arguments passed to the command specified in the "command" field. These override the default arguments from the container image, and behave like command-line parameters.
	Args []string `json:"args"`
}

// CreateDomainRequest: create domain request.
type CreateDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ContainerID: unique ID of the container the domain will be assigned to.
	ContainerID string `json:"container_id"`

	// Hostname: domain assigned to the container.
	Hostname string `json:"hostname"`

	// Tags: a list of arbitrary tags associated with the domain.
	Tags []string `json:"tags"`
}

// CreateNamespaceRequest: create namespace request.
type CreateNamespaceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: unique ID of the Project the namespace belongs to.
	ProjectID string `json:"project_id"`

	// Name: namespace name.
	Name string `json:"name"`

	// Description: namespace description.
	Description *string `json:"description,omitempty"`

	// EnvironmentVariables: namespace environment variables.
	EnvironmentVariables map[string]string `json:"environment_variables"`

	// SecretEnvironmentVariables: namespace secret environment variables.
	SecretEnvironmentVariables map[string]string `json:"secret_environment_variables"`

	// Tags: a list of arbitrary tags associated with the namespace.
	Tags []string `json:"tags"`
}

// CreateTriggerRequest: create trigger request.
type CreateTriggerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ContainerID: ID of the container to trigger.
	ContainerID string `json:"container_id"`

	// Name: name of the trigger.
	Name string `json:"name"`

	// Description: description of the trigger.
	Description *string `json:"description,omitempty"`

	// Tags: tags of the trigger.
	Tags []string `json:"tags"`

	// DestinationConfig: configuration of the destination to trigger.
	DestinationConfig *CreateTriggerRequestDestinationConfig `json:"destination_config,omitempty"`

	// CronConfig: configuration for a cron source.
	// Precisely one of CronConfig, SqsConfig, NatsConfig must be set.
	CronConfig *CreateTriggerRequestCronConfig `json:"cron_config,omitempty"`

	// SqsConfig: configuration for an SQS queue source.
	// Precisely one of CronConfig, SqsConfig, NatsConfig must be set.
	SqsConfig *CreateTriggerRequestSQSConfig `json:"sqs_config,omitempty"`

	// NatsConfig: configuration for a NATS source.
	// Precisely one of CronConfig, SqsConfig, NatsConfig must be set.
	NatsConfig *CreateTriggerRequestNATSConfig `json:"nats_config,omitempty"`
}

// DeleteContainerRequest: delete container request.
type DeleteContainerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ContainerID: UUID of the container to delete.
	ContainerID string `json:"-"`
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

// DeleteTriggerRequest: delete trigger request.
type DeleteTriggerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TriggerID: ID of the trigger to delete.
	TriggerID string `json:"-"`
}

// GetContainerRequest: get container request.
type GetContainerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ContainerID string `json:"-"`
}

// GetDomainRequest: get domain request.
type GetDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DomainID string `json:"-"`
}

// GetNamespaceRequest: get namespace request.
type GetNamespaceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	NamespaceID string `json:"-"`
}

// GetServiceInfoRequest: get service info request.
type GetServiceInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// GetTriggerRequest: get trigger request.
type GetTriggerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	TriggerID string `json:"-"`
}

// ListContainersRequest: list containers request.
type ListContainersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListContainersRequestOrderBy `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`

	NamespaceID *string `json:"-"`

	Name *string `json:"-"`
}

// ListContainersResponse: list containers response.
type ListContainersResponse struct {
	Containers []*Container `json:"containers"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListContainersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListContainersResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListContainersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Containers = append(r.Containers, results.Containers...)
	r.TotalCount += uint64(len(results.Containers))
	return uint64(len(results.Containers)), nil
}

// ListDomainsRequest: list domains request.
type ListDomainsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListDomainsRequestOrderBy `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`

	NamespaceID *string `json:"-"`

	ContainerID *string `json:"-"`
}

// ListDomainsResponse: list domains response.
type ListDomainsResponse struct {
	Domains []*Domain `json:"domains"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListDomainsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Domains = append(r.Domains, results.Domains...)
	r.TotalCount += uint64(len(results.Domains))
	return uint64(len(results.Domains)), nil
}

// ListNamespacesRequest: list namespaces request.
type ListNamespacesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListNamespacesRequestOrderBy `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`

	Name *string `json:"-"`
}

// ListNamespacesResponse: list namespaces response.
type ListNamespacesResponse struct {
	Namespaces []*Namespace `json:"namespaces"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListNamespacesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Namespaces = append(r.Namespaces, results.Namespaces...)
	r.TotalCount += uint64(len(results.Namespaces))
	return uint64(len(results.Namespaces)), nil
}

// ListTriggersRequest: list triggers request.
type ListTriggersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListTriggersRequestOrderBy `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`

	NamespaceID *string `json:"-"`

	ContainerID *string `json:"-"`
}

// ListTriggersResponse: list triggers response.
type ListTriggersResponse struct {
	Triggers []*Trigger `json:"triggers"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTriggersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTriggersResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListTriggersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Triggers = append(r.Triggers, results.Triggers...)
	r.TotalCount += uint64(len(results.Triggers))
	return uint64(len(results.Triggers)), nil
}

// RedeployContainerRequest: redeploy container request.
type RedeployContainerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ContainerID: ID of the container to redeploy.
	ContainerID string `json:"-"`
}

// UpdateContainerRequest: update container request.
type UpdateContainerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ContainerID: UUID of the container to update.
	ContainerID string `json:"-"`

	// EnvironmentVariables: environment variables of the container.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// SecretEnvironmentVariables: secret environment variables of the container.
	SecretEnvironmentVariables *map[string]string `json:"secret_environment_variables,omitempty"`

	// MinScale: minimum number of instances to scale the container to.
	MinScale *uint32 `json:"min_scale,omitempty"`

	// MaxScale: maximum number of instances to scale the container to.
	MaxScale *uint32 `json:"max_scale,omitempty"`

	// MemoryLimitBytes: memory limit of the container in bytes.
	MemoryLimitBytes *scw.Size `json:"memory_limit_bytes,omitempty"`

	// MvcpuLimit: CPU limit of the container in mvCPU.
	MvcpuLimit *uint32 `json:"mvcpu_limit,omitempty"`

	// Timeout: processing time limit for the container.
	Timeout *scw.Duration `json:"timeout,omitempty"`

	// Privacy: privacy policy of the container.
	// Default value: unknown_privacy
	Privacy ContainerPrivacy `json:"privacy"`

	// Description: container description.
	Description *string `json:"description,omitempty"`

	// Image: image reference (e.g. "rg.fr-par.scw.cloud/my-registry-namespace/image:tag").
	Image *string `json:"image,omitempty"`

	// Protocol: protocol the container uses.
	// Default value: unknown_protocol
	Protocol ContainerProtocol `json:"protocol"`

	// Port: port the container listens on.
	Port *uint32 `json:"port,omitempty"`

	// HTTPSConnectionOnly: if true, it will allow only HTTPS connections to access your container to prevent it from being triggered by insecure connections (HTTP).
	HTTPSConnectionOnly *bool `json:"https_connection_only,omitempty"`

	// Sandbox: execution environment of the container.
	// Default value: unknown_sandbox
	Sandbox ContainerSandbox `json:"sandbox"`

	// LocalStorageLimitBytes: local storage limit of the container (in bytes).
	LocalStorageLimitBytes *scw.Size `json:"local_storage_limit_bytes,omitempty"`

	// ScalingOption: parameter used to decide when to scale up or down.
	ScalingOption *ContainerScalingOption `json:"scaling_option,omitempty"`

	// LivenessProbe: if the liveness probe fails, the container will be restarted.
	// It is performed periodically during the container's lifetime. The liveness probe is not executed until the startup probe (if defined) is successful.
	//
	// Possible check types:
	// - http: Perform HTTP check on the container with the specified path.
	// - tcp: Perform TCP check on the container.
	LivenessProbe *ContainerProbe `json:"liveness_probe,omitempty"`

	// StartupProbe: if the startup probe fails, the container will be restarted.
	// This check is useful for applications that might take a long time to start. It is only performed when the container is starting.
	//
	// Possible check types:
	// - http: Perform HTTP check on the container with the specified path.
	// - tcp: Perform TCP check on the container.
	StartupProbe *UpdateContainerRequestProbe `json:"startup_probe,omitempty"`

	// Tags: tags of the Serverless Container.
	Tags *[]string `json:"tags,omitempty"`

	// PrivateNetworkID: when connected to a Private Network, the container can access other Scaleway resources in this Private Network.
	PrivateNetworkID *string `json:"private_network_id,omitempty"`

	// Command: command executed when the container starts. This overrides the default command defined in the container image. This is usually the main executable, or ENTRYPOINT script to run.
	Command *[]string `json:"command,omitempty"`

	// Args: arguments passed to the command specified in the "command" field. These override the default arguments from the container image, and behave like command-line parameters.
	Args *[]string `json:"args,omitempty"`
}

// UpdateDomainRequest: update domain request.
type UpdateDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainID: UUID of the domain to update.
	DomainID string `json:"-"`

	// Tags: a list of arbitrary tags associated with the domain.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateNamespaceRequest: update namespace request.
type UpdateNamespaceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NamespaceID: UUID of the namespace to update.
	NamespaceID string `json:"-"`

	// Description: namespace description.
	Description *string `json:"description,omitempty"`

	// EnvironmentVariables: namespace environment variables.
	EnvironmentVariables *map[string]string `json:"environment_variables,omitempty"`

	// SecretEnvironmentVariables: namespace secret environment variables.
	SecretEnvironmentVariables *map[string]string `json:"secret_environment_variables,omitempty"`

	// Tags: a list of arbitrary tags associated with the namespace.
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

	// Tags: tags of the trigger.
	Tags *[]string `json:"tags,omitempty"`

	// DestinationConfig: configuration of the destination to trigger.
	DestinationConfig *UpdateTriggerRequestDestinationConfig `json:"destination_config,omitempty"`

	// CronConfig: configuration for a cron source.
	// Precisely one of CronConfig, SqsConfig, NatsConfig must be set.
	CronConfig *UpdateTriggerRequestCronConfig `json:"cron_config,omitempty"`

	// SqsConfig: configuration for an SQS queue source.
	// Precisely one of CronConfig, SqsConfig, NatsConfig must be set.
	SqsConfig *UpdateTriggerRequestSQSConfig `json:"sqs_config,omitempty"`

	// NatsConfig: configuration for a NATS source.
	// Precisely one of CronConfig, SqsConfig, NatsConfig must be set.
	NatsConfig *UpdateTriggerRequestNATSConfig `json:"nats_config,omitempty"`
}

// Easily run containers on the cloud with a single command.
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

// GetServiceInfo:
func (s *API) GetServiceInfo(req *GetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateNamespace: Namespace name must be unique inside a project.
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
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
		NamespaceStatusUpdating:  {},
		NamespaceStatusDeleting:  {},
		NamespaceStatusLocking:   {},
		NamespaceStatusCreating:  {},
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

// ListNamespaces: By default, the namespaces listed are ordered by creation date in ascending order. This can be modified via the `order_by` field.
//
// Additional parameters can be set in the query to filter, such as `organization_id`, `project_id`, and `name`.
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Query:  query,
	}

	var resp ListNamespacesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateNamespace: Only fields present in the request are updated; others are left untouched.
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
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

// DeleteNamespace: It also deletes in cascade any resource inside the namespace.
//
// This action **cannot** be undone.
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateContainer: Name must be unique inside the given namespace.
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/containers",
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "",
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
		ContainerStatusUpdating:  {},
		ContainerStatusDeleting:  {},
		ContainerStatusLocking:   {},
		ContainerStatusCreating:  {},
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

// ListContainers: By default, the containers listed are ordered by creation date in ascending order. This can be modified via the `order_by` field.
//
// Additional parameters can be set in the query to filter, such as `organization_id`, `project_id`, and `name`.
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/containers",
		Query:  query,
	}

	var resp ListContainersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateContainer: Only fields present in the request are updated; others are left untouched.
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "",
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

// DeleteContainer: It also deletes in cascade any resource linked to the container (crons, tokens, etc.).
//
// This action **cannot** be undone.
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "",
	}

	var resp Container

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDomain: Create a new custom domain for the container with the specified ID.
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/domains",
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

// GetDomain: Get the custom domain associated with the specified ID.
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
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
		DomainStatusCreating:  {},
		DomainStatusUpdating:  {},
		DomainStatusDeleting:  {},
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

// ListDomains: By default, the custom domains listed are ordered by creation date in ascending order. This can be modified via the `order_by` field.
//
// Additional parameters can be set in the query to filter the output, such as `organization_id`, `project_id`, `namespace_id`, or `container_id`.
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "container_id", req.ContainerID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Query:  query,
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDomain: Only fields present in the request are updated; others are left untouched.
func (s *API) UpdateDomain(req *UpdateDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
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
		Method: "PATCH",
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
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

// DeleteDomain: Delete the custom domain associated with the specified ID.
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RedeployContainer: Performs a rollout of the container by creating new instances with the latest image version and terminating the old instances.
// When using mutable registry image references (e.g. `my-registry-namespace/image:tag`), this endpoint can be used to force the container to use
// the most recent image version available in the registry.
func (s *API) RedeployContainer(req *RedeployContainerRequest, opts ...scw.RequestOption) (*Container, error) {
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/containers/" + fmt.Sprint(req.ContainerID) + "/redeploy",
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

// CreateTrigger: Create a new trigger for the container with the specified ID.
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/triggers",
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

// GetTrigger: Get the trigger associated with the specified ID.
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
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
		TriggerStatusUpdating:  {},
		TriggerStatusCreating:  {},
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

// ListTriggers: By default, the triggers listed are ordered by creation date in ascending order. This can be modified via the `order_by` field.
//
// Additional parameters can be set in the query to filter, such as `organization_id`, `project_id`, `namespace_id`, or `container_id`.
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "container_id", req.ContainerID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/triggers",
		Query:  query,
	}

	var resp ListTriggersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateTrigger: When updating a trigger, you cannot specify a different source type than the one already set.
// Only fields present in the request are updated; others are left untouched.
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
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

// DeleteTrigger: This action **cannot** be undone.
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
		Path:   "/containers/v1/regions/" + fmt.Sprint(req.Region) + "/triggers/" + fmt.Sprint(req.TriggerID) + "",
	}

	var resp Trigger

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

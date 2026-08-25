// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package autoscaling provides methods and message types of the autoscaling v1alpha2 API.
package autoscaling

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
	defaultAutoscalingRetryInterval = 15 * time.Second
	defaultAutoscalingTimeout       = 5 * time.Minute
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

type AlertAlertType string

const (
	// Unknown alert type.
	AlertAlertTypeUnknownAlertType = AlertAlertType("unknown_alert_type")
	// Quotas exceeded during scaling.
	AlertAlertTypeQuotasExceeded = AlertAlertType("quotas_exceeded")
	// Out of stock error during scaling out.
	AlertAlertTypeOutOfStock = AlertAlertType("out_of_stock")
	// Invalid template configuration.
	AlertAlertTypeInvalidTemplate = AlertAlertType("invalid_template")
	// Template not found.
	AlertAlertTypeTemplateNotFound = AlertAlertType("template_not_found")
	// Invalid instance configuration.
	AlertAlertTypeInvalidInstance = AlertAlertType("invalid_instance")
	// Template permissions denied.
	AlertAlertTypeTemplatePermissionsDenied = AlertAlertType("template_permissions_denied")
	// Load balancer not found.
	AlertAlertTypeLoadBalancerNotFound = AlertAlertType("load_balancer_not_found")
	// Load balancer permissions denied.
	AlertAlertTypeLoadBalancerPermissionsDenied = AlertAlertType("load_balancer_permissions_denied")
	// Load balancer backend not found.
	AlertAlertTypeBackendNotFound = AlertAlertType("backend_not_found")
	// Load balancer backend permissions denied.
	AlertAlertTypeBackendPermissionsDenied = AlertAlertType("backend_permissions_denied")
)

func (enum AlertAlertType) String() string {
	if enum == "" {
		// return default value if empty
		return string(AlertAlertTypeUnknownAlertType)
	}
	return string(enum)
}

func (enum AlertAlertType) Values() []AlertAlertType {
	return []AlertAlertType{
		"unknown_alert_type",
		"quotas_exceeded",
		"out_of_stock",
		"invalid_template",
		"template_not_found",
		"invalid_instance",
		"template_permissions_denied",
		"load_balancer_not_found",
		"load_balancer_permissions_denied",
		"backend_not_found",
		"backend_permissions_denied",
	}
}

func (enum AlertAlertType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AlertAlertType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AlertAlertType(AlertAlertType(tmp).String())
	return nil
}

type GroupGroupStatus string

const (
	// Unknown status.
	GroupGroupStatusUnknownGroupStatus = GroupGroupStatus("unknown_group_status")
	// Group is active and operating normally.
	GroupGroupStatusActive = GroupGroupStatus("active")
	// Group is scaling out (adding instances).
	GroupGroupStatusScalingOut = GroupGroupStatus("scaling_out")
	// Group is scaling in (removing instances).
	GroupGroupStatusScalingIn = GroupGroupStatus("scaling_in")
	// Group is refreshing its configuration.
	GroupGroupStatusRefreshing = GroupGroupStatus("refreshing")
	// Group is healing failed instances.
	GroupGroupStatusHealing = GroupGroupStatus("healing")
	// Group encountered a scaling failure.
	GroupGroupStatusScalingFailure = GroupGroupStatus("scaling_failure")
	// Group is being deleted.
	GroupGroupStatusDeleting = GroupGroupStatus("deleting")
)

func (enum GroupGroupStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(GroupGroupStatusUnknownGroupStatus)
	}
	return string(enum)
}

func (enum GroupGroupStatus) Values() []GroupGroupStatus {
	return []GroupGroupStatus{
		"unknown_group_status",
		"active",
		"scaling_out",
		"scaling_in",
		"refreshing",
		"healing",
		"scaling_failure",
		"deleting",
	}
}

func (enum GroupGroupStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *GroupGroupStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = GroupGroupStatus(GroupGroupStatus(tmp).String())
	return nil
}

type GroupLoadBalancerConfigurationBackendAddressFamily string

const (
	GroupLoadBalancerConfigurationBackendAddressFamilyUnknownAddressFamily = GroupLoadBalancerConfigurationBackendAddressFamily("unknown_address_family")
	GroupLoadBalancerConfigurationBackendAddressFamilyIPv4                 = GroupLoadBalancerConfigurationBackendAddressFamily("ipv4")
	GroupLoadBalancerConfigurationBackendAddressFamilyIPv6                 = GroupLoadBalancerConfigurationBackendAddressFamily("ipv6")
)

func (enum GroupLoadBalancerConfigurationBackendAddressFamily) String() string {
	if enum == "" {
		// return default value if empty
		return string(GroupLoadBalancerConfigurationBackendAddressFamilyUnknownAddressFamily)
	}
	return string(enum)
}

func (enum GroupLoadBalancerConfigurationBackendAddressFamily) Values() []GroupLoadBalancerConfigurationBackendAddressFamily {
	return []GroupLoadBalancerConfigurationBackendAddressFamily{
		"unknown_address_family",
		"ipv4",
		"ipv6",
	}
}

func (enum GroupLoadBalancerConfigurationBackendAddressFamily) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *GroupLoadBalancerConfigurationBackendAddressFamily) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = GroupLoadBalancerConfigurationBackendAddressFamily(GroupLoadBalancerConfigurationBackendAddressFamily(tmp).String())
	return nil
}

type GroupSummaryScalingPolicyTargetType string

const (
	GroupSummaryScalingPolicyTargetTypeUnknownScalingPolicyTargetType = GroupSummaryScalingPolicyTargetType("unknown_scaling_policy_target_type")
	GroupSummaryScalingPolicyTargetTypeFixedSize                      = GroupSummaryScalingPolicyTargetType("fixed_size")
	GroupSummaryScalingPolicyTargetTypeCPUTarget                      = GroupSummaryScalingPolicyTargetType("cpu_target")
	GroupSummaryScalingPolicyTargetTypeMemoryTarget                   = GroupSummaryScalingPolicyTargetType("memory_target")
)

func (enum GroupSummaryScalingPolicyTargetType) String() string {
	if enum == "" {
		// return default value if empty
		return string(GroupSummaryScalingPolicyTargetTypeUnknownScalingPolicyTargetType)
	}
	return string(enum)
}

func (enum GroupSummaryScalingPolicyTargetType) Values() []GroupSummaryScalingPolicyTargetType {
	return []GroupSummaryScalingPolicyTargetType{
		"unknown_scaling_policy_target_type",
		"fixed_size",
		"cpu_target",
		"memory_target",
	}
}

func (enum GroupSummaryScalingPolicyTargetType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *GroupSummaryScalingPolicyTargetType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = GroupSummaryScalingPolicyTargetType(GroupSummaryScalingPolicyTargetType(tmp).String())
	return nil
}

type ListGroupsRequestOrderBy string

const (
	ListGroupsRequestOrderByCreatedAtDesc = ListGroupsRequestOrderBy("created_at_desc")
	ListGroupsRequestOrderByCreatedAtAsc  = ListGroupsRequestOrderBy("created_at_asc")
)

func (enum ListGroupsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListGroupsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListGroupsRequestOrderBy) Values() []ListGroupsRequestOrderBy {
	return []ListGroupsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
	}
}

func (enum ListGroupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListGroupsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListGroupsRequestOrderBy(ListGroupsRequestOrderBy(tmp).String())
	return nil
}

type LogLogLevel string

const (
	LogLogLevelUnknownLogLevel = LogLogLevel("unknown_log_level")
	LogLogLevelInfo            = LogLogLevel("info")
	LogLogLevelWarning         = LogLogLevel("warning")
	LogLogLevelError           = LogLogLevel("error")
)

func (enum LogLogLevel) String() string {
	if enum == "" {
		// return default value if empty
		return string(LogLogLevelUnknownLogLevel)
	}
	return string(enum)
}

func (enum LogLogLevel) Values() []LogLogLevel {
	return []LogLogLevel{
		"unknown_log_level",
		"info",
		"warning",
		"error",
	}
}

func (enum LogLogLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LogLogLevel) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LogLogLevel(LogLogLevel(tmp).String())
	return nil
}

// LoadBalancerConfigurationSpecAutoHealing: load balancer configuration spec auto healing.
type LoadBalancerConfigurationSpecAutoHealing struct {
	// Enabled: whether auto-healing is enabled.
	Enabled *bool `json:"enabled"`

	// GracePeriod: grace period for health checks.
	GracePeriod *scw.Duration `json:"grace_period"`
}

// LoadBalancerConfigurationSpecBackend: load balancer configuration spec backend.
type LoadBalancerConfigurationSpecBackend struct {
	// BackendID: ID of the load balancer backend.
	BackendID string `json:"backend_id"`

	// AddressFamily: IP address family (IPv4 or IPv6).
	// Default value: unknown_address_family
	AddressFamily GroupLoadBalancerConfigurationBackendAddressFamily `json:"address_family"`

	// PrivateNetworkID: optional private network ID.
	PrivateNetworkID *string `json:"private_network_id"`
}

// GroupScalingPolicyScalingPolicyCPUTarget: group scaling policy scaling policy cpu target.
type GroupScalingPolicyScalingPolicyCPUTarget struct {
	TargetAvgPercent uint32 `json:"target_avg_percent"`
}

// GroupScalingPolicyScalingPolicyFixedSize: group scaling policy scaling policy fixed size.
type GroupScalingPolicyScalingPolicyFixedSize struct {
	Size uint32 `json:"size"`
}

// GroupScalingPolicyScalingPolicyMemoryTarget: group scaling policy scaling policy memory target.
type GroupScalingPolicyScalingPolicyMemoryTarget struct {
	TargetAvgPercent uint32 `json:"target_avg_percent"`
}

// GroupLoadBalancerConfigurationAutoHealing: group load balancer configuration auto healing.
type GroupLoadBalancerConfigurationAutoHealing struct {
	Enabled bool `json:"enabled"`

	GracePeriod *scw.Duration `json:"grace_period"`
}

// GroupLoadBalancerConfigurationBackend: group load balancer configuration backend.
type GroupLoadBalancerConfigurationBackend struct {
	BackendID string `json:"backend_id"`

	// AddressFamily: default value: unknown_address_family
	AddressFamily GroupLoadBalancerConfigurationBackendAddressFamily `json:"address_family"`

	PrivateNetworkID *string `json:"private_network_id"`
}

// Alert: alert.
type Alert struct {
	// Type: type of alert.
	// Default value: unknown_alert_type
	Type AlertAlertType `json:"type"`

	// OpenedAt: timestamp when the alert was opened.
	OpenedAt *time.Time `json:"opened_at"`

	// ClosedAt: optional timestamp when the alert was closed.
	ClosedAt *time.Time `json:"closed_at"`

	// GroupID: ID of the group.
	GroupID string `json:"group_id"`

	// FailingQuotas: list of quota names that are exceeded (for quotas_exceeded alerts).
	FailingQuotas []string `json:"failing_quotas"`
}

// LoadBalancerConfigurationSpec: load balancer configuration spec.
type LoadBalancerConfigurationSpec struct {
	// LoadBalancerID: ID of the load balancer (set to empty to disable).
	LoadBalancerID *string `json:"load_balancer_id"`

	// Backends: list of load balancer backend configurations.
	Backends []*LoadBalancerConfigurationSpecBackend `json:"backends"`

	// AutoHealing: auto-healing configuration.
	AutoHealing *LoadBalancerConfigurationSpecAutoHealing `json:"auto_healing"`
}

// ScalingPolicySpec: scaling policy spec.
type ScalingPolicySpec struct {
	// MinimumSize: minimum number of instances in the group.
	MinimumSize *uint32 `json:"minimum_size"`

	// MaximumSize: maximum number of instances in the group.
	MaximumSize *uint32 `json:"maximum_size"`

	// ScaleOutCooldown: cooldown period after a scale out event.
	ScaleOutCooldown *scw.Duration `json:"scale_out_cooldown"`

	// ScaleInCooldown: cooldown period after a scale in event.
	ScaleInCooldown *scw.Duration `json:"scale_in_cooldown"`

	// ScaleInStep: number of instances to remove in a single scale in event.
	ScaleInStep *uint32 `json:"scale_in_step"`

	// ScaleOutStep: number of instances to add in a single scale out event.
	ScaleOutStep *uint32 `json:"scale_out_step"`

	// FixedSize: fixed size scaling policy configuration.
	// Precisely one of FixedSize, CPUTarget, MemoryTarget must be set.
	FixedSize *GroupScalingPolicyScalingPolicyFixedSize `json:"fixed_size,omitempty"`

	// CPUTarget: CPU target scaling policy configuration.
	// Precisely one of FixedSize, CPUTarget, MemoryTarget must be set.
	CPUTarget *GroupScalingPolicyScalingPolicyCPUTarget `json:"cpu_target,omitempty"`

	// MemoryTarget: memory target scaling policy configuration.
	// Precisely one of FixedSize, CPUTarget, MemoryTarget must be set.
	MemoryTarget *GroupScalingPolicyScalingPolicyMemoryTarget `json:"memory_target,omitempty"`
}

// GroupLoadBalancerConfiguration: group load balancer configuration.
type GroupLoadBalancerConfiguration struct {
	LoadBalancerID string `json:"load_balancer_id"`

	Backends []*GroupLoadBalancerConfigurationBackend `json:"backends"`

	AutoHealing *GroupLoadBalancerConfigurationAutoHealing `json:"auto_healing"`
}

// GroupScalingPolicy: group scaling policy.
type GroupScalingPolicy struct {
	MinimumSize uint32 `json:"minimum_size"`

	MaximumSize uint32 `json:"maximum_size"`

	ScaleOutCooldown *scw.Duration `json:"scale_out_cooldown"`

	ScaleInCooldown *scw.Duration `json:"scale_in_cooldown"`

	ScaleInStep uint32 `json:"scale_in_step"`

	ScaleOutStep uint32 `json:"scale_out_step"`

	// Precisely one of FixedSize, CPUTarget, MemoryTarget must be set.
	FixedSize *GroupScalingPolicyScalingPolicyFixedSize `json:"fixed_size,omitempty"`

	// Precisely one of FixedSize, CPUTarget, MemoryTarget must be set.
	CPUTarget *GroupScalingPolicyScalingPolicyCPUTarget `json:"cpu_target,omitempty"`

	// Precisely one of FixedSize, CPUTarget, MemoryTarget must be set.
	MemoryTarget *GroupScalingPolicyScalingPolicyMemoryTarget `json:"memory_target,omitempty"`
}

// GroupSummary: group summary.
type GroupSummary struct {
	// ProjectID: project ID owning this group.
	ProjectID string `json:"project_id"`

	// ID: unique identifier of the autoscaling group.
	ID string `json:"id"`

	// Name: name of the autoscaling group.
	Name string `json:"name"`

	// Tags: tags associated with the group.
	Tags []string `json:"tags"`

	// Status: current status of the autoscaling group.
	// Default value: unknown_group_status
	Status GroupGroupStatus `json:"status"`

	// CreatedAt: creation timestamp of the group.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update timestamp of the group.
	UpdatedAt *time.Time `json:"updated_at"`

	// TemplateID: the Instance template ID used to create instances in this group.
	TemplateID string `json:"template_id"`

	// LoadBalancerID: optional load balancer ID.
	LoadBalancerID *string `json:"load_balancer_id"`

	// CurrentSize: current number of instances in the group.
	CurrentSize uint32 `json:"current_size"`

	// LatestOpenAlert: most recent active alert if any.
	LatestOpenAlert *Alert `json:"latest_open_alert"`

	// MinimumSize: the minimum number of instances in the group.
	MinimumSize uint32 `json:"minimum_size"`

	// MaximumSize: the maximum number of instances in the group.
	MaximumSize uint32 `json:"maximum_size"`

	// ScalingPolicyTargetType: defines which metric the group uses for scaling (cpu, memory, or fixed size).
	// Default value: unknown_scaling_policy_target_type
	ScalingPolicyTargetType GroupSummaryScalingPolicyTargetType `json:"scaling_policy_target_type"`

	// Zone: availability zone of the group.
	Zone scw.Zone `json:"zone"`
}

// Log: log.
type Log struct {
	// Timestamp: timestamp of the log entry.
	Timestamp *time.Time `json:"timestamp"`

	// Level: log level (info, warning, error).
	// Default value: unknown_log_level
	Level LogLogLevel `json:"level"`

	// Message: log message content.
	Message string `json:"message"`
}

// Server: server.
type Server struct {
	// ServerID: ID of the instance server.
	ServerID string `json:"server_id"`
}

// CreateGroupRequest: create group request.
type CreateGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ProjectID: project ID to create the group in.
	ProjectID string `json:"project_id"`

	// Name: name of the autoscaling group.
	Name string `json:"name"`

	// Tags: tags associated with the group.
	Tags []string `json:"tags"`

	// TemplateID: template ID for instances in this group.
	TemplateID string `json:"template_id"`

	// ScalingPolicySpec: scaling policy configuration.
	ScalingPolicySpec *ScalingPolicySpec `json:"scaling_policy_spec,omitempty"`

	// LoadBalancerConfigurationSpec: optional load balancer configuration.
	LoadBalancerConfigurationSpec *LoadBalancerConfigurationSpec `json:"load_balancer_configuration_spec,omitempty"`
}

// DeleteGroupRequest: delete group request.
type DeleteGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GroupID: ID of the group to delete.
	GroupID string `json:"-"`
}

// GetGroupRequest: get group request.
type GetGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GroupID: ID of the group to get.
	GroupID string `json:"-"`
}

// Group: group.
type Group struct {
	// ID: unique identifier of the autoscaling group.
	ID string `json:"id"`

	// ProjectID: project ID owning this group.
	ProjectID string `json:"project_id"`

	// Name: name of the autoscaling group.
	Name string `json:"name"`

	// Tags: tags associated with the group.
	Tags []string `json:"tags"`

	// CreatedAt: creation timestamp of the group.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update timestamp of the group.
	UpdatedAt *time.Time `json:"updated_at"`

	// Status: current status of the autoscaling group.
	// Default value: unknown_group_status
	Status GroupGroupStatus `json:"status"`

	// OpenAlerts: list of active alerts on the group.
	OpenAlerts []*Alert `json:"open_alerts"`

	// CurrentSize: current number of instances in the group.
	CurrentSize uint32 `json:"current_size"`

	// TargetSize: target number of instances the group is scaling to.
	TargetSize uint32 `json:"target_size"`

	// LastScaleOutAt: timestamp of the last scale out event.
	LastScaleOutAt *time.Time `json:"last_scale_out_at"`

	// LastScaleInAt: timestamp of the last scale in event.
	LastScaleInAt *time.Time `json:"last_scale_in_at"`

	// TemplateID: the Instance template ID used to create instances in this group.
	TemplateID string `json:"template_id"`

	// ScalingPolicy: scaling policy configuration.
	ScalingPolicy *GroupScalingPolicy `json:"scaling_policy"`

	// LoadBalancerConfiguration: optional load balancer configuration.
	LoadBalancerConfiguration *GroupLoadBalancerConfiguration `json:"load_balancer_configuration"`
}

// ListAlertsRequest: list alerts request.
type ListAlertsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Precisely one of GroupID, ProjectID must be set.
	GroupID *string `json:"group_id,omitempty"`

	PageToken *string `json:"-"`

	PageSize *uint32 `json:"-"`

	// Precisely one of GroupID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`
}

// ListAlertsResponse: list alerts response.
type ListAlertsResponse struct {
	// Alerts: list of alerts.
	Alerts []*Alert `json:"alerts"`

	// NextPageToken: token for the next page of results.
	NextPageToken *string `json:"next_page_token"`

	// TotalCount: total number of alerts.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListAlertsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListAlertsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListAlertsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Alerts = append(r.Alerts, results.Alerts...)
	r.TotalCount += uint64(len(results.Alerts))
	return uint64(len(results.Alerts)), nil
}

// ListGroupsRequest: list groups request.
type ListGroupsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: order criteria for listing groups.
	// Default value: created_at_desc
	OrderBy ListGroupsRequestOrderBy `json:"-"`

	// PageSize: page size for pagination.
	PageSize *uint32 `json:"-"`

	// PageToken: token for pagination.
	PageToken *string `json:"-"`

	// ProjectID: project ID to filter groups.
	ProjectID string `json:"-"`

	// TemplateID: template ID to filter groups.
	TemplateID *string `json:"-"`

	// LoadBalancerID: load balancer ID to filter groups.
	LoadBalancerID *string `json:"-"`
}

// ListGroupsResponse: list groups response.
type ListGroupsResponse struct {
	// GroupSummaries: list of group summaries.
	GroupSummaries []*GroupSummary `json:"group_summaries"`

	// NextPageToken: token for the next page of results.
	NextPageToken *string `json:"next_page_token"`

	// TotalCount: total number of groups.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListGroupsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListGroupsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.GroupSummaries = append(r.GroupSummaries, results.GroupSummaries...)
	r.TotalCount += uint64(len(results.GroupSummaries))
	return uint64(len(results.GroupSummaries)), nil
}

// ListLogsRequest: list logs request.
type ListLogsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	GroupID string `json:"-"`

	PageToken *string `json:"-"`

	PageSize *uint32 `json:"-"`

	StartTime *time.Time `json:"-"`

	EndTime *time.Time `json:"-"`
}

// ListLogsResponse: list logs response.
type ListLogsResponse struct {
	// Logs: list of log entries.
	Logs []*Log `json:"logs"`

	// NextPageToken: token for the next page of results.
	NextPageToken *string `json:"next_page_token"`

	// TotalCount: total number of logs.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListLogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Logs = append(r.Logs, results.Logs...)
	r.TotalCount += uint64(len(results.Logs))
	return uint64(len(results.Logs)), nil
}

// ListServersRequest: list servers request.
type ListServersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	GroupID string `json:"-"`

	PageToken *string `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListServersResponse: list servers response.
type ListServersResponse struct {
	// Servers: list of servers.
	Servers []*Server `json:"servers"`

	// NextPageToken: token for the next page of results.
	NextPageToken *string `json:"next_page_token"`

	// TotalCount: total number of servers.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint64(len(results.Servers))
	return uint64(len(results.Servers)), nil
}

// UpdateGroupRequest: update group request.
type UpdateGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GroupID: ID of the group to update.
	GroupID string `json:"-"`

	// Name: new name for the group.
	Name *string `json:"name,omitempty"`

	// Tags: new tags for the group.
	Tags *[]string `json:"tags,omitempty"`

	// TemplateID: new template ID for instances.
	TemplateID *string `json:"template_id,omitempty"`

	// ScalingPolicySpec: new scaling policy configuration.
	ScalingPolicySpec *ScalingPolicySpec `json:"scaling_policy_spec,omitempty"`

	// LoadBalancerConfigurationSpec: new load balancer configuration.
	LoadBalancerConfigurationSpec *LoadBalancerConfigurationSpec `json:"load_balancer_configuration_spec,omitempty"`
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

func (s *API) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3}
}

// ListGroups: List all autoscaling groups in a project.
func (s *API) ListGroups(req *ListGroupsRequest, opts ...scw.RequestOption) (*ListGroupsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

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
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "template_id", req.TemplateID)
	parameter.AddToQuery(query, "load_balancer_id", req.LoadBalancerID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/autoscaling/v1alpha2/zones/" + fmt.Sprint(req.Zone) + "/groups",
		Query:  query,
	}

	var resp ListGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGroup: Get details of a specified autoscaling group including its
// configuration, current size, and status.
func (s *API) GetGroup(req *GetGroupRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/autoscaling/v1alpha2/zones/" + fmt.Sprint(req.Zone) + "/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForGroupRequest is used by WaitForGroup method.
type WaitForGroupRequest struct {
	Zone          scw.Zone
	GroupID       string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForGroup waits for the Group to reach a terminal state.
func (s *API) WaitForGroup(req *WaitForGroupRequest, opts ...scw.RequestOption) (*Group, error) {
	timeout := defaultAutoscalingTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultAutoscalingRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[GroupGroupStatus]struct{}{
		GroupGroupStatusScalingOut: {},
		GroupGroupStatusScalingIn:  {},
		GroupGroupStatusRefreshing: {},
		GroupGroupStatusHealing:    {},
		GroupGroupStatusDeleting:   {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetGroup(&GetGroupRequest{
				Zone:    req.Zone,
				GroupID: req.GroupID,
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
		return nil, errors.Wrap(err, "waiting for Group failed")
	}

	return res.(*Group), nil
}

// CreateGroup: Create a new autoscaling group with the specified configuration
// including template, scaling policy, and optional load balancer
// settings.
func (s *API) CreateGroup(req *CreateGroupRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/autoscaling/v1alpha2/zones/" + fmt.Sprint(req.Zone) + "/groups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateGroup: Update the configuration of a specified autoscaling group including
// name, tags, template, scaling policy, and load balancer settings.
func (s *API) UpdateGroup(req *UpdateGroupRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/autoscaling/v1alpha2/zones/" + fmt.Sprint(req.Zone) + "/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteGroup: Delete a specified autoscaling group and all its associated
// resources.
func (s *API) DeleteGroup(req *DeleteGroupRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/autoscaling/v1alpha2/zones/" + fmt.Sprint(req.Zone) + "/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListLogs: List logs for a specified autoscaling group to view scaling events
// and activities.
func (s *API) ListLogs(req *ListLogsRequest, opts ...scw.RequestOption) (*ListLogsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "group_id", req.GroupID)
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "start_time", req.StartTime)
	parameter.AddToQuery(query, "end_time", req.EndTime)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/autoscaling/v1alpha2/zones/" + fmt.Sprint(req.Zone) + "/logs",
		Query:  query,
	}

	var resp ListLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServers: List all Instances belonging to a specified autoscaling group.
func (s *API) ListServers(req *ListServersRequest, opts ...scw.RequestOption) (*ListServersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "group_id", req.GroupID)
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/autoscaling/v1alpha2/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:  query,
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListAlerts: List active and historical alerts for a specified autoscaling group.
func (s *API) ListAlerts(req *ListAlertsRequest, opts ...scw.RequestOption) (*ListAlertsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.GroupID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "group_id", req.GroupID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/autoscaling/v1alpha2/zones/" + fmt.Sprint(req.Zone) + "/alerts",
		Query:  query,
	}

	var resp ListAlertsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

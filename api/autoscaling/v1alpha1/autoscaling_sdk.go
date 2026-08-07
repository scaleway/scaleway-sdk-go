// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package autoscaling provides methods and message types of the autoscaling v1alpha1 API.
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

type InstanceGroupEventLevel string

const (
	// Informational log about the Instance group.
	InstanceGroupEventLevelInfo = InstanceGroupEventLevel("info")
	// Success log for Instance scaling or Load Balancer configuration.
	InstanceGroupEventLevelSuccess = InstanceGroupEventLevel("success")
	// Error log for Instance group configuration or Instance group operation.
	InstanceGroupEventLevelError = InstanceGroupEventLevel("error")
)

func (enum InstanceGroupEventLevel) String() string {
	if enum == "" {
		// return default value if empty
		return string(InstanceGroupEventLevelInfo)
	}
	return string(enum)
}

func (enum InstanceGroupEventLevel) Values() []InstanceGroupEventLevel {
	return []InstanceGroupEventLevel{
		"info",
		"success",
		"error",
	}
}

func (enum InstanceGroupEventLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceGroupEventLevel) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceGroupEventLevel(InstanceGroupEventLevel(tmp).String())
	return nil
}

type InstanceGroupEventSource string

const (
	// Unknown source event.
	InstanceGroupEventSourceUnknownSource = InstanceGroupEventSource("unknown_source")
	// Responsible for collecting and analyzing metrics.
	InstanceGroupEventSourceWatcher = InstanceGroupEventSource("watcher")
	// Responsible for processing the policy and generating Instance scale up/down tasks.
	InstanceGroupEventSourceScaler = InstanceGroupEventSource("scaler")
	// Responsible for managing the creation and deletion of Instances, and configuring the Load Balancer attached to the Instance group.
	InstanceGroupEventSourceInstanceManager = InstanceGroupEventSource("instance_manager")
	// Global Instance groups supervisor.
	InstanceGroupEventSourceSupervisor = InstanceGroupEventSource("supervisor")
)

func (enum InstanceGroupEventSource) String() string {
	if enum == "" {
		// return default value if empty
		return string(InstanceGroupEventSourceUnknownSource)
	}
	return string(enum)
}

func (enum InstanceGroupEventSource) Values() []InstanceGroupEventSource {
	return []InstanceGroupEventSource{
		"unknown_source",
		"watcher",
		"scaler",
		"instance_manager",
		"supervisor",
	}
}

func (enum InstanceGroupEventSource) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceGroupEventSource) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceGroupEventSource(InstanceGroupEventSource(tmp).String())
	return nil
}

type InstancePolicyAction string

const (
	// Action of unknown type.
	InstancePolicyActionUnknownAction = InstancePolicyAction("unknown_action")
	// Create one or many new Instances (based on policy type and value).
	InstancePolicyActionScaleUp = InstancePolicyAction("scale_up")
	// Stop and remove one or many Instances (based on policy type and value).
	InstancePolicyActionScaleDown = InstancePolicyAction("scale_down")
)

func (enum InstancePolicyAction) String() string {
	if enum == "" {
		// return default value if empty
		return string(InstancePolicyActionUnknownAction)
	}
	return string(enum)
}

func (enum InstancePolicyAction) Values() []InstancePolicyAction {
	return []InstancePolicyAction{
		"unknown_action",
		"scale_up",
		"scale_down",
	}
}

func (enum InstancePolicyAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstancePolicyAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstancePolicyAction(InstancePolicyAction(tmp).String())
	return nil
}

type InstancePolicyType string

const (
	// Action of unknown type.
	InstancePolicyTypeUnknownType = InstancePolicyType("unknown_type")
	// Add or delete a fixed number of Instances. This must be a positive integer.
	InstancePolicyTypeFlatCount = InstancePolicyType("flat_count")
	// Add or delete a percentage of the total Instance group size.
	InstancePolicyTypePercentOfTotalGroup = InstancePolicyType("percent_of_total_group")
	// Set the total number of Instances in the group to this arbitrary number. This must be a positive integer.
	InstancePolicyTypeSetTotalGroup = InstancePolicyType("set_total_group")
)

func (enum InstancePolicyType) String() string {
	if enum == "" {
		// return default value if empty
		return string(InstancePolicyTypeUnknownType)
	}
	return string(enum)
}

func (enum InstancePolicyType) Values() []InstancePolicyType {
	return []InstancePolicyType{
		"unknown_type",
		"flat_count",
		"percent_of_total_group",
		"set_total_group",
	}
}

func (enum InstancePolicyType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstancePolicyType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstancePolicyType(InstancePolicyType(tmp).String())
	return nil
}

type InstanceTemplateStatus string

const (
	// Unknown status.
	InstanceTemplateStatusUnknownStatus = InstanceTemplateStatus("unknown_status")
	// Template is ready and healthy.
	InstanceTemplateStatusReady = InstanceTemplateStatus("ready")
	// Template is in error state, see events log.
	InstanceTemplateStatusError = InstanceTemplateStatus("error")
)

func (enum InstanceTemplateStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(InstanceTemplateStatusUnknownStatus)
	}
	return string(enum)
}

func (enum InstanceTemplateStatus) Values() []InstanceTemplateStatus {
	return []InstanceTemplateStatus{
		"unknown_status",
		"ready",
		"error",
	}
}

func (enum InstanceTemplateStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceTemplateStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceTemplateStatus(InstanceTemplateStatus(tmp).String())
	return nil
}

type ListInstanceGroupEventsRequestOrderBy string

const (
	// Order by creation date (descending chronological order).
	ListInstanceGroupEventsRequestOrderByCreatedAtDesc = ListInstanceGroupEventsRequestOrderBy("created_at_desc")
	// Order by creation date (ascending chronological order).
	ListInstanceGroupEventsRequestOrderByCreatedAtAsc = ListInstanceGroupEventsRequestOrderBy("created_at_asc")
)

func (enum ListInstanceGroupEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListInstanceGroupEventsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListInstanceGroupEventsRequestOrderBy) Values() []ListInstanceGroupEventsRequestOrderBy {
	return []ListInstanceGroupEventsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
	}
}

func (enum ListInstanceGroupEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInstanceGroupEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInstanceGroupEventsRequestOrderBy(ListInstanceGroupEventsRequestOrderBy(tmp).String())
	return nil
}

type ListInstanceGroupsRequestOrderBy string

const (
	// Order by creation date (descending chronological order).
	ListInstanceGroupsRequestOrderByCreatedAtDesc = ListInstanceGroupsRequestOrderBy("created_at_desc")
	// Order by creation date (ascending chronological order).
	ListInstanceGroupsRequestOrderByCreatedAtAsc = ListInstanceGroupsRequestOrderBy("created_at_asc")
)

func (enum ListInstanceGroupsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListInstanceGroupsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListInstanceGroupsRequestOrderBy) Values() []ListInstanceGroupsRequestOrderBy {
	return []ListInstanceGroupsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
	}
}

func (enum ListInstanceGroupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInstanceGroupsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInstanceGroupsRequestOrderBy(ListInstanceGroupsRequestOrderBy(tmp).String())
	return nil
}

type ListInstancePoliciesRequestOrderBy string

const (
	// Order by creation date (descending chronological order).
	ListInstancePoliciesRequestOrderByCreatedAtDesc = ListInstancePoliciesRequestOrderBy("created_at_desc")
	// Order by creation date (ascending chronological order).
	ListInstancePoliciesRequestOrderByCreatedAtAsc = ListInstancePoliciesRequestOrderBy("created_at_asc")
)

func (enum ListInstancePoliciesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListInstancePoliciesRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListInstancePoliciesRequestOrderBy) Values() []ListInstancePoliciesRequestOrderBy {
	return []ListInstancePoliciesRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
	}
}

func (enum ListInstancePoliciesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInstancePoliciesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInstancePoliciesRequestOrderBy(ListInstancePoliciesRequestOrderBy(tmp).String())
	return nil
}

type ListInstanceTemplatesRequestOrderBy string

const (
	// Order by creation date (descending chronological order).
	ListInstanceTemplatesRequestOrderByCreatedAtDesc = ListInstanceTemplatesRequestOrderBy("created_at_desc")
	// Order by creation date (ascending chronological order).
	ListInstanceTemplatesRequestOrderByCreatedAtAsc = ListInstanceTemplatesRequestOrderBy("created_at_asc")
)

func (enum ListInstanceTemplatesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListInstanceTemplatesRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListInstanceTemplatesRequestOrderBy) Values() []ListInstanceTemplatesRequestOrderBy {
	return []ListInstanceTemplatesRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
	}
}

func (enum ListInstanceTemplatesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInstanceTemplatesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInstanceTemplatesRequestOrderBy(ListInstanceTemplatesRequestOrderBy(tmp).String())
	return nil
}

type MetricAggregate string

const (
	MetricAggregateAggregateUnknown = MetricAggregate("aggregate_unknown")
	// The average value of all sample points during the specified interval.
	MetricAggregateAggregateAverage = MetricAggregate("aggregate_average")
	// The maximum value of all sample points during the specified interval.
	MetricAggregateAggregateMax = MetricAggregate("aggregate_max")
	// The minimum value of all sample points during the specified interval.
	MetricAggregateAggregateMin = MetricAggregate("aggregate_min")
	// The sum of all sample values during the specified interval.
	MetricAggregateAggregateSum = MetricAggregate("aggregate_sum")
)

func (enum MetricAggregate) String() string {
	if enum == "" {
		// return default value if empty
		return string(MetricAggregateAggregateUnknown)
	}
	return string(enum)
}

func (enum MetricAggregate) Values() []MetricAggregate {
	return []MetricAggregate{
		"aggregate_unknown",
		"aggregate_average",
		"aggregate_max",
		"aggregate_min",
		"aggregate_sum",
	}
}

func (enum MetricAggregate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MetricAggregate) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MetricAggregate(MetricAggregate(tmp).String())
	return nil
}

type MetricManagedMetric string

const (
	MetricManagedMetricManagedMetricUnknown = MetricManagedMetric("managed_metric_unknown")
	// Percentage of CPU usage across the whole Instance group.
	MetricManagedMetricManagedMetricInstanceCPU = MetricManagedMetric("managed_metric_instance_cpu")
	// Rate input, in bytes/sec, of the Instance's public interface.
	MetricManagedMetricManagedMetricInstanceNetworkIn = MetricManagedMetric("managed_metric_instance_network_in")
	// Rate output, in bytes/sec, of the Instance's public interface.
	MetricManagedMetricManagedMetricInstanceNetworkOut = MetricManagedMetric("managed_metric_instance_network_out")
	// Cumulative number of connections established to the Load Balancer backend related to the Instance group.
	MetricManagedMetricManagedLoadbalancerBackendConnectionsRate = MetricManagedMetric("managed_loadbalancer_backend_connections_rate")
	// Rate, in bytes/sec, of all traffic forwarded to all Load Balancer backend servers.
	MetricManagedMetricManagedLoadbalancerBackendThroughput = MetricManagedMetric("managed_loadbalancer_backend_throughput")
)

func (enum MetricManagedMetric) String() string {
	if enum == "" {
		// return default value if empty
		return string(MetricManagedMetricManagedMetricUnknown)
	}
	return string(enum)
}

func (enum MetricManagedMetric) Values() []MetricManagedMetric {
	return []MetricManagedMetric{
		"managed_metric_unknown",
		"managed_metric_instance_cpu",
		"managed_metric_instance_network_in",
		"managed_metric_instance_network_out",
		"managed_loadbalancer_backend_connections_rate",
		"managed_loadbalancer_backend_throughput",
	}
}

func (enum MetricManagedMetric) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MetricManagedMetric) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MetricManagedMetric(MetricManagedMetric(tmp).String())
	return nil
}

type MetricOperator string

const (
	MetricOperatorOperatorUnknown = MetricOperator("operator_unknown")
	// Equivalent of a greater than symbol (>).
	MetricOperatorOperatorGreaterThan = MetricOperator("operator_greater_than")
	// Equivalent of a less than symbol (<).
	MetricOperatorOperatorLessThan = MetricOperator("operator_less_than")
)

func (enum MetricOperator) String() string {
	if enum == "" {
		// return default value if empty
		return string(MetricOperatorOperatorUnknown)
	}
	return string(enum)
}

func (enum MetricOperator) Values() []MetricOperator {
	return []MetricOperator{
		"operator_unknown",
		"operator_greater_than",
		"operator_less_than",
	}
}

func (enum MetricOperator) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MetricOperator) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MetricOperator(MetricOperator(tmp).String())
	return nil
}

type UpdateInstancePolicyRequestMetricAggregate string

const (
	UpdateInstancePolicyRequestMetricAggregateAggregateUnknown = UpdateInstancePolicyRequestMetricAggregate("aggregate_unknown")
	// The average value of all sample points during the specified interval.
	UpdateInstancePolicyRequestMetricAggregateAggregateAverage = UpdateInstancePolicyRequestMetricAggregate("aggregate_average")
	// The maximum value of all sample points during the specified interval.
	UpdateInstancePolicyRequestMetricAggregateAggregateMax = UpdateInstancePolicyRequestMetricAggregate("aggregate_max")
	// The minimum value of all sample points during the specified interval.
	UpdateInstancePolicyRequestMetricAggregateAggregateMin = UpdateInstancePolicyRequestMetricAggregate("aggregate_min")
	// The sum of all sample values during the specified interval.
	UpdateInstancePolicyRequestMetricAggregateAggregateSum = UpdateInstancePolicyRequestMetricAggregate("aggregate_sum")
)

func (enum UpdateInstancePolicyRequestMetricAggregate) String() string {
	if enum == "" {
		// return default value if empty
		return string(UpdateInstancePolicyRequestMetricAggregateAggregateUnknown)
	}
	return string(enum)
}

func (enum UpdateInstancePolicyRequestMetricAggregate) Values() []UpdateInstancePolicyRequestMetricAggregate {
	return []UpdateInstancePolicyRequestMetricAggregate{
		"aggregate_unknown",
		"aggregate_average",
		"aggregate_max",
		"aggregate_min",
		"aggregate_sum",
	}
}

func (enum UpdateInstancePolicyRequestMetricAggregate) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UpdateInstancePolicyRequestMetricAggregate) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UpdateInstancePolicyRequestMetricAggregate(UpdateInstancePolicyRequestMetricAggregate(tmp).String())
	return nil
}

type UpdateInstancePolicyRequestMetricManagedMetric string

const (
	UpdateInstancePolicyRequestMetricManagedMetricManagedMetricUnknown = UpdateInstancePolicyRequestMetricManagedMetric("managed_metric_unknown")
	// Percentage of CPU usage across the whole Instance group.
	UpdateInstancePolicyRequestMetricManagedMetricManagedMetricInstanceCPU = UpdateInstancePolicyRequestMetricManagedMetric("managed_metric_instance_cpu")
	// Rate input, in bytes/sec, of the Instance's public interface.
	UpdateInstancePolicyRequestMetricManagedMetricManagedMetricInstanceNetworkIn = UpdateInstancePolicyRequestMetricManagedMetric("managed_metric_instance_network_in")
	// Rate output, in bytes/sec, of the Instance's public interface.
	UpdateInstancePolicyRequestMetricManagedMetricManagedMetricInstanceNetworkOut = UpdateInstancePolicyRequestMetricManagedMetric("managed_metric_instance_network_out")
	// Cumulative number of connections established to the Load Balancer backend related to the Instance group.
	UpdateInstancePolicyRequestMetricManagedMetricManagedLoadbalancerBackendConnectionsRate = UpdateInstancePolicyRequestMetricManagedMetric("managed_loadbalancer_backend_connections_rate")
	// Rate, in bytes/sec, of all traffic forwarded to all Load Balancer backend servers.
	UpdateInstancePolicyRequestMetricManagedMetricManagedLoadbalancerBackendThroughput = UpdateInstancePolicyRequestMetricManagedMetric("managed_loadbalancer_backend_throughput")
)

func (enum UpdateInstancePolicyRequestMetricManagedMetric) String() string {
	if enum == "" {
		// return default value if empty
		return string(UpdateInstancePolicyRequestMetricManagedMetricManagedMetricUnknown)
	}
	return string(enum)
}

func (enum UpdateInstancePolicyRequestMetricManagedMetric) Values() []UpdateInstancePolicyRequestMetricManagedMetric {
	return []UpdateInstancePolicyRequestMetricManagedMetric{
		"managed_metric_unknown",
		"managed_metric_instance_cpu",
		"managed_metric_instance_network_in",
		"managed_metric_instance_network_out",
		"managed_loadbalancer_backend_connections_rate",
		"managed_loadbalancer_backend_throughput",
	}
}

func (enum UpdateInstancePolicyRequestMetricManagedMetric) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UpdateInstancePolicyRequestMetricManagedMetric) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UpdateInstancePolicyRequestMetricManagedMetric(UpdateInstancePolicyRequestMetricManagedMetric(tmp).String())
	return nil
}

type UpdateInstancePolicyRequestMetricOperator string

const (
	UpdateInstancePolicyRequestMetricOperatorOperatorUnknown = UpdateInstancePolicyRequestMetricOperator("operator_unknown")
	// Equivalent of a greater than symbol (>).
	UpdateInstancePolicyRequestMetricOperatorOperatorGreaterThan = UpdateInstancePolicyRequestMetricOperator("operator_greater_than")
	// Equivalent of a less than symbol (<).
	UpdateInstancePolicyRequestMetricOperatorOperatorLessThan = UpdateInstancePolicyRequestMetricOperator("operator_less_than")
)

func (enum UpdateInstancePolicyRequestMetricOperator) String() string {
	if enum == "" {
		// return default value if empty
		return string(UpdateInstancePolicyRequestMetricOperatorOperatorUnknown)
	}
	return string(enum)
}

func (enum UpdateInstancePolicyRequestMetricOperator) Values() []UpdateInstancePolicyRequestMetricOperator {
	return []UpdateInstancePolicyRequestMetricOperator{
		"operator_unknown",
		"operator_greater_than",
		"operator_less_than",
	}
}

func (enum UpdateInstancePolicyRequestMetricOperator) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UpdateInstancePolicyRequestMetricOperator) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UpdateInstancePolicyRequestMetricOperator(UpdateInstancePolicyRequestMetricOperator(tmp).String())
	return nil
}

type VolumeInstanceTemplateVolumeType string

const (
	// Unknown volume type.
	VolumeInstanceTemplateVolumeTypeUnknownVolumeType = VolumeInstanceTemplateVolumeType("unknown_volume_type")
	// Local storage.
	VolumeInstanceTemplateVolumeTypeLSSD = VolumeInstanceTemplateVolumeType("l_ssd")
	// Block storage.
	VolumeInstanceTemplateVolumeTypeSbs = VolumeInstanceTemplateVolumeType("sbs")
)

func (enum VolumeInstanceTemplateVolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return string(VolumeInstanceTemplateVolumeTypeUnknownVolumeType)
	}
	return string(enum)
}

func (enum VolumeInstanceTemplateVolumeType) Values() []VolumeInstanceTemplateVolumeType {
	return []VolumeInstanceTemplateVolumeType{
		"unknown_volume_type",
		"l_ssd",
		"sbs",
	}
}

func (enum VolumeInstanceTemplateVolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeInstanceTemplateVolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeInstanceTemplateVolumeType(VolumeInstanceTemplateVolumeType(tmp).String())
	return nil
}

// VolumeInstanceTemplateFromEmpty: volume instance template from empty.
type VolumeInstanceTemplateFromEmpty struct {
	Size scw.Size `json:"size"`
}

// VolumeInstanceTemplateFromSnapshot: volume instance template from snapshot.
type VolumeInstanceTemplateFromSnapshot struct {
	Size *scw.Size `json:"size"`

	SnapshotID string `json:"snapshot_id"`
}

// Capacity: capacity.
type Capacity struct {
	// MaxReplicas: maximum count of Instances for the Instance group.
	MaxReplicas uint32 `json:"max_replicas"`

	// MinReplicas: minimum count of Instances for the Instance group.
	MinReplicas uint32 `json:"min_replicas"`

	// CooldownDelay: time (in seconds) after a scaling action during which requests to carry out a new scaling action will be denied.
	CooldownDelay *scw.Duration `json:"cooldown_delay"`
}

// Loadbalancer: loadbalancer.
type Loadbalancer struct {
	// ID: load Balancer ID.
	ID string `json:"id"`

	// BackendIDs: load Balancer backend IDs.
	BackendIDs []string `json:"backend_ids"`

	// PrivateNetworkID: ID of the Private Network attached to the Load Balancer.
	PrivateNetworkID string `json:"private_network_id"`
}

// Metric: metric.
type Metric struct {
	// Name: name or description of the metric policy.
	Name string `json:"name"`

	// ManagedMetric: managed metric to use for this policy. These are available by default in Cockpit without any configuration or `node_exporter`. The chosen metric forms the basis of the condition that will be checked to determine whether a scaling action should be triggered.
	// Default value: managed_metric_unknown
	// Precisely one of ManagedMetric, CockpitMetricName must be set.
	ManagedMetric *MetricManagedMetric `json:"managed_metric,omitempty"`

	// CockpitMetricName: custom metric to use for this policy. This must be stored in Scaleway Cockpit. The metric forms the basis of the condition that will be checked to determine whether a scaling action should be triggered.
	// Precisely one of ManagedMetric, CockpitMetricName must be set.
	CockpitMetricName *string `json:"cockpit_metric_name,omitempty"`

	// Operator: operator used when comparing the threshold value of the chosen `metric` to the actual sampled and aggregated value.
	// Default value: operator_unknown
	Operator MetricOperator `json:"operator"`

	// Aggregate: how the values sampled for the `metric` should be aggregated.
	// Default value: aggregate_unknown
	Aggregate MetricAggregate `json:"aggregate"`

	// SamplingRangeMin: interval of time, in minutes, during which metric is sampled.
	SamplingRangeMin uint32 `json:"sampling_range_min"`

	// Threshold: threshold value to measure the aggregated sampled `metric` value against. Combined with the `operator` field, determines whether a scaling action should be triggered.
	Threshold float32 `json:"threshold"`
}

// VolumeInstanceTemplate: volume instance template.
type VolumeInstanceTemplate struct {
	// Name: name of the volume.
	Name string `json:"name"`

	// PerfIops: the maximum IO/s expected, according to the different options available in stock (`5000 | 15000`).
	// Precisely one of PerfIops must be set.
	PerfIops *uint32 `json:"perf_iops,omitempty"`

	// FromEmpty: specify the size of the new volume if creating a new one from scratch.
	// Precisely one of FromEmpty, FromSnapshot must be set.
	FromEmpty *VolumeInstanceTemplateFromEmpty `json:"from_empty,omitempty"`

	// FromSnapshot: specify the snapshot ID of the original snapshot.
	// Precisely one of FromEmpty, FromSnapshot must be set.
	FromSnapshot *VolumeInstanceTemplateFromSnapshot `json:"from_snapshot,omitempty"`

	// Tags: list of tags assigned to the volume.
	Tags []string `json:"tags"`

	// Boot: force the Instance to boot on this volume.
	Boot bool `json:"boot"`

	// VolumeType: type of the volume.
	// Default value: unknown_volume_type
	VolumeType VolumeInstanceTemplateVolumeType `json:"volume_type"`
}

// InstanceGroupEvent: instance group event.
type InstanceGroupEvent struct {
	// ID: instance group event ID.
	ID string `json:"id"`

	// Source: log source.
	// Default value: unknown_source
	Source InstanceGroupEventSource `json:"source"`

	// Level: the severity of the log.
	// Default value: info
	Level InstanceGroupEventLevel `json:"level"`

	// Name: log title.
	Name string `json:"name"`

	// CreatedAt: date and time of the log.
	CreatedAt *time.Time `json:"created_at"`

	// Details: full text of the log.
	Details *string `json:"details"`
}

// InstanceGroup: instance group.
type InstanceGroup struct {
	// ID: instance group ID.
	ID string `json:"id"`

	// ProjectID: project ID of the Instance group.
	ProjectID string `json:"project_id"`

	// Name: name of the Instance group.
	Name string `json:"name"`

	// Tags: instance group tags.
	Tags []string `json:"tags"`

	// InstanceTemplateID: template ID (ID of the Instance template to attach to the Instance group).
	InstanceTemplateID string `json:"instance_template_id"`

	// Capacity: specification of the minimum and maximum replicas for the Instance group, and the cooldown interval between two scaling events.
	Capacity *Capacity `json:"capacity"`

	// Loadbalancer: specification of the Load Balancer linked to the Instance group.
	Loadbalancer *Loadbalancer `json:"loadbalancer"`

	// ErrorMessages: any configuration errors for dependencies (Load Balancer, Private Network, Instance template etc.).
	ErrorMessages []string `json:"error_messages"`

	// CreatedAt: date on which the Instance group was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date on which the Instance group was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// Zone: zone for this resource.
	Zone scw.Zone `json:"zone"`
}

// InstancePolicy: instance policy.
type InstancePolicy struct {
	// ID: scaling policy ID.
	ID string `json:"id"`

	// Name: name of scaling policy.
	Name string `json:"name"`

	// Metric: managed metric to use for this policy. These are available by default in Cockpit without any configuration or `node_exporter`. The chosen metric forms the basis of the condition that will be checked to determine whether a scaling action should be triggered.
	// Precisely one of Metric must be set.
	Metric *Metric `json:"metric,omitempty"`

	// Action: action to execute when the metric-based condition is met.
	// Default value: unknown_action
	Action InstancePolicyAction `json:"action"`

	// Type: how to use the number defined in `value` when determining by how many Instances to scale up/down.
	// Default value: unknown_type
	Type InstancePolicyType `json:"type"`

	// Value: number representing the magnitude of the scaling action to take for the Instance group.
	Value uint32 `json:"value"`

	// Priority: priority of this policy compared to all other scaling policies. The lower the number, the higher the priority (higher priority will be processed sooner in the order).
	Priority uint32 `json:"priority"`

	// InstanceGroupID: instance group ID related to this policy.
	InstanceGroupID string `json:"instance_group_id"`

	// Zone: zone for this resource.
	Zone scw.Zone `json:"zone"`
}

// InstanceTemplate: instance template.
type InstanceTemplate struct {
	// ID: ID of Instance template resource.
	ID string `json:"id"`

	// CommercialType: name of Instance commercial type.
	CommercialType string `json:"commercial_type"`

	// ImageID: instance image ID. Can be an ID of a marketplace or personal image. This image must be compatible with `volume` and `commercial_type` template.
	ImageID *string `json:"image_id"`

	// Volumes: template of Instance volume.
	Volumes map[string]*VolumeInstanceTemplate `json:"volumes"`

	// Tags: list of tags for the Instance template.
	Tags []string `json:"tags"`

	// SecurityGroupID: instance security group ID (optional).
	SecurityGroupID *string `json:"security_group_id"`

	// PlacementGroupID: instance placement group ID. This is optional, but it is highly recommended to set a preference for Instance location within Availability Zone.
	PlacementGroupID *string `json:"placement_group_id"`

	// PublicIPsV4Count: number of flexible IPv4 addresses to attach to the new Instance.
	PublicIPsV4Count *uint32 `json:"public_ips_v4_count"`

	// PublicIPsV6Count: number of flexible IPv6 addresses to attach to the new Instance.
	PublicIPsV6Count *uint32 `json:"public_ips_v6_count"`

	// ProjectID: ID of the Project containing the Instance template resource.
	ProjectID string `json:"project_id"`

	// Name: name of Instance template.
	Name string `json:"name"`

	// PrivateNetworkIDs: private Network IDs to attach to the new Instance.
	PrivateNetworkIDs []string `json:"private_network_ids"`

	// Status: status of Instance template.
	// Default value: unknown_status
	Status InstanceTemplateStatus `json:"status"`

	// CloudInit: cloud-config file must be passed in Base64 format. Cloud-config files are special scripts designed to be run by the cloud-init process. These are generally used for initial configuration on the very first boot of a server.
	CloudInit *[]byte `json:"cloud_init"`

	// CreatedAt: date on which the Instance template was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date on which the Instance template was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// Zone: zone for this resource.
	Zone scw.Zone `json:"zone"`
}

// UpdateInstanceGroupRequestCapacity: update instance group request capacity.
type UpdateInstanceGroupRequestCapacity struct {
	// MaxReplicas: maximum count of Instances for the Instance group.
	MaxReplicas *uint32 `json:"max_replicas"`

	// MinReplicas: minimum count of Instances for the Instance group.
	MinReplicas *uint32 `json:"min_replicas"`

	// CooldownDelay: time (in seconds) after a scaling action during which requests to carry out a new scaling action will be denied.
	CooldownDelay *scw.Duration `json:"cooldown_delay"`
}

// UpdateInstanceGroupRequestLoadbalancer: update instance group request loadbalancer.
type UpdateInstanceGroupRequestLoadbalancer struct {
	// BackendIDs: load Balancer backend IDs.
	BackendIDs *[]string `json:"backend_ids"`
}

// UpdateInstancePolicyRequestMetric: update instance policy request metric.
type UpdateInstancePolicyRequestMetric struct {
	// Name: name or description of your metric policy.
	Name *string `json:"name"`

	// ManagedMetric: managed metric to use for this policy. These are available by default in Cockpit without any configuration or `node_exporter`. The chosen metric forms the basis of the condition that will be checked to determine whether a scaling action should be triggered.
	// Default value: managed_metric_unknown
	// Precisely one of ManagedMetric, CockpitMetricName must be set.
	ManagedMetric *UpdateInstancePolicyRequestMetricManagedMetric `json:"managed_metric,omitempty"`

	// CockpitMetricName: custom metric to use for this policy. This must be stored in Scaleway Cockpit. The metric forms the basis of the condition that will be checked to determine whether a scaling action should be triggered.
	// Precisely one of ManagedMetric, CockpitMetricName must be set.
	CockpitMetricName *string `json:"cockpit_metric_name,omitempty"`

	// Operator: operator used when comparing the threshold value of the chosen `metric` to the actual sampled and aggregated value.
	// Default value: operator_unknown
	Operator UpdateInstancePolicyRequestMetricOperator `json:"operator"`

	// Aggregate: how the values sampled for the `metric` should be aggregated.
	// Default value: aggregate_unknown
	Aggregate UpdateInstancePolicyRequestMetricAggregate `json:"aggregate"`

	// SamplingRangeMin: interval of time, in minutes, during which metric is sampled.
	SamplingRangeMin *uint32 `json:"sampling_range_min"`

	// Threshold: threshold value to measure the aggregated sampled `metric` value against. Combined with the `operator` field, determines whether a scaling action should be triggered.
	Threshold *float32 `json:"threshold"`
}

// CreateInstanceGroupRequest: create instance group request.
type CreateInstanceGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ProjectID: project ID to filter for, only Instance groups from this Project will be returned.
	ProjectID string `json:"project_id"`

	// Name: name of Instance group.
	Name string `json:"name"`

	// Tags: list of tags for the Instance group.
	Tags []string `json:"tags"`

	// TemplateID: template ID (ID of the Instance template to attach to the Instance group).
	TemplateID string `json:"template_id"`

	// Capacity: specification of the minimum and maximum replicas for the Instance group, and the cooldown interval between two scaling events.
	Capacity *Capacity `json:"capacity"`

	// Loadbalancer: specification of the Load Balancer to link to the Instance group.
	Loadbalancer *Loadbalancer `json:"loadbalancer"`
}

// CreateInstancePolicyRequest: create instance policy request.
type CreateInstancePolicyRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Name: name of the policy.
	Name string `json:"name"`

	// Metric: cockpit metric to use when determining whether to trigger a scale up/down action.
	// Precisely one of Metric must be set.
	Metric *Metric `json:"metric,omitempty"`

	// Action: action to execute when the metric-based condition is met.
	// Default value: unknown_action
	Action InstancePolicyAction `json:"action"`

	// Type: how to use the number defined in `value` when determining by how many Instances to scale up/down.
	// Default value: unknown_type
	Type InstancePolicyType `json:"type"`

	// Value: value representing the magnitude of the scaling action to take for the Instance group. Depending on the `type` parameter, this number could represent a total number of Instances in the group, a number of Instances to add, or a percentage to scale the group by.
	Value uint32 `json:"value"`

	// Priority: priority of this policy compared to all other scaling policies. This determines the processing order. The lower the number, the higher the priority.
	Priority uint32 `json:"priority"`

	// InstanceGroupID: instance group ID related to this policy.
	InstanceGroupID string `json:"instance_group_id"`
}

// CreateInstanceTemplateRequest: create instance template request.
type CreateInstanceTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// CommercialType: name of Instance commercial type.
	CommercialType string `json:"commercial_type"`

	// ImageID: instance image ID. Can be an ID of a marketplace or personal image. This image must be compatible with `volume` and `commercial_type` template.
	ImageID *string `json:"image_id,omitempty"`

	// Volumes: template of Instance volume.
	Volumes map[string]*VolumeInstanceTemplate `json:"volumes"`

	// Tags: list of tags for the Instance template.
	Tags []string `json:"tags"`

	// SecurityGroupID: instance security group ID (optional).
	SecurityGroupID *string `json:"security_group_id,omitempty"`

	// PlacementGroupID: instance placement group ID. This is optional, but it is highly recommended to set a preference for Instance location within Availability Zone.
	PlacementGroupID *string `json:"placement_group_id,omitempty"`

	// PublicIPsV4Count: number of flexible IPv4 addresses to attach to the new Instance.
	PublicIPsV4Count *uint32 `json:"public_ips_v4_count,omitempty"`

	// PublicIPsV6Count: number of flexible IPv6 addresses to attach to the new Instance.
	PublicIPsV6Count *uint32 `json:"public_ips_v6_count,omitempty"`

	// ProjectID: ID of the Project containing the Instance template resource.
	ProjectID string `json:"project_id"`

	// Name: name of Instance template.
	Name string `json:"name"`

	// PrivateNetworkIDs: private Network IDs to attach to the new Instance.
	PrivateNetworkIDs []string `json:"private_network_ids"`

	// CloudInit: cloud-config file must be passed in Base64 format. Cloud-config files are special scripts designed to be run by the cloud-init process. These are generally used for initial configuration on the very first boot of a server.
	CloudInit *[]byte `json:"cloud_init,omitempty"`
}

// DeleteInstanceGroupRequest: delete instance group request.
type DeleteInstanceGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// InstanceGroupID: ID of the Instance group to delete.
	InstanceGroupID string `json:"-"`
}

// DeleteInstancePolicyRequest: delete instance policy request.
type DeleteInstancePolicyRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PolicyID: ID of the policy to delete.
	PolicyID string `json:"-"`
}

// DeleteInstanceTemplateRequest: delete instance template request.
type DeleteInstanceTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: ID of the template to delete.
	TemplateID string `json:"-"`
}

// GetInstanceGroupRequest: get instance group request.
type GetInstanceGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// InstanceGroupID: ID of the requested Instance group.
	InstanceGroupID string `json:"-"`
}

// GetInstancePolicyRequest: get instance policy request.
type GetInstancePolicyRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PolicyID: policy ID.
	PolicyID string `json:"-"`
}

// GetInstanceTemplateRequest: get instance template request.
type GetInstanceTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: template ID of the resource.
	TemplateID string `json:"-"`
}

// ListInstanceGroupEventsRequest: list instance group events request.
type ListInstanceGroupEventsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// InstanceGroupID: list all event logs for the Instance group ID.
	InstanceGroupID string `json:"-"`

	// OrderBy: sort order of Instance groups in the response.
	// Default value: created_at_desc
	OrderBy ListInstanceGroupEventsRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of Instance groups to return per page.
	PageSize *uint32 `json:"-"`
}

// ListInstanceGroupEventsResponse: list instance group events response.
type ListInstanceGroupEventsResponse struct {
	// InstanceEvents: paginated list of Instance groups.
	InstanceEvents []*InstanceGroupEvent `json:"instance_events"`

	// TotalCount: count of all Instance groups matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInstanceGroupEventsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInstanceGroupEventsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListInstanceGroupEventsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.InstanceEvents = append(r.InstanceEvents, results.InstanceEvents...)
	r.TotalCount += uint64(len(results.InstanceEvents))
	return uint64(len(results.InstanceEvents)), nil
}

// ListInstanceGroupsRequest: list instance groups request.
type ListInstanceGroupsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: sort order of Instance groups in the response.
	// Default value: created_at_desc
	OrderBy ListInstanceGroupsRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of Instance groups to return per page.
	PageSize *uint32 `json:"-"`
}

// ListInstanceGroupsResponse: list instance groups response.
type ListInstanceGroupsResponse struct {
	// InstanceGroups: paginated list of Instance groups.
	InstanceGroups []*InstanceGroup `json:"instance_groups"`

	// TotalCount: count of all Instance groups matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInstanceGroupsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInstanceGroupsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListInstanceGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.InstanceGroups = append(r.InstanceGroups, results.InstanceGroups...)
	r.TotalCount += uint64(len(results.InstanceGroups))
	return uint64(len(results.InstanceGroups)), nil
}

// ListInstancePoliciesRequest: list instance policies request.
type ListInstancePoliciesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: sort order of Instance groups in the response.
	// Default value: created_at_desc
	OrderBy ListInstancePoliciesRequestOrderBy `json:"-"`

	// InstanceGroupID: instance group ID.
	InstanceGroupID string `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of scaling policies to return per page.
	PageSize *uint32 `json:"-"`
}

// ListInstancePoliciesResponse: list instance policies response.
type ListInstancePoliciesResponse struct {
	// Policies: paginated list of policies.
	Policies []*InstancePolicy `json:"policies"`

	// TotalCount: count of all policies matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInstancePoliciesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInstancePoliciesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListInstancePoliciesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Policies = append(r.Policies, results.Policies...)
	r.TotalCount += uint64(len(results.Policies))
	return uint64(len(results.Policies)), nil
}

// ListInstanceTemplatesRequest: list instance templates request.
type ListInstanceTemplatesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: sort order of Instance groups in the response.
	// Default value: created_at_desc
	OrderBy ListInstanceTemplatesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of Instance groups to return per page.
	PageSize *uint32 `json:"-"`
}

// ListInstanceTemplatesResponse: list instance templates response.
type ListInstanceTemplatesResponse struct {
	// TotalCount: count of all templates matching the requested criteria.
	TotalCount uint64 `json:"total_count"`

	// InstanceTemplates: paginated list of Instance templates.
	InstanceTemplates []*InstanceTemplate `json:"instance_templates"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInstanceTemplatesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInstanceTemplatesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListInstanceTemplatesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.InstanceTemplates = append(r.InstanceTemplates, results.InstanceTemplates...)
	r.TotalCount += uint64(len(results.InstanceTemplates))
	return uint64(len(results.InstanceTemplates)), nil
}

// UpdateInstanceGroupRequest: update instance group request.
type UpdateInstanceGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// InstanceGroupID: instance group ID to update.
	InstanceGroupID string `json:"-"`

	// Name: name of Instance group.
	Name *string `json:"name,omitempty"`

	// Tags: list of tags for the Load Balancer.
	Tags *[]string `json:"tags,omitempty"`

	// Capacity: specification of the minimum and maximum replicas for the Instance group, and the cooldown interval between two scaling events.
	Capacity *UpdateInstanceGroupRequestCapacity `json:"capacity,omitempty"`

	// Loadbalancer: specification of the Load Balancer to link to the Instance group.
	Loadbalancer *UpdateInstanceGroupRequestLoadbalancer `json:"loadbalancer,omitempty"`
}

// UpdateInstancePolicyRequest: update instance policy request.
type UpdateInstancePolicyRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PolicyID: policy ID to update.
	PolicyID string `json:"-"`

	// Name: policy name to update.
	Name *string `json:"name,omitempty"`

	// Metric: metric specification to update (Cockpit metric to use when determining whether to trigger a scale up/down action).
	// Precisely one of Metric must be set.
	Metric *UpdateInstancePolicyRequestMetric `json:"metric,omitempty"`

	// Action: action to update (action to execute when the metric-based condition is met).
	// Default value: unknown_action
	Action InstancePolicyAction `json:"action"`

	// Type: type to update (how to use the number defined in `value` when determining by how many Instances to scale up/down).
	// Default value: unknown_type
	Type InstancePolicyType `json:"type"`

	// Value: value to update (number representing the magnitude of the scaling action to take for the Instance group).
	Value *uint32 `json:"value,omitempty"`

	// Priority: priority to update (priority of this policy compared to all other scaling policies. The lower the number, the higher the priority).
	Priority *uint32 `json:"priority,omitempty"`
}

// UpdateInstanceTemplateRequest: update instance template request.
type UpdateInstanceTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: template ID of the resource.
	TemplateID string `json:"-"`

	// CommercialType: name of Instance commercial type.
	CommercialType *string `json:"commercial_type,omitempty"`

	// ImageID: instance image ID. Can be an ID of a marketplace or personal image. This image must be compatible with `volume` and `commercial_type` template.
	ImageID *string `json:"image_id,omitempty"`

	// Volumes: template of Instance volume.
	Volumes map[string]*VolumeInstanceTemplate `json:"volumes"`

	// Tags: list of tags for the Instance template.
	Tags *[]string `json:"tags,omitempty"`

	// SecurityGroupID: instance security group ID (optional).
	SecurityGroupID *string `json:"security_group_id,omitempty"`

	// PlacementGroupID: instance placement group ID. This is optional, but it is highly recommended to set a preference for Instance location within Availability Zone.
	PlacementGroupID *string `json:"placement_group_id,omitempty"`

	// PublicIPsV4Count: number of flexible IPv4 addresses to attach to the new Instance.
	PublicIPsV4Count *uint32 `json:"public_ips_v4_count,omitempty"`

	// PublicIPsV6Count: number of flexible IPv6 addresses to attach to the new Instance.
	PublicIPsV6Count *uint32 `json:"public_ips_v6_count,omitempty"`

	// Name: name of Instance template.
	Name *string `json:"name,omitempty"`

	// PrivateNetworkIDs: private Network IDs to attach to the new Instance.
	PrivateNetworkIDs *[]string `json:"private_network_ids,omitempty"`

	// CloudInit: cloud-config file must be passed in Base64 format. Cloud-config files are special scripts designed to be run by the cloud-init process. These are generally used for initial configuration on the very first boot of a server.
	CloudInit *[]byte `json:"cloud_init,omitempty"`
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
	return []scw.Zone{}
}

// GetInstanceGroup: Retrieve information about an existing Instance group, specified by its `instance_group_id`. Its full details, including errors, are returned in the response object.
func (s *API) GetInstanceGroup(req *GetInstanceGroupRequest, opts ...scw.RequestOption) (*InstanceGroup, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceGroupID) == "" {
		return nil, errors.New("field InstanceGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-groups/" + fmt.Sprint(req.InstanceGroupID) + "",
	}

	var resp InstanceGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateInstanceGroup: Create a new Instance group. You must specify a `template_id`, capacity and Load Balancer object.
func (s *API) CreateInstanceGroup(req *CreateInstanceGroupRequest, opts ...scw.RequestOption) (*InstanceGroup, error) {
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
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-groups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp InstanceGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInstanceGroups: List all Instance groups, for a Scaleway Organization or Scaleway Project. By default, the Instance groups returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListInstanceGroups(req *ListInstanceGroupsRequest, opts ...scw.RequestOption) (*ListInstanceGroupsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-groups",
		Query:  query,
	}

	var resp ListInstanceGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateInstanceGroup: Update the parameters of an existing Instance group, specified by its `instance_group_id`.
func (s *API) UpdateInstanceGroup(req *UpdateInstanceGroupRequest, opts ...scw.RequestOption) (*InstanceGroup, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceGroupID) == "" {
		return nil, errors.New("field InstanceGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-groups/" + fmt.Sprint(req.InstanceGroupID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp InstanceGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteInstanceGroup: Delete an existing Instance group, specified by its `instance_group_id`. Deleting an Instance group is permanent, and cannot be undone.
func (s *API) DeleteInstanceGroup(req *DeleteInstanceGroupRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceGroupID) == "" {
		return errors.New("field InstanceGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-groups/" + fmt.Sprint(req.InstanceGroupID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateInstanceTemplate: Create a new Instance template. This specifies the details of the Instance (commercial type, zone, image, volumes etc.) that will be in the Instance group.
func (s *API) CreateInstanceTemplate(req *CreateInstanceTemplateRequest, opts ...scw.RequestOption) (*InstanceTemplate, error) {
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
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-templates",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp InstanceTemplate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateInstanceTemplate: Update an Instance template, such as its commercial offer type, image or volume template.
func (s *API) UpdateInstanceTemplate(req *UpdateInstanceTemplateRequest, opts ...scw.RequestOption) (*InstanceTemplate, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return nil, errors.New("field TemplateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-templates/" + fmt.Sprint(req.TemplateID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp InstanceTemplate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInstanceTemplate: Get an existing Instance template from its `template_id`.
func (s *API) GetInstanceTemplate(req *GetInstanceTemplateRequest, opts ...scw.RequestOption) (*InstanceTemplate, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return nil, errors.New("field TemplateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-templates/" + fmt.Sprint(req.TemplateID) + "",
	}

	var resp InstanceTemplate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteInstanceTemplate: Delete an existing Instance template. This action is permanent and cannot be undone.
func (s *API) DeleteInstanceTemplate(req *DeleteInstanceTemplateRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return errors.New("field TemplateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-templates/" + fmt.Sprint(req.TemplateID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListInstanceTemplates: List all Instance templates, for a Scaleway Organization or Scaleway Project. By default, the Instance templates returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListInstanceTemplates(req *ListInstanceTemplatesRequest, opts ...scw.RequestOption) (*ListInstanceTemplatesResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-templates",
		Query:  query,
	}

	var resp ListInstanceTemplatesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateInstancePolicy: Create a new scaling policy. You must specify a `policy_id`, capacity and Load Balancer object.
func (s *API) CreateInstancePolicy(req *CreateInstancePolicyRequest, opts ...scw.RequestOption) (*InstancePolicy, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-policies",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp InstancePolicy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateInstancePolicy: Update the parameters of an existing scaling policy, specified by its `policy_id`.
func (s *API) UpdateInstancePolicy(req *UpdateInstancePolicyRequest, opts ...scw.RequestOption) (*InstancePolicy, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PolicyID) == "" {
		return nil, errors.New("field PolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-policies/" + fmt.Sprint(req.PolicyID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp InstancePolicy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInstancePolicies: List all scaling policies, for a Scaleway Organization or Scaleway Project. By default, the policies returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListInstancePolicies(req *ListInstancePoliciesRequest, opts ...scw.RequestOption) (*ListInstancePoliciesResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "instance_group_id", req.InstanceGroupID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-policies",
		Query:  query,
	}

	var resp ListInstancePoliciesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInstancePolicy: Retrieve information about an existing scaling policy, specified by its `policy_id`. Its full details are returned in the response object.
func (s *API) GetInstancePolicy(req *GetInstancePolicyRequest, opts ...scw.RequestOption) (*InstancePolicy, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PolicyID) == "" {
		return nil, errors.New("field PolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-policies/" + fmt.Sprint(req.PolicyID) + "",
	}

	var resp InstancePolicy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteInstancePolicy: Delete an existing scaling policy, specified by its `policy_id`. Deleting a scaling policy is permanent, and cannot be undone.
func (s *API) DeleteInstancePolicy(req *DeleteInstancePolicyRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PolicyID) == "" {
		return errors.New("field PolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-policies/" + fmt.Sprint(req.PolicyID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListInstanceGroupEvents: List all events for a given Instance group. By default, the events are ordered by creation date in descending order, though this can be modified via the `order_by` field.
func (s *API) ListInstanceGroupEvents(req *ListInstanceGroupEventsRequest, opts ...scw.RequestOption) (*ListInstanceGroupEventsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceGroupID) == "" {
		return nil, errors.New("field InstanceGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/autoscaling/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/instance-groups/" + fmt.Sprint(req.InstanceGroupID) + "/events",
		Query:  query,
	}

	var resp ListInstanceGroupEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

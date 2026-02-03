// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package k8s provides methods and message types of the k8s v1 API.
package k8s

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
	defaultK8sRetryInterval = 15 * time.Second
	defaultK8sTimeout       = 15 * time.Minute
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

type AutoscalerEstimator string

const (
	AutoscalerEstimatorUnknownEstimator = AutoscalerEstimator("unknown_estimator")
	AutoscalerEstimatorBinpacking       = AutoscalerEstimator("binpacking")
)

func (enum AutoscalerEstimator) String() string {
	if enum == "" {
		// return default value if empty
		return string(AutoscalerEstimatorUnknownEstimator)
	}
	return string(enum)
}

func (enum AutoscalerEstimator) Values() []AutoscalerEstimator {
	return []AutoscalerEstimator{
		"unknown_estimator",
		"binpacking",
	}
}

func (enum AutoscalerEstimator) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AutoscalerEstimator) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AutoscalerEstimator(AutoscalerEstimator(tmp).String())
	return nil
}

// Kubernetes autoscaler strategy to fit pods into cluster nodes (https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders).
type AutoscalerExpander string

const (
	AutoscalerExpanderUnknownExpander = AutoscalerExpander("unknown_expander")
	AutoscalerExpanderRandom          = AutoscalerExpander("random")
	AutoscalerExpanderMostPods        = AutoscalerExpander("most_pods")
	AutoscalerExpanderLeastWaste      = AutoscalerExpander("least_waste")
	AutoscalerExpanderPriority        = AutoscalerExpander("priority")
	AutoscalerExpanderPrice           = AutoscalerExpander("price")
)

func (enum AutoscalerExpander) String() string {
	if enum == "" {
		// return default value if empty
		return string(AutoscalerExpanderUnknownExpander)
	}
	return string(enum)
}

func (enum AutoscalerExpander) Values() []AutoscalerExpander {
	return []AutoscalerExpander{
		"unknown_expander",
		"random",
		"most_pods",
		"least_waste",
		"priority",
		"price",
	}
}

func (enum AutoscalerExpander) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AutoscalerExpander) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AutoscalerExpander(AutoscalerExpander(tmp).String())
	return nil
}

type CNI string

const (
	CNIUnknownCni = CNI("unknown_cni")
	// Cilium CNI will be configured (https://github.com/cilium/cilium).
	CNICilium = CNI("cilium")
	// Calico CNI will be configured (https://github.com/projectcalico/calico).
	CNICalico  = CNI("calico")
	CNIWeave   = CNI("weave")
	CNIFlannel = CNI("flannel")
	// Kilo CNI will be configured (https://github.com/squat/kilo/). Note that this CNI is only available for Kosmos clusters.
	CNIKilo = CNI("kilo")
	// Does not install any CNI. This feature is only available through a ticket and is not covered by support.
	CNINone = CNI("none")
	// Cilium CNI will be configured in native routing mode (https://docs.cilium.io/en/stable/network/concepts/routing/#native-routing).
	CNICiliumNative = CNI("cilium_native")
)

func (enum CNI) String() string {
	if enum == "" {
		// return default value if empty
		return string(CNIUnknownCni)
	}
	return string(enum)
}

func (enum CNI) Values() []CNI {
	return []CNI{
		"unknown_cni",
		"cilium",
		"calico",
		"weave",
		"flannel",
		"kilo",
		"none",
		"cilium_native",
	}
}

func (enum CNI) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CNI) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CNI(CNI(tmp).String())
	return nil
}

type ClusterStatus string

const (
	ClusterStatusUnknown = ClusterStatus("unknown")
	// Cluster is provisioning.
	ClusterStatusCreating = ClusterStatus("creating")
	// Cluster is ready to use.
	ClusterStatusReady = ClusterStatus("ready")
	// Cluster is waiting to be processed for deletion.
	ClusterStatusDeleting = ClusterStatus("deleting")
	ClusterStatusDeleted  = ClusterStatus("deleted")
	// Cluster is updating its own configuration, it can be a version upgrade too.
	ClusterStatusUpdating = ClusterStatus("updating")
	// Cluster is locked because an abuse has been detected or reported.
	ClusterStatusLocked = ClusterStatus("locked")
	// Cluster has no associated pool and has been shutdown.
	ClusterStatusPoolRequired = ClusterStatus("pool_required")
)

func (enum ClusterStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ClusterStatusUnknown)
	}
	return string(enum)
}

func (enum ClusterStatus) Values() []ClusterStatus {
	return []ClusterStatus{
		"unknown",
		"creating",
		"ready",
		"deleting",
		"deleted",
		"updating",
		"locked",
		"pool_required",
	}
}

func (enum ClusterStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ClusterStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ClusterStatus(ClusterStatus(tmp).String())
	return nil
}

type ClusterTypeAvailability string

const (
	// Type is available in quantity.
	ClusterTypeAvailabilityAvailable = ClusterTypeAvailability("available")
	// Limited availability.
	ClusterTypeAvailabilityScarce = ClusterTypeAvailability("scarce")
	// Out of stock.
	ClusterTypeAvailabilityShortage = ClusterTypeAvailability("shortage")
)

func (enum ClusterTypeAvailability) String() string {
	if enum == "" {
		// return default value if empty
		return string(ClusterTypeAvailabilityAvailable)
	}
	return string(enum)
}

func (enum ClusterTypeAvailability) Values() []ClusterTypeAvailability {
	return []ClusterTypeAvailability{
		"available",
		"scarce",
		"shortage",
	}
}

func (enum ClusterTypeAvailability) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ClusterTypeAvailability) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ClusterTypeAvailability(ClusterTypeAvailability(tmp).String())
	return nil
}

type ClusterTypeResiliency string

const (
	ClusterTypeResiliencyUnknownResiliency = ClusterTypeResiliency("unknown_resiliency")
	// The control plane is rescheduled on other machines in case of failure of a lower layer.
	ClusterTypeResiliencyStandard = ClusterTypeResiliency("standard")
	// The control plane has replicas to ensure service continuity in case of failure of a lower layer.
	ClusterTypeResiliencyHighAvailability = ClusterTypeResiliency("high_availability")
)

func (enum ClusterTypeResiliency) String() string {
	if enum == "" {
		// return default value if empty
		return string(ClusterTypeResiliencyUnknownResiliency)
	}
	return string(enum)
}

func (enum ClusterTypeResiliency) Values() []ClusterTypeResiliency {
	return []ClusterTypeResiliency{
		"unknown_resiliency",
		"standard",
		"high_availability",
	}
}

func (enum ClusterTypeResiliency) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ClusterTypeResiliency) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ClusterTypeResiliency(ClusterTypeResiliency(tmp).String())
	return nil
}

type ListClustersRequestOrderBy string

const (
	ListClustersRequestOrderByCreatedAtAsc  = ListClustersRequestOrderBy("created_at_asc")
	ListClustersRequestOrderByCreatedAtDesc = ListClustersRequestOrderBy("created_at_desc")
	ListClustersRequestOrderByUpdatedAtAsc  = ListClustersRequestOrderBy("updated_at_asc")
	ListClustersRequestOrderByUpdatedAtDesc = ListClustersRequestOrderBy("updated_at_desc")
	ListClustersRequestOrderByNameAsc       = ListClustersRequestOrderBy("name_asc")
	ListClustersRequestOrderByNameDesc      = ListClustersRequestOrderBy("name_desc")
	ListClustersRequestOrderByStatusAsc     = ListClustersRequestOrderBy("status_asc")
	ListClustersRequestOrderByStatusDesc    = ListClustersRequestOrderBy("status_desc")
	ListClustersRequestOrderByVersionAsc    = ListClustersRequestOrderBy("version_asc")
	ListClustersRequestOrderByVersionDesc   = ListClustersRequestOrderBy("version_desc")
)

func (enum ListClustersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListClustersRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListClustersRequestOrderBy) Values() []ListClustersRequestOrderBy {
	return []ListClustersRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
		"name_asc",
		"name_desc",
		"status_asc",
		"status_desc",
		"version_asc",
		"version_desc",
	}
}

func (enum ListClustersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListClustersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListClustersRequestOrderBy(ListClustersRequestOrderBy(tmp).String())
	return nil
}

type ListNodesRequestOrderBy string

const (
	ListNodesRequestOrderByCreatedAtAsc  = ListNodesRequestOrderBy("created_at_asc")
	ListNodesRequestOrderByCreatedAtDesc = ListNodesRequestOrderBy("created_at_desc")
	ListNodesRequestOrderByUpdatedAtAsc  = ListNodesRequestOrderBy("updated_at_asc")
	ListNodesRequestOrderByUpdatedAtDesc = ListNodesRequestOrderBy("updated_at_desc")
	ListNodesRequestOrderByNameAsc       = ListNodesRequestOrderBy("name_asc")
	ListNodesRequestOrderByNameDesc      = ListNodesRequestOrderBy("name_desc")
	ListNodesRequestOrderByStatusAsc     = ListNodesRequestOrderBy("status_asc")
	ListNodesRequestOrderByStatusDesc    = ListNodesRequestOrderBy("status_desc")
	ListNodesRequestOrderByVersionAsc    = ListNodesRequestOrderBy("version_asc")
	ListNodesRequestOrderByVersionDesc   = ListNodesRequestOrderBy("version_desc")
)

func (enum ListNodesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListNodesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListNodesRequestOrderBy) Values() []ListNodesRequestOrderBy {
	return []ListNodesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
		"name_asc",
		"name_desc",
		"status_asc",
		"status_desc",
		"version_asc",
		"version_desc",
	}
}

func (enum ListNodesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNodesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNodesRequestOrderBy(ListNodesRequestOrderBy(tmp).String())
	return nil
}

type ListPoolsRequestOrderBy string

const (
	ListPoolsRequestOrderByCreatedAtAsc  = ListPoolsRequestOrderBy("created_at_asc")
	ListPoolsRequestOrderByCreatedAtDesc = ListPoolsRequestOrderBy("created_at_desc")
	ListPoolsRequestOrderByUpdatedAtAsc  = ListPoolsRequestOrderBy("updated_at_asc")
	ListPoolsRequestOrderByUpdatedAtDesc = ListPoolsRequestOrderBy("updated_at_desc")
	ListPoolsRequestOrderByNameAsc       = ListPoolsRequestOrderBy("name_asc")
	ListPoolsRequestOrderByNameDesc      = ListPoolsRequestOrderBy("name_desc")
	ListPoolsRequestOrderByStatusAsc     = ListPoolsRequestOrderBy("status_asc")
	ListPoolsRequestOrderByStatusDesc    = ListPoolsRequestOrderBy("status_desc")
	ListPoolsRequestOrderByVersionAsc    = ListPoolsRequestOrderBy("version_asc")
	ListPoolsRequestOrderByVersionDesc   = ListPoolsRequestOrderBy("version_desc")
)

func (enum ListPoolsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPoolsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListPoolsRequestOrderBy) Values() []ListPoolsRequestOrderBy {
	return []ListPoolsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
		"name_asc",
		"name_desc",
		"status_asc",
		"status_desc",
		"version_asc",
		"version_desc",
	}
}

func (enum ListPoolsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPoolsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPoolsRequestOrderBy(ListPoolsRequestOrderBy(tmp).String())
	return nil
}

type MaintenanceWindowDayOfTheWeek string

const (
	MaintenanceWindowDayOfTheWeekAny       = MaintenanceWindowDayOfTheWeek("any")
	MaintenanceWindowDayOfTheWeekMonday    = MaintenanceWindowDayOfTheWeek("monday")
	MaintenanceWindowDayOfTheWeekTuesday   = MaintenanceWindowDayOfTheWeek("tuesday")
	MaintenanceWindowDayOfTheWeekWednesday = MaintenanceWindowDayOfTheWeek("wednesday")
	MaintenanceWindowDayOfTheWeekThursday  = MaintenanceWindowDayOfTheWeek("thursday")
	MaintenanceWindowDayOfTheWeekFriday    = MaintenanceWindowDayOfTheWeek("friday")
	MaintenanceWindowDayOfTheWeekSaturday  = MaintenanceWindowDayOfTheWeek("saturday")
	MaintenanceWindowDayOfTheWeekSunday    = MaintenanceWindowDayOfTheWeek("sunday")
)

func (enum MaintenanceWindowDayOfTheWeek) String() string {
	if enum == "" {
		// return default value if empty
		return string(MaintenanceWindowDayOfTheWeekAny)
	}
	return string(enum)
}

func (enum MaintenanceWindowDayOfTheWeek) Values() []MaintenanceWindowDayOfTheWeek {
	return []MaintenanceWindowDayOfTheWeek{
		"any",
		"monday",
		"tuesday",
		"wednesday",
		"thursday",
		"friday",
		"saturday",
		"sunday",
	}
}

func (enum MaintenanceWindowDayOfTheWeek) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MaintenanceWindowDayOfTheWeek) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MaintenanceWindowDayOfTheWeek(MaintenanceWindowDayOfTheWeek(tmp).String())
	return nil
}

type NodeStatus string

const (
	NodeStatusUnknown = NodeStatus("unknown")
	// Node is provisioning.
	NodeStatusCreating = NodeStatus("creating")
	// Node is unable to connect to apiserver.
	NodeStatusNotReady = NodeStatus("not_ready")
	// Node is ready to execute workload (marked schedulable by k8s scheduler).
	NodeStatusReady = NodeStatus("ready")
	// Node is waiting to be processed for deletion.
	NodeStatusDeleting = NodeStatus("deleting")
	NodeStatusDeleted  = NodeStatus("deleted")
	// Node is locked because an abuse has been detected or reported.
	NodeStatusLocked = NodeStatus("locked")
	// Node is rebooting.
	NodeStatusRebooting     = NodeStatus("rebooting")
	NodeStatusCreationError = NodeStatus("creation_error")
	NodeStatusUpgrading     = NodeStatus("upgrading")
	NodeStatusStarting      = NodeStatus("starting")
	NodeStatusRegistering   = NodeStatus("registering")
)

func (enum NodeStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(NodeStatusUnknown)
	}
	return string(enum)
}

func (enum NodeStatus) Values() []NodeStatus {
	return []NodeStatus{
		"unknown",
		"creating",
		"not_ready",
		"ready",
		"deleting",
		"deleted",
		"locked",
		"rebooting",
		"creation_error",
		"upgrading",
		"starting",
		"registering",
	}
}

func (enum NodeStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NodeStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NodeStatus(NodeStatus(tmp).String())
	return nil
}

type PoolStatus string

const (
	PoolStatusUnknown = PoolStatus("unknown")
	// Pool has the right amount of nodes and is ready to process the workload.
	PoolStatusReady = PoolStatus("ready")
	// Pool is waiting to be processed for deletion.
	PoolStatusDeleting = PoolStatus("deleting")
	PoolStatusDeleted  = PoolStatus("deleted")
	// Pool is growing or shrinking.
	PoolStatusScaling = PoolStatus("scaling")
	// Pool has some issues, check nodes.
	PoolStatusWarning = PoolStatus("warning")
	// Pool is locked because an abuse has been detected or reported.
	PoolStatusLocked = PoolStatus("locked")
	// Pool is upgrading its Kubernetes version.
	PoolStatusUpgrading = PoolStatus("upgrading")
)

func (enum PoolStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(PoolStatusUnknown)
	}
	return string(enum)
}

func (enum PoolStatus) Values() []PoolStatus {
	return []PoolStatus{
		"unknown",
		"ready",
		"deleting",
		"deleted",
		"scaling",
		"warning",
		"locked",
		"upgrading",
	}
}

func (enum PoolStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PoolStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PoolStatus(PoolStatus(tmp).String())
	return nil
}

type PoolVolumeType string

const (
	PoolVolumeTypeDefaultVolumeType = PoolVolumeType("default_volume_type")
	// Local Block Storage: your system is stored locally on your node hypervisor.
	PoolVolumeTypeLSSD = PoolVolumeType("l_ssd")
	// Remote Block Storage: your system is stored on a centralized and resilient cluster (deprecated: will use sbs_5k instead).
	PoolVolumeTypeBSSD = PoolVolumeType("b_ssd")
	// Remote Block Storage: your system is stored on a centralized and resilient cluster with up to 5k IOPS.
	PoolVolumeTypeSbs5k = PoolVolumeType("sbs_5k")
	// Remote Block Storage: your system is stored on a centralized and resilient cluster with up to 15k IOPS.
	PoolVolumeTypeSbs15k = PoolVolumeType("sbs_15k")
)

func (enum PoolVolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return string(PoolVolumeTypeDefaultVolumeType)
	}
	return string(enum)
}

func (enum PoolVolumeType) Values() []PoolVolumeType {
	return []PoolVolumeType{
		"default_volume_type",
		"l_ssd",
		"b_ssd",
		"sbs_5k",
		"sbs_15k",
	}
}

func (enum PoolVolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PoolVolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PoolVolumeType(PoolVolumeType(tmp).String())
	return nil
}

type Runtime string

const (
	RuntimeUnknownRuntime = Runtime("unknown_runtime")
	RuntimeDocker         = Runtime("docker")
	// Containerd Runtime will be configured (https://github.com/containerd/containerd).
	RuntimeContainerd = Runtime("containerd")
	RuntimeCrio       = Runtime("crio")
)

func (enum Runtime) String() string {
	if enum == "" {
		// return default value if empty
		return string(RuntimeUnknownRuntime)
	}
	return string(enum)
}

func (enum Runtime) Values() []Runtime {
	return []Runtime{
		"unknown_runtime",
		"docker",
		"containerd",
		"crio",
	}
}

func (enum Runtime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Runtime) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Runtime(Runtime(tmp).String())
	return nil
}

// MaintenanceWindow: maintenance window.
type MaintenanceWindow struct {
	// StartHour: start time of the two-hour maintenance window.
	StartHour uint32 `json:"start_hour"`

	// Day: day of the week for the maintenance window.
	// Default value: any
	Day MaintenanceWindowDayOfTheWeek `json:"day"`
}

// CreateClusterRequestPoolConfigUpgradePolicy: create cluster request pool config upgrade policy.
type CreateClusterRequestPoolConfigUpgradePolicy struct {
	// MaxUnavailable: the maximum number of nodes that can be not ready at the same time.
	MaxUnavailable *uint32 `json:"max_unavailable"`

	// MaxSurge: the maximum number of nodes to be created during the upgrade.
	MaxSurge *uint32 `json:"max_surge"`
}

// ClusterAutoUpgrade: cluster auto upgrade.
type ClusterAutoUpgrade struct {
	// Enabled: defines whether auto upgrade is enabled for the cluster.
	Enabled bool `json:"enabled"`

	// MaintenanceWindow: maintenance window of the cluster auto upgrades.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window"`
}

// ClusterAutoscalerConfig: cluster autoscaler config.
type ClusterAutoscalerConfig struct {
	// ScaleDownDisabled: forbid cluster autoscaler to scale down the cluster, defaults to false.
	ScaleDownDisabled bool `json:"scale_down_disabled"`

	// ScaleDownDelayAfterAdd: how long after scale up the scale down evaluation resumes.
	ScaleDownDelayAfterAdd string `json:"scale_down_delay_after_add"`

	// Estimator: type of resource estimator to be used in scale up.
	// Default value: unknown_estimator
	Estimator AutoscalerEstimator `json:"estimator"`

	// Expander: kubernetes autoscaler strategy to fit pods into nodes, see https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders for details.
	// Default value: unknown_expander
	Expander AutoscalerExpander `json:"expander"`

	// IgnoreDaemonsetsUtilization: ignore DaemonSet pods when calculating resource utilization for scaling down, defaults to false.
	IgnoreDaemonsetsUtilization bool `json:"ignore_daemonsets_utilization"`

	// BalanceSimilarNodeGroups: detect similar node groups and balance the number of nodes between them, defaults to false.
	BalanceSimilarNodeGroups bool `json:"balance_similar_node_groups"`

	// ExpendablePodsPriorityCutoff: pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they won't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
	ExpendablePodsPriorityCutoff int32 `json:"expendable_pods_priority_cutoff"`

	// ScaleDownUnneededTime: how long a node should be unneeded before it is eligible for scale down, defaults to 10 minutes.
	ScaleDownUnneededTime string `json:"scale_down_unneeded_time"`

	// ScaleDownUtilizationThreshold: node utilization level, defined as a sum of requested resources divided by allocatable capacity, below which a node can be considered for scale down.
	ScaleDownUtilizationThreshold float32 `json:"scale_down_utilization_threshold"`

	// MaxGracefulTerminationSec: maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node, defaults to 600 (10 minutes).
	MaxGracefulTerminationSec uint32 `json:"max_graceful_termination_sec"`
}

// ClusterOpenIDConnectConfig: cluster open id connect config.
type ClusterOpenIDConnectConfig struct {
	// IssuerURL: URL of the provider which allows the API server to discover public signing keys. Only URLs using the `https://` scheme are accepted. This is typically the provider's discovery URL without a path, for example "https://accounts.google.com" or "https://login.salesforce.com".
	IssuerURL string `json:"issuer_url"`

	// ClientID: a client ID that all tokens must be issued for.
	ClientID string `json:"client_id"`

	// UsernameClaim: jWT claim to use as the user name. The default is `sub`, which is expected to be the end user's unique identifier. Admins can choose other claims, such as `email` or `name`, depending on their provider. However, claims other than `email` will be prefixed with the issuer URL to prevent name collision.
	UsernameClaim string `json:"username_claim"`

	// UsernamePrefix: prefix prepended to username claims to prevent name collision (such as `system:` users). For example, the value `oidc:` will create usernames like `oidc:jane.doe`. If this flag is not provided and `username_claim` is a value other than `email`, the prefix defaults to `( Issuer URL )#` where `( Issuer URL )` is the value of `issuer_url`. The value `-` can be used to disable all prefixing.
	UsernamePrefix string `json:"username_prefix"`

	// GroupsClaim: jWT claim to use as the user's group.
	GroupsClaim []string `json:"groups_claim"`

	// GroupsPrefix: prefix prepended to group claims to prevent name collision (such as `system:` groups). For example, the value `oidc:` will create group names like `oidc:engineering` and `oidc:infra`.
	GroupsPrefix string `json:"groups_prefix"`

	// RequiredClaim: multiple key=value pairs describing a required claim in the ID token. If set, the claims are verified to be present in the ID token with a matching value.
	RequiredClaim []string `json:"required_claim"`
}

// PoolUpgradePolicy: pool upgrade policy.
type PoolUpgradePolicy struct {
	MaxUnavailable uint32 `json:"max_unavailable"`

	MaxSurge uint32 `json:"max_surge"`
}

// ACLRuleRequest: acl rule request.
type ACLRuleRequest struct {
	// IP: IP subnet to allow.
	// Precisely one of IP, ScalewayRanges must be set.
	IP *scw.IPNet `json:"ip,omitempty"`

	// ScalewayRanges: only one rule with this field set to true can be added.
	// Precisely one of IP, ScalewayRanges must be set.
	ScalewayRanges *bool `json:"scaleway_ranges,omitempty"`

	// Description: description of the ACL.
	Description string `json:"description"`
}

// ACLRule: acl rule.
type ACLRule struct {
	// ID: ID of the ACL rule.
	ID string `json:"id"`

	// IP: IP subnet to allow.
	// Precisely one of IP, ScalewayRanges must be set.
	IP *scw.IPNet `json:"ip,omitempty"`

	// ScalewayRanges: only one rule with this field set to true can be added.
	// Precisely one of IP, ScalewayRanges must be set.
	ScalewayRanges *bool `json:"scaleway_ranges,omitempty"`

	// Description: description of the ACL.
	Description string `json:"description"`
}

// CreateClusterRequestAutoUpgrade: create cluster request auto upgrade.
type CreateClusterRequestAutoUpgrade struct {
	// Enable: defines whether auto upgrade is enabled for the cluster.
	Enable bool `json:"enable"`

	// MaintenanceWindow: maintenance window of the cluster auto upgrades.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window"`
}

// CreateClusterRequestAutoscalerConfig: create cluster request autoscaler config.
type CreateClusterRequestAutoscalerConfig struct {
	// ScaleDownDisabled: forbid cluster autoscaler to scale down the cluster, defaults to false.
	ScaleDownDisabled *bool `json:"scale_down_disabled"`

	// ScaleDownDelayAfterAdd: how long after scale up the scale down evaluation resumes.
	ScaleDownDelayAfterAdd *string `json:"scale_down_delay_after_add"`

	// Estimator: type of resource estimator to be used in scale up.
	// Default value: unknown_estimator
	Estimator AutoscalerEstimator `json:"estimator"`

	// Expander: kubernetes autoscaler strategy to fit pods into nodes, see https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders for details.
	// Default value: unknown_expander
	Expander AutoscalerExpander `json:"expander"`

	// IgnoreDaemonsetsUtilization: ignore DaemonSet pods when calculating resource utilization for scaling down, defaults to false.
	IgnoreDaemonsetsUtilization *bool `json:"ignore_daemonsets_utilization"`

	// BalanceSimilarNodeGroups: detect similar node groups and balance the number of nodes between them, defaults to false.
	BalanceSimilarNodeGroups *bool `json:"balance_similar_node_groups"`

	// ExpendablePodsPriorityCutoff: pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they won't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
	ExpendablePodsPriorityCutoff *int32 `json:"expendable_pods_priority_cutoff"`

	// ScaleDownUnneededTime: how long a node should be unneeded before it is eligible for scale down, defaults to 10 minutes.
	ScaleDownUnneededTime *string `json:"scale_down_unneeded_time"`

	// ScaleDownUtilizationThreshold: node utilization level, defined as a sum of requested resources divided by allocatable capacity, below which a node can be considered for scale down.
	ScaleDownUtilizationThreshold *float32 `json:"scale_down_utilization_threshold"`

	// MaxGracefulTerminationSec: maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node, defaults to 600 (10 minutes).
	MaxGracefulTerminationSec *uint32 `json:"max_graceful_termination_sec"`
}

// CreateClusterRequestOpenIDConnectConfig: create cluster request open id connect config.
type CreateClusterRequestOpenIDConnectConfig struct {
	// IssuerURL: URL of the provider which allows the API server to discover public signing keys. Only URLs using the `https://` scheme are accepted. This is typically the provider's discovery URL without a path, for example "https://accounts.google.com" or "https://login.salesforce.com".
	IssuerURL string `json:"issuer_url"`

	// ClientID: a client ID that all tokens must be issued for.
	ClientID string `json:"client_id"`

	// UsernameClaim: jWT claim to use as the user name. The default is `sub`, which is expected to be the end user's unique identifier. Admins can choose other claims, such as `email` or `name`, depending on their provider. However, claims other than `email` will be prefixed with the issuer URL to prevent name collision.
	UsernameClaim *string `json:"username_claim"`

	// UsernamePrefix: prefix prepended to username claims to prevent name collision (such as `system:` users). For example, the value `oidc:` will create usernames like `oidc:jane.doe`. If this flag is not provided and `username_claim` is a value other than `email`, the prefix defaults to `( Issuer URL )#` where `( Issuer URL )` is the value of `issuer_url`. The value `-` can be used to disable all prefixing.
	UsernamePrefix *string `json:"username_prefix"`

	// GroupsClaim: jWT claim to use as the user's group.
	GroupsClaim *[]string `json:"groups_claim"`

	// GroupsPrefix: prefix prepended to group claims to prevent name collision (such as `system:` groups). For example, the value `oidc:` will create group names like `oidc:engineering` and `oidc:infra`.
	GroupsPrefix *string `json:"groups_prefix"`

	// RequiredClaim: multiple key=value pairs describing a required claim in the ID token. If set, the claims are verified to be present in the ID token with a matching value.
	RequiredClaim *[]string `json:"required_claim"`
}

// CreateClusterRequestPoolConfig: create cluster request pool config.
type CreateClusterRequestPoolConfig struct {
	// Name: name of the pool.
	Name string `json:"name"`

	// NodeType: node type is the type of Scaleway Instance wanted for the pool. Nodes with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). 'external' is a special node type used to provision instances from other cloud providers in a Kosmos Cluster.
	NodeType string `json:"node_type"`

	// PlacementGroupID: placement group ID in which all the nodes of the pool will be created, placement groups are limited to 20 instances.
	PlacementGroupID *string `json:"placement_group_id"`

	// Autoscaling: defines whether the autoscaling feature is enabled for the pool.
	Autoscaling bool `json:"autoscaling"`

	// Size: size (number of nodes) of the pool.
	Size uint32 `json:"size"`

	// MinSize: defines the minimum size of the pool. Note that this field is only used when autoscaling is enabled on the pool.
	MinSize *uint32 `json:"min_size"`

	// MaxSize: defines the maximum size of the pool. Note that this field is only used when autoscaling is enabled on the pool.
	MaxSize *uint32 `json:"max_size"`

	// ContainerRuntime: customization of the container runtime is available for each pool.
	// Default value: unknown_runtime
	ContainerRuntime Runtime `json:"container_runtime"`

	// Autohealing: defines whether the autohealing feature is enabled for the pool.
	Autohealing bool `json:"autohealing"`

	// Tags: tags associated with the pool, see [managing tags](https://www.scaleway.com/en/docs/kubernetes/api-cli/managing-tags).
	Tags []string `json:"tags"`

	// KubeletArgs: kubelet arguments to be used by this pool. Note that this feature is experimental.
	KubeletArgs map[string]string `json:"kubelet_args"`

	// UpgradePolicy: pool upgrade policy.
	UpgradePolicy *CreateClusterRequestPoolConfigUpgradePolicy `json:"upgrade_policy"`

	// Zone: zone in which the pool's nodes will be spawned.
	Zone scw.Zone `json:"zone"`

	// RootVolumeType: * `l_ssd` is a local block storage which means your system is stored locally on your node's hypervisor. This type is not available for all node types
	// * `sbs_5k` is a remote block storage which means your system is stored on a centralized and resilient cluster with 5k IOPS limits
	// * `sbs_15k` is a faster remote block storage which means your system is stored on a centralized and resilient cluster with 15k IOPS limits
	// * `b_ssd` is the legacy remote block storage which means your system is stored on a centralized and resilient cluster. Not available for new pools, use `sbs_5k` or `sbs_15k` instead.
	// Default value: default_volume_type
	RootVolumeType PoolVolumeType `json:"root_volume_type"`

	// RootVolumeSize: system volume disk size.
	RootVolumeSize *scw.Size `json:"root_volume_size"`

	// PublicIPDisabled: defines if the public IP should be removed from Nodes. To use this feature, your Cluster must have an attached Private Network set up with a Public Gateway.
	PublicIPDisabled bool `json:"public_ip_disabled"`

	// SecurityGroupID: security group ID in which all the nodes of the pool will be created. If unset, the pool will use default Kapsule security group in current zone.
	SecurityGroupID *string `json:"security_group_id"`
}

// CreatePoolRequestUpgradePolicy: create pool request upgrade policy.
type CreatePoolRequestUpgradePolicy struct {
	MaxUnavailable *uint32 `json:"max_unavailable"`

	MaxSurge *uint32 `json:"max_surge"`
}

// ExternalNodeCoreV1Taint: external node core v1 taint.
type ExternalNodeCoreV1Taint struct {
	Key string `json:"key"`

	Value string `json:"value"`

	Effect string `json:"effect"`
}

// ClusterType: cluster type.
type ClusterType struct {
	// Name: cluster type name.
	Name string `json:"name"`

	// Availability: cluster type availability.
	// Default value: available
	Availability ClusterTypeAvailability `json:"availability"`

	// MaxNodes: maximum number of nodes supported by the offer.
	MaxNodes uint32 `json:"max_nodes"`

	// CommitmentDelay: time period during which you can no longer switch to a lower offer.
	CommitmentDelay *scw.Duration `json:"commitment_delay"`

	// SLA: value of the Service Level Agreement of the offer.
	SLA float32 `json:"sla"`

	// Resiliency: resiliency offered by the offer.
	// Default value: unknown_resiliency
	Resiliency ClusterTypeResiliency `json:"resiliency"`

	// Memory: max RAM allowed for the control plane.
	Memory scw.Size `json:"memory"`

	// Dedicated: returns information if this offer uses dedicated resources.
	Dedicated bool `json:"dedicated"`

	// AuditLogsSupported: true if the offer allows activation of the audit log functionality. Please note that audit logs are sent to Cockpit.
	AuditLogsSupported bool `json:"audit_logs_supported"`

	// MaxEtcdSize: maximum amount of data that can be stored in etcd for the offer.
	MaxEtcdSize scw.Size `json:"max_etcd_size"`
}

// Version: version.
type Version struct {
	// Name: name of the Kubernetes version.
	Name string `json:"name"`

	// Label: label of the Kubernetes version.
	Label string `json:"label"`

	// Region: region in which this version is available.
	Region scw.Region `json:"region"`

	// AvailableCnis: supported Container Network Interface (CNI) plugins for this version.
	AvailableCnis []CNI `json:"available_cnis"`

	// AvailableContainerRuntimes: supported container runtimes for this version.
	AvailableContainerRuntimes []Runtime `json:"available_container_runtimes"`

	// AvailableFeatureGates: supported feature gates for this version.
	AvailableFeatureGates []string `json:"available_feature_gates"`

	// AvailableAdmissionPlugins: supported admission plugins for this version.
	AvailableAdmissionPlugins []string `json:"available_admission_plugins"`

	// AvailableKubeletArgs: supported kubelet arguments for this version.
	AvailableKubeletArgs map[string]string `json:"available_kubelet_args"`

	// DeprecatedAt: date from which this version will no longer be available for provisioning.
	DeprecatedAt *time.Time `json:"deprecated_at"`

	// EndOfLifeAt: date from which any remaining clusters on this version will begin to be forcibly upgraded to the next minor version.
	EndOfLifeAt *time.Time `json:"end_of_life_at"`

	// ReleasedAt: date at which this version was made available by Kapsule product.
	ReleasedAt *time.Time `json:"released_at"`
}

// Cluster: cluster.
type Cluster struct {
	// ID: cluster ID.
	ID string `json:"id"`

	// Type: cluster type.
	Type string `json:"type"`

	// Name: cluster name.
	Name string `json:"name"`

	// Status: status of the cluster.
	// Default value: unknown
	Status ClusterStatus `json:"status"`

	// Version: kubernetes version of the cluster.
	Version string `json:"version"`

	// Region: region in which the cluster is deployed.
	Region scw.Region `json:"region"`

	// OrganizationID: ID of the Organization owning the cluster.
	OrganizationID string `json:"organization_id"`

	// ProjectID: ID of the Project owning the cluster.
	ProjectID string `json:"project_id"`

	// Tags: tags associated with the cluster.
	Tags []string `json:"tags"`

	// Cni: container Network Interface (CNI) plugin running in the cluster.
	// Default value: unknown_cni
	Cni CNI `json:"cni"`

	// Description: cluster description.
	Description string `json:"description"`

	// ClusterURL: kubernetes API server URL of the cluster.
	ClusterURL string `json:"cluster_url"`

	// DNSWildcard: wildcard DNS resolving all the ready cluster nodes.
	DNSWildcard string `json:"dns_wildcard"`

	// CreatedAt: date on which the cluster was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date on which the cluster was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// AutoscalerConfig: autoscaler configuration for the cluster, see https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md for details.
	AutoscalerConfig *ClusterAutoscalerConfig `json:"autoscaler_config"`

	// AutoUpgrade: auto upgrade Kubernetes version of the cluster.
	AutoUpgrade *ClusterAutoUpgrade `json:"auto_upgrade"`

	// UpgradeAvailable: defines whether a new Kubernetes version is available.
	UpgradeAvailable bool `json:"upgrade_available"`

	// FeatureGates: list of enabled feature gates.
	FeatureGates []string `json:"feature_gates"`

	// AdmissionPlugins: list of enabled admission plugins.
	AdmissionPlugins []string `json:"admission_plugins"`

	// OpenIDConnectConfig: this configuration enables to update the OpenID Connect configuration of the Kubernetes API server.
	OpenIDConnectConfig *ClusterOpenIDConnectConfig `json:"open_id_connect_config"`

	// ApiserverCertSans: additional Subject Alternative Names for the Kubernetes API server certificate.
	ApiserverCertSans []string `json:"apiserver_cert_sans"`

	// PrivateNetworkID: private network ID for internal cluster communication.
	PrivateNetworkID *string `json:"private_network_id"`

	// CommitmentEndsAt: date on which it will be possible to switch to a smaller offer.
	CommitmentEndsAt *time.Time `json:"commitment_ends_at"`

	// Deprecated: ACLAvailable: defines whether ACL is available on the cluster.
	ACLAvailable *bool `json:"acl_available,omitempty"`

	// IamNodesGroupID: iAM group that nodes are members of (this field might be empty during early stage of cluster creation).
	IamNodesGroupID string `json:"iam_nodes_group_id"`

	// PodCidr: subnet used for the Pod CIDR.
	PodCidr scw.IPNet `json:"pod_cidr"`

	// ServiceCidr: subnet used for the Service CIDR.
	ServiceCidr scw.IPNet `json:"service_cidr"`

	// ServiceDNSIP: IP used for the DNS Service.
	ServiceDNSIP net.IP `json:"service_dns_ip"`
}

// Node: node.
type Node struct {
	// ID: node ID.
	ID string `json:"id"`

	// PoolID: pool ID of the node.
	PoolID string `json:"pool_id"`

	// ClusterID: cluster ID of the node.
	ClusterID string `json:"cluster_id"`

	// ProviderID: underlying instance ID. It is prefixed by instance type and location information (see https://pkg.go.dev/k8s.io/api/core/v1#NodeSpec.ProviderID).
	ProviderID string `json:"provider_id"`

	// Region: cluster region of the node.
	Region scw.Region `json:"region"`

	// Name: name of the node.
	Name string `json:"name"`

	// Deprecated: PublicIPV4: public IPv4 address of the node.
	PublicIPV4 *net.IP `json:"public_ip_v4,omitempty"`

	// Deprecated: PublicIPV6: public IPv6 address of the node.
	PublicIPV6 *net.IP `json:"public_ip_v6,omitempty"`

	// Deprecated: Conditions: conditions of the node. These conditions contain the Node Problem Detector conditions, as well as some in house conditions.
	Conditions *map[string]string `json:"conditions,omitempty"`

	// Status: status of the node.
	// Default value: unknown
	Status NodeStatus `json:"status"`

	// ErrorMessage: details of the error, if any occurred when managing the node.
	ErrorMessage *string `json:"error_message"`

	// CreatedAt: date on which the node was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date on which the node was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
}

// Pool: pool.
type Pool struct {
	// ID: pool ID.
	ID string `json:"id"`

	// ClusterID: cluster ID of the pool.
	ClusterID string `json:"cluster_id"`

	// CreatedAt: date on which the pool was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date on which the pool was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// Name: pool name.
	Name string `json:"name"`

	// Status: pool status.
	// Default value: unknown
	Status PoolStatus `json:"status"`

	// Version: pool version.
	Version string `json:"version"`

	// NodeType: node type is the type of Scaleway Instance wanted for the pool. Nodes with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). 'external' is a special node type used to provision instances from other cloud providers in a Kosmos Cluster.
	NodeType string `json:"node_type"`

	// Autoscaling: defines whether the autoscaling feature is enabled for the pool.
	Autoscaling bool `json:"autoscaling"`

	// Size: size (number of nodes) of the pool.
	Size uint32 `json:"size"`

	// MinSize: defines the minimum size of the pool. Note that this field is only used when autoscaling is enabled on the pool.
	MinSize uint32 `json:"min_size"`

	// MaxSize: defines the maximum size of the pool. Note that this field is only used when autoscaling is enabled on the pool.
	MaxSize uint32 `json:"max_size"`

	// ContainerRuntime: customization of the container runtime is available for each pool.
	// Default value: unknown_runtime
	ContainerRuntime Runtime `json:"container_runtime"`

	// Autohealing: defines whether the autohealing feature is enabled for the pool.
	Autohealing bool `json:"autohealing"`

	// Tags: tags associated with the pool, see [managing tags](https://www.scaleway.com/en/docs/kubernetes/api-cli/managing-tags).
	Tags []string `json:"tags"`

	// PlacementGroupID: placement group ID in which all the nodes of the pool will be created, placement groups are limited to 20 instances.
	PlacementGroupID *string `json:"placement_group_id"`

	// KubeletArgs: kubelet arguments to be used by this pool. Note that this feature is experimental.
	KubeletArgs map[string]string `json:"kubelet_args"`

	// UpgradePolicy: pool upgrade policy.
	UpgradePolicy *PoolUpgradePolicy `json:"upgrade_policy"`

	// Zone: zone in which the pool's nodes will be spawned.
	Zone scw.Zone `json:"zone"`

	// RootVolumeType: * `l_ssd` is a local block storage which means your system is stored locally on your node's hypervisor. This type is not available for all node types
	// * `sbs_5k` is a remote block storage which means your system is stored on a centralized and resilient cluster with 5k IOPS limits
	// * `sbs_15k` is a faster remote block storage which means your system is stored on a centralized and resilient cluster with 15k IOPS limits
	// * `b_ssd` is the legacy remote block storage which means your system is stored on a centralized and resilient cluster. Not available for new pools, use `sbs_5k` or `sbs_15k` instead.
	// Default value: default_volume_type
	RootVolumeType PoolVolumeType `json:"root_volume_type"`

	// RootVolumeSize: system volume disk size.
	RootVolumeSize *scw.Size `json:"root_volume_size"`

	// PublicIPDisabled: defines if the public IP should be removed from Nodes. To use this feature, your Cluster must have an attached Private Network set up with a Public Gateway.
	PublicIPDisabled bool `json:"public_ip_disabled"`

	// SecurityGroupID: security group ID in which all the nodes of the pool will be created. If unset, the pool will use default Kapsule security group in current zone.
	SecurityGroupID string `json:"security_group_id"`

	// Region: cluster region of the pool.
	Region scw.Region `json:"region"`
}

// NodeMetadataCoreV1Taint: node metadata core v1 taint.
type NodeMetadataCoreV1Taint struct {
	Key string `json:"key"`

	Value string `json:"value"`

	Effect string `json:"effect"`
}

// UpdateClusterRequestAutoUpgrade: update cluster request auto upgrade.
type UpdateClusterRequestAutoUpgrade struct {
	// Enable: defines whether auto upgrade is enabled for the cluster.
	Enable *bool `json:"enable"`

	// MaintenanceWindow: maintenance window of the cluster auto upgrades.
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window"`
}

// UpdateClusterRequestAutoscalerConfig: update cluster request autoscaler config.
type UpdateClusterRequestAutoscalerConfig struct {
	// ScaleDownDisabled: forbid cluster autoscaler to scale down the cluster, defaults to false.
	ScaleDownDisabled *bool `json:"scale_down_disabled"`

	// ScaleDownDelayAfterAdd: how long after scale up the scale down evaluation resumes.
	ScaleDownDelayAfterAdd *string `json:"scale_down_delay_after_add"`

	// Estimator: type of resource estimator to be used in scale up.
	// Default value: unknown_estimator
	Estimator AutoscalerEstimator `json:"estimator"`

	// Expander: kubernetes autoscaler strategy to fit pods into nodes, see https://github.com/kubernetes/autoscaler/blob/master/cluster-autoscaler/FAQ.md#what-are-expanders for details.
	// Default value: unknown_expander
	Expander AutoscalerExpander `json:"expander"`

	// IgnoreDaemonsetsUtilization: ignore DaemonSet pods when calculating resource utilization for scaling down, defaults to false.
	IgnoreDaemonsetsUtilization *bool `json:"ignore_daemonsets_utilization"`

	// BalanceSimilarNodeGroups: detect similar node groups and balance the number of nodes between them, defaults to false.
	BalanceSimilarNodeGroups *bool `json:"balance_similar_node_groups"`

	// ExpendablePodsPriorityCutoff: pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they won't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
	ExpendablePodsPriorityCutoff *int32 `json:"expendable_pods_priority_cutoff"`

	// ScaleDownUnneededTime: how long a node should be unneeded before it is eligible for scale down, defaults to 10 minutes.
	ScaleDownUnneededTime *string `json:"scale_down_unneeded_time"`

	// ScaleDownUtilizationThreshold: node utilization level, defined as a sum of requested resources divided by allocatable capacity, below which a node can be considered for scale down.
	ScaleDownUtilizationThreshold *float32 `json:"scale_down_utilization_threshold"`

	// MaxGracefulTerminationSec: maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node, defaults to 600 (10 minutes).
	MaxGracefulTerminationSec *uint32 `json:"max_graceful_termination_sec"`
}

// UpdateClusterRequestOpenIDConnectConfig: update cluster request open id connect config.
type UpdateClusterRequestOpenIDConnectConfig struct {
	// IssuerURL: URL of the provider which allows the API server to discover public signing keys. Only URLs using the `https://` scheme are accepted. This is typically the provider's discovery URL without a path, for example "https://accounts.google.com" or "https://login.salesforce.com".
	IssuerURL *string `json:"issuer_url"`

	// ClientID: a client ID that all tokens must be issued for.
	ClientID *string `json:"client_id"`

	// UsernameClaim: jWT claim to use as the user name. The default is `sub`, which is expected to be the end user's unique identifier. Admins can choose other claims, such as `email` or `name`, depending on their provider. However, claims other than `email` will be prefixed with the issuer URL to prevent name collision.
	UsernameClaim *string `json:"username_claim"`

	// UsernamePrefix: prefix prepended to username claims to prevent name collision (such as `system:` users). For example, the value `oidc:` will create usernames like `oidc:jane.doe`. If this flag is not provided and `username_claim` is a value other than `email`, the prefix defaults to `( Issuer URL )#` where `( Issuer URL )` is the value of `issuer_url`. The value `-` can be used to disable all prefixing.
	UsernamePrefix *string `json:"username_prefix"`

	// GroupsClaim: jWT claim to use as the user's group.
	GroupsClaim *[]string `json:"groups_claim"`

	// GroupsPrefix: prefix prepended to group claims to prevent name collision (such as `system:` groups). For example, the value `oidc:` will create group names like `oidc:engineering` and `oidc:infra`.
	GroupsPrefix *string `json:"groups_prefix"`

	// RequiredClaim: multiple key=value pairs describing a required claim in the ID token. If set, the claims are verified to be present in the ID token with a matching value.
	RequiredClaim *[]string `json:"required_claim"`
}

// UpdatePoolRequestUpgradePolicy: update pool request upgrade policy.
type UpdatePoolRequestUpgradePolicy struct {
	MaxUnavailable *uint32 `json:"max_unavailable"`

	MaxSurge *uint32 `json:"max_surge"`
}

// AddClusterACLRulesRequest: add cluster acl rules request.
type AddClusterACLRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster whose ACLs will be added.
	ClusterID string `json:"-"`

	// ACLs: aCLs to add.
	ACLs []*ACLRuleRequest `json:"acls"`
}

// AddClusterACLRulesResponse: add cluster acl rules response.
type AddClusterACLRulesResponse struct {
	// Rules: aCLs that were added.
	Rules []*ACLRule `json:"rules"`
}

// AuthExternalNodeRequest: auth external node request.
type AuthExternalNodeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PoolID: pool the node will be attached to.
	PoolID string `json:"-"`
}

// CreateClusterRequest: create cluster request.
type CreateClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Deprecated: OrganizationID: organization ID in which the cluster will be created.
	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// ProjectID: project ID in which the cluster will be created.
	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// Type: type of the cluster. See [list available cluster types](#list-available-cluster-types-for-a-cluster) for a list of valid types.
	Type string `json:"type"`

	// Name: cluster name.
	Name string `json:"name"`

	// Description: cluster description.
	Description string `json:"description"`

	// Tags: tags associated with the cluster.
	Tags []string `json:"tags"`

	// Version: kubernetes version of the cluster.
	Version string `json:"version"`

	// Cni: container Network Interface (CNI) plugin running in the cluster.
	// Default value: unknown_cni
	Cni CNI `json:"cni"`

	// Pools: pools created along with the cluster.
	Pools []*CreateClusterRequestPoolConfig `json:"pools"`

	// AutoscalerConfig: autoscaler configuration for the cluster. It allows you to set (to an extent) your preferred autoscaler configuration, which is an implementation of the cluster-autoscaler (https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler/).
	AutoscalerConfig *CreateClusterRequestAutoscalerConfig `json:"autoscaler_config,omitempty"`

	// AutoUpgrade: auto upgrade configuration of the cluster. This configuration enables to set a specific 2-hour time window in which the cluster can be automatically updated to the latest patch version.
	AutoUpgrade *CreateClusterRequestAutoUpgrade `json:"auto_upgrade,omitempty"`

	// FeatureGates: list of feature gates to enable.
	FeatureGates []string `json:"feature_gates"`

	// AdmissionPlugins: list of admission plugins to enable.
	AdmissionPlugins []string `json:"admission_plugins"`

	// OpenIDConnectConfig: openID Connect configuration of the cluster. This configuration enables to update the OpenID Connect configuration of the Kubernetes API server.
	OpenIDConnectConfig *CreateClusterRequestOpenIDConnectConfig `json:"open_id_connect_config,omitempty"`

	// ApiserverCertSans: additional Subject Alternative Names for the Kubernetes API server certificate.
	ApiserverCertSans []string `json:"apiserver_cert_sans"`

	// PrivateNetworkID: private network ID for internal cluster communication (cannot be changed later).
	PrivateNetworkID *string `json:"private_network_id,omitempty"`

	// PodCidr: subnet used for the Pod CIDR (cannot be changed later).
	PodCidr *scw.IPNet `json:"pod_cidr,omitempty"`

	// ServiceCidr: subnet used for the Service CIDR (cannot be changed later).
	ServiceCidr *scw.IPNet `json:"service_cidr,omitempty"`

	// ServiceDNSIP: IP used for the DNS Service (cannot be changes later). If unset, default to Service CIDR's network + 10.
	ServiceDNSIP *net.IP `json:"service_dns_ip,omitempty"`
}

// CreateExternalNodeRequest: create external node request.
type CreateExternalNodeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	PoolID string `json:"-"`
}

// CreatePoolRequest: create pool request.
type CreatePoolRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: cluster ID to which the pool will be attached.
	ClusterID string `json:"-"`

	// Name: pool name.
	Name string `json:"name"`

	// NodeType: node type is the type of Scaleway Instance wanted for the pool. Nodes with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). 'external' is a special node type used to provision instances from other cloud providers in a Kosmos Cluster.
	NodeType string `json:"node_type"`

	// PlacementGroupID: placement group ID in which all the nodes of the pool will be created, placement groups are limited to 20 instances.
	PlacementGroupID *string `json:"placement_group_id,omitempty"`

	// Autoscaling: defines whether the autoscaling feature is enabled for the pool.
	Autoscaling bool `json:"autoscaling"`

	// Size: size (number of nodes) of the pool.
	Size uint32 `json:"size"`

	// MinSize: defines the minimum size of the pool. Note that this field is only used when autoscaling is enabled on the pool.
	MinSize *uint32 `json:"min_size,omitempty"`

	// MaxSize: defines the maximum size of the pool. Note that this field is only used when autoscaling is enabled on the pool.
	MaxSize *uint32 `json:"max_size,omitempty"`

	// ContainerRuntime: customization of the container runtime is available for each pool.
	// Default value: unknown_runtime
	ContainerRuntime Runtime `json:"container_runtime"`

	// Autohealing: defines whether the autohealing feature is enabled for the pool.
	Autohealing bool `json:"autohealing"`

	// Tags: tags associated with the pool, see [managing tags](https://www.scaleway.com/en/docs/kubernetes/api-cli/managing-tags).
	Tags []string `json:"tags"`

	// KubeletArgs: kubelet arguments to be used by this pool. Note that this feature is experimental.
	KubeletArgs map[string]string `json:"kubelet_args"`

	// UpgradePolicy: pool upgrade policy.
	UpgradePolicy *CreatePoolRequestUpgradePolicy `json:"upgrade_policy,omitempty"`

	// Zone: zone in which the pool's nodes will be spawned.
	Zone scw.Zone `json:"zone"`

	// RootVolumeType: * `l_ssd` is a local block storage which means your system is stored locally on your node's hypervisor. This type is not available for all node types
	// * `sbs_5k` is a remote block storage which means your system is stored on a centralized and resilient cluster with 5k IOPS limits
	// * `sbs_15k` is a faster remote block storage which means your system is stored on a centralized and resilient cluster with 15k IOPS limits
	// * `b_ssd` is the legacy remote block storage which means your system is stored on a centralized and resilient cluster. Not available for new pools, use `sbs_5k` or `sbs_15k` instead.
	// Default value: default_volume_type
	RootVolumeType PoolVolumeType `json:"root_volume_type"`

	// RootVolumeSize: system volume disk size.
	RootVolumeSize *scw.Size `json:"root_volume_size,omitempty"`

	// PublicIPDisabled: defines if the public IP should be removed from Nodes. To use this feature, your Cluster must have an attached Private Network set up with a Public Gateway.
	PublicIPDisabled bool `json:"public_ip_disabled"`

	// SecurityGroupID: security group ID in which all the nodes of the pool will be created. If unset, the pool will use default Kapsule security group in current zone.
	SecurityGroupID *string `json:"security_group_id,omitempty"`
}

// DeleteACLRuleRequest: delete acl rule request.
type DeleteACLRuleRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ACLID: ID of the ACL rule to delete.
	ACLID string `json:"-"`
}

// DeleteClusterRequest: delete cluster request.
type DeleteClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster to delete.
	ClusterID string `json:"-"`

	// WithAdditionalResources: defines whether all volumes (including retain volume type), empty Private Networks and Load Balancers with a name starting with the cluster ID will also be deleted.
	WithAdditionalResources bool `json:"with_additional_resources"`
}

// DeleteNodeRequest: delete node request.
type DeleteNodeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NodeID: ID of the node to replace.
	NodeID string `json:"-"`

	// SkipDrain: skip draining node from its workload (Note: this parameter is currently inactive).
	SkipDrain bool `json:"-"`

	// Replace: add a new node after the deletion of this node.
	Replace bool `json:"-"`
}

// DeletePoolRequest: delete pool request.
type DeletePoolRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PoolID: ID of the pool to delete.
	PoolID string `json:"-"`
}

// ExternalNode: external node.
type ExternalNode struct {
	ID string `json:"id"`

	Name string `json:"name"`

	ClusterURL string `json:"cluster_url"`

	PoolVersion string `json:"pool_version"`

	ClusterCa string `json:"cluster_ca"`

	KubeToken string `json:"kube_token"`

	KubeletConfig string `json:"kubelet_config"`

	ExternalIP string `json:"external_ip"`

	ContainerdVersion string `json:"containerd_version"`

	RuncVersion string `json:"runc_version"`

	CniPluginsVersion string `json:"cni_plugins_version"`

	NodeLabels map[string]string `json:"node_labels"`

	NodeTaints []*ExternalNodeCoreV1Taint `json:"node_taints"`

	IamToken string `json:"iam_token"`
}

// ExternalNodeAuth: external node auth.
type ExternalNodeAuth struct {
	NodeSecretKey string `json:"node_secret_key"`

	MetadataURL string `json:"metadata_url"`
}

// GetClusterKubeConfigRequest: get cluster kube config request.
type GetClusterKubeConfigRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: cluster ID for which to download the kubeconfig.
	ClusterID string `json:"-"`

	// Redacted: hide the legacy token from the kubeconfig.
	Redacted *bool `json:"redacted,omitempty"`
}

// GetClusterRequest: get cluster request.
type GetClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the requested cluster.
	ClusterID string `json:"-"`
}

// GetNodeMetadataRequest: get node metadata request.
type GetNodeMetadataRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// GetNodeRequest: get node request.
type GetNodeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NodeID: ID of the requested node.
	NodeID string `json:"-"`
}

// GetPoolRequest: get pool request.
type GetPoolRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PoolID: ID of the requested pool.
	PoolID string `json:"-"`
}

// GetVersionRequest: get version request.
type GetVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VersionName: requested version name.
	VersionName string `json:"-"`
}

// ListClusterACLRulesRequest: list cluster acl rules request.
type ListClusterACLRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster whose ACLs will be listed.
	ClusterID string `json:"-"`

	// Page: page number for the returned ACLs.
	Page *int32 `json:"-"`

	// PageSize: maximum number of ACLs per page.
	PageSize *uint32 `json:"-"`
}

// ListClusterACLRulesResponse: list cluster acl rules response.
type ListClusterACLRulesResponse struct {
	// TotalCount: total number of ACLs that exist for the cluster.
	TotalCount uint64 `json:"total_count"`

	// Rules: paginated returned ACLs.
	Rules []*ACLRule `json:"rules"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClusterACLRulesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClusterACLRulesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListClusterACLRulesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Rules = append(r.Rules, results.Rules...)
	r.TotalCount += uint64(len(results.Rules))
	return uint64(len(results.Rules)), nil
}

// ListClusterAvailableTypesRequest: list cluster available types request.
type ListClusterAvailableTypesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: cluster ID for which the available Kubernetes types will be listed.
	ClusterID string `json:"-"`
}

// ListClusterAvailableTypesResponse: list cluster available types response.
type ListClusterAvailableTypesResponse struct {
	// ClusterTypes: available cluster types for the cluster.
	ClusterTypes []*ClusterType `json:"cluster_types"`

	// TotalCount: total number of types.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClusterAvailableTypesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClusterAvailableTypesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListClusterAvailableTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ClusterTypes = append(r.ClusterTypes, results.ClusterTypes...)
	r.TotalCount += uint64(len(results.ClusterTypes))
	return uint64(len(results.ClusterTypes)), nil
}

// ListClusterAvailableVersionsRequest: list cluster available versions request.
type ListClusterAvailableVersionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: cluster ID for which the available Kubernetes versions will be listed.
	ClusterID string `json:"-"`
}

// ListClusterAvailableVersionsResponse: list cluster available versions response.
type ListClusterAvailableVersionsResponse struct {
	// Versions: available Kubernetes versions for the cluster.
	Versions []*Version `json:"versions"`
}

// ListClusterTypesRequest: list cluster types request.
type ListClusterTypesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number, from the paginated results, to return for cluster-types.
	Page *int32 `json:"-"`

	// PageSize: maximum number of clusters per page.
	PageSize *uint32 `json:"-"`
}

// ListClusterTypesResponse: list cluster types response.
type ListClusterTypesResponse struct {
	// TotalCount: total number of cluster-types.
	TotalCount uint64 `json:"total_count"`

	// ClusterTypes: paginated returned cluster-types.
	ClusterTypes []*ClusterType `json:"cluster_types"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClusterTypesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClusterTypesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListClusterTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ClusterTypes = append(r.ClusterTypes, results.ClusterTypes...)
	r.TotalCount += uint64(len(results.ClusterTypes))
	return uint64(len(results.ClusterTypes)), nil
}

// ListClustersRequest: list clusters request.
type ListClustersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrganizationID: organization ID on which to filter the returned clusters.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID on which to filter the returned clusters.
	ProjectID *string `json:"-"`

	// OrderBy: sort order of returned clusters.
	// Default value: created_at_asc
	OrderBy ListClustersRequestOrderBy `json:"-"`

	// Page: page number to return for clusters, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: maximum number of clusters per page.
	PageSize *uint32 `json:"-"`

	// Name: name to filter on, only clusters containing this substring in their name will be returned.
	Name *string `json:"-"`

	// Status: status to filter on, only clusters with this status will be returned.
	// Default value: unknown
	Status ClusterStatus `json:"-"`

	// Type: type to filter on, only clusters with this type will be returned.
	Type *string `json:"-"`

	// PrivateNetworkID: private Network ID to filter on, only clusters within this Private Network will be returned.
	PrivateNetworkID *string `json:"-"`
}

// ListClustersResponse: list clusters response.
type ListClustersResponse struct {
	// TotalCount: total number of clusters.
	TotalCount uint64 `json:"total_count"`

	// Clusters: paginated returned clusters.
	Clusters []*Cluster `json:"clusters"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClustersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClustersResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListClustersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Clusters = append(r.Clusters, results.Clusters...)
	r.TotalCount += uint64(len(results.Clusters))
	return uint64(len(results.Clusters)), nil
}

// ListNodesRequest: list nodes request.
type ListNodesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: cluster ID from which the nodes will be listed from.
	ClusterID string `json:"-"`

	// PoolID: pool ID on which to filter the returned nodes.
	PoolID *string `json:"-"`

	// OrderBy: sort order of the returned nodes.
	// Default value: created_at_asc
	OrderBy ListNodesRequestOrderBy `json:"-"`

	// Page: page number for the returned nodes.
	Page *int32 `json:"-"`

	// PageSize: maximum number of nodes per page.
	PageSize *uint32 `json:"-"`

	// Name: name to filter on, only nodes containing this substring in their name will be returned.
	Name *string `json:"-"`

	// Status: status to filter on, only nodes with this status will be returned.
	// Default value: unknown
	Status NodeStatus `json:"-"`
}

// ListNodesResponse: list nodes response.
type ListNodesResponse struct {
	// TotalCount: total number of nodes.
	TotalCount uint64 `json:"total_count"`

	// Nodes: paginated returned nodes.
	Nodes []*Node `json:"nodes"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNodesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNodesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListNodesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Nodes = append(r.Nodes, results.Nodes...)
	r.TotalCount += uint64(len(results.Nodes))
	return uint64(len(results.Nodes)), nil
}

// ListPoolsRequest: list pools request.
type ListPoolsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster whose pools will be listed.
	ClusterID string `json:"-"`

	// OrderBy: sort order of returned pools.
	// Default value: created_at_asc
	OrderBy ListPoolsRequestOrderBy `json:"-"`

	// Page: page number for the returned pools.
	Page *int32 `json:"-"`

	// PageSize: maximum number of pools per page.
	PageSize *uint32 `json:"-"`

	// Name: name to filter on, only pools containing this substring in their name will be returned.
	Name *string `json:"-"`

	// Status: status to filter on, only pools with this status will be returned.
	// Default value: unknown
	Status PoolStatus `json:"-"`
}

// ListPoolsResponse: list pools response.
type ListPoolsResponse struct {
	// TotalCount: total number of pools that exists for the cluster.
	TotalCount uint64 `json:"total_count"`

	// Pools: paginated returned pools.
	Pools []*Pool `json:"pools"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPoolsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPoolsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListPoolsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Pools = append(r.Pools, results.Pools...)
	r.TotalCount += uint64(len(results.Pools))
	return uint64(len(results.Pools)), nil
}

// ListVersionsRequest: list versions request.
type ListVersionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// ListVersionsResponse: list versions response.
type ListVersionsResponse struct {
	// Versions: available Kubernetes versions.
	Versions []*Version `json:"versions"`
}

// NodeMetadata: node metadata.
type NodeMetadata struct {
	ID string `json:"id"`

	Name string `json:"name"`

	ClusterURL string `json:"cluster_url"`

	ClusterCa string `json:"cluster_ca"`

	CredentialProviderConfig string `json:"credential_provider_config"`

	PoolVersion string `json:"pool_version"`

	KubeletConfig string `json:"kubelet_config"`

	NodeLabels map[string]string `json:"node_labels"`

	NodeTaints []*NodeMetadataCoreV1Taint `json:"node_taints"`

	ProviderID string `json:"provider_id"`

	ResolvconfPath string `json:"resolvconf_path"`

	TemplateArgs map[string]string `json:"template_args"`

	HasGpu bool `json:"has_gpu"`

	ExternalIP string `json:"external_ip"`

	RepoURI string `json:"repo_uri"`

	InstallerTags []string `json:"installer_tags"`

	UpdaterBinURL string `json:"updater_bin_url"`

	UpdaterBinVersion string `json:"updater_bin_version"`

	UpdaterBinPath string `json:"updater_bin_path"`
}

// RebootNodeRequest: reboot node request.
type RebootNodeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NodeID: ID of the node to reboot.
	NodeID string `json:"-"`
}

// ReplaceNodeRequest: replace node request.
type ReplaceNodeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NodeID: ID of the node to replace.
	NodeID string `json:"-"`
}

// ResetClusterAdminTokenRequest: reset cluster admin token request.
type ResetClusterAdminTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: cluster ID on which the admin token will be renewed.
	ClusterID string `json:"-"`
}

// SetClusterACLRulesRequest: set cluster acl rules request.
type SetClusterACLRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster whose ACLs will be set.
	ClusterID string `json:"-"`

	// ACLs: aCLs to set.
	ACLs []*ACLRuleRequest `json:"acls"`
}

// SetClusterACLRulesResponse: set cluster acl rules response.
type SetClusterACLRulesResponse struct {
	// Rules: aCLs that were set.
	Rules []*ACLRule `json:"rules"`
}

// SetClusterTypeRequest: set cluster type request.
type SetClusterTypeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster to migrate from one type to another.
	ClusterID string `json:"-"`

	// Type: type of the cluster. Note that some migrations are not possible (please refer to product documentation).
	Type string `json:"type"`
}

// UpdateClusterRequest: update cluster request.
type UpdateClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster to update.
	ClusterID string `json:"-"`

	// Name: new external name for the cluster.
	Name *string `json:"name,omitempty"`

	// Description: new description for the cluster.
	Description *string `json:"description,omitempty"`

	// Tags: new tags associated with the cluster.
	Tags *[]string `json:"tags,omitempty"`

	// AutoscalerConfig: new autoscaler config for the cluster.
	AutoscalerConfig *UpdateClusterRequestAutoscalerConfig `json:"autoscaler_config,omitempty"`

	// AutoUpgrade: new auto upgrade configuration for the cluster. Note that all fields needs to be set.
	AutoUpgrade *UpdateClusterRequestAutoUpgrade `json:"auto_upgrade,omitempty"`

	// FeatureGates: list of feature gates to enable.
	FeatureGates *[]string `json:"feature_gates,omitempty"`

	// AdmissionPlugins: list of admission plugins to enable.
	AdmissionPlugins *[]string `json:"admission_plugins,omitempty"`

	// OpenIDConnectConfig: openID Connect configuration of the cluster. This configuration enables to update the OpenID Connect configuration of the Kubernetes API server.
	OpenIDConnectConfig *UpdateClusterRequestOpenIDConnectConfig `json:"open_id_connect_config,omitempty"`

	// ApiserverCertSans: additional Subject Alternative Names for the Kubernetes API server certificate.
	ApiserverCertSans *[]string `json:"apiserver_cert_sans,omitempty"`
}

// UpdatePoolRequest: update pool request.
type UpdatePoolRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PoolID: ID of the pool to update.
	PoolID string `json:"-"`

	// Autoscaling: new value for the pool autoscaling enablement.
	Autoscaling *bool `json:"autoscaling,omitempty"`

	// Size: new desired pool size.
	Size *uint32 `json:"size,omitempty"`

	// MinSize: new minimum size for the pool.
	MinSize *uint32 `json:"min_size,omitempty"`

	// MaxSize: new maximum size for the pool.
	MaxSize *uint32 `json:"max_size,omitempty"`

	// Autohealing: new value for the pool autohealing enablement.
	Autohealing *bool `json:"autohealing,omitempty"`

	// Tags: new tags associated with the pool.
	Tags *[]string `json:"tags,omitempty"`

	// KubeletArgs: new Kubelet arguments to be used by this pool. Note that this feature is experimental.
	KubeletArgs *map[string]string `json:"kubelet_args,omitempty"`

	// UpgradePolicy: new upgrade policy for the pool.
	UpgradePolicy *UpdatePoolRequestUpgradePolicy `json:"upgrade_policy,omitempty"`

	// SecurityGroupID: security group ID in which all the nodes of the pool will be moved.
	SecurityGroupID *string `json:"security_group_id,omitempty"`
}

// UpgradeClusterRequest: upgrade cluster request.
type UpgradeClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster to upgrade.
	ClusterID string `json:"-"`

	// Version: new Kubernetes version of the cluster. Note that the version should either be a higher patch version of the same minor version or the direct minor version after the current one.
	Version string `json:"version"`

	// UpgradePools: defines whether pools will also be upgraded once the control plane is upgraded.
	UpgradePools bool `json:"upgrade_pools"`
}

// UpgradePoolRequest: upgrade pool request.
type UpgradePoolRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PoolID: ID of the pool to upgrade.
	PoolID string `json:"-"`

	// Version: new Kubernetes version for the pool.
	Version string `json:"version"`
}

// This API allows you to manage Kubernetes Kapsule and Kosmos clusters.
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

// ListClusters: List all existing Kubernetes clusters in a specific region.
func (s *API) ListClusters(req *ListClustersRequest, opts ...scw.RequestOption) (*ListClustersResponse, error) {
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "private_network_id", req.PrivateNetworkID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters",
		Query:  query,
	}

	var resp ListClustersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCluster: Create a new Kubernetes cluster in a Scaleway region.
func (s *API) CreateCluster(req *CreateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("k8s")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCluster: Retrieve information about a specific Kubernetes cluster.
func (s *API) GetCluster(req *GetClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForClusterRequest is used by WaitForCluster method.
type WaitForClusterRequest struct {
	Region        scw.Region
	ClusterID     string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForCluster waits for the Cluster to reach a terminal state.
func (s *API) WaitForCluster(req *WaitForClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	timeout := defaultK8sTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultK8sRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[ClusterStatus]struct{}{
		ClusterStatusCreating: {},
		ClusterStatusDeleting: {},
		ClusterStatusUpdating: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetCluster(&GetClusterRequest{
				Region:    req.Region,
				ClusterID: req.ClusterID,
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
		return nil, errors.Wrap(err, "waiting for Cluster failed")
	}

	return res.(*Cluster), nil
}

// UpdateCluster: Update information on a specific Kubernetes cluster. You can update details such as its name, description, tags and configuration. To upgrade a cluster, you will need to use the dedicated endpoint.
func (s *API) UpdateCluster(req *UpdateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCluster: Delete a specific Kubernetes cluster and all its associated pools and nodes, and possibly its associated Load Balancers or Block Volumes.
func (s *API) DeleteCluster(req *DeleteClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "with_additional_resources", req.WithAdditionalResources)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
		Query:  query,
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpgradeCluster: Upgrade a specific Kubernetes cluster and possibly its associated pools to a specific and supported Kubernetes version.
func (s *API) UpgradeCluster(req *UpgradeClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/upgrade",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetClusterType: Change the type of a specific Kubernetes cluster. To see the possible values you can enter for the `type` field, [list available cluster types](#list-available-cluster-types-for-a-cluster).
func (s *API) SetClusterType(req *SetClusterTypeRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/set-type",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterAvailableVersions: List the versions that a specific Kubernetes cluster is allowed to upgrade to. Results will include every patch version greater than the current patch, as well as one minor version ahead of the current version. Any upgrade skipping a minor version will not work.
func (s *API) ListClusterAvailableVersions(req *ListClusterAvailableVersionsRequest, opts ...scw.RequestOption) (*ListClusterAvailableVersionsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/available-versions",
	}

	var resp ListClusterAvailableVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterAvailableTypes: List the cluster types that a specific Kubernetes cluster is allowed to switch to.
func (s *API) ListClusterAvailableTypes(req *ListClusterAvailableTypesRequest, opts ...scw.RequestOption) (*ListClusterAvailableTypesResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/available-types",
	}

	var resp ListClusterAvailableTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// getClusterKubeConfig: Download the Kubernetes cluster config file (also known as `kubeconfig`) for a specific cluster to use it with `kubectl`.
// Tip: add `?dl=1` at the end of the URL to directly retrieve the base64 decoded kubeconfig. If you choose not to, the kubeconfig will be base64 encoded.
func (s *API) getClusterKubeConfig(req *GetClusterKubeConfigRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "redacted", req.Redacted)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/kubeconfig",
		Query:  query,
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResetClusterAdminToken: Reset the admin token for a specific Kubernetes cluster. This will revoke the old admin token (which will not be usable afterwards) and create a new one. Note that you will need to download the kubeconfig again to keep interacting with the cluster.
func (s *API) ResetClusterAdminToken(req *ResetClusterAdminTokenRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/reset-admin-token",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListClusterACLRules: List ACLs for a specific cluster.
func (s *API) ListClusterACLRules(req *ListClusterACLRulesRequest, opts ...scw.RequestOption) (*ListClusterACLRulesResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/acls",
		Query:  query,
	}

	var resp ListClusterACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddClusterACLRules: Add new ACL rules for a specific cluster.
func (s *API) AddClusterACLRules(req *AddClusterACLRulesRequest, opts ...scw.RequestOption) (*AddClusterACLRulesResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddClusterACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetClusterACLRules: Set new ACL rules for a specific cluster.
func (s *API) SetClusterACLRules(req *SetClusterACLRulesRequest, opts ...scw.RequestOption) (*SetClusterACLRulesResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetClusterACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteACLRule: Delete an existing ACL.
func (s *API) DeleteACLRule(req *DeleteACLRuleRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/acls/" + fmt.Sprint(req.ACLID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListPools: List all the existing pools for a specific Kubernetes cluster.
func (s *API) ListPools(req *ListPoolsRequest, opts ...scw.RequestOption) (*ListPoolsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/pools",
		Query:  query,
	}

	var resp ListPoolsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePool: Create a new pool in a specific Kubernetes cluster.
func (s *API) CreatePool(req *CreatePoolRequest, opts ...scw.RequestOption) (*Pool, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("pool")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/pools",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Pool

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPool: Retrieve details about a specific pool in a Kubernetes cluster.
func (s *API) GetPool(req *GetPoolRequest, opts ...scw.RequestOption) (*Pool, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PoolID) == "" {
		return nil, errors.New("field PoolID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "",
	}

	var resp Pool

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForPoolRequest is used by WaitForPool method.
type WaitForPoolRequest struct {
	Region        scw.Region
	PoolID        string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForPool waits for the Pool to reach a terminal state.
func (s *API) WaitForPool(req *WaitForPoolRequest, opts ...scw.RequestOption) (*Pool, error) {
	timeout := defaultK8sTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultK8sRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[PoolStatus]struct{}{
		PoolStatusDeleting:  {},
		PoolStatusScaling:   {},
		PoolStatusUpgrading: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetPool(&GetPoolRequest{
				Region: req.Region,
				PoolID: req.PoolID,
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
		return nil, errors.Wrap(err, "waiting for Pool failed")
	}

	return res.(*Pool), nil
}

// UpgradePool: Upgrade the Kubernetes version of a specific pool. Note that it only works if the targeted version matches the cluster's version.
// This will drain and replace the nodes in that pool.
func (s *API) UpgradePool(req *UpgradePoolRequest, opts ...scw.RequestOption) (*Pool, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PoolID) == "" {
		return nil, errors.New("field PoolID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "/upgrade",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Pool

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePool: Update the attributes of a specific pool, such as its desired size, autoscaling settings, and tags. To upgrade a pool, you will need to use the dedicated endpoint.
func (s *API) UpdatePool(req *UpdatePoolRequest, opts ...scw.RequestOption) (*Pool, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PoolID) == "" {
		return nil, errors.New("field PoolID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Pool

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePool: Delete a specific pool from a cluster. Note that all the pool's nodes will also be deleted.
func (s *API) DeletePool(req *DeletePoolRequest, opts ...scw.RequestOption) (*Pool, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PoolID) == "" {
		return nil, errors.New("field PoolID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "",
	}

	var resp Pool

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNodeMetadata: Rerieve metadata to instantiate a Kapsule/Kosmos node. This method is not intended to be called by end users but rather programmatically by the node-installer.
func (s *API) GetNodeMetadata(req *GetNodeMetadataRequest, opts ...scw.RequestOption) (*NodeMetadata, error) {
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
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/node-metadata",
	}

	var resp NodeMetadata

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AuthExternalNode: Creates a newer Kosmos node and returns its token. This method is not intended to be called by end users but rather programmatically by the node-installer.
func (s *API) AuthExternalNode(req *AuthExternalNodeRequest, opts ...scw.RequestOption) (*ExternalNodeAuth, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PoolID) == "" {
		return nil, errors.New("field PoolID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "/external-nodes/auth",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ExternalNodeAuth

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateExternalNode: Retrieve metadata for a Kosmos node. This method is not intended to be called by end users but rather programmatically by the kapsule-node-agent.
func (s *API) CreateExternalNode(req *CreateExternalNodeRequest, opts ...scw.RequestOption) (*ExternalNode, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PoolID) == "" {
		return nil, errors.New("field PoolID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "/external-nodes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ExternalNode

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodes: List all the existing nodes for a specific Kubernetes cluster.
func (s *API) ListNodes(req *ListNodesRequest, opts ...scw.RequestOption) (*ListNodesResponse, error) {
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
	parameter.AddToQuery(query, "pool_id", req.PoolID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/nodes",
		Query:  query,
	}

	var resp ListNodesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNode: Retrieve details about a specific Kubernetes Node.
func (s *API) GetNode(req *GetNodeRequest, opts ...scw.RequestOption) (*Node, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NodeID) == "" {
		return nil, errors.New("field NodeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "",
	}

	var resp Node

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForNodeRequest is used by WaitForNode method.
type WaitForNodeRequest struct {
	Region        scw.Region
	NodeID        string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForNode waits for the Node to reach a terminal state.
func (s *API) WaitForNode(req *WaitForNodeRequest, opts ...scw.RequestOption) (*Node, error) {
	timeout := defaultK8sTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultK8sRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[NodeStatus]struct{}{
		NodeStatusCreating:    {},
		NodeStatusDeleting:    {},
		NodeStatusRebooting:   {},
		NodeStatusUpgrading:   {},
		NodeStatusStarting:    {},
		NodeStatusRegistering: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetNode(&GetNodeRequest{
				Region: req.Region,
				NodeID: req.NodeID,
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
		return nil, errors.Wrap(err, "waiting for Node failed")
	}

	return res.(*Node), nil
}

// Deprecated: ReplaceNode: Replace a specific Node. The node will first be drained and pods will be rescheduled onto another node. Note that when there is not enough space to reschedule all the pods (such as in a one-node cluster, or with specific constraints), disruption of your applications may occur.
func (s *API) ReplaceNode(req *ReplaceNodeRequest, opts ...scw.RequestOption) (*Node, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NodeID) == "" {
		return nil, errors.New("field NodeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "/replace",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Node

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RebootNode: Reboot a specific Node. The node will first be drained and pods will be rescheduled onto another node. Note that when there is not enough space to reschedule all the pods (such as in a one-node cluster, or with specific constraints), disruption of your applications may occur.
func (s *API) RebootNode(req *RebootNodeRequest, opts ...scw.RequestOption) (*Node, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NodeID) == "" {
		return nil, errors.New("field NodeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "/reboot",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Node

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNode: Delete a specific Node. The node will first be drained and pods will be rescheduled onto another node. Note that when there is not enough space to reschedule all the pods (such as in a one-node cluster, or with specific constraints), disruption of your applications may occur.
func (s *API) DeleteNode(req *DeleteNodeRequest, opts ...scw.RequestOption) (*Node, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "skip_drain", req.SkipDrain)
	parameter.AddToQuery(query, "replace", req.Replace)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NodeID) == "" {
		return nil, errors.New("field NodeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "",
		Query:  query,
	}

	var resp Node

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVersions: List all available versions for the creation of a new Kubernetes cluster.
func (s *API) ListVersions(req *ListVersionsRequest, opts ...scw.RequestOption) (*ListVersionsResponse, error) {
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
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/versions",
	}

	var resp ListVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVersion: Retrieve a specific Kubernetes version and its details.
func (s *API) GetVersion(req *GetVersionRequest, opts ...scw.RequestOption) (*Version, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VersionName) == "" {
		return nil, errors.New("field VersionName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/versions/" + fmt.Sprint(req.VersionName) + "",
	}

	var resp Version

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterTypes: List available cluster types and their technical details.
func (s *API) ListClusterTypes(req *ListClusterTypesRequest, opts ...scw.RequestOption) (*ListClusterTypesResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/cluster-types",
		Query:  query,
	}

	var resp ListClusterTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

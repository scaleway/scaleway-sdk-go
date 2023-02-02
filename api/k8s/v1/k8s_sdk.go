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

// API: kapsule API
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type AutoscalerEstimator string

const (
	// AutoscalerEstimatorUnknownEstimator is [insert doc].
	AutoscalerEstimatorUnknownEstimator = AutoscalerEstimator("unknown_estimator")
	// AutoscalerEstimatorBinpacking is [insert doc].
	AutoscalerEstimatorBinpacking = AutoscalerEstimator("binpacking")
)

func (enum AutoscalerEstimator) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_estimator"
	}
	return string(enum)
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

type AutoscalerExpander string

const (
	// AutoscalerExpanderUnknownExpander is [insert doc].
	AutoscalerExpanderUnknownExpander = AutoscalerExpander("unknown_expander")
	// AutoscalerExpanderRandom is [insert doc].
	AutoscalerExpanderRandom = AutoscalerExpander("random")
	// AutoscalerExpanderMostPods is [insert doc].
	AutoscalerExpanderMostPods = AutoscalerExpander("most_pods")
	// AutoscalerExpanderLeastWaste is [insert doc].
	AutoscalerExpanderLeastWaste = AutoscalerExpander("least_waste")
	// AutoscalerExpanderPriority is [insert doc].
	AutoscalerExpanderPriority = AutoscalerExpander("priority")
	// AutoscalerExpanderPrice is [insert doc].
	AutoscalerExpanderPrice = AutoscalerExpander("price")
)

func (enum AutoscalerExpander) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_expander"
	}
	return string(enum)
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
	// CNIUnknownCni is [insert doc].
	CNIUnknownCni = CNI("unknown_cni")
	// CNICilium is [insert doc].
	CNICilium = CNI("cilium")
	// CNICalico is [insert doc].
	CNICalico = CNI("calico")
	// CNIWeave is [insert doc].
	CNIWeave = CNI("weave")
	// CNIFlannel is [insert doc].
	CNIFlannel = CNI("flannel")
	// CNIKilo is [insert doc].
	CNIKilo = CNI("kilo")
)

func (enum CNI) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_cni"
	}
	return string(enum)
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
	// ClusterStatusUnknown is [insert doc].
	ClusterStatusUnknown = ClusterStatus("unknown")
	// ClusterStatusCreating is [insert doc].
	ClusterStatusCreating = ClusterStatus("creating")
	// ClusterStatusReady is [insert doc].
	ClusterStatusReady = ClusterStatus("ready")
	// ClusterStatusDeleting is [insert doc].
	ClusterStatusDeleting = ClusterStatus("deleting")
	// ClusterStatusDeleted is [insert doc].
	ClusterStatusDeleted = ClusterStatus("deleted")
	// ClusterStatusUpdating is [insert doc].
	ClusterStatusUpdating = ClusterStatus("updating")
	// ClusterStatusLocked is [insert doc].
	ClusterStatusLocked = ClusterStatus("locked")
	// ClusterStatusPoolRequired is [insert doc].
	ClusterStatusPoolRequired = ClusterStatus("pool_required")
)

func (enum ClusterStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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

type Ingress string

const (
	// IngressUnknownIngress is [insert doc].
	IngressUnknownIngress = Ingress("unknown_ingress")
	// IngressNone is [insert doc].
	IngressNone = Ingress("none")
	// IngressNginx is [insert doc].
	IngressNginx = Ingress("nginx")
	// IngressTraefik is [insert doc].
	IngressTraefik = Ingress("traefik")
	// IngressTraefik2 is [insert doc].
	IngressTraefik2 = Ingress("traefik2")
)

func (enum Ingress) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_ingress"
	}
	return string(enum)
}

func (enum Ingress) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Ingress) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Ingress(Ingress(tmp).String())
	return nil
}

type ListClustersRequestOrderBy string

const (
	// ListClustersRequestOrderByCreatedAtAsc is [insert doc].
	ListClustersRequestOrderByCreatedAtAsc = ListClustersRequestOrderBy("created_at_asc")
	// ListClustersRequestOrderByCreatedAtDesc is [insert doc].
	ListClustersRequestOrderByCreatedAtDesc = ListClustersRequestOrderBy("created_at_desc")
	// ListClustersRequestOrderByUpdatedAtAsc is [insert doc].
	ListClustersRequestOrderByUpdatedAtAsc = ListClustersRequestOrderBy("updated_at_asc")
	// ListClustersRequestOrderByUpdatedAtDesc is [insert doc].
	ListClustersRequestOrderByUpdatedAtDesc = ListClustersRequestOrderBy("updated_at_desc")
	// ListClustersRequestOrderByNameAsc is [insert doc].
	ListClustersRequestOrderByNameAsc = ListClustersRequestOrderBy("name_asc")
	// ListClustersRequestOrderByNameDesc is [insert doc].
	ListClustersRequestOrderByNameDesc = ListClustersRequestOrderBy("name_desc")
	// ListClustersRequestOrderByStatusAsc is [insert doc].
	ListClustersRequestOrderByStatusAsc = ListClustersRequestOrderBy("status_asc")
	// ListClustersRequestOrderByStatusDesc is [insert doc].
	ListClustersRequestOrderByStatusDesc = ListClustersRequestOrderBy("status_desc")
	// ListClustersRequestOrderByVersionAsc is [insert doc].
	ListClustersRequestOrderByVersionAsc = ListClustersRequestOrderBy("version_asc")
	// ListClustersRequestOrderByVersionDesc is [insert doc].
	ListClustersRequestOrderByVersionDesc = ListClustersRequestOrderBy("version_desc")
)

func (enum ListClustersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
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
	// ListNodesRequestOrderByCreatedAtAsc is [insert doc].
	ListNodesRequestOrderByCreatedAtAsc = ListNodesRequestOrderBy("created_at_asc")
	// ListNodesRequestOrderByCreatedAtDesc is [insert doc].
	ListNodesRequestOrderByCreatedAtDesc = ListNodesRequestOrderBy("created_at_desc")
)

func (enum ListNodesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
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
	// ListPoolsRequestOrderByCreatedAtAsc is [insert doc].
	ListPoolsRequestOrderByCreatedAtAsc = ListPoolsRequestOrderBy("created_at_asc")
	// ListPoolsRequestOrderByCreatedAtDesc is [insert doc].
	ListPoolsRequestOrderByCreatedAtDesc = ListPoolsRequestOrderBy("created_at_desc")
	// ListPoolsRequestOrderByUpdatedAtAsc is [insert doc].
	ListPoolsRequestOrderByUpdatedAtAsc = ListPoolsRequestOrderBy("updated_at_asc")
	// ListPoolsRequestOrderByUpdatedAtDesc is [insert doc].
	ListPoolsRequestOrderByUpdatedAtDesc = ListPoolsRequestOrderBy("updated_at_desc")
	// ListPoolsRequestOrderByNameAsc is [insert doc].
	ListPoolsRequestOrderByNameAsc = ListPoolsRequestOrderBy("name_asc")
	// ListPoolsRequestOrderByNameDesc is [insert doc].
	ListPoolsRequestOrderByNameDesc = ListPoolsRequestOrderBy("name_desc")
	// ListPoolsRequestOrderByStatusAsc is [insert doc].
	ListPoolsRequestOrderByStatusAsc = ListPoolsRequestOrderBy("status_asc")
	// ListPoolsRequestOrderByStatusDesc is [insert doc].
	ListPoolsRequestOrderByStatusDesc = ListPoolsRequestOrderBy("status_desc")
	// ListPoolsRequestOrderByVersionAsc is [insert doc].
	ListPoolsRequestOrderByVersionAsc = ListPoolsRequestOrderBy("version_asc")
	// ListPoolsRequestOrderByVersionDesc is [insert doc].
	ListPoolsRequestOrderByVersionDesc = ListPoolsRequestOrderBy("version_desc")
)

func (enum ListPoolsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
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
	// MaintenanceWindowDayOfTheWeekAny is [insert doc].
	MaintenanceWindowDayOfTheWeekAny = MaintenanceWindowDayOfTheWeek("any")
	// MaintenanceWindowDayOfTheWeekMonday is [insert doc].
	MaintenanceWindowDayOfTheWeekMonday = MaintenanceWindowDayOfTheWeek("monday")
	// MaintenanceWindowDayOfTheWeekTuesday is [insert doc].
	MaintenanceWindowDayOfTheWeekTuesday = MaintenanceWindowDayOfTheWeek("tuesday")
	// MaintenanceWindowDayOfTheWeekWednesday is [insert doc].
	MaintenanceWindowDayOfTheWeekWednesday = MaintenanceWindowDayOfTheWeek("wednesday")
	// MaintenanceWindowDayOfTheWeekThursday is [insert doc].
	MaintenanceWindowDayOfTheWeekThursday = MaintenanceWindowDayOfTheWeek("thursday")
	// MaintenanceWindowDayOfTheWeekFriday is [insert doc].
	MaintenanceWindowDayOfTheWeekFriday = MaintenanceWindowDayOfTheWeek("friday")
	// MaintenanceWindowDayOfTheWeekSaturday is [insert doc].
	MaintenanceWindowDayOfTheWeekSaturday = MaintenanceWindowDayOfTheWeek("saturday")
	// MaintenanceWindowDayOfTheWeekSunday is [insert doc].
	MaintenanceWindowDayOfTheWeekSunday = MaintenanceWindowDayOfTheWeek("sunday")
)

func (enum MaintenanceWindowDayOfTheWeek) String() string {
	if enum == "" {
		// return default value if empty
		return "any"
	}
	return string(enum)
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
	// NodeStatusUnknown is [insert doc].
	NodeStatusUnknown = NodeStatus("unknown")
	// NodeStatusCreating is [insert doc].
	NodeStatusCreating = NodeStatus("creating")
	// NodeStatusNotReady is [insert doc].
	NodeStatusNotReady = NodeStatus("not_ready")
	// NodeStatusReady is [insert doc].
	NodeStatusReady = NodeStatus("ready")
	// NodeStatusDeleting is [insert doc].
	NodeStatusDeleting = NodeStatus("deleting")
	// NodeStatusDeleted is [insert doc].
	NodeStatusDeleted = NodeStatus("deleted")
	// NodeStatusLocked is [insert doc].
	NodeStatusLocked = NodeStatus("locked")
	// NodeStatusRebooting is [insert doc].
	NodeStatusRebooting = NodeStatus("rebooting")
	// NodeStatusCreationError is [insert doc].
	NodeStatusCreationError = NodeStatus("creation_error")
	// NodeStatusUpgrading is [insert doc].
	NodeStatusUpgrading = NodeStatus("upgrading")
	// NodeStatusStarting is [insert doc].
	NodeStatusStarting = NodeStatus("starting")
	// NodeStatusRegistering is [insert doc].
	NodeStatusRegistering = NodeStatus("registering")
)

func (enum NodeStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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
	// PoolStatusUnknown is [insert doc].
	PoolStatusUnknown = PoolStatus("unknown")
	// PoolStatusReady is [insert doc].
	PoolStatusReady = PoolStatus("ready")
	// PoolStatusDeleting is [insert doc].
	PoolStatusDeleting = PoolStatus("deleting")
	// PoolStatusDeleted is [insert doc].
	PoolStatusDeleted = PoolStatus("deleted")
	// PoolStatusScaling is [insert doc].
	PoolStatusScaling = PoolStatus("scaling")
	// PoolStatusWarning is [insert doc].
	PoolStatusWarning = PoolStatus("warning")
	// PoolStatusLocked is [insert doc].
	PoolStatusLocked = PoolStatus("locked")
	// PoolStatusUpgrading is [insert doc].
	PoolStatusUpgrading = PoolStatus("upgrading")
)

func (enum PoolStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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
	// PoolVolumeTypeDefaultVolumeType is [insert doc].
	PoolVolumeTypeDefaultVolumeType = PoolVolumeType("default_volume_type")
	// PoolVolumeTypeLSSD is [insert doc].
	PoolVolumeTypeLSSD = PoolVolumeType("l_ssd")
	// PoolVolumeTypeBSSD is [insert doc].
	PoolVolumeTypeBSSD = PoolVolumeType("b_ssd")
)

func (enum PoolVolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return "default_volume_type"
	}
	return string(enum)
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
	// RuntimeUnknownRuntime is [insert doc].
	RuntimeUnknownRuntime = Runtime("unknown_runtime")
	// RuntimeDocker is [insert doc].
	RuntimeDocker = Runtime("docker")
	// RuntimeContainerd is [insert doc].
	RuntimeContainerd = Runtime("containerd")
	// RuntimeCrio is [insert doc].
	RuntimeCrio = Runtime("crio")
)

func (enum Runtime) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_runtime"
	}
	return string(enum)
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

// Cluster: cluster
type Cluster struct {
	// ID: the ID of the cluster
	ID string `json:"id"`
	// Type: the type of the cluster
	Type string `json:"type"`
	// Name: the name of the cluster
	Name string `json:"name"`
	// Status: the status of the cluster
	//
	// Default value: unknown
	Status ClusterStatus `json:"status"`
	// Version: the Kubernetes version of the cluster
	Version string `json:"version"`
	// Region: the region in which the cluster is
	Region scw.Region `json:"region"`
	// OrganizationID: the ID of the organization owning the cluster
	OrganizationID string `json:"organization_id"`
	// ProjectID: the ID of the project owning the cluster
	ProjectID string `json:"project_id"`
	// Tags: the tags associated with the cluster
	Tags []string `json:"tags"`
	// Cni: the Container Network Interface (CNI) plugin running in the cluster
	//
	// Default value: unknown_cni
	Cni CNI `json:"cni"`
	// Description: the description of the cluster
	Description string `json:"description"`
	// ClusterURL: the Kubernetes API server URL of the cluster
	ClusterURL string `json:"cluster_url"`
	// DNSWildcard: the DNS wildcard resovling all the ready nodes of the cluster
	DNSWildcard string `json:"dns_wildcard"`
	// CreatedAt: the date at which the cluster was created
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: the date at which the cluster was last updated
	UpdatedAt *time.Time `json:"updated_at"`
	// AutoscalerConfig: the autoscaler config for the cluster
	AutoscalerConfig *ClusterAutoscalerConfig `json:"autoscaler_config"`
	// Deprecated: DashboardEnabled: the enablement of the Kubernetes Dashboard in the cluster
	DashboardEnabled *bool `json:"dashboard_enabled,omitempty"`
	// Deprecated: Ingress: the ingress controller used in the cluster
	//
	// Default value: unknown_ingress
	Ingress *Ingress `json:"ingress,omitempty"`
	// AutoUpgrade: the auto upgrade configuration of the cluster
	AutoUpgrade *ClusterAutoUpgrade `json:"auto_upgrade"`
	// UpgradeAvailable: true if a new Kubernetes version is available
	UpgradeAvailable bool `json:"upgrade_available"`
	// FeatureGates: list of enabled feature gates
	FeatureGates []string `json:"feature_gates"`
	// AdmissionPlugins: list of enabled admission plugins
	AdmissionPlugins []string `json:"admission_plugins"`
	// OpenIDConnectConfig: aLPHA - The OpenID Connect configuration of the cluster
	//
	// This feature is in ALPHA state, it may be deleted or modified. This configuration is the [OpenID Connect configuration](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens) of the Kubernetes API server.
	OpenIDConnectConfig *ClusterOpenIDConnectConfig `json:"open_id_connect_config"`
	// ApiserverCertSans: additional Subject Alternative Names for the Kubernetes API server certificate
	ApiserverCertSans []string `json:"apiserver_cert_sans"`
}

// ClusterAutoUpgrade: cluster. auto upgrade
type ClusterAutoUpgrade struct {
	// Enabled: whether or not auto upgrade is enabled for the cluster
	Enabled bool `json:"enabled"`
	// MaintenanceWindow: the maintenance window of the cluster auto upgrades
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window"`
}

// ClusterAutoscalerConfig: cluster. autoscaler config
type ClusterAutoscalerConfig struct {
	// ScaleDownDisabled: disable the cluster autoscaler
	ScaleDownDisabled bool `json:"scale_down_disabled"`
	// ScaleDownDelayAfterAdd: how long after scale up that scale down evaluation resumes
	ScaleDownDelayAfterAdd string `json:"scale_down_delay_after_add"`
	// Estimator: type of resource estimator to be used in scale up
	//
	// Default value: unknown_estimator
	Estimator AutoscalerEstimator `json:"estimator"`
	// Expander: type of node group expander to be used in scale up
	//
	// Default value: unknown_expander
	Expander AutoscalerExpander `json:"expander"`
	// IgnoreDaemonsetsUtilization: ignore DaemonSet pods when calculating resource utilization for scaling down
	IgnoreDaemonsetsUtilization bool `json:"ignore_daemonsets_utilization"`
	// BalanceSimilarNodeGroups: detect similar node groups and balance the number of nodes between them
	BalanceSimilarNodeGroups bool `json:"balance_similar_node_groups"`
	// ExpendablePodsPriorityCutoff: pods with priority below cutoff will be expendable
	//
	// Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they don't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
	ExpendablePodsPriorityCutoff int32 `json:"expendable_pods_priority_cutoff"`
	// ScaleDownUnneededTime: how long a node should be unneeded before it is eligible for scale down
	ScaleDownUnneededTime string `json:"scale_down_unneeded_time"`
	// ScaleDownUtilizationThreshold: node utilization level, defined as sum of requested resources divided by capacity, below which a node can be considered for scale down
	ScaleDownUtilizationThreshold float32 `json:"scale_down_utilization_threshold"`
	// MaxGracefulTerminationSec: maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node
	MaxGracefulTerminationSec uint32 `json:"max_graceful_termination_sec"`
}

// ClusterOpenIDConnectConfig: cluster. open id connect config
type ClusterOpenIDConnectConfig struct {
	// IssuerURL: URL of the provider which allows the API server to discover public signing keys
	//
	// URL of the provider which allows the API server to discover public signing keys. Only URLs which use the `https://` scheme are accepted. This is typically the provider's discovery URL without a path, for example "https://accounts.google.com" or "https://login.salesforce.com". This URL should point to the level below .well-known/openid-configuration.
	//
	IssuerURL string `json:"issuer_url"`
	// ClientID: a client id that all tokens must be issued for
	ClientID string `json:"client_id"`
	// UsernameClaim: jWT claim to use as the user name
	//
	// JWT claim to use as the user name. By default `sub`, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as `email` or `name`, depending on their provider. However, claims other than `email` will be prefixed with the issuer URL to prevent naming clashes with other plugins.
	//
	UsernameClaim string `json:"username_claim"`
	// UsernamePrefix: prefix prepended to username
	//
	// Prefix prepended to username claims to prevent clashes with existing names (such as `system:` users). For example, the value `oidc:` will create usernames like `oidc:jane.doe`. If this flag isn't provided and `username_claim` is a value other than `email` the prefix defaults to `( Issuer URL )#` where `( Issuer URL )` is the value of `issuer_url`. The value `-` can be used to disable all prefixing.
	//
	UsernamePrefix string `json:"username_prefix"`
	// GroupsClaim: jWT claim to use as the user's group
	GroupsClaim []string `json:"groups_claim"`
	// GroupsPrefix: prefix prepended to group claims
	//
	// Prefix prepended to group claims to prevent clashes with existing names (such as `system:` groups). For example, the value `oidc:` will create group names like `oidc:engineering` and `oidc:infra`.
	//
	GroupsPrefix string `json:"groups_prefix"`
	// RequiredClaim: multiple key=value pairs that describes a required claim in the ID Token
	//
	// Multiple key=value pairs that describes a required claim in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value.
	//
	RequiredClaim []string `json:"required_claim"`
}

// CreateClusterRequestAutoUpgrade: create cluster request. auto upgrade
type CreateClusterRequestAutoUpgrade struct {
	// Enable: whether or not auto upgrade is enabled for the cluster
	Enable bool `json:"enable"`
	// MaintenanceWindow: the maintenance window of the cluster auto upgrades
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window"`
}

// CreateClusterRequestAutoscalerConfig: create cluster request. autoscaler config
type CreateClusterRequestAutoscalerConfig struct {
	// ScaleDownDisabled: disable the cluster autoscaler
	ScaleDownDisabled *bool `json:"scale_down_disabled"`
	// ScaleDownDelayAfterAdd: how long after scale up that scale down evaluation resumes
	ScaleDownDelayAfterAdd *string `json:"scale_down_delay_after_add"`
	// Estimator: type of resource estimator to be used in scale up
	//
	// Default value: unknown_estimator
	Estimator AutoscalerEstimator `json:"estimator"`
	// Expander: type of node group expander to be used in scale up
	//
	// Default value: unknown_expander
	Expander AutoscalerExpander `json:"expander"`
	// IgnoreDaemonsetsUtilization: ignore DaemonSet pods when calculating resource utilization for scaling down
	IgnoreDaemonsetsUtilization *bool `json:"ignore_daemonsets_utilization"`
	// BalanceSimilarNodeGroups: detect similar node groups and balance the number of nodes between them
	BalanceSimilarNodeGroups *bool `json:"balance_similar_node_groups"`
	// ExpendablePodsPriorityCutoff: pods with priority below cutoff will be expendable
	//
	// Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they don't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
	ExpendablePodsPriorityCutoff *int32 `json:"expendable_pods_priority_cutoff"`
	// ScaleDownUnneededTime: how long a node should be unneeded before it is eligible for scale down
	ScaleDownUnneededTime *string `json:"scale_down_unneeded_time"`
	// ScaleDownUtilizationThreshold: node utilization level, defined as sum of requested resources divided by capacity, below which a node can be considered for scale down
	ScaleDownUtilizationThreshold *float32 `json:"scale_down_utilization_threshold"`
	// MaxGracefulTerminationSec: maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node
	MaxGracefulTerminationSec *uint32 `json:"max_graceful_termination_sec"`
}

// CreateClusterRequestOpenIDConnectConfig: create cluster request. open id connect config
type CreateClusterRequestOpenIDConnectConfig struct {
	// IssuerURL: URL of the provider which allows the API server to discover public signing keys
	//
	// URL of the provider which allows the API server to discover public signing keys. Only URLs which use the `https://` scheme are accepted. This is typically the provider's discovery URL without a path, for example "https://accounts.google.com" or "https://login.salesforce.com". This URL should point to the level below .well-known/openid-configuration.
	//
	IssuerURL string `json:"issuer_url"`
	// ClientID: a client id that all tokens must be issued for
	ClientID string `json:"client_id"`
	// UsernameClaim: jWT claim to use as the user name
	//
	// JWT claim to use as the user name. By default `sub`, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as `email` or `name`, depending on their provider. However, claims other than `email` will be prefixed with the issuer URL to prevent naming clashes with other plugins.
	//
	UsernameClaim *string `json:"username_claim"`
	// UsernamePrefix: prefix prepended to username
	//
	// Prefix prepended to username claims to prevent clashes with existing names (such as `system:` users). For example, the value `oidc:` will create usernames like `oidc:jane.doe`. If this flag isn't provided and `username_claim` is a value other than `email` the prefix defaults to `( Issuer URL )#` where `( Issuer URL )` is the value of `issuer_url`. The value `-` can be used to disable all prefixing.
	//
	UsernamePrefix *string `json:"username_prefix"`
	// GroupsClaim: jWT claim to use as the user's group
	GroupsClaim *[]string `json:"groups_claim"`
	// GroupsPrefix: prefix prepended to group claims
	//
	// Prefix prepended to group claims to prevent clashes with existing names (such as `system:` groups). For example, the value `oidc:` will create group names like `oidc:engineering` and `oidc:infra`.
	//
	GroupsPrefix *string `json:"groups_prefix"`
	// RequiredClaim: multiple key=value pairs that describes a required claim in the ID Token
	//
	// Multiple key=value pairs that describes a required claim in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value.
	//
	RequiredClaim *[]string `json:"required_claim"`
}

// CreateClusterRequestPoolConfig: create cluster request. pool config
type CreateClusterRequestPoolConfig struct {
	// Name: the name of the pool
	Name string `json:"name"`
	// NodeType: the node type is the type of Scaleway Instance wanted for the pool
	//
	// The node type is the type of Scaleway Instance wanted for the pool. Nodes with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). 'external' is a special node type used to provision instances from other cloud providers.
	NodeType string `json:"node_type"`
	// PlacementGroupID: the placement group ID in which all the nodes of the pool will be created
	PlacementGroupID *string `json:"placement_group_id"`
	// Autoscaling: the enablement of the autoscaling feature for the pool
	Autoscaling bool `json:"autoscaling"`
	// Size: the size (number of nodes) of the pool
	Size uint32 `json:"size"`
	// MinSize: the minimum size of the pool
	//
	// The minimum size of the pool. Note that this field will be used only when autoscaling is enabled.
	MinSize *uint32 `json:"min_size"`
	// MaxSize: the maximum size of the pool
	//
	// The maximum size of the pool. Note that this field will be used only when autoscaling is enabled.
	MaxSize *uint32 `json:"max_size"`
	// ContainerRuntime: the container runtime for the nodes of the pool
	//
	// The customization of the container runtime is available for each pool. Note that `docker` is deprecated since 1.20 and will be removed in 1.24.
	//
	// Default value: unknown_runtime
	ContainerRuntime Runtime `json:"container_runtime"`
	// Autohealing: the enablement of the autohealing feature for the pool
	Autohealing bool `json:"autohealing"`
	// Tags: the tags associated with the pool
	Tags []string `json:"tags"`
	// KubeletArgs: the Kubelet arguments to be used by this pool. Note that this feature is to be considered as experimental
	KubeletArgs map[string]string `json:"kubelet_args"`
	// UpgradePolicy: the Pool upgrade policy
	UpgradePolicy *CreateClusterRequestPoolConfigUpgradePolicy `json:"upgrade_policy"`
	// Zone: the Zone in which the Pool's node will be spawn in
	Zone scw.Zone `json:"zone"`
	// RootVolumeType: the system volume disk type
	//
	// The system volume disk type, we provide two different types of volume (`volume_type`):
	//   - `l_ssd` is a local block storage: your system is stored locally on
	//     the hypervisor of your node.
	//   - `b_ssd` is a remote block storage: your system is stored on a
	//     centralised and resilient cluster.
	//
	// Default value: default_volume_type
	RootVolumeType PoolVolumeType `json:"root_volume_type"`
	// RootVolumeSize: the system volume disk size
	RootVolumeSize *scw.Size `json:"root_volume_size"`
}

// CreateClusterRequestPoolConfigUpgradePolicy: create cluster request. pool config. upgrade policy
type CreateClusterRequestPoolConfigUpgradePolicy struct {
	// MaxUnavailable: the maximum number of nodes that can be not ready at the same time
	MaxUnavailable *uint32 `json:"max_unavailable"`
	// MaxSurge: the maximum number of nodes to be created during the upgrade
	MaxSurge *uint32 `json:"max_surge"`
}

type CreatePoolRequestUpgradePolicy struct {
	MaxUnavailable *uint32 `json:"max_unavailable"`

	MaxSurge *uint32 `json:"max_surge"`
}

// ListClusterAvailableVersionsResponse: list cluster available versions response
type ListClusterAvailableVersionsResponse struct {
	// Versions: the available Kubernetes version for the cluster
	Versions []*Version `json:"versions"`
}

// ListClustersResponse: list clusters response
type ListClustersResponse struct {
	// TotalCount: the total number of clusters
	TotalCount uint32 `json:"total_count"`
	// Clusters: the paginated returned clusters
	Clusters []*Cluster `json:"clusters"`
}

// ListNodesResponse: list nodes response
type ListNodesResponse struct {
	// TotalCount: the total number of nodes
	TotalCount uint32 `json:"total_count"`
	// Nodes: the paginated returned nodes
	Nodes []*Node `json:"nodes"`
}

// ListPoolsResponse: list pools response
type ListPoolsResponse struct {
	// TotalCount: the total number of pools that exists for the cluster
	TotalCount uint32 `json:"total_count"`
	// Pools: the paginated returned pools
	Pools []*Pool `json:"pools"`
}

// ListVersionsResponse: list versions response
type ListVersionsResponse struct {
	// Versions: the available Kubernetes versions
	Versions []*Version `json:"versions"`
}

// MaintenanceWindow: maintenance window
type MaintenanceWindow struct {
	// StartHour: the start hour of the 2-hour maintenance window
	StartHour uint32 `json:"start_hour"`
	// Day: the day of the week for the maintenance window
	//
	// Default value: any
	Day MaintenanceWindowDayOfTheWeek `json:"day"`
}

// Node: node
type Node struct {
	// ID: the ID of the node
	ID string `json:"id"`
	// PoolID: the pool ID of the node
	PoolID string `json:"pool_id"`
	// ClusterID: the cluster ID of the node
	ClusterID string `json:"cluster_id"`
	// ProviderID: the underlying instance ID
	//
	// It is prefixed by instance type and location information (see https://pkg.go.dev/k8s.io/api/core/v1#NodeSpec.ProviderID).
	ProviderID string `json:"provider_id"`
	// Region: the cluster region of the node
	Region scw.Region `json:"region"`
	// Name: the name of the node
	Name string `json:"name"`
	// Deprecated: PublicIPV4: the public IPv4 address of the node
	PublicIPV4 *net.IP `json:"public_ip_v4,omitempty"`
	// Deprecated: PublicIPV6: the public IPv6 address of the node
	PublicIPV6 *net.IP `json:"public_ip_v6,omitempty"`
	// Deprecated: Conditions: the conditions of the node
	//
	// These conditions contains the Node Problem Detector conditions, as well as some in house conditions.
	Conditions *map[string]string `json:"conditions,omitempty"`
	// Status: the status of the node
	//
	// Default value: unknown
	Status NodeStatus `json:"status"`
	// ErrorMessage: details of the error, if any occured when managing the node
	ErrorMessage *string `json:"error_message"`
	// CreatedAt: the date at which the node was created
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: the date at which the node was last updated
	UpdatedAt *time.Time `json:"updated_at"`
}

// Pool: pool
type Pool struct {
	// ID: the ID of the pool
	ID string `json:"id"`
	// ClusterID: the cluster ID of the pool
	ClusterID string `json:"cluster_id"`
	// CreatedAt: the date at which the pool was created
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: the date at which the pool was last updated
	UpdatedAt *time.Time `json:"updated_at"`
	// Name: the name of the pool
	Name string `json:"name"`
	// Status: the status of the pool
	//
	// Default value: unknown
	Status PoolStatus `json:"status"`
	// Version: the version of the pool
	Version string `json:"version"`
	// NodeType: the node type is the type of Scaleway Instance wanted for the pool
	//
	// The node type is the type of Scaleway Instance wanted for the pool. Nodes with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). 'external' is a special node type used to provision instances from other cloud providers.
	NodeType string `json:"node_type"`
	// Autoscaling: the enablement of the autoscaling feature for the pool
	Autoscaling bool `json:"autoscaling"`
	// Size: the size (number of nodes) of the pool
	Size uint32 `json:"size"`
	// MinSize: the minimum size of the pool
	//
	// The minimum size of the pool. Note that this field will be used only when autoscaling is enabled.
	MinSize uint32 `json:"min_size"`
	// MaxSize: the maximum size of the pool
	//
	// The maximum size of the pool. Note that this field will be used only when autoscaling is enabled.
	MaxSize uint32 `json:"max_size"`
	// ContainerRuntime: the container runtime for the nodes of the pool
	//
	// The customization of the container runtime is available for each pool. Note that `docker` is deprecated since 1.20 and will be removed in 1.24.
	//
	// Default value: unknown_runtime
	ContainerRuntime Runtime `json:"container_runtime"`
	// Autohealing: the enablement of the autohealing feature for the pool
	Autohealing bool `json:"autohealing"`
	// Tags: the tags associated with the pool
	Tags []string `json:"tags"`
	// PlacementGroupID: the placement group ID in which all the nodes of the pool will be created
	PlacementGroupID *string `json:"placement_group_id"`
	// KubeletArgs: the Kubelet arguments to be used by this pool. Note that this feature is to be considered as experimental
	KubeletArgs map[string]string `json:"kubelet_args"`
	// UpgradePolicy: the Pool upgrade policy
	UpgradePolicy *PoolUpgradePolicy `json:"upgrade_policy"`
	// Zone: the Zone in which the Pool's node will be spawn in
	Zone scw.Zone `json:"zone"`
	// RootVolumeType: the system volume disk type
	//
	// The system volume disk type, we provide two different types of volume (`volume_type`):
	//   - `l_ssd` is a local block storage: your system is stored locally on
	//     the hypervisor of your node.
	//   - `b_ssd` is a remote block storage: your system is stored on a
	//     centralised and resilient cluster.
	//
	// Default value: default_volume_type
	RootVolumeType PoolVolumeType `json:"root_volume_type"`
	// RootVolumeSize: the system volume disk size
	RootVolumeSize *scw.Size `json:"root_volume_size"`
	// Region: the cluster region of the pool
	Region scw.Region `json:"region"`
}

type PoolUpgradePolicy struct {
	MaxUnavailable uint32 `json:"max_unavailable"`

	MaxSurge uint32 `json:"max_surge"`
}

// UpdateClusterRequestAutoUpgrade: update cluster request. auto upgrade
type UpdateClusterRequestAutoUpgrade struct {
	// Enable: whether or not auto upgrade is enabled for the cluster
	Enable *bool `json:"enable"`
	// MaintenanceWindow: the maintenance window of the cluster auto upgrades
	MaintenanceWindow *MaintenanceWindow `json:"maintenance_window"`
}

// UpdateClusterRequestAutoscalerConfig: update cluster request. autoscaler config
type UpdateClusterRequestAutoscalerConfig struct {
	// ScaleDownDisabled: disable the cluster autoscaler
	ScaleDownDisabled *bool `json:"scale_down_disabled"`
	// ScaleDownDelayAfterAdd: how long after scale up that scale down evaluation resumes
	ScaleDownDelayAfterAdd *string `json:"scale_down_delay_after_add"`
	// Estimator: type of resource estimator to be used in scale up
	//
	// Default value: unknown_estimator
	Estimator AutoscalerEstimator `json:"estimator"`
	// Expander: type of node group expander to be used in scale up
	//
	// Default value: unknown_expander
	Expander AutoscalerExpander `json:"expander"`
	// IgnoreDaemonsetsUtilization: ignore DaemonSet pods when calculating resource utilization for scaling down
	IgnoreDaemonsetsUtilization *bool `json:"ignore_daemonsets_utilization"`
	// BalanceSimilarNodeGroups: detect similar node groups and balance the number of nodes between them
	BalanceSimilarNodeGroups *bool `json:"balance_similar_node_groups"`
	// ExpendablePodsPriorityCutoff: pods with priority below cutoff will be expendable
	//
	// Pods with priority below cutoff will be expendable. They can be killed without any consideration during scale down and they don't cause scale up. Pods with null priority (PodPriority disabled) are non expendable.
	ExpendablePodsPriorityCutoff *int32 `json:"expendable_pods_priority_cutoff"`
	// ScaleDownUnneededTime: how long a node should be unneeded before it is eligible for scale down
	ScaleDownUnneededTime *string `json:"scale_down_unneeded_time"`
	// ScaleDownUtilizationThreshold: node utilization level, defined as sum of requested resources divided by capacity, below which a node can be considered for scale down
	ScaleDownUtilizationThreshold *float32 `json:"scale_down_utilization_threshold"`
	// MaxGracefulTerminationSec: maximum number of seconds the cluster autoscaler waits for pod termination when trying to scale down a node
	MaxGracefulTerminationSec *uint32 `json:"max_graceful_termination_sec"`
}

// UpdateClusterRequestOpenIDConnectConfig: update cluster request. open id connect config
type UpdateClusterRequestOpenIDConnectConfig struct {
	// IssuerURL: URL of the provider which allows the API server to discover public signing keys
	//
	// URL of the provider which allows the API server to discover public signing keys. Only URLs which use the `https://` scheme are accepted. This is typically the provider's discovery URL without a path, for example "https://accounts.google.com" or "https://login.salesforce.com". This URL should point to the level below .well-known/openid-configuration.
	//
	IssuerURL *string `json:"issuer_url"`
	// ClientID: a client id that all tokens must be issued for
	ClientID *string `json:"client_id"`
	// UsernameClaim: jWT claim to use as the user name
	//
	// JWT claim to use as the user name. By default `sub`, which is expected to be a unique identifier of the end user. Admins can choose other claims, such as `email` or `name`, depending on their provider. However, claims other than `email` will be prefixed with the issuer URL to prevent naming clashes with other plugins.
	//
	UsernameClaim *string `json:"username_claim"`
	// UsernamePrefix: prefix prepended to username
	//
	// Prefix prepended to username claims to prevent clashes with existing names (such as `system:` users). For example, the value `oidc:` will create usernames like `oidc:jane.doe`. If this flag isn't provided and `username_claim` is a value other than `email` the prefix defaults to `( Issuer URL )#` where `( Issuer URL )` is the value of `issuer_url`. The value `-` can be used to disable all prefixing.
	//
	UsernamePrefix *string `json:"username_prefix"`
	// GroupsClaim: jWT claim to use as the user's group
	GroupsClaim *[]string `json:"groups_claim"`
	// GroupsPrefix: prefix prepended to group claims
	//
	// Prefix prepended to group claims to prevent clashes with existing names (such as `system:` groups). For example, the value `oidc:` will create group names like `oidc:engineering` and `oidc:infra`.
	//
	GroupsPrefix *string `json:"groups_prefix"`
	// RequiredClaim: multiple key=value pairs that describes a required claim in the ID Token
	//
	// Multiple key=value pairs that describes a required claim in the ID Token. If set, the claims are verified to be present in the ID Token with a matching value.
	//
	RequiredClaim *[]string `json:"required_claim"`
}

type UpdatePoolRequestUpgradePolicy struct {
	MaxUnavailable *uint32 `json:"max_unavailable"`

	MaxSurge *uint32 `json:"max_surge"`
}

// Version: version
type Version struct {
	// Name: the name of the Kubernetes version
	Name string `json:"name"`
	// Label: the label of the Kubernetes version
	Label string `json:"label"`
	// Region: the region in which this version is available
	Region scw.Region `json:"region"`
	// AvailableCnis: the supported Container Network Interface (CNI) plugins for this version
	AvailableCnis []CNI `json:"available_cnis"`
	// Deprecated: AvailableIngresses: the supported Ingress Controllers for this version
	AvailableIngresses *[]Ingress `json:"available_ingresses,omitempty"`
	// AvailableContainerRuntimes: the supported container runtimes for this version
	AvailableContainerRuntimes []Runtime `json:"available_container_runtimes"`
	// AvailableFeatureGates: the supported feature gates for this version
	AvailableFeatureGates []string `json:"available_feature_gates"`
	// AvailableAdmissionPlugins: the supported admission plugins for this version
	AvailableAdmissionPlugins []string `json:"available_admission_plugins"`
	// AvailableKubeletArgs: the supported kubelet arguments for this version
	AvailableKubeletArgs map[string]string `json:"available_kubelet_args"`
}

// Service API

// Regions list localities the api is available in
func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

type ListClustersRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// OrganizationID: the organization ID on which to filter the returned clusters
	OrganizationID *string `json:"-"`
	// ProjectID: the project ID on which to filter the returned clusters
	ProjectID *string `json:"-"`
	// OrderBy: the sort order of the returned clusters
	//
	// Default value: created_at_asc
	OrderBy ListClustersRequestOrderBy `json:"-"`
	// Page: the page number for the returned clusters
	Page *int32 `json:"-"`
	// PageSize: the maximum number of clusters per page
	PageSize *uint32 `json:"-"`
	// Name: the name on which to filter the returned clusters
	Name *string `json:"-"`
	// Status: the status on which to filter the returned clusters
	//
	// Default value: unknown
	Status ClusterStatus `json:"-"`
	// Type: the type on which to filter the returned clusters
	Type *string `json:"-"`
}

// ListClusters: list all the clusters
//
// This method allows to list all the existing Kubernetes clusters in an account.
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListClustersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateClusterRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// Deprecated: OrganizationID: the organization ID where the cluster will be created
	// Precisely one of OrganizationID, ProjectID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: the project ID where the cluster will be created
	// Precisely one of OrganizationID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`
	// Type: the type of the cluster
	//
	// The type of the cluster (possible values are kapsule, multicloud).
	Type string `json:"type"`
	// Name: the name of the cluster
	Name string `json:"name"`
	// Description: the description of the cluster
	Description string `json:"description"`
	// Tags: the tags associated with the cluster
	Tags []string `json:"tags"`
	// Version: the Kubernetes version of the cluster
	Version string `json:"version"`
	// Cni: the Container Network Interface (CNI) plugin that will run in the cluster
	//
	// Default value: unknown_cni
	Cni CNI `json:"cni"`
	// Deprecated: EnableDashboard: the enablement of the Kubernetes Dashboard in the cluster
	EnableDashboard *bool `json:"enable_dashboard,omitempty"`
	// Deprecated: Ingress: the Ingress Controller that will run in the cluster
	//
	// Default value: unknown_ingress
	Ingress *Ingress `json:"ingress,omitempty"`
	// Pools: the pools to be created along with the cluster
	Pools []*CreateClusterRequestPoolConfig `json:"pools"`
	// AutoscalerConfig: the autoscaler config for the cluster
	//
	// This field allows to specify some configuration for the autoscaler, which is an implementation of the [cluster-autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler/).
	AutoscalerConfig *CreateClusterRequestAutoscalerConfig `json:"autoscaler_config"`
	// AutoUpgrade: the auto upgrade configuration of the cluster
	//
	// This configuration enables to set a specific 2-hour time window in which the cluster can be automatically updated to the latest patch version in the current minor one.
	AutoUpgrade *CreateClusterRequestAutoUpgrade `json:"auto_upgrade"`
	// FeatureGates: list of feature gates to enable
	FeatureGates []string `json:"feature_gates"`
	// AdmissionPlugins: list of admission plugins to enable
	AdmissionPlugins []string `json:"admission_plugins"`
	// OpenIDConnectConfig: aLPHA - The OpenID Connect configuration of the cluster
	//
	// This feature is in ALPHA state, it may be deleted or modified. This configuration enables to set the [OpenID Connect configuration](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens) of the Kubernetes API server.
	OpenIDConnectConfig *CreateClusterRequestOpenIDConnectConfig `json:"open_id_connect_config"`
	// ApiserverCertSans: additional Subject Alternative Names for the Kubernetes API server certificate
	ApiserverCertSans []string `json:"apiserver_cert_sans"`
}

// CreateCluster: create a new cluster
//
// This method allows to create a new Kubernetes cluster on an account.
func (s *API) CreateCluster(req *CreateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("k8s")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters",
		Headers: http.Header{},
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

type GetClusterRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// ClusterID: the ID of the requested cluster
	ClusterID string `json:"-"`
}

// GetCluster: get a cluster
//
// This method allows to get details about a specific Kubernetes cluster.
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
		Method:  "GET",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
		Headers: http.Header{},
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateClusterRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// ClusterID: the ID of the cluster to update
	ClusterID string `json:"-"`
	// Name: the new name of the cluster
	//
	// This field allows to update the external name of the cluster. The internal name (used for instance in hostname) won't change.
	Name *string `json:"name"`
	// Description: the new description of the cluster
	Description *string `json:"description"`
	// Tags: the new tags associated with the cluster
	Tags *[]string `json:"tags"`
	// AutoscalerConfig: the new autoscaler config for the cluster
	//
	// This field allows to update some configuration for the autoscaler, which is an implementation of the [cluster-autoscaler](https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler/).
	AutoscalerConfig *UpdateClusterRequestAutoscalerConfig `json:"autoscaler_config"`
	// Deprecated: EnableDashboard: the new value of the Kubernetes Dashboard enablement
	EnableDashboard *bool `json:"enable_dashboard,omitempty"`
	// Deprecated: Ingress: the new Ingress Controller for the cluster
	//
	// Default value: unknown_ingress
	Ingress *Ingress `json:"ingress,omitempty"`
	// AutoUpgrade: the new auto upgrade configuration of the cluster
	//
	// The new auto upgrade configuration of the cluster. Note that all fields need to be set.
	AutoUpgrade *UpdateClusterRequestAutoUpgrade `json:"auto_upgrade"`
	// FeatureGates: list of feature gates to enable
	FeatureGates *[]string `json:"feature_gates"`
	// AdmissionPlugins: list of admission plugins to enable
	AdmissionPlugins *[]string `json:"admission_plugins"`
	// OpenIDConnectConfig: aLPHA - The new OpenID Connect configuration of the cluster
	//
	// This feature is in ALPHA state, it may be deleted or modified. This configuration enables to update the [OpenID Connect configuration](https://kubernetes.io/docs/reference/access-authn-authz/authentication/#openid-connect-tokens) of the Kubernetes API server.
	OpenIDConnectConfig *UpdateClusterRequestOpenIDConnectConfig `json:"open_id_connect_config"`
	// ApiserverCertSans: additional Subject Alternative Names for the Kubernetes API server certificate
	ApiserverCertSans *[]string `json:"apiserver_cert_sans"`
}

// UpdateCluster: update a cluster
//
// This method allows to update a specific Kubernetes cluster. Note that this method is not made to upgrade a Kubernetes cluster.
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
		Method:  "PATCH",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
		Headers: http.Header{},
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

type DeleteClusterRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// ClusterID: the ID of the cluster to delete
	ClusterID string `json:"-"`
	// WithAdditionalResources: set true if you want to delete all volumes (including retain volume type) and loadbalancers whose name start with cluster ID
	WithAdditionalResources bool `json:"-"`
}

// DeleteCluster: delete a cluster
//
// This method allows to delete a specific cluster and all its associated pools and nodes. Note that this method will not delete any Load Balancers or Block Volumes that are associated with the cluster.
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
		Method:  "DELETE",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
		Query:   query,
		Headers: http.Header{},
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpgradeClusterRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// ClusterID: the ID of the cluster to upgrade
	ClusterID string `json:"-"`
	// Version: the new Kubernetes version of the cluster
	//
	// The new Kubernetes version of the cluster. Note that the version shoud either be a higher patch version of the same minor version or the direct minor version after the current one.
	Version string `json:"version"`
	// UpgradePools: the enablement of the pools upgrade
	//
	// This field makes the upgrade upgrades the pool once the Kubernetes master in upgrade.
	UpgradePools bool `json:"upgrade_pools"`
}

// UpgradeCluster: upgrade a cluster
//
// This method allows to upgrade a specific Kubernetes cluster and/or its associated pools to a specific and supported Kubernetes version.
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
		Method:  "POST",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/upgrade",
		Headers: http.Header{},
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

type ListClusterAvailableVersionsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// ClusterID: the ID of the cluster which the available Kuberentes versions will be listed from
	ClusterID string `json:"-"`
}

// ListClusterAvailableVersions: list available versions for a cluster
//
// This method allows to list the versions that a specific Kubernetes cluster is allowed to upgrade to. Note that it will be every patch version greater than the actual one as well a one minor version ahead of the actual one. Upgrades skipping a minor version will not work.
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
		Method:  "GET",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/available-versions",
		Headers: http.Header{},
	}

	var resp ListClusterAvailableVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetClusterKubeConfigRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// ClusterID: the ID of the cluster to download the kubeconfig from
	ClusterID string `json:"-"`
}

// getClusterKubeConfig: download the kubeconfig for a cluster
//
// This method allows to download the Kubernetes cluster config file (AKA kubeconfig) for a specific cluster in order to use it with, for instance, `kubectl`. Tips: add `?dl=1` at the end of the URL to directly get the base64 decoded kubeconfig. If not, the kubeconfig will be base64 encoded.
//
func (s *API) getClusterKubeConfig(req *GetClusterKubeConfigRequest, opts ...scw.RequestOption) (*scw.File, error) {
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
		Method:  "GET",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/kubeconfig",
		Headers: http.Header{},
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ResetClusterAdminTokenRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// ClusterID: the ID of the cluster of which the admin token will be renewed
	ClusterID string `json:"-"`
}

// ResetClusterAdminToken: reset the admin token of a cluster
//
// This method allows to reset the admin token for a specific Kubernetes cluster. This will invalidate the old admin token (which will not be usable after) and create a new one. Note that the redownload of the kubeconfig will be necessary to keep interacting with the cluster (if the old admin token was used).
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
		Method:  "POST",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/reset-admin-token",
		Headers: http.Header{},
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

type ListPoolsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// ClusterID: the ID of the cluster from which the pools will be listed from
	ClusterID string `json:"-"`
	// OrderBy: the sort order of the returned pools
	//
	// Default value: created_at_asc
	OrderBy ListPoolsRequestOrderBy `json:"-"`
	// Page: the page number for the returned pools
	Page *int32 `json:"-"`
	// PageSize: the maximum number of pools per page
	PageSize *uint32 `json:"-"`
	// Name: the name on which to filter the returned pools
	Name *string `json:"-"`
	// Status: the status on which to filter the returned pools
	//
	// Default value: unknown
	Status PoolStatus `json:"-"`
}

// ListPools: list all the pools in a cluster
//
// This method allows to list all the existing pools for a specific Kubernetes cluster.
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
		Method:  "GET",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/pools",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListPoolsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreatePoolRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// ClusterID: the ID of the cluster in which the pool will be created
	ClusterID string `json:"-"`
	// Name: the name of the pool
	Name string `json:"name"`
	// NodeType: the node type is the type of Scaleway Instance wanted for the pool
	//
	// The node type is the type of Scaleway Instance wanted for the pool. Nodes with insufficient memory are not eligible (DEV1-S, PLAY2-PICO, STARDUST). 'external' is a special node type used to provision instances from other cloud providers.
	NodeType string `json:"node_type"`
	// PlacementGroupID: the placement group ID in which all the nodes of the pool will be created
	PlacementGroupID *string `json:"placement_group_id"`
	// Autoscaling: the enablement of the autoscaling feature for the pool
	Autoscaling bool `json:"autoscaling"`
	// Size: the size (number of nodes) of the pool
	Size uint32 `json:"size"`
	// MinSize: the minimum size of the pool
	//
	// The minimum size of the pool. Note that this field will be used only when autoscaling is enabled.
	MinSize *uint32 `json:"min_size"`
	// MaxSize: the maximum size of the pool
	//
	// The maximum size of the pool. Note that this field will be used only when autoscaling is enabled.
	MaxSize *uint32 `json:"max_size"`
	// ContainerRuntime: the container runtime for the nodes of the pool
	//
	// The customization of the container runtime is available for each pool. Note that `docker` is deprecated since 1.20 and will be removed in 1.24.
	//
	// Default value: unknown_runtime
	ContainerRuntime Runtime `json:"container_runtime"`
	// Autohealing: the enablement of the autohealing feature for the pool
	Autohealing bool `json:"autohealing"`
	// Tags: the tags associated with the pool
	Tags []string `json:"tags"`
	// KubeletArgs: the Kubelet arguments to be used by this pool. Note that this feature is to be considered as experimental
	KubeletArgs map[string]string `json:"kubelet_args"`
	// UpgradePolicy: the Pool upgrade policy
	UpgradePolicy *CreatePoolRequestUpgradePolicy `json:"upgrade_policy"`
	// Zone: the Zone in which the Pool's node will be spawn in
	Zone scw.Zone `json:"zone"`
	// RootVolumeType: the system volume disk type
	//
	// The system volume disk type, we provide two different types of volume (`volume_type`):
	//   - `l_ssd` is a local block storage: your system is stored locally on
	//     the hypervisor of your node.
	//   - `b_ssd` is a remote block storage: your system is stored on a
	//     centralised and resilient cluster.
	//
	// Default value: default_volume_type
	RootVolumeType PoolVolumeType `json:"root_volume_type"`
	// RootVolumeSize: the system volume disk size
	RootVolumeSize *scw.Size `json:"root_volume_size"`
}

// CreatePool: create a new pool in a cluster
//
// This method allows to create a new pool in a specific Kubernetes cluster.
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
		Method:  "POST",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/pools",
		Headers: http.Header{},
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

type GetPoolRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// PoolID: the ID of the requested pool
	PoolID string `json:"-"`
}

// GetPool: get a pool in a cluster
//
// This method allows to get details about a specific pool.
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
		Method:  "GET",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "",
		Headers: http.Header{},
	}

	var resp Pool

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpgradePoolRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// PoolID: the ID of the pool to upgrade
	PoolID string `json:"-"`
	// Version: the new Kubernetes version for the pool
	Version string `json:"version"`
}

// UpgradePool: upgrade a pool in a cluster
//
// This method allows to upgrade the Kubernetes version of a specific pool. Note that this will work when the targeted version is the same than the version of the cluster.
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
		Method:  "POST",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "/upgrade",
		Headers: http.Header{},
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

type UpdatePoolRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// PoolID: the ID of the pool to update
	PoolID string `json:"-"`
	// Autoscaling: the new value for the enablement of autoscaling for the pool
	Autoscaling *bool `json:"autoscaling"`
	// Size: the new size for the pool
	Size *uint32 `json:"size"`
	// MinSize: the new minimun size for the pool
	MinSize *uint32 `json:"min_size"`
	// MaxSize: the new maximum size for the pool
	MaxSize *uint32 `json:"max_size"`
	// Autohealing: the new value for the enablement of autohealing for the pool
	Autohealing *bool `json:"autohealing"`
	// Tags: the new tags associated with the pool
	Tags *[]string `json:"tags"`
	// KubeletArgs: the new Kubelet arguments to be used by this pool. Note that this feature is to be considered as experimental
	KubeletArgs *map[string]string `json:"kubelet_args"`
	// UpgradePolicy: the Pool upgrade policy
	UpgradePolicy *UpdatePoolRequestUpgradePolicy `json:"upgrade_policy"`
}

// UpdatePool: update a pool in a cluster
//
// This method allows to update some attributes of a specific pool such as the size, the autoscaling enablement, the tags, ...
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
		Method:  "PATCH",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "",
		Headers: http.Header{},
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

type DeletePoolRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// PoolID: the ID of the pool to delete
	PoolID string `json:"-"`
}

// DeletePool: delete a pool in a cluster
//
// This method allows to delete a specific pool from a cluster, deleting all the nodes associated with it.
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
		Method:  "DELETE",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/pools/" + fmt.Sprint(req.PoolID) + "",
		Headers: http.Header{},
	}

	var resp Pool

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListNodesRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// ClusterID: the cluster ID from which the nodes will be listed from
	ClusterID string `json:"-"`
	// PoolID: the pool ID on which to filter the returned nodes
	PoolID *string `json:"-"`
	// OrderBy: the sort order of the returned nodes
	//
	// Default value: created_at_asc
	OrderBy ListNodesRequestOrderBy `json:"-"`
	// Page: the page number for the returned nodes
	Page *int32 `json:"-"`
	// PageSize: the maximum number of nodes per page
	PageSize *uint32 `json:"-"`
	// Name: the name on which to filter the returned nodes
	Name *string `json:"-"`
	// Status: the status on which to filter the returned nodes
	//
	// Default value: unknown
	Status NodeStatus `json:"-"`
}

// ListNodes: list all the nodes in a cluster
//
// This method allows to list all the existing nodes for a specific Kubernetes cluster.
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
		Method:  "GET",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/nodes",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListNodesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetNodeRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// NodeID: the ID of the requested node
	NodeID string `json:"-"`
}

// GetNode: get a node in a cluster
//
// This method allows to get details about a specific Kubernetes node.
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
		Method:  "GET",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "",
		Headers: http.Header{},
	}

	var resp Node

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ReplaceNodeRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// NodeID: the ID of the node to replace
	NodeID string `json:"-"`
}

// Deprecated: ReplaceNode: replace a node in a cluster
//
// This method allows to replace a specific node. The node will be set cordoned, meaning that scheduling will be disabled. Then the existing pods on the node will be drained and reschedule onto another schedulable node. Then the node will be deleted, and a new one will be created after the deletion. Note that when there is not enough space to reschedule all the pods (in a one node cluster for instance), you may experience some disruption of your applications.
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
		Method:  "POST",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "/replace",
		Headers: http.Header{},
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

type RebootNodeRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// NodeID: the ID of the node to reboot
	NodeID string `json:"-"`
}

// RebootNode: reboot a node in a cluster
//
// This method allows to reboot a specific node. This node will frist be cordoned, meaning that scheduling will be disabled. Then the existing pods on the node will be drained and reschedule onto another schedulable node. Note that when there is not enough space to reschedule all the pods (in a one node cluster for instance), you may experience some disruption of your applications.
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
		Method:  "POST",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "/reboot",
		Headers: http.Header{},
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

type DeleteNodeRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// NodeID: the ID of the node to replace
	NodeID string `json:"-"`
	// SkipDrain: skip draining node from its workload
	SkipDrain bool `json:"-"`
	// Replace: add a new node after the deletion of this node
	Replace bool `json:"-"`
}

// DeleteNode: delete a node in a cluster
//
// This method allows to delete a specific node. Note that when there is not enough space to reschedule all the pods (in a one node cluster for instance), you may experience some disruption of your applications.
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
		Method:  "DELETE",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "",
		Query:   query,
		Headers: http.Header{},
	}

	var resp Node

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListVersionsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
}

// ListVersions: list all available versions
//
// This method allows to list all available versions for the creation of a new Kubernetes cluster.
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
		Method:  "GET",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/versions",
		Headers: http.Header{},
	}

	var resp ListVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetVersionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// VersionName: the requested version name
	VersionName string `json:"-"`
}

// GetVersion: get details about a specific version
//
// This method allows to get a specific Kubernetes version and the details about the version.
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
		Method:  "GET",
		Path:    "/k8s/v1/regions/" + fmt.Sprint(req.Region) + "/versions/" + fmt.Sprint(req.VersionName) + "",
		Headers: http.Header{},
	}

	var resp Version

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClustersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClustersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListClustersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Clusters = append(r.Clusters, results.Clusters...)
	r.TotalCount += uint32(len(results.Clusters))
	return uint32(len(results.Clusters)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPoolsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPoolsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPoolsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Pools = append(r.Pools, results.Pools...)
	r.TotalCount += uint32(len(results.Pools))
	return uint32(len(results.Pools)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNodesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNodesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNodesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Nodes = append(r.Nodes, results.Nodes...)
	r.TotalCount += uint32(len(results.Nodes))
	return uint32(len(results.Nodes)), nil
}

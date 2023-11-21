// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package rdb_admin provides methods and message types of the rdb_admin v1 API.
package rdb_admin

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
	lb_v1 "github.com/scaleway/scaleway-sdk-go/api/lb/v1"
	rdb_v1 "github.com/scaleway/scaleway-sdk-go/api/rdb/v1"
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

type CleanupResourcesRequestResourceType string

const (
	CleanupResourcesRequestResourceTypeAll             = CleanupResourcesRequestResourceType("all")
	CleanupResourcesRequestResourceTypeInstances       = CleanupResourcesRequestResourceType("instances")
	CleanupResourcesRequestResourceTypeVolumes         = CleanupResourcesRequestResourceType("volumes")
	CleanupResourcesRequestResourceTypeFlexibleIPs     = CleanupResourcesRequestResourceType("flexible_ips")
	CleanupResourcesRequestResourceTypePlacementGroups = CleanupResourcesRequestResourceType("placement_groups")
	CleanupResourcesRequestResourceTypeSecurityGroups  = CleanupResourcesRequestResourceType("security_groups")
	CleanupResourcesRequestResourceTypeLoadBalancers   = CleanupResourcesRequestResourceType("load_balancers")
	CleanupResourcesRequestResourceTypeDNS             = CleanupResourcesRequestResourceType("dns")
	CleanupResourcesRequestResourceTypeBackups         = CleanupResourcesRequestResourceType("backups")
	CleanupResourcesRequestResourceTypeSnapshots       = CleanupResourcesRequestResourceType("snapshots")
)

func (enum CleanupResourcesRequestResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "all"
	}
	return string(enum)
}

func (enum CleanupResourcesRequestResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CleanupResourcesRequestResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CleanupResourcesRequestResourceType(CleanupResourcesRequestResourceType(tmp).String())
	return nil
}

type InstanceEventResourceType string

const (
	InstanceEventResourceTypeNone         = InstanceEventResourceType("none")
	InstanceEventResourceTypeBackups      = InstanceEventResourceType("backups")
	InstanceEventResourceTypeInstances    = InstanceEventResourceType("instances")
	InstanceEventResourceTypeSnapshots    = InstanceEventResourceType("snapshots")
	InstanceEventResourceTypeReadReplicas = InstanceEventResourceType("read_replicas")
)

func (enum InstanceEventResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "none"
	}
	return string(enum)
}

func (enum InstanceEventResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceEventResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceEventResourceType(InstanceEventResourceType(tmp).String())
	return nil
}

type InstanceEventStatus string

const (
	InstanceEventStatusInfo    = InstanceEventStatus("info")
	InstanceEventStatusSuccess = InstanceEventStatus("success")
	InstanceEventStatusError   = InstanceEventStatus("error")
)

func (enum InstanceEventStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "info"
	}
	return string(enum)
}

func (enum InstanceEventStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceEventStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceEventStatus(InstanceEventStatus(tmp).String())
	return nil
}

type ListInstanceEventsRequestOrderBy string

const (
	ListInstanceEventsRequestOrderByCreatedAtAsc  = ListInstanceEventsRequestOrderBy("created_at_asc")
	ListInstanceEventsRequestOrderByCreatedAtDesc = ListInstanceEventsRequestOrderBy("created_at_desc")
)

func (enum ListInstanceEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListInstanceEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInstanceEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInstanceEventsRequestOrderBy(ListInstanceEventsRequestOrderBy(tmp).String())
	return nil
}

type ListInstancesRequestOrderBy string

const (
	ListInstancesRequestOrderByCreatedAtAsc  = ListInstancesRequestOrderBy("created_at_asc")
	ListInstancesRequestOrderByCreatedAtDesc = ListInstancesRequestOrderBy("created_at_desc")
	ListInstancesRequestOrderByNameAsc       = ListInstancesRequestOrderBy("name_asc")
	ListInstancesRequestOrderByNameDesc      = ListInstancesRequestOrderBy("name_desc")
	ListInstancesRequestOrderByRegion        = ListInstancesRequestOrderBy("region")
)

func (enum ListInstancesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListInstancesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInstancesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInstancesRequestOrderBy(ListInstancesRequestOrderBy(tmp).String())
	return nil
}

type ListMaintenancesRequestOrderBy string

const (
	ListMaintenancesRequestOrderByCreatedAtAsc  = ListMaintenancesRequestOrderBy("created_at_asc")
	ListMaintenancesRequestOrderByCreatedAtDesc = ListMaintenancesRequestOrderBy("created_at_desc")
	ListMaintenancesRequestOrderByStartsAtAsc   = ListMaintenancesRequestOrderBy("starts_at_asc")
	ListMaintenancesRequestOrderByStartsAtDesc  = ListMaintenancesRequestOrderBy("starts_at_desc")
)

func (enum ListMaintenancesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListMaintenancesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListMaintenancesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListMaintenancesRequestOrderBy(ListMaintenancesRequestOrderBy(tmp).String())
	return nil
}

type MaintenanceStatus string

const (
	MaintenanceStatusUnknown  = MaintenanceStatus("unknown")
	MaintenanceStatusPending  = MaintenanceStatus("pending")
	MaintenanceStatusDone     = MaintenanceStatus("done")
	MaintenanceStatusCanceled = MaintenanceStatus("canceled")
)

func (enum MaintenanceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum MaintenanceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MaintenanceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MaintenanceStatus(MaintenanceStatus(tmp).String())
	return nil
}

type ReadReplicaStatus string

const (
	ReadReplicaStatusUnknown      = ReadReplicaStatus("unknown")
	ReadReplicaStatusProvisioning = ReadReplicaStatus("provisioning")
	ReadReplicaStatusInitializing = ReadReplicaStatus("initializing")
	ReadReplicaStatusReady        = ReadReplicaStatus("ready")
	ReadReplicaStatusDeleting     = ReadReplicaStatus("deleting")
	ReadReplicaStatusError        = ReadReplicaStatus("error")
	ReadReplicaStatusLocked       = ReadReplicaStatus("locked")
	ReadReplicaStatusConfiguring  = ReadReplicaStatus("configuring")
)

func (enum ReadReplicaStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ReadReplicaStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReadReplicaStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReadReplicaStatus(ReadReplicaStatus(tmp).String())
	return nil
}

// ConsulServer: consul server.
type ConsulServer struct {
	PublicIP string `json:"public_ip"`

	Hostname string `json:"hostname"`

	SecurityGroup string `json:"security_group"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`

	ConsulClusterKey string `json:"consul_cluster_key"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// BackupScheduleDetail: backup schedule detail.
type BackupScheduleDetail struct {
	Frequency uint32 `json:"frequency"`

	Retention uint32 `json:"retention"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	NextUpdate *time.Time `json:"next_update"`

	Disabled bool `json:"disabled"`
}

// ConsulCluster: consul cluster.
type ConsulCluster struct {
	ServiceIP string `json:"service_ip"`

	Port uint64 `json:"port"`

	Datacenter string `json:"datacenter"`

	ConsulClusterID string `json:"consul_cluster_id"`

	ConsulServers []*ConsulServer `json:"consul_servers"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// LB: lb.
type LB struct {
	LBID string `json:"lb_id"`

	LBBackendID string `json:"lb_backend_id"`

	LBFrontendID string `json:"lb_frontend_id"`
}

// Node: node.
type Node struct {
	ComputeKey string `json:"compute_key"`

	Name string `json:"name"`

	ImageKey string `json:"image_key"`

	PublicIPKey string `json:"public_ip_key"`

	PublicIP string `json:"public_ip"`

	PrivateIP string `json:"private_ip"`

	VolumeKey string `json:"volume_key"`

	VolumeSize uint64 `json:"volume_size"`

	CommercialType string `json:"commercial_type"`

	LocationZoneID string `json:"location_zone_id"`

	LocationInformation string `json:"location_information"`

	InstanceKey string `json:"instance_key"`

	ID string `json:"id"`

	ReadReplicaKey string `json:"read_replica_key"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// SecurityGroup: security group.
type SecurityGroup struct {
	Name string `json:"name"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`

	Description string `json:"description"`

	ComputeSgKey string `json:"compute_sg_key"`

	InstanceKey string `json:"instance_key"`

	Rules []*rdb_v1.ACLRule `json:"rules"`
}

// EngineMinorVersion: engine minor version.
type EngineMinorVersion struct {
	ID string `json:"id"`

	MinorVersion string `json:"minor_version"`

	Revision uint32 `json:"revision"`

	Enabled bool `json:"enabled"`
}

// EndpointDirectAccessDetails: endpoint direct access details.
type EndpointDirectAccessDetails struct {
}

// EndpointLoadBalancerDetails: endpoint load balancer details.
type EndpointLoadBalancerDetails struct {
}

// EndpointPrivateNetworkDetails: endpoint private network details.
type EndpointPrivateNetworkDetails struct {
	PrivateNetworkID string `json:"private_network_id"`

	ServiceIP scw.IPNet `json:"service_ip"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
}

// InstanceEvent: instance event.
type InstanceEvent struct {
	ID string `json:"id"`

	// Status: default value: info
	Status InstanceEventStatus `json:"status"`

	Name string `json:"name"`

	Details *string `json:"details"`

	CreatedAt *time.Time `json:"created_at"`

	ResourceID *string `json:"resource_id"`

	// ResourceType: default value: none
	ResourceType InstanceEventResourceType `json:"resource_type"`

	Internal *bool `json:"internal"`
}

// Instance: instance.
type Instance struct {
	PublicInstance *rdb_v1.Instance `json:"public_instance"`

	Nodes []*Node `json:"nodes"`

	SecurityGroups []*SecurityGroup `json:"security_groups"`

	ConsulCluster *ConsulCluster `json:"consul_cluster"`

	LB *LB `json:"lb"`

	LockReason string `json:"lock_reason"`

	BackupScheduleDetail *BackupScheduleDetail `json:"backup_schedule_detail"`
}

// Maintenance: maintenance.
type Maintenance struct {
	ID string `json:"id"`

	InstanceID string `json:"instance_id"`

	StartsAt *time.Time `json:"starts_at"`

	StopsAt *time.Time `json:"stops_at"`

	ClosedAt *time.Time `json:"closed_at"`

	Reason string `json:"reason"`

	// Status: default value: unknown
	Status MaintenanceStatus `json:"status"`

	TargetEngineMinorVersion *EngineMinorVersion `json:"target_engine_minor_version"`
}

// Endpoint: endpoint.
type Endpoint struct {
	// Precisely one of IP, Hostname must be set.
	IP *net.IP `json:"ip,omitempty"`

	Port uint32 `json:"port"`

	Name *string `json:"name"`

	// Precisely one of IP, Hostname must be set.
	Hostname *string `json:"hostname,omitempty"`

	ID string `json:"id"`

	// Precisely one of LoadBalancer, PrivateNetwork, DirectAccess must be set.
	LoadBalancer *EndpointLoadBalancerDetails `json:"load_balancer,omitempty"`

	// Precisely one of LoadBalancer, PrivateNetwork, DirectAccess must be set.
	PrivateNetwork *EndpointPrivateNetworkDetails `json:"private_network,omitempty"`

	// Precisely one of LoadBalancer, PrivateNetwork, DirectAccess must be set.
	DirectAccess *EndpointDirectAccessDetails `json:"direct_access,omitempty"`
}

// CleanupResourcesRequest: cleanup resources request.
type CleanupResourcesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ResourceType: default value: all
	ResourceType CleanupResourcesRequestResourceType `json:"resource_type"`
}

// ConsulClusterDetailed: consul cluster detailed.
type ConsulClusterDetailed struct {
	ServiceIP string `json:"service_ip"`

	Port uint64 `json:"port"`

	Datacenter string `json:"datacenter"`

	EncryptionKey string `json:"encryption_key"`

	CaSsl string `json:"ca_ssl"`

	CertSsl string `json:"cert_ssl"`

	KeySsl string `json:"key_ssl"`

	MasterToken string `json:"master_token"`

	ConsulClusterID string `json:"consul_cluster_id"`

	ConsulServers []*ConsulServer `json:"consul_servers"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// CreateConsulClusterRequest: create consul cluster request.
type CreateConsulClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ServiceIP string `json:"service_ip"`

	Port uint32 `json:"port"`

	Datacenter string `json:"datacenter"`

	EncryptionKey string `json:"encryption_key"`

	CaSsl string `json:"ca_ssl"`

	CertSsl string `json:"cert_ssl"`

	KeySsl string `json:"key_ssl"`

	MasterToken string `json:"master_token"`
}

// CreateConsulServerRequest: create consul server request.
type CreateConsulServerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	PublicIP string `json:"public_ip"`

	Hostname string `json:"hostname"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`

	SecurityGroup string `json:"security_group"`

	ConsulClusterKey string `json:"consul_cluster_key"`
}

// CreateMaintenanceRequest: create maintenance request.
type CreateMaintenanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"instance_id"`

	StartsAt *time.Time `json:"starts_at,omitempty"`

	StopsAt *time.Time `json:"stops_at,omitempty"`

	Reason string `json:"reason"`

	TargetEngineMinorVersion *string `json:"target_engine_minor_version,omitempty"`
}

// DatabaseBackup: database backup.
type DatabaseBackup struct {
	PublicDatabaseBackup *rdb_v1.DatabaseBackup `json:"public_database_backup"`
}

// DeleteMaintenanceRequest: delete maintenance request.
type DeleteMaintenanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	MaintenanceID string `json:"-"`
}

// DowngradeInstanceRequest: downgrade instance request.
type DowngradeInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`

	NodeType string `json:"node_type"`
}

// GetBackupRequest: get backup request.
type GetBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	BackupID string `json:"-"`
}

// GetConsulClusterRequest: get consul cluster request.
type GetConsulClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ConsulClusterID string `json:"-"`
}

// GetInstanceRequest: get instance request.
type GetInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`
}

// GetLBRequest: get lb request.
type GetLBRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	LBID string `json:"-"`
}

// GetMaintenanceRequest: get maintenance request.
type GetMaintenanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	MaintenanceID string `json:"-"`
}

// GetNodeRequest: get node request.
type GetNodeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	NodeID string `json:"-"`
}

// GetSecurityGroupsRequest: get security groups request.
type GetSecurityGroupsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	SecurityGroupID string `json:"-"`
}

// GetServiceInfoRequest: get service info request.
type GetServiceInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// HotReplaceNodeRequest: hot replace node request.
type HotReplaceNodeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	NodeID string `json:"-"`
}

// ListInstanceEventsRequest: list instance events request.
type ListInstanceEventsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListInstanceEventsRequestOrderBy `json:"-"`
}

// ListInstanceEventsResponse: list instance events response.
type ListInstanceEventsResponse struct {
	TotalCount uint32 `json:"total_count"`

	InstanceEvents []*InstanceEvent `json:"instance_events"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInstanceEventsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInstanceEventsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListInstanceEventsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.InstanceEvents = append(r.InstanceEvents, results.InstanceEvents...)
	r.TotalCount += uint32(len(results.InstanceEvents))
	return uint32(len(results.InstanceEvents)), nil
}

// ListInstancesRequest: list instances request.
type ListInstancesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListInstancesRequestOrderBy `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListInstancesResponse: list instances response.
type ListInstancesResponse struct {
	Instances []*Instance `json:"instances"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInstancesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInstancesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListInstancesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Instances = append(r.Instances, results.Instances...)
	r.TotalCount += uint32(len(results.Instances))
	return uint32(len(results.Instances)), nil
}

// ListMaintenancesRequest: list maintenances request.
type ListMaintenancesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListMaintenancesRequestOrderBy `json:"-"`

	InstanceID *string `json:"-"`

	IncludeClosed bool `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListMaintenancesResponse: list maintenances response.
type ListMaintenancesResponse struct {
	Maintenances []*Maintenance `json:"maintenances"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListMaintenancesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListMaintenancesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListMaintenancesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Maintenances = append(r.Maintenances, results.Maintenances...)
	r.TotalCount += uint32(len(results.Maintenances))
	return uint32(len(results.Maintenances)), nil
}

// PublicLB: public lb.
type PublicLB struct {
	PublicLB *lb_v1.LB `json:"public_lb"`
}

// ReadReplica: read replica.
type ReadReplica struct {
	ID string `json:"id"`

	Endpoints []*Endpoint `json:"endpoints"`

	// Status: default value: unknown
	Status ReadReplicaStatus `json:"status"`

	SameZone bool `json:"same_zone"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
}

// RegenerateInstanceSettingsRequest: regenerate instance settings request.
type RegenerateInstanceSettingsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`

	ForceRestart bool `json:"force_restart"`
}

// ReplayBackupRequest: replay backup request.
type ReplayBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	BackupID string `json:"-"`
}

// ReplayHotReplaceNodeRequest: replay hot replace node request.
type ReplayHotReplaceNodeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	NodeID string `json:"-"`
}

// ReplayInstanceCloneRequest: replay instance clone request.
type ReplayInstanceCloneRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`

	// Precisely one of InstanceSourceID, InstanceBackupID, SnapshotID must be set.
	InstanceSourceID *string `json:"instance_source_id,omitempty"`

	// Precisely one of InstanceSourceID, InstanceBackupID, SnapshotID must be set.
	InstanceBackupID *string `json:"instance_backup_id,omitempty"`

	// Precisely one of InstanceSourceID, InstanceBackupID, SnapshotID must be set.
	SnapshotID *string `json:"snapshot_id,omitempty"`
}

// ReplayInstanceDeletionRequest: replay instance deletion request.
type ReplayInstanceDeletionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`
}

// ReplayInstanceHAUpgradeRequest: replay instance ha upgrade request.
type ReplayInstanceHAUpgradeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`
}

// ReplayInstanceProvisioningRequest: replay instance provisioning request.
type ReplayInstanceProvisioningRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`

	UserName *string `json:"user_name,omitempty"`

	SnapshotID *string `json:"snapshot_id,omitempty"`
}

// ReplayInstanceUpgradeRequest: replay instance upgrade request.
type ReplayInstanceUpgradeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`

	OldNodeType *string `json:"old_node_type,omitempty"`

	NewNodeType *string `json:"new_node_type,omitempty"`
}

// ReplayInstanceVolumeSizeUpgradeRequest: replay instance volume size upgrade request.
type ReplayInstanceVolumeSizeUpgradeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`
}

// ReplayInstanceVolumeTypeChangeRequest: replay instance volume type change request.
type ReplayInstanceVolumeTypeChangeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`
}

// ReplayMigrateEndpointRequest: replay migrate endpoint request.
type ReplayMigrateEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	EndpointID string `json:"-"`

	InstanceID string `json:"instance_id"`
}

// ReplayReadReplicaCreateEndpointRequest: replay read replica create endpoint request.
type ReplayReadReplicaCreateEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ReadReplicaID string `json:"-"`

	EndpointID string `json:"endpoint_id"`
}

// ReplayReadReplicaDeletionRequest: replay read replica deletion request.
type ReplayReadReplicaDeletionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ReadReplicaID string `json:"-"`
}

// ReplayReadReplicaProvisioningRequest: replay read replica provisioning request.
type ReplayReadReplicaProvisioningRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ReadReplicaID string `json:"-"`
}

// ReplaySnapshotRequest: replay snapshot request.
type ReplaySnapshotRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	SnapshotID string `json:"-"`
}

// RestartInstanceDockerRequest: restart instance docker request.
type RestartInstanceDockerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`
}

// RestartInstanceRequest: restart instance request.
type RestartInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`
}

// RestoreDatabaseBackupRequest: restore database backup request.
type RestoreDatabaseBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	BackupID string `json:"-"`

	InstanceID string `json:"instance_id"`

	DatabaseName *string `json:"database_name,omitempty"`
}

// Snapshot: snapshot.
type Snapshot struct {
	PublicSnapshot *rdb_v1.Snapshot `json:"public_snapshot"`
}

// SyncLBACLsRequest: sync lbac ls request.
type SyncLBACLsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`
}

// SyncLBACLsResponse: sync lbac ls response.
type SyncLBACLsResponse struct {
	Rules []*rdb_v1.ACLRule `json:"rules"`
}

// UpdateBackupRequest: update backup request.
type UpdateBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	BackupID string `json:"-"`

	Status string `json:"status"`

	Size *int64 `json:"size,omitempty"`

	ExpiresAt *time.Time `json:"expires_at,omitempty"`

	CreatedAt *time.Time `json:"created_at,omitempty"`

	UpdatedAt *time.Time `json:"updated_at,omitempty"`
}

// UpdateBackupScheduleRequest: update backup schedule request.
type UpdateBackupScheduleRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`

	Frequency *int32 `json:"frequency,omitempty"`

	Retention *int32 `json:"retention,omitempty"`

	Disabled *bool `json:"disabled,omitempty"`

	NextUpdate *time.Time `json:"next_update,omitempty"`
}

// UpdateConsulClusterRequest: update consul cluster request.
type UpdateConsulClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ConsulClusterID string `json:"-"`

	ServiceIP string `json:"service_ip"`

	Port uint32 `json:"port"`

	Datacenter string `json:"datacenter"`

	EncryptionKey string `json:"encryption_key"`

	CaSsl string `json:"ca_ssl"`

	CertSsl string `json:"cert_ssl"`

	KeySsl string `json:"key_ssl"`

	MasterToken string `json:"master_token"`
}

// UpdateConsulServerRequest: update consul server request.
type UpdateConsulServerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	PublicIP string `json:"-"`

	Hostname string `json:"hostname"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`

	SecurityGroup string `json:"security_group"`

	ConsulClusterKey string `json:"consul_cluster_key"`
}

// UpdateInstanceRequest: update instance request.
type UpdateInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`

	Status string `json:"status"`
}

// UpdateMaintenanceRequest: update maintenance request.
type UpdateMaintenanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	MaintenanceID string `json:"-"`

	StartsAt *time.Time `json:"starts_at,omitempty"`

	StopsAt *time.Time `json:"stops_at,omitempty"`

	ClosedAt *time.Time `json:"closed_at,omitempty"`

	Reason *string `json:"reason,omitempty"`

	TargetEngineMinorVersion *string `json:"target_engine_minor_version,omitempty"`

	// Status: default value: unknown
	Status MaintenanceStatus `json:"status"`
}

// UpdateReadReplicaRequest: update read replica request.
type UpdateReadReplicaRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ReadReplicaID string `json:"-"`

	// Status: default value: unknown
	Status ReadReplicaStatus `json:"status"`
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
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInstances:
func (s *API) ListInstances(req *ListInstancesRequest, opts ...scw.RequestOption) (*ListInstancesResponse, error) {
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances",
		Query:  query,
	}

	var resp ListInstancesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInstance:
func (s *API) GetInstance(req *GetInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateInstance:
func (s *API) UpdateInstance(req *UpdateInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInstanceEvents:
func (s *API) ListInstanceEvents(req *ListInstanceEventsRequest, opts ...scw.RequestOption) (*ListInstanceEventsResponse, error) {
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

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/events",
		Query:  query,
	}

	var resp ListInstanceEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplayInstanceProvisioning:
func (s *API) ReplayInstanceProvisioning(req *ReplayInstanceProvisioningRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/replay-provisioning",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplayInstanceDeletion:
func (s *API) ReplayInstanceDeletion(req *ReplayInstanceDeletionRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/replay-deletion",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplayInstanceUpgrade:
func (s *API) ReplayInstanceUpgrade(req *ReplayInstanceUpgradeRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/replay-upgrade",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DowngradeInstance:
func (s *API) DowngradeInstance(req *DowngradeInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/downgrade",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplayInstanceClone:
func (s *API) ReplayInstanceClone(req *ReplayInstanceCloneRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/replay-clone",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplayInstanceHAUpgrade:
func (s *API) ReplayInstanceHAUpgrade(req *ReplayInstanceHAUpgradeRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/replay-ha-upgrade",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplayInstanceVolumeSizeUpgrade:
func (s *API) ReplayInstanceVolumeSizeUpgrade(req *ReplayInstanceVolumeSizeUpgradeRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/replay-volume-size-upgrade",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplayInstanceVolumeTypeChange:
func (s *API) ReplayInstanceVolumeTypeChange(req *ReplayInstanceVolumeTypeChangeRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/replay-volume-type-change",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestartInstanceDocker:
func (s *API) RestartInstanceDocker(req *RestartInstanceDockerRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/restart-docker",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBackupSchedule:
func (s *API) UpdateBackupSchedule(req *UpdateBackupScheduleRequest, opts ...scw.RequestOption) (*BackupScheduleDetail, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/backup-schedule",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp BackupScheduleDetail

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateReadReplica:
func (s *API) UpdateReadReplica(req *UpdateReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ReadReplicaID) == "" {
		return nil, errors.New("field ReadReplicaID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ReadReplica

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplayReadReplicaProvisioning:
func (s *API) ReplayReadReplicaProvisioning(req *ReplayReadReplicaProvisioningRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ReadReplicaID) == "" {
		return nil, errors.New("field ReadReplicaID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "/replay-provisioning",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ReadReplica

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplayReadReplicaDeletion:
func (s *API) ReplayReadReplicaDeletion(req *ReplayReadReplicaDeletionRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ReadReplicaID) == "" {
		return nil, errors.New("field ReadReplicaID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "/replay-deletion",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ReadReplica

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplayReadReplicaCreateEndpoint:
func (s *API) ReplayReadReplicaCreateEndpoint(req *ReplayReadReplicaCreateEndpointRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ReadReplicaID) == "" {
		return nil, errors.New("field ReadReplicaID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "/replay-create-endpoint",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ReadReplica

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBackup:
func (s *API) GetBackup(req *GetBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackupID) == "" {
		return nil, errors.New("field BackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.BackupID) + "",
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBackup:
func (s *API) UpdateBackup(req *UpdateBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackupID) == "" {
		return nil, errors.New("field BackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.BackupID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplayBackup:
func (s *API) ReplayBackup(req *ReplayBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackupID) == "" {
		return nil, errors.New("field BackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.BackupID) + "/replay",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestoreDatabaseBackup:
func (s *API) RestoreDatabaseBackup(req *RestoreDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackupID) == "" {
		return nil, errors.New("field BackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.BackupID) + "/restore",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplaySnapshot:
func (s *API) ReplaySnapshot(req *ReplaySnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "/replay",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetConsulCluster:
func (s *API) GetConsulCluster(req *GetConsulClusterRequest, opts ...scw.RequestOption) (*ConsulClusterDetailed, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ConsulClusterID) == "" {
		return nil, errors.New("field ConsulClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/consul-clusters/" + fmt.Sprint(req.ConsulClusterID) + "",
	}

	var resp ConsulClusterDetailed

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateConsulCluster:
func (s *API) UpdateConsulCluster(req *UpdateConsulClusterRequest, opts ...scw.RequestOption) (*ConsulClusterDetailed, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ConsulClusterID) == "" {
		return nil, errors.New("field ConsulClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/consul-clusters/" + fmt.Sprint(req.ConsulClusterID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ConsulClusterDetailed

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateConsulServer:
func (s *API) UpdateConsulServer(req *UpdateConsulServerRequest, opts ...scw.RequestOption) (*ConsulServer, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PublicIP) == "" {
		return nil, errors.New("field PublicIP cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/consul-servers/" + fmt.Sprint(req.PublicIP) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ConsulServer

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateConsulCluster:
func (s *API) CreateConsulCluster(req *CreateConsulClusterRequest, opts ...scw.RequestOption) (*ConsulClusterDetailed, error) {
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
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/consul-clusters",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ConsulClusterDetailed

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateConsulServer:
func (s *API) CreateConsulServer(req *CreateConsulServerRequest, opts ...scw.RequestOption) (*ConsulServer, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/consul-servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ConsulServer

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNode:
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
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "",
	}

	var resp Node

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// HotReplaceNode:
func (s *API) HotReplaceNode(req *HotReplaceNodeRequest, opts ...scw.RequestOption) (*Node, error) {
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
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "/hot-replace",
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

// ReplayHotReplaceNode:
func (s *API) ReplayHotReplaceNode(req *ReplayHotReplaceNodeRequest, opts ...scw.RequestOption) (*Node, error) {
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
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/nodes/" + fmt.Sprint(req.NodeID) + "/replay-hot-replace",
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

// GetSecurityGroups:
func (s *API) GetSecurityGroups(req *GetSecurityGroupsRequest, opts ...scw.RequestOption) (*SecurityGroup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/security-groups/" + fmt.Sprint(req.SecurityGroupID) + "",
	}

	var resp SecurityGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLB:
func (s *API) GetLB(req *GetLBRequest, opts ...scw.RequestOption) (*PublicLB, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "",
	}

	var resp PublicLB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncLBACLs:
func (s *API) SyncLBACLs(req *SyncLBACLsRequest, opts ...scw.RequestOption) (*SyncLBACLsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/sync-acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SyncLBACLsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RegenerateInstanceSettings:
func (s *API) RegenerateInstanceSettings(req *RegenerateInstanceSettingsRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/regenerate-settings",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CleanupResources:
func (s *API) CleanupResources(req *CleanupResourcesRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/cleanup",
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

// RestartInstance:
func (s *API) RestartInstance(req *RestartInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/restart",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListMaintenances:
func (s *API) ListMaintenances(req *ListMaintenancesRequest, opts ...scw.RequestOption) (*ListMaintenancesResponse, error) {
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
	parameter.AddToQuery(query, "instance_id", req.InstanceID)
	parameter.AddToQuery(query, "include_closed", req.IncludeClosed)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/maintenances",
		Query:  query,
	}

	var resp ListMaintenancesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMaintenance:
func (s *API) GetMaintenance(req *GetMaintenanceRequest, opts ...scw.RequestOption) (*Maintenance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.MaintenanceID) == "" {
		return nil, errors.New("field MaintenanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/maintenances/" + fmt.Sprint(req.MaintenanceID) + "",
	}

	var resp Maintenance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateMaintenance:
func (s *API) CreateMaintenance(req *CreateMaintenanceRequest, opts ...scw.RequestOption) (*Maintenance, error) {
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
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/maintenances",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Maintenance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateMaintenance:
func (s *API) UpdateMaintenance(req *UpdateMaintenanceRequest, opts ...scw.RequestOption) (*Maintenance, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.MaintenanceID) == "" {
		return nil, errors.New("field MaintenanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/maintenances/" + fmt.Sprint(req.MaintenanceID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Maintenance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteMaintenance:
func (s *API) DeleteMaintenance(req *DeleteMaintenanceRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.MaintenanceID) == "" {
		return errors.New("field MaintenanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/maintenances/" + fmt.Sprint(req.MaintenanceID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ReplayMigrateEndpoint:
func (s *API) ReplayMigrateEndpoint(req *ReplayMigrateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return nil, errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/rdb-admin/v1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "/migrate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Endpoint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

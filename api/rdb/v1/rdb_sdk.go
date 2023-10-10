// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package rdb provides methods and message types of the rdb v1 API.
package rdb

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

// API: managed Database for PostgreSQL and MySQL API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type ACLRuleAction string

const (
	ACLRuleActionAllow = ACLRuleAction("allow")
	ACLRuleActionDeny  = ACLRuleAction("deny")
)

func (enum ACLRuleAction) String() string {
	if enum == "" {
		// return default value if empty
		return "allow"
	}
	return string(enum)
}

func (enum ACLRuleAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ACLRuleAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ACLRuleAction(ACLRuleAction(tmp).String())
	return nil
}

type ACLRuleDirection string

const (
	ACLRuleDirectionInbound  = ACLRuleDirection("inbound")
	ACLRuleDirectionOutbound = ACLRuleDirection("outbound")
)

func (enum ACLRuleDirection) String() string {
	if enum == "" {
		// return default value if empty
		return "inbound"
	}
	return string(enum)
}

func (enum ACLRuleDirection) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ACLRuleDirection) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ACLRuleDirection(ACLRuleDirection(tmp).String())
	return nil
}

type ACLRuleProtocol string

const (
	ACLRuleProtocolTCP  = ACLRuleProtocol("tcp")
	ACLRuleProtocolUDP  = ACLRuleProtocol("udp")
	ACLRuleProtocolIcmp = ACLRuleProtocol("icmp")
)

func (enum ACLRuleProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "tcp"
	}
	return string(enum)
}

func (enum ACLRuleProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ACLRuleProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ACLRuleProtocol(ACLRuleProtocol(tmp).String())
	return nil
}

type DatabaseBackupStatus string

const (
	DatabaseBackupStatusUnknown   = DatabaseBackupStatus("unknown")
	DatabaseBackupStatusCreating  = DatabaseBackupStatus("creating")
	DatabaseBackupStatusReady     = DatabaseBackupStatus("ready")
	DatabaseBackupStatusRestoring = DatabaseBackupStatus("restoring")
	DatabaseBackupStatusDeleting  = DatabaseBackupStatus("deleting")
	DatabaseBackupStatusError     = DatabaseBackupStatus("error")
	DatabaseBackupStatusExporting = DatabaseBackupStatus("exporting")
	DatabaseBackupStatusLocked    = DatabaseBackupStatus("locked")
)

func (enum DatabaseBackupStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DatabaseBackupStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DatabaseBackupStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DatabaseBackupStatus(DatabaseBackupStatus(tmp).String())
	return nil
}

type EngineSettingPropertyType string

const (
	EngineSettingPropertyTypeBOOLEAN = EngineSettingPropertyType("BOOLEAN")
	EngineSettingPropertyTypeINT     = EngineSettingPropertyType("INT")
	EngineSettingPropertyTypeSTRING  = EngineSettingPropertyType("STRING")
	EngineSettingPropertyTypeFLOAT   = EngineSettingPropertyType("FLOAT")
)

func (enum EngineSettingPropertyType) String() string {
	if enum == "" {
		// return default value if empty
		return "BOOLEAN"
	}
	return string(enum)
}

func (enum EngineSettingPropertyType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EngineSettingPropertyType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EngineSettingPropertyType(EngineSettingPropertyType(tmp).String())
	return nil
}

type InstanceLogStatus string

const (
	InstanceLogStatusUnknown  = InstanceLogStatus("unknown")
	InstanceLogStatusReady    = InstanceLogStatus("ready")
	InstanceLogStatusCreating = InstanceLogStatus("creating")
	InstanceLogStatusError    = InstanceLogStatus("error")
)

func (enum InstanceLogStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum InstanceLogStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceLogStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceLogStatus(InstanceLogStatus(tmp).String())
	return nil
}

type InstanceStatus string

const (
	InstanceStatusUnknown      = InstanceStatus("unknown")
	InstanceStatusReady        = InstanceStatus("ready")
	InstanceStatusProvisioning = InstanceStatus("provisioning")
	InstanceStatusConfiguring  = InstanceStatus("configuring")
	InstanceStatusDeleting     = InstanceStatus("deleting")
	InstanceStatusError        = InstanceStatus("error")
	InstanceStatusAutohealing  = InstanceStatus("autohealing")
	InstanceStatusLocked       = InstanceStatus("locked")
	InstanceStatusInitializing = InstanceStatus("initializing")
	InstanceStatusDiskFull     = InstanceStatus("disk_full")
	InstanceStatusBackuping    = InstanceStatus("backuping")
	InstanceStatusSnapshotting = InstanceStatus("snapshotting")
	InstanceStatusRestarting   = InstanceStatus("restarting")
)

func (enum InstanceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum InstanceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceStatus(InstanceStatus(tmp).String())
	return nil
}

type ListDatabaseBackupsRequestOrderBy string

const (
	ListDatabaseBackupsRequestOrderByCreatedAtAsc  = ListDatabaseBackupsRequestOrderBy("created_at_asc")
	ListDatabaseBackupsRequestOrderByCreatedAtDesc = ListDatabaseBackupsRequestOrderBy("created_at_desc")
	ListDatabaseBackupsRequestOrderByNameAsc       = ListDatabaseBackupsRequestOrderBy("name_asc")
	ListDatabaseBackupsRequestOrderByNameDesc      = ListDatabaseBackupsRequestOrderBy("name_desc")
	ListDatabaseBackupsRequestOrderByStatusAsc     = ListDatabaseBackupsRequestOrderBy("status_asc")
	ListDatabaseBackupsRequestOrderByStatusDesc    = ListDatabaseBackupsRequestOrderBy("status_desc")
)

func (enum ListDatabaseBackupsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListDatabaseBackupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDatabaseBackupsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDatabaseBackupsRequestOrderBy(ListDatabaseBackupsRequestOrderBy(tmp).String())
	return nil
}

type ListDatabasesRequestOrderBy string

const (
	ListDatabasesRequestOrderByNameAsc  = ListDatabasesRequestOrderBy("name_asc")
	ListDatabasesRequestOrderByNameDesc = ListDatabasesRequestOrderBy("name_desc")
	ListDatabasesRequestOrderBySizeAsc  = ListDatabasesRequestOrderBy("size_asc")
	ListDatabasesRequestOrderBySizeDesc = ListDatabasesRequestOrderBy("size_desc")
)

func (enum ListDatabasesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListDatabasesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDatabasesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDatabasesRequestOrderBy(ListDatabasesRequestOrderBy(tmp).String())
	return nil
}

type ListInstanceLogsRequestOrderBy string

const (
	ListInstanceLogsRequestOrderByCreatedAtAsc  = ListInstanceLogsRequestOrderBy("created_at_asc")
	ListInstanceLogsRequestOrderByCreatedAtDesc = ListInstanceLogsRequestOrderBy("created_at_desc")
)

func (enum ListInstanceLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListInstanceLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInstanceLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInstanceLogsRequestOrderBy(ListInstanceLogsRequestOrderBy(tmp).String())
	return nil
}

type ListInstancesRequestOrderBy string

const (
	ListInstancesRequestOrderByCreatedAtAsc  = ListInstancesRequestOrderBy("created_at_asc")
	ListInstancesRequestOrderByCreatedAtDesc = ListInstancesRequestOrderBy("created_at_desc")
	ListInstancesRequestOrderByNameAsc       = ListInstancesRequestOrderBy("name_asc")
	ListInstancesRequestOrderByNameDesc      = ListInstancesRequestOrderBy("name_desc")
	ListInstancesRequestOrderByRegion        = ListInstancesRequestOrderBy("region")
	ListInstancesRequestOrderByStatusAsc     = ListInstancesRequestOrderBy("status_asc")
	ListInstancesRequestOrderByStatusDesc    = ListInstancesRequestOrderBy("status_desc")
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

type ListPrivilegesRequestOrderBy string

const (
	ListPrivilegesRequestOrderByUserNameAsc      = ListPrivilegesRequestOrderBy("user_name_asc")
	ListPrivilegesRequestOrderByUserNameDesc     = ListPrivilegesRequestOrderBy("user_name_desc")
	ListPrivilegesRequestOrderByDatabaseNameAsc  = ListPrivilegesRequestOrderBy("database_name_asc")
	ListPrivilegesRequestOrderByDatabaseNameDesc = ListPrivilegesRequestOrderBy("database_name_desc")
)

func (enum ListPrivilegesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "user_name_asc"
	}
	return string(enum)
}

func (enum ListPrivilegesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPrivilegesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPrivilegesRequestOrderBy(ListPrivilegesRequestOrderBy(tmp).String())
	return nil
}

type ListSnapshotsRequestOrderBy string

const (
	ListSnapshotsRequestOrderByCreatedAtAsc  = ListSnapshotsRequestOrderBy("created_at_asc")
	ListSnapshotsRequestOrderByCreatedAtDesc = ListSnapshotsRequestOrderBy("created_at_desc")
	ListSnapshotsRequestOrderByNameAsc       = ListSnapshotsRequestOrderBy("name_asc")
	ListSnapshotsRequestOrderByNameDesc      = ListSnapshotsRequestOrderBy("name_desc")
	ListSnapshotsRequestOrderByExpiresAtAsc  = ListSnapshotsRequestOrderBy("expires_at_asc")
	ListSnapshotsRequestOrderByExpiresAtDesc = ListSnapshotsRequestOrderBy("expires_at_desc")
)

func (enum ListSnapshotsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSnapshotsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSnapshotsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSnapshotsRequestOrderBy(ListSnapshotsRequestOrderBy(tmp).String())
	return nil
}

type ListUsersRequestOrderBy string

const (
	ListUsersRequestOrderByNameAsc     = ListUsersRequestOrderBy("name_asc")
	ListUsersRequestOrderByNameDesc    = ListUsersRequestOrderBy("name_desc")
	ListUsersRequestOrderByIsAdminAsc  = ListUsersRequestOrderBy("is_admin_asc")
	ListUsersRequestOrderByIsAdminDesc = ListUsersRequestOrderBy("is_admin_desc")
)

func (enum ListUsersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListUsersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListUsersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListUsersRequestOrderBy(ListUsersRequestOrderBy(tmp).String())
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

type NodeTypeGeneration string

const (
	NodeTypeGenerationUnknownGeneration = NodeTypeGeneration("unknown_generation")
	NodeTypeGenerationGenerationV1      = NodeTypeGeneration("generation_v1")
	NodeTypeGenerationGenerationV2      = NodeTypeGeneration("generation_v2")
)

func (enum NodeTypeGeneration) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_generation"
	}
	return string(enum)
}

func (enum NodeTypeGeneration) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NodeTypeGeneration) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NodeTypeGeneration(NodeTypeGeneration(tmp).String())
	return nil
}

type NodeTypeStock string

const (
	NodeTypeStockUnknown    = NodeTypeStock("unknown")
	NodeTypeStockLowStock   = NodeTypeStock("low_stock")
	NodeTypeStockOutOfStock = NodeTypeStock("out_of_stock")
	NodeTypeStockAvailable  = NodeTypeStock("available")
)

func (enum NodeTypeStock) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NodeTypeStock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NodeTypeStock) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NodeTypeStock(NodeTypeStock(tmp).String())
	return nil
}

type Permission string

const (
	PermissionReadonly  = Permission("readonly")
	PermissionReadwrite = Permission("readwrite")
	PermissionAll       = Permission("all")
	PermissionCustom    = Permission("custom")
	PermissionNone      = Permission("none")
)

func (enum Permission) String() string {
	if enum == "" {
		// return default value if empty
		return "readonly"
	}
	return string(enum)
}

func (enum Permission) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Permission) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Permission(Permission(tmp).String())
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
	ReadReplicaStatusPromoting    = ReadReplicaStatus("promoting")
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

type SnapshotStatus string

const (
	SnapshotStatusUnknown   = SnapshotStatus("unknown")
	SnapshotStatusCreating  = SnapshotStatus("creating")
	SnapshotStatusReady     = SnapshotStatus("ready")
	SnapshotStatusRestoring = SnapshotStatus("restoring")
	SnapshotStatusDeleting  = SnapshotStatus("deleting")
	SnapshotStatusError     = SnapshotStatus("error")
	SnapshotStatusLocked    = SnapshotStatus("locked")
)

func (enum SnapshotStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum SnapshotStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SnapshotStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SnapshotStatus(SnapshotStatus(tmp).String())
	return nil
}

type StorageClass string

const (
	StorageClassUnknownStorageClass = StorageClass("unknown_storage_class")
	StorageClassLssd                = StorageClass("lssd")
	StorageClassBssd                = StorageClass("bssd")
	StorageClassSbs                 = StorageClass("sbs")
)

func (enum StorageClass) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_storage_class"
	}
	return string(enum)
}

func (enum StorageClass) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *StorageClass) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = StorageClass(StorageClass(tmp).String())
	return nil
}

type VolumeType string

const (
	VolumeTypeLssd   = VolumeType("lssd")
	VolumeTypeBssd   = VolumeType("bssd")
	VolumeTypeSbs5k  = VolumeType("sbs_5k")
	VolumeTypeSbs15k = VolumeType("sbs_15k")
)

func (enum VolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return "lssd"
	}
	return string(enum)
}

func (enum VolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeType(VolumeType(tmp).String())
	return nil
}

type ACLRule struct {
	IP scw.IPNet `json:"ip"`
	// Deprecated
	Port *uint32 `json:"port,omitempty"`
	// Protocol: default value: tcp
	Protocol ACLRuleProtocol `json:"protocol"`
	// Direction: default value: inbound
	Direction ACLRuleDirection `json:"direction"`
	// Action: default value: allow
	Action ACLRuleAction `json:"action"`

	Description string `json:"description"`
}

type ACLRuleRequest struct {
	IP scw.IPNet `json:"ip"`

	Description string `json:"description"`
}

// AddInstanceACLRulesResponse: add instance acl rules response.
type AddInstanceACLRulesResponse struct {
	// Rules: ACL Rules enabled for the Database Instance.
	Rules []*ACLRule `json:"rules"`
}

// AddInstanceSettingsResponse: add instance settings response.
type AddInstanceSettingsResponse struct {
	// Settings: settings available on the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// BackupSchedule: backup schedule.
type BackupSchedule struct {
	// Frequency: frequency of the backup schedule (in hours).
	Frequency uint32 `json:"frequency"`
	// Retention: default retention period of backups (in days).
	Retention uint32 `json:"retention"`
	// Disabled: defines whether the backup schedule feature is disabled.
	Disabled bool `json:"disabled"`
	// NextRunAt: next run of the backup schedule (accurate to 10 minutes).
	NextRunAt *time.Time `json:"next_run_at"`
}

// Database: database.
type Database struct {
	// Name: name of the database.
	Name string `json:"name"`
	// Owner: name of the database owner.
	Owner string `json:"owner"`
	// Managed: defines whether the database is managed or not.
	Managed bool `json:"managed"`
	// Size: size of the database.
	Size scw.Size `json:"size"`
}

// DatabaseBackup: database backup.
type DatabaseBackup struct {
	// ID: UUID of the database backup.
	ID string `json:"id"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"instance_id"`
	// DatabaseName: name of backed up database.
	DatabaseName string `json:"database_name"`
	// Name: name of the backup.
	Name string `json:"name"`
	// Status: status of the backup.
	// Default value: unknown
	Status DatabaseBackupStatus `json:"status"`
	// Size: size of the database backup.
	Size *scw.Size `json:"size"`
	// ExpiresAt: expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at"`
	// CreatedAt: creation date (must follow the ISO 8601 format).
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: updated date (must follow the ISO 8601 format).
	UpdatedAt *time.Time `json:"updated_at"`
	// InstanceName: name of the Database Instance of the backup.
	InstanceName string `json:"instance_name"`
	// DownloadURL: URL you can download the backup from.
	DownloadURL *string `json:"download_url"`
	// DownloadURLExpiresAt: expiration date of the download link.
	DownloadURLExpiresAt *time.Time `json:"download_url_expires_at"`
	// Region: region of the database backup.
	Region scw.Region `json:"region"`
	// SameRegion: store logical backups in the same region as the source Database Instance.
	SameRegion bool `json:"same_region"`
}

// DatabaseEngine: database engine.
type DatabaseEngine struct {
	// Name: engine name.
	Name string `json:"name"`
	// LogoURL: engine logo URL.
	LogoURL string `json:"logo_url"`
	// Versions: available versions.
	Versions []*EngineVersion `json:"versions"`
	// Region: region of this Database Instance.
	Region scw.Region `json:"region"`
}

// DeleteInstanceACLRulesResponse: delete instance acl rules response.
type DeleteInstanceACLRulesResponse struct {
	// Rules: IP addresses defined in the ACL rules of the Database Instance.
	Rules []*ACLRule `json:"rules"`
}

// DeleteInstanceSettingsResponse: delete instance settings response.
type DeleteInstanceSettingsResponse struct {
	// Settings: settings names to delete from the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// Endpoint: endpoint.
type Endpoint struct {
	// ID: UUID of the endpoint.
	ID string `json:"id"`
	// IP: iPv4 address of the endpoint.
	// Precisely one of Hostname, IP must be set.
	IP *net.IP `json:"ip,omitempty"`
	// Port: TCP port of the endpoint.
	Port uint32 `json:"port"`
	// Name: name of the endpoint.
	Name *string `json:"name"`
	// PrivateNetwork: private Network details. One maximum per Database Instance or Read Replica (a Database Instance and its Read Replica can have different Private Networks). Cannot be updated (has to be deleted and recreated).
	// Precisely one of DirectAccess, LoadBalancer, PrivateNetwork must be set.
	PrivateNetwork *EndpointPrivateNetworkDetails `json:"private_network,omitempty"`
	// LoadBalancer: load balancer details. Public endpoint for Database Instance which is systematically present. One per Database Instance.
	// Precisely one of DirectAccess, LoadBalancer, PrivateNetwork must be set.
	LoadBalancer *EndpointLoadBalancerDetails `json:"load_balancer,omitempty"`
	// DirectAccess: direct access details. Public endpoint reserved for Read Replicas. One per Read Replica.
	// Precisely one of DirectAccess, LoadBalancer, PrivateNetwork must be set.
	DirectAccess *EndpointDirectAccessDetails `json:"direct_access,omitempty"`
	// Hostname: hostname of the endpoint.
	// Precisely one of Hostname, IP must be set.
	Hostname *string `json:"hostname,omitempty"`
}

type EndpointDirectAccessDetails struct {
}

type EndpointLoadBalancerDetails struct {
}

// EndpointPrivateNetworkDetails: endpoint. private network details.
type EndpointPrivateNetworkDetails struct {
	// PrivateNetworkID: UUID of the private network.
	PrivateNetworkID string `json:"private_network_id"`
	// ServiceIP: cIDR notation of the endpoint IPv4 address.
	ServiceIP scw.IPNet `json:"service_ip"`
	// Zone: private network zone.
	Zone scw.Zone `json:"zone"`
}

// EndpointSpec: endpoint spec.
type EndpointSpec struct {
	// LoadBalancer: load balancer endpoint specifications. Public endpoint for Database Instance which is systematically present. One per RDB instance.
	// Precisely one of LoadBalancer, PrivateNetwork must be set.
	LoadBalancer *EndpointSpecLoadBalancer `json:"load_balancer,omitempty"`
	// PrivateNetwork: private Network endpoint specifications. One maximum per Database Instance or Read Replica (a Database Instance and its Read Replica can have different Private Networks). Cannot be updated (has to be deleted and recreated).
	// Precisely one of LoadBalancer, PrivateNetwork must be set.
	PrivateNetwork *EndpointSpecPrivateNetwork `json:"private_network,omitempty"`
}

type EndpointSpecLoadBalancer struct {
}

// EndpointSpecPrivateNetwork: endpoint spec. private network.
type EndpointSpecPrivateNetwork struct {
	// PrivateNetworkID: UUID of the Private Network to be connected to the Database Instance.
	PrivateNetworkID string `json:"private_network_id"`
	// ServiceIP: endpoint IPv4 address with a CIDR notation. Refer to the official Scaleway documentation to learn more about IP and subnet limitations.
	// Precisely one of IpamConfig, ServiceIP must be set.
	ServiceIP *scw.IPNet `json:"service_ip,omitempty"`
	// IpamConfig: automated configuration of your Private Network endpoint with Scaleway IPAM service. One at the most per Database Instance or Read Replica (a Database Instance and its Read Replica can have different Private Networks). Cannot be updated (has to be deleted and recreated).
	// Precisely one of IpamConfig, ServiceIP must be set.
	IpamConfig *EndpointSpecPrivateNetworkIpamConfig `json:"ipam_config,omitempty"`
}

type EndpointSpecPrivateNetworkIpamConfig struct {
}

// EngineSetting: engine setting.
type EngineSetting struct {
	// Name: setting name from the database engine.
	Name string `json:"name"`
	// DefaultValue: value set when not specified.
	DefaultValue string `json:"default_value"`
	// HotConfigurable: setting can be applied without restarting.
	HotConfigurable bool `json:"hot_configurable"`
	// Description: setting description.
	Description string `json:"description"`
	// PropertyType: setting type.
	// Default value: BOOLEAN
	PropertyType EngineSettingPropertyType `json:"property_type"`
	// Unit: setting base unit.
	Unit *string `json:"unit"`
	// StringConstraint: validation regex for string type settings.
	StringConstraint *string `json:"string_constraint"`
	// IntMin: minimum value for int types.
	IntMin *int32 `json:"int_min"`
	// IntMax: maximum value for int types.
	IntMax *int32 `json:"int_max"`
	// FloatMin: minimum value for float types.
	FloatMin *float32 `json:"float_min"`
	// FloatMax: maximum value for float types.
	FloatMax *float32 `json:"float_max"`
}

// EngineVersion: engine version.
type EngineVersion struct {
	// Version: database engine version.
	Version string `json:"version"`
	// Name: database engine name.
	Name string `json:"name"`
	// EndOfLife: end of life date.
	EndOfLife *time.Time `json:"end_of_life"`
	// AvailableSettings: engine settings available to be set.
	AvailableSettings []*EngineSetting `json:"available_settings"`
	// Disabled: disabled versions cannot be created.
	Disabled bool `json:"disabled"`
	// Beta: beta status of engine version.
	Beta bool `json:"beta"`
	// AvailableInitSettings: engine settings available to be set at database initialization.
	AvailableInitSettings []*EngineSetting `json:"available_init_settings"`
}

// Instance: instance.
type Instance struct {
	// CreatedAt: creation date (must follow the ISO 8601 format).
	CreatedAt *time.Time `json:"created_at"`
	// Volume: volumes of the Database Instance.
	Volume *Volume `json:"volume"`
	// Region: region the Database Instance is in.
	Region scw.Region `json:"region"`
	// ID: UUID of the Database Instance.
	ID string `json:"id"`
	// Name: name of the Database Instance.
	Name string `json:"name"`
	// OrganizationID: organization ID the Database Instance belongs to.
	OrganizationID string `json:"organization_id"`
	// ProjectID: project ID the Database Instance belongs to.
	ProjectID string `json:"project_id"`
	// Status: status of the Database Instance.
	// Default value: unknown
	Status InstanceStatus `json:"status"`
	// Engine: database engine of the database (PostgreSQL, MySQL, ...).
	Engine string `json:"engine"`
	// UpgradableVersion: available database engine versions for upgrade.
	UpgradableVersion []*UpgradableVersion `json:"upgradable_version"`
	// Deprecated: Endpoint: endpoint of the Database Instance.
	Endpoint *Endpoint `json:"endpoint,omitempty"`
	// Tags: list of tags applied to the Database Instance.
	Tags []string `json:"tags"`
	// Settings: advanced settings of the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
	// BackupSchedule: backup schedule of the Database Instance.
	BackupSchedule *BackupSchedule `json:"backup_schedule"`
	// IsHaCluster: defines whether or not High-Availability is enabled.
	IsHaCluster bool `json:"is_ha_cluster"`
	// ReadReplicas: read Replicas of the Database Instance.
	ReadReplicas []*ReadReplica `json:"read_replicas"`
	// NodeType: node type of the Database Instance.
	NodeType string `json:"node_type"`
	// InitSettings: list of engine settings to be set at database initialization.
	InitSettings []*InstanceSetting `json:"init_settings"`
	// Endpoints: list of Database Instance endpoints.
	Endpoints []*Endpoint `json:"endpoints"`
	// LogsPolicy: logs policy of the Database Instance.
	LogsPolicy *LogsPolicy `json:"logs_policy"`
	// BackupSameRegion: store logical backups in the same region as the Database Instance.
	BackupSameRegion bool `json:"backup_same_region"`
	// Maintenances: list of Database Instance maintenance events.
	Maintenances []*Maintenance `json:"maintenances"`
}

// InstanceLog: instance log.
type InstanceLog struct {
	// DownloadURL: presigned S3 URL to download your log file.
	DownloadURL *string `json:"download_url"`
	// ID: UUID of the Database Instance log.
	ID string `json:"id"`
	// Status: status of the logs in a Database Instance.
	// Default value: unknown
	Status InstanceLogStatus `json:"status"`
	// NodeName: name of the underlying node.
	NodeName string `json:"node_name"`
	// ExpiresAt: expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at"`
	// CreatedAt: creation date (must follow the ISO 8601 format).
	CreatedAt *time.Time `json:"created_at"`
	// Region: region the Database Instance is in.
	Region scw.Region `json:"region"`
}

// InstanceMetrics: instance metrics.
type InstanceMetrics struct {
	// Timeseries: time series of metrics of a Database Instance.
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

type InstanceSetting struct {
	Name string `json:"name"`

	Value string `json:"value"`
}

// ListDatabaseBackupsResponse: list database backups response.
type ListDatabaseBackupsResponse struct {
	// DatabaseBackups: list of database backups.
	DatabaseBackups []*DatabaseBackup `json:"database_backups"`
	// TotalCount: total count of database backups available.
	TotalCount uint32 `json:"total_count"`
}

// ListDatabaseEnginesResponse: list database engines response.
type ListDatabaseEnginesResponse struct {
	// Engines: list of the available database engines.
	Engines []*DatabaseEngine `json:"engines"`
	// TotalCount: total count of database engines available.
	TotalCount uint32 `json:"total_count"`
}

// ListDatabasesResponse: list databases response.
type ListDatabasesResponse struct {
	// Databases: list of the databases.
	Databases []*Database `json:"databases"`
	// TotalCount: total count of databases present on a Database Instance.
	TotalCount uint32 `json:"total_count"`
}

// ListInstanceACLRulesResponse: list instance acl rules response.
type ListInstanceACLRulesResponse struct {
	// Rules: list of ACL rules present on a Database Instance.
	Rules []*ACLRule `json:"rules"`
	// TotalCount: total count of ACL rules present on a Database Instance.
	TotalCount uint32 `json:"total_count"`
}

// ListInstanceLogsDetailsResponse: list instance logs details response.
type ListInstanceLogsDetailsResponse struct {
	// Details: remote Database Instance logs details.
	Details []*ListInstanceLogsDetailsResponseInstanceLogDetail `json:"details"`
}

type ListInstanceLogsDetailsResponseInstanceLogDetail struct {
	LogName string `json:"log_name"`

	Size uint64 `json:"size"`
}

// ListInstanceLogsResponse: list instance logs response.
type ListInstanceLogsResponse struct {
	// InstanceLogs: available logs in a Database Instance.
	InstanceLogs []*InstanceLog `json:"instance_logs"`
}

// ListInstancesResponse: list instances response.
type ListInstancesResponse struct {
	// Instances: list of all Database Instances available in an Organization or Project.
	Instances []*Instance `json:"instances"`
	// TotalCount: total count of Database Instances available in a Organization or Project.
	TotalCount uint32 `json:"total_count"`
}

// ListNodeTypesResponse: list node types response.
type ListNodeTypesResponse struct {
	// NodeTypes: types of the node.
	NodeTypes []*NodeType `json:"node_types"`
	// TotalCount: total count of node-types available.
	TotalCount uint32 `json:"total_count"`
}

// ListPrivilegesResponse: list privileges response.
type ListPrivilegesResponse struct {
	// Privileges: privileges of a user in a database in a Database Instance.
	Privileges []*Privilege `json:"privileges"`
	// TotalCount: total count of privileges present on a database.
	TotalCount uint32 `json:"total_count"`
}

// ListSnapshotsResponse: list snapshots response.
type ListSnapshotsResponse struct {
	// Snapshots: list of snapshots.
	Snapshots []*Snapshot `json:"snapshots"`
	// TotalCount: total count of snapshots available.
	TotalCount uint32 `json:"total_count"`
}

// ListUsersResponse: list users response.
type ListUsersResponse struct {
	// Users: list of users in a Database Instance.
	Users []*User `json:"users"`
	// TotalCount: total count of users present on a Database Instance.
	TotalCount uint32 `json:"total_count"`
}

// LogsPolicy: logs policy.
type LogsPolicy struct {
	// MaxAgeRetention: max age (in days) of remote logs to keep on the Database Instance.
	MaxAgeRetention *uint32 `json:"max_age_retention"`
	// TotalDiskRetention: max disk size of remote logs to keep on the Database Instance.
	TotalDiskRetention *scw.Size `json:"total_disk_retention"`
}

// Maintenance: maintenance.
type Maintenance struct {
	// StartsAt: start date of the maintenance window.
	StartsAt *time.Time `json:"starts_at"`
	// StopsAt: end date of the maintenance window.
	StopsAt *time.Time `json:"stops_at"`
	// ClosedAt: closed maintenance date.
	ClosedAt *time.Time `json:"closed_at"`
	// Reason: maintenance information message.
	Reason string `json:"reason"`
	// Status: status of the maintenance.
	// Default value: unknown
	Status MaintenanceStatus `json:"status"`
}

// NodeType: node type.
type NodeType struct {
	// Name: node Type name identifier.
	Name string `json:"name"`
	// StockStatus: current stock status for the Node Type.
	// Default value: unknown
	StockStatus NodeTypeStock `json:"stock_status"`
	// Description: current specs of the offer.
	Description string `json:"description"`
	// Vcpus: number of virtual CPUs.
	Vcpus uint32 `json:"vcpus"`
	// Memory: quantity of RAM.
	Memory scw.Size `json:"memory"`
	// Deprecated: VolumeConstraint: [deprecated] Node Type volume constraints.
	VolumeConstraint *NodeTypeVolumeConstraintSizes `json:"volume_constraint,omitempty"`
	// Deprecated: IsBssdCompatible: the Node Type is compliant with Block Storage.
	IsBssdCompatible *bool `json:"is_bssd_compatible,omitempty"`
	// Disabled: the Node Type is currently disabled.
	Disabled bool `json:"disabled"`
	// Beta: the Node Type is currently in beta.
	Beta bool `json:"beta"`
	// AvailableVolumeTypes: available storage options for the Node Type.
	AvailableVolumeTypes []*NodeTypeVolumeType `json:"available_volume_types"`
	// IsHaRequired: the Node Type can be used only with high availability option.
	IsHaRequired bool `json:"is_ha_required"`
	// Generation: generation associated with the NodeType offer.
	// Default value: unknown_generation
	Generation NodeTypeGeneration `json:"generation"`
	// InstanceRange: instance range associated with the NodeType offer.
	InstanceRange string `json:"instance_range"`
	// Region: region the Node Type is in.
	Region scw.Region `json:"region"`
}

// NodeTypeVolumeConstraintSizes: node type. volume constraint sizes.
type NodeTypeVolumeConstraintSizes struct {
	// MinSize: [deprecated] Mimimum size required for the Volume.
	MinSize scw.Size `json:"min_size"`
	// MaxSize: [deprecated] Maximum size required for the Volume.
	MaxSize scw.Size `json:"max_size"`
}

// NodeTypeVolumeType: node type. volume type.
type NodeTypeVolumeType struct {
	// Type: volume Type.
	// Default value: lssd
	Type VolumeType `json:"type"`
	// Description: the description of the Volume.
	Description string `json:"description"`
	// MinSize: mimimum size required for the Volume.
	MinSize scw.Size `json:"min_size"`
	// MaxSize: maximum size required for the Volume.
	MaxSize scw.Size `json:"max_size"`
	// ChunkSize: minimum increment level for a Block Storage volume size.
	ChunkSize scw.Size `json:"chunk_size"`
	// Class: the storage class of the volume.
	// Default value: unknown_storage_class
	Class StorageClass `json:"class"`
}

// PrepareInstanceLogsResponse: prepare instance logs response.
type PrepareInstanceLogsResponse struct {
	// InstanceLogs: instance logs for a Database Instance between a start and an end date.
	InstanceLogs []*InstanceLog `json:"instance_logs"`
}

// Privilege: privilege.
type Privilege struct {
	// Permission: permission (Read, Read/Write, All, Custom).
	// Default value: readonly
	Permission Permission `json:"permission"`
	// DatabaseName: name of the database.
	DatabaseName string `json:"database_name"`
	// UserName: name of the user.
	UserName string `json:"user_name"`
}

// ReadReplica: read replica.
type ReadReplica struct {
	// ID: UUID of the Read Replica.
	ID string `json:"id"`
	// Endpoints: display Read Replica connection information.
	Endpoints []*Endpoint `json:"endpoints"`
	// Status: read replica status.
	// Default value: unknown
	Status ReadReplicaStatus `json:"status"`
	// Region: region the Read Replica is in.
	Region scw.Region `json:"region"`
	// SameZone: whether the replica is in the same availability zone as the main instance nodes or not.
	SameZone bool `json:"same_zone"`
}

// ReadReplicaEndpointSpec: read replica endpoint spec.
type ReadReplicaEndpointSpec struct {
	// DirectAccess: direct access endpoint specifications. Public endpoint reserved for Read Replicas. One per Read Replica.
	// Precisely one of DirectAccess, PrivateNetwork must be set.
	DirectAccess *ReadReplicaEndpointSpecDirectAccess `json:"direct_access,omitempty"`
	// PrivateNetwork: private Network endpoint specifications. One at the most per Read Replica. Cannot be updated (has to be deleted and recreated).
	// Precisely one of DirectAccess, PrivateNetwork must be set.
	PrivateNetwork *ReadReplicaEndpointSpecPrivateNetwork `json:"private_network,omitempty"`
}

type ReadReplicaEndpointSpecDirectAccess struct {
}

// ReadReplicaEndpointSpecPrivateNetwork: read replica endpoint spec. private network.
type ReadReplicaEndpointSpecPrivateNetwork struct {
	// PrivateNetworkID: UUID of the Private Network to be connected to the Read Replica.
	PrivateNetworkID string `json:"private_network_id"`
	// ServiceIP: endpoint IPv4 address with a CIDR notation. Refer to the official Scaleway documentation to learn more about IP and subnet limitations.
	// Precisely one of IpamConfig, ServiceIP must be set.
	ServiceIP *scw.IPNet `json:"service_ip,omitempty"`
	// IpamConfig: automated configuration of your Private Network endpoint with Scaleway IPAM service. One at the most per Database Instance or Read Replica (a Database Instance and its Read Replica can have different private networks). Cannot be updated (has to be deleted and recreated).
	// Precisely one of IpamConfig, ServiceIP must be set.
	IpamConfig *ReadReplicaEndpointSpecPrivateNetworkIpamConfig `json:"ipam_config,omitempty"`
}

type ReadReplicaEndpointSpecPrivateNetworkIpamConfig struct {
}

// SetInstanceACLRulesResponse: set instance acl rules response.
type SetInstanceACLRulesResponse struct {
	// Rules: aCLs rules configured for a Database Instance.
	Rules []*ACLRule `json:"rules"`
}

// SetInstanceSettingsResponse: set instance settings response.
type SetInstanceSettingsResponse struct {
	// Settings: settings configured for a Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// Snapshot: snapshot.
type Snapshot struct {
	// ID: UUID of the snapshot.
	ID string `json:"id"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"instance_id"`
	// Name: name of the snapshot.
	Name string `json:"name"`
	// Status: status of the snapshot.
	// Default value: unknown
	Status SnapshotStatus `json:"status"`
	// Size: size of the snapshot.
	Size *scw.Size `json:"size"`
	// ExpiresAt: expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at"`
	// CreatedAt: creation date (must follow the ISO 8601 format).
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: updated date (must follow the ISO 8601 format).
	UpdatedAt *time.Time `json:"updated_at"`
	// InstanceName: name of the Database Instance of the snapshot.
	InstanceName string `json:"instance_name"`
	// NodeType: source node type.
	NodeType string `json:"node_type"`
	// Region: region of this snapshot.
	Region scw.Region `json:"region"`
}

type UpgradableVersion struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Version string `json:"version"`

	MinorVersion string `json:"minor_version"`
}

// User: user.
type User struct {
	// Name: name of the user (Length must be between 1 and 63 characters for PostgreSQL and between 1 and 32 characters for MySQL. First character must be an alphabet character (a-zA-Z). Your username cannot start with '_rdb' or in PostgreSQL, 'pg_'. Only a-zA-Z0-9_$- characters are accepted).
	Name string `json:"name"`
	// IsAdmin: defines whether or not a user got administrative privileges on the Database Instance.
	IsAdmin bool `json:"is_admin"`
}

type Volume struct {
	// Type: default value: lssd
	Type VolumeType `json:"type"`

	Size scw.Size `json:"size"`
	// Class: default value: unknown_storage_class
	Class StorageClass `json:"class"`
}

// Service API

// Regions list localities the api is available in
func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

type ListDatabaseEnginesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// Name: name of the database engine.
	Name *string `json:"-"`
	// Version: version of the database engine.
	Version *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListDatabaseEngines: list available database engines.
// List the PostgreSQL and MySQL database engines available at Scaleway.
func (s *API) ListDatabaseEngines(req *ListDatabaseEnginesRequest, opts ...scw.RequestOption) (*ListDatabaseEnginesResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "version", req.Version)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/database-engines",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListDatabaseEnginesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListNodeTypesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// IncludeDisabledTypes: defines whether or not to include disabled types.
	IncludeDisabledTypes bool `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListNodeTypes: list available node types.
// List all available node types. By default, the node types returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListNodeTypes(req *ListNodeTypesRequest, opts ...scw.RequestOption) (*ListNodeTypesResponse, error) {
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
	parameter.AddToQuery(query, "include_disabled_types", req.IncludeDisabledTypes)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/node-types",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListNodeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListDatabaseBackupsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// Name: name of the database backups.
	Name *string `json:"-"`
	// OrderBy: criteria to use when ordering database backups listing.
	// Default value: created_at_asc
	OrderBy ListDatabaseBackupsRequestOrderBy `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID *string `json:"-"`
	// OrganizationID: organization ID of the Organization the database backups belong to.
	OrganizationID *string `json:"-"`
	// ProjectID: project ID of the Project the database backups belong to.
	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListDatabaseBackups: list database backups.
// List all backups in a specified region, for a given Scaleway Organization or Scaleway Project. By default, the backups listed are ordered by creation date in ascending order. This can be modified via the `order_by` field.
func (s *API) ListDatabaseBackups(req *ListDatabaseBackupsRequest, opts ...scw.RequestOption) (*ListDatabaseBackupsResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "instance_id", req.InstanceID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListDatabaseBackupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateDatabaseBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"instance_id"`
	// DatabaseName: name of the database you want to back up.
	DatabaseName string `json:"database_name"`
	// Name: name of the backup.
	Name string `json:"name"`
	// ExpiresAt: expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at"`
}

// CreateDatabaseBackup: create a database backup.
// Create a new backup. You must set the `instance_id`, `database_name`, `name` and `expires_at` parameters.
func (s *API) CreateDatabaseBackup(req *CreateDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("bkp")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups",
		Headers: http.Header{},
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

type GetDatabaseBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// DatabaseBackupID: UUID of the database backup.
	DatabaseBackupID string `json:"-"`
}

// GetDatabaseBackup: get a database backup.
// Retrieve information about a given backup, specified by its database backup ID and region. Full details about the backup, like size, URL and expiration date, are returned in the response.
func (s *API) GetDatabaseBackup(req *GetDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseBackupID) == "" {
		return nil, errors.New("field DatabaseBackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.DatabaseBackupID) + "",
		Headers: http.Header{},
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateDatabaseBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// DatabaseBackupID: UUID of the database backup to update.
	DatabaseBackupID string `json:"-"`
	// Name: name of the Database Backup.
	Name *string `json:"name"`
	// ExpiresAt: expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at"`
}

// UpdateDatabaseBackup: update a database backup.
// Update the parameters of a backup, including name and expiration date.
func (s *API) UpdateDatabaseBackup(req *UpdateDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseBackupID) == "" {
		return nil, errors.New("field DatabaseBackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.DatabaseBackupID) + "",
		Headers: http.Header{},
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

type DeleteDatabaseBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// DatabaseBackupID: UUID of the database backup to delete.
	DatabaseBackupID string `json:"-"`
}

// DeleteDatabaseBackup: delete a database backup.
// Delete a backup, specified by its database backup ID and region. Deleting a backup is permanent, and cannot be undone.
func (s *API) DeleteDatabaseBackup(req *DeleteDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseBackupID) == "" {
		return nil, errors.New("field DatabaseBackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.DatabaseBackupID) + "",
		Headers: http.Header{},
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RestoreDatabaseBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// DatabaseBackupID: backup of a logical database.
	DatabaseBackupID string `json:"-"`
	// DatabaseName: defines the destination database to restore into a specified database (the default destination is set to the origin database of the backup).
	DatabaseName *string `json:"database_name"`
	// InstanceID: defines the Database Instance where the backup has to be restored.
	InstanceID string `json:"instance_id"`
}

// RestoreDatabaseBackup: restore a database backup.
// Launch the process of restoring database backup. You must specify the `instance_id` of the Database Instance of destination, where the backup will be restored. Note that large database backups can take up to several hours to restore.
func (s *API) RestoreDatabaseBackup(req *RestoreDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseBackupID) == "" {
		return nil, errors.New("field DatabaseBackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.DatabaseBackupID) + "/restore",
		Headers: http.Header{},
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

type ExportDatabaseBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// DatabaseBackupID: UUID of the database backup you want to export.
	DatabaseBackupID string `json:"-"`
}

// ExportDatabaseBackup: export a database backup.
// Export a backup, specified by the `database_backup_id` and the `region` parameters. The download URL is returned in the response.
func (s *API) ExportDatabaseBackup(req *ExportDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseBackupID) == "" {
		return nil, errors.New("field DatabaseBackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.DatabaseBackupID) + "/export",
		Headers: http.Header{},
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

type UpgradeInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to upgrade.
	InstanceID string `json:"-"`
	// NodeType: node type of the Database Instance you want to upgrade to.
	// Precisely one of EnableHa, NodeType, UpgradableVersionID, VolumeSize, VolumeType must be set.
	NodeType *string `json:"node_type,omitempty"`
	// EnableHa: defines whether or not high availability should be enabled on the Database Instance.
	// Precisely one of EnableHa, NodeType, UpgradableVersionID, VolumeSize, VolumeType must be set.
	EnableHa *bool `json:"enable_ha,omitempty"`
	// VolumeSize: increase your block storage volume size.
	// Precisely one of EnableHa, NodeType, UpgradableVersionID, VolumeSize, VolumeType must be set.
	VolumeSize *uint64 `json:"volume_size,omitempty"`
	// VolumeType: change your Database Instance storage type.
	// Default value: lssd
	// Precisely one of EnableHa, NodeType, UpgradableVersionID, VolumeSize, VolumeType must be set.
	VolumeType *VolumeType `json:"volume_type,omitempty"`
	// UpgradableVersionID: update your database engine to a newer version.
	// This will create a new Database Instance with same specifications as the current one and perform a Database Engine upgrade.
	// Precisely one of EnableHa, NodeType, UpgradableVersionID, VolumeSize, VolumeType must be set.
	UpgradableVersionID *string `json:"upgradable_version_id,omitempty"`
}

// UpgradeInstance: upgrade a Database Instance.
// Upgrade your current Database Instance specifications like node type, high availability, volume, or the database engine version. Note that upon upgrade the `enable_ha` parameter can only be set to `true`.
func (s *API) UpgradeInstance(req *UpgradeInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/upgrade",
		Headers: http.Header{},
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

type ListInstancesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// Tags: list Database Instances that have a given tag.
	Tags []string `json:"-"`
	// Name: lists Database Instances that match a name pattern.
	Name *string `json:"-"`
	// OrderBy: criteria to use when ordering Database Instance listings.
	// Default value: created_at_asc
	OrderBy ListInstancesRequestOrderBy `json:"-"`
	// OrganizationID: please use project_id instead.
	OrganizationID *string `json:"-"`
	// ProjectID: project ID to list the Database Instance of.
	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListInstances: list Database Instances.
// List all Database Instances in the specified region, for a given Scaleway Organization or Scaleway Project. By default, the Database Instances returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `tags` and `name`. For the `name` parameter, the value you include will be checked against the whole name string to see if it includes the string you put in the parameter.
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
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListInstancesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
}

// GetInstance: get a Database Instance.
// Retrieve information about a given Database Instance, specified by the `region` and `instance_id` parameters. Its full details, including name, status, IP address and port, are returned in the response object.
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
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
		Headers: http.Header{},
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// Deprecated: OrganizationID: please use project_id instead.
	// Precisely one of OrganizationID, ProjectID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`
	// ProjectID: the Project ID on which the Database Instance will be created.
	// Precisely one of OrganizationID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`
	// Name: name of the Database Instance.
	Name string `json:"name"`
	// Engine: database engine of the Database Instance (PostgreSQL, MySQL, ...).
	Engine string `json:"engine"`
	// UserName: username created when the Database Instance is created.
	UserName string `json:"user_name"`
	// Password: password of the user.
	Password string `json:"password"`
	// NodeType: type of node to use for the Database Instance.
	NodeType string `json:"node_type"`
	// IsHaCluster: defines whether or not High-Availability is enabled.
	IsHaCluster bool `json:"is_ha_cluster"`
	// DisableBackup: defines whether or not backups are disabled.
	DisableBackup bool `json:"disable_backup"`
	// Tags: tags to apply to the Database Instance.
	Tags []string `json:"tags"`
	// InitSettings: list of engine settings to be set upon Database Instance initialization.
	InitSettings []*InstanceSetting `json:"init_settings"`
	// VolumeType: type of volume where data is stored (lssd, bssd, ...).
	// Default value: lssd
	VolumeType VolumeType `json:"volume_type"`
	// VolumeSize: volume size when volume_type is not lssd.
	VolumeSize scw.Size `json:"volume_size"`
	// InitEndpoints: one or multiple EndpointSpec used to expose your Database Instance. A load_balancer public endpoint is systematically created.
	InitEndpoints []*EndpointSpec `json:"init_endpoints"`
	// BackupSameRegion: defines whether to or not to store logical backups in the same region as the Database Instance.
	BackupSameRegion bool `json:"backup_same_region"`
}

// CreateInstance: create a Database Instance.
// Create a new Database Instance. You must set the `engine`, `user_name`, `password` and `node_type` parameters. Optionally, you can specify the volume type and size.
func (s *API) CreateInstance(req *CreateInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		req.Name = namegenerator.GetRandomName("ins")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances",
		Headers: http.Header{},
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

type UpdateInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to update.
	InstanceID string `json:"-"`
	// BackupScheduleFrequency: in hours.
	BackupScheduleFrequency *uint32 `json:"backup_schedule_frequency"`
	// BackupScheduleRetention: in days.
	BackupScheduleRetention *uint32 `json:"backup_schedule_retention"`
	// IsBackupScheduleDisabled: defines whether or not the backup schedule is disabled.
	IsBackupScheduleDisabled *bool `json:"is_backup_schedule_disabled"`
	// Name: name of the Database Instance.
	Name *string `json:"name"`
	// Tags: tags of a Database Instance.
	Tags *[]string `json:"tags"`
	// LogsPolicy: logs policy of the Database Instance.
	LogsPolicy *LogsPolicy `json:"logs_policy"`
	// BackupSameRegion: store logical backups in the same region as the Database Instance.
	BackupSameRegion *bool `json:"backup_same_region"`
	// BackupScheduleStartHour: defines the start time of the autobackup.
	BackupScheduleStartHour *uint32 `json:"backup_schedule_start_hour"`
}

// UpdateInstance: update a Database Instance.
// Update the parameters of a Database Instance, including name, tags and backup schedule details.
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
		Method:  "PATCH",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
		Headers: http.Header{},
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

type DeleteInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to delete.
	InstanceID string `json:"-"`
}

// DeleteInstance: delete a Database Instance.
// Delete a given Database Instance, specified by the `region` and `instance_id` parameters. Deleting a Database Instance is permanent, and cannot be undone. Note that upon deletion all your data will be lost.
func (s *API) DeleteInstance(req *DeleteInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Method:  "DELETE",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
		Headers: http.Header{},
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CloneInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to clone.
	InstanceID string `json:"-"`
	// Name: name of the Database Instance clone.
	Name string `json:"name"`
	// NodeType: node type of the clone.
	NodeType *string `json:"node_type"`
}

// CloneInstance: clone a Database Instance.
// Clone a given Database Instance, specified by the `region` and `instance_id` parameters. The clone feature allows you to create a new Database Instance from an existing one. The clone includes all existing databases, users and permissions. You can create a clone on a Database Instance bigger than your current one.
func (s *API) CloneInstance(req *CloneInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/clone",
		Headers: http.Header{},
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

type RestartInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to restart.
	InstanceID string `json:"-"`
}

// RestartInstance: restart Database Instance.
// Restart a given Database Instance, specified by the `region` and `instance_id` parameters. The status of the Database Instance returned in the response.
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/restart",
		Headers: http.Header{},
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

type GetInstanceCertificateRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
}

// GetInstanceCertificate: get the TLS certificate of a Database Instance.
// Retrieve information about the TLS certificate of a given Database Instance. Details like name and content are returned in the response.
func (s *API) GetInstanceCertificate(req *GetInstanceCertificateRequest, opts ...scw.RequestOption) (*scw.File, error) {
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
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/certificate",
		Headers: http.Header{},
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RenewInstanceCertificateRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
}

// RenewInstanceCertificate: renew the TLS certificate of a Database Instance.
// Renew a TLS for a Database Instance. Renewing a certificate means that you will not be able to connect to your Database Instance using the previous certificate. You will also need to download and update the new certificate for all database clients.
func (s *API) RenewInstanceCertificate(req *RenewInstanceCertificateRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/renew-certificate",
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

type GetInstanceMetricsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// StartDate: start date to gather metrics from.
	StartDate *time.Time `json:"-"`
	// EndDate: end date to gather metrics from.
	EndDate *time.Time `json:"-"`
	// MetricName: name of the metric to gather.
	MetricName *string `json:"-"`
}

// GetInstanceMetrics: get Database Instance metrics.
// Retrieve the time series metrics of a given Database Instance. You can define the period from which to retrieve metrics by specifying the `start_date` and `end_date`.
func (s *API) GetInstanceMetrics(req *GetInstanceMetricsRequest, opts ...scw.RequestOption) (*InstanceMetrics, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)
	parameter.AddToQuery(query, "metric_name", req.MetricName)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/metrics",
		Query:   query,
		Headers: http.Header{},
	}

	var resp InstanceMetrics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateReadReplicaRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to create a Read Replica from.
	InstanceID string `json:"instance_id"`
	// EndpointSpec: specification of the endpoint you want to create.
	EndpointSpec []*ReadReplicaEndpointSpec `json:"endpoint_spec"`
	// SameZone: defines whether to create the replica in the same availability zone as the main instance nodes or not.
	SameZone *bool `json:"same_zone"`
}

// CreateReadReplica: create a Read Replica.
// Create a new Read Replica of a Database Instance. You must specify the `region` and the `instance_id`. You can only create a maximum of 3 Read Replicas per Database Instance.
func (s *API) CreateReadReplica(req *CreateReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
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
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas",
		Headers: http.Header{},
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

type GetReadReplicaRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
}

// GetReadReplica: get a Read Replica.
// Retrieve information about a Database Instance Read Replica. Full details about the Read Replica, like `endpoints`, `status`  and `region` are returned in the response.
func (s *API) GetReadReplica(req *GetReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
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
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "",
		Headers: http.Header{},
	}

	var resp ReadReplica

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteReadReplicaRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
}

// DeleteReadReplica: delete a Read Replica.
// Delete a Read Replica of a Database Instance. You must specify the `region` and `read_replica_id` parameters of the Read Replica you want to delete.
func (s *API) DeleteReadReplica(req *DeleteReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
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
		Method:  "DELETE",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "",
		Headers: http.Header{},
	}

	var resp ReadReplica

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ResetReadReplicaRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
}

// ResetReadReplica: resync a Read Replica.
// When you resync a Read Replica, first it is reset, then its data is resynchronized from the primary node. Your Read Replica remains unavailable during the resync process. The duration of this process is proportional to the size of your Database Instance.
// The configured endpoints do not change.
func (s *API) ResetReadReplica(req *ResetReadReplicaRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "/reset",
		Headers: http.Header{},
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

type PromoteReadReplicaRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
}

// PromoteReadReplica: promote a Read Replica.
// Promote a Read Replica to Database Instance automatically.
func (s *API) PromoteReadReplica(req *PromoteReadReplicaRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "/promote",
		Headers: http.Header{},
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

type CreateReadReplicaEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// ReadReplicaID: UUID of the Read Replica.
	ReadReplicaID string `json:"-"`
	// EndpointSpec: specification of the endpoint you want to create.
	EndpointSpec []*ReadReplicaEndpointSpec `json:"endpoint_spec"`
}

// CreateReadReplicaEndpoint: create an endpoint for a Read Replica.
// Create a new endpoint for a Read Replica. Read Replicas can have at most one direct access and one Private Network endpoint.
func (s *API) CreateReadReplicaEndpoint(req *CreateReadReplicaEndpointRequest, opts ...scw.RequestOption) (*ReadReplica, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/read-replicas/" + fmt.Sprint(req.ReadReplicaID) + "/endpoints",
		Headers: http.Header{},
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

type PrepareInstanceLogsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
	// StartDate: start datetime of your log. (RFC 3339 format).
	StartDate *time.Time `json:"start_date"`
	// EndDate: end datetime of your log. (RFC 3339 format).
	EndDate *time.Time `json:"end_date"`
}

// PrepareInstanceLogs: prepare logs of a Database Instance.
// Prepare your Database Instance logs. You can define the `start_date` and `end_date` parameters for your query. The download URL is returned in the response. Logs are recorded from 00h00 to 23h59 and then aggregated in a `.log` file once a day. Therefore, even if you specify a timeframe from which you want to get the logs, you will receive logs from the full 24 hours.
func (s *API) PrepareInstanceLogs(req *PrepareInstanceLogsRequest, opts ...scw.RequestOption) (*PrepareInstanceLogsResponse, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/prepare-logs",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrepareInstanceLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListInstanceLogsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
	// OrderBy: criteria to use when ordering Database Instance logs listing.
	// Default value: created_at_asc
	OrderBy ListInstanceLogsRequestOrderBy `json:"-"`
}

// ListInstanceLogs: list available logs of a Database Instance.
// List the available logs of a Database Instance. By default, the logs returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.
func (s *API) ListInstanceLogs(req *ListInstanceLogsRequest, opts ...scw.RequestOption) (*ListInstanceLogsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/logs",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListInstanceLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetInstanceLogRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceLogID: UUID of the instance_log you want.
	InstanceLogID string `json:"-"`
}

// GetInstanceLog: get given logs of a Database Instance.
// Retrieve information about the logs of a Database Instance. Specify the `instance_log_id` and `region` in your request to get information such as `download_url`, `status`, `expires_at` and `created_at` about your logs in the response.
func (s *API) GetInstanceLog(req *GetInstanceLogRequest, opts ...scw.RequestOption) (*InstanceLog, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceLogID) == "" {
		return nil, errors.New("field InstanceLogID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/logs/" + fmt.Sprint(req.InstanceLogID) + "",
		Headers: http.Header{},
	}

	var resp InstanceLog

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type PurgeInstanceLogsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
	// LogName: given log name to purge.
	LogName *string `json:"log_name"`
}

// PurgeInstanceLogs: purge remote Database Instance logs.
// Purge a given remote log from a Database Instance. You can specify the `log_name` of the log you wish to clean from your Database Instance.
func (s *API) PurgeInstanceLogs(req *PurgeInstanceLogsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/purge-logs",
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

type ListInstanceLogsDetailsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want logs of.
	InstanceID string `json:"-"`
}

// ListInstanceLogsDetails: list remote Database Instance logs details.
// List remote log details. By default, the details returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.
func (s *API) ListInstanceLogsDetails(req *ListInstanceLogsDetailsRequest, opts ...scw.RequestOption) (*ListInstanceLogsDetailsResponse, error) {
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
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/logs-details",
		Headers: http.Header{},
	}

	var resp ListInstanceLogsDetailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddInstanceSettingsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to add settings to.
	InstanceID string `json:"-"`
	// Settings: settings to add to the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// AddInstanceSettings: add Database Instance advanced settings.
// Add an advanced setting to a Database Instance. You must set the `name` and the `value` of each setting.
func (s *API) AddInstanceSettings(req *AddInstanceSettingsRequest, opts ...scw.RequestOption) (*AddInstanceSettingsResponse, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/settings",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddInstanceSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteInstanceSettingsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to delete settings from.
	InstanceID string `json:"-"`
	// SettingNames: settings names to delete.
	SettingNames []string `json:"setting_names"`
}

// DeleteInstanceSettings: delete Database Instance advanced settings.
// Delete an advanced setting in a Database Instance. You must specify the names of the settings you want to delete in the request.
func (s *API) DeleteInstanceSettings(req *DeleteInstanceSettingsRequest, opts ...scw.RequestOption) (*DeleteInstanceSettingsResponse, error) {
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
		Method:  "DELETE",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/settings",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DeleteInstanceSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetInstanceSettingsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance where the settings must be set.
	InstanceID string `json:"-"`
	// Settings: settings to define for the Database Instance.
	Settings []*InstanceSetting `json:"settings"`
}

// SetInstanceSettings: set Database Instance advanced settings.
// Update an advanced setting for a Database Instance. Settings added upon database engine initalization can only be defined once, and cannot, therefore, be updated.
func (s *API) SetInstanceSettings(req *SetInstanceSettingsRequest, opts ...scw.RequestOption) (*SetInstanceSettingsResponse, error) {
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
		Method:  "PUT",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/settings",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetInstanceSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListInstanceACLRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListInstanceACLRules: list ACL rules of a Database Instance.
// List the ACL rules for a given Database Instance. The response is an array of ACL objects, each one representing an ACL that denies, allows or redirects traffic based on certain conditions.
func (s *API) ListInstanceACLRules(req *ListInstanceACLRulesRequest, opts ...scw.RequestOption) (*ListInstanceACLRulesResponse, error) {
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

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/acls",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListInstanceACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddInstanceACLRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to add ACL rules to.
	InstanceID string `json:"-"`
	// Rules: ACL rules to add to the Database Instance.
	Rules []*ACLRuleRequest `json:"rules"`
}

// AddInstanceACLRules: add an ACL rule to a Database Instance.
// Add an additional ACL rule to a Database Instance.
func (s *API) AddInstanceACLRules(req *AddInstanceACLRulesRequest, opts ...scw.RequestOption) (*AddInstanceACLRulesResponse, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/acls",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddInstanceACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetInstanceACLRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance where the ACL rules must be set.
	InstanceID string `json:"-"`
	// Rules: ACL rules to define for the Database Instance.
	Rules []*ACLRuleRequest `json:"rules"`
}

// SetInstanceACLRules: set ACL rules for a Database Instance.
// Replace all the ACL rules of a Database Instance.
func (s *API) SetInstanceACLRules(req *SetInstanceACLRulesRequest, opts ...scw.RequestOption) (*SetInstanceACLRulesResponse, error) {
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
		Method:  "PUT",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/acls",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetInstanceACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteInstanceACLRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you want to delete an ACL rule from.
	InstanceID string `json:"-"`
	// ACLRuleIPs: IP addresses defined in the ACL rules of the Database Instance.
	ACLRuleIPs []string `json:"acl_rule_ips"`
}

// DeleteInstanceACLRules: delete ACL rules of a Database Instance.
// Delete one or more ACL rules of a Database Instance.
func (s *API) DeleteInstanceACLRules(req *DeleteInstanceACLRulesRequest, opts ...scw.RequestOption) (*DeleteInstanceACLRulesResponse, error) {
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
		Method:  "DELETE",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/acls",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DeleteInstanceACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListUsersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// Name: name of the user.
	Name *string `json:"-"`
	// OrderBy: criteria to use when requesting user listing.
	// Default value: name_asc
	OrderBy ListUsersRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListUsers: list users of a Database Instance.
// List all users of a given Database Instance. By default, the users returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.
func (s *API) ListUsers(req *ListUsersRequest, opts ...scw.RequestOption) (*ListUsersResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance in which you want to create a user.
	InstanceID string `json:"-"`
	// Name: name of the user you want to create.
	Name string `json:"name"`
	// Password: password of the user you want to create.
	Password string `json:"password"`
	// IsAdmin: defines whether the user will have administrative privileges.
	IsAdmin bool `json:"is_admin"`
}

// CreateUser: create a user for a Database Instance.
// Create a new user for a Database Instance. You must define the `name`, `password` and `is_admin` parameters.
func (s *API) CreateUser(req *CreateUserRequest, opts ...scw.RequestOption) (*User, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance the user belongs to.
	InstanceID string `json:"-"`
	// Name: name of the database user.
	Name string `json:"-"`
	// Password: password of the database user.
	Password *string `json:"password"`
	// IsAdmin: defines whether or not this user got administrative privileges.
	IsAdmin *bool `json:"is_admin"`
}

// UpdateUser: update a user on a Database Instance.
// Update the parameters of a user on a Database Instance. You can update the `password` and `is_admin` parameters, but you cannot change the name of the user.
func (s *API) UpdateUser(req *UpdateUserRequest, opts ...scw.RequestOption) (*User, error) {
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

	if fmt.Sprint(req.Name) == "" {
		return nil, errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users/" + fmt.Sprint(req.Name) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to delete the user from.
	InstanceID string `json:"-"`
	// Name: name of the user.
	Name string `json:"-"`
}

// DeleteUser: delete a user on a Database Instance.
// Delete a given user on a Database Instance. You must specify, in the endpoint,  the `region`, `instance_id` and `name` parameters of the user you want to delete.
func (s *API) DeleteUser(req *DeleteUserRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return errors.New("field InstanceID cannot be empty in request")
	}

	if fmt.Sprint(req.Name) == "" {
		return errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users/" + fmt.Sprint(req.Name) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListDatabasesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance to list the databases of.
	InstanceID string `json:"-"`
	// Name: name of the database.
	Name *string `json:"-"`
	// Managed: defines whether or not the database is managed.
	Managed *bool `json:"-"`
	// Owner: user that owns this database.
	Owner *string `json:"-"`
	// OrderBy: criteria to use when ordering database listing.
	// Default value: name_asc
	OrderBy ListDatabasesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListDatabases: list databases in a Database Instance.
// List all databases of a given Database Instance. By default, the databases returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `name`, `managed` and `owner`.
func (s *API) ListDatabases(req *ListDatabasesRequest, opts ...scw.RequestOption) (*ListDatabasesResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "managed", req.Managed)
	parameter.AddToQuery(query, "owner", req.Owner)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/databases",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListDatabasesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateDatabaseRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance where to create the database.
	InstanceID string `json:"-"`
	// Name: name of the database.
	Name string `json:"name"`
}

// CreateDatabase: create a database in a Database Instance.
// Create a new database. You must define the `name` parameter in the request.
func (s *API) CreateDatabase(req *CreateDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/databases",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Database

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteDatabaseRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance where to delete the database.
	InstanceID string `json:"-"`
	// Name: name of the database to delete.
	Name string `json:"-"`
}

// DeleteDatabase: delete a database in a Database Instance.
// Delete a given database on a Database Instance. You must specify, in the endpoint, the `region`, `instance_id` and `name` parameters of the database you want to delete.
func (s *API) DeleteDatabase(req *DeleteDatabaseRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return errors.New("field InstanceID cannot be empty in request")
	}

	if fmt.Sprint(req.Name) == "" {
		return errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/databases/" + fmt.Sprint(req.Name) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListPrivilegesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// OrderBy: criteria to use when ordering privileges listing.
	// Default value: user_name_asc
	OrderBy ListPrivilegesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// DatabaseName: name of the database.
	DatabaseName *string `json:"-"`
	// UserName: name of the user.
	UserName *string `json:"-"`
}

// ListPrivileges: list user privileges for a database.
// List privileges of a user on a database. By default, the details returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `database_name` and `user_name`.
func (s *API) ListPrivileges(req *ListPrivilegesRequest, opts ...scw.RequestOption) (*ListPrivilegesResponse, error) {
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
	parameter.AddToQuery(query, "database_name", req.DatabaseName)
	parameter.AddToQuery(query, "user_name", req.UserName)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/privileges",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListPrivilegesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetPrivilegeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// DatabaseName: name of the database.
	DatabaseName string `json:"database_name"`
	// UserName: name of the user.
	UserName string `json:"user_name"`
	// Permission: permission to set (Read, Read/Write, All, Custom).
	// Default value: readonly
	Permission Permission `json:"permission"`
}

// SetPrivilege: set user privileges for a database.
// Set the privileges of a user on a database. You must define `database_name`, `user_name` and `permission` in the request body.
func (s *API) SetPrivilege(req *SetPrivilegeRequest, opts ...scw.RequestOption) (*Privilege, error) {
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
		Method:  "PUT",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/privileges",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Privilege

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListSnapshotsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// Name: name of the snapshot.
	Name *string `json:"-"`
	// OrderBy: criteria to use when ordering snapshot listing.
	// Default value: created_at_asc
	OrderBy ListSnapshotsRequestOrderBy `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID *string `json:"-"`
	// OrganizationID: organization ID the snapshots belongs to.
	OrganizationID *string `json:"-"`
	// ProjectID: project ID the snapshots belongs to.
	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListSnapshots: list snapshots.
// List snapshots. You can include the `instance_id` or `project_id` in your query to get the list of snapshots for specific Database Instances and/or Projects. By default, the details returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListSnapshots(req *ListSnapshotsRequest, opts ...scw.RequestOption) (*ListSnapshotsResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "instance_id", req.InstanceID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/snapshots",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListSnapshotsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetSnapshotRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SnapshotID: UUID of the snapshot.
	SnapshotID string `json:"-"`
}

// GetSnapshot: get a Database Instance snapshot.
// Retrieve information about a given snapshot, specified by its `snapshot_id` and `region`. Full details about the snapshot, like size and expiration date, are returned in the response.
func (s *API) GetSnapshot(req *GetSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
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
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
		Headers: http.Header{},
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateSnapshotRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
	// Name: name of the snapshot.
	Name string `json:"name"`
	// ExpiresAt: expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at"`
}

// CreateSnapshot: create a Database Instance snapshot.
// Create a new snapshot of a Database Instance. You must define the `name` parameter in the request.
func (s *API) CreateSnapshot(req *CreateSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("snp")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.InstanceID) == "" {
		return nil, errors.New("field InstanceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/snapshots",
		Headers: http.Header{},
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

type UpdateSnapshotRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SnapshotID: UUID of the snapshot to update.
	SnapshotID string `json:"-"`
	// Name: name of the snapshot.
	Name *string `json:"name"`
	// ExpiresAt: expiration date (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at"`
}

// UpdateSnapshot: update a Database Instance snapshot.
// Update the parameters of a snapshot of a Database Instance. You can update the `name` and `expires_at` parameters.
func (s *API) UpdateSnapshot(req *UpdateSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
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
		Method:  "PATCH",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
		Headers: http.Header{},
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

type DeleteSnapshotRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SnapshotID: UUID of the snapshot to delete.
	SnapshotID string `json:"-"`
}

// DeleteSnapshot: delete a Database Instance snapshot.
// Delete a given snapshot of a Database Instance. You must specify, in the endpoint,  the `region` and `snapshot_id` parameters of the snapshot you want to delete.
func (s *API) DeleteSnapshot(req *DeleteSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
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
		Method:  "DELETE",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
		Headers: http.Header{},
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateInstanceFromSnapshotRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SnapshotID: block snapshot of the Database Instance.
	SnapshotID string `json:"-"`
	// InstanceName: name of the Database Instance created with the snapshot.
	InstanceName string `json:"instance_name"`
	// IsHaCluster: defines whether or not High-Availability is enabled on the new Database Instance.
	IsHaCluster *bool `json:"is_ha_cluster"`
	// NodeType: the node type used to restore the snapshot.
	NodeType *string `json:"node_type"`
}

// CreateInstanceFromSnapshot: create a new Database Instance from a snapshot.
// Restore a snapshot. When you restore a snapshot, a new Instance is created and billed to your account. Note that is possible to select a larger node type for your new Database Instance. However, the Block volume size will be the same as the size of the restored snapshot. All Instance settings will be restored if you chose a node type with the same or more memory size than the initial Instance. Settings will be reset to the default if your node type has less memory.
func (s *API) CreateInstanceFromSnapshot(req *CreateInstanceFromSnapshotRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "/create-instance",
		Headers: http.Header{},
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

type CreateEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the Database Instance you to which you want to add an endpoint.
	InstanceID string `json:"-"`
	// EndpointSpec: specification of the endpoint you want to create.
	EndpointSpec *EndpointSpec `json:"endpoint_spec"`
}

// CreateEndpoint: create a new Database Instance endpoint.
// Create a new endpoint for a Database Instance. You can add `load_balancer` and `private_network` specifications to the body of the request.
func (s *API) CreateEndpoint(req *CreateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/endpoints",
		Headers: http.Header{},
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

type DeleteEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// EndpointID: UUID of the endpoint you want to delete.
	// This endpoint can also be used to delete a Read Replica endpoint.
	EndpointID string `json:"-"`
}

// DeleteEndpoint: delete a Database Instance endpoint.
// Delete the endpoint of a Database Instance. You must specify the `region` and `endpoint_id` parameters of the endpoint you want to delete. Note that might need to update any environment configurations that point to the deleted endpoint.
func (s *API) DeleteEndpoint(req *DeleteEndpointRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// EndpointID: UUID of the endpoint you want to get.
	EndpointID string `json:"-"`
}

// GetEndpoint: get a Database Instance endpoint.
// Retrieve information about a Database Instance endpoint. Full details about the endpoint, like `ip`, `port`, `private_network` and `load_balancer` specifications are returned in the response.
func (s *API) GetEndpoint(req *GetEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
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
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
		Headers: http.Header{},
	}

	var resp Endpoint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type MigrateEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// EndpointID: UUID of the endpoint you want to migrate.
	EndpointID string `json:"-"`
	// InstanceID: UUID of the instance you want to attach the endpoint to.
	InstanceID string `json:"instance_id"`
}

// MigrateEndpoint: migrate an existing instance endpoint to another instance.
func (s *API) MigrateEndpoint(req *MigrateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
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
		Method:  "POST",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "/migrate",
		Headers: http.Header{},
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

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatabaseEnginesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatabaseEnginesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDatabaseEnginesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Engines = append(r.Engines, results.Engines...)
	r.TotalCount += uint32(len(results.Engines))
	return uint32(len(results.Engines)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNodeTypesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNodeTypesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNodeTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.NodeTypes = append(r.NodeTypes, results.NodeTypes...)
	r.TotalCount += uint32(len(results.NodeTypes))
	return uint32(len(results.NodeTypes)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatabaseBackupsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatabaseBackupsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDatabaseBackupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.DatabaseBackups = append(r.DatabaseBackups, results.DatabaseBackups...)
	r.TotalCount += uint32(len(results.DatabaseBackups))
	return uint32(len(results.DatabaseBackups)), nil
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

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInstanceACLRulesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInstanceACLRulesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListInstanceACLRulesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Rules = append(r.Rules, results.Rules...)
	r.TotalCount += uint32(len(results.Rules))
	return uint32(len(results.Rules)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Users = append(r.Users, results.Users...)
	r.TotalCount += uint32(len(results.Users))
	return uint32(len(results.Users)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatabasesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatabasesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDatabasesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Databases = append(r.Databases, results.Databases...)
	r.TotalCount += uint32(len(results.Databases))
	return uint32(len(results.Databases)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPrivilegesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPrivilegesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPrivilegesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Privileges = append(r.Privileges, results.Privileges...)
	r.TotalCount += uint32(len(results.Privileges))
	return uint32(len(results.Privileges)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSnapshotsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Snapshots = append(r.Snapshots, results.Snapshots...)
	r.TotalCount += uint32(len(results.Snapshots))
	return uint32(len(results.Snapshots)), nil
}

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

// API: database RDB API
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
	// ACLRuleActionAllow is [insert doc].
	ACLRuleActionAllow = ACLRuleAction("allow")
	// ACLRuleActionDeny is [insert doc].
	ACLRuleActionDeny = ACLRuleAction("deny")
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
	// ACLRuleDirectionInbound is [insert doc].
	ACLRuleDirectionInbound = ACLRuleDirection("inbound")
	// ACLRuleDirectionOutbound is [insert doc].
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
	// ACLRuleProtocolTCP is [insert doc].
	ACLRuleProtocolTCP = ACLRuleProtocol("tcp")
	// ACLRuleProtocolUDP is [insert doc].
	ACLRuleProtocolUDP = ACLRuleProtocol("udp")
	// ACLRuleProtocolIcmp is [insert doc].
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
	// DatabaseBackupStatusUnknown is [insert doc].
	DatabaseBackupStatusUnknown = DatabaseBackupStatus("unknown")
	// DatabaseBackupStatusCreating is [insert doc].
	DatabaseBackupStatusCreating = DatabaseBackupStatus("creating")
	// DatabaseBackupStatusReady is [insert doc].
	DatabaseBackupStatusReady = DatabaseBackupStatus("ready")
	// DatabaseBackupStatusRestoring is [insert doc].
	DatabaseBackupStatusRestoring = DatabaseBackupStatus("restoring")
	// DatabaseBackupStatusDeleting is [insert doc].
	DatabaseBackupStatusDeleting = DatabaseBackupStatus("deleting")
	// DatabaseBackupStatusError is [insert doc].
	DatabaseBackupStatusError = DatabaseBackupStatus("error")
	// DatabaseBackupStatusExporting is [insert doc].
	DatabaseBackupStatusExporting = DatabaseBackupStatus("exporting")
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
	// EngineSettingPropertyTypeBOOLEAN is [insert doc].
	EngineSettingPropertyTypeBOOLEAN = EngineSettingPropertyType("BOOLEAN")
	// EngineSettingPropertyTypeINT is [insert doc].
	EngineSettingPropertyTypeINT = EngineSettingPropertyType("INT")
	// EngineSettingPropertyTypeSTRING is [insert doc].
	EngineSettingPropertyTypeSTRING = EngineSettingPropertyType("STRING")
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
	// InstanceLogStatusUnknown is [insert doc].
	InstanceLogStatusUnknown = InstanceLogStatus("unknown")
	// InstanceLogStatusReady is [insert doc].
	InstanceLogStatusReady = InstanceLogStatus("ready")
	// InstanceLogStatusCreating is [insert doc].
	InstanceLogStatusCreating = InstanceLogStatus("creating")
	// InstanceLogStatusError is [insert doc].
	InstanceLogStatusError = InstanceLogStatus("error")
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
	// InstanceStatusUnknown is [insert doc].
	InstanceStatusUnknown = InstanceStatus("unknown")
	// InstanceStatusReady is [insert doc].
	InstanceStatusReady = InstanceStatus("ready")
	// InstanceStatusProvisioning is [insert doc].
	InstanceStatusProvisioning = InstanceStatus("provisioning")
	// InstanceStatusConfiguring is [insert doc].
	InstanceStatusConfiguring = InstanceStatus("configuring")
	// InstanceStatusDeleting is [insert doc].
	InstanceStatusDeleting = InstanceStatus("deleting")
	// InstanceStatusError is [insert doc].
	InstanceStatusError = InstanceStatus("error")
	// InstanceStatusAutohealing is [insert doc].
	InstanceStatusAutohealing = InstanceStatus("autohealing")
	// InstanceStatusLocked is [insert doc].
	InstanceStatusLocked = InstanceStatus("locked")
	// InstanceStatusInitializing is [insert doc].
	InstanceStatusInitializing = InstanceStatus("initializing")
	// InstanceStatusDiskFull is [insert doc].
	InstanceStatusDiskFull = InstanceStatus("disk_full")
	// InstanceStatusBackuping is [insert doc].
	InstanceStatusBackuping = InstanceStatus("backuping")
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
	// ListDatabaseBackupsRequestOrderByCreatedAtAsc is [insert doc].
	ListDatabaseBackupsRequestOrderByCreatedAtAsc = ListDatabaseBackupsRequestOrderBy("created_at_asc")
	// ListDatabaseBackupsRequestOrderByCreatedAtDesc is [insert doc].
	ListDatabaseBackupsRequestOrderByCreatedAtDesc = ListDatabaseBackupsRequestOrderBy("created_at_desc")
	// ListDatabaseBackupsRequestOrderByNameAsc is [insert doc].
	ListDatabaseBackupsRequestOrderByNameAsc = ListDatabaseBackupsRequestOrderBy("name_asc")
	// ListDatabaseBackupsRequestOrderByNameDesc is [insert doc].
	ListDatabaseBackupsRequestOrderByNameDesc = ListDatabaseBackupsRequestOrderBy("name_desc")
	// ListDatabaseBackupsRequestOrderByStatusAsc is [insert doc].
	ListDatabaseBackupsRequestOrderByStatusAsc = ListDatabaseBackupsRequestOrderBy("status_asc")
	// ListDatabaseBackupsRequestOrderByStatusDesc is [insert doc].
	ListDatabaseBackupsRequestOrderByStatusDesc = ListDatabaseBackupsRequestOrderBy("status_desc")
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
	// ListDatabasesRequestOrderByNameAsc is [insert doc].
	ListDatabasesRequestOrderByNameAsc = ListDatabasesRequestOrderBy("name_asc")
	// ListDatabasesRequestOrderByNameDesc is [insert doc].
	ListDatabasesRequestOrderByNameDesc = ListDatabasesRequestOrderBy("name_desc")
	// ListDatabasesRequestOrderBySizeAsc is [insert doc].
	ListDatabasesRequestOrderBySizeAsc = ListDatabasesRequestOrderBy("size_asc")
	// ListDatabasesRequestOrderBySizeDesc is [insert doc].
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
	// ListInstanceLogsRequestOrderByCreatedAtAsc is [insert doc].
	ListInstanceLogsRequestOrderByCreatedAtAsc = ListInstanceLogsRequestOrderBy("created_at_asc")
	// ListInstanceLogsRequestOrderByCreatedAtDesc is [insert doc].
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
	// ListInstancesRequestOrderByCreatedAtAsc is [insert doc].
	ListInstancesRequestOrderByCreatedAtAsc = ListInstancesRequestOrderBy("created_at_asc")
	// ListInstancesRequestOrderByCreatedAtDesc is [insert doc].
	ListInstancesRequestOrderByCreatedAtDesc = ListInstancesRequestOrderBy("created_at_desc")
	// ListInstancesRequestOrderByNameAsc is [insert doc].
	ListInstancesRequestOrderByNameAsc = ListInstancesRequestOrderBy("name_asc")
	// ListInstancesRequestOrderByNameDesc is [insert doc].
	ListInstancesRequestOrderByNameDesc = ListInstancesRequestOrderBy("name_desc")
	// ListInstancesRequestOrderByRegion is [insert doc].
	ListInstancesRequestOrderByRegion = ListInstancesRequestOrderBy("region")
	// ListInstancesRequestOrderByStatusAsc is [insert doc].
	ListInstancesRequestOrderByStatusAsc = ListInstancesRequestOrderBy("status_asc")
	// ListInstancesRequestOrderByStatusDesc is [insert doc].
	ListInstancesRequestOrderByStatusDesc = ListInstancesRequestOrderBy("status_desc")
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
	// ListPrivilegesRequestOrderByUserNameAsc is [insert doc].
	ListPrivilegesRequestOrderByUserNameAsc = ListPrivilegesRequestOrderBy("user_name_asc")
	// ListPrivilegesRequestOrderByUserNameDesc is [insert doc].
	ListPrivilegesRequestOrderByUserNameDesc = ListPrivilegesRequestOrderBy("user_name_desc")
	// ListPrivilegesRequestOrderByDatabaseNameAsc is [insert doc].
	ListPrivilegesRequestOrderByDatabaseNameAsc = ListPrivilegesRequestOrderBy("database_name_asc")
	// ListPrivilegesRequestOrderByDatabaseNameDesc is [insert doc].
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

type ListUsersRequestOrderBy string

const (
	// ListUsersRequestOrderByNameAsc is [insert doc].
	ListUsersRequestOrderByNameAsc = ListUsersRequestOrderBy("name_asc")
	// ListUsersRequestOrderByNameDesc is [insert doc].
	ListUsersRequestOrderByNameDesc = ListUsersRequestOrderBy("name_desc")
	// ListUsersRequestOrderByIsAdminAsc is [insert doc].
	ListUsersRequestOrderByIsAdminAsc = ListUsersRequestOrderBy("is_admin_asc")
	// ListUsersRequestOrderByIsAdminDesc is [insert doc].
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

type NodeTypeStock string

const (
	// NodeTypeStockUnknown is [insert doc].
	NodeTypeStockUnknown = NodeTypeStock("unknown")
	// NodeTypeStockLowStock is [insert doc].
	NodeTypeStockLowStock = NodeTypeStock("low_stock")
	// NodeTypeStockOutOfStock is [insert doc].
	NodeTypeStockOutOfStock = NodeTypeStock("out_of_stock")
	// NodeTypeStockAvailable is [insert doc].
	NodeTypeStockAvailable = NodeTypeStock("available")
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
	// PermissionReadonly is [insert doc].
	PermissionReadonly = Permission("readonly")
	// PermissionReadwrite is [insert doc].
	PermissionReadwrite = Permission("readwrite")
	// PermissionAll is [insert doc].
	PermissionAll = Permission("all")
	// PermissionCustom is [insert doc].
	PermissionCustom = Permission("custom")
	// PermissionNone is [insert doc].
	PermissionNone = Permission("none")
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

type VolumeType string

const (
	// VolumeTypeLssd is [insert doc].
	VolumeTypeLssd = VolumeType("lssd")
	// VolumeTypeBssd is [insert doc].
	VolumeTypeBssd = VolumeType("bssd")
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

	Port uint32 `json:"port"`
	// Protocol:
	//
	// Default value: tcp
	Protocol ACLRuleProtocol `json:"protocol"`
	// Direction:
	//
	// Default value: inbound
	Direction ACLRuleDirection `json:"direction"`
	// Action:
	//
	// Default value: allow
	Action ACLRuleAction `json:"action"`

	Description string `json:"description"`
}

type ACLRuleRequest struct {
	IP scw.IPNet `json:"ip"`

	Description string `json:"description"`
}

// AddInstanceACLRulesResponse: add instance acl rules response
type AddInstanceACLRulesResponse struct {
	// Rules: rules enabled on the instance
	Rules []*ACLRule `json:"rules"`
}

// AddInstanceSettingsResponse: add instance settings response
type AddInstanceSettingsResponse struct {
	// Settings: settings available on the instance
	Settings []*InstanceSetting `json:"settings"`
}

type BackupSchedule struct {
	Frequency uint32 `json:"frequency"`

	Retention uint32 `json:"retention"`

	Disabled bool `json:"disabled"`
}

// Database: database
type Database struct {
	// Name: name of the database
	Name string `json:"name"`
	// Owner: name of the owner of the database
	Owner string `json:"owner"`
	// Managed: whether or not the database is managed or not
	Managed bool `json:"managed"`
	// Size: size of the database
	Size scw.Size `json:"size"`
}

// DatabaseBackup: database backup
type DatabaseBackup struct {
	// ID: UUID of the database backup
	ID string `json:"id"`
	// InstanceID: UUID of the instance
	InstanceID string `json:"instance_id"`
	// DatabaseName: name of the database of this backup
	DatabaseName string `json:"database_name"`
	// Name: name of the backup
	Name string `json:"name"`
	// Status: status of the backup
	//
	// Default value: unknown
	Status DatabaseBackupStatus `json:"status"`
	// Size: size of the database backup
	Size *scw.Size `json:"size"`
	// ExpiresAt: expiration date (Format ISO 8601)
	ExpiresAt *time.Time `json:"expires_at"`
	// CreatedAt: creation date (Format ISO 8601)
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: updated date (Format ISO 8601)
	UpdatedAt *time.Time `json:"updated_at"`
	// InstanceName: name of the instance of the backup
	InstanceName string `json:"instance_name"`
	// DownloadURL: URL you can download the backup from
	DownloadURL *string `json:"download_url"`
	// DownloadURLExpiresAt: expiration date of the download link
	DownloadURLExpiresAt *time.Time `json:"download_url_expires_at"`
	// Region: region of this database backup
	Region scw.Region `json:"region"`
}

type DatabaseEngine struct {
	Name string `json:"name"`

	LogoURL string `json:"logo_url"`

	Versions []*EngineVersion `json:"versions"`

	Region scw.Region `json:"region"`
}

// DeleteInstanceACLRulesResponse: delete instance acl rules response
type DeleteInstanceACLRulesResponse struct {
	// Rules: ACL rules present on the instance
	Rules []*ACLRule `json:"rules"`
}

// DeleteInstanceSettingsResponse: delete instance settings response
type DeleteInstanceSettingsResponse struct {
	// Settings: settings names to delete from the instance
	Settings []*InstanceSetting `json:"settings"`
}

// Endpoint: endpoint
type Endpoint struct {
	// IP: iPv4 address of the endpoint
	IP *net.IP `json:"ip"`
	// Port: TCP port of the endpoint
	Port uint32 `json:"port"`
	// Name: name of the endpoint
	Name *string `json:"name"`
}

type EngineSetting struct {
	Name string `json:"name"`

	DefaultValue string `json:"default_value"`

	HotConfigurable bool `json:"hot_configurable"`

	Description string `json:"description"`
	// PropertyType:
	//
	// Default value: BOOLEAN
	PropertyType EngineSettingPropertyType `json:"property_type"`

	Unit *string `json:"unit"`

	StringConstraint *string `json:"string_constraint"`

	IntMin *int32 `json:"int_min"`

	IntMax *int32 `json:"int_max"`
}

type EngineVersion struct {
	Version string `json:"version"`

	Name string `json:"name"`

	EndOfLife *time.Time `json:"end_of_life"`

	AvailableSettings []*EngineSetting `json:"available_settings"`

	Disabled bool `json:"disabled"`

	Beta bool `json:"beta"`
}

// Instance: instance
type Instance struct {
	// CreatedAt: creation date (Format ISO 8601)
	CreatedAt *time.Time `json:"created_at"`
	// Volume: volumes of the instance
	Volume *Volume `json:"volume"`
	// Region: region the instance is in
	Region scw.Region `json:"region"`
	// ID: UUID of the instance
	ID string `json:"id"`
	// Name: name of the instance
	Name string `json:"name"`
	// OrganizationID: organization ID the instance belongs to
	OrganizationID string `json:"organization_id"`
	// Status: status of the instance
	//
	// Default value: unknown
	Status InstanceStatus `json:"status"`
	// Engine: database engine of the database (PostgreSQL, MySQL, ...)
	Engine string `json:"engine"`
	// Endpoint: endpoint of the instance
	Endpoint *Endpoint `json:"endpoint"`
	// Tags: list of tags applied to the instance
	Tags []string `json:"tags"`
	// Settings: advanced settings of the instance
	Settings []*InstanceSetting `json:"settings"`
	// BackupSchedule: backup schedule of the instance
	BackupSchedule *BackupSchedule `json:"backup_schedule"`
	// IsHaCluster: whether or not High-Availability is enabled
	IsHaCluster bool `json:"is_ha_cluster"`
	// ReadReplicas: read replicas of the instance
	ReadReplicas []*Endpoint `json:"read_replicas"`
	// NodeType: node type of the instance
	NodeType string `json:"node_type"`
}

// InstanceLog: instance log
type InstanceLog struct {
	// DownloadURL: presigned S3 URL to download your log file
	DownloadURL *string `json:"download_url"`
	// ID: UUID of the instance log
	ID string `json:"id"`
	// Status: status of the logs in a given instance
	//
	// Default value: unknown
	Status InstanceLogStatus `json:"status"`
	// NodeName: name of the undelying node
	NodeName string `json:"node_name"`
	// ExpiresAt: expiration date (Format ISO 8601)
	ExpiresAt *time.Time `json:"expires_at"`
	// CreatedAt: creation date (Format ISO 8601)
	CreatedAt *time.Time `json:"created_at"`
	// Region: region the instance is in
	Region scw.Region `json:"region"`
}

// InstanceMetrics: instance metrics
type InstanceMetrics struct {
	// Timeseries: time series of metrics of a given instance
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

type InstanceSetting struct {
	Name string `json:"name"`

	Value string `json:"value"`
}

// ListDatabaseBackupsResponse: list database backups response
type ListDatabaseBackupsResponse struct {
	// DatabaseBackups: list of database backups
	DatabaseBackups []*DatabaseBackup `json:"database_backups"`
	// TotalCount: total count of database backups available
	TotalCount uint32 `json:"total_count"`
}

// ListDatabaseEnginesResponse: list database engines response
type ListDatabaseEnginesResponse struct {
	// Engines: list of the available database engines
	Engines []*DatabaseEngine `json:"engines"`
	// TotalCount: total count of database engines available
	TotalCount uint32 `json:"total_count"`
}

// ListDatabasesResponse: list databases response
type ListDatabasesResponse struct {
	// Databases: list of the databases
	Databases []*Database `json:"databases"`
	// TotalCount: total count of databases present on a given instance
	TotalCount uint32 `json:"total_count"`
}

// ListInstanceACLRulesResponse: list instance acl rules response
type ListInstanceACLRulesResponse struct {
	// Rules: list of the ACL rules present on a given instance
	Rules []*ACLRule `json:"rules"`
	// TotalCount: total count of ACL rules present on a given instance
	TotalCount uint32 `json:"total_count"`
}

// ListInstanceLogsResponse: list instance logs response
type ListInstanceLogsResponse struct {
	// InstanceLogs: available logs in a given instance
	InstanceLogs []*InstanceLog `json:"instance_logs"`
}

// ListInstancesResponse: list instances response
type ListInstancesResponse struct {
	// Instances: list all instances available in a given organization
	Instances []*Instance `json:"instances"`
	// TotalCount: total count of instances available in a given organization
	TotalCount uint32 `json:"total_count"`
}

// ListNodeTypesResponse: list node types response
type ListNodeTypesResponse struct {
	// NodeTypes: types of the node
	NodeTypes []*NodeType `json:"node_types"`
	// TotalCount: total count of node-types available
	TotalCount uint32 `json:"total_count"`
}

// ListPrivilegesResponse: list privileges response
type ListPrivilegesResponse struct {
	// Privileges: privileges of a given user in a given database in a given instance
	Privileges []*Privilege `json:"privileges"`
	// TotalCount: total count of privileges present on a given database
	TotalCount uint32 `json:"total_count"`
}

// ListUsersResponse: list users response
type ListUsersResponse struct {
	// Users: list of users in a given instance
	Users []*User `json:"users"`
	// TotalCount: total count of users present on a given instance
	TotalCount uint32 `json:"total_count"`
}

type NodeType struct {
	Name string `json:"name"`
	// StockStatus:
	//
	// Default value: unknown
	StockStatus NodeTypeStock `json:"stock_status"`

	Description string `json:"description"`

	Vcpus uint32 `json:"vcpus"`

	Memory scw.Size `json:"memory"`

	VolumeConstraint *NodeTypeVolumeConstraintSizes `json:"volume_constraint"`

	IsBssdCompatible bool `json:"is_bssd_compatible"`

	Disabled bool `json:"disabled"`

	Beta bool `json:"beta"`

	Region scw.Region `json:"region"`
}

type NodeTypeVolumeConstraintSizes struct {
	MinSize scw.Size `json:"min_size"`

	MaxSize scw.Size `json:"max_size"`
}

// PrepareInstanceLogsResponse: prepare instance logs response
type PrepareInstanceLogsResponse struct {
	// InstanceLogs: instance logs for a given instance between a start and an end date
	InstanceLogs []*InstanceLog `json:"instance_logs"`
}

// Privilege: privilege
type Privilege struct {
	// Permission: permission (Read, Read/Write, All, Custom)
	//
	// Default value: readonly
	Permission Permission `json:"permission"`
	// DatabaseName: name of the database
	DatabaseName string `json:"database_name"`
	// UserName: name of the user
	UserName string `json:"user_name"`
}

// SetInstanceACLRulesResponse: set instance acl rules response
type SetInstanceACLRulesResponse struct {
	// Rules: aCLs rules configured for an instance
	Rules []*ACLRule `json:"rules"`
}

// SetInstanceSettingsResponse: set instance settings response
type SetInstanceSettingsResponse struct {
	// Settings: settings configured for a given instance
	Settings []*InstanceSetting `json:"settings"`
}

// User: user
type User struct {
	// Name: name of the user (Length must be between 1 and 63 characters, First character must be an alphabet character (a-zA-Z), Your Username cannot start with '_rdb', Only a-zA-Z0-9_$- characters are accepted)
	Name string `json:"name"`
	// IsAdmin: whether or not a user got administrative privileges on the database instance
	IsAdmin bool `json:"is_admin"`
}

type Volume struct {
	// Type:
	//
	// Default value: lssd
	Type VolumeType `json:"type"`

	Size scw.Size `json:"size"`
}

// Service API

type GetServiceInfoRequest struct {
	Region scw.Region `json:"-"`
}

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
		Method:  "GET",
		Path:    "/rdb/v1/regions/" + fmt.Sprint(req.Region) + "",
		Headers: http.Header{},
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListDatabaseBackupsRequest struct {
	Region scw.Region `json:"-"`
	// Name: name of the database backups
	Name *string `json:"-"`
	// OrderBy: criteria to use when ordering database backups listing
	//
	// Default value: created_at_asc
	OrderBy ListDatabaseBackupsRequestOrderBy `json:"-"`
	// InstanceID: UUID of the instance
	InstanceID *string `json:"-"`
	// OrganizationID: organization ID the database backups belongs to
	OrganizationID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListDatabaseBackups: list database backups
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

type CreateDatabaseBackupRequest struct {
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance
	InstanceID string `json:"instance_id"`
	// DatabaseName: name of the database you want to make a backup of
	DatabaseName string `json:"database_name"`
	// Name: name of the backup
	Name string `json:"name"`
	// ExpiresAt: expiration date (Format ISO 8601)
	ExpiresAt *time.Time `json:"expires_at"`
}

// CreateDatabaseBackup: create a database backup
func (s *API) CreateDatabaseBackup(req *CreateDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
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
	Region scw.Region `json:"-"`
	// DatabaseBackupID: UUID of the database backup
	DatabaseBackupID string `json:"-"`
}

// GetDatabaseBackup: get a database backup
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
	Region scw.Region `json:"-"`
	// DatabaseBackupID: UUID of the database backup to update
	DatabaseBackupID string `json:"-"`
	// Name: name of the Database Backup
	Name *string `json:"name"`
	// ExpiresAt: expiration date (Format ISO 8601)
	ExpiresAt *time.Time `json:"expires_at"`
}

// UpdateDatabaseBackup: update a database backup
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
	Region scw.Region `json:"-"`
	// DatabaseBackupID: UUID of the database backup to delete
	DatabaseBackupID string `json:"-"`
}

// DeleteDatabaseBackup: delete a database backup
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
	Region scw.Region `json:"-"`
	// DatabaseBackupID: backup of a logical database
	DatabaseBackupID string `json:"-"`
	// DatabaseName: defines the destination database in order to restore into a specified database, the default destination is set to the origin database of the backup
	DatabaseName *string `json:"database_name"`
	// InstanceID: defines the rdb instance where the backup has to be restored
	InstanceID string `json:"instance_id"`
}

// RestoreDatabaseBackup: restore a database backup
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
	Region scw.Region `json:"-"`
	// DatabaseBackupID: UUID of the database backup you want to export
	DatabaseBackupID string `json:"-"`
}

// ExportDatabaseBackup: export a database backup
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

type CloneInstanceRequest struct {
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance you want to clone
	InstanceID string `json:"-"`
	// Name: name of the clone instance
	Name string `json:"name"`
	// NodeType: node type of the clone
	NodeType *string `json:"node_type"`
}

// CloneInstance: clone an instance
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

type ListDatabaseEnginesRequest struct {
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListDatabaseEngines: list available database engines
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

type UpgradeInstanceRequest struct {
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance you want to upgrade
	InstanceID string `json:"-"`
	// NodeType: node type of the instance you want to upgrade to
	// Precisely one of EnableHa, NodeType must be set.
	NodeType *string `json:"node_type,omitempty"`
	// EnableHa: set to true to enable high availability on your instance
	// Precisely one of EnableHa, NodeType must be set.
	EnableHa *bool `json:"enable_ha,omitempty"`
}

// UpgradeInstance: upgrade an instance to an higher instance type
//
// Upgrade your current `node_type` or enable high availability on your standalone database instance.
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
	Region scw.Region `json:"-"`
	// Tags: list instance that have a given tags
	Tags []string `json:"-"`
	// Name: list instance that match a given name pattern
	Name *string `json:"-"`
	// OrderBy: criteria to use when ordering instance listing
	//
	// Default value: created_at_asc
	OrderBy ListInstancesRequestOrderBy `json:"-"`
	// OrganizationID: organization ID to list the instance of
	OrganizationID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListInstances: list instances
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

type GetInstanceRequest struct {
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance
	InstanceID string `json:"-"`
}

// GetInstance: get an instance
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
	Region scw.Region `json:"-"`
	// OrganizationID: the organization ID on which to create the instance
	OrganizationID string `json:"organization_id"`
	// Name: name of the instance
	Name string `json:"name"`
	// Engine: database engine of the database (PostgreSQL, MySQL, ...)
	Engine string `json:"engine"`
	// UserName: name of the user created when the instance is created
	UserName string `json:"user_name"`
	// Password: password of the user
	Password string `json:"password"`
	// NodeType: type of node to use for the instance
	NodeType string `json:"node_type"`
	// IsHaCluster: whether or not High-Availability is enabled
	IsHaCluster bool `json:"is_ha_cluster"`
	// DisableBackup: whether or not backups are disabled
	DisableBackup bool `json:"disable_backup"`
	// Tags: tags to apply to the instance
	Tags []string `json:"tags"`
}

// CreateInstance: create an instance
func (s *API) CreateInstance(req *CreateInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance to update
	InstanceID string `json:"-"`
	// BackupScheduleFrequency: in hours
	BackupScheduleFrequency *uint32 `json:"backup_schedule_frequency"`
	// BackupScheduleRetention: in days
	BackupScheduleRetention *uint32 `json:"backup_schedule_retention"`
	// IsBackupScheduleDisabled: whether or not the backup schedule is disabled
	IsBackupScheduleDisabled *bool `json:"is_backup_schedule_disabled"`
	// Name: name of the instance
	Name *string `json:"name"`
	// Tags: tags of a given instance
	Tags *[]string `json:"tags"`
}

// UpdateInstance: update an instance
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance to delete
	InstanceID string `json:"-"`
}

// DeleteInstance: delete an instance
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

type GetInstanceCertificateRequest struct {
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance
	InstanceID string `json:"-"`
}

// GetInstanceCertificate: get the TLS certificate of an instance
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

type PrepareInstanceLogsRequest struct {
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance you want logs of
	InstanceID string `json:"-"`
	// StartDate: start datetime of your log. Format: `{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z`
	StartDate *time.Time `json:"start_date"`
	// EndDate: end datetime of your log. Format: `{year}-{month}-{day}T{hour}:{min}:{sec}[.{frac_sec}]Z`
	EndDate *time.Time `json:"end_date"`
}

// PrepareInstanceLogs: prepare logs of a given instance
//
// Prepare your instance logs. Logs will be grouped on a minimum interval of a day.
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance you want logs of
	InstanceID string `json:"-"`
	// OrderBy: criteria to use when ordering instance logs listing
	//
	// Default value: created_at_asc
	OrderBy ListInstanceLogsRequestOrderBy `json:"-"`
}

// ListInstanceLogs: list available logs of a given instance
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
	Region scw.Region `json:"-"`
	// InstanceLogID: UUID of the instance_log you want
	InstanceLogID string `json:"-"`
}

// GetInstanceLog: get specific logs of a given instance
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

type GetInstanceMetricsRequest struct {
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance
	InstanceID string `json:"-"`
	// StartDate: start date to gather metrics from
	StartDate *time.Time `json:"-"`
	// EndDate: end date to gather metrics from
	EndDate *time.Time `json:"-"`
	// MetricName: name of the metric to gather
	MetricName *string `json:"-"`
}

// GetInstanceMetrics: get instance metrics
//
// Get database instance metrics.
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

type AddInstanceSettingsRequest struct {
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance you want to add settings to
	InstanceID string `json:"-"`
	// Settings: settings to add on the instance
	Settings []*InstanceSetting `json:"settings"`
}

// AddInstanceSettings: add an instance setting
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance to delete settings from
	InstanceID string `json:"-"`
	// SettingNames: settings names to delete
	SettingNames []string `json:"setting_names"`
}

// DeleteInstanceSettings: delete an instance setting
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance where the settings has to be set
	InstanceID string `json:"-"`
	// Settings: settings to define for the instance
	Settings []*InstanceSetting `json:"settings"`
}

// SetInstanceSettings: set a given instance setting
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance
	InstanceID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListInstanceACLRules: list ACL rules of a given instance
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

type AddInstanceACLRulesRequest struct {
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance you want to add acl rules to
	InstanceID string `json:"-"`
	// Rules: aCLs rules to add to the instance
	Rules []*ACLRuleRequest `json:"rules"`
}

// AddInstanceACLRules: add an ACL instance to a given instance
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance where the ACL rules has to be set
	InstanceID string `json:"-"`
	// Rules: ACL rules to define for the instance
	Rules []*ACLRuleRequest `json:"rules"`
}

// SetInstanceACLRules: set ACL rules for a given instance
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance you want to delete an ACL rules from
	InstanceID string `json:"-"`
	// ACLRuleIPs: ACL rules IP present on the instance
	ACLRuleIPs []string `json:"acl_rule_ips"`
}

// DeleteInstanceACLRules: delete ACL rules of a given instance
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance
	InstanceID string `json:"-"`
	// Name: name of the user
	Name *string `json:"-"`
	// OrderBy: criteria to use when ordering users listing
	//
	// Default value: name_asc
	OrderBy ListUsersRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListUsers: list users of a given instance
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

type CreateUserRequest struct {
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance you want to create a user in
	InstanceID string `json:"-"`
	// Name: name of the user you want to create
	Name string `json:"name"`
	// Password: password of the user you want to create
	Password string `json:"password"`
	// IsAdmin: whether the user you want to create will have administrative privileges
	IsAdmin bool `json:"is_admin"`
}

// CreateUser: create an user in a given instance
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance the user belongs to
	InstanceID string `json:"-"`
	// Name: name of the database user
	Name string `json:"-"`
	// Password: password of the database user
	Password *string `json:"password"`
	// IsAdmin: whether or not this user got administrative privileges
	IsAdmin *bool `json:"is_admin"`
}

// UpdateUser: update an user in a given instance
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance to delete a user from
	InstanceID string `json:"-"`
	// Name: name of the user
	Name string `json:"-"`
}

// DeleteUser: delete an user in a given instance
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance to list database of
	InstanceID string `json:"-"`
	// Name: name of the database
	Name *string `json:"-"`
	// Managed: whether or not the database is managed
	Managed *bool `json:"-"`
	// Owner: user that owns this database
	Owner *string `json:"-"`
	// OrderBy: criteria to use when ordering database listing
	//
	// Default value: name_asc
	OrderBy ListDatabasesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListDatabases: list all database in a given instance
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

type CreateDatabaseRequest struct {
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance where to create the database
	InstanceID string `json:"-"`
	// Name: name of the database
	Name string `json:"name"`
}

// CreateDatabase: create a database in a given instance
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance where to delete the database
	InstanceID string `json:"-"`
	// Name: name of the database to delete
	Name string `json:"-"`
}

// DeleteDatabase: delete a database in a given instance
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
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance
	InstanceID string `json:"-"`
	// OrderBy: criteria to use when ordering privileges listing
	//
	// Default value: user_name_asc
	OrderBy ListPrivilegesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// DatabaseName: name of the database
	DatabaseName *string `json:"-"`
	// UserName: name of the user
	UserName *string `json:"-"`
}

// ListPrivileges: list privileges of a given user in a given database in a given instance
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

type SetPrivilegeRequest struct {
	Region scw.Region `json:"-"`
	// InstanceID: UUID of the instance
	InstanceID string `json:"-"`
	// DatabaseName: name of the database
	DatabaseName string `json:"database_name"`
	// UserName: name of the user
	UserName string `json:"user_name"`
	// Permission: permission to set (Read, Read/Write, All, Custom)
	//
	// Default value: readonly
	Permission Permission `json:"permission"`
}

// SetPrivilege: set privileges of a given user in a given database in a given instance
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

type ListNodeTypesRequest struct {
	Region scw.Region `json:"-"`
	// IncludeDisabledTypes: whether or not to include disabled types
	IncludeDisabledTypes bool `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListNodeTypes: list available node types
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

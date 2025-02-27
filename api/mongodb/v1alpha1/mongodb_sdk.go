// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package mongodb provides methods and message types of the mongodb v1alpha1 API.
package mongodb

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

type InstanceStatus string

const (
	InstanceStatusUnknownStatus = InstanceStatus("unknown_status")
	InstanceStatusReady         = InstanceStatus("ready")
	InstanceStatusProvisioning  = InstanceStatus("provisioning")
	InstanceStatusConfiguring   = InstanceStatus("configuring")
	InstanceStatusDeleting      = InstanceStatus("deleting")
	InstanceStatusError         = InstanceStatus("error")
	InstanceStatusInitializing  = InstanceStatus("initializing")
	InstanceStatusLocked        = InstanceStatus("locked")
	InstanceStatusSnapshotting  = InstanceStatus("snapshotting")
)

func (enum InstanceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum InstanceStatus) Values() []InstanceStatus {
	return []InstanceStatus{
		"unknown_status",
		"ready",
		"provisioning",
		"configuring",
		"deleting",
		"error",
		"initializing",
		"locked",
		"snapshotting",
	}
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

type ListInstancesRequestOrderBy string

const (
	ListInstancesRequestOrderByCreatedAtAsc  = ListInstancesRequestOrderBy("created_at_asc")
	ListInstancesRequestOrderByCreatedAtDesc = ListInstancesRequestOrderBy("created_at_desc")
	ListInstancesRequestOrderByNameAsc       = ListInstancesRequestOrderBy("name_asc")
	ListInstancesRequestOrderByNameDesc      = ListInstancesRequestOrderBy("name_desc")
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

func (enum ListInstancesRequestOrderBy) Values() []ListInstancesRequestOrderBy {
	return []ListInstancesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
		"status_asc",
		"status_desc",
	}
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

func (enum ListSnapshotsRequestOrderBy) Values() []ListSnapshotsRequestOrderBy {
	return []ListSnapshotsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
		"expires_at_asc",
		"expires_at_desc",
	}
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
	ListUsersRequestOrderByNameAsc  = ListUsersRequestOrderBy("name_asc")
	ListUsersRequestOrderByNameDesc = ListUsersRequestOrderBy("name_desc")
)

func (enum ListUsersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListUsersRequestOrderBy) Values() []ListUsersRequestOrderBy {
	return []ListUsersRequestOrderBy{
		"name_asc",
		"name_desc",
	}
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
	NodeTypeStockUnknownStock = NodeTypeStock("unknown_stock")
	NodeTypeStockLowStock     = NodeTypeStock("low_stock")
	NodeTypeStockOutOfStock   = NodeTypeStock("out_of_stock")
	NodeTypeStockAvailable    = NodeTypeStock("available")
)

func (enum NodeTypeStock) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_stock"
	}
	return string(enum)
}

func (enum NodeTypeStock) Values() []NodeTypeStock {
	return []NodeTypeStock{
		"unknown_stock",
		"low_stock",
		"out_of_stock",
		"available",
	}
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

type SettingPropertyType string

const (
	SettingPropertyTypeBOOLEAN = SettingPropertyType("BOOLEAN")
	SettingPropertyTypeINT     = SettingPropertyType("INT")
	SettingPropertyTypeSTRING  = SettingPropertyType("STRING")
	SettingPropertyTypeFLOAT   = SettingPropertyType("FLOAT")
)

func (enum SettingPropertyType) String() string {
	if enum == "" {
		// return default value if empty
		return "BOOLEAN"
	}
	return string(enum)
}

func (enum SettingPropertyType) Values() []SettingPropertyType {
	return []SettingPropertyType{
		"BOOLEAN",
		"INT",
		"STRING",
		"FLOAT",
	}
}

func (enum SettingPropertyType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SettingPropertyType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SettingPropertyType(SettingPropertyType(tmp).String())
	return nil
}

type SnapshotStatus string

const (
	SnapshotStatusUnknownStatus = SnapshotStatus("unknown_status")
	SnapshotStatusCreating      = SnapshotStatus("creating")
	SnapshotStatusReady         = SnapshotStatus("ready")
	SnapshotStatusRestoring     = SnapshotStatus("restoring")
	SnapshotStatusDeleting      = SnapshotStatus("deleting")
	SnapshotStatusError         = SnapshotStatus("error")
	SnapshotStatusLocked        = SnapshotStatus("locked")
)

func (enum SnapshotStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum SnapshotStatus) Values() []SnapshotStatus {
	return []SnapshotStatus{
		"unknown_status",
		"creating",
		"ready",
		"restoring",
		"deleting",
		"error",
		"locked",
	}
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

type UserRoleRole string

const (
	UserRoleRoleUnknownRole = UserRoleRole("unknown_role")
	UserRoleRoleRead        = UserRoleRole("read")
	UserRoleRoleReadWrite   = UserRoleRole("read_write")
	UserRoleRoleDbAdmin     = UserRoleRole("db_admin")
)

func (enum UserRoleRole) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_role"
	}
	return string(enum)
}

func (enum UserRoleRole) Values() []UserRoleRole {
	return []UserRoleRole{
		"unknown_role",
		"read",
		"read_write",
		"db_admin",
	}
}

func (enum UserRoleRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UserRoleRole) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UserRoleRole(UserRoleRole(tmp).String())
	return nil
}

type VolumeType string

const (
	VolumeTypeUnknownType = VolumeType("unknown_type")
	VolumeTypeSbs5k       = VolumeType("sbs_5k")
	VolumeTypeSbs15k      = VolumeType("sbs_15k")
)

func (enum VolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum VolumeType) Values() []VolumeType {
	return []VolumeType{
		"unknown_type",
		"sbs_5k",
		"sbs_15k",
	}
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

// EndpointPrivateNetworkDetails: Private Network details.
type EndpointPrivateNetworkDetails struct {
	// PrivateNetworkID: UUID of the Private Network.
	PrivateNetworkID string `json:"private_network_id"`
}

// EndpointPublicDetails: endpoint public details.
type EndpointPublicDetails struct {
}

// EndpointSpecPrivateNetworkDetails: endpoint spec private network details.
type EndpointSpecPrivateNetworkDetails struct {
	// PrivateNetworkID: UUID of the Private Network.
	PrivateNetworkID string `json:"private_network_id"`
}

// EndpointSpecPublicDetails: endpoint spec public details.
type EndpointSpecPublicDetails struct {
}

// Endpoint: endpoint.
type Endpoint struct {
	// ID: UUID of the endpoint.
	ID string `json:"id"`

	// IPs: list of IPv4 addresses of the endpoint.
	IPs []net.IP `json:"ips"`

	// DNSRecords: list of DNS records of the endpoint.
	DNSRecords []string `json:"dns_records"`

	// Port: TCP port of the endpoint.
	Port uint32 `json:"port"`

	// PrivateNetwork: private Network endpoint details.
	// Precisely one of PrivateNetwork, Public must be set.
	PrivateNetwork *EndpointPrivateNetworkDetails `json:"private_network,omitempty"`

	// Public: public endpoint details.
	// Precisely one of PrivateNetwork, Public must be set.
	Public *EndpointPublicDetails `json:"public,omitempty"`
}

// InstanceSetting: instance setting.
type InstanceSetting struct {
	// Name: name of the settings.
	Name string `json:"name"`

	// Value: value of the settings.
	Value string `json:"value"`
}

// Volume: volume.
type Volume struct {
	// Type: type of volume where data is stored.
	// Default value: unknown_type
	Type VolumeType `json:"type"`

	// Size: volume size.
	Size scw.Size `json:"size"`
}

// NodeTypeVolumeType: node type volume type.
type NodeTypeVolumeType struct {
	// Type: volume Type.
	// Default value: unknown_type
	Type VolumeType `json:"type"`

	// Description: the description of the volume.
	Description string `json:"description"`

	// MinSize: mimimum size required for the volume.
	MinSize scw.Size `json:"min_size"`

	// MaxSize: maximum size required for the volume.
	MaxSize scw.Size `json:"max_size"`

	// ChunkSize: minimum increment level for a Block Storage volume size.
	ChunkSize scw.Size `json:"chunk_size"`
}

// SnapshotVolumeType: snapshot volume type.
type SnapshotVolumeType struct {
	// Type: default value: unknown_type
	Type VolumeType `json:"type"`
}

// UserRole: user role.
type UserRole struct {
	// Role: default value: unknown_role
	Role UserRoleRole `json:"role"`

	// Precisely one of Database, AnyDatabase must be set.
	Database *string `json:"database,omitempty"`

	// Precisely one of Database, AnyDatabase must be set.
	AnyDatabase *bool `json:"any_database,omitempty"`
}

// Setting: setting.
type Setting struct {
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
	PropertyType SettingPropertyType `json:"property_type"`

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

// EndpointSpec: endpoint spec.
type EndpointSpec struct {
	// Precisely one of Public, PrivateNetwork must be set.
	Public *EndpointSpecPublicDetails `json:"public,omitempty"`

	// Precisely one of Public, PrivateNetwork must be set.
	PrivateNetwork *EndpointSpecPrivateNetworkDetails `json:"private_network,omitempty"`
}

// CreateInstanceRequestVolumeDetails: create instance request volume details.
type CreateInstanceRequestVolumeDetails struct {
	// VolumeSize: volume size.
	VolumeSize scw.Size `json:"volume_size"`

	// VolumeType: type of volume where data is stored.
	// Default value: unknown_type
	VolumeType VolumeType `json:"volume_type"`
}

// Instance: instance.
type Instance struct {
	// ID: UUID of the Database Instance.
	ID string `json:"id"`

	// Name: name of the Database Instance.
	Name string `json:"name"`

	// ProjectID: project ID the Database Instance belongs to.
	ProjectID string `json:"project_id"`

	// Status: status of the Database Instance.
	// Default value: unknown_status
	Status InstanceStatus `json:"status"`

	// Version: mongoDB® engine version of the Database Instance.
	Version string `json:"version"`

	// Tags: list of tags applied to the Database Instance.
	Tags []string `json:"tags"`

	// Settings: advanced settings of the Database Instance.
	Settings []*InstanceSetting `json:"settings"`

	// NodeNumber: number of node in the Database Instance.
	NodeNumber uint32 `json:"node_number"`

	// NodeType: node type of the Database Instance.
	NodeType string `json:"node_type"`

	// Volume: volumes of the Database Instance.
	Volume *Volume `json:"volume"`

	// Endpoints: list of Database Instance endpoints.
	Endpoints []*Endpoint `json:"endpoints"`

	// CreatedAt: creation date (must follow the ISO 8601 format).
	CreatedAt *time.Time `json:"created_at"`

	// Region: region the Database Instance is in.
	Region scw.Region `json:"region"`
}

// NodeType: node type.
type NodeType struct {
	// Name: node type name identifier.
	Name string `json:"name"`

	// StockStatus: current stock status for the node type.
	// Default value: unknown_stock
	StockStatus NodeTypeStock `json:"stock_status"`

	// Description: current specs of the offer.
	Description string `json:"description"`

	// Vcpus: number of virtual CPUs.
	Vcpus uint32 `json:"vcpus"`

	// Memory: quantity of RAM.
	Memory scw.Size `json:"memory"`

	// AvailableVolumeTypes: available storage options for the node type.
	AvailableVolumeTypes []*NodeTypeVolumeType `json:"available_volume_types"`

	// Disabled: the node type is currently disabled.
	Disabled bool `json:"disabled"`

	// Beta: the node type is currently in beta.
	Beta bool `json:"beta"`

	// InstanceRange: instance range associated with the node type offer.
	InstanceRange string `json:"instance_range"`
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
	// Default value: unknown_status
	Status SnapshotStatus `json:"status"`

	// Size: size of the snapshot.
	Size scw.Size `json:"size"`

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

	// VolumeType: type of volume where data is stored - sbs_5k or sbs_15k.
	VolumeType *SnapshotVolumeType `json:"volume_type"`

	// Region: region of the snapshot.
	Region scw.Region `json:"region"`
}

// User: user.
type User struct {
	// Name: name of the user (Length must be between 1 and 63 characters. First character must be an alphabet character (a-zA-Z). Only a-zA-Z0-9_$- characters are accepted).
	Name string `json:"name"`

	// Roles: list of roles assigned to the user, along with the corresponding database where each role is granted.
	Roles []*UserRole `json:"roles"`
}

// Version: version.
type Version struct {
	// Version: mongoDB® engine version.
	Version string `json:"version"`

	// EndOfLifeAt: date of End of Life.
	EndOfLifeAt *time.Time `json:"end_of_life_at"`

	// AvailableSettings: instance settings available to be updated.
	AvailableSettings []*Setting `json:"available_settings"`
}

// RestoreSnapshotRequestVolumeDetails: restore snapshot request volume details.
type RestoreSnapshotRequestVolumeDetails struct {
	// VolumeType: type of volume where data is stored.
	// Default value: unknown_type
	VolumeType VolumeType `json:"volume_type"`
}

// CreateEndpointRequest: create endpoint request.
type CreateEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"instance_id"`

	// Endpoint: endpointSpec used to expose your Database Instance.
	Endpoint *EndpointSpec `json:"endpoint"`
}

// CreateInstanceRequest: create instance request.
type CreateInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: the Project ID on which the Database Instance will be created.
	ProjectID string `json:"project_id"`

	// Name: name of the Database Instance.
	Name string `json:"name"`

	// Version: version of the MongoDB® engine.
	Version string `json:"version"`

	// Tags: tags to apply to the Database Instance.
	Tags []string `json:"tags"`

	// NodeNumber: number of node to use for the Database Instance.
	NodeNumber uint32 `json:"node_number"`

	// NodeType: type of node to use for the Database Instance.
	NodeType string `json:"node_type"`

	// UserName: username created when the Database Instance is created.
	UserName string `json:"user_name"`

	// Password: password of the initial user.
	Password string `json:"password"`

	// Volume: instance volume information.
	Volume *CreateInstanceRequestVolumeDetails `json:"volume,omitempty"`

	// Endpoints: one or multiple EndpointSpec used to expose your Database Instance.
	Endpoints []*EndpointSpec `json:"endpoints"`
}

// CreateSnapshotRequest: create snapshot request.
type CreateSnapshotRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// InstanceID: UUID of the Database Instance to snapshot.
	InstanceID string `json:"-"`

	// Name: name of the snapshot.
	Name string `json:"name"`

	// ExpiresAt: expiration date of the snapshot (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// CreateUserRequest: create user request.
type CreateUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// InstanceID: UUID of the Database Instance the user belongs to.
	InstanceID string `json:"-"`

	// Name: name of the database user.
	Name string `json:"name"`

	// Password: password of the database user.
	Password string `json:"password"`
}

// DeleteEndpointRequest: delete endpoint request.
type DeleteEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// EndpointID: UUID of the Endpoint to delete.
	EndpointID string `json:"-"`
}

// DeleteInstanceRequest: delete instance request.
type DeleteInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// InstanceID: UUID of the Database Instance to delete.
	InstanceID string `json:"-"`
}

// DeleteSnapshotRequest: delete snapshot request.
type DeleteSnapshotRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SnapshotID: UUID of the snapshot.
	SnapshotID string `json:"-"`
}

// DeleteUserRequest: delete user request.
type DeleteUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// InstanceID: UUID of the Database Instance the user belongs to.
	InstanceID string `json:"-"`

	// Name: name of the database user.
	Name string `json:"-"`
}

// GetInstanceCertificateRequest: get instance certificate request.
type GetInstanceCertificateRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
}

// GetInstanceRequest: get instance request.
type GetInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// InstanceID: UUID of the Database Instance.
	InstanceID string `json:"-"`
}

// GetSnapshotRequest: get snapshot request.
type GetSnapshotRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SnapshotID: UUID of the snapshot.
	SnapshotID string `json:"-"`
}

// ListInstancesRequest: list instances request.
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

	// OrganizationID: organization ID of the Database Instance.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID.
	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListInstancesResponse: list instances response.
type ListInstancesResponse struct {
	// Instances: list of all Database Instances available in an Organization or Project.
	Instances []*Instance `json:"instances"`

	// TotalCount: total count of Database Instances available in an Organization or Project.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInstancesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInstancesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListInstancesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Instances = append(r.Instances, results.Instances...)
	r.TotalCount += uint64(len(results.Instances))
	return uint64(len(results.Instances)), nil
}

// ListNodeTypesRequest: list node types request.
type ListNodeTypesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IncludeDisabledTypes: defines whether or not to include disabled types.
	IncludeDisabledTypes *bool `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListNodeTypesResponse: list node types response.
type ListNodeTypesResponse struct {
	// NodeTypes: types of the node.
	NodeTypes []*NodeType `json:"node_types"`

	// TotalCount: total count of node-types available.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNodeTypesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNodeTypesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListNodeTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.NodeTypes = append(r.NodeTypes, results.NodeTypes...)
	r.TotalCount += uint64(len(results.NodeTypes))
	return uint64(len(results.NodeTypes)), nil
}

// ListSnapshotsRequest: list snapshots request.
type ListSnapshotsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// InstanceID: instance ID the snapshots belongs to.
	InstanceID *string `json:"-"`

	// Name: lists database snapshots that match a name pattern.
	Name *string `json:"-"`

	// OrderBy: criteria to use when ordering snapshot listings.
	// Default value: created_at_asc
	OrderBy ListSnapshotsRequestOrderBy `json:"-"`

	// OrganizationID: organization ID the snapshots belongs to.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID to list the snapshots of.
	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListSnapshotsResponse: list snapshots response.
type ListSnapshotsResponse struct {
	// Snapshots: list of all database snapshots available in an Organization or Project.
	Snapshots []*Snapshot `json:"snapshots"`

	// TotalCount: total count of database snapshots available in a Organization or Project.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListSnapshotsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Snapshots = append(r.Snapshots, results.Snapshots...)
	r.TotalCount += uint64(len(results.Snapshots))
	return uint64(len(results.Snapshots)), nil
}

// ListUsersRequest: list users request.
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

// ListUsersResponse: list users response.
type ListUsersResponse struct {
	// Users: list of users in a Database Instance.
	Users []*User `json:"users"`

	// TotalCount: total count of users present on a Database Instance.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Users = append(r.Users, results.Users...)
	r.TotalCount += uint64(len(results.Users))
	return uint64(len(results.Users)), nil
}

// ListVersionsRequest: list versions request.
type ListVersionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Version *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListVersionsResponse: list versions response.
type ListVersionsResponse struct {
	// Versions: available MongoDB® engine version.
	Versions []*Version `json:"versions"`

	// TotalCount: total count of MongoDB® engine version available.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVersionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVersionsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Versions = append(r.Versions, results.Versions...)
	r.TotalCount += uint64(len(results.Versions))
	return uint64(len(results.Versions)), nil
}

// RestoreSnapshotRequest: restore snapshot request.
type RestoreSnapshotRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SnapshotID: UUID of the snapshot.
	SnapshotID string `json:"-"`

	// InstanceName: name of the new Database Instance.
	InstanceName string `json:"instance_name"`

	// NodeType: node type to use for the new Database Instance.
	NodeType string `json:"node_type"`

	// NodeNumber: number of nodes to use for the new Database Instance.
	NodeNumber uint32 `json:"node_number"`

	// Volume: instance volume information.
	Volume *RestoreSnapshotRequestVolumeDetails `json:"volume"`
}

// SetUserRoleRequest: set user role request.
type SetUserRoleRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// InstanceID: UUID of the Database Instance the user belongs to.
	InstanceID string `json:"-"`

	// UserName: name of the database user.
	UserName string `json:"user_name"`

	// Roles: list of roles assigned to the user, along with the corresponding database where each role is granted.
	Roles []*UserRole `json:"roles"`
}

// UpdateInstanceRequest: update instance request.
type UpdateInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// InstanceID: UUID of the Database Instance to update.
	InstanceID string `json:"-"`

	// Name: name of the Database Instance.
	Name *string `json:"name,omitempty"`

	// Tags: tags of a Database Instance.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateSnapshotRequest: update snapshot request.
type UpdateSnapshotRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SnapshotID: UUID of the Snapshot.
	SnapshotID string `json:"-"`

	// Name: name of the snapshot.
	Name *string `json:"name,omitempty"`

	// ExpiresAt: expiration date of the snapshot (must follow the ISO 8601 format).
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// UpdateUserRequest: update user request.
type UpdateUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// InstanceID: UUID of the Database Instance the user belongs to.
	InstanceID string `json:"-"`

	// Name: name of the database user.
	Name string `json:"-"`

	// Password: password of the database user.
	Password *string `json:"password,omitempty"`
}

// UpgradeInstanceRequest: upgrade instance request.
type UpgradeInstanceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// InstanceID: UUID of the Database Instance you want to upgrade.
	InstanceID string `json:"-"`

	// VolumeSize: increase your Block Storage volume size.
	// Precisely one of VolumeSize must be set.
	VolumeSize *scw.Size `json:"volume_size,omitempty"`
}

// This API allows you to manage your Managed Databases for MongoDB®.
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
	return []scw.Region{scw.RegionFrPar}
}

// ListNodeTypes: List available node types.
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
		Method: "GET",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/node-types",
		Query:  query,
	}

	var resp ListNodeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVersions: List available MongoDB® versions.
func (s *API) ListVersions(req *ListVersionsRequest, opts ...scw.RequestOption) (*ListVersionsResponse, error) {
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
	parameter.AddToQuery(query, "version", req.Version)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/versions",
		Query:  query,
	}

	var resp ListVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInstances: List all MongoDB® Database Instances in the specified region. By default, the MongoDB® Database Instances returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `tags` and `name`. For the `name` parameter, the value you include will be checked against the whole name string to see if it includes the string you put in the parameter.
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
		Method: "GET",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances",
		Query:  query,
	}

	var resp ListInstancesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInstance: Retrieve information about a given MongoDB® Database Instance, specified by the `region` and `instance_id` parameters. Its full details, including name, status, IP address and port, are returned in the response object.
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
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateInstance: Create a new MongoDB® Database Instance.
func (s *API) CreateInstance(req *CreateInstanceRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		req.Name = namegenerator.GetRandomName("mgdb")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances",
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

// UpdateInstance: Update the parameters of a MongoDB® Database Instance.
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
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
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

// DeleteInstance: Delete a given MongoDB® Database Instance, specified by the `region` and `instance_id` parameters. Deleting a MongoDB® Database Instance is permanent, and cannot be undone. Note that upon deletion all your data will be lost.
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
		Method: "DELETE",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "",
	}

	var resp Instance

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpgradeInstance: Upgrade your current Database Instance specifications like volume size.
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
		Method: "POST",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/upgrade",
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

// GetInstanceCertificate: Retrieve the certificate of a given Database Instance, specified by the `instance_id` parameter.
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
		Method: "GET",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/certificate",
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSnapshot: Create a new snapshot of a Database Instance. You must define the `name` and `instance_id` parameters in the request.
func (s *API) CreateSnapshot(req *CreateSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
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
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/snapshots",
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

// GetSnapshot: Retrieve information about a given snapshot of a Database Instance. You must specify, in the endpoint, the `snapshot_id` parameter of the snapshot you want to retrieve.
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
		Method: "GET",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSnapshot: Update the parameters of a snapshot of a Database Instance. You can update the `name` and `expires_at` parameters.
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
		Method: "POST",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
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

// RestoreSnapshot: Restore a given snapshot of a Database Instance. You must specify, in the endpoint, the `snapshot_id` parameter of the snapshot you want to restore, the `instance_name` of the new Database Instance, `node_type` of the new Database Instance and `node_number` of the new Database Instance.
func (s *API) RestoreSnapshot(req *RestoreSnapshotRequest, opts ...scw.RequestOption) (*Instance, error) {
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
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "/restore",
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

// ListSnapshots: List snapshots. You can include the `instance_id` or `project_id` in your query to get the list of snapshots for specific Database Instances and/or Projects. By default, the details returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
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
	parameter.AddToQuery(query, "instance_id", req.InstanceID)
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
		Method: "GET",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/snapshots",
		Query:  query,
	}

	var resp ListSnapshotsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSnapshot: Delete a given snapshot of a Database Instance. You must specify, in the endpoint, the `snapshot_id` parameter of the snapshot you want to delete.
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
		Method: "DELETE",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListUsers: List all users of a given Database Instance.
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
		Method: "GET",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users",
		Query:  query,
	}

	var resp ListUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateUser: Create an user on a Database Instance. You must define the `name`, `password` of the user and `instance_id` parameters in the request.
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
		Method: "POST",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users",
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

// UpdateUser: Update the parameters of a user on a Database Instance. You can update the `password` parameter, but you cannot change the name of the user.
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
		Method: "PATCH",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users/" + fmt.Sprint(req.Name) + "",
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

// DeleteUser: Delete an existing user on a Database Instance.
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
		Method: "DELETE",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/users/" + fmt.Sprint(req.Name) + "",
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

// SetUserRole:
func (s *API) SetUserRole(req *SetUserRoleRequest, opts ...scw.RequestOption) (*User, error) {
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
		Method: "PUT",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/instances/" + fmt.Sprint(req.InstanceID) + "/roles",
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

// DeleteEndpoint: Delete the endpoint of a Database Instance. You must specify the `endpoint_id` parameter of the endpoint you want to delete. Note that you might need to update any environment configurations that point to the deleted endpoint.
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
		Method: "DELETE",
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateEndpoint: Create a new endpoint for a MongoDB® Database Instance. You can add `public_network` or `private_network` specifications to the body of the request.
func (s *API) CreateEndpoint(req *CreateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
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
		Path:   "/mongodb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/endpoints",
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

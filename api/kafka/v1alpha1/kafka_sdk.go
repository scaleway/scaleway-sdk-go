// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package kafka provides methods and message types of the kafka v1alpha1 API.
package kafka

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
	defaultKafkaRetryInterval = 15 * time.Second
	defaultKafkaTimeout       = 15 * time.Minute
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

type ClusterStatus string

const (
	ClusterStatusUnknownStatus = ClusterStatus("unknown_status")
	ClusterStatusReady         = ClusterStatus("ready")
	ClusterStatusCreating      = ClusterStatus("creating")
	ClusterStatusConfiguring   = ClusterStatus("configuring")
	ClusterStatusDeleting      = ClusterStatus("deleting")
	ClusterStatusError         = ClusterStatus("error")
	ClusterStatusLocked        = ClusterStatus("locked")
	ClusterStatusStopped       = ClusterStatus("stopped")
)

func (enum ClusterStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ClusterStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ClusterStatus) Values() []ClusterStatus {
	return []ClusterStatus{
		"unknown_status",
		"ready",
		"creating",
		"configuring",
		"deleting",
		"error",
		"locked",
		"stopped",
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

type ListClustersRequestOrderBy string

const (
	ListClustersRequestOrderByCreatedAtAsc  = ListClustersRequestOrderBy("created_at_asc")
	ListClustersRequestOrderByCreatedAtDesc = ListClustersRequestOrderBy("created_at_desc")
	ListClustersRequestOrderByNameAsc       = ListClustersRequestOrderBy("name_asc")
	ListClustersRequestOrderByNameDesc      = ListClustersRequestOrderBy("name_desc")
	ListClustersRequestOrderByStatusAsc     = ListClustersRequestOrderBy("status_asc")
	ListClustersRequestOrderByStatusDesc    = ListClustersRequestOrderBy("status_desc")
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
		"name_asc",
		"name_desc",
		"status_asc",
		"status_desc",
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

type ListUsersRequestOrderBy string

const (
	ListUsersRequestOrderByNameAsc  = ListUsersRequestOrderBy("name_asc")
	ListUsersRequestOrderByNameDesc = ListUsersRequestOrderBy("name_desc")
)

func (enum ListUsersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListUsersRequestOrderByNameAsc)
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
		return string(NodeTypeStockUnknownStock)
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

type VolumeType string

const (
	VolumeTypeUnknownType = VolumeType("unknown_type")
	VolumeTypeSbs5k       = VolumeType("sbs_5k")
	VolumeTypeSbs15k      = VolumeType("sbs_15k")
)

func (enum VolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return string(VolumeTypeUnknownType)
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

// EndpointPublicDetails: Public Access details.
type EndpointPublicDetails struct{}

// VersionAvailableSettingBooleanProperty: version available setting boolean property.
type VersionAvailableSettingBooleanProperty struct {
	DefaultValue bool `json:"default_value"`
}

// VersionAvailableSettingFloatProperty: version available setting float property.
type VersionAvailableSettingFloatProperty struct {
	Min float32 `json:"min"`

	Max float32 `json:"max"`

	DefaultValue float32 `json:"default_value"`

	Unit *string `json:"unit"`
}

// VersionAvailableSettingIntegerProperty: version available setting integer property.
type VersionAvailableSettingIntegerProperty struct {
	Min int32 `json:"min"`

	Max int32 `json:"max"`

	DefaultValue int32 `json:"default_value"`

	Unit *string `json:"unit"`
}

// VersionAvailableSettingStringProperty: version available setting string property.
type VersionAvailableSettingStringProperty struct {
	StringConstraint *string `json:"string_constraint"`

	DefaultValue string `json:"default_value"`
}

// EndpointSpecPrivateNetworkDetails: endpoint spec private network details.
type EndpointSpecPrivateNetworkDetails struct {
	// PrivateNetworkID: UUID of the Private Network.
	PrivateNetworkID string `json:"private_network_id"`
}

// EndpointSpecPublicDetails: endpoint spec public details.
type EndpointSpecPublicDetails struct{}

// ClusterSetting: cluster setting.
type ClusterSetting struct {
	// Name: name of the setting.
	Name string `json:"name"`

	// BoolValue: boolean value of the setting.
	// Precisely one of BoolValue, StringValue, IntValue, FloatValue must be set.
	BoolValue *bool `json:"bool_value,omitempty"`

	// StringValue: string value of the setting.
	// Precisely one of BoolValue, StringValue, IntValue, FloatValue must be set.
	StringValue *string `json:"string_value,omitempty"`

	// IntValue: integer value of the setting.
	// Precisely one of BoolValue, StringValue, IntValue, FloatValue must be set.
	IntValue *int32 `json:"int_value,omitempty"`

	// FloatValue: float value of the setting.
	// Precisely one of BoolValue, StringValue, IntValue, FloatValue must be set.
	FloatValue *float32 `json:"float_value,omitempty"`
}

// Endpoint: endpoint.
type Endpoint struct {
	// ID: UUID of the endpoint.
	ID string `json:"id"`

	// DNSRecords: list of DNS records of the endpoint.
	DNSRecords []string `json:"dns_records"`

	// Port: TCP port of the endpoint.
	Port uint32 `json:"port"`

	// PrivateNetwork: private Network endpoint details.
	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PrivateNetwork *EndpointPrivateNetworkDetails `json:"private_network,omitempty"`

	// PublicNetwork: public endpoint details.
	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PublicNetwork *EndpointPublicDetails `json:"public_network,omitempty"`
}

// Volume: volume.
type Volume struct {
	// Type: type of volume where data is stored.
	// Default value: unknown_type
	Type VolumeType `json:"type"`

	// SizeBytes: volume size.
	SizeBytes scw.Size `json:"size_bytes"`
}

// NodeTypeVolumeType: node type volume type.
type NodeTypeVolumeType struct {
	// Type: volume type.
	// Default value: unknown_type
	Type VolumeType `json:"type"`

	// Description: the description of the volume.
	Description string `json:"description"`

	// MinSizeBytes: minimum size required for the volume.
	MinSizeBytes scw.Size `json:"min_size_bytes"`

	// MaxSizeBytes: maximum size required for the volume.
	MaxSizeBytes scw.Size `json:"max_size_bytes"`

	// ChunkSizeBytes: minimum increment level for a Block Storage volume size.
	ChunkSizeBytes scw.Size `json:"chunk_size_bytes"`
}

// VersionAvailableSetting: version available setting.
type VersionAvailableSetting struct {
	// Name: kafka cluster setting name.
	Name string `json:"name"`

	// HotConfigurable: defines whether this setting can be applied without needing a restart.
	HotConfigurable bool `json:"hot_configurable"`

	// Description: setting description.
	Description string `json:"description"`

	// BoolProperty: boolean property, if the setting is a boolean.
	// Precisely one of BoolProperty, StringProperty, IntProperty, FloatProperty must be set.
	BoolProperty *VersionAvailableSettingBooleanProperty `json:"bool_property,omitempty"`

	// StringProperty: string property, if the setting is a string.
	// Precisely one of BoolProperty, StringProperty, IntProperty, FloatProperty must be set.
	StringProperty *VersionAvailableSettingStringProperty `json:"string_property,omitempty"`

	// IntProperty: integer property, if the setting is an integer.
	// Precisely one of BoolProperty, StringProperty, IntProperty, FloatProperty must be set.
	IntProperty *VersionAvailableSettingIntegerProperty `json:"int_property,omitempty"`

	// FloatProperty: float property, if the setting is a float.
	// Precisely one of BoolProperty, StringProperty, IntProperty, FloatProperty must be set.
	FloatProperty *VersionAvailableSettingFloatProperty `json:"float_property,omitempty"`
}

// CreateClusterRequestVolumeSpec: create cluster request volume spec.
type CreateClusterRequestVolumeSpec struct {
	// SizeBytes: volume size.
	SizeBytes scw.Size `json:"size_bytes"`

	// Type: type of volume where data is stored.
	// Default value: unknown_type
	Type VolumeType `json:"type"`
}

// EndpointSpec: endpoint spec.
type EndpointSpec struct {
	// Precisely one of PublicNetwork, PrivateNetwork must be set.
	PublicNetwork *EndpointSpecPublicDetails `json:"public_network,omitempty"`

	// Precisely one of PublicNetwork, PrivateNetwork must be set.
	PrivateNetwork *EndpointSpecPrivateNetworkDetails `json:"private_network,omitempty"`
}

// Cluster: cluster.
type Cluster struct {
	// ID: UUID of the Kafka cluster.
	ID string `json:"id"`

	// Name: name of the Kafka cluster.
	Name string `json:"name"`

	// ProjectID: project ID the Kafka cluster belongs to.
	ProjectID string `json:"project_id"`

	// OrganizationID: organisation ID the Kafka cluster belongs to.
	OrganizationID string `json:"organization_id"`

	// Status: status of the Kafka cluster.
	// Default value: unknown_status
	Status ClusterStatus `json:"status"`

	// Version: kafka version of the Kafka cluster.
	Version string `json:"version"`

	// Tags: list of tags applied to the Kafka cluster.
	Tags []string `json:"tags"`

	// Settings: advanced settings of the Kafka cluster.
	Settings []*ClusterSetting `json:"settings"`

	// NodeAmount: number of nodes in Kafka cluster.
	NodeAmount uint32 `json:"node_amount"`

	// NodeType: node type of the Kafka cluster.
	NodeType string `json:"node_type"`

	// Volume: volumes of the Kafka cluster.
	Volume *Volume `json:"volume"`

	// Endpoints: list of Kafka cluster endpoints.
	Endpoints []*Endpoint `json:"endpoints"`

	// CreatedAt: creation date (must follow the ISO 8601 format).
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date (must follow the ISO 8601 format).
	UpdatedAt *time.Time `json:"updated_at"`

	// Region: region the Kafka cluster is in.
	Region scw.Region `json:"region"`
}

// NodeType: node type.
type NodeType struct {
	// Name: node type name identifier.
	Name string `json:"name"`

	// StockStatus: current stock status for the node type.
	// Default value: unknown_stock
	StockStatus NodeTypeStock `json:"stock_status"`

	// Description: current specifications of the node type offer.
	Description string `json:"description"`

	// Vcpus: number of virtual CPUs of the node type.
	Vcpus uint32 `json:"vcpus"`

	// MemoryBytes: quantity of RAM.
	MemoryBytes scw.Size `json:"memory_bytes"`

	// AvailableVolumeTypes: available storage options for the node type.
	AvailableVolumeTypes []*NodeTypeVolumeType `json:"available_volume_types"`

	// Disabled: defines whether the node type is currently disabled.
	Disabled bool `json:"disabled"`

	// Beta: defines whether the node type is currently in beta.
	Beta bool `json:"beta"`

	// ClusterRange: cluster range associated with the node type offer.
	ClusterRange string `json:"cluster_range"`
}

// User: user.
type User struct {
	Username string `json:"username"`
}

// Version: version.
type Version struct {
	// Version: kafka version.
	Version string `json:"version"`

	// EndOfLifeAt: date of End of Life for the version.
	EndOfLifeAt *time.Time `json:"end_of_life_at"`

	// AvailableSettings: cluster configuration settings you are able to change for clusters running this version. Each item in `available_settings` describes one configurable cluster setting.
	AvailableSettings []*VersionAvailableSetting `json:"available_settings"`
}

// CreateClusterRequest: create cluster request.
type CreateClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: the ID of the Project in which the Kafka cluster will be created.
	ProjectID string `json:"project_id"`

	// Name: name of the Kafka cluster.
	Name string `json:"name"`

	// Version: version of Kafka.
	Version string `json:"version"`

	// Tags: tags to apply to the Kafka cluster.
	Tags []string `json:"tags"`

	// NodeAmount: number of nodes to use for the Kafka cluster.
	NodeAmount uint32 `json:"node_amount"`

	// NodeType: type of node to use for the Kafka cluster.
	NodeType string `json:"node_type"`

	// Volume: kafka volume information.
	Volume *CreateClusterRequestVolumeSpec `json:"volume,omitempty"`

	// Endpoints: one or multiple EndpointSpec used to expose your Kafka cluster.
	Endpoints []*EndpointSpec `json:"endpoints"`

	// UserName: username for the kafka user.
	UserName *string `json:"user_name,omitempty"`

	// Password: password for the kafka user.
	Password *string `json:"password,omitempty"`
}

// CreateEndpointRequest: create endpoint request.
type CreateEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: UUID of the Kafka Cluster.
	ClusterID string `json:"cluster_id"`

	// Endpoint: endpoint object (`EndpointSpec`) used to expose your Kafka EndpointSpec.
	Endpoint *EndpointSpec `json:"endpoint"`
}

// DeleteClusterRequest: delete cluster request.
type DeleteClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: UUID of the Kafka Cluster to delete.
	ClusterID string `json:"-"`
}

// DeleteEndpointRequest: delete endpoint request.
type DeleteEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// EndpointID: UUID of the endpoint to delete.
	EndpointID string `json:"-"`
}

// GetClusterCertificateAuthorityRequest: get cluster certificate authority request.
type GetClusterCertificateAuthorityRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: UUID of the Kafka Cluster.
	ClusterID string `json:"-"`
}

// GetClusterRequest: get cluster request.
type GetClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: UUID of the Kafka Cluster.
	ClusterID string `json:"-"`
}

// ListClustersRequest: list clusters request.
type ListClustersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Tags: list Kafka cluster with a given tag.
	Tags []string `json:"-"`

	// Name: lists Kafka clusters that match a name pattern.
	Name *string `json:"-"`

	// OrderBy: criteria to use when ordering Kafka cluster listings.
	// Default value: created_at_asc
	OrderBy ListClustersRequestOrderBy `json:"-"`

	// OrganizationID: organization ID of the Kafka cluster.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID.
	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListClustersResponse: list clusters response.
type ListClustersResponse struct {
	// Clusters: list of all Kafka cluster available in an Organization or Project.
	Clusters []*Cluster `json:"clusters"`

	// TotalCount: total count of Kafka cluster available in an Organization or Project.
	TotalCount uint64 `json:"total_count"`
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

	// TotalCount: total count of node types available.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNodeTypesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNodeTypesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListNodeTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.NodeTypes = append(r.NodeTypes, results.NodeTypes...)
	r.TotalCount += uint64(len(results.NodeTypes))
	return uint64(len(results.NodeTypes)), nil
}

// ListUsersRequest: list users request.
type ListUsersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ClusterID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListUsersRequestOrderBy `json:"-"`

	Name *string `json:"-"`
}

// ListUsersResponse: list users response.
type ListUsersResponse struct {
	Users []*User `json:"users"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeAppend(res any) (uint64, error) {
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

	// Version: kafka version to filter for.
	Version *string `json:"-"`

	// Page: the page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: the number of items to return.
	PageSize *uint32 `json:"-"`
}

// ListVersionsResponse: list versions response.
type ListVersionsResponse struct {
	// Versions: available Kafka versions.
	Versions []*Version `json:"versions"`

	// TotalCount: total count of Kafka versions available.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVersionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVersionsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Versions = append(r.Versions, results.Versions...)
	r.TotalCount += uint64(len(results.Versions))
	return uint64(len(results.Versions)), nil
}

// RenewClusterCertificateAuthorityRequest: renew cluster certificate authority request.
type RenewClusterCertificateAuthorityRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: UUID of the Kafka Cluster.
	ClusterID string `json:"-"`
}

// UpdateClusterRequest: update cluster request.
type UpdateClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: UUID of the Kafka Clusters to update.
	ClusterID string `json:"-"`

	// Name: name of the Kafka Cluster.
	Name *string `json:"name,omitempty"`

	// Tags: tags of a Kafka Cluster.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateUserRequest: Update a user of a Kafka cluster.
type UpdateUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster in which to update the user's password.
	ClusterID string `json:"-"`

	// Username: username of the Kafka cluster user.
	Username string `json:"-"`

	// Password: new password for the Kafka cluster user.
	Password *string `json:"password,omitempty"`
}

// This API allows you to manage your Clusters for Apache Kafka®. This product is currently in Private Beta.
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
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/node-types",
		Query:  query,
	}

	var resp ListNodeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVersions: List all available versions of Kafka at the current time.
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
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/versions",
		Query:  query,
	}

	var resp ListVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusters: List all Kafka clusters in the specified region. By default, the clusters returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `tags` and `name`. For the `name` parameter, the value you include will be checked against the whole name string to see if it includes the string you put in the parameter.
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
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/clusters",
		Query:  query,
	}

	var resp ListClustersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCluster: Retrieve information about a given Kafka cluster, specified by the `region` and `cluster_id` parameters. Its full details, including name, status, IP address and port, are returned in the response object.
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
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
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
	timeout := defaultKafkaTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultKafkaRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[ClusterStatus]struct{}{
		ClusterStatusCreating:    {},
		ClusterStatusConfiguring: {},
		ClusterStatusDeleting:    {},
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

// CreateCluster: Create a new Kafka cluster.
func (s *API) CreateCluster(req *CreateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
		req.Name = namegenerator.GetRandomName("kafk")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/clusters",
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

// UpdateCluster: Update the parameters of a Kafka cluster.
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
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
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

// DeleteCluster: Delete a given Kafka cluster, specified by the `region` and `cluster_id` parameters. Deleting a Kafka cluster is permanent, and cannot be undone. Note that upon deletion all your data will be lost.
func (s *API) DeleteCluster(req *DeleteClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
		Method: "DELETE",
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetClusterCertificateAuthority: Retrieve certificate authority for a given Kafka cluster, specified by the `region` and `cluster_id` parameters. The response object contains the certificate in PEM format. The certificate is required to validate the sever from the client side during TLS connection.
func (s *API) GetClusterCertificateAuthority(req *GetClusterCertificateAuthorityRequest, opts ...scw.RequestOption) (*scw.File, error) {
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
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/certificate-authority",
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RenewClusterCertificateAuthority: Request to renew the certificate authority for a given Kafka cluster, specified by the `region` and `cluster_id` parameters. The certificate authority will be renewed within a few minutes.
func (s *API) RenewClusterCertificateAuthority(req *RenewClusterCertificateAuthorityRequest, opts ...scw.RequestOption) error {
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
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/renew-certificate-authority",
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

// DeleteEndpoint: Delete the endpoint of a Kafka cluster. You must specify the `endpoint_id` parameter of the endpoint you want to delete. Note that you might need to update any environment configurations that point to the deleted endpoint.
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
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateEndpoint: Create a new endpoint for a Kafka cluster. You can add `public_network` or `private_network` specifications to the body of the request. Note that currently only `private_network` is supported.
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
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/endpoints",
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

// ListUsers: Retrieve a list of deployment users.
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/users",
		Query:  query,
	}

	var resp ListUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateUser: Update an existing user.
func (s *API) UpdateUser(req *UpdateUserRequest, opts ...scw.RequestOption) (*User, error) {
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

	if fmt.Sprint(req.Username) == "" {
		return nil, errors.New("field Username cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/kafka/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/users/" + fmt.Sprint(req.Username) + "",
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

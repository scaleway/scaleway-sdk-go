// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package redis provides methods and message types of the redis v1alpha1 API.
package redis

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

// API: managed Database for Redis™ API
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type AvailableClusterSettingPropertyType string

const (
	// AvailableClusterSettingPropertyTypeBOOLEAN is [insert doc].
	AvailableClusterSettingPropertyTypeBOOLEAN = AvailableClusterSettingPropertyType("BOOLEAN")
	// AvailableClusterSettingPropertyTypeINT is [insert doc].
	AvailableClusterSettingPropertyTypeINT = AvailableClusterSettingPropertyType("INT")
	// AvailableClusterSettingPropertyTypeSTRING is [insert doc].
	AvailableClusterSettingPropertyTypeSTRING = AvailableClusterSettingPropertyType("STRING")
)

func (enum AvailableClusterSettingPropertyType) String() string {
	if enum == "" {
		// return default value if empty
		return "BOOLEAN"
	}
	return string(enum)
}

func (enum AvailableClusterSettingPropertyType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AvailableClusterSettingPropertyType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AvailableClusterSettingPropertyType(AvailableClusterSettingPropertyType(tmp).String())
	return nil
}

type ClusterStatus string

const (
	// ClusterStatusUnknown is [insert doc].
	ClusterStatusUnknown = ClusterStatus("unknown")
	// ClusterStatusReady is [insert doc].
	ClusterStatusReady = ClusterStatus("ready")
	// ClusterStatusProvisioning is [insert doc].
	ClusterStatusProvisioning = ClusterStatus("provisioning")
	// ClusterStatusConfiguring is [insert doc].
	ClusterStatusConfiguring = ClusterStatus("configuring")
	// ClusterStatusDestroying is [insert doc].
	ClusterStatusDestroying = ClusterStatus("destroying")
	// ClusterStatusError is [insert doc].
	ClusterStatusError = ClusterStatus("error")
	// ClusterStatusAutohealing is [insert doc].
	ClusterStatusAutohealing = ClusterStatus("autohealing")
	// ClusterStatusLocked is [insert doc].
	ClusterStatusLocked = ClusterStatus("locked")
	// ClusterStatusSuspended is [insert doc].
	ClusterStatusSuspended = ClusterStatus("suspended")
	// ClusterStatusInitializing is [insert doc].
	ClusterStatusInitializing = ClusterStatus("initializing")
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

type ListClustersRequestOrderBy string

const (
	// ListClustersRequestOrderByCreatedAtAsc is [insert doc].
	ListClustersRequestOrderByCreatedAtAsc = ListClustersRequestOrderBy("created_at_asc")
	// ListClustersRequestOrderByCreatedAtDesc is [insert doc].
	ListClustersRequestOrderByCreatedAtDesc = ListClustersRequestOrderBy("created_at_desc")
	// ListClustersRequestOrderByNameAsc is [insert doc].
	ListClustersRequestOrderByNameAsc = ListClustersRequestOrderBy("name_asc")
	// ListClustersRequestOrderByNameDesc is [insert doc].
	ListClustersRequestOrderByNameDesc = ListClustersRequestOrderBy("name_desc")
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

// ACLRule: acl rule
type ACLRule struct {
	// ID: ID of the rule
	ID string `json:"id"`
	// IP: iPv4 network address of the rule
	IP *scw.IPNet `json:"ip"`
	// Description: description of the rule
	Description *string `json:"description"`
}

// ACLRuleSpec: acl rule spec
type ACLRuleSpec struct {
	// IP: iPv4 network address of the rule
	IP scw.IPNet `json:"ip"`
	// Description: description of the rule
	Description string `json:"description"`
}

// ACLRulesResponse: acl rules response
type ACLRulesResponse struct {
	// ACLRules: ACL Rules enabled on the cluster
	ACLRules []*ACLRule `json:"acl_rules"`
}

// AddEndpointsResponse: add endpoints response
type AddEndpointsResponse struct {
	// Endpoints: endpoints defined on the cluster
	Endpoints []*Endpoint `json:"endpoints"`
	// TotalCount: total count of endpoints of the cluster
	TotalCount uint32 `json:"total_count"`
}

// AvailableClusterSetting: available cluster setting
type AvailableClusterSetting struct {
	// Name: name of the setting
	Name string `json:"name"`
	// DefaultValue: default value of the setting
	DefaultValue *string `json:"default_value"`
	// Type: type of the setting
	//
	// Default value: BOOLEAN
	Type AvailableClusterSettingPropertyType `json:"type"`
	// Description: description of the setting
	Description string `json:"description"`
	// MaxValue: optional maximum value of the setting
	MaxValue *int64 `json:"max_value"`
	// MinValue: optional minimum value of the setting
	MinValue *int64 `json:"min_value"`
	// Regex: optional validation rule of the setting
	Regex *string `json:"regex"`
	// Deprecated: whether the setting is deprecated
	Deprecated bool `json:"deprecated"`
}

// Cluster: cluster
type Cluster struct {
	// ID: UUID of the cluster
	ID string `json:"id"`
	// Name: name of the cluster
	Name string `json:"name"`
	// ProjectID: project ID the cluster belongs to
	ProjectID string `json:"project_id"`
	// Status: status of the cluster
	//
	// Default value: unknown
	Status ClusterStatus `json:"status"`
	// Version: redis™ engine version of the cluster
	Version string `json:"version"`
	// Endpoints: list of cluster endpoints
	Endpoints []*Endpoint `json:"endpoints"`
	// Tags: list of tags applied to the cluster
	Tags []string `json:"tags"`
	// NodeType: node type of the cluster
	NodeType string `json:"node_type"`
	// CreatedAt: creation date (Format ISO 8601)
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: update date (Format ISO 8601)
	UpdatedAt *time.Time `json:"updated_at"`
	// TLSEnabled: whether or not TLS is enabled
	TLSEnabled bool `json:"tls_enabled"`
	// ClusterSettings: list of cluster settings
	ClusterSettings []*ClusterSetting `json:"cluster_settings"`
	// ACLRules: list of acl rules
	ACLRules []*ACLRule `json:"acl_rules"`
	// ClusterSize: number of nodes of the cluster
	ClusterSize uint32 `json:"cluster_size"`
	// Zone: zone of the cluster
	Zone scw.Zone `json:"zone"`
}

// ClusterMetricsResponse: cluster metrics response
type ClusterMetricsResponse struct {
	// Timeseries: time series of metrics of a given cluster
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

// ClusterSetting: cluster setting
type ClusterSetting struct {
	// Value: value of the setting
	Value string `json:"value"`
	// Name: name of the setting
	Name string `json:"name"`
}

// ClusterSettingsResponse: cluster settings response
type ClusterSettingsResponse struct {
	// Settings: settings configured for a given cluster
	Settings []*ClusterSetting `json:"settings"`
}

// ClusterVersion: cluster version
type ClusterVersion struct {
	// Version: redis™ engine version
	Version string `json:"version"`
	// EolDate: end of life date
	EolDate *time.Time `json:"eol_date"`
	// AvailableSettings: cluster settings available to be set
	AvailableSettings []*AvailableClusterSetting `json:"available_settings"`
	// LogoURL: redis™ logo url
	LogoURL string `json:"logo_url"`
	// Zone: zone of the Managed Database for Redis™
	Zone scw.Zone `json:"zone"`
}

// Endpoint: endpoint
type Endpoint struct {
	// Port: TCP port of the endpoint
	Port uint32 `json:"port"`
	// PrivateNetwork: private network details
	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PrivateNetwork *PrivateNetwork `json:"private_network,omitempty"`
	// PublicNetwork: public network details
	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PublicNetwork *EndpointPublicNetwork `json:"public_network,omitempty"`
	// IPs: lis of IPv4 address of the endpoint
	IPs []net.IP `json:"ips"`
	// ID: UUID of the endpoint
	ID string `json:"id"`
}

type EndpointPublicNetwork struct {
}

// EndpointSpec: endpoint spec
type EndpointSpec struct {
	// PrivateNetwork: private network spec details
	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PrivateNetwork *EndpointSpecPrivateNetworkSpec `json:"private_network,omitempty"`
	// PublicNetwork: public network spec details
	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PublicNetwork *EndpointSpecPublicNetworkSpec `json:"public_network,omitempty"`
}

// EndpointSpecPrivateNetworkSpec: endpoint spec. private network spec
type EndpointSpecPrivateNetworkSpec struct {
	// ID: UUID of the private network to be connected to the cluster
	ID string `json:"id"`
	// ServiceIPs: endpoint IPv4 adress with a CIDR notation. You must provide at least one IPv4 per node. Check documentation about IP and subnet limitation.
	ServiceIPs []scw.IPNet `json:"service_ips"`
}

// EndpointSpecPublicNetworkSpec: endpoint spec. public network spec
type EndpointSpecPublicNetworkSpec struct {
}

// ListClustersResponse: list clusters response
type ListClustersResponse struct {
	// Clusters: list all clusters
	Clusters []*Cluster `json:"clusters"`
	// TotalCount: total count of clusters
	TotalCount uint32 `json:"total_count"`
}

// ListNodeTypesResponse: list node types response
type ListNodeTypesResponse struct {
	// NodeTypes: types of the node
	NodeTypes []*NodeType `json:"node_types"`
	// TotalCount: total count of node-types available
	TotalCount uint32 `json:"total_count"`
}

// ListVersionsResponse: list versions response
type ListVersionsResponse struct {
	// Versions: list of the available Redis™ engine versions
	Versions []*ClusterVersion `json:"versions"`
	// TotalCount: total count of available Redis™ engine versions
	TotalCount uint32 `json:"total_count"`
}

// NodeType: node type
type NodeType struct {
	// Name: node Type name identifier
	Name string `json:"name"`
	// StockStatus: current stock status for the Node Type
	//
	// Default value: unknown
	StockStatus NodeTypeStock `json:"stock_status"`
	// Description: current specs of the offer
	Description string `json:"description"`
	// Vcpus: number of virtual CPUs
	Vcpus uint32 `json:"vcpus"`
	// Memory: quantity of RAM
	Memory scw.Size `json:"memory"`
	// Disabled: the Node Type is currently disabled
	Disabled bool `json:"disabled"`
	// Beta: the Node Type is currently in beta
	Beta bool `json:"beta"`
	// Zone: zone the Node Type is in
	Zone scw.Zone `json:"zone"`
}

// PrivateNetwork: private network
type PrivateNetwork struct {
	// ID: UUID of the private network
	ID string `json:"id"`
	// ServiceIPs: list of IPv4 CIDR notation addresses of the endpoint
	ServiceIPs []scw.IPNet `json:"service_ips"`
	// Zone: private network zone
	Zone scw.Zone `json:"zone"`
}

// SetEndpointsResponse: set endpoints response
type SetEndpointsResponse struct {
	// Endpoints: endpoints defined on the cluster
	Endpoints []*Endpoint `json:"endpoints"`
	// TotalCount: total count of endpoints of the cluster
	TotalCount uint32 `json:"total_count"`
}

// Service API

type CreateClusterRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ProjectID: the project ID on which to create the cluster
	ProjectID string `json:"project_id"`
	// Name: name of the cluster
	Name string `json:"name"`
	// Version: redis™ engine version of the cluster
	Version string `json:"version"`
	// Tags: tags to apply to the cluster
	Tags []string `json:"tags"`
	// NodeType: type of node to use for the cluster
	NodeType string `json:"node_type"`
	// UserName: name of the user created when the cluster is created
	UserName string `json:"user_name"`
	// Password: password of the user
	Password string `json:"password"`
	// ClusterSize: number of nodes for the cluster
	ClusterSize *int32 `json:"cluster_size"`
	// ACLRules: list of ACLRuleSpec used to secure your publicly exposed cluster
	ACLRules []*ACLRuleSpec `json:"acl_rules"`
	// Endpoints: zero or multiple EndpointSpec used to expose your cluster publicly and inside private networks
	//
	// Zero or multiple EndpointSpec used to expose your cluster publicly and inside private networks. If no EndpoindSpec is given the cluster will be publicly exposed by default.
	Endpoints []*EndpointSpec `json:"endpoints"`
	// TLSEnabled: whether or not TLS is enabled
	TLSEnabled bool `json:"tls_enabled"`
	// ClusterSettings: list of cluster settings to be set at cluster initialisation
	ClusterSettings []*ClusterSetting `json:"cluster_settings"`
}

// CreateCluster: create a cluster
func (s *API) CreateCluster(req *CreateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("ins")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters",
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

type UpdateClusterRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster to update
	ClusterID string `json:"-"`
	// Name: name of the cluster
	Name *string `json:"name"`
	// Tags: tags of a given cluster
	Tags *[]string `json:"tags"`
	// UserName: name of the cluster user
	UserName *string `json:"user_name"`
	// Password: password of the cluster user
	Password *string `json:"password"`
}

// UpdateCluster: update a cluster
func (s *API) UpdateCluster(req *UpdateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
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
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster
	ClusterID string `json:"-"`
}

// GetCluster: get a cluster
func (s *API) GetCluster(req *GetClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
		Headers: http.Header{},
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListClustersRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// Tags: tags of the clusters to filter upon
	Tags []string `json:"-"`
	// Name: name of the clusters to filter upon
	Name *string `json:"-"`
	// OrderBy: criteria to use when ordering cluster listing
	//
	// Default value: created_at_asc
	OrderBy ListClustersRequestOrderBy `json:"-"`
	// ProjectID: project ID to list the cluster of
	ProjectID *string `json:"-"`
	// OrganizationID: organization ID to list the cluster of
	OrganizationID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListClusters: list clusters
func (s *API) ListClusters(req *ListClustersRequest, opts ...scw.RequestOption) (*ListClustersResponse, error) {
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
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters",
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

type MigrateClusterRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster to update
	ClusterID string `json:"-"`
	// Version: redis™ engine version of the cluster
	// Precisely one of ClusterSize, NodeType, Version must be set.
	Version *string `json:"version,omitempty"`
	// NodeType: type of node to use for the cluster
	// Precisely one of ClusterSize, NodeType, Version must be set.
	NodeType *string `json:"node_type,omitempty"`
	// ClusterSize: number of nodes for the cluster
	// Precisely one of ClusterSize, NodeType, Version must be set.
	ClusterSize *uint32 `json:"cluster_size,omitempty"`
}

// MigrateCluster: migrate your cluster architecture
//
// Upgrade your Database for Redis® cluster to a new version or scale it vertically / horizontally. Please note: scaling horizontally your Database for Redis® cluster won't renew its TLS certificate. In order to refresh the SSL certificate, you have to use the dedicated api route.
func (s *API) MigrateCluster(req *MigrateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/migrate",
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
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster to delete
	ClusterID string `json:"-"`
}

// DeleteCluster: delete a cluster
func (s *API) DeleteCluster(req *DeleteClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
		Headers: http.Header{},
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetClusterMetricsRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster
	ClusterID string `json:"-"`
	// StartDate: start date to gather metrics from
	StartDate *time.Time `json:"-"`
	// EndDate: end date to gather metrics from
	EndDate *time.Time `json:"-"`
	// MetricName: name of the metric to gather
	MetricName *string `json:"-"`
}

// GetClusterMetrics: get metrics of a cluster
func (s *API) GetClusterMetrics(req *GetClusterMetricsRequest, opts ...scw.RequestOption) (*ClusterMetricsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)
	parameter.AddToQuery(query, "metric_name", req.MetricName)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/metrics",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ClusterMetricsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListNodeTypesRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// IncludeDisabledTypes: whether or not to include disabled types
	IncludeDisabledTypes bool `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListNodeTypes: list available node types
func (s *API) ListNodeTypes(req *ListNodeTypesRequest, opts ...scw.RequestOption) (*ListNodeTypesResponse, error) {
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
	parameter.AddToQuery(query, "include_disabled_types", req.IncludeDisabledTypes)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/node-types",
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

type ListVersionsRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// IncludeDisabled: whether or not to include disabled Redis™ engine versions
	IncludeDisabled bool `json:"-"`
	// IncludeBeta: whether or not to include beta Redis™ engine versions
	IncludeBeta bool `json:"-"`
	// IncludeDeprecated: whether or not to include deprecated Redis™ engine versions
	IncludeDeprecated bool `json:"-"`
	// VersionName: list Redis™ engine versions that match a given name pattern
	VersionName *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListVersions: list available Redis™ versions
func (s *API) ListVersions(req *ListVersionsRequest, opts ...scw.RequestOption) (*ListVersionsResponse, error) {
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
	parameter.AddToQuery(query, "include_disabled", req.IncludeDisabled)
	parameter.AddToQuery(query, "include_beta", req.IncludeBeta)
	parameter.AddToQuery(query, "include_deprecated", req.IncludeDeprecated)
	parameter.AddToQuery(query, "version_name", req.VersionName)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/versions",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetClusterCertificateRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster
	ClusterID string `json:"-"`
}

// GetClusterCertificate: get the TLS certificate of a cluster
func (s *API) GetClusterCertificate(req *GetClusterCertificateRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/certificate",
		Headers: http.Header{},
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RenewClusterCertificateRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster
	ClusterID string `json:"-"`
}

// RenewClusterCertificate: renew the TLS certificate of a cluster
func (s *API) RenewClusterCertificate(req *RenewClusterCertificateRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/renew-certificate",
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

type AddClusterSettingsRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster you want to add settings to
	ClusterID string `json:"-"`
	// Settings: settings to add on the cluster
	Settings []*ClusterSetting `json:"settings"`
}

// AddClusterSettings: add cluster settings
func (s *API) AddClusterSettings(req *AddClusterSettingsRequest, opts ...scw.RequestOption) (*ClusterSettingsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/settings",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ClusterSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteClusterSettingRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster where the settings has to be set
	ClusterID string `json:"-"`
	// SettingsName: setting name to delete
	SettingsName string `json:"settings_name"`
}

// DeleteClusterSetting: delete a cluster setting
func (s *API) DeleteClusterSetting(req *DeleteClusterSettingRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/settings",
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

type SetClusterSettingsRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster where the settings has to be set
	ClusterID string `json:"-"`
	// Settings: settings to define for the cluster
	Settings []*ClusterSetting `json:"settings"`
}

// SetClusterSettings: set cluster settings
func (s *API) SetClusterSettings(req *SetClusterSettingsRequest, opts ...scw.RequestOption) (*ClusterSettingsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/settings",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ClusterSettingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetACLRulesRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster where the ACL rules has to be set
	ClusterID string `json:"-"`
	// ACLRules: aCLs rules to define for the cluster
	ACLRules []*ACLRuleSpec `json:"acl_rules"`
}

// SetACLRules: set ACL rules for a given cluster
func (s *API) SetACLRules(req *SetACLRulesRequest, opts ...scw.RequestOption) (*ACLRulesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/acls",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddACLRulesRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster you want to add acl rules to
	ClusterID string `json:"-"`
	// ACLRules: aCLs rules to add to the cluster
	ACLRules []*ACLRuleSpec `json:"acl_rules"`
}

// AddACLRules: add ACL rules for a given cluster
func (s *API) AddACLRules(req *AddACLRulesRequest, opts ...scw.RequestOption) (*ACLRulesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/acls",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteACLRuleRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ACLID: UUID of the acl rule you want to delete
	ACLID string `json:"-"`
}

// DeleteACLRule: delete an ACL rule for a given cluster
func (s *API) DeleteACLRule(req *DeleteACLRuleRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/acls/" + fmt.Sprint(req.ACLID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetACLRuleRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ACLID: UUID of the acl rule you want to get
	ACLID string `json:"-"`
}

// GetACLRule: get an ACL rule
func (s *API) GetACLRule(req *GetACLRuleRequest, opts ...scw.RequestOption) (*ACLRule, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return nil, errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/acls/" + fmt.Sprint(req.ACLID) + "",
		Headers: http.Header{},
	}

	var resp ACLRule

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetEndpointsRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster where the endpoints has to be set
	ClusterID string `json:"-"`
	// Endpoints: endpoints to define for the cluster
	Endpoints []*EndpointSpec `json:"endpoints"`
}

// SetEndpoints: set endpoints for a given cluster
func (s *API) SetEndpoints(req *SetEndpointsRequest, opts ...scw.RequestOption) (*SetEndpointsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/endpoints",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetEndpointsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type AddEndpointsRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ClusterID: UUID of the cluster you want to add endpoints to
	ClusterID string `json:"-"`
	// Endpoints: endpoints to add to the cluster
	Endpoints []*EndpointSpec `json:"endpoints"`
}

// AddEndpoints: add endpoints for a given cluster
func (s *API) AddEndpoints(req *AddEndpointsRequest, opts ...scw.RequestOption) (*AddEndpointsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/endpoints",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddEndpointsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteEndpointRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// EndpointID: UUID of the endpoint you want to delete
	EndpointID string `json:"-"`
}

// DeleteEndpoint: delete an endpoint for a given cluster
func (s *API) DeleteEndpoint(req *DeleteEndpointRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetEndpointRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// EndpointID: UUID of the endpoint you want to get
	EndpointID string `json:"-"`
}

// GetEndpoint: get an endpoint
func (s *API) GetEndpoint(req *GetEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return nil, errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
		Headers: http.Header{},
	}

	var resp Endpoint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateEndpointRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`

	EndpointID string `json:"-"`

	Endpoint *EndpointSpec `json:"endpoint"`
}

func (s *API) UpdateEndpoint(req *UpdateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.EndpointID) == "" {
		return nil, errors.New("field EndpointID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
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
func (r *ListVersionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVersionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Versions = append(r.Versions, results.Versions...)
	r.TotalCount += uint32(len(results.Versions))
	return uint32(len(results.Versions)), nil
}

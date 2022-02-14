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

// API: database Redis API
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
	ID string `json:"id"`

	IP *scw.IPNet `json:"ip"`

	Description *string `json:"description"`
}

// ACLRuleResponse: acl rule response
type ACLRuleResponse struct {
	ACLRule *ACLRule `json:"acl_rule"`
}

// ACLRuleSpec: acl rule spec
type ACLRuleSpec struct {
	IP scw.IPNet `json:"ip"`

	Description string `json:"description"`
}

// ACLRulesResponse: acl rules response
type ACLRulesResponse struct {
	ACLRules []*ACLRule `json:"acl_rules"`
}

// AvailableClusterSetting: available cluster setting
type AvailableClusterSetting struct {
	Name string `json:"name"`

	DefaultValue *string `json:"default_value"`
	// Type:
	//
	// Default value: BOOLEAN
	Type AvailableClusterSettingPropertyType `json:"type"`

	Description string `json:"description"`

	MaxValue *int64 `json:"max_value"`

	MinValue *int64 `json:"min_value"`

	Regex *string `json:"regex"`

	Deprecated bool `json:"deprecated"`
}

// Cluster: cluster
type Cluster struct {
	ID string `json:"id"`

	Name string `json:"name"`

	ProjectID string `json:"project_id"`
	// Status:
	//
	// Default value: unknown
	Status ClusterStatus `json:"status"`

	Version string `json:"version"`

	Endpoints []*Endpoint `json:"endpoints"`

	Tags []string `json:"tags"`

	NodeType string `json:"node_type"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	TLSEnabled bool `json:"tls_enabled"`

	ClusterSettings []*ClusterSetting `json:"cluster_settings"`

	ACLRules []*ACLRule `json:"acl_rules"`

	ClusterSize int32 `json:"cluster_size"`

	Zone scw.Zone `json:"zone"`
}

// ClusterMetricsResponse: cluster metrics response
type ClusterMetricsResponse struct {
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

// ClusterSetting: cluster setting
type ClusterSetting struct {
	Value string `json:"value"`

	Name string `json:"name"`
}

// ClusterSettingsResponse: cluster settings response
type ClusterSettingsResponse struct {
	Settings []*ClusterSetting `json:"settings"`
}

// ClusterVersion: cluster version
type ClusterVersion struct {
	Version string `json:"version"`

	EolDate *time.Time `json:"eol_date"`

	AvailableSettings []*AvailableClusterSetting `json:"available_settings"`

	LogoURL string `json:"logo_url"`

	Zone scw.Zone `json:"zone"`
}

// Endpoint: endpoint
type Endpoint struct {
	// Port: port to connect to redis server
	Port uint32 `json:"port"`
	// PrivateNetwork: private network details for endpoint
	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PrivateNetwork *PrivateNetwork `json:"private_network,omitempty"`
	// ID: UUID of current endpoint (Used for edit or delete PN endpoint)
	ID string `json:"id"`
	// PublicNetwork: public network details for endpoint
	// Precisely one of PrivateNetwork, PublicNetwork must be set.
	PublicNetwork *PublicDetails `json:"public_network,omitempty"`
}

// EndpointSpec: endpoint spec
type EndpointSpec struct {
	// ID: put UUID of the endpoint you want to edit or do not fill this field to create a new endpoint
	ID *string `json:"id"`
	// PrivateNetwork: private network details for endpoint
	PrivateNetwork *PrivateNetworkSpec `json:"private_network"`
}

// ListClustersResponse: list clusters response
type ListClustersResponse struct {
	Clusters []*Cluster `json:"clusters"`

	TotalCount uint32 `json:"total_count"`
}

// ListNodeTypesResponse: list node types response
type ListNodeTypesResponse struct {
	NodeTypes []*NodeType `json:"node_types"`

	TotalCount uint32 `json:"total_count"`
}

// ListVersionsResponse: list versions response
type ListVersionsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Versions []*ClusterVersion `json:"versions"`
}

// NodeType: node type
type NodeType struct {
	Name string `json:"name"`
	// StockStatus:
	//
	// Default value: unknown
	StockStatus NodeTypeStock `json:"stock_status"`

	Description string `json:"description"`

	Vcpus uint32 `json:"vcpus"`

	Memory scw.Size `json:"memory"`

	Disabled bool `json:"disabled"`

	Beta bool `json:"beta"`

	Zone scw.Zone `json:"zone"`
}

type PrivateNetwork struct {
	ID string `json:"id"`

	ServiceIPs []scw.IPNet `json:"service_ips"`

	Zone scw.Zone `json:"zone"`
}

// PrivateNetworkSpec: private network spec
type PrivateNetworkSpec struct {
	// ID: put UUID of the endpoint you want to connect cluster
	ID string `json:"id"`
	// ServiceIPs: put a list of IPv4 in CIDR format to expose the cluster in the private network. You must provide one IPv4 per node
	ServiceIPs []scw.IPNet `json:"service_ips"`
}

// PublicDetails: public details
type PublicDetails struct {
	IPs []net.IP `json:"ips"`
}

// Service API

type GetServiceInfoRequest struct {
	Zone scw.Zone `json:"-"`
}

func (s *API) GetServiceInfo(req *GetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "",
		Headers: http.Header{},
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateClusterRequest struct {
	Zone scw.Zone `json:"-"`

	ProjectID string `json:"project_id"`

	Name string `json:"name"`

	Version string `json:"version"`

	Tags []string `json:"tags"`

	NodeType string `json:"node_type"`

	UserName string `json:"user_name"`

	Password string `json:"password"`

	ClusterSize *int32 `json:"cluster_size"`

	ACLRules []*ACLRuleSpec `json:"acl_rules"`

	Endpoints []*EndpointSpec `json:"endpoints"`

	TLSEnabled bool `json:"tls_enabled"`

	ClusterSettings []*ClusterSetting `json:"cluster_settings"`
}

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
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	Name *string `json:"name"`

	Tags []string `json:"tags"`

	UserName *string `json:"user_name"`

	Password *string `json:"password"`
}

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
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`
}

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
	Zone scw.Zone `json:"-"`

	Tags []string `json:"-"`

	Name *string `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListClustersRequestOrderBy `json:"-"`

	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

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

type AlterClusterRequest struct {
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	Version *string `json:"version"`

	NodeType *string `json:"node_type"`

	ClusterSize *int32 `json:"cluster_size"`

	Endpoints []*EndpointSpec `json:"endpoints"`
}

func (s *API) AlterCluster(req *AlterClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/alter",
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
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`
}

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

type ListNodeTypesRequest struct {
	Zone scw.Zone `json:"-"`

	IncludeDisabledTypes bool `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

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
	Zone scw.Zone `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	IncludeDisabled bool `json:"-"`

	IncludeBeta bool `json:"-"`

	IncludeDeprecated bool `json:"-"`

	VersionName *string `json:"-"`
}

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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "include_disabled", req.IncludeDisabled)
	parameter.AddToQuery(query, "include_beta", req.IncludeBeta)
	parameter.AddToQuery(query, "include_deprecated", req.IncludeDeprecated)
	parameter.AddToQuery(query, "version_name", req.VersionName)

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
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`
}

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
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`
}

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

type GetClusterMetricsRequest struct {
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	StartDate *time.Time `json:"-"`

	EndDate *time.Time `json:"-"`

	MetricName *string `json:"-"`
}

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

type AddClusterSettingsRequest struct {
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	Settings []*ClusterSetting `json:"settings"`
}

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

type DeleteClusterSettingsRequest struct {
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	SettingsName []string `json:"settings_name"`
}

func (s *API) DeleteClusterSettings(req *DeleteClusterSettingsRequest, opts ...scw.RequestOption) (*ClusterSettingsResponse, error) {
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

type SetClusterSettingsRequest struct {
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	Settings []*ClusterSetting `json:"settings"`
}

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
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	ACLRules []*ACLRuleSpec `json:"acl_rules"`
}

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
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	ACLRules []*ACLRuleSpec `json:"acl_rules"`
}

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
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	ACLID string `json:"-"`
}

func (s *API) DeleteACLRule(req *DeleteACLRuleRequest, opts ...scw.RequestOption) (*ACLRuleResponse, error) {
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

	if fmt.Sprint(req.ACLID) == "" {
		return nil, errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/redis/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/acls/" + fmt.Sprint(req.ACLID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ACLRuleResponse

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

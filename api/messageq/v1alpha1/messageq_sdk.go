// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package messageq provides methods and message types of the messageq v1alpha1 API.
package messageq

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
	defaultMessageqRetryInterval = 15 * time.Second
	defaultMessageqTimeout       = 15 * time.Minute
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

type DeploymentStatus string

const (
	DeploymentStatusUnknownStatus = DeploymentStatus("unknown_status")
	DeploymentStatusReady         = DeploymentStatus("ready")
	DeploymentStatusCreating      = DeploymentStatus("creating")
	DeploymentStatusInitializing  = DeploymentStatus("initializing")
	DeploymentStatusUpgrading     = DeploymentStatus("upgrading")
	DeploymentStatusDeleting      = DeploymentStatus("deleting")
	DeploymentStatusError         = DeploymentStatus("error")
	DeploymentStatusLocked        = DeploymentStatus("locked")
	DeploymentStatusLocking       = DeploymentStatus("locking")
	DeploymentStatusUnlocking     = DeploymentStatus("unlocking")
)

func (enum DeploymentStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(DeploymentStatusUnknownStatus)
	}
	return string(enum)
}

func (enum DeploymentStatus) Values() []DeploymentStatus {
	return []DeploymentStatus{
		"unknown_status",
		"ready",
		"creating",
		"initializing",
		"upgrading",
		"deleting",
		"error",
		"locked",
		"locking",
		"unlocking",
	}
}

func (enum DeploymentStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DeploymentStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DeploymentStatus(DeploymentStatus(tmp).String())
	return nil
}

type ListDeploymentsRequestOrderBy string

const (
	ListDeploymentsRequestOrderByCreatedAtAsc  = ListDeploymentsRequestOrderBy("created_at_asc")
	ListDeploymentsRequestOrderByCreatedAtDesc = ListDeploymentsRequestOrderBy("created_at_desc")
	ListDeploymentsRequestOrderByNameAsc       = ListDeploymentsRequestOrderBy("name_asc")
	ListDeploymentsRequestOrderByNameDesc      = ListDeploymentsRequestOrderBy("name_desc")
	ListDeploymentsRequestOrderByUpdatedAtAsc  = ListDeploymentsRequestOrderBy("updated_at_asc")
	ListDeploymentsRequestOrderByUpdatedAtDesc = ListDeploymentsRequestOrderBy("updated_at_desc")
)

func (enum ListDeploymentsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListDeploymentsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListDeploymentsRequestOrderBy) Values() []ListDeploymentsRequestOrderBy {
	return []ListDeploymentsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
		"updated_at_asc",
		"updated_at_desc",
	}
}

func (enum ListDeploymentsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDeploymentsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDeploymentsRequestOrderBy(ListDeploymentsRequestOrderBy(tmp).String())
	return nil
}

type ListNodeTypesRequestOrderBy string

const (
	ListNodeTypesRequestOrderByNameAsc    = ListNodeTypesRequestOrderBy("name_asc")
	ListNodeTypesRequestOrderByNameDesc   = ListNodeTypesRequestOrderBy("name_desc")
	ListNodeTypesRequestOrderByVcpusAsc   = ListNodeTypesRequestOrderBy("vcpus_asc")
	ListNodeTypesRequestOrderByVcpusDesc  = ListNodeTypesRequestOrderBy("vcpus_desc")
	ListNodeTypesRequestOrderByMemoryAsc  = ListNodeTypesRequestOrderBy("memory_asc")
	ListNodeTypesRequestOrderByMemoryDesc = ListNodeTypesRequestOrderBy("memory_desc")
)

func (enum ListNodeTypesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListNodeTypesRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListNodeTypesRequestOrderBy) Values() []ListNodeTypesRequestOrderBy {
	return []ListNodeTypesRequestOrderBy{
		"name_asc",
		"name_desc",
		"vcpus_asc",
		"vcpus_desc",
		"memory_asc",
		"memory_desc",
	}
}

func (enum ListNodeTypesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNodeTypesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNodeTypesRequestOrderBy(ListNodeTypesRequestOrderBy(tmp).String())
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

type ListVersionsRequestOrderBy string

const (
	ListVersionsRequestOrderByVersionAsc  = ListVersionsRequestOrderBy("version_asc")
	ListVersionsRequestOrderByVersionDesc = ListVersionsRequestOrderBy("version_desc")
)

func (enum ListVersionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListVersionsRequestOrderByVersionAsc)
	}
	return string(enum)
}

func (enum ListVersionsRequestOrderBy) Values() []ListVersionsRequestOrderBy {
	return []ListVersionsRequestOrderBy{
		"version_asc",
		"version_desc",
	}
}

func (enum ListVersionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListVersionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListVersionsRequestOrderBy(ListVersionsRequestOrderBy(tmp).String())
	return nil
}

type NodeTypeStockStatus string

const (
	NodeTypeStockStatusUnknownStock = NodeTypeStockStatus("unknown_stock")
	NodeTypeStockStatusLowStock     = NodeTypeStockStatus("low_stock")
	NodeTypeStockStatusOutOfStock   = NodeTypeStockStatus("out_of_stock")
	NodeTypeStockStatusAvailable    = NodeTypeStockStatus("available")
)

func (enum NodeTypeStockStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(NodeTypeStockStatusUnknownStock)
	}
	return string(enum)
}

func (enum NodeTypeStockStatus) Values() []NodeTypeStockStatus {
	return []NodeTypeStockStatus{
		"unknown_stock",
		"low_stock",
		"out_of_stock",
		"available",
	}
}

func (enum NodeTypeStockStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NodeTypeStockStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NodeTypeStockStatus(NodeTypeStockStatus(tmp).String())
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

// EndpointPrivateNetworkDetails: endpoint private network details.
type EndpointPrivateNetworkDetails struct {
	PrivateNetworkID string `json:"private_network_id"`
}

// EndpointPublicDetails: endpoint public details.
type EndpointPublicDetails struct{}

// EndpointService: endpoint service.
type EndpointService struct {
	Name string `json:"name"`

	Port uint32 `json:"port"`

	URL string `json:"url"`
}

// EndpointSpecPrivateNetworkDetails: endpoint spec private network details.
type EndpointSpecPrivateNetworkDetails struct {
	PrivateNetworkID string `json:"private_network_id"`
}

// EndpointSpecPublicDetails: endpoint spec public details.
type EndpointSpecPublicDetails struct{}

// Endpoint: endpoint.
type Endpoint struct {
	// ID: unique identifier of the endpoint.
	ID string `json:"id"`

	// Deprecated: DNSRecord: DNS record for service access. Now deprecated. Use the `url` field from `services` field instead.
	DNSRecord *string `json:"dns_record,omitempty"`

	// Services: list of available services, their ports and URLs.
	Services []*EndpointService `json:"services"`

	// Precisely one of Public, PrivateNetwork must be set.
	Public *EndpointPublicDetails `json:"public,omitempty"`

	// Precisely one of Public, PrivateNetwork must be set.
	PrivateNetwork *EndpointPrivateNetworkDetails `json:"private_network,omitempty"`
}

// Volume: volume.
type Volume struct {
	// Type: type of the Volume.
	// Default value: unknown_type
	Type VolumeType `json:"type"`

	// SizeBytes: size of the Volume.
	SizeBytes scw.Size `json:"size_bytes"`
}

// NodeTypeVolumeType: node type volume type.
type NodeTypeVolumeType struct {
	// Type: default value: unknown_type
	Type VolumeType `json:"type"`

	Description string `json:"description"`

	MinSizeBytes scw.Size `json:"min_size_bytes"`

	MaxSizeBytes scw.Size `json:"max_size_bytes"`

	ChunkSizeBytes scw.Size `json:"chunk_size_bytes"`
}

// EndpointSpec: endpoint spec.
type EndpointSpec struct {
	// Precisely one of Public, PrivateNetwork must be set.
	Public *EndpointSpecPublicDetails `json:"public,omitempty"`

	// Precisely one of Public, PrivateNetwork must be set.
	PrivateNetwork *EndpointSpecPrivateNetworkDetails `json:"private_network,omitempty"`
}

// Deployment: deployment.
type Deployment struct {
	// ID: unique identifier of the deployment.
	ID string `json:"id"`

	// Name: name of the deployment.
	Name string `json:"name"`

	// OrganizationID: ID of the Organization containing the deployment.
	OrganizationID string `json:"organization_id"`

	// ProjectID: ID of the Project containing the deployment.
	ProjectID string `json:"project_id"`

	// Status: status of the deployment.
	// Default value: unknown_status
	Status DeploymentStatus `json:"status"`

	// Tags: tags of the deployment.
	Tags []string `json:"tags"`

	// NodeCount: number of nodes.
	NodeCount uint32 `json:"node_count"`

	// NodeType: node type used.
	NodeType string `json:"node_type"`

	// Volume: volume type and size.
	Volume *Volume `json:"volume"`

	// Endpoints: exposed endpoints.
	Endpoints []*Endpoint `json:"endpoints"`

	// CreatedAt: date the deployment was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the deployment was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// Version: messageQ version of the deployment.
	Version string `json:"version"`

	// Region: region of the deployment.
	Region scw.Region `json:"region"`
}

// NodeType: node type.
type NodeType struct {
	// StockStatus: stock status for the node type.
	// Default value: unknown_stock
	StockStatus NodeTypeStockStatus `json:"stock_status"`

	// Name: name of the node type.
	Name string `json:"name"`

	// Description: description of the node type.
	Description string `json:"description"`

	// Vcpus: number of vCPUs available.
	Vcpus uint32 `json:"vcpus"`

	// MemoryBytes: amount of memory available.
	MemoryBytes scw.Size `json:"memory_bytes"`

	// Disabled: defines whether the node type is disabled.
	Disabled bool `json:"disabled"`

	// Beta: defines whether the node type is in beta.
	Beta bool `json:"beta"`

	// InstanceRange: instance range associated with the node type offer.
	InstanceRange string `json:"instance_range"`

	// AvailableVolumeTypes: available storage options for the node type.
	AvailableVolumeTypes []*NodeTypeVolumeType `json:"available_volume_types"`
}

// User: user.
type User struct {
	Username string `json:"username"`
}

// Version: version.
type Version struct {
	// Version: messageQ version.
	Version string `json:"version"`

	// EndOfLife: date the version support ends.
	EndOfLife *time.Time `json:"end_of_life"`

	// Disabled: defines whether the version is disabled.
	Disabled bool `json:"disabled"`

	// Beta: defines whether the version is in beta.
	Beta bool `json:"beta"`
}

// CreateDeploymentRequest: Request to create a new deployment.
type CreateDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project ID in which to create the deployment.
	ProjectID string `json:"project_id"`

	// Name: name of the deployment.
	Name string `json:"name"`

	// Tags: tags.
	Tags []string `json:"tags"`

	// NodeCount: number of nodes.
	NodeCount uint32 `json:"node_count"`

	// NodeType: node type to use.
	NodeType string `json:"node_type"`

	// UserName: username for the deployment user.
	UserName *string `json:"user_name,omitempty"`

	// Password: password for the deployment user.
	Password *string `json:"password,omitempty"`

	// Volume: volume for storing data.
	Volume *Volume `json:"volume,omitempty"`

	// Endpoints: endpoints to access the deployment.
	Endpoints []*EndpointSpec `json:"endpoints"`

	// Version: the MessageQ version to use.
	Version string `json:"version"`
}

// CreateEndpointRequest: Create an endpoint for a specific deployment.
type CreateEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment.
	DeploymentID string `json:"deployment_id"`

	// EndpointSpec: specification of the endpoint you want to create.
	EndpointSpec *EndpointSpec `json:"endpoint_spec,omitempty"`
}

// CreateUserRequest: Create a user for a deployment.
type CreateUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment.
	DeploymentID string `json:"-"`

	// Username: username of the deployment user.
	Username string `json:"username"`

	// Password: password of the deployment user.
	Password string `json:"password"`
}

// DeleteDeploymentRequest: Delete a deployment specified by the ID.
type DeleteDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment.
	DeploymentID string `json:"-"`
}

// DeleteEndpointRequest: Delete an endpoint from a specific deployment.
type DeleteEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// EndpointID: ID of the endpoint.
	EndpointID string `json:"-"`
}

// DeleteUserRequest: Delete a user from a deployment.
type DeleteUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment.
	DeploymentID string `json:"-"`

	// Username: username of the deployment user.
	Username string `json:"-"`
}

// GetDeploymentRequest: Retrieve a deployment specified by the ID.
type GetDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment.
	DeploymentID string `json:"-"`
}

// ListDeploymentsRequest: Retrieve a list of deployments.
type ListDeploymentsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrganizationID: organization ID to filter for, only deployments from this Organization will be returned.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID to filter for, only deployments from this Project will be returned.
	ProjectID *string `json:"-"`

	// OrderBy: sort order for deployments in the response.
	// Default value: created_at_asc
	OrderBy ListDeploymentsRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of deployments to return per page.
	PageSize *uint32 `json:"-"`

	// Tags: tags to filter for, only deployments with one or more matching tags will be returned.
	Tags []string `json:"-"`

	// Name: deployment name to filter for, only deployments with this string within their name will be returned.
	Name *string `json:"-"`

	// Version: engine version to filter for, only deployments with this version will be returned.
	Version *string `json:"-"`
}

// ListDeploymentsResponse: list deployments response.
type ListDeploymentsResponse struct {
	// Deployments: list of deployments.
	Deployments []*Deployment `json:"deployments"`

	// TotalCount: number of deployments in result set.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDeploymentsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDeploymentsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListDeploymentsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Deployments = append(r.Deployments, results.Deployments...)
	r.TotalCount += uint64(len(results.Deployments))
	return uint64(len(results.Deployments)), nil
}

// ListNodeTypesRequest: Retrieve a list of available node types for a MessageQ deployment.
type ListNodeTypesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: sort order for versions in the response.
	// Default value: name_asc
	OrderBy ListNodeTypesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of node types to return per page.
	PageSize *uint32 `json:"-"`
}

// ListNodeTypesResponse: list node types response.
type ListNodeTypesResponse struct {
	// NodeTypes: list of node types compatible.
	NodeTypes []*NodeType `json:"node_types"`

	// TotalCount: number of node types in result set.
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

	DeploymentID string `json:"-"`

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

// ListVersionsRequest: Retrieve a list of available versions.
type ListVersionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: sort order for versions in the response.
	// Default value: version_asc
	OrderBy ListVersionsRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of versions to return per page.
	PageSize *uint32 `json:"-"`

	// Version: engine version to filter for, only versions with this version will be returned.
	Version *string `json:"-"`
}

// ListVersionsResponse: list versions response.
type ListVersionsResponse struct {
	// Versions: list of versions.
	Versions []*Version `json:"versions"`

	// TotalCount: number of versions in result set.
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

// UpdateDeploymentRequest: update deployment request.
type UpdateDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment.
	DeploymentID string `json:"-"`

	// Name: new name for the deployment.
	Name *string `json:"name,omitempty"`

	// Tags: tags to update.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateUserRequest: Update a deployment user.
type UpdateUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment.
	DeploymentID string `json:"-"`

	// Username: username of the deployment user.
	Username string `json:"-"`

	// Password: password of the deployment user.
	Password *string `json:"password,omitempty"`
}

// UpgradeDeploymentRequest: upgrade deployment request.
type UpgradeDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment.
	DeploymentID string `json:"-"`

	// NodeCount: target number of nodes.
	// Precisely one of NodeCount, VolumeSizeBytes must be set.
	NodeCount *uint32 `json:"node_count,omitempty"`

	// VolumeSizeBytes: target volume size.
	// Precisely one of NodeCount, VolumeSizeBytes must be set.
	VolumeSizeBytes *scw.Size `json:"volume_size_bytes,omitempty"`
}

// The MessageQ API allows you to manage your MessageQ deployments.
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

// CreateDeployment: Create a new MessageQ deployment.
func (s *API) CreateDeployment(req *CreateDeploymentRequest, opts ...scw.RequestOption) (*Deployment, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Deployment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDeployment: Update a MessageQ deployment.
func (s *API) UpdateDeployment(req *UpdateDeploymentRequest, opts ...scw.RequestOption) (*Deployment, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeploymentID) == "" {
		return nil, errors.New("field DeploymentID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Deployment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpgradeDeployment: Upgrade a MessageQ deployment.
func (s *API) UpgradeDeployment(req *UpgradeDeploymentRequest, opts ...scw.RequestOption) (*Deployment, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeploymentID) == "" {
		return nil, errors.New("field DeploymentID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/upgrade",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Deployment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDeployment: Retrieve a specific MessageQ deployment.
func (s *API) GetDeployment(req *GetDeploymentRequest, opts ...scw.RequestOption) (*Deployment, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeploymentID) == "" {
		return nil, errors.New("field DeploymentID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
	}

	var resp Deployment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForDeploymentRequest is used by WaitForDeployment method.
type WaitForDeploymentRequest struct {
	Region        scw.Region
	DeploymentID  string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForDeployment waits for the Deployment to reach a terminal state.
func (s *API) WaitForDeployment(req *WaitForDeploymentRequest, opts ...scw.RequestOption) (*Deployment, error) {
	timeout := defaultMessageqTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultMessageqRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[DeploymentStatus]struct{}{
		DeploymentStatusCreating:     {},
		DeploymentStatusInitializing: {},
		DeploymentStatusUpgrading:    {},
		DeploymentStatusDeleting:     {},
		DeploymentStatusLocking:      {},
		DeploymentStatusUnlocking:    {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetDeployment(&GetDeploymentRequest{
				Region:       req.Region,
				DeploymentID: req.DeploymentID,
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
		return nil, errors.Wrap(err, "waiting for Deployment failed")
	}

	return res.(*Deployment), nil
}

// DeleteDeployment: Delete a MessageQ deployment.
func (s *API) DeleteDeployment(req *DeleteDeploymentRequest, opts ...scw.RequestOption) (*Deployment, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeploymentID) == "" {
		return nil, errors.New("field DeploymentID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
	}

	var resp Deployment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDeployments: Retrieve a list of MessageQ deployments.
func (s *API) ListDeployments(req *ListDeploymentsRequest, opts ...scw.RequestOption) (*ListDeploymentsResponse, error) {
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
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "version", req.Version)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments",
		Query:  query,
	}

	var resp ListDeploymentsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVersions: List available MessageQ versions.
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "version", req.Version)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/versions",
		Query:  query,
	}

	var resp ListVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodeTypes: Retrieve a list of available node types.
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/node-types",
		Query:  query,
	}

	var resp ListNodeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEndpoint: Create a new endpoint for a deployment.
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
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/endpoints",
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

// DeleteEndpoint: Delete an existing endpoint.
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
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
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

	if fmt.Sprint(req.DeploymentID) == "" {
		return nil, errors.New("field DeploymentID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/users",
		Query:  query,
	}

	var resp ListUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateUser: Create a new user.
func (s *API) CreateUser(req *CreateUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeploymentID) == "" {
		return nil, errors.New("field DeploymentID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/users",
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

	if fmt.Sprint(req.DeploymentID) == "" {
		return nil, errors.New("field DeploymentID cannot be empty in request")
	}

	if fmt.Sprint(req.Username) == "" {
		return nil, errors.New("field Username cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/users/" + fmt.Sprint(req.Username) + "",
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

// DeleteUser: Delete an existing user.
func (s *API) DeleteUser(req *DeleteUserRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeploymentID) == "" {
		return errors.New("field DeploymentID cannot be empty in request")
	}

	if fmt.Sprint(req.Username) == "" {
		return errors.New("field Username cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/messageq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/users/" + fmt.Sprint(req.Username) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

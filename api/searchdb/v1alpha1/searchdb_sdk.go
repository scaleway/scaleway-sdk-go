// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package searchdb provides methods and message types of the searchdb v1alpha1 API.
package searchdb

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
	defaultSearchdbRetryInterval = 15 * time.Second
	defaultSearchdbTimeout       = 5 * time.Minute
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

// Endpoint: Refers to an Endpoint.
type Endpoint struct {
	// ID: unique identifier of the Endpoint.
	ID string `json:"id"`

	// Deprecated: DNSRecord: DNS entry to access to the service. Now deprecated. Use the `url` field from `services` field instead.
	DNSRecord *string `json:"dns_record,omitempty"`

	// Services: list of available services, their ports and URLs.
	Services []*EndpointService `json:"services"`

	// Precisely one of Public, PrivateNetwork must be set.
	Public *EndpointPublicDetails `json:"public,omitempty"`

	// Precisely one of Public, PrivateNetwork must be set.
	PrivateNetwork *EndpointPrivateNetworkDetails `json:"private_network,omitempty"`
}

// Volume: Volume.
type Volume struct {
	// Type: define the type of the Volume.
	// Default value: unknown_type
	Type VolumeType `json:"type"`

	// SizeBytes: define the size of the Volume.
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

// Deployment: Refers to a Deployment.
type Deployment struct {
	// ID: unique identifier of the Deployment.
	ID string `json:"id"`

	// Name: name of the Deployment.
	Name string `json:"name"`

	// OrganizationID: ID of the Organization containing the Deployment.
	OrganizationID string `json:"organization_id"`

	// ProjectID: ID of the Project containing the Deployment.
	ProjectID string `json:"project_id"`

	// Status: status of the Deployment.
	// Default value: unknown_status
	Status DeploymentStatus `json:"status"`

	// Tags: tags of the Deployment.
	Tags []string `json:"tags"`

	// NodeAmount: number of nodes allocated per deployment.
	NodeAmount uint32 `json:"node_amount"`

	// NodeType: node type used in deployment.
	NodeType string `json:"node_type"`

	// Volume: volume type and size.
	Volume *Volume `json:"volume"`

	// Endpoints: exposed endpoints.
	Endpoints []*Endpoint `json:"endpoints"`

	// CreatedAt: creation date of the Deployment.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date when last update was done to the Deployment.
	UpdatedAt *time.Time `json:"updated_at"`

	// Version: opensearch version of the Deployment.
	Version string `json:"version"`

	// Region: region the Deployment is located.
	Region scw.Region `json:"region"`
}

// NodeType: Node type.
type NodeType struct {
	// StockStatus: stock status of the node type.
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

	// InstanceRange: instance range associated with the NodeType offer.
	InstanceRange string `json:"instance_range"`

	// AvailableVolumeTypes: available storage options for the Node Type.
	AvailableVolumeTypes []*NodeTypeVolumeType `json:"available_volume_types"`
}

// User: user.
type User struct {
	Username string `json:"username"`
}

// Version: Opensearch Version.
type Version struct {
	// Version: opensearch Version.
	Version string `json:"version"`

	// EndOfLife: end of life date of the version.
	EndOfLife *time.Time `json:"end_of_life"`

	// Disabled: parameter that tell if the version is disabled.
	Disabled bool `json:"disabled"`

	// Beta: parameter that tell if the version is in beta.
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

	// NodeAmount: number of nodes.
	NodeAmount uint32 `json:"node_amount"`

	// NodeType: node type.
	NodeType string `json:"node_type"`

	// UserName: username for the deployment user.
	UserName *string `json:"user_name,omitempty"`

	// Password: password for the deployment user.
	Password *string `json:"password,omitempty"`

	// Volume: volume.
	Volume *Volume `json:"volume,omitempty"`

	// Endpoints: endpoints to access the deployment.
	Endpoints []*EndpointSpec `json:"endpoints"`

	// Version: the Opensearch version to use.
	Version string `json:"version"`
}

// CreateEndpointRequest: Create an endpoint for a specific deployment.
type CreateEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment for which to create an endpoint.
	DeploymentID string `json:"deployment_id"`

	// EndpointSpec: specification of the endpoint you want to create.
	EndpointSpec *EndpointSpec `json:"endpoint_spec,omitempty"`
}

// CreateUserRequest: Create a user in an deployment.
type CreateUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment in which to create the user.
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

	// EndpointID: ID of the endpoint to delete.
	EndpointID string `json:"-"`
}

// DeleteUserRequest: Delete a user from a deployment.
type DeleteUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment in which to create the user.
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

	// OrganizationID: ID of the Organization containing the deployments.
	OrganizationID *string `json:"-"`

	// ProjectID: ID of the Project containing the deployments.
	ProjectID *string `json:"-"`

	// OrderBy: define the order of the returned deployments.
	// Default value: created_at_asc
	OrderBy ListDeploymentsRequestOrderBy `json:"-"`

	// Page: the page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of deployments to return.
	PageSize *uint32 `json:"-"`

	// Tags: filter by tag, only deployments with one or more matching tags will be returned.
	Tags []string `json:"-"`

	// Name: deployment name to filter for.
	Name *string `json:"-"`

	// Version: engine version to filter for.
	Version *string `json:"-"`
}

// ListDeploymentsResponse: Retrieve a list of deployments.
type ListDeploymentsResponse struct {
	// Deployments: list of deployments available.
	Deployments []*Deployment `json:"deployments"`

	// TotalCount: total number of objects.
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

// ListNodeTypesRequest: Retrieve a list of available node types for a Cloud Essentials for OpenSearch cluster.
type ListNodeTypesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: sort order of nodes in the response (name, vcpus or memory).
	// Default value: name_asc
	OrderBy ListNodeTypesRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of node types to return.
	PageSize *uint32 `json:"-"`
}

// ListNodeTypesResponse: Returns a list of node types available for a Cloud Essentials for OpenSearch cluster.
type ListNodeTypesResponse struct {
	// NodeTypes: node types compatible with the cluster.
	NodeTypes []*NodeType `json:"node_types"`

	// TotalCount: number of available node types to return.
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

	// OrderBy: define the order of the returned version.
	// Default value: version_asc
	OrderBy ListVersionsRequestOrderBy `json:"-"`

	// Page: the page number to return, form the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of version to return.
	PageSize *uint32 `json:"-"`

	// Version: filter by version.
	Version *string `json:"-"`
}

// ListVersionsResponse: Retrieve a list of version.
type ListVersionsResponse struct {
	// Versions: list of versions.
	Versions []*Version `json:"versions"`

	// TotalCount: number of versions in the list.
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

	// DeploymentID: UUID of the deployment to update.
	DeploymentID string `json:"-"`

	// Name: name of the deployment.
	Name *string `json:"name,omitempty"`

	// Tags: tags of a deployment.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateUserRequest: Update a user in an deployment.
type UpdateUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment in which to create the user.
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

	// DeploymentID: UUID of the Deployment to upgrade.
	DeploymentID string `json:"-"`

	// NodeAmount: amount of node upgrade target.
	// Precisely one of NodeAmount, VolumeSizeBytes must be set.
	NodeAmount *uint32 `json:"node_amount,omitempty"`

	// VolumeSizeBytes: volume size upgrade target.
	// Precisely one of NodeAmount, VolumeSizeBytes must be set.
	VolumeSizeBytes *uint64 `json:"volume_size_bytes,omitempty"`
}

// The Cloud Essentials for Opensearch API allows you to manage your Opensearch resources.
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

// CreateDeployment: Create a new Cloud Essentials for OpenSearch deployment.
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments",
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

// UpdateDeployment: Update a Cloud Essentials for OpenSearch deployment.
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
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

// UpgradeDeployment: Upgrade a Cloud Essentials for OpenSearch deployment.
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/upgrade",
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

// GetDeployment: Retrieve a specific Cloud Essentials for OpenSearch deployment.
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
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
	timeout := defaultSearchdbTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultSearchdbRetryInterval
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

// DeleteDeployment: Delete a Cloud Essentials for OpenSearch deployment.
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
	}

	var resp Deployment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDeployments: Retrieve a list of Cloud Essentials for OpenSearch deployments.
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments",
		Query:  query,
	}

	var resp ListDeploymentsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVersions: List available Cloud Essentials for OpenSearch versions.
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/versions",
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/node-types",
		Query:  query,
	}

	var resp ListNodeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEndpoint: Create a new endpoint on a deployment.
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/endpoints",
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/users",
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/users",
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/users/" + fmt.Sprint(req.Username) + "",
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
		Path:   "/searchdb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/users/" + fmt.Sprint(req.Username) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

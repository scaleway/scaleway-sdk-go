// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package datawarehouse provides methods and message types of the datawarehouse v1beta1 API.
package datawarehouse

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
	defaultDatawarehouseRetryInterval = 15 * time.Second
	defaultDatawarehouseTimeout       = 5 * time.Minute
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
	DeploymentStatusConfiguring   = DeploymentStatus("configuring")
	DeploymentStatusDeleting      = DeploymentStatus("deleting")
	DeploymentStatusError         = DeploymentStatus("error")
	DeploymentStatusLocked        = DeploymentStatus("locked")
	DeploymentStatusLocking       = DeploymentStatus("locking")
	DeploymentStatusUnlocking     = DeploymentStatus("unlocking")
	DeploymentStatusDeploying     = DeploymentStatus("deploying")
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
		"configuring",
		"deleting",
		"error",
		"locked",
		"locking",
		"unlocking",
		"deploying",
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

type EndpointServiceProtocol string

const (
	EndpointServiceProtocolUnknownProtocol = EndpointServiceProtocol("unknown_protocol")
	EndpointServiceProtocolTCP             = EndpointServiceProtocol("tcp")
	EndpointServiceProtocolHTTPS           = EndpointServiceProtocol("https")
	EndpointServiceProtocolMysql           = EndpointServiceProtocol("mysql")
)

func (enum EndpointServiceProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return string(EndpointServiceProtocolUnknownProtocol)
	}
	return string(enum)
}

func (enum EndpointServiceProtocol) Values() []EndpointServiceProtocol {
	return []EndpointServiceProtocol{
		"unknown_protocol",
		"tcp",
		"https",
		"mysql",
	}
}

func (enum EndpointServiceProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EndpointServiceProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EndpointServiceProtocol(EndpointServiceProtocol(tmp).String())
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
		return string(ListDatabasesRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListDatabasesRequestOrderBy) Values() []ListDatabasesRequestOrderBy {
	return []ListDatabasesRequestOrderBy{
		"name_asc",
		"name_desc",
		"size_asc",
		"size_desc",
	}
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

type ListDeploymentsRequestOrderBy string

const (
	ListDeploymentsRequestOrderByCreatedAtDesc = ListDeploymentsRequestOrderBy("created_at_desc")
	ListDeploymentsRequestOrderByCreatedAtAsc  = ListDeploymentsRequestOrderBy("created_at_asc")
	ListDeploymentsRequestOrderByNameAsc       = ListDeploymentsRequestOrderBy("name_asc")
	ListDeploymentsRequestOrderByNameDesc      = ListDeploymentsRequestOrderBy("name_desc")
)

func (enum ListDeploymentsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListDeploymentsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListDeploymentsRequestOrderBy) Values() []ListDeploymentsRequestOrderBy {
	return []ListDeploymentsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"name_asc",
		"name_desc",
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

// EndpointPrivateNetworkDetails: endpoint private network details.
type EndpointPrivateNetworkDetails struct {
	PrivateNetworkID string `json:"private_network_id"`
}

// EndpointPublicDetails: endpoint public details.
type EndpointPublicDetails struct{}

// EndpointService: endpoint service.
type EndpointService struct {
	// Protocol: default value: unknown_protocol
	Protocol EndpointServiceProtocol `json:"protocol"`

	Port uint32 `json:"port"`
}

// EndpointSpecPrivateNetworkDetails: endpoint spec private network details.
type EndpointSpecPrivateNetworkDetails struct {
	// PrivateNetworkID: UUID of the Private Network.
	PrivateNetworkID string `json:"private_network_id"`
}

// EndpointSpecPublicDetails: endpoint spec public details.
type EndpointSpecPublicDetails struct{}

// Endpoint: endpoint.
type Endpoint struct {
	// ID: unique identifier of the endpoint.
	ID string `json:"id"`

	// DNSRecord: DNS record associated with the endpoint.
	DNSRecord string `json:"dns_record"`

	// Services: list of services associated with the endpoint.
	Services []*EndpointService `json:"services"`

	// PrivateNetwork: private Network endpoint details.
	// Precisely one of PrivateNetwork, Public must be set.
	PrivateNetwork *EndpointPrivateNetworkDetails `json:"private_network,omitempty"`

	// Public: public endpoint details.
	// Precisely one of PrivateNetwork, Public must be set.
	Public *EndpointPublicDetails `json:"public,omitempty"`
}

// EndpointSpec: endpoint spec.
type EndpointSpec struct {
	// Precisely one of Public, PrivateNetwork must be set.
	Public *EndpointSpecPublicDetails `json:"public,omitempty"`

	// Precisely one of Public, PrivateNetwork must be set.
	PrivateNetwork *EndpointSpecPrivateNetworkDetails `json:"private_network,omitempty"`
}

// Database: database.
type Database struct {
	// Name: name of the database.
	Name string `json:"name"`

	// Size: size of the database.
	Size scw.Size `json:"size"`
}

// Deployment: deployment.
type Deployment struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// Name: name of the deployment.
	Name string `json:"name"`

	// OrganizationID: organization ID.
	OrganizationID string `json:"organization_id"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// Status: status of the deployment.
	// Default value: unknown_status
	Status DeploymentStatus `json:"status"`

	// Tags: list of tags applied to the deployment.
	Tags []string `json:"tags"`

	// CreatedAt: creation date of the deployment.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the deployment.
	UpdatedAt *time.Time `json:"updated_at"`

	// Version: clickHouse® version.
	Version string `json:"version"`

	// ReplicaCount: number of replicas for the deployment.
	ReplicaCount uint32 `json:"replica_count"`

	// CPUMin: minimum CPU count for the deployment.
	CPUMin uint32 `json:"cpu_min"`

	// CPUMax: maximum CPU count for the deployment.
	CPUMax uint32 `json:"cpu_max"`

	// Endpoints: list of endpoints associated with the deployment.
	Endpoints []*Endpoint `json:"endpoints"`

	// RAMPerCPU: RAM per CPU count for the deployment (in GB).
	RAMPerCPU uint32 `json:"ram_per_cpu"`

	// Region: region of the deployment.
	Region scw.Region `json:"region"`
}

// Preset: preset.
type Preset struct {
	// Name: name of the preset.
	Name string `json:"name"`

	// Category: category of the preset.
	Category string `json:"category"`

	// CPUMin: minimum CPU count for the preset.
	CPUMin uint32 `json:"cpu_min"`

	// CPUMax: maximum CPU count for the preset.
	CPUMax uint32 `json:"cpu_max"`

	// RAMPerCPU: RAM per CPU count for the preset (in GB).
	RAMPerCPU uint32 `json:"ram_per_cpu"`

	// ReplicaCount: number of replicas for the preset.
	ReplicaCount uint32 `json:"replica_count"`
}

// User: user.
type User struct {
	// Name: name of the user.
	Name string `json:"name"`

	// IsAdmin: indicates if the user is an administrator.
	IsAdmin bool `json:"is_admin"`
}

// Version: version.
type Version struct {
	// Version: deployment version.
	Version string `json:"version"`

	// EndOfLifeAt: date of End of Life.
	EndOfLifeAt *time.Time `json:"end_of_life_at"`
}

// CreateDatabaseRequest: create database request.
type CreateDatabaseRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: UUID of the deployment.
	DeploymentID string `json:"-"`

	// Name: name of the database.
	Name string `json:"name"`
}

// CreateDeploymentRequest: create deployment request.
type CreateDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: the Project ID on which the deployment will be created.
	ProjectID string `json:"project_id"`

	// Name: name of the deployment.
	Name string `json:"name"`

	// Tags: tags to apply to the deployment.
	Tags []string `json:"tags"`

	// Version: clickHouse® version to use for the deployment.
	Version string `json:"version"`

	// ReplicaCount: number of replicas for the deployment.
	ReplicaCount uint32 `json:"replica_count"`

	// Password: password for the initial user.
	Password string `json:"password"`

	// CPUMin: minimum CPU count for the deployment.
	CPUMin uint32 `json:"cpu_min"`

	// CPUMax: maximum CPU count for the deployment.
	CPUMax uint32 `json:"cpu_max"`

	// Endpoints: endpoints to associate with the deployment.
	Endpoints []*EndpointSpec `json:"endpoints"`

	// RAMPerCPU: RAM per CPU count for the deployment (in GB).
	RAMPerCPU uint32 `json:"ram_per_cpu"`
}

// CreateEndpointRequest: create endpoint request.
type CreateEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: UUID of the deployment.
	DeploymentID string `json:"deployment_id"`

	// Endpoint: endpoint specification.
	Endpoint *EndpointSpec `json:"endpoint,omitempty"`
}

// CreateUserRequest: create user request.
type CreateUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: UUID of the deployment.
	DeploymentID string `json:"-"`

	// Name: name of the user.
	Name string `json:"name"`

	// Password: password for the user.
	Password string `json:"password"`

	// IsAdmin: indicates if the user is an administrator.
	IsAdmin bool `json:"is_admin"`
}

// DeleteDatabaseRequest: delete database request.
type DeleteDatabaseRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: UUID of the deployment.
	DeploymentID string `json:"-"`

	// Name: name of the database to delete.
	Name string `json:"-"`
}

// DeleteDeploymentRequest: delete deployment request.
type DeleteDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: UUID of the deployment to delete.
	DeploymentID string `json:"-"`
}

// DeleteEndpointRequest: delete endpoint request.
type DeleteEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// EndpointID: UUID of the Endpoint to delete.
	EndpointID string `json:"-"`
}

// DeleteUserRequest: delete user request.
type DeleteUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: UUID of the deployment.
	DeploymentID string `json:"-"`

	// Name: name of the user to delete.
	Name string `json:"-"`
}

// GetDeploymentCertificateRequest: get deployment certificate request.
type GetDeploymentCertificateRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DeploymentID string `json:"-"`
}

// GetDeploymentRequest: get deployment request.
type GetDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: UUID of the deployment.
	DeploymentID string `json:"-"`
}

// ListDatabasesRequest: list databases request.
type ListDatabasesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: UUID of the deployment.
	DeploymentID string `json:"-"`

	// Name: name of the database to filter by.
	Name *string `json:"-"`

	// OrderBy: criteria to use when ordering database listings.
	// Default value: name_asc
	OrderBy ListDatabasesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListDatabasesResponse: list databases response.
type ListDatabasesResponse struct {
	// Databases: list of databases associated with the deployment.
	Databases []*Database `json:"databases"`

	// TotalCount: total count of databases associated with the deployment.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatabasesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatabasesResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListDatabasesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Databases = append(r.Databases, results.Databases...)
	r.TotalCount += uint32(len(results.Databases))
	return uint32(len(results.Databases)), nil
}

// ListDeploymentsRequest: list deployments request.
type ListDeploymentsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Tags: list deployments with a given tag.
	Tags []string `json:"-"`

	// Name: lists deployments that match a name pattern.
	Name *string `json:"-"`

	// OrderBy: criteria to use when ordering deployment listings.
	// Default value: created_at_desc
	OrderBy ListDeploymentsRequestOrderBy `json:"-"`

	// OrganizationID: organization ID the deployment belongs to.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID the deployment belongs to.
	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListDeploymentsResponse: list deployments response.
type ListDeploymentsResponse struct {
	// Deployments: list of all deployments available in an Organization or Project.
	Deployments []*Deployment `json:"deployments"`

	// TotalCount: total count of deployments available in an Organization or Project.
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

// ListPresetsRequest: list presets request.
type ListPresetsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListPresetsResponse: list presets response.
type ListPresetsResponse struct {
	// Presets: list of available presets.
	Presets []*Preset `json:"presets"`

	// TotalCount: total count of presets available.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPresetsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPresetsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListPresetsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Presets = append(r.Presets, results.Presets...)
	r.TotalCount += uint64(len(results.Presets))
	return uint64(len(results.Presets)), nil
}

// ListUsersRequest: list users request.
type ListUsersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: UUID of the deployment.
	DeploymentID string `json:"-"`

	// Name: name of the user to filter by.
	Name *string `json:"-"`

	// OrderBy: criteria to use when ordering user listings.
	// Default value: name_asc
	OrderBy ListUsersRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListUsersResponse: list users response.
type ListUsersResponse struct {
	// Users: list of users associated with the deployment.
	Users []*User `json:"users"`

	// TotalCount: total count of users associated with the deployment.
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

	Version *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListVersionsResponse: list versions response.
type ListVersionsResponse struct {
	// Versions: available deployment version.
	Versions []*Version `json:"versions"`

	// TotalCount: total count of deployment version available.
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

	// CPUMin: minimum CPU count for the deployment.
	CPUMin *uint32 `json:"cpu_min,omitempty"`

	// CPUMax: maximum CPU count for the deployment.
	CPUMax *uint32 `json:"cpu_max,omitempty"`

	// ReplicaCount: number of replicas for the deployment.
	ReplicaCount *uint32 `json:"replica_count,omitempty"`
}

// UpdateUserRequest: update user request.
type UpdateUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: UUID of the deployment.
	DeploymentID string `json:"-"`

	// Name: name of the user.
	Name string `json:"-"`

	// Password: new password for the user.
	Password *string `json:"password,omitempty"`

	// IsAdmin: updates the user administrator permissions.
	IsAdmin *bool `json:"is_admin,omitempty"`
}

// This API allows you to manage your Data Warehouse.
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

// ListPresets: List available presets.
func (s *API) ListPresets(req *ListPresetsRequest, opts ...scw.RequestOption) (*ListPresetsResponse, error) {
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
		Method: "GET",
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/presets",
		Query:  query,
	}

	var resp ListPresetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVersions: List available ClickHouse® versions.
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
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/versions",
		Query:  query,
	}

	var resp ListVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDeployments: List all deployments in the specified region, for a given Scaleway Project. By default, the deployments returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. You can define additional parameters for your query, such as `tags` and `name`. For the `name` parameter, the value you provide will be checked against the whole name string to see if it includes the string you put in the parameter.
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
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments",
		Query:  query,
	}

	var resp ListDeploymentsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDeployment: Retrieve information about a given deployment, specified by the `region` and `deployment_id` parameters. Its full details, including name, status are returned in the response object.
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
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
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
	timeout := defaultDatawarehouseTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultDatawarehouseRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[DeploymentStatus]struct{}{
		DeploymentStatusCreating:    {},
		DeploymentStatusConfiguring: {},
		DeploymentStatusDeleting:    {},
		DeploymentStatusLocking:     {},
		DeploymentStatusUnlocking:   {},
		DeploymentStatusDeploying:   {},
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

// CreateDeployment: Create a new deployment.
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
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments",
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

// UpdateDeployment: Update the parameters of a deployment.
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
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
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

// DeleteDeployment: Delete a given deployment, specified by the `region` and `deployment_id` parameters. Deleting a deployment is permanent, and cannot be undone. Upon deletion, all your data will be lost.
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
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
	}

	var resp Deployment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDeploymentCertificate:
func (s *API) GetDeploymentCertificate(req *GetDeploymentCertificateRequest, opts ...scw.RequestOption) (*scw.File, error) {
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
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/certificate",
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListUsers: List users associated with a deployment.
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

	if fmt.Sprint(req.DeploymentID) == "" {
		return nil, errors.New("field DeploymentID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/users",
		Query:  query,
	}

	var resp ListUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateUser: Create a new user for a deployment.
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
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/users",
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

// UpdateUser: Update an existing user for a deployment.
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

	if fmt.Sprint(req.Name) == "" {
		return nil, errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/users/" + fmt.Sprint(req.Name) + "",
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

// DeleteUser: Delete a user from a deployment.
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

	if fmt.Sprint(req.Name) == "" {
		return errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/users/" + fmt.Sprint(req.Name) + "",
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

// DeleteEndpoint: Delete an endpoint from a deployment.
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
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
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
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/endpoints",
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

// ListDatabases: List databases within a deployment.
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeploymentID) == "" {
		return nil, errors.New("field DeploymentID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/databases",
		Query:  query,
	}

	var resp ListDatabasesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDatabase: Create a new database within a deployment.
func (s *API) CreateDatabase(req *CreateDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
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
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/databases",
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

// DeleteDatabase: Delete a database from a deployment.
func (s *API) DeleteDatabase(req *DeleteDatabaseRequest, opts ...scw.RequestOption) error {
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

	if fmt.Sprint(req.Name) == "" {
		return errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/datawarehouse/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/databases/" + fmt.Sprint(req.Name) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package inference provides methods and message types of the inference v1beta1 API.
package inference

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

type DeploymentStatus string

const (
	DeploymentStatusUnknownStatus = DeploymentStatus("unknown_status")
	DeploymentStatusCreating      = DeploymentStatus("creating")
	DeploymentStatusDeploying     = DeploymentStatus("deploying")
	DeploymentStatusReady         = DeploymentStatus("ready")
	DeploymentStatusError         = DeploymentStatus("error")
	DeploymentStatusDeleting      = DeploymentStatus("deleting")
	DeploymentStatusLocked        = DeploymentStatus("locked")
)

func (enum DeploymentStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum DeploymentStatus) Values() []DeploymentStatus {
	return []DeploymentStatus{
		"unknown_status",
		"creating",
		"deploying",
		"ready",
		"error",
		"deleting",
		"locked",
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
	ListDeploymentsRequestOrderByCreatedAtDesc = ListDeploymentsRequestOrderBy("created_at_desc")
	ListDeploymentsRequestOrderByCreatedAtAsc  = ListDeploymentsRequestOrderBy("created_at_asc")
	ListDeploymentsRequestOrderByNameAsc       = ListDeploymentsRequestOrderBy("name_asc")
	ListDeploymentsRequestOrderByNameDesc      = ListDeploymentsRequestOrderBy("name_desc")
)

func (enum ListDeploymentsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_desc"
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

type ListModelsRequestOrderBy string

const (
	ListModelsRequestOrderByDisplayRankAsc = ListModelsRequestOrderBy("display_rank_asc")
	ListModelsRequestOrderByCreatedAtAsc   = ListModelsRequestOrderBy("created_at_asc")
	ListModelsRequestOrderByCreatedAtDesc  = ListModelsRequestOrderBy("created_at_desc")
	ListModelsRequestOrderByNameAsc        = ListModelsRequestOrderBy("name_asc")
	ListModelsRequestOrderByNameDesc       = ListModelsRequestOrderBy("name_desc")
)

func (enum ListModelsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "display_rank_asc"
	}
	return string(enum)
}

func (enum ListModelsRequestOrderBy) Values() []ListModelsRequestOrderBy {
	return []ListModelsRequestOrderBy{
		"display_rank_asc",
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListModelsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListModelsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListModelsRequestOrderBy(ListModelsRequestOrderBy(tmp).String())
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

// EndpointPrivateNetworkDetails: endpoint private network details.
type EndpointPrivateNetworkDetails struct {
	// PrivateNetworkID: ID of the Private Network.
	PrivateNetworkID string `json:"private_network_id"`
}

// EndpointPublicAccessDetails: endpoint public access details.
type EndpointPublicAccessDetails struct {
}

// EndpointSpecPrivateNetwork: endpoint spec private network.
type EndpointSpecPrivateNetwork struct {
	// PrivateNetworkID: ID of the Private Network.
	PrivateNetworkID string `json:"private_network_id"`
}

// EndpointSpecPublic: endpoint spec public.
type EndpointSpecPublic struct {
}

// Endpoint: endpoint.
type Endpoint struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// URL: for private endpoints, the URL will be accessible only from the Private Network.
	// In addition, private endpoints will expose a CA certificate that can be used to verify the server's identity.
	// This CA certificate can be retrieved using the `GetDeploymentCertificate` API call.
	URL string `json:"url"`

	// PublicAccess: defines whether the endpoint is public.
	// Precisely one of PublicAccess, PrivateNetwork must be set.
	PublicAccess *EndpointPublicAccessDetails `json:"public_access,omitempty"`

	// PrivateNetwork: details of the Private Network.
	// Precisely one of PublicAccess, PrivateNetwork must be set.
	PrivateNetwork *EndpointPrivateNetworkDetails `json:"private_network,omitempty"`

	// DisableAuth: defines whether the authentication is disabled.
	DisableAuth bool `json:"disable_auth"`
}

// ModelS3Model: model s3 model.
type ModelS3Model struct {
	S3URL string `json:"s3_url"`

	PythonDependencies map[string]string `json:"python_dependencies"`

	NodeType *string `json:"node_type"`

	TritonServerVersion *string `json:"triton_server_version"`
}

// ACLRuleRequest: acl rule request.
type ACLRuleRequest struct {
	// IP: it can be specified as a single IP address or a range of IP addresses in CIDR notation.
	IP scw.IPNet `json:"ip"`

	// Description: description of the ACL rule.
	Description string `json:"description"`
}

// ACLRule: acl rule.
type ACLRule struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// IP: allowed IP address or CIDR range.
	IP scw.IPNet `json:"ip"`

	// Description: description of the ACL rule.
	Description string `json:"description"`
}

// EndpointSpec: endpoint spec.
type EndpointSpec struct {
	// Public: set the endpoint as public.
	// Precisely one of Public, PrivateNetwork must be set.
	Public *EndpointSpecPublic `json:"public,omitempty"`

	// PrivateNetwork: private endpoints are only accessible from the Private Network.
	// Precisely one of Public, PrivateNetwork must be set.
	PrivateNetwork *EndpointSpecPrivateNetwork `json:"private_network,omitempty"`

	// DisableAuth: by default, deployments are protected by IAM authentication.
	// When setting this field to true, the authentication will be disabled.
	DisableAuth bool `json:"disable_auth"`
}

// Deployment: deployment.
type Deployment struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// Name: name of the deployment.
	Name string `json:"name"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// Status: status of the deployment.
	// Default value: unknown_status
	Status DeploymentStatus `json:"status"`

	// Tags: list of tags applied to the deployment.
	Tags []string `json:"tags"`

	// NodeType: node type of the deployment.
	NodeType string `json:"node_type"`

	// Endpoints: list of endpoints.
	Endpoints []*Endpoint `json:"endpoints"`

	// Size: current size of the pool.
	Size uint32 `json:"size"`

	// MinSize: defines the minimum size of the pool.
	MinSize uint32 `json:"min_size"`

	// MaxSize: defines the maximum size of the pool.
	MaxSize uint32 `json:"max_size"`

	// ErrorMessage: displays information if your deployment is in error state.
	ErrorMessage *string `json:"error_message"`

	// ModelName: the inference model used for the deployment.
	ModelName string `json:"model_name"`

	// ModelID: ID of the model used for the deployment.
	ModelID string `json:"model_id"`

	// CreatedAt: creation date of the deployment.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the deployment.
	UpdatedAt *time.Time `json:"updated_at"`

	// Region: region of the deployment.
	Region scw.Region `json:"region"`
}

// Model: model.
type Model struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// Name: unique Name identifier.
	Name string `json:"name"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// Provider: name of the model provider.
	Provider string `json:"provider"`

	// Tags: list of tags applied to the model.
	Tags []string `json:"tags"`

	// Description: purpose of the model.
	Description string `json:"description"`

	// HasEula: defines whether the model has an end user license agreement.
	HasEula bool `json:"has_eula"`

	// CreatedAt: creation date of the model.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the model.
	UpdatedAt *time.Time `json:"updated_at"`

	// Region: region of the model.
	Region scw.Region `json:"region"`

	// S3Model: object Storage URL pointing to the model source weight.
	// Precisely one of S3Model must be set.
	S3Model *ModelS3Model `json:"s3_model,omitempty"`

	// IsPublic: defines whether the model is public or not.
	IsPublic bool `json:"is_public"`

	// CompatibleNodeTypes: names of the node types compatible with the model.
	CompatibleNodeTypes []string `json:"compatible_node_types"`

	// QuantizationLevel: quantization level of the model.
	QuantizationLevel string `json:"quantization_level"`
}

// NodeType: node type.
type NodeType struct {
	// Name: name of the node type.
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

	// Vram: quantity of GPU RAM.
	Vram scw.Size `json:"vram"`

	// Disabled: the node type is currently disabled.
	Disabled bool `json:"disabled"`

	// Beta: the node type is currently in beta.
	Beta bool `json:"beta"`

	// CreatedAt: creation date of the node type.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the node type.
	UpdatedAt *time.Time `json:"updated_at"`

	// Gpus: number of GPUs.
	Gpus uint32 `json:"gpus"`

	// Region: region of the node type.
	Region scw.Region `json:"region"`
}

// AddDeploymentACLRulesRequest: add deployment acl rules request.
type AddDeploymentACLRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment to add ACL rules to.
	DeploymentID string `json:"-"`

	// ACLs: list of ACL rules to add.
	ACLs []*ACLRuleRequest `json:"acls"`
}

// AddDeploymentACLRulesResponse: add deployment acl rules response.
type AddDeploymentACLRulesResponse struct {
	// Rules: list of ACL rules added.
	Rules []*ACLRule `json:"rules"`
}

// CreateDeploymentRequest: create deployment request.
type CreateDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: name of the deployment.
	Name string `json:"name"`

	// ProjectID: ID of the Project to create the deployment in.
	ProjectID string `json:"project_id"`

	// ModelName: name of the model to use.
	ModelName string `json:"model_name"`

	// AcceptEula: if the model has an EULA, you must accept it before proceeding.
	// The terms of the EULA can be retrieved using the `GetModelEula` API call.
	AcceptEula *bool `json:"accept_eula,omitempty"`

	// NodeType: name of the node type to use.
	NodeType string `json:"node_type"`

	// Tags: list of tags to apply to the deployment.
	Tags []string `json:"tags"`

	// MinSize: defines the minimum size of the pool.
	MinSize *uint32 `json:"min_size,omitempty"`

	// MaxSize: defines the maximum size of the pool.
	MaxSize *uint32 `json:"max_size,omitempty"`

	// Endpoints: list of endpoints to create.
	Endpoints []*EndpointSpec `json:"endpoints"`
}

// CreateEndpointRequest: create endpoint request.
type CreateEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment to create the endpoint for.
	DeploymentID string `json:"deployment_id"`

	// Endpoint: specification of the endpoint.
	Endpoint *EndpointSpec `json:"endpoint"`
}

// DeleteDeploymentACLRuleRequest: delete deployment acl rule request.
type DeleteDeploymentACLRuleRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ACLID: ID of the ACL rule to delete.
	ACLID string `json:"-"`
}

// DeleteDeploymentRequest: delete deployment request.
type DeleteDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment to delete.
	DeploymentID string `json:"-"`
}

// DeleteEndpointRequest: delete endpoint request.
type DeleteEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// EndpointID: ID of the endpoint to delete.
	EndpointID string `json:"-"`
}

// Eula: eula.
type Eula struct {
	// Content: content of the end user license agreement.
	Content string `json:"content"`
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

	// DeploymentID: ID of the deployment to get.
	DeploymentID string `json:"-"`
}

// GetModelEulaRequest: get model eula request.
type GetModelEulaRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ModelID string `json:"-"`
}

// GetModelRequest: get model request.
type GetModelRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ModelID: ID of the model to get.
	ModelID string `json:"-"`
}

// ListDeploymentACLRulesRequest: list deployment acl rules request.
type ListDeploymentACLRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment to list ACL rules for.
	DeploymentID string `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of ACL rules to return per page.
	PageSize *uint32 `json:"-"`
}

// ListDeploymentACLRulesResponse: list deployment acl rules response.
type ListDeploymentACLRulesResponse struct {
	// Rules: list of ACL rules on the current page.
	Rules []*ACLRule `json:"rules"`

	// TotalCount: total number of ACL rules.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDeploymentACLRulesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDeploymentACLRulesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListDeploymentACLRulesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Rules = append(r.Rules, results.Rules...)
	r.TotalCount += uint64(len(results.Rules))
	return uint64(len(results.Rules)), nil
}

// ListDeploymentsRequest: list deployments request.
type ListDeploymentsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of deployments to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_desc
	OrderBy ListDeploymentsRequestOrderBy `json:"-"`

	// ProjectID: filter by Project ID.
	ProjectID *string `json:"-"`

	// OrganizationID: filter by Organization ID.
	OrganizationID *string `json:"-"`

	// Name: filter by deployment name.
	Name *string `json:"-"`

	// Tags: filter by tags.
	Tags []string `json:"-"`
}

// ListDeploymentsResponse: list deployments response.
type ListDeploymentsResponse struct {
	// Deployments: list of deployments on the current page.
	Deployments []*Deployment `json:"deployments"`

	// TotalCount: total number of deployments.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDeploymentsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDeploymentsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListDeploymentsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Deployments = append(r.Deployments, results.Deployments...)
	r.TotalCount += uint64(len(results.Deployments))
	return uint64(len(results.Deployments)), nil
}

// ListModelsRequest: list models request.
type ListModelsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: display_rank_asc
	OrderBy ListModelsRequestOrderBy `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of models to return per page.
	PageSize *uint32 `json:"-"`

	// ProjectID: filter by Project ID.
	ProjectID *string `json:"-"`

	// Name: filter by model name.
	Name *string `json:"-"`

	// Tags: filter by tags.
	Tags []string `json:"-"`
}

// ListModelsResponse: list models response.
type ListModelsResponse struct {
	// Models: list of models on the current page.
	Models []*Model `json:"models"`

	// TotalCount: total number of models.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListModelsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListModelsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListModelsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Models = append(r.Models, results.Models...)
	r.TotalCount += uint64(len(results.Models))
	return uint64(len(results.Models)), nil
}

// ListNodeTypesRequest: list node types request.
type ListNodeTypesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of node types to return per page.
	PageSize *uint32 `json:"-"`

	// IncludeDisabledTypes: include disabled node types in the response.
	IncludeDisabledTypes bool `json:"-"`
}

// ListNodeTypesResponse: list node types response.
type ListNodeTypesResponse struct {
	// NodeTypes: list of node types.
	NodeTypes []*NodeType `json:"node_types"`

	// TotalCount: total number of node types.
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

// SetDeploymentACLRulesRequest: set deployment acl rules request.
type SetDeploymentACLRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment to set ACL rules for.
	DeploymentID string `json:"-"`

	// ACLs: all existing ACL rules will be replaced by the new ones.
	ACLs []*ACLRuleRequest `json:"acls"`
}

// SetDeploymentACLRulesResponse: set deployment acl rules response.
type SetDeploymentACLRulesResponse struct {
	// Rules: list of ACL rules that were set.
	Rules []*ACLRule `json:"rules"`
}

// UpdateDeploymentRequest: update deployment request.
type UpdateDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeploymentID: ID of the deployment to update.
	DeploymentID string `json:"-"`

	// Name: name of the deployment.
	Name *string `json:"name,omitempty"`

	// Tags: list of tags to apply to the deployment.
	Tags *[]string `json:"tags,omitempty"`

	// MinSize: defines the new minimum size of the pool.
	MinSize *uint32 `json:"min_size,omitempty"`

	// MaxSize: defines the new maximum size of the pool.
	MaxSize *uint32 `json:"max_size,omitempty"`
}

// UpdateEndpointRequest: update endpoint request.
type UpdateEndpointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// EndpointID: ID of the endpoint to update.
	EndpointID string `json:"-"`

	// DisableAuth: by default, deployments are protected by IAM authentication.
	// When setting this field to true, the authentication will be disabled.
	DisableAuth *bool `json:"disable_auth,omitempty"`
}

// This API allows you to manage your Inference services.
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

// ListDeployments: List all your inference deployments.
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments",
		Query:  query,
	}

	var resp ListDeploymentsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDeployment: Get the deployment for the given ID.
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
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
	}

	var resp Deployment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDeployment: Create a new inference deployment related to a specific model.
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

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("inference")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments",
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

// UpdateDeployment: Update an existing inference deployment.
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
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
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

// DeleteDeployment: Delete an existing inference deployment.
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
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
	}

	var resp Deployment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDeploymentCertificate: Get the CA certificate used for the deployment of private endpoints.
// The CA certificate will be returned as a PEM file.
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
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/certificate",
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEndpoint: Create a new Endpoint related to a specific deployment.
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
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/endpoints",
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

// UpdateEndpoint: Update an existing Endpoint.
func (s *API) UpdateEndpoint(req *UpdateEndpointRequest, opts ...scw.RequestOption) (*Endpoint, error) {
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
		Method: "PATCH",
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
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

// DeleteEndpoint: Delete an existing Endpoint.
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
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/endpoints/" + fmt.Sprint(req.EndpointID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListDeploymentACLRules: List ACLs for a specific deployment.
func (s *API) ListDeploymentACLRules(req *ListDeploymentACLRulesRequest, opts ...scw.RequestOption) (*ListDeploymentACLRulesResponse, error) {
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

	if fmt.Sprint(req.DeploymentID) == "" {
		return nil, errors.New("field DeploymentID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/acls",
		Query:  query,
	}

	var resp ListDeploymentACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddDeploymentACLRules: Add new ACL rules for a specific deployment.
func (s *API) AddDeploymentACLRules(req *AddDeploymentACLRulesRequest, opts ...scw.RequestOption) (*AddDeploymentACLRulesResponse, error) {
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
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddDeploymentACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetDeploymentACLRules: Set new ACL rules for a specific deployment.
func (s *API) SetDeploymentACLRules(req *SetDeploymentACLRulesRequest, opts ...scw.RequestOption) (*SetDeploymentACLRulesResponse, error) {
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
		Method: "PUT",
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/acls",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetDeploymentACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDeploymentACLRule: Delete an exising ACL.
func (s *API) DeleteDeploymentACLRule(req *DeleteDeploymentACLRuleRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/acls/" + fmt.Sprint(req.ACLID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListModels: List all available models.
func (s *API) ListModels(req *ListModelsRequest, opts ...scw.RequestOption) (*ListModelsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/models",
		Query:  query,
	}

	var resp ListModelsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetModel: Get the model for the given ID.
func (s *API) GetModel(req *GetModelRequest, opts ...scw.RequestOption) (*Model, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ModelID) == "" {
		return nil, errors.New("field ModelID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/models/" + fmt.Sprint(req.ModelID) + "",
	}

	var resp Model

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetModelEula:
func (s *API) GetModelEula(req *GetModelEulaRequest, opts ...scw.RequestOption) (*Eula, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ModelID) == "" {
		return nil, errors.New("field ModelID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/models/" + fmt.Sprint(req.ModelID) + "/eula",
	}

	var resp Eula

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodeTypes: List all available node types. By default, the node types returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "include_disabled_types", req.IncludeDisabledTypes)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/node-types",
		Query:  query,
	}

	var resp ListNodeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

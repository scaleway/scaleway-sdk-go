// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package llm_inference provides methods and message types of the llm_inference v1beta1 API.
package llm_inference

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/scw"
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
	ListModelsRequestOrderByCreatedAtAsc  = ListModelsRequestOrderBy("created_at_asc")
	ListModelsRequestOrderByCreatedAtDesc = ListModelsRequestOrderBy("created_at_desc")
	ListModelsRequestOrderByNameAsc       = ListModelsRequestOrderBy("name_asc")
	ListModelsRequestOrderByNameDesc      = ListModelsRequestOrderBy("name_desc")
)

func (enum ListModelsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
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
	PrivateNetworkID string `json:"private_network_id"`
}

// EndpointPublicDetails: endpoint public details.
type EndpointPublicDetails struct {
}

// ModelS3Model: model s3 model.
type ModelS3Model struct {
	S3URL string `json:"s3_url"`

	PythonDependencies map[string]string `json:"python_dependencies"`

	NodeType *string `json:"node_type"`

	TritonServerVersion *string `json:"triton_server_version"`
}

// EndpointSpecPrivateNetwork: endpoint spec private network.
type EndpointSpecPrivateNetwork struct {
	PrivateNetworkID string `json:"private_network_id"`
}

// EndpointSpecPublic: endpoint spec public.
type EndpointSpecPublic struct {
}

// Endpoint: endpoint.
type Endpoint struct {
	ID string `json:"id"`

	URL string `json:"url"`

	// Precisely one of PublicAccess, PrivateNetwork must be set.
	PublicAccess *EndpointPublicDetails `json:"public_access,omitempty"`

	// Precisely one of PublicAccess, PrivateNetwork must be set.
	PrivateNetwork *EndpointPrivateNetworkDetails `json:"private_network,omitempty"`

	AuthRequired bool `json:"auth_required"`
}

// Model: model.
type Model struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// Name: name of the model.
	Name string `json:"name"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// Tags: list of tags applied to the model.
	Tags []string `json:"tags"`

	// Description: purpose of the model.
	Description string `json:"description"`

	// S3Model: s3 URL pointing to the model source weight.
	// Precisely one of S3Model must be set.
	S3Model *ModelS3Model `json:"s3_model,omitempty"`

	// IsPublic: defines whether the model is public or not.
	IsPublic bool `json:"is_public"`

	// CreatedAt: creation date of the model.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the model.
	UpdatedAt *time.Time `json:"updated_at"`

	// Region: region of the model.
	Region scw.Region `json:"region"`
}

// ACLRuleRequest: acl rule request.
type ACLRuleRequest struct {
	IP scw.IPNet `json:"ip"`

	Description string `json:"description"`
}

// ACLRule: acl rule.
type ACLRule struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// IP: IP address to be allowed.
	IP scw.IPNet `json:"ip"`

	// Description: description of the IP address.
	Description string `json:"description"`
}

// EndpointSpec: endpoint spec.
type EndpointSpec struct {
	// Precisely one of Public, PrivateNetwork must be set.
	Public *EndpointSpecPublic `json:"public,omitempty"`

	// Precisely one of Public, PrivateNetwork must be set.
	PrivateNetwork *EndpointSpecPrivateNetwork `json:"private_network,omitempty"`

	AuthRequired bool `json:"auth_required"`
}

// Token: token.
type Token struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// BearerToken: authentication token used to perform authentication on the endpoint.
	BearerToken *string `json:"bearer_token"`

	// Description: description of the authentication token usage.
	Description string `json:"description"`

	// CreatedAt: creation date of the authentication token.
	CreatedAt *time.Time `json:"created_at"`

	// ExpiredAt: expiration date of the authentication token.
	ExpiredAt *time.Time `json:"expired_at"`
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

	// Endpoints: list of deployments.
	Endpoints []*Endpoint `json:"endpoints"`

	// Size: current size of the pool.
	Size uint32 `json:"size"`

	// MinSize: defines the minimum size of the pool.
	MinSize uint32 `json:"min_size"`

	// MaxSize: defines the maximum size of the pool.
	MaxSize uint32 `json:"max_size"`

	// ErrorMessage: displays information if your deployment is in error status.
	ErrorMessage *string `json:"error_message"`

	// Model: the inference model used for the deployment.
	Model *Model `json:"model"`

	// CreatedAt: creation date of the deployment.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the deployment.
	UpdatedAt *time.Time `json:"updated_at"`

	// Region: region of the deployment.
	Region scw.Region `json:"region"`
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

	// Region: region of the node type.
	Region scw.Region `json:"region"`
}

// AddDeploymentACLRulesRequest: add deployment acl rules request.
type AddDeploymentACLRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DeploymentID string `json:"-"`

	ACLs []*ACLRuleRequest `json:"acls"`
}

// AddDeploymentACLRulesResponse: add deployment acl rules response.
type AddDeploymentACLRulesResponse struct {
	Rules []*ACLRule `json:"rules"`
}

// AddDeploymentTokenRequest: add deployment token request.
type AddDeploymentTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DeploymentID string `json:"-"`

	Description string `json:"description"`
}

// CreateDeploymentRequest: create deployment request.
type CreateDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ProjectID string `json:"project_id"`

	Name string `json:"name"`

	Tags []string `json:"tags"`

	NodeType string `json:"node_type"`

	ModelID string `json:"model_id"`

	MinSize *uint32 `json:"min_size,omitempty"`

	MaxSize *uint32 `json:"max_size,omitempty"`

	Endpoints []*EndpointSpec `json:"endpoints"`
}

// DeleteDeploymentACLRuleRequest: delete deployment acl rule request.
type DeleteDeploymentACLRuleRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ACLID string `json:"-"`
}

// DeleteDeploymentRequest: delete deployment request.
type DeleteDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DeploymentID string `json:"-"`
}

// DeleteTokenRequest: delete token request.
type DeleteTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	TokenID string `json:"-"`
}

// GetDeploymentRequest: get deployment request.
type GetDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DeploymentID string `json:"-"`
}

// GetModelRequest: get model request.
type GetModelRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ModelID string `json:"-"`
}

// ListDeploymentACLRulesRequest: list deployment acl rules request.
type ListDeploymentACLRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DeploymentID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListDeploymentACLRulesResponse: list deployment acl rules response.
type ListDeploymentACLRulesResponse struct {
	Rules []*ACLRule `json:"rules"`

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

// ListDeploymentTokensRequest: list deployment tokens request.
type ListDeploymentTokensRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DeploymentID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListDeploymentTokensResponse: list deployment tokens response.
type ListDeploymentTokensResponse struct {
	Tokens []*Token `json:"tokens"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDeploymentTokensResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDeploymentTokensResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListDeploymentTokensResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tokens = append(r.Tokens, results.Tokens...)
	r.TotalCount += uint64(len(results.Tokens))
	return uint64(len(results.Tokens)), nil
}

// ListDeploymentsRequest: list deployments request.
type ListDeploymentsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Name *string `json:"-"`

	DeploymentID *string `json:"-"`

	ProjectID *string `json:"-"`

	OrganizationID *string `json:"-"`

	Tags []string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_desc
	OrderBy ListDeploymentsRequestOrderBy `json:"-"`
}

// ListDeploymentsResponse: list deployments response.
type ListDeploymentsResponse struct {
	Deployments []*Deployment `json:"deployments"`

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

	Tags []string `json:"-"`

	Name *string `json:"-"`

	ProjectID *string `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListModelsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListModelsResponse: list models response.
type ListModelsResponse struct {
	Models []*Model `json:"models"`

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

	IncludeDisabledTypes bool `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListNodeTypesResponse: list node types response.
type ListNodeTypesResponse struct {
	NodeTypes []*NodeType `json:"node_types"`

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

	DeploymentID string `json:"-"`

	ACLs []*ACLRuleRequest `json:"acls"`
}

// SetDeploymentACLRulesResponse: set deployment acl rules response.
type SetDeploymentACLRulesResponse struct {
	Rules []*ACLRule `json:"rules"`
}

// UpdateDeploymentRequest: update deployment request.
type UpdateDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DeploymentID string `json:"-"`

	Name *string `json:"name,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	Endpoints []*Endpoint `json:"endpoints"`
}

// Inference API.
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "deployment_id", req.DeploymentID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments",
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
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
	}

	var resp Deployment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDeployment: Create a new inference Deployment related to a specific model.
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
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments",
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

// UpdateDeployment: Update an existing inference Deployment.
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
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
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

// DeleteDeployment: Delete an existing inference Deployment.
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
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
	}

	var resp Deployment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDeploymentTokens: List authentication tokens for a specific deployment.
func (s *API) ListDeploymentTokens(req *ListDeploymentTokensRequest, opts ...scw.RequestOption) (*ListDeploymentTokensResponse, error) {
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
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/auth-tokens",
		Query:  query,
	}

	var resp ListDeploymentTokensResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddDeploymentToken: Add a new authentication token to a specific deployment.
func (s *API) AddDeploymentToken(req *AddDeploymentTokenRequest, opts ...scw.RequestOption) (*Token, error) {
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
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/auth-tokens",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteToken: Delete an existing authentication token.
func (s *API) DeleteToken(req *DeleteTokenRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TokenID) == "" {
		return errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/auth-tokens/" + fmt.Sprint(req.TokenID) + "",
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
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/acls",
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
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/acls",
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
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/acls",
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
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/acls/" + fmt.Sprint(req.ACLID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListModels: List all available LLM models.
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
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/models",
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
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/models/" + fmt.Sprint(req.ModelID) + "",
	}

	var resp Model

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
	parameter.AddToQuery(query, "include_disabled_types", req.IncludeDisabledTypes)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/llm-inference/v1beta1/regions/" + fmt.Sprint(req.Region) + "/node-types",
		Query:  query,
	}

	var resp ListNodeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

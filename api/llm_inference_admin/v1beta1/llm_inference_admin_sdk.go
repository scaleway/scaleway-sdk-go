// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package llm_inference_admin provides methods and message types of the llm_inference_admin v1beta1 API.
package llm_inference_admin

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
	DeploymentStatusDeleted       = DeploymentStatus("deleted")
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

// EndpointPrivateNetworkDetails: endpoint private network details.
type EndpointPrivateNetworkDetails struct {
	PrivateNetworkID string `json:"private_network_id"`
}

// EndpointPublicAccessDetails: endpoint public access details.
type EndpointPublicAccessDetails struct {
}

// Endpoint: endpoint.
type Endpoint struct {
	ID string `json:"id"`

	URL string `json:"url"`

	// Precisely one of PublicAccess, PrivateNetwork must be set.
	PublicAccess *EndpointPublicAccessDetails `json:"public_access,omitempty"`

	// Precisely one of PublicAccess, PrivateNetwork must be set.
	PrivateNetwork *EndpointPrivateNetworkDetails `json:"private_network,omitempty"`

	AuthRequired bool `json:"auth_required"`

	LBID string `json:"lb_id"`
}

// KapsuleCluster: kapsule cluster.
type KapsuleCluster struct {
	// ID: unique identifier.
	ID string `json:"id"`

	// PoolID: node pool unique identifier.
	PoolID string `json:"pool_id"`
}

// Model: model.
type Model struct {
	// Name: name of the model.
	Name string `json:"name"`

	// Tags: list of tags applied to the model.
	Tags []string `json:"tags"`

	// Description: purpose of the model.
	Description string `json:"description"`

	// CreatedAt: creation date of the model.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the model.
	UpdatedAt *time.Time `json:"updated_at"`
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

	// DeletedAt: deployment deletion date.
	DeletedAt *time.Time `json:"deleted_at"`

	// KapsuleCluster: information about the kapsule cluster when deployment is running.
	KapsuleCluster *KapsuleCluster `json:"kapsule_cluster"`

	// Region: region of the deployment.
	Region scw.Region `json:"region"`
}

// ListNodeTypeHistoryResponseEntry: list node type history response entry.
type ListNodeTypeHistoryResponseEntry struct {
	ID string `json:"id"`

	NodeCount uint32 `json:"node_count"`

	NodeType string `json:"node_type"`

	CreatedAt *time.Time `json:"created_at"`

	AvailableAt *time.Time `json:"available_at"`

	ReplacedAt *time.Time `json:"replaced_at"`
}

// GetDeploymentRequest: get deployment request.
type GetDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DeploymentID string `json:"-"`
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

// ListDeploymentsRequest: list deployments request.
type ListDeploymentsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Name *string `json:"-"`

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

// ListNodeTypeHistoryRequest: list node type history request.
type ListNodeTypeHistoryRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DeploymentID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListNodeTypeHistoryResponse: list node type history response.
type ListNodeTypeHistoryResponse struct {
	Entries []*ListNodeTypeHistoryResponseEntry `json:"entries"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNodeTypeHistoryResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNodeTypeHistoryResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListNodeTypeHistoryResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Entries = append(r.Entries, results.Entries...)
	r.TotalCount += uint64(len(results.Entries))
	return uint64(len(results.Entries)), nil
}

// ReplayDeploymentRequest: replay deployment request.
type ReplayDeploymentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DeploymentID string `json:"-"`
}

// ReplayUpdateKapsulePoolRequest: replay update kapsule pool request.
type ReplayUpdateKapsulePoolRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DeploymentID string `json:"-"`
}

// Inference Admin API.
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
		Path:   "/llm-inference-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments",
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
		Path:   "/llm-inference-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "",
	}

	var resp Deployment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReplayDeployment: Redeploy an existing inference Deployment related to a specific model.
func (s *API) ReplayDeployment(req *ReplayDeploymentRequest, opts ...scw.RequestOption) (*Deployment, error) {
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
		Path:   "/llm-inference-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/replay-deployment",
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

// ReplayUpdateKapsulePool: Redeploy an update KapsulePool operation.
func (s *API) ReplayUpdateKapsulePool(req *ReplayUpdateKapsulePoolRequest, opts ...scw.RequestOption) (*Deployment, error) {
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
		Path:   "/llm-inference-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/replay-kapsule-update",
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

// ListNodeTypeHistory:
func (s *API) ListNodeTypeHistory(req *ListNodeTypeHistoryRequest, opts ...scw.RequestOption) (*ListNodeTypeHistoryResponse, error) {
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
		Path:   "/llm-inference-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/node-history",
		Query:  query,
	}

	var resp ListNodeTypeHistoryResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDeploymentACLRules:
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
		Path:   "/llm-inference-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deployments/" + fmt.Sprint(req.DeploymentID) + "/acls",
		Query:  query,
	}

	var resp ListDeploymentACLRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

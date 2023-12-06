// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package redis_admin provides methods and message types of the redis_admin v1 API.
package redis_admin

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
	redis_v1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/redis/v1"
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
	ClusterStatusUnknown      = ClusterStatus("unknown")
	ClusterStatusReady        = ClusterStatus("ready")
	ClusterStatusProvisioning = ClusterStatus("provisioning")
	ClusterStatusConfiguring  = ClusterStatus("configuring")
	ClusterStatusDestroying   = ClusterStatus("destroying")
	ClusterStatusError        = ClusterStatus("error")
	ClusterStatusAutohealing  = ClusterStatus("autohealing")
	ClusterStatusLocked       = ClusterStatus("locked")
	ClusterStatusSuspended    = ClusterStatus("suspended")
	ClusterStatusDestroyed    = ClusterStatus("destroyed")
	ClusterStatusInitializing = ClusterStatus("initializing")
	ClusterStatusDeleting     = ClusterStatus("deleting")
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

type DeploymentStatus string

const (
	DeploymentStatusUnknown    = DeploymentStatus("unknown")
	DeploymentStatusQueued     = DeploymentStatus("queued")
	DeploymentStatusPending    = DeploymentStatus("pending")
	DeploymentStatusRunning    = DeploymentStatus("running")
	DeploymentStatusFailed     = DeploymentStatus("failed")
	DeploymentStatusSucceeded  = DeploymentStatus("succeeded")
	DeploymentStatusRecovering = DeploymentStatus("recovering")
)

func (enum DeploymentStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
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

type ListClusterDeploymentsRequestOrderBy string

const (
	ListClusterDeploymentsRequestOrderByCreatedAtAsc  = ListClusterDeploymentsRequestOrderBy("created_at_asc")
	ListClusterDeploymentsRequestOrderByCreatedAtDesc = ListClusterDeploymentsRequestOrderBy("created_at_desc")
)

func (enum ListClusterDeploymentsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListClusterDeploymentsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListClusterDeploymentsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListClusterDeploymentsRequestOrderBy(ListClusterDeploymentsRequestOrderBy(tmp).String())
	return nil
}

type ListClustersRequestOrderBy string

const (
	ListClustersRequestOrderByCreatedAtAsc  = ListClustersRequestOrderBy("created_at_asc")
	ListClustersRequestOrderByCreatedAtDesc = ListClustersRequestOrderBy("created_at_desc")
	ListClustersRequestOrderByNameAsc       = ListClustersRequestOrderBy("name_asc")
	ListClustersRequestOrderByNameDesc      = ListClustersRequestOrderBy("name_desc")
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

type UpdateClusterRequestStatus string

const (
	UpdateClusterRequestStatusUnknown      = UpdateClusterRequestStatus("unknown")
	UpdateClusterRequestStatusReady        = UpdateClusterRequestStatus("ready")
	UpdateClusterRequestStatusProvisioning = UpdateClusterRequestStatus("provisioning")
	UpdateClusterRequestStatusConfiguring  = UpdateClusterRequestStatus("configuring")
	UpdateClusterRequestStatusDestroying   = UpdateClusterRequestStatus("destroying")
	UpdateClusterRequestStatusError        = UpdateClusterRequestStatus("error")
	UpdateClusterRequestStatusAutohealing  = UpdateClusterRequestStatus("autohealing")
	UpdateClusterRequestStatusLocked       = UpdateClusterRequestStatus("locked")
	UpdateClusterRequestStatusSuspended    = UpdateClusterRequestStatus("suspended")
	UpdateClusterRequestStatusDestroyed    = UpdateClusterRequestStatus("destroyed")
	UpdateClusterRequestStatusInitializing = UpdateClusterRequestStatus("initializing")
)

func (enum UpdateClusterRequestStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum UpdateClusterRequestStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UpdateClusterRequestStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UpdateClusterRequestStatus(UpdateClusterRequestStatus(tmp).String())
	return nil
}

// ACLRule: acl rule.
type ACLRule struct {
	// ID: ID of the rule.
	ID string `json:"id"`

	// IP: iPv4 network address of the rule.
	IP *scw.IPNet `json:"ip"`

	// Description: description of the rule.
	Description *string `json:"description"`
}

// ClusterSetting: cluster setting.
type ClusterSetting struct {
	Name string `json:"name"`

	Value string `json:"value"`
}

// Node: node.
type Node struct {
	ClusterID string `json:"cluster_id"`

	ComputeID string `json:"compute_id"`

	ComputeType string `json:"compute_type"`

	DeletedAt *time.Time `json:"deleted_at"`

	ID string `json:"id"`

	Name string `json:"name"`

	PublicIP net.IP `json:"public_ip"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
}

// Deployment: deployment.
type Deployment struct {
	CreatedAt *time.Time `json:"created_at"`

	ID string `json:"id"`

	LatestMessage string `json:"latest_message"`

	// Status: default value: unknown
	Status DeploymentStatus `json:"status"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// ClusterEvent: cluster event.
type ClusterEvent struct {
	CreatedAt *time.Time `json:"created_at"`

	Event string `json:"event"`

	ID string `json:"id"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// Cluster: cluster.
type Cluster struct {
	// ID: UUID of the Database Instance.
	ID string `json:"id"`

	// Name: name of the Database Instance.
	Name string `json:"name"`

	// ProjectID: project ID the Database Instance belongs to.
	ProjectID string `json:"project_id"`

	OrganizationID string `json:"organization_id"`

	// Status: status of the Database Instance.
	// Default value: unknown
	Status ClusterStatus `json:"status"`

	// Version: redis™ engine version of the Database Instance.
	Version string `json:"version"`

	// Endpoints: list of Database Instance endpoints.
	Endpoints []*redis_v1.Endpoint `json:"endpoints"`

	// Tags: list of tags applied to the Database Instance.
	Tags []string `json:"tags"`

	Nodes []*Node `json:"nodes"`

	// NodeType: node type of the Database Instance.
	NodeType string `json:"node_type"`

	// CreatedAt: creation date (Format ISO 8601).
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: update date (Format ISO 8601).
	UpdatedAt *time.Time `json:"updated_at"`

	// TLSEnabled: defines whether or not TLS is enabled.
	TLSEnabled bool `json:"tls_enabled"`

	// ClusterSettings: list of Database Instance settings.
	ClusterSettings []*ClusterSetting `json:"cluster_settings"`

	// ACLRules: list of ACL rules.
	ACLRules []*ACLRule `json:"acl_rules"`

	// ClusterSize: number of nodes of the Database Instance cluster.
	ClusterSize uint32 `json:"cluster_size"`

	// Zone: zone of the Database Instance.
	Zone scw.Zone `json:"zone"`

	// Username: name of the user associated to the cluster.
	Username string `json:"username"`

	AvailableAt *time.Time `json:"available_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	LockReason string `json:"lock_reason"`
}

// Resource: resource.
type Resource struct {
	ResourceID string `json:"resource_id"`

	ResourceType string `json:"resource_type"`
}

// DeleteClusterRequest: delete cluster request.
type DeleteClusterRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`
}

// FailoverClusterRequest: failover cluster request.
type FailoverClusterRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	NodeID string `json:"node_id"`
}

// GetClusterRequest: get cluster request.
type GetClusterRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`
}

// HotNodeReplaceRequest: hot node replace request.
type HotNodeReplaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	NodeID string `json:"node_id"`
}

// ListClusterDeploymentsRequest: list cluster deployments request.
type ListClusterDeploymentsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListClusterDeploymentsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListClusterDeploymentsResponse: list cluster deployments response.
type ListClusterDeploymentsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Deployments []*Deployment `json:"deployments"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClusterDeploymentsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClusterDeploymentsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListClusterDeploymentsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Deployments = append(r.Deployments, results.Deployments...)
	r.TotalCount += uint32(len(results.Deployments))
	return uint32(len(results.Deployments)), nil
}

// ListClusterEventsRequest: list cluster events request.
type ListClusterEventsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListClusterEventsResponse: list cluster events response.
type ListClusterEventsResponse struct {
	TotalCount uint32 `json:"total_count"`

	ClusterEvents []*ClusterEvent `json:"cluster_events"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClusterEventsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClusterEventsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListClusterEventsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ClusterEvents = append(r.ClusterEvents, results.ClusterEvents...)
	r.TotalCount += uint32(len(results.ClusterEvents))
	return uint32(len(results.ClusterEvents)), nil
}

// ListClustersRequest: list clusters request.
type ListClustersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ExcludeStatuses []ClusterStatus `json:"-"`

	IncludeStatuses []ClusterStatus `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListClustersRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListClustersResponse: list clusters response.
type ListClustersResponse struct {
	TotalCount uint32 `json:"total_count"`

	Clusters []*Cluster `json:"clusters"`
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

// ListResourcesRequest: list resources request.
type ListResourcesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListResourcesResponse: list resources response.
type ListResourcesResponse struct {
	ClusterID string `json:"cluster_id"`

	Resources []*Resource `json:"resources"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListResourcesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListResourcesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListResourcesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Resources = append(r.Resources, results.Resources...)
	r.TotalCount += uint32(len(results.Resources))
	return uint32(len(results.Resources)), nil
}

// RedeployClusterRequest: redeploy cluster request.
type RedeployClusterRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`
}

// UpdateClusterRequest: update cluster request.
type UpdateClusterRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	// Status: default value: unknown
	Status UpdateClusterRequestStatus `json:"status"`
}

type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}
func (s *API) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar2}
}

// RedeployCluster:
func (s *API) RedeployCluster(req *RedeployClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
		Method: "POST",
		Path:   "/redis-admin/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/redeploy",
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

// ListResources:
func (s *API) ListResources(req *ListResourcesRequest, opts ...scw.RequestOption) (*ListResourcesResponse, error) {
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

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/redis-admin/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/resources",
		Query:  query,
	}

	var resp ListResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// HotNodeReplace:
func (s *API) HotNodeReplace(req *HotNodeReplaceRequest, opts ...scw.RequestOption) (*Node, error) {
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
		Method: "POST",
		Path:   "/redis-admin/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/hot-node-replace",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Node

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCluster: Delete a Redis™ Database Instance (Redis™ cluster), specified by the `region` and `cluster_id` parameters. Deleting a Database Instance is permanent, and cannot be undone. Note that upon deletion all your data will be lost.
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
		Method: "DELETE",
		Path:   "/redis-admin/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCluster: Retrieve information about a Redis™ Database Instance (Redis™ cluster). Specify the `cluster_id` and `region` in your request to get information such as `id`, `status`, `version`, `tls_enabled`, `cluster_settings`, `upgradable_versions` and `endpoints` about your cluster in the response.
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
		Method: "GET",
		Path:   "/redis-admin/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusters:
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
	parameter.AddToQuery(query, "exclude_statuses", req.ExcludeStatuses)
	parameter.AddToQuery(query, "include_statuses", req.IncludeStatuses)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/redis-admin/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters",
		Query:  query,
	}

	var resp ListClustersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterEvents:
func (s *API) ListClusterEvents(req *ListClusterEventsRequest, opts ...scw.RequestOption) (*ListClusterEventsResponse, error) {
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

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/redis-admin/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/events",
		Query:  query,
	}

	var resp ListClusterEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterDeployments:
func (s *API) ListClusterDeployments(req *ListClusterDeploymentsRequest, opts ...scw.RequestOption) (*ListClusterDeploymentsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/redis-admin/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/deployments",
		Query:  query,
	}

	var resp ListClusterDeploymentsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCluster: Update the parameters of a Redis™ Database Instance (Redis™ cluster), including `name`, `tags`, `user_name` and `password`.
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
		Method: "PATCH",
		Path:   "/redis-admin/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
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

// FailoverCluster:
func (s *API) FailoverCluster(req *FailoverClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
		Method: "POST",
		Path:   "/redis-admin/v1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/failover",
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

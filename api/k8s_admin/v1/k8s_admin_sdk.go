// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package k8s_admin provides methods and message types of the k8s_admin v1 API.
package k8s_admin

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
	k8s_v1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/k8s/v1"
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

type ListClustersRequestOrderBy string

const (
	ListClustersRequestOrderByCreatedAtAsc  = ListClustersRequestOrderBy("created_at_asc")
	ListClustersRequestOrderByCreatedAtDesc = ListClustersRequestOrderBy("created_at_desc")
	ListClustersRequestOrderByUpdatedAtAsc  = ListClustersRequestOrderBy("updated_at_asc")
	ListClustersRequestOrderByUpdatedAtDesc = ListClustersRequestOrderBy("updated_at_desc")
	ListClustersRequestOrderByNameAsc       = ListClustersRequestOrderBy("name_asc")
	ListClustersRequestOrderByNameDesc      = ListClustersRequestOrderBy("name_desc")
	ListClustersRequestOrderByStatusAsc     = ListClustersRequestOrderBy("status_asc")
	ListClustersRequestOrderByStatusDesc    = ListClustersRequestOrderBy("status_desc")
	ListClustersRequestOrderByVersionAsc    = ListClustersRequestOrderBy("version_asc")
	ListClustersRequestOrderByVersionDesc   = ListClustersRequestOrderBy("version_desc")
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

// ClusterMemoryLimits: cluster memory limits.
type ClusterMemoryLimits struct {
	APIServer string `json:"api_server"`

	ControllerManager string `json:"controller_manager"`
}

// MasterStatusPodContainer: master status pod container.
type MasterStatusPodContainer struct {
	Name string `json:"name"`

	State string `json:"state"`

	Restarts uint32 `json:"restarts"`

	CPUUsage string `json:"cpu_usage"`

	MemoryUsage string `json:"memory_usage"`
}

// Cluster: cluster.
type Cluster struct {
	EtcdID *string `json:"etcd_id"`

	LockReason string `json:"lock_reason"`

	MemoryLimits *ClusterMemoryLimits `json:"memory_limits"`

	PublicCluster *k8s_v1.Cluster `json:"public_cluster"`

	SeedID *string `json:"seed_id"`
}

// Etcd: etcd.
type Etcd struct {
	Capacity uint32 `json:"capacity"`

	ID string `json:"id"`

	Members []string `json:"members"`

	Name string `json:"name"`

	Provisionable bool `json:"provisionable"`

	Type string `json:"type"`

	Usage uint32 `json:"usage"`
}

// Seed: seed.
type Seed struct {
	Capacity uint32 `json:"capacity"`

	ClusterTypes []string `json:"cluster_types"`

	ID string `json:"id"`

	KubeServer string `json:"kube_server"`

	Name string `json:"name"`

	Provisionable bool `json:"provisionable"`

	Usage uint32 `json:"usage"`
}

// MasterStatusPod: master status pod.
type MasterStatusPod struct {
	Name string `json:"name"`

	Containers []*MasterStatusPodContainer `json:"containers"`
}

// UpdateClusterRequestMemoryLimits: update cluster request memory limits.
type UpdateClusterRequestMemoryLimits struct {
	APIServer *string `json:"api_server"`

	ControllerManager *string `json:"controller_manager"`
}

// CreateClusterRequest: create cluster request.
type CreateClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	CreateCluster *k8s_v1.CreateClusterRequest `json:"create_cluster,omitempty"`

	EtcdID *string `json:"etcd_id,omitempty"`

	SeedID *string `json:"seed_id,omitempty"`
}

// EmptySeedRequest: empty seed request.
type EmptySeedRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	SeedID string `json:"-"`
}

// EmptySeedResponse: empty seed response.
type EmptySeedResponse struct {
	TaskID string `json:"task_id"`
}

// EnsureClusterRequest: ensure cluster request.
type EnsureClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ClusterID string `json:"-"`

	Level int32 `json:"level"`
}

// EnsureClusterResponse: ensure cluster response.
type EnsureClusterResponse struct {
	TaskID string `json:"task_id"`
}

// GetClusterRequest: get cluster request.
type GetClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ClusterID string `json:"-"`
}

// GetEtcdRequest: get etcd request.
type GetEtcdRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	EtcdID string `json:"-"`
}

// GetMasterStatusRequest: get master status request.
type GetMasterStatusRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ClusterID string `json:"-"`
}

// GetNodeMetadataRequest: get node metadata request.
type GetNodeMetadataRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	InstanceID string `json:"-"`

	SourcePort uint32 `json:"-"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// GetSeedRequest: get seed request.
type GetSeedRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	SeedID string `json:"-"`
}

// ListClustersRequest: list clusters request.
type ListClustersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	EtcdID *string `json:"-"`

	IncludeDeleted bool `json:"-"`

	Name *string `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListClustersRequestOrderBy `json:"-"`

	OrganizationID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	PrivateNetworkID *string `json:"-"`

	ProjectID *string `json:"-"`

	SeedID *string `json:"-"`

	// Status: default value: unknown
	Status k8s_v1.ClusterStatus `json:"-"`
}

// ListClustersResponse: list clusters response.
type ListClustersResponse struct {
	Clusters []*Cluster `json:"clusters"`

	TotalCount uint32 `json:"total_count"`
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

// ListEtcdsRequest: list etcds request.
type ListEtcdsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Name *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Provisionable *bool `json:"-"`

	Type *string `json:"-"`
}

// ListEtcdsResponse: list etcds response.
type ListEtcdsResponse struct {
	Etcds []*Etcd `json:"etcds"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListEtcdsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListEtcdsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListEtcdsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Etcds = append(r.Etcds, results.Etcds...)
	r.TotalCount += uint32(len(results.Etcds))
	return uint32(len(results.Etcds)), nil
}

// ListSeedsRequest: list seeds request.
type ListSeedsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Name *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Provisionable *bool `json:"-"`
}

// ListSeedsResponse: list seeds response.
type ListSeedsResponse struct {
	Seeds []*Seed `json:"seeds"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSeedsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSeedsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSeedsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Seeds = append(r.Seeds, results.Seeds...)
	r.TotalCount += uint32(len(results.Seeds))
	return uint32(len(results.Seeds)), nil
}

// MasterStatus: master status.
type MasterStatus struct {
	Pods []*MasterStatusPod `json:"pods"`
}

// MoveClusterRequest: move cluster request.
type MoveClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ClusterID string `json:"-"`

	SeedID string `json:"seed_id"`
}

// MoveClusterResponse: move cluster response.
type MoveClusterResponse struct {
	TaskID string `json:"task_id"`
}

// NodeMetadata: node metadata.
type NodeMetadata struct {
	ClusterID string `json:"cluster_id"`
}

// SetClusterTypeRequest: set cluster type request.
type SetClusterTypeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ClusterID string `json:"-"`

	Type string `json:"type"`
}

// UpdateClusterRequest: update cluster request.
type UpdateClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ClusterID string `json:"-"`

	MemoryLimits *UpdateClusterRequestMemoryLimits `json:"memory_limits,omitempty"`
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
func (s *API) Regions() []scw.Region {
	return []scw.Region{}
}

// ListClusters:
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
	parameter.AddToQuery(query, "etcd_id", req.EtcdID)
	parameter.AddToQuery(query, "include_deleted", req.IncludeDeleted)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "private_network_id", req.PrivateNetworkID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "seed_id", req.SeedID)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/clusters",
		Query:  query,
	}

	var resp ListClustersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCluster:
func (s *API) CreateCluster(req *CreateClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/clusters",
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

// GetCluster:
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
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnsureCluster:
func (s *API) EnsureCluster(req *EnsureClusterRequest, opts ...scw.RequestOption) (*EnsureClusterResponse, error) {
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
		Method: "POST",
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/ensure",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EnsureClusterResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MoveCluster:
func (s *API) MoveCluster(req *MoveClusterRequest, opts ...scw.RequestOption) (*MoveClusterResponse, error) {
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
		Method: "POST",
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/move",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp MoveClusterResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMasterStatus:
func (s *API) GetMasterStatus(req *GetMasterStatusRequest, opts ...scw.RequestOption) (*MasterStatus, error) {
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
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/master-status",
	}

	var resp MasterStatus

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCluster:
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
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
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

// SetClusterType:
func (s *API) SetClusterType(req *SetClusterTypeRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
		Method: "POST",
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/set-type",
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

// GetNodeMetadata:
func (s *API) GetNodeMetadata(req *GetNodeMetadataRequest, opts ...scw.RequestOption) (*NodeMetadata, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "instance_id", req.InstanceID)
	parameter.AddToQuery(query, "source_port", req.SourcePort)
	parameter.AddToQuery(query, "zone", req.Zone)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/node-metadata",
		Query:  query,
	}

	var resp NodeMetadata

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSeeds:
func (s *API) ListSeeds(req *ListSeedsRequest, opts ...scw.RequestOption) (*ListSeedsResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "provisionable", req.Provisionable)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/seeds",
		Query:  query,
	}

	var resp ListSeedsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSeed:
func (s *API) GetSeed(req *GetSeedRequest, opts ...scw.RequestOption) (*Seed, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SeedID) == "" {
		return nil, errors.New("field SeedID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/seeds/" + fmt.Sprint(req.SeedID) + "",
	}

	var resp Seed

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EmptySeed:
func (s *API) EmptySeed(req *EmptySeedRequest, opts ...scw.RequestOption) (*EmptySeedResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SeedID) == "" {
		return nil, errors.New("field SeedID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/seeds/" + fmt.Sprint(req.SeedID) + "/empty",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EmptySeedResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListEtcds:
func (s *API) ListEtcds(req *ListEtcdsRequest, opts ...scw.RequestOption) (*ListEtcdsResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "provisionable", req.Provisionable)
	parameter.AddToQuery(query, "type", req.Type)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/etcds",
		Query:  query,
	}

	var resp ListEtcdsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEtcd:
func (s *API) GetEtcd(req *GetEtcdRequest, opts ...scw.RequestOption) (*Etcd, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.EtcdID) == "" {
		return nil, errors.New("field EtcdID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s-admin/v1/regions/" + fmt.Sprint(req.Region) + "/etcds/" + fmt.Sprint(req.EtcdID) + "",
	}

	var resp Etcd

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

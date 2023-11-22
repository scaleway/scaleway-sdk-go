// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package k8s_private provides methods and message types of the k8s_private v1beta3 API.
package k8s_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/k8s_private/v1beta3/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/k8s_private/v1beta3/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/k8s_private/v1beta3/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/k8s_private/v1beta3/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/k8s_private/v1beta3/scw"
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

type EasyDeployStatus string

const (
	EasyDeployStatusUnknown  = EasyDeployStatus("unknown")
	EasyDeployStatusPending  = EasyDeployStatus("pending")
	EasyDeployStatusDeployed = EasyDeployStatus("deployed")
	EasyDeployStatusError    = EasyDeployStatus("error")
	EasyDeployStatusDeleted  = EasyDeployStatus("deleted")
	EasyDeployStatusDeleting = EasyDeployStatus("deleting")
)

func (enum EasyDeployStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum EasyDeployStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EasyDeployStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EasyDeployStatus(EasyDeployStatus(tmp).String())
	return nil
}

// EasyDeployEnvVar: easy deploy env var.
type EasyDeployEnvVar struct {
	// Name: env name.
	Name string `json:"name"`

	// Value: env value.
	Value string `json:"value"`
}

// EasyDeployCron: easy deploy cron.
type EasyDeployCron struct {
	// Image: docker image (eg. namespace/image:tag).
	Image string `json:"image"`

	// ScheduleMinute: minute (0 - 59).
	ScheduleMinute string `json:"schedule_minute"`

	// ScheduleHour: hour (0 - 23).
	ScheduleHour string `json:"schedule_hour"`

	// ScheduleMonthday: day of the month (1 - 31).
	ScheduleMonthday string `json:"schedule_monthday"`

	// ScheduleMonth: month (1 - 12).
	ScheduleMonth string `json:"schedule_month"`

	// ScheduleWeekday: day of the week (0 - 6).
	ScheduleWeekday string `json:"schedule_weekday"`

	Environments []*EasyDeployEnvVar `json:"environments"`

	// Command: command.
	Command string `json:"command"`

	// Args: a list of arguments.
	Args []string `json:"args"`
}

// EasyDeployDeploy: easy deploy deploy.
type EasyDeployDeploy struct {
	// Image: docker image (eg. namespace/image:tag).
	Image string `json:"image"`

	// Replicas: number of replicas (0,N).
	Replicas uint32 `json:"replicas"`

	// Ports: list of pods to expose.
	Ports []string `json:"ports"`

	// Loadbalancer: enable exposure through a loadbalancer.
	Loadbalancer bool `json:"loadbalancer"`

	Environments []*EasyDeployEnvVar `json:"environments"`

	// Command: command.
	Command string `json:"command"`

	// Args: a list of arguments.
	Args []string `json:"args"`
}

// EasyDeployHelm: easy deploy helm.
type EasyDeployHelm struct {
	// Chart: chart package name (eg. stable/nginx).
	Chart string `json:"chart"`

	// Version: version.
	Version string `json:"version"`

	// Values: a YAML that override default value.yaml.
	Values string `json:"values"`
}

// CatalogItem: catalog item.
type CatalogItem struct {
	Name string `json:"name"`

	Description string `json:"description"`

	Icon string `json:"icon"`

	Readme string `json:"readme"`

	Version string `json:"version"`

	ProposedValue string `json:"proposed_value"`

	Tags []string `json:"tags"`

	ID string `json:"id"`
}

// EasyDeploy: easy deploy.
type EasyDeploy struct {
	// Region: region.
	Region scw.Region `json:"region"`

	// ID: iD.
	ID string `json:"id"`

	// ClusterID: cluster ID.
	ClusterID string `json:"cluster_id"`

	// CreatedAt: created At.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: updated At.
	UpdatedAt *time.Time `json:"updated_at"`

	// Name: name.
	Name string `json:"name"`

	// Namespace: namespace.
	Namespace string `json:"namespace"`

	// Status: status.
	// Default value: unknown
	Status EasyDeployStatus `json:"status"`

	// Note: note.
	Note string `json:"note"`

	// Precisely one of Helm, Deploy, Cron must be set.
	Helm *EasyDeployHelm `json:"helm,omitempty"`

	// Precisely one of Helm, Deploy, Cron must be set.
	Deploy *EasyDeployDeploy `json:"deploy,omitempty"`

	// Precisely one of Helm, Deploy, Cron must be set.
	Cron *EasyDeployCron `json:"cron,omitempty"`
}

// EasyDeployAPICreateEasyDeployRequest: easy deploy api create easy deploy request.
type EasyDeployAPICreateEasyDeployRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster.
	ClusterID string `json:"-"`

	// Name: name.
	Name string `json:"name"`

	// Namespace: namespace.
	Namespace *string `json:"namespace,omitempty"`

	// Precisely one of Helm, Deploy, Cron must be set.
	Helm *EasyDeployHelm `json:"helm,omitempty"`

	// Precisely one of Helm, Deploy, Cron must be set.
	Deploy *EasyDeployDeploy `json:"deploy,omitempty"`

	// Precisely one of Helm, Deploy, Cron must be set.
	Cron *EasyDeployCron `json:"cron,omitempty"`
}

// EasyDeployAPIDeleteEasyDeployRequest: easy deploy api delete easy deploy request.
type EasyDeployAPIDeleteEasyDeployRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster.
	ClusterID string `json:"-"`

	// DeployID: ID of the deploy.
	DeployID string `json:"-"`
}

// EasyDeployAPIGetEasyDeployRequest: easy deploy api get easy deploy request.
type EasyDeployAPIGetEasyDeployRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster.
	ClusterID string `json:"-"`

	// DeployID: ID of the deploy.
	DeployID string `json:"-"`
}

// EasyDeployAPIListCatalogRequest: easy deploy api list catalog request.
type EasyDeployAPIListCatalogRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Tags []string `json:"-"`

	Name *string `json:"-"`

	IncludeHidden bool `json:"-"`
}

// EasyDeployAPIListCatalogTagsRequest: easy deploy api list catalog tags request.
type EasyDeployAPIListCatalogTagsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// EasyDeployAPIListEasyDeployRequest: easy deploy api list easy deploy request.
type EasyDeployAPIListEasyDeployRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster.
	ClusterID string `json:"-"`

	// Page: page.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// Name: name.
	Name *string `json:"-"`

	// Namespace: namespace.
	Namespace *string `json:"-"`
}

// EasyDeployAPIUpdateEasyDeployRequest: easy deploy api update easy deploy request.
type EasyDeployAPIUpdateEasyDeployRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClusterID: ID of the cluster.
	ClusterID string `json:"-"`

	// DeployID: ID of the deploy.
	DeployID string `json:"-"`

	// Precisely one of Helm, Deploy, Cron must be set.
	Helm *EasyDeployHelm `json:"helm,omitempty"`

	// Precisely one of Helm, Deploy, Cron must be set.
	Deploy *EasyDeployDeploy `json:"deploy,omitempty"`

	// Precisely one of Helm, Deploy, Cron must be set.
	Cron *EasyDeployCron `json:"cron,omitempty"`
}

// ListCatalogResponse: list catalog response.
type ListCatalogResponse struct {
	TotalCount uint32 `json:"total_count"`

	Catalog []*CatalogItem `json:"catalog"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCatalogResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCatalogResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCatalogResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Catalog = append(r.Catalog, results.Catalog...)
	r.TotalCount += uint32(len(results.Catalog))
	return uint32(len(results.Catalog)), nil
}

// ListCatalogTagsResponse: list catalog tags response.
type ListCatalogTagsResponse struct {
	Tags []string `json:"tags"`
}

// ListEasyDeployResponse: list easy deploy response.
type ListEasyDeployResponse struct {
	// TotalCount: total count.
	TotalCount uint32 `json:"total_count"`

	// EasyDeploys: easy deploys.
	EasyDeploys []*EasyDeploy `json:"easy_deploys"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListEasyDeployResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListEasyDeployResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListEasyDeployResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.EasyDeploys = append(r.EasyDeploys, results.EasyDeploys...)
	r.TotalCount += uint32(len(results.EasyDeploys))
	return uint32(len(results.EasyDeploys)), nil
}

// Kapsule EasyDeploy API.
type EasyDeployAPI struct {
	client *scw.Client
}

// NewEasyDeployAPI returns a EasyDeployAPI object from a Scaleway client.
func NewEasyDeployAPI(client *scw.Client) *EasyDeployAPI {
	return &EasyDeployAPI{
		client: client,
	}
}
func (s *EasyDeployAPI) Regions() []scw.Region {
	return []scw.Region{}
}

// ListEasyDeploy: List the easy deploys.
func (s *EasyDeployAPI) ListEasyDeploy(req *EasyDeployAPIListEasyDeployRequest, opts ...scw.RequestOption) (*ListEasyDeployResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "namespace", req.Namespace)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s-private/v1beta3/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/easy-deploy",
		Query:  query,
	}

	var resp ListEasyDeployResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateEasyDeploy: Create an easy deploy.
func (s *EasyDeployAPI) CreateEasyDeploy(req *EasyDeployAPICreateEasyDeployRequest, opts ...scw.RequestOption) (*EasyDeploy, error) {
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
		Path:   "/k8s-private/v1beta3/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/easy-deploy",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EasyDeploy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetEasyDeploy: Get an easy deploy.
func (s *EasyDeployAPI) GetEasyDeploy(req *EasyDeployAPIGetEasyDeployRequest, opts ...scw.RequestOption) (*EasyDeploy, error) {
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

	if fmt.Sprint(req.DeployID) == "" {
		return nil, errors.New("field DeployID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s-private/v1beta3/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/easy-deploy/" + fmt.Sprint(req.DeployID) + "",
	}

	var resp EasyDeploy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateEasyDeploy: Update an easy deploy.
func (s *EasyDeployAPI) UpdateEasyDeploy(req *EasyDeployAPIUpdateEasyDeployRequest, opts ...scw.RequestOption) (*EasyDeploy, error) {
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

	if fmt.Sprint(req.DeployID) == "" {
		return nil, errors.New("field DeployID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/k8s-private/v1beta3/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/easy-deploy/" + fmt.Sprint(req.DeployID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EasyDeploy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteEasyDeploy: Delete an easy deploy.
func (s *EasyDeployAPI) DeleteEasyDeploy(req *EasyDeployAPIDeleteEasyDeployRequest, opts ...scw.RequestOption) (*EasyDeploy, error) {
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

	if fmt.Sprint(req.DeployID) == "" {
		return nil, errors.New("field DeployID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/k8s-private/v1beta3/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/easy-deploy/" + fmt.Sprint(req.DeployID) + "",
	}

	var resp EasyDeploy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCatalog:
func (s *EasyDeployAPI) ListCatalog(req *EasyDeployAPIListCatalogRequest, opts ...scw.RequestOption) (*ListCatalogResponse, error) {
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
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "include_hidden", req.IncludeHidden)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s-private/v1beta3/regions/" + fmt.Sprint(req.Region) + "/catalog/easy-deploy",
		Query:  query,
	}

	var resp ListCatalogResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCatalogTags:
func (s *EasyDeployAPI) ListCatalogTags(req *EasyDeployAPIListCatalogTagsRequest, opts ...scw.RequestOption) (*ListCatalogTagsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s-private/v1beta3/regions/" + fmt.Sprint(req.Region) + "/catalog/easy-deploy-tags",
	}

	var resp ListCatalogTagsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

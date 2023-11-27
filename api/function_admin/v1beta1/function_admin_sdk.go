// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package function_admin provides methods and message types of the function_admin v1beta1 API.
package function_admin

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

type ClusterStatus string

const (
	ClusterStatusUnknown      = ClusterStatus("unknown")
	ClusterStatusProvisioned  = ClusterStatus("provisioned")
	ClusterStatusProvisioning = ClusterStatus("provisioning")
	ClusterStatusError        = ClusterStatus("error")
	ClusterStatusCreating     = ClusterStatus("creating")
	ClusterStatusCreated      = ClusterStatus("created")
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

type FunctionPrivacy string

const (
	FunctionPrivacyUnknownPrivacy = FunctionPrivacy("unknown_privacy")
	FunctionPrivacyPublic         = FunctionPrivacy("public")
	FunctionPrivacyPrivate        = FunctionPrivacy("private")
)

func (enum FunctionPrivacy) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_privacy"
	}
	return string(enum)
}

func (enum FunctionPrivacy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionPrivacy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionPrivacy(FunctionPrivacy(tmp).String())
	return nil
}

type FunctionStatus string

const (
	FunctionStatusUnknown  = FunctionStatus("unknown")
	FunctionStatusReady    = FunctionStatus("ready")
	FunctionStatusDeleting = FunctionStatus("deleting")
	FunctionStatusError    = FunctionStatus("error")
	FunctionStatusLocked   = FunctionStatus("locked")
	FunctionStatusCreating = FunctionStatus("creating")
	FunctionStatusPending  = FunctionStatus("pending")
	FunctionStatusCreated  = FunctionStatus("created")
)

func (enum FunctionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum FunctionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionStatus(FunctionStatus(tmp).String())
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

type ListFunctionsRequestOrderBy string

const (
	ListFunctionsRequestOrderByCreatedAtAsc  = ListFunctionsRequestOrderBy("created_at_asc")
	ListFunctionsRequestOrderByCreatedAtDesc = ListFunctionsRequestOrderBy("created_at_desc")
	ListFunctionsRequestOrderByNameAsc       = ListFunctionsRequestOrderBy("name_asc")
	ListFunctionsRequestOrderByNameDesc      = ListFunctionsRequestOrderBy("name_desc")
)

func (enum ListFunctionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListFunctionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFunctionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFunctionsRequestOrderBy(ListFunctionsRequestOrderBy(tmp).String())
	return nil
}

type ListLogsRequestOrderBy string

const (
	ListLogsRequestOrderByTimestampDesc = ListLogsRequestOrderBy("timestamp_desc")
	ListLogsRequestOrderByTimestampAsc  = ListLogsRequestOrderBy("timestamp_asc")
)

func (enum ListLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "timestamp_desc"
	}
	return string(enum)
}

func (enum ListLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListLogsRequestOrderBy(ListLogsRequestOrderBy(tmp).String())
	return nil
}

type ListNamespacesRequestOrderBy string

const (
	ListNamespacesRequestOrderByCreatedAtAsc  = ListNamespacesRequestOrderBy("created_at_asc")
	ListNamespacesRequestOrderByCreatedAtDesc = ListNamespacesRequestOrderBy("created_at_desc")
	ListNamespacesRequestOrderByNameAsc       = ListNamespacesRequestOrderBy("name_asc")
	ListNamespacesRequestOrderByNameDesc      = ListNamespacesRequestOrderBy("name_desc")
)

func (enum ListNamespacesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNamespacesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNamespacesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNamespacesRequestOrderBy(ListNamespacesRequestOrderBy(tmp).String())
	return nil
}

type LogStream string

const (
	LogStreamUnknown = LogStream("unknown")
	LogStreamStdout  = LogStream("stdout")
	LogStreamStderr  = LogStream("stderr")
)

func (enum LogStream) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum LogStream) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LogStream) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LogStream(LogStream(tmp).String())
	return nil
}

type NamespaceStatus string

const (
	NamespaceStatusUnknown  = NamespaceStatus("unknown")
	NamespaceStatusReady    = NamespaceStatus("ready")
	NamespaceStatusDeleting = NamespaceStatus("deleting")
	NamespaceStatusError    = NamespaceStatus("error")
	NamespaceStatusLocked   = NamespaceStatus("locked")
	NamespaceStatusCreating = NamespaceStatus("creating")
	NamespaceStatusPending  = NamespaceStatus("pending")
)

func (enum NamespaceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NamespaceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NamespaceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NamespaceStatus(NamespaceStatus(tmp).String())
	return nil
}

type ReconcileTriggerStatusesResponseStatus string

const (
	ReconcileTriggerStatusesResponseStatusUnknownStatus = ReconcileTriggerStatusesResponseStatus("unknown_status")
	ReconcileTriggerStatusesResponseStatusReady         = ReconcileTriggerStatusesResponseStatus("ready")
	ReconcileTriggerStatusesResponseStatusDeleting      = ReconcileTriggerStatusesResponseStatus("deleting")
	ReconcileTriggerStatusesResponseStatusError         = ReconcileTriggerStatusesResponseStatus("error")
	ReconcileTriggerStatusesResponseStatusCreating      = ReconcileTriggerStatusesResponseStatus("creating")
	ReconcileTriggerStatusesResponseStatusPending       = ReconcileTriggerStatusesResponseStatus("pending")
)

func (enum ReconcileTriggerStatusesResponseStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum ReconcileTriggerStatusesResponseStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReconcileTriggerStatusesResponseStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReconcileTriggerStatusesResponseStatus(ReconcileTriggerStatusesResponseStatus(tmp).String())
	return nil
}

// FunctionDebugDataPodContainer: function debug data pod container.
type FunctionDebugDataPodContainer struct {
	Name string `json:"name"`

	Status string `json:"status"`

	Restarts uint32 `json:"restarts"`

	Reason string `json:"reason"`

	StartedAt string `json:"started_at"`
}

// SecretHashedValue: secret hashed value.
type SecretHashedValue struct {
	HashedValue string `json:"hashed_value"`

	Key string `json:"key"`
}

// NamespaceCluster: namespace cluster.
type NamespaceCluster struct {
	ID string `json:"id"`

	KapsuleID string `json:"kapsule_id"`

	Name string `json:"name"`

	Version string `json:"version"`
}

// FunctionDebugDataBuild: function debug data build.
type FunctionDebugDataBuild struct {
	Name string `json:"name"`

	Succeeded bool `json:"succeeded"`

	Reason string `json:"reason"`
}

// FunctionDebugDataKsvc: function debug data ksvc.
type FunctionDebugDataKsvc struct {
	Name string `json:"name"`

	Ready bool `json:"ready"`

	Reason string `json:"reason"`

	Revision string `json:"revision"`
}

// FunctionDebugDataPod: function debug data pod.
type FunctionDebugDataPod struct {
	Name string `json:"name"`

	Containers []*FunctionDebugDataPodContainer `json:"containers"`
}

// Cluster: cluster.
type Cluster struct {
	Billable bool `json:"billable"`

	CanProvision bool `json:"can_provision"`

	CreatedAt *time.Time `json:"created_at"`

	Environment string `json:"environment"`

	ErrorMessage *string `json:"error_message"`

	EtcdID *string `json:"etcd_id"`

	ID string `json:"id"`

	KapsuleID *string `json:"kapsule_id"`

	KapsuleType *string `json:"kapsule_type"`

	LoadBalancerIP *string `json:"load_balancer_ip"`

	Name string `json:"name"`

	ProvisionedAt *time.Time `json:"provisioned_at"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`

	SeedID *string `json:"seed_id"`

	// Status: default value: unknown
	Status ClusterStatus `json:"status"`

	Version uint32 `json:"version"`
}

// Function: function.
type Function struct {
	CPULimit uint32 `json:"cpu_limit"`

	CreatedAt *time.Time `json:"created_at"`

	Description *string `json:"description"`

	DomainName string `json:"domain_name"`

	EnvironmentVariables map[string]string `json:"environment_variables"`

	ErrorMessage *string `json:"error_message"`

	Handler string `json:"handler"`

	ID string `json:"id"`

	MaxScale uint32 `json:"max_scale"`

	MemoryLimit uint32 `json:"memory_limit"`

	MinScale uint32 `json:"min_scale"`

	Name string `json:"name"`

	NamespaceID string `json:"namespace_id"`

	// Privacy: default value: unknown_privacy
	Privacy FunctionPrivacy `json:"privacy"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`

	RegistryImage string `json:"registry_image"`

	Runtime string `json:"runtime"`

	RuntimeMessage string `json:"runtime_message"`

	SecretEnvironmentVariables []*SecretHashedValue `json:"secret_environment_variables"`

	// Status: default value: unknown
	Status FunctionStatus `json:"status"`

	Timeout *scw.Duration `json:"timeout"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// Log: log.
type Log struct {
	// Message: message of the log.
	Message string `json:"message"`

	// Timestamp: timestamp of the log.
	Timestamp *time.Time `json:"timestamp"`

	// ID: UUID of the log.
	ID string `json:"id"`

	// Level: severity of the log (info, debug, error etc.).
	Level string `json:"level"`

	// Source: source of the log (core runtime or user code).
	Source string `json:"source"`

	// Stream: can be stdout or stderr.
	// Default value: unknown
	Stream LogStream `json:"stream"`
}

// Namespace: namespace.
type Namespace struct {
	// ID: UUID of the namespace.
	ID string `json:"id"`

	// Name: name of the namespace.
	Name string `json:"name"`

	// EnvironmentVariables: environment variables of the namespace.
	EnvironmentVariables map[string]string `json:"environment_variables"`

	// OrganizationID: UUID of the Organization the namespace belongs to.
	OrganizationID string `json:"organization_id"`

	// ProjectID: UUID of the Project the namespace belongs to.
	ProjectID string `json:"project_id"`

	// Status: status of the namespace.
	// Default value: unknown
	Status NamespaceStatus `json:"status"`

	// RegistryNamespaceID: UUID of the registry namespace.
	RegistryNamespaceID string `json:"registry_namespace_id"`

	// ErrorMessage: error message if the namespace is in "error" state.
	ErrorMessage *string `json:"error_message"`

	// RegistryEndpoint: registry endpoint of the namespace.
	RegistryEndpoint string `json:"registry_endpoint"`

	// Description: description of the namespace.
	Description *string `json:"description"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	// SecretEnvironmentVariables: secret environment variables of the namespace.
	SecretEnvironmentVariables []*SecretHashedValue `json:"secret_environment_variables"`

	Cluster *NamespaceCluster `json:"cluster"`

	// Region: region in which the namespace is located.
	Region scw.Region `json:"region"`
}

// ReconcileTriggerStatusesResponseTriggerStatus: reconcile trigger statuses response trigger status.
type ReconcileTriggerStatusesResponseTriggerStatus struct {
	TriggerID string `json:"trigger_id"`

	// StatusBefore: default value: unknown_status
	StatusBefore ReconcileTriggerStatusesResponseStatus `json:"status_before"`

	ErrorMessageBefore string `json:"error_message_before"`

	ErrorMessageAfter string `json:"error_message_after"`
}

// CreateClusterRequest: create cluster request.
type CreateClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	KapsuleType *string `json:"kapsule_type,omitempty"`

	Version uint32 `json:"version"`
}

// DeleteClusterRequest: delete cluster request.
type DeleteClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ClusterID string `json:"-"`
}

// FunctionDebugData: function debug data.
type FunctionDebugData struct {
	Build *FunctionDebugDataBuild `json:"build"`

	Ksvc *FunctionDebugDataKsvc `json:"ksvc"`

	Pods []*FunctionDebugDataPod `json:"pods"`
}

// GetClusterRequest: get cluster request.
type GetClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ClusterID string `json:"-"`
}

// GetFunctionDebugDataRequest: get function debug data request.
type GetFunctionDebugDataRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`
}

// GetFunctionRequest: get function request.
type GetFunctionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`
}

// GetNamespaceRequest: get namespace request.
type GetNamespaceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NamespaceID: UUID of the namespace.
	NamespaceID string `json:"-"`
}

// ListClustersRequest: list clusters request.
type ListClustersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Name *string `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListClustersRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
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

// ListFunctionsRequest: list functions request.
type ListFunctionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: number of functions per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the functions.
	// Default value: created_at_asc
	OrderBy ListFunctionsRequestOrderBy `json:"-"`

	// NamespaceID: UUID of the namespace the function belongs to.
	NamespaceID string `json:"-"`

	// Name: name of the function.
	Name *string `json:"-"`

	// OrganizationID: UUID of the Organziation the function belongs to.
	OrganizationID *string `json:"-"`

	// ProjectID: UUID of the Project the function belongs to.
	ProjectID *string `json:"-"`
}

// ListFunctionsResponse: list functions response.
type ListFunctionsResponse struct {
	Functions []*Function `json:"functions"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFunctionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFunctionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListFunctionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Functions = append(r.Functions, results.Functions...)
	r.TotalCount += uint32(len(results.Functions))
	return uint32(len(results.Functions)), nil
}

// ListLogsRequest: list logs request.
type ListLogsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FunctionID: UUID of the function to get the logs for.
	FunctionID string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: number of logs per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the logs.
	// Default value: timestamp_desc
	OrderBy ListLogsRequestOrderBy `json:"-"`
}

// ListLogsResponse: list logs response.
type ListLogsResponse struct {
	Logs []*Log `json:"logs"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListLogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Logs = append(r.Logs, results.Logs...)
	r.TotalCount += uint32(len(results.Logs))
	return uint32(len(results.Logs)), nil
}

// ListNamespacesRequest: list namespaces request.
type ListNamespacesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: number of namespaces per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the namespaces.
	// Default value: created_at_asc
	OrderBy ListNamespacesRequestOrderBy `json:"-"`

	// Name: name of the namespace.
	Name *string `json:"-"`

	// OrganizationID: UUID of the Organization the namespace belongs to.
	OrganizationID *string `json:"-"`

	// ProjectID: UUID of the Project the namespace belongs to.
	ProjectID *string `json:"-"`
}

// ListNamespacesResponse: list namespaces response.
type ListNamespacesResponse struct {
	Namespaces []*Namespace `json:"namespaces"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNamespacesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Namespaces = append(r.Namespaces, results.Namespaces...)
	r.TotalCount += uint32(len(results.Namespaces))
	return uint32(len(results.Namespaces)), nil
}

// OrchestratorAPIGetServiceInfoRequest: orchestrator api get service info request.
type OrchestratorAPIGetServiceInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// OrchestratorAPIReconcileTriggerStatusesRequest: orchestrator api reconcile trigger statuses request.
type OrchestratorAPIReconcileTriggerStatusesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	DryRun bool `json:"dry_run"`
}

// OrchestratorAPIRedeployTriggersRequest: orchestrator api redeploy triggers request.
type OrchestratorAPIRedeployTriggersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Precisely one of ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`
}

// ReconcileTriggerStatusesResponse: reconcile trigger statuses response.
type ReconcileTriggerStatusesResponse struct {
	TriggerStatuses []*ReconcileTriggerStatusesResponseTriggerStatus `json:"trigger_statuses"`
}

// RedeployTriggersResponse: redeploy triggers response.
type RedeployTriggersResponse struct {
	TriggerIDs []string `json:"trigger_ids"`
}

// UpdateClusterRequest: update cluster request.
type UpdateClusterRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ClusterID string `json:"-"`

	Billable *bool `json:"billable,omitempty"`

	CanProvision *bool `json:"can_provision,omitempty"`

	Upgrade *bool `json:"upgrade,omitempty"`

	Version *uint32 `json:"version,omitempty"`
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

// ListNamespaces:
func (s *API) ListNamespaces(req *ListNamespacesRequest, opts ...scw.RequestOption) (*ListNamespacesResponse, error) {
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Query:  query,
	}

	var resp ListNamespacesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNamespace:
func (s *API) GetNamespace(req *GetNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFunctions:
func (s *API) ListFunctions(req *ListFunctionsRequest, opts ...scw.RequestOption) (*ListFunctionsResponse, error) {
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
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions",
		Query:  query,
	}

	var resp ListFunctionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFunction:
func (s *API) GetFunction(req *GetFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetFunctionDebugData:
func (s *API) GetFunctionDebugData(req *GetFunctionDebugDataRequest, opts ...scw.RequestOption) (*FunctionDebugData, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/debug-data",
	}

	var resp FunctionDebugData

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListLogs:
func (s *API) ListLogs(req *ListLogsRequest, opts ...scw.RequestOption) (*ListLogsResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/logs",
		Query:  query,
	}

	var resp ListLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/clusters",
		Query:  query,
	}

	var resp ListClustersResponse

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
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
	}

	var resp Cluster

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
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/clusters",
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
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
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

// DeleteCluster:
func (s *API) DeleteCluster(req *DeleteClusterRequest, opts ...scw.RequestOption) (*Cluster, error) {
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
		Method: "DELETE",
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "",
	}

	var resp Cluster

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type OrchestratorAPI struct {
	client *scw.Client
}

// NewOrchestratorAPI returns a OrchestratorAPI object from a Scaleway client.
func NewOrchestratorAPI(client *scw.Client) *OrchestratorAPI {
	return &OrchestratorAPI{
		client: client,
	}
}
func (s *OrchestratorAPI) Regions() []scw.Region {
	return []scw.Region{}
}

// GetServiceInfo:
func (s *OrchestratorAPI) GetServiceInfo(req *OrchestratorAPIGetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
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
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/orchestrator",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RedeployTriggers:
func (s *OrchestratorAPI) RedeployTriggers(req *OrchestratorAPIRedeployTriggersRequest, opts ...scw.RequestOption) (*RedeployTriggersResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/orchestrator/redeploy-triggers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RedeployTriggersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReconcileTriggerStatuses:
func (s *OrchestratorAPI) ReconcileTriggerStatuses(req *OrchestratorAPIReconcileTriggerStatusesRequest, opts ...scw.RequestOption) (*ReconcileTriggerStatusesResponse, error) {
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
		Path:   "/functions-admin/v1beta1/regions/" + fmt.Sprint(req.Region) + "/orchestrator/reconcile-trigger-statuses",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ReconcileTriggerStatusesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package datalab provides methods and message types of the datalab v1beta1 API.
package datalab

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
	defaultDatalabRetryInterval = 15 * time.Second
	defaultDatalabTimeout       = 15 * time.Minute
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

type DatalabStatus string

const (
	DatalabStatusUnknownStatus = DatalabStatus("unknown_status")
	DatalabStatusCreating      = DatalabStatus("creating")
	DatalabStatusUpdating      = DatalabStatus("updating")
	DatalabStatusReady         = DatalabStatus("ready")
	DatalabStatusError         = DatalabStatus("error")
	DatalabStatusDeleting      = DatalabStatus("deleting")
	DatalabStatusLocked        = DatalabStatus("locked")
	DatalabStatusDeleted       = DatalabStatus("deleted")
)

func (enum DatalabStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(DatalabStatusUnknownStatus)
	}
	return string(enum)
}

func (enum DatalabStatus) Values() []DatalabStatus {
	return []DatalabStatus{
		"unknown_status",
		"creating",
		"updating",
		"ready",
		"error",
		"deleting",
		"locked",
		"deleted",
	}
}

func (enum DatalabStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DatalabStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DatalabStatus(DatalabStatus(tmp).String())
	return nil
}

type ListClusterVersionsRequestOrderBy string

const (
	ListClusterVersionsRequestOrderByNameAsc  = ListClusterVersionsRequestOrderBy("name_asc")
	ListClusterVersionsRequestOrderByNameDesc = ListClusterVersionsRequestOrderBy("name_desc")
)

func (enum ListClusterVersionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListClusterVersionsRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListClusterVersionsRequestOrderBy) Values() []ListClusterVersionsRequestOrderBy {
	return []ListClusterVersionsRequestOrderBy{
		"name_asc",
		"name_desc",
	}
}

func (enum ListClusterVersionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListClusterVersionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListClusterVersionsRequestOrderBy(ListClusterVersionsRequestOrderBy(tmp).String())
	return nil
}

type ListDatalabsRequestOrderBy string

const (
	ListDatalabsRequestOrderByNameAsc       = ListDatalabsRequestOrderBy("name_asc")
	ListDatalabsRequestOrderByNameDesc      = ListDatalabsRequestOrderBy("name_desc")
	ListDatalabsRequestOrderByCreatedAtAsc  = ListDatalabsRequestOrderBy("created_at_asc")
	ListDatalabsRequestOrderByCreatedAtDesc = ListDatalabsRequestOrderBy("created_at_desc")
	ListDatalabsRequestOrderByUpdatedAtAsc  = ListDatalabsRequestOrderBy("updated_at_asc")
	ListDatalabsRequestOrderByUpdatedAtDesc = ListDatalabsRequestOrderBy("updated_at_desc")
)

func (enum ListDatalabsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListDatalabsRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListDatalabsRequestOrderBy) Values() []ListDatalabsRequestOrderBy {
	return []ListDatalabsRequestOrderBy{
		"name_asc",
		"name_desc",
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
	}
}

func (enum ListDatalabsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDatalabsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDatalabsRequestOrderBy(ListDatalabsRequestOrderBy(tmp).String())
	return nil
}

type ListNodeTypesRequestOrderBy string

const (
	ListNodeTypesRequestOrderByNameAsc             = ListNodeTypesRequestOrderBy("name_asc")
	ListNodeTypesRequestOrderByNameDesc            = ListNodeTypesRequestOrderBy("name_desc")
	ListNodeTypesRequestOrderByVcpusAsc            = ListNodeTypesRequestOrderBy("vcpus_asc")
	ListNodeTypesRequestOrderByVcpusDesc           = ListNodeTypesRequestOrderBy("vcpus_desc")
	ListNodeTypesRequestOrderByMemoryGigabytesAsc  = ListNodeTypesRequestOrderBy("memory_gigabytes_asc")
	ListNodeTypesRequestOrderByMemoryGigabytesDesc = ListNodeTypesRequestOrderBy("memory_gigabytes_desc")
	ListNodeTypesRequestOrderByVramBytesAsc        = ListNodeTypesRequestOrderBy("vram_bytes_asc")
	ListNodeTypesRequestOrderByVramBytesDesc       = ListNodeTypesRequestOrderBy("vram_bytes_desc")
	ListNodeTypesRequestOrderByGpusAsc             = ListNodeTypesRequestOrderBy("gpus_asc")
	ListNodeTypesRequestOrderByGpusDesc            = ListNodeTypesRequestOrderBy("gpus_desc")
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
		"memory_gigabytes_asc",
		"memory_gigabytes_desc",
		"vram_bytes_asc",
		"vram_bytes_desc",
		"gpus_asc",
		"gpus_desc",
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

type ListNodeTypesRequestResourceType string

const (
	ListNodeTypesRequestResourceTypeAll = ListNodeTypesRequestResourceType("all")
	ListNodeTypesRequestResourceTypeGpu = ListNodeTypesRequestResourceType("gpu")
	ListNodeTypesRequestResourceTypeCPU = ListNodeTypesRequestResourceType("cpu")
)

func (enum ListNodeTypesRequestResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListNodeTypesRequestResourceTypeAll)
	}
	return string(enum)
}

func (enum ListNodeTypesRequestResourceType) Values() []ListNodeTypesRequestResourceType {
	return []ListNodeTypesRequestResourceType{
		"all",
		"gpu",
		"cpu",
	}
}

func (enum ListNodeTypesRequestResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNodeTypesRequestResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNodeTypesRequestResourceType(ListNodeTypesRequestResourceType(tmp).String())
	return nil
}

type ListNotebookVersionsRequestOrderBy string

const (
	ListNotebookVersionsRequestOrderByNameAsc  = ListNotebookVersionsRequestOrderBy("name_asc")
	ListNotebookVersionsRequestOrderByNameDesc = ListNotebookVersionsRequestOrderBy("name_desc")
)

func (enum ListNotebookVersionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListNotebookVersionsRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListNotebookVersionsRequestOrderBy) Values() []ListNotebookVersionsRequestOrderBy {
	return []ListNotebookVersionsRequestOrderBy{
		"name_asc",
		"name_desc",
	}
}

func (enum ListNotebookVersionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNotebookVersionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNotebookVersionsRequestOrderBy(ListNotebookVersionsRequestOrderBy(tmp).String())
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
		return string(NodeTypeStockUnknownStock)
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

type NodeTypeTarget string

const (
	NodeTypeTargetUnknownTarget = NodeTypeTarget("unknown_target")
	NodeTypeTargetNotebook      = NodeTypeTarget("notebook")
	NodeTypeTargetWorker        = NodeTypeTarget("worker")
)

func (enum NodeTypeTarget) String() string {
	if enum == "" {
		// return default value if empty
		return string(NodeTypeTargetUnknownTarget)
	}
	return string(enum)
}

func (enum NodeTypeTarget) Values() []NodeTypeTarget {
	return []NodeTypeTarget{
		"unknown_target",
		"notebook",
		"worker",
	}
}

func (enum NodeTypeTarget) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NodeTypeTarget) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NodeTypeTarget(NodeTypeTarget(tmp).String())
	return nil
}

type VolumeType string

const (
	VolumeTypeUnknownType = VolumeType("unknown_type")
	VolumeTypeSbs5k       = VolumeType("sbs_5k")
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

// Volume: volume.
type Volume struct {
	// Type: default value: unknown_type
	Type VolumeType `json:"type"`

	Size scw.Size `json:"size"`
}

// ClusterVersion: A cluster version.
type ClusterVersion struct {
	// Version: the version of the cluster.
	Version string `json:"version"`

	// EndOfLife: the end of life date of the cluster version.
	EndOfLife *time.Time `json:"end_of_life"`

	// CreatedAt: the creation date of the cluster version.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: the last update date of the cluster version.
	UpdatedAt *time.Time `json:"updated_at"`

	// Disabled: whether the cluster version is disabled.
	Disabled bool `json:"disabled"`

	// Beta: whether the cluster version is in beta.
	Beta bool `json:"beta"`
}

// DatalabSparkMain: datalab spark main.
type DatalabSparkMain struct {
	NodeType string `json:"node_type"`

	SparkUIURL string `json:"spark_ui_url"`

	SparkMasterURL string `json:"spark_master_url"`

	RootVolume *Volume `json:"root_volume"`
}

// DatalabSparkWorker: datalab spark worker.
type DatalabSparkWorker struct {
	NodeType string `json:"node_type"`

	NodeCount uint32 `json:"node_count"`

	RootVolume *Volume `json:"root_volume"`
}

// NotebookVersion: A notebook version.
type NotebookVersion struct {
	// Version: the version of the notebook.
	Version string `json:"version"`

	// EndOfLife: the end of life date of the notebook version.
	EndOfLife *time.Time `json:"end_of_life"`

	// CreatedAt: the creation date of the notebook version.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: the last update date of the notebook version.
	UpdatedAt *time.Time `json:"updated_at"`

	// Disabled: whether the notebook version is disabled.
	Disabled bool `json:"disabled"`

	// Beta: whether the notebook version is in beta.
	Beta bool `json:"beta"`
}

// CreateDatalabRequestSparkMain: create datalab request spark main.
type CreateDatalabRequestSparkMain struct {
	NodeType string `json:"node_type"`
}

// CreateDatalabRequestSparkWorker: create datalab request spark worker.
type CreateDatalabRequestSparkWorker struct {
	NodeType string `json:"node_type"`

	NodeCount uint32 `json:"node_count"`
}

// Cluster: A cluster.
type Cluster struct {
	// Name: the name of the cluster.
	Name string `json:"name"`

	// Description: the description of the cluster.
	Description string `json:"description"`

	// Versions: the versions of the cluster.
	Versions []*ClusterVersion `json:"versions"`
}

// Datalab: A Data Lab resource.
type Datalab struct {
	// ID: the unique identifier of the Data Lab.
	ID string `json:"id"`

	// ProjectID: the identifier of the project where the Data Lab has been created.
	ProjectID string `json:"project_id"`

	// Name: the name of the Data Lab.
	Name string `json:"name"`

	// Description: the description of the Data Lab.
	Description string `json:"description"`

	// Tags: the tags of the Data Lab.
	Tags []string `json:"tags"`

	// Main: the Spark Main node specification of Data lab. It holds the parameters `node_type` the compute node type of the main node, `spark_ui_url` where the Spark UI is available, `spark_master_url` with which one can connect to the cluster from within one's VPC, `root_volume` the size of the volume assigned to the main node.
	Main *DatalabSparkMain `json:"main"`

	// Worker: the worker node specification of the Data Lab. It presents the parameters `node_type` the compute node type of each worker node, `node_count` the number of worker nodes currently in the cluster, `root_volume` the root volume size of each executor.
	Worker *DatalabSparkWorker `json:"worker"`

	// Status: the status of the Data Lab. For a working Data Lab this should be `ready`.
	// Default value: unknown_status
	Status DatalabStatus `json:"status"`

	// CreatedAt: the creation timestamp of the Data Lab.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: the last update date of the Data Lab.
	UpdatedAt *time.Time `json:"updated_at"`

	// Region: the region of the Data Lab.
	Region scw.Region `json:"region"`

	// HasNotebook: whether a JupyterLab notebook is associated with the Data Lab or not.
	HasNotebook bool `json:"has_notebook"`

	// NotebookURL: the URL of said notebook if exists.
	NotebookURL *string `json:"notebook_url"`

	// SparkVersion: the version of Spark running inside the Data Lab.
	SparkVersion string `json:"spark_version"`

	// TotalStorage: the total storage selected by the user for Spark.
	TotalStorage *Volume `json:"total_storage"`

	// PrivateNetworkID: the private network to which the data lab is connected. This is important for accessing the Spark Master URL.
	PrivateNetworkID string `json:"private_network_id"`

	// NotebookMasterURL: the URL to the Spark Master endpoint from, and only from the perspective of the JupyterLab Notebook. This is NOT the URL to use for accessing the cluster from a private server.
	NotebookMasterURL *string `json:"notebook_master_url"`
}

// NodeType: A node type.
type NodeType struct {
	// StockStatus: the stock status of the node type.
	// Default value: unknown_stock
	StockStatus NodeTypeStock `json:"stock_status"`

	// Name: the name of the node type.
	Name string `json:"name"`

	// Description: the description of the node type.
	Description string `json:"description"`

	// Vcpus: the number of vCPUs.
	Vcpus uint32 `json:"vcpus"`

	// MemoryGigabytes: the amount of memory.
	MemoryGigabytes uint64 `json:"memory_gigabytes"`

	// VramGigabytes: the amount of VRAM.
	VramGigabytes uint64 `json:"vram_gigabytes"`

	// Gpus: the number of GPUs.
	Gpus uint32 `json:"gpus"`

	// Disabled: whether the node type is disabled.
	Disabled bool `json:"disabled"`

	// Beta: whether the node type is in beta.
	Beta bool `json:"beta"`

	// CreatedAt: the creation date of the node type.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: the last update date of the node type.
	UpdatedAt *time.Time `json:"updated_at"`

	// Targets: the targets of the node type.
	Targets []NodeTypeTarget `json:"targets"`
}

// Notebook: A notebook.
type Notebook struct {
	// Name: the name of the notebook.
	Name string `json:"name"`

	// Description: the description of the notebook.
	Description string `json:"description"`

	// Versions: the versions of the notebook.
	Versions []*NotebookVersion `json:"versions"`
}

// CreateDatalabRequest: A request to create a Data Lab.
type CreateDatalabRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: the unique identifier of the project where the Data Lab will be created.
	ProjectID string `json:"project_id"`

	// Name: the name of the Data Lab.
	Name string `json:"name"`

	// Description: the description of the Data Lab.
	Description string `json:"description"`

	// Tags: the tags of the Data Lab.
	Tags []string `json:"tags"`

	// Main: the Spark main node configuration of the Data Lab, has one parameter `node_type` which specifies the compute node type of the main node. See ListNodeTypes for available options.
	Main *CreateDatalabRequestSparkMain `json:"main,omitempty"`

	// Worker: the Spark worker node configuration of the Data Lab, has two parameters `node_type` for selecting the type of the worker node, and `node_count` for specifying the amount of nodes.
	Worker *CreateDatalabRequestSparkWorker `json:"worker,omitempty"`

	// HasNotebook: whether a JupyterLab notebook shall be created with the Data Lab or not.
	HasNotebook bool `json:"has_notebook"`

	// SparkVersion: the version of Spark running inside the Data Lab, available options can be viewed at ListClusterVersions.
	SparkVersion string `json:"spark_version"`

	// TotalStorage: the total storage selected by the user for Spark workers. This means the workers will not use more then this amount for their workload.
	TotalStorage *Volume `json:"total_storage,omitempty"`

	// PrivateNetworkID: the private network to which the Data Lab is connected. Important for accessing the Spark Master URL from a private cluster.
	PrivateNetworkID string `json:"private_network_id"`
}

// DeleteDatalabRequest: A request to delete a Data Lab.
type DeleteDatalabRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DatalabID: the unique identifier of the Data Lab.
	DatalabID string `json:"-"`
}

// GetDatalabRequest: A request to get information about a Data Lab.
type GetDatalabRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DatalabID: the unique identifier of the Data Lab.
	DatalabID string `json:"-"`
}

// ListClusterVersionsRequest: A request to list cluster versions.
type ListClusterVersionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: the page number.
	Page *int32 `json:"-"`

	// PageSize: the page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: the order by field.
	// Default value: name_asc
	OrderBy ListClusterVersionsRequestOrderBy `json:"-"`
}

// ListClusterVersionsResponse: A response to list cluster versions.
type ListClusterVersionsResponse struct {
	// Clusters: the list of cluster versions.
	Clusters []*Cluster `json:"clusters"`

	// TotalCount: the total count of cluster versions.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListClusterVersionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListClusterVersionsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListClusterVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Clusters = append(r.Clusters, results.Clusters...)
	r.TotalCount += uint64(len(results.Clusters))
	return uint64(len(results.Clusters)), nil
}

// ListDatalabsRequest: A request to list Datalabs.
type ListDatalabsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrganizationID: the unique identifier of the organization whose Data Labs you want to list.
	OrganizationID *string `json:"-"`

	// ProjectID: the unique identifier of the project whose Data Labs you want to list.
	ProjectID *string `json:"-"`

	// Name: the name of the Data Lab you want to list.
	Name *string `json:"-"`

	// Tags: the tags associated with the Data Lab you want to list.
	Tags []string `json:"-"`

	// Page: the page number for pagination.
	Page *int32 `json:"-"`

	// PageSize: the page size for pagination.
	PageSize *uint32 `json:"-"`

	// OrderBy: the order by field, available options are `name_asc`, `name_desc`, `created_at_asc`, `created_at_desc`, `updated_at_asc`, `updated_at_desc`.
	// Default value: name_asc
	OrderBy ListDatalabsRequestOrderBy `json:"-"`
}

// ListDatalabsResponse: A response to list Data Labs.
type ListDatalabsResponse struct {
	// Datalabs: the list of Data Labs. This is a list composed of messages of type `DataLab`.
	Datalabs []*Datalab `json:"datalabs"`

	// TotalCount: the total count of Datalabs.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatalabsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatalabsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListDatalabsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Datalabs = append(r.Datalabs, results.Datalabs...)
	r.TotalCount += uint64(len(results.Datalabs))
	return uint64(len(results.Datalabs)), nil
}

// ListNodeTypesRequest: A request to list node types.
type ListNodeTypesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: the page number.
	Page *int32 `json:"-"`

	// PageSize: the page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: the order by field. Available fields are `name_asc`, `name_desc`, `vcpus_asc`, `vcpus_desc`, `memory_gigabytes_asc`, `memory_gigabytes_desc`, `vram_bytes_asc`, `vram_bytes_desc`, `gpus_asc`, `gpus_desc`.
	// Default value: name_asc
	OrderBy ListNodeTypesRequestOrderBy `json:"-"`

	// Targets: filter on the wanted targets, whether it's for main node or worker.
	Targets []NodeTypeTarget `json:"-"`

	// ResourceType: filter based on node type ( `cpu`/`gpu`/`all` ).
	// Default value: all
	ResourceType ListNodeTypesRequestResourceType `json:"-"`
}

// ListNodeTypesResponse: A response to list node types.
type ListNodeTypesResponse struct {
	// NodeTypes: the list of node types.
	NodeTypes []*NodeType `json:"node_types"`

	// TotalCount: the total count of node types.
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

// ListNotebookVersionsRequest: A request to list notebook versions.
type ListNotebookVersionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: the page number.
	Page *int32 `json:"-"`

	// PageSize: the page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: the order by field. Available options are `name_asc` and `name_desc`.
	// Default value: name_asc
	OrderBy ListNotebookVersionsRequestOrderBy `json:"-"`
}

// ListNotebookVersionsResponse: A response to list notebook versions.
type ListNotebookVersionsResponse struct {
	// Notebooks: the list of notebook versions.
	Notebooks []*Notebook `json:"notebooks"`

	// TotalCount: the total count of notebook versions.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNotebookVersionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNotebookVersionsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListNotebookVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Notebooks = append(r.Notebooks, results.Notebooks...)
	r.TotalCount += uint64(len(results.Notebooks))
	return uint64(len(results.Notebooks)), nil
}

// UpdateDatalabRequest: A request to update a Data Lab.
type UpdateDatalabRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DatalabID: the unique identifier of the Data Lab.
	DatalabID string `json:"-"`

	// Name: the updated name of the Data Lab.
	Name *string `json:"name,omitempty"`

	// Description: the updated description of the Data Lab.
	Description *string `json:"description,omitempty"`

	// Tags: the updated tags of the Data Lab.
	Tags []string `json:"tags"`

	// NodeCount: the updated node count of the Data Lab. Scale up or down the number of worker nodes.
	NodeCount *uint32 `json:"node_count,omitempty"`
}

// This API allows you to manage your Data Lab resources.
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

// CreateDatalab: Create a new Data Lab. In this call, one can personalize the node counts, add a notebook, choose the private network, define the persistent volume storage capacity.
func (s *API) CreateDatalab(req *CreateDatalabRequest, opts ...scw.RequestOption) (*Datalab, error) {
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
		Path:   "/datalab/v1beta1/regions/" + fmt.Sprint(req.Region) + "/datalabs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Datalab

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDatalab: Retrieve information about a given Data Lab cluster, specified by the `region` and `datalab_id` parameters. Its full details, including name, status, node counts, are returned in the response object.
func (s *API) GetDatalab(req *GetDatalabRequest, opts ...scw.RequestOption) (*Datalab, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatalabID) == "" {
		return nil, errors.New("field DatalabID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/datalab/v1beta1/regions/" + fmt.Sprint(req.Region) + "/datalabs/" + fmt.Sprint(req.DatalabID) + "",
	}

	var resp Datalab

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForDatalabRequest is used by WaitForDatalab method.
type WaitForDatalabRequest struct {
	Region        scw.Region
	DatalabID     string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForDatalab waits for the Datalab to reach a terminal state.
func (s *API) WaitForDatalab(req *WaitForDatalabRequest, opts ...scw.RequestOption) (*Datalab, error) {
	timeout := defaultDatalabTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultDatalabRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[DatalabStatus]struct{}{
		DatalabStatusCreating: {},
		DatalabStatusUpdating: {},
		DatalabStatusDeleting: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetDatalab(&GetDatalabRequest{
				Region:    req.Region,
				DatalabID: req.DatalabID,
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
		return nil, errors.Wrap(err, "waiting for Datalab failed")
	}

	return res.(*Datalab), nil
}

// ListDatalabs: List information about Data Lab cluster within a project or an organization.
func (s *API) ListDatalabs(req *ListDatalabsRequest, opts ...scw.RequestOption) (*ListDatalabsResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/datalab/v1beta1/regions/" + fmt.Sprint(req.Region) + "/datalabs",
		Query:  query,
	}

	var resp ListDatalabsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDatalab: Update a Data Labs node counts. Allows for up- and downscaling on demand, depending on the expected workload.
func (s *API) UpdateDatalab(req *UpdateDatalabRequest, opts ...scw.RequestOption) (*Datalab, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatalabID) == "" {
		return nil, errors.New("field DatalabID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/datalab/v1beta1/regions/" + fmt.Sprint(req.Region) + "/datalabs/" + fmt.Sprint(req.DatalabID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Datalab

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDatalab: Delete a Data Lab based on its region and id.
func (s *API) DeleteDatalab(req *DeleteDatalabRequest, opts ...scw.RequestOption) (*Datalab, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatalabID) == "" {
		return nil, errors.New("field DatalabID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/datalab/v1beta1/regions/" + fmt.Sprint(req.Region) + "/datalabs/" + fmt.Sprint(req.DatalabID) + "",
	}

	var resp Datalab

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNodeTypes: List the available compute node types upon which a Data Lab can be created.
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "targets", req.Targets)
	parameter.AddToQuery(query, "resource_type", req.ResourceType)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/datalab/v1beta1/regions/" + fmt.Sprint(req.Region) + "/node-types",
		Query:  query,
	}

	var resp ListNodeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNotebookVersions: List available notebook versions.
func (s *API) ListNotebookVersions(req *ListNotebookVersionsRequest, opts ...scw.RequestOption) (*ListNotebookVersionsResponse, error) {
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

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/datalab/v1beta1/regions/" + fmt.Sprint(req.Region) + "/notebook-versions",
		Query:  query,
	}

	var resp ListNotebookVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusterVersions: List the Spark versions the product is compatible with.
func (s *API) ListClusterVersions(req *ListClusterVersionsRequest, opts ...scw.RequestOption) (*ListClusterVersionsResponse, error) {
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

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/datalab/v1beta1/regions/" + fmt.Sprint(req.Region) + "/cluster-versions",
		Query:  query,
	}

	var resp ListClusterVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package search provides methods and message types of the search v1alpha1 API.
package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

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

type Locality string

const (
	LocalityUnknownLocality = Locality("unknown_locality")
	LocalityGlobal          = Locality("global")
	LocalityFrRz            = Locality("fr_rz")
	LocalityFrSrr           = Locality("fr_srr")
	LocalityFrSrr1          = Locality("fr_srr_1")
	LocalityFrPar           = Locality("fr_par")
	LocalityFrPar1          = Locality("fr_par_1")
	LocalityFrPar2          = Locality("fr_par_2")
	LocalityFrPar3          = Locality("fr_par_3")
	LocalityFrPar4          = Locality("fr_par_4")
	LocalityNlAms           = Locality("nl_ams")
	LocalityNlAms1          = Locality("nl_ams_1")
	LocalityNlAms2          = Locality("nl_ams_2")
	LocalityNlAms3          = Locality("nl_ams_3")
	LocalityPlWaw           = Locality("pl_waw")
	LocalityPlWaw1          = Locality("pl_waw_1")
	LocalityPlWaw2          = Locality("pl_waw_2")
	LocalityPlWaw3          = Locality("pl_waw_3")
	LocalityFrInt           = Locality("fr_int")
	LocalityFrInt1          = Locality("fr_int_1")
	LocalityFrLab           = Locality("fr_lab")
	LocalityFrLab1          = Locality("fr_lab_1")
	LocalityItMil           = Locality("it_mil")
	LocalityItMil1          = Locality("it_mil_1")
)

func (enum Locality) String() string {
	if enum == "" {
		// return default value if empty
		return string(LocalityUnknownLocality)
	}
	return string(enum)
}

func (enum Locality) Values() []Locality {
	return []Locality{
		"unknown_locality",
		"global",
		"fr_rz",
		"fr_srr",
		"fr_srr_1",
		"fr_par",
		"fr_par_1",
		"fr_par_2",
		"fr_par_3",
		"fr_par_4",
		"nl_ams",
		"nl_ams_1",
		"nl_ams_2",
		"nl_ams_3",
		"pl_waw",
		"pl_waw_1",
		"pl_waw_2",
		"pl_waw_3",
		"fr_int",
		"fr_int_1",
		"fr_lab",
		"fr_lab_1",
		"it_mil",
		"it_mil_1",
	}
}

func (enum Locality) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Locality) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Locality(Locality(tmp).String())
	return nil
}

type ObsDatasourceInfoDataType string

const (
	ObsDatasourceInfoDataTypeUnknownDataType = ObsDatasourceInfoDataType("unknown_data_type")
	ObsDatasourceInfoDataTypeMetrics         = ObsDatasourceInfoDataType("metrics")
	ObsDatasourceInfoDataTypeLogs            = ObsDatasourceInfoDataType("logs")
	ObsDatasourceInfoDataTypeTraces          = ObsDatasourceInfoDataType("traces")
)

func (enum ObsDatasourceInfoDataType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ObsDatasourceInfoDataTypeUnknownDataType)
	}
	return string(enum)
}

func (enum ObsDatasourceInfoDataType) Values() []ObsDatasourceInfoDataType {
	return []ObsDatasourceInfoDataType{
		"unknown_data_type",
		"metrics",
		"logs",
		"traces",
	}
}

func (enum ObsDatasourceInfoDataType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ObsDatasourceInfoDataType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ObsDatasourceInfoDataType(ObsDatasourceInfoDataType(tmp).String())
	return nil
}

type ObsExporterInfoDestinationType string

const (
	ObsExporterInfoDestinationTypeUnknownDestinationType = ObsExporterInfoDestinationType("unknown_destination_type")
	ObsExporterInfoDestinationTypeDatadog                = ObsExporterInfoDestinationType("datadog")
	ObsExporterInfoDestinationTypeOtlp                   = ObsExporterInfoDestinationType("otlp")
)

func (enum ObsExporterInfoDestinationType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ObsExporterInfoDestinationTypeUnknownDestinationType)
	}
	return string(enum)
}

func (enum ObsExporterInfoDestinationType) Values() []ObsExporterInfoDestinationType {
	return []ObsExporterInfoDestinationType{
		"unknown_destination_type",
		"datadog",
		"otlp",
	}
}

func (enum ObsExporterInfoDestinationType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ObsExporterInfoDestinationType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ObsExporterInfoDestinationType(ObsExporterInfoDestinationType(tmp).String())
	return nil
}

type ResourceType string

const (
	// Unknown type.
	ResourceTypeUnknownType = ResourceType("unknown_type")
	// Instance server.
	ResourceTypeInstanceServer = ResourceType("instance_server")
	// Instance volume.
	ResourceTypeInstanceVolume = ResourceType("instance_volume")
	// Instance image.
	ResourceTypeInstanceImage          = ResourceType("instance_image")
	ResourceTypeInstanceSecurityGroup  = ResourceType("instance_security_group")
	ResourceTypeInstancePrivateNic     = ResourceType("instance_private_nic")
	ResourceTypeInstanceSnapshot       = ResourceType("instance_snapshot")
	ResourceTypeInstancePlacementGroup = ResourceType("instance_placement_group")
	// K8S cluster.
	ResourceTypeK8sCluster = ResourceType("k8s_cluster")
	// K8S pool.
	ResourceTypeK8sPool = ResourceType("k8s_pool")
	// K8S node.
	ResourceTypeK8sNode = ResourceType("k8s_node")
	// Domain.
	ResourceTypeDomainDomain = ResourceType("domain_domain")
	// DNS zone.
	ResourceTypeDNSZone = ResourceType("dns_zone")
	// VPC private network.
	ResourceTypeVpcPrivateNetwork = ResourceType("vpc_private_network")
	// VPC.
	ResourceTypeVpcVpc = ResourceType("vpc_vpc")
	// VPG.
	ResourceTypeVpgGateway = ResourceType("vpg_gateway")
	// Apple Silicon server.
	ResourceTypeAppleSiliconServer = ResourceType("apple_silicon_server")
	// RDB instance.
	ResourceTypeRdbInstance = ResourceType("rdb_instance")
	// RDB snapshot.
	ResourceTypeRdbSnapshot = ResourceType("rdb_snapshot")
	// RDB backup.
	ResourceTypeRdbBackup = ResourceType("rdb_backup")
	// Baremetal server.
	ResourceTypeBaremetalServer = ResourceType("baremetal_server")
	ResourceTypeTemDomain       = ResourceType("tem_domain")
	// LB server.
	ResourceTypeLBServer = ResourceType("lb_server")
	// Serverless functions function.
	ResourceTypeServerlessFunctionsFunction = ResourceType("serverless_functions_function")
	// Serverless containers container.
	ResourceTypeServerlessContainersContainer = ResourceType("serverless_containers_container")
	// Web Hosting.
	ResourceTypeWbhHosting = ResourceType("wbh_hosting")
	// Redis cluster.
	ResourceTypeRedisCluster = ResourceType("redis_cluster")
	// Secret manager secret.
	ResourceTypeSmSecret = ResourceType("sm_secret")
	// Key manager key.
	ResourceTypeKmsKey      = ResourceType("kms_key")
	ResourceTypeEdgPipeline = ResourceType("edg_pipeline")
	// Messaging and Queuing NATS Account.
	ResourceTypeMnqNatsAccount = ResourceType("mnq_nats_account")
	// Block storage volume.
	ResourceTypeSbsVolume = ResourceType("sbs_volume")
	// Block storage snapshot.
	ResourceTypeSbsSnapshot             = ResourceType("sbs_snapshot")
	ResourceTypeServerlessJobDefinition = ResourceType("serverless_job_definition")
	ResourceTypeServerlessSqldbDatabase = ResourceType("serverless_sqldb_database")
	ResourceTypeServerlessSqldbBackup   = ResourceType("serverless_sqldb_backup")
	// Distributed Data Lab.
	ResourceTypeDdlDatalab   = ResourceType("ddl_datalab")
	ResourceTypeMgdbInstance = ResourceType("mgdb_instance")
	ResourceTypeMgdbSnapshot = ResourceType("mgdb_snapshot")
	// Managed inference deployment.
	ResourceTypeIfrDeployment = ResourceType("ifr_deployment")
	// Managed inference model.
	ResourceTypeIfrModel  = ResourceType("ifr_model")
	ResourceTypeGapiBatch = ResourceType("gapi_batch")
	// DataWarehouse for Clickhouse deployment.
	ResourceTypeDtwhDeployment      = ResourceType("dtwh_deployment")
	ResourceTypeObsDatasource       = ResourceType("obs_datasource")
	ResourceTypeObsExporter         = ResourceType("obs_exporter")
	ResourceTypeSvpnVpnGateway      = ResourceType("svpn_vpn_gateway")
	ResourceTypeSvpnCustomerGateway = ResourceType("svpn_customer_gateway")
	ResourceTypeSvpnConnection      = ResourceType("svpn_connection")
	ResourceTypeSvpnRoutingPolicy   = ResourceType("svpn_routing_policy")
)

func (enum ResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ResourceTypeUnknownType)
	}
	return string(enum)
}

func (enum ResourceType) Values() []ResourceType {
	return []ResourceType{
		"unknown_type",
		"instance_server",
		"instance_volume",
		"instance_image",
		"instance_security_group",
		"instance_private_nic",
		"instance_snapshot",
		"instance_placement_group",
		"k8s_cluster",
		"k8s_pool",
		"k8s_node",
		"domain_domain",
		"dns_zone",
		"vpc_private_network",
		"vpc_vpc",
		"vpg_gateway",
		"apple_silicon_server",
		"rdb_instance",
		"rdb_snapshot",
		"rdb_backup",
		"baremetal_server",
		"tem_domain",
		"lb_server",
		"serverless_functions_function",
		"serverless_containers_container",
		"wbh_hosting",
		"redis_cluster",
		"sm_secret",
		"kms_key",
		"edg_pipeline",
		"mnq_nats_account",
		"sbs_volume",
		"sbs_snapshot",
		"serverless_job_definition",
		"serverless_sqldb_database",
		"serverless_sqldb_backup",
		"ddl_datalab",
		"mgdb_instance",
		"mgdb_snapshot",
		"ifr_deployment",
		"ifr_model",
		"gapi_batch",
		"dtwh_deployment",
		"obs_datasource",
		"obs_exporter",
		"svpn_vpn_gateway",
		"svpn_customer_gateway",
		"svpn_connection",
		"svpn_routing_policy",
	}
}

func (enum ResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ResourceType(ResourceType(tmp).String())
	return nil
}

// BrmServerInfo: brm server info.
type BrmServerInfo struct {
	IP string `json:"ip"`
}

// ObsDatasourceInfo: obs datasource info.
type ObsDatasourceInfo struct {
	// Type: default value: unknown_data_type
	Type ObsDatasourceInfoDataType `json:"type"`
}

// ObsExporterInfo: obs exporter info.
type ObsExporterInfo struct {
	// DestinationType: default value: unknown_destination_type
	DestinationType ObsExporterInfoDestinationType `json:"destination_type"`
}

// ServerlessContainersContainerInfo: serverless containers container info.
type ServerlessContainersContainerInfo struct {
	// NamespaceID: ID of the Namespace the Serverless Container belongs to.
	NamespaceID string `json:"namespace_id"`
}

// ServerlessFunctionsFunctionInfo: serverless functions function info.
type ServerlessFunctionsFunctionInfo struct {
	// NamespaceID: ID of the Namespace the Serverless Function belongs to.
	NamespaceID string `json:"namespace_id"`
}

// ServerlessSqldbBackupInfo: serverless sqldb backup info.
type ServerlessSqldbBackupInfo struct {
	DatabaseID string `json:"database_id"`
}

// VpcPrivateNetworkInfo: vpc private network info.
type VpcPrivateNetworkInfo struct {
	// VpcID: ID of the VPC the Private Network belongs to.
	VpcID string `json:"vpc_id"`
}

// Resource: resource.
type Resource struct {
	// ID: ID of the resource.
	ID string `json:"id"`

	// Name: name of the resource.
	Name *string `json:"name"`

	// ProjectID: ID of the Project the resource belongs to.
	ProjectID *string `json:"project_id"`

	// OrganizationID: ID of the Organization the resource belongs to.
	OrganizationID string `json:"organization_id"`

	// Type: type of the resource.
	// Default value: unknown_type
	Type ResourceType `json:"type"`

	// Global: locality the resource is in if this is a global resource.
	// Precisely one of Global, Zone, Region must be set.
	Global *bool `json:"global,omitempty"`

	// Zone: locality the resource is in if this is a zonal resource.
	// Precisely one of Global, Zone, Region must be set.
	Zone *scw.Zone `json:"zone,omitempty"`

	// Region: locality the resource is in if this is a regional resource.
	// Precisely one of Global, Zone, Region must be set.
	Region *scw.Region `json:"region,omitempty"`

	// VpcPrivateNetworkInfo: additional information for a VPC Private Network.
	// Precisely one of VpcPrivateNetworkInfo, ServerlessFunctionsFunctionInfo, ServerlessContainersContainerInfo, BaremetalServerInfo, ServerlessSqldbBackupInfo, ObsDatasourceInfo, ObsExporterInfo must be set.
	VpcPrivateNetworkInfo *VpcPrivateNetworkInfo `json:"vpc_private_network_info,omitempty"`

	// ServerlessFunctionsFunctionInfo: additional information for a Serverless Function.
	// Precisely one of VpcPrivateNetworkInfo, ServerlessFunctionsFunctionInfo, ServerlessContainersContainerInfo, BaremetalServerInfo, ServerlessSqldbBackupInfo, ObsDatasourceInfo, ObsExporterInfo must be set.
	ServerlessFunctionsFunctionInfo *ServerlessFunctionsFunctionInfo `json:"serverless_functions_function_info,omitempty"`

	// ServerlessContainersContainerInfo: additional information for a Serverless Container.
	// Precisely one of VpcPrivateNetworkInfo, ServerlessFunctionsFunctionInfo, ServerlessContainersContainerInfo, BaremetalServerInfo, ServerlessSqldbBackupInfo, ObsDatasourceInfo, ObsExporterInfo must be set.
	ServerlessContainersContainerInfo *ServerlessContainersContainerInfo `json:"serverless_containers_container_info,omitempty"`

	// Precisely one of VpcPrivateNetworkInfo, ServerlessFunctionsFunctionInfo, ServerlessContainersContainerInfo, BaremetalServerInfo, ServerlessSqldbBackupInfo, ObsDatasourceInfo, ObsExporterInfo must be set.
	BaremetalServerInfo *BrmServerInfo `json:"baremetal_server_info,omitempty"`

	// Precisely one of VpcPrivateNetworkInfo, ServerlessFunctionsFunctionInfo, ServerlessContainersContainerInfo, BaremetalServerInfo, ServerlessSqldbBackupInfo, ObsDatasourceInfo, ObsExporterInfo must be set.
	ServerlessSqldbBackupInfo *ServerlessSqldbBackupInfo `json:"serverless_sqldb_backup_info,omitempty"`

	// Precisely one of VpcPrivateNetworkInfo, ServerlessFunctionsFunctionInfo, ServerlessContainersContainerInfo, BaremetalServerInfo, ServerlessSqldbBackupInfo, ObsDatasourceInfo, ObsExporterInfo must be set.
	ObsDatasourceInfo *ObsDatasourceInfo `json:"obs_datasource_info,omitempty"`

	// Precisely one of VpcPrivateNetworkInfo, ServerlessFunctionsFunctionInfo, ServerlessContainersContainerInfo, BaremetalServerInfo, ServerlessSqldbBackupInfo, ObsDatasourceInfo, ObsExporterInfo must be set.
	ObsExporterInfo *ObsExporterInfo `json:"obs_exporter_info,omitempty"`
}

// SearchResourcesRequest: search resources request.
type SearchResourcesRequest struct {
	// Query: search query.
	Query string `json:"-"`

	// OrganizationID: ID of the Organization to search in.
	OrganizationID string `json:"-"`

	// ProjectIDs: list of Project IDs to filter the resources by.
	ProjectIDs []string `json:"-"`

	// Types: list of resource types to filter the resources by.
	Types []ResourceType `json:"-"`

	// Localities: list of scopes (zones, regions, or global) to filter the resources by.
	Localities []Locality `json:"-"`

	// CreatedAfter: filter resources created after this timestamp.
	CreatedAfter *time.Time `json:"-"`

	// CreatedBefore: filter resources created before this timestamp.
	CreatedBefore *time.Time `json:"-"`

	// ModifiedAfter: filter resources modified after this timestamp.
	ModifiedAfter *time.Time `json:"-"`

	// ModifiedBefore: filter resources modified before this timestamp.
	ModifiedBefore *time.Time `json:"-"`
}

// SearchResourcesResponse: search resources response.
type SearchResourcesResponse struct {
	// Resources: top resources found.
	Resources []*Resource `json:"resources"`
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

// SearchResources: Search API.
func (s *API) SearchResources(req *SearchResourcesRequest, opts ...scw.RequestOption) (*SearchResourcesResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "query", req.Query)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_ids", req.ProjectIDs)
	parameter.AddToQuery(query, "types", req.Types)
	parameter.AddToQuery(query, "localities", req.Localities)
	parameter.AddToQuery(query, "created_after", req.CreatedAfter)
	parameter.AddToQuery(query, "created_before", req.CreatedBefore)
	parameter.AddToQuery(query, "modified_after", req.ModifiedAfter)
	parameter.AddToQuery(query, "modified_before", req.ModifiedBefore)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/search/v1alpha1/resources",
		Query:  query,
	}

	var resp SearchResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package audit_trail provides methods and message types of the audit_trail v1alpha1 API.
package audit_trail

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	std "github.com/scaleway/scaleway-sdk-go/api/std"
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

type AuthenticationEventFailureReason string

const (
	AuthenticationEventFailureReasonUnknownFailureReason = AuthenticationEventFailureReason("unknown_failure_reason")
	AuthenticationEventFailureReasonInvalidMfa           = AuthenticationEventFailureReason("invalid_mfa")
	AuthenticationEventFailureReasonInvalidPassword      = AuthenticationEventFailureReason("invalid_password")
)

func (enum AuthenticationEventFailureReason) String() string {
	if enum == "" {
		// return default value if empty
		return string(AuthenticationEventFailureReasonUnknownFailureReason)
	}
	return string(enum)
}

func (enum AuthenticationEventFailureReason) Values() []AuthenticationEventFailureReason {
	return []AuthenticationEventFailureReason{
		"unknown_failure_reason",
		"invalid_mfa",
		"invalid_password",
	}
}

func (enum AuthenticationEventFailureReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AuthenticationEventFailureReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AuthenticationEventFailureReason(AuthenticationEventFailureReason(tmp).String())
	return nil
}

type AuthenticationEventMFAType string

const (
	AuthenticationEventMFATypeUnknownMfaType = AuthenticationEventMFAType("unknown_mfa_type")
	AuthenticationEventMFATypeTotp           = AuthenticationEventMFAType("totp")
)

func (enum AuthenticationEventMFAType) String() string {
	if enum == "" {
		// return default value if empty
		return string(AuthenticationEventMFATypeUnknownMfaType)
	}
	return string(enum)
}

func (enum AuthenticationEventMFAType) Values() []AuthenticationEventMFAType {
	return []AuthenticationEventMFAType{
		"unknown_mfa_type",
		"totp",
	}
}

func (enum AuthenticationEventMFAType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AuthenticationEventMFAType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AuthenticationEventMFAType(AuthenticationEventMFAType(tmp).String())
	return nil
}

type AuthenticationEventMethod string

const (
	AuthenticationEventMethodUnknownMethod      = AuthenticationEventMethod("unknown_method")
	AuthenticationEventMethodPassword           = AuthenticationEventMethod("password")
	AuthenticationEventMethodAuthenticationCode = AuthenticationEventMethod("authentication_code")
	AuthenticationEventMethodOauth2             = AuthenticationEventMethod("oauth2")
	AuthenticationEventMethodSaml               = AuthenticationEventMethod("saml")
)

func (enum AuthenticationEventMethod) String() string {
	if enum == "" {
		// return default value if empty
		return string(AuthenticationEventMethodUnknownMethod)
	}
	return string(enum)
}

func (enum AuthenticationEventMethod) Values() []AuthenticationEventMethod {
	return []AuthenticationEventMethod{
		"unknown_method",
		"password",
		"authentication_code",
		"oauth2",
		"saml",
	}
}

func (enum AuthenticationEventMethod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AuthenticationEventMethod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AuthenticationEventMethod(AuthenticationEventMethod(tmp).String())
	return nil
}

type AuthenticationEventOrigin string

const (
	AuthenticationEventOriginUnknownOrigin = AuthenticationEventOrigin("unknown_origin")
	AuthenticationEventOriginPublicAPI     = AuthenticationEventOrigin("public_api")
	AuthenticationEventOriginAdminAPI      = AuthenticationEventOrigin("admin_api")
)

func (enum AuthenticationEventOrigin) String() string {
	if enum == "" {
		// return default value if empty
		return string(AuthenticationEventOriginUnknownOrigin)
	}
	return string(enum)
}

func (enum AuthenticationEventOrigin) Values() []AuthenticationEventOrigin {
	return []AuthenticationEventOrigin{
		"unknown_origin",
		"public_api",
		"admin_api",
	}
}

func (enum AuthenticationEventOrigin) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AuthenticationEventOrigin) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AuthenticationEventOrigin(AuthenticationEventOrigin(tmp).String())
	return nil
}

type AuthenticationEventResult string

const (
	AuthenticationEventResultUnknownResult = AuthenticationEventResult("unknown_result")
	AuthenticationEventResultSuccess       = AuthenticationEventResult("success")
	AuthenticationEventResultFailure       = AuthenticationEventResult("failure")
)

func (enum AuthenticationEventResult) String() string {
	if enum == "" {
		// return default value if empty
		return string(AuthenticationEventResultUnknownResult)
	}
	return string(enum)
}

func (enum AuthenticationEventResult) Values() []AuthenticationEventResult {
	return []AuthenticationEventResult{
		"unknown_result",
		"success",
		"failure",
	}
}

func (enum AuthenticationEventResult) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AuthenticationEventResult) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AuthenticationEventResult(AuthenticationEventResult(tmp).String())
	return nil
}

type ExportJobStatusCode string

const (
	ExportJobStatusCodeUnknownCode = ExportJobStatusCode("unknown_code")
	ExportJobStatusCodeSuccess     = ExportJobStatusCode("success")
	ExportJobStatusCodeFailure     = ExportJobStatusCode("failure")
)

func (enum ExportJobStatusCode) String() string {
	if enum == "" {
		// return default value if empty
		return string(ExportJobStatusCodeUnknownCode)
	}
	return string(enum)
}

func (enum ExportJobStatusCode) Values() []ExportJobStatusCode {
	return []ExportJobStatusCode{
		"unknown_code",
		"success",
		"failure",
	}
}

func (enum ExportJobStatusCode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ExportJobStatusCode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ExportJobStatusCode(ExportJobStatusCode(tmp).String())
	return nil
}

type ListAuthenticationEventsRequestOrderBy string

const (
	ListAuthenticationEventsRequestOrderByRecordedAtDesc = ListAuthenticationEventsRequestOrderBy("recorded_at_desc")
	ListAuthenticationEventsRequestOrderByRecordedAtAsc  = ListAuthenticationEventsRequestOrderBy("recorded_at_asc")
)

func (enum ListAuthenticationEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListAuthenticationEventsRequestOrderByRecordedAtDesc)
	}
	return string(enum)
}

func (enum ListAuthenticationEventsRequestOrderBy) Values() []ListAuthenticationEventsRequestOrderBy {
	return []ListAuthenticationEventsRequestOrderBy{
		"recorded_at_desc",
		"recorded_at_asc",
	}
}

func (enum ListAuthenticationEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListAuthenticationEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListAuthenticationEventsRequestOrderBy(ListAuthenticationEventsRequestOrderBy(tmp).String())
	return nil
}

type ListCombinedEventsRequestOrderBy string

const (
	ListCombinedEventsRequestOrderByRecordedAtDesc = ListCombinedEventsRequestOrderBy("recorded_at_desc")
	ListCombinedEventsRequestOrderByRecordedAtAsc  = ListCombinedEventsRequestOrderBy("recorded_at_asc")
)

func (enum ListCombinedEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListCombinedEventsRequestOrderByRecordedAtDesc)
	}
	return string(enum)
}

func (enum ListCombinedEventsRequestOrderBy) Values() []ListCombinedEventsRequestOrderBy {
	return []ListCombinedEventsRequestOrderBy{
		"recorded_at_desc",
		"recorded_at_asc",
	}
}

func (enum ListCombinedEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCombinedEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCombinedEventsRequestOrderBy(ListCombinedEventsRequestOrderBy(tmp).String())
	return nil
}

type ListEventsRequestOrderBy string

const (
	ListEventsRequestOrderByRecordedAtDesc = ListEventsRequestOrderBy("recorded_at_desc")
	ListEventsRequestOrderByRecordedAtAsc  = ListEventsRequestOrderBy("recorded_at_asc")
)

func (enum ListEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListEventsRequestOrderByRecordedAtDesc)
	}
	return string(enum)
}

func (enum ListEventsRequestOrderBy) Values() []ListEventsRequestOrderBy {
	return []ListEventsRequestOrderBy{
		"recorded_at_desc",
		"recorded_at_asc",
	}
}

func (enum ListEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListEventsRequestOrderBy(ListEventsRequestOrderBy(tmp).String())
	return nil
}

type ListExportJobsRequestOrderBy string

const (
	ListExportJobsRequestOrderByNameAsc       = ListExportJobsRequestOrderBy("name_asc")
	ListExportJobsRequestOrderByNameDesc      = ListExportJobsRequestOrderBy("name_desc")
	ListExportJobsRequestOrderByCreatedAtAsc  = ListExportJobsRequestOrderBy("created_at_asc")
	ListExportJobsRequestOrderByCreatedAtDesc = ListExportJobsRequestOrderBy("created_at_desc")
)

func (enum ListExportJobsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListExportJobsRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListExportJobsRequestOrderBy) Values() []ListExportJobsRequestOrderBy {
	return []ListExportJobsRequestOrderBy{
		"name_asc",
		"name_desc",
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListExportJobsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListExportJobsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListExportJobsRequestOrderBy(ListExportJobsRequestOrderBy(tmp).String())
	return nil
}

type ResourceType string

const (
	ResourceTypeUnknownType              = ResourceType("unknown_type")
	ResourceTypeSecmSecret               = ResourceType("secm_secret")
	ResourceTypeSecmSecretVersion        = ResourceType("secm_secret_version")
	ResourceTypeKubeCluster              = ResourceType("kube_cluster")
	ResourceTypeKubePool                 = ResourceType("kube_pool")
	ResourceTypeKubeNode                 = ResourceType("kube_node")
	ResourceTypeKubeACL                  = ResourceType("kube_acl")
	ResourceTypeKeymKey                  = ResourceType("keym_key")
	ResourceTypeIamUser                  = ResourceType("iam_user")
	ResourceTypeIamApplication           = ResourceType("iam_application")
	ResourceTypeIamGroup                 = ResourceType("iam_group")
	ResourceTypeIamPolicy                = ResourceType("iam_policy")
	ResourceTypeIamAPIKey                = ResourceType("iam_api_key")
	ResourceTypeIamSSHKey                = ResourceType("iam_ssh_key")
	ResourceTypeIamRule                  = ResourceType("iam_rule")
	ResourceTypeIamSaml                  = ResourceType("iam_saml")
	ResourceTypeIamSamlCertificate       = ResourceType("iam_saml_certificate")
	ResourceTypeIamScim                  = ResourceType("iam_scim")
	ResourceTypeIamScimToken             = ResourceType("iam_scim_token")
	ResourceTypeSecretManagerSecret      = ResourceType("secret_manager_secret")
	ResourceTypeSecretManagerVersion     = ResourceType("secret_manager_version")
	ResourceTypeKeyManagerKey            = ResourceType("key_manager_key")
	ResourceTypeAccountUser              = ResourceType("account_user")
	ResourceTypeAccountOrganization      = ResourceType("account_organization")
	ResourceTypeAccountProject           = ResourceType("account_project")
	ResourceTypeAccountContractSignature = ResourceType("account_contract_signature")
	ResourceTypeInstanceServer           = ResourceType("instance_server")
	ResourceTypeInstancePlacementGroup   = ResourceType("instance_placement_group")
	ResourceTypeInstanceSecurityGroup    = ResourceType("instance_security_group")
	ResourceTypeInstanceVolume           = ResourceType("instance_volume")
	ResourceTypeInstanceSnapshot         = ResourceType("instance_snapshot")
	ResourceTypeInstanceImage            = ResourceType("instance_image")
	ResourceTypeInstanceTemplate         = ResourceType("instance_template")
	ResourceTypeAppleSiliconServer       = ResourceType("apple_silicon_server")
	ResourceTypeBaremetalServer          = ResourceType("baremetal_server")
	ResourceTypeBaremetalSetting         = ResourceType("baremetal_setting")
	ResourceTypeIpamIP                   = ResourceType("ipam_ip")
	ResourceTypeSbsVolume                = ResourceType("sbs_volume")
	ResourceTypeSbsSnapshot              = ResourceType("sbs_snapshot")
	ResourceTypeLoadBalancerLB           = ResourceType("load_balancer_lb")
	ResourceTypeLoadBalancerIP           = ResourceType("load_balancer_ip")
	ResourceTypeLoadBalancerFrontend     = ResourceType("load_balancer_frontend")
	ResourceTypeLoadBalancerBackend      = ResourceType("load_balancer_backend")
	ResourceTypeLoadBalancerRoute        = ResourceType("load_balancer_route")
	ResourceTypeLoadBalancerACL          = ResourceType("load_balancer_acl")
	ResourceTypeLoadBalancerCertificate  = ResourceType("load_balancer_certificate")
	ResourceTypeSfsFilesystem            = ResourceType("sfs_filesystem")
	ResourceTypeVpcPrivateNetwork        = ResourceType("vpc_private_network")
	ResourceTypeVpcVpc                   = ResourceType("vpc_vpc")
	ResourceTypeVpcSubnet                = ResourceType("vpc_subnet")
	ResourceTypeVpcRoute                 = ResourceType("vpc_route")
	ResourceTypeVpcACL                   = ResourceType("vpc_acl")
	ResourceTypeEdgeServicesPlan         = ResourceType("edge_services_plan")
	ResourceTypeEdgeServicesPipeline     = ResourceType("edge_services_pipeline")
	ResourceTypeEdgeServicesDNSStage     = ResourceType("edge_services_dns_stage")
	ResourceTypeEdgeServicesTLSStage     = ResourceType("edge_services_tls_stage")
	ResourceTypeEdgeServicesCacheStage   = ResourceType("edge_services_cache_stage")
	ResourceTypeEdgeServicesRouteStage   = ResourceType("edge_services_route_stage")
	ResourceTypeEdgeServicesRouteRules   = ResourceType("edge_services_route_rules")
	ResourceTypeEdgeServicesWafStage     = ResourceType("edge_services_waf_stage")
	ResourceTypeEdgeServicesBackendStage = ResourceType("edge_services_backend_stage")
	ResourceTypeS2sVpnGateway            = ResourceType("s2s_vpn_gateway")
	ResourceTypeS2sCustomerGateway       = ResourceType("s2s_customer_gateway")
	ResourceTypeS2sRoutingPolicy         = ResourceType("s2s_routing_policy")
	ResourceTypeS2sConnection            = ResourceType("s2s_connection")
	ResourceTypeVpcGwGateway             = ResourceType("vpc_gw_gateway")
	ResourceTypeVpcGwGatewayNetwork      = ResourceType("vpc_gw_gateway_network")
	ResourceTypeVpcGwDHCP                = ResourceType("vpc_gw_dhcp")
	ResourceTypeVpcGwDHCPEntry           = ResourceType("vpc_gw_dhcp_entry")
	ResourceTypeVpcGwPatRule             = ResourceType("vpc_gw_pat_rule")
	ResourceTypeVpcGwIP                  = ResourceType("vpc_gw_ip")
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
		"secm_secret",
		"secm_secret_version",
		"kube_cluster",
		"kube_pool",
		"kube_node",
		"kube_acl",
		"keym_key",
		"iam_user",
		"iam_application",
		"iam_group",
		"iam_policy",
		"iam_api_key",
		"iam_ssh_key",
		"iam_rule",
		"iam_saml",
		"iam_saml_certificate",
		"iam_scim",
		"iam_scim_token",
		"secret_manager_secret",
		"secret_manager_version",
		"key_manager_key",
		"account_user",
		"account_organization",
		"account_project",
		"account_contract_signature",
		"instance_server",
		"instance_placement_group",
		"instance_security_group",
		"instance_volume",
		"instance_snapshot",
		"instance_image",
		"instance_template",
		"apple_silicon_server",
		"baremetal_server",
		"baremetal_setting",
		"ipam_ip",
		"sbs_volume",
		"sbs_snapshot",
		"load_balancer_lb",
		"load_balancer_ip",
		"load_balancer_frontend",
		"load_balancer_backend",
		"load_balancer_route",
		"load_balancer_acl",
		"load_balancer_certificate",
		"sfs_filesystem",
		"vpc_private_network",
		"vpc_vpc",
		"vpc_subnet",
		"vpc_route",
		"vpc_acl",
		"edge_services_plan",
		"edge_services_pipeline",
		"edge_services_dns_stage",
		"edge_services_tls_stage",
		"edge_services_cache_stage",
		"edge_services_route_stage",
		"edge_services_route_rules",
		"edge_services_waf_stage",
		"edge_services_backend_stage",
		"s2s_vpn_gateway",
		"s2s_customer_gateway",
		"s2s_routing_policy",
		"s2s_connection",
		"vpc_gw_gateway",
		"vpc_gw_gateway_network",
		"vpc_gw_dhcp",
		"vpc_gw_dhcp_entry",
		"vpc_gw_pat_rule",
		"vpc_gw_ip",
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

type SystemEventKind string

const (
	SystemEventKindUnknownKind  = SystemEventKind("unknown_kind")
	SystemEventKindCron         = SystemEventKind("cron")
	SystemEventKindNotification = SystemEventKind("notification")
)

func (enum SystemEventKind) String() string {
	if enum == "" {
		// return default value if empty
		return string(SystemEventKindUnknownKind)
	}
	return string(enum)
}

func (enum SystemEventKind) Values() []SystemEventKind {
	return []SystemEventKind{
		"unknown_kind",
		"cron",
		"notification",
	}
}

func (enum SystemEventKind) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SystemEventKind) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SystemEventKind(SystemEventKind(tmp).String())
	return nil
}

// AccountContractSignatureInfoAccountContractInfo: account contract signature info account contract info.
type AccountContractSignatureInfoAccountContractInfo struct {
	ID string `json:"id"`

	Type string `json:"type"`

	Name string `json:"name"`

	Version int32 `json:"version"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// AccountContractSignatureInfo: account contract signature info.
type AccountContractSignatureInfo struct {
	SignedAt *time.Time `json:"signed_at"`

	SignedByAccountRootUserID string `json:"signed_by_account_root_user_id"`

	ExpiresAt *time.Time `json:"expires_at"`

	Contract *AccountContractSignatureInfoAccountContractInfo `json:"contract"`
}

// AccountOrganizationInfo: account organization info.
type AccountOrganizationInfo struct{}

// AccountProjectInfo: account project info.
type AccountProjectInfo struct {
	Description string `json:"description"`
}

// AccountUserInfo: account user info.
type AccountUserInfo struct {
	Email string `json:"email"`

	PhoneNumber *string `json:"phone_number"`
}

// AppleSiliconServerInfo: apple silicon server info.
type AppleSiliconServerInfo struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

// BaremetalServerInfo: baremetal server info.
type BaremetalServerInfo struct {
	Description string `json:"description"`

	Tags []string `json:"tags"`
}

// BaremetalSettingInfo: baremetal setting info.
type BaremetalSettingInfo struct {
	Type string `json:"type"`
}

// EdgeServicesBackendStageInfo: edge services backend stage info.
type EdgeServicesBackendStageInfo struct {
	PipelineID *string `json:"pipeline_id"`
}

// EdgeServicesCacheStageInfo: edge services cache stage info.
type EdgeServicesCacheStageInfo struct {
	PipelineID *string `json:"pipeline_id"`
}

// EdgeServicesDNSStageInfo: edge services dns stage info.
type EdgeServicesDNSStageInfo struct {
	PipelineID *string `json:"pipeline_id"`
}

// EdgeServicesPipelineInfo: edge services pipeline info.
type EdgeServicesPipelineInfo struct {
	Name string `json:"name"`
}

// EdgeServicesPlanInfo: edge services plan info.
type EdgeServicesPlanInfo struct{}

// EdgeServicesRouteRulesInfo: edge services route rules info.
type EdgeServicesRouteRulesInfo struct {
	RouteStageID string `json:"route_stage_id"`
}

// EdgeServicesRouteStageInfo: edge services route stage info.
type EdgeServicesRouteStageInfo struct {
	PipelineID *string `json:"pipeline_id"`
}

// EdgeServicesTLSStageInfo: edge services tls stage info.
type EdgeServicesTLSStageInfo struct {
	PipelineID *string `json:"pipeline_id"`
}

// EdgeServicesWAFStageInfo: edge services waf stage info.
type EdgeServicesWAFStageInfo struct {
	PipelineID *string `json:"pipeline_id"`
}

// InstanceServerInfo: instance server info.
type InstanceServerInfo struct {
	Name string `json:"name"`
}

// IpamIPInfo: ipam ip info.
type IpamIPInfo struct {
	Address scw.IPNet `json:"address"`
}

// KeyManagerKeyInfo: key manager key info.
type KeyManagerKeyInfo struct{}

// KubernetesACLInfo: kubernetes acl info.
type KubernetesACLInfo struct{}

// KubernetesClusterInfo: kubernetes cluster info.
type KubernetesClusterInfo struct{}

// KubernetesNodeInfo: kubernetes node info.
type KubernetesNodeInfo struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

// KubernetesPoolInfo: kubernetes pool info.
type KubernetesPoolInfo struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

// LoadBalancerACLInfo: load balancer acl info.
type LoadBalancerACLInfo struct {
	FrontendID string `json:"frontend_id"`
}

// LoadBalancerBackendInfo: load balancer backend info.
type LoadBalancerBackendInfo struct {
	LBID string `json:"lb_id"`

	Name string `json:"name"`
}

// LoadBalancerCertificateInfo: load balancer certificate info.
type LoadBalancerCertificateInfo struct {
	LBID string `json:"lb_id"`

	Name string `json:"name"`
}

// LoadBalancerFrontendInfo: load balancer frontend info.
type LoadBalancerFrontendInfo struct {
	LBID string `json:"lb_id"`

	Name string `json:"name"`
}

// LoadBalancerIPInfo: load balancer ip info.
type LoadBalancerIPInfo struct {
	IPAddress string `json:"ip_address"`

	LBID *string `json:"lb_id"`
}

// LoadBalancerLBInfo: load balancer lb info.
type LoadBalancerLBInfo struct {
	Name string `json:"name"`
}

// LoadBalancerRouteInfo: load balancer route info.
type LoadBalancerRouteInfo struct {
	FrontendID string `json:"frontend_id"`

	BackendID string `json:"backend_id"`
}

// SecretManagerSecretInfo: secret manager secret info.
type SecretManagerSecretInfo struct {
	Path string `json:"path"`

	KeyID *string `json:"key_id"`
}

// SecretManagerSecretVersionInfo: secret manager secret version info.
type SecretManagerSecretVersionInfo struct {
	Revision uint32 `json:"revision"`
}

// VpcPrivateNetworkInfo: vpc private network info.
type VpcPrivateNetworkInfo struct {
	VpcID string `json:"vpc_id"`

	PushDefaultRoute bool `json:"push_default_route"`
}

// VpcRouteInfo: vpc route info.
type VpcRouteInfo struct {
	VpcID string `json:"vpc_id"`

	Destination scw.IPNet `json:"destination"`

	NexthopResourceKey *string `json:"nexthop_resource_key"`

	NexthopPrivateNetworkKey *string `json:"nexthop_private_network_key"`
}

// VpcSubnetInfo: vpc subnet info.
type VpcSubnetInfo struct {
	SubnetCidr scw.IPNet `json:"subnet_cidr"`

	VpcID string `json:"vpc_id"`
}

// Resource: resource.
type Resource struct {
	ID string `json:"id"`

	// Type: default value: unknown_type
	Type ResourceType `json:"type"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	Name *string `json:"name"`

	// Deprecated
	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	SecmSecretInfo *SecretManagerSecretInfo `json:"secm_secret_info,omitempty"`

	// Deprecated
	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	SecmSecretVersionInfo *SecretManagerSecretVersionInfo `json:"secm_secret_version_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	KubeClusterInfo *KubernetesClusterInfo `json:"kube_cluster_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	KubePoolInfo *KubernetesPoolInfo `json:"kube_pool_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	KubeNodeInfo *KubernetesNodeInfo `json:"kube_node_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	KubeACLInfo *KubernetesACLInfo `json:"kube_acl_info,omitempty"`

	// Deprecated
	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	KeymKeyInfo *KeyManagerKeyInfo `json:"keym_key_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	SecretManagerSecretInfo *SecretManagerSecretInfo `json:"secret_manager_secret_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	SecretManagerVersionInfo *SecretManagerSecretVersionInfo `json:"secret_manager_version_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	KeyManagerKeyInfo *KeyManagerKeyInfo `json:"key_manager_key_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	AccountUserInfo *AccountUserInfo `json:"account_user_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	AccountOrganizationInfo *AccountOrganizationInfo `json:"account_organization_info,omitempty"`

	// Deprecated
	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	InstanceServerInfo *InstanceServerInfo `json:"instance_server_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	AppleSiliconServerInfo *AppleSiliconServerInfo `json:"apple_silicon_server_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	AccountProjectInfo *AccountProjectInfo `json:"account_project_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	BaremetalServerInfo *BaremetalServerInfo `json:"baremetal_server_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	BaremetalSettingInfo *BaremetalSettingInfo `json:"baremetal_setting_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	IpamIPInfo *IpamIPInfo `json:"ipam_ip_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	LoadBalancerLBInfo *LoadBalancerLBInfo `json:"load_balancer_lb_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	LoadBalancerIPInfo *LoadBalancerIPInfo `json:"load_balancer_ip_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	LoadBalancerFrontendInfo *LoadBalancerFrontendInfo `json:"load_balancer_frontend_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	LoadBalancerBackendInfo *LoadBalancerBackendInfo `json:"load_balancer_backend_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	LoadBalancerRouteInfo *LoadBalancerRouteInfo `json:"load_balancer_route_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	LoadBalancerACLInfo *LoadBalancerACLInfo `json:"load_balancer_acl_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	LoadBalancerCertificateInfo *LoadBalancerCertificateInfo `json:"load_balancer_certificate_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	EdgeServicesPlanInfo *EdgeServicesPlanInfo `json:"edge_services_plan_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	EdgeServicesPipelineInfo *EdgeServicesPipelineInfo `json:"edge_services_pipeline_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	EdgeServicesDNSStageInfo *EdgeServicesDNSStageInfo `json:"edge_services_dns_stage_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	EdgeServicesTLSStageInfo *EdgeServicesTLSStageInfo `json:"edge_services_tls_stage_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	EdgeServicesCacheStageInfo *EdgeServicesCacheStageInfo `json:"edge_services_cache_stage_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	EdgeServicesRouteStageInfo *EdgeServicesRouteStageInfo `json:"edge_services_route_stage_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	EdgeServicesRouteRulesInfo *EdgeServicesRouteRulesInfo `json:"edge_services_route_rules_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	EdgeServicesWafStageInfo *EdgeServicesWAFStageInfo `json:"edge_services_waf_stage_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	EdgeServicesBackendStageInfo *EdgeServicesBackendStageInfo `json:"edge_services_backend_stage_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	AccountContractSignatureInfo *AccountContractSignatureInfo `json:"account_contract_signature_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	VpcSubnetInfo *VpcSubnetInfo `json:"vpc_subnet_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	VpcRouteInfo *VpcRouteInfo `json:"vpc_route_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo, InstanceServerInfo, AppleSiliconServerInfo, AccountProjectInfo, BaremetalServerInfo, BaremetalSettingInfo, IpamIPInfo, LoadBalancerLBInfo, LoadBalancerIPInfo, LoadBalancerFrontendInfo, LoadBalancerBackendInfo, LoadBalancerRouteInfo, LoadBalancerACLInfo, LoadBalancerCertificateInfo, EdgeServicesPlanInfo, EdgeServicesPipelineInfo, EdgeServicesDNSStageInfo, EdgeServicesTLSStageInfo, EdgeServicesCacheStageInfo, EdgeServicesRouteStageInfo, EdgeServicesRouteRulesInfo, EdgeServicesWafStageInfo, EdgeServicesBackendStageInfo, AccountContractSignatureInfo, VpcSubnetInfo, VpcRouteInfo, VpcPrivateNetworkInfo must be set.
	VpcPrivateNetworkInfo *VpcPrivateNetworkInfo `json:"vpc_private_network_info,omitempty"`
}

// EventPrincipal: event principal.
type EventPrincipal struct {
	ID string `json:"id"`
}

// AuthenticationEvent: authentication event.
type AuthenticationEvent struct {
	// ID: ID of the event.
	ID string `json:"id"`

	// RecordedAt: timestamp of the event.
	RecordedAt *time.Time `json:"recorded_at"`

	// OrganizationID: organization ID containing the event.
	OrganizationID string `json:"organization_id"`

	// SourceIP: IP address at the origin of the event.
	SourceIP net.IP `json:"source_ip"`

	// UserAgent: user Agent at the origin of the event.
	UserAgent *string `json:"user_agent"`

	// Resources: resources attached to the event.
	Resources []*Resource `json:"resources"`

	// Result: result of the authentication attempt.
	// Default value: unknown_result
	Result AuthenticationEventResult `json:"result"`

	// FailureReason: (Optional) Reason for authentication failure.
	// Default value: unknown_failure_reason
	FailureReason *AuthenticationEventFailureReason `json:"failure_reason"`

	// CountryCode: (Optional) ISO 3166-1 alpha-2 country code of the source IP.
	// Default value: unknown_country_code
	CountryCode *std.CountryCode `json:"country_code"`

	// Method: authentication method used.
	// Default value: unknown_method
	Method AuthenticationEventMethod `json:"method"`

	// Origin: origin of the authentication attempt.
	// Default value: unknown_origin
	Origin AuthenticationEventOrigin `json:"origin"`

	// MfaType: (Optional) MFA type used for the authentication attempt.
	// Default value: unknown_mfa_type
	MfaType *AuthenticationEventMFAType `json:"mfa_type"`
}

// Event: event.
type Event struct {
	// ID: ID of the event.
	ID string `json:"id"`

	// RecordedAt: timestamp of the event.
	RecordedAt *time.Time `json:"recorded_at"`

	// Locality: locality of the resource attached to the event.
	Locality string `json:"locality"`

	// Principal: user or IAM application at the origin of the event.
	// Precisely one of Principal must be set.
	Principal *EventPrincipal `json:"principal,omitempty"`

	// OrganizationID: organization ID containing the event.
	OrganizationID string `json:"organization_id"`

	// ProjectID: (Optional) Project of the resource attached to the event.
	ProjectID *string `json:"project_id"`

	// SourceIP: IP address at the origin of the event.
	SourceIP net.IP `json:"source_ip"`

	// UserAgent: user Agent at the origin of the event.
	UserAgent *string `json:"user_agent"`

	// ProductName: product name of the resource attached to the event.
	ProductName string `json:"product_name"`

	// ServiceName: API name called to trigger the event.
	ServiceName string `json:"service_name"`

	// MethodName: API method called to trigger the event.
	MethodName string `json:"method_name"`

	// Resources: resources attached to the event.
	Resources []*Resource `json:"resources"`

	// RequestID: unique identifier of the request at the origin of the event.
	RequestID string `json:"request_id"`

	// RequestBody: request at the origin of the event.
	RequestBody *scw.JSONObject `json:"request_body"`

	// StatusCode: HTTP status code resulting of the API call.
	StatusCode uint32 `json:"status_code"`
}

// SystemEvent: system event.
type SystemEvent struct {
	ID string `json:"id"`

	RecordedAt *time.Time `json:"recorded_at"`

	Locality string `json:"locality"`

	OrganizationID string `json:"organization_id"`

	ProjectID *string `json:"project_id"`

	Source string `json:"source"`

	SystemName string `json:"system_name"`

	Resources []*Resource `json:"resources"`

	// Kind: default value: unknown_kind
	Kind SystemEventKind `json:"kind"`

	ProductName string `json:"product_name"`
}

// ExportJobS3: export job s3.
type ExportJobS3 struct {
	Bucket string `json:"bucket"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`

	Prefix *string `json:"prefix"`

	ProjectID *string `json:"project_id"`
}

// ExportJobStatus: export job status.
type ExportJobStatus struct {
	// Code: default value: unknown_code
	Code ExportJobStatusCode `json:"code"`

	Message *string `json:"message"`
}

// ProductService: product service.
type ProductService struct {
	Name string `json:"name"`

	Methods []string `json:"methods"`
}

// ListCombinedEventsResponseCombinedEvent: list combined events response combined event.
type ListCombinedEventsResponseCombinedEvent struct {
	// Precisely one of API, Auth, System must be set.
	API *Event `json:"api,omitempty"`

	// Precisely one of API, Auth, System must be set.
	Auth *AuthenticationEvent `json:"auth,omitempty"`

	// Precisely one of API, Auth, System must be set.
	System *SystemEvent `json:"system,omitempty"`
}

// ExportJob: export job.
type ExportJob struct {
	// ID: ID of the export job.
	ID string `json:"id"`

	// OrganizationID: ID of the targeted Organization.
	OrganizationID string `json:"organization_id"`

	// Name: name of the export job.
	Name string `json:"name"`

	// S3: destination in an S3 storage.
	// Precisely one of S3 must be set.
	S3 *ExportJobS3 `json:"s3,omitempty"`

	// CreatedAt: export job creation date.
	CreatedAt *time.Time `json:"created_at"`

	// LastRunAt: last run of export job.
	LastRunAt *time.Time `json:"last_run_at"`

	// Tags: tags of the export job.
	Tags []string `json:"tags"`

	// LastStatus: status of last export job.
	LastStatus *ExportJobStatus `json:"last_status"`
}

// Product: product.
type Product struct {
	// Title: product title.
	Title string `json:"title"`

	// Name: product name.
	Name string `json:"name"`

	// Services: specifies the API versions of the products integrated with Audit Trail. Each version defines the methods logged by Audit Trail.
	Services []*ProductService `json:"services"`
}

// CreateExportJobRequest: create export job request.
type CreateExportJobRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrganizationID: ID of the Organization to target.
	OrganizationID string `json:"organization_id"`

	// Name: name of the export.
	Name string `json:"name"`

	// S3: the configuration specifying the bucket where the audit trail events will be exported.
	// Precisely one of S3 must be set.
	S3 *ExportJobS3 `json:"s3,omitempty"`

	// Tags: tags of the export.
	Tags []string `json:"tags"`
}

// DeleteExportJobRequest: delete export job request.
type DeleteExportJobRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ExportJobID: ID of the export job.
	ExportJobID string `json:"-"`
}

// ListAuthenticationEventsRequest: list authentication events request.
type ListAuthenticationEventsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	OrganizationID string `json:"-"`

	RecordedAfter *time.Time `json:"-"`

	RecordedBefore *time.Time `json:"-"`

	// OrderBy: default value: recorded_at_desc
	OrderBy ListAuthenticationEventsRequestOrderBy `json:"-"`

	PageSize *uint32 `json:"-"`

	PageToken *string `json:"-"`
}

// ListAuthenticationEventsResponse: list authentication events response.
type ListAuthenticationEventsResponse struct {
	Events []*AuthenticationEvent `json:"events"`

	NextPageToken *string `json:"next_page_token"`
}

// ListCombinedEventsRequest: list combined events request.
type ListCombinedEventsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	OrganizationID string `json:"-"`

	ProjectID *string `json:"-"`

	// ResourceType: default value: unknown_type
	ResourceType ResourceType `json:"-"`

	RecordedAfter *time.Time `json:"-"`

	RecordedBefore *time.Time `json:"-"`

	// OrderBy: default value: recorded_at_desc
	OrderBy ListCombinedEventsRequestOrderBy `json:"-"`

	PageSize *uint32 `json:"-"`

	PageToken *string `json:"-"`
}

// ListCombinedEventsResponse: list combined events response.
type ListCombinedEventsResponse struct {
	Events []*ListCombinedEventsResponseCombinedEvent `json:"events"`

	NextPageToken *string `json:"next_page_token"`
}

// ListEventsRequest: list events request.
type ListEventsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: (Optional) ID of the Project containing the Audit Trail events.
	ProjectID *string `json:"-"`

	// OrganizationID: ID of the Organization containing the Audit Trail events.
	OrganizationID string `json:"-"`

	// ResourceType: (Optional) Type of the Scaleway resource.
	// Default value: unknown_type
	ResourceType ResourceType `json:"-"`

	// MethodName: (Optional) Name of the method of the API call performed.
	MethodName *string `json:"-"`

	// Status: (Optional) HTTP status code of the request. Returns either `200` if the request was successful or `403` if the permission was denied.
	Status *uint32 `json:"-"`

	// RecordedAfter: (Optional) The `recorded_after` parameter defines the earliest timestamp from which Audit Trail events are retrieved. Returns `one hour ago` by default.
	RecordedAfter *time.Time `json:"-"`

	// RecordedBefore: (Optional) The `recorded_before` parameter defines the latest timestamp up to which Audit Trail events are retrieved. Returns `now` by default.
	RecordedBefore *time.Time `json:"-"`

	// OrderBy: default value: recorded_at_desc
	OrderBy ListEventsRequestOrderBy `json:"-"`

	PageSize *uint32 `json:"-"`

	PageToken *string `json:"-"`

	// ProductName: (Optional) Name of the Scaleway product in a hyphenated format.
	ProductName *string `json:"-"`

	// ServiceName: (Optional) Name of the service of the API call performed.
	ServiceName *string `json:"-"`

	// ResourceID: (Optional) ID of the Scaleway resource.
	ResourceID *string `json:"-"`

	// PrincipalID: (Optional) ID of the User or IAM application at the origin of the event.
	PrincipalID *string `json:"-"`

	// SourceIP: (Optional) IP address at the origin of the event.
	SourceIP *string `json:"-"`
}

// ListEventsResponse: list events response.
type ListEventsResponse struct {
	// Events: single page of events matching the requested criteria.
	Events []*Event `json:"events"`

	// NextPageToken: page token to use in following calls to keep listing.
	NextPageToken *string `json:"next_page_token"`
}

// ListExportJobsRequest: list export jobs request.
type ListExportJobsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrganizationID: filter by Organization ID.
	OrganizationID string `json:"-"`

	// Name: (Optional) Filter by export name.
	Name *string `json:"-"`

	// Tags: (Optional) List of tags to filter on.
	Tags []string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListExportJobsRequestOrderBy `json:"-"`
}

// ListExportJobsResponse: list export jobs response.
type ListExportJobsResponse struct {
	// ExportJobs: single page of export jobs matching the requested criteria.
	ExportJobs []*ExportJob `json:"export_jobs"`

	// TotalCount: total count of export jobs matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListExportJobsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListExportJobsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListExportJobsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ExportJobs = append(r.ExportJobs, results.ExportJobs...)
	r.TotalCount += uint64(len(results.ExportJobs))
	return uint64(len(results.ExportJobs)), nil
}

// ListProductsRequest: list products request.
type ListProductsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrganizationID: ID of the Organization containing the Audit Trail events.
	OrganizationID string `json:"organization_id"`
}

// ListProductsResponse: list products response.
type ListProductsResponse struct {
	// Products: list of all products integrated with Audit Trail.
	Products []*Product `json:"products"`

	// TotalCount: number of integrated products.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListProductsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListProductsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListProductsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Products = append(r.Products, results.Products...)
	r.TotalCount += uint64(len(results.Products))
	return uint64(len(results.Products)), nil
}

// This API allows you to ensure accountability and security by recording events and changes performed within your Scaleway Organization.
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
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms}
}

// ListEvents: Retrieve the list of Audit Trail events for a Scaleway Organization and/or Project. You must specify the `organization_id` and optionally, the `project_id`.
func (s *API) ListEvents(req *ListEventsRequest, opts ...scw.RequestOption) (*ListEventsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "resource_type", req.ResourceType)
	parameter.AddToQuery(query, "method_name", req.MethodName)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "recorded_after", req.RecordedAfter)
	parameter.AddToQuery(query, "recorded_before", req.RecordedBefore)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "product_name", req.ProductName)
	parameter.AddToQuery(query, "service_name", req.ServiceName)
	parameter.AddToQuery(query, "resource_id", req.ResourceID)
	parameter.AddToQuery(query, "principal_id", req.PrincipalID)
	parameter.AddToQuery(query, "source_ip", req.SourceIP)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/audit-trail/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/events",
		Query:  query,
	}

	var resp ListEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListAuthenticationEvents: Retrieve the list of Audit Trail authentication events for a Scaleway Organization. You must specify the `organization_id`.
func (s *API) ListAuthenticationEvents(req *ListAuthenticationEventsRequest, opts ...scw.RequestOption) (*ListAuthenticationEventsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "recorded_after", req.RecordedAfter)
	parameter.AddToQuery(query, "recorded_before", req.RecordedBefore)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page_token", req.PageToken)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/audit-trail/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/authentication-events",
		Query:  query,
	}

	var resp ListAuthenticationEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCombinedEvents:
func (s *API) ListCombinedEvents(req *ListCombinedEventsRequest, opts ...scw.RequestOption) (*ListCombinedEventsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "resource_type", req.ResourceType)
	parameter.AddToQuery(query, "recorded_after", req.RecordedAfter)
	parameter.AddToQuery(query, "recorded_before", req.RecordedBefore)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page_token", req.PageToken)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/audit-trail/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/combined-events",
		Query:  query,
	}

	var resp ListCombinedEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListProducts: Retrieve the list of Scaleway resources for which you have Audit Trail events.
func (s *API) ListProducts(req *ListProductsRequest, opts ...scw.RequestOption) (*ListProductsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/audit-trail/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/products",
		Query:  query,
	}

	var resp ListProductsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateExportJob: Create an export job for a specified organization. This allows you to export audit trail events to a destination, such as an S3 bucket. The request requires the organization ID, a name for the export, and a destination configuration.
func (s *API) CreateExportJob(req *CreateExportJobRequest, opts ...scw.RequestOption) (*ExportJob, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/audit-trail/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/export-jobs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ExportJob

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteExportJob: Deletes an export job for a specified id.
func (s *API) DeleteExportJob(req *DeleteExportJobRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ExportJobID) == "" {
		return errors.New("field ExportJobID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/audit-trail/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/export-jobs/" + fmt.Sprint(req.ExportJobID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListExportJobs:
func (s *API) ListExportJobs(req *ListExportJobsRequest, opts ...scw.RequestOption) (*ListExportJobsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
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
		Path:   "/audit-trail/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/export-jobs",
		Query:  query,
	}

	var resp ListExportJobsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

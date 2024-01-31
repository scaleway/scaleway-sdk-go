// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package trustandsafety_private provides methods and message types of the trustandsafety_private v1 API.
package trustandsafety_private

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
	resource_v1alpha1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/resource/v1alpha1"
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

type LockReason string

const (
	LockReasonUnknown               = LockReason("unknown")
	LockReasonLockedForAbuse        = LockReason("locked_for_abuse")
	LockReasonInvoicePaymentFailure = LockReason("invoice_payment_failure")
	LockReasonIllegalActivity       = LockReason("illegal_activity")
	LockReasonLegalRequisition      = LockReason("legal_requisition")
	LockReasonMaintenance           = LockReason("maintenance")
	LockReasonFlood                 = LockReason("flood")
	LockReasonCriticalSecurityIssue = LockReason("critical_security_issue")
	LockReasonSecurityIssue         = LockReason("security_issue")
	LockReasonUltraSecurityIssue    = LockReason("ultra_security_issue")
)

func (enum LockReason) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum LockReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LockReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LockReason(LockReason(tmp).String())
	return nil
}

type ResourceStatus string

const (
	ResourceStatusUnlocked  = ResourceStatus("unlocked")
	ResourceStatusLocked    = ResourceStatus("locked")
	ResourceStatusLocking   = ResourceStatus("locking")
	ResourceStatusUnlocking = ResourceStatus("unlocking")
)

func (enum ResourceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unlocked"
	}
	return string(enum)
}

func (enum ResourceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ResourceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ResourceStatus(ResourceStatus(tmp).String())
	return nil
}

type ResourceType string

const (
	ResourceTypeUnknown                       = ResourceType("unknown")
	ResourceTypeInstanceServer                = ResourceType("instance_server")
	ResourceTypeInstanceIP                    = ResourceType("instance_ip")
	ResourceTypeInstanceSecurityGroup         = ResourceType("instance_security_group")
	ResourceTypeInstanceVolume                = ResourceType("instance_volume")
	ResourceTypeInstanceSnapshot              = ResourceType("instance_snapshot")
	ResourceTypeInstanceImage                 = ResourceType("instance_image")
	ResourceTypeLBLB                          = ResourceType("lb_lb")
	ResourceTypeLBIP                          = ResourceType("lb_ip")
	ResourceTypeKubernetesCluster             = ResourceType("kubernetes_cluster")
	ResourceTypeBaremetalServer               = ResourceType("baremetal_server")
	ResourceTypeBaremetalIPFailover           = ResourceType("baremetal_ip_failover")
	ResourceTypeObjectBucket                  = ResourceType("object_bucket")
	ResourceTypeObjectObject                  = ResourceType("object_object")
	ResourceTypeDomainDomain                  = ResourceType("domain_domain")
	ResourceTypeDomainDNS                     = ResourceType("domain_dns")
	ResourceTypeRegistryNamespace             = ResourceType("registry_namespace")
	ResourceTypeRegistryImage                 = ResourceType("registry_image")
	ResourceTypeRegistryTag                   = ResourceType("registry_tag")
	ResourceTypeVpcPrivateNetwork             = ResourceType("vpc_private_network")
	ResourceTypeRdbInstance                   = ResourceType("rdb_instance")
	ResourceTypeRdbBackup                     = ResourceType("rdb_backup")
	ResourceTypeFlexibleIPFip                 = ResourceType("flexible_ip_fip")
	ResourceTypeRdbSnapshot                   = ResourceType("rdb_snapshot")
	ResourceTypeAsServer                      = ResourceType("as_server")
	ResourceTypeVpcGwGateway                  = ResourceType("vpc_gw_gateway")
	ResourceTypeVpcGwDHCP                     = ResourceType("vpc_gw_dhcp")
	ResourceTypeVpcGwIP                       = ResourceType("vpc_gw_ip")
	ResourceTypeIotIP                         = ResourceType("iot_ip")
	ResourceTypeIotHub                        = ResourceType("iot_hub")
	ResourceTypeFncFunction                   = ResourceType("fnc_function")
	ResourceTypeFncNamespace                  = ResourceType("fnc_namespace")
	ResourceTypeCtnContainer                  = ResourceType("ctn_container")
	ResourceTypeCtnNamespace                  = ResourceType("ctn_namespace")
	ResourceTypeBaremetalPrivateNetworkMember = ResourceType("baremetal_private_network_member")
	ResourceTypeTemDomain                     = ResourceType("tem_domain")
	ResourceTypeRkvCluster                    = ResourceType("rkv_cluster")
	ResourceTypeWbhHosting                    = ResourceType("wbh_hosting")
	ResourceTypeObsCockpit                    = ResourceType("obs_cockpit")
	ResourceTypeSmSecret                      = ResourceType("sm_secret")
	ResourceTypeKmsKey                        = ResourceType("kms_key")
	ResourceTypeMnqNamespace                  = ResourceType("mnq_namespace")
	ResourceTypeIpamIP                        = ResourceType("ipam_ip")
	ResourceTypeIpfsCid                       = ResourceType("ipfs_cid")
	ResourceTypeIpfsVolume                    = ResourceType("ipfs_volume")
	ResourceTypeSdbInstance                   = ResourceType("sdb_instance")
	ResourceTypeSljJob                        = ResourceType("slj_job")
	ResourceTypeVpcVpc                        = ResourceType("vpc_vpc")
	ResourceTypeSbsVolume                     = ResourceType("sbs_volume")
	ResourceTypeSbsSnapshot                   = ResourceType("sbs_snapshot")
	ResourceTypeEdgPipeline                   = ResourceType("edg_pipeline")
	ResourceTypeIpfsNames                     = ResourceType("ipfs_names")
	ResourceTypeQaasSessions                  = ResourceType("qaas_sessions")
	ResourceTypeServerlessSqldbDatabase       = ResourceType("serverless_sqldb_database")
	ResourceTypeServerlessSqldbBackup         = ResourceType("serverless_sqldb_backup")
)

func (enum ResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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

// Resource: resource.
type Resource struct {
	ID string `json:"id"`

	// Type: default value: unknown
	Type ResourceType `json:"type"`

	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// Status: default value: unlocked
	Status ResourceStatus `json:"status"`

	// LockReason: default value: unknown
	LockReason LockReason `json:"lock_reason"`

	HideLockReason bool `json:"hide_lock_reason"`

	Locality string `json:"locality"`
}

// DeleteResourceRequest: delete resource request.
type DeleteResourceRequest struct {
	ResourceID string `json:"resource_id"`

	// ResourceType: default value: unknown
	ResourceType ResourceType `json:"resource_type"`

	// Locality: default value: unknown_locality
	Locality resource_v1alpha1.ResourceLocality `json:"locality"`
}

// DeleteResourceResponse: delete resource response.
type DeleteResourceResponse struct {
}

// DeleteResourcesRequest: delete resources request.
type DeleteResourcesRequest struct {
	OrganizationID string `json:"organization_id"`
}

// DeleteResourcesResponse: delete resources response.
type DeleteResourcesResponse struct {
}

// HasResourcesRequest: has resources request.
type HasResourcesRequest struct {
	// Precisely one of OrganizationID, ProjectID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// Precisely one of OrganizationID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`
}

// HasResourcesResponse: has resources response.
type HasResourcesResponse struct {
	HasResources bool `json:"has_resources"`
}

// LockResourceRequest: lock resource request.
type LockResourceRequest struct {
	ResourceID string `json:"resource_id"`

	// ResourceType: default value: unknown
	ResourceType ResourceType `json:"resource_type"`

	// Reason: default value: unknown
	Reason LockReason `json:"reason"`

	HideReason bool `json:"hide_reason"`

	// Locality: default value: unknown_locality
	Locality resource_v1alpha1.ResourceLocality `json:"locality"`
}

// LockResourceResponse: lock resource response.
type LockResourceResponse struct {
}

// LockResourcesRequest: lock resources request.
type LockResourcesRequest struct {
	OrganizationID string `json:"organization_id"`

	// Reason: default value: unknown
	Reason LockReason `json:"reason"`

	HideReason bool `json:"hide_reason"`
}

// LockResourcesResponse: lock resources response.
type LockResourcesResponse struct {
}

// LookupRequest: lookup request.
type LookupRequest struct {
	// Precisely one of IP, URL must be set.
	IP *string `json:"ip,omitempty"`

	// Precisely one of IP, URL must be set.
	URL *string `json:"url,omitempty"`

	Date *time.Time `json:"date,omitempty"`
}

// LookupResourcesRequest: lookup resources request.
type LookupResourcesRequest struct {
	// Precisely one of IP, URL must be set.
	IP *string `json:"ip,omitempty"`

	// Precisely one of IP, URL must be set.
	URL *string `json:"url,omitempty"`

	Date *time.Time `json:"date,omitempty"`
}

// LookupResourcesResponse: lookup resources response.
type LookupResourcesResponse struct {
	Resources []*Resource `json:"resources"`
}

// ShutdownResourceRequest: shutdown resource request.
type ShutdownResourceRequest struct {
	ResourceID string `json:"resource_id"`

	// ResourceType: default value: unknown
	ResourceType ResourceType `json:"resource_type"`

	// Locality: default value: unknown_locality
	Locality resource_v1alpha1.ResourceLocality `json:"locality"`
}

// ShutdownResourceResponse: shutdown resource response.
type ShutdownResourceResponse struct {
}

// ShutdownResourcesRequest: shutdown resources request.
type ShutdownResourcesRequest struct {
	OrganizationID string `json:"organization_id"`
}

// ShutdownResourcesResponse: shutdown resources response.
type ShutdownResourcesResponse struct {
}

// UnlockResourceRequest: unlock resource request.
type UnlockResourceRequest struct {
	ResourceID string `json:"resource_id"`

	// ResourceType: default value: unknown
	ResourceType ResourceType `json:"resource_type"`

	// Locality: default value: unknown_locality
	Locality resource_v1alpha1.ResourceLocality `json:"locality"`
}

// UnlockResourceResponse: unlock resource response.
type UnlockResourceResponse struct {
}

// UnlockResourcesRequest: unlock resources request.
type UnlockResourcesRequest struct {
	OrganizationID string `json:"organization_id"`
}

// UnlockResourcesResponse: unlock resources response.
type UnlockResourcesResponse struct {
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

// Lookup:
func (s *API) Lookup(req *LookupRequest, opts ...scw.RequestOption) (*Resource, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1/lookup",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Resource

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LookupResources:
func (s *API) LookupResources(req *LookupResourcesRequest, opts ...scw.RequestOption) (*LookupResourcesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1/lookup-resources",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LookupResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// HasResources:
func (s *API) HasResources(req *HasResourcesRequest, opts ...scw.RequestOption) (*HasResourcesResponse, error) {
	var err error

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1/has-resources",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp HasResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LockResource:
func (s *API) LockResource(req *LockResourceRequest, opts ...scw.RequestOption) (*LockResourceResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1/lock-resource",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LockResourceResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LockResources:
func (s *API) LockResources(req *LockResourcesRequest, opts ...scw.RequestOption) (*LockResourcesResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1/lock-resources",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LockResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnlockResource:
func (s *API) UnlockResource(req *UnlockResourceRequest, opts ...scw.RequestOption) (*UnlockResourceResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1/unlock-resource",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UnlockResourceResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnlockResources:
func (s *API) UnlockResources(req *UnlockResourcesRequest, opts ...scw.RequestOption) (*UnlockResourcesResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1/unlock-resources",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UnlockResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShutdownResource:
func (s *API) ShutdownResource(req *ShutdownResourceRequest, opts ...scw.RequestOption) (*ShutdownResourceResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1/shutdown-resource",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ShutdownResourceResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ShutdownResources:
func (s *API) ShutdownResources(req *ShutdownResourcesRequest, opts ...scw.RequestOption) (*ShutdownResourcesResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1/shutdown-resources",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ShutdownResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteResource:
func (s *API) DeleteResource(req *DeleteResourceRequest, opts ...scw.RequestOption) (*DeleteResourceResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1/delete-resource",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DeleteResourceResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteResources:
func (s *API) DeleteResources(req *DeleteResourcesRequest, opts ...scw.RequestOption) (*DeleteResourcesResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1/delete-resources",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DeleteResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

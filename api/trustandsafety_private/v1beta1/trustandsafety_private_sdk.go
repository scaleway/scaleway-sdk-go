// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package trustandsafety_private provides methods and message types of the trustandsafety_private v1beta1 API.
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
	ResourceTypeBaremetalPrivateNetworkMember = ResourceType("baremetal_private_network_member")
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

type SnapshotStatus string

const (
	SnapshotStatusReady    = SnapshotStatus("ready")
	SnapshotStatusCreating = SnapshotStatus("creating")
	SnapshotStatusFailed   = SnapshotStatus("failed")
)

func (enum SnapshotStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "ready"
	}
	return string(enum)
}

func (enum SnapshotStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SnapshotStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SnapshotStatus(SnapshotStatus(tmp).String())
	return nil
}

// Resource: resource.
type Resource struct {
	// ID: the ID of the resource.
	ID string `json:"id"`

	// Type: the type of the resource.
	// Default value: unknown
	Type ResourceType `json:"type"`

	// ProjectID: the project ID of the resource.
	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// OrganizationID: the organization ID of the resource.
	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// Status: the status of the resource.
	// Default value: unlocked
	Status ResourceStatus `json:"status"`

	// LockReason: the lock reason of the resource.
	// Default value: unknown
	LockReason LockReason `json:"lock_reason"`

	// HideLockReason: indicates whether the lock should be known to the user.
	HideLockReason bool `json:"hide_lock_reason"`
}

// DeleteResourceResponse: delete resource response.
type DeleteResourceResponse struct {
}

// DeleteResourcesResponse: delete resources response.
type DeleteResourcesResponse struct {
}

// ListResourcesResponse: list resources response.
type ListResourcesResponse struct {
	// TotalCount: the total number of resources.
	TotalCount uint32 `json:"total_count"`

	// Resources: the paginated returned cluster.
	Resources []*Resource `json:"resources"`
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

// LockResourceResponse: lock resource response.
type LockResourceResponse struct {
}

// LockResourcesResponse: lock resources response.
type LockResourcesResponse struct {
}

// LookupResponse: lookup response.
type LookupResponse struct {
	// Resources: the returned resources.
	Resources []*Resource `json:"resources"`
}

// ProductAPICreateSnapshotRequest: product api create snapshot request.
type ProductAPICreateSnapshotRequest struct {
	// ResourceID: the ID of the resource that will be snapshotted.
	ResourceID string `json:"resource_id"`

	// ResourceType: the type of the resource.
	// Default value: unknown
	ResourceType ResourceType `json:"resource_type"`
}

// ProductAPIDeleteResourceRequest: product api delete resource request.
type ProductAPIDeleteResourceRequest struct {
	// ResourceID: the ID of the resource that will be deleted.
	ResourceID string `json:"resource_id"`

	// ResourceType: the type of the resource.
	// Default value: unknown
	ResourceType ResourceType `json:"resource_type"`
}

// ProductAPIDeleteResourcesRequest: product api delete resources request.
type ProductAPIDeleteResourcesRequest struct {
	// OrganizationID: the organization ID of the resource.
	// Precisely one of OrganizationID, ProjectID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// ProjectID: the project ID of the resource.
	// Precisely one of OrganizationID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`
}

// ProductAPIGetSnapshotRequest: product api get snapshot request.
type ProductAPIGetSnapshotRequest struct {
	// SnapshotID: the ID of the snapshot.
	SnapshotID string `json:"snapshot_id"`
}

// ProductAPIListResourcesRequest: product api list resources request.
type ProductAPIListResourcesRequest struct {
	// ProjectID: filter resources by project ID.
	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// OrganizationID: filter resources by organization ID.
	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// Locked: filter resources by locked status.
	Locked *bool `json:"locked,omitempty"`

	// Page: the page number for the returned resources.
	Page *int32 `json:"page,omitempty"`

	// PageSize: the maximum number of resources per page.
	PageSize *uint32 `json:"page_size,omitempty"`
}

// ProductAPILockResourceRequest: product api lock resource request.
type ProductAPILockResourceRequest struct {
	// ResourceID: the ID of the resource that will be locked.
	ResourceID string `json:"resource_id"`

	// ResourceType: the type of the resource.
	// Default value: unknown
	ResourceType ResourceType `json:"resource_type"`

	// HideReason: indicates whether the reason for locking the resource should be hidden to the user.
	HideReason bool `json:"hide_reason"`

	// Reason: the reason for locking the resource.
	// Default value: unknown
	Reason LockReason `json:"reason"`
}

// ProductAPILockResourcesRequest: product api lock resources request.
type ProductAPILockResourcesRequest struct {
	// OrganizationID: the organization ID of the resource.
	// Precisely one of OrganizationID, ProjectID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// ProjectID: the project ID of the resource.
	// Precisely one of OrganizationID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// HideReason: indicates whether the reason for locking the resource should be hidden to the user.
	HideReason bool `json:"hide_reason"`

	// Reason: the reason for locking the resource.
	// Default value: unknown
	Reason LockReason `json:"reason"`
}

// ProductAPILookupRequest: product api lookup request.
type ProductAPILookupRequest struct {
	// IP: filter resources by IP.
	// Precisely one of IP, URL, ResourceID must be set.
	IP *string `json:"ip,omitempty"`

	// URL: filter resources by URL.
	// Precisely one of IP, URL, ResourceID must be set.
	URL *string `json:"url,omitempty"`

	// ResourceID: filter resources by resource ID.
	// Precisely one of IP, URL, ResourceID must be set.
	ResourceID *string `json:"resource_id,omitempty"`

	// Date: filter resources by date.
	Date *time.Time `json:"date,omitempty"`
}

// ProductAPIShutdownResourceRequest: product api shutdown resource request.
type ProductAPIShutdownResourceRequest struct {
	// ResourceID: the ID of the resource that will be shutdown.
	ResourceID string `json:"resource_id"`

	// ResourceType: the type of the resource.
	// Default value: unknown
	ResourceType ResourceType `json:"resource_type"`
}

// ProductAPIShutdownResourcesRequest: product api shutdown resources request.
type ProductAPIShutdownResourcesRequest struct {
	// OrganizationID: the organization ID of the resource.
	// Precisely one of OrganizationID, ProjectID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// ProjectID: the project ID of the resource.
	// Precisely one of OrganizationID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`
}

// ProductAPIUnlockResourceRequest: product api unlock resource request.
type ProductAPIUnlockResourceRequest struct {
	// ResourceID: the ID of the resource that will be unlocked.
	ResourceID string `json:"resource_id"`

	// ResourceType: the type of the resource.
	// Default value: unknown
	ResourceType ResourceType `json:"resource_type"`
}

// ProductAPIUnlockResourcesRequest: product api unlock resources request.
type ProductAPIUnlockResourcesRequest struct {
	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`
}

// ShutdownResourceResponse: shutdown resource response.
type ShutdownResourceResponse struct {
}

// ShutdownResourcesResponse: shutdown resources response.
type ShutdownResourcesResponse struct {
}

// Snapshot: snapshot.
type Snapshot struct {
	// ID: display the snapshot ID.
	ID string `json:"id"`

	// URL: display the snapshot URL.
	URL string `json:"url"`

	// Status: display the status of the snapshot.
	// Default value: ready
	Status SnapshotStatus `json:"status"`
}

// UnlockResourceResponse: unlock resource response.
type UnlockResourceResponse struct {
}

// UnlockResourcesResponse: unlock resources response.
type UnlockResourcesResponse struct {
}

// This service should be implemented by services to comply with T&S requirements.
type ProductAPI struct {
	client *scw.Client
}

// NewProductAPI returns a ProductAPI object from a Scaleway client.
func NewProductAPI(client *scw.Client) *ProductAPI {
	return &ProductAPI{
		client: client,
	}
}

// Lookup: Used by worker-abuse to find where suspicious IPs / URLs come from.
func (s *ProductAPI) Lookup(req *ProductAPILookupRequest, opts ...scw.RequestOption) (*LookupResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1beta1/lookup",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LookupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListResources: This method is used for 2 cases: project deletion, resource unlocking.
//
// ## Project deletion.
//
// Before deleting a project from account API, all products in all localities will receive a ListResources call in order to assert no resource are left on a deleted project.
//
// NB: this is a first iteration of project deletion. Next iteration could be an asynchronous deletion of resources.
//
// ## Resource unlocking.
//
// Account workflow relies on ListResources and the implementation of the `locked` filter to unlock a resource.
// Its process is the following:
// - List all locked resources accross all products.
// - Look for recent unlocked organizations in db_world.
// - Call UnlockResource one by one when some resources are left unlocked when they shouldn't.
func (s *ProductAPI) ListResources(req *ProductAPIListResourcesRequest, opts ...scw.RequestOption) (*ListResourcesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1beta1/list-resources",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ListResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LockResource: Can be used in the admin console to lock a specific resource with a specific reason.
func (s *ProductAPI) LockResource(req *ProductAPILockResourceRequest, opts ...scw.RequestOption) (*LockResourceResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1beta1/lock-resource",
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

// LockResources: Used by account workflow to lock resources on an organization.
func (s *ProductAPI) LockResources(req *ProductAPILockResourcesRequest, opts ...scw.RequestOption) (*LockResourcesResponse, error) {
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
		Path:   "/trustandsafety-private/v1beta1/lock-resources",
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

// UnlockResource: Cf. see ListResources documentation.
func (s *ProductAPI) UnlockResource(req *ProductAPIUnlockResourceRequest, opts ...scw.RequestOption) (*UnlockResourceResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1beta1/unlock-resource",
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

// UnlockResources: Not used yet.
func (s *ProductAPI) UnlockResources(req *ProductAPIUnlockResourcesRequest, opts ...scw.RequestOption) (*UnlockResourcesResponse, error) {
	var err error

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1beta1/unlock-resources",
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

// DeleteResource: Used by account workflow to delete resources of a deleted organization.
func (s *ProductAPI) DeleteResource(req *ProductAPIDeleteResourceRequest, opts ...scw.RequestOption) (*DeleteResourceResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1beta1/delete-resource",
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

// DeleteResources: Used by account workflow to delete resources of a deleted organization.
func (s *ProductAPI) DeleteResources(req *ProductAPIDeleteResourcesRequest, opts ...scw.RequestOption) (*DeleteResourcesResponse, error) {
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
		Path:   "/trustandsafety-private/v1beta1/delete-resources",
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

// ShutdownResource: Instance-specific method.
func (s *ProductAPI) ShutdownResource(req *ProductAPIShutdownResourceRequest, opts ...scw.RequestOption) (*ShutdownResourceResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1beta1/shutdown-resource",
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
func (s *ProductAPI) ShutdownResources(req *ProductAPIShutdownResourcesRequest, opts ...scw.RequestOption) (*ShutdownResourcesResponse, error) {
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
		Path:   "/trustandsafety-private/v1beta1/shutdown-resources",
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

// CreateSnapshot: Instance-specific method used when the authorities ask for a data export of a client's server (rare).
//
// This could be implemented for other data-related products.
func (s *ProductAPI) CreateSnapshot(req *ProductAPICreateSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1beta1/create-snapshot",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSnapshot: Instance-specific method used when the authorities ask for a data export of a client's server (rare).
//
// This could be implemented for other data-related products.
func (s *ProductAPI) GetSnapshot(req *ProductAPIGetSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/trustandsafety-private/v1beta1/get-snapshot",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

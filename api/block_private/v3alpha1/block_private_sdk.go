// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package block_private provides methods and message types of the block_private v3alpha1 API.
package block_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	block_v1alpha1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/block/v1alpha1"
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

type ResourceType string

const (
	// If unspecified, the ResourceType is unknown by default.
	ResourceTypeUnknownType = ResourceType("unknown_type")
	// Resource of type Volume.
	ResourceTypeVolume = ResourceType("volume")
	// Resource of type Snapshot.
	ResourceTypeSnapshot = ResourceType("snapshot")
)

func (enum ResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
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

// ReferenceWithHandle: reference with handle.
type ReferenceWithHandle struct {
	Reference *block_v1alpha1.Reference `json:"reference"`

	HandleID string `json:"handle_id"`
}

// SnapshotLink: snapshot link.
type SnapshotLink struct {
	CephSnapshotName string `json:"ceph_snapshot_name"`

	DisplayName string `json:"display_name"`

	Size uint64 `json:"size"`

	ProjectID string `json:"project_id"`

	CreatedAt *time.Time `json:"created_at"`

	References []*block_v1alpha1.Reference `json:"references"`

	SnapshotID string `json:"snapshot_id"`

	DeletedAt *time.Time `json:"deleted_at"`

	Tags []string `json:"tags"`

	OrganizationID string `json:"organization_id"`
}

// Usage: usage.
type Usage struct {
	StartedAt *time.Time `json:"started_at"`

	Size uint64 `json:"size"`
}

// LinkedSnapshot: linked snapshot.
type LinkedSnapshot struct {
	Snapshot *block_v1alpha1.Snapshot `json:"snapshot"`

	References []*ReferenceWithHandle `json:"references"`
}

// LinkedVolume: linked volume.
type LinkedVolume struct {
	Volume *block_v1alpha1.Volume `json:"volume"`

	Reference *ReferenceWithHandle `json:"reference"`
}

// Cluster: cluster.
type Cluster struct {
	// ID: UUID of the Cluster.
	ID string `json:"id"`

	// Name: name of the Ceph cluster.
	Name string `json:"name"`

	// Hwgen: hardware Generation description of the cluster.
	Hwgen string `json:"hwgen"`

	// HandledVolumeTypes: list of the volume-type available on the cluster.
	HandledVolumeTypes []string `json:"handled_volume_types"`
}

// VolumeTypeUsage: volume type usage.
type VolumeTypeUsage struct {
	// Type: name for this volume type.
	Type string `json:"type"`

	// Class: storage class for this volume type.
	// Default value: unknown_storage_class
	Class block_v1alpha1.StorageClass `json:"class"`

	// VolumeCount: total number of volumes for this volume type.
	VolumeCount uint32 `json:"volume_count"`

	// VolumeSize: total size used (in bytes) for this volume type.
	VolumeSize uint64 `json:"volume_size"`
}

// CreateReferenceRequest: create reference request.
type CreateReferenceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ResourceType: default value: unknown_type
	ResourceType ResourceType `json:"resource_type"`

	ResourceID string `json:"resource_id"`

	ProductResourceType string `json:"product_resource_type"`

	ProductResourceID string `json:"product_resource_id"`

	// ReferenceType: default value: unknown_type
	ReferenceType block_v1alpha1.ReferenceType `json:"reference_type"`

	AuthID string `json:"auth_id"`
}

// CreateSnapshotWithReferenceRequest: create snapshot with reference request.
type CreateSnapshotWithReferenceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SnapshotRequest *block_v1alpha1.CreateSnapshotRequest `json:"snapshot_request,omitempty"`

	ProductResourceType string `json:"product_resource_type"`

	ProductResourceID string `json:"product_resource_id"`

	// ReferenceType: default value: unknown_type
	ReferenceType block_v1alpha1.ReferenceType `json:"reference_type"`

	AuthID string `json:"auth_id"`
}

// CreateSnapshotWithReferenceResponse: create snapshot with reference response.
type CreateSnapshotWithReferenceResponse struct {
	Snapshot *block_v1alpha1.Snapshot `json:"snapshot"`

	Reference *ReferenceWithHandle `json:"reference"`
}

// CreateVolumeOnClusterRequest: create volume on cluster request.
type CreateVolumeOnClusterRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ClusterID string `json:"-"`

	VolumeRequest *block_v1alpha1.CreateVolumeRequest `json:"volume_request,omitempty"`
}

// CreateVolumeWithReferenceRequest: create volume with reference request.
type CreateVolumeWithReferenceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	VolumeRequest *block_v1alpha1.CreateVolumeRequest `json:"volume_request,omitempty"`

	ProductResourceType string `json:"product_resource_type"`

	ProductResourceID string `json:"product_resource_id"`

	// ReferenceType: default value: unknown_type
	ReferenceType block_v1alpha1.ReferenceType `json:"reference_type"`

	AuthID string `json:"auth_id"`

	SnapshotIsPublic bool `json:"snapshot_is_public"`
}

// CreateVolumeWithReferenceResponse: create volume with reference response.
type CreateVolumeWithReferenceResponse struct {
	Volume *block_v1alpha1.Volume `json:"volume"`

	Reference *ReferenceWithHandle `json:"reference"`
}

// Credentials: credentials.
type Credentials struct {
	ClusterName string `json:"cluster_name"`

	ResourcePath string `json:"resource_path"`

	ReadOnly bool `json:"read_only"`

	CephFsid string `json:"ceph_fsid"`

	Monitors []string `json:"monitors"`

	MonitorsIPv4 []string `json:"monitors_ipv4"`

	MonitorsIPv6 []string `json:"monitors_ipv6"`

	PublicCidrv4 []string `json:"public_cidrv4"`

	PublicCidrv6 []string `json:"public_cidrv6"`

	ClientID string `json:"client_id"`

	Keyring string `json:"keyring"`

	ReadIopsLimit uint32 `json:"read_iops_limit"`

	WriteIopsLimit uint32 `json:"write_iops_limit"`

	MixedIopsLimit uint32 `json:"mixed_iops_limit"`
}

// DeleteReferenceRequest: delete reference request.
type DeleteReferenceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	HandleID string `json:"-"`
}

// GetReferenceCredentialsRequest: get reference credentials request.
type GetReferenceCredentialsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// HandleID: UUID of the Handle the reference belongs to.
	HandleID string `json:"-"`
}

// GetUsageSummaryRequest: get usage summary request.
type GetUsageSummaryRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`
}

// LinkVolumeAndSnapshotsDryRunRequest: link volume and snapshots dry run request.
type LinkVolumeAndSnapshotsDryRunRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// CephClusterName: name of the Ceph cluster.
	CephClusterName string `json:"ceph_cluster_name"`

	// CephPoolName: name of the Ceph pool.
	CephPoolName string `json:"ceph_pool_name"`

	// CephNamespaceName: name of the Ceph namespace.
	CephNamespaceName string `json:"ceph_namespace_name"`

	// CephVolumeName: name of the Ceph volume.
	CephVolumeName string `json:"ceph_volume_name"`

	// DisplayName: displayed name in Scaleway console.
	DisplayName string `json:"display_name"`

	// Usages: usage history for this volume.
	Usages []*Usage `json:"usages"`

	// ProjectID: project linked to the volume.
	ProjectID string `json:"project_id"`

	// Reference: list of references.
	Reference *block_v1alpha1.Reference `json:"reference,omitempty"`

	// VolumeID: volume UUID known in database.
	VolumeID string `json:"volume_id"`

	// DeletedAt: date of deletion by the customer.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// Snapshots: list of snapshots created from the volume.
	Snapshots []*SnapshotLink `json:"snapshots"`

	// ParentID: snapshot UUID when volume created from a snapshot.
	ParentID *string `json:"parent_id,omitempty"`

	// Tags: list of tags assigned to the volume.
	Tags []string `json:"tags"`

	// OrganizationID: organization linked to the volume.
	OrganizationID string `json:"organization_id"`
}

// LinkVolumeAndSnapshotsRequest: link volume and snapshots request.
type LinkVolumeAndSnapshotsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// CephClusterName: name of the Ceph cluster.
	CephClusterName string `json:"ceph_cluster_name"`

	// CephPoolName: name of the Ceph pool.
	CephPoolName string `json:"ceph_pool_name"`

	// CephNamespaceName: name of the Ceph namespace.
	CephNamespaceName string `json:"ceph_namespace_name"`

	// CephVolumeName: name of the Ceph volume.
	CephVolumeName string `json:"ceph_volume_name"`

	// DisplayName: displayed name in Scaleway console.
	DisplayName string `json:"display_name"`

	// Usages: usage history for this volume.
	Usages []*Usage `json:"usages"`

	// ProjectID: project linked to the volume.
	ProjectID string `json:"project_id"`

	// Reference: list of references.
	Reference *block_v1alpha1.Reference `json:"reference,omitempty"`

	// VolumeID: volume UUID known in database.
	VolumeID string `json:"volume_id"`

	// DeletedAt: date of deletion by the customer.
	DeletedAt *time.Time `json:"deleted_at,omitempty"`

	// Snapshots: list of snapshots created from the volume.
	Snapshots []*SnapshotLink `json:"snapshots"`

	// ParentID: snapshot UUID when volume created from a snapshot.
	ParentID *string `json:"parent_id,omitempty"`

	// Tags: list of tags assigned to the volume.
	Tags []string `json:"tags"`

	// OrganizationID: organization linked to the volume.
	OrganizationID string `json:"organization_id"`
}

// LinkVolumeAndSnapshotsResponse: link volume and snapshots response.
type LinkVolumeAndSnapshotsResponse struct {
	// LinkedVolume: volume to be linked.
	LinkedVolume *LinkedVolume `json:"linked_volume"`

	// LinkedSnapshots: list of snapshots linked.
	LinkedSnapshots []*LinkedSnapshot `json:"linked_snapshots"`
}

// ListCLustersResponse: list c lusters response.
type ListCLustersResponse struct {
	Clusters []*Cluster `json:"clusters"`
}

// ListClustersRequest: list clusters request.
type ListClustersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// UpdateReferenceRequest: update reference request.
type UpdateReferenceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	HandleID string `json:"-"`

	// Status: default value: unknown_status
	Status block_v1alpha1.ReferenceStatus `json:"status"`
}

// UsageSummary: usage summary.
type UsageSummary struct {
	// TotalVolumeCount: total number of block volumes in the zone.
	TotalVolumeCount uint32 `json:"total_volume_count"`

	// TotalVolumeSize: total size used (in bytes) for block volumes.
	TotalVolumeSize uint64 `json:"total_volume_size"`

	// TotalSnapshotCount: total number of snapshots in the zone.
	TotalSnapshotCount uint32 `json:"total_snapshot_count"`

	// TotalSnapshotSize: total size used (in bytes) for snapshots.
	TotalSnapshotSize uint64 `json:"total_snapshot_size"`

	// VolumeTypes: repartition of the volume usage per volume type.
	VolumeTypes []*VolumeTypeUsage `json:"volume_types"`
}

// API Block is used mainly by Instance and BMaaS to manage the Block volumes and snapshots.
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

// CreateReference: A reference can be created only by Scaleway products such as Instance or BMaaS.
// It is usually similar for a Product point of view as a "Attach" operation.
// Once a reference is created, a HandleID is provided which is a kind of secret provided to the Product consumer to identify and manage the reference.
func (s *API) CreateReference(req *CreateReferenceRequest, opts ...scw.RequestOption) (*ReferenceWithHandle, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/block-private/v3alpha1/zones/" + fmt.Sprint(req.Zone) + "/references",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ReferenceWithHandle

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteReference: It is usually similar for a Product point of view as a "Detach" operation.
func (s *API) DeleteReference(req *DeleteReferenceRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.HandleID) == "" {
		return errors.New("field HandleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/block-private/v3alpha1/zones/" + fmt.Sprint(req.Zone) + "/references/" + fmt.Sprint(req.HandleID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateReference: This operation is going to be deprecated.
func (s *API) UpdateReference(req *UpdateReferenceRequest, opts ...scw.RequestOption) (*block_v1alpha1.Reference, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.HandleID) == "" {
		return nil, errors.New("field HandleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/block-private/v3alpha1/zones/" + fmt.Sprint(req.Zone) + "/references/" + fmt.Sprint(req.HandleID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp block_v1alpha1.Reference

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetReferenceCredentials: Should be used by Instance or Block team when mounting a volume on a server.
func (s *API) GetReferenceCredentials(req *GetReferenceCredentialsRequest, opts ...scw.RequestOption) (*Credentials, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.HandleID) == "" {
		return nil, errors.New("field HandleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block-private/v3alpha1/zones/" + fmt.Sprint(req.Zone) + "/references/" + fmt.Sprint(req.HandleID) + "/credentials",
	}

	var resp Credentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LinkVolumeAndSnapshots: Act as an import feature, this method should be used only by Instance team for legacy volumes.
// When this query is successful, the legacy data owned by legacy Instance DB should be erased.
func (s *API) LinkVolumeAndSnapshots(req *LinkVolumeAndSnapshotsRequest, opts ...scw.RequestOption) (*LinkVolumeAndSnapshotsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/block-private/v3alpha1/zones/" + fmt.Sprint(req.Zone) + "/link-volume",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LinkVolumeAndSnapshotsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LinkVolumeAndSnapshotsDryRun: Should be used by Instance team to check quota, content and format of the query and so on.
func (s *API) LinkVolumeAndSnapshotsDryRun(req *LinkVolumeAndSnapshotsDryRunRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/block-private/v3alpha1/zones/" + fmt.Sprint(req.Zone) + "/link-volume-dry-run",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateVolumeOnCluster: This feature is needed by DBaaS team when they have to ensure some resiliency.
func (s *API) CreateVolumeOnCluster(req *CreateVolumeOnClusterRequest, opts ...scw.RequestOption) (*block_v1alpha1.Volume, error) {
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
		Path:   "/block-private/v3alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/volumes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp block_v1alpha1.Volume

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVolumeWithReference: This feature is needed by Instance team when creating a new Instance with no Local Storage.
func (s *API) CreateVolumeWithReference(req *CreateVolumeWithReferenceRequest, opts ...scw.RequestOption) (*CreateVolumeWithReferenceResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/block-private/v3alpha1/zones/" + fmt.Sprint(req.Zone) + "/volumes-with-reference",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateVolumeWithReferenceResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSnapshotWithReference:
func (s *API) CreateSnapshotWithReference(req *CreateSnapshotWithReferenceRequest, opts ...scw.RequestOption) (*CreateSnapshotWithReferenceResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/block-private/v3alpha1/zones/" + fmt.Sprint(req.Zone) + "/snapshots-with-reference",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateSnapshotWithReferenceResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUsageSummary: It contains the number and the allocated size of the snapshots, the number and the allocated size of the volumes.
// This feature is needed by Instance team to detect the current usage of an organization or a project, during the migration period.
func (s *API) GetUsageSummary(req *GetUsageSummaryRequest, opts ...scw.RequestOption) (*UsageSummary, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block-private/v3alpha1/zones/" + fmt.Sprint(req.Zone) + "/usage",
		Query:  query,
	}

	var resp UsageSummary

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListClusters: List all the Ceph clusters manageable through this SBS API.
func (s *API) ListClusters(req *ListClustersRequest, opts ...scw.RequestOption) (*ListCLustersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block-private/v3alpha1/zones/" + fmt.Sprint(req.Zone) + "/clusters",
	}

	var resp ListCLustersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

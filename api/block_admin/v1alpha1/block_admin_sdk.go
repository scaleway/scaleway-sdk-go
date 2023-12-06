// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package block_admin provides methods and message types of the block_admin v1alpha1 API.
package block_admin

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
	block_v1alpha1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/block/v1alpha1"
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

type ListAuditLogsRequestOrderBy string

const (
	ListAuditLogsRequestOrderByCreatedAtAsc  = ListAuditLogsRequestOrderBy("created_at_asc")
	ListAuditLogsRequestOrderByCreatedAtDesc = ListAuditLogsRequestOrderBy("created_at_desc")
)

func (enum ListAuditLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListAuditLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListAuditLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListAuditLogsRequestOrderBy(ListAuditLogsRequestOrderBy(tmp).String())
	return nil
}

type ListSnapshotsRequestOrderBy string

const (
	ListSnapshotsRequestOrderByCreatedAtAsc  = ListSnapshotsRequestOrderBy("created_at_asc")
	ListSnapshotsRequestOrderByCreatedAtDesc = ListSnapshotsRequestOrderBy("created_at_desc")
	ListSnapshotsRequestOrderByNameAsc       = ListSnapshotsRequestOrderBy("name_asc")
	ListSnapshotsRequestOrderByNameDesc      = ListSnapshotsRequestOrderBy("name_desc")
)

func (enum ListSnapshotsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSnapshotsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSnapshotsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSnapshotsRequestOrderBy(ListSnapshotsRequestOrderBy(tmp).String())
	return nil
}

type ListVolumesRequestOrderBy string

const (
	ListVolumesRequestOrderByCreatedAtAsc  = ListVolumesRequestOrderBy("created_at_asc")
	ListVolumesRequestOrderByCreatedAtDesc = ListVolumesRequestOrderBy("created_at_desc")
	ListVolumesRequestOrderByNameAsc       = ListVolumesRequestOrderBy("name_asc")
	ListVolumesRequestOrderByNameDesc      = ListVolumesRequestOrderBy("name_desc")
)

func (enum ListVolumesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListVolumesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListVolumesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListVolumesRequestOrderBy(ListVolumesRequestOrderBy(tmp).String())
	return nil
}

// SnapshotParentVolume: snapshot parent volume.
type SnapshotParentVolume struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Type string `json:"type"`

	// Status: default value: unknown_status
	Status block_v1alpha1.VolumeStatus `json:"status"`
}

// VolumeSpecifications: volume specifications.
type VolumeSpecifications struct {
	// PerfIops: the maximum IO/s expected, according to the different options available in stock (`5000 | 15000`).
	PerfIops *uint32 `json:"perf_iops"`

	// Class: the storage class of the volume.
	// Default value: unknown_storage_class
	Class block_v1alpha1.StorageClass `json:"class"`
}

// AuditLog: audit log.
type AuditLog struct {
	Action string `json:"action"`

	CreatedAt *time.Time `json:"created_at"`

	Detail *scw.JSONObject `json:"detail"`

	ID string `json:"id"`
}

// SnapshotSummary: snapshot summary.
type SnapshotSummary struct {
	// ID: UUID of the snapshot.
	ID string `json:"id"`

	// Name: name of the snapshot.
	Name string `json:"name"`

	// ParentVolume: if the parent volume has been deleted, value is null.
	ParentVolume *SnapshotParentVolume `json:"parent_volume"`

	// Size: size of the snapshot in bytes.
	Size scw.Size `json:"size"`

	// ProjectID: UUID of the project the snapshot belongs to.
	ProjectID string `json:"project_id"`

	// CreatedAt: creation date of the snapshot.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the properties of a snapshot.
	UpdatedAt *time.Time `json:"updated_at"`

	// Status: current status of the snapshot (available, in_use, ...).
	// Default value: unknown_status
	Status block_v1alpha1.SnapshotStatus `json:"status"`

	// Tags: list of tags assigned to the volume.
	Tags []string `json:"tags"`

	// Zone: snapshot Availability Zone.
	Zone scw.Zone `json:"zone"`

	// Class: storage class of the snapshot.
	// Default value: unknown_storage_class
	Class block_v1alpha1.StorageClass `json:"class"`

	OrganizationID string `json:"organization_id"`

	CephName *string `json:"ceph_name"`

	// Namespace: ceph namespace of the volume.
	Namespace string `json:"namespace"`

	// Pool: ceph pool of the volume.
	Pool string `json:"pool"`

	UsedSize *scw.Size `json:"used_size"`

	DeletedAt *time.Time `json:"deleted_at"`

	PurgedAt *time.Time `json:"purged_at"`
}

// VolumeType: volume type.
type VolumeType struct {
	// Type: volume type.
	Type string `json:"type"`

	// Pricing: price of the volume billed in GB/hour.
	Pricing *scw.Money `json:"pricing"`

	// SnapshotPricing: price of the snapshot billed in GB/hour.
	SnapshotPricing *scw.Money `json:"snapshot_pricing"`

	// Specs: volume specifications of the volume type.
	Specs *VolumeSpecifications `json:"specs"`
}

// VolumeSummary: volume summary.
type VolumeSummary struct {
	CephName *string `json:"ceph_name"`

	CreatedAt *time.Time `json:"created_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	ID string `json:"id"`

	Name string `json:"name"`

	OrganizationID string `json:"organization_id"`

	ParentSnapshotID *string `json:"parent_snapshot_id"`

	ProjectID string `json:"project_id"`

	PurgedAt *time.Time `json:"purged_at"`

	Size scw.Size `json:"size"`

	Specs *VolumeSpecifications `json:"specs"`

	// Status: default value: unknown_status
	Status block_v1alpha1.VolumeStatus `json:"status"`

	Tags []string `json:"tags"`

	Type string `json:"type"`

	UpdatedAt *time.Time `json:"updated_at"`

	UsedSize *scw.Size `json:"used_size"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`

	Namespace string `json:"namespace"`

	Pool string `json:"pool"`
}

// Reference: reference.
type Reference struct {
	// ID: UUID of the reference.
	ID string `json:"id"`

	// ProductResourceType: type of resoruce to which the reference is associated (snapshot or volume).
	ProductResourceType string `json:"product_resource_type"`

	// ProductResourceID: UUID of the volume or the snapshot it refers to (according to the product_resource_type).
	ProductResourceID string `json:"product_resource_id"`

	// CreatedAt: creation date of the reference.
	CreatedAt *time.Time `json:"created_at"`

	// Type: type of reference (link, exclusive, read_only).
	// Default value: unknown_type
	Type block_v1alpha1.ReferenceType `json:"type"`

	// Status: status of reference (attaching, attached, detaching).
	// Default value: unknown_status
	Status block_v1alpha1.ReferenceStatus `json:"status"`
}

// DeleteSnapshotRequest: delete snapshot request.
type DeleteSnapshotRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// SnapshotID: UUID of the snapshot.
	SnapshotID string `json:"-"`
}

// DeleteVolumeRequest: delete volume request.
type DeleteVolumeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// VolumeID: UUID of the volume.
	VolumeID string `json:"-"`
}

// GetSnapshotRequest: get snapshot request.
type GetSnapshotRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// SnapshotID: UUID of the snapshot.
	SnapshotID string `json:"-"`
}

// GetVolumeRequest: get volume request.
type GetVolumeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// VolumeID: UUID of the volume.
	VolumeID string `json:"-"`
}

// ListAuditLogsRequest: list audit logs request.
type ListAuditLogsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ResourceID string `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListAuditLogsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListAuditLogsResponse: list audit logs response.
type ListAuditLogsResponse struct {
	AuditLogs []*AuditLog `json:"audit_logs"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListAuditLogsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListAuditLogsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListAuditLogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.AuditLogs = append(r.AuditLogs, results.AuditLogs...)
	r.TotalCount += uint64(len(results.AuditLogs))
	return uint64(len(results.AuditLogs)), nil
}

// ListSnapshotsRequest: list snapshots request.
type ListSnapshotsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: criteria to use when ordering the list.
	// Default value: created_at_asc
	OrderBy ListSnapshotsRequestOrderBy `json:"-"`

	// ProjectID: filter by Project ID.
	ProjectID *string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size, defines how many entries are returned in one page, must be lower or equal to 100.
	PageSize *uint32 `json:"-"`

	// VolumeID: filter snapshots by the ID of the original volume.
	VolumeID *string `json:"-"`

	// Name: filter snapshots by their names.
	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	IncludeDeleted bool `json:"-"`

	// Status: filter snapshots by their status.
	// Default value: unknown_status
	Status block_v1alpha1.SnapshotStatus `json:"-"`
}

// ListSnapshotsResponse: list snapshots response.
type ListSnapshotsResponse struct {
	// Snapshots: paginated returned list of snapshots.
	Snapshots []*SnapshotSummary `json:"snapshots"`

	// TotalCount: total number of snpashots in the project.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListSnapshotsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Snapshots = append(r.Snapshots, results.Snapshots...)
	r.TotalCount += uint64(len(results.Snapshots))
	return uint64(len(results.Snapshots)), nil
}

// ListVolumeTypesRequest: list volume types request.
type ListVolumeTypesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size, defines how many entries are returned in one page, must be lower or equal to 100.
	PageSize *uint32 `json:"-"`
}

// ListVolumeTypesResponse: list volume types response.
type ListVolumeTypesResponse struct {
	// VolumeTypes: returns paginated list of volume-types.
	VolumeTypes []*VolumeType `json:"volume_types"`

	// TotalCount: total number of volume-types currently available in stock.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVolumeTypesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVolumeTypesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListVolumeTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.VolumeTypes = append(r.VolumeTypes, results.VolumeTypes...)
	r.TotalCount += uint64(len(results.VolumeTypes))
	return uint64(len(results.VolumeTypes)), nil
}

// ListVolumesRequest: list volumes request.
type ListVolumesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: criteria to use when ordering the list.
	// Default value: created_at_asc
	OrderBy ListVolumesRequestOrderBy `json:"-"`

	// ProjectID: filter by Project ID.
	ProjectID *string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size, defines how many entries are returned in one page, must be lower or equal to 100.
	PageSize *uint32 `json:"-"`

	// Name: filter the return volumes by their names.
	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	IncludeDeleted bool `json:"-"`

	// Status: filter the volumes returned by their status.
	// Default value: unknown_status
	Status block_v1alpha1.VolumeStatus `json:"-"`
}

// ListVolumesResponse: list volumes response.
type ListVolumesResponse struct {
	// Volumes: paginated returned list of volumes.
	Volumes []*VolumeSummary `json:"volumes"`

	// TotalCount: total number of volumes in the project.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVolumesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVolumesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListVolumesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Volumes = append(r.Volumes, results.Volumes...)
	r.TotalCount += uint64(len(results.Volumes))
	return uint64(len(results.Volumes)), nil
}

// Snapshot: snapshot.
type Snapshot struct {
	// ID: UUID of the snapshot.
	ID string `json:"id"`

	// Name: name of the snapshot.
	Name string `json:"name"`

	// ParentVolume: if the parent volume was deleted, value is null.
	ParentVolume *SnapshotParentVolume `json:"parent_volume"`

	// Size: size in bytes of the snapshot.
	Size scw.Size `json:"size"`

	// ProjectID: UUID of the project the snapshot belongs to.
	ProjectID string `json:"project_id"`

	// CreatedAt: creation date of the snapshot.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the properties of a snapshot.
	UpdatedAt *time.Time `json:"updated_at"`

	// References: list of the references to the snapshot.
	References []*Reference `json:"references"`

	// Status: current status of the snapshot (available, in_use, ...).
	// Default value: unknown_status
	Status block_v1alpha1.SnapshotStatus `json:"status"`

	// Tags: list of tags assigned to the volume.
	Tags []string `json:"tags"`

	// Zone: snapshot zone.
	Zone scw.Zone `json:"zone"`

	// Class: storage class of the snapshot.
	// Default value: unknown_storage_class
	Class block_v1alpha1.StorageClass `json:"class"`

	DeletedAt *time.Time `json:"deleted_at"`

	UsedSize *scw.Size `json:"used_size"`

	CephName *string `json:"ceph_name"`

	// Namespace: ceph namespace of the volume.
	Namespace string `json:"namespace"`

	// Pool: ceph pool of the volume.
	Pool string `json:"pool"`

	OrganizationID string `json:"organization_id"`

	PurgedAt *time.Time `json:"purged_at"`
}

// Volume: volume.
type Volume struct {
	// ID: UUID of the volume.
	ID string `json:"id"`

	// Name: name of the volume.
	Name string `json:"name"`

	// Type: volume type.
	Type string `json:"type"`

	// Size: volume size in bytes.
	Size scw.Size `json:"size"`

	// ProjectID: UUID of the project to which the volume belongs.
	ProjectID string `json:"project_id"`

	// CreatedAt: creation date of the volume.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update of the properties of a volume.
	UpdatedAt *time.Time `json:"updated_at"`

	// References: list of the references to the volume.
	References []*Reference `json:"references"`

	// ParentSnapshotID: when a volume is created from a snapshot, is the UUID of the snapshot from which the volume has been created.
	ParentSnapshotID *string `json:"parent_snapshot_id"`

	// Status: current status of the volume (available, in_use, ...).
	// Default value: unknown_status
	Status block_v1alpha1.VolumeStatus `json:"status"`

	// Tags: list of tags assigned to the volume.
	Tags []string `json:"tags"`

	// Zone: volume zone.
	Zone scw.Zone `json:"zone"`

	// Specs: specifications of the volume.
	Specs *VolumeSpecifications `json:"specs"`

	OrganizationID string `json:"organization_id"`

	CephName *string `json:"ceph_name"`

	UsedSize *scw.Size `json:"used_size"`

	// Namespace: ceph namespace of the volume.
	Namespace string `json:"namespace"`

	// Pool: ceph pool of the volume.
	Pool string `json:"pool"`

	DeletedAt *time.Time `json:"deleted_at"`

	PurgedAt *time.Time `json:"purged_at"`
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
	return []scw.Zone{scw.ZoneFrPar1}
}

// ListVolumeTypes:
func (s *API) ListVolumeTypes(req *ListVolumeTypesRequest, opts ...scw.RequestOption) (*ListVolumeTypesResponse, error) {
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

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/volume-types",
		Query:  query,
	}

	var resp ListVolumeTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVolumes:
func (s *API) ListVolumes(req *ListVolumesRequest, opts ...scw.RequestOption) (*ListVolumesResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "include_deleted", req.IncludeDeleted)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/volumes",
		Query:  query,
	}

	var resp ListVolumesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolume:
func (s *API) GetVolume(req *GetVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return nil, errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
	}

	var resp Volume

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVolume:
func (s *API) DeleteVolume(req *DeleteVolumeRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/block-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListSnapshots:
func (s *API) ListSnapshots(req *ListSnapshotsRequest, opts ...scw.RequestOption) (*ListSnapshotsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "volume_id", req.VolumeID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "include_deleted", req.IncludeDeleted)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/snapshots",
		Query:  query,
	}

	var resp ListSnapshotsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSnapshot:
func (s *API) GetSnapshot(req *GetSnapshotRequest, opts ...scw.RequestOption) (*Snapshot, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	var resp Snapshot

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSnapshot:
func (s *API) DeleteSnapshot(req *DeleteSnapshotRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/block-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListAuditLogs:
func (s *API) ListAuditLogs(req *ListAuditLogsRequest, opts ...scw.RequestOption) (*ListAuditLogsResponse, error) {
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

	if fmt.Sprint(req.ResourceID) == "" {
		return nil, errors.New("field ResourceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/block-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/resources/" + fmt.Sprint(req.ResourceID) + "/audit-logs",
		Query:  query,
	}

	var resp ListAuditLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

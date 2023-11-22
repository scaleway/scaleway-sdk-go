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

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/block_admin/v1alpha1/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/block_admin/v1alpha1/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/block_admin/v1alpha1/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/block_admin/v1alpha1/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/block_admin/v1alpha1/scw"
	block_v1alpha1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/block_admin/v1alpha1/api/block/v1alpha1"
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
	PerfIops *uint32 `json:"perf_iops"`

	// Class: default value: unknown_storage_class
	Class block_v1alpha1.StorageClass `json:"class"`
}

// AuditLog: audit log.
type AuditLog struct {
	ID string `json:"id"`

	Action string `json:"action"`

	Detail *scw.JSONObject `json:"detail"`

	CreatedAt *time.Time `json:"created_at"`
}

// SnapshotSummary: snapshot summary.
type SnapshotSummary struct {
	ID string `json:"id"`

	Name string `json:"name"`

	ParentVolume *SnapshotParentVolume `json:"parent_volume"`

	Size scw.Size `json:"size"`

	ProjectID string `json:"project_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	// Status: default value: unknown_status
	Status block_v1alpha1.SnapshotStatus `json:"status"`

	Tags []string `json:"tags"`

	// Class: default value: unknown_storage_class
	Class block_v1alpha1.StorageClass `json:"class"`

	OrganizationID string `json:"organization_id"`

	CephName *string `json:"ceph_name"`

	UsedSize *scw.Size `json:"used_size"`

	PurgedAt *time.Time `json:"purged_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
}

// VolumeType: volume type.
type VolumeType struct {
	Type string `json:"type"`

	Specs *VolumeSpecifications `json:"specs"`

	Pricing *scw.Money `json:"pricing"`

	SnapshotPricing *scw.Money `json:"snapshot_pricing"`
}

// VolumeSummary: volume summary.
type VolumeSummary struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Type string `json:"type"`

	Size scw.Size `json:"size"`

	ProjectID string `json:"project_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	ParentSnapshotID *string `json:"parent_snapshot_id"`

	// Status: default value: unknown_status
	Status block_v1alpha1.VolumeStatus `json:"status"`

	Tags []string `json:"tags"`

	Specs *VolumeSpecifications `json:"specs"`

	OrganizationID string `json:"organization_id"`

	CephName *string `json:"ceph_name"`

	UsedSize *scw.Size `json:"used_size"`

	PurgedAt *time.Time `json:"purged_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
}

// Reference: reference.
type Reference struct {
	ID string `json:"id"`

	ProductResourceType string `json:"product_resource_type"`

	ProductResourceID string `json:"product_resource_id"`

	CreatedAt *time.Time `json:"created_at"`

	// Type: default value: unknown_type
	Type block_v1alpha1.ReferenceType `json:"type"`

	// Status: default value: unknown_status
	Status block_v1alpha1.ReferenceStatus `json:"status"`
}

// DeleteSnapshotRequest: delete snapshot request.
type DeleteSnapshotRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SnapshotID string `json:"-"`
}

// DeleteVolumeRequest: delete volume request.
type DeleteVolumeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	VolumeID string `json:"-"`
}

// GetServiceInfoRequest: get service info request.
type GetServiceInfoRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// GetSnapshotRequest: get snapshot request.
type GetSnapshotRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SnapshotID string `json:"-"`
}

// GetVolumeRequest: get volume request.
type GetVolumeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	VolumeID string `json:"-"`
}

// ListAuditLogsRequest: list audit logs request.
type ListAuditLogsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ResourceID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListAuditLogsRequestOrderBy `json:"-"`
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

	// OrderBy: default value: created_at_asc
	OrderBy ListSnapshotsRequestOrderBy `json:"-"`

	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	VolumeID *string `json:"-"`

	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	IncludeDeleted bool `json:"-"`
}

// ListSnapshotsResponse: list snapshots response.
type ListSnapshotsResponse struct {
	Snapshots []*SnapshotSummary `json:"snapshots"`

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

	PageSize *uint32 `json:"-"`

	Page *int32 `json:"-"`
}

// ListVolumeTypesResponse: list volume types response.
type ListVolumeTypesResponse struct {
	VolumeTypes []*VolumeType `json:"volume_types"`

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

	// OrderBy: default value: created_at_asc
	OrderBy ListVolumesRequestOrderBy `json:"-"`

	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	IncludeDeleted bool `json:"-"`
}

// ListVolumesResponse: list volumes response.
type ListVolumesResponse struct {
	Volumes []*VolumeSummary `json:"volumes"`

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
	ID string `json:"id"`

	Name string `json:"name"`

	ParentVolume *SnapshotParentVolume `json:"parent_volume"`

	Size scw.Size `json:"size"`

	ProjectID string `json:"project_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	References []*Reference `json:"references"`

	// Status: default value: unknown_status
	Status block_v1alpha1.SnapshotStatus `json:"status"`

	Tags []string `json:"tags"`

	// Class: default value: unknown_storage_class
	Class block_v1alpha1.StorageClass `json:"class"`

	OrganizationID string `json:"organization_id"`

	CephName *string `json:"ceph_name"`

	UsedSize *scw.Size `json:"used_size"`

	PurgedAt *time.Time `json:"purged_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
}

// Volume: volume.
type Volume struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Type string `json:"type"`

	Size scw.Size `json:"size"`

	ProjectID string `json:"project_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	References []*Reference `json:"references"`

	ParentSnapshotID *string `json:"parent_snapshot_id"`

	// Status: default value: unknown_status
	Status block_v1alpha1.VolumeStatus `json:"status"`

	Tags []string `json:"tags"`

	Specs *VolumeSpecifications `json:"specs"`

	OrganizationID string `json:"organization_id"`

	CephName *string `json:"ceph_name"`

	UsedSize *scw.Size `json:"used_size"`

	PurgedAt *time.Time `json:"purged_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
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

// GetServiceInfo:
func (s *API) GetServiceInfo(req *GetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
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
		Path:   "/block-admin/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)

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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

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

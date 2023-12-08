// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package ipfs_admin provides methods and message types of the ipfs_admin v1 API.
package ipfs_admin

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

type ListNamesRequestOrderBy string

const (
	ListNamesRequestOrderByCreatedAtAsc  = ListNamesRequestOrderBy("created_at_asc")
	ListNamesRequestOrderByCreatedAtDesc = ListNamesRequestOrderBy("created_at_desc")
)

func (enum ListNamesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNamesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNamesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNamesRequestOrderBy(ListNamesRequestOrderBy(tmp).String())
	return nil
}

type ListPinsRequestOrderBy string

const (
	ListPinsRequestOrderByCreatedAtAsc  = ListPinsRequestOrderBy("created_at_asc")
	ListPinsRequestOrderByCreatedAtDesc = ListPinsRequestOrderBy("created_at_desc")
)

func (enum ListPinsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListPinsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPinsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPinsRequestOrderBy(ListPinsRequestOrderBy(tmp).String())
	return nil
}

type ListVolumesRequestOrderBy string

const (
	ListVolumesRequestOrderByCreatedAtAsc  = ListVolumesRequestOrderBy("created_at_asc")
	ListVolumesRequestOrderByCreatedAtDesc = ListVolumesRequestOrderBy("created_at_desc")
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

type NameLockStatus string

const (
	NameLockStatusUnlocked = NameLockStatus("unlocked")
	NameLockStatusLocked   = NameLockStatus("locked")
)

func (enum NameLockStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unlocked"
	}
	return string(enum)
}

func (enum NameLockStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NameLockStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NameLockStatus(NameLockStatus(tmp).String())
	return nil
}

type NameProvide string

const (
	NameProvideUnknownProvide = NameProvide("unknown_provide")
	NameProvideBan            = NameProvide("ban")
	NameProvideUnban          = NameProvide("unban")
)

func (enum NameProvide) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_provide"
	}
	return string(enum)
}

func (enum NameProvide) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NameProvide) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NameProvide(NameProvide(tmp).String())
	return nil
}

type NameStatus string

const (
	NameStatusUnknownStatus = NameStatus("unknown_status")
	NameStatusQueued        = NameStatus("queued")
	NameStatusPublishing    = NameStatus("publishing")
	NameStatusFailed        = NameStatus("failed")
	NameStatusPublished     = NameStatus("published")
)

func (enum NameStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum NameStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NameStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NameStatus(NameStatus(tmp).String())
	return nil
}

type PinDetails string

const (
	PinDetailsUnknownDetails                  = PinDetails("unknown_details")
	PinDetailsPinningLookingForProvider       = PinDetails("pinning_looking_for_provider")
	PinDetailsPinningInProgress               = PinDetails("pinning_in_progress")
	PinDetailsPinningBlocksFetched            = PinDetails("pinning_blocks_fetched")
	PinDetailsPinningFetchingURLData          = PinDetails("pinning_fetching_url_data")
	PinDetailsPinnedOk                        = PinDetails("pinned_ok")
	PinDetailsUnpinnedOk                      = PinDetails("unpinned_ok")
	PinDetailsUnpinningInProgress             = PinDetails("unpinning_in_progress")
	PinDetailsFailedContainsBannedCid         = PinDetails("failed_contains_banned_cid")
	PinDetailsFailedPinning                   = PinDetails("failed_pinning")
	PinDetailsFailedPinningNoProvider         = PinDetails("failed_pinning_no_provider")
	PinDetailsFailedPinningBadCidFormat       = PinDetails("failed_pinning_bad_cid_format")
	PinDetailsFailedPinningTimeout            = PinDetails("failed_pinning_timeout")
	PinDetailsFailedPinningTooBigContent      = PinDetails("failed_pinning_too_big_content")
	PinDetailsFailedPinningUnreachableURL     = PinDetails("failed_pinning_unreachable_url")
	PinDetailsFailedPinningBadURLFormat       = PinDetails("failed_pinning_bad_url_format")
	PinDetailsFailedPinningNoURLContentLength = PinDetails("failed_pinning_no_url_content_length")
	PinDetailsFailedPinningBadURLStatusCode   = PinDetails("failed_pinning_bad_url_status_code")
	PinDetailsFailedUnpinning                 = PinDetails("failed_unpinning")
	PinDetailsCheckingCoherence               = PinDetails("checking_coherence")
	PinDetailsRescheduled                     = PinDetails("rescheduled")
)

func (enum PinDetails) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_details"
	}
	return string(enum)
}

func (enum PinDetails) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PinDetails) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PinDetails(PinDetails(tmp).String())
	return nil
}

type PinProvide string

const (
	PinProvideUnknownProvide = PinProvide("unknown_provide")
	PinProvideProvide        = PinProvide("provide")
	PinProvideNoProvide      = PinProvide("no_provide")
	PinProvideBan            = PinProvide("ban")
	PinProvideUnban          = PinProvide("unban")
)

func (enum PinProvide) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_provide"
	}
	return string(enum)
}

func (enum PinProvide) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PinProvide) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PinProvide(PinProvide(tmp).String())
	return nil
}

type PinStatus string

const (
	PinStatusUnknownStatus = PinStatus("unknown_status")
	PinStatusQueued        = PinStatus("queued")
	PinStatusPinning       = PinStatus("pinning")
	PinStatusFailed        = PinStatus("failed")
	PinStatusPinned        = PinStatus("pinned")
)

func (enum PinStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum PinStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PinStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PinStatus(PinStatus(tmp).String())
	return nil
}

type VolumeLockStatus string

const (
	VolumeLockStatusUnlocked = VolumeLockStatus("unlocked")
	VolumeLockStatusLocked   = VolumeLockStatus("locked")
)

func (enum VolumeLockStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unlocked"
	}
	return string(enum)
}

func (enum VolumeLockStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeLockStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeLockStatus(VolumeLockStatus(tmp).String())
	return nil
}

// PinCIDMeta: pin cid meta.
type PinCIDMeta struct {
	ID *string `json:"id"`
}

// PinCID: pin cid.
type PinCID struct {
	Cid *string `json:"cid"`

	Name *string `json:"name"`

	Origins []string `json:"origins"`

	Meta *PinCIDMeta `json:"meta"`
}

// PinInfo: pin info.
type PinInfo struct {
	ID *string `json:"id"`

	URL *string `json:"url"`

	Size *uint64 `json:"size"`

	Progress *uint32 `json:"progress"`

	// StatusDetails: default value: unknown_details
	StatusDetails PinDetails `json:"status_details"`
}

// Name: name.
type Name struct {
	ID string `json:"id"`

	CreatedAt *time.Time `json:"created_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	DeletingAt *time.Time `json:"deleting_at"`

	// LockStatus: default value: unlocked
	LockStatus NameLockStatus `json:"lock_status"`

	Name string `json:"name"`

	Key string `json:"key"`

	Value string `json:"value"`

	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`

	Tags []string `json:"tags"`

	UpdatedAt *time.Time `json:"updated_at"`

	// Status: default value: unknown_status
	Status NameStatus `json:"status"`

	// StatusProvide: default value: unknown_provide
	StatusProvide NameProvide `json:"status_provide"`
}

// Pin: pin.
type Pin struct {
	Cid *PinCID `json:"cid"`

	CreatedAt *time.Time `json:"created_at"`

	Delegates []string `json:"delegates"`

	DeletedAt *time.Time `json:"deleted_at"`

	DeletingAt *time.Time `json:"deleting_at"`

	Info *PinInfo `json:"info"`

	OrganizationID string `json:"organization_id"`

	PinID string `json:"pin_id"`

	ProjectID string `json:"project_id"`

	// Status: default value: unknown_status
	Status PinStatus `json:"status"`

	// StatusProvide: default value: unknown_provide
	StatusProvide PinProvide `json:"status_provide"`

	VolumeID string `json:"volume_id"`
}

// Volume: volume.
type Volume struct {
	CountPin uint64 `json:"count_pin"`

	CreatedAt *time.Time `json:"created_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	DeletingAt *time.Time `json:"deleting_at"`

	ID string `json:"id"`

	// LockStatus: default value: unlocked
	LockStatus VolumeLockStatus `json:"lock_status"`

	Name string `json:"name"`

	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`

	Size *uint32 `json:"size"`

	Tags []string `json:"tags"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// BanCIDRequest: ban cid request.
type BanCIDRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Cid string `json:"-"`
}

// GetPinRequest: get pin request.
type GetPinRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PinID: pin ID of which you want to obtain information.
	PinID string `json:"-"`
}

// GetVolumeRequest: get volume request.
type GetVolumeRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VolumeID: volume ID.
	VolumeID string `json:"-"`
}

// IpnsAPIGetNameRequest: ipns api get name request.
type IpnsAPIGetNameRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NameID: name ID.
	NameID string `json:"-"`
}

// IpnsAPIListNamesRequest: ipns api list names request.
type IpnsAPIListNamesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project ID, only names belonging to this project will be listed.
	ProjectID *string `json:"-"`

	// OrderBy: sort the order of the returned names.
	// Default value: created_at_asc
	OrderBy ListNamesRequestOrderBy `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of names to return per page.
	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`
}

// IpnsAPIUpdateNameRequest: ipns api update name request.
type IpnsAPIUpdateNameRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	NameID string `json:"-"`

	// Status: default value: unknown_status
	Status NameStatus `json:"status"`
}

// ListNamesResponse: list names response.
type ListNamesResponse struct {
	Names []*Name `json:"names"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNamesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNamesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListNamesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Names = append(r.Names, results.Names...)
	r.TotalCount += uint64(len(results.Names))
	return uint64(len(results.Names)), nil
}

// ListPinsRequest: list pins request.
type ListPinsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VolumeID: volume ID of which you want to list the pins.
	VolumeID *string `json:"-"`

	Cid *string `json:"-"`

	// ProjectID: project ID.
	ProjectID *string `json:"-"`

	// OrganizationID: organization ID.
	OrganizationID *string `json:"-"`

	// OrderBy: sort order of the returned Volume.
	// Default value: created_at_asc
	OrderBy ListPinsRequestOrderBy `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of volumes to return per page.
	PageSize *uint32 `json:"-"`

	// Status: list pins by status.
	// Default value: unknown_status
	Status PinStatus `json:"-"`
}

// ListPinsResponse: list pins response.
type ListPinsResponse struct {
	TotalCount uint64 `json:"total_count"`

	Pins []*Pin `json:"pins"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPinsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPinsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPinsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Pins = append(r.Pins, results.Pins...)
	r.TotalCount += uint64(len(results.Pins))
	return uint64(len(results.Pins)), nil
}

// ListVolumesRequest: list volumes request.
type ListVolumesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project ID, only volumes belonging to this project will be listed.
	ProjectID *string `json:"-"`

	// OrderBy: sort the order of the returned volumes.
	// Default value: created_at_asc
	OrderBy ListVolumesRequestOrderBy `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of volumes to return per page.
	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`
}

// ListVolumesResponse: list volumes response.
type ListVolumesResponse struct {
	Volumes []*Volume `json:"volumes"`

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

// ProvideCIDRequest: provide cid request.
type ProvideCIDRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Cid string `json:"-"`
}

// StopProvideCIDRequest: stop provide cid request.
type StopProvideCIDRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Cid string `json:"-"`
}

// UnbanCIDRequest: unban cid request.
type UnbanCIDRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Cid string `json:"-"`
}

// UpdatePinRequest: update pin request.
type UpdatePinRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	PinID string `json:"-"`

	// Status: default value: unknown_status
	Status PinStatus `json:"status"`
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

// GetPin: Retrieve information about the provided **pin ID**, such as status, last modification, and CID.
func (s *API) GetPin(req *GetPinRequest, opts ...scw.RequestOption) (*Pin, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PinID) == "" {
		return nil, errors.New("field PinID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/pins/" + fmt.Sprint(req.PinID) + "",
	}

	var resp Pin

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPins: Retrieve information about all pins within a volume.
func (s *API) ListPins(req *ListPinsRequest, opts ...scw.RequestOption) (*ListPinsResponse, error) {
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
	parameter.AddToQuery(query, "volume_id", req.VolumeID)
	parameter.AddToQuery(query, "cid", req.Cid)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/pins",
		Query:  query,
	}

	var resp ListPinsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVolume: Retrieve information about a specific volume.
func (s *API) GetVolume(req *GetVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return nil, errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
	}

	var resp Volume

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVolumes: Retrieve information about all volumes from a Project ID.
func (s *API) ListVolumes(req *ListVolumesRequest, opts ...scw.RequestOption) (*ListVolumesResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/volumes",
		Query:  query,
	}

	var resp ListVolumesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePin:
func (s *API) UpdatePin(req *UpdatePinRequest, opts ...scw.RequestOption) (*Pin, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PinID) == "" {
		return nil, errors.New("field PinID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/ipfs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/pins/" + fmt.Sprint(req.PinID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Pin

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopProvideCID:
func (s *API) StopProvideCID(req *StopProvideCIDRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.Cid) == "" {
		return errors.New("field Cid cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipfs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/cids/" + fmt.Sprint(req.Cid) + "/stop-provide",
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

// BanCID:
func (s *API) BanCID(req *BanCIDRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.Cid) == "" {
		return errors.New("field Cid cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipfs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/cids/" + fmt.Sprint(req.Cid) + "/ban",
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

// ProvideCID:
func (s *API) ProvideCID(req *ProvideCIDRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.Cid) == "" {
		return errors.New("field Cid cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipfs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/cids/" + fmt.Sprint(req.Cid) + "/provide",
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

// UnbanCID:
func (s *API) UnbanCID(req *UnbanCIDRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.Cid) == "" {
		return errors.New("field Cid cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/ipfs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/cids/" + fmt.Sprint(req.Cid) + "/unban",
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

type IpnsAPI struct {
	client *scw.Client
}

// NewIpnsAPI returns a IpnsAPI object from a Scaleway client.
func NewIpnsAPI(client *scw.Client) *IpnsAPI {
	return &IpnsAPI{
		client: client,
	}
}

// GetName: Retrieve information about a specific name.
func (s *IpnsAPI) GetName(req *IpnsAPIGetNameRequest, opts ...scw.RequestOption) (*Name, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NameID) == "" {
		return nil, errors.New("field NameID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/names/" + fmt.Sprint(req.NameID) + "",
	}

	var resp Name

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNames: Retrieve information about all names from a Project ID.
func (s *IpnsAPI) ListNames(req *IpnsAPIListNamesRequest, opts ...scw.RequestOption) (*ListNamesResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/ipfs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/names",
		Query:  query,
	}

	var resp ListNamesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateName:
func (s *IpnsAPI) UpdateName(req *IpnsAPIUpdateNameRequest, opts ...scw.RequestOption) (*Name, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NameID) == "" {
		return nil, errors.New("field NameID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/ipfs-admin/v1/regions/" + fmt.Sprint(req.Region) + "/names/" + fmt.Sprint(req.NameID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Name

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

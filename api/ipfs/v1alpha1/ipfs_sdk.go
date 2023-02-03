// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package ipfs provides methods and message types of the ipfs v1alpha1 API.
package ipfs

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
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

// API: pinning service ipfs API for Scaleway
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type ListPinsRequestOrderBy string

const (
	// ListPinsRequestOrderByCreatedAtAsc is [insert doc].
	ListPinsRequestOrderByCreatedAtAsc = ListPinsRequestOrderBy("created_at_asc")
	// ListPinsRequestOrderByCreatedAtDesc is [insert doc].
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
	// ListVolumesRequestOrderByCreatedAtAsc is [insert doc].
	ListVolumesRequestOrderByCreatedAtAsc = ListVolumesRequestOrderBy("created_at_asc")
	// ListVolumesRequestOrderByCreatedAtDesc is [insert doc].
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

type PinStatus string

const (
	// PinStatusUnknownStatus is [insert doc].
	PinStatusUnknownStatus = PinStatus("unknown_status")
	// PinStatusQueued is [insert doc].
	PinStatusQueued = PinStatus("queued")
	// PinStatusPinning is [insert doc].
	PinStatusPinning = PinStatus("pinning")
	// PinStatusFailed is [insert doc].
	PinStatusFailed = PinStatus("failed")
	// PinStatusPinned is [insert doc].
	PinStatusPinned = PinStatus("pinned")
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

type ListPinsResponse struct {
	TotalCount uint64 `json:"total_count"`

	Pins []*Pin `json:"pins"`
}

type ListVolumesResponse struct {
	Volumes []*Volume `json:"volumes"`

	TotalCount uint64 `json:"total_count"`
}

type Pin struct {
	PinID string `json:"pin_id"`
	// Status:
	//
	// Default value: unknown_status
	Status PinStatus `json:"status"`

	CreatedAt *time.Time `json:"created_at"`

	Cid *PinCID `json:"cid"`

	Delegates []string `json:"delegates"`

	Info *PinInfo `json:"info"`
}

type PinCID struct {
	Cid string `json:"cid"`

	Name *string `json:"name"`

	Origins []string `json:"origins"`

	Meta *PinCIDMeta `json:"meta"`
}

type PinCIDMeta struct {
	AppID string `json:"app_id"`
}

type PinInfo struct {
	StatusDetails *string `json:"status_details"`
}

type PinOptions struct {
	RequiredZones []string `json:"required_zones"`

	ReplicationCount uint32 `json:"replication_count"`
}

type ReplacePinResponse struct {
	Pin *Pin `json:"pin"`
}

type Volume struct {
	ID string `json:"id"`

	ProjectID string `json:"project_id"`

	Region scw.Region `json:"region"`

	CountPin uint64 `json:"count_pin"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	Tags []string `json:"tags"`

	Name string `json:"name"`
}

// Service API

// Regions list localities the api is available in
func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

type CreateVolumeRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	ProjectID string `json:"project_id"`
}

// CreateVolume: create volume in S3 bucket
func (s *API) CreateVolume(req *CreateVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/volumes",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Volume

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetVolumeRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	VolumeID string `json:"-"`
}

// GetVolume: get information about volume
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
		Method:  "GET",
		Path:    "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
		Headers: http.Header{},
	}

	var resp Volume

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListVolumesRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	ProjectID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListVolumesRequestOrderBy `json:"-"`
}

// ListVolumes: list volumes in project-id
func (s *API) ListVolumes(req *ListVolumesRequest, opts ...scw.RequestOption) (*ListVolumesResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/volumes",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListVolumesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateVolumeRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	VolumeID string `json:"-"`

	Name *string `json:"name"`

	Tags *[]string `json:"tags"`
}

// UpdateVolume: update volume name or tag
func (s *API) UpdateVolume(req *UpdateVolumeRequest, opts ...scw.RequestOption) (*Volume, error) {
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
		Method:  "PATCH",
		Path:    "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Volume

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteVolumeRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	VolumeID string `json:"-"`
}

// DeleteVolume: delete volume in S3 bucket
func (s *API) DeleteVolume(req *DeleteVolumeRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type CreatePinByURLRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	VolumeID string `json:"volume_id"`

	URL string `json:"url"`

	Name *string `json:"name"`

	PinOptions *PinOptions `json:"pin_options"`
}

// CreatePinByURL: add content in s3 bucket
func (s *API) CreatePinByURL(req *CreatePinByURLRequest, opts ...scw.RequestOption) (*Pin, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins/create-by-url",
		Headers: http.Header{},
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

type CreatePinByCIDRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	VolumeID string `json:"volume_id"`

	Cid string `json:"cid"`

	Name *string `json:"name"`

	Origins []string `json:"origins"`

	Meta *PinCIDMeta `json:"meta"`

	PinOptions *PinOptions `json:"pin_options"`
}

// CreatePinByCID: add content in s3 bucket
func (s *API) CreatePinByCID(req *CreatePinByCIDRequest, opts ...scw.RequestOption) (*Pin, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins/create-by-cid",
		Headers: http.Header{},
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

type CreatePinByRawRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	VolumeID string `json:"volume_id"`

	Content []byte `json:"content"`

	MimeType *string `json:"mime_type"`

	Name *string `json:"name"`

	PinOptions *PinOptions `json:"pin_options"`
}

// CreatePinByRaw: add content in s3 bucket
func (s *API) CreatePinByRaw(req *CreatePinByRawRequest, opts ...scw.RequestOption) (*Pin, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins/create-by-raw",
		Headers: http.Header{},
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

type ReplacePinRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	PinID string `json:"-"`

	VolumeID string `json:"volume_id"`

	Cid string `json:"cid"`

	Name *string `json:"name"`

	Origins []string `json:"origins"`

	Meta *PinCIDMeta `json:"meta"`

	PinOptions *PinOptions `json:"pin_options"`
}

func (s *API) ReplacePin(req *ReplacePinRequest, opts ...scw.RequestOption) (*ReplacePinResponse, error) {
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
		Method:  "POST",
		Path:    "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins/" + fmt.Sprint(req.PinID) + "/replace",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ReplacePinResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetPinRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	PinID string `json:"-"`

	VolumeID string `json:"-"`
}

// GetPin: get pin id create when content is add in s3 bucket
func (s *API) GetPin(req *GetPinRequest, opts ...scw.RequestOption) (*Pin, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "volume_id", req.VolumeID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PinID) == "" {
		return nil, errors.New("field PinID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins/" + fmt.Sprint(req.PinID) + "",
		Query:   query,
		Headers: http.Header{},
	}

	var resp Pin

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListPinsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	VolumeID string `json:"-"`

	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListPinsRequestOrderBy `json:"-"`
	// Status:
	//
	// Default value: unknown_status
	Status PinStatus `json:"-"`
}

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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListPinsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeletePinRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	PinID string `json:"-"`

	VolumeID string `json:"-"`
}

// DeletePin: remove by pin id
func (s *API) DeletePin(req *DeletePinRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "volume_id", req.VolumeID)

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PinID) == "" {
		return errors.New("field PinID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/ipfs/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/pins/" + fmt.Sprint(req.PinID) + "",
		Query:   query,
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
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

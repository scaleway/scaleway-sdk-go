// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package object_admin provides methods and message types of the object_admin v1 API.
package object_admin

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

type Backend string

const (
	BackendUnknownBackend = Backend("unknown_backend")
	BackendOpenio         = Backend("openio")
	BackendHive           = Backend("hive")
)

func (enum Backend) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_backend"
	}
	return string(enum)
}

func (enum Backend) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Backend) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Backend(Backend(tmp).String())
	return nil
}

type StorageClass string

const (
	StorageClassSTANDARD  = StorageClass("STANDARD")
	StorageClassGLACIER   = StorageClass("GLACIER")
	StorageClassONEZONEIA = StorageClass("ONEZONE_IA")
)

func (enum StorageClass) String() string {
	if enum == "" {
		// return default value if empty
		return "STANDARD"
	}
	return string(enum)
}

func (enum StorageClass) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *StorageClass) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = StorageClass(StorageClass(tmp).String())
	return nil
}

// Usage: usage.
type Usage struct {
	Date string `json:"date"`

	UsageMax float64 `json:"usage_max"`
}

// Bucket: bucket.
type Bucket struct {
	Name string `json:"name"`

	Bytes float64 `json:"bytes"`

	Objects float64 `json:"objects"`

	Status string `json:"status"`

	Dtime string `json:"dtime"`

	Mtime string `json:"mtime"`

	DamagedObjects int64 `json:"damaged_objects"`

	MissingChuncks int64 `json:"missing_chuncks"`

	StorageUsageStandard int64 `json:"storage_usage_standard"`

	StorageUsageGlacier int64 `json:"storage_usage_glacier"`

	OpenioMeta1Status string `json:"openio_meta1_status"`

	OpenioMeta2Status string `json:"openio_meta2_status"`

	HasLifecycleRules bool `json:"has_lifecycle_rules"`

	LifecycleIsHealthy bool `json:"lifecycle_is_healthy"`

	Deleted bool `json:"deleted"`

	Frozen bool `json:"frozen"`

	HasFrozenObjects bool `json:"has_frozen_objects"`

	// Backend: default value: unknown_backend
	Backend Backend `json:"backend"`

	StorageUsageOnezoneIa int64 `json:"storage_usage_onezone_ia"`
}

// PeerMeta2: peer meta2.
type PeerMeta2 struct {
	Host string `json:"host"`

	Version string `json:"version"`
}

// Bandwidth: bandwidth.
type Bandwidth struct {
	Date string `json:"date"`

	IPType string `json:"ip_type"`

	Bandwidth int64 `json:"bandwidth"`
}

// ListUsagesPerBucketResponseBucket: list usages per bucket response bucket.
type ListUsagesPerBucketResponseBucket struct {
	BucketName string `json:"bucket_name"`

	Usages []*Usage `json:"usages"`
}

// GetBucketRequest: get bucket request.
type GetBucketRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	BucketName string `json:"-"`
}

// GetBucketResponse: get bucket response.
type GetBucketResponse struct {
	OrganizationID string `json:"organization_id"`

	Bucket *Bucket `json:"bucket"`

	Peers []*PeerMeta2 `json:"peers"`

	IsSync bool `json:"is_sync"`
}

// GetUsageRequest: get usage request.
type GetUsageRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	OrganizationID string `json:"-"`

	// StorageClass: default value: STANDARD
	StorageClass StorageClass `json:"-"`
}

// GetUsageResponse: get usage response.
type GetUsageResponse struct {
	Usages []*Usage `json:"usages"`
}

// ListBandwidthsRequest: list bandwidths request.
type ListBandwidthsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	OrganizationID string `json:"organization_id"`
}

// ListBandwidthsResponse: list bandwidths response.
type ListBandwidthsResponse struct {
	Bandwidths []*Bandwidth `json:"bandwidths"`
}

// ListBucketsRequest: list buckets request.
type ListBucketsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	OrganizationID string `json:"organization_id"`
}

// ListBucketsResponse: list buckets response.
type ListBucketsResponse struct {
	Buckets []*Bucket `json:"buckets"`
}

// ListFrozenObjectsRequest: list frozen objects request.
type ListFrozenObjectsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	BucketName string `json:"-"`
}

// ListFrozenObjectsResponse: list frozen objects response.
type ListFrozenObjectsResponse struct {
	Objects []string `json:"objects"`
}

// ListUsagesPerBucketRequest: list usages per bucket request.
type ListUsagesPerBucketRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	OrganizationID string `json:"-"`

	// StorageClass: default value: STANDARD
	StorageClass StorageClass `json:"-"`
}

// ListUsagesPerBucketResponse: list usages per bucket response.
type ListUsagesPerBucketResponse struct {
	Buckets []*ListUsagesPerBucketResponseBucket `json:"buckets"`
}

// ListUsagesRequest: list usages request.
type ListUsagesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	BucketName string `json:"-"`

	OrganizationID string `json:"-"`

	// StorageClass: default value: STANDARD
	StorageClass StorageClass `json:"-"`
}

// ListUsagesResponse: list usages response.
type ListUsagesResponse struct {
	Usages []*Usage `json:"usages"`
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

// GetBucket:
func (s *API) GetBucket(req *GetBucketRequest, opts ...scw.RequestOption) (*GetBucketResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BucketName) == "" {
		return nil, errors.New("field BucketName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/object-admin/v1/regions/" + fmt.Sprint(req.Region) + "/buckets/" + fmt.Sprint(req.BucketName) + "",
	}

	var resp GetBucketResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListBuckets:
func (s *API) ListBuckets(req *ListBucketsRequest, opts ...scw.RequestOption) (*ListBucketsResponse, error) {
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
		Path:   "/object-admin/v1/regions/" + fmt.Sprint(req.Region) + "/buckets",
		Query:  query,
	}

	var resp ListBucketsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListBandwidths:
func (s *API) ListBandwidths(req *ListBandwidthsRequest, opts ...scw.RequestOption) (*ListBandwidthsResponse, error) {
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
		Path:   "/object-admin/v1/regions/" + fmt.Sprint(req.Region) + "/bandwidths",
		Query:  query,
	}

	var resp ListBandwidthsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUsage:
func (s *API) GetUsage(req *GetUsageRequest, opts ...scw.RequestOption) (*GetUsageResponse, error) {
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
	parameter.AddToQuery(query, "storage_class", req.StorageClass)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/object-admin/v1/regions/" + fmt.Sprint(req.Region) + "/usage",
		Query:  query,
	}

	var resp GetUsageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListUsagesPerBucket:
func (s *API) ListUsagesPerBucket(req *ListUsagesPerBucketRequest, opts ...scw.RequestOption) (*ListUsagesPerBucketResponse, error) {
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
	parameter.AddToQuery(query, "storage_class", req.StorageClass)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/object-admin/v1/regions/" + fmt.Sprint(req.Region) + "/usages",
		Query:  query,
	}

	var resp ListUsagesPerBucketResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListUsages:
func (s *API) ListUsages(req *ListUsagesRequest, opts ...scw.RequestOption) (*ListUsagesResponse, error) {
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
	parameter.AddToQuery(query, "storage_class", req.StorageClass)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BucketName) == "" {
		return nil, errors.New("field BucketName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/object-admin/v1/regions/" + fmt.Sprint(req.Region) + "/usages/" + fmt.Sprint(req.BucketName) + "",
		Query:  query,
	}

	var resp ListUsagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFrozenObjects:
func (s *API) ListFrozenObjects(req *ListFrozenObjectsRequest, opts ...scw.RequestOption) (*ListFrozenObjectsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BucketName) == "" {
		return nil, errors.New("field BucketName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/object-admin/v1/regions/" + fmt.Sprint(req.Region) + "/frozen-objects/" + fmt.Sprint(req.BucketName) + "",
	}

	var resp ListFrozenObjectsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

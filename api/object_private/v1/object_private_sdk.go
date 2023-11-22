// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package object_private provides methods and message types of the object_private v1 API.
package object_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/object_private/v1/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/object_private/v1/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/object_private/v1/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/object_private/v1/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/object_private/v1/scw"
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

type BucketInfoBackend string

const (
	BucketInfoBackendUnknownBackend = BucketInfoBackend("unknown_backend")
	BucketInfoBackendOld            = BucketInfoBackend("old")
	BucketInfoBackendNew            = BucketInfoBackend("new")
)

func (enum BucketInfoBackend) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_backend"
	}
	return string(enum)
}

func (enum BucketInfoBackend) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BucketInfoBackend) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BucketInfoBackend(BucketInfoBackend(tmp).String())
	return nil
}

type BucketInfoStatus string

const (
	BucketInfoStatusUnknown     = BucketInfoStatus("unknown")
	BucketInfoStatusEnabled     = BucketInfoStatus("enabled")
	BucketInfoStatusFrozen      = BucketInfoStatus("frozen")
	BucketInfoStatusSafedeleted = BucketInfoStatus("safedeleted")
	BucketInfoStatusDeleted     = BucketInfoStatus("deleted")
)

func (enum BucketInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum BucketInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BucketInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BucketInfoStatus(BucketInfoStatus(tmp).String())
	return nil
}

// BucketInfo: bucket info.
type BucketInfo struct {
	CurrentSize float64 `json:"current_size"`

	CurrentObjects float64 `json:"current_objects"`

	UpdatedAt *time.Time `json:"updated_at"`

	IsPublic bool `json:"is_public"`

	// Status: default value: unknown
	Status BucketInfoStatus `json:"status"`

	CurrentSegments float64 `json:"current_segments"`

	// Backend: default value: unknown_backend
	Backend BucketInfoBackend `json:"backend"`
}

// BucketMetrics: bucket metrics.
type BucketMetrics struct {
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

// DeleteBucketRequest: delete bucket request.
type DeleteBucketRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Deprecated
	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	BucketName string `json:"bucket_name"`

	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`
}

// GetBucketMetricsRequest: get bucket metrics request.
type GetBucketMetricsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	BucketID string `json:"-"`

	StartDate *time.Time `json:"-"`

	EndDate *time.Time `json:"-"`

	MetricName *string `json:"-"`
}

// GetBucketsInfoRequest: get buckets info request.
type GetBucketsInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Deprecated
	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	BucketsName []string `json:"buckets_name"`

	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`
}

// GetBucketsInfoResponse: get buckets info response.
type GetBucketsInfoResponse struct {
	// Deprecated
	CurrentSize *float64 `json:"current_size,omitempty"`

	// Deprecated
	CurrentObjects *float64 `json:"current_objects,omitempty"`

	// Deprecated
	QuotaSize *float64 `json:"quota_size,omitempty"`

	// Deprecated
	QuotaObjects *float64 `json:"quota_objects,omitempty"`

	// Deprecated
	QuotaBuckets *float64 `json:"quota_buckets,omitempty"`

	Buckets map[string]*BucketInfo `json:"buckets"`
}

// GetProjectInfoRequest: get project info request.
type GetProjectInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ProjectID string `json:"project_id"`
}

// GetProjectInfoResponse: get project info response.
type GetProjectInfoResponse struct {
	BucketCount float64 `json:"bucket_count"`

	ObjectCount float64 `json:"object_count"`

	StorageUsage float64 `json:"storage_usage"`

	QuotaStorage float64 `json:"quota_storage"`

	QuotaObjects float64 `json:"quota_objects"`

	QuotaBuckets float64 `json:"quota_buckets"`
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
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms}
}

// GetBucketsInfo:
func (s *API) GetBucketsInfo(req *GetBucketsInfoRequest, opts ...scw.RequestOption) (*GetBucketsInfoResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/object-private/v1/regions/" + fmt.Sprint(req.Region) + "/buckets-info",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp GetBucketsInfoResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetProjectInfo:
func (s *API) GetProjectInfo(req *GetProjectInfoRequest, opts ...scw.RequestOption) (*GetProjectInfoResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/object-private/v1/regions/" + fmt.Sprint(req.Region) + "/project-info",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp GetProjectInfoResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBucket:
func (s *API) DeleteBucket(req *DeleteBucketRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/object-private/v1/regions/" + fmt.Sprint(req.Region) + "/delete-bucket",
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

// GetBucketMetrics:
func (s *API) GetBucketMetrics(req *GetBucketMetricsRequest, opts ...scw.RequestOption) (*BucketMetrics, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)
	parameter.AddToQuery(query, "metric_name", req.MetricName)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BucketID) == "" {
		return nil, errors.New("field BucketID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/object-private/v1/regions/" + fmt.Sprint(req.Region) + "/buckets/" + fmt.Sprint(req.BucketID) + "/metrics",
		Query:  query,
	}

	var resp BucketMetrics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

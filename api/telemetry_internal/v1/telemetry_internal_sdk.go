// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package telemetry_internal provides methods and message types of the telemetry_internal v1 API.
package telemetry_internal

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

type EventType string

const (
	EventTypeStart = EventType("start")
	EventTypeStop  = EventType("stop")
)

func (enum EventType) String() string {
	if enum == "" {
		// return default value if empty
		return "start"
	}
	return string(enum)
}

func (enum EventType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EventType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EventType(EventType(tmp).String())
	return nil
}

type State string

const (
	StateNew          = State("new")
	StateProcessed    = State("processed")
	StateInvalid      = State("invalid")
	StateDangling     = State("dangling")
	StateIgnored      = State("ignored")
	StateUnauthorized = State("unauthorized")
)

func (enum State) String() string {
	if enum == "" {
		// return default value if empty
		return "new"
	}
	return string(enum)
}

func (enum State) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *State) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = State(State(tmp).String())
	return nil
}

// TelemetrySummary: telemetry summary.
type TelemetrySummary struct {
	ID string `json:"id"`

	OperationPath string `json:"operation_path"`

	OrganizationID string `json:"organization_id"`

	ResourceID string `json:"resource_id"`

	AccountingTime *time.Time `json:"accounting_time"`

	// EventType: default value: start
	EventType EventType `json:"event_type"`

	Quantity string `json:"quantity"`

	// State: default value: new
	State State `json:"state"`

	Unit string `json:"unit"`
}

// TelemetryRequest: telemetry request.
type TelemetryRequest struct {
	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`

	ResourceID string `json:"resource_id"`

	AccountingTime *time.Time `json:"accounting_time"`

	OperationPath string `json:"operation_path"`

	// EventType: default value: start
	EventType EventType `json:"event_type"`

	Quantity string `json:"quantity"`

	Unit string `json:"unit"`
}

// TelemetryResponse: telemetry response.
type TelemetryResponse struct {
	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`

	ResourceID string `json:"resource_id"`

	AccountingTime *time.Time `json:"accounting_time"`

	OperationPath string `json:"operation_path"`

	// EventType: default value: start
	EventType EventType `json:"event_type"`

	EventTime *time.Time `json:"event_time"`

	Quantity string `json:"quantity"`

	Unit string `json:"unit"`

	DupsTotal uint64 `json:"dups_total"`

	DupsSkipped uint64 `json:"dups_skipped"`

	DupsEmitted uint64 `json:"dups_emitted"`
}

// GetTelemetryRequest: get telemetry request.
type GetTelemetryRequest struct {
	OrganizationID string `json:"-"`

	ResourceID string `json:"-"`

	AccountingTime time.Time `json:"-"`

	Key string `json:"-"`
}

// GetTelemetryResponse: get telemetry response.
type GetTelemetryResponse struct {
	ID string `json:"id"`

	OperationPath string `json:"operation_path"`

	OrganizationID string `json:"organization_id"`

	ResourceID string `json:"resource_id"`

	AccountingTime *time.Time `json:"accounting_time"`

	EventTime *time.Time `json:"event_time"`

	ConsolidationTime *time.Time `json:"consolidation_time"`

	// EventType: default value: start
	EventType EventType `json:"event_type"`

	Quantity string `json:"quantity"`

	Unit string `json:"unit"`

	// State: default value: new
	State State `json:"state"`
}

// ListTelemetriesRequest: list telemetries request.
type ListTelemetriesRequest struct {
	OrganizationID string `json:"-"`

	ResourceID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListTelemetriesResponse: list telemetries response.
type ListTelemetriesResponse struct {
	Telemetries []*TelemetrySummary `json:"telemetries"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTelemetriesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTelemetriesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTelemetriesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Telemetries = append(r.Telemetries, results.Telemetries...)
	r.TotalCount += uint32(len(results.Telemetries))
	return uint32(len(results.Telemetries)), nil
}

// PushTelemetriesRequest: push telemetries request.
type PushTelemetriesRequest struct {
	Deduplicate *bool `json:"deduplicate,omitempty"`

	Telemetries []*TelemetryRequest `json:"telemetries"`
}

// PushTelemetriesResponse: push telemetries response.
type PushTelemetriesResponse struct {
	Telemetries []*TelemetryResponse `json:"telemetries"`
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
	return []scw.Region{scw.RegionFrPar}
}

// ListTelemetries:
func (s *API) ListTelemetries(req *ListTelemetriesRequest, opts ...scw.RequestOption) (*ListTelemetriesResponse, error) {
	var err error

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
	parameter.AddToQuery(query, "resource_id", req.ResourceID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/telemetry-internal/v1/telemetries",
		Query:  query,
	}

	var resp ListTelemetriesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTelemetry:
func (s *API) GetTelemetry(req *GetTelemetryRequest, opts ...scw.RequestOption) (*GetTelemetryResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	if fmt.Sprint(req.ResourceID) == "" {
		return nil, errors.New("field ResourceID cannot be empty in request")
	}

	if fmt.Sprint(req.AccountingTime) == "" {
		return nil, errors.New("field AccountingTime cannot be empty in request")
	}

	if fmt.Sprint(req.Key) == "" {
		return nil, errors.New("field Key cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/telemetry-internal/v1/telemetries/" + fmt.Sprint(req.OrganizationID) + "/" + fmt.Sprint(req.ResourceID) + "/" + fmt.Sprint(req.AccountingTime) + "/" + fmt.Sprint(req.Key) + "",
	}

	var resp GetTelemetryResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PushTelemetries:
func (s *API) PushTelemetries(req *PushTelemetriesRequest, opts ...scw.RequestOption) (*PushTelemetriesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/telemetry-internal/v1/telemetries",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PushTelemetriesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

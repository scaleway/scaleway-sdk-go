// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package abuse_admin provides methods and message types of the abuse_admin v1 API.
package abuse_admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/abuse_admin/v1/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/abuse_admin/v1/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/abuse_admin/v1/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/abuse_admin/v1/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/abuse_admin/v1/scw"
	abuse_private_v1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/abuse_admin/v1/api/abuse_private/v1"
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

type AbuseAbuseStatus string

const (
	AbuseAbuseStatusUnknownStatus = AbuseAbuseStatus("unknown_status")
	AbuseAbuseStatusNew           = AbuseAbuseStatus("new")
	AbuseAbuseStatusClosed        = AbuseAbuseStatus("closed")
	AbuseAbuseStatusCancelled     = AbuseAbuseStatus("cancelled")
	AbuseAbuseStatusValidated     = AbuseAbuseStatus("validated")
	AbuseAbuseStatusConfirmed     = AbuseAbuseStatus("confirmed")
	AbuseAbuseStatusReopened      = AbuseAbuseStatus("reopened")
)

func (enum AbuseAbuseStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum AbuseAbuseStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AbuseAbuseStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AbuseAbuseStatus(AbuseAbuseStatus(tmp).String())
	return nil
}

type AbuseService string

const (
	AbuseServiceUnknownService = AbuseService("unknown_service")
	AbuseServiceDedibox        = AbuseService("dedibox")
	AbuseServiceElements       = AbuseService("elements")
	AbuseServiceMutu           = AbuseService("mutu")
	AbuseServiceStorage        = AbuseService("storage")
	AbuseServiceRegistry       = AbuseService("registry")
	AbuseServiceCloudCompute   = AbuseService("cloud_compute")
	AbuseServiceLbaas          = AbuseService("lbaas")
	AbuseServiceBmaas          = AbuseService("bmaas")
)

func (enum AbuseService) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_service"
	}
	return string(enum)
}

func (enum AbuseService) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AbuseService) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AbuseService(AbuseService(tmp).String())
	return nil
}

type ListAbusesRequestOrderBy string

const (
	ListAbusesRequestOrderByCreatedAtAsc    = ListAbusesRequestOrderBy("created_at_asc")
	ListAbusesRequestOrderByCreatedAtDesc   = ListAbusesRequestOrderBy("created_at_desc")
	ListAbusesRequestOrderByValidatedAtAsc  = ListAbusesRequestOrderBy("validated_at_asc")
	ListAbusesRequestOrderByValidatedAtDesc = ListAbusesRequestOrderBy("validated_at_desc")
	ListAbusesRequestOrderByResolvedAtAsc   = ListAbusesRequestOrderBy("resolved_at_asc")
	ListAbusesRequestOrderByResolvedAtDesc  = ListAbusesRequestOrderBy("resolved_at_desc")
	ListAbusesRequestOrderByObservedAtAsc   = ListAbusesRequestOrderBy("observed_at_asc")
	ListAbusesRequestOrderByObservedAtDesc  = ListAbusesRequestOrderBy("observed_at_desc")
)

func (enum ListAbusesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListAbusesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListAbusesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListAbusesRequestOrderBy(ListAbusesRequestOrderBy(tmp).String())
	return nil
}

// Abuse: abuse.
type Abuse struct {
	// ConfirmedAt: tODO.
	ConfirmedAt *time.Time `json:"confirmed_at"`

	// Details: tODO.
	Details map[string]string `json:"details"`

	// ID: tODO.
	ID string `json:"id"`

	// LegacyID: tODO.
	LegacyID *string `json:"legacy_id"`

	// LinkedReports: tODO.
	LinkedReports []*abuse_private_v1.Report `json:"linked_reports"`

	// ObserverIP: tODO.
	ObserverIP *string `json:"observer_ip"`

	// OffendingCustomerID: tODO.
	OffendingCustomerID *string `json:"offending_customer_id"`

	// Report: tODO.
	Report *abuse_private_v1.Report `json:"report"`

	// ResolveAnswer: tODO.
	ResolveAnswer *string `json:"resolve_answer"`

	// ResolveDescription: tODO.
	ResolveDescription *string `json:"resolve_description"`

	// ResolvedAt: tODO.
	ResolvedAt *time.Time `json:"resolved_at"`

	// Resolver: tODO.
	Resolver *string `json:"resolver"`

	// Service: tODO.
	// Default value: unknown_service
	Service AbuseService `json:"service"`

	// Status: tODO.
	// Default value: unknown_status
	Status AbuseAbuseStatus `json:"status"`

	// Token: tODO.
	Token *string `json:"token"`

	// ValidatedAt: tODO.
	ValidatedAt *time.Time `json:"validated_at"`
}

// CreateReportRequest: create report request.
type CreateReportRequest struct {
	AbuseID *string `json:"abuse_id,omitempty"`

	// CompanyName: tODO.
	CompanyName *string `json:"company_name,omitempty"`

	// CreatorType: tODO.
	// Default value: unknown_creator_type
	CreatorType abuse_private_v1.ReportCreatorType `json:"creator_type"`

	// ObservedAt: tODO.
	ObservedAt *time.Time `json:"observed_at,omitempty"`

	// ObserverDescription: tODO.
	ObserverDescription *string `json:"observer_description,omitempty"`

	// ObserverEmail: tODO.
	ObserverEmail *string `json:"observer_email,omitempty"`

	// ObserverVisible: tODO.
	ObserverVisible bool `json:"observer_visible"`

	// OffendingResource: tODO.
	OffendingResource *string `json:"offending_resource,omitempty"`

	// OffendingResourceType: tODO.
	// Default value: unknown_resource_type
	OffendingResourceType abuse_private_v1.ReportResourceType `json:"offending_resource_type"`

	// OffendingURL: tODO.
	OffendingURL *string `json:"offending_url,omitempty"`

	// Type: tODO.
	// Default value: unknown_type
	Type abuse_private_v1.ReportType `json:"type"`
}

// GetAbuseRequest: TODO.
type GetAbuseRequest struct {
	// AbuseID: tODO.
	AbuseID string `json:"-"`
}

// ListAbusesRequest: list abuses request.
type ListAbusesRequest struct {
	// OffendingCustomerID: tODO.
	OffendingCustomerID *string `json:"-"`

	// OffendingResource: tODO.
	OffendingResource *string `json:"-"`

	// OffendingResourceType: tODO.
	// Default value: unknown_resource_type
	OffendingResourceType abuse_private_v1.ReportResourceType `json:"-"`

	// OrderBy: tODO.
	// Default value: created_at_asc
	OrderBy ListAbusesRequestOrderBy `json:"-"`

	// Origin: tODO.
	Origin *string `json:"-"`

	// Page: tODO.
	Page *int32 `json:"-"`

	// PageSize: tODO.
	PageSize *uint32 `json:"-"`

	// Service: tODO.
	// Default value: unknown_service
	Service AbuseService `json:"-"`

	// Status: tODO.
	Status []string `json:"-"`

	// Type: tODO.
	// Default value: unknown_type
	Type abuse_private_v1.ReportType `json:"-"`
}

// ListAbusesResponse: TODO.
type ListAbusesResponse struct {
	// Abuses: tODO.
	Abuses []*Abuse `json:"abuses"`

	// TotalCount: tODO.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListAbusesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListAbusesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListAbusesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Abuses = append(r.Abuses, results.Abuses...)
	r.TotalCount += uint32(len(results.Abuses))
	return uint32(len(results.Abuses)), nil
}

// UpdateAbuseRequest: update abuse request.
type UpdateAbuseRequest struct {
	// AbuseID: tODO.
	AbuseID string `json:"-"`

	// ConfirmedAt: tODO.
	ConfirmedAt *time.Time `json:"confirmed_at,omitempty"`

	// ObservedAt: tODO.
	ObservedAt *time.Time `json:"observed_at,omitempty"`

	// ObserverDescription: tODO.
	ObserverDescription *string `json:"observer_description,omitempty"`

	// ObserverEmail: tODO.
	ObserverEmail *string `json:"observer_email,omitempty"`

	// ObserverVisible: tODO.
	ObserverVisible bool `json:"observer_visible"`

	// OffendingResource: tODO.
	OffendingResource *string `json:"offending_resource,omitempty"`

	// OffendingResourceType: tODO.
	// Default value: unknown_resource_type
	OffendingResourceType abuse_private_v1.ReportResourceType `json:"offending_resource_type"`

	// ReplyText: tODO.
	ReplyText *string `json:"reply_text,omitempty"`

	// ResolveDescription: tODO.
	ResolveDescription *string `json:"resolve_description,omitempty"`

	// Resolver: tODO.
	Resolver *string `json:"resolver,omitempty"`

	// Status: tODO.
	// Default value: unknown_status
	Status AbuseAbuseStatus `json:"status"`

	// Type: tODO.
	// Default value: unknown_type
	Type abuse_private_v1.ReportType `json:"type"`
}

// This API allows you to create report, list and update abuses.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// CreateReport: Create a report.
func (s *API) CreateReport(req *CreateReportRequest, opts ...scw.RequestOption) (*abuse_private_v1.Report, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/abuse-admin/v1/reports",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp abuse_private_v1.Report

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAbuse: Update the abuse.
func (s *API) UpdateAbuse(req *UpdateAbuseRequest, opts ...scw.RequestOption) (*Abuse, error) {
	var err error

	if fmt.Sprint(req.AbuseID) == "" {
		return nil, errors.New("field AbuseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/abuse-admin/v1/abuses/" + fmt.Sprint(req.AbuseID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Abuse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListAbuses: List organization abuse.
func (s *API) ListAbuses(req *ListAbusesRequest, opts ...scw.RequestOption) (*ListAbusesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "offending_customer_id", req.OffendingCustomerID)
	parameter.AddToQuery(query, "offending_resource", req.OffendingResource)
	parameter.AddToQuery(query, "offending_resource_type", req.OffendingResourceType)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "origin", req.Origin)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "service", req.Service)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "type", req.Type)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse-admin/v1/abuses",
		Query:  query,
	}

	var resp ListAbusesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAbuse: Get an abuse details.
func (s *API) GetAbuse(req *GetAbuseRequest, opts ...scw.RequestOption) (*Abuse, error) {
	var err error

	if fmt.Sprint(req.AbuseID) == "" {
		return nil, errors.New("field AbuseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse-admin/v1/abuses/" + fmt.Sprint(req.AbuseID) + "",
	}

	var resp Abuse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package abuse_private provides methods and message types of the abuse_private v1 API.
package abuse_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/abuse_private/v1/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/abuse_private/v1/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/abuse_private/v1/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/abuse_private/v1/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/abuse_private/v1/scw"
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

type AbuseCreatorType string

const (
	AbuseCreatorTypeUnknownCreatorType = AbuseCreatorType("unknown_creator_type")
	AbuseCreatorTypeOwner              = AbuseCreatorType("owner")
	AbuseCreatorTypeRepresentative     = AbuseCreatorType("representative")
)

func (enum AbuseCreatorType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_creator_type"
	}
	return string(enum)
}

func (enum AbuseCreatorType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AbuseCreatorType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AbuseCreatorType(AbuseCreatorType(tmp).String())
	return nil
}

type AbuseStatus string

const (
	AbuseStatusUnknownStatus = AbuseStatus("unknown_status")
	AbuseStatusNew           = AbuseStatus("new")
	AbuseStatusClosed        = AbuseStatus("closed")
	AbuseStatusCancelled     = AbuseStatus("cancelled")
	AbuseStatusValidated     = AbuseStatus("validated")
	AbuseStatusConfirmed     = AbuseStatus("confirmed")
	AbuseStatusReopened      = AbuseStatus("reopened")
)

func (enum AbuseStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum AbuseStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AbuseStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AbuseStatus(AbuseStatus(tmp).String())
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

type ReportCreatorType string

const (
	ReportCreatorTypeUnknownCreatorType = ReportCreatorType("unknown_creator_type")
	ReportCreatorTypeOwner              = ReportCreatorType("owner")
	ReportCreatorTypeRepresentative     = ReportCreatorType("representative")
)

func (enum ReportCreatorType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_creator_type"
	}
	return string(enum)
}

func (enum ReportCreatorType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReportCreatorType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReportCreatorType(ReportCreatorType(tmp).String())
	return nil
}

type ReportResourceType string

const (
	ReportResourceTypeUnknownResourceType = ReportResourceType("unknown_resource_type")
	ReportResourceTypeEmail               = ReportResourceType("email")
	ReportResourceTypeDomain              = ReportResourceType("domain")
	ReportResourceTypeURL                 = ReportResourceType("url")
	ReportResourceTypeIPv4                = ReportResourceType("ipv4")
	ReportResourceTypeIPv6                = ReportResourceType("ipv6")
	ReportResourceTypeBucket              = ReportResourceType("bucket")
	ReportResourceTypeS3Object            = ReportResourceType("s3_object")
)

func (enum ReportResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_resource_type"
	}
	return string(enum)
}

func (enum ReportResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReportResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReportResourceType(ReportResourceType(tmp).String())
	return nil
}

type ReportStatus string

const (
	ReportStatusUnknownStatus = ReportStatus("unknown_status")
	ReportStatusNew           = ReportStatus("new")
	ReportStatusClosed        = ReportStatus("closed")
	ReportStatusCancelled     = ReportStatus("cancelled")
	ReportStatusValidated     = ReportStatus("validated")
	ReportStatusConfirmed     = ReportStatus("confirmed")
	ReportStatusReopened      = ReportStatus("reopened")
)

func (enum ReportStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum ReportStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReportStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReportStatus(ReportStatus(tmp).String())
	return nil
}

type ReportType string

const (
	ReportTypeUnknownType     = ReportType("unknown_type")
	ReportTypeBruteforce      = ReportType("bruteforce")
	ReportTypeCopyright       = ReportType("copyright")
	ReportTypeIrc             = ReportType("irc")
	ReportTypeOpenrelay       = ReportType("openrelay")
	ReportTypeSecurityhole    = ReportType("securityhole")
	ReportTypeVirus           = ReportType("virus")
	ReportTypeSpam            = ReportType("spam")
	ReportTypeBotnet          = ReportType("botnet")
	ReportTypePhishing        = ReportType("phishing")
	ReportTypeDos             = ReportType("dos")
	ReportTypeIllicitecontent = ReportType("illicitecontent")
)

func (enum ReportType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum ReportType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReportType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReportType(ReportType(tmp).String())
	return nil
}

// Abuse: abuse.
type Abuse struct {
	ID string `json:"id"`

	Type string `json:"type"`

	Description *string `json:"description"`

	ReplyText *string `json:"reply_text"`

	ResolveDescription *string `json:"resolve_description"`

	ObserverEmail *string `json:"observer_email"`

	ObserverVisible bool `json:"observer_visible"`

	ObserverDescription *string `json:"observer_description"`

	Details map[string]string `json:"details"`

	ConfirmedAt *time.Time `json:"confirmed_at"`

	ValidatedAt *time.Time `json:"validated_at"`

	ResolvedAt *time.Time `json:"resolved_at"`

	// Status: default value: unknown_status
	Status AbuseStatus `json:"status"`

	OffendingResource string `json:"offending_resource"`

	OffendingURL *string `json:"offending_url"`

	// OwnerType: default value: unknown_creator_type
	OwnerType AbuseCreatorType `json:"owner_type"`

	CompanyName *string `json:"company_name"`
}

// AbuseAPIGetAbuseRequest: abuse api get abuse request.
type AbuseAPIGetAbuseRequest struct {
	AbuseID string `json:"-"`
}

// AbuseAPIListAbusesRequest: abuse api list abuses request.
type AbuseAPIListAbusesRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListAbusesRequestOrderBy `json:"-"`

	Status []AbuseStatus `json:"-"`

	OrganizationID string `json:"-"`
}

// AbuseAPIUpdateAbuseRequest: abuse api update abuse request.
type AbuseAPIUpdateAbuseRequest struct {
	AbuseID string `json:"-"`

	// Status: default value: unknown_status
	Status AbuseStatus `json:"status"`

	ReplyText *string `json:"reply_text,omitempty"`

	ResolveDescription *string `json:"resolve_description,omitempty"`
}

// CreateReportRequest: create report request.
type CreateReportRequest struct {
	// OffendingResourceType: default value: unknown_resource_type
	OffendingResourceType ReportResourceType `json:"offending_resource_type"`

	// Type: default value: unknown_type
	Type ReportType `json:"type"`

	ObserverEmail *string `json:"observer_email,omitempty"`

	ObserverVisible bool `json:"observer_visible"`

	ObserverDescription *string `json:"observer_description,omitempty"`

	ObservedAt *time.Time `json:"observed_at,omitempty"`

	OffendingResource *string `json:"offending_resource,omitempty"`

	Captcha *string `json:"captcha,omitempty"`

	OffendingURL *string `json:"offending_url,omitempty"`

	// CreatorType: default value: unknown_creator_type
	CreatorType ReportCreatorType `json:"creator_type"`

	CompanyName *string `json:"company_name,omitempty"`
}

// ListAbusesResponse: list abuses response.
type ListAbusesResponse struct {
	TotalCount uint32 `json:"total_count"`

	Abuses []*Abuse `json:"abuses"`
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

// Report: report.
type Report struct {
	ID string `json:"id"`

	// OffendingResourceType: default value: unknown_resource_type
	OffendingResourceType ReportResourceType `json:"offending_resource_type"`

	// Status: default value: unknown_status
	Status ReportStatus `json:"status"`

	// Type: default value: unknown_type
	Type ReportType `json:"type"`

	ObserverEmail *string `json:"observer_email"`

	ObserverVisible bool `json:"observer_visible"`

	ObserverDescription *string `json:"observer_description"`

	ObservedAt *time.Time `json:"observed_at"`

	OffendingResource *string `json:"offending_resource"`

	CreatedAt *time.Time `json:"created_at"`

	AbuseID *string `json:"abuse_id"`

	OffendingURL *string `json:"offending_url"`

	// CreatorType: default value: unknown_creator_type
	CreatorType ReportCreatorType `json:"creator_type"`

	CompanyName *string `json:"company_name"`
}

// This API allows you to create.
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
func (s *API) CreateReport(req *CreateReportRequest, opts ...scw.RequestOption) (*Report, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/abuse-private/v1/reports",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Report

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to list and update report.
type AbuseAPI struct {
	client *scw.Client
}

// NewAbuseAPI returns a AbuseAPI object from a Scaleway client.
func NewAbuseAPI(client *scw.Client) *AbuseAPI {
	return &AbuseAPI{
		client: client,
	}
}

// UpdateAbuse: Update the abuse.
func (s *AbuseAPI) UpdateAbuse(req *AbuseAPIUpdateAbuseRequest, opts ...scw.RequestOption) (*Abuse, error) {
	var err error

	if fmt.Sprint(req.AbuseID) == "" {
		return nil, errors.New("field AbuseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/abuse-private/v1/abuses/" + fmt.Sprint(req.AbuseID) + "",
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
func (s *AbuseAPI) ListAbuses(req *AbuseAPIListAbusesRequest, opts ...scw.RequestOption) (*ListAbusesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse-private/v1/abuses",
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
func (s *AbuseAPI) GetAbuse(req *AbuseAPIGetAbuseRequest, opts ...scw.RequestOption) (*Abuse, error) {
	var err error

	if fmt.Sprint(req.AbuseID) == "" {
		return nil, errors.New("field AbuseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse-private/v1/abuses/" + fmt.Sprint(req.AbuseID) + "",
	}

	var resp Abuse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

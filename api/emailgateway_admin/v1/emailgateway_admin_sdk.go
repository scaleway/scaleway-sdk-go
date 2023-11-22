// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package emailgateway_admin provides methods and message types of the emailgateway_admin v1 API.
package emailgateway_admin

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

type EmailProvider string

const (
	// Unknown provider.
	EmailProviderUnknownProvider = EmailProvider("unknown_provider")
	// Sendgrid.
	EmailProviderSendgrid = EmailProvider("sendgrid")
	// Mailjet.
	EmailProviderMailjet = EmailProvider("mailjet")
)

func (enum EmailProvider) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_provider"
	}
	return string(enum)
}

func (enum EmailProvider) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EmailProvider) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EmailProvider(EmailProvider(tmp).String())
	return nil
}

type EmailStatus string

const (
	// Unknown status.
	EmailStatusUnknownStatus = EmailStatus("unknown_status")
	// Queued.
	EmailStatusQueued = EmailStatus("queued")
	// Processing.
	EmailStatusProcessing = EmailStatus("processing")
	// Delivered.
	EmailStatusDelivered = EmailStatus("delivered")
	// Failed.
	EmailStatusFailed = EmailStatus("failed")
)

func (enum EmailStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum EmailStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EmailStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EmailStatus(EmailStatus(tmp).String())
	return nil
}

type EventType string

const (
	// Unknown type.
	EventTypeUnknownType = EventType("unknown_type")
	// Processed.
	EventTypeProcessed = EventType("processed")
	// Dropped.
	EventTypeDropped = EventType("dropped")
	// Delivered.
	EventTypeDelivered = EventType("delivered")
	// Deferred.
	EventTypeDeferred = EventType("deferred")
	// Bounced.
	EventTypeBounced = EventType("bounced")
	// Opened.
	EventTypeOpened = EventType("opened")
	// Clicked.
	EventTypeClicked = EventType("clicked")
	// Spam report.
	EventTypeSpamReport = EventType("spam_report")
	// Unsubscribed.
	EventTypeUnsubscribed = EventType("unsubscribed")
)

func (enum EventType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
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

type ListEmailEventsRequestOrderBy string

const (
	// Time ascending.
	ListEmailEventsRequestOrderByTimeAsc = ListEmailEventsRequestOrderBy("time_asc")
	// Time descending.
	ListEmailEventsRequestOrderByTimeDesc = ListEmailEventsRequestOrderBy("time_desc")
)

func (enum ListEmailEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "time_asc"
	}
	return string(enum)
}

func (enum ListEmailEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListEmailEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListEmailEventsRequestOrderBy(ListEmailEventsRequestOrderBy(tmp).String())
	return nil
}

type ListEmailsRequestOrderBy string

const (
	// Sent date ascending.
	ListEmailsRequestOrderBySentAtAsc = ListEmailsRequestOrderBy("sent_at_asc")
	// Sent date descending.
	ListEmailsRequestOrderBySentAtDesc = ListEmailsRequestOrderBy("sent_at_desc")
)

func (enum ListEmailsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "sent_at_asc"
	}
	return string(enum)
}

func (enum ListEmailsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListEmailsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListEmailsRequestOrderBy(ListEmailsRequestOrderBy(tmp).String())
	return nil
}

// Event: event.
type Event struct {
	ID string `json:"id"`

	EmailID string `json:"email_id"`

	Time *time.Time `json:"time"`

	// Type: default value: unknown_type
	Type EventType `json:"type"`

	Payload *scw.JSONObject `json:"payload"`
}

// Email: email.
type Email struct {
	ID string `json:"id"`

	UserID *string `json:"user_id"`

	TemplateName string `json:"template_name"`

	Subject string `json:"subject"`

	Content string `json:"content"`

	// Status: default value: unknown_status
	Status EmailStatus `json:"status"`

	ErrorMessage *string `json:"error_message"`

	SentAt *time.Time `json:"sent_at"`

	// Provider: default value: unknown_provider
	Provider EmailProvider `json:"provider"`
}

// ListEmailEventsRequest: list email events request.
type ListEmailEventsRequest struct {
	EmailID string `json:"-"`

	// OrderBy: default value: time_asc
	OrderBy ListEmailEventsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListEmailEventsResponse: list email events response.
type ListEmailEventsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Events []*Event `json:"events"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListEmailEventsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListEmailEventsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListEmailEventsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Events = append(r.Events, results.Events...)
	r.TotalCount += uint32(len(results.Events))
	return uint32(len(results.Events)), nil
}

// ListEmailsRequest: list emails request.
type ListEmailsRequest struct {
	UserID *string `json:"-"`

	TemplateNames []string `json:"-"`

	Statuses []EmailStatus `json:"-"`

	// OrderBy: default value: sent_at_asc
	OrderBy ListEmailsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Address *string `json:"-"`
}

// ListEmailsResponse: list emails response.
type ListEmailsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Emails []*Email `json:"emails"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListEmailsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListEmailsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListEmailsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Emails = append(r.Emails, results.Emails...)
	r.TotalCount += uint32(len(results.Emails))
	return uint32(len(results.Emails)), nil
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

// ListEmails:
func (s *API) ListEmails(req *ListEmailsRequest, opts ...scw.RequestOption) (*ListEmailsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "user_id", req.UserID)
	parameter.AddToQuery(query, "template_names", req.TemplateNames)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "address", req.Address)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/emailgateway-admin/v1/emails",
		Query:  query,
	}

	var resp ListEmailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListEmailEvents:
func (s *API) ListEmailEvents(req *ListEmailEventsRequest, opts ...scw.RequestOption) (*ListEmailEventsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.EmailID) == "" {
		return nil, errors.New("field EmailID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/emailgateway-admin/v1/emails/" + fmt.Sprint(req.EmailID) + "/events",
		Query:  query,
	}

	var resp ListEmailEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

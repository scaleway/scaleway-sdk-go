// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package support_admin provides methods and message types of the support_admin v1 API.
package support_admin

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

type ListTicketsRequestOrderBy string

const (
	ListTicketsRequestOrderByCreatedAtAsc  = ListTicketsRequestOrderBy("created_at_asc")
	ListTicketsRequestOrderByCreatedAtDesc = ListTicketsRequestOrderBy("created_at_desc")
	ListTicketsRequestOrderByClosedAtAsc   = ListTicketsRequestOrderBy("closed_at_asc")
	ListTicketsRequestOrderByClosedAtDesc  = ListTicketsRequestOrderBy("closed_at_desc")
	ListTicketsRequestOrderByUpdatedAtAsc  = ListTicketsRequestOrderBy("updated_at_asc")
	ListTicketsRequestOrderByUpdatedAtDesc = ListTicketsRequestOrderBy("updated_at_desc")
)

func (enum ListTicketsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListTicketsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTicketsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTicketsRequestOrderBy(ListTicketsRequestOrderBy(tmp).String())
	return nil
}

type Severity string

const (
	SeverityUnknownSeverity = Severity("unknown_severity")
	SeverityHigh            = Severity("high")
	SeverityMedium          = Severity("medium")
	SeverityLow             = Severity("low")
)

func (enum Severity) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_severity"
	}
	return string(enum)
}

func (enum Severity) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Severity) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Severity(Severity(tmp).String())
	return nil
}

type TicketSide string

const (
	TicketSideUnknownTicketSide = TicketSide("unknown_ticket_side")
	TicketSideAwaitingAgent     = TicketSide("awaiting_agent")
	TicketSideAwaitingUser      = TicketSide("awaiting_user")
)

func (enum TicketSide) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_ticket_side"
	}
	return string(enum)
}

func (enum TicketSide) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TicketSide) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TicketSide(TicketSide(tmp).String())
	return nil
}

type TicketStatus string

const (
	TicketStatusUnknownTicketStatus = TicketStatus("unknown_ticket_status")
	TicketStatusOpen                = TicketStatus("open")
	TicketStatusClosed              = TicketStatus("closed")
)

func (enum TicketStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_ticket_status"
	}
	return string(enum)
}

func (enum TicketStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TicketStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TicketStatus(TicketStatus(tmp).String())
	return nil
}

// User: user.
type User struct {
	ID string `json:"id"`

	IsAgent bool `json:"is_agent"`

	Firstname string `json:"firstname"`

	Lastname string `json:"lastname"`

	Fullname string `json:"fullname"`

	Gravatar *string `json:"gravatar"`

	Email *string `json:"email"`

	UserID string `json:"user_id"`

	OrganizationID string `json:"organization_id"`
}

// TicketSummary: ticket summary.
type TicketSummary struct {
	ID string `json:"id"`

	Subject string `json:"subject"`

	CreatedBy *User `json:"created_by"`

	// Status: default value: unknown_ticket_status
	Status TicketStatus `json:"status"`

	// Side: default value: unknown_ticket_side
	Side TicketSide `json:"side"`

	Tags []string `json:"tags"`

	// Severity: default value: unknown_severity
	Severity Severity `json:"severity"`

	Product *string `json:"product"`

	ResourceID *string `json:"resource_id"`

	EscalationLevel int32 `json:"escalation_level"`

	CreatedAt *time.Time `json:"created_at"`

	ClosedAt *time.Time `json:"closed_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	DisplayID string `json:"display_id"`
}

// AgentNote: agent note.
type AgentNote struct {
	TicketID string `json:"ticket_id"`

	Sender *User `json:"sender"`

	Content string `json:"content"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// Answer: answer.
type Answer struct {
	TicketID string `json:"ticket_id"`

	Sender *User `json:"sender"`

	Content string `json:"content"`

	IsAgentSide bool `json:"is_agent_side"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// AddTicketAgentNoteRequest: add ticket agent note request.
type AddTicketAgentNoteRequest struct {
	TicketID string `json:"-"`

	Content string `json:"content"`

	EscalationLevel uint32 `json:"escalation_level"`

	UserID string `json:"user_id"`
}

// AddTicketAnswerRequest: add ticket answer request.
type AddTicketAnswerRequest struct {
	TicketID string `json:"-"`

	Content string `json:"content"`

	// TicketSide: default value: unknown_ticket_side
	TicketSide TicketSide `json:"ticket_side"`

	UserID string `json:"user_id"`
}

// CreateTicketRequest: create ticket request.
type CreateTicketRequest struct {
	OrganizationID string `json:"organization_id"`

	UserID string `json:"user_id"`

	Subject string `json:"subject"`

	Content string `json:"content"`

	Tags []string `json:"tags"`

	// Severity: default value: unknown_severity
	Severity Severity `json:"severity"`

	ResourceID *string `json:"resource_id,omitempty"`

	Product *string `json:"product,omitempty"`
}

// GetTicketRequest: get ticket request.
type GetTicketRequest struct {
	TicketID string `json:"-"`
}

// ListTicketsRequest: list tickets request.
type ListTicketsRequest struct {
	OrganizationID *string `json:"-"`

	ResourceID *string `json:"-"`

	Subject *string `json:"-"`

	// Status: default value: unknown_ticket_status
	Status TicketStatus `json:"-"`

	// Side: default value: unknown_ticket_side
	Side TicketSide `json:"-"`

	Product *string `json:"-"`

	Tag *string `json:"-"`

	// Severity: default value: unknown_severity
	Severity Severity `json:"-"`

	EscalationLevel *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListTicketsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListTicketsResponse: list tickets response.
type ListTicketsResponse struct {
	Tickets []*TicketSummary `json:"tickets"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTicketsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTicketsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTicketsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tickets = append(r.Tickets, results.Tickets...)
	r.TotalCount += uint32(len(results.Tickets))
	return uint32(len(results.Tickets)), nil
}

// Ticket: ticket.
type Ticket struct {
	ID string `json:"id"`

	Subject string `json:"subject"`

	CreatedBy *User `json:"created_by"`

	// Status: default value: unknown_ticket_status
	Status TicketStatus `json:"status"`

	// Side: default value: unknown_ticket_side
	Side TicketSide `json:"side"`

	Tags []string `json:"tags"`

	// Severity: default value: unknown_severity
	Severity Severity `json:"severity"`

	Product *string `json:"product"`

	ResourceID *string `json:"resource_id"`

	EscalationLevel int32 `json:"escalation_level"`

	Answers []*Answer `json:"answers"`

	AgentNotes []*AgentNote `json:"agent_notes"`

	CreatedAt *time.Time `json:"created_at"`

	ClosedAt *time.Time `json:"closed_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	DisplayID string `json:"display_id"`
}

// UpdateTicketRequest: update ticket request.
type UpdateTicketRequest struct {
	TicketID string `json:"-"`

	Subject *string `json:"subject,omitempty"`

	Tags []string `json:"tags"`

	// Severity: default value: unknown_severity
	Severity Severity `json:"severity"`

	Product *string `json:"product,omitempty"`

	// Status: default value: unknown_ticket_status
	Status TicketStatus `json:"status"`

	// Side: default value: unknown_ticket_side
	Side TicketSide `json:"side"`

	EscalationLevel *uint32 `json:"escalation_level,omitempty"`

	ResourceID *string `json:"resource_id,omitempty"`
}

// This service implement ticket processing between Customers and Support.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ListTickets: List all tickets. Filterable.
func (s *API) ListTickets(req *ListTicketsRequest, opts ...scw.RequestOption) (*ListTicketsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "resource_id", req.ResourceID)
	parameter.AddToQuery(query, "subject", req.Subject)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "side", req.Side)
	parameter.AddToQuery(query, "product", req.Product)
	parameter.AddToQuery(query, "tag", req.Tag)
	parameter.AddToQuery(query, "severity", req.Severity)
	parameter.AddToQuery(query, "escalation_level", req.EscalationLevel)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/support-admin/v1/tickets",
		Query:  query,
	}

	var resp ListTicketsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTicket: Get a ticket's data.
func (s *API) GetTicket(req *GetTicketRequest, opts ...scw.RequestOption) (*Ticket, error) {
	var err error

	if fmt.Sprint(req.TicketID) == "" {
		return nil, errors.New("field TicketID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/support-admin/v1/tickets/" + fmt.Sprint(req.TicketID) + "",
	}

	var resp Ticket

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTicket: Create a ticket.
func (s *API) CreateTicket(req *CreateTicketRequest, opts ...scw.RequestOption) (*Ticket, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/support-admin/v1/tickets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Ticket

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateTicket: Edit a ticket's metadata.
func (s *API) UpdateTicket(req *UpdateTicketRequest, opts ...scw.RequestOption) (*Ticket, error) {
	var err error

	if fmt.Sprint(req.TicketID) == "" {
		return nil, errors.New("field TicketID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/support-admin/v1/tickets/" + fmt.Sprint(req.TicketID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Ticket

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddTicketAnswer: Add an answer to a ticket's conversation. Optionally change ticket's side.
func (s *API) AddTicketAnswer(req *AddTicketAnswerRequest, opts ...scw.RequestOption) (*Ticket, error) {
	var err error

	if fmt.Sprint(req.TicketID) == "" {
		return nil, errors.New("field TicketID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/support-admin/v1/tickets/" + fmt.Sprint(req.TicketID) + "/answer",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Ticket

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddTicketAgentNote: Add an internal note to a ticket's conversation. Optionally escalate ticket.
func (s *API) AddTicketAgentNote(req *AddTicketAgentNoteRequest, opts ...scw.RequestOption) (*Ticket, error) {
	var err error

	if fmt.Sprint(req.TicketID) == "" {
		return nil, errors.New("field TicketID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/support-admin/v1/tickets/" + fmt.Sprint(req.TicketID) + "/agentnote",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Ticket

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

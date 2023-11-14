// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package support provides methods and message types of the support v2 API.
package support

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

// Answer: answer.
type Answer struct {
	TicketID string `json:"ticket_id"`

	Content string `json:"content"`

	IsAgentSide bool `json:"is_agent_side"`

	CreatedByID string `json:"created_by_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	CreatedByName string `json:"created_by_name"`
}

// Ticket: ticket.
type Ticket struct {
	ID string `json:"id"`

	DisplayID string `json:"display_id"`

	OrganizationID *string `json:"organization_id"`

	UserID *string `json:"user_id"`

	DediboxID *string `json:"dedibox_id"`

	Subject string `json:"subject"`

	// Status: default value: unknown_ticket_status
	Status TicketStatus `json:"status"`

	// Side: default value: unknown_ticket_side
	Side TicketSide `json:"side"`

	// Severity: default value: unknown_severity
	Severity Severity `json:"severity"`

	Product *string `json:"product"`

	ResourceID *string `json:"resource_id"`

	EscalationLevel int32 `json:"escalation_level"`

	CsatScore *uint32 `json:"csat_score"`

	CsatComment *string `json:"csat_comment"`

	Answers []*Answer `json:"answers"`

	CreatedAt *time.Time `json:"created_at"`

	ClosedAt *time.Time `json:"closed_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// AddTicketAnswerRequest: add ticket answer request.
type AddTicketAnswerRequest struct {
	TicketID string `json:"-"`

	Content string `json:"content"`

	// TicketSide: default value: unknown_ticket_side
	TicketSide TicketSide `json:"ticket_side"`

	UserID string `json:"user_id"`
}

// CloseTicketRequest: close ticket request.
type CloseTicketRequest struct {
	TicketID string `json:"-"`

	CsatScore *uint32 `json:"csat_score,omitempty"`

	CsatComment *string `json:"csat_comment,omitempty"`
}

// CreateJwtRequest: create jwt request.
type CreateJwtRequest struct {
	OrganizationID string `json:"organization_id"`

	UserID string `json:"user_id"`
}

// CreateJwtResponse: create jwt response.
type CreateJwtResponse struct {
	Jwt string `json:"jwt"`
}

// CreateTicketRequest: create ticket request.
type CreateTicketRequest struct {
	OrganizationID string `json:"organization_id"`

	UserID string `json:"user_id"`

	Subject string `json:"subject"`

	Content string `json:"content"`

	// Severity: default value: unknown_severity
	Severity Severity `json:"severity"`

	ResourceID *string `json:"resource_id,omitempty"`

	Product *string `json:"product,omitempty"`
}

// EscalateTicketRequest: escalate ticket request.
type EscalateTicketRequest struct {
	TicketID string `json:"-"`

	EscalationLevel int32 `json:"escalation_level"`
}

// GetTicketRequest: get ticket request.
type GetTicketRequest struct {
	TicketID string `json:"-"`
}

// ListTicketsRequest: list tickets request.
type ListTicketsRequest struct {
	OrganizationID string `json:"-"`

	ResourceID *string `json:"-"`

	Subject *string `json:"-"`

	// Status: default value: unknown_ticket_status
	Status TicketStatus `json:"-"`

	// Side: default value: unknown_ticket_side
	Side TicketSide `json:"-"`

	Product *string `json:"-"`

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
	Tickets []*Ticket `json:"tickets"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTicketsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTicketsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListTicketsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tickets = append(r.Tickets, results.Tickets...)
	r.TotalCount += uint64(len(results.Tickets))
	return uint64(len(results.Tickets)), nil
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

// CreateJwt: Creates a JWT for the Livechat.
func (s *API) CreateJwt(req *CreateJwtRequest, opts ...scw.RequestOption) (*CreateJwtResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/support/v2/jwt",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateJwtResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTickets: List all tickets for a customer's organization. Filterable.
func (s *API) ListTickets(req *ListTicketsRequest, opts ...scw.RequestOption) (*ListTicketsResponse, error) {
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
	parameter.AddToQuery(query, "subject", req.Subject)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "side", req.Side)
	parameter.AddToQuery(query, "product", req.Product)
	parameter.AddToQuery(query, "severity", req.Severity)
	parameter.AddToQuery(query, "escalation_level", req.EscalationLevel)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/support/v2/tickets",
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
		Path:   "/support/v2/tickets/" + fmt.Sprint(req.TicketID) + "",
	}

	var resp Ticket

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTicket: Create a new ticket.
func (s *API) CreateTicket(req *CreateTicketRequest, opts ...scw.RequestOption) (*Ticket, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/support/v2/tickets",
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
		Path:   "/support/v2/tickets/" + fmt.Sprint(req.TicketID) + "/add-answer",
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

// CloseTicket: Closes a ticket, specifying customer satisfaction.
func (s *API) CloseTicket(req *CloseTicketRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.TicketID) == "" {
		return errors.New("field TicketID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/support/v2/tickets/" + fmt.Sprint(req.TicketID) + "/close",
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

// EscalateTicket: Sets the escalation level for a ticket.
func (s *API) EscalateTicket(req *EscalateTicketRequest, opts ...scw.RequestOption) (*Ticket, error) {
	var err error

	if fmt.Sprint(req.TicketID) == "" {
		return nil, errors.New("field TicketID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/support/v2/tickets/" + fmt.Sprint(req.TicketID) + "/escalate",
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

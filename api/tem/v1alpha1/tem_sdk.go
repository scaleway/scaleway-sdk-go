// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package tem provides methods and message types of the tem v1alpha1 API.
package tem

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

// API: tem
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type DomainStatus string

const (
	// DomainStatusUnknown is [insert doc].
	DomainStatusUnknown = DomainStatus("unknown")
	// DomainStatusChecked is [insert doc].
	DomainStatusChecked = DomainStatus("checked")
	// DomainStatusUnchecked is [insert doc].
	DomainStatusUnchecked = DomainStatus("unchecked")
	// DomainStatusInvalid is [insert doc].
	DomainStatusInvalid = DomainStatus("invalid")
	// DomainStatusLocked is [insert doc].
	DomainStatusLocked = DomainStatus("locked")
	// DomainStatusRevoked is [insert doc].
	DomainStatusRevoked = DomainStatus("revoked")
	// DomainStatusPending is [insert doc].
	DomainStatusPending = DomainStatus("pending")
)

func (enum DomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DomainStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainStatus(DomainStatus(tmp).String())
	return nil
}

type EmailRcptType string

const (
	// EmailRcptTypeUnknownRcptType is [insert doc].
	EmailRcptTypeUnknownRcptType = EmailRcptType("unknown_rcpt_type")
	// EmailRcptTypeTo is [insert doc].
	EmailRcptTypeTo = EmailRcptType("to")
	// EmailRcptTypeCc is [insert doc].
	EmailRcptTypeCc = EmailRcptType("cc")
	// EmailRcptTypeBcc is [insert doc].
	EmailRcptTypeBcc = EmailRcptType("bcc")
)

func (enum EmailRcptType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_rcpt_type"
	}
	return string(enum)
}

func (enum EmailRcptType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EmailRcptType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EmailRcptType(EmailRcptType(tmp).String())
	return nil
}

type EmailStatus string

const (
	// EmailStatusUnknown is [insert doc].
	EmailStatusUnknown = EmailStatus("unknown")
	// EmailStatusNew is [insert doc].
	EmailStatusNew = EmailStatus("new")
	// EmailStatusSending is [insert doc].
	EmailStatusSending = EmailStatus("sending")
	// EmailStatusSent is [insert doc].
	EmailStatusSent = EmailStatus("sent")
	// EmailStatusFailed is [insert doc].
	EmailStatusFailed = EmailStatus("failed")
	// EmailStatusCanceled is [insert doc].
	EmailStatusCanceled = EmailStatus("canceled")
)

func (enum EmailStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
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

// CreateEmailRequestAddress: create email request. address
type CreateEmailRequestAddress struct {
	// Email: email address
	Email string `json:"email"`
	// Name: optional display name
	Name *string `json:"name"`
}

// CreateEmailRequestAttachment: create email request. attachment
type CreateEmailRequestAttachment struct {
	// Name: filename of the attachment
	Name string `json:"name"`
	// Type: mIME type of the attachment (Currently only allow, text files, pdf and html files)
	Type string `json:"type"`
	// Content: content of the attachment, encoded in base64
	Content []byte `json:"content"`
}

// CreateEmailResponse: create email response
type CreateEmailResponse struct {
	// Emails: single page of emails matching the requested criteria
	Emails []*Email `json:"emails"`
}

// Domain: domain
type Domain struct {
	// ID: ID of the domain
	ID string `json:"id"`
	// OrganizationID: ID of the organization to which the domain belongs
	OrganizationID string `json:"organization_id"`
	// ProjectID: ID of the project
	ProjectID string `json:"project_id"`
	// Name: domain name (example.com)
	Name string `json:"name"`
	// Status: status of the domain
	//
	// Default value: unknown
	Status DomainStatus `json:"status"`
	// CreatedAt: date and time of domain's creation
	CreatedAt *time.Time `json:"created_at"`
	// NextCheckAt: date and time of the next scheduled check
	NextCheckAt *time.Time `json:"next_check_at"`
	// LastValidAt: date and time the domain was last found to be valid
	LastValidAt *time.Time `json:"last_valid_at"`
	// RevokedAt: date and time of the revocation of the domain
	RevokedAt *time.Time `json:"revoked_at"`
	// LastError: error message if the last check failed
	LastError *string `json:"last_error"`
	// SpfConfig: snippet of the SPF record that should be registered in the DNS zone
	SpfConfig string `json:"spf_config"`
	// DkimConfig: dKIM public key, as should be recorded in the DNS zone
	DkimConfig string `json:"dkim_config"`
	// Statistics: domain's statistics
	Statistics *DomainStatistics `json:"statistics"`

	Region scw.Region `json:"region"`
}

type DomainStatistics struct {
	TotalCount uint32 `json:"total_count"`

	SentCount uint32 `json:"sent_count"`

	FailedCount uint32 `json:"failed_count"`

	CanceledCount uint32 `json:"canceled_count"`
}

// Email: email
type Email struct {
	// ID: technical ID of the email
	ID string `json:"id"`
	// MessageID: messageID of the email
	MessageID string `json:"message_id"`
	// ProjectID: ID of the project to which the email belongs
	ProjectID string `json:"project_id"`
	// MailFrom: email address of the sender
	MailFrom string `json:"mail_from"`
	// RcptTo: email address of the recipient
	RcptTo string `json:"rcpt_to"`
	// RcptType: type of the recipient
	//
	// Default value: unknown_rcpt_type
	RcptType EmailRcptType `json:"rcpt_type"`
	// CreatedAt: creation date of the email object
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: last update time of the email object
	UpdatedAt *time.Time `json:"updated_at"`
	// Status: status of the email
	//
	// Default value: unknown
	Status EmailStatus `json:"status"`
	// StatusDetails: additional information on the status
	StatusDetails *string `json:"status_details"`
	// TryCount: total number of attempts to send the email
	TryCount uint32 `json:"try_count"`
	// LastTries: informations about the latest three attempts to send the email
	LastTries []*EmailTry `json:"last_tries"`
}

// EmailTry: email. try
type EmailTry struct {
	// Rank: rank number of this attempt to send the email
	Rank uint32 `json:"rank"`
	// TriedAt: date of the attempt
	TriedAt *time.Time `json:"tried_at"`
	// Code: the SMTP status code received after the attempt. 0 if the attempt did not reach an SMTP server.
	Code int32 `json:"code"`
	// Message: the SMTP message received, if any. If the attempt did not reach an SMTP server, the message says why.
	Message string `json:"message"`
}

// ListDomainsResponse: list domains response
type ListDomainsResponse struct {
	// TotalCount: total number of domains matching the request (without pagination)
	TotalCount uint32 `json:"total_count"`

	Domains []*Domain `json:"domains"`
}

// ListEmailsResponse: list emails response
type ListEmailsResponse struct {
	// TotalCount: count of all emails matching the requested criteria
	TotalCount uint32 `json:"total_count"`
	// Emails: single page of emails matching the requested criteria
	Emails []*Email `json:"emails"`
}

// Statistics: statistics
type Statistics struct {
	// TotalCount: total number of emails matching the request criteria
	TotalCount uint32 `json:"total_count"`
	// NewCount: number of emails still in the `new` transient state (received from the API, not yet processed)
	NewCount uint32 `json:"new_count"`
	// SendingCount: number of emails still in the `sending` transient state (received from the API, not yet in their final status)
	SendingCount uint32 `json:"sending_count"`
	// SentCount: number of emails in the final `sent` state (have been delivered to the target mail system)
	SentCount uint32 `json:"sent_count"`
	// FailedCount: number of emails in the final `failed` state (refused by the target mail system with a final error status)
	FailedCount uint32 `json:"failed_count"`
	// CanceledCount: number of emails in the final `canceled` state (canceled by customer's request)
	CanceledCount uint32 `json:"canceled_count"`
}

// Service API

// Regions list localities the api is available in
func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar}
}

type CreateEmailRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// From: sender information (must be from a checked domain declared in the project)
	From *CreateEmailRequestAddress `json:"from"`
	// To: array of recipient information (limited to 1 recipient)
	To []*CreateEmailRequestAddress `json:"to"`
	// Cc: array of recipient information (unimplemented)
	Cc []*CreateEmailRequestAddress `json:"cc"`
	// Bcc: array of recipient information (unimplemented)
	Bcc []*CreateEmailRequestAddress `json:"bcc"`
	// Subject: message subject
	Subject string `json:"subject"`
	// Text: text content
	Text string `json:"text"`
	// HTML: HTML content
	HTML string `json:"html"`
	// ProjectID: ID of the project in which to create the email
	ProjectID string `json:"project_id"`
	// Attachments: array of attachments
	Attachments []*CreateEmailRequestAttachment `json:"attachments"`
	// SendBefore: maximum date to deliver mail
	SendBefore *time.Time `json:"send_before"`
}

// CreateEmail: send an email
func (s *API) CreateEmail(req *CreateEmailRequest, opts ...scw.RequestOption) (*CreateEmailResponse, error) {
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
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateEmailResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetEmailRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// EmailID: ID of the email to retrieve
	EmailID string `json:"-"`
}

// GetEmail: get information about an email
func (s *API) GetEmail(req *GetEmailRequest, opts ...scw.RequestOption) (*Email, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.EmailID) == "" {
		return nil, errors.New("field EmailID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails/" + fmt.Sprint(req.EmailID) + "",
		Headers: http.Header{},
	}

	var resp Email

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListEmailsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// ProjectID: optional ID of the project in which to list the emails
	ProjectID *string `json:"-"`
	// DomainID: optional ID of the domain for which to list the emails
	DomainID *string `json:"-"`
	// MessageID: optional ID of the message for which to list the emails
	MessageID *string `json:"-"`
	// Since: optional, list emails created after this date
	Since *time.Time `json:"-"`
	// Until: optional, list emails created before this date
	Until *time.Time `json:"-"`
	// MailFrom: optional, list emails sent with this `mail_from` sender's address
	MailFrom *string `json:"-"`
	// MailTo: optional, list emails sent with this `mail_to` recipient's address
	MailTo *string `json:"-"`
	// Statuses: optional, list emails having any of this status
	Statuses []EmailStatus `json:"-"`
}

// ListEmails: list emails sent from a domain and/or for a project and/or for an organization
func (s *API) ListEmails(req *ListEmailsRequest, opts ...scw.RequestOption) (*ListEmailsResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "domain_id", req.DomainID)
	parameter.AddToQuery(query, "message_id", req.MessageID)
	parameter.AddToQuery(query, "since", req.Since)
	parameter.AddToQuery(query, "until", req.Until)
	parameter.AddToQuery(query, "mail_from", req.MailFrom)
	parameter.AddToQuery(query, "mail_to", req.MailTo)
	parameter.AddToQuery(query, "statuses", req.Statuses)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListEmailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetStatisticsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// ProjectID: optional, count emails for this project
	ProjectID *string `json:"-"`
	// DomainID: optional, count emails send from this domain (must be coherent with the `project_id` and the `organization_id`)
	DomainID *string `json:"-"`
	// Since: optional, count emails created after this date
	Since *time.Time `json:"-"`
	// Until: optional, count emails created before this date
	Until *time.Time `json:"-"`
	// MailFrom: optional, count emails sent with this `mail_from` sender's address
	MailFrom *string `json:"-"`
}

// GetStatistics: get statistics on the email statuses
func (s *API) GetStatistics(req *GetStatisticsRequest, opts ...scw.RequestOption) (*Statistics, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "domain_id", req.DomainID)
	parameter.AddToQuery(query, "since", req.Since)
	parameter.AddToQuery(query, "until", req.Until)
	parameter.AddToQuery(query, "mail_from", req.MailFrom)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/statistics",
		Query:   query,
		Headers: http.Header{},
	}

	var resp Statistics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CancelEmailRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// EmailID: ID of the email to cancel
	EmailID string `json:"-"`
}

// CancelEmail: try to cancel an email if it has not yet been sent
func (s *API) CancelEmail(req *CancelEmailRequest, opts ...scw.RequestOption) (*Email, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.EmailID) == "" {
		return nil, errors.New("field EmailID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails/" + fmt.Sprint(req.EmailID) + "/cancel",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Email

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateDomainRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`

	ProjectID string `json:"project_id"`

	DomainName string `json:"domain_name"`
}

// CreateDomain: register a domain in a project
func (s *API) CreateDomain(req *CreateDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
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
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDomainRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// DomainID: ID of the domain
	DomainID string `json:"-"`
}

// GetDomain: get information about a domain
func (s *API) GetDomain(req *GetDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
		Headers: http.Header{},
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListDomainsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// Page: page number (1 for the first page)
	Page *int32 `json:"-"`
	// PageSize: page size
	PageSize *uint32 `json:"-"`

	ProjectID *string `json:"-"`

	Status []DomainStatus `json:"-"`

	OrganizationID *string `json:"-"`

	Name *string `json:"-"`
}

// ListDomains: list domains in a project and/or in an organization
func (s *API) ListDomains(req *ListDomainsRequest, opts ...scw.RequestOption) (*ListDomainsResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RevokeDomainRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// DomainID: ID of the domain to revoke
	DomainID string `json:"-"`
}

// RevokeDomain: revoke a domain
func (s *API) RevokeDomain(req *RevokeDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/revoke",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CheckDomainRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// DomainID: ID of the domain to check
	DomainID string `json:"-"`
}

// CheckDomain: ask for an immediate check of a domain (DNS check)
func (s *API) CheckDomain(req *CheckDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/check",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDomainsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Domains = append(r.Domains, results.Domains...)
	r.TotalCount += uint32(len(results.Domains))
	return uint32(len(results.Domains)), nil
}

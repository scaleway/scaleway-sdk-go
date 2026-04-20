// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package mailbox provides methods and message types of the mailbox v1alpha1 API.
package mailbox

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultMailboxRetryInterval = 15 * time.Second
	defaultMailboxTimeout       = 5 * time.Minute
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

type DomainRecordDNSType string

const (
	// If unspecified, the record's DNS type is unknown by default.
	DomainRecordDNSTypeUnknownDNSType = DomainRecordDNSType("unknown_dns_type")
	// CNAME DNS record type.
	DomainRecordDNSTypeCnameDNSType = DomainRecordDNSType("cname_dns_type")
	// MX DNS record type.
	DomainRecordDNSTypeMxDNSType = DomainRecordDNSType("mx_dns_type")
	// SRV DNS record type.
	DomainRecordDNSTypeSrvDNSType = DomainRecordDNSType("srv_dns_type")
	// TXT DNS record type.
	DomainRecordDNSTypeTxtDNSType = DomainRecordDNSType("txt_dns_type")
)

func (enum DomainRecordDNSType) String() string {
	if enum == "" {
		// return default value if empty
		return string(DomainRecordDNSTypeUnknownDNSType)
	}
	return string(enum)
}

func (enum DomainRecordDNSType) Values() []DomainRecordDNSType {
	return []DomainRecordDNSType{
		"unknown_dns_type",
		"cname_dns_type",
		"mx_dns_type",
		"srv_dns_type",
		"txt_dns_type",
	}
}

func (enum DomainRecordDNSType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainRecordDNSType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainRecordDNSType(DomainRecordDNSType(tmp).String())
	return nil
}

type DomainRecordLevel string

const (
	// If unspecified, the record level is unknown by default.
	DomainRecordLevelUnknownLevel = DomainRecordLevel("unknown_level")
	// A record is required to validate the domain.
	DomainRecordLevelRequired = DomainRecordLevel("required")
	// To have the best experience available and ease the integration of your mailboxes.
	DomainRecordLevelRecommended = DomainRecordLevel("recommended")
	// A record is optional.
	DomainRecordLevelOptional = DomainRecordLevel("optional")
)

func (enum DomainRecordLevel) String() string {
	if enum == "" {
		// return default value if empty
		return string(DomainRecordLevelUnknownLevel)
	}
	return string(enum)
}

func (enum DomainRecordLevel) Values() []DomainRecordLevel {
	return []DomainRecordLevel{
		"unknown_level",
		"required",
		"recommended",
		"optional",
	}
}

func (enum DomainRecordLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainRecordLevel) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainRecordLevel(DomainRecordLevel(tmp).String())
	return nil
}

type DomainRecordStatus string

const (
	// If unspecified, the status is unknown by default.
	DomainRecordStatusUnknownStatus = DomainRecordStatus("unknown_status")
	// The record is in the process of being validated.
	DomainRecordStatusValidating = DomainRecordStatus("validating")
	// The record is found in the domain's DNS zone and is valid.
	DomainRecordStatusValid = DomainRecordStatus("valid")
	// A detailed error can be found in the "error" field of a DNS record.
	DomainRecordStatusInvalid = DomainRecordStatus("invalid")
	// The record cannot be found in the domain's DNS zone.
	DomainRecordStatusNotFound = DomainRecordStatus("not_found")
)

func (enum DomainRecordStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(DomainRecordStatusUnknownStatus)
	}
	return string(enum)
}

func (enum DomainRecordStatus) Values() []DomainRecordStatus {
	return []DomainRecordStatus{
		"unknown_status",
		"validating",
		"valid",
		"invalid",
		"not_found",
	}
}

func (enum DomainRecordStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainRecordStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainRecordStatus(DomainRecordStatus(tmp).String())
	return nil
}

type DomainStatus string

const (
	// If unspecified, the status is unknown by default.
	DomainStatusUnknownStatus = DomainStatus("unknown_status")
	// The domain was just registered, but the validation records are not yet available.
	DomainStatusCreating = DomainStatus("creating")
	// Mandatory DNS records must be set on the domain's zone and a validation must be triggered to validate the domain.
	DomainStatusWaitingValidation = DomainStatus("waiting_validation")
	// The domain is in the process of validating.
	DomainStatusValidating       = DomainStatus("validating")
	DomainStatusValidationFailed = DomainStatus("validation_failed")
	// The domain is validated and in the process of being provisioned.
	DomainStatusProvisioning = DomainStatus("provisioning")
	// The domain is ready to be used and accept mailbox creations.
	DomainStatusReady = DomainStatus("ready")
	// The domain is in the process of being deleted.
	DomainStatusDeleting = DomainStatus("deleting")
)

func (enum DomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(DomainStatusUnknownStatus)
	}
	return string(enum)
}

func (enum DomainStatus) Values() []DomainStatus {
	return []DomainStatus{
		"unknown_status",
		"creating",
		"waiting_validation",
		"validating",
		"validation_failed",
		"provisioning",
		"ready",
		"deleting",
	}
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

type ListDomainsRequestOrderBy string

const (
	// Order by creation date (descending chronological order).
	ListDomainsRequestOrderByCreatedAtDesc = ListDomainsRequestOrderBy("created_at_desc")
	// Order by creation date (ascending chronological order).
	ListDomainsRequestOrderByCreatedAtAsc = ListDomainsRequestOrderBy("created_at_asc")
	// Order by last update date (descending chronological order).
	ListDomainsRequestOrderByUpdatedAtDesc = ListDomainsRequestOrderBy("updated_at_desc")
	// Order by last update date (ascending chronological order).
	ListDomainsRequestOrderByUpdatedAtAsc = ListDomainsRequestOrderBy("updated_at_asc")
	// Order by name (descending alphabetical order).
	ListDomainsRequestOrderByNameDesc = ListDomainsRequestOrderBy("name_desc")
	// Order by name (ascending alphabetical order).
	ListDomainsRequestOrderByNameAsc = ListDomainsRequestOrderBy("name_asc")
	// Order by mailbox total count (descending numerical order).
	ListDomainsRequestOrderByMailboxTotalCountDesc = ListDomainsRequestOrderBy("mailbox_total_count_desc")
	// Order by name (ascending numerical order).
	ListDomainsRequestOrderByMailboxTotalCountAsc = ListDomainsRequestOrderBy("mailbox_total_count_asc")
)

func (enum ListDomainsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListDomainsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListDomainsRequestOrderBy) Values() []ListDomainsRequestOrderBy {
	return []ListDomainsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"updated_at_desc",
		"updated_at_asc",
		"name_desc",
		"name_asc",
		"mailbox_total_count_desc",
		"mailbox_total_count_asc",
	}
}

func (enum ListDomainsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDomainsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDomainsRequestOrderBy(ListDomainsRequestOrderBy(tmp).String())
	return nil
}

type ListMailboxesRequestOrderBy string

const (
	ListMailboxesRequestOrderByCreatedAtDesc = ListMailboxesRequestOrderBy("created_at_desc")
	ListMailboxesRequestOrderByCreatedAtAsc  = ListMailboxesRequestOrderBy("created_at_asc")
	ListMailboxesRequestOrderByUpdatedAtDesc = ListMailboxesRequestOrderBy("updated_at_desc")
	ListMailboxesRequestOrderByUpdatedAtAsc  = ListMailboxesRequestOrderBy("updated_at_asc")
	ListMailboxesRequestOrderByEmailDesc     = ListMailboxesRequestOrderBy("email_desc")
	ListMailboxesRequestOrderByEmailAsc      = ListMailboxesRequestOrderBy("email_asc")
)

func (enum ListMailboxesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListMailboxesRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListMailboxesRequestOrderBy) Values() []ListMailboxesRequestOrderBy {
	return []ListMailboxesRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"updated_at_desc",
		"updated_at_asc",
		"email_desc",
		"email_asc",
	}
}

func (enum ListMailboxesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListMailboxesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListMailboxesRequestOrderBy(ListMailboxesRequestOrderBy(tmp).String())
	return nil
}

type MailboxStatus string

const (
	// If unspecified, the status is unknown by default.
	MailboxStatusUnknownStatus = MailboxStatus("unknown_status")
	// The mailbox is being created.
	MailboxStatusCreating = MailboxStatus("creating")
	// The mailbox is waiting for payment.
	MailboxStatusWaitingPayment = MailboxStatus("waiting_payment")
	// The mailbox is waiting for the domain to be properly configured.
	MailboxStatusWaitingDomain = MailboxStatus("waiting_domain")
	// The mailbox is ready to use.
	MailboxStatusReady = MailboxStatus("ready")
	// The mailbox is planned to be deleted.
	MailboxStatusDeletionScheduled = MailboxStatus("deletion_scheduled")
	// The mailbox is locked (e.g billing issue).
	MailboxStatusLocked = MailboxStatus("locked")
	// The mailbox is renewing the payment for a new period.
	MailboxStatusRenewing  = MailboxStatus("renewing")
	MailboxStatusDeleting  = MailboxStatus("deleting")
	MailboxStatusRestoring = MailboxStatus("restoring")
	// The payment of the mailbox failed.
	MailboxStatusPaymentFailed = MailboxStatus("payment_failed")
)

func (enum MailboxStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(MailboxStatusUnknownStatus)
	}
	return string(enum)
}

func (enum MailboxStatus) Values() []MailboxStatus {
	return []MailboxStatus{
		"unknown_status",
		"creating",
		"waiting_payment",
		"waiting_domain",
		"ready",
		"deletion_scheduled",
		"locked",
		"renewing",
		"deleting",
		"restoring",
		"payment_failed",
	}
}

func (enum MailboxStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MailboxStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MailboxStatus(MailboxStatus(tmp).String())
	return nil
}

type MailboxSubscriptionPeriod string

const (
	// If unspecified, the subscription period is unknown by default.
	MailboxSubscriptionPeriodUnknownSubscriptionPeriod = MailboxSubscriptionPeriod("unknown_subscription_period")
	// The subscription has been canceled.
	MailboxSubscriptionPeriodCanceled = MailboxSubscriptionPeriod("canceled")
	// The subscription is monthly.
	MailboxSubscriptionPeriodMonthly = MailboxSubscriptionPeriod("monthly")
	// The subscription is yearly.
	MailboxSubscriptionPeriodYearly = MailboxSubscriptionPeriod("yearly")
)

func (enum MailboxSubscriptionPeriod) String() string {
	if enum == "" {
		// return default value if empty
		return string(MailboxSubscriptionPeriodUnknownSubscriptionPeriod)
	}
	return string(enum)
}

func (enum MailboxSubscriptionPeriod) Values() []MailboxSubscriptionPeriod {
	return []MailboxSubscriptionPeriod{
		"unknown_subscription_period",
		"canceled",
		"monthly",
		"yearly",
	}
}

func (enum MailboxSubscriptionPeriod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MailboxSubscriptionPeriod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MailboxSubscriptionPeriod(MailboxSubscriptionPeriod(tmp).String())
	return nil
}

// BatchCreateMailboxesRequestMailboxParameters: batch create mailboxes request mailbox parameters.
type BatchCreateMailboxesRequestMailboxParameters struct {
	LocalPart string `json:"local_part"`

	Password string `json:"password"`
}

// Mailbox: mailbox.
type Mailbox struct {
	// ID: unique identifier of the mailbox.
	ID string `json:"id"`

	// DomainID: ID of the domain to which the mailbox belongs.
	DomainID string `json:"domain_id"`

	// Email: email address of the mailbox as local_part@domain.
	Email string `json:"email"`

	// Status: status of the mailbox.
	// Default value: unknown_status
	Status MailboxStatus `json:"status"`

	// SubscriptionPeriod: subscription renewal period, it can be monthly, yearly or canceled if the subscription has been canceled.
	// Default value: unknown_subscription_period
	SubscriptionPeriod MailboxSubscriptionPeriod `json:"subscription_period"`

	// SubscriptionPeriodStartedAt: date and time of subscription period start.
	SubscriptionPeriodStartedAt *time.Time `json:"subscription_period_started_at"`

	// NextSubscriptionPeriod: next subscription renewal period, it can be monthly or yearly or canceled if the subscription has been canceled.
	// Default value: unknown_subscription_period
	NextSubscriptionPeriod MailboxSubscriptionPeriod `json:"next_subscription_period"`

	// NextSubscriptionPeriodStartsAt: date and time when the next subscription period starts.
	NextSubscriptionPeriodStartsAt *time.Time `json:"next_subscription_period_starts_at"`

	// CreatedAt: date and time of mailbox creation.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date and time when the mailbox was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// DeletionScheduledAt: date and time of the unrecoverable mailbox deletion.
	DeletionScheduledAt *time.Time `json:"deletion_scheduled_at"`
}

// DomainRecord: domain record.
type DomainRecord struct {
	// ID: unique identifier of the DNS record.
	ID string `json:"id"`

	// DomainID: ID of the domain to which the record belongs.
	DomainID string `json:"domain_id"`

	// Status: status of the record.
	// Default value: unknown_status
	Status DomainRecordStatus `json:"status"`

	// Level: level of requirement of the record.
	// Default value: unknown_level
	Level DomainRecordLevel `json:"level"`

	// DNSType: DNS type of the record.
	// Default value: unknown_dns_type
	DNSType DomainRecordDNSType `json:"dns_type"`

	// DNSName: DNS name of the record.
	DNSName string `json:"dns_name"`

	// DNSValue: DNS value of the record.
	DNSValue string `json:"dns_value"`

	// Error: error associated in case the record is not valid.
	Error *string `json:"error"`

	// CreatedAt: date and time of record creation.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date and time of record last updated.
	UpdatedAt *time.Time `json:"updated_at"`
}

// Domain: domain.
type Domain struct {
	// ID: unique identifier of the domain.
	ID string `json:"id"`

	// ProjectID: ID of the Project to which the domain belongs.
	ProjectID string `json:"project_id"`

	// Name: fully qualified domain name.
	Name string `json:"name"`

	// Status: status of the domain.
	// Default value: unknown_status
	Status DomainStatus `json:"status"`

	// MailboxTotalCount: number of mailboxes of the domain.
	MailboxTotalCount uint64 `json:"mailbox_total_count"`

	// CreatedAt: date and time of domain creation.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date and time of the domain's last update.
	UpdatedAt *time.Time `json:"updated_at"`

	// WebmailURL: URL of the domain's webmail.
	WebmailURL string `json:"webmail_url"`

	// ImapURL: URL of the domain's IMAP service.
	ImapURL string `json:"imap_url"`

	// Pop3URL: URL of the domain's POP3 service.
	Pop3URL string `json:"pop3_url"`

	// SMTPURL: URL of the domain's SMTP service.
	SMTPURL string `json:"smtp_url"`
}

// BatchCreateMailboxesRequest: batch create mailboxes request.
type BatchCreateMailboxesRequest struct {
	// Mailboxes: parameters for the mailboxes to create.
	Mailboxes []*BatchCreateMailboxesRequestMailboxParameters `json:"mailboxes"`

	// DomainID: ID of the domain in which to create the mailboxes.
	DomainID string `json:"domain_id"`

	// SubscriptionPeriod: subscription renewal period, it can be monthly or yearly.
	// Default value: unknown_subscription_period
	SubscriptionPeriod MailboxSubscriptionPeriod `json:"subscription_period"`
}

// BatchCreateMailboxesResponse: batch create mailboxes response.
type BatchCreateMailboxesResponse struct {
	// Mailboxes: mailboxes created.
	Mailboxes []*Mailbox `json:"mailboxes"`
}

// CreateDomainRequest: create domain request.
type CreateDomainRequest struct {
	// ProjectID: ID of the project to which the domain belongs.
	ProjectID string `json:"project_id"`

	// Name: fully qualified domain name.
	Name string `json:"name"`
}

// DeleteDomainRequest: delete domain request.
type DeleteDomainRequest struct {
	// DomainID: ID of the domain to delete.
	DomainID string `json:"-"`
}

// DeleteMailboxRequest: delete mailbox request.
type DeleteMailboxRequest struct {
	// MailboxID: ID of the mailbox to delete.
	MailboxID string `json:"-"`
}

// GetDomainRecordsRequest: get domain records request.
type GetDomainRecordsRequest struct {
	// DomainID: (Optional) ID of the domain in which to get the records.
	DomainID string `json:"-"`
}

// GetDomainRecordsResponse: get domain records response.
type GetDomainRecordsResponse struct {
	// Autoconfig: record designed to automatically configure an email client with the appropriate mail server settings using a standardized XML file format.
	Autoconfig *DomainRecord `json:"autoconfig"`

	// Autodiscover: record designed to automate the discovery of server settings to configure email clients like Outlook.
	Autodiscover *DomainRecord `json:"autodiscover"`

	// Caldav: record that allows clients to interact with calendar data stored on a server.
	Caldav *DomainRecord `json:"caldav"`

	// Carddav: record that allows clients to interact with contact data stored on a server.
	Carddav *DomainRecord `json:"carddav"`

	// Dkim: record that allows the DKIM email authentication method to be employed to verify the authenticity of an email message.
	Dkim *DomainRecord `json:"dkim"`

	// Dmarc: record that provides a mechanism for email receivers to determine if incoming messages are legitimate and were sent from authorized sources.
	Dmarc *DomainRecord `json:"dmarc"`

	// DomainValidation: record that provides a unique token to validate a domain ownership.
	DomainValidation *DomainRecord `json:"domain_validation"`

	// Imap: record that allows accessing the mailbox with the IMAP protocol.
	Imap *DomainRecord `json:"imap"`

	// Mx: record that directs emails to a mail server.
	Mx *DomainRecord `json:"mx"`

	// Pop3: record that allows accessing the mailbox with the POP3 protocol.
	Pop3 *DomainRecord `json:"pop3"`

	// Spf: record that lists all the servers authorized to send emails from a particular domain.
	Spf *DomainRecord `json:"spf"`

	// Submission: record that allows the use of the SMTP submission mechanism.
	Submission *DomainRecord `json:"submission"`
}

// GetDomainRequest: get domain request.
type GetDomainRequest struct {
	// DomainID: ID of the domain to get.
	DomainID string `json:"-"`
}

// GetMailboxRequest: get mailbox request.
type GetMailboxRequest struct {
	// MailboxID: ID of the mailbox to get.
	MailboxID string `json:"-"`
}

// ListDomainsRequest: list domains request.
type ListDomainsRequest struct {
	// OrderBy: default value: created_at_desc
	OrderBy ListDomainsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	ProjectID *string `json:"-"`

	Statuses []DomainStatus `json:"-"`

	Search *string `json:"-"`
}

// ListDomainsResponse: list domains response.
type ListDomainsResponse struct {
	// TotalCount: number of domains that match the request (without pagination).
	TotalCount uint64 `json:"total_count"`

	// Domains: single page of domains matching the requested criteria.
	Domains []*Domain `json:"domains"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDomainsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListDomainsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Domains = append(r.Domains, results.Domains...)
	r.TotalCount += uint64(len(results.Domains))
	return uint64(len(results.Domains)), nil
}

// ListMailboxesRequest: list mailboxes request.
type ListMailboxesRequest struct {
	// OrderBy: order matching mailbox by different criteria.
	// Default value: created_at_desc
	OrderBy ListMailboxesRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`

	// PageSize: requested page size. Value must be between 1 and 1000.
	PageSize *uint32 `json:"-"`

	// DomainID: (Optional) ID of the domain in which to list the mailboxes.
	DomainID *string `json:"-"`

	// Statuses: (Optional) Filter mailboxes by their statuses.
	Statuses []MailboxStatus `json:"-"`

	// Search: (Optional) Search term to filter mailboxes on name and local_part.
	Search *string `json:"-"`
}

// ListMailboxesResponse: list mailboxes response.
type ListMailboxesResponse struct {
	// TotalCount: number of mailboxes that match the request (without pagination).
	TotalCount uint64 `json:"total_count"`

	// Mailboxes: mailboxes matching the requested criteria.
	Mailboxes []*Mailbox `json:"mailboxes"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListMailboxesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListMailboxesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListMailboxesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Mailboxes = append(r.Mailboxes, results.Mailboxes...)
	r.TotalCount += uint64(len(results.Mailboxes))
	return uint64(len(results.Mailboxes)), nil
}

// RestoreMailboxRequest: restore mailbox request.
type RestoreMailboxRequest struct {
	// MailboxID: ID of the mailbox to restore.
	MailboxID string `json:"-"`
}

// UpdateMailboxRequest: update mailbox request.
type UpdateMailboxRequest struct {
	// MailboxID: ID of the mailbox to update.
	MailboxID string `json:"-"`

	// SubscriptionPeriod: (Optional) New subscription period for the mailbox.
	// Default value: unknown_subscription_period
	SubscriptionPeriod *MailboxSubscriptionPeriod `json:"subscription_period,omitempty"`

	// NewPassword: (Optional) New password of the mailbox.
	NewPassword *string `json:"new_password,omitempty"`
}

// ValidateDomainRecordsRequest: validate domain records request.
type ValidateDomainRecordsRequest struct {
	DomainID string `json:"-"`
}

// This API allows you to manage your Mailbox services.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// CreateDomain: You must specify a `project_id` and a `domain_name` to register a domain in a specific Project.
func (s *API) CreateDomain(req *CreateDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mailbox/v1alpha1/domains",
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

// ListDomains: The return list can be filtered with request parameters.
func (s *API) ListDomains(req *ListDomainsRequest, opts ...scw.RequestOption) (*ListDomainsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "search", req.Search)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mailbox/v1alpha1/domains",
		Query:  query,
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomain: Get a domain by its ID.
func (s *API) GetDomain(req *GetDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mailbox/v1alpha1/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForDomainRequest is used by WaitForDomain method.
type WaitForDomainRequest struct {
	DomainID      string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForDomain waits for the Domain to reach a terminal state.
func (s *API) WaitForDomain(req *WaitForDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	timeout := defaultMailboxTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultMailboxRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[DomainStatus]struct{}{
		DomainStatusCreating:     {},
		DomainStatusValidating:   {},
		DomainStatusProvisioning: {},
		DomainStatusDeleting:     {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetDomain(&GetDomainRequest{
				DomainID: req.DomainID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTransient := transientStatuses[res.Status]

			return res, !isTransient, nil
		},
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
		Timeout:          timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for Domain failed")
	}

	return res.(*Domain), nil
}

// DeleteDomain: Delete a domain by its ID.
func (s *API) DeleteDomain(req *DeleteDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/mailbox/v1alpha1/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomainRecords: Get domain records by its ID.
func (s *API) GetDomainRecords(req *GetDomainRecordsRequest, opts ...scw.RequestOption) (*GetDomainRecordsResponse, error) {
	var err error

	if fmt.Sprint(req.DomainID) == "" {
		return nil, errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mailbox/v1alpha1/domains/" + fmt.Sprint(req.DomainID) + "/records",
	}

	var resp GetDomainRecordsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateDomainRecords:
func (s *API) ValidateDomainRecords(req *ValidateDomainRecordsRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.DomainID) == "" {
		return errors.New("field DomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mailbox/v1alpha1/domains/" + fmt.Sprint(req.DomainID) + "/validate-records",
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

// BatchCreateMailboxes:
func (s *API) BatchCreateMailboxes(req *BatchCreateMailboxesRequest, opts ...scw.RequestOption) (*BatchCreateMailboxesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mailbox/v1alpha1/batch-create-mailboxes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp BatchCreateMailboxesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListMailboxes: The return list can be filtered with request parameters.
func (s *API) ListMailboxes(req *ListMailboxesRequest, opts ...scw.RequestOption) (*ListMailboxesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "domain_id", req.DomainID)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "search", req.Search)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mailbox/v1alpha1/mailboxes",
		Query:  query,
	}

	var resp ListMailboxesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMailbox: Get a mailbox by its ID.
func (s *API) GetMailbox(req *GetMailboxRequest, opts ...scw.RequestOption) (*Mailbox, error) {
	var err error

	if fmt.Sprint(req.MailboxID) == "" {
		return nil, errors.New("field MailboxID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mailbox/v1alpha1/mailboxes/" + fmt.Sprint(req.MailboxID) + "",
	}

	var resp Mailbox

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForMailboxRequest is used by WaitForMailbox method.
type WaitForMailboxRequest struct {
	MailboxID     string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForMailbox waits for the Mailbox to reach a terminal state.
func (s *API) WaitForMailbox(req *WaitForMailboxRequest, opts ...scw.RequestOption) (*Mailbox, error) {
	timeout := defaultMailboxTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultMailboxRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[MailboxStatus]struct{}{
		MailboxStatusCreating:       {},
		MailboxStatusWaitingPayment: {},
		MailboxStatusWaitingDomain:  {},
		MailboxStatusRenewing:       {},
		MailboxStatusDeleting:       {},
		MailboxStatusRestoring:      {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetMailbox(&GetMailboxRequest{
				MailboxID: req.MailboxID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTransient := transientStatuses[res.Status]

			return res, !isTransient, nil
		},
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
		Timeout:          timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for Mailbox failed")
	}

	return res.(*Mailbox), nil
}

// UpdateMailbox: Update a mailbox subscription period or password with its ID.
func (s *API) UpdateMailbox(req *UpdateMailboxRequest, opts ...scw.RequestOption) (*Mailbox, error) {
	var err error

	if fmt.Sprint(req.MailboxID) == "" {
		return nil, errors.New("field MailboxID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/mailbox/v1alpha1/mailboxes/" + fmt.Sprint(req.MailboxID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Mailbox

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteMailbox: Delete a mailbox by its ID.
func (s *API) DeleteMailbox(req *DeleteMailboxRequest, opts ...scw.RequestOption) (*Mailbox, error) {
	var err error

	if fmt.Sprint(req.MailboxID) == "" {
		return nil, errors.New("field MailboxID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/mailbox/v1alpha1/mailboxes/" + fmt.Sprint(req.MailboxID) + "",
	}

	var resp Mailbox

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestoreMailbox: Restore a mailbox in deletion scheduled status by its ID.
func (s *API) RestoreMailbox(req *RestoreMailboxRequest, opts ...scw.RequestOption) (*Mailbox, error) {
	var err error

	if fmt.Sprint(req.MailboxID) == "" {
		return nil, errors.New("field MailboxID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mailbox/v1alpha1/mailboxes/" + fmt.Sprint(req.MailboxID) + "/restore",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Mailbox

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package tem_admin provides methods and message types of the tem_admin v1alpha1 API.
package tem_admin

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

type DomainBlockReason string

const (
	// If unspecified, block reason is unknown by default.
	DomainBlockReasonUnknownBlockReason = DomainBlockReason("unknown_block_reason")
	// The domain is blocked because it is blacklisted.
	DomainBlockReasonBlacklist = DomainBlockReason("blacklist")
	// The domain is blocked due to spamming.
	DomainBlockReasonSpam = DomainBlockReason("spam")
	// The domain is blocked due to a low score.
	DomainBlockReasonLowDomainScore = DomainBlockReason("low_domain_score")
	// The domain was blocked manually.
	DomainBlockReasonManual = DomainBlockReason("manual")
)

func (enum DomainBlockReason) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_block_reason"
	}
	return string(enum)
}

func (enum DomainBlockReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainBlockReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainBlockReason(DomainBlockReason(tmp).String())
	return nil
}

type DomainBypass string

const (
	// The domain does not benefit from bypass.
	DomainBypassNo = DomainBypass("no")
	// The domain benefits from bypass.
	DomainBypassYes = DomainBypass("yes")
	// The domain is in purgatory zone. This means your domain is blocked but it can be unblocked after a while.
	DomainBypassPurgatory = DomainBypass("purgatory")
)

func (enum DomainBypass) String() string {
	if enum == "" {
		// return default value if empty
		return "no"
	}
	return string(enum)
}

func (enum DomainBypass) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainBypass) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainBypass(DomainBypass(tmp).String())
	return nil
}

type DomainLastStatusRecordStatus string

const (
	// If unspecified, the status of the domain's record is unknown by default.
	DomainLastStatusRecordStatusUnknownRecordStatus = DomainLastStatusRecordStatus("unknown_record_status")
	// The record is valid.
	DomainLastStatusRecordStatusValid = DomainLastStatusRecordStatus("valid")
	// The record is invalid.
	DomainLastStatusRecordStatusInvalid = DomainLastStatusRecordStatus("invalid")
	// The record was not found.
	DomainLastStatusRecordStatusNotFound = DomainLastStatusRecordStatus("not_found")
)

func (enum DomainLastStatusRecordStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_record_status"
	}
	return string(enum)
}

func (enum DomainLastStatusRecordStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainLastStatusRecordStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainLastStatusRecordStatus(DomainLastStatusRecordStatus(tmp).String())
	return nil
}

type DomainReputationStatus string

const (
	// If unspecified, the domain's reputation is unknown by default.
	DomainReputationStatusUnknownStatus = DomainReputationStatus("unknown_status")
	// The domain has an excellent reputation.
	DomainReputationStatusExcellent = DomainReputationStatus("excellent")
	// The domain has a good reputation.
	DomainReputationStatusGood = DomainReputationStatus("good")
	// The domain has an average reputation.
	DomainReputationStatusAverage = DomainReputationStatus("average")
	// The domain has a bad reputation.
	DomainReputationStatusBad = DomainReputationStatus("bad")
)

func (enum DomainReputationStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum DomainReputationStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainReputationStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainReputationStatus(DomainReputationStatus(tmp).String())
	return nil
}

type DomainStatus string

const (
	// If unspecified, the status of the domain is unknown by default.
	DomainStatusUnknown = DomainStatus("unknown")
	// The domain is checked.
	DomainStatusChecked = DomainStatus("checked")
	// The domain is unchecked.
	DomainStatusUnchecked = DomainStatus("unchecked")
	// The domain is invalid.
	DomainStatusInvalid = DomainStatus("invalid")
	// The domain is locked.
	DomainStatusLocked = DomainStatus("locked")
	// The domain is revoked.
	DomainStatusRevoked = DomainStatus("revoked")
	// The domain is pending, waiting to be checked.
	DomainStatusPending = DomainStatus("pending")
	// The domain is blocked.
	DomainStatusBlocked = DomainStatus("blocked")
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

type EmailFlag string

const (
	// If unspecified, the flag type is unknown by default.
	EmailFlagUnknownFlag = EmailFlag("unknown_flag")
	// Refers to a non critical error received while sending the email(s). Soft bounced emails are retried.
	EmailFlagSoftBounce = EmailFlag("soft_bounce")
	// Refers to a critical error that happened while sending the email(s).
	EmailFlagHardBounce = EmailFlag("hard_bounce")
	// Refers to an email considered as spam.
	EmailFlagSpam = EmailFlag("spam")
	// Refers to an undelivered email because the recipient mailbox is full.
	EmailFlagMailboxFull = EmailFlag("mailbox_full")
	// Refers to an undelivered email because the recipient mailbox does not exist.
	EmailFlagMailboxNotFound = EmailFlag("mailbox_not_found")
	// Refers to an email slightly delayed by the recipient to ensure that Scaleway is not sending spam.
	EmailFlagGreylisted = EmailFlag("greylisted")
	// Refers to an email with a `send-before` tag to indicate the maximum time limit for the email to be sent.
	EmailFlagSendBeforeExpiration = EmailFlag("send_before_expiration")
	// (Admin only) Refers to an email with a return message indicating a bad reputation or that our sending IP is blacklisted.
	EmailFlagIPBlacklisted = EmailFlag("ip_blacklisted")
)

func (enum EmailFlag) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_flag"
	}
	return string(enum)
}

func (enum EmailFlag) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EmailFlag) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EmailFlag(EmailFlag(tmp).String())
	return nil
}

type EmailRcptType string

const (
	// If unspecified, the recipient type is unknown by default.
	EmailRcptTypeUnknownRcptType = EmailRcptType("unknown_rcpt_type")
	// Primary recipient.
	EmailRcptTypeTo = EmailRcptType("to")
	// Carbon copy recipient.
	EmailRcptTypeCc = EmailRcptType("cc")
	// Blind carbon copy recipient.
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
	// If unspecified, the status of the email is unknown by default.
	EmailStatusUnknown = EmailStatus("unknown")
	// The email is new.
	EmailStatusNew = EmailStatus("new")
	// The email is in the process of being sent.
	EmailStatusSending = EmailStatus("sending")
	// The email was sent.
	EmailStatusSent = EmailStatus("sent")
	// The sending of the email failed.
	EmailStatusFailed = EmailStatus("failed")
	// The sending of the email was canceled.
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

type ListEmailsRequestOrderBy string

const (
	// Order by creation date (descending chronological order).
	ListEmailsRequestOrderByCreatedAtDesc = ListEmailsRequestOrderBy("created_at_desc")
	// Order by creation date (ascending chronological order).
	ListEmailsRequestOrderByCreatedAtAsc = ListEmailsRequestOrderBy("created_at_asc")
	// Order by last update date (descending chronological order).
	ListEmailsRequestOrderByUpdatedAtDesc = ListEmailsRequestOrderBy("updated_at_desc")
	// Order by last update date (ascending chronological order).
	ListEmailsRequestOrderByUpdatedAtAsc = ListEmailsRequestOrderBy("updated_at_asc")
	// Order by status (descending alphabetical order).
	ListEmailsRequestOrderByStatusDesc = ListEmailsRequestOrderBy("status_desc")
	// Order by status (ascending alphabetical order).
	ListEmailsRequestOrderByStatusAsc = ListEmailsRequestOrderBy("status_asc")
	// Order by mail_from (descending alphabetical order).
	ListEmailsRequestOrderByMailFromDesc = ListEmailsRequestOrderBy("mail_from_desc")
	// Order by mail_from (ascending alphabetical order).
	ListEmailsRequestOrderByMailFromAsc = ListEmailsRequestOrderBy("mail_from_asc")
	// Order by mail recipient (descending alphabetical order).
	ListEmailsRequestOrderByMailRcptDesc = ListEmailsRequestOrderBy("mail_rcpt_desc")
	// Order by mail recipient (ascending alphabetical order).
	ListEmailsRequestOrderByMailRcptAsc = ListEmailsRequestOrderBy("mail_rcpt_asc")
	// Order by subject (descending alphabetical order).
	ListEmailsRequestOrderBySubjectDesc = ListEmailsRequestOrderBy("subject_desc")
	// Order by subject (ascending alphabetical order).
	ListEmailsRequestOrderBySubjectAsc = ListEmailsRequestOrderBy("subject_asc")
)

func (enum ListEmailsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_desc"
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

// DomainReputation: domain reputation.
type DomainReputation struct {
	// Status: status of your domain's reputation.
	// Default value: unknown_status
	Status DomainReputationStatus `json:"status"`

	// Score: a range from 0 to 100 that determines your domain's reputation score. A score of `0` means a bad domain reputation and a score of `100` means an excellent domain reputation.
	Score uint32 `json:"score"`

	// ScoredAt: time and date the score was calculated.
	ScoredAt *time.Time `json:"scored_at"`

	// PreviousScore: the previously-calculated domain's reputation score.
	PreviousScore *uint32 `json:"previous_score"`

	// PreviousScoredAt: time and date the previous reputation score was calculated.
	PreviousScoredAt *time.Time `json:"previous_scored_at"`
}

// DomainStatistics: domain statistics.
type DomainStatistics struct {
	TotalCount uint32 `json:"total_count"`

	SentCount uint32 `json:"sent_count"`

	FailedCount uint32 `json:"failed_count"`

	CanceledCount uint32 `json:"canceled_count"`
}

// EmailTry: email try.
type EmailTry struct {
	// Rank: rank number of this attempt to send the email.
	Rank uint32 `json:"rank"`

	// TriedAt: date of the attempt.
	TriedAt *time.Time `json:"tried_at"`

	// Code: the SMTP status code received after the attempt. 0 if the attempt did not reach an SMTP server.
	Code int32 `json:"code"`

	// Message: the SMTP message received, if any. If the attempt did not reach an SMTP server, the message says why.
	Message string `json:"message"`
}

// DomainLastStatusDkimRecord: domain last status dkim record.
type DomainLastStatusDkimRecord struct {
	// Status: status of the DKIM record's configurartion.
	// Default value: unknown_record_status
	Status DomainLastStatusRecordStatus `json:"status"`

	// LastValidAt: time and date the DKIM record was last valid.
	LastValidAt *time.Time `json:"last_valid_at"`

	// Error: an error text displays in case the record is not valid.
	Error *string `json:"error"`
}

// DomainLastStatusDmarcRecord: domain last status dmarc record.
type DomainLastStatusDmarcRecord struct {
	// Status: status of the DMARC record's configuration.
	// Default value: unknown_record_status
	Status DomainLastStatusRecordStatus `json:"status"`

	// LastValidAt: time and date the DMARC record was last valid.
	LastValidAt *time.Time `json:"last_valid_at"`

	// Error: an error text displays in case the record is not valid.
	Error *string `json:"error"`
}

// DomainLastStatusSpfRecord: domain last status spf record.
type DomainLastStatusSpfRecord struct {
	// Status: status of the SPF record's configurartion.
	// Default value: unknown_record_status
	Status DomainLastStatusRecordStatus `json:"status"`

	// LastValidAt: time and date the SPF record was last valid.
	LastValidAt *time.Time `json:"last_valid_at"`

	// Error: an error text displays in case the record is not valid.
	Error *string `json:"error"`
}

// Domain: domain.
type Domain struct {
	// ID: ID of the domain.
	ID string `json:"id"`

	// OrganizationID: ID of the organization to which the domain belongs.
	OrganizationID string `json:"organization_id"`

	// ProjectID: ID of the project.
	ProjectID string `json:"project_id"`

	// Name: domain name (example.com).
	Name string `json:"name"`

	// Status: status of the domain.
	// Default value: unknown
	Status DomainStatus `json:"status"`

	// CreatedAt: date and time of domain's creation.
	CreatedAt *time.Time `json:"created_at"`

	// NextCheckAt: date and time of the next scheduled check.
	NextCheckAt *time.Time `json:"next_check_at"`

	// LastValidAt: date and time the domain was last found to be valid.
	LastValidAt *time.Time `json:"last_valid_at"`

	// RevokedAt: date and time of the revocation of the domain.
	RevokedAt *time.Time `json:"revoked_at"`

	// Deprecated: LastError: error message if the last check failed.
	LastError *string `json:"last_error,omitempty"`

	// SpfConfig: snippet of the SPF record that should be registered in the DNS zone.
	SpfConfig string `json:"spf_config"`

	// DkimConfig: dKIM public key, as should be recorded in the DNS zone.
	DkimConfig string `json:"dkim_config"`

	// Statistics: domain's statistics.
	Statistics *DomainStatistics `json:"statistics"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`

	// BlockReason: the reason why the domain is blocked.
	// Default value: unknown_block_reason
	BlockReason DomainBlockReason `json:"block_reason"`

	// Reputation: the domain's reputation is available when your domain is checked and has sent enough emails.
	Reputation *DomainReputation `json:"reputation"`

	// Bypass: default value: no
	Bypass DomainBypass `json:"bypass"`
}

// Email: email.
type Email struct {
	// ID: technical ID of the email.
	ID string `json:"id"`

	// MessageID: messageID of the email.
	MessageID string `json:"message_id"`

	// ProjectID: ID of the project to which the email belongs.
	ProjectID string `json:"project_id"`

	// MailFrom: email address of the sender.
	MailFrom string `json:"mail_from"`

	// Deprecated: RcptTo: email address of the recipient.
	RcptTo *string `json:"rcpt_to,omitempty"`

	// MailRcpt: email address of the recipient.
	MailRcpt string `json:"mail_rcpt"`

	// RcptType: type of the recipient.
	// Default value: unknown_rcpt_type
	RcptType EmailRcptType `json:"rcpt_type"`

	// Subject: subject of the email.
	Subject string `json:"subject"`

	// CreatedAt: creation date of the email object.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update time of the email object.
	UpdatedAt *time.Time `json:"updated_at"`

	// Status: status of the email.
	// Default value: unknown
	Status EmailStatus `json:"status"`

	// StatusDetails: additional information on the status.
	StatusDetails *string `json:"status_details"`

	// TryCount: total number of attempts to send the email.
	TryCount uint32 `json:"try_count"`

	// LastTries: information about the latest three attempts to send the email.
	LastTries []*EmailTry `json:"last_tries"`

	// Flags: flags categorize emails. They allow you to obtain more information about recurring errors, for example.
	Flags []EmailFlag `json:"flags"`
}

// BlockDomainRequest: block domain request.
type BlockDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainID: ID of the domain to block.
	DomainID string `json:"-"`

	// Reason: the reason why the domain is blocked.
	// Default value: unknown_block_reason
	Reason DomainBlockReason `json:"reason"`
}

// CancelEmailRequest: cancel email request.
type CancelEmailRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// EmailID: ID of the email to cancel.
	EmailID string `json:"-"`
}

// CheckDomainRequest: check domain request.
type CheckDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainID: ID of the domain to check.
	DomainID string `json:"-"`
}

// DomainLastStatus: domain last status.
type DomainLastStatus struct {
	// DomainID: the id of the domain.
	DomainID string `json:"domain_id"`

	// DomainName: the domain name (example.com).
	DomainName string `json:"domain_name"`

	// SpfRecord: the SPF record verification data.
	SpfRecord *DomainLastStatusSpfRecord `json:"spf_record"`

	// DkimRecord: the DKIM record verification data.
	DkimRecord *DomainLastStatusDkimRecord `json:"dkim_record"`

	// DmarcRecord: the DMARC record verification data.
	DmarcRecord *DomainLastStatusDmarcRecord `json:"dmarc_record"`
}

// GetDomainLastStatusRequest: get domain last status request.
type GetDomainLastStatusRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainID: ID of the domain to delete.
	DomainID string `json:"-"`
}

// GetDomainRequest: get domain request.
type GetDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainID: ID of the domain.
	DomainID string `json:"-"`
}

// GetEmailRequest: get email request.
type GetEmailRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	EmailID string `json:"-"`
}

// GetStatisticsRequest: get statistics request.
type GetStatisticsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: (Optional) Number of emails sent for this project.
	ProjectID *string `json:"-"`

	// DomainID: (Optional) Number of emails sent from this domain (must be coherent with the `project_id` and the `organization_id`).
	DomainID *string `json:"-"`

	// Since: (Optional) Number of emails created after this date.
	Since *time.Time `json:"-"`

	// Until: (Optional) Number of emails created before this date.
	Until *time.Time `json:"-"`

	// MailFrom: (Optional) Number of emails sent with this sender's email address.
	MailFrom *string `json:"-"`
}

// ListDomainsRequest: list domains request.
type ListDomainsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number (1 for the first page).
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	ProjectID *string `json:"-"`

	Status []DomainStatus `json:"-"`

	OrganizationID *string `json:"-"`

	Name *string `json:"-"`
}

// ListDomainsResponse: list domains response.
type ListDomainsResponse struct {
	// TotalCount: total number of domains matching the request (without pagination).
	TotalCount uint32 `json:"total_count"`

	Domains []*Domain `json:"domains"`
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

// ListEmailsRequest: list emails request.
type ListEmailsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// ProjectID: (Optional) ID of the project in which to list the emails.
	ProjectID *string `json:"-"`

	// DomainID: (Optional) ID of the domain for which to list the emails.
	DomainID *string `json:"-"`

	// MessageID: (Optional) ID of the message for which to list the emails.
	MessageID *string `json:"-"`

	// Since: (Optional) List emails created after this date.
	Since *time.Time `json:"-"`

	// Until: (Optional) List emails created before this date.
	Until *time.Time `json:"-"`

	// MailFrom: (Optional) List emails sent with this sender's email address.
	MailFrom *string `json:"-"`

	// Deprecated: MailTo: list emails sent to a specific email address.
	MailTo *string `json:"-"`

	// MailRcpt: (Optional) List emails sent to this recipient's email address.
	MailRcpt *string `json:"-"`

	// Statuses: (Optional) List emails with any of these statuses.
	Statuses []EmailStatus `json:"-"`

	// Subject: (Optional) List emails with this subject.
	Subject *string `json:"-"`

	// Search: (Optional) List emails by searching to all fields.
	Search *string `json:"-"`

	// OrderBy: (Optional) List emails corresponding to specific criteria.
	// Default value: created_at_desc
	OrderBy ListEmailsRequestOrderBy `json:"-"`

	// Flags: (Optional) List emails containing only specific flags.
	Flags []EmailFlag `json:"-"`
}

// ListEmailsResponse: list emails response.
type ListEmailsResponse struct {
	// TotalCount: number of emails matching the requested criteria.
	TotalCount uint32 `json:"total_count"`

	// Emails: single page of emails matching the requested criteria.
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

// RevokeDomainRequest: revoke domain request.
type RevokeDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainID: ID of the domain to revoke.
	DomainID string `json:"-"`
}

// Statistics: statistics.
type Statistics struct {
	// TotalCount: total number of emails matching the request criteria.
	TotalCount uint32 `json:"total_count"`

	// NewCount: number of emails still in the `new` transient state (received from the API, not yet processed).
	NewCount uint32 `json:"new_count"`

	// SendingCount: number of emails still in the `sending` transient state (received from the API, not yet in their final status).
	SendingCount uint32 `json:"sending_count"`

	// SentCount: number of emails in the final `sent` state (have been delivered to the target mail system).
	SentCount uint32 `json:"sent_count"`

	// FailedCount: number of emails in the final `failed` state (refused by the target mail system with a final error status).
	FailedCount uint32 `json:"failed_count"`

	// CanceledCount: number of emails in the final `canceled` state (canceled by customer's request).
	CanceledCount uint32 `json:"canceled_count"`
}

// UnblockDomainRequest: unblock domain request.
type UnblockDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainID: ID of the domain to unblock.
	DomainID string `json:"-"`
}

// Tem-admin.
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

// GetEmail: Get information about an email.
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
		Method: "GET",
		Path:   "/transactional-email-admin/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails/" + fmt.Sprint(req.EmailID) + "",
	}

	var resp Email

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListEmails: Retrieve the list of emails sent from a specific domain or for a specific Project or Organization. You must specify the `region`.
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
	parameter.AddToQuery(query, "mail_rcpt", req.MailRcpt)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "subject", req.Subject)
	parameter.AddToQuery(query, "search", req.Search)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "flags", req.Flags)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/transactional-email-admin/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails",
		Query:  query,
	}

	var resp ListEmailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetStatistics: Get statistics on the email statuses.
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
		Method: "GET",
		Path:   "/transactional-email-admin/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/statistics",
		Query:  query,
	}

	var resp Statistics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelEmail: Try to cancel an email if it has not yet been sent.
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
		Method: "POST",
		Path:   "/transactional-email-admin/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails/" + fmt.Sprint(req.EmailID) + "/cancel",
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

// GetDomain: Get information about a domain.
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
		Method: "GET",
		Path:   "/transactional-email-admin/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDomains: List domains in a project and/or in an organization.
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
		Method: "GET",
		Path:   "/transactional-email-admin/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Query:  query,
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevokeDomain: Revoke a domain.
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
		Method: "POST",
		Path:   "/transactional-email-admin/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/revoke",
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

// CheckDomain: Ask for an immediate check of a domain (DNS check).
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
		Method: "POST",
		Path:   "/transactional-email-admin/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/check",
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

// GetDomainLastStatus: Display SPF and DKIM records status and potential errors, including the found records to make debugging easier.
func (s *API) GetDomainLastStatus(req *GetDomainLastStatusRequest, opts ...scw.RequestOption) (*DomainLastStatus, error) {
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
		Method: "GET",
		Path:   "/transactional-email-admin/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/verification",
	}

	var resp DomainLastStatus

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BlockDomain: Block the domain and delete every pending emails.
func (s *API) BlockDomain(req *BlockDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
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
		Method: "POST",
		Path:   "/transactional-email-admin/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/block",
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

// UnblockDomain: Unblock the domain as it can send emails again.
func (s *API) UnblockDomain(req *UnblockDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
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
		Method: "POST",
		Path:   "/transactional-email-admin/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/unblock",
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

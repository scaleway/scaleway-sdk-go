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

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
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

type DomainLastStatusAutoconfigStateReason string

const (
	// If not specified, the auto-configuration state is unknown by default.
	DomainLastStatusAutoconfigStateReasonUnknownReason = DomainLastStatusAutoconfigStateReason("unknown_reason")
	// The token doesn't have the necessary permissions to manage the domain's DNS records.
	DomainLastStatusAutoconfigStateReasonPermissionDenied = DomainLastStatusAutoconfigStateReason("permission_denied")
	// The domain does not exist or isn't manageable by the token.
	DomainLastStatusAutoconfigStateReasonDomainNotFound = DomainLastStatusAutoconfigStateReason("domain_not_found")
)

func (enum DomainLastStatusAutoconfigStateReason) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_reason"
	}
	return string(enum)
}

func (enum DomainLastStatusAutoconfigStateReason) Values() []DomainLastStatusAutoconfigStateReason {
	return []DomainLastStatusAutoconfigStateReason{
		"unknown_reason",
		"permission_denied",
		"domain_not_found",
	}
}

func (enum DomainLastStatusAutoconfigStateReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainLastStatusAutoconfigStateReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainLastStatusAutoconfigStateReason(DomainLastStatusAutoconfigStateReason(tmp).String())
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

func (enum DomainLastStatusRecordStatus) Values() []DomainLastStatusRecordStatus {
	return []DomainLastStatusRecordStatus{
		"unknown_record_status",
		"valid",
		"invalid",
		"not_found",
	}
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
	// If unspecified, the status of the domain's reputation is unknown by default.
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

func (enum DomainReputationStatus) Values() []DomainReputationStatus {
	return []DomainReputationStatus{
		"unknown_status",
		"excellent",
		"good",
		"average",
		"bad",
	}
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
	// The domain is in process of auto-configuration of the domain's DNS zone.
	DomainStatusAutoconfiguring = DomainStatus("autoconfiguring")
)

func (enum DomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DomainStatus) Values() []DomainStatus {
	return []DomainStatus{
		"unknown",
		"checked",
		"unchecked",
		"invalid",
		"locked",
		"revoked",
		"pending",
		"autoconfiguring",
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
)

func (enum EmailFlag) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_flag"
	}
	return string(enum)
}

func (enum EmailFlag) Values() []EmailFlag {
	return []EmailFlag{
		"unknown_flag",
		"soft_bounce",
		"hard_bounce",
		"spam",
		"mailbox_full",
		"mailbox_not_found",
		"greylisted",
		"send_before_expiration",
	}
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

func (enum EmailRcptType) Values() []EmailRcptType {
	return []EmailRcptType{
		"unknown_rcpt_type",
		"to",
		"cc",
		"bcc",
	}
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

func (enum EmailStatus) Values() []EmailStatus {
	return []EmailStatus{
		"unknown",
		"new",
		"sending",
		"sent",
		"failed",
		"canceled",
	}
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

func (enum ListEmailsRequestOrderBy) Values() []ListEmailsRequestOrderBy {
	return []ListEmailsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"updated_at_desc",
		"updated_at_asc",
		"status_desc",
		"status_asc",
		"mail_from_desc",
		"mail_from_asc",
		"mail_rcpt_desc",
		"mail_rcpt_asc",
		"subject_desc",
		"subject_asc",
	}
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

type ListWebhookEventsRequestOrderBy string

const (
	ListWebhookEventsRequestOrderByCreatedAtDesc = ListWebhookEventsRequestOrderBy("created_at_desc")
	ListWebhookEventsRequestOrderByCreatedAtAsc  = ListWebhookEventsRequestOrderBy("created_at_asc")
)

func (enum ListWebhookEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_desc"
	}
	return string(enum)
}

func (enum ListWebhookEventsRequestOrderBy) Values() []ListWebhookEventsRequestOrderBy {
	return []ListWebhookEventsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
	}
}

func (enum ListWebhookEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListWebhookEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListWebhookEventsRequestOrderBy(ListWebhookEventsRequestOrderBy(tmp).String())
	return nil
}

type ListWebhooksRequestOrderBy string

const (
	ListWebhooksRequestOrderByCreatedAtDesc = ListWebhooksRequestOrderBy("created_at_desc")
	ListWebhooksRequestOrderByCreatedAtAsc  = ListWebhooksRequestOrderBy("created_at_asc")
)

func (enum ListWebhooksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_desc"
	}
	return string(enum)
}

func (enum ListWebhooksRequestOrderBy) Values() []ListWebhooksRequestOrderBy {
	return []ListWebhooksRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
	}
}

func (enum ListWebhooksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListWebhooksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListWebhooksRequestOrderBy(ListWebhooksRequestOrderBy(tmp).String())
	return nil
}

type ProjectSettingsPeriodicReportFrequency string

const (
	// If unspecified, the frequency is unknown by default.
	ProjectSettingsPeriodicReportFrequencyUnknownFrequency = ProjectSettingsPeriodicReportFrequency("unknown_frequency")
	// The periodic report is sent once a month.
	ProjectSettingsPeriodicReportFrequencyMonthly = ProjectSettingsPeriodicReportFrequency("monthly")
	// The periodic report is sent once a week.
	ProjectSettingsPeriodicReportFrequencyWeekly = ProjectSettingsPeriodicReportFrequency("weekly")
	// The periodic report is sent once a day.
	ProjectSettingsPeriodicReportFrequencyDaily = ProjectSettingsPeriodicReportFrequency("daily")
)

func (enum ProjectSettingsPeriodicReportFrequency) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_frequency"
	}
	return string(enum)
}

func (enum ProjectSettingsPeriodicReportFrequency) Values() []ProjectSettingsPeriodicReportFrequency {
	return []ProjectSettingsPeriodicReportFrequency{
		"unknown_frequency",
		"monthly",
		"weekly",
		"daily",
	}
}

func (enum ProjectSettingsPeriodicReportFrequency) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ProjectSettingsPeriodicReportFrequency) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ProjectSettingsPeriodicReportFrequency(ProjectSettingsPeriodicReportFrequency(tmp).String())
	return nil
}

type WebhookEventStatus string

const (
	// If unspecified, the status of the Webhook event is unknown by default.
	WebhookEventStatusUnknownStatus = WebhookEventStatus("unknown_status")
	// The Webhook event is being sent.
	WebhookEventStatusSending = WebhookEventStatus("sending")
	// The Webhook event was sent.
	WebhookEventStatusSent = WebhookEventStatus("sent")
	// The Webhook event cannot be sent after multiple retries.
	WebhookEventStatusFailed = WebhookEventStatus("failed")
)

func (enum WebhookEventStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum WebhookEventStatus) Values() []WebhookEventStatus {
	return []WebhookEventStatus{
		"unknown_status",
		"sending",
		"sent",
		"failed",
	}
}

func (enum WebhookEventStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *WebhookEventStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = WebhookEventStatus(WebhookEventStatus(tmp).String())
	return nil
}

type WebhookEventType string

const (
	// If unspecified, the type of the Webhook Event is unknown by default.
	WebhookEventTypeUnknownType = WebhookEventType("unknown_type")
	// The email was received and is in preparation to be sent to the recipient servers.
	WebhookEventTypeEmailQueued = WebhookEventType("email_queued")
	// The email was sent but hard-bounced by the recipient server.
	WebhookEventTypeEmailDropped = WebhookEventType("email_dropped")
	// The email was sent but soft-bounced by the recipient server. In this case, the sending of the email will be automatically retried.
	WebhookEventTypeEmailDeferred = WebhookEventType("email_deferred")
	// The email was successfully sent.
	WebhookEventTypeEmailDelivered = WebhookEventType("email_delivered")
	// The email resource was identified as spam by Scaleway or by the recipient server.
	WebhookEventTypeEmailSpam = WebhookEventType("email_spam")
	// The email hard-bounced with a "mailbox not found" error.
	WebhookEventTypeEmailMailboxNotFound = WebhookEventType("email_mailbox_not_found")
	// The email was blocked before it was sent, as the recipient matches a blocklist.
	WebhookEventTypeEmailBlocklisted = WebhookEventType("email_blocklisted")
	// A new blocklist is created.
	WebhookEventTypeBlocklistCreated = WebhookEventType("blocklist_created")
)

func (enum WebhookEventType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum WebhookEventType) Values() []WebhookEventType {
	return []WebhookEventType{
		"unknown_type",
		"email_queued",
		"email_dropped",
		"email_deferred",
		"email_delivered",
		"email_spam",
		"email_mailbox_not_found",
		"email_blocklisted",
		"blocklist_created",
	}
}

func (enum WebhookEventType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *WebhookEventType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = WebhookEventType(WebhookEventType(tmp).String())
	return nil
}

// DomainRecordsDMARC: domain records dmarc.
type DomainRecordsDMARC struct {
	// Name: name of the DMARC TXT record.
	Name string `json:"name"`

	// Value: value of the DMARC TXT record.
	Value string `json:"value"`
}

// EmailTry: email try.
type EmailTry struct {
	// Rank: rank number of this attempt to send the email.
	Rank uint32 `json:"rank"`

	// TriedAt: date of the attempt to send the email.
	TriedAt *time.Time `json:"tried_at"`

	// Code: the SMTP status code received after the attempt. 0 if the attempt did not reach an SMTP server.
	Code int32 `json:"code"`

	// Message: the SMTP message received. If the attempt did not reach an SMTP server, the message returned explains what happened.
	Message string `json:"message"`
}

// DomainRecords: domain records.
type DomainRecords struct {
	// Dmarc: dMARC TXT record specification.
	Dmarc *DomainRecordsDMARC `json:"dmarc"`
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

// CreateEmailRequestAddress: create email request address.
type CreateEmailRequestAddress struct {
	// Email: email address.
	Email string `json:"email"`

	// Name: (Optional) Name displayed.
	Name *string `json:"name"`
}

// CreateEmailRequestAttachment: create email request attachment.
type CreateEmailRequestAttachment struct {
	// Name: filename of the attachment.
	Name string `json:"name"`

	// Type: mIME type of the attachment.
	Type string `json:"type"`

	// Content: content of the attachment encoded in base64.
	Content []byte `json:"content"`
}

// CreateEmailRequestHeader: create email request header.
type CreateEmailRequestHeader struct {
	// Key: email header key.
	Key string `json:"key"`

	// Value: email header value.
	Value string `json:"value"`
}

// Email: email.
type Email struct {
	// ID: technical ID of the email.
	ID string `json:"id"`

	// MessageID: message ID of the email.
	MessageID string `json:"message_id"`

	// ProjectID: ID of the Project to which the email belongs.
	ProjectID string `json:"project_id"`

	// MailFrom: email address of the sender.
	MailFrom string `json:"mail_from"`

	// Deprecated: RcptTo: email address of the recipient.
	RcptTo *string `json:"rcpt_to,omitempty"`

	// MailRcpt: email address of the recipient.
	MailRcpt string `json:"mail_rcpt"`

	// RcptType: type of recipient.
	// Default value: unknown_rcpt_type
	RcptType EmailRcptType `json:"rcpt_type"`

	// Subject: subject of the email.
	Subject string `json:"subject"`

	// CreatedAt: creation date of the email object.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update of the email object.
	UpdatedAt *time.Time `json:"updated_at"`

	// Status: status of the email.
	// Default value: unknown
	Status EmailStatus `json:"status"`

	// StatusDetails: additional status information.
	StatusDetails *string `json:"status_details"`

	// TryCount: number of attempts to send the email.
	TryCount uint32 `json:"try_count"`

	// LastTries: information about the last three attempts to send the email.
	LastTries []*EmailTry `json:"last_tries"`

	// Flags: flags categorize emails. They allow you to obtain more information about recurring errors, for example.
	Flags []EmailFlag `json:"flags"`
}

// DomainLastStatusAutoconfigState: domain last status autoconfig state.
type DomainLastStatusAutoconfigState struct {
	// Enabled: enable or disable the auto-configuration of domain DNS records.
	Enabled bool `json:"enabled"`

	// Autoconfigurable: whether the domain can be auto-configured or not.
	Autoconfigurable bool `json:"autoconfigurable"`

	// Reason: the reason that the domain cannot be auto-configurable.
	// Default value: unknown_reason
	Reason *DomainLastStatusAutoconfigStateReason `json:"reason"`
}

// DomainLastStatusDkimRecord: domain last status dkim record.
type DomainLastStatusDkimRecord struct {
	// Status: status of the DKIM record's configuration.
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
	// Status: status of the SPF record's configuration.
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

	// OrganizationID: ID of the domain's Organization.
	OrganizationID string `json:"organization_id"`

	// ProjectID: ID of the domain's Project.
	ProjectID string `json:"project_id"`

	// Name: domain name (example.com).
	Name string `json:"name"`

	// Status: status of the domain.
	// Default value: unknown
	Status DomainStatus `json:"status"`

	// CreatedAt: date and time of domain creation.
	CreatedAt *time.Time `json:"created_at"`

	// NextCheckAt: date and time of the next scheduled check.
	NextCheckAt *time.Time `json:"next_check_at"`

	// LastValidAt: date and time the domain was last valid.
	LastValidAt *time.Time `json:"last_valid_at"`

	// RevokedAt: date and time of the domain's deletion.
	RevokedAt *time.Time `json:"revoked_at"`

	// Deprecated: LastError: error message returned if the last check failed.
	LastError *string `json:"last_error,omitempty"`

	// SpfConfig: snippet of the SPF record to register in the DNS zone.
	SpfConfig string `json:"spf_config"`

	// DkimConfig: dKIM public key to record in the DNS zone.
	DkimConfig string `json:"dkim_config"`

	// Statistics: domain's statistics.
	Statistics *DomainStatistics `json:"statistics"`

	// Reputation: the domain's reputation is available when your domain is checked and has sent enough emails.
	Reputation *DomainReputation `json:"reputation"`

	// Records: list of records to configure to validate a domain.
	Records *DomainRecords `json:"records"`

	// Autoconfig: status of auto-configuration for the domain's DNS zone.
	Autoconfig bool `json:"autoconfig"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// WebhookEvent: webhook event.
type WebhookEvent struct {
	// ID: ID of the Webhook Event.
	ID string `json:"id"`

	// WebhookID: ID of the Webhook that triggers the Event.
	WebhookID string `json:"webhook_id"`

	// OrganizationID: ID of the Webhook Event Organization.
	OrganizationID string `json:"organization_id"`

	// ProjectID: ID of the Webhook Event Project.
	ProjectID string `json:"project_id"`

	// DomainID: ID of the webhook event domain.
	DomainID string `json:"domain_id"`

	// Type: type of the Webhook Event.
	// Default value: unknown_type
	Type WebhookEventType `json:"type"`

	// Status: status of the Webhook Event.
	// Default value: unknown_status
	Status WebhookEventStatus `json:"status"`

	// Data: data sent to the Webhook destination.
	Data string `json:"data"`

	// CreatedAt: date and time of the Webhook Event creation.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date and time of last Webhook Event updates.
	UpdatedAt *time.Time `json:"updated_at"`

	// EmailID: optional Email ID if the event is triggered by an Email resource.
	EmailID *string `json:"email_id"`
}

// Webhook: webhook.
type Webhook struct {
	// ID: ID of the Webhook.
	ID string `json:"id"`

	// DomainID: ID of the Domain to watch for triggering events.
	DomainID string `json:"domain_id"`

	// OrganizationID: ID of the Webhook Organization.
	OrganizationID string `json:"organization_id"`

	// ProjectID: ID of the Webhook Project.
	ProjectID string `json:"project_id"`

	// Name: name of the Webhook.
	Name string `json:"name"`

	// EventTypes: list of event types that will trigger a Webhook Event.
	EventTypes []WebhookEventType `json:"event_types"`

	// SnsArn: scaleway SNS ARN topic to push the events to.
	SnsArn string `json:"sns_arn"`

	// CreatedAt: date and time of the Webhook creation.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date and time of last Webhook updates.
	UpdatedAt *time.Time `json:"updated_at"`
}

// ProjectSettingsPeriodicReport: project settings periodic report.
type ProjectSettingsPeriodicReport struct {
	// Enabled: enable or disable periodic report notifications.
	Enabled bool `json:"enabled"`

	// Frequency: at which frequency you receive periodic report notifications.
	// Default value: unknown_frequency
	Frequency ProjectSettingsPeriodicReportFrequency `json:"frequency"`

	// SendingHour: at which hour you receive periodic report notifications.
	SendingHour uint32 `json:"sending_hour"`

	// SendingDay: on which day you receive periodic report notifications (1-7 weekly, 1-28 monthly).
	SendingDay uint32 `json:"sending_day"`
}

// UpdateProjectSettingsRequestUpdatePeriodicReport: update project settings request update periodic report.
type UpdateProjectSettingsRequestUpdatePeriodicReport struct {
	// Enabled: (Optional) Enable or disable periodic report notifications.
	Enabled *bool `json:"enabled"`

	// Frequency: (Optional) Frequency at which you receive periodic report notifications.
	// Default value: unknown_frequency
	Frequency *ProjectSettingsPeriodicReportFrequency `json:"frequency"`

	// SendingHour: (Optional) Hour at which you receive periodic report notifications.
	SendingHour *uint32 `json:"sending_hour"`

	// SendingDay: (Optional) On which day you receive periodic report notifications (1-7 weekly, 1-28 monthly).
	SendingDay *uint32 `json:"sending_day"`
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

// CreateDomainRequest: create domain request.
type CreateDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the project to which the domain belongs.
	ProjectID string `json:"project_id"`

	// DomainName: fully qualified domain dame.
	DomainName string `json:"domain_name"`

	// AcceptTos: accept Scaleway's Terms of Service.
	AcceptTos bool `json:"accept_tos"`

	// Autoconfig: activate auto-configuration of the domain's DNS zone.
	Autoconfig bool `json:"autoconfig"`
}

// CreateEmailRequest: create email request.
type CreateEmailRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// From: sender information. Must be from a checked domain declared in the Project.
	From *CreateEmailRequestAddress `json:"from"`

	// To: an array of the primary recipient's information.
	To []*CreateEmailRequestAddress `json:"to"`

	// Cc: an array of the carbon copy recipient's information.
	Cc []*CreateEmailRequestAddress `json:"cc"`

	// Bcc: an array of the blind carbon copy recipient's information.
	Bcc []*CreateEmailRequestAddress `json:"bcc"`

	// Subject: subject of the email.
	Subject string `json:"subject"`

	// Text: text content.
	Text string `json:"text"`

	// HTML: HTML content.
	HTML string `json:"html"`

	// ProjectID: ID of the Project in which to create the email.
	ProjectID string `json:"project_id"`

	// Attachments: array of attachments.
	Attachments []*CreateEmailRequestAttachment `json:"attachments"`

	// SendBefore: maximum date to deliver the email.
	SendBefore *time.Time `json:"send_before,omitempty"`

	// AdditionalHeaders: array of additional headers as key-value.
	AdditionalHeaders []*CreateEmailRequestHeader `json:"additional_headers"`
}

// CreateEmailResponse: create email response.
type CreateEmailResponse struct {
	// Emails: single page of emails matching the requested criteria.
	Emails []*Email `json:"emails"`
}

// CreateWebhookRequest: create webhook request.
type CreateWebhookRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainID: ID of the Domain to watch for triggering events.
	DomainID string `json:"domain_id"`

	// ProjectID: ID of the project to which the Webhook belongs.
	ProjectID string `json:"project_id"`

	// Name: name of the Webhook.
	Name string `json:"name"`

	// EventTypes: list of event types that will trigger an event.
	EventTypes []WebhookEventType `json:"event_types"`

	// SnsArn: scaleway SNS ARN topic to push the events to.
	SnsArn string `json:"sns_arn"`
}

// DeleteWebhookRequest: delete webhook request.
type DeleteWebhookRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WebhookID: ID of the Webhook to delete.
	WebhookID string `json:"-"`
}

// DomainLastStatus: domain last status.
type DomainLastStatus struct {
	// DomainID: the ID of the domain.
	DomainID string `json:"domain_id"`

	// DomainName: the domain name (example.com).
	DomainName string `json:"domain_name"`

	// SpfRecord: the SPF record verification data.
	SpfRecord *DomainLastStatusSpfRecord `json:"spf_record"`

	// DkimRecord: the DKIM record verification data.
	DkimRecord *DomainLastStatusDkimRecord `json:"dkim_record"`

	// DmarcRecord: the DMARC record verification data.
	DmarcRecord *DomainLastStatusDmarcRecord `json:"dmarc_record"`

	// AutoconfigState: the verification state of domain auto-configuration.
	AutoconfigState *DomainLastStatusAutoconfigState `json:"autoconfig_state"`
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

	// EmailID: ID of the email to retrieve.
	EmailID string `json:"-"`
}

// GetProjectSettingsRequest: get project settings request.
type GetProjectSettingsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the project.
	ProjectID string `json:"-"`
}

// GetStatisticsRequest: get statistics request.
type GetStatisticsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: (Optional) Number of emails for this Project.
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

// GetWebhookRequest: get webhook request.
type GetWebhookRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WebhookID: ID of the Webhook to check.
	WebhookID string `json:"-"`
}

// ListDomainsRequest: list domains request.
type ListDomainsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: requested page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`

	// PageSize: requested page size. Value must be between 1 and 1000.
	PageSize *uint32 `json:"-"`

	// ProjectID: (Optional) ID of the Project in which to list the domains.
	ProjectID *string `json:"-"`

	// Status: (Optional) List domains under specific statuses.
	Status []DomainStatus `json:"-"`

	// OrganizationID: (Optional) ID of the Organization in which to list the domains.
	OrganizationID *string `json:"-"`

	// Name: (Optional) Names of the domains to list.
	Name *string `json:"-"`
}

// ListDomainsResponse: list domains response.
type ListDomainsResponse struct {
	// TotalCount: number of domains that match the request (without pagination).
	TotalCount uint32 `json:"total_count"`

	// Domains: single page of domains matching the requested criteria.
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

	// ProjectID: (Optional) ID of the Project in which to list the emails.
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

	// Deprecated: MailTo: list emails sent to this recipient's email address.
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

// ListWebhookEventsRequest: list webhook events request.
type ListWebhookEventsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WebhookID: ID of the Webhook linked to the events.
	WebhookID string `json:"-"`

	// OrderBy: (Optional) List Webhook events corresponding to specific criteria.
	// Default value: created_at_desc
	OrderBy ListWebhookEventsRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`

	// PageSize: requested page size. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// EmailID: ID of the email linked to the events.
	EmailID *string `json:"-"`

	// EventTypes: list of event types linked to the events.
	EventTypes []WebhookEventType `json:"-"`

	// Statuses: list of event statuses.
	Statuses []WebhookEventStatus `json:"-"`

	// ProjectID: ID of the webhook Project.
	ProjectID *string `json:"-"`

	// OrganizationID: ID of the webhook Organization.
	OrganizationID *string `json:"-"`

	// DomainID: ID of the domain to watch for triggering events.
	DomainID *string `json:"-"`
}

// ListWebhookEventsResponse: list webhook events response.
type ListWebhookEventsResponse struct {
	// TotalCount: number of Webhook events matching the requested criteria.
	TotalCount uint64 `json:"total_count"`

	// WebhookEvents: single page of Webhook events matching the requested criteria.
	WebhookEvents []*WebhookEvent `json:"webhook_events"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListWebhookEventsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListWebhookEventsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListWebhookEventsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.WebhookEvents = append(r.WebhookEvents, results.WebhookEvents...)
	r.TotalCount += uint64(len(results.WebhookEvents))
	return uint64(len(results.WebhookEvents)), nil
}

// ListWebhooksRequest: list webhooks request.
type ListWebhooksRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: (Optional) List Webhooks corresponding to specific criteria.
	// Default value: created_at_desc
	OrderBy ListWebhooksRequestOrderBy `json:"-"`

	// Page: (Optional) Requested page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`

	// PageSize: (Optional) Requested page size. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// ProjectID: (Optional) ID of the Project for which to list the Webhooks.
	ProjectID *string `json:"-"`

	// OrganizationID: (Optional) ID of the Organization for which to list the Webhooks.
	OrganizationID *string `json:"-"`

	// DomainID: (Optional) ID of the Domain for which to list the Webhooks.
	DomainID *string `json:"-"`
}

// ListWebhooksResponse: list webhooks response.
type ListWebhooksResponse struct {
	// TotalCount: number of Webhooks matching the requested criteria.
	TotalCount uint64 `json:"total_count"`

	// Webhooks: single page of Webhooks matching the requested criteria.
	Webhooks []*Webhook `json:"webhooks"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListWebhooksResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListWebhooksResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListWebhooksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Webhooks = append(r.Webhooks, results.Webhooks...)
	r.TotalCount += uint64(len(results.Webhooks))
	return uint64(len(results.Webhooks)), nil
}

// ProjectSettings: project settings.
type ProjectSettings struct {
	// PeriodicReport: information about your periodic report.
	PeriodicReport *ProjectSettingsPeriodicReport `json:"periodic_report"`
}

// RevokeDomainRequest: revoke domain request.
type RevokeDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainID: ID of the domain to delete.
	DomainID string `json:"-"`
}

// Statistics: statistics.
type Statistics struct {
	// TotalCount: total number of emails matching the requested criteria.
	TotalCount uint32 `json:"total_count"`

	// NewCount: number of emails still in the `new` transient state. This means emails received from the API but not yet processed.
	NewCount uint32 `json:"new_count"`

	// SendingCount: number of emails still in the `sending` transient state. This means emails received from the API but not yet in their final status.
	SendingCount uint32 `json:"sending_count"`

	// SentCount: number of emails in the final `sent` state. This means emails that have been delivered to the target mail system.
	SentCount uint32 `json:"sent_count"`

	// FailedCount: number of emails in the final `failed` state. This means emails that have been refused by the target mail system with a final error status.
	FailedCount uint32 `json:"failed_count"`

	// CanceledCount: number of emails in the final `canceled` state. This means emails that have been canceled upon request.
	CanceledCount uint32 `json:"canceled_count"`
}

// UpdateDomainRequest: update domain request.
type UpdateDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainID: ID of the domain to update.
	DomainID string `json:"-"`

	// Autoconfig: (Optional) If set to true, activate auto-configuration of the domain's DNS zone.
	Autoconfig *bool `json:"autoconfig,omitempty"`
}

// UpdateProjectSettingsRequest: update project settings request.
type UpdateProjectSettingsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the project.
	ProjectID string `json:"-"`

	// PeriodicReport: periodic report update details - all fields are optional.
	PeriodicReport *UpdateProjectSettingsRequestUpdatePeriodicReport `json:"periodic_report,omitempty"`
}

// UpdateWebhookRequest: update webhook request.
type UpdateWebhookRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// WebhookID: ID of the Webhook to update.
	WebhookID string `json:"-"`

	// Name: name of the Webhook to update.
	Name *string `json:"name,omitempty"`

	// EventTypes: list of event types to update.
	EventTypes []WebhookEventType `json:"event_types"`

	// SnsArn: scaleway SNS ARN topic to update.
	SnsArn *string `json:"sns_arn,omitempty"`
}

// This API allows you to manage your Transactional Email services.
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

// CreateEmail: You must specify the `region`, the sender and the recipient's information and the `project_id` to send an email from a checked domain. The subject of the email must contain at least 6 characters.
func (s *API) CreateEmail(req *CreateEmailRequest, opts ...scw.RequestOption) (*CreateEmailResponse, error) {
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails",
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

// GetEmail: Retrieve information about a specific email using the `email_id` and `region` parameters.
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails/" + fmt.Sprint(req.EmailID) + "",
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails",
		Query:  query,
	}

	var resp ListEmailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetStatistics: Get information on your emails' statuses.
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/statistics",
		Query:  query,
	}

	var resp Statistics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelEmail: You can cancel the sending of an email if it has not been sent yet. You must specify the `region` and the `email_id` of the email you want to cancel.
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/emails/" + fmt.Sprint(req.EmailID) + "/cancel",
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

// CreateDomain: You must specify the `region`, `project_id` and `domain_name` to register a domain in a specific Project.
func (s *API) CreateDomain(req *CreateDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains",
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

// GetDomain: Retrieve information about a specific domain using the `region` and `domain_id` parameters. Monitor your domain's reputation and improve **average** and **bad** reputation statuses, using your domain's **Email activity** tab on the [Scaleway console](https://console.scaleway.com/transactional-email/domains) to get a more detailed report. Check out our [dedicated documentation](https://www.scaleway.com/en/docs/managed-services/transactional-email/reference-content/understanding-tem-reputation-score/) to improve your domain's reputation.
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDomains: Retrieve domains in a specific Project or in a specific Organization using the `region` parameter.
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains",
		Query:  query,
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RevokeDomain: You must specify the domain you want to delete by the `region` and `domain_id`. Deleting a domain is permanent and cannot be undone.
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/revoke",
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

// CheckDomain: Perform an immediate DNS check of a domain using the `region` and `domain_id` parameters.
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/check",
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "/verification",
	}

	var resp DomainLastStatus

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDomain: Update a domain auto-configuration.
func (s *API) UpdateDomain(req *UpdateDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
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
		Method: "PATCH",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainID) + "",
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

// CreateWebhook: Create a new Webhook triggered by a list of event types and pushed to a Scaleway SNS ARN.
func (s *API) CreateWebhook(req *CreateWebhookRequest, opts ...scw.RequestOption) (*Webhook, error) {
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
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/webhooks",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Webhook

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListWebhooks: Retrieve Webhooks in a specific Project or in a specific Organization using the `region` parameter.
func (s *API) ListWebhooks(req *ListWebhooksRequest, opts ...scw.RequestOption) (*ListWebhooksResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "domain_id", req.DomainID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/webhooks",
		Query:  query,
	}

	var resp ListWebhooksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWebhook: Retrieve information about a specific Webhook using the `webhook_id` and `region` parameters.
func (s *API) GetWebhook(req *GetWebhookRequest, opts ...scw.RequestOption) (*Webhook, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WebhookID) == "" {
		return nil, errors.New("field WebhookID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/webhooks/" + fmt.Sprint(req.WebhookID) + "",
	}

	var resp Webhook

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateWebhook: Update a Webhook events type, SNS ARN or name.
func (s *API) UpdateWebhook(req *UpdateWebhookRequest, opts ...scw.RequestOption) (*Webhook, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WebhookID) == "" {
		return nil, errors.New("field WebhookID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/webhooks/" + fmt.Sprint(req.WebhookID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Webhook

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteWebhook: You must specify the Webhook you want to delete by the `region` and `webhook_id`. Deleting a Webhook is permanent and cannot be undone.
func (s *API) DeleteWebhook(req *DeleteWebhookRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WebhookID) == "" {
		return errors.New("field WebhookID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/webhooks/" + fmt.Sprint(req.WebhookID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListWebhookEvents: Retrieve the list of Webhook events triggered from a specific Webhook or for a specific Project or Organization. You must specify the `region`.
func (s *API) ListWebhookEvents(req *ListWebhookEventsRequest, opts ...scw.RequestOption) (*ListWebhookEventsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "email_id", req.EmailID)
	parameter.AddToQuery(query, "event_types", req.EventTypes)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "domain_id", req.DomainID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.WebhookID) == "" {
		return nil, errors.New("field WebhookID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/webhooks/" + fmt.Sprint(req.WebhookID) + "/events",
		Query:  query,
	}

	var resp ListWebhookEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetProjectSettings: Retrieve the project settings including periodic reports.
func (s *API) GetProjectSettings(req *GetProjectSettingsRequest, opts ...scw.RequestOption) (*ProjectSettings, error) {
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

	if fmt.Sprint(req.ProjectID) == "" {
		return nil, errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/project/" + fmt.Sprint(req.ProjectID) + "/settings",
	}

	var resp ProjectSettings

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateProjectSettings: Update the project settings including periodic reports.
func (s *API) UpdateProjectSettings(req *UpdateProjectSettingsRequest, opts ...scw.RequestOption) (*ProjectSettings, error) {
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

	if fmt.Sprint(req.ProjectID) == "" {
		return nil, errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/transactional-email/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/project/" + fmt.Sprint(req.ProjectID) + "/settings",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ProjectSettings

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

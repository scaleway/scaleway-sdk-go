// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package account_admin provides methods and message types of the account_admin v3 API.
package account_admin

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
	std "github.com/scaleway/scaleway-sdk-go/api/std"
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

type ContractType string

const (
	// Unknown contract type.
	ContractTypeUnknownType = ContractType("unknown_type")
	// A term of services contract.
	ContractTypeTos = ContractType("tos")
	// A product contract.
	ContractTypeProduct = ContractType("product")
	// A legal contract.
	ContractTypeLegal = ContractType("legal")
)

func (enum ContractType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum ContractType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContractType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContractType(ContractType(tmp).String())
	return nil
}

type ListAccountsRequestOrderBy string

const (
	// Creation date ascending.
	ListAccountsRequestOrderByCreatedAtAsc = ListAccountsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListAccountsRequestOrderByCreatedAtDesc = ListAccountsRequestOrderBy("created_at_desc")
)

func (enum ListAccountsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListAccountsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListAccountsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListAccountsRequestOrderBy(ListAccountsRequestOrderBy(tmp).String())
	return nil
}

type ListBlacklistedDomainsRequestOrderBy string

const (
	// Creation date ascending.
	ListBlacklistedDomainsRequestOrderByCreatedAtAsc = ListBlacklistedDomainsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListBlacklistedDomainsRequestOrderByCreatedAtDesc = ListBlacklistedDomainsRequestOrderBy("created_at_desc")
	// Domain name ascending.
	ListBlacklistedDomainsRequestOrderByDomainAsc = ListBlacklistedDomainsRequestOrderBy("domain_asc")
	// Domain name descending.
	ListBlacklistedDomainsRequestOrderByDomainDesc = ListBlacklistedDomainsRequestOrderBy("domain_desc")
)

func (enum ListBlacklistedDomainsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListBlacklistedDomainsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListBlacklistedDomainsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListBlacklistedDomainsRequestOrderBy(ListBlacklistedDomainsRequestOrderBy(tmp).String())
	return nil
}

type ListContractSignaturesRequestOrderBy string

const (
	// Signing date ascending.
	ListContractSignaturesRequestOrderBySignedAtAsc = ListContractSignaturesRequestOrderBy("signed_at_asc")
	// Signing date descending.
	ListContractSignaturesRequestOrderBySignedAtDesc = ListContractSignaturesRequestOrderBy("signed_at_desc")
	// Expiration date ascending.
	ListContractSignaturesRequestOrderByExpiresAtAsc = ListContractSignaturesRequestOrderBy("expires_at_asc")
	// Expiration date descending.
	ListContractSignaturesRequestOrderByExpiresAtDesc = ListContractSignaturesRequestOrderBy("expires_at_desc")
)

func (enum ListContractSignaturesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "signed_at_asc"
	}
	return string(enum)
}

func (enum ListContractSignaturesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListContractSignaturesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListContractSignaturesRequestOrderBy(ListContractSignaturesRequestOrderBy(tmp).String())
	return nil
}

type ListContractsRequestOrderBy string

const (
	// Creation date ascending.
	ListContractsRequestOrderByCreatedAtAsc = ListContractsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListContractsRequestOrderByCreatedAtDesc = ListContractsRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListContractsRequestOrderByUpdatedAtAsc = ListContractsRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListContractsRequestOrderByUpdatedAtDesc = ListContractsRequestOrderBy("updated_at_desc")
)

func (enum ListContractsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListContractsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListContractsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListContractsRequestOrderBy(ListContractsRequestOrderBy(tmp).String())
	return nil
}

type ListInvitationsRequestOrderBy string

const (
	// Creation date ascending.
	ListInvitationsRequestOrderByCreatedAtAsc = ListInvitationsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListInvitationsRequestOrderByCreatedAtDesc = ListInvitationsRequestOrderBy("created_at_desc")
)

func (enum ListInvitationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListInvitationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInvitationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInvitationsRequestOrderBy(ListInvitationsRequestOrderBy(tmp).String())
	return nil
}

type ListMailingListSubscriptionsRequestOrderBy string

const (
	// Name ascending.
	ListMailingListSubscriptionsRequestOrderByNameAsc = ListMailingListSubscriptionsRequestOrderBy("name_asc")
	// Name descending.
	ListMailingListSubscriptionsRequestOrderByNameDesc = ListMailingListSubscriptionsRequestOrderBy("name_desc")
)

func (enum ListMailingListSubscriptionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListMailingListSubscriptionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListMailingListSubscriptionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListMailingListSubscriptionsRequestOrderBy(ListMailingListSubscriptionsRequestOrderBy(tmp).String())
	return nil
}

type ListOrganizationDeactivationLogsRequestOrderBy string

const (
	// Creation date ascending.
	ListOrganizationDeactivationLogsRequestOrderByCreatedAtAsc = ListOrganizationDeactivationLogsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListOrganizationDeactivationLogsRequestOrderByCreatedAtDesc = ListOrganizationDeactivationLogsRequestOrderBy("created_at_desc")
)

func (enum ListOrganizationDeactivationLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListOrganizationDeactivationLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOrganizationDeactivationLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOrganizationDeactivationLogsRequestOrderBy(ListOrganizationDeactivationLogsRequestOrderBy(tmp).String())
	return nil
}

type ListPhoneValidationsRequestOrderBy string

const (
	// Sent date ascending.
	ListPhoneValidationsRequestOrderBySentAtAsc = ListPhoneValidationsRequestOrderBy("sent_at_asc")
	// Sent date descending.
	ListPhoneValidationsRequestOrderBySentAtDesc = ListPhoneValidationsRequestOrderBy("sent_at_desc")
)

func (enum ListPhoneValidationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "sent_at_asc"
	}
	return string(enum)
}

func (enum ListPhoneValidationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPhoneValidationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPhoneValidationsRequestOrderBy(ListPhoneValidationsRequestOrderBy(tmp).String())
	return nil
}

type ListProjectsRequestOrderBy string

const (
	// Creation date ascending.
	ListProjectsRequestOrderByCreatedAtAsc = ListProjectsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListProjectsRequestOrderByCreatedAtDesc = ListProjectsRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListProjectsRequestOrderByUpdatedAtAsc = ListProjectsRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListProjectsRequestOrderByUpdatedAtDesc = ListProjectsRequestOrderBy("updated_at_desc")
	// Name ascending.
	ListProjectsRequestOrderByNameAsc = ListProjectsRequestOrderBy("name_asc")
	// Name descending.
	ListProjectsRequestOrderByNameDesc = ListProjectsRequestOrderBy("name_desc")
)

func (enum ListProjectsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListProjectsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListProjectsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListProjectsRequestOrderBy(ListProjectsRequestOrderBy(tmp).String())
	return nil
}

type ListSupportPlansRequestOrderBy string

const (
	// Start date ascending.
	ListSupportPlansRequestOrderByStartedAtAsc = ListSupportPlansRequestOrderBy("started_at_asc")
	// Start date descending.
	ListSupportPlansRequestOrderByStartedAtDesc = ListSupportPlansRequestOrderBy("started_at_desc")
	// Stop date ascending.
	ListSupportPlansRequestOrderByStoppedAtAsc = ListSupportPlansRequestOrderBy("stopped_at_asc")
	// Stop date descending.
	ListSupportPlansRequestOrderByStoppedAtDesc = ListSupportPlansRequestOrderBy("stopped_at_desc")
)

func (enum ListSupportPlansRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "started_at_asc"
	}
	return string(enum)
}

func (enum ListSupportPlansRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSupportPlansRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSupportPlansRequestOrderBy(ListSupportPlansRequestOrderBy(tmp).String())
	return nil
}

type ListWarningLogsRequestOrderBy string

const (
	// Creation date ascending.
	ListWarningLogsRequestOrderByCreatedAtAsc = ListWarningLogsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListWarningLogsRequestOrderByCreatedAtDesc = ListWarningLogsRequestOrderBy("created_at_desc")
)

func (enum ListWarningLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListWarningLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListWarningLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListWarningLogsRequestOrderBy(ListWarningLogsRequestOrderBy(tmp).String())
	return nil
}

type ListWarningsRequestOrderBy string

const (
	// Open date ascending.
	ListWarningsRequestOrderByOpenedAtAsc = ListWarningsRequestOrderBy("opened_at_asc")
	// Open date descending.
	ListWarningsRequestOrderByOpenedAtDesc = ListWarningsRequestOrderBy("opened_at_desc")
	// Update date ascending.
	ListWarningsRequestOrderByUpdatedAtAsc = ListWarningsRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListWarningsRequestOrderByUpdatedAtDesc = ListWarningsRequestOrderBy("updated_at_desc")
	// Lock date ascending.
	ListWarningsRequestOrderByLockedAtAsc = ListWarningsRequestOrderBy("locked_at_asc")
	// Lock date descending.
	ListWarningsRequestOrderByLockedAtDesc = ListWarningsRequestOrderBy("locked_at_desc")
)

func (enum ListWarningsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "opened_at_asc"
	}
	return string(enum)
}

func (enum ListWarningsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListWarningsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListWarningsRequestOrderBy(ListWarningsRequestOrderBy(tmp).String())
	return nil
}

type OrganizationDeactivationLogReason string

const (
	// Unknown reason.
	OrganizationDeactivationLogReasonUnknownReason = OrganizationDeactivationLogReason("unknown_reason")
	// The client no longer need Scaleway's products.
	OrganizationDeactivationLogReasonNoLongerNeeded = OrganizationDeactivationLogReason("no_longer_needed")
	// The client finds it too expensive.
	OrganizationDeactivationLogReasonTooExpensive = OrganizationDeactivationLogReason("too_expensive")
	// The client needs some features we don't have.
	OrganizationDeactivationLogReasonMissingFeature = OrganizationDeactivationLogReason("missing_feature")
	// The client has had technical issues.
	OrganizationDeactivationLogReasonTechnicalIssue = OrganizationDeactivationLogReason("technical_issue")
	// The client just wanted to test the service.
	OrganizationDeactivationLogReasonTestOnly = OrganizationDeactivationLogReason("test_only")
	// The client finds it difficult to use.
	OrganizationDeactivationLogReasonDifficultToUse = OrganizationDeactivationLogReason("difficult_to_use")
	// The client is using another account.
	OrganizationDeactivationLogReasonUsingAnotherAccount = OrganizationDeactivationLogReason("using_another_account")
	// Other reason.
	OrganizationDeactivationLogReasonOther = OrganizationDeactivationLogReason("other")
	// Closed automatically.
	OrganizationDeactivationLogReasonClosedAutomatically = OrganizationDeactivationLogReason("closed_automatically")
)

func (enum OrganizationDeactivationLogReason) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_reason"
	}
	return string(enum)
}

func (enum OrganizationDeactivationLogReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationDeactivationLogReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationDeactivationLogReason(OrganizationDeactivationLogReason(tmp).String())
	return nil
}

type SupportPlanLevel string

const (
	// Unknown level.
	SupportPlanLevelUnknownLevel = SupportPlanLevel("unknown_level")
	// Basic.
	SupportPlanLevelBasic = SupportPlanLevel("basic")
	// Developer.
	SupportPlanLevelDeveloper = SupportPlanLevel("developer")
	// Business.
	SupportPlanLevelBusiness = SupportPlanLevel("business")
	// Enterprise.
	SupportPlanLevelEnterprise = SupportPlanLevel("enterprise")
	// Bronze.
	SupportPlanLevelBronze = SupportPlanLevel("bronze")
	// Silver.
	SupportPlanLevelSilver = SupportPlanLevel("silver")
	// Gold.
	SupportPlanLevelGold = SupportPlanLevel("gold")
	// Platinum.
	SupportPlanLevelPlatinum = SupportPlanLevel("platinum")
)

func (enum SupportPlanLevel) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_level"
	}
	return string(enum)
}

func (enum SupportPlanLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SupportPlanLevel) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SupportPlanLevel(SupportPlanLevel(tmp).String())
	return nil
}

type WarningLogAction string

const (
	// Unknown action.
	WarningLogActionUnknownAction = WarningLogAction("unknown_action")
	// The warning log was created.
	WarningLogActionCreated = WarningLogAction("created")
	// The warning log was updated.
	WarningLogActionUpdated = WarningLogAction("updated")
)

func (enum WarningLogAction) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_action"
	}
	return string(enum)
}

func (enum WarningLogAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *WarningLogAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = WarningLogAction(WarningLogAction(tmp).String())
	return nil
}

type WarningReason string

const (
	// Unknown reason.
	WarningReasonUnknownReason = WarningReason("unknown_reason")
	// Security concerns.
	WarningReasonSecurityConcerns = WarningReason("security_concerns")
	// Security issue.
	WarningReasonSecurityIssue = WarningReason("security_issue")
	// Critical security issue.
	WarningReasonCriticalSecurityIssue = WarningReason("critical_security_issue")
	// Locked for abuse.
	WarningReasonLockedForAbuse = WarningReason("locked_for_abuse")
	// Network abuse.
	WarningReasonNetworkAbuse = WarningReason("network_abuse")
	// Rate limiting.
	WarningReasonRateLimiting = WarningReason("rate_limiting")
	// Locked by reseller.
	WarningReasonLockedByReseller = WarningReason("locked_by_reseller")
	// Reseller locked.
	WarningReasonResellerLocked = WarningReason("reseller_locked")
	// Invoice payment failure.
	WarningReasonInvoicePaymentFailure = WarningReason("invoice_payment_failure")
	// Got chargeback.
	WarningReasonGotChargeback = WarningReason("got_chargeback")
	// Magic code rate limiting.
	WarningReasonMagicCodeRateLimiting = WarningReason("magic_code_rate_limiting")
	// Validate TOS.
	WarningReasonValidateTos = WarningReason("validate_tos")
	// Validate email.
	WarningReasonValidateEmail = WarningReason("validate_email")
	// Confirm email change.
	WarningReasonConfirmEmailChange = WarningReason("confirm_email_change")
	// Phone missing.
	WarningReasonPhoneMissing = WarningReason("phone_missing")
	// Payment info missing.
	WarningReasonPaymentInfoMissing = WarningReason("payment_info_missing")
	// Billing info missing.
	WarningReasonBillingInfoMissing = WarningReason("billing_info_missing")
	// Account closed.
	WarningReasonAccountClosed = WarningReason("account_closed")
	// Card expired or soon.
	WarningReasonCardExpiredOrSoon = WarningReason("card_expired_or_soon")
	// Magic code validation timeout.
	WarningReasonMagicCodeValidationTimeout = WarningReason("magic_code_validation_timeout")
	// Account anonymized.
	WarningReasonAccountAnonymized = WarningReason("account_anonymized")
	// Budget delete resources.
	WarningReasonBudgetDeleteResources = WarningReason("budget_delete_resources")
	// Budget forbid new resources.
	WarningReasonBudgetForbidNewResources = WarningReason("budget_forbid_new_resources")
	// Validate KYC.
	WarningReasonValidateKyc = WarningReason("validate_kyc")
	// Trusted level security check.
	WarningReasonTrustedLevelSecurityCheck = WarningReason("trusted_level_security_check")
	// GDPR delete.
	WarningReasonGdprDelete = WarningReason("gdpr_delete")
	// Online invoice payment failure.
	WarningReasonOnlineInvoicePaymentFailure = WarningReason("online_invoice_payment_failure")
)

func (enum WarningReason) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_reason"
	}
	return string(enum)
}

func (enum WarningReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *WarningReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = WarningReason(WarningReason(tmp).String())
	return nil
}

// Contract: contract.
type Contract struct {
	// ID: ID of the contract.
	ID string `json:"id"`

	// ContractType: the type of the contract.
	// Default value: unknown_type
	ContractType ContractType `json:"contract_type"`

	// Name: the name of the contract.
	Name string `json:"name"`

	// Version: the version of the contract.
	Version uint32 `json:"version"`

	// ThirdPartyID: the third party organization bound to the contract.
	ThirdPartyID string `json:"third_party_id"`

	// CreatedAt: the creation date of the contract.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: the last modification date of the contract.
	UpdatedAt *time.Time `json:"updated_at"`
}

// SupportPlan: support plan.
type SupportPlan struct {
	// ID: ID of the support plan.
	ID string `json:"id"`

	// Level: the level of the support plan.
	// Default value: unknown_level
	Level SupportPlanLevel `json:"level"`

	// StartedAt: creation date of the support plan.
	StartedAt *time.Time `json:"started_at"`

	// StoppedAt: end date of the support plan.
	StoppedAt *time.Time `json:"stopped_at"`

	// OrganizationID: ID of the organization.
	OrganizationID string `json:"organization_id"`

	// Price: the price of the support plan. Only for Platinum.
	Price *scw.Money `json:"price"`
}

// TechnicalAccountManager: technical account manager.
type TechnicalAccountManager struct {
	// ID: ID of the TAM.
	ID string `json:"id"`

	// LastName: last name of the TAM.
	LastName string `json:"last_name"`

	// FirstName: first name of the TAM.
	FirstName string `json:"first_name"`

	// Email: email address of the TAM.
	Email string `json:"email"`

	// PhoneNumber: phone number of the TAM.
	PhoneNumber string `json:"phone_number"`
}

// Warning: warning.
type Warning struct {
	// ID: ID of the warning.
	ID string `json:"id"`

	// OrganizationID: the organization ID linked to the warning.
	OrganizationID string `json:"organization_id"`

	// Reason: the reason of the warning.
	// Default value: unknown_reason
	Reason WarningReason `json:"reason"`

	// OpenedAt: the creation date of the warning.
	OpenedAt *time.Time `json:"opened_at"`

	// LockedAt: the lock date of the warning.
	LockedAt *time.Time `json:"locked_at"`

	// ClosedAt: the closure date of the warning.
	ClosedAt *time.Time `json:"closed_at"`

	// ClosableByUser: if true the user is able to close the warning.
	ClosableByUser bool `json:"closable_by_user"`

	// ClosedFromIP: IP used to close the warning.
	ClosedFromIP string `json:"closed_from_ip"`

	// Code: a random string that is requested by the API to close this warning for non-admin users.
	Code string `json:"code"`

	// Note: a note associated to the warning.
	Note string `json:"note"`

	// UpdatedAt: the last update date of the warning.
	UpdatedAt *time.Time `json:"updated_at"`
}

// Account: account.
type Account struct {
	// ID: the ID of the account.
	ID string `json:"id"`

	// Email: the email associated with the account.
	Email string `json:"email"`

	// Request: content of the account creation request.
	Request *scw.JSONObject `json:"request"`

	// CreatedAt: the account was created at this date.
	CreatedAt *time.Time `json:"created_at"`

	// ExpiresAt: the account expires at this date.
	ExpiresAt *time.Time `json:"expires_at"`

	// DeletedAt: the account was deleted at this date.
	DeletedAt *time.Time `json:"deleted_at"`

	// NotifiedAt: the account was last notified at this date.
	NotifiedAt *time.Time `json:"notified_at"`
}

// BlacklistedDomain: blacklisted domain.
type BlacklistedDomain struct {
	// ID: the ID of the blacklisted domain.
	ID string `json:"id"`

	// Domain: the domain name.
	Domain string `json:"domain"`

	// CreatedAt: the creation date of the blacklisted domain.
	CreatedAt *time.Time `json:"created_at"`

	// BlockAccountCreation: if true the creation of account will be blocked.
	BlockAccountCreation bool `json:"block_account_creation"`
}

// ContractSignature: contract signature.
type ContractSignature struct {
	// ID: ID of the contract signature.
	ID string `json:"id"`

	// OrganizationID: the organization ID which signed the contract.
	OrganizationID string `json:"organization_id"`

	// CreatedAt: the creation date of the contract signature.
	CreatedAt *time.Time `json:"created_at"`

	// SignedAt: the signing date of the contract signature.
	SignedAt *time.Time `json:"signed_at"`

	// ExpiresAt: the expiration date of the contract signature.
	ExpiresAt *time.Time `json:"expires_at"`

	// Contract: the contract signed.
	Contract *Contract `json:"contract"`
}

// Invitation: invitation.
type Invitation struct {
	ID string `json:"id"`

	Email string `json:"email"`

	OrganizationID string `json:"organization_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	ExpiresAt *time.Time `json:"expires_at"`

	DeletedAt *time.Time `json:"deleted_at"`
}

// MailingListSubscription: mailing list subscription.
type MailingListSubscription struct {
	// ID: the ID of the mailing list.
	ID string `json:"id"`

	// Name: the name of the mailing list.
	Name string `json:"name"`

	// Description: the description of the mailing list.
	Description string `json:"description"`

	// Subscribed: the status of the subscription.
	Subscribed bool `json:"subscribed"`
}

// OrganizationDeactivationLog: organization deactivation log.
type OrganizationDeactivationLog struct {
	// ID: ID of the organization deactivation log.
	ID string `json:"id"`

	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"organization_id"`

	// CreatedAt: creation date of the organization deactivation log.
	CreatedAt *time.Time `json:"created_at"`

	// Reasons: reasons of the organization deactivation.
	Reasons []OrganizationDeactivationLogReason `json:"reasons"`

	// Details: more details about why the Organization is being closed.
	Details string `json:"details"`
}

// PhoneValidation: phone validation.
type PhoneValidation struct {
	// ID: the ID of the phone validation.
	ID string `json:"id"`

	// PhoneNumber: the phone number in international format (E.164).
	PhoneNumber string `json:"phone_number"`

	// Code: the code required to validate the phone number.
	Code string `json:"code"`

	// CarrierType: the type of carrier for this phone number (mobile, landline, voip, etc.).
	CarrierType string `json:"carrier_type"`

	// SentAt: the validation request has been sent at this date.
	SentAt *time.Time `json:"sent_at"`

	// ValidatedAt: the phone number has been validated at this date.
	ValidatedAt *time.Time `json:"validated_at"`
}

// Project: project.
type Project struct {
	// ID: the ID of the project.
	ID string `json:"id"`

	// Name: the name of the project.
	Name string `json:"name"`

	// OrganizationID: the organization ID of the project.
	OrganizationID string `json:"organization_id"`

	// CreatedAt: the creation date of the project.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: the update date of the project.
	UpdatedAt *time.Time `json:"updated_at"`

	// DeletedAt: the deletion date of the project.
	DeletedAt *time.Time `json:"deleted_at"`

	// Description: the description of the project.
	Description string `json:"description"`
}

// WarningLog: warning log.
type WarningLog struct {
	// ID: ID of the warning log.
	ID string `json:"id"`

	// CreatedAt: the creation date of the warning log.
	CreatedAt *time.Time `json:"created_at"`

	// OrganizationWarningID: the organization warning linked to that log entry.
	OrganizationWarningID string `json:"organization_warning_id"`

	// FromPrincipalID: the principal ID that performed the action.
	FromPrincipalID string `json:"from_principal_id"`

	// Action: action performed on the organization warning.
	// Default value: unknown_action
	Action WarningLogAction `json:"action"`

	// Data: current values of fields modified by the action on the organization warning.
	Data *scw.JSONObject `json:"data"`

	// PreviousData: old values of fields before the action on the organization warning.
	PreviousData *scw.JSONObject `json:"previous_data"`
}

// StopMailingListSubscriptionsRequestUnsubscription: stop mailing list subscriptions request unsubscription.
type StopMailingListSubscriptionsRequestUnsubscription struct {
	OrganizationID string `json:"organization_id"`

	UnsubscribedAt *time.Time `json:"unsubscribed_at"`
}

// ActiveSupportPlan: active support plan.
type ActiveSupportPlan struct {
	// SupportPlan: active support plan information.
	SupportPlan *SupportPlan `json:"support_plan"`

	// SupportID: support id of the organization.
	SupportID string `json:"support_id"`

	// SupportPin: support pin of the organization.
	SupportPin string `json:"support_pin"`

	// TechnicalAccountManager: technical account manager attached to the organization.
	TechnicalAccountManager *TechnicalAccountManager `json:"technical_account_manager"`
}

// BlacklistedDomainAPICreateBlacklistedDomainRequest: blacklisted domain api create blacklisted domain request.
type BlacklistedDomainAPICreateBlacklistedDomainRequest struct {
	// Domain: the domain name.
	Domain string `json:"domain"`

	// BlockAccountCreation: if true the creation of account will be blocked.
	BlockAccountCreation bool `json:"block_account_creation"`
}

// BlacklistedDomainAPIDeleteBlacklistedDomainRequest: blacklisted domain api delete blacklisted domain request.
type BlacklistedDomainAPIDeleteBlacklistedDomainRequest struct {
	// BlacklistedDomainID: the ID of the blacklisted domain.
	BlacklistedDomainID string `json:"-"`
}

// BlacklistedDomainAPIListBlacklistedDomainsRequest: blacklisted domain api list blacklisted domains request.
type BlacklistedDomainAPIListBlacklistedDomainsRequest struct {
	// Page: the page number for the returned blacklisted domaines.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of blacklisted domains per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the blacklisted domaines are ordered in the response.
	// Default value: created_at_asc
	OrderBy ListBlacklistedDomainsRequestOrderBy `json:"-"`

	// BlockAccountCreation: filter on domains blocking the creation of accounts.
	BlockAccountCreation *bool `json:"-"`
}

// CloseAllWarningsResponse: close all warnings response.
type CloseAllWarningsResponse struct {
	// Warnings: the closed warnings.
	Warnings []*Warning `json:"warnings"`
}

// ContractAPICreateContractRequest: contract api create contract request.
type ContractAPICreateContractRequest struct {
	// ContractType: the type of the contract.
	// Default value: unknown_type
	ContractType ContractType `json:"contract_type"`

	// Name: the name for the contract.
	Name string `json:"name"`

	// ThirdPartyID: the third party organization bound to the contract.
	ThirdPartyID string `json:"third_party_id"`

	// LocalizedContentFrFr: the localized fr_fr content of the contract.
	LocalizedContentFrFr []byte `json:"localized_content_fr_fr"`

	// LocalizedContentEnUs: the localized en_us content of the contract.
	LocalizedContentEnUs []byte `json:"localized_content_en_us"`
}

// ContractAPIDownloadContractRequest: contract api download contract request.
type ContractAPIDownloadContractRequest struct {
	// ContractID: the contract ID.
	ContractID string `json:"-"`

	// Locale: the locale requested for the content of the contract.
	// Default value: unknown_language_code
	Locale std.LanguageCode `json:"-"`
}

// ContractAPIGetContractRequest: contract api get contract request.
type ContractAPIGetContractRequest struct {
	ContractID string `json:"-"`
}

// ContractAPIListContractSignaturesRequest: contract api list contract signatures request.
type ContractAPIListContractSignaturesRequest struct {
	// Page: the page number for the returned contracts.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of contracts per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the contracts are ordered in the response.
	// Default value: signed_at_asc
	OrderBy ListContractSignaturesRequestOrderBy `json:"-"`

	// OrganizationID: filter on organization ID.
	OrganizationID string `json:"-"`
}

// ContractAPIListContractsRequest: contract api list contracts request.
type ContractAPIListContractsRequest struct {
	// Page: the page number for the returned contracts.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of contracts per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the contracts are ordered in the response.
	// Default value: created_at_asc
	OrderBy ListContractsRequestOrderBy `json:"-"`

	// ContractType: filter on a contract type.
	// Default value: unknown_type
	ContractType ContractType `json:"-"`
}

// GetMailingListRecipientsResponse: get mailing list recipients response.
type GetMailingListRecipientsResponse struct {
	// UserIDs: list of user IDs.
	UserIDs []string `json:"user_ids"`
}

// ListAccountsResponse: list accounts response.
type ListAccountsResponse struct {
	// TotalCount: the total number of accounts.
	TotalCount uint64 `json:"total_count"`

	// Accounts: the paginated returned accounts.
	Accounts []*Account `json:"accounts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListAccountsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListAccountsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListAccountsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Accounts = append(r.Accounts, results.Accounts...)
	r.TotalCount += uint64(len(results.Accounts))
	return uint64(len(results.Accounts)), nil
}

// ListBlacklistedDomainsResponse: list blacklisted domains response.
type ListBlacklistedDomainsResponse struct {
	// TotalCount: the total number of blacklisted domains.
	TotalCount uint64 `json:"total_count"`

	// BlacklistedDomains: the paginated returned blacklisted domains.
	BlacklistedDomains []*BlacklistedDomain `json:"blacklisted_domains"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBlacklistedDomainsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBlacklistedDomainsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListBlacklistedDomainsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.BlacklistedDomains = append(r.BlacklistedDomains, results.BlacklistedDomains...)
	r.TotalCount += uint64(len(results.BlacklistedDomains))
	return uint64(len(results.BlacklistedDomains)), nil
}

// ListContractSignaturesResponse: list contract signatures response.
type ListContractSignaturesResponse struct {
	// TotalCount: the total number of contract signatures.
	TotalCount uint64 `json:"total_count"`

	// ContractSignatures: the paginated returned contract signatures.
	ContractSignatures []*ContractSignature `json:"contract_signatures"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListContractSignaturesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListContractSignaturesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListContractSignaturesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ContractSignatures = append(r.ContractSignatures, results.ContractSignatures...)
	r.TotalCount += uint64(len(results.ContractSignatures))
	return uint64(len(results.ContractSignatures)), nil
}

// ListContractsResponse: list contracts response.
type ListContractsResponse struct {
	// TotalCount: the total number of contracts.
	TotalCount uint64 `json:"total_count"`

	// Contracts: the paginated returned contracts.
	Contracts []*Contract `json:"contracts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListContractsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListContractsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListContractsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Contracts = append(r.Contracts, results.Contracts...)
	r.TotalCount += uint64(len(results.Contracts))
	return uint64(len(results.Contracts)), nil
}

// ListInvitationsResponse: list invitations response.
type ListInvitationsResponse struct {
	// TotalCount: the total number of invitations.
	TotalCount uint64 `json:"total_count"`

	// Invitations: the paginated returned invitations.
	Invitations []*Invitation `json:"invitations"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInvitationsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInvitationsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListInvitationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Invitations = append(r.Invitations, results.Invitations...)
	r.TotalCount += uint64(len(results.Invitations))
	return uint64(len(results.Invitations)), nil
}

// ListMailingListSubscriptionsResponse: list mailing list subscriptions response.
type ListMailingListSubscriptionsResponse struct {
	// TotalCount: the number of susbcriptions.
	TotalCount uint64 `json:"total_count"`

	// MailingListSubscriptions: a list of mailing list subscriptions.
	MailingListSubscriptions []*MailingListSubscription `json:"mailing_list_subscriptions"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListMailingListSubscriptionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListMailingListSubscriptionsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListMailingListSubscriptionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.MailingListSubscriptions = append(r.MailingListSubscriptions, results.MailingListSubscriptions...)
	r.TotalCount += uint64(len(results.MailingListSubscriptions))
	return uint64(len(results.MailingListSubscriptions)), nil
}

// ListOrganizationDeactivationLogsResponse: list organization deactivation logs response.
type ListOrganizationDeactivationLogsResponse struct {
	// OrganizationDeactivationLogs: list of organization deactivation logs.
	OrganizationDeactivationLogs []*OrganizationDeactivationLog `json:"organization_deactivation_logs"`

	// TotalCount: total count of organization deactivation logs.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOrganizationDeactivationLogsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOrganizationDeactivationLogsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListOrganizationDeactivationLogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.OrganizationDeactivationLogs = append(r.OrganizationDeactivationLogs, results.OrganizationDeactivationLogs...)
	r.TotalCount += uint64(len(results.OrganizationDeactivationLogs))
	return uint64(len(results.OrganizationDeactivationLogs)), nil
}

// ListPhoneValidationsResponse: list phone validations response.
type ListPhoneValidationsResponse struct {
	// TotalCount: the total number of phone validations.
	TotalCount uint64 `json:"total_count"`

	// PhoneValidations: the paginated returned phone validations.
	PhoneValidations []*PhoneValidation `json:"phone_validations"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPhoneValidationsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPhoneValidationsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPhoneValidationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PhoneValidations = append(r.PhoneValidations, results.PhoneValidations...)
	r.TotalCount += uint64(len(results.PhoneValidations))
	return uint64(len(results.PhoneValidations)), nil
}

// ListProjectsResponse: list projects response.
type ListProjectsResponse struct {
	// TotalCount: the total number of projects.
	TotalCount uint64 `json:"total_count"`

	// Projects: the paginated returned projects.
	Projects []*Project `json:"projects"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListProjectsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListProjectsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListProjectsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Projects = append(r.Projects, results.Projects...)
	r.TotalCount += uint64(len(results.Projects))
	return uint64(len(results.Projects)), nil
}

// ListSupportPlansResponse: list support plans response.
type ListSupportPlansResponse struct {
	// SupportPlans: list of support plans.
	SupportPlans []*SupportPlan `json:"support_plans"`

	// TotalCount: total count of support plans.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSupportPlansResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSupportPlansResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListSupportPlansResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SupportPlans = append(r.SupportPlans, results.SupportPlans...)
	r.TotalCount += uint64(len(results.SupportPlans))
	return uint64(len(results.SupportPlans)), nil
}

// ListWarningLogsResponse: list warning logs response.
type ListWarningLogsResponse struct {
	// WarningLogs: the paginated returned warning logs.
	WarningLogs []*WarningLog `json:"warning_logs"`

	// TotalCount: the total number of warning logs.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListWarningLogsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListWarningLogsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListWarningLogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.WarningLogs = append(r.WarningLogs, results.WarningLogs...)
	r.TotalCount += uint64(len(results.WarningLogs))
	return uint64(len(results.WarningLogs)), nil
}

// ListWarningsResponse: list warnings response.
type ListWarningsResponse struct {
	// TotalCount: the total number of warnings.
	TotalCount uint64 `json:"total_count"`

	// Warnings: the paginated returned warnings.
	Warnings []*Warning `json:"warnings"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListWarningsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListWarningsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListWarningsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Warnings = append(r.Warnings, results.Warnings...)
	r.TotalCount += uint64(len(results.Warnings))
	return uint64(len(results.Warnings)), nil
}

// OrganizationAPIChangeSupportPlanRequest: organization api change support plan request.
type OrganizationAPIChangeSupportPlanRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`

	// Level: the level of support plan. Can only be basic, silver or gold.
	// Default value: unknown_level
	Level SupportPlanLevel `json:"level"`

	// Price: the price of the support plan. Only for Platinum.
	Price *scw.Money `json:"price,omitempty"`

	// DisableEmail: if true disable email notification to the customer.
	DisableEmail bool `json:"disable_email"`
}

// OrganizationAPICloseAllWarningsRequest: organization api close all warnings request.
type OrganizationAPICloseAllWarningsRequest struct {
	// OrganizationID: the ID of the organization.
	OrganizationID string `json:"organization_id"`

	// Reason: reason to look for.
	// Default value: unknown_reason
	Reason WarningReason `json:"reason"`

	// Note: add a note on warning.
	Note string `json:"note"`
}

// OrganizationAPICloseOrganizationRequest: organization api close organization request.
type OrganizationAPICloseOrganizationRequest struct {
	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"-"`

	// Reasons: reasons to close the Organization.
	Reasons []OrganizationDeactivationLogReason `json:"reasons"`

	// Details: more details about why the Organization is being closed.
	Details *string `json:"details,omitempty"`
}

// OrganizationAPICloseWarningRequest: organization api close warning request.
type OrganizationAPICloseWarningRequest struct {
	// WarningID: the ID of the warning.
	WarningID string `json:"-"`

	// Note: add a note on warning.
	Note string `json:"note"`
}

// OrganizationAPICreateWarningRequest: organization api create warning request.
type OrganizationAPICreateWarningRequest struct {
	// OrganizationID: the ID of the organization.
	OrganizationID string `json:"organization_id"`

	// Reason: the reason of the warning.
	// Default value: unknown_reason
	Reason WarningReason `json:"reason"`

	// Note: add a note on the warning.
	Note string `json:"note"`

	// IfNotExists: create the warning only if there is not already one with the same reason.
	IfNotExists bool `json:"if_not_exists"`

	// LockingNow: toggle if the warning is locking now.
	// Precisely one of LockingNow, LockedAt must be set.
	LockingNow *bool `json:"locking_now,omitempty"`

	// LockedAt: the lock date of the warning in the future.
	// Precisely one of LockingNow, LockedAt must be set.
	LockedAt *time.Time `json:"locked_at,omitempty"`
}

// OrganizationAPIGetActiveSupportPlanRequest: organization api get active support plan request.
type OrganizationAPIGetActiveSupportPlanRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`
}

// OrganizationAPIGetWarningLogRequest: organization api get warning log request.
type OrganizationAPIGetWarningLogRequest struct {
	WarningLogID string `json:"-"`
}

// OrganizationAPIGetWarningRequest: organization api get warning request.
type OrganizationAPIGetWarningRequest struct {
	// WarningID: the ID of the warning.
	WarningID string `json:"-"`
}

// OrganizationAPIListOrganizationDeactivationLogsRequest: organization api list organization deactivation logs request.
type OrganizationAPIListOrganizationDeactivationLogsRequest struct {
	// OrganizationID: filter by organization ID.
	OrganizationID string `json:"-"`

	// OrderBy: sort order of the organization deactivation logs.
	// Default value: created_at_asc
	OrderBy ListOrganizationDeactivationLogsRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equals to 1.
	Page *int32 `json:"-"`

	// PageSize: number of items per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`
}

// OrganizationAPIListSupportPlansRequest: organization api list support plans request.
type OrganizationAPIListSupportPlansRequest struct {
	// OrganizationID: filter by organization ID.
	OrganizationID string `json:"-"`

	// OrderBy: sort order of support plans.
	// Default value: started_at_asc
	OrderBy ListSupportPlansRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equals to 1.
	Page *int32 `json:"-"`

	// PageSize: number of items per page. Value must be between 1 and 1000.
	PageSize *uint32 `json:"-"`
}

// OrganizationAPIListWarningLogsRequest: organization api list warning logs request.
type OrganizationAPIListWarningLogsRequest struct {
	// OrderBy: how the warning logs are ordered in the response.
	// Default value: created_at_asc
	OrderBy ListWarningLogsRequestOrderBy `json:"-"`

	// Page: the page number for the returned warning logs.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of warning logs per page.
	PageSize *uint32 `json:"-"`

	// CreatedAfter: filter on created after a given timestamp.
	CreatedAfter *time.Time `json:"-"`

	// CreatedBefore: filter on created before a given timestamp.
	CreatedBefore *time.Time `json:"-"`

	// OrganizationWarningID: filter on an organization warning uuid.
	OrganizationWarningID *string `json:"-"`

	// FromPrincipalID: filter on an principal uuid.
	FromPrincipalID *string `json:"-"`

	// Action: filter on an action.
	// Default value: unknown_action
	Action WarningLogAction `json:"-"`

	// OrganizationID: filter on an organization uuid.
	OrganizationID *string `json:"-"`
}

// OrganizationAPIListWarningsRequest: organization api list warnings request.
type OrganizationAPIListWarningsRequest struct {
	// Page: the page number for the returned warnings.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of warnings per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the warnings are ordered in the response.
	// Default value: opened_at_asc
	OrderBy ListWarningsRequestOrderBy `json:"-"`

	// OrganizationID: filter on an organization uuid.
	OrganizationID *string `json:"-"`

	// Reason: filter on a warning reason.
	// Default value: unknown_reason
	Reason WarningReason `json:"-"`

	// IsClosed: filter on status closed or not.
	IsClosed *bool `json:"-"`

	// IsLocking: filter on status locking or not.
	IsLocking *bool `json:"-"`

	// UpdatedAfter: filter on updated after a given timestamp.
	UpdatedAfter *time.Time `json:"-"`

	// Reasons: filter on multiple reasons.
	Reasons []WarningReason `json:"-"`

	// LockedBefore: filter on locked before a given timestamp.
	LockedBefore *time.Time `json:"-"`

	// LockedAfter: filter on locked after a given timestamp.
	LockedAfter *time.Time `json:"-"`
}

// OrganizationAPIUpdateWarningRequest: organization api update warning request.
type OrganizationAPIUpdateWarningRequest struct {
	// WarningID: the ID of the warning.
	WarningID string `json:"-"`

	// Reason: edit the reason of the warning.
	// Default value: unknown_reason
	Reason WarningReason `json:"reason"`

	// Note: edit the note on warning.
	Note *string `json:"note,omitempty"`

	// Locking: toggle if the warning is locking.
	// Precisely one of Locking, LockingAt must be set.
	Locking *bool `json:"locking,omitempty"`

	// LockingAt: the lock date of the warning in the future.
	// Precisely one of Locking, LockingAt must be set.
	LockingAt *time.Time `json:"locking_at,omitempty"`
}

// ProjectAPIGetProjectRequest: project api get project request.
type ProjectAPIGetProjectRequest struct {
	// ProjectID: the project ID of the project.
	ProjectID string `json:"-"`

	// IncludeDeleted: include deleted projects in the response.
	IncludeDeleted bool `json:"-"`
}

// ProjectAPIListProjectsRequest: project api list projects request.
type ProjectAPIListProjectsRequest struct {
	// OrganizationID: return only projects that are part of this organization.
	OrganizationID *string `json:"-"`

	// UpdatedAfter: return only the projects updated after this time.
	UpdatedAfter *time.Time `json:"-"`

	// Name: return only the project with this name.
	Name *string `json:"-"`

	// Page: the page number for the returned projects.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of projects per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the projects are ordered in the response.
	// Default value: created_at_asc
	OrderBy ListProjectsRequestOrderBy `json:"-"`

	// IncludeDeleted: include deleted projects in the response.
	IncludeDeleted *bool `json:"-"`

	// ProjectIDs: filter out by a list of project ID.
	ProjectIDs []string `json:"-"`
}

// SetMailingListSubscriptionsResponse: set mailing list subscriptions response.
type SetMailingListSubscriptionsResponse struct {
	// MailingListSubscriptions: a list of mailing list subscriptions.
	MailingListSubscriptions []*MailingListSubscription `json:"mailing_list_subscriptions"`
}

// StopMailingListSubscriptionsResponse: stop mailing list subscriptions response.
type StopMailingListSubscriptionsResponse struct {
	// UnsubscriptionCount: the number of organizations which were actually unsubscribed.
	UnsubscriptionCount uint32 `json:"unsubscription_count"`
}

// UserAPIConfirmPhoneValidationRequest: user api confirm phone validation request.
type UserAPIConfirmPhoneValidationRequest struct {
	// PhoneValidationID: the ID of the phone validation.
	PhoneValidationID string `json:"-"`
}

// UserAPIDisableMFAOTPRequest: user api disable mfaotp request.
type UserAPIDisableMFAOTPRequest struct {
	// UserID: the ID of the user to disable the mfa for.
	UserID string `json:"-"`
}

// UserAPIGetAccountRequest: user api get account request.
type UserAPIGetAccountRequest struct {
	AccountID string `json:"-"`
}

// UserAPIGetMailingListRecipientsRequest: user api get mailing list recipients request.
type UserAPIGetMailingListRecipientsRequest struct {
	// OrganizationID: the ID of the organization.
	OrganizationID string `json:"-"`

	// MailingListID: the ID of the mailing list.
	MailingListID string `json:"-"`
}

// UserAPIListAccountsRequest: user api list accounts request.
type UserAPIListAccountsRequest struct {
	// Deleted: toggle to display deleted accounts or not.
	Deleted *bool `json:"-"`

	// Page: the page number.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of items per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the accounts are ordered in the response.
	// Default value: created_at_asc
	OrderBy ListAccountsRequestOrderBy `json:"-"`
}

// UserAPIListInvitationsRequest: user api list invitations request.
type UserAPIListInvitationsRequest struct {
	// Page: the page number.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of items per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the invitations are ordered in the response.
	// Default value: created_at_asc
	OrderBy ListInvitationsRequestOrderBy `json:"-"`

	// IncludeDeleted: toggle to display deleted invitations or not.
	IncludeDeleted bool `json:"-"`

	// UserID: the ID of the user.
	UserID *string `json:"-"`

	// Email: the email that received the invitations.
	Email *string `json:"-"`
}

// UserAPIListMailingListSubscriptionsRequest: user api list mailing list subscriptions request.
type UserAPIListMailingListSubscriptionsRequest struct {
	// UserID: the ID of the user.
	UserID string `json:"-"`

	// Page: the page number.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of items per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: the sort order.
	// Default value: name_asc
	OrderBy ListMailingListSubscriptionsRequestOrderBy `json:"-"`

	// OrganizationID: the ID of the organization.
	OrganizationID *string `json:"-"`
}

// UserAPIListPhoneValidationsRequest: user api list phone validations request.
type UserAPIListPhoneValidationsRequest struct {
	// UserID: the ID of the user.
	UserID string `json:"-"`

	// Page: the page number.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of items per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the phone validations are ordered in the response.
	// Default value: sent_at_asc
	OrderBy ListPhoneValidationsRequestOrderBy `json:"-"`
}

// UserAPISetMailingListSubscriptionsRequest: user api set mailing list subscriptions request.
type UserAPISetMailingListSubscriptionsRequest struct {
	// UserID: the ID of the user.
	UserID string `json:"-"`

	// MailingListSubscribed: a map of mailing list ids with their subscription status.
	MailingListSubscribed map[string]bool `json:"mailing_list_subscribed"`

	// OrganizationID: the ID of the organization.
	OrganizationID *string `json:"organization_id,omitempty"`
}

// UserAPIStopMailingListSubscriptionsRequest: user api stop mailing list subscriptions request.
type UserAPIStopMailingListSubscriptionsRequest struct {
	// MailingListID: the ID of the mailing list.
	MailingListID string `json:"mailing_list_id"`

	// Unsubscriptions: the list of organizations to unsubscribe.
	Unsubscriptions []*StopMailingListSubscriptionsRequestUnsubscription `json:"unsubscriptions"`
}

// Account Admin Blacklisted Domain API.
type BlacklistedDomainAPI struct {
	client *scw.Client
}

// NewBlacklistedDomainAPI returns a BlacklistedDomainAPI object from a Scaleway client.
func NewBlacklistedDomainAPI(client *scw.Client) *BlacklistedDomainAPI {
	return &BlacklistedDomainAPI{
		client: client,
	}
}

// CreateBlacklistedDomain: Create a blacklisted domain.
func (s *BlacklistedDomainAPI) CreateBlacklistedDomain(req *BlacklistedDomainAPICreateBlacklistedDomainRequest, opts ...scw.RequestOption) (*BlacklistedDomain, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v3/blacklisted-domains",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp BlacklistedDomain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBlacklistedDomain: Delete a blacklisted domain.
func (s *BlacklistedDomainAPI) DeleteBlacklistedDomain(req *BlacklistedDomainAPIDeleteBlacklistedDomainRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.BlacklistedDomainID) == "" {
		return errors.New("field BlacklistedDomainID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account-admin/v3/blacklisted-domains/" + fmt.Sprint(req.BlacklistedDomainID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListBlacklistedDomains: List blacklisted domains.
func (s *BlacklistedDomainAPI) ListBlacklistedDomains(req *BlacklistedDomainAPIListBlacklistedDomainsRequest, opts ...scw.RequestOption) (*ListBlacklistedDomainsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "block_account_creation", req.BlockAccountCreation)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/blacklisted-domains",
		Query:  query,
	}

	var resp ListBlacklistedDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// The Contract API allows you to manage contracts.
type ContractAPI struct {
	client *scw.Client
}

// NewContractAPI returns a ContractAPI object from a Scaleway client.
func NewContractAPI(client *scw.Client) *ContractAPI {
	return &ContractAPI{
		client: client,
	}
}

// ListContracts: List contracts.
func (s *ContractAPI) ListContracts(req *ContractAPIListContractsRequest, opts ...scw.RequestOption) (*ListContractsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "contract_type", req.ContractType)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/contracts",
		Query:  query,
	}

	var resp ListContractsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateContract: Create a new contract.
func (s *ContractAPI) CreateContract(req *ContractAPICreateContractRequest, opts ...scw.RequestOption) (*Contract, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v3/contracts",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Contract

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetContract: Get a contract from its ID.
func (s *ContractAPI) GetContract(req *ContractAPIGetContractRequest, opts ...scw.RequestOption) (*Contract, error) {
	var err error

	if fmt.Sprint(req.ContractID) == "" {
		return nil, errors.New("field ContractID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/contracts/" + fmt.Sprint(req.ContractID) + "",
	}

	var resp Contract

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DownloadContract: Download a contract content.
func (s *ContractAPI) DownloadContract(req *ContractAPIDownloadContractRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "locale", req.Locale)

	if fmt.Sprint(req.ContractID) == "" {
		return nil, errors.New("field ContractID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/contracts/" + fmt.Sprint(req.ContractID) + "/download",
		Query:  query,
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListContractSignatures: List contract signatures for an organization.
func (s *ContractAPI) ListContractSignatures(req *ContractAPIListContractSignaturesRequest, opts ...scw.RequestOption) (*ListContractSignaturesResponse, error) {
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/contract-signatures",
		Query:  query,
	}

	var resp ListContractSignaturesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// The Organization API allows you to manage the personal data of an organization.
type OrganizationAPI struct {
	client *scw.Client
}

// NewOrganizationAPI returns a OrganizationAPI object from a Scaleway client.
func NewOrganizationAPI(client *scw.Client) *OrganizationAPI {
	return &OrganizationAPI{
		client: client,
	}
}

// ListWarnings: List organization warnings.
func (s *OrganizationAPI) ListWarnings(req *OrganizationAPIListWarningsRequest, opts ...scw.RequestOption) (*ListWarningsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "reason", req.Reason)
	parameter.AddToQuery(query, "is_closed", req.IsClosed)
	parameter.AddToQuery(query, "is_locking", req.IsLocking)
	parameter.AddToQuery(query, "updated_after", req.UpdatedAfter)
	parameter.AddToQuery(query, "reasons", req.Reasons)
	parameter.AddToQuery(query, "locked_before", req.LockedBefore)
	parameter.AddToQuery(query, "locked_after", req.LockedAfter)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/warnings",
		Query:  query,
	}

	var resp ListWarningsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateWarning: Create an organization warning.
func (s *OrganizationAPI) CreateWarning(req *OrganizationAPICreateWarningRequest, opts ...scw.RequestOption) (*Warning, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v3/warnings",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Warning

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWarning: Get a specific organization warning.
func (s *OrganizationAPI) GetWarning(req *OrganizationAPIGetWarningRequest, opts ...scw.RequestOption) (*Warning, error) {
	var err error

	if fmt.Sprint(req.WarningID) == "" {
		return nil, errors.New("field WarningID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/warnings/" + fmt.Sprint(req.WarningID) + "",
	}

	var resp Warning

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateWarning: Update a specific organization warning.
func (s *OrganizationAPI) UpdateWarning(req *OrganizationAPIUpdateWarningRequest, opts ...scw.RequestOption) (*Warning, error) {
	var err error

	if fmt.Sprint(req.WarningID) == "" {
		return nil, errors.New("field WarningID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account-admin/v3/warnings/" + fmt.Sprint(req.WarningID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Warning

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CloseWarning: Close a specific organization warning.
func (s *OrganizationAPI) CloseWarning(req *OrganizationAPICloseWarningRequest, opts ...scw.RequestOption) (*Warning, error) {
	var err error

	if fmt.Sprint(req.WarningID) == "" {
		return nil, errors.New("field WarningID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v3/warnings/" + fmt.Sprint(req.WarningID) + "/close",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Warning

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CloseAllWarnings: Close all warnings opened with the specified reason in the organization.
func (s *OrganizationAPI) CloseAllWarnings(req *OrganizationAPICloseAllWarningsRequest, opts ...scw.RequestOption) (*CloseAllWarningsResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v3/warnings/close-all",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CloseAllWarningsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListWarningLogs: List organization warning logs.
func (s *OrganizationAPI) ListWarningLogs(req *OrganizationAPIListWarningLogsRequest, opts ...scw.RequestOption) (*ListWarningLogsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "created_after", req.CreatedAfter)
	parameter.AddToQuery(query, "created_before", req.CreatedBefore)
	parameter.AddToQuery(query, "organization_warning_id", req.OrganizationWarningID)
	parameter.AddToQuery(query, "from_principal_id", req.FromPrincipalID)
	parameter.AddToQuery(query, "action", req.Action)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/warning-logs",
		Query:  query,
	}

	var resp ListWarningLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetWarningLog: Get a specific organization warning log.
func (s *OrganizationAPI) GetWarningLog(req *OrganizationAPIGetWarningLogRequest, opts ...scw.RequestOption) (*WarningLog, error) {
	var err error

	if fmt.Sprint(req.WarningLogID) == "" {
		return nil, errors.New("field WarningLogID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/warning-logs/" + fmt.Sprint(req.WarningLogID) + "",
	}

	var resp WarningLog

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSupportPlan: Change the support plan by supplying the level of support wanted and the organization ID.
func (s *OrganizationAPI) ChangeSupportPlan(req *OrganizationAPIChangeSupportPlanRequest, opts ...scw.RequestOption) (*SupportPlan, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v3/organizations/" + fmt.Sprint(req.OrganizationID) + "/support-plans",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SupportPlan

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetActiveSupportPlan: Get the support plan for the organization along with the Technical Account Manager contact information.
func (s *OrganizationAPI) GetActiveSupportPlan(req *OrganizationAPIGetActiveSupportPlanRequest, opts ...scw.RequestOption) (*ActiveSupportPlan, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/organizations/" + fmt.Sprint(req.OrganizationID) + "/active-support-plan",
	}

	var resp ActiveSupportPlan

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSupportPlans: List the support plans for the organization.
func (s *OrganizationAPI) ListSupportPlans(req *OrganizationAPIListSupportPlansRequest, opts ...scw.RequestOption) (*ListSupportPlansResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/organizations/" + fmt.Sprint(req.OrganizationID) + "/support-plans",
		Query:  query,
	}

	var resp ListSupportPlansResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CloseOrganization: Close the Organization associated with the given ID.
func (s *OrganizationAPI) CloseOrganization(req *OrganizationAPICloseOrganizationRequest, opts ...scw.RequestOption) error {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v3/organizations/" + fmt.Sprint(req.OrganizationID) + "/close",
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

// ListOrganizationDeactivationLogs:
func (s *OrganizationAPI) ListOrganizationDeactivationLogs(req *OrganizationAPIListOrganizationDeactivationLogsRequest, opts ...scw.RequestOption) (*ListOrganizationDeactivationLogsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/organizations/" + fmt.Sprint(req.OrganizationID) + "/deactivation-logs",
		Query:  query,
	}

	var resp ListOrganizationDeactivationLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Account Admin Project API.
type ProjectAPI struct {
	client *scw.Client
}

// NewProjectAPI returns a ProjectAPI object from a Scaleway client.
func NewProjectAPI(client *scw.Client) *ProjectAPI {
	return &ProjectAPI{
		client: client,
	}
}

// ListProjects: List projects.
func (s *ProjectAPI) ListProjects(req *ProjectAPIListProjectsRequest, opts ...scw.RequestOption) (*ListProjectsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "updated_after", req.UpdatedAfter)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "include_deleted", req.IncludeDeleted)
	parameter.AddToQuery(query, "project_ids", req.ProjectIDs)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/projects",
		Query:  query,
	}

	var resp ListProjectsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetProject: Get project.
func (s *ProjectAPI) GetProject(req *ProjectAPIGetProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "include_deleted", req.IncludeDeleted)

	if fmt.Sprint(req.ProjectID) == "" {
		return nil, errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/projects/" + fmt.Sprint(req.ProjectID) + "",
		Query:  query,
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// The User API allows you to manage the personal data of a user.
type UserAPI struct {
	client *scw.Client
}

// NewUserAPI returns a UserAPI object from a Scaleway client.
func NewUserAPI(client *scw.Client) *UserAPI {
	return &UserAPI{
		client: client,
	}
}

// ListMailingListSubscriptions: List mailing list subscriptions.
func (s *UserAPI) ListMailingListSubscriptions(req *UserAPIListMailingListSubscriptionsRequest, opts ...scw.RequestOption) (*ListMailingListSubscriptionsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/users/" + fmt.Sprint(req.UserID) + "/mailing-list-subscriptions",
		Query:  query,
	}

	var resp ListMailingListSubscriptionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetMailingListSubscriptions: Set mailing list subscriptions.
func (s *UserAPI) SetMailingListSubscriptions(req *UserAPISetMailingListSubscriptionsRequest, opts ...scw.RequestOption) (*SetMailingListSubscriptionsResponse, error) {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/account-admin/v3/users/" + fmt.Sprint(req.UserID) + "/mailing-list-subscriptions",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetMailingListSubscriptionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMailingListRecipients: Get mailing list recipients.
func (s *UserAPI) GetMailingListRecipients(req *UserAPIGetMailingListRecipientsRequest, opts ...scw.RequestOption) (*GetMailingListRecipientsResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "mailing_list_id", req.MailingListID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/mailing-list-recipients",
		Query:  query,
	}

	var resp GetMailingListRecipientsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StopMailingListSubscriptions: Stop mailing list subscriptions.
func (s *UserAPI) StopMailingListSubscriptions(req *UserAPIStopMailingListSubscriptionsRequest, opts ...scw.RequestOption) (*StopMailingListSubscriptionsResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v3/stop-mailing-list-subscriptions",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp StopMailingListSubscriptionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPhoneValidations: List phone number validations.
func (s *UserAPI) ListPhoneValidations(req *UserAPIListPhoneValidationsRequest, opts ...scw.RequestOption) (*ListPhoneValidationsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "user_id", req.UserID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/phone-validations",
		Query:  query,
	}

	var resp ListPhoneValidationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ConfirmPhoneValidation: Confirm the phone number validation.
func (s *UserAPI) ConfirmPhoneValidation(req *UserAPIConfirmPhoneValidationRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.PhoneValidationID) == "" {
		return errors.New("field PhoneValidationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v3/phone-validations/" + fmt.Sprint(req.PhoneValidationID) + "/confirm",
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

// DisableMFAOTP: Disable MFA OTP.
func (s *UserAPI) DisableMFAOTP(req *UserAPIDisableMFAOTPRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v3/users/" + fmt.Sprint(req.UserID) + "/disable-mfa",
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

// ListAccounts: List account requests.
func (s *UserAPI) ListAccounts(req *UserAPIListAccountsRequest, opts ...scw.RequestOption) (*ListAccountsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "deleted", req.Deleted)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/accounts",
		Query:  query,
	}

	var resp ListAccountsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAccount: Get a single account request.
func (s *UserAPI) GetAccount(req *UserAPIGetAccountRequest, opts ...scw.RequestOption) (*Account, error) {
	var err error

	if fmt.Sprint(req.AccountID) == "" {
		return nil, errors.New("field AccountID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/accounts/" + fmt.Sprint(req.AccountID) + "",
	}

	var resp Account

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInvitations: List invitations.
func (s *UserAPI) ListInvitations(req *UserAPIListInvitationsRequest, opts ...scw.RequestOption) (*ListInvitationsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "include_deleted", req.IncludeDeleted)
	parameter.AddToQuery(query, "user_id", req.UserID)
	parameter.AddToQuery(query, "email", req.Email)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v3/invitations",
		Query:  query,
	}

	var resp ListInvitationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

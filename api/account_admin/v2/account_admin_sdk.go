// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package account_admin provides methods and message types of the account_admin v2 API.
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

type GDPRRequestStatus string

const (
	// Unknown status.
	GDPRRequestStatusUnknownStatus = GDPRRequestStatus("unknown_status")
	// Pending.
	GDPRRequestStatusPending = GDPRRequestStatus("pending")
	// Completed.
	GDPRRequestStatusCompleted = GDPRRequestStatus("completed")
	// Failed.
	GDPRRequestStatusFailed = GDPRRequestStatus("failed")
	// Processing.
	GDPRRequestStatusProcessing = GDPRRequestStatus("processing")
)

func (enum GDPRRequestStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum GDPRRequestStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *GDPRRequestStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = GDPRRequestStatus(GDPRRequestStatus(tmp).String())
	return nil
}

type GDPRRequestType string

const (
	// Unknown type.
	GDPRRequestTypeUnknownType = GDPRRequestType("unknown_type")
	// Data export.
	GDPRRequestTypeDataExport = GDPRRequestType("data_export")
	// Data information.
	GDPRRequestTypeDataInformation = GDPRRequestType("data_information")
	// Data limit.
	GDPRRequestTypeDataLimit = GDPRRequestType("data_limit")
	// Data portability.
	GDPRRequestTypeDataPortability = GDPRRequestType("data_portability")
	// Delete account.
	GDPRRequestTypeDeleteAccount = GDPRRequestType("delete_account")
)

func (enum GDPRRequestType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum GDPRRequestType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *GDPRRequestType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = GDPRRequestType(GDPRRequestType(tmp).String())
	return nil
}

type ListEligibleUsersRequestOrderBy string

const (
	// Creation date ascending.
	ListEligibleUsersRequestOrderByCreatedAtAsc = ListEligibleUsersRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListEligibleUsersRequestOrderByCreatedAtDesc = ListEligibleUsersRequestOrderBy("created_at_desc")
	// Lastname ascending.
	ListEligibleUsersRequestOrderByLastnameAsc = ListEligibleUsersRequestOrderBy("lastname_asc")
	// Lastname descending.
	ListEligibleUsersRequestOrderByLastnameDesc = ListEligibleUsersRequestOrderBy("lastname_desc")
)

func (enum ListEligibleUsersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListEligibleUsersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListEligibleUsersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListEligibleUsersRequestOrderBy(ListEligibleUsersRequestOrderBy(tmp).String())
	return nil
}

type ListGDPRRequestsRequestOrderBy string

const (
	// Type ascending.
	ListGDPRRequestsRequestOrderByTypeAsc = ListGDPRRequestsRequestOrderBy("type_asc")
	// Type descending.
	ListGDPRRequestsRequestOrderByTypeDesc = ListGDPRRequestsRequestOrderBy("type_desc")
	// Status ascending.
	ListGDPRRequestsRequestOrderByStatusAsc = ListGDPRRequestsRequestOrderBy("status_asc")
	// Status descending.
	ListGDPRRequestsRequestOrderByStatusDesc = ListGDPRRequestsRequestOrderBy("status_desc")
	// Creation date ascending.
	ListGDPRRequestsRequestOrderByCreatedAtAsc = ListGDPRRequestsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListGDPRRequestsRequestOrderByCreatedAtDesc = ListGDPRRequestsRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListGDPRRequestsRequestOrderByUpdatedAtAsc = ListGDPRRequestsRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListGDPRRequestsRequestOrderByUpdatedAtDesc = ListGDPRRequestsRequestOrderBy("updated_at_desc")
)

func (enum ListGDPRRequestsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "type_asc"
	}
	return string(enum)
}

func (enum ListGDPRRequestsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListGDPRRequestsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListGDPRRequestsRequestOrderBy(ListGDPRRequestsRequestOrderBy(tmp).String())
	return nil
}

type ListOrganizationTechnicalAccountManagersHistoryRequestOrderBy string

const (
	// Start date ascending.
	ListOrganizationTechnicalAccountManagersHistoryRequestOrderByStartDateAsc = ListOrganizationTechnicalAccountManagersHistoryRequestOrderBy("start_date_asc")
	// Start date descending.
	ListOrganizationTechnicalAccountManagersHistoryRequestOrderByStartDateDesc = ListOrganizationTechnicalAccountManagersHistoryRequestOrderBy("start_date_desc")
	// Stop date ascending.
	ListOrganizationTechnicalAccountManagersHistoryRequestOrderByStopDateAsc = ListOrganizationTechnicalAccountManagersHistoryRequestOrderBy("stop_date_asc")
	// Stop date descending.
	ListOrganizationTechnicalAccountManagersHistoryRequestOrderByStopDateDesc = ListOrganizationTechnicalAccountManagersHistoryRequestOrderBy("stop_date_desc")
)

func (enum ListOrganizationTechnicalAccountManagersHistoryRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "start_date_asc"
	}
	return string(enum)
}

func (enum ListOrganizationTechnicalAccountManagersHistoryRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOrganizationTechnicalAccountManagersHistoryRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOrganizationTechnicalAccountManagersHistoryRequestOrderBy(ListOrganizationTechnicalAccountManagersHistoryRequestOrderBy(tmp).String())
	return nil
}

type ListOrganizationsRequestOrderBy string

const (
	// Creation date ascending.
	ListOrganizationsRequestOrderByCreatedAtAsc = ListOrganizationsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListOrganizationsRequestOrderByCreatedAtDesc = ListOrganizationsRequestOrderBy("created_at_desc")
	// Contacted by active TAM date ascending.
	ListOrganizationsRequestOrderByContactedByActiveTamAtAsc = ListOrganizationsRequestOrderBy("contacted_by_active_tam_at_asc")
	// Contacted by active TAM date descending.
	ListOrganizationsRequestOrderByContactedByActiveTamAtDesc = ListOrganizationsRequestOrderBy("contacted_by_active_tam_at_desc")
)

func (enum ListOrganizationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListOrganizationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOrganizationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOrganizationsRequestOrderBy(ListOrganizationsRequestOrderBy(tmp).String())
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

type ListSSHKeysRequestOrderBy string

const (
	// Creation date ascending.
	ListSSHKeysRequestOrderByCreatedAtAsc = ListSSHKeysRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListSSHKeysRequestOrderByCreatedAtDesc = ListSSHKeysRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListSSHKeysRequestOrderByUpdatedAtAsc = ListSSHKeysRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListSSHKeysRequestOrderByUpdatedAtDesc = ListSSHKeysRequestOrderBy("updated_at_desc")
	// Name ascending.
	ListSSHKeysRequestOrderByNameAsc = ListSSHKeysRequestOrderBy("name_asc")
	// Name descending.
	ListSSHKeysRequestOrderByNameDesc = ListSSHKeysRequestOrderBy("name_desc")
)

func (enum ListSSHKeysRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSSHKeysRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSSHKeysRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSSHKeysRequestOrderBy(ListSSHKeysRequestOrderBy(tmp).String())
	return nil
}

type ListSupportSubscriptionsRequestOrderBy string

const (
	// Creation date ascending.
	ListSupportSubscriptionsRequestOrderByCreatedAtAsc = ListSupportSubscriptionsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListSupportSubscriptionsRequestOrderByCreatedAtDesc = ListSupportSubscriptionsRequestOrderBy("created_at_desc")
	// Start date ascending.
	ListSupportSubscriptionsRequestOrderByStartedAtAsc = ListSupportSubscriptionsRequestOrderBy("started_at_asc")
	// Start date descending.
	ListSupportSubscriptionsRequestOrderByStartedAtDesc = ListSupportSubscriptionsRequestOrderBy("started_at_desc")
	// Stop date ascending.
	ListSupportSubscriptionsRequestOrderByStoppedAtAsc = ListSupportSubscriptionsRequestOrderBy("stopped_at_asc")
	// Stop date descending.
	ListSupportSubscriptionsRequestOrderByStoppedAtDesc = ListSupportSubscriptionsRequestOrderBy("stopped_at_desc")
)

func (enum ListSupportSubscriptionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSupportSubscriptionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSupportSubscriptionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSupportSubscriptionsRequestOrderBy(ListSupportSubscriptionsRequestOrderBy(tmp).String())
	return nil
}

type ListTechnicalAccountManagersRequestOrderBy string

const (
	// Creation date ascending.
	ListTechnicalAccountManagersRequestOrderByCreatedAtAsc = ListTechnicalAccountManagersRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListTechnicalAccountManagersRequestOrderByCreatedAtDesc = ListTechnicalAccountManagersRequestOrderBy("created_at_desc")
	// Lastname ascending.
	ListTechnicalAccountManagersRequestOrderByLastnameAsc = ListTechnicalAccountManagersRequestOrderBy("lastname_asc")
	// Lastname descending.
	ListTechnicalAccountManagersRequestOrderByLastnameDesc = ListTechnicalAccountManagersRequestOrderBy("lastname_desc")
	// Number of assigned customers ascending.
	ListTechnicalAccountManagersRequestOrderByNbAssignedCustomersAsc = ListTechnicalAccountManagersRequestOrderBy("nb_assigned_customers_asc")
	// Number of assigned customers descending.
	ListTechnicalAccountManagersRequestOrderByNbAssignedCustomersDesc = ListTechnicalAccountManagersRequestOrderBy("nb_assigned_customers_desc")
)

func (enum ListTechnicalAccountManagersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListTechnicalAccountManagersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTechnicalAccountManagersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTechnicalAccountManagersRequestOrderBy(ListTechnicalAccountManagersRequestOrderBy(tmp).String())
	return nil
}

type ListTrustEventsRequestOrderBy string

const (
	// Creation date ascending.
	ListTrustEventsRequestOrderByCreatedAtAsc = ListTrustEventsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListTrustEventsRequestOrderByCreatedAtDesc = ListTrustEventsRequestOrderBy("created_at_desc")
)

func (enum ListTrustEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListTrustEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTrustEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTrustEventsRequestOrderBy(ListTrustEventsRequestOrderBy(tmp).String())
	return nil
}

type ListUserAuthenticationLogsRequestOrderBy string

const (
	// Creation date ascending.
	ListUserAuthenticationLogsRequestOrderByCreatedAtAsc = ListUserAuthenticationLogsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListUserAuthenticationLogsRequestOrderByCreatedAtDesc = ListUserAuthenticationLogsRequestOrderBy("created_at_desc")
)

func (enum ListUserAuthenticationLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListUserAuthenticationLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListUserAuthenticationLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListUserAuthenticationLogsRequestOrderBy(ListUserAuthenticationLogsRequestOrderBy(tmp).String())
	return nil
}

type OrganizationCorporateIndustry string

const (
	// Unknown corporate industry.
	OrganizationCorporateIndustryUnknownCorporateIndustry = OrganizationCorporateIndustry("unknown_corporate_industry")
	// Consulting & Services.
	OrganizationCorporateIndustryConsultingServices = OrganizationCorporateIndustry("consulting_services")
	// Cybersecurity & software.
	OrganizationCorporateIndustryCybersecuritySoftware = OrganizationCorporateIndustry("cybersecurity_software")
	// E-commerce & retail.
	OrganizationCorporateIndustryEcommerceRetail = OrganizationCorporateIndustry("ecommerce_retail")
	// Education.
	OrganizationCorporateIndustryEducation = OrganizationCorporateIndustry("education")
	// Energy.
	OrganizationCorporateIndustryEnergy = OrganizationCorporateIndustry("energy")
	// Financial Services & Insurance.
	OrganizationCorporateIndustryFinancialServicesInsurance = OrganizationCorporateIndustry("financial_services_insurance")
	// Gaming & Entertainment.
	OrganizationCorporateIndustryGamingEntertainment = OrganizationCorporateIndustry("gaming_entertainment")
	// Hospitality & Leisure.
	OrganizationCorporateIndustryHospitalityLeisure = OrganizationCorporateIndustry("hospitality_leisure")
	// Lifescience, Healthcare & Pharmaceuticals.
	OrganizationCorporateIndustryLifescienceHealthcarePharmaceuticals = OrganizationCorporateIndustry("lifescience_healthcare_pharmaceuticals")
	// Manufacturing.
	OrganizationCorporateIndustryManufacturing = OrganizationCorporateIndustry("manufacturing")
	// Media - Press & TV.
	OrganizationCorporateIndustryMediaPressTv = OrganizationCorporateIndustry("media_press_tv")
	// Public Sector.
	OrganizationCorporateIndustryPublicSector = OrganizationCorporateIndustry("public_sector")
	// Telecommunications.
	OrganizationCorporateIndustryTelecommunications = OrganizationCorporateIndustry("telecommunications")
	// Technology.
	OrganizationCorporateIndustryTechnology = OrganizationCorporateIndustry("technology")
	// Other.
	OrganizationCorporateIndustryOther = OrganizationCorporateIndustry("other")
)

func (enum OrganizationCorporateIndustry) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_corporate_industry"
	}
	return string(enum)
}

func (enum OrganizationCorporateIndustry) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationCorporateIndustry) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationCorporateIndustry(OrganizationCorporateIndustry(tmp).String())
	return nil
}

type OrganizationCustomerClass string

const (
	// Unknown customer class.
	OrganizationCustomerClassUnknownCustomerClass = OrganizationCustomerClass("unknown_customer_class")
	// Individual.
	OrganizationCustomerClassIndividual = OrganizationCustomerClass("individual")
	// Corporate.
	OrganizationCustomerClassCorporate = OrganizationCustomerClass("corporate")
	// Internal.
	OrganizationCustomerClassInternal = OrganizationCustomerClass("internal")
	// Intragroup.
	OrganizationCustomerClassIntragroup = OrganizationCustomerClass("intragroup")
)

func (enum OrganizationCustomerClass) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_customer_class"
	}
	return string(enum)
}

func (enum OrganizationCustomerClass) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationCustomerClass) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationCustomerClass(OrganizationCustomerClass(tmp).String())
	return nil
}

type OrganizationPaymentMode string

const (
	// Unknown payment mode.
	OrganizationPaymentModeUnknownPaymentMode = OrganizationPaymentMode("unknown_payment_mode")
	// Card.
	OrganizationPaymentModeCard = OrganizationPaymentMode("card")
	// SEPA.
	OrganizationPaymentModeSepa = OrganizationPaymentMode("sepa")
	// Wire transfer.
	OrganizationPaymentModeWireTransfer = OrganizationPaymentMode("wire_transfer")
)

func (enum OrganizationPaymentMode) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_payment_mode"
	}
	return string(enum)
}

func (enum OrganizationPaymentMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationPaymentMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationPaymentMode(OrganizationPaymentMode(tmp).String())
	return nil
}

type OrganizationWarningReason string

const (
	// Unknown reason.
	OrganizationWarningReasonUnknownReason = OrganizationWarningReason("unknown_reason")
	// Security concerns.
	OrganizationWarningReasonSecurityConcerns = OrganizationWarningReason("security_concerns")
	// Security issue.
	OrganizationWarningReasonSecurityIssue = OrganizationWarningReason("security_issue")
	// Critical security issue.
	OrganizationWarningReasonCriticalSecurityIssue = OrganizationWarningReason("critical_security_issue")
	// Locked for abuse.
	OrganizationWarningReasonLockedForAbuse = OrganizationWarningReason("locked_for_abuse")
	// Network abuse.
	OrganizationWarningReasonNetworkAbuse = OrganizationWarningReason("network_abuse")
	// Rate limiting.
	OrganizationWarningReasonRateLimiting = OrganizationWarningReason("rate_limiting")
	// Locked by reseller.
	OrganizationWarningReasonLockedByReseller = OrganizationWarningReason("locked_by_reseller")
	// Reseller locked.
	OrganizationWarningReasonResellerLocked = OrganizationWarningReason("reseller_locked")
	// Invoice payment failure.
	OrganizationWarningReasonInvoicePaymentFailure = OrganizationWarningReason("invoice_payment_failure")
	// Got chargeback.
	OrganizationWarningReasonGotChargeback = OrganizationWarningReason("got_chargeback")
	// Magic code rate limiting.
	OrganizationWarningReasonMagicCodeRateLimiting = OrganizationWarningReason("magic_code_rate_limiting")
	// Validate TOS.
	OrganizationWarningReasonValidateTos = OrganizationWarningReason("validate_tos")
	// Validate email.
	OrganizationWarningReasonValidateEmail = OrganizationWarningReason("validate_email")
	// Confirm email change.
	OrganizationWarningReasonConfirmEmailChange = OrganizationWarningReason("confirm_email_change")
	// Phone missing.
	OrganizationWarningReasonPhoneMissing = OrganizationWarningReason("phone_missing")
	// Payment info missing.
	OrganizationWarningReasonPaymentInfoMissing = OrganizationWarningReason("payment_info_missing")
	// Billing info missing.
	OrganizationWarningReasonBillingInfoMissing = OrganizationWarningReason("billing_info_missing")
	// Account closed.
	OrganizationWarningReasonAccountClosed = OrganizationWarningReason("account_closed")
	// Card expired or soon.
	OrganizationWarningReasonCardExpiredOrSoon = OrganizationWarningReason("card_expired_or_soon")
	// Magic code validation timeout.
	OrganizationWarningReasonMagicCodeValidationTimeout = OrganizationWarningReason("magic_code_validation_timeout")
	// Account anonymized.
	OrganizationWarningReasonAccountAnonymized = OrganizationWarningReason("account_anonymized")
	// Budget delete resources.
	OrganizationWarningReasonBudgetDeleteResources = OrganizationWarningReason("budget_delete_resources")
	// Budget forbid new resources.
	OrganizationWarningReasonBudgetForbidNewResources = OrganizationWarningReason("budget_forbid_new_resources")
	// Validate KYC.
	OrganizationWarningReasonValidateKyc = OrganizationWarningReason("validate_kyc")
	// Online invoice payment failure.
	OrganizationWarningReasonOnlineInvoicePaymentFailure = OrganizationWarningReason("online_invoice_payment_failure")
)

func (enum OrganizationWarningReason) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_reason"
	}
	return string(enum)
}

func (enum OrganizationWarningReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationWarningReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationWarningReason(OrganizationWarningReason(tmp).String())
	return nil
}

type RoleName string

const (
	// Access his bucket.
	RoleNameAccessHisBucket = RoleName("access_his_bucket")
	// Account workflow.
	RoleNameAccountWorkflow = RoleName("account_workflow")
	// Admin.
	RoleNameAdmin = RoleName("admin")
	// Admin server edit.
	RoleNameAdminServerEdit = RoleName("admin_server_edit")
	// Admin server reader.
	RoleNameAdminServerReader = RoleName("admin_server_reader")
	// Admin stop abuse.
	RoleNameAdminStopAbuse = RoleName("admin_stop_abuse")
	// Admin task creator.
	RoleNameAdminTaskCreator = RoleName("admin_task_creator")
	// Admin user reader.
	RoleNameAdminUserReader = RoleName("admin_user_reader")
	// AutoDNS.
	RoleNameAutodns = RoleName("autodns")
	// Billing.
	RoleNameBilling = RoleName("billing")
	// Delegated billing.
	RoleNameDelegatedBilling = RoleName("delegated_billing")
	// Delegated owner.
	RoleNameDelegatedOwner = RoleName("delegated_owner")
	// Gotty.
	RoleNameGotty = RoleName("gotty")
	// Mrproper.
	RoleNameMrproper = RoleName("mrproper")
	// Ops.
	RoleNameOps = RoleName("ops")
	// Owner.
	RoleNameOwner = RoleName("owner")
	// Project owner.
	RoleNameProjectOwner = RoleName("project_owner")
	// Scaleway devops.
	RoleNameScalewayDevops = RoleName("scaleway_devops")
	// Super admin.
	RoleNameSuperAdmin = RoleName("super_admin")
	// Super reader.
	RoleNameSuperReader = RoleName("super_reader")
	// Support L1.
	RoleNameSupportL1 = RoleName("support_l1")
	// Support L2.
	RoleNameSupportL2 = RoleName("support_l2")
	// Support manager.
	RoleNameSupportManager = RoleName("support_manager")
	// Trustsafety.
	RoleNameTrustsafety = RoleName("trustsafety")
	// Worker OS release.
	RoleNameWorkerOsRelease = RoleName("worker_os_release")
	// Hardware failure.
	RoleNameHardwareFailure = RoleName("hardware_failure")
	// Reseller.
	RoleNameReseller = RoleName("reseller")
)

func (enum RoleName) String() string {
	if enum == "" {
		// return default value if empty
		return "access_his_bucket"
	}
	return string(enum)
}

func (enum RoleName) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RoleName) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RoleName(RoleName(tmp).String())
	return nil
}

type ScalewayOrganizationCustomerClass string

const (
	// Unknown customer class.
	ScalewayOrganizationCustomerClassUnknownCustomerClass = ScalewayOrganizationCustomerClass("unknown_customer_class")
	// Individual.
	ScalewayOrganizationCustomerClassIndividual = ScalewayOrganizationCustomerClass("individual")
	// Corporate.
	ScalewayOrganizationCustomerClassCorporate = ScalewayOrganizationCustomerClass("corporate")
	// Internal.
	ScalewayOrganizationCustomerClassInternal = ScalewayOrganizationCustomerClass("internal")
	// Intragroup.
	ScalewayOrganizationCustomerClassIntragroup = ScalewayOrganizationCustomerClass("intragroup")
)

func (enum ScalewayOrganizationCustomerClass) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_customer_class"
	}
	return string(enum)
}

func (enum ScalewayOrganizationCustomerClass) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ScalewayOrganizationCustomerClass) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ScalewayOrganizationCustomerClass(ScalewayOrganizationCustomerClass(tmp).String())
	return nil
}

type SupportSubscriptionPlan string

const (
	// Unknown plan.
	SupportSubscriptionPlanUnknownPlan = SupportSubscriptionPlan("unknown_plan")
	// Basic.
	SupportSubscriptionPlanBasic = SupportSubscriptionPlan("basic")
	// Developer.
	SupportSubscriptionPlanDeveloper = SupportSubscriptionPlan("developer")
	// Business.
	SupportSubscriptionPlanBusiness = SupportSubscriptionPlan("business")
	// Enterprise.
	SupportSubscriptionPlanEnterprise = SupportSubscriptionPlan("enterprise")
	// Bronze.
	SupportSubscriptionPlanBronze = SupportSubscriptionPlan("bronze")
	// Silver.
	SupportSubscriptionPlanSilver = SupportSubscriptionPlan("silver")
	// Gold.
	SupportSubscriptionPlanGold = SupportSubscriptionPlan("gold")
	// Platinum.
	SupportSubscriptionPlanPlatinum = SupportSubscriptionPlan("platinum")
)

func (enum SupportSubscriptionPlan) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_plan"
	}
	return string(enum)
}

func (enum SupportSubscriptionPlan) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SupportSubscriptionPlan) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SupportSubscriptionPlan(SupportSubscriptionPlan(tmp).String())
	return nil
}

type SupportSubscriptionStatus string

const (
	// Unknown status.
	SupportSubscriptionStatusUnknownStatus = SupportSubscriptionStatus("unknown_status")
	// Expired.
	SupportSubscriptionStatusExpired = SupportSubscriptionStatus("expired")
	// Active.
	SupportSubscriptionStatusActive = SupportSubscriptionStatus("active")
	// Pending.
	SupportSubscriptionStatusPending = SupportSubscriptionStatus("pending")
)

func (enum SupportSubscriptionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum SupportSubscriptionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SupportSubscriptionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SupportSubscriptionStatus(SupportSubscriptionStatus(tmp).String())
	return nil
}

type TrustEventMischief string

const (
	// Unknown mischief.
	TrustEventMischiefUnknownMischief = TrustEventMischief("unknown_mischief")
	// Duplicate payment device.
	TrustEventMischiefDuplicatePaymentDevice = TrustEventMischief("duplicate_payment_device")
	// Invalid payment device.
	TrustEventMischiefInvalidPaymentDevice = TrustEventMischief("invalid_payment_device")
	// Duplicate SSH key.
	TrustEventMischiefDuplicateSSHKey = TrustEventMischief("duplicate_ssh_key")
)

func (enum TrustEventMischief) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_mischief"
	}
	return string(enum)
}

func (enum TrustEventMischief) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TrustEventMischief) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TrustEventMischief(TrustEventMischief(tmp).String())
	return nil
}

type UserAuthenticationLogFailureReason string

const (
	// Unknown failure reason.
	UserAuthenticationLogFailureReasonUnknownFailureReason = UserAuthenticationLogFailureReason("unknown_failure_reason")
	// Invalid password.
	UserAuthenticationLogFailureReasonInvalidPassword = UserAuthenticationLogFailureReason("invalid_password")
	// Invalid MFA.
	UserAuthenticationLogFailureReasonInvalidMfa = UserAuthenticationLogFailureReason("invalid_mfa")
	// Invalid captcha.
	UserAuthenticationLogFailureReasonInvalidCaptcha = UserAuthenticationLogFailureReason("invalid_captcha")
	// Invalid magic link.
	UserAuthenticationLogFailureReasonInvalidMagicLink = UserAuthenticationLogFailureReason("invalid_magic_link")
)

func (enum UserAuthenticationLogFailureReason) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_failure_reason"
	}
	return string(enum)
}

func (enum UserAuthenticationLogFailureReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UserAuthenticationLogFailureReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UserAuthenticationLogFailureReason(UserAuthenticationLogFailureReason(tmp).String())
	return nil
}

type UserAuthenticationLogMethod string

const (
	// Unknown method.
	UserAuthenticationLogMethodUnknownMethod = UserAuthenticationLogMethod("unknown_method")
	// Password.
	UserAuthenticationLogMethodPassword = UserAuthenticationLogMethod("password")
	// Magic link.
	UserAuthenticationLogMethodMagicLink = UserAuthenticationLogMethod("magic_link")
	// JWT.
	UserAuthenticationLogMethodJwt = UserAuthenticationLogMethod("jwt")
)

func (enum UserAuthenticationLogMethod) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_method"
	}
	return string(enum)
}

func (enum UserAuthenticationLogMethod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UserAuthenticationLogMethod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UserAuthenticationLogMethod(UserAuthenticationLogMethod(tmp).String())
	return nil
}

type UserAuthenticationLogOrigin string

const (
	// Unknown origin.
	UserAuthenticationLogOriginUnknownOrigin = UserAuthenticationLogOrigin("unknown_origin")
	// Public API.
	UserAuthenticationLogOriginPublicAPI = UserAuthenticationLogOrigin("public_api")
	// Admin API.
	UserAuthenticationLogOriginAdminAPI = UserAuthenticationLogOrigin("admin_api")
)

func (enum UserAuthenticationLogOrigin) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_origin"
	}
	return string(enum)
}

func (enum UserAuthenticationLogOrigin) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UserAuthenticationLogOrigin) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UserAuthenticationLogOrigin(UserAuthenticationLogOrigin(tmp).String())
	return nil
}

type UserType string

const (
	// Unknown type.
	UserTypeUnknownType = UserType("unknown_type")
	// Basic.
	UserTypeBasic = UserType("basic")
	// Internal.
	UserTypeInternal = UserType("internal")
)

func (enum UserType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum UserType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UserType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UserType(UserType(tmp).String())
	return nil
}

// OrganizationSummary: organization summary.
type OrganizationSummary struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

// Role: role.
type Role struct {
	// Name: default value: access_his_bucket
	Name RoleName `json:"name"`

	Organization *OrganizationSummary `json:"organization"`
}

// Address: address.
type Address struct {
	Street string `json:"street"`

	Complementary string `json:"complementary"`

	City string `json:"city"`

	PostalCode string `json:"postal_code"`

	Country string `json:"country"`
}

// ScalewayUser: scaleway user.
type ScalewayUser struct {
	ID string `json:"id"`

	Firstname string `json:"firstname"`

	Lastname string `json:"lastname"`

	Email string `json:"email"`

	PhoneNumber string `json:"phone_number"`
}

// OrganizationDediboxAccount: organization dedibox account.
type OrganizationDediboxAccount struct {
	ID uint32 `json:"id"`

	LinkedAt *time.Time `json:"linked_at"`
}

// OrganizationWarning: organization warning.
type OrganizationWarning struct {
	ID string `json:"id"`

	OrganizationID string `json:"organization_id"`

	// Reason: default value: unknown_reason
	Reason OrganizationWarningReason `json:"reason"`

	OpenedAt *time.Time `json:"opened_at"`

	LockedAt *time.Time `json:"locked_at"`

	ClosedAt *time.Time `json:"closed_at"`

	ClosableByUser bool `json:"closable_by_user"`

	ClosedBy *string `json:"closed_by"`

	Code *string `json:"code"`

	Note *string `json:"note"`
}

// User: user.
type User struct {
	ID string `json:"id"`

	Email string `json:"email"`

	Firstname string `json:"firstname"`

	Lastname string `json:"lastname"`

	// Locale: default value: unknown_language_code
	Locale std.LanguageCode `json:"locale"`

	Roles []*Role `json:"roles"`

	// Type: default value: unknown_type
	Type UserType `json:"type"`

	Organizations []*OrganizationSummary `json:"organizations"`

	PhoneNumber string `json:"phone_number"`
}

// UserSummary: user summary.
type UserSummary struct {
	ID string `json:"id"`

	Email string `json:"email"`

	Firstname string `json:"firstname"`

	Lastname string `json:"lastname"`

	PhoneNumber string `json:"phone_number"`
}

// SSHKeyCreationInfo: ssh key creation info.
type SSHKeyCreationInfo struct {
	Address string `json:"address"`

	UserAgent string `json:"user_agent"`

	// CountryCode: default value: unknown_country_code
	CountryCode std.CountryCode `json:"country_code"`
}

// InternalUser: internal user.
type InternalUser struct {
	// AccountRootUserID: the account root user ID.
	AccountRootUserID string `json:"account_root_user_id"`

	// Firstname: the first name of user.
	Firstname string `json:"firstname"`

	// Lastname: the last name of user.
	Lastname string `json:"lastname"`

	// CreatedAt: the creation date.
	CreatedAt *time.Time `json:"created_at"`
}

// GDPRRequest: gdpr request.
type GDPRRequest struct {
	// ID: the ID of request.
	ID string `json:"id"`

	// AccountRootUserID: the account root user ID.
	AccountRootUserID string `json:"account_root_user_id"`

	// Type: type of request.
	// Default value: unknown_type
	Type GDPRRequestType `json:"type"`

	// Status: status of request.
	// Default value: unknown_status
	Status GDPRRequestStatus `json:"status"`

	// Comment: optional comment.
	Comment *string `json:"comment"`

	// ErrorMessage: the error message.
	ErrorMessage *string `json:"error_message"`

	// ResultMessage: the result message.
	ResultMessage *string `json:"result_message"`

	// CreatedAt: the creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: the update date.
	UpdatedAt *time.Time `json:"updated_at"`

	// DeletedAt: the deletion date.
	DeletedAt *time.Time `json:"deleted_at"`

	// ScheduledAt: the schedule date.
	ScheduledAt *time.Time `json:"scheduled_at"`
}

// ListInactiveUsersResponseInactiveUser: list inactive users response inactive user.
type ListInactiveUsersResponseInactiveUser struct {
	// ID: the inactive user ID.
	ID string `json:"id"`

	// OrganizationID: the ID of the organization created by the user.
	OrganizationID string `json:"organization_id"`
}

// TechnicalAccountManagerAssignment: technical account manager assignment.
type TechnicalAccountManagerAssignment struct {
	// AccountRootUserID: the account root user ID.
	AccountRootUserID string `json:"account_root_user_id"`

	// Firstname: the first name of user.
	Firstname string `json:"firstname"`

	// Lastname: the last name of user.
	Lastname string `json:"lastname"`

	// Email: the email of user.
	Email string `json:"email"`

	// StartDate: the start date.
	StartDate *time.Time `json:"start_date"`

	// StopDate: the end date if defined.
	StopDate *time.Time `json:"stop_date"`
}

// ScalewayOrganization: scaleway organization.
type ScalewayOrganization struct {
	ID string `json:"id"`

	Name string `json:"name"`

	// CustomerClass: default value: unknown_customer_class
	CustomerClass ScalewayOrganizationCustomerClass `json:"customer_class"`

	Address *Address `json:"address"`

	VatNumber *string `json:"vat_number"`

	BillingEmail *string `json:"billing_email"`

	Creator *ScalewayUser `json:"creator"`
}

// OrganizationLockingStatus: organization locking status.
type OrganizationLockingStatus struct {
	ID string `json:"id"`

	Locked bool `json:"locked"`

	NoQuota bool `json:"no_quota"`
}

// Organization: organization.
type Organization struct {
	ID string `json:"id"`

	Name string `json:"name"`

	SupportID string `json:"support_id"`

	SupportPin string `json:"support_pin"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	AddressLine1 string `json:"address_line1"`

	AddressLine2 string `json:"address_line2"`

	ZipCode string `json:"zip_code"`

	City string `json:"city"`

	// CountryCode: default value: unknown_country_code
	CountryCode std.CountryCode `json:"country_code"`

	SubdivisionCode string `json:"subdivision_code"`

	Creator *User `json:"creator"`

	CustomerLevels []string `json:"customer_levels"`

	// SupportLevel: default value: unknown_plan
	SupportLevel SupportSubscriptionPlan `json:"support_level"`

	TechnicalAccountManager *UserSummary `json:"technical_account_manager"`

	ResellerID *string `json:"reseller_id"`

	DediboxAccount *OrganizationDediboxAccount `json:"dedibox_account"`

	// CustomerClass: default value: unknown_customer_class
	CustomerClass OrganizationCustomerClass `json:"customer_class"`

	MfaEnforced bool `json:"mfa_enforced"`

	Score string `json:"score"`

	Whitelisted bool `json:"whitelisted"`

	Comment string `json:"comment"`

	Billable bool `json:"billable"`

	IsInvoiceable bool `json:"is_invoiceable"`

	Warnings []*OrganizationWarning `json:"warnings"`

	ContactedByActiveTamAt *time.Time `json:"contacted_by_active_tam_at"`

	ActiveTamNote string `json:"active_tam_note"`

	// PaymentMode: default value: unknown_payment_mode
	PaymentMode OrganizationPaymentMode `json:"payment_mode"`

	Origin string `json:"origin"`

	// CorporateIndustry: default value: unknown_corporate_industry
	CorporateIndustry OrganizationCorporateIndustry `json:"corporate_industry"`

	BillingEmail string `json:"billing_email"`

	SirenNumber string `json:"siren_number"`

	VatNumber string `json:"vat_number"`
}

// PhoneValidation: phone validation.
type PhoneValidation struct {
	ID string `json:"id"`

	PhoneNumber string `json:"phone_number"`

	Code string `json:"code"`

	CarrierType string `json:"carrier_type"`

	SentAt *time.Time `json:"sent_at"`

	ValidatedAt *time.Time `json:"validated_at"`
}

// Project: project.
type Project struct {
	ID string `json:"id"`

	Name string `json:"name"`

	OrganizationID string `json:"organization_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	Description string `json:"description"`
}

// SSHKey: ssh key.
type SSHKey struct {
	ID string `json:"id"`

	Name string `json:"name"`

	PublicKey string `json:"public_key"`

	Fingerprint string `json:"fingerprint"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	CreationInfo *SSHKeyCreationInfo `json:"creation_info"`

	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`
}

// SupportSubscription: support subscription.
type SupportSubscription struct {
	SupportSubscriptionID string `json:"support_subscription_id"`

	// Plan: default value: unknown_plan
	Plan SupportSubscriptionPlan `json:"plan"`

	StartDate *time.Time `json:"start_date"`

	StopDate *time.Time `json:"stop_date"`

	OrganizationID string `json:"organization_id"`

	// Status: default value: unknown_status
	Status SupportSubscriptionStatus `json:"status"`

	Price *scw.Money `json:"price"`
}

// TechnicalAccountManager: technical account manager.
type TechnicalAccountManager struct {
	// AccountRootUserID: the account root user ID.
	AccountRootUserID string `json:"account_root_user_id"`

	// Firstname: the first name of user.
	Firstname string `json:"firstname"`

	// Lastname: the last name of user.
	Lastname string `json:"lastname"`

	// Email: the email of user.
	Email string `json:"email"`

	// PhoneNumber: the phone number of user.
	PhoneNumber string `json:"phone_number"`

	// NbAssignedCustomers: the number of customers assigned to this technical account manager.
	NbAssignedCustomers uint32 `json:"nb_assigned_customers"`

	// CreatedAt: the creation date.
	CreatedAt *time.Time `json:"created_at"`

	// DeletedAt: the removal date. Empty if still active.
	DeletedAt *time.Time `json:"deleted_at"`
}

// TrustEvent: trust event.
type TrustEvent struct {
	Key string `json:"key"`

	OrganizationID string `json:"organization_id"`

	// Mischief: default value: unknown_mischief
	Mischief TrustEventMischief `json:"mischief"`

	StructureType string `json:"structure_type"`

	StructureID *string `json:"structure_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	Details *string `json:"details"`

	JSON *scw.JSONObject `json:"json"`
}

// UserAuthenticationLog: user authentication log.
type UserAuthenticationLog struct {
	// ID: the authentication log ID.
	ID string `json:"id"`

	// UserID: the user ID.
	UserID string `json:"user_id"`

	// CreatedAt: the date of the authentication.
	CreatedAt *time.Time `json:"created_at"`

	// IP: the IP address of the user.
	IP net.IP `json:"ip"`

	// Port: the port of the user.
	Port uint32 `json:"port"`

	// Country: the country of the user.
	Country *string `json:"country"`

	// UserAgent: the user agent of the user.
	UserAgent string `json:"user_agent"`

	// Origin: the origin of the request (public or admin console).
	// Default value: unknown_origin
	Origin UserAuthenticationLogOrigin `json:"origin"`

	// Method: the authentication method used (password or magic link).
	// Default value: unknown_method
	Method UserAuthenticationLogMethod `json:"method"`

	// MfaEnabled: whether MFA was enabled for the authentication.
	MfaEnabled bool `json:"mfa_enabled"`

	// Success: the authentication success status.
	Success bool `json:"success"`

	// FailureReason: the authentication failure reason.
	// Default value: unknown_failure_reason
	FailureReason UserAuthenticationLogFailureReason `json:"failure_reason"`
}

// JWT: jwt.
type JWT struct {
	// Jti: jTI of the JWT.
	Jti string `json:"jti"`

	// IssuerID: iAM user ID of the JWT issuer.
	IssuerID string `json:"issuer_id"`

	// AudienceID: iAM user ID of the JWT audience.
	AudienceID string `json:"audience_id"`

	// CreatedAt: creation date of the JWT.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: renewal date of the JWT.
	UpdatedAt *time.Time `json:"updated_at"`

	// ExpiresAt: expiration date of the JWT.
	ExpiresAt *time.Time `json:"expires_at"`

	// IP: source IP address for the JWT creation.
	IP net.IP `json:"ip"`

	// UserAgent: user agent for the JWT creation.
	UserAgent string `json:"user_agent"`
}

// ChangeSupportPlanRequest: change support plan request.
type ChangeSupportPlanRequest struct {
	OrganizationID string `json:"-"`

	// Plan: default value: unknown_plan
	Plan SupportSubscriptionPlan `json:"plan"`

	Price *scw.Money `json:"price,omitempty"`

	DisableEmail bool `json:"disable_email"`
}

// CreateInternalOrganizationRequest: create internal organization request.
type CreateInternalOrganizationRequest struct {
	Email string `json:"email"`

	Firstname string `json:"firstname"`

	Lastname string `json:"lastname"`
}

// CreateTrustEventRequest: create trust event request.
type CreateTrustEventRequest struct {
	OrganizationID string `json:"organization_id"`

	// Mischief: default value: unknown_mischief
	Mischief TrustEventMischief `json:"mischief"`

	StructureType string `json:"structure_type"`

	StructureID string `json:"structure_id"`

	Details string `json:"details"`

	JSON *scw.JSONObject `json:"json,omitempty"`
}

// DediboxToElementsAPIListOrganizationsLinkedToDediboxRequest: dedibox to elements api list organizations linked to dedibox request.
type DediboxToElementsAPIListOrganizationsLinkedToDediboxRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	UpdatedAfter *time.Time `json:"-"`
}

// DeleteInternalOrganizationRequest: delete internal organization request.
type DeleteInternalOrganizationRequest struct {
	OrganizationID string `json:"-"`
}

// DeleteOrganizationRequest: delete organization request.
type DeleteOrganizationRequest struct {
	OrganizationID string `json:"-"`
}

// DeleteUserRequest: delete user request.
type DeleteUserRequest struct {
	UserID string `json:"-"`

	OrganizationID string `json:"-"`
}

// GDPRApiCreateGDPRRequestRequest: gdpr api create gdpr request request.
type GDPRApiCreateGDPRRequestRequest struct {
	// AccountRootUserID: the ID of account user.
	AccountRootUserID string `json:"account_root_user_id"`

	// Comment: comment for GDPR request.
	Comment *string `json:"comment,omitempty"`

	// Type: type of request.
	// Default value: unknown_type
	Type GDPRRequestType `json:"type"`

	// DisableDeletionDelay: disable the default deletion delay of 72h.
	DisableDeletionDelay bool `json:"disable_deletion_delay"`
}

// GDPRApiDeleteGDPRRequestRequest: gdpr api delete gdpr request request.
type GDPRApiDeleteGDPRRequestRequest struct {
	// RequestID: ID of GDPR request.
	RequestID string `json:"-"`
}

// GDPRApiGetGDPRRequestRequest: gdpr api get gdpr request request.
type GDPRApiGetGDPRRequestRequest struct {
	// RequestID: ID of GDPR request.
	RequestID string `json:"-"`
}

// GDPRApiListGDPRRequestsRequest: gdpr api list gdpr requests request.
type GDPRApiListGDPRRequestsRequest struct {
	// OrderBy: order criteria.
	// Default value: type_asc
	OrderBy ListGDPRRequestsRequestOrderBy `json:"-"`

	// Page: the page number.
	Page *int32 `json:"-"`

	// PageSize: the size of page for pagination.
	PageSize *uint32 `json:"-"`

	// AccountRootUserID: the ID of account user.
	AccountRootUserID *string `json:"-"`

	// Status: status of request.
	// Default value: unknown_status
	Status GDPRRequestStatus `json:"-"`

	// Type: type of request.
	// Default value: unknown_type
	Type GDPRRequestType `json:"-"`
}

// GDPRApiRetryGDPRRequestRequest: gdpr api retry gdpr request request.
type GDPRApiRetryGDPRRequestRequest struct {
	// RequestID: ID of GDPR request.
	RequestID string `json:"-"`

	// ScheduledAt: optional schedule date.
	ScheduledAt *time.Time `json:"scheduled_at,omitempty"`
}

// GetOrganizationRequest: get organization request.
type GetOrganizationRequest struct {
	OrganizationID string `json:"-"`
}

// GetProjectRequest: get project request.
type GetProjectRequest struct {
	// ProjectID: the project ID of the project.
	ProjectID string `json:"-"`
}

// GetUserRequest: get user request.
type GetUserRequest struct {
	UserID string `json:"-"`
}

// ListEligibleUsersResponse: list eligible users response.
type ListEligibleUsersResponse struct {
	// TotalCount: the total number of eligible users.
	TotalCount uint32 `json:"total_count"`

	// InternalUsers: the list of eligible users.
	InternalUsers []*InternalUser `json:"internal_users"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListEligibleUsersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListEligibleUsersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListEligibleUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.InternalUsers = append(r.InternalUsers, results.InternalUsers...)
	r.TotalCount += uint32(len(results.InternalUsers))
	return uint32(len(results.InternalUsers)), nil
}

// ListGDPRRequestsResponse: list gdpr requests response.
type ListGDPRRequestsResponse struct {
	// GdprRequests: list of requests.
	GdprRequests []*GDPRRequest `json:"gdpr_requests"`

	// TotalCount: total number of requests.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListGDPRRequestsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListGDPRRequestsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListGDPRRequestsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.GdprRequests = append(r.GdprRequests, results.GdprRequests...)
	r.TotalCount += uint32(len(results.GdprRequests))
	return uint32(len(results.GdprRequests)), nil
}

// ListInactiveUsersRequest: list inactive users request.
type ListInactiveUsersRequest struct {
	// Page: the page number for the returned users.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of users per page.
	PageSize *uint32 `json:"-"`
}

// ListInactiveUsersResponse: list inactive users response.
type ListInactiveUsersResponse struct {
	// TotalCount: the total number of inactive users.
	TotalCount uint32 `json:"total_count"`

	// InactiveUsers: the list of inactive users.
	InactiveUsers []*ListInactiveUsersResponseInactiveUser `json:"inactive_users"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInactiveUsersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInactiveUsersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListInactiveUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.InactiveUsers = append(r.InactiveUsers, results.InactiveUsers...)
	r.TotalCount += uint32(len(results.InactiveUsers))
	return uint32(len(results.InactiveUsers)), nil
}

// ListOrganizationTechnicalAccountManagersHistoryResponse: list organization technical account managers history response.
type ListOrganizationTechnicalAccountManagersHistoryResponse struct {
	// TotalCount: the total number of TAM assignments.
	TotalCount uint32 `json:"total_count"`

	// TechnicalAccountManagersHistory: the list of TAM assignments.
	TechnicalAccountManagersHistory []*TechnicalAccountManagerAssignment `json:"technical_account_managers_history"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOrganizationTechnicalAccountManagersHistoryResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOrganizationTechnicalAccountManagersHistoryResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListOrganizationTechnicalAccountManagersHistoryResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.TechnicalAccountManagersHistory = append(r.TechnicalAccountManagersHistory, results.TechnicalAccountManagersHistory...)
	r.TotalCount += uint32(len(results.TechnicalAccountManagersHistory))
	return uint32(len(results.TechnicalAccountManagersHistory)), nil
}

// ListOrganizationsLinkedToDediboxResponse: list organizations linked to dedibox response.
type ListOrganizationsLinkedToDediboxResponse struct {
	TotalCount uint32 `json:"total_count"`

	Organizations []*ScalewayOrganization `json:"organizations"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOrganizationsLinkedToDediboxResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOrganizationsLinkedToDediboxResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListOrganizationsLinkedToDediboxResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Organizations = append(r.Organizations, results.Organizations...)
	r.TotalCount += uint32(len(results.Organizations))
	return uint32(len(results.Organizations)), nil
}

// ListOrganizationsLockingStatusRequest: list organizations locking status request.
type ListOrganizationsLockingStatusRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	UpdatedAfter *time.Time `json:"-"`
}

// ListOrganizationsLockingStatusResponse: list organizations locking status response.
type ListOrganizationsLockingStatusResponse struct {
	TotalCount uint32 `json:"total_count"`

	OrganizationsLockingStatus []*OrganizationLockingStatus `json:"organizations_locking_status"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOrganizationsLockingStatusResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOrganizationsLockingStatusResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListOrganizationsLockingStatusResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.OrganizationsLockingStatus = append(r.OrganizationsLockingStatus, results.OrganizationsLockingStatus...)
	r.TotalCount += uint32(len(results.OrganizationsLockingStatus))
	return uint32(len(results.OrganizationsLockingStatus)), nil
}

// ListOrganizationsRequest: list organizations request.
type ListOrganizationsRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	SupportID *string `json:"-"`

	TechnicalAccountManagerID *string `json:"-"`

	// Deprecated: SupportLevel: default value: unknown_plan
	SupportLevel *SupportSubscriptionPlan `json:"-"`

	ResellerID *string `json:"-"`

	CustomerLevel *string `json:"-"`

	OrganizationIDs []string `json:"-"`

	UpdatedAfter *time.Time `json:"-"`

	UpdatedBefore *time.Time `json:"-"`

	Name *string `json:"-"`

	// PaymentMode: default value: unknown_payment_mode
	PaymentMode OrganizationPaymentMode `json:"-"`

	CountryCode *string `json:"-"`

	TrustScore *string `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListOrganizationsRequestOrderBy `json:"-"`

	SupportLevels []SupportSubscriptionPlan `json:"-"`

	// CorporateIndustry: default value: unknown_corporate_industry
	CorporateIndustry OrganizationCorporateIndustry `json:"-"`

	// CustomerClass: default value: unknown_customer_class
	CustomerClass OrganizationCustomerClass `json:"-"`
}

// ListOrganizationsResponse: list organizations response.
type ListOrganizationsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Organizations []*Organization `json:"organizations"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOrganizationsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOrganizationsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListOrganizationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Organizations = append(r.Organizations, results.Organizations...)
	r.TotalCount += uint32(len(results.Organizations))
	return uint32(len(results.Organizations)), nil
}

// ListPhoneValidationsRequest: list phone validations request.
type ListPhoneValidationsRequest struct {
	UserID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: sent_at_asc
	OrderBy ListPhoneValidationsRequestOrderBy `json:"-"`
}

// ListPhoneValidationsResponse: list phone validations response.
type ListPhoneValidationsResponse struct {
	TotalCount uint32 `json:"total_count"`

	PhoneValidations []*PhoneValidation `json:"phone_validations"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPhoneValidationsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPhoneValidationsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPhoneValidationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PhoneValidations = append(r.PhoneValidations, results.PhoneValidations...)
	r.TotalCount += uint32(len(results.PhoneValidations))
	return uint32(len(results.PhoneValidations)), nil
}

// ListProjectsRequest: list projects request.
type ListProjectsRequest struct {
	// UpdatedAfter: return only the projects updated after this time.
	UpdatedAfter *time.Time `json:"-"`

	// OrderBy: how the projects are ordered in the response.
	// Default value: created_at_asc
	OrderBy ListProjectsRequestOrderBy `json:"-"`

	// Page: the page number for the returned projects.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of projects per page.
	PageSize *uint32 `json:"-"`

	// Name: return only the project with this name.
	Name *string `json:"-"`

	// OrganizationID: return only projects that are part of this organization.
	OrganizationID *string `json:"-"`

	// IncludeDeleted: include deleted projects in the response.
	IncludeDeleted *bool `json:"-"`
}

// ListProjectsResponse: list projects response.
type ListProjectsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Projects []*Project `json:"projects"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListProjectsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListProjectsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListProjectsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Projects = append(r.Projects, results.Projects...)
	r.TotalCount += uint32(len(results.Projects))
	return uint32(len(results.Projects)), nil
}

// ListSSHKeysRequest: list ssh keys request.
type ListSSHKeysRequest struct {
	// OrderBy: how are SSH keys ordered in the response.
	// Default value: created_at_asc
	OrderBy ListSSHKeysRequestOrderBy `json:"-"`

	// Page: the page number for the returned SSH keys.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of SSH keys per page.
	PageSize *uint32 `json:"-"`

	// Name: return only the SSH keys that with this name.
	Name *string `json:"-"`

	// OrganizationID: return only SSH keys that are part of this organization.
	OrganizationID *string `json:"-"`

	// ProjectID: return only SSH keys that are part of this project.
	ProjectID *string `json:"-"`
}

// ListSSHKeysResponse: list ssh keys response.
type ListSSHKeysResponse struct {
	SSHKeys []*SSHKey `json:"ssh_keys"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSSHKeysResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSSHKeysResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSSHKeysResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SSHKeys = append(r.SSHKeys, results.SSHKeys...)
	r.TotalCount += uint32(len(results.SSHKeys))
	return uint32(len(results.SSHKeys)), nil
}

// ListSupportSubscriptionsRequest: list support subscriptions request.
type ListSupportSubscriptionsRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListSupportSubscriptionsRequestOrderBy `json:"-"`

	// Plan: default value: unknown_plan
	Plan SupportSubscriptionPlan `json:"-"`

	// Status: default value: unknown_status
	Status SupportSubscriptionStatus `json:"-"`
}

// ListSupportSubscriptionsResponse: list support subscriptions response.
type ListSupportSubscriptionsResponse struct {
	TotalCount uint32 `json:"total_count"`

	SupportSubscriptions []*SupportSubscription `json:"support_subscriptions"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSupportSubscriptionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSupportSubscriptionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSupportSubscriptionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SupportSubscriptions = append(r.SupportSubscriptions, results.SupportSubscriptions...)
	r.TotalCount += uint32(len(results.SupportSubscriptions))
	return uint32(len(results.SupportSubscriptions)), nil
}

// ListTechnicalAccountManagersResponse: list technical account managers response.
type ListTechnicalAccountManagersResponse struct {
	// TotalCount: the total number of TAM users.
	TotalCount uint32 `json:"total_count"`

	// TechnicalAccountManagers: the list of TAM users.
	TechnicalAccountManagers []*TechnicalAccountManager `json:"technical_account_managers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTechnicalAccountManagersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTechnicalAccountManagersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTechnicalAccountManagersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.TechnicalAccountManagers = append(r.TechnicalAccountManagers, results.TechnicalAccountManagers...)
	r.TotalCount += uint32(len(results.TechnicalAccountManagers))
	return uint32(len(results.TechnicalAccountManagers)), nil
}

// ListTrustEventsRequest: list trust events request.
type ListTrustEventsRequest struct {
	// OrderBy: default value: created_at_asc
	OrderBy ListTrustEventsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	// Mischief: default value: unknown_mischief
	Mischief TrustEventMischief `json:"-"`

	StructureType *string `json:"-"`

	StructureID *string `json:"-"`
}

// ListTrustEventsResponse: list trust events response.
type ListTrustEventsResponse struct {
	TotalCount uint32 `json:"total_count"`

	TrustEvents []*TrustEvent `json:"trust_events"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTrustEventsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTrustEventsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTrustEventsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.TrustEvents = append(r.TrustEvents, results.TrustEvents...)
	r.TotalCount += uint32(len(results.TrustEvents))
	return uint32(len(results.TrustEvents)), nil
}

// ListUserAuthenticationLogsRequest: list user authentication logs request.
type ListUserAuthenticationLogsRequest struct {
	// UserID: the user ID.
	UserID string `json:"-"`

	// Page: the page number for the returned authentication logs.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of authentication logs per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the authentication logs are ordered in the response.
	// Default value: created_at_asc
	OrderBy ListUserAuthenticationLogsRequestOrderBy `json:"-"`

	// Origin: return only the authentication logs with this origin.
	// Default value: unknown_origin
	Origin UserAuthenticationLogOrigin `json:"-"`

	// Method: return only the authentication logs with this method.
	// Default value: unknown_method
	Method UserAuthenticationLogMethod `json:"-"`

	// MfaEnabled: return only the authentication logs with this MFA status.
	MfaEnabled *bool `json:"-"`

	// Success: return only the authentication logs with this success status.
	Success *bool `json:"-"`

	// FailureReason: return only the authentication logs with this failure reason.
	// Default value: unknown_failure_reason
	FailureReason UserAuthenticationLogFailureReason `json:"-"`
}

// ListUserAuthenticationLogsResponse: list user authentication logs response.
type ListUserAuthenticationLogsResponse struct {
	// TotalCount: the total number of authentication logs.
	TotalCount uint32 `json:"total_count"`

	// UserAuthenticationLogs: the list of user authentication logs.
	UserAuthenticationLogs []*UserAuthenticationLog `json:"user_authentication_logs"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListUserAuthenticationLogsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListUserAuthenticationLogsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListUserAuthenticationLogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.UserAuthenticationLogs = append(r.UserAuthenticationLogs, results.UserAuthenticationLogs...)
	r.TotalCount += uint32(len(results.UserAuthenticationLogs))
	return uint32(len(results.UserAuthenticationLogs)), nil
}

// ListUsersRequest: list users request.
type ListUsersRequest struct {
	// Page: the page number for the returned users.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of users per page.
	PageSize *uint32 `json:"-"`

	// Type: return only the organizations of that type.
	// Default value: unknown_type
	Type UserType `json:"-"`

	// PhoneNumber: return only the organizations that have this phone numnber.
	PhoneNumber *string `json:"-"`

	// IsTechnicalAccountManager: return only users that can be technical account managers.
	IsTechnicalAccountManager bool `json:"-"`

	// OrganizationID: return only users that are part of this organization.
	OrganizationID *string `json:"-"`

	// UpdatedAfter: return only users that have been updated after this date.
	UpdatedAfter *time.Time `json:"-"`

	// UpdatedBefore: return only users that have been updated before this date.
	UpdatedBefore *time.Time `json:"-"`
}

// ListUsersResponse: list users response.
type ListUsersResponse struct {
	TotalCount uint32 `json:"total_count"`

	Users []*User `json:"users"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Users = append(r.Users, results.Users...)
	r.TotalCount += uint32(len(results.Users))
	return uint32(len(results.Users)), nil
}

// LoginResponse: login response.
type LoginResponse struct {
	// Jwt: data of the JWT created by the login request.
	Jwt *JWT `json:"jwt"`

	// Token: encoded JWT token.
	Token string `json:"token"`

	// RenewToken: encoded JWT renew token.
	RenewToken string `json:"renew_token"`
}

// SendTOSEmailRequest: send tos email request.
type SendTOSEmailRequest struct {
	OrganizationID string `json:"-"`

	Email *string `json:"email,omitempty"`
}

// TechnicalAccountManagerAPIAssignTechnicalAccountManagerRequest: technical account manager api assign technical account manager request.
type TechnicalAccountManagerAPIAssignTechnicalAccountManagerRequest struct {
	// OrganizationID: the organization ID.
	OrganizationID string `json:"-"`

	// AccountRootUserID: the account root user ID.
	AccountRootUserID *string `json:"account_root_user_id,omitempty"`

	// StartDate: the start date if defined (NOW by default).
	StartDate *time.Time `json:"start_date,omitempty"`
}

// TechnicalAccountManagerAPICreateTechnicalAccountManagerRequest: technical account manager api create technical account manager request.
type TechnicalAccountManagerAPICreateTechnicalAccountManagerRequest struct {
	// AccountRootUserID: the ID of user to add to the technical account manager team.
	AccountRootUserID string `json:"account_root_user_id"`
}

// TechnicalAccountManagerAPIDeleteTechnicalAccountManagerRequest: technical account manager api delete technical account manager request.
type TechnicalAccountManagerAPIDeleteTechnicalAccountManagerRequest struct {
	// AccountRootUserID: the ID of user to remove from the technical account manager team.
	AccountRootUserID string `json:"-"`
}

// TechnicalAccountManagerAPIListEligibleUsersRequest: technical account manager api list eligible users request.
type TechnicalAccountManagerAPIListEligibleUsersRequest struct {
	// OrderBy: how are users ordered in the response.
	// Default value: created_at_asc
	OrderBy ListEligibleUsersRequestOrderBy `json:"-"`

	// Page: the page number for the returned users.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of users per page (1000 by default).
	PageSize *uint32 `json:"-"`

	// Name: search by user name, usually concatenation of firstname and lastname.
	Name *string `json:"-"`
}

// TechnicalAccountManagerAPIListOrganizationTechnicalAccountManagersHistoryRequest: technical account manager api list organization technical account managers history request.
type TechnicalAccountManagerAPIListOrganizationTechnicalAccountManagersHistoryRequest struct {
	// OrganizationID: the ID of organization.
	OrganizationID string `json:"-"`

	// OrderBy: how are users ordered in the response.
	// Default value: start_date_asc
	OrderBy ListOrganizationTechnicalAccountManagersHistoryRequestOrderBy `json:"-"`

	// Page: the page number for the returned TAM users.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of TAM users per page (1000 by default).
	PageSize *uint32 `json:"-"`
}

// TechnicalAccountManagerAPIListTechnicalAccountManagersRequest: technical account manager api list technical account managers request.
type TechnicalAccountManagerAPIListTechnicalAccountManagersRequest struct {
	// OrderBy: how are users ordered in the response.
	// Default value: created_at_asc
	OrderBy ListTechnicalAccountManagersRequestOrderBy `json:"-"`

	// Page: the page number for the returned TAM users.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of TAM users per page (1000 by default).
	PageSize *uint32 `json:"-"`

	// AccountRootUserID: the ID of user to select.
	AccountRootUserID *string `json:"-"`

	// IncludeDeleted: true to select also removed TAM users.
	IncludeDeleted *bool `json:"-"`

	// Name: search by user name, usually concatenation of firstname and lastname.
	Name *string `json:"-"`
}

// TechnicalAccountManagerAPIUnassignTechnicalAccountManagerRequest: technical account manager api unassign technical account manager request.
type TechnicalAccountManagerAPIUnassignTechnicalAccountManagerRequest struct {
	// OrganizationID: the organization ID.
	OrganizationID string `json:"-"`

	// StopDate: the stop date if defined (NOW by default).
	StopDate *time.Time `json:"-"`
}

// UnauthenticatedUserAPILoginRequest: unauthenticated user api login request.
type UnauthenticatedUserAPILoginRequest struct {
	// Email: email address to log in with.
	Email string `json:"email"`

	// Password: password to log in with.
	// Precisely one of Password, Passwordless must be set.
	Password *string `json:"password,omitempty"`

	// Passwordless: passwordless token to log in with.
	// Precisely one of Password, Passwordless must be set.
	Passwordless *string `json:"passwordless,omitempty"`

	// Otp: oTP to log in with.
	// Precisely one of Otp must be set.
	Otp *string `json:"otp,omitempty"`

	// Captcha: captcha to log in with.
	Captcha *string `json:"captcha,omitempty"`
}

// UnsetUserPasswordRequest: unset user password request.
type UnsetUserPasswordRequest struct {
	// UserID: the user ID.
	UserID string `json:"-"`
}

// UpdateOrganizationRequest: update organization request.
type UpdateOrganizationRequest struct {
	// OrganizationID: the organization ID.
	OrganizationID string `json:"-"`

	// CustomerLevels: update the customer level.
	CustomerLevels *[]string `json:"customer_levels,omitempty"`

	// TechnicalAccountManagerID: attach or detach a technical account manager (TAM) to an organization.
	// Only organizations that are subscribed to a `silver` or `gold` support plans can have a technical account manager.
	TechnicalAccountManagerID *string `json:"technical_account_manager_id,omitempty"`

	// DediboxUserID: attach or Detach a Scaleway organization to/from a Dedibox account.
	// If set to `0` it will detach the organization from the Dedibox account.
	DediboxUserID *uint32 `json:"dedibox_user_id,omitempty"`

	// MfaEnforced: enforce MFA for the organization.
	// If set to `true` it will enforce MFA for the organization and all its users.
	// If set to `false` it will disable MFA requirement for the organization and all its users.
	MfaEnforced *bool `json:"mfa_enforced,omitempty"`

	// ContactedByActiveTamAt: organization with a subscription to GOLD or PLATINUM support plan have a Technical Account Manager assigned.
	// The TAM usually have to contact the owner of organization. This field allows to keep the date of contact.
	ContactedByActiveTamAt *time.Time `json:"contacted_by_active_tam_at,omitempty"`

	// ActiveTamNote: note of the last contact.
	ActiveTamNote *string `json:"active_tam_note,omitempty"`
}

// ValidatePhoneValidationRequest: validate phone validation request.
type ValidatePhoneValidationRequest struct {
	PhoneValidationID string `json:"-"`
}

// Account Admin API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// GetOrganization: Get organization.
func (s *API) GetOrganization(req *GetOrganizationRequest, opts ...scw.RequestOption) (*Organization, error) {
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
		Path:   "/account-admin/v2/organizations/" + fmt.Sprint(req.OrganizationID) + "",
	}

	var resp Organization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOrganization: Update organization.
func (s *API) UpdateOrganization(req *UpdateOrganizationRequest, opts ...scw.RequestOption) (*Organization, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account-admin/v2/organizations/" + fmt.Sprint(req.OrganizationID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Organization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOrganizations: List organization.
func (s *API) ListOrganizations(req *ListOrganizationsRequest, opts ...scw.RequestOption) (*ListOrganizationsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "support_id", req.SupportID)
	parameter.AddToQuery(query, "technical_account_manager_id", req.TechnicalAccountManagerID)
	parameter.AddToQuery(query, "support_level", req.SupportLevel)
	parameter.AddToQuery(query, "reseller_id", req.ResellerID)
	parameter.AddToQuery(query, "customer_level", req.CustomerLevel)
	parameter.AddToQuery(query, "organization_ids", req.OrganizationIDs)
	parameter.AddToQuery(query, "updated_after", req.UpdatedAfter)
	parameter.AddToQuery(query, "updated_before", req.UpdatedBefore)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "payment_mode", req.PaymentMode)
	parameter.AddToQuery(query, "country_code", req.CountryCode)
	parameter.AddToQuery(query, "trust_score", req.TrustScore)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "support_levels", req.SupportLevels)
	parameter.AddToQuery(query, "corporate_industry", req.CorporateIndustry)
	parameter.AddToQuery(query, "customer_class", req.CustomerClass)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/organizations",
		Query:  query,
	}

	var resp ListOrganizationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOrganizationsLockingStatus: List organizations locking status.
func (s *API) ListOrganizationsLockingStatus(req *ListOrganizationsLockingStatusRequest, opts ...scw.RequestOption) (*ListOrganizationsLockingStatusResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "updated_after", req.UpdatedAfter)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/organizations-locking-status",
		Query:  query,
	}

	var resp ListOrganizationsLockingStatusResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteOrganization: Delete organization.
func (s *API) DeleteOrganization(req *DeleteOrganizationRequest, opts ...scw.RequestOption) error {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account-admin/v2/organizations/" + fmt.Sprint(req.OrganizationID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: ListProjects: Deprecated in favor of Account Admin Project API v3. List projects.
func (s *API) ListProjects(req *ListProjectsRequest, opts ...scw.RequestOption) (*ListProjectsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "updated_after", req.UpdatedAfter)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "include_deleted", req.IncludeDeleted)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/projects",
		Query:  query,
	}

	var resp ListProjectsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: GetProject: Deprecated in favor of Account Admin Project API v3. Get project.
func (s *API) GetProject(req *GetProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return nil, errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTrustEvents:
func (s *API) ListTrustEvents(req *ListTrustEventsRequest, opts ...scw.RequestOption) (*ListTrustEventsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "mischief", req.Mischief)
	parameter.AddToQuery(query, "structure_type", req.StructureType)
	parameter.AddToQuery(query, "structure_id", req.StructureID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/trust-events",
		Query:  query,
	}

	var resp ListTrustEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTrustEvent:
func (s *API) CreateTrustEvent(req *CreateTrustEventRequest, opts ...scw.RequestOption) (*TrustEvent, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v2/trust-events",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TrustEvent

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSupportSubscriptions: List support subscriptions.
//
// Note that ended support plan will also be listed by default, use filters to change that.
func (s *API) ListSupportSubscriptions(req *ListSupportSubscriptionsRequest, opts ...scw.RequestOption) (*ListSupportSubscriptionsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "plan", req.Plan)
	parameter.AddToQuery(query, "status", req.Status)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/support-subscriptions",
		Query:  query,
	}

	var resp ListSupportSubscriptionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeSupportPlan: Change the support plan of an organization.
// All organizations comes with a default basic support plan.
//
// Note that a support plan upgrade starts immediately and a support plan downgrade starts the next month.
func (s *API) ChangeSupportPlan(req *ChangeSupportPlanRequest, opts ...scw.RequestOption) (*SupportSubscription, error) {
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
		Path:   "/account-admin/v2/organizations/" + fmt.Sprint(req.OrganizationID) + "/support-plan",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SupportSubscription

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListUsers: List all types of Scaleway users.
func (s *API) ListUsers(req *ListUsersRequest, opts ...scw.RequestOption) (*ListUsersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "phone_number", req.PhoneNumber)
	parameter.AddToQuery(query, "is_technical_account_manager", req.IsTechnicalAccountManager)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "updated_after", req.UpdatedAfter)
	parameter.AddToQuery(query, "updated_before", req.UpdatedBefore)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/users",
		Query:  query,
	}

	var resp ListUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUser:
func (s *API) GetUser(req *GetUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/users/" + fmt.Sprint(req.UserID) + "",
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteUser:
func (s *API) DeleteUser(req *DeleteUserRequest, opts ...scw.RequestOption) error {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.UserID) == "" {
		return errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account-admin/v2/users/" + fmt.Sprint(req.UserID) + "",
		Query:  query,
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UnsetUserPassword: Unset user password.
func (s *API) UnsetUserPassword(req *UnsetUserPasswordRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v2/users/" + fmt.Sprint(req.UserID) + "/unset-password",
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

// CreateInternalOrganization:
func (s *API) CreateInternalOrganization(req *CreateInternalOrganizationRequest, opts ...scw.RequestOption) (*Organization, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v2/internal-organizations",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Organization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteInternalOrganization:
func (s *API) DeleteInternalOrganization(req *DeleteInternalOrganizationRequest, opts ...scw.RequestOption) error {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account-admin/v2/internal-organizations/" + fmt.Sprint(req.OrganizationID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListSSHKeys: List all ssh-keys.
func (s *API) ListSSHKeys(req *ListSSHKeysRequest, opts ...scw.RequestOption) (*ListSSHKeysResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/ssh-keys",
		Query:  query,
	}

	var resp ListSSHKeysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendTOSEmail:
func (s *API) SendTOSEmail(req *SendTOSEmailRequest, opts ...scw.RequestOption) error {
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
		Path:   "/account-admin/v2/organizations/" + fmt.Sprint(req.OrganizationID) + "/send-tos-email",
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

// ListPhoneValidations:
func (s *API) ListPhoneValidations(req *ListPhoneValidationsRequest, opts ...scw.RequestOption) (*ListPhoneValidationsResponse, error) {
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
		Path:   "/account-admin/v2/phone-validations",
		Query:  query,
	}

	var resp ListPhoneValidationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidatePhoneValidation:
func (s *API) ValidatePhoneValidation(req *ValidatePhoneValidationRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.PhoneValidationID) == "" {
		return errors.New("field PhoneValidationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v2/phone-validations/" + fmt.Sprint(req.PhoneValidationID) + "/validate",
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

// ListInactiveUsers: List inactive users.
func (s *API) ListInactiveUsers(req *ListInactiveUsersRequest, opts ...scw.RequestOption) (*ListInactiveUsersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/inactive-users",
		Query:  query,
	}

	var resp ListInactiveUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListUserAuthenticationLogs: List user authentication logs.
func (s *API) ListUserAuthenticationLogs(req *ListUserAuthenticationLogsRequest, opts ...scw.RequestOption) (*ListUserAuthenticationLogsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "origin", req.Origin)
	parameter.AddToQuery(query, "method", req.Method)
	parameter.AddToQuery(query, "mfa_enabled", req.MfaEnabled)
	parameter.AddToQuery(query, "success", req.Success)
	parameter.AddToQuery(query, "failure_reason", req.FailureReason)

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/users/" + fmt.Sprint(req.UserID) + "/authentication-logs",
		Query:  query,
	}

	var resp ListUserAuthenticationLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DediboxToElementsAPI struct {
	client *scw.Client
}

// NewDediboxToElementsAPI returns a DediboxToElementsAPI object from a Scaleway client.
func NewDediboxToElementsAPI(client *scw.Client) *DediboxToElementsAPI {
	return &DediboxToElementsAPI{
		client: client,
	}
}

// ListOrganizationsLinkedToDedibox:
func (s *DediboxToElementsAPI) ListOrganizationsLinkedToDedibox(req *DediboxToElementsAPIListOrganizationsLinkedToDediboxRequest, opts ...scw.RequestOption) (*ListOrganizationsLinkedToDediboxResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "updated_after", req.UpdatedAfter)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/dedibox-to-elements/organizations",
		Query:  query,
	}

	var resp ListOrganizationsLinkedToDediboxResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GDPRAPI struct {
	client *scw.Client
}

// NewGDPRAPI returns a GDPRAPI object from a Scaleway client.
func NewGDPRAPI(client *scw.Client) *GDPRAPI {
	return &GDPRAPI{
		client: client,
	}
}

// CreateGDPRRequest: Create a new GDPR request.
func (s *GDPRAPI) CreateGDPRRequest(req *GDPRApiCreateGDPRRequestRequest, opts ...scw.RequestOption) (*GDPRRequest, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v2/gdpr-requests",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp GDPRRequest

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListGDPRRequests: List all GDPR requests.
func (s *GDPRAPI) ListGDPRRequests(req *GDPRApiListGDPRRequestsRequest, opts ...scw.RequestOption) (*ListGDPRRequestsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "account_root_user_id", req.AccountRootUserID)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "type", req.Type)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/gdpr-requests",
		Query:  query,
	}

	var resp ListGDPRRequestsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGDPRRequest: Get a GDPR request by ID.
func (s *GDPRAPI) GetGDPRRequest(req *GDPRApiGetGDPRRequestRequest, opts ...scw.RequestOption) (*GDPRRequest, error) {
	var err error

	if fmt.Sprint(req.RequestID) == "" {
		return nil, errors.New("field RequestID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/gdpr-requests/" + fmt.Sprint(req.RequestID) + "",
	}

	var resp GDPRRequest

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RetryGDPRRequest: Retry a GDPR request by ID.
func (s *GDPRAPI) RetryGDPRRequest(req *GDPRApiRetryGDPRRequestRequest, opts ...scw.RequestOption) (*GDPRRequest, error) {
	var err error

	if fmt.Sprint(req.RequestID) == "" {
		return nil, errors.New("field RequestID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v2/gdpr-requests/" + fmt.Sprint(req.RequestID) + "/retry",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp GDPRRequest

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteGDPRRequest: Delete a GDPR request by ID.
func (s *GDPRAPI) DeleteGDPRRequest(req *GDPRApiDeleteGDPRRequestRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.RequestID) == "" {
		return errors.New("field RequestID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account-admin/v2/gdpr-requests/" + fmt.Sprint(req.RequestID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// TAM Admin API.
type TechnicalAccountManagerAPI struct {
	client *scw.Client
}

// NewTechnicalAccountManagerAPI returns a TechnicalAccountManagerAPI object from a Scaleway client.
func NewTechnicalAccountManagerAPI(client *scw.Client) *TechnicalAccountManagerAPI {
	return &TechnicalAccountManagerAPI{
		client: client,
	}
}

// ListTechnicalAccountManagers: List users of the technical account manager team.
func (s *TechnicalAccountManagerAPI) ListTechnicalAccountManagers(req *TechnicalAccountManagerAPIListTechnicalAccountManagersRequest, opts ...scw.RequestOption) (*ListTechnicalAccountManagersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "account_root_user_id", req.AccountRootUserID)
	parameter.AddToQuery(query, "include_deleted", req.IncludeDeleted)
	parameter.AddToQuery(query, "name", req.Name)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/technical-account-managers",
		Query:  query,
	}

	var resp ListTechnicalAccountManagersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTechnicalAccountManager: Add an user to the technical account manager team.
func (s *TechnicalAccountManagerAPI) CreateTechnicalAccountManager(req *TechnicalAccountManagerAPICreateTechnicalAccountManagerRequest, opts ...scw.RequestOption) (*TechnicalAccountManager, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v2/technical-account-managers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TechnicalAccountManager

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteTechnicalAccountManager: Remove an user from the technical account manager team.
func (s *TechnicalAccountManagerAPI) DeleteTechnicalAccountManager(req *TechnicalAccountManagerAPIDeleteTechnicalAccountManagerRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.AccountRootUserID) == "" {
		return errors.New("field AccountRootUserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account-admin/v2/technical-account-managers/" + fmt.Sprint(req.AccountRootUserID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListOrganizationTechnicalAccountManagersHistory: List the history of TAM assignments for an organization.
func (s *TechnicalAccountManagerAPI) ListOrganizationTechnicalAccountManagersHistory(req *TechnicalAccountManagerAPIListOrganizationTechnicalAccountManagersHistoryRequest, opts ...scw.RequestOption) (*ListOrganizationTechnicalAccountManagersHistoryResponse, error) {
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
		Path:   "/account-admin/v2/technical-account-managers/history/" + fmt.Sprint(req.OrganizationID) + "",
		Query:  query,
	}

	var resp ListOrganizationTechnicalAccountManagersHistoryResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AssignTechnicalAccountManager: Assign a technical account manager to an organization.
func (s *TechnicalAccountManagerAPI) AssignTechnicalAccountManager(req *TechnicalAccountManagerAPIAssignTechnicalAccountManagerRequest, opts ...scw.RequestOption) (*TechnicalAccountManagerAssignment, error) {
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
		Path:   "/account-admin/v2/technical-account-managers/assignment/" + fmt.Sprint(req.OrganizationID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TechnicalAccountManagerAssignment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnassignTechnicalAccountManager: Unassign a technical account manager from an organization.
func (s *TechnicalAccountManagerAPI) UnassignTechnicalAccountManager(req *TechnicalAccountManagerAPIUnassignTechnicalAccountManagerRequest, opts ...scw.RequestOption) (*TechnicalAccountManagerAssignment, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "stop_date", req.StopDate)

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account-admin/v2/technical-account-managers/assignment/" + fmt.Sprint(req.OrganizationID) + "",
		Query:  query,
	}

	var resp TechnicalAccountManagerAssignment

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListEligibleUsers: List internal users without active TAM assignation.
func (s *TechnicalAccountManagerAPI) ListEligibleUsers(req *TechnicalAccountManagerAPIListEligibleUsersRequest, opts ...scw.RequestOption) (*ListEligibleUsersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2/technical-account-managers/available-users",
		Query:  query,
	}

	var resp ListEligibleUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows to manage a user without authentication.
type UnauthenticatedUserAPI struct {
	client *scw.Client
}

// NewUnauthenticatedUserAPI returns a UnauthenticatedUserAPI object from a Scaleway client.
func NewUnauthenticatedUserAPI(client *scw.Client) *UnauthenticatedUserAPI {
	return &UnauthenticatedUserAPI{
		client: client,
	}
}

// Login: Login user.
func (s *UnauthenticatedUserAPI) Login(req *UnauthenticatedUserAPILoginRequest, opts ...scw.RequestOption) (*LoginResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v2/login",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LoginResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

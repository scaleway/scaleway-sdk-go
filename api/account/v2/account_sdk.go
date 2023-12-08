// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package account provides methods and message types of the account v2 API.
package account

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
	std "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/std"
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

type CaptchaProviderName string

const (
	// Unknown name.
	CaptchaProviderNameUnknownName = CaptchaProviderName("unknown_name")
	// ReCAPTCHA v2.
	CaptchaProviderNameRecaptchaV2 = CaptchaProviderName("recaptcha_v2")
	// Friendly Captcha.
	CaptchaProviderNameFriendlyCaptcha = CaptchaProviderName("friendly_captcha")
	// HCaptcha.
	CaptchaProviderNameHcaptcha = CaptchaProviderName("hcaptcha")
)

func (enum CaptchaProviderName) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_name"
	}
	return string(enum)
}

func (enum CaptchaProviderName) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CaptchaProviderName) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CaptchaProviderName(CaptchaProviderName(tmp).String())
	return nil
}

type CreateUserRequestProfessionalOrganizationCorporateIndustry string

const (
	// Unknown corporate industry.
	CreateUserRequestProfessionalOrganizationCorporateIndustryUnknownCorporateIndustry = CreateUserRequestProfessionalOrganizationCorporateIndustry("unknown_corporate_industry")
	// Consulting & Services.
	CreateUserRequestProfessionalOrganizationCorporateIndustryConsultingServices = CreateUserRequestProfessionalOrganizationCorporateIndustry("consulting_services")
	// Cybersecurity & software.
	CreateUserRequestProfessionalOrganizationCorporateIndustryCybersecuritySoftware = CreateUserRequestProfessionalOrganizationCorporateIndustry("cybersecurity_software")
	// E-commerce & retail.
	CreateUserRequestProfessionalOrganizationCorporateIndustryEcommerceRetail = CreateUserRequestProfessionalOrganizationCorporateIndustry("ecommerce_retail")
	// Education.
	CreateUserRequestProfessionalOrganizationCorporateIndustryEducation = CreateUserRequestProfessionalOrganizationCorporateIndustry("education")
	// Energy.
	CreateUserRequestProfessionalOrganizationCorporateIndustryEnergy = CreateUserRequestProfessionalOrganizationCorporateIndustry("energy")
	// Financial Services & Insurance.
	CreateUserRequestProfessionalOrganizationCorporateIndustryFinancialServicesInsurance = CreateUserRequestProfessionalOrganizationCorporateIndustry("financial_services_insurance")
	// Gaming & Entertainment.
	CreateUserRequestProfessionalOrganizationCorporateIndustryGamingEntertainment = CreateUserRequestProfessionalOrganizationCorporateIndustry("gaming_entertainment")
	// Hospitality & Leisure.
	CreateUserRequestProfessionalOrganizationCorporateIndustryHospitalityLeisure = CreateUserRequestProfessionalOrganizationCorporateIndustry("hospitality_leisure")
	// Lifescience, Healthcare & Pharmaceuticals.
	CreateUserRequestProfessionalOrganizationCorporateIndustryLifescienceHealthcarePharmaceuticals = CreateUserRequestProfessionalOrganizationCorporateIndustry("lifescience_healthcare_pharmaceuticals")
	// Manufacturing.
	CreateUserRequestProfessionalOrganizationCorporateIndustryManufacturing = CreateUserRequestProfessionalOrganizationCorporateIndustry("manufacturing")
	// Media - Press & TV.
	CreateUserRequestProfessionalOrganizationCorporateIndustryMediaPressTv = CreateUserRequestProfessionalOrganizationCorporateIndustry("media_press_tv")
	// Public Sector.
	CreateUserRequestProfessionalOrganizationCorporateIndustryPublicSector = CreateUserRequestProfessionalOrganizationCorporateIndustry("public_sector")
	// Telecommunications.
	CreateUserRequestProfessionalOrganizationCorporateIndustryTelecommunications = CreateUserRequestProfessionalOrganizationCorporateIndustry("telecommunications")
	// Technology.
	CreateUserRequestProfessionalOrganizationCorporateIndustryTechnology = CreateUserRequestProfessionalOrganizationCorporateIndustry("technology")
	// Other.
	CreateUserRequestProfessionalOrganizationCorporateIndustryOther = CreateUserRequestProfessionalOrganizationCorporateIndustry("other")
)

func (enum CreateUserRequestProfessionalOrganizationCorporateIndustry) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_corporate_industry"
	}
	return string(enum)
}

func (enum CreateUserRequestProfessionalOrganizationCorporateIndustry) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CreateUserRequestProfessionalOrganizationCorporateIndustry) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CreateUserRequestProfessionalOrganizationCorporateIndustry(CreateUserRequestProfessionalOrganizationCorporateIndustry(tmp).String())
	return nil
}

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

type ListGDPRExportsRequestOrderBy string

const (
	// Creation date ascending.
	ListGDPRExportsRequestOrderByCreatedAtAsc = ListGDPRExportsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListGDPRExportsRequestOrderByCreatedAtDesc = ListGDPRExportsRequestOrderBy("created_at_desc")
)

func (enum ListGDPRExportsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListGDPRExportsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListGDPRExportsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListGDPRExportsRequestOrderBy(ListGDPRExportsRequestOrderBy(tmp).String())
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

type ListProjectsRequestOrderBy string

const (
	// Creation date ascending.
	ListProjectsRequestOrderByCreatedAtAsc = ListProjectsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListProjectsRequestOrderByCreatedAtDesc = ListProjectsRequestOrderBy("created_at_desc")
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

type OrganizationCurrency string

const (
	// Unknown currency.
	OrganizationCurrencyUnknownCurrency = OrganizationCurrency("unknown_currency")
	// Euro.
	OrganizationCurrencyEur = OrganizationCurrency("eur")
)

func (enum OrganizationCurrency) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_currency"
	}
	return string(enum)
}

func (enum OrganizationCurrency) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationCurrency) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationCurrency(OrganizationCurrency(tmp).String())
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

type OrganizationType string

const (
	// Unknown type.
	OrganizationTypeUnknownType = OrganizationType("unknown_type")
	// Personal.
	OrganizationTypePersonal = OrganizationType("personal")
	// Professional.
	OrganizationTypeProfessional = OrganizationType("professional")
	// Internal.
	OrganizationTypeInternal = OrganizationType("internal")
)

func (enum OrganizationType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum OrganizationType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationType(OrganizationType(tmp).String())
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
	// Invoice payment failure.
	OrganizationWarningReasonInvoicePaymentFailure = OrganizationWarningReason("invoice_payment_failure")
	// Got chargeback.
	OrganizationWarningReasonGotChargeback = OrganizationWarningReason("got_chargeback")
	// Validate TOS.
	OrganizationWarningReasonValidateTos = OrganizationWarningReason("validate_tos")
	// Validate KYC.
	OrganizationWarningReasonValidateKyc = OrganizationWarningReason("validate_kyc")
	// Payment info missing.
	OrganizationWarningReasonPaymentInfoMissing = OrganizationWarningReason("payment_info_missing")
	// Billing info missing.
	OrganizationWarningReasonBillingInfoMissing = OrganizationWarningReason("billing_info_missing")
	// Account closed.
	OrganizationWarningReasonAccountClosed = OrganizationWarningReason("account_closed")
	// Card expired or soon.
	OrganizationWarningReasonCardExpiredOrSoon = OrganizationWarningReason("card_expired_or_soon")
	// Confirm email change.
	OrganizationWarningReasonConfirmEmailChange = OrganizationWarningReason("confirm_email_change")
	// GDPR delete.
	OrganizationWarningReasonGdprDelete = OrganizationWarningReason("gdpr_delete")
	// Rate limiting.
	OrganizationWarningReasonRateLimiting = OrganizationWarningReason("rate_limiting")
	// Trusted level security check.
	OrganizationWarningReasonTrustedLevelSecurityCheck = OrganizationWarningReason("trusted_level_security_check")
	// Online invoice payment failure.
	OrganizationWarningReasonOnlineInvoicePaymentFailure = OrganizationWarningReason("online_invoice_payment_failure")
	// Validate email address.
	OrganizationWarningReasonValidateEmail = OrganizationWarningReason("validate_email")
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

// UserOrganizationSummary: user organization summary.
type UserOrganizationSummary struct {
	ID string `json:"id"`

	Name string `json:"name"`

	IsOwner bool `json:"is_owner"`

	Locked bool `json:"locked"`

	MfaEnforced bool `json:"mfa_enforced"`
}

// OrganizationCurrentSupportPlanTechnicalAccountManager: organization current support plan technical account manager.
type OrganizationCurrentSupportPlanTechnicalAccountManager struct {
	ID string `json:"id"`

	LastName string `json:"last_name"`

	FirstName string `json:"first_name"`

	Email string `json:"email"`

	PhoneNumber string `json:"phone_number"`
}

// CreateUserRequestPersonalOrganization: create user request personal organization.
type CreateUserRequestPersonalOrganization struct {
}

// CreateUserRequestProfessionalOrganization: create user request professional organization.
type CreateUserRequestProfessionalOrganization struct {
	OrganizationName string `json:"organization_name"`

	IsStartup bool `json:"is_startup"`

	// CorporateIndustry: default value: unknown_corporate_industry
	CorporateIndustry CreateUserRequestProfessionalOrganizationCorporateIndustry `json:"corporate_industry"`
}

// User: user.
type User struct {
	// AccountRootUserID: ID of the user.
	AccountRootUserID string `json:"account_root_user_id"`

	// FirstName: first name of the user.
	FirstName string `json:"first_name"`

	// LastName: last name of the user.
	LastName string `json:"last_name"`

	// Email: email of the user.
	Email string `json:"email"`

	// PhoneNumber: phone number of the user.
	PhoneNumber string `json:"phone_number"`

	// CreatedAt: creation date of the user.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the user.
	UpdatedAt *time.Time `json:"updated_at"`

	// MfaEnabled: multi-factor authentication enabled for the user.
	MfaEnabled bool `json:"mfa_enabled"`

	// Locale: locale.
	// Default value: unknown_language_code
	Locale std.LanguageCode `json:"locale"`

	// HasPassword: set to true if the user has set a password.
	HasPassword bool `json:"has_password"`

	// Organizations: organizations of the user.
	Organizations []*UserOrganizationSummary `json:"organizations"`
}

// SupportPlan: support plan.
type SupportPlan struct {
	// ID: ID of the support plan.
	ID string `json:"id"`

	// Level: sku for the support plan.
	// Default value: unknown_level
	Level SupportPlanLevel `json:"level"`

	// StartedAt: creation date of the support plan.
	StartedAt *time.Time `json:"started_at"`

	// StoppedAt: end date of the support plan.
	StoppedAt *time.Time `json:"stopped_at"`
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

// GDPRExport: gdpr export.
type GDPRExport struct {
	// Filename: name of file.
	Filename string `json:"filename"`

	// Size: size of file.
	Size scw.Size `json:"size"`

	// CreatedAt: creation date of export file.
	CreatedAt *time.Time `json:"created_at"`

	// ExpiresAt: deletion date of export file.
	ExpiresAt *time.Time `json:"expires_at"`
}

// GDPRRequest: gdpr request.
type GDPRRequest struct {
	// ID: ID of GDPR request.
	ID string `json:"id"`

	// AccountRootUserID: ID of the owner of request.
	AccountRootUserID string `json:"account_root_user_id"`

	// Type: type of request.
	// Default value: unknown_type
	Type GDPRRequestType `json:"type"`

	// Status: status of request.
	// Default value: unknown_status
	Status GDPRRequestStatus `json:"status"`

	// Comment: optional comment.
	Comment *string `json:"comment"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: update date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// Project: project.
type Project struct {
	// ID: ID of the Project.
	ID string `json:"id"`

	// Name: name of the Project.
	Name string `json:"name"`

	// OrganizationID: organization ID of the Project.
	OrganizationID string `json:"organization_id"`

	// CreatedAt: creation date of the Project.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: update date of the Project.
	UpdatedAt *time.Time `json:"updated_at"`

	// Description: description of the Project.
	Description string `json:"description"`
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

// OrganizationAddress: organization address.
type OrganizationAddress struct {
	AddressLine1 string `json:"address_line1"`

	AddressLine2 string `json:"address_line2"`

	CityName string `json:"city_name"`

	ZipCode string `json:"zip_code"`

	SubdivisionCode string `json:"subdivision_code"`

	CountryCode string `json:"country_code"`
}

// OrganizationCurrentSupportPlan: organization current support plan.
type OrganizationCurrentSupportPlan struct {
	// Level: default value: unknown_level
	Level SupportPlanLevel `json:"level"`

	SupportID string `json:"support_id"`

	SupportPin string `json:"support_pin"`

	TechnicalAccountManager *OrganizationCurrentSupportPlanTechnicalAccountManager `json:"technical_account_manager"`
}

// OrganizationDediboxAccount: organization dedibox account.
type OrganizationDediboxAccount struct {
	ID uint32 `json:"id"`

	LinkedAt *time.Time `json:"linked_at"`
}

// OrganizationWarning: organization warning.
type OrganizationWarning struct {
	ID string `json:"id"`

	// Reason: default value: unknown_reason
	Reason OrganizationWarningReason `json:"reason"`

	Locking bool `json:"locking"`

	Closed bool `json:"closed"`
}

// CaptchaProvider: captcha provider.
type CaptchaProvider struct {
	// Name: default value: unknown_name
	Name CaptchaProviderName `json:"name"`
}

// CreateProjectRequest: create project request.
type CreateProjectRequest struct {
	// Name: name of the Project.
	Name string `json:"name"`

	// OrganizationID: organization ID of the Project.
	OrganizationID string `json:"organization_id"`

	// Description: description of the Project.
	Description *string `json:"description,omitempty"`
}

// CreateUserResponse: create user response.
type CreateUserResponse struct {
	// User: user created.
	User *User `json:"user"`

	// Secret: secret authenticating the user.
	Secret string `json:"secret"`
}

// DeleteProjectRequest: delete project request.
type DeleteProjectRequest struct {
	// ProjectID: project ID of the Project.
	ProjectID string `json:"-"`
}

// GDPRApiCreateGDPRRequestRequest: gdpr api create gdpr request request.
type GDPRApiCreateGDPRRequestRequest struct {
	// AccountRootUserID: ID of the requester.
	AccountRootUserID string `json:"account_root_user_id"`

	// Type: type of request.
	// Default value: unknown_type
	Type GDPRRequestType `json:"type"`

	// Comment: optional comment.
	Comment *string `json:"comment,omitempty"`
}

// GDPRApiDownloadGDPRExportRequest: gdpr api download gdpr export request.
type GDPRApiDownloadGDPRExportRequest struct {
	// Filename: name of file to download.
	Filename string `json:"-"`

	// AccountRootUserID: ID of the export owner.
	AccountRootUserID string `json:"-"`
}

// GDPRApiListGDPRExportsRequest: gdpr api list gdpr exports request.
type GDPRApiListGDPRExportsRequest struct {
	// OrderBy: order by criteria.
	// Default value: created_at_asc
	OrderBy ListGDPRExportsRequestOrderBy `json:"-"`

	// Page: number of page.
	Page *int32 `json:"-"`

	// PageSize: number of items to return.
	PageSize *uint32 `json:"-"`

	// AccountRootUserID: ID of the owner of exports.
	AccountRootUserID string `json:"-"`
}

// GDPRApiListGDPRRequestsRequest: gdpr api list gdpr requests request.
type GDPRApiListGDPRRequestsRequest struct {
	// OrderBy: order by criteria.
	// Default value: type_asc
	OrderBy ListGDPRRequestsRequestOrderBy `json:"-"`

	// Page: number of page.
	Page *int32 `json:"-"`

	// PageSize: number of items to return.
	PageSize *uint32 `json:"-"`

	// AccountRootUserID: ID of the owner of requests.
	AccountRootUserID string `json:"-"`

	// Status: status of GDPR request.
	// Default value: unknown_status
	Status GDPRRequestStatus `json:"-"`
}

// GetActiveSupportPlanResponse: get active support plan response.
type GetActiveSupportPlanResponse struct {
	// ActiveSupportPlan: active support plan information.
	ActiveSupportPlan *SupportPlan `json:"active_support_plan"`

	// SupportID: support id of the organization.
	SupportID string `json:"support_id"`

	// SupportPin: support pin of the organization.
	SupportPin string `json:"support_pin"`

	// TechnicalAccountManager: technical account manager attached to the organization.
	TechnicalAccountManager *TechnicalAccountManager `json:"technical_account_manager"`
}

// GetProjectRequest: get project request.
type GetProjectRequest struct {
	// ProjectID: project ID of the Project.
	ProjectID string `json:"-"`
}

// ListGDPRExportsResponse: list gdpr exports response.
type ListGDPRExportsResponse struct {
	// GdprExports: list of exports.
	GdprExports []*GDPRExport `json:"gdpr_exports"`

	// TotalCount: total number of exports.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListGDPRExportsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListGDPRExportsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListGDPRExportsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.GdprExports = append(r.GdprExports, results.GdprExports...)
	r.TotalCount += uint32(len(results.GdprExports))
	return uint32(len(results.GdprExports)), nil
}

// ListGDPRRequestsResponse: list gdpr requests response.
type ListGDPRRequestsResponse struct {
	// GdprRequests: list of GDPR requests.
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

// ListProjectsRequest: list projects request.
type ListProjectsRequest struct {
	// OrganizationID: organization ID of the Project.
	OrganizationID string `json:"-"`

	// Name: name of the Project.
	Name *string `json:"-"`

	// Page: page number for the returned Projects.
	Page *int32 `json:"-"`

	// PageSize: maximum number of Project per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned Projects.
	// Default value: created_at_asc
	OrderBy ListProjectsRequestOrderBy `json:"-"`

	// ProjectIDs: project IDs to filter for. The results will be limited to any Projects with an ID in this array.
	ProjectIDs []string `json:"-"`
}

// ListProjectsResponse: list projects response.
type ListProjectsResponse struct {
	// TotalCount: total number of Projects.
	TotalCount uint32 `json:"total_count"`

	// Projects: paginated returned Projects.
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

// ListSupportPlansResponse: list support plans response.
type ListSupportPlansResponse struct {
	// SupportPlans: list of support plans.
	SupportPlans []*SupportPlan `json:"support_plans"`

	// TotalCount: total count of support plans.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSupportPlansResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSupportPlansResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSupportPlansResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SupportPlans = append(r.SupportPlans, results.SupportPlans...)
	r.TotalCount += uint32(len(results.SupportPlans))
	return uint32(len(results.SupportPlans)), nil
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

// Organization: organization.
type Organization struct {
	// ID: ID of the Organization.
	ID string `json:"id"`

	// Name: name of the Organization.
	Name string `json:"name"`

	// CreatedAt: creation date of the organization.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the organization.
	UpdatedAt *time.Time `json:"updated_at"`

	// Address: postal address of the organization.
	Address *OrganizationAddress `json:"address"`

	// VatNumber: vAT number for tax exemption.
	VatNumber *string `json:"vat_number"`

	// Type: type of this organization (personal or professional).
	// Default value: unknown_type
	Type OrganizationType `json:"type"`

	// Currency: currency defined in the signed Terms of Service.
	// Default value: unknown_currency
	Currency OrganizationCurrency `json:"currency"`

	// Warnings: all warnings attached to this organization.
	Warnings []*OrganizationWarning `json:"warnings"`

	// SupportPlan: current support plan information for this organization.
	SupportPlan *OrganizationCurrentSupportPlan `json:"support_plan"`

	// PaymentMode: payment mode used for this Organization (card or SEPA).
	// Default value: unknown_payment_mode
	PaymentMode OrganizationPaymentMode `json:"payment_mode"`

	// BillingEmail: email to which invoices are sent. If empty, the Owner's email is used.
	BillingEmail string `json:"billing_email"`

	// MfaEnforced: set to true if MFA is required for the organization.
	MfaEnforced bool `json:"mfa_enforced"`

	// DediboxAccount: contains ID and link creation date if a dedibox account is linked.
	DediboxAccount *OrganizationDediboxAccount `json:"dedibox_account"`

	// IsStartup: set to true if the organization is identified as a startup.
	IsStartup bool `json:"is_startup"`

	// CorporateIndustry: corporate industry of the organization.
	// Default value: unknown_corporate_industry
	CorporateIndustry OrganizationCorporateIndustry `json:"corporate_industry"`
}

// OrganizationAPICreateSupportPlanRequest: organization api create support plan request.
type OrganizationAPICreateSupportPlanRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"organization_id"`

	// Level: wanted level of support plan. Can only be basic, silver or gold.
	// Default value: unknown_level
	Level SupportPlanLevel `json:"level"`
}

// OrganizationAPIGetActiveSupportPlanRequest: organization api get active support plan request.
type OrganizationAPIGetActiveSupportPlanRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`
}

// OrganizationAPIGetOrganizationRequest: organization api get organization request.
type OrganizationAPIGetOrganizationRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`
}

// OrganizationAPIListSupportPlansRequest: organization api list support plans request.
type OrganizationAPIListSupportPlansRequest struct {
	// OrderBy: sort order of support plans.
	// Default value: started_at_asc
	OrderBy ListSupportPlansRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equals to 1.
	Page *int32 `json:"-"`

	// PageSize: number of items per page. Value must be between 1 and 1000.
	PageSize *uint32 `json:"-"`

	// OrganizationID: filter by organization ID.
	OrganizationID string `json:"-"`
}

// OrganizationAPISendMFAEnableReminderEmailRequest: organization api send mfa enable reminder email request.
type OrganizationAPISendMFAEnableReminderEmailRequest struct {
	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"-"`

	// SenderAccountRootUserID: ID of the user sending the email.
	SenderAccountRootUserID string `json:"sender_account_root_user_id"`

	// AccountRootUserID: ID of the user receiving the email.
	AccountRootUserID string `json:"account_root_user_id"`
}

// OrganizationAPIUpdateOrganizationPaymentModeRequest: organization api update organization payment mode request.
type OrganizationAPIUpdateOrganizationPaymentModeRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`

	// PaymentMode: new default payment mode. Permissible payment modes are card or SEPA.
	// Default value: unknown_payment_mode
	PaymentMode OrganizationPaymentMode `json:"payment_mode"`
}

// OrganizationAPIUpdateOrganizationRequest: organization api update organization request.
type OrganizationAPIUpdateOrganizationRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`

	// Name: name of the organization.
	Name *string `json:"name,omitempty"`

	// Type: type of this organization (personal or professional).
	// Default value: unknown_type
	Type OrganizationType `json:"type"`

	// BillingEmail: email to which invoices are sent. If empty, the Owner's email is used.
	BillingEmail *string `json:"billing_email,omitempty"`

	// VatNumber: vAT number for tax exemption.
	VatNumber *string `json:"vat_number,omitempty"`

	// AddressLine1: first address line.
	AddressLine1 *string `json:"address_line1,omitempty"`

	// AddressLine2: second address line.
	AddressLine2 *string `json:"address_line2,omitempty"`

	// AddressCountryCode: address Country Code of the organization. It uses the ISO 3166-1 alpha-2 standard.
	AddressCountryCode *string `json:"address_country_code,omitempty"`

	// AddressCityName: address City Name of the organization.
	AddressCityName *string `json:"address_city_name,omitempty"`

	// AddressZipCode: address Zip Code of the organization.
	AddressZipCode *string `json:"address_zip_code,omitempty"`

	// AddressSubdivisionCode: address Subdivision Code of the Organization. It uses the ISO 3166-2 standard.
	AddressSubdivisionCode *string `json:"address_subdivision_code,omitempty"`

	// MfaEnforced: set to true to ensure MFA is required for the organization.
	MfaEnforced *bool `json:"mfa_enforced,omitempty"`
}

// PhoneValidation: phone validation.
type PhoneValidation struct {
	// ID: the ID of the phone validation.
	ID string `json:"id"`
}

// PhoneValidationAPISendPhoneValidationRequest: phone validation api send phone validation request.
type PhoneValidationAPISendPhoneValidationRequest struct {
	// UserID: the ID of the user.
	UserID string `json:"user_id"`

	// PhoneNumber: the phone number to validate.
	PhoneNumber string `json:"phone_number"`
}

// PhoneValidationAPIValidatePhoneValidationRequest: phone validation api validate phone validation request.
type PhoneValidationAPIValidatePhoneValidationRequest struct {
	// PhoneValidationID: the ID of the phone validation.
	PhoneValidationID string `json:"-"`

	// Code: the code to validate the phone numbers.
	Code string `json:"code"`
}

// UnauthenticatedOrganizationAPIValidateAddressUpdateRequest: unauthenticated organization api validate address update request.
type UnauthenticatedOrganizationAPIValidateAddressUpdateRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`

	// Token: token received to validate the address update.
	Token string `json:"token"`
}

// UnauthenticatedOrganizationAPIValidateContractsRequest: unauthenticated organization api validate contracts request.
type UnauthenticatedOrganizationAPIValidateContractsRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`

	// Token: token received to validate the contracts.
	Token string `json:"token"`

	// ContractIDs: iDs of the contracts to validate.
	ContractIDs []string `json:"contract_ids"`
}

// UnauthenticatedUserAPICreateUserRequest: unauthenticated user api create user request.
type UnauthenticatedUserAPICreateUserRequest struct {
	// Email: email of the user.
	Email string `json:"email"`

	// FirstName: first name of the user.
	FirstName string `json:"first_name"`

	// LastName: last name of the user.
	LastName string `json:"last_name"`

	// Locale: locale.
	// Default value: unknown_language_code
	Locale std.LanguageCode `json:"locale"`

	// PersonalOrganization: organization type. Personal is the default value.
	// Precisely one of PersonalOrganization, ProfessionalOrganization must be set.
	PersonalOrganization *CreateUserRequestPersonalOrganization `json:"personal_organization,omitempty"`

	// ProfessionalOrganization: organization type. If this is a corporate account, it might be tagged as professional.
	// Precisely one of PersonalOrganization, ProfessionalOrganization must be set.
	ProfessionalOrganization *CreateUserRequestProfessionalOrganization `json:"professional_organization,omitempty"`

	// Captcha: captcha validation code. This is required as an anti-fraud measure.
	Captcha string `json:"captcha"`

	// CaptchaProviderName: captcha provider name.
	// Default value: unknown_name
	CaptchaProviderName CaptchaProviderName `json:"captcha_provider_name"`
}

// UnauthenticatedUserAPIGetCaptchaProviderRequest: unauthenticated user api get captcha provider request.
type UnauthenticatedUserAPIGetCaptchaProviderRequest struct {
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

// UnauthenticatedUserAPIRejectEmailUpdateRequest: unauthenticated user api reject email update request.
type UnauthenticatedUserAPIRejectEmailUpdateRequest struct {
	// AccountRootUserID: ID of the user.
	AccountRootUserID string `json:"-"`

	// Token: token received to reject the email update.
	Token string `json:"token"`
}

// UnauthenticatedUserAPIResetPasswordRequest: unauthenticated user api reset password request.
type UnauthenticatedUserAPIResetPasswordRequest struct {
	// Token: token received to reset the password.
	Token string `json:"token"`

	// Password: new password.
	Password string `json:"password"`
}

// UnauthenticatedUserAPISendResetPasswordEmailRequest: unauthenticated user api send reset password email request.
type UnauthenticatedUserAPISendResetPasswordEmailRequest struct {
	// Email: email associated to the user account.
	Email string `json:"email"`
}

// UnauthenticatedUserAPIValidateAccountRequest: unauthenticated user api validate account request.
type UnauthenticatedUserAPIValidateAccountRequest struct {
	// Email: email used for the account request.
	Email string `json:"-"`

	// Token: token received to validate the account.
	Token string `json:"token"`
}

// UnauthenticatedUserAPIValidateEmailUpdateRequest: unauthenticated user api validate email update request.
type UnauthenticatedUserAPIValidateEmailUpdateRequest struct {
	// AccountRootUserID: ID of the user.
	AccountRootUserID string `json:"-"`

	// Token: token received to validate the email update.
	Token string `json:"token"`
}

// UpdateProjectRequest: update project request.
type UpdateProjectRequest struct {
	// ProjectID: project ID of the Project.
	ProjectID string `json:"-"`

	// Name: name of the Project.
	Name *string `json:"name,omitempty"`

	// Description: description of the Project.
	Description *string `json:"description,omitempty"`
}

// UserAPIGetUserRequest: user api get user request.
type UserAPIGetUserRequest struct {
	// AccountRootUserID: ID of the user.
	AccountRootUserID string `json:"-"`
}

// UserAPIUpdateUserRequest: user api update user request.
type UserAPIUpdateUserRequest struct {
	// AccountRootUserID: ID of the user.
	AccountRootUserID string `json:"-"`

	// FirstName: first name of the user.
	FirstName *string `json:"first_name,omitempty"`

	// LastName: last name of the user.
	LastName *string `json:"last_name,omitempty"`

	// Email: email of the user.
	Email *string `json:"email,omitempty"`

	// Password: new password to set. This is used to update the password.
	Password *string `json:"password,omitempty"`

	// CurrentPassword: current user password. This is used to unset the password.
	CurrentPassword *string `json:"current_password,omitempty"`

	// Locale: locale of the user.
	// Default value: unknown_language_code
	Locale std.LanguageCode `json:"locale"`
}

// ValidateAccountResponse: validate account response.
type ValidateAccountResponse struct {
	// User: user created.
	User *User `json:"user"`

	// Secret: secret authenticating the user.
	Secret string `json:"secret"`
}

// ValidateContractsResponse: validate contracts response.
type ValidateContractsResponse struct {
	// Email: email address of the Owner of the Organization.
	Email string `json:"email"`

	// OrganizationType: type of the Organization.
	// Default value: unknown_type
	OrganizationType OrganizationType `json:"organization_type"`
}

// This API allows you to manage projects.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// Deprecated: CreateProject: Deprecated in favor of Account API v3.
// Generate a new Project for an Organization, specifying its configuration including name and description.
func (s *API) CreateProject(req *CreateProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("proj")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/projects",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: ListProjects: Deprecated in favor of Account API v3.
// List all Projects of an Organization. The response will include the total number of Projects as well as their associated Organizations, names and IDs. Other information include the creation and update date of the Project.
func (s *API) ListProjects(req *ListProjectsRequest, opts ...scw.RequestOption) (*ListProjectsResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_ids", req.ProjectIDs)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v2/projects",
		Query:  query,
	}

	var resp ListProjectsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: GetProject: Deprecated in favor of Account API v3.
// Retrieve information about an existing Project, specified by its Project ID. Its full details, including ID, name and description, are returned in the response object.
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
		Path:   "/account/v2/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: DeleteProject: Deprecated in favor of Account API v3.
// Delete an existing Project, specified by its Project ID. The Project needs to be empty (meaning there are no resources left in it) to be deleted effectively. Note that deleting a Project is permanent, and cannot be undone.
func (s *API) DeleteProject(req *DeleteProjectRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account/v2/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: UpdateProject: Deprecated in favor of Account API v3.
// Update the parameters of an existing Project, specified by its Project ID. These parameters include the name and description.
func (s *API) UpdateProject(req *UpdateProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return nil, errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account/v2/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GDPR Private API.
type GDPRAPI struct {
	client *scw.Client
}

// NewGDPRAPI returns a GDPRAPI object from a Scaleway client.
func NewGDPRAPI(client *scw.Client) *GDPRAPI {
	return &GDPRAPI{
		client: client,
	}
}

// CreateGDPRRequest: Create a new GDPR request for user.
func (s *GDPRAPI) CreateGDPRRequest(req *GDPRApiCreateGDPRRequestRequest, opts ...scw.RequestOption) (*GDPRRequest, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/gdpr-requests",
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

// ListGDPRRequests: List the GDPR requests created by user.
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

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v2/gdpr-requests",
		Query:  query,
	}

	var resp ListGDPRRequestsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListGDPRExports: List the available GDPR exports for user.
func (s *GDPRAPI) ListGDPRExports(req *GDPRApiListGDPRExportsRequest, opts ...scw.RequestOption) (*ListGDPRExportsResponse, error) {
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

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v2/gdpr-exports",
		Query:  query,
	}

	var resp ListGDPRExportsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DownloadGDPRExport: Download a ZIP file including a JSON and a HTML file with user data.
func (s *GDPRAPI) DownloadGDPRExport(req *GDPRApiDownloadGDPRExportRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "account_root_user_id", req.AccountRootUserID)

	if fmt.Sprint(req.Filename) == "" {
		return nil, errors.New("field Filename cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v2/gdpr-exports/" + fmt.Sprint(req.Filename) + "/download",
		Query:  query,
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage a Scaleway Organization.
type OrganizationAPI struct {
	client *scw.Client
}

// NewOrganizationAPI returns a OrganizationAPI object from a Scaleway client.
func NewOrganizationAPI(client *scw.Client) *OrganizationAPI {
	return &OrganizationAPI{
		client: client,
	}
}

// GetOrganization: Get the Organization associated with the given ID.
func (s *OrganizationAPI) GetOrganization(req *OrganizationAPIGetOrganizationRequest, opts ...scw.RequestOption) (*Organization, error) {
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
		Path:   "/account/v2/organizations/" + fmt.Sprint(req.OrganizationID) + "",
	}

	var resp Organization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOrganization: Patch the Organization associated with the given ID. You can update parameters including address details and billing email.
func (s *OrganizationAPI) UpdateOrganization(req *OrganizationAPIUpdateOrganizationRequest, opts ...scw.RequestOption) (*Organization, error) {
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
		Path:   "/account/v2/organizations/" + fmt.Sprint(req.OrganizationID) + "",
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

// UpdateOrganizationPaymentMode: Update the Organization payment mode associated with the given ID. Permissible payment modes are card and SEPA only.
func (s *OrganizationAPI) UpdateOrganizationPaymentMode(req *OrganizationAPIUpdateOrganizationPaymentModeRequest, opts ...scw.RequestOption) (*Organization, error) {
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
		Path:   "/account/v2/organizations/" + fmt.Sprint(req.OrganizationID) + "/payment-mode",
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

// ListSupportPlans: List an organizations support plans.
func (s *OrganizationAPI) ListSupportPlans(req *OrganizationAPIListSupportPlansRequest, opts ...scw.RequestOption) (*ListSupportPlansResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v2/support-plans",
		Query:  query,
	}

	var resp ListSupportPlansResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSupportPlan: Create a support plan by specifying the required support plan level, and the Organization ID.
func (s *OrganizationAPI) CreateSupportPlan(req *OrganizationAPICreateSupportPlanRequest, opts ...scw.RequestOption) (*SupportPlan, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/support-plans",
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

// GetActiveSupportPlan: Get the existing support plan of the Organization associated with the given ID, including Technical Account Manager and other details.
func (s *OrganizationAPI) GetActiveSupportPlan(req *OrganizationAPIGetActiveSupportPlanRequest, opts ...scw.RequestOption) (*GetActiveSupportPlanResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v2/support-plans/active-support-plan",
		Query:  query,
	}

	var resp GetActiveSupportPlanResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendMFAEnableReminderEmail: Send an email to a user to remind them to enable MFA on their account. Useful for Organizations where MFA is enforced for members.
func (s *OrganizationAPI) SendMFAEnableReminderEmail(req *OrganizationAPISendMFAEnableReminderEmailRequest, opts ...scw.RequestOption) error {
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
		Path:   "/account/v2/organizations/" + fmt.Sprint(req.OrganizationID) + "/send-mfa-enable-reminder-email",
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

// This API allows you to validate phone numbers.
type PhoneValidationAPI struct {
	client *scw.Client
}

// NewPhoneValidationAPI returns a PhoneValidationAPI object from a Scaleway client.
func NewPhoneValidationAPI(client *scw.Client) *PhoneValidationAPI {
	return &PhoneValidationAPI{
		client: client,
	}
}

// SendPhoneValidation: Send Phone Validation.
func (s *PhoneValidationAPI) SendPhoneValidation(req *PhoneValidationAPISendPhoneValidationRequest, opts ...scw.RequestOption) (*PhoneValidation, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/phone-validations",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PhoneValidation

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidatePhoneValidation: Validate Phone Number.
func (s *PhoneValidationAPI) ValidatePhoneValidation(req *PhoneValidationAPIValidatePhoneValidationRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.PhoneValidationID) == "" {
		return errors.New("field PhoneValidationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/phone-validations/" + fmt.Sprint(req.PhoneValidationID) + "/validate",
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

// This API allows to manage an organization without authentication.
type UnauthenticatedOrganizationAPI struct {
	client *scw.Client
}

// NewUnauthenticatedOrganizationAPI returns a UnauthenticatedOrganizationAPI object from a Scaleway client.
func NewUnauthenticatedOrganizationAPI(client *scw.Client) *UnauthenticatedOrganizationAPI {
	return &UnauthenticatedOrganizationAPI{
		client: client,
	}
}

// ValidateContracts: Validate contracts.
func (s *UnauthenticatedOrganizationAPI) ValidateContracts(req *UnauthenticatedOrganizationAPIValidateContractsRequest, opts ...scw.RequestOption) (*ValidateContractsResponse, error) {
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
		Path:   "/account/v2/organizations/" + fmt.Sprint(req.OrganizationID) + "/validate-contracts",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ValidateContractsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateAddressUpdate: Validate address update.
func (s *UnauthenticatedOrganizationAPI) ValidateAddressUpdate(req *UnauthenticatedOrganizationAPIValidateAddressUpdateRequest, opts ...scw.RequestOption) error {
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
		Path:   "/account/v2/organizations/" + fmt.Sprint(req.OrganizationID) + "/validate-address-update",
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

// This API allows to manage a user.
type UserAPI struct {
	client *scw.Client
}

// NewUserAPI returns a UserAPI object from a Scaleway client.
func NewUserAPI(client *scw.Client) *UserAPI {
	return &UserAPI{
		client: client,
	}
}

// GetUser: Get the user associated with the given ID.
func (s *UserAPI) GetUser(req *UserAPIGetUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error

	if fmt.Sprint(req.AccountRootUserID) == "" {
		return nil, errors.New("field AccountRootUserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v2/users/" + fmt.Sprint(req.AccountRootUserID) + "",
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateUser: Patch the user associated with the given ID.
func (s *UserAPI) UpdateUser(req *UserAPIUpdateUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error

	if fmt.Sprint(req.AccountRootUserID) == "" {
		return nil, errors.New("field AccountRootUserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account/v2/users/" + fmt.Sprint(req.AccountRootUserID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp User

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

// SendResetPasswordEmail: Send reset password link to email.
func (s *UnauthenticatedUserAPI) SendResetPasswordEmail(req *UnauthenticatedUserAPISendResetPasswordEmailRequest, opts ...scw.RequestOption) error {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/send-reset-password-email",
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

// ResetPassword: Reset password.
func (s *UnauthenticatedUserAPI) ResetPassword(req *UnauthenticatedUserAPIResetPasswordRequest, opts ...scw.RequestOption) error {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/reset-password",
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

// CreateUser: Create a new user.
func (s *UnauthenticatedUserAPI) CreateUser(req *UnauthenticatedUserAPICreateUserRequest, opts ...scw.RequestOption) (*CreateUserResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/users",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateUserResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCaptchaProvider: Get a Captcha provider.
func (s *UnauthenticatedUserAPI) GetCaptchaProvider(req *UnauthenticatedUserAPIGetCaptchaProviderRequest, opts ...scw.RequestOption) (*CaptchaProvider, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v2/captcha-provider",
	}

	var resp CaptchaProvider

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateEmailUpdate: Validate email update.
func (s *UnauthenticatedUserAPI) ValidateEmailUpdate(req *UnauthenticatedUserAPIValidateEmailUpdateRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.AccountRootUserID) == "" {
		return errors.New("field AccountRootUserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/users/" + fmt.Sprint(req.AccountRootUserID) + "/validate-email-update",
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

// RejectEmailUpdate: Reject email update.
func (s *UnauthenticatedUserAPI) RejectEmailUpdate(req *UnauthenticatedUserAPIRejectEmailUpdateRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.AccountRootUserID) == "" {
		return errors.New("field AccountRootUserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/users/" + fmt.Sprint(req.AccountRootUserID) + "/reject-email-update",
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

// Login: Login user.
func (s *UnauthenticatedUserAPI) Login(req *UnauthenticatedUserAPILoginRequest, opts ...scw.RequestOption) (*LoginResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/login",
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

// ValidateAccount: Validate the user request.
func (s *UnauthenticatedUserAPI) ValidateAccount(req *UnauthenticatedUserAPIValidateAccountRequest, opts ...scw.RequestOption) (*ValidateAccountResponse, error) {
	var err error

	if fmt.Sprint(req.Email) == "" {
		return nil, errors.New("field Email cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v2/accounts/" + fmt.Sprint(req.Email) + "/validate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ValidateAccountResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

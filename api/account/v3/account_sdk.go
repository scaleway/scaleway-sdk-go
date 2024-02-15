// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package account provides methods and message types of the account v3 API.
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

type CloseOrganizationRequestReason string

const (
	// Unknown reason.
	CloseOrganizationRequestReasonUnknownReason = CloseOrganizationRequestReason("unknown_reason")
	// I no longer need Scaleway's products.
	CloseOrganizationRequestReasonNoLongerNeeded = CloseOrganizationRequestReason("no_longer_needed")
	// It's too expensive.
	CloseOrganizationRequestReasonTooExpensive = CloseOrganizationRequestReason("too_expensive")
	// I need some features you don't have.
	CloseOrganizationRequestReasonMissingFeature = CloseOrganizationRequestReason("missing_feature")
	// I've had technical issues.
	CloseOrganizationRequestReasonTechnicalIssue = CloseOrganizationRequestReason("technical_issue")
	// I just wanted to test the service.
	CloseOrganizationRequestReasonTestOnly = CloseOrganizationRequestReason("test_only")
	// I find it difficult to use.
	CloseOrganizationRequestReasonDifficultToUse = CloseOrganizationRequestReason("difficult_to_use")
	// I'm using another account.
	CloseOrganizationRequestReasonUsingAnotherAccount = CloseOrganizationRequestReason("using_another_account")
	// Other reason.
	CloseOrganizationRequestReasonOther = CloseOrganizationRequestReason("other")
)

func (enum CloseOrganizationRequestReason) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_reason"
	}
	return string(enum)
}

func (enum CloseOrganizationRequestReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CloseOrganizationRequestReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CloseOrganizationRequestReason(CloseOrganizationRequestReason(tmp).String())
	return nil
}

type CreateAccountRequestProfessionalOrganizationCorporateIndustry string

const (
	// Unknown corporate industry.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryUnknownCorporateIndustry = CreateAccountRequestProfessionalOrganizationCorporateIndustry("unknown_corporate_industry")
	// Consulting & Services.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryConsultingServices = CreateAccountRequestProfessionalOrganizationCorporateIndustry("consulting_services")
	// Cybersecurity & software.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryCybersecuritySoftware = CreateAccountRequestProfessionalOrganizationCorporateIndustry("cybersecurity_software")
	// E-commerce & retail.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryEcommerceRetail = CreateAccountRequestProfessionalOrganizationCorporateIndustry("ecommerce_retail")
	// Education.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryEducation = CreateAccountRequestProfessionalOrganizationCorporateIndustry("education")
	// Energy.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryEnergy = CreateAccountRequestProfessionalOrganizationCorporateIndustry("energy")
	// Financial Services & Insurance.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryFinancialServicesInsurance = CreateAccountRequestProfessionalOrganizationCorporateIndustry("financial_services_insurance")
	// Gaming & Entertainment.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryGamingEntertainment = CreateAccountRequestProfessionalOrganizationCorporateIndustry("gaming_entertainment")
	// Hospitality & Leisure.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryHospitalityLeisure = CreateAccountRequestProfessionalOrganizationCorporateIndustry("hospitality_leisure")
	// Lifescience, Healthcare & Pharmaceuticals.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryLifescienceHealthcarePharmaceuticals = CreateAccountRequestProfessionalOrganizationCorporateIndustry("lifescience_healthcare_pharmaceuticals")
	// Manufacturing.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryManufacturing = CreateAccountRequestProfessionalOrganizationCorporateIndustry("manufacturing")
	// Media - Press & TV.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryMediaPressTv = CreateAccountRequestProfessionalOrganizationCorporateIndustry("media_press_tv")
	// Public Sector.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryPublicSector = CreateAccountRequestProfessionalOrganizationCorporateIndustry("public_sector")
	// Telecommunications.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryTelecommunications = CreateAccountRequestProfessionalOrganizationCorporateIndustry("telecommunications")
	// Technology.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryTechnology = CreateAccountRequestProfessionalOrganizationCorporateIndustry("technology")
	// Other.
	CreateAccountRequestProfessionalOrganizationCorporateIndustryOther = CreateAccountRequestProfessionalOrganizationCorporateIndustry("other")
)

func (enum CreateAccountRequestProfessionalOrganizationCorporateIndustry) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_corporate_industry"
	}
	return string(enum)
}

func (enum CreateAccountRequestProfessionalOrganizationCorporateIndustry) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CreateAccountRequestProfessionalOrganizationCorporateIndustry) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CreateAccountRequestProfessionalOrganizationCorporateIndustry(CreateAccountRequestProfessionalOrganizationCorporateIndustry(tmp).String())
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
	// Name ascending.
	ListContractSignaturesRequestOrderByNameAsc = ListContractSignaturesRequestOrderBy("name_asc")
	// Name descending.
	ListContractSignaturesRequestOrderByNameDesc = ListContractSignaturesRequestOrderBy("name_desc")
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

type ListInvitationsRequestOrderBy string

const (
	// Creation date ascending.
	ListInvitationsRequestOrderByCreatedAtAsc = ListInvitationsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListInvitationsRequestOrderByCreatedAtDesc = ListInvitationsRequestOrderBy("created_at_desc")
	// Expiration date ascending.
	ListInvitationsRequestOrderByExpiresAtAsc = ListInvitationsRequestOrderBy("expires_at_asc")
	// Expiration date descending.
	ListInvitationsRequestOrderByExpiresAtDesc = ListInvitationsRequestOrderBy("expires_at_desc")
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

type ListMFAOTPsRequestOrderBy string

const (
	// Creation date ascending.
	ListMFAOTPsRequestOrderByCreatedAtAsc = ListMFAOTPsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListMFAOTPsRequestOrderByCreatedAtDesc = ListMFAOTPsRequestOrderBy("created_at_desc")
)

func (enum ListMFAOTPsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListMFAOTPsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListMFAOTPsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListMFAOTPsRequestOrderBy(ListMFAOTPsRequestOrderBy(tmp).String())
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

type PasswordStrengthScore string

const (
	// Unknown score.
	PasswordStrengthScoreUnknownScore = PasswordStrengthScore("unknown_score")
	// Very weak.
	PasswordStrengthScoreVeryWeak = PasswordStrengthScore("very_weak")
	// Weak.
	PasswordStrengthScoreWeak = PasswordStrengthScore("weak")
	// Medium.
	PasswordStrengthScoreMedium = PasswordStrengthScore("medium")
	// Strong.
	PasswordStrengthScoreStrong = PasswordStrengthScore("strong")
	// Very strong.
	PasswordStrengthScoreVeryStrong = PasswordStrengthScore("very_strong")
)

func (enum PasswordStrengthScore) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_score"
	}
	return string(enum)
}

func (enum PasswordStrengthScore) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PasswordStrengthScore) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PasswordStrengthScore(PasswordStrengthScore(tmp).String())
	return nil
}

type SupportPlanLevel string

const (
	// Unknown level.
	SupportPlanLevelUnknownLevel = SupportPlanLevel("unknown_level")
	// Basic (deprecated).
	SupportPlanLevelBasic = SupportPlanLevel("basic")
	// Developer (deprecated).
	SupportPlanLevelDeveloper = SupportPlanLevel("developer")
	// Business (deprecated).
	SupportPlanLevelBusiness = SupportPlanLevel("business")
	// Enterprise (deprecated).
	SupportPlanLevelEnterprise = SupportPlanLevel("enterprise")
	// Bronze (deprecated).
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

type UserPreferencesConsoleTheme string

const (
	// Unknown console theme.
	UserPreferencesConsoleThemeUnknownConsoleTheme = UserPreferencesConsoleTheme("unknown_console_theme")
	// Light.
	UserPreferencesConsoleThemeLight = UserPreferencesConsoleTheme("light")
	// Dark.
	UserPreferencesConsoleThemeDark = UserPreferencesConsoleTheme("dark")
	// System.
	UserPreferencesConsoleThemeSystem = UserPreferencesConsoleTheme("system")
	// Darker.
	UserPreferencesConsoleThemeDarker = UserPreferencesConsoleTheme("darker")
)

func (enum UserPreferencesConsoleTheme) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_console_theme"
	}
	return string(enum)
}

func (enum UserPreferencesConsoleTheme) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UserPreferencesConsoleTheme) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UserPreferencesConsoleTheme(UserPreferencesConsoleTheme(tmp).String())
	return nil
}

// Contract: contract.
type Contract struct {
	// ID: ID of the contract.
	ID string `json:"id"`

	// Name: the name of the contract.
	Name string `json:"name"`

	// Version: the version of the contract.
	Version uint32 `json:"version"`

	// CreatedAt: the creation date of the contract.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: the last modification date of the contract.
	UpdatedAt *time.Time `json:"updated_at"`
}

// InvitationOrganization: invitation organization.
type InvitationOrganization struct {
	ID string `json:"id"`

	Name string `json:"name"`

	MfaEnforced bool `json:"mfa_enforced"`
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

// CreateAccountRequestPersonalOrganization: create account request personal organization.
type CreateAccountRequestPersonalOrganization struct {
}

// CreateAccountRequestProfessionalOrganization: create account request professional organization.
type CreateAccountRequestProfessionalOrganization struct {
	OrganizationName string `json:"organization_name"`

	Startup bool `json:"startup"`

	// CorporateIndustry: default value: unknown_corporate_industry
	CorporateIndustry CreateAccountRequestProfessionalOrganizationCorporateIndustry `json:"corporate_industry"`
}

// ContractSignature: contract signature.
type ContractSignature struct {
	// ID: ID of the contract signature.
	ID string `json:"id"`

	// OrganizationID: the Organization ID which signed the contract.
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

// ListCountriesResponseCountry: list countries response country.
type ListCountriesResponseCountry struct {
	// Name: name of the country.
	Name string `json:"name"`

	// Code: two letters representing the country.
	Code string `json:"code"`

	// Flag: flag representation of the country.
	Flag string `json:"flag"`

	// VatNumberRequired: true if the country should have a VAT identification number, false otherwise.
	VatNumberRequired bool `json:"vat_number_required"`
}

// ListCountrySubdivisionsResponseSubdivision: list country subdivisions response subdivision.
type ListCountrySubdivisionsResponseSubdivision struct {
	// Name: subdivision name.
	Name string `json:"name"`

	// Code: subdivision code is a few numbers and/or letters completing the ISO 3166-1 code of the country.
	Code string `json:"code"`
}

// Invitation: invitation.
type Invitation struct {
	// ID: ID of the invitation.
	ID string `json:"id"`

	// Organization: the Organization of the invitation.
	Organization *InvitationOrganization `json:"organization"`

	// CreatedAt: creation date of the invitation.
	CreatedAt *time.Time `json:"created_at"`

	// ExpiresAt: expiration date of the invitation.
	ExpiresAt *time.Time `json:"expires_at"`
}

// MFAOTP: mfaotp.
type MFAOTP struct {
	// ID: ID of the MFA OTP.
	ID string `json:"id"`

	// Secret: secret of the MFA OTP.
	Secret *string `json:"secret"`
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

// Account: account.
type Account struct {
	// ID: request creation id.
	ID string `json:"id"`

	// Email: email used for the account request.
	Email string `json:"email"`
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

// ContractAPIDownloadContractSignatureRequest: contract api download contract signature request.
type ContractAPIDownloadContractSignatureRequest struct {
	// ContractSignatureID: the contract signature ID.
	ContractSignatureID string `json:"-"`

	// Locale: the locale requested for the content of the contract.
	// Default value: unknown_language_code
	Locale std.LanguageCode `json:"-"`
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

	// OrganizationID: filter on Organization ID.
	OrganizationID string `json:"-"`
}

// ContractAPIValidateContractSignatureRequest: contract api validate contract signature request.
type ContractAPIValidateContractSignatureRequest struct {
	// ContractSignatureID: the contract linked to your Organization you want to sign.
	ContractSignatureID string `json:"-"`
}

// IsoCodeAPIListCountriesRequest: iso code api list countries request.
type IsoCodeAPIListCountriesRequest struct {
}

// IsoCodeAPIListCountrySubdivisionsRequest: iso code api list country subdivisions request.
type IsoCodeAPIListCountrySubdivisionsRequest struct {
	// CountryCode: the country code.
	CountryCode string `json:"-"`
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

// ListCountriesResponse: list countries response.
type ListCountriesResponse struct {
	// Countries: list of countries.
	Countries []*ListCountriesResponseCountry `json:"countries"`
}

// ListCountrySubdivisionsResponse: list country subdivisions response.
type ListCountrySubdivisionsResponse struct {
	// Subdivisions: list of subdivisions for a country code.
	Subdivisions []*ListCountrySubdivisionsResponseSubdivision `json:"subdivisions"`
}

// ListInvitationsResponse: list invitations response.
type ListInvitationsResponse struct {
	// TotalCount: the number of invitations.
	TotalCount uint64 `json:"total_count"`

	// Invitations: a list of invitations.
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

// ListMFAOTPsResponse: list mfaot ps response.
type ListMFAOTPsResponse struct {
	// TotalCount: total number of MFA OTPs.
	TotalCount uint64 `json:"total_count"`

	// MfaOtps: paginated returned MFA OTPs.
	MfaOtps []*MFAOTP `json:"mfa_otps"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListMFAOTPsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListMFAOTPsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListMFAOTPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.MfaOtps = append(r.MfaOtps, results.MfaOtps...)
	r.TotalCount += uint64(len(results.MfaOtps))
	return uint64(len(results.MfaOtps)), nil
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

// ListProjectsResponse: list projects response.
type ListProjectsResponse struct {
	// TotalCount: total number of Projects.
	TotalCount uint64 `json:"total_count"`

	// Projects: paginated returned Projects.
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

// OrganizationAPIChangeSupportPlanRequest: organization api change support plan request.
type OrganizationAPIChangeSupportPlanRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`

	// Level: the level of support plan. Can only be basic, silver or gold.
	// Default value: unknown_level
	Level SupportPlanLevel `json:"level"`
}

// OrganizationAPICloseOrganizationRequest: organization api close organization request.
type OrganizationAPICloseOrganizationRequest struct {
	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"-"`

	// Reasons: reasons to close the Organization.
	Reasons []CloseOrganizationRequestReason `json:"reasons"`

	// Details: more details about why the Organization is being closed.
	Details *string `json:"details,omitempty"`
}

// OrganizationAPIGetActiveSupportPlanRequest: organization api get active support plan request.
type OrganizationAPIGetActiveSupportPlanRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`
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

// PasswordStrength: password strength.
type PasswordStrength struct {
	// Score: default value: unknown_score
	Score PasswordStrengthScore `json:"score"`
}

// PhoneValidation: phone validation.
type PhoneValidation struct {
	// ID: the ID of the phone validation.
	ID string `json:"id"`
}

// ProjectAPICreateProjectRequest: project api create project request.
type ProjectAPICreateProjectRequest struct {
	// Name: name of the Project.
	Name string `json:"name"`

	// OrganizationID: organization ID of the Project.
	OrganizationID string `json:"organization_id"`

	// Description: description of the Project.
	Description string `json:"description"`
}

// ProjectAPIDeleteProjectRequest: project api delete project request.
type ProjectAPIDeleteProjectRequest struct {
	// ProjectID: project ID of the Project.
	ProjectID string `json:"-"`
}

// ProjectAPIGetProjectRequest: project api get project request.
type ProjectAPIGetProjectRequest struct {
	// ProjectID: project ID of the Project.
	ProjectID string `json:"-"`
}

// ProjectAPIListProjectsRequest: project api list projects request.
type ProjectAPIListProjectsRequest struct {
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

// ProjectAPIUpdateProjectRequest: project api update project request.
type ProjectAPIUpdateProjectRequest struct {
	// ProjectID: project ID of the Project.
	ProjectID string `json:"-"`

	// Name: name of the Project.
	Name *string `json:"name,omitempty"`

	// Description: description of the Project.
	Description *string `json:"description,omitempty"`
}

// SendEmailValidationResponse: send email validation response.
type SendEmailValidationResponse struct {
	// Email: email address to validate.
	Email string `json:"email"`
}

// SetMailingListSubscriptionsResponse: set mailing list subscriptions response.
type SetMailingListSubscriptionsResponse struct {
	// MailingListSubscriptions: a list of mailing list subscriptions.
	MailingListSubscriptions []*MailingListSubscription `json:"mailing_list_subscriptions"`
}

// UnauthenticatedAPIComputePasswordStrengthRequest: unauthenticated api compute password strength request.
type UnauthenticatedAPIComputePasswordStrengthRequest struct {
	Password string `json:"password"`

	UserInputs []string `json:"user_inputs"`
}

// UnauthenticatedOrganizationAPIValidateAddressUpdateRequest: unauthenticated organization api validate address update request.
type UnauthenticatedOrganizationAPIValidateAddressUpdateRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`

	// Token: token received to validate the address update.
	Token string `json:"token"`
}

// UnauthenticatedUserAPICreateAccountRequest: unauthenticated user api create account request.
type UnauthenticatedUserAPICreateAccountRequest struct {
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
	PersonalOrganization *CreateAccountRequestPersonalOrganization `json:"personal_organization,omitempty"`

	// ProfessionalOrganization: organization type. If this is a corporate account, it might be tagged as professional.
	// Precisely one of PersonalOrganization, ProfessionalOrganization must be set.
	ProfessionalOrganization *CreateAccountRequestProfessionalOrganization `json:"professional_organization,omitempty"`

	// Captcha: captcha validation code. This is required as an anti-fraud measure.
	Captcha string `json:"captcha"`
}

// UnauthenticatedUserAPINotifyAccountRequest: unauthenticated user api notify account request.
type UnauthenticatedUserAPINotifyAccountRequest struct {
	// Email: email associated to the user account.
	Email string `json:"-"`
}

// UnauthenticatedUserAPIResetPasswordRequest: unauthenticated user api reset password request.
type UnauthenticatedUserAPIResetPasswordRequest struct {
	// Token: token received to reset the password.
	Token string `json:"token"`

	// Password: new password.
	Password string `json:"password"`
}

// UnauthenticatedUserAPISendPasswordlessAuthRequest: unauthenticated user api send passwordless auth request.
type UnauthenticatedUserAPISendPasswordlessAuthRequest struct {
	// Email: email address to send the passwordless auth to.
	Email string `json:"email"`

	// RedirectURL: URL to redirect to after login.
	RedirectURL string `json:"redirect_url"`
}

// UnauthenticatedUserAPIValidateEmailRequest: unauthenticated user api validate email request.
type UnauthenticatedUserAPIValidateEmailRequest struct {
	// Email: email address to validate.
	Email string `json:"-"`

	// Token: token received to validate the email address.
	Token string `json:"token"`
}

// UserAPIAcceptInvitationRequest: user api accept invitation request.
type UserAPIAcceptInvitationRequest struct {
	// InvitationID: ID of the invitation to accept.
	InvitationID string `json:"-"`
}

// UserAPIConfirmPhoneValidationRequest: user api confirm phone validation request.
type UserAPIConfirmPhoneValidationRequest struct {
	// PhoneValidationID: the ID of the phone validation.
	PhoneValidationID string `json:"-"`

	// Code: the code required to validate the phone number.
	Code string `json:"code"`
}

// UserAPICreateMFAOTPRequest: user api create mfaotp request.
type UserAPICreateMFAOTPRequest struct {
	// UserID: user ID of the MFA OTP.
	UserID string `json:"user_id"`
}

// UserAPIDeleteMFAOTPRequest: user api delete mfaotp request.
type UserAPIDeleteMFAOTPRequest struct {
	// MfaOtpID: ID of the MFA OTP.
	MfaOtpID string `json:"-"`
}

// UserAPIGetUserPreferencesRequest: user api get user preferences request.
type UserAPIGetUserPreferencesRequest struct {
	// UserID: ID of the user.
	UserID string `json:"-"`
}

// UserAPIListInvitationsRequest: user api list invitations request.
type UserAPIListInvitationsRequest struct {
	// UserID: ID of the user.
	UserID string `json:"-"`

	// Page: the page number.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of items per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: the sort order.
	// Default value: created_at_asc
	OrderBy ListInvitationsRequestOrderBy `json:"-"`
}

// UserAPIListMFAOTPsRequest: user api list mfaot ps request.
type UserAPIListMFAOTPsRequest struct {
	// Page: page number for the returned MFA OTPs.
	Page *int32 `json:"-"`

	// PageSize: maximum number of MFA OTP per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned MFA OTPs.
	// Default value: created_at_asc
	OrderBy ListMFAOTPsRequestOrderBy `json:"-"`

	// UserID: filter out by a user ID.
	UserID string `json:"-"`
}

// UserAPIListMailingListSubscriptionsRequest: user api list mailing list subscriptions request.
type UserAPIListMailingListSubscriptionsRequest struct {
	// UserID: the ID of the user.
	UserID string `json:"-"`

	// OrganizationID: the ID of the organization. Must be defined if you want to list mailing lists related to the organization.
	OrganizationID *string `json:"-"`

	// Page: the page number.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of items per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: the sort order.
	// Default value: name_asc
	OrderBy ListMailingListSubscriptionsRequestOrderBy `json:"-"`
}

// UserAPIRefuseInvitationRequest: user api refuse invitation request.
type UserAPIRefuseInvitationRequest struct {
	// InvitationID: ID of the invitation to refuse.
	InvitationID string `json:"-"`
}

// UserAPISendEmailValidationRequest: user api send email validation request.
type UserAPISendEmailValidationRequest struct {
	// UserID: ID of the user.
	UserID string `json:"-"`
}

// UserAPISendPhoneValidationRequest: user api send phone validation request.
type UserAPISendPhoneValidationRequest struct {
	// UserID: the ID of the user.
	UserID string `json:"user_id"`

	// PhoneNumber: the phone number to validate. Must be in international format (E.164).
	PhoneNumber string `json:"phone_number"`
}

// UserAPISetMailingListSubscriptionsRequest: user api set mailing list subscriptions request.
type UserAPISetMailingListSubscriptionsRequest struct {
	// UserID: the ID of the user.
	UserID string `json:"-"`

	// MailingListSubscribed: a map of mailing list ids with their subscription status.
	MailingListSubscribed map[string]bool `json:"mailing_list_subscribed"`

	// OrganizationID: the ID of the organization. Must be defined if you want to set mailing lists subscriptions related to the organization.
	OrganizationID *string `json:"organization_id,omitempty"`
}

// UserAPIUpdateUserPreferencesRequest: user api update user preferences request.
type UserAPIUpdateUserPreferencesRequest struct {
	// UserID: ID of the user.
	UserID string `json:"-"`

	// ConsoleTheme: console theme preference.
	// Default value: unknown_console_theme
	ConsoleTheme UserPreferencesConsoleTheme `json:"console_theme"`
}

// UserAPIValidateMFAOTPRequest: user api validate mfaotp request.
type UserAPIValidateMFAOTPRequest struct {
	// MfaOtpID: ID of the MFA OTP.
	MfaOtpID string `json:"-"`

	// Otp: one time password of the MFA OTP.
	Otp string `json:"otp"`
}

// UserPreferences: user preferences.
type UserPreferences struct {
	// UserID: ID of the user.
	UserID string `json:"user_id"`

	// ConsoleTheme: console theme preference.
	// Default value: unknown_console_theme
	ConsoleTheme UserPreferencesConsoleTheme `json:"console_theme"`
}

// ValidateMFAOTPResponse: validate mfaotp response.
type ValidateMFAOTPResponse struct {
	// BackupCodes: backup codes of the MFA OTP.
	BackupCodes []string `json:"backup_codes"`
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

// DownloadContractSignature: Download a contract content.
func (s *ContractAPI) DownloadContractSignature(req *ContractAPIDownloadContractSignatureRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "locale", req.Locale)

	if fmt.Sprint(req.ContractSignatureID) == "" {
		return nil, errors.New("field ContractSignatureID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/contract-signatures/" + fmt.Sprint(req.ContractSignatureID) + "/download",
		Query:  query,
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateContractSignature: Sign a contract for your Organization.
func (s *ContractAPI) ValidateContractSignature(req *ContractAPIValidateContractSignatureRequest, opts ...scw.RequestOption) (*ContractSignature, error) {
	var err error

	if fmt.Sprint(req.ContractSignatureID) == "" {
		return nil, errors.New("field ContractSignatureID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/contract-signatures/" + fmt.Sprint(req.ContractSignatureID) + "/validate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ContractSignature

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListContractSignatures: List contract signatures for an Organization.
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
		Path:   "/account/v3/contract-signatures",
		Query:  query,
	}

	var resp ListContractSignaturesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// The IsoCode API list countries or subdivisions of a country.
type IsoCodeAPI struct {
	client *scw.Client
}

// NewIsoCodeAPI returns a IsoCodeAPI object from a Scaleway client.
func NewIsoCodeAPI(client *scw.Client) *IsoCodeAPI {
	return &IsoCodeAPI{
		client: client,
	}
}

// GetServiceInfo:
func (s *IsoCodeAPI) GetServiceInfo(opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/iso-codes",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCountries: List all countries from ISO 3166-1 alpha-2.
func (s *IsoCodeAPI) ListCountries(req *IsoCodeAPIListCountriesRequest, opts ...scw.RequestOption) (*ListCountriesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/iso-codes/countries",
	}

	var resp ListCountriesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCountrySubdivisions: Method to list all subdivisions (from ISO 3166-2 alpha-2) for a country code.
func (s *IsoCodeAPI) ListCountrySubdivisions(req *IsoCodeAPIListCountrySubdivisionsRequest, opts ...scw.RequestOption) (*ListCountrySubdivisionsResponse, error) {
	var err error

	if fmt.Sprint(req.CountryCode) == "" {
		return nil, errors.New("field CountryCode cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/iso-codes/countries/" + fmt.Sprint(req.CountryCode) + "/subdivisions",
	}

	var resp ListCountrySubdivisionsResponse

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
		Path:   "/account/v3/organizations/" + fmt.Sprint(req.OrganizationID) + "/support-plans",
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
		Path:   "/account/v3/organizations/" + fmt.Sprint(req.OrganizationID) + "/active-support-plan",
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
		Path:   "/account/v3/organizations/" + fmt.Sprint(req.OrganizationID) + "/support-plans",
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
		Path:   "/account/v3/organizations/" + fmt.Sprint(req.OrganizationID) + "/close",
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

// This API allows you to manage projects.
type ProjectAPI struct {
	client *scw.Client
}

// NewProjectAPI returns a ProjectAPI object from a Scaleway client.
func NewProjectAPI(client *scw.Client) *ProjectAPI {
	return &ProjectAPI{
		client: client,
	}
}

// CreateProject: Generate a new Project for an Organization, specifying its configuration including name and description.
func (s *ProjectAPI) CreateProject(req *ProjectAPICreateProjectRequest, opts ...scw.RequestOption) (*Project, error) {
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
		Path:   "/account/v3/projects",
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

// ListProjects: List all Projects of an Organization. The response will include the total number of Projects as well as their associated Organizations, names, and IDs. Other information includes the creation and update date of the Project.
func (s *ProjectAPI) ListProjects(req *ProjectAPIListProjectsRequest, opts ...scw.RequestOption) (*ListProjectsResponse, error) {
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
		Path:   "/account/v3/projects",
		Query:  query,
	}

	var resp ListProjectsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetProject: Retrieve information about an existing Project, specified by its Project ID. Its full details, including ID, name and description, are returned in the response object.
func (s *ProjectAPI) GetProject(req *ProjectAPIGetProjectRequest, opts ...scw.RequestOption) (*Project, error) {
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
		Path:   "/account/v3/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteProject: Delete an existing Project, specified by its Project ID. The Project needs to be empty (meaning there are no resources left in it) to be deleted effectively. Note that deleting a Project is permanent, and cannot be undone.
func (s *ProjectAPI) DeleteProject(req *ProjectAPIDeleteProjectRequest, opts ...scw.RequestOption) error {
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
		Path:   "/account/v3/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProject: Update the parameters of an existing Project, specified by its Project ID. These parameters include the name and description.
func (s *ProjectAPI) UpdateProject(req *ProjectAPIUpdateProjectRequest, opts ...scw.RequestOption) (*Project, error) {
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
		Path:   "/account/v3/projects/" + fmt.Sprint(req.ProjectID) + "",
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

type UnauthenticatedAPI struct {
	client *scw.Client
}

// NewUnauthenticatedAPI returns a UnauthenticatedAPI object from a Scaleway client.
func NewUnauthenticatedAPI(client *scw.Client) *UnauthenticatedAPI {
	return &UnauthenticatedAPI{
		client: client,
	}
}

// ComputePasswordStrength:
func (s *UnauthenticatedAPI) ComputePasswordStrength(req *UnauthenticatedAPIComputePasswordStrengthRequest, opts ...scw.RequestOption) (*PasswordStrength, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/compute-password-strength",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PasswordStrength

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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
		Path:   "/account/v3/organizations/" + fmt.Sprint(req.OrganizationID) + "/validate-address-update",
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

// CreateAccount: Create a new user request.
func (s *UnauthenticatedUserAPI) CreateAccount(req *UnauthenticatedUserAPICreateAccountRequest, opts ...scw.RequestOption) (*Account, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/accounts",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Account

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// NotifyAccount: Send the validation email for the user request.
func (s *UnauthenticatedUserAPI) NotifyAccount(req *UnauthenticatedUserAPINotifyAccountRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.Email) == "" {
		return errors.New("field Email cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/accounts/" + fmt.Sprint(req.Email) + "/notify",
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

// ValidateEmail: Validate the email address of an existing user.
func (s *UnauthenticatedUserAPI) ValidateEmail(req *UnauthenticatedUserAPIValidateEmailRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.Email) == "" {
		return errors.New("field Email cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/users/" + fmt.Sprint(req.Email) + "/validate-email",
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
		Path:   "/account/v3/reset-password",
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

// SendPasswordlessAuth: Send passwordless auth.
func (s *UnauthenticatedUserAPI) SendPasswordlessAuth(req *UnauthenticatedUserAPISendPasswordlessAuthRequest, opts ...scw.RequestOption) error {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/send-passwordless-auth",
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/users/" + fmt.Sprint(req.UserID) + "/mailing-list-subscriptions",
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
		Path:   "/account/v3/users/" + fmt.Sprint(req.UserID) + "/mailing-list-subscriptions",
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

// SendPhoneValidation: Send a phone number for validation.
func (s *UserAPI) SendPhoneValidation(req *UserAPISendPhoneValidationRequest, opts ...scw.RequestOption) (*PhoneValidation, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/phone-validations",
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

// ConfirmPhoneValidation: Confirm the phone number validation.
func (s *UserAPI) ConfirmPhoneValidation(req *UserAPIConfirmPhoneValidationRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.PhoneValidationID) == "" {
		return errors.New("field PhoneValidationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/phone-validations/" + fmt.Sprint(req.PhoneValidationID) + "/confirm",
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

// GetUserPreferences: Get User Preferences.
func (s *UserAPI) GetUserPreferences(req *UserAPIGetUserPreferencesRequest, opts ...scw.RequestOption) (*UserPreferences, error) {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/users/" + fmt.Sprint(req.UserID) + "/preferences",
	}

	var resp UserPreferences

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateUserPreferences: Update User Preferences.
func (s *UserAPI) UpdateUserPreferences(req *UserAPIUpdateUserPreferencesRequest, opts ...scw.RequestOption) (*UserPreferences, error) {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account/v3/users/" + fmt.Sprint(req.UserID) + "/preferences",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UserPreferences

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

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/users/" + fmt.Sprint(req.UserID) + "/invitations",
		Query:  query,
	}

	var resp ListInvitationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AcceptInvitation: Accept an invitation.
func (s *UserAPI) AcceptInvitation(req *UserAPIAcceptInvitationRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.InvitationID) == "" {
		return errors.New("field InvitationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/invitations/" + fmt.Sprint(req.InvitationID) + "/accept",
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

// RefuseInvitation: Refuse an invitation.
func (s *UserAPI) RefuseInvitation(req *UserAPIRefuseInvitationRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.InvitationID) == "" {
		return errors.New("field InvitationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/invitations/" + fmt.Sprint(req.InvitationID) + "/refuse",
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

// CreateMFAOTP: Create MFA OTP.
func (s *UserAPI) CreateMFAOTP(req *UserAPICreateMFAOTPRequest, opts ...scw.RequestOption) (*MFAOTP, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/mfa/otps",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp MFAOTP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteMFAOTP: Delete MFA OTP.
func (s *UserAPI) DeleteMFAOTP(req *UserAPIDeleteMFAOTPRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.MfaOtpID) == "" {
		return errors.New("field MfaOtpID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account/v3/mfa/otps/" + fmt.Sprint(req.MfaOtpID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListMFAOTPs: List MFA OTPs.
func (s *UserAPI) ListMFAOTPs(req *UserAPIListMFAOTPsRequest, opts ...scw.RequestOption) (*ListMFAOTPsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "user_id", req.UserID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/mfa/otps",
		Query:  query,
	}

	var resp ListMFAOTPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateMFAOTP: Validate MFA OTP.
func (s *UserAPI) ValidateMFAOTP(req *UserAPIValidateMFAOTPRequest, opts ...scw.RequestOption) (*ValidateMFAOTPResponse, error) {
	var err error

	if fmt.Sprint(req.MfaOtpID) == "" {
		return nil, errors.New("field MfaOtpID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/mfa/otps/" + fmt.Sprint(req.MfaOtpID) + "/validate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ValidateMFAOTPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendEmailValidation:
func (s *UserAPI) SendEmailValidation(req *UserAPISendEmailValidationRequest, opts ...scw.RequestOption) (*SendEmailValidationResponse, error) {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/users/" + fmt.Sprint(req.UserID) + "/send-email-validation",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SendEmailValidationResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

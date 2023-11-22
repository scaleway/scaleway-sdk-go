// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package domain_admin provides methods and message types of the domain_admin v2beta1 API.
package domain_admin

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/domain_admin/v2beta1/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/domain_admin/v2beta1/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/domain_admin/v2beta1/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/domain_admin/v2beta1/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/domain_admin/v2beta1/scw"
	std "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/domain_admin/v2beta1/api/std"
	domain_v2beta1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/domain_admin/v2beta1/api/domain/v2beta1"
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

type ContactEmailStatus string

const (
	ContactEmailStatusEmailStatusUnknown = ContactEmailStatus("email_status_unknown")
	ContactEmailStatusValidated          = ContactEmailStatus("validated")
	ContactEmailStatusNotValidated       = ContactEmailStatus("not_validated")
	ContactEmailStatusInvalidEmail       = ContactEmailStatus("invalid_email")
)

func (enum ContactEmailStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "email_status_unknown"
	}
	return string(enum)
}

func (enum ContactEmailStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContactEmailStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContactEmailStatus(ContactEmailStatus(tmp).String())
	return nil
}

type ContactExtensionFRMode string

const (
	ContactExtensionFRModeModeUnknown               = ContactExtensionFRMode("mode_unknown")
	ContactExtensionFRModeIndividual                = ContactExtensionFRMode("individual")
	ContactExtensionFRModeCompanyIdentificationCode = ContactExtensionFRMode("company_identification_code")
	ContactExtensionFRModeDuns                      = ContactExtensionFRMode("duns")
	ContactExtensionFRModeLocal                     = ContactExtensionFRMode("local")
	ContactExtensionFRModeAssociation               = ContactExtensionFRMode("association")
	ContactExtensionFRModeTrademark                 = ContactExtensionFRMode("trademark")
	ContactExtensionFRModeCodeAuthAfnic             = ContactExtensionFRMode("code_auth_afnic")
)

func (enum ContactExtensionFRMode) String() string {
	if enum == "" {
		// return default value if empty
		return "mode_unknown"
	}
	return string(enum)
}

func (enum ContactExtensionFRMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContactExtensionFRMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContactExtensionFRMode(ContactExtensionFRMode(tmp).String())
	return nil
}

type ContactExtensionNLLegalForm string

const (
	ContactExtensionNLLegalFormLegalFormUnknown                      = ContactExtensionNLLegalForm("legal_form_unknown")
	ContactExtensionNLLegalFormOther                                 = ContactExtensionNLLegalForm("other")
	ContactExtensionNLLegalFormNonDutchEuCompany                     = ContactExtensionNLLegalForm("non_dutch_eu_company")
	ContactExtensionNLLegalFormNonDutchLegalFormEnterpriseSubsidiary = ContactExtensionNLLegalForm("non_dutch_legal_form_enterprise_subsidiary")
	ContactExtensionNLLegalFormLimitedCompany                        = ContactExtensionNLLegalForm("limited_company")
	ContactExtensionNLLegalFormLimitedCompanyInFormation             = ContactExtensionNLLegalForm("limited_company_in_formation")
	ContactExtensionNLLegalFormCooperative                           = ContactExtensionNLLegalForm("cooperative")
	ContactExtensionNLLegalFormLimitedPartnership                    = ContactExtensionNLLegalForm("limited_partnership")
	ContactExtensionNLLegalFormSoleCompany                           = ContactExtensionNLLegalForm("sole_company")
	ContactExtensionNLLegalFormEuropeanEconomicInterestGroup         = ContactExtensionNLLegalForm("european_economic_interest_group")
	ContactExtensionNLLegalFormReligiousEntity                       = ContactExtensionNLLegalForm("religious_entity")
	ContactExtensionNLLegalFormPartnership                           = ContactExtensionNLLegalForm("partnership")
	ContactExtensionNLLegalFormPublicCompany                         = ContactExtensionNLLegalForm("public_company")
	ContactExtensionNLLegalFormMutualBenefitCompany                  = ContactExtensionNLLegalForm("mutual_benefit_company")
	ContactExtensionNLLegalFormResidential                           = ContactExtensionNLLegalForm("residential")
	ContactExtensionNLLegalFormShippingCompany                       = ContactExtensionNLLegalForm("shipping_company")
	ContactExtensionNLLegalFormFoundation                            = ContactExtensionNLLegalForm("foundation")
	ContactExtensionNLLegalFormAssociation                           = ContactExtensionNLLegalForm("association")
	ContactExtensionNLLegalFormTradingPartnership                    = ContactExtensionNLLegalForm("trading_partnership")
)

func (enum ContactExtensionNLLegalForm) String() string {
	if enum == "" {
		// return default value if empty
		return "legal_form_unknown"
	}
	return string(enum)
}

func (enum ContactExtensionNLLegalForm) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContactExtensionNLLegalForm) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContactExtensionNLLegalForm(ContactExtensionNLLegalForm(tmp).String())
	return nil
}

type ContactLegalForm string

const (
	ContactLegalFormLegalFormUnknown = ContactLegalForm("legal_form_unknown")
	ContactLegalFormIndividual       = ContactLegalForm("individual")
	ContactLegalFormCorporate        = ContactLegalForm("corporate")
	ContactLegalFormAssociation      = ContactLegalForm("association")
	ContactLegalFormOther            = ContactLegalForm("other")
)

func (enum ContactLegalForm) String() string {
	if enum == "" {
		// return default value if empty
		return "legal_form_unknown"
	}
	return string(enum)
}

func (enum ContactLegalForm) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContactLegalForm) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContactLegalForm(ContactLegalForm(tmp).String())
	return nil
}

type DNSZoneStatus string

const (
	DNSZoneStatusUnknown = DNSZoneStatus("unknown")
	DNSZoneStatusActive  = DNSZoneStatus("active")
	DNSZoneStatusPending = DNSZoneStatus("pending")
	DNSZoneStatusError   = DNSZoneStatus("error")
	DNSZoneStatusLocked  = DNSZoneStatus("locked")
)

func (enum DNSZoneStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DNSZoneStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DNSZoneStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DNSZoneStatus(DNSZoneStatus(tmp).String())
	return nil
}

type DSRecordAlgorithm string

const (
	DSRecordAlgorithmRsamd5           = DSRecordAlgorithm("rsamd5")
	DSRecordAlgorithmDh               = DSRecordAlgorithm("dh")
	DSRecordAlgorithmDsa              = DSRecordAlgorithm("dsa")
	DSRecordAlgorithmRsasha1          = DSRecordAlgorithm("rsasha1")
	DSRecordAlgorithmDsaNsec3Sha1     = DSRecordAlgorithm("dsa_nsec3_sha1")
	DSRecordAlgorithmRsasha1Nsec3Sha1 = DSRecordAlgorithm("rsasha1_nsec3_sha1")
	DSRecordAlgorithmRsasha256        = DSRecordAlgorithm("rsasha256")
	DSRecordAlgorithmRsasha512        = DSRecordAlgorithm("rsasha512")
	DSRecordAlgorithmEccGost          = DSRecordAlgorithm("ecc_gost")
	DSRecordAlgorithmEcdsap256sha256  = DSRecordAlgorithm("ecdsap256sha256")
	DSRecordAlgorithmEcdsap384sha384  = DSRecordAlgorithm("ecdsap384sha384")
	DSRecordAlgorithmEd25519          = DSRecordAlgorithm("ed25519")
	DSRecordAlgorithmEd448            = DSRecordAlgorithm("ed448")
)

func (enum DSRecordAlgorithm) String() string {
	if enum == "" {
		// return default value if empty
		return "rsamd5"
	}
	return string(enum)
}

func (enum DSRecordAlgorithm) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DSRecordAlgorithm) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DSRecordAlgorithm(DSRecordAlgorithm(tmp).String())
	return nil
}

type DSRecordDigestType string

const (
	DSRecordDigestTypeSha1          = DSRecordDigestType("sha_1")
	DSRecordDigestTypeSha256        = DSRecordDigestType("sha_256")
	DSRecordDigestTypeGostR34_11_94 = DSRecordDigestType("gost_r_34_11_94")
	DSRecordDigestTypeSha384        = DSRecordDigestType("sha_384")
)

func (enum DSRecordDigestType) String() string {
	if enum == "" {
		// return default value if empty
		return "sha_1"
	}
	return string(enum)
}

func (enum DSRecordDigestType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DSRecordDigestType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DSRecordDigestType(DSRecordDigestType(tmp).String())
	return nil
}

type DomainFeatureStatus string

const (
	DomainFeatureStatusFeatureStatusUnknown = DomainFeatureStatus("feature_status_unknown")
	DomainFeatureStatusEnabling             = DomainFeatureStatus("enabling")
	DomainFeatureStatusEnabled              = DomainFeatureStatus("enabled")
	DomainFeatureStatusDisabling            = DomainFeatureStatus("disabling")
	DomainFeatureStatusDisabled             = DomainFeatureStatus("disabled")
)

func (enum DomainFeatureStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "feature_status_unknown"
	}
	return string(enum)
}

func (enum DomainFeatureStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainFeatureStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainFeatureStatus(DomainFeatureStatus(tmp).String())
	return nil
}

type DomainRegistrationStatusTransferStatus string

const (
	DomainRegistrationStatusTransferStatusStatusUnknown = DomainRegistrationStatusTransferStatus("status_unknown")
	DomainRegistrationStatusTransferStatusPending       = DomainRegistrationStatusTransferStatus("pending")
	DomainRegistrationStatusTransferStatusWaitingVote   = DomainRegistrationStatusTransferStatus("waiting_vote")
	DomainRegistrationStatusTransferStatusRejected      = DomainRegistrationStatusTransferStatus("rejected")
	DomainRegistrationStatusTransferStatusProcessing    = DomainRegistrationStatusTransferStatus("processing")
	DomainRegistrationStatusTransferStatusDone          = DomainRegistrationStatusTransferStatus("done")
)

func (enum DomainRegistrationStatusTransferStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "status_unknown"
	}
	return string(enum)
}

func (enum DomainRegistrationStatusTransferStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainRegistrationStatusTransferStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainRegistrationStatusTransferStatus(DomainRegistrationStatusTransferStatus(tmp).String())
	return nil
}

type DomainStatus string

const (
	DomainStatusStatusUnknown = DomainStatus("status_unknown")
	DomainStatusActive        = DomainStatus("active")
	DomainStatusCreating      = DomainStatus("creating")
	DomainStatusCreateError   = DomainStatus("create_error")
	DomainStatusRenewing      = DomainStatus("renewing")
	DomainStatusRenewError    = DomainStatus("renew_error")
	DomainStatusXfering       = DomainStatus("xfering")
	DomainStatusXferError     = DomainStatus("xfer_error")
	DomainStatusRestoring     = DomainStatus("restoring")
	DomainStatusRestoreError  = DomainStatus("restore_error")
	DomainStatusExpired       = DomainStatus("expired")
	DomainStatusExpiring      = DomainStatus("expiring")
	DomainStatusUpdating      = DomainStatus("updating")
	DomainStatusChecking      = DomainStatus("checking")
	DomainStatusLocked        = DomainStatus("locked")
	DomainStatusDeleting      = DomainStatus("deleting")
)

func (enum DomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "status_unknown"
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

type ListDNSZoneRecordsRequestOrderBy string

const (
	ListDNSZoneRecordsRequestOrderByNameAsc  = ListDNSZoneRecordsRequestOrderBy("name_asc")
	ListDNSZoneRecordsRequestOrderByNameDesc = ListDNSZoneRecordsRequestOrderBy("name_desc")
)

func (enum ListDNSZoneRecordsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListDNSZoneRecordsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDNSZoneRecordsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDNSZoneRecordsRequestOrderBy(ListDNSZoneRecordsRequestOrderBy(tmp).String())
	return nil
}

type ListDNSZonesRequestOrderBy string

const (
	ListDNSZonesRequestOrderByDomainAsc     = ListDNSZonesRequestOrderBy("domain_asc")
	ListDNSZonesRequestOrderByDomainDesc    = ListDNSZonesRequestOrderBy("domain_desc")
	ListDNSZonesRequestOrderBySubdomainAsc  = ListDNSZonesRequestOrderBy("subdomain_asc")
	ListDNSZonesRequestOrderBySubdomainDesc = ListDNSZonesRequestOrderBy("subdomain_desc")
)

func (enum ListDNSZonesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "domain_asc"
	}
	return string(enum)
}

func (enum ListDNSZonesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDNSZonesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDNSZonesRequestOrderBy(ListDNSZonesRequestOrderBy(tmp).String())
	return nil
}

type ListDomainsRequestOrderBy string

const (
	ListDomainsRequestOrderByDomainAsc      = ListDomainsRequestOrderBy("domain_asc")
	ListDomainsRequestOrderByDomainDesc     = ListDomainsRequestOrderBy("domain_desc")
	ListDomainsRequestOrderByLastUpdateAsc  = ListDomainsRequestOrderBy("last_update_asc")
	ListDomainsRequestOrderByLastUpdateDesc = ListDomainsRequestOrderBy("last_update_desc")
	ListDomainsRequestOrderByCreatedAtAsc   = ListDomainsRequestOrderBy("created_at_asc")
	ListDomainsRequestOrderByCreatedAtDesc  = ListDomainsRequestOrderBy("created_at_desc")
	ListDomainsRequestOrderByUpdatedAtAsc   = ListDomainsRequestOrderBy("updated_at_asc")
	ListDomainsRequestOrderByUpdatedAtDesc  = ListDomainsRequestOrderBy("updated_at_desc")
)

func (enum ListDomainsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "domain_asc"
	}
	return string(enum)
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

type ListTasksRequestOrderBy string

const (
	ListTasksRequestOrderByUpdatedAtAsc  = ListTasksRequestOrderBy("updated_at_asc")
	ListTasksRequestOrderByUpdatedAtDesc = ListTasksRequestOrderBy("updated_at_desc")
)

func (enum ListTasksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "updated_at_asc"
	}
	return string(enum)
}

func (enum ListTasksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTasksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTasksRequestOrderBy(ListTasksRequestOrderBy(tmp).String())
	return nil
}

type RawFormat string

const (
	RawFormatUnknownRawFormat = RawFormat("unknown_raw_format")
	RawFormatNetbox           = RawFormat("netbox")
)

func (enum RawFormat) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_raw_format"
	}
	return string(enum)
}

func (enum RawFormat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RawFormat) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RawFormat(RawFormat(tmp).String())
	return nil
}

type RecordHTTPServiceConfigStrategy string

const (
	RecordHTTPServiceConfigStrategyRandom = RecordHTTPServiceConfigStrategy("random")
	RecordHTTPServiceConfigStrategyHashed = RecordHTTPServiceConfigStrategy("hashed")
)

func (enum RecordHTTPServiceConfigStrategy) String() string {
	if enum == "" {
		// return default value if empty
		return "random"
	}
	return string(enum)
}

func (enum RecordHTTPServiceConfigStrategy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RecordHTTPServiceConfigStrategy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RecordHTTPServiceConfigStrategy(RecordHTTPServiceConfigStrategy(tmp).String())
	return nil
}

type RecordType string

const (
	RecordTypeUnknown = RecordType("unknown")
	RecordTypeA       = RecordType("A")
	RecordTypeAAAA    = RecordType("AAAA")
	RecordTypeCNAME   = RecordType("CNAME")
	RecordTypeTXT     = RecordType("TXT")
	RecordTypeSRV     = RecordType("SRV")
	RecordTypeTLSA    = RecordType("TLSA")
	RecordTypeMX      = RecordType("MX")
	RecordTypeNS      = RecordType("NS")
	RecordTypePTR     = RecordType("PTR")
	RecordTypeCAA     = RecordType("CAA")
	RecordTypeALIAS   = RecordType("ALIAS")
	RecordTypeLOC     = RecordType("LOC")
	RecordTypeSSHFP   = RecordType("SSHFP")
	RecordTypeHINFO   = RecordType("HINFO")
	RecordTypeRP      = RecordType("RP")
	RecordTypeURI     = RecordType("URI")
	RecordTypeDS      = RecordType("DS")
	RecordTypeNAPTR   = RecordType("NAPTR")
	RecordTypeDNAME   = RecordType("DNAME")
)

func (enum RecordType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum RecordType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RecordType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RecordType(RecordType(tmp).String())
	return nil
}

type ReverseIPScope string

const (
	ReverseIPScopeUnknownScope = ReverseIPScope("unknown_scope")
	ReverseIPScopeOnline       = ReverseIPScope("online")
	ReverseIPScopeScaleway     = ReverseIPScope("scaleway")
)

func (enum ReverseIPScope) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_scope"
	}
	return string(enum)
}

func (enum ReverseIPScope) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReverseIPScope) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReverseIPScope(ReverseIPScope(tmp).String())
	return nil
}

type ReverseIPStatus string

const (
	ReverseIPStatusUnknownStatus = ReverseIPStatus("unknown_status")
	ReverseIPStatusPending       = ReverseIPStatus("pending")
	ReverseIPStatusActive        = ReverseIPStatus("active")
	ReverseIPStatusError         = ReverseIPStatus("error")
)

func (enum ReverseIPStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum ReverseIPStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReverseIPStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReverseIPStatus(ReverseIPStatus(tmp).String())
	return nil
}

type SSLCertificateStatus string

const (
	SSLCertificateStatusUnknownStatus = SSLCertificateStatus("unknown_status")
	SSLCertificateStatusNew           = SSLCertificateStatus("new")
	SSLCertificateStatusPending       = SSLCertificateStatus("pending")
	SSLCertificateStatusSuccess       = SSLCertificateStatus("success")
	SSLCertificateStatusError         = SSLCertificateStatus("error")
)

func (enum SSLCertificateStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum SSLCertificateStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SSLCertificateStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SSLCertificateStatus(SSLCertificateStatus(tmp).String())
	return nil
}

// ContactExtensionFRAssociationInfo: contact extension fr association info.
type ContactExtensionFRAssociationInfo struct {
	PublicationJo *time.Time `json:"publication_jo"`

	PublicationJoPage uint32 `json:"publication_jo_page"`
}

// ContactExtensionFRCodeAuthAfnicInfo: contact extension fr code auth afnic info.
type ContactExtensionFRCodeAuthAfnicInfo struct {
	CodeAuthAfnic string `json:"code_auth_afnic"`
}

// ContactExtensionFRDunsInfo: contact extension fr duns info.
type ContactExtensionFRDunsInfo struct {
	DunsID string `json:"duns_id"`

	LocalID string `json:"local_id"`
}

// ContactExtensionFRIndividualInfo: contact extension fr individual info.
type ContactExtensionFRIndividualInfo struct {
	WhoisOptIn bool `json:"whois_opt_in"`
}

// ContactExtensionFRTrademarkInfo: contact extension fr trademark info.
type ContactExtensionFRTrademarkInfo struct {
	TrademarkInpi string `json:"trademark_inpi"`
}

// DSRecordDigest: ds record digest.
type DSRecordDigest struct {
	// Type: default value: sha_1
	Type DSRecordDigestType `json:"type"`

	Digest string `json:"digest"`
}

// DSRecordPublicKey: ds record public key.
type DSRecordPublicKey struct {
	Key string `json:"key"`
}

// RecordGeoIPConfigMatch: record geo ip config match.
type RecordGeoIPConfigMatch struct {
	Countries []string `json:"countries"`

	Continents []string `json:"continents"`

	Data string `json:"data"`
}

// RecordViewConfigView: record view config view.
type RecordViewConfigView struct {
	Subnet string `json:"subnet"`

	Data string `json:"data"`
}

// RecordWeightedConfigWeightedIP: record weighted config weighted ip.
type RecordWeightedConfigWeightedIP struct {
	IP net.IP `json:"ip"`

	Weight uint32 `json:"weight"`
}

// ContactExtensionEU: contact extension eu.
type ContactExtensionEU struct {
	EuropeanCitizenship string `json:"european_citizenship"`
}

// ContactExtensionFR: contact extension fr.
type ContactExtensionFR struct {
	// Mode: default value: mode_unknown
	Mode ContactExtensionFRMode `json:"mode"`

	// Precisely one of IndividualInfo, DunsInfo, AssociationInfo, TrademarkInfo, CodeAuthAfnicInfo must be set.
	IndividualInfo *ContactExtensionFRIndividualInfo `json:"individual_info,omitempty"`

	// Precisely one of IndividualInfo, DunsInfo, AssociationInfo, TrademarkInfo, CodeAuthAfnicInfo must be set.
	DunsInfo *ContactExtensionFRDunsInfo `json:"duns_info,omitempty"`

	// Precisely one of IndividualInfo, DunsInfo, AssociationInfo, TrademarkInfo, CodeAuthAfnicInfo must be set.
	AssociationInfo *ContactExtensionFRAssociationInfo `json:"association_info,omitempty"`

	// Precisely one of IndividualInfo, DunsInfo, AssociationInfo, TrademarkInfo, CodeAuthAfnicInfo must be set.
	TrademarkInfo *ContactExtensionFRTrademarkInfo `json:"trademark_info,omitempty"`

	// Precisely one of IndividualInfo, DunsInfo, AssociationInfo, TrademarkInfo, CodeAuthAfnicInfo must be set.
	CodeAuthAfnicInfo *ContactExtensionFRCodeAuthAfnicInfo `json:"code_auth_afnic_info,omitempty"`
}

// ContactExtensionNL: contact extension nl.
type ContactExtensionNL struct {
	// LegalForm: default value: legal_form_unknown
	LegalForm ContactExtensionNLLegalForm `json:"legal_form"`

	LegalFormRegistrationNumber string `json:"legal_form_registration_number"`
}

// ContactQuestion: contact question.
type ContactQuestion struct {
	Question string `json:"question"`

	Answer string `json:"answer"`
}

// DSRecord: ds record.
type DSRecord struct {
	KeyID uint32 `json:"key_id"`

	// Algorithm: default value: rsamd5
	Algorithm DSRecordAlgorithm `json:"algorithm"`

	// Precisely one of Digest, PublicKey must be set.
	Digest *DSRecordDigest `json:"digest,omitempty"`

	// Precisely one of Digest, PublicKey must be set.
	PublicKey *DSRecordPublicKey `json:"public_key,omitempty"`
}

// RecordGeoIPConfig: record geo ip config.
type RecordGeoIPConfig struct {
	Matches []*RecordGeoIPConfigMatch `json:"matches"`

	Default string `json:"default"`
}

// RecordHTTPServiceConfig: record http service config.
type RecordHTTPServiceConfig struct {
	IPs []net.IP `json:"ips"`

	MustContain *string `json:"must_contain"`

	URL string `json:"url"`

	UserAgent *string `json:"user_agent"`

	// Strategy: default value: random
	Strategy RecordHTTPServiceConfigStrategy `json:"strategy"`
}

// RecordViewConfig: record view config.
type RecordViewConfig struct {
	Views []*RecordViewConfigView `json:"views"`
}

// RecordWeightedConfig: record weighted config.
type RecordWeightedConfig struct {
	WeightedIPs []*RecordWeightedConfigWeightedIP `json:"weighted_ips"`
}

// Contact: contact.
type Contact struct {
	ID string `json:"id"`

	// LegalForm: default value: legal_form_unknown
	LegalForm ContactLegalForm `json:"legal_form"`

	Firstname string `json:"firstname"`

	Lastname string `json:"lastname"`

	CompanyName string `json:"company_name"`

	Email string `json:"email"`

	EmailAlt string `json:"email_alt"`

	PhoneNumber string `json:"phone_number"`

	FaxNumber string `json:"fax_number"`

	AddressLine1 string `json:"address_line_1"`

	AddressLine2 string `json:"address_line_2"`

	Zip string `json:"zip"`

	City string `json:"city"`

	Country string `json:"country"`

	VatIDentificationCode string `json:"vat_identification_code"`

	CompanyIDentificationCode string `json:"company_identification_code"`

	// Lang: default value: unknown_language_code
	Lang std.LanguageCode `json:"lang"`

	Resale bool `json:"resale"`

	// Deprecated
	Questions *[]*ContactQuestion `json:"questions,omitempty"`

	ExtensionFr *ContactExtensionFR `json:"extension_fr"`

	ExtensionEu *ContactExtensionEU `json:"extension_eu"`

	WhoisOptIn bool `json:"whois_opt_in"`

	// EmailStatus: default value: email_status_unknown
	EmailStatus ContactEmailStatus `json:"email_status"`

	State string `json:"state"`

	ExtensionNl *ContactExtensionNL `json:"extension_nl"`
}

// DNSZone: dns zone.
type DNSZone struct {
	Domain string `json:"domain"`

	Subdomain string `json:"subdomain"`

	Ns []string `json:"ns"`

	NsDefault []string `json:"ns_default"`

	NsMaster []string `json:"ns_master"`

	OrganizationID string `json:"organization_id"`

	// Status: default value: unknown
	Status DNSZoneStatus `json:"status"`

	Message *string `json:"message"`

	UpdatedAt *time.Time `json:"updated_at"`

	Internal bool `json:"internal"`

	ProjectID string `json:"project_id"`
}

// DomainDNSSEC: domain dnssec.
type DomainDNSSEC struct {
	// Status: default value: feature_status_unknown
	Status DomainFeatureStatus `json:"status"`

	DsRecords []*DSRecord `json:"ds_records"`
}

// DomainRegistrationStatusExternalDomain: domain registration status external domain.
type DomainRegistrationStatusExternalDomain struct {
	ValidationToken string `json:"validation_token"`
}

// DomainRegistrationStatusTransfer: domain registration status transfer.
type DomainRegistrationStatusTransfer struct {
	// Status: default value: status_unknown
	Status DomainRegistrationStatusTransferStatus `json:"status"`

	VoteCurrentOwner bool `json:"vote_current_owner"`

	VoteNewOwner bool `json:"vote_new_owner"`
}

// Record: record.
type Record struct {
	Data string `json:"data"`

	Name string `json:"name"`

	Priority uint32 `json:"priority"`

	TTL uint32 `json:"ttl"`

	// Type: default value: unknown
	Type RecordType `json:"type"`

	Comment *string `json:"comment"`

	// Precisely one of GeoIPConfig, HTTPServiceConfig, WeightedConfig, ViewConfig must be set.
	GeoIPConfig *RecordGeoIPConfig `json:"geo_ip_config,omitempty"`

	// Precisely one of GeoIPConfig, HTTPServiceConfig, WeightedConfig, ViewConfig must be set.
	HTTPServiceConfig *RecordHTTPServiceConfig `json:"http_service_config,omitempty"`

	// Precisely one of GeoIPConfig, HTTPServiceConfig, WeightedConfig, ViewConfig must be set.
	WeightedConfig *RecordWeightedConfig `json:"weighted_config,omitempty"`

	// Precisely one of GeoIPConfig, HTTPServiceConfig, WeightedConfig, ViewConfig must be set.
	ViewConfig *RecordViewConfig `json:"view_config,omitempty"`

	ID string `json:"id"`
}

// DomainSummary: domain summary.
type DomainSummary struct {
	Domain string `json:"domain"`

	// AutoRenewStatus: default value: feature_status_unknown
	AutoRenewStatus DomainFeatureStatus `json:"auto_renew_status"`

	// DnssecStatus: default value: feature_status_unknown
	DnssecStatus DomainFeatureStatus `json:"dnssec_status"`

	EppCode []string `json:"epp_code"`

	ExpiredAt *time.Time `json:"expired_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	SyncedAt *time.Time `json:"synced_at"`

	Registrar string `json:"registrar"`

	// Status: default value: status_unknown
	Status DomainStatus `json:"status"`

	IsExternal bool `json:"is_external"`

	DNSZonesCount uint32 `json:"dns_zones_count"`

	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`

	CreatedAt *time.Time `json:"created_at"`
}

// InstancesSource: instances source.
type InstancesSource struct {
	Az string `json:"az"`

	URL string `json:"url"`
}

// ReverseIP: reverse ip.
type ReverseIP struct {
	IP net.IP `json:"ip"`

	Hostname string `json:"hostname"`

	// Status: default value: unknown_status
	Status ReverseIPStatus `json:"status"`

	MasterDNSPublic string `json:"master_dns_public"`

	MasterDNSInternal string `json:"master_dns_internal"`

	// Scope: default value: unknown_scope
	Scope ReverseIPScope `json:"scope"`
}

// SSLCertificate: ssl certificate.
type SSLCertificate struct {
	DNSZone string `json:"dns_zone"`

	AlternativeDNSZones []string `json:"alternative_dns_zones"`

	// Status: default value: unknown_status
	Status SSLCertificateStatus `json:"status"`

	PrivateKey string `json:"private_key"`

	CertificateChain string `json:"certificate_chain"`

	CreatedAt *time.Time `json:"created_at"`

	ExpiredAt *time.Time `json:"expired_at"`
}

// Slave: slave.
type Slave struct {
	Configuration string `json:"configuration"`

	Name string `json:"name"`

	Driver string `json:"driver"`

	URI string `json:"uri"`
}

// AddInstancesSourceRequest: add instances source request.
type AddInstancesSourceRequest struct {
	Az string `json:"-"`

	URL string `json:"url"`

	Token string `json:"token"`
}

// AddSlaveRequest: add slave request.
type AddSlaveRequest struct {
	Configuration string `json:"-"`

	Name string `json:"name"`

	Driver string `json:"driver"`

	URI string `json:"uri"`

	Login string `json:"login"`

	Password string `json:"password"`
}

// CancelTaskRequest: cancel task request.
type CancelTaskRequest struct {
	TaskID string `json:"-"`
}

// CancelTaskResponse: cancel task response.
type CancelTaskResponse struct {
}

// CreateSSLCertificateRequest: create ssl certificate request.
type CreateSSLCertificateRequest struct {
	DNSZone string `json:"dns_zone"`

	AlternativeDNSZones []string `json:"alternative_dns_zones"`
}

// DeleteInstancesSourceRequest: delete instances source request.
type DeleteInstancesSourceRequest struct {
	Az string `json:"-"`
}

// DeleteSSLCertificateRequest: delete ssl certificate request.
type DeleteSSLCertificateRequest struct {
	DNSZone string `json:"-"`
}

// DeleteSSLCertificateResponse: delete ssl certificate response.
type DeleteSSLCertificateResponse struct {
}

// DeleteSlaveRequest: delete slave request.
type DeleteSlaveRequest struct {
	Configuration string `json:"-"`

	Name string `json:"-"`
}

// DeleteTaskRequest: delete task request.
type DeleteTaskRequest struct {
	TaskID string `json:"-"`
}

// DeleteTaskResponse: delete task response.
type DeleteTaskResponse struct {
}

// Domain: domain.
type Domain struct {
	Domain string `json:"domain"`

	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`

	// AutoRenewStatus: default value: feature_status_unknown
	AutoRenewStatus DomainFeatureStatus `json:"auto_renew_status"`

	Dnssec *DomainDNSSEC `json:"dnssec"`

	EppCode []string `json:"epp_code"`

	ExpiredAt *time.Time `json:"expired_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	Registrar string `json:"registrar"`

	IsExternal bool `json:"is_external"`

	// Status: default value: status_unknown
	Status DomainStatus `json:"status"`

	DNSZones []*DNSZone `json:"dns_zones"`

	OwnerContact *Contact `json:"owner_contact"`

	TechnicalContact *Contact `json:"technical_contact"`

	AdministrativeContact *Contact `json:"administrative_contact"`

	// Precisely one of ExternalDomainRegistrationStatus, TransferRegistrationStatus must be set.
	ExternalDomainRegistrationStatus *DomainRegistrationStatusExternalDomain `json:"external_domain_registration_status,omitempty"`

	// Precisely one of ExternalDomainRegistrationStatus, TransferRegistrationStatus must be set.
	TransferRegistrationStatus *DomainRegistrationStatusTransfer `json:"transfer_registration_status,omitempty"`
}

// GetDomainRequest: get domain request.
type GetDomainRequest struct {
	Domain string `json:"-"`
}

// GetSSLCertificateRequest: get ssl certificate request.
type GetSSLCertificateRequest struct {
	DNSZone string `json:"-"`
}

// ImportRawDNSZoneRequest: import raw dns zone request.
type ImportRawDNSZoneRequest struct {
	DNSZone string `json:"-"`

	// Format: default value: unknown_raw_format
	Format RawFormat `json:"format"`

	Content string `json:"content"`

	ProjectID string `json:"project_id"`
}

// ImportRawDNSZoneResponse: import raw dns zone response.
type ImportRawDNSZoneResponse struct {
}

// ListDNSZoneRecordsRequest: list dns zone records request.
type ListDNSZoneRecordsRequest struct {
	DNSZone string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListDNSZoneRecordsRequestOrderBy `json:"-"`

	Name string `json:"-"`

	// Type: default value: unknown
	Type RecordType `json:"-"`
}

// ListDNSZoneRecordsResponse: list dns zone records response.
type ListDNSZoneRecordsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Records []*Record `json:"records"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDNSZoneRecordsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDNSZoneRecordsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDNSZoneRecordsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Records = append(r.Records, results.Records...)
	r.TotalCount += uint32(len(results.Records))
	return uint32(len(results.Records)), nil
}

// ListDNSZonesRequest: list dns zones request.
type ListDNSZonesRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: domain_asc
	OrderBy ListDNSZonesRequestOrderBy `json:"-"`

	Domain string `json:"-"`

	DNSZone string `json:"-"`

	ProjectID *string `json:"-"`

	OrganizationID *string `json:"-"`
}

// ListDNSZonesResponse: list dns zones response.
type ListDNSZonesResponse struct {
	TotalCount uint32 `json:"total_count"`

	DNSZones []*DNSZone `json:"dns_zones"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDNSZonesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDNSZonesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDNSZonesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.DNSZones = append(r.DNSZones, results.DNSZones...)
	r.TotalCount += uint32(len(results.DNSZones))
	return uint32(len(results.DNSZones)), nil
}

// ListDomainsRequest: list domains request.
type ListDomainsRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: domain_asc
	OrderBy ListDomainsRequestOrderBy `json:"-"`

	Domain *string `json:"-"`

	Registrar *string `json:"-"`

	// Status: default value: status_unknown
	Status DomainStatus `json:"-"`

	ProjectID *string `json:"-"`

	OrganizationID *string `json:"-"`

	IsExternal *bool `json:"-"`
}

// ListDomainsResponse: list domains response.
type ListDomainsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Domains []*DomainSummary `json:"domains"`
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

// ListInstancesSourcesRequest: list instances sources request.
type ListInstancesSourcesRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListInstancesSourcesResponse: list instances sources response.
type ListInstancesSourcesResponse struct {
	TotalCount uint32 `json:"total_count"`

	Sources []*InstancesSource `json:"sources"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInstancesSourcesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInstancesSourcesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListInstancesSourcesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Sources = append(r.Sources, results.Sources...)
	r.TotalCount += uint32(len(results.Sources))
	return uint32(len(results.Sources)), nil
}

// ListReverseIPsRequest: list reverse i ps request.
type ListReverseIPsRequest struct {
	Searches []string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListReverseIPsResponse: list reverse i ps response.
type ListReverseIPsResponse struct {
	TotalCount uint32 `json:"total_count"`

	ReverseIPs []*ReverseIP `json:"reverse_ips"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListReverseIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListReverseIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListReverseIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ReverseIPs = append(r.ReverseIPs, results.ReverseIPs...)
	r.TotalCount += uint32(len(results.ReverseIPs))
	return uint32(len(results.ReverseIPs)), nil
}

// ListSSLCertificatesRequest: list ssl certificates request.
type ListSSLCertificatesRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	DNSZone string `json:"-"`

	ProjectID *string `json:"-"`
}

// ListSSLCertificatesResponse: list ssl certificates response.
type ListSSLCertificatesResponse struct {
	TotalCount uint32 `json:"total_count"`

	Certificates []*SSLCertificate `json:"certificates"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSSLCertificatesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSSLCertificatesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSSLCertificatesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Certificates = append(r.Certificates, results.Certificates...)
	r.TotalCount += uint32(len(results.Certificates))
	return uint32(len(results.Certificates)), nil
}

// ListSlavesRequest: list slaves request.
type ListSlavesRequest struct {
	Configuration string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListSlavesResponse: list slaves response.
type ListSlavesResponse struct {
	TotalCount uint32 `json:"total_count"`

	Slaves []*Slave `json:"slaves"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSlavesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSlavesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSlavesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Slaves = append(r.Slaves, results.Slaves...)
	r.TotalCount += uint32(len(results.Slaves))
	return uint32(len(results.Slaves)), nil
}

// ListTasksRequest: list tasks request.
type ListTasksRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: updated_at_asc
	OrderBy ListTasksRequestOrderBy `json:"-"`

	Domain *string `json:"-"`

	ProjectID *string `json:"-"`

	// Type: default value: unknown
	Type domain_v2beta1.TaskType `json:"-"`

	// Status: default value: unavailable
	Status domain_v2beta1.TaskStatus `json:"-"`
}

// ListTasksResponse: list tasks response.
type ListTasksResponse struct {
	TotalCount uint32 `json:"total_count"`

	Tasks []*domain_v2beta1.Task `json:"tasks"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTasksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTasksResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTasksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tasks = append(r.Tasks, results.Tasks...)
	r.TotalCount += uint32(len(results.Tasks))
	return uint32(len(results.Tasks)), nil
}

// ResetReverseIPRequest: reset reverse ip request.
type ResetReverseIPRequest struct {
	ReverseIP net.IP `json:"-"`
}

// ResetReverseIPResponse: reset reverse ip response.
type ResetReverseIPResponse struct {
	ReverseIP *ReverseIP `json:"reverse_ip"`
}

// UpdateReverseIPRequest: update reverse ip request.
type UpdateReverseIPRequest struct {
	ReverseIP net.IP `json:"-"`

	Hostname string `json:"hostname"`
}

// Admin management for domains.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ListDomains: Returns a list of all domains.
// You can filter the list by:
//   - Domain name (domain)
//   - Organization ID (organization_id)
//   - Domain status (domain_status)
//   - Registrar (registrar)
//
// You can order the list, ascending and descending, by:
//   - created_at (date of domain order)
//   - updated_at (date of last domain update).
func (s *API) ListDomains(req *ListDomainsRequest, opts ...scw.RequestOption) (*ListDomainsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "domain", req.Domain)
	parameter.AddToQuery(query, "registrar", req.Registrar)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "is_external", req.IsExternal)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-admin/v2beta1/domains",
		Query:  query,
	}

	var resp ListDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomain: Get a domain.
func (s *API) GetDomain(req *GetDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-admin/v2beta1/domains/" + fmt.Sprint(req.Domain) + "",
	}

	var resp Domain

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDNSZones: Returns a list of all DNS zones.
func (s *API) ListDNSZones(req *ListDNSZonesRequest, opts ...scw.RequestOption) (*ListDNSZonesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "domain", req.Domain)
	parameter.AddToQuery(query, "dns_zone", req.DNSZone)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-admin/v2beta1/dns-zones",
		Query:  query,
	}

	var resp ListDNSZonesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDNSZoneRecords: Get records for a DNS zone.
func (s *API) ListDNSZoneRecords(req *ListDNSZoneRecordsRequest, opts ...scw.RequestOption) (*ListDNSZoneRecordsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "type", req.Type)

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-admin/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/records",
		Query:  query,
	}

	var resp ListDNSZoneRecordsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTasks: Returns a list of all tasks.
// You can filter the list by:
//   - Domain name (domain)
//   - Organization ID (organization_id)
//   - Task type (type)
//   - Task status (status).
func (s *API) ListTasks(req *ListTasksRequest, opts ...scw.RequestOption) (*ListTasksResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "domain", req.Domain)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "status", req.Status)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-admin/v2beta1/tasks",
		Query:  query,
	}

	var resp ListTasksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelTask: Cancel a task and rollback all its steps.
// Task types you can cancel with it's required status:
//   - `create_domain`: `error` with step before registrar validation.
func (s *API) CancelTask(req *CancelTaskRequest, opts ...scw.RequestOption) (*CancelTaskResponse, error) {
	var err error

	if fmt.Sprint(req.TaskID) == "" {
		return nil, errors.New("field TaskID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain-admin/v2beta1/tasks/" + fmt.Sprint(req.TaskID) + "/cancel",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CancelTaskResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteTask: Delete a success or in error task from the DB.
func (s *API) DeleteTask(req *DeleteTaskRequest, opts ...scw.RequestOption) (*DeleteTaskResponse, error) {
	var err error

	if fmt.Sprint(req.TaskID) == "" {
		return nil, errors.New("field TaskID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/domain-admin/v2beta1/tasks/" + fmt.Sprint(req.TaskID) + "",
	}

	var resp DeleteTaskResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListReverseIPs: List reverse IP with a list of search values.
// You can search by `IPv4`, `IPv6`, `Hostname` or `Network subnet`.
func (s *API) ListReverseIPs(req *ListReverseIPsRequest, opts ...scw.RequestOption) (*ListReverseIPsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "searches", req.Searches)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-admin/v2beta1/reverse-ips",
		Query:  query,
	}

	var resp ListReverseIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateReverseIP: Replace a reverse IP with the given IPv4 or IPv6 and an Hostname.
func (s *API) UpdateReverseIP(req *UpdateReverseIPRequest, opts ...scw.RequestOption) (*ReverseIP, error) {
	var err error

	if fmt.Sprint(req.ReverseIP) == "" {
		return nil, errors.New("field ReverseIP cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/domain-admin/v2beta1/reverse-ips/" + fmt.Sprint(req.ReverseIP) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ReverseIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResetReverseIP: Reset a reverse IP with the given IPv4 or IPv6 and set the default reverse back.
func (s *API) ResetReverseIP(req *ResetReverseIPRequest, opts ...scw.RequestOption) (*ResetReverseIPResponse, error) {
	var err error

	if fmt.Sprint(req.ReverseIP) == "" {
		return nil, errors.New("field ReverseIP cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain-admin/v2beta1/reverse-ips/" + fmt.Sprint(req.ReverseIP) + "/reset",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ResetReverseIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSlave:
func (s *API) AddSlave(req *AddSlaveRequest, opts ...scw.RequestOption) (*Slave, error) {
	var err error

	if fmt.Sprint(req.Configuration) == "" {
		return nil, errors.New("field Configuration cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain-admin/v2beta1/slaves/" + fmt.Sprint(req.Configuration) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Slave

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSlaves:
func (s *API) ListSlaves(req *ListSlavesRequest, opts ...scw.RequestOption) (*ListSlavesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Configuration) == "" {
		return nil, errors.New("field Configuration cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-admin/v2beta1/slaves/" + fmt.Sprint(req.Configuration) + "",
		Query:  query,
	}

	var resp ListSlavesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSlave:
func (s *API) DeleteSlave(req *DeleteSlaveRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.Configuration) == "" {
		return errors.New("field Configuration cannot be empty in request")
	}

	if fmt.Sprint(req.Name) == "" {
		return errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/domain-admin/v2beta1/slaves/" + fmt.Sprint(req.Configuration) + "/" + fmt.Sprint(req.Name) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// AddInstancesSource:
func (s *API) AddInstancesSource(req *AddInstancesSourceRequest, opts ...scw.RequestOption) (*InstancesSource, error) {
	var err error

	if fmt.Sprint(req.Az) == "" {
		return nil, errors.New("field Az cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain-admin/v2beta1/instances-sources/" + fmt.Sprint(req.Az) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp InstancesSource

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInstancesSources:
func (s *API) ListInstancesSources(req *ListInstancesSourcesRequest, opts ...scw.RequestOption) (*ListInstancesSourcesResponse, error) {
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
		Path:   "/domain-admin/v2beta1/instances-sources",
		Query:  query,
	}

	var resp ListInstancesSourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteInstancesSource:
func (s *API) DeleteInstancesSource(req *DeleteInstancesSourceRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.Az) == "" {
		return errors.New("field Az cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/domain-admin/v2beta1/instances-sources/" + fmt.Sprint(req.Az) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ImportRawDNSZone:
func (s *API) ImportRawDNSZone(req *ImportRawDNSZoneRequest, opts ...scw.RequestOption) (*ImportRawDNSZoneResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain-admin/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/raw",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ImportRawDNSZoneResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSSLCertificate: Get the zone TLS certificate if it exists.
func (s *API) GetSSLCertificate(req *GetSSLCertificateRequest, opts ...scw.RequestOption) (*SSLCertificate, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-admin/v2beta1/ssl-certificates/" + fmt.Sprint(req.DNSZone) + "",
	}

	var resp SSLCertificate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSSLCertificate: Create or return the zone TLS certificate.
func (s *API) CreateSSLCertificate(req *CreateSSLCertificateRequest, opts ...scw.RequestOption) (*SSLCertificate, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain-admin/v2beta1/ssl-certificates",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SSLCertificate

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSSLCertificates: List all user TLS certificates.
func (s *API) ListSSLCertificates(req *ListSSLCertificatesRequest, opts ...scw.RequestOption) (*ListSSLCertificatesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "dns_zone", req.DNSZone)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-admin/v2beta1/ssl-certificates",
		Query:  query,
	}

	var resp ListSSLCertificatesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSSLCertificate: Delete an TLS certificate.
func (s *API) DeleteSSLCertificate(req *DeleteSSLCertificateRequest, opts ...scw.RequestOption) (*DeleteSSLCertificateResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/domain-admin/v2beta1/ssl-certificates/" + fmt.Sprint(req.DNSZone) + "",
	}

	var resp DeleteSSLCertificateResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

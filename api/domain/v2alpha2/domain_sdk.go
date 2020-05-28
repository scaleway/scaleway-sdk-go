// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package domain provides methods and message types of the domain v2alpha2 API.
package domain

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

// API: domain API
//
// Manage your domain DNS zones and records.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// SearchAPI: domain API
//
// Manage your domain zones and records.
type SearchAPI struct {
	client *scw.Client
}

// NewSearchAPI returns a SearchAPI object from a Scaleway client.
func NewSearchAPI(client *scw.Client) *SearchAPI {
	return &SearchAPI{
		client: client,
	}
}

type Civility string

const (
	// CivilityCivilityUnknown is [insert doc].
	CivilityCivilityUnknown = Civility("civility_unknown")
	// CivilityM is [insert doc].
	CivilityM = Civility("m")
	// CivilityMme is [insert doc].
	CivilityMme = Civility("mme")
)

func (enum Civility) String() string {
	if enum == "" {
		// return default value if empty
		return "civility_unknown"
	}
	return string(enum)
}

func (enum Civility) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Civility) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Civility(Civility(tmp).String())
	return nil
}

type ContactEmailStatus string

const (
	// ContactEmailStatusEmailStatusUnknown is [insert doc].
	ContactEmailStatusEmailStatusUnknown = ContactEmailStatus("email_status_unknown")
	// ContactEmailStatusValidated is [insert doc].
	ContactEmailStatusValidated = ContactEmailStatus("validated")
	// ContactEmailStatusNotValidated is [insert doc].
	ContactEmailStatusNotValidated = ContactEmailStatus("not_validated")
	// ContactEmailStatusInvalidEmail is [insert doc].
	ContactEmailStatusInvalidEmail = ContactEmailStatus("invalid_email")
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

type DNSZoneStatus string

const (
	// DNSZoneStatusUnknown is [insert doc].
	DNSZoneStatusUnknown = DNSZoneStatus("unknown")
	// DNSZoneStatusActive is [insert doc].
	DNSZoneStatusActive = DNSZoneStatus("active")
	// DNSZoneStatusPending is [insert doc].
	DNSZoneStatusPending = DNSZoneStatus("pending")
	// DNSZoneStatusError is [insert doc].
	DNSZoneStatusError = DNSZoneStatus("error")
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

type DSRecordDNSSECAlgorithm string

const (
	// DSRecordDNSSECAlgorithmRsamd5 is [insert doc].
	DSRecordDNSSECAlgorithmRsamd5 = DSRecordDNSSECAlgorithm("rsamd5")
	// DSRecordDNSSECAlgorithmDh is [insert doc].
	DSRecordDNSSECAlgorithmDh = DSRecordDNSSECAlgorithm("dh")
	// DSRecordDNSSECAlgorithmDsa is [insert doc].
	DSRecordDNSSECAlgorithmDsa = DSRecordDNSSECAlgorithm("dsa")
	// DSRecordDNSSECAlgorithmRsasha1 is [insert doc].
	DSRecordDNSSECAlgorithmRsasha1 = DSRecordDNSSECAlgorithm("rsasha1")
	// DSRecordDNSSECAlgorithmDsaNsec3Sha1 is [insert doc].
	DSRecordDNSSECAlgorithmDsaNsec3Sha1 = DSRecordDNSSECAlgorithm("dsa_nsec3_sha1")
	// DSRecordDNSSECAlgorithmRsasha1Nsec3Sha1 is [insert doc].
	DSRecordDNSSECAlgorithmRsasha1Nsec3Sha1 = DSRecordDNSSECAlgorithm("rsasha1_nsec3_sha1")
	// DSRecordDNSSECAlgorithmRsasha256 is [insert doc].
	DSRecordDNSSECAlgorithmRsasha256 = DSRecordDNSSECAlgorithm("rsasha256")
	// DSRecordDNSSECAlgorithmRsasha512 is [insert doc].
	DSRecordDNSSECAlgorithmRsasha512 = DSRecordDNSSECAlgorithm("rsasha512")
	// DSRecordDNSSECAlgorithmEccGost is [insert doc].
	DSRecordDNSSECAlgorithmEccGost = DSRecordDNSSECAlgorithm("ecc_gost")
	// DSRecordDNSSECAlgorithmEcdsap256sha256 is [insert doc].
	DSRecordDNSSECAlgorithmEcdsap256sha256 = DSRecordDNSSECAlgorithm("ecdsap256sha256")
	// DSRecordDNSSECAlgorithmEcdsap384sha384 is [insert doc].
	DSRecordDNSSECAlgorithmEcdsap384sha384 = DSRecordDNSSECAlgorithm("ecdsap384sha384")
	// DSRecordDNSSECAlgorithmEd25519 is [insert doc].
	DSRecordDNSSECAlgorithmEd25519 = DSRecordDNSSECAlgorithm("ed25519")
	// DSRecordDNSSECAlgorithmEd448 is [insert doc].
	DSRecordDNSSECAlgorithmEd448 = DSRecordDNSSECAlgorithm("ed448")
)

func (enum DSRecordDNSSECAlgorithm) String() string {
	if enum == "" {
		// return default value if empty
		return "rsamd5"
	}
	return string(enum)
}

func (enum DSRecordDNSSECAlgorithm) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DSRecordDNSSECAlgorithm) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DSRecordDNSSECAlgorithm(DSRecordDNSSECAlgorithm(tmp).String())
	return nil
}

type DigestDNSSECDigestType string

const (
	// DigestDNSSECDigestTypeSha1 is [insert doc].
	DigestDNSSECDigestTypeSha1 = DigestDNSSECDigestType("sha_1")
	// DigestDNSSECDigestTypeSha256 is [insert doc].
	DigestDNSSECDigestTypeSha256 = DigestDNSSECDigestType("sha_256")
	// DigestDNSSECDigestTypeGostR34_11_94 is [insert doc].
	DigestDNSSECDigestTypeGostR34_11_94 = DigestDNSSECDigestType("gost_r_34_11_94")
	// DigestDNSSECDigestTypeSha384 is [insert doc].
	DigestDNSSECDigestTypeSha384 = DigestDNSSECDigestType("sha_384")
)

func (enum DigestDNSSECDigestType) String() string {
	if enum == "" {
		// return default value if empty
		return "sha_1"
	}
	return string(enum)
}

func (enum DigestDNSSECDigestType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DigestDNSSECDigestType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DigestDNSSECDigestType(DigestDNSSECDigestType(tmp).String())
	return nil
}

type DomainFeatureStatus string

const (
	// DomainFeatureStatusUnavailable is [insert doc].
	DomainFeatureStatusUnavailable = DomainFeatureStatus("unavailable")
	// DomainFeatureStatusEnabling is [insert doc].
	DomainFeatureStatusEnabling = DomainFeatureStatus("enabling")
	// DomainFeatureStatusEnabled is [insert doc].
	DomainFeatureStatusEnabled = DomainFeatureStatus("enabled")
	// DomainFeatureStatusDisabling is [insert doc].
	DomainFeatureStatusDisabling = DomainFeatureStatus("disabling")
	// DomainFeatureStatusDisabled is [insert doc].
	DomainFeatureStatusDisabled = DomainFeatureStatus("disabled")
)

func (enum DomainFeatureStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unavailable"
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

type DomainStatus string

const (
	// DomainStatusUnknown is [insert doc].
	DomainStatusUnknown = DomainStatus("unknown")
	// DomainStatusActive is [insert doc].
	DomainStatusActive = DomainStatus("active")
	// DomainStatusToCreate is [insert doc].
	DomainStatusToCreate = DomainStatus("to_create")
	// DomainStatusCreating is [insert doc].
	DomainStatusCreating = DomainStatus("creating")
	// DomainStatusCreateError is [insert doc].
	DomainStatusCreateError = DomainStatus("create_error")
	// DomainStatusToRenew is [insert doc].
	DomainStatusToRenew = DomainStatus("to_renew")
	// DomainStatusRenewing is [insert doc].
	DomainStatusRenewing = DomainStatus("renewing")
	// DomainStatusRenewError is [insert doc].
	DomainStatusRenewError = DomainStatus("renew_error")
	// DomainStatusToXfer is [insert doc].
	DomainStatusToXfer = DomainStatus("to_xfer")
	// DomainStatusXfering is [insert doc].
	DomainStatusXfering = DomainStatus("xfering")
	// DomainStatusXferError is [insert doc].
	DomainStatusXferError = DomainStatus("xfer_error")
	// DomainStatusToRestore is [insert doc].
	DomainStatusToRestore = DomainStatus("to_restore")
	// DomainStatusRestoring is [insert doc].
	DomainStatusRestoring = DomainStatus("restoring")
	// DomainStatusRestoreError is [insert doc].
	DomainStatusRestoreError = DomainStatus("restore_error")
	// DomainStatusToDelete is [insert doc].
	DomainStatusToDelete = DomainStatus("to_delete")
	// DomainStatusExpired is [insert doc].
	DomainStatusExpired = DomainStatus("expired")
	// DomainStatusExpiring is [insert doc].
	DomainStatusExpiring = DomainStatus("expiring")
	// DomainStatusUpdating is [insert doc].
	DomainStatusUpdating = DomainStatus("updating")
	// DomainStatusChecking is [insert doc].
	DomainStatusChecking = DomainStatus("checking")
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

type ExtensionFRMode string

const (
	// ExtensionFRModeModeUnknown is [insert doc].
	ExtensionFRModeModeUnknown = ExtensionFRMode("mode_unknown")
	// ExtensionFRModeParticular is [insert doc].
	ExtensionFRModeParticular = ExtensionFRMode("particular")
	// ExtensionFRModeSiret is [insert doc].
	ExtensionFRModeSiret = ExtensionFRMode("siret")
	// ExtensionFRModeDuns is [insert doc].
	ExtensionFRModeDuns = ExtensionFRMode("duns")
	// ExtensionFRModeLocal is [insert doc].
	ExtensionFRModeLocal = ExtensionFRMode("local")
	// ExtensionFRModeAssociation is [insert doc].
	ExtensionFRModeAssociation = ExtensionFRMode("association")
	// ExtensionFRModeBrand is [insert doc].
	ExtensionFRModeBrand = ExtensionFRMode("brand")
	// ExtensionFRModeCodeAuthAfnic is [insert doc].
	ExtensionFRModeCodeAuthAfnic = ExtensionFRMode("code_auth_afnic")
)

func (enum ExtensionFRMode) String() string {
	if enum == "" {
		// return default value if empty
		return "mode_unknown"
	}
	return string(enum)
}

func (enum ExtensionFRMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ExtensionFRMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ExtensionFRMode(ExtensionFRMode(tmp).String())
	return nil
}

type Lang string

const (
	// LangLangUnknown is [insert doc].
	LangLangUnknown = Lang("lang_unknown")
	// LangFrench is [insert doc].
	LangFrench = Lang("french")
	// LangEnglish is [insert doc].
	LangEnglish = Lang("english")
	// LangSpanish is [insert doc].
	LangSpanish = Lang("spanish")
	// LangGerman is [insert doc].
	LangGerman = Lang("german")
)

func (enum Lang) String() string {
	if enum == "" {
		// return default value if empty
		return "lang_unknown"
	}
	return string(enum)
}

func (enum Lang) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Lang) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Lang(Lang(tmp).String())
	return nil
}

type LegalForm string

const (
	// LegalFormLegalFormUnknown is [insert doc].
	LegalFormLegalFormUnknown = LegalForm("legal_form_unknown")
	// LegalFormParticular is [insert doc].
	LegalFormParticular = LegalForm("particular")
	// LegalFormSociety is [insert doc].
	LegalFormSociety = LegalForm("society")
	// LegalFormAssociation is [insert doc].
	LegalFormAssociation = LegalForm("association")
	// LegalFormOther is [insert doc].
	LegalFormOther = LegalForm("other")
)

func (enum LegalForm) String() string {
	if enum == "" {
		// return default value if empty
		return "legal_form_unknown"
	}
	return string(enum)
}

func (enum LegalForm) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LegalForm) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LegalForm(LegalForm(tmp).String())
	return nil
}

type ListDNSZoneRecordsRequestOrderBy string

const (
	// ListDNSZoneRecordsRequestOrderByNameAsc is [insert doc].
	ListDNSZoneRecordsRequestOrderByNameAsc = ListDNSZoneRecordsRequestOrderBy("name_asc")
	// ListDNSZoneRecordsRequestOrderByNameDesc is [insert doc].
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
	// ListDNSZonesRequestOrderByDomainAsc is [insert doc].
	ListDNSZonesRequestOrderByDomainAsc = ListDNSZonesRequestOrderBy("domain_asc")
	// ListDNSZonesRequestOrderByDomainDesc is [insert doc].
	ListDNSZonesRequestOrderByDomainDesc = ListDNSZonesRequestOrderBy("domain_desc")
	// ListDNSZonesRequestOrderBySubdomainAsc is [insert doc].
	ListDNSZonesRequestOrderBySubdomainAsc = ListDNSZonesRequestOrderBy("subdomain_asc")
	// ListDNSZonesRequestOrderBySubdomainDesc is [insert doc].
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
	// ListDomainsRequestOrderByDomainAsc is [insert doc].
	ListDomainsRequestOrderByDomainAsc = ListDomainsRequestOrderBy("domain_asc")
	// ListDomainsRequestOrderByDomainDesc is [insert doc].
	ListDomainsRequestOrderByDomainDesc = ListDomainsRequestOrderBy("domain_desc")
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

type ListTldsRequestOrderBy string

const (
	// ListTldsRequestOrderByPopularityAsc is [insert doc].
	ListTldsRequestOrderByPopularityAsc = ListTldsRequestOrderBy("popularity_asc")
	// ListTldsRequestOrderByPopularityDesc is [insert doc].
	ListTldsRequestOrderByPopularityDesc = ListTldsRequestOrderBy("popularity_desc")
	// ListTldsRequestOrderByNameAsc is [insert doc].
	ListTldsRequestOrderByNameAsc = ListTldsRequestOrderBy("name_asc")
	// ListTldsRequestOrderByNameDesc is [insert doc].
	ListTldsRequestOrderByNameDesc = ListTldsRequestOrderBy("name_desc")
)

func (enum ListTldsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "popularity_asc"
	}
	return string(enum)
}

func (enum ListTldsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTldsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTldsRequestOrderBy(ListTldsRequestOrderBy(tmp).String())
	return nil
}

type RecordServiceUPConfigStrategy string

const (
	// RecordServiceUPConfigStrategyRandom is [insert doc].
	RecordServiceUPConfigStrategyRandom = RecordServiceUPConfigStrategy("random")
	// RecordServiceUPConfigStrategyHashed is [insert doc].
	RecordServiceUPConfigStrategyHashed = RecordServiceUPConfigStrategy("hashed")
)

func (enum RecordServiceUPConfigStrategy) String() string {
	if enum == "" {
		// return default value if empty
		return "random"
	}
	return string(enum)
}

func (enum RecordServiceUPConfigStrategy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RecordServiceUPConfigStrategy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RecordServiceUPConfigStrategy(RecordServiceUPConfigStrategy(tmp).String())
	return nil
}

type RecordType string

const (
	// RecordTypeUnknown is [insert doc].
	RecordTypeUnknown = RecordType("unknown")
	// RecordTypeA is [insert doc].
	RecordTypeA = RecordType("A")
	// RecordTypeAAAA is [insert doc].
	RecordTypeAAAA = RecordType("AAAA")
	// RecordTypeCNAME is [insert doc].
	RecordTypeCNAME = RecordType("CNAME")
	// RecordTypeTXT is [insert doc].
	RecordTypeTXT = RecordType("TXT")
	// RecordTypeSRV is [insert doc].
	RecordTypeSRV = RecordType("SRV")
	// RecordTypeTLSA is [insert doc].
	RecordTypeTLSA = RecordType("TLSA")
	// RecordTypeMX is [insert doc].
	RecordTypeMX = RecordType("MX")
	// RecordTypeNS is [insert doc].
	RecordTypeNS = RecordType("NS")
	// RecordTypePTR is [insert doc].
	RecordTypePTR = RecordType("PTR")
	// RecordTypeCAA is [insert doc].
	RecordTypeCAA = RecordType("CAA")
	// RecordTypeFUNCMYIPA is [insert doc].
	RecordTypeFUNCMYIPA = RecordType("FUNC_MYIP_A")
	// RecordTypeFUNCMYIPAAAA is [insert doc].
	RecordTypeFUNCMYIPAAAA = RecordType("FUNC_MYIP_AAAA")
	// RecordTypeFUNCURLUPA is [insert doc].
	RecordTypeFUNCURLUPA = RecordType("FUNC_URLUP_A")
	// RecordTypeFUNCURLUPAAAA is [insert doc].
	RecordTypeFUNCURLUPAAAA = RecordType("FUNC_URLUP_AAAA")
	// RecordTypeFUNCPORTUPA is [insert doc].
	RecordTypeFUNCPORTUPA = RecordType("FUNC_PORTUP_A")
	// RecordTypeFUNCPORTUPAAAA is [insert doc].
	RecordTypeFUNCPORTUPAAAA = RecordType("FUNC_PORTUP_AAAA")
	// RecordTypeFUNCVIEWA is [insert doc].
	RecordTypeFUNCVIEWA = RecordType("FUNC_VIEW_A")
	// RecordTypeFUNCVIEWAAAA is [insert doc].
	RecordTypeFUNCVIEWAAAA = RecordType("FUNC_VIEW_AAAA")
	// RecordTypeFUNCVIEWCNAME is [insert doc].
	RecordTypeFUNCVIEWCNAME = RecordType("FUNC_VIEW_CNAME")
	// RecordTypeFUNCVIEWTXT is [insert doc].
	RecordTypeFUNCVIEWTXT = RecordType("FUNC_VIEW_TXT")
	// RecordTypeALIAS is [insert doc].
	RecordTypeALIAS = RecordType("ALIAS")
	// RecordTypeFUNCTION is [insert doc].
	RecordTypeFUNCTION = RecordType("FUNCTION")
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

type TaskStatus string

const (
	// TaskStatusUnavailable is [insert doc].
	TaskStatusUnavailable = TaskStatus("unavailable")
	// TaskStatusNew is [insert doc].
	TaskStatusNew = TaskStatus("new")
	// TaskStatusWaitingPayment is [insert doc].
	TaskStatusWaitingPayment = TaskStatus("waiting_payment")
	// TaskStatusPending is [insert doc].
	TaskStatusPending = TaskStatus("pending")
	// TaskStatusSuccess is [insert doc].
	TaskStatusSuccess = TaskStatus("success")
	// TaskStatusError is [insert doc].
	TaskStatusError = TaskStatus("error")
)

func (enum TaskStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unavailable"
	}
	return string(enum)
}

func (enum TaskStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TaskStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TaskStatus(TaskStatus(tmp).String())
	return nil
}

type TaskType string

const (
	// TaskTypeUnknown is [insert doc].
	TaskTypeUnknown = TaskType("unknown")
	// TaskTypeCreateDomain is [insert doc].
	TaskTypeCreateDomain = TaskType("create_domain")
	// TaskTypeCreateExternalDomain is [insert doc].
	TaskTypeCreateExternalDomain = TaskType("create_external_domain")
	// TaskTypeRenewDomain is [insert doc].
	TaskTypeRenewDomain = TaskType("renew_domain")
	// TaskTypeRestoreDomain is [insert doc].
	TaskTypeRestoreDomain = TaskType("restore_domain")
	// TaskTypeTransferDomain is [insert doc].
	TaskTypeTransferDomain = TaskType("transfer_domain")
	// TaskTypeTradeDomain is [insert doc].
	TaskTypeTradeDomain = TaskType("trade_domain")
	// TaskTypeLockDomainTransfer is [insert doc].
	TaskTypeLockDomainTransfer = TaskType("lock_domain_transfer")
	// TaskTypeUnlockDomainTransfer is [insert doc].
	TaskTypeUnlockDomainTransfer = TaskType("unlock_domain_transfer")
	// TaskTypeEnableDnssec is [insert doc].
	TaskTypeEnableDnssec = TaskType("enable_dnssec")
	// TaskTypeDisableDnssec is [insert doc].
	TaskTypeDisableDnssec = TaskType("disable_dnssec")
	// TaskTypeUpdateDomain is [insert doc].
	TaskTypeUpdateDomain = TaskType("update_domain")
	// TaskTypeUpdateContact is [insert doc].
	TaskTypeUpdateContact = TaskType("update_contact")
	// TaskTypeDeleteDomain is [insert doc].
	TaskTypeDeleteDomain = TaskType("delete_domain")
	// TaskTypeCancelTask is [insert doc].
	TaskTypeCancelTask = TaskType("cancel_task")
	// TaskTypeGenerateSslCertificate is [insert doc].
	TaskTypeGenerateSslCertificate = TaskType("generate_ssl_certificate")
	// TaskTypeRenewSslCertificate is [insert doc].
	TaskTypeRenewSslCertificate = TaskType("renew_ssl_certificate")
	// TaskTypeSendMessage is [insert doc].
	TaskTypeSendMessage = TaskType("send_message")
	// TaskTypeDeleteDomainExpired is [insert doc].
	TaskTypeDeleteDomainExpired = TaskType("delete_domain_expired")
	// TaskTypeDeleteExternalDomain is [insert doc].
	TaskTypeDeleteExternalDomain = TaskType("delete_external_domain")
)

func (enum TaskType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum TaskType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TaskType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TaskType(TaskType(tmp).String())
	return nil
}

type ZoneSSLStatus string

const (
	// ZoneSSLStatusUnknown is [insert doc].
	ZoneSSLStatusUnknown = ZoneSSLStatus("unknown")
	// ZoneSSLStatusNew is [insert doc].
	ZoneSSLStatusNew = ZoneSSLStatus("new")
	// ZoneSSLStatusPending is [insert doc].
	ZoneSSLStatusPending = ZoneSSLStatus("pending")
	// ZoneSSLStatusSuccess is [insert doc].
	ZoneSSLStatusSuccess = ZoneSSLStatus("success")
	// ZoneSSLStatusError is [insert doc].
	ZoneSSLStatusError = ZoneSSLStatus("error")
)

func (enum ZoneSSLStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ZoneSSLStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ZoneSSLStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ZoneSSLStatus(ZoneSSLStatus(tmp).String())
	return nil
}

// AvailableDomain: available domain
type AvailableDomain struct {
	// Domain: domain name
	Domain string `json:"domain"`
	// Available: availability
	Available bool `json:"available"`
	// Creation: price for a creation
	// Precisely one of Creation, Renew, Transfer must be set.
	Creation *scw.Money `json:"creation,omitempty"`
	// Renew: price for a renew
	// Precisely one of Creation, Renew, Transfer must be set.
	Renew *scw.Money `json:"renew,omitempty"`
	// Transfer: price for a transfer
	// Precisely one of Creation, Renew, Transfer must be set.
	Transfer *scw.Money `json:"transfer,omitempty"`
	// Tld: tLD of the domain
	Tld *Tld `json:"tld"`
	// DNSOptions: DNS options availability
	DNSOptions *DNSOptions `json:"dns_options"`
}

// ClearDNSZoneRecordsResponse: clear dns zone records response
type ClearDNSZoneRecordsResponse struct {
}

// Contact: contact
type Contact struct {
	ID string `json:"id"`
	// LegalForm:
	//
	// Default value: legal_form_unknown
	LegalForm LegalForm `json:"legal_form"`
	// Civility:
	//
	// Default value: civility_unknown
	Civility Civility `json:"civility"`

	Firstname string `json:"firstname"`

	Lastname string `json:"lastname"`

	CompanyName string `json:"company_name"`

	Email string `json:"email"`

	EmailAlt string `json:"email_alt"`

	PhoneNumber string `json:"phone_number"`

	FaxNumber string `json:"fax_number"`

	Address1 string `json:"address1"`

	Address2 string `json:"address2"`

	Zip string `json:"zip"`

	City string `json:"city"`

	Country string `json:"country"`

	Vat string `json:"vat"`

	Siret string `json:"siret"`
	// Lang:
	//
	// Default value: lang_unknown
	Lang Lang `json:"lang"`

	Resale bool `json:"resale"`

	Question1 string `json:"question1"`

	Answer1 string `json:"answer1"`

	Question2 string `json:"question2"`

	Answer2 string `json:"answer2"`

	Question3 string `json:"question3"`

	Answer3 string `json:"answer3"`

	Question4 string `json:"question4"`

	Answer4 string `json:"answer4"`

	Question5 string `json:"question5"`

	Answer5 string `json:"answer5"`

	ExtensionFr *ExtensionFR `json:"extension_fr"`

	ExtensionEu *ExtensionEU `json:"extension_eu"`

	WhoisOptOut bool `json:"whois_opt_out"`
	// EmailStatus:
	//
	// Default value: email_status_unknown
	EmailStatus ContactEmailStatus `json:"email_status"`
}

type ContactRoles struct {
	Contact *Contact `json:"contact"`

	Roles map[string]*ContactRolesRoles `json:"roles"`
}

type ContactRolesRoles struct {
	IsOwner bool `json:"is_owner"`

	IsAdminitrative bool `json:"is_adminitrative"`

	IsTechnical bool `json:"is_technical"`
}

// DNSOptions: summary on availables options
type DNSOptions struct {
	BasicInternal *scw.Money `json:"basic_internal"`

	BasicExternal *scw.Money `json:"basic_external"`

	PremiumInternal *scw.Money `json:"premium_internal"`

	PremiumExternal *scw.Money `json:"premium_external"`
}

type DNSZone struct {
	Domain string `json:"domain"`

	Subdomain string `json:"subdomain"`

	Ns []string `json:"ns"`

	NsDefault []string `json:"ns_default"`

	NsMaster []string `json:"ns_master"`
	// Status:
	//
	// Default value: unknown
	Status DNSZoneStatus `json:"status"`

	Message *string `json:"message"`

	UpdatedAt time.Time `json:"updated_at"`

	OrganizationIDs []string `json:"organization_ids"`
}

type DSRecord struct {
	KeyID uint32 `json:"key_id"`
	// Algorithm:
	//
	// Default value: rsamd5
	Algorithm DSRecordDNSSECAlgorithm `json:"algorithm"`

	// Precisely one of Digest, PublicKey must be set.
	Digest *Digest `json:"digest,omitempty"`

	// Precisely one of Digest, PublicKey must be set.
	PublicKey *PublicKey `json:"public_key,omitempty"`
}

// DeleteDNSZoneResponse: delete dns zone response
type DeleteDNSZoneResponse struct {
}

// DeleteExternalDomainResponse: delete external domain response
type DeleteExternalDomainResponse struct {
}

// DeleteSSLCertificateResponse: delete ssl certificate response
type DeleteSSLCertificateResponse struct {
}

type Digest struct {
	// Type:
	//
	// Default value: sha_1
	Type DigestDNSSECDigestType `json:"type"`

	Digest string `json:"digest"`
}

// Domain: domain
type Domain struct {
	Domain string `json:"domain"`

	OrganizationID string `json:"organization_id"`
	// AutoRenewStatus:
	//
	// Default value: unavailable
	AutoRenewStatus DomainFeatureStatus `json:"auto_renew_status"`
	// DnssecStatus:
	//
	// Default value: unavailable
	DnssecStatus DomainFeatureStatus `json:"dnssec_status"`

	DsRecords []*DSRecord `json:"ds_records"`

	Epp []string `json:"epp"`

	ExpiredAt time.Time `json:"expired_at"`

	UpdatedAt time.Time `json:"updated_at"`

	Registrar string `json:"registrar"`

	IsExternal bool `json:"is_external"`
	// Status:
	//
	// Default value: unknown
	Status DomainStatus `json:"status"`

	DNSZones []*DNSZone `json:"dns_zones"`

	DNSZoneCount uint32 `json:"dns_zone_count"`

	TrademarkProtection *DomainTrademarkProtectionConfig `json:"trademark_protection"`

	OwnerContact *Contact `json:"owner_contact"`

	TechnicalContact *Contact `json:"technical_contact"`

	AdministrativeContact *Contact `json:"administrative_contact"`
}

type DomainSummary struct {
	Domain string `json:"domain"`

	OrganizationID string `json:"organization_id"`
	// AutoRenewStatus:
	//
	// Default value: unavailable
	AutoRenewStatus DomainFeatureStatus `json:"auto_renew_status"`
	// DnssecStatus:
	//
	// Default value: unavailable
	DnssecStatus DomainFeatureStatus `json:"dnssec_status"`

	Epp []string `json:"epp"`

	ExpiredAt time.Time `json:"expired_at"`

	UpdatedAt time.Time `json:"updated_at"`

	Registrar string `json:"registrar"`

	IsExternal bool `json:"is_external"`
	// Status:
	//
	// Default value: unknown
	Status DomainStatus `json:"status"`

	DNSZoneCount uint32 `json:"dns_zone_count"`
}

type DomainTrademarkProtectionConfig struct {
}

type ExtensionEU struct {
	EuropeanCitizenship string `json:"european_citizenship"`
}

type ExtensionFR struct {
	// Mode:
	//
	// Default value: mode_unknown
	Mode ExtensionFRMode `json:"mode"`

	// Precisely one of AssociationInfos, BrandInfos, CodeAuthAfnicInfos, DunsInfos, ParticularInfos must be set.
	ParticularInfos *ExtensionFRParticularInfos `json:"particular_infos,omitempty"`

	// Precisely one of AssociationInfos, BrandInfos, CodeAuthAfnicInfos, DunsInfos, ParticularInfos must be set.
	DunsInfos *ExtensionFRDunsInfos `json:"duns_infos,omitempty"`

	// Precisely one of AssociationInfos, BrandInfos, CodeAuthAfnicInfos, DunsInfos, ParticularInfos must be set.
	AssociationInfos *ExtensionFRAssociationInfos `json:"association_infos,omitempty"`

	// Precisely one of AssociationInfos, BrandInfos, CodeAuthAfnicInfos, DunsInfos, ParticularInfos must be set.
	BrandInfos *ExtensionFRBrandInfos `json:"brand_infos,omitempty"`

	// Precisely one of AssociationInfos, BrandInfos, CodeAuthAfnicInfos, DunsInfos, ParticularInfos must be set.
	CodeAuthAfnicInfos *ExtensionFRCodeAuthAfnicInfos `json:"code_auth_afnic_infos,omitempty"`
}

type ExtensionFRAssociationInfos struct {
	PublicationJo time.Time `json:"publication_jo"`

	PublicationJoPage uint32 `json:"publication_jo_page"`
}

type ExtensionFRBrandInfos struct {
	BrandInpi string `json:"brand_inpi"`
}

type ExtensionFRCodeAuthAfnicInfos struct {
	CodeAuthAfnic string `json:"code_auth_afnic"`
}

type ExtensionFRDunsInfos struct {
	DunsID string `json:"duns_id"`

	LocalID string `json:"local_id"`
}

type ExtensionFRParticularInfos struct {
	WhoisOptOut bool `json:"whois_opt_out"`
}

// GetDNSZoneTsigKeyResponse: get dns zone tsig key response
type GetDNSZoneTsigKeyResponse struct {
	ID string `json:"id"`

	Key string `json:"key"`

	Algorithm string `json:"algorithm"`
}

// GetDNSZoneVersionDiffResponse: get dns zone version diff response
type GetDNSZoneVersionDiffResponse struct {
	Changes []*RecordChange `json:"changes"`
}

// GetDomainAuthCodeResponse: get domain auth code response
type GetDomainAuthCodeResponse struct {
	AuthCode string `json:"auth_code"`
}

// GetDomainResponse: get domain response
type GetDomainResponse struct {
	Domain *Domain `json:"domain"`
}

type ImportProviderDNSZoneRequestOnlineV1 struct {
	Token string `json:"token"`
}

// ImportProviderDNSZoneResponse: import provider dns zone response
type ImportProviderDNSZoneResponse struct {
	Records []*Record `json:"records"`
}

// ImportRawDNSZoneResponse: import raw dns zone response
type ImportRawDNSZoneResponse struct {
	Records []*Record `json:"records"`
}

// ListContactsResponse: list contacts response
type ListContactsResponse struct {
	Contacts []*ContactRoles `json:"contacts"`
}

// ListDNSZoneNameserversResponse: list dns zone nameservers response
type ListDNSZoneNameserversResponse struct {
	Ns []*Nameserver `json:"ns"`
}

// ListDNSZoneRecordsResponse: list dns zone records response
type ListDNSZoneRecordsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Records []*Record `json:"records"`
}

// ListDNSZoneVersionRecordsResponse: list dns zone version records response
type ListDNSZoneVersionRecordsResponse struct {
	Records []*Record `json:"records"`
}

// ListDNSZoneVersionsResponse: list dns zone versions response
type ListDNSZoneVersionsResponse struct {
	Versions []*Version `json:"versions"`
}

// ListDNSZonesResponse: list dns zones response
type ListDNSZonesResponse struct {
	TotalCount uint32 `json:"total_count"`

	DNSZones []*DNSZone `json:"dns_zones"`
}

// ListDomainsResponse: list domains response
type ListDomainsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Domains []*DomainSummary `json:"domains"`
}

// ListSSLCertificatesResponse: list ssl certificates response
type ListSSLCertificatesResponse struct {
	Certificates []*ZoneSSL `json:"certificates"`
}

// ListTasksResponse: list tasks response
type ListTasksResponse struct {
	TotalCount uint32 `json:"total_count"`

	Tasks []*Task `json:"tasks"`
}

// ListTldsResponse: list tlds response
type ListTldsResponse struct {
	// TotalCount: number of tlds
	TotalCount uint32 `json:"total_count"`
	// Tlds: array of tlds summaries
	Tlds []*Tld `json:"tlds"`
}

type Nameserver struct {
	Name string `json:"name"`

	IP []string `json:"ip"`
}

type NewContact struct {
	// LegalForm:
	//
	// Default value: legal_form_unknown
	LegalForm LegalForm `json:"legal_form"`
	// Civility:
	//
	// Default value: civility_unknown
	Civility Civility `json:"civility"`

	Firstname string `json:"firstname"`

	Lastname string `json:"lastname"`

	CompanyName string `json:"company_name"`

	Email string `json:"email"`

	EmailAlt string `json:"email_alt"`

	PhoneNumber string `json:"phone_number"`

	FaxNumber string `json:"fax_number"`

	Address1 string `json:"address1"`

	Address2 string `json:"address2"`

	Zip string `json:"zip"`

	City string `json:"city"`

	Country string `json:"country"`

	Vat string `json:"vat"`

	Siret string `json:"siret"`
	// Lang:
	//
	// Default value: lang_unknown
	Lang Lang `json:"lang"`

	Resale bool `json:"resale"`

	Question1 string `json:"question1"`

	Answer1 string `json:"answer1"`

	Question2 string `json:"question2"`

	Answer2 string `json:"answer2"`

	Question3 string `json:"question3"`

	Answer3 string `json:"answer3"`

	Question4 string `json:"question4"`

	Answer4 string `json:"answer4"`

	Question5 string `json:"question5"`

	Answer5 string `json:"answer5"`

	ExtensionFr *ExtensionFR `json:"extension_fr"`

	ExtensionEu *ExtensionEU `json:"extension_eu"`

	WhoisOptOut bool `json:"whois_opt_out"`
}

type PublicKey struct {
	Key string `json:"key"`
}

type Record struct {
	Data string `json:"data"`

	Name string `json:"name"`

	Priority uint32 `json:"priority"`

	TTL uint32 `json:"ttl"`
	// Type:
	//
	// Default value: unknown
	Type RecordType `json:"type"`

	Comment *string `json:"comment"`

	// Precisely one of GeoIPConfig, ServiceUpConfig, ViewConfig, WeightedConfig must be set.
	GeoIPConfig *RecordGeoIPConfig `json:"geo_ip_config,omitempty"`

	// Precisely one of GeoIPConfig, ServiceUpConfig, ViewConfig, WeightedConfig must be set.
	ServiceUpConfig *RecordServiceUPConfig `json:"service_up_config,omitempty"`

	// Precisely one of GeoIPConfig, ServiceUpConfig, ViewConfig, WeightedConfig must be set.
	WeightedConfig *RecordWeightedConfig `json:"weighted_config,omitempty"`

	// Precisely one of GeoIPConfig, ServiceUpConfig, ViewConfig, WeightedConfig must be set.
	ViewConfig *RecordViewConfig `json:"view_config,omitempty"`
}

type RecordChange struct {

	// Precisely one of Add, Clear, Delete, Set must be set.
	Add *RecordChangeAdd `json:"add,omitempty"`

	// Precisely one of Add, Clear, Delete, Set must be set.
	Set *RecordChangeSet `json:"set,omitempty"`

	// Precisely one of Add, Clear, Delete, Set must be set.
	Delete *RecordChangeDelete `json:"delete,omitempty"`

	// Precisely one of Add, Clear, Delete, Set must be set.
	Clear *RecordChangeClear `json:"clear,omitempty"`
}

type RecordChangeAdd struct {
	Records []*Record `json:"records"`
}

type RecordChangeClear struct {
}

type RecordChangeDelete struct {
	Data string `json:"data"`

	Name string `json:"name"`
	// Type:
	//
	// Default value: unknown
	Type RecordType `json:"type"`
}

type RecordChangeSet struct {
	Data string `json:"data"`

	Name string `json:"name"`

	Records []*Record `json:"records"`

	TTL uint32 `json:"ttl"`
	// Type:
	//
	// Default value: unknown
	Type RecordType `json:"type"`
}

type RecordGeoIPConfig struct {
	Filters []*RecordGeoIPConfigFilter `json:"filters"`

	Default string `json:"default"`
}

type RecordGeoIPConfigFilter struct {
	Countries []string `json:"countries"`

	Continents []string `json:"continents"`

	Data string `json:"data"`
}

type RecordServiceUPConfig struct {
	IPs []net.IP `json:"ips"`

	MustContain *string `json:"must_contain"`

	URL string `json:"url"`

	UserAgent *string `json:"user_agent"`
	// Strategy:
	//
	// Default value: random
	Strategy RecordServiceUPConfigStrategy `json:"strategy"`
}

type RecordViewConfig struct {
	Views []*RecordViewConfigView `json:"views"`
}

type RecordViewConfigView struct {
	Subnet string `json:"subnet"`

	Data string `json:"data"`
}

type RecordWeightedConfig struct {
	WeightedIPs []*RecordWeightedConfigWeightedIP `json:"weighted_ips"`
}

type RecordWeightedConfigWeightedIP struct {
	IP net.IP `json:"ip"`

	Weight uint32 `json:"weight"`
}

// RefreshDNSZoneResponse: refresh dns zone response
type RefreshDNSZoneResponse struct {
	DNSZones []*DNSZone `json:"dns_zones"`
}

type RegisterExternalDomainResponse struct {
	Domain string `json:"domain"`

	OrganizationID string `json:"organization_id"`

	ValidationToken string `json:"validation_token"`

	CreatedAt time.Time `json:"created_at"`
}

// RestoreDNSZoneVersionResponse: restore dns zone version response
type RestoreDNSZoneVersionResponse struct {
}

// SearchAvailableDomainsResponse: search available domains response
type SearchAvailableDomainsResponse struct {
	// Domains: array of domains summaries
	Domains []*AvailableDomain `json:"domains"`
}

type Task struct {
	Domain *string `json:"domain"`
	// Type:
	//
	// Default value: unknown
	Type TaskType `json:"type"`
	// Status:
	//
	// Default value: unavailable
	Status TaskStatus `json:"status"`

	StartedAt time.Time `json:"started_at"`

	UpdatedAt time.Time `json:"updated_at"`

	ID string `json:"id"`

	Message string `json:"message"`
}

// Tld: tld
type Tld struct {
	// Name: tLD extension
	Name string `json:"name"`
	// Prices: prices summary
	Prices *TldPrice `json:"prices"`
	// Infos: infos summary
	Infos *TldInfo `json:"infos"`
}

type TldInfo struct {
	Dnssec bool `json:"dnssec"`

	PeriodMin uint32 `json:"period_min"`

	PeriodMax uint32 `json:"period_max"`

	IsDiscount bool `json:"is_discount"`

	RealPrice *TldPrice `json:"real_price"`

	IdnSupport bool `json:"idn_support"`
}

type TldPrice struct {
	Creation *scw.Money `json:"creation"`

	Renew *scw.Money `json:"renew"`

	Transfer *scw.Money `json:"transfer"`

	Trade *scw.Money `json:"trade"`

	Restore *scw.Money `json:"restore"`
}

// UpdateDNSZoneNameserversResponse: update dns zone nameservers response
type UpdateDNSZoneNameserversResponse struct {
	Ns []*Nameserver `json:"ns"`
}

// UpdateDNSZoneRecordsResponse: update dns zone records response
type UpdateDNSZoneRecordsResponse struct {
	Records []*Record `json:"records"`
}

type Version struct {
	CreatedAt time.Time `json:"created_at"`
}

type ZoneSSL struct {
	DNSZone string `json:"dns_zone"`

	AlternativeDNSZones []string `json:"alternative_dns_zones"`
	// Status:
	//
	// Default value: unknown
	Status ZoneSSLStatus `json:"status"`

	PrivateKey []string `json:"private_key"`

	Certificate []string `json:"certificate"`

	CreatedAt time.Time `json:"created_at"`

	ExpiredAt time.Time `json:"expired_at"`
}

// Service API

type ListTasksRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Domain string `json:"-"`
}

// ListTasks: list tasks
//
// List all account tasks.
// You can filter the list by a domain name.
//
func (s *API) ListTasks(req *ListTasksRequest, opts ...scw.RequestOption) (*ListTasksResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "domain", req.Domain)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/tasks",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListTasksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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

type BuyDomainRequest struct {
	Domain string `json:"domain"`

	Period uint32 `json:"period"`

	OrganizationID string `json:"organization_id"`

	// Precisely one of Contact, ContactID must be set.
	ContactID *string `json:"contact_id,omitempty"`

	// Precisely one of Contact, ContactID must be set.
	Contact *NewContact `json:"contact,omitempty"`

	// Precisely one of AdministrativeContact, AdministrativeContactID must be set.
	AdministrativeContactID *string `json:"administrative_contact_id,omitempty"`

	// Precisely one of AdministrativeContact, AdministrativeContactID must be set.
	AdministrativeContact *NewContact `json:"administrative_contact,omitempty"`

	// Precisely one of TechnicalContact, TechnicalContactID must be set.
	TechnicalContactID *string `json:"technical_contact_id,omitempty"`

	// Precisely one of TechnicalContact, TechnicalContactID must be set.
	TechnicalContact *NewContact `json:"technical_contact,omitempty"`
}

// BuyDomain: buy a domain
//
// Request the registration of a domain name.
// You can give an owned domain contact or a full new contact.
//
func (s *API) BuyDomain(req *BuyDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/domains",
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

type RenewDomainRequest struct {
	Domain string `json:"-"`

	Period uint32 `json:"period"`
}

// RenewDomain: renew a domain
//
// Request the renew of a domain name.
//
func (s *API) RenewDomain(req *RenewDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/domains/" + fmt.Sprint(req.Domain) + "/renew",
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

type TransferDomainRequest struct {
	Domain string `json:"domain"`

	AuthCode string `json:"auth_code"`

	OrganizationID string `json:"organization_id"`

	// Precisely one of Contact, ContactID must be set.
	ContactID *string `json:"contact_id,omitempty"`

	// Precisely one of Contact, ContactID must be set.
	Contact *NewContact `json:"contact,omitempty"`

	// Precisely one of AdministrativeContact, AdministrativeContactID must be set.
	AdministrativeContactID *string `json:"administrative_contact_id,omitempty"`

	// Precisely one of AdministrativeContact, AdministrativeContactID must be set.
	AdministrativeContact *NewContact `json:"administrative_contact,omitempty"`

	// Precisely one of TechnicalContact, TechnicalContactID must be set.
	TechnicalContactID *string `json:"technical_contact_id,omitempty"`

	// Precisely one of TechnicalContact, TechnicalContactID must be set.
	TechnicalContact *NewContact `json:"technical_contact,omitempty"`
}

// TransferDomain: transfer a domain
//
// Request the transfer from an other registrar domain to Scaleway.
//
func (s *API) TransferDomain(req *TransferDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/domains/transfer",
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

type TradeDomainRequest struct {
	Domain string `json:"-"`

	OrganizationID *string `json:"organization_id"`

	// Precisely one of Contact, ContactID must be set.
	ContactID *string `json:"contact_id,omitempty"`

	// Precisely one of Contact, ContactID must be set.
	Contact *NewContact `json:"contact,omitempty"`
}

// TradeDomain: trade a domain contact
//
// Request a trade for the contact owner.<br/>
// If an `organization_id` is given, the change is a from the current Scaleway account to another Scaleway account.<br/>
// If no contact is given, the first contact of the other Scaleway account is taken.<br/>
// If the other Scaleway account has no contact. An error occurred.
//
func (s *API) TradeDomain(req *TradeDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/domains/" + fmt.Sprint(req.Domain) + "/trade",
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

type RegisterExternalDomainRequest struct {
	Domain string `json:"domain"`

	OrganizationID string `json:"organization_id"`
}

// RegisterExternalDomain: register an external domain
//
// Request the register of an external domain name.
//
func (s *API) RegisterExternalDomain(req *RegisterExternalDomainRequest, opts ...scw.RequestOption) (*RegisterExternalDomainResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/domains/external-domain",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RegisterExternalDomainResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteExternalDomainRequest struct {
	Domain string `json:"-"`
}

// DeleteExternalDomain: delete an external domain
//
// Delete an external domain name.
//
func (s *API) DeleteExternalDomain(req *DeleteExternalDomainRequest, opts ...scw.RequestOption) (*DeleteExternalDomainResponse, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/domain/v2alpha2/domains/external-domain/" + fmt.Sprint(req.Domain) + "",
		Headers: http.Header{},
	}

	var resp DeleteExternalDomainResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListContactsRequest struct {
	Domain *string `json:"-"`
}

// ListContacts: list contacts
//
// Return a list of contacts with their domains and roles.
// You can filter the list by a domain name.
//
func (s *API) ListContacts(req *ListContactsRequest, opts ...scw.RequestOption) (*ListContactsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "domain", req.Domain)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/contacts",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListContactsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetContactRequest struct {
	ContactID string `json:"-"`
}

// GetContact: get a contact
//
// Return a contact details retrieve from the registrar by a given contact ID.
func (s *API) GetContact(req *GetContactRequest, opts ...scw.RequestOption) (*Contact, error) {
	var err error

	if fmt.Sprint(req.ContactID) == "" {
		return nil, errors.New("field ContactID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/contacts/" + fmt.Sprint(req.ContactID) + "",
		Headers: http.Header{},
	}

	var resp Contact

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateContactRequest struct {
	ContactID string `json:"-"`

	Email *string `json:"email"`

	EmailAlt *string `json:"email_alt"`

	PhoneNumber *string `json:"phone_number"`

	FaxNumber *string `json:"fax_number"`

	Address1 *string `json:"address1"`

	Address2 *string `json:"address2"`

	Zip *string `json:"zip"`

	City *string `json:"city"`

	Country *string `json:"country"`

	Vat *string `json:"vat"`

	Siret *string `json:"siret"`
	// Lang:
	//
	// Default value: lang_unknown
	Lang Lang `json:"lang"`

	Resale *bool `json:"resale"`

	Question1 *string `json:"question1"`

	Answer1 *string `json:"answer1"`

	Question2 *string `json:"question2"`

	Answer2 *string `json:"answer2"`

	Question3 *string `json:"question3"`

	Answer3 *string `json:"answer3"`

	Question4 *string `json:"question4"`

	Answer4 *string `json:"answer4"`

	Question5 *string `json:"question5"`

	Answer5 *string `json:"answer5"`

	ExtensionFr *ExtensionFR `json:"extension_fr"`

	ExtensionEu *ExtensionEU `json:"extension_eu"`

	WhoisOptOut *bool `json:"whois_opt_out"`
}

// UpdateContact: update contact
//
// You can edit the contact coordinates.
func (s *API) UpdateContact(req *UpdateContactRequest, opts ...scw.RequestOption) (*Contact, error) {
	var err error

	if fmt.Sprint(req.ContactID) == "" {
		return nil, errors.New("field ContactID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/domain/v2alpha2/contacts/" + fmt.Sprint(req.ContactID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Contact

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListDomainsRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: domain_asc
	OrderBy ListDomainsRequestOrderBy `json:"-"`

	Registrar string `json:"-"`
	// Status:
	//
	// Default value: unknown
	Status DomainStatus `json:"-"`
}

// ListDomains: list domains
//
// Returns a list of domains owned by the user.
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
	parameter.AddToQuery(query, "registrar", req.Registrar)
	parameter.AddToQuery(query, "status", req.Status)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/domains",
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

type GetDomainRequest struct {
	Domain string `json:"-"`
}

// GetDomain: get domain
//
// Returns a the domain with more informations.
func (s *API) GetDomain(req *GetDomainRequest, opts ...scw.RequestOption) (*GetDomainResponse, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/domains/" + fmt.Sprint(req.Domain) + "",
		Headers: http.Header{},
	}

	var resp GetDomainResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateDomainRequest struct {
	Domain string `json:"-"`

	// Precisely one of TechnicalContact, TechnicalContactID must be set.
	TechnicalContactID *string `json:"technical_contact_id,omitempty"`

	// Precisely one of TechnicalContact, TechnicalContactID must be set.
	TechnicalContact *NewContact `json:"technical_contact,omitempty"`

	// Precisely one of OwnerContact, OwnerContactID must be set.
	OwnerContactID *string `json:"owner_contact_id,omitempty"`

	// Precisely one of OwnerContact, OwnerContactID must be set.
	OwnerContact *NewContact `json:"owner_contact,omitempty"`

	// Precisely one of AdministrativeContact, AdministrativeContactID must be set.
	AdministrativeContactID *string `json:"administrative_contact_id,omitempty"`

	// Precisely one of AdministrativeContact, AdministrativeContactID must be set.
	AdministrativeContact *NewContact `json:"administrative_contact,omitempty"`
}

// UpdateDomain: update a domain
//
// Update the domain contacts or create a new one.<br/>
// If you add the same contact for multiple roles. Only one ID will be created and use for all of them.
//
func (s *API) UpdateDomain(req *UpdateDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/domain/v2alpha2/domains/" + fmt.Sprint(req.Domain) + "",
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

type LockDomainTransferRequest struct {
	Domain string `json:"-"`
}

// LockDomainTransfer: lock domain transfer
//
// Lock domain transfer, a lock domain transfer can't be transferred and the auth code can't be requested.
//
func (s *API) LockDomainTransfer(req *LockDomainTransferRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/domains/" + fmt.Sprint(req.Domain) + "/lock-transfer",
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

type UnlockDomainTransferRequest struct {
	Domain string `json:"-"`
}

// UnlockDomainTransfer: unlock domain transfer
//
// Unlock domain transfer, an unlock domain transfer can be transferred and the auth code can be requested.
//
func (s *API) UnlockDomainTransfer(req *UnlockDomainTransferRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/domains/" + fmt.Sprint(req.Domain) + "/unlock-transfer",
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

type EnableDomainAutoRenewRequest struct {
	Domain string `json:"-"`
}

// EnableDomainAutoRenew: enable domain auto renew
func (s *API) EnableDomainAutoRenew(req *EnableDomainAutoRenewRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/domains/" + fmt.Sprint(req.Domain) + "/enable-auto-renew",
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

type DisableDomainAutoRenewRequest struct {
	Domain string `json:"-"`
}

// DisableDomainAutoRenew: disable domain auto renew
func (s *API) DisableDomainAutoRenew(req *DisableDomainAutoRenewRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/domains/" + fmt.Sprint(req.Domain) + "/disable-auto-renew",
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

type GetDomainAuthCodeRequest struct {
	Domain string `json:"-"`
}

// GetDomainAuthCode: return domain auth code
//
// If possible, return the auth code for an unlock domain transfer, or an error if the domain is locked.
// Some TLD may have a different procedure to retrieve the auth code, in that case, the informations is given in the message field.
//
func (s *API) GetDomainAuthCode(req *GetDomainAuthCodeRequest, opts ...scw.RequestOption) (*GetDomainAuthCodeResponse, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/domains/" + fmt.Sprint(req.Domain) + "/auth-code",
		Headers: http.Header{},
	}

	var resp GetDomainAuthCodeResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EnableDomainDNSSECRequest struct {
	Domain string `json:"-"`

	DsRecord *DSRecord `json:"ds_record"`
}

// EnableDomainDNSSEC: update domain DNSSEC
//
// If your domain has the default Scaleway NS and uses another registrar, you have to update the DS record manually.
// For the algorithm, here the code number for each type:
//   - 1: RSAMD5
//   - 2: DIFFIE_HELLMAN
//   - 3: DSA_SHA1
//   - 5: RSA_SHA1
//   - 6: DSA_NSEC3_SHA1
//   - 7: RSASHA1_NSEC3_SHA1
//   - 8: RSASHA256
//   - 10: RSASHA512
//   - 12: ECC_GOST
//   - 13: ECDSAP256SHA256
//   - 14: ECDSAP384SHA384
//
// And for the digest type:
//   - 1: SHA_1
//   - 2: SHA_256
//   - 3: GOST_R_34_11_94
//   - 4: SHA_384
//
func (s *API) EnableDomainDNSSEC(req *EnableDomainDNSSECRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/domains/" + fmt.Sprint(req.Domain) + "/enable-dnssec",
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

type DisableDomainDNSSECRequest struct {
	Domain string `json:"-"`
}

// DisableDomainDNSSEC: disable domain DNSSEC
func (s *API) DisableDomainDNSSEC(req *DisableDomainDNSSECRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/domains/" + fmt.Sprint(req.Domain) + "/disable-dnssec",
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

type ListDNSZonesRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: domain_asc
	OrderBy ListDNSZonesRequestOrderBy `json:"-"`

	Domain string `json:"-"`

	DNSZone string `json:"-"`
}

// ListDNSZones: list DNS zones
//
// Returns a list of manageable DNS zones.
// You can filter the DNS zones by a domain name.
//
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

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/dns-zones",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListDNSZonesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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

type CreateDNSZoneRequest struct {
	Domain string `json:"domain"`

	Subdomain string `json:"subdomain"`

	OrganizationIDs []string `json:"organization_ids"`
}

// CreateDNSZone: create a DNS zone
//
// Create a new DNS zone.
func (s *API) CreateDNSZone(req *CreateDNSZoneRequest, opts ...scw.RequestOption) (*DNSZone, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/dns-zones",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DNSZone

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateDNSZoneRequest struct {
	DNSZone string `json:"-"`

	NewDNSZone *string `json:"new_dns_zone"`

	OrganizationIDs *[]string `json:"organization_ids"`
}

// UpdateDNSZone: update a DNS zone
//
// Update the name and/or the organizations for a DNS zone.
func (s *API) UpdateDNSZone(req *UpdateDNSZoneRequest, opts ...scw.RequestOption) (*DNSZone, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DNSZone

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CopyDNSZoneRequest struct {
	DNSZone string `json:"-"`

	NewDNSZone *string `json:"new_dns_zone"`

	Overwrite *bool `json:"overwrite"`
}

// CopyDNSZone: copy a DNS zone
//
// Copy an existed DNS zone with all its records into a new one.
func (s *API) CopyDNSZone(req *CopyDNSZoneRequest, opts ...scw.RequestOption) (*DNSZone, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/copy",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DNSZone

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteDNSZoneRequest struct {
	DNSZone string `json:"-"`
}

// DeleteDNSZone: delete DNS zone
//
// Delete a DNS zone and all it's records.
func (s *API) DeleteDNSZone(req *DeleteDNSZoneRequest, opts ...scw.RequestOption) (*DeleteDNSZoneResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "",
		Headers: http.Header{},
	}

	var resp DeleteDNSZoneResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListDNSZoneRecordsRequest struct {
	DNSZone string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: name_asc
	OrderBy ListDNSZoneRecordsRequestOrderBy `json:"-"`

	Name string `json:"-"`
	// Type:
	//
	// Default value: unknown
	Type RecordType `json:"-"`
}

// ListDNSZoneRecords: list DNS zone records
//
// Returns a list of DNS records about a DNS zone with default NS.
// You can filter the records by a type and a name.
//
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
		Method:  "GET",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/records",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListDNSZoneRecordsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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

type UpdateDNSZoneRecordsRequest struct {
	DNSZone string `json:"-"`

	Changes []*RecordChange `json:"changes"`

	ReturnAllRecords *bool `json:"return_all_records"`
}

// UpdateDNSZoneRecords: update DNS zone records
//
// Only available with default NS.<br/>
// Send a list of actions and records.
//
// Action can be:
//  - add:
//   - Add new record
//   - Can be more specific and add a new IP to an existing A record for example
//  - set:
//   - Edit a record
//   - Can be more specific and edit an IP from an existing A record for example
//  - delete:
//   - Delete a record
//   - Can be more specific and delete an IP from an existing A record for example
//  - clear:
//   - Delete all records from a DNS zone
//
// Any edition will be versioned.
//
func (s *API) UpdateDNSZoneRecords(req *UpdateDNSZoneRecordsRequest, opts ...scw.RequestOption) (*UpdateDNSZoneRecordsResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/records",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateDNSZoneRecordsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListDNSZoneNameserversRequest struct {
	DNSZone string `json:"-"`
}

// ListDNSZoneNameservers: list DNS zone nameservers
//
// Returns a list of Nameservers and there optionnal glue records for a DNS zone.
func (s *API) ListDNSZoneNameservers(req *ListDNSZoneNameserversRequest, opts ...scw.RequestOption) (*ListDNSZoneNameserversResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/nameservers",
		Headers: http.Header{},
	}

	var resp ListDNSZoneNameserversResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateDNSZoneNameserversRequest struct {
	DNSZone string `json:"-"`

	Ns []*Nameserver `json:"ns"`
}

// UpdateDNSZoneNameservers: update DNS zone nameservers
//
// Update DNS zone nameservers and set optionnal glue records.
func (s *API) UpdateDNSZoneNameservers(req *UpdateDNSZoneNameserversRequest, opts ...scw.RequestOption) (*UpdateDNSZoneNameserversResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/nameservers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateDNSZoneNameserversResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ClearDNSZoneRecordsRequest struct {
	DNSZone string `json:"-"`
}

// ClearDNSZoneRecords: clear DNS zone records
//
// Only available with default NS.<br/>
// Delete all the records from a DNS zone.
// Any edition will be versioned.
//
func (s *API) ClearDNSZoneRecords(req *ClearDNSZoneRecordsRequest, opts ...scw.RequestOption) (*ClearDNSZoneRecordsResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/records",
		Headers: http.Header{},
	}

	var resp ClearDNSZoneRecordsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ExportRawDNSZoneRequest struct {
	DNSZone string `json:"-"`

	Format string `json:"-"`
}

// ExportRawDNSZone: export raw DNS zone
//
// Get DNS zone from a given format with default NS.
func (s *API) ExportRawDNSZone(req *ExportRawDNSZoneRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "format", req.Format)

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/raw",
		Query:   query,
		Headers: http.Header{},
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ImportRawDNSZoneRequest struct {
	DNSZone string `json:"-"`

	Format string `json:"format"`

	Content string `json:"content"`
}

// ImportRawDNSZone: import raw DNS zone
//
// Import and replace records from a given provider format with default NS.
func (s *API) ImportRawDNSZone(req *ImportRawDNSZoneRequest, opts ...scw.RequestOption) (*ImportRawDNSZoneResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/raw",
		Headers: http.Header{},
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

type ImportProviderDNSZoneRequest struct {
	DNSZone string `json:"-"`

	// Precisely one of OnlineV1 must be set.
	OnlineV1 *ImportProviderDNSZoneRequestOnlineV1 `json:"online_v1,omitempty"`
}

// ImportProviderDNSZone: import provider DNS zone
//
// Import and replace records from a given provider format with default NS.
func (s *API) ImportProviderDNSZone(req *ImportProviderDNSZoneRequest, opts ...scw.RequestOption) (*ImportProviderDNSZoneResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/import-provider",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ImportProviderDNSZoneResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RefreshDNSZoneRequest struct {
	DNSZone string `json:"-"`

	RecreateDNSZone bool `json:"recreate_dns_zone"`

	RecreateSubDNSZone bool `json:"recreate_sub_dns_zone"`
}

// RefreshDNSZone: refresh DNS zone
//
// Refresh SOA DNS zone.
// You can recreate the given DNS zone and it's sub DNS zone if needed.
//
func (s *API) RefreshDNSZone(req *RefreshDNSZoneRequest, opts ...scw.RequestOption) (*RefreshDNSZoneResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/refresh",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RefreshDNSZoneResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListDNSZoneVersionsRequest struct {
	DNSZone string `json:"-"`
}

// ListDNSZoneVersions: list DNS zone versions
//
// Get a list of DNS zone versions.<br/>
// You are limited to 100 versions.<br/>
// Beyound that, the most older version will be erased after each new modification.
//
func (s *API) ListDNSZoneVersions(req *ListDNSZoneVersionsRequest, opts ...scw.RequestOption) (*ListDNSZoneVersionsResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/versions",
		Headers: http.Header{},
	}

	var resp ListDNSZoneVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListDNSZoneVersionRecordsRequest struct {
	DNSZone string `json:"-"`

	Version string `json:"-"`
}

// ListDNSZoneVersionRecords: list DNS zone version records
//
// Get a list of records from an old DNS zone version.
func (s *API) ListDNSZoneVersionRecords(req *ListDNSZoneVersionRecordsRequest, opts ...scw.RequestOption) (*ListDNSZoneVersionRecordsResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	if fmt.Sprint(req.Version) == "" {
		return nil, errors.New("field Version cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/versions/" + fmt.Sprint(req.Version) + "/records",
		Headers: http.Header{},
	}

	var resp ListDNSZoneVersionRecordsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDNSZoneVersionDiffRequest struct {
	DNSZone string `json:"-"`

	Version string `json:"-"`
}

// GetDNSZoneVersionDiff: get DNS zone version diff
//
// Get all differences from an old DNS zone version.
func (s *API) GetDNSZoneVersionDiff(req *GetDNSZoneVersionDiffRequest, opts ...scw.RequestOption) (*GetDNSZoneVersionDiffResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	if fmt.Sprint(req.Version) == "" {
		return nil, errors.New("field Version cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/versions/" + fmt.Sprint(req.Version) + "/diff",
		Headers: http.Header{},
	}

	var resp GetDNSZoneVersionDiffResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RestoreDNSZoneVersionRequest struct {
	DNSZone string `json:"-"`

	Version string `json:"-"`
}

// RestoreDNSZoneVersion: restore DNS zone version
//
// Restore and active old DNS zone version.
func (s *API) RestoreDNSZoneVersion(req *RestoreDNSZoneVersionRequest, opts ...scw.RequestOption) (*RestoreDNSZoneVersionResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	if fmt.Sprint(req.Version) == "" {
		return nil, errors.New("field Version cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/versions/" + fmt.Sprint(req.Version) + "/restore",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RestoreDNSZoneVersionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateSSLCertificateRequest struct {
	DNSZone string `json:"dns_zone"`

	AlternativeDNSZones []string `json:"alternative_dns_zones"`
}

// CreateSSLCertificate: create a new SSL certificate
func (s *API) CreateSSLCertificate(req *CreateSSLCertificateRequest, opts ...scw.RequestOption) (*ZoneSSL, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/domain/v2alpha2/ssl-certificates",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ZoneSSL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListSSLCertificatesRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	DNSZone string `json:"-"`
}

// ListSSLCertificates: list all user SSL certificate
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

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/ssl-certificates",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListSSLCertificatesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteSSLCertificateRequest struct {
	DNSZone string `json:"-"`
}

// DeleteSSLCertificate: delete an SSL certificate
func (s *API) DeleteSSLCertificate(req *DeleteSSLCertificateRequest, opts ...scw.RequestOption) (*DeleteSSLCertificateResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/domain/v2alpha2/ssl-certificates/" + fmt.Sprint(req.DNSZone) + "",
		Headers: http.Header{},
	}

	var resp DeleteSSLCertificateResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDNSZoneTsigKeyRequest struct {
	DNSZone string `json:"-"`
}

// GetDNSZoneTsigKey: get the DNS zone TSIG Key
//
// Get the DNS zone TSIG Key to allow AXFR request.
func (s *API) GetDNSZoneTsigKey(req *GetDNSZoneTsigKeyRequest, opts ...scw.RequestOption) (*GetDNSZoneTsigKeyResponse, error) {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/tsig-key",
		Headers: http.Header{},
	}

	var resp GetDNSZoneTsigKeyResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteDNSZoneTsigKeyRequest struct {
	DNSZone string `json:"-"`

	ID string `json:"-"`
}

// DeleteDNSZoneTsigKey: delete the DNS zone TSIG Key
func (s *API) DeleteDNSZoneTsigKey(req *DeleteDNSZoneTsigKeyRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.DNSZone) == "" {
		return errors.New("field DNSZone cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/domain/v2alpha2/dns-zones/" + fmt.Sprint(req.DNSZone) + "/tsig-key/" + fmt.Sprint(req.ID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Service SearchAPI

type SearchAPIListTldsRequest struct {
	// Page: page required
	Page *int32 `json:"-"`
	// PageSize: number of tlds by page
	PageSize *uint32 `json:"-"`
	// OrderBy: field to use to order
	//
	// Default value: popularity_asc
	OrderBy ListTldsRequestOrderBy `json:"-"`
	// Tlds: array of tlds to search on
	Tlds []string `json:"-"`
}

// ListTlds: list available TLDs
//
// Returns a list of all Top-Level-Domain available.
func (s *SearchAPI) ListTlds(req *SearchAPIListTldsRequest, opts ...scw.RequestOption) (*ListTldsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "tlds", req.Tlds)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/tlds",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListTldsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTldsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTldsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTldsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tlds = append(r.Tlds, results.Tlds...)
	r.TotalCount += uint32(len(results.Tlds))
	return uint32(len(results.Tlds)), nil
}

type SearchAPISearchAvailableDomainsRequest struct {
	// Search: string to search on
	Search string `json:"-"`
	// Tlds: array of tlds to search on
	Tlds []string `json:"-"`
}

// SearchAvailableDomains: search available domains
//
// Search a domain (you can search multiple domains by separate them with a comma).
// With their corresponding price:
//  - `creation` price for a new domain.
//  - `renew` price for an owned domain.
//  - `transfer` price for a domain not owned on this registrar.
//
// If the TLD list is empty or not set. The search returns the results from the most popular TLDs.
//
func (s *SearchAPI) SearchAvailableDomains(req *SearchAPISearchAvailableDomainsRequest, opts ...scw.RequestOption) (*SearchAvailableDomainsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "search", req.Search)
	parameter.AddToQuery(query, "tlds", req.Tlds)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/domain/v2alpha2/available-domains",
		Query:   query,
		Headers: http.Header{},
	}

	var resp SearchAvailableDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

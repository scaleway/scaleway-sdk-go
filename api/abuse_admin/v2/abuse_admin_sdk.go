// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package abuse_admin provides methods and message types of the abuse_admin v2 API.
package abuse_admin

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
	trustandsafety_private_v1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/trustandsafety_private/v1"
	resource_v1alpha1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/resource/v1alpha1"
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

type CaseClass string

const (
	// Unknown class.
	CaseClassUnknownClass = CaseClass("unknown_class")
	// The family of complaint type is activity.
	CaseClassActivity = CaseClass("activity")
	// The family of complaint type is content.
	CaseClassContent = CaseClass("content")
)

func (enum CaseClass) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_class"
	}
	return string(enum)
}

func (enum CaseClass) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CaseClass) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CaseClass(CaseClass(tmp).String())
	return nil
}

type CaseCloseReason string

const (
	// Unknown close reason.
	CaseCloseReasonUnknownCloseReason = CaseCloseReason("unknown_close_reason")
	// Case has been closed manually.
	CaseCloseReasonClosedManually = CaseCloseReason("closed_manually")
	// Case has been closed automatically by the playbook.
	CaseCloseReasonClosedByPlaybook = CaseCloseReason("closed_by_playbook")
	// The owner of the resource changed.
	CaseCloseReasonResourceOwnerChanged = CaseCloseReason("resource_owner_changed")
	// The resource is not available anymore.
	CaseCloseReasonResourceNotAvailable = CaseCloseReason("resource_not_available")
)

func (enum CaseCloseReason) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_close_reason"
	}
	return string(enum)
}

func (enum CaseCloseReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CaseCloseReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CaseCloseReason(CaseCloseReason(tmp).String())
	return nil
}

type CaseOrigin string

const (
	// Unknown origin.
	CaseOriginUnknownOrigin = CaseOrigin("unknown_origin")
	// Complaint has been submitted through the webform.
	CaseOriginWebform = CaseOrigin("webform")
	// Complaint has been submitted through the abuse mailbox.
	CaseOriginMailbox = CaseOrigin("mailbox")
	// Complaint has been manually submitted.
	CaseOriginManual = CaseOrigin("manual")
)

func (enum CaseOrigin) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_origin"
	}
	return string(enum)
}

func (enum CaseOrigin) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CaseOrigin) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CaseOrigin(CaseOrigin(tmp).String())
	return nil
}

type CaseStatus string

const (
	// Unknown status.
	CaseStatusUnknownStatus = CaseStatus("unknown_status")
	// The case is open and active.
	CaseStatusOpen = CaseStatus("open")
	// The case is closed.
	CaseStatusClosed = CaseStatus("closed")
	// The case doesn't have an assigned playbook.
	CaseStatusNoPlaybookAssigned = CaseStatus("no_playbook_assigned")
)

func (enum CaseStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum CaseStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CaseStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CaseStatus(CaseStatus(tmp).String())
	return nil
}

type CaseType string

const (
	// Unknown type.
	CaseTypeUnknownType = CaseType("unknown_type")
	// Case complaint is about bruteforce.
	CaseTypeBruteforce = CaseType("bruteforce")
	// Case complaint is about botnet.
	CaseTypeBotnet = CaseType("botnet")
	// Case complaint is about copyright.
	CaseTypeCopyright = CaseType("copyright")
	// Case complaint is about malware.
	CaseTypeMalware = CaseType("malware")
	// Case complaint is about phishing.
	CaseTypePhishing = CaseType("phishing")
	// Case complaint is about spam.
	CaseTypeSpam = CaseType("spam")
	// Case complaint is about virus.
	CaseTypeVirus = CaseType("virus")
	// Case complaint is about iptv.
	CaseTypeIptv = CaseType("iptv")
	// Case complaint is about animal welfare.
	CaseTypeAnimalWelfare = CaseType("animal_welfare")
	// Case complaint is about consumer information infringements.
	CaseTypeConsumerInformationInfringement = CaseType("consumer_information_infringement")
	// Case complaint is about child sexual abuse material (CSAM).
	CaseTypeCsam = CaseType("csam")
	// Case complaint is about denial of service (DOS).
	CaseTypeDos = CaseType("dos")
	// Case complaint is about data protection and privacy violation (GDPR).
	CaseTypeGdpr = CaseType("gdpr")
	// Case complaint is about illegal or harmful speech.
	CaseTypeIllegalHarmfulSpeech = CaseType("illegal_harmful_speech")
	// Case complaint is about protection of minors.
	CaseTypeMinorsProtection = CaseType("minors_protection")
	// Case complaint is about negative effects on civic discourse or elections.
	CaseTypeNegativeEffectsCivic = CaseType("negative_effects_civic")
	// Case complaint is about unsafe, non-compliant or prohibited products.
	CaseTypeNonCompliantProduct = CaseType("non_compliant_product")
	// Case complaint is about non-consensual behaviour.
	CaseTypeNonConsensualBehaviour = CaseType("non_consensual_behaviour")
	// Case complaint is about open relay.
	CaseTypeOpenRelay = CaseType("open_relay")
	// Case complaint is about pornography or sexualised content.
	CaseTypePornographySexualisedContent = CaseType("pornography_sexualised_content")
	// Case complaint is about risk for public security.
	CaseTypeRiskPublicSecurity = CaseType("risk_public_security")
	// Case complaint is about scams and/or fraud.
	CaseTypeScamFraud = CaseType("scam_fraud")
	// Case complaint is about security hole.
	CaseTypeSecurityHole = CaseType("security_hole")
	// Case complaint is about self harm.
	CaseTypeSelfHarm = CaseType("self_harm")
	// Case complaint is about spamvertisement.
	CaseTypeSpamvertisement = CaseType("spamvertisement")
	// Case complaint is about violence.
	CaseTypeViolence = CaseType("violence")
	// Case complaint is about another type not listed.
	CaseTypeOther = CaseType("other")
)

func (enum CaseType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum CaseType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CaseType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CaseType(CaseType(tmp).String())
	return nil
}

type ListCasesRequestOrderBy string

const (
	// Last time we received a report for the case descending.
	ListCasesRequestOrderByLastReportedAtDesc = ListCasesRequestOrderBy("last_reported_at_desc")
	// Last time we received a report for the case ascending.
	ListCasesRequestOrderByLastReportedAtAsc = ListCasesRequestOrderBy("last_reported_at_asc")
	// Number of report in a case ascending.
	ListCasesRequestOrderByReportCountAsc = ListCasesRequestOrderBy("report_count_asc")
	// Number of report in a case descending.
	ListCasesRequestOrderByReportCountDesc = ListCasesRequestOrderBy("report_count_desc")
	// Resource value ascending.
	ListCasesRequestOrderByResourceAsc = ListCasesRequestOrderBy("resource_asc")
	// Resource value descending.
	ListCasesRequestOrderByResourceDesc = ListCasesRequestOrderBy("resource_desc")
	// Date of the first report received in the case ascending.
	ListCasesRequestOrderByFirstReportedAtAsc = ListCasesRequestOrderBy("first_reported_at_asc")
	// Date of the first report received in the case descending.
	ListCasesRequestOrderByFirstReportedAtDesc = ListCasesRequestOrderBy("first_reported_at_desc")
	// ID of the case ascending.
	ListCasesRequestOrderByCaseIDAsc = ListCasesRequestOrderBy("case_id_asc")
	// ID of the case descending.
	ListCasesRequestOrderByCaseIDDesc = ListCasesRequestOrderBy("case_id_desc")
	// Creation date ascending.
	ListCasesRequestOrderByCreatedAtAsc = ListCasesRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListCasesRequestOrderByCreatedAtDesc = ListCasesRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListCasesRequestOrderByUpdatedAtAsc = ListCasesRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListCasesRequestOrderByUpdatedAtDesc = ListCasesRequestOrderBy("updated_at_desc")
)

func (enum ListCasesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "last_reported_at_desc"
	}
	return string(enum)
}

func (enum ListCasesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCasesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCasesRequestOrderBy(ListCasesRequestOrderBy(tmp).String())
	return nil
}

type ListPlaybooksRequestOrderBy string

const (
	// Name ascending.
	ListPlaybooksRequestOrderByNameAsc = ListPlaybooksRequestOrderBy("name_asc")
	// Name descending.
	ListPlaybooksRequestOrderByNameDesc = ListPlaybooksRequestOrderBy("name_desc")
	// Creation date ascending.
	ListPlaybooksRequestOrderByCreatedAtAsc = ListPlaybooksRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListPlaybooksRequestOrderByCreatedAtDesc = ListPlaybooksRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListPlaybooksRequestOrderByUpdatedAtAsc = ListPlaybooksRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListPlaybooksRequestOrderByUpdatedAtDesc = ListPlaybooksRequestOrderBy("updated_at_desc")
)

func (enum ListPlaybooksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListPlaybooksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPlaybooksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPlaybooksRequestOrderBy(ListPlaybooksRequestOrderBy(tmp).String())
	return nil
}

type ListReportsRequestOrderBy string

const (
	// Report reception date descending.
	ListReportsRequestOrderByReportedAtDesc = ListReportsRequestOrderBy("reported_at_desc")
	// Report reception date ascending.
	ListReportsRequestOrderByReportedAtAsc = ListReportsRequestOrderBy("reported_at_asc")
	// Creation date ascending.
	ListReportsRequestOrderByCreatedAtAsc = ListReportsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListReportsRequestOrderByCreatedAtDesc = ListReportsRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListReportsRequestOrderByUpdatedAtAsc = ListReportsRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListReportsRequestOrderByUpdatedAtDesc = ListReportsRequestOrderBy("updated_at_desc")
	// Date when the event has been seen by the complainer descending.
	ListReportsRequestOrderByObservedAtDesc = ListReportsRequestOrderBy("observed_at_desc")
	// Date when the event has been seen by the complainer ascending.
	ListReportsRequestOrderByObservedAtAsc = ListReportsRequestOrderBy("observed_at_asc")
)

func (enum ListReportsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "reported_at_desc"
	}
	return string(enum)
}

func (enum ListReportsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListReportsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListReportsRequestOrderBy(ListReportsRequestOrderBy(tmp).String())
	return nil
}

type ListStepsRequestOrderBy string

const (
	// Name ascending.
	ListStepsRequestOrderByNameAsc = ListStepsRequestOrderBy("name_asc")
	// Name descending.
	ListStepsRequestOrderByNameDesc = ListStepsRequestOrderBy("name_desc")
	// Creation date ascending.
	ListStepsRequestOrderByCreatedAtAsc = ListStepsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListStepsRequestOrderByCreatedAtDesc = ListStepsRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListStepsRequestOrderByUpdatedAtAsc = ListStepsRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListStepsRequestOrderByUpdatedAtDesc = ListStepsRequestOrderBy("updated_at_desc")
)

func (enum ListStepsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListStepsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListStepsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListStepsRequestOrderBy(ListStepsRequestOrderBy(tmp).String())
	return nil
}

type ListTransitionsRequestOrderBy string

const (
	// Name ascending.
	ListTransitionsRequestOrderByNameAsc = ListTransitionsRequestOrderBy("name_asc")
	// Name descending.
	ListTransitionsRequestOrderByNameDesc = ListTransitionsRequestOrderBy("name_desc")
	// ID of the source step ascending.
	ListTransitionsRequestOrderByStepSourceIDAsc = ListTransitionsRequestOrderBy("step_source_id_asc")
	// ID of the source step descending.
	ListTransitionsRequestOrderByStepSourceIDDesc = ListTransitionsRequestOrderBy("step_source_id_desc")
	// ID of the destination step ascending.
	ListTransitionsRequestOrderByStepDestinationIDAsc = ListTransitionsRequestOrderBy("step_destination_id_asc")
	// ID of the destination step descending.
	ListTransitionsRequestOrderByStepDestinationIDDesc = ListTransitionsRequestOrderBy("step_destination_id_desc")
	// Type ascending.
	ListTransitionsRequestOrderByTypeAsc = ListTransitionsRequestOrderBy("type_asc")
	// Type descending.
	ListTransitionsRequestOrderByTypeDesc = ListTransitionsRequestOrderBy("type_desc")
	// Creation date ascending.
	ListTransitionsRequestOrderByCreatedAtAsc = ListTransitionsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListTransitionsRequestOrderByCreatedAtDesc = ListTransitionsRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListTransitionsRequestOrderByUpdatedAtAsc = ListTransitionsRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListTransitionsRequestOrderByUpdatedAtDesc = ListTransitionsRequestOrderBy("updated_at_desc")
)

func (enum ListTransitionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListTransitionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTransitionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTransitionsRequestOrderBy(ListTransitionsRequestOrderBy(tmp).String())
	return nil
}

type ListWebhooksRequestOrderBy string

const (
	// Name ascending.
	ListWebhooksRequestOrderByNameAsc = ListWebhooksRequestOrderBy("name_asc")
	// Name descending.
	ListWebhooksRequestOrderByNameDesc = ListWebhooksRequestOrderBy("name_desc")
	// URL ascending.
	ListWebhooksRequestOrderByURLAsc = ListWebhooksRequestOrderBy("url_asc")
	// URL descending.
	ListWebhooksRequestOrderByURLDesc = ListWebhooksRequestOrderBy("url_desc")
	// Method ascending.
	ListWebhooksRequestOrderByMethodAsc = ListWebhooksRequestOrderBy("method_asc")
	// Method descending.
	ListWebhooksRequestOrderByMethodDesc = ListWebhooksRequestOrderBy("method_desc")
	// Creation date ascending.
	ListWebhooksRequestOrderByCreatedAtAsc = ListWebhooksRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListWebhooksRequestOrderByCreatedAtDesc = ListWebhooksRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListWebhooksRequestOrderByUpdatedAtAsc = ListWebhooksRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListWebhooksRequestOrderByUpdatedAtDesc = ListWebhooksRequestOrderBy("updated_at_desc")
)

func (enum ListWebhooksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
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

type ResourceBusinessUnit string

const (
	// Unknown business_unit.
	ResourceBusinessUnitUnknownBusinessUnit = ResourceBusinessUnit("unknown_business_unit")
	// Resource is coming from Elements.
	ResourceBusinessUnitElements = ResourceBusinessUnit("elements")
	// Resource is coming from Dedibox.
	ResourceBusinessUnitDedibox = ResourceBusinessUnit("dedibox")
)

func (enum ResourceBusinessUnit) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_business_unit"
	}
	return string(enum)
}

func (enum ResourceBusinessUnit) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ResourceBusinessUnit) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ResourceBusinessUnit(ResourceBusinessUnit(tmp).String())
	return nil
}

type TransitionType string

const (
	// Unknown type.
	TransitionTypeUnknownType = TransitionType("unknown_type")
	// The transition has to be triggered manually.
	TransitionTypeManual = TransitionType("manual")
	// The transition will be triggered automatically if the precondition is valid.
	TransitionTypeAutomatic = TransitionType("automatic")
)

func (enum TransitionType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum TransitionType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TransitionType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TransitionType(TransitionType(tmp).String())
	return nil
}

type WebhookCredentialType string

const (
	// Unknown type.
	WebhookCredentialTypeUnknownType = WebhookCredentialType("unknown_type")
	// Use BASIC as auth header.
	WebhookCredentialTypeBasic = WebhookCredentialType("basic")
	// Use BEARER as auth header.
	WebhookCredentialTypeBearer = WebhookCredentialType("bearer")
)

func (enum WebhookCredentialType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum WebhookCredentialType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *WebhookCredentialType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = WebhookCredentialType(WebhookCredentialType(tmp).String())
	return nil
}

type WebhookMethod string

const (
	// Unknown method.
	WebhookMethodUnknownMethod = WebhookMethod("unknown_method")
	// Method DELETE has to be use to call the url.
	WebhookMethodDelete = WebhookMethod("delete")
	// Method PATCH has to be use to call the url.
	WebhookMethodPatch = WebhookMethod("patch")
	// Method POST has to be use to call the url.
	WebhookMethodPost = WebhookMethod("post")
	// Method PUT has to be use to call the url.
	WebhookMethodPut = WebhookMethod("put")
)

func (enum WebhookMethod) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_method"
	}
	return string(enum)
}

func (enum WebhookMethod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *WebhookMethod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = WebhookMethod(WebhookMethod(tmp).String())
	return nil
}

// Resource: resource.
type Resource struct {
	// ID: ID of the resource targeted by the case.
	ID string `json:"id"`

	// Type: type of the resource targeted by the case.
	// Default value: unknown
	Type trustandsafety_private_v1.ResourceType `json:"type"`

	// Locality: locality of the resource targeted by the case.
	// Default value: unknown_locality
	Locality resource_v1alpha1.ResourceLocality `json:"locality"`

	// Value: representation of the resource targeted by the case.
	Value string `json:"value"`

	// BusinessUnit: business unit of the resource, Elements or Dedibox.
	// Default value: unknown_business_unit
	BusinessUnit ResourceBusinessUnit `json:"business_unit"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// WebhookCredential: webhook credential.
type WebhookCredential struct {
	// ID: ID of the object.
	ID string `json:"id"`

	// Name: name of the credential.
	Name string `json:"name"`

	// Type: type of the credential.
	// Default value: unknown_type
	Type WebhookCredentialType `json:"type"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// Playbook: playbook.
type Playbook struct {
	// ID: ID of the object.
	ID string `json:"id"`

	// Name: name of the playbook.
	Name string `json:"name"`

	// Description: description of the playbook.
	Description string `json:"description"`

	// Enabled: define if the object is enable or not.
	Enabled bool `json:"enabled"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date.
	UpdatedAt *time.Time `json:"updated_at"`

	// OpenedCaseCount: number of open cases associated with this playbook.
	OpenedCaseCount uint32 `json:"opened_case_count"`
}

// Step: step.
type Step struct {
	// ID: ID of the object.
	ID string `json:"id"`

	// Name: name of the step.
	Name string `json:"name"`

	// Enabled: define if the object is enable or not.
	Enabled bool `json:"enabled"`

	// PlaybookID: id of the playbook to which the step is linked.
	PlaybookID string `json:"playbook_id"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date.
	UpdatedAt *time.Time `json:"updated_at"`

	// OpenedCaseCount: number of opened cases associated with this step.
	OpenedCaseCount uint32 `json:"opened_case_count"`
}

// Transition: transition.
type Transition struct {
	// ID: ID of the object.
	ID string `json:"id"`

	// Name: name of the transition.
	Name string `json:"name"`

	// Enabled: define if the object is enable or not.
	Enabled bool `json:"enabled"`

	// StepSourceID: step source of the transition.
	StepSourceID string `json:"step_source_id"`

	// StepDestinationID: step destination of the transition.
	StepDestinationID string `json:"step_destination_id"`

	// Type: type of the transition.
	// Default value: unknown_type
	Type TransitionType `json:"type"`

	// Precondition: condition which need to be valid to allow automatic transition.
	Precondition string `json:"precondition"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Description: description of the transition.
	Description string `json:"description"`
}

// Case: case.
type Case struct {
	// ID: iD.
	ID string `json:"id"`

	// Class: class of the abuse related to the case.
	// Default value: unknown_class
	Class CaseClass `json:"class"`

	// LastReportedAt: datetime of the last report received for the case.
	LastReportedAt *time.Time `json:"last_reported_at"`

	// ReportCount: number of reports associated to the case.
	ReportCount uint32 `json:"report_count"`

	// FirstReportedAt: datetime at which the first report related to the case has been received.
	FirstReportedAt *time.Time `json:"first_reported_at"`

	// Resource: resource targeted by the case.
	Resource *Resource `json:"resource"`

	// Status: status.
	// Default value: unknown_status
	Status CaseStatus `json:"status"`

	// Type: type of the abuse related to the case.
	// Default value: unknown_type
	Type CaseType `json:"type"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date.
	UpdatedAt *time.Time `json:"updated_at"`

	// PlaybookID: id of the associated playbook.
	PlaybookID *string `json:"playbook_id"`

	// CurrentStepID: current step of the case.
	CurrentStepID string `json:"current_step_id"`

	// CloseReason: explain why the case is closed.
	// Default value: unknown_close_reason
	CloseReason CaseCloseReason `json:"close_reason"`

	// ProjectID: ID of the project linked to the case.
	ProjectID string `json:"project_id"`

	// OrganizationID: ID of the organization linked to the case.
	OrganizationID string `json:"organization_id"`

	// OrganizationCaseCount: number of cases for the case owner.
	OrganizationCaseCount uint32 `json:"organization_case_count"`

	// EstimatedClosingAt: if available teh date at which the case will be automatically resolved.
	// Precisely one of EstimatedClosingAt must be set.
	EstimatedClosingAt *time.Time `json:"estimated_closing_at,omitempty"`

	// Origin: origin of the complaint.
	// Default value: unknown_origin
	Origin CaseOrigin `json:"origin"`
}

// Report: report.
type Report struct {
	// ID: iD.
	ID string `json:"id"`

	// ReportedAt: datetime at which the report has been received.
	ReportedAt *time.Time `json:"reported_at"`

	// ObservedAt: datetime at which the report has been see by the complainer.
	ObservedAt *time.Time `json:"observed_at"`

	// Content: content of the report.
	Content string `json:"content"`

	// SenderEmail: sender email.
	SenderEmail string `json:"sender_email"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// Webhook: webhook.
type Webhook struct {
	// ID: ID of the object.
	ID string `json:"id"`

	// Name: name of the webhook.
	Name string `json:"name"`

	// URL: URL of the webhook.
	URL string `json:"url"`

	// Method: method to use to call the webhook.
	// Default value: unknown_method
	Method WebhookMethod `json:"method"`

	// Template: payload to use for the call.
	Template string `json:"template"`

	// Enabled: define if the object is enable or not.
	Enabled bool `json:"enabled"`

	// Credentials: list of credentials needed for the call.
	Credentials []*WebhookCredential `json:"credentials"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// GetPlaybookResponse: get playbook response.
type GetPlaybookResponse struct {
	// Playbook: playbook with the provided ID.
	Playbook *Playbook `json:"playbook"`

	// Steps: list of steps belonging to the playbook.
	Steps []*Step `json:"steps"`

	// Transitions: list of transitions belonging to the playbook.
	Transitions []*Transition `json:"transitions"`
}

// ListCasesResponse: list cases response.
type ListCasesResponse struct {
	// Cases: list of cases.
	Cases []*Case `json:"cases"`

	// TotalCount: total count of cases.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCasesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCasesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListCasesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Cases = append(r.Cases, results.Cases...)
	r.TotalCount += uint64(len(results.Cases))
	return uint64(len(results.Cases)), nil
}

// ListPlaybooksResponse: list playbooks response.
type ListPlaybooksResponse struct {
	// Playbooks: list of playbooks.
	Playbooks []*Playbook `json:"playbooks"`

	// TotalCount: total count of playbooks.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPlaybooksResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPlaybooksResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPlaybooksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Playbooks = append(r.Playbooks, results.Playbooks...)
	r.TotalCount += uint64(len(results.Playbooks))
	return uint64(len(results.Playbooks)), nil
}

// ListReportsResponse: list reports response.
type ListReportsResponse struct {
	// Reports: list of reports.
	Reports []*Report `json:"reports"`

	// TotalCount: total count of reports.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListReportsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListReportsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListReportsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Reports = append(r.Reports, results.Reports...)
	r.TotalCount += uint64(len(results.Reports))
	return uint64(len(results.Reports)), nil
}

// ListStepsResponse: list steps response.
type ListStepsResponse struct {
	// Steps: list of steps.
	Steps []*Step `json:"steps"`

	// TotalCount: total count of steps.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListStepsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListStepsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListStepsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Steps = append(r.Steps, results.Steps...)
	r.TotalCount += uint64(len(results.Steps))
	return uint64(len(results.Steps)), nil
}

// ListTransitionsResponse: list transitions response.
type ListTransitionsResponse struct {
	// Transitions: list of transitions.
	Transitions []*Transition `json:"transitions"`

	// TotalCount: total count of transitions.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTransitionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTransitionsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListTransitionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Transitions = append(r.Transitions, results.Transitions...)
	r.TotalCount += uint64(len(results.Transitions))
	return uint64(len(results.Transitions)), nil
}

// ListWebhooksResponse: list webhooks response.
type ListWebhooksResponse struct {
	// Webhooks: list of webhooks.
	Webhooks []*Webhook `json:"webhooks"`

	// TotalCount: total count of webhooks.
	TotalCount uint64 `json:"total_count"`
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

// WorkflowAPIAddTimeToCaseRequest: workflow api add time to case request.
type WorkflowAPIAddTimeToCaseRequest struct {
	// CaseID: case for which we want to add time.
	CaseID string `json:"-"`

	// Delay: delay in second to add.
	Delay uint32 `json:"delay"`
}

// WorkflowAPICloseCaseRequest: workflow api close case request.
type WorkflowAPICloseCaseRequest struct {
	// CaseID: ID of the case to close.
	CaseID string `json:"-"`
}

// WorkflowAPIGetCaseRequest: workflow api get case request.
type WorkflowAPIGetCaseRequest struct {
	// CaseID: ID of the case to retrieve.
	CaseID string `json:"-"`
}

// WorkflowAPIGetPlaybookRequest: workflow api get playbook request.
type WorkflowAPIGetPlaybookRequest struct {
	// PlaybookID: ID of the playbook to retrieve.
	PlaybookID string `json:"-"`
}

// WorkflowAPIGetReportRequest: workflow api get report request.
type WorkflowAPIGetReportRequest struct {
	// ReportID: ID of the report to retrieve.
	ReportID string `json:"-"`
}

// WorkflowAPIListCasesRequest: workflow api list cases request.
type WorkflowAPIListCasesRequest struct {
	// OrderBy: sort order of cases.
	// Default value: last_reported_at_desc
	OrderBy ListCasesRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equals to 1.
	Page *int32 `json:"-"`

	// PageSize: number of items per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// OrganizationID: filter by organization ID.
	OrganizationID *string `json:"-"`

	// ObservedAfter: filter after observed_at field.
	ObservedAfter *time.Time `json:"-"`

	// ObservedBefore: filter before observed_at field.
	ObservedBefore *time.Time `json:"-"`

	// Type: filter by type of case.
	// Default value: unknown_type
	Type CaseType `json:"-"`

	// Class: filter by class of case.
	// Default value: unknown_class
	Class CaseClass `json:"-"`

	// CloseReason: filter by close reason of case.
	// Default value: unknown_close_reason
	CloseReason CaseCloseReason `json:"-"`

	// Origin: filter by origin of case.
	// Default value: unknown_origin
	Origin CaseOrigin `json:"-"`

	// WaitForAction: filter cases which need a manual trigger.
	WaitForAction *bool `json:"-"`

	// ResourceValue: filter cases by the value of the attached resource.
	ResourceValue *string `json:"-"`

	// PlaybookID: filter cases by their playbook id.
	PlaybookID *string `json:"-"`

	// StepID: filter cases by their step id.
	StepID *string `json:"-"`
}

// WorkflowAPIListPlaybooksRequest: workflow api list playbooks request.
type WorkflowAPIListPlaybooksRequest struct {
	// OrderBy: sort order of playbooks.
	// Default value: name_asc
	OrderBy ListPlaybooksRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equals to 1.
	Page *int32 `json:"-"`

	// PageSize: number of items per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// Enabled: filter enabled or disabled playbooks.
	Enabled *bool `json:"-"`
}

// WorkflowAPIListReportsRequest: workflow api list reports request.
type WorkflowAPIListReportsRequest struct {
	// OrderBy: sort order of reports.
	// Default value: reported_at_desc
	OrderBy ListReportsRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equals to 1.
	Page *int32 `json:"-"`

	// PageSize: number of items per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// CaseID: case ID linked to the reports.
	CaseID string `json:"-"`
}

// WorkflowAPIListStepsRequest: workflow api list steps request.
type WorkflowAPIListStepsRequest struct {
	// OrderBy: sort order of steps.
	// Default value: name_asc
	OrderBy ListStepsRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equals to 1.
	Page *int32 `json:"-"`

	// PageSize: number of items per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// PlaybookID: filter by Playbook ID.
	PlaybookID *string `json:"-"`

	// Enabled: filter enabled or disabled Steps.
	Enabled *bool `json:"-"`
}

// WorkflowAPIListTransitionsRequest: workflow api list transitions request.
type WorkflowAPIListTransitionsRequest struct {
	// OrderBy: sort order of transitions.
	// Default value: name_asc
	OrderBy ListTransitionsRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equals to 1.
	Page *int32 `json:"-"`

	// PageSize: number of items per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// Enabled: filter enabled or disabled Transitions.
	Enabled *bool `json:"-"`

	// StepSourceID: filter by source Step ID.
	StepSourceID *string `json:"-"`

	// StepDestinationID: filter by destination Step ID.
	StepDestinationID *string `json:"-"`
}

// WorkflowAPIListWebhooksRequest: workflow api list webhooks request.
type WorkflowAPIListWebhooksRequest struct {
	// OrderBy: sort order of webhooks.
	// Default value: name_asc
	OrderBy ListWebhooksRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equals to 1.
	Page *int32 `json:"-"`

	// PageSize: number of items per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// TransitionID: filter by Transition ID.
	TransitionID *string `json:"-"`

	// Enabled: filter enabled or disabled Webhooks.
	Enabled *bool `json:"-"`
}

// WorkflowAPITriggerTransitionRequest: workflow api trigger transition request.
type WorkflowAPITriggerTransitionRequest struct {
	// CaseID: case for which we want to trigger a transition.
	CaseID string `json:"-"`

	// TransitionID: specific ID of the transition to trigger.
	TransitionID string `json:"transition_id"`
}

// WorkflowAPIUpdatePlaybookRequest: workflow api update playbook request.
type WorkflowAPIUpdatePlaybookRequest struct {
	// PlaybookID: ID of the object to toggle.
	PlaybookID string `json:"-"`

	// Enabled: enable or disable the object.
	Enabled *bool `json:"enabled,omitempty"`
}

// WorkflowAPIUpdateStepRequest: workflow api update step request.
type WorkflowAPIUpdateStepRequest struct {
	// StepID: ID of the object to toggle.
	StepID string `json:"-"`

	// Enabled: enable or disable the object.
	Enabled *bool `json:"enabled,omitempty"`
}

// WorkflowAPIUpdateTransitionRequest: workflow api update transition request.
type WorkflowAPIUpdateTransitionRequest struct {
	// TransitionID: ID of the object to toggle.
	TransitionID string `json:"-"`

	// Enabled: enable or disable the object.
	Enabled *bool `json:"enabled,omitempty"`
}

// WorkflowAPIUpdateWebhookRequest: workflow api update webhook request.
type WorkflowAPIUpdateWebhookRequest struct {
	// WebhookID: ID of the object to toggle.
	WebhookID string `json:"-"`

	// Enabled: enable or disable the object.
	Enabled *bool `json:"enabled,omitempty"`
}

// This API aims to be used by the admin console to manage all objects related to Abuse workflow.
type WorkflowAPI struct {
	client *scw.Client
}

// NewWorkflowAPI returns a WorkflowAPI object from a Scaleway client.
func NewWorkflowAPI(client *scw.Client) *WorkflowAPI {
	return &WorkflowAPI{
		client: client,
	}
}

// ListCases: List all cases for the provided organization.
func (s *WorkflowAPI) ListCases(req *WorkflowAPIListCasesRequest, opts ...scw.RequestOption) (*ListCasesResponse, error) {
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
	parameter.AddToQuery(query, "observed_after", req.ObservedAfter)
	parameter.AddToQuery(query, "observed_before", req.ObservedBefore)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "class", req.Class)
	parameter.AddToQuery(query, "close_reason", req.CloseReason)
	parameter.AddToQuery(query, "origin", req.Origin)
	parameter.AddToQuery(query, "wait_for_action", req.WaitForAction)
	parameter.AddToQuery(query, "resource_value", req.ResourceValue)
	parameter.AddToQuery(query, "playbook_id", req.PlaybookID)
	parameter.AddToQuery(query, "step_id", req.StepID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse_admin/v2/cases",
		Query:  query,
	}

	var resp ListCasesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCase: Get the details of one case.
func (s *WorkflowAPI) GetCase(req *WorkflowAPIGetCaseRequest, opts ...scw.RequestOption) (*Case, error) {
	var err error

	if fmt.Sprint(req.CaseID) == "" {
		return nil, errors.New("field CaseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse_admin/v2/cases/" + fmt.Sprint(req.CaseID) + "",
	}

	var resp Case

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListReports: List reports associated with the provided case.
func (s *WorkflowAPI) ListReports(req *WorkflowAPIListReportsRequest, opts ...scw.RequestOption) (*ListReportsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "case_id", req.CaseID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse_admin/v2/reports",
		Query:  query,
	}

	var resp ListReportsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetReport: Get the details of a report.
func (s *WorkflowAPI) GetReport(req *WorkflowAPIGetReportRequest, opts ...scw.RequestOption) (*Report, error) {
	var err error

	if fmt.Sprint(req.ReportID) == "" {
		return nil, errors.New("field ReportID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse_admin/v2/reports/" + fmt.Sprint(req.ReportID) + "",
	}

	var resp Report

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTransitions: List the transitions linked to a step.
func (s *WorkflowAPI) ListTransitions(req *WorkflowAPIListTransitionsRequest, opts ...scw.RequestOption) (*ListTransitionsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "enabled", req.Enabled)
	parameter.AddToQuery(query, "step_source_id", req.StepSourceID)
	parameter.AddToQuery(query, "step_destination_id", req.StepDestinationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse_admin/v2/transitions",
		Query:  query,
	}

	var resp ListTransitionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateTransition: Enable or disable a transition.
func (s *WorkflowAPI) UpdateTransition(req *WorkflowAPIUpdateTransitionRequest, opts ...scw.RequestOption) (*Transition, error) {
	var err error

	if fmt.Sprint(req.TransitionID) == "" {
		return nil, errors.New("field TransitionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/abuse_admin/v2/transitions/" + fmt.Sprint(req.TransitionID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Transition

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListWebhooks: List the webhooks linked to a transition.
func (s *WorkflowAPI) ListWebhooks(req *WorkflowAPIListWebhooksRequest, opts ...scw.RequestOption) (*ListWebhooksResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "transition_id", req.TransitionID)
	parameter.AddToQuery(query, "enabled", req.Enabled)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse_admin/v2/webhooks",
		Query:  query,
	}

	var resp ListWebhooksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateWebhook: Enable or disable a webhook.
func (s *WorkflowAPI) UpdateWebhook(req *WorkflowAPIUpdateWebhookRequest, opts ...scw.RequestOption) (*Webhook, error) {
	var err error

	if fmt.Sprint(req.WebhookID) == "" {
		return nil, errors.New("field WebhookID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/abuse_admin/v2/webhooks/" + fmt.Sprint(req.WebhookID) + "",
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

// ListSteps: List the steps linked to a playbook.
func (s *WorkflowAPI) ListSteps(req *WorkflowAPIListStepsRequest, opts ...scw.RequestOption) (*ListStepsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "playbook_id", req.PlaybookID)
	parameter.AddToQuery(query, "enabled", req.Enabled)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse_admin/v2/steps",
		Query:  query,
	}

	var resp ListStepsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateStep: Enable or disable a step.
func (s *WorkflowAPI) UpdateStep(req *WorkflowAPIUpdateStepRequest, opts ...scw.RequestOption) (*Step, error) {
	var err error

	if fmt.Sprint(req.StepID) == "" {
		return nil, errors.New("field StepID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/abuse_admin/v2/steps/" + fmt.Sprint(req.StepID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Step

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPlaybooks: List playbooks.
func (s *WorkflowAPI) ListPlaybooks(req *WorkflowAPIListPlaybooksRequest, opts ...scw.RequestOption) (*ListPlaybooksResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "enabled", req.Enabled)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse_admin/v2/playbooks",
		Query:  query,
	}

	var resp ListPlaybooksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPlaybook: Get detailed playbook including steps and transitions.
func (s *WorkflowAPI) GetPlaybook(req *WorkflowAPIGetPlaybookRequest, opts ...scw.RequestOption) (*GetPlaybookResponse, error) {
	var err error

	if fmt.Sprint(req.PlaybookID) == "" {
		return nil, errors.New("field PlaybookID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse_admin/v2/playbooks/" + fmt.Sprint(req.PlaybookID) + "",
	}

	var resp GetPlaybookResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePlaybook: Enable or disable a playbook.
func (s *WorkflowAPI) UpdatePlaybook(req *WorkflowAPIUpdatePlaybookRequest, opts ...scw.RequestOption) (*Playbook, error) {
	var err error

	if fmt.Sprint(req.PlaybookID) == "" {
		return nil, errors.New("field PlaybookID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/abuse_admin/v2/playbooks/" + fmt.Sprint(req.PlaybookID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Playbook

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CloseCase: Close a case and skip the remaining steps.
func (s *WorkflowAPI) CloseCase(req *WorkflowAPICloseCaseRequest, opts ...scw.RequestOption) (*Case, error) {
	var err error

	if fmt.Sprint(req.CaseID) == "" {
		return nil, errors.New("field CaseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/abuse_admin/v2/cases/" + fmt.Sprint(req.CaseID) + "/close",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Case

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddTimeToCase: Increase in second the lifetime of a case.
func (s *WorkflowAPI) AddTimeToCase(req *WorkflowAPIAddTimeToCaseRequest, opts ...scw.RequestOption) (*Case, error) {
	var err error

	if fmt.Sprint(req.CaseID) == "" {
		return nil, errors.New("field CaseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/abuse_admin/v2/cases/" + fmt.Sprint(req.CaseID) + "/add-time",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Case

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TriggerTransition: Trigger a transition for the provided case.
func (s *WorkflowAPI) TriggerTransition(req *WorkflowAPITriggerTransitionRequest, opts ...scw.RequestOption) (*Case, error) {
	var err error

	if fmt.Sprint(req.CaseID) == "" {
		return nil, errors.New("field CaseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/abuse_admin/v2/cases/" + fmt.Sprint(req.CaseID) + "/trigger-transition",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Case

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

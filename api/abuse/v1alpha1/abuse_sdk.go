// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package abuse provides methods and message types of the abuse v1alpha1 API.
package abuse

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

type CaseClass string

const (
	// Unknown class.
	CaseClassUnknownClass = CaseClass("unknown_class")
	// The family of Case type is activity.
	CaseClassActivity = CaseClass("activity")
	// The family of Case type is content.
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

type CaseStatus string

const (
	// Unknown status.
	CaseStatusUnknownStatus = CaseStatus("unknown_status")
	// Case is open and active.
	CaseStatusOpen = CaseStatus("open")
	// Case is closed.
	CaseStatusClosed = CaseStatus("closed")
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
	// Case is about bruteforce.
	CaseTypeBruteforce = CaseType("bruteforce")
	// Case is about botnet.
	CaseTypeBotnet = CaseType("botnet")
	// Case is about copyright.
	CaseTypeCopyright = CaseType("copyright")
	// Case is about malware.
	CaseTypeMalware = CaseType("malware")
	// Case is about phishing.
	CaseTypePhishing = CaseType("phishing")
	// Case is about spam.
	CaseTypeSpam = CaseType("spam")
	// Case is about virus.
	CaseTypeVirus = CaseType("virus")
	// Case is about iptv.
	CaseTypeIptv = CaseType("iptv")
	// Case is about animal welfare.
	CaseTypeAnimalWelfare = CaseType("animal_welfare")
	// Case is about consumer information infringements.
	CaseTypeConsumerInformationInfringement = CaseType("consumer_information_infringement")
	// Case is about child sexual abuse material (CSAM).
	CaseTypeCsam = CaseType("csam")
	// Case is about denial of service (DOS).
	CaseTypeDos = CaseType("dos")
	// Case is about data protection and privacy violation (GDPR).
	CaseTypeGdpr = CaseType("gdpr")
	// Case is about illegal or harmful speech.
	CaseTypeIllegalHarmfulSpeech = CaseType("illegal_harmful_speech")
	// Case is about protection of minors.
	CaseTypeMinorsProtection = CaseType("minors_protection")
	// Case is about negative effects on civic discourse or elections.
	CaseTypeNegativeEffectsCivic = CaseType("negative_effects_civic")
	// Case is about unsafe, non-compliant or prohibited products.
	CaseTypeNonCompliantProduct = CaseType("non_compliant_product")
	// Case is about non-consensual behaviour.
	CaseTypeNonConsensualBehaviour = CaseType("non_consensual_behaviour")
	// Case is about open relay.
	CaseTypeOpenRelay = CaseType("open_relay")
	// Case is about pornography or sexualised content.
	CaseTypePornographySexualisedContent = CaseType("pornography_sexualised_content")
	// Case is about risk for public security.
	CaseTypeRiskPublicSecurity = CaseType("risk_public_security")
	// Case is about scams and/or fraud.
	CaseTypeScamFraud = CaseType("scam_fraud")
	// Case is about security hole.
	CaseTypeSecurityHole = CaseType("security_hole")
	// Case is about self harm.
	CaseTypeSelfHarm = CaseType("self_harm")
	// Case is about spamvertisement.
	CaseTypeSpamvertisement = CaseType("spamvertisement")
	// Case is about violence.
	CaseTypeViolence = CaseType("violence")
	// Case is about another type not listed.
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

type ComplaintClass string

const (
	ComplaintClassUnknownClass = ComplaintClass("unknown_class")
	ComplaintClassActivity     = ComplaintClass("activity")
	ComplaintClassContent      = ComplaintClass("content")
)

func (enum ComplaintClass) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_class"
	}
	return string(enum)
}

func (enum ComplaintClass) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ComplaintClass) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ComplaintClass(ComplaintClass(tmp).String())
	return nil
}

type ComplaintStatus string

const (
	ComplaintStatusUnknownStatus = ComplaintStatus("unknown_status")
	ComplaintStatusNew           = ComplaintStatus("new")
	ComplaintStatusForwarded     = ComplaintStatus("forwarded")
	ComplaintStatusProcessing    = ComplaintStatus("processing")
	ComplaintStatusNotSupported  = ComplaintStatus("not_supported")
	ComplaintStatusError         = ComplaintStatus("error")
)

func (enum ComplaintStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum ComplaintStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ComplaintStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ComplaintStatus(ComplaintStatus(tmp).String())
	return nil
}

type ComplaintType string

const (
	ComplaintTypeUnknownType                     = ComplaintType("unknown_type")
	ComplaintTypeBruteforce                      = ComplaintType("bruteforce")
	ComplaintTypeBotnet                          = ComplaintType("botnet")
	ComplaintTypeCopyright                       = ComplaintType("copyright")
	ComplaintTypeMalware                         = ComplaintType("malware")
	ComplaintTypePhishing                        = ComplaintType("phishing")
	ComplaintTypeSpam                            = ComplaintType("spam")
	ComplaintTypeVirus                           = ComplaintType("virus")
	ComplaintTypeIptv                            = ComplaintType("iptv")
	ComplaintTypeAnimalWelfare                   = ComplaintType("animal_welfare")
	ComplaintTypeConsumerInformationInfringement = ComplaintType("consumer_information_infringement")
	ComplaintTypeCsam                            = ComplaintType("csam")
	ComplaintTypeDos                             = ComplaintType("dos")
	ComplaintTypeGdpr                            = ComplaintType("gdpr")
	ComplaintTypeIllegalHarmfulSpeech            = ComplaintType("illegal_harmful_speech")
	ComplaintTypeMinorsProtection                = ComplaintType("minors_protection")
	ComplaintTypeNegativeEffectsCivic            = ComplaintType("negative_effects_civic")
	ComplaintTypeNonCompliantProduct             = ComplaintType("non_compliant_product")
	ComplaintTypeNonConsensualBehaviour          = ComplaintType("non_consensual_behaviour")
	ComplaintTypeOpenRelay                       = ComplaintType("open_relay")
	ComplaintTypePornographySexualisedContent    = ComplaintType("pornography_sexualised_content")
	ComplaintTypeRiskPublicSecurity              = ComplaintType("risk_public_security")
	ComplaintTypeScamFraud                       = ComplaintType("scam_fraud")
	ComplaintTypeSecurityHole                    = ComplaintType("security_hole")
	ComplaintTypeSelfHarm                        = ComplaintType("self_harm")
	ComplaintTypeSpamvertisement                 = ComplaintType("spamvertisement")
	ComplaintTypeViolence                        = ComplaintType("violence")
	ComplaintTypeOther                           = ComplaintType("other")
)

func (enum ComplaintType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum ComplaintType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ComplaintType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ComplaintType(ComplaintType(tmp).String())
	return nil
}

type ListCaseReportsRequestOrderBy string

const (
	// Reported date descending.
	ListCaseReportsRequestOrderByReportedAtDesc = ListCaseReportsRequestOrderBy("reported_at_desc")
	// Reported date ascending.
	ListCaseReportsRequestOrderByReportedAtAsc = ListCaseReportsRequestOrderBy("reported_at_asc")
	// Observed date descending.
	ListCaseReportsRequestOrderByObservedAtDesc = ListCaseReportsRequestOrderBy("observed_at_desc")
	// Observed date ascending.
	ListCaseReportsRequestOrderByObservedAtAsc = ListCaseReportsRequestOrderBy("observed_at_asc")
)

func (enum ListCaseReportsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "reported_at_desc"
	}
	return string(enum)
}

func (enum ListCaseReportsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCaseReportsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCaseReportsRequestOrderBy(ListCaseReportsRequestOrderBy(tmp).String())
	return nil
}

type ListCasesRequestOrderBy string

const (
	// Date of the last reported report in the case descending.
	ListCasesRequestOrderByLastReportedAtDesc = ListCasesRequestOrderBy("last_reported_at_desc")
	// Date of the last reported report in the case ascending.
	ListCasesRequestOrderByLastReportedAtAsc = ListCasesRequestOrderBy("last_reported_at_asc")
	// Number of report per case ascending.
	ListCasesRequestOrderByReportCountAsc = ListCasesRequestOrderBy("report_count_asc")
	// Number of report per case descending.
	ListCasesRequestOrderByReportCountDesc = ListCasesRequestOrderBy("report_count_desc")
	// Resource ID ascending.
	ListCasesRequestOrderByResourceAsc = ListCasesRequestOrderBy("resource_asc")
	// Resource ID descending.
	ListCasesRequestOrderByResourceDesc = ListCasesRequestOrderBy("resource_desc")
	// Reported date ascending.
	ListCasesRequestOrderByReportedAtAsc = ListCasesRequestOrderBy("reported_at_asc")
	// Reported date descending.
	ListCasesRequestOrderByReportedAtDesc = ListCasesRequestOrderBy("reported_at_desc")
	// Case ID ascending.
	ListCasesRequestOrderByCaseIDAsc = ListCasesRequestOrderBy("case_id_asc")
	// Case ID descending.
	ListCasesRequestOrderByCaseIDDesc = ListCasesRequestOrderBy("case_id_desc")
	// Observed date ascending.
	ListCasesRequestOrderByObservedAtAsc = ListCasesRequestOrderBy("observed_at_asc")
	// Observed date descending.
	ListCasesRequestOrderByObservedAtDesc = ListCasesRequestOrderBy("observed_at_desc")
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

// Resource: resource.
type Resource struct {
	// ID: ID of the resource targeted by the case.
	ID string `json:"id"`

	// Type: type of the resource targeted by the case.
	Type string `json:"type"`

	// Locality: locality of the resource targeted by the case.
	Locality string `json:"locality"`

	// Value: representation of the resource targeted by the case.
	Value string `json:"value"`
}

// Report: report.
type Report struct {
	// ReportedAt: datetime at which the report has been received.
	ReportedAt *time.Time `json:"reported_at"`

	// ObservedAt: datetime at which the infringement as been seen.
	ObservedAt *time.Time `json:"observed_at"`

	// Content: content of the report.
	Content string `json:"content"`
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

	// ReportedAt: datetime at which the first report related to the case has been received.
	ReportedAt *time.Time `json:"reported_at"`

	// ObservedAt: datetime at which the infringement as been seen.
	ObservedAt *time.Time `json:"observed_at"`

	// Resource: resource targeted by the case.
	Resource *Resource `json:"resource"`

	// Status: status.
	// Default value: unknown_status
	Status CaseStatus `json:"status"`

	// Type: type of the abuse related to the case.
	// Default value: unknown_type
	Type CaseType `json:"type"`
}

// Complaint: complaint.
type Complaint struct {
	// ID: iD.
	ID string `json:"id"`

	// Evidence: content of the complaint explaining why the complaint is made.
	Evidence string `json:"evidence"`

	// ObservedAt: datetime at which the infringement as been seen.
	ObservedAt *time.Time `json:"observed_at"`

	// ObserverEmail: email of the person or the entity creating the complaint.
	ObserverEmail string `json:"observer_email"`

	// ObserverName: optional name of the person or the entity creating the complaint (optional in case of CSAM).
	ObserverName *string `json:"observer_name"`

	// ResourceValue: representation of the resource targeted by the complaint.
	ResourceValue string `json:"resource_value"`

	// Status: status.
	// Default value: unknown_status
	Status ComplaintStatus `json:"status"`

	// Type: type of the abuse related to the complaint.
	// Default value: unknown_type
	Type ComplaintType `json:"type"`

	// Class: class of the abuse related to the complaint.
	// Default value: unknown_class
	Class ComplaintClass `json:"class"`

	// ReportedAt: datetime at which the complaint has been made.
	ReportedAt *time.Time `json:"reported_at"`
}

// ComplaintAPICreateComplaintRequest: complaint api create complaint request.
type ComplaintAPICreateComplaintRequest struct {
	// Captcha: value of the captcha challenge.
	Captcha string `json:"captcha"`

	// Evidence: content of the complaint explaining why the complaint is made.
	Evidence *string `json:"evidence,omitempty"`

	// ObservedAt: datetime at which the infringement as been seen.
	ObservedAt *time.Time `json:"observed_at,omitempty"`

	// ObserverEmail: email of the person or the entity creating the complaint.
	ObserverEmail string `json:"observer_email"`

	// ObserverName: optional name of the person or the entity creating the complaint (optional in case of CSAM).
	ObserverName *string `json:"observer_name,omitempty"`

	// ResourceValue: representation of the resource targeted by the complaint.
	ResourceValue string `json:"resource_value"`

	// Type: type of the abuse related to the complaint.
	// Default value: unknown_type
	Type ComplaintType `json:"type"`
}

// ListCaseReportsResponse: list case reports response.
type ListCaseReportsResponse struct {
	// Reports: list of reports.
	Reports []*Report `json:"reports"`

	// TotalCount: total count of reports.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCaseReportsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCaseReportsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListCaseReportsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Reports = append(r.Reports, results.Reports...)
	r.TotalCount += uint64(len(results.Reports))
	return uint64(len(results.Reports)), nil
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

// WorkflowAPIGetCaseRequest: workflow api get case request.
type WorkflowAPIGetCaseRequest struct {
	// CaseID: ID of the case to retrieve.
	CaseID string `json:"-"`
}

// WorkflowAPIListCaseReportsRequest: workflow api list case reports request.
type WorkflowAPIListCaseReportsRequest struct {
	// CaseID: case ID linked to the reports.
	CaseID string `json:"-"`

	// OrderBy: sort order of reports.
	// Default value: reported_at_desc
	OrderBy ListCaseReportsRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equals to 1.
	Page *int32 `json:"-"`

	// PageSize: number of items per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`
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
	OrganizationID string `json:"-"`
}

// This API is used to create complaints.
type ComplaintAPI struct {
	client *scw.Client
}

// NewComplaintAPI returns a ComplaintAPI object from a Scaleway client.
func NewComplaintAPI(client *scw.Client) *ComplaintAPI {
	return &ComplaintAPI{
		client: client,
	}
}

// CreateComplaint: Create complaint.
func (s *ComplaintAPI) CreateComplaint(req *ComplaintAPICreateComplaintRequest, opts ...scw.RequestOption) (*Complaint, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/abuse/v1alpha1/complaints",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Complaint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allow to list and display data related to Cases and Reports.
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
		Path:   "/abuse/v1alpha1/cases",
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
		Path:   "/abuse/v1alpha1/cases/" + fmt.Sprint(req.CaseID) + "",
	}

	var resp Case

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCaseReports: List reports associated with the provided case.
func (s *WorkflowAPI) ListCaseReports(req *WorkflowAPIListCaseReportsRequest, opts ...scw.RequestOption) (*ListCaseReportsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.CaseID) == "" {
		return nil, errors.New("field CaseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/abuse/v1alpha1/cases/" + fmt.Sprint(req.CaseID) + "/reports",
		Query:  query,
	}

	var resp ListCaseReportsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

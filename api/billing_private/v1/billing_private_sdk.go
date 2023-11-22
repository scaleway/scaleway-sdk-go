// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package billing_private provides methods and message types of the billing_private v1 API.
package billing_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/billing_private/v1/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/billing_private/v1/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/billing_private/v1/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/billing_private/v1/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/billing_private/v1/scw"
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

type ConsumptionGroupCategory string

const (
	ConsumptionGroupCategoryUnknown             = ConsumptionGroupCategory("unknown")
	ConsumptionGroupCategoryInstance            = ConsumptionGroupCategory("instance")
	ConsumptionGroupCategoryInstanceGpu         = ConsumptionGroupCategory("instance_gpu")
	ConsumptionGroupCategoryTools               = ConsumptionGroupCategory("tools")
	ConsumptionGroupCategoryStorage             = ConsumptionGroupCategory("storage")
	ConsumptionGroupCategoryNetwork             = ConsumptionGroupCategory("network")
	ConsumptionGroupCategoryDiscount            = ConsumptionGroupCategory("discount")
	ConsumptionGroupCategoryBaremetal           = ConsumptionGroupCategory("baremetal")
	ConsumptionGroupCategorySubscription        = ConsumptionGroupCategory("subscription")
	ConsumptionGroupCategorySupport             = ConsumptionGroupCategory("support")
	ConsumptionGroupCategoryIot                 = ConsumptionGroupCategory("iot")
	ConsumptionGroupCategoryAppleSilicon        = ConsumptionGroupCategory("apple_silicon")
	ConsumptionGroupCategoryK8s                 = ConsumptionGroupCategory("k8s")
	ConsumptionGroupCategoryServerless          = ConsumptionGroupCategory("serverless")
	ConsumptionGroupCategoryCompute             = ConsumptionGroupCategory("compute")
	ConsumptionGroupCategoryContainers          = ConsumptionGroupCategory("containers")
	ConsumptionGroupCategoryManagedServices     = ConsumptionGroupCategory("managed_services")
	ConsumptionGroupCategoryManagedDatabases    = ConsumptionGroupCategory("managed_databases")
	ConsumptionGroupCategoryLabs                = ConsumptionGroupCategory("labs")
	ConsumptionGroupCategoryFake                = ConsumptionGroupCategory("fake")
	ConsumptionGroupCategoryObservability       = ConsumptionGroupCategory("observability")
	ConsumptionGroupCategorySecurityAndIdentity = ConsumptionGroupCategory("security_and_identity")
	ConsumptionGroupCategoryObjectStorage       = ConsumptionGroupCategory("object_storage")
)

func (enum ConsumptionGroupCategory) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ConsumptionGroupCategory) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ConsumptionGroupCategory) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ConsumptionGroupCategory(ConsumptionGroupCategory(tmp).String())
	return nil
}

type DownloadConsumptionReportRequestFileType string

const (
	DownloadConsumptionReportRequestFileTypePdf = DownloadConsumptionReportRequestFileType("pdf")
)

func (enum DownloadConsumptionReportRequestFileType) String() string {
	if enum == "" {
		// return default value if empty
		return "pdf"
	}
	return string(enum)
}

func (enum DownloadConsumptionReportRequestFileType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DownloadConsumptionReportRequestFileType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DownloadConsumptionReportRequestFileType(DownloadConsumptionReportRequestFileType(tmp).String())
	return nil
}

type ListOfferActivationsRequestOrderBy string

const (
	ListOfferActivationsRequestOrderByActivationDateDesc = ListOfferActivationsRequestOrderBy("activation_date_desc")
	ListOfferActivationsRequestOrderByActivationDateAsc  = ListOfferActivationsRequestOrderBy("activation_date_asc")
)

func (enum ListOfferActivationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "activation_date_desc"
	}
	return string(enum)
}

func (enum ListOfferActivationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOfferActivationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOfferActivationsRequestOrderBy(ListOfferActivationsRequestOrderBy(tmp).String())
	return nil
}

// ConsumptionGroupLine: consumption group line.
type ConsumptionGroupLine struct {
	// Label: description of the product consumed.
	Label string `json:"label"`

	// Value: consumption price.
	Value *scw.Money `json:"value"`
}

// ConsumptionGroup: consumption group.
type ConsumptionGroup struct {
	// Category: consumption category.
	// Default value: unknown
	Category ConsumptionGroupCategory `json:"category"`

	// Total: total consumption price.
	Total *scw.Money `json:"total"`

	// Lines: all lines corresponding to the consumption.
	Lines []*ConsumptionGroupLine `json:"lines"`
}

// ConsumptionReport: consumption report.
type ConsumptionReport struct {
	// OrganizationID: the organization ID of the resellee.
	OrganizationID string `json:"organization_id"`

	// TotalUntaxed: the total undiscounted amount of the consumption report.
	TotalUntaxed *scw.Money `json:"total_untaxed"`

	// Currency: the currency of the consumption report.
	Currency string `json:"currency"`

	// StartDate: the start date of the consumption report.
	StartDate *time.Time `json:"start_date"`

	// StopDate: the stop date of the consumption report.
	StopDate *time.Time `json:"stop_date"`
}

// OfferActivation: offer activation.
type OfferActivation struct {
	// OrganizationID: organization_id of the offer activation.
	OrganizationID string `json:"organization_id"`

	// OfferSku: sku of the offer.
	OfferSku string `json:"offer_sku"`

	// StartDate: date at which the offer will have effect.
	StartDate *time.Time `json:"start_date"`

	// StopDate: date at which the offer's effect will stop.
	StopDate *time.Time `json:"stop_date"`

	// ActivationDate: date at which the offer activation was triggered.
	ActivationDate *time.Time `json:"activation_date"`

	// Value: consumed amount associated with the offer activation.
	Value string `json:"value"`

	// Unit: unit of the consumed amount.
	Unit string `json:"unit"`
}

// ConsoleAPICreateOfferActivationRequest: console api create offer activation request.
type ConsoleAPICreateOfferActivationRequest struct {
	// OrganizationID: organization_id of the offer activation.
	OrganizationID string `json:"-"`

	// OfferSku: sku of the offer.
	OfferSku string `json:"-"`
}

// ConsoleAPIDownloadConsumptionReportRequest: console api download consumption report request.
type ConsoleAPIDownloadConsumptionReportRequest struct {
	// OrganizationID: the organization ID of the resellee.
	OrganizationID *string `json:"-"`

	// StartDate: the start date of the consumption report.
	StartDate *time.Time `json:"-"`

	// FileType: file type of the report.
	// Default value: pdf
	FileType DownloadConsumptionReportRequestFileType `json:"-"`
}

// ConsoleAPIGetCurrentConsumptionRequest: console api get current consumption request.
type ConsoleAPIGetCurrentConsumptionRequest struct {
	// OrganizationID: get consumption for this organization ID.
	OrganizationID string `json:"-"`
}

// ConsoleAPIGetOfferEligibilityRequest: console api get offer eligibility request.
type ConsoleAPIGetOfferEligibilityRequest struct {
	// OrganizationID: the organization that is tested.
	OrganizationID string `json:"-"`

	// OfferSku: the offer that is tested.
	OfferSku string `json:"-"`
}

// ConsoleAPIListConsumptionReportsRequest: console api list consumption reports request.
type ConsoleAPIListConsumptionReportsRequest struct {
	// OrganizationID: the organization ID of the resellee.
	OrganizationID *string `json:"-"`

	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: a positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// StartedAfter: a timestamp to filter consumption reports by date range.
	StartedAfter *time.Time `json:"-"`

	// StartedBefore: a timestamp to filter consumption reports by date range.
	StartedBefore *time.Time `json:"-"`
}

// ConsoleAPIListOfferActivationsRequest: console api list offer activations request.
type ConsoleAPIListOfferActivationsRequest struct {
	// OrderBy: the sort order for the returned offer activations.
	// Default value: activation_date_desc
	OrderBy ListOfferActivationsRequestOrderBy `json:"-"`

	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: a positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// OrganizationID: organization ID to filter for.
	OrganizationID string `json:"-"`

	// OfferSku: offer sku to filter for.
	OfferSku *string `json:"-"`
}

// CurrentConsumption: current consumption.
type CurrentConsumption struct {
	// ConsumptionGroups: all consumption combined.
	ConsumptionGroups []*ConsumptionGroup `json:"consumption_groups"`

	// UpdatedAt: most recent date on which a consumption was updated.
	UpdatedAt *time.Time `json:"updated_at"`
}

// ListConsumptionReportsResponse: list consumption reports response.
type ListConsumptionReportsResponse struct {
	// ConsumptionReports: the paginated returned consumption reports.
	ConsumptionReports []*ConsumptionReport `json:"consumption_reports"`

	// TotalCount: the total number of consumption reports.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListConsumptionReportsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListConsumptionReportsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListConsumptionReportsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ConsumptionReports = append(r.ConsumptionReports, results.ConsumptionReports...)
	r.TotalCount += uint32(len(results.ConsumptionReports))
	return uint32(len(results.ConsumptionReports)), nil
}

// ListOfferActivationsResponse: list offer activations response.
type ListOfferActivationsResponse struct {
	// OfferActivations: offerActivations that match filters.
	OfferActivations []*OfferActivation `json:"offer_activations"`

	// TotalCount: total number of offer activations that match the selected filters.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOfferActivationsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOfferActivationsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListOfferActivationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.OfferActivations = append(r.OfferActivations, results.OfferActivations...)
	r.TotalCount += uint64(len(results.OfferActivations))
	return uint64(len(results.OfferActivations)), nil
}

// OfferEligibility: offer eligibility.
type OfferEligibility struct {
	// EligibilityState: the state of the elegibility of the organization to the offer at the given date.
	EligibilityState bool `json:"eligibility_state"`
}

// Billing APIs dedicated for the console.
type ConsoleAPI struct {
	client *scw.Client
}

// NewConsoleAPI returns a ConsoleAPI object from a Scaleway client.
func NewConsoleAPI(client *scw.Client) *ConsoleAPI {
	return &ConsoleAPI{
		client: client,
	}
}

// GetCurrentConsumption:
func (s *ConsoleAPI) GetCurrentConsumption(req *ConsoleAPIGetCurrentConsumptionRequest, opts ...scw.RequestOption) (*CurrentConsumption, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-private/v1/current-consumption",
		Query:  query,
	}

	var resp CurrentConsumption

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListConsumptionReports:
func (s *ConsoleAPI) ListConsumptionReports(req *ConsoleAPIListConsumptionReportsRequest, opts ...scw.RequestOption) (*ListConsumptionReportsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "started_after", req.StartedAfter)
	parameter.AddToQuery(query, "started_before", req.StartedBefore)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-private/v1/consumption-reports",
		Query:  query,
	}

	var resp ListConsumptionReportsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DownloadConsumptionReport:
func (s *ConsoleAPI) DownloadConsumptionReport(req *ConsoleAPIDownloadConsumptionReportRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "file_type", req.FileType)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-private/v1/download-consumption-report",
		Query:  query,
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOfferActivation:
func (s *ConsoleAPI) CreateOfferActivation(req *ConsoleAPICreateOfferActivationRequest, opts ...scw.RequestOption) (*OfferActivation, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "offer_sku", req.OfferSku)

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/billing-private/v1/offer-activations",
		Query:  query,
	}

	var resp OfferActivation

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOfferActivations:
func (s *ConsoleAPI) ListOfferActivations(req *ConsoleAPIListOfferActivationsRequest, opts ...scw.RequestOption) (*ListOfferActivationsResponse, error) {
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
	parameter.AddToQuery(query, "offer_sku", req.OfferSku)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-private/v1/offer-activations",
		Query:  query,
	}

	var resp ListOfferActivationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOfferEligibility:
func (s *ConsoleAPI) GetOfferEligibility(req *ConsoleAPIGetOfferEligibilityRequest, opts ...scw.RequestOption) (*OfferEligibility, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "offer_sku", req.OfferSku)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-private/v1/offer-eligibility",
		Query:  query,
	}

	var resp OfferEligibility

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

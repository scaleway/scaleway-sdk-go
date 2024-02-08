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

// Consumption: consumption.
type Consumption struct {
	// TotalAmount: the total amount of the consumption.
	TotalAmount float32 `json:"total_amount"`

	// TotalUntaxed: the total untaxed amount of the consumption.
	TotalUntaxed float32 `json:"total_untaxed"`

	// TotalDiscounted: the total discount amount of the consumption.
	TotalDiscounted float32 `json:"total_discounted"`

	// Currency: the currency of the consumption.
	Currency string `json:"currency"`
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

// ListConsumptionsByCategoryResponseConsumptionByCategory: list consumptions by category response consumption by category.
type ListConsumptionsByCategoryResponseConsumptionByCategory struct {
	CategoryName string `json:"category_name"`

	Consumption *Consumption `json:"consumption"`
}

// ListConsumptionsByProjectResponseConsumptionByProject: list consumptions by project response consumption by project.
type ListConsumptionsByProjectResponseConsumptionByProject struct {
	CategoryName string `json:"category_name"`

	ProductName string `json:"product_name"`

	ProjectID string `json:"project_id"`

	Consumption *Consumption `json:"consumption"`
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

// ListProductsByCategoryResponseProductsByCategory: list products by category response products by category.
type ListProductsByCategoryResponseProductsByCategory struct {
	CategoryName string `json:"category_name"`

	ProductNames []string `json:"product_names"`
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

// ConsoleAPIGetConsumptionRequest: console api get consumption request.
type ConsoleAPIGetConsumptionRequest struct {
	// OrganizationID: get consumption for this organization ID.
	OrganizationID string `json:"-"`

	// BillingPeriod: get consumption on this billing_period.
	// Precisely one of BillingPeriod, Year must be set.
	BillingPeriod *string `json:"billing_period,omitempty"`

	// Year: get consumption on this year.
	// Precisely one of BillingPeriod, Year must be set.
	Year *int64 `json:"year,omitempty"`
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

// ConsoleAPIListConsumptionsByCategoryRequest: console api list consumptions by category request.
type ConsoleAPIListConsumptionsByCategoryRequest struct {
	// OrganizationID: get consumption for this organization ID aggregated by category.
	OrganizationID string `json:"-"`

	// CategoryName: filter the consumption list by category name.
	CategoryName *string `json:"-"`

	// ProductName: filter the consumption list by product name.
	ProductName *string `json:"-"`

	// ProjectID: filter the consumption list by project ID.
	ProjectID *string `json:"-"`

	// Year: filter the consumption list by year.
	Year int64 `json:"-"`
}

// ConsoleAPIListConsumptionsByProjectRequest: console api list consumptions by project request.
type ConsoleAPIListConsumptionsByProjectRequest struct {
	// OrganizationID: get consumption for this organization ID aggregated by project.
	OrganizationID string `json:"-"`

	// CategoryName: filter the consumption list by category name.
	CategoryName *string `json:"-"`

	// ProductName: filter the consumption list by product name.
	ProductName *string `json:"-"`

	// ProjectID: filter the consumption list by project ID.
	ProjectID *string `json:"-"`

	// Year: filter the consumption list by year.
	Year int64 `json:"-"`

	// Page: positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`
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

// ConsoleAPIListProductsByCategoryRequest: console api list products by category request.
type ConsoleAPIListProductsByCategoryRequest struct {
	// OrganizationID: get products consumed by this organization ID, over the last five years, aggregated by category.
	OrganizationID string `json:"-"`
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

// ListConsumptionsByCategoryResponse: list consumptions by category response.
type ListConsumptionsByCategoryResponse struct {
	// TotalCount: total number of returned items.
	TotalCount int64 `json:"total_count"`

	// ConsumptionByCategory: list of consumptions aggregated by category.
	ConsumptionByCategory []*ListConsumptionsByCategoryResponseConsumptionByCategory `json:"consumption_by_category"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListConsumptionsByCategoryResponse) UnsafeGetTotalCount() int64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListConsumptionsByCategoryResponse) UnsafeAppend(res interface{}) (int64, error) {
	results, ok := res.(*ListConsumptionsByCategoryResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ConsumptionByCategory = append(r.ConsumptionByCategory, results.ConsumptionByCategory...)
	r.TotalCount += int64(len(results.ConsumptionByCategory))
	return int64(len(results.ConsumptionByCategory)), nil
}

// ListConsumptionsByProjectResponse: list consumptions by project response.
type ListConsumptionsByProjectResponse struct {
	// TotalCount: total number of returned items.
	TotalCount int64 `json:"total_count"`

	// ConsumptionByProject: list of consumptions aggregated by project.
	ConsumptionByProject []*ListConsumptionsByProjectResponseConsumptionByProject `json:"consumption_by_project"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListConsumptionsByProjectResponse) UnsafeGetTotalCount() int64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListConsumptionsByProjectResponse) UnsafeAppend(res interface{}) (int64, error) {
	results, ok := res.(*ListConsumptionsByProjectResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ConsumptionByProject = append(r.ConsumptionByProject, results.ConsumptionByProject...)
	r.TotalCount += int64(len(results.ConsumptionByProject))
	return int64(len(results.ConsumptionByProject)), nil
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

// ListProductsByCategoryResponse: list products by category response.
type ListProductsByCategoryResponse struct {
	// TotalCount: total number of returned items.
	TotalCount int64 `json:"total_count"`

	// ProductsByCategory: list of product names aggregated by category.
	ProductsByCategory []*ListProductsByCategoryResponseProductsByCategory `json:"products_by_category"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListProductsByCategoryResponse) UnsafeGetTotalCount() int64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListProductsByCategoryResponse) UnsafeAppend(res interface{}) (int64, error) {
	results, ok := res.(*ListProductsByCategoryResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ProductsByCategory = append(r.ProductsByCategory, results.ProductsByCategory...)
	r.TotalCount += int64(len(results.ProductsByCategory))
	return int64(len(results.ProductsByCategory)), nil
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

// GetConsumption:
func (s *ConsoleAPI) GetConsumption(req *ConsoleAPIGetConsumptionRequest, opts ...scw.RequestOption) (*Consumption, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "billing_period", req.BillingPeriod)
	parameter.AddToQuery(query, "year", req.Year)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-private/v1/consumption",
		Query:  query,
	}

	var resp Consumption

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListProductsByCategory:
func (s *ConsoleAPI) ListProductsByCategory(req *ConsoleAPIListProductsByCategoryRequest, opts ...scw.RequestOption) (*ListProductsByCategoryResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-private/v1/products-by-category",
		Query:  query,
	}

	var resp ListProductsByCategoryResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListConsumptionsByCategory:
func (s *ConsoleAPI) ListConsumptionsByCategory(req *ConsoleAPIListConsumptionsByCategoryRequest, opts ...scw.RequestOption) (*ListConsumptionsByCategoryResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "category_name", req.CategoryName)
	parameter.AddToQuery(query, "product_name", req.ProductName)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "year", req.Year)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-private/v1/consumptions-by-category",
		Query:  query,
	}

	var resp ListConsumptionsByCategoryResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListConsumptionsByProject:
func (s *ConsoleAPI) ListConsumptionsByProject(req *ConsoleAPIListConsumptionsByProjectRequest, opts ...scw.RequestOption) (*ListConsumptionsByProjectResponse, error) {
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
	parameter.AddToQuery(query, "category_name", req.CategoryName)
	parameter.AddToQuery(query, "product_name", req.ProductName)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "year", req.Year)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-private/v1/consumptions-by-project",
		Query:  query,
	}

	var resp ListConsumptionsByProjectResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

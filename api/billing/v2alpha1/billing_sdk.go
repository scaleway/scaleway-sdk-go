// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package billing provides methods and message types of the billing v2alpha1 API.
package billing

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
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

type DiscountDiscountMode string

const (
	// Unknown discount mode.
	DiscountDiscountModeUnknownDiscountMode = DiscountDiscountMode("unknown_discount_mode")
	// A rate discount that reduces each customer bill by the discount value percentage.
	DiscountDiscountModeDiscountModeRate = DiscountDiscountMode("discount_mode_rate")
	// A value discount that reduces the amount of the customer bill by the discount value.
	DiscountDiscountModeDiscountModeValue = DiscountDiscountMode("discount_mode_value")
	// A fixed sum to be deducted from the user's bills.
	DiscountDiscountModeDiscountModeSplittable = DiscountDiscountMode("discount_mode_splittable")
)

func (enum DiscountDiscountMode) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_discount_mode"
	}
	return string(enum)
}

func (enum DiscountDiscountMode) Values() []DiscountDiscountMode {
	return []DiscountDiscountMode{
		"unknown_discount_mode",
		"discount_mode_rate",
		"discount_mode_value",
		"discount_mode_splittable",
	}
}

func (enum DiscountDiscountMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DiscountDiscountMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DiscountDiscountMode(DiscountDiscountMode(tmp).String())
	return nil
}

type DiscountFilterType string

const (
	// Unknown filter type.
	DiscountFilterTypeUnknownType = DiscountFilterType("unknown_type")
	// Product category, such as Compute, Network, Observability.
	DiscountFilterTypeProductCategory = DiscountFilterType("product_category")
	// Products within the Product category. For example, VPC, Private Networks, and Public Gateways are products in the Network category.
	DiscountFilterTypeProduct = DiscountFilterType("product")
	// The offer of a product. For example, "VPC Public Gateway S", "VPC Public Gateway M" for the VPC product.
	DiscountFilterTypeProductOffer = DiscountFilterType("product_offer")
	// Identifies the reference based on category, product, range, size, region, and zone. It can sometimes include different product options, such as licenses and monthly payments.
	DiscountFilterTypeProductReference = DiscountFilterType("product_reference")
	// Region name like "FR-PAR", "NL-AMS", "PL-WAW".
	DiscountFilterTypeRegion = DiscountFilterType("region")
	// Zone name like "FR-PAR-1", "FR-PAR-2", "FR-PAR-3".
	DiscountFilterTypeZone = DiscountFilterType("zone")
)

func (enum DiscountFilterType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum DiscountFilterType) Values() []DiscountFilterType {
	return []DiscountFilterType{
		"unknown_type",
		"product_category",
		"product",
		"product_offer",
		"product_reference",
		"region",
		"zone",
	}
}

func (enum DiscountFilterType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DiscountFilterType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DiscountFilterType(DiscountFilterType(tmp).String())
	return nil
}

type DownloadInvoiceRequestFileType string

const (
	DownloadInvoiceRequestFileTypePdf = DownloadInvoiceRequestFileType("pdf")
)

func (enum DownloadInvoiceRequestFileType) String() string {
	if enum == "" {
		// return default value if empty
		return "pdf"
	}
	return string(enum)
}

func (enum DownloadInvoiceRequestFileType) Values() []DownloadInvoiceRequestFileType {
	return []DownloadInvoiceRequestFileType{
		"pdf",
	}
}

func (enum DownloadInvoiceRequestFileType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DownloadInvoiceRequestFileType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DownloadInvoiceRequestFileType(DownloadInvoiceRequestFileType(tmp).String())
	return nil
}

type InvoiceType string

const (
	InvoiceTypeUnknownType = InvoiceType("unknown_type")
	InvoiceTypePeriodic    = InvoiceType("periodic")
	InvoiceTypePurchase    = InvoiceType("purchase")
)

func (enum InvoiceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum InvoiceType) Values() []InvoiceType {
	return []InvoiceType{
		"unknown_type",
		"periodic",
		"purchase",
	}
}

func (enum InvoiceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InvoiceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InvoiceType(InvoiceType(tmp).String())
	return nil
}

type ListDiscountsRequestOrderBy string

const (
	// Order discounts by creation date (descending chronological order).
	ListDiscountsRequestOrderByCreationDateDesc = ListDiscountsRequestOrderBy("creation_date_desc")
	// Order discounts by creation date (ascending chronological order).
	ListDiscountsRequestOrderByCreationDateAsc = ListDiscountsRequestOrderBy("creation_date_asc")
)

func (enum ListDiscountsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "creation_date_desc"
	}
	return string(enum)
}

func (enum ListDiscountsRequestOrderBy) Values() []ListDiscountsRequestOrderBy {
	return []ListDiscountsRequestOrderBy{
		"creation_date_desc",
		"creation_date_asc",
	}
}

func (enum ListDiscountsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDiscountsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDiscountsRequestOrderBy(ListDiscountsRequestOrderBy(tmp).String())
	return nil
}

type ListInvoicesRequestOrderBy string

const (
	ListInvoicesRequestOrderByInvoiceNumberDesc = ListInvoicesRequestOrderBy("invoice_number_desc")
	ListInvoicesRequestOrderByInvoiceNumberAsc  = ListInvoicesRequestOrderBy("invoice_number_asc")
	ListInvoicesRequestOrderByStartDateDesc     = ListInvoicesRequestOrderBy("start_date_desc")
	ListInvoicesRequestOrderByStartDateAsc      = ListInvoicesRequestOrderBy("start_date_asc")
	ListInvoicesRequestOrderByIssuedDateDesc    = ListInvoicesRequestOrderBy("issued_date_desc")
	ListInvoicesRequestOrderByIssuedDateAsc     = ListInvoicesRequestOrderBy("issued_date_asc")
	ListInvoicesRequestOrderByDueDateDesc       = ListInvoicesRequestOrderBy("due_date_desc")
	ListInvoicesRequestOrderByDueDateAsc        = ListInvoicesRequestOrderBy("due_date_asc")
	ListInvoicesRequestOrderByTotalUntaxedDesc  = ListInvoicesRequestOrderBy("total_untaxed_desc")
	ListInvoicesRequestOrderByTotalUntaxedAsc   = ListInvoicesRequestOrderBy("total_untaxed_asc")
	ListInvoicesRequestOrderByTotalTaxedDesc    = ListInvoicesRequestOrderBy("total_taxed_desc")
	ListInvoicesRequestOrderByTotalTaxedAsc     = ListInvoicesRequestOrderBy("total_taxed_asc")
	ListInvoicesRequestOrderByInvoiceTypeDesc   = ListInvoicesRequestOrderBy("invoice_type_desc")
	ListInvoicesRequestOrderByInvoiceTypeAsc    = ListInvoicesRequestOrderBy("invoice_type_asc")
)

func (enum ListInvoicesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "invoice_number_desc"
	}
	return string(enum)
}

func (enum ListInvoicesRequestOrderBy) Values() []ListInvoicesRequestOrderBy {
	return []ListInvoicesRequestOrderBy{
		"invoice_number_desc",
		"invoice_number_asc",
		"start_date_desc",
		"start_date_asc",
		"issued_date_desc",
		"issued_date_asc",
		"due_date_desc",
		"due_date_asc",
		"total_untaxed_desc",
		"total_untaxed_asc",
		"total_taxed_desc",
		"total_taxed_asc",
		"invoice_type_desc",
		"invoice_type_asc",
	}
}

func (enum ListInvoicesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInvoicesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInvoicesRequestOrderBy(ListInvoicesRequestOrderBy(tmp).String())
	return nil
}

// DiscountCoupon: discount coupon.
type DiscountCoupon struct {
	// Description: the description of the coupon.
	Description *string `json:"description"`
}

// DiscountFilter: discount filter.
type DiscountFilter struct {
	// Type: type of the filter.
	// Default value: unknown_type
	Type DiscountFilterType `json:"type"`

	// Value: value of filter, it can be a product/range/region/zone value.
	Value string `json:"value"`
}

// GetConsumptionResponseConsumption: get consumption response consumption.
type GetConsumptionResponseConsumption struct {
	// Value: monetary value of the consumption.
	Value *scw.Money `json:"value"`

	// Description: description of the consumption.
	Description string `json:"description"`

	// ProjectID: project ID of the consumption.
	ProjectID string `json:"project_id"`

	// Category: category of the consumption.
	Category string `json:"category"`

	// OperationPath: unique identifier of the product.
	OperationPath string `json:"operation_path"`
}

// Discount: discount.
type Discount struct {
	// ID: the ID of the discount.
	ID string `json:"id"`

	// CreationDate: the creation date of the discount.
	CreationDate *time.Time `json:"creation_date"`

	// OrganizationID: the organization ID of the discount.
	OrganizationID string `json:"organization_id"`

	// Description: the description of the discount.
	Description string `json:"description"`

	// Value: the initial value of the discount.
	Value float64 `json:"value"`

	// ValueUsed: the value indicating how much of the discount has been used.
	ValueUsed float64 `json:"value_used"`

	// ValueRemaining: the remaining value of the discount.
	ValueRemaining float64 `json:"value_remaining"`

	// Mode: the mode of the discount.
	// Default value: unknown_discount_mode
	Mode DiscountDiscountMode `json:"mode"`

	// StartDate: the start date of the discount.
	StartDate *time.Time `json:"start_date"`

	// StopDate: the stop date of the discount.
	StopDate *time.Time `json:"stop_date"`

	// Coupon: the description of the coupon.
	Coupon *DiscountCoupon `json:"coupon"`

	// Filters: list of products/ranges/regions/zones to limit the usability of discounts.
	Filters []*DiscountFilter `json:"filters"`
}

// Invoice: invoice.
type Invoice struct {
	// ID: invoice ID.
	ID string `json:"id"`

	// StartDate: start date of the billing period.
	StartDate *time.Time `json:"start_date"`

	// IssuedDate: date when the invoice was sent to the customer.
	IssuedDate *time.Time `json:"issued_date"`

	// DueDate: payment time limit, set according to the Organization's payment conditions.
	DueDate *time.Time `json:"due_date"`

	// TotalUntaxed: total amount, untaxed.
	TotalUntaxed *scw.Money `json:"total_untaxed"`

	// TotalTaxed: total amount, taxed.
	TotalTaxed *scw.Money `json:"total_taxed"`

	// InvoiceType: type of invoice.
	// Default value: unknown_type
	InvoiceType InvoiceType `json:"invoice_type"`

	// Number: invoice number.
	Number int32 `json:"number"`
}

// DownloadInvoiceRequest: download invoice request.
type DownloadInvoiceRequest struct {
	// InvoiceID: invoice ID.
	InvoiceID string `json:"-"`

	// FileType: wanted file type.
	// Default value: pdf
	FileType DownloadInvoiceRequestFileType `json:"-"`
}

// GetConsumptionRequest: get consumption request.
type GetConsumptionRequest struct {
	// OrganizationID: filter by organization ID.
	OrganizationID string `json:"-"`
}

// GetConsumptionResponse: get consumption response.
type GetConsumptionResponse struct {
	// Consumptions: detailed consumption list.
	Consumptions []*GetConsumptionResponseConsumption `json:"consumptions"`

	// UpdatedAt: last consumption update date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// ListDiscountsRequest: list discounts request.
type ListDiscountsRequest struct {
	// OrderBy: order discounts in the response by their description.
	// Default value: creation_date_desc
	OrderBy ListDiscountsRequestOrderBy `json:"-"`

	// Page: positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// OrganizationID: ID of the organization.
	OrganizationID *string `json:"-"`
}

// ListDiscountsResponse: list discounts response.
type ListDiscountsResponse struct {
	// TotalCount: total number of discounts.
	TotalCount uint64 `json:"total_count"`

	// Discounts: paginated returned discounts.
	Discounts []*Discount `json:"discounts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDiscountsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDiscountsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListDiscountsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Discounts = append(r.Discounts, results.Discounts...)
	r.TotalCount += uint64(len(results.Discounts))
	return uint64(len(results.Discounts)), nil
}

// ListInvoicesRequest: list invoices request.
type ListInvoicesRequest struct {
	// OrganizationID: organization ID to filter for, only invoices from this Organization will be returned.
	OrganizationID *string `json:"-"`

	// StartedAfter: invoice's `start_date` is greater or equal to `started_after`.
	StartedAfter *time.Time `json:"-"`

	// StartedBefore: invoice's `start_date` precedes `started_before`.
	StartedBefore *time.Time `json:"-"`

	// InvoiceType: invoice type. It can either be `periodic` or `purchase`.
	// Default value: unknown_type
	InvoiceType InvoiceType `json:"-"`

	// Page: positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// OrderBy: how invoices are ordered in the response.
	// Default value: invoice_number_desc
	OrderBy ListInvoicesRequestOrderBy `json:"-"`
}

// ListInvoicesResponse: list invoices response.
type ListInvoicesResponse struct {
	// TotalCount: total number of invoices.
	TotalCount uint32 `json:"total_count"`

	// Invoices: paginated returned invoices.
	Invoices []*Invoice `json:"invoices"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInvoicesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInvoicesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListInvoicesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Invoices = append(r.Invoices, results.Invoices...)
	r.TotalCount += uint32(len(results.Invoices))
	return uint32(len(results.Invoices)), nil
}

// This API allows you to manage and query your Scaleway billing and consumption.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// GetConsumption: The consumption reflects the amount of money you have spent for the products you have used.
// The consumption value is monetary and is not computed in real time.
func (s *API) GetConsumption(req *GetConsumptionRequest, opts ...scw.RequestOption) (*GetConsumptionResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2alpha1/consumption",
		Query:  query,
	}

	var resp GetConsumptionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInvoices: List all your invoices, filtering by `start_date` and `invoice_type`. Each invoice has its own ID.
func (s *API) ListInvoices(req *ListInvoicesRequest, opts ...scw.RequestOption) (*ListInvoicesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "started_after", req.StartedAfter)
	parameter.AddToQuery(query, "started_before", req.StartedBefore)
	parameter.AddToQuery(query, "invoice_type", req.InvoiceType)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2alpha1/invoices",
		Query:  query,
	}

	var resp ListInvoicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DownloadInvoice: Download a specific invoice, specified by its ID.
func (s *API) DownloadInvoice(req *DownloadInvoiceRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "file_type", req.FileType)

	if fmt.Sprint(req.InvoiceID) == "" {
		return nil, errors.New("field InvoiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2alpha1/invoices/" + fmt.Sprint(req.InvoiceID) + "/download",
		Query:  query,
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDiscounts: List all discounts for an organization and usable categories/products/offers/references/regions/zones where the discount can be applied.
func (s *API) ListDiscounts(req *ListDiscountsRequest, opts ...scw.RequestOption) (*ListDiscountsResponse, error) {
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

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2alpha1/discounts",
		Query:  query,
	}

	var resp ListDiscountsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

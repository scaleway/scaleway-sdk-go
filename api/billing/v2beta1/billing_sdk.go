// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package billing provides methods and message types of the billing v2beta1 API.
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
		return string(DiscountDiscountModeUnknownDiscountMode)
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
	DiscountFilterTypeCategoryName = DiscountFilterType("category_name")
	// Products within the Product category. For example, VPC, Private Networks, and Public Gateways are products in the Network category.
	DiscountFilterTypeProductName = DiscountFilterType("product_name")
	// The range of a product. For example, `Production Optimized` or `Cost Optimized` for an `instance` product of the `compute` category.
	DiscountFilterTypeProductRange = DiscountFilterType("product_range")
	// Identifies the reference based on category, product, range, size, region, and zone. It can sometimes include different product options, such as licenses and monthly payments.
	DiscountFilterTypeResourceName = DiscountFilterType("resource_name")
	// Region name like "FR-PAR", "NL-AMS", "PL-WAW".
	DiscountFilterTypeRegion = DiscountFilterType("region")
	// Zone name like "FR-PAR-1", "FR-PAR-2", "FR-PAR-3".
	DiscountFilterTypeZone = DiscountFilterType("zone")
)

func (enum DiscountFilterType) String() string {
	if enum == "" {
		// return default value if empty
		return string(DiscountFilterTypeUnknownType)
	}
	return string(enum)
}

func (enum DiscountFilterType) Values() []DiscountFilterType {
	return []DiscountFilterType{
		"unknown_type",
		"category_name",
		"product_name",
		"product_range",
		"resource_name",
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
		return string(DownloadInvoiceRequestFileTypePdf)
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

type ExportInvoicesRequestFileType string

const (
	ExportInvoicesRequestFileTypeCsv = ExportInvoicesRequestFileType("csv")
)

func (enum ExportInvoicesRequestFileType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ExportInvoicesRequestFileTypeCsv)
	}
	return string(enum)
}

func (enum ExportInvoicesRequestFileType) Values() []ExportInvoicesRequestFileType {
	return []ExportInvoicesRequestFileType{
		"csv",
	}
}

func (enum ExportInvoicesRequestFileType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ExportInvoicesRequestFileType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ExportInvoicesRequestFileType(ExportInvoicesRequestFileType(tmp).String())
	return nil
}

type ExportInvoicesRequestOrderBy string

const (
	ExportInvoicesRequestOrderByInvoiceNumberDesc = ExportInvoicesRequestOrderBy("invoice_number_desc")
	ExportInvoicesRequestOrderByInvoiceNumberAsc  = ExportInvoicesRequestOrderBy("invoice_number_asc")
	ExportInvoicesRequestOrderByStartDateDesc     = ExportInvoicesRequestOrderBy("start_date_desc")
	ExportInvoicesRequestOrderByStartDateAsc      = ExportInvoicesRequestOrderBy("start_date_asc")
	ExportInvoicesRequestOrderByIssuedDateDesc    = ExportInvoicesRequestOrderBy("issued_date_desc")
	ExportInvoicesRequestOrderByIssuedDateAsc     = ExportInvoicesRequestOrderBy("issued_date_asc")
	ExportInvoicesRequestOrderByDueDateDesc       = ExportInvoicesRequestOrderBy("due_date_desc")
	ExportInvoicesRequestOrderByDueDateAsc        = ExportInvoicesRequestOrderBy("due_date_asc")
	ExportInvoicesRequestOrderByTotalUntaxedDesc  = ExportInvoicesRequestOrderBy("total_untaxed_desc")
	ExportInvoicesRequestOrderByTotalUntaxedAsc   = ExportInvoicesRequestOrderBy("total_untaxed_asc")
	ExportInvoicesRequestOrderByTotalTaxedDesc    = ExportInvoicesRequestOrderBy("total_taxed_desc")
	ExportInvoicesRequestOrderByTotalTaxedAsc     = ExportInvoicesRequestOrderBy("total_taxed_asc")
	ExportInvoicesRequestOrderByInvoiceTypeDesc   = ExportInvoicesRequestOrderBy("invoice_type_desc")
	ExportInvoicesRequestOrderByInvoiceTypeAsc    = ExportInvoicesRequestOrderBy("invoice_type_asc")
)

func (enum ExportInvoicesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ExportInvoicesRequestOrderByInvoiceNumberDesc)
	}
	return string(enum)
}

func (enum ExportInvoicesRequestOrderBy) Values() []ExportInvoicesRequestOrderBy {
	return []ExportInvoicesRequestOrderBy{
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

func (enum ExportInvoicesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ExportInvoicesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ExportInvoicesRequestOrderBy(ExportInvoicesRequestOrderBy(tmp).String())
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
		return string(InvoiceTypeUnknownType)
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

type ListConsumptionsRequestOrderBy string

const (
	// Order consumptions by update date (descending chronological order).
	ListConsumptionsRequestOrderByUpdatedAtDesc = ListConsumptionsRequestOrderBy("updated_at_desc")
	// Order consumptions by update date (ascending chronological order).
	ListConsumptionsRequestOrderByUpdatedAtAsc = ListConsumptionsRequestOrderBy("updated_at_asc")
	// Order consumptions by category name (descending alphabetical order).
	ListConsumptionsRequestOrderByCategoryNameDesc = ListConsumptionsRequestOrderBy("category_name_desc")
	// Order consumptions by category name (ascending alphabetical order).
	ListConsumptionsRequestOrderByCategoryNameAsc = ListConsumptionsRequestOrderBy("category_name_asc")
)

func (enum ListConsumptionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListConsumptionsRequestOrderByUpdatedAtDesc)
	}
	return string(enum)
}

func (enum ListConsumptionsRequestOrderBy) Values() []ListConsumptionsRequestOrderBy {
	return []ListConsumptionsRequestOrderBy{
		"updated_at_desc",
		"updated_at_asc",
		"category_name_desc",
		"category_name_asc",
	}
}

func (enum ListConsumptionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListConsumptionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListConsumptionsRequestOrderBy(ListConsumptionsRequestOrderBy(tmp).String())
	return nil
}

type ListDiscountsRequestOrderBy string

const (
	// Order discounts by creation date (descending chronological order).
	ListDiscountsRequestOrderByCreationDateDesc = ListDiscountsRequestOrderBy("creation_date_desc")
	// Order discounts by creation date (ascending chronological order).
	ListDiscountsRequestOrderByCreationDateAsc = ListDiscountsRequestOrderBy("creation_date_asc")
	// Order discounts by start date (descending chronological order).
	ListDiscountsRequestOrderByStartDateDesc = ListDiscountsRequestOrderBy("start_date_desc")
	// Order discounts by start date (ascending chronological order).
	ListDiscountsRequestOrderByStartDateAsc = ListDiscountsRequestOrderBy("start_date_asc")
	// Order discounts by stop date (descending chronological order).
	ListDiscountsRequestOrderByStopDateDesc = ListDiscountsRequestOrderBy("stop_date_desc")
	// Order discounts by stop date (ascending chronological order).
	ListDiscountsRequestOrderByStopDateAsc = ListDiscountsRequestOrderBy("stop_date_asc")
)

func (enum ListDiscountsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListDiscountsRequestOrderByCreationDateDesc)
	}
	return string(enum)
}

func (enum ListDiscountsRequestOrderBy) Values() []ListDiscountsRequestOrderBy {
	return []ListDiscountsRequestOrderBy{
		"creation_date_desc",
		"creation_date_asc",
		"start_date_desc",
		"start_date_asc",
		"stop_date_desc",
		"stop_date_asc",
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
		return string(ListInvoicesRequestOrderByInvoiceNumberDesc)
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

type ListTaxesRequestOrderBy string

const (
	// Order consumptions by update date (descending chronological order).
	ListTaxesRequestOrderByUpdatedAtDesc = ListTaxesRequestOrderBy("updated_at_desc")
	// Order consumptions by update date (ascending chronological order).
	ListTaxesRequestOrderByUpdatedAtAsc = ListTaxesRequestOrderBy("updated_at_asc")
	// Order consumptions by category name (descending alphabetical order).
	ListTaxesRequestOrderByCategoryNameDesc = ListTaxesRequestOrderBy("category_name_desc")
	// Order consumptions by category name (ascending alphabetical order).
	ListTaxesRequestOrderByCategoryNameAsc = ListTaxesRequestOrderBy("category_name_asc")
)

func (enum ListTaxesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListTaxesRequestOrderByUpdatedAtDesc)
	}
	return string(enum)
}

func (enum ListTaxesRequestOrderBy) Values() []ListTaxesRequestOrderBy {
	return []ListTaxesRequestOrderBy{
		"updated_at_desc",
		"updated_at_asc",
		"category_name_desc",
		"category_name_asc",
	}
}

func (enum ListTaxesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTaxesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTaxesRequestOrderBy(ListTaxesRequestOrderBy(tmp).String())
	return nil
}

// DiscountCoupon: discount coupon.
type DiscountCoupon struct {
	// Description: the description of the coupon.
	Description *string `json:"description"`
}

// DiscountFilter: discount filter.
type DiscountFilter struct {
	// Type: type of the filter (category name, product name, product range, resource name, region or zone).
	// Default value: unknown_type
	Type DiscountFilterType `json:"type"`

	// Value: value of filter.
	Value string `json:"value"`

	// Exclude: boolean to describe if filter is an excluding filter.
	Exclude bool `json:"exclude"`
}

// ListConsumptionsResponseConsumption: list consumptions response consumption.
type ListConsumptionsResponseConsumption struct {
	// Value: monetary value of the consumption.
	Value *scw.Money `json:"value"`

	// ProductName: the product name. For example, "VPC Public Gateway S", "VPC Public Gateway M" for the VPC product.
	ProductName string `json:"product_name"`

	// ResourceName: identifies the reference based on the category.
	ResourceName string `json:"resource_name"`

	// Sku: unique identifier of the product.
	Sku string `json:"sku"`

	// ProjectID: project ID of the consumption.
	ProjectID string `json:"project_id"`

	// CategoryName: name of consumption category.
	CategoryName string `json:"category_name"`

	// Unit: unit of consumed quantity.
	Unit string `json:"unit"`

	// BilledQuantity: consumed quantity.
	BilledQuantity string `json:"billed_quantity"`
}

// Discount: discount.
type Discount struct {
	// ID: the ID of the discount.
	ID string `json:"id"`

	// CreationDate: the creation date of the discount.
	CreationDate *time.Time `json:"creation_date"`

	// OrganizationID: the Organization ID of the discount.
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

	// Filters: list of the discount scopes.
	Filters []*DiscountFilter `json:"filters"`
}

// Invoice: invoice.
type Invoice struct {
	// ID: invoice ID.
	ID string `json:"id"`

	OrganizationID string `json:"organization_id"`

	OrganizationName string `json:"organization_name"`

	// StartDate: start date of the billing period.
	StartDate *time.Time `json:"start_date"`

	StopDate *time.Time `json:"stop_date"`

	// BillingPeriod: the billing period of the invoice in the YYYY-MM format.
	BillingPeriod *time.Time `json:"billing_period"`

	// IssuedDate: date when the invoice was sent to the customer.
	IssuedDate *time.Time `json:"issued_date"`

	// DueDate: payment time limit, set according to the Organization's payment conditions.
	DueDate *time.Time `json:"due_date"`

	// TotalUntaxed: total amount, untaxed.
	TotalUntaxed *scw.Money `json:"total_untaxed"`

	// TotalTaxed: total amount, taxed.
	TotalTaxed *scw.Money `json:"total_taxed"`

	// TotalTax: the total tax amount of the invoice.
	TotalTax *scw.Money `json:"total_tax"`

	// TotalDiscount: the total discount amount of the invoice.
	TotalDiscount *scw.Money `json:"total_discount"`

	// TotalUndiscount: the total amount of the invoice before applying the discount.
	TotalUndiscount *scw.Money `json:"total_undiscount"`

	// Type: type of invoice, either periodic or purchase.
	// Default value: unknown_type
	Type InvoiceType `json:"type"`

	// State: the state of the Invoice.
	State string `json:"state"`

	// Number: invoice number.
	Number int32 `json:"number"`

	// SellerName: the name of the seller (Scaleway).
	SellerName string `json:"seller_name"`
}

// ListTaxesResponseTax: list taxes response tax.
type ListTaxesResponseTax struct {
	// Description: description of the tax applied.
	Description string `json:"description"`

	// Currency: the three-letter currency code.
	Currency string `json:"currency"`

	// Rate: applied tax rate (0.2 means a VAT of 20%).
	Rate *float64 `json:"rate"`

	// TotalTaxValue: the total tax value of the consumption.
	TotalTaxValue *float64 `json:"total_tax_value"`
}

// DownloadInvoiceRequest: download invoice request.
type DownloadInvoiceRequest struct {
	// InvoiceID: invoice ID.
	InvoiceID string `json:"-"`

	// FileType: file type. PDF by default.
	// Default value: pdf
	FileType DownloadInvoiceRequestFileType `json:"-"`
}

// ExportInvoicesRequest: export invoices request.
type ExportInvoicesRequest struct {
	// OrganizationID: organization ID. If specified, only invoices from this Organization will be returned.
	OrganizationID *string `json:"-"`

	// BillingPeriodStartAfter: return only invoice with start date greater than billing_period_start.
	BillingPeriodStartAfter *time.Time `json:"-"`

	// BillingPeriodStartBefore: return only invoice with start date less than billing_period_start.
	BillingPeriodStartBefore *time.Time `json:"-"`

	// InvoiceType: invoice type. It can either be `periodic` or `purchase`.
	// Default value: unknown_type
	InvoiceType InvoiceType `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// OrderBy: how invoices are ordered in the response.
	// Default value: invoice_number_desc
	OrderBy ExportInvoicesRequestOrderBy `json:"-"`

	// FileType: file format for exporting the invoice list.
	// Default value: csv
	FileType ExportInvoicesRequestFileType `json:"-"`
}

// GetInvoiceRequest: get invoice request.
type GetInvoiceRequest struct {
	// InvoiceID: invoice ID.
	InvoiceID string `json:"-"`
}

// ListConsumptionsRequest: list consumptions request.
type ListConsumptionsRequest struct {
	// OrderBy: order consumptions list in the response by their update date.
	// Default value: updated_at_desc
	OrderBy ListConsumptionsRequestOrderBy `json:"-"`

	// Page: positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// OrganizationID: filter by Organization ID.
	// Precisely one of OrganizationID, ProjectID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// ProjectID: filter by Project ID.
	// Precisely one of OrganizationID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// CategoryName: filter by name of a Category as they are shown in the invoice (Compute, Network, Observability).
	CategoryName *string `json:"-"`

	// BillingPeriod: filter by the billing period in the YYYY-MM format. If it is empty the current billing period will be used as default.
	BillingPeriod *string `json:"-"`
}

// ListConsumptionsResponse: list consumptions response.
type ListConsumptionsResponse struct {
	// Consumptions: detailed consumption list.
	Consumptions []*ListConsumptionsResponseConsumption `json:"consumptions"`

	// TotalCount: total number of returned items.
	TotalCount uint64 `json:"total_count"`

	// TotalDiscountUntaxedValue: sum of all discounts, displayed only when no category or project ID filter is applied.
	TotalDiscountUntaxedValue float64 `json:"total_discount_untaxed_value"`

	// UpdatedAt: last consumption update date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListConsumptionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListConsumptionsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListConsumptionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Consumptions = append(r.Consumptions, results.Consumptions...)
	r.TotalCount += uint64(len(results.Consumptions))
	return uint64(len(results.Consumptions)), nil
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
	// OrganizationID: organization ID. If specified, only invoices from this Organization will be returned.
	OrganizationID *string `json:"-"`

	// BillingPeriodStartAfter: return only invoice with start date greater than billing_period_start.
	BillingPeriodStartAfter *time.Time `json:"-"`

	// BillingPeriodStartBefore: return only invoice with start date less than billing_period_start.
	BillingPeriodStartBefore *time.Time `json:"-"`

	// InvoiceType: invoice type. It can either be `periodic` or `purchase`.
	// Default value: unknown_type
	InvoiceType InvoiceType `json:"-"`

	// Page: page number.
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
	TotalCount uint64 `json:"total_count"`

	// Invoices: paginated returned invoices.
	Invoices []*Invoice `json:"invoices"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInvoicesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInvoicesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListInvoicesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Invoices = append(r.Invoices, results.Invoices...)
	r.TotalCount += uint64(len(results.Invoices))
	return uint64(len(results.Invoices)), nil
}

// ListTaxesRequest: list taxes request.
type ListTaxesRequest struct {
	// OrderBy: order consumed taxes list in the response by their update date.
	// Default value: updated_at_desc
	OrderBy ListTaxesRequestOrderBy `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// OrganizationID: filter by Organization ID.
	OrganizationID string `json:"-"`

	// BillingPeriod: filter by the billing period in the YYYY-MM format. If it is empty the current billing period will be used as default.
	BillingPeriod *string `json:"-"`
}

// ListTaxesResponse: list taxes response.
type ListTaxesResponse struct {
	// Taxes: detailed consumption tax.
	Taxes []*ListTaxesResponseTax `json:"taxes"`

	// TotalCount: total number of returned items.
	TotalCount uint64 `json:"total_count"`

	// UpdatedAt: last consumption update date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTaxesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTaxesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListTaxesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Taxes = append(r.Taxes, results.Taxes...)
	r.TotalCount += uint64(len(results.Taxes))
	return uint64(len(results.Taxes)), nil
}

// RedeemCouponRequest: redeem coupon request.
type RedeemCouponRequest struct {
	// OrganizationID: the Organization ID of the discount.
	OrganizationID string `json:"-"`

	// Code: the code of the coupon to redeem.
	Code string `json:"-"`
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

// ListConsumptions: Consumption allows you to retrieve your past or current consumption cost, by project or category.
func (s *API) ListConsumptions(req *ListConsumptionsRequest, opts ...scw.RequestOption) (*ListConsumptionsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "category_name", req.CategoryName)
	parameter.AddToQuery(query, "billing_period", req.BillingPeriod)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2beta1/consumptions",
		Query:  query,
	}

	var resp ListConsumptionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListTaxes: Consumption Tax allows you to retrieve your past or current tax charges, by project or category.
func (s *API) ListTaxes(req *ListTaxesRequest, opts ...scw.RequestOption) (*ListTaxesResponse, error) {
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
	parameter.AddToQuery(query, "billing_period", req.BillingPeriod)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2beta1/taxes",
		Query:  query,
	}

	var resp ListTaxesResponse

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
	parameter.AddToQuery(query, "billing_period_start_after", req.BillingPeriodStartAfter)
	parameter.AddToQuery(query, "billing_period_start_before", req.BillingPeriodStartBefore)
	parameter.AddToQuery(query, "invoice_type", req.InvoiceType)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2beta1/invoices",
		Query:  query,
	}

	var resp ListInvoicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportInvoices: Export invoices in a CSV file.
func (s *API) ExportInvoices(req *ExportInvoicesRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "billing_period_start_after", req.BillingPeriodStartAfter)
	parameter.AddToQuery(query, "billing_period_start_before", req.BillingPeriodStartBefore)
	parameter.AddToQuery(query, "invoice_type", req.InvoiceType)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "file_type", req.FileType)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2beta1/export-invoices",
		Query:  query,
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInvoice: Get a specific invoice, specified by its ID.
func (s *API) GetInvoice(req *GetInvoiceRequest, opts ...scw.RequestOption) (*Invoice, error) {
	var err error

	if fmt.Sprint(req.InvoiceID) == "" {
		return nil, errors.New("field InvoiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2beta1/invoices/" + fmt.Sprint(req.InvoiceID) + "",
	}

	var resp Invoice

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
		Path:   "/billing/v2beta1/invoices/" + fmt.Sprint(req.InvoiceID) + "/download",
		Query:  query,
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDiscounts: List all discounts for your Organization and usable categories, products, offers, references, regions and zones where the discount can be applied. As a reseller:
// - If you do not specify an `organization_id` you will list the discounts applied to your own Organization and your customers
// - If you indicate your `organization_id` you will list only the discounts applied to your Organization
// - If you indicate `the organization_id` of one of your customers, you will list the discounts applied to their Organization.
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
		Path:   "/billing/v2beta1/discounts",
		Query:  query,
	}

	var resp ListDiscountsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RedeemCoupon: Redeem a coupon given the related code.
func (s *API) RedeemCoupon(req *RedeemCouponRequest, opts ...scw.RequestOption) (*Discount, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "code", req.Code)

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/billing/v2beta1/redeem-coupon",
		Query:  query,
	}

	var resp Discount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

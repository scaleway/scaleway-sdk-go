// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package billing_admin provides methods and message types of the billing_admin v1 API.
package billing_admin

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

type DiscountDiscountMode string

const (
	DiscountDiscountModeUnknownDiscountMode    = DiscountDiscountMode("unknown_discount_mode")
	DiscountDiscountModeDiscountModeRate       = DiscountDiscountMode("discount_mode_rate")
	DiscountDiscountModeDiscountModeValue      = DiscountDiscountMode("discount_mode_value")
	DiscountDiscountModeDiscountModeSplittable = DiscountDiscountMode("discount_mode_splittable")
)

func (enum DiscountDiscountMode) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_discount_mode"
	}
	return string(enum)
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

type InvoiceState string

const (
	InvoiceStateUnknownState  = InvoiceState("unknown_state")
	InvoiceStateDraft         = InvoiceState("draft")
	InvoiceStateOutdated      = InvoiceState("outdated")
	InvoiceStateStopped       = InvoiceState("stopped")
	InvoiceStateIncomplete    = InvoiceState("incomplete")
	InvoiceStateIssuing       = InvoiceState("issuing")
	InvoiceStateIssued        = InvoiceState("issued")
	InvoiceStatePaid          = InvoiceState("paid")
	InvoiceStatePendingUpdate = InvoiceState("pending_update")
)

func (enum InvoiceState) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_state"
	}
	return string(enum)
}

func (enum InvoiceState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InvoiceState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InvoiceState(InvoiceState(tmp).String())
	return nil
}

type ListFormat string

const (
	ListFormatJSON = ListFormat("json")
	ListFormatCsv  = ListFormat("csv")
)

func (enum ListFormat) String() string {
	if enum == "" {
		// return default value if empty
		return "json"
	}
	return string(enum)
}

func (enum ListFormat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFormat) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFormat(ListFormat(tmp).String())
	return nil
}

// Discount: discount.
type Discount struct {
	// ID: the ID of the discount.
	ID string `json:"id"`

	// CreationDate: the creation date of the discount.
	CreationDate *time.Time `json:"creation_date"`

	// Organization: the organization ID of the discount.
	Organization string `json:"organization"`

	// Description: the public description of the discount.
	Description string `json:"description"`

	// InternalDescription: the internal description of the discount.
	InternalDescription string `json:"internal_description"`

	// Value: the initial value of the discount.
	Value float64 `json:"value"`

	// ValueUsed: the used value of the discount.
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

	// Deleted: indicates if the discount is deleted or not.
	Deleted bool `json:"deleted"`

	// DeletionDate: the deletion date of the discount.
	DeletionDate *time.Time `json:"deletion_date"`

	// CouponID: the coupon ID of the discount.
	CouponID string `json:"coupon_id"`
}

// IndexInvoice: An invoice in the list.
type IndexInvoice struct {
	// ID: the ID of the invoice.
	ID string `json:"id"`

	// OrganizationID: the organization ID of the invoice consumer.
	OrganizationID string `json:"organization_id"`

	// StartDate: the start date of the discount.
	StartDate *time.Time `json:"start_date"`

	// StopDate: the stop date of the discount.
	StopDate *time.Time `json:"stop_date"`

	// OrganizationName: the name of the invoice payer.
	OrganizationName string `json:"organization_name"`

	// Number: the number of the invoice.
	Number int32 `json:"number"`

	// Currency: the currency of the invoice.
	Currency string `json:"currency"`

	// State: the state of the invoice.
	// Default value: unknown_state
	State InvoiceState `json:"state"`

	// BillingPeriod: the billing period of the invoice in the format YYYY-MM.
	BillingPeriod string `json:"billing_period"`

	// LastUpdate: the last update date of the invoice.
	LastUpdate *time.Time `json:"last_update"`

	// IssuedDate: the issued date of the invoice.
	IssuedDate *time.Time `json:"issued_date"`

	// DueDate: the payment date limit of the invoice.
	DueDate *time.Time `json:"due_date"`

	// TotalUndiscounted: the total undiscounted amount of the invoice.
	TotalUndiscounted float64 `json:"total_undiscounted"`

	// TotalUntaxed: the total untaxed amount of the invoice.
	TotalUntaxed float64 `json:"total_untaxed"`

	// TotalTax: the total tax amount of the invoice.
	TotalTax float64 `json:"total_tax"`

	// TotalTaxed: the total taxed amount of the invoice.
	TotalTaxed float64 `json:"total_taxed"`

	// TotalDiscount: the total discount amount of the invoice.
	TotalDiscount float64 `json:"total_discount"`

	// PaymentMethod: the payment method of the invoice.
	PaymentMethod string `json:"payment_method"`
}

// IndexUsage: A usage in the list.
type IndexUsage struct {
	// ID: the ID of the usage.
	ID string `json:"id"`

	// OrganizationID: the organization ID of the usage.
	OrganizationID string `json:"organization_id"`

	// ProductID: the product ID of the usage.
	ProductID string `json:"product_id"`

	// OperationPath: the operation path of the usage.
	OperationPath string `json:"operation_path"`

	// StartDate: the start date of the usage.
	StartDate *time.Time `json:"start_date"`

	// StopDate: the stop date of the usage.
	StopDate *time.Time `json:"stop_date"`

	// ResourceID: the resource ID of the usage.
	ResourceID string `json:"resource_id"`

	// Quantity: the quantity of the usage.
	Quantity string `json:"quantity"`

	// Unit: the unit of the usage.
	Unit string `json:"unit"`

	// Granularity: the granularity of the usage.
	Granularity string `json:"granularity"`

	// InvoiceID: the invoice ID of the usage.
	InvoiceID string `json:"invoice_id"`

	// State: the state of the usage.
	State string `json:"state"`
}

// ListDiscountsRequest: list discounts request.
type ListDiscountsRequest struct {
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`

	// Organization: filter discounts by organization ID.
	Organization *string `json:"-"`
}

// ListDiscountsResponse: list discounts response.
type ListDiscountsResponse struct {
	// TotalCount: the total number of discounts.
	TotalCount uint32 `json:"total_count"`

	// Discounts: the paginated returned discounts.
	Discounts []*Discount `json:"discounts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDiscountsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDiscountsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDiscountsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Discounts = append(r.Discounts, results.Discounts...)
	r.TotalCount += uint32(len(results.Discounts))
	return uint32(len(results.Discounts)), nil
}

// ListInvoicesRequest: list invoices request.
type ListInvoicesRequest struct {
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`

	// OrganizationID: filter invoices by organization ID.
	OrganizationID *string `json:"-"`

	// StartDate: filter invoices by start date.
	StartDate *time.Time `json:"-"`

	// Format: format for exporting the invoice list.
	// Default value: json
	Format ListFormat `json:"-"`

	// XAuthToken: user token.
	XAuthToken *string `json:"-"`
}

// ListInvoicesResponse: list invoices response.
type ListInvoicesResponse struct {
	// TotalCount: the total number of invoices.
	TotalCount uint32 `json:"total_count"`

	// Invoices: the paginated returned invoices.
	Invoices []*IndexInvoice `json:"invoices"`
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

// ListUsagesRequest: list usages request.
type ListUsagesRequest struct {
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`

	// OrganizationID: filter usages by organization ID.
	OrganizationID *string `json:"-"`

	// ProductID: filter usages by product ID.
	ProductID *string `json:"-"`

	// ResourceID: filter usages by resource ID.
	ResourceID *string `json:"-"`

	// Format: format for exporting the usage list.
	// Default value: json
	Format ListFormat `json:"-"`

	// XAuthToken: user token.
	XAuthToken *string `json:"-"`
}

// ListUsagesResponse: list usages response.
type ListUsagesResponse struct {
	// TotalCount: the total number of usages.
	TotalCount uint32 `json:"total_count"`

	// Usages: the paginated returned usages.
	Usages []*IndexUsage `json:"usages"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListUsagesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListUsagesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListUsagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Usages = append(r.Usages, results.Usages...)
	r.TotalCount += uint32(len(results.Usages))
	return uint32(len(results.Usages)), nil
}

type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ListDiscounts: List discounts.
func (s *API) ListDiscounts(req *ListDiscountsRequest, opts ...scw.RequestOption) (*ListDiscountsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "organization", req.Organization)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-admin/v1/discounts",
		Query:  query,
	}

	var resp ListDiscountsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInvoices: List invoices.
func (s *API) ListInvoices(req *ListInvoicesRequest, opts ...scw.RequestOption) (*ListInvoicesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "format", req.Format)
	parameter.AddToQuery(query, "x_auth_token", req.XAuthToken)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-admin/v1/invoices",
		Query:  query,
	}

	var resp ListInvoicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListUsages: List usages.
func (s *API) ListUsages(req *ListUsagesRequest, opts ...scw.RequestOption) (*ListUsagesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "product_id", req.ProductID)
	parameter.AddToQuery(query, "resource_id", req.ResourceID)
	parameter.AddToQuery(query, "format", req.Format)
	parameter.AddToQuery(query, "x_auth_token", req.XAuthToken)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-admin/v1/usages",
		Query:  query,
	}

	var resp ListUsagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

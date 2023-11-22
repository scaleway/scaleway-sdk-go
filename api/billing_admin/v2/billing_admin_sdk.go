// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package billing_admin provides methods and message types of the billing_admin v2 API.
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

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/billing_admin/v2/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/billing_admin/v2/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/billing_admin/v2/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/billing_admin/v2/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/billing_admin/v2/scw"
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

type CouponDiscountDiscountMode string

const (
	CouponDiscountDiscountModeUnknown    = CouponDiscountDiscountMode("unknown")
	CouponDiscountDiscountModeRate       = CouponDiscountDiscountMode("rate")
	CouponDiscountDiscountModeValueMode  = CouponDiscountDiscountMode("value_mode")
	CouponDiscountDiscountModeSplittable = CouponDiscountDiscountMode("splittable")
)

func (enum CouponDiscountDiscountMode) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum CouponDiscountDiscountMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CouponDiscountDiscountMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CouponDiscountDiscountMode(CouponDiscountDiscountMode(tmp).String())
	return nil
}

type InvoiceMetadataType string

const (
	InvoiceMetadataTypeUnknownType           = InvoiceMetadataType("unknown_type")
	InvoiceMetadataTypeExtExternalIdentifier = InvoiceMetadataType("ext_external_identifier")
	InvoiceMetadataTypeExtPurchaseOrder      = InvoiceMetadataType("ext_purchase_order")
	InvoiceMetadataTypeSalesContactName      = InvoiceMetadataType("sales_contact_name")
	InvoiceMetadataTypeSalesContactPhone     = InvoiceMetadataType("sales_contact_phone")
	InvoiceMetadataTypeSalesContactEmail     = InvoiceMetadataType("sales_contact_email")
	InvoiceMetadataTypeCatalogID             = InvoiceMetadataType("catalog_id")
	InvoiceMetadataTypePaymentTerm           = InvoiceMetadataType("payment_term")
)

func (enum InvoiceMetadataType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum InvoiceMetadataType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InvoiceMetadataType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InvoiceMetadataType(InvoiceMetadataType(tmp).String())
	return nil
}

type InvoiceState string

const (
	InvoiceStateUnknownState = InvoiceState("unknown_state")
	InvoiceStateDraft        = InvoiceState("draft")
	InvoiceStateOutdated     = InvoiceState("outdated")
	InvoiceStateStopped      = InvoiceState("stopped")
	InvoiceStateIncomplete   = InvoiceState("incomplete")
	InvoiceStateIssuing      = InvoiceState("issuing")
	InvoiceStateIssued       = InvoiceState("issued")
	InvoiceStatePaid         = InvoiceState("paid")
	InvoiceStateNodue        = InvoiceState("nodue")
	InvoiceStateCancelled    = InvoiceState("cancelled")
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

type ListCouponsRequestOrderBy string

const (
	ListCouponsRequestOrderByCreationDateDesc = ListCouponsRequestOrderBy("creation_date_desc")
	ListCouponsRequestOrderByCreationDateAsc  = ListCouponsRequestOrderBy("creation_date_asc")
	ListCouponsRequestOrderByStopDateAsc      = ListCouponsRequestOrderBy("stop_date_asc")
	ListCouponsRequestOrderByStopDateDesc     = ListCouponsRequestOrderBy("stop_date_desc")
	ListCouponsRequestOrderByStartDateAsc     = ListCouponsRequestOrderBy("start_date_asc")
	ListCouponsRequestOrderByStartDateDesc    = ListCouponsRequestOrderBy("start_date_desc")
	ListCouponsRequestOrderByCodeAsc          = ListCouponsRequestOrderBy("code_asc")
	ListCouponsRequestOrderByCodeDesc         = ListCouponsRequestOrderBy("code_desc")
	ListCouponsRequestOrderByActivationAsc    = ListCouponsRequestOrderBy("activation_asc")
	ListCouponsRequestOrderByActivationDesc   = ListCouponsRequestOrderBy("activation_desc")
)

func (enum ListCouponsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "creation_date_desc"
	}
	return string(enum)
}

func (enum ListCouponsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCouponsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCouponsRequestOrderBy(ListCouponsRequestOrderBy(tmp).String())
	return nil
}

// CouponDiscount: coupon discount.
type CouponDiscount struct {
	// StartDate: the start date, cannot be set with billing period.
	StartDate *time.Time `json:"start_date"`

	// StopDate: the end date, cannot be set with billing period.
	StopDate *time.Time `json:"stop_date"`

	// Value: desired value, percentage or absolute depending on the mode.
	Value *scw.Money `json:"value"`

	// Mode: discount's mode can be rate, value, or splittable.
	// Default value: unknown
	Mode CouponDiscountDiscountMode `json:"mode"`

	// Deprecated
	Description *string `json:"description,omitempty"`

	// BillingPeriodCount: the billing period range to which the discount is appliable.
	BillingPeriodCount *uint32 `json:"billing_period_count"`

	// Families: list of family keys to limit the usability of discounts.
	Families []string `json:"families"`

	// Products: list of products keys to limit the usability of discounts.
	Products []string `json:"products"`

	// Variants: list of variants keys to limit the usability of discounts.
	Variants []string `json:"variants"`

	// Ranges: list of ranges to limit the usability of discounts.
	Ranges []string `json:"ranges"`

	// Regions: list of regions where the discount applies.
	Regions []string `json:"regions"`

	// Zones: list of AZs where the discount applies.
	Zones []string `json:"zones"`

	// ExcludeFamilies: bool indicating whether we include or exclude the families.
	ExcludeFamilies bool `json:"exclude_families"`

	// ExcludeProducts: bool indicating whether we include or exclude the products.
	ExcludeProducts bool `json:"exclude_products"`

	// ExcludeVariants: bool indicating whether we include or exclude the variants.
	ExcludeVariants bool `json:"exclude_variants"`

	// ExcludeRanges: bool indicating whether we include or exclude the ranges.
	ExcludeRanges bool `json:"exclude_ranges"`

	// ExcludeRegions: bool indicating whether we include or exclude the regions.
	ExcludeRegions bool `json:"exclude_regions"`

	// ExcludeZones: bool indicating whether we include or exclude the zones.
	ExcludeZones bool `json:"exclude_zones"`
}

// CouponUpdateDiscountUpdate: coupon update discount update.
type CouponUpdateDiscountUpdate struct {
	// Deprecated
	Description *string `json:"description,omitempty"`

	// StartDate: cannot be set with the billing_period_count field.
	StartDate *time.Time `json:"start_date"`

	// StopDate: the end date, cannot be set with billing period.
	StopDate *time.Time `json:"stop_date"`

	// BillingPeriodCount: the billing period range to which the discount is appliable.
	BillingPeriodCount *uint32 `json:"billing_period_count"`

	// Families: list of family keys to limit the usability of discounts.
	Families []string `json:"families"`

	// Products: list of products keys to limit the usability of discounts.
	Products []string `json:"products"`

	// Variants: list of variants keys to limit the usability of discounts.
	Variants []string `json:"variants"`

	// Ranges: list of ranges to limit the usability of discounts.
	Ranges []string `json:"ranges"`

	// Regions: list of regions where the discount applies.
	Regions []string `json:"regions"`

	// Zones: list of AZs where the discount applies.
	Zones []string `json:"zones"`

	// ExcludeFamilies: bool indicating whether we include or exclude the families.
	ExcludeFamilies bool `json:"exclude_families"`

	// ExcludeProducts: bool indicating whether we include or exclude the products.
	ExcludeProducts bool `json:"exclude_products"`

	// ExcludeVariants: bool indicating whether we include or exclude the variants.
	ExcludeVariants bool `json:"exclude_variants"`

	// ExcludeRanges: bool indicating whether we include or exclude the ranges.
	ExcludeRanges bool `json:"exclude_ranges"`

	// ExcludeRegions: bool indicating whether we include or exclude the regions.
	ExcludeRegions bool `json:"exclude_regions"`

	// ExcludeZones: bool indicating whether we include or exclude the zones.
	ExcludeZones bool `json:"exclude_zones"`
}

// InvoiceMetadataValues: invoice metadata values.
type InvoiceMetadataValues struct {
	Values []string `json:"values"`
}

// Coupon: coupon.
type Coupon struct {
	// StartDate: coupon's start date.
	StartDate *time.Time `json:"start_date"`

	// StopDate: coupon's stop date.
	StopDate *time.Time `json:"stop_date"`

	// Description: the description that the end user will see.
	Description string `json:"description"`

	// InternalDescription: the description that the admin will see.
	InternalDescription string `json:"internal_description"`

	// Code: the code can only be alphanumeric.
	Code string `json:"code"`

	// Active: is the coupon currently usable.
	Active bool `json:"active"`

	// NewOrganizationOnly: apply only to organizations with no invoices.
	NewOrganizationOnly bool `json:"new_organization_only"`

	// Label: non-public label.
	Label string `json:"label"`

	// UsageLimit: counter to track the limit coupon's usages.
	UsageLimit uint32 `json:"usage_limit"`

	// ID: coupon's key.
	ID *string `json:"id"`

	Discount *CouponDiscount `json:"discount"`
}

// CouponUpdate: coupon update.
type CouponUpdate struct {
	// StartDate: coupon's start date.
	StartDate *time.Time `json:"start_date"`

	// StopDate: coupon's stop date.
	StopDate *time.Time `json:"stop_date"`

	// Description: the description that the end user will see.
	Description string `json:"description"`

	// InternalDescription: the description that the admin will see.
	InternalDescription string `json:"internal_description"`

	// NewOrganizationOnly: apply only to organizations with no invoices.
	NewOrganizationOnly bool `json:"new_organization_only"`

	// Label: non-public label.
	Label string `json:"label"`

	// UsageLimit: counter to track the limit coupon's usages.
	UsageLimit uint32 `json:"usage_limit"`

	// ID: coupon's key.
	ID *string `json:"id"`

	Discount *CouponUpdateDiscountUpdate `json:"discount"`
}

// InvoiceLine: invoice line.
type InvoiceLine struct {
	// ID: ID of invoice line.
	ID string `json:"id"`

	// InvoiceID: ID of related invoice.
	InvoiceID string `json:"invoice_id"`

	// ConsumerID: organization ID of the consumer.
	ConsumerID string `json:"consumer_id"`

	// ProjectID: the ID of the project.
	ProjectID string `json:"project_id"`

	// OperationPath: operation path of the service.
	OperationPath string `json:"operation_path"`

	// OperationName: operation name of the service.
	OperationName string `json:"operation_name"`

	// ResourceCount: the number of consumed resources.
	ResourceCount uint32 `json:"resource_count"`

	// StartDate: the start date of the invoice line.
	StartDate *time.Time `json:"start_date"`

	// StopDate: the stop date of the invoice line.
	StopDate *time.Time `json:"stop_date"`

	// UnitPrice: unit price.
	UnitPrice *scw.Money `json:"unit_price"`

	// UntaxedValue: the untaxed amount.
	UntaxedValue *scw.Money `json:"untaxed_value"`

	// Unit: the unit of the quantity.
	Unit string `json:"unit"`

	// ServiceCode: the servive code.
	ServiceCode string `json:"service_code"`
}

// Invoice: An invoice in the list.
type Invoice struct {
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

	// Email: the email of the invoice payer.
	Email string `json:"email"`

	// PhoneNumber: phone number of the invoice payer.
	PhoneNumber string `json:"phone_number"`

	// Locale: locale.
	Locale string `json:"locale"`

	// VatNumber: vat_number.
	VatNumber string `json:"vat_number"`

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

	// AddressLine1: address line of the invoice payer.
	AddressLine1 string `json:"address_line1"`

	// AddressLine2: address line 2 of the invoice payer.
	AddressLine2 string `json:"address_line2"`

	// AddressPostalCode: postal code of the invoice payer.
	AddressPostalCode string `json:"address_postal_code"`

	// AddressCityName: city of the invoice payer.
	AddressCityName string `json:"address_city_name"`

	// AddressCountryCode: country code of the invoice payer.
	AddressCountryCode string `json:"address_country_code"`

	// AddressSubdivisionCode: address subdivision code of the invoice payer.
	AddressSubdivisionCode string `json:"address_subdivision_code"`

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

	// DocumentFilename: filename of the invoice document.
	DocumentFilename string `json:"document_filename"`

	// DocumentChecksum: checksum of the invoice document.
	DocumentChecksum string `json:"document_checksum"`

	// PaymentMethod: the payment method of the invoice.
	PaymentMethod string `json:"payment_method"`

	// Metadata: some metadata related to the invoice.
	Metadata map[string]*InvoiceMetadataValues `json:"metadata"`
}

// ConsoleAPICreateCouponRequest: console api create coupon request.
type ConsoleAPICreateCouponRequest struct {
	Coupon *Coupon `json:"coupon,omitempty"`
}

// ConsoleAPIDeleteCouponRequest: console api delete coupon request.
type ConsoleAPIDeleteCouponRequest struct {
	CouponID string `json:"-"`
}

// ConsoleAPIListCouponsRequest: console api list coupons request.
type ConsoleAPIListCouponsRequest struct {
	// OrganizationID: when set display all coupons activable by the organization.
	OrganizationID string `json:"-"`

	// Page: number of the page.
	Page *int32 `json:"-"`

	// PageSize: size of the page.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the coupons are ordered in the response.
	// Default value: creation_date_desc
	OrderBy ListCouponsRequestOrderBy `json:"-"`

	// Activation: retrieve only activated/desactivated coupons.
	Activation *bool `json:"-"`

	// Label: filter coupons by label.
	Label *string `json:"-"`

	// Value: filter coupons by value.
	Value *string `json:"-"`
}

// ConsoleAPIUpdateCouponRequest: console api update coupon request.
type ConsoleAPIUpdateCouponRequest struct {
	Coupon *CouponUpdate `json:"coupon,omitempty"`
}

// CreateCouponResponse: create coupon response.
type CreateCouponResponse struct {
	Coupon *Coupon `json:"coupon"`
}

// CreateOrUpdateInvoiceMetadataRequest: create or update invoice metadata request.
type CreateOrUpdateInvoiceMetadataRequest struct {
	// PayerOrganizationID: the organization ID of the payer.
	PayerOrganizationID string `json:"payer_organization_id"`

	// ConsumerOrganizationID: the organization ID of the customer.
	ConsumerOrganizationID *string `json:"consumer_organization_id,omitempty"`

	// StartAt: the invoice metadata start date.
	StartAt *time.Time `json:"start_at,omitempty"`

	// Type: the invoice metdata type.
	// Default value: unknown_type
	Type InvoiceMetadataType `json:"type"`

	// Value: the invoice metadata value.
	Value string `json:"value"`

	// StopAt: the invoice metadata stop date.
	StopAt *time.Time `json:"stop_at,omitempty"`
}

// InvoiceMetadata: An Invoice Metadata.
type InvoiceMetadata struct {
	// Key: the Invoice Metadata id.
	Key string `json:"key"`

	// ConsumerOrganizationID: the organization ID of the customer.
	ConsumerOrganizationID *string `json:"consumer_organization_id"`

	// PayerOrganizationID: the organization ID of the payer.
	PayerOrganizationID *string `json:"payer_organization_id"`

	// Type: the Invoice Metadata type.
	// Default value: unknown_type
	Type InvoiceMetadataType `json:"type"`

	// Value: the Invoice Metadata value associated to the type.
	Value string `json:"value"`

	// StartAt: invoice Metadata start date.
	StartAt *time.Time `json:"start_at"`

	// StopAt: invoice Metadata stop date.
	StopAt *time.Time `json:"stop_at"`

	// CreatedAt: invoice Metadata creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: invoice Metadata update date.
	UpdatedAt *time.Time `json:"updated_at"`

	// DeletedAt: invoice Metadata deletion date.
	DeletedAt *time.Time `json:"deleted_at"`
}

// ListCouponsResponse: list coupons response.
type ListCouponsResponse struct {
	Coupons []*Coupon `json:"coupons"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCouponsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCouponsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCouponsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Coupons = append(r.Coupons, results.Coupons...)
	r.TotalCount += uint32(len(results.Coupons))
	return uint32(len(results.Coupons)), nil
}

// ListInvoiceLinesRequest: list invoice lines request.
type ListInvoiceLinesRequest struct {
	// InvoiceID: ID of the invoice.
	InvoiceID string `json:"-"`

	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: a positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// ConsumerID: the organization id of the consumer.
	ConsumerID *string `json:"-"`
}

// ListInvoiceLinesResponse: list invoice lines response.
type ListInvoiceLinesResponse struct {
	// TotalCount: the number of returned invoice lines.
	TotalCount uint32 `json:"total_count"`

	// Lines: returned invoice lines.
	Lines []*InvoiceLine `json:"lines"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInvoiceLinesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInvoiceLinesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListInvoiceLinesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Lines = append(r.Lines, results.Lines...)
	r.TotalCount += uint32(len(results.Lines))
	return uint32(len(results.Lines)), nil
}

// ListInvoicesRequest: list invoices request.
type ListInvoicesRequest struct {
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: a positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// OrganizationID: filter invoices by organization ID.
	OrganizationID *string `json:"-"`

	// StartedAfter: filter invoices where invoice's start_date is greater or equal to started_after.
	StartedAfter *time.Time `json:"-"`

	// StartedBefore: filter invoices where invoice's start_date is less to started_before.
	StartedBefore *time.Time `json:"-"`
}

// ListInvoicesResponse: list invoices response.
type ListInvoicesResponse struct {
	// TotalCount: the total number of invoices.
	TotalCount uint32 `json:"total_count"`

	// Invoices: the paginated returned invoices.
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

// UpdateCouponResponse: update coupon response.
type UpdateCouponResponse struct {
	Coupon *Coupon `json:"coupon"`
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

// ListInvoices: List invoices.
func (s *API) ListInvoices(req *ListInvoicesRequest, opts ...scw.RequestOption) (*ListInvoicesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "started_after", req.StartedAfter)
	parameter.AddToQuery(query, "started_before", req.StartedBefore)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-admin/v2/invoices",
		Query:  query,
	}

	var resp ListInvoicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInvoiceLines: List invoice lines.
func (s *API) ListInvoiceLines(req *ListInvoiceLinesRequest, opts ...scw.RequestOption) (*ListInvoiceLinesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "consumer_id", req.ConsumerID)

	if fmt.Sprint(req.InvoiceID) == "" {
		return nil, errors.New("field InvoiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-admin/v2/invoices/" + fmt.Sprint(req.InvoiceID) + "/lines",
		Query:  query,
	}

	var resp ListInvoiceLinesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOrUpdateInvoiceMetadata: Create an invoice metadata.
func (s *API) CreateOrUpdateInvoiceMetadata(req *CreateOrUpdateInvoiceMetadataRequest, opts ...scw.RequestOption) (*InvoiceMetadata, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/billing-admin/v2/invoice-metadata",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp InvoiceMetadata

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Billing APIs dedicated for the admin console.
type ConsoleAPI struct {
	client *scw.Client
}

// NewConsoleAPI returns a ConsoleAPI object from a Scaleway client.
func NewConsoleAPI(client *scw.Client) *ConsoleAPI {
	return &ConsoleAPI{
		client: client,
	}
}

// ListCoupons:
func (s *ConsoleAPI) ListCoupons(req *ConsoleAPIListCouponsRequest, opts ...scw.RequestOption) (*ListCouponsResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "activation", req.Activation)
	parameter.AddToQuery(query, "label", req.Label)
	parameter.AddToQuery(query, "value", req.Value)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-admin/v2/coupons",
		Query:  query,
	}

	var resp ListCouponsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCoupon:
func (s *ConsoleAPI) CreateCoupon(req *ConsoleAPICreateCouponRequest, opts ...scw.RequestOption) (*CreateCouponResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/billing-admin/v2/coupons",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateCouponResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCoupon:
func (s *ConsoleAPI) UpdateCoupon(req *ConsoleAPIUpdateCouponRequest, opts ...scw.RequestOption) (*UpdateCouponResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/billing-admin/v2/coupons",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateCouponResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCoupon:
func (s *ConsoleAPI) DeleteCoupon(req *ConsoleAPIDeleteCouponRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.CouponID) == "" {
		return errors.New("field CouponID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/billing-admin/v2/coupons/" + fmt.Sprint(req.CouponID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

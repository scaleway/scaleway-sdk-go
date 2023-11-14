// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package billing_admin provides methods and message types of the billing_admin v3 API.
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

type DiscountFilterType string

const (
	DiscountFilterTypeUnknownType = DiscountFilterType("unknown_type")
	DiscountFilterTypeFamily      = DiscountFilterType("family")
	DiscountFilterTypeProduct     = DiscountFilterType("product")
	DiscountFilterTypeVariant     = DiscountFilterType("variant")
	DiscountFilterTypeRange       = DiscountFilterType("range")
	DiscountFilterTypeRegion      = DiscountFilterType("region")
	DiscountFilterTypeZone        = DiscountFilterType("zone")
)

func (enum DiscountFilterType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
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

type DiscountMode string

const (
	DiscountModeUnknownMode = DiscountMode("unknown_mode")
	DiscountModeRate        = DiscountMode("rate")
	DiscountModeValueMode   = DiscountMode("value_mode")
	DiscountModeSplittable  = DiscountMode("splittable")
)

func (enum DiscountMode) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_mode"
	}
	return string(enum)
}

func (enum DiscountMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DiscountMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DiscountMode(DiscountMode(tmp).String())
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
	InvoiceMetadataTypeOrganizationName      = InvoiceMetadataType("organization_name")
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

type ListInvoicesMetadataRequestOrderBy string

const (
	ListInvoicesMetadataRequestOrderByStartDateDesc              = ListInvoicesMetadataRequestOrderBy("start_date_desc")
	ListInvoicesMetadataRequestOrderByStartDateAsc               = ListInvoicesMetadataRequestOrderBy("start_date_asc")
	ListInvoicesMetadataRequestOrderByConsumerOrganizationIDDesc = ListInvoicesMetadataRequestOrderBy("consumer_organization_id_desc")
	ListInvoicesMetadataRequestOrderByConsumerOrganizationIDAsc  = ListInvoicesMetadataRequestOrderBy("consumer_organization_id_asc")
	ListInvoicesMetadataRequestOrderByPayerOrganizationIDDesc    = ListInvoicesMetadataRequestOrderBy("payer_organization_id_desc")
	ListInvoicesMetadataRequestOrderByPayerOrganizationIDAsc     = ListInvoicesMetadataRequestOrderBy("payer_organization_id_asc")
)

func (enum ListInvoicesMetadataRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "start_date_desc"
	}
	return string(enum)
}

func (enum ListInvoicesMetadataRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInvoicesMetadataRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInvoicesMetadataRequestOrderBy(ListInvoicesMetadataRequestOrderBy(tmp).String())
	return nil
}

type ListInvoicesRequestOrderBy string

const (
	ListInvoicesRequestOrderByInvoiceNumberDesc    = ListInvoicesRequestOrderBy("invoice_number_desc")
	ListInvoicesRequestOrderByInvoiceNumberAsc     = ListInvoicesRequestOrderBy("invoice_number_asc")
	ListInvoicesRequestOrderByPaymentDateAsc       = ListInvoicesRequestOrderBy("payment_date_asc")
	ListInvoicesRequestOrderByPaymentDateDesc      = ListInvoicesRequestOrderBy("payment_date_desc")
	ListInvoicesRequestOrderByTotalUntaxedAsc      = ListInvoicesRequestOrderBy("total_untaxed_asc")
	ListInvoicesRequestOrderByTotalUntaxedDesc     = ListInvoicesRequestOrderBy("total_untaxed_desc")
	ListInvoicesRequestOrderByOrganizationNameAsc  = ListInvoicesRequestOrderBy("organization_name_asc")
	ListInvoicesRequestOrderByOrganizationNameDesc = ListInvoicesRequestOrderBy("organization_name_desc")
	ListInvoicesRequestOrderByOrganizationIDAsc    = ListInvoicesRequestOrderBy("organization_id_asc")
	ListInvoicesRequestOrderByOrganizationIDDesc   = ListInvoicesRequestOrderBy("organization_id_desc")
	ListInvoicesRequestOrderByTotalTaxedAsc        = ListInvoicesRequestOrderBy("total_taxed_asc")
	ListInvoicesRequestOrderByTotalTaxedDesc       = ListInvoicesRequestOrderBy("total_taxed_desc")
	ListInvoicesRequestOrderByBillingPeriodAsc     = ListInvoicesRequestOrderBy("billing_period_asc")
	ListInvoicesRequestOrderByBillingPeriodDesc    = ListInvoicesRequestOrderBy("billing_period_desc")
)

func (enum ListInvoicesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "invoice_number_desc"
	}
	return string(enum)
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

type ListInvoicesRequestPaymentMethod string

const (
	ListInvoicesRequestPaymentMethodUnknownPaymentMethod = ListInvoicesRequestPaymentMethod("unknown_payment_method")
	ListInvoicesRequestPaymentMethodCard                 = ListInvoicesRequestPaymentMethod("card")
	ListInvoicesRequestPaymentMethodSepa                 = ListInvoicesRequestPaymentMethod("sepa")
	ListInvoicesRequestPaymentMethodTransfer             = ListInvoicesRequestPaymentMethod("transfer")
)

func (enum ListInvoicesRequestPaymentMethod) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_payment_method"
	}
	return string(enum)
}

func (enum ListInvoicesRequestPaymentMethod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInvoicesRequestPaymentMethod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInvoicesRequestPaymentMethod(ListInvoicesRequestPaymentMethod(tmp).String())
	return nil
}

type ListInvoicesRequestPaymentState string

const (
	ListInvoicesRequestPaymentStateUnknownPaymentState = ListInvoicesRequestPaymentState("unknown_payment_state")
	ListInvoicesRequestPaymentStatePending             = ListInvoicesRequestPaymentState("pending")
	ListInvoicesRequestPaymentStateProcessing          = ListInvoicesRequestPaymentState("processing")
	ListInvoicesRequestPaymentStateClosed              = ListInvoicesRequestPaymentState("closed")
)

func (enum ListInvoicesRequestPaymentState) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_payment_state"
	}
	return string(enum)
}

func (enum ListInvoicesRequestPaymentState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInvoicesRequestPaymentState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInvoicesRequestPaymentState(ListInvoicesRequestPaymentState(tmp).String())
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

type ListSamplesRequestOrderBy string

const (
	ListSamplesRequestOrderByProbeTimeDesc = ListSamplesRequestOrderBy("probe_time_desc")
)

func (enum ListSamplesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "probe_time_desc"
	}
	return string(enum)
}

func (enum ListSamplesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSamplesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSamplesRequestOrderBy(ListSamplesRequestOrderBy(tmp).String())
	return nil
}

type SampleOrigin string

const (
	SampleOriginOriginUnknown = SampleOrigin("origin_unknown")
	SampleOriginOriginAPI     = SampleOrigin("origin_api")
	SampleOriginOriginCutoff  = SampleOrigin("origin_cutoff")
	SampleOriginOriginManual  = SampleOrigin("origin_manual")
)

func (enum SampleOrigin) String() string {
	if enum == "" {
		// return default value if empty
		return "origin_unknown"
	}
	return string(enum)
}

func (enum SampleOrigin) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SampleOrigin) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SampleOrigin(SampleOrigin(tmp).String())
	return nil
}

// InvoiceMetadataValues: invoice metadata values.
type InvoiceMetadataValues struct {
	Values []string `json:"values"`
}

// CreateDiscountRequest: create discount request.
type CreateDiscountRequest struct {
	// OrganizationID: ID of the organization to which the created discount applies.
	OrganizationID string `json:"organization_id"`

	// StartDate: date time aligned on the beginning of a month.
	StartDate *time.Time `json:"start_date,omitempty"`

	// StopDate: date time aligned on the end of a month.
	StopDate *time.Time `json:"stop_date,omitempty"`

	// PublicDescription: public description which will be displayed on the invoice and on the frontend.
	PublicDescription string `json:"public_description"`

	// InternalDescription: internal description which will be displayed on the admin console and on reports.
	InternalDescription string `json:"internal_description"`

	// Value: value carried by the discount, could be a currency value or a percentage value.
	Value *scw.Money `json:"value,omitempty"`

	// Mode: discount mode which can be value, splittable or rate.
	// Default value: unknown_mode
	Mode DiscountMode `json:"mode"`

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

// DiscountFilter: discount filter.
type DiscountFilter struct {
	ID string `json:"id"`

	// Type: default value: unknown_type
	Type DiscountFilterType `json:"type"`

	Value string `json:"value"`

	Exclude bool `json:"exclude"`
}

// Sample: A sample in the list.
type Sample struct {
	// Sku: the SKU of the sample.
	Sku string `json:"sku"`

	// ProjectID: the project ID of the sample.
	ProjectID string `json:"project_id"`

	// ResourceID: the resource ID of the sample.
	ResourceID string `json:"resource_id"`

	// ProbeTime: the time at which the the measurement was made.
	ProbeTime *time.Time `json:"probe_time"`

	// Quantity: the measured quantity of the sample.
	Quantity string `json:"quantity"`

	// Unit: the quantity unit of the sample.
	Unit string `json:"unit"`

	// Origin: the origin of the sample.
	// Default value: origin_unknown
	Origin SampleOrigin `json:"origin"`

	// CreatedAt: the time at which the sample was received.
	CreatedAt *time.Time `json:"created_at"`
}

// InvoiceMetadata: An Invoice Metadata.
type InvoiceMetadata struct {
	// ID: the Invoice Metadata id.
	ID string `json:"id"`

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

	// AccountingLineID: the accounting line related to the invoice.
	AccountingLineID string `json:"accounting_line_id"`
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
}

// CreateAutomatedDiscountRequest: create automated discount request.
type CreateAutomatedDiscountRequest struct {
	// Request: a CreateDiscountRequest.
	Request *CreateDiscountRequest `json:"request,omitempty"`
}

// CreateDiscountFilterRequest: create discount filter request.
type CreateDiscountFilterRequest struct {
	// DiscountID: the id of the associated discount.
	DiscountID string `json:"discount_id"`

	// Type: the discount filter type.
	// Default value: unknown_type
	Type DiscountFilterType `json:"type"`

	// Value: the id or name to filter.
	Value string `json:"value"`

	// Exclude: a boolean that indicates if we exclude or not the filtered value.
	Exclude bool `json:"exclude"`
}

// DeleteDiscountFilterRequest: delete discount filter request.
type DeleteDiscountFilterRequest struct {
	// DiscountFilterID: the id of the discount filter to delete.
	DiscountFilterID string `json:"-"`
}

// DeleteInvoiceMetadataRequest: delete invoice metadata request.
type DeleteInvoiceMetadataRequest struct {
	// InvoiceMetadataID: the ID of invoice metadata to delete.
	InvoiceMetadataID string `json:"-"`
}

// Discount: discount.
type Discount struct {
	// ID: date time aligned on the beginning of a month.
	ID string `json:"id"`

	// OrganizationID: ID of the organization to which the created discount applies.
	OrganizationID string `json:"organization_id"`

	// StartDate: date time aligned on the beginning of a month.
	StartDate *time.Time `json:"start_date"`

	// StopDate: date time aligned on the end of a month.
	StopDate *time.Time `json:"stop_date"`

	// PublicDescription: public description which will be displayed on the invoice and on the frontend.
	PublicDescription string `json:"public_description"`

	// InternalDescription: internal description which will be displayed on the admin console and on reports.
	InternalDescription string `json:"internal_description"`

	// Value: value carried by the discount, could be a currency value or a percentage value.
	Value *scw.Money `json:"value"`

	// Mode: discount mode which can be value, splittable or rate.
	// Default value: unknown_mode
	Mode DiscountMode `json:"mode"`

	// Filters: list of discount filters to limit the usability of discounts.
	Filters []*DiscountFilter `json:"filters"`
}

// ListActiveOrganizationsRequest: list active organizations request.
type ListActiveOrganizationsRequest struct {
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: a positive integer to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// Since: filter active organizations since this date.
	Since *time.Time `json:"-"`
}

// ListActiveOrganizationsResponse: list active organizations response.
type ListActiveOrganizationsResponse struct {
	// TotalCount: total count of active organizations.
	TotalCount uint32 `json:"total_count"`

	// OrganizationIDs: list of active organization IDs.
	OrganizationIDs []string `json:"organization_ids"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListActiveOrganizationsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListActiveOrganizationsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListActiveOrganizationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.OrganizationIDs = append(r.OrganizationIDs, results.OrganizationIDs...)
	r.TotalCount += uint32(len(results.OrganizationIDs))
	return uint32(len(results.OrganizationIDs)), nil
}

// ListInvoicesMetadataRequest: list invoices metadata request.
type ListInvoicesMetadataRequest struct {
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: a positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// ConsumerOrganizationID: the organization ID of the customer.
	ConsumerOrganizationID *string `json:"-"`

	// PayerOrganizationID: the organization ID of the payer.
	PayerOrganizationID *string `json:"-"`

	// Value: returns only the invoice metadata with this value.
	Value *string `json:"-"`

	// Type: returns only the invoice metadata with this type.
	Type *string `json:"-"`

	// OrderBy: how the invoice metadata are ordered in the response.
	// Default value: start_date_desc
	OrderBy ListInvoicesMetadataRequestOrderBy `json:"-"`
}

// ListInvoicesMetadataResponse: list invoices metadata response.
type ListInvoicesMetadataResponse struct {
	// TotalCount: total count of invoice metadata.
	TotalCount uint32 `json:"total_count"`

	// InvoiceMetadata: list of returned invoice metadata.
	InvoiceMetadata []*InvoiceMetadata `json:"invoice_metadata"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInvoicesMetadataResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInvoicesMetadataResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListInvoicesMetadataResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.InvoiceMetadata = append(r.InvoiceMetadata, results.InvoiceMetadata...)
	r.TotalCount += uint32(len(results.InvoiceMetadata))
	return uint32(len(results.InvoiceMetadata)), nil
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

	// StartedBefore: filter invoices where invoice's start_date is before started_before.
	StartedBefore *time.Time `json:"-"`

	// PaymentMethod: filter invoices by the payment method used, card, sepa or wire transfer.
	// Default value: unknown_payment_method
	PaymentMethod ListInvoicesRequestPaymentMethod `json:"-"`

	// PaymentState: filter invoices by the current state of the payment, pending, processing or closed.
	// Default value: unknown_payment_state
	PaymentState ListInvoicesRequestPaymentState `json:"-"`

	// InvoiceNumber: filter to find the invoice corresponding to this invoice number.
	InvoiceNumber *int32 `json:"-"`

	// PaymentDate: filter invoices by their payment date.
	PaymentDate *time.Time `json:"-"`

	// Amount: filter invoices where their amounts are greater than or equal to this amount.
	Amount *float64 `json:"-"`

	// OrganizationName: filter invoices to find the name of the invoice payer.
	OrganizationName *string `json:"-"`

	// OrderBy: how invoices are ordered in the response.
	// Default value: invoice_number_desc
	OrderBy ListInvoicesRequestOrderBy `json:"-"`

	// BillingPeriod: filter invoices by billing_period.
	BillingPeriod *string `json:"-"`

	// State: filter invoices according to their state.
	// Default value: unknown_state
	State InvoiceState `json:"-"`
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

// ListSamplesResponse: list samples response.
type ListSamplesResponse struct {
	// TotalCount: the total number of samples.
	TotalCount uint64 `json:"total_count"`

	// Samples: the paginated returned samples.
	Samples []*Sample `json:"samples"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSamplesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSamplesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListSamplesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Samples = append(r.Samples, results.Samples...)
	r.TotalCount += uint64(len(results.Samples))
	return uint64(len(results.Samples)), nil
}

// OfferAPICreateOfferActivationRequest: offer api create offer activation request.
type OfferAPICreateOfferActivationRequest struct {
	// OrganizationID: organization_id of the offer activation.
	OrganizationID string `json:"-"`

	// OfferSku: sku of the offer.
	OfferSku string `json:"-"`
}

// OfferAPIGetOfferEligibilityRequest: offer api get offer eligibility request.
type OfferAPIGetOfferEligibilityRequest struct {
	// OrganizationID: the organization that is tested.
	OrganizationID string `json:"-"`

	// OfferSku: the offer that is tested.
	OfferSku string `json:"-"`
}

// OfferAPIListOfferActivationsRequest: offer api list offer activations request.
type OfferAPIListOfferActivationsRequest struct {
	// OrderBy: the sort order for the returned offer activations.
	// Default value: activation_date_desc
	OrderBy ListOfferActivationsRequestOrderBy `json:"-"`

	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: a positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// OrganizationID: organization ID to filter for.
	OrganizationID *string `json:"-"`

	// OfferSku: offer sku to filter for.
	OfferSku *string `json:"-"`
}

// OfferEligibility: offer eligibility.
type OfferEligibility struct {
	// EligibilityState: the state of the elegibility of the organization to the offer at the given date.
	EligibilityState bool `json:"eligibility_state"`
}

// SampleAPIIgnoreSamplesRequest: sample api ignore samples request.
type SampleAPIIgnoreSamplesRequest struct {
	// Samples: the samples to be ignored. This is limited a maximum of 1000 samples at a time.
	Samples []*Sample `json:"samples"`
}

// SampleAPIListSamplesRequest: sample api list samples request.
type SampleAPIListSamplesRequest struct {
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: a positive integer lower or equal to 1000 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned samples.
	// Default value: probe_time_desc
	OrderBy ListSamplesRequestOrderBy `json:"-"`

	// ProjectIDs: project IDs to filter for. The results will be limited to any project with an ID in this array.
	ProjectIDs *[]string `json:"-"`

	// ProbeTimeAfter: only keep samples whose probe times are after this time. This boundary is included in the interval.
	ProbeTimeAfter *time.Time `json:"-"`

	// ProbeTimeBefore: only keep samples whose probe times are before this time. This boundary is included in the interval.
	ProbeTimeBefore *time.Time `json:"-"`

	// Skus: sKU to filter for. The results will be limited to any sample with an SKU specified in this array.
	Skus *[]string `json:"-"`

	// ResourceIDs: resource IDs to filter for. The results will be limited to any resource with an ID in this array.
	ResourceIDs *[]string `json:"-"`

	// Origins: origins to filter for. The results will be limited to any origin specified in this array.
	Origins []SampleOrigin `json:"-"`
}

// UpdateDiscountFilterRequest: update discount filter request.
type UpdateDiscountFilterRequest struct {
	// DiscountFilterID: the id of the discount filter to update.
	DiscountFilterID string `json:"-"`

	// Type: the discount filter type.
	// Default value: unknown_type
	Type DiscountFilterType `json:"type"`

	// Value: the id or name to filter.
	Value *string `json:"value,omitempty"`

	// Exclude: a boolean that indicates if we exclude or not the filtered value.
	Exclude *bool `json:"exclude,omitempty"`
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
	parameter.AddToQuery(query, "payment_method", req.PaymentMethod)
	parameter.AddToQuery(query, "payment_state", req.PaymentState)
	parameter.AddToQuery(query, "invoice_number", req.InvoiceNumber)
	parameter.AddToQuery(query, "payment_date", req.PaymentDate)
	parameter.AddToQuery(query, "amount", req.Amount)
	parameter.AddToQuery(query, "organization_name", req.OrganizationName)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "billing_period", req.BillingPeriod)
	parameter.AddToQuery(query, "state", req.State)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-admin/v3/invoices",
		Query:  query,
	}

	var resp ListInvoicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDiscount: Create a new discount.
func (s *API) CreateDiscount(req *CreateDiscountRequest, opts ...scw.RequestOption) (*Discount, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/billing-admin/v3/discounts",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Discount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAutomatedDiscount: Create automated discount.
func (s *API) CreateAutomatedDiscount(req *CreateAutomatedDiscountRequest, opts ...scw.RequestOption) (*Discount, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/billing-admin/v3/automated-discounts",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Discount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDiscountFilter: Create a discount filter.
func (s *API) CreateDiscountFilter(req *CreateDiscountFilterRequest, opts ...scw.RequestOption) (*DiscountFilter, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/billing-admin/v3/discount-filters",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DiscountFilter

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDiscountFilter: Update a discount filter.
func (s *API) UpdateDiscountFilter(req *UpdateDiscountFilterRequest, opts ...scw.RequestOption) (*DiscountFilter, error) {
	var err error

	if fmt.Sprint(req.DiscountFilterID) == "" {
		return nil, errors.New("field DiscountFilterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/billing-admin/v3/discount-filters/" + fmt.Sprint(req.DiscountFilterID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DiscountFilter

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDiscountFilter: Delete a discount filter.
func (s *API) DeleteDiscountFilter(req *DeleteDiscountFilterRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.DiscountFilterID) == "" {
		return errors.New("field DiscountFilterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/billing-admin/v3/discount-filters/" + fmt.Sprint(req.DiscountFilterID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListInvoicesMetadata: List invoice metadata.
func (s *API) ListInvoicesMetadata(req *ListInvoicesMetadataRequest, opts ...scw.RequestOption) (*ListInvoicesMetadataResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "consumer_organization_id", req.ConsumerOrganizationID)
	parameter.AddToQuery(query, "payer_organization_id", req.PayerOrganizationID)
	parameter.AddToQuery(query, "value", req.Value)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-admin/v3/invoices-metadata",
		Query:  query,
	}

	var resp ListInvoicesMetadataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteInvoiceMetadata: Delete an invoice metadata.
func (s *API) DeleteInvoiceMetadata(req *DeleteInvoiceMetadataRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.InvoiceMetadataID) == "" {
		return errors.New("field InvoiceMetadataID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/billing-admin/v3/invoices-metadata/" + fmt.Sprint(req.InvoiceMetadataID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListActiveOrganizations: List active organizations.
func (s *API) ListActiveOrganizations(req *ListActiveOrganizationsRequest, opts ...scw.RequestOption) (*ListActiveOrganizationsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "since", req.Since)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-admin/v3/active-organizations",
		Query:  query,
	}

	var resp ListActiveOrganizationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type OfferAPI struct {
	client *scw.Client
}

// NewOfferAPI returns a OfferAPI object from a Scaleway client.
func NewOfferAPI(client *scw.Client) *OfferAPI {
	return &OfferAPI{
		client: client,
	}
}

// CreateOfferActivation: Create offer activation.
func (s *OfferAPI) CreateOfferActivation(req *OfferAPICreateOfferActivationRequest, opts ...scw.RequestOption) (*OfferActivation, error) {
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
		Path:   "/billing-admin/v3/offer-activations",
		Query:  query,
	}

	var resp OfferActivation

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOfferActivations: List offer activations.
func (s *OfferAPI) ListOfferActivations(req *OfferAPIListOfferActivationsRequest, opts ...scw.RequestOption) (*ListOfferActivationsResponse, error) {
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
	parameter.AddToQuery(query, "offer_sku", req.OfferSku)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-admin/v3/offer-activations",
		Query:  query,
	}

	var resp ListOfferActivationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOfferEligibility: Get offer eligibility.
func (s *OfferAPI) GetOfferEligibility(req *OfferAPIGetOfferEligibilityRequest, opts ...scw.RequestOption) (*OfferEligibility, error) {
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
		Path:   "/billing-admin/v3/offer-eligibility",
		Query:  query,
	}

	var resp OfferEligibility

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SampleAPI struct {
	client *scw.Client
}

// NewSampleAPI returns a SampleAPI object from a Scaleway client.
func NewSampleAPI(client *scw.Client) *SampleAPI {
	return &SampleAPI{
		client: client,
	}
}
func (s *SampleAPI) Regions() []scw.Region {
	return []scw.Region{}
}

// ListSamples: List samples.
func (s *SampleAPI) ListSamples(req *SampleAPIListSamplesRequest, opts ...scw.RequestOption) (*ListSamplesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_ids", req.ProjectIDs)
	parameter.AddToQuery(query, "probe_time_after", req.ProbeTimeAfter)
	parameter.AddToQuery(query, "probe_time_before", req.ProbeTimeBefore)
	parameter.AddToQuery(query, "skus", req.Skus)
	parameter.AddToQuery(query, "resource_ids", req.ResourceIDs)
	parameter.AddToQuery(query, "origins", req.Origins)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing-admin/v3/samples",
		Query:  query,
	}

	var resp ListSamplesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// IgnoreSamples: Ignore samples.
func (s *SampleAPI) IgnoreSamples(req *SampleAPIIgnoreSamplesRequest, opts ...scw.RequestOption) error {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/billing-admin/v3/samples/ignore",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

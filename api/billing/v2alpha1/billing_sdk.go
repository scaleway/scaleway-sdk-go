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

// API: this API allows you to query your consumption.
// Billing API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
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

// GetConsumptionResponse: get consumption response.
type GetConsumptionResponse struct {
	// Consumptions: detailed consumption list.
	Consumptions []*GetConsumptionResponseConsumption `json:"consumptions"`
	// UpdatedAt: last consumption update date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// GetConsumptionResponseConsumption: get consumption response. consumption.
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

// ListInvoicesResponse: list invoices response.
type ListInvoicesResponse struct {
	// TotalCount: total number of invoices.
	TotalCount uint32 `json:"total_count"`
	// Invoices: paginated returned invoices.
	Invoices []*Invoice `json:"invoices"`
}

// Service API

type GetConsumptionRequest struct {
	// OrganizationID: filter by organization ID.
	OrganizationID string `json:"-"`
}

// GetConsumption: get current month's consumption.
// The consumption reflects the amount of money you have spent for the products you have used.
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
		Method:  "GET",
		Path:    "/billing/v2alpha1/consumption",
		Query:   query,
		Headers: http.Header{},
	}

	var resp GetConsumptionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

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
	// Default value: 20
	PageSize *uint32 `json:"-"`
	// OrderBy: how invoices are ordered in the response.
	// Default value: invoice_number_desc
	OrderBy ListInvoicesRequestOrderBy `json:"-"`
}

// ListInvoices: list invoices.
// List all your invoices, filtering by `start_date` and `invoice_type`. Each invoice has its own ID.
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
		Method:  "GET",
		Path:    "/billing/v2alpha1/invoices",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListInvoicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DownloadInvoiceRequest struct {
	// InvoiceID: invoice ID.
	InvoiceID string `json:"-"`
	// FileType: wanted file type.
	// Default value: pdf
	FileType DownloadInvoiceRequestFileType `json:"-"`
}

// DownloadInvoice: download an invoice.
// Download a specific invoice, specified by its ID.
func (s *API) DownloadInvoice(req *DownloadInvoiceRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "file_type", req.FileType)

	if fmt.Sprint(req.InvoiceID) == "" {
		return nil, errors.New("field InvoiceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/billing/v2alpha1/invoices/" + fmt.Sprint(req.InvoiceID) + "/download",
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

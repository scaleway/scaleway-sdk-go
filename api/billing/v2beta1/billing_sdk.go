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

type ListConsumptionsRequestOrderBy string

const (
	// Order consumptions by update date (descending chronological order).
	ListConsumptionsRequestOrderByUpdatedAtDateDesc = ListConsumptionsRequestOrderBy("updated_at_date_desc")
	// Order consumptions by update date (ascending chronological order).
	ListConsumptionsRequestOrderByUpdatedAtDateAsc = ListConsumptionsRequestOrderBy("updated_at_date_asc")
	// Order consumptions by category name (descending alphabetical order).
	ListConsumptionsRequestOrderByCategoryNameDesc = ListConsumptionsRequestOrderBy("category_name_desc")
	// Order consumptions by category name (ascending alphabetical order).
	ListConsumptionsRequestOrderByCategoryNameAsc = ListConsumptionsRequestOrderBy("category_name_asc")
)

func (enum ListConsumptionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "updated_at_date_desc"
	}
	return string(enum)
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

type ListTaxesRequestOrderBy string

const (
	// Order consumptions by update date (descending chronological order).
	ListTaxesRequestOrderByUpdatedAtDateDesc = ListTaxesRequestOrderBy("updated_at_date_desc")
	// Order consumptions by update date (ascending chronological order).
	ListTaxesRequestOrderByUpdatedAtDateAsc = ListTaxesRequestOrderBy("updated_at_date_asc")
	// Order consumptions by category name (descending alphabetical order).
	ListTaxesRequestOrderByCategoryNameDesc = ListTaxesRequestOrderBy("category_name_desc")
	// Order consumptions by category name (ascending alphabetical order).
	ListTaxesRequestOrderByCategoryNameAsc = ListTaxesRequestOrderBy("category_name_asc")
)

func (enum ListTaxesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "updated_at_date_desc"
	}
	return string(enum)
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

// ListConsumptionsRequest: list consumptions request.
type ListConsumptionsRequest struct {
	// OrderBy: order consumptions list in the response by their update date.
	// Default value: updated_at_date_desc
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

// ListTaxesRequest: list taxes request.
type ListTaxesRequest struct {
	// OrderBy: order consumed taxes list in the response by their update date.
	// Default value: updated_at_date_desc
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

// This API allows you to query your consumption.
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

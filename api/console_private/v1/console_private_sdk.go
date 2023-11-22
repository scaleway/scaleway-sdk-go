// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package console_private provides methods and message types of the console_private v1 API.
package console_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

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

type ListCurrentQuotasRequestOrderBy string

const (
	ListCurrentQuotasRequestOrderByNameAsc  = ListCurrentQuotasRequestOrderBy("name_asc")
	ListCurrentQuotasRequestOrderByNameDesc = ListCurrentQuotasRequestOrderBy("name_desc")
)

func (enum ListCurrentQuotasRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListCurrentQuotasRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCurrentQuotasRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCurrentQuotasRequestOrderBy(ListCurrentQuotasRequestOrderBy(tmp).String())
	return nil
}

type ListProductsRequestOrderBy string

const (
	ListProductsRequestOrderByNameAsc  = ListProductsRequestOrderBy("name_asc")
	ListProductsRequestOrderByNameDesc = ListProductsRequestOrderBy("name_desc")
)

func (enum ListProductsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListProductsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListProductsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListProductsRequestOrderBy(ListProductsRequestOrderBy(tmp).String())
	return nil
}

// Locality: locality.
type Locality struct {
	// Region: region to target. If none is passed will use default region from the config.
	// Precisely one of Region, Zone must be set.
	Region *scw.Region `json:"region,omitempty"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	// Precisely one of Region, Zone must be set.
	Zone *scw.Zone `json:"zone,omitempty"`
}

// LocalityCurrentCount: locality current count.
type LocalityCurrentCount struct {
	// Locality: regions or AZ.
	Locality *Locality `json:"locality"`

	// Count: number of existing resources.
	Count uint64 `json:"count"`

	// HasError: true if an error occured while retrieving the number of resources.
	HasError bool `json:"has_error"`
}

// FeatureQuota: feature quota.
type FeatureQuota struct {
	// FeatureKey: key describing the product feature.
	FeatureKey string `json:"feature_key"`

	// Limit: product quotas. -1 means unlimited.
	Limit int64 `json:"limit"`

	// CurrentCounts: number of existing resources grouped by locality.
	CurrentCounts []*LocalityCurrentCount `json:"current_counts"`
}

// Quota: quota.
type Quota struct {
	// FeatureQuotas: product quotas by feature.
	FeatureQuotas []*FeatureQuota `json:"feature_quotas"`
}

// Localities: localities.
type Localities struct {
	Localities []*Locality `json:"localities"`
}

// ListCurrentQuotasResponse: list current quotas response.
type ListCurrentQuotasResponse struct {
	// TotalCount: total quota count.
	TotalCount uint32 `json:"total_count"`

	// QuotaList: quotas and count of existing resources grouped by product.
	QuotaList map[string]*Quota `json:"quota_list"`
}

// ListProductsResponse: list products response.
type ListProductsResponse struct {
	// TotalCount: total count of products.
	TotalCount uint32 `json:"total_count"`

	// Products: product aliases and their localities (regions or AZ).
	Products map[string]*Localities `json:"products"`
}

// ProductAPIListProductsRequest: product api list products request.
type ProductAPIListProductsRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: ascending (asc) or descending (desc) order.
	// Default value: name_asc
	OrderBy ListProductsRequestOrderBy `json:"-"`
}

// QuotaAPIListCurrentQuotasRequest: quota api list current quotas request.
type QuotaAPIListCurrentQuotasRequest struct {
	// Page: the number of page to fetch.
	Page *int32 `json:"-"`

	// PageSize: the number of items per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: ascending (asc) or descending (desc) order.
	// Default value: name_asc
	OrderBy ListCurrentQuotasRequestOrderBy `json:"-"`

	// ProjectID: project quotas. Exclude organization_id.
	ProjectID string `json:"-"`

	// Products: show quota for a selected list of products (names comma-separated).
	Products *string `json:"-"`

	// NoResourceCounting: by default for each quota returned the application fetch the number of consumed resources of the same type per region.
	//
	// When It's set to true this feature can be disabled.
	NoResourceCounting bool `json:"-"`
}

// Provides a list of available products to resellers and partners.
type ProductAPI struct {
	client *scw.Client
}

// NewProductAPI returns a ProductAPI object from a Scaleway client.
func NewProductAPI(client *scw.Client) *ProductAPI {
	return &ProductAPI{
		client: client,
	}
}

// ListProducts: Returns the list of available products and their regions or AZs. Useful for filtering ListCurrentQuotas responses.
func (s *ProductAPI) ListProducts(req *ProductAPIListProductsRequest, opts ...scw.RequestOption) (*ListProductsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/console-private/v1/products",
		Query:  query,
	}

	var resp ListProductsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Provides quotas and the number of existing resources by Organization or Project to resellers and partners.
type QuotaAPI struct {
	client *scw.Client
}

// NewQuotaAPI returns a QuotaAPI object from a Scaleway client.
func NewQuotaAPI(client *scw.Client) *QuotaAPI {
	return &QuotaAPI{
		client: client,
	}
}

// ListCurrentQuotas: Returns quota by product. The response value includes the product's overall quota and the number of existing resources by locality (ie regions or AZ).
func (s *QuotaAPI) ListCurrentQuotas(req *QuotaAPIListCurrentQuotasRequest, opts ...scw.RequestOption) (*ListCurrentQuotasResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "products", req.Products)
	parameter.AddToQuery(query, "no_resource_counting", req.NoResourceCounting)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/console-private/v1/current-quotas",
		Query:  query,
	}

	var resp ListCurrentQuotasResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

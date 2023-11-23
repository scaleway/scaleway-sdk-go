// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package product_marketplace_private provides methods and message types of the product_marketplace_private v1 API.
package product_marketplace_private

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

// Value: value.
type Value struct {
	// Name: description of the value.
	Name string `json:"name"`

	// Value: value.
	Value string `json:"value"`
}

// Specification: specification.
type Specification struct {
	// Name: name of the specification.
	Name string `json:"name"`

	// Value: value of the specification.
	Value []*Value `json:"value"`
}

// Specifications: specifications.
type Specifications struct {
	// Variant: additional information of the variant.
	Variant []*Specification `json:"variant"`

	// Range: additional information of the range.
	Range []*Specification `json:"range"`

	// Product: additional information of the product.
	Product []*Specification `json:"product"`

	// Availability: additional information of the availability.
	Availability []*Specification `json:"availability"`
}

// Variant: variant.
type Variant struct {
	// VariantID: unique identifier of the variant.
	VariantID string `json:"variant_id"`

	// TimeUnit: unit of time size.
	TimeUnit string `json:"time_unit"`

	// TimeSize: period of time (or slice of time) used to split and compute the price of the usage.
	TimeSize int32 `json:"time_size"`

	// CountableUnit: unit of countable_size.
	CountableUnit string `json:"countable_unit"`

	// CountableSize: quantity used to split and compute the price of the usage.
	CountableSize int32 `json:"countable_size"`

	// Price: monetary value of the variant.
	Price *scw.Money `json:"price"`

	// Product: name of the product.
	Product string `json:"product"`

	// Availability: availability zone in which the product resides. Prices may differ depending on the availability zone. An availability can be eather a zone (i.e. fr-par-1) or a region (i.e. nl-ams) or global for worldwide products.
	Availability string `json:"availability"`

	// Range: the range regroups products of the same type but with variations like different configurations of options.
	Range string `json:"range"`
}

// ConsoleAPIListDynamicPricesRequest: console api list dynamic prices request.
type ConsoleAPIListDynamicPricesRequest struct {
	// SiblingKey: sibling key.
	SiblingKey string `json:"-"`
}

// ListDynamicPricesResponse: list dynamic prices response.
type ListDynamicPricesResponse struct {
	// Data: list of related variants.
	Data []*Variant `json:"data"`

	// Specifications: range, product and zone specifications.
	Specifications []*Specifications `json:"specifications"`
}

// Product Marketplace API is used to display product information in the Console.
type ConsoleAPI struct {
	client *scw.Client
}

// NewConsoleAPI returns a ConsoleAPI object from a Scaleway client.
func NewConsoleAPI(client *scw.Client) *ConsoleAPI {
	return &ConsoleAPI{
		client: client,
	}
}

// ListDynamicPrices:
func (s *ConsoleAPI) ListDynamicPrices(req *ConsoleAPIListDynamicPricesRequest, opts ...scw.RequestOption) (*ListDynamicPricesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "sibling_key", req.SiblingKey)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/product-marketplace-private/v1/dynamic-prices",
		Query:  query,
	}

	var resp ListDynamicPricesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

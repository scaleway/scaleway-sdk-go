// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package product_catalog provides methods and message types of the product_catalog v2alpha1 API.
package product_catalog

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

type DurationUnit string

const (
	DurationUnitUnknownUnit = DurationUnit("unknown_unit")
	DurationUnitHour        = DurationUnit("hour")
	DurationUnitDay         = DurationUnit("day")
	DurationUnitMonth       = DurationUnit("month")
)

func (enum DurationUnit) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_unit"
	}
	return string(enum)
}

func (enum DurationUnit) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DurationUnit) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DurationUnit(DurationUnit(tmp).String())
	return nil
}

// Duration: duration.
type Duration struct {
	Value uint32 `json:"value"`

	// Unit: default value: unknown_unit
	Unit DurationUnit `json:"unit"`
}

// EstimateCostRequestItem: estimate cost request item.
type EstimateCostRequestItem struct {
	Sku string `json:"sku"`

	Quantity uint64 `json:"quantity"`

	Multiplier uint32 `json:"multiplier"`
}

// EstimateCostResponseItem: estimate cost response item.
type EstimateCostResponseItem struct {
	Price *scw.Money `json:"price"`
}

// EstimateCostRequest: estimate cost request.
type EstimateCostRequest struct {
	Duration *Duration `json:"duration"`

	Currency string `json:"currency"`

	Items []*EstimateCostRequestItem `json:"items"`

	PriceRounded bool `json:"price_rounded"`
}

// EstimateCostResponse: estimate cost response.
type EstimateCostResponse struct {
	Items []*EstimateCostResponseItem `json:"items"`
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

// EstimateCost:
func (s *API) EstimateCost(req *EstimateCostRequest, opts ...scw.RequestOption) (*EstimateCostResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/product-catalog/v2alpha1/estimate-cost",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EstimateCostResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

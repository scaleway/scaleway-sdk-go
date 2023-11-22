// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package it_automation provides methods and message types of the it_automation v1 API.
package it_automation

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

// Card: card.
type Card struct {
	// Bin: the bin of the card.
	Bin string `json:"bin"`

	// OrganizationID: the id of the organization owning the card.
	OrganizationID string `json:"organization_id"`

	// Locked: wether the organization owning the card is locked.
	Locked *bool `json:"locked"`

	// LockedAt: the date at which the organization owning the card has been (or will be) locked.
	LockedAt *time.Time `json:"locked_at"`

	// OrganizationCreatedAt: the creation date of the organization owning the card.
	OrganizationCreatedAt *time.Time `json:"organization_created_at"`

	// OrganizationCountryCode: the country in which the organization owning the card is located.
	OrganizationCountryCode string `json:"organization_country_code"`

	// CardCountryCode: the country of the card.
	CardCountryCode string `json:"card_country_code"`

	// Revenue: the revenue that the organization owning the card has generated.
	Revenue *scw.Money `json:"revenue"`

	// Discounted: the amount the organization owning the card has been discounted.
	Discounted *scw.Money `json:"discounted"`
}

// CardAPIListCardsRequest: card api list cards request.
type CardAPIListCardsRequest struct {
	// Bin: the BIN the cards must match.
	Bin string `json:"-"`

	// CountryCode: the country the card come from (default null).
	CountryCode string `json:"-"`

	// Page: the page number.
	Page *int32 `json:"-"`

	// PageSize: the number of items per page.
	PageSize *uint32 `json:"-"`
}

// ListCardsResponse: list cards response.
type ListCardsResponse struct {
	// Cards: the list of cards requested.
	Cards []*Card `json:"cards"`

	// TotalCount: the number of cards.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCardsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCardsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListCardsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Cards = append(r.Cards, results.Cards...)
	r.TotalCount += uint64(len(results.Cards))
	return uint64(len(results.Cards)), nil
}

// An API giving acces to credit cards information.
type CardAPI struct {
	client *scw.Client
}

// NewCardAPI returns a CardAPI object from a Scaleway client.
func NewCardAPI(client *scw.Client) *CardAPI {
	return &CardAPI{
		client: client,
	}
}

// ListCards: List cards.
func (s *CardAPI) ListCards(req *CardAPIListCardsRequest, opts ...scw.RequestOption) (*ListCardsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "bin", req.Bin)
	parameter.AddToQuery(query, "country_code", req.CountryCode)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/it-automation_admin/v1/cards",
		Query:  query,
	}

	var resp ListCardsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

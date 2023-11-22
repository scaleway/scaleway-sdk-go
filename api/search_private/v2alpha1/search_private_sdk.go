// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package search_private provides methods and message types of the search_private v2alpha1 API.
package search_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/search_private/v2alpha1/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/search_private/v2alpha1/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/search_private/v2alpha1/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/search_private/v2alpha1/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/search_private/v2alpha1/scw"
	resource_v1alpha1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/search_private/v2alpha1/api/resource/v1alpha1"
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

// SearchRequest: search request.
type SearchRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Q string `json:"-"`

	// Type: default value: unknown_type
	Type resource_v1alpha1.ResourceType `json:"-"`
}

// SearchResponse: search response.
type SearchResponse struct {
	Resource []*resource_v1alpha1.Resource `json:"resource"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *SearchResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *SearchResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*SearchResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Resource = append(r.Resource, results.Resource...)
	r.TotalCount += uint32(len(results.Resource))
	return uint32(len(results.Resource)), nil
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

// Search: Search API.
func (s *API) Search(req *SearchRequest, opts ...scw.RequestOption) (*SearchResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "q", req.Q)
	parameter.AddToQuery(query, "type", req.Type)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/search-private/v2alpha1/search",
		Query:  query,
	}

	var resp SearchResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

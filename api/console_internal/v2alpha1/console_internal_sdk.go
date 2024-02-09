// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package console_internal provides methods and message types of the console_internal v2alpha1 API.
package console_internal

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
	std "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/std"
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

type MotdVariant string

const (
	MotdVariantUnknownVariant = MotdVariant("unknown_variant")
	MotdVariantPromotional    = MotdVariant("promotional")
	MotdVariantIntroduction   = MotdVariant("introduction")
	MotdVariantWarning        = MotdVariant("warning")
)

func (enum MotdVariant) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_variant"
	}
	return string(enum)
}

func (enum MotdVariant) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MotdVariant) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MotdVariant(MotdVariant(tmp).String())
	return nil
}

// MotdLink: motd link.
type MotdLink struct {
	Label string `json:"label"`

	Target string `json:"target"`

	URL string `json:"url"`
}

// Motd: motd.
type Motd struct {
	Message string `json:"message"`

	// Variant: default value: unknown_variant
	Variant MotdVariant `json:"variant"`

	Link *MotdLink `json:"link"`
}

// Changelog: changelog.
type Changelog struct {
	Category string `json:"category"`

	Product string `json:"product"`

	Title string `json:"title"`

	Description string `json:"description"`

	URL string `json:"url"`

	PublishedAt *time.Time `json:"published_at"`
}

// News: news.
type News struct {
	Title string `json:"title"`

	Description string `json:"description"`

	URL string `json:"url"`

	PublishedAt *time.Time `json:"published_at"`

	ImageURL string `json:"image_url"`

	Author string `json:"author"`
}

// CustomerIP: customer ip.
type CustomerIP struct {
	IP net.IP `json:"ip"`
}

// GetCustomerIPRequest: get customer ip request.
type GetCustomerIPRequest struct {
}

// GetMotdsRequest: get motds request.
type GetMotdsRequest struct {
	// Locale: default value: unknown_language_code
	Locale std.LanguageCode `json:"-"`
}

// GetMotdsResponse: get motds response.
type GetMotdsResponse struct {
	Motd map[string]*Motd `json:"motd"`
}

// ListChangelogsRequest: list changelogs request.
type ListChangelogsRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListChangelogsResponse: list changelogs response.
type ListChangelogsResponse struct {
	Changelogs []*Changelog `json:"changelogs"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListChangelogsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListChangelogsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListChangelogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Changelogs = append(r.Changelogs, results.Changelogs...)
	r.TotalCount += uint64(len(results.Changelogs))
	return uint64(len(results.Changelogs)), nil
}

// ListNewsRequest: list news request.
type ListNewsRequest struct {
}

// ListNewsResponse: list news response.
type ListNewsResponse struct {
	News []*News `json:"news"`
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

// GetCustomerIP:
func (s *API) GetCustomerIP(req *GetCustomerIPRequest, opts ...scw.RequestOption) (*CustomerIP, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/console-internal/v2alpha1/ip",
	}

	var resp CustomerIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMotds:
func (s *API) GetMotds(req *GetMotdsRequest, opts ...scw.RequestOption) (*GetMotdsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "locale", req.Locale)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/console-internal/v2alpha1/motds",
		Query:  query,
	}

	var resp GetMotdsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNews:
func (s *API) ListNews(req *ListNewsRequest, opts ...scw.RequestOption) (*ListNewsResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/console-internal/v2alpha1/news",
	}

	var resp ListNewsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListChangelogs:
func (s *API) ListChangelogs(req *ListChangelogsRequest, opts ...scw.RequestOption) (*ListChangelogsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/console-internal/v2alpha1/changelogs",
		Query:  query,
	}

	var resp ListChangelogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

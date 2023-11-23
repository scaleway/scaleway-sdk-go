// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package console_internal provides methods and message types of the console_internal v1 API.
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

type MotdType string

const (
	MotdTypeUnknown = MotdType("unknown")
	MotdTypeInfo    = MotdType("info")
	MotdTypeWarning = MotdType("warning")
	MotdTypeAlert   = MotdType("alert")
)

func (enum MotdType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum MotdType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MotdType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MotdType(MotdType(tmp).String())
	return nil
}

// MotdLink: motd link.
type MotdLink struct {
	Title string `json:"title"`

	Target string `json:"target"`

	URL string `json:"url"`
}

// Motd: motd.
type Motd struct {
	Message string `json:"message"`

	// Type: default value: unknown
	Type MotdType `json:"type"`

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

// GetMotdsConsoleRequest: get motds console request.
type GetMotdsConsoleRequest struct {
	// Locale: default value: unknown_language_code
	Locale std.LanguageCode `json:"-"`
}

// GetMotdsConsoleResponse: get motds console response.
type GetMotdsConsoleResponse struct {
	Global *Motd `json:"global"`

	Instance *Motd `json:"instance"`

	ElasticMetal *Motd `json:"elastic_metal"`

	Kapsule *Motd `json:"kapsule"`

	ObjectStorage *Motd `json:"object_storage"`

	Rdb *Motd `json:"rdb"`

	LB *Motd `json:"lb"`

	Domains *Motd `json:"domains"`

	Registry *Motd `json:"registry"`

	Iot *Motd `json:"iot"`

	AppleSilicon *Motd `json:"apple_silicon"`

	Functions *Motd `json:"functions"`

	Containers *Motd `json:"containers"`

	PrivateNetwork *Motd `json:"private_network"`

	PublicGateway *Motd `json:"public_gateway"`

	Dedibox *Motd `json:"dedibox"`

	Iam *Motd `json:"iam"`

	Messaging *Motd `json:"messaging"`

	Observability *Motd `json:"observability"`

	Redis *Motd `json:"redis"`

	Tem *Motd `json:"tem"`

	Webhosting *Motd `json:"webhosting"`

	SecretManager *Motd `json:"secret_manager"`

	Vpc *Motd `json:"vpc"`

	DocumentDb *Motd `json:"document_db"`

	ServerlessSqldb *Motd `json:"serverless_sqldb"`
}

// ListChangelogsRequest: list changelogs request.
type ListChangelogsRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListChangelogsResponse: list changelogs response.
type ListChangelogsResponse struct {
	Changelogs []*Changelog `json:"changelogs"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListChangelogsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListChangelogsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListChangelogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Changelogs = append(r.Changelogs, results.Changelogs...)
	r.TotalCount += uint32(len(results.Changelogs))
	return uint32(len(results.Changelogs)), nil
}

// ListNewsRequest: list news request.
type ListNewsRequest struct {
}

// ListNewsResponse: list news response.
type ListNewsResponse struct {
	News []*News `json:"news"`
}

// Console internal API.
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
		Path:   "/console-internal/v1/ip",
	}

	var resp CustomerIP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMotdsConsole:
func (s *API) GetMotdsConsole(req *GetMotdsConsoleRequest, opts ...scw.RequestOption) (*GetMotdsConsoleResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "locale", req.Locale)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/console-internal/v1/motds-console",
		Query:  query,
	}

	var resp GetMotdsConsoleResponse

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
		Path:   "/console-internal/v1/news",
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
		Path:   "/console-internal/v1/changelogs",
		Query:  query,
	}

	var resp ListChangelogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

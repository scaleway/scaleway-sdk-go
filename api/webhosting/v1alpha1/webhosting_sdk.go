// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package webhosting provides methods and message types of the webhosting v1alpha1 API.
package webhosting

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

// API: webhosting API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type DNSRecordStatus string

const (
	DNSRecordStatusUnknownStatus = DNSRecordStatus("unknown_status")
	DNSRecordStatusValid         = DNSRecordStatus("valid")
	DNSRecordStatusInvalid       = DNSRecordStatus("invalid")
)

func (enum DNSRecordStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum DNSRecordStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DNSRecordStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DNSRecordStatus(DNSRecordStatus(tmp).String())
	return nil
}

type DNSRecordType string

const (
	DNSRecordTypeUnknownType = DNSRecordType("unknown_type")
	DNSRecordTypeA           = DNSRecordType("a")
	DNSRecordTypeCname       = DNSRecordType("cname")
	DNSRecordTypeMx          = DNSRecordType("mx")
	DNSRecordTypeTxt         = DNSRecordType("txt")
	DNSRecordTypeNs          = DNSRecordType("ns")
	DNSRecordTypeAaaa        = DNSRecordType("aaaa")
)

func (enum DNSRecordType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum DNSRecordType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DNSRecordType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DNSRecordType(DNSRecordType(tmp).String())
	return nil
}

type DNSRecordsStatus string

const (
	DNSRecordsStatusUnknown = DNSRecordsStatus("unknown")
	DNSRecordsStatusValid   = DNSRecordsStatus("valid")
	DNSRecordsStatusInvalid = DNSRecordsStatus("invalid")
)

func (enum DNSRecordsStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DNSRecordsStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DNSRecordsStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DNSRecordsStatus(DNSRecordsStatus(tmp).String())
	return nil
}

type HostingDNSStatus string

const (
	HostingDNSStatusUnknownDNSStatus = HostingDNSStatus("unknown_dns_status")
	HostingDNSStatusValid            = HostingDNSStatus("valid")
	HostingDNSStatusInvalid          = HostingDNSStatus("invalid")
)

func (enum HostingDNSStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_dns_status"
	}
	return string(enum)
}

func (enum HostingDNSStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HostingDNSStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HostingDNSStatus(HostingDNSStatus(tmp).String())
	return nil
}

type HostingStatus string

const (
	HostingStatusUnknownStatus = HostingStatus("unknown_status")
	HostingStatusDelivering    = HostingStatus("delivering")
	HostingStatusReady         = HostingStatus("ready")
	HostingStatusDeleting      = HostingStatus("deleting")
	HostingStatusError         = HostingStatus("error")
	HostingStatusLocked        = HostingStatus("locked")
)

func (enum HostingStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum HostingStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HostingStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HostingStatus(HostingStatus(tmp).String())
	return nil
}

type ListHostingsRequestOrderBy string

const (
	ListHostingsRequestOrderByCreatedAtAsc  = ListHostingsRequestOrderBy("created_at_asc")
	ListHostingsRequestOrderByCreatedAtDesc = ListHostingsRequestOrderBy("created_at_desc")
)

func (enum ListHostingsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListHostingsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListHostingsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListHostingsRequestOrderBy(ListHostingsRequestOrderBy(tmp).String())
	return nil
}

type ListOffersRequestOrderBy string

const (
	ListOffersRequestOrderByPriceAsc = ListOffersRequestOrderBy("price_asc")
)

func (enum ListOffersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "price_asc"
	}
	return string(enum)
}

func (enum ListOffersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOffersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOffersRequestOrderBy(ListOffersRequestOrderBy(tmp).String())
	return nil
}

type NameserverStatus string

const (
	NameserverStatusUnknownStatus = NameserverStatus("unknown_status")
	NameserverStatusValid         = NameserverStatus("valid")
	NameserverStatusInvalid       = NameserverStatus("invalid")
)

func (enum NameserverStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum NameserverStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NameserverStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NameserverStatus(NameserverStatus(tmp).String())
	return nil
}

// DNSRecord: dns record.
type DNSRecord struct {
	// Name: record name.
	Name string `json:"name"`
	// Type: record type.
	// Default value: unknown_type
	Type DNSRecordType `json:"type"`
	// TTL: record time to live.
	TTL uint32 `json:"ttl"`
	// Value: record value.
	Value string `json:"value"`
	// Priority: record priority level.
	Priority *uint32 `json:"priority"`
	// Status: record status.
	// Default value: unknown_status
	Status DNSRecordStatus `json:"status"`
}

// DNSRecords: dns records.
type DNSRecords struct {
	// Records: list of DNS records.
	Records []*DNSRecord `json:"records"`
	// NameServers: list of nameservers.
	NameServers []*Nameserver `json:"name_servers"`
	// Status: status of the records.
	// Default value: unknown
	Status DNSRecordsStatus `json:"status"`
}

// Hosting: hosting.
type Hosting struct {
	// ID: ID of the hosting.
	ID string `json:"id"`
	// OrganizationID: organization ID of the hosting.
	OrganizationID string `json:"organization_id"`
	// ProjectID: project ID of the hosting.
	ProjectID string `json:"project_id"`
	// UpdatedAt: last update date.
	UpdatedAt *time.Time `json:"updated_at"`
	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`
	// Status: the hosting status.
	// Default value: unknown_status
	Status HostingStatus `json:"status"`
	// PlatformHostname: hostname of the host platform.
	PlatformHostname string `json:"platform_hostname"`
	// PlatformNumber: number of the host platform.
	PlatformNumber *int32 `json:"platform_number"`
	// OfferID: ID of the active offer.
	OfferID string `json:"offer_id"`
	// OfferName: name of the active offer.
	OfferName string `json:"offer_name"`
	// Domain: main domain of the hosting.
	Domain string `json:"domain"`
	// Tags: tags of the hosting.
	Tags []string `json:"tags"`
	// Options: active options of the hosting.
	Options []*HostingOption `json:"options"`
	// DNSStatus: DNS status of the hosting.
	// Default value: unknown_dns_status
	DNSStatus HostingDNSStatus `json:"dns_status"`
	// CpanelURLs: URL to connect to cPanel Dashboard and to Webmail interface.
	CpanelURLs *HostingCpanelURLs `json:"cpanel_urls"`
	// Username: main hosting cPanel username.
	Username string `json:"username"`
	// Region: region of the hosting.
	Region scw.Region `json:"region"`
}

type HostingCpanelURLs struct {
	Dashboard string `json:"dashboard"`

	Webmail string `json:"webmail"`
}

// HostingOption: hosting. option.
type HostingOption struct {
	// ID: option ID.
	ID string `json:"id"`
	// Name: option name.
	Name string `json:"name"`
}

// ListHostingsResponse: list hostings response.
type ListHostingsResponse struct {
	// TotalCount: number of returned hostings.
	TotalCount uint32 `json:"total_count"`
	// Hostings: list of hostings.
	Hostings []*Hosting `json:"hostings"`
}

// ListOffersResponse: list offers response.
type ListOffersResponse struct {
	// Offers: list of returned offers.
	Offers []*Offer `json:"offers"`
}

// Nameserver: nameserver.
type Nameserver struct {
	// Hostname: hostname of the nameserver.
	Hostname string `json:"hostname"`
	// Status: status of the nameserver.
	// Default value: unknown_status
	Status NameserverStatus `json:"status"`
	// IsDefault: if the nameserver is the default.
	IsDefault bool `json:"is_default"`
}

// Offer: offer.
type Offer struct {
	// ID: offer ID.
	ID string `json:"id"`
	// BillingOperationPath: unique identifier used for billing.
	BillingOperationPath string `json:"billing_operation_path"`
	// Product: offer product.
	Product *OfferProduct `json:"product"`
	// Price: offer price.
	Price *scw.Money `json:"price"`
}

// OfferProduct: offer. product.
type OfferProduct struct {
	// Name: product name.
	Name string `json:"name"`
	// Option: product option.
	Option bool `json:"option"`

	EmailAccountsQuota int32 `json:"email_accounts_quota"`

	EmailStorageQuota int32 `json:"email_storage_quota"`

	DatabasesQuota int32 `json:"databases_quota"`

	HostingStorageQuota uint32 `json:"hosting_storage_quota"`

	SupportIncluded bool `json:"support_included"`

	VCPU uint32 `json:"v_cpu"`

	RAM uint32 `json:"ram"`
}

// Service API

// Regions list localities the api is available in
func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar}
}

type CreateHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// OfferID: ID of the selected offer for the hosting.
	OfferID string `json:"offer_id"`
	// ProjectID: project ID of the hosting.
	ProjectID string `json:"project_id"`
	// Email: contact email of the client for the hosting.
	Email *string `json:"email"`
	// Tags: the tags of the hosting.
	Tags []string `json:"tags"`
	// Domain: the domain name of the hosting.
	Domain string `json:"domain"`
	// OptionIDs: iDs of the selected options for the hosting.
	OptionIDs []string `json:"option_ids"`
}

// CreateHosting: create a hosting.
func (s *API) CreateHosting(req *CreateHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListHostingsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`
	// PageSize: a positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`
	// OrderBy: define the order of the returned hostings.
	// Default value: created_at_asc
	OrderBy ListHostingsRequestOrderBy `json:"-"`
	// Tags: return hostings with these tags.
	Tags *[]string `json:"-"`
	// Statuses: return hostings with these statuses.
	Statuses []HostingStatus `json:"-"`
	// Domain: return hostings with this domain.
	Domain *string `json:"-"`
	// ProjectID: return hostings from this project ID.
	ProjectID *string `json:"-"`
	// OrganizationID: return hostings from this organization ID.
	OrganizationID *string `json:"-"`
}

// ListHostings: list all hostings.
func (s *API) ListHostings(req *ListHostingsRequest, opts ...scw.RequestOption) (*ListHostingsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "domain", req.Domain)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListHostingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// HostingID: hosting ID.
	HostingID string `json:"-"`
}

// GetHosting: get the details of a Hosting with the given ID.
func (s *API) GetHosting(req *GetHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
		Headers: http.Header{},
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// HostingID: hosting ID.
	HostingID string `json:"-"`
	// Email: new contact email for the hosting.
	Email *string `json:"email"`
	// Tags: new tags for the hosting.
	Tags *[]string `json:"tags"`
	// OptionIDs: new options IDs for the hosting.
	OptionIDs *[]string `json:"option_ids"`
	// OfferID: new offer ID for the hosting.
	OfferID *string `json:"offer_id"`
}

// UpdateHosting: update a hosting.
func (s *API) UpdateHosting(req *UpdateHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// HostingID: hosting ID.
	HostingID string `json:"-"`
}

// DeleteHosting: delete a hosting with the given ID.
func (s *API) DeleteHosting(req *DeleteHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
		Headers: http.Header{},
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RestoreHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// HostingID: hosting ID.
	HostingID string `json:"-"`
}

// RestoreHosting: restore a hosting with the given ID.
func (s *API) RestoreHosting(req *RestoreHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/restore",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDomainDNSRecordsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// Domain: domain associated to the DNS records.
	Domain string `json:"-"`
}

// GetDomainDNSRecords: get the DNS records of a specified domain.
func (s *API) GetDomainDNSRecords(req *GetDomainDNSRecordsRequest, opts ...scw.RequestOption) (*DNSRecords, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.Domain) + "/dns-records",
		Headers: http.Header{},
	}

	var resp DNSRecords

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListOffersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// OrderBy: define the order of the returned hostings.
	// Default value: price_asc
	OrderBy ListOffersRequestOrderBy `json:"-"`
	// WithoutOptions: select only offers, no options.
	WithoutOptions bool `json:"-"`
	// OnlyOptions: select only options.
	OnlyOptions bool `json:"-"`
}

// ListOffers: list all offers.
func (s *API) ListOffers(req *ListOffersRequest, opts ...scw.RequestOption) (*ListOffersResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "without_options", req.WithoutOptions)
	parameter.AddToQuery(query, "only_options", req.OnlyOptions)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/offers",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListOffersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListHostingsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListHostingsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListHostingsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Hostings = append(r.Hostings, results.Hostings...)
	r.TotalCount += uint32(len(results.Hostings))
	return uint32(len(results.Hostings)), nil
}

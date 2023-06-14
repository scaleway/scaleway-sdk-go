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

// API: web Hosting API.
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
	HostingStatusMigrating     = HostingStatus("migrating")
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

type OfferQuotaWarning string

const (
	OfferQuotaWarningUnknownQuotaWarning   = OfferQuotaWarning("unknown_quota_warning")
	OfferQuotaWarningEmailCountExceeded    = OfferQuotaWarning("email_count_exceeded")
	OfferQuotaWarningDatabaseCountExceeded = OfferQuotaWarning("database_count_exceeded")
	OfferQuotaWarningDiskUsageExceeded     = OfferQuotaWarning("disk_usage_exceeded")
)

func (enum OfferQuotaWarning) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_quota_warning"
	}
	return string(enum)
}

func (enum OfferQuotaWarning) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferQuotaWarning) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferQuotaWarning(OfferQuotaWarning(tmp).String())
	return nil
}

// DNSRecord: dns record.
type DNSRecord struct {
	// Name: record name.
	Name string `json:"name"`
	// Type: record type.
	// Default value: unknown_type
	Type DNSRecordType `json:"type"`
	// TTL: record time-to-live.
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
	// ID: ID of the Web Hosting plan.
	ID string `json:"id"`
	// OrganizationID: ID of the Scaleway Organization the Web Hosting plan belongs to.
	OrganizationID string `json:"organization_id"`
	// ProjectID: ID of the Scaleway Project the Web Hosting plan belongs to.
	ProjectID string `json:"project_id"`
	// UpdatedAt: date on which the Web Hosting plan was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// CreatedAt: date on which the Web Hosting plan was created.
	CreatedAt *time.Time `json:"created_at"`
	// Status: status of the Web Hosting plan.
	// Default value: unknown_status
	Status HostingStatus `json:"status"`
	// PlatformHostname: hostname of the host platform.
	PlatformHostname string `json:"platform_hostname"`
	// PlatformNumber: number of the host platform.
	PlatformNumber *int32 `json:"platform_number"`
	// OfferID: ID of the active offer for the Web Hosting plan.
	OfferID string `json:"offer_id"`
	// OfferName: name of the active offer for the Web Hosting plan.
	OfferName string `json:"offer_name"`
	// Domain: main domain associated with the Web Hosting plan.
	Domain string `json:"domain"`
	// Tags: list of tags associated with the Web Hosting plan.
	Tags []string `json:"tags"`
	// Options: array of any options activated for the Web Hosting plan.
	Options []*HostingOption `json:"options"`
	// DNSStatus: DNS status of the Web Hosting plan.
	// Default value: unknown_dns_status
	DNSStatus HostingDNSStatus `json:"dns_status"`
	// CpanelURLs: URL to connect to cPanel dashboard and to Webmail interface.
	CpanelURLs *HostingCpanelURLs `json:"cpanel_urls"`
	// Username: main Web Hosting cPanel username.
	Username string `json:"username"`
	// Region: region where the Web Hosting plan is hosted.
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
	// TotalCount: number of Web Hosting plans returned.
	TotalCount uint32 `json:"total_count"`
	// Hostings: list of Web Hosting plans.
	Hostings []*Hosting `json:"hostings"`
}

// ListOffersResponse: list offers response.
type ListOffersResponse struct {
	// Offers: list of offers.
	Offers []*Offer `json:"offers"`
}

// Nameserver: nameserver.
type Nameserver struct {
	// Hostname: hostname of the nameserver.
	Hostname string `json:"hostname"`
	// Status: status of the nameserver.
	// Default value: unknown_status
	Status NameserverStatus `json:"status"`
	// IsDefault: defines whether the nameserver is the default one.
	IsDefault bool `json:"is_default"`
}

// Offer: offer.
type Offer struct {
	// ID: offer ID.
	ID string `json:"id"`
	// BillingOperationPath: unique identifier used for billing.
	BillingOperationPath string `json:"billing_operation_path"`
	// Product: product constituting this offer.
	Product *OfferProduct `json:"product"`
	// Price: price of this offer.
	Price *scw.Money `json:"price"`
	// Available: if a hosting_id was specified in the call, defines whether this offer is available for that Web Hosting plan to migrate (update) to.
	Available bool `json:"available"`
	// QuotaWarnings: quota warnings, if the offer is not available for the specified hosting_id.
	QuotaWarnings []OfferQuotaWarning `json:"quota_warnings"`
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
	// OfferID: ID of the selected offer for the Web Hosting plan.
	OfferID string `json:"offer_id"`
	// ProjectID: ID of the Scaleway Project in which to create the Web Hosting plan.
	ProjectID string `json:"project_id"`
	// Email: contact email for the Web Hosting client.
	Email *string `json:"email"`
	// Tags: list of tags for the Web Hosting plan.
	Tags []string `json:"tags"`
	// Domain: domain name to link to the Web Hosting plan. You must already own this domain name, and have completed the DNS validation process beforehand.
	Domain string `json:"domain"`
	// OptionIDs: iDs of any selected additional options for the Web Hosting plan.
	OptionIDs []string `json:"option_ids"`
}

// CreateHosting: order a Web Hosting plan.
// Order a Web Hosting plan, specifying the offer type required via the `offer_id` parameter.
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
	// Page: page number to return, from the paginated results (must be a positive integer).
	Page *int32 `json:"-"`
	// PageSize: number of Web Hosting plans to return (must be a positive integer lower or equal to 100).
	PageSize *uint32 `json:"-"`
	// OrderBy: sort order for Web Hosting plans in the response.
	// Default value: created_at_asc
	OrderBy ListHostingsRequestOrderBy `json:"-"`
	// Tags: tags to filter for, only Web Hosting plans with matching tags will be returned.
	Tags *[]string `json:"-"`
	// Statuses: statuses to filter for, only Web Hosting plans with matching statuses will be returned.
	Statuses []HostingStatus `json:"-"`
	// Domain: domain to filter for, only Web Hosting plans associated with this domain will be returned.
	Domain *string `json:"-"`
	// ProjectID: project ID to filter for, only Web Hosting plans from this Project will be returned.
	ProjectID *string `json:"-"`
	// OrganizationID: organization ID to filter for, only Web Hosting plans from this Organization will be returned.
	OrganizationID *string `json:"-"`
}

// ListHostings: list all Web Hosting plans.
// List all of your existing Web Hosting plans. Various filters are available to limit the results, including filtering by domain, status, tag and Project ID.
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

// GetHosting: get a Web Hosting plan.
// Get the details of one of your existing Web Hosting plans, specified by its `hosting_id`.
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
	// Email: new contact email for the Web Hosting plan.
	Email *string `json:"email"`
	// Tags: new tags for the Web Hosting plan.
	Tags *[]string `json:"tags"`
	// OptionIDs: iDs of the new options for the Web Hosting plan.
	OptionIDs *[]string `json:"option_ids"`
	// OfferID: ID of the new offer for the Web Hosting plan.
	OfferID *string `json:"offer_id"`
}

// UpdateHosting: update a Web Hosting plan.
// Update the details of one of your existing Web Hosting plans, specified by its `hosting_id`. You can update parameters including the contact email address, tags, options and offer.
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

// DeleteHosting: delete a Web Hosting plan.
// Delete a Web Hosting plan, specified by its `hosting_id`. Note that deletion is not immediate: it will take place at the end of the calendar month, after which time your Web Hosting plan and all its data (files and emails) will be irreversibly lost.
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

// RestoreHosting: restore a Web Hosting plan.
// When you [delete a Web Hosting plan](#path-hostings-delete-a-hosting), definitive deletion does not take place until the end of the calendar month. In the time between initiating the deletion, and definitive deletion at the end of the month, you can choose to **restore** the Web Hosting plan, using this endpoint and specifying its `hosting_id`.
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
	// Domain: domain associated with the DNS records.
	Domain string `json:"-"`
}

// GetDomainDNSRecords: get DNS records.
// Get the set of DNS records of a specified domain associated with a Web Hosting plan.
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
	// OrderBy: sort order of offers in the response.
	// Default value: price_asc
	OrderBy ListOffersRequestOrderBy `json:"-"`
	// WithoutOptions: defines whether the response should consist of offers only, without options.
	WithoutOptions bool `json:"-"`
	// OnlyOptions: defines whether the response should consist of options only, without offers.
	OnlyOptions bool `json:"-"`
	// HostingID: ID of a Web Hosting plan, to check compatibility with returned offers (in case of wanting to update the plan).
	HostingID *string `json:"-"`
}

// ListOffers: list all offers.
// List the different Web Hosting offers, and their options, available to order from Scaleway.
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
	parameter.AddToQuery(query, "hosting_id", req.HostingID)

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

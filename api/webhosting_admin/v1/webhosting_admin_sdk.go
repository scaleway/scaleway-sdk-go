// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package webhosting_admin provides methods and message types of the webhosting_admin v1 API.
package webhosting_admin

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
	HostingStatusUnknownStatus  = HostingStatus("unknown_status")
	HostingStatusDelivering     = HostingStatus("delivering")
	HostingStatusReady          = HostingStatus("ready")
	HostingStatusDeleting       = HostingStatus("deleting")
	HostingStatusError          = HostingStatus("error")
	HostingStatusLocked         = HostingStatus("locked")
	HostingStatusDeleted        = HostingStatus("deleted")
	HostingStatusDomainReserved = HostingStatus("domain_reserved")
	HostingStatusMigrating      = HostingStatus("migrating")
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

type ListPlatformsRequestOrderBy string

const (
	ListPlatformsRequestOrderByCreatedAtAsc  = ListPlatformsRequestOrderBy("created_at_asc")
	ListPlatformsRequestOrderByCreatedAtDesc = ListPlatformsRequestOrderBy("created_at_desc")
)

func (enum ListPlatformsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListPlatformsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPlatformsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPlatformsRequestOrderBy(ListPlatformsRequestOrderBy(tmp).String())
	return nil
}

type PlatformStatus string

const (
	PlatformStatusUnknownStatus = PlatformStatus("unknown_status")
	PlatformStatusCreating      = PlatformStatus("creating")
	PlatformStatusActive        = PlatformStatus("active")
	PlatformStatusDeleting      = PlatformStatus("deleting")
	PlatformStatusDeleted       = PlatformStatus("deleted")
	PlatformStatusError         = PlatformStatus("error")
)

func (enum PlatformStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum PlatformStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlatformStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlatformStatus(PlatformStatus(tmp).String())
	return nil
}

// HostingOption: hosting option.
type HostingOption struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

// Blacklist: blacklist.
type Blacklist struct {
	// Domain: blacklisted domain.
	Domain *string `json:"domain"`

	// Email: blacklisted email.
	Email *string `json:"email"`

	// IP: blacklisted ip.
	IP *string `json:"ip"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Region: region of the blacklist.
	Region scw.Region `json:"region"`
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

	PlatformIP string `json:"platform_ip"`

	// OfferEndOfLife: indicates if the hosting offer has reached its end of life.
	OfferEndOfLife bool `json:"offer_end_of_life"`

	// Region: region where the Web Hosting plan is hosted.
	Region scw.Region `json:"region"`
}

// Platform: platform.
type Platform struct {
	CanProvisionHostings bool `json:"can_provision_hostings"`

	CreatedAt *time.Time `json:"created_at"`

	ErrorMessage string `json:"error_message"`

	Hostname string `json:"hostname"`

	ID string `json:"id"`

	IP string `json:"ip"`

	PlatformNumber int32 `json:"platform_number"`

	// Status: default value: unknown_status
	Status PlatformStatus `json:"status"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// BlacklistAPIBlockIncomingEmailsFromDomainRequest: blacklist api block incoming emails from domain request.
type BlacklistAPIBlockIncomingEmailsFromDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Domain: domain to block mails received from.
	Domain string `json:"-"`
}

// BlacklistAPIBlockIncomingEmailsFromIPRequest: blacklist api block incoming emails from ip request.
type BlacklistAPIBlockIncomingEmailsFromIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPAddress: ip to block mails received from.
	IPAddress string `json:"-"`
}

// BlacklistAPIBlockIncomingEmailsRequest: blacklist api block incoming emails request.
type BlacklistAPIBlockIncomingEmailsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Email: email to block mails received from.
	Email string `json:"-"`
}

// BlacklistAPIListBlacklistsRequest: blacklist api list blacklists request.
type BlacklistAPIListBlacklistsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: a positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`
}

// BlacklistAPIUnblockIncomingEmailsFromDomainRequest: blacklist api unblock incoming emails from domain request.
type BlacklistAPIUnblockIncomingEmailsFromDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Domain: domain to unblock mails received from.
	Domain string `json:"-"`
}

// BlacklistAPIUnblockIncomingEmailsFromIPRequest: blacklist api unblock incoming emails from ip request.
type BlacklistAPIUnblockIncomingEmailsFromIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPAddress: ip to unblock mails received from.
	IPAddress string `json:"-"`
}

// BlacklistAPIUnblockIncomingEmailsRequest: blacklist api unblock incoming emails request.
type BlacklistAPIUnblockIncomingEmailsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Email: email to unblock mails received from.
	Email string `json:"-"`
}

// BlockIncomingEmailsFromDomainRequest: block incoming emails from domain request.
type BlockIncomingEmailsFromDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Domain: domain to block mails received from.
	Domain string `json:"-"`
}

// BlockIncomingEmailsFromIPRequest: block incoming emails from ip request.
type BlockIncomingEmailsFromIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPAddress: ip to block mails received from.
	IPAddress string `json:"-"`
}

// CreateSudoSessionRequest: create sudo session request.
type CreateSudoSessionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	HostingID string `json:"-"`
}

// CreateSudoSessionResponse: create sudo session response.
type CreateSudoSessionResponse struct {
	URL string `json:"url"`
}

// DeleteHostingRequest: delete hosting request.
type DeleteHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	HostingID string `json:"-"`
}

// GetHostingRequest: get hosting request.
type GetHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	HostingID string `json:"-"`
}

// GetPlatformRequest: get platform request.
type GetPlatformRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	PlatformID string `json:"-"`
}

// IncomingEmailsResponse: incoming emails response.
type IncomingEmailsResponse struct {
}

// ListBlacklistsResponse: list blacklists response.
type ListBlacklistsResponse struct {
	// TotalCount: number of returned hostings.
	TotalCount uint32 `json:"total_count"`

	// Blacklists: list of returned blacklists.
	Blacklists []*Blacklist `json:"blacklists"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBlacklistsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBlacklistsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListBlacklistsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Blacklists = append(r.Blacklists, results.Blacklists...)
	r.TotalCount += uint32(len(results.Blacklists))
	return uint32(len(results.Blacklists)), nil
}

// ListHostingsRequest: list hostings request.
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

// ListHostingsResponse: list hostings response.
type ListHostingsResponse struct {
	// TotalCount: number of Web Hosting plans returned.
	TotalCount uint32 `json:"total_count"`

	// Hostings: list of Web Hosting plans.
	Hostings []*Hosting `json:"hostings"`
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

// ListPlatformsRequest: list platforms request.
type ListPlatformsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	CanProvisionHostings *bool `json:"-"`

	Hostname *string `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListPlatformsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	PlatformNumber *int32 `json:"-"`

	// Status: default value: unknown_status
	Status PlatformStatus `json:"-"`
}

// ListPlatformsResponse: list platforms response.
type ListPlatformsResponse struct {
	Platforms []*Platform `json:"platforms"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPlatformsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPlatformsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPlatformsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Platforms = append(r.Platforms, results.Platforms...)
	r.TotalCount += uint32(len(results.Platforms))
	return uint32(len(results.Platforms)), nil
}

// MoveToPremiumRequest: move to premium request.
type MoveToPremiumRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	HostingID string `json:"-"`
}

// OutgoingEmailsResponse: outgoing emails response.
type OutgoingEmailsResponse struct {
}

// RelaunchCreateHostingRequest: relaunch create hosting request.
type RelaunchCreateHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	HostingID string `json:"-"`
}

// RemoveFromPremiumRequest: remove from premium request.
type RemoveFromPremiumRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	HostingID string `json:"-"`
}

// ResetHostingPasswordRequest: reset hosting password request.
type ResetHostingPasswordRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	HostingID string `json:"-"`
}

// ResetHostingPasswordResponse: reset hosting password response.
type ResetHostingPasswordResponse struct {
	DashboardURL string `json:"dashboard_url"`

	Password string `json:"password"`

	Username string `json:"username"`
}

// SuspendOutgoingEmailsFromUserRequest: suspend outgoing emails from user request.
type SuspendOutgoingEmailsFromUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Username string `json:"-"`
}

// TrackEmailRecipientsRequest: track email recipients request.
type TrackEmailRecipientsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Sender string `json:"sender"`
}

// TrackEmailRecipientsResponse: track email recipients response.
type TrackEmailRecipientsResponse struct {
	Recipients []string `json:"recipients"`
}

// UnblockIncomingEmailsFromDomainRequest: unblock incoming emails from domain request.
type UnblockIncomingEmailsFromDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Domain: domain to unblock mails received from.
	Domain string `json:"-"`
}

// UnblockIncomingEmailsFromIPRequest: unblock incoming emails from ip request.
type UnblockIncomingEmailsFromIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// IPAddress: ip to unblock mails received from.
	IPAddress string `json:"-"`
}

// UnsuspendOutgoingEmailsFromUserRequest: unsuspend outgoing emails from user request.
type UnsuspendOutgoingEmailsFromUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Username string `json:"-"`
}

type BlacklistAPI struct {
	client *scw.Client
}

// NewBlacklistAPI returns a BlacklistAPI object from a Scaleway client.
func NewBlacklistAPI(client *scw.Client) *BlacklistAPI {
	return &BlacklistAPI{
		client: client,
	}
}
func (s *BlacklistAPI) Regions() []scw.Region {
	return []scw.Region{}
}

// ListBlacklists: Lists blacklisted incoming emails from domains, emails and ips.
func (s *BlacklistAPI) ListBlacklists(req *BlacklistAPIListBlacklistsRequest, opts ...scw.RequestOption) (*ListBlacklistsResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/blacklist/blacklists",
		Query:  query,
	}

	var resp ListBlacklistsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BlockIncomingEmails: Block incoming emails from specific email.
func (s *BlacklistAPI) BlockIncomingEmails(req *BlacklistAPIBlockIncomingEmailsRequest, opts ...scw.RequestOption) (*Blacklist, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.Email) == "" {
		return nil, errors.New("field Email cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/blacklist/emails/" + fmt.Sprint(req.Email) + "/block",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Blacklist

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnblockIncomingEmails: Unblock incoming emails from specific email.
func (s *BlacklistAPI) UnblockIncomingEmails(req *BlacklistAPIUnblockIncomingEmailsRequest, opts ...scw.RequestOption) (*Blacklist, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.Email) == "" {
		return nil, errors.New("field Email cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/blacklist/emails/" + fmt.Sprint(req.Email) + "/unblock",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Blacklist

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BlockIncomingEmailsFromDomain: Block incoming emails from specific domain.
func (s *BlacklistAPI) BlockIncomingEmailsFromDomain(req *BlacklistAPIBlockIncomingEmailsFromDomainRequest, opts ...scw.RequestOption) (*Blacklist, error) {
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
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/blacklist/domains/" + fmt.Sprint(req.Domain) + "/block",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Blacklist

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnblockIncomingEmailsFromDomain: Unblock incoming emails from specific domain.
func (s *BlacklistAPI) UnblockIncomingEmailsFromDomain(req *BlacklistAPIUnblockIncomingEmailsFromDomainRequest, opts ...scw.RequestOption) (*Blacklist, error) {
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
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/blacklist/domains/" + fmt.Sprint(req.Domain) + "/unblock",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Blacklist

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BlockIncomingEmailsFromIP: Block incoming emails from specific ip.
func (s *BlacklistAPI) BlockIncomingEmailsFromIP(req *BlacklistAPIBlockIncomingEmailsFromIPRequest, opts ...scw.RequestOption) (*Blacklist, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPAddress) == "" {
		return nil, errors.New("field IPAddress cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/blacklist/ips/" + fmt.Sprint(req.IPAddress) + "/block",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Blacklist

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnblockIncomingEmailsFromIP: Unblock incoming emails from specific ip.
func (s *BlacklistAPI) UnblockIncomingEmailsFromIP(req *BlacklistAPIUnblockIncomingEmailsFromIPRequest, opts ...scw.RequestOption) (*Blacklist, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPAddress) == "" {
		return nil, errors.New("field IPAddress cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/blacklist/ips/" + fmt.Sprint(req.IPAddress) + "/unblock",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Blacklist

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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
func (s *API) Regions() []scw.Region {
	return []scw.Region{}
}

// ListPlatforms:
func (s *API) ListPlatforms(req *ListPlatformsRequest, opts ...scw.RequestOption) (*ListPlatformsResponse, error) {
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
	parameter.AddToQuery(query, "can_provision_hostings", req.CanProvisionHostings)
	parameter.AddToQuery(query, "hostname", req.Hostname)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "platform_number", req.PlatformNumber)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/platforms",
		Query:  query,
	}

	var resp ListPlatformsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPlatform:
func (s *API) GetPlatform(req *GetPlatformRequest, opts ...scw.RequestOption) (*Platform, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PlatformID) == "" {
		return nil, errors.New("field PlatformID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/platforms/" + fmt.Sprint(req.PlatformID) + "",
	}

	var resp Platform

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHosting: Get the details of one of your existing Web Hosting plans, specified by its `hosting_id`.
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
		Method: "GET",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListHostings: List all of your existing Web Hosting plans. Various filters are available to limit the results, including filtering by domain, status, tag and Project ID.
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
		Method: "GET",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/hostings",
		Query:  query,
	}

	var resp ListHostingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteHosting: Delete a Web Hosting plan, specified by its `hosting_id`. Note that deletion is not immediate: it will take place at the end of the calendar month, after which time your Web Hosting plan and all its data (files and emails) will be irreversibly lost.
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
		Method: "DELETE",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResetHostingPassword:
func (s *API) ResetHostingPassword(req *ResetHostingPasswordRequest, opts ...scw.RequestOption) (*ResetHostingPasswordResponse, error) {
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
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/reset-password",
	}

	var resp ResetHostingPasswordResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSudoSession:
func (s *API) CreateSudoSession(req *CreateSudoSessionRequest, opts ...scw.RequestOption) (*CreateSudoSessionResponse, error) {
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
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/create-sudo-session",
	}

	var resp CreateSudoSessionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BlockIncomingEmailsFromDomain:
func (s *API) BlockIncomingEmailsFromDomain(req *BlockIncomingEmailsFromDomainRequest, opts ...scw.RequestOption) (*IncomingEmailsResponse, error) {
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
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/emails/domains/" + fmt.Sprint(req.Domain) + "/block",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IncomingEmailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnblockIncomingEmailsFromDomain:
func (s *API) UnblockIncomingEmailsFromDomain(req *UnblockIncomingEmailsFromDomainRequest, opts ...scw.RequestOption) (*IncomingEmailsResponse, error) {
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
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/emails/domains/" + fmt.Sprint(req.Domain) + "/unblock",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IncomingEmailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SuspendOutgoingEmailsFromUser:
func (s *API) SuspendOutgoingEmailsFromUser(req *SuspendOutgoingEmailsFromUserRequest, opts ...scw.RequestOption) (*OutgoingEmailsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.Username) == "" {
		return nil, errors.New("field Username cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/emails/users/" + fmt.Sprint(req.Username) + "/suspend",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp OutgoingEmailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsuspendOutgoingEmailsFromUser:
func (s *API) UnsuspendOutgoingEmailsFromUser(req *UnsuspendOutgoingEmailsFromUserRequest, opts ...scw.RequestOption) (*OutgoingEmailsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.Username) == "" {
		return nil, errors.New("field Username cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/emails/users/" + fmt.Sprint(req.Username) + "/unsuspend",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp OutgoingEmailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BlockIncomingEmailsFromIP:
func (s *API) BlockIncomingEmailsFromIP(req *BlockIncomingEmailsFromIPRequest, opts ...scw.RequestOption) (*IncomingEmailsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPAddress) == "" {
		return nil, errors.New("field IPAddress cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/emails/ips/" + fmt.Sprint(req.IPAddress) + "/block",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IncomingEmailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnblockIncomingEmailsFromIP:
func (s *API) UnblockIncomingEmailsFromIP(req *UnblockIncomingEmailsFromIPRequest, opts ...scw.RequestOption) (*IncomingEmailsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPAddress) == "" {
		return nil, errors.New("field IPAddress cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/emails/ips/" + fmt.Sprint(req.IPAddress) + "/unblock",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IncomingEmailsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TrackEmailRecipients:
func (s *API) TrackEmailRecipients(req *TrackEmailRecipientsRequest, opts ...scw.RequestOption) (*TrackEmailRecipientsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/emails/track",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TrackEmailRecipientsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RelaunchCreateHosting:
func (s *API) RelaunchCreateHosting(req *RelaunchCreateHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
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
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/relaunch-create",
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

// MoveToPremium:
func (s *API) MoveToPremium(req *MoveToPremiumRequest, opts ...scw.RequestOption) (*Hosting, error) {
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
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/move-to-premium",
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

// RemoveFromPremium:
func (s *API) RemoveFromPremium(req *RemoveFromPremiumRequest, opts ...scw.RequestOption) (*Hosting, error) {
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
		Method: "POST",
		Path:   "/webhosting-admin/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/remove-from-premium",
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

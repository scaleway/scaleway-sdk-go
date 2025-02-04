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

	std "github.com/scaleway/scaleway-sdk-go/api/std"
	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
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

func (enum DNSRecordStatus) Values() []DNSRecordStatus {
	return []DNSRecordStatus{
		"unknown_status",
		"valid",
		"invalid",
	}
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

func (enum DNSRecordType) Values() []DNSRecordType {
	return []DNSRecordType{
		"unknown_type",
		"a",
		"cname",
		"mx",
		"txt",
		"ns",
		"aaaa",
	}
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

func (enum DNSRecordsStatus) Values() []DNSRecordsStatus {
	return []DNSRecordsStatus{
		"unknown",
		"valid",
		"invalid",
	}
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

func (enum HostingDNSStatus) Values() []HostingDNSStatus {
	return []HostingDNSStatus{
		"unknown_dns_status",
		"valid",
		"invalid",
	}
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

func (enum HostingStatus) Values() []HostingStatus {
	return []HostingStatus{
		"unknown_status",
		"delivering",
		"ready",
		"deleting",
		"error",
		"locked",
		"migrating",
	}
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

func (enum ListHostingsRequestOrderBy) Values() []ListHostingsRequestOrderBy {
	return []ListHostingsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
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

func (enum ListOffersRequestOrderBy) Values() []ListOffersRequestOrderBy {
	return []ListOffersRequestOrderBy{
		"price_asc",
	}
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

func (enum NameserverStatus) Values() []NameserverStatus {
	return []NameserverStatus{
		"unknown_status",
		"valid",
		"invalid",
	}
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
	OfferQuotaWarningUnknownQuotaWarning      = OfferQuotaWarning("unknown_quota_warning")
	OfferQuotaWarningEmailCountExceeded       = OfferQuotaWarning("email_count_exceeded")
	OfferQuotaWarningDatabaseCountExceeded    = OfferQuotaWarning("database_count_exceeded")
	OfferQuotaWarningDiskUsageExceeded        = OfferQuotaWarning("disk_usage_exceeded")
	OfferQuotaWarningAddonDomainCountExceeded = OfferQuotaWarning("addon_domain_count_exceeded")
)

func (enum OfferQuotaWarning) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_quota_warning"
	}
	return string(enum)
}

func (enum OfferQuotaWarning) Values() []OfferQuotaWarning {
	return []OfferQuotaWarning{
		"unknown_quota_warning",
		"email_count_exceeded",
		"database_count_exceeded",
		"disk_usage_exceeded",
		"addon_domain_count_exceeded",
	}
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

// HostingCpanelURLs: hosting cpanel ur ls.
type HostingCpanelURLs struct {
	Dashboard string `json:"dashboard"`

	Webmail string `json:"webmail"`
}

// HostingOption: hosting option.
type HostingOption struct {
	// ID: option ID.
	ID string `json:"id"`

	// Name: option name.
	Name string `json:"name"`
}

// EmailAddress: email address.
type EmailAddress struct {
	// Domain: domain part of the mailbox address.
	Domain string `json:"domain"`

	// Login: username part address of the mailbox address.
	Login string `json:"login"`
}

// OfferProduct: offer product.
type OfferProduct struct {
	// Name: product name.
	Name string `json:"name"`

	// Option: product option.
	Option bool `json:"option"`

	// EmailAccountsQuota: limit number of email accounts.
	EmailAccountsQuota int32 `json:"email_accounts_quota"`

	// EmailStorageQuota: limit quantity of email storage in gigabytes.
	EmailStorageQuota int32 `json:"email_storage_quota"`

	// DatabasesQuota: limit number of databases.
	DatabasesQuota int32 `json:"databases_quota"`

	// HostingStorageQuota: limit quantity of hosting storage in gigabytes.
	HostingStorageQuota uint32 `json:"hosting_storage_quota"`

	// SupportIncluded: whether or not support is included.
	SupportIncluded bool `json:"support_included"`

	// VCPU: limit number of virtual CPU.
	VCPU uint32 `json:"v_cpu"`

	// RAM: limit quantity of memory in gigabytes.
	RAM uint32 `json:"ram"`

	// MaxAddonDomains: limit number of add-on domains.
	MaxAddonDomains int32 `json:"max_addon_domains"`
}

// CreateHostingRequestDomainConfiguration: create hosting request domain configuration.
type CreateHostingRequestDomainConfiguration struct {
	UpdateNameservers bool `json:"update_nameservers"`

	UpdateWebRecord bool `json:"update_web_record"`

	UpdateMailRecord bool `json:"update_mail_record"`

	UpdateAllRecords bool `json:"update_all_records"`
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

// ControlPanel: control panel.
type ControlPanel struct {
	// Name: control panel name.
	Name string `json:"name"`

	// Available: define if the control panel type is available to order.
	Available bool `json:"available"`

	// LogoURL: URL of this control panel's logo.
	LogoURL string `json:"logo_url"`

	// AvailableLanguages: list of available languages for the control panel.
	AvailableLanguages []std.LanguageCode `json:"available_languages"`
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

	// OfferEndOfLife: indicates if the hosting offer has reached its end of life.
	OfferEndOfLife bool `json:"offer_end_of_life"`

	// ControlPanelName: name of the control panel.
	ControlPanelName string `json:"control_panel_name"`

	// PlatformGroup: group of the hosting's host server/platform.
	PlatformGroup string `json:"platform_group"`

	// IPv4: iPv4 address of the hosting's host server.
	IPv4 string `json:"ipv4"`

	// IPv6: iPv6 address of the hosting's host server.
	IPv6 string `json:"ipv6"`

	// Protected: whether the hosting is protected or not.
	Protected bool `json:"protected"`

	// OneTimePassword: one-time-password used for the first login or reset password, empty after first use.
	OneTimePassword string `json:"one_time_password"`

	// Region: region where the Web Hosting plan is hosted.
	Region scw.Region `json:"region"`
}

// Mailbox: mailbox.
type Mailbox struct {
	// MailboxID: the ID of the mailbox.
	MailboxID uint32 `json:"mailbox_id"`

	// Email: the email address of the mailbox.
	Email *EmailAddress `json:"email"`
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

	// EndOfLife: indicates if the offer has reached its end of life.
	EndOfLife bool `json:"end_of_life"`

	// ControlPanelName: name of the control panel.
	ControlPanelName string `json:"control_panel_name"`
}

// CheckUserOwnsDomainRequest: check user owns domain request.
type CheckUserOwnsDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Domain: domain for which ownership is to be verified.
	Domain string `json:"-"`

	// ProjectID: ID of the project currently in use.
	ProjectID string `json:"project_id"`
}

// CheckUserOwnsDomainResponse: check user owns domain response.
type CheckUserOwnsDomainResponse struct {
	// OwnsDomain: indicates whether the specified project owns the domain.
	OwnsDomain bool `json:"owns_domain"`
}

// ClassicMailAPICreateMailboxRequest: classic mail api create mailbox request.
type ClassicMailAPICreateMailboxRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OnlineID: the Online hosting ID.
	OnlineID uint32 `json:"-"`

	// Email: the email address of the mailbox.
	Email *EmailAddress `json:"email,omitempty"`

	// Password: password for the new mailbox.
	Password string `json:"password"`
}

// ClassicMailAPIDeleteMailboxRequest: classic mail api delete mailbox request.
type ClassicMailAPIDeleteMailboxRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OnlineID: the Online hosting ID.
	OnlineID uint32 `json:"-"`

	// MailboxID: the ID of the mailbox to delete.
	MailboxID uint32 `json:"-"`
}

// ClassicMailAPIGetMailboxRequest: classic mail api get mailbox request.
type ClassicMailAPIGetMailboxRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OnlineID: the Online hosting ID.
	OnlineID uint32 `json:"-"`

	// MailboxID: the ID of the mailbox to get.
	MailboxID uint32 `json:"-"`
}

// ClassicMailAPIListMailboxesRequest: classic mail api list mailboxes request.
type ClassicMailAPIListMailboxesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OnlineID: the Online hosting ID.
	OnlineID uint32 `json:"-"`

	// Page: page number (must be a positive integer).
	Page *int32 `json:"-"`

	// PageSize: number of mailboxes to return (must be a positive integer lower or equal to 100).
	PageSize *uint32 `json:"-"`

	// Domain: domain to filter the mailboxes.
	Domain *string `json:"-"`
}

// ClassicMailAPIUpdateMailboxRequest: classic mail api update mailbox request.
type ClassicMailAPIUpdateMailboxRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OnlineID: the Online hosting ID.
	OnlineID uint32 `json:"-"`

	// MailboxID: the ID of the mailbox to update.
	MailboxID uint32 `json:"-"`

	// Password: new password for the mailbox.
	Password *string `json:"password,omitempty"`
}

// CreateHostingRequest: create hosting request.
type CreateHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OfferID: ID of the selected offer for the Web Hosting plan.
	OfferID string `json:"offer_id"`

	// ProjectID: ID of the Scaleway Project in which to create the Web Hosting plan.
	ProjectID string `json:"project_id"`

	// Email: contact email for the Web Hosting client.
	Email *string `json:"email,omitempty"`

	// Tags: list of tags for the Web Hosting plan.
	Tags []string `json:"tags"`

	// Domain: domain name to link to the Web Hosting plan. You must already own this domain name, and have completed the DNS validation process beforehand.
	Domain string `json:"domain"`

	// OptionIDs: iDs of any selected additional options for the Web Hosting plan.
	OptionIDs []string `json:"option_ids"`

	// Language: default language for the control panel interface.
	// Default value: unknown_language_code
	Language std.LanguageCode `json:"language"`

	// DomainConfiguration: indicates whether to update hosting domain name servers and DNS records for domains managed by Scaleway Elements.
	DomainConfiguration *CreateHostingRequestDomainConfiguration `json:"domain_configuration,omitempty"`
}

// CreateSessionRequest: create session request.
type CreateSessionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: hosting ID.
	HostingID string `json:"-"`
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

// DeleteHostingRequest: delete hosting request.
type DeleteHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: hosting ID.
	HostingID string `json:"-"`
}

// GetDomainDNSRecordsRequest: get domain dns records request.
type GetDomainDNSRecordsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Domain: domain associated with the DNS records.
	Domain string `json:"-"`
}

// GetHostingRequest: get hosting request.
type GetHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: hosting ID.
	HostingID string `json:"-"`
}

// ListControlPanelsRequest: list control panels request.
type ListControlPanelsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return, from the paginated results (must be a positive integer).
	Page *int32 `json:"-"`

	// PageSize: number of control panels to return (must be a positive integer lower or equal to 100).
	PageSize *uint32 `json:"-"`
}

// ListControlPanelsResponse: list control panels response.
type ListControlPanelsResponse struct {
	// TotalCount: number of control panels returned.
	TotalCount uint64 `json:"total_count"`

	// ControlPanels: list of control panels.
	ControlPanels []*ControlPanel `json:"control_panels"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListControlPanelsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListControlPanelsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListControlPanelsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ControlPanels = append(r.ControlPanels, results.ControlPanels...)
	r.TotalCount += uint64(len(results.ControlPanels))
	return uint64(len(results.ControlPanels)), nil
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
	Tags []string `json:"-"`

	// Statuses: statuses to filter for, only Web Hosting plans with matching statuses will be returned.
	Statuses []HostingStatus `json:"-"`

	// Domain: domain to filter for, only Web Hosting plans associated with this domain will be returned.
	Domain *string `json:"-"`

	// ProjectID: project ID to filter for, only Web Hosting plans from this Project will be returned.
	ProjectID *string `json:"-"`

	// OrganizationID: organization ID to filter for, only Web Hosting plans from this Organization will be returned.
	OrganizationID *string `json:"-"`

	// ControlPanels: name of the control panel to filter for, only Web Hosting plans from this control panel will be returned.
	ControlPanels []string `json:"-"`
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

// ListMailboxesResponse: list mailboxes response.
type ListMailboxesResponse struct {
	// TotalCount: total number of mailboxes.
	TotalCount uint64 `json:"total_count"`

	// Mailboxes: list of mailboxes.
	Mailboxes []*Mailbox `json:"mailboxes"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListMailboxesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListMailboxesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListMailboxesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Mailboxes = append(r.Mailboxes, results.Mailboxes...)
	r.TotalCount += uint64(len(results.Mailboxes))
	return uint64(len(results.Mailboxes)), nil
}

// ListOffersRequest: list offers request.
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

	// ControlPanels: name of the control panel to filter for.
	ControlPanels []string `json:"-"`
}

// ListOffersResponse: list offers response.
type ListOffersResponse struct {
	// Offers: list of offers.
	Offers []*Offer `json:"offers"`
}

// ResetHostingPasswordRequest: reset hosting password request.
type ResetHostingPasswordRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting.
	HostingID string `json:"-"`
}

// ResetHostingPasswordResponse: reset hosting password response.
type ResetHostingPasswordResponse struct {
	// Password: new password.
	Password string `json:"password"`
}

// RestoreHostingRequest: restore hosting request.
type RestoreHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: hosting ID.
	HostingID string `json:"-"`
}

// Session: session.
type Session struct {
	// URL: logged user's session URL.
	URL string `json:"url"`
}

// UpdateHostingRequest: update hosting request.
type UpdateHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: hosting ID.
	HostingID string `json:"-"`

	// Email: new contact email for the Web Hosting plan.
	Email *string `json:"email,omitempty"`

	// Tags: new tags for the Web Hosting plan.
	Tags *[]string `json:"tags,omitempty"`

	// OptionIDs: iDs of the new options for the Web Hosting plan.
	OptionIDs *[]string `json:"option_ids,omitempty"`

	// OfferID: ID of the new offer for the Web Hosting plan.
	OfferID *string `json:"offer_id,omitempty"`

	// Protected: whether the hosting is protected or not.
	Protected *bool `json:"protected,omitempty"`
}

// This API allows you to manage your Web Hosting services.
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
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms}
}

// CreateHosting: Order a Web Hosting plan, specifying the offer type required via the `offer_id` parameter.
func (s *API) CreateHosting(req *CreateHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings",
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
	parameter.AddToQuery(query, "control_panels", req.ControlPanels)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings",
		Query:  query,
	}

	var resp ListHostingsResponse

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
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHosting: Update the details of one of your existing Web Hosting plans, specified by its `hosting_id`. You can update parameters including the contact email address, tags, options and offer.
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
		Method: "PATCH",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
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
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestoreHosting: When you [delete a Web Hosting plan](#path-hostings-delete-a-hosting), definitive deletion does not take place until the end of the calendar month. In the time between initiating the deletion, and definitive deletion at the end of the month, you can choose to **restore** the Web Hosting plan, using this endpoint and specifying its `hosting_id`.
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
		Method: "POST",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/restore",
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

// GetDomainDNSRecords: Get the set of DNS records of a specified domain associated with a Web Hosting plan.
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
		Method: "GET",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.Domain) + "/dns-records",
	}

	var resp DNSRecords

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckUserOwnsDomain: "Check whether you own this domain or not.".
func (s *API) CheckUserOwnsDomain(req *CheckUserOwnsDomainRequest, opts ...scw.RequestOption) (*CheckUserOwnsDomainResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.Domain) + "/check-ownership",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CheckUserOwnsDomainResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOffers: List the different Web Hosting offers, and their options, available to order from Scaleway.
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
	parameter.AddToQuery(query, "control_panels", req.ControlPanels)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/offers",
		Query:  query,
	}

	var resp ListOffersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListControlPanels: "List the control panels type: cpanel or plesk.".
func (s *API) ListControlPanels(req *ListControlPanelsRequest, opts ...scw.RequestOption) (*ListControlPanelsResponse, error) {
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
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/control-panels",
		Query:  query,
	}

	var resp ListControlPanelsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSession: Create a user session.
func (s *API) CreateSession(req *CreateSessionRequest, opts ...scw.RequestOption) (*Session, error) {
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
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/sessions",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Session

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
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/reset-password",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ResetHostingPasswordResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage your mailboxes for your Web Hosting services.
type ClassicMailAPI struct {
	client *scw.Client
}

// NewClassicMailAPI returns a ClassicMailAPI object from a Scaleway client.
func NewClassicMailAPI(client *scw.Client) *ClassicMailAPI {
	return &ClassicMailAPI{
		client: client,
	}
}
func (s *ClassicMailAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// CreateMailbox: Create a new mailbox within your hosting plan.
func (s *ClassicMailAPI) CreateMailbox(req *ClassicMailAPICreateMailboxRequest, opts ...scw.RequestOption) (*Mailbox, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.OnlineID) == "" {
		return nil, errors.New("field OnlineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/classic-hostings/" + fmt.Sprint(req.OnlineID) + "/mailboxes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Mailbox

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMailbox: Get a mailbox by id within your hosting plan.
func (s *ClassicMailAPI) GetMailbox(req *ClassicMailAPIGetMailboxRequest, opts ...scw.RequestOption) (*Mailbox, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.OnlineID) == "" {
		return nil, errors.New("field OnlineID cannot be empty in request")
	}

	if fmt.Sprint(req.MailboxID) == "" {
		return nil, errors.New("field MailboxID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/classic-hostings/" + fmt.Sprint(req.OnlineID) + "/mailboxes/" + fmt.Sprint(req.MailboxID) + "",
	}

	var resp Mailbox

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListMailboxes: List all mailboxes within your hosting plan.
func (s *ClassicMailAPI) ListMailboxes(req *ClassicMailAPIListMailboxesRequest, opts ...scw.RequestOption) (*ListMailboxesResponse, error) {
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
	parameter.AddToQuery(query, "domain", req.Domain)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.OnlineID) == "" {
		return nil, errors.New("field OnlineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/classic-hostings/" + fmt.Sprint(req.OnlineID) + "/mailboxes",
		Query:  query,
	}

	var resp ListMailboxesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteMailbox:
func (s *ClassicMailAPI) DeleteMailbox(req *ClassicMailAPIDeleteMailboxRequest, opts ...scw.RequestOption) (*Mailbox, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.OnlineID) == "" {
		return nil, errors.New("field OnlineID cannot be empty in request")
	}

	if fmt.Sprint(req.MailboxID) == "" {
		return nil, errors.New("field MailboxID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/classic-hostings/" + fmt.Sprint(req.OnlineID) + "/mailboxes/" + fmt.Sprint(req.MailboxID) + "",
	}

	var resp Mailbox

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateMailbox: Update the mailbox within your hosting plan.
func (s *ClassicMailAPI) UpdateMailbox(req *ClassicMailAPIUpdateMailboxRequest, opts ...scw.RequestOption) (*Mailbox, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.OnlineID) == "" {
		return nil, errors.New("field OnlineID cannot be empty in request")
	}

	if fmt.Sprint(req.MailboxID) == "" {
		return nil, errors.New("field MailboxID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/webhosting/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/classic-hostings/" + fmt.Sprint(req.OnlineID) + "/mailboxes/" + fmt.Sprint(req.MailboxID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Mailbox

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

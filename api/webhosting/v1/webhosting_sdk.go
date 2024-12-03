// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package webhosting provides methods and message types of the webhosting v1 API.
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

type HostingSummaryStatus string

const (
	HostingSummaryStatusUnknownStatus = HostingSummaryStatus("unknown_status")
	HostingSummaryStatusDelivering    = HostingSummaryStatus("delivering")
	HostingSummaryStatusReady         = HostingSummaryStatus("ready")
	HostingSummaryStatusDeleting      = HostingSummaryStatus("deleting")
	HostingSummaryStatusError         = HostingSummaryStatus("error")
	HostingSummaryStatusLocked        = HostingSummaryStatus("locked")
	HostingSummaryStatusMigrating     = HostingSummaryStatus("migrating")
)

func (enum HostingSummaryStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum HostingSummaryStatus) Values() []HostingSummaryStatus {
	return []HostingSummaryStatus{
		"unknown_status",
		"delivering",
		"ready",
		"deleting",
		"error",
		"locked",
		"migrating",
	}
}

func (enum HostingSummaryStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HostingSummaryStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HostingSummaryStatus(HostingSummaryStatus(tmp).String())
	return nil
}

type ListDatabaseUsersRequestOrderBy string

const (
	ListDatabaseUsersRequestOrderByUsernameAsc  = ListDatabaseUsersRequestOrderBy("username_asc")
	ListDatabaseUsersRequestOrderByUsernameDesc = ListDatabaseUsersRequestOrderBy("username_desc")
)

func (enum ListDatabaseUsersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "username_asc"
	}
	return string(enum)
}

func (enum ListDatabaseUsersRequestOrderBy) Values() []ListDatabaseUsersRequestOrderBy {
	return []ListDatabaseUsersRequestOrderBy{
		"username_asc",
		"username_desc",
	}
}

func (enum ListDatabaseUsersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDatabaseUsersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDatabaseUsersRequestOrderBy(ListDatabaseUsersRequestOrderBy(tmp).String())
	return nil
}

type ListDatabasesRequestOrderBy string

const (
	ListDatabasesRequestOrderByDatabaseNameAsc  = ListDatabasesRequestOrderBy("database_name_asc")
	ListDatabasesRequestOrderByDatabaseNameDesc = ListDatabasesRequestOrderBy("database_name_desc")
)

func (enum ListDatabasesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "database_name_asc"
	}
	return string(enum)
}

func (enum ListDatabasesRequestOrderBy) Values() []ListDatabasesRequestOrderBy {
	return []ListDatabasesRequestOrderBy{
		"database_name_asc",
		"database_name_desc",
	}
}

func (enum ListDatabasesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDatabasesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDatabasesRequestOrderBy(ListDatabasesRequestOrderBy(tmp).String())
	return nil
}

type ListFtpAccountsRequestOrderBy string

const (
	ListFtpAccountsRequestOrderByUsernameAsc  = ListFtpAccountsRequestOrderBy("username_asc")
	ListFtpAccountsRequestOrderByUsernameDesc = ListFtpAccountsRequestOrderBy("username_desc")
)

func (enum ListFtpAccountsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "username_asc"
	}
	return string(enum)
}

func (enum ListFtpAccountsRequestOrderBy) Values() []ListFtpAccountsRequestOrderBy {
	return []ListFtpAccountsRequestOrderBy{
		"username_asc",
		"username_desc",
	}
}

func (enum ListFtpAccountsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFtpAccountsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFtpAccountsRequestOrderBy(ListFtpAccountsRequestOrderBy(tmp).String())
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

type ListMailAccountsRequestOrderBy string

const (
	ListMailAccountsRequestOrderByUsernameAsc  = ListMailAccountsRequestOrderBy("username_asc")
	ListMailAccountsRequestOrderByUsernameDesc = ListMailAccountsRequestOrderBy("username_desc")
	ListMailAccountsRequestOrderByDomainAsc    = ListMailAccountsRequestOrderBy("domain_asc")
	ListMailAccountsRequestOrderByDomainDesc   = ListMailAccountsRequestOrderBy("domain_desc")
)

func (enum ListMailAccountsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "username_asc"
	}
	return string(enum)
}

func (enum ListMailAccountsRequestOrderBy) Values() []ListMailAccountsRequestOrderBy {
	return []ListMailAccountsRequestOrderBy{
		"username_asc",
		"username_desc",
		"domain_asc",
		"domain_desc",
	}
}

func (enum ListMailAccountsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListMailAccountsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListMailAccountsRequestOrderBy(ListMailAccountsRequestOrderBy(tmp).String())
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

type ListWebsitesRequestOrderBy string

const (
	ListWebsitesRequestOrderByDomainAsc  = ListWebsitesRequestOrderBy("domain_asc")
	ListWebsitesRequestOrderByDomainDesc = ListWebsitesRequestOrderBy("domain_desc")
)

func (enum ListWebsitesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "domain_asc"
	}
	return string(enum)
}

func (enum ListWebsitesRequestOrderBy) Values() []ListWebsitesRequestOrderBy {
	return []ListWebsitesRequestOrderBy{
		"domain_asc",
		"domain_desc",
	}
}

func (enum ListWebsitesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListWebsitesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListWebsitesRequestOrderBy(ListWebsitesRequestOrderBy(tmp).String())
	return nil
}

type OfferOptionName string

const (
	OfferOptionNameUnknownName = OfferOptionName("unknown_name")
	OfferOptionNameDomainCount = OfferOptionName("domain_count")
	OfferOptionNameEmailCount  = OfferOptionName("email_count")
	OfferOptionNameStorageGb   = OfferOptionName("storage_gb")
	OfferOptionNameVcpuCount   = OfferOptionName("vcpu_count")
	OfferOptionNameRAMGb       = OfferOptionName("ram_gb")
	OfferOptionNameBackup      = OfferOptionName("backup")
	OfferOptionNameDedicatedIP = OfferOptionName("dedicated_ip")
)

func (enum OfferOptionName) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_name"
	}
	return string(enum)
}

func (enum OfferOptionName) Values() []OfferOptionName {
	return []OfferOptionName{
		"unknown_name",
		"domain_count",
		"email_count",
		"storage_gb",
		"vcpu_count",
		"ram_gb",
		"backup",
		"dedicated_ip",
	}
}

func (enum OfferOptionName) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferOptionName) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferOptionName(OfferOptionName(tmp).String())
	return nil
}

type OfferOptionWarning string

const (
	OfferOptionWarningUnknownWarning       = OfferOptionWarning("unknown_warning")
	OfferOptionWarningQuotaExceededWarning = OfferOptionWarning("quota_exceeded_warning")
)

func (enum OfferOptionWarning) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_warning"
	}
	return string(enum)
}

func (enum OfferOptionWarning) Values() []OfferOptionWarning {
	return []OfferOptionWarning{
		"unknown_warning",
		"quota_exceeded_warning",
	}
}

func (enum OfferOptionWarning) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferOptionWarning) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferOptionWarning(OfferOptionWarning(tmp).String())
	return nil
}

// OfferOption: offer option.
type OfferOption struct {
	// ID: option ID.
	ID string `json:"id"`

	// Name: name of the option.
	// Default value: unknown_name
	Name OfferOptionName `json:"name"`

	// BillingOperationPath: unique identifier used for billing.
	BillingOperationPath string `json:"billing_operation_path"`

	// MinValue: minimum value for the option in the offer.
	MinValue int32 `json:"min_value"`

	// CurrentValue: if a hosting_id was specified in the call, defines the current value of the option in the hosting.
	CurrentValue int32 `json:"current_value"`

	// MaxValue: maximum value for the option in the offer.
	MaxValue int32 `json:"max_value"`

	// QuotaWarning: defines a warning if the maximum value for the option has been reached.
	// Default value: unknown_warning
	QuotaWarning OfferOptionWarning `json:"quota_warning"`
}

// CreateHostingRequestDomainConfiguration: create hosting request domain configuration.
type CreateHostingRequestDomainConfiguration struct {
	UpdateNameservers bool `json:"update_nameservers"`

	UpdateWebRecord bool `json:"update_web_record"`

	UpdateMailRecord bool `json:"update_mail_record"`

	UpdateAllRecords bool `json:"update_all_records"`
}

// OfferOptionRequest: offer option request.
type OfferOptionRequest struct {
	// ID: offer option ID.
	ID string `json:"id"`

	// Quantity: the option requested quantity to set for the Web Hosting plan.
	Quantity int64 `json:"quantity"`
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
	// Default value: unknown_name
	Name OfferOptionName `json:"name"`

	// Quantity: option quantity.
	Quantity int32 `json:"quantity"`
}

// ControlPanel: control panel.
type ControlPanel struct {
	// Name: control panel name.
	Name string `json:"name"`

	// Available: define if the control panel type is available to order.
	Available bool `json:"available"`

	// LogoURL: URL of the control panel's logo.
	LogoURL string `json:"logo_url"`

	// AvailableLanguages: list of available languages for the control panel.
	AvailableLanguages []std.LanguageCode `json:"available_languages"`
}

// DatabaseUser: database user.
type DatabaseUser struct {
	// Username: name of the database user.
	Username string `json:"username"`

	// Databases: list of databases accessible by the user.
	Databases []string `json:"databases"`
}

// Database: database.
type Database struct {
	// DatabaseName: name of the database.
	DatabaseName string `json:"database_name"`

	// Users: list of users who have access to the database.
	Users []string `json:"users"`
}

// FtpAccount: ftp account.
type FtpAccount struct {
	// Username: the username of the FTP account.
	Username string `json:"username"`

	// Path: the path associated with the FTP account.
	Path string `json:"path"`
}

// HostingSummary: hosting summary.
type HostingSummary struct {
	// ID: ID of the Web Hosting plan.
	ID string `json:"id"`

	// ProjectID: ID of the Scaleway Project the Web Hosting plan belongs to.
	ProjectID string `json:"project_id"`

	// CreatedAt: date on which the Web Hosting plan was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date on which the Web Hosting plan was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// Status: status of the Web Hosting plan.
	// Default value: unknown_status
	Status HostingSummaryStatus `json:"status"`

	// Domain: main domain associated with the Web Hosting plan.
	Domain string `json:"domain"`

	// Protected: whether the hosting is protected or not.
	Protected bool `json:"protected"`

	// Region: region where the Web Hosting plan is hosted.
	Region scw.Region `json:"region"`
}

// MailAccount: mail account.
type MailAccount struct {
	// Domain: domain part of the mail account address.
	Domain string `json:"domain"`

	// Username: username part address of the mail account address.
	Username string `json:"username"`
}

// Offer: offer.
type Offer struct {
	// ID: offer ID.
	ID string `json:"id"`

	// BillingOperationPath: unique identifier used for billing.
	BillingOperationPath string `json:"billing_operation_path"`

	// Options: options available for the offer.
	Options []*OfferOption `json:"options"`

	// Price: price of the offer.
	Price *scw.Money `json:"price"`

	// Available: if a hosting_id was specified in the call, defines whether the offer is available for a specified hosting plan to migrate (update) to.
	Available bool `json:"available"`

	// ControlPanelName: name of the control panel.
	ControlPanelName string `json:"control_panel_name"`

	// EndOfLife: indicates if the offer has reached its end of life.
	EndOfLife bool `json:"end_of_life"`
}

// Website: website.
type Website struct {
	// Domain: the domain of the website.
	Domain string `json:"domain"`

	// Path: the directory path of the website.
	Path string `json:"path"`

	// SslStatus: the SSL status of the website.
	SslStatus bool `json:"ssl_status"`
}

// ControlPanelAPIListControlPanelsRequest: control panel api list control panels request.
type ControlPanelAPIListControlPanelsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number (must be a positive integer).
	Page *int32 `json:"-"`

	// PageSize: number of control panels to return (must be a positive integer lower or equal to 100).
	PageSize *uint32 `json:"-"`
}

// DatabaseAPIAssignDatabaseUserRequest: database api assign database user request.
type DatabaseAPIAssignDatabaseUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// DatabaseName: name of the database to be assigned.
	DatabaseName string `json:"-"`

	// Username: name of the user to assign.
	Username string `json:"username"`
}

// DatabaseAPIChangeDatabaseUserPasswordRequest: database api change database user password request.
type DatabaseAPIChangeDatabaseUserPasswordRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Username: name of the user to update.
	Username string `json:"-"`

	// Password: new password.
	Password string `json:"password"`
}

// DatabaseAPICreateDatabaseRequest: database api create database request.
type DatabaseAPICreateDatabaseRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan where the database will be created.
	HostingID string `json:"-"`

	// DatabaseName: name of the database to be created.
	DatabaseName string `json:"database_name"`
}

// DatabaseAPICreateDatabaseUserRequest: database api create database user request.
type DatabaseAPICreateDatabaseUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Username: name of the user to create.
	Username string `json:"username"`

	// Password: password of the user to create.
	Password string `json:"password"`
}

// DatabaseAPIDeleteDatabaseRequest: database api delete database request.
type DatabaseAPIDeleteDatabaseRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// DatabaseName: name of the database to delete.
	DatabaseName string `json:"-"`
}

// DatabaseAPIDeleteDatabaseUserRequest: database api delete database user request.
type DatabaseAPIDeleteDatabaseUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Username: name of the database user to delete.
	Username string `json:"-"`
}

// DatabaseAPIGetDatabaseRequest: database api get database request.
type DatabaseAPIGetDatabaseRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// DatabaseName: name of the database.
	DatabaseName string `json:"-"`
}

// DatabaseAPIGetDatabaseUserRequest: database api get database user request.
type DatabaseAPIGetDatabaseUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Username: name of the database user to retrieve details.
	Username string `json:"-"`
}

// DatabaseAPIListDatabaseUsersRequest: database api list database users request.
type DatabaseAPIListDatabaseUsersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Page: page number (must be a positive integer).
	Page *int32 `json:"-"`

	// PageSize: number of database users to return (must be a positive integer lower or equal to 100).
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of database users in the response.
	// Default value: username_asc
	OrderBy ListDatabaseUsersRequestOrderBy `json:"-"`
}

// DatabaseAPIListDatabasesRequest: database api list databases request.
type DatabaseAPIListDatabasesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Page: page number (must be a positive integer).
	Page *int32 `json:"-"`

	// PageSize: number of databases to return (must be a positive integer lower or equal to 100).
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of databases in the response.
	// Default value: database_name_asc
	OrderBy ListDatabasesRequestOrderBy `json:"-"`
}

// DatabaseAPIUnassignDatabaseUserRequest: database api unassign database user request.
type DatabaseAPIUnassignDatabaseUserRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// DatabaseName: name of the database to be unassigned.
	DatabaseName string `json:"-"`

	// Username: name of the user to unassign.
	Username string `json:"username"`
}

// FtpAccountAPIChangeFtpAccountPasswordRequest: ftp account api change ftp account password request.
type FtpAccountAPIChangeFtpAccountPasswordRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Username: username of the FTP account.
	Username string `json:"-"`

	// Password: new password for the FTP account.
	Password string `json:"password"`
}

// FtpAccountAPICreateFtpAccountRequest: ftp account api create ftp account request.
type FtpAccountAPICreateFtpAccountRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Username: username for the new FTP account.
	Username string `json:"username"`

	// Path: path for the new FTP account.
	Path string `json:"path"`

	// Password: password for the new FTP account.
	Password string `json:"password"`
}

// FtpAccountAPIListFtpAccountsRequest: ftp account api list ftp accounts request.
type FtpAccountAPIListFtpAccountsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Page: page number (must be a positive integer).
	Page *int32 `json:"-"`

	// PageSize: number of FTP accounts to return (must be a positive integer lower or equal to 100).
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of FTP accounts in the response.
	// Default value: username_asc
	OrderBy ListFtpAccountsRequestOrderBy `json:"-"`

	// Domain: domain to filter the FTP accounts.
	Domain *string `json:"-"`
}

// FtpAccountAPIRemoveFtpAccountRequest: ftp account api remove ftp account request.
type FtpAccountAPIRemoveFtpAccountRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Username: username of the FTP account to be deleted.
	Username string `json:"-"`
}

// Hosting: hosting.
type Hosting struct {
	// ID: ID of the Web Hosting plan.
	ID string `json:"id"`

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
	PlatformNumber int32 `json:"platform_number"`

	// OfferID: ID of the active offer for the Web Hosting plan.
	OfferID string `json:"offer_id"`

	// OfferName: name of the active offer for the Web Hosting plan.
	OfferName string `json:"offer_name"`

	// Domain: main domain associated with the Web Hosting plan.
	Domain string `json:"domain"`

	// Tags: list of tags associated with the Web Hosting plan.
	Tags []string `json:"tags"`

	// Options: list of the Web Hosting plan options.
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

	// ContactEmail: contact email used for the hosting.
	ContactEmail string `json:"contact_email"`

	// Region: region where the Web Hosting plan is hosted.
	Region scw.Region `json:"region"`
}

// HostingAPICreateHostingRequest: hosting api create hosting request.
type HostingAPICreateHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OfferID: ID of the selected offer for the Web Hosting plan.
	OfferID string `json:"offer_id"`

	// ProjectID: ID of the Scaleway Project in which to create the Web Hosting plan.
	ProjectID string `json:"project_id"`

	// Email: contact email for the Web Hosting client.
	Email string `json:"email"`

	// Tags: list of tags for the Web Hosting plan.
	Tags []string `json:"tags"`

	// Domain: domain name to link to the Web Hosting plan. You must already own this domain name, and have completed the DNS validation process beforehand.
	Domain string `json:"domain"`

	// OfferOptions: list of the Web Hosting plan options IDs with their quantities.
	OfferOptions []*OfferOptionRequest `json:"offer_options"`

	// Language: default language for the control panel interface.
	// Default value: unknown_language_code
	Language std.LanguageCode `json:"language"`

	// DomainConfiguration: indicates whether to update hosting domain name servers and DNS records for domains managed by Scaleway Elements.
	DomainConfiguration *CreateHostingRequestDomainConfiguration `json:"domain_configuration,omitempty"`

	// SkipWelcomeEmail: indicates whether to skip a welcome email to the contact email containing hosting info.
	SkipWelcomeEmail *bool `json:"skip_welcome_email,omitempty"`
}

// HostingAPICreateSessionRequest: hosting api create session request.
type HostingAPICreateSessionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: hosting ID.
	HostingID string `json:"-"`
}

// HostingAPIDeleteHostingRequest: hosting api delete hosting request.
type HostingAPIDeleteHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: hosting ID.
	HostingID string `json:"-"`
}

// HostingAPIGetHostingRequest: hosting api get hosting request.
type HostingAPIGetHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: hosting ID.
	HostingID string `json:"-"`
}

// HostingAPIGetResourceSummaryRequest: hosting api get resource summary request.
type HostingAPIGetResourceSummaryRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: hosting ID.
	HostingID string `json:"-"`
}

// HostingAPIListHostingsRequest: hosting api list hostings request.
type HostingAPIListHostingsRequest struct {
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

// HostingAPIResetHostingPasswordRequest: hosting api reset hosting password request.
type HostingAPIResetHostingPasswordRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting.
	HostingID string `json:"-"`
}

// HostingAPIUpdateHostingRequest: hosting api update hosting request.
type HostingAPIUpdateHostingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: hosting ID.
	HostingID string `json:"-"`

	// Email: new contact email for the Web Hosting plan.
	Email *string `json:"email,omitempty"`

	// Tags: new tags for the Web Hosting plan.
	Tags *[]string `json:"tags,omitempty"`

	// OfferOptions: list of the Web Hosting plan options IDs with their quantities.
	OfferOptions []*OfferOptionRequest `json:"offer_options"`

	// OfferID: ID of the new offer for the Web Hosting plan.
	OfferID *string `json:"offer_id,omitempty"`

	// Protected: whether the hosting is protected or not.
	Protected *bool `json:"protected,omitempty"`
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

// ListDatabaseUsersResponse: list database users response.
type ListDatabaseUsersResponse struct {
	// TotalCount: total number of database users.
	TotalCount uint64 `json:"total_count"`

	// Users: list of database users.
	Users []*DatabaseUser `json:"users"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatabaseUsersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatabaseUsersResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListDatabaseUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Users = append(r.Users, results.Users...)
	r.TotalCount += uint64(len(results.Users))
	return uint64(len(results.Users)), nil
}

// ListDatabasesResponse: list databases response.
type ListDatabasesResponse struct {
	// TotalCount: total number of databases.
	TotalCount uint64 `json:"total_count"`

	// Databases: list of databases.
	Databases []*Database `json:"databases"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatabasesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatabasesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListDatabasesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Databases = append(r.Databases, results.Databases...)
	r.TotalCount += uint64(len(results.Databases))
	return uint64(len(results.Databases)), nil
}

// ListFtpAccountsResponse: list ftp accounts response.
type ListFtpAccountsResponse struct {
	// TotalCount: total number of FTP accounts.
	TotalCount uint64 `json:"total_count"`

	// FtpAccounts: list of FTP accounts.
	FtpAccounts []*FtpAccount `json:"ftp_accounts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFtpAccountsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFtpAccountsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListFtpAccountsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.FtpAccounts = append(r.FtpAccounts, results.FtpAccounts...)
	r.TotalCount += uint64(len(results.FtpAccounts))
	return uint64(len(results.FtpAccounts)), nil
}

// ListHostingsResponse: list hostings response.
type ListHostingsResponse struct {
	// TotalCount: number of Web Hosting plans returned.
	TotalCount uint64 `json:"total_count"`

	// Hostings: list of Web Hosting plans.
	Hostings []*HostingSummary `json:"hostings"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListHostingsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListHostingsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListHostingsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Hostings = append(r.Hostings, results.Hostings...)
	r.TotalCount += uint64(len(results.Hostings))
	return uint64(len(results.Hostings)), nil
}

// ListMailAccountsResponse: list mail accounts response.
type ListMailAccountsResponse struct {
	// TotalCount: total number of mail accounts.
	TotalCount uint64 `json:"total_count"`

	// MailAccounts: list of mail accounts.
	MailAccounts []*MailAccount `json:"mail_accounts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListMailAccountsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListMailAccountsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListMailAccountsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.MailAccounts = append(r.MailAccounts, results.MailAccounts...)
	r.TotalCount += uint64(len(results.MailAccounts))
	return uint64(len(results.MailAccounts)), nil
}

// ListOffersResponse: list offers response.
type ListOffersResponse struct {
	// TotalCount: total number of offers.
	TotalCount uint64 `json:"total_count"`

	// Offers: list of offers.
	Offers []*Offer `json:"offers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOffersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOffersResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListOffersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Offers = append(r.Offers, results.Offers...)
	r.TotalCount += uint64(len(results.Offers))
	return uint64(len(results.Offers)), nil
}

// ListWebsitesResponse: list websites response.
type ListWebsitesResponse struct {
	// TotalCount: total number of websites.
	TotalCount uint64 `json:"total_count"`

	// Websites: list of websites.
	Websites []*Website `json:"websites"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListWebsitesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListWebsitesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListWebsitesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Websites = append(r.Websites, results.Websites...)
	r.TotalCount += uint64(len(results.Websites))
	return uint64(len(results.Websites)), nil
}

// MailAccountAPIChangeMailAccountPasswordRequest: mail account api change mail account password request.
type MailAccountAPIChangeMailAccountPasswordRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Domain: domain part of the mail account address.
	Domain string `json:"domain"`

	// Username: username part of the mail account address.
	Username string `json:"username"`

	// Password: new password for the mail account.
	Password string `json:"password"`
}

// MailAccountAPICreateMailAccountRequest: mail account api create mail account request.
type MailAccountAPICreateMailAccountRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Domain: domain part of the mail account address.
	Domain string `json:"domain"`

	// Username: username part address of the mail account address.
	Username string `json:"username"`

	// Password: password for the new mail account.
	Password string `json:"password"`
}

// MailAccountAPIListMailAccountsRequest: mail account api list mail accounts request.
type MailAccountAPIListMailAccountsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Page: page number (must be a positive integer).
	Page *int32 `json:"-"`

	// PageSize: number of mail accounts to return (must be a positive integer lower or equal to 100).
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of mail accounts in the response.
	// Default value: username_asc
	OrderBy ListMailAccountsRequestOrderBy `json:"-"`

	// Domain: domain to filter the mail accounts.
	Domain *string `json:"-"`
}

// MailAccountAPIRemoveMailAccountRequest: mail account api remove mail account request.
type MailAccountAPIRemoveMailAccountRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Domain: domain part of the mail account address.
	Domain string `json:"domain"`

	// Username: username part of the mail account address.
	Username string `json:"username"`
}

// OfferAPIListOffersRequest: offer api list offers request.
type OfferAPIListOffersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number (must be a positive integer).
	Page *int32 `json:"-"`

	// PageSize: number of websites to return (must be a positive integer lower or equal to 100).
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order for Web Hosting offers in the response.
	// Default value: price_asc
	OrderBy ListOffersRequestOrderBy `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID *string `json:"-"`

	// ControlPanels: name of the control panel(s) to filter for.
	ControlPanels []string `json:"-"`
}

// ResetHostingPasswordResponse: reset hosting password response.
type ResetHostingPasswordResponse struct {
	// OneTimePassword: new temporary password.
	OneTimePassword string `json:"one_time_password"`
}

// ResourceSummary: resource summary.
type ResourceSummary struct {
	// DatabasesCount: total number of active databases in the Web Hosting plan.
	DatabasesCount uint32 `json:"databases_count"`

	// MailAccountsCount: total number of active email accounts in the Web Hosting plan.
	MailAccountsCount uint32 `json:"mail_accounts_count"`

	// FtpAccountsCount: total number of active FTP accounts in the Web Hosting plan.
	FtpAccountsCount uint32 `json:"ftp_accounts_count"`

	// WebsitesCount: total number of active domains in the the Web Hosting plan.
	WebsitesCount uint32 `json:"websites_count"`
}

// Session: session.
type Session struct {
	// URL: logged user's session URL.
	URL string `json:"url"`
}

// WebsiteAPIListWebsitesRequest: website api list websites request.
type WebsiteAPIListWebsitesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting plan.
	HostingID string `json:"-"`

	// Page: page number (must be a positive integer).
	Page *int32 `json:"-"`

	// PageSize: number of websites to return (must be a positive integer lower or equal to 100).
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order for Web Hosting websites in the response.
	// Default value: domain_asc
	OrderBy ListWebsitesRequestOrderBy `json:"-"`
}

// This API allows you to manage your Web Hosting services.
type ControlPanelAPI struct {
	client *scw.Client
}

// NewControlPanelAPI returns a ControlPanelAPI object from a Scaleway client.
func NewControlPanelAPI(client *scw.Client) *ControlPanelAPI {
	return &ControlPanelAPI{
		client: client,
	}
}
func (s *ControlPanelAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms}
}

// ListControlPanels: "List the control panels type: cpanel or plesk.".
func (s *ControlPanelAPI) ListControlPanels(req *ControlPanelAPIListControlPanelsRequest, opts ...scw.RequestOption) (*ListControlPanelsResponse, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/control-panels",
		Query:  query,
	}

	var resp ListControlPanelsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage your databases and database users for your Web Hosting services.
type DatabaseAPI struct {
	client *scw.Client
}

// NewDatabaseAPI returns a DatabaseAPI object from a Scaleway client.
func NewDatabaseAPI(client *scw.Client) *DatabaseAPI {
	return &DatabaseAPI{
		client: client,
	}
}
func (s *DatabaseAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// CreateDatabase: "Create a new database within your hosting plan".
func (s *DatabaseAPI) CreateDatabase(req *DatabaseAPICreateDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/databases",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Database

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDatabases: "List all databases within your hosting plan".
func (s *DatabaseAPI) ListDatabases(req *DatabaseAPIListDatabasesRequest, opts ...scw.RequestOption) (*ListDatabasesResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/databases",
		Query:  query,
	}

	var resp ListDatabasesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDatabase: "Get details of a database within your hosting plan".
func (s *DatabaseAPI) GetDatabase(req *DatabaseAPIGetDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
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

	if fmt.Sprint(req.DatabaseName) == "" {
		return nil, errors.New("field DatabaseName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/databases/" + fmt.Sprint(req.DatabaseName) + "",
	}

	var resp Database

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDatabase: "Delete a database within your hosting plan".
func (s *DatabaseAPI) DeleteDatabase(req *DatabaseAPIDeleteDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
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

	if fmt.Sprint(req.DatabaseName) == "" {
		return nil, errors.New("field DatabaseName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/databases/" + fmt.Sprint(req.DatabaseName) + "",
	}

	var resp Database

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDatabaseUser: "Create a new database user".
func (s *DatabaseAPI) CreateDatabaseUser(req *DatabaseAPICreateDatabaseUserRequest, opts ...scw.RequestOption) (*DatabaseUser, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/databases-users",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseUser

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDatabaseUsers: "List all database users".
func (s *DatabaseAPI) ListDatabaseUsers(req *DatabaseAPIListDatabaseUsersRequest, opts ...scw.RequestOption) (*ListDatabaseUsersResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/database-users",
		Query:  query,
	}

	var resp ListDatabaseUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDatabaseUser: "Get details of a database user".
func (s *DatabaseAPI) GetDatabaseUser(req *DatabaseAPIGetDatabaseUserRequest, opts ...scw.RequestOption) (*DatabaseUser, error) {
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

	if fmt.Sprint(req.Username) == "" {
		return nil, errors.New("field Username cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/databases-users/" + fmt.Sprint(req.Username) + "",
	}

	var resp DatabaseUser

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDatabaseUser: "Delete a database user".
func (s *DatabaseAPI) DeleteDatabaseUser(req *DatabaseAPIDeleteDatabaseUserRequest, opts ...scw.RequestOption) (*DatabaseUser, error) {
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

	if fmt.Sprint(req.Username) == "" {
		return nil, errors.New("field Username cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/database-users/" + fmt.Sprint(req.Username) + "",
	}

	var resp DatabaseUser

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeDatabaseUserPassword: "Change the password of a database user".
func (s *DatabaseAPI) ChangeDatabaseUserPassword(req *DatabaseAPIChangeDatabaseUserPasswordRequest, opts ...scw.RequestOption) (*DatabaseUser, error) {
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

	if fmt.Sprint(req.Username) == "" {
		return nil, errors.New("field Username cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/databases-users/" + fmt.Sprint(req.Username) + "/change-password",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseUser

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AssignDatabaseUser: "Assign a database user to a database".
func (s *DatabaseAPI) AssignDatabaseUser(req *DatabaseAPIAssignDatabaseUserRequest, opts ...scw.RequestOption) (*DatabaseUser, error) {
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

	if fmt.Sprint(req.DatabaseName) == "" {
		return nil, errors.New("field DatabaseName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/databases/" + fmt.Sprint(req.DatabaseName) + "/assign-user",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseUser

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnassignDatabaseUser: "Unassign a database user from a database".
func (s *DatabaseAPI) UnassignDatabaseUser(req *DatabaseAPIUnassignDatabaseUserRequest, opts ...scw.RequestOption) (*DatabaseUser, error) {
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

	if fmt.Sprint(req.DatabaseName) == "" {
		return nil, errors.New("field DatabaseName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/databases/" + fmt.Sprint(req.DatabaseName) + "/unassign-user",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseUser

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage your offer for your Web Hosting services.
type OfferAPI struct {
	client *scw.Client
}

// NewOfferAPI returns a OfferAPI object from a Scaleway client.
func NewOfferAPI(client *scw.Client) *OfferAPI {
	return &OfferAPI{
		client: client,
	}
}
func (s *OfferAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// ListOffers: List all available hosting offers along with their specific options.
func (s *OfferAPI) ListOffers(req *OfferAPIListOffersRequest, opts ...scw.RequestOption) (*ListOffersResponse, error) {
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
	parameter.AddToQuery(query, "hosting_id", req.HostingID)
	parameter.AddToQuery(query, "control_panels", req.ControlPanels)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/offers",
		Query:  query,
	}

	var resp ListOffersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage your Web Hosting services.
type HostingAPI struct {
	client *scw.Client
}

// NewHostingAPI returns a HostingAPI object from a Scaleway client.
func NewHostingAPI(client *scw.Client) *HostingAPI {
	return &HostingAPI{
		client: client,
	}
}
func (s *HostingAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms}
}

// CreateHosting: Order a Web Hosting plan, specifying the offer type required via the `offer_id` parameter.
func (s *HostingAPI) CreateHosting(req *HostingAPICreateHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings",
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
func (s *HostingAPI) ListHostings(req *HostingAPIListHostingsRequest, opts ...scw.RequestOption) (*ListHostingsResponse, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings",
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
func (s *HostingAPI) GetHosting(req *HostingAPIGetHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHosting: Update the details of one of your existing Web Hosting plans, specified by its `hosting_id`. You can update parameters including the contact email address, tags, options and offer.
func (s *HostingAPI) UpdateHosting(req *HostingAPIUpdateHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
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
func (s *HostingAPI) DeleteHosting(req *HostingAPIDeleteHostingRequest, opts ...scw.RequestOption) (*Hosting, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "",
	}

	var resp Hosting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSession: Create a user session.
func (s *HostingAPI) CreateSession(req *HostingAPICreateSessionRequest, opts ...scw.RequestOption) (*Session, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/sessions",
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

// ResetHostingPassword: Reset a Web Hosting plan password.
func (s *HostingAPI) ResetHostingPassword(req *HostingAPIResetHostingPasswordRequest, opts ...scw.RequestOption) (*ResetHostingPasswordResponse, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/reset-password",
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

// GetResourceSummary: Get the total counts of websites, databases, email accounts, and FTP accounts of a Web Hosting plan.
func (s *HostingAPI) GetResourceSummary(req *HostingAPIGetResourceSummaryRequest, opts ...scw.RequestOption) (*ResourceSummary, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/resource-summary",
	}

	var resp ResourceSummary

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage your FTP accounts for your Web Hosting services.
type FtpAccountAPI struct {
	client *scw.Client
}

// NewFtpAccountAPI returns a FtpAccountAPI object from a Scaleway client.
func NewFtpAccountAPI(client *scw.Client) *FtpAccountAPI {
	return &FtpAccountAPI{
		client: client,
	}
}
func (s *FtpAccountAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// CreateFtpAccount: Create a new FTP account within your hosting plan.
func (s *FtpAccountAPI) CreateFtpAccount(req *FtpAccountAPICreateFtpAccountRequest, opts ...scw.RequestOption) (*FtpAccount, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/ftp-accounts",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FtpAccount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFtpAccounts: List all FTP accounts within your hosting plan.
func (s *FtpAccountAPI) ListFtpAccounts(req *FtpAccountAPIListFtpAccountsRequest, opts ...scw.RequestOption) (*ListFtpAccountsResponse, error) {
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
	parameter.AddToQuery(query, "domain", req.Domain)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/ftp-accounts",
		Query:  query,
	}

	var resp ListFtpAccountsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveFtpAccount: Delete a specific FTP account within your hosting plan.
func (s *FtpAccountAPI) RemoveFtpAccount(req *FtpAccountAPIRemoveFtpAccountRequest, opts ...scw.RequestOption) (*FtpAccount, error) {
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

	if fmt.Sprint(req.Username) == "" {
		return nil, errors.New("field Username cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/ftp-accounts/" + fmt.Sprint(req.Username) + "",
	}

	var resp FtpAccount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeFtpAccountPassword:
func (s *FtpAccountAPI) ChangeFtpAccountPassword(req *FtpAccountAPIChangeFtpAccountPasswordRequest, opts ...scw.RequestOption) (*FtpAccount, error) {
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

	if fmt.Sprint(req.Username) == "" {
		return nil, errors.New("field Username cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/ftp-accounts/" + fmt.Sprint(req.Username) + "/change-password",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FtpAccount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage your mail accounts for your Web Hosting services.
type MailAccountAPI struct {
	client *scw.Client
}

// NewMailAccountAPI returns a MailAccountAPI object from a Scaleway client.
func NewMailAccountAPI(client *scw.Client) *MailAccountAPI {
	return &MailAccountAPI{
		client: client,
	}
}
func (s *MailAccountAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// CreateMailAccount: Create a new mail account within your hosting plan.
func (s *MailAccountAPI) CreateMailAccount(req *MailAccountAPICreateMailAccountRequest, opts ...scw.RequestOption) (*MailAccount, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/mail-accounts",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp MailAccount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListMailAccounts: List all mail accounts within your hosting plan.
func (s *MailAccountAPI) ListMailAccounts(req *MailAccountAPIListMailAccountsRequest, opts ...scw.RequestOption) (*ListMailAccountsResponse, error) {
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
	parameter.AddToQuery(query, "domain", req.Domain)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/mail-accounts",
		Query:  query,
	}

	var resp ListMailAccountsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveMailAccount: Delete a mail account within your hosting plan.
func (s *MailAccountAPI) RemoveMailAccount(req *MailAccountAPIRemoveMailAccountRequest, opts ...scw.RequestOption) (*MailAccount, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/remove-mail-account",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp MailAccount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ChangeMailAccountPassword: Update the password of a mail account within your hosting plan.
func (s *MailAccountAPI) ChangeMailAccountPassword(req *MailAccountAPIChangeMailAccountPasswordRequest, opts ...scw.RequestOption) (*MailAccount, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/change-mail-password",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp MailAccount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage your websites for your Web Hosting services.
type WebsiteAPI struct {
	client *scw.Client
}

// NewWebsiteAPI returns a WebsiteAPI object from a Scaleway client.
func NewWebsiteAPI(client *scw.Client) *WebsiteAPI {
	return &WebsiteAPI{
		client: client,
	}
}
func (s *WebsiteAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// ListWebsites: List all websites for a specific hosting.
func (s *WebsiteAPI) ListWebsites(req *WebsiteAPIListWebsitesRequest, opts ...scw.RequestOption) (*ListWebsitesResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/websites",
		Query:  query,
	}

	var resp ListWebsitesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

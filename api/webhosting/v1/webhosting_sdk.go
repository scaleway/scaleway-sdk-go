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

type BackupItemType string

const (
	// Default type returned when the backup item type is not specified.
	BackupItemTypeUnknownBackupItemType = BackupItemType("unknown_backup_item_type")
	// Complete system backup.
	BackupItemTypeFull = BackupItemType("full")
	// Backup item containing website-related files.
	BackupItemTypeWeb = BackupItemType("web")
	// Backup item for mail service data.
	BackupItemTypeMail = BackupItemType("mail")
	// Backup item for databases.
	BackupItemTypeDb = BackupItemType("db")
	// Backup item for database user accounts.
	BackupItemTypeDbUser = BackupItemType("db_user")
	// Backup item for FTP user accounts.
	BackupItemTypeFtpUser = BackupItemType("ftp_user")
	// Backup item for DNS zone configurations.
	BackupItemTypeDNSZone = BackupItemType("dns_zone")
	// Backup item for scheduled cron jobs.
	BackupItemTypeCronJob = BackupItemType("cron_job")
	// Backup item for SSL certificates.
	BackupItemTypeSslCertificate = BackupItemType("ssl_certificate")
)

func (enum BackupItemType) String() string {
	if enum == "" {
		// return default value if empty
		return string(BackupItemTypeUnknownBackupItemType)
	}
	return string(enum)
}

func (enum BackupItemType) Values() []BackupItemType {
	return []BackupItemType{
		"unknown_backup_item_type",
		"full",
		"web",
		"mail",
		"db",
		"db_user",
		"ftp_user",
		"dns_zone",
		"cron_job",
		"ssl_certificate",
	}
}

func (enum BackupItemType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BackupItemType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BackupItemType(BackupItemType(tmp).String())
	return nil
}

type BackupStatus string

const (
	// Default status returned when the backup status is not specified.
	BackupStatusUnknownBackupStatus = BackupStatus("unknown_backup_status")
	// Backup is active and available.
	BackupStatusActive = BackupStatus("active")
	// Backup is locked and cannot be changed.
	BackupStatusLocked = BackupStatus("locked")
	// Backup is currently disabled.
	BackupStatusDisabled = BackupStatus("disabled")
	// Backup is incomplete, damaged, or corrupted.
	BackupStatusDamaged = BackupStatus("damaged")
	// Backup is being restored.
	BackupStatusRestoring = BackupStatus("restoring")
)

func (enum BackupStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(BackupStatusUnknownBackupStatus)
	}
	return string(enum)
}

func (enum BackupStatus) Values() []BackupStatus {
	return []BackupStatus{
		"unknown_backup_status",
		"active",
		"locked",
		"disabled",
		"damaged",
		"restoring",
	}
}

func (enum BackupStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BackupStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BackupStatus(BackupStatus(tmp).String())
	return nil
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
		return string(DNSRecordStatusUnknownStatus)
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
		return string(DNSRecordTypeUnknownType)
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
	DNSRecordsStatusUnknownStatus = DNSRecordsStatus("unknown_status")
	DNSRecordsStatusValid         = DNSRecordsStatus("valid")
	DNSRecordsStatusInvalid       = DNSRecordsStatus("invalid")
)

func (enum DNSRecordsStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(DNSRecordsStatusUnknownStatus)
	}
	return string(enum)
}

func (enum DNSRecordsStatus) Values() []DNSRecordsStatus {
	return []DNSRecordsStatus{
		"unknown_status",
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

type DomainAction string

const (
	// Unknown action.
	DomainActionUnknownAction = DomainAction("unknown_action")
	// Action to transfer the domain to Scaleway.
	DomainActionTransfer = DomainAction("transfer")
	// Action to manage domain zone at Scaleway.
	DomainActionManageExternal = DomainAction("manage_external")
	// Action to renew the domain's registration.
	DomainActionRenew = DomainAction("renew")
)

func (enum DomainAction) String() string {
	if enum == "" {
		// return default value if empty
		return string(DomainActionUnknownAction)
	}
	return string(enum)
}

func (enum DomainAction) Values() []DomainAction {
	return []DomainAction{
		"unknown_action",
		"transfer",
		"manage_external",
		"renew",
	}
}

func (enum DomainAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainAction(DomainAction(tmp).String())
	return nil
}

type DomainAvailabilityAction string

const (
	// Unknown action.
	DomainAvailabilityActionUnknownAction = DomainAvailabilityAction("unknown_action")
	// Action to register the domain at Scaleway.
	DomainAvailabilityActionRegister = DomainAvailabilityAction("register")
	// Action to transfer the domain from another registrar to Scaleway.
	DomainAvailabilityActionTransfer = DomainAvailabilityAction("transfer")
	// Action to manage the domain as an external domain at Scaleway.
	DomainAvailabilityActionManageExternal = DomainAvailabilityAction("manage_external")
)

func (enum DomainAvailabilityAction) String() string {
	if enum == "" {
		// return default value if empty
		return string(DomainAvailabilityActionUnknownAction)
	}
	return string(enum)
}

func (enum DomainAvailabilityAction) Values() []DomainAvailabilityAction {
	return []DomainAvailabilityAction{
		"unknown_action",
		"register",
		"transfer",
		"manage_external",
	}
}

func (enum DomainAvailabilityAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainAvailabilityAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainAvailabilityAction(DomainAvailabilityAction(tmp).String())
	return nil
}

type DomainAvailabilityStatus string

const (
	// The status of the domain is unknown.
	DomainAvailabilityStatusUnknownStatus = DomainAvailabilityStatus("unknown_status")
	// The domain is available for registration.
	DomainAvailabilityStatusAvailable = DomainAvailabilityStatus("available")
	// The domain is not available for registration.
	DomainAvailabilityStatusNotAvailable = DomainAvailabilityStatus("not_available")
	// The domain is already owned by the user.
	DomainAvailabilityStatusOwned = DomainAvailabilityStatus("owned")
	// The domain is being validated.
	DomainAvailabilityStatusValidating = DomainAvailabilityStatus("validating")
	// The domain is in an error status, it should be checked in Domain&DNS.
	DomainAvailabilityStatusError = DomainAvailabilityStatus("error")
)

func (enum DomainAvailabilityStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(DomainAvailabilityStatusUnknownStatus)
	}
	return string(enum)
}

func (enum DomainAvailabilityStatus) Values() []DomainAvailabilityStatus {
	return []DomainAvailabilityStatus{
		"unknown_status",
		"available",
		"not_available",
		"owned",
		"validating",
		"error",
	}
}

func (enum DomainAvailabilityStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainAvailabilityStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainAvailabilityStatus(DomainAvailabilityStatus(tmp).String())
	return nil
}

type DomainDNSAction string

const (
	// Unknown DNS action.
	DomainDNSActionUnknownDNSAction = DomainDNSAction("unknown_dns_action")
	// Automatically configure all DNS records for the domain.
	DomainDNSActionAutoConfigAllRecords = DomainDNSAction("auto_config_all_records")
	// Automatically configure web-related DNS records (e.g., A, CNAME).
	DomainDNSActionAutoConfigWebRecords = DomainDNSAction("auto_config_web_records")
	// Automatically configure mail-related DNS records (e.g., MX, SPF, DKIM).
	DomainDNSActionAutoConfigMailRecords = DomainDNSAction("auto_config_mail_records")
	// Automatically configure the domain's name servers to point to Web Hosting name servers.
	DomainDNSActionAutoConfigNameservers = DomainDNSAction("auto_config_nameservers")
	// No automatic domain configuration. Users must configure their domain for the Web Hosting to work.
	DomainDNSActionAutoConfigNone = DomainDNSAction("auto_config_none")
)

func (enum DomainDNSAction) String() string {
	if enum == "" {
		// return default value if empty
		return string(DomainDNSActionUnknownDNSAction)
	}
	return string(enum)
}

func (enum DomainDNSAction) Values() []DomainDNSAction {
	return []DomainDNSAction{
		"unknown_dns_action",
		"auto_config_all_records",
		"auto_config_web_records",
		"auto_config_mail_records",
		"auto_config_nameservers",
		"auto_config_none",
	}
}

func (enum DomainDNSAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainDNSAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainDNSAction(DomainDNSAction(tmp).String())
	return nil
}

type DomainStatus string

const (
	// The status of the domain is unknown.
	DomainStatusUnknownStatus = DomainStatus("unknown_status")
	// The domain is valid and operational.
	DomainStatusValid = DomainStatus("valid")
	// The domain is invalid, potentially due to incorrect settings or configuration issues.
	DomainStatusInvalid = DomainStatus("invalid")
	// The domain is in the process of being validated.
	DomainStatusValidating = DomainStatus("validating")
	// The domain is in an error status, it should be checked in Domain&DNS.
	DomainStatusError = DomainStatus("error")
)

func (enum DomainStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(DomainStatusUnknownStatus)
	}
	return string(enum)
}

func (enum DomainStatus) Values() []DomainStatus {
	return []DomainStatus{
		"unknown_status",
		"valid",
		"invalid",
		"validating",
		"error",
	}
}

func (enum DomainStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainStatus(DomainStatus(tmp).String())
	return nil
}

type DomainZoneOwner string

const (
	// The zone owner is unknown or not specified.
	DomainZoneOwnerUnknownZoneOwner = DomainZoneOwner("unknown_zone_owner")
	// The domain's DNS is managed externally.
	DomainZoneOwnerExternal = DomainZoneOwner("external")
	// The domain's DNS is managed by Scaleway.
	DomainZoneOwnerScaleway = DomainZoneOwner("scaleway")
	// The domain's DNS is managed by Online.
	DomainZoneOwnerOnline = DomainZoneOwner("online")
	// The domain's DNS is managed by the Web Hosting platform.
	DomainZoneOwnerWebhosting = DomainZoneOwner("webhosting")
)

func (enum DomainZoneOwner) String() string {
	if enum == "" {
		// return default value if empty
		return string(DomainZoneOwnerUnknownZoneOwner)
	}
	return string(enum)
}

func (enum DomainZoneOwner) Values() []DomainZoneOwner {
	return []DomainZoneOwner{
		"unknown_zone_owner",
		"external",
		"scaleway",
		"online",
		"webhosting",
	}
}

func (enum DomainZoneOwner) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DomainZoneOwner) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DomainZoneOwner(DomainZoneOwner(tmp).String())
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
	HostingStatusUpdating      = HostingStatus("updating")
)

func (enum HostingStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(HostingStatusUnknownStatus)
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
		"updating",
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

type ListBackupsRequestOrderBy string

const (
	// Order by backup creation date, most recent first (default).
	ListBackupsRequestOrderByCreatedAtDesc = ListBackupsRequestOrderBy("created_at_desc")
	// Order by backup creation date, oldest first.
	ListBackupsRequestOrderByCreatedAtAsc = ListBackupsRequestOrderBy("created_at_asc")
)

func (enum ListBackupsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListBackupsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListBackupsRequestOrderBy) Values() []ListBackupsRequestOrderBy {
	return []ListBackupsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
	}
}

func (enum ListBackupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListBackupsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListBackupsRequestOrderBy(ListBackupsRequestOrderBy(tmp).String())
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
		return string(ListDatabaseUsersRequestOrderByUsernameAsc)
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
		return string(ListDatabasesRequestOrderByDatabaseNameAsc)
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
		return string(ListFtpAccountsRequestOrderByUsernameAsc)
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
		return string(ListHostingsRequestOrderByCreatedAtAsc)
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
		return string(ListMailAccountsRequestOrderByUsernameAsc)
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
		return string(ListOffersRequestOrderByPriceAsc)
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
		return string(ListWebsitesRequestOrderByDomainAsc)
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

type NameserverStatus string

const (
	NameserverStatusUnknownStatus = NameserverStatus("unknown_status")
	NameserverStatusValid         = NameserverStatus("valid")
	NameserverStatusInvalid       = NameserverStatus("invalid")
)

func (enum NameserverStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(NameserverStatusUnknownStatus)
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

type OfferOptionName string

const (
	OfferOptionNameUnknownName    = OfferOptionName("unknown_name")
	OfferOptionNameDomainCount    = OfferOptionName("domain_count")
	OfferOptionNameEmailCount     = OfferOptionName("email_count")
	OfferOptionNameStorageGb      = OfferOptionName("storage_gb")
	OfferOptionNameVcpuCount      = OfferOptionName("vcpu_count")
	OfferOptionNameRAMGb          = OfferOptionName("ram_gb")
	OfferOptionNameBackup         = OfferOptionName("backup")
	OfferOptionNameDedicatedIP    = OfferOptionName("dedicated_ip")
	OfferOptionNameEmailStorageGb = OfferOptionName("email_storage_gb")
	OfferOptionNameDatabaseCount  = OfferOptionName("database_count")
	OfferOptionNameSupport        = OfferOptionName("support")
)

func (enum OfferOptionName) String() string {
	if enum == "" {
		// return default value if empty
		return string(OfferOptionNameUnknownName)
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
		"email_storage_gb",
		"database_count",
		"support",
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
	OfferOptionWarningUsageLowWarning      = OfferOptionWarning("usage_low_warning")
)

func (enum OfferOptionWarning) String() string {
	if enum == "" {
		// return default value if empty
		return string(OfferOptionWarningUnknownWarning)
	}
	return string(enum)
}

func (enum OfferOptionWarning) Values() []OfferOptionWarning {
	return []OfferOptionWarning{
		"unknown_warning",
		"quota_exceeded_warning",
		"usage_low_warning",
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

type PlatformPlatformGroup string

const (
	PlatformPlatformGroupUnknownGroup = PlatformPlatformGroup("unknown_group")
	PlatformPlatformGroupDefault      = PlatformPlatformGroup("default")
	PlatformPlatformGroupPremium      = PlatformPlatformGroup("premium")
)

func (enum PlatformPlatformGroup) String() string {
	if enum == "" {
		// return default value if empty
		return string(PlatformPlatformGroupUnknownGroup)
	}
	return string(enum)
}

func (enum PlatformPlatformGroup) Values() []PlatformPlatformGroup {
	return []PlatformPlatformGroup{
		"unknown_group",
		"default",
		"premium",
	}
}

func (enum PlatformPlatformGroup) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlatformPlatformGroup) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlatformPlatformGroup(PlatformPlatformGroup(tmp).String())
	return nil
}

// AutoConfigDomainDNS: auto config domain dns.
type AutoConfigDomainDNS struct {
	// Nameservers: whether or not to synchronize domain nameservers.
	Nameservers bool `json:"nameservers"`

	// WebRecords: whether or not to synchronize web records.
	WebRecords bool `json:"web_records"`

	// MailRecords: whether or not to synchronize mail records.
	MailRecords bool `json:"mail_records"`

	// AllRecords: whether or not to synchronize all types of records. Takes priority over the other fields.
	AllRecords bool `json:"all_records"`

	// None: no automatic domain configuration. Users must configure their domain for the Web Hosting to work.
	None bool `json:"none"`
}

// PlatformControlPanelURLs: platform control panel ur ls.
type PlatformControlPanelURLs struct {
	// Dashboard: URL to connect to the hosting control panel dashboard.
	Dashboard string `json:"dashboard"`

	// Webmail: URL to connect to the hosting Webmail interface.
	Webmail string `json:"webmail"`
}

// HostingDomainCustomDomain: hosting domain custom domain.
type HostingDomainCustomDomain struct {
	// Domain: custom domain linked to the hosting plan.
	Domain string `json:"domain"`

	// DomainStatus: status of the custom domain verification.
	// Default value: unknown_status
	DomainStatus DomainStatus `json:"domain_status"`

	// DNSStatus: status of the DNS configuration for the custom domain.
	// Default value: unknown_status
	DNSStatus DNSRecordsStatus `json:"dns_status"`

	// AutoConfigDomainDNS: indicates whether to auto-configure DNS for this domain.
	AutoConfigDomainDNS *AutoConfigDomainDNS `json:"auto_config_domain_dns"`
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

	// Price: price of the option for 1 value.
	Price *scw.Money `json:"price"`
}

// PlatformControlPanel: platform control panel.
type PlatformControlPanel struct {
	// Name: name of the control panel.
	Name string `json:"name"`

	// URLs: URL to connect to control panel dashboard and to Webmail interface.
	URLs *PlatformControlPanelURLs `json:"urls"`
}

// BackupItem: backup item.
type BackupItem struct {
	// ID: ID of the item.
	ID string `json:"id"`

	// Name: name of the item (e.g., `database name`, `email address`).
	Name string `json:"name"`

	// Type: type of the item (e.g., email, database, FTP).
	// Default value: unknown_backup_item_type
	Type BackupItemType `json:"type"`

	// Size: size of the item in bytes.
	Size scw.Size `json:"size"`

	// Status: status of the item. Available values are `active`, `damaged`, and `restoring`.
	// Default value: unknown_backup_status
	Status BackupStatus `json:"status"`

	// CreatedAt: date and time at which this item was backed up.
	CreatedAt *time.Time `json:"created_at"`
}

// HostingDomain: hosting domain.
type HostingDomain struct {
	// Subdomain: optional free subdomain linked to the Web Hosting plan.
	Subdomain string `json:"subdomain"`

	// CustomDomain: optional custom domain linked to the Web Hosting plan.
	CustomDomain *HostingDomainCustomDomain `json:"custom_domain"`
}

// CreateDatabaseRequestUser: create database request user.
type CreateDatabaseRequestUser struct {
	Username string `json:"username"`

	Password string `json:"password"`
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

// SyncDomainDNSRecordsRequestRecord: sync domain dns records request record.
type SyncDomainDNSRecordsRequestRecord struct {
	Name string `json:"name"`

	// Type: default value: unknown_type
	Type DNSRecordType `json:"type"`
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

	// RawData: record representation as it appears in the zone file or DNS management system.
	RawData string `json:"raw_data"`
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

// HostingUser: hosting user.
type HostingUser struct {
	// Username: main Web Hosting control panel username.
	Username string `json:"username"`

	// Deprecated: OneTimePassword: one-time-password used for the first login to the control panel, cleared after first use (deprecated, use password_b64 instead).
	OneTimePassword *string `json:"one_time_password,omitempty"`

	// ContactEmail: contact email used for the hosting.
	ContactEmail string `json:"contact_email"`

	// OneTimePasswordB64: one-time-password used for the first login to the control panel, cleared after first use, encoded in base64.
	OneTimePasswordB64 *string `json:"one_time_password_b64"`
}

// Offer: offer.
type Offer struct {
	// ID: offer ID.
	ID string `json:"id"`

	// Name: offer name.
	Name string `json:"name"`

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

	// QuotaWarning: defines a warning if the maximum value for an option in the offer is exceeded.
	// Default value: unknown_warning
	QuotaWarning OfferOptionWarning `json:"quota_warning"`
}

// Platform: platform.
type Platform struct {
	// Hostname: hostname of the host platform.
	Hostname string `json:"hostname"`

	// Number: number of the host platform.
	Number int32 `json:"number"`

	// GroupName: group name of the hosting's host platform.
	// Default value: unknown_group
	GroupName PlatformPlatformGroup `json:"group_name"`

	// IPv4: iPv4 address of the hosting's host platform.
	IPv4 net.IP `json:"ipv4"`

	// IPv6: iPv6 address of the hosting's host platform.
	IPv6 net.IP `json:"ipv6"`

	// ControlPanel: details of the platform control panel.
	ControlPanel *PlatformControlPanel `json:"control_panel"`
}

// BackupItemGroup: backup item group.
type BackupItemGroup struct {
	// Type: type of items (e.g., email, database, FTP).
	// Default value: unknown_backup_item_type
	Type BackupItemType `json:"type"`

	// Items: list of individual backup items of this type.
	Items []*BackupItem `json:"items"`
}

// Backup: backup.
type Backup struct {
	// ID: ID of the backup.
	ID string `json:"id"`

	// Size: total size of the backup in bytes.
	Size scw.Size `json:"size"`

	// CreatedAt: creation date of the backup.
	CreatedAt *time.Time `json:"created_at"`

	// Status: status of the backup. Available values are `active`, `locked`, and `restoring`.
	// Default value: unknown_backup_status
	Status BackupStatus `json:"status"`

	// TotalItems: total number of restorable items in the backup.
	TotalItems uint32 `json:"total_items"`
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
	Status HostingStatus `json:"status"`

	// Deprecated: Domain: main domain associated with the Web Hosting plan (deprecated, use domain_info).
	Domain *string `json:"domain,omitempty"`

	// Protected: whether the hosting is protected or not.
	Protected bool `json:"protected"`

	// Deprecated: DNSStatus: DNS status of the Web Hosting plan.
	// Default value: unknown_status
	DNSStatus *DNSRecordsStatus `json:"dns_status,omitempty"`

	// OfferName: name of the active offer for the Web Hosting plan.
	OfferName string `json:"offer_name"`

	// Deprecated: DomainStatus: main domain status of the Web Hosting plan.
	// Default value: unknown_status
	DomainStatus *DomainStatus `json:"domain_status,omitempty"`

	// Region: region where the Web Hosting plan is hosted.
	Region scw.Region `json:"region"`

	// DomainInfo: domain configuration block (subdomain, optional custom domain, and DNS settings).
	DomainInfo *HostingDomain `json:"domain_info"`
}

// MailAccount: mail account.
type MailAccount struct {
	// Domain: domain part of the mail account address.
	Domain string `json:"domain"`

	// Username: username part address of the mail account address.
	Username string `json:"username"`
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

// DomainAvailability: domain availability.
type DomainAvailability struct {
	// Name: fully qualified domain name (FQDN).
	Name string `json:"name"`

	// ZoneName: DNS zone associated with the domain.
	ZoneName string `json:"zone_name"`

	// Status: availability status of the domain.
	// Default value: unknown_status
	Status DomainAvailabilityStatus `json:"status"`

	// AvailableActions: a list of actions that can be performed on the domain.
	AvailableActions []DomainAvailabilityAction `json:"available_actions"`

	// CanCreateHosting: whether a hosting can be created for this domain.
	CanCreateHosting bool `json:"can_create_hosting"`

	// Price: price for registering the domain.
	Price *scw.Money `json:"price"`
}

// BackupAPIGetBackupRequest: backup api get backup request.
type BackupAPIGetBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting account.
	HostingID string `json:"-"`

	// BackupID: ID of the backup to retrieve.
	BackupID string `json:"-"`
}

// BackupAPIListBackupItemsRequest: backup api list backup items request.
type BackupAPIListBackupItemsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting account.
	HostingID string `json:"-"`

	// BackupID: ID of the backup to list items from.
	BackupID string `json:"backup_id"`
}

// BackupAPIListBackupsRequest: backup api list backups request.
type BackupAPIListBackupsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting account.
	HostingID string `json:"-"`

	// Page: page number to retrieve.
	Page *int32 `json:"-"`

	// PageSize: number of backups to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return the list of backups.
	// Default value: created_at_desc
	OrderBy ListBackupsRequestOrderBy `json:"-"`
}

// BackupAPIRestoreBackupItemsRequest: backup api restore backup items request.
type BackupAPIRestoreBackupItemsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting account.
	HostingID string `json:"-"`

	// ItemIDs: list of backup item IDs to restore individually.
	ItemIDs []string `json:"item_ids"`
}

// BackupAPIRestoreBackupRequest: backup api restore backup request.
type BackupAPIRestoreBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: UUID of the hosting account.
	HostingID string `json:"-"`

	// BackupID: ID of the backup to fully restore.
	BackupID string `json:"-"`
}

// CheckUserOwnsDomainResponse: check user owns domain response.
type CheckUserOwnsDomainResponse struct {
	// OwnsDomain: indicates whether the specified project owns the domain.
	OwnsDomain bool `json:"owns_domain"`
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

// DNSAPICheckUserOwnsDomainRequest: dnsapi check user owns domain request.
type DNSAPICheckUserOwnsDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Domain: domain for which ownership is to be verified.
	Domain string `json:"-"`

	// ProjectID: ID of the project currently in use.
	ProjectID string `json:"project_id"`
}

// DNSAPIGetDomainDNSRecordsRequest: dnsapi get domain dns records request.
type DNSAPIGetDomainDNSRecordsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Domain: domain associated with the DNS records.
	Domain string `json:"-"`
}

// DNSAPIGetDomainRequest: dnsapi get domain request.
type DNSAPIGetDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainName: domain name to get.
	DomainName string `json:"-"`

	// ProjectID: ID of the Scaleway Project in which to get the domain to create the Web Hosting plan.
	ProjectID string `json:"project_id"`
}

// DNSAPISearchDomainsRequest: dnsapi search domains request.
type DNSAPISearchDomainsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DomainName: domain name to search.
	DomainName string `json:"-"`

	// ProjectID: ID of the Scaleway Project in which to search the domain to create the Web Hosting plan.
	ProjectID string `json:"-"`
}

// DNSAPISyncDomainDNSRecordsRequest: dnsapi sync domain dns records request.
type DNSAPISyncDomainDNSRecordsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Domain: domain for which the DNS records will be synchronized.
	Domain string `json:"-"`

	// Deprecated: UpdateWebRecords: whether or not to synchronize the web records (deprecated, use auto_config_domain_dns).
	UpdateWebRecords *bool `json:"update_web_records,omitempty"`

	// Deprecated: UpdateMailRecords: whether or not to synchronize the mail records (deprecated, use auto_config_domain_dns).
	UpdateMailRecords *bool `json:"update_mail_records,omitempty"`

	// Deprecated: UpdateAllRecords: whether or not to synchronize all types of records. This one has priority (deprecated, use auto_config_domain_dns).
	UpdateAllRecords *bool `json:"update_all_records,omitempty"`

	// Deprecated: UpdateNameservers: whether or not to synchronize domain nameservers (deprecated, use auto_config_domain_dns).
	UpdateNameservers *bool `json:"update_nameservers,omitempty"`

	// Deprecated: CustomRecords: custom records to synchronize.
	CustomRecords *[]*SyncDomainDNSRecordsRequestRecord `json:"custom_records,omitempty"`

	// AutoConfigDomainDNS: whether or not to synchronize each types of records.
	AutoConfigDomainDNS *AutoConfigDomainDNS `json:"auto_config_domain_dns,omitempty"`
}

// DNSRecords: dns records.
type DNSRecords struct {
	// Records: list of DNS records.
	Records []*DNSRecord `json:"records"`

	// NameServers: list of nameservers.
	NameServers []*Nameserver `json:"name_servers"`

	// Status: status of the records.
	// Default value: unknown_status
	Status DNSRecordsStatus `json:"status"`

	// Deprecated: DNSConfig: records dns auto configuration settings (deprecated, use auto_config_domain_dns).
	DNSConfig *[]DomainDNSAction `json:"dns_config,omitempty"`

	// AutoConfigDomainDNS: whether or not to synchronize each types of records.
	AutoConfigDomainDNS *AutoConfigDomainDNS `json:"auto_config_domain_dns"`
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

	// NewUser: (Optional) Username and password to create a user and link to the database.
	// Precisely one of NewUser, ExistingUsername must be set.
	NewUser *CreateDatabaseRequestUser `json:"new_user,omitempty"`

	// ExistingUsername: (Optional) Username to link an existing user to the database.
	// Precisely one of NewUser, ExistingUsername must be set.
	ExistingUsername *string `json:"existing_username,omitempty"`
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

// Domain: domain.
type Domain struct {
	// Name: name of the domain.
	Name string `json:"name"`

	// Status: current status of the domain.
	// Default value: unknown_status
	Status DomainStatus `json:"status"`

	// Owner: zone owner of the domain.
	// Default value: unknown_zone_owner
	Owner DomainZoneOwner `json:"owner"`

	// ZoneDomainName: main domain for this zone.
	ZoneDomainName string `json:"zone_domain_name"`

	// AvailableActions: a list of actions that can be performed on the domain.
	AvailableActions []DomainAction `json:"available_actions"`

	// Deprecated: AvailableDNSActions: a list of DNS-related actions that can be auto configured for the domain (deprecated, use auto_config_domain_dns instead).
	AvailableDNSActions *[]DomainDNSAction `json:"available_dns_actions,omitempty"`

	// AutoConfigDomainDNS: whether or not to synchronize each type of record.
	AutoConfigDomainDNS *AutoConfigDomainDNS `json:"auto_config_domain_dns"`
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

	// Deprecated: Domain: main domain associated with the Web Hosting plan (deprecated, use domain_info).
	Domain *string `json:"domain,omitempty"`

	// Offer: details of the Web Hosting plan offer and options.
	Offer *Offer `json:"offer"`

	// Platform: details of the hosting platform.
	Platform *Platform `json:"platform"`

	// Tags: list of tags associated with the Web Hosting plan.
	Tags []string `json:"tags"`

	// Deprecated: DNSStatus: DNS status of the Web Hosting plan (deprecated, use domain_info).
	// Default value: unknown_status
	DNSStatus *DNSRecordsStatus `json:"dns_status,omitempty"`

	// IPv4: current IPv4 address of the hosting.
	IPv4 net.IP `json:"ipv4"`

	// Protected: whether the hosting is protected or not.
	Protected bool `json:"protected"`

	// User: details of the hosting user.
	User *HostingUser `json:"user"`

	// Deprecated: DomainStatus: main domain status of the Web Hosting plan (deprecated, use domain_info).
	// Default value: unknown_status
	DomainStatus *DomainStatus `json:"domain_status,omitempty"`

	// Region: region where the Web Hosting plan is hosted.
	Region scw.Region `json:"region"`

	// DomainInfo: domain configuration block (subdomain, optional custom domain, and DNS settings).
	DomainInfo *HostingDomain `json:"domain_info"`
}

// HostingAPIAddCustomDomainRequest: hosting api add custom domain request.
type HostingAPIAddCustomDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: hosting ID to which the custom domain is attached to.
	HostingID string `json:"-"`

	// DomainName: the custom domain name to attach to the hosting.
	DomainName string `json:"domain_name"`
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

	// Subdomain: the name prefix to use as a free subdomain (for example, `mysite`) assigned to the Web Hosting plan. The full domain will be automatically created by adding it to the fixed base domain (e.g. `mysite.scw.site`). You do not need to include the base domain yourself.
	Subdomain *string `json:"subdomain,omitempty"`

	// OfferOptions: list of the Web Hosting plan options IDs with their quantities.
	OfferOptions []*OfferOptionRequest `json:"offer_options"`

	// Language: default language for the control panel interface.
	// Default value: unknown_language_code
	Language std.LanguageCode `json:"language"`

	// Deprecated: DomainConfiguration: indicates whether to update hosting domain name servers and DNS records for domains managed by Scaleway Elements (deprecated, use auto_config_domain_dns instead).
	DomainConfiguration *CreateHostingRequestDomainConfiguration `json:"domain_configuration,omitempty"`

	// SkipWelcomeEmail: indicates whether to skip a welcome email to the contact email containing hosting info.
	SkipWelcomeEmail *bool `json:"skip_welcome_email,omitempty"`

	// AutoConfigDomainDNS: indicates whether to update hosting domain name servers and DNS records for domains managed by Scaleway Elements (deprecated, use auto_update_* fields instead).
	AutoConfigDomainDNS *AutoConfigDomainDNS `json:"auto_config_domain_dns,omitempty"`
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

	// Subdomain: optional free subdomain linked to the Web Hosting plan.
	Subdomain *string `json:"-"`
}

// HostingAPIRemoveCustomDomainRequest: hosting api remove custom domain request.
type HostingAPIRemoveCustomDomainRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HostingID: hosting ID to which the custom domain is detached from.
	HostingID string `json:"-"`
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

// ListBackupItemsResponse: list backup items response.
type ListBackupItemsResponse struct {
	// TotalCount: total number of backup item groups.
	TotalCount uint64 `json:"total_count"`

	// Groups: list of backup item groups categorized by type.
	Groups []*BackupItemGroup `json:"groups"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBackupItemsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBackupItemsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListBackupItemsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Groups = append(r.Groups, results.Groups...)
	r.TotalCount += uint64(len(results.Groups))
	return uint64(len(results.Groups)), nil
}

// ListBackupsResponse: list backups response.
type ListBackupsResponse struct {
	// TotalCount: total number of available backups.
	TotalCount uint64 `json:"total_count"`

	// Backups: list of available backups.
	Backups []*Backup `json:"backups"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBackupsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBackupsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListBackupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Backups = append(r.Backups, results.Backups...)
	r.TotalCount += uint64(len(results.Backups))
	return uint64(len(results.Backups)), nil
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
func (r *ListControlPanelsResponse) UnsafeAppend(res any) (uint64, error) {
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
func (r *ListDatabaseUsersResponse) UnsafeAppend(res any) (uint64, error) {
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
func (r *ListDatabasesResponse) UnsafeAppend(res any) (uint64, error) {
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
func (r *ListFtpAccountsResponse) UnsafeAppend(res any) (uint64, error) {
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
func (r *ListHostingsResponse) UnsafeAppend(res any) (uint64, error) {
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
func (r *ListMailAccountsResponse) UnsafeAppend(res any) (uint64, error) {
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
func (r *ListOffersResponse) UnsafeAppend(res any) (uint64, error) {
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
func (r *ListWebsitesResponse) UnsafeAppend(res any) (uint64, error) {
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
	// Deprecated: OneTimePassword: new temporary password (deprecated, use password_b64 instead).
	OneTimePassword *string `json:"one_time_password,omitempty"`

	// OneTimePasswordB64: new temporary password, encoded in base64.
	OneTimePasswordB64 string `json:"one_time_password_b64"`
}

// ResourceSummary: resource summary.
type ResourceSummary struct {
	// DatabasesCount: total number of active databases in the Web Hosting plan.
	DatabasesCount uint32 `json:"databases_count"`

	// MailAccountsCount: total number of active email accounts in the Web Hosting plan.
	MailAccountsCount uint32 `json:"mail_accounts_count"`

	// FtpAccountsCount: total number of active FTP accounts in the Web Hosting plan.
	FtpAccountsCount uint32 `json:"ftp_accounts_count"`

	// WebsitesCount: total number of active domains in the Web Hosting plan.
	WebsitesCount uint32 `json:"websites_count"`
}

// RestoreBackupItemsResponse: restore backup items response.
type RestoreBackupItemsResponse struct{}

// RestoreBackupResponse: restore backup response.
type RestoreBackupResponse struct{}

// SearchDomainsResponse: search domains response.
type SearchDomainsResponse struct {
	// DomainsAvailable: list of domains availability.
	DomainsAvailable []*DomainAvailability `json:"domains_available"`
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

// This API allows you to list and restore backups for your cPanel and WordPress Web Hosting service.
type BackupAPI struct {
	client *scw.Client
}

// NewBackupAPI returns a BackupAPI object from a Scaleway client.
func NewBackupAPI(client *scw.Client) *BackupAPI {
	return &BackupAPI{
		client: client,
	}
}

func (s *BackupAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// ListBackups: List all available backups for a hosting account.
func (s *BackupAPI) ListBackups(req *BackupAPIListBackupsRequest, opts ...scw.RequestOption) (*ListBackupsResponse, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/backups",
		Query:  query,
	}

	var resp ListBackupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBackup: Get info about a backup specified by the backup ID.
func (s *BackupAPI) GetBackup(req *BackupAPIGetBackupRequest, opts ...scw.RequestOption) (*Backup, error) {
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

	if fmt.Sprint(req.BackupID) == "" {
		return nil, errors.New("field BackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/backups/" + fmt.Sprint(req.BackupID) + "",
	}

	var resp Backup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestoreBackup: Restore an entire backup to your hosting environment.
func (s *BackupAPI) RestoreBackup(req *BackupAPIRestoreBackupRequest, opts ...scw.RequestOption) (*RestoreBackupResponse, error) {
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

	if fmt.Sprint(req.BackupID) == "" {
		return nil, errors.New("field BackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/backups/" + fmt.Sprint(req.BackupID) + "/restore",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RestoreBackupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListBackupItems: List items within a specific backup, grouped by type.
func (s *BackupAPI) ListBackupItems(req *BackupAPIListBackupItemsRequest, opts ...scw.RequestOption) (*ListBackupItemsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "backup_id", req.BackupID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HostingID) == "" {
		return nil, errors.New("field HostingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/backup-items",
		Query:  query,
	}

	var resp ListBackupItemsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestoreBackupItems: Restore specific items from a backup (e.g., a database or mailbox).
func (s *BackupAPI) RestoreBackupItems(req *BackupAPIRestoreBackupItemsRequest, opts ...scw.RequestOption) (*RestoreBackupItemsResponse, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/restore-backup-items",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RestoreBackupItemsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
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

// This API allows you to manage your Web Hosting services.
type DnsAPI struct {
	client *scw.Client
}

// NewDnsAPI returns a DnsAPI object from a Scaleway client.
func NewDnsAPI(client *scw.Client) *DnsAPI {
	return &DnsAPI{
		client: client,
	}
}

func (s *DnsAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// GetDomainDNSRecords: Get the set of DNS records of a specified domain associated with a Web Hosting plan's domain.
func (s *DnsAPI) GetDomainDNSRecords(req *DNSAPIGetDomainDNSRecordsRequest, opts ...scw.RequestOption) (*DNSRecords, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.Domain) + "/dns-records",
	}

	var resp DNSRecords

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: CheckUserOwnsDomain: Check whether you own this domain or not.
func (s *DnsAPI) CheckUserOwnsDomain(req *DNSAPICheckUserOwnsDomainRequest, opts ...scw.RequestOption) (*CheckUserOwnsDomainResponse, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.Domain) + "/check-ownership",
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

// SyncDomainDNSRecords: Synchronize your DNS records on the Elements Console and on cPanel.
func (s *DnsAPI) SyncDomainDNSRecords(req *DNSAPISyncDomainDNSRecordsRequest, opts ...scw.RequestOption) (*DNSRecords, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.Domain) + "/sync-domain-dns-records",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DNSRecords

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SearchDomains: Search for available domains based on domain name.
func (s *DnsAPI) SearchDomains(req *DNSAPISearchDomainsRequest, opts ...scw.RequestOption) (*SearchDomainsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "domain_name", req.DomainName)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/search-domains",
		Query:  query,
	}

	var resp SearchDomainsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDomain: Retrieve detailed information about a specific domain, including its status, DNS configuration, and ownership.
func (s *DnsAPI) GetDomain(req *DNSAPIGetDomainRequest, opts ...scw.RequestOption) (*Domain, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DomainName) == "" {
		return nil, errors.New("field DomainName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/domains/" + fmt.Sprint(req.DomainName) + "",
		Query:  query,
	}

	var resp Domain

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
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
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
	parameter.AddToQuery(query, "subdomain", req.Subdomain)

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

// AddCustomDomain: Attach a custom domain to a webhosting.
func (s *HostingAPI) AddCustomDomain(req *HostingAPIAddCustomDomainRequest, opts ...scw.RequestOption) (*HostingSummary, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/add-custom-domain",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp HostingSummary

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveCustomDomain: Detach a custom domain from a webhosting.
func (s *HostingAPI) RemoveCustomDomain(req *HostingAPIRemoveCustomDomainRequest, opts ...scw.RequestOption) (*HostingSummary, error) {
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
		Path:   "/webhosting/v1/regions/" + fmt.Sprint(req.Region) + "/hostings/" + fmt.Sprint(req.HostingID) + "/remove-custom-domain",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp HostingSummary

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

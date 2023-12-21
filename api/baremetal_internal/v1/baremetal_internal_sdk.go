// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package baremetal_internal provides methods and message types of the baremetal_internal v1 API.
package baremetal_internal

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
	baremetal_v1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/baremetal/v1"
	trustandsafety_private_v1beta1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/trustandsafety_private/v1beta1"
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

type IPIPKind string

const (
	IPIPKindIPKindUnknown = IPIPKind("ip_kind_unknown")
	IPIPKindIPKindPublic  = IPIPKind("ip_kind_public")
	IPIPKindIPKindIpmi    = IPIPKind("ip_kind_ipmi")
)

func (enum IPIPKind) String() string {
	if enum == "" {
		// return default value if empty
		return "ip_kind_unknown"
	}
	return string(enum)
}

func (enum IPIPKind) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPIPKind) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPIPKind(IPIPKind(tmp).String())
	return nil
}

type IPPingStatus string

const (
	IPPingStatusPingStatusUnknown = IPPingStatus("ping_status_unknown")
	IPPingStatusPingStatusUp      = IPPingStatus("ping_status_up")
	IPPingStatusPingStatusDown    = IPPingStatus("ping_status_down")
)

func (enum IPPingStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "ping_status_unknown"
	}
	return string(enum)
}

func (enum IPPingStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPPingStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPPingStatus(IPPingStatus(tmp).String())
	return nil
}

type IPReverseStatus string

const (
	IPReverseStatusUnknown = IPReverseStatus("unknown")
	IPReverseStatusPending = IPReverseStatus("pending")
	IPReverseStatusActive  = IPReverseStatus("active")
	IPReverseStatusError   = IPReverseStatus("error")
)

func (enum IPReverseStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum IPReverseStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPReverseStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPReverseStatus(IPReverseStatus(tmp).String())
	return nil
}

type IPVersion string

const (
	IPVersionIPv4 = IPVersion("IPv4")
	IPVersionIPv6 = IPVersion("IPv6")
)

func (enum IPVersion) String() string {
	if enum == "" {
		// return default value if empty
		return "IPv4"
	}
	return string(enum)
}

func (enum IPVersion) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IPVersion) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IPVersion(IPVersion(tmp).String())
	return nil
}

type ListIPsRequestVersion string

const (
	ListIPsRequestVersionAll  = ListIPsRequestVersion("All")
	ListIPsRequestVersionIPv4 = ListIPsRequestVersion("IPv4")
	ListIPsRequestVersionIPv6 = ListIPsRequestVersion("IPv6")
)

func (enum ListIPsRequestVersion) String() string {
	if enum == "" {
		// return default value if empty
		return "All"
	}
	return string(enum)
}

func (enum ListIPsRequestVersion) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListIPsRequestVersion) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListIPsRequestVersion(ListIPsRequestVersion(tmp).String())
	return nil
}

type ListOSRequestStatus string

const (
	ListOSRequestStatusAll     = ListOSRequestStatus("all")
	ListOSRequestStatusEnable  = ListOSRequestStatus("enable")
	ListOSRequestStatusDisable = ListOSRequestStatus("disable")
)

func (enum ListOSRequestStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "all"
	}
	return string(enum)
}

func (enum ListOSRequestStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOSRequestStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOSRequestStatus(ListOSRequestStatus(tmp).String())
	return nil
}

type ListOffersRequestOrderBy string

const (
	ListOffersRequestOrderByNameAsc             = ListOffersRequestOrderBy("name_asc")
	ListOffersRequestOrderByNameDesc            = ListOffersRequestOrderBy("name_desc")
	ListOffersRequestOrderByCommercialRangeAsc  = ListOffersRequestOrderBy("commercial_range_asc")
	ListOffersRequestOrderByCommercialRangeDesc = ListOffersRequestOrderBy("commercial_range_desc")
	ListOffersRequestOrderByStockAsc            = ListOffersRequestOrderBy("stock_asc")
	ListOffersRequestOrderByStockDesc           = ListOffersRequestOrderBy("stock_desc")
)

func (enum ListOffersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
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

type ListOffersRequestStatus string

const (
	ListOffersRequestStatusAll     = ListOffersRequestStatus("all")
	ListOffersRequestStatusEnable  = ListOffersRequestStatus("enable")
	ListOffersRequestStatusDisable = ListOffersRequestStatus("disable")
)

func (enum ListOffersRequestStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "all"
	}
	return string(enum)
}

func (enum ListOffersRequestStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOffersRequestStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOffersRequestStatus(ListOffersRequestStatus(tmp).String())
	return nil
}

type ListServerEventsRequestOrderBy string

const (
	ListServerEventsRequestOrderByCreatedAtAsc  = ListServerEventsRequestOrderBy("created_at_asc")
	ListServerEventsRequestOrderByCreatedAtDesc = ListServerEventsRequestOrderBy("created_at_desc")
)

func (enum ListServerEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListServerEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServerEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServerEventsRequestOrderBy(ListServerEventsRequestOrderBy(tmp).String())
	return nil
}

type ListServersArchivedRequestOrderBy string

const (
	ListServersArchivedRequestOrderByCreatedAtAsc  = ListServersArchivedRequestOrderBy("created_at_asc")
	ListServersArchivedRequestOrderByCreatedAtDesc = ListServersArchivedRequestOrderBy("created_at_desc")
)

func (enum ListServersArchivedRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListServersArchivedRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServersArchivedRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServersArchivedRequestOrderBy(ListServersArchivedRequestOrderBy(tmp).String())
	return nil
}

type ListServersRequestOrderBy string

const (
	ListServersRequestOrderByCreatedAtAsc  = ListServersRequestOrderBy("created_at_asc")
	ListServersRequestOrderByCreatedAtDesc = ListServersRequestOrderBy("created_at_desc")
)

func (enum ListServersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListServersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServersRequestOrderBy(ListServersRequestOrderBy(tmp).String())
	return nil
}

type OfferSubscriptionPeriod string

const (
	OfferSubscriptionPeriodUnknownSubscriptionPeriod = OfferSubscriptionPeriod("unknown_subscription_period")
	OfferSubscriptionPeriodHourly                    = OfferSubscriptionPeriod("hourly")
	OfferSubscriptionPeriodMonthly                   = OfferSubscriptionPeriod("monthly")
)

func (enum OfferSubscriptionPeriod) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_subscription_period"
	}
	return string(enum)
}

func (enum OfferSubscriptionPeriod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OfferSubscriptionPeriod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OfferSubscriptionPeriod(OfferSubscriptionPeriod(tmp).String())
	return nil
}

type ServerBootType string

const (
	ServerBootTypeUnknownBootType = ServerBootType("unknown_boot_type")
	ServerBootTypeNormal          = ServerBootType("normal")
	ServerBootTypeRescue          = ServerBootType("rescue")
)

func (enum ServerBootType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_boot_type"
	}
	return string(enum)
}

func (enum ServerBootType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerBootType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerBootType(ServerBootType(tmp).String())
	return nil
}

type ServerInstallStatus string

const (
	ServerInstallStatusUnknown    = ServerInstallStatus("unknown")
	ServerInstallStatusToInstall  = ServerInstallStatus("to_install")
	ServerInstallStatusInstalling = ServerInstallStatus("installing")
	ServerInstallStatusInstalled  = ServerInstallStatus("installed")
	ServerInstallStatusError      = ServerInstallStatus("error")
)

func (enum ServerInstallStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ServerInstallStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerInstallStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerInstallStatus(ServerInstallStatus(tmp).String())
	return nil
}

type ServerOptionOptionStatus string

const (
	ServerOptionOptionStatusOptionStatusUnknown   = ServerOptionOptionStatus("option_status_unknown")
	ServerOptionOptionStatusOptionStatusEnable    = ServerOptionOptionStatus("option_status_enable")
	ServerOptionOptionStatusOptionStatusEnabling  = ServerOptionOptionStatus("option_status_enabling")
	ServerOptionOptionStatusOptionStatusDisabling = ServerOptionOptionStatus("option_status_disabling")
	ServerOptionOptionStatusOptionStatusError     = ServerOptionOptionStatus("option_status_error")
)

func (enum ServerOptionOptionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "option_status_unknown"
	}
	return string(enum)
}

func (enum ServerOptionOptionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerOptionOptionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerOptionOptionStatus(ServerOptionOptionStatus(tmp).String())
	return nil
}

type ServerPingStatus string

const (
	ServerPingStatusPingStatusUnknown = ServerPingStatus("ping_status_unknown")
	ServerPingStatusPingStatusUp      = ServerPingStatus("ping_status_up")
	ServerPingStatusPingStatusDown    = ServerPingStatus("ping_status_down")
)

func (enum ServerPingStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "ping_status_unknown"
	}
	return string(enum)
}

func (enum ServerPingStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerPingStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerPingStatus(ServerPingStatus(tmp).String())
	return nil
}

type ServerStatus string

const (
	ServerStatusUnknown    = ServerStatus("unknown")
	ServerStatusDelivering = ServerStatus("delivering")
	ServerStatusReady      = ServerStatus("ready")
	ServerStatusStopping   = ServerStatus("stopping")
	ServerStatusStopped    = ServerStatus("stopped")
	ServerStatusStarting   = ServerStatus("starting")
	ServerStatusError      = ServerStatus("error")
	ServerStatusDeleting   = ServerStatus("deleting")
	ServerStatusOutOfStock = ServerStatus("out_of_stock")
	ServerStatusOrdered    = ServerStatus("ordered")
	ServerStatusCleaning   = ServerStatus("cleaning")
	ServerStatusResetting  = ServerStatus("resetting")
)

func (enum ServerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ServerStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerStatus(ServerStatus(tmp).String())
	return nil
}

type SettingType string

const (
	SettingTypeUnknown = SettingType("unknown")
	SettingTypeSMTP    = SettingType("smtp")
)

func (enum SettingType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum SettingType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SettingType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SettingType(SettingType(tmp).String())
	return nil
}

// OSOSField: osos field.
type OSOSField struct {
	Editable bool `json:"editable"`

	Required bool `json:"required"`

	DefaultValue *string `json:"default_value"`
}

// CPU: cpu.
type CPU struct {
	Name string `json:"name"`

	CoreCount uint32 `json:"core_count"`

	ThreadCount uint32 `json:"thread_count"`

	Frequency uint32 `json:"frequency"`

	ID string `json:"id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	Benchmark string `json:"benchmark"`
}

// Disk: disk.
type Disk struct {
	Capacity uint64 `json:"capacity"`

	Type string `json:"type"`

	ID string `json:"id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// Memory: memory.
type Memory struct {
	Capacity uint64 `json:"capacity"`

	Type string `json:"type"`

	Frequency uint32 `json:"frequency"`

	IsEcc bool `json:"is_ecc"`

	ID string `json:"id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// OfferOptionOffer: offer option offer.
type OfferOptionOffer struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Price *scw.Money `json:"price"`

	Enabled bool `json:"enabled"`

	// SubscriptionPeriod: default value: unknown_subscription_period
	SubscriptionPeriod OfferSubscriptionPeriod `json:"subscription_period"`

	Manageable bool `json:"manageable"`

	OsID *string `json:"os_id"`
}

// PersistentMemory: persistent memory.
type PersistentMemory struct {
	Capacity uint64 `json:"capacity"`

	Type string `json:"type"`

	Frequency uint32 `json:"frequency"`

	ID string `json:"id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// RaidController: raid controller.
type RaidController struct {
	ID string `json:"id"`

	RaidLevel []string `json:"raid_level"`

	Model string `json:"model"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// HwProduct: hw product.
type HwProduct struct {
	Count int32 `json:"count"`

	ID string `json:"id"`
}

// IP: ip.
type IP struct {
	// ID: ID of the IP.
	ID string `json:"id"`

	// Address: address of the IP.
	Address net.IP `json:"address"`

	// Reverse: reverse IP value.
	Reverse string `json:"reverse"`

	// Version: version of IP (v4 or v6).
	// Default value: IPv4
	Version IPVersion `json:"version"`

	// ReverseStatus: status of the reverse.
	// Default value: unknown
	ReverseStatus IPReverseStatus `json:"reverse_status"`

	// ReverseStatusMessage: a message related to the reverse status, in case of an error for example.
	ReverseStatusMessage string `json:"reverse_status_message"`

	// Kind: the kind of ip (public or ipmi).
	// Default value: ip_kind_unknown
	Kind IPIPKind `json:"kind"`

	// PingStatus: IP status of ping.
	// Default value: ping_status_unknown
	PingStatus IPPingStatus `json:"ping_status"`
}

// ServerInstall: server install.
type ServerInstall struct {
	// OsID: ID of the OS to install.
	OsID string `json:"os_id"`

	// Hostname: hostname of the server.
	Hostname string `json:"hostname"`

	// SSHKeyIDs: array of SSH keys IDs to use.
	SSHKeyIDs []string `json:"ssh_key_ids"`

	// Status: status of the installation.
	// Default value: unknown
	Status ServerInstallStatus `json:"status"`

	// InternalError: error message.
	InternalError string `json:"internal_error"`

	// StartedAt: start date of the install.
	StartedAt *time.Time `json:"started_at"`

	// EndedAt: end date of the install.
	EndedAt *time.Time `json:"ended_at"`
}

// ServerOption: server option.
type ServerOption struct {
	// ID: ID of the option.
	ID string `json:"id"`

	// Name: name of the option.
	Name string `json:"name"`

	// Manageable: is false if the option could not be added or removed.
	Manageable bool `json:"manageable"`

	// Status: status of the option.
	// Default value: option_status_unknown
	Status ServerOptionOptionStatus `json:"status"`

	// ExpiresAt: auto expiration date for compatible options.
	ExpiresAt *time.Time `json:"expires_at"`
}

// CreateServerRequestInstall: create server request install.
type CreateServerRequestInstall struct {
	// OsID: ID of the OS to install on the server.
	OsID string `json:"os_id"`

	// Hostname: hostname of the server.
	Hostname string `json:"hostname"`

	// SSHKeyIDs: SSH key IDs authorized on the server.
	SSHKeyIDs []string `json:"ssh_key_ids"`

	// User: user used for the installation.
	User *string `json:"user"`

	// Password: password used for the installation.
	Password *string `json:"password"`

	// ServiceUser: user used for the service to install.
	ServiceUser *string `json:"service_user"`

	// ServicePassword: password used for the service to install.
	ServicePassword *string `json:"service_password"`

	// UserData: cloud-init file.
	UserData *scw.File `json:"user_data"`
}

// ExportUsersResponseUser: export users response user.
type ExportUsersResponseUser struct {
	OrganizationID string `json:"organization_id"`

	IP string `json:"ip"`
}

// OrganizationConsumption: organization consumption.
type OrganizationConsumption struct {
	OrganizationID string `json:"organization_id"`

	ServersCount uint32 `json:"servers_count"`
}

// OS: os.
type OS struct {
	// ID: ID of the OS.
	ID string `json:"id"`

	// Name: name of the OS.
	Name string `json:"name"`

	// Version: version of the OS.
	Version string `json:"version"`

	// ConsoleID: ID in Online Console.
	ConsoleID int32 `json:"console_id"`

	// Enable: state of OS.
	Enable bool `json:"enable"`

	// ReservedTo: array of organization_ids which have access to the OS, if null OS is available for every body.
	ReservedTo []string `json:"reserved_to"`

	// UpdatedAt: date of last update of the OS.
	UpdatedAt *time.Time `json:"updated_at"`

	// CreatedAt: creation date of the OS.
	CreatedAt *time.Time `json:"created_at"`

	// SSH: define the SSH requirements to install the OS.
	SSH *OSOSField `json:"ssh"`

	// User: define the username requirements to install the OS.
	User *OSOSField `json:"user"`

	// Password: define the password requirements to install the OS.
	Password *OSOSField `json:"password"`

	// ServiceUser: define the username requirements to install the service.
	ServiceUser *OSOSField `json:"service_user"`

	// ServicePassword: define the password requirements to install the service.
	ServicePassword *OSOSField `json:"service_password"`

	// LicenseRequired: license required (check server options for pricing details).
	LicenseRequired bool `json:"license_required"`
}

// Offer: offer.
type Offer struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Stock uint32 `json:"stock"`

	Bandwidth uint32 `json:"bandwidth"`

	CommercialRange string `json:"commercial_range"`

	Disks []*Disk `json:"disks"`

	Enable bool `json:"enable"`

	CPUs []*CPU `json:"cpus"`

	Memories []*Memory `json:"memories"`

	QuotaName string `json:"quota_name"`

	PricePerHour *scw.Money `json:"price_per_hour"`

	PricePerMonth *scw.Money `json:"price_per_month"`

	BillingOperationPath string `json:"billing_operation_path"`

	ReservedTo []string `json:"reserved_to"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	PersistentMemories []*PersistentMemory `json:"persistent_memories"`

	RaidControllers []*RaidController `json:"raid_controllers"`

	// SubscriptionPeriod: default value: unknown_subscription_period
	SubscriptionPeriod OfferSubscriptionPeriod `json:"subscription_period"`

	OperationPath string `json:"operation_path"`

	Fee *scw.Money `json:"fee"`

	Options []*OfferOptionOffer `json:"options"`

	PrivateBandwidth uint64 `json:"private_bandwidth"`

	SharedBandwidth bool `json:"shared_bandwidth"`

	Tags []string `json:"tags"`

	ProductIDs []string `json:"product_ids"`

	CompatibleOsIDs []string `json:"compatible_os_ids"`
}

// Partitioning: partitioning.
type Partitioning struct {
	ID string `json:"id"`

	ProductID string `json:"product_id"`

	Content string `json:"content"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// Product: product.
type Product struct {
	ID string `json:"id"`

	ConsoleProductID int32 `json:"console_product_id"`

	Datacenter int32 `json:"datacenter"`

	Support string `json:"support"`

	StockLowLevel int32 `json:"stock_low_level"`

	RescueImage string `json:"rescue_image"`

	Bandwidth uint32 `json:"bandwidth"`

	CPUsProduct []*HwProduct `json:"cpus_product"`

	MemoriesProduct []*HwProduct `json:"memories_product"`

	DisksProduct []*HwProduct `json:"disks_product"`

	IncompatibleOsIDs []string `json:"incompatible_os_ids"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	PersistentMemoriesProduct []*HwProduct `json:"persistent_memories_product"`

	RaidControllersProduct []*HwProduct `json:"raid_controllers_product"`
}

// ServerEvent: server event.
type ServerEvent struct {
	// ID: ID of the server for whom the action will be applied.
	ID string `json:"id"`

	// Action: the action that will be applied to the server.
	Action string `json:"action"`

	// Description: optional message describing the event status.
	Description string `json:"description"`

	// Public: bool indicating whether the event is visible to users.
	Public bool `json:"public"`

	// UpdatedAt: date of last modification of the action.
	UpdatedAt *time.Time `json:"updated_at"`

	// CreatedAt: date of creation of the action.
	CreatedAt *time.Time `json:"created_at"`

	// ResourceID: id of the external resource linked to this event.
	ResourceID *string `json:"resource_id"`

	// ResourceType: type of the external resource linked to this event.
	ResourceType *string `json:"resource_type"`

	// LogURL: path of the log file of the event.
	LogURL string `json:"log_url"`
}

// ServerArchived: server archived.
type ServerArchived struct {
	// ID: ID of the server.
	ID string `json:"id"`

	// OrganizationID: organization ID linked with the server.
	OrganizationID string `json:"organization_id"`

	// ProjectID: project ID linked with the server.
	ProjectID string `json:"project_id"`

	// Name: name of the server.
	Name string `json:"name"`

	// Description: description of the server.
	Description string `json:"description"`

	// UpdatedAt: date of last update of the server.
	UpdatedAt *time.Time `json:"updated_at"`

	// CreatedAt: creation date of the server.
	CreatedAt *time.Time `json:"created_at"`

	// DeletedAt: deleted date of the server.
	DeletedAt *time.Time `json:"deleted_at"`

	// OfferID: offer reference of the server.
	OfferID string `json:"offer_id"`

	// ConsoleID: ID in online console.
	ConsoleID *int32 `json:"console_id"`

	// ConsoleOrderID: order ID in online console.
	ConsoleOrderID *int32 `json:"console_order_id"`
}

// Server: server.
type Server struct {
	// ID: ID of the server.
	ID string `json:"id"`

	// OrganizationID: organization ID linked with the server.
	OrganizationID string `json:"organization_id"`

	// ProjectID: project ID linked with the server.
	ProjectID string `json:"project_id"`

	// Name: name of the server.
	Name string `json:"name"`

	// Description: description of the server.
	Description string `json:"description"`

	// UpdatedAt: date of last update of the server.
	UpdatedAt *time.Time `json:"updated_at"`

	// CreatedAt: creation date of the server.
	CreatedAt *time.Time `json:"created_at"`

	// Status: status of the server.
	// Default value: unknown
	Status ServerStatus `json:"status"`

	// OfferID: offer reference of the server.
	OfferID string `json:"offer_id"`

	// Install: object to describe the specifications of installation.
	Install *ServerInstall `json:"install"`

	// Tags: array of customs tags.
	Tags []string `json:"tags"`

	// IPs: array of IPs.
	IPs []*IP `json:"ips"`

	// Domain: the server domain.
	Domain string `json:"domain"`

	// BootType: boot type.
	// Default value: unknown_boot_type
	BootType ServerBootType `json:"boot_type"`

	// ConsoleID: ID in online console.
	ConsoleID *int32 `json:"console_id"`

	// ConsoleOrderID: order ID in online console.
	ConsoleOrderID *int32 `json:"console_order_id"`

	// Locked: boolean to know if server is locked.
	Locked bool `json:"locked"`

	// LockedReason: reason of lock.
	// Default value: unknown
	LockedReason trustandsafety_private_v1beta1.LockReason `json:"locked_reason"`

	// NetboxID: netbox ID of server.
	NetboxID *int32 `json:"netbox_id"`

	// Zone: the zone in which is the server.
	Zone scw.Zone `json:"zone"`

	// PingStatus: server status of ping.
	// Default value: ping_status_unknown
	PingStatus ServerPingStatus `json:"ping_status"`

	// HideLock: bool to hide the lock for user.
	HideLock bool `json:"hide_lock"`

	// Options: options enabled on server.
	Options []*ServerOption `json:"options"`

	// ServiceURL: the address of the installed service.
	ServiceURL string `json:"service_url"`

	// FeeOrderID: ID of the order to pay setup fees.
	FeeOrderID string `json:"fee_order_id"`

	// AssetID: ID of the asset.
	AssetID string `json:"asset_id"`
}

// AddOptionServerRequest: add option server request.
type AddOptionServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server.
	ServerID string `json:"-"`

	// OptionID: ID of the option to add.
	OptionID string `json:"-"`

	// ExpiresAt: auto expire the option after this date.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// CreateCPURequest: create cpu request.
type CreateCPURequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Name string `json:"name"`

	CoreCount uint32 `json:"core_count"`

	ThreadCount uint32 `json:"thread_count"`

	Frequency uint32 `json:"frequency"`
}

// CreateDiskRequest: create disk request.
type CreateDiskRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Type string `json:"type"`

	Capacity uint64 `json:"capacity"`
}

// CreateMemoryRequest: create memory request.
type CreateMemoryRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Type string `json:"type"`

	Capacity uint64 `json:"capacity"`

	Frequency uint32 `json:"frequency"`

	IsEcc bool `json:"is_ecc"`
}

// CreateOSRequest: create os request.
type CreateOSRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Name: name of the OS.
	Name string `json:"name"`

	// Version: version of the OS.
	Version string `json:"version"`

	// Enable: state of OS.
	Enable bool `json:"enable"`

	// ConsoleID: ID in Online Console.
	ConsoleID int32 `json:"console_id"`

	// ReservedTo: array of organization_ids which have access to the OS, if null OS is available for every body.
	ReservedTo []string `json:"reserved_to"`
}

// CreateOfferRequest: create offer request.
type CreateOfferRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ProductID string `json:"product_id"`

	Name string `json:"name"`

	CommercialRange string `json:"commercial_range"`

	Enable bool `json:"enable"`

	QuotaName string `json:"quota_name"`

	PricePerHour *scw.Money `json:"price_per_hour,omitempty"`

	PricePerMonth *scw.Money `json:"price_per_month,omitempty"`

	BillingOperationPath string `json:"billing_operation_path"`

	ReservedTo []string `json:"reserved_to"`
}

// CreatePartitioningRequest: create partitioning request.
type CreatePartitioningRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Content string `json:"content"`

	ProductID string `json:"product_id"`
}

// CreatePersistentMemoryRequest: create persistent memory request.
type CreatePersistentMemoryRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Type string `json:"type"`

	Capacity uint64 `json:"capacity"`

	Frequency uint32 `json:"frequency"`
}

// CreateProductRequest: create product request.
type CreateProductRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ConsoleProductID int32 `json:"console_product_id"`

	Datacenter int32 `json:"datacenter"`

	Support string `json:"support"`

	StockLowLevel int32 `json:"stock_low_level"`

	RescueImage string `json:"rescue_image"`

	Bandwidth int32 `json:"bandwidth"`

	CPUsProduct []*HwProduct `json:"cpus_product"`

	MemoriesProduct []*HwProduct `json:"memories_product"`

	DisksProduct []*HwProduct `json:"disks_product"`

	IncompatibleOsIDs []string `json:"incompatible_os_ids"`

	PersistentMemoriesProduct []*HwProduct `json:"persistent_memories_product"`

	RaidControllersProduct []*HwProduct `json:"raid_controllers_product"`
}

// CreateRaidControllerRequest: create raid controller request.
type CreateRaidControllerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	RaidLevel []string `json:"raid_level"`

	Model string `json:"model"`
}

// CreateServerRequest: create server request.
type CreateServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OfferID: offer ID of the new server.
	OfferID string `json:"offer_id"`

	// Deprecated: OrganizationID: organization ID with which the server will be created.
	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// ProjectID: project ID with which the server will be created.
	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// ProductID: if specified, will create a server using this Product ID.
	ProductID string `json:"product_id"`

	// Name: name of the server (≠hostname).
	Name string `json:"name"`

	// Description: description associated to the server, max 255 characters.
	Description string `json:"description"`

	// Tags: tags to associate to the server.
	Tags []string `json:"tags"`

	Install *CreateServerRequestInstall `json:"install,omitempty"`

	// OptionIDs: iDs of options to enable on server.
	OptionIDs []string `json:"option_ids"`

	// SkipFeeBill: skip billing fees when we create a server.
	SkipFeeBill bool `json:"skip_fee_bill"`

	// Comment: comment to describe why we used an internal endpoint to create a server for a customer.
	Comment string `json:"comment"`

	// ReplacedServerID: the id of the server to replace.
	ReplacedServerID string `json:"replaced_server_id"`
}

// CreateTaskRequest: create task request.
type CreateTaskRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Name: name of the task.
	Name string `json:"name"`

	// ResourceIDs: array of the resource IDs to execute the task.
	ResourceIDs []string `json:"resource_ids"`
}

// CreateTaskResponse: create task response.
type CreateTaskResponse struct {
	TaskName string `json:"task_name"`

	Launched map[string]string `json:"launched"`

	Failed map[string]string `json:"failed"`
}

// DeleteOptionServerRequest: delete option server request.
type DeleteOptionServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server.
	ServerID string `json:"-"`

	// OptionID: ID of the option to delete.
	OptionID string `json:"-"`
}

// DeleteSSHKeysRequest: delete ssh keys request.
type DeleteSSHKeysRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ProjectID string `json:"project_id"`
}

// DeleteSSHKeysResponse: delete ssh keys response.
type DeleteSSHKeysResponse struct {
}

// ExportUsersRequest: export users request.
type ExportUsersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Page *int32 `json:"page,omitempty"`

	PageSize *uint32 `json:"page_size,omitempty"`

	DediboxIDs []int32 `json:"dedibox_ids"`

	OrganizationIDs []string `json:"organization_ids"`

	ServerIDs []string `json:"server_ids"`
}

// ExportUsersResponse: export users response.
type ExportUsersResponse struct {
	TotalCount uint32 `json:"total_count"`

	Users []*ExportUsersResponseUser `json:"users"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ExportUsersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ExportUsersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ExportUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Users = append(r.Users, results.Users...)
	r.TotalCount += uint32(len(results.Users))
	return uint32(len(results.Users)), nil
}

// GetCPURequest: get cpu request.
type GetCPURequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	CPUID string `json:"-"`
}

// GetDiskRequest: get disk request.
type GetDiskRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	DiskID string `json:"-"`
}

// GetMemoryRequest: get memory request.
type GetMemoryRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	MemoryID string `json:"-"`
}

// GetOSRequest: get os request.
type GetOSRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OsID: ID of the researched OS.
	OsID string `json:"-"`
}

// GetOfferRequest: get offer request.
type GetOfferRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	OfferID string `json:"-"`
}

// GetOrganizationsConsumptionRequest: get organizations consumption request.
type GetOrganizationsConsumptionRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	FromDate *string `json:"-"`

	ToDate *string `json:"-"`

	Actif *bool `json:"-"`
}

// GetOrganizationsConsumptionResponse: get organizations consumption response.
type GetOrganizationsConsumptionResponse struct {
	TotalCount uint32 `json:"total_count"`

	OrganizationConsumptions []*OrganizationConsumption `json:"organization_consumptions"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *GetOrganizationsConsumptionResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *GetOrganizationsConsumptionResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*GetOrganizationsConsumptionResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.OrganizationConsumptions = append(r.OrganizationConsumptions, results.OrganizationConsumptions...)
	r.TotalCount += uint32(len(results.OrganizationConsumptions))
	return uint32(len(results.OrganizationConsumptions)), nil
}

// GetPartitioningRequest: get partitioning request.
type GetPartitioningRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PartitioningID string `json:"-"`
}

// GetPersistentMemoryRequest: get persistent memory request.
type GetPersistentMemoryRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PersistentMemoryID string `json:"-"`
}

// GetProductRequest: get product request.
type GetProductRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ProductID string `json:"-"`
}

// GetRaidControllerRequest: get raid controller request.
type GetRaidControllerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	RaidControllerID string `json:"-"`
}

// GetSLORequest: get slo request.
type GetSLORequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// GetSLOResponse: get slo response.
type GetSLOResponse struct {
	UpCount int32 `json:"up_count"`

	DownCount int32 `json:"down_count"`
}

// GetServerRequest: get server request.
type GetServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server.
	ServerID string `json:"-"`
}

// ListCPUsRequest: list cp us request.
type ListCPUsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Name *string `json:"-"`
}

// ListCPUsResponse: list cp us response.
type ListCPUsResponse struct {
	TotalCount uint32 `json:"total_count"`

	CPUs []*CPU `json:"cpus"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCPUsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCPUsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCPUsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.CPUs = append(r.CPUs, results.CPUs...)
	r.TotalCount += uint32(len(results.CPUs))
	return uint32(len(results.CPUs)), nil
}

// ListDisksRequest: list disks request.
type ListDisksRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Capacity *uint64 `json:"-"`

	Type *string `json:"-"`
}

// ListDisksResponse: list disks response.
type ListDisksResponse struct {
	TotalCount uint32 `json:"total_count"`

	Disks []*Disk `json:"disks"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDisksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDisksResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDisksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Disks = append(r.Disks, results.Disks...)
	r.TotalCount += uint32(len(results.Disks))
	return uint32(len(results.Disks)), nil
}

// ListIPsRequest: list i ps request.
type ListIPsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// Version: default value: All
	Version ListIPsRequestVersion `json:"-"`

	// Precisely one of OrganizationID, ProjectID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// Precisely one of OrganizationID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`
}

// ListIPsResponse: list i ps response.
type ListIPsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Addresses []string `json:"addresses"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Addresses = append(r.Addresses, results.Addresses...)
	r.TotalCount += uint32(len(results.Addresses))
	return uint32(len(results.Addresses)), nil
}

// ListMemoriesRequest: list memories request.
type ListMemoriesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Capacity *uint64 `json:"-"`

	Type *string `json:"-"`

	IsEcc *bool `json:"-"`
}

// ListMemoriesResponse: list memories response.
type ListMemoriesResponse struct {
	TotalCount uint32 `json:"total_count"`

	Memories []*Memory `json:"memories"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListMemoriesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListMemoriesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListMemoriesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Memories = append(r.Memories, results.Memories...)
	r.TotalCount += uint32(len(results.Memories))
	return uint32(len(results.Memories)), nil
}

// ListOSRequest: list os request.
type ListOSRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: number of OS per page.
	PageSize *uint32 `json:"-"`

	// Status: filter OS by status.
	// Default value: all
	Status ListOSRequestStatus `json:"-"`

	// Name: filter OS by name.
	Name *string `json:"-"`
}

// ListOSResponse: list os response.
type ListOSResponse struct {
	// TotalCount: total count of matching OS.
	TotalCount uint32 `json:"total_count"`

	// Os: oS that match filters.
	Os []*OS `json:"os"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOSResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOSResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListOSResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Os = append(r.Os, results.Os...)
	r.TotalCount += uint32(len(results.Os))
	return uint32(len(results.Os)), nil
}

// ListOffersRequest: list offers request.
type ListOffersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// Status: default value: all
	Status ListOffersRequestStatus `json:"-"`

	Name *string `json:"-"`

	QuotaName *string `json:"-"`

	// SubscriptionPeriod: default value: unknown_subscription_period
	SubscriptionPeriod OfferSubscriptionPeriod `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListOffersRequestOrderBy `json:"-"`

	CommercialRange *string `json:"-"`
}

// ListOffersResponse: list offers response.
type ListOffersResponse struct {
	TotalCount uint32 `json:"total_count"`

	Offers []*Offer `json:"offers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOffersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOffersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListOffersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Offers = append(r.Offers, results.Offers...)
	r.TotalCount += uint32(len(results.Offers))
	return uint32(len(results.Offers)), nil
}

// ListPartitioningsRequest: list partitionings request.
type ListPartitioningsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListPartitioningsResponse: list partitionings response.
type ListPartitioningsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Partitionings []*Partitioning `json:"partitionings"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPartitioningsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPartitioningsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPartitioningsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Partitionings = append(r.Partitionings, results.Partitionings...)
	r.TotalCount += uint32(len(results.Partitionings))
	return uint32(len(results.Partitionings)), nil
}

// ListPersistentMemoriesRequest: list persistent memories request.
type ListPersistentMemoriesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Capacity *uint64 `json:"-"`

	Type *string `json:"-"`
}

// ListPersistentMemoriesResponse: list persistent memories response.
type ListPersistentMemoriesResponse struct {
	TotalCount uint32 `json:"total_count"`

	PersistentMemories []*PersistentMemory `json:"persistent_memories"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPersistentMemoriesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPersistentMemoriesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPersistentMemoriesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PersistentMemories = append(r.PersistentMemories, results.PersistentMemories...)
	r.TotalCount += uint32(len(results.PersistentMemories))
	return uint32(len(results.PersistentMemories)), nil
}

// ListProductsRequest: list products request.
type ListProductsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListProductsResponse: list products response.
type ListProductsResponse struct {
	TotalCount uint32 `json:"total_count"`

	Products []*Product `json:"products"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListProductsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListProductsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListProductsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Products = append(r.Products, results.Products...)
	r.TotalCount += uint32(len(results.Products))
	return uint32(len(results.Products)), nil
}

// ListRaidControllersRequest: list raid controllers request.
type ListRaidControllersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListRaidControllersResponse: list raid controllers response.
type ListRaidControllersResponse struct {
	TotalCount uint32 `json:"total_count"`

	RaidControllers []*RaidController `json:"raid_controllers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRaidControllersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRaidControllersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListRaidControllersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.RaidControllers = append(r.RaidControllers, results.RaidControllers...)
	r.TotalCount += uint32(len(results.RaidControllers))
	return uint32(len(results.RaidControllers)), nil
}

// ListServerEventsRequest: list server events request.
type ListServerEventsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server events searched.
	ServerID string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: number of server events per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the server events.
	// Default value: created_at_asc
	OrderBy ListServerEventsRequestOrderBy `json:"-"`
}

// ListServerEventsResponse: list server events response.
type ListServerEventsResponse struct {
	// TotalCount: total count of matching events.
	TotalCount uint32 `json:"total_count"`

	// Events: server events that match filters.
	Events []*ServerEvent `json:"events"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServerEventsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServerEventsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListServerEventsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Events = append(r.Events, results.Events...)
	r.TotalCount += uint32(len(results.Events))
	return uint32(len(results.Events)), nil
}

// ListServersArchivedRequest: list servers archived request.
type ListServersArchivedRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: set the maximum list size.
	PageSize *uint32 `json:"-"`

	// OrderBy: order the response.
	// Default value: created_at_asc
	OrderBy ListServersArchivedRequestOrderBy `json:"-"`

	// Name: filter by name.
	Name *string `json:"-"`

	// OrganizationID: filter by organization.
	OrganizationID *string `json:"-"`

	// ProjectID: filter by project.
	ProjectID *string `json:"-"`
}

// ListServersArchivedResponse: list servers archived response.
type ListServersArchivedResponse struct {
	// TotalCount: total count of matching items (total_count≠page_size, it is the global count).
	TotalCount uint32 `json:"total_count"`

	// Servers: array of archived servers.
	Servers []*ServerArchived `json:"servers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServersArchivedResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServersArchivedResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListServersArchivedResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint32(len(results.Servers))
	return uint32(len(results.Servers)), nil
}

// ListServersRequest: list servers request.
type ListServersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: set the maximum list size.
	PageSize *uint32 `json:"-"`

	// OrderBy: order the response.
	// Default value: created_at_asc
	OrderBy ListServersRequestOrderBy `json:"-"`

	// Tags: filter by tags.
	Tags []string `json:"-"`

	// Status: filter by status.
	Status []string `json:"-"`

	// Name: filter by name.
	Name *string `json:"-"`

	// OrganizationID: filter by organization.
	OrganizationID *string `json:"-"`

	// ProjectID: filter by project.
	ProjectID *string `json:"-"`

	// OfferID: filter by offer.
	OfferID *string `json:"-"`

	// OptionID: filter by option ID.
	OptionID *string `json:"-"`
}

// ListServersResponse: list servers response.
type ListServersResponse struct {
	// TotalCount: total count of matching items (total_count≠page_size, it is the global count).
	TotalCount uint32 `json:"total_count"`

	// Servers: listing of servers.
	Servers []*Server `json:"servers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint32(len(results.Servers))
	return uint32(len(results.Servers)), nil
}

// ListTasksRequest: list tasks request.
type ListTasksRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// ListTasksResponse: list tasks response.
type ListTasksResponse struct {
	// TaskNames: array of task names available.
	TaskNames []string `json:"task_names"`
}

// MigrateServersWithOfferRequest: migrate servers with offer request.
type MigrateServersWithOfferRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	OldOfferID string `json:"old_offer_id"`

	NewOfferID string `json:"new_offer_id"`

	StopOldOfferAt *time.Time `json:"stop_old_offer_at,omitempty"`

	StartNewOfferAt *time.Time `json:"start_new_offer_at,omitempty"`
}

// MigrateServersWithOfferResponse: migrate servers with offer response.
type MigrateServersWithOfferResponse struct {
}

// SendTelemetriesRequest: send telemetries request.
type SendTelemetriesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Since: beginning date to send the telemetries.
	Since *time.Time `json:"since,omitempty"`
}

// SendTelemetriesResponse: send telemetries response.
type SendTelemetriesResponse struct {
}

// Setting: setting.
type Setting struct {
	// ID: ID of the setting.
	ID string `json:"id"`

	// ProjectID: ID of the project ID.
	ProjectID string `json:"project_id"`

	// Type: type of the setting.
	// Default value: unknown
	Type SettingType `json:"type"`

	// Enabled: the setting is enable or disable.
	Enabled bool `json:"enabled"`
}

// StartBMCAccessRequest: start bmc access request.
type StartBMCAccessRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server.
	ServerID string `json:"-"`

	// IP: the IP authorized to connect to the server.
	IP net.IP `json:"ip"`
}

// UpdateCPURequest: update cpu request.
type UpdateCPURequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	CPUID string `json:"-"`

	Name *string `json:"name,omitempty"`

	CoreCount *uint32 `json:"core_count,omitempty"`

	ThreadCount *uint32 `json:"thread_count,omitempty"`

	Frequency *uint32 `json:"frequency,omitempty"`
}

// UpdateDiskRequest: update disk request.
type UpdateDiskRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	DiskID string `json:"-"`

	Type *string `json:"type,omitempty"`

	Capacity *uint64 `json:"capacity,omitempty"`
}

// UpdateMemoryRequest: update memory request.
type UpdateMemoryRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	MemoryID string `json:"-"`

	Type *string `json:"type,omitempty"`

	Capacity *uint64 `json:"capacity,omitempty"`

	Frequency *uint32 `json:"frequency,omitempty"`

	IsEcc *bool `json:"is_ecc,omitempty"`
}

// UpdateOSRequest: update os request.
type UpdateOSRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OsID: ID of the OS to update.
	OsID string `json:"-"`

	// Name: name of the OS.
	Name *string `json:"name,omitempty"`

	// Version: version of the OS.
	Version *string `json:"version,omitempty"`

	// Enable: state of OS.
	Enable *bool `json:"enable,omitempty"`

	// ConsoleID: ID in Online Console.
	ConsoleID *int32 `json:"console_id,omitempty"`

	// ReservedTo: array of organization_ids which have access to the OS, if null OS is available for every body.
	ReservedTo *[]string `json:"reserved_to,omitempty"`
}

// UpdateOfferRequest: update offer request.
type UpdateOfferRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	OfferID string `json:"-"`

	ProductID *string `json:"product_id,omitempty"`

	Name *string `json:"name,omitempty"`

	CommercialRange *string `json:"commercial_range,omitempty"`

	Enable *bool `json:"enable,omitempty"`

	QuotaName *string `json:"quota_name,omitempty"`

	PricePerHour *scw.Money `json:"price_per_hour,omitempty"`

	PricePerMonth *scw.Money `json:"price_per_month,omitempty"`

	BillingOperationPath *string `json:"billing_operation_path,omitempty"`

	ReservedTo *[]string `json:"reserved_to,omitempty"`
}

// UpdatePartitioningRequest: update partitioning request.
type UpdatePartitioningRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PartitioningID string `json:"-"`

	ProductID *string `json:"product_id,omitempty"`

	Content *string `json:"content,omitempty"`
}

// UpdatePersistentMemoryRequest: update persistent memory request.
type UpdatePersistentMemoryRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PersistentMemoryID string `json:"-"`

	Type *string `json:"type,omitempty"`

	Capacity *uint64 `json:"capacity,omitempty"`

	Frequency *uint32 `json:"frequency,omitempty"`
}

// UpdateProductRequest: update product request.
type UpdateProductRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ProductID string `json:"-"`

	ConsoleProductID *int32 `json:"console_product_id,omitempty"`

	Datacenter *int32 `json:"datacenter,omitempty"`

	Support *string `json:"support,omitempty"`

	StockLowLevel *int32 `json:"stock_low_level,omitempty"`

	RescueImage *string `json:"rescue_image,omitempty"`

	Bandwidth *int32 `json:"bandwidth,omitempty"`

	CPUsProduct []*HwProduct `json:"cpus_product"`

	MemoriesProduct []*HwProduct `json:"memories_product"`

	DisksProduct []*HwProduct `json:"disks_product"`

	IncompatibleOsIDs *[]string `json:"incompatible_os_ids,omitempty"`

	PersistentMemoriesProduct []*HwProduct `json:"persistent_memories_product"`

	RaidControllersProduct []*HwProduct `json:"raid_controllers_product"`
}

// UpdateRaidControllerRequest: update raid controller request.
type UpdateRaidControllerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	RaidControllerID string `json:"-"`

	RaidLevel *[]string `json:"raid_level,omitempty"`

	Model *string `json:"model,omitempty"`
}

// UpdateSettingRequest: update setting request.
type UpdateSettingRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// SettingID: ID of the setting.
	SettingID string `json:"-"`

	// Enabled: enable/Disable the setting.
	Enabled *bool `json:"enabled,omitempty"`
}

// This internal API allows to manage all servers.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}
func (s *API) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar2}
}

// ListTasks: List all avaible tasks for admin intervention.
func (s *API) ListTasks(req *ListTasksRequest, opts ...scw.RequestOption) (*ListTasksResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/tasks",
	}

	var resp ListTasksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTask: Create task associated with the given name.
func (s *API) CreateTask(req *CreateTaskRequest, opts ...scw.RequestOption) (*CreateTaskResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/tasks",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateTaskResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendTelemetries: Send telemetries from the given date.
func (s *API) SendTelemetries(req *SendTelemetriesRequest, opts ...scw.RequestOption) (*SendTelemetriesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/telemetries",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SendTelemetriesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServers: List all servers.
func (s *API) ListServers(req *ListServersRequest, opts ...scw.RequestOption) (*ListServersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
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
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "offer_id", req.OfferID)
	parameter.AddToQuery(query, "option_id", req.OptionID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:  query,
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServersArchived: List all archived servers.
func (s *API) ListServersArchived(req *ListServersArchivedRequest, opts ...scw.RequestOption) (*ListServersArchivedResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers-archived",
		Query:  query,
	}

	var resp ListServersArchivedResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartBMCAccess: Unvailable - Start BMC (Baseboard Management Controller) access associated with the given ID.
func (s *API) StartBMCAccess(req *StartBMCAccessRequest, opts ...scw.RequestOption) (*baremetal_v1.BMCAccess, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/bmc-access",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp baremetal_v1.BMCAccess

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateServer: Create server associated with the offer.
func (s *API) CreateServer(req *CreateServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServer:
func (s *API) GetServer(req *GetServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServerEvents: List all events associated to the given server ID.
func (s *API) ListServerEvents(req *ListServerEventsRequest, opts ...scw.RequestOption) (*ListServerEventsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/events",
		Query:  query,
	}

	var resp ListServerEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListIPs:
func (s *API) ListIPs(req *ListIPsRequest, opts ...scw.RequestOption) (*ListIPsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "version", req.Version)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/ips",
		Query:  query,
	}

	var resp ListIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSSHKeys:
func (s *API) DeleteSSHKeys(req *DeleteSSHKeysRequest, opts ...scw.RequestOption) (*DeleteSSHKeysResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/ssh-keys",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DeleteSSHKeysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// MigrateServersWithOffer:
func (s *API) MigrateServersWithOffer(req *MigrateServersWithOfferRequest, opts ...scw.RequestOption) (*MigrateServersWithOfferResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/offers/migrate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp MigrateServersWithOfferResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOffers:
func (s *API) ListOffers(req *ListOffersRequest, opts ...scw.RequestOption) (*ListOffersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "quota_name", req.QuotaName)
	parameter.AddToQuery(query, "subscription_period", req.SubscriptionPeriod)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "commercial_range", req.CommercialRange)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/offers",
		Query:  query,
	}

	var resp ListOffersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOffer:
func (s *API) GetOffer(req *GetOfferRequest, opts ...scw.RequestOption) (*Offer, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OfferID) == "" {
		return nil, errors.New("field OfferID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/offers/" + fmt.Sprint(req.OfferID) + "",
	}

	var resp Offer

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOffer:
func (s *API) CreateOffer(req *CreateOfferRequest, opts ...scw.RequestOption) (*Offer, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/offers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Offer

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOffer:
func (s *API) UpdateOffer(req *UpdateOfferRequest, opts ...scw.RequestOption) (*Offer, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OfferID) == "" {
		return nil, errors.New("field OfferID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/offers/" + fmt.Sprint(req.OfferID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Offer

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOS: List all OS, could filter by name or status.
func (s *API) ListOS(req *ListOSRequest, opts ...scw.RequestOption) (*ListOSResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/os",
		Query:  query,
	}

	var resp ListOSResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOS: Create a new OS with the given params.
func (s *API) CreateOS(req *CreateOSRequest, opts ...scw.RequestOption) (*OS, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/os",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp OS

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOS: Return specific OS for the given ID.
func (s *API) GetOS(req *GetOSRequest, opts ...scw.RequestOption) (*OS, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OsID) == "" {
		return nil, errors.New("field OsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/os/" + fmt.Sprint(req.OsID) + "",
	}

	var resp OS

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOS: Update OS with the given params.
func (s *API) UpdateOS(req *UpdateOSRequest, opts ...scw.RequestOption) (*OS, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OsID) == "" {
		return nil, errors.New("field OsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/os/" + fmt.Sprint(req.OsID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp OS

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSLO:
func (s *API) GetSLO(req *GetSLORequest, opts ...scw.RequestOption) (*GetSLOResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/slo",
	}

	var resp GetSLOResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDisks:
func (s *API) ListDisks(req *ListDisksRequest, opts ...scw.RequestOption) (*ListDisksResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "capacity", req.Capacity)
	parameter.AddToQuery(query, "type", req.Type)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/disks",
		Query:  query,
	}

	var resp ListDisksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDisk:
func (s *API) CreateDisk(req *CreateDiskRequest, opts ...scw.RequestOption) (*Disk, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/disks",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Disk

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDisk:
func (s *API) GetDisk(req *GetDiskRequest, opts ...scw.RequestOption) (*Disk, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.DiskID) == "" {
		return nil, errors.New("field DiskID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/disks/" + fmt.Sprint(req.DiskID) + "",
	}

	var resp Disk

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDisk:
func (s *API) UpdateDisk(req *UpdateDiskRequest, opts ...scw.RequestOption) (*Disk, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.DiskID) == "" {
		return nil, errors.New("field DiskID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/disks/" + fmt.Sprint(req.DiskID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Disk

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListMemories:
func (s *API) ListMemories(req *ListMemoriesRequest, opts ...scw.RequestOption) (*ListMemoriesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "capacity", req.Capacity)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "is_ecc", req.IsEcc)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/memories",
		Query:  query,
	}

	var resp ListMemoriesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateMemory:
func (s *API) CreateMemory(req *CreateMemoryRequest, opts ...scw.RequestOption) (*Memory, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/memories",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Memory

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMemory:
func (s *API) GetMemory(req *GetMemoryRequest, opts ...scw.RequestOption) (*Memory, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.MemoryID) == "" {
		return nil, errors.New("field MemoryID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/memories/" + fmt.Sprint(req.MemoryID) + "",
	}

	var resp Memory

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateMemory:
func (s *API) UpdateMemory(req *UpdateMemoryRequest, opts ...scw.RequestOption) (*Memory, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.MemoryID) == "" {
		return nil, errors.New("field MemoryID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/memories/" + fmt.Sprint(req.MemoryID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Memory

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCPUs:
func (s *API) ListCPUs(req *ListCPUsRequest, opts ...scw.RequestOption) (*ListCPUsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/cpus",
		Query:  query,
	}

	var resp ListCPUsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCPU:
func (s *API) CreateCPU(req *CreateCPURequest, opts ...scw.RequestOption) (*CPU, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/cpus",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CPU

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCPU:
func (s *API) GetCPU(req *GetCPURequest, opts ...scw.RequestOption) (*CPU, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.CPUID) == "" {
		return nil, errors.New("field CPUID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/cpus/" + fmt.Sprint(req.CPUID) + "",
	}

	var resp CPU

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCPU:
func (s *API) UpdateCPU(req *UpdateCPURequest, opts ...scw.RequestOption) (*CPU, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.CPUID) == "" {
		return nil, errors.New("field CPUID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/cpus/" + fmt.Sprint(req.CPUID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CPU

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListProducts:
func (s *API) ListProducts(req *ListProductsRequest, opts ...scw.RequestOption) (*ListProductsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/products",
		Query:  query,
	}

	var resp ListProductsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateProduct:
func (s *API) CreateProduct(req *CreateProductRequest, opts ...scw.RequestOption) (*Product, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/products",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Product

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetProduct:
func (s *API) GetProduct(req *GetProductRequest, opts ...scw.RequestOption) (*Product, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ProductID) == "" {
		return nil, errors.New("field ProductID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/products/" + fmt.Sprint(req.ProductID) + "",
	}

	var resp Product

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateProduct:
func (s *API) UpdateProduct(req *UpdateProductRequest, opts ...scw.RequestOption) (*Product, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ProductID) == "" {
		return nil, errors.New("field ProductID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/products/" + fmt.Sprint(req.ProductID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Product

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPersistentMemories:
func (s *API) ListPersistentMemories(req *ListPersistentMemoriesRequest, opts ...scw.RequestOption) (*ListPersistentMemoriesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "capacity", req.Capacity)
	parameter.AddToQuery(query, "type", req.Type)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/persistent-memories",
		Query:  query,
	}

	var resp ListPersistentMemoriesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePersistentMemory:
func (s *API) CreatePersistentMemory(req *CreatePersistentMemoryRequest, opts ...scw.RequestOption) (*PersistentMemory, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/persistent-memories",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PersistentMemory

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPersistentMemory:
func (s *API) GetPersistentMemory(req *GetPersistentMemoryRequest, opts ...scw.RequestOption) (*PersistentMemory, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PersistentMemoryID) == "" {
		return nil, errors.New("field PersistentMemoryID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/persistent-memories/" + fmt.Sprint(req.PersistentMemoryID) + "",
	}

	var resp PersistentMemory

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePersistentMemory:
func (s *API) UpdatePersistentMemory(req *UpdatePersistentMemoryRequest, opts ...scw.RequestOption) (*PersistentMemory, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PersistentMemoryID) == "" {
		return nil, errors.New("field PersistentMemoryID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/persistent-memories/" + fmt.Sprint(req.PersistentMemoryID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PersistentMemory

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRaidControllers:
func (s *API) ListRaidControllers(req *ListRaidControllersRequest, opts ...scw.RequestOption) (*ListRaidControllersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/raid-controllers",
		Query:  query,
	}

	var resp ListRaidControllersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRaidController:
func (s *API) CreateRaidController(req *CreateRaidControllerRequest, opts ...scw.RequestOption) (*RaidController, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/raid-controllers",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RaidController

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRaidController:
func (s *API) GetRaidController(req *GetRaidControllerRequest, opts ...scw.RequestOption) (*RaidController, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.RaidControllerID) == "" {
		return nil, errors.New("field RaidControllerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/raid-controllers/" + fmt.Sprint(req.RaidControllerID) + "",
	}

	var resp RaidController

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateRaidController:
func (s *API) UpdateRaidController(req *UpdateRaidControllerRequest, opts ...scw.RequestOption) (*RaidController, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.RaidControllerID) == "" {
		return nil, errors.New("field RaidControllerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/raid-controllers/" + fmt.Sprint(req.RaidControllerID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RaidController

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPartitionings:
func (s *API) ListPartitionings(req *ListPartitioningsRequest, opts ...scw.RequestOption) (*ListPartitioningsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/partitionings",
		Query:  query,
	}

	var resp ListPartitioningsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePartitioning:
func (s *API) CreatePartitioning(req *CreatePartitioningRequest, opts ...scw.RequestOption) (*Partitioning, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/partitionings",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Partitioning

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPartitioning:
func (s *API) GetPartitioning(req *GetPartitioningRequest, opts ...scw.RequestOption) (*Partitioning, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PartitioningID) == "" {
		return nil, errors.New("field PartitioningID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/partitionings/" + fmt.Sprint(req.PartitioningID) + "",
	}

	var resp Partitioning

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePartitioning:
func (s *API) UpdatePartitioning(req *UpdatePartitioningRequest, opts ...scw.RequestOption) (*Partitioning, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PartitioningID) == "" {
		return nil, errors.New("field PartitioningID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/partitionings/" + fmt.Sprint(req.PartitioningID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Partitioning

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOrganizationsConsumption:
func (s *API) GetOrganizationsConsumption(req *GetOrganizationsConsumptionRequest, opts ...scw.RequestOption) (*GetOrganizationsConsumptionResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "from_date", req.FromDate)
	parameter.AddToQuery(query, "to_date", req.ToDate)
	parameter.AddToQuery(query, "actif", req.Actif)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/organizations/consumption",
		Query:  query,
	}

	var resp GetOrganizationsConsumptionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportUsers:
func (s *API) ExportUsers(req *ExportUsersRequest, opts ...scw.RequestOption) (*ExportUsersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/exports/users",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ExportUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddOptionServer: Add an option to a specific server.
func (s *API) AddOptionServer(req *AddOptionServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.OptionID) == "" {
		return nil, errors.New("field OptionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/options/" + fmt.Sprint(req.OptionID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteOptionServer: Delete an option from a specific server.
func (s *API) DeleteOptionServer(req *DeleteOptionServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.OptionID) == "" {
		return nil, errors.New("field OptionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/options/" + fmt.Sprint(req.OptionID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSetting: Update setting for a project ID (enable or disable).
func (s *API) UpdateSetting(req *UpdateSettingRequest, opts ...scw.RequestOption) (*Setting, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SettingID) == "" {
		return nil, errors.New("field SettingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/baremetal-internal/v1/zones/" + fmt.Sprint(req.Zone) + "/settings/" + fmt.Sprint(req.SettingID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Setting

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

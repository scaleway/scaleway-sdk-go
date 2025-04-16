// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package applesilicon provides methods and message types of the applesilicon v1alpha1 API.
package applesilicon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

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

type CommitmentType string

const (
	CommitmentTypeDuration24h    = CommitmentType("duration_24h")
	CommitmentTypeRenewedMonthly = CommitmentType("renewed_monthly")
	CommitmentTypeNone           = CommitmentType("none")
)

func (enum CommitmentType) String() string {
	if enum == "" {
		// return default value if empty
		return string(CommitmentTypeDuration24h)
	}
	return string(enum)
}

func (enum CommitmentType) Values() []CommitmentType {
	return []CommitmentType{
		"duration_24h",
		"renewed_monthly",
		"none",
	}
}

func (enum CommitmentType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CommitmentType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CommitmentType(CommitmentType(tmp).String())
	return nil
}

type ConnectivityDiagnosticActionType string

const (
	ConnectivityDiagnosticActionTypeRebootServer    = ConnectivityDiagnosticActionType("reboot_server")
	ConnectivityDiagnosticActionTypeReinstallServer = ConnectivityDiagnosticActionType("reinstall_server")
)

func (enum ConnectivityDiagnosticActionType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ConnectivityDiagnosticActionTypeRebootServer)
	}
	return string(enum)
}

func (enum ConnectivityDiagnosticActionType) Values() []ConnectivityDiagnosticActionType {
	return []ConnectivityDiagnosticActionType{
		"reboot_server",
		"reinstall_server",
	}
}

func (enum ConnectivityDiagnosticActionType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ConnectivityDiagnosticActionType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ConnectivityDiagnosticActionType(ConnectivityDiagnosticActionType(tmp).String())
	return nil
}

type ConnectivityDiagnosticDiagnosticStatus string

const (
	ConnectivityDiagnosticDiagnosticStatusUnknownStatus = ConnectivityDiagnosticDiagnosticStatus("unknown_status")
	ConnectivityDiagnosticDiagnosticStatusProcessing    = ConnectivityDiagnosticDiagnosticStatus("processing")
	ConnectivityDiagnosticDiagnosticStatusError         = ConnectivityDiagnosticDiagnosticStatus("error")
	ConnectivityDiagnosticDiagnosticStatusCompleted     = ConnectivityDiagnosticDiagnosticStatus("completed")
)

func (enum ConnectivityDiagnosticDiagnosticStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ConnectivityDiagnosticDiagnosticStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ConnectivityDiagnosticDiagnosticStatus) Values() []ConnectivityDiagnosticDiagnosticStatus {
	return []ConnectivityDiagnosticDiagnosticStatus{
		"unknown_status",
		"processing",
		"error",
		"completed",
	}
}

func (enum ConnectivityDiagnosticDiagnosticStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ConnectivityDiagnosticDiagnosticStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ConnectivityDiagnosticDiagnosticStatus(ConnectivityDiagnosticDiagnosticStatus(tmp).String())
	return nil
}

type ListServerPrivateNetworksRequestOrderBy string

const (
	ListServerPrivateNetworksRequestOrderByCreatedAtAsc  = ListServerPrivateNetworksRequestOrderBy("created_at_asc")
	ListServerPrivateNetworksRequestOrderByCreatedAtDesc = ListServerPrivateNetworksRequestOrderBy("created_at_desc")
	ListServerPrivateNetworksRequestOrderByUpdatedAtAsc  = ListServerPrivateNetworksRequestOrderBy("updated_at_asc")
	ListServerPrivateNetworksRequestOrderByUpdatedAtDesc = ListServerPrivateNetworksRequestOrderBy("updated_at_desc")
)

func (enum ListServerPrivateNetworksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListServerPrivateNetworksRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListServerPrivateNetworksRequestOrderBy) Values() []ListServerPrivateNetworksRequestOrderBy {
	return []ListServerPrivateNetworksRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
	}
}

func (enum ListServerPrivateNetworksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServerPrivateNetworksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServerPrivateNetworksRequestOrderBy(ListServerPrivateNetworksRequestOrderBy(tmp).String())
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
		return string(ListServersRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListServersRequestOrderBy) Values() []ListServersRequestOrderBy {
	return []ListServersRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
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

type ServerPrivateNetworkServerStatus string

const (
	ServerPrivateNetworkServerStatusUnknownStatus = ServerPrivateNetworkServerStatus("unknown_status")
	ServerPrivateNetworkServerStatusAttaching     = ServerPrivateNetworkServerStatus("attaching")
	ServerPrivateNetworkServerStatusAttached      = ServerPrivateNetworkServerStatus("attached")
	ServerPrivateNetworkServerStatusError         = ServerPrivateNetworkServerStatus("error")
	ServerPrivateNetworkServerStatusDetaching     = ServerPrivateNetworkServerStatus("detaching")
	ServerPrivateNetworkServerStatusLocked        = ServerPrivateNetworkServerStatus("locked")
)

func (enum ServerPrivateNetworkServerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerPrivateNetworkServerStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ServerPrivateNetworkServerStatus) Values() []ServerPrivateNetworkServerStatus {
	return []ServerPrivateNetworkServerStatus{
		"unknown_status",
		"attaching",
		"attached",
		"error",
		"detaching",
		"locked",
	}
}

func (enum ServerPrivateNetworkServerStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerPrivateNetworkServerStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerPrivateNetworkServerStatus(ServerPrivateNetworkServerStatus(tmp).String())
	return nil
}

type ServerPrivateNetworkStatus string

const (
	ServerPrivateNetworkStatusVpcUnknownStatus = ServerPrivateNetworkStatus("vpc_unknown_status")
	ServerPrivateNetworkStatusVpcEnabled       = ServerPrivateNetworkStatus("vpc_enabled")
	ServerPrivateNetworkStatusVpcUpdating      = ServerPrivateNetworkStatus("vpc_updating")
	ServerPrivateNetworkStatusVpcDisabled      = ServerPrivateNetworkStatus("vpc_disabled")
)

func (enum ServerPrivateNetworkStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerPrivateNetworkStatusVpcUnknownStatus)
	}
	return string(enum)
}

func (enum ServerPrivateNetworkStatus) Values() []ServerPrivateNetworkStatus {
	return []ServerPrivateNetworkStatus{
		"vpc_unknown_status",
		"vpc_enabled",
		"vpc_updating",
		"vpc_disabled",
	}
}

func (enum ServerPrivateNetworkStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerPrivateNetworkStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerPrivateNetworkStatus(ServerPrivateNetworkStatus(tmp).String())
	return nil
}

type ServerStatus string

const (
	ServerStatusUnknownStatus = ServerStatus("unknown_status")
	ServerStatusStarting      = ServerStatus("starting")
	ServerStatusReady         = ServerStatus("ready")
	ServerStatusError         = ServerStatus("error")
	ServerStatusRebooting     = ServerStatus("rebooting")
	ServerStatusUpdating      = ServerStatus("updating")
	ServerStatusLocking       = ServerStatus("locking")
	ServerStatusLocked        = ServerStatus("locked")
	ServerStatusUnlocking     = ServerStatus("unlocking")
	ServerStatusReinstalling  = ServerStatus("reinstalling")
	ServerStatusBusy          = ServerStatus("busy")
)

func (enum ServerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ServerStatus) Values() []ServerStatus {
	return []ServerStatus{
		"unknown_status",
		"starting",
		"ready",
		"error",
		"rebooting",
		"updating",
		"locking",
		"locked",
		"unlocking",
		"reinstalling",
		"busy",
	}
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

type ServerTypeStock string

const (
	ServerTypeStockUnknownStock = ServerTypeStock("unknown_stock")
	ServerTypeStockNoStock      = ServerTypeStock("no_stock")
	ServerTypeStockLowStock     = ServerTypeStock("low_stock")
	ServerTypeStockHighStock    = ServerTypeStock("high_stock")
)

func (enum ServerTypeStock) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerTypeStockUnknownStock)
	}
	return string(enum)
}

func (enum ServerTypeStock) Values() []ServerTypeStock {
	return []ServerTypeStock{
		"unknown_stock",
		"no_stock",
		"low_stock",
		"high_stock",
	}
}

func (enum ServerTypeStock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerTypeStock) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerTypeStock(ServerTypeStock(tmp).String())
	return nil
}

// OS: os.
type OS struct {
	// ID: unique ID of the OS.
	ID string `json:"id"`

	// Name: oS name.
	Name string `json:"name"`

	// Label: oS name as it should be displayed.
	Label string `json:"label"`

	// ImageURL: URL of the image.
	ImageURL string `json:"image_url"`

	// Family: the OS family to which this OS belongs, eg. 13 or 14.
	Family string `json:"family"`

	// IsBeta: describes if the OS is in beta.
	IsBeta bool `json:"is_beta"`

	// Version: the OS version number, eg. Sonoma has version number 14.3.
	Version string `json:"version"`

	// XcodeVersion: the current xcode version for this OS.
	XcodeVersion string `json:"xcode_version"`

	// CompatibleServerTypes: list of compatible server types.
	CompatibleServerTypes []string `json:"compatible_server_types"`
}

// ServerTypeCPU: server type cpu.
type ServerTypeCPU struct {
	Name string `json:"name"`

	CoreCount uint32 `json:"core_count"`

	Frequency uint64 `json:"frequency"`
}

// ServerTypeDisk: server type disk.
type ServerTypeDisk struct {
	Capacity scw.Size `json:"capacity"`

	Type string `json:"type"`
}

// ServerTypeGPU: server type gpu.
type ServerTypeGPU struct {
	Count uint64 `json:"count"`
}

// ServerTypeMemory: server type memory.
type ServerTypeMemory struct {
	Capacity scw.Size `json:"capacity"`

	Type string `json:"type"`
}

// ServerTypeNetwork: server type network.
type ServerTypeNetwork struct {
	PublicBandwidthBps uint64 `json:"public_bandwidth_bps"`
}

// Commitment: commitment.
type Commitment struct {
	// Type: default value: duration_24h
	Type CommitmentType `json:"type"`

	Cancelled bool `json:"cancelled"`
}

// ConnectivityDiagnosticServerHealth: connectivity diagnostic server health.
type ConnectivityDiagnosticServerHealth struct {
	LastCheckinDate *time.Time `json:"last_checkin_date"`

	IsServerAlive bool `json:"is_server_alive"`

	IsAgentAlive bool `json:"is_agent_alive"`

	IsMdmAlive bool `json:"is_mdm_alive"`

	IsSSHPortUp bool `json:"is_ssh_port_up"`

	IsVncPortUp bool `json:"is_vnc_port_up"`
}

// ServerPrivateNetwork: server private network.
type ServerPrivateNetwork struct {
	// ID: ID of the Server-to-Private Network mapping.
	ID string `json:"id"`

	// ProjectID: private Network Project ID.
	ProjectID string `json:"project_id"`

	// ServerID: apple silicon server ID.
	ServerID string `json:"server_id"`

	// PrivateNetworkID: private Network ID.
	PrivateNetworkID string `json:"private_network_id"`

	// Vlan: ID of the VLAN associated with the Private Network.
	Vlan *uint32 `json:"vlan"`

	// Status: configuration status of the Private Network.
	// Default value: unknown_status
	Status ServerPrivateNetworkServerStatus `json:"status"`

	// CreatedAt: private Network creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the Private Network was last modified.
	UpdatedAt *time.Time `json:"updated_at"`

	// IpamIPIDs: iPAM IP IDs of the server, if it has any.
	IpamIPIDs []string `json:"ipam_ip_ids"`
}

// ServerType: server type.
type ServerType struct {
	// CPU: CPU description.
	CPU *ServerTypeCPU `json:"cpu"`

	// Disk: size of the local disk of the server.
	Disk *ServerTypeDisk `json:"disk"`

	// Name: name of the type.
	Name string `json:"name"`

	// Memory: size of memory available.
	Memory *ServerTypeMemory `json:"memory"`

	// Stock: current stock.
	// Default value: unknown_stock
	Stock ServerTypeStock `json:"stock"`

	// MinimumLeaseDuration: minimum duration of the lease in seconds (example. 3.4s).
	MinimumLeaseDuration *scw.Duration `json:"minimum_lease_duration"`

	// Gpu: gPU description.
	Gpu *ServerTypeGPU `json:"gpu"`

	// Network: network description.
	Network *ServerTypeNetwork `json:"network"`

	// DefaultOs: the default OS for this server type.
	DefaultOs *OS `json:"default_os"`
}

// Server: server.
type Server struct {
	// ID: UUID of the server.
	ID string `json:"id"`

	// Type: type of the server.
	Type string `json:"type"`

	// Name: name of the server.
	Name string `json:"name"`

	// ProjectID: project this server is associated with.
	ProjectID string `json:"project_id"`

	// OrganizationID: organization this server is associated with.
	OrganizationID string `json:"organization_id"`

	// IP: iPv4 address of the server.
	IP net.IP `json:"ip"`

	// VncURL: vnc:// URL to access Apple Remote Desktop.
	VncURL string `json:"vnc_url"`

	// SSHUsername: SSH Username for remote shell.
	SSHUsername string `json:"ssh_username"`

	// SudoPassword: admin password required to execute commands.
	SudoPassword string `json:"sudo_password"`

	// VncPort: vNC port to use for remote desktop connection.
	VncPort uint32 `json:"vnc_port"`

	// Os: initially installed OS, this does not necessarily reflect the current OS version.
	Os *OS `json:"os"`

	// Status: current status of the server.
	// Default value: unknown_status
	Status ServerStatus `json:"status"`

	// CreatedAt: date on which the server was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date on which the server was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// DeletableAt: date from which the server can be deleted.
	DeletableAt *time.Time `json:"deletable_at"`

	// DeletionScheduled: set to true to mark the server for automatic deletion depending on `deletable_at` date. Set to false to cancel an existing deletion schedule. Leave unset otherwise.
	DeletionScheduled bool `json:"deletion_scheduled"`

	// Zone: zone of the server.
	Zone scw.Zone `json:"zone"`

	// Delivered: set to true once the server has completed its provisioning steps and is ready to use. Some OS configurations might require a reinstallation of the server before delivery depending on the available stock. A reinstallation after the initial delivery will not change this flag and can be tracked using the server status.
	Delivered bool `json:"delivered"`

	// VpcStatus: activation status of optional Private Network feature support for this server.
	// Default value: vpc_unknown_status
	VpcStatus ServerPrivateNetworkStatus `json:"vpc_status"`

	// Commitment: commitment scheme applied to this server.
	Commitment *Commitment `json:"commitment"`
}

// CommitmentTypeValue: commitment type value.
type CommitmentTypeValue struct {
	// CommitmentType: default value: duration_24h
	CommitmentType CommitmentType `json:"commitment_type"`
}

// ConnectivityDiagnostic: connectivity diagnostic.
type ConnectivityDiagnostic struct {
	ID string `json:"id"`

	// Status: default value: unknown_status
	Status ConnectivityDiagnosticDiagnosticStatus `json:"status"`

	IsHealthy bool `json:"is_healthy"`

	HealthDetails *ConnectivityDiagnosticServerHealth `json:"health_details"`

	SupportedActions []ConnectivityDiagnosticActionType `json:"supported_actions"`

	ErrorMessage string `json:"error_message"`
}

// CreateServerRequest: create server request.
type CreateServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Name: create a server with this given name.
	Name string `json:"name"`

	// ProjectID: create a server in the given project ID.
	ProjectID string `json:"project_id"`

	// Type: create a server of the given type.
	Type string `json:"type"`

	// OsID: create a server & install the given os_id, when no os_id provided the default OS for this server type is chosen. Requesting a non-default OS will induce an extended delivery time.
	OsID *string `json:"os_id,omitempty"`

	// EnableVpc: activate the Private Network feature for this server. This feature is configured through the Apple Silicon - Private Networks API.
	EnableVpc bool `json:"enable_vpc"`

	// CommitmentType: activate commitment for this server. If not specified, there is a 24h commitment due to Apple licensing (commitment_type `duration_24h`). It can be updated with the Update Server request. Available commitment depends on server type.
	// Default value: duration_24h
	CommitmentType CommitmentType `json:"commitment_type"`
}

// DeleteServerRequest: delete server request.
type DeleteServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: UUID of the server you want to delete.
	ServerID string `json:"-"`
}

// GetConnectivityDiagnosticRequest: get connectivity diagnostic request.
type GetConnectivityDiagnosticRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	DiagnosticID string `json:"-"`
}

// GetOSRequest: get os request.
type GetOSRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OsID: UUID of the OS you want to get.
	OsID string `json:"-"`
}

// GetServerRequest: get server request.
type GetServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: UUID of the server you want to get.
	ServerID string `json:"-"`
}

// GetServerTypeRequest: get server type request.
type GetServerTypeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerType: server type identifier.
	ServerType string `json:"-"`
}

// ListOSRequest: list os request.
type ListOSRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Page: positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`

	// ServerType: list of compatible server types.
	ServerType *string `json:"-"`

	// Name: filter OS by name (note that "11.1" will return "11.1.2" and "11.1" but not "12")).
	Name *string `json:"-"`
}

// ListOSResponse: list os response.
type ListOSResponse struct {
	// TotalCount: total number of OS.
	TotalCount uint32 `json:"total_count"`

	// Os: list of OS.
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

// ListServerPrivateNetworksResponse: list server private networks response.
type ListServerPrivateNetworksResponse struct {
	ServerPrivateNetworks []*ServerPrivateNetwork `json:"server_private_networks"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServerPrivateNetworksResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServerPrivateNetworksResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListServerPrivateNetworksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ServerPrivateNetworks = append(r.ServerPrivateNetworks, results.ServerPrivateNetworks...)
	r.TotalCount += uint64(len(results.ServerPrivateNetworks))
	return uint64(len(results.ServerPrivateNetworks)), nil
}

// ListServerTypesRequest: list server types request.
type ListServerTypesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// ListServerTypesResponse: list server types response.
type ListServerTypesResponse struct {
	// ServerTypes: available server types.
	ServerTypes []*ServerType `json:"server_types"`
}

// ListServersRequest: list servers request.
type ListServersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: sort order of the returned servers.
	// Default value: created_at_asc
	OrderBy ListServersRequestOrderBy `json:"-"`

	// ProjectID: only list servers of this project ID.
	ProjectID *string `json:"-"`

	// OrganizationID: only list servers of this Organization ID.
	OrganizationID *string `json:"-"`

	// Page: positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PageSize: positive integer lower or equal to 100 to select the number of items to return.
	PageSize *uint32 `json:"-"`
}

// ListServersResponse: list servers response.
type ListServersResponse struct {
	// TotalCount: total number of servers.
	TotalCount uint32 `json:"total_count"`

	// Servers: paginated returned servers.
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

// PrivateNetworkAPIAddServerPrivateNetworkRequest: private network api add server private network request.
type PrivateNetworkAPIAddServerPrivateNetworkRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server.
	ServerID string `json:"-"`

	// PrivateNetworkID: ID of the Private Network.
	PrivateNetworkID string `json:"private_network_id"`

	// IpamIPIDs: iPAM IDs of IPs to attach to the server.
	IpamIPIDs []string `json:"ipam_ip_ids"`
}

// PrivateNetworkAPIDeleteServerPrivateNetworkRequest: private network api delete server private network request.
type PrivateNetworkAPIDeleteServerPrivateNetworkRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server.
	ServerID string `json:"-"`

	// PrivateNetworkID: ID of the Private Network.
	PrivateNetworkID string `json:"-"`
}

// PrivateNetworkAPIGetServerPrivateNetworkRequest: private network api get server private network request.
type PrivateNetworkAPIGetServerPrivateNetworkRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	PrivateNetworkID string `json:"-"`
}

// PrivateNetworkAPIListServerPrivateNetworksRequest: private network api list server private networks request.
type PrivateNetworkAPIListServerPrivateNetworksRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: sort order for the returned Private Networks.
	// Default value: created_at_asc
	OrderBy ListServerPrivateNetworksRequestOrderBy `json:"-"`

	// Page: page number for the returned Private Networks.
	Page *int32 `json:"-"`

	// PageSize: maximum number of Private Networks per page.
	PageSize *uint32 `json:"-"`

	// ServerID: filter Private Networks by server ID.
	ServerID *string `json:"-"`

	// PrivateNetworkID: filter Private Networks by Private Network ID.
	PrivateNetworkID *string `json:"-"`

	// OrganizationID: filter Private Networks by Organization ID.
	OrganizationID *string `json:"-"`

	// ProjectID: filter Private Networks by Project ID.
	ProjectID *string `json:"-"`

	// IpamIPIDs: filter Private Networks by IPAM IP IDs.
	IpamIPIDs []string `json:"-"`
}

// PrivateNetworkAPISetServerPrivateNetworksRequest: private network api set server private networks request.
type PrivateNetworkAPISetServerPrivateNetworksRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server.
	ServerID string `json:"-"`

	// PerPrivateNetworkIpamIPIDs: object where the keys are the IDs of Private Networks and the values are arrays of IPAM IDs representing the IPs to assign to this Apple silicon server on the Private Network. If the array supplied for a Private Network is empty, the next available IP from the Private Network's CIDR block will automatically be used for attachment.
	PerPrivateNetworkIpamIPIDs map[string]*[]string `json:"per_private_network_ipam_ip_ids"`
}

// RebootServerRequest: reboot server request.
type RebootServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: UUID of the server you want to reboot.
	ServerID string `json:"-"`
}

// ReinstallServerRequest: reinstall server request.
type ReinstallServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: UUID of the server you want to reinstall.
	ServerID string `json:"-"`

	// OsID: reinstall the server with the target OS, when no os_id provided the default OS for the server type is used.
	OsID *string `json:"os_id,omitempty"`
}

// SetServerPrivateNetworksResponse: set server private networks response.
type SetServerPrivateNetworksResponse struct {
	ServerPrivateNetworks []*ServerPrivateNetwork `json:"server_private_networks"`
}

// StartConnectivityDiagnosticRequest: start connectivity diagnostic request.
type StartConnectivityDiagnosticRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"server_id"`
}

// StartConnectivityDiagnosticResponse: start connectivity diagnostic response.
type StartConnectivityDiagnosticResponse struct {
	DiagnosticID string `json:"diagnostic_id"`
}

// UpdateServerRequest: update server request.
type UpdateServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: UUID of the server you want to update.
	ServerID string `json:"-"`

	// Name: updated name for your server.
	Name *string `json:"name,omitempty"`

	// ScheduleDeletion: specify whether the server should be flagged for automatic deletion.
	ScheduleDeletion *bool `json:"schedule_deletion,omitempty"`

	// EnableVpc: activate or deactivate Private Network support for this server.
	EnableVpc *bool `json:"enable_vpc,omitempty"`

	// CommitmentType: change commitment. Use 'none' to automatically cancel a renewing commitment.
	CommitmentType *CommitmentTypeValue `json:"commitment_type,omitempty"`
}

// This API allows you to manage your Apple silicon machines.
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
	return []scw.Zone{scw.ZoneFrPar3}
}

// ListServerTypes: List all technical details about Apple silicon server types available in the specified zone. Since there is only one Availability Zone for Apple silicon servers, the targeted value is `fr-par-3`.
func (s *API) ListServerTypes(req *ListServerTypesRequest, opts ...scw.RequestOption) (*ListServerTypesResponse, error) {
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
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/server-types",
	}

	var resp ListServerTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServerType: Get technical details (CPU, disk size etc.) of a server type.
func (s *API) GetServerType(req *GetServerTypeRequest, opts ...scw.RequestOption) (*ServerType, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerType) == "" {
		return nil, errors.New("field ServerType cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/server-type/" + fmt.Sprint(req.ServerType) + "",
	}

	var resp ServerType

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateServer: Create a new server in the targeted zone, specifying its configuration including name and type.
func (s *API) CreateServer(req *CreateServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("as")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers",
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

// ListServers: List all servers in the specified zone. By default, returned servers in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:  query,
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOS: List all Operating Systems (OS). The response will include the total number of OS as well as their associated IDs, names and labels.
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
	parameter.AddToQuery(query, "server_type", req.ServerType)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/os",
		Query:  query,
	}

	var resp ListOSResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOS: Get an Operating System (OS).  The response will include the OS's unique ID as well as its name and label.
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
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/os/" + fmt.Sprint(req.OsID) + "",
	}

	var resp OS

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetServer: Retrieve information about an existing Apple silicon server, specified by its server ID. Its full details, including name, status and IP address, are returned in the response object.
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
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateServer: Update the parameters of an existing Apple silicon server, specified by its server ID.
func (s *API) UpdateServer(req *UpdateServerRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Method: "PATCH",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
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

// DeleteServer: Delete an existing Apple silicon server, specified by its server ID. Deleting a server is permanent, and cannot be undone. Note that the minimum allocation period for Apple silicon-as-a-service is 24 hours, meaning you cannot delete your server prior to that.
func (s *API) DeleteServer(req *DeleteServerRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// RebootServer: Reboot an existing Apple silicon server, specified by its server ID.
func (s *API) RebootServer(req *RebootServerRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/reboot",
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

// ReinstallServer: Reinstall an existing Apple silicon server (specified by its server ID) from a new image (OS). All the data on the disk is deleted and all configuration is reset to the defailt configuration values of the image (OS).
func (s *API) ReinstallServer(req *ReinstallServerRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/reinstall",
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

// StartConnectivityDiagnostic:
func (s *API) StartConnectivityDiagnostic(req *StartConnectivityDiagnosticRequest, opts ...scw.RequestOption) (*StartConnectivityDiagnosticResponse, error) {
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
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/connectivity-diagnostics",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp StartConnectivityDiagnosticResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetConnectivityDiagnostic:
func (s *API) GetConnectivityDiagnostic(req *GetConnectivityDiagnosticRequest, opts ...scw.RequestOption) (*ConnectivityDiagnostic, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.DiagnosticID) == "" {
		return nil, errors.New("field DiagnosticID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/connectivity-diagnostics/" + fmt.Sprint(req.DiagnosticID) + "",
	}

	var resp ConnectivityDiagnostic

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Apple silicon - Private Networks API.
type PrivateNetworkAPI struct {
	client *scw.Client
}

// NewPrivateNetworkAPI returns a PrivateNetworkAPI object from a Scaleway client.
func NewPrivateNetworkAPI(client *scw.Client) *PrivateNetworkAPI {
	return &PrivateNetworkAPI{
		client: client,
	}
}

func (s *PrivateNetworkAPI) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar3}
}

// GetServerPrivateNetwork:
func (s *PrivateNetworkAPI) GetServerPrivateNetwork(req *PrivateNetworkAPIGetServerPrivateNetworkRequest, opts ...scw.RequestOption) (*ServerPrivateNetwork, error) {
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

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	var resp ServerPrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddServerPrivateNetwork: Add an Apple silicon server to a Private Network.
func (s *PrivateNetworkAPI) AddServerPrivateNetwork(req *PrivateNetworkAPIAddServerPrivateNetworkRequest, opts ...scw.RequestOption) (*ServerPrivateNetwork, error) {
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
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private-networks",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ServerPrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetServerPrivateNetworks: Configure multiple Private Networks on an Apple silicon server.
func (s *PrivateNetworkAPI) SetServerPrivateNetworks(req *PrivateNetworkAPISetServerPrivateNetworksRequest, opts ...scw.RequestOption) (*SetServerPrivateNetworksResponse, error) {
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
		Method: "PUT",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private-networks",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetServerPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServerPrivateNetworks: List the Private Networks of an Apple silicon server.
func (s *PrivateNetworkAPI) ListServerPrivateNetworks(req *PrivateNetworkAPIListServerPrivateNetworksRequest, opts ...scw.RequestOption) (*ListServerPrivateNetworksResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "server_id", req.ServerID)
	parameter.AddToQuery(query, "private_network_id", req.PrivateNetworkID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "ipam_ip_ids", req.IpamIPIDs)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/server-private-networks",
		Query:  query,
	}

	var resp ListServerPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteServerPrivateNetwork: Delete a Private Network.
func (s *PrivateNetworkAPI) DeleteServerPrivateNetwork(req *PrivateNetworkAPIDeleteServerPrivateNetworkRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

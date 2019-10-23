// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package instance

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
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

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ scw.File
	_ = parameter.AddToQuery
	_ = namegenerator.GetRandomName
)

// API instance API
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type Arch string

const (
	// ArchX86_64 is [insert doc].
	ArchX86_64 = Arch("x86_64")
	// ArchArm is [insert doc].
	ArchArm = Arch("arm")
)

func (enum Arch) String() string {
	if enum == "" {
		// return default value if empty
		return "x86_64"
	}
	return string(enum)
}

func (enum Arch) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Arch) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Arch(Arch(tmp).String())
	return nil
}

type ImageState string

const (
	// ImageStateAvailable is [insert doc].
	ImageStateAvailable = ImageState("available")
	// ImageStateCreating is [insert doc].
	ImageStateCreating = ImageState("creating")
	// ImageStateError is [insert doc].
	ImageStateError = ImageState("error")
)

func (enum ImageState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum ImageState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ImageState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ImageState(ImageState(tmp).String())
	return nil
}

type PlacementGroupPolicyMode string

const (
	// PlacementGroupPolicyModeOptional is [insert doc].
	PlacementGroupPolicyModeOptional = PlacementGroupPolicyMode("optional")
	// PlacementGroupPolicyModeEnforced is [insert doc].
	PlacementGroupPolicyModeEnforced = PlacementGroupPolicyMode("enforced")
)

func (enum PlacementGroupPolicyMode) String() string {
	if enum == "" {
		// return default value if empty
		return "optional"
	}
	return string(enum)
}

func (enum PlacementGroupPolicyMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlacementGroupPolicyMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlacementGroupPolicyMode(PlacementGroupPolicyMode(tmp).String())
	return nil
}

type PlacementGroupPolicyType string

const (
	// PlacementGroupPolicyTypeMaxAvailability is [insert doc].
	PlacementGroupPolicyTypeMaxAvailability = PlacementGroupPolicyType("max_availability")
	// PlacementGroupPolicyTypeLowLatency is [insert doc].
	PlacementGroupPolicyTypeLowLatency = PlacementGroupPolicyType("low_latency")
)

func (enum PlacementGroupPolicyType) String() string {
	if enum == "" {
		// return default value if empty
		return "max_availability"
	}
	return string(enum)
}

func (enum PlacementGroupPolicyType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlacementGroupPolicyType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlacementGroupPolicyType(PlacementGroupPolicyType(tmp).String())
	return nil
}

type SecurityGroupPolicy string

const (
	// SecurityGroupPolicyAccept is [insert doc].
	SecurityGroupPolicyAccept = SecurityGroupPolicy("accept")
	// SecurityGroupPolicyDrop is [insert doc].
	SecurityGroupPolicyDrop = SecurityGroupPolicy("drop")
)

func (enum SecurityGroupPolicy) String() string {
	if enum == "" {
		// return default value if empty
		return "accept"
	}
	return string(enum)
}

func (enum SecurityGroupPolicy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupPolicy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupPolicy(SecurityGroupPolicy(tmp).String())
	return nil
}

type SecurityGroupRuleAction string

const (
	// SecurityGroupRuleActionAccept is [insert doc].
	SecurityGroupRuleActionAccept = SecurityGroupRuleAction("accept")
	// SecurityGroupRuleActionDrop is [insert doc].
	SecurityGroupRuleActionDrop = SecurityGroupRuleAction("drop")
)

func (enum SecurityGroupRuleAction) String() string {
	if enum == "" {
		// return default value if empty
		return "accept"
	}
	return string(enum)
}

func (enum SecurityGroupRuleAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupRuleAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupRuleAction(SecurityGroupRuleAction(tmp).String())
	return nil
}

type SecurityGroupRuleDirection string

const (
	// SecurityGroupRuleDirectionInbound is [insert doc].
	SecurityGroupRuleDirectionInbound = SecurityGroupRuleDirection("inbound")
	// SecurityGroupRuleDirectionOutbound is [insert doc].
	SecurityGroupRuleDirectionOutbound = SecurityGroupRuleDirection("outbound")
)

func (enum SecurityGroupRuleDirection) String() string {
	if enum == "" {
		// return default value if empty
		return "inbound"
	}
	return string(enum)
}

func (enum SecurityGroupRuleDirection) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupRuleDirection) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupRuleDirection(SecurityGroupRuleDirection(tmp).String())
	return nil
}

type SecurityGroupRuleProtocol string

const (
	// SecurityGroupRuleProtocolTCP is [insert doc].
	SecurityGroupRuleProtocolTCP = SecurityGroupRuleProtocol("TCP")
	// SecurityGroupRuleProtocolUDP is [insert doc].
	SecurityGroupRuleProtocolUDP = SecurityGroupRuleProtocol("UDP")
	// SecurityGroupRuleProtocolICMP is [insert doc].
	SecurityGroupRuleProtocolICMP = SecurityGroupRuleProtocol("ICMP")
	// SecurityGroupRuleProtocolANY is [insert doc].
	SecurityGroupRuleProtocolANY = SecurityGroupRuleProtocol("ANY")
)

func (enum SecurityGroupRuleProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "TCP"
	}
	return string(enum)
}

func (enum SecurityGroupRuleProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupRuleProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupRuleProtocol(SecurityGroupRuleProtocol(tmp).String())
	return nil
}

type ServerAction string

const (
	// ServerActionPoweron is [insert doc].
	ServerActionPoweron = ServerAction("poweron")
	// ServerActionBackup is [insert doc].
	ServerActionBackup = ServerAction("backup")
	// ServerActionStopInPlace is [insert doc].
	ServerActionStopInPlace = ServerAction("stop_in_place")
	// ServerActionPoweroff is [insert doc].
	ServerActionPoweroff = ServerAction("poweroff")
	// ServerActionTerminate is [insert doc].
	ServerActionTerminate = ServerAction("terminate")
	// ServerActionReboot is [insert doc].
	ServerActionReboot = ServerAction("reboot")
)

func (enum ServerAction) String() string {
	if enum == "" {
		// return default value if empty
		return "poweron"
	}
	return string(enum)
}

func (enum ServerAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerAction(ServerAction(tmp).String())
	return nil
}

type ServerBootType string

const (
	// ServerBootTypeLocal is [insert doc].
	ServerBootTypeLocal = ServerBootType("local")
)

func (enum ServerBootType) String() string {
	if enum == "" {
		// return default value if empty
		return "local"
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

type ServerState string

const (
	// ServerStateRunning is [insert doc].
	ServerStateRunning = ServerState("running")
	// ServerStateStopped is [insert doc].
	ServerStateStopped = ServerState("stopped")
	// ServerStateStoppedInPlace is [insert doc].
	ServerStateStoppedInPlace = ServerState("stopped in place")
	// ServerStateStarting is [insert doc].
	ServerStateStarting = ServerState("starting")
	// ServerStateStopping is [insert doc].
	ServerStateStopping = ServerState("stopping")
	// ServerStateLocked is [insert doc].
	ServerStateLocked = ServerState("locked")
)

func (enum ServerState) String() string {
	if enum == "" {
		// return default value if empty
		return "running"
	}
	return string(enum)
}

func (enum ServerState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerState(ServerState(tmp).String())
	return nil
}

type ServerTypesAvailability string

const (
	// ServerTypesAvailabilityAvailable is [insert doc].
	ServerTypesAvailabilityAvailable = ServerTypesAvailability("available")
	// ServerTypesAvailabilityScarce is [insert doc].
	ServerTypesAvailabilityScarce = ServerTypesAvailability("scarce")
	// ServerTypesAvailabilityShortage is [insert doc].
	ServerTypesAvailabilityShortage = ServerTypesAvailability("shortage")
)

func (enum ServerTypesAvailability) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum ServerTypesAvailability) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerTypesAvailability) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerTypesAvailability(ServerTypesAvailability(tmp).String())
	return nil
}

type SnapshotState string

const (
	// SnapshotStateAvailable is [insert doc].
	SnapshotStateAvailable = SnapshotState("available")
	// SnapshotStateSnapshotting is [insert doc].
	SnapshotStateSnapshotting = SnapshotState("snapshotting")
	// SnapshotStateError is [insert doc].
	SnapshotStateError = SnapshotState("error")
)

func (enum SnapshotState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum SnapshotState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SnapshotState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SnapshotState(SnapshotState(tmp).String())
	return nil
}

type TaskStatus string

const (
	// TaskStatusPending is [insert doc].
	TaskStatusPending = TaskStatus("pending")
	// TaskStatusStarted is [insert doc].
	TaskStatusStarted = TaskStatus("started")
	// TaskStatusSuccess is [insert doc].
	TaskStatusSuccess = TaskStatus("success")
	// TaskStatusFailure is [insert doc].
	TaskStatusFailure = TaskStatus("failure")
	// TaskStatusRetry is [insert doc].
	TaskStatusRetry = TaskStatus("retry")
)

func (enum TaskStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "pending"
	}
	return string(enum)
}

func (enum TaskStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TaskStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TaskStatus(TaskStatus(tmp).String())
	return nil
}

type VolumeState string

const (
	// VolumeStateAvailable is [insert doc].
	VolumeStateAvailable = VolumeState("available")
	// VolumeStateSnapshotting is [insert doc].
	VolumeStateSnapshotting = VolumeState("snapshotting")
	// VolumeStateError is [insert doc].
	VolumeStateError = VolumeState("error")
)

func (enum VolumeState) String() string {
	if enum == "" {
		// return default value if empty
		return "available"
	}
	return string(enum)
}

func (enum VolumeState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeState(VolumeState(tmp).String())
	return nil
}

type VolumeType string

const (
	// VolumeTypeLSSD is [insert doc].
	VolumeTypeLSSD = VolumeType("l_ssd")
	// VolumeTypeBSSD is [insert doc].
	VolumeTypeBSSD = VolumeType("b_ssd")
)

func (enum VolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return "l_ssd"
	}
	return string(enum)
}

func (enum VolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VolumeType(VolumeType(tmp).String())
	return nil
}

// Bootscript bootscript
type Bootscript struct {
	// Bootcmdargs display the bootscript parameters
	Bootcmdargs string `json:"bootcmdargs"`
	// Default dispmay if the bootscript is the default bootscript if no other boot option is configured
	Default bool `json:"default"`
	// Dtb provide information regarding a Device Tree Binary (dtb) for use with C1 servers
	Dtb string `json:"dtb"`
	// ID display the bootscripts ID
	ID string `json:"id"`
	// Initrd display the initrd (initial ramdisk) configuration
	Initrd string `json:"initrd"`
	// Kernel display the server kernel version
	Kernel string `json:"kernel"`
	// Organization display the bootscripts organization
	Organization string `json:"organization"`
	// Public provide information if the bootscript is public
	Public bool `json:"public"`
	// Title display the bootscripts title
	Title string `json:"title"`
	// Arch display the bootscripts arch
	//
	// Default value: x86_64
	Arch Arch `json:"arch"`
}

type CreateIPResponse struct {
	IP *IP `json:"ip"`

	Location string `json:"Location"`
}

type CreateImageResponse struct {
	Image *Image `json:"image"`

	Location string `json:"Location"`
}

type CreatePlacementGroupResponse struct {
	PlacementGroup *PlacementGroup `json:"placement_group"`
}

type CreateSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group"`
}

type CreateSecurityGroupRuleResponse struct {
	Rule *SecurityGroupRule `json:"rule"`
}

type CreateServerResponse struct {
	Server *Server `json:"server"`
}

type CreateSnapshotResponse struct {
	Snapshot *Snapshot `json:"snapshot"`
}

type CreateVolumeResponse struct {
	Volume *Volume `json:"volume"`

	Location string `json:"Location"`
}

type Dashboard struct {
	VolumesCount uint32 `json:"volumes_count"`

	RunningServersCount uint32 `json:"running_servers_count"`

	ServersByTypes map[string]uint32 `json:"servers_by_types"`

	ImagesCount uint32 `json:"images_count"`

	SnapshotsCount uint32 `json:"snapshots_count"`

	ServersCount uint32 `json:"servers_count"`

	IPsCount uint32 `json:"ips_count"`

	SecurityGroupsCount uint32 `json:"security_groups_count"`

	IPsUnused uint32 `json:"ips_unused"`
}

type GetBootscriptResponse struct {
	Bootscript *Bootscript `json:"bootscript"`
}

type GetDashboardResponse struct {
	Dashboard *Dashboard `json:"dashboard"`
}

type GetIPResponse struct {
	IP *IP `json:"ip"`
}

type GetImageResponse struct {
	Image *Image `json:"image"`
}

type GetPlacementGroupResponse struct {
	PlacementGroup *PlacementGroup `json:"placement_group"`
}

type GetPlacementGroupServersResponse struct {
	Servers []*PlacementGroupServer `json:"servers"`
}

type GetSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group"`
}

type GetSecurityGroupRuleResponse struct {
	Rule *SecurityGroupRule `json:"rule"`
}

type GetServerResponse struct {
	Server *Server `json:"server"`
}

type GetServerTypesAvailabilityResponse struct {
	Servers map[string]ServerTypesAvailability `json:"servers"`
}

type GetSnapshotResponse struct {
	Snapshot *Snapshot `json:"snapshot"`
}

type GetVolumeResponse struct {
	Volume *Volume `json:"volume"`
}

type IP struct {
	ID string `json:"id"`

	Address net.IP `json:"address"`

	Reverse *string `json:"reverse"`

	Server *ServerSummary `json:"server"`

	Organization string `json:"organization"`
}

type Image struct {
	ID string `json:"id"`

	Name string `json:"name"`
	// Arch
	//
	// Default value: x86_64
	Arch Arch `json:"arch"`

	CreationDate time.Time `json:"creation_date"`

	ModificationDate time.Time `json:"modification_date"`

	DefaultBootscript *Bootscript `json:"default_bootscript"`

	ExtraVolumes map[string]*Volume `json:"extra_volumes"`

	FromServer string `json:"from_server"`

	Organization string `json:"organization"`

	Public bool `json:"public"`

	RootVolume *VolumeSummary `json:"root_volume"`
	// State
	//
	// Default value: available
	State ImageState `json:"state"`
}

type ListBootscriptsResponse struct {
	Bootscripts []*Bootscript `json:"bootscripts"`

	TotalCount uint32 `json:"total_count"`
}

type ListIPsResponse struct {
	IPs []*IP `json:"ips"`

	TotalCount uint32 `json:"total_count"`
}

type ListImagesResponse struct {
	Images []*Image `json:"images"`

	TotalCount uint32 `json:"total_count"`
}

type ListPlacementGroupsResponse struct {
	PlacementGroups []*PlacementGroup `json:"placement_groups"`

	TotalCount uint32 `json:"total_count"`
}

type ListSecurityGroupRulesResponse struct {
	Rules []*SecurityGroupRule `json:"rules"`

	TotalCount uint32 `json:"total_count"`
}

type ListSecurityGroupsResponse struct {
	SecurityGroups []*SecurityGroup `json:"security_groups"`

	TotalCount uint32 `json:"total_count"`
}

type ListServerActionsResponse struct {
	Actions []ServerAction `json:"actions"`
}

type ListServerUserDataResponse struct {
	UserData []string `json:"user_data"`
}

type ListServersResponse struct {
	Servers []*Server `json:"servers"`

	TotalCount uint32 `json:"total_count"`
}

type ListServersTypesResponse struct {
	Servers map[string]*ServerType `json:"servers"`

	TotalCount uint32 `json:"total_count"`
}

type ListSnapshotsResponse struct {
	Snapshots []*Snapshot `json:"snapshots"`

	TotalCount uint32 `json:"total_count"`
}

type ListVolumesResponse struct {
	Volumes []*Volume `json:"volumes"`

	TotalCount uint32 `json:"total_count"`
}

type NullableStringValue struct {
	Null bool `json:"null,omitempty"`

	Value string `json:"value,omitempty"`
}

// PlacementGroup placement group
type PlacementGroup struct {
	// ID display placement group unique ID
	ID string `json:"id"`
	// Name display placement group name
	Name string `json:"name"`
	// Organization display placement group organization
	Organization string `json:"organization"`
	// PolicyMode select the failling mode when the placement cannot be  respected, either optional or enforced
	//
	// Default value: optional
	PolicyMode PlacementGroupPolicyMode `json:"policy_mode"`
	// PolicyType select the behavior of the placement group, either low_latency (group) or max_availability (spread)
	//
	// Default value: max_availability
	PolicyType PlacementGroupPolicyType `json:"policy_type"`
	// PolicyRespected returns true if the policy is respected, false otherwise
	PolicyRespected bool `json:"policy_respected"`
}

type PlacementGroupServer struct {
	ID string `json:"id"`

	Name string `json:"name"`

	PolicyRespected bool `json:"policy_respected"`
}

// SecurityGroup security group
type SecurityGroup struct {
	// ID display the security groups' unique ID
	ID string `json:"id"`
	// Name display the security groups name
	Name string `json:"name"`
	// Description display the security groups description
	Description string `json:"description"`
	// EnableDefaultSecurity display if the security group is set as default
	EnableDefaultSecurity bool `json:"enable_default_security"`
	// InboundDefaultPolicy display the default inbound policy
	//
	// Default value: accept
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy"`
	// OutboundDefaultPolicy display the default outbound policy
	//
	// Default value: accept
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy"`
	// Organization display the security groups organization ID
	Organization string `json:"organization"`
	// OrganizationDefault display if the security group is set as organization default
	OrganizationDefault bool `json:"organization_default"`
	// CreationDate display the security group creation date
	CreationDate time.Time `json:"creation_date"`
	// ModificationDate display the security group modification date
	ModificationDate time.Time `json:"modification_date"`
	// Servers list of servers attached to this security group
	Servers []*ServerSummary `json:"servers"`
	// Stateful true if the security group is stateful
	Stateful bool `json:"stateful"`
}

type SecurityGroupRule struct {
	ID string `json:"id"`
	// Protocol
	//
	// Default value: TCP
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	// Direction
	//
	// Default value: inbound
	Direction SecurityGroupRuleDirection `json:"direction"`
	// Action
	//
	// Default value: accept
	Action SecurityGroupRuleAction `json:"action"`

	IPRange string `json:"ip_range"`

	DestPortFrom *uint32 `json:"dest_port_from"`

	DestPortTo *uint32 `json:"dest_port_to"`

	Position uint32 `json:"position"`

	Editable bool `json:"editable"`
}

type SecurityGroupSummary struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

type SecurityGroupTemplate struct {
	ID string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

// Server server
type Server struct {
	// ID display the server unique ID
	ID string `json:"id"`
	// Name display the server name
	Name string `json:"name"`
	// Organization display the server organization
	Organization string `json:"organization"`
	// AllowedActions provide as list of allowed actions on the server
	AllowedActions []ServerAction `json:"allowed_actions"`
	// Tags display the server associated tags
	Tags []string `json:"tags"`
	// CommercialType display the server commercial type (eg. GP1-M)
	CommercialType string `json:"commercial_type"`
	// CreationDate display the server creation date
	CreationDate time.Time `json:"creation_date"`
	// DynamicIPRequired display if a dynamic IP is required
	DynamicIPRequired bool `json:"dynamic_ip_required"`
	// EnableIPv6 display if IPv6 is enabled
	EnableIPv6 bool `json:"enable_ipv6"`
	// Hostname display the server host name
	Hostname string `json:"hostname"`
	// Image provide information on the server image
	Image *Image `json:"image"`
	// Protected display the server protection option is activated
	Protected bool `json:"protected"`
	// PrivateIP display the server private IP address
	PrivateIP *string `json:"private_ip"`
	// PublicIP display the server public IP address
	PublicIP *ServerIP `json:"public_ip"`
	// ModificationDate display the server modification date
	ModificationDate time.Time `json:"modification_date"`
	// State display the server state
	//
	// Default value: running
	State ServerState `json:"state"`
	// Location display the server location
	Location *ServerLocation `json:"location"`
	// IPv6 display the server IPv6 address
	IPv6 *ServerIPv6 `json:"ipv6"`
	// Bootscript display the server bootscript
	Bootscript *Bootscript `json:"bootscript"`
	// BootType display the server boot type
	//
	// Default value: local
	BootType ServerBootType `json:"boot_type"`
	// Volumes display the server volumes
	Volumes map[string]*Volume `json:"volumes"`
	// SecurityGroup display the server security group
	SecurityGroup *SecurityGroupSummary `json:"security_group"`
	// Maintenances display the server planned maintenances
	Maintenances []*ServerMaintenance `json:"maintenances"`
	// StateDetail display the server state_detail
	StateDetail string `json:"state_detail"`
	// Arch display the server arch
	//
	// Default value: x86_64
	Arch Arch `json:"arch"`
	// PlacementGroup display the server placement group
	PlacementGroup *PlacementGroup `json:"placement_group"`
}

type ServerActionResponse struct {
	Task *Task `json:"task"`
}

// ServerIP server. ip
type ServerIP struct {
	// ID display the unique ID of the IP address
	ID string `json:"id"`
	// Address display the server public IPv4 IP-Address
	Address net.IP `json:"address"`
	// Dynamic display information if the IP address will be considered as dynamic
	Dynamic bool `json:"dynamic"`
}

// ServerIPv6 server. ipv6
type ServerIPv6 struct {
	// Address display the server IPv6 IP-Address
	Address net.IP `json:"address"`
	// Gateway display the IPv6 IP-addresses gateway
	Gateway net.IP `json:"gateway"`
	// Netmask display the IPv6 IP-addresses CIDR netmask
	Netmask string `json:"netmask"`
}

type ServerLocation struct {
	ClusterID string `json:"cluster_id"`

	HypervisorID string `json:"hypervisor_id"`

	NodeID string `json:"node_id"`

	PlatformID string `json:"platform_id"`

	ZoneID string `json:"zone_id"`
}

type ServerMaintenance struct {
}

type ServerSummary struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

type ServerType struct {
	MonthlyPrice float32 `json:"monthly_price"`

	HourlyPrice float32 `json:"hourly_price"`

	AltNames []string `json:"alt_names"`

	PerVolumeConstraint *ServerTypeVolumeConstraintsByType `json:"per_volume_constraint"`

	VolumesConstraint *ServerTypeVolumeConstraintSizes `json:"volumes_constraint"`

	Ncpus uint32 `json:"ncpus"`

	Gpu *uint64 `json:"gpu"`

	RAM uint64 `json:"ram"`
	// Arch
	//
	// Default value: x86_64
	Arch Arch `json:"arch"`

	Baremetal bool `json:"baremetal"`

	Network *ServerTypeNetwork `json:"network"`
}

type ServerTypeNetwork struct {
	Interfaces []*ServerTypeNetworkInterface `json:"interfaces"`

	SumInternalBandwidth *uint64 `json:"sum_internal_bandwidth"`

	SumInternetBandwidth *uint64 `json:"sum_internet_bandwidth"`

	IPv6Support bool `json:"ipv6_support"`
}

type ServerTypeNetworkInterface struct {
	InternalBandwidth *uint64 `json:"internal_bandwidth"`

	InternetBandwidth *uint64 `json:"internet_bandwidth"`
}

type ServerTypeVolumeConstraintSizes struct {
	MinSize scw.Size `json:"min_size"`

	MaxSize scw.Size `json:"max_size"`
}

type ServerTypeVolumeConstraintsByType struct {
	LSSD *ServerTypeVolumeConstraintSizes `json:"l_ssd"`
}

type SetPlacementGroupResponse struct {
	PlacementGroupID string `json:"placement_group_id"`
}

type SetPlacementGroupServersResponse struct {
	Servers []*PlacementGroupServer `json:"servers"`
}

type Snapshot struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Organization string `json:"organization"`
	// VolumeType
	//
	// Default value: l_ssd
	VolumeType VolumeType `json:"volume_type"`

	Size scw.Size `json:"size"`
	// State
	//
	// Default value: available
	State SnapshotState `json:"state"`

	BaseVolume *SnapshotBaseVolume `json:"base_volume"`

	CreationDate time.Time `json:"creation_date"`

	ModificationDate time.Time `json:"modification_date"`
}

type SnapshotBaseVolume struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

// Task task
type Task struct {
	// ID the unique ID of the task
	ID string `json:"id"`
	// Description the description of the task
	Description string `json:"description"`
	// Progress show the progress of the task in percent
	Progress int32 `json:"progress"`
	// StartedAt display the task start date
	StartedAt time.Time `json:"started_at"`
	// TerminatedAt display the task end date
	TerminatedAt time.Time `json:"terminated_at"`
	// Status display the task status
	//
	// Default value: pending
	Status TaskStatus `json:"status"`

	HrefFrom string `json:"href_from"`

	HrefResult string `json:"href_result"`
}

type UpdateIPResponse struct {
	IP *IP `json:"ip"`
}

type UpdatePlacementGroupResponse struct {
	PlacementGroupID string `json:"placement_group_id"`
}

type UpdatePlacementGroupServersResponse struct {
	Servers []*PlacementGroupServer `json:"servers"`
}

type UpdateServerResponse struct {
	Server *Server `json:"server"`
}

// Volume volume
type Volume struct {
	// ID display the volumes unique ID
	ID string `json:"id"`
	// Name display the volumes names
	Name string `json:"name"`
	// ExportURI show the volumes NBD export URI
	ExportURI string `json:"export_uri"`
	// Size display the volumes disk size
	Size scw.Size `json:"size"`
	// VolumeType display the volumes type
	//
	// Default value: l_ssd
	VolumeType VolumeType `json:"volume_type"`
	// CreationDate display the volumes creation date
	CreationDate time.Time `json:"creation_date"`
	// ModificationDate display the volumes modification date
	ModificationDate time.Time `json:"modification_date"`
	// Organization display the volumes organization
	Organization string `json:"organization"`
	// Server display information about the server attached to the volume
	Server *ServerSummary `json:"server"`
	// State display the volumes state
	//
	// Default value: available
	State VolumeState `json:"state"`
}

type VolumeSummary struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Size scw.Size `json:"size"`
	// VolumeType
	//
	// Default value: l_ssd
	VolumeType VolumeType `json:"volume_type"`
}

// VolumeTemplate volume template
type VolumeTemplate struct {
	// ID display the volumes unique ID
	ID string `json:"id,omitempty"`
	// Name display the volumes name
	Name string `json:"name,omitempty"`
	// Size display the volumes disk size
	Size scw.Size `json:"size,omitempty"`
	// VolumeType display the volumes type
	//
	// Default value: l_ssd
	VolumeType VolumeType `json:"volume_type,omitempty"`
	// Organization the organization ID
	Organization string `json:"organization,omitempty"`
}

// setIPResponse set ip response
type setIPResponse struct {
	IP *IP `json:"ip"`
}

// setImageResponse set image response
type setImageResponse struct {
	Image *Image `json:"image"`
}

// setSecurityGroupResponse set security group response
type setSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group"`
}

// setSecurityGroupRuleResponse set security group rule response
type setSecurityGroupRuleResponse struct {
	Rule *SecurityGroupRule `json:"rule"`
}

// setServerResponse set server response
type setServerResponse struct {
	Server *Server `json:"server"`
}

// setSnapshotResponse set snapshot response
type setSnapshotResponse struct {
	Snapshot *Snapshot `json:"snapshot"`
}

// Service API

type GetServerTypesAvailabilityRequest struct {
	Zone scw.Zone `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`
}

// GetServerTypesAvailability get availability
//
// Get availibility for all server types
func (s *API) GetServerTypesAvailability(req *GetServerTypesAvailabilityRequest, opts ...scw.RequestOption) (*GetServerTypesAvailabilityResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/products/servers/availability",
		Query:   query,
		Headers: http.Header{},
	}

	var resp GetServerTypesAvailabilityResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServersTypesRequest struct {
	Zone scw.Zone `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`
}

// ListServersTypes list server types
//
// Get server types technical details
func (s *API) ListServersTypes(req *ListServersTypesRequest, opts ...scw.RequestOption) (*ListServersTypesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/products/servers",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListServersTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServersRequest struct {
	Zone scw.Zone `json:"-"`
	// PerPage a positive integer lower or equal to 100 to select the number of items to display
	//
	// Default value: 20
	PerPage *uint32 `json:"-"`
	// Page a positive integer to choose the page to display
	Page *int32 `json:"-"`
	// Organization list only servers of this organization
	Organization *string `json:"-"`
	// Name filter servers by name (for eg. "server1" will return "server100" and "server1" but not "foo")
	Name *string `json:"-"`
	// PrivateIP list servers by private_ip
	PrivateIP *net.IP `json:"-"`
	// WithoutIP list servers that are not attached to a public IP
	WithoutIP *bool `json:"-"`
	// CommercialType list servers of this commercial type
	CommercialType *string `json:"-"`
}

// ListServers list servers
func (s *API) ListServers(req *ListServersRequest, opts ...scw.RequestOption) (*ListServersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "private_ip", req.PrivateIP)
	parameter.AddToQuery(query, "without_ip", req.WithoutIP)
	parameter.AddToQuery(query, "commercial_type", req.CommercialType)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeAppend(res interface{}) (uint32, scw.SdkError) {
	results, ok := res.(*ListServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint32(len(results.Servers))
	return uint32(len(results.Servers)), nil
}

type CreateServerRequest struct {
	Zone scw.Zone `json:"-"`
	// Name display the server name
	Name string `json:"name,omitempty"`
	// DynamicIPRequired define if a dynamic IP is required for the instance
	DynamicIPRequired *bool `json:"dynamic_ip_required,omitempty"`
	// CommercialType define the server commercial type (i.e. GP1-S)
	CommercialType string `json:"commercial_type,omitempty"`
	// Image the server image ID or label
	Image string `json:"image,omitempty"`
	// Volumes the volumes attached to the server
	Volumes map[string]*VolumeTemplate `json:"volumes,omitempty"`
	// EnableIPv6 true if IPv6 is enabled on the server
	EnableIPv6 bool `json:"enable_ipv6,omitempty"`
	// PublicIP the public IPv4 attached to the server
	PublicIP string `json:"public_ip,omitempty"`
	// Organization the server organization ID
	Organization string `json:"organization,omitempty"`
	// Tags the server tags
	Tags []string `json:"tags,omitempty"`
	// SecurityGroup the security group ID
	SecurityGroup string `json:"security_group,omitempty"`
	// PlacementGroup placement group ID if server must be part of a placement group
	PlacementGroup string `json:"placement_group,omitempty"`
}

// createServer create server
func (s *API) createServer(req *CreateServerRequest, opts ...scw.RequestOption) (*CreateServerResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("srv")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateServerResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteServerRequest struct {
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

// DeleteServer delete server
//
// Delete a server with the given id
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
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetServerRequest struct {
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

// GetServer get server
//
// Get the details of a specified Server
func (s *API) GetServer(req *GetServerRequest, opts ...scw.RequestOption) (*GetServerResponse, error) {
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
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
		Headers: http.Header{},
	}

	var resp GetServerResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type setServerRequest struct {
	Zone scw.Zone `json:"-"`
	// ID display the server unique ID
	ID string `json:"-"`
	// Name display the server name
	Name string `json:"name"`
	// Organization display the server organization
	Organization string `json:"organization"`
	// AllowedActions provide as list of allowed actions on the server
	AllowedActions []ServerAction `json:"allowed_actions"`
	// Tags display the server associated tags
	Tags []string `json:"tags"`
	// CommercialType display the server commercial type (eg. GP1-M)
	CommercialType string `json:"commercial_type"`
	// CreationDate display the server creation date
	CreationDate time.Time `json:"creation_date"`
	// DynamicIPRequired display if a dynamic IP is required
	DynamicIPRequired bool `json:"dynamic_ip_required"`
	// EnableIPv6 display if IPv6 is enabled
	EnableIPv6 bool `json:"enable_ipv6"`
	// Hostname display the server host name
	Hostname string `json:"hostname"`
	// Image provide information on the server image
	Image *Image `json:"image"`
	// Protected display the server protection option is activated
	Protected bool `json:"protected"`
	// PrivateIP display the server private IP address
	PrivateIP *string `json:"private_ip"`
	// PublicIP display the server public IP address
	PublicIP *ServerIP `json:"public_ip"`
	// ModificationDate display the server modification date
	ModificationDate time.Time `json:"modification_date"`
	// State display the server state
	//
	// Default value: running
	State ServerState `json:"state"`
	// Location display the server location
	Location *ServerLocation `json:"location"`
	// IPv6 display the server IPv6 address
	IPv6 *ServerIPv6 `json:"ipv6"`
	// Bootscript display the server bootscript
	Bootscript *Bootscript `json:"bootscript"`
	// BootType display the server boot type
	//
	// Default value: local
	BootType ServerBootType `json:"boot_type"`
	// Volumes display the server volumes
	Volumes map[string]*Volume `json:"volumes"`
	// SecurityGroup display the server security group
	SecurityGroup *SecurityGroupSummary `json:"security_group"`
	// Maintenances display the server planned maintenances
	Maintenances []*ServerMaintenance `json:"maintenances"`
	// StateDetail display the server state_detail
	StateDetail string `json:"state_detail"`
	// Arch display the server arch
	//
	// Default value: x86_64
	Arch Arch `json:"arch"`
	// PlacementGroup display the server placement group
	PlacementGroup *PlacementGroup `json:"placement_group"`
}

func (s *API) setServer(req *setServerRequest, opts ...scw.RequestOption) (*setServerResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setServerResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateServerRequest struct {
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	Name *string `json:"name,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	Volumes *map[string]*VolumeTemplate `json:"volumes,omitempty"`

	Bootscript string `json:"bootscript,omitempty"`

	DynamicIPRequired *bool `json:"dynamic_ip_required,omitempty"`

	EnableIPv6 *bool `json:"enable_ipv6,omitempty"`

	Protected *bool `json:"protected,omitempty"`

	SecurityGroup *SecurityGroupTemplate `json:"security_group,omitempty"`
	// PlacementGroup placement group ID if server must be part of a placement group
	PlacementGroup *NullableStringValue `json:"placement_group,omitempty"`
}

// updateServer update server
func (s *API) updateServer(req *UpdateServerRequest, opts ...scw.RequestOption) (*UpdateServerResponse, error) {
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
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateServerResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServerActionsRequest struct {
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

// ListServerActions list server actions
//
// Liste all actions that can currently be performed on a server
func (s *API) ListServerActions(req *ListServerActionsRequest, opts ...scw.RequestOption) (*ListServerActionsResponse, error) {
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
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/action",
		Headers: http.Header{},
	}

	var resp ListServerActionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ServerActionRequest struct {
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
	// Action
	//
	// Default value: poweron
	Action ServerAction `json:"action"`
}

// ServerAction perform action
//
// Perform power related actions on a server
func (s *API) ServerAction(req *ServerActionRequest, opts ...scw.RequestOption) (*ServerActionResponse, error) {
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
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/action",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ServerActionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServerUserDataRequest struct {
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

// ListServerUserData list user data
//
// List all user data keys register on a given server
func (s *API) ListServerUserData(req *ListServerUserDataRequest, opts ...scw.RequestOption) (*ListServerUserDataResponse, error) {
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
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user_data",
		Headers: http.Header{},
	}

	var resp ListServerUserDataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteServerUserDataRequest struct {
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	Key string `json:"-"`
}

// DeleteServerUserData delete user data
//
// Delete the given key from a server user data
func (s *API) DeleteServerUserData(req *DeleteServerUserDataRequest, opts ...scw.RequestOption) error {
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

	if fmt.Sprint(req.Key) == "" {
		return errors.New("field Key cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user_data/" + fmt.Sprint(req.Key) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListImagesRequest struct {
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`

	Name *string `json:"-"`

	Public *bool `json:"-"`

	Arch *string `json:"-"`
}

// ListImages list images
//
// List all images available in an account
func (s *API) ListImages(req *ListImagesRequest, opts ...scw.RequestOption) (*ListImagesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "public", req.Public)
	parameter.AddToQuery(query, "arch", req.Arch)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListImagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListImagesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListImagesResponse) UnsafeAppend(res interface{}) (uint32, scw.SdkError) {
	results, ok := res.(*ListImagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Images = append(r.Images, results.Images...)
	r.TotalCount += uint32(len(results.Images))
	return uint32(len(results.Images)), nil
}

type GetImageRequest struct {
	Zone scw.Zone `json:"-"`

	ImageID string `json:"-"`
}

// GetImage get image
//
// Get details of an image with the given id
func (s *API) GetImage(req *GetImageRequest, opts ...scw.RequestOption) (*GetImageResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.ImageID) + "",
		Headers: http.Header{},
	}

	var resp GetImageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateImageRequest struct {
	Zone scw.Zone `json:"-"`

	Name string `json:"name,omitempty"`

	RootVolume string `json:"root_volume,omitempty"`
	// Arch
	//
	// Default value: x86_64
	Arch Arch `json:"arch"`

	DefaultBootscript string `json:"default_bootscript,omitempty"`

	ExtraVolumes map[string]*VolumeTemplate `json:"extra_volumes,omitempty"`

	Organization string `json:"organization,omitempty"`

	Public bool `json:"public,omitempty"`
}

// CreateImage create image
func (s *API) CreateImage(req *CreateImageRequest, opts ...scw.RequestOption) (*CreateImageResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateImageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetImageRequest struct {
	Zone scw.Zone `json:"-"`

	ID string `json:"-"`

	Name string `json:"name"`
	// Arch
	//
	// Default value: x86_64
	Arch Arch `json:"arch"`

	CreationDate time.Time `json:"creation_date"`

	ModificationDate time.Time `json:"modification_date"`

	DefaultBootscript *Bootscript `json:"default_bootscript"`

	ExtraVolumes map[string]*Volume `json:"extra_volumes"`

	FromServer string `json:"from_server"`

	Organization string `json:"organization"`

	Public bool `json:"public"`

	RootVolume *VolumeSummary `json:"root_volume"`
	// State
	//
	// Default value: available
	State ImageState `json:"state"`
}

// setImage update image
//
// Replace all image properties with an image message
func (s *API) setImage(req *SetImageRequest, opts ...scw.RequestOption) (*setImageResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.ID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setImageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteImageRequest struct {
	Zone scw.Zone `json:"-"`

	ImageID string `json:"-"`
}

// DeleteImage delete image
//
// Delete the image with the given id
func (s *API) DeleteImage(req *DeleteImageRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ImageID) == "" {
		return errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.ImageID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListSnapshotsRequest struct {
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`

	Name *string `json:"-"`
}

// ListSnapshots list snapshots
func (s *API) ListSnapshots(req *ListSnapshotsRequest, opts ...scw.RequestOption) (*ListSnapshotsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListSnapshotsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSnapshotsResponse) UnsafeAppend(res interface{}) (uint32, scw.SdkError) {
	results, ok := res.(*ListSnapshotsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Snapshots = append(r.Snapshots, results.Snapshots...)
	r.TotalCount += uint32(len(results.Snapshots))
	return uint32(len(results.Snapshots)), nil
}

type CreateSnapshotRequest struct {
	Zone scw.Zone `json:"-"`

	VolumeID string `json:"volume_id,omitempty"`

	Organization string `json:"organization,omitempty"`

	Name string `json:"name,omitempty"`
}

// CreateSnapshot create snapshot
func (s *API) CreateSnapshot(req *CreateSnapshotRequest, opts ...scw.RequestOption) (*CreateSnapshotResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateSnapshotResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetSnapshotRequest struct {
	Zone scw.Zone `json:"-"`

	SnapshotID string `json:"-"`
}

// GetSnapshot get snapshot
//
// Get details of a snapshot with the given id
func (s *API) GetSnapshot(req *GetSnapshotRequest, opts ...scw.RequestOption) (*GetSnapshotResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return nil, errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
		Headers: http.Header{},
	}

	var resp GetSnapshotResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetSnapshotRequest struct {
	Zone scw.Zone `json:"-"`

	ID string `json:"-"`

	Name string `json:"name"`

	Organization string `json:"organization"`
	// VolumeType
	//
	// Default value: l_ssd
	VolumeType VolumeType `json:"volume_type"`

	Size scw.Size `json:"size"`
	// State
	//
	// Default value: available
	State SnapshotState `json:"state"`

	BaseVolume *SnapshotBaseVolume `json:"base_volume"`

	CreationDate time.Time `json:"creation_date"`

	ModificationDate time.Time `json:"modification_date"`
}

// setSnapshot update snapshot
//
// Replace all snapshot properties with a snapshot message
func (s *API) setSnapshot(req *SetSnapshotRequest, opts ...scw.RequestOption) (*setSnapshotResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.ID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setSnapshotResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteSnapshotRequest struct {
	Zone scw.Zone `json:"-"`

	SnapshotID string `json:"-"`
}

// DeleteSnapshot delete snapshot
//
// Delete the snapshot with the given id
func (s *API) DeleteSnapshot(req *DeleteSnapshotRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SnapshotID) == "" {
		return errors.New("field SnapshotID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListVolumesRequest struct {
	Zone scw.Zone `json:"-"`
	// VolumeType filter by volume type
	//
	// Default value: l_ssd
	VolumeType VolumeType `json:"-"`
	// PerPage a positive integer lower or equal to 100 to select the number of items to display
	//
	// Default value: 20
	PerPage *uint32 `json:"-"`
	// Page a positive integer to choose the page to display
	Page *int32 `json:"-"`
	// Organization display volumes of this organization
	Organization *string `json:"-"`
	// Name filter volume by name (for eg. "vol" will return "myvolume" but not "data")
	Name *string `json:"-"`
}

// ListVolumes list volumes
func (s *API) ListVolumes(req *ListVolumesRequest, opts ...scw.RequestOption) (*ListVolumesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "volume_type", req.VolumeType)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListVolumesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVolumesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVolumesResponse) UnsafeAppend(res interface{}) (uint32, scw.SdkError) {
	results, ok := res.(*ListVolumesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Volumes = append(r.Volumes, results.Volumes...)
	r.TotalCount += uint32(len(results.Volumes))
	return uint32(len(results.Volumes)), nil
}

type CreateVolumeRequest struct {
	Zone scw.Zone `json:"-"`

	Name string `json:"name,omitempty"`

	Organization string `json:"organization,omitempty"`
	// VolumeType
	//
	// Default value: l_ssd
	VolumeType VolumeType `json:"volume_type"`

	// Precisely one of BaseSnapshot, BaseVolume, Size must be set.
	Size *scw.Size `json:"size,omitempty"`

	// Precisely one of BaseSnapshot, BaseVolume, Size must be set.
	BaseVolume *string `json:"base_volume,omitempty"`

	// Precisely one of BaseSnapshot, BaseVolume, Size must be set.
	BaseSnapshot *string `json:"base_snapshot,omitempty"`
}

func (m *CreateVolumeRequest) GetFrom() From {
	switch {
	case m.Size != nil:
		return FromSize{*m.Size}
	case m.BaseVolume != nil:
		return FromBaseVolume{*m.BaseVolume}
	case m.BaseSnapshot != nil:
		return FromBaseSnapshot{*m.BaseSnapshot}
	}
	return nil
}

// CreateVolume create volume
func (s *API) CreateVolume(req *CreateVolumeRequest, opts ...scw.RequestOption) (*CreateVolumeResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateVolumeResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetVolumeRequest struct {
	Zone scw.Zone `json:"-"`

	VolumeID string `json:"-"`
}

// GetVolume get volume
//
// Get details of a volume with the given id
func (s *API) GetVolume(req *GetVolumeRequest, opts ...scw.RequestOption) (*GetVolumeResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return nil, errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
		Headers: http.Header{},
	}

	var resp GetVolumeResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteVolumeRequest struct {
	Zone scw.Zone `json:"-"`

	VolumeID string `json:"-"`
}

// DeleteVolume delete volume
//
// Delete the volume with the given id
func (s *API) DeleteVolume(req *DeleteVolumeRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.VolumeID) == "" {
		return errors.New("field VolumeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListSecurityGroupsRequest struct {
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`

	Name *string `json:"-"`
}

// ListSecurityGroups list security groups
//
// List all security groups available in an account
func (s *API) ListSecurityGroups(req *ListSecurityGroupsRequest, opts ...scw.RequestOption) (*ListSecurityGroupsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListSecurityGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecurityGroupsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecurityGroupsResponse) UnsafeAppend(res interface{}) (uint32, scw.SdkError) {
	results, ok := res.(*ListSecurityGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SecurityGroups = append(r.SecurityGroups, results.SecurityGroups...)
	r.TotalCount += uint32(len(results.SecurityGroups))
	return uint32(len(results.SecurityGroups)), nil
}

type CreateSecurityGroupRequest struct {
	Zone scw.Zone `json:"-"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Organization string `json:"organization,omitempty"`
	// OrganizationDefault
	//
	// Default value: false
	OrganizationDefault bool `json:"organization_default,omitempty"`
	// Stateful
	//
	// Default value: false
	Stateful bool `json:"stateful,omitempty"`
	// InboundDefaultPolicy
	//
	// Default value: accept
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy"`
	// OutboundDefaultPolicy
	//
	// Default value: accept
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy"`
}

// CreateSecurityGroup create security group
func (s *API) CreateSecurityGroup(req *CreateSecurityGroupRequest, opts ...scw.RequestOption) (*CreateSecurityGroupResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("sg")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateSecurityGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetSecurityGroupRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`
}

// GetSecurityGroup get security group
//
// Get the details of a Security Group with the given id
func (s *API) GetSecurityGroup(req *GetSecurityGroupRequest, opts ...scw.RequestOption) (*GetSecurityGroupResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "",
		Headers: http.Header{},
	}

	var resp GetSecurityGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteSecurityGroupRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`
}

// DeleteSecurityGroup delete security group
func (s *API) DeleteSecurityGroup(req *DeleteSecurityGroupRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type setSecurityGroupRequest struct {
	Zone scw.Zone `json:"-"`
	// ID display the security groups' unique ID
	ID string `json:"-"`
	// Name display the security groups name
	Name string `json:"name"`
	// Description display the security groups description
	Description string `json:"description"`
	// EnableDefaultSecurity display if the security group is set as default
	EnableDefaultSecurity bool `json:"enable_default_security"`
	// InboundDefaultPolicy display the default inbound policy
	//
	// Default value: accept
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy"`
	// OutboundDefaultPolicy display the default outbound policy
	//
	// Default value: accept
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy"`
	// Organization display the security groups organization ID
	Organization string `json:"organization"`
	// OrganizationDefault display if the security group is set as organization default
	OrganizationDefault bool `json:"organization_default"`
	// CreationDate display the security group creation date
	CreationDate time.Time `json:"creation_date"`
	// ModificationDate display the security group modification date
	ModificationDate time.Time `json:"modification_date"`
	// Servers list of servers attached to this security group
	Servers []*ServerSummary `json:"servers"`
	// Stateful true if the security group is stateful
	Stateful bool `json:"stateful"`
}

// setSecurityGroup update security group
//
// Replace all security group properties with a security group message
func (s *API) setSecurityGroup(req *setSecurityGroupRequest, opts ...scw.RequestOption) (*setSecurityGroupResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.ID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setSecurityGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListSecurityGroupRulesRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`
}

// ListSecurityGroupRules list rules
func (s *API) ListSecurityGroupRules(req *ListSecurityGroupRulesRequest, opts ...scw.RequestOption) (*ListSecurityGroupRulesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListSecurityGroupRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecurityGroupRulesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecurityGroupRulesResponse) UnsafeAppend(res interface{}) (uint32, scw.SdkError) {
	results, ok := res.(*ListSecurityGroupRulesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Rules = append(r.Rules, results.Rules...)
	r.TotalCount += uint32(len(results.Rules))
	return uint32(len(results.Rules)), nil
}

type CreateSecurityGroupRuleRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`
	// Protocol
	//
	// Default value: TCP
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	// Direction
	//
	// Default value: inbound
	Direction SecurityGroupRuleDirection `json:"direction"`
	// Action
	//
	// Default value: accept
	Action SecurityGroupRuleAction `json:"action"`

	IPRange string `json:"ip_range,omitempty"`

	DestPortFrom *uint32 `json:"dest_port_from,omitempty"`

	DestPortTo *uint32 `json:"dest_port_to,omitempty"`

	Position uint32 `json:"position,omitempty"`

	Editable bool `json:"editable,omitempty"`
}

// CreateSecurityGroupRule create rule
func (s *API) CreateSecurityGroupRule(req *CreateSecurityGroupRuleRequest, opts ...scw.RequestOption) (*CreateSecurityGroupRuleResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateSecurityGroupRuleResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteSecurityGroupRuleRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`

	SecurityGroupRuleID string `json:"-"`
}

// DeleteSecurityGroupRule delete rule
//
// Delete a security group rule with the given id
func (s *API) DeleteSecurityGroupRule(req *DeleteSecurityGroupRuleRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return errors.New("field SecurityGroupID cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupRuleID) == "" {
		return errors.New("field SecurityGroupRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules/" + fmt.Sprint(req.SecurityGroupRuleID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetSecurityGroupRuleRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`

	SecurityGroupRuleID string `json:"-"`
}

// GetSecurityGroupRule get rule
//
// Get details of a security group rule with the given id
func (s *API) GetSecurityGroupRule(req *GetSecurityGroupRuleRequest, opts ...scw.RequestOption) (*GetSecurityGroupRuleResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupRuleID) == "" {
		return nil, errors.New("field SecurityGroupRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules/" + fmt.Sprint(req.SecurityGroupRuleID) + "",
		Headers: http.Header{},
	}

	var resp GetSecurityGroupRuleResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type setSecurityGroupRuleRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`

	SecurityGroupRuleID string `json:"-"`

	ID string `json:"id"`
	// Protocol
	//
	// Default value: TCP
	Protocol SecurityGroupRuleProtocol `json:"protocol"`
	// Direction
	//
	// Default value: inbound
	Direction SecurityGroupRuleDirection `json:"direction"`
	// Action
	//
	// Default value: accept
	Action SecurityGroupRuleAction `json:"action"`

	IPRange string `json:"ip_range"`

	DestPortFrom *uint32 `json:"dest_port_from"`

	DestPortTo *uint32 `json:"dest_port_to"`

	Position uint32 `json:"position"`

	Editable bool `json:"editable"`
}

// setSecurityGroupRule update security group rule
func (s *API) setSecurityGroupRule(req *setSecurityGroupRuleRequest, opts ...scw.RequestOption) (*setSecurityGroupRuleResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupID) == "" {
		return nil, errors.New("field SecurityGroupID cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupRuleID) == "" {
		return nil, errors.New("field SecurityGroupRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupID) + "/rules/" + fmt.Sprint(req.SecurityGroupRuleID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setSecurityGroupRuleResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListPlacementGroupsRequest struct {
	Zone scw.Zone `json:"-"`
	// PerPage a positive integer lower or equal to 100 to select the number of items to display
	//
	// Default value: 20
	PerPage *uint32 `json:"-"`
	// Page a positive integer to choose the page to display
	Page *int32 `json:"-"`
	// Organization list only placement groups of this organization
	Organization *string `json:"-"`
	// Name filter placement groups by name (for eg. "cluster1" will return "cluster100" and "cluster1" but not "foo")
	Name *string `json:"-"`
}

// ListPlacementGroups list placement groups
//
// List all placement groups
func (s *API) ListPlacementGroups(req *ListPlacementGroupsRequest, opts ...scw.RequestOption) (*ListPlacementGroupsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListPlacementGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPlacementGroupsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPlacementGroupsResponse) UnsafeAppend(res interface{}) (uint32, scw.SdkError) {
	results, ok := res.(*ListPlacementGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PlacementGroups = append(r.PlacementGroups, results.PlacementGroups...)
	r.TotalCount += uint32(len(results.PlacementGroups))
	return uint32(len(results.PlacementGroups)), nil
}

type CreatePlacementGroupRequest struct {
	Zone scw.Zone `json:"-"`

	Name string `json:"name,omitempty"`

	Organization string `json:"organization,omitempty"`
	// PolicyMode
	//
	// Default value: optional
	PolicyMode PlacementGroupPolicyMode `json:"policy_mode"`
	// PolicyType
	//
	// Default value: max_availability
	PolicyType PlacementGroupPolicyType `json:"policy_type"`
}

// CreatePlacementGroup create placement group
//
// Create a new placement group
func (s *API) CreatePlacementGroup(req *CreatePlacementGroupRequest, opts ...scw.RequestOption) (*CreatePlacementGroupResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreatePlacementGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetPlacementGroupRequest struct {
	Zone scw.Zone `json:"-"`

	PlacementGroupID string `json:"-"`
}

// GetPlacementGroup get placement group
//
// Get the given placement group
func (s *API) GetPlacementGroup(req *GetPlacementGroupRequest, opts ...scw.RequestOption) (*GetPlacementGroupResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "",
		Headers: http.Header{},
	}

	var resp GetPlacementGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetPlacementGroupRequest struct {
	Zone scw.Zone `json:"-"`

	PlacementGroupID string `json:"-"`

	Name string `json:"name"`

	Organization string `json:"organization"`
	// PolicyMode
	//
	// Default value: optional
	PolicyMode PlacementGroupPolicyMode `json:"policy_mode"`
	// PolicyType
	//
	// Default value: max_availability
	PolicyType PlacementGroupPolicyType `json:"policy_type"`
}

// SetPlacementGroup set placement group
//
// Set all parameters of the given placement group
func (s *API) SetPlacementGroup(req *SetPlacementGroupRequest, opts ...scw.RequestOption) (*SetPlacementGroupResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetPlacementGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdatePlacementGroupRequest struct {
	Zone scw.Zone `json:"-"`

	PlacementGroupID string `json:"-"`

	Name *string `json:"name,omitempty"`

	Organization *string `json:"organization,omitempty"`
	// PolicyMode
	//
	// Default value: optional
	PolicyMode PlacementGroupPolicyMode `json:"policy_mode"`
	// PolicyType
	//
	// Default value: max_availability
	PolicyType PlacementGroupPolicyType `json:"policy_type"`
}

// UpdatePlacementGroup update placement group
//
// Update one or more parameter of the given placement group
func (s *API) UpdatePlacementGroup(req *UpdatePlacementGroupRequest, opts ...scw.RequestOption) (*UpdatePlacementGroupResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdatePlacementGroupResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeletePlacementGroupRequest struct {
	Zone scw.Zone `json:"-"`

	PlacementGroupID string `json:"-"`
}

// DeletePlacementGroup delete the given placement group
//
// Delete the given placement group
func (s *API) DeletePlacementGroup(req *DeletePlacementGroupRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetPlacementGroupServersRequest struct {
	Zone scw.Zone `json:"-"`

	PlacementGroupID string `json:"-"`
}

// GetPlacementGroupServers get placement group servers
//
// Get all servers belonging to the given placement group
func (s *API) GetPlacementGroupServers(req *GetPlacementGroupServersRequest, opts ...scw.RequestOption) (*GetPlacementGroupServersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "/servers",
		Headers: http.Header{},
	}

	var resp GetPlacementGroupServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetPlacementGroupServersRequest struct {
	Zone scw.Zone `json:"-"`

	PlacementGroupID string `json:"-"`
}

// SetPlacementGroupServers set placement group servers
//
// Set all servers belonging to the given placement group
func (s *API) SetPlacementGroupServers(req *SetPlacementGroupServersRequest, opts ...scw.RequestOption) (*SetPlacementGroupServersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "/servers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetPlacementGroupServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdatePlacementGroupServersRequest struct {
	Zone scw.Zone `json:"-"`

	PlacementGroupID string `json:"-"`
}

// UpdatePlacementGroupServers update placement group servers
//
// Update all servers belonging to the given placement group
func (s *API) UpdatePlacementGroupServers(req *UpdatePlacementGroupServersRequest, opts ...scw.RequestOption) (*UpdatePlacementGroupServersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return nil, errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "/servers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdatePlacementGroupServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeletePlacementGroupServersRequest struct {
	Zone scw.Zone `json:"-"`

	PlacementGroupID string `json:"-"`
}

// DeletePlacementGroupServers delete placement group servers
//
// Delete all servers from the given placement group
func (s *API) DeletePlacementGroupServers(req *DeletePlacementGroupServersRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PlacementGroupID) == "" {
		return errors.New("field PlacementGroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/placement_groups/" + fmt.Sprint(req.PlacementGroupID) + "/servers",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListIPsRequest struct {
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`

	Name *string `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`
}

// ListIPs list IPs
func (s *API) ListIPs(req *ListIPsRequest, opts ...scw.RequestOption) (*ListIPsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeAppend(res interface{}) (uint32, scw.SdkError) {
	results, ok := res.(*ListIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.IPs = append(r.IPs, results.IPs...)
	r.TotalCount += uint32(len(results.IPs))
	return uint32(len(results.IPs)), nil
}

type CreateIPRequest struct {
	Zone scw.Zone `json:"-"`

	Organization string `json:"organization,omitempty"`

	Server *string `json:"server,omitempty"`
}

// CreateIP reseve an IP
func (s *API) CreateIP(req *CreateIPRequest, opts ...scw.RequestOption) (*CreateIPResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetIPRequest struct {
	Zone scw.Zone `json:"-"`

	IPID string `json:"-"`
}

// GetIP get IP
//
// Get details of an IP with the given id
func (s *API) GetIP(req *GetIPRequest, opts ...scw.RequestOption) (*GetIPResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IPID) + "",
		Headers: http.Header{},
	}

	var resp GetIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetIPRequest struct {
	Zone scw.Zone `json:"-"`

	ID string `json:"-"`

	Address net.IP `json:"address"`

	Reverse *string `json:"reverse"`

	Server *ServerSummary `json:"server"`

	Organization string `json:"organization"`
}

func (s *API) setIP(req *SetIPRequest, opts ...scw.RequestOption) (*setIPResponse, error) {
	var err error

	if req.Organization == "" {
		defaultOrganization, _ := s.client.GetDefaultOrganizationID()
		req.Organization = defaultOrganization
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.ID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp setIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type updateIPRequest struct {
	Zone scw.Zone `json:"-"`

	IPID string `json:"-"`

	Reverse *NullableStringValue `json:"reverse,omitempty"`

	Server *NullableStringValue `json:"server,omitempty"`
}

// updateIP update IP
func (s *API) updateIP(req *updateIPRequest, opts ...scw.RequestOption) (*UpdateIPResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IPID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteIPRequest struct {
	Zone scw.Zone `json:"-"`

	IPID string `json:"-"`
}

// DeleteIP delete IP
//
// Delete the IP with the given id
func (s *API) DeleteIP(req *DeleteIPRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IPID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListBootscriptsRequest struct {
	Zone scw.Zone `json:"-"`

	Arch *string `json:"-"`

	Title *string `json:"-"`

	Default *bool `json:"-"`

	Public *bool `json:"-"`

	PerPage *uint32 `json:"-"`

	Page *int32 `json:"-"`
}

// ListBootscripts list bootscripts
func (s *API) ListBootscripts(req *ListBootscriptsRequest, opts ...scw.RequestOption) (*ListBootscriptsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPerPage, exist := s.client.GetDefaultPageSize()
	if (req.PerPage == nil || *req.PerPage == 0) && exist {
		req.PerPage = &defaultPerPage
	}

	query := url.Values{}
	parameter.AddToQuery(query, "arch", req.Arch)
	parameter.AddToQuery(query, "title", req.Title)
	parameter.AddToQuery(query, "default", req.Default)
	parameter.AddToQuery(query, "public", req.Public)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/bootscripts",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListBootscriptsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBootscriptsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBootscriptsResponse) UnsafeAppend(res interface{}) (uint32, scw.SdkError) {
	results, ok := res.(*ListBootscriptsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Bootscripts = append(r.Bootscripts, results.Bootscripts...)
	r.TotalCount += uint32(len(results.Bootscripts))
	return uint32(len(results.Bootscripts)), nil
}

type GetBootscriptRequest struct {
	Zone scw.Zone `json:"-"`

	BootscriptID string `json:"-"`
}

// GetBootscript get bootscripts
//
// Get details of a bootscript with the given id
func (s *API) GetBootscript(req *GetBootscriptRequest, opts ...scw.RequestOption) (*GetBootscriptResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.BootscriptID) == "" {
		return nil, errors.New("field BootscriptID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/bootscripts/" + fmt.Sprint(req.BootscriptID) + "",
		Headers: http.Header{},
	}

	var resp GetBootscriptResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDashboardRequest struct {
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`
}

func (s *API) GetDashboard(req *GetDashboardRequest, opts ...scw.RequestOption) (*GetDashboardResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/dashboard",
		Query:   query,
		Headers: http.Header{},
	}

	var resp GetDashboardResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type From interface {
	isFrom()
}

type FromSize struct {
	Value scw.Size
}

func (FromSize) isFrom() {
}

type FromBaseVolume struct {
	Value string
}

func (FromBaseVolume) isFrom() {
}

type FromBaseSnapshot struct {
	Value string
}

func (FromBaseSnapshot) isFrom() {
}

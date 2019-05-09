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

	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/utils"
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
	_ utils.File
	_ = parameter.AddToQuery
)

// Api: instance API
type Api struct {
	client *scw.Client
}

// NewApi returns a Api object from a Scaleway client.
func NewApi(client *scw.Client) *Api {
	return &Api{
		client: client,
	}
}

type Arch string

const (
	// ArchX8664 is [insert doc].
	ArchX8664 = Arch("x86_64")
	// ArchArm is [insert doc].
	ArchArm = Arch("arm")
)

type GetServerTypesAvailabilityResponseAvailability string

const (
	// GetServerTypesAvailabilityResponseAvailabilityAvailable is [insert doc].
	GetServerTypesAvailabilityResponseAvailabilityAvailable = GetServerTypesAvailabilityResponseAvailability("available")
	// GetServerTypesAvailabilityResponseAvailabilityScarce is [insert doc].
	GetServerTypesAvailabilityResponseAvailabilityScarce = GetServerTypesAvailabilityResponseAvailability("scarce")
	// GetServerTypesAvailabilityResponseAvailabilityShortage is [insert doc].
	GetServerTypesAvailabilityResponseAvailabilityShortage = GetServerTypesAvailabilityResponseAvailability("shortage")
)

type ImageState string

const (
	// ImageStateAvailable is [insert doc].
	ImageStateAvailable = ImageState("available")
	// ImageStateCreating is [insert doc].
	ImageStateCreating = ImageState("creating")
	// ImageStateError is [insert doc].
	ImageStateError = ImageState("error")
)

type SecurityGroupPolicy string

const (
	// SecurityGroupPolicyAccept is [insert doc].
	SecurityGroupPolicyAccept = SecurityGroupPolicy("accept")
	// SecurityGroupPolicyDrop is [insert doc].
	SecurityGroupPolicyDrop = SecurityGroupPolicy("drop")
)

type SecurityRuleAction string

const (
	// SecurityRuleActionAccept is [insert doc].
	SecurityRuleActionAccept = SecurityRuleAction("accept")
	// SecurityRuleActionDrop is [insert doc].
	SecurityRuleActionDrop = SecurityRuleAction("drop")
)

type SecurityRuleDirection string

const (
	// SecurityRuleDirectionInbound is [insert doc].
	SecurityRuleDirectionInbound = SecurityRuleDirection("inbound")
	// SecurityRuleDirectionOutbound is [insert doc].
	SecurityRuleDirectionOutbound = SecurityRuleDirection("outbound")
)

type SecurityRuleProtocol string

const (
	// SecurityRuleProtocolTcp is [insert doc].
	SecurityRuleProtocolTcp = SecurityRuleProtocol("tcp")
	// SecurityRuleProtocolUdp is [insert doc].
	SecurityRuleProtocolUdp = SecurityRuleProtocol("udp")
	// SecurityRuleProtocolIcmp is [insert doc].
	SecurityRuleProtocolIcmp = SecurityRuleProtocol("icmp")
)

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

type ServerBootType string

const (
	// ServerBootTypeLocal is [insert doc].
	ServerBootTypeLocal = ServerBootType("local")
)

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

type SnapshotState string

const (
	// SnapshotStateAvailable is [insert doc].
	SnapshotStateAvailable = SnapshotState("available")
	// SnapshotStateSnapshotting is [insert doc].
	SnapshotStateSnapshotting = SnapshotState("snapshotting")
	// SnapshotStateError is [insert doc].
	SnapshotStateError = SnapshotState("error")
)

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

type VolumeState string

const (
	// VolumeStateAvailable is [insert doc].
	VolumeStateAvailable = VolumeState("available")
	// VolumeStateSnapshotting is [insert doc].
	VolumeStateSnapshotting = VolumeState("snapshotting")
	// VolumeStateError is [insert doc].
	VolumeStateError = VolumeState("error")
)

type VolumeType string

const (
	// VolumeTypeLSsd is [insert doc].
	VolumeTypeLSsd = VolumeType("l_ssd")
	// VolumeTypeLHdd is [insert doc].
	VolumeTypeLHdd = VolumeType("l_hdd")
	// VolumeTypeRSsd is [insert doc].
	VolumeTypeRSsd = VolumeType("r_ssd")
)

type Bootscript struct {
	// Arch: display the bootscripts arch
	Arch Arch `json:"arch,omitempty"`
	// Bootcmdargs: display the bootscript parameters
	Bootcmdargs string `json:"bootcmdargs,omitempty"`
	// Default: dispmay if the bootscript is the default bootscript if no other boot option is configured
	Default bool `json:"default,omitempty"`
	// Dtb: provide information regarding a Device Tree Binary (dtb) for use with C1 servers
	Dtb string `json:"dtb,omitempty"`
	// Id: display the bootscripts ID
	Id string `json:"id,omitempty"`
	// Initrd: display the initrd (initial ramdisk) configuration
	Initrd string `json:"initrd,omitempty"`
	// Kernel: display the server kernel version
	Kernel string `json:"kernel,omitempty"`
	// Organization: display the bootscripts organization
	Organization string `json:"organization,omitempty"`
	// Public: provide information if the bootscript is public
	Public bool `json:"public,omitempty"`
	// Title: display the bootscripts title
	Title string `json:"title,omitempty"`
}

type CreateImageResponse struct {
	Image *Image `json:"image,omitempty"`

	Location string `json:"Location,omitempty"`
}

type CreateIpResponse struct {
	Ip *Ip `json:"ip,omitempty"`

	Location string `json:"Location,omitempty"`
}

type CreateSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group,omitempty"`
}

type CreateSecurityGroupRuleResponse struct {
	SecurityRule *SecurityRule `json:"security_rule,omitempty"`
}

type CreateServerResponse struct {
	Server *Server `json:"server,omitempty"`
}

type CreateSnapshotResponse struct {
	Snapshot *Snapshot `json:"snapshot,omitempty"`
}

type CreateVolumeResponse struct {
	Volume *Volume `json:"volume,omitempty"`

	Location string `json:"Location,omitempty"`
}

type Dashboard struct {
	VolumesCount uint32 `json:"volumes_count,omitempty"`

	RunningServersCount uint32 `json:"running_servers_count,omitempty"`

	ServersByTypes map[string]uint32 `json:"servers_by_types,omitempty"`

	ImagesCount uint32 `json:"images_count,omitempty"`

	SnapshotsCount uint32 `json:"snapshots_count,omitempty"`

	ServersCount uint32 `json:"servers_count,omitempty"`

	IpsCount uint32 `json:"ips_count,omitempty"`

	SecurityGroupsCount uint32 `json:"security_groups_count,omitempty"`

	IpsUnused uint32 `json:"ips_unused,omitempty"`
}

type GetBootscriptResponse struct {
	Bootscript *Bootscript `json:"bootscript,omitempty"`
}

type GetDashboardResponse struct {
	Dashboard *Dashboard `json:"dashboard,omitempty"`
}

type GetImageResponse struct {
	Image *Image `json:"image,omitempty"`
}

type GetIpResponse struct {
	Ip *Ip `json:"ip,omitempty"`
}

type GetSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group,omitempty"`
}

type GetSecurityGroupRuleResponse struct {
	SecurityRule *SecurityRule `json:"security_rule,omitempty"`
}

type GetServerResponse struct {
	Server *Server `json:"server,omitempty"`
}

type GetServerTypesAvailabilityResponse struct {
	Servers map[string]GetServerTypesAvailabilityResponseAvailability `json:"servers,omitempty"`
}

type GetServiceInfoResponse struct {
	Api string `json:"api,omitempty"`

	Description string `json:"description,omitempty"`

	Version string `json:"version,omitempty"`
}

type GetSnapshotResponse struct {
	Snapshot *Snapshot `json:"snapshot,omitempty"`
}

type GetVolumeResponse struct {
	Volume *Volume `json:"volume,omitempty"`
}

type Image struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Arch Arch `json:"arch,omitempty"`

	CreationDate time.Time `json:"creation_date,omitempty"`

	ModificationDate time.Time `json:"modification_date,omitempty"`

	DefaultBootscript *Bootscript `json:"default_bootscript,omitempty"`

	ExtraVolumes map[string]*Volume `json:"extra_volumes,omitempty"`

	FromServer *ServerSummary `json:"from_server,omitempty"`

	Organization string `json:"organization,omitempty"`

	Public bool `json:"public,omitempty"`

	RootVolume *VolumeTemplate `json:"root_volume,omitempty"`

	State ImageState `json:"state,omitempty"`
}

type Ip struct {
	Id string `json:"id,omitempty"`

	Address net.IP `json:"address,omitempty"`

	Reverse string `json:"reverse,omitempty"`

	Server string `json:"server,omitempty"`

	Organization string `json:"organization,omitempty"`
}

type ListBootscriptsResponse struct {
	Bootscripts []*Bootscript `json:"bootscripts,omitempty"`
}

type ListImagesResponse struct {
	Images []*Image `json:"images,omitempty"`
}

type ListIpsResponse struct {
	Ips []*Ip `json:"ips,omitempty"`
}

type ListSecurityGroupRulesResponse struct {
	SecurityRules []*SecurityRule `json:"security_rules,omitempty"`
}

type ListSecurityGroupsResponse struct {
	SecurityGroups []*SecurityGroup `json:"security_groups,omitempty"`
}

type ListServerActionsResponse struct {
	Actions []ServerAction `json:"actions,omitempty"`
}

type ListServerUserDataResponse struct {
	UserData []string `json:"user_data,omitempty"`
}

type ListServersResponse struct {
	Servers []*Server `json:"servers,omitempty"`
}

type ListServersTypesResponse struct {
	Servers map[string]*ServerTypeDefinition `json:"servers,omitempty"`
}

type ListSnapshotsResponse struct {
	Snapshots []*Snapshot `json:"snapshots,omitempty"`
}

type ListVolumesResponse struct {
	Volumes []*Volume `json:"volumes,omitempty"`
}

type SecurityGroup struct {
	// Id: display the security groups' unique ID
	Id string `json:"id,omitempty"`
	// Name: display the security groups name
	Name string `json:"name,omitempty"`
	// CreationDate: display the security group creation date
	CreationDate time.Time `json:"creation_date,omitempty"`
	// ModificationDate: display the security group modification date
	ModificationDate time.Time `json:"modification_date,omitempty"`
	// Description: display the security groups description
	Description string `json:"description,omitempty"`
	// EnableDefaultSecurity: display if the security group is set as default
	EnableDefaultSecurity bool `json:"enable_default_security,omitempty"`
	// InboundDefaultPolicy: display the default inbound policy
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy,omitempty"`
	// Organization: display the security groups organization ID
	Organization string `json:"organization,omitempty"`
	// OrganizationDefault: display if the security group is set as organization default
	OrganizationDefault bool `json:"organization_default,omitempty"`
	// OutboundDefaultPolicy: display the default outbound policy
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy,omitempty"`
	// Servers: list of servers attached to this security group
	Servers map[string]*ServerSummary `json:"servers,omitempty"`
	// Stateful: true if the security group is stateful
	Stateful bool `json:"stateful,omitempty"`
}

type SecurityGroupSummary struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

type SecurityRule struct {
	Id string `json:"id,omitempty"`

	Protocol SecurityRuleProtocol `json:"protocol,omitempty"`

	Direction SecurityRuleDirection `json:"direction,omitempty"`

	Action SecurityRuleAction `json:"action,omitempty"`

	IpRange string `json:"ip_range,omitempty"`

	DestPortFrom uint32 `json:"dest_port_from,omitempty"`

	DestPortTo uint32 `json:"dest_port_to,omitempty"`

	Position uint32 `json:"position,omitempty"`

	Editable bool `json:"editable,omitempty"`
}

type Server struct {
	// Id: display the server unique ID
	Id string `json:"id,omitempty"`
	// Image: provide information on the server image
	Image *Image `json:"image,omitempty"`
	// Name: display the server name
	Name string `json:"name,omitempty"`
	// Organization: display the server organization
	Organization string `json:"organization,omitempty"`
	// PrivateIp: display the server private IP address
	PrivateIp *string `json:"private_ip,omitempty"`
	// PublicIp: display the server public IP address
	PublicIp *ServerIp `json:"public_ip,omitempty"`
	// State: display the server state
	State ServerState `json:"state,omitempty"`
	// BootType: display the server boot type
	BootType ServerBootType `json:"boot_type,omitempty"`
	// Tags: display the server associated tags
	Tags []string `json:"tags,omitempty"`
	// Volumes: display the server volumes
	Volumes map[string]*Volume `json:"volumes,omitempty"`
	// Bootscript: display the server bootscript
	Bootscript *Bootscript `json:"bootscript,omitempty"`
	// DynamicPublicIp: display the server dynamic public IP
	DynamicPublicIp bool `json:"dynamic_public_ip,omitempty"`
	// CommercialType: display the server commercial type (e.g. GP1-M)
	CommercialType string `json:"commercial_type,omitempty"`
	// CreationDate: display the server creation date
	CreationDate time.Time `json:"creation_date,omitempty"`
	// DynamicIpRequired: display if a dynamic IP is required
	DynamicIpRequired bool `json:"dynamic_ip_required,omitempty"`
	// EnableIpv6: display if IPv6 is enabled
	EnableIpv6 bool `json:"enable_ipv6,omitempty"`
	// ExtraNetworks: display information about additional network interfaces
	ExtraNetworks []string `json:"extra_networks,omitempty"`
	// Hostname: display the server host name
	Hostname string `json:"hostname,omitempty"`
	// AllowedActions: provide as list of allowed actions on the server
	AllowedActions []ServerAction `json:"allowed_actions,omitempty"`
	// Arch: display the server arch
	Arch Arch `json:"arch,omitempty"`
	// Ipv6: display the server IPv6 address
	Ipv6 *ServerIpv6 `json:"ipv6,omitempty"`
	// Location: display the server location
	Location *ServerLocation `json:"location,omitempty"`
	// Maintenances: display the server planned maintenances
	Maintenances []*ServerMaintenance `json:"maintenances,omitempty"`
	// ModificationDate: display the server modification date
	ModificationDate time.Time `json:"modification_date,omitempty"`
	// Protected: display the server protection option is activated
	Protected bool `json:"protected,omitempty"`
	// SecurityGroup: display the server security group
	SecurityGroup *SecurityGroupSummary `json:"security_group,omitempty"`
	// StateDetail: display the server state_detail
	StateDetail string `json:"state_detail,omitempty"`
}

type ServerActionResponse struct {
	Task *Task `json:"task,omitempty"`
}

type ServerIp struct {
	// Id: display the unique ID of the IP address
	Id string `json:"id,omitempty"`
	// Address: display the server public IPv4 IP-Address
	Address net.IP `json:"address,omitempty"`
	// Dynamic: display information if the IP address will be considered as dynamic
	Dynamic bool `json:"dynamic,omitempty"`
}

type ServerIpv6 struct {
	// Address: display the server IPv6 IP-Address
	Address net.IP `json:"address,omitempty"`
	// Gateway: display the IPv6 IP-addresses gateway
	Gateway string `json:"gateway,omitempty"`
	// Netmask: display the IPv6 IP-addresses CIDR netmask
	Netmask string `json:"netmask,omitempty"`
}

type ServerLocation struct {
	ClusterId string `json:"cluster_id,omitempty"`

	HypervisorId string `json:"hypervisor_id,omitempty"`

	NodeId string `json:"node_id,omitempty"`

	PlatformId string `json:"platform_id,omitempty"`

	ZoneId string `json:"zone_id,omitempty"`
}

type ServerMaintenance struct {
}

type ServerSummary struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

type ServerTypeDefinition struct {
	MonthlyPrice float32 `json:"monthly_price,omitempty"`

	HourlyPrice float32 `json:"hourly_price,omitempty"`

	AltNames map[uint32]string `json:"alt_names,omitempty"`

	PerVolumeConstraint map[string]*ServerTypeDefinitionVolumeConstraintSizes `json:"per_volume_constraint,omitempty"`

	VolumesConstraint *ServerTypeDefinitionVolumeConstraintSizes `json:"volumes_constraint,omitempty"`

	Ncpus uint32 `json:"ncpus,omitempty"`

	Gpu *uint64 `json:"gpu,omitempty"`

	Ram uint64 `json:"ram,omitempty"`

	Arch Arch `json:"arch,omitempty"`

	Baremetal bool `json:"baremetal,omitempty"`

	Network *ServerTypeDefinitionNetwork `json:"network,omitempty"`
}

type ServerTypeDefinitionNetwork struct {
	Interfaces []*ServerTypeDefinitionNetworkInterface `json:"interfaces,omitempty"`

	SumInternalBandwidth *uint64 `json:"sum_internal_bandwidth,omitempty"`

	SumInternetBandwidth *uint64 `json:"sum_internet_bandwidth,omitempty"`

	Ipv6Support bool `json:"ipv6_support,omitempty"`
}

type ServerTypeDefinitionNetworkInterface struct {
	InternalBandwidth *uint64 `json:"internal_bandwidth,omitempty"`

	InternetBandwidth *uint64 `json:"internet_bandwidth,omitempty"`
}

type ServerTypeDefinitionVolumeConstraintSizes struct {
	MinSize uint64 `json:"min_size,omitempty"`

	MaxSize uint64 `json:"max_size,omitempty"`
}

type SetImageResponse struct {
	Image *Image `json:"image,omitempty"`
}

type SetIpResponse struct {
	Ip *Ip `json:"ip,omitempty"`
}

type SetServerResponse struct {
	Server *Server `json:"server,omitempty"`
}

type SetSnapshotResponse struct {
	Snapshot *Snapshot `json:"snapshot,omitempty"`
}

type SetVolumeResponse struct {
	Volume *Volume `json:"volume,omitempty"`
}

type Snapshot struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Organization string `json:"organization,omitempty"`

	VolumeType VolumeType `json:"volume_type,omitempty"`

	Size uint64 `json:"size,omitempty"`

	State SnapshotState `json:"state,omitempty"`

	BaseVolume *SnapshotBaseVolume `json:"base_volume,omitempty"`

	CreationDate time.Time `json:"creation_date,omitempty"`

	ModificationDate time.Time `json:"modification_date,omitempty"`
}

type SnapshotBaseVolume struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
}

type Task struct {
	// Id: the unique ID of the task
	Id string `json:"id,omitempty"`
	// Description: the description of the task
	Description string `json:"description,omitempty"`

	HrefFrom string `json:"href_from,omitempty"`

	HrefResult string `json:"href_result,omitempty"`
	// Progress: show the progress of the task in percent
	Progress int32 `json:"progress,omitempty"`
	// StartedAt: display the task start date
	StartedAt time.Time `json:"started_at,omitempty"`
	// Status: display the task status
	Status TaskStatus `json:"status,omitempty"`
	// TerminatedAt: display the task end date
	TerminatedAt time.Time `json:"terminated_at,omitempty"`
}

type UpdateIpResponse struct {
	Ip *Ip `json:"ip,omitempty"`
}

type UpdateSecurityGroupResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group,omitempty"`
}

type UpdateServerResponse struct {
	Server *Server `json:"server,omitempty"`
}

type Volume struct {
	// Id: display the volumes unique ID
	Id string `json:"id,omitempty"`
	// Name: display the volumes names
	Name string `json:"name,omitempty"`
	// ExportUri: show the volumes NBD export URI
	ExportUri string `json:"export_uri,omitempty"`
	// Organization: display the volumes organization
	Organization string `json:"organization,omitempty"`
	// Server: display information about the server attached to the volume
	Server *ServerSummary `json:"server,omitempty"`
	// Size: display the volumes disk size
	Size uint64 `json:"size,omitempty"`
	// VolumeType: display the volumes type
	VolumeType VolumeType `json:"volume_type,omitempty"`
	// CreationDate: display the volumes creation date
	CreationDate time.Time `json:"creation_date,omitempty"`
	// ModificationDate: display the volumes modification date
	ModificationDate time.Time `json:"modification_date,omitempty"`
	// State: display the volumes state
	State VolumeState `json:"state,omitempty"`
}

type VolumeTemplate struct {
	// Id: display the volumes unique ID
	Id string `json:"id,omitempty"`
	// Name: display the volumes name
	Name string `json:"name,omitempty"`
	// Size: display the volumes disk size
	Size uint64 `json:"size,omitempty"`
	// VolumeType: display the volumes type
	VolumeType VolumeType `json:"volume_type,omitempty"`
	// Organization: the organization ID
	Organization string `json:"organization,omitempty"`
}

// Service Api

type GetServerTypesAvailabilityRequest struct {
	Zone scw.Zone `json:"-"`

	PerPage *int32 `json:"-"`

	Page *int32 `json:"-"`
}

// GetServerTypesAvailability: get availability
//
// Get availibility for all server types
func (s *Api) GetServerTypesAvailability(req *GetServerTypesAvailabilityRequest) (*GetServerTypesAvailabilityResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}
	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/products/servers/availability",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetServerTypesAvailabilityResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServersTypesRequest struct {
	Zone scw.Zone `json:"-"`

	PerPage *int32 `json:"-"`

	Page *int32 `json:"-"`
}

// ListServersTypes: list server types
//
// Get server types technical details
func (s *Api) ListServersTypes(req *ListServersTypesRequest) (*ListServersTypesResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}
	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/products/servers",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListServersTypesResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServersRequest struct {
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`

	PerPage *int32 `json:"-"`

	Page *int32 `json:"-"`
}

// ListServers: list servers
func (s *Api) ListServers(req *ListServersRequest) (*ListServersResponse, error) {
	var err error

	val := s.client.GetDefaultOrganizationID()
	if req.Organization == nil || *req.Organization == "" {
		req.Organization = &val
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}
	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListServersResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateServerRequest struct {
	Zone scw.Zone `json:"-"`
	// Name: display the server name
	Name string `json:"name,omitempty"`
	// DynamicIpRequired: define if a dynamic IP is required for the instance
	DynamicIpRequired bool `json:"dynamic_ip_required,omitempty"`
	// CommercialType: define the server commercial type (i.e. GP1-S)
	CommercialType string `json:"commercial_type,omitempty"`
	// Image: define the server image id
	Image string `json:"image,omitempty"`
	// Volumes: define the volumes attached to the server
	Volumes map[string]*VolumeTemplate `json:"volumes,omitempty"`
	// EnableIpv6: define if IPv6 is enabled on the server
	EnableIpv6 bool `json:"enable_ipv6,omitempty"`
	// PublicIp: define the public IPv4 attached to the server
	PublicIp string `json:"public_ip,omitempty"`
	// BootType: define the boot type you want to use
	BootType ServerBootType `json:"boot_type,omitempty"`
	// Organization: define the server organization
	Organization string `json:"organization,omitempty"`
}

// CreateServer: create server
func (s *Api) CreateServer(req *CreateServerRequest) (*CreateServerResponse, error) {
	var err error

	if req.Organization == "" {
		req.Organization = s.client.GetDefaultOrganizationID()
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp CreateServerResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteServerRequest struct {
	Zone scw.Zone `json:"-"`

	ServerId string `json:"-"`
}

// DeleteServer: delete server
//
// Delete a server with the given id
func (s *Api) DeleteServer(req *DeleteServerRequest) error {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerId) + "",
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq)

	if err != nil {
		return err
	}
	return nil
}

type GetServerRequest struct {
	Zone scw.Zone `json:"-"`

	ServerId string `json:"-"`
}

// GetServer: get server
//
// Get the details of a specified Server
func (s *Api) GetServer(req *GetServerRequest) (*GetServerResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetServerResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetServerRequest struct {
	Zone scw.Zone `json:"-"`
	// Id: display the server unique ID
	Id string `json:"-"`
	// Name: display the server name
	Name string `json:"name,omitempty"`
	// Organization: display the server organization
	Organization string `json:"organization,omitempty"`
	// AllowedActions: provide as list of allowed actions on the server
	AllowedActions []ServerAction `json:"allowed_actions,omitempty"`
	// Tags: display the server associated tags
	Tags []string `json:"tags,omitempty"`
	// CommercialType: display the server commercial type (e.g. GP1-M)
	CommercialType string `json:"commercial_type,omitempty"`
	// CreationDate: display the server creation date
	CreationDate time.Time `json:"creation_date,omitempty"`
	// DynamicIpRequired: display if a dynamic IP is required
	DynamicIpRequired bool `json:"dynamic_ip_required,omitempty"`
	// DynamicPublicIp: display the server dynamic public IP
	DynamicPublicIp bool `json:"dynamic_public_ip,omitempty"`
	// EnableIpv6: display if IPv6 is enabled
	EnableIpv6 bool `json:"enable_ipv6,omitempty"`
	// ExtraNetworks: display information about additional network interfaces
	ExtraNetworks []string `json:"extra_networks,omitempty"`
	// Hostname: display the server host name
	Hostname string `json:"hostname,omitempty"`
	// Image: provide information on the server image
	Image *Image `json:"image,omitempty"`
	// Protected: display the server protection option is activated
	Protected bool `json:"protected,omitempty"`
	// PrivateIp: display the server private IP address
	PrivateIp *string `json:"private_ip,omitempty"`
	// PublicIp: display the server public IP address
	PublicIp *ServerIp `json:"public_ip,omitempty"`
	// ModificationDate: display the server modification date
	ModificationDate time.Time `json:"modification_date,omitempty"`
	// State: display the server state
	State ServerState `json:"state,omitempty"`
	// Location: display the server location
	Location *ServerLocation `json:"location,omitempty"`
	// Ipv6: display the server IPv6 address
	Ipv6 *ServerIpv6 `json:"ipv6,omitempty"`
	// Bootscript: display the server bootscript
	Bootscript *Bootscript `json:"bootscript,omitempty"`
	// BootType: display the server boot type
	BootType ServerBootType `json:"boot_type,omitempty"`
	// Volumes: display the server volumes
	Volumes map[string]*Volume `json:"volumes,omitempty"`
	// SecurityGroup: display the server security group
	SecurityGroup *SecurityGroupSummary `json:"security_group,omitempty"`
	// Maintenances: display the server planned maintenances
	Maintenances []*ServerMaintenance `json:"maintenances,omitempty"`
	// StateDetail: display the server state_detail
	StateDetail string `json:"state_detail,omitempty"`
	// Arch: display the server arch
	Arch Arch `json:"arch,omitempty"`
}

func (s *Api) SetServer(req *SetServerRequest) (*SetServerResponse, error) {
	var err error

	if req.Organization == "" {
		req.Organization = s.client.GetDefaultOrganizationID()
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.Id) + "",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp SetServerResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateServerRequest struct {
	Zone scw.Zone `json:"-"`

	ServerId string `json:"-"`

	Name *string `json:"name,omitempty"`

	BootType ServerBootType `json:"boot_type,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	Volumes map[string]*VolumeTemplate `json:"volumes,omitempty"`

	Bootscript *Bootscript `json:"bootscript,omitempty"`

	DynamicIpRequired *bool `json:"dynamic_ip_required,omitempty"`

	EnableIpv6 *bool `json:"enable_ipv6,omitempty"`

	ExtraNetworks *[]string `json:"extra_networks,omitempty"`

	Protected bool `json:"protected,omitempty"`

	SecurityGroup *SecurityGroupSummary `json:"security_group,omitempty"`
}

// UpdateServer: update server
func (s *Api) UpdateServer(req *UpdateServerRequest) (*UpdateServerResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerId) + "",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp UpdateServerResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServerActionsRequest struct {
	Zone scw.Zone `json:"-"`

	ServerId string `json:"-"`
}

// ListServerActions: list server actions
//
// Liste all actions that can currently be performed on a server
func (s *Api) ListServerActions(req *ListServerActionsRequest) (*ListServerActionsResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerId) + "/action",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListServerActionsResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ServerActionRequest struct {
	Zone scw.Zone `json:"-"`

	ServerId string `json:"-"`

	Action ServerAction `json:"action,omitempty"`
}

// ServerAction: perform action
//
// Perform power related actions on a server
func (s *Api) ServerAction(req *ServerActionRequest) (*ServerActionResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerId) + "/action",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ServerActionResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServerUserDataRequest struct {
	Zone scw.Zone `json:"-"`

	ServerId string `json:"-"`
}

// ListServerUserData: list user data
//
// List all user data keys register on a given server
func (s *Api) ListServerUserData(req *ListServerUserDataRequest) (*ListServerUserDataResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerId) + "/user_data",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListServerUserDataResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteServerUserDataRequest struct {
	Zone scw.Zone `json:"-"`

	ServerId string `json:"-"`

	Key string `json:"-"`
}

// DeleteServerUserData: delete user data
//
// Delete the given key from a server user data
func (s *Api) DeleteServerUserData(req *DeleteServerUserDataRequest) error {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerId) + "/user_data/" + fmt.Sprint(req.Key) + "",
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq)

	if err != nil {
		return err
	}
	return nil
}

type SetServerUserDataRequest struct {
	Zone scw.Zone `json:"-"`

	ServerId string `json:"-"`

	Key string `json:"-"`

	Content *utils.File
}

// SetServerUserData: add/Set user data
//
// Add or update a user data with the given key on a server
func (s *Api) SetServerUserData(req *SetServerUserDataRequest) error {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerId) + "/user_data/" + fmt.Sprint(req.Key) + "",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req.Content)
	if err != nil {
		return err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	_, err = s.client.Do(scwReq)

	if err != nil {
		return err
	}
	return nil
}

type GetServerUserDataRequest struct {
	Zone scw.Zone `json:"-"`

	ServerId string `json:"-"`

	Key string `json:"-"`
}

// GetServerUserData: get user data
//
// Get the content of a user data with the given key on a server
func (s *Api) GetServerUserData(req *GetServerUserDataRequest) (*utils.File, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerId) + "/user_data/" + fmt.Sprint(req.Key) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp utils.File
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListImagesRequest struct {
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`

	PerPage *int32 `json:"-"`

	Page *int32 `json:"-"`

	Name *string `json:"-"`

	Public bool `json:"-"`

	Arch *string `json:"-"`
}

// ListImages: list images
//
// List all images available in an account
func (s *Api) ListImages(req *ListImagesRequest) (*ListImagesResponse, error) {
	var err error

	val := s.client.GetDefaultOrganizationID()
	if req.Organization == nil || *req.Organization == "" {
		req.Organization = &val
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}
	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "public", req.Public)
	parameter.AddToQuery(query, "arch", req.Arch)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListImagesResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetImageRequest struct {
	Zone scw.Zone `json:"-"`

	ImageId string `json:"-"`
}

// GetImage: get image
//
// Get details of an image with the given id
func (s *Api) GetImage(req *GetImageRequest) (*GetImageResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.ImageId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetImageResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateImageRequest struct {
	Zone scw.Zone `json:"-"`

	Name string `json:"name,omitempty"`

	Arch Arch `json:"arch,omitempty"`

	DefaultBootscript *Bootscript `json:"default_bootscript,omitempty"`

	ExtraVolumes map[string]*Volume `json:"extra_volumes,omitempty"`

	Organization string `json:"organization,omitempty"`

	Public bool `json:"public,omitempty"`
}

// CreateImage: create image
func (s *Api) CreateImage(req *CreateImageRequest) (*CreateImageResponse, error) {
	var err error

	if req.Organization == "" {
		req.Organization = s.client.GetDefaultOrganizationID()
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp CreateImageResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetImageRequest struct {
	Zone scw.Zone `json:"-"`

	Id string `json:"-"`

	Name string `json:"name,omitempty"`

	Arch Arch `json:"arch,omitempty"`

	CreationDate time.Time `json:"creation_date,omitempty"`

	ModificationDate time.Time `json:"modification_date,omitempty"`

	DefaultBootscript *Bootscript `json:"default_bootscript,omitempty"`

	ExtraVolumes map[string]*Volume `json:"extra_volumes,omitempty"`

	FromServer *ServerSummary `json:"from_server,omitempty"`

	Organization string `json:"organization,omitempty"`

	Public bool `json:"public,omitempty"`

	RootVolume *VolumeTemplate `json:"root_volume,omitempty"`

	State ImageState `json:"state,omitempty"`
}

// SetImage: update image
//
// Replace all image properties with an image message
func (s *Api) SetImage(req *SetImageRequest) (*SetImageResponse, error) {
	var err error

	if req.Organization == "" {
		req.Organization = s.client.GetDefaultOrganizationID()
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.Id) + "",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp SetImageResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteImageRequest struct {
	Zone scw.Zone `json:"-"`

	ImageId string `json:"-"`
}

// DeleteImage: delete image
//
// Delete the image with the given id
func (s *Api) DeleteImage(req *DeleteImageRequest) error {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/images/" + fmt.Sprint(req.ImageId) + "",
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq)

	if err != nil {
		return err
	}
	return nil
}

type ListSnapshotsRequest struct {
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`

	PerPage *int32 `json:"-"`

	Page *int32 `json:"-"`

	Name *string `json:"-"`
}

// ListSnapshots: list snapshots
func (s *Api) ListSnapshots(req *ListSnapshotsRequest) (*ListSnapshotsResponse, error) {
	var err error

	val := s.client.GetDefaultOrganizationID()
	if req.Organization == nil || *req.Organization == "" {
		req.Organization = &val
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}
	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "name", req.Name)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListSnapshotsResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateSnapshotRequest struct {
	Zone scw.Zone `json:"-"`

	VolumeId string `json:"volume_id,omitempty"`

	Organization string `json:"organization,omitempty"`

	Name string `json:"name,omitempty"`
}

// CreateSnapshot: create snapshot
func (s *Api) CreateSnapshot(req *CreateSnapshotRequest) (*CreateSnapshotResponse, error) {
	var err error

	if req.Organization == "" {
		req.Organization = s.client.GetDefaultOrganizationID()
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp CreateSnapshotResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetSnapshotRequest struct {
	Zone scw.Zone `json:"-"`

	SnapshotId string `json:"-"`
}

// GetSnapshot: get snapshot
//
// Get details of a snapshot with the given id
func (s *Api) GetSnapshot(req *GetSnapshotRequest) (*GetSnapshotResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetSnapshotResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetSnapshotRequest struct {
	Zone scw.Zone `json:"-"`

	Id string `json:"-"`

	Name string `json:"name,omitempty"`

	Organization string `json:"organization,omitempty"`

	VolumeType VolumeType `json:"volume_type,omitempty"`

	Size uint64 `json:"size,omitempty"`

	State SnapshotState `json:"state,omitempty"`

	BaseVolume *SnapshotBaseVolume `json:"base_volume,omitempty"`

	CreationDate time.Time `json:"creation_date,omitempty"`

	ModificationDate time.Time `json:"modification_date,omitempty"`
}

// SetSnapshot: update snapshot
//
// Replace all snapshot properties with a snapshot message
func (s *Api) SetSnapshot(req *SetSnapshotRequest) (*SetSnapshotResponse, error) {
	var err error

	if req.Organization == "" {
		req.Organization = s.client.GetDefaultOrganizationID()
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.Id) + "",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp SetSnapshotResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteSnapshotRequest struct {
	Zone scw.Zone `json:"-"`

	SnapshotId string `json:"-"`
}

// DeleteSnapshot: delete snapshot
//
// Delete the snapshot with the given id
func (s *Api) DeleteSnapshot(req *DeleteSnapshotRequest) error {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/snapshots/" + fmt.Sprint(req.SnapshotId) + "",
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq)

	if err != nil {
		return err
	}
	return nil
}

type ListVolumesRequest struct {
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`

	PerPage *int32 `json:"-"`

	Page *int32 `json:"-"`

	Name *string `json:"-"`
}

// ListVolumes: list volumes
func (s *Api) ListVolumes(req *ListVolumesRequest) (*ListVolumesResponse, error) {
	var err error

	val := s.client.GetDefaultOrganizationID()
	if req.Organization == nil || *req.Organization == "" {
		req.Organization = &val
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}
	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "name", req.Name)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListVolumesResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateVolumeRequest struct {
	Zone scw.Zone `json:"-"`

	Name string `json:"name,omitempty"`

	Organization string `json:"organization,omitempty"`

	VolumeType VolumeType `json:"volume_type,omitempty"`

	// Precisely one of BaseSnapshot, BaseVolume, Size must be set.
	Size *uint64 `json:"size,omitempty"`

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

// CreateVolume: create volume
func (s *Api) CreateVolume(req *CreateVolumeRequest) (*CreateVolumeResponse, error) {
	var err error

	if req.Organization == "" {
		req.Organization = s.client.GetDefaultOrganizationID()
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp CreateVolumeResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetVolumeRequest struct {
	Zone scw.Zone `json:"-"`

	VolumeId string `json:"-"`
}

// GetVolume: get volume
//
// Get details of a volume with the given id
func (s *Api) GetVolume(req *GetVolumeRequest) (*GetVolumeResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetVolumeResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetVolumeRequest struct {
	Zone scw.Zone `json:"-"`
	// Id: display the volumes unique ID
	Id string `json:"-"`
	// Name: display the volumes names
	Name string `json:"name,omitempty"`
	// ExportUri: show the volumes NBD export URI
	ExportUri string `json:"export_uri,omitempty"`
	// Size: display the volumes disk size
	Size uint64 `json:"size,omitempty"`
	// VolumeType: display the volumes type
	VolumeType VolumeType `json:"volume_type,omitempty"`
	// CreationDate: display the volumes creation date
	CreationDate time.Time `json:"creation_date,omitempty"`
	// ModificationDate: display the volumes modification date
	ModificationDate time.Time `json:"modification_date,omitempty"`
	// Organization: display the volumes organization
	Organization string `json:"organization,omitempty"`
	// Server: display information about the server attached to the volume
	Server *ServerSummary `json:"server,omitempty"`
	// State: display the volumes state
	State VolumeState `json:"state,omitempty"`
}

// SetVolume: update volume
//
// Replace all volume properties with a volume message
func (s *Api) SetVolume(req *SetVolumeRequest) (*SetVolumeResponse, error) {
	var err error

	if req.Organization == "" {
		req.Organization = s.client.GetDefaultOrganizationID()
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.Id) + "",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp SetVolumeResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteVolumeRequest struct {
	Zone scw.Zone `json:"-"`

	VolumeId string `json:"-"`
}

// DeleteVolume: delete volume
//
// Delete the volume with the given id
func (s *Api) DeleteVolume(req *DeleteVolumeRequest) error {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/volumes/" + fmt.Sprint(req.VolumeId) + "",
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq)

	if err != nil {
		return err
	}
	return nil
}

type ListSecurityGroupsRequest struct {
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`

	PerPage *int32 `json:"-"`

	Page *int32 `json:"-"`
}

// ListSecurityGroups: list security groups
//
// List all security groups available in an account
func (s *Api) ListSecurityGroups(req *ListSecurityGroupsRequest) (*ListSecurityGroupsResponse, error) {
	var err error

	val := s.client.GetDefaultOrganizationID()
	if req.Organization == nil || *req.Organization == "" {
		req.Organization = &val
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}
	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListSecurityGroupsResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateSecurityGroupRequest struct {
	Zone scw.Zone `json:"-"`

	Name string `json:"name,omitempty"`

	OrganizationKey string `json:"organization_key,omitempty"`

	OrganizationDefault bool `json:"organization_default,omitempty"`

	Stateful bool `json:"stateful,omitempty"`

	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy,omitempty"`

	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy,omitempty"`
}

// CreateSecurityGroup: create security group
func (s *Api) CreateSecurityGroup(req *CreateSecurityGroupRequest) (*CreateSecurityGroupResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp CreateSecurityGroupResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetSecurityGroupRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupId string `json:"-"`
}

// GetSecurityGroup: get security group
//
// Get the details of a Security Group with the given id
func (s *Api) GetSecurityGroup(req *GetSecurityGroupRequest) (*GetSecurityGroupResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetSecurityGroupResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteSecurityGroupRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupId string `json:"-"`
}

// DeleteSecurityGroup: delete security group
func (s *Api) DeleteSecurityGroup(req *DeleteSecurityGroupRequest) error {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupId) + "",
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq)

	if err != nil {
		return err
	}
	return nil
}

type SetSecurityGroupRequest struct {
	Zone scw.Zone `json:"-"`
	// Id: display the security groups' unique ID
	Id string `json:"-"`
	// Name: display the security groups name
	Name string `json:"name,omitempty"`
	// Description: display the security groups description
	Description string `json:"description,omitempty"`
	// EnableDefaultSecurity: display if the security group is set as default
	EnableDefaultSecurity bool `json:"enable_default_security,omitempty"`
	// InboundDefaultPolicy: display the default inbound policy
	InboundDefaultPolicy SecurityGroupPolicy `json:"inbound_default_policy,omitempty"`
	// OutboundDefaultPolicy: display the default outbound policy
	OutboundDefaultPolicy SecurityGroupPolicy `json:"outbound_default_policy,omitempty"`
	// Organization: display the security groups organization ID
	Organization string `json:"organization,omitempty"`
	// OrganizationDefault: display if the security group is set as organization default
	OrganizationDefault bool `json:"organization_default,omitempty"`
	// CreationDate: display the security group creation date
	CreationDate time.Time `json:"creation_date,omitempty"`
	// ModificationDate: display the security group modification date
	ModificationDate time.Time `json:"modification_date,omitempty"`
	// Servers: list of servers attached to this security group
	Servers map[string]*ServerSummary `json:"servers,omitempty"`
	// Stateful: true if the security group is stateful
	Stateful bool `json:"stateful,omitempty"`
}

// SetSecurityGroup: update security group
//
// Replace all security group properties with a security group message
func (s *Api) SetSecurityGroup(req *SetSecurityGroupRequest) (*UpdateSecurityGroupResponse, error) {
	var err error

	if req.Organization == "" {
		req.Organization = s.client.GetDefaultOrganizationID()
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.Id) + "",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp UpdateSecurityGroupResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListSecurityGroupRulesRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupId string `json:"-"`

	PerPage *int32 `json:"-"`

	Page *int32 `json:"-"`
}

// ListSecurityGroupRules: list rules
func (s *Api) ListSecurityGroupRules(req *ListSecurityGroupRulesRequest) (*ListSecurityGroupRulesResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}
	query := url.Values{}
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupId) + "/rules",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListSecurityGroupRulesResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateSecurityGroupRuleRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupId string `json:"-"`

	Protocol SecurityRuleProtocol `json:"protocol,omitempty"`

	Direction SecurityRuleDirection `json:"direction,omitempty"`

	Action SecurityRuleAction `json:"action,omitempty"`

	IpRange string `json:"ip_range,omitempty"`

	DestPortFrom uint32 `json:"dest_port_from,omitempty"`

	DestPortTo uint32 `json:"dest_port_to,omitempty"`

	Position uint32 `json:"position,omitempty"`

	Editable bool `json:"editable,omitempty"`
}

// CreateSecurityGroupRule: create rule
func (s *Api) CreateSecurityGroupRule(req *CreateSecurityGroupRuleRequest) (*CreateSecurityGroupRuleResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupId) + "/rules",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp CreateSecurityGroupRuleResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteSecurityGroupRuleRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupId string `json:"-"`

	SecurityRuleId string `json:"-"`
}

// DeleteSecurityGroupRule: delete rule
//
// Delete a security group rule with the given id
func (s *Api) DeleteSecurityGroupRule(req *DeleteSecurityGroupRuleRequest) error {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupId) + "/rules/" + fmt.Sprint(req.SecurityRuleId) + "",
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq)

	if err != nil {
		return err
	}
	return nil
}

type GetSecurityGroupRuleRequest struct {
	Zone scw.Zone `json:"-"`

	SecurityGroupId string `json:"-"`

	SecurityRuleId string `json:"-"`
}

// GetSecurityGroupRule: get rule
//
// Get details of a security group rule with the given id
func (s *Api) GetSecurityGroupRule(req *GetSecurityGroupRuleRequest) (*GetSecurityGroupRuleResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/security_groups/" + fmt.Sprint(req.SecurityGroupId) + "/rules/" + fmt.Sprint(req.SecurityRuleId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetSecurityGroupRuleResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListIpsRequest struct {
	Zone scw.Zone `json:"-"`

	Organization string `json:"-"`

	Name *string `json:"-"`
}

// ListIps: list IPs
func (s *Api) ListIps(req *ListIpsRequest) (*ListIpsResponse, error) {
	var err error

	if req.Organization == "" {
		req.Organization = s.client.GetDefaultOrganizationID()
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}
	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "name", req.Name)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListIpsResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateIpRequest struct {
	Zone scw.Zone `json:"-"`

	Organization string `json:"organization,omitempty"`

	Server *string `json:"server,omitempty"`
}

// CreateIp: reseve an IP
func (s *Api) CreateIp(req *CreateIpRequest) (*CreateIpResponse, error) {
	var err error

	if req.Organization == "" {
		req.Organization = s.client.GetDefaultOrganizationID()
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp CreateIpResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetIpRequest struct {
	Zone scw.Zone `json:"-"`

	IpId string `json:"-"`
}

// GetIp: get IP
//
// Get details of an IP with the given id
func (s *Api) GetIp(req *GetIpRequest) (*GetIpResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IpId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetIpResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetIpRequest struct {
	Zone scw.Zone `json:"-"`

	Id string `json:"-"`

	Address net.IP `json:"address,omitempty"`

	Reverse string `json:"reverse,omitempty"`

	Server string `json:"server,omitempty"`

	Organization string `json:"organization,omitempty"`
}

func (s *Api) SetIp(req *SetIpRequest) (*SetIpResponse, error) {
	var err error

	if req.Organization == "" {
		req.Organization = s.client.GetDefaultOrganizationID()
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.Id) + "",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp SetIpResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateIpRequest struct {
	Zone scw.Zone `json:"-"`

	IpId string `json:"-"`

	Reverse *string `json:"reverse,omitempty"`

	Server *string `json:"server,omitempty"`
}

// UpdateIp: update IP
func (s *Api) UpdateIp(req *UpdateIpRequest) (*UpdateIpResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IpId) + "",
		Headers: http.Header{},
	}
	body, err := marshaler.MarshalBody(req)
	if err != nil {
		return nil, err
	}
	scwReq.Headers.Add("Content-Type", body.ContentType())
	scwReq.Body = body

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp UpdateIpResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteIpRequest struct {
	Zone scw.Zone `json:"-"`

	IpId string `json:"-"`
}

// DeleteIp: delete IP
//
// Delete the IP with the given id
func (s *Api) DeleteIp(req *DeleteIpRequest) error {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IpId) + "",
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq)

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
}

// ListBootscripts: list bootscripts
func (s *Api) ListBootscripts(req *ListBootscriptsRequest) (*ListBootscriptsResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}
	query := url.Values{}
	parameter.AddToQuery(query, "arch", req.Arch)
	parameter.AddToQuery(query, "title", req.Title)
	parameter.AddToQuery(query, "default", req.Default)
	parameter.AddToQuery(query, "public", req.Public)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/bootscripts",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListBootscriptsResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetBootscriptRequest struct {
	Zone scw.Zone `json:"-"`

	BootscriptId string `json:"-"`
}

// GetBootscript: get bootscripts
//
// Get details of a bootscript with the given id
func (s *Api) GetBootscript(req *GetBootscriptRequest) (*GetBootscriptResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/bootscripts/" + fmt.Sprint(req.BootscriptId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetBootscriptResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetServiceInfoRequest struct {
	Zone scw.Zone `json:"-"`
}

func (s *Api) GetServiceInfo(req *GetServiceInfoRequest) (*GetServiceInfoResponse, error) {
	var err error

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetServiceInfoResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDashboardRequest struct {
	Zone scw.Zone `json:"-"`

	Organization *string `json:"-"`
}

func (s *Api) GetDashboard(req *GetDashboardRequest) (*GetDashboardResponse, error) {
	var err error

	val := s.client.GetDefaultOrganizationID()
	if req.Organization == nil || *req.Organization == "" {
		req.Organization = &val
	}

	if req.Zone == "" {
		req.Zone = s.client.GetDefaultZone()
	}
	query := url.Values{}
	parameter.AddToQuery(query, "organization", req.Organization)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/instance/v1/zones/" + fmt.Sprint(req.Zone) + "/dashboard",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp GetDashboardResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type From interface {
	isFrom()
}

type FromSize struct {
	Value uint64
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

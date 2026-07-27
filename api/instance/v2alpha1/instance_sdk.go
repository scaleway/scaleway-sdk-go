// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package instance provides methods and message types of the instance v2alpha1 API.
package instance

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
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultInstanceRetryInterval = 15 * time.Second
	defaultInstanceTimeout       = 5 * time.Minute
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

type CreateServerRequestBookIPIPType string

const (
	CreateServerRequestBookIPIPTypeUnknownIPType = CreateServerRequestBookIPIPType("unknown_ip_type")
	CreateServerRequestBookIPIPTypeZonalIPv4     = CreateServerRequestBookIPIPType("zonal_ipv4")
	CreateServerRequestBookIPIPTypeZonalIPv6     = CreateServerRequestBookIPIPType("zonal_ipv6")
)

func (enum CreateServerRequestBookIPIPType) String() string {
	if enum == "" {
		// return default value if empty
		return string(CreateServerRequestBookIPIPTypeUnknownIPType)
	}
	return string(enum)
}

func (enum CreateServerRequestBookIPIPType) Values() []CreateServerRequestBookIPIPType {
	return []CreateServerRequestBookIPIPType{
		"unknown_ip_type",
		"zonal_ipv4",
		"zonal_ipv6",
	}
}

func (enum CreateServerRequestBookIPIPType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CreateServerRequestBookIPIPType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CreateServerRequestBookIPIPType(CreateServerRequestBookIPIPType(tmp).String())
	return nil
}

type CreateServerRequestServerVolumeVolumeType string

const (
	CreateServerRequestServerVolumeVolumeTypeUnknownVolumeType = CreateServerRequestServerVolumeVolumeType("unknown_volume_type")
	CreateServerRequestServerVolumeVolumeTypeLSSD              = CreateServerRequestServerVolumeVolumeType("l_ssd")
	CreateServerRequestServerVolumeVolumeTypeSbs               = CreateServerRequestServerVolumeVolumeType("sbs")
	CreateServerRequestServerVolumeVolumeTypeScratch           = CreateServerRequestServerVolumeVolumeType("scratch")
)

func (enum CreateServerRequestServerVolumeVolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return string(CreateServerRequestServerVolumeVolumeTypeUnknownVolumeType)
	}
	return string(enum)
}

func (enum CreateServerRequestServerVolumeVolumeType) Values() []CreateServerRequestServerVolumeVolumeType {
	return []CreateServerRequestServerVolumeVolumeType{
		"unknown_volume_type",
		"l_ssd",
		"sbs",
		"scratch",
	}
}

func (enum CreateServerRequestServerVolumeVolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CreateServerRequestServerVolumeVolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CreateServerRequestServerVolumeVolumeType(CreateServerRequestServerVolumeVolumeType(tmp).String())
	return nil
}

type ListPlacementGroupsRequestOrderBy string

const (
	// Created at descending.
	ListPlacementGroupsRequestOrderByCreatedAtDesc = ListPlacementGroupsRequestOrderBy("created_at_desc")
	// Created at ascending.
	ListPlacementGroupsRequestOrderByCreatedAtAsc = ListPlacementGroupsRequestOrderBy("created_at_asc")
	// Updated at descending.
	ListPlacementGroupsRequestOrderByUpdatedAtDesc = ListPlacementGroupsRequestOrderBy("updated_at_desc")
	// Updated at ascending.
	ListPlacementGroupsRequestOrderByUpdatedAtAsc = ListPlacementGroupsRequestOrderBy("updated_at_asc")
)

func (enum ListPlacementGroupsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPlacementGroupsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListPlacementGroupsRequestOrderBy) Values() []ListPlacementGroupsRequestOrderBy {
	return []ListPlacementGroupsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"updated_at_desc",
		"updated_at_asc",
	}
}

func (enum ListPlacementGroupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPlacementGroupsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPlacementGroupsRequestOrderBy(ListPlacementGroupsRequestOrderBy(tmp).String())
	return nil
}

type ListPrivateNetworkInterfacesRequestOrderBy string

const (
	// Created at descending.
	ListPrivateNetworkInterfacesRequestOrderByCreatedAtDesc = ListPrivateNetworkInterfacesRequestOrderBy("created_at_desc")
	// Created at ascending.
	ListPrivateNetworkInterfacesRequestOrderByCreatedAtAsc = ListPrivateNetworkInterfacesRequestOrderBy("created_at_asc")
	// Updated at descending.
	ListPrivateNetworkInterfacesRequestOrderByUpdatedAtDesc = ListPrivateNetworkInterfacesRequestOrderBy("updated_at_desc")
	// Updated at ascending.
	ListPrivateNetworkInterfacesRequestOrderByUpdatedAtAsc = ListPrivateNetworkInterfacesRequestOrderBy("updated_at_asc")
)

func (enum ListPrivateNetworkInterfacesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPrivateNetworkInterfacesRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListPrivateNetworkInterfacesRequestOrderBy) Values() []ListPrivateNetworkInterfacesRequestOrderBy {
	return []ListPrivateNetworkInterfacesRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"updated_at_desc",
		"updated_at_asc",
	}
}

func (enum ListPrivateNetworkInterfacesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPrivateNetworkInterfacesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPrivateNetworkInterfacesRequestOrderBy(ListPrivateNetworkInterfacesRequestOrderBy(tmp).String())
	return nil
}

type ListSecurityGroupsRequestOrderBy string

const (
	// Created at descending.
	ListSecurityGroupsRequestOrderByCreatedAtDesc = ListSecurityGroupsRequestOrderBy("created_at_desc")
	// Created at ascending.
	ListSecurityGroupsRequestOrderByCreatedAtAsc = ListSecurityGroupsRequestOrderBy("created_at_asc")
	// Updated at descending.
	ListSecurityGroupsRequestOrderByUpdatedAtDesc = ListSecurityGroupsRequestOrderBy("updated_at_desc")
	// Updated at ascending.
	ListSecurityGroupsRequestOrderByUpdatedAtAsc = ListSecurityGroupsRequestOrderBy("updated_at_asc")
)

func (enum ListSecurityGroupsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListSecurityGroupsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListSecurityGroupsRequestOrderBy) Values() []ListSecurityGroupsRequestOrderBy {
	return []ListSecurityGroupsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"updated_at_desc",
		"updated_at_asc",
	}
}

func (enum ListSecurityGroupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSecurityGroupsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSecurityGroupsRequestOrderBy(ListSecurityGroupsRequestOrderBy(tmp).String())
	return nil
}

type ListServersRequestOrderBy string

const (
	ListServersRequestOrderByCreatedAtDesc = ListServersRequestOrderBy("created_at_desc")
	ListServersRequestOrderByCreatedAtAsc  = ListServersRequestOrderBy("created_at_asc")
	ListServersRequestOrderByUpdatedAtDesc = ListServersRequestOrderBy("updated_at_desc")
	ListServersRequestOrderByUpdatedAtAsc  = ListServersRequestOrderBy("updated_at_asc")
)

func (enum ListServersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListServersRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListServersRequestOrderBy) Values() []ListServersRequestOrderBy {
	return []ListServersRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"updated_at_desc",
		"updated_at_asc",
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

type ListTemplatesRequestOrderBy string

const (
	ListTemplatesRequestOrderByCreatedAtDesc = ListTemplatesRequestOrderBy("created_at_desc")
	ListTemplatesRequestOrderByCreatedAtAsc  = ListTemplatesRequestOrderBy("created_at_asc")
	ListTemplatesRequestOrderByUpdatedAtDesc = ListTemplatesRequestOrderBy("updated_at_desc")
	ListTemplatesRequestOrderByUpdatedAtAsc  = ListTemplatesRequestOrderBy("updated_at_asc")
)

func (enum ListTemplatesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListTemplatesRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListTemplatesRequestOrderBy) Values() []ListTemplatesRequestOrderBy {
	return []ListTemplatesRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
		"updated_at_desc",
		"updated_at_asc",
	}
}

func (enum ListTemplatesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTemplatesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTemplatesRequestOrderBy(ListTemplatesRequestOrderBy(tmp).String())
	return nil
}

type PlacementGroupPolicyType string

const (
	// Unknown policy type.
	PlacementGroupPolicyTypeUnknownPolicyType = PlacementGroupPolicyType("unknown_policy_type")
	// Ensures that all servers in the placement group are placed on the same physical host to minimize latency.
	PlacementGroupPolicyTypeLowLatency = PlacementGroupPolicyType("low_latency")
	// Distributes servers across multiple physical hosts to maximize availability.
	PlacementGroupPolicyTypeMaxAvailability = PlacementGroupPolicyType("max_availability")
)

func (enum PlacementGroupPolicyType) String() string {
	if enum == "" {
		// return default value if empty
		return string(PlacementGroupPolicyTypeUnknownPolicyType)
	}
	return string(enum)
}

func (enum PlacementGroupPolicyType) Values() []PlacementGroupPolicyType {
	return []PlacementGroupPolicyType{
		"unknown_policy_type",
		"low_latency",
		"max_availability",
	}
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

type PrivateNetworkInterfaceStatus string

const (
	// Unknown status.
	PrivateNetworkInterfaceStatusUnknownStatus = PrivateNetworkInterfaceStatus("unknown_status")
	// Interface is available.
	PrivateNetworkInterfaceStatusAvailable = PrivateNetworkInterfaceStatus("available")
	// Interface is being attached.
	PrivateNetworkInterfaceStatusAttaching = PrivateNetworkInterfaceStatus("attaching")
	// Interface is being detached.
	PrivateNetworkInterfaceStatusDetaching = PrivateNetworkInterfaceStatus("detaching")
	// Interface is being synchronized.
	PrivateNetworkInterfaceStatusSyncing = PrivateNetworkInterfaceStatus("syncing")
)

func (enum PrivateNetworkInterfaceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(PrivateNetworkInterfaceStatusUnknownStatus)
	}
	return string(enum)
}

func (enum PrivateNetworkInterfaceStatus) Values() []PrivateNetworkInterfaceStatus {
	return []PrivateNetworkInterfaceStatus{
		"unknown_status",
		"available",
		"attaching",
		"detaching",
		"syncing",
	}
}

func (enum PrivateNetworkInterfaceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PrivateNetworkInterfaceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PrivateNetworkInterfaceStatus(PrivateNetworkInterfaceStatus(tmp).String())
	return nil
}

type SecurityGroupAction string

const (
	// Unknown action.
	SecurityGroupActionUnknownAction = SecurityGroupAction("unknown_action")
	// Accept the traffic.
	SecurityGroupActionAccept = SecurityGroupAction("accept")
	// Drop the traffic.
	SecurityGroupActionDrop = SecurityGroupAction("drop")
)

func (enum SecurityGroupAction) String() string {
	if enum == "" {
		// return default value if empty
		return string(SecurityGroupActionUnknownAction)
	}
	return string(enum)
}

func (enum SecurityGroupAction) Values() []SecurityGroupAction {
	return []SecurityGroupAction{
		"unknown_action",
		"accept",
		"drop",
	}
}

func (enum SecurityGroupAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecurityGroupAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecurityGroupAction(SecurityGroupAction(tmp).String())
	return nil
}

type SecurityGroupRuleAction string

const (
	// Unknown action.
	SecurityGroupRuleActionUnknownAction = SecurityGroupRuleAction("unknown_action")
	// Accept the traffic.
	SecurityGroupRuleActionAccept = SecurityGroupRuleAction("accept")
	// Drop the traffic.
	SecurityGroupRuleActionDrop = SecurityGroupRuleAction("drop")
)

func (enum SecurityGroupRuleAction) String() string {
	if enum == "" {
		// return default value if empty
		return string(SecurityGroupRuleActionUnknownAction)
	}
	return string(enum)
}

func (enum SecurityGroupRuleAction) Values() []SecurityGroupRuleAction {
	return []SecurityGroupRuleAction{
		"unknown_action",
		"accept",
		"drop",
	}
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
	// Unknown direction.
	SecurityGroupRuleDirectionUnknownDirection = SecurityGroupRuleDirection("unknown_direction")
	// Inbound traffic.
	SecurityGroupRuleDirectionInbound = SecurityGroupRuleDirection("inbound")
	// Outbound traffic.
	SecurityGroupRuleDirectionOutbound = SecurityGroupRuleDirection("outbound")
	// Both inbound and outbound traffic.
	SecurityGroupRuleDirectionBoth = SecurityGroupRuleDirection("both")
)

func (enum SecurityGroupRuleDirection) String() string {
	if enum == "" {
		// return default value if empty
		return string(SecurityGroupRuleDirectionUnknownDirection)
	}
	return string(enum)
}

func (enum SecurityGroupRuleDirection) Values() []SecurityGroupRuleDirection {
	return []SecurityGroupRuleDirection{
		"unknown_direction",
		"inbound",
		"outbound",
		"both",
	}
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
	// Unknown protocol.
	SecurityGroupRuleProtocolUnknownProtocol = SecurityGroupRuleProtocol("unknown_protocol")
	// TCP protocol.
	SecurityGroupRuleProtocolTCP = SecurityGroupRuleProtocol("tcp")
	// UDP protocol.
	SecurityGroupRuleProtocolUDP = SecurityGroupRuleProtocol("udp")
	// ICMP protocol.
	SecurityGroupRuleProtocolIcmp = SecurityGroupRuleProtocol("icmp")
	// Any protocol.
	SecurityGroupRuleProtocolAny = SecurityGroupRuleProtocol("any")
)

func (enum SecurityGroupRuleProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return string(SecurityGroupRuleProtocolUnknownProtocol)
	}
	return string(enum)
}

func (enum SecurityGroupRuleProtocol) Values() []SecurityGroupRuleProtocol {
	return []SecurityGroupRuleProtocol{
		"unknown_protocol",
		"tcp",
		"udp",
		"icmp",
		"any",
	}
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

type ServerArchitecture string

const (
	// Architecture is unknown.
	ServerArchitectureUnknownArchitecture = ServerArchitecture("unknown_architecture")
	// X86_64 architecture.
	ServerArchitectureX86_64 = ServerArchitecture("x86_64")
	// AArch64 architecture.
	ServerArchitectureAarch64 = ServerArchitecture("aarch64")
)

func (enum ServerArchitecture) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerArchitectureUnknownArchitecture)
	}
	return string(enum)
}

func (enum ServerArchitecture) Values() []ServerArchitecture {
	return []ServerArchitecture{
		"unknown_architecture",
		"x86_64",
		"aarch64",
	}
}

func (enum ServerArchitecture) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerArchitecture) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerArchitecture(ServerArchitecture(tmp).String())
	return nil
}

type ServerFilesystemStatus string

const (
	ServerFilesystemStatusUnknownStatus = ServerFilesystemStatus("unknown_status")
	ServerFilesystemStatusAttaching     = ServerFilesystemStatus("attaching")
	ServerFilesystemStatusAvailable     = ServerFilesystemStatus("available")
	ServerFilesystemStatusDetaching     = ServerFilesystemStatus("detaching")
)

func (enum ServerFilesystemStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerFilesystemStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ServerFilesystemStatus) Values() []ServerFilesystemStatus {
	return []ServerFilesystemStatus{
		"unknown_status",
		"attaching",
		"available",
		"detaching",
	}
}

func (enum ServerFilesystemStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerFilesystemStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerFilesystemStatus(ServerFilesystemStatus(tmp).String())
	return nil
}

type ServerIPStatus string

const (
	// Status is unknown.
	ServerIPStatusUnknownStatus = ServerIPStatus("unknown_status")
	// IP is detached.
	ServerIPStatusDetached = ServerIPStatus("detached")
	// IP is attached.
	ServerIPStatusAttached = ServerIPStatus("attached")
	// IP is pending.
	ServerIPStatusPending = ServerIPStatus("pending")
	// IP is in error state.
	ServerIPStatusError = ServerIPStatus("error")
)

func (enum ServerIPStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerIPStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ServerIPStatus) Values() []ServerIPStatus {
	return []ServerIPStatus{
		"unknown_status",
		"detached",
		"attached",
		"pending",
		"error",
	}
}

func (enum ServerIPStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerIPStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerIPStatus(ServerIPStatus(tmp).String())
	return nil
}

type ServerPrivateNetworkInterfaceStatus string

const (
	// Status is unknown.
	ServerPrivateNetworkInterfaceStatusUnknownStatus = ServerPrivateNetworkInterfaceStatus("unknown_status")
	// Interface is available.
	ServerPrivateNetworkInterfaceStatusAvailable = ServerPrivateNetworkInterfaceStatus("available")
	// Interface is being attached.
	ServerPrivateNetworkInterfaceStatusAttaching = ServerPrivateNetworkInterfaceStatus("attaching")
	// Interface is being detached.
	ServerPrivateNetworkInterfaceStatusDetaching = ServerPrivateNetworkInterfaceStatus("detaching")
	// Interface is syncing.
	ServerPrivateNetworkInterfaceStatusSyncing = ServerPrivateNetworkInterfaceStatus("syncing")
)

func (enum ServerPrivateNetworkInterfaceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerPrivateNetworkInterfaceStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ServerPrivateNetworkInterfaceStatus) Values() []ServerPrivateNetworkInterfaceStatus {
	return []ServerPrivateNetworkInterfaceStatus{
		"unknown_status",
		"available",
		"attaching",
		"detaching",
		"syncing",
	}
}

func (enum ServerPrivateNetworkInterfaceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerPrivateNetworkInterfaceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerPrivateNetworkInterfaceStatus(ServerPrivateNetworkInterfaceStatus(tmp).String())
	return nil
}

type ServerPublicNetworkInterfaceStatus string

const (
	// Status is unknown.
	ServerPublicNetworkInterfaceStatusUnknownStatus = ServerPublicNetworkInterfaceStatus("unknown_status")
	// Interface is available.
	ServerPublicNetworkInterfaceStatusAvailable = ServerPublicNetworkInterfaceStatus("available")
	// Interface is syncing.
	ServerPublicNetworkInterfaceStatusSyncing = ServerPublicNetworkInterfaceStatus("syncing")
)

func (enum ServerPublicNetworkInterfaceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerPublicNetworkInterfaceStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ServerPublicNetworkInterfaceStatus) Values() []ServerPublicNetworkInterfaceStatus {
	return []ServerPublicNetworkInterfaceStatus{
		"unknown_status",
		"available",
		"syncing",
	}
}

func (enum ServerPublicNetworkInterfaceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerPublicNetworkInterfaceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerPublicNetworkInterfaceStatus(ServerPublicNetworkInterfaceStatus(tmp).String())
	return nil
}

type ServerStatus string

const (
	// Status is unknown.
	ServerStatusUnknownStatus = ServerStatus("unknown_status")
	// Server is running.
	ServerStatusStarted = ServerStatus("started")
	// Server is stopped.
	ServerStatusStopped = ServerStatus("stopped")
	// Server is paused.
	ServerStatusPaused = ServerStatus("paused")
	// Server is starting.
	ServerStatusStarting = ServerStatus("starting")
	// Server is stopping.
	ServerStatusStopping = ServerStatus("stopping")
	// Server is pausing.
	ServerStatusPausing = ServerStatus("pausing")
	// Server is locked.
	ServerStatusLocked = ServerStatus("locked")
	// Server is rebooting.
	ServerStatusRebooting = ServerStatus("rebooting")
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
		"started",
		"stopped",
		"paused",
		"starting",
		"stopping",
		"pausing",
		"locked",
		"rebooting",
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

type ServerTypeArchitecture string

const (
	// Architecture is unknown.
	ServerTypeArchitectureUnknownArchitecture = ServerTypeArchitecture("unknown_architecture")
	// X86_64 architecture.
	ServerTypeArchitectureX86_64 = ServerTypeArchitecture("x86_64")
	// AArch64 architecture.
	ServerTypeArchitectureAarch64 = ServerTypeArchitecture("aarch64")
)

func (enum ServerTypeArchitecture) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerTypeArchitectureUnknownArchitecture)
	}
	return string(enum)
}

func (enum ServerTypeArchitecture) Values() []ServerTypeArchitecture {
	return []ServerTypeArchitecture{
		"unknown_architecture",
		"x86_64",
		"aarch64",
	}
}

func (enum ServerTypeArchitecture) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerTypeArchitecture) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerTypeArchitecture(ServerTypeArchitecture(tmp).String())
	return nil
}

type ServerTypeAvailability string

const (
	// Availability is unknown.
	ServerTypeAvailabilityUnknownAvailability = ServerTypeAvailability("unknown_availability")
	// Server type is available.
	ServerTypeAvailabilityAvailable = ServerTypeAvailability("available")
	// Server type is in low stock.
	ServerTypeAvailabilityLowStock = ServerTypeAvailability("low_stock")
	// Server type is out of stock.
	ServerTypeAvailabilityOutOfStock = ServerTypeAvailability("out_of_stock")
)

func (enum ServerTypeAvailability) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerTypeAvailabilityUnknownAvailability)
	}
	return string(enum)
}

func (enum ServerTypeAvailability) Values() []ServerTypeAvailability {
	return []ServerTypeAvailability{
		"unknown_availability",
		"available",
		"low_stock",
		"out_of_stock",
	}
}

func (enum ServerTypeAvailability) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerTypeAvailability) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerTypeAvailability(ServerTypeAvailability(tmp).String())
	return nil
}

type ServerVolumeVolumeType string

const (
	// Volume type is unknown.
	ServerVolumeVolumeTypeUnknownVolumeType = ServerVolumeVolumeType("unknown_volume_type")
	// Local SSD volume.
	ServerVolumeVolumeTypeLSSD = ServerVolumeVolumeType("l_ssd")
	// Scaleway Block Storage volume.
	ServerVolumeVolumeTypeSbs = ServerVolumeVolumeType("sbs")
	// Scratch volume.
	ServerVolumeVolumeTypeScratch = ServerVolumeVolumeType("scratch")
)

func (enum ServerVolumeVolumeType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerVolumeVolumeTypeUnknownVolumeType)
	}
	return string(enum)
}

func (enum ServerVolumeVolumeType) Values() []ServerVolumeVolumeType {
	return []ServerVolumeVolumeType{
		"unknown_volume_type",
		"l_ssd",
		"sbs",
		"scratch",
	}
}

func (enum ServerVolumeVolumeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerVolumeVolumeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerVolumeVolumeType(ServerVolumeVolumeType(tmp).String())
	return nil
}

// SecurityGroupRulePortRange: security group rule port range.
type SecurityGroupRulePortRange struct {
	// Start: start of the port range.
	Start uint32 `json:"start"`

	// End: end of the port range.
	End uint32 `json:"end"`
}

// CreateServerRequestBookIP: create server request book ip.
type CreateServerRequestBookIP struct {
	// Type: type of IP to book.
	// Default value: unknown_ip_type
	Type CreateServerRequestBookIPIPType `json:"type"`

	// Tags: tags to associate with the IP.
	Tags []string `json:"tags"`
}

// SecurityGroupRule: security group rule.
type SecurityGroupRule struct {
	// ID: unique ID of the rule.
	ID string `json:"id"`

	// Protocol: protocol this rule applies to.
	// Default value: unknown_protocol
	Protocol SecurityGroupRuleProtocol `json:"protocol"`

	// Direction: direction of traffic this rule applies to.
	// Default value: unknown_direction
	Direction SecurityGroupRuleDirection `json:"direction"`

	// Action: action to take when the rule matches.
	// Default value: unknown_action
	Action SecurityGroupRuleAction `json:"action"`

	// SourceIPRange: source IP range for the rule.
	SourceIPRange scw.IPNet `json:"source_ip_range"`

	// DestinationIPRange: destination IP range for the rule.
	DestinationIPRange scw.IPNet `json:"destination_ip_range"`

	// SourcePorts: source port range for the rule.
	SourcePorts *SecurityGroupRulePortRange `json:"source_ports"`

	// DestinationPorts: destination port range for the rule.
	DestinationPorts *SecurityGroupRulePortRange `json:"destination_ports"`

	// Position: position of the rule in the list.
	Position uint32 `json:"position"`
}

// CreateServerRequestServerIP: create server request server ip.
type CreateServerRequestServerIP struct {
	// IpamIPID: ID of the IPAM IP to attach.
	// Precisely one of IpamIPID, NewIP must be set.
	IpamIPID *string `json:"ipam_ip_id,omitempty"`

	// NewIP: configuration for a new IP to book.
	// Precisely one of IpamIPID, NewIP must be set.
	NewIP *CreateServerRequestBookIP `json:"new_ip,omitempty"`
}

// CreateServerRequestCreateVolume: create server request create volume.
type CreateServerRequestCreateVolume struct {
	// Name: name of the volume.
	Name string `json:"name"`

	// Tags: tags to associate with the volume.
	Tags []string `json:"tags"`

	// Size: size of the volume.
	Size *scw.Size `json:"size"`

	// BaseSnapshotID: ID of the base snapshot for the volume.
	// Precisely one of BaseSnapshotID, ImageLabel must be set.
	BaseSnapshotID *string `json:"base_snapshot_id,omitempty"`

	// ImageLabel: label of the image to use for the volume.
	// Precisely one of BaseSnapshotID, ImageLabel must be set.
	ImageLabel *string `json:"image_label,omitempty"`

	// PerfIops: performance IOPS for the volume.
	PerfIops *uint32 `json:"perf_iops"`
}

// ServerTypeGpuInfo: server type gpu info.
type ServerTypeGpuInfo struct {
	// Manufacturer: manufacturer of the GPU.
	Manufacturer string `json:"manufacturer"`

	// Name: name of the GPU.
	Name string `json:"name"`

	// Memory: memory of the GPU.
	Memory scw.Size `json:"memory"`
}

// ServerTypeLimits: server type limits.
type ServerTypeLimits struct {
	// PrivateNetworkCount: maximum number of Private Networks.
	PrivateNetworkCount uint32 `json:"private_network_count"`

	// FileSystemCount: maximum number of filesystems.
	FileSystemCount uint32 `json:"file_system_count"`

	// PrivateNetworkBandwidth: maximum Private Network bandwidth.
	PrivateNetworkBandwidth uint64 `json:"private_network_bandwidth"`

	// BlockBandwidth: maximum block storage bandwidth.
	BlockBandwidth uint64 `json:"block_bandwidth"`

	// InternetBandwidth: maximum internet bandwidth.
	InternetBandwidth uint64 `json:"internet_bandwidth"`

	// LSSDSize: maximum size of local SSD.
	LSSDSize scw.Size `json:"l_ssd_size"`

	// ScratchSize: maximum size of scratch storage.
	ScratchSize scw.Size `json:"scratch_size"`

	// ScratchVolumesCount: maximum number of scratch volumes.
	ScratchVolumesCount uint32 `json:"scratch_volumes_count"`

	// IPCount: maximum number of IPs.
	IPCount uint32 `json:"ip_count"`

	// VolumeCount: maximum number of volumes.
	VolumeCount uint32 `json:"volume_count"`
}

// ServerIP: server ip.
type ServerIP struct {
	ID string `json:"id"`

	Dynamic bool `json:"dynamic"`

	// Status: default value: unknown_status
	Status ServerIPStatus `json:"status"`

	Default bool `json:"default"`
}

// CreateTemplateRequestPrivateNetworkTemplate: create template request private network template.
type CreateTemplateRequestPrivateNetworkTemplate struct {
	// PrivateNetworkID: ID of the private network.
	PrivateNetworkID string `json:"private_network_id"`
}

// CreateTemplateRequestVolumeTemplate: create template request volume template.
type CreateTemplateRequestVolumeTemplate struct {
	// VolumeType: type of the volume.
	// Default value: unknown_volume_type
	VolumeType CreateServerRequestServerVolumeVolumeType `json:"volume_type"`

	// Name: name of the volume.
	Name string `json:"name"`

	// Tags: tags associated with the volume.
	Tags []string `json:"tags"`

	// Size: size of the volume in bytes.
	Size *scw.Size `json:"size"`

	// BaseSnapshotID: ID of the base snapshot for the volume.
	// Precisely one of BaseSnapshotID, ImageLabel must be set.
	BaseSnapshotID *string `json:"base_snapshot_id,omitempty"`

	// ImageLabel: label of the image used as base for the volume.
	// Precisely one of BaseSnapshotID, ImageLabel must be set.
	ImageLabel *string `json:"image_label,omitempty"`

	// PerfIops: performance IOPS for the volume.
	PerfIops *uint32 `json:"perf_iops"`
}

// SecurityGroupRuleConfig: security group rule config.
type SecurityGroupRuleConfig struct {
	// Protocol: protocol for the rule.
	// Default value: unknown_protocol
	Protocol SecurityGroupRuleProtocol `json:"protocol"`

	// Direction: direction of traffic for the rule.
	// Default value: unknown_direction
	Direction SecurityGroupRuleDirection `json:"direction"`

	// Action: action to take when the rule matches.
	// Default value: unknown_action
	Action SecurityGroupRuleAction `json:"action"`

	// SourceIPRange: source IP range for the rule.
	SourceIPRange scw.IPNet `json:"source_ip_range"`

	// DestinationIPRange: destination IP range for the rule.
	DestinationIPRange scw.IPNet `json:"destination_ip_range"`

	// SourcePorts: source port range for the rule.
	SourcePorts *SecurityGroupRulePortRange `json:"source_ports"`

	// DestinationPorts: destination port range for the rule.
	DestinationPorts *SecurityGroupRulePortRange `json:"destination_ports"`

	// Position: position of the rule in the list.
	Position int32 `json:"position"`
}

// SecurityGroup: security group.
type SecurityGroup struct {
	// ID: unique ID of the security group.
	ID string `json:"id"`

	// Name: name of the security group.
	Name string `json:"name"`

	// Description: description of the security group.
	Description string `json:"description"`

	// ProjectID: project ID the security group belongs to.
	ProjectID string `json:"project_id"`

	// Tags: tags associated with the security group.
	Tags []string `json:"tags"`

	// DisableDefaultRules: true if default rules are disabled.
	DisableDefaultRules bool `json:"disable_default_rules"`

	// ProjectDefault: true if this is the default security group for the project.
	ProjectDefault bool `json:"project_default"`

	// InboundDefaultAction: default action for inbound rules.
	// Default value: unknown_action
	InboundDefaultAction SecurityGroupAction `json:"inbound_default_action"`

	// OutboundDefaultAction: default action for outbound rules.
	// Default value: unknown_action
	OutboundDefaultAction SecurityGroupAction `json:"outbound_default_action"`

	// Stateless: true if the security group is stateless.
	Stateless bool `json:"stateless"`

	// CreatedAt: creation timestamp of the security group.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update timestamp of the security group.
	UpdatedAt *time.Time `json:"updated_at"`

	// DefaultRules: list of default rules applied to the security group.
	DefaultRules []*SecurityGroupRule `json:"default_rules"`

	// Rules: list of custom rules applied to the security group.
	Rules []*SecurityGroupRule `json:"rules"`

	// Zone: zone in which the security group is located.
	Zone scw.Zone `json:"zone"`
}

// CreateServerRequestPublicNetworkInterface: create server request public network interface.
type CreateServerRequestPublicNetworkInterface struct {
	// SecurityGroupID: ID of the security group for the interface.
	SecurityGroupID *string `json:"security_group_id"`

	// IPs: list of IPs to attach to the interface.
	IPs []*CreateServerRequestServerIP `json:"ips"`
}

// CreateServerRequestServerVolume: create server request server volume.
type CreateServerRequestServerVolume struct {
	// VolumeType: type of the volume.
	// Default value: unknown_volume_type
	VolumeType CreateServerRequestServerVolumeVolumeType `json:"volume_type"`

	// VolumeID: ID of the volume to attach.
	// Precisely one of VolumeID, NewVolume must be set.
	VolumeID *string `json:"volume_id,omitempty"`

	// NewVolume: configuration for a new volume to create.
	// Precisely one of VolumeID, NewVolume must be set.
	NewVolume *CreateServerRequestCreateVolume `json:"new_volume,omitempty"`
}

// PlacementGroup: placement group.
type PlacementGroup struct {
	// ID: placement group unique ID.
	ID string `json:"id"`

	// ProjectID: placement group Project ID.
	ProjectID string `json:"project_id"`

	// Name: placement group name.
	Name string `json:"name"`

	// PolicyType: select the behavior of the placement group, either low_latency (group) or max_availability (spread).
	// Default value: unknown_policy_type
	PolicyType PlacementGroupPolicyType `json:"policy_type"`

	// Tags: placement group tags.
	Tags []string `json:"tags"`

	// CreatedAt: placement group creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: placement group modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Zone: zone in which the placement group is located.
	Zone scw.Zone `json:"zone"`
}

// PrivateNetworkInterfaceSummary: private network interface summary.
type PrivateNetworkInterfaceSummary struct {
	// ID: unique ID of the private network interface.
	ID string `json:"id"`

	// PrivateNetworkID: ID of the Private Network this interface is attached to.
	PrivateNetworkID string `json:"private_network_id"`

	// ProjectID: project ID the private network interface belongs to.
	ProjectID string `json:"project_id"`

	// ServerID: ID of the Instance this interface is attached to.
	ServerID string `json:"server_id"`

	// MacAddress: mAC address of the private network interface.
	MacAddress string `json:"mac_address"`

	// Status: current status of the private network interface.
	// Default value: unknown_status
	Status PrivateNetworkInterfaceStatus `json:"status"`

	// IPIDs: list of IP IDs attached to this interface.
	IPIDs []string `json:"ip_ids"`

	// Tags: tags associated with the private network interface.
	Tags []string `json:"tags"`

	// CreatedAt: creation timestamp of the private network interface.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update timestamp of the private network interface.
	UpdatedAt *time.Time `json:"updated_at"`
}

// SecurityGroupSummary: security group summary.
type SecurityGroupSummary struct {
	// ID: unique ID of the security group.
	ID string `json:"id"`

	// Name: name of the security group.
	Name string `json:"name"`

	// Description: description of the security group.
	Description string `json:"description"`

	// ProjectID: project ID the security group belongs to.
	ProjectID string `json:"project_id"`

	// Tags: tags associated with the security group.
	Tags []string `json:"tags"`

	// DisableDefaultRules: true if default rules are disabled.
	DisableDefaultRules bool `json:"disable_default_rules"`

	// ProjectDefault: true if this is the default security group for the project.
	ProjectDefault bool `json:"project_default"`

	// InboundDefaultAction: default action for inbound rules.
	// Default value: unknown_action
	InboundDefaultAction SecurityGroupAction `json:"inbound_default_action"`

	// OutboundDefaultAction: default action for outbound rules.
	// Default value: unknown_action
	OutboundDefaultAction SecurityGroupAction `json:"outbound_default_action"`

	// Stateless: true if the security group is stateless.
	Stateless bool `json:"stateless"`

	// CreatedAt: creation timestamp of the security group.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update timestamp of the security group.
	UpdatedAt *time.Time `json:"updated_at"`
}

// ServerType: server type.
type ServerType struct {
	// Name: name of the server type.
	Name string `json:"name"`

	// VcpuCount: number of vCPUs.
	VcpuCount uint32 `json:"vcpu_count"`

	// GpuCount: number of GPUs.
	GpuCount uint32 `json:"gpu_count"`

	// Memory: amount of memory.
	Memory scw.Size `json:"memory"`

	// Architecture: architecture of the server type.
	// Default value: unknown_architecture
	Architecture ServerTypeArchitecture `json:"architecture"`

	// Availability: availability status of the server type.
	// Default value: unknown_availability
	Availability ServerTypeAvailability `json:"availability"`

	// Limits: limits for the server type.
	Limits *ServerTypeLimits `json:"limits"`

	// GpuInfo: gPU information for the server type.
	GpuInfo *ServerTypeGpuInfo `json:"gpu_info"`

	// EndOfService: whether the server type has reached end of service.
	EndOfService bool `json:"end_of_service"`
}

// ServerSummary: server summary.
type ServerSummary struct {
	// ID: unique ID of the server.
	ID string `json:"id"`

	// Name: name of the server.
	Name string `json:"name"`

	// ProjectID: project ID to which the server belongs.
	ProjectID string `json:"project_id"`

	// Tags: tags associated with the server.
	Tags []string `json:"tags"`

	// ServerType: type of the server.
	ServerType string `json:"server_type"`

	// PlacementGroupID: ID of the placement group the server belongs to.
	PlacementGroupID *string `json:"placement_group_id"`

	// Status: current status of the server.
	// Default value: unknown_status
	Status ServerStatus `json:"status"`

	// Architecture: architecture of the server.
	// Default value: unknown_architecture
	Architecture ServerArchitecture `json:"architecture"`

	// CreatedAt: creation timestamp of the server.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update timestamp of the server.
	UpdatedAt *time.Time `json:"updated_at"`

	// RescueMode: whether the server is in rescue mode.
	RescueMode bool `json:"rescue_mode"`
}

// TemplateSummary: template summary.
type TemplateSummary struct {
	// ProjectID: project ID associated with the template.
	ProjectID string `json:"project_id"`

	// ID: unique ID of the template.
	ID string `json:"id"`

	// Name: name of the template.
	Name string `json:"name"`

	// Tags: tags associated with the template.
	Tags []string `json:"tags"`

	// ServerTags: tags associated with servers created from this template.
	ServerTags []string `json:"server_tags"`

	// ServerType: commercial type of the server defined by the template.
	ServerType string `json:"server_type"`

	// SecurityGroupID: security group ID associated with the template.
	SecurityGroupID *string `json:"security_group_id"`

	// PlacementGroupID: placement group ID associated with the template.
	PlacementGroupID *string `json:"placement_group_id"`

	// PublicIPV4Count: number of IPv4 public IPs to attach to servers created from this template.
	PublicIPV4Count uint32 `json:"public_ip_v4_count"`

	// PublicIPV6Count: number of IPv6 public IPs to attach to servers created from this template.
	PublicIPV6Count uint32 `json:"public_ip_v6_count"`

	// FilesystemIDs: list of Filesystem IDs associated with the template.
	FilesystemIDs []string `json:"filesystem_ids"`

	// CreatedAt: creation timestamp of the template.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update timestamp of the template.
	UpdatedAt *time.Time `json:"updated_at"`

	// Zone: zone in which the template is located.
	Zone scw.Zone `json:"zone"`
}

// ServerFilesystem: server filesystem.
type ServerFilesystem struct {
	ID string `json:"id"`

	// Status: default value: unknown_status
	Status ServerFilesystemStatus `json:"status"`
}

// ServerPrivateNetworkInterface: server private network interface.
type ServerPrivateNetworkInterface struct {
	ID string `json:"id"`

	PrivateNetworkID string `json:"private_network_id"`

	MacAddress string `json:"mac_address"`

	// Status: default value: unknown_status
	Status ServerPrivateNetworkInterfaceStatus `json:"status"`

	IPIDs []string `json:"ip_ids"`

	SecurityGroupID string `json:"security_group_id"`
}

// ServerPublicNetworkInterface: server public network interface.
type ServerPublicNetworkInterface struct {
	// Status: default value: unknown_status
	Status ServerPublicNetworkInterfaceStatus `json:"status"`

	MacAddress string `json:"mac_address"`

	SecurityGroupID string `json:"security_group_id"`

	IPs []*ServerIP `json:"ips"`

	DNS string `json:"dns"`
}

// ServerRDPPassword: server rdp password.
type ServerRDPPassword struct {
	EncryptedPassword string `json:"encrypted_password"`

	RdpSSHKeyID string `json:"rdp_ssh_key_id"`
}

// ServerVolume: server volume.
type ServerVolume struct {
	ID string `json:"id"`

	// VolumeType: default value: unknown_volume_type
	VolumeType ServerVolumeVolumeType `json:"volume_type"`
}

// UpdateServerRequestPublicNetworkInterface: update server request public network interface.
type UpdateServerRequestPublicNetworkInterface struct {
	// SecurityGroupID: ID of the security group for the interface.
	SecurityGroupID *string `json:"security_group_id"`
}

// UpdateTemplateRequestUpdatePrivateNetworks: update template request update private networks.
type UpdateTemplateRequestUpdatePrivateNetworks struct {
	// PrivateNetworks: list of updated private networks.
	PrivateNetworks []*CreateTemplateRequestPrivateNetworkTemplate `json:"private_networks"`
}

// UpdateTemplateRequestUpdateVolumes: update template request update volumes.
type UpdateTemplateRequestUpdateVolumes struct {
	// Volumes: list of updated volume templates.
	Volumes []*CreateTemplateRequestVolumeTemplate `json:"volumes"`
}

// AddSecurityGroupRulesRequest: add security group rules request.
type AddSecurityGroupRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// SecurityGroupID: ID of the security group to add rules to.
	SecurityGroupID string `json:"security_group_id"`

	// SecurityGroupRules: list of rules to add.
	SecurityGroupRules []*SecurityGroupRuleConfig `json:"security_group_rules"`
}

// AddSecurityGroupRulesResponse: add security group rules response.
type AddSecurityGroupRulesResponse struct {
	// SecurityGroup: updated security group.
	SecurityGroup *SecurityGroup `json:"security_group"`

	// AddedRules: list of rules that were added.
	AddedRules []*SecurityGroupRule `json:"added_rules"`
}

// AttachServerFileSystemRequest: attach server file system request.
type AttachServerFileSystemRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to attach the filesystem to.
	ServerID string `json:"-"`

	// FilesystemID: ID of the filesystem to attach.
	FilesystemID string `json:"filesystem_id"`
}

// AttachServerIPRequest: attach server ip request.
type AttachServerIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to attach the IP to.
	ServerID string `json:"-"`

	// IPID: ID of the IP to attach.
	IPID string `json:"ip_id"`

	// Default: whether the IP should be the default IP.
	Default bool `json:"default"`

	// MoveAllowed: whether moving the IP is allowed.
	MoveAllowed bool `json:"move_allowed"`
}

// AttachServerPrivateNetworkInterfaceRequest: attach server private network interface request.
type AttachServerPrivateNetworkInterfaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to attach the private network interface to.
	ServerID string `json:"-"`

	// PrivateNetworkInterfaceID: ID of the private network interface to attach.
	PrivateNetworkInterfaceID string `json:"private_network_interface_id"`
}

// AttachServerVolumeRequest: attach server volume request.
type AttachServerVolumeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to attach the volume to.
	ServerID string `json:"-"`

	// VolumeID: ID of the volume to attach.
	VolumeID string `json:"volume_id"`

	// VolumeType: type of the volume.
	// Default value: unknown_volume_type
	VolumeType ServerVolumeVolumeType `json:"volume_type"`

	// BootVolume: whether the volume should be used as the boot volume.
	BootVolume bool `json:"boot_volume"`
}

// CheckTemplateRequest: check template request.
type CheckTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: unique ID of the template to check.
	TemplateID string `json:"-"`
}

// CreatePlacementGroupRequest: create placement group request.
type CreatePlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ProjectID: project ID of the placement group.
	ProjectID string `json:"project_id"`

	// Name: name of the placement group.
	Name string `json:"name"`

	// PolicyType: policy type of the placement group.
	// Default value: unknown_policy_type
	PolicyType PlacementGroupPolicyType `json:"policy_type"`

	// Tags: tags of the placement group.
	Tags []string `json:"tags"`
}

// CreatePrivateNetworkInterfaceRequest: create private network interface request.
type CreatePrivateNetworkInterfaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PrivateNetworkID: ID of the Private Network to attach to.
	PrivateNetworkID string `json:"private_network_id"`

	// ProjectID: project ID for the private network interface.
	ProjectID string `json:"project_id"`

	// ServerID: ID of the Instance to attach the interface to.
	ServerID *string `json:"server_id,omitempty"`

	// IPIDs: list of IP IDs to attach to the interface.
	IPIDs []string `json:"ip_ids"`

	// Tags: tags to assign to the private network interface.
	Tags []string `json:"tags"`
}

// CreateSecurityGroupRequest: create security group request.
type CreateSecurityGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Name: name of the security group.
	Name string `json:"name"`

	// Description: description of the security group.
	Description string `json:"description"`

	// DisableDefaultRules: whether to disable default rules.
	DisableDefaultRules bool `json:"disable_default_rules"`

	// ProjectID: project ID the security group belongs to.
	ProjectID string `json:"project_id"`

	// Tags: tags for the security group.
	Tags []string `json:"tags"`

	// ProjectDefault: whether this should be the default security group for the project.
	ProjectDefault bool `json:"project_default"`

	// InboundDefaultAction: default action for inbound rules.
	// Default value: unknown_action
	InboundDefaultAction SecurityGroupAction `json:"inbound_default_action"`

	// OutboundDefaultAction: default action for outbound rules.
	// Default value: unknown_action
	OutboundDefaultAction SecurityGroupAction `json:"outbound_default_action"`

	// Stateless: whether the security group should be stateless.
	Stateless bool `json:"stateless"`
}

// CreateServerFromTemplateRequest: create server from template request.
type CreateServerFromTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: unique ID of the template to use.
	TemplateID string `json:"-"`

	// Name: name of the new server.
	Name string `json:"name"`
}

// CreateServerRequest: create server request.
type CreateServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ProjectID: project ID for the server.
	ProjectID string `json:"project_id"`

	// Name: name of the server.
	Name string `json:"name"`

	// Tags: tags to associate with the server.
	Tags []string `json:"tags"`

	// ServerType: type of the server.
	ServerType string `json:"server_type"`

	// PlacementGroupID: ID of the placement group the server belongs to.
	PlacementGroupID *string `json:"placement_group_id,omitempty"`

	// Volumes: volumes to attach to the server.
	Volumes []*CreateServerRequestServerVolume `json:"volumes"`

	// WindowsRdpSSHKeyID: iAM ID of the SSH key used to encrypt the Windows `Administrator` password for RDP use.
	WindowsRdpSSHKeyID *string `json:"windows_rdp_ssh_key_id,omitempty"`

	// PublicNetworkInterface: public network interface configuration.
	PublicNetworkInterface *CreateServerRequestPublicNetworkInterface `json:"public_network_interface,omitempty"`
}

// CreateTemplateRequest: create template request.
type CreateTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ProjectID: project ID for the template.
	ProjectID string `json:"project_id"`

	// Name: name of the template.
	Name string `json:"name"`

	// Tags: tags to associate with the template.
	Tags []string `json:"tags"`

	// ServerTags: tags to associate with servers created from the template.
	ServerTags []string `json:"server_tags"`

	// ServerType: commercial type of the server defined by the template.
	ServerType string `json:"server_type"`

	// SecurityGroupID: security group ID for the template.
	SecurityGroupID *string `json:"security_group_id,omitempty"`

	// PlacementGroupID: placement group ID for the template.
	PlacementGroupID *string `json:"placement_group_id,omitempty"`

	// Volumes: list of volume templates to define volumes for servers.
	Volumes []*CreateTemplateRequestVolumeTemplate `json:"volumes"`

	// PrivateNetworks: list of private networks to associate with the template.
	PrivateNetworks []*CreateTemplateRequestPrivateNetworkTemplate `json:"private_networks"`

	// FilesystemIDs: list of filesystem IDs to associate with the template.
	FilesystemIDs []string `json:"filesystem_ids"`

	// PublicIPV4Count: number of IPv4 public IPs to attach to servers created from this template.
	PublicIPV4Count uint32 `json:"public_ip_v4_count"`

	// PublicIPV6Count: number of IPv6 public IPs to attach to servers created from this template.
	PublicIPV6Count uint32 `json:"public_ip_v6_count"`

	// WindowsRdpSSHKeyID: iAM ID of the SSH key used to encrypt the Windows `Administrator` password for RDP use.
	WindowsRdpSSHKeyID *string `json:"windows_rdp_ssh_key_id,omitempty"`
}

// DeletePlacementGroupRequest: delete placement group request.
type DeletePlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PlacementGroupID: UUID of the placement group you want to delete.
	PlacementGroupID string `json:"-"`
}

// DeletePrivateNetworkInterfaceRequest: delete private network interface request.
type DeletePrivateNetworkInterfaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PrivateNetworkInterfaceID: ID of the private network interface to delete.
	PrivateNetworkInterfaceID string `json:"-"`
}

// DeleteSecurityGroupRequest: delete security group request.
type DeleteSecurityGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// SecurityGroupID: ID of the security group to delete.
	SecurityGroupID string `json:"-"`
}

// DeleteSecurityGroupRulesRequest: delete security group rules request.
type DeleteSecurityGroupRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// SecurityGroupRuleIDs: list of rule IDs to delete.
	SecurityGroupRuleIDs []string `json:"security_group_rule_ids"`
}

// DeleteServerRequest: delete server request.
type DeleteServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to delete.
	ServerID string `json:"-"`

	// DeleteAllIPs: whether to delete all IPs attached to the server.
	// Precisely one of DeleteAllIPs, DeleteIPIDs must be set.
	DeleteAllIPs *bool `json:"delete_all_ips,omitempty"`

	// DeleteIPIDs: list of IP IDs to delete.
	// Precisely one of DeleteAllIPs, DeleteIPIDs must be set.
	DeleteIPIDs *[]string `json:"delete_ip_ids,omitempty"`

	// DeleteAllVolumes: whether to delete all volumes attached to the server.
	// Precisely one of DeleteAllVolumes, DeleteVolumeIDs must be set.
	DeleteAllVolumes *bool `json:"delete_all_volumes,omitempty"`

	// DeleteVolumeIDs: list of volume IDs to delete.
	// Precisely one of DeleteAllVolumes, DeleteVolumeIDs must be set.
	DeleteVolumeIDs *[]string `json:"delete_volume_ids,omitempty"`

	// KeepAllPrivateNics: whether to keep all private network interfaces.
	// Precisely one of KeepAllPrivateNics, DeletePrivateNicIDs must be set.
	KeepAllPrivateNics *bool `json:"keep_all_private_nics,omitempty"`

	// DeletePrivateNicIDs: list of private network interface IDs to delete.
	// Precisely one of KeepAllPrivateNics, DeletePrivateNicIDs must be set.
	DeletePrivateNicIDs *[]string `json:"delete_private_nic_ids,omitempty"`
}

// DeleteTemplateRequest: delete template request.
type DeleteTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: unique ID of the template to delete.
	TemplateID string `json:"-"`
}

// DeleteTemplateUserDataRequest: delete template user data request.
type DeleteTemplateUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: unique ID of the template.
	TemplateID string `json:"-"`

	// Key: key of the user data to delete.
	Key string `json:"-"`
}

// DeleteUserDataRequest: delete user data request.
type DeleteUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: the ID of the server.
	ServerID string `json:"-"`

	// Key: the key of the user data to delete.
	Key string `json:"-"`
}

// DetachServerFileSystemRequest: detach server file system request.
type DetachServerFileSystemRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to detach the filesystem from.
	ServerID string `json:"-"`

	// FilesystemID: ID of the filesystem to detach.
	FilesystemID string `json:"filesystem_id"`
}

// DetachServerIPRequest: detach server ip request.
type DetachServerIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to detach the IP from.
	ServerID string `json:"-"`

	// IPID: ID of the IP to detach.
	IPID string `json:"ip_id"`
}

// DetachServerPrivateNetworkInterfaceRequest: detach server private network interface request.
type DetachServerPrivateNetworkInterfaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to detach the private network interface from.
	ServerID string `json:"-"`

	// PrivateNetworkInterfaceID: ID of the private network interface to detach.
	PrivateNetworkInterfaceID string `json:"private_network_interface_id"`
}

// DetachServerVolumeRequest: detach server volume request.
type DetachServerVolumeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to detach the volume from.
	ServerID string `json:"-"`

	// VolumeID: ID of the volume to detach.
	VolumeID string `json:"volume_id"`
}

// GetPlacementGroupRequest: get placement group request.
type GetPlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PlacementGroupID: UUID of the placement group you want to get.
	PlacementGroupID string `json:"-"`
}

// GetPrivateNetworkInterfaceRequest: get private network interface request.
type GetPrivateNetworkInterfaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PrivateNetworkInterfaceID: ID of the private network interface to retrieve.
	PrivateNetworkInterfaceID string `json:"-"`
}

// GetResourceCountsRequest: get resource counts request.
type GetResourceCountsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrganizationID: organization ID to filter resource counts.
	// Precisely one of OrganizationID, ProjectID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// ProjectID: project ID to filter resource counts.
	// Precisely one of OrganizationID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`
}

// GetSecurityGroupRequest: get security group request.
type GetSecurityGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// SecurityGroupID: ID of the security group to retrieve.
	SecurityGroupID string `json:"-"`
}

// GetServerCloudInitRequest: get server cloud init request.
type GetServerCloudInitRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: the ID of the server.
	ServerID string `json:"-"`
}

// GetServerRequest: get server request.
type GetServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to retrieve.
	ServerID string `json:"-"`
}

// GetServiceInfoRequest: get service info request.
type GetServiceInfoRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// GetTemplateCloudInitRequest: get template cloud init request.
type GetTemplateCloudInitRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: unique ID of the template.
	TemplateID string `json:"-"`
}

// GetTemplateRequest: get template request.
type GetTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: unique ID of the template to retrieve.
	TemplateID string `json:"-"`
}

// GetTemplateUserDataRequest: get template user data request.
type GetTemplateUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: unique ID of the template.
	TemplateID string `json:"-"`

	// Key: key of the user data to retrieve.
	Key string `json:"-"`
}

// GetUserDataRequest: get user data request.
type GetUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: the ID of the server.
	ServerID string `json:"-"`

	// Key: the key of the user data to retrieve.
	Key string `json:"-"`
}

// ListPlacementGroupsRequest: list placement groups request.
type ListPlacementGroupsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PageToken: the initial pagination token to start from.
	PageToken *string `json:"-"`

	// PageSize: the maximum number of placement groups to return.
	PageSize *uint32 `json:"-"`

	// OrderBy: the field by which to order the result list.
	// Default value: created_at_desc
	OrderBy ListPlacementGroupsRequestOrderBy `json:"-"`

	// ProjectID: list only placement groups of this Project ID.
	ProjectID string `json:"-"`

	// PlacementGroupIDs: list only placement groups with these IDs.
	PlacementGroupIDs []string `json:"-"`

	// Name: filter placement groups by name.
	Name *string `json:"-"`

	// Tags: list placement groups with these exact tags.
	Tags []string `json:"-"`
}

// ListPlacementGroupsResponse: list placement groups response.
type ListPlacementGroupsResponse struct {
	// PlacementGroups: list of placement groups.
	PlacementGroups []*PlacementGroup `json:"placement_groups"`

	// NextPageToken: the pagination token, use it to get the next page of results.
	NextPageToken *string `json:"next_page_token"`

	// TotalCount: total number of placement groups.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPlacementGroupsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPlacementGroupsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListPlacementGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PlacementGroups = append(r.PlacementGroups, results.PlacementGroups...)
	r.TotalCount += uint64(len(results.PlacementGroups))
	return uint64(len(results.PlacementGroups)), nil
}

// ListPrivateNetworkInterfacesRequest: list private network interfaces request.
type ListPrivateNetworkInterfacesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PageToken: token for pagination.
	PageToken *string `json:"-"`

	// PageSize: number of items to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: field to order results by.
	// Default value: created_at_desc
	OrderBy ListPrivateNetworkInterfacesRequestOrderBy `json:"-"`

	// ProjectID: filter by Project ID.
	ProjectID string `json:"-"`

	// ServerIDs: filter by server IDs.
	ServerIDs []string `json:"-"`

	// PrivateNetworkIDs: filter by Private Network IDs.
	PrivateNetworkIDs []string `json:"-"`

	// Tags: filter by tags.
	Tags []string `json:"-"`
}

// ListPrivateNetworkInterfacesResponse: list private network interfaces response.
type ListPrivateNetworkInterfacesResponse struct {
	// PrivateNetworkInterfaces: list of private network interfaces.
	PrivateNetworkInterfaces []*PrivateNetworkInterfaceSummary `json:"private_network_interfaces"`

	// NextPageToken: token for the next page.
	NextPageToken *string `json:"next_page_token"`

	// TotalCount: total number of items.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPrivateNetworkInterfacesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPrivateNetworkInterfacesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListPrivateNetworkInterfacesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PrivateNetworkInterfaces = append(r.PrivateNetworkInterfaces, results.PrivateNetworkInterfaces...)
	r.TotalCount += uint64(len(results.PrivateNetworkInterfaces))
	return uint64(len(results.PrivateNetworkInterfaces)), nil
}

// ListSecurityGroupsRequest: list security groups request.
type ListSecurityGroupsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PageToken: token for pagination.
	PageToken *string `json:"-"`

	// PageSize: number of items to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: field and direction to sort by.
	// Default value: created_at_desc
	OrderBy ListSecurityGroupsRequestOrderBy `json:"-"`

	// ProjectID: filter by Project ID.
	ProjectID string `json:"-"`

	// Name: filter by name.
	Name *string `json:"-"`

	// Tags: filter by tags.
	Tags []string `json:"-"`

	// SecurityGroupIDs: filter by specific security group IDs.
	SecurityGroupIDs []string `json:"-"`
}

// ListSecurityGroupsResponse: list security groups response.
type ListSecurityGroupsResponse struct {
	// SecurityGroups: list of security groups.
	SecurityGroups []*SecurityGroupSummary `json:"security_groups"`

	// NextPageToken: token for the next page.
	NextPageToken *string `json:"next_page_token"`

	// TotalCount: total number of items.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecurityGroupsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecurityGroupsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListSecurityGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SecurityGroups = append(r.SecurityGroups, results.SecurityGroups...)
	r.TotalCount += uint64(len(results.SecurityGroups))
	return uint64(len(results.SecurityGroups)), nil
}

// ListServerTypesRequest: list server types request.
type ListServerTypesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PageToken: token for pagination.
	PageToken *string `json:"-"`

	// PageSize: number of server types to return per page.
	PageSize *uint32 `json:"-"`
}

// ListServerTypesResponse: list server types response.
type ListServerTypesResponse struct {
	// ServerTypes: list of server types.
	ServerTypes []*ServerType `json:"server_types"`

	// NextPageToken: token for the next page.
	NextPageToken *string `json:"next_page_token"`

	// TotalCount: total number of server types.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServerTypesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServerTypesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListServerTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ServerTypes = append(r.ServerTypes, results.ServerTypes...)
	r.TotalCount += uint64(len(results.ServerTypes))
	return uint64(len(results.ServerTypes)), nil
}

// ListServersRequest: list servers request.
type ListServersRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PageToken: token for pagination.
	PageToken *string `json:"-"`

	// PageSize: number of servers to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the returned servers.
	// Default value: created_at_desc
	OrderBy ListServersRequestOrderBy `json:"-"`

	// ProjectID: project ID to filter servers.
	ProjectID string `json:"-"`

	// ServerIDs: list of server IDs to filter.
	ServerIDs []string `json:"-"`

	// Name: name to filter servers.
	Name *string `json:"-"`

	// ServerType: server type to filter.
	ServerType *string `json:"-"`

	// Tags: tags to filter servers.
	Tags []string `json:"-"`

	// SecurityGroupIDs: security group IDs to filter servers.
	SecurityGroupIDs []string `json:"-"`

	// PlacementGroupIDs: placement group IDs to filter servers.
	PlacementGroupIDs []string `json:"-"`

	// PrivateNetworkIDs: private Network IDs to filter servers.
	PrivateNetworkIDs []string `json:"-"`

	// MacAddresses: mAC addresses to filter servers.
	MacAddresses []string `json:"-"`
}

// ListServersResponse: list servers response.
type ListServersResponse struct {
	// Servers: list of servers.
	Servers []*ServerSummary `json:"servers"`

	// NextPageToken: token for the next page.
	NextPageToken *string `json:"next_page_token"`

	// TotalCount: total number of servers.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint64(len(results.Servers))
	return uint64(len(results.Servers)), nil
}

// ListTemplateUserDataKeysRequest: list template user data keys request.
type ListTemplateUserDataKeysRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: unique ID of the template.
	TemplateID string `json:"-"`

	// PageToken: token for pagination.
	PageToken *string `json:"-"`

	// PageSize: number of items to return per page.
	PageSize *uint32 `json:"-"`
}

// ListTemplateUserDataKeysResponse: list template user data keys response.
type ListTemplateUserDataKeysResponse struct {
	// Keys: list of user data keys associated with the template.
	Keys []string `json:"keys"`

	// NextPageToken: token for the next page.
	NextPageToken *string `json:"next_page_token"`

	// TotalCount: total number of items.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTemplateUserDataKeysResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTemplateUserDataKeysResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListTemplateUserDataKeysResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Keys = append(r.Keys, results.Keys...)
	r.TotalCount += uint64(len(results.Keys))
	return uint64(len(results.Keys)), nil
}

// ListTemplatesRequest: list templates request.
type ListTemplatesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PageToken: token for pagination.
	PageToken *string `json:"-"`

	// PageSize: number of items to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: field to sort results by.
	// Default value: created_at_desc
	OrderBy ListTemplatesRequestOrderBy `json:"-"`

	// ProjectID: filter by Project ID.
	ProjectID string `json:"-"`

	// TemplateIDs: filter by specific template IDs.
	TemplateIDs []string `json:"-"`

	// Name: filter by template name.
	Name *string `json:"-"`

	// Tags: filter by tags.
	Tags []string `json:"-"`

	// ServerTags: filter by server tags.
	ServerTags []string `json:"-"`

	// SecurityGroupIDs: filter by security group IDs.
	SecurityGroupIDs []string `json:"-"`

	// PlacementGroupIDs: filter by placement group IDs.
	PlacementGroupIDs []string `json:"-"`
}

// ListTemplatesResponse: list templates response.
type ListTemplatesResponse struct {
	// Templates: list of template summaries.
	Templates []*TemplateSummary `json:"templates"`

	// NextPageToken: token for the next page.
	NextPageToken *string `json:"next_page_token"`

	// TotalCount: total number of items.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTemplatesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTemplatesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListTemplatesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Templates = append(r.Templates, results.Templates...)
	r.TotalCount += uint64(len(results.Templates))
	return uint64(len(results.Templates)), nil
}

// ListUserDataKeysRequest: list user data keys request.
type ListUserDataKeysRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: the ID of the server.
	ServerID string `json:"-"`

	// PageToken: page token for pagination.
	PageToken *string `json:"-"`

	// PageSize: number of items to return per page.
	PageSize *uint32 `json:"-"`
}

// ListUserDataKeysResponse: list user data keys response.
type ListUserDataKeysResponse struct {
	// Keys: list of user data keys.
	Keys []string `json:"keys"`

	// NextPageToken: token for the next page.
	NextPageToken *string `json:"next_page_token"`

	// TotalCount: total number of items.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListUserDataKeysResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListUserDataKeysResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListUserDataKeysResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Keys = append(r.Keys, results.Keys...)
	r.TotalCount += uint64(len(results.Keys))
	return uint64(len(results.Keys)), nil
}

// PauseServerRequest: pause server request.
type PauseServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to pause.
	ServerID string `json:"-"`
}

// PrivateNetworkInterface: private network interface.
type PrivateNetworkInterface struct {
	// ID: unique ID of the private network interface.
	ID string `json:"id"`

	// PrivateNetworkID: ID of the Private Network this interface is attached to.
	PrivateNetworkID string `json:"private_network_id"`

	// ProjectID: project ID the private network interface belongs to.
	ProjectID string `json:"project_id"`

	// ServerID: ID of the Instance this interface is attached to.
	ServerID string `json:"server_id"`

	// MacAddress: mAC address of the private network interface.
	MacAddress string `json:"mac_address"`

	// Status: current status of the private network interface.
	// Default value: unknown_status
	Status PrivateNetworkInterfaceStatus `json:"status"`

	// IPIDs: list of IP IDs attached to this interface.
	IPIDs []string `json:"ip_ids"`

	// Tags: tags associated with the private network interface.
	Tags []string `json:"tags"`

	// CreatedAt: creation timestamp of the private network interface.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update timestamp of the private network interface.
	UpdatedAt *time.Time `json:"updated_at"`
}

// RebootServerRequest: reboot server request.
type RebootServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to reboot.
	ServerID string `json:"-"`
}

// ResourceCounts: resource counts.
type ResourceCounts struct {
	// Servers: number of servers.
	Servers uint32 `json:"servers"`

	// GpuServers: number of GPU servers.
	GpuServers uint32 `json:"gpu_servers"`

	// ServersByType: map of server types with their counts.
	ServersByType map[string]uint32 `json:"servers_by_type"`

	// SecurityGroups: number of security groups.
	SecurityGroups uint32 `json:"security_groups"`

	// PlacementGroups: number of placement groups.
	PlacementGroups uint32 `json:"placement_groups"`

	// Snapshots: number of snapshots.
	Snapshots uint32 `json:"snapshots"`

	// Volumes: number of volumes.
	Volumes uint32 `json:"volumes"`

	// VolumesLSSD: number of local SSD volumes.
	VolumesLSSD uint32 `json:"volumes_l_ssd"`

	// VolumesLSSDTotalSize: total size of local SSD volumes in bytes.
	VolumesLSSDTotalSize uint64 `json:"volumes_l_ssd_total_size"`

	// VolumesScratch: number of scratch volumes.
	VolumesScratch uint32 `json:"volumes_scratch"`

	// PrivateNetworkInterfaces: number of private network interfaces.
	PrivateNetworkInterfaces uint32 `json:"private_network_interfaces"`
}

// Server: server.
type Server struct {
	// ID: unique ID of the server.
	ID string `json:"id"`

	// Name: name of the server.
	Name string `json:"name"`

	// ProjectID: project ID to which the server belongs.
	ProjectID string `json:"project_id"`

	// Tags: tags associated with the server.
	Tags []string `json:"tags"`

	// ServerType: type of the server.
	ServerType string `json:"server_type"`

	// PlacementGroupID: ID of the placement group the server belongs to.
	PlacementGroupID *string `json:"placement_group_id"`

	// Status: current status of the server.
	// Default value: unknown_status
	Status ServerStatus `json:"status"`

	// Volumes: list of volumes attached to the server.
	Volumes []*ServerVolume `json:"volumes"`

	// Filesystems: list of filesystems attached to the server.
	Filesystems []*ServerFilesystem `json:"filesystems"`

	// Architecture: architecture of the server.
	// Default value: unknown_architecture
	Architecture ServerArchitecture `json:"architecture"`

	// CreatedAt: creation timestamp of the server.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update timestamp of the server.
	UpdatedAt *time.Time `json:"updated_at"`

	// PrivateNetworkInterfaces: list of private network interfaces attached to the server.
	PrivateNetworkInterfaces []*ServerPrivateNetworkInterface `json:"private_network_interfaces"`

	// RescueMode: whether the server is in rescue mode.
	RescueMode bool `json:"rescue_mode"`

	// BootVolumeID: ID of the boot volume.
	BootVolumeID *string `json:"boot_volume_id"`

	// StatusDetail: detailed status information of the server.
	StatusDetail string `json:"status_detail"`

	// WindowsRdpPassword: encrypted RDP password for Windows servers. The encryption scheme is RSA-PKCS1-v1_5, using the public part of the SSH key supplied in `windows_rdp_ssh_key_id`.
	WindowsRdpPassword *ServerRDPPassword `json:"windows_rdp_password"`

	// PublicNetworkInterface: public network interface of the server.
	PublicNetworkInterface *ServerPublicNetworkInterface `json:"public_network_interface"`

	// Zone: zone in which the server is located.
	Zone scw.Zone `json:"zone"`
}

// SetSecurityGroupRulesRequest: set security group rules request.
type SetSecurityGroupRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// SecurityGroupID: ID of the security group to set rules for.
	SecurityGroupID string `json:"security_group_id"`

	// SecurityGroupRules: list of rules to set.
	SecurityGroupRules []*SecurityGroupRuleConfig `json:"security_group_rules"`
}

// SetServerCloudInitRequest: set server cloud init request.
type SetServerCloudInitRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: the ID of the server.
	ServerID string `json:"-"`

	// Content: the cloud-init configuration content.
	Content []byte `json:"content"`
}

// SetServerDefaultIPRequest: set server default ip request.
type SetServerDefaultIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to set the default IP for.
	ServerID string `json:"-"`

	// IPID: ID of the IP to set as default.
	IPID string `json:"ip_id"`
}

// SetTemplateCloudInitRequest: set template cloud init request.
type SetTemplateCloudInitRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: unique ID of the template.
	TemplateID string `json:"-"`

	// Content: cloud-init configuration content.
	Content []byte `json:"content"`
}

// SetTemplateUserDataRequest: set template user data request.
type SetTemplateUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: unique ID of the template.
	TemplateID string `json:"-"`

	// Key: key of the user data to set.
	Key string `json:"-"`

	// Content: content of the user data.
	Content []byte `json:"content"`
}

// SetUserDataRequest: set user data request.
type SetUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: the ID of the server.
	ServerID string `json:"-"`

	// Key: the key of the user data to set.
	Key string `json:"-"`

	// Content: the content to set for the user data.
	Content []byte `json:"content"`
}

// StartServerRequest: start server request.
type StartServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to start.
	ServerID string `json:"-"`
}

// StopAndDeleteServerRequest: stop and delete server request.
type StopAndDeleteServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to stop and delete.
	ServerID string `json:"-"`

	// DeleteAllIPs: whether to delete all IPs attached to the server.
	// Precisely one of DeleteAllIPs, DeleteIPIDs must be set.
	DeleteAllIPs *bool `json:"delete_all_ips,omitempty"`

	// DeleteIPIDs: list of IP IDs to delete.
	// Precisely one of DeleteAllIPs, DeleteIPIDs must be set.
	DeleteIPIDs *[]string `json:"delete_ip_ids,omitempty"`

	// DeleteAllVolumes: whether to delete all volumes attached to the server.
	// Precisely one of DeleteAllVolumes, DeleteVolumeIDs must be set.
	DeleteAllVolumes *bool `json:"delete_all_volumes,omitempty"`

	// DeleteVolumeIDs: list of volume IDs to delete.
	// Precisely one of DeleteAllVolumes, DeleteVolumeIDs must be set.
	DeleteVolumeIDs *[]string `json:"delete_volume_ids,omitempty"`

	// KeepAllPrivateNics: whether to keep all private network interfaces.
	// Precisely one of KeepAllPrivateNics, DeletePrivateNicIDs must be set.
	KeepAllPrivateNics *bool `json:"keep_all_private_nics,omitempty"`

	// DeletePrivateNicIDs: list of private network interface IDs to delete.
	// Precisely one of KeepAllPrivateNics, DeletePrivateNicIDs must be set.
	DeletePrivateNicIDs *[]string `json:"delete_private_nic_ids,omitempty"`
}

// StopServerRequest: stop server request.
type StopServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to stop.
	ServerID string `json:"-"`
}

// Template: template.
type Template struct {
	// ProjectID: project ID associated with the template.
	ProjectID string `json:"project_id"`

	// ID: unique ID of the template.
	ID string `json:"id"`

	// Name: name of the template.
	Name string `json:"name"`

	// Tags: tags associated with the template.
	Tags []string `json:"tags"`

	// ServerTags: tags associated with servers created from this template.
	ServerTags []string `json:"server_tags"`

	// ServerType: commercial type of the server defined by the template.
	ServerType string `json:"server_type"`

	// SecurityGroupID: security group ID associated with the template.
	SecurityGroupID *string `json:"security_group_id"`

	// PlacementGroupID: placement group ID associated with the template.
	PlacementGroupID *string `json:"placement_group_id"`

	// PublicIPV4Count: number of IPv4 public IPs to attach to servers created from this template.
	PublicIPV4Count uint32 `json:"public_ip_v4_count"`

	// PublicIPV6Count: number of IPv6 public IPs to attach to servers created from this template.
	PublicIPV6Count uint32 `json:"public_ip_v6_count"`

	// Volumes: list of volume templates used to create volumes for servers.
	Volumes []*CreateTemplateRequestVolumeTemplate `json:"volumes"`

	// PrivateNetworks: list of private network associated with the template.
	PrivateNetworks []*CreateTemplateRequestPrivateNetworkTemplate `json:"private_networks"`

	// FilesystemIDs: list of filesystem IDs associated with the template.
	FilesystemIDs []string `json:"filesystem_ids"`

	// CreatedAt: creation timestamp of the template.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update timestamp of the template.
	UpdatedAt *time.Time `json:"updated_at"`

	// WindowsRdpSSHKeyID: iAM ID of the SSH key used to encrypt the Windows `Administrator` password for RDP use.
	WindowsRdpSSHKeyID *string `json:"windows_rdp_ssh_key_id"`

	// Zone: zone in which the template is located.
	Zone scw.Zone `json:"zone"`
}

// UpdatePlacementGroupRequest: update placement group request.
type UpdatePlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PlacementGroupID: UUID of the placement group.
	PlacementGroupID string `json:"-"`

	// Name: name of the placement group.
	Name *string `json:"name,omitempty"`

	// PolicyType: policy type of the placement group.
	// Default value: unknown_policy_type
	PolicyType PlacementGroupPolicyType `json:"policy_type"`

	// Tags: tags of the placement group.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdatePrivateNetworkInterfaceRequest: update private network interface request.
type UpdatePrivateNetworkInterfaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PrivateNetworkInterfaceID: ID of the private network interface to update.
	PrivateNetworkInterfaceID string `json:"-"`

	// Tags: new tags to assign to the private network interface.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateSecurityGroupRequest: update security group request.
type UpdateSecurityGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// SecurityGroupID: ID of the security group to update.
	SecurityGroupID string `json:"-"`

	// Name: new name for the security group.
	Name *string `json:"name,omitempty"`

	// Description: new description for the security group.
	Description *string `json:"description,omitempty"`

	// DisableDefaultRules: whether to disable default rules.
	DisableDefaultRules *bool `json:"disable_default_rules,omitempty"`

	// Tags: new tags for the security group.
	Tags *[]string `json:"tags,omitempty"`

	// ProjectDefault: whether this should be the default security group for the project.
	ProjectDefault *bool `json:"project_default,omitempty"`

	// InboundDefaultAction: new default action for inbound rules.
	// Default value: unknown_action
	InboundDefaultAction SecurityGroupAction `json:"inbound_default_action"`

	// OutboundDefaultAction: new default action for outbound rules.
	// Default value: unknown_action
	OutboundDefaultAction SecurityGroupAction `json:"outbound_default_action"`

	// Stateless: whether the security group should be stateless.
	Stateless *bool `json:"stateless,omitempty"`
}

// UpdateSecurityGroupRuleRequest: update security group rule request.
type UpdateSecurityGroupRuleRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// SecurityGroupRuleID: ID of the rule to update.
	SecurityGroupRuleID string `json:"-"`

	// Protocol: new protocol for the rule.
	// Default value: unknown_protocol
	Protocol SecurityGroupRuleProtocol `json:"protocol"`

	// Direction: new direction for the rule.
	// Default value: unknown_direction
	Direction SecurityGroupRuleDirection `json:"direction"`

	// Action: new action for the rule.
	// Default value: unknown_action
	Action SecurityGroupRuleAction `json:"action"`

	// SourceIPRange: new source IP range for the rule.
	SourceIPRange *scw.IPNet `json:"source_ip_range,omitempty"`

	// DestinationIPRange: new destination IP range for the rule.
	DestinationIPRange *scw.IPNet `json:"destination_ip_range,omitempty"`

	// SourcePorts: new source port range for the rule.
	SourcePorts *SecurityGroupRulePortRange `json:"source_ports,omitempty"`

	// DestinationPorts: new destination port range for the rule.
	DestinationPorts *SecurityGroupRulePortRange `json:"destination_ports,omitempty"`

	// Position: new position for the rule.
	Position *int32 `json:"position,omitempty"`
}

// UpdateServerRequest: update server request.
type UpdateServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: ID of the server to update.
	ServerID string `json:"-"`

	// Name: new name for the server.
	Name *string `json:"name,omitempty"`

	// Tags: new tags for the server.
	Tags *[]string `json:"tags,omitempty"`

	// ServerType: new server type.
	ServerType *string `json:"server_type,omitempty"`

	// PlacementGroupID: new placement group ID.
	PlacementGroupID *string `json:"placement_group_id,omitempty"`

	// RescueMode: new rescue mode setting.
	RescueMode *bool `json:"rescue_mode,omitempty"`

	// BootVolumeID: new boot volume ID.
	BootVolumeID *string `json:"boot_volume_id,omitempty"`

	// WindowsRdpSSHKeyID: new IAM ID of the SSH key used to encrypt the Windows `Administrator` password for RDP use.
	WindowsRdpSSHKeyID *string `json:"windows_rdp_ssh_key_id,omitempty"`

	// Protected: protection status of the server.
	Protected *bool `json:"protected,omitempty"`

	// PublicNetworkInterface: new public network interface configuration.
	PublicNetworkInterface *UpdateServerRequestPublicNetworkInterface `json:"public_network_interface,omitempty"`
}

// UpdateTemplateRequest: update template request.
type UpdateTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// TemplateID: unique ID of the template to update.
	TemplateID string `json:"-"`

	// Name: new name for the template.
	Name *string `json:"name,omitempty"`

	// Tags: new tags for the template.
	Tags *[]string `json:"tags,omitempty"`

	// ServerTags: new server tags for the template.
	ServerTags *[]string `json:"server_tags,omitempty"`

	// ServerType: new server type for the template.
	ServerType *string `json:"server_type,omitempty"`

	// SecurityGroupID: new security group ID for the template.
	SecurityGroupID *string `json:"security_group_id,omitempty"`

	// PlacementGroupID: new placement group ID for the template.
	PlacementGroupID *string `json:"placement_group_id,omitempty"`

	// UpdateVolumes: updated volume templates for the template.
	UpdateVolumes *UpdateTemplateRequestUpdateVolumes `json:"update_volumes,omitempty"`

	// UpdatePrivateNetworks: updated private networks list for the template.
	UpdatePrivateNetworks *UpdateTemplateRequestUpdatePrivateNetworks `json:"update_private_networks,omitempty"`

	// FilesystemIDs: new list of filesystem IDs for the template.
	FilesystemIDs *[]string `json:"filesystem_ids,omitempty"`

	// PublicIPV4Count: new number of IPv4 public IPs to attach to servers.
	PublicIPV4Count *uint32 `json:"public_ip_v4_count,omitempty"`

	// PublicIPV6Count: new number of IPv6 public IPs to attach to servers.
	PublicIPV6Count *uint32 `json:"public_ip_v6_count,omitempty"`

	// WindowsRdpSSHKeyID: new IAM ID of the SSH key used to encrypt the Windows `Administrator` password for RDP use.
	WindowsRdpSSHKeyID *string `json:"windows_rdp_ssh_key_id,omitempty"`
}

// UserData: user data.
type UserData struct {
	// Key: the key of the user data.
	Key string `json:"key"`

	// Content: the content of the user data.
	Content []byte `json:"content"`
}

// This API allows you to manage Instance local and scratch volumes.
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
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3, scw.ZoneItMil1}
}

// GetServiceInfo:
func (s *API) GetServiceInfo(req *GetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetResourceCounts:
func (s *API) GetResourceCounts(req *GetResourceCountsRequest, opts ...scw.RequestOption) (*ResourceCounts, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/resource-counts",
		Query:  query,
	}

	var resp ResourceCounts

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListServers:
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

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "server_ids", req.ServerIDs)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "server_type", req.ServerType)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "security_group_ids", req.SecurityGroupIDs)
	parameter.AddToQuery(query, "placement_group_ids", req.PlacementGroupIDs)
	parameter.AddToQuery(query, "private_network_ids", req.PrivateNetworkIDs)
	parameter.AddToQuery(query, "mac_addresses", req.MacAddresses)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:  query,
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateServer:
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

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers",
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForServerRequest is used by WaitForServer method.
type WaitForServerRequest struct {
	Zone          scw.Zone
	ServerID      string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForServer waits for the Server to reach a terminal state.
func (s *API) WaitForServer(req *WaitForServerRequest, opts ...scw.RequestOption) (*Server, error) {
	timeout := defaultInstanceTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultInstanceRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[ServerStatus]struct{}{
		ServerStatusStarting:  {},
		ServerStatusStopping:  {},
		ServerStatusPausing:   {},
		ServerStatusRebooting: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetServer(&GetServerRequest{
				Zone:     req.Zone,
				ServerID: req.ServerID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTransient := transientStatuses[res.Status]

			return res, !isTransient, nil
		},
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
		Timeout:          timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for Server failed")
	}

	return res.(*Server), nil
}

// UpdateServer:
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
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

// DeleteServer:
func (s *API) DeleteServer(req *DeleteServerRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "delete_all_ips", req.DeleteAllIPs)
	parameter.AddToQuery(query, "delete_ip_ids", req.DeleteIPIDs)
	parameter.AddToQuery(query, "delete_all_volumes", req.DeleteAllVolumes)
	parameter.AddToQuery(query, "delete_volume_ids", req.DeleteVolumeIDs)
	parameter.AddToQuery(query, "keep_all_private_nics", req.KeepAllPrivateNics)
	parameter.AddToQuery(query, "delete_private_nic_ids", req.DeletePrivateNicIDs)

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
		Query:  query,
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListServerTypes:
func (s *API) ListServerTypes(req *ListServerTypesRequest, opts ...scw.RequestOption) (*ListServerTypesResponse, error) {
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
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/server-types",
		Query:  query,
	}

	var resp ListServerTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartServer:
func (s *API) StartServer(req *StartServerRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/start",
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

// RebootServer:
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/reboot",
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

// PauseServer:
func (s *API) PauseServer(req *PauseServerRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/pause",
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

// StopServer:
func (s *API) StopServer(req *StopServerRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/stop",
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

// StopAndDeleteServer:
func (s *API) StopAndDeleteServer(req *StopAndDeleteServerRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/stop-and-delete",
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

// AttachServerVolume:
func (s *API) AttachServerVolume(req *AttachServerVolumeRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/attach-volume",
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

// DetachServerVolume:
func (s *API) DetachServerVolume(req *DetachServerVolumeRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/detach-volume",
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

// AttachServerFileSystem:
func (s *API) AttachServerFileSystem(req *AttachServerFileSystemRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/attach-filesystem",
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

// DetachServerFileSystem:
func (s *API) DetachServerFileSystem(req *DetachServerFileSystemRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/detach-filesystem",
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

// AttachServerIP:
func (s *API) AttachServerIP(req *AttachServerIPRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/attach-ip",
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

// DetachServerIP:
func (s *API) DetachServerIP(req *DetachServerIPRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/detach-ip",
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

// SetServerDefaultIP:
func (s *API) SetServerDefaultIP(req *SetServerDefaultIPRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/set-default-ip",
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

// AttachServerPrivateNetworkInterface:
func (s *API) AttachServerPrivateNetworkInterface(req *AttachServerPrivateNetworkInterfaceRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/attach-private-network-interface",
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

// DetachServerPrivateNetworkInterface:
func (s *API) DetachServerPrivateNetworkInterface(req *DetachServerPrivateNetworkInterfaceRequest, opts ...scw.RequestOption) (*Server, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/detach-private-network-interface",
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

// ListPrivateNetworkInterfaces:
func (s *API) ListPrivateNetworkInterfaces(req *ListPrivateNetworkInterfacesRequest, opts ...scw.RequestOption) (*ListPrivateNetworkInterfacesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "server_ids", req.ServerIDs)
	parameter.AddToQuery(query, "private_network_ids", req.PrivateNetworkIDs)
	parameter.AddToQuery(query, "tags", req.Tags)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/private-network-interfaces",
		Query:  query,
	}

	var resp ListPrivateNetworkInterfacesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePrivateNetworkInterface:
func (s *API) CreatePrivateNetworkInterface(req *CreatePrivateNetworkInterfaceRequest, opts ...scw.RequestOption) (*PrivateNetworkInterface, error) {
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
		Method: "POST",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/private-network-interfaces",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNetworkInterface

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrivateNetworkInterface:
func (s *API) GetPrivateNetworkInterface(req *GetPrivateNetworkInterfaceRequest, opts ...scw.RequestOption) (*PrivateNetworkInterface, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkInterfaceID) == "" {
		return nil, errors.New("field PrivateNetworkInterfaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/private-network-interfaces/" + fmt.Sprint(req.PrivateNetworkInterfaceID) + "",
	}

	var resp PrivateNetworkInterface

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForPrivateNetworkInterfaceRequest is used by WaitForPrivateNetworkInterface method.
type WaitForPrivateNetworkInterfaceRequest struct {
	Zone                      scw.Zone
	PrivateNetworkInterfaceID string
	Timeout                   *time.Duration
	RetryInterval             *time.Duration
}

// WaitForPrivateNetworkInterface waits for the PrivateNetworkInterface to reach a terminal state.
func (s *API) WaitForPrivateNetworkInterface(req *WaitForPrivateNetworkInterfaceRequest, opts ...scw.RequestOption) (*PrivateNetworkInterface, error) {
	timeout := defaultInstanceTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultInstanceRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[PrivateNetworkInterfaceStatus]struct{}{
		PrivateNetworkInterfaceStatusAttaching: {},
		PrivateNetworkInterfaceStatusDetaching: {},
		PrivateNetworkInterfaceStatusSyncing:   {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetPrivateNetworkInterface(&GetPrivateNetworkInterfaceRequest{
				Zone:                      req.Zone,
				PrivateNetworkInterfaceID: req.PrivateNetworkInterfaceID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTransient := transientStatuses[res.Status]

			return res, !isTransient, nil
		},
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
		Timeout:          timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for PrivateNetworkInterface failed")
	}

	return res.(*PrivateNetworkInterface), nil
}

// UpdatePrivateNetworkInterface:
func (s *API) UpdatePrivateNetworkInterface(req *UpdatePrivateNetworkInterfaceRequest, opts ...scw.RequestOption) (*PrivateNetworkInterface, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkInterfaceID) == "" {
		return nil, errors.New("field PrivateNetworkInterfaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/private-network-interfaces/" + fmt.Sprint(req.PrivateNetworkInterfaceID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNetworkInterface

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePrivateNetworkInterface:
func (s *API) DeletePrivateNetworkInterface(req *DeletePrivateNetworkInterfaceRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkInterfaceID) == "" {
		return errors.New("field PrivateNetworkInterfaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/private-network-interfaces/" + fmt.Sprint(req.PrivateNetworkInterfaceID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListPlacementGroups:
func (s *API) ListPlacementGroups(req *ListPlacementGroupsRequest, opts ...scw.RequestOption) (*ListPlacementGroupsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "placement_group_ids", req.PlacementGroupIDs)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/placement-groups",
		Query:  query,
	}

	var resp ListPlacementGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePlacementGroup:
func (s *API) CreatePlacementGroup(req *CreatePlacementGroupRequest, opts ...scw.RequestOption) (*PlacementGroup, error) {
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
		Method: "POST",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/placement-groups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PlacementGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPlacementGroup:
func (s *API) GetPlacementGroup(req *GetPlacementGroupRequest, opts ...scw.RequestOption) (*PlacementGroup, error) {
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
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/placement-groups/" + fmt.Sprint(req.PlacementGroupID) + "",
	}

	var resp PlacementGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePlacementGroup:
func (s *API) UpdatePlacementGroup(req *UpdatePlacementGroupRequest, opts ...scw.RequestOption) (*PlacementGroup, error) {
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
		Method: "PATCH",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/placement-groups/" + fmt.Sprint(req.PlacementGroupID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PlacementGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePlacementGroup:
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
		Method: "DELETE",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/placement-groups/" + fmt.Sprint(req.PlacementGroupID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListSecurityGroups:
func (s *API) ListSecurityGroups(req *ListSecurityGroupsRequest, opts ...scw.RequestOption) (*ListSecurityGroupsResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "security_group_ids", req.SecurityGroupIDs)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/security-groups",
		Query:  query,
	}

	var resp ListSecurityGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSecurityGroup:
func (s *API) CreateSecurityGroup(req *CreateSecurityGroupRequest, opts ...scw.RequestOption) (*SecurityGroup, error) {
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
		Method: "POST",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/security-groups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecurityGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSecurityGroup:
func (s *API) GetSecurityGroup(req *GetSecurityGroupRequest, opts ...scw.RequestOption) (*SecurityGroup, error) {
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
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/security-groups/" + fmt.Sprint(req.SecurityGroupID) + "",
	}

	var resp SecurityGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSecurityGroup:
func (s *API) UpdateSecurityGroup(req *UpdateSecurityGroupRequest, opts ...scw.RequestOption) (*SecurityGroup, error) {
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
		Method: "PATCH",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/security-groups/" + fmt.Sprint(req.SecurityGroupID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecurityGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSecurityGroup:
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
		Method: "DELETE",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/security-groups/" + fmt.Sprint(req.SecurityGroupID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// AddSecurityGroupRules:
func (s *API) AddSecurityGroupRules(req *AddSecurityGroupRulesRequest, opts ...scw.RequestOption) (*AddSecurityGroupRulesResponse, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/security-group-rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddSecurityGroupRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetSecurityGroupRules:
func (s *API) SetSecurityGroupRules(req *SetSecurityGroupRulesRequest, opts ...scw.RequestOption) (*SecurityGroup, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/security-group-rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecurityGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSecurityGroupRule:
func (s *API) UpdateSecurityGroupRule(req *UpdateSecurityGroupRuleRequest, opts ...scw.RequestOption) (*SecurityGroup, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.SecurityGroupRuleID) == "" {
		return nil, errors.New("field SecurityGroupRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/security-group-rules/" + fmt.Sprint(req.SecurityGroupRuleID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecurityGroup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSecurityGroupRules:
func (s *API) DeleteSecurityGroupRules(req *DeleteSecurityGroupRulesRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/security-group-rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListUserDataKeys:
func (s *API) ListUserDataKeys(req *ListUserDataKeysRequest, opts ...scw.RequestOption) (*ListUserDataKeysResponse, error) {
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
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user-data",
		Query:  query,
	}

	var resp ListUserDataKeysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUserData:
func (s *API) GetUserData(req *GetUserDataRequest, opts ...scw.RequestOption) (*UserData, error) {
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

	if fmt.Sprint(req.Key) == "" {
		return nil, errors.New("field Key cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user-data/" + fmt.Sprint(req.Key) + "",
	}

	var resp UserData

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetUserData:
func (s *API) SetUserData(req *SetUserDataRequest, opts ...scw.RequestOption) error {
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
		Method: "PUT",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user-data/" + fmt.Sprint(req.Key) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteUserData:
func (s *API) DeleteUserData(req *DeleteUserDataRequest, opts ...scw.RequestOption) error {
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
		Method: "DELETE",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user-data/" + fmt.Sprint(req.Key) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetServerCloudInit:
func (s *API) GetServerCloudInit(req *GetServerCloudInitRequest, opts ...scw.RequestOption) (*UserData, error) {
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
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user-data/cloud-init",
	}

	var resp UserData

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetServerCloudInit:
func (s *API) SetServerCloudInit(req *SetServerCloudInitRequest, opts ...scw.RequestOption) error {
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
		Method: "PUT",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/user-data/cloud-init",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListTemplates:
func (s *API) ListTemplates(req *ListTemplatesRequest, opts ...scw.RequestOption) (*ListTemplatesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "template_ids", req.TemplateIDs)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "server_tags", req.ServerTags)
	parameter.AddToQuery(query, "security_group_ids", req.SecurityGroupIDs)
	parameter.AddToQuery(query, "placement_group_ids", req.PlacementGroupIDs)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates",
		Query:  query,
	}

	var resp ListTemplatesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateTemplate:
func (s *API) CreateTemplate(req *CreateTemplateRequest, opts ...scw.RequestOption) (*Template, error) {
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
		Method: "POST",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Template

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTemplate:
func (s *API) GetTemplate(req *GetTemplateRequest, opts ...scw.RequestOption) (*Template, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return nil, errors.New("field TemplateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates/" + fmt.Sprint(req.TemplateID) + "",
	}

	var resp Template

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateTemplate:
func (s *API) UpdateTemplate(req *UpdateTemplateRequest, opts ...scw.RequestOption) (*Template, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return nil, errors.New("field TemplateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates/" + fmt.Sprint(req.TemplateID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Template

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteTemplate:
func (s *API) DeleteTemplate(req *DeleteTemplateRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return errors.New("field TemplateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates/" + fmt.Sprint(req.TemplateID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListTemplateUserDataKeys:
func (s *API) ListTemplateUserDataKeys(req *ListTemplateUserDataKeysRequest, opts ...scw.RequestOption) (*ListTemplateUserDataKeysResponse, error) {
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
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return nil, errors.New("field TemplateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates/" + fmt.Sprint(req.TemplateID) + "/user-data",
		Query:  query,
	}

	var resp ListTemplateUserDataKeysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetTemplateUserData:
func (s *API) GetTemplateUserData(req *GetTemplateUserDataRequest, opts ...scw.RequestOption) (*UserData, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return nil, errors.New("field TemplateID cannot be empty in request")
	}

	if fmt.Sprint(req.Key) == "" {
		return nil, errors.New("field Key cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates/" + fmt.Sprint(req.TemplateID) + "/user-data/" + fmt.Sprint(req.Key) + "",
	}

	var resp UserData

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetTemplateUserData:
func (s *API) SetTemplateUserData(req *SetTemplateUserDataRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return errors.New("field TemplateID cannot be empty in request")
	}

	if fmt.Sprint(req.Key) == "" {
		return errors.New("field Key cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates/" + fmt.Sprint(req.TemplateID) + "/user-data/" + fmt.Sprint(req.Key) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteTemplateUserData:
func (s *API) DeleteTemplateUserData(req *DeleteTemplateUserDataRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return errors.New("field TemplateID cannot be empty in request")
	}

	if fmt.Sprint(req.Key) == "" {
		return errors.New("field Key cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates/" + fmt.Sprint(req.TemplateID) + "/user-data/" + fmt.Sprint(req.Key) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetTemplateCloudInit:
func (s *API) GetTemplateCloudInit(req *GetTemplateCloudInitRequest, opts ...scw.RequestOption) (*UserData, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return nil, errors.New("field TemplateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates/" + fmt.Sprint(req.TemplateID) + "/user-data/cloud-init",
	}

	var resp UserData

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetTemplateCloudInit:
func (s *API) SetTemplateCloudInit(req *SetTemplateCloudInitRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return errors.New("field TemplateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates/" + fmt.Sprint(req.TemplateID) + "/user-data/cloud-init",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CheckTemplate:
func (s *API) CheckTemplate(req *CheckTemplateRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return errors.New("field TemplateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates/" + fmt.Sprint(req.TemplateID) + "/check",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateServerFromTemplate:
func (s *API) CreateServerFromTemplate(req *CreateServerFromTemplateRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.TemplateID) == "" {
		return nil, errors.New("field TemplateID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/instance/v2alpha1/zones/" + fmt.Sprint(req.Zone) + "/templates/" + fmt.Sprint(req.TemplateID) + "/create-server",
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

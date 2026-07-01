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
	ListPlacementGroupsRequestOrderByCreatedAtDesc = ListPlacementGroupsRequestOrderBy("created_at_desc")
	ListPlacementGroupsRequestOrderByCreatedAtAsc  = ListPlacementGroupsRequestOrderBy("created_at_asc")
	ListPlacementGroupsRequestOrderByUpdatedAtDesc = ListPlacementGroupsRequestOrderBy("updated_at_desc")
	ListPlacementGroupsRequestOrderByUpdatedAtAsc  = ListPlacementGroupsRequestOrderBy("updated_at_asc")
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
	ListPrivateNetworkInterfacesRequestOrderByCreatedAtDesc = ListPrivateNetworkInterfacesRequestOrderBy("created_at_desc")
	ListPrivateNetworkInterfacesRequestOrderByCreatedAtAsc  = ListPrivateNetworkInterfacesRequestOrderBy("created_at_asc")
	ListPrivateNetworkInterfacesRequestOrderByUpdatedAtDesc = ListPrivateNetworkInterfacesRequestOrderBy("updated_at_desc")
	ListPrivateNetworkInterfacesRequestOrderByUpdatedAtAsc  = ListPrivateNetworkInterfacesRequestOrderBy("updated_at_asc")
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
	ListSecurityGroupsRequestOrderByCreatedAtDesc = ListSecurityGroupsRequestOrderBy("created_at_desc")
	ListSecurityGroupsRequestOrderByCreatedAtAsc  = ListSecurityGroupsRequestOrderBy("created_at_asc")
	ListSecurityGroupsRequestOrderByUpdatedAtDesc = ListSecurityGroupsRequestOrderBy("updated_at_desc")
	ListSecurityGroupsRequestOrderByUpdatedAtAsc  = ListSecurityGroupsRequestOrderBy("updated_at_asc")
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
	PlacementGroupPolicyTypeUnknownPolicyType = PlacementGroupPolicyType("unknown_policy_type")
	PlacementGroupPolicyTypeLowLatency        = PlacementGroupPolicyType("low_latency")
	PlacementGroupPolicyTypeMaxAvailability   = PlacementGroupPolicyType("max_availability")
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
	PrivateNetworkInterfaceStatusUnknownStatus = PrivateNetworkInterfaceStatus("unknown_status")
	PrivateNetworkInterfaceStatusAvailable     = PrivateNetworkInterfaceStatus("available")
	PrivateNetworkInterfaceStatusAttaching     = PrivateNetworkInterfaceStatus("attaching")
	PrivateNetworkInterfaceStatusDetaching     = PrivateNetworkInterfaceStatus("detaching")
	PrivateNetworkInterfaceStatusSyncing       = PrivateNetworkInterfaceStatus("syncing")
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
	SecurityGroupActionUnknownAction = SecurityGroupAction("unknown_action")
	SecurityGroupActionAccept        = SecurityGroupAction("accept")
	SecurityGroupActionDrop          = SecurityGroupAction("drop")
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
	SecurityGroupRuleActionUnknownAction = SecurityGroupRuleAction("unknown_action")
	SecurityGroupRuleActionAccept        = SecurityGroupRuleAction("accept")
	SecurityGroupRuleActionDrop          = SecurityGroupRuleAction("drop")
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
	SecurityGroupRuleDirectionUnknownDirection = SecurityGroupRuleDirection("unknown_direction")
	SecurityGroupRuleDirectionInbound          = SecurityGroupRuleDirection("inbound")
	SecurityGroupRuleDirectionOutbound         = SecurityGroupRuleDirection("outbound")
	SecurityGroupRuleDirectionBoth             = SecurityGroupRuleDirection("both")
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
	SecurityGroupRuleProtocolUnknownProtocol = SecurityGroupRuleProtocol("unknown_protocol")
	SecurityGroupRuleProtocolTCP             = SecurityGroupRuleProtocol("tcp")
	SecurityGroupRuleProtocolUDP             = SecurityGroupRuleProtocol("udp")
	SecurityGroupRuleProtocolIcmp            = SecurityGroupRuleProtocol("icmp")
	SecurityGroupRuleProtocolAny             = SecurityGroupRuleProtocol("any")
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
	ServerArchitectureUnknownArchitecture = ServerArchitecture("unknown_architecture")
	ServerArchitectureX86_64              = ServerArchitecture("x86_64")
	ServerArchitectureAarch64             = ServerArchitecture("aarch64")
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
	ServerIPStatusUnknownStatus = ServerIPStatus("unknown_status")
	ServerIPStatusDetached      = ServerIPStatus("detached")
	ServerIPStatusAttached      = ServerIPStatus("attached")
	ServerIPStatusPending       = ServerIPStatus("pending")
	ServerIPStatusError         = ServerIPStatus("error")
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
	ServerPrivateNetworkInterfaceStatusUnknownStatus = ServerPrivateNetworkInterfaceStatus("unknown_status")
	ServerPrivateNetworkInterfaceStatusAvailable     = ServerPrivateNetworkInterfaceStatus("available")
	ServerPrivateNetworkInterfaceStatusAttaching     = ServerPrivateNetworkInterfaceStatus("attaching")
	ServerPrivateNetworkInterfaceStatusDetaching     = ServerPrivateNetworkInterfaceStatus("detaching")
	ServerPrivateNetworkInterfaceStatusSyncing       = ServerPrivateNetworkInterfaceStatus("syncing")
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
	ServerPublicNetworkInterfaceStatusUnknownStatus = ServerPublicNetworkInterfaceStatus("unknown_status")
	ServerPublicNetworkInterfaceStatusAvailable     = ServerPublicNetworkInterfaceStatus("available")
	ServerPublicNetworkInterfaceStatusSyncing       = ServerPublicNetworkInterfaceStatus("syncing")
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
	ServerStatusUnknownStatus = ServerStatus("unknown_status")
	ServerStatusStarted       = ServerStatus("started")
	ServerStatusStopped       = ServerStatus("stopped")
	ServerStatusPaused        = ServerStatus("paused")
	ServerStatusStarting      = ServerStatus("starting")
	ServerStatusStopping      = ServerStatus("stopping")
	ServerStatusPausing       = ServerStatus("pausing")
	ServerStatusLocked        = ServerStatus("locked")
	ServerStatusRebooting     = ServerStatus("rebooting")
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
	ServerTypeArchitectureUnknownArchitecture = ServerTypeArchitecture("unknown_architecture")
	ServerTypeArchitectureX86_64              = ServerTypeArchitecture("x86_64")
	ServerTypeArchitectureAarch64             = ServerTypeArchitecture("aarch64")
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
	ServerTypeAvailabilityUnknownAvailability = ServerTypeAvailability("unknown_availability")
	ServerTypeAvailabilityAvailable           = ServerTypeAvailability("available")
	ServerTypeAvailabilityLowStock            = ServerTypeAvailability("low_stock")
	ServerTypeAvailabilityOutOfStock          = ServerTypeAvailability("out_of_stock")
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
	ServerVolumeVolumeTypeUnknownVolumeType = ServerVolumeVolumeType("unknown_volume_type")
	ServerVolumeVolumeTypeLSSD              = ServerVolumeVolumeType("l_ssd")
	ServerVolumeVolumeTypeSbs               = ServerVolumeVolumeType("sbs")
	ServerVolumeVolumeTypeScratch           = ServerVolumeVolumeType("scratch")
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
	Start uint32 `json:"start"`

	End uint32 `json:"end"`
}

// CreateServerRequestBookIP: create server request book ip.
type CreateServerRequestBookIP struct {
	// Type: default value: unknown_ip_type
	Type CreateServerRequestBookIPIPType `json:"type"`

	Tags []string `json:"tags"`
}

// SecurityGroupRule: security group rule.
type SecurityGroupRule struct {
	ID string `json:"id"`

	// Protocol: default value: unknown_protocol
	Protocol SecurityGroupRuleProtocol `json:"protocol"`

	// Direction: default value: unknown_direction
	Direction SecurityGroupRuleDirection `json:"direction"`

	// Action: default value: unknown_action
	Action SecurityGroupRuleAction `json:"action"`

	SourceIPRange scw.IPNet `json:"source_ip_range"`

	DestinationIPRange scw.IPNet `json:"destination_ip_range"`

	SourcePorts *SecurityGroupRulePortRange `json:"source_ports"`

	DestinationPorts *SecurityGroupRulePortRange `json:"destination_ports"`

	Position uint32 `json:"position"`
}

// CreateServerRequestServerIP: create server request server ip.
type CreateServerRequestServerIP struct {
	// Precisely one of IpamIPID, NewIP must be set.
	IpamIPID *string `json:"ipam_ip_id,omitempty"`

	// Precisely one of IpamIPID, NewIP must be set.
	NewIP *CreateServerRequestBookIP `json:"new_ip,omitempty"`
}

// CreateServerRequestCreateVolume: create server request create volume.
type CreateServerRequestCreateVolume struct {
	Name string `json:"name"`

	Tags []string `json:"tags"`

	Size *scw.Size `json:"size"`

	// Precisely one of BaseSnapshotID, ImageLabel must be set.
	BaseSnapshotID *string `json:"base_snapshot_id,omitempty"`

	// Precisely one of BaseSnapshotID, ImageLabel must be set.
	ImageLabel *string `json:"image_label,omitempty"`

	PerfIops *uint32 `json:"perf_iops"`
}

// ServerTypeGpuInfo: server type gpu info.
type ServerTypeGpuInfo struct {
	Manufacturer string `json:"manufacturer"`

	Name string `json:"name"`

	Memory scw.Size `json:"memory"`
}

// ServerTypeLimits: server type limits.
type ServerTypeLimits struct {
	PrivateNetworkCount uint32 `json:"private_network_count"`

	FileSystemCount uint32 `json:"file_system_count"`

	PrivateNetworkBandwidth uint64 `json:"private_network_bandwidth"`

	BlockBandwidth uint64 `json:"block_bandwidth"`

	InternetBandwidth uint64 `json:"internet_bandwidth"`

	LSSDSize scw.Size `json:"l_ssd_size"`

	ScratchSize scw.Size `json:"scratch_size"`

	IPCount uint32 `json:"ip_count"`

	VolumeCount uint32 `json:"volume_count"`

	ScratchVolumesCount uint32 `json:"scratch_volumes_count"`
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
	PrivateNetworkID string `json:"private_network_id"`
}

// CreateTemplateRequestVolumeTemplate: create template request volume template.
type CreateTemplateRequestVolumeTemplate struct {
	// VolumeType: default value: unknown_volume_type
	VolumeType CreateServerRequestServerVolumeVolumeType `json:"volume_type"`

	Name string `json:"name"`

	Tags []string `json:"tags"`

	Size *scw.Size `json:"size"`

	// Precisely one of BaseSnapshotID, ImageLabel must be set.
	BaseSnapshotID *string `json:"base_snapshot_id,omitempty"`

	// Precisely one of BaseSnapshotID, ImageLabel must be set.
	ImageLabel *string `json:"image_label,omitempty"`

	PerfIops *uint32 `json:"perf_iops"`
}

// SecurityGroupRuleConfig: security group rule config.
type SecurityGroupRuleConfig struct {
	// Protocol: default value: unknown_protocol
	Protocol SecurityGroupRuleProtocol `json:"protocol"`

	// Direction: default value: unknown_direction
	Direction SecurityGroupRuleDirection `json:"direction"`

	// Action: default value: unknown_action
	Action SecurityGroupRuleAction `json:"action"`

	SourceIPRange scw.IPNet `json:"source_ip_range"`

	DestinationIPRange scw.IPNet `json:"destination_ip_range"`

	SourcePorts *SecurityGroupRulePortRange `json:"source_ports"`

	DestinationPorts *SecurityGroupRulePortRange `json:"destination_ports"`

	Position int32 `json:"position"`
}

// SecurityGroup: security group.
type SecurityGroup struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	ProjectID string `json:"project_id"`

	Tags []string `json:"tags"`

	DisableDefaultRules bool `json:"disable_default_rules"`

	ProjectDefault bool `json:"project_default"`

	// InboundDefaultAction: default value: unknown_action
	InboundDefaultAction SecurityGroupAction `json:"inbound_default_action"`

	// OutboundDefaultAction: default value: unknown_action
	OutboundDefaultAction SecurityGroupAction `json:"outbound_default_action"`

	Stateless bool `json:"stateless"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	DefaultRules []*SecurityGroupRule `json:"default_rules"`

	Rules []*SecurityGroupRule `json:"rules"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
}

// CreateServerRequestPublicNetworkInterface: create server request public network interface.
type CreateServerRequestPublicNetworkInterface struct {
	SecurityGroupID *string `json:"security_group_id"`

	IPs []*CreateServerRequestServerIP `json:"ips"`
}

// CreateServerRequestServerVolume: create server request server volume.
type CreateServerRequestServerVolume struct {
	// VolumeType: default value: unknown_volume_type
	VolumeType CreateServerRequestServerVolumeVolumeType `json:"volume_type"`

	// Precisely one of VolumeID, NewVolume must be set.
	VolumeID *string `json:"volume_id,omitempty"`

	// Precisely one of VolumeID, NewVolume must be set.
	NewVolume *CreateServerRequestCreateVolume `json:"new_volume,omitempty"`
}

// PlacementGroup: placement group.
type PlacementGroup struct {
	ProjectID string `json:"project_id"`

	Name string `json:"name"`

	ID string `json:"id"`

	// PolicyType: default value: unknown_policy_type
	PolicyType PlacementGroupPolicyType `json:"policy_type"`

	Tags []string `json:"tags"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
}

// PrivateNetworkInterface: private network interface.
type PrivateNetworkInterface struct {
	ID string `json:"id"`

	PrivateNetworkID string `json:"private_network_id"`

	ProjectID string `json:"project_id"`

	ServerID string `json:"server_id"`

	SecurityGroupID string `json:"security_group_id"`

	MacAddress string `json:"mac_address"`

	// Status: default value: unknown_status
	Status PrivateNetworkInterfaceStatus `json:"status"`

	IPIDs []string `json:"ip_ids"`

	Tags []string `json:"tags"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// SecurityGroupSummary: security group summary.
type SecurityGroupSummary struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	ProjectID string `json:"project_id"`

	Tags []string `json:"tags"`

	DisableDefaultRules bool `json:"disable_default_rules"`

	ProjectDefault bool `json:"project_default"`

	// InboundDefaultAction: default value: unknown_action
	InboundDefaultAction SecurityGroupAction `json:"inbound_default_action"`

	// OutboundDefaultAction: default value: unknown_action
	OutboundDefaultAction SecurityGroupAction `json:"outbound_default_action"`

	Stateless bool `json:"stateless"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// ServerType: server type.
type ServerType struct {
	Name string `json:"name"`

	VcpuCount uint32 `json:"vcpu_count"`

	GpuCount uint32 `json:"gpu_count"`

	Memory scw.Size `json:"memory"`

	// Architecture: default value: unknown_architecture
	Architecture ServerTypeArchitecture `json:"architecture"`

	// Availability: default value: unknown_availability
	Availability ServerTypeAvailability `json:"availability"`

	Limits *ServerTypeLimits `json:"limits"`

	GpuInfo *ServerTypeGpuInfo `json:"gpu_info"`

	EndOfService bool `json:"end_of_service"`
}

// ServerSummary: server summary.
type ServerSummary struct {
	ProjectID string `json:"project_id"`

	ID string `json:"id"`

	Name string `json:"name"`

	Tags []string `json:"tags"`

	ServerType string `json:"server_type"`

	PlacementGroupID *string `json:"placement_group_id"`

	// Status: default value: unknown_status
	Status ServerStatus `json:"status"`

	// Architecture: default value: unknown_architecture
	Architecture ServerArchitecture `json:"architecture"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	RescueMode bool `json:"rescue_mode"`
}

// TemplateSummary: template summary.
type TemplateSummary struct {
	ProjectID string `json:"project_id"`

	ID string `json:"id"`

	Name string `json:"name"`

	Tags []string `json:"tags"`

	ServerTags []string `json:"server_tags"`

	ServerType string `json:"server_type"`

	SecurityGroupID *string `json:"security_group_id"`

	PlacementGroupID *string `json:"placement_group_id"`

	PublicIPV4Count uint32 `json:"public_ip_v4_count"`

	PublicIPV6Count uint32 `json:"public_ip_v6_count"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	FilesystemIDs []string `json:"filesystem_ids"`

	// Zone: zone to target. If none is passed will use default zone from the config.
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
	SecurityGroupID *string `json:"security_group_id"`
}

// UpdateTemplateRequestUpdatePrivateNetworks: update template request update private networks.
type UpdateTemplateRequestUpdatePrivateNetworks struct {
	PrivateNetworks []*CreateTemplateRequestPrivateNetworkTemplate `json:"private_networks"`
}

// UpdateTemplateRequestUpdateVolumes: update template request update volumes.
type UpdateTemplateRequestUpdateVolumes struct {
	Volumes []*CreateTemplateRequestVolumeTemplate `json:"volumes"`
}

// AddSecurityGroupRulesRequest: add security group rules request.
type AddSecurityGroupRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"security_group_id"`

	SecurityGroupRules []*SecurityGroupRuleConfig `json:"security_group_rules"`
}

// AddSecurityGroupRulesResponse: add security group rules response.
type AddSecurityGroupRulesResponse struct {
	SecurityGroup *SecurityGroup `json:"security_group"`

	AddedRules []*SecurityGroupRule `json:"added_rules"`
}

// AttachServerFileSystemRequest: attach server file system request.
type AttachServerFileSystemRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	FilesystemID string `json:"filesystem_id"`
}

// AttachServerIPRequest: attach server ip request.
type AttachServerIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	IPID string `json:"ip_id"`

	Default bool `json:"default"`

	MoveAllowed bool `json:"move_allowed"`
}

// AttachServerPrivateNetworkInterfaceRequest: attach server private network interface request.
type AttachServerPrivateNetworkInterfaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	PrivateNetworkInterfaceID string `json:"private_network_interface_id"`
}

// AttachServerVolumeRequest: attach server volume request.
type AttachServerVolumeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	VolumeID string `json:"volume_id"`

	// VolumeType: default value: unknown_volume_type
	VolumeType ServerVolumeVolumeType `json:"volume_type"`

	BootVolume bool `json:"boot_volume"`
}

// CheckTemplateRequest: check template request.
type CheckTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	TemplateID string `json:"-"`
}

// CreatePlacementGroupRequest: create placement group request.
type CreatePlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ProjectID string `json:"project_id"`

	Name string `json:"name"`

	// PolicyType: default value: unknown_policy_type
	PolicyType PlacementGroupPolicyType `json:"policy_type"`

	Tags []string `json:"tags"`
}

// CreatePrivateNetworkInterfaceRequest: create private network interface request.
type CreatePrivateNetworkInterfaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PrivateNetworkID string `json:"private_network_id"`

	ProjectID string `json:"project_id"`

	ServerID *string `json:"server_id,omitempty"`

	SecurityGroupID *string `json:"security_group_id,omitempty"`

	IPIDs []string `json:"ip_ids"`

	Tags []string `json:"tags"`
}

// CreateSecurityGroupRequest: create security group request.
type CreateSecurityGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	Name string `json:"name"`

	Description string `json:"description"`

	DisableDefaultRules bool `json:"disable_default_rules"`

	ProjectID string `json:"project_id"`

	Tags []string `json:"tags"`

	ProjectDefault bool `json:"project_default"`

	// InboundDefaultAction: default value: unknown_action
	InboundDefaultAction SecurityGroupAction `json:"inbound_default_action"`

	// OutboundDefaultAction: default value: unknown_action
	OutboundDefaultAction SecurityGroupAction `json:"outbound_default_action"`

	Stateless bool `json:"stateless"`
}

// CreateServerFromTemplateRequest: create server from template request.
type CreateServerFromTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	TemplateID string `json:"-"`

	Name string `json:"name"`
}

// CreateServerRequest: create server request.
type CreateServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ProjectID string `json:"project_id"`

	Name string `json:"name"`

	Tags []string `json:"tags"`

	ServerType string `json:"server_type"`

	PlacementGroupID *string `json:"placement_group_id,omitempty"`

	Volumes []*CreateServerRequestServerVolume `json:"volumes"`

	WindowsRdpSSHKeyID *string `json:"windows_rdp_ssh_key_id,omitempty"`

	PublicNetworkInterface *CreateServerRequestPublicNetworkInterface `json:"public_network_interface,omitempty"`
}

// CreateTemplateRequest: create template request.
type CreateTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ProjectID string `json:"project_id"`

	Name string `json:"name"`

	Tags []string `json:"tags"`

	ServerTags []string `json:"server_tags"`

	ServerType string `json:"server_type"`

	SecurityGroupID *string `json:"security_group_id,omitempty"`

	PlacementGroupID *string `json:"placement_group_id,omitempty"`

	Volumes []*CreateTemplateRequestVolumeTemplate `json:"volumes"`

	PrivateNetworks []*CreateTemplateRequestPrivateNetworkTemplate `json:"private_networks"`

	PublicIPV4Count uint32 `json:"public_ip_v4_count"`

	PublicIPV6Count uint32 `json:"public_ip_v6_count"`

	WindowsRdpSSHKeyID *string `json:"windows_rdp_ssh_key_id,omitempty"`

	FilesystemIDs []string `json:"filesystem_ids"`
}

// DeletePlacementGroupRequest: delete placement group request.
type DeletePlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PlacementGroupID string `json:"-"`
}

// DeletePrivateNetworkInterfaceRequest: delete private network interface request.
type DeletePrivateNetworkInterfaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PrivateNetworkInterfaceID string `json:"-"`
}

// DeleteSecurityGroupRequest: delete security group request.
type DeleteSecurityGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`
}

// DeleteSecurityGroupRulesRequest: delete security group rules request.
type DeleteSecurityGroupRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SecurityGroupRuleIDs []string `json:"security_group_rule_ids"`
}

// DeleteServerRequest: delete server request.
type DeleteServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	// Precisely one of DeleteAllIPs, DeleteIPIDs must be set.
	DeleteAllIPs *bool `json:"delete_all_ips,omitempty"`

	// Precisely one of DeleteAllIPs, DeleteIPIDs must be set.
	DeleteIPIDs *[]string `json:"delete_ip_ids,omitempty"`

	// Precisely one of DeleteAllVolumes, DeleteVolumeIDs must be set.
	DeleteAllVolumes *bool `json:"delete_all_volumes,omitempty"`

	// Precisely one of DeleteAllVolumes, DeleteVolumeIDs must be set.
	DeleteVolumeIDs *[]string `json:"delete_volume_ids,omitempty"`

	// Precisely one of KeepAllPrivateNics, DeletePrivateNicIDs must be set.
	KeepAllPrivateNics *bool `json:"keep_all_private_nics,omitempty"`

	// Precisely one of KeepAllPrivateNics, DeletePrivateNicIDs must be set.
	DeletePrivateNicIDs *[]string `json:"delete_private_nic_ids,omitempty"`
}

// DeleteTemplateRequest: delete template request.
type DeleteTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	TemplateID string `json:"-"`
}

// DeleteTemplateUserDataRequest: delete template user data request.
type DeleteTemplateUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	TemplateID string `json:"-"`

	Key string `json:"-"`
}

// DeleteUserDataRequest: delete user data request.
type DeleteUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	Key string `json:"-"`
}

// DetachServerFileSystemRequest: detach server file system request.
type DetachServerFileSystemRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	FilesystemID string `json:"filesystem_id"`
}

// DetachServerIPRequest: detach server ip request.
type DetachServerIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	IPID string `json:"ip_id"`
}

// DetachServerPrivateNetworkInterfaceRequest: detach server private network interface request.
type DetachServerPrivateNetworkInterfaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	PrivateNetworkInterfaceID string `json:"private_network_interface_id"`
}

// DetachServerVolumeRequest: detach server volume request.
type DetachServerVolumeRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	VolumeID string `json:"volume_id"`
}

// GetPlacementGroupRequest: get placement group request.
type GetPlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PlacementGroupID string `json:"-"`
}

// GetPrivateNetworkInterfaceRequest: get private network interface request.
type GetPrivateNetworkInterfaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PrivateNetworkInterfaceID string `json:"-"`
}

// GetResourceCountsRequest: get resource counts request.
type GetResourceCountsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Precisely one of OrganizationID, ProjectID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// Precisely one of OrganizationID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`
}

// GetSecurityGroupRequest: get security group request.
type GetSecurityGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`
}

// GetServerCloudInitRequest: get server cloud init request.
type GetServerCloudInitRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

// GetServerRequest: get server request.
type GetServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

// GetTemplateCloudInitRequest: get template cloud init request.
type GetTemplateCloudInitRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	TemplateID string `json:"-"`
}

// GetTemplateRequest: get template request.
type GetTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	TemplateID string `json:"-"`
}

// GetTemplateUserDataRequest: get template user data request.
type GetTemplateUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	TemplateID string `json:"-"`

	Key string `json:"-"`
}

// GetUserDataRequest: get user data request.
type GetUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	Key string `json:"-"`
}

// ListPlacementGroupsRequest: list placement groups request.
type ListPlacementGroupsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PageToken *string `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_desc
	OrderBy ListPlacementGroupsRequestOrderBy `json:"-"`

	ProjectID string `json:"-"`

	PlacementGroupIDs []string `json:"-"`

	Name *string `json:"-"`

	Tags []string `json:"-"`
}

// ListPlacementGroupsResponse: list placement groups response.
type ListPlacementGroupsResponse struct {
	PlacementGroups []*PlacementGroup `json:"placement_groups"`

	NextPageToken *string `json:"next_page_token"`

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

	PageToken *string `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_desc
	OrderBy ListPrivateNetworkInterfacesRequestOrderBy `json:"-"`

	ProjectID string `json:"-"`

	ServerIDs []string `json:"-"`

	SecurityGroupIDs []string `json:"-"`

	PrivateNetworkIDs []string `json:"-"`

	Tags []string `json:"-"`
}

// ListPrivateNetworkInterfacesResponse: list private network interfaces response.
type ListPrivateNetworkInterfacesResponse struct {
	PrivateNetworkInterfaces []*PrivateNetworkInterface `json:"private_network_interfaces"`

	NextPageToken *string `json:"next_page_token"`

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

	PageToken *string `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_desc
	OrderBy ListSecurityGroupsRequestOrderBy `json:"-"`

	ProjectID string `json:"-"`

	Name *string `json:"-"`

	Tags []string `json:"-"`

	SecurityGroupIDs []string `json:"-"`
}

// ListSecurityGroupsResponse: list security groups response.
type ListSecurityGroupsResponse struct {
	SecurityGroups []*SecurityGroupSummary `json:"security_groups"`

	NextPageToken *string `json:"next_page_token"`

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

	PageToken *string `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListServerTypesResponse: list server types response.
type ListServerTypesResponse struct {
	ServerTypes []*ServerType `json:"server_types"`

	NextPageToken *string `json:"next_page_token"`

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

	PageToken *string `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_desc
	OrderBy ListServersRequestOrderBy `json:"-"`

	ProjectID string `json:"-"`

	ServerIDs []string `json:"-"`

	Name *string `json:"-"`

	ServerType *string `json:"-"`

	Tags []string `json:"-"`

	SecurityGroupIDs []string `json:"-"`

	PlacementGroupIDs []string `json:"-"`

	PrivateNetworkIDs []string `json:"-"`

	MacAddresses []string `json:"-"`
}

// ListServersResponse: list servers response.
type ListServersResponse struct {
	Servers []*ServerSummary `json:"servers"`

	NextPageToken *string `json:"next_page_token"`

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

	TemplateID string `json:"-"`

	PageToken *string `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListTemplateUserDataKeysResponse: list template user data keys response.
type ListTemplateUserDataKeysResponse struct {
	Keys []string `json:"keys"`

	NextPageToken *string `json:"next_page_token"`

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

	PageToken *string `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_desc
	OrderBy ListTemplatesRequestOrderBy `json:"-"`

	ProjectID string `json:"-"`

	TemplateIDs []string `json:"-"`

	Name *string `json:"-"`

	Tags []string `json:"-"`

	ServerTags []string `json:"-"`

	SecurityGroupIDs []string `json:"-"`

	PlacementGroupIDs []string `json:"-"`
}

// ListTemplatesResponse: list templates response.
type ListTemplatesResponse struct {
	Templates []*TemplateSummary `json:"templates"`

	NextPageToken *string `json:"next_page_token"`

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

	ServerID string `json:"-"`

	PageToken *string `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListUserDataKeysResponse: list user data keys response.
type ListUserDataKeysResponse struct {
	Keys []string `json:"keys"`

	NextPageToken *string `json:"next_page_token"`

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

	ServerID string `json:"-"`
}

// RebootServerRequest: reboot server request.
type RebootServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

// ResourceCounts: resource counts.
type ResourceCounts struct {
	Servers uint32 `json:"servers"`

	GpuServers uint32 `json:"gpu_servers"`

	ServersByType map[string]uint32 `json:"servers_by_type"`

	SecurityGroups uint32 `json:"security_groups"`

	PlacementGroups uint32 `json:"placement_groups"`

	Snapshots uint32 `json:"snapshots"`

	Volumes uint32 `json:"volumes"`

	VolumesLSSDTotalSize uint64 `json:"volumes_l_ssd_total_size"`

	PrivateNetworkInterfaces uint32 `json:"private_network_interfaces"`

	VolumesScratch uint32 `json:"volumes_scratch"`

	VolumesLSSD uint32 `json:"volumes_l_ssd"`
}

// Server: server.
type Server struct {
	ProjectID string `json:"project_id"`

	ID string `json:"id"`

	Name string `json:"name"`

	Tags []string `json:"tags"`

	ServerType string `json:"server_type"`

	PlacementGroupID *string `json:"placement_group_id"`

	// Status: default value: unknown_status
	Status ServerStatus `json:"status"`

	Volumes []*ServerVolume `json:"volumes"`

	// Architecture: default value: unknown_architecture
	Architecture ServerArchitecture `json:"architecture"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	PrivateNetworkInterfaces []*ServerPrivateNetworkInterface `json:"private_network_interfaces"`

	RescueMode bool `json:"rescue_mode"`

	BootVolumeID *string `json:"boot_volume_id"`

	StatusDetail string `json:"status_detail"`

	WindowsRdpPassword *ServerRDPPassword `json:"windows_rdp_password"`

	Filesystems []*ServerFilesystem `json:"filesystems"`

	PublicNetworkInterface *ServerPublicNetworkInterface `json:"public_network_interface"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
}

// SetSecurityGroupRulesRequest: set security group rules request.
type SetSecurityGroupRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"security_group_id"`

	SecurityGroupRules []*SecurityGroupRuleConfig `json:"security_group_rules"`
}

// SetServerCloudInitRequest: set server cloud init request.
type SetServerCloudInitRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	Content []byte `json:"content"`
}

// SetServerDefaultIPRequest: set server default ip request.
type SetServerDefaultIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	IPID string `json:"ip_id"`
}

// SetTemplateCloudInitRequest: set template cloud init request.
type SetTemplateCloudInitRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	TemplateID string `json:"-"`

	Content []byte `json:"content"`
}

// SetTemplateUserDataRequest: set template user data request.
type SetTemplateUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	TemplateID string `json:"-"`

	Key string `json:"-"`

	Content []byte `json:"content"`
}

// SetUserDataRequest: set user data request.
type SetUserDataRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	Key string `json:"-"`

	Content []byte `json:"content"`
}

// StartServerRequest: start server request.
type StartServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

// StopAndDeleteServerRequest: stop and delete server request.
type StopAndDeleteServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	// Precisely one of DeleteAllIPs, DeleteIPIDs must be set.
	DeleteAllIPs *bool `json:"delete_all_ips,omitempty"`

	// Precisely one of DeleteAllIPs, DeleteIPIDs must be set.
	DeleteIPIDs *[]string `json:"delete_ip_ids,omitempty"`

	// Precisely one of DeleteAllVolumes, DeleteVolumeIDs must be set.
	DeleteAllVolumes *bool `json:"delete_all_volumes,omitempty"`

	// Precisely one of DeleteAllVolumes, DeleteVolumeIDs must be set.
	DeleteVolumeIDs *[]string `json:"delete_volume_ids,omitempty"`

	// Precisely one of KeepAllPrivateNics, DeletePrivateNicIDs must be set.
	KeepAllPrivateNics *bool `json:"keep_all_private_nics,omitempty"`

	// Precisely one of KeepAllPrivateNics, DeletePrivateNicIDs must be set.
	DeletePrivateNicIDs *[]string `json:"delete_private_nic_ids,omitempty"`
}

// StopServerRequest: stop server request.
type StopServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`
}

// Template: template.
type Template struct {
	ProjectID string `json:"project_id"`

	ID string `json:"id"`

	Name string `json:"name"`

	Tags []string `json:"tags"`

	ServerTags []string `json:"server_tags"`

	ServerType string `json:"server_type"`

	SecurityGroupID *string `json:"security_group_id"`

	PlacementGroupID *string `json:"placement_group_id"`

	PublicIPV4Count uint32 `json:"public_ip_v4_count"`

	PublicIPV6Count uint32 `json:"public_ip_v6_count"`

	Volumes []*CreateTemplateRequestVolumeTemplate `json:"volumes"`

	PrivateNetworks []*CreateTemplateRequestPrivateNetworkTemplate `json:"private_networks"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	WindowsRdpSSHKeyID *string `json:"windows_rdp_ssh_key_id"`

	FilesystemIDs []string `json:"filesystem_ids"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
}

// UpdatePlacementGroupRequest: update placement group request.
type UpdatePlacementGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PlacementGroupID string `json:"-"`

	Name *string `json:"name,omitempty"`

	// PolicyType: default value: unknown_policy_type
	PolicyType PlacementGroupPolicyType `json:"policy_type"`

	Tags *[]string `json:"tags,omitempty"`
}

// UpdatePrivateNetworkInterfaceRequest: update private network interface request.
type UpdatePrivateNetworkInterfaceRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	PrivateNetworkInterfaceID string `json:"-"`

	SecurityGroupID *string `json:"security_group_id,omitempty"`

	Tags *[]string `json:"tags,omitempty"`
}

// UpdateSecurityGroupRequest: update security group request.
type UpdateSecurityGroupRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SecurityGroupID string `json:"-"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	DisableDefaultRules *bool `json:"disable_default_rules,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	ProjectDefault *bool `json:"project_default,omitempty"`

	// InboundDefaultAction: default value: unknown_action
	InboundDefaultAction SecurityGroupAction `json:"inbound_default_action"`

	// OutboundDefaultAction: default value: unknown_action
	OutboundDefaultAction SecurityGroupAction `json:"outbound_default_action"`

	Stateless *bool `json:"stateless,omitempty"`
}

// UpdateSecurityGroupRuleRequest: update security group rule request.
type UpdateSecurityGroupRuleRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	SecurityGroupRuleID string `json:"-"`

	// Protocol: default value: unknown_protocol
	Protocol SecurityGroupRuleProtocol `json:"protocol"`

	// Direction: default value: unknown_direction
	Direction SecurityGroupRuleDirection `json:"direction"`

	// Action: default value: unknown_action
	Action SecurityGroupRuleAction `json:"action"`

	SourceIPRange *scw.IPNet `json:"source_ip_range,omitempty"`

	DestinationIPRange *scw.IPNet `json:"destination_ip_range,omitempty"`

	SourcePorts *SecurityGroupRulePortRange `json:"source_ports,omitempty"`

	DestinationPorts *SecurityGroupRulePortRange `json:"destination_ports,omitempty"`

	Position *int32 `json:"position,omitempty"`
}

// UpdateServerRequest: update server request.
type UpdateServerRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	ServerID string `json:"-"`

	Name *string `json:"name,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	ServerType *string `json:"server_type,omitempty"`

	PlacementGroupID *string `json:"placement_group_id,omitempty"`

	RescueMode *bool `json:"rescue_mode,omitempty"`

	BootVolumeID *string `json:"boot_volume_id,omitempty"`

	WindowsRdpSSHKeyID *string `json:"windows_rdp_ssh_key_id,omitempty"`

	Protected *bool `json:"protected,omitempty"`

	PublicNetworkInterface *UpdateServerRequestPublicNetworkInterface `json:"public_network_interface,omitempty"`
}

// UpdateTemplateRequest: update template request.
type UpdateTemplateRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	TemplateID string `json:"-"`

	Name *string `json:"name,omitempty"`

	Tags *[]string `json:"tags,omitempty"`

	ServerTags *[]string `json:"server_tags,omitempty"`

	ServerType *string `json:"server_type,omitempty"`

	SecurityGroupID *string `json:"security_group_id,omitempty"`

	PlacementGroupID *string `json:"placement_group_id,omitempty"`

	UpdateVolumes *UpdateTemplateRequestUpdateVolumes `json:"update_volumes,omitempty"`

	UpdatePrivateNetworks *UpdateTemplateRequestUpdatePrivateNetworks `json:"update_private_networks,omitempty"`

	PublicIPV4Count *uint32 `json:"public_ip_v4_count,omitempty"`

	PublicIPV6Count *uint32 `json:"public_ip_v6_count,omitempty"`

	WindowsRdpSSHKeyID *string `json:"windows_rdp_ssh_key_id,omitempty"`

	FilesystemIDs *[]string `json:"filesystem_ids,omitempty"`
}

// UserData: user data.
type UserData struct {
	Key string `json:"key"`

	Content []byte `json:"content"`
}

// This API allows you to manage your CPU and GPU Instances.
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
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneFrPar3, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3}
}

// GetResourceCounts: Get counts of various resources (e.g. servers, volumes).
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

// ListServers: List all Instances.
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

// CreateServer: Create a new Instance of a specified server_type.
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

// GetServer: Get the details of a specified Instance.
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

// UpdateServer: Update the properties of a specified Instance information, such as name, rescue_mode, or tags.
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

// DeleteServer: Delete a specified Instance.
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

// ListServerTypes: List available Instance types and their technical details.
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

// StartServer: Start a stopped or paused Instance.
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

// RebootServer: Reboot a running or paused Instance.
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

// PauseServer: Pause a running Instance.
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

// StopServer: Stop a running or paused Instance.
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

// StopAndDeleteServer: Stop and delete a running or paused Instance.
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

// AttachServerVolume: Attach a l_ssd or SBS volume to an Instance.
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

// DetachServerVolume: Detach a volume from an Instance.
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

// AttachServerFileSystem: Attach a filesystem volume to an Instance.
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

// DetachServerFileSystem: Detach a filesystem volume from an Instance.
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

// AttachServerIP: Attach an IP to an Instance.
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

// DetachServerIP: Detach an IP from an Instance.
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

// SetServerDefaultIP: Set the default IP for an Instance.
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

// AttachServerPrivateNetworkInterface: Attach a private network interface to an Instance.
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

// DetachServerPrivateNetworkInterface: Detach a private network interface from an Instance.
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

// ListPrivateNetworkInterfaces: List all private network interfaces.
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
	parameter.AddToQuery(query, "security_group_ids", req.SecurityGroupIDs)
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

// CreatePrivateNetworkInterface: Create a private network interface linked to a Private Network. It can be attached to an Instance.
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

// GetPrivateNetworkInterface: Get details of a specified private network interface.
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

// UpdatePrivateNetworkInterface: Update the properties of a specified private network interface.
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

// DeletePrivateNetworkInterface: Delete a specified private network interface.
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

// ListPlacementGroups: List all placement groups.
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

// CreatePlacementGroup: Create a new placement group.
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

// GetPlacementGroup: Get a specified placement group.
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

// UpdatePlacementGroup: Update the properties of a specified placement group.
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

// DeletePlacementGroup: Delete a specified placement group.
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

// ListSecurityGroups: List all security groups.
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

// CreateSecurityGroup: Create a security group with a specified name and description.
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

// GetSecurityGroup: Get the details of a specified security group.
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

// UpdateSecurityGroup: Update the properties of a security group.
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

// DeleteSecurityGroup: Delete a specified security group.
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

// AddSecurityGroupRules: Add one or more rules to a security group.
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

// SetSecurityGroupRules: Replace all rules of a specified security group with the provided rules.
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

// UpdateSecurityGroupRule: Update the properties of a rule from a specified security group.
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

// DeleteSecurityGroupRules: Delete specified security groups.
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

// ListUserDataKeys: List all user data keys registered on a specified Instance.
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

// GetUserData: Get the content of a user data with a specified key on an Instance.
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

// SetUserData: Add or update a user data with a specified key on an Instance.
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

// DeleteUserData: Delete a specified key from an Instance's user data.
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

// GetServerCloudInit: Get the cloud-init configuration of a specified Instance.
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

// SetServerCloudInit: Set the cloud-init configuration for a specified Instance.
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

// ListTemplates: List all available templates.
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

// CreateTemplate: Create a new template from an Instance.
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

// GetTemplate: Get details of a specified template.
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

// UpdateTemplate: Update the properties of a template.
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

// DeleteTemplate: Delete a specified template.
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

// ListTemplateUserDataKeys: List all user data keys of a template.
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

// GetTemplateUserData: Get a specific user data key of a template.
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

// SetTemplateUserData: Set a user data key of a template.
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

// DeleteTemplateUserData: Delete a specific user data key of a template.
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

// GetTemplateCloudInit: Get the cloud-init configuration of a template.
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

// SetTemplateCloudInit: Set the cloud-init configuration of a template.
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

// CheckTemplate: Validate that a template is usable.
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

// CreateServerFromTemplate: Create a new Instance using a specified template.
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

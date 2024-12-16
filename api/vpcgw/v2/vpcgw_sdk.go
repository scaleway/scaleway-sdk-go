// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package vpcgw provides methods and message types of the vpcgw v2 API.
package vpcgw

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

type GatewayNetworkStatus string

const (
	GatewayNetworkStatusUnknownStatus = GatewayNetworkStatus("unknown_status")
	GatewayNetworkStatusCreated       = GatewayNetworkStatus("created")
	GatewayNetworkStatusAttaching     = GatewayNetworkStatus("attaching")
	GatewayNetworkStatusConfiguring   = GatewayNetworkStatus("configuring")
	GatewayNetworkStatusReady         = GatewayNetworkStatus("ready")
	GatewayNetworkStatusDetaching     = GatewayNetworkStatus("detaching")
)

func (enum GatewayNetworkStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum GatewayNetworkStatus) Values() []GatewayNetworkStatus {
	return []GatewayNetworkStatus{
		"unknown_status",
		"created",
		"attaching",
		"configuring",
		"ready",
		"detaching",
	}
}

func (enum GatewayNetworkStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *GatewayNetworkStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = GatewayNetworkStatus(GatewayNetworkStatus(tmp).String())
	return nil
}

type GatewayStatus string

const (
	GatewayStatusUnknownStatus = GatewayStatus("unknown_status")
	GatewayStatusStopped       = GatewayStatus("stopped")
	GatewayStatusAllocating    = GatewayStatus("allocating")
	GatewayStatusConfiguring   = GatewayStatus("configuring")
	GatewayStatusRunning       = GatewayStatus("running")
	GatewayStatusStopping      = GatewayStatus("stopping")
	GatewayStatusFailed        = GatewayStatus("failed")
	GatewayStatusDeleting      = GatewayStatus("deleting")
	GatewayStatusLocked        = GatewayStatus("locked")
)

func (enum GatewayStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum GatewayStatus) Values() []GatewayStatus {
	return []GatewayStatus{
		"unknown_status",
		"stopped",
		"allocating",
		"configuring",
		"running",
		"stopping",
		"failed",
		"deleting",
		"locked",
	}
}

func (enum GatewayStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *GatewayStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = GatewayStatus(GatewayStatus(tmp).String())
	return nil
}

type ListGatewayNetworksRequestOrderBy string

const (
	ListGatewayNetworksRequestOrderByCreatedAtAsc  = ListGatewayNetworksRequestOrderBy("created_at_asc")
	ListGatewayNetworksRequestOrderByCreatedAtDesc = ListGatewayNetworksRequestOrderBy("created_at_desc")
	ListGatewayNetworksRequestOrderByStatusAsc     = ListGatewayNetworksRequestOrderBy("status_asc")
	ListGatewayNetworksRequestOrderByStatusDesc    = ListGatewayNetworksRequestOrderBy("status_desc")
)

func (enum ListGatewayNetworksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListGatewayNetworksRequestOrderBy) Values() []ListGatewayNetworksRequestOrderBy {
	return []ListGatewayNetworksRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"status_asc",
		"status_desc",
	}
}

func (enum ListGatewayNetworksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListGatewayNetworksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListGatewayNetworksRequestOrderBy(ListGatewayNetworksRequestOrderBy(tmp).String())
	return nil
}

type ListGatewaysRequestOrderBy string

const (
	ListGatewaysRequestOrderByCreatedAtAsc  = ListGatewaysRequestOrderBy("created_at_asc")
	ListGatewaysRequestOrderByCreatedAtDesc = ListGatewaysRequestOrderBy("created_at_desc")
	ListGatewaysRequestOrderByNameAsc       = ListGatewaysRequestOrderBy("name_asc")
	ListGatewaysRequestOrderByNameDesc      = ListGatewaysRequestOrderBy("name_desc")
	ListGatewaysRequestOrderByTypeAsc       = ListGatewaysRequestOrderBy("type_asc")
	ListGatewaysRequestOrderByTypeDesc      = ListGatewaysRequestOrderBy("type_desc")
	ListGatewaysRequestOrderByStatusAsc     = ListGatewaysRequestOrderBy("status_asc")
	ListGatewaysRequestOrderByStatusDesc    = ListGatewaysRequestOrderBy("status_desc")
)

func (enum ListGatewaysRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListGatewaysRequestOrderBy) Values() []ListGatewaysRequestOrderBy {
	return []ListGatewaysRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
		"type_asc",
		"type_desc",
		"status_asc",
		"status_desc",
	}
}

func (enum ListGatewaysRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListGatewaysRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListGatewaysRequestOrderBy(ListGatewaysRequestOrderBy(tmp).String())
	return nil
}

type ListIPsRequestOrderBy string

const (
	ListIPsRequestOrderByCreatedAtAsc  = ListIPsRequestOrderBy("created_at_asc")
	ListIPsRequestOrderByCreatedAtDesc = ListIPsRequestOrderBy("created_at_desc")
	ListIPsRequestOrderByAddressAsc    = ListIPsRequestOrderBy("address_asc")
	ListIPsRequestOrderByAddressDesc   = ListIPsRequestOrderBy("address_desc")
	ListIPsRequestOrderByReverseAsc    = ListIPsRequestOrderBy("reverse_asc")
	ListIPsRequestOrderByReverseDesc   = ListIPsRequestOrderBy("reverse_desc")
)

func (enum ListIPsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListIPsRequestOrderBy) Values() []ListIPsRequestOrderBy {
	return []ListIPsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"address_asc",
		"address_desc",
		"reverse_asc",
		"reverse_desc",
	}
}

func (enum ListIPsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListIPsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListIPsRequestOrderBy(ListIPsRequestOrderBy(tmp).String())
	return nil
}

type ListPatRulesRequestOrderBy string

const (
	ListPatRulesRequestOrderByCreatedAtAsc   = ListPatRulesRequestOrderBy("created_at_asc")
	ListPatRulesRequestOrderByCreatedAtDesc  = ListPatRulesRequestOrderBy("created_at_desc")
	ListPatRulesRequestOrderByPublicPortAsc  = ListPatRulesRequestOrderBy("public_port_asc")
	ListPatRulesRequestOrderByPublicPortDesc = ListPatRulesRequestOrderBy("public_port_desc")
)

func (enum ListPatRulesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListPatRulesRequestOrderBy) Values() []ListPatRulesRequestOrderBy {
	return []ListPatRulesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"public_port_asc",
		"public_port_desc",
	}
}

func (enum ListPatRulesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPatRulesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPatRulesRequestOrderBy(ListPatRulesRequestOrderBy(tmp).String())
	return nil
}

type PatRuleProtocol string

const (
	PatRuleProtocolUnknownProtocol = PatRuleProtocol("unknown_protocol")
	PatRuleProtocolBoth            = PatRuleProtocol("both")
	PatRuleProtocolTCP             = PatRuleProtocol("tcp")
	PatRuleProtocolUDP             = PatRuleProtocol("udp")
)

func (enum PatRuleProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_protocol"
	}
	return string(enum)
}

func (enum PatRuleProtocol) Values() []PatRuleProtocol {
	return []PatRuleProtocol{
		"unknown_protocol",
		"both",
		"tcp",
		"udp",
	}
}

func (enum PatRuleProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PatRuleProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PatRuleProtocol(PatRuleProtocol(tmp).String())
	return nil
}

// GatewayNetwork: gateway network.
type GatewayNetwork struct {
	// ID: ID of the Public Gateway-Private Network connection.
	ID string `json:"id"`

	// CreatedAt: connection creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: connection last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// GatewayID: ID of the connected Public Gateway.
	GatewayID string `json:"gateway_id"`

	// PrivateNetworkID: ID of the connected Private Network.
	PrivateNetworkID string `json:"private_network_id"`

	// MacAddress: mAC address of the gateway in the Private Network (if the gateway is up and running).
	MacAddress *string `json:"mac_address"`

	// MasqueradeEnabled: defines whether the gateway masquerades traffic for this Private Network (Dynamic NAT).
	MasqueradeEnabled bool `json:"masquerade_enabled"`

	// Status: current status of the Public Gateway's connection to the Private Network.
	// Default value: unknown_status
	Status GatewayNetworkStatus `json:"status"`

	// PushDefaultRoute: enabling the default route also enables masquerading.
	PushDefaultRoute bool `json:"push_default_route"`

	// IpamIPID: use this IPAM-booked IP ID as the Gateway's IP in this Private Network.
	IpamIPID string `json:"ipam_ip_id"`

	// Zone: zone of the GatewayNetwork connection.
	Zone scw.Zone `json:"zone"`
}

// IP: ip.
type IP struct {
	// ID: IP address ID.
	ID string `json:"id"`

	// OrganizationID: owning Organization.
	OrganizationID string `json:"organization_id"`

	// ProjectID: owning Project.
	ProjectID string `json:"project_id"`

	// CreatedAt: IP address creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: IP address last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Tags: tags associated with the IP address.
	Tags []string `json:"tags"`

	// Address: the IP address itself.
	Address net.IP `json:"address"`

	// Reverse: reverse domain name for the IP address.
	Reverse *string `json:"reverse"`

	// GatewayID: public Gateway associated with the IP address.
	GatewayID *string `json:"gateway_id"`

	// Zone: zone of the IP address.
	Zone scw.Zone `json:"zone"`
}

// GatewayType: gateway type.
type GatewayType struct {
	// Name: public Gateway type name.
	Name string `json:"name"`

	// Bandwidth: bandwidth, in bps, of the Public Gateway. This is the public bandwidth to the outer Internet, and the internal bandwidth to each connected Private Networks.
	Bandwidth uint64 `json:"bandwidth"`

	// Zone: zone the Public Gateway type is available in.
	Zone scw.Zone `json:"zone"`
}

// Gateway: gateway.
type Gateway struct {
	// ID: ID of the gateway.
	ID string `json:"id"`

	// OrganizationID: owning Organization.
	OrganizationID string `json:"organization_id"`

	// ProjectID: owning Project.
	ProjectID string `json:"project_id"`

	// CreatedAt: gateway creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: gateway last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Type: gateway type name (commercial offer).
	Type string `json:"type"`

	// Bandwidth: bandwidth available of the gateway.
	Bandwidth uint64 `json:"bandwidth"`

	// Status: current status of the gateway.
	// Default value: unknown_status
	Status GatewayStatus `json:"status"`

	// Name: name of the gateway.
	Name string `json:"name"`

	// Tags: tags associated with the gateway.
	Tags []string `json:"tags"`

	// IPv4: public IPv4 address of the gateway.
	IPv4 *IP `json:"ipv4"`

	// GatewayNetworks: gatewayNetwork objects attached to the gateway (each one represents a connection to a Private Network).
	GatewayNetworks []*GatewayNetwork `json:"gateway_networks"`

	// Version: version of the running gateway software.
	Version *string `json:"version"`

	// CanUpgradeTo: newly available gateway software version that can be updated to.
	CanUpgradeTo *string `json:"can_upgrade_to"`

	// BastionEnabled: defines whether SSH bastion is enabled on the gateway.
	BastionEnabled bool `json:"bastion_enabled"`

	// BastionPort: port of the SSH bastion.
	BastionPort uint32 `json:"bastion_port"`

	// SMTPEnabled: defines whether SMTP traffic is allowed to pass through the gateway.
	SMTPEnabled bool `json:"smtp_enabled"`

	// IsLegacy: defines whether the gateway uses non-IPAM IP configurations.
	IsLegacy bool `json:"is_legacy"`

	// Zone: zone of the gateway.
	Zone scw.Zone `json:"zone"`
}

// PatRule: pat rule.
type PatRule struct {
	// ID: pAT rule ID.
	ID string `json:"id"`

	// GatewayID: gateway the PAT rule applies to.
	GatewayID string `json:"gateway_id"`

	// CreatedAt: pAT rule creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: pAT rule last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// PublicPort: public port to listen on.
	PublicPort uint32 `json:"public_port"`

	// PrivateIP: private IP address to forward data to.
	PrivateIP net.IP `json:"private_ip"`

	// PrivatePort: private port to translate to.
	PrivatePort uint32 `json:"private_port"`

	// Protocol: protocol the rule applies to.
	// Default value: unknown_protocol
	Protocol PatRuleProtocol `json:"protocol"`

	// Zone: zone of the PAT rule.
	Zone scw.Zone `json:"zone"`
}

// SetPatRulesRequestRule: set pat rules request rule.
type SetPatRulesRequestRule struct {
	// PublicPort: public port to listen on. Uniquely identifies the rule, and a matching rule will be updated with the new parameters.
	PublicPort uint32 `json:"public_port"`

	// PrivateIP: private IP to forward data to.
	PrivateIP net.IP `json:"private_ip"`

	// PrivatePort: private port to translate to.
	PrivatePort uint32 `json:"private_port"`

	// Protocol: protocol the rule should apply to.
	// Default value: unknown_protocol
	Protocol PatRuleProtocol `json:"protocol"`
}

// CreateGatewayNetworkRequest: create gateway network request.
type CreateGatewayNetworkRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GatewayID: public Gateway to connect.
	GatewayID string `json:"gateway_id"`

	// PrivateNetworkID: private Network to connect.
	PrivateNetworkID string `json:"private_network_id"`

	// EnableMasquerade: defines whether to enable masquerade (dynamic NAT) on the GatewayNetwork.
	EnableMasquerade bool `json:"enable_masquerade"`

	// PushDefaultRoute: enabling the default route also enables masquerading.
	PushDefaultRoute bool `json:"push_default_route"`

	// IpamIPID: use this IPAM-booked IP ID as the Gateway's IP in this Private Network.
	IpamIPID *string `json:"ipam_ip_id,omitempty"`
}

// CreateGatewayRequest: create gateway request.
type CreateGatewayRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ProjectID: scaleway Project to create the gateway in.
	ProjectID string `json:"project_id"`

	// Name: name for the gateway.
	Name string `json:"name"`

	// Tags: tags for the gateway.
	Tags []string `json:"tags"`

	// Type: gateway type (commercial offer type).
	Type string `json:"type"`

	// IPID: existing IP address to attach to the gateway.
	IPID *string `json:"ip_id,omitempty"`

	// EnableSMTP: defines whether SMTP traffic should be allowed pass through the gateway.
	EnableSMTP bool `json:"enable_smtp"`

	// EnableBastion: defines whether SSH bastion should be enabled the gateway.
	EnableBastion bool `json:"enable_bastion"`

	// BastionPort: port of the SSH bastion.
	BastionPort *uint32 `json:"bastion_port,omitempty"`
}

// CreateIPRequest: create ip request.
type CreateIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ProjectID: project to create the IP address in.
	ProjectID string `json:"project_id"`

	// Tags: tags to give to the IP address.
	Tags []string `json:"tags"`
}

// CreatePatRuleRequest: create pat rule request.
type CreatePatRuleRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GatewayID: ID of the Gateway on which to create the rule.
	GatewayID string `json:"gateway_id"`

	// PublicPort: public port to listen on.
	PublicPort uint32 `json:"public_port"`

	// PrivateIP: private IP to forward data to.
	PrivateIP net.IP `json:"private_ip"`

	// PrivatePort: private port to translate to.
	PrivatePort uint32 `json:"private_port"`

	// Protocol: protocol the rule should apply to.
	// Default value: unknown_protocol
	Protocol PatRuleProtocol `json:"protocol"`
}

// DeleteGatewayNetworkRequest: delete gateway network request.
type DeleteGatewayNetworkRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GatewayNetworkID: ID of the GatewayNetwork to delete.
	GatewayNetworkID string `json:"-"`
}

// DeleteGatewayRequest: delete gateway request.
type DeleteGatewayRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GatewayID: ID of the gateway to delete.
	GatewayID string `json:"-"`
}

// DeleteIPRequest: delete ip request.
type DeleteIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// IPID: ID of the IP address to delete.
	IPID string `json:"-"`
}

// DeletePatRuleRequest: delete pat rule request.
type DeletePatRuleRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PatRuleID: ID of the PAT rule to delete.
	PatRuleID string `json:"-"`
}

// GetGatewayNetworkRequest: get gateway network request.
type GetGatewayNetworkRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GatewayNetworkID: ID of the GatewayNetwork to fetch.
	GatewayNetworkID string `json:"-"`
}

// GetGatewayRequest: get gateway request.
type GetGatewayRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GatewayID: ID of the gateway to fetch.
	GatewayID string `json:"-"`
}

// GetIPRequest: get ip request.
type GetIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// IPID: ID of the IP address to get.
	IPID string `json:"-"`
}

// GetPatRuleRequest: get pat rule request.
type GetPatRuleRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PatRuleID: ID of the PAT rule to get.
	PatRuleID string `json:"-"`
}

// ListGatewayNetworksRequest: list gateway networks request.
type ListGatewayNetworksRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListGatewayNetworksRequestOrderBy `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: gatewayNetworks per page.
	PageSize *uint32 `json:"-"`

	// Status: filter for GatewayNetworks with these status. Use `unknown` to include all statuses.
	Status []GatewayNetworkStatus `json:"-"`

	// GatewayIDs: filter for GatewayNetworks connected to these gateways.
	GatewayIDs []string `json:"-"`

	// PrivateNetworkIDs: filter for GatewayNetworks connected to these Private Networks.
	PrivateNetworkIDs []string `json:"-"`

	// MasqueradeEnabled: filter for GatewayNetworks with this `enable_masquerade` setting.
	MasqueradeEnabled *bool `json:"-"`
}

// ListGatewayNetworksResponse: list gateway networks response.
type ListGatewayNetworksResponse struct {
	// GatewayNetworks: gatewayNetworks on this page.
	GatewayNetworks []*GatewayNetwork `json:"gateway_networks"`

	// TotalCount: total GatewayNetworks count matching the filter.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListGatewayNetworksResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListGatewayNetworksResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListGatewayNetworksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.GatewayNetworks = append(r.GatewayNetworks, results.GatewayNetworks...)
	r.TotalCount += uint64(len(results.GatewayNetworks))
	return uint64(len(results.GatewayNetworks)), nil
}

// ListGatewayTypesRequest: list gateway types request.
type ListGatewayTypesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// ListGatewayTypesResponse: list gateway types response.
type ListGatewayTypesResponse struct {
	// Types: available types of Public Gateway.
	Types []*GatewayType `json:"types"`
}

// ListGatewaysRequest: list gateways request.
type ListGatewaysRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListGatewaysRequestOrderBy `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: gateways per page.
	PageSize *uint32 `json:"-"`

	// OrganizationID: include only gateways in this Organization.
	OrganizationID *string `json:"-"`

	// ProjectID: include only gateways in this Project.
	ProjectID *string `json:"-"`

	// Name: filter for gateways which have this search term in their name.
	Name *string `json:"-"`

	// Tags: filter for gateways with these tags.
	Tags []string `json:"-"`

	// Types: filter for gateways of these types.
	Types []string `json:"-"`

	// Status: filter for gateways with these status. Use `unknown` to include all statuses.
	Status []GatewayStatus `json:"-"`

	// PrivateNetworkIDs: filter for gateways attached to these Private Networks.
	PrivateNetworkIDs []string `json:"-"`

	// IncludeLegacy: include also legacy gateways.
	IncludeLegacy *bool `json:"-"`
}

// ListGatewaysResponse: list gateways response.
type ListGatewaysResponse struct {
	// Gateways: gateways on this page.
	Gateways []*Gateway `json:"gateways"`

	// TotalCount: total count of gateways matching the filter.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListGatewaysResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListGatewaysResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListGatewaysResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Gateways = append(r.Gateways, results.Gateways...)
	r.TotalCount += uint64(len(results.Gateways))
	return uint64(len(results.Gateways)), nil
}

// ListIPsRequest: list i ps request.
type ListIPsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListIPsRequestOrderBy `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: IP addresses per page.
	PageSize *uint32 `json:"-"`

	// OrganizationID: include only gateways in this Organization.
	OrganizationID *string `json:"-"`

	// ProjectID: filter for IP addresses in this Project.
	ProjectID *string `json:"-"`

	// Tags: filter for IP addresses with these tags.
	Tags []string `json:"-"`

	// Reverse: filter for IP addresses that have a reverse containing this string.
	Reverse *string `json:"-"`

	// IsFree: filter based on whether the IP is attached to a gateway or not.
	IsFree *bool `json:"-"`
}

// ListIPsResponse: list i ps response.
type ListIPsResponse struct {
	// IPs: IP addresses on this page.
	IPs []*IP `json:"ips"`

	// TotalCount: total count of IP addresses matching the filter.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIPsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListIPsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.IPs = append(r.IPs, results.IPs...)
	r.TotalCount += uint64(len(results.IPs))
	return uint64(len(results.IPs)), nil
}

// ListPatRulesRequest: list pat rules request.
type ListPatRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListPatRulesRequestOrderBy `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: pAT rules per page.
	PageSize *uint32 `json:"-"`

	// GatewayIDs: filter for PAT rules on these gateways.
	GatewayIDs []string `json:"-"`

	// PrivateIPs: filter for PAT rules targeting these private ips.
	PrivateIPs []string `json:"-"`

	// Protocol: filter for PAT rules with this protocol.
	// Default value: unknown_protocol
	Protocol PatRuleProtocol `json:"-"`
}

// ListPatRulesResponse: list pat rules response.
type ListPatRulesResponse struct {
	// PatRules: array of PAT rules matching the filter.
	PatRules []*PatRule `json:"pat_rules"`

	// TotalCount: total count of PAT rules matching the filter.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPatRulesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPatRulesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPatRulesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PatRules = append(r.PatRules, results.PatRules...)
	r.TotalCount += uint64(len(results.PatRules))
	return uint64(len(results.PatRules)), nil
}

// RefreshSSHKeysRequest: refresh ssh keys request.
type RefreshSSHKeysRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GatewayID: ID of the gateway to refresh SSH keys on.
	GatewayID string `json:"-"`
}

// SetPatRulesRequest: set pat rules request.
type SetPatRulesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GatewayID: ID of the gateway on which to set the PAT rules.
	GatewayID string `json:"gateway_id"`

	// PatRules: new list of PAT rules.
	PatRules []*SetPatRulesRequestRule `json:"pat_rules"`
}

// SetPatRulesResponse: set pat rules response.
type SetPatRulesResponse struct {
	// PatRules: list of PAT rules.
	PatRules []*PatRule `json:"pat_rules"`
}

// UpdateGatewayNetworkRequest: update gateway network request.
type UpdateGatewayNetworkRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GatewayNetworkID: ID of the GatewayNetwork to update.
	GatewayNetworkID string `json:"-"`

	// EnableMasquerade: defines whether to enable masquerade (dynamic NAT) on the GatewayNetwork.
	EnableMasquerade *bool `json:"enable_masquerade,omitempty"`

	// PushDefaultRoute: enabling the default route also enables masquerading.
	PushDefaultRoute *bool `json:"push_default_route,omitempty"`

	// IpamIPID: use this IPAM-booked IP ID as the Gateway's IP in this Private Network.
	IpamIPID *string `json:"ipam_ip_id,omitempty"`
}

// UpdateGatewayRequest: update gateway request.
type UpdateGatewayRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GatewayID: ID of the gateway to update.
	GatewayID string `json:"-"`

	// Name: name for the gateway.
	Name *string `json:"name,omitempty"`

	// Tags: tags for the gateway.
	Tags *[]string `json:"tags,omitempty"`

	// EnableBastion: defines whether SSH bastion should be enabled the gateway.
	EnableBastion *bool `json:"enable_bastion,omitempty"`

	// BastionPort: port of the SSH bastion.
	BastionPort *uint32 `json:"bastion_port,omitempty"`

	// EnableSMTP: defines whether SMTP traffic should be allowed to pass through the gateway.
	EnableSMTP *bool `json:"enable_smtp,omitempty"`
}

// UpdateIPRequest: update ip request.
type UpdateIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// IPID: ID of the IP address to update.
	IPID string `json:"-"`

	// Tags: tags to give to the IP address.
	Tags *[]string `json:"tags,omitempty"`

	// Reverse: reverse to set on the address. Empty string to unset.
	Reverse *string `json:"reverse,omitempty"`

	// GatewayID: gateway to attach the IP address to. Empty string to detach.
	GatewayID *string `json:"gateway_id,omitempty"`
}

// UpdatePatRuleRequest: update pat rule request.
type UpdatePatRuleRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PatRuleID: ID of the PAT rule to update.
	PatRuleID string `json:"-"`

	// PublicPort: public port to listen on.
	PublicPort *uint32 `json:"public_port,omitempty"`

	// PrivateIP: private IP to forward data to.
	PrivateIP *net.IP `json:"private_ip,omitempty"`

	// PrivatePort: private port to translate to.
	PrivatePort *uint32 `json:"private_port,omitempty"`

	// Protocol: protocol the rule should apply to.
	// Default value: unknown_protocol
	Protocol PatRuleProtocol `json:"protocol"`
}

// UpgradeGatewayRequest: upgrade gateway request.
type UpgradeGatewayRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// GatewayID: ID of the gateway to upgrade.
	GatewayID string `json:"-"`

	// Type: gateway type (commercial offer).
	Type *string `json:"type,omitempty"`
}

// This API allows you to manage your Public Gateways.
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
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZoneNlAms3, scw.ZonePlWaw1, scw.ZonePlWaw2, scw.ZonePlWaw3}
}

// ListGateways: List Public Gateways in a given Scaleway Organization or Project. By default, results are displayed in ascending order of creation date.
func (s *API) ListGateways(req *ListGatewaysRequest, opts ...scw.RequestOption) (*ListGatewaysResponse, error) {
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "types", req.Types)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "private_network_ids", req.PrivateNetworkIDs)
	parameter.AddToQuery(query, "include_legacy", req.IncludeLegacy)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateways",
		Query:  query,
	}

	var resp ListGatewaysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGateway: Get details of a Public Gateway, specified by its gateway ID. The response object contains full details of the gateway, including its **name**, **type**, **status** and more.
func (s *API) GetGateway(req *GetGatewayRequest, opts ...scw.RequestOption) (*Gateway, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayID) == "" {
		return nil, errors.New("field GatewayID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateways/" + fmt.Sprint(req.GatewayID) + "",
	}

	var resp Gateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateGateway: Create a new Public Gateway in the specified Scaleway Project, defining its **name**, **type** and other configuration details such as whether to enable SSH bastion.
func (s *API) CreateGateway(req *CreateGatewayRequest, opts ...scw.RequestOption) (*Gateway, error) {
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
		req.Name = namegenerator.GetRandomName("gw")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateways",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Gateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateGateway: Update the parameters of an existing Public Gateway, for example, its **name**, **tags**, **SSH bastion configuration**, and **DNS servers**.
func (s *API) UpdateGateway(req *UpdateGatewayRequest, opts ...scw.RequestOption) (*Gateway, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayID) == "" {
		return nil, errors.New("field GatewayID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateways/" + fmt.Sprint(req.GatewayID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Gateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteGateway: Delete an existing Public Gateway, specified by its gateway ID. This action is irreversible.
func (s *API) DeleteGateway(req *DeleteGatewayRequest, opts ...scw.RequestOption) (*Gateway, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayID) == "" {
		return nil, errors.New("field GatewayID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateways/" + fmt.Sprint(req.GatewayID) + "",
	}

	var resp Gateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpgradeGateway: Upgrade a given Public Gateway to the newest software version or to a different commercial offer type. This applies the latest bugfixes and features to your Public Gateway. Note that gateway service will be interrupted during the update.
func (s *API) UpgradeGateway(req *UpgradeGatewayRequest, opts ...scw.RequestOption) (*Gateway, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayID) == "" {
		return nil, errors.New("field GatewayID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateways/" + fmt.Sprint(req.GatewayID) + "/upgrade",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Gateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListGatewayNetworks: List the connections between Public Gateways and Private Networks (a connection = a GatewayNetwork). You can choose to filter by `gateway-id` to list all Private Networks attached to the specified Public Gateway, or by `private_network_id` to list all Public Gateways attached to the specified Private Network. Other query parameters are also available. The result is an array of GatewayNetwork objects, each giving details of the connection between a given Public Gateway and a given Private Network.
func (s *API) ListGatewayNetworks(req *ListGatewayNetworksRequest, opts ...scw.RequestOption) (*ListGatewayNetworksResponse, error) {
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
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "gateway_ids", req.GatewayIDs)
	parameter.AddToQuery(query, "private_network_ids", req.PrivateNetworkIDs)
	parameter.AddToQuery(query, "masquerade_enabled", req.MasqueradeEnabled)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateway-networks",
		Query:  query,
	}

	var resp ListGatewayNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGatewayNetwork: Get details of a given connection between a Public Gateway and a Private Network (this connection = a GatewayNetwork), specified by its `gateway_network_id`. The response object contains details of the connection including the IDs of the Public Gateway and Private Network, the dates the connection was created/updated and its configuration settings.
func (s *API) GetGatewayNetwork(req *GetGatewayNetworkRequest, opts ...scw.RequestOption) (*GatewayNetwork, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayNetworkID) == "" {
		return nil, errors.New("field GatewayNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateway-networks/" + fmt.Sprint(req.GatewayNetworkID) + "",
	}

	var resp GatewayNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateGatewayNetwork: Attach a specific Public Gateway to a specific Private Network (create a GatewayNetwork). You can configure parameters for the connection including whether to enable masquerade (dynamic NAT), and more.
func (s *API) CreateGatewayNetwork(req *CreateGatewayNetworkRequest, opts ...scw.RequestOption) (*GatewayNetwork, error) {
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
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateway-networks",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp GatewayNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateGatewayNetwork: Update the configuration parameters of a connection between a given Public Gateway and Private Network (the connection = a GatewayNetwork). Updatable parameters include whether to enable traffic masquerade (dynamic NAT).
func (s *API) UpdateGatewayNetwork(req *UpdateGatewayNetworkRequest, opts ...scw.RequestOption) (*GatewayNetwork, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayNetworkID) == "" {
		return nil, errors.New("field GatewayNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateway-networks/" + fmt.Sprint(req.GatewayNetworkID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp GatewayNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteGatewayNetwork: Detach a given Public Gateway from a given Private Network, i.e. delete a GatewayNetwork specified by a gateway_network_id.
func (s *API) DeleteGatewayNetwork(req *DeleteGatewayNetworkRequest, opts ...scw.RequestOption) (*GatewayNetwork, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayNetworkID) == "" {
		return nil, errors.New("field GatewayNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateway-networks/" + fmt.Sprint(req.GatewayNetworkID) + "",
	}

	var resp GatewayNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPatRules: List PAT rules. You can filter by gateway ID to list all PAT rules for a particular gateway, or filter for PAT rules targeting a specific IP address or using a specific protocol.
func (s *API) ListPatRules(req *ListPatRulesRequest, opts ...scw.RequestOption) (*ListPatRulesResponse, error) {
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
	parameter.AddToQuery(query, "gateway_ids", req.GatewayIDs)
	parameter.AddToQuery(query, "private_ips", req.PrivateIPs)
	parameter.AddToQuery(query, "protocol", req.Protocol)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/pat-rules",
		Query:  query,
	}

	var resp ListPatRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPatRule: Get a PAT rule, specified by its PAT rule ID. The response object gives full details of the PAT rule, including the Public Gateway it belongs to and the configuration settings in terms of public / private ports, private IP and protocol.
func (s *API) GetPatRule(req *GetPatRuleRequest, opts ...scw.RequestOption) (*PatRule, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PatRuleID) == "" {
		return nil, errors.New("field PatRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/pat-rules/" + fmt.Sprint(req.PatRuleID) + "",
	}

	var resp PatRule

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePatRule: Create a new PAT rule on a specified Public Gateway, defining the protocol to use, public port to listen on, and private port / IP address to map to.
func (s *API) CreatePatRule(req *CreatePatRuleRequest, opts ...scw.RequestOption) (*PatRule, error) {
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
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/pat-rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PatRule

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePatRule: Update a PAT rule, specified by its PAT rule ID. Configuration settings including private/public port, private IP address and protocol can all be updated.
func (s *API) UpdatePatRule(req *UpdatePatRuleRequest, opts ...scw.RequestOption) (*PatRule, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PatRuleID) == "" {
		return nil, errors.New("field PatRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/pat-rules/" + fmt.Sprint(req.PatRuleID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PatRule

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetPatRules: Set a definitive list of PAT rules attached to a Public Gateway. Each rule is identified by its public port and protocol. This will sync the current PAT rule list on the gateway with the new list, creating, updating or deleting PAT rules accordingly.
func (s *API) SetPatRules(req *SetPatRulesRequest, opts ...scw.RequestOption) (*SetPatRulesResponse, error) {
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
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/pat-rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetPatRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePatRule: Delete a PAT rule, identified by its PAT rule ID. This action is irreversible.
func (s *API) DeletePatRule(req *DeletePatRuleRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PatRuleID) == "" {
		return errors.New("field PatRuleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/pat-rules/" + fmt.Sprint(req.PatRuleID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListGatewayTypes: List the different Public Gateway commercial offer types available at Scaleway. The response is an array of objects describing the name and technical details of each available gateway type.
func (s *API) ListGatewayTypes(req *ListGatewayTypesRequest, opts ...scw.RequestOption) (*ListGatewayTypesResponse, error) {
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
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateway-types",
	}

	var resp ListGatewayTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListIPs: List Public Gateway flexible IP addresses. A number of filter options are available for limiting results in the response.
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

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "reverse", req.Reverse)
	parameter.AddToQuery(query, "is_free", req.IsFree)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/ips",
		Query:  query,
	}

	var resp ListIPsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetIP: Get details of a Public Gateway flexible IP address, identified by its IP ID. The response object contains information including which (if any) Public Gateway using this IP address, the reverse and various other metadata.
func (s *API) GetIP(req *GetIPRequest, opts ...scw.RequestOption) (*IP, error) {
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
		Method: "GET",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateIP: Create (reserve) a new flexible IP address that can be used for a Public Gateway in a specified Scaleway Project.
func (s *API) CreateIP(req *CreateIPRequest, opts ...scw.RequestOption) (*IP, error) {
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
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/ips",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateIP: Update details of an existing flexible IP address, including its tags, reverse and the Public Gateway it is assigned to.
func (s *API) UpdateIP(req *UpdateIPRequest, opts ...scw.RequestOption) (*IP, error) {
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
		Method: "PATCH",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteIP: Delete a flexible IP address from your account. This action is irreversible.
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
		Method: "DELETE",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/ips/" + fmt.Sprint(req.IPID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// RefreshSSHKeys: Refresh the SSH keys of a given Public Gateway, specified by its gateway ID. This adds any new SSH keys in the gateway's Scaleway Project to the gateway itself.
func (s *API) RefreshSSHKeys(req *RefreshSSHKeysRequest, opts ...scw.RequestOption) (*Gateway, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayID) == "" {
		return nil, errors.New("field GatewayID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc-gw/v2/zones/" + fmt.Sprint(req.Zone) + "/gateways/" + fmt.Sprint(req.GatewayID) + "/refresh-ssh-keys",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Gateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

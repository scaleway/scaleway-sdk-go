// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package vpc provides methods and message types of the vpc v2 API.
package vpc

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

type ACLRuleProtocol string

const (
	ACLRuleProtocolANY  = ACLRuleProtocol("ANY")
	ACLRuleProtocolTCP  = ACLRuleProtocol("TCP")
	ACLRuleProtocolUDP  = ACLRuleProtocol("UDP")
	ACLRuleProtocolICMP = ACLRuleProtocol("ICMP")
)

func (enum ACLRuleProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return string(ACLRuleProtocolANY)
	}
	return string(enum)
}

func (enum ACLRuleProtocol) Values() []ACLRuleProtocol {
	return []ACLRuleProtocol{
		"ANY",
		"TCP",
		"UDP",
		"ICMP",
	}
}

func (enum ACLRuleProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ACLRuleProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ACLRuleProtocol(ACLRuleProtocol(tmp).String())
	return nil
}

type Action string

const (
	ActionUnknownAction = Action("unknown_action")
	ActionAccept        = Action("accept")
	ActionDrop          = Action("drop")
)

func (enum Action) String() string {
	if enum == "" {
		// return default value if empty
		return string(ActionUnknownAction)
	}
	return string(enum)
}

func (enum Action) Values() []Action {
	return []Action{
		"unknown_action",
		"accept",
		"drop",
	}
}

func (enum Action) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Action) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Action(Action(tmp).String())
	return nil
}

type ListPrivateNetworksRequestOrderBy string

const (
	ListPrivateNetworksRequestOrderByCreatedAtAsc  = ListPrivateNetworksRequestOrderBy("created_at_asc")
	ListPrivateNetworksRequestOrderByCreatedAtDesc = ListPrivateNetworksRequestOrderBy("created_at_desc")
	ListPrivateNetworksRequestOrderByNameAsc       = ListPrivateNetworksRequestOrderBy("name_asc")
	ListPrivateNetworksRequestOrderByNameDesc      = ListPrivateNetworksRequestOrderBy("name_desc")
)

func (enum ListPrivateNetworksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPrivateNetworksRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListPrivateNetworksRequestOrderBy) Values() []ListPrivateNetworksRequestOrderBy {
	return []ListPrivateNetworksRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListPrivateNetworksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPrivateNetworksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPrivateNetworksRequestOrderBy(ListPrivateNetworksRequestOrderBy(tmp).String())
	return nil
}

type ListRoutesWithNexthopRequestOrderBy string

const (
	ListRoutesWithNexthopRequestOrderByCreatedAtAsc    = ListRoutesWithNexthopRequestOrderBy("created_at_asc")
	ListRoutesWithNexthopRequestOrderByCreatedAtDesc   = ListRoutesWithNexthopRequestOrderBy("created_at_desc")
	ListRoutesWithNexthopRequestOrderByDestinationAsc  = ListRoutesWithNexthopRequestOrderBy("destination_asc")
	ListRoutesWithNexthopRequestOrderByDestinationDesc = ListRoutesWithNexthopRequestOrderBy("destination_desc")
	ListRoutesWithNexthopRequestOrderByPrefixLenAsc    = ListRoutesWithNexthopRequestOrderBy("prefix_len_asc")
	ListRoutesWithNexthopRequestOrderByPrefixLenDesc   = ListRoutesWithNexthopRequestOrderBy("prefix_len_desc")
)

func (enum ListRoutesWithNexthopRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRoutesWithNexthopRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRoutesWithNexthopRequestOrderBy) Values() []ListRoutesWithNexthopRequestOrderBy {
	return []ListRoutesWithNexthopRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"destination_asc",
		"destination_desc",
		"prefix_len_asc",
		"prefix_len_desc",
	}
}

func (enum ListRoutesWithNexthopRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRoutesWithNexthopRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRoutesWithNexthopRequestOrderBy(ListRoutesWithNexthopRequestOrderBy(tmp).String())
	return nil
}

type ListSubnetsRequestOrderBy string

const (
	ListSubnetsRequestOrderByCreatedAtAsc  = ListSubnetsRequestOrderBy("created_at_asc")
	ListSubnetsRequestOrderByCreatedAtDesc = ListSubnetsRequestOrderBy("created_at_desc")
)

func (enum ListSubnetsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListSubnetsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListSubnetsRequestOrderBy) Values() []ListSubnetsRequestOrderBy {
	return []ListSubnetsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListSubnetsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSubnetsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSubnetsRequestOrderBy(ListSubnetsRequestOrderBy(tmp).String())
	return nil
}

type ListVPCConnectorsRequestOrderBy string

const (
	ListVPCConnectorsRequestOrderByCreatedAtAsc  = ListVPCConnectorsRequestOrderBy("created_at_asc")
	ListVPCConnectorsRequestOrderByCreatedAtDesc = ListVPCConnectorsRequestOrderBy("created_at_desc")
	ListVPCConnectorsRequestOrderByNameAsc       = ListVPCConnectorsRequestOrderBy("name_asc")
	ListVPCConnectorsRequestOrderByNameDesc      = ListVPCConnectorsRequestOrderBy("name_desc")
)

func (enum ListVPCConnectorsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListVPCConnectorsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListVPCConnectorsRequestOrderBy) Values() []ListVPCConnectorsRequestOrderBy {
	return []ListVPCConnectorsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListVPCConnectorsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListVPCConnectorsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListVPCConnectorsRequestOrderBy(ListVPCConnectorsRequestOrderBy(tmp).String())
	return nil
}

type ListVPCsRequestOrderBy string

const (
	ListVPCsRequestOrderByCreatedAtAsc  = ListVPCsRequestOrderBy("created_at_asc")
	ListVPCsRequestOrderByCreatedAtDesc = ListVPCsRequestOrderBy("created_at_desc")
	ListVPCsRequestOrderByNameAsc       = ListVPCsRequestOrderBy("name_asc")
	ListVPCsRequestOrderByNameDesc      = ListVPCsRequestOrderBy("name_desc")
)

func (enum ListVPCsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListVPCsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListVPCsRequestOrderBy) Values() []ListVPCsRequestOrderBy {
	return []ListVPCsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListVPCsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListVPCsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListVPCsRequestOrderBy(ListVPCsRequestOrderBy(tmp).String())
	return nil
}

type RouteType string

const (
	RouteTypeUnknownRouteType = RouteType("unknown_route_type")
	RouteTypeSubnet           = RouteType("subnet")
	RouteTypeDefault          = RouteType("default")
	RouteTypeCustom           = RouteType("custom")
	RouteTypeInterlink        = RouteType("interlink")
	RouteTypeS2sVpn           = RouteType("s2s_vpn")
)

func (enum RouteType) String() string {
	if enum == "" {
		// return default value if empty
		return string(RouteTypeUnknownRouteType)
	}
	return string(enum)
}

func (enum RouteType) Values() []RouteType {
	return []RouteType{
		"unknown_route_type",
		"subnet",
		"default",
		"custom",
		"interlink",
		"s2s_vpn",
	}
}

func (enum RouteType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RouteType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RouteType(RouteType(tmp).String())
	return nil
}

type RouteWithNexthopResourceType string

const (
	RouteWithNexthopResourceTypeUnknownType            = RouteWithNexthopResourceType("unknown_type")
	RouteWithNexthopResourceTypeVpcGatewayNetwork      = RouteWithNexthopResourceType("vpc_gateway_network")
	RouteWithNexthopResourceTypeInstancePrivateNic     = RouteWithNexthopResourceType("instance_private_nic")
	RouteWithNexthopResourceTypeBaremetalPrivateNic    = RouteWithNexthopResourceType("baremetal_private_nic")
	RouteWithNexthopResourceTypeAppleSiliconPrivateNic = RouteWithNexthopResourceType("apple_silicon_private_nic")
)

func (enum RouteWithNexthopResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return string(RouteWithNexthopResourceTypeUnknownType)
	}
	return string(enum)
}

func (enum RouteWithNexthopResourceType) Values() []RouteWithNexthopResourceType {
	return []RouteWithNexthopResourceType{
		"unknown_type",
		"vpc_gateway_network",
		"instance_private_nic",
		"baremetal_private_nic",
		"apple_silicon_private_nic",
	}
}

func (enum RouteWithNexthopResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RouteWithNexthopResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RouteWithNexthopResourceType(RouteWithNexthopResourceType(tmp).String())
	return nil
}

type VPCConnectorStatus string

const (
	VPCConnectorStatusUnknownVpcConnectorStatus = VPCConnectorStatus("unknown_vpc_connector_status")
	VPCConnectorStatusOrphan                    = VPCConnectorStatus("orphan")
	VPCConnectorStatusPeered                    = VPCConnectorStatus("peered")
	VPCConnectorStatusConflict                  = VPCConnectorStatus("conflict")
)

func (enum VPCConnectorStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(VPCConnectorStatusUnknownVpcConnectorStatus)
	}
	return string(enum)
}

func (enum VPCConnectorStatus) Values() []VPCConnectorStatus {
	return []VPCConnectorStatus{
		"unknown_vpc_connector_status",
		"orphan",
		"peered",
		"conflict",
	}
}

func (enum VPCConnectorStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VPCConnectorStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VPCConnectorStatus(VPCConnectorStatus(tmp).String())
	return nil
}

// Subnet: subnet.
type Subnet struct {
	// ID: ID of the subnet.
	ID string `json:"id"`

	// CreatedAt: subnet creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: subnet last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Subnet: subnet CIDR.
	Subnet scw.IPNet `json:"subnet"`

	// ProjectID: scaleway Project the subnet belongs to.
	ProjectID string `json:"project_id"`

	// PrivateNetworkID: private Network the subnet belongs to.
	PrivateNetworkID string `json:"private_network_id"`

	// VpcID: vPC the subnet belongs to.
	VpcID string `json:"vpc_id"`
}

// PrivateNetwork: private network.
type PrivateNetwork struct {
	// ID: private Network ID.
	ID string `json:"id"`

	// Name: private Network name.
	Name string `json:"name"`

	// OrganizationID: scaleway Organization the Private Network belongs to.
	OrganizationID string `json:"organization_id"`

	// ProjectID: scaleway Project the Private Network belongs to.
	ProjectID string `json:"project_id"`

	// Region: region in which the Private Network is available.
	Region scw.Region `json:"region"`

	// Tags: tags of the Private Network.
	Tags []string `json:"tags"`

	// CreatedAt: date the Private Network was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the Private Network was last modified.
	UpdatedAt *time.Time `json:"updated_at"`

	// Subnets: private Network subnets.
	Subnets []*Subnet `json:"subnets"`

	// VpcID: vPC the Private Network belongs to.
	VpcID string `json:"vpc_id"`

	// DHCPEnabled: defines whether managed DHCP is enabled for this Private Network.
	DHCPEnabled bool `json:"dhcp_enabled"`

	// DefaultRoutePropagationEnabled: defines whether default v4 and v6 routes are propagated for this Private Network.
	DefaultRoutePropagationEnabled bool `json:"default_route_propagation_enabled"`
}

// Route: route.
type Route struct {
	// ID: route ID.
	ID string `json:"id"`

	// Description: route description.
	Description string `json:"description"`

	// Tags: tags of the Route.
	Tags []string `json:"tags"`

	// VpcID: vPC the Route belongs to.
	VpcID string `json:"vpc_id"`

	// Destination: destination of the Route.
	Destination scw.IPNet `json:"destination"`

	// NexthopResourceID: ID of the nexthop resource.
	NexthopResourceID *string `json:"nexthop_resource_id"`

	// NexthopPrivateNetworkID: ID of the nexthop private network.
	NexthopPrivateNetworkID *string `json:"nexthop_private_network_id"`

	// NexthopVpcConnectorID: ID of the nexthop VPC connector.
	NexthopVpcConnectorID *string `json:"nexthop_vpc_connector_id"`

	// CreatedAt: date the Route was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the Route was last modified.
	UpdatedAt *time.Time `json:"updated_at"`

	// IsReadOnly: defines whether the route can be modified or deleted by the user.
	IsReadOnly bool `json:"is_read_only"`

	// Type: type of the Route.
	// Default value: unknown_route_type
	Type *RouteType `json:"type"`

	// Region: region of the Route.
	Region scw.Region `json:"region"`
}

// VPCConnectorPeerInfo: vpc connector peer info.
type VPCConnectorPeerInfo struct {
	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`

	VpcName string `json:"vpc_name"`
}

// ACLRule: acl rule.
type ACLRule struct {
	// Protocol: protocol to which this rule applies.
	// Default value: ANY
	Protocol ACLRuleProtocol `json:"protocol"`

	// Source: source IP range to which this rule applies (CIDR notation with subnet mask).
	Source scw.IPNet `json:"source"`

	// SrcPortLow: starting port of the source port range to which this rule applies (inclusive).
	SrcPortLow uint32 `json:"src_port_low"`

	// SrcPortHigh: ending port of the source port range to which this rule applies (inclusive).
	SrcPortHigh uint32 `json:"src_port_high"`

	// Destination: destination IP range to which this rule applies (CIDR notation with subnet mask).
	Destination scw.IPNet `json:"destination"`

	// DstPortLow: starting port of the destination port range to which this rule applies (inclusive).
	DstPortLow uint32 `json:"dst_port_low"`

	// DstPortHigh: ending port of the destination port range to which this rule applies (inclusive).
	DstPortHigh uint32 `json:"dst_port_high"`

	// Action: policy to apply to the packet.
	// Default value: unknown_action
	Action Action `json:"action"`

	// Description: rule description.
	Description *string `json:"description"`
}

// RouteWithNexthop: route with nexthop.
type RouteWithNexthop struct {
	// Route: route.
	Route *Route `json:"route"`

	// NexthopIP: IP of the route's next hop.
	NexthopIP *net.IP `json:"nexthop_ip"`

	// NexthopName: name of the route's next hop.
	NexthopName *string `json:"nexthop_name"`

	// NexthopResourceType: resource type of the route's next hop.
	// Default value: unknown_type
	NexthopResourceType RouteWithNexthopResourceType `json:"nexthop_resource_type"`
}

// VPCConnector: vpc connector.
type VPCConnector struct {
	// ID: vPC connector ID.
	ID string `json:"id"`

	// Name: vPC connector name.
	Name string `json:"name"`

	// OrganizationID: scaleway Organization the VPC connector belongs to.
	OrganizationID string `json:"organization_id"`

	// ProjectID: scaleway Project the VPC connector belongs to.
	ProjectID string `json:"project_id"`

	// VpcID: vPC the VPC connector belongs to (origin VPC).
	VpcID string `json:"vpc_id"`

	// TargetVpcID: vPC with which the VPC connector is peered (target VPC).
	TargetVpcID string `json:"target_vpc_id"`

	// Status: status of the VPC connector.
	// Default value: unknown_vpc_connector_status
	Status VPCConnectorStatus `json:"status"`

	// PeerInfo: peer info of target VPC. Available when status is Peered.
	PeerInfo *VPCConnectorPeerInfo `json:"peer_info"`

	// Region: region of the VPC connector.
	Region scw.Region `json:"region"`

	// Tags: tags for the VPC connector.
	Tags []string `json:"tags"`

	// CreatedAt: date the VPC connector was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the VPC connector was last modified.
	UpdatedAt *time.Time `json:"updated_at"`
}

// VPC: vpc.
type VPC struct {
	// ID: vPC ID.
	ID string `json:"id"`

	// Name: vPC name.
	Name string `json:"name"`

	// OrganizationID: scaleway Organization the VPC belongs to.
	OrganizationID string `json:"organization_id"`

	// ProjectID: scaleway Project the VPC belongs to.
	ProjectID string `json:"project_id"`

	// Region: region of the VPC.
	Region scw.Region `json:"region"`

	// Tags: tags for the VPC.
	Tags []string `json:"tags"`

	// IsDefault: defines whether the VPC is the default one for its Project.
	IsDefault bool `json:"is_default"`

	// CreatedAt: date the VPC was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the VPC was last modified.
	UpdatedAt *time.Time `json:"updated_at"`

	// PrivateNetworkCount: number of Private Networks within this VPC.
	PrivateNetworkCount uint32 `json:"private_network_count"`

	// RoutingEnabled: defines whether the VPC routes traffic between its Private Networks.
	RoutingEnabled bool `json:"routing_enabled"`

	// CustomRoutesPropagationEnabled: defines whether the VPC advertises custom routes between its Private Networks.
	CustomRoutesPropagationEnabled bool `json:"custom_routes_propagation_enabled"`
}

// AddSubnetsRequest: add subnets request.
type AddSubnetsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PrivateNetworkID: private Network ID.
	PrivateNetworkID string `json:"-"`

	// Subnets: private Network subnets CIDR.
	Subnets []scw.IPNet `json:"subnets"`
}

// AddSubnetsResponse: add subnets response.
type AddSubnetsResponse struct {
	Subnets []scw.IPNet `json:"subnets"`
}

// CreatePrivateNetworkRequest: create private network request.
type CreatePrivateNetworkRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: name for the Private Network.
	Name string `json:"name"`

	// ProjectID: scaleway Project in which to create the Private Network.
	ProjectID string `json:"project_id"`

	// Tags: tags for the Private Network.
	Tags []string `json:"tags"`

	// Subnets: private Network subnets CIDR.
	Subnets []scw.IPNet `json:"subnets"`

	// VpcID: vPC in which to create the Private Network.
	VpcID *string `json:"vpc_id,omitempty"`

	// DefaultRoutePropagationEnabled: defines whether default v4 and v6 routes are propagated for this Private Network.
	DefaultRoutePropagationEnabled bool `json:"default_route_propagation_enabled"`
}

// CreateRouteRequest: create route request.
type CreateRouteRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Description: route description.
	Description string `json:"description"`

	// Tags: tags of the Route.
	Tags []string `json:"tags"`

	// VpcID: vPC the Route belongs to.
	VpcID string `json:"vpc_id"`

	// Destination: destination of the Route.
	Destination scw.IPNet `json:"destination"`

	// NexthopResourceID: ID of the nexthop resource.
	NexthopResourceID *string `json:"nexthop_resource_id,omitempty"`

	// NexthopPrivateNetworkID: ID of the nexthop private network.
	NexthopPrivateNetworkID *string `json:"nexthop_private_network_id,omitempty"`

	// NexthopVpcConnectorID: ID of the nexthop VPC Connector.
	NexthopVpcConnectorID *string `json:"nexthop_vpc_connector_id,omitempty"`
}

// CreateVPCConnectorRequest: create vpc connector request.
type CreateVPCConnectorRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: name for the VPC connector.
	Name string `json:"name"`

	// Tags: tags for the VPC connector.
	Tags []string `json:"tags"`

	// VpcID: vPC ID to filter for. Only connectors belonging to this VPC will be returned.
	VpcID string `json:"vpc_id"`

	// TargetVpcID: target VPC ID to filter for. Only connectors belonging to this target VPC will be returned.
	TargetVpcID string `json:"target_vpc_id"`
}

// CreateVPCRequest: create vpc request.
type CreateVPCRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: name for the VPC.
	Name string `json:"name"`

	// ProjectID: scaleway Project in which to create the VPC.
	ProjectID string `json:"project_id"`

	// Tags: tags for the VPC.
	Tags []string `json:"tags"`

	// EnableRouting: enable routing between Private Networks in the VPC.
	EnableRouting bool `json:"enable_routing"`
}

// DeletePrivateNetworkRequest: delete private network request.
type DeletePrivateNetworkRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PrivateNetworkID: private Network ID.
	PrivateNetworkID string `json:"-"`
}

// DeleteRouteRequest: delete route request.
type DeleteRouteRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// RouteID: route ID.
	RouteID string `json:"-"`
}

// DeleteSubnetsRequest: delete subnets request.
type DeleteSubnetsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PrivateNetworkID: private Network ID.
	PrivateNetworkID string `json:"-"`

	// Subnets: private Network subnets CIDR.
	Subnets []scw.IPNet `json:"subnets"`
}

// DeleteSubnetsResponse: delete subnets response.
type DeleteSubnetsResponse struct {
	Subnets []scw.IPNet `json:"subnets"`
}

// DeleteVPCConnectorRequest: delete vpc connector request.
type DeleteVPCConnectorRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VpcConnectorID: vPC connector ID.
	VpcConnectorID string `json:"-"`
}

// DeleteVPCRequest: delete vpc request.
type DeleteVPCRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VpcID: vPC ID.
	VpcID string `json:"-"`
}

// EnableCustomRoutesPropagationRequest: enable custom routes propagation request.
type EnableCustomRoutesPropagationRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VpcID: vPC ID.
	VpcID string `json:"-"`
}

// EnableDHCPRequest: enable dhcp request.
type EnableDHCPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PrivateNetworkID: private Network ID.
	PrivateNetworkID string `json:"-"`
}

// EnableRoutingRequest: enable routing request.
type EnableRoutingRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VpcID: vPC ID.
	VpcID string `json:"-"`
}

// GetACLRequest: get acl request.
type GetACLRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VpcID: ID of the Network ACL's VPC.
	VpcID string `json:"-"`

	// IsIPv6: defines whether this set of ACL rules is for IPv6 (false = IPv4). Each Network ACL can have rules for only one IP type.
	IsIPv6 bool `json:"is_ipv6"`
}

// GetACLResponse: get acl response.
type GetACLResponse struct {
	Rules []*ACLRule `json:"rules"`

	// DefaultPolicy: default value: unknown_action
	DefaultPolicy Action `json:"default_policy"`
}

// GetPrivateNetworkRequest: get private network request.
type GetPrivateNetworkRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PrivateNetworkID: private Network ID.
	PrivateNetworkID string `json:"-"`
}

// GetRouteRequest: get route request.
type GetRouteRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// RouteID: route ID.
	RouteID string `json:"-"`
}

// GetVPCConnectorRequest: get vpc connector request.
type GetVPCConnectorRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VpcConnectorID: vPC connector ID.
	VpcConnectorID string `json:"-"`
}

// GetVPCRequest: get vpc request.
type GetVPCRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VpcID: vPC ID.
	VpcID string `json:"-"`
}

// ListPrivateNetworksRequest: list private networks request.
type ListPrivateNetworksRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: sort order of the returned Private Networks.
	// Default value: created_at_asc
	OrderBy ListPrivateNetworksRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: maximum number of Private Networks to return per page.
	PageSize *uint32 `json:"-"`

	// Name: name to filter for. Only Private Networks with names containing this string will be returned.
	Name *string `json:"-"`

	// Tags: tags to filter for. Only Private Networks with one or more matching tags will be returned.
	Tags []string `json:"-"`

	// OrganizationID: organization ID to filter for. Only Private Networks belonging to this Organization will be returned.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID to filter for. Only Private Networks belonging to this Project will be returned.
	ProjectID *string `json:"-"`

	// PrivateNetworkIDs: private Network IDs to filter for. Only Private Networks with one of these IDs will be returned.
	PrivateNetworkIDs []string `json:"-"`

	// VpcID: vPC ID to filter for. Only Private Networks belonging to this VPC will be returned.
	VpcID *string `json:"-"`

	// DHCPEnabled: DHCP status to filter for. When true, only Private Networks with managed DHCP enabled will be returned.
	DHCPEnabled *bool `json:"-"`
}

// ListPrivateNetworksResponse: list private networks response.
type ListPrivateNetworksResponse struct {
	PrivateNetworks []*PrivateNetwork `json:"private_networks"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPrivateNetworksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPrivateNetworksResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListPrivateNetworksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PrivateNetworks = append(r.PrivateNetworks, results.PrivateNetworks...)
	r.TotalCount += uint32(len(results.PrivateNetworks))
	return uint32(len(results.PrivateNetworks)), nil
}

// ListRoutesWithNexthopResponse: list routes with nexthop response.
type ListRoutesWithNexthopResponse struct {
	// Routes: list of routes.
	Routes []*RouteWithNexthop `json:"routes"`

	// TotalCount: total number of routes.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRoutesWithNexthopResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRoutesWithNexthopResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListRoutesWithNexthopResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Routes = append(r.Routes, results.Routes...)
	r.TotalCount += uint64(len(results.Routes))
	return uint64(len(results.Routes)), nil
}

// ListSubnetsRequest: list subnets request.
type ListSubnetsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: sort order of the returned subnets.
	// Default value: created_at_asc
	OrderBy ListSubnetsRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: maximum number of Private Networks to return per page.
	PageSize *uint32 `json:"-"`

	// OrganizationID: organization ID to filter for. Only subnets belonging to this Organization will be returned.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID to filter for. Only subnets belonging to this Project will be returned.
	ProjectID *string `json:"-"`

	// SubnetIDs: subnet IDs to filter for. Only subnets matching the specified IDs will be returned.
	SubnetIDs []string `json:"-"`

	// VpcID: vPC ID to filter for. Only subnets belonging to this VPC will be returned.
	VpcID *string `json:"-"`
}

// ListSubnetsResponse: list subnets response.
type ListSubnetsResponse struct {
	Subnets []*Subnet `json:"subnets"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSubnetsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSubnetsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListSubnetsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Subnets = append(r.Subnets, results.Subnets...)
	r.TotalCount += uint32(len(results.Subnets))
	return uint32(len(results.Subnets)), nil
}

// ListVPCConnectorsRequest: list vpc connectors request.
type ListVPCConnectorsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: sort order of the returned VPC connectors.
	// Default value: created_at_asc
	OrderBy ListVPCConnectorsRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: maximum number of VPC connectors to return per page.
	PageSize *uint32 `json:"-"`

	// Name: name to filter for. Only connectors with names containing this string will be returned.
	Name *string `json:"-"`

	// Tags: tags to filter for. Only connectors with one or more matching tags will be returned.
	Tags []string `json:"-"`

	// OrganizationID: organization ID to filter for. Only connectors belonging to this Organization will be returned.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID to filter for. Only connectors belonging to this Project will be returned.
	ProjectID *string `json:"-"`

	// VpcID: vPC ID to filter for. Only connectors belonging to this VPC will be returned.
	VpcID *string `json:"-"`

	// TargetVpcID: target VPC ID to filter for. Only connectors belonging to this target VPC will be returned.
	TargetVpcID *string `json:"-"`

	// Status: status of the VPC connector.
	// Default value: unknown_vpc_connector_status
	Status *VPCConnectorStatus `json:"-"`
}

// ListVPCConnectorsResponse: list vpc connectors response.
type ListVPCConnectorsResponse struct {
	VpcConnectors []*VPCConnector `json:"vpc_connectors"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVPCConnectorsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVPCConnectorsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListVPCConnectorsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.VpcConnectors = append(r.VpcConnectors, results.VpcConnectors...)
	r.TotalCount += uint32(len(results.VpcConnectors))
	return uint32(len(results.VpcConnectors)), nil
}

// ListVPCsRequest: list vp cs request.
type ListVPCsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: sort order of the returned VPCs.
	// Default value: created_at_asc
	OrderBy ListVPCsRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: maximum number of VPCs to return per page.
	PageSize *uint32 `json:"-"`

	// Name: name to filter for. Only VPCs with names containing this string will be returned.
	Name *string `json:"-"`

	// Tags: tags to filter for. Only VPCs with one or more matching tags will be returned.
	Tags []string `json:"-"`

	// OrganizationID: organization ID to filter for. Only VPCs belonging to this Organization will be returned.
	OrganizationID *string `json:"-"`

	// ProjectID: project ID to filter for. Only VPCs belonging to this Project will be returned.
	ProjectID *string `json:"-"`

	// IsDefault: defines whether to filter only for VPCs which are the default one for their Project.
	IsDefault *bool `json:"-"`

	// RoutingEnabled: defines whether to filter only for VPCs which route traffic between their Private Networks.
	RoutingEnabled *bool `json:"-"`
}

// ListVPCsResponse: list vp cs response.
type ListVPCsResponse struct {
	Vpcs []*VPC `json:"vpcs"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVPCsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVPCsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListVPCsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Vpcs = append(r.Vpcs, results.Vpcs...)
	r.TotalCount += uint32(len(results.Vpcs))
	return uint32(len(results.Vpcs)), nil
}

// RoutesWithNexthopAPIListRoutesWithNexthopRequest: routes with nexthop api list routes with nexthop request.
type RoutesWithNexthopAPIListRoutesWithNexthopRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: sort order of the returned routes.
	// Default value: created_at_asc
	OrderBy ListRoutesWithNexthopRequestOrderBy `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: maximum number of routes to return per page.
	PageSize *uint32 `json:"-"`

	// VpcID: vPC to filter for. Only routes within this VPC will be returned.
	VpcID *string `json:"-"`

	// NexthopResourceID: next hop resource ID to filter for. Only routes with a matching next hop resource ID will be returned.
	NexthopResourceID *string `json:"-"`

	// NexthopPrivateNetworkID: next hop private network ID to filter for. Only routes with a matching next hop private network ID will be returned.
	NexthopPrivateNetworkID *string `json:"-"`

	// NexthopResourceType: next hop resource type to filter for. Only Routes with a matching next hop resource type will be returned.
	// Default value: unknown_type
	NexthopResourceType RouteWithNexthopResourceType `json:"-"`

	// NexthopVpcConnectorID: next hop VPC connector ID to filter for. Only routes with a matching next hop VPC connector ID will be returned.
	NexthopVpcConnectorID *string `json:"-"`

	// Contains: only routes whose destination is contained in this subnet will be returned.
	Contains *scw.IPNet `json:"-"`

	// Tags: tags to filter for, only routes with one or more matching tags will be returned.
	Tags []string `json:"-"`

	// IsIPv6: only routes with an IPv6 destination will be returned.
	IsIPv6 *bool `json:"-"`
}

// SetACLRequest: set acl request.
type SetACLRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VpcID: ID of the Network ACL's VPC.
	VpcID string `json:"-"`

	// Rules: list of Network ACL rules.
	Rules []*ACLRule `json:"rules"`

	// IsIPv6: defines whether this set of ACL rules is for IPv6 (false = IPv4). Each Network ACL can have rules for only one IP type.
	IsIPv6 bool `json:"is_ipv6"`

	// DefaultPolicy: action to take for packets which do not match any rules.
	// Default value: unknown_action
	DefaultPolicy Action `json:"default_policy"`
}

// SetACLResponse: set acl response.
type SetACLResponse struct {
	Rules []*ACLRule `json:"rules"`

	// DefaultPolicy: default value: unknown_action
	DefaultPolicy Action `json:"default_policy"`
}

// UpdatePrivateNetworkRequest: update private network request.
type UpdatePrivateNetworkRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PrivateNetworkID: private Network ID.
	PrivateNetworkID string `json:"-"`

	// Name: name for the Private Network.
	Name *string `json:"name,omitempty"`

	// Tags: tags for the Private Network.
	Tags *[]string `json:"tags,omitempty"`

	// DefaultRoutePropagationEnabled: defines whether default v4 and v6 routes are propagated for this Private Network.
	DefaultRoutePropagationEnabled *bool `json:"default_route_propagation_enabled,omitempty"`
}

// UpdateRouteRequest: update route request.
type UpdateRouteRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// RouteID: route ID.
	RouteID string `json:"-"`

	// Description: route description.
	Description *string `json:"description,omitempty"`

	// Tags: tags of the Route.
	Tags *[]string `json:"tags,omitempty"`

	// Destination: destination of the Route.
	Destination *scw.IPNet `json:"destination,omitempty"`

	// NexthopResourceID: ID of the nexthop resource.
	NexthopResourceID *string `json:"nexthop_resource_id,omitempty"`

	// NexthopPrivateNetworkID: ID of the nexthop private network.
	NexthopPrivateNetworkID *string `json:"nexthop_private_network_id,omitempty"`

	// NexthopVpcConnectorID: ID of the nexthop VPC connector.
	NexthopVpcConnectorID *string `json:"nexthop_vpc_connector_id,omitempty"`
}

// UpdateVPCConnectorRequest: update vpc connector request.
type UpdateVPCConnectorRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VpcConnectorID: vPC connector ID.
	VpcConnectorID string `json:"-"`

	// Name: name for the VPC connector.
	Name *string `json:"name,omitempty"`

	// Tags: tags for the VPC connector.
	Tags *[]string `json:"tags,omitempty"`
}

// UpdateVPCRequest: update vpc request.
type UpdateVPCRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// VpcID: vPC ID.
	VpcID string `json:"-"`

	// Name: name for the VPC.
	Name *string `json:"name,omitempty"`

	// Tags: tags for the VPC.
	Tags *[]string `json:"tags,omitempty"`
}

// This API allows you to manage your Virtual Private Clouds (VPCs) and Private Networks.
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
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// ListVPCs: List existing VPCs in the specified region.
func (s *API) ListVPCs(req *ListVPCsRequest, opts ...scw.RequestOption) (*ListVPCsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "is_default", req.IsDefault)
	parameter.AddToQuery(query, "routing_enabled", req.RoutingEnabled)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs",
		Query:  query,
	}

	var resp ListVPCsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVPC: Create a new VPC in the specified region.
func (s *API) CreateVPC(req *CreateVPCRequest, opts ...scw.RequestOption) (*VPC, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("vpc")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp VPC

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVPC: Retrieve details of an existing VPC, specified by its VPC ID.
func (s *API) GetVPC(req *GetVPCRequest, opts ...scw.RequestOption) (*VPC, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcID) == "" {
		return nil, errors.New("field VpcID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs/" + fmt.Sprint(req.VpcID) + "",
	}

	var resp VPC

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVPC: Update parameters including name and tags of the specified VPC.
func (s *API) UpdateVPC(req *UpdateVPCRequest, opts ...scw.RequestOption) (*VPC, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcID) == "" {
		return nil, errors.New("field VpcID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs/" + fmt.Sprint(req.VpcID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp VPC

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVPC: Delete a VPC specified by its VPC ID.
func (s *API) DeleteVPC(req *DeleteVPCRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcID) == "" {
		return errors.New("field VpcID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs/" + fmt.Sprint(req.VpcID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListPrivateNetworks: List existing Private Networks in the specified region. By default, the Private Networks returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field.
func (s *API) ListPrivateNetworks(req *ListPrivateNetworksRequest, opts ...scw.RequestOption) (*ListPrivateNetworksResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "private_network_ids", req.PrivateNetworkIDs)
	parameter.AddToQuery(query, "vpc_id", req.VpcID)
	parameter.AddToQuery(query, "dhcp_enabled", req.DHCPEnabled)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks",
		Query:  query,
	}

	var resp ListPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePrivateNetwork: Create a new Private Network. Once created, you can attach Scaleway resources which are in the same region.
func (s *API) CreatePrivateNetwork(req *CreatePrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("pn")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrivateNetwork: Retrieve information about an existing Private Network, specified by its Private Network ID. Its full details are returned in the response object.
func (s *API) GetPrivateNetwork(req *GetPrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePrivateNetwork: Update parameters (such as name or tags) of an existing Private Network, specified by its Private Network ID.
func (s *API) UpdatePrivateNetwork(req *UpdatePrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePrivateNetwork: Delete an existing Private Network. Note that you must first detach all resources from the network, in order to delete it.
func (s *API) DeletePrivateNetwork(req *DeletePrivateNetworkRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// EnableDHCP: Enable DHCP managed on an existing Private Network. Note that you will not be able to deactivate it afterwards.
func (s *API) EnableDHCP(req *EnableDHCPRequest, opts ...scw.RequestOption) (*PrivateNetwork, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "/enable-dhcp",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp PrivateNetwork

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableRouting: Enable routing on an existing VPC. Note that you will not be able to deactivate it afterwards.
func (s *API) EnableRouting(req *EnableRoutingRequest, opts ...scw.RequestOption) (*VPC, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcID) == "" {
		return nil, errors.New("field VpcID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs/" + fmt.Sprint(req.VpcID) + "/enable-routing",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp VPC

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableCustomRoutesPropagation: Enable custom routes propagation on an existing VPC. Note that you will not be able to deactivate it afterwards.
func (s *API) EnableCustomRoutesPropagation(req *EnableCustomRoutesPropagationRequest, opts ...scw.RequestOption) (*VPC, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcID) == "" {
		return nil, errors.New("field VpcID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs/" + fmt.Sprint(req.VpcID) + "/enable-custom-routes-propagation",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp VPC

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSubnets: List any Private Network's subnets. See ListPrivateNetworks to list a specific Private Network's subnets.
func (s *API) ListSubnets(req *ListSubnetsRequest, opts ...scw.RequestOption) (*ListSubnetsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "subnet_ids", req.SubnetIDs)
	parameter.AddToQuery(query, "vpc_id", req.VpcID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/subnets",
		Query:  query,
	}

	var resp ListSubnetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSubnets: Add new subnets to an existing Private Network.
func (s *API) AddSubnets(req *AddSubnetsRequest, opts ...scw.RequestOption) (*AddSubnetsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "/subnets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddSubnetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSubnets: Delete the specified subnets from a Private Network.
func (s *API) DeleteSubnets(req *DeleteSubnetsRequest, opts ...scw.RequestOption) (*DeleteSubnetsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "/subnets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DeleteSubnetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRoute: Create a new custom Route.
func (s *API) CreateRoute(req *CreateRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/routes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRoute: Retrieve details of an existing Route, specified by its Route ID.
func (s *API) GetRoute(req *GetRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.RouteID) == "" {
		return nil, errors.New("field RouteID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateRoute: Update parameters of the specified Route.
func (s *API) UpdateRoute(req *UpdateRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.RouteID) == "" {
		return nil, errors.New("field RouteID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteRoute: Delete a Route specified by its Route ID.
func (s *API) DeleteRoute(req *DeleteRouteRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.RouteID) == "" {
		return errors.New("field RouteID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetACL: Retrieve a list of ACL rules for a VPC, specified by its VPC ID.
func (s *API) GetACL(req *GetACLRequest, opts ...scw.RequestOption) (*GetACLResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "is_ipv6", req.IsIPv6)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcID) == "" {
		return nil, errors.New("field VpcID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs/" + fmt.Sprint(req.VpcID) + "/acl-rules",
		Query:  query,
	}

	var resp GetACLResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetACL: Set the list of ACL rules and the default routing policy for a VPC.
func (s *API) SetACL(req *SetACLRequest, opts ...scw.RequestOption) (*SetACLResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcID) == "" {
		return nil, errors.New("field VpcID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpcs/" + fmt.Sprint(req.VpcID) + "/acl-rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetACLResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVPCConnectors: List existing VPC connectors in the specified region.
func (s *API) ListVPCConnectors(req *ListVPCConnectorsRequest, opts ...scw.RequestOption) (*ListVPCConnectorsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "vpc_id", req.VpcID)
	parameter.AddToQuery(query, "target_vpc_id", req.TargetVpcID)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpc-connectors",
		Query:  query,
	}

	var resp ListVPCConnectorsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVPCConnector: Create a new VPC connector in the specified region.
func (s *API) CreateVPCConnector(req *CreateVPCConnectorRequest, opts ...scw.RequestOption) (*VPCConnector, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("VPCConnector")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpc-connectors",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp VPCConnector

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVPCConnector: Retrieve details of an existing VPC connector, specified by its VPC connector ID.
func (s *API) GetVPCConnector(req *GetVPCConnectorRequest, opts ...scw.RequestOption) (*VPCConnector, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcConnectorID) == "" {
		return nil, errors.New("field VpcConnectorID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpc-connectors/" + fmt.Sprint(req.VpcConnectorID) + "",
	}

	var resp VPCConnector

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVPCConnector: Update parameters including name and tags of the specified VPC connector.
func (s *API) UpdateVPCConnector(req *UpdateVPCConnectorRequest, opts ...scw.RequestOption) (*VPCConnector, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcConnectorID) == "" {
		return nil, errors.New("field VpcConnectorID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpc-connectors/" + fmt.Sprint(req.VpcConnectorID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp VPCConnector

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVPCConnector: Delete a VPC connector specified by its VPC connector ID.
func (s *API) DeleteVPCConnector(req *DeleteVPCConnectorRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.VpcConnectorID) == "" {
		return errors.New("field VpcConnectorID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/vpc-connectors/" + fmt.Sprint(req.VpcConnectorID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type RoutesWithNexthopAPI struct {
	client *scw.Client
}

// NewRoutesWithNexthopAPI returns a RoutesWithNexthopAPI object from a Scaleway client.
func NewRoutesWithNexthopAPI(client *scw.Client) *RoutesWithNexthopAPI {
	return &RoutesWithNexthopAPI{
		client: client,
	}
}

// ListRoutesWithNexthop: Return routes with associated next hop data.
func (s *RoutesWithNexthopAPI) ListRoutesWithNexthop(req *RoutesWithNexthopAPIListRoutesWithNexthopRequest, opts ...scw.RequestOption) (*ListRoutesWithNexthopResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "vpc_id", req.VpcID)
	parameter.AddToQuery(query, "nexthop_resource_id", req.NexthopResourceID)
	parameter.AddToQuery(query, "nexthop_private_network_id", req.NexthopPrivateNetworkID)
	parameter.AddToQuery(query, "nexthop_resource_type", req.NexthopResourceType)
	parameter.AddToQuery(query, "nexthop_vpc_connector_id", req.NexthopVpcConnectorID)
	parameter.AddToQuery(query, "contains", req.Contains)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "is_ipv6", req.IsIPv6)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc/v2/regions/" + fmt.Sprint(req.Region) + "/routes-with-nexthop",
		Query:  query,
	}

	var resp ListRoutesWithNexthopResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

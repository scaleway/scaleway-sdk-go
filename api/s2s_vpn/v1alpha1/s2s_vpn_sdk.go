// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package s2s_vpn provides methods and message types of the s2s_vpn v1alpha1 API.
package s2s_vpn

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
	defaultS2sVpnRetryInterval = 15 * time.Second
	defaultS2sVpnTimeout       = 5 * time.Minute
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

type BgpSessionStatus string

const (
	BgpSessionStatusUnknownStatus = BgpSessionStatus("unknown_status")
	BgpSessionStatusUp            = BgpSessionStatus("up")
	BgpSessionStatusDown          = BgpSessionStatus("down")
	BgpSessionStatusDisabled      = BgpSessionStatus("disabled")
)

func (enum BgpSessionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(BgpSessionStatusUnknownStatus)
	}
	return string(enum)
}

func (enum BgpSessionStatus) Values() []BgpSessionStatus {
	return []BgpSessionStatus{
		"unknown_status",
		"up",
		"down",
		"disabled",
	}
}

func (enum BgpSessionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BgpSessionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BgpSessionStatus(BgpSessionStatus(tmp).String())
	return nil
}

type ConnectionDhGroup string

const (
	ConnectionDhGroupUnknownDhgroup = ConnectionDhGroup("unknown_dhgroup")
	ConnectionDhGroupModp2048       = ConnectionDhGroup("modp2048")
	ConnectionDhGroupModp3072       = ConnectionDhGroup("modp3072")
	ConnectionDhGroupModp4096       = ConnectionDhGroup("modp4096")
	ConnectionDhGroupEcp256         = ConnectionDhGroup("ecp256")
	ConnectionDhGroupEcp384         = ConnectionDhGroup("ecp384")
	ConnectionDhGroupEcp521         = ConnectionDhGroup("ecp521")
	ConnectionDhGroupCurve25519     = ConnectionDhGroup("curve25519")
)

func (enum ConnectionDhGroup) String() string {
	if enum == "" {
		// return default value if empty
		return string(ConnectionDhGroupUnknownDhgroup)
	}
	return string(enum)
}

func (enum ConnectionDhGroup) Values() []ConnectionDhGroup {
	return []ConnectionDhGroup{
		"unknown_dhgroup",
		"modp2048",
		"modp3072",
		"modp4096",
		"ecp256",
		"ecp384",
		"ecp521",
		"curve25519",
	}
}

func (enum ConnectionDhGroup) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ConnectionDhGroup) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ConnectionDhGroup(ConnectionDhGroup(tmp).String())
	return nil
}

type ConnectionEncryption string

const (
	ConnectionEncryptionUnknownEncryption = ConnectionEncryption("unknown_encryption")
	ConnectionEncryptionAes128            = ConnectionEncryption("aes128")
	ConnectionEncryptionAes192            = ConnectionEncryption("aes192")
	ConnectionEncryptionAes256            = ConnectionEncryption("aes256")
	ConnectionEncryptionAes128gcm         = ConnectionEncryption("aes128gcm")
	ConnectionEncryptionAes192gcm         = ConnectionEncryption("aes192gcm")
	ConnectionEncryptionAes256gcm         = ConnectionEncryption("aes256gcm")
	ConnectionEncryptionAes128ccm         = ConnectionEncryption("aes128ccm")
	ConnectionEncryptionAes256ccm         = ConnectionEncryption("aes256ccm")
	ConnectionEncryptionChacha20poly1305  = ConnectionEncryption("chacha20poly1305")
)

func (enum ConnectionEncryption) String() string {
	if enum == "" {
		// return default value if empty
		return string(ConnectionEncryptionUnknownEncryption)
	}
	return string(enum)
}

func (enum ConnectionEncryption) Values() []ConnectionEncryption {
	return []ConnectionEncryption{
		"unknown_encryption",
		"aes128",
		"aes192",
		"aes256",
		"aes128gcm",
		"aes192gcm",
		"aes256gcm",
		"aes128ccm",
		"aes256ccm",
		"chacha20poly1305",
	}
}

func (enum ConnectionEncryption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ConnectionEncryption) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ConnectionEncryption(ConnectionEncryption(tmp).String())
	return nil
}

type ConnectionInitiationPolicy string

const (
	ConnectionInitiationPolicyUnknownInitiationPolicy = ConnectionInitiationPolicy("unknown_initiation_policy")
	ConnectionInitiationPolicyVpnGateway              = ConnectionInitiationPolicy("vpn_gateway")
	ConnectionInitiationPolicyCustomerGateway         = ConnectionInitiationPolicy("customer_gateway")
)

func (enum ConnectionInitiationPolicy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ConnectionInitiationPolicyUnknownInitiationPolicy)
	}
	return string(enum)
}

func (enum ConnectionInitiationPolicy) Values() []ConnectionInitiationPolicy {
	return []ConnectionInitiationPolicy{
		"unknown_initiation_policy",
		"vpn_gateway",
		"customer_gateway",
	}
}

func (enum ConnectionInitiationPolicy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ConnectionInitiationPolicy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ConnectionInitiationPolicy(ConnectionInitiationPolicy(tmp).String())
	return nil
}

type ConnectionIntegrity string

const (
	ConnectionIntegrityUnknownIntegrity = ConnectionIntegrity("unknown_integrity")
	ConnectionIntegritySha256           = ConnectionIntegrity("sha256")
	ConnectionIntegritySha384           = ConnectionIntegrity("sha384")
	ConnectionIntegritySha512           = ConnectionIntegrity("sha512")
)

func (enum ConnectionIntegrity) String() string {
	if enum == "" {
		// return default value if empty
		return string(ConnectionIntegrityUnknownIntegrity)
	}
	return string(enum)
}

func (enum ConnectionIntegrity) Values() []ConnectionIntegrity {
	return []ConnectionIntegrity{
		"unknown_integrity",
		"sha256",
		"sha384",
		"sha512",
	}
}

func (enum ConnectionIntegrity) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ConnectionIntegrity) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ConnectionIntegrity(ConnectionIntegrity(tmp).String())
	return nil
}

type ConnectionStatus string

const (
	ConnectionStatusUnknownStatus       = ConnectionStatus("unknown_status")
	ConnectionStatusActive              = ConnectionStatus("active")
	ConnectionStatusLimitedConnectivity = ConnectionStatus("limited_connectivity")
	ConnectionStatusDown                = ConnectionStatus("down")
	ConnectionStatusLocked              = ConnectionStatus("locked")
)

func (enum ConnectionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ConnectionStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ConnectionStatus) Values() []ConnectionStatus {
	return []ConnectionStatus{
		"unknown_status",
		"active",
		"limited_connectivity",
		"down",
		"locked",
	}
}

func (enum ConnectionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ConnectionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ConnectionStatus(ConnectionStatus(tmp).String())
	return nil
}

type CreateConnectionRequestInitiationPolicy string

const (
	CreateConnectionRequestInitiationPolicyUnknownInitiationPolicy = CreateConnectionRequestInitiationPolicy("unknown_initiation_policy")
	CreateConnectionRequestInitiationPolicyVpnGateway              = CreateConnectionRequestInitiationPolicy("vpn_gateway")
	CreateConnectionRequestInitiationPolicyCustomerGateway         = CreateConnectionRequestInitiationPolicy("customer_gateway")
)

func (enum CreateConnectionRequestInitiationPolicy) String() string {
	if enum == "" {
		// return default value if empty
		return string(CreateConnectionRequestInitiationPolicyUnknownInitiationPolicy)
	}
	return string(enum)
}

func (enum CreateConnectionRequestInitiationPolicy) Values() []CreateConnectionRequestInitiationPolicy {
	return []CreateConnectionRequestInitiationPolicy{
		"unknown_initiation_policy",
		"vpn_gateway",
		"customer_gateway",
	}
}

func (enum CreateConnectionRequestInitiationPolicy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CreateConnectionRequestInitiationPolicy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CreateConnectionRequestInitiationPolicy(CreateConnectionRequestInitiationPolicy(tmp).String())
	return nil
}

type ListConnectionsRequestOrderBy string

const (
	ListConnectionsRequestOrderByCreatedAtAsc  = ListConnectionsRequestOrderBy("created_at_asc")
	ListConnectionsRequestOrderByCreatedAtDesc = ListConnectionsRequestOrderBy("created_at_desc")
	ListConnectionsRequestOrderByNameAsc       = ListConnectionsRequestOrderBy("name_asc")
	ListConnectionsRequestOrderByNameDesc      = ListConnectionsRequestOrderBy("name_desc")
	ListConnectionsRequestOrderByStatusAsc     = ListConnectionsRequestOrderBy("status_asc")
	ListConnectionsRequestOrderByStatusDesc    = ListConnectionsRequestOrderBy("status_desc")
)

func (enum ListConnectionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListConnectionsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListConnectionsRequestOrderBy) Values() []ListConnectionsRequestOrderBy {
	return []ListConnectionsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
		"status_asc",
		"status_desc",
	}
}

func (enum ListConnectionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListConnectionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListConnectionsRequestOrderBy(ListConnectionsRequestOrderBy(tmp).String())
	return nil
}

type ListCustomerGatewaysRequestOrderBy string

const (
	ListCustomerGatewaysRequestOrderByCreatedAtAsc  = ListCustomerGatewaysRequestOrderBy("created_at_asc")
	ListCustomerGatewaysRequestOrderByCreatedAtDesc = ListCustomerGatewaysRequestOrderBy("created_at_desc")
	ListCustomerGatewaysRequestOrderByNameAsc       = ListCustomerGatewaysRequestOrderBy("name_asc")
	ListCustomerGatewaysRequestOrderByNameDesc      = ListCustomerGatewaysRequestOrderBy("name_desc")
)

func (enum ListCustomerGatewaysRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListCustomerGatewaysRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListCustomerGatewaysRequestOrderBy) Values() []ListCustomerGatewaysRequestOrderBy {
	return []ListCustomerGatewaysRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListCustomerGatewaysRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCustomerGatewaysRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCustomerGatewaysRequestOrderBy(ListCustomerGatewaysRequestOrderBy(tmp).String())
	return nil
}

type ListRoutingPoliciesRequestOrderBy string

const (
	ListRoutingPoliciesRequestOrderByCreatedAtAsc  = ListRoutingPoliciesRequestOrderBy("created_at_asc")
	ListRoutingPoliciesRequestOrderByCreatedAtDesc = ListRoutingPoliciesRequestOrderBy("created_at_desc")
	ListRoutingPoliciesRequestOrderByNameAsc       = ListRoutingPoliciesRequestOrderBy("name_asc")
	ListRoutingPoliciesRequestOrderByNameDesc      = ListRoutingPoliciesRequestOrderBy("name_desc")
)

func (enum ListRoutingPoliciesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRoutingPoliciesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListRoutingPoliciesRequestOrderBy) Values() []ListRoutingPoliciesRequestOrderBy {
	return []ListRoutingPoliciesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListRoutingPoliciesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRoutingPoliciesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRoutingPoliciesRequestOrderBy(ListRoutingPoliciesRequestOrderBy(tmp).String())
	return nil
}

type ListVpnGatewaysRequestOrderBy string

const (
	ListVpnGatewaysRequestOrderByCreatedAtAsc  = ListVpnGatewaysRequestOrderBy("created_at_asc")
	ListVpnGatewaysRequestOrderByCreatedAtDesc = ListVpnGatewaysRequestOrderBy("created_at_desc")
	ListVpnGatewaysRequestOrderByNameAsc       = ListVpnGatewaysRequestOrderBy("name_asc")
	ListVpnGatewaysRequestOrderByNameDesc      = ListVpnGatewaysRequestOrderBy("name_desc")
	ListVpnGatewaysRequestOrderByTypeAsc       = ListVpnGatewaysRequestOrderBy("type_asc")
	ListVpnGatewaysRequestOrderByTypeDesc      = ListVpnGatewaysRequestOrderBy("type_desc")
	ListVpnGatewaysRequestOrderByStatusAsc     = ListVpnGatewaysRequestOrderBy("status_asc")
	ListVpnGatewaysRequestOrderByStatusDesc    = ListVpnGatewaysRequestOrderBy("status_desc")
)

func (enum ListVpnGatewaysRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListVpnGatewaysRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListVpnGatewaysRequestOrderBy) Values() []ListVpnGatewaysRequestOrderBy {
	return []ListVpnGatewaysRequestOrderBy{
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

func (enum ListVpnGatewaysRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListVpnGatewaysRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListVpnGatewaysRequestOrderBy(ListVpnGatewaysRequestOrderBy(tmp).String())
	return nil
}

type TunnelStatus string

const (
	TunnelStatusUnknownTunnelStatus = TunnelStatus("unknown_tunnel_status")
	TunnelStatusUp                  = TunnelStatus("up")
	TunnelStatusDown                = TunnelStatus("down")
)

func (enum TunnelStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(TunnelStatusUnknownTunnelStatus)
	}
	return string(enum)
}

func (enum TunnelStatus) Values() []TunnelStatus {
	return []TunnelStatus{
		"unknown_tunnel_status",
		"up",
		"down",
	}
}

func (enum TunnelStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TunnelStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TunnelStatus(TunnelStatus(tmp).String())
	return nil
}

type VpnGatewayStatus string

const (
	VpnGatewayStatusUnknownStatus  = VpnGatewayStatus("unknown_status")
	VpnGatewayStatusConfiguring    = VpnGatewayStatus("configuring")
	VpnGatewayStatusFailed         = VpnGatewayStatus("failed")
	VpnGatewayStatusProvisioning   = VpnGatewayStatus("provisioning")
	VpnGatewayStatusActive         = VpnGatewayStatus("active")
	VpnGatewayStatusDeprovisioning = VpnGatewayStatus("deprovisioning")
	VpnGatewayStatusLocked         = VpnGatewayStatus("locked")
)

func (enum VpnGatewayStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(VpnGatewayStatusUnknownStatus)
	}
	return string(enum)
}

func (enum VpnGatewayStatus) Values() []VpnGatewayStatus {
	return []VpnGatewayStatus{
		"unknown_status",
		"configuring",
		"failed",
		"provisioning",
		"active",
		"deprovisioning",
		"locked",
	}
}

func (enum VpnGatewayStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *VpnGatewayStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = VpnGatewayStatus(VpnGatewayStatus(tmp).String())
	return nil
}

// BgpSession: bgp session.
type BgpSession struct {
	RoutingPolicyID string `json:"routing_policy_id"`

	PrivateIP scw.IPNet `json:"private_ip"`

	PeerPrivateIP scw.IPNet `json:"peer_private_ip"`
}

// ConnectionCipher: connection cipher.
type ConnectionCipher struct {
	// Encryption: default value: unknown_encryption
	Encryption ConnectionEncryption `json:"encryption"`

	// Integrity: default value: unknown_integrity
	Integrity *ConnectionIntegrity `json:"integrity"`

	// DhGroup: default value: unknown_dhgroup
	DhGroup *ConnectionDhGroup `json:"dh_group"`
}

// VpnGatewayPrivateConfig: vpn gateway private config.
type VpnGatewayPrivateConfig struct{}

// VpnGatewayPublicConfig: vpn gateway public config.
type VpnGatewayPublicConfig struct {
	IpamIPv4ID *string `json:"ipam_ipv4_id"`

	IpamIPv6ID *string `json:"ipam_ipv6_id"`
}

// CreateConnectionRequestBgpConfig: create connection request bgp config.
type CreateConnectionRequestBgpConfig struct {
	RoutingPolicyID string `json:"routing_policy_id"`

	PrivateIP *scw.IPNet `json:"private_ip"`

	PeerPrivateIP *scw.IPNet `json:"peer_private_ip"`
}

// Connection: connection.
type Connection struct {
	// ID: unique identifier of the connection.
	ID string `json:"id"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// OrganizationID: organization ID.
	OrganizationID string `json:"organization_id"`

	// Name: name of the connection.
	Name string `json:"name"`

	// Tags: list of tags applied to the connection.
	Tags []string `json:"tags"`

	// CreatedAt: creation date of the connection.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the connection.
	UpdatedAt *time.Time `json:"updated_at"`

	// Status: status of the connection.
	// Default value: unknown_status
	Status ConnectionStatus `json:"status"`

	// IsIPv6: IP version of the IPSec Tunnel.
	IsIPv6 bool `json:"is_ipv6"`

	// InitiationPolicy: who initiates the IPsec tunnel.
	// Default value: unknown_initiation_policy
	InitiationPolicy ConnectionInitiationPolicy `json:"initiation_policy"`

	// SecretID: ID of the secret in Secret Manager which contains the PSK.
	SecretID string `json:"secret_id"`

	// SecretRevision: version number of the secret in Secret Manager which contains the PSK.
	SecretRevision uint32 `json:"secret_revision"`

	// Ikev2Ciphers: list of IKE v2 ciphers proposed for the IPsec tunnel.
	Ikev2Ciphers []*ConnectionCipher `json:"ikev2_ciphers"`

	// EspCiphers: list of ESP ciphers proposed for the IPsec tunnel.
	EspCiphers []*ConnectionCipher `json:"esp_ciphers"`

	// RoutePropagationEnabled: defines whether route propagation is enabled or not.
	RoutePropagationEnabled bool `json:"route_propagation_enabled"`

	// VpnGatewayID: ID of the VPN gateway attached to the connection.
	VpnGatewayID string `json:"vpn_gateway_id"`

	// CustomerGatewayID: ID of the customer gateway attached to the connection.
	CustomerGatewayID string `json:"customer_gateway_id"`

	// TunnelStatus: status of the IPsec tunnel.
	// Default value: unknown_tunnel_status
	TunnelStatus TunnelStatus `json:"tunnel_status"`

	// Deprecated: TunnelStatusIPv4: status of the IPv4 IPsec tunnel.
	// Default value: unknown_tunnel_status
	TunnelStatusIPv4 *TunnelStatus `json:"tunnel_status_ipv4,omitempty"`

	// Deprecated: TunnelStatusIPv6: status of the IPv6 IPsec tunnel.
	// Default value: unknown_tunnel_status
	TunnelStatusIPv6 *TunnelStatus `json:"tunnel_status_ipv6,omitempty"`

	// BgpStatusIPv4: status of the BGP IPv4 session.
	// Default value: unknown_status
	BgpStatusIPv4 BgpSessionStatus `json:"bgp_status_ipv4"`

	// BgpStatusIPv6: status of the BGP IPv6 session.
	// Default value: unknown_status
	BgpStatusIPv6 BgpSessionStatus `json:"bgp_status_ipv6"`

	// BgpSessionIPv4: bGP IPv4 session, including status, interco private IPv4 subnet and attached routing policy.
	BgpSessionIPv4 *BgpSession `json:"bgp_session_ipv4"`

	// BgpSessionIPv6: bGP IPv6 session, including status, interco private IPv6 subnet and attached routing policy.
	BgpSessionIPv6 *BgpSession `json:"bgp_session_ipv6"`

	// Region: region of the connection.
	Region scw.Region `json:"region"`
}

// CreateVpnGatewayRequestPublicConfig: create vpn gateway request public config.
type CreateVpnGatewayRequestPublicConfig struct {
	IpamIPv4ID *string `json:"ipam_ipv4_id"`

	IpamIPv6ID *string `json:"ipam_ipv6_id"`
}

// CustomerGateway: customer gateway.
type CustomerGateway struct {
	// ID: unique identifier of the customer gateway.
	ID string `json:"id"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// OrganizationID: organization ID.
	OrganizationID string `json:"organization_id"`

	// Name: name of the customer gateway.
	Name string `json:"name"`

	// Tags: list of tags applied to the customer gateway.
	Tags []string `json:"tags"`

	// CreatedAt: creation date of the customer gateway.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the customer gateway.
	UpdatedAt *time.Time `json:"updated_at"`

	// PublicIPv4: public IPv4 address of the customer gateway.
	PublicIPv4 *net.IP `json:"public_ipv4"`

	// PublicIPv6: public IPv6 address of the customer gateway.
	PublicIPv6 *net.IP `json:"public_ipv6"`

	// Asn: aS Number of the customer gateway.
	Asn uint32 `json:"asn"`
}

// RoutingPolicy: routing policy.
type RoutingPolicy struct {
	// ID: unique identifier of the routing policy.
	ID string `json:"id"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// OrganizationID: organization ID.
	OrganizationID string `json:"organization_id"`

	// Name: name of the routing policy.
	Name string `json:"name"`

	// Tags: list of tags associated with the routing policy.
	Tags []string `json:"tags"`

	// CreatedAt: creation date of the routing policy.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the routing policy.
	UpdatedAt *time.Time `json:"updated_at"`

	// IsIPv6: IP prefixes version of the routing policy.
	IsIPv6 bool `json:"is_ipv6"`

	// PrefixFilterIn: IP prefixes to accept from the peer (ranges of route announcements to accept).
	PrefixFilterIn []scw.IPNet `json:"prefix_filter_in"`

	// PrefixFilterOut: IP prefix filters to advertise to the peer (ranges of routes to advertise).
	PrefixFilterOut []scw.IPNet `json:"prefix_filter_out"`

	// Region: region of the routing policy.
	Region scw.Region `json:"region"`
}

// GatewayType: gateway type.
type GatewayType struct {
	Name string `json:"name"`

	Bandwidth uint64 `json:"bandwidth"`

	AllowedConnections uint64 `json:"allowed_connections"`

	Zones []string `json:"zones"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// VpnGateway: vpn gateway.
type VpnGateway struct {
	// ID: unique identifier of the VPN gateway.
	ID string `json:"id"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// OrganizationID: organization ID.
	OrganizationID string `json:"organization_id"`

	// Name: name of the VPN gateway.
	Name string `json:"name"`

	// Tags: list of tags applied to the VPN gateway.
	Tags []string `json:"tags"`

	// CreatedAt: creation date of the VPN gateway.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the VPN gateway.
	UpdatedAt *time.Time `json:"updated_at"`

	// Status: status of the VPN gateway.
	// Default value: unknown_status
	Status VpnGatewayStatus `json:"status"`

	// GatewayType: gateway type of the VPN gateway.
	GatewayType string `json:"gateway_type"`

	// PublicConfig: public endpoint configuration of the VPN gateway.
	// Precisely one of PublicConfig, PrivateConfig must be set.
	PublicConfig *VpnGatewayPublicConfig `json:"public_config,omitempty"`

	// PrivateNetworkID: ID of the Private Network attached to the VPN gateway.
	PrivateNetworkID string `json:"private_network_id"`

	// Precisely one of PublicConfig, PrivateConfig must be set.
	PrivateConfig *VpnGatewayPrivateConfig `json:"private_config,omitempty"`

	// IpamPrivateIPv4ID: ID of the IPAM private IPv4 address attached to the VPN gateway.
	IpamPrivateIPv4ID string `json:"ipam_private_ipv4_id"`

	// IpamPrivateIPv6ID: ID of the IPAM private IPv6 address attached to the VPN gateway.
	IpamPrivateIPv6ID string `json:"ipam_private_ipv6_id"`

	// Asn: autonomous System Number (ASN) of the VPN gateway, used by Border Gateway Protocol (BGP) to exchange routing information with the customer gateway.
	Asn uint32 `json:"asn"`

	// Zone: zone where the VPN gateway resource is currently provisioned.
	Zone scw.Zone `json:"zone"`

	// Region: region of the VPN gateway.
	Region scw.Region `json:"region"`
}

// CreateConnectionRequest: create connection request.
type CreateConnectionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project to create the connection in.
	ProjectID string `json:"project_id"`

	// Name: name of the connection.
	Name string `json:"name"`

	// Tags: list of tags to apply to the connection.
	Tags []string `json:"tags"`

	// IsIPv6: defines IP version of the IPSec Tunnel.
	IsIPv6 bool `json:"is_ipv6"`

	// InitiationPolicy: who initiates the IPsec tunnel.
	// Default value: unknown_initiation_policy
	InitiationPolicy CreateConnectionRequestInitiationPolicy `json:"initiation_policy"`

	// Ikev2Ciphers: list of IKE v2 ciphers proposed for the IPsec tunnel.
	Ikev2Ciphers []*ConnectionCipher `json:"ikev2_ciphers"`

	// EspCiphers: list of ESP ciphers proposed for the IPsec tunnel.
	EspCiphers []*ConnectionCipher `json:"esp_ciphers"`

	// EnableRoutePropagation: defines whether route propagation is enabled or not.
	EnableRoutePropagation bool `json:"enable_route_propagation"`

	// VpnGatewayID: ID of the VPN gateway to attach to the connection.
	VpnGatewayID string `json:"vpn_gateway_id"`

	// CustomerGatewayID: ID of the customer gateway to attach to the connection.
	CustomerGatewayID string `json:"customer_gateway_id"`

	// BgpConfigIPv4: bGP config of IPv4 session, including interco private IPv4 subnet (first IP assigned to the VPN Gateway, second IP to the Customer Gateway) and attached routing policy.
	BgpConfigIPv4 *CreateConnectionRequestBgpConfig `json:"bgp_config_ipv4,omitempty"`

	// BgpConfigIPv6: bGP config of IPv6 session, including interco private IPv6 subnet (first IP assigned to the VPN Gateway, second IP to the Customer Gateway) and attached routing policy.
	BgpConfigIPv6 *CreateConnectionRequestBgpConfig `json:"bgp_config_ipv6,omitempty"`
}

// CreateConnectionResponse: create connection response.
type CreateConnectionResponse struct {
	// Connection: this connection.
	Connection *Connection `json:"connection"`

	// Deprecated: PreSharedKey: deprecated, use secret_id & secret_revision fields.
	PreSharedKey *string `json:"pre_shared_key,omitempty"`
}

// CreateCustomerGatewayRequest: create customer gateway request.
type CreateCustomerGatewayRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project to create the customer gateway in.
	ProjectID string `json:"project_id"`

	// Name: name of the customer gateway.
	Name string `json:"name"`

	// Tags: list of tags to apply to the customer gateway.
	Tags []string `json:"tags"`

	// IPv4Public: public IPv4 address of the customer gateway.
	IPv4Public *scw.IPNet `json:"ipv4_public,omitempty"`

	// IPv6Public: public IPv6 address of the customer gateway.
	IPv6Public *scw.IPNet `json:"ipv6_public,omitempty"`

	// Asn: aS Number of the customer gateway.
	Asn uint32 `json:"asn"`
}

// CreateRoutingPolicyRequest: create routing policy request.
type CreateRoutingPolicyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project to create the routing policy in.
	ProjectID string `json:"project_id"`

	// Name: name of the routing policy.
	Name string `json:"name"`

	// Tags: list of tags to apply to the routing policy.
	Tags []string `json:"tags"`

	// IsIPv6: IP prefixes version of the routing policy.
	IsIPv6 bool `json:"is_ipv6"`

	// PrefixFilterIn: IP prefixes to accept from the peer (ranges of route announcements to accept).
	PrefixFilterIn []scw.IPNet `json:"prefix_filter_in"`

	// PrefixFilterOut: IP prefix filters to advertise to the peer (ranges of routes to advertise).
	PrefixFilterOut []scw.IPNet `json:"prefix_filter_out"`
}

// CreateVpnGatewayRequest: create vpn gateway request.
type CreateVpnGatewayRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project to create the VPN gateway in.
	ProjectID string `json:"project_id"`

	// Name: name of the VPN gateway.
	Name string `json:"name"`

	// Tags: list of tags to apply to the VPN gateway.
	Tags []string `json:"tags"`

	// GatewayType: vPN gateway type (commercial offer type).
	GatewayType string `json:"gateway_type"`

	// PublicConfig: public endpoint configuration of the VPN gateway.
	// Precisely one of PublicConfig must be set.
	PublicConfig *CreateVpnGatewayRequestPublicConfig `json:"public_config,omitempty"`

	// PrivateNetworkID: ID of the Private Network to attach to the VPN gateway.
	PrivateNetworkID string `json:"private_network_id"`

	// IpamPrivateIPv4ID: ID of the IPAM private IPv4 address to attach to the VPN gateway.
	IpamPrivateIPv4ID *string `json:"ipam_private_ipv4_id,omitempty"`

	// IpamPrivateIPv6ID: ID of the IPAM private IPv6 address to attach to the VPN gateway.
	IpamPrivateIPv6ID *string `json:"ipam_private_ipv6_id,omitempty"`

	// Zone: availability Zone where the VPN gateway should be provisioned. If no zone is specified, the VPN gateway will be automatically placed.
	Zone *scw.Zone `json:"zone,omitempty"`
}

// DeleteConnectionRequest: delete connection request.
type DeleteConnectionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ConnectionID: ID of the connection to delete.
	ConnectionID string `json:"-"`
}

// DeleteCustomerGatewayRequest: delete customer gateway request.
type DeleteCustomerGatewayRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// GatewayID: ID of the customer gateway to delete.
	GatewayID string `json:"-"`
}

// DeleteRoutingPolicyRequest: delete routing policy request.
type DeleteRoutingPolicyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// RoutingPolicyID: ID of the routing policy to delete.
	RoutingPolicyID string `json:"-"`
}

// DeleteVpnGatewayRequest: delete vpn gateway request.
type DeleteVpnGatewayRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// GatewayID: ID of the VPN gateway to delete.
	GatewayID string `json:"-"`
}

// DetachRoutingPolicyRequest: detach routing policy request.
type DetachRoutingPolicyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ConnectionID: ID of the connection from which routing policy is being detached.
	ConnectionID string `json:"-"`

	// RoutingPolicyV4: ID of the routing policy to detach from the BGP IPv4 session.
	// Precisely one of RoutingPolicyV4, RoutingPolicyV6 must be set.
	RoutingPolicyV4 *string `json:"routing_policy_v4,omitempty"`

	// RoutingPolicyV6: ID of the routing policy to detach from the BGP IPv6 session.
	// Precisely one of RoutingPolicyV4, RoutingPolicyV6 must be set.
	RoutingPolicyV6 *string `json:"routing_policy_v6,omitempty"`
}

// DisableRoutePropagationRequest: disable route propagation request.
type DisableRoutePropagationRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ConnectionID: ID of the connection on which to disable route propagation.
	ConnectionID string `json:"-"`
}

// EnableRoutePropagationRequest: enable route propagation request.
type EnableRoutePropagationRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ConnectionID: ID of the connection on which to enable route propagation.
	ConnectionID string `json:"-"`
}

// GetConnectionRequest: get connection request.
type GetConnectionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ConnectionID: ID of the requested connection.
	ConnectionID string `json:"-"`
}

// GetCustomerGatewayRequest: get customer gateway request.
type GetCustomerGatewayRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// GatewayID: ID of the requested customer gateway.
	GatewayID string `json:"-"`
}

// GetRoutingPolicyRequest: get routing policy request.
type GetRoutingPolicyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// RoutingPolicyID: ID of the routing policy to get.
	RoutingPolicyID string `json:"-"`
}

// GetVpnGatewayRequest: get vpn gateway request.
type GetVpnGatewayRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// GatewayID: ID of the requested VPN gateway.
	GatewayID string `json:"-"`
}

// ListConnectionsRequest: list connections request.
type ListConnectionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of connections to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListConnectionsRequestOrderBy `json:"-"`

	// ProjectID: project ID to filter for.
	ProjectID *string `json:"-"`

	// Name: connection name to filter for.
	Name *string `json:"-"`

	// Tags: tags to filter for.
	Tags []string `json:"-"`

	// Statuses: connection statuses to filter for.
	Statuses []ConnectionStatus `json:"-"`

	// IsIPv6: filter connections with IP version of IPSec tunnel.
	IsIPv6 *bool `json:"-"`

	// RoutingPolicyIDs: filter for connections using these routing policies.
	RoutingPolicyIDs []string `json:"-"`

	// RoutePropagationEnabled: filter for connections with route propagation enabled.
	RoutePropagationEnabled *bool `json:"-"`

	// VpnGatewayIDs: filter for connections attached to these VPN gateways.
	VpnGatewayIDs []string `json:"-"`

	// CustomerGatewayIDs: filter for connections attached to these customer gateways.
	CustomerGatewayIDs []string `json:"-"`
}

// ListConnectionsResponse: list connections response.
type ListConnectionsResponse struct {
	// Connections: list of connections on the current page.
	Connections []*Connection `json:"connections"`

	// TotalCount: total number of connections.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListConnectionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListConnectionsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListConnectionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Connections = append(r.Connections, results.Connections...)
	r.TotalCount += uint64(len(results.Connections))
	return uint64(len(results.Connections)), nil
}

// ListCustomerGatewaysRequest: list customer gateways request.
type ListCustomerGatewaysRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of customer gateways to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListCustomerGatewaysRequestOrderBy `json:"-"`

	// ProjectID: project ID to filter for.
	ProjectID *string `json:"-"`

	// Name: customer gateway name to filter for.
	Name *string `json:"-"`

	// Tags: tags to filter for.
	Tags []string `json:"-"`
}

// ListCustomerGatewaysResponse: list customer gateways response.
type ListCustomerGatewaysResponse struct {
	// Gateways: list of customer gateways on the current page.
	Gateways []*CustomerGateway `json:"gateways"`

	// TotalCount: total number of customer gateways.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCustomerGatewaysResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCustomerGatewaysResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListCustomerGatewaysResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Gateways = append(r.Gateways, results.Gateways...)
	r.TotalCount += uint64(len(results.Gateways))
	return uint64(len(results.Gateways)), nil
}

// ListRoutingPoliciesRequest: list routing policies request.
type ListRoutingPoliciesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of routing policies to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListRoutingPoliciesRequestOrderBy `json:"-"`

	// ProjectID: project ID to filter for.
	ProjectID *string `json:"-"`

	// Name: routing policy name to filter for.
	Name *string `json:"-"`

	// Tags: tags to filter for.
	Tags []string `json:"-"`

	// IPv6: filter for the routing policies based on IP prefixes version.
	IPv6 *bool `json:"-"`
}

// ListRoutingPoliciesResponse: list routing policies response.
type ListRoutingPoliciesResponse struct {
	RoutingPolicies []*RoutingPolicy `json:"routing_policies"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRoutingPoliciesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRoutingPoliciesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListRoutingPoliciesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.RoutingPolicies = append(r.RoutingPolicies, results.RoutingPolicies...)
	r.TotalCount += uint64(len(results.RoutingPolicies))
	return uint64(len(results.RoutingPolicies)), nil
}

// ListVpnGatewayTypesRequest: list vpn gateway types request.
type ListVpnGatewayTypesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of VPN gateway types to return per page.
	PageSize *uint32 `json:"-"`
}

// ListVpnGatewayTypesResponse: list vpn gateway types response.
type ListVpnGatewayTypesResponse struct {
	// GatewayTypes: list of VPN gateway types on the current page.
	GatewayTypes []*GatewayType `json:"gateway_types"`

	// TotalCount: total number of gateway types.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVpnGatewayTypesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVpnGatewayTypesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListVpnGatewayTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.GatewayTypes = append(r.GatewayTypes, results.GatewayTypes...)
	r.TotalCount += uint64(len(results.GatewayTypes))
	return uint64(len(results.GatewayTypes)), nil
}

// ListVpnGatewaysRequest: list vpn gateways request.
type ListVpnGatewaysRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of VPN gateways to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListVpnGatewaysRequestOrderBy `json:"-"`

	// ProjectID: project ID to filter for.
	ProjectID *string `json:"-"`

	// Name: vPN gateway name to filter for.
	Name *string `json:"-"`

	// Tags: tags to filter for.
	Tags []string `json:"-"`

	// Statuses: vPN gateway statuses to filter for.
	Statuses []VpnGatewayStatus `json:"-"`

	// GatewayTypes: filter for VPN gateways of these types.
	GatewayTypes []string `json:"-"`

	// PrivateNetworkIDs: filter for VPN gateways attached to these private networks.
	PrivateNetworkIDs []string `json:"-"`
}

// ListVpnGatewaysResponse: list vpn gateways response.
type ListVpnGatewaysResponse struct {
	// Gateways: list of VPN gateways on the current page.
	Gateways []*VpnGateway `json:"gateways"`

	// TotalCount: total number of VPN gateways.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVpnGatewaysResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVpnGatewaysResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListVpnGatewaysResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Gateways = append(r.Gateways, results.Gateways...)
	r.TotalCount += uint64(len(results.Gateways))
	return uint64(len(results.Gateways)), nil
}

// RenewConnectionPskRequest: renew connection psk request.
type RenewConnectionPskRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ConnectionID: ID of the connection to renew the PSK.
	ConnectionID string `json:"-"`
}

// RenewConnectionPskResponse: renew connection psk response.
type RenewConnectionPskResponse struct {
	// Connection: this connection.
	Connection *Connection `json:"connection"`

	// Deprecated: PreSharedKey: deprecated, use secret_id & secret_revision fields.
	PreSharedKey *string `json:"pre_shared_key,omitempty"`
}

// SetRoutingPolicyRequest: set routing policy request.
type SetRoutingPolicyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ConnectionID: ID of the connection whose routing policy is being updated.
	ConnectionID string `json:"-"`

	// RoutingPolicyV4: ID of the routing policy to set for the BGP IPv4 session.
	// Precisely one of RoutingPolicyV4, RoutingPolicyV6 must be set.
	RoutingPolicyV4 *string `json:"routing_policy_v4,omitempty"`

	// RoutingPolicyV6: ID of the routing policy to set for the BGP IPv6 session.
	// Precisely one of RoutingPolicyV4, RoutingPolicyV6 must be set.
	RoutingPolicyV6 *string `json:"routing_policy_v6,omitempty"`
}

// UpdateConnectionRequest: update connection request.
type UpdateConnectionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ConnectionID: ID of the connection to update.
	ConnectionID string `json:"-"`

	// Name: name of the connection.
	Name *string `json:"name,omitempty"`

	// Tags: list of tags to apply to the connection.
	Tags *[]string `json:"tags,omitempty"`

	// InitiationPolicy: who initiates the IPsec tunnel.
	// Default value: unknown_initiation_policy
	InitiationPolicy CreateConnectionRequestInitiationPolicy `json:"initiation_policy"`

	// Ikev2Ciphers: list of IKE v2 ciphers proposed for the IPsec tunnel.
	Ikev2Ciphers []*ConnectionCipher `json:"ikev2_ciphers"`

	// EspCiphers: list of ESP ciphers proposed for the IPsec tunnel.
	EspCiphers []*ConnectionCipher `json:"esp_ciphers"`
}

// UpdateCustomerGatewayRequest: update customer gateway request.
type UpdateCustomerGatewayRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// GatewayID: ID of the customer gateway to update.
	GatewayID string `json:"-"`

	// Name: name of the customer gateway.
	Name *string `json:"name,omitempty"`

	// Tags: list of tags to apply to the customer gateway.
	Tags *[]string `json:"tags,omitempty"`

	// IPv4Public: public IPv4 address of the customer gateway.
	IPv4Public *scw.IPNet `json:"ipv4_public,omitempty"`

	// IPv6Public: public IPv6 address of the customer gateway.
	IPv6Public *scw.IPNet `json:"ipv6_public,omitempty"`

	// Asn: aS Number of the customer gateway.
	Asn *uint32 `json:"asn,omitempty"`
}

// UpdateRoutingPolicyRequest: update routing policy request.
type UpdateRoutingPolicyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// RoutingPolicyID: ID of the routing policy to update.
	RoutingPolicyID string `json:"-"`

	// Name: name of the routing policy.
	Name *string `json:"name,omitempty"`

	// Tags: list of tags to apply to the routing policy.
	Tags *[]string `json:"tags,omitempty"`

	// PrefixFilterIn: IP prefixes to accept from the peer (ranges of route announcements to accept).
	PrefixFilterIn *[]string `json:"prefix_filter_in,omitempty"`

	// PrefixFilterOut: IP prefix filters for routes to advertise to the peer (ranges of routes to advertise).
	PrefixFilterOut *[]string `json:"prefix_filter_out,omitempty"`
}

// UpdateVpnGatewayRequest: update vpn gateway request.
type UpdateVpnGatewayRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// GatewayID: ID of the VPN gateway to update.
	GatewayID string `json:"-"`

	// Name: name of the VPN gateway.
	Name *string `json:"name,omitempty"`

	// Tags: list of tags to apply to the VPN Gateway.
	Tags *[]string `json:"tags,omitempty"`
}

// This API allows you to manage your Site-to-Site VPN.
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

// ListVpnGatewayTypes: List the different VPN gateway commercial offer types available at Scaleway. The response is an array of objects describing the name and technical details of each available VPN gateway type.
func (s *API) ListVpnGatewayTypes(req *ListVpnGatewayTypesRequest, opts ...scw.RequestOption) (*ListVpnGatewayTypesResponse, error) {
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
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/vpn-gateway-types",
		Query:  query,
	}

	var resp ListVpnGatewayTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListVpnGateways: List all your VPN gateways. A number of filters are available, including Project ID, name, tags and status.
func (s *API) ListVpnGateways(req *ListVpnGatewaysRequest, opts ...scw.RequestOption) (*ListVpnGatewaysResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "gateway_types", req.GatewayTypes)
	parameter.AddToQuery(query, "private_network_ids", req.PrivateNetworkIDs)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/vpn-gateways",
		Query:  query,
	}

	var resp ListVpnGatewaysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVpnGateway: Get a VPN gateway for the given VPN gateway ID.
func (s *API) GetVpnGateway(req *GetVpnGatewayRequest, opts ...scw.RequestOption) (*VpnGateway, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayID) == "" {
		return nil, errors.New("field GatewayID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/vpn-gateways/" + fmt.Sprint(req.GatewayID) + "",
	}

	var resp VpnGateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForVpnGatewayRequest is used by WaitForVpnGateway method.
type WaitForVpnGatewayRequest struct {
	Region        scw.Region
	GatewayID     string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForVpnGateway waits for the VpnGateway to reach a terminal state.
func (s *API) WaitForVpnGateway(req *WaitForVpnGatewayRequest, opts ...scw.RequestOption) (*VpnGateway, error) {
	timeout := defaultS2sVpnTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultS2sVpnRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[VpnGatewayStatus]struct{}{
		VpnGatewayStatusConfiguring:    {},
		VpnGatewayStatusProvisioning:   {},
		VpnGatewayStatusDeprovisioning: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetVpnGateway(&GetVpnGatewayRequest{
				Region:    req.Region,
				GatewayID: req.GatewayID,
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
		return nil, errors.Wrap(err, "waiting for VpnGateway failed")
	}

	return res.(*VpnGateway), nil
}

// CreateVpnGateway: Create VPN gateway.
func (s *API) CreateVpnGateway(req *CreateVpnGatewayRequest, opts ...scw.RequestOption) (*VpnGateway, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	defaultZone, exist := s.client.GetDefaultZone()
	if (req.Zone == nil || *req.Zone == "") && exist {
		req.Zone = &defaultZone
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/vpn-gateways",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp VpnGateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVpnGateway: Update an existing VPN gateway, specified by its VPN gateway ID. Only its name and tags can be updated.
func (s *API) UpdateVpnGateway(req *UpdateVpnGatewayRequest, opts ...scw.RequestOption) (*VpnGateway, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayID) == "" {
		return nil, errors.New("field GatewayID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/vpn-gateways/" + fmt.Sprint(req.GatewayID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp VpnGateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVpnGateway: Delete an existing VPN gateway, specified by its VPN gateway ID.
func (s *API) DeleteVpnGateway(req *DeleteVpnGatewayRequest, opts ...scw.RequestOption) (*VpnGateway, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayID) == "" {
		return nil, errors.New("field GatewayID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/vpn-gateways/" + fmt.Sprint(req.GatewayID) + "",
	}

	var resp VpnGateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListConnections: List all your connections. A number of filters are available, including Project ID, name, tags and status.
func (s *API) ListConnections(req *ListConnectionsRequest, opts ...scw.RequestOption) (*ListConnectionsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "is_ipv6", req.IsIPv6)
	parameter.AddToQuery(query, "routing_policy_ids", req.RoutingPolicyIDs)
	parameter.AddToQuery(query, "route_propagation_enabled", req.RoutePropagationEnabled)
	parameter.AddToQuery(query, "vpn_gateway_ids", req.VpnGatewayIDs)
	parameter.AddToQuery(query, "customer_gateway_ids", req.CustomerGatewayIDs)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/connections",
		Query:  query,
	}

	var resp ListConnectionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetConnection: Get a connection for the given connection ID. The response object includes information about the connection's various configuration details.
func (s *API) GetConnection(req *GetConnectionRequest, opts ...scw.RequestOption) (*Connection, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ConnectionID) == "" {
		return nil, errors.New("field ConnectionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/connections/" + fmt.Sprint(req.ConnectionID) + "",
	}

	var resp Connection

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateConnection: Create a connection.
func (s *API) CreateConnection(req *CreateConnectionRequest, opts ...scw.RequestOption) (*CreateConnectionResponse, error) {
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
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/connections",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateConnectionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateConnection: Update an existing connection, specified by its connection ID.
func (s *API) UpdateConnection(req *UpdateConnectionRequest, opts ...scw.RequestOption) (*Connection, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ConnectionID) == "" {
		return nil, errors.New("field ConnectionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/connections/" + fmt.Sprint(req.ConnectionID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Connection

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteConnection: Delete an existing connection, specified by its connection ID.
func (s *API) DeleteConnection(req *DeleteConnectionRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ConnectionID) == "" {
		return errors.New("field ConnectionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/connections/" + fmt.Sprint(req.ConnectionID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// RenewConnectionPsk: Renew pre-shared key for a given connection.
func (s *API) RenewConnectionPsk(req *RenewConnectionPskRequest, opts ...scw.RequestOption) (*RenewConnectionPskResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ConnectionID) == "" {
		return nil, errors.New("field ConnectionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/connections/" + fmt.Sprint(req.ConnectionID) + "/renew-psk",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RenewConnectionPskResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetRoutingPolicy: Set a new routing policy on a connection, overriding the existing one if present, specified by its connection ID.
func (s *API) SetRoutingPolicy(req *SetRoutingPolicyRequest, opts ...scw.RequestOption) (*Connection, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ConnectionID) == "" {
		return nil, errors.New("field ConnectionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/connections/" + fmt.Sprint(req.ConnectionID) + "/set-routing-policy",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Connection

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachRoutingPolicy: Detach an existing routing policy from a connection, specified by its connection ID.
func (s *API) DetachRoutingPolicy(req *DetachRoutingPolicyRequest, opts ...scw.RequestOption) (*Connection, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ConnectionID) == "" {
		return nil, errors.New("field ConnectionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/connections/" + fmt.Sprint(req.ConnectionID) + "/detach-routing-policy",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Connection

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableRoutePropagation: Enable all allowed prefixes (defined in a routing policy) to be announced in the BGP session. This allows traffic to flow between the attached VPC and the on-premises infrastructure along the announced routes. Note that by default, even when route propagation is enabled, all routes are blocked. It is essential to attach a routing policy to define the ranges of routes to announce.
func (s *API) EnableRoutePropagation(req *EnableRoutePropagationRequest, opts ...scw.RequestOption) (*Connection, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ConnectionID) == "" {
		return nil, errors.New("field ConnectionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/connections/" + fmt.Sprint(req.ConnectionID) + "/enable-route-propagation",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Connection

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableRoutePropagation: Prevent any prefixes from being announced in the BGP session. Traffic will not be able to flow over the VPN Gateway until route propagation is re-enabled.
func (s *API) DisableRoutePropagation(req *DisableRoutePropagationRequest, opts ...scw.RequestOption) (*Connection, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ConnectionID) == "" {
		return nil, errors.New("field ConnectionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/connections/" + fmt.Sprint(req.ConnectionID) + "/disable-route-propagation",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Connection

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCustomerGateways: List all your customer gateways. A number of filters are available, including Project ID, name, and tags.
func (s *API) ListCustomerGateways(req *ListCustomerGatewaysRequest, opts ...scw.RequestOption) (*ListCustomerGatewaysResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/customer-gateways",
		Query:  query,
	}

	var resp ListCustomerGatewaysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCustomerGateway: Get a customer gateway for the given customer gateway ID.
func (s *API) GetCustomerGateway(req *GetCustomerGatewayRequest, opts ...scw.RequestOption) (*CustomerGateway, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayID) == "" {
		return nil, errors.New("field GatewayID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/customer-gateways/" + fmt.Sprint(req.GatewayID) + "",
	}

	var resp CustomerGateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCustomerGateway: Create a customer gateway.
func (s *API) CreateCustomerGateway(req *CreateCustomerGatewayRequest, opts ...scw.RequestOption) (*CustomerGateway, error) {
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
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/customer-gateways",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CustomerGateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCustomerGateway: Update an existing customer gateway, specified by its customer gateway ID. You can update its name, tags, public IPv4 & IPv6 address and AS Number.
func (s *API) UpdateCustomerGateway(req *UpdateCustomerGatewayRequest, opts ...scw.RequestOption) (*CustomerGateway, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayID) == "" {
		return nil, errors.New("field GatewayID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/customer-gateways/" + fmt.Sprint(req.GatewayID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CustomerGateway

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCustomerGateway: Delete an existing customer gateway, specified by its customer gateway ID.
func (s *API) DeleteCustomerGateway(req *DeleteCustomerGatewayRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.GatewayID) == "" {
		return errors.New("field GatewayID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/customer-gateways/" + fmt.Sprint(req.GatewayID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListRoutingPolicies: List all routing policies in a given region. A routing policy can be attached to one or multiple connections (S2S VPN connections).
func (s *API) ListRoutingPolicies(req *ListRoutingPoliciesRequest, opts ...scw.RequestOption) (*ListRoutingPoliciesResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "ipv6", req.IPv6)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/routing-policies",
		Query:  query,
	}

	var resp ListRoutingPoliciesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRoutingPolicy: Get a routing policy for the given routing policy ID. The response object gives information including the policy's name, tags and prefix filters.
func (s *API) GetRoutingPolicy(req *GetRoutingPolicyRequest, opts ...scw.RequestOption) (*RoutingPolicy, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.RoutingPolicyID) == "" {
		return nil, errors.New("field RoutingPolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/routing-policies/" + fmt.Sprint(req.RoutingPolicyID) + "",
	}

	var resp RoutingPolicy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRoutingPolicy: Create a routing policy. Routing policies allow you to set IP prefix filters to define the incoming route announcements to accept from the customer gateway, and the outgoing routes to announce to the customer gateway.
func (s *API) CreateRoutingPolicy(req *CreateRoutingPolicyRequest, opts ...scw.RequestOption) (*RoutingPolicy, error) {
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
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/routing-policies",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RoutingPolicy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateRoutingPolicy: Update an existing routing policy, specified by its routing policy ID. Its name, tags and incoming/outgoing prefix filters can be updated.
func (s *API) UpdateRoutingPolicy(req *UpdateRoutingPolicyRequest, opts ...scw.RequestOption) (*RoutingPolicy, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.RoutingPolicyID) == "" {
		return nil, errors.New("field RoutingPolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/routing-policies/" + fmt.Sprint(req.RoutingPolicyID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RoutingPolicy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteRoutingPolicy: Delete an existing routing policy, specified by its routing policy ID.
func (s *API) DeleteRoutingPolicy(req *DeleteRoutingPolicyRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.RoutingPolicyID) == "" {
		return errors.New("field RoutingPolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/s2s-vpn/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/routing-policies/" + fmt.Sprint(req.RoutingPolicyID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

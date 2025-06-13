// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package interlink provides methods and message types of the interlink v1beta1 API.
package interlink

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

type BgpStatus string

const (
	BgpStatusUnknownBgpStatus = BgpStatus("unknown_bgp_status")
	BgpStatusUp               = BgpStatus("up")
	BgpStatusDown             = BgpStatus("down")
)

func (enum BgpStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(BgpStatusUnknownBgpStatus)
	}
	return string(enum)
}

func (enum BgpStatus) Values() []BgpStatus {
	return []BgpStatus{
		"unknown_bgp_status",
		"up",
		"down",
	}
}

func (enum BgpStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BgpStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BgpStatus(BgpStatus(tmp).String())
	return nil
}

type DedicatedConnectionStatus string

const (
	DedicatedConnectionStatusUnknownStatus = DedicatedConnectionStatus("unknown_status")
	DedicatedConnectionStatusCreated       = DedicatedConnectionStatus("created")
	DedicatedConnectionStatusConfiguring   = DedicatedConnectionStatus("configuring")
	DedicatedConnectionStatusFailed        = DedicatedConnectionStatus("failed")
	DedicatedConnectionStatusActive        = DedicatedConnectionStatus("active")
	DedicatedConnectionStatusDisabled      = DedicatedConnectionStatus("disabled")
	DedicatedConnectionStatusDeleted       = DedicatedConnectionStatus("deleted")
	DedicatedConnectionStatusLocked        = DedicatedConnectionStatus("locked")
)

func (enum DedicatedConnectionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(DedicatedConnectionStatusUnknownStatus)
	}
	return string(enum)
}

func (enum DedicatedConnectionStatus) Values() []DedicatedConnectionStatus {
	return []DedicatedConnectionStatus{
		"unknown_status",
		"created",
		"configuring",
		"failed",
		"active",
		"disabled",
		"deleted",
		"locked",
	}
}

func (enum DedicatedConnectionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DedicatedConnectionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DedicatedConnectionStatus(DedicatedConnectionStatus(tmp).String())
	return nil
}

type LinkKind string

const (
	LinkKindHosted     = LinkKind("hosted")
	LinkKindSelfHosted = LinkKind("self_hosted")
)

func (enum LinkKind) String() string {
	if enum == "" {
		// return default value if empty
		return string(LinkKindHosted)
	}
	return string(enum)
}

func (enum LinkKind) Values() []LinkKind {
	return []LinkKind{
		"hosted",
		"self_hosted",
	}
}

func (enum LinkKind) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LinkKind) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LinkKind(LinkKind(tmp).String())
	return nil
}

type LinkStatus string

const (
	LinkStatusUnknownLinkStatus   = LinkStatus("unknown_link_status")
	LinkStatusConfiguring         = LinkStatus("configuring")
	LinkStatusFailed              = LinkStatus("failed")
	LinkStatusRequested           = LinkStatus("requested")
	LinkStatusRefused             = LinkStatus("refused")
	LinkStatusExpired             = LinkStatus("expired")
	LinkStatusProvisioning        = LinkStatus("provisioning")
	LinkStatusActive              = LinkStatus("active")
	LinkStatusLimitedConnectivity = LinkStatus("limited_connectivity")
	LinkStatusAllDown             = LinkStatus("all_down")
	LinkStatusDeprovisioning      = LinkStatus("deprovisioning")
	LinkStatusDeleted             = LinkStatus("deleted")
	LinkStatusLocked              = LinkStatus("locked")
)

func (enum LinkStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(LinkStatusUnknownLinkStatus)
	}
	return string(enum)
}

func (enum LinkStatus) Values() []LinkStatus {
	return []LinkStatus{
		"unknown_link_status",
		"configuring",
		"failed",
		"requested",
		"refused",
		"expired",
		"provisioning",
		"active",
		"limited_connectivity",
		"all_down",
		"deprovisioning",
		"deleted",
		"locked",
	}
}

func (enum LinkStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LinkStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LinkStatus(LinkStatus(tmp).String())
	return nil
}

type ListDedicatedConnectionsRequestOrderBy string

const (
	ListDedicatedConnectionsRequestOrderByCreatedAtAsc  = ListDedicatedConnectionsRequestOrderBy("created_at_asc")
	ListDedicatedConnectionsRequestOrderByCreatedAtDesc = ListDedicatedConnectionsRequestOrderBy("created_at_desc")
	ListDedicatedConnectionsRequestOrderByUpdatedAtAsc  = ListDedicatedConnectionsRequestOrderBy("updated_at_asc")
	ListDedicatedConnectionsRequestOrderByUpdatedAtDesc = ListDedicatedConnectionsRequestOrderBy("updated_at_desc")
	ListDedicatedConnectionsRequestOrderByNameAsc       = ListDedicatedConnectionsRequestOrderBy("name_asc")
	ListDedicatedConnectionsRequestOrderByNameDesc      = ListDedicatedConnectionsRequestOrderBy("name_desc")
	ListDedicatedConnectionsRequestOrderByStatusAsc     = ListDedicatedConnectionsRequestOrderBy("status_asc")
	ListDedicatedConnectionsRequestOrderByStatusDesc    = ListDedicatedConnectionsRequestOrderBy("status_desc")
)

func (enum ListDedicatedConnectionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListDedicatedConnectionsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListDedicatedConnectionsRequestOrderBy) Values() []ListDedicatedConnectionsRequestOrderBy {
	return []ListDedicatedConnectionsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
		"name_asc",
		"name_desc",
		"status_asc",
		"status_desc",
	}
}

func (enum ListDedicatedConnectionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDedicatedConnectionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDedicatedConnectionsRequestOrderBy(ListDedicatedConnectionsRequestOrderBy(tmp).String())
	return nil
}

type ListLinksRequestOrderBy string

const (
	ListLinksRequestOrderByCreatedAtAsc  = ListLinksRequestOrderBy("created_at_asc")
	ListLinksRequestOrderByCreatedAtDesc = ListLinksRequestOrderBy("created_at_desc")
	ListLinksRequestOrderByNameAsc       = ListLinksRequestOrderBy("name_asc")
	ListLinksRequestOrderByNameDesc      = ListLinksRequestOrderBy("name_desc")
	ListLinksRequestOrderByStatusAsc     = ListLinksRequestOrderBy("status_asc")
	ListLinksRequestOrderByStatusDesc    = ListLinksRequestOrderBy("status_desc")
)

func (enum ListLinksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListLinksRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListLinksRequestOrderBy) Values() []ListLinksRequestOrderBy {
	return []ListLinksRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
		"status_asc",
		"status_desc",
	}
}

func (enum ListLinksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListLinksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListLinksRequestOrderBy(ListLinksRequestOrderBy(tmp).String())
	return nil
}

type ListPartnersRequestOrderBy string

const (
	ListPartnersRequestOrderByNameAsc  = ListPartnersRequestOrderBy("name_asc")
	ListPartnersRequestOrderByNameDesc = ListPartnersRequestOrderBy("name_desc")
)

func (enum ListPartnersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPartnersRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListPartnersRequestOrderBy) Values() []ListPartnersRequestOrderBy {
	return []ListPartnersRequestOrderBy{
		"name_asc",
		"name_desc",
	}
}

func (enum ListPartnersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPartnersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPartnersRequestOrderBy(ListPartnersRequestOrderBy(tmp).String())
	return nil
}

type ListPopsRequestOrderBy string

const (
	ListPopsRequestOrderByNameAsc  = ListPopsRequestOrderBy("name_asc")
	ListPopsRequestOrderByNameDesc = ListPopsRequestOrderBy("name_desc")
)

func (enum ListPopsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPopsRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListPopsRequestOrderBy) Values() []ListPopsRequestOrderBy {
	return []ListPopsRequestOrderBy{
		"name_asc",
		"name_desc",
	}
}

func (enum ListPopsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPopsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPopsRequestOrderBy(ListPopsRequestOrderBy(tmp).String())
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

// BgpConfig: bgp config.
type BgpConfig struct {
	// Asn: aS Number of the BGP peer.
	Asn uint32 `json:"asn"`

	// IPv4: iPv4 address of the BGP peer.
	IPv4 scw.IPNet `json:"ipv4"`

	// IPv6: iPv6 address of the BGP peer.
	IPv6 scw.IPNet `json:"ipv6"`
}

// PartnerHost: partner host.
type PartnerHost struct {
	// PartnerID: ID of the partner facilitating the link.
	PartnerID string `json:"partner_id"`

	// PairingKey: used to identify a link from a user or partner's point of view.
	PairingKey string `json:"pairing_key"`

	// DisapprovedReason: reason given by partner to explain why they did not approve the request for a hosted link.
	DisapprovedReason *string `json:"disapproved_reason"`
}

// SelfHost: self host.
type SelfHost struct {
	// ConnectionID: dedicated physical connection supporting the link.
	ConnectionID string `json:"connection_id"`
}

// DedicatedConnection: dedicated connection.
type DedicatedConnection struct {
	// ID: unique identifier of the dedicated connection.
	ID string `json:"id"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// OrganizationID: organization ID.
	OrganizationID string `json:"organization_id"`

	// Status: status of the dedicated connection.
	// Default value: unknown_status
	Status DedicatedConnectionStatus `json:"status"`

	// Name: name of the dedicated connection.
	Name string `json:"name"`

	// Tags: list of tags associated with the dedicated connection.
	Tags []string `json:"tags"`

	// PopID: ID of the PoP where the dedicated connection is located.
	PopID string `json:"pop_id"`

	// BandwidthMbps: bandwidth size of the dedicated connection.
	BandwidthMbps uint64 `json:"bandwidth_mbps"`

	// AvailableLinkBandwidths: size of the links supported on this dedicated connection.
	AvailableLinkBandwidths []uint64 `json:"available_link_bandwidths"`

	// CreatedAt: creation date of the dedicated connection.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the dedicated connection.
	UpdatedAt *time.Time `json:"updated_at"`

	// DemarcationInfo: demarcation details required by the data center to set up the supporting Cross Connect. This generally includes the physical space in the facility, the cabinet or rack the connection should land in, the patch panel to go in, the port designation, and the media type.
	DemarcationInfo *string `json:"demarcation_info"`

	// Region: region of the dedicated connection.
	Region scw.Region `json:"region"`
}

// Link: link.
type Link struct {
	// ID: unique identifier of the link.
	ID string `json:"id"`

	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// OrganizationID: organization ID.
	OrganizationID string `json:"organization_id"`

	// Name: name of the link.
	Name string `json:"name"`

	// Tags: list of tags associated with the link.
	Tags []string `json:"tags"`

	// PopID: ID of the PoP where the link's corresponding connection is located.
	PopID string `json:"pop_id"`

	// BandwidthMbps: rate limited bandwidth of the link.
	BandwidthMbps uint64 `json:"bandwidth_mbps"`

	// Status: status of the link.
	// Default value: unknown_link_status
	Status LinkStatus `json:"status"`

	// BgpV4Status: status of the link's BGP IPv4 session.
	// Default value: unknown_bgp_status
	BgpV4Status BgpStatus `json:"bgp_v4_status"`

	// BgpV6Status: status of the link's BGP IPv6 session.
	// Default value: unknown_bgp_status
	BgpV6Status BgpStatus `json:"bgp_v6_status"`

	// VpcID: ID of the Scaleway VPC attached to the link.
	VpcID *string `json:"vpc_id"`

	// Deprecated: RoutingPolicyID: deprecated. Use routing_policy_v4_id or routing_policy_v6_id instead.
	RoutingPolicyID *string `json:"routing_policy_id,omitempty"`

	// EnableRoutePropagation: defines whether route propagation is enabled or not. To enable or disable route propagation, use the dedicated endpoint.
	EnableRoutePropagation bool `json:"enable_route_propagation"`

	// CreatedAt: creation date of the link.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the link.
	UpdatedAt *time.Time `json:"updated_at"`

	// Partner: partner host information.
	// Precisely one of Partner, Self must be set.
	Partner *PartnerHost `json:"partner,omitempty"`

	// Self: self-host information.
	// Precisely one of Partner, Self must be set.
	Self *SelfHost `json:"self,omitempty"`

	// Vlan: vLAN of the link.
	Vlan uint32 `json:"vlan"`

	// ScwBgpConfig: bGP configuration on Scaleway's side.
	ScwBgpConfig *BgpConfig `json:"scw_bgp_config"`

	// PeerBgpConfig: bGP configuration on peer's side (on-premises or other hosting provider).
	PeerBgpConfig *BgpConfig `json:"peer_bgp_config"`

	// RoutingPolicyV4ID: ID of the routing policy IPv4 attached to the link.
	RoutingPolicyV4ID *string `json:"routing_policy_v4_id"`

	// RoutingPolicyV6ID: ID of the routing policy IPv6 attached to the link.
	RoutingPolicyV6ID *string `json:"routing_policy_v6_id"`

	// Region: region of the link.
	Region scw.Region `json:"region"`
}

// Partner: partner.
type Partner struct {
	// ID: unique identifier of the partner.
	ID string `json:"id"`

	// Name: name of the partner.
	Name string `json:"name"`

	// ContactEmail: contact email address of partner.
	ContactEmail string `json:"contact_email"`

	// LogoURL: image URL of the partner's logo.
	LogoURL string `json:"logo_url"`

	// PortalURL: URL of the partner's portal.
	PortalURL string `json:"portal_url"`

	// CreatedAt: creation date of the partner.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the partner.
	UpdatedAt *time.Time `json:"updated_at"`
}

// Pop: pop.
type Pop struct {
	// ID: unique identifier of the PoP.
	ID string `json:"id"`

	// Name: name of the PoP. It is the common reference of Hosting DC (ex: TH2).
	Name string `json:"name"`

	// HostingProviderName: name of the PoP's hosting provider, e.g. Telehouse for TH2 or OpCore for DC3.
	HostingProviderName string `json:"hosting_provider_name"`

	// Address: physical address of the PoP.
	Address string `json:"address"`

	// City: city where PoP is located.
	City string `json:"city"`

	// LogoURL: image URL of the PoP's logo.
	LogoURL string `json:"logo_url"`

	// AvailableLinkBandwidthsMbps: available bandwidth in Mbits/s for future hosted links from available connections in this PoP.
	AvailableLinkBandwidthsMbps []uint64 `json:"available_link_bandwidths_mbps"`

	// Region: region of the PoP.
	Region scw.Region `json:"region"`
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

	// PrefixFilterIn: IP prefixes to accept from the peer (ranges of route announcements to accept).
	PrefixFilterIn []scw.IPNet `json:"prefix_filter_in"`

	// PrefixFilterOut: IP prefix filters to advertise to the peer (ranges of routes to advertise).
	PrefixFilterOut []scw.IPNet `json:"prefix_filter_out"`

	// CreatedAt: creation date of the routing policy.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the routing policy.
	UpdatedAt *time.Time `json:"updated_at"`

	// IsIPv6: IP prefixes version of the routing policy.
	IsIPv6 bool `json:"is_ipv6"`

	// Region: region of the routing policy.
	Region scw.Region `json:"region"`
}

// AttachRoutingPolicyRequest: attach routing policy request.
type AttachRoutingPolicyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// LinkID: ID of the link to attach a routing policy to.
	LinkID string `json:"-"`

	// RoutingPolicyID: ID of the routing policy to be attached.
	RoutingPolicyID string `json:"routing_policy_id"`
}

// AttachVpcRequest: attach vpc request.
type AttachVpcRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// LinkID: ID of the link to attach VPC to.
	LinkID string `json:"-"`

	// VpcID: ID of the VPC to attach.
	VpcID string `json:"vpc_id"`
}

// CreateLinkRequest: create link request.
type CreateLinkRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project to create the link in.
	ProjectID string `json:"project_id"`

	// Name: name of the link.
	Name string `json:"name"`

	// Tags: list of tags to apply to the link.
	Tags []string `json:"tags"`

	// PopID: poP (location) where the link will be created.
	PopID string `json:"pop_id"`

	// BandwidthMbps: desired bandwidth for the link. Must be compatible with available link bandwidths and remaining bandwidth capacity of the connection.
	BandwidthMbps uint64 `json:"bandwidth_mbps"`

	// ConnectionID: if set, creates a self-hosted link using this dedicated physical connection. As the customer, specify the ID of the physical connection you already have for this link.
	// Precisely one of ConnectionID, PartnerID must be set.
	ConnectionID *string `json:"connection_id,omitempty"`

	// PartnerID: if set, creates a hosted link on a partner's connection. Specify the ID of the chosen partner, who already has a shared connection with available bandwidth.
	// Precisely one of ConnectionID, PartnerID must be set.
	PartnerID *string `json:"partner_id,omitempty"`

	// PeerAsn: for self-hosted links we need the peer AS Number to establish BGP session. If not given, a default one will be assigned.
	PeerAsn *uint32 `json:"peer_asn,omitempty"`
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

	// PrefixFilterIn: IP prefixes to accept from the peer (ranges of route announcements to accept).
	PrefixFilterIn []scw.IPNet `json:"prefix_filter_in"`

	// PrefixFilterOut: IP prefix filters to advertise to the peer (ranges of routes to advertise).
	PrefixFilterOut []scw.IPNet `json:"prefix_filter_out"`

	// IsIPv6: IP prefixes version of the routing policy.
	IsIPv6 bool `json:"is_ipv6"`
}

// DeleteLinkRequest: delete link request.
type DeleteLinkRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// LinkID: ID of the link to delete.
	LinkID string `json:"-"`
}

// DeleteRoutingPolicyRequest: delete routing policy request.
type DeleteRoutingPolicyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// RoutingPolicyID: ID of the routing policy to delete.
	RoutingPolicyID string `json:"-"`
}

// DetachRoutingPolicyRequest: detach routing policy request.
type DetachRoutingPolicyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// LinkID: ID of the link to detach a routing policy from.
	LinkID string `json:"-"`

	// RoutingPolicyID: ID of the routing policy to be detached.
	RoutingPolicyID string `json:"routing_policy_id"`
}

// DetachVpcRequest: detach vpc request.
type DetachVpcRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// LinkID: ID of the link to detach the VPC from.
	LinkID string `json:"-"`
}

// DisableRoutePropagationRequest: disable route propagation request.
type DisableRoutePropagationRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// LinkID: ID of the link on which to disable route propagation.
	LinkID string `json:"-"`
}

// EnableRoutePropagationRequest: enable route propagation request.
type EnableRoutePropagationRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// LinkID: ID of the link on which to enable route propagation.
	LinkID string `json:"-"`
}

// GetDedicatedConnectionRequest: get dedicated connection request.
type GetDedicatedConnectionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ConnectionID: ID of connection to get.
	ConnectionID string `json:"-"`
}

// GetLinkRequest: get link request.
type GetLinkRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// LinkID: ID of the link to get.
	LinkID string `json:"-"`
}

// GetPartnerRequest: get partner request.
type GetPartnerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PartnerID: ID of partner to get.
	PartnerID string `json:"-"`
}

// GetPopRequest: get pop request.
type GetPopRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// PopID: ID of PoP to get.
	PopID string `json:"-"`
}

// GetRoutingPolicyRequest: get routing policy request.
type GetRoutingPolicyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// RoutingPolicyID: ID of the routing policy to get.
	RoutingPolicyID string `json:"-"`
}

// ListDedicatedConnectionsRequest: list dedicated connections request.
type ListDedicatedConnectionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListDedicatedConnectionsRequestOrderBy `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of connections to return per page.
	PageSize *uint32 `json:"-"`

	// ProjectID: project ID to filter for.
	ProjectID *string `json:"-"`

	// OrganizationID: organization ID to filter for.
	OrganizationID *string `json:"-"`

	// Name: link name to filter for.
	Name *string `json:"-"`

	// Tags: tags to filter for.
	Tags []string `json:"-"`

	// Status: connection status to filter for.
	// Default value: unknown_status
	Status DedicatedConnectionStatus `json:"-"`

	// BandwidthMbps: filter for dedicated connections with this bandwidth size.
	BandwidthMbps *uint64 `json:"-"`

	// PopID: filter for dedicated connections present in this PoP.
	PopID *string `json:"-"`
}

// ListDedicatedConnectionsResponse: list dedicated connections response.
type ListDedicatedConnectionsResponse struct {
	// Connections: list of connections on current page.
	Connections []*DedicatedConnection `json:"connections"`

	// TotalCount: total number of connections returned.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDedicatedConnectionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDedicatedConnectionsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListDedicatedConnectionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Connections = append(r.Connections, results.Connections...)
	r.TotalCount += uint64(len(results.Connections))
	return uint64(len(results.Connections)), nil
}

// ListLinksRequest: list links request.
type ListLinksRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListLinksRequestOrderBy `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of links to return per page.
	PageSize *uint32 `json:"-"`

	// ProjectID: project ID to filter for.
	ProjectID *string `json:"-"`

	// OrganizationID: organization ID to filter for.
	OrganizationID *string `json:"-"`

	// Name: link name to filter for.
	Name *string `json:"-"`

	// Tags: tags to filter for.
	Tags []string `json:"-"`

	// Status: link status to filter for.
	// Default value: unknown_link_status
	Status LinkStatus `json:"-"`

	// BgpV4Status: bGP IPv4 status to filter for.
	// Default value: unknown_bgp_status
	BgpV4Status BgpStatus `json:"-"`

	// BgpV6Status: bGP IPv6 status to filter for.
	// Default value: unknown_bgp_status
	BgpV6Status BgpStatus `json:"-"`

	// PopID: filter for links attached to this PoP (via connections).
	PopID *string `json:"-"`

	// BandwidthMbps: filter for link bandwidth (in Mbps).
	BandwidthMbps *uint64 `json:"-"`

	// PartnerID: filter for links hosted by this partner.
	PartnerID *string `json:"-"`

	// VpcID: filter for links attached to this VPC.
	VpcID *string `json:"-"`

	// RoutingPolicyID: filter for links using this routing policy.
	RoutingPolicyID *string `json:"-"`

	// PairingKey: filter for the link with this pairing_key.
	PairingKey *string `json:"-"`

	// Kind: filter for hosted or self-hosted links.
	// Default value: hosted
	Kind *LinkKind `json:"-"`

	// ConnectionID: filter for links self-hosted on this connection.
	ConnectionID *string `json:"-"`
}

// ListLinksResponse: list links response.
type ListLinksResponse struct {
	// Links: list of links on the current page.
	Links []*Link `json:"links"`

	// TotalCount: total number of links.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLinksResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLinksResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListLinksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Links = append(r.Links, results.Links...)
	r.TotalCount += uint64(len(results.Links))
	return uint64(len(results.Links)), nil
}

// ListPartnersRequest: list partners request.
type ListPartnersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: name_asc
	OrderBy ListPartnersRequestOrderBy `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of partners to return per page.
	PageSize *uint32 `json:"-"`

	// PopIDs: filter for partners present (offering a connection) in one of these PoPs.
	PopIDs []string `json:"-"`
}

// ListPartnersResponse: list partners response.
type ListPartnersResponse struct {
	// Partners: list of partners on current page.
	Partners []*Partner `json:"partners"`

	// TotalCount: total number of partners returned.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPartnersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPartnersResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListPartnersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Partners = append(r.Partners, results.Partners...)
	r.TotalCount += uint64(len(results.Partners))
	return uint64(len(results.Partners)), nil
}

// ListPopsRequest: list pops request.
type ListPopsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: name_asc
	OrderBy ListPopsRequestOrderBy `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of PoPs to return per page.
	PageSize *uint32 `json:"-"`

	// Name: poP name to filter for.
	Name *string `json:"-"`

	// HostingProviderName: hosting provider name to filter for.
	HostingProviderName *string `json:"-"`

	// PartnerID: filter for PoPs hosting an available shared connection from this partner.
	PartnerID *string `json:"-"`

	// LinkBandwidthMbps: filter for PoPs with a shared connection allowing this bandwidth size. Note that we cannot guarantee that PoPs returned will have available capacity.
	LinkBandwidthMbps *uint64 `json:"-"`

	// DedicatedAvailable: filter for PoPs with a dedicated connection available for self-hosted links.
	DedicatedAvailable *bool `json:"-"`
}

// ListPopsResponse: list pops response.
type ListPopsResponse struct {
	// Pops: list of PoPs on the current page.
	Pops []*Pop `json:"pops"`

	// TotalCount: total number of PoPs.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPopsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPopsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListPopsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Pops = append(r.Pops, results.Pops...)
	r.TotalCount += uint64(len(results.Pops))
	return uint64(len(results.Pops)), nil
}

// ListRoutingPoliciesRequest: list routing policies request.
type ListRoutingPoliciesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListRoutingPoliciesRequestOrderBy `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of routing policies to return per page.
	PageSize *uint32 `json:"-"`

	// ProjectID: project ID to filter for.
	ProjectID *string `json:"-"`

	// OrganizationID: organization ID to filter for.
	OrganizationID *string `json:"-"`

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

// UpdateLinkRequest: update link request.
type UpdateLinkRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// LinkID: ID of the link to update.
	LinkID string `json:"-"`

	// Name: name of the link.
	Name *string `json:"name,omitempty"`

	// Tags: list of tags to apply to the link.
	Tags *[]string `json:"tags,omitempty"`

	// PeerAsn: for self-hosted links, AS Number to establish BGP session.
	PeerAsn *uint32 `json:"peer_asn,omitempty"`
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

// This API allows you to manage your Scaleway InterLink, to connect your on-premises infrastructure with your Scaleway VPC.
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

// ListDedicatedConnections: For self-hosted users, list their dedicated physical connections in a given region. By default, the connections returned in the list are ordered by name in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListDedicatedConnections(req *ListDedicatedConnectionsRequest, opts ...scw.RequestOption) (*ListDedicatedConnectionsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "bandwidth_mbps", req.BandwidthMbps)
	parameter.AddToQuery(query, "pop_id", req.PopID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/dedicated-connections",
		Query:  query,
	}

	var resp ListDedicatedConnectionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDedicatedConnection: For self-hosted users, get a dedicated physical connection corresponding to the given ID. The response object includes information such as the connection's name, status and total bandwidth.
func (s *API) GetDedicatedConnection(req *GetDedicatedConnectionRequest, opts ...scw.RequestOption) (*DedicatedConnection, error) {
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
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/dedicated-connections/" + fmt.Sprint(req.ConnectionID) + "",
	}

	var resp DedicatedConnection

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPartners: List all available partners. By default, the partners returned in the list are ordered by name in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListPartners(req *ListPartnersRequest, opts ...scw.RequestOption) (*ListPartnersResponse, error) {
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
	parameter.AddToQuery(query, "pop_ids", req.PopIDs)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/partners",
		Query:  query,
	}

	var resp ListPartnersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPartner: Get a partner for the given partner IP. The response object includes information such as the partner's name, email address and portal URL.
func (s *API) GetPartner(req *GetPartnerRequest, opts ...scw.RequestOption) (*Partner, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PartnerID) == "" {
		return nil, errors.New("field PartnerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/partners/" + fmt.Sprint(req.PartnerID) + "",
	}

	var resp Partner

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPops: List all available PoPs (locations) for a given region. By default, the results are returned in ascending alphabetical order by name.
func (s *API) ListPops(req *ListPopsRequest, opts ...scw.RequestOption) (*ListPopsResponse, error) {
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
	parameter.AddToQuery(query, "hosting_provider_name", req.HostingProviderName)
	parameter.AddToQuery(query, "partner_id", req.PartnerID)
	parameter.AddToQuery(query, "link_bandwidth_mbps", req.LinkBandwidthMbps)
	parameter.AddToQuery(query, "dedicated_available", req.DedicatedAvailable)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/pops",
		Query:  query,
	}

	var resp ListPopsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPop: Get a PoP for the given PoP ID. The response object includes the PoP's name and information about its physical location.
func (s *API) GetPop(req *GetPopRequest, opts ...scw.RequestOption) (*Pop, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.PopID) == "" {
		return nil, errors.New("field PopID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/pops/" + fmt.Sprint(req.PopID) + "",
	}

	var resp Pop

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListLinks: List all your links (InterLink connections). A number of filters are available, including Project ID, name, tags and status.
func (s *API) ListLinks(req *ListLinksRequest, opts ...scw.RequestOption) (*ListLinksResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "bgp_v4_status", req.BgpV4Status)
	parameter.AddToQuery(query, "bgp_v6_status", req.BgpV6Status)
	parameter.AddToQuery(query, "pop_id", req.PopID)
	parameter.AddToQuery(query, "bandwidth_mbps", req.BandwidthMbps)
	parameter.AddToQuery(query, "partner_id", req.PartnerID)
	parameter.AddToQuery(query, "vpc_id", req.VpcID)
	parameter.AddToQuery(query, "routing_policy_id", req.RoutingPolicyID)
	parameter.AddToQuery(query, "pairing_key", req.PairingKey)
	parameter.AddToQuery(query, "kind", req.Kind)
	parameter.AddToQuery(query, "connection_id", req.ConnectionID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/links",
		Query:  query,
	}

	var resp ListLinksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLink: Get a link (InterLink session / logical InterLink resource) for the given link ID. The response object includes information about the link's various configuration details.
func (s *API) GetLink(req *GetLinkRequest, opts ...scw.RequestOption) (*Link, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LinkID) == "" {
		return nil, errors.New("field LinkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/links/" + fmt.Sprint(req.LinkID) + "",
	}

	var resp Link

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateLink: Create a link (InterLink session / logical InterLink resource) in a given PoP, specifying its various configuration details. Links can either be hosted (facilitated by partners' shared physical connections) or self-hosted (for users who have purchased a dedicated physical connection).
func (s *API) CreateLink(req *CreateLinkRequest, opts ...scw.RequestOption) (*Link, error) {
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
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/links",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Link

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateLink: Update an existing link, specified by its link ID. Only its name and tags can be updated.
func (s *API) UpdateLink(req *UpdateLinkRequest, opts ...scw.RequestOption) (*Link, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LinkID) == "" {
		return nil, errors.New("field LinkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/links/" + fmt.Sprint(req.LinkID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Link

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteLink: Delete an existing link, specified by its link ID. Note that as well as deleting the link here on the Scaleway side, it is also necessary to request deletion from the partner on their side. Only when this action has been carried out on both sides will the resource be completely deleted.
func (s *API) DeleteLink(req *DeleteLinkRequest, opts ...scw.RequestOption) (*Link, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LinkID) == "" {
		return nil, errors.New("field LinkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/links/" + fmt.Sprint(req.LinkID) + "",
	}

	var resp Link

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachVpc: Attach a VPC to an existing link. This facilitates communication between the resources in your Scaleway VPC, and your on-premises infrastructure.
func (s *API) AttachVpc(req *AttachVpcRequest, opts ...scw.RequestOption) (*Link, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LinkID) == "" {
		return nil, errors.New("field LinkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/links/" + fmt.Sprint(req.LinkID) + "/attach-vpc",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Link

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachVpc: Detach a VPC from an existing link.
func (s *API) DetachVpc(req *DetachVpcRequest, opts ...scw.RequestOption) (*Link, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LinkID) == "" {
		return nil, errors.New("field LinkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/links/" + fmt.Sprint(req.LinkID) + "/detach-vpc",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Link

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AttachRoutingPolicy: Attach a routing policy to an existing link. As all routes across the link are blocked by default, you must attach a routing policy to set IP prefix filters for allowed routes, facilitating traffic flow.
func (s *API) AttachRoutingPolicy(req *AttachRoutingPolicyRequest, opts ...scw.RequestOption) (*Link, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LinkID) == "" {
		return nil, errors.New("field LinkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/links/" + fmt.Sprint(req.LinkID) + "/attach-routing-policy",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Link

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DetachRoutingPolicy: Detach a routing policy from an existing link. Without a routing policy, all routes across the link are blocked by default.
func (s *API) DetachRoutingPolicy(req *DetachRoutingPolicyRequest, opts ...scw.RequestOption) (*Link, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LinkID) == "" {
		return nil, errors.New("field LinkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/links/" + fmt.Sprint(req.LinkID) + "/detach-routing-policy",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Link

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableRoutePropagation: Enable all allowed prefixes (defined in a routing policy) to be announced in the BGP session. This allows traffic to flow between the attached VPC and the on-premises infrastructure along the announced routes. Note that by default, even when route propagation is enabled, all routes are blocked. It is essential to attach a routing policy to define the ranges of routes to announce.
func (s *API) EnableRoutePropagation(req *EnableRoutePropagationRequest, opts ...scw.RequestOption) (*Link, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LinkID) == "" {
		return nil, errors.New("field LinkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/links/" + fmt.Sprint(req.LinkID) + "/enable-route-propagation",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Link

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableRoutePropagation: Prevent any prefixes from being announced in the BGP session. Traffic will not be able to flow over the InterLink until route propagation is re-enabled.
func (s *API) DisableRoutePropagation(req *DisableRoutePropagationRequest, opts ...scw.RequestOption) (*Link, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LinkID) == "" {
		return nil, errors.New("field LinkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/links/" + fmt.Sprint(req.LinkID) + "/disable-route-propagation",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Link

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRoutingPolicies: List all routing policies in a given region. A routing policy can be attached to one or multiple links (InterLink connections).
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "ipv6", req.IPv6)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routing-policies",
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
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routing-policies/" + fmt.Sprint(req.RoutingPolicyID) + "",
	}

	var resp RoutingPolicy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRoutingPolicy: Create a routing policy. Routing policies allow you to set IP prefix filters to define the incoming route announcements to accept from the peer, and the outgoing routes to announce to the peer.
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
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routing-policies",
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
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routing-policies/" + fmt.Sprint(req.RoutingPolicyID) + "",
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
		Path:   "/interlink/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routing-policies/" + fmt.Sprint(req.RoutingPolicyID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

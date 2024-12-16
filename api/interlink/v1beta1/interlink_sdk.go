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
		return "unknown_bgp_status"
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
		return "unknown_link_status"
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
		return "created_at_asc"
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
		return "name_asc"
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
		return "name_asc"
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
		return "created_at_asc"
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

	// PopID: ID of the PoP where the link's corresponding port is located.
	PopID string `json:"pop_id"`

	// PartnerID: ID of the partner facilitating this link.
	PartnerID *string `json:"partner_id"`

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

	// RoutingPolicyID: ID of the routing policy attached to the link.
	RoutingPolicyID *string `json:"routing_policy_id"`

	// EnableRoutePropagation: defines whether route propagation is enabled or not. To enable or disable route propagation, use the dedicated endpoint.
	EnableRoutePropagation bool `json:"enable_route_propagation"`

	// CreatedAt: creation date of the link.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last modification date of the link.
	UpdatedAt *time.Time `json:"updated_at"`

	// PairingKey: used to identify a link from a user or partner's point of view.
	PairingKey string `json:"pairing_key"`

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

	// BandwidthMbps: desired bandwidth for the link. Must be compatible with available link bandwidths and remaining bandwidth capacity of the port.
	BandwidthMbps uint64 `json:"bandwidth_mbps"`

	// Dedicated: if true, a dedicated link (1 link per port, dedicated to one customer) will be crated. It is not necessary to specify a `port_id` or `partner_id`. A new port will created and assigned to the link. Note that Scaleway has not yet enabled the creation of dedicated links, this field is reserved for future use.
	// Precisely one of Dedicated, PortID, PartnerID must be set.
	Dedicated *bool `json:"dedicated,omitempty"`

	// PortID: if set, a shared link (N links per port, one of which is this customer's port) will be created. As the customer, specify the ID of the port you already have for this link. Note that shared links are not currently available. Note that Scaleway has not yet enabled the creation of shared links, this field is reserved for future use.
	// Precisely one of Dedicated, PortID, PartnerID must be set.
	PortID *string `json:"port_id,omitempty"`

	// PartnerID: if set, a hosted link (N links per port on a partner port) will be created. Specify the ID of the chosen partner, who already has a shareable port with available bandwidth. Note that this is currently the only type of link offered by Scaleway, and therefore this field must be set when creating a link.
	// Precisely one of Dedicated, PortID, PartnerID must be set.
	PartnerID *string `json:"partner_id,omitempty"`
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

	// PopID: filter for links attached to this PoP (via ports).
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
func (r *ListLinksResponse) UnsafeAppend(res interface{}) (uint64, error) {
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

	// PopIDs: filter for partners present (offering a port) in one of these PoPs.
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
func (r *ListPartnersResponse) UnsafeAppend(res interface{}) (uint64, error) {
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

	// PartnerID: filter for PoPs hosting an available shared port from this partner.
	PartnerID *string `json:"-"`

	// LinkBandwidthMbps: filter for PoPs with a shared port allowing this bandwidth size. Note that we cannot guarantee that PoPs returned will have available capacity.
	LinkBandwidthMbps *uint64 `json:"-"`
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
func (r *ListPopsResponse) UnsafeAppend(res interface{}) (uint64, error) {
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
func (r *ListRoutingPoliciesResponse) UnsafeAppend(res interface{}) (uint64, error) {
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

// GetLink: Get a link (InterLink connection) for the given link ID. The response object includes information about the link's various configuration details.
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

// CreateLink: Create a link (InterLink connection) in a given PoP, specifying its various configuration details. For the moment only hosted links (faciliated by partners) are available, though in the future dedicated and shared links will also be possible.
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

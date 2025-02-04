// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package baremetal provides methods and message types of the baremetal v3 API.
package baremetal

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
		return "created_at_asc"
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

type ServerPrivateNetworkStatus string

const (
	ServerPrivateNetworkStatusUnknownStatus = ServerPrivateNetworkStatus("unknown_status")
	ServerPrivateNetworkStatusAttaching     = ServerPrivateNetworkStatus("attaching")
	ServerPrivateNetworkStatusAttached      = ServerPrivateNetworkStatus("attached")
	ServerPrivateNetworkStatusError         = ServerPrivateNetworkStatus("error")
	ServerPrivateNetworkStatusDetaching     = ServerPrivateNetworkStatus("detaching")
	ServerPrivateNetworkStatusLocked        = ServerPrivateNetworkStatus("locked")
)

func (enum ServerPrivateNetworkStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum ServerPrivateNetworkStatus) Values() []ServerPrivateNetworkStatus {
	return []ServerPrivateNetworkStatus{
		"unknown_status",
		"attaching",
		"attached",
		"error",
		"detaching",
		"locked",
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

// ServerPrivateNetwork: server private network.
type ServerPrivateNetwork struct {
	// ID: UUID of the Server-to-Private Network mapping.
	ID string `json:"id"`

	// ProjectID: private Network Project UUID.
	ProjectID string `json:"project_id"`

	// ServerID: server UUID.
	ServerID string `json:"server_id"`

	// PrivateNetworkID: private Network UUID.
	PrivateNetworkID string `json:"private_network_id"`

	// Vlan: vLAN UUID associated with the Private Network.
	Vlan *uint32 `json:"vlan"`

	// Status: configuration status of the Private Network.
	// Default value: unknown_status
	Status ServerPrivateNetworkStatus `json:"status"`

	// CreatedAt: private Network creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the Private Network was last modified.
	UpdatedAt *time.Time `json:"updated_at"`

	// IpamIPIDs: iPAM IP IDs of the server, if it has any.
	IpamIPIDs []string `json:"ipam_ip_ids"`
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

// PrivateNetworkAPIAddServerPrivateNetworkRequest: private network api add server private network request.
type PrivateNetworkAPIAddServerPrivateNetworkRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: UUID of the server.
	ServerID string `json:"-"`

	// PrivateNetworkID: UUID of the Private Network.
	PrivateNetworkID string `json:"private_network_id"`

	// IpamIPIDs: iPAM IDs of an IPs to attach to the server.
	IpamIPIDs []string `json:"ipam_ip_ids"`
}

// PrivateNetworkAPIDeleteServerPrivateNetworkRequest: private network api delete server private network request.
type PrivateNetworkAPIDeleteServerPrivateNetworkRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: UUID of the server.
	ServerID string `json:"-"`

	// PrivateNetworkID: UUID of the Private Network.
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

	// ServerID: filter Private Networks by server UUID.
	ServerID *string `json:"-"`

	// PrivateNetworkID: filter Private Networks by Private Network UUID.
	PrivateNetworkID *string `json:"-"`

	// OrganizationID: filter Private Networks by organization UUID.
	OrganizationID *string `json:"-"`

	// ProjectID: filter Private Networks by project UUID.
	ProjectID *string `json:"-"`

	// IpamIPIDs: filter Private Networks by IPAM IP UUIDs.
	IpamIPIDs []string `json:"-"`
}

// PrivateNetworkAPISetServerPrivateNetworksRequest: private network api set server private networks request.
type PrivateNetworkAPISetServerPrivateNetworksRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// ServerID: UUID of the server.
	ServerID string `json:"-"`

	// PerPrivateNetworkIpamIPIDs: object where the keys are the UUIDs of Private Networks and the values are arrays of IPAM IDs representing the IPs to assign to this Elastic Metal server on the Private Network. If the array supplied for a Private Network is empty, the next available IP from the Private Network's CIDR block will automatically be used for attachment.
	PerPrivateNetworkIpamIPIDs map[string]*[]string `json:"per_private_network_ipam_ip_ids"`
}

// SetServerPrivateNetworksResponse: set server private networks response.
type SetServerPrivateNetworksResponse struct {
	ServerPrivateNetworks []*ServerPrivateNetwork `json:"server_private_networks"`
}

// Elastic Metal - Private Networks API.
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
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZoneNlAms2, scw.ZonePlWaw2, scw.ZonePlWaw3}
}

// AddServerPrivateNetwork: Add an Elastic Metal server to a Private Network.
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
		Path:   "/baremetal/v3/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private-networks",
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

// SetServerPrivateNetworks: Configure multiple Private Networks on an Elastic Metal server.
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
		Path:   "/baremetal/v3/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private-networks",
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

// ListServerPrivateNetworks: List the Private Networks of an Elastic Metal server.
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
		Path:   "/baremetal/v3/zones/" + fmt.Sprint(req.Zone) + "/server-private-networks",
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
		Path:   "/baremetal/v3/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package vpc_private provides methods and message types of the vpc_private v1 API.
package vpc_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	vpc_v1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/vpc/v1"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/scw"
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

type PrivateNICStatus string

const (
	PrivateNICStatusUnknownPrivateNicStatus = PrivateNICStatus("unknown_private_nic_status")
	PrivateNICStatusStopped                 = PrivateNICStatus("stopped")
	PrivateNICStatusAvailable               = PrivateNICStatus("available")
	PrivateNICStatusConfiguring             = PrivateNICStatus("configuring")
	PrivateNICStatusError                   = PrivateNICStatus("error")
)

func (enum PrivateNICStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_private_nic_status"
	}
	return string(enum)
}

func (enum PrivateNICStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PrivateNICStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PrivateNICStatus(PrivateNICStatus(tmp).String())
	return nil
}

// PrivateNICMetadata: private nic metadata.
type PrivateNICMetadata struct {
	ProductType string `json:"product_type"`

	ResourceType string `json:"resource_type"`

	ResourceID string `json:"resource_id"`
}

// PrivateNIC: private nic.
type PrivateNIC struct {
	ID string `json:"id"`

	PrivateNetworkID string `json:"private_network_id"`

	ServerID string `json:"server_id"`

	MacAddress string `json:"mac_address"`

	Metadata *PrivateNICMetadata `json:"metadata"`

	// Status: default value: unknown_private_nic_status
	Status PrivateNICStatus `json:"status"`

	Vlan uint32 `json:"vlan"`
}

// PrivateNetworkWithNics: private network with nics.
type PrivateNetworkWithNics struct {
	PrivateNetwork *vpc_v1.PrivateNetwork `json:"private_network"`

	PrivateNics []*PrivateNIC `json:"private_nics"`
}

// ConsoleAPIGetPrivateNetworkRequest: console api get private network request.
type ConsoleAPIGetPrivateNetworkRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// PrivateNetworkID: private Network ID.
	PrivateNetworkID string `json:"-"`
}

// ConsoleAPIGetServiceInfoRequest: console api get service info request.
type ConsoleAPIGetServiceInfoRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// ConsoleAPIListPrivateNetworksRequest: console api list private networks request.
type ConsoleAPIListPrivateNetworksRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// OrderBy: sort order of the returned Private Networks.
	// Default value: created_at_asc
	OrderBy vpc_v1.ListPrivateNetworksRequestOrderBy `json:"-"`

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

	// IncludeRegional: defines whether to include regional Private Networks in the response.
	IncludeRegional *bool `json:"-"`
}

// ListPrivateNetworksResponse: list private networks response.
type ListPrivateNetworksResponse struct {
	PrivateNetworks []*PrivateNetworkWithNics `json:"private_networks"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPrivateNetworksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPrivateNetworksResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPrivateNetworksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PrivateNetworks = append(r.PrivateNetworks, results.PrivateNetworks...)
	r.TotalCount += uint32(len(results.PrivateNetworks))
	return uint32(len(results.PrivateNetworks)), nil
}

type ConsoleAPI struct {
	client *scw.Client
}

// NewConsoleAPI returns a ConsoleAPI object from a Scaleway client.
func NewConsoleAPI(client *scw.Client) *ConsoleAPI {
	return &ConsoleAPI{
		client: client,
	}
}

// GetServiceInfo:
func (s *ConsoleAPI) GetServiceInfo(req *ConsoleAPIGetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
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
		Path:   "/vpc-private/v1/zones/" + fmt.Sprint(req.Zone) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPrivateNetworks:
func (s *ConsoleAPI) ListPrivateNetworks(req *ConsoleAPIListPrivateNetworksRequest, opts ...scw.RequestOption) (*ListPrivateNetworksResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "private_network_ids", req.PrivateNetworkIDs)
	parameter.AddToQuery(query, "include_regional", req.IncludeRegional)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-private/v1/zones/" + fmt.Sprint(req.Zone) + "/private-networks",
		Query:  query,
	}

	var resp ListPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPrivateNetwork:
func (s *ConsoleAPI) GetPrivateNetwork(req *ConsoleAPIGetPrivateNetworkRequest, opts ...scw.RequestOption) (*PrivateNetworkWithNics, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.PrivateNetworkID) == "" {
		return nil, errors.New("field PrivateNetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vpc-private/v1/zones/" + fmt.Sprint(req.Zone) + "/private-networks/" + fmt.Sprint(req.PrivateNetworkID) + "",
	}

	var resp PrivateNetworkWithNics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package k8s_private provides methods and message types of the k8s_private v1beta2 API.
package k8s_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/k8s_private/v1beta2/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/k8s_private/v1beta2/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/k8s_private/v1beta2/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/k8s_private/v1beta2/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/k8s_private/v1beta2/scw"
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

// CommercialTypeDetail: commercial type detail.
type CommercialTypeDetail struct {
	// Region: region.
	Region scw.Region `json:"region"`

	// ID: iD.
	ID string `json:"id"`

	// Label: name of commercial type.
	Label string `json:"label"`

	// Zone: zone.
	Zone scw.Zone `json:"zone"`
}

// ClusterDashboardInfo: cluster dashboard info.
type ClusterDashboardInfo struct {
	URL string `json:"url"`

	Cookie string `json:"cookie"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// ConsoleAPIGetClusterDashboardInfoRequest: console api get cluster dashboard info request.
type ConsoleAPIGetClusterDashboardInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ClusterID string `json:"-"`
}

// ConsoleAPIListCommercialTypesRequest: console api list commercial types request.
type ConsoleAPIListCommercialTypesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Zone: in which zone to get the commercial types.
	Zone scw.Zone `json:"zone"`
}

// ListCommercialTypesResponse: list commercial types response.
type ListCommercialTypesResponse struct {
	// CommercialTypes: list of available commercial types.
	CommercialTypes []*CommercialTypeDetail `json:"commercial_types"`
}

// Kapsule Console API.
type ConsoleAPI struct {
	client *scw.Client
}

// NewConsoleAPI returns a ConsoleAPI object from a Scaleway client.
func NewConsoleAPI(client *scw.Client) *ConsoleAPI {
	return &ConsoleAPI{
		client: client,
	}
}
func (s *ConsoleAPI) Regions() []scw.Region {
	return []scw.Region{}
}

// ListCommercialTypes: List available commercial types.
func (s *ConsoleAPI) ListCommercialTypes(req *ConsoleAPIListCommercialTypesRequest, opts ...scw.RequestOption) (*ListCommercialTypesResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "zone", req.Zone)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s-private/v1beta2/regions/" + fmt.Sprint(req.Region) + "/catalog/commercial-types",
		Query:  query,
	}

	var resp ListCommercialTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetClusterDashboardInfo:
func (s *ConsoleAPI) GetClusterDashboardInfo(req *ConsoleAPIGetClusterDashboardInfoRequest, opts ...scw.RequestOption) (*ClusterDashboardInfo, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ClusterID) == "" {
		return nil, errors.New("field ClusterID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/k8s-private/v1beta2/regions/" + fmt.Sprint(req.Region) + "/clusters/" + fmt.Sprint(req.ClusterID) + "/dashboard-info",
	}

	var resp ClusterDashboardInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

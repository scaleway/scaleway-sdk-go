// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package search provides methods and message types of the search v1alpha1 API.
package search

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

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

type ResourceType string

const (
	// Unknown type.
	ResourceTypeUnknownType = ResourceType("unknown_type")
	// Instance server.
	ResourceTypeInstanceServer = ResourceType("instance_server")
	// K8S cluster.
	ResourceTypeK8sCluster = ResourceType("k8s_cluster")
	// K8S pool.
	ResourceTypeK8sPool = ResourceType("k8s_pool")
	// K8S node.
	ResourceTypeK8sNode = ResourceType("k8s_node")
	// DNS zone.
	ResourceTypeDNSZone = ResourceType("dns_zone")
	// VPC private network.
	ResourceTypeVpcPrivateNetwork = ResourceType("vpc_private_network")
	// RDB instance.
	ResourceTypeRdbInstance = ResourceType("rdb_instance")
	// Baremetal server.
	ResourceTypeBaremetalServer = ResourceType("baremetal_server")
	// LB server.
	ResourceTypeLBServer = ResourceType("lb_server")
	// Serverless functions function.
	ResourceTypeServerlessFunctionsFunction = ResourceType("serverless_functions_function")
	// Serverless containers container.
	ResourceTypeServerlessContainersContainer = ResourceType("serverless_containers_container")
	// Redis cluster.
	ResourceTypeRedisCluster = ResourceType("redis_cluster")
)

func (enum ResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum ResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ResourceType(ResourceType(tmp).String())
	return nil
}

// Resource: resource.
type Resource struct {
	// ID: ID of the resource.
	ID string `json:"id"`

	// Name: name of the resource.
	Name *string `json:"name"`

	// ProjectID: ID of the Project the resource belongs to.
	ProjectID *string `json:"project_id"`

	// OrganizationID: ID of the Organization the resource belongs to.
	OrganizationID string `json:"organization_id"`

	// Type: type of the resource.
	// Default value: unknown_type
	Type ResourceType `json:"type"`

	// Locality: locality the resource is in.
	Locality string `json:"locality"`
}

// SearchResourcesRequest: search resources request.
type SearchResourcesRequest struct {
	// Query: search query.
	Query string `json:"-"`

	// OrganizationID: ID of the Organization to search in.
	OrganizationID string `json:"-"`
}

// SearchResourcesResponse: search resources response.
type SearchResourcesResponse struct {
	// Resources: top resources found.
	Resources []*Resource `json:"resources"`
}

type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// SearchResources: Search API.
func (s *API) SearchResources(req *SearchResourcesRequest, opts ...scw.RequestOption) (*SearchResourcesResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "query", req.Query)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/search/v1alpha1/resources",
		Query:  query,
	}

	var resp SearchResourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
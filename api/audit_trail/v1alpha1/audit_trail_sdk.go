// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package audit_trail provides methods and message types of the audit_trail v1alpha1 API.
package audit_trail

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

type ListEventsRequestOrderBy string

const (
	ListEventsRequestOrderByRecordedAtDesc = ListEventsRequestOrderBy("recorded_at_desc")
	ListEventsRequestOrderByRecordedAtAsc  = ListEventsRequestOrderBy("recorded_at_asc")
)

func (enum ListEventsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListEventsRequestOrderByRecordedAtDesc)
	}
	return string(enum)
}

func (enum ListEventsRequestOrderBy) Values() []ListEventsRequestOrderBy {
	return []ListEventsRequestOrderBy{
		"recorded_at_desc",
		"recorded_at_asc",
	}
}

func (enum ListEventsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListEventsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListEventsRequestOrderBy(ListEventsRequestOrderBy(tmp).String())
	return nil
}

type ResourceType string

const (
	ResourceTypeUnknownType          = ResourceType("unknown_type")
	ResourceTypeSecmSecret           = ResourceType("secm_secret")
	ResourceTypeSecmSecretVersion    = ResourceType("secm_secret_version")
	ResourceTypeKubeCluster          = ResourceType("kube_cluster")
	ResourceTypeKubePool             = ResourceType("kube_pool")
	ResourceTypeKubeNode             = ResourceType("kube_node")
	ResourceTypeKubeACL              = ResourceType("kube_acl")
	ResourceTypeKeymKey              = ResourceType("keym_key")
	ResourceTypeIamUser              = ResourceType("iam_user")
	ResourceTypeIamApplication       = ResourceType("iam_application")
	ResourceTypeIamGroup             = ResourceType("iam_group")
	ResourceTypeIamPolicy            = ResourceType("iam_policy")
	ResourceTypeIamAPIKey            = ResourceType("iam_api_key")
	ResourceTypeIamSSHKey            = ResourceType("iam_ssh_key")
	ResourceTypeIamRule              = ResourceType("iam_rule")
	ResourceTypeSecretManagerSecret  = ResourceType("secret_manager_secret")
	ResourceTypeSecretManagerVersion = ResourceType("secret_manager_version")
	ResourceTypeKeyManagerKey        = ResourceType("key_manager_key")
	ResourceTypeAccountUser          = ResourceType("account_user")
	ResourceTypeAccountOrganization  = ResourceType("account_organization")
)

func (enum ResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ResourceTypeUnknownType)
	}
	return string(enum)
}

func (enum ResourceType) Values() []ResourceType {
	return []ResourceType{
		"unknown_type",
		"secm_secret",
		"secm_secret_version",
		"kube_cluster",
		"kube_pool",
		"kube_node",
		"kube_acl",
		"keym_key",
		"iam_user",
		"iam_application",
		"iam_group",
		"iam_policy",
		"iam_api_key",
		"iam_ssh_key",
		"iam_rule",
		"secret_manager_secret",
		"secret_manager_version",
		"key_manager_key",
		"account_user",
		"account_organization",
	}
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

// AccountOrganizationInfo: account organization info.
type AccountOrganizationInfo struct{}

// AccountUserInfo: account user info.
type AccountUserInfo struct {
	Email string `json:"email"`
}

// KeyManagerKeyInfo: key manager key info.
type KeyManagerKeyInfo struct{}

// KubernetesACLInfo: kubernetes acl info.
type KubernetesACLInfo struct{}

// KubernetesClusterInfo: kubernetes cluster info.
type KubernetesClusterInfo struct{}

// KubernetesNodeInfo: kubernetes node info.
type KubernetesNodeInfo struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

// KubernetesPoolInfo: kubernetes pool info.
type KubernetesPoolInfo struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

// SecretManagerSecretInfo: secret manager secret info.
type SecretManagerSecretInfo struct {
	Path string `json:"path"`
}

// SecretManagerSecretVersionInfo: secret manager secret version info.
type SecretManagerSecretVersionInfo struct {
	Revision uint32 `json:"revision"`
}

// EventPrincipal: event principal.
type EventPrincipal struct {
	ID string `json:"id"`
}

// Resource: resource.
type Resource struct {
	ID string `json:"id"`

	// Type: default value: unknown_type
	Type ResourceType `json:"type"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	DeletedAt *time.Time `json:"deleted_at"`

	Name *string `json:"name"`

	// Deprecated
	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo must be set.
	SecmSecretInfo *SecretManagerSecretInfo `json:"secm_secret_info,omitempty"`

	// Deprecated
	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo must be set.
	SecmSecretVersionInfo *SecretManagerSecretVersionInfo `json:"secm_secret_version_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo must be set.
	KubeClusterInfo *KubernetesClusterInfo `json:"kube_cluster_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo must be set.
	KubePoolInfo *KubernetesPoolInfo `json:"kube_pool_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo must be set.
	KubeNodeInfo *KubernetesNodeInfo `json:"kube_node_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo must be set.
	KubeACLInfo *KubernetesACLInfo `json:"kube_acl_info,omitempty"`

	// Deprecated
	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo must be set.
	KeymKeyInfo *KeyManagerKeyInfo `json:"keym_key_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo must be set.
	SecretManagerSecretInfo *SecretManagerSecretInfo `json:"secret_manager_secret_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo must be set.
	SecretManagerVersionInfo *SecretManagerSecretVersionInfo `json:"secret_manager_version_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo must be set.
	KeyManagerKeyInfo *KeyManagerKeyInfo `json:"key_manager_key_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo must be set.
	AccountUserInfo *AccountUserInfo `json:"account_user_info,omitempty"`

	// Precisely one of SecmSecretInfo, SecmSecretVersionInfo, KubeClusterInfo, KubePoolInfo, KubeNodeInfo, KubeACLInfo, KeymKeyInfo, SecretManagerSecretInfo, SecretManagerVersionInfo, KeyManagerKeyInfo, AccountUserInfo, AccountOrganizationInfo must be set.
	AccountOrganizationInfo *AccountOrganizationInfo `json:"account_organization_info,omitempty"`
}

// ProductService: product service.
type ProductService struct {
	Name string `json:"name"`

	Methods []string `json:"methods"`
}

// Event: event.
type Event struct {
	// ID: ID of the event.
	ID string `json:"id"`

	// RecordedAt: timestamp of the event.
	RecordedAt *time.Time `json:"recorded_at"`

	// Locality: locality of the resource attached to the event.
	Locality string `json:"locality"`

	// Principal: user or IAM application at the origin of the event.
	Principal *EventPrincipal `json:"principal"`

	// OrganizationID: organization ID containing the event.
	OrganizationID string `json:"organization_id"`

	// ProjectID: (Optional) Project of the resource attached to the event.
	ProjectID *string `json:"project_id"`

	// SourceIP: IP address at the origin of the event.
	SourceIP net.IP `json:"source_ip"`

	// UserAgent: user Agent at the origin of the event.
	UserAgent *string `json:"user_agent"`

	// ProductName: product name of the resource attached to the event.
	ProductName string `json:"product_name"`

	// ServiceName: API name called to trigger the event.
	ServiceName string `json:"service_name"`

	// MethodName: API method called to trigger the event.
	MethodName string `json:"method_name"`

	// Deprecated: Resource: resource attached to the event.
	Resource *Resource `json:"resource,omitempty"`

	// Resources: resources attached to the event.
	Resources []*Resource `json:"resources"`

	// RequestID: unique identifier of the request at the origin of the event.
	RequestID string `json:"request_id"`

	// RequestBody: request at the origin of the event.
	RequestBody *scw.JSONObject `json:"request_body"`

	// StatusCode: HTTP status code resulting of the API call.
	StatusCode uint32 `json:"status_code"`
}

// Product: product.
type Product struct {
	// Title: product title.
	Title string `json:"title"`

	// Name: product name.
	Name string `json:"name"`

	// Services: specifies the API versions of the products integrated with Audit Trail. Each version defines the methods logged by Audit Trail.
	Services []*ProductService `json:"services"`
}

// ListEventsRequest: list events request.
type ListEventsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: (Optional) ID of the Project containing the Audit Trail events.
	ProjectID *string `json:"-"`

	// OrganizationID: ID of the Organization containing the Audit Trail events.
	OrganizationID string `json:"-"`

	// ResourceType: (Optional) Returns a paginated list of Scaleway resources' features.
	// Default value: unknown_type
	ResourceType ResourceType `json:"-"`

	// MethodName: (Optional) Name of the method of the API call performed.
	MethodName *string `json:"-"`

	// Status: (Optional) HTTP status code of the request. Returns either `200` if the request was successful or `403` if the permission was denied.
	Status *uint32 `json:"-"`

	// RecordedAfter: (Optional) The `recorded_after` parameter defines the earliest timestamp from which Audit Trail events are retrieved. Returns `one hour ago` by default.
	RecordedAfter *time.Time `json:"-"`

	// RecordedBefore: (Optional) The `recorded_before` parameter defines the latest timestamp up to which Audit Trail events are retrieved. Returns `now` by default.
	RecordedBefore *time.Time `json:"-"`

	// OrderBy: default value: recorded_at_desc
	OrderBy ListEventsRequestOrderBy `json:"-"`

	PageSize *uint32 `json:"-"`

	PageToken *string `json:"-"`

	// ProductName: (Optional) Name of the Scaleway resource in a hyphenated format.
	ProductName *string `json:"-"`

	// ServiceName: (Optional) Name of the service of the API call performed.
	ServiceName *string `json:"-"`
}

// ListEventsResponse: list events response.
type ListEventsResponse struct {
	// Events: single page of events matching the requested criteria.
	Events []*Event `json:"events"`

	// NextPageToken: page token to use in following calls to keep listing.
	NextPageToken *string `json:"next_page_token"`
}

// ListProductsRequest: list products request.
type ListProductsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrganizationID: ID of the Organization containing the Audit Trail events.
	OrganizationID string `json:"organization_id"`
}

// ListProductsResponse: list products response.
type ListProductsResponse struct {
	// Products: list of all products integrated with Audit Trail.
	Products []*Product `json:"products"`

	// TotalCount: number of integrated products.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListProductsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListProductsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListProductsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Products = append(r.Products, results.Products...)
	r.TotalCount += uint64(len(results.Products))
	return uint64(len(results.Products)), nil
}

// This API allows you to ensure accountability and security by recording events and changes performed within your Scaleway Organization.
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
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms}
}

// ListEvents: Retrieve the list of Audit Trail events for a Scaleway Organization and/or Project. You must specify the `organization_id` and optionally, the `project_id`.
func (s *API) ListEvents(req *ListEventsRequest, opts ...scw.RequestOption) (*ListEventsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "resource_type", req.ResourceType)
	parameter.AddToQuery(query, "method_name", req.MethodName)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "recorded_after", req.RecordedAfter)
	parameter.AddToQuery(query, "recorded_before", req.RecordedBefore)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page_token", req.PageToken)
	parameter.AddToQuery(query, "product_name", req.ProductName)
	parameter.AddToQuery(query, "service_name", req.ServiceName)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/audit-trail/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/events",
		Query:  query,
	}

	var resp ListEventsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListProducts: Retrieve the list of Scaleway resources for which you have Audit Trail events.
func (s *API) ListProducts(req *ListProductsRequest, opts ...scw.RequestOption) (*ListProductsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/audit-trail/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/products",
		Query:  query,
	}

	var resp ListProductsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

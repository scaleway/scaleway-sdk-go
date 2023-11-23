// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package secret_admin provides methods and message types of the secret_admin v1 API.
package secret_admin

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
	secret_v1alpha1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/secret/v1alpha1"
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

type ListSecretsRequestOrderBy string

const (
	ListSecretsRequestOrderByNameAsc       = ListSecretsRequestOrderBy("name_asc")
	ListSecretsRequestOrderByNameDesc      = ListSecretsRequestOrderBy("name_desc")
	ListSecretsRequestOrderByCreatedAtAsc  = ListSecretsRequestOrderBy("created_at_asc")
	ListSecretsRequestOrderByCreatedAtDesc = ListSecretsRequestOrderBy("created_at_desc")
	ListSecretsRequestOrderByUpdatedAtAsc  = ListSecretsRequestOrderBy("updated_at_asc")
	ListSecretsRequestOrderByUpdatedAtDesc = ListSecretsRequestOrderBy("updated_at_desc")
)

func (enum ListSecretsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListSecretsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSecretsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSecretsRequestOrderBy(ListSecretsRequestOrderBy(tmp).String())
	return nil
}

// CreateManagedSecretRequest: create managed secret request.
type CreateManagedSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ClientProjectID: client Project holding the Secret.
	ClientProjectID string `json:"client_project_id"`

	// ClientOrganizationID: client Organization holding the Secret.
	ClientOrganizationID string `json:"client_organization_id"`

	// Deprecated: ProductName: (Deprecated: use `product` field) Name of the product to add.
	ProductName *string `json:"product_name,omitempty"`

	// Product: see `Product` enum for description of values.
	// Default value: unknown
	Product secret_v1alpha1.Product `json:"product"`

	// Name: name of the secret.
	Name string `json:"name"`

	// Tags: list of the secret's tags.
	Tags []string `json:"tags"`

	// Description: description of the secret.
	Description *string `json:"description,omitempty"`

	// Type: (Optional.) See `SecretType` enum for description of values. If not specified, the type is `Opaque`.
	// Default value: unknown_secret_type
	Type secret_v1alpha1.SecretType `json:"type"`

	// Path: location of the secret in the directory structure.
	Path *string `json:"path,omitempty"`
}

// CreateManagedSecretResponse: create managed secret response.
type CreateManagedSecretResponse struct {
	// Secret: secret.
	Secret *secret_v1alpha1.Secret `json:"secret"`
}

// GetSecretRequest: get secret request.
type GetSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`
}

// GetSecretResponse: get secret response.
type GetSecretResponse struct {
	// Secret: secret.
	Secret *secret_v1alpha1.Secret `json:"secret"`
}

// GetSecretVersionRequest: get secret version request.
type GetSecretVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	// Revision: version number. The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

// GetSecretVersionResponse: get secret version response.
type GetSecretVersionResponse struct {
	// Version: secret version.
	Version *secret_v1alpha1.SecretVersion `json:"version"`
}

// ListSecretVersionsRequest: list secret versions request.
type ListSecretVersionsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// Status: filter results by status.
	Status []secret_v1alpha1.SecretVersionStatus `json:"-"`
}

// ListSecretVersionsResponse: list secret versions response.
type ListSecretVersionsResponse struct {
	// Versions: single page of versions.
	Versions []*secret_v1alpha1.SecretVersion `json:"versions"`

	// TotalCount: number of versions.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecretVersionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecretVersionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSecretVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Versions = append(r.Versions, results.Versions...)
	r.TotalCount += uint32(len(results.Versions))
	return uint32(len(results.Versions)), nil
}

// ListSecretsRequest: list secrets request.
type ListSecretsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrganizationID: filter by Organization ID (optional).
	OrganizationID *string `json:"-"`

	// ProjectID: filter by Project ID (optional).
	ProjectID *string `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListSecretsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// Tags: list of tags to filter on (optional).
	Tags []string `json:"-"`

	// Name: filter by secret name (optional).
	Name *string `json:"-"`

	// IsManaged: filter by managed / not managed (optional).
	IsManaged *bool `json:"-"`

	// Path: filter by path.
	Path *string `json:"-"`
}

// ListSecretsResponse: list secrets response.
type ListSecretsResponse struct {
	// Secrets: single page of secrets matching the requested criteria.
	Secrets []*secret_v1alpha1.Secret `json:"secrets"`

	// TotalCount: count of all secrets matching the requested criteria.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecretsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecretsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSecretsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Secrets = append(r.Secrets, results.Secrets...)
	r.TotalCount += uint32(len(results.Secrets))
	return uint32(len(results.Secrets)), nil
}

// This API is for the admin console of Secret Manager.
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
	return []scw.Region{scw.RegionFrPar}
}

// CreateManagedSecret: Create a managed secret.
func (s *API) CreateManagedSecret(req *CreateManagedSecretRequest, opts ...scw.RequestOption) (*CreateManagedSecretResponse, error) {
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
		Path:   "/secret-manager-admin/v1/regions/" + fmt.Sprint(req.Region) + "/managed-secrets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateManagedSecretResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSecret: Retrieve the metadata of a secret specified by the `region` and the `secret_name` parameters.
func (s *API) GetSecret(req *GetSecretRequest, opts ...scw.RequestOption) (*GetSecretResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager-admin/v1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
	}

	var resp GetSecretResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSecrets: Retrieve the list of secrets created within an Organization and/or Project. You must specify either the `organization_id` or the `project_id` and the `region`.
func (s *API) ListSecrets(req *ListSecretsRequest, opts ...scw.RequestOption) (*ListSecretsResponse, error) {
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "is_managed", req.IsManaged)
	parameter.AddToQuery(query, "path", req.Path)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager-admin/v1/regions/" + fmt.Sprint(req.Region) + "/secrets",
		Query:  query,
	}

	var resp ListSecretsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSecretVersion: Retrieve the metadata of a secret's given version specified by the `region`, `secret_id` and `revision` parameters.
func (s *API) GetSecretVersion(req *GetSecretVersionRequest, opts ...scw.RequestOption) (*GetSecretVersionResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	if fmt.Sprint(req.Revision) == "" {
		return nil, errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager-admin/v1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "",
	}

	var resp GetSecretVersionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSecretVersions: Retrieve the list of a given secret's versions specified by the `secret_id` and `region` parameters.
func (s *API) ListSecretVersions(req *ListSecretVersionsRequest, opts ...scw.RequestOption) (*ListSecretVersionsResponse, error) {
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
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return nil, errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager-admin/v1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions",
		Query:  query,
	}

	var resp ListSecretVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

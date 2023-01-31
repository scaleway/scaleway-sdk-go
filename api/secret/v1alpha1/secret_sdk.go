// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package secret provides methods and message types of the secret v1alpha1 API.
package secret

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
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

// API: this API allows you to conveniently store, access and share sensitive data
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type ListSecretsRequestOrderBy string

const (
	// ListSecretsRequestOrderByNameAsc is [insert doc].
	ListSecretsRequestOrderByNameAsc = ListSecretsRequestOrderBy("name_asc")
	// ListSecretsRequestOrderByNameDesc is [insert doc].
	ListSecretsRequestOrderByNameDesc = ListSecretsRequestOrderBy("name_desc")
	// ListSecretsRequestOrderByCreatedAtAsc is [insert doc].
	ListSecretsRequestOrderByCreatedAtAsc = ListSecretsRequestOrderBy("created_at_asc")
	// ListSecretsRequestOrderByCreatedAtDesc is [insert doc].
	ListSecretsRequestOrderByCreatedAtDesc = ListSecretsRequestOrderBy("created_at_desc")
	// ListSecretsRequestOrderByUpdatedAtAsc is [insert doc].
	ListSecretsRequestOrderByUpdatedAtAsc = ListSecretsRequestOrderBy("updated_at_asc")
	// ListSecretsRequestOrderByUpdatedAtDesc is [insert doc].
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

type SecretStatus string

const (
	// SecretStatusReady is [insert doc].
	SecretStatusReady = SecretStatus("ready")
	// SecretStatusLocked is [insert doc].
	SecretStatusLocked = SecretStatus("locked")
)

func (enum SecretStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "ready"
	}
	return string(enum)
}

func (enum SecretStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecretStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecretStatus(SecretStatus(tmp).String())
	return nil
}

type SecretVersionStatus string

const (
	// SecretVersionStatusUnknown is [insert doc].
	SecretVersionStatusUnknown = SecretVersionStatus("unknown")
	// SecretVersionStatusEnabled is [insert doc].
	SecretVersionStatusEnabled = SecretVersionStatus("enabled")
	// SecretVersionStatusDisabled is [insert doc].
	SecretVersionStatusDisabled = SecretVersionStatus("disabled")
	// SecretVersionStatusDestroyed is [insert doc].
	SecretVersionStatusDestroyed = SecretVersionStatus("destroyed")
)

func (enum SecretVersionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum SecretVersionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecretVersionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecretVersionStatus(SecretVersionStatus(tmp).String())
	return nil
}

// AccessSecretVersionResponse: access secret version response
type AccessSecretVersionResponse struct {
	// SecretID: ID of the Secret
	SecretID string `json:"secret_id"`
	// Revision: revision of the SecretVersion
	Revision uint32 `json:"revision"`
	// Data: the base64-encoded secret payload of the SecretVersion
	Data []byte `json:"data"`
}

// ListSecretVersionsResponse: list secret versions response
type ListSecretVersionsResponse struct {
	// TotalCount: count of all SecretVersions
	TotalCount uint32 `json:"total_count"`
	// Versions: single page of SecretVersions
	Versions []*SecretVersion `json:"versions"`
}

// ListSecretsResponse: list secrets response
type ListSecretsResponse struct {
	// TotalCount: count of all Secrets matching the requested criteria
	TotalCount uint32 `json:"total_count"`
	// Secrets: single page of Secrets matching the requested criteria
	Secrets []*Secret `json:"secrets"`
}

// Secret: secret
type Secret struct {
	// ID: ID of the Secret
	ID string `json:"id"`
	// ProjectID: ID of the project containing the Secret
	ProjectID string `json:"project_id"`
	// Name: name of the Secret
	Name string `json:"name"`
	// Status: current status of the Secret
	//
	// * `ready`: the Secret is ready.
	// * `locked`: the Secret is locked.
	//
	// Default value: ready
	Status SecretStatus `json:"status"`
	// CreatedAt: the time at which the Secret was created
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: the time at which the Secret was updated
	UpdatedAt *time.Time `json:"updated_at"`
	// Tags: list of tags associated to this Secret
	Tags []string `json:"tags"`
	// Region: region of the Secret
	Region scw.Region `json:"region"`
	// VersionCount: the number of versions for this Secret
	VersionCount uint32 `json:"version_count"`
	// Description: description of the Secret
	Description *string `json:"description"`
}

// SecretVersion: secret version
type SecretVersion struct {
	// SecretID: ID of the Secret
	SecretID string `json:"secret_id"`
	// Revision: revision of the SecretVersion
	Revision uint32 `json:"revision"`
	// Status: current status of the SecretVersion
	//
	// * `unknown`: the SecretVersion is in an invalid state.
	// * `enabled`: the SecretVersion is accessible.
	// * `disabled`: the SecretVersion is not accessible but can be enabled.
	// * `destroyed`: the SecretVersion is permanently destroyed.
	//
	// Default value: unknown
	Status SecretVersionStatus `json:"status"`
	// CreatedAt: the time at which the SecretVersion was created
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: the time at which the SecretVersion was updated
	UpdatedAt *time.Time `json:"updated_at"`
	// Description: description of the SecretVersion
	Description *string `json:"description"`
}

// Service API

// Regions list localities the api is available in
func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar}
}

type CreateSecretRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// ProjectID: ID of the project containing the Secret
	ProjectID string `json:"project_id"`
	// Name: name of the Secret
	Name string `json:"name"`
	// Tags: list of tags associated to this Secret
	Tags []string `json:"tags"`
	// Description: description of the Secret
	Description *string `json:"description"`
}

// CreateSecret: create a Secret containing no versions
func (s *API) CreateSecret(req *CreateSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetSecretRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// SecretID: ID of the Secret
	SecretID string `json:"-"`
}

// GetSecret: get metadata of a Secret
func (s *API) GetSecret(req *GetSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
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
		Method:  "GET",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
		Headers: http.Header{},
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateSecretRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// SecretID: ID of the Secret
	SecretID string `json:"-"`
	// Name: new name of the Secret (optional)
	Name *string `json:"name"`
	// Tags: new list of tags associated to this Secret (optional)
	Tags *[]string `json:"tags"`
	// Description: description of the Secret
	Description *string `json:"description"`
}

// UpdateSecret: update metadata of a Secret
func (s *API) UpdateSecret(req *UpdateSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
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
		Method:  "PATCH",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListSecretsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// OrganizationID: ID of an organization to filter on (optional)
	OrganizationID *string `json:"-"`
	// ProjectID: ID of a project to filter on (optional)
	ProjectID *string `json:"-"`
	// Tags: list of tags to filter on (optional)
	Tags []string `json:"-"`
	// OrderBy:
	//
	// Default value: name_asc
	OrderBy ListSecretsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListSecrets: list Secrets
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
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListSecretsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteSecretRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// SecretID: ID of the Secret
	SecretID string `json:"-"`
}

// DeleteSecret: delete a secret
func (s *API) DeleteSecret(req *DeleteSecretRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretID) == "" {
		return errors.New("field SecretID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type CreateSecretVersionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// SecretID: ID of the Secret
	SecretID string `json:"-"`
	// Data: the base64-encoded secret payload of the SecretVersion
	Data []byte `json:"data"`
	// Description: description of the SecretVersion
	Description *string `json:"description"`
}

// CreateSecretVersion: create a SecretVersion
func (s *API) CreateSecretVersion(req *CreateSecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
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
		Method:  "POST",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetSecretVersionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// SecretID: ID of the Secret
	SecretID string `json:"-"`
	// Revision: revision of the SecretVersion (may be a number or "latest")
	Revision string `json:"-"`
}

// GetSecretVersion: get metadata of a SecretVersion
func (s *API) GetSecretVersion(req *GetSecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
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
		Method:  "GET",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "",
		Headers: http.Header{},
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateSecretVersionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// SecretID: ID of the Secret
	SecretID string `json:"-"`
	// Revision: revision of the SecretVersion (may be a number or "latest")
	Revision string `json:"-"`
	// Description: description of the SecretVersion
	Description *string `json:"description"`
}

// UpdateSecretVersion: update metadata of a SecretVersion
func (s *API) UpdateSecretVersion(req *UpdateSecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
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
		Method:  "PATCH",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListSecretVersionsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// SecretID: ID of the Secret
	SecretID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// Status: filter results by status
	Status []SecretVersionStatus `json:"-"`
}

// ListSecretVersions: list versions of a secret, not returning any sensitive data
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
		Method:  "GET",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListSecretVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DestroySecretVersionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// SecretID: ID of the Secret
	SecretID string `json:"-"`
	// Revision: revision of the SecretVersion (may be a number or "latest")
	Revision string `json:"-"`
}

// DestroySecretVersion: destroy a SecretVersion, permanently destroying the sensitive data
func (s *API) DestroySecretVersion(req *DestroySecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
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
		Method:  "POST",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/destroy",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type EnableSecretVersionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// SecretID: ID of the Secret
	SecretID string `json:"-"`
	// Revision: revision of the SecretVersion (may be a number or "latest")
	Revision string `json:"-"`
}

// EnableSecretVersion: enable a SecretVersion
func (s *API) EnableSecretVersion(req *EnableSecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
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
		Method:  "POST",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/enable",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DisableSecretVersionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// SecretID: ID of the Secret
	SecretID string `json:"-"`
	// Revision: revision of the SecretVersion (may be a number or "latest")
	Revision string `json:"-"`
}

// DisableSecretVersion: disable a SecretVersion
func (s *API) DisableSecretVersion(req *DisableSecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
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
		Method:  "POST",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/disable",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type AccessSecretVersionRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// SecretID: ID of the Secret
	SecretID string `json:"-"`
	// Revision: revision of the SecretVersion (may be a number or "latest")
	Revision string `json:"-"`
}

// AccessSecretVersion: access a SecretVersion, returning the sensitive data
func (s *API) AccessSecretVersion(req *AccessSecretVersionRequest, opts ...scw.RequestOption) (*AccessSecretVersionResponse, error) {
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
		Method:  "GET",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/access",
		Headers: http.Header{},
	}

	var resp AccessSecretVersionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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

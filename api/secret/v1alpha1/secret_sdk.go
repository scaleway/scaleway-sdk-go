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

// API: this API allows you to conveniently store, access and share sensitive data.
// Secret Manager API documentation.
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

type SecretStatus string

const (
	SecretStatusReady  = SecretStatus("ready")
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
	SecretVersionStatusUnknown   = SecretVersionStatus("unknown")
	SecretVersionStatusEnabled   = SecretVersionStatus("enabled")
	SecretVersionStatusDisabled  = SecretVersionStatus("disabled")
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

// AccessSecretVersionResponse: access secret version response.
type AccessSecretVersionResponse struct {
	// SecretID: ID of the secret.
	SecretID string `json:"secret_id"`
	// Revision: version number.
	// The first version of the secret is numbered 1, and all subsequent revisions augment by 1.
	Revision uint32 `json:"revision"`
	// Data: the base64-encoded secret payload of the version.
	Data []byte `json:"data"`
	// DataCrc32: the CRC32 checksum of the data as a base-10 integer.
	// This field is present only if a CRC32 was supplied during the creation of the version.
	DataCrc32 *uint32 `json:"data_crc32"`
}

// ListSecretVersionsResponse: list secret versions response.
type ListSecretVersionsResponse struct {
	// Versions: single page of versions.
	Versions []*SecretVersion `json:"versions"`
	// TotalCount: number of versions.
	TotalCount uint32 `json:"total_count"`
}

// ListSecretsResponse: list secrets response.
type ListSecretsResponse struct {
	// Secrets: single page of secrets matching the requested criteria.
	Secrets []*Secret `json:"secrets"`
	// TotalCount: count of all secrets matching the requested criteria.
	TotalCount uint32 `json:"total_count"`
}

// PasswordGenerationParams: password generation params.
type PasswordGenerationParams struct {
	// Length: length of the password to generate (between 1 and 1024).
	Length uint32 `json:"length"`
	// NoLowercaseLetters: do not include lower case letters by default in the alphabet.
	NoLowercaseLetters bool `json:"no_lowercase_letters"`
	// NoUppercaseLetters: do not include upper case letters by default in the alphabet.
	NoUppercaseLetters bool `json:"no_uppercase_letters"`
	// NoDigits: do not include digits by default in the alphabet.
	NoDigits bool `json:"no_digits"`
	// AdditionalChars: additional ascii characters to be included in the alphabet.
	AdditionalChars string `json:"additional_chars"`
}

// Secret: secret.
type Secret struct {
	// ID: ID of the secret.
	ID string `json:"id"`
	// ProjectID: ID of the Project containing the secret.
	ProjectID string `json:"project_id"`
	// Name: name of the secret.
	Name string `json:"name"`
	// Status: current status of the secret.
	// * `ready`: the secret is ready.
	// * `locked`: the secret is locked.
	// Default value: ready
	Status SecretStatus `json:"status"`
	// CreatedAt: date and time of the secret's creation.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: last update of the secret.
	UpdatedAt *time.Time `json:"updated_at"`
	// Tags: list of the secret's tags.
	Tags []string `json:"tags"`
	// VersionCount: number of versions for this secret.
	VersionCount uint32 `json:"version_count"`
	// Description: updated description of the secret.
	Description *string `json:"description"`
	// IsManaged: true for secrets that are managed by another product.
	IsManaged bool `json:"is_managed"`
	// Region: region of the secret.
	Region scw.Region `json:"region"`
}

// SecretVersion: secret version.
type SecretVersion struct {
	// Revision: version number.
	// The first version of the secret is numbered 1, and all subsequent revisions augment by 1.
	Revision uint32 `json:"revision"`
	// SecretID: ID of the secret.
	SecretID string `json:"secret_id"`
	// Status: current status of the version.
	// * `unknown`: the version is in an invalid state.
	// * `enabled`: the version is accessible.
	// * `disabled`: the version is not accessible but can be enabled.
	// * `destroyed`: the version is permanently deleted. It is not possible to recover it.
	// Default value: unknown
	Status SecretVersionStatus `json:"status"`
	// CreatedAt: date and time of the version's creation.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: last update of the version.
	UpdatedAt *time.Time `json:"updated_at"`
	// Description: description of the version.
	Description *string `json:"description"`
	// IsLatest: true if the version is the latest one.
	IsLatest bool `json:"is_latest"`
}

// Service API

// Regions list localities the api is available in
func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar}
}

type CreateSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// ProjectID: ID of the Project containing the secret.
	ProjectID string `json:"project_id"`
	// Name: name of the secret.
	Name string `json:"name"`
	// Tags: list of the secret's tags.
	Tags []string `json:"tags"`
	// Description: description of the secret.
	Description *string `json:"description"`
}

// CreateSecret: create a secret.
// You must sepcify the `region` to create a secret.
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
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
}

// GetSecret: get metadata using the secret's name.
// Retrieve the metadata of a secret specified by the `region` and the `secret_name` parameters.
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

type GetSecretByNameRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretName: name of the secret.
	SecretName string `json:"-"`
}

// GetSecretByName: get metadata using the secret's ID.
// Retrieve the metadata of a secret specified by the `region` and the `secret_id` parameters.
func (s *API) GetSecretByName(req *GetSecretByNameRequest, opts ...scw.RequestOption) (*Secret, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretName) == "" {
		return nil, errors.New("field SecretName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-name/" + fmt.Sprint(req.SecretName) + "",
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
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Name: secret's updated name (optional).
	Name *string `json:"name"`
	// Tags: secret's updated list of tags (optional).
	Tags *[]string `json:"tags"`
	// Description: description of the secret.
	Description *string `json:"description"`
}

// UpdateSecret: update metadata of a secret.
// Edit a secret's metadata such as name, tag(s) and description. The secret to update is specified by the `secret_id` and `region` parameters.
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

type AddSecretOwnerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// ProductName: name of the product to add.
	ProductName string `json:"product_name"`
}

// AddSecretOwner: allow another product to use the secret.
func (s *API) AddSecretOwner(req *AddSecretOwnerRequest, opts ...scw.RequestOption) error {
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
		Method:  "POST",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/add-owner",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return err
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListSecretsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// OrganizationID: filter by Organization ID (optional).
	OrganizationID *string `json:"-"`
	// ProjectID: filter by Project ID (optional).
	ProjectID *string `json:"-"`
	// Name: filter by secret name (optional).
	Name *string `json:"-"`
	// Tags: list of tags to filter on (optional).
	Tags []string `json:"-"`
	// IsManaged: filter by managed / not managed (optional).
	IsManaged *bool `json:"-"`
	// OrderBy: default value: name_asc
	OrderBy ListSecretsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListSecrets: list secrets.
// Retrieve the list of secrets created within an Organization and/or Project. You must specify either the `organization_id` or the `project_id` and the `region`.
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
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
}

// DeleteSecret: delete a secret.
// Delete a given secret specified by the `region` and `secret_id` parameters.
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
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Data: the base64-encoded secret payload of the version.
	Data []byte `json:"data"`
	// Description: description of the version.
	Description *string `json:"description"`
	// DisablePrevious: disable the previous secret version.
	// Optional. If there is no previous version or if the previous version was already disabled, does nothing.
	DisablePrevious *bool `json:"disable_previous"`
	// PasswordGeneration: options to generate a password.
	// Optional. If specified, a random password will be generated. The data and data_crc32 fields must be empty. By default, the generator will use upper and lower case letters, and digits. This behavior can be tuned using the generation parameters.
	// Precisely one of PasswordGeneration must be set.
	PasswordGeneration *PasswordGenerationParams `json:"password_generation,omitempty"`
	// DataCrc32: the CRC32 checksum of the data as a base-10 integer.
	// Optional. If specified, the Secret Manager will verify the integrity of the data received against the given CRC32. An error is returned if the CRC32 does not match. Otherwise, the CRC32 will be stored and returned along with the SecretVersion on futur accesses.
	DataCrc32 *uint32 `json:"data_crc32"`
}

// CreateSecretVersion: create a version.
// Create a version of a given secret specified by the `region` and `secret_id` parameters.
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
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Revision: version number.
	// The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

// GetSecretVersion: get metadata of a secret's version using the secret's ID.
// Retrieve the metadata of a secret's given version specified by the `region`, `secret_id` and `revision` parameters.
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

type GetSecretVersionByNameRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretName: name of the secret.
	SecretName string `json:"-"`
	// Revision: version number.
	// The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

// GetSecretVersionByName: get metadata of a secret's version using the secret's name.
// Retrieve the metadata of a secret's given version specified by the `region`, `secret_name` and `revision` parameters.
func (s *API) GetSecretVersionByName(req *GetSecretVersionByNameRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretName) == "" {
		return nil, errors.New("field SecretName cannot be empty in request")
	}

	if fmt.Sprint(req.Revision) == "" {
		return nil, errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-name/" + fmt.Sprint(req.SecretName) + "/versions/" + fmt.Sprint(req.Revision) + "",
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
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Revision: version number.
	// The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
	// Description: description of the version.
	Description *string `json:"description"`
}

// UpdateSecretVersion: update metadata of a version.
// Edit the metadata of a secret's given version, specified by the `region`, `secret_id` and `revision` parameters.
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
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// Status: filter results by status.
	Status []SecretVersionStatus `json:"-"`
}

// ListSecretVersions: list versions of a secret using the secret's ID.
// Retrieve the list of a given secret's versions specified by the `secret_id` and `region` parameters.
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

type ListSecretVersionsByNameRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretName: name of the secret.
	SecretName string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// Status: filter results by status.
	Status []SecretVersionStatus `json:"-"`
}

// ListSecretVersionsByName: list versions of a secret using the secret's name.
// Retrieve the list of a given secret's versions specified by the `secret_name` and `region` parameters.
func (s *API) ListSecretVersionsByName(req *ListSecretVersionsByNameRequest, opts ...scw.RequestOption) (*ListSecretVersionsResponse, error) {
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

	if fmt.Sprint(req.SecretName) == "" {
		return nil, errors.New("field SecretName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-name/" + fmt.Sprint(req.SecretName) + "/versions",
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

type EnableSecretVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Revision: version number.
	// The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

// EnableSecretVersion: enable a version.
// Make a specific version accessible. You must specify the `region`, `secret_id` and `revision` parameters.
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
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Revision: version number.
	// The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

// DisableSecretVersion: disable a version.
// Make a specific version inaccessible. You must specify the `region`, `secret_id` and `revision` parameters.
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
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Revision: version number.
	// The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

// AccessSecretVersion: access a secret's version using the secret's ID.
// Access sensitive data in a secret's version specified by the `region`, `secret_id` and `revision` parameters.
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

type AccessSecretVersionByNameRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretName: name of the secret.
	SecretName string `json:"-"`
	// Revision: version number.
	// The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

// AccessSecretVersionByName: access a secret's version using the secret's name.
// Access sensitive data in a secret's version specified by the `region`, `secret_name` and `revision` parameters.
func (s *API) AccessSecretVersionByName(req *AccessSecretVersionByNameRequest, opts ...scw.RequestOption) (*AccessSecretVersionResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretName) == "" {
		return nil, errors.New("field SecretName cannot be empty in request")
	}

	if fmt.Sprint(req.Revision) == "" {
		return nil, errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-name/" + fmt.Sprint(req.SecretName) + "/versions/" + fmt.Sprint(req.Revision) + "/access",
		Headers: http.Header{},
	}

	var resp AccessSecretVersionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DestroySecretVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
	// SecretID: ID of the secret.
	SecretID string `json:"-"`
	// Revision: version number.
	// The first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be a number or "latest".
	Revision string `json:"-"`
}

// DestroySecretVersion: delete a version.
// Delete a secret's version and the sensitive data contained in it. Deleting a version is permanent and cannot be undone.
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

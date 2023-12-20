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

type EphemeralPolicyAction string

const (
	EphemeralPolicyActionUnknownAction = EphemeralPolicyAction("unknown_action")
	// The version is deleted once it expires.
	EphemeralPolicyActionDelete = EphemeralPolicyAction("delete")
	// The version is disabled once it expires.
	EphemeralPolicyActionDisable = EphemeralPolicyAction("disable")
)

func (enum EphemeralPolicyAction) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_action"
	}
	return string(enum)
}

func (enum EphemeralPolicyAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EphemeralPolicyAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EphemeralPolicyAction(EphemeralPolicyAction(tmp).String())
	return nil
}

type ListFoldersRequestOrderBy string

const (
	ListFoldersRequestOrderByCreatedAtAsc  = ListFoldersRequestOrderBy("created_at_asc")
	ListFoldersRequestOrderByCreatedAtDesc = ListFoldersRequestOrderBy("created_at_desc")
	ListFoldersRequestOrderByNameAsc       = ListFoldersRequestOrderBy("name_asc")
	ListFoldersRequestOrderByNameDesc      = ListFoldersRequestOrderBy("name_desc")
)

func (enum ListFoldersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListFoldersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFoldersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFoldersRequestOrderBy(ListFoldersRequestOrderBy(tmp).String())
	return nil
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

type Product string

const (
	ProductUnknown      = Product("unknown")
	ProductEdgeServices = Product("edge_services")
)

func (enum Product) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum Product) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Product) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Product(Product(tmp).String())
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

type SecretType string

const (
	SecretTypeUnknownSecretType = SecretType("unknown_secret_type")
	// Default type.
	SecretTypeOpaque = SecretType("opaque")
	// List of concatenated PEM blocks. They can contain certificates, private keys or any other PEM block types.
	SecretTypeCertificate = SecretType("certificate")
	// Flat JSON that allows you to set any number of first level key and a scalar type as a value (string, numeric, boolean).
	SecretTypeKeyValue = SecretType("key_value")
)

func (enum SecretType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_secret_type"
	}
	return string(enum)
}

func (enum SecretType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SecretType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SecretType(SecretType(tmp).String())
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

// EphemeralProperties: ephemeral properties.
type EphemeralProperties struct {
	// ExpiresAt: (Optional.) If not specified, the version does not have an expiration date.
	ExpiresAt *time.Time `json:"expires_at"`

	// ExpiresOnceAccessed: (Optional.) If not specified, the version can be accessed an unlimited amount of times.
	ExpiresOnceAccessed *bool `json:"expires_once_accessed"`

	// Action: see `EphemeralPolicy.Action` enum for a description of values.
	// Default value: unknown_action
	Action EphemeralPolicyAction `json:"action"`
}

// EphemeralPolicy: ephemeral policy.
type EphemeralPolicy struct {
	// TimeToLive: time frame, from one second and up to one year, during which the secret's versions are valid.
	TimeToLive *scw.Duration `json:"time_to_live"`

	// ExpiresOnceAccessed: returns `true` if the version expires after a single user access.
	ExpiresOnceAccessed *bool `json:"expires_once_accessed"`

	// Action: see the `EphemeralPolicy.Action` enum for a description of values.
	// Default value: unknown_action
	Action EphemeralPolicyAction `json:"action"`
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

// Folder: folder.
type Folder struct {
	// ID: ID of the folder.
	ID string `json:"id"`

	// ProjectID: ID of the Project containing the folder.
	ProjectID string `json:"project_id"`

	// Name: name of the folder.
	Name string `json:"name"`

	// Path: location of the folder in the directory structure.
	Path string `json:"path"`

	// CreatedAt: date and time of the folder's creation.
	CreatedAt *time.Time `json:"created_at"`

	// Region: region of the folder.
	Region scw.Region `json:"region"`
}

// SecretVersion: secret version.
type SecretVersion struct {
	// Revision: the first version of the secret is numbered 1, and all subsequent revisions augment by 1.
	Revision uint32 `json:"revision"`

	// SecretID: ID of the secret.
	SecretID string `json:"secret_id"`

	// Status: * `unknown`: the version is in an invalid state.
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

	// IsLatest: returns `true` if the version is the latest.
	IsLatest bool `json:"is_latest"`

	// EphemeralProperties: returns the version's expiration date, whether it expires after being accessed once, and the action to perform (disable or delete) once the version expires.
	EphemeralProperties *EphemeralProperties `json:"ephemeral_properties"`
}

// Secret: secret.
type Secret struct {
	// ID: ID of the secret.
	ID string `json:"id"`

	// ProjectID: ID of the Project containing the secret.
	ProjectID string `json:"project_id"`

	// Name: name of the secret.
	Name string `json:"name"`

	// Status: * `ready`: the secret can be read, modified and deleted.
	// * `locked`: no action can be performed on the secret. This status can only be applied and removed by Scaleway.
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

	// IsManaged: returns `true` for secrets that are managed by another product.
	IsManaged bool `json:"is_managed"`

	// IsProtected: returns `true` for protected secrets that cannot be deleted.
	IsProtected bool `json:"is_protected"`

	// Type: see `Secret.Type` enum for description of values.
	// Default value: unknown_secret_type
	Type SecretType `json:"type"`

	// Path: location of the secret in the directory structure.
	Path string `json:"path"`

	// EphemeralPolicy: (Optional.) Policy that defines whether/when a secret's versions expire. By default, the policy is applied to all the secret's versions.
	EphemeralPolicy *EphemeralPolicy `json:"ephemeral_policy"`

	// Region: region of the secret.
	Region scw.Region `json:"region"`
}

// AccessSecretVersionByNameRequest: access secret version by name request.
type AccessSecretVersionByNameRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretName: name of the secret.
	SecretName string `json:"-"`

	// Revision: the first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be either:
	// - a number (the revision number)
	// - "latest" (the latest revision)
	// - "latest_enabled" (the latest enabled revision).
	Revision string `json:"-"`

	// ProjectID: (Optional.) If not specified, Secret Manager will look for the secret version in all Projects.
	ProjectID *string `json:"project_id,omitempty"`
}

// AccessSecretVersionRequest: access secret version request.
type AccessSecretVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	// Revision: the first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be either:
	// - a number (the revision number)
	// - "latest" (the latest revision)
	// - "latest_enabled" (the latest enabled revision).
	Revision string `json:"-"`
}

// AccessSecretVersionResponse: access secret version response.
type AccessSecretVersionResponse struct {
	// SecretID: ID of the secret.
	SecretID string `json:"secret_id"`

	// Revision: the first version of the secret is numbered 1, and all subsequent revisions augment by 1.
	Revision uint32 `json:"revision"`

	// Data: the base64-encoded secret payload of the version.
	Data []byte `json:"data"`

	// DataCrc32: this field is only available if a CRC32 was supplied during the creation of the version.
	DataCrc32 *uint32 `json:"data_crc32"`
}

// AddSecretOwnerRequest: add secret owner request.
type AddSecretOwnerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	// Deprecated: ProductName: (Deprecated: use `product` field) Name of the product to add.
	ProductName *string `json:"product_name,omitempty"`

	// Product: see `Product` enum for description of values.
	// Default value: unknown
	Product Product `json:"product"`
}

// CreateFolderRequest: create folder request.
type CreateFolderRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project containing the folder.
	ProjectID string `json:"project_id"`

	// Name: name of the folder.
	Name string `json:"name"`

	// Path: (Optional.) Location of the folder in the directory structure. If not specified, the path is `/`.
	Path *string `json:"path,omitempty"`
}

// CreateSecretRequest: create secret request.
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
	Description *string `json:"description,omitempty"`

	// Type: (Optional.) See `Secret.Type` enum for description of values. If not specified, the type is `Opaque`.
	// Default value: unknown_secret_type
	Type SecretType `json:"type"`

	// Path: (Optional.) Location of the secret in the directory structure. If not specified, the path is `/`.
	Path *string `json:"path,omitempty"`

	// EphemeralPolicy: (Optional.) Policy that defines whether/when a secret's versions expire. By default, the policy is applied to all the secret's versions.
	EphemeralPolicy *EphemeralPolicy `json:"ephemeral_policy,omitempty"`
}

// CreateSecretVersionRequest: create secret version request.
type CreateSecretVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	// Data: the base64-encoded secret payload of the version.
	Data []byte `json:"data"`

	// Description: description of the version.
	Description *string `json:"description,omitempty"`

	// DisablePrevious: (Optional.) If there is no previous version or if the previous version was already disabled, does nothing.
	DisablePrevious *bool `json:"disable_previous,omitempty"`

	// Deprecated: PasswordGeneration: (Optional.) If specified, a random password will be generated. The `data` and `data_crc32` fields must be empty. By default, the generator will use upper and lower case letters, and digits. This behavior can be tuned using the generation parameters.
	PasswordGeneration *PasswordGenerationParams `json:"password_generation,omitempty"`

	// DataCrc32: if specified, Secret Manager will verify the integrity of the data received against the given CRC32 checksum. An error is returned if the CRC32 does not match. If, however, the CRC32 matches, it will be stored and returned along with the SecretVersion on future access requests.
	DataCrc32 *uint32 `json:"data_crc32,omitempty"`
}

// DeleteFolderRequest: delete folder request.
type DeleteFolderRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FolderID: ID of the folder.
	FolderID string `json:"-"`
}

// DeleteSecretRequest: delete secret request.
type DeleteSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`
}

// DestroySecretVersionRequest: destroy secret version request.
type DestroySecretVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	// Revision: the first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be either:
	// - a number (the revision number)
	// - "latest" (the latest revision)
	// - "latest_enabled" (the latest enabled revision).
	Revision string `json:"-"`
}

// DisableSecretVersionRequest: disable secret version request.
type DisableSecretVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	// Revision: the first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be either:
	// - a number (the revision number)
	// - "latest" (the latest revision)
	// - "latest_enabled" (the latest enabled revision).
	Revision string `json:"-"`
}

// EnableSecretVersionRequest: enable secret version request.
type EnableSecretVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	// Revision: the first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be either:
	// - a number (the revision number)
	// - "latest" (the latest revision)
	// - "latest_enabled" (the latest enabled revision).
	Revision string `json:"-"`
}

// GeneratePasswordRequest: generate password request.
type GeneratePasswordRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	// Description: description of the version.
	Description *string `json:"description,omitempty"`

	// DisablePrevious: this has no effect if there is no previous version or if the previous version was already disabled.
	DisablePrevious *bool `json:"disable_previous,omitempty"`

	// Length: length of the password to generate (between 1 and 1024 characters).
	Length uint32 `json:"length"`

	// NoLowercaseLetters: (Optional.) Exclude lower case letters by default in the password character set.
	NoLowercaseLetters *bool `json:"no_lowercase_letters,omitempty"`

	// NoUppercaseLetters: (Optional.) Exclude upper case letters by default in the password character set.
	NoUppercaseLetters *bool `json:"no_uppercase_letters,omitempty"`

	// NoDigits: (Optional.) Exclude digits by default in the password character set.
	NoDigits *bool `json:"no_digits,omitempty"`

	// AdditionalChars: (Optional.) Additional ASCII characters to be included in the password character set.
	AdditionalChars *string `json:"additional_chars,omitempty"`
}

// GetSecretByNameRequest: get secret by name request.
type GetSecretByNameRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretName: name of the secret.
	SecretName string `json:"-"`

	// ProjectID: (Optional.) If not specified, Secret Manager will look for the secret in all Projects.
	ProjectID *string `json:"project_id,omitempty"`
}

// GetSecretRequest: get secret request.
type GetSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`
}

// GetSecretVersionByNameRequest: get secret version by name request.
type GetSecretVersionByNameRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretName: name of the secret.
	SecretName string `json:"-"`

	// Revision: the first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be either:
	// - a number (the revision number)
	// - "latest" (the latest revision)
	// - "latest_enabled" (the latest enabled revision).
	Revision string `json:"-"`

	// ProjectID: (Optional.) If not specified, Secret Manager will look for the secret version in all Projects.
	ProjectID *string `json:"project_id,omitempty"`
}

// GetSecretVersionRequest: get secret version request.
type GetSecretVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	// Revision: the first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be either:
	// - a number (the revision number)
	// - "latest" (the latest revision)
	// - "latest_enabled" (the latest enabled revision).
	Revision string `json:"-"`
}

// ListFoldersRequest: list folders request.
type ListFoldersRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: filter by Project ID (optional).
	ProjectID *string `json:"-"`

	// Path: filter by path (optional).
	Path *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListFoldersRequestOrderBy `json:"-"`
}

// ListFoldersResponse: list folders response.
type ListFoldersResponse struct {
	// Folders: list of folders.
	Folders []*Folder `json:"folders"`

	// TotalCount: count of all folders matching the requested criteria.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFoldersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFoldersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListFoldersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Folders = append(r.Folders, results.Folders...)
	r.TotalCount += uint32(len(results.Folders))
	return uint32(len(results.Folders)), nil
}

// ListSecretVersionsByNameRequest: list secret versions by name request.
type ListSecretVersionsByNameRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretName: name of the secret.
	SecretName string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// Status: filter results by status.
	Status []SecretVersionStatus `json:"-"`

	// ProjectID: (Optional.) If not specified, Secret Manager will look for the secret in all Projects.
	ProjectID *string `json:"-"`
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
	Status []SecretVersionStatus `json:"-"`
}

// ListSecretVersionsResponse: list secret versions response.
type ListSecretVersionsResponse struct {
	// Versions: single page of versions.
	Versions []*SecretVersion `json:"versions"`

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

	// Path: filter by path (optional).
	Path *string `json:"-"`

	// IsEphemeral: filter by ephemeral / not ephemeral (optional).
	IsEphemeral *bool `json:"-"`
}

// ListSecretsResponse: list secrets response.
type ListSecretsResponse struct {
	// Secrets: single page of secrets matching the requested criteria.
	Secrets []*Secret `json:"secrets"`

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

// ListTagsRequest: list tags request.
type ListTagsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: (Optional.) If not specified, Secret Manager will look for tags in all Projects.
	ProjectID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListTagsResponse: list tags response.
type ListTagsResponse struct {
	// Tags: list of tags.
	Tags []string `json:"tags"`

	// TotalCount: count of all tags matching the requested criteria.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTagsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTagsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTagsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tags = append(r.Tags, results.Tags...)
	r.TotalCount += uint32(len(results.Tags))
	return uint32(len(results.Tags)), nil
}

// ProtectSecretRequest: protect secret request.
type ProtectSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret to protect.
	SecretID string `json:"-"`
}

// UnprotectSecretRequest: unprotect secret request.
type UnprotectSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret to unprotect.
	SecretID string `json:"-"`
}

// UpdateSecretRequest: update secret request.
type UpdateSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	// Name: secret's updated name (optional).
	Name *string `json:"name,omitempty"`

	// Tags: secret's updated list of tags (optional).
	Tags *[]string `json:"tags,omitempty"`

	// Description: description of the secret.
	Description *string `json:"description,omitempty"`

	// Path: (Optional.) Location of the folder in the directory structure. If not specified, the path is `/`.
	Path *string `json:"path,omitempty"`

	// EphemeralPolicy: (Optional.) Policy that defines whether/when a secret's versions expire.
	EphemeralPolicy *EphemeralPolicy `json:"ephemeral_policy,omitempty"`
}

// UpdateSecretVersionRequest: update secret version request.
type UpdateSecretVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	// Revision: the first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be either:
	// - a number (the revision number)
	// - "latest" (the latest revision)
	// - "latest_enabled" (the latest enabled revision).
	Revision string `json:"-"`

	// Description: description of the version.
	Description *string `json:"description,omitempty"`

	// EphemeralProperties: (Optional.) Properties that defines the version's expiration date, whether it expires after being accessed once, and the action to perform (disable or delete) once the version expires.
	EphemeralProperties *EphemeralProperties `json:"ephemeral_properties,omitempty"`
}

// This API allows you to conveniently store, access and share sensitive data.
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

// CreateSecret: You must specify the `region` to create a secret.
func (s *API) CreateSecret(req *CreateSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
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
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets",
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

// CreateFolder: Create folder.
func (s *API) CreateFolder(req *CreateFolderRequest, opts ...scw.RequestOption) (*Folder, error) {
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
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/folders",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Folder

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSecret: Retrieve the metadata of a secret specified by the `region` and `secret_id` parameters.
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
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: GetSecretByName: Retrieve the metadata of a secret specified by the `region` and `secret_name` parameters.
//
// GetSecretByName usage is now deprecated.
//
// Scaleway recommends that you use the `ListSecrets` request with the `name` filter.
func (s *API) GetSecretByName(req *GetSecretByNameRequest, opts ...scw.RequestOption) (*Secret, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretName) == "" {
		return nil, errors.New("field SecretName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-name/" + fmt.Sprint(req.SecretName) + "",
		Query:  query,
	}

	var resp Secret

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSecret: Edit a secret's metadata such as name, tag(s), description and ephemeral policy. The secret to update is specified by the `secret_id` and `region` parameters.
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
		Method: "PATCH",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
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
	parameter.AddToQuery(query, "is_ephemeral", req.IsEphemeral)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets",
		Query:  query,
	}

	var resp ListSecretsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFolders: Retrieve the list of folders created within a Project.
func (s *API) ListFolders(req *ListFoldersRequest, opts ...scw.RequestOption) (*ListFoldersResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "path", req.Path)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/folders",
		Query:  query,
	}

	var resp ListFoldersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSecret: Delete a given secret specified by the `region` and `secret_id` parameters.
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
		Method: "DELETE",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteFolder: Delete a given folder specified by the `region` and `folder_id` parameters.
func (s *API) DeleteFolder(req *DeleteFolderRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FolderID) == "" {
		return errors.New("field FolderID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/folders/" + fmt.Sprint(req.FolderID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ProtectSecret: Protect a given secret specified by the `secret_id` parameter. A protected secret can be read and modified but cannot be deleted.
func (s *API) ProtectSecret(req *ProtectSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
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
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/protect",
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

// UnprotectSecret: Unprotect a given secret specified by the `secret_id` parameter. An unprotected secret can be read, modified and deleted.
func (s *API) UnprotectSecret(req *UnprotectSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
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
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/unprotect",
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

// AddSecretOwner: Allow a product to use the secret.
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
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/add-owner",
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

// CreateSecretVersion: Create a version of a given secret specified by the `region` and `secret_id` parameters.
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
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions",
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

// GeneratePassword: Generate a password for the given secret specified by the `region` and `secret_id` parameters. This will also create a new version of the secret that will store the password.
func (s *API) GeneratePassword(req *GeneratePasswordRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
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
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/generate-password",
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

// GetSecretVersion: Retrieve the metadata of a secret's given version specified by the `region`, `secret_id` and `revision` parameters.
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
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "",
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: GetSecretVersionByName: Retrieve the metadata of a secret's given version specified by the `region`, `secret_name`, `revision` and `project_id` parameters.
//
// This method is deprecated.
//
// Scaleway recommends that you use the `ListSecrets` request with the `name` filter to specify the secret version desired, then use the `GetSecretVersion` request.
func (s *API) GetSecretVersionByName(req *GetSecretVersionByNameRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

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
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-name/" + fmt.Sprint(req.SecretName) + "/versions/" + fmt.Sprint(req.Revision) + "",
		Query:  query,
	}

	var resp SecretVersion

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSecretVersion: Edit the metadata of a secret's given version, specified by the `region`, `secret_id` and `revision` parameters.
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
		Method: "PATCH",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "",
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
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions",
		Query:  query,
	}

	var resp ListSecretVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: ListSecretVersionsByName: Retrieve the list of a given secret's versions specified by the `secret_name`,`region` and `project_id` parameters.
//
// This method is deprecated.
//
// Scaleway recommends that you use the `ListSecrets` request with the `name` filter to specify the secret version desired, then use the `ListSecretVersions` request.
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SecretName) == "" {
		return nil, errors.New("field SecretName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-name/" + fmt.Sprint(req.SecretName) + "/versions",
		Query:  query,
	}

	var resp ListSecretVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableSecretVersion: Make a specific version accessible. You must specify the `region`, `secret_id` and `revision` parameters.
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
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/enable",
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

// DisableSecretVersion: Make a specific version inaccessible. You must specify the `region`, `secret_id` and `revision` parameters.
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
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/disable",
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

// AccessSecretVersion: Access sensitive data in a secret's version specified by the `region`, `secret_id` and `revision` parameters.
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
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/access",
	}

	var resp AccessSecretVersionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: AccessSecretVersionByName: Access sensitive data in a secret's version specified by the `region`, `secret_name`, `revision` and `project_id` parameters.
//
// This method is deprecated.
//
// Scaleway recommends that you use the `ListSecrets` request with the `name` filter to specify the secret version desired, then use the `AccessSecretVersion` request.
func (s *API) AccessSecretVersionByName(req *AccessSecretVersionByNameRequest, opts ...scw.RequestOption) (*AccessSecretVersionResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

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
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-name/" + fmt.Sprint(req.SecretName) + "/versions/" + fmt.Sprint(req.Revision) + "/access",
		Query:  query,
	}

	var resp AccessSecretVersionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DestroySecretVersion: Delete a secret's version and the sensitive data contained in it. Deleting a version is permanent and cannot be undone.
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
		Method: "POST",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/destroy",
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

// ListTags: List all tags associated with secrets within a given Project.
func (s *API) ListTags(req *ListTagsRequest, opts ...scw.RequestOption) (*ListTagsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/tags",
		Query:  query,
	}

	var resp ListTagsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

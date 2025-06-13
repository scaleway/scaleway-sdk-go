// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package secret provides methods and message types of the secret v1beta1 API.
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

type BrowseSecretsRequestOrderBy string

const (
	BrowseSecretsRequestOrderByNameAsc       = BrowseSecretsRequestOrderBy("name_asc")
	BrowseSecretsRequestOrderByNameDesc      = BrowseSecretsRequestOrderBy("name_desc")
	BrowseSecretsRequestOrderByCreatedAtAsc  = BrowseSecretsRequestOrderBy("created_at_asc")
	BrowseSecretsRequestOrderByCreatedAtDesc = BrowseSecretsRequestOrderBy("created_at_desc")
	BrowseSecretsRequestOrderByUpdatedAtAsc  = BrowseSecretsRequestOrderBy("updated_at_asc")
	BrowseSecretsRequestOrderByUpdatedAtDesc = BrowseSecretsRequestOrderBy("updated_at_desc")
)

func (enum BrowseSecretsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(BrowseSecretsRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum BrowseSecretsRequestOrderBy) Values() []BrowseSecretsRequestOrderBy {
	return []BrowseSecretsRequestOrderBy{
		"name_asc",
		"name_desc",
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
	}
}

func (enum BrowseSecretsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BrowseSecretsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BrowseSecretsRequestOrderBy(BrowseSecretsRequestOrderBy(tmp).String())
	return nil
}

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
		return string(EphemeralPolicyActionUnknownAction)
	}
	return string(enum)
}

func (enum EphemeralPolicyAction) Values() []EphemeralPolicyAction {
	return []EphemeralPolicyAction{
		"unknown_action",
		"delete",
		"disable",
	}
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
		return string(ListSecretsRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListSecretsRequestOrderBy) Values() []ListSecretsRequestOrderBy {
	return []ListSecretsRequestOrderBy{
		"name_asc",
		"name_desc",
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
	}
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
	ProductUnknownProduct = Product("unknown_product")
	ProductEdgeServices   = Product("edge_services")
)

func (enum Product) String() string {
	if enum == "" {
		// return default value if empty
		return string(ProductUnknownProduct)
	}
	return string(enum)
}

func (enum Product) Values() []Product {
	return []Product{
		"unknown_product",
		"edge_services",
	}
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
	SecretStatusUnknownStatus = SecretStatus("unknown_status")
	SecretStatusReady         = SecretStatus("ready")
	SecretStatusLocked        = SecretStatus("locked")
)

func (enum SecretStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(SecretStatusUnknownStatus)
	}
	return string(enum)
}

func (enum SecretStatus) Values() []SecretStatus {
	return []SecretStatus{
		"unknown_status",
		"ready",
		"locked",
	}
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
	SecretTypeUnknownType = SecretType("unknown_type")
	// Default type.
	SecretTypeOpaque = SecretType("opaque")
	// List of concatenated PEM blocks. They can contain certificates, private keys or any other PEM block types.
	SecretTypeCertificate = SecretType("certificate")
	// Flat JSON that allows you to set as many first level keys and scalar types as values (string, numeric, boolean) as you need.
	SecretTypeKeyValue = SecretType("key_value")
	// Flat JSON that allows you to set a username and a password.
	SecretTypeBasicCredentials = SecretType("basic_credentials")
	// Flat JSON that allows you to set an engine, username, password, host, database name, and port.
	SecretTypeDatabaseCredentials = SecretType("database_credentials")
	// Flat JSON that allows you to set an SSH key.
	SecretTypeSSHKey = SecretType("ssh_key")
)

func (enum SecretType) String() string {
	if enum == "" {
		// return default value if empty
		return string(SecretTypeUnknownType)
	}
	return string(enum)
}

func (enum SecretType) Values() []SecretType {
	return []SecretType{
		"unknown_type",
		"opaque",
		"certificate",
		"key_value",
		"basic_credentials",
		"database_credentials",
		"ssh_key",
	}
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
	SecretVersionStatusUnknownStatus        = SecretVersionStatus("unknown_status")
	SecretVersionStatusEnabled              = SecretVersionStatus("enabled")
	SecretVersionStatusDisabled             = SecretVersionStatus("disabled")
	SecretVersionStatusDeleted              = SecretVersionStatus("deleted")
	SecretVersionStatusScheduledForDeletion = SecretVersionStatus("scheduled_for_deletion")
)

func (enum SecretVersionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(SecretVersionStatusUnknownStatus)
	}
	return string(enum)
}

func (enum SecretVersionStatus) Values() []SecretVersionStatus {
	return []SecretVersionStatus{
		"unknown_status",
		"enabled",
		"disabled",
		"deleted",
		"scheduled_for_deletion",
	}
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

// BrowseSecretsResponseItemFolderDetails: browse secrets response item folder details.
type BrowseSecretsResponseItemFolderDetails struct{}

// BrowseSecretsResponseItemSecretDetails: browse secrets response item secret details.
type BrowseSecretsResponseItemSecretDetails struct {
	ID string `json:"id"`

	Tags []string `json:"tags"`

	VersionCount uint32 `json:"version_count"`

	Protected bool `json:"protected"`

	EphemeralPolicy *EphemeralPolicy `json:"ephemeral_policy"`

	// Type: default value: unknown_type
	Type SecretType `json:"type"`
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

// BrowseSecretsResponseItem: browse secrets response item.
type BrowseSecretsResponseItem struct {
	Name string `json:"name"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	// Precisely one of Secret, Folder must be set.
	Secret *BrowseSecretsResponseItemSecretDetails `json:"secret,omitempty"`

	// Precisely one of Secret, Folder must be set.
	Folder *BrowseSecretsResponseItemFolderDetails `json:"folder,omitempty"`
}

// SecretVersion: secret version.
type SecretVersion struct {
	// Revision: the first version of the secret is numbered 1, and all subsequent revisions augment by 1.
	Revision uint32 `json:"revision"`

	// SecretID: ID of the secret.
	SecretID string `json:"secret_id"`

	// Status: * `unknown_status`: the version is in an invalid state.
	// * `enabled`: the version is accessible.
	// * `disabled`: the version is not accessible but can be enabled.
	// * `scheduled_for_deletion`: the version is scheduled for deletion. It will be deleted in 7 days.
	// * `deleted`: the version is permanently deleted. It is not possible to recover it.
	// Default value: unknown_status
	Status SecretVersionStatus `json:"status"`

	// CreatedAt: date and time of the version's creation.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update of the version.
	UpdatedAt *time.Time `json:"updated_at"`

	// DeletedAt: date and time of the version's deletion.
	DeletedAt *time.Time `json:"deleted_at"`

	// Description: description of the version.
	Description *string `json:"description"`

	// Latest: returns `true` if the version is the latest.
	Latest bool `json:"latest"`

	// EphemeralProperties: returns the version's expiration date, whether it expires after being accessed once, and the action to perform (disable or delete) once the version expires.
	EphemeralProperties *EphemeralProperties `json:"ephemeral_properties"`

	// DeletionRequestedAt: returns the time at which deletion was requested.
	DeletionRequestedAt *time.Time `json:"deletion_requested_at"`
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
	// Default value: unknown_status
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

	// Managed: returns `true` for secrets that are managed by another product.
	Managed bool `json:"managed"`

	// Protected: returns `true` for protected secrets that cannot be deleted.
	Protected bool `json:"protected"`

	// Type: see the `Secret.Type` enum for a description of values.
	// Default value: unknown_type
	Type SecretType `json:"type"`

	// Path: location of the secret in the directory structure.
	Path string `json:"path"`

	// EphemeralPolicy: (Optional.) Policy that defines whether/when a secret's versions expire. By default, the policy is applied to all the secret's versions.
	EphemeralPolicy *EphemeralPolicy `json:"ephemeral_policy"`

	// UsedBy: list of Scaleway resources that can access and manage the secret.
	UsedBy []Product `json:"used_by"`

	// DeletionRequestedAt: returns the time at which deletion was requested.
	DeletionRequestedAt *time.Time `json:"deletion_requested_at"`

	// KeyID: (Optional.) The Scaleway Key Manager key ID used to encrypt and decrypt secret versions.
	KeyID *string `json:"key_id"`

	// Region: region of the secret.
	Region scw.Region `json:"region"`
}

// AccessSecretVersionByPathRequest: access secret version by path request.
type AccessSecretVersionByPathRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Revision: the first version of the secret is numbered 1, and all subsequent revisions augment by 1. Value can be either:
	// - an integer (the revision number)
	// - "latest" (the latest revision)
	// - "latest_enabled" (the latest enabled revision).
	Revision string `json:"-"`

	// SecretPath: secret's path.
	SecretPath string `json:"-"`

	// SecretName: secret's name.
	SecretName string `json:"-"`

	// ProjectID: ID of the Project to target.
	ProjectID string `json:"-"`
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

	// Type: see the `Secret.Type` enum for a description of values.
	// Default value: unknown_type
	Type SecretType `json:"type"`
}

// AddSecretOwnerRequest: add secret owner request.
type AddSecretOwnerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`

	// Product: see `Product` enum for description of values.
	// Default value: unknown_product
	Product Product `json:"product"`
}

// BasicCredentials: basic credentials.
type BasicCredentials struct {
	// Username: the username or identifier associated with the credentials.
	Username string `json:"username"`

	// Password: the password associated with the credentials.
	Password string `json:"password"`
}

// BrowseSecretsRequest: browse secrets request.
type BrowseSecretsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: filter by Project ID (optional).
	ProjectID *string `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy BrowseSecretsRequestOrderBy `json:"-"`

	// Prefix: filter secrets and folders for a given prefix (default /).
	Prefix string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// Tags: filter secrets by tags.
	Tags []string `json:"-"`

	// Type: filter by secret type (optional).
	// Default value: unknown_type
	Type SecretType `json:"-"`
}

// BrowseSecretsResponse: browse secrets response.
type BrowseSecretsResponse struct {
	// Items: repeated item of type secret or folder, sorted by the request parameter.
	Items []*BrowseSecretsResponseItem `json:"items"`

	// CurrentPath: current path for the given prefix.
	CurrentPath string `json:"current_path"`

	// TotalCount: count of all secrets and folders matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *BrowseSecretsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *BrowseSecretsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*BrowseSecretsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Items = append(r.Items, results.Items...)
	r.TotalCount += uint64(len(results.Items))
	return uint64(len(results.Items)), nil
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

	// Type: (Optional.) See the `Secret.Type` enum for a description of values. If not specified, the type is `Opaque`.
	// Default value: unknown_type
	Type SecretType `json:"type"`

	// Path: (Optional.) Location of the secret in the directory structure. If not specified, the path is `/`.
	Path *string `json:"path,omitempty"`

	// EphemeralPolicy: (Optional.) Policy that defines whether/when a secret's versions expire. By default, the policy is applied to all the secret's versions.
	EphemeralPolicy *EphemeralPolicy `json:"ephemeral_policy,omitempty"`

	// Protected: a protected secret cannot be deleted.
	Protected bool `json:"protected"`

	// KeyID: (Optional.) The Scaleway Key Manager key ID will be used to encrypt and decrypt secret versions. If not specified, Secret Manager will use a Key Manager internal key.
	KeyID *string `json:"key_id,omitempty"`
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

	// DataCrc32: if specified, Secret Manager will verify the integrity of the data received against the given CRC32 checksum. An error is returned if the CRC32 does not match. If, however, the CRC32 matches, it will be stored and returned along with the SecretVersion on future access requests.
	DataCrc32 *uint32 `json:"data_crc32,omitempty"`
}

// DatabaseCredentials: database credentials.
type DatabaseCredentials struct {
	// Engine: supported database engines are: 'postgres', 'mysql', 'other'.
	Engine string `json:"engine"`

	// Username: the username used to authenticate to the database server.
	Username string `json:"username"`

	// Password: the password used to authenticate to the database server.
	Password string `json:"password"`

	// Host: the hostname or resolvable DNS name of the database server.
	Host string `json:"host"`

	// Dbname: the name of the database to connect to.
	Dbname string `json:"dbname"`

	// Port: the port must be an integer ranging from 0 to 65535.
	Port string `json:"port"`
}

// DeleteSecretRequest: delete secret request.
type DeleteSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`
}

// DeleteSecretVersionRequest: delete secret version request.
type DeleteSecretVersionRequest struct {
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

// GetSecretRequest: get secret request.
type GetSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret.
	SecretID string `json:"-"`
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

// ListSecretTypesRequest: list secret types request.
type ListSecretTypesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project to target.
	ProjectID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListSecretTypesResponse: list secret types response.
type ListSecretTypesResponse struct {
	// Types: list of secret types.
	Types []SecretType `json:"types"`

	// TotalCount: count of all secret types matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecretTypesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecretTypesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListSecretTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Types = append(r.Types, results.Types...)
	r.TotalCount += uint64(len(results.Types))
	return uint64(len(results.Types)), nil
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
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecretVersionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecretVersionsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListSecretVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Versions = append(r.Versions, results.Versions...)
	r.TotalCount += uint64(len(results.Versions))
	return uint64(len(results.Versions)), nil
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

	// Path: filter by exact path (optional).
	Path *string `json:"-"`

	// Ephemeral: filter by ephemeral / not ephemeral (optional).
	Ephemeral *bool `json:"-"`

	// Type: filter by secret type (optional).
	// Default value: unknown_type
	Type SecretType `json:"-"`

	// ScheduledForDeletion: filter by whether the secret was scheduled for deletion / not scheduled for deletion. By default, it will display only not scheduled for deletion secrets.
	ScheduledForDeletion bool `json:"-"`
}

// ListSecretsResponse: list secrets response.
type ListSecretsResponse struct {
	// Secrets: single page of secrets matching the requested criteria.
	Secrets []*Secret `json:"secrets"`

	// TotalCount: count of all secrets matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSecretsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSecretsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListSecretsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Secrets = append(r.Secrets, results.Secrets...)
	r.TotalCount += uint64(len(results.Secrets))
	return uint64(len(results.Secrets)), nil
}

// ListTagsRequest: list tags request.
type ListTagsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project to target.
	ProjectID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListTagsResponse: list tags response.
type ListTagsResponse struct {
	// Tags: list of tags.
	Tags []string `json:"tags"`

	// TotalCount: count of all tags matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTagsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTagsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListTagsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tags = append(r.Tags, results.Tags...)
	r.TotalCount += uint64(len(results.Tags))
	return uint64(len(results.Tags)), nil
}

// ProtectSecretRequest: protect secret request.
type ProtectSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret to enable secret protection for.
	SecretID string `json:"-"`
}

// RestoreSecretRequest: restore secret request.
type RestoreSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	SecretID string `json:"-"`
}

// RestoreSecretVersionRequest: restore secret version request.
type RestoreSecretVersionRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	SecretID string `json:"-"`

	Revision string `json:"-"`
}

// SSHKey: ssh key.
type SSHKey struct {
	// SSHPrivateKey: the private SSH key.
	SSHPrivateKey string `json:"ssh_private_key"`
}

// UnprotectSecretRequest: unprotect secret request.
type UnprotectSecretRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SecretID: ID of the secret to disable secret protection for.
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

// This API allows you to manage your Secret Manager services, for storing, accessing and sharing sensitive data such as passwords, API keys and certificates.
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

// CreateSecret: Create a secret in a given region specified by the `region` parameter.
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets",
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
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
	parameter.AddToQuery(query, "path", req.Path)
	parameter.AddToQuery(query, "ephemeral", req.Ephemeral)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "scheduled_for_deletion", req.ScheduledForDeletion)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets",
		Query:  query,
	}

	var resp ListSecretsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// BrowseSecrets: Retrieve the list of secrets and folders for the given prefix. You must specify either the `organization_id` or the `project_id` and the `region`.
func (s *API) BrowseSecrets(req *BrowseSecretsRequest, opts ...scw.RequestOption) (*BrowseSecretsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "prefix", req.Prefix)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "tags", req.Tags)
	parameter.AddToQuery(query, "type", req.Type)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/browse",
		Query:  query,
	}

	var resp BrowseSecretsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ProtectSecret: Enable secret protection for a given secret specified by the `secret_id` parameter. Enabling secret protection means that your secret can be read and modified, but it cannot be deleted.
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/protect",
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

// UnprotectSecret: Disable secret protection for a given secret specified by the `secret_id` parameter. Disabling secret protection means that your secret can be read, modified and deleted.
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/unprotect",
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/add-owner",
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions",
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "",
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "",
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

// DeleteSecretVersion: Delete a secret's version and the sensitive data contained in it. Deleting a version is permanent and cannot be undone.
func (s *API) DeleteSecretVersion(req *DeleteSecretVersionRequest, opts ...scw.RequestOption) error {
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

	if fmt.Sprint(req.Revision) == "" {
		return errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions",
		Query:  query,
	}

	var resp ListSecretVersionsResponse

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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/access",
	}

	var resp AccessSecretVersionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AccessSecretVersionByPath: Access sensitive data in a secret's version specified by the `region`, `secret_name`, `secret_path` and `revision` parameters.
func (s *API) AccessSecretVersionByPath(req *AccessSecretVersionByPathRequest, opts ...scw.RequestOption) (*AccessSecretVersionResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "secret_path", req.SecretPath)
	parameter.AddToQuery(query, "secret_name", req.SecretName)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.Revision) == "" {
		return nil, errors.New("field Revision cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets-by-path/versions/" + fmt.Sprint(req.Revision) + "/access",
		Query:  query,
	}

	var resp AccessSecretVersionResponse

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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/enable",
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/disable",
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

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/tags",
		Query:  query,
	}

	var resp ListTagsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSecretTypes: List all secret types created within a given Project.
func (s *API) ListSecretTypes(req *ListSecretTypesRequest, opts ...scw.RequestOption) (*ListSecretTypesResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secret-types",
		Query:  query,
	}

	var resp ListSecretTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestoreSecretVersion: Restore a secret's version specified by the `region`, `secret_id` and `revision` parameters.
func (s *API) RestoreSecretVersion(req *RestoreSecretVersionRequest, opts ...scw.RequestOption) (*SecretVersion, error) {
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/versions/" + fmt.Sprint(req.Revision) + "/restore",
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

// RestoreSecret: Restore a secret and all its versions scheduled for deletion specified by the `region` and `secret_id` parameters.
func (s *API) RestoreSecret(req *RestoreSecretRequest, opts ...scw.RequestOption) (*Secret, error) {
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
		Path:   "/secret-manager/v1beta1/regions/" + fmt.Sprint(req.Region) + "/secrets/" + fmt.Sprint(req.SecretID) + "/restore",
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

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package key_manager provides methods and message types of the key_manager v1alpha1 API.
package key_manager

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

type DataKeyAlgorithmSymmetricEncryption string

const (
	DataKeyAlgorithmSymmetricEncryptionUnknownSymmetricEncryption = DataKeyAlgorithmSymmetricEncryption("unknown_symmetric_encryption")
	// Key Manager currently only supports the `AES-GCM` (256-bits) data key algorithm.
	DataKeyAlgorithmSymmetricEncryptionAes256Gcm = DataKeyAlgorithmSymmetricEncryption("aes_256_gcm")
)

func (enum DataKeyAlgorithmSymmetricEncryption) String() string {
	if enum == "" {
		// return default value if empty
		return string(DataKeyAlgorithmSymmetricEncryptionUnknownSymmetricEncryption)
	}
	return string(enum)
}

func (enum DataKeyAlgorithmSymmetricEncryption) Values() []DataKeyAlgorithmSymmetricEncryption {
	return []DataKeyAlgorithmSymmetricEncryption{
		"unknown_symmetric_encryption",
		"aes_256_gcm",
	}
}

func (enum DataKeyAlgorithmSymmetricEncryption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DataKeyAlgorithmSymmetricEncryption) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DataKeyAlgorithmSymmetricEncryption(DataKeyAlgorithmSymmetricEncryption(tmp).String())
	return nil
}

type KeyAlgorithmSymmetricEncryption string

const (
	KeyAlgorithmSymmetricEncryptionUnknownSymmetricEncryption = KeyAlgorithmSymmetricEncryption("unknown_symmetric_encryption")
	// Key Manager currently only supports the `AES-GCM` (256-bits) key algorithm.
	KeyAlgorithmSymmetricEncryptionAes256Gcm = KeyAlgorithmSymmetricEncryption("aes_256_gcm")
)

func (enum KeyAlgorithmSymmetricEncryption) String() string {
	if enum == "" {
		// return default value if empty
		return string(KeyAlgorithmSymmetricEncryptionUnknownSymmetricEncryption)
	}
	return string(enum)
}

func (enum KeyAlgorithmSymmetricEncryption) Values() []KeyAlgorithmSymmetricEncryption {
	return []KeyAlgorithmSymmetricEncryption{
		"unknown_symmetric_encryption",
		"aes_256_gcm",
	}
}

func (enum KeyAlgorithmSymmetricEncryption) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *KeyAlgorithmSymmetricEncryption) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = KeyAlgorithmSymmetricEncryption(KeyAlgorithmSymmetricEncryption(tmp).String())
	return nil
}

type KeyOrigin string

const (
	KeyOriginUnknownOrigin = KeyOrigin("unknown_origin")
	// Key Manager generates the key material upon key creation.
	KeyOriginScalewayKms = KeyOrigin("scaleway_kms")
	// Key Manager creates a key with key material coming from an external source.
	KeyOriginExternal = KeyOrigin("external")
)

func (enum KeyOrigin) String() string {
	if enum == "" {
		// return default value if empty
		return string(KeyOriginUnknownOrigin)
	}
	return string(enum)
}

func (enum KeyOrigin) Values() []KeyOrigin {
	return []KeyOrigin{
		"unknown_origin",
		"scaleway_kms",
		"external",
	}
}

func (enum KeyOrigin) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *KeyOrigin) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = KeyOrigin(KeyOrigin(tmp).String())
	return nil
}

type KeyState string

const (
	KeyStateUnknownState = KeyState("unknown_state")
	// The key can be used for cryptographic operations.
	KeyStateEnabled = KeyState("enabled")
	// The key cannot be used for cryptographic operations.
	KeyStateDisabled = KeyState("disabled")
	// Key material must be imported before the key can be used for cryptographic operations.
	KeyStatePendingKeyMaterial = KeyState("pending_key_material")
)

func (enum KeyState) String() string {
	if enum == "" {
		// return default value if empty
		return string(KeyStateUnknownState)
	}
	return string(enum)
}

func (enum KeyState) Values() []KeyState {
	return []KeyState{
		"unknown_state",
		"enabled",
		"disabled",
		"pending_key_material",
	}
}

func (enum KeyState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *KeyState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = KeyState(KeyState(tmp).String())
	return nil
}

type ListKeysRequestOrderBy string

const (
	ListKeysRequestOrderByNameAsc       = ListKeysRequestOrderBy("name_asc")
	ListKeysRequestOrderByNameDesc      = ListKeysRequestOrderBy("name_desc")
	ListKeysRequestOrderByCreatedAtAsc  = ListKeysRequestOrderBy("created_at_asc")
	ListKeysRequestOrderByCreatedAtDesc = ListKeysRequestOrderBy("created_at_desc")
	ListKeysRequestOrderByUpdatedAtAsc  = ListKeysRequestOrderBy("updated_at_asc")
	ListKeysRequestOrderByUpdatedAtDesc = ListKeysRequestOrderBy("updated_at_desc")
)

func (enum ListKeysRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListKeysRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListKeysRequestOrderBy) Values() []ListKeysRequestOrderBy {
	return []ListKeysRequestOrderBy{
		"name_asc",
		"name_desc",
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
	}
}

func (enum ListKeysRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListKeysRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListKeysRequestOrderBy(ListKeysRequestOrderBy(tmp).String())
	return nil
}

// KeyRotationPolicy: key rotation policy.
type KeyRotationPolicy struct {
	// RotationPeriod: time interval between two key rotations. The minimum duration is 24 hours and the maximum duration is 1 year (876000 hours).
	RotationPeriod *scw.Duration `json:"rotation_period"`

	// NextRotationAt: timestamp indicating the next scheduled rotation.
	NextRotationAt *time.Time `json:"next_rotation_at"`
}

// KeyUsage: key usage.
type KeyUsage struct {
	// SymmetricEncryption: see the `Key.Algorithm.SymmetricEncryption` enum for a description of values.
	// Default value: unknown_symmetric_encryption
	// Precisely one of SymmetricEncryption must be set.
	SymmetricEncryption *KeyAlgorithmSymmetricEncryption `json:"symmetric_encryption,omitempty"`
}

// Key: key.
type Key struct {
	// ID: ID of the key.
	ID string `json:"id"`

	// ProjectID: ID of the Project containing the key.
	ProjectID string `json:"project_id"`

	// Name: name of the key.
	Name string `json:"name"`

	// Usage: keys with a usage set to `symmetric_encryption` can encrypt and decrypt data using the `AES-256-GCM` key algorithm. Key Manager currently only supports `AES-256-GCM`.
	Usage *KeyUsage `json:"usage"`

	// State: see the `Key.State` enum for a description of possible values.
	// Default value: unknown_state
	State KeyState `json:"state"`

	// RotationCount: the rotation count tracks the number of times the key has been rotated.
	RotationCount uint32 `json:"rotation_count"`

	// CreatedAt: key creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: key last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Protected: returns `true` if key protection is applied to the key.
	Protected bool `json:"protected"`

	// Locked: returns `true` if the key is locked.
	Locked bool `json:"locked"`

	// Description: description of the key.
	Description *string `json:"description"`

	// Tags: list of the key's tags.
	Tags []string `json:"tags"`

	// RotatedAt: key last rotation date.
	RotatedAt *time.Time `json:"rotated_at"`

	// RotationPolicy: key rotation policy.
	RotationPolicy *KeyRotationPolicy `json:"rotation_policy"`

	// Origin: refer to the `Key.Origin` enum for a description of values.
	// Default value: unknown_origin
	Origin KeyOrigin `json:"origin"`

	// Region: region where the key is stored.
	Region scw.Region `json:"region"`
}

// CreateKeyRequest: create key request.
type CreateKeyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project containing the key.
	ProjectID string `json:"project_id"`

	// Name: (Optional) Name of the key.
	Name *string `json:"name,omitempty"`

	// Usage: see the `Key.Algorithm.SymmetricEncryption` enum for a description of values.
	Usage *KeyUsage `json:"usage,omitempty"`

	// Description: (Optional) Description of the key.
	Description *string `json:"description,omitempty"`

	// Tags: (Optional) List of the key's tags.
	Tags []string `json:"tags"`

	// RotationPolicy: if not specified, no rotation policy will be applied to the key.
	RotationPolicy *KeyRotationPolicy `json:"rotation_policy,omitempty"`

	// Unprotected: default value is `false`.
	Unprotected bool `json:"unprotected"`

	// Origin: refer to the `Key.Origin` enum for a description of values.
	// Default value: unknown_origin
	Origin KeyOrigin `json:"origin"`
}

// DataKey: data key.
type DataKey struct {
	// KeyID: ID of the data encryption key.
	KeyID string `json:"key_id"`

	// Algorithm: symmetric encryption algorithm of the data encryption key (`AES-256-GCM`).
	// Default value: unknown_symmetric_encryption
	Algorithm DataKeyAlgorithmSymmetricEncryption `json:"algorithm"`

	// Ciphertext: your data encryption key's ciphertext can be stored safely. It can only be decrypted through the keys you create in Key Manager, using the relevant key ID.
	Ciphertext []byte `json:"ciphertext"`

	// Plaintext: (Optional) Your data encryption key's plaintext allows you to use the key immediately upon creation. It must neither be stored or shared.
	Plaintext *[]byte `json:"plaintext"`

	// CreatedAt: data encryption key creation date.
	CreatedAt *time.Time `json:"created_at"`
}

// DecryptRequest: decrypt request.
type DecryptRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key to decrypt.
	KeyID string `json:"-"`

	// Ciphertext: data size must be between 1 and 131071 bytes.
	Ciphertext []byte `json:"ciphertext"`

	// AssociatedData: the additional data must match the value passed in the encryption request.
	AssociatedData *[]byte `json:"associated_data,omitempty"`
}

// DecryptResponse: decrypt response.
type DecryptResponse struct {
	// KeyID: ID of the key used for decryption.
	KeyID string `json:"key_id"`

	// Plaintext: key's decrypted data.
	Plaintext []byte `json:"plaintext"`

	// Ciphertext: if the data was already encrypted with the latest key rotation, no output will be returned in the response object.
	Ciphertext *[]byte `json:"ciphertext"`
}

// DeleteKeyMaterialRequest: delete key material request.
type DeleteKeyMaterialRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key of which to delete the key material.
	KeyID string `json:"-"`
}

// DeleteKeyRequest: delete key request.
type DeleteKeyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key to delete.
	KeyID string `json:"-"`
}

// DisableKeyRequest: disable key request.
type DisableKeyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key to disable.
	KeyID string `json:"-"`
}

// EnableKeyRequest: enable key request.
type EnableKeyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key to enable.
	KeyID string `json:"-"`
}

// EncryptRequest: encrypt request.
type EncryptRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key to encrypt.
	KeyID string `json:"-"`

	// Plaintext: data size must be between 1 and 65535 bytes.
	Plaintext []byte `json:"plaintext"`

	// AssociatedData: additional data which will not be encrypted, but authenticated and appended to the encrypted payload.
	AssociatedData *[]byte `json:"associated_data,omitempty"`
}

// EncryptResponse: encrypt response.
type EncryptResponse struct {
	// KeyID: ID of the key used for encryption.
	KeyID string `json:"key_id"`

	// Ciphertext: key's encrypted data.
	Ciphertext []byte `json:"ciphertext"`
}

// GenerateDataKeyRequest: generate data key request.
type GenerateDataKeyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key.
	KeyID string `json:"-"`

	// Algorithm: see the `DataKey.Algorithm.SymmetricEncryption` enum for a description of values.
	// Default value: unknown_symmetric_encryption
	Algorithm DataKeyAlgorithmSymmetricEncryption `json:"algorithm"`

	// WithoutPlaintext: default value is `false`, meaning that the plaintext is returned.
	// Set it to `true` if you do not wish the plaintext to be returned in the response object.
	WithoutPlaintext bool `json:"without_plaintext"`
}

// GetKeyRequest: get key request.
type GetKeyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key to target.
	KeyID string `json:"-"`
}

// GetPublicKeyRequest: get public key request.
type GetPublicKeyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key.
	KeyID string `json:"-"`
}

// ImportKeyMaterialRequest: import key material request.
type ImportKeyMaterialRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: the key's origin must be `external`.
	KeyID string `json:"-"`

	// KeyMaterial: the key material The key material is a random sequence of bytes used to derive a cryptographic key.
	KeyMaterial []byte `json:"key_material"`

	// Salt: a salt is random data added to key material to ensure unique derived keys, even if the input is similar. It helps strengthen security when the key material has low randomness (low entropy).
	Salt *[]byte `json:"salt,omitempty"`
}

// ListKeysRequest: list keys request.
type ListKeysRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrganizationID: (Optional) Filter by Organization ID.
	OrganizationID *string `json:"-"`

	// ProjectID: (Optional) Filter by Project ID.
	ProjectID *string `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListKeysRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// Tags: (Optional) List of tags to filter on.
	Tags []string `json:"-"`

	// Name: (Optional) Filter by key name.
	Name *string `json:"-"`
}

// ListKeysResponse: list keys response.
type ListKeysResponse struct {
	// Keys: single page of keys matching the requested criteria.
	Keys []*Key `json:"keys"`

	// TotalCount: total count of keys matching the requested criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListKeysResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListKeysResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListKeysResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Keys = append(r.Keys, results.Keys...)
	r.TotalCount += uint64(len(results.Keys))
	return uint64(len(results.Keys)), nil
}

// ProtectKeyRequest: protect key request.
type ProtectKeyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key to apply key protection to.
	KeyID string `json:"-"`
}

// PublicKey: public key.
type PublicKey struct {
	Pem string `json:"pem"`
}

// RotateKeyRequest: rotate key request.
type RotateKeyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key to rotate.
	KeyID string `json:"-"`
}

// UnprotectKeyRequest: unprotect key request.
type UnprotectKeyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key to remove key protection from.
	KeyID string `json:"-"`
}

// UpdateKeyRequest: update key request.
type UpdateKeyRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// KeyID: ID of the key to update.
	KeyID string `json:"-"`

	// Name: (Optional) Updated name of the key.
	Name *string `json:"name,omitempty"`

	// Description: (Optional) Updated description of the key.
	Description *string `json:"description,omitempty"`

	// Tags: (Optional) Updated list of the key's tags.
	Tags *[]string `json:"tags,omitempty"`

	// RotationPolicy: if not specified, the key's existing rotation policy applies.
	RotationPolicy *KeyRotationPolicy `json:"rotation_policy,omitempty"`
}

// This API allows you to create, manage and use cryptographic keys in a centralized and secure service.
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

// CreateKey: Create a key in a given region specified by the `region` parameter. Keys only support symmetric encryption. You can use keys to encrypt or decrypt arbitrary payloads, or to generate data encryption keys. **Data encryption keys are not stored in Key Manager**.
func (s *API) CreateKey(req *CreateKeyRequest, opts ...scw.RequestOption) (*Key, error) {
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
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Key

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetKey: Retrieve metadata for a specified key using the `region` and `key_id` parameters.
func (s *API) GetKey(req *GetKeyRequest, opts ...scw.RequestOption) (*Key, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "",
	}

	var resp Key

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPublicKey: Retrieves the public portion of an asymmetric cryptographic key in PEM format.
func (s *API) GetPublicKey(req *GetPublicKeyRequest, opts ...scw.RequestOption) (*PublicKey, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "/public-key",
	}

	var resp PublicKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateKey: Modify a key's metadata including name, description and tags, specified by the `key_id` and `region` parameters.
func (s *API) UpdateKey(req *UpdateKeyRequest, opts ...scw.RequestOption) (*Key, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Key

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteKey: Permanently delete a key specified by the `region` and `key_id` parameters. This action is irreversible. Any data encrypted with this key, including data encryption keys, will no longer be decipherable.
func (s *API) DeleteKey(req *DeleteKeyRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// RotateKey: Generate a new version of an existing key with new key material. Previous key versions remain usable to decrypt previously encrypted data, but the key's new version will be used for subsequent encryption operations and data key generation.
func (s *API) RotateKey(req *RotateKeyRequest, opts ...scw.RequestOption) (*Key, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "/rotate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Key

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ProtectKey: Apply protection to a given key specified by the `key_id` parameter. Applying key protection means that your key can be used and modified, but it cannot be deleted.
func (s *API) ProtectKey(req *ProtectKeyRequest, opts ...scw.RequestOption) (*Key, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "/protect",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Key

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnprotectKey: Remove key protection from a given key specified by the `key_id` parameter. Removing key protection means that your key can be deleted anytime.
func (s *API) UnprotectKey(req *UnprotectKeyRequest, opts ...scw.RequestOption) (*Key, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "/unprotect",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Key

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableKey: Enable a given key to be used for cryptographic operations. Enabling a key allows you to make a disabled key usable again. You must specify the `region` and `key_id` parameters.
func (s *API) EnableKey(req *EnableKeyRequest, opts ...scw.RequestOption) (*Key, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "/enable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Key

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableKey: Disable a given key, preventing it to be used for cryptographic operations. Disabling a key renders it unusable. You must specify the `region` and `key_id` parameters.
func (s *API) DisableKey(req *DisableKeyRequest, opts ...scw.RequestOption) (*Key, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "/disable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Key

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListKeys: Retrieve a list of keys across all Projects in an Organization or within a specific Project. You must specify the `region`, and either the `organization_id` or the `project_id`.
func (s *API) ListKeys(req *ListKeysRequest, opts ...scw.RequestOption) (*ListKeysResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys",
		Query:  query,
	}

	var resp ListKeysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GenerateDataKey: Create a new data encryption key for cryptographic operations outside of Key Manager. The data encryption key is encrypted and must be decrypted using the key you have created in Key Manager.
//
// The data encryption key is returned in plaintext and ciphertext but it should only be stored in its encrypted form (ciphertext). Key Manager does not store your data encryption key. To retrieve your key's plaintext, use the `Decrypt` method with your key's ID and ciphertext.
func (s *API) GenerateDataKey(req *GenerateDataKeyRequest, opts ...scw.RequestOption) (*DataKey, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "/generate-data-key",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DataKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Encrypt: Encrypt a payload using an existing key, specified by the `key_id` parameter. Only keys with a usage set to `symmetric_encryption` are supported by this method. The maximum payload size that can be encrypted is 64 KB of plaintext.
func (s *API) Encrypt(req *EncryptRequest, opts ...scw.RequestOption) (*EncryptResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "/encrypt",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EncryptResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Decrypt: Decrypt an encrypted payload using an existing key, specified by the `key_id` parameter. The maximum payload size that can be decrypted is equivalent to the encrypted output of 64 KB of data (around 131 KB).
func (s *API) Decrypt(req *DecryptRequest, opts ...scw.RequestOption) (*DecryptResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "/decrypt",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DecryptResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ImportKeyMaterial: Import externally generated key material into Key Manager to derive a new cryptographic key. The key's origin must be `external`.
func (s *API) ImportKeyMaterial(req *ImportKeyMaterialRequest, opts ...scw.RequestOption) (*Key, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "/import-key-material",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Key

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteKeyMaterial: Delete previously imported key material. This renders the associated cryptographic key unusable for any operation. The key's origin must be `external`.
func (s *API) DeleteKeyMaterial(req *DeleteKeyMaterialRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.KeyID) == "" {
		return errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/key-manager/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/keys/" + fmt.Sprint(req.KeyID) + "/delete-key-material",
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

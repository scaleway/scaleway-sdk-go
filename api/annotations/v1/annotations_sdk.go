// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package annotations provides methods and message types of the annotations v1 API.
package annotations

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

// ListAllKeysAndValuesResponseValue: list all keys and values response value.
type ListAllKeysAndValuesResponseValue struct {
	// ID: ID of the value.
	ID string `json:"id"`

	// Name: name of the value.
	Name string `json:"name"`

	// Description: description of the value.
	Description string `json:"description"`
}

// BindingKey: binding key.
type BindingKey struct {
	// ID: ID of the key.
	ID string `json:"id"`

	// Name: name of the key.
	Name string `json:"name"`
}

// BindingValue: binding value.
type BindingValue struct {
	// ID: ID of the value.
	ID string `json:"id"`

	// Name: name of the value.
	Name string `json:"name"`
}

// ListAllKeysAndValuesResponseKey: list all keys and values response key.
type ListAllKeysAndValuesResponseKey struct {
	// ID: ID of the key.
	ID string `json:"id"`

	// Name: name of the key.
	Name string `json:"name"`

	// Description: description of the key.
	Description string `json:"description"`

	// Values: list of values associated with the key, sorted alphabetically by name.
	Values []*ListAllKeysAndValuesResponseValue `json:"values"`
}

// Binding: binding.
type Binding struct {
	// ID: ID of the binding.
	ID string `json:"id"`

	// Srn: scaleway Resource Number associated to the binding.
	Srn string `json:"srn"`

	// Key: key associated to the binding.
	Key *BindingKey `json:"key"`

	// Value: value associated to the binding.
	Value *BindingValue `json:"value"`
}

// Key: key.
type Key struct {
	// ID: ID of the annotation key.
	ID string `json:"id"`

	// Name: name of the annotation key.
	Name string `json:"name"`

	// Description: description of the annotation key.
	Description string `json:"description"`
}

// Value: value.
type Value struct {
	// ID: ID of the value.
	ID string `json:"id"`

	// KeyID: ID of the key the value is associated to.
	KeyID string `json:"key_id"`

	// Name: name of the value (e.g. "production" for a key "environment").
	Name string `json:"name"`

	// Description: description of the value.
	Description string `json:"description"`
}

// CreateBindingRequest: create binding request.
type CreateBindingRequest struct {
	// Srn: scaleway Resource Number to associate.
	Srn string `json:"srn"`

	// ValueID: ID of the value to associate.
	ValueID string `json:"value_id"`
}

// CreateKeyRequest: create key request.
type CreateKeyRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"organization_id"`

	// Name: name of the annotation key.
	Name string `json:"name"`

	// Description: description of the annotation key.
	Description string `json:"description"`
}

// CreateValueRequest: create value request.
type CreateValueRequest struct {
	// KeyID: ID of the key the value will be bound to.
	KeyID string `json:"key_id"`

	// Name: name of the value.
	Name string `json:"name"`

	// Description: description of the value.
	Description string `json:"description"`
}

// DeleteAllBindingsMatchingSRNRequest: delete all bindings matching srn request.
type DeleteAllBindingsMatchingSRNRequest struct {
	// Srn: scaleway Resource Number for which all bindings should be deleted.
	Srn string `json:"-"`

	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`
}

// DeleteAllBindingsMatchingSRNResponse: delete all bindings matching srn response.
type DeleteAllBindingsMatchingSRNResponse struct {
	// TotalDeleted: total number of bindings deleted.
	TotalDeleted uint64 `json:"total_deleted"`
}

// DeleteAllBindingsMatchingValueRequest: delete all bindings matching value request.
type DeleteAllBindingsMatchingValueRequest struct {
	// ValueID: ID of the value for which all bindings should be deleted.
	ValueID string `json:"-"`
}

// DeleteAllBindingsMatchingValueResponse: delete all bindings matching value response.
type DeleteAllBindingsMatchingValueResponse struct {
	// TotalDeleted: total number of bindings deleted.
	TotalDeleted uint64 `json:"total_deleted"`
}

// DeleteAllValuesMatchingKeyRequest: delete all values matching key request.
type DeleteAllValuesMatchingKeyRequest struct {
	// KeyID: ID of the key for which to delete all values.
	KeyID string `json:"-"`
}

// DeleteAllValuesMatchingKeyResponse: delete all values matching key response.
type DeleteAllValuesMatchingKeyResponse struct {
	// TotalDeleted: total number of bindings deleted.
	TotalDeleted uint64 `json:"total_deleted"`
}

// DeleteBindingRequest: delete binding request.
type DeleteBindingRequest struct {
	// BindingID: ID of the binding to delete.
	BindingID string `json:"-"`
}

// DeleteKeyRequest: delete key request.
type DeleteKeyRequest struct {
	// KeyID: ID of the key to delete.
	KeyID string `json:"-"`
}

// DeleteValueRequest: delete value request.
type DeleteValueRequest struct {
	// ValueID: ID of the value to delete.
	ValueID string `json:"-"`
}

// GetKeyRequest: get key request.
type GetKeyRequest struct {
	// KeyID: ID of the key to retrieve.
	KeyID string `json:"-"`
}

// GetValueRequest: get value request.
type GetValueRequest struct {
	// ValueID: ID of the value to retrieve.
	ValueID string `json:"-"`
}

// ListAllKeysAndValuesRequest: list all keys and values request.
type ListAllKeysAndValuesRequest struct {
	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`
}

// ListAllKeysAndValuesResponse: list all keys and values response.
type ListAllKeysAndValuesResponse struct {
	// Keys: list of keys with values for an organization, sorted alphabetically by name.
	Keys []*ListAllKeysAndValuesResponseKey `json:"keys"`
}

// ListBindingsRequest: list bindings request.
type ListBindingsRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of bindings on the page.
	PageSize *uint32 `json:"-"`

	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`

	// Srn: scaleway Resource Number for which to list all bindings.
	Srn *string `json:"-"`

	// ValueID: value ID for which to list all bindings.
	ValueID *string `json:"-"`
}

// ListBindingsResponse: list bindings response.
type ListBindingsResponse struct {
	// Bindings: list of bindings for the organization. Response order by ID.
	Bindings []*Binding `json:"bindings"`

	// TotalCount: total number of bindings returned.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBindingsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBindingsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListBindingsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Bindings = append(r.Bindings, results.Bindings...)
	r.TotalCount += uint64(len(results.Bindings))
	return uint64(len(results.Bindings)), nil
}

// ListKeysRequest: list keys request.
type ListKeysRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of keys on the page.
	PageSize *uint32 `json:"-"`

	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`
}

// ListKeysResponse: list keys response.
type ListKeysResponse struct {
	// Keys: list of keys for an organization, sorted alphabetically by name.
	Keys []*Key `json:"keys"`

	// TotalCount: total number of keys returned.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListKeysResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListKeysResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListKeysResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Keys = append(r.Keys, results.Keys...)
	r.TotalCount += uint64(len(results.Keys))
	return uint64(len(results.Keys)), nil
}

// ListValuesRequest: list values request.
type ListValuesRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: maximum number of values on the page.
	PageSize *uint32 `json:"-"`

	// OrganizationID: ID of the organization.
	OrganizationID string `json:"-"`

	// KeyID: ID of the key to list the values for.
	KeyID *string `json:"-"`
}

// ListValuesResponse: list values response.
type ListValuesResponse struct {
	// Values: list of values for a key, sorted alphabetically by name.
	Values []*Value `json:"values"`

	// TotalCount: total number of values returned.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListValuesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListValuesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListValuesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Values = append(r.Values, results.Values...)
	r.TotalCount += uint64(len(results.Values))
	return uint64(len(results.Values)), nil
}

// UpdateKeyRequest: update key request.
type UpdateKeyRequest struct {
	// KeyID: ID of the key to update.
	KeyID string `json:"-"`

	// Name: new name of the key.
	Name *string `json:"name,omitempty"`

	// Description: new description of the key.
	Description *string `json:"description,omitempty"`
}

// UpdateValueRequest: update value request.
type UpdateValueRequest struct {
	// ValueID: ID of the value to update.
	ValueID string `json:"-"`

	// Name: new name of the value.
	Name *string `json:"name,omitempty"`

	// Description: new description of the value.
	Description *string `json:"description,omitempty"`
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

// CreateKey: Create an annotation key.
func (s *API) CreateKey(req *CreateKeyRequest, opts ...scw.RequestOption) (*Key, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/annotations/v1/keys",
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

// ListKeys: List of keys.
func (s *API) ListKeys(req *ListKeysRequest, opts ...scw.RequestOption) (*ListKeysResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/annotations/v1/keys",
		Query:  query,
	}

	var resp ListKeysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetKey: Retrieve a specific key.
func (s *API) GetKey(req *GetKeyRequest, opts ...scw.RequestOption) (*Key, error) {
	var err error

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/annotations/v1/keys/" + fmt.Sprint(req.KeyID) + "",
	}

	var resp Key

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateKey: Update name or description. All associated resources will immediately display the new name.
func (s *API) UpdateKey(req *UpdateKeyRequest, opts ...scw.RequestOption) (*Key, error) {
	var err error

	if fmt.Sprint(req.KeyID) == "" {
		return nil, errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/annotations/v1/keys/" + fmt.Sprint(req.KeyID) + "",
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

// DeleteKey: Delete a key definition. Fails if the key has any associated values.
func (s *API) DeleteKey(req *DeleteKeyRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.KeyID) == "" {
		return errors.New("field KeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/annotations/v1/keys/" + fmt.Sprint(req.KeyID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateValue: Add a value definition to a key.
func (s *API) CreateValue(req *CreateValueRequest, opts ...scw.RequestOption) (*Value, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/annotations/v1/values",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Value

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListValues: List all values for a key, sorted alphabetically by name.
func (s *API) ListValues(req *ListValuesRequest, opts ...scw.RequestOption) (*ListValuesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "key_id", req.KeyID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/annotations/v1/values",
		Query:  query,
	}

	var resp ListValuesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetValue: Retrieve a specific value.
func (s *API) GetValue(req *GetValueRequest, opts ...scw.RequestOption) (*Value, error) {
	var err error

	if fmt.Sprint(req.ValueID) == "" {
		return nil, errors.New("field ValueID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/annotations/v1/values/" + fmt.Sprint(req.ValueID) + "",
	}

	var resp Value

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateValue: Update name or description. Global update.
func (s *API) UpdateValue(req *UpdateValueRequest, opts ...scw.RequestOption) (*Value, error) {
	var err error

	if fmt.Sprint(req.ValueID) == "" {
		return nil, errors.New("field ValueID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/annotations/v1/values/" + fmt.Sprint(req.ValueID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Value

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteValue: Delete a value definition. Fails if the value is currently bound to any resource.
func (s *API) DeleteValue(req *DeleteValueRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.ValueID) == "" {
		return errors.New("field ValueID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/annotations/v1/values/" + fmt.Sprint(req.ValueID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAllValuesMatchingKey: Delete ALL values associated with a key. Fails if any of these values are currently bound to any resource.
func (s *API) DeleteAllValuesMatchingKey(req *DeleteAllValuesMatchingKeyRequest, opts ...scw.RequestOption) (*DeleteAllValuesMatchingKeyResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "key_id", req.KeyID)

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/annotations/v1/values/delete-all-matching-key",
		Query:  query,
	}

	var resp DeleteAllValuesMatchingKeyResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListAllKeysAndValues: List all keys and values for an organization, sorted alphabetically by key name and value name.
func (s *API) ListAllKeysAndValues(req *ListAllKeysAndValuesRequest, opts ...scw.RequestOption) (*ListAllKeysAndValuesResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/annotations/v1/all-keys-and-values",
		Query:  query,
	}

	var resp ListAllKeysAndValuesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBinding: Attach a value to a resource. Fails if the resource already has a value for this key.
func (s *API) CreateBinding(req *CreateBindingRequest, opts ...scw.RequestOption) (*Binding, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/annotations/v1/bindings",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Binding

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListBindings: List all bindings, or filter by Scaleway Resource Number or value ID. Response order by ID.
func (s *API) ListBindings(req *ListBindingsRequest, opts ...scw.RequestOption) (*ListBindingsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "srn", req.Srn)
	parameter.AddToQuery(query, "value_id", req.ValueID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/annotations/v1/bindings",
		Query:  query,
	}

	var resp ListBindingsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBinding: Detach an annotation from a resource.
func (s *API) DeleteBinding(req *DeleteBindingRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.BindingID) == "" {
		return errors.New("field BindingID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/annotations/v1/bindings/" + fmt.Sprint(req.BindingID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteAllBindingsMatchingValue: Delete ALL bindings associated with a value.
func (s *API) DeleteAllBindingsMatchingValue(req *DeleteAllBindingsMatchingValueRequest, opts ...scw.RequestOption) (*DeleteAllBindingsMatchingValueResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "value_id", req.ValueID)

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/annotations/v1/bindings/delete-all-matching-value",
		Query:  query,
	}

	var resp DeleteAllBindingsMatchingValueResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAllBindingsMatchingSRN: Delete ALL bindings associated with a Scaleway Resource Number.
func (s *API) DeleteAllBindingsMatchingSRN(req *DeleteAllBindingsMatchingSRNRequest, opts ...scw.RequestOption) (*DeleteAllBindingsMatchingSRNResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "srn", req.Srn)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/annotations/v1/bindings/delete-all-matching-srn",
		Query:  query,
	}

	var resp DeleteAllBindingsMatchingSRNResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

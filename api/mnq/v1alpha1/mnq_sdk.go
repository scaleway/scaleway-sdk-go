// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package mnq provides methods and message types of the mnq v1alpha1 API.
package mnq

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

// API: this API allows you to manage Messaging or Queueing brokers
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type ListCredentialsRequestOrderBy string

const (
	// ListCredentialsRequestOrderByIDAsc is [insert doc].
	ListCredentialsRequestOrderByIDAsc = ListCredentialsRequestOrderBy("id_asc")
	// ListCredentialsRequestOrderByIDDesc is [insert doc].
	ListCredentialsRequestOrderByIDDesc = ListCredentialsRequestOrderBy("id_desc")
	// ListCredentialsRequestOrderByNameAsc is [insert doc].
	ListCredentialsRequestOrderByNameAsc = ListCredentialsRequestOrderBy("name_asc")
	// ListCredentialsRequestOrderByNameDesc is [insert doc].
	ListCredentialsRequestOrderByNameDesc = ListCredentialsRequestOrderBy("name_desc")
)

func (enum ListCredentialsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "id_asc"
	}
	return string(enum)
}

func (enum ListCredentialsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCredentialsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCredentialsRequestOrderBy(ListCredentialsRequestOrderBy(tmp).String())
	return nil
}

type ListNamespacesRequestOrderBy string

const (
	// ListNamespacesRequestOrderByCreatedAtAsc is [insert doc].
	ListNamespacesRequestOrderByCreatedAtAsc = ListNamespacesRequestOrderBy("created_at_asc")
	// ListNamespacesRequestOrderByCreatedAtDesc is [insert doc].
	ListNamespacesRequestOrderByCreatedAtDesc = ListNamespacesRequestOrderBy("created_at_desc")
	// ListNamespacesRequestOrderByUpdatedAtAsc is [insert doc].
	ListNamespacesRequestOrderByUpdatedAtAsc = ListNamespacesRequestOrderBy("updated_at_asc")
	// ListNamespacesRequestOrderByUpdatedAtDesc is [insert doc].
	ListNamespacesRequestOrderByUpdatedAtDesc = ListNamespacesRequestOrderBy("updated_at_desc")
	// ListNamespacesRequestOrderByIDAsc is [insert doc].
	ListNamespacesRequestOrderByIDAsc = ListNamespacesRequestOrderBy("id_asc")
	// ListNamespacesRequestOrderByIDDesc is [insert doc].
	ListNamespacesRequestOrderByIDDesc = ListNamespacesRequestOrderBy("id_desc")
	// ListNamespacesRequestOrderByNameAsc is [insert doc].
	ListNamespacesRequestOrderByNameAsc = ListNamespacesRequestOrderBy("name_asc")
	// ListNamespacesRequestOrderByNameDesc is [insert doc].
	ListNamespacesRequestOrderByNameDesc = ListNamespacesRequestOrderBy("name_desc")
	// ListNamespacesRequestOrderByProjectIDAsc is [insert doc].
	ListNamespacesRequestOrderByProjectIDAsc = ListNamespacesRequestOrderBy("project_id_asc")
	// ListNamespacesRequestOrderByProjectIDDesc is [insert doc].
	ListNamespacesRequestOrderByProjectIDDesc = ListNamespacesRequestOrderBy("project_id_desc")
)

func (enum ListNamespacesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNamespacesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNamespacesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNamespacesRequestOrderBy(ListNamespacesRequestOrderBy(tmp).String())
	return nil
}

type NamespaceProtocol string

const (
	// NamespaceProtocolUnknown is [insert doc].
	NamespaceProtocolUnknown = NamespaceProtocol("unknown")
	// NamespaceProtocolNats is [insert doc].
	NamespaceProtocolNats = NamespaceProtocol("nats")
	// NamespaceProtocolSqsSns is [insert doc].
	NamespaceProtocolSqsSns = NamespaceProtocol("sqs_sns")
)

func (enum NamespaceProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NamespaceProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NamespaceProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NamespaceProtocol(NamespaceProtocol(tmp).String())
	return nil
}

// Credential: credential
type Credential struct {
	// ID: credential ID
	ID string `json:"id"`
	// Name: credential name
	Name string `json:"name"`
	// NamespaceID: namespace containing the Credential
	NamespaceID string `json:"namespace_id"`
	// Protocol: protocol associated to the Credential
	//
	// Default value: unknown
	Protocol NamespaceProtocol `json:"protocol"`
	// NatsCredentials: credentials file used to connect to the NATS service
	// Precisely one of NatsCredentials, SqsSnsCredentials must be set.
	NatsCredentials *CredentialNATSCredsFile `json:"nats_credentials,omitempty"`
	// SqsSnsCredentials: credential used to connect to the SQS/SNS service
	// Precisely one of NatsCredentials, SqsSnsCredentials must be set.
	SqsSnsCredentials *CredentialSQSSNSCreds `json:"sqs_sns_credentials,omitempty"`
}

// CredentialNATSCredsFile: credential.nats creds file
type CredentialNATSCredsFile struct {
	// Content: raw content of the NATS credentials file
	Content string `json:"content"`
}

// CredentialSQSSNSCreds: credential.sqssns creds
type CredentialSQSSNSCreds struct {
	// AccessKey: ID of the key
	AccessKey string `json:"access_key"`
	// SecretKey: secret value of the key
	SecretKey *string `json:"secret_key"`
	// Permissions: list of permissions associated to this Credential
	Permissions *Permissions `json:"permissions"`
}

// CredentialSummary: credential summary
type CredentialSummary struct {
	// ID: credential ID
	ID string `json:"id"`
	// Name: credential name
	Name string `json:"name"`
	// NamespaceID: namespace containing the Credential
	NamespaceID string `json:"namespace_id"`
	// Protocol: protocol associated to the Credential
	//
	// Default value: unknown
	Protocol NamespaceProtocol `json:"protocol"`
	// SqsSnsCredentials: credential used to connect to the SQS/SNS service
	// Precisely one of SqsSnsCredentials must be set.
	SqsSnsCredentials *CredentialSummarySQSSNSCreds `json:"sqs_sns_credentials,omitempty"`
}

// CredentialSummarySQSSNSCreds: credential summary.sqssns creds
type CredentialSummarySQSSNSCreds struct {
	// AccessKey: ID of the key
	AccessKey string `json:"access_key"`
	// Permissions: list of permissions associated to this Credential
	Permissions *Permissions `json:"permissions"`
}

// ListCredentialsResponse: list credentials response
type ListCredentialsResponse struct {
	// TotalCount: total number of existing Credentials
	TotalCount uint32 `json:"total_count"`
	// Credentials: a page of Credentials
	Credentials []*CredentialSummary `json:"credentials"`
}

// ListNamespacesResponse: list namespaces response
type ListNamespacesResponse struct {
	// TotalCount: total number of existing Namespaces
	TotalCount uint32 `json:"total_count"`
	// Namespaces: a page of Namespaces
	Namespaces []*Namespace `json:"namespaces"`
}

// Namespace: namespace
type Namespace struct {
	// ID: namespace ID
	ID string `json:"id"`
	// Name: namespace name
	Name string `json:"name"`
	// Endpoint: endpoint of the service matching the Namespace protocol
	Endpoint string `json:"endpoint"`
	// Protocol: namespace protocol
	//
	// Default value: unknown
	Protocol NamespaceProtocol `json:"protocol"`
	// ProjectID: project containing the Namespace
	ProjectID string `json:"project_id"`
	// Region: region where the Namespace is deployed
	Region scw.Region `json:"region"`
	// CreatedAt: namespace creation date
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: namespace last modification date
	UpdatedAt *time.Time `json:"updated_at"`
}

// Permissions: permissions
type Permissions struct {
	// CanPublish: defines if user can publish messages to the service
	CanPublish *bool `json:"can_publish"`
	// CanReceive: defines if user can receive messages from the service
	CanReceive *bool `json:"can_receive"`
	// CanManage: defines if user can manage the associated resource(s)
	CanManage *bool `json:"can_manage"`
}

// Service API

// Regions list localities the api is available in
func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar}
}

type ListNamespacesRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// OrganizationID: will list only the Namespaces owned by the specified organization
	OrganizationID *string `json:"-"`
	// ProjectID: will list only the Namespaces contained into the specified project
	ProjectID *string `json:"-"`
	// Page: indicate the page number of results to be returned
	Page *int32 `json:"-"`
	// PageSize: maximum number of results returned by page
	PageSize *uint32 `json:"-"`
	// OrderBy: field used for sorting results
	//
	// Default value: created_at_asc
	OrderBy ListNamespacesRequestOrderBy `json:"-"`
}

// ListNamespaces: list namespaces
func (s *API) ListNamespaces(req *ListNamespacesRequest, opts ...scw.RequestOption) (*ListNamespacesResponse, error) {
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListNamespacesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateNamespaceRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// Name: namespace name
	Name string `json:"name"`
	// Protocol: namespace protocol
	//
	// Default value: unknown
	Protocol NamespaceProtocol `json:"protocol"`
	// ProjectID: project containing the Namespace
	ProjectID string `json:"project_id"`
}

// CreateNamespace: create a namespace
func (s *API) CreateNamespace(req *CreateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("mnq")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateNamespaceRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// NamespaceID: ID of the Namespace to update
	NamespaceID string `json:"namespace_id"`
	// Name: namespace name
	Name *string `json:"name"`
}

// UpdateNamespace: update the name of a namespace
func (s *API) UpdateNamespace(req *UpdateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetNamespaceRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// NamespaceID: ID of the Namespace to get
	NamespaceID string `json:"-"`
}

// GetNamespace: get a namespace
func (s *API) GetNamespace(req *GetNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return nil, errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
		Headers: http.Header{},
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteNamespaceRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// NamespaceID: ID of the Namespace to delete
	NamespaceID string `json:"-"`
}

// DeleteNamespace: delete a namespace
func (s *API) DeleteNamespace(req *DeleteNamespaceRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NamespaceID) == "" {
		return errors.New("field NamespaceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type CreateCredentialRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// NamespaceID: namespace containing the Credential
	NamespaceID string `json:"namespace_id"`
	// Name: credential name
	Name string `json:"name"`
	// Permissions: list of permissions associated to this Credential
	// Precisely one of Permissions must be set.
	Permissions *Permissions `json:"permissions,omitempty"`
}

// CreateCredential: create a set of credentials
//
// Create a set of credentials for a specific namespace.
func (s *API) CreateCredential(req *CreateCredentialRequest, opts ...scw.RequestOption) (*Credential, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("mnq")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/credentials",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Credential

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteCredentialRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// CredentialID: ID of the Credential to delete
	CredentialID string `json:"-"`
}

// DeleteCredential: delete credentials
func (s *API) DeleteCredential(req *DeleteCredentialRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CredentialID) == "" {
		return errors.New("field CredentialID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/credentials/" + fmt.Sprint(req.CredentialID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListCredentialsRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// NamespaceID: namespace containing the Credential
	NamespaceID *string `json:"-"`
	// Page: indicate the page number of results to be returned
	Page *int32 `json:"-"`
	// PageSize: maximum number of results returned by page
	PageSize *uint32 `json:"-"`
	// OrderBy: field used for sorting results
	//
	// Default value: id_asc
	OrderBy ListCredentialsRequestOrderBy `json:"-"`
}

// ListCredentials: list credentials
func (s *API) ListCredentials(req *ListCredentialsRequest, opts ...scw.RequestOption) (*ListCredentialsResponse, error) {
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
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/credentials",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListCredentialsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateCredentialRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// CredentialID: ID of the Credential to update
	CredentialID string `json:"-"`
	// Name: credential name
	Name *string `json:"name"`
	// Permissions: list of permissions associated to this Credential
	// Precisely one of Permissions must be set.
	Permissions *Permissions `json:"permissions,omitempty"`
}

// UpdateCredential: update a set of credentials
//
// Update a set of credentials.
func (s *API) UpdateCredential(req *UpdateCredentialRequest, opts ...scw.RequestOption) (*Credential, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CredentialID) == "" {
		return nil, errors.New("field CredentialID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/credentials/" + fmt.Sprint(req.CredentialID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Credential

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetCredentialRequest struct {
	// Region:
	//
	// Region to target. If none is passed will use default region from the config
	Region scw.Region `json:"-"`
	// CredentialID: ID of the Credential to get
	CredentialID string `json:"-"`
}

// GetCredential: get a set of credentials
func (s *API) GetCredential(req *GetCredentialRequest, opts ...scw.RequestOption) (*Credential, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CredentialID) == "" {
		return nil, errors.New("field CredentialID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/mnq/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/credentials/" + fmt.Sprint(req.CredentialID) + "",
		Headers: http.Header{},
	}

	var resp Credential

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNamespacesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNamespacesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Namespaces = append(r.Namespaces, results.Namespaces...)
	r.TotalCount += uint32(len(results.Namespaces))
	return uint32(len(results.Namespaces)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCredentialsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCredentialsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCredentialsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Credentials = append(r.Credentials, results.Credentials...)
	r.TotalCount += uint32(len(results.Credentials))
	return uint32(len(results.Credentials)), nil
}

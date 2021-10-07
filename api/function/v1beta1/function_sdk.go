// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package function provides methods and message types of the function v1beta1 API.
package function

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

// API: serverless functions API
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type CronStatus string

const (
	// CronStatusUnknown is [insert doc].
	CronStatusUnknown = CronStatus("unknown")
	// CronStatusReady is [insert doc].
	CronStatusReady = CronStatus("ready")
	// CronStatusDeleting is [insert doc].
	CronStatusDeleting = CronStatus("deleting")
	// CronStatusError is [insert doc].
	CronStatusError = CronStatus("error")
	// CronStatusLocked is [insert doc].
	CronStatusLocked = CronStatus("locked")
	// CronStatusCreating is [insert doc].
	CronStatusCreating = CronStatus("creating")
	// CronStatusPending is [insert doc].
	CronStatusPending = CronStatus("pending")
)

func (enum CronStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum CronStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CronStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CronStatus(CronStatus(tmp).String())
	return nil
}

type FunctionPrivacy string

const (
	// FunctionPrivacyUnknownPrivacy is [insert doc].
	FunctionPrivacyUnknownPrivacy = FunctionPrivacy("unknown_privacy")
	// FunctionPrivacyPublic is [insert doc].
	FunctionPrivacyPublic = FunctionPrivacy("public")
	// FunctionPrivacyPrivate is [insert doc].
	FunctionPrivacyPrivate = FunctionPrivacy("private")
)

func (enum FunctionPrivacy) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_privacy"
	}
	return string(enum)
}

func (enum FunctionPrivacy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionPrivacy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionPrivacy(FunctionPrivacy(tmp).String())
	return nil
}

type FunctionRuntime string

const (
	// FunctionRuntimeUnknownRuntime is [insert doc].
	FunctionRuntimeUnknownRuntime = FunctionRuntime("unknown_runtime")
	// FunctionRuntimeGolang is [insert doc].
	FunctionRuntimeGolang = FunctionRuntime("golang")
	// FunctionRuntimePython is [insert doc].
	FunctionRuntimePython = FunctionRuntime("python")
	// FunctionRuntimePython3 is [insert doc].
	FunctionRuntimePython3 = FunctionRuntime("python3")
	// FunctionRuntimeNode8 is [insert doc].
	FunctionRuntimeNode8 = FunctionRuntime("node8")
	// FunctionRuntimeNode10 is [insert doc].
	FunctionRuntimeNode10 = FunctionRuntime("node10")
	// FunctionRuntimeNode14 is [insert doc].
	FunctionRuntimeNode14 = FunctionRuntime("node14")
)

func (enum FunctionRuntime) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_runtime"
	}
	return string(enum)
}

func (enum FunctionRuntime) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionRuntime) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionRuntime(FunctionRuntime(tmp).String())
	return nil
}

type FunctionStatus string

const (
	// FunctionStatusUnknown is [insert doc].
	FunctionStatusUnknown = FunctionStatus("unknown")
	// FunctionStatusReady is [insert doc].
	FunctionStatusReady = FunctionStatus("ready")
	// FunctionStatusDeleting is [insert doc].
	FunctionStatusDeleting = FunctionStatus("deleting")
	// FunctionStatusError is [insert doc].
	FunctionStatusError = FunctionStatus("error")
	// FunctionStatusLocked is [insert doc].
	FunctionStatusLocked = FunctionStatus("locked")
	// FunctionStatusCreating is [insert doc].
	FunctionStatusCreating = FunctionStatus("creating")
	// FunctionStatusPending is [insert doc].
	FunctionStatusPending = FunctionStatus("pending")
	// FunctionStatusCreated is [insert doc].
	FunctionStatusCreated = FunctionStatus("created")
)

func (enum FunctionStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum FunctionStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FunctionStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FunctionStatus(FunctionStatus(tmp).String())
	return nil
}

type ListCronsRequestOrderBy string

const (
	// ListCronsRequestOrderByCreatedAtAsc is [insert doc].
	ListCronsRequestOrderByCreatedAtAsc = ListCronsRequestOrderBy("created_at_asc")
	// ListCronsRequestOrderByCreatedAtDesc is [insert doc].
	ListCronsRequestOrderByCreatedAtDesc = ListCronsRequestOrderBy("created_at_desc")
)

func (enum ListCronsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListCronsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCronsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCronsRequestOrderBy(ListCronsRequestOrderBy(tmp).String())
	return nil
}

type ListFunctionsRequestOrderBy string

const (
	// ListFunctionsRequestOrderByCreatedAtAsc is [insert doc].
	ListFunctionsRequestOrderByCreatedAtAsc = ListFunctionsRequestOrderBy("created_at_asc")
	// ListFunctionsRequestOrderByCreatedAtDesc is [insert doc].
	ListFunctionsRequestOrderByCreatedAtDesc = ListFunctionsRequestOrderBy("created_at_desc")
	// ListFunctionsRequestOrderByNameAsc is [insert doc].
	ListFunctionsRequestOrderByNameAsc = ListFunctionsRequestOrderBy("name_asc")
	// ListFunctionsRequestOrderByNameDesc is [insert doc].
	ListFunctionsRequestOrderByNameDesc = ListFunctionsRequestOrderBy("name_desc")
)

func (enum ListFunctionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListFunctionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFunctionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFunctionsRequestOrderBy(ListFunctionsRequestOrderBy(tmp).String())
	return nil
}

type ListLogsRequestOrderBy string

const (
	// ListLogsRequestOrderByTimestampDesc is [insert doc].
	ListLogsRequestOrderByTimestampDesc = ListLogsRequestOrderBy("timestamp_desc")
	// ListLogsRequestOrderByTimestampAsc is [insert doc].
	ListLogsRequestOrderByTimestampAsc = ListLogsRequestOrderBy("timestamp_asc")
)

func (enum ListLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "timestamp_desc"
	}
	return string(enum)
}

func (enum ListLogsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListLogsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListLogsRequestOrderBy(ListLogsRequestOrderBy(tmp).String())
	return nil
}

type ListNamespacesRequestOrderBy string

const (
	// ListNamespacesRequestOrderByCreatedAtAsc is [insert doc].
	ListNamespacesRequestOrderByCreatedAtAsc = ListNamespacesRequestOrderBy("created_at_asc")
	// ListNamespacesRequestOrderByCreatedAtDesc is [insert doc].
	ListNamespacesRequestOrderByCreatedAtDesc = ListNamespacesRequestOrderBy("created_at_desc")
	// ListNamespacesRequestOrderByNameAsc is [insert doc].
	ListNamespacesRequestOrderByNameAsc = ListNamespacesRequestOrderBy("name_asc")
	// ListNamespacesRequestOrderByNameDesc is [insert doc].
	ListNamespacesRequestOrderByNameDesc = ListNamespacesRequestOrderBy("name_desc")
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

type NamespaceStatus string

const (
	// NamespaceStatusUnknown is [insert doc].
	NamespaceStatusUnknown = NamespaceStatus("unknown")
	// NamespaceStatusReady is [insert doc].
	NamespaceStatusReady = NamespaceStatus("ready")
	// NamespaceStatusDeleting is [insert doc].
	NamespaceStatusDeleting = NamespaceStatus("deleting")
	// NamespaceStatusError is [insert doc].
	NamespaceStatusError = NamespaceStatus("error")
	// NamespaceStatusLocked is [insert doc].
	NamespaceStatusLocked = NamespaceStatus("locked")
	// NamespaceStatusCreating is [insert doc].
	NamespaceStatusCreating = NamespaceStatus("creating")
	// NamespaceStatusPending is [insert doc].
	NamespaceStatusPending = NamespaceStatus("pending")
)

func (enum NamespaceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum NamespaceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NamespaceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NamespaceStatus(NamespaceStatus(tmp).String())
	return nil
}

type NullValue string

const (
	// NullValueNULLVALUE is [insert doc].
	NullValueNULLVALUE = NullValue("NULL_VALUE")
)

func (enum NullValue) String() string {
	if enum == "" {
		// return default value if empty
		return "NULL_VALUE"
	}
	return string(enum)
}

func (enum NullValue) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NullValue) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NullValue(NullValue(tmp).String())
	return nil
}

// Cron: cron
type Cron struct {
	ID string `json:"id"`

	FunctionID string `json:"function_id"`

	Schedule string `json:"schedule"`

	Args []byte `json:"args"`
	// Status:
	//
	// Default value: unknown
	Status CronStatus `json:"status"`
}

type DownloadURL struct {
	URL string `json:"url"`

	Headers map[string]*[]string `json:"headers"`
}

// Function: function
type Function struct {
	ID string `json:"id"`

	Name string `json:"name"`

	NamespaceID string `json:"namespace_id"`
	// Status:
	//
	// Default value: unknown
	Status FunctionStatus `json:"status"`

	EnvironmentVariables map[string]string `json:"environment_variables"`

	MinScale uint32 `json:"min_scale"`

	MaxScale uint32 `json:"max_scale"`
	// Runtime:
	//
	// Default value: unknown_runtime
	Runtime FunctionRuntime `json:"runtime"`

	MemoryLimit uint32 `json:"memory_limit"`

	CPULimit uint32 `json:"cpu_limit"`

	Timeout *scw.Duration `json:"timeout"`

	Handler string `json:"handler"`

	ErrorMessage *string `json:"error_message"`
	// Privacy:
	//
	// Default value: unknown_privacy
	Privacy FunctionPrivacy `json:"privacy"`

	Description *string `json:"description"`

	DomainName string `json:"domain_name"`

	Region scw.Region `json:"region"`
}

// ListCronsResponse: list crons response
type ListCronsResponse struct {
	Crons []*Cron `json:"crons"`

	TotalCount uint32 `json:"total_count"`
}

// ListFunctionRuntimesResponse: list function runtimes response
type ListFunctionRuntimesResponse struct {
	Runtimes []FunctionRuntime `json:"runtimes"`
}

// ListFunctionsResponse: list functions response
type ListFunctionsResponse struct {
	Functions []*Function `json:"functions"`

	TotalCount uint32 `json:"total_count"`
}

// ListLogsResponse: list logs response
type ListLogsResponse struct {
	Logs []*Log `json:"logs"`

	TotalCount uint32 `json:"total_count"`
}

// ListNamespacesResponse: list namespaces response
type ListNamespacesResponse struct {
	Namespaces []*Namespace `json:"namespaces"`

	TotalCount uint32 `json:"total_count"`
}

// Log: log
type Log struct {
	Message string `json:"message"`

	Timestamp *time.Time `json:"timestamp"`

	ID string `json:"id"`
}

// Namespace: namespace
type Namespace struct {
	ID string `json:"id"`

	Name string `json:"name"`

	EnvironmentVariables map[string]string `json:"environment_variables"`

	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`
	// Status:
	//
	// Default value: unknown
	Status NamespaceStatus `json:"status"`

	RegistryNamespaceID string `json:"registry_namespace_id"`

	ErrorMessage *string `json:"error_message"`

	RegistryEndpoint string `json:"registry_endpoint"`

	Description *string `json:"description"`

	Region scw.Region `json:"region"`
}

type Token struct {
	Token string `json:"token"`

	PublicKey string `json:"public_key"`
}

// UploadURL: upload url
type UploadURL struct {
	URL string `json:"url"`

	Headers map[string]*[]string `json:"headers"`
}

// Service API

type ListNamespacesRequest struct {
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListNamespacesRequestOrderBy `json:"-"`

	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// ListNamespaces: list all your namespaces
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
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

type GetNamespaceRequest struct {
	Region scw.Region `json:"-"`

	NamespaceID string `json:"-"`
}

// GetNamespace: get a namespace
//
// Get the namespace associated with the given id.
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
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
		Headers: http.Header{},
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateNamespaceRequest struct {
	Region scw.Region `json:"-"`

	Name string `json:"name"`

	EnvironmentVariables *map[string]string `json:"environment_variables"`

	ProjectID string `json:"project_id"`

	Description *string `json:"description"`
}

// CreateNamespace: create a new namespace
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces",
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
	Region scw.Region `json:"-"`

	NamespaceID string `json:"-"`

	EnvironmentVariables *map[string]string `json:"environment_variables"`

	Description *string `json:"description"`
}

// UpdateNamespace: update an existing namespace
//
// Update the space associated with the given id.
func (s *API) UpdateNamespace(req *UpdateNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
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
		Method:  "PATCH",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
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

type DeleteNamespaceRequest struct {
	Region scw.Region `json:"-"`

	NamespaceID string `json:"-"`
}

// DeleteNamespace: delete an existing namespace
//
// Delete the namespace associated with the given id.
func (s *API) DeleteNamespace(req *DeleteNamespaceRequest, opts ...scw.RequestOption) (*Namespace, error) {
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
		Method:  "DELETE",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/namespaces/" + fmt.Sprint(req.NamespaceID) + "",
		Headers: http.Header{},
	}

	var resp Namespace

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListFunctionsRequest struct {
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListFunctionsRequestOrderBy `json:"-"`

	NamespaceID string `json:"-"`

	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// ListFunctions: list all your functions
func (s *API) ListFunctions(req *ListFunctionsRequest, opts ...scw.RequestOption) (*ListFunctionsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListFunctionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetFunctionRequest struct {
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`
}

// GetFunction: get a function
//
// Get the function associated with the given id.
func (s *API) GetFunction(req *GetFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
		Headers: http.Header{},
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateFunctionRequest struct {
	Region scw.Region `json:"-"`

	NamespaceID string `json:"namespace_id"`

	Name string `json:"name"`

	EnvironmentVariables *map[string]string `json:"environment_variables"`

	MinScale *uint32 `json:"min_scale"`

	MaxScale *uint32 `json:"max_scale"`
	// Runtime:
	//
	// Default value: unknown_runtime
	Runtime FunctionRuntime `json:"runtime"`

	MemoryLimit *uint32 `json:"memory_limit"`

	Timeout *scw.Duration `json:"timeout"`

	Handler *string `json:"handler"`
	// Privacy:
	//
	// Default value: unknown_privacy
	Privacy FunctionPrivacy `json:"privacy"`

	Description *string `json:"description"`

	DomainName *string `json:"domain_name"`
}

// CreateFunction: create a new function
func (s *API) CreateFunction(req *CreateFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateFunctionRequest struct {
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`

	EnvironmentVariables *map[string]string `json:"environment_variables"`

	MinScale *uint32 `json:"min_scale"`

	MaxScale *uint32 `json:"max_scale"`

	MemoryLimit *uint32 `json:"memory_limit"`

	Timeout *scw.Duration `json:"timeout"`

	Redeploy *bool `json:"redeploy"`

	Handler *string `json:"handler"`
	// Privacy:
	//
	// Default value: unknown_privacy
	Privacy FunctionPrivacy `json:"privacy"`

	Description *string `json:"description"`

	DomainName *string `json:"domain_name"`
}

// UpdateFunction: update an existing function
//
// Update the function associated with the given id.
func (s *API) UpdateFunction(req *UpdateFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteFunctionRequest struct {
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`
}

// DeleteFunction: delete a function
//
// Delete the function associated with the given id.
func (s *API) DeleteFunction(req *DeleteFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "",
		Headers: http.Header{},
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeployFunctionRequest struct {
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`
}

func (s *API) DeployFunction(req *DeployFunctionRequest, opts ...scw.RequestOption) (*Function, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/deploy",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Function

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListFunctionRuntimesRequest struct {
	Region scw.Region `json:"-"`
}

// ListFunctionRuntimes: list function runtimes
//
// List available function runtimes.
func (s *API) ListFunctionRuntimes(req *ListFunctionRuntimesRequest, opts ...scw.RequestOption) (*ListFunctionRuntimesResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/runtimes",
		Headers: http.Header{},
	}

	var resp ListFunctionRuntimesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetFunctionUploadURLRequest struct {
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`

	ContentLength uint64 `json:"-"`
}

func (s *API) GetFunctionUploadURL(req *GetFunctionUploadURLRequest, opts ...scw.RequestOption) (*UploadURL, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "content_length", req.ContentLength)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/upload-url",
		Query:   query,
		Headers: http.Header{},
	}

	var resp UploadURL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetFunctionDownloadURLRequest struct {
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`
}

func (s *API) GetFunctionDownloadURL(req *GetFunctionDownloadURLRequest, opts ...scw.RequestOption) (*DownloadURL, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/download-url",
		Headers: http.Header{},
	}

	var resp DownloadURL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListCronsRequest struct {
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListCronsRequestOrderBy `json:"-"`

	FunctionID string `json:"-"`
}

// ListCrons: list all your crons
func (s *API) ListCrons(req *ListCronsRequest, opts ...scw.RequestOption) (*ListCronsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "function_id", req.FunctionID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListCronsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetCronRequest struct {
	Region scw.Region `json:"-"`

	CronID string `json:"-"`
}

// GetCron: get a cron
//
// Get the cron associated with the given id.
func (s *API) GetCron(req *GetCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CronID) == "" {
		return nil, errors.New("field CronID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
		Headers: http.Header{},
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateCronRequest struct {
	Region scw.Region `json:"-"`

	FunctionID string `json:"function_id"`

	Schedule string `json:"schedule"`

	Args []byte `json:"args"`
}

// CreateCron: create a new cron
func (s *API) CreateCron(req *CreateCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateCronRequest struct {
	Region scw.Region `json:"-"`

	CronID string `json:"-"`

	FunctionID *string `json:"function_id"`

	Schedule *string `json:"schedule"`

	Args []byte `json:"args"`
}

// UpdateCron: update an existing cron
//
// Update the cron associated with the given id.
func (s *API) UpdateCron(req *UpdateCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CronID) == "" {
		return nil, errors.New("field CronID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteCronRequest struct {
	Region scw.Region `json:"-"`

	CronID string `json:"-"`
}

// DeleteCron: delete an existing cron
//
// Delete the cron associated with the given id.
func (s *API) DeleteCron(req *DeleteCronRequest, opts ...scw.RequestOption) (*Cron, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.CronID) == "" {
		return nil, errors.New("field CronID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/crons/" + fmt.Sprint(req.CronID) + "",
		Headers: http.Header{},
	}

	var resp Cron

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListLogsRequest struct {
	Region scw.Region `json:"-"`

	FunctionID string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: timestamp_desc
	OrderBy ListLogsRequestOrderBy `json:"-"`
}

// ListLogs: list your application logs
func (s *API) ListLogs(req *ListLogsRequest, opts ...scw.RequestOption) (*ListLogsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FunctionID) == "" {
		return nil, errors.New("field FunctionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/functions/" + fmt.Sprint(req.FunctionID) + "/logs",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type IssueJWTRequest struct {
	Region scw.Region `json:"-"`

	FunctionID *string `json:"-"`

	NamespaceID *string `json:"-"`

	ExpiresAt *time.Time `json:"-"`
}

func (s *API) IssueJWT(req *IssueJWTRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "function_id", req.FunctionID)
	parameter.AddToQuery(query, "namespace_id", req.NamespaceID)
	parameter.AddToQuery(query, "expires_at", req.ExpiresAt)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/functions/v1beta1/regions/" + fmt.Sprint(req.Region) + "/issue-jwt",
		Query:   query,
		Headers: http.Header{},
	}

	var resp Token

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
func (r *ListFunctionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFunctionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListFunctionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Functions = append(r.Functions, results.Functions...)
	r.TotalCount += uint32(len(results.Functions))
	return uint32(len(results.Functions)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCronsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCronsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCronsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Crons = append(r.Crons, results.Crons...)
	r.TotalCount += uint32(len(results.Crons))
	return uint32(len(results.Crons)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListLogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Logs = append(r.Logs, results.Logs...)
	r.TotalCount += uint32(len(results.Logs))
	return uint32(len(results.Logs)), nil
}

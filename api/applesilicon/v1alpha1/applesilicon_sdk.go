// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package applesilicon provides methods and message types of the applesilicon v1alpha1 API.
package applesilicon

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

// API: apple Mac mini M1 as a service
//
// Scaleway Apple silicon M1 as-a-Service is built using the latest generation of Apple Mac mini hardware (fifth generation).
//
// These dedicated Mac mini M1s are designed for developing, building, testing, and signing applications for Apple devices, including iPhones, iPads, Mac computers and much more.
//
// Get set to explore, learn and build on a dedicated Mac mini M1 with more performance and speed than you ever thought possible.
//
// **Apple silicon as a Service comes with a minimum allocation period of 24 hours**.
//
// Mac mini and macOS are trademarks of Apple Inc., registered in the U.S. and other countries and regions.
// IOS is a trademark or registered trademark of Cisco in the U.S. and other countries and is used by Apple under license.
// Scaleway is not affiliated with Apple Inc.
//
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type ListServersRequestOrderBy string

const (
	// ListServersRequestOrderByCreatedAtAsc is [insert doc].
	ListServersRequestOrderByCreatedAtAsc = ListServersRequestOrderBy("created_at_asc")
	// ListServersRequestOrderByCreatedAtDesc is [insert doc].
	ListServersRequestOrderByCreatedAtDesc = ListServersRequestOrderBy("created_at_desc")
)

func (enum ListServersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListServersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListServersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListServersRequestOrderBy(ListServersRequestOrderBy(tmp).String())
	return nil
}

type ServerStatus string

const (
	// ServerStatusUnknownStatus is [insert doc].
	ServerStatusUnknownStatus = ServerStatus("unknown_status")
	// ServerStatusStarting is [insert doc].
	ServerStatusStarting = ServerStatus("starting")
	// ServerStatusReady is [insert doc].
	ServerStatusReady = ServerStatus("ready")
	// ServerStatusError is [insert doc].
	ServerStatusError = ServerStatus("error")
	// ServerStatusRebooting is [insert doc].
	ServerStatusRebooting = ServerStatus("rebooting")
	// ServerStatusUpdating is [insert doc].
	ServerStatusUpdating = ServerStatus("updating")
	// ServerStatusLocking is [insert doc].
	ServerStatusLocking = ServerStatus("locking")
	// ServerStatusLocked is [insert doc].
	ServerStatusLocked = ServerStatus("locked")
	// ServerStatusUnlocking is [insert doc].
	ServerStatusUnlocking = ServerStatus("unlocking")
	// ServerStatusReinstalling is [insert doc].
	ServerStatusReinstalling = ServerStatus("reinstalling")
)

func (enum ServerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum ServerStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerStatus(ServerStatus(tmp).String())
	return nil
}

type ServerTypeStock string

const (
	// ServerTypeStockUnknownStock is [insert doc].
	ServerTypeStockUnknownStock = ServerTypeStock("unknown_stock")
	// ServerTypeStockNoStock is [insert doc].
	ServerTypeStockNoStock = ServerTypeStock("no_stock")
	// ServerTypeStockLowStock is [insert doc].
	ServerTypeStockLowStock = ServerTypeStock("low_stock")
	// ServerTypeStockHighStock is [insert doc].
	ServerTypeStockHighStock = ServerTypeStock("high_stock")
)

func (enum ServerTypeStock) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_stock"
	}
	return string(enum)
}

func (enum ServerTypeStock) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServerTypeStock) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServerTypeStock(ServerTypeStock(tmp).String())
	return nil
}

// ListOSResponse: list os response
type ListOSResponse struct {
	// TotalCount: total number of os
	TotalCount uint32 `json:"total_count"`
	// Os: list of OS
	Os []*OS `json:"os"`
}

// ListServerTypesResponse: list server types response
type ListServerTypesResponse struct {
	// ServerTypes: the available server types
	ServerTypes []*ServerType `json:"server_types"`
}

// ListServersResponse: list servers response
type ListServersResponse struct {
	// TotalCount: the total number of servers
	TotalCount uint32 `json:"total_count"`
	// Servers: the paginated returned servers
	Servers []*Server `json:"servers"`
}

// OS: os
type OS struct {
	// ID: the OS unique ID
	ID string `json:"id"`
	// Name: the OS name
	Name string `json:"name"`
	// Label: the OS name as it should be displayed
	Label string `json:"label"`
	// ImageURL: URL of the image
	ImageURL string `json:"image_url"`
	// CompatibleServerTypes: list of compatible server types
	CompatibleServerTypes []string `json:"compatible_server_types"`
}

// Server: server
type Server struct {
	// ID: UUID of the server
	ID string `json:"id"`
	// Type: type of the server
	Type string `json:"type"`
	// Name: name of the server
	Name string `json:"name"`
	// ProjectID: project this server is associated with
	ProjectID string `json:"project_id"`
	// OrganizationID: organization this server is associated with
	OrganizationID string `json:"organization_id"`
	// IP: iPv4 address of the server
	IP net.IP `json:"ip"`
	// VncURL: URL of the VNC
	VncURL string `json:"vnc_url"`
	// Status: current status of the server
	//
	// Default value: unknown_status
	Status ServerStatus `json:"status"`
	// CreatedAt: the date at which the server was created
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: the date at which the server was last updated
	UpdatedAt *time.Time `json:"updated_at"`
	// DeletableAt: the date at which the server was last deleted
	DeletableAt *time.Time `json:"deletable_at"`
	// Zone: the zone of the server
	Zone scw.Zone `json:"zone"`
}

// ServerType: server type
type ServerType struct {
	// CPU: CPU description
	CPU *ServerTypeCPU `json:"cpu"`
	// Disk: size of the local disk of the server
	Disk *ServerTypeDisk `json:"disk"`
	// Name: name of the type
	Name string `json:"name"`
	// Memory: size of memory available
	Memory *ServerTypeMemory `json:"memory"`
	// Stock: current stock
	//
	// Default value: unknown_stock
	Stock ServerTypeStock `json:"stock"`
	// MinimumLeaseDuration: minimum duration of the lease in seconds
	//
	// Minimum duration of the lease in seconds (example. 3.4s).
	MinimumLeaseDuration *scw.Duration `json:"minimum_lease_duration"`
}

type ServerTypeCPU struct {
	Name string `json:"name"`

	CoreCount uint32 `json:"core_count"`
}

type ServerTypeDisk struct {
	Capacity scw.Size `json:"capacity"`

	Type string `json:"type"`
}

type ServerTypeMemory struct {
	Capacity scw.Size `json:"capacity"`

	Type string `json:"type"`
}

// Service API

// Zones list localities the api is available in
func (s *API) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar3}
}

type ListServerTypesRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
}

// ListServerTypes: list server types
//
// List all server types technical details.
func (s *API) ListServerTypes(req *ListServerTypesRequest, opts ...scw.RequestOption) (*ListServerTypesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/server-types",
		Headers: http.Header{},
	}

	var resp ListServerTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetServerTypeRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ServerType: server type identifier
	ServerType string `json:"-"`
}

// GetServerType: get a server type
//
// Get a server technical details.
func (s *API) GetServerType(req *GetServerTypeRequest, opts ...scw.RequestOption) (*ServerType, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerType) == "" {
		return nil, errors.New("field ServerType cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/server-type/" + fmt.Sprint(req.ServerType) + "",
		Headers: http.Header{},
	}

	var resp ServerType

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateServerRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// Name: create a server with this given name
	Name string `json:"name"`
	// ProjectID: create a server in the given project ID
	ProjectID string `json:"project_id"`
	// Type: create a server of the given type
	Type string `json:"type"`
}

// CreateServer: create a server
//
// Create a server.
func (s *API) CreateServer(req *CreateServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("as")
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListServersRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// OrderBy: the sort order of the returned servers
	//
	// Default value: created_at_asc
	OrderBy ListServersRequestOrderBy `json:"-"`
	// ProjectID: list only servers of this project ID
	ProjectID *string `json:"-"`
	// OrganizationID: list only servers of this organization ID
	OrganizationID *string `json:"-"`
	// Page: a positive integer to choose the page to return
	Page *int32 `json:"-"`
	// PageSize: a positive integer lower or equal to 100 to select the number of items to return
	//
	// Default value: 50
	PageSize *uint32 `json:"-"`
}

// ListServers: list all servers
//
// List all servers.
func (s *API) ListServers(req *ListServersRequest, opts ...scw.RequestOption) (*ListServersResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListOSRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// Page: a positive integer to choose the page to return
	Page *int32 `json:"-"`
	// PageSize: a positive integer lower or equal to 100 to select the number of items to return
	//
	// Default value: 50
	PageSize *uint32 `json:"-"`
	// ServerType: list of compatible server type
	ServerType *string `json:"-"`
	// Name: filter os by name (for eg. "11.1" will return "11.1.2" and "11.1" but not "12")
	Name *string `json:"-"`
}

// ListOS: list all Operating System (OS)
//
// List all Operating System (OS).
func (s *API) ListOS(req *ListOSRequest, opts ...scw.RequestOption) (*ListOSResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "server_type", req.ServerType)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/os",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListOSResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetOSRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// OsID: UUID of the OS you want to get
	OsID string `json:"-"`
}

// GetOS: get an Operating System (OS)
//
// Get an Operating System (OS).
func (s *API) GetOS(req *GetOSRequest, opts ...scw.RequestOption) (*OS, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.OsID) == "" {
		return nil, errors.New("field OsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/os/" + fmt.Sprint(req.OsID) + "",
		Headers: http.Header{},
	}

	var resp OS

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetServerRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to get
	ServerID string `json:"-"`
}

// GetServer: get a server
//
// Get a server.
func (s *API) GetServer(req *GetServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
		Headers: http.Header{},
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateServerRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to update
	ServerID string `json:"-"`
	// Name: updated name for your server
	Name *string `json:"name"`
}

// UpdateServer: update a server
//
// Update a server.
func (s *API) UpdateServer(req *UpdateServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteServerRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to delete
	ServerID string `json:"-"`
}

// DeleteServer: delete a server
//
// Delete a server.
func (s *API) DeleteServer(req *DeleteServerRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type RebootServerRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to reboot
	ServerID string `json:"-"`
}

// RebootServer: reboot a server
//
// Reboot a server.
func (s *API) RebootServer(req *RebootServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/reboot",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ReinstallServerRequest struct {
	// Zone:
	//
	// Zone to target. If none is passed will use default zone from the config
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to reinstall
	ServerID string `json:"-"`
}

// ReinstallServer: reinstall a server
//
// Reinstall a server.
func (s *API) ReinstallServer(req *ReinstallServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/apple-silicon/v1alpha1/zones/" + fmt.Sprint(req.Zone) + "/servers/" + fmt.Sprint(req.ServerID) + "/reinstall",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint32(len(results.Servers))
	return uint32(len(results.Servers)), nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOSResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOSResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListOSResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Os = append(r.Os, results.Os...)
	r.TotalCount += uint32(len(results.Os))
	return uint32(len(results.Os)), nil
}

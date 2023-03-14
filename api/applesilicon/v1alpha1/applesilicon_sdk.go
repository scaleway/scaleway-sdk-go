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

// API: apple Mac mini as a service.
// Scaleway Apple silicon as-a-Service is built using the latest generation of Apple Mac mini hardware (fifth generation).
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
	ListServersRequestOrderByCreatedAtAsc  = ListServersRequestOrderBy("created_at_asc")
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
	ServerStatusUnknownStatus = ServerStatus("unknown_status")
	ServerStatusStarting      = ServerStatus("starting")
	ServerStatusReady         = ServerStatus("ready")
	ServerStatusError         = ServerStatus("error")
	ServerStatusRebooting     = ServerStatus("rebooting")
	ServerStatusUpdating      = ServerStatus("updating")
	ServerStatusLocking       = ServerStatus("locking")
	ServerStatusLocked        = ServerStatus("locked")
	ServerStatusUnlocking     = ServerStatus("unlocking")
	ServerStatusReinstalling  = ServerStatus("reinstalling")
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
	ServerTypeStockUnknownStock = ServerTypeStock("unknown_stock")
	ServerTypeStockNoStock      = ServerTypeStock("no_stock")
	ServerTypeStockLowStock     = ServerTypeStock("low_stock")
	ServerTypeStockHighStock    = ServerTypeStock("high_stock")
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

// ListOSResponse: list os response.
type ListOSResponse struct {
	// TotalCount: total number of OS.
	TotalCount uint32 `json:"total_count"`
	// Os: list of OS.
	Os []*OS `json:"os"`
}

// ListServerTypesResponse: list server types response.
type ListServerTypesResponse struct {
	// ServerTypes: available server types.
	ServerTypes []*ServerType `json:"server_types"`
}

// ListServersResponse: list servers response.
type ListServersResponse struct {
	// TotalCount: total number of servers.
	TotalCount uint32 `json:"total_count"`
	// Servers: paginated returned servers.
	Servers []*Server `json:"servers"`
}

// OS: os.
type OS struct {
	// ID: unique ID of the OS.
	ID string `json:"id"`
	// Name: oS name.
	Name string `json:"name"`
	// Label: oS name as it should be displayed.
	Label string `json:"label"`
	// ImageURL: URL of the image.
	ImageURL string `json:"image_url"`
	// CompatibleServerTypes: list of compatible server types.
	CompatibleServerTypes []string `json:"compatible_server_types"`
}

// Server: server.
type Server struct {
	// ID: UUID of the server.
	ID string `json:"id"`
	// Type: type of the server.
	Type string `json:"type"`
	// Name: name of the server.
	Name string `json:"name"`
	// ProjectID: project this server is associated with.
	ProjectID string `json:"project_id"`
	// OrganizationID: organization this server is associated with.
	OrganizationID string `json:"organization_id"`
	// IP: iPv4 address of the server.
	IP net.IP `json:"ip"`
	// VncURL: URL of the VNC.
	VncURL string `json:"vnc_url"`
	// Status: current status of the server.
	// Default value: unknown_status
	Status ServerStatus `json:"status"`
	// CreatedAt: date on which the server was created.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: date on which the server was last updated.
	UpdatedAt *time.Time `json:"updated_at"`
	// DeletableAt: date on which the server was last deleted.
	DeletableAt *time.Time `json:"deletable_at"`
	// Zone: zone of the server.
	Zone scw.Zone `json:"zone"`
}

// ServerType: server type.
type ServerType struct {
	// CPU: CPU description.
	CPU *ServerTypeCPU `json:"cpu"`
	// Disk: size of the local disk of the server.
	Disk *ServerTypeDisk `json:"disk"`
	// Name: name of the type.
	Name string `json:"name"`
	// Memory: size of memory available.
	Memory *ServerTypeMemory `json:"memory"`
	// Stock: current stock.
	// Default value: unknown_stock
	Stock ServerTypeStock `json:"stock"`
	// MinimumLeaseDuration: minimum duration of the lease in seconds.
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
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// ListServerTypes: list server types.
// List all technical details about Apple silicon server types available in the specified zone. Since there is only one Availability Zone for Apple silicon servers, the targeted value is `fr-par-3`.
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
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ServerType: server type identifier.
	ServerType string `json:"-"`
}

// GetServerType: get a server type.
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
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// Name: create a server with this given name.
	Name string `json:"name"`
	// ProjectID: create a server in the given project ID.
	ProjectID string `json:"project_id"`
	// Type: create a server of the given type.
	Type string `json:"type"`
}

// CreateServer: create a server.
// Create a new server in the targeted zone, specifying its configuration including name and type.
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
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// OrderBy: sort order of the returned servers.
	// Default value: created_at_asc
	OrderBy ListServersRequestOrderBy `json:"-"`
	// ProjectID: only list servers of this project ID.
	ProjectID *string `json:"-"`
	// OrganizationID: only list servers of this Organization ID.
	OrganizationID *string `json:"-"`
	// Page: positive integer to choose the page to return.
	Page *int32 `json:"-"`
	// PageSize: positive integer lower or equal to 100 to select the number of items to return.
	// Default value: 50
	PageSize *uint32 `json:"-"`
}

// ListServers: list all servers.
// List all servers in the specified zone. By default, returned servers in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
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
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// Page: positive integer to choose the page to return.
	Page *int32 `json:"-"`
	// PageSize: positive integer lower or equal to 100 to select the number of items to return.
	// Default value: 50
	PageSize *uint32 `json:"-"`
	// ServerType: list of compatible server types.
	ServerType *string `json:"-"`
	// Name: filter OS by name (note that "11.1" will return "11.1.2" and "11.1" but not "12")).
	Name *string `json:"-"`
}

// ListOS: list all Operating System (OS).
// List all Operating System (OS). The response will include the total number of OS as well as their associated IDs, names and labels.
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
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// OsID: UUID of the OS you want to get.
	OsID string `json:"-"`
}

// GetOS: get an Operating System (OS).
// Get an Operating System (OS).  The response will include the OS's unique ID as well as its name and label.
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
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to get.
	ServerID string `json:"-"`
}

// GetServer: get a server.
// Retrieve information about an existing Apple silicon server, specified by its server ID. Its full details, including name, status and IP address, are returned in the response object.
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
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to update.
	ServerID string `json:"-"`
	// Name: updated name for your server.
	Name *string `json:"name"`
}

// UpdateServer: update a server.
// Update the parameters of an existing Apple silicon server, specified by its server ID.
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
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to delete.
	ServerID string `json:"-"`
}

// DeleteServer: delete a server.
// Delete an existing Apple silicon server, specified by its server ID. Deleting a server is permanent, and cannot be undone. Note that the minimum allocation period for Apple silicon-as-a-service is 24 hours, meaning you cannot delete your server prior to that.
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
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to reboot.
	ServerID string `json:"-"`
}

// RebootServer: reboot a server.
// Reboot an existing Apple silicon server, specified by its server ID.
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
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
	// ServerID: UUID of the server you want to reinstall.
	ServerID string `json:"-"`
}

// ReinstallServer: reinstall a server.
// Reinstall an existing Apple silicon server (specified by its server ID) from a new image (OS). All the data on the disk is deleted and all configuration is reset to the defailt configuration values of the image (OS).
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

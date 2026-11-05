// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package vps provides methods and message types of the vps v1 API.
package vps

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
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultVpsRetryInterval = 15 * time.Second
	defaultVpsTimeout       = 5 * time.Minute
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

type ManageServerPowerRequestAction string

const (
	ManageServerPowerRequestActionUnknownAction = ManageServerPowerRequestAction("unknown_action")
	ManageServerPowerRequestActionPowerOff      = ManageServerPowerRequestAction("power_off")
	ManageServerPowerRequestActionPowerOn       = ManageServerPowerRequestAction("power_on")
	ManageServerPowerRequestActionReboot        = ManageServerPowerRequestAction("reboot")
)

func (enum ManageServerPowerRequestAction) String() string {
	if enum == "" {
		// return default value if empty
		return string(ManageServerPowerRequestActionUnknownAction)
	}
	return string(enum)
}

func (enum ManageServerPowerRequestAction) Values() []ManageServerPowerRequestAction {
	return []ManageServerPowerRequestAction{
		"unknown_action",
		"power_off",
		"power_on",
		"reboot",
	}
}

func (enum ManageServerPowerRequestAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ManageServerPowerRequestAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ManageServerPowerRequestAction(ManageServerPowerRequestAction(tmp).String())
	return nil
}

type ServerStatus string

const (
	ServerStatusUnknownStatus = ServerStatus("unknown_status")
	ServerStatusOrdered       = ServerStatus("ordered")
	ServerStatusInstalling    = ServerStatus("installing")
	ServerStatusActive        = ServerStatus("active")
	ServerStatusStopped       = ServerStatus("stopped")
	ServerStatusLocked        = ServerStatus("locked")
	ServerStatusSuspended     = ServerStatus("suspended")
	ServerStatusStopping      = ServerStatus("stopping")
	ServerStatusStarting      = ServerStatus("starting")
	ServerStatusDeleting      = ServerStatus("deleting")
)

func (enum ServerStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServerStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ServerStatus) Values() []ServerStatus {
	return []ServerStatus{
		"unknown_status",
		"ordered",
		"installing",
		"active",
		"stopped",
		"locked",
		"suspended",
		"stopping",
		"starting",
		"deleting",
	}
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

// HardwareSpecs: hardware specs.
type HardwareSpecs struct {
	// CoreCount: number of cores.
	CoreCount int32 `json:"core_count"`

	// StorageBytes: storage size.
	StorageBytes int64 `json:"storage_bytes"`

	// RAMBytes: RAM size.
	RAMBytes int64 `json:"ram_bytes"`
}

// NetworkSpecs: network specs.
type NetworkSpecs struct {
	// BandwidthMax: network bandwidth (GB).
	BandwidthMax int64 `json:"bandwidth_max"`

	// BandwidthSpeedGbps: bandwidth speed.
	BandwidthSpeedGbps int64 `json:"bandwidth_speed_gbps"`
}

// Location: location.
type Location struct {
	// ID: iD.
	ID string `json:"id"`

	// GeographicArea: geographic area.
	GeographicArea string `json:"geographic_area"`

	// CountryCode: country code.
	CountryCode string `json:"country_code"`

	// CityName: city name.
	CityName string `json:"city_name"`

	// AvailableOfferIDs: available offer IDs.
	AvailableOfferIDs []string `json:"available_offer_ids"`

	// Srn: sRN of the location.
	Srn string `json:"srn"`
}

//nolint:unused,nolintlint // generated, no check on whether it's used or not
func (m *Location) setSRN(platform string) {
	if m.Srn != "" {
		// if the field is set server-side, trust the server
		return
	}

	// We do not check that *m.XYZ != "", as there are currently no use cases for an
	// optional value in an SRN where the value set to the empty string makes sense.

	if fmt.Sprint(m.ID) != "" {
		m.Srn = fmt.Sprintf("srn://vps.%s/locations/%s", platform, fmt.Sprint(m.ID))
		return
	}
}

// OS: os.
type OS struct {
	// ID: iD.
	ID string `json:"id"`

	// Name: name.
	Name string `json:"name"`

	// MainOs: oS family.
	MainOs string `json:"main_os"`

	// Description: description.
	Description *string `json:"description"`

	// Srn: sRN of the OS.
	Srn string `json:"srn"`
}

//nolint:unused,nolintlint // generated, no check on whether it's used or not
func (m *OS) setSRN(platform string) {
	if m.Srn != "" {
		// if the field is set server-side, trust the server
		return
	}

	// We do not check that *m.XYZ != "", as there are currently no use cases for an
	// optional value in an SRN where the value set to the empty string makes sense.

	if fmt.Sprint(m.ID) != "" {
		m.Srn = fmt.Sprintf("srn://vps.%s/os/%s", platform, fmt.Sprint(m.ID))
		return
	}
}

// Offer: offer.
type Offer struct {
	// ID: ID of the offer.
	ID string `json:"id"`

	// Name: name of the offer.
	Name string `json:"name"`

	// Available: whether the offer is available.
	Available bool `json:"available"`

	// HardwareSpecs: hardware specifications.
	HardwareSpecs *HardwareSpecs `json:"hardware_specs"`

	// NetworkSpecs: network specifications.
	NetworkSpecs *NetworkSpecs `json:"network_specs"`

	// Srn: sRN of the offer.
	Srn string `json:"srn"`
}

//nolint:unused,nolintlint // generated, no check on whether it's used or not
func (m *Offer) setSRN(platform string) {
	if m.Srn != "" {
		// if the field is set server-side, trust the server
		return
	}

	// We do not check that *m.XYZ != "", as there are currently no use cases for an
	// optional value in an SRN where the value set to the empty string makes sense.

	if fmt.Sprint(m.ID) != "" {
		m.Srn = fmt.Sprintf("srn://vps.%s/offers/%s", platform, fmt.Sprint(m.ID))
		return
	}
}

// Server: server.
type Server struct {
	// ID: ID of the server.
	ID string `json:"id"`

	// Hostname: hostname of the server.
	Hostname string `json:"hostname"`

	// OsID: oS ID installed on server.
	OsID string `json:"os_id"`

	// LocationID: location ID of the server.
	LocationID string `json:"location_id"`

	// Status: status of the server.
	// Default value: unknown_status
	Status ServerStatus `json:"status"`

	// IPs: list of IPs of the server.
	IPs []string `json:"ips"`

	// OrganizationID: organization ID owning the server.
	OrganizationID string `json:"organization_id"`

	// ProjectID: project ID of the server.
	ProjectID string `json:"project_id"`

	// OfferID: offer ID of the server.
	OfferID string `json:"offer_id"`

	// Srn: sRN of the server.
	Srn string `json:"srn"`
}

//nolint:unused,nolintlint // generated, no check on whether it's used or not
func (m *Server) setSRN(platform string) {
	if m.Srn != "" {
		// if the field is set server-side, trust the server
		return
	}

	// We do not check that *m.XYZ != "", as there are currently no use cases for an
	// optional value in an SRN where the value set to the empty string makes sense.

	if fmt.Sprint(m.ID) != "" {
		m.Srn = fmt.Sprintf("srn://vps.%s/servers/%s", platform, fmt.Sprint(m.ID))
		return
	}
}

// ListLocationsResponse: list locations response.
type ListLocationsResponse struct {
	// TotalCount: total count.
	TotalCount uint64 `json:"total_count"`

	// Locations: list of locations available for the given offer.
	Locations []*Location `json:"locations"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLocationsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLocationsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListLocationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Locations = append(r.Locations, results.Locations...)
	r.TotalCount += uint64(len(results.Locations))
	return uint64(len(results.Locations)), nil
}

// ListOSResponse: list os response.
type ListOSResponse struct {
	// TotalCount: total count.
	TotalCount uint64 `json:"total_count"`

	// Os: list of OS.
	Os []*OS `json:"os"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOSResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOSResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListOSResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Os = append(r.Os, results.Os...)
	r.TotalCount += uint64(len(results.Os))
	return uint64(len(results.Os)), nil
}

// ListOffersResponse: list offers response.
type ListOffersResponse struct {
	// TotalCount: total count of matching offers.
	TotalCount uint64 `json:"total_count"`

	// Offers: list of offers.
	Offers []*Offer `json:"offers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOffersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOffersResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListOffersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Offers = append(r.Offers, results.Offers...)
	r.TotalCount += uint64(len(results.Offers))
	return uint64(len(results.Offers)), nil
}

// ListServersResponse: list servers response.
type ListServersResponse struct {
	// TotalCount: total count of matching servers.
	TotalCount uint64 `json:"total_count"`

	// Servers: servers that match filters.
	Servers []*Server `json:"servers"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListServersResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListServersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Servers = append(r.Servers, results.Servers...)
	r.TotalCount += uint64(len(results.Servers))
	return uint64(len(results.Servers)), nil
}

// VpsAPIChangeServerHostnameRequest: vps api change server hostname request.
type VpsAPIChangeServerHostnameRequest struct {
	// ServerID: ID of the server.
	ServerID string `json:"-"`

	// Hostname: new hostname of the server.
	Hostname string `json:"hostname"`
}

// VpsAPIChangeServerPasswordRequest: vps api change server password request.
type VpsAPIChangeServerPasswordRequest struct {
	// ServerID: ID of the server.
	ServerID string `json:"-"`

	// Password: new root password of the server.
	Password string `json:"password"`
}

// VpsAPICreateServerRequest: vps api create server request.
type VpsAPICreateServerRequest struct {
	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// OfferID: offer ID of the server.
	OfferID string `json:"offer_id"`

	// OsID: oS ID to install on the server.
	OsID string `json:"os_id"`

	// LocationID: location ID of the server.
	LocationID string `json:"location_id"`

	// Hostname: hostname of the server.
	Hostname string `json:"hostname"`

	// SSHKeyID: SSH key to install on the server.
	SSHKeyID *string `json:"ssh_key_id,omitempty"`
}

// VpsAPIDeleteServerRequest: vps api delete server request.
type VpsAPIDeleteServerRequest struct {
	// ServerID: ID of the server to delete.
	ServerID string `json:"-"`
}

// VpsAPIGetOfferRequest: vps api get offer request.
type VpsAPIGetOfferRequest struct {
	// OfferID: ID of the offer.
	OfferID string `json:"-"`
}

// VpsAPIGetServerRequest: vps api get server request.
type VpsAPIGetServerRequest struct {
	// ServerID: ID of the server.
	ServerID string `json:"-"`
}

// VpsAPIListLocationsRequest: vps api list locations request.
type VpsAPIListLocationsRequest struct {
	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of locations per page.
	PageSize *uint32 `json:"-"`
}

// VpsAPIListOSRequest: vps api list os request.
type VpsAPIListOSRequest struct {
	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of OS per page.
	PageSize *uint32 `json:"-"`
}

// VpsAPIListOffersRequest: vps api list offers request.
type VpsAPIListOffersRequest struct {
	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of offers per page.
	PageSize *uint32 `json:"-"`
}

// VpsAPIListServersRequest: vps api list servers request.
type VpsAPIListServersRequest struct {
	// OrganizationID: organization ID.
	// Precisely one of OrganizationID, ProjectID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// ProjectID: project ID.
	// Precisely one of OrganizationID, ProjectID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// Page: page number.
	Page *uint32 `json:"-"`

	// PageSize: number of servers per page.
	PageSize *uint32 `json:"-"`
}

// VpsAPIManageServerPowerRequest: vps api manage server power request.
type VpsAPIManageServerPowerRequest struct {
	// ServerID: ID of the server.
	ServerID string `json:"-"`

	// Action: action to perform (power_off, power_on, reboot).
	// Default value: unknown_action
	Action ManageServerPowerRequestAction `json:"action"`
}

// VpsAPIReinstallServerRequest: vps api reinstall server request.
type VpsAPIReinstallServerRequest struct {
	// ServerID: ID of the server.
	ServerID string `json:"-"`

	// OsID: oS ID to reinstall on the server.
	OsID string `json:"os_id"`

	// SSHKeyID: SSH key to install on the server.
	SSHKeyID *string `json:"ssh_key_id,omitempty"`
}

// Virtual private server API.
type VpsAPI struct {
	client *scw.Client
}

// NewVpsAPI returns a VpsAPI object from a Scaleway client.
func NewVpsAPI(client *scw.Client) *VpsAPI {
	return &VpsAPI{
		client: client,
	}
}

// ListServers: List VPS for project.
func (s *VpsAPI) ListServers(req *VpsAPIListServersRequest, opts ...scw.RequestOption) (*ListServersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.OrganizationID == nil && req.ProjectID == nil {
		req.ProjectID = &defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vps/v1/servers",
		Query:  query,
	}

	var resp ListServersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		for _, el := range resp.Servers {
			el.setSRN(apiMetadata.Domain)
		}
	}
	return &resp, nil
}

// GetServer: Get the VPS associated with the given ID.
func (s *VpsAPI) GetServer(req *VpsAPIGetServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vps/v1/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// WaitForServerRequest is used by WaitForServer method.
type WaitForServerRequest struct {
	ServerID      string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForServer waits for the Server to reach a terminal state.
func (s *VpsAPI) WaitForServer(req *WaitForServerRequest, opts ...scw.RequestOption) (*Server, error) {
	timeout := defaultVpsTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultVpsRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[ServerStatus]struct{}{
		ServerStatusOrdered:    {},
		ServerStatusInstalling: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetServer(&VpsAPIGetServerRequest{
				ServerID: req.ServerID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTransient := transientStatuses[res.Status]

			return res, !isTransient, nil
		},
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
		Timeout:          timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for Server failed")
	}

	return res.(*Server), nil
}

// CreateServer: Create a new VPS. The order return you a service ID to follow the provisioning status with GetServer.
func (s *VpsAPI) CreateServer(req *VpsAPICreateServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vps/v1/servers",
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
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// DeleteServer: Delete the VPS associated with the given ID.
func (s *VpsAPI) DeleteServer(req *VpsAPIDeleteServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/vps/v1/servers/" + fmt.Sprint(req.ServerID) + "",
	}

	var resp Server

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// ChangeServerHostname: Change a VPS's hostname.
func (s *VpsAPI) ChangeServerHostname(req *VpsAPIChangeServerHostnameRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vps/v1/servers/" + fmt.Sprint(req.ServerID) + "/change-hostname",
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
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// ChangeServerPassword: Change a VPS's root password.
func (s *VpsAPI) ChangeServerPassword(req *VpsAPIChangeServerPasswordRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vps/v1/servers/" + fmt.Sprint(req.ServerID) + "/change-password",
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
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// ManageServerPower: Switch VM off/on or reboot.
func (s *VpsAPI) ManageServerPower(req *VpsAPIManageServerPowerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vps/v1/servers/" + fmt.Sprint(req.ServerID) + "/manage-power",
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
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// ReinstallServer: Reinstall a given OS to the VPS.
func (s *VpsAPI) ReinstallServer(req *VpsAPIReinstallServerRequest, opts ...scw.RequestOption) (*Server, error) {
	var err error

	if fmt.Sprint(req.ServerID) == "" {
		return nil, errors.New("field ServerID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/vps/v1/servers/" + fmt.Sprint(req.ServerID) + "/reinstall",
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
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// ListOffers: List VPS offers.
func (s *VpsAPI) ListOffers(req *VpsAPIListOffersRequest, opts ...scw.RequestOption) (*ListOffersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vps/v1/offers",
		Query:  query,
	}

	var resp ListOffersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		for _, el := range resp.Offers {
			el.setSRN(apiMetadata.Domain)
		}
	}
	return &resp, nil
}

// GetOffer: Return specific offer for the given ID.
func (s *VpsAPI) GetOffer(req *VpsAPIGetOfferRequest, opts ...scw.RequestOption) (*Offer, error) {
	var err error

	if fmt.Sprint(req.OfferID) == "" {
		return nil, errors.New("field OfferID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vps/v1/offers/" + fmt.Sprint(req.OfferID) + "",
	}

	var resp Offer

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		resp.setSRN(apiMetadata.Domain)
	}
	return &resp, nil
}

// ListOS: List all available OS.
func (s *VpsAPI) ListOS(req *VpsAPIListOSRequest, opts ...scw.RequestOption) (*ListOSResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vps/v1/os",
		Query:  query,
	}

	var resp ListOSResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		for _, el := range resp.Os {
			el.setSRN(apiMetadata.Domain)
		}
	}
	return &resp, nil
}

// ListLocations: List available locations and their available offers.
func (s *VpsAPI) ListLocations(req *VpsAPIListLocationsRequest, opts ...scw.RequestOption) (*ListLocationsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/vps/v1/locations",
		Query:  query,
	}

	var resp ListLocationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	apiMetadata, err := s.client.GetAPIMetadata()
	if err == nil {
		for _, el := range resp.Locations {
			el.setSRN(apiMetadata.Domain)
		}
	}
	return &resp, nil
}

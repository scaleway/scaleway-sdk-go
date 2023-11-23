// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package domain_private provides methods and message types of the domain_private v2beta1 API.
package domain_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/scw"
	domain_v2beta1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/domain/v2beta1"
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

type GeographicArea string

const (
	GeographicAreaUnknownGeographicArea = GeographicArea("unknown_geographic_area")
	GeographicAreaFrench                = GeographicArea("french")
	GeographicAreaEuropean              = GeographicArea("european")
	GeographicAreaAsian                 = GeographicArea("asian")
	GeographicAreaEuropeanEconomicArea  = GeographicArea("european_economic_area")
	GeographicAreaAllCountries          = GeographicArea("all_countries")
)

func (enum GeographicArea) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_geographic_area"
	}
	return string(enum)
}

func (enum GeographicArea) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *GeographicArea) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = GeographicArea(GeographicArea(tmp).String())
	return nil
}

type ListFrModesResponseFrModeFieldFrModeFieldType string

const (
	ListFrModesResponseFrModeFieldFrModeFieldTypeUnknownFrModeFieldType = ListFrModesResponseFrModeFieldFrModeFieldType("unknown_fr_mode_field_type")
	ListFrModesResponseFrModeFieldFrModeFieldTypeString                 = ListFrModesResponseFrModeFieldFrModeFieldType("string")
	ListFrModesResponseFrModeFieldFrModeFieldTypeNumber                 = ListFrModesResponseFrModeFieldFrModeFieldType("number")
	ListFrModesResponseFrModeFieldFrModeFieldTypeBoolean                = ListFrModesResponseFrModeFieldFrModeFieldType("boolean")
	ListFrModesResponseFrModeFieldFrModeFieldTypeDate                   = ListFrModesResponseFrModeFieldFrModeFieldType("date")
	ListFrModesResponseFrModeFieldFrModeFieldTypeRegex                  = ListFrModesResponseFrModeFieldFrModeFieldType("regex")
)

func (enum ListFrModesResponseFrModeFieldFrModeFieldType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_fr_mode_field_type"
	}
	return string(enum)
}

func (enum ListFrModesResponseFrModeFieldFrModeFieldType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFrModesResponseFrModeFieldFrModeFieldType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFrModesResponseFrModeFieldFrModeFieldType(ListFrModesResponseFrModeFieldFrModeFieldType(tmp).String())
	return nil
}

type MessageStatus string

const (
	MessageStatusStatusUnknown = MessageStatus("status_unknown")
	MessageStatusPending       = MessageStatus("pending")
	MessageStatusError         = MessageStatus("error")
	MessageStatusSent          = MessageStatus("sent")
)

func (enum MessageStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "status_unknown"
	}
	return string(enum)
}

func (enum MessageStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MessageStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MessageStatus(MessageStatus(tmp).String())
	return nil
}

type MessageType string

const (
	MessageTypeTypeUnknown                        = MessageType("type_unknown")
	MessageTypeBuyDomain                          = MessageType("buy_domain")
	MessageTypeRegisterExternalDomainToken        = MessageType("register_external_domain_token")
	MessageTypeRegisterExternalDomainConfirmation = MessageType("register_external_domain_confirmation")
	MessageTypeDeleteExternalDomainWarning        = MessageType("delete_external_domain_warning")
	MessageTypeDeleteExternalDomainWarning24h     = MessageType("delete_external_domain_warning_24h")
	MessageTypeDeleteExternalDomain               = MessageType("delete_external_domain")
	MessageTypeDeleteDomainExpired                = MessageType("delete_domain_expired")
	MessageTypeEmailNotValid                      = MessageType("email_not_valid")
	MessageTypeDomainCreated                      = MessageType("domain_created")
	MessageTypeDomainCreateFailed                 = MessageType("domain_create_failed")
	MessageTypeDomainRenewed                      = MessageType("domain_renewed")
	MessageTypeDomainRenewFailed                  = MessageType("domain_renew_failed")
	MessageTypeAutoRenewSuspended                 = MessageType("auto_renew_suspended")
	MessageTypeDomainLateRenewed                  = MessageType("domain_late_renewed")
	MessageTypeDomainLateRenewFailed              = MessageType("domain_late_renew_failed")
	MessageTypeDomainTradeCanceled                = MessageType("domain_trade_canceled")
	MessageTypeDomainTransfered                   = MessageType("domain_transfered")
	MessageTypeDomainTransferFailed               = MessageType("domain_transfer_failed")
)

func (enum MessageType) String() string {
	if enum == "" {
		// return default value if empty
		return "type_unknown"
	}
	return string(enum)
}

func (enum MessageType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MessageType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MessageType(MessageType(tmp).String())
	return nil
}

type RawFormat string

const (
	RawFormatUnknownRawFormat = RawFormat("unknown_raw_format")
	RawFormatNetbox           = RawFormat("netbox")
)

func (enum RawFormat) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_raw_format"
	}
	return string(enum)
}

func (enum RawFormat) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RawFormat) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RawFormat(RawFormat(tmp).String())
	return nil
}

type ReverseIPScope string

const (
	ReverseIPScopeScopeUnknown = ReverseIPScope("scope_unknown")
	ReverseIPScopeOnline       = ReverseIPScope("online")
	ReverseIPScopeScaleway     = ReverseIPScope("scaleway")
)

func (enum ReverseIPScope) String() string {
	if enum == "" {
		// return default value if empty
		return "scope_unknown"
	}
	return string(enum)
}

func (enum ReverseIPScope) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReverseIPScope) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReverseIPScope(ReverseIPScope(tmp).String())
	return nil
}

type ReverseIPStatus string

const (
	ReverseIPStatusStatusUnknown = ReverseIPStatus("status_unknown")
	ReverseIPStatusPending       = ReverseIPStatus("pending")
	ReverseIPStatusActive        = ReverseIPStatus("active")
	ReverseIPStatusError         = ReverseIPStatus("error")
)

func (enum ReverseIPStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "status_unknown"
	}
	return string(enum)
}

func (enum ReverseIPStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReverseIPStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReverseIPStatus(ReverseIPStatus(tmp).String())
	return nil
}

// FormField: form field.
type FormField struct {
	Code string `json:"code"`

	Name string `json:"name"`
}

// ListFrModesResponseFrModeField: list fr modes response fr mode field.
type ListFrModesResponseFrModeField struct {
	FieldName string `json:"field_name"`

	Example string `json:"example"`

	// Type: default value: unknown_fr_mode_field_type
	Type ListFrModesResponseFrModeFieldFrModeFieldType `json:"type"`

	TypeDetail *string `json:"type_detail"`
}

// Nameserver: nameserver.
type Nameserver struct {
	Name string `json:"name"`

	IP []string `json:"ip"`
}

// FormFieldCountry: form field country.
type FormFieldCountry struct {
	Code string `json:"code"`

	Name string `json:"name"`

	Continent *FormField `json:"continent"`
}

// ListFrModesResponseFrMode: list fr modes response fr mode.
type ListFrModesResponseFrMode struct {
	ModeName string `json:"mode_name"`

	ModeAPIName string `json:"mode_api_name"`

	Fields []*ListFrModesResponseFrModeField `json:"fields"`
}

// InstancesSource: instances source.
type InstancesSource struct {
	Az string `json:"az"`

	URL string `json:"url"`
}

// Message: message.
type Message struct {
	Domain string `json:"domain"`

	// Type: default value: type_unknown
	Type MessageType `json:"type"`

	Comment *string `json:"comment"`

	// Status: default value: status_unknown
	Status MessageStatus `json:"status"`

	CreatedAt *time.Time `json:"created_at"`
}

// ReverseIP: reverse ip.
type ReverseIP struct {
	IP net.IP `json:"ip"`

	Hostname string `json:"hostname"`

	// Status: default value: status_unknown
	Status ReverseIPStatus `json:"status"`

	MasterDNSPublic string `json:"master_dns_public"`

	MasterDNSInternal string `json:"master_dns_internal"`

	// Scope: default value: scope_unknown
	Scope ReverseIPScope `json:"scope"`
}

// Slave: slave.
type Slave struct {
	Configuration string `json:"configuration"`

	Name string `json:"name"`

	Driver string `json:"driver"`

	URI string `json:"uri"`
}

// FormFieldState: form field state.
type FormFieldState struct {
	Code string `json:"code"`

	Name string `json:"name"`

	Continent *FormField `json:"continent"`
}

// RDAP: rdap.
type RDAP struct {
	Domain string `json:"domain"`

	// DnssecStatus: default value: feature_status_unknown
	DnssecStatus domain_v2beta1.DomainFeatureStatus `json:"dnssec_status"`

	// DnssecDelegationStatus: default value: feature_status_unknown
	DnssecDelegationStatus domain_v2beta1.DomainFeatureStatus `json:"dnssec_delegation_status"`

	DsRecords []*domain_v2beta1.DSRecord `json:"ds_records"`

	EppCode []string `json:"epp_code"`

	ExpiredAt *time.Time `json:"expired_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	Registrar string `json:"registrar"`

	Nameservers []*Nameserver `json:"nameservers"`
}

// AddInstancesSourceRequest: add instances source request.
type AddInstancesSourceRequest struct {
	Az string `json:"-"`

	URL string `json:"url"`

	Token string `json:"token"`
}

// AddSlaveRequest: add slave request.
type AddSlaveRequest struct {
	Configuration string `json:"-"`

	Name string `json:"name"`

	Driver string `json:"driver"`

	URI string `json:"uri"`

	Login string `json:"login"`

	Password string `json:"password"`
}

// ConsoleAPIGetRdapRequest: console api get rdap request.
type ConsoleAPIGetRdapRequest struct {
	Domain string `json:"-"`
}

// ConsoleAPIListContinentsRequest: console api list continents request.
type ConsoleAPIListContinentsRequest struct {
}

// ConsoleAPIListCountriesRequest: console api list countries request.
type ConsoleAPIListCountriesRequest struct {
	Continent *string `json:"-"`

	// Area: default value: unknown_geographic_area
	Area GeographicArea `json:"-"`
}

// ConsoleAPIListRegionsRequest: console api list regions request.
type ConsoleAPIListRegionsRequest struct {
}

// ConsoleAPIListStatesRequest: console api list states request.
type ConsoleAPIListStatesRequest struct {
	Continent *string `json:"-"`
}

// CreateMessageRequest: create message request.
type CreateMessageRequest struct {
	Domain string `json:"domain"`

	// Type: default value: type_unknown
	Type MessageType `json:"type"`

	Comment *string `json:"comment,omitempty"`
}

// DeleteInstancesSourceRequest: delete instances source request.
type DeleteInstancesSourceRequest struct {
	Az string `json:"-"`
}

// DeleteSlaveRequest: delete slave request.
type DeleteSlaveRequest struct {
	Configuration string `json:"-"`

	Name string `json:"-"`
}

// ImportRawDNSZoneRequest: import raw dns zone request.
type ImportRawDNSZoneRequest struct {
	DNSZone string `json:"-"`

	// Format: default value: unknown_raw_format
	Format RawFormat `json:"format"`

	Content string `json:"content"`

	ProjectID string `json:"project_id"`
}

// ImportRawDNSZoneResponse: import raw dns zone response.
type ImportRawDNSZoneResponse struct {
}

// ListContinentsResponse: list continents response.
type ListContinentsResponse struct {
	Continents []*FormField `json:"continents"`
}

// ListCountriesResponse: list countries response.
type ListCountriesResponse struct {
	Countries []*FormFieldCountry `json:"countries"`
}

// ListFrModesResponse: list fr modes response.
type ListFrModesResponse struct {
	Modes []*ListFrModesResponseFrMode `json:"modes"`
}

// ListInstancesSourcesRequest: list instances sources request.
type ListInstancesSourcesRequest struct {
}

// ListInstancesSourcesResponse: list instances sources response.
type ListInstancesSourcesResponse struct {
	Sources []*InstancesSource `json:"sources"`
}

// ListMessagesRequest: list messages request.
type ListMessagesRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Domain string `json:"-"`

	// Type: default value: type_unknown
	Type MessageType `json:"-"`

	// Status: default value: status_unknown
	Status MessageStatus `json:"-"`
}

// ListMessagesResponse: list messages response.
type ListMessagesResponse struct {
	TotalCount uint32 `json:"total_count"`

	Messages []*Message `json:"messages"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListMessagesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListMessagesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListMessagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Messages = append(r.Messages, results.Messages...)
	r.TotalCount += uint32(len(results.Messages))
	return uint32(len(results.Messages)), nil
}

// ListNlLegalFormsResponse: list nl legal forms response.
type ListNlLegalFormsResponse struct {
	LegalForms []string `json:"legal_forms"`
}

// ListRegionsResponse: list regions response.
type ListRegionsResponse struct {
	Regions []*FormField `json:"regions"`
}

// ListReverseIPRequest: list reverse ip request.
type ListReverseIPRequest struct {
	Search []string `json:"-"`
}

// ListReverseIPResponse: list reverse ip response.
type ListReverseIPResponse struct {
	ReverseIP []*ReverseIP `json:"reverse_ip"`
}

// ListSlavesRequest: list slaves request.
type ListSlavesRequest struct {
	Configuration string `json:"-"`
}

// ListSlavesResponse: list slaves response.
type ListSlavesResponse struct {
	Slaves []*Slave `json:"slaves"`
}

// ListStatesResponse: list states response.
type ListStatesResponse struct {
	States []*FormFieldState `json:"states"`
}

// RdapResponse: rdap response.
type RdapResponse struct {
	DomainRdap *RDAP `json:"domain_rdap"`
}

// ResetReverseIPRequest: reset reverse ip request.
type ResetReverseIPRequest struct {
	ReverseIP net.IP `json:"-"`
}

// ResetReverseIPResponse: reset reverse ip response.
type ResetReverseIPResponse struct {
	ReverseIP *ReverseIP `json:"reverse_ip"`
}

// UpdateReverseIPRequest: update reverse ip request.
type UpdateReverseIPRequest struct {
	ReverseIP net.IP `json:"-"`

	Hostname string `json:"hostname"`
}

// UpdateReverseIPResponse: update reverse ip response.
type UpdateReverseIPResponse struct {
	ReverseIP *ReverseIP `json:"reverse_ip"`
}

// Domain private API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ListReverseIP: List reverse IP with a list of search values.
// You can search by `IPv4`, `IPv6`, `Hostname` or `Network subnet`.
func (s *API) ListReverseIP(req *ListReverseIPRequest, opts ...scw.RequestOption) (*ListReverseIPResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "search", req.Search)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-private/v2beta1/reverse-ips",
		Query:  query,
	}

	var resp ListReverseIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateReverseIP: Replace a reverse IP with the given IPv4 or IPv6 and an Hostname.
func (s *API) UpdateReverseIP(req *UpdateReverseIPRequest, opts ...scw.RequestOption) (*UpdateReverseIPResponse, error) {
	var err error

	if fmt.Sprint(req.ReverseIP) == "" {
		return nil, errors.New("field ReverseIP cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/domain-private/v2beta1/reverse-ips/" + fmt.Sprint(req.ReverseIP) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UpdateReverseIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResetReverseIP: Reset a reverse IP with the given IPv4 or IPv6 and set the default reverse back.
func (s *API) ResetReverseIP(req *ResetReverseIPRequest, opts ...scw.RequestOption) (*ResetReverseIPResponse, error) {
	var err error

	if fmt.Sprint(req.ReverseIP) == "" {
		return nil, errors.New("field ReverseIP cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain-private/v2beta1/reverse-ips/" + fmt.Sprint(req.ReverseIP) + "/reset",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ResetReverseIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateMessage: Create a message for a domain.
func (s *API) CreateMessage(req *CreateMessageRequest, opts ...scw.RequestOption) (*Message, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain-private/v2beta1/messages",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Message

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListMessages: List messages by domain, type and/or status.
func (s *API) ListMessages(req *ListMessagesRequest, opts ...scw.RequestOption) (*ListMessagesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "domain", req.Domain)
	parameter.AddToQuery(query, "type", req.Type)
	parameter.AddToQuery(query, "status", req.Status)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-private/v2beta1/messages",
		Query:  query,
	}

	var resp ListMessagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ImportRawDNSZone: Import records from external source (example netbox).
func (s *API) ImportRawDNSZone(req *ImportRawDNSZoneRequest, opts ...scw.RequestOption) (*ImportRawDNSZoneResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.DNSZone) == "" {
		return nil, errors.New("field DNSZone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain-private/v2beta1/dns-zones/" + fmt.Sprint(req.DNSZone) + "/raw",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ImportRawDNSZoneResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddSlave:
func (s *API) AddSlave(req *AddSlaveRequest, opts ...scw.RequestOption) (*Slave, error) {
	var err error

	if fmt.Sprint(req.Configuration) == "" {
		return nil, errors.New("field Configuration cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain-private/v2beta1/slaves/" + fmt.Sprint(req.Configuration) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Slave

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSlaves:
func (s *API) ListSlaves(req *ListSlavesRequest, opts ...scw.RequestOption) (*ListSlavesResponse, error) {
	var err error

	if fmt.Sprint(req.Configuration) == "" {
		return nil, errors.New("field Configuration cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-private/v2beta1/slaves/" + fmt.Sprint(req.Configuration) + "",
	}

	var resp ListSlavesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSlave:
func (s *API) DeleteSlave(req *DeleteSlaveRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.Configuration) == "" {
		return errors.New("field Configuration cannot be empty in request")
	}

	if fmt.Sprint(req.Name) == "" {
		return errors.New("field Name cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/domain-private/v2beta1/slaves/" + fmt.Sprint(req.Configuration) + "/" + fmt.Sprint(req.Name) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// AddInstancesSource:
func (s *API) AddInstancesSource(req *AddInstancesSourceRequest, opts ...scw.RequestOption) (*InstancesSource, error) {
	var err error

	if fmt.Sprint(req.Az) == "" {
		return nil, errors.New("field Az cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/domain-private/v2beta1/instances-sources/" + fmt.Sprint(req.Az) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp InstancesSource

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListInstancesSources:
func (s *API) ListInstancesSources(req *ListInstancesSourcesRequest, opts ...scw.RequestOption) (*ListInstancesSourcesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-private/v2beta1/instances-sources",
	}

	var resp ListInstancesSourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteInstancesSource:
func (s *API) DeleteInstancesSource(req *DeleteInstancesSourceRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.Az) == "" {
		return errors.New("field Az cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/domain-private/v2beta1/instances-sources/" + fmt.Sprint(req.Az) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Domain private API.
type ConsoleAPI struct {
	client *scw.Client
}

// NewConsoleAPI returns a ConsoleAPI object from a Scaleway client.
func NewConsoleAPI(client *scw.Client) *ConsoleAPI {
	return &ConsoleAPI{
		client: client,
	}
}

// GetRdap: Get Rdap answer for a Domain, if Rdap not available fallback to whois.
func (s *ConsoleAPI) GetRdap(req *ConsoleAPIGetRdapRequest, opts ...scw.RequestOption) (*RdapResponse, error) {
	var err error

	if fmt.Sprint(req.Domain) == "" {
		return nil, errors.New("field Domain cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-private/v2beta1/rdaps/" + fmt.Sprint(req.Domain) + "",
	}

	var resp RdapResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRegions: List regions available.
func (s *ConsoleAPI) ListRegions(req *ConsoleAPIListRegionsRequest, opts ...scw.RequestOption) (*ListRegionsResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-private/v2beta1/regions",
	}

	var resp ListRegionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListContinents: List all continents available for form request.<br/>
// For example: This list can be use to fill the GeoIP continents field.
func (s *ConsoleAPI) ListContinents(req *ConsoleAPIListContinentsRequest, opts ...scw.RequestOption) (*ListContinentsResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-private/v2beta1/continents",
	}

	var resp ListContinentsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCountries: List all countries available for form request.<br/>
// For example: This list can be use to fill the GeoIP countries field.
func (s *ConsoleAPI) ListCountries(req *ConsoleAPIListCountriesRequest, opts ...scw.RequestOption) (*ListCountriesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "continent", req.Continent)
	parameter.AddToQuery(query, "area", req.Area)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-private/v2beta1/countries",
		Query:  query,
	}

	var resp ListCountriesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListStates: List all states available for form request.
func (s *ConsoleAPI) ListStates(req *ConsoleAPIListStatesRequest, opts ...scw.RequestOption) (*ListStatesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "continent", req.Continent)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-private/v2beta1/states",
		Query:  query,
	}

	var resp ListStatesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNlLegalForms: List legal forms for .NL domain.
func (s *ConsoleAPI) ListNlLegalForms(opts ...scw.RequestOption) (*ListNlLegalFormsResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-private/v2beta1/nl-legal-forms",
	}

	var resp ListNlLegalFormsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFrModes: List fields linked to a contact mode for .FR domain.
func (s *ConsoleAPI) ListFrModes(opts ...scw.RequestOption) (*ListFrModesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/domain-private/v2beta1/fr-modes",
	}

	var resp ListFrModesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package account_private provides methods and message types of the account_private v1 API.
package account_private

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

type LegalForm string

const (
	// Unknown legal form.
	LegalFormUnknownLegalForm = LegalForm("unknown_legal_form")
	// Individual.
	LegalFormIndividual = LegalForm("individual")
	// Corporate.
	LegalFormCorporate = LegalForm("corporate")
)

func (enum LegalForm) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_legal_form"
	}
	return string(enum)
}

func (enum LegalForm) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LegalForm) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LegalForm(LegalForm(tmp).String())
	return nil
}

type MigrationStateScalewayOrganizationScalewayOrganizationState string

const (
	// Unknown state.
	MigrationStateScalewayOrganizationScalewayOrganizationStateUnknownState = MigrationStateScalewayOrganizationScalewayOrganizationState("unknown_state")
	// Available.
	MigrationStateScalewayOrganizationScalewayOrganizationStateAvailable = MigrationStateScalewayOrganizationScalewayOrganizationState("available")
	// Unavailable.
	MigrationStateScalewayOrganizationScalewayOrganizationStateUnavailable = MigrationStateScalewayOrganizationScalewayOrganizationState("unavailable")
	// Already linked.
	MigrationStateScalewayOrganizationScalewayOrganizationStateAlreadyLinked = MigrationStateScalewayOrganizationScalewayOrganizationState("already_linked")
)

func (enum MigrationStateScalewayOrganizationScalewayOrganizationState) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_state"
	}
	return string(enum)
}

func (enum MigrationStateScalewayOrganizationScalewayOrganizationState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MigrationStateScalewayOrganizationScalewayOrganizationState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MigrationStateScalewayOrganizationScalewayOrganizationState(MigrationStateScalewayOrganizationScalewayOrganizationState(tmp).String())
	return nil
}

type MigrationStateStatus string

const (
	// Unknown status.
	MigrationStateStatusUnknownStatus = MigrationStateStatus("unknown_status")
	// Missing Dedibox account.
	MigrationStateStatusMissingDediboxAccount = MigrationStateStatus("missing_dedibox_account")
	// Missing Scaleway user.
	MigrationStateStatusMissingScalewayUser = MigrationStateStatus("missing_scaleway_user")
	// Missing Scaleway organization.
	MigrationStateStatusMissingScalewayOrganization = MigrationStateStatus("missing_scaleway_organization")
	// Missing source of data.
	MigrationStateStatusMissingSourceOfData = MigrationStateStatus("missing_source_of_data")
	// Ready to commit.
	MigrationStateStatusReadyToCommit = MigrationStateStatus("ready_to_commit")
	// Error.
	MigrationStateStatusError = MigrationStateStatus("error")
)

func (enum MigrationStateStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum MigrationStateStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MigrationStateStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MigrationStateStatus(MigrationStateStatus(tmp).String())
	return nil
}

type SourceOfData string

const (
	// Unknown source of data.
	SourceOfDataUnknownSourceOfData = SourceOfData("unknown_source_of_data")
	// Scaleway.
	SourceOfDataScaleway = SourceOfData("scaleway")
	// Dedibox.
	SourceOfDataDedibox = SourceOfData("dedibox")
)

func (enum SourceOfData) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_source_of_data"
	}
	return string(enum)
}

func (enum SourceOfData) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SourceOfData) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SourceOfData(SourceOfData(tmp).String())
	return nil
}

// Address: address.
type Address struct {
	Street string `json:"street"`

	Complementary string `json:"complementary"`

	City string `json:"city"`

	PostalCode string `json:"postal_code"`

	Country string `json:"country"`
}

// ListCountryCodesResponseCountry: list country codes response country.
type ListCountryCodesResponseCountry struct {
	Name string `json:"name"`

	Code string `json:"code"`

	Flag string `json:"flag"`
}

// ListSubdivisionCodesResponseSubdivision: list subdivision codes response subdivision.
type ListSubdivisionCodesResponseSubdivision struct {
	Name string `json:"name"`

	Code string `json:"code"`
}

// MigrationStateDediboxAccount: migration state dedibox account.
type MigrationStateDediboxAccount struct {
	ID uint32 `json:"id"`

	Login string `json:"login"`

	Email string `json:"email"`

	Lastname string `json:"lastname"`

	Firstname string `json:"firstname"`

	PhoneNumber string `json:"phone_number"`

	Address *Address `json:"address"`

	Company *string `json:"company"`

	VatNumber *string `json:"vat_number"`

	// LegalForm: default value: unknown_legal_form
	LegalForm LegalForm `json:"legal_form"`
}

// MigrationStateScalewayOrganization: migration state scaleway organization.
type MigrationStateScalewayOrganization struct {
	ID string `json:"id"`

	Name string `json:"name"`

	// State: default value: unknown_state
	State MigrationStateScalewayOrganizationScalewayOrganizationState `json:"state"`

	Address *Address `json:"address"`

	VatNumber *string `json:"vat_number"`

	// LegalForm: default value: unknown_legal_form
	LegalForm LegalForm `json:"legal_form"`
}

// MigrationStateScalewayUser: migration state scaleway user.
type MigrationStateScalewayUser struct {
	ID string `json:"id"`

	Email string `json:"email"`

	Lastname string `json:"lastname"`

	Firstname string `json:"firstname"`

	PhoneNumber string `json:"phone_number"`
}

// CommitMigrationResponse: commit migration response.
type CommitMigrationResponse struct {
	Email string `json:"email"`

	MagicLinkToken string `json:"magic_link_token"`

	OrganizationID string `json:"organization_id"`
}

// ConsoleAPIEnableIAMRequest: console api enable iam request.
type ConsoleAPIEnableIAMRequest struct {
	OrganizationID string `json:"organization_id"`
}

// ConsoleAPIInviteUserRequest: console api invite user request.
type ConsoleAPIInviteUserRequest struct {
	OrganizationID string `json:"organization_id"`

	Email string `json:"email"`
}

// ConsoleAPIReinviteUserRequest: console api reinvite user request.
type ConsoleAPIReinviteUserRequest struct {
	OrganizationID string `json:"organization_id"`

	Email string `json:"email"`
}

// DediboxToElementsAPICommitMigrationRequest: dedibox to elements api commit migration request.
type DediboxToElementsAPICommitMigrationRequest struct {
	Jwt string `json:"jwt"`
}

// DediboxToElementsAPIDeleteScalewayOrganizationRequest: dedibox to elements api delete scaleway organization request.
type DediboxToElementsAPIDeleteScalewayOrganizationRequest struct {
	Jwt string `json:"-"`
}

// DediboxToElementsAPIDeleteScalewayUserRequest: dedibox to elements api delete scaleway user request.
type DediboxToElementsAPIDeleteScalewayUserRequest struct {
	Jwt string `json:"-"`
}

// DediboxToElementsAPIGetMigrationStateRequest: dedibox to elements api get migration state request.
type DediboxToElementsAPIGetMigrationStateRequest struct {
	Jwt string `json:"-"`
}

// DediboxToElementsAPISelectScalewayOrganizationRequest: dedibox to elements api select scaleway organization request.
type DediboxToElementsAPISelectScalewayOrganizationRequest struct {
	Jwt string `json:"jwt"`

	OrganizationID *string `json:"organization_id,omitempty"`
}

// DediboxToElementsAPISelectScalewayUserRequest: dedibox to elements api select scaleway user request.
type DediboxToElementsAPISelectScalewayUserRequest struct {
	Jwt string `json:"jwt"`

	ScalewaySessionToken *string `json:"scaleway_session_token,omitempty"`
}

// DediboxToElementsAPISelectSourceOfDataRequest: dedibox to elements api select source of data request.
type DediboxToElementsAPISelectSourceOfDataRequest struct {
	Jwt string `json:"jwt"`

	// SourceOfData: default value: unknown_source_of_data
	SourceOfData SourceOfData `json:"source_of_data"`
}

// ISOCodesAPIListCountryCodesRequest: iso codes api list country codes request.
type ISOCodesAPIListCountryCodesRequest struct {
}

// ISOCodesAPIListSubdivisionCodesRequest: iso codes api list subdivision codes request.
type ISOCodesAPIListSubdivisionCodesRequest struct {
	CountryCode string `json:"-"`
}

// InviteUserResponse: invite user response.
type InviteUserResponse struct {
	UserID string `json:"user_id"`
}

// ListCountryCodesResponse: list country codes response.
type ListCountryCodesResponse struct {
	Countries []*ListCountryCodesResponseCountry `json:"countries"`
}

// ListSubdivisionCodesResponse: list subdivision codes response.
type ListSubdivisionCodesResponse struct {
	Subdivisions []*ListSubdivisionCodesResponseSubdivision `json:"subdivisions"`
}

// MigrationState: migration state.
type MigrationState struct {
	// Status: default value: unknown_status
	Status MigrationStateStatus `json:"status"`

	DediboxAccount *MigrationStateDediboxAccount `json:"dedibox_account"`

	ScalewayUser *MigrationStateScalewayUser `json:"scaleway_user"`

	SelectedScalewayOrganization *MigrationStateScalewayOrganization `json:"selected_scaleway_organization"`

	ScalewayOrganizations []*MigrationStateScalewayOrganization `json:"scaleway_organizations"`

	// SourceOfData: default value: unknown_source_of_data
	SourceOfData SourceOfData `json:"source_of_data"`

	ErrorMessage *string `json:"error_message"`

	ErrorKey *string `json:"error_key"`
}

// ScalewayOrganizationResponse: scaleway organization response.
type ScalewayOrganizationResponse struct {
	Jwt string `json:"jwt"`
}

// ScalewayUserResponse: scaleway user response.
type ScalewayUserResponse struct {
	Jwt string `json:"jwt"`
}

// SelectSourceOfDataResponse: select source of data response.
type SelectSourceOfDataResponse struct {
	Jwt string `json:"jwt"`
}

// StartMigrationResponse: start migration response.
type StartMigrationResponse struct {
	Jwt string `json:"jwt"`
}

type ConsoleAPI struct {
	client *scw.Client
}

// NewConsoleAPI returns a ConsoleAPI object from a Scaleway client.
func NewConsoleAPI(client *scw.Client) *ConsoleAPI {
	return &ConsoleAPI{
		client: client,
	}
}

// GetServiceInfo:
func (s *ConsoleAPI) GetServiceInfo(opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-private/v1",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// InviteUser:
func (s *ConsoleAPI) InviteUser(req *ConsoleAPIInviteUserRequest, opts ...scw.RequestOption) (*InviteUserResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-private/v1/invite-user",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp InviteUserResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ReinviteUser:
func (s *ConsoleAPI) ReinviteUser(req *ConsoleAPIReinviteUserRequest, opts ...scw.RequestOption) error {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-private/v1/reinvite-user",
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

// EnableIAM:
func (s *ConsoleAPI) EnableIAM(req *ConsoleAPIEnableIAMRequest, opts ...scw.RequestOption) error {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-private/v1/enable-iam",
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

type DediboxToElementsAPI struct {
	client *scw.Client
}

// NewDediboxToElementsAPI returns a DediboxToElementsAPI object from a Scaleway client.
func NewDediboxToElementsAPI(client *scw.Client) *DediboxToElementsAPI {
	return &DediboxToElementsAPI{
		client: client,
	}
}

// GetServiceInfo:
func (s *DediboxToElementsAPI) GetServiceInfo(opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-private/v1/dedibox-to-elements",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartMigration:
func (s *DediboxToElementsAPI) StartMigration(opts ...scw.RequestOption) (*StartMigrationResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-private/v1/dedibox-to-elements/start",
	}

	var resp StartMigrationResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CommitMigration:
func (s *DediboxToElementsAPI) CommitMigration(req *DediboxToElementsAPICommitMigrationRequest, opts ...scw.RequestOption) (*CommitMigrationResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-private/v1/dedibox-to-elements/commit",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CommitMigrationResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetMigrationState:
func (s *DediboxToElementsAPI) GetMigrationState(req *DediboxToElementsAPIGetMigrationStateRequest, opts ...scw.RequestOption) (*MigrationState, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "jwt", req.Jwt)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-private/v1/dedibox-to-elements/state",
		Query:  query,
	}

	var resp MigrationState

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SelectScalewayUser:
func (s *DediboxToElementsAPI) SelectScalewayUser(req *DediboxToElementsAPISelectScalewayUserRequest, opts ...scw.RequestOption) (*ScalewayUserResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-private/v1/dedibox-to-elements/scaleway-user",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ScalewayUserResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteScalewayUser:
func (s *DediboxToElementsAPI) DeleteScalewayUser(req *DediboxToElementsAPIDeleteScalewayUserRequest, opts ...scw.RequestOption) (*ScalewayUserResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "jwt", req.Jwt)

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account-private/v1/dedibox-to-elements/scaleway-user",
		Query:  query,
	}

	var resp ScalewayUserResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SelectScalewayOrganization:
func (s *DediboxToElementsAPI) SelectScalewayOrganization(req *DediboxToElementsAPISelectScalewayOrganizationRequest, opts ...scw.RequestOption) (*ScalewayOrganizationResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-private/v1/dedibox-to-elements/scaleway-organization",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ScalewayOrganizationResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteScalewayOrganization:
func (s *DediboxToElementsAPI) DeleteScalewayOrganization(req *DediboxToElementsAPIDeleteScalewayOrganizationRequest, opts ...scw.RequestOption) (*ScalewayOrganizationResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "jwt", req.Jwt)

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account-private/v1/dedibox-to-elements/scaleway-organization",
		Query:  query,
	}

	var resp ScalewayOrganizationResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SelectSourceOfData:
func (s *DediboxToElementsAPI) SelectSourceOfData(req *DediboxToElementsAPISelectSourceOfDataRequest, opts ...scw.RequestOption) (*SelectSourceOfDataResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-private/v1/dedibox-to-elements/source-of-data",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SelectSourceOfDataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ISOCodesAPI struct {
	client *scw.Client
}

// NewISOCodesAPI returns a ISOCodesAPI object from a Scaleway client.
func NewISOCodesAPI(client *scw.Client) *ISOCodesAPI {
	return &ISOCodesAPI{
		client: client,
	}
}

// GetServiceInfo:
func (s *ISOCodesAPI) GetServiceInfo(opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-private/v1/iso-codes",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCountryCodes:
func (s *ISOCodesAPI) ListCountryCodes(req *ISOCodesAPIListCountryCodesRequest, opts ...scw.RequestOption) (*ListCountryCodesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-private/v1/iso-codes/countries",
	}

	var resp ListCountryCodesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSubdivisionCodes:
func (s *ISOCodesAPI) ListSubdivisionCodes(req *ISOCodesAPIListSubdivisionCodesRequest, opts ...scw.RequestOption) (*ListSubdivisionCodesResponse, error) {
	var err error

	if fmt.Sprint(req.CountryCode) == "" {
		return nil, errors.New("field CountryCode cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-private/v1/iso-codes/subdivisions/" + fmt.Sprint(req.CountryCode) + "",
	}

	var resp ListSubdivisionCodesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

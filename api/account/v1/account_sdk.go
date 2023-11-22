// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package account provides methods and message types of the account v1 API.
package account

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

type ListTokensRequestOrderBy string

const (
	// Creation date ascending.
	ListTokensRequestOrderByCreatedAtAsc = ListTokensRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListTokensRequestOrderByCreatedAtDesc = ListTokensRequestOrderBy("created_at_desc")
)

func (enum ListTokensRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListTokensRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTokensRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTokensRequestOrderBy(ListTokensRequestOrderBy(tmp).String())
	return nil
}

type OrganizationCurrency string

const (
	// Euro.
	OrganizationCurrencyEur = OrganizationCurrency("eur")
)

func (enum OrganizationCurrency) String() string {
	if enum == "" {
		// return default value if empty
		return "eur"
	}
	return string(enum)
}

func (enum OrganizationCurrency) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationCurrency) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationCurrency(OrganizationCurrency(tmp).String())
	return nil
}

type OrganizationCustomerClass string

const (
	// Individual.
	OrganizationCustomerClassIndividual = OrganizationCustomerClass("individual")
	// Corporate.
	OrganizationCustomerClassCorporate = OrganizationCustomerClass("corporate")
	OrganizationCustomerClassInternal  = OrganizationCustomerClass("internal")
)

func (enum OrganizationCustomerClass) String() string {
	if enum == "" {
		// return default value if empty
		return "individual"
	}
	return string(enum)
}

func (enum OrganizationCustomerClass) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationCustomerClass) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationCustomerClass(OrganizationCustomerClass(tmp).String())
	return nil
}

type SupportSubscriptionLevel string

const (
	// Basic.
	SupportSubscriptionLevelBasic = SupportSubscriptionLevel("basic")
	// Developer.
	SupportSubscriptionLevelDeveloper = SupportSubscriptionLevel("developer")
	// Business.
	SupportSubscriptionLevelBusiness = SupportSubscriptionLevel("business")
	// Enterprise.
	SupportSubscriptionLevelEnterprise = SupportSubscriptionLevel("enterprise")
)

func (enum SupportSubscriptionLevel) String() string {
	if enum == "" {
		// return default value if empty
		return "basic"
	}
	return string(enum)
}

func (enum SupportSubscriptionLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SupportSubscriptionLevel) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SupportSubscriptionLevel(SupportSubscriptionLevel(tmp).String())
	return nil
}

type TokenCategory string

const (
	// Session.
	TokenCategorySession = TokenCategory("session")
	// User created.
	TokenCategoryUserCreated = TokenCategory("user_created")
)

func (enum TokenCategory) String() string {
	if enum == "" {
		// return default value if empty
		return "session"
	}
	return string(enum)
}

func (enum TokenCategory) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TokenCategory) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TokenCategory(TokenCategory(tmp).String())
	return nil
}

// OrganizationAddress: organization address.
type OrganizationAddress struct {
	Line1 string `json:"line1"`

	Line2 string `json:"line2"`

	CountryCode string `json:"country_code"`

	CityName string `json:"city_name"`

	PostalCode string `json:"postal_code"`

	SubdivisionCode string `json:"subdivision_code"`
}

// TokenPermissionsResources: token permissions resources.
type TokenPermissionsResources struct {
	Details []string `json:"details"`
}

// Abuse: abuse.
type Abuse struct {
	ID string `json:"id"`
}

// Organization: organization.
type Organization struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Locale string `json:"locale"`

	// Currency: default value: eur
	Currency OrganizationCurrency `json:"currency"`

	// CustomerClass: default value: individual
	CustomerClass OrganizationCustomerClass `json:"customer_class"`

	Address *OrganizationAddress `json:"address"`

	VatNumber *string `json:"vat_number"`

	Timezone string `json:"timezone"`

	// SupportLevel: default value: basic
	SupportLevel SupportSubscriptionLevel `json:"support_level"`

	SupportID string `json:"support_id"`

	SupportPin string `json:"support_pin"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// SupportSubscription: support subscription.
type SupportSubscription struct {
	ID string `json:"id"`

	// Level: default value: basic
	Level SupportSubscriptionLevel `json:"level"`

	StartTime *time.Time `json:"start_time"`

	EndTime *time.Time `json:"end_time"`
}

// Token: token.
type Token struct {
	ID string `json:"id"`

	// SecretKey: use this to authenticate.
	SecretKey string `json:"secret_key"`

	UserID string `json:"user_id"`

	// Category: default value: session
	Category TokenCategory `json:"category"`

	CreatedAt *time.Time `json:"created_at"`

	Description string `json:"description"`

	AccessKey string `json:"access_key"`

	InheritsUserPerms bool `json:"inherits_user_perms"`

	CreationIP string `json:"creation_ip"`

	ExpiresAt *time.Time `json:"expires_at"`

	DeletedAt *time.Time `json:"deleted_at"`
}

// User: user.
type User struct {
	ID string `json:"id"`

	FirstName string `json:"first_name"`

	LastName string `json:"last_name"`

	Email string `json:"email"`

	PhoneNumber *string `json:"phone_number"`

	TwoFactorEnabled bool `json:"two_factor_enabled"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// TokenPermissionsService: token permissions service.
type TokenPermissionsService struct {
	Resources map[string]*TokenPermissionsResources `json:"resources"`
}

// BackupCodes: backup codes.
type BackupCodes struct {
	BackupCodes []string `json:"backup_codes"`
}

// CloseAbuseRequest: close abuse request.
type CloseAbuseRequest struct {
	ID string `json:"-"`
}

// CompleteResetPasswordRequest: complete reset password request.
type CompleteResetPasswordRequest struct {
	Token string `json:"token"`

	NewPassword string `json:"new_password"`
}

// CreateSupportSubscriptionRequest: create support subscription request.
type CreateSupportSubscriptionRequest struct {
	OrganizationID string `json:"organization_id"`

	// Level: default value: basic
	Level SupportSubscriptionLevel `json:"level"`
}

// CreateTokenRequest: create token request.
type CreateTokenRequest struct {
	Email string `json:"email"`

	Expires bool `json:"expires"`

	Description string `json:"description"`
}

// CreateUserRequest: create user request.
type CreateUserRequest struct {
	Email string `json:"email"`

	Password string `json:"password"`
}

// DeleteTokenRequest: delete token request.
type DeleteTokenRequest struct {
	AccessKey string `json:"-"`
}

// DisableTwoFactorAuthRequest: disable two factor auth request.
type DisableTwoFactorAuthRequest struct {
	UserID string `json:"-"`
}

// EnableTwoFactorAuthRequest: enable two factor auth request.
type EnableTwoFactorAuthRequest struct {
	UserID string `json:"-"`
}

// EnableTwoFactorAuthResponse: enable two factor auth response.
type EnableTwoFactorAuthResponse struct {
	Secret string `json:"secret"`
}

// GetOrganizationRequest: get organization request.
type GetOrganizationRequest struct {
	ID string `json:"-"`
}

// GetSupportSubscriptionRequest: get support subscription request.
type GetSupportSubscriptionRequest struct {
	SupportSubscriptionID string `json:"-"`

	OrganizationID string `json:"-"`
}

// GetTokenPermissionsRequest: get token permissions request.
type GetTokenPermissionsRequest struct {
	AccessKey string `json:"-"`

	OrganizationID *string `json:"-"`

	ResourceID *string `json:"-"`

	Service *string `json:"-"`

	Name *string `json:"-"`

	Permission *string `json:"-"`
}

// GetTokenRequest: get token request.
type GetTokenRequest struct {
	AccessKey string `json:"-"`
}

// GetUserRequest: get user request.
type GetUserRequest struct {
	ID string `json:"-"`
}

// ListAbusesRequest: list abuses request.
type ListAbusesRequest struct {
	OrganizationID string `json:"-"`
}

// ListAbusesResponse: list abuses response.
type ListAbusesResponse struct {
	Abuses []*Abuse `json:"abuses"`
}

// ListOrganizationsRequest: list organizations request.
type ListOrganizationsRequest struct {
	UserID string `json:"-"`
}

// ListOrganizationsResponse: list organizations response.
type ListOrganizationsResponse struct {
	Organizations []*Organization `json:"organizations"`
}

// ListQuotasRequest: list quotas request.
type ListQuotasRequest struct {
	OrganizationID string `json:"-"`
}

// ListQuotasResponse: list quotas response.
type ListQuotasResponse struct {
	Quotas map[string]*int32 `json:"quotas"`
}

// ListSupportSubscriptionsRequest: list support subscriptions request.
type ListSupportSubscriptionsRequest struct {
	OrganizationID string `json:"-"`
}

// ListSupportSubscriptionsResponse: list support subscriptions response.
type ListSupportSubscriptionsResponse struct {
	SupportSubscriptions []*SupportSubscription `json:"support_subscriptions"`
}

// ListTokensRequest: Request to list your tokens.
type ListTokensRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListTokensRequestOrderBy `json:"-"`
}

// ListTokensResponse: list tokens response.
type ListTokensResponse struct {
	TotalCount uint32 `json:"total_count"`

	Tokens []*Token `json:"tokens"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTokensResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTokensResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListTokensResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tokens = append(r.Tokens, results.Tokens...)
	r.TotalCount += uint32(len(results.Tokens))
	return uint32(len(results.Tokens)), nil
}

// ListUsersRequest: list users request.
type ListUsersRequest struct {
	OrganizationID string `json:"-"`
}

// ListUsersResponse: list users response.
type ListUsersResponse struct {
	Users []*User `json:"users"`
}

// LoginRequest: login request.
type LoginRequest struct {
	Email string `json:"email"`

	Password string `json:"password"`

	TwoFactorToken string `json:"two_factor_token"`
}

// StartResetPasswordRequest: start reset password request.
type StartResetPasswordRequest struct {
	Email string `json:"email"`
}

// TokenPermissions: token permissions.
type TokenPermissions struct {
	Services map[string]*TokenPermissionsService `json:"services"`
}

// UpdateOrganizationRequest: update organization request.
type UpdateOrganizationRequest struct {
	ID string `json:"-"`

	Name string `json:"name"`

	Locale string `json:"locale"`

	// Currency: default value: eur
	Currency OrganizationCurrency `json:"currency"`

	Timezone string `json:"timezone"`

	Address *OrganizationAddress `json:"address,omitempty"`

	VatNumber *string `json:"vat_number,omitempty"`
}

// UpdateTokenRequest: update token request.
type UpdateTokenRequest struct {
	AccessKey string `json:"-"`

	Description string `json:"description"`
}

// UpdateUserRequest: update user request.
type UpdateUserRequest struct {
	ID string `json:"-"`

	FirstName *string `json:"first_name,omitempty"`

	LastName *string `json:"last_name,omitempty"`

	Email *string `json:"email,omitempty"`

	PhoneNumber *string `json:"phone_number,omitempty"`
}

// ValidateTwoFactorAuthRequest: validate two factor auth request.
type ValidateTwoFactorAuthRequest struct {
	UserID string `json:"-"`

	TwoFactorToken string `json:"two_factor_token"`
}

// VerifyUserEmailRequest: verify user email request.
type VerifyUserEmailRequest struct {
	UserID string `json:"-"`
}

// VerifyUserPhoneRequest: verify user phone request.
type VerifyUserPhoneRequest struct {
	UserID string `json:"-"`

	PhoneNumber string `json:"phone_number"`
}

// This API allows you to access all user related data.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ListTokens: List all your tokens.
func (s *API) ListTokens(req *ListTokensRequest, opts ...scw.RequestOption) (*ListTokensResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v1/tokens",
		Query:  query,
	}

	var resp ListTokensResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetToken: Get the token associated with the given access key.
func (s *API) GetToken(req *GetTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if fmt.Sprint(req.AccessKey) == "" {
		return nil, errors.New("field AccessKey cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v1/tokens/" + fmt.Sprint(req.AccessKey) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateToken: Create a new token.
func (s *API) CreateToken(req *CreateTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v1/tokens",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateToken: Update the token associated with the given access key.
func (s *API) UpdateToken(req *UpdateTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if fmt.Sprint(req.AccessKey) == "" {
		return nil, errors.New("field AccessKey cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account/v1/tokens/" + fmt.Sprint(req.AccessKey) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteToken: Delete the token associated with the given access key.
func (s *API) DeleteToken(req *DeleteTokenRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.AccessKey) == "" {
		return errors.New("field AccessKey cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account/v1/tokens/" + fmt.Sprint(req.AccessKey) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetTokenPermissions: Get permissions associated with the given access key.
func (s *API) GetTokenPermissions(req *GetTokenPermissionsRequest, opts ...scw.RequestOption) (*TokenPermissions, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "resource_id", req.ResourceID)
	parameter.AddToQuery(query, "service", req.Service)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "permission", req.Permission)

	if fmt.Sprint(req.AccessKey) == "" {
		return nil, errors.New("field AccessKey cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v1/tokens/" + fmt.Sprint(req.AccessKey) + "/permissions",
		Query:  query,
	}

	var resp TokenPermissions

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListUsers:
func (s *API) ListUsers(req *ListUsersRequest, opts ...scw.RequestOption) (*ListUsersResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v1/users",
		Query:  query,
	}

	var resp ListUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUser:
func (s *API) GetUser(req *GetUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v1/users/" + fmt.Sprint(req.ID) + "",
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateUser:
func (s *API) CreateUser(req *CreateUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v1/users",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateUser:
func (s *API) UpdateUser(req *UpdateUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account/v1/users/" + fmt.Sprint(req.ID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// VerifyUserEmail:
func (s *API) VerifyUserEmail(req *VerifyUserEmailRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v1/users/" + fmt.Sprint(req.UserID) + "/verify-email",
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

// VerifyUserPhone:
func (s *API) VerifyUserPhone(req *VerifyUserPhoneRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v1/users/" + fmt.Sprint(req.UserID) + "/verify-phone",
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

// EnableTwoFactorAuth: Enable 2FA.
func (s *API) EnableTwoFactorAuth(req *EnableTwoFactorAuthRequest, opts ...scw.RequestOption) (*EnableTwoFactorAuthResponse, error) {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v1/users/" + fmt.Sprint(req.UserID) + "/enable-two-factor-auth",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EnableTwoFactorAuthResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateTwoFactorAuth: Validate 2FA.
func (s *API) ValidateTwoFactorAuth(req *ValidateTwoFactorAuthRequest, opts ...scw.RequestOption) (*BackupCodes, error) {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v1/users/" + fmt.Sprint(req.UserID) + "/validate-two-factor-auth",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp BackupCodes

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableTwoFactorAuth: Disable 2FA.
func (s *API) DisableTwoFactorAuth(req *DisableTwoFactorAuthRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v1/users/" + fmt.Sprint(req.UserID) + "/disable-two-factor-auth",
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

// ListOrganizations:
func (s *API) ListOrganizations(req *ListOrganizationsRequest, opts ...scw.RequestOption) (*ListOrganizationsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "user_id", req.UserID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v1/organizations",
		Query:  query,
	}

	var resp ListOrganizationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetOrganization:
func (s *API) GetOrganization(req *GetOrganizationRequest, opts ...scw.RequestOption) (*Organization, error) {
	var err error

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v1/organizations/" + fmt.Sprint(req.ID) + "",
	}

	var resp Organization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOrganization:
func (s *API) UpdateOrganization(req *UpdateOrganizationRequest, opts ...scw.RequestOption) (*Organization, error) {
	var err error

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account/v1/organizations/" + fmt.Sprint(req.ID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Organization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListAbuses:
func (s *API) ListAbuses(req *ListAbusesRequest, opts ...scw.RequestOption) (*ListAbusesResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v1/abuses",
		Query:  query,
	}

	var resp ListAbusesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CloseAbuse:
func (s *API) CloseAbuse(req *CloseAbuseRequest, opts ...scw.RequestOption) (*interface{}, error) {
	var err error

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account/v1/abuses/" + fmt.Sprint(req.ID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp interface{}

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSupportSubscriptions:
func (s *API) ListSupportSubscriptions(req *ListSupportSubscriptionsRequest, opts ...scw.RequestOption) (*ListSupportSubscriptionsResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v1/support-subscriptions",
		Query:  query,
	}

	var resp ListSupportSubscriptionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSupportSubscription:
func (s *API) GetSupportSubscription(req *GetSupportSubscriptionRequest, opts ...scw.RequestOption) (*SupportSubscription, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.SupportSubscriptionID) == "" {
		return nil, errors.New("field SupportSubscriptionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v1/support-subscriptions/" + fmt.Sprint(req.SupportSubscriptionID) + "",
		Query:  query,
	}

	var resp SupportSubscription

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSupportSubscription:
func (s *API) CreateSupportSubscription(req *CreateSupportSubscriptionRequest, opts ...scw.RequestOption) (*SupportSubscription, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v1/support-subscriptions",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SupportSubscription

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListQuotas:
func (s *API) ListQuotas(req *ListQuotasRequest, opts ...scw.RequestOption) (*ListQuotasResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v1/quotas",
		Query:  query,
	}

	var resp ListQuotasResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Login: Login user.
func (s *API) Login(req *LoginRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v1/login",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartResetPassword: Start reset-password.
func (s *API) StartResetPassword(req *StartResetPasswordRequest, opts ...scw.RequestOption) error {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v1/start-reset-password",
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

// CompleteResetPassword: Complete reset-password.
func (s *API) CompleteResetPassword(req *CompleteResetPasswordRequest, opts ...scw.RequestOption) error {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account/v1/complete-reset-password",
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

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package account_private provides methods and message types of the account_private v1beta1 API.
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

type CheckPermissionsResponseResponseDecision string

const (
	// Deny.
	CheckPermissionsResponseResponseDecisionDeny = CheckPermissionsResponseResponseDecision("deny")
	// Allow.
	CheckPermissionsResponseResponseDecisionAllow = CheckPermissionsResponseResponseDecision("allow")
)

func (enum CheckPermissionsResponseResponseDecision) String() string {
	if enum == "" {
		// return default value if empty
		return "deny"
	}
	return string(enum)
}

func (enum CheckPermissionsResponseResponseDecision) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CheckPermissionsResponseResponseDecision) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CheckPermissionsResponseResponseDecision(CheckPermissionsResponseResponseDecision(tmp).String())
	return nil
}

type IdentityVerificationStatus string

const (
	// Unknown.
	IdentityVerificationStatusUnknown = IdentityVerificationStatus("unknown")
	// Pending.
	IdentityVerificationStatusPending = IdentityVerificationStatus("pending")
	// Validated.
	IdentityVerificationStatusValidated = IdentityVerificationStatus("validated")
	// Waiting for approval.
	IdentityVerificationStatusWaitingForApproval = IdentityVerificationStatus("waiting_for_approval")
	// Rejected.
	IdentityVerificationStatusRejected = IdentityVerificationStatus("rejected")
	// Failed.
	IdentityVerificationStatusFailed = IdentityVerificationStatus("failed")
)

func (enum IdentityVerificationStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum IdentityVerificationStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IdentityVerificationStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IdentityVerificationStatus(IdentityVerificationStatus(tmp).String())
	return nil
}

type ListIdentityVerificationsRequestOrderBy string

const (
	// Creation date ascending.
	ListIdentityVerificationsRequestOrderByCreatedAtAsc = ListIdentityVerificationsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListIdentityVerificationsRequestOrderByCreatedAtDesc = ListIdentityVerificationsRequestOrderBy("created_at_desc")
)

func (enum ListIdentityVerificationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListIdentityVerificationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListIdentityVerificationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListIdentityVerificationsRequestOrderBy(ListIdentityVerificationsRequestOrderBy(tmp).String())
	return nil
}

type OrganizationUserRole string

const (
	// Unknown.
	OrganizationUserRoleUnknown = OrganizationUserRole("unknown")
	// Owner.
	OrganizationUserRoleOwner = OrganizationUserRole("owner")
	// Admin.
	OrganizationUserRoleAdmin = OrganizationUserRole("admin")
	// Ops.
	OrganizationUserRoleOps = OrganizationUserRole("ops")
	// Billing.
	OrganizationUserRoleBilling = OrganizationUserRole("billing")
	// Delegated billing.
	OrganizationUserRoleDelegatedBilling = OrganizationUserRole("delegated_billing")
	// Delegated owner.
	OrganizationUserRoleDelegatedOwner = OrganizationUserRole("delegated_owner")
)

func (enum OrganizationUserRole) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum OrganizationUserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationUserRole) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationUserRole(OrganizationUserRole(tmp).String())
	return nil
}

type OrganizationUserState string

const (
	// Invited.
	OrganizationUserStateInvited = OrganizationUserState("invited")
	// Accepted.
	OrganizationUserStateAccepted = OrganizationUserState("accepted")
)

func (enum OrganizationUserState) String() string {
	if enum == "" {
		// return default value if empty
		return "invited"
	}
	return string(enum)
}

func (enum OrganizationUserState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationUserState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationUserState(OrganizationUserState(tmp).String())
	return nil
}

// IdentityVerificationJumioVerification: identity verification jumio verification.
type IdentityVerificationJumioVerification struct {
	IframeURL string `json:"iframe_url"`
}

// AddUsersToOrganizationRequestUser: add users to organization request user.
type AddUsersToOrganizationRequestUser struct {
	// Email: email of the user that will be added to the organization.
	Email string `json:"email"`

	// Role: role the user will have in the organization.
	// Default value: unknown
	Role OrganizationUserRole `json:"role"`
}

// OrganizationUser: organization user.
type OrganizationUser struct {
	// Email: display the user email.
	Email string `json:"email"`

	// TwoFa: display the state of the 2FA user.
	TwoFa bool `json:"two_fa"`

	// Role: display the user role.
	// Default value: unknown
	Role OrganizationUserRole `json:"role"`

	// LastLogin: display last connection.
	LastLogin *time.Time `json:"last_login"`

	// State: display invitation state.
	// Default value: invited
	State OrganizationUserState `json:"state"`
}

// CheckPermissionsRequestPermission: check permissions request permission.
type CheckPermissionsRequestPermission struct {
	Service string `json:"service"`

	Name string `json:"name"`

	Action string `json:"action"`

	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`
}

// CheckPermissionsResponseResponse: check permissions response response.
type CheckPermissionsResponseResponse struct {
	// Decision: default value: deny
	Decision CheckPermissionsResponseResponseDecision `json:"decision"`
}

// IdentityVerification: identity verification.
type IdentityVerification struct {
	ID string `json:"id"`

	// Status: default value: unknown
	Status IdentityVerificationStatus `json:"status"`

	OrganizationID string `json:"organization_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	// Precisely one of Jumio must be set.
	Jumio *IdentityVerificationJumioVerification `json:"jumio,omitempty"`
}

// AddUsersToOrganizationResponse: add users to organization response.
type AddUsersToOrganizationResponse struct {
	OrganizationUsers []*OrganizationUser `json:"organization_users"`
}

// CheckPermissionsResponse: check permissions response.
type CheckPermissionsResponse struct {
	Responses []*CheckPermissionsResponseResponse `json:"responses"`
}

// ConsoleAPIAddUsersToOrganizationRequest: console api add users to organization request.
type ConsoleAPIAddUsersToOrganizationRequest struct {
	// OrganizationID: organization receiving new users.
	OrganizationID string `json:"-"`

	// Users: emails to be added to the organization along with the desired role for each of them.
	Users []*AddUsersToOrganizationRequestUser `json:"users"`
}

// ConsoleAPICheckPermissionsRequest: console api check permissions request.
type ConsoleAPICheckPermissionsRequest struct {
	Permissions []*CheckPermissionsRequestPermission `json:"permissions"`
}

// ConsoleAPIListOrganizationUsersRequest: console api list organization users request.
type ConsoleAPIListOrganizationUsersRequest struct {
	// OrganizationID: filter user in organization.
	OrganizationID string `json:"-"`
}

// ConsoleAPIRemoveUserFromOrganizationRequest: console api remove user from organization request.
type ConsoleAPIRemoveUserFromOrganizationRequest struct {
	// OrganizationID: organization that will be left.
	OrganizationID string `json:"-"`

	// Email: email of the user that will be removed.
	Email string `json:"-"`
}

// ConsoleAPISendTOSRequest: console api send tos request.
type ConsoleAPISendTOSRequest struct {
	OrganizationID string `json:"-"`
}

// ConsoleAPIUpdateOrganizationUserRequest: console api update organization user request.
type ConsoleAPIUpdateOrganizationUserRequest struct {
	// OrganizationID: organization of the user.
	OrganizationID string `json:"-"`

	// Email: email of the user that will be updated.
	Email string `json:"-"`

	// Role: new role.
	// Default value: unknown
	Role OrganizationUserRole `json:"role"`
}

// KYCApiListIdentityVerificationsRequest: kyc api list identity verifications request.
type KYCApiListIdentityVerificationsRequest struct {
	// OrderBy: default value: created_at_asc
	OrderBy ListIdentityVerificationsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	// Status: default value: unknown
	Status IdentityVerificationStatus `json:"-"`
}

// KYCApiStartIdentityVerificationRequest: kyc api start identity verification request.
type KYCApiStartIdentityVerificationRequest struct {
	OrganizationID string `json:"organization_id"`
}

// ListIdentityVerificationsResponse: list identity verifications response.
type ListIdentityVerificationsResponse struct {
	IDentityVerifications []*IdentityVerification `json:"identity_verifications"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIdentityVerificationsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIdentityVerificationsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListIdentityVerificationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.IDentityVerifications = append(r.IDentityVerifications, results.IDentityVerifications...)
	r.TotalCount += uint32(len(results.IDentityVerifications))
	return uint32(len(results.IDentityVerifications)), nil
}

// ListOrganizationUsersResponse: list organization users response.
type ListOrganizationUsersResponse struct {
	OrganizationUsers []*OrganizationUser `json:"organization_users"`
}

// RemoveUserFromOrganizationsResponse: remove user from organizations response.
type RemoveUserFromOrganizationsResponse struct {
	OrganizationUsers []*OrganizationUser `json:"organization_users"`
}

// Account Console API.
type ConsoleAPI struct {
	client *scw.Client
}

// NewConsoleAPI returns a ConsoleAPI object from a Scaleway client.
func NewConsoleAPI(client *scw.Client) *ConsoleAPI {
	return &ConsoleAPI{
		client: client,
	}
}

// ListOrganizationUsers: List all users in an organization.
func (s *ConsoleAPI) ListOrganizationUsers(req *ConsoleAPIListOrganizationUsersRequest, opts ...scw.RequestOption) (*ListOrganizationUsersResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-private/v1beta1/organizations/" + fmt.Sprint(req.OrganizationID) + "/users",
	}

	var resp ListOrganizationUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddUsersToOrganization: Add user in an organization.
func (s *ConsoleAPI) AddUsersToOrganization(req *ConsoleAPIAddUsersToOrganizationRequest, opts ...scw.RequestOption) (*AddUsersToOrganizationResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-private/v1beta1/organizations/" + fmt.Sprint(req.OrganizationID) + "/users",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AddUsersToOrganizationResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RemoveUserFromOrganization: Delete a user from an organization.
func (s *ConsoleAPI) RemoveUserFromOrganization(req *ConsoleAPIRemoveUserFromOrganizationRequest, opts ...scw.RequestOption) (*RemoveUserFromOrganizationsResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	if fmt.Sprint(req.Email) == "" {
		return nil, errors.New("field Email cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account-private/v1beta1/organizations/" + fmt.Sprint(req.OrganizationID) + "/users/" + fmt.Sprint(req.Email) + "",
	}

	var resp RemoveUserFromOrganizationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOrganizationUser: Update the user's role in an organization.
func (s *ConsoleAPI) UpdateOrganizationUser(req *ConsoleAPIUpdateOrganizationUserRequest, opts ...scw.RequestOption) (*OrganizationUser, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	if fmt.Sprint(req.Email) == "" {
		return nil, errors.New("field Email cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account-private/v1beta1/organizations/" + fmt.Sprint(req.OrganizationID) + "/users/" + fmt.Sprint(req.Email) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp OrganizationUser

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendTOS:
func (s *ConsoleAPI) SendTOS(req *ConsoleAPISendTOSRequest, opts ...scw.RequestOption) error {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-private/v1beta1/organizations/" + fmt.Sprint(req.OrganizationID) + "/send-tos",
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

// CheckPermissions: Check permissions for feature selection.
func (s *ConsoleAPI) CheckPermissions(req *ConsoleAPICheckPermissionsRequest, opts ...scw.RequestOption) (*CheckPermissionsResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-private/v1beta1/check-permissions",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CheckPermissionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Account Private KYC API.
type KYCAPI struct {
	client *scw.Client
}

// NewKYCAPI returns a KYCAPI object from a Scaleway client.
func NewKYCAPI(client *scw.Client) *KYCAPI {
	return &KYCAPI{
		client: client,
	}
}

// StartIdentityVerification: Start an identity verification.
func (s *KYCAPI) StartIdentityVerification(req *KYCApiStartIdentityVerificationRequest, opts ...scw.RequestOption) (*IdentityVerification, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-private/v1beta1/kyc/identity-verifications",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IdentityVerification

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListIdentityVerifications: List identity verifications.
func (s *KYCAPI) ListIdentityVerifications(req *KYCApiListIdentityVerificationsRequest, opts ...scw.RequestOption) (*ListIdentityVerificationsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "status", req.Status)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-private/v1beta1/kyc/identity-verifications",
		Query:  query,
	}

	var resp ListIdentityVerificationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

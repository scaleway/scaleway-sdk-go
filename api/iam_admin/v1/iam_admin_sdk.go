// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package iam_admin provides methods and message types of the iam_admin v1 API.
package iam_admin

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

type APIKeyCategory string

const (
	// Unknown category.
	APIKeyCategoryUnknownCategory = APIKeyCategory("unknown_category")
	// User.
	APIKeyCategoryUser = APIKeyCategory("user")
	// Admin.
	APIKeyCategoryAdmin = APIKeyCategory("admin")
	// Reseller.
	APIKeyCategoryReseller = APIKeyCategory("reseller")
)

func (enum APIKeyCategory) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_category"
	}
	return string(enum)
}

func (enum APIKeyCategory) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *APIKeyCategory) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = APIKeyCategory(APIKeyCategory(tmp).String())
	return nil
}

type CreateProductAPIKeyRequestProductName string

const (
	// Unknown product name.
	CreateProductAPIKeyRequestProductNameUnknownProductName = CreateProductAPIKeyRequestProductName("unknown_product_name")
	// Kapsule.
	CreateProductAPIKeyRequestProductNameKapsule = CreateProductAPIKeyRequestProductName("kapsule")
	// Instances.
	CreateProductAPIKeyRequestProductNameInstances = CreateProductAPIKeyRequestProductName("instances")
	// Video Transcoder.
	CreateProductAPIKeyRequestProductNameVideoTranscoder = CreateProductAPIKeyRequestProductName("video_transcoder")
	// Serverless Functions.
	CreateProductAPIKeyRequestProductNameServerlessFunctions = CreateProductAPIKeyRequestProductName("serverless_functions")
	// Serverless Containers.
	CreateProductAPIKeyRequestProductNameServerlessContainers = CreateProductAPIKeyRequestProductName("serverless_containers")
	// VPC.
	CreateProductAPIKeyRequestProductNameVpc        = CreateProductAPIKeyRequestProductName("vpc")
	CreateProductAPIKeyRequestProductNameEdge       = CreateProductAPIKeyRequestProductName("edge")
	CreateProductAPIKeyRequestProductNameSbs        = CreateProductAPIKeyRequestProductName("sbs")
	CreateProductAPIKeyRequestProductNameWebhosting = CreateProductAPIKeyRequestProductName("webhosting")
)

func (enum CreateProductAPIKeyRequestProductName) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_product_name"
	}
	return string(enum)
}

func (enum CreateProductAPIKeyRequestProductName) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CreateProductAPIKeyRequestProductName) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CreateProductAPIKeyRequestProductName(CreateProductAPIKeyRequestProductName(tmp).String())
	return nil
}

type ListAPIKeysRequestOrderBy string

const (
	// Creation date ascending.
	ListAPIKeysRequestOrderByCreatedAtAsc = ListAPIKeysRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListAPIKeysRequestOrderByCreatedAtDesc = ListAPIKeysRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListAPIKeysRequestOrderByUpdatedAtAsc = ListAPIKeysRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListAPIKeysRequestOrderByUpdatedAtDesc = ListAPIKeysRequestOrderBy("updated_at_desc")
	// Expiration date ascending.
	ListAPIKeysRequestOrderByExpiresAtAsc = ListAPIKeysRequestOrderBy("expires_at_asc")
	// Expiration date descending.
	ListAPIKeysRequestOrderByExpiresAtDesc = ListAPIKeysRequestOrderBy("expires_at_desc")
	// Access key ascending.
	ListAPIKeysRequestOrderByAccessKeyAsc = ListAPIKeysRequestOrderBy("access_key_asc")
	// Access key descending.
	ListAPIKeysRequestOrderByAccessKeyDesc = ListAPIKeysRequestOrderBy("access_key_desc")
)

func (enum ListAPIKeysRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListAPIKeysRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListAPIKeysRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListAPIKeysRequestOrderBy(ListAPIKeysRequestOrderBy(tmp).String())
	return nil
}

type ListApplicationsRequestOrderBy string

const (
	// Creation date ascending.
	ListApplicationsRequestOrderByCreatedAtAsc = ListApplicationsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListApplicationsRequestOrderByCreatedAtDesc = ListApplicationsRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListApplicationsRequestOrderByUpdatedAtAsc = ListApplicationsRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListApplicationsRequestOrderByUpdatedAtDesc = ListApplicationsRequestOrderBy("updated_at_desc")
	// Name ascending.
	ListApplicationsRequestOrderByNameAsc = ListApplicationsRequestOrderBy("name_asc")
	// Name descending.
	ListApplicationsRequestOrderByNameDesc = ListApplicationsRequestOrderBy("name_desc")
)

func (enum ListApplicationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListApplicationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListApplicationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListApplicationsRequestOrderBy(ListApplicationsRequestOrderBy(tmp).String())
	return nil
}

type ListCustomerLevelsRequestOrderBy string

const (
	// Name ascending.
	ListCustomerLevelsRequestOrderByNameAsc = ListCustomerLevelsRequestOrderBy("name_asc")
	// Name descending.
	ListCustomerLevelsRequestOrderByNameDesc = ListCustomerLevelsRequestOrderBy("name_desc")
)

func (enum ListCustomerLevelsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListCustomerLevelsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListCustomerLevelsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListCustomerLevelsRequestOrderBy(ListCustomerLevelsRequestOrderBy(tmp).String())
	return nil
}

type ListGroupsRequestOrderBy string

const (
	// Creation date ascending.
	ListGroupsRequestOrderByCreatedAtAsc = ListGroupsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListGroupsRequestOrderByCreatedAtDesc = ListGroupsRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListGroupsRequestOrderByUpdatedAtAsc = ListGroupsRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListGroupsRequestOrderByUpdatedAtDesc = ListGroupsRequestOrderBy("updated_at_desc")
	// Name ascending.
	ListGroupsRequestOrderByNameAsc = ListGroupsRequestOrderBy("name_asc")
	// Name descending.
	ListGroupsRequestOrderByNameDesc = ListGroupsRequestOrderBy("name_desc")
)

func (enum ListGroupsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListGroupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListGroupsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListGroupsRequestOrderBy(ListGroupsRequestOrderBy(tmp).String())
	return nil
}

type ListInternalOrganizationsRequestOrderBy string

const (
	ListInternalOrganizationsRequestOrderByCreatedAtAsc  = ListInternalOrganizationsRequestOrderBy("created_at_asc")
	ListInternalOrganizationsRequestOrderByCreatedAtDesc = ListInternalOrganizationsRequestOrderBy("created_at_desc")
)

func (enum ListInternalOrganizationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListInternalOrganizationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListInternalOrganizationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListInternalOrganizationsRequestOrderBy(ListInternalOrganizationsRequestOrderBy(tmp).String())
	return nil
}

type ListJWTsRequestOrderBy string

const (
	// Creation date ascending.
	ListJWTsRequestOrderByCreatedAtAsc = ListJWTsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListJWTsRequestOrderByCreatedAtDesc = ListJWTsRequestOrderBy("created_at_desc")
)

func (enum ListJWTsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListJWTsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListJWTsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListJWTsRequestOrderBy(ListJWTsRequestOrderBy(tmp).String())
	return nil
}

type ListLimitsRequestOrderBy string

const (
	// Quota name ascending.
	ListLimitsRequestOrderByQuotumNameAsc = ListLimitsRequestOrderBy("quotum_name_asc")
	// Quota name descending.
	ListLimitsRequestOrderByQuotumNameDesc = ListLimitsRequestOrderBy("quotum_name_desc")
	// Customer level name ascending.
	ListLimitsRequestOrderByCustomerLevelNameAsc = ListLimitsRequestOrderBy("customer_level_name_asc")
	// Customer level name descending.
	ListLimitsRequestOrderByCustomerLevelNameDesc = ListLimitsRequestOrderBy("customer_level_name_desc")
)

func (enum ListLimitsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "quotum_name_asc"
	}
	return string(enum)
}

func (enum ListLimitsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListLimitsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListLimitsRequestOrderBy(ListLimitsRequestOrderBy(tmp).String())
	return nil
}

type ListLogsRequestOrderBy string

const (
	// Creation date ascending.
	ListLogsRequestOrderByCreatedAtAsc = ListLogsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListLogsRequestOrderByCreatedAtDesc = ListLogsRequestOrderBy("created_at_desc")
)

func (enum ListLogsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
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

type ListOrganizationsRequestOrderBy string

const (
	// Creation date ascending.
	ListOrganizationsRequestOrderByCreatedAtAsc = ListOrganizationsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListOrganizationsRequestOrderByCreatedAtDesc = ListOrganizationsRequestOrderBy("created_at_desc")
)

func (enum ListOrganizationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListOrganizationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOrganizationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOrganizationsRequestOrderBy(ListOrganizationsRequestOrderBy(tmp).String())
	return nil
}

type ListPermissionSetsRequestOrderBy string

const (
	// Name ascending.
	ListPermissionSetsRequestOrderByNameAsc = ListPermissionSetsRequestOrderBy("name_asc")
	// Name descending.
	ListPermissionSetsRequestOrderByNameDesc = ListPermissionSetsRequestOrderBy("name_desc")
	// Creation date ascending.
	ListPermissionSetsRequestOrderByCreatedAtAsc = ListPermissionSetsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListPermissionSetsRequestOrderByCreatedAtDesc = ListPermissionSetsRequestOrderBy("created_at_desc")
)

func (enum ListPermissionSetsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListPermissionSetsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPermissionSetsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPermissionSetsRequestOrderBy(ListPermissionSetsRequestOrderBy(tmp).String())
	return nil
}

type ListPoliciesRequestOrderBy string

const (
	// Policy name ascending.
	ListPoliciesRequestOrderByPolicyNameAsc = ListPoliciesRequestOrderBy("policy_name_asc")
	// Policy name descending.
	ListPoliciesRequestOrderByPolicyNameDesc = ListPoliciesRequestOrderBy("policy_name_desc")
	// Creation date ascending.
	ListPoliciesRequestOrderByCreatedAtAsc = ListPoliciesRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListPoliciesRequestOrderByCreatedAtDesc = ListPoliciesRequestOrderBy("created_at_desc")
)

func (enum ListPoliciesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "policy_name_asc"
	}
	return string(enum)
}

func (enum ListPoliciesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPoliciesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPoliciesRequestOrderBy(ListPoliciesRequestOrderBy(tmp).String())
	return nil
}

type ListQuotaRequestOrderBy string

const (
	// Name ascending.
	ListQuotaRequestOrderByNameAsc = ListQuotaRequestOrderBy("name_asc")
	// Name descending.
	ListQuotaRequestOrderByNameDesc = ListQuotaRequestOrderBy("name_desc")
)

func (enum ListQuotaRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListQuotaRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListQuotaRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListQuotaRequestOrderBy(ListQuotaRequestOrderBy(tmp).String())
	return nil
}

type ListSSHKeysRequestOrderBy string

const (
	// Creation date ascending.
	ListSSHKeysRequestOrderByCreatedAtAsc = ListSSHKeysRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListSSHKeysRequestOrderByCreatedAtDesc = ListSSHKeysRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListSSHKeysRequestOrderByUpdatedAtAsc = ListSSHKeysRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListSSHKeysRequestOrderByUpdatedAtDesc = ListSSHKeysRequestOrderBy("updated_at_desc")
	// Name ascending.
	ListSSHKeysRequestOrderByNameAsc = ListSSHKeysRequestOrderBy("name_asc")
	// Name descending.
	ListSSHKeysRequestOrderByNameDesc = ListSSHKeysRequestOrderBy("name_desc")
)

func (enum ListSSHKeysRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSSHKeysRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSSHKeysRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSSHKeysRequestOrderBy(ListSSHKeysRequestOrderBy(tmp).String())
	return nil
}

type ListUsersRequestOrderBy string

const (
	// Creation date ascending.
	ListUsersRequestOrderByCreatedAtAsc = ListUsersRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListUsersRequestOrderByCreatedAtDesc = ListUsersRequestOrderBy("created_at_desc")
	// Update date ascending.
	ListUsersRequestOrderByUpdatedAtAsc = ListUsersRequestOrderBy("updated_at_asc")
	// Update date descending.
	ListUsersRequestOrderByUpdatedAtDesc = ListUsersRequestOrderBy("updated_at_desc")
	// Email ascending.
	ListUsersRequestOrderByEmailAsc = ListUsersRequestOrderBy("email_asc")
	// Email descending.
	ListUsersRequestOrderByEmailDesc = ListUsersRequestOrderBy("email_desc")
	// Last login ascending.
	ListUsersRequestOrderByLastLoginAsc = ListUsersRequestOrderBy("last_login_asc")
	// Last login descending.
	ListUsersRequestOrderByLastLoginDesc = ListUsersRequestOrderBy("last_login_desc")
)

func (enum ListUsersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListUsersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListUsersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListUsersRequestOrderBy(ListUsersRequestOrderBy(tmp).String())
	return nil
}

type LogAction string

const (
	// Unknown action.
	LogActionUnknownAction = LogAction("unknown_action")
	// Created.
	LogActionCreated = LogAction("created")
	// Updated.
	LogActionUpdated = LogAction("updated")
	// Deleted.
	LogActionDeleted = LogAction("deleted")
)

func (enum LogAction) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_action"
	}
	return string(enum)
}

func (enum LogAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LogAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LogAction(LogAction(tmp).String())
	return nil
}

type LogResourceType string

const (
	// Unknown resource type.
	LogResourceTypeUnknownResourceType = LogResourceType("unknown_resource_type")
	// JWT.
	LogResourceTypeJwt = LogResourceType("jwt")
	// API key.
	LogResourceTypeAPIKey = LogResourceType("api_key")
	// User.
	LogResourceTypeUser = LogResourceType("user")
	// Application.
	LogResourceTypeApplication = LogResourceType("application")
	// Group.
	LogResourceTypeGroup = LogResourceType("group")
	// Policy.
	LogResourceTypePolicy = LogResourceType("policy")
	// Organization customer levels.
	LogResourceTypeOrganizationCustomerLevels = LogResourceType("organization__customer_levels")
)

func (enum LogResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_resource_type"
	}
	return string(enum)
}

func (enum LogResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LogResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LogResourceType(LogResourceType(tmp).String())
	return nil
}

type PermissionSetScopeType string

const (
	// Unknown scope type.
	PermissionSetScopeTypeUnknownScopeType = PermissionSetScopeType("unknown_scope_type")
	// Projects.
	PermissionSetScopeTypeProjects = PermissionSetScopeType("projects")
	// Organization.
	PermissionSetScopeTypeOrganization = PermissionSetScopeType("organization")
	// Account root user.
	PermissionSetScopeTypeAccountRootUser = PermissionSetScopeType("account_root_user")
	// Admin.
	PermissionSetScopeTypeAdmin = PermissionSetScopeType("admin")
)

func (enum PermissionSetScopeType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_scope_type"
	}
	return string(enum)
}

func (enum PermissionSetScopeType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PermissionSetScopeType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PermissionSetScopeType(PermissionSetScopeType(tmp).String())
	return nil
}

type UserStatus string

const (
	// Unknown status.
	UserStatusUnknownStatus = UserStatus("unknown_status")
	// Invitation pending.
	UserStatusInvitationPending = UserStatus("invitation_pending")
	// Activated.
	UserStatusActivated = UserStatus("activated")
)

func (enum UserStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum UserStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UserStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UserStatus(UserStatus(tmp).String())
	return nil
}

// InternalOrganizationInternalUser: internal organization internal user.
type InternalOrganizationInternalUser struct {
	ID string `json:"id"`

	AccountRootUserID string `json:"account_root_user_id"`

	Email string `json:"email"`
}

// RuleSpecs: rule specs.
type RuleSpecs struct {
	// PermissionSetNames: names of permission sets bound to the rule.
	PermissionSetNames *[]string `json:"permission_set_names"`

	// ProjectIDs: list of Project IDs the rule is scoped to.
	// Precisely one of ProjectIDs, OrganizationID must be set.
	ProjectIDs *[]string `json:"project_ids,omitempty"`

	// OrganizationID: ID of Organization the rule is scoped to.
	// Precisely one of ProjectIDs, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`
}

// JWT: jwt.
type JWT struct {
	// Jti: jWT ID.
	Jti string `json:"jti"`

	// IssuerID: ID of the user who issued the JWT.
	IssuerID string `json:"issuer_id"`

	// AudienceID: ID of the user targeted by the JWT.
	AudienceID string `json:"audience_id"`

	// CreatedAt: creation date of the JWT.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date of the JWT.
	UpdatedAt *time.Time `json:"updated_at"`

	// ExpiresAt: expiration date of the JWT.
	ExpiresAt *time.Time `json:"expires_at"`

	// IP: IP address used during the creation of the JWT.
	IP net.IP `json:"ip"`

	// UserAgent: user-agent used during the creation of the JWT.
	UserAgent string `json:"user_agent"`

	Staff bool `json:"staff"`

	Renewable bool `json:"renewable"`
}

// PermissionSet: permission set.
type PermissionSet struct {
	// ID: id of the permission set.
	ID string `json:"id"`

	// Name: name of the permission set.
	Name string `json:"name"`

	// ScopeType: scope of the permission set.
	// Default value: unknown_scope_type
	ScopeType PermissionSetScopeType `json:"scope_type"`

	// Description: description of the permission set.
	Description string `json:"description"`

	// Categories: categories of the permission set.
	Categories *[]string `json:"categories"`

	// CreatedAt: when the permission set was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: when the permission set was updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// NbPermissions: number or permissions inside the permission set.
	NbPermissions uint32 `json:"nb_permissions"`
}

// Permission: permission.
type Permission struct {
	// Precisely one of AccountRootUserID, AllOrganizations, OrganizationID, ProjectIDs must be set.
	AccountRootUserID *string `json:"account_root_user_id,omitempty"`

	// Precisely one of AccountRootUserID, AllOrganizations, OrganizationID, ProjectIDs must be set.
	AllOrganizations *bool `json:"all_organizations,omitempty"`

	Name string `json:"name"`

	// Precisely one of AccountRootUserID, AllOrganizations, OrganizationID, ProjectIDs must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// Precisely one of AccountRootUserID, AllOrganizations, OrganizationID, ProjectIDs must be set.
	ProjectIDs *[]string `json:"project_ids,omitempty"`
}

// APIKey: api key.
type APIKey struct {
	// AccessKey: access key of the API key.
	AccessKey string `json:"access_key"`

	// SecretKey: secret key of the API Key.
	SecretKey *string `json:"secret_key"`

	// ApplicationID: ID of application that bears the API key.
	// Precisely one of ApplicationID, UserID must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// UserID: ID of user that bears the API key.
	// Precisely one of ApplicationID, UserID must be set.
	UserID *string `json:"user_id,omitempty"`

	// Description: description of API key.
	Description string `json:"description"`

	// CreatedAt: date and time of API key creation.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date and time of last API key update.
	UpdatedAt *time.Time `json:"updated_at"`

	// ExpiredAt: date and time of API key expiration.
	ExpiredAt *time.Time `json:"expired_at"`

	// DefaultProjectID: default Project ID specified for this API key.
	DefaultProjectID string `json:"default_project_id"`

	// Editable: defines whether or not the API key is editable.
	Editable bool `json:"editable"`

	// IsStaff: if an API key comes from staff.
	IsStaff bool `json:"is_staff"`

	// Category: default value: unknown_category
	Category APIKeyCategory `json:"category"`

	Visible bool `json:"visible"`
}

// Application: application.
type Application struct {
	// ID: ID of the application.
	ID string `json:"id"`

	// Name: name of the application.
	Name string `json:"name"`

	// Description: description of the application.
	Description string `json:"description"`

	// CreatedAt: date and time application was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date and time of last application update.
	UpdatedAt *time.Time `json:"updated_at"`

	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"organization_id"`

	// Editable: defines whether or not the application is editable.
	Editable bool `json:"editable"`

	// NbAPIKeys: number of API keys attributed to the application.
	NbAPIKeys uint32 `json:"nb_api_keys"`

	// Visible: whether an application is visible or not.
	Visible bool `json:"visible"`
}

// CustomerLevel: customer level.
type CustomerLevel struct {
	CreatedAt *time.Time `json:"created_at"`

	ID string `json:"id"`

	Name string `json:"name"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// Group: group.
type Group struct {
	// ID: ID of the group.
	ID string `json:"id"`

	// CreatedAt: date and time of group creation.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date and time of last group update.
	UpdatedAt *time.Time `json:"updated_at"`

	// OrganizationID: ID of Organization linked to the group.
	OrganizationID string `json:"organization_id"`

	// Name: name of the group.
	Name string `json:"name"`

	// Description: description of the group.
	Description string `json:"description"`

	// UserIDs: iDs of users attached to this group.
	UserIDs []string `json:"user_ids"`

	// ApplicationIDs: iDs of applications attached to this group.
	ApplicationIDs []string `json:"application_ids"`

	// Editable: defines whether or not the group is editable.
	Editable bool `json:"editable"`

	// Visible: defines whether or not the group is visible.
	Visible bool `json:"visible"`

	// Managed: defines whether or not the group is managed.
	Managed bool `json:"managed"`
}

// InternalOrganization: internal organization.
type InternalOrganization struct {
	CreatedAt *time.Time `json:"created_at"`

	Creator *InternalOrganizationInternalUser `json:"creator"`

	ID string `json:"id"`

	PermissionSetIDs []string `json:"permission_set_ids"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// Limit: limit.
type Limit struct {
	CustomerLevelID string `json:"customer_level_id"`

	CustomerLevelName string `json:"customer_level_name"`

	// Precisely one of Disabled, Limit, Unlimited must be set.
	Disabled *bool `json:"disabled,omitempty"`

	ID string `json:"id"`

	// Precisely one of Disabled, Limit, Unlimited must be set.
	Limit *uint64 `json:"limit,omitempty"`

	QuotumID string `json:"quotum_id"`

	QuotumName string `json:"quotum_name"`

	// Precisely one of Disabled, Limit, Unlimited must be set.
	Unlimited *bool `json:"unlimited,omitempty"`
}

// Log: log.
type Log struct {
	// ID: log ID.
	ID string `json:"id"`

	// CreatedAt: creation date of the log.
	CreatedAt *time.Time `json:"created_at"`

	// IP: IP address of the HTTP request linked to the log.
	IP net.IP `json:"ip"`

	// UserAgent: user-Agent of the HTTP request linked to the log.
	UserAgent string `json:"user_agent"`

	// Action: action linked to the log.
	// Default value: unknown_action
	Action LogAction `json:"action"`

	// FromPrincipalID: ID of the principal at the origin of the log.
	FromPrincipalID string `json:"from_principal_id"`

	// OrganizationID: ID of Organization linked to the log.
	OrganizationID string `json:"organization_id"`

	// ResourceType: type of the resource linked to the log.
	// Default value: unknown_resource_type
	ResourceType LogResourceType `json:"resource_type"`

	// ResourceID: ID of the resource linked  to the log.
	ResourceID string `json:"resource_id"`

	Data *scw.JSONObject `json:"data"`

	Visible bool `json:"visible"`

	Country string `json:"country"`
}

// Organization: organization.
type Organization struct {
	CreatedAt *time.Time `json:"created_at"`

	CustomerClass string `json:"customer_class"`

	ID string `json:"id"`

	Locked bool `json:"locked"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// Policy: policy.
type Policy struct {
	// ID: id of the policy.
	ID string `json:"id"`

	// Name: name of the policy.
	Name string `json:"name"`

	// Description: description of the policy.
	Description string `json:"description"`

	// OrganizationID: organization ID of the policy.
	OrganizationID string `json:"organization_id"`

	// CreatedAt: date and time of policy creation.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date and time of last policy update.
	UpdatedAt *time.Time `json:"updated_at"`

	// Editable: defines whether or not a policy is editable.
	Editable bool `json:"editable"`

	Visible bool `json:"visible"`

	// NbRules: number of rules of the policy.
	NbRules uint32 `json:"nb_rules"`

	// NbScopes: number of policy scopes.
	NbScopes uint32 `json:"nb_scopes"`

	// NbPermissionSets: number of permission sets of the policy.
	NbPermissionSets uint32 `json:"nb_permission_sets"`

	// UserID: ID of the user attributed to the policy.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	UserID *string `json:"user_id,omitempty"`

	// GroupID: ID of the group attributed to the policy.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	GroupID *string `json:"group_id,omitempty"`

	// ApplicationID: ID of the application attributed to the policy.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// NoPrincipal: defines whether or not a policy is attributed to a principal.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	NoPrincipal *bool `json:"no_principal,omitempty"`
}

// Quotum: quotum.
type Quotum struct {
	ID string `json:"id"`

	// Name: name of the quota.
	Name string `json:"name"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// Rule: rule.
type Rule struct {
	// ID: id of rule.
	ID string `json:"id"`

	// PermissionSetNames: names of permission sets bound to the rule.
	PermissionSetNames *[]string `json:"permission_set_names"`

	// ProjectIDs: list of Project IDs the rule is scoped to.
	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID, AllOrganizations must be set.
	ProjectIDs *[]string `json:"project_ids,omitempty"`

	// OrganizationID: ID of Organization the rule is scoped to.
	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID, AllOrganizations must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// AccountRootUserID: ID of account root user the rule is scoped to.
	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID, AllOrganizations must be set.
	AccountRootUserID *string `json:"account_root_user_id,omitempty"`

	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID, AllOrganizations must be set.
	AllOrganizations *bool `json:"all_organizations,omitempty"`
}

// SSHKey: ssh key.
type SSHKey struct {
	// ID: ID of SSH key.
	ID string `json:"id"`

	// Name: name of SSH key.
	Name string `json:"name"`

	// PublicKey: public key of SSH key.
	PublicKey string `json:"public_key"`

	// Fingerprint: fingerprint of the SSH key.
	Fingerprint string `json:"fingerprint"`

	// CreatedAt: creation date of SSH key.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date of SSH key.
	UpdatedAt *time.Time `json:"updated_at"`

	// OrganizationID: ID of Organization linked to the SSH key.
	OrganizationID string `json:"organization_id"`

	// ProjectID: ID of Project linked to the SSH key.
	ProjectID string `json:"project_id"`

	// Disabled: SSH key status.
	Disabled bool `json:"disabled"`
}

// User: user.
type User struct {
	// ID: ID of user.
	ID string `json:"id"`

	// Email: email of user.
	Email string `json:"email"`

	// CreatedAt: date user was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date of last user update.
	UpdatedAt *time.Time `json:"updated_at"`

	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"organization_id"`

	// Deletable: deletion status of user. Owners cannot be deleted.
	Deletable bool `json:"deletable"`

	// Status: status of user invitation.
	// Default value: unknown_status
	Status UserStatus `json:"status"`

	// AccountRootUserID: ID of the account root user associated with the user.
	AccountRootUserID *string `json:"account_root_user_id"`
}

// AddCustomerLevelPermissionSetRequest: add customer level permission set request.
type AddCustomerLevelPermissionSetRequest struct {
	CustomerLevelID string `json:"-"`

	PermissionSetID string `json:"permission_set_id"`
}

// AddInternalOrganizationPermissionSetRequest: add internal organization permission set request.
type AddInternalOrganizationPermissionSetRequest struct {
	OrganizationID string `json:"-"`

	PermissionSetID string `json:"permission_set_id"`
}

// AddOrganizationCustomerLevelRequest: add organization customer level request.
type AddOrganizationCustomerLevelRequest struct {
	CustomerLevelID string `json:"customer_level_id"`

	OrganizationID string `json:"organization_id"`
}

// CreateAPIKeyRequest: create api key request.
type CreateAPIKeyRequest struct {
	// ApplicationID: ID of the application.
	// Precisely one of ApplicationID, UserID must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// UserID: ID of the user.
	// Precisely one of ApplicationID, UserID must be set.
	UserID *string `json:"user_id,omitempty"`

	// ExpiresAt: expiration date of the API key.
	ExpiresAt *time.Time `json:"expires_at,omitempty"`

	// Description: description of the API key (max length is 200 characters).
	Description string `json:"description"`

	Editable bool `json:"editable"`

	Visible bool `json:"visible"`

	Staff bool `json:"staff"`
}

// CreateCustomerLevelRequest: create customer level request.
type CreateCustomerLevelRequest struct {
	Name string `json:"name"`
}

// CreateGroupRequest: create group request.
type CreateGroupRequest struct {
	// OrganizationID: ID of Organization linked to the group.
	OrganizationID string `json:"organization_id"`

	// Name: name of the group to create (max length is 64 chars). MUST be unique inside an Organization.
	Name string `json:"name"`

	// Description: description of the group to create (max length is 200 chars).
	Description string `json:"description"`
}

// CreateLimitRequest: create limit request.
type CreateLimitRequest struct {
	CustomerLevelID string `json:"customer_level_id"`

	// Precisely one of Disabled, Limit, Unlimited must be set.
	Disabled *bool `json:"disabled,omitempty"`

	// Precisely one of Disabled, Limit, Unlimited must be set.
	Limit *uint64 `json:"limit,omitempty"`

	QuotumID string `json:"quotum_id"`

	// Precisely one of Disabled, Limit, Unlimited must be set.
	Unlimited *bool `json:"unlimited,omitempty"`
}

// CreatePolicyRequest: create policy request.
type CreatePolicyRequest struct {
	// Name: name of the policy to create (max length is 64 characters).
	Name string `json:"name"`

	// Description: description of the policy to create (max length is 200 characters).
	Description string `json:"description"`

	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"organization_id"`

	// Rules: rules of the policy to create.
	Rules []*RuleSpecs `json:"rules"`

	// UserID: ID of user attributed to the policy.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	UserID *string `json:"user_id,omitempty"`

	// GroupID: ID of group attributed to the policy.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	GroupID *string `json:"group_id,omitempty"`

	// ApplicationID: ID of application attributed to the policy.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// NoPrincipal: defines whether or not a policy is attributed to a principal.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	NoPrincipal *bool `json:"no_principal,omitempty"`
}

// CreateProductAPIKeyRequest: create product api key request.
type CreateProductAPIKeyRequest struct {
	Description string `json:"description"`

	ExpiresAt *time.Time `json:"expires_at,omitempty"`

	PermissionSetNames []string `json:"permission_set_names"`

	// ProductName: default value: unknown_product_name
	ProductName CreateProductAPIKeyRequestProductName `json:"product_name"`

	ProjectID string `json:"project_id"`

	ResourceID string `json:"resource_id"`
}

// CreateQuotumRequest: create quotum request.
type CreateQuotumRequest struct {
	Name string `json:"name"`
}

// CreateSSHKeyRequest: create ssh key request.
type CreateSSHKeyRequest struct {
	// Name: name of the SSH key. Max length is 1000.
	Name string `json:"name"`

	// PublicKey: SSH public key. Currently only the ssh-rsa, ssh-dss (DSA), ssh-ed25519 and ecdsa keys with NIST curves are supported. Max length is 65000.
	PublicKey string `json:"public_key"`

	// ProjectID: project the resource is attributed to.
	ProjectID string `json:"project_id"`
}

// DeleteAPIKeyRequest: delete api key request.
type DeleteAPIKeyRequest struct {
	// AccessKey: access key to delete.
	AccessKey string `json:"-"`
}

// DeleteApplicationRequest: delete application request.
type DeleteApplicationRequest struct {
	// ApplicationID: ID of the application to delete.
	ApplicationID string `json:"-"`
}

// DeleteCustomerLevelPermissionSetRequest: delete customer level permission set request.
type DeleteCustomerLevelPermissionSetRequest struct {
	CustomerLevelID string `json:"-"`

	PermissionSetID string `json:"-"`
}

// DeleteCustomerLevelRequest: delete customer level request.
type DeleteCustomerLevelRequest struct {
	CustomerLevelID string `json:"-"`
}

// DeleteGroupRequest: delete group request.
type DeleteGroupRequest struct {
	// GroupID: ID of the group to delete.
	GroupID string `json:"-"`
}

// DeleteInternalOrganizationPermissionSetRequest: delete internal organization permission set request.
type DeleteInternalOrganizationPermissionSetRequest struct {
	OrganizationID string `json:"-"`

	PermissionSetID string `json:"-"`
}

// DeleteJWTRequest: delete jwt request.
type DeleteJWTRequest struct {
	// Jti: jWT ID of the JWT to delete.
	Jti string `json:"-"`
}

// DeleteLimitRequest: delete limit request.
type DeleteLimitRequest struct {
	LimitID string `json:"-"`
}

// DeleteOrganizationCustomerLevelRequest: delete organization customer level request.
type DeleteOrganizationCustomerLevelRequest struct {
	CustomerLevelID string `json:"-"`

	OrganizationID string `json:"-"`
}

// DeletePolicyRequest: delete policy request.
type DeletePolicyRequest struct {
	// PolicyID: id of policy to delete.
	PolicyID string `json:"-"`
}

// DeleteProductAPIKeyRequest: delete product api key request.
type DeleteProductAPIKeyRequest struct {
	AccessKey string `json:"-"`
}

// DeleteQuotumRequest: delete quotum request.
type DeleteQuotumRequest struct {
	QuotumID string `json:"-"`
}

// DeleteSSHKeyRequest: delete ssh key request.
type DeleteSSHKeyRequest struct {
	SSHKeyID string `json:"-"`
}

// EncodedJWT: encoded jwt.
type EncodedJWT struct {
	Jwt *JWT `json:"jwt"`

	Token string `json:"token"`

	RenewToken string `json:"renew_token"`
}

// GetAPIKeyRequest: get api key request.
type GetAPIKeyRequest struct {
	// AccessKey: access key to search for.
	AccessKey string `json:"-"`
}

// GetApplicationRequest: get application request.
type GetApplicationRequest struct {
	// ApplicationID: ID of the application to find.
	ApplicationID string `json:"-"`
}

// GetGroupRequest: get group request.
type GetGroupRequest struct {
	// GroupID: ID of the group.
	GroupID string `json:"-"`
}

// GetInternalOrganizationRequest: get internal organization request.
type GetInternalOrganizationRequest struct {
	OrganizationID string `json:"-"`
}

// GetJWTRequest: get jwt request.
type GetJWTRequest struct {
	Jti string `json:"-"`
}

// GetLogRequest: get log request.
type GetLogRequest struct {
	// LogID: ID of the log.
	LogID string `json:"-"`
}

// GetPermissionSetRequest: get permission set request.
type GetPermissionSetRequest struct {
	PermissionSetName string `json:"-"`
}

// GetPermissionSetResponse: get permission set response.
type GetPermissionSetResponse struct {
	PermissionSet *PermissionSet `json:"permission_set"`

	Permissions []string `json:"permissions"`
}

// GetPermissionsRequest: get permissions request.
type GetPermissionsRequest struct {
	// Precisely one of AuthID, PrincipalID must be set.
	AuthID *string `json:"auth_id,omitempty"`

	// Precisely one of AuthID, PrincipalID must be set.
	PrincipalID *string `json:"principal_id,omitempty"`
}

// GetPermissionsResponse: get permissions response.
type GetPermissionsResponse struct {
	Permissions []*Permission `json:"permissions"`
}

// GetPolicyRequest: get policy request.
type GetPolicyRequest struct {
	// PolicyID: id of policy to search.
	PolicyID string `json:"-"`
}

// GetSSHKeyRequest: get ssh key request.
type GetSSHKeyRequest struct {
	// SSHKeyID: ID of the SSH key.
	SSHKeyID string `json:"-"`
}

// GetUserRequest: get user request.
type GetUserRequest struct {
	// UserID: ID of the user to find.
	UserID string `json:"-"`
}

// ListAPIKeysRequest: list api keys request.
type ListAPIKeysRequest struct {
	// OrderBy: criteria for sorting results.
	// Default value: created_at_asc
	OrderBy ListAPIKeysRequestOrderBy `json:"-"`

	// Page: page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`

	// PageSize: number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// OrganizationID: ID of Organization.
	OrganizationID *string `json:"-"`

	// ApplicationID: ID of application that bears the API key.
	// Precisely one of ApplicationID, UserID must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// UserID: ID of user that bears the API key.
	// Precisely one of ApplicationID, UserID must be set.
	UserID *string `json:"user_id,omitempty"`

	// Editable: defines whether to filter out editable API keys or not.
	Editable *bool `json:"-"`

	Visible *bool `json:"-"`

	// Category: default value: unknown_category
	Category APIKeyCategory `json:"-"`
}

// ListAPIKeysResponse: list api keys response.
type ListAPIKeysResponse struct {
	// APIKeys: list of API keys.
	APIKeys []*APIKey `json:"api_keys"`

	// TotalCount: total count of API Keys.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListAPIKeysResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListAPIKeysResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListAPIKeysResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.APIKeys = append(r.APIKeys, results.APIKeys...)
	r.TotalCount += uint32(len(results.APIKeys))
	return uint32(len(results.APIKeys)), nil
}

// ListApplicationsRequest: list applications request.
type ListApplicationsRequest struct {
	// OrderBy: criteria for sorting results.
	// Default value: created_at_asc
	OrderBy ListApplicationsRequestOrderBy `json:"-"`

	// PageSize: number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// Page: page number. Value must be greater than 1.
	Page *int32 `json:"-"`

	// Name: name of the application to filter.
	Name *string `json:"-"`

	// OrganizationID: ID of the Organization to filter.
	OrganizationID *string `json:"-"`

	// Editable: defines whether to filter out editable applications or not.
	Editable *bool `json:"-"`

	// Visible: defines whether to filter out visible applications or not.
	Visible *bool `json:"-"`
}

// ListApplicationsResponse: list applications response.
type ListApplicationsResponse struct {
	// Applications: list of applications.
	Applications []*Application `json:"applications"`

	// TotalCount: total count of applications.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListApplicationsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListApplicationsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListApplicationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Applications = append(r.Applications, results.Applications...)
	r.TotalCount += uint32(len(results.Applications))
	return uint32(len(results.Applications)), nil
}

// ListCustomerLevelsRequest: list customer levels request.
type ListCustomerLevelsRequest struct {
	CustomerLevelNames []string `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListCustomerLevelsRequestOrderBy `json:"-"`

	OrganizationID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListCustomerLevelsResponse: list customer levels response.
type ListCustomerLevelsResponse struct {
	CustomerLevels []*CustomerLevel `json:"customer_levels"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCustomerLevelsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCustomerLevelsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListCustomerLevelsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.CustomerLevels = append(r.CustomerLevels, results.CustomerLevels...)
	r.TotalCount += uint64(len(results.CustomerLevels))
	return uint64(len(results.CustomerLevels)), nil
}

// ListGroupsRequest: list groups request.
type ListGroupsRequest struct {
	// OrderBy: sort order of groups.
	// Default value: created_at_asc
	OrderBy ListGroupsRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`

	// PageSize: number of items per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// OrganizationID: filter by Organization ID.
	OrganizationID *string `json:"-"`

	// Name: name of group to find.
	Name *string `json:"-"`

	// ApplicationID: filter by a list of application IDs.
	// Precisely one of ApplicationID, UserID must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// UserID: filter by a list of user IDs.
	// Precisely one of ApplicationID, UserID must be set.
	UserID *string `json:"user_id,omitempty"`

	// Editable: defines whether to filter out editable groups or not.
	Editable *bool `json:"-"`

	// Visible: defines whether to filter out visible groups or not.
	Visible *bool `json:"-"`

	// Managed: defines whether to filter out managed groups or not.
	Managed *bool `json:"-"`
}

// ListGroupsResponse: list groups response.
type ListGroupsResponse struct {
	// Groups: list of groups.
	Groups []*Group `json:"groups"`

	// TotalCount: total count of groups.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListGroupsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListGroupsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListGroupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Groups = append(r.Groups, results.Groups...)
	r.TotalCount += uint32(len(results.Groups))
	return uint32(len(results.Groups)), nil
}

// ListInternalOrganizationsRequest: list internal organizations request.
type ListInternalOrganizationsRequest struct {
	// OrderBy: default value: created_at_asc
	OrderBy ListInternalOrganizationsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListInternalOrganizationsResponse: list internal organizations response.
type ListInternalOrganizationsResponse struct {
	Organizations []*InternalOrganization `json:"organizations"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListInternalOrganizationsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListInternalOrganizationsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListInternalOrganizationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Organizations = append(r.Organizations, results.Organizations...)
	r.TotalCount += uint64(len(results.Organizations))
	return uint64(len(results.Organizations)), nil
}

// ListJWTsRequest: list jw ts request.
type ListJWTsRequest struct {
	// OrderBy: criteria for sorting results.
	// Default value: created_at_asc
	OrderBy ListJWTsRequestOrderBy `json:"-"`

	// AudienceID: ID of the user to search.
	AudienceID *string `json:"-"`

	// PageSize: number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// Page: page number. Value must be greater to 1.
	Page *int32 `json:"-"`
}

// ListJWTsResponse: list jw ts response.
type ListJWTsResponse struct {
	Jwts []*JWT `json:"jwts"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListJWTsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListJWTsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListJWTsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Jwts = append(r.Jwts, results.Jwts...)
	r.TotalCount += uint32(len(results.Jwts))
	return uint32(len(results.Jwts)), nil
}

// ListLimitsRequest: list limits request.
type ListLimitsRequest struct {
	CustomerLevelID *string `json:"-"`

	// OrderBy: default value: quotum_name_asc
	OrderBy ListLimitsRequestOrderBy `json:"-"`

	OrganizationID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	QuotumID *string `json:"-"`
}

// ListLimitsResponse: list limits response.
type ListLimitsResponse struct {
	Limits []*Limit `json:"limits"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLimitsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLimitsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListLimitsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Limits = append(r.Limits, results.Limits...)
	r.TotalCount += uint64(len(results.Limits))
	return uint64(len(results.Limits)), nil
}

// ListLogsRequest: list logs request.
type ListLogsRequest struct {
	// OrderBy: criteria for sorting results.
	// Default value: created_at_asc
	OrderBy ListLogsRequestOrderBy `json:"-"`

	// OrganizationID: filter by Organization ID.
	OrganizationID *string `json:"-"`

	// PageSize: number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// Page: page number. Value must be greater to 1.
	Page *int32 `json:"-"`

	// CreatedAfter: defined whether or not to filter out logs created after this timestamp.
	CreatedAfter *time.Time `json:"-"`

	// CreatedBefore: defined whether or not to filter out logs created before this timestamp.
	CreatedBefore *time.Time `json:"-"`

	// Action: defined whether or not to filter out by a specific action.
	// Default value: unknown_action
	Action LogAction `json:"-"`

	// ResourceType: defined whether or not to filter out by a specific type of resource.
	// Default value: unknown_resource_type
	ResourceType LogResourceType `json:"-"`

	IP *string `json:"-"`

	Country *string `json:"-"`

	FromPrincipalID *string `json:"-"`

	ResourceID *string `json:"-"`
}

// ListLogsResponse: list logs response.
type ListLogsResponse struct {
	// Logs: list of logs.
	Logs []*Log `json:"logs"`

	// TotalCount: total count of logs.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLogsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListLogsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Logs = append(r.Logs, results.Logs...)
	r.TotalCount += uint64(len(results.Logs))
	return uint64(len(results.Logs)), nil
}

// ListOrganizationsRequest: list organizations request.
type ListOrganizationsRequest struct {
	CustomerLevelID *string `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListOrganizationsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListOrganizationsResponse: list organizations response.
type ListOrganizationsResponse struct {
	Organizations []*Organization `json:"organizations"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOrganizationsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOrganizationsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListOrganizationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Organizations = append(r.Organizations, results.Organizations...)
	r.TotalCount += uint64(len(results.Organizations))
	return uint64(len(results.Organizations)), nil
}

// ListPermissionSetsRequest: list permission sets request.
type ListPermissionSetsRequest struct {
	// OrderBy: criteria for sorting results.
	// Default value: name_asc
	OrderBy ListPermissionSetsRequestOrderBy `json:"-"`

	// PageSize: number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// Page: page number. Value must be greater than 1.
	Page *int32 `json:"-"`

	// OrganizationID: filter by Organization ID.
	OrganizationID *string `json:"-"`

	PermissionSetNames []string `json:"-"`

	PermissionSetIDs []string `json:"-"`

	CustomerLevelID *string `json:"-"`
}

// ListPermissionSetsResponse: list permission sets response.
type ListPermissionSetsResponse struct {
	// PermissionSets: list of permission sets.
	PermissionSets []*PermissionSet `json:"permission_sets"`

	// TotalCount: total count of permission sets.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPermissionSetsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPermissionSetsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPermissionSetsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PermissionSets = append(r.PermissionSets, results.PermissionSets...)
	r.TotalCount += uint32(len(results.PermissionSets))
	return uint32(len(results.PermissionSets)), nil
}

// ListPoliciesRequest: list policies request.
type ListPoliciesRequest struct {
	// OrderBy: criteria for sorting results.
	// Default value: policy_name_asc
	OrderBy ListPoliciesRequestOrderBy `json:"-"`

	// PageSize: number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// Page: page number. Value must be greater than 1.
	Page *int32 `json:"-"`

	// OrganizationID: ID of the Organization to filter.
	OrganizationID *string `json:"-"`

	// Editable: defines whether or not filter out editable policies.
	Editable *bool `json:"-"`

	// UserID: defines whether or not to filter by list of user IDs.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	UserID *string `json:"user_id,omitempty"`

	// GroupID: defines whether or not to filter by list of group IDs.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	GroupID *string `json:"group_id,omitempty"`

	// ApplicationID: filter by a list of application IDs.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// NoPrincipal: defines whether or not the policy is attributed to a principal.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	NoPrincipal *bool `json:"no_principal,omitempty"`

	// Visible: defines if the policy is visible or not.
	Visible *bool `json:"-"`
}

// ListPoliciesResponse: list policies response.
type ListPoliciesResponse struct {
	// Policies: list of policies.
	Policies []*Policy `json:"policies"`

	// TotalCount: total count of policies.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPoliciesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPoliciesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPoliciesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Policies = append(r.Policies, results.Policies...)
	r.TotalCount += uint32(len(results.Policies))
	return uint32(len(results.Policies)), nil
}

// ListQuotaRequest: list quota request.
type ListQuotaRequest struct {
	// OrderBy: criteria for sorting results.
	// Default value: name_asc
	OrderBy ListQuotaRequestOrderBy `json:"-"`

	// PageSize: number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// Page: page number. Value must be greater than 1.
	Page *int32 `json:"-"`
}

// ListQuotaResponse: list quota response.
type ListQuotaResponse struct {
	// Quota: list of quota.
	Quota []*Quotum `json:"quota"`

	// TotalCount: total count of quota.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListQuotaResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListQuotaResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListQuotaResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Quota = append(r.Quota, results.Quota...)
	r.TotalCount += uint64(len(results.Quota))
	return uint64(len(results.Quota)), nil
}

// ListRulesRequest: list rules request.
type ListRulesRequest struct {
	// PolicyID: id of policy to search.
	PolicyID *string `json:"-"`

	// PageSize: number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// Page: page number. Value must be greater than 1.
	Page *int32 `json:"-"`
}

// ListRulesResponse: list rules response.
type ListRulesResponse struct {
	// Rules: rules of the policy.
	Rules []*Rule `json:"rules"`

	// TotalCount: total count of rules.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRulesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRulesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListRulesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Rules = append(r.Rules, results.Rules...)
	r.TotalCount += uint32(len(results.Rules))
	return uint32(len(results.Rules)), nil
}

// ListSSHKeysRequest: list ssh keys request.
type ListSSHKeysRequest struct {
	// OrderBy: sort order of the SSH keys.
	// Default value: created_at_asc
	OrderBy ListSSHKeysRequestOrderBy `json:"-"`

	// Page: requested page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`

	// PageSize: number of items per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// OrganizationID: filter by Organization ID.
	OrganizationID *string `json:"-"`

	// Name: name of group to find.
	Name *string `json:"-"`

	// ProjectID: filter by Project ID.
	ProjectID *string `json:"-"`

	// Disabled: defines whether to include disabled SSH keys or not.
	Disabled *bool `json:"-"`
}

// ListSSHKeysResponse: list ssh keys response.
type ListSSHKeysResponse struct {
	// SSHKeys: list of SSH keys.
	SSHKeys []*SSHKey `json:"ssh_keys"`

	// TotalCount: total count of SSH keys.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSSHKeysResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSSHKeysResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSSHKeysResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SSHKeys = append(r.SSHKeys, results.SSHKeys...)
	r.TotalCount += uint32(len(results.SSHKeys))
	return uint32(len(results.SSHKeys)), nil
}

// ListUsersRequest: list users request.
type ListUsersRequest struct {
	// OrderBy: criteria for sorting results.
	// Default value: created_at_asc
	OrderBy ListUsersRequestOrderBy `json:"-"`

	// PageSize: number of results per page. Value must be between 1 and 100.
	PageSize *uint32 `json:"-"`

	// Page: page number. Value must be greater or equal to 1.
	Page *int32 `json:"-"`

	// OrganizationID: ID of the Organization to filter.
	OrganizationID *string `json:"-"`

	// AccountRootUserID: user ID of the root account.
	AccountRootUserID *string `json:"-"`
}

// ListUsersResponse: list users response.
type ListUsersResponse struct {
	// Users: list of users.
	Users []*User `json:"users"`

	// TotalCount: total count of users.
	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Users = append(r.Users, results.Users...)
	r.TotalCount += uint32(len(results.Users))
	return uint32(len(results.Users)), nil
}

// OrganizationCustomerLevels: organization customer levels.
type OrganizationCustomerLevels struct {
	CustomerLevelIDs []string `json:"customer_level_ids"`

	OrganizationID string `json:"organization_id"`
}

// SetCustomerLevelPermissionSetsRequest: set customer level permission sets request.
type SetCustomerLevelPermissionSetsRequest struct {
	CustomerLevelID string `json:"-"`

	PermissionSetIDs []string `json:"permission_set_ids"`
}

// SetInternalOrganizationPermissionSetsRequest: set internal organization permission sets request.
type SetInternalOrganizationPermissionSetsRequest struct {
	OrganizationID string `json:"-"`

	PermissionSetIDs []string `json:"permission_set_ids"`
}

// SetOrganizationCustomerLevelsRequest: set organization customer levels request.
type SetOrganizationCustomerLevelsRequest struct {
	CustomerLevelIDs []string `json:"customer_level_ids"`

	OrganizationID string `json:"organization_id"`
}

// SetRulesRequest: set rules request.
type SetRulesRequest struct {
	// PolicyID: id of policy to update.
	PolicyID string `json:"policy_id"`

	// Rules: rules of the policy to set.
	Rules []*RuleSpecs `json:"rules"`
}

// SetRulesResponse: set rules response.
type SetRulesResponse struct {
	// Rules: rules of the policy.
	Rules []*Rule `json:"rules"`
}

// UnauthenticatedAPIRenewJWTRequest: unauthenticated api renew jwt request.
type UnauthenticatedAPIRenewJWTRequest struct {
	Jti string `json:"-"`

	RenewToken string `json:"renew_token"`
}

// UpdateAPIKeyRequest: update api key request.
type UpdateAPIKeyRequest struct {
	// AccessKey: access key to update.
	AccessKey string `json:"-"`

	// DefaultProjectID: new default Project ID to set.
	DefaultProjectID *string `json:"default_project_id,omitempty"`

	// Description: new description to update.
	Description *string `json:"description,omitempty"`

	// Editable: whether the access key is editable.
	Editable *bool `json:"editable,omitempty"`

	// Visible: make the api key visible or not.
	Visible *bool `json:"visible,omitempty"`
}

// UpdateApplicationRequest: update application request.
type UpdateApplicationRequest struct {
	// ApplicationID: ID of the application to update.
	ApplicationID string `json:"-"`

	// Name: new name for the application (max length is 64 chars).
	Name *string `json:"name,omitempty"`

	// Description: new description for the application (max length is 200 chars).
	Description *string `json:"description,omitempty"`

	Editable *bool `json:"editable,omitempty"`

	Visible *bool `json:"visible,omitempty"`
}

// UpdateGroupRequest: update group request.
type UpdateGroupRequest struct {
	// GroupID: ID of the group to update.
	GroupID string `json:"-"`

	// Name: new name for the group (max length is 64 chars). MUST be unique inside an Organization.
	Name *string `json:"name,omitempty"`

	// Description: new description for the group (max length is 200 chars).
	Description *string `json:"description,omitempty"`

	// Editable: defines whether or not the group is editable.
	Editable *bool `json:"editable,omitempty"`

	// Visible: defines whether or not the group is visible.
	Visible *bool `json:"visible,omitempty"`

	// Managed: defines whether or not the group is managed.
	Managed *bool `json:"managed,omitempty"`
}

// UpdateLimitRequest: update limit request.
type UpdateLimitRequest struct {
	LimitID string `json:"-"`

	// Precisely one of Disabled, Limit, Unlimited must be set.
	Disabled *bool `json:"disabled,omitempty"`

	// Precisely one of Disabled, Limit, Unlimited must be set.
	Limit *uint64 `json:"limit,omitempty"`

	// Precisely one of Disabled, Limit, Unlimited must be set.
	Unlimited *bool `json:"unlimited,omitempty"`
}

// UpdatePolicyRequest: update policy request.
type UpdatePolicyRequest struct {
	// PolicyID: id of policy to update.
	PolicyID string `json:"-"`

	// Name: new name for the policy (max length is 64 characters).
	Name *string `json:"name,omitempty"`

	// Description: new description of policy (max length is 200 characters).
	Description *string `json:"description,omitempty"`

	// UserID: new ID of user attributed to the policy.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	UserID *string `json:"user_id,omitempty"`

	// GroupID: new ID of group attributed to the policy.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	GroupID *string `json:"group_id,omitempty"`

	// ApplicationID: new ID of application attributed to the policy.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// NoPrincipal: defines whether or not the policy is attributed to a principal.
	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	NoPrincipal *bool `json:"no_principal,omitempty"`
}

// UpdateQuotumRequest: update quotum request.
type UpdateQuotumRequest struct {
	QuotumID string `json:"-"`

	Name *string `json:"name,omitempty"`
}

// UpdateSSHKeyRequest: update ssh key request.
type UpdateSSHKeyRequest struct {
	SSHKeyID string `json:"-"`

	// Name: name of the SSH key. Max length is 1000.
	Name *string `json:"name,omitempty"`

	// Disabled: enable or disable the SSH key.
	Disabled *bool `json:"disabled,omitempty"`
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

// CreateCustomerLevel:
func (s *API) CreateCustomerLevel(req *CreateCustomerLevelRequest, opts ...scw.RequestOption) (*CustomerLevel, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam-admin/v1/customer-levels",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CustomerLevel

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCustomerLevels:
func (s *API) ListCustomerLevels(req *ListCustomerLevelsRequest, opts ...scw.RequestOption) (*ListCustomerLevelsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "customer_level_names", req.CustomerLevelNames)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/customer-levels",
		Query:  query,
	}

	var resp ListCustomerLevelsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCustomerLevel:
func (s *API) DeleteCustomerLevel(req *DeleteCustomerLevelRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.CustomerLevelID) == "" {
		return errors.New("field CustomerLevelID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/customer-levels/" + fmt.Sprint(req.CustomerLevelID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// SetCustomerLevelPermissionSets:
func (s *API) SetCustomerLevelPermissionSets(req *SetCustomerLevelPermissionSetsRequest, opts ...scw.RequestOption) (*CustomerLevel, error) {
	var err error

	if fmt.Sprint(req.CustomerLevelID) == "" {
		return nil, errors.New("field CustomerLevelID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/iam-admin/v1/customer-levels/" + fmt.Sprint(req.CustomerLevelID) + "/permission-sets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CustomerLevel

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddCustomerLevelPermissionSet:
func (s *API) AddCustomerLevelPermissionSet(req *AddCustomerLevelPermissionSetRequest, opts ...scw.RequestOption) (*CustomerLevel, error) {
	var err error

	if fmt.Sprint(req.CustomerLevelID) == "" {
		return nil, errors.New("field CustomerLevelID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam-admin/v1/customer-levels/" + fmt.Sprint(req.CustomerLevelID) + "/permission-sets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CustomerLevel

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCustomerLevelPermissionSet:
func (s *API) DeleteCustomerLevelPermissionSet(req *DeleteCustomerLevelPermissionSetRequest, opts ...scw.RequestOption) (*CustomerLevel, error) {
	var err error

	if fmt.Sprint(req.CustomerLevelID) == "" {
		return nil, errors.New("field CustomerLevelID cannot be empty in request")
	}

	if fmt.Sprint(req.PermissionSetID) == "" {
		return nil, errors.New("field PermissionSetID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/customer-levels/" + fmt.Sprint(req.CustomerLevelID) + "/permission-sets/" + fmt.Sprint(req.PermissionSetID) + "",
	}

	var resp CustomerLevel

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetOrganizationCustomerLevels:
func (s *API) SetOrganizationCustomerLevels(req *SetOrganizationCustomerLevelsRequest, opts ...scw.RequestOption) (*OrganizationCustomerLevels, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/iam-admin/v1/organization-customer-levels",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp OrganizationCustomerLevels

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddOrganizationCustomerLevel:
func (s *API) AddOrganizationCustomerLevel(req *AddOrganizationCustomerLevelRequest, opts ...scw.RequestOption) (*OrganizationCustomerLevels, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam-admin/v1/organization-customer-levels",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp OrganizationCustomerLevels

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteOrganizationCustomerLevel:
func (s *API) DeleteOrganizationCustomerLevel(req *DeleteOrganizationCustomerLevelRequest, opts ...scw.RequestOption) (*OrganizationCustomerLevels, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "customer_level_id", req.CustomerLevelID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/organization-customer-levels",
		Query:  query,
	}

	var resp OrganizationCustomerLevels

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListQuota:
func (s *API) ListQuota(req *ListQuotaRequest, opts ...scw.RequestOption) (*ListQuotaResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/quota",
		Query:  query,
	}

	var resp ListQuotaResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateQuotum:
func (s *API) CreateQuotum(req *CreateQuotumRequest, opts ...scw.RequestOption) (*Quotum, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam-admin/v1/quota",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Quotum

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateQuotum:
func (s *API) UpdateQuotum(req *UpdateQuotumRequest, opts ...scw.RequestOption) (*Quotum, error) {
	var err error

	if fmt.Sprint(req.QuotumID) == "" {
		return nil, errors.New("field QuotumID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iam-admin/v1/quota/" + fmt.Sprint(req.QuotumID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Quotum

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteQuotum:
func (s *API) DeleteQuotum(req *DeleteQuotumRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.QuotumID) == "" {
		return errors.New("field QuotumID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/quota/" + fmt.Sprint(req.QuotumID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListLimits:
func (s *API) ListLimits(req *ListLimitsRequest, opts ...scw.RequestOption) (*ListLimitsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "customer_level_id", req.CustomerLevelID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "quotum_id", req.QuotumID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/limits",
		Query:  query,
	}

	var resp ListLimitsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateLimit:
func (s *API) CreateLimit(req *CreateLimitRequest, opts ...scw.RequestOption) (*Limit, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam-admin/v1/limits",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Limit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateLimit:
func (s *API) UpdateLimit(req *UpdateLimitRequest, opts ...scw.RequestOption) (*Limit, error) {
	var err error

	if fmt.Sprint(req.LimitID) == "" {
		return nil, errors.New("field LimitID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iam-admin/v1/limits/" + fmt.Sprint(req.LimitID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Limit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteLimit:
func (s *API) DeleteLimit(req *DeleteLimitRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.LimitID) == "" {
		return errors.New("field LimitID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/limits/" + fmt.Sprint(req.LimitID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListSSHKeys:
func (s *API) ListSSHKeys(req *ListSSHKeysRequest, opts ...scw.RequestOption) (*ListSSHKeysResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "disabled", req.Disabled)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/ssh-keys",
		Query:  query,
	}

	var resp ListSSHKeysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSSHKey:
func (s *API) CreateSSHKey(req *CreateSSHKeyRequest, opts ...scw.RequestOption) (*SSHKey, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("key")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam-admin/v1/ssh-keys",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SSHKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSSHKey:
func (s *API) GetSSHKey(req *GetSSHKeyRequest, opts ...scw.RequestOption) (*SSHKey, error) {
	var err error

	if fmt.Sprint(req.SSHKeyID) == "" {
		return nil, errors.New("field SSHKeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/ssh-keys/" + fmt.Sprint(req.SSHKeyID) + "",
	}

	var resp SSHKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateSSHKey: Update the parameters of an SSH key, including `name` and `disable`.
func (s *API) UpdateSSHKey(req *UpdateSSHKeyRequest, opts ...scw.RequestOption) (*SSHKey, error) {
	var err error

	if fmt.Sprint(req.SSHKeyID) == "" {
		return nil, errors.New("field SSHKeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iam-admin/v1/ssh-keys/" + fmt.Sprint(req.SSHKeyID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SSHKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSSHKey:
func (s *API) DeleteSSHKey(req *DeleteSSHKeyRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.SSHKeyID) == "" {
		return errors.New("field SSHKeyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/ssh-keys/" + fmt.Sprint(req.SSHKeyID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListUsers:
func (s *API) ListUsers(req *ListUsersRequest, opts ...scw.RequestOption) (*ListUsersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "account_root_user_id", req.AccountRootUserID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/users",
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

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/users/" + fmt.Sprint(req.UserID) + "",
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListApplications:
func (s *API) ListApplications(req *ListApplicationsRequest, opts ...scw.RequestOption) (*ListApplicationsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "editable", req.Editable)
	parameter.AddToQuery(query, "visible", req.Visible)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/applications",
		Query:  query,
	}

	var resp ListApplicationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetApplication:
func (s *API) GetApplication(req *GetApplicationRequest, opts ...scw.RequestOption) (*Application, error) {
	var err error

	if fmt.Sprint(req.ApplicationID) == "" {
		return nil, errors.New("field ApplicationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/applications/" + fmt.Sprint(req.ApplicationID) + "",
	}

	var resp Application

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateApplication:
func (s *API) UpdateApplication(req *UpdateApplicationRequest, opts ...scw.RequestOption) (*Application, error) {
	var err error

	if fmt.Sprint(req.ApplicationID) == "" {
		return nil, errors.New("field ApplicationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iam-admin/v1/applications/" + fmt.Sprint(req.ApplicationID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Application

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteApplication:
func (s *API) DeleteApplication(req *DeleteApplicationRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.ApplicationID) == "" {
		return errors.New("field ApplicationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/applications/" + fmt.Sprint(req.ApplicationID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListGroups:
func (s *API) ListGroups(req *ListGroupsRequest, opts ...scw.RequestOption) (*ListGroupsResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "editable", req.Editable)
	parameter.AddToQuery(query, "visible", req.Visible)
	parameter.AddToQuery(query, "managed", req.Managed)
	parameter.AddToQuery(query, "application_id", req.ApplicationID)
	parameter.AddToQuery(query, "user_id", req.UserID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/groups",
		Query:  query,
	}

	var resp ListGroupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateGroup:
func (s *API) CreateGroup(req *CreateGroupRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("grp")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam-admin/v1/groups",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGroup:
func (s *API) GetGroup(req *GetGroupRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateGroup:
func (s *API) UpdateGroup(req *UpdateGroupRequest, opts ...scw.RequestOption) (*Group, error) {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return nil, errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iam-admin/v1/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Group

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteGroup:
func (s *API) DeleteGroup(req *DeleteGroupRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.GroupID) == "" {
		return errors.New("field GroupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/groups/" + fmt.Sprint(req.GroupID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListPolicies:
func (s *API) ListPolicies(req *ListPoliciesRequest, opts ...scw.RequestOption) (*ListPoliciesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "editable", req.Editable)
	parameter.AddToQuery(query, "visible", req.Visible)
	parameter.AddToQuery(query, "user_id", req.UserID)
	parameter.AddToQuery(query, "group_id", req.GroupID)
	parameter.AddToQuery(query, "application_id", req.ApplicationID)
	parameter.AddToQuery(query, "no_principal", req.NoPrincipal)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/policies",
		Query:  query,
	}

	var resp ListPoliciesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreatePolicy:
func (s *API) CreatePolicy(req *CreatePolicyRequest, opts ...scw.RequestOption) (*Policy, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("pol")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam-admin/v1/policies",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Policy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPolicy:
func (s *API) GetPolicy(req *GetPolicyRequest, opts ...scw.RequestOption) (*Policy, error) {
	var err error

	if fmt.Sprint(req.PolicyID) == "" {
		return nil, errors.New("field PolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/policies/" + fmt.Sprint(req.PolicyID) + "",
	}

	var resp Policy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePolicy:
func (s *API) UpdatePolicy(req *UpdatePolicyRequest, opts ...scw.RequestOption) (*Policy, error) {
	var err error

	if fmt.Sprint(req.PolicyID) == "" {
		return nil, errors.New("field PolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iam-admin/v1/policies/" + fmt.Sprint(req.PolicyID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Policy

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeletePolicy:
func (s *API) DeletePolicy(req *DeletePolicyRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.PolicyID) == "" {
		return errors.New("field PolicyID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/policies/" + fmt.Sprint(req.PolicyID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// SetRules:
func (s *API) SetRules(req *SetRulesRequest, opts ...scw.RequestOption) (*SetRulesResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/iam-admin/v1/rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRules:
func (s *API) ListRules(req *ListRulesRequest, opts ...scw.RequestOption) (*ListRulesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "policy_id", req.PolicyID)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/rules",
		Query:  query,
	}

	var resp ListRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPermissionSets:
func (s *API) ListPermissionSets(req *ListPermissionSetsRequest, opts ...scw.RequestOption) (*ListPermissionSetsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "permission_set_names", req.PermissionSetNames)
	parameter.AddToQuery(query, "permission_set_ids", req.PermissionSetIDs)
	parameter.AddToQuery(query, "customer_level_id", req.CustomerLevelID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/permission-sets",
		Query:  query,
	}

	var resp ListPermissionSetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPermissionSet:
func (s *API) GetPermissionSet(req *GetPermissionSetRequest, opts ...scw.RequestOption) (*GetPermissionSetResponse, error) {
	var err error

	if fmt.Sprint(req.PermissionSetName) == "" {
		return nil, errors.New("field PermissionSetName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/permission-sets/" + fmt.Sprint(req.PermissionSetName) + "",
	}

	var resp GetPermissionSetResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetPermissions:
func (s *API) GetPermissions(req *GetPermissionsRequest, opts ...scw.RequestOption) (*GetPermissionsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "auth_id", req.AuthID)
	parameter.AddToQuery(query, "principal_id", req.PrincipalID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/permissions",
		Query:  query,
	}

	var resp GetPermissionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListAPIKeys:
func (s *API) ListAPIKeys(req *ListAPIKeysRequest, opts ...scw.RequestOption) (*ListAPIKeysResponse, error) {
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
	parameter.AddToQuery(query, "editable", req.Editable)
	parameter.AddToQuery(query, "visible", req.Visible)
	parameter.AddToQuery(query, "category", req.Category)
	parameter.AddToQuery(query, "application_id", req.ApplicationID)
	parameter.AddToQuery(query, "user_id", req.UserID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/api-keys",
		Query:  query,
	}

	var resp ListAPIKeysResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateAPIKey:
func (s *API) CreateAPIKey(req *CreateAPIKeyRequest, opts ...scw.RequestOption) (*APIKey, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam-admin/v1/api-keys",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp APIKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateProductAPIKey:
func (s *API) CreateProductAPIKey(req *CreateProductAPIKeyRequest, opts ...scw.RequestOption) (*APIKey, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam-admin/v1/product-api-keys",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp APIKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetAPIKey:
func (s *API) GetAPIKey(req *GetAPIKeyRequest, opts ...scw.RequestOption) (*APIKey, error) {
	var err error

	if fmt.Sprint(req.AccessKey) == "" {
		return nil, errors.New("field AccessKey cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/api-keys/" + fmt.Sprint(req.AccessKey) + "",
	}

	var resp APIKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAPIKey:
func (s *API) UpdateAPIKey(req *UpdateAPIKeyRequest, opts ...scw.RequestOption) (*APIKey, error) {
	var err error

	if fmt.Sprint(req.AccessKey) == "" {
		return nil, errors.New("field AccessKey cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iam-admin/v1/api-keys/" + fmt.Sprint(req.AccessKey) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp APIKey

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteAPIKey:
func (s *API) DeleteAPIKey(req *DeleteAPIKeyRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.AccessKey) == "" {
		return errors.New("field AccessKey cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/api-keys/" + fmt.Sprint(req.AccessKey) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// DeleteProductAPIKey:
func (s *API) DeleteProductAPIKey(req *DeleteProductAPIKeyRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.AccessKey) == "" {
		return errors.New("field AccessKey cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/product-api-keys/" + fmt.Sprint(req.AccessKey) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListJWTs:
func (s *API) ListJWTs(req *ListJWTsRequest, opts ...scw.RequestOption) (*ListJWTsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "audience_id", req.AudienceID)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/jwts",
		Query:  query,
	}

	var resp ListJWTsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetJWT:
func (s *API) GetJWT(req *GetJWTRequest, opts ...scw.RequestOption) (*JWT, error) {
	var err error

	if fmt.Sprint(req.Jti) == "" {
		return nil, errors.New("field Jti cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/jwts/" + fmt.Sprint(req.Jti) + "",
	}

	var resp JWT

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteJWT:
func (s *API) DeleteJWT(req *DeleteJWTRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.Jti) == "" {
		return errors.New("field Jti cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/jwts/" + fmt.Sprint(req.Jti) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListInternalOrganizations:
func (s *API) ListInternalOrganizations(req *ListInternalOrganizationsRequest, opts ...scw.RequestOption) (*ListInternalOrganizationsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/internal-organizations",
		Query:  query,
	}

	var resp ListInternalOrganizationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetInternalOrganization:
func (s *API) GetInternalOrganization(req *GetInternalOrganizationRequest, opts ...scw.RequestOption) (*InternalOrganization, error) {
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
		Path:   "/iam-admin/v1/internal-organizations/" + fmt.Sprint(req.OrganizationID) + "",
	}

	var resp InternalOrganization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetInternalOrganizationPermissionSets:
func (s *API) SetInternalOrganizationPermissionSets(req *SetInternalOrganizationPermissionSetsRequest, opts ...scw.RequestOption) (*InternalOrganization, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/iam-admin/v1/internal-organizations/" + fmt.Sprint(req.OrganizationID) + "/permission-sets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp InternalOrganization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// AddInternalOrganizationPermissionSet:
func (s *API) AddInternalOrganizationPermissionSet(req *AddInternalOrganizationPermissionSetRequest, opts ...scw.RequestOption) (*InternalOrganization, error) {
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
		Path:   "/iam-admin/v1/internal-organizations/" + fmt.Sprint(req.OrganizationID) + "/permission-sets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp InternalOrganization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteInternalOrganizationPermissionSet:
func (s *API) DeleteInternalOrganizationPermissionSet(req *DeleteInternalOrganizationPermissionSetRequest, opts ...scw.RequestOption) (*InternalOrganization, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	if fmt.Sprint(req.PermissionSetID) == "" {
		return nil, errors.New("field PermissionSetID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iam-admin/v1/internal-organizations/" + fmt.Sprint(req.OrganizationID) + "/permission-sets/" + fmt.Sprint(req.PermissionSetID) + "",
	}

	var resp InternalOrganization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOrganizations:
func (s *API) ListOrganizations(req *ListOrganizationsRequest, opts ...scw.RequestOption) (*ListOrganizationsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "customer_level_id", req.CustomerLevelID)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/organizations",
		Query:  query,
	}

	var resp ListOrganizationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListLogs:
func (s *API) ListLogs(req *ListLogsRequest, opts ...scw.RequestOption) (*ListLogsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "created_after", req.CreatedAfter)
	parameter.AddToQuery(query, "created_before", req.CreatedBefore)
	parameter.AddToQuery(query, "action", req.Action)
	parameter.AddToQuery(query, "resource_type", req.ResourceType)
	parameter.AddToQuery(query, "ip", req.IP)
	parameter.AddToQuery(query, "country", req.Country)
	parameter.AddToQuery(query, "from_principal_id", req.FromPrincipalID)
	parameter.AddToQuery(query, "resource_id", req.ResourceID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/logs",
		Query:  query,
	}

	var resp ListLogsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLog:
func (s *API) GetLog(req *GetLogRequest, opts ...scw.RequestOption) (*Log, error) {
	var err error

	if fmt.Sprint(req.LogID) == "" {
		return nil, errors.New("field LogID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1/logs/" + fmt.Sprint(req.LogID) + "",
	}

	var resp Log

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UnauthenticatedAPI struct {
	client *scw.Client
}

// NewUnauthenticatedAPI returns a UnauthenticatedAPI object from a Scaleway client.
func NewUnauthenticatedAPI(client *scw.Client) *UnauthenticatedAPI {
	return &UnauthenticatedAPI{
		client: client,
	}
}

// RenewJWT:
func (s *UnauthenticatedAPI) RenewJWT(req *UnauthenticatedAPIRenewJWTRequest, opts ...scw.RequestOption) (*EncodedJWT, error) {
	var err error

	if fmt.Sprint(req.Jti) == "" {
		return nil, errors.New("field Jti cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam-admin/v1/jwts/" + fmt.Sprint(req.Jti) + "/renew",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EncodedJWT

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

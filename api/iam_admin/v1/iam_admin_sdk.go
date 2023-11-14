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
	CreateProductAPIKeyRequestProductNameVpc  = CreateProductAPIKeyRequestProductName("vpc")
	CreateProductAPIKeyRequestProductNameEdge = CreateProductAPIKeyRequestProductName("edge")
	CreateProductAPIKeyRequestProductNameSbs  = CreateProductAPIKeyRequestProductName("sbs")
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
	// Creation date ascending.
	ListInternalOrganizationsRequestOrderByCreatedAtAsc = ListInternalOrganizationsRequestOrderBy("created_at_asc")
	// Creation date descending.
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
	PermissionSetNames *[]string `json:"permission_set_names"`

	// Precisely one of ProjectIDs, OrganizationID must be set.
	ProjectIDs *[]string `json:"project_ids,omitempty"`

	// Precisely one of ProjectIDs, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`
}

// JWT: jwt.
type JWT struct {
	Jti string `json:"jti"`

	IssuerID string `json:"issuer_id"`

	AudienceID string `json:"audience_id"`

	Staff bool `json:"staff"`

	Renewable bool `json:"renewable"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	ExpiresAt *time.Time `json:"expires_at"`

	IP net.IP `json:"ip"`

	UserAgent string `json:"user_agent"`
}

// PermissionSet: permission set.
type PermissionSet struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	// ScopeType: default value: unknown_scope_type
	ScopeType PermissionSetScopeType `json:"scope_type"`

	NbPermissions uint32 `json:"nb_permissions"`

	Categories *[]string `json:"categories"`
}

// Permission: permission.
type Permission struct {
	Name string `json:"name"`

	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID, AllOrganizations must be set.
	ProjectIDs *[]string `json:"project_ids,omitempty"`

	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID, AllOrganizations must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID, AllOrganizations must be set.
	AccountRootUserID *string `json:"account_root_user_id,omitempty"`

	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID, AllOrganizations must be set.
	AllOrganizations *bool `json:"all_organizations,omitempty"`
}

// APIKey: api key.
type APIKey struct {
	AccessKey string `json:"access_key"`

	SecretKey *string `json:"secret_key"`

	Description string `json:"description"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	ExpiredAt *time.Time `json:"expired_at"`

	DefaultProjectID string `json:"default_project_id"`

	IsStaff bool `json:"is_staff"`

	// Category: default value: unknown_category
	Category APIKeyCategory `json:"category"`

	Editable bool `json:"editable"`

	Visible bool `json:"visible"`

	// Precisely one of UserID, ApplicationID must be set.
	UserID *string `json:"user_id,omitempty"`

	// Precisely one of UserID, ApplicationID must be set.
	ApplicationID *string `json:"application_id,omitempty"`
}

// Application: application.
type Application struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	OrganizationID string `json:"organization_id"`

	Editable bool `json:"editable"`

	Visible bool `json:"visible"`

	NbAPIKeys uint32 `json:"nb_api_keys"`
}

// CustomerLevel: customer level.
type CustomerLevel struct {
	ID string `json:"id"`

	Name string `json:"name"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// Group: group.
type Group struct {
	ID string `json:"id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	OrganizationID string `json:"organization_id"`

	Name string `json:"name"`

	Description string `json:"description"`

	UserIDs []string `json:"user_ids"`

	ApplicationIDs []string `json:"application_ids"`
}

// InternalOrganization: internal organization.
type InternalOrganization struct {
	ID string `json:"id"`

	Creator *InternalOrganizationInternalUser `json:"creator"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	PermissionSetIDs []string `json:"permission_set_ids"`
}

// Limit: limit.
type Limit struct {
	ID string `json:"id"`

	QuotumID string `json:"quotum_id"`

	QuotumName string `json:"quotum_name"`

	CustomerLevelID string `json:"customer_level_id"`

	CustomerLevelName string `json:"customer_level_name"`

	// Precisely one of Limit, Unlimited, Disabled must be set.
	Limit *uint64 `json:"limit,omitempty"`

	// Precisely one of Limit, Unlimited, Disabled must be set.
	Unlimited *bool `json:"unlimited,omitempty"`

	// Precisely one of Limit, Unlimited, Disabled must be set.
	Disabled *bool `json:"disabled,omitempty"`
}

// Log: log.
type Log struct {
	ID string `json:"id"`

	CreatedAt *time.Time `json:"created_at"`

	IP net.IP `json:"ip"`

	UserAgent string `json:"user_agent"`

	Country string `json:"country"`

	// Action: default value: unknown_action
	Action LogAction `json:"action"`

	FromPrincipalID string `json:"from_principal_id"`

	OrganizationID string `json:"organization_id"`

	// ResourceType: default value: unknown_resource_type
	ResourceType LogResourceType `json:"resource_type"`

	ResourceID string `json:"resource_id"`

	Data *scw.JSONObject `json:"data"`

	Visible bool `json:"visible"`
}

// Organization: organization.
type Organization struct {
	ID string `json:"id"`

	CustomerClass string `json:"customer_class"`

	Locked bool `json:"locked"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// Policy: policy.
type Policy struct {
	ID string `json:"id"`

	Name string `json:"name"`

	OrganizationID string `json:"organization_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	Editable bool `json:"editable"`

	NbRules uint32 `json:"nb_rules"`

	NbScopes uint32 `json:"nb_scopes"`

	NbPermissionSets uint32 `json:"nb_permission_sets"`

	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	UserID *string `json:"user_id,omitempty"`

	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	GroupID *string `json:"group_id,omitempty"`

	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	NoPrincipal *bool `json:"no_principal,omitempty"`

	Description string `json:"description"`

	Visible bool `json:"visible"`
}

// Quotum: quotum.
type Quotum struct {
	ID string `json:"id"`

	Name string `json:"name"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`
}

// Rule: rule.
type Rule struct {
	ID string `json:"id"`

	PermissionSetNames *[]string `json:"permission_set_names"`

	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID, AllOrganizations must be set.
	ProjectIDs *[]string `json:"project_ids,omitempty"`

	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID, AllOrganizations must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID, AllOrganizations must be set.
	AccountRootUserID *string `json:"account_root_user_id,omitempty"`

	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID, AllOrganizations must be set.
	AllOrganizations *bool `json:"all_organizations,omitempty"`
}

// SSHKey: ssh key.
type SSHKey struct {
	ID string `json:"id"`

	Name string `json:"name"`

	PublicKey string `json:"public_key"`

	Fingerprint string `json:"fingerprint"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	OrganizationID string `json:"organization_id"`

	ProjectID string `json:"project_id"`

	Disabled bool `json:"disabled"`
}

// User: user.
type User struct {
	ID string `json:"id"`

	AccountRootUserID *string `json:"account_root_user_id"`

	Email string `json:"email"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	OrganizationID string `json:"organization_id"`

	Deletable bool `json:"deletable"`

	// Status: default value: unknown_status
	Status UserStatus `json:"status"`
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
	OrganizationID string `json:"organization_id"`

	CustomerLevelID string `json:"customer_level_id"`
}

// CreateAPIKeyRequest: create api key request.
type CreateAPIKeyRequest struct {
	// Precisely one of UserID, ApplicationID must be set.
	UserID *string `json:"user_id,omitempty"`

	// Precisely one of UserID, ApplicationID must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	Description string `json:"description"`

	ExpiresAt *time.Time `json:"expires_at,omitempty"`

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
	OrganizationID string `json:"organization_id"`

	Name string `json:"name"`

	Description string `json:"description"`
}

// CreateLimitRequest: create limit request.
type CreateLimitRequest struct {
	CustomerLevelID string `json:"customer_level_id"`

	QuotumID string `json:"quotum_id"`

	// Precisely one of Limit, Unlimited, Disabled must be set.
	Limit *uint64 `json:"limit,omitempty"`

	// Precisely one of Limit, Unlimited, Disabled must be set.
	Unlimited *bool `json:"unlimited,omitempty"`

	// Precisely one of Limit, Unlimited, Disabled must be set.
	Disabled *bool `json:"disabled,omitempty"`
}

// CreatePolicyRequest: create policy request.
type CreatePolicyRequest struct {
	Name string `json:"name"`

	OrganizationID string `json:"organization_id"`

	Rules []*RuleSpecs `json:"rules"`

	Description string `json:"description"`

	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	UserID *string `json:"user_id,omitempty"`

	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	GroupID *string `json:"group_id,omitempty"`

	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	NoPrincipal *bool `json:"no_principal,omitempty"`
}

// CreateProductAPIKeyRequest: create product api key request.
type CreateProductAPIKeyRequest struct {
	ResourceID string `json:"resource_id"`

	// ProductName: default value: unknown_product_name
	ProductName CreateProductAPIKeyRequestProductName `json:"product_name"`

	ProjectID string `json:"project_id"`

	PermissionSetNames []string `json:"permission_set_names"`

	Description string `json:"description"`

	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// CreateQuotumRequest: create quotum request.
type CreateQuotumRequest struct {
	Name string `json:"name"`
}

// CreateSSHKeyRequest: create ssh key request.
type CreateSSHKeyRequest struct {
	Name string `json:"name"`

	PublicKey string `json:"public_key"`

	ProjectID string `json:"project_id"`
}

// DeleteAPIKeyRequest: delete api key request.
type DeleteAPIKeyRequest struct {
	AccessKey string `json:"-"`
}

// DeleteApplicationRequest: delete application request.
type DeleteApplicationRequest struct {
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
	GroupID string `json:"-"`
}

// DeleteInternalOrganizationPermissionSetRequest: delete internal organization permission set request.
type DeleteInternalOrganizationPermissionSetRequest struct {
	OrganizationID string `json:"-"`

	PermissionSetID string `json:"-"`
}

// DeleteJWTRequest: delete jwt request.
type DeleteJWTRequest struct {
	Jti string `json:"-"`
}

// DeleteLimitRequest: delete limit request.
type DeleteLimitRequest struct {
	LimitID string `json:"-"`
}

// DeleteOrganizationCustomerLevelRequest: delete organization customer level request.
type DeleteOrganizationCustomerLevelRequest struct {
	OrganizationID string `json:"-"`

	CustomerLevelID string `json:"-"`
}

// DeletePolicyRequest: delete policy request.
type DeletePolicyRequest struct {
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
	AccessKey string `json:"-"`
}

// GetApplicationRequest: get application request.
type GetApplicationRequest struct {
	ApplicationID string `json:"-"`
}

// GetGroupRequest: get group request.
type GetGroupRequest struct {
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
	// Precisely one of PrincipalID, AuthID must be set.
	PrincipalID *string `json:"principal_id,omitempty"`

	// Precisely one of PrincipalID, AuthID must be set.
	AuthID *string `json:"auth_id,omitempty"`
}

// GetPermissionsResponse: get permissions response.
type GetPermissionsResponse struct {
	Permissions []*Permission `json:"permissions"`
}

// GetPolicyRequest: get policy request.
type GetPolicyRequest struct {
	PolicyID string `json:"-"`
}

// GetSSHKeyRequest: get ssh key request.
type GetSSHKeyRequest struct {
	SSHKeyID string `json:"-"`
}

// GetUserRequest: get user request.
type GetUserRequest struct {
	UserID string `json:"-"`
}

// ListAPIKeysRequest: list api keys request.
type ListAPIKeysRequest struct {
	// OrderBy: default value: created_at_asc
	OrderBy ListAPIKeysRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	Editable *bool `json:"-"`

	Visible *bool `json:"-"`

	// Category: default value: unknown_category
	Category APIKeyCategory `json:"-"`

	// Precisely one of UserID, ApplicationID must be set.
	UserID *string `json:"user_id,omitempty"`

	// Precisely one of UserID, ApplicationID must be set.
	ApplicationID *string `json:"application_id,omitempty"`
}

// ListAPIKeysResponse: list api keys response.
type ListAPIKeysResponse struct {
	APIKeys []*APIKey `json:"api_keys"`

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
	// OrderBy: default value: created_at_asc
	OrderBy ListApplicationsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	Editable *bool `json:"-"`

	Visible *bool `json:"-"`
}

// ListApplicationsResponse: list applications response.
type ListApplicationsResponse struct {
	Applications []*Application `json:"applications"`

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
	// OrderBy: default value: name_asc
	OrderBy ListCustomerLevelsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	CustomerLevelNames []string `json:"-"`
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
	// OrderBy: default value: created_at_asc
	OrderBy ListGroupsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	Name *string `json:"-"`

	// Precisely one of ApplicationID, UserID must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// Precisely one of ApplicationID, UserID must be set.
	UserID *string `json:"user_id,omitempty"`
}

// ListGroupsResponse: list groups response.
type ListGroupsResponse struct {
	Groups []*Group `json:"groups"`

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
	// OrderBy: default value: created_at_asc
	OrderBy ListJWTsRequestOrderBy `json:"-"`

	AudienceID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
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
	// OrderBy: default value: quotum_name_asc
	OrderBy ListLimitsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	CustomerLevelID *string `json:"-"`

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
	// OrderBy: default value: created_at_asc
	OrderBy ListLogsRequestOrderBy `json:"-"`

	CreatedAfter *time.Time `json:"-"`

	CreatedBefore *time.Time `json:"-"`

	IP *string `json:"-"`

	Country *string `json:"-"`

	// Action: default value: unknown_action
	Action LogAction `json:"-"`

	FromPrincipalID *string `json:"-"`

	OrganizationID *string `json:"-"`

	// ResourceType: default value: unknown_resource_type
	ResourceType LogResourceType `json:"-"`

	ResourceID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListLogsResponse: list logs response.
type ListLogsResponse struct {
	Logs []*Log `json:"logs"`

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
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListOrganizationsRequestOrderBy `json:"-"`

	CustomerLevelID *string `json:"-"`
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
	// OrderBy: default value: name_asc
	OrderBy ListPermissionSetsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	PermissionSetNames []string `json:"-"`

	PermissionSetIDs []string `json:"-"`

	CustomerLevelID *string `json:"-"`
}

// ListPermissionSetsResponse: list permission sets response.
type ListPermissionSetsResponse struct {
	PermissionSets []*PermissionSet `json:"permission_sets"`

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
	// OrderBy: default value: policy_name_asc
	OrderBy ListPoliciesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	Editable *bool `json:"-"`

	Visible *bool `json:"-"`

	// Precisely one of GroupID, UserID, ApplicationID, NoPrincipal must be set.
	GroupID *string `json:"group_id,omitempty"`

	// Precisely one of GroupID, UserID, ApplicationID, NoPrincipal must be set.
	UserID *string `json:"user_id,omitempty"`

	// Precisely one of GroupID, UserID, ApplicationID, NoPrincipal must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// Precisely one of GroupID, UserID, ApplicationID, NoPrincipal must be set.
	NoPrincipal *bool `json:"no_principal,omitempty"`
}

// ListPoliciesResponse: list policies response.
type ListPoliciesResponse struct {
	Policies []*Policy `json:"policies"`

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
	// OrderBy: default value: name_asc
	OrderBy ListQuotaRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListQuotaResponse: list quota response.
type ListQuotaResponse struct {
	Quota []*Quotum `json:"quota"`

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
	PolicyID *string `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ListRulesResponse: list rules response.
type ListRulesResponse struct {
	Rules []*Rule `json:"rules"`

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
	// OrderBy: default value: created_at_asc
	OrderBy ListSSHKeysRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	Name *string `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`

	Disabled *bool `json:"-"`
}

// ListSSHKeysResponse: list ssh keys response.
type ListSSHKeysResponse struct {
	SSHKeys []*SSHKey `json:"ssh_keys"`

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
	// OrderBy: default value: created_at_asc
	OrderBy ListUsersRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	AccountRootUserID *string `json:"-"`
}

// ListUsersResponse: list users response.
type ListUsersResponse struct {
	Users []*User `json:"users"`

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
	OrganizationID string `json:"organization_id"`

	CustomerLevelIDs []string `json:"customer_level_ids"`
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
	OrganizationID string `json:"organization_id"`

	CustomerLevelIDs []string `json:"customer_level_ids"`
}

// SetRulesRequest: set rules request.
type SetRulesRequest struct {
	PolicyID string `json:"policy_id"`

	Rules []*RuleSpecs `json:"rules"`
}

// SetRulesResponse: set rules response.
type SetRulesResponse struct {
	Rules []*Rule `json:"rules"`
}

// UnauthenticatedAPIRenewJWTRequest: unauthenticated api renew jwt request.
type UnauthenticatedAPIRenewJWTRequest struct {
	Jti string `json:"-"`

	RenewToken string `json:"renew_token"`
}

// UpdateAPIKeyRequest: update api key request.
type UpdateAPIKeyRequest struct {
	AccessKey string `json:"-"`

	DefaultProjectID *string `json:"default_project_id,omitempty"`

	Description *string `json:"description,omitempty"`

	Editable *bool `json:"editable,omitempty"`

	Visible *bool `json:"visible,omitempty"`
}

// UpdateApplicationRequest: update application request.
type UpdateApplicationRequest struct {
	ApplicationID string `json:"-"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Editable *bool `json:"editable,omitempty"`

	Visible *bool `json:"visible,omitempty"`
}

// UpdateGroupRequest: update group request.
type UpdateGroupRequest struct {
	GroupID string `json:"-"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`
}

// UpdateLimitRequest: update limit request.
type UpdateLimitRequest struct {
	LimitID string `json:"-"`

	// Precisely one of Limit, Unlimited, Disabled must be set.
	Limit *uint64 `json:"limit,omitempty"`

	// Precisely one of Limit, Unlimited, Disabled must be set.
	Unlimited *bool `json:"unlimited,omitempty"`

	// Precisely one of Limit, Unlimited, Disabled must be set.
	Disabled *bool `json:"disabled,omitempty"`
}

// UpdatePolicyRequest: update policy request.
type UpdatePolicyRequest struct {
	PolicyID string `json:"-"`

	Name *string `json:"name,omitempty"`

	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	UserID *string `json:"user_id,omitempty"`

	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	GroupID *string `json:"group_id,omitempty"`

	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	ApplicationID *string `json:"application_id,omitempty"`

	// Precisely one of UserID, GroupID, ApplicationID, NoPrincipal must be set.
	NoPrincipal *bool `json:"no_principal,omitempty"`

	Description *string `json:"description,omitempty"`
}

// UpdateQuotumRequest: update quotum request.
type UpdateQuotumRequest struct {
	QuotumID string `json:"-"`

	Name *string `json:"name,omitempty"`
}

// UpdateSSHKeyRequest: update ssh key request.
type UpdateSSHKeyRequest struct {
	SSHKeyID string `json:"-"`

	Name *string `json:"name,omitempty"`

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

// GetServiceInfo:
func (s *API) GetServiceInfo(opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-admin/v1",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "customer_level_names", req.CustomerLevelNames)

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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "customer_level_id", req.CustomerLevelID)

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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "customer_level_id", req.CustomerLevelID)
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
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

// UpdateSSHKey:
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "editable", req.Editable)
	parameter.AddToQuery(query, "visible", req.Visible)
	parameter.AddToQuery(query, "group_id", req.GroupID)
	parameter.AddToQuery(query, "user_id", req.UserID)
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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
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
	parameter.AddToQuery(query, "principal_id", req.PrincipalID)
	parameter.AddToQuery(query, "auth_id", req.AuthID)

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
	parameter.AddToQuery(query, "user_id", req.UserID)
	parameter.AddToQuery(query, "application_id", req.ApplicationID)

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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

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
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "customer_level_id", req.CustomerLevelID)

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
	parameter.AddToQuery(query, "created_after", req.CreatedAfter)
	parameter.AddToQuery(query, "created_before", req.CreatedBefore)
	parameter.AddToQuery(query, "ip", req.IP)
	parameter.AddToQuery(query, "country", req.Country)
	parameter.AddToQuery(query, "action", req.Action)
	parameter.AddToQuery(query, "from_principal_id", req.FromPrincipalID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "resource_type", req.ResourceType)
	parameter.AddToQuery(query, "resource_id", req.ResourceID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

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

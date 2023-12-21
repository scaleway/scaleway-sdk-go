// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package iam_private provides methods and message types of the iam_private v1 API.
package iam_private

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
	iam_v1alpha1 "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/iam/v1alpha1"
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

type ListPermissionSetCategoriesRequestOrderBy string

const (
	// Name ascending.
	ListPermissionSetCategoriesRequestOrderByNameAsc = ListPermissionSetCategoriesRequestOrderBy("name_asc")
	// Name descending.
	ListPermissionSetCategoriesRequestOrderByNameDesc = ListPermissionSetCategoriesRequestOrderBy("name_desc")
)

func (enum ListPermissionSetCategoriesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListPermissionSetCategoriesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPermissionSetCategoriesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPermissionSetCategoriesRequestOrderBy(ListPermissionSetCategoriesRequestOrderBy(tmp).String())
	return nil
}

type ListPrincipalPermissionSetsRequestOrderBy string

const (
	// Name ascending.
	ListPrincipalPermissionSetsRequestOrderByNameAsc = ListPrincipalPermissionSetsRequestOrderBy("name_asc")
	// Name descending.
	ListPrincipalPermissionSetsRequestOrderByNameDesc = ListPrincipalPermissionSetsRequestOrderBy("name_desc")
)

func (enum ListPrincipalPermissionSetsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListPrincipalPermissionSetsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPrincipalPermissionSetsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPrincipalPermissionSetsRequestOrderBy(ListPrincipalPermissionSetsRequestOrderBy(tmp).String())
	return nil
}

// PermissionSetCategory: permission set category.
type PermissionSetCategory struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

// PrincipalPermissionSet: principal permission set.
type PrincipalPermissionSet struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	PolicyID string `json:"policy_id"`

	GroupID *string `json:"group_id"`

	// ScopeType: default value: unknown_scope_type
	ScopeType iam_v1alpha1.PermissionSetScopeType `json:"scope_type"`

	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID must be set.
	ProjectIDs *[]string `json:"project_ids,omitempty"`

	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	// Precisely one of ProjectIDs, OrganizationID, AccountRootUserID must be set.
	AccountRootUserID *string `json:"account_root_user_id,omitempty"`
}

// ConsoleAPIListPermissionSetCategoriesRequest: console api list permission set categories request.
type ConsoleAPIListPermissionSetCategoriesRequest struct {
	// OrderBy: default value: name_asc
	OrderBy ListPermissionSetCategoriesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID string `json:"-"`
}

// ConsoleAPIListPrincipalPermissionSetsRequest: console api list principal permission sets request.
type ConsoleAPIListPrincipalPermissionSetsRequest struct {
	PrincipalID string `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListPrincipalPermissionSetsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	CategoryIDs []string `json:"-"`

	ProjectIDs []string `json:"-"`
}

// ConsoleAPISwitchJWTOrganizationRequest: console api switch jwt organization request.
type ConsoleAPISwitchJWTOrganizationRequest struct {
	Jti string `json:"-"`

	OrganizationID string `json:"organization_id"`
}

// ListPermissionSetCategoriesResponse: list permission set categories response.
type ListPermissionSetCategoriesResponse struct {
	PermissionSetCategories []*PermissionSetCategory `json:"permission_set_categories"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPermissionSetCategoriesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPermissionSetCategoriesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPermissionSetCategoriesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PermissionSetCategories = append(r.PermissionSetCategories, results.PermissionSetCategories...)
	r.TotalCount += uint64(len(results.PermissionSetCategories))
	return uint64(len(results.PermissionSetCategories)), nil
}

// ListPrincipalPermissionSetsResponse: list principal permission sets response.
type ListPrincipalPermissionSetsResponse struct {
	PrincipalPermissionSets []*PrincipalPermissionSet `json:"principal_permission_sets"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPrincipalPermissionSetsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPrincipalPermissionSetsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPrincipalPermissionSetsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PrincipalPermissionSets = append(r.PrincipalPermissionSets, results.PrincipalPermissionSets...)
	r.TotalCount += uint64(len(results.PrincipalPermissionSets))
	return uint64(len(results.PrincipalPermissionSets)), nil
}

// SwitchJWTOrganizationResponse: switch jwt organization response.
type SwitchJWTOrganizationResponse struct {
	Token string `json:"token"`

	RenewToken string `json:"renew_token"`

	// Deprecated
	Jti *string `json:"jti,omitempty"`

	// Deprecated
	UserID *string `json:"user_id,omitempty"`

	Jwt *iam_v1alpha1.JWT `json:"jwt"`
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

// SwitchJWTOrganization:
func (s *ConsoleAPI) SwitchJWTOrganization(req *ConsoleAPISwitchJWTOrganizationRequest, opts ...scw.RequestOption) (*SwitchJWTOrganizationResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.Jti) == "" {
		return nil, errors.New("field Jti cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iam-private/v1/jwts/" + fmt.Sprint(req.Jti) + "/switch-organization",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SwitchJWTOrganizationResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPermissionSetCategories:
func (s *ConsoleAPI) ListPermissionSetCategories(req *ConsoleAPIListPermissionSetCategoriesRequest, opts ...scw.RequestOption) (*ListPermissionSetCategoriesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-private/v1/permission-set-categories",
		Query:  query,
	}

	var resp ListPermissionSetCategoriesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPrincipalPermissionSets:
func (s *ConsoleAPI) ListPrincipalPermissionSets(req *ConsoleAPIListPrincipalPermissionSetsRequest, opts ...scw.RequestOption) (*ListPrincipalPermissionSetsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "category_ids", req.CategoryIDs)
	parameter.AddToQuery(query, "project_ids", req.ProjectIDs)

	if fmt.Sprint(req.PrincipalID) == "" {
		return nil, errors.New("field PrincipalID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iam-private/v1/principals/" + fmt.Sprint(req.PrincipalID) + "/permission-sets",
		Query:  query,
	}

	var resp ListPrincipalPermissionSetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

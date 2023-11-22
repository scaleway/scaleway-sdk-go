// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package observability provides methods and message types of the observability v1alpha1 API.
package observability

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

type CockpitStatus string

const (
	CockpitStatusUnknownStatus = CockpitStatus("unknown_status")
	CockpitStatusCreating      = CockpitStatus("creating")
	CockpitStatusReady         = CockpitStatus("ready")
	CockpitStatusDeleting      = CockpitStatus("deleting")
	CockpitStatusUpdating      = CockpitStatus("updating")
	CockpitStatusError         = CockpitStatus("error")
)

func (enum CockpitStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum CockpitStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CockpitStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CockpitStatus(CockpitStatus(tmp).String())
	return nil
}

type GrafanaUserRole string

const (
	GrafanaUserRoleUnknownRole = GrafanaUserRole("unknown_role")
	GrafanaUserRoleEditor      = GrafanaUserRole("editor")
	GrafanaUserRoleViewer      = GrafanaUserRole("viewer")
)

func (enum GrafanaUserRole) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_role"
	}
	return string(enum)
}

func (enum GrafanaUserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *GrafanaUserRole) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = GrafanaUserRole(GrafanaUserRole(tmp).String())
	return nil
}

type ListGrafanaUsersRequestOrderBy string

const (
	ListGrafanaUsersRequestOrderByLoginAsc  = ListGrafanaUsersRequestOrderBy("login_asc")
	ListGrafanaUsersRequestOrderByLoginDesc = ListGrafanaUsersRequestOrderBy("login_desc")
)

func (enum ListGrafanaUsersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "login_asc"
	}
	return string(enum)
}

func (enum ListGrafanaUsersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListGrafanaUsersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListGrafanaUsersRequestOrderBy(ListGrafanaUsersRequestOrderBy(tmp).String())
	return nil
}

type ListTokensRequestOrderBy string

const (
	ListTokensRequestOrderByCreatedAtAsc  = ListTokensRequestOrderBy("created_at_asc")
	ListTokensRequestOrderByCreatedAtDesc = ListTokensRequestOrderBy("created_at_desc")
	ListTokensRequestOrderByNameAsc       = ListTokensRequestOrderBy("name_asc")
	ListTokensRequestOrderByNameDesc      = ListTokensRequestOrderBy("name_desc")
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

// ContactPointEmail: contact point email.
type ContactPointEmail struct {
	To string `json:"to"`
}

// TokenScopes: token scopes.
type TokenScopes struct {
	QueryMetrics bool `json:"query_metrics"`

	WriteMetrics bool `json:"write_metrics"`

	SetupMetricsRules bool `json:"setup_metrics_rules"`

	QueryLogs bool `json:"query_logs"`

	WriteLogs bool `json:"write_logs"`

	SetupLogsRules bool `json:"setup_logs_rules"`

	SetupAlerts bool `json:"setup_alerts"`
}

// CockpitEndpoints: cockpit endpoints.
type CockpitEndpoints struct {
	MetricsURL string `json:"metrics_url"`

	LogsURL string `json:"logs_url"`

	AlertmanagerURL string `json:"alertmanager_url"`

	GrafanaURL string `json:"grafana_url"`
}

// ContactPoint: Alert contact point.
type ContactPoint struct {
	// Email: alert contact point configuration.
	// Precisely one of Email must be set.
	Email *ContactPointEmail `json:"email,omitempty"`
}

// GrafanaUser: Grafana user.
type GrafanaUser struct {
	ID uint32 `json:"id"`

	Login string `json:"login"`

	// Role: default value: unknown_role
	Role GrafanaUserRole `json:"role"`

	Password *string `json:"password"`
}

// Token: token.
type Token struct {
	ID string `json:"id"`

	ProjectID string `json:"project_id"`

	Name string `json:"name"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	Scopes *TokenScopes `json:"scopes"`

	SecretKey *string `json:"secret_key"`
}

// ActivateCockpitRequest: activate cockpit request.
type ActivateCockpitRequest struct {
	ProjectID string `json:"project_id"`
}

// Cockpit: Cockpit.
type Cockpit struct {
	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// CreatedAt: created at.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: updated at.
	UpdatedAt *time.Time `json:"updated_at"`

	// Endpoints: endpoints.
	Endpoints *CockpitEndpoints `json:"endpoints"`

	// Status: status.
	// Default value: unknown_status
	Status CockpitStatus `json:"status"`

	// ManagedAlertsEnabled: managed alerts enabled.
	ManagedAlertsEnabled bool `json:"managed_alerts_enabled"`
}

// CreateContactPointRequest: Create contact point request.
type CreateContactPointRequest struct {
	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// ContactPoint: contact point to create.
	ContactPoint *ContactPoint `json:"contact_point,omitempty"`
}

// CreateGrafanaUserRequest: Create grafana user request.
type CreateGrafanaUserRequest struct {
	ProjectID string `json:"project_id"`

	Login string `json:"login"`

	// Role: default value: unknown_role
	Role GrafanaUserRole `json:"role"`
}

// CreateTokenRequest: create token request.
type CreateTokenRequest struct {
	ProjectID string `json:"project_id"`

	Name string `json:"name"`

	Scopes *TokenScopes `json:"scopes,omitempty"`
}

// DeactivateCockpitRequest: deactivate cockpit request.
type DeactivateCockpitRequest struct {
	ProjectID string `json:"project_id"`
}

// DeleteContactPointRequest: Delete contact point request.
type DeleteContactPointRequest struct {
	ProjectID string `json:"project_id"`

	// ContactPoint: contact point to delete.
	ContactPoint *ContactPoint `json:"contact_point,omitempty"`
}

// DeleteGrafanaUserRequest: Delete grafana user request.
type DeleteGrafanaUserRequest struct {
	GrafanaUserID uint32 `json:"-"`

	ProjectID string `json:"project_id"`
}

// DeleteTokenRequest: delete token request.
type DeleteTokenRequest struct {
	TokenID string `json:"-"`
}

// DisableManagedAlertsRequest: Disable managed alerts request.
type DisableManagedAlertsRequest struct {
	ProjectID string `json:"project_id"`
}

// EnableManagedAlertsRequest: Enable managed alerts request.
type EnableManagedAlertsRequest struct {
	ProjectID string `json:"project_id"`
}

// GetCockpitRequest: get cockpit request.
type GetCockpitRequest struct {
	ProjectID string `json:"-"`
}

// GetTokenRequest: get token request.
type GetTokenRequest struct {
	TokenID string `json:"-"`
}

// ListContactPointsRequest: List contact points request.
type ListContactPointsRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// ProjectID: project ID.
	ProjectID string `json:"-"`
}

// ListContactPointsResponse: List contact points response.
type ListContactPointsResponse struct {
	// TotalCount: total count of contact points.
	TotalCount uint32 `json:"total_count"`

	// ContactPoints: contact points array.
	ContactPoints []*ContactPoint `json:"contact_points"`

	// HasAdditionalReceivers: has receivers other than default.
	HasAdditionalReceivers bool `json:"has_additional_receivers"`

	// HasAdditionalContactPoints: has unmanaged contact points.
	HasAdditionalContactPoints bool `json:"has_additional_contact_points"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListContactPointsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListContactPointsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListContactPointsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ContactPoints = append(r.ContactPoints, results.ContactPoints...)
	r.TotalCount += uint32(len(results.ContactPoints))
	return uint32(len(results.ContactPoints)), nil
}

// ListGrafanaUsersRequest: List grafana users request.
type ListGrafanaUsersRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: login_asc
	OrderBy ListGrafanaUsersRequestOrderBy `json:"-"`

	ProjectID string `json:"-"`
}

// ListGrafanaUsersResponse: List grafana users response.
type ListGrafanaUsersResponse struct {
	TotalCount uint32 `json:"total_count"`

	GrafanaUsers []*GrafanaUser `json:"grafana_users"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListGrafanaUsersResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListGrafanaUsersResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListGrafanaUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.GrafanaUsers = append(r.GrafanaUsers, results.GrafanaUsers...)
	r.TotalCount += uint32(len(results.GrafanaUsers))
	return uint32(len(results.GrafanaUsers)), nil
}

// ListTokensRequest: list tokens request.
type ListTokensRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListTokensRequestOrderBy `json:"-"`

	ProjectID string `json:"-"`
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

// ResetCockpitGrafanaRequest: reset cockpit grafana request.
type ResetCockpitGrafanaRequest struct {
	ProjectID string `json:"project_id"`
}

// ResetGrafanaUserPasswordRequest: Reset grafana user password request.
type ResetGrafanaUserPasswordRequest struct {
	GrafanaUserID uint32 `json:"-"`

	ProjectID string `json:"project_id"`
}

// TriggerTestAlertRequest: trigger test alert request.
type TriggerTestAlertRequest struct {
	ProjectID string `json:"project_id"`
}

// This API allows to manage Cockpits.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ActivateCockpit:
func (s *API) ActivateCockpit(req *ActivateCockpitRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/observability/v1alpha1/activate-cockpit",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cockpit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCockpit: Get the cockpit associated with the given project ID.
func (s *API) GetCockpit(req *GetCockpitRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/observability/v1alpha1/cockpit",
		Query:  query,
	}

	var resp Cockpit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeactivateCockpit:
func (s *API) DeactivateCockpit(req *DeactivateCockpitRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/observability/v1alpha1/deactivate-cockpit",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cockpit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ResetCockpitGrafana: Reset the Grafana your cockpit associated with the given project ID.
func (s *API) ResetCockpitGrafana(req *ResetCockpitGrafanaRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/observability/v1alpha1/reset-cockpit-grafana",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Cockpit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateToken: Create token associated with the given project ID.
func (s *API) CreateToken(req *CreateTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/observability/v1alpha1/tokens",
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

// ListTokens: List tokens associated with the given project ID.
func (s *API) ListTokens(req *ListTokensRequest, opts ...scw.RequestOption) (*ListTokensResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/observability/v1alpha1/tokens",
		Query:  query,
	}

	var resp ListTokensResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetToken: Get the token associated with the given ID.
func (s *API) GetToken(req *GetTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if fmt.Sprint(req.TokenID) == "" {
		return nil, errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/observability/v1alpha1/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteToken: Delete the token associated with the given ID.
func (s *API) DeleteToken(req *DeleteTokenRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.TokenID) == "" {
		return errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/observability/v1alpha1/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateContactPoint: Create an alert contact point for the default receiver.
func (s *API) CreateContactPoint(req *CreateContactPointRequest, opts ...scw.RequestOption) (*ContactPoint, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/observability/v1alpha1/contact-points",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ContactPoint

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListContactPoints: List alert contact points associated with the given cockpit ID.
func (s *API) ListContactPoints(req *ListContactPointsRequest, opts ...scw.RequestOption) (*ListContactPointsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/observability/v1alpha1/contact-points",
		Query:  query,
	}

	var resp ListContactPointsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteContactPoint: Delete an alert contact point for the default receiver.
func (s *API) DeleteContactPoint(req *DeleteContactPointRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/observability/v1alpha1/delete-contact-point",
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

// EnableManagedAlerts: Enable managed alerts.
func (s *API) EnableManagedAlerts(req *EnableManagedAlertsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/observability/v1alpha1/enable-managed-alerts",
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

// DisableManagedAlerts: Disable managed alerts.
func (s *API) DisableManagedAlerts(req *DisableManagedAlertsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/observability/v1alpha1/disable-managed-alerts",
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

// TriggerTestAlert: Trigger a test alert to all receivers.
func (s *API) TriggerTestAlert(req *TriggerTestAlertRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/observability/v1alpha1/trigger-test-alert",
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

// CreateGrafanaUser: Create a grafana user for your grafana instance.
func (s *API) CreateGrafanaUser(req *CreateGrafanaUserRequest, opts ...scw.RequestOption) (*GrafanaUser, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/observability/v1alpha1/grafana-users",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp GrafanaUser

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListGrafanaUsers: List grafana users who are able to connect to your grafana instance.
func (s *API) ListGrafanaUsers(req *ListGrafanaUsersRequest, opts ...scw.RequestOption) (*ListGrafanaUsersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/observability/v1alpha1/grafana-users",
		Query:  query,
	}

	var resp ListGrafanaUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteGrafanaUser: Delete a grafana user from your grafana instance.
func (s *API) DeleteGrafanaUser(req *DeleteGrafanaUserRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.GrafanaUserID) == "" {
		return errors.New("field GrafanaUserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/observability/v1alpha1/grafana-users/" + fmt.Sprint(req.GrafanaUserID) + "/delete",
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

// ResetGrafanaUserPassword: Reset the Grafana user password from your grafana instance.
func (s *API) ResetGrafanaUserPassword(req *ResetGrafanaUserPasswordRequest, opts ...scw.RequestOption) (*GrafanaUser, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.GrafanaUserID) == "" {
		return nil, errors.New("field GrafanaUserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/observability/v1alpha1/grafana-users/" + fmt.Sprint(req.GrafanaUserID) + "/reset-password",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp GrafanaUser

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

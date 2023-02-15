// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package cockpit provides methods and message types of the cockpit v1beta1 API.
package cockpit

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

// API: this API allows to manage Cockpits
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

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

// Cockpit: cockpit
type Cockpit struct {
	// ProjectID: project ID
	ProjectID string `json:"project_id"`
	// CreatedAt: created at
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: updated at
	UpdatedAt *time.Time `json:"updated_at"`
	// Endpoints: endpoints
	Endpoints *CockpitEndpoints `json:"endpoints"`
	// Status: status
	//
	// Default value: unknown_status
	Status CockpitStatus `json:"status"`
	// ManagedAlertsEnabled: managed alerts enabled
	ManagedAlertsEnabled bool `json:"managed_alerts_enabled"`
}

// CockpitEndpoints: cockpit. endpoints
type CockpitEndpoints struct {
	MetricsURL string `json:"metrics_url"`

	LogsURL string `json:"logs_url"`

	AlertmanagerURL string `json:"alertmanager_url"`

	GrafanaURL string `json:"grafana_url"`
}

// ContactPoint: alert contact point
type ContactPoint struct {
	// Email: alert contact point configuration
	// Precisely one of Email must be set.
	Email *ContactPointEmail `json:"email,omitempty"`
}

type ContactPointEmail struct {
	To string `json:"to"`
}

// GrafanaUser: grafana user
type GrafanaUser struct {
	ID uint32 `json:"id"`

	Login string `json:"login"`
	// Role:
	//
	// Default value: unknown_role
	Role GrafanaUserRole `json:"role"`

	Password *string `json:"password"`
}

// ListContactPointsResponse: list contact points response
type ListContactPointsResponse struct {
	// TotalCount: total count of contact points
	TotalCount uint32 `json:"total_count"`
	// ContactPoints: contact points array
	ContactPoints []*ContactPoint `json:"contact_points"`
	// HasAdditionalReceivers: has receivers other than default
	HasAdditionalReceivers bool `json:"has_additional_receivers"`
	// HasAdditionalContactPoints: has unmanaged contact points
	HasAdditionalContactPoints bool `json:"has_additional_contact_points"`
}

// ListGrafanaUsersResponse: list grafana users response
type ListGrafanaUsersResponse struct {
	TotalCount uint32 `json:"total_count"`

	GrafanaUsers []*GrafanaUser `json:"grafana_users"`
}

// ListTokensResponse: list tokens response
type ListTokensResponse struct {
	TotalCount uint32 `json:"total_count"`

	Tokens []*Token `json:"tokens"`
}

// Token: token
type Token struct {
	ID string `json:"id"`

	ProjectID string `json:"project_id"`

	Name string `json:"name"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	Scopes *TokenScopes `json:"scopes"`

	SecretKey *string `json:"secret_key"`
}

// TokenScopes: token scopes
type TokenScopes struct {
	QueryMetrics bool `json:"query_metrics"`

	WriteMetrics bool `json:"write_metrics"`

	SetupMetricsRules bool `json:"setup_metrics_rules"`

	QueryLogs bool `json:"query_logs"`

	WriteLogs bool `json:"write_logs"`

	SetupLogsRules bool `json:"setup_logs_rules"`

	SetupAlerts bool `json:"setup_alerts"`
}

// Service API

type ActivateCockpitRequest struct {
	ProjectID string `json:"project_id"`
}

func (s *API) ActivateCockpit(req *ActivateCockpitRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/cockpit/v1beta1/activate",
		Headers: http.Header{},
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

type GetCockpitRequest struct {
	ProjectID string `json:"-"`
}

// GetCockpit: get cockpit
//
// Get the cockpit associated with the given project ID.
func (s *API) GetCockpit(req *GetCockpitRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/cockpit/v1beta1/cockpit",
		Query:   query,
		Headers: http.Header{},
	}

	var resp Cockpit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeactivateCockpitRequest struct {
	ProjectID string `json:"project_id"`
}

func (s *API) DeactivateCockpit(req *DeactivateCockpitRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/cockpit/v1beta1/deactivate",
		Headers: http.Header{},
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

type ResetCockpitGrafanaRequest struct {
	ProjectID string `json:"project_id"`
}

// ResetCockpitGrafana: reset Grafana
//
// Reset the Grafana your cockpit associated with the given project ID.
func (s *API) ResetCockpitGrafana(req *ResetCockpitGrafanaRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/cockpit/v1beta1/reset-grafana",
		Headers: http.Header{},
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

type CreateTokenRequest struct {
	ProjectID string `json:"project_id"`

	Name string `json:"name"`

	Scopes *TokenScopes `json:"scopes"`
}

// CreateToken: create a token
//
// Create token associated with the given project ID.
func (s *API) CreateToken(req *CreateTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/cockpit/v1beta1/tokens",
		Headers: http.Header{},
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

type ListTokensRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: created_at_asc
	OrderBy ListTokensRequestOrderBy `json:"-"`

	ProjectID string `json:"-"`
}

// ListTokens: list tokens
//
// List tokens associated with the given project ID.
func (s *API) ListTokens(req *ListTokensRequest, opts ...scw.RequestOption) (*ListTokensResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/cockpit/v1beta1/tokens",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListTokensResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetTokenRequest struct {
	TokenID string `json:"-"`
}

// GetToken: get token
//
// Get the token associated with the given ID.
func (s *API) GetToken(req *GetTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if fmt.Sprint(req.TokenID) == "" {
		return nil, errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/cockpit/v1beta1/tokens/" + fmt.Sprint(req.TokenID) + "",
		Headers: http.Header{},
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteTokenRequest struct {
	TokenID string `json:"-"`
}

// DeleteToken: delete token
//
// Delete the token associated with the given ID.
func (s *API) DeleteToken(req *DeleteTokenRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.TokenID) == "" {
		return errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/cockpit/v1beta1/tokens/" + fmt.Sprint(req.TokenID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type CreateContactPointRequest struct {
	// ProjectID: project ID
	ProjectID string `json:"project_id"`
	// ContactPoint: contact point to create
	ContactPoint *ContactPoint `json:"contact_point"`
}

// CreateContactPoint: create an alert contact point
//
// Create an alert contact point for the default receiver.
func (s *API) CreateContactPoint(req *CreateContactPointRequest, opts ...scw.RequestOption) (*ContactPoint, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/cockpit/v1beta1/contact-points",
		Headers: http.Header{},
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

type ListContactPointsRequest struct {
	// Page: page number
	Page *int32 `json:"-"`
	// PageSize: page size
	PageSize *uint32 `json:"-"`
	// ProjectID: project ID
	ProjectID string `json:"-"`
}

// ListContactPoints: list alert contact points
//
// List alert contact points associated with the given cockpit ID.
func (s *API) ListContactPoints(req *ListContactPointsRequest, opts ...scw.RequestOption) (*ListContactPointsResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/cockpit/v1beta1/contact-points",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListContactPointsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteContactPointRequest struct {
	ProjectID string `json:"project_id"`
	// ContactPoint: contact point to delete
	ContactPoint *ContactPoint `json:"contact_point"`
}

// DeleteContactPoint: delete an alert contact point
//
// Delete an alert contact point for the default receiver.
func (s *API) DeleteContactPoint(req *DeleteContactPointRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/cockpit/v1beta1/delete-contact-point",
		Headers: http.Header{},
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

type EnableManagedAlertsRequest struct {
	ProjectID string `json:"project_id"`
}

// EnableManagedAlerts: enable managed alerts
//
// Enable managed alerts.
func (s *API) EnableManagedAlerts(req *EnableManagedAlertsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/cockpit/v1beta1/enable-managed-alerts",
		Headers: http.Header{},
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

type DisableManagedAlertsRequest struct {
	ProjectID string `json:"project_id"`
}

// DisableManagedAlerts: disable managed alerts
//
// Disable managed alerts.
func (s *API) DisableManagedAlerts(req *DisableManagedAlertsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/cockpit/v1beta1/disable-managed-alerts",
		Headers: http.Header{},
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

type TriggerTestAlertRequest struct {
	ProjectID string `json:"project_id"`
}

// TriggerTestAlert: trigger a test alert
//
// Trigger a test alert to all receivers.
func (s *API) TriggerTestAlert(req *TriggerTestAlertRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/cockpit/v1beta1/trigger-test-alert",
		Headers: http.Header{},
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

type CreateGrafanaUserRequest struct {
	ProjectID string `json:"project_id"`

	Login string `json:"login"`
	// Role:
	//
	// Default value: unknown_role
	Role GrafanaUserRole `json:"role"`
}

// CreateGrafanaUser: create a grafana user
//
// Create a grafana user for your grafana instance.
func (s *API) CreateGrafanaUser(req *CreateGrafanaUserRequest, opts ...scw.RequestOption) (*GrafanaUser, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/cockpit/v1beta1/grafana-users",
		Headers: http.Header{},
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

type ListGrafanaUsersRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
	// OrderBy:
	//
	// Default value: login_asc
	OrderBy ListGrafanaUsersRequestOrderBy `json:"-"`

	ProjectID string `json:"-"`
}

// ListGrafanaUsers: list grafana users
//
// List grafana users who are able to connect to your grafana instance.
func (s *API) ListGrafanaUsers(req *ListGrafanaUsersRequest, opts ...scw.RequestOption) (*ListGrafanaUsersResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/cockpit/v1beta1/grafana-users",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListGrafanaUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteGrafanaUserRequest struct {
	GrafanaUserID uint32 `json:"-"`

	ProjectID string `json:"project_id"`
}

// DeleteGrafanaUser: delete a grafana user
//
// Delete a grafana user from your grafana instance.
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
		Method:  "POST",
		Path:    "/cockpit/v1beta1/grafana-users/" + fmt.Sprint(req.GrafanaUserID) + "/delete",
		Headers: http.Header{},
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

type ResetGrafanaUserPasswordRequest struct {
	GrafanaUserID uint32 `json:"-"`

	ProjectID string `json:"project_id"`
}

// ResetGrafanaUserPassword: reset Grafana user password
//
// Reset the Grafana user password from your grafana instance.
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
		Method:  "POST",
		Path:    "/cockpit/v1beta1/grafana-users/" + fmt.Sprint(req.GrafanaUserID) + "/reset-password",
		Headers: http.Header{},
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

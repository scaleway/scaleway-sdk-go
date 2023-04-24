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

// API: cockpit API.
// Cockpit's API allows you to activate your Cockpit on your Projects. Scaleway's Cockpit stores metrics and logs and provides a dedicated Grafana for dashboarding to visualize them.
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

type ListPlansRequestOrderBy string

const (
	ListPlansRequestOrderByNameAsc  = ListPlansRequestOrderBy("name_asc")
	ListPlansRequestOrderByNameDesc = ListPlansRequestOrderBy("name_desc")
)

func (enum ListPlansRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListPlansRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPlansRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPlansRequestOrderBy(ListPlansRequestOrderBy(tmp).String())
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

type PlanName string

const (
	PlanNameUnknownName = PlanName("unknown_name")
	PlanNameFree        = PlanName("free")
	PlanNamePremium     = PlanName("premium")
	PlanNameCustom      = PlanName("custom")
)

func (enum PlanName) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_name"
	}
	return string(enum)
}

func (enum PlanName) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlanName) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlanName(PlanName(tmp).String())
	return nil
}

// Cockpit: cockpit.
type Cockpit struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`
	// CreatedAt: date and time of the Cockpit's creation.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: date and time of the Cockpit's last update.
	UpdatedAt *time.Time `json:"updated_at"`
	// Endpoints: endpoints of the Cockpit.
	Endpoints *CockpitEndpoints `json:"endpoints"`
	// Status: status of the Cockpit.
	// Default value: unknown_status
	Status CockpitStatus `json:"status"`
	// ManagedAlertsEnabled: specifies whether managed alerts are enabled or disabled.
	ManagedAlertsEnabled bool `json:"managed_alerts_enabled"`
	// Plan: pricing plan information.
	Plan *Plan `json:"plan"`
}

// CockpitEndpoints: cockpit. endpoints.
type CockpitEndpoints struct {
	// MetricsURL: URL for metrics.
	MetricsURL string `json:"metrics_url"`
	// LogsURL: URL for logs.
	LogsURL string `json:"logs_url"`
	// AlertmanagerURL: URL for the alert manager.
	AlertmanagerURL string `json:"alertmanager_url"`
	// GrafanaURL: URL for the Grafana dashboard.
	GrafanaURL string `json:"grafana_url"`
}

// CockpitMetrics: metrics for a given Cockpit.
// Cockpit metrics.
type CockpitMetrics struct {
	// Timeseries: time series array.
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

// ContactPoint: contact point.
type ContactPoint struct {
	// Email: contact point configuration.
	// Precisely one of Email must be set.
	Email *ContactPointEmail `json:"email,omitempty"`
}

type ContactPointEmail struct {
	To string `json:"to"`
}

// GrafanaUser: grafana user.
type GrafanaUser struct {
	// ID: ID of the Grafana user.
	ID uint32 `json:"id"`
	// Login: username of the Grafana user.
	Login string `json:"login"`
	// Role: role assigned to the Grafana user.
	// Default value: unknown_role
	Role GrafanaUserRole `json:"role"`
	// Password: the Grafana user's password.
	Password *string `json:"password"`
}

// ListContactPointsResponse: response returned when listing contact points.
// List contact points response.
type ListContactPointsResponse struct {
	// TotalCount: count of all contact points created.
	TotalCount uint32 `json:"total_count"`
	// ContactPoints: array of contact points.
	ContactPoints []*ContactPoint `json:"contact_points"`
	// HasAdditionalReceivers: specifies whether the contact point has other receivers than the default receiver.
	HasAdditionalReceivers bool `json:"has_additional_receivers"`
	// HasAdditionalContactPoints: specifies whether there are unmanaged contact points.
	HasAdditionalContactPoints bool `json:"has_additional_contact_points"`
}

// ListGrafanaUsersResponse: response returned when listing Grafana users.
// List grafana users response.
type ListGrafanaUsersResponse struct {
	// TotalCount: count of all Grafana users.
	TotalCount uint32 `json:"total_count"`
	// GrafanaUsers: information on all Grafana users.
	GrafanaUsers []*GrafanaUser `json:"grafana_users"`
}

// ListPlansResponse: response returned when listing all pricing plans.
// List plans response.
type ListPlansResponse struct {
	// TotalCount: count of all pricing plans.
	TotalCount uint64 `json:"total_count"`
	// Plans: information on plans.
	Plans []*Plan `json:"plans"`
}

// ListTokensResponse: list tokens response.
type ListTokensResponse struct {
	// TotalCount: count of all tokens created.
	TotalCount uint32 `json:"total_count"`
	// Tokens: list of all tokens created.
	Tokens []*Token `json:"tokens"`
}

// Plan: pricing plan.
// Plan.
type Plan struct {
	// ID: ID of a given pricing plan.
	ID string `json:"id"`
	// Name: name of a given pricing plan.
	// Default value: unknown_name
	Name PlanName `json:"name"`
	// RetentionMetricsInterval: retention for metrics.
	RetentionMetricsInterval *scw.Duration `json:"retention_metrics_interval"`
	// RetentionLogsInterval: retention for logs.
	RetentionLogsInterval *scw.Duration `json:"retention_logs_interval"`
	// SampleIngestionPrice: ingestion price for 1 million samples in cents.
	SampleIngestionPrice uint32 `json:"sample_ingestion_price"`
	// LogsIngestionPrice: ingestion price for 1 GB of logs in cents.
	LogsIngestionPrice uint32 `json:"logs_ingestion_price"`
	// RetentionPrice: retention price in euros per month.
	RetentionPrice uint32 `json:"retention_price"`
}

// SelectPlanResponse: response returned when selecting a pricing plan.
// Select plan response.
type SelectPlanResponse struct {
}

// Token: token.
type Token struct {
	// ID: ID of the token.
	ID string `json:"id"`
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
	// Name: name of the token.
	Name string `json:"name"`
	// CreatedAt: date and time of the token's creation.
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: date and time of the token's last update.
	UpdatedAt *time.Time `json:"updated_at"`
	// Scopes: token's permissions.
	Scopes *TokenScopes `json:"scopes"`
	// SecretKey: token's secret key.
	SecretKey *string `json:"secret_key"`
}

// TokenScopes: token scopes.
type TokenScopes struct {
	// QueryMetrics: permission to fetch metrics.
	QueryMetrics bool `json:"query_metrics"`
	// WriteMetrics: permission to write metrics.
	WriteMetrics bool `json:"write_metrics"`
	// SetupMetricsRules: permission to setup metrics rules.
	SetupMetricsRules bool `json:"setup_metrics_rules"`
	// QueryLogs: permission to fetch logs.
	QueryLogs bool `json:"query_logs"`
	// WriteLogs: permission to write logs.
	WriteLogs bool `json:"write_logs"`
	// SetupLogsRules: permission to setup logs rules.
	SetupLogsRules bool `json:"setup_logs_rules"`
	// SetupAlerts: permission to setup alerts.
	SetupAlerts bool `json:"setup_alerts"`
}

// Service API

type ActivateCockpitRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`
}

// ActivateCockpit: activate a Cockpit.
// Activate the Cockpit of the specified Project ID.
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
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"-"`
}

// GetCockpit: get a Cockpit.
// Retrieve the Cockpit of the specified Project ID.
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

type GetCockpitMetricsRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"-"`
	// StartDate: desired time range's start date for the metrics.
	StartDate *time.Time `json:"-"`
	// EndDate: desired time range's end date for the metrics.
	EndDate *time.Time `json:"-"`
	// MetricName: name of the metric requested.
	MetricName *string `json:"-"`
}

// GetCockpitMetrics: get Cockpit metrics.
// Get metrics from your Cockpit with the specified Project ID.
func (s *API) GetCockpitMetrics(req *GetCockpitMetricsRequest, opts ...scw.RequestOption) (*CockpitMetrics, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)
	parameter.AddToQuery(query, "metric_name", req.MetricName)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/cockpit/v1beta1/cockpit/metrics",
		Query:   query,
		Headers: http.Header{},
	}

	var resp CockpitMetrics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeactivateCockpitRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`
}

// DeactivateCockpit: deactivate a Cockpit.
// Deactivate the Cockpit of the specified Project ID.
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
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`
}

// ResetCockpitGrafana: reset a Grafana.
// Reset your Cockpit's Grafana associated with the specified Project ID.
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
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
	// Name: name of the token.
	Name string `json:"name"`
	// Scopes: token's permissions.
	Scopes *TokenScopes `json:"scopes"`
}

// CreateToken: create a token.
// Create a token associated with the specified Project ID.
func (s *API) CreateToken(req *CreateTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("token")
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
	// Page: page number.
	Page *int32 `json:"-"`
	// PageSize: page size.
	PageSize *uint32 `json:"-"`
	// OrderBy: default value: created_at_asc
	OrderBy ListTokensRequestOrderBy `json:"-"`
	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// ListTokens: list tokens.
// Get a list of tokens associated with the specified Project ID.
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
	// TokenID: ID of the token.
	TokenID string `json:"-"`
}

// GetToken: get a token.
// Retrieve the token associated with the specified token ID.
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
	// TokenID: ID of the token.
	TokenID string `json:"-"`
}

// DeleteToken: delete a token.
// Delete the token associated with the specified token ID.
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
	// ProjectID: ID of the Project in which to create the contact point.
	ProjectID string `json:"project_id"`
	// ContactPoint: contact point to create.
	ContactPoint *ContactPoint `json:"contact_point"`
}

// CreateContactPoint: create a contact point.
// Create a contact point to receive alerts for the default receiver.
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
	// Page: page number.
	Page *int32 `json:"-"`
	// PageSize: page size.
	PageSize *uint32 `json:"-"`
	// ProjectID: ID of the Project from which to list the contact points.
	ProjectID string `json:"-"`
}

// ListContactPoints: list contact points.
// Get a list of contact points for the Cockpit associated with the specified Project ID.
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
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
	// ContactPoint: contact point to delete.
	ContactPoint *ContactPoint `json:"contact_point"`
}

// DeleteContactPoint: delete an alert contact point.
// Delete a contact point for the default receiver.
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
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// EnableManagedAlerts: enable managed alerts.
// Enable the sending of managed alerts for the specified Project's Cockpit.
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
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// DisableManagedAlerts: disable managed alerts.
// Disable the sending of managed alerts for the specified Project's Cockpit.
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

// TriggerTestAlert: trigger a test alert.
// Trigger a test alert to all of the Cockpit's receivers.
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
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
	// Login: username of the Grafana user.
	Login string `json:"login"`
	// Role: role assigned to the Grafana user.
	// Default value: unknown_role
	Role GrafanaUserRole `json:"role"`
}

// CreateGrafanaUser: create a Grafana user.
// Create a Grafana user for your Cockpit's Grafana instance. Make sure you save the automatically-generated password and the Grafana user ID.
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
	// Page: page number.
	Page *int32 `json:"-"`
	// PageSize: page size.
	PageSize *uint32 `json:"-"`
	// OrderBy: default value: login_asc
	OrderBy ListGrafanaUsersRequestOrderBy `json:"-"`
	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// ListGrafanaUsers: list Grafana users.
// Get a list of Grafana users who are able to connect to the Cockpit's Grafana instance.
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
	// GrafanaUserID: ID of the Grafana user.
	GrafanaUserID uint32 `json:"-"`
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// DeleteGrafanaUser: delete a Grafana user.
// Delete a Grafana user from a Grafana instance, specified by the Cockpit's Project ID and the Grafana user ID.
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
	// GrafanaUserID: ID of the Grafana user.
	GrafanaUserID uint32 `json:"-"`
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// ResetGrafanaUserPassword: reset a Grafana user's password.
// Reset a Grafana user's password specified by the Cockpit's Project ID and the Grafana user ID.
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

type ListPlansRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`
	// PageSize: page size.
	PageSize *uint32 `json:"-"`
	// OrderBy: default value: name_asc
	OrderBy ListPlansRequestOrderBy `json:"-"`
}

// ListPlans: list pricing plans.
// Get a list of all pricing plans available.
func (s *API) ListPlans(req *ListPlansRequest, opts ...scw.RequestOption) (*ListPlansResponse, error) {
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
		Method:  "GET",
		Path:    "/cockpit/v1beta1/plans",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListPlansResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SelectPlanRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
	// PlanID: ID of the pricing plan.
	PlanID string `json:"plan_id"`
}

// SelectPlan: select pricing plan.
// Select your chosen pricing plan for your Cockpit, specifying the Cockpit's Project ID and the pricing plan's ID in the request.
func (s *API) SelectPlan(req *SelectPlanRequest, opts ...scw.RequestOption) (*SelectPlanResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/cockpit/v1beta1/select-plan",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SelectPlanResponse

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

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPlansResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPlansResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPlansResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Plans = append(r.Plans, results.Plans...)
	r.TotalCount += uint64(len(results.Plans))
	return uint64(len(results.Plans)), nil
}

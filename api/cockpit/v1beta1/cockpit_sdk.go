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

type DatasourceType string

const (
	DatasourceTypeUnknownDatasourceType = DatasourceType("unknown_datasource_type")
	DatasourceTypeMetrics               = DatasourceType("metrics")
	DatasourceTypeLogs                  = DatasourceType("logs")
	DatasourceTypeTraces                = DatasourceType("traces")
	DatasourceTypeAlerts                = DatasourceType("alerts")
)

func (enum DatasourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_datasource_type"
	}
	return string(enum)
}

func (enum DatasourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DatasourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DatasourceType(DatasourceType(tmp).String())
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

type ListDatasourcesRequestOrderBy string

const (
	ListDatasourcesRequestOrderByCreatedAtAsc  = ListDatasourcesRequestOrderBy("created_at_asc")
	ListDatasourcesRequestOrderByCreatedAtDesc = ListDatasourcesRequestOrderBy("created_at_desc")
	ListDatasourcesRequestOrderByNameAsc       = ListDatasourcesRequestOrderBy("name_asc")
	ListDatasourcesRequestOrderByNameDesc      = ListDatasourcesRequestOrderBy("name_desc")
)

func (enum ListDatasourcesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListDatasourcesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDatasourcesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDatasourcesRequestOrderBy(ListDatasourcesRequestOrderBy(tmp).String())
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
	// The free plan is the default plan.
	PlanNameFree = PlanName("free")
	// The premium plan with a longer retention.
	PlanNamePremium = PlanName("premium")
	// The custom plan.
	PlanNameCustom = PlanName("custom")
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

// ContactPointEmail: contact point email.
type ContactPointEmail struct {
	To string `json:"to"`
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

	// SetupLogsRules: permission to set up logs rules.
	SetupLogsRules bool `json:"setup_logs_rules"`

	// SetupAlerts: permission to set up alerts.
	SetupAlerts bool `json:"setup_alerts"`

	// QueryTraces: permission to fetch traces.
	QueryTraces bool `json:"query_traces"`

	// WriteTraces: permission to write traces.
	WriteTraces bool `json:"write_traces"`
}

// CockpitEndpoints: cockpit endpoints.
type CockpitEndpoints struct {
	// MetricsURL: URL for metrics.
	MetricsURL string `json:"metrics_url"`

	// LogsURL: URL for logs.
	LogsURL string `json:"logs_url"`

	// TracesURL: URL for traces.
	TracesURL string `json:"traces_url"`

	// AlertmanagerURL: URL for the alert manager.
	AlertmanagerURL string `json:"alertmanager_url"`

	// GrafanaURL: URL for the Grafana dashboard.
	GrafanaURL string `json:"grafana_url"`
}

// Plan: Pricing plan.
type Plan struct {
	// ID: ID of a given pricing plan.
	ID string `json:"id"`

	// Name: name of a given pricing plan.
	// Default value: unknown_name
	Name PlanName `json:"name"`

	// RetentionMetricsInterval: interval of time during which Scaleway's Cockpit keeps your metrics.
	RetentionMetricsInterval *scw.Duration `json:"retention_metrics_interval"`

	// RetentionLogsInterval: interval of time during which Scaleway's Cockpit keeps your logs.
	RetentionLogsInterval *scw.Duration `json:"retention_logs_interval"`

	// RetentionTracesInterval: interval of time during which Scaleway's Cockpit keeps your traces.
	RetentionTracesInterval *scw.Duration `json:"retention_traces_interval"`

	// SampleIngestionPrice: ingestion price in cents for 1 million samples.
	SampleIngestionPrice uint32 `json:"sample_ingestion_price"`

	// LogsIngestionPrice: ingestion price in cents for 1 GB of logs.
	LogsIngestionPrice uint32 `json:"logs_ingestion_price"`

	// TracesIngestionPrice: ingestion price in cents for 1 GB of traces.
	TracesIngestionPrice uint32 `json:"traces_ingestion_price"`

	// RetentionPrice: retention price in euros per month.
	RetentionPrice uint32 `json:"retention_price"`
}

// ContactPoint: Contact point.
type ContactPoint struct {
	// Email: contact point configuration.
	// Precisely one of Email must be set.
	Email *ContactPointEmail `json:"email,omitempty"`
}

// Datasource: Data source.
type Datasource struct {
	// ID: ID of the data source.
	ID string `json:"id"`

	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`

	// Name: data source name.
	Name string `json:"name"`

	// URL: data source URL.
	URL string `json:"url"`

	// Type: data source type.
	// Default value: unknown_datasource_type
	Type DatasourceType `json:"type"`

	// IsManagedByScaleway: specifies that the data source receives data from Scaleway products and is managed by Scaleway.
	IsManagedByScaleway bool `json:"is_managed_by_scaleway"`
}

// GrafanaProductDashboard: Grafana dashboard.
type GrafanaProductDashboard struct {
	// DashboardName: name of the dashboard.
	DashboardName string `json:"dashboard_name"`

	// Title: title of the dashboard.
	Title string `json:"title"`

	// URL: URL of the dashboard.
	URL string `json:"url"`

	// Tags: tags of the dashboard.
	Tags []string `json:"tags"`

	// Variables: variables of the dashboard.
	Variables []string `json:"variables"`
}

// GrafanaUser: Grafana user.
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

// ActivateCockpitRequest: activate cockpit request.
type ActivateCockpitRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`
}

// Cockpit: Cockpit.
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

// CockpitMetrics: Metrics for a given Cockpit.
type CockpitMetrics struct {
	// Timeseries: time series array.
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

// CreateContactPointRequest: Request to create a contact point.
type CreateContactPointRequest struct {
	// ProjectID: ID of the Project in which to create the contact point.
	ProjectID string `json:"project_id"`

	// ContactPoint: contact point to create.
	ContactPoint *ContactPoint `json:"contact_point,omitempty"`
}

// CreateDatasourceRequest: Request to create a data source.
type CreateDatasourceRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`

	// Name: data source name.
	Name string `json:"name"`

	// Type: data source type.
	// Default value: unknown_datasource_type
	Type DatasourceType `json:"type"`

	// IsDefault: specifies that the returned output is the default data source per type.
	IsDefault bool `json:"is_default"`
}

// CreateGrafanaUserRequest: Request to create a Grafana user.
type CreateGrafanaUserRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`

	// Login: username of the Grafana user.
	Login string `json:"login"`

	// Role: role assigned to the Grafana user.
	// Default value: unknown_role
	Role GrafanaUserRole `json:"role"`
}

// CreateTokenRequest: create token request.
type CreateTokenRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`

	// Name: name of the token.
	Name string `json:"name"`

	// Scopes: token's permissions.
	Scopes *TokenScopes `json:"scopes,omitempty"`
}

// DeactivateCockpitRequest: deactivate cockpit request.
type DeactivateCockpitRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`
}

// DeleteContactPointRequest: Request to delete a contact point.
type DeleteContactPointRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`

	// ContactPoint: contact point to delete.
	ContactPoint *ContactPoint `json:"contact_point,omitempty"`
}

// DeleteDatasourceRequest: Request to delete a data source.
type DeleteDatasourceRequest struct {
	// DatasourceID: ID of the data source.
	DatasourceID string `json:"-"`
}

// DeleteGrafanaUserRequest: Request to delete a Grafana user.
type DeleteGrafanaUserRequest struct {
	// GrafanaUserID: ID of the Grafana user.
	GrafanaUserID uint32 `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// DeleteTokenRequest: delete token request.
type DeleteTokenRequest struct {
	// TokenID: ID of the token.
	TokenID string `json:"-"`
}

// DisableManagedAlertsRequest: Request to disable the sending of managed alerts.
type DisableManagedAlertsRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// EnableManagedAlertsRequest: Request to enable the sending of managed alerts.
type EnableManagedAlertsRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// GetCockpitMetricsRequest: Request to get a given Cockpit's metrics.
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

// GetCockpitRequest: get cockpit request.
type GetCockpitRequest struct {
	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"-"`
}

// GetGrafanaProductDashboardRequest: Request to get a dashboard.
type GetGrafanaProductDashboardRequest struct {
	// DashboardName: name of the dashboard.
	DashboardName string `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// GetTokenRequest: get token request.
type GetTokenRequest struct {
	// TokenID: ID of the token.
	TokenID string `json:"-"`
}

// ListContactPointsRequest: Request to list all contact points.
type ListContactPointsRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// ProjectID: ID of the Project from which to list the contact points.
	ProjectID string `json:"-"`
}

// ListContactPointsResponse: Response returned when listing contact points.
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

// ListDatasourcesRequest: list datasources request.
type ListDatasourcesRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the response is ordered.
	// Default value: created_at_asc
	OrderBy ListDatasourcesRequestOrderBy `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`

	// Types: filter by datasource types.
	Types []DatasourceType `json:"-"`

	// IsManagedByScaleway: filter by managed datasources.
	IsManagedByScaleway *bool `json:"-"`
}

// ListDatasourcesResponse: list datasources response.
type ListDatasourcesResponse struct {
	// TotalCount: count of all datasources corresponding to the request.
	TotalCount uint32 `json:"total_count"`

	// Datasources: list of the datasources within the pagination.
	Datasources []*Datasource `json:"datasources"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatasourcesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatasourcesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDatasourcesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Datasources = append(r.Datasources, results.Datasources...)
	r.TotalCount += uint32(len(results.Datasources))
	return uint32(len(results.Datasources)), nil
}

// ListGrafanaProductDashboardsRequest: Request to get a list of dashboards.
type ListGrafanaProductDashboardsRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// Tags: tags to filter the dashboards.
	Tags []string `json:"-"`
}

// ListGrafanaProductDashboardsResponse: Response returned when getting a list of dashboards.
type ListGrafanaProductDashboardsResponse struct {
	// TotalCount: count of grafana dasboards.
	TotalCount uint64 `json:"total_count"`

	// Dashboards: information on grafana dashboards.
	Dashboards []*GrafanaProductDashboard `json:"dashboards"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListGrafanaProductDashboardsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListGrafanaProductDashboardsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListGrafanaProductDashboardsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Dashboards = append(r.Dashboards, results.Dashboards...)
	r.TotalCount += uint64(len(results.Dashboards))
	return uint64(len(results.Dashboards)), nil
}

// ListGrafanaUsersRequest: Request to list all Grafana users.
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

// ListGrafanaUsersResponse: Response returned when listing Grafana users.
type ListGrafanaUsersResponse struct {
	// TotalCount: count of all Grafana users.
	TotalCount uint32 `json:"total_count"`

	// GrafanaUsers: information on all Grafana users.
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

// ListPlansRequest: Request to list all pricing plans.
type ListPlansRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListPlansRequestOrderBy `json:"-"`
}

// ListPlansResponse: Response returned when listing all pricing plans.
type ListPlansResponse struct {
	// TotalCount: count of all pricing plans.
	TotalCount uint64 `json:"total_count"`

	// Plans: information on plans.
	Plans []*Plan `json:"plans"`
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

// ListTokensRequest: list tokens request.
type ListTokensRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the response is ordered.
	// Default value: created_at_asc
	OrderBy ListTokensRequestOrderBy `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// ListTokensResponse: list tokens response.
type ListTokensResponse struct {
	// TotalCount: count of all tokens created.
	TotalCount uint32 `json:"total_count"`

	// Tokens: list of all tokens created.
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

// ResetGrafanaUserPasswordRequest: Request to reset a Grafana user's password.
type ResetGrafanaUserPasswordRequest struct {
	// GrafanaUserID: ID of the Grafana user.
	GrafanaUserID uint32 `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// SelectPlanRequest: Request to select a specific pricing plan.
type SelectPlanRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`

	// PlanID: ID of the pricing plan.
	PlanID string `json:"plan_id"`
}

// SelectPlanResponse: Response returned when selecting a pricing plan.
type SelectPlanResponse struct {
}

// TriggerTestAlertRequest: trigger test alert request.
type TriggerTestAlertRequest struct {
	ProjectID string `json:"project_id"`
}

// The Cockpit API allows you to activate your Cockpit to store metrics and logs. It also provides you with a dedicated Grafana for dashboarding to visualize your metrics and logs.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ActivateCockpit: Activate the Cockpit of a given Project specified by the Project ID.
func (s *API) ActivateCockpit(req *ActivateCockpitRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/activate",
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

// GetCockpit: Retrieve the Cockpit of a given Project specified by the Project ID.
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
		Path:   "/cockpit/v1beta1/cockpit",
		Query:  query,
	}

	var resp Cockpit

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCockpitMetrics: Retrieve metrics from your Cockpit specified by the ID of the Project the Cockpit belongs to.
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
		Method: "GET",
		Path:   "/cockpit/v1beta1/cockpit/metrics",
		Query:  query,
	}

	var resp CockpitMetrics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeactivateCockpit: Deactivate the Cockpit of a given Project specified by the Project ID.
func (s *API) DeactivateCockpit(req *DeactivateCockpitRequest, opts ...scw.RequestOption) (*Cockpit, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/deactivate",
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

// CreateDatasource: Create a data source for a given Project specified by the Project ID and the data source type.
func (s *API) CreateDatasource(req *CreateDatasourceRequest, opts ...scw.RequestOption) (*Datasource, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/datasources",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Datasource

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDatasource: Delete a given data source specified by the data source ID.
func (s *API) DeleteDatasource(req *DeleteDatasourceRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.DatasourceID) == "" {
		return errors.New("field DatasourceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/cockpit/v1beta1/datasources/" + fmt.Sprint(req.DatasourceID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListDatasources: Get a list of data sources for the specified Project ID.
func (s *API) ListDatasources(req *ListDatasourcesRequest, opts ...scw.RequestOption) (*ListDatasourcesResponse, error) {
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
	parameter.AddToQuery(query, "types", req.Types)
	parameter.AddToQuery(query, "is_managed_by_scaleway", req.IsManagedByScaleway)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/datasources",
		Query:  query,
	}

	var resp ListDatasourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateToken: Create a token in a given Project specified by the Project ID.
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
		Method: "POST",
		Path:   "/cockpit/v1beta1/tokens",
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

// ListTokens: Get a list of tokens in a given Project specified by the Project ID.
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
		Path:   "/cockpit/v1beta1/tokens",
		Query:  query,
	}

	var resp ListTokensResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetToken: Retrieve a given token specified by the token ID.
func (s *API) GetToken(req *GetTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if fmt.Sprint(req.TokenID) == "" {
		return nil, errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteToken: Delete a given token specified by the token ID.
func (s *API) DeleteToken(req *DeleteTokenRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.TokenID) == "" {
		return errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/cockpit/v1beta1/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateContactPoint: Create a contact point associated with the default receiver, to receive alerts.
func (s *API) CreateContactPoint(req *CreateContactPointRequest, opts ...scw.RequestOption) (*ContactPoint, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/contact-points",
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

// ListContactPoints: Get a list of contact points created for a given Cockpit, specified by the ID of the Project the Cockpit belongs to.
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
		Path:   "/cockpit/v1beta1/contact-points",
		Query:  query,
	}

	var resp ListContactPointsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteContactPoint: Delete a contact point associated with the default receiver.
func (s *API) DeleteContactPoint(req *DeleteContactPointRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/delete-contact-point",
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

// EnableManagedAlerts: Enable the sending of managed alerts for a given Cockpit, specified by the ID of the Project the Cockpit belongs to.
func (s *API) EnableManagedAlerts(req *EnableManagedAlertsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/enable-managed-alerts",
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

// DisableManagedAlerts: Disable the sending of managed alerts for a given Cockpit, specified by the ID of the Project the Cockpit belongs to.
func (s *API) DisableManagedAlerts(req *DisableManagedAlertsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/disable-managed-alerts",
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

// TriggerTestAlert: Send a test alert to make sure your contact points get notified when an actual alert is triggered.
func (s *API) TriggerTestAlert(req *TriggerTestAlertRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/trigger-test-alert",
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

// CreateGrafanaUser: Create a Grafana user for your Cockpit's Grafana. Make sure you save the automatically-generated password and the Grafana user ID.
func (s *API) CreateGrafanaUser(req *CreateGrafanaUserRequest, opts ...scw.RequestOption) (*GrafanaUser, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/grafana-users",
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

// ListGrafanaUsers: Get a list of all Grafana users created in your Cockpit's Grafana.
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
		Path:   "/cockpit/v1beta1/grafana-users",
		Query:  query,
	}

	var resp ListGrafanaUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteGrafanaUser: Delete a Grafana user from your Cockpit's Grafana, specified by the ID of the Project the Cockpit belongs to, and the ID of the Grafana user.
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
		Path:   "/cockpit/v1beta1/grafana-users/" + fmt.Sprint(req.GrafanaUserID) + "/delete",
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

// ResetGrafanaUserPassword: Reset the password of a Grafana user, specified by the ID of the Project the Cockpit belongs to, and the ID of the Grafana user.
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
		Path:   "/cockpit/v1beta1/grafana-users/" + fmt.Sprint(req.GrafanaUserID) + "/reset-password",
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

// ListPlans: Get a list of all pricing plans available.
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
		Method: "GET",
		Path:   "/cockpit/v1beta1/plans",
		Query:  query,
	}

	var resp ListPlansResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SelectPlan: Select your chosen pricing plan for your Cockpit, specifying the Cockpit's Project ID and the pricing plan's ID in the request.
func (s *API) SelectPlan(req *SelectPlanRequest, opts ...scw.RequestOption) (*SelectPlanResponse, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1beta1/select-plan",
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

// ListGrafanaProductDashboards: Get a list of available product dashboards.
func (s *API) ListGrafanaProductDashboards(req *ListGrafanaProductDashboardsRequest, opts ...scw.RequestOption) (*ListGrafanaProductDashboardsResponse, error) {
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "tags", req.Tags)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/grafana-product-dashboards",
		Query:  query,
	}

	var resp ListGrafanaProductDashboardsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGrafanaProductDashboard: Get a product dashboard specified by the dashboard ID.
func (s *API) GetGrafanaProductDashboard(req *GetGrafanaProductDashboardRequest, opts ...scw.RequestOption) (*GrafanaProductDashboard, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.DashboardName) == "" {
		return nil, errors.New("field DashboardName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1beta1/grafana-product-dashboards/" + fmt.Sprint(req.DashboardName) + "",
		Query:  query,
	}

	var resp GrafanaProductDashboard

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

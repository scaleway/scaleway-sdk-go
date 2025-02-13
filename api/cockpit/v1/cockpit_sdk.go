// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package cockpit provides methods and message types of the cockpit v1 API.
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

	"github.com/scaleway/scaleway-sdk-go/errors"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
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

type DataSourceOrigin string

const (
	// Unknown data source origin.
	DataSourceOriginUnknownOrigin = DataSourceOrigin("unknown_origin")
	// Data source managed by Scaleway, used to store and query metrics and logs from Scaleway resources.
	DataSourceOriginScaleway = DataSourceOrigin("scaleway")
	// Data source created by the user, used to store and query metrics, logs and traces from user's custom resources.
	DataSourceOriginExternal = DataSourceOrigin("external")
	// Data source created by the user, used to store and query metrics, logs and traces from user's custom resources.
	DataSourceOriginCustom = DataSourceOrigin("custom")
)

func (enum DataSourceOrigin) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_origin"
	}
	return string(enum)
}

func (enum DataSourceOrigin) Values() []DataSourceOrigin {
	return []DataSourceOrigin{
		"unknown_origin",
		"scaleway",
		"external",
		"custom",
	}
}

func (enum DataSourceOrigin) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DataSourceOrigin) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DataSourceOrigin(DataSourceOrigin(tmp).String())
	return nil
}

type DataSourceType string

const (
	// Unknown data source type.
	DataSourceTypeUnknownType = DataSourceType("unknown_type")
	// Metrics data source type, used to store and query metrics using Grafana Mimir.
	DataSourceTypeMetrics = DataSourceType("metrics")
	// Logs data source type, used to store and query logs using Grafana Loki.
	DataSourceTypeLogs = DataSourceType("logs")
	// Traces data source type, used to store and query traces using Grafana Tempo.
	DataSourceTypeTraces = DataSourceType("traces")
)

func (enum DataSourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum DataSourceType) Values() []DataSourceType {
	return []DataSourceType{
		"unknown_type",
		"metrics",
		"logs",
		"traces",
	}
}

func (enum DataSourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DataSourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DataSourceType(DataSourceType(tmp).String())
	return nil
}

type GrafanaUserRole string

const (
	// Unknown role.
	GrafanaUserRoleUnknownRole = GrafanaUserRole("unknown_role")
	// Editor role.
	GrafanaUserRoleEditor = GrafanaUserRole("editor")
	// Viewer role.
	GrafanaUserRoleViewer = GrafanaUserRole("viewer")
)

func (enum GrafanaUserRole) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_role"
	}
	return string(enum)
}

func (enum GrafanaUserRole) Values() []GrafanaUserRole {
	return []GrafanaUserRole{
		"unknown_role",
		"editor",
		"viewer",
	}
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

type ListDataSourcesRequestOrderBy string

const (
	ListDataSourcesRequestOrderByCreatedAtAsc  = ListDataSourcesRequestOrderBy("created_at_asc")
	ListDataSourcesRequestOrderByCreatedAtDesc = ListDataSourcesRequestOrderBy("created_at_desc")
	ListDataSourcesRequestOrderByNameAsc       = ListDataSourcesRequestOrderBy("name_asc")
	ListDataSourcesRequestOrderByNameDesc      = ListDataSourcesRequestOrderBy("name_desc")
	ListDataSourcesRequestOrderByTypeAsc       = ListDataSourcesRequestOrderBy("type_asc")
	ListDataSourcesRequestOrderByTypeDesc      = ListDataSourcesRequestOrderBy("type_desc")
)

func (enum ListDataSourcesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListDataSourcesRequestOrderBy) Values() []ListDataSourcesRequestOrderBy {
	return []ListDataSourcesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
		"type_asc",
		"type_desc",
	}
}

func (enum ListDataSourcesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDataSourcesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDataSourcesRequestOrderBy(ListDataSourcesRequestOrderBy(tmp).String())
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

func (enum ListGrafanaUsersRequestOrderBy) Values() []ListGrafanaUsersRequestOrderBy {
	return []ListGrafanaUsersRequestOrderBy{
		"login_asc",
		"login_desc",
	}
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

type ListManagedAlertsRequestOrderBy string

const (
	ListManagedAlertsRequestOrderByCreatedAtAsc  = ListManagedAlertsRequestOrderBy("created_at_asc")
	ListManagedAlertsRequestOrderByCreatedAtDesc = ListManagedAlertsRequestOrderBy("created_at_desc")
	ListManagedAlertsRequestOrderByNameAsc       = ListManagedAlertsRequestOrderBy("name_asc")
	ListManagedAlertsRequestOrderByNameDesc      = ListManagedAlertsRequestOrderBy("name_desc")
	ListManagedAlertsRequestOrderByTypeAsc       = ListManagedAlertsRequestOrderBy("type_asc")
	ListManagedAlertsRequestOrderByTypeDesc      = ListManagedAlertsRequestOrderBy("type_desc")
)

func (enum ListManagedAlertsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListManagedAlertsRequestOrderBy) Values() []ListManagedAlertsRequestOrderBy {
	return []ListManagedAlertsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
		"type_asc",
		"type_desc",
	}
}

func (enum ListManagedAlertsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListManagedAlertsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListManagedAlertsRequestOrderBy(ListManagedAlertsRequestOrderBy(tmp).String())
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

func (enum ListPlansRequestOrderBy) Values() []ListPlansRequestOrderBy {
	return []ListPlansRequestOrderBy{
		"name_asc",
		"name_desc",
	}
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

func (enum ListTokensRequestOrderBy) Values() []ListTokensRequestOrderBy {
	return []ListTokensRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
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

func (enum PlanName) Values() []PlanName {
	return []PlanName{
		"unknown_name",
		"free",
		"premium",
		"custom",
	}
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

type TokenScope string

const (
	// The token's permission scope is unknown.
	TokenScopeUnknownScope = TokenScope("unknown_scope")
	// Token has permission to read data from metrics data sources.
	TokenScopeReadOnlyMetrics = TokenScope("read_only_metrics")
	// Token has permission to write data in metrics data sources.
	TokenScopeWriteOnlyMetrics = TokenScope("write_only_metrics")
	// Token has permission to read and write prometheus rules in metrics data sources.
	TokenScopeFullAccessMetricsRules = TokenScope("full_access_metrics_rules")
	// Token has permission to read data from logs data sources.
	TokenScopeReadOnlyLogs = TokenScope("read_only_logs")
	// Token has permission to write data in logs data sources.
	TokenScopeWriteOnlyLogs = TokenScope("write_only_logs")
	// Token has permission to read and write prometheus rules in logs data sources.
	TokenScopeFullAccessLogsRules = TokenScope("full_access_logs_rules")
	// Token has permission to read and write the Alert manager API.
	TokenScopeFullAccessAlertManager = TokenScope("full_access_alert_manager")
	// Token has permission to read data from traces data sources.
	TokenScopeReadOnlyTraces = TokenScope("read_only_traces")
	// Token has permission to write data in traces data sources.
	TokenScopeWriteOnlyTraces = TokenScope("write_only_traces")
)

func (enum TokenScope) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_scope"
	}
	return string(enum)
}

func (enum TokenScope) Values() []TokenScope {
	return []TokenScope{
		"unknown_scope",
		"read_only_metrics",
		"write_only_metrics",
		"full_access_metrics_rules",
		"read_only_logs",
		"write_only_logs",
		"full_access_logs_rules",
		"full_access_alert_manager",
		"read_only_traces",
		"write_only_traces",
	}
}

func (enum TokenScope) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TokenScope) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TokenScope(TokenScope(tmp).String())
	return nil
}

type UsageUnit string

const (
	UsageUnitUnknownUnit = UsageUnit("unknown_unit")
	UsageUnitBytes       = UsageUnit("bytes")
	UsageUnitSamples     = UsageUnit("samples")
)

func (enum UsageUnit) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_unit"
	}
	return string(enum)
}

func (enum UsageUnit) Values() []UsageUnit {
	return []UsageUnit{
		"unknown_unit",
		"bytes",
		"samples",
	}
}

func (enum UsageUnit) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UsageUnit) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UsageUnit(UsageUnit(tmp).String())
	return nil
}

// ContactPointEmail: contact point email.
type ContactPointEmail struct {
	To string `json:"to"`
}

// GetConfigResponseRetention: get config response retention.
type GetConfigResponseRetention struct {
	MinDays uint32 `json:"min_days"`

	MaxDays uint32 `json:"max_days"`

	DefaultDays uint32 `json:"default_days"`
}

// ContactPoint: Contact point.
type ContactPoint struct {
	// Email: email address to send alerts to.
	// Precisely one of Email must be set.
	Email *ContactPointEmail `json:"email,omitempty"`

	// Region: region.
	Region scw.Region `json:"region"`

	// ReceiveResolvedNotifications: send an email notification when an alert is marked as resolved.
	ReceiveResolvedNotifications bool `json:"receive_resolved_notifications"`
}

// DataSource: Data source.
type DataSource struct {
	// ID: ID of the data source.
	ID string `json:"id"`

	// ProjectID: ID of the Project the data source belongs to.
	ProjectID string `json:"project_id"`

	// Name: data source name.
	Name string `json:"name"`

	// URL: data source URL.
	URL string `json:"url"`

	// Type: data source type.
	// Default value: unknown_type
	Type DataSourceType `json:"type"`

	// Origin: data source origin.
	// Default value: unknown_origin
	Origin DataSourceOrigin `json:"origin"`

	// CreatedAt: date the data source was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date the data source was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// SynchronizedWithGrafana: indicates whether the data source is synchronized with Grafana.
	SynchronizedWithGrafana bool `json:"synchronized_with_grafana"`

	// RetentionDays: bETA - Duration for which the data will be retained in the data source.
	RetentionDays uint32 `json:"retention_days"`

	// Region: region of the data source.
	Region scw.Region `json:"region"`
}

// GrafanaProductDashboard: Grafana dashboard.
type GrafanaProductDashboard struct {
	// Name: dashboard name.
	Name string `json:"name"`

	// Title: dashboard title.
	Title string `json:"title"`

	// URL: dashboard URL.
	URL string `json:"url"`

	// Tags: dashboard tags.
	Tags []string `json:"tags"`

	// Variables: dashboard variables.
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

	// Password: grafana user's password.
	Password *string `json:"password"`
}

// Alert: alert.
type Alert struct {
	ProductFamily string `json:"product_family"`

	Product string `json:"product"`

	Name string `json:"name"`

	Rule string `json:"rule"`

	Description string `json:"description"`
}

// Plan: Type of pricing plan.
type Plan struct {
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

	// MonthlyPrice: retention price in euros per month.
	MonthlyPrice uint32 `json:"monthly_price"`
}

// Token: Token.
type Token struct {
	// ID: ID of the token.
	ID string `json:"id"`

	// ProjectID: ID of the Project the token belongs to.
	ProjectID string `json:"project_id"`

	// Name: name of the token.
	Name string `json:"name"`

	// CreatedAt: token creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: token last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Scopes: token permission scopes.
	Scopes []TokenScope `json:"scopes"`

	// SecretKey: token secret key.
	SecretKey *string `json:"secret_key"`

	// Region: regions where the token is located.
	Region scw.Region `json:"region"`
}

// Usage: Data source usage.
type Usage struct {
	// DataSourceID: ID of the data source.
	DataSourceID *string `json:"data_source_id"`

	// ProjectID: ID of the Project the data source belongs to.
	ProjectID string `json:"project_id"`

	// DataSourceOrigin: origin of the data source.
	// Default value: unknown_origin
	DataSourceOrigin DataSourceOrigin `json:"data_source_origin"`

	// DataSourceType: type of the data source.
	// Default value: unknown_type
	DataSourceType DataSourceType `json:"data_source_type"`

	// Unit: unit of the data source usage.
	// Default value: unknown_unit
	Unit UsageUnit `json:"unit"`

	// Interval: interval for the data source usage.
	Interval *scw.Duration `json:"interval"`

	// QuantityOverInterval: data source usage for the given interval.
	QuantityOverInterval uint64 `json:"quantity_over_interval"`

	// Region: region of the data source usage.
	Region scw.Region `json:"region"`
}

// AlertManager: Alert manager information.
type AlertManager struct {
	// AlertManagerURL: alert manager URL.
	AlertManagerURL *string `json:"alert_manager_url"`

	// AlertManagerEnabled: the Alert manager is enabled.
	AlertManagerEnabled bool `json:"alert_manager_enabled"`

	// ManagedAlertsEnabled: managed alerts are enabled.
	ManagedAlertsEnabled bool `json:"managed_alerts_enabled"`

	// Region: regions where the Alert manager is enabled.
	Region scw.Region `json:"region"`
}

// GetConfigResponse: Cockpit configuration.
type GetConfigResponse struct {
	// CustomMetricsRetention: custom metrics retention configuration.
	CustomMetricsRetention *GetConfigResponseRetention `json:"custom_metrics_retention"`

	// CustomLogsRetention: custom logs retention configuration.
	CustomLogsRetention *GetConfigResponseRetention `json:"custom_logs_retention"`

	// CustomTracesRetention: custom traces retention configuration.
	CustomTracesRetention *GetConfigResponseRetention `json:"custom_traces_retention"`

	// ProductMetricsRetention: scaleway metrics retention configuration.
	ProductMetricsRetention *GetConfigResponseRetention `json:"product_metrics_retention"`

	// ProductLogsRetention: scaleway logs retention configuration.
	ProductLogsRetention *GetConfigResponseRetention `json:"product_logs_retention"`
}

// GlobalAPICreateGrafanaUserRequest: Create a Grafana user.
type GlobalAPICreateGrafanaUserRequest struct {
	// ProjectID: ID of the Project in which to create the Grafana user.
	ProjectID string `json:"project_id"`

	// Login: username of the Grafana user. Note that the `admin` username is not available for creation.
	Login string `json:"login"`

	// Role: role assigned to the Grafana user.
	// Default value: unknown_role
	Role GrafanaUserRole `json:"role"`
}

// GlobalAPIDeleteGrafanaUserRequest: Delete a Grafana user.
type GlobalAPIDeleteGrafanaUserRequest struct {
	// GrafanaUserID: ID of the Grafana user.
	GrafanaUserID uint32 `json:"-"`

	// ProjectID: ID of the Project to target.
	ProjectID string `json:"-"`
}

// GlobalAPIGetCurrentPlanRequest: Retrieve a pricing plan for the given Project.
type GlobalAPIGetCurrentPlanRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// GlobalAPIGetGrafanaProductDashboardRequest: Retrieve a specific dashboard.
type GlobalAPIGetGrafanaProductDashboardRequest struct {
	// DashboardName: name of the dashboard.
	DashboardName string `json:"-"`

	// ProjectID: ID of the Project the dashboard belongs to.
	ProjectID string `json:"-"`
}

// GlobalAPIGetGrafanaRequest: Request a Grafana.
type GlobalAPIGetGrafanaRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// GlobalAPIListGrafanaProductDashboardsRequest: Retrieve a list of available product dashboards.
type GlobalAPIListGrafanaProductDashboardsRequest struct {
	// ProjectID: ID of the Project to target.
	ProjectID string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// Tags: tags to filter for.
	Tags []string `json:"-"`
}

// GlobalAPIListGrafanaUsersRequest: List all Grafana users.
type GlobalAPIListGrafanaUsersRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: order of the Grafana users.
	// Default value: login_asc
	OrderBy ListGrafanaUsersRequestOrderBy `json:"-"`

	// ProjectID: ID of the Project to target.
	ProjectID string `json:"-"`
}

// GlobalAPIListPlansRequest: Retrieve a list of available pricing plans.
type GlobalAPIListPlansRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListPlansRequestOrderBy `json:"-"`
}

// GlobalAPIResetGrafanaUserPasswordRequest: Reset a Grafana user's password.
type GlobalAPIResetGrafanaUserPasswordRequest struct {
	// GrafanaUserID: ID of the Grafana user.
	GrafanaUserID uint32 `json:"-"`

	// ProjectID: ID of the Project to target.
	ProjectID string `json:"project_id"`
}

// GlobalAPISelectPlanRequest: Select a specific pricing plan.
type GlobalAPISelectPlanRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`

	// PlanName: name of the pricing plan.
	// Default value: unknown_name
	PlanName PlanName `json:"plan_name"`
}

// GlobalAPISyncGrafanaDataSourcesRequest: Trigger the synchronization of all data sources created in the relevant regions.
type GlobalAPISyncGrafanaDataSourcesRequest struct {
	// ProjectID: ID of the Project to target.
	ProjectID string `json:"project_id"`
}

// Grafana: Grafana user.
type Grafana struct {
	// GrafanaURL: URL to access your Cockpit's Grafana.
	GrafanaURL string `json:"grafana_url"`
}

// ListContactPointsResponse: Response returned when listing contact points.
type ListContactPointsResponse struct {
	// TotalCount: total count of contact points associated with the default receiver.
	TotalCount uint64 `json:"total_count"`

	// ContactPoints: list of contact points associated with the default receiver.
	ContactPoints []*ContactPoint `json:"contact_points"`

	// HasAdditionalReceivers: indicates whether the Alert manager has other receivers than the default one.
	HasAdditionalReceivers bool `json:"has_additional_receivers"`

	// HasAdditionalContactPoints: indicates whether there are unmanaged contact points on the default receiver.
	HasAdditionalContactPoints bool `json:"has_additional_contact_points"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListContactPointsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListContactPointsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListContactPointsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ContactPoints = append(r.ContactPoints, results.ContactPoints...)
	r.TotalCount += uint64(len(results.ContactPoints))
	return uint64(len(results.ContactPoints)), nil
}

// ListDataSourcesResponse: Response returned when listing data sources.
type ListDataSourcesResponse struct {
	// TotalCount: total count of data sources matching the request.
	TotalCount uint64 `json:"total_count"`

	// DataSources: data sources matching the request within the pagination.
	DataSources []*DataSource `json:"data_sources"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDataSourcesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDataSourcesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListDataSourcesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.DataSources = append(r.DataSources, results.DataSources...)
	r.TotalCount += uint64(len(results.DataSources))
	return uint64(len(results.DataSources)), nil
}

// ListGrafanaProductDashboardsResponse: Output returned when listing dashboards.
type ListGrafanaProductDashboardsResponse struct {
	// TotalCount: total count of Grafana dashboards.
	TotalCount uint64 `json:"total_count"`

	// Dashboards: grafana dashboards information.
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

// ListGrafanaUsersResponse: Ouptut returned when listing Grafana users.
type ListGrafanaUsersResponse struct {
	// TotalCount: total count of Grafana users.
	TotalCount uint64 `json:"total_count"`

	// GrafanaUsers: grafana users information.
	GrafanaUsers []*GrafanaUser `json:"grafana_users"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListGrafanaUsersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListGrafanaUsersResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListGrafanaUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.GrafanaUsers = append(r.GrafanaUsers, results.GrafanaUsers...)
	r.TotalCount += uint64(len(results.GrafanaUsers))
	return uint64(len(results.GrafanaUsers)), nil
}

// ListManagedAlertsResponse: Response returned when listing data sources.
type ListManagedAlertsResponse struct {
	// TotalCount: total count of data sources matching the request.
	TotalCount uint64 `json:"total_count"`

	// Alerts: alerts matching the request within the pagination.
	Alerts []*Alert `json:"alerts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListManagedAlertsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListManagedAlertsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListManagedAlertsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Alerts = append(r.Alerts, results.Alerts...)
	r.TotalCount += uint64(len(results.Alerts))
	return uint64(len(results.Alerts)), nil
}

// ListPlansResponse: Output returned when listing pricing plans.
type ListPlansResponse struct {
	// TotalCount: total count of available pricing plans.
	TotalCount uint64 `json:"total_count"`

	// Plans: plan types information.
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

// ListTokensResponse: Response returned when listing tokens.
type ListTokensResponse struct {
	// TotalCount: total count of tokens matching the request.
	TotalCount uint64 `json:"total_count"`

	// Tokens: tokens matching the request within the pagination.
	Tokens []*Token `json:"tokens"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTokensResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTokensResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListTokensResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Tokens = append(r.Tokens, results.Tokens...)
	r.TotalCount += uint64(len(results.Tokens))
	return uint64(len(results.Tokens)), nil
}

// RegionalAPICreateContactPointRequest: Create a contact point.
type RegionalAPICreateContactPointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project to create the contact point in.
	ProjectID string `json:"project_id"`

	// Email: email address of the contact point to create.
	// Precisely one of Email must be set.
	Email *ContactPointEmail `json:"email,omitempty"`

	// ReceiveResolvedNotifications: send an email notification when an alert is marked as resolved.
	ReceiveResolvedNotifications *bool `json:"receive_resolved_notifications,omitempty"`
}

// RegionalAPICreateDataSourceRequest: Create a data source.
type RegionalAPICreateDataSourceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project the data source belongs to.
	ProjectID string `json:"project_id"`

	// Name: data source name.
	Name string `json:"name"`

	// Type: data source type.
	// Default value: unknown_type
	Type DataSourceType `json:"type"`

	// RetentionDays: default values are 30 days for metrics, 7 days for logs and traces.
	RetentionDays *uint32 `json:"retention_days,omitempty"`
}

// RegionalAPICreateTokenRequest: Create a token.
type RegionalAPICreateTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project the token belongs to.
	ProjectID string `json:"project_id"`

	// Name: name of the token.
	Name string `json:"name"`

	// TokenScopes: token permission scopes.
	TokenScopes []TokenScope `json:"token_scopes"`
}

// RegionalAPIDeleteContactPointRequest: Delete a contact point.
type RegionalAPIDeleteContactPointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project containing the contact point to delete.
	ProjectID string `json:"project_id"`

	// Email: email address of the contact point to delete.
	// Precisely one of Email must be set.
	Email *ContactPointEmail `json:"email,omitempty"`
}

// RegionalAPIDeleteDataSourceRequest: Delete a data source.
type RegionalAPIDeleteDataSourceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DataSourceID: ID of the data source to delete.
	DataSourceID string `json:"-"`
}

// RegionalAPIDeleteTokenRequest: Delete a token.
type RegionalAPIDeleteTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TokenID: ID of the token to delete.
	TokenID string `json:"-"`
}

// RegionalAPIDisableAlertManagerRequest: Disable the Alert manager.
type RegionalAPIDisableAlertManagerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project to disable the Alert manager in.
	ProjectID string `json:"project_id"`
}

// RegionalAPIDisableManagedAlertsRequest: Disable the sending of managed alerts.
type RegionalAPIDisableManagedAlertsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// RegionalAPIEnableAlertManagerRequest: Enable the Alert manager.
type RegionalAPIEnableAlertManagerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project to enable the Alert manager in.
	ProjectID string `json:"project_id"`
}

// RegionalAPIEnableManagedAlertsRequest: Enable the sending of managed alerts.
type RegionalAPIEnableManagedAlertsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// RegionalAPIGetAlertManagerRequest: Get the Alert manager.
type RegionalAPIGetAlertManagerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project ID of the requested Alert manager.
	ProjectID string `json:"project_id"`
}

// RegionalAPIGetConfigRequest: Get Cockpit configuration.
type RegionalAPIGetConfigRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// RegionalAPIGetDataSourceRequest: Retrieve a data source.
type RegionalAPIGetDataSourceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DataSourceID: ID of the relevant data source.
	DataSourceID string `json:"-"`
}

// RegionalAPIGetTokenRequest: Get a token.
type RegionalAPIGetTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TokenID: token ID.
	TokenID string `json:"-"`
}

// RegionalAPIGetUsageOverviewRequest: regional api get usage overview request.
type RegionalAPIGetUsageOverviewRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	ProjectID string `json:"-"`

	Interval *scw.Duration `json:"-"`
}

// RegionalAPIListContactPointsRequest: List contact points.
type RegionalAPIListContactPointsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: total count of contact points to return per page.
	PageSize *uint32 `json:"-"`

	// ProjectID: ID of the Project containing the contact points to list.
	ProjectID string `json:"-"`
}

// RegionalAPIListDataSourcesRequest: List data sources.
type RegionalAPIListDataSourcesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of data sources to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order for data sources in the response.
	// Default value: created_at_asc
	OrderBy ListDataSourcesRequestOrderBy `json:"-"`

	// ProjectID: project ID to filter for, only data sources from this Project will be returned.
	ProjectID string `json:"-"`

	// Origin: origin to filter for, only data sources with matching origin will be returned.
	// Default value: unknown_origin
	Origin DataSourceOrigin `json:"-"`

	// Types: types to filter for, only data sources with matching types will be returned.
	Types []DataSourceType `json:"-"`
}

// RegionalAPIListManagedAlertsRequest: Enable the sending of managed alerts.
type RegionalAPIListManagedAlertsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of data sources to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order for data sources in the response.
	// Default value: created_at_asc
	OrderBy ListManagedAlertsRequestOrderBy `json:"-"`

	// ProjectID: project ID to filter for, only data sources from this Project will be returned.
	ProjectID string `json:"-"`
}

// RegionalAPIListTokensRequest: List tokens.
type RegionalAPIListTokensRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of tokens to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListTokensRequestOrderBy `json:"-"`

	// ProjectID: ID of the Project the tokens belong to.
	ProjectID string `json:"-"`

	// TokenScopes: token scopes to filter for.
	TokenScopes []TokenScope `json:"-"`
}

// RegionalAPITriggerTestAlertRequest: Request to trigger a test alert.
type RegionalAPITriggerTestAlertRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// RegionalAPIUpdateContactPointRequest: Update a contact point.
type RegionalAPIUpdateContactPointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project containing the contact point to update.
	ProjectID string `json:"project_id"`

	// Email: email address of the contact point to update.
	// Precisely one of Email must be set.
	Email *ContactPointEmail `json:"email,omitempty"`

	// ReceiveResolvedNotifications: enable or disable notifications when alert is resolved.
	ReceiveResolvedNotifications *bool `json:"receive_resolved_notifications,omitempty"`
}

// RegionalAPIUpdateDataSourceRequest: Update a data source name.
type RegionalAPIUpdateDataSourceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DataSourceID: ID of the data source to update.
	DataSourceID string `json:"-"`

	// Name: updated name of the data source.
	Name *string `json:"name,omitempty"`

	// RetentionDays: bETA - Duration for which the data will be retained in the data source.
	RetentionDays *uint32 `json:"retention_days,omitempty"`
}

// UsageOverview: usage overview.
type UsageOverview struct {
	ScalewayMetricsUsage *Usage `json:"scaleway_metrics_usage"`

	ScalewayLogsUsage *Usage `json:"scaleway_logs_usage"`

	ExternalMetricsUsage *Usage `json:"external_metrics_usage"`

	ExternalLogsUsage *Usage `json:"external_logs_usage"`

	ExternalTracesUsage *Usage `json:"external_traces_usage"`
}

// The Cockpit Global API allows you to manage your Cockpit's Grafana and plans.
type GlobalAPI struct {
	client *scw.Client
}

// NewGlobalAPI returns a GlobalAPI object from a Scaleway client.
func NewGlobalAPI(client *scw.Client) *GlobalAPI {
	return &GlobalAPI{
		client: client,
	}
}

// GetGrafana: Retrieve information on your Cockpit's Grafana, specified by the ID of the Project the Cockpit belongs to.
// The output returned displays the URL to access your Cockpit's Grafana.
func (s *GlobalAPI) GetGrafana(req *GlobalAPIGetGrafanaRequest, opts ...scw.RequestOption) (*Grafana, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/grafana",
		Query:  query,
	}

	var resp Grafana

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncGrafanaDataSources: Trigger the synchronization of all your data sources and the alert manager in the relevant regions. The alert manager will only be synchronized if you have enabled it.
func (s *GlobalAPI) SyncGrafanaDataSources(req *GlobalAPISyncGrafanaDataSourcesRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1/grafana/sync-data-sources",
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

// CreateGrafanaUser: Create a Grafana user to connect to your Cockpit's Grafana. Upon creation, your user password displays only once, so make sure that you save it.
// Each Grafana user is associated with a role: viewer or editor. A viewer can only view dashboards, whereas an editor can create and edit dashboards. Note that the `admin` username is not available for creation.
func (s *GlobalAPI) CreateGrafanaUser(req *GlobalAPICreateGrafanaUserRequest, opts ...scw.RequestOption) (*GrafanaUser, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1/grafana/users",
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

// ListGrafanaUsers: List all Grafana users created in your Cockpit's Grafana. By default, the Grafana users returned in the list are ordered in ascending order.
func (s *GlobalAPI) ListGrafanaUsers(req *GlobalAPIListGrafanaUsersRequest, opts ...scw.RequestOption) (*ListGrafanaUsersResponse, error) {
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
		Path:   "/cockpit/v1/grafana/users",
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
func (s *GlobalAPI) DeleteGrafanaUser(req *GlobalAPIDeleteGrafanaUserRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.GrafanaUserID) == "" {
		return errors.New("field GrafanaUserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/cockpit/v1/grafana/users/" + fmt.Sprint(req.GrafanaUserID) + "",
		Query:  query,
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ResetGrafanaUserPassword: Reset the password of a Grafana user, specified by the ID of the Project the Cockpit belongs to, and the ID of the Grafana user.
// A new password regenerates and only displays once. Make sure that you save it.
func (s *GlobalAPI) ResetGrafanaUserPassword(req *GlobalAPIResetGrafanaUserPasswordRequest, opts ...scw.RequestOption) (*GrafanaUser, error) {
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
		Path:   "/cockpit/v1/grafana/users/" + fmt.Sprint(req.GrafanaUserID) + "/reset-password",
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

// ListGrafanaProductDashboards: Retrieve a list of available dashboards in Grafana, for all Scaleway resources which are integrated with Cockpit.
func (s *GlobalAPI) ListGrafanaProductDashboards(req *GlobalAPIListGrafanaProductDashboardsRequest, opts ...scw.RequestOption) (*ListGrafanaProductDashboardsResponse, error) {
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
		Path:   "/cockpit/v1/grafana/product-dashboards",
		Query:  query,
	}

	var resp ListGrafanaProductDashboardsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetGrafanaProductDashboard: Retrieve information about the dashboard of a Scaleway resource in Grafana, specified by the ID of the Project the Cockpit belongs to, and the name of the dashboard.
func (s *GlobalAPI) GetGrafanaProductDashboard(req *GlobalAPIGetGrafanaProductDashboardRequest, opts ...scw.RequestOption) (*GrafanaProductDashboard, error) {
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
		Path:   "/cockpit/v1/grafana/product-dashboards/" + fmt.Sprint(req.DashboardName) + "",
		Query:  query,
	}

	var resp GrafanaProductDashboard

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: ListPlans: Retrieve a list of available pricing plan types.
// Deprecated, retention is now managed at the data source level.
func (s *GlobalAPI) ListPlans(req *GlobalAPIListPlansRequest, opts ...scw.RequestOption) (*ListPlansResponse, error) {
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
		Path:   "/cockpit/v1/plans",
		Query:  query,
	}

	var resp ListPlansResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: SelectPlan: Apply a pricing plan on a given Project. You must specify the ID of the pricing plan type. Note that you will be billed for the plan you apply.
// Deprecated, retention is now managed at the data source level.
func (s *GlobalAPI) SelectPlan(req *GlobalAPISelectPlanRequest, opts ...scw.RequestOption) (*Plan, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/cockpit/v1/plans",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Plan

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: GetCurrentPlan: Retrieve a pricing plan for the given Project, specified by the ID of the Project.
// Deprecated, retention is now managed at the data source level.
func (s *GlobalAPI) GetCurrentPlan(req *GlobalAPIGetCurrentPlanRequest, opts ...scw.RequestOption) (*Plan, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/current-plan",
		Query:  query,
	}

	var resp Plan

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// The Cockpit Regional API allows you to create data sources and tokens to store and query data types such as metrics, logs, and traces. You can also push your data into Cockpit, and send alerts to your contact points when your resources may require your attention, using the regional Alert manager.
type RegionalAPI struct {
	client *scw.Client
}

// NewRegionalAPI returns a RegionalAPI object from a Scaleway client.
func NewRegionalAPI(client *scw.Client) *RegionalAPI {
	return &RegionalAPI{
		client: client,
	}
}
func (s *RegionalAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// GetConfig: Get the Cockpit configuration.
func (s *RegionalAPI) GetConfig(req *RegionalAPIGetConfigRequest, opts ...scw.RequestOption) (*GetConfigResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/config",
	}

	var resp GetConfigResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDataSource: You must specify the data source type upon creation. Available data source types include:
//   - metrics
//   - logs
//   - traces
//
// The name of the data source will then be used as reference to name the associated Grafana data source.
func (s *RegionalAPI) CreateDataSource(req *RegionalAPICreateDataSourceRequest, opts ...scw.RequestOption) (*DataSource, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/data-sources",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DataSource

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDataSource: Retrieve information about a given data source, specified by the data source ID. The data source's information such as its name, type, URL, origin, and retention period, is returned.
func (s *RegionalAPI) GetDataSource(req *RegionalAPIGetDataSourceRequest, opts ...scw.RequestOption) (*DataSource, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DataSourceID) == "" {
		return nil, errors.New("field DataSourceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/data-sources/" + fmt.Sprint(req.DataSourceID) + "",
	}

	var resp DataSource

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDataSource: Delete a given data source, specified by the data source ID. Note that deleting a data source is irreversible, and cannot be undone.
func (s *RegionalAPI) DeleteDataSource(req *RegionalAPIDeleteDataSourceRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DataSourceID) == "" {
		return errors.New("field DataSourceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/data-sources/" + fmt.Sprint(req.DataSourceID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListDataSources: Retrieve the list of data sources available in the specified region. By default, the data sources returned in the list are ordered by creation date, in ascending order.
// You can list data sources by Project, type and origin.
func (s *RegionalAPI) ListDataSources(req *RegionalAPIListDataSourcesRequest, opts ...scw.RequestOption) (*ListDataSourcesResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

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
	parameter.AddToQuery(query, "origin", req.Origin)
	parameter.AddToQuery(query, "types", req.Types)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/data-sources",
		Query:  query,
	}

	var resp ListDataSourcesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDataSource: Update a given data source name, specified by the data source ID.
func (s *RegionalAPI) UpdateDataSource(req *RegionalAPIUpdateDataSourceRequest, opts ...scw.RequestOption) (*DataSource, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DataSourceID) == "" {
		return nil, errors.New("field DataSourceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/data-sources/" + fmt.Sprint(req.DataSourceID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DataSource

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUsageOverview: Retrieve the data source usage overview per type for the specified Project.
func (s *RegionalAPI) GetUsageOverview(req *RegionalAPIGetUsageOverviewRequest, opts ...scw.RequestOption) (*UsageOverview, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "interval", req.Interval)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/usage-overview",
		Query:  query,
	}

	var resp UsageOverview

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateToken: Give your token the relevant scopes to ensure it has the right permissions to interact with your data sources and the Alert manager. Make sure that you create your token in the same regions as the data sources you want to use it for.
// Upon creation, your token's secret key display only once. Make sure that you save it.
func (s *RegionalAPI) CreateToken(req *RegionalAPICreateTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/tokens",
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

// ListTokens: Retrieve a list of all tokens in the specified region. By default, tokens returned in the list are ordered by creation date, in ascending order.
// You can filter tokens by Project ID and token scopes.
func (s *RegionalAPI) ListTokens(req *RegionalAPIListTokensRequest, opts ...scw.RequestOption) (*ListTokensResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

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
	parameter.AddToQuery(query, "token_scopes", req.TokenScopes)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/tokens",
		Query:  query,
	}

	var resp ListTokensResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetToken: Retrieve information about a given token, specified by the token ID. The token's information such as its scopes, is returned.
func (s *RegionalAPI) GetToken(req *RegionalAPIGetTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TokenID) == "" {
		return nil, errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	var resp Token

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteToken: Delete a given token, specified by the token ID. Deleting a token is irreversible and cannot be undone.
func (s *RegionalAPI) DeleteToken(req *RegionalAPIDeleteTokenRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TokenID) == "" {
		return errors.New("field TokenID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/tokens/" + fmt.Sprint(req.TokenID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetAlertManager: Retrieve information about the Alert manager which is unique per Project and region. By default the Alert manager is disabled.
// The output returned displays a URL to access the Alert manager, and whether the Alert manager and managed alerts are enabled.
func (s *RegionalAPI) GetAlertManager(req *RegionalAPIGetAlertManagerRequest, opts ...scw.RequestOption) (*AlertManager, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alert-manager",
		Query:  query,
	}

	var resp AlertManager

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableAlertManager: Enabling the Alert manager allows you to enable managed alerts and create contact points in the specified Project and region, to be notified when your Scaleway resources may require your attention.
func (s *RegionalAPI) EnableAlertManager(req *RegionalAPIEnableAlertManagerRequest, opts ...scw.RequestOption) (*AlertManager, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alert-manager/enable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AlertManager

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableAlertManager: Disabling the Alert manager deletes the contact points you have created and disables managed alerts in the specified Project and region.
func (s *RegionalAPI) DisableAlertManager(req *RegionalAPIDisableAlertManagerRequest, opts ...scw.RequestOption) (*AlertManager, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alert-manager/disable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AlertManager

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateContactPoint: Contact points are email addresses associated with the default receiver, that the Alert manager sends alerts to.
// The source of the alerts are data sources within the same Project and region as the Alert manager.
// If you need to receive alerts for other receivers, you can create additional contact points and receivers in Grafana. Make sure that you select the Scaleway Alert manager.
func (s *RegionalAPI) CreateContactPoint(req *RegionalAPICreateContactPointRequest, opts ...scw.RequestOption) (*ContactPoint, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alert-manager/contact-points",
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

// ListContactPoints: Retrieve a list of contact points for the specified Project. The response lists all contact points and receivers created in Grafana or via the API.
func (s *RegionalAPI) ListContactPoints(req *RegionalAPIListContactPointsRequest, opts ...scw.RequestOption) (*ListContactPointsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alert-manager/contact-points",
		Query:  query,
	}

	var resp ListContactPointsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateContactPoint:
func (s *RegionalAPI) UpdateContactPoint(req *RegionalAPIUpdateContactPointRequest, opts ...scw.RequestOption) (*ContactPoint, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alert-manager/contact-points",
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

// DeleteContactPoint: Delete a contact point associated with the default receiver.
func (s *RegionalAPI) DeleteContactPoint(req *RegionalAPIDeleteContactPointRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alert-manager/contact-points/delete",
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

// ListManagedAlerts: List all managed alerts for the specified Project.
func (s *RegionalAPI) ListManagedAlerts(req *RegionalAPIListManagedAlertsRequest, opts ...scw.RequestOption) (*ListManagedAlertsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/managed-alerts",
		Query:  query,
	}

	var resp ListManagedAlertsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableManagedAlerts: Enable the sending of managed alerts for the specified Project. Managed alerts are predefined alerts that apply to Scaleway recources integrated with Cockpit by default.
func (s *RegionalAPI) EnableManagedAlerts(req *RegionalAPIEnableManagedAlertsRequest, opts ...scw.RequestOption) (*AlertManager, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alert-manager/managed-alerts/enable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AlertManager

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableManagedAlerts: Disable the sending of managed alerts for the specified Project.
func (s *RegionalAPI) DisableManagedAlerts(req *RegionalAPIDisableManagedAlertsRequest, opts ...scw.RequestOption) (*AlertManager, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alert-manager/managed-alerts/disable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AlertManager

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// TriggerTestAlert: Send a test alert to the Alert manager to make sure your contact points get notified.
func (s *RegionalAPI) TriggerTestAlert(req *RegionalAPITriggerTestAlertRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alert-manager/trigger-test-alert",
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

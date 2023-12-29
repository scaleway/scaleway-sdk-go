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

type DataSourceOrigin string

const (
	// Unknown data source origin.
	DataSourceOriginUnknownOrigin = DataSourceOrigin("unknown_origin")
	// Data source which received data from Scaleway Products.
	DataSourceOriginScaleway = DataSourceOrigin("scaleway")
	// Data source which received data from external sources.
	DataSourceOriginExternal = DataSourceOrigin("external")
)

func (enum DataSourceOrigin) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_origin"
	}
	return string(enum)
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
	// Metrics data source.
	DataSourceTypeMetrics = DataSourceType("metrics")
	// Logs data source.
	DataSourceTypeLogs = DataSourceType("logs")
	// Traces data source.
	DataSourceTypeTraces = DataSourceType("traces")
)

func (enum DataSourceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
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

type ListPlanTypesRequestOrderBy string

const (
	ListPlanTypesRequestOrderByNameAsc  = ListPlanTypesRequestOrderBy("name_asc")
	ListPlanTypesRequestOrderByNameDesc = ListPlanTypesRequestOrderBy("name_desc")
)

func (enum ListPlanTypesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListPlanTypesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPlanTypesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPlanTypesRequestOrderBy(ListPlanTypesRequestOrderBy(tmp).String())
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

type ListUsagesRequestOrderBy string

const (
	ListUsagesRequestOrderByCreatedAtAsc  = ListUsagesRequestOrderBy("created_at_asc")
	ListUsagesRequestOrderByCreatedAtDesc = ListUsagesRequestOrderBy("created_at_desc")
)

func (enum ListUsagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListUsagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListUsagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListUsagesRequestOrderBy(ListUsagesRequestOrderBy(tmp).String())
	return nil
}

type PlanTypeName string

const (
	PlanTypeNameUnknownName = PlanTypeName("unknown_name")
	PlanTypeNameFree        = PlanTypeName("free")
	PlanTypeNamePremium     = PlanTypeName("premium")
	PlanTypeNameCustom      = PlanTypeName("custom")
)

func (enum PlanTypeName) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_name"
	}
	return string(enum)
}

func (enum PlanTypeName) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PlanTypeName) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PlanTypeName(PlanTypeName(tmp).String())
	return nil
}

type TokenScope string

const (
	// Unknown token scope.
	TokenScopeUnknownScope = TokenScope("unknown_scope")
	// Permission to read data from metrics data sources.
	TokenScopeReadOnlyMetrics = TokenScope("read_only_metrics")
	// Permission to write data to metrics data sources.
	TokenScopeWriteOnlyMetrics = TokenScope("write_only_metrics")
	// Permission to setup prometheus rules to metrics data sources.
	TokenScopeFullAccessMetricsRules = TokenScope("full_access_metrics_rules")
	// Permission to read data from logs data sources.
	TokenScopeReadOnlyLogs = TokenScope("read_only_logs")
	// Permission to write data to logs data sources.
	TokenScopeWriteOnlyLogs = TokenScope("write_only_logs")
	// Permission to setup prometheus rules to logs data sources.
	TokenScopeFullAccessLogsRules = TokenScope("full_access_logs_rules")
	// Permission to configure the alert manager.
	TokenScopeFullAccessAlertManager = TokenScope("full_access_alert_manager")
	// Permission to read data from traces data sources.
	TokenScopeReadOnlyTraces = TokenScope("read_only_traces")
	// Permission to write data to traces data sources.
	TokenScopeWriteOnlyTraces = TokenScope("write_only_traces")
)

func (enum TokenScope) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_scope"
	}
	return string(enum)
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

type UsageInterval string

const (
	UsageIntervalUnknownInterval = UsageInterval("unknown_interval")
	UsageIntervalCurrentMonth    = UsageInterval("current_month")
	UsageIntervalThirtyOneDays   = UsageInterval("thirty_one_days")
	UsageIntervalRetentionPeriod = UsageInterval("retention_period")
)

func (enum UsageInterval) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_interval"
	}
	return string(enum)
}

func (enum UsageInterval) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UsageInterval) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UsageInterval(UsageInterval(tmp).String())
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

// ContactPoint: Contact point.
type ContactPoint struct {
	// Email: contact point configuration.
	// Precisely one of Email must be set.
	Email *ContactPointEmail `json:"email,omitempty"`

	// Region: contact point region.
	Region scw.Region `json:"region"`
}

// DataSource: Datasource.
type DataSource struct {
	// ID: ID of the data source.
	ID string `json:"id"`

	// ProjectID: ID of the Project the Cockpit belongs to.
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

	// CreatedAt: creation date of the data source.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date of the data source.
	UpdatedAt *time.Time `json:"updated_at"`

	// Region: data source region.
	Region scw.Region `json:"region"`
}

// GrafanaProductDashboard: Grafana dashboard.
type GrafanaProductDashboard struct {
	// Name: name of the dashboard.
	Name string `json:"name"`

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

	// Password: password of the Grafana user.
	Password *string `json:"password"`
}

// PlanType: Type of pricing plan.
type PlanType struct {
	// ID: ID of a given pricing plan.
	ID string `json:"id"`

	// Name: name of a given pricing plan.
	// Default value: unknown_name
	Name PlanTypeName `json:"name"`

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

// Plan: Matching pricing plan with project.
type Plan struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`

	// PlanTypeID: ID of the pricing plan.
	PlanTypeID string `json:"plan_type_id"`
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
	Scopes []TokenScope `json:"scopes"`

	// SecretKey: token's secret key returned only at token creation.
	SecretKey *string `json:"secret_key"`

	// Region: token's region.
	Region scw.Region `json:"region"`
}

// Usage: usage.
type Usage struct {
	DataSourceID string `json:"data_source_id"`

	ProjectID string `json:"project_id"`

	// DataSourceType: default value: unknown_type
	DataSourceType DataSourceType `json:"data_source_type"`

	// DataSourceOrigin: default value: unknown_origin
	DataSourceOrigin DataSourceOrigin `json:"data_source_origin"`

	// Unit: default value: unknown_unit
	Unit UsageUnit `json:"unit"`

	// Interval: default value: unknown_interval
	Interval UsageInterval `json:"interval"`

	QuantityOverInterval uint64 `json:"quantity_over_interval"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// AlertManager: Alert manager.
type AlertManager struct {
	// AlertManagerURL: alert manager URL.
	AlertManagerURL string `json:"alert_manager_url"`

	// AlertManagerEnabled: specifies whether the alert manager is enabled.
	AlertManagerEnabled bool `json:"alert_manager_enabled"`

	// ManagedAlertsEnabled: specifies whether managed alerts are enabled.
	ManagedAlertsEnabled bool `json:"managed_alerts_enabled"`

	// Region: alert manager region.
	Region scw.Region `json:"region"`
}

// GlobalAPICreateGrafanaUserRequest: Request to create a Grafana user.
type GlobalAPICreateGrafanaUserRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`

	// Login: username of the Grafana user.
	Login string `json:"login"`

	// Role: role assigned to the Grafana user.
	// Default value: unknown_role
	Role GrafanaUserRole `json:"role"`
}

// GlobalAPIDeleteGrafanaUserRequest: Request to delete a Grafana user.
type GlobalAPIDeleteGrafanaUserRequest struct {
	// GrafanaUserID: ID of the Grafana user.
	GrafanaUserID uint32 `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// GlobalAPIGetGrafanaProductDashboardRequest: Request to get a dashboard.
type GlobalAPIGetGrafanaProductDashboardRequest struct {
	// DashboardName: name of the dashboard.
	DashboardName string `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// GlobalAPIGetGrafanaRequest: Request to get a Grafana instance.
type GlobalAPIGetGrafanaRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// GlobalAPIListGrafanaProductDashboardsRequest: Request to get a list of dashboards.
type GlobalAPIListGrafanaProductDashboardsRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// Tags: tags to filter the dashboards.
	Tags []string `json:"-"`
}

// GlobalAPIListGrafanaUsersRequest: Request to list all Grafana users.
type GlobalAPIListGrafanaUsersRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: order by.
	// Default value: login_asc
	OrderBy ListGrafanaUsersRequestOrderBy `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"-"`
}

// GlobalAPIListPlanTypesRequest: Request to list all pricing plans.
type GlobalAPIListPlanTypesRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListPlanTypesRequestOrderBy `json:"-"`
}

// GlobalAPIListPlansRequest: Request to list all pricing plans.
type GlobalAPIListPlansRequest struct {
	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: order by.
	// Default value: name_asc
	OrderBy ListPlansRequestOrderBy `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID *string `json:"-"`
}

// GlobalAPIResetGrafanaUserPasswordRequest: Request to reset a Grafana user's password.
type GlobalAPIResetGrafanaUserPasswordRequest struct {
	// GrafanaUserID: ID of the Grafana user.
	GrafanaUserID uint32 `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// GlobalAPISelectPlanRequest: Request to select a specific pricing plan.
type GlobalAPISelectPlanRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`

	// PlanTypeID: ID of the pricing plan.
	PlanTypeID string `json:"plan_type_id"`
}

// GlobalAPISyncGrafanaDataSourcesRequest: Request to synchronize all data-sources created in the Regional API.
type GlobalAPISyncGrafanaDataSourcesRequest struct {
	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// Grafana: Grafana user.
type Grafana struct {
	// GrafanaURL: URL of the Grafana instance.
	GrafanaURL string `json:"grafana_url"`
}

// ListContactPointsResponse: Response returned when listing contact points.
type ListContactPointsResponse struct {
	// TotalCount: count of all contact points created.
	TotalCount uint64 `json:"total_count"`

	// ContactPoints: array of contact points.
	ContactPoints []*ContactPoint `json:"contact_points"`

	// HasAdditionalReceivers: specifies whether the contact point has other receivers than the default receiver.
	HasAdditionalReceivers bool `json:"has_additional_receivers"`

	// HasAdditionalContactPoints: specifies whether there are unmanaged contact points.
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
	// TotalCount: count of all datasources corresponding to the request.
	TotalCount uint64 `json:"total_count"`

	// DataSources: list of the datasources within the pagination.
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

// ListGrafanaProductDashboardsResponse: Response returned when getting a list of dashboards.
type ListGrafanaProductDashboardsResponse struct {
	// TotalCount: count of grafana dashboards.
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

// ListGrafanaUsersResponse: Response returned when listing Grafana users.
type ListGrafanaUsersResponse struct {
	// TotalCount: count of all Grafana users.
	TotalCount uint64 `json:"total_count"`

	// GrafanaUsers: information on all Grafana users.
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

// ListPlanTypesResponse: Response returned when listing all pricing plans.
type ListPlanTypesResponse struct {
	// TotalCount: count of all pricing plans.
	TotalCount uint64 `json:"total_count"`

	// PlanTypes: information on plan types.
	PlanTypes []*PlanType `json:"plan_types"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPlanTypesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPlanTypesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPlanTypesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PlanTypes = append(r.PlanTypes, results.PlanTypes...)
	r.TotalCount += uint64(len(results.PlanTypes))
	return uint64(len(results.PlanTypes)), nil
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

// ListTokensResponse: list tokens response.
type ListTokensResponse struct {
	// TotalCount: count of all tokens created.
	TotalCount uint64 `json:"total_count"`

	// Tokens: list of all tokens created.
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

// ListUsagesResponse: list usages response.
type ListUsagesResponse struct {
	TotalCount uint64 `json:"total_count"`

	Usages []*Usage `json:"usages"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListUsagesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListUsagesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListUsagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Usages = append(r.Usages, results.Usages...)
	r.TotalCount += uint64(len(results.Usages))
	return uint64(len(results.Usages)), nil
}

// RegionalAPICreateContactPointRequest: Request to create a contact point.
type RegionalAPICreateContactPointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project in which to create the contact point.
	ProjectID string `json:"project_id"`

	// Email: email address of the contact point.
	// Precisely one of Email must be set.
	Email *ContactPointEmail `json:"email,omitempty"`
}

// RegionalAPICreateDataSourceRequest: Request to create a data source.
type RegionalAPICreateDataSourceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project the Cockpit belongs to.
	ProjectID string `json:"project_id"`

	// Name: data source name.
	Name string `json:"name"`

	// Type: data source type.
	// Default value: unknown_type
	Type DataSourceType `json:"type"`
}

// RegionalAPICreateTokenRequest: regional api create token request.
type RegionalAPICreateTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`

	// Name: name of the token.
	Name string `json:"name"`

	// TokenScopes: token's permission scopes.
	TokenScopes []TokenScope `json:"token_scopes"`
}

// RegionalAPIDeleteContactPointRequest: Request to delete a contact point.
type RegionalAPIDeleteContactPointRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`

	// ContactPoint: contact point to delete.
	ContactPoint *ContactPoint `json:"contact_point,omitempty"`
}

// RegionalAPIDeleteDataSourceRequest: Request to delete a data source.
type RegionalAPIDeleteDataSourceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DataSourceID: ID of the data source.
	DataSourceID string `json:"-"`
}

// RegionalAPIDeleteTokenRequest: regional api delete token request.
type RegionalAPIDeleteTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TokenID: ID of the token.
	TokenID string `json:"-"`
}

// RegionalAPIDisableAlertManagerRequest: Request to disable the alert manager.
type RegionalAPIDisableAlertManagerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// RegionalAPIDisableManagedAlertsRequest: Request to disable the sending of managed alerts.
type RegionalAPIDisableManagedAlertsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// RegionalAPIEnableAlertManagerRequest: Request to enable the alert manager.
type RegionalAPIEnableAlertManagerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// RegionalAPIEnableManagedAlertsRequest: Request to enable the sending of managed alerts.
type RegionalAPIEnableManagedAlertsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// RegionalAPIGetAlertManagerRequest: Request to get the alert manager.
type RegionalAPIGetAlertManagerRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// RegionalAPIGetDataSourceRequest: Request to get a data source.
type RegionalAPIGetDataSourceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DataSourceID: ID of the data source.
	DataSourceID string `json:"-"`
}

// RegionalAPIGetTokenRequest: regional api get token request.
type RegionalAPIGetTokenRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TokenID: ID of the token.
	TokenID string `json:"-"`
}

// RegionalAPIListContactPointsRequest: Request to list all contact points.
type RegionalAPIListContactPointsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// ProjectID: ID of the Project from which to list the contact points.
	ProjectID string `json:"-"`
}

// RegionalAPIListDataSourcesRequest: Request to list all data sources.
type RegionalAPIListDataSourcesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the response is ordered.
	// Default value: created_at_asc
	OrderBy ListDataSourcesRequestOrderBy `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID *string `json:"-"`

	// Origins: filter by data source origins.
	Origins []DataSourceOrigin `json:"-"`

	// Types: filter by data source types.
	Types []DataSourceType `json:"-"`
}

// RegionalAPIListTokensRequest: regional api list tokens request.
type RegionalAPIListTokensRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the response is ordered.
	// Default value: created_at_asc
	OrderBy ListTokensRequestOrderBy `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID *string `json:"-"`

	// TokenScopes: filter by token scopes.
	TokenScopes []TokenScope `json:"-"`
}

// RegionalAPIListUsagesRequest: regional api list usages request.
type RegionalAPIListUsagesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListUsagesRequestOrderBy `json:"-"`

	ProjectID string `json:"-"`

	DataSourceIDs []string `json:"-"`

	Origins []DataSourceOrigin `json:"-"`

	Types []DataSourceType `json:"-"`

	// Interval: default value: unknown_interval
	Interval UsageInterval `json:"-"`
}

// RegionalAPITriggerTestAlertRequest: Request to trigger a test alert.
type RegionalAPITriggerTestAlertRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`
}

// The Global Cockpit API allows you to manage your Cockpit instance.
type GlobalAPI struct {
	client *scw.Client
}

// NewGlobalAPI returns a GlobalAPI object from a Scaleway client.
func NewGlobalAPI(client *scw.Client) *GlobalAPI {
	return &GlobalAPI{
		client: client,
	}
}
func (s *GlobalAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// GetGrafana: Retrieve the Cockpit of the specified Project ID.
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

// SyncGrafanaDataSources: Synchronize all data sources created in the Regional API.
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

// CreateGrafanaUser: Create a Grafana user for your Cockpit's Grafana instance. Make sure you save the automatically-generated password and the Grafana user ID.
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

// ListGrafanaUsers: Get a list of Grafana users who are able to connect to the Cockpit's Grafana instance.
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

// DeleteGrafanaUser: Delete a Grafana user from a Grafana instance, specified by the Cockpit's Project ID and the Grafana user ID.
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

// ResetGrafanaUserPassword: Reset a Grafana user's password specified by the Cockpit's Project ID and the Grafana user ID.
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

// ListGrafanaProductDashboards: Get a list of available product dashboards.
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

// GetGrafanaProductDashboard: Get a product dashboard specified by the dashboard ID.
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

// ListPlanTypes: Get a list of all pricing plans available.
func (s *GlobalAPI) ListPlanTypes(req *GlobalAPIListPlanTypesRequest, opts ...scw.RequestOption) (*ListPlanTypesResponse, error) {
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
		Path:   "/cockpit/v1/plan-types",
		Query:  query,
	}

	var resp ListPlanTypesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SelectPlan: Select your chosen pricing plan for your Cockpit, specifying the Cockpit's Project ID and the pricing plan's ID in the request.
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

// ListPlans: Get a list of all pricing plans available.
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
	parameter.AddToQuery(query, "project_id", req.ProjectID)

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

// Regional Cockpit's API allows you to create datasources and tokens to store and query Metrics, Logs & Traces. Cockpit comes with a regional Alertmanager as well to send Alerts to your contact points. Scaleway Products which are integrated with Cockpit send Metrics and Logs for your Projects and propose managed alerts.
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

// CreateDataSource: Create a data source for the specified Project ID and the given type.
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

// GetDataSource: Retrieve the data source associated with the specified datasource ID.
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

// DeleteDataSource: Delete the data source associated with the specified datasource ID.
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

// ListDataSources: Get a list of data sources for the specified Project ID.
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

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "origins", req.Origins)
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

// ListUsages: Get a list of data sources for the specified Project ID.
func (s *RegionalAPI) ListUsages(req *RegionalAPIListUsagesRequest, opts ...scw.RequestOption) (*ListUsagesResponse, error) {
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
	parameter.AddToQuery(query, "data_source_ids", req.DataSourceIDs)
	parameter.AddToQuery(query, "origins", req.Origins)
	parameter.AddToQuery(query, "types", req.Types)
	parameter.AddToQuery(query, "interval", req.Interval)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/usages",
		Query:  query,
	}

	var resp ListUsagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateToken: Create a token associated with the specified Project ID.
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

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("token")
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

// ListTokens: Get a list of tokens associated with the specified Project ID.
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

// GetToken: Retrieve the token associated with the specified token ID.
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

// DeleteToken: Delete the token associated with the specified token ID.
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

// GetAlertManager: Create a contact point to receive alerts for the default receiver.
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

// EnableAlertManager: Create a contact point to receive alerts for the default receiver.
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

// DisableAlertManager: Create a contact point to receive alerts for the default receiver.
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

// CreateContactPoint: Create a contact point to receive alerts for the default receiver.
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

// ListContactPoints: Get a list of contact points for the Cockpit associated with the specified Project ID.
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

// DeleteContactPoint: Delete a contact point for the default receiver.
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

// EnableManagedAlerts: Enable the sending of managed alerts for the specified Project's Cockpit.
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

// DisableManagedAlerts: Disable the sending of managed alerts for the specified Project's Cockpit.
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

// TriggerTestAlert: Trigger a test alert to all of the Cockpit's receivers.
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

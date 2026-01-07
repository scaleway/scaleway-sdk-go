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

type AlertState string

const (
	AlertStateUnknownState = AlertState("unknown_state")
	// The alert is inactive and may transition to `pending` or `firing` if its conditions are met.
	AlertStateInactive = AlertState("inactive")
	// The alert's conditions are met. They must persist for the configured duration before transitioning to the `firing` state.
	AlertStatePending = AlertState("pending")
	// The alert's conditions, including the required duration, have been fully met.
	AlertStateFiring = AlertState("firing")
)

func (enum AlertState) String() string {
	if enum == "" {
		// return default value if empty
		return string(AlertStateUnknownState)
	}
	return string(enum)
}

func (enum AlertState) Values() []AlertState {
	return []AlertState{
		"unknown_state",
		"inactive",
		"pending",
		"firing",
	}
}

func (enum AlertState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AlertState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AlertState(AlertState(tmp).String())
	return nil
}

type AlertStatus string

const (
	AlertStatusUnknownStatus = AlertStatus("unknown_status")
	// The alert is enabled and may trigger based on its conditions.
	AlertStatusEnabled = AlertStatus("enabled")
	// The alert is disabled. It will never trigger, and will not be evaluated.
	AlertStatusDisabled = AlertStatus("disabled")
	// The alert has been marked for activation. It will be enabled momentarily.
	AlertStatusEnabling = AlertStatus("enabling")
	// The alert has been marked for deactivation. It will be disabled momentarily.
	AlertStatusDisabling = AlertStatus("disabling")
)

func (enum AlertStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(AlertStatusUnknownStatus)
	}
	return string(enum)
}

func (enum AlertStatus) Values() []AlertStatus {
	return []AlertStatus{
		"unknown_status",
		"enabled",
		"disabled",
		"enabling",
		"disabling",
	}
}

func (enum AlertStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AlertStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AlertStatus(AlertStatus(tmp).String())
	return nil
}

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
		return string(DataSourceOriginUnknownOrigin)
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
		return string(DataSourceTypeUnknownType)
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
		return string(GrafanaUserRoleUnknownRole)
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
		return string(ListDataSourcesRequestOrderByCreatedAtAsc)
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
		return string(ListGrafanaUsersRequestOrderByLoginAsc)
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

type ListPlansRequestOrderBy string

const (
	ListPlansRequestOrderByNameAsc  = ListPlansRequestOrderBy("name_asc")
	ListPlansRequestOrderByNameDesc = ListPlansRequestOrderBy("name_desc")
)

func (enum ListPlansRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPlansRequestOrderByNameAsc)
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

type ListProductsRequestOrderBy string

const (
	ListProductsRequestOrderByCreatedAtAsc    = ListProductsRequestOrderBy("created_at_asc")
	ListProductsRequestOrderByCreatedAtDesc   = ListProductsRequestOrderBy("created_at_desc")
	ListProductsRequestOrderByDisplayNameAsc  = ListProductsRequestOrderBy("display_name_asc")
	ListProductsRequestOrderByDisplayNameDesc = ListProductsRequestOrderBy("display_name_desc")
	ListProductsRequestOrderByFamilyNameAsc   = ListProductsRequestOrderBy("family_name_asc")
	ListProductsRequestOrderByFamilyNameDesc  = ListProductsRequestOrderBy("family_name_desc")
)

func (enum ListProductsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListProductsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListProductsRequestOrderBy) Values() []ListProductsRequestOrderBy {
	return []ListProductsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"display_name_asc",
		"display_name_desc",
		"family_name_asc",
		"family_name_desc",
	}
}

func (enum ListProductsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListProductsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListProductsRequestOrderBy(ListProductsRequestOrderBy(tmp).String())
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
		return string(ListTokensRequestOrderByCreatedAtAsc)
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
		return string(PlanNameUnknownName)
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
		return string(TokenScopeUnknownScope)
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
		return string(UsageUnitUnknownUnit)
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

// PreconfiguredAlertData: Structure for additional data relative to preconfigured alerts.
type PreconfiguredAlertData struct {
	// PreconfiguredRuleID: ID of the preconfigured rule if the alert is preconfigured.
	PreconfiguredRuleID string `json:"preconfigured_rule_id"`

	// DisplayName: human readable name of the alert.
	DisplayName string `json:"display_name"`

	// DisplayDescription: human readable description of the alert.
	DisplayDescription string `json:"display_description"`

	// ProductName: product associated with the alert.
	ProductName string `json:"product_name"`

	// ProductFamily: family of the product associated with the alert.
	ProductFamily string `json:"product_family"`
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

// RulesCount: rules count.
type RulesCount struct {
	// DataSourceID: ID of the data source.
	DataSourceID string `json:"data_source_id"`

	// DataSourceName: name of the data source.
	DataSourceName string `json:"data_source_name"`

	// RulesCount: total count of rules associated with this data source.
	RulesCount int32 `json:"rules_count"`
}

// Alert: Structure representing an alert.
type Alert struct {
	// Region: the region in which the alert is defined.
	Region scw.Region `json:"region"`

	// Preconfigured: indicates if the alert is preconfigured or custom.
	Preconfigured bool `json:"preconfigured"`

	// Name: name of the alert.
	Name string `json:"name"`

	// Rule: rule defining the alert condition.
	Rule string `json:"rule"`

	// Duration: duration for which the alert must be active before firing. The format of this duration follows the prometheus duration format.
	Duration string `json:"duration"`

	// RuleStatus: indicates if the alert is enabled, enabling, disabled or disabling. Preconfigured alerts can have any of these values, whereas custom alerts can only have the status "enabled".
	// Default value: unknown_status
	RuleStatus AlertStatus `json:"rule_status"`

	// State: current state of the alert. Possible states are `inactive`, `pending`, and `firing`.
	// Default value: unknown_state
	State *AlertState `json:"state"`

	// Annotations: annotations for the alert, used to provide additional information about the alert.
	Annotations map[string]string `json:"annotations"`

	// PreconfiguredData: contains additional data for preconfigured alerts, such as the rule ID, display name, and description. Only present if the alert is preconfigured.
	PreconfiguredData *PreconfiguredAlertData `json:"preconfigured_data"`

	// DataSourceID: ID of the data source containing the alert rule.
	DataSourceID string `json:"data_source_id"`
}

// ContactPoint: Contact point.
type ContactPoint struct {
	// Email: email address to send alerts to.
	// Precisely one of Email must be set.
	Email *ContactPointEmail `json:"email,omitempty"`

	// Region: region.
	Region scw.Region `json:"region"`

	// SendResolvedNotifications: send an email notification when an alert is marked as resolved.
	SendResolvedNotifications bool `json:"send_resolved_notifications"`
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

	// RetentionDays: duration for which the data will be retained in the data source.
	RetentionDays uint32 `json:"retention_days"`

	// Region: region of the data source.
	Region scw.Region `json:"region"`

	// CurrentMonthUsage: usage of the month in bytes.
	CurrentMonthUsage *scw.Size `json:"current_month_usage"`
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

// Product: product.
type Product struct {
	Name string `json:"name"`

	DisplayName string `json:"display_name"`

	FamilyName string `json:"family_name"`

	ResourceTypes []string `json:"resource_types"`
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

// DisableAlertRulesResponse: Output returned when alert rules are disabled.
type DisableAlertRulesResponse struct {
	// DisabledRuleIDs: only newly disabled rules are listed. Rules that were already disabled are not returned in the output.
	DisabledRuleIDs []string `json:"disabled_rule_ids"`
}

// EnableAlertRulesResponse: Output returned when alert rules are enabled.
type EnableAlertRulesResponse struct {
	// EnabledRuleIDs: only newly enabled rules are listed. Rules that were already enabled are not returned in the output.
	EnabledRuleIDs []string `json:"enabled_rule_ids"`
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

// GetRulesCountResponse: get rules count response.
type GetRulesCountResponse struct {
	// RulesCountByDatasource: total count of rules grouped by data source.
	RulesCountByDatasource []*RulesCount `json:"rules_count_by_datasource"`

	// PreconfiguredRulesCount: total count of preconfigured rules.
	PreconfiguredRulesCount int32 `json:"preconfigured_rules_count"`

	// CustomRulesCount: total count of custom rules.
	CustomRulesCount int32 `json:"custom_rules_count"`
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

// ListAlertsResponse: Retrieve a list of alerts matching the request.
type ListAlertsResponse struct {
	// TotalCount: total count of alerts matching the request.
	TotalCount uint64 `json:"total_count"`

	// Alerts: list of alerts matching the applied filters.
	Alerts []*Alert `json:"alerts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListAlertsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListAlertsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListAlertsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Alerts = append(r.Alerts, results.Alerts...)
	r.TotalCount += uint64(len(results.Alerts))
	return uint64(len(results.Alerts)), nil
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
func (r *ListContactPointsResponse) UnsafeAppend(res any) (uint64, error) {
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
func (r *ListDataSourcesResponse) UnsafeAppend(res any) (uint64, error) {
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
func (r *ListGrafanaProductDashboardsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListGrafanaProductDashboardsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Dashboards = append(r.Dashboards, results.Dashboards...)
	r.TotalCount += uint64(len(results.Dashboards))
	return uint64(len(results.Dashboards)), nil
}

// ListGrafanaUsersResponse: Output returned when listing Grafana users.
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
func (r *ListGrafanaUsersResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListGrafanaUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.GrafanaUsers = append(r.GrafanaUsers, results.GrafanaUsers...)
	r.TotalCount += uint64(len(results.GrafanaUsers))
	return uint64(len(results.GrafanaUsers)), nil
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
func (r *ListPlansResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListPlansResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Plans = append(r.Plans, results.Plans...)
	r.TotalCount += uint64(len(results.Plans))
	return uint64(len(results.Plans)), nil
}

// ListProductsResponse: list products response.
type ListProductsResponse struct {
	ProductsList []*Product `json:"products_list"`

	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListProductsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListProductsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListProductsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ProductsList = append(r.ProductsList, results.ProductsList...)
	r.TotalCount += uint64(len(results.ProductsList))
	return uint64(len(results.ProductsList)), nil
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
func (r *ListTokensResponse) UnsafeAppend(res any) (uint64, error) {
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

	// SendResolvedNotifications: send an email notification when an alert is marked as resolved.
	SendResolvedNotifications *bool `json:"send_resolved_notifications,omitempty"`
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

	// RetentionDays: default values are 31 days for metrics, 7 days for logs and traces.
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

// RegionalAPIDisableAlertRulesRequest: regional api disable alert rules request.
type RegionalAPIDisableAlertRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`

	// RuleIDs: list of IDs of the rules to enable. If empty, disables all preconfigured rules.
	RuleIDs []string `json:"rule_ids"`
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

// RegionalAPIEnableAlertRulesRequest: regional api enable alert rules request.
type RegionalAPIEnableAlertRulesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project.
	ProjectID string `json:"project_id"`

	// RuleIDs: list of IDs of the rules to enable. If empty, enables all preconfigured rules.
	RuleIDs []string `json:"rule_ids"`
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

// RegionalAPIGetRulesCountRequest: regional api get rules count request.
type RegionalAPIGetRulesCountRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project to retrieve the rule count for.
	ProjectID string `json:"project_id"`
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

// RegionalAPIListAlertsRequest: Retrieve a list of alerts.
type RegionalAPIListAlertsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project ID to filter for, only alerts from this Project will be returned.
	ProjectID string `json:"-"`

	// RuleStatus: returns only alerts with the given activation status. If omitted, no alert filtering is applied. Other filters may still apply.
	// Default value: unknown_status
	RuleStatus *AlertStatus `json:"-"`

	// IsPreconfigured: true returns only preconfigured alerts. False returns only custom alerts. If omitted, no filtering is applied on alert types. Other filters may still apply.
	IsPreconfigured *bool `json:"-"`

	// State: valid values to filter on are `inactive`, `pending` and `firing`. If omitted, no filtering is applied on alert states. Other filters may still apply.
	// Default value: unknown_state
	State *AlertState `json:"-"`

	// DataSourceID: if omitted, only alerts from the default Scaleway metrics data source will be listed.
	DataSourceID *string `json:"-"`
}

// RegionalAPIListContactPointsRequest: List contact points.
type RegionalAPIListContactPointsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project containing the contact points to list.
	ProjectID string `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: total count of contact points to return per page.
	PageSize *uint32 `json:"-"`
}

// RegionalAPIListDataSourcesRequest: List data sources.
type RegionalAPIListDataSourcesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project ID to filter for, only data sources from this Project will be returned.
	ProjectID string `json:"-"`

	// Origin: origin to filter for, only data sources with matching origin will be returned. If omitted, all types will be returned.
	// Default value: unknown_origin
	Origin DataSourceOrigin `json:"-"`

	// Types: types to filter for (metrics, logs, traces), only data sources with matching types will be returned. If omitted, all types will be returned.
	Types []DataSourceType `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of data sources to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order for data sources in the response.
	// Default value: created_at_asc
	OrderBy ListDataSourcesRequestOrderBy `json:"-"`
}

// RegionalAPIListProductsRequest: List all Scaleway products that send metrics and/or logs to Cockpit.
type RegionalAPIListProductsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of products to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order for products in the response.
	// Default value: created_at_asc
	OrderBy ListProductsRequestOrderBy `json:"-"`
}

// RegionalAPIListTokensRequest: List tokens.
type RegionalAPIListTokensRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: ID of the Project the tokens belong to.
	ProjectID string `json:"-"`

	// TokenScopes: token scopes to filter for.
	TokenScopes []TokenScope `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of tokens to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListTokensRequestOrderBy `json:"-"`
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

	// SendResolvedNotifications: enable or disable notifications when alert is resolved.
	SendResolvedNotifications *bool `json:"send_resolved_notifications,omitempty"`
}

// RegionalAPIUpdateDataSourceRequest: Update a data source name.
type RegionalAPIUpdateDataSourceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DataSourceID: ID of the data source to update.
	DataSourceID string `json:"-"`

	// Name: updated name of the data source.
	Name *string `json:"name,omitempty"`

	// RetentionDays: duration for which the data will be retained in the data source.
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

// The Cockpit Global API allows you to manage your Cockpit's Grafana.
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

// Deprecated: CreateGrafanaUser: Create a Grafana user
// Create a Grafana user to connect to your Cockpit's Grafana. Upon creation, your user password displays only once, so make sure that you save it.
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

// Deprecated: ListGrafanaUsers: List Grafana users
// List all Grafana users created in your Cockpit's Grafana. By default, the Grafana users returned in the list are ordered in ascending order.
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

// Deprecated: DeleteGrafanaUser: Delete a Grafana user
// Delete a Grafana user from your Cockpit's Grafana, specified by the ID of the Project the Cockpit belongs to, and the ID of the Grafana user.
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

// Deprecated: ResetGrafanaUserPassword: Reset a Grafana user password
// Reset the password of a Grafana user, specified by the ID of the Project the Cockpit belongs to, and the ID of the Grafana user.
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
// Deprecated due to retention now being managed at the data source level.
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
// Deprecated due to retention now being managed at the data source level.
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
// Deprecated due to retention now being managed at the data source level.
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

// The Cockpit API allows you to create data sources and Cockpit tokens to store and query data types such as metrics, logs, and traces. You can also push your data into Cockpit, and send alerts to your contact points when your resources may require your attention, using the regional Alert manager.
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

// CreateDataSource: You must specify the data source name and type (metrics, logs, traces) upon creation.
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

// DeleteDataSource: Delete a given data source. Note that this action will permanently delete this data source and any data associated with it.
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
func (s *RegionalAPI) ListDataSources(req *RegionalAPIListDataSourcesRequest, opts ...scw.RequestOption) (*ListDataSourcesResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

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
	parameter.AddToQuery(query, "origin", req.Origin)
	parameter.AddToQuery(query, "types", req.Types)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

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

// UpdateDataSource: Update a given data source attributes (name and/or retention_days).
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

// GetUsageOverview: Retrieve the volume of data ingested for each of your data sources in the specified project and region.
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
	parameter.AddToQuery(query, "token_scopes", req.TokenScopes)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

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

// ListProducts: List all Scaleway products that send metrics and/or logs to Cockpit.
// Note that all of those products send at least metrics, but only a subset send logs to Cockpit.
// For more information, see https://www.scaleway.com/en/docs/cockpit/reference-content/cockpit-product-integration/.
func (s *RegionalAPI) ListProducts(req *RegionalAPIListProductsRequest, opts ...scw.RequestOption) (*ListProductsResponse, error) {
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

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/products",
		Query:  query,
	}

	var resp ListProductsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
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

// GetRulesCount: Get a detailed count of enabled rules in the specified Project. Includes preconfigured and custom alerting and recording rules.
func (s *RegionalAPI) GetRulesCount(req *RegionalAPIGetRulesCountRequest, opts ...scw.RequestOption) (*GetRulesCountResponse, error) {
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
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/rules/count",
		Query:  query,
	}

	var resp GetRulesCountResponse

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

// ListAlerts: List preconfigured and/or custom alerts for the specified Project and data source.
func (s *RegionalAPI) ListAlerts(req *RegionalAPIListAlertsRequest, opts ...scw.RequestOption) (*ListAlertsResponse, error) {
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
	parameter.AddToQuery(query, "rule_status", req.RuleStatus)
	parameter.AddToQuery(query, "is_preconfigured", req.IsPreconfigured)
	parameter.AddToQuery(query, "state", req.State)
	parameter.AddToQuery(query, "data_source_id", req.DataSourceID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alerts",
		Query:  query,
	}

	var resp ListAlertsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: EnableManagedAlerts: Enable the sending of managed alerts for the specified Project. Managed alerts are predefined alerts that apply to Scaleway resources integrated with Cockpit by default.
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

// Deprecated: DisableManagedAlerts: Disable the sending of managed alerts for the specified Project.
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

// EnableAlertRules: Enable alert rules from the list of available preconfigured rules.
func (s *RegionalAPI) EnableAlertRules(req *RegionalAPIEnableAlertRulesRequest, opts ...scw.RequestOption) (*EnableAlertRulesResponse, error) {
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
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alert-manager/enable-alert-rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp EnableAlertRulesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableAlertRules: Disable alert rules from the list of available preconfigured rules.
func (s *RegionalAPI) DisableAlertRules(req *RegionalAPIDisableAlertRulesRequest, opts ...scw.RequestOption) (*DisableAlertRulesResponse, error) {
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
		Path:   "/cockpit/v1/regions/" + fmt.Sprint(req.Region) + "/alert-manager/disable-alert-rules",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DisableAlertRulesResponse

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

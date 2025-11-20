// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package iot provides methods and message types of the iot v1 API.
package iot

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
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultIotRetryInterval = 15 * time.Second
	defaultIotTimeout       = 15 * time.Minute
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

type DeviceMessageFiltersRulePolicy string

const (
	DeviceMessageFiltersRulePolicyUnknown = DeviceMessageFiltersRulePolicy("unknown")
	DeviceMessageFiltersRulePolicyAccept  = DeviceMessageFiltersRulePolicy("accept")
	DeviceMessageFiltersRulePolicyReject  = DeviceMessageFiltersRulePolicy("reject")
)

func (enum DeviceMessageFiltersRulePolicy) String() string {
	if enum == "" {
		// return default value if empty
		return string(DeviceMessageFiltersRulePolicyUnknown)
	}
	return string(enum)
}

func (enum DeviceMessageFiltersRulePolicy) Values() []DeviceMessageFiltersRulePolicy {
	return []DeviceMessageFiltersRulePolicy{
		"unknown",
		"accept",
		"reject",
	}
}

func (enum DeviceMessageFiltersRulePolicy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DeviceMessageFiltersRulePolicy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DeviceMessageFiltersRulePolicy(DeviceMessageFiltersRulePolicy(tmp).String())
	return nil
}

type DeviceStatus string

const (
	DeviceStatusUnknown  = DeviceStatus("unknown")
	DeviceStatusError    = DeviceStatus("error")
	DeviceStatusEnabled  = DeviceStatus("enabled")
	DeviceStatusDisabled = DeviceStatus("disabled")
)

func (enum DeviceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(DeviceStatusUnknown)
	}
	return string(enum)
}

func (enum DeviceStatus) Values() []DeviceStatus {
	return []DeviceStatus{
		"unknown",
		"error",
		"enabled",
		"disabled",
	}
}

func (enum DeviceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DeviceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DeviceStatus(DeviceStatus(tmp).String())
	return nil
}

type HubProductPlan string

const (
	HubProductPlanPlanUnknown   = HubProductPlan("plan_unknown")
	HubProductPlanPlanShared    = HubProductPlan("plan_shared")
	HubProductPlanPlanDedicated = HubProductPlan("plan_dedicated")
	HubProductPlanPlanHa        = HubProductPlan("plan_ha")
)

func (enum HubProductPlan) String() string {
	if enum == "" {
		// return default value if empty
		return string(HubProductPlanPlanUnknown)
	}
	return string(enum)
}

func (enum HubProductPlan) Values() []HubProductPlan {
	return []HubProductPlan{
		"plan_unknown",
		"plan_shared",
		"plan_dedicated",
		"plan_ha",
	}
}

func (enum HubProductPlan) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HubProductPlan) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HubProductPlan(HubProductPlan(tmp).String())
	return nil
}

type HubStatus string

const (
	HubStatusUnknown   = HubStatus("unknown")
	HubStatusError     = HubStatus("error")
	HubStatusEnabling  = HubStatus("enabling")
	HubStatusReady     = HubStatus("ready")
	HubStatusDisabling = HubStatus("disabling")
	HubStatusDisabled  = HubStatus("disabled")
)

func (enum HubStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(HubStatusUnknown)
	}
	return string(enum)
}

func (enum HubStatus) Values() []HubStatus {
	return []HubStatus{
		"unknown",
		"error",
		"enabling",
		"ready",
		"disabling",
		"disabled",
	}
}

func (enum HubStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HubStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HubStatus(HubStatus(tmp).String())
	return nil
}

type ListDevicesRequestOrderBy string

const (
	ListDevicesRequestOrderByNameAsc           = ListDevicesRequestOrderBy("name_asc")
	ListDevicesRequestOrderByNameDesc          = ListDevicesRequestOrderBy("name_desc")
	ListDevicesRequestOrderByStatusAsc         = ListDevicesRequestOrderBy("status_asc")
	ListDevicesRequestOrderByStatusDesc        = ListDevicesRequestOrderBy("status_desc")
	ListDevicesRequestOrderByHubIDAsc          = ListDevicesRequestOrderBy("hub_id_asc")
	ListDevicesRequestOrderByHubIDDesc         = ListDevicesRequestOrderBy("hub_id_desc")
	ListDevicesRequestOrderByCreatedAtAsc      = ListDevicesRequestOrderBy("created_at_asc")
	ListDevicesRequestOrderByCreatedAtDesc     = ListDevicesRequestOrderBy("created_at_desc")
	ListDevicesRequestOrderByUpdatedAtAsc      = ListDevicesRequestOrderBy("updated_at_asc")
	ListDevicesRequestOrderByUpdatedAtDesc     = ListDevicesRequestOrderBy("updated_at_desc")
	ListDevicesRequestOrderByAllowInsecureAsc  = ListDevicesRequestOrderBy("allow_insecure_asc")
	ListDevicesRequestOrderByAllowInsecureDesc = ListDevicesRequestOrderBy("allow_insecure_desc")
)

func (enum ListDevicesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListDevicesRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListDevicesRequestOrderBy) Values() []ListDevicesRequestOrderBy {
	return []ListDevicesRequestOrderBy{
		"name_asc",
		"name_desc",
		"status_asc",
		"status_desc",
		"hub_id_asc",
		"hub_id_desc",
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
		"allow_insecure_asc",
		"allow_insecure_desc",
	}
}

func (enum ListDevicesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDevicesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDevicesRequestOrderBy(ListDevicesRequestOrderBy(tmp).String())
	return nil
}

type ListHubsRequestOrderBy string

const (
	ListHubsRequestOrderByNameAsc         = ListHubsRequestOrderBy("name_asc")
	ListHubsRequestOrderByNameDesc        = ListHubsRequestOrderBy("name_desc")
	ListHubsRequestOrderByStatusAsc       = ListHubsRequestOrderBy("status_asc")
	ListHubsRequestOrderByStatusDesc      = ListHubsRequestOrderBy("status_desc")
	ListHubsRequestOrderByProductPlanAsc  = ListHubsRequestOrderBy("product_plan_asc")
	ListHubsRequestOrderByProductPlanDesc = ListHubsRequestOrderBy("product_plan_desc")
	ListHubsRequestOrderByCreatedAtAsc    = ListHubsRequestOrderBy("created_at_asc")
	ListHubsRequestOrderByCreatedAtDesc   = ListHubsRequestOrderBy("created_at_desc")
	ListHubsRequestOrderByUpdatedAtAsc    = ListHubsRequestOrderBy("updated_at_asc")
	ListHubsRequestOrderByUpdatedAtDesc   = ListHubsRequestOrderBy("updated_at_desc")
)

func (enum ListHubsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListHubsRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListHubsRequestOrderBy) Values() []ListHubsRequestOrderBy {
	return []ListHubsRequestOrderBy{
		"name_asc",
		"name_desc",
		"status_asc",
		"status_desc",
		"product_plan_asc",
		"product_plan_desc",
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
	}
}

func (enum ListHubsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListHubsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListHubsRequestOrderBy(ListHubsRequestOrderBy(tmp).String())
	return nil
}

type ListNetworksRequestOrderBy string

const (
	ListNetworksRequestOrderByNameAsc       = ListNetworksRequestOrderBy("name_asc")
	ListNetworksRequestOrderByNameDesc      = ListNetworksRequestOrderBy("name_desc")
	ListNetworksRequestOrderByTypeAsc       = ListNetworksRequestOrderBy("type_asc")
	ListNetworksRequestOrderByTypeDesc      = ListNetworksRequestOrderBy("type_desc")
	ListNetworksRequestOrderByCreatedAtAsc  = ListNetworksRequestOrderBy("created_at_asc")
	ListNetworksRequestOrderByCreatedAtDesc = ListNetworksRequestOrderBy("created_at_desc")
)

func (enum ListNetworksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListNetworksRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListNetworksRequestOrderBy) Values() []ListNetworksRequestOrderBy {
	return []ListNetworksRequestOrderBy{
		"name_asc",
		"name_desc",
		"type_asc",
		"type_desc",
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListNetworksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNetworksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNetworksRequestOrderBy(ListNetworksRequestOrderBy(tmp).String())
	return nil
}

type ListRoutesRequestOrderBy string

const (
	ListRoutesRequestOrderByNameAsc       = ListRoutesRequestOrderBy("name_asc")
	ListRoutesRequestOrderByNameDesc      = ListRoutesRequestOrderBy("name_desc")
	ListRoutesRequestOrderByHubIDAsc      = ListRoutesRequestOrderBy("hub_id_asc")
	ListRoutesRequestOrderByHubIDDesc     = ListRoutesRequestOrderBy("hub_id_desc")
	ListRoutesRequestOrderByTypeAsc       = ListRoutesRequestOrderBy("type_asc")
	ListRoutesRequestOrderByTypeDesc      = ListRoutesRequestOrderBy("type_desc")
	ListRoutesRequestOrderByCreatedAtAsc  = ListRoutesRequestOrderBy("created_at_asc")
	ListRoutesRequestOrderByCreatedAtDesc = ListRoutesRequestOrderBy("created_at_desc")
)

func (enum ListRoutesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListRoutesRequestOrderByNameAsc)
	}
	return string(enum)
}

func (enum ListRoutesRequestOrderBy) Values() []ListRoutesRequestOrderBy {
	return []ListRoutesRequestOrderBy{
		"name_asc",
		"name_desc",
		"hub_id_asc",
		"hub_id_desc",
		"type_asc",
		"type_desc",
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListRoutesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListRoutesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListRoutesRequestOrderBy(ListRoutesRequestOrderBy(tmp).String())
	return nil
}

type NetworkNetworkType string

const (
	NetworkNetworkTypeUnknown = NetworkNetworkType("unknown")
	NetworkNetworkTypeSigfox  = NetworkNetworkType("sigfox")
	NetworkNetworkTypeRest    = NetworkNetworkType("rest")
)

func (enum NetworkNetworkType) String() string {
	if enum == "" {
		// return default value if empty
		return string(NetworkNetworkTypeUnknown)
	}
	return string(enum)
}

func (enum NetworkNetworkType) Values() []NetworkNetworkType {
	return []NetworkNetworkType{
		"unknown",
		"sigfox",
		"rest",
	}
}

func (enum NetworkNetworkType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *NetworkNetworkType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = NetworkNetworkType(NetworkNetworkType(tmp).String())
	return nil
}

type RouteDatabaseConfigEngine string

const (
	RouteDatabaseConfigEngineUnknown    = RouteDatabaseConfigEngine("unknown")
	RouteDatabaseConfigEnginePostgresql = RouteDatabaseConfigEngine("postgresql")
	RouteDatabaseConfigEngineMysql      = RouteDatabaseConfigEngine("mysql")
)

func (enum RouteDatabaseConfigEngine) String() string {
	if enum == "" {
		// return default value if empty
		return string(RouteDatabaseConfigEngineUnknown)
	}
	return string(enum)
}

func (enum RouteDatabaseConfigEngine) Values() []RouteDatabaseConfigEngine {
	return []RouteDatabaseConfigEngine{
		"unknown",
		"postgresql",
		"mysql",
	}
}

func (enum RouteDatabaseConfigEngine) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RouteDatabaseConfigEngine) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RouteDatabaseConfigEngine(RouteDatabaseConfigEngine(tmp).String())
	return nil
}

type RouteRestConfigHTTPVerb string

const (
	RouteRestConfigHTTPVerbUnknown = RouteRestConfigHTTPVerb("unknown")
	RouteRestConfigHTTPVerbGet     = RouteRestConfigHTTPVerb("get")
	RouteRestConfigHTTPVerbPost    = RouteRestConfigHTTPVerb("post")
	RouteRestConfigHTTPVerbPut     = RouteRestConfigHTTPVerb("put")
	RouteRestConfigHTTPVerbPatch   = RouteRestConfigHTTPVerb("patch")
	RouteRestConfigHTTPVerbDelete  = RouteRestConfigHTTPVerb("delete")
)

func (enum RouteRestConfigHTTPVerb) String() string {
	if enum == "" {
		// return default value if empty
		return string(RouteRestConfigHTTPVerbUnknown)
	}
	return string(enum)
}

func (enum RouteRestConfigHTTPVerb) Values() []RouteRestConfigHTTPVerb {
	return []RouteRestConfigHTTPVerb{
		"unknown",
		"get",
		"post",
		"put",
		"patch",
		"delete",
	}
}

func (enum RouteRestConfigHTTPVerb) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RouteRestConfigHTTPVerb) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RouteRestConfigHTTPVerb(RouteRestConfigHTTPVerb(tmp).String())
	return nil
}

type RouteRouteType string

const (
	RouteRouteTypeUnknown  = RouteRouteType("unknown")
	RouteRouteTypeS3       = RouteRouteType("s3")
	RouteRouteTypeDatabase = RouteRouteType("database")
	RouteRouteTypeRest     = RouteRouteType("rest")
)

func (enum RouteRouteType) String() string {
	if enum == "" {
		// return default value if empty
		return string(RouteRouteTypeUnknown)
	}
	return string(enum)
}

func (enum RouteRouteType) Values() []RouteRouteType {
	return []RouteRouteType{
		"unknown",
		"s3",
		"database",
		"rest",
	}
}

func (enum RouteRouteType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RouteRouteType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RouteRouteType(RouteRouteType(tmp).String())
	return nil
}

type RouteS3ConfigS3Strategy string

const (
	RouteS3ConfigS3StrategyUnknown    = RouteS3ConfigS3Strategy("unknown")
	RouteS3ConfigS3StrategyPerTopic   = RouteS3ConfigS3Strategy("per_topic")
	RouteS3ConfigS3StrategyPerMessage = RouteS3ConfigS3Strategy("per_message")
)

func (enum RouteS3ConfigS3Strategy) String() string {
	if enum == "" {
		// return default value if empty
		return string(RouteS3ConfigS3StrategyUnknown)
	}
	return string(enum)
}

func (enum RouteS3ConfigS3Strategy) Values() []RouteS3ConfigS3Strategy {
	return []RouteS3ConfigS3Strategy{
		"unknown",
		"per_topic",
		"per_message",
	}
}

func (enum RouteS3ConfigS3Strategy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RouteS3ConfigS3Strategy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RouteS3ConfigS3Strategy(RouteS3ConfigS3Strategy(tmp).String())
	return nil
}

// DeviceMessageFiltersRule: device message filters rule.
type DeviceMessageFiltersRule struct {
	// Policy: if set to `accept`, all topics in the topics list will be allowed, with all other topics being denied.
	// If set to `reject`, all topics in the topics list will be denied, with all other topics being allowed.
	// Default value: unknown
	Policy DeviceMessageFiltersRulePolicy `json:"policy"`

	// Topics: list of topics to accept or reject. It must be valid MQTT topics and up to 65535 characters.
	Topics *[]string `json:"topics"`
}

// DeviceMessageFilters: device message filters.
type DeviceMessageFilters struct {
	// Publish: filtering rule to restrict topics the device can publish to.
	Publish *DeviceMessageFiltersRule `json:"publish"`

	// Subscribe: filtering rule to restrict topics the device can subscribe to.
	Subscribe *DeviceMessageFiltersRule `json:"subscribe"`
}

// HubTwinsGraphiteConfig: hub twins graphite config.
type HubTwinsGraphiteConfig struct {
	PushURI string `json:"push_uri"`
}

// Certificate: certificate.
type Certificate struct {
	Crt string `json:"crt"`

	Key string `json:"key"`
}

// Device: device.
type Device struct {
	// ID: device ID, also used as MQTT Client ID or username.
	ID string `json:"id"`

	// Name: device name.
	Name string `json:"name"`

	// Description: device description.
	Description string `json:"description"`

	// Status: device status.
	// Default value: unknown
	Status DeviceStatus `json:"status"`

	// HubID: hub ID.
	HubID string `json:"hub_id"`

	// LastActivityAt: last connection/activity date of a device.
	LastActivityAt *time.Time `json:"last_activity_at"`

	// IsConnected: defines whether the device is connected to the Hub.
	IsConnected bool `json:"is_connected"`

	// AllowInsecure: defines whether to allow the device to connect to the Hub without TLS mutual authentication.
	AllowInsecure bool `json:"allow_insecure"`

	// AllowMultipleConnections: defines whether to allow multiple physical devices to connect to the Hub with this device's credentials.
	AllowMultipleConnections bool `json:"allow_multiple_connections"`

	// MessageFilters: filter-sets to restrict the topics the device can publish/subscribe to.
	MessageFilters *DeviceMessageFilters `json:"message_filters"`

	// HasCustomCertificate: assigning a custom certificate allows a device to authenticate using that specific certificate without checking the Hub's CA certificate.
	HasCustomCertificate bool `json:"has_custom_certificate"`

	// CreatedAt: date at which the device was added.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date at which the device was last modified.
	UpdatedAt *time.Time `json:"updated_at"`

	// Region: region of the device.
	Region scw.Region `json:"region"`
}

// Network: network.
type Network struct {
	// ID: network ID.
	ID string `json:"id"`

	// Name: network name.
	Name string `json:"name"`

	// Type: type of network to connect with.
	// Default value: unknown
	Type NetworkNetworkType `json:"type"`

	// Endpoint: endpoint to use for interacting with the network.
	Endpoint string `json:"endpoint"`

	// HubID: hub ID to connect the Network to.
	HubID string `json:"hub_id"`

	// CreatedAt: date at which the network was created.
	CreatedAt *time.Time `json:"created_at"`

	// TopicPrefix: this prefix will be prepended to all topics for this Network.
	TopicPrefix string `json:"topic_prefix"`

	// Region: region of the network.
	Region scw.Region `json:"region"`
}

// CreateRouteRequestDatabaseConfig: create route request database config.
type CreateRouteRequestDatabaseConfig struct {
	Host string `json:"host"`

	Port uint32 `json:"port"`

	Dbname string `json:"dbname"`

	Username string `json:"username"`

	Password string `json:"password"`

	Query string `json:"query"`

	// Engine: default value: unknown
	Engine RouteDatabaseConfigEngine `json:"engine"`
}

// CreateRouteRequestRestConfig: create route request rest config.
type CreateRouteRequestRestConfig struct {
	// Verb: default value: unknown
	Verb RouteRestConfigHTTPVerb `json:"verb"`

	URI string `json:"uri"`

	Headers map[string]string `json:"headers"`
}

// CreateRouteRequestS3Config: create route request s3 config.
type CreateRouteRequestS3Config struct {
	BucketRegion string `json:"bucket_region"`

	BucketName string `json:"bucket_name"`

	ObjectPrefix string `json:"object_prefix"`

	// Strategy: default value: unknown
	Strategy RouteS3ConfigS3Strategy `json:"strategy"`
}

// Hub: hub.
type Hub struct {
	// ID: hub ID.
	ID string `json:"id"`

	// Name: hub name.
	Name string `json:"name"`

	// Status: current status of the Hub.
	// Default value: unknown
	Status HubStatus `json:"status"`

	// ProductPlan: hub feature set.
	// Default value: plan_unknown
	ProductPlan HubProductPlan `json:"product_plan"`

	// Enabled: defines whether the hub has been enabled.
	Enabled bool `json:"enabled"`

	// DeviceCount: number of registered devices.
	DeviceCount uint64 `json:"device_count"`

	// ConnectedDeviceCount: number of currently connected devices.
	ConnectedDeviceCount uint64 `json:"connected_device_count"`

	// Endpoint: devices should be connected to this host. Port may be 1883 (MQTT), 8883 (MQTT over TLS), 80 (MQTT over Websocket) or 443 (MQTT over Websocket over TLS).
	Endpoint string `json:"endpoint"`

	// DisableEvents: defines whether to disable Hub events.
	DisableEvents bool `json:"disable_events"`

	// EventsTopicPrefix: hub events topic prefix.
	EventsTopicPrefix string `json:"events_topic_prefix"`

	// Region: region of the Hub.
	Region scw.Region `json:"region"`

	// CreatedAt: hub creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: hub last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// ProjectID: project owning the resource.
	ProjectID string `json:"project_id"`

	// OrganizationID: organization owning the resource.
	OrganizationID string `json:"organization_id"`

	// EnableDeviceAutoProvisioning: when an unknown device connects to your hub using a valid certificate chain, it will be automatically provisioned inside your Hub. The Hub uses the common name of the device certificate to find out if a device with the same name already exists. This setting can only be enabled on a hub with a custom certificate authority.
	EnableDeviceAutoProvisioning bool `json:"enable_device_auto_provisioning"`

	// HasCustomCa: flag is automatically set to `false` after Hub creation, as Hub certificates are managed by Scaleway. Once a custom certificate authority is set, the flag will be set to `true`.
	HasCustomCa bool `json:"has_custom_ca"`

	// TwinsGraphiteConfig: bETA - not implemented yet.
	// Precisely one of TwinsGraphiteConfig must be set.
	TwinsGraphiteConfig *HubTwinsGraphiteConfig `json:"twins_graphite_config,omitempty"`
}

// RouteSummary: route summary.
type RouteSummary struct {
	// ID: route ID.
	ID string `json:"id"`

	// Name: route name.
	Name string `json:"name"`

	// HubID: hub ID of the route.
	HubID string `json:"hub_id"`

	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters.
	Topic string `json:"topic"`

	// Type: route type.
	// Default value: unknown
	Type RouteRouteType `json:"type"`

	// CreatedAt: date at which the route was created.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: date at which the route was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// Region: region of the route.
	Region scw.Region `json:"region"`
}

// ListTwinDocumentsResponseDocumentSummary: list twin documents response document summary.
type ListTwinDocumentsResponseDocumentSummary struct {
	// DocumentName: name of the document.
	DocumentName string `json:"document_name"`
}

// RouteDatabaseConfig: route database config.
type RouteDatabaseConfig struct {
	// Engine: database engine the route will connect to. If not specified, the default database will be 'PostgreSQL'.
	// Default value: unknown
	Engine RouteDatabaseConfigEngine `json:"engine"`

	// Host: database host.
	Host string `json:"host"`

	// Port: database port.
	Port uint32 `json:"port"`

	// Dbname: database name.
	Dbname string `json:"dbname"`

	// Username: database username. Make sure this account can execute the provided query.
	Username string `json:"username"`

	// Password: database password.
	Password string `json:"password"`

	// Query: SQL query to be executed ($TOPIC and $PAYLOAD variables are available, see documentation).
	Query string `json:"query"`
}

// RouteRestConfig: route rest config.
type RouteRestConfig struct {
	// Verb: HTTP verb used to call REST URI.
	// Default value: unknown
	Verb RouteRestConfigHTTPVerb `json:"verb"`

	// URI: URI of the REST endpoint.
	URI string `json:"uri"`

	// Headers: HTTP call extra headers.
	Headers map[string]string `json:"headers"`
}

// RouteS3Config: route s3 config.
type RouteS3Config struct {
	// BucketRegion: region of the Amazon S3 route's destination bucket (e.g., 'fr-par').
	BucketRegion string `json:"bucket_region"`

	// BucketName: destination bucket name of the Amazon S3 route.
	BucketName string `json:"bucket_name"`

	// ObjectPrefix: optional string to prefix object names with.
	ObjectPrefix string `json:"object_prefix"`

	// Strategy: how the Amazon S3 route's objects will be created: one per topic or one per message.
	// Default value: unknown
	Strategy RouteS3ConfigS3Strategy `json:"strategy"`
}

// UpdateRouteRequestDatabaseConfig: update route request database config.
type UpdateRouteRequestDatabaseConfig struct {
	Host *string `json:"host"`

	Port *uint32 `json:"port"`

	Dbname *string `json:"dbname"`

	Username *string `json:"username"`

	Password *string `json:"password"`

	Query *string `json:"query"`

	// Engine: default value: unknown
	Engine RouteDatabaseConfigEngine `json:"engine"`
}

// UpdateRouteRequestRestConfig: update route request rest config.
type UpdateRouteRequestRestConfig struct {
	// Verb: default value: unknown
	Verb RouteRestConfigHTTPVerb `json:"verb"`

	URI *string `json:"uri"`

	Headers *map[string]string `json:"headers"`
}

// UpdateRouteRequestS3Config: update route request s3 config.
type UpdateRouteRequestS3Config struct {
	BucketRegion *string `json:"bucket_region"`

	BucketName *string `json:"bucket_name"`

	ObjectPrefix *string `json:"object_prefix"`

	// Strategy: default value: unknown
	Strategy RouteS3ConfigS3Strategy `json:"strategy"`
}

// CreateDeviceRequest: create device request.
type CreateDeviceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: device name.
	Name string `json:"name"`

	// HubID: hub ID of the device.
	HubID string `json:"hub_id"`

	// AllowInsecure: defines whether to allow plain and server-authenticated SSL connections in addition to mutually-authenticated ones.
	AllowInsecure bool `json:"allow_insecure"`

	// AllowMultipleConnections: defines whether to allow multiple physical devices to connect with this device's credentials.
	AllowMultipleConnections bool `json:"allow_multiple_connections"`

	// MessageFilters: filter-sets to authorize or deny the device to publish/subscribe to specific topics.
	MessageFilters *DeviceMessageFilters `json:"message_filters,omitempty"`

	// Description: device description.
	Description *string `json:"description,omitempty"`
}

// CreateDeviceResponse: create device response.
type CreateDeviceResponse struct {
	// Device: information related to the created device.
	Device *Device `json:"device"`

	// Certificate: device certificate.
	Certificate *Certificate `json:"certificate"`
}

// CreateHubRequest: create hub request.
type CreateHubRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: hub name (up to 255 characters).
	Name string `json:"name"`

	// ProjectID: project/Organization ID to filter for, only Hubs from this Project/Organization will be returned.
	ProjectID string `json:"project_id"`

	// ProductPlan: hub product plan.
	// Default value: plan_unknown
	ProductPlan HubProductPlan `json:"product_plan"`

	// DisableEvents: disable Hub events.
	DisableEvents *bool `json:"disable_events,omitempty"`

	// EventsTopicPrefix: topic prefix (default '$SCW/events') of Hub events.
	EventsTopicPrefix *string `json:"events_topic_prefix,omitempty"`

	// TwinsGraphiteConfig: bETA - not implemented yet.
	// Precisely one of TwinsGraphiteConfig must be set.
	TwinsGraphiteConfig *HubTwinsGraphiteConfig `json:"twins_graphite_config,omitempty"`
}

// CreateNetworkRequest: create network request.
type CreateNetworkRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: network name.
	Name string `json:"name"`

	// Type: type of network to connect with.
	// Default value: unknown
	Type NetworkNetworkType `json:"type"`

	// HubID: hub ID to connect the Network to.
	HubID string `json:"hub_id"`

	// TopicPrefix: topic prefix for the Network.
	TopicPrefix string `json:"topic_prefix"`
}

// CreateNetworkResponse: create network response.
type CreateNetworkResponse struct {
	// Network: information related to the created network.
	Network *Network `json:"network"`

	// Secret: endpoint Key to keep secret. This cannot be retrieved later.
	Secret string `json:"secret"`
}

// CreateRouteRequest: create route request.
type CreateRouteRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: route name.
	Name string `json:"name"`

	// HubID: hub ID of the route.
	HubID string `json:"hub_id"`

	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters.
	Topic string `json:"topic"`

	// S3Config: if creating Amazon S3 Routes, Amazon S3-specific configuration fields.
	// Precisely one of S3Config, DbConfig, RestConfig must be set.
	S3Config *CreateRouteRequestS3Config `json:"s3_config,omitempty"`

	// DbConfig: if creating Database Route, DB-specific configuration fields.
	// Precisely one of S3Config, DbConfig, RestConfig must be set.
	DbConfig *CreateRouteRequestDatabaseConfig `json:"db_config,omitempty"`

	// RestConfig: if creating Rest Route, Rest-specific configuration fields.
	// Precisely one of S3Config, DbConfig, RestConfig must be set.
	RestConfig *CreateRouteRequestRestConfig `json:"rest_config,omitempty"`
}

// DeleteDeviceRequest: delete device request.
type DeleteDeviceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeviceID: device ID.
	DeviceID string `json:"-"`
}

// DeleteHubRequest: delete hub request.
type DeleteHubRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HubID: hub ID.
	HubID string `json:"-"`

	// DeleteDevices: defines whether to force the deletion of devices added to this Hub or reject the operation.
	DeleteDevices *bool `json:"delete_devices,omitempty"`
}

// DeleteNetworkRequest: delete network request.
type DeleteNetworkRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NetworkID: network ID.
	NetworkID string `json:"-"`
}

// DeleteRouteRequest: delete route request.
type DeleteRouteRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// RouteID: route ID.
	RouteID string `json:"-"`
}

// DeleteTwinDocumentRequest: delete twin document request.
type DeleteTwinDocumentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TwinID: twin ID.
	TwinID string `json:"-"`

	// DocumentName: name of the document.
	DocumentName string `json:"-"`
}

// DeleteTwinDocumentsRequest: delete twin documents request.
type DeleteTwinDocumentsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TwinID: twin ID.
	TwinID string `json:"-"`
}

// DisableDeviceRequest: disable device request.
type DisableDeviceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeviceID: device ID.
	DeviceID string `json:"-"`
}

// DisableHubRequest: disable hub request.
type DisableHubRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HubID: hub ID.
	HubID string `json:"-"`
}

// EnableDeviceRequest: enable device request.
type EnableDeviceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeviceID: device ID.
	DeviceID string `json:"-"`
}

// EnableHubRequest: enable hub request.
type EnableHubRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HubID: hub ID.
	HubID string `json:"-"`
}

// GetDeviceCertificateRequest: get device certificate request.
type GetDeviceCertificateRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeviceID: device ID.
	DeviceID string `json:"-"`
}

// GetDeviceCertificateResponse: get device certificate response.
type GetDeviceCertificateResponse struct {
	// Device: information related to the created device.
	Device *Device `json:"device"`

	// CertificatePem: device certificate.
	CertificatePem string `json:"certificate_pem"`
}

// GetDeviceMetricsRequest: get device metrics request.
type GetDeviceMetricsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeviceID: device ID.
	DeviceID string `json:"-"`

	// StartDate: start date used to compute the best scale for the returned metrics.
	StartDate *time.Time `json:"start_date,omitempty"`
}

// GetDeviceMetricsResponse: get device metrics response.
type GetDeviceMetricsResponse struct {
	// Metrics: metrics for a device over the requested period.
	Metrics []*scw.TimeSeries `json:"metrics"`
}

// GetDeviceRequest: get device request.
type GetDeviceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeviceID: device ID.
	DeviceID string `json:"-"`
}

// GetHubCARequest: get hub ca request.
type GetHubCARequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	HubID string `json:"-"`
}

// GetHubCAResponse: get hub ca response.
type GetHubCAResponse struct {
	CaCertPem string `json:"ca_cert_pem"`
}

// GetHubMetricsRequest: get hub metrics request.
type GetHubMetricsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HubID: hub ID.
	HubID string `json:"-"`

	// StartDate: start date used to compute the best scale for returned metrics.
	StartDate *time.Time `json:"start_date,omitempty"`
}

// GetHubMetricsResponse: get hub metrics response.
type GetHubMetricsResponse struct {
	// Metrics: metrics for a Hub over the requested period.
	Metrics []*scw.TimeSeries `json:"metrics"`
}

// GetHubRequest: get hub request.
type GetHubRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HubID: hub ID.
	HubID string `json:"-"`
}

// GetNetworkRequest: get network request.
type GetNetworkRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NetworkID: network ID.
	NetworkID string `json:"-"`
}

// GetRouteRequest: get route request.
type GetRouteRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// RouteID: route ID.
	RouteID string `json:"-"`
}

// GetTwinDocumentRequest: get twin document request.
type GetTwinDocumentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TwinID: twin ID.
	TwinID string `json:"-"`

	// DocumentName: name of the document.
	DocumentName string `json:"-"`
}

// ListDevicesRequest: list devices request.
type ListDevicesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of devices to return within a page. Maximum value is 100.
	PageSize *uint32 `json:"-"`

	// OrderBy: ordering of requested devices.
	// Default value: name_asc
	OrderBy ListDevicesRequestOrderBy `json:"-"`

	// Name: name to filter for, only devices with this name will be returned.
	Name *string `json:"-"`

	// HubID: hub ID to filter for, only devices attached to this Hub will be returned.
	HubID *string `json:"-"`

	// AllowInsecure: defines whether to filter the allow_insecure flag.
	AllowInsecure *bool `json:"-"`

	// Status: device status (enabled, disabled, etc.).
	// Default value: unknown
	Status DeviceStatus `json:"-"`
}

// ListDevicesResponse: list devices response.
type ListDevicesResponse struct {
	// TotalCount: total number of devices.
	TotalCount uint32 `json:"total_count"`

	// Devices: page of devices.
	Devices []*Device `json:"devices"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDevicesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDevicesResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListDevicesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Devices = append(r.Devices, results.Devices...)
	r.TotalCount += uint32(len(results.Devices))
	return uint32(len(results.Devices)), nil
}

// ListHubsRequest: list hubs request.
type ListHubsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of Hubs to return within a page. Maximum value is 100.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of Hubs in the response.
	// Default value: name_asc
	OrderBy ListHubsRequestOrderBy `json:"-"`

	// ProjectID: only list Hubs of this Project ID.
	ProjectID *string `json:"-"`

	// OrganizationID: only list Hubs of this Organization ID.
	OrganizationID *string `json:"-"`

	// Name: hub name.
	Name *string `json:"-"`
}

// ListHubsResponse: list hubs response.
type ListHubsResponse struct {
	// TotalCount: total number of Hubs.
	TotalCount uint32 `json:"total_count"`

	// Hubs: a page of hubs.
	Hubs []*Hub `json:"hubs"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListHubsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListHubsResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListHubsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Hubs = append(r.Hubs, results.Hubs...)
	r.TotalCount += uint32(len(results.Hubs))
	return uint32(len(results.Hubs)), nil
}

// ListNetworksRequest: list networks request.
type ListNetworksRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of networks to return. The maximum value is 100.
	PageSize *uint32 `json:"-"`

	// OrderBy: ordering of requested routes.
	// Default value: name_asc
	OrderBy ListNetworksRequestOrderBy `json:"-"`

	// Name: network name to filter for.
	Name *string `json:"-"`

	// HubID: hub ID to filter for.
	HubID *string `json:"-"`

	// TopicPrefix: topic prefix to filter for.
	TopicPrefix *string `json:"-"`
}

// ListNetworksResponse: list networks response.
type ListNetworksResponse struct {
	// TotalCount: total number of Networks.
	TotalCount uint32 `json:"total_count"`

	// Networks: page of networks.
	Networks []*Network `json:"networks"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNetworksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNetworksResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListNetworksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Networks = append(r.Networks, results.Networks...)
	r.TotalCount += uint32(len(results.Networks))
	return uint32(len(results.Networks)), nil
}

// ListRoutesRequest: list routes request.
type ListRoutesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Page: page number to return, from the paginated results.
	Page *int32 `json:"-"`

	// PageSize: number of routes to return within a page. Maximum value is 100.
	PageSize *uint32 `json:"-"`

	// OrderBy: ordering of requested routes.
	// Default value: name_asc
	OrderBy ListRoutesRequestOrderBy `json:"-"`

	// HubID: hub ID to filter for.
	HubID *string `json:"-"`

	// Name: route name to filter for.
	Name *string `json:"-"`
}

// ListRoutesResponse: list routes response.
type ListRoutesResponse struct {
	// TotalCount: total number of routes.
	TotalCount uint32 `json:"total_count"`

	// Routes: page of routes.
	Routes []*RouteSummary `json:"routes"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRoutesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRoutesResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListRoutesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Routes = append(r.Routes, results.Routes...)
	r.TotalCount += uint32(len(results.Routes))
	return uint32(len(results.Routes)), nil
}

// ListTwinDocumentsRequest: list twin documents request.
type ListTwinDocumentsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TwinID: twin ID.
	TwinID string `json:"-"`
}

// ListTwinDocumentsResponse: list twin documents response.
type ListTwinDocumentsResponse struct {
	// Documents: list of the twin document.
	Documents []*ListTwinDocumentsResponseDocumentSummary `json:"documents"`
}

// PatchTwinDocumentRequest: patch twin document request.
type PatchTwinDocumentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TwinID: twin ID.
	TwinID string `json:"-"`

	// DocumentName: name of the document.
	DocumentName string `json:"-"`

	// Version: if set, ensures that the current version of the document matches before persisting the update.
	Version *uint32 `json:"version,omitempty"`

	// Data: a json data that will be applied on the document's current data.
	// Patching rules:
	// * The patch goes recursively through the patch objects.
	// * If the patch object property is null, it is removed from the final object.
	// * If the patch object property is a value (number, strings, bool, arrays), it is replaced.
	// * If the patch object property is an object, the previous rules will be applied recursively on it.
	Data *scw.JSONObject `json:"data,omitempty"`
}

// PutTwinDocumentRequest: put twin document request.
type PutTwinDocumentRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// TwinID: twin ID.
	TwinID string `json:"-"`

	// DocumentName: name of the document.
	DocumentName string `json:"-"`

	// Version: if set, ensures that the current version of the document matches before persisting the update.
	Version *uint32 `json:"version,omitempty"`

	// Data: new data that will replace the contents of the document.
	Data *scw.JSONObject `json:"data,omitempty"`
}

// RenewDeviceCertificateRequest: renew device certificate request.
type RenewDeviceCertificateRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeviceID: device ID.
	DeviceID string `json:"-"`
}

// RenewDeviceCertificateResponse: renew device certificate response.
type RenewDeviceCertificateResponse struct {
	// Device: information related to the created device.
	Device *Device `json:"device"`

	// Certificate: device certificate.
	Certificate *Certificate `json:"certificate"`
}

// Route: route.
type Route struct {
	// ID: route ID.
	ID string `json:"id"`

	// Name: route name.
	Name string `json:"name"`

	// HubID: hub ID of the route.
	HubID string `json:"hub_id"`

	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters.
	Topic string `json:"topic"`

	// Type: route type.
	// Default value: unknown
	Type RouteRouteType `json:"type"`

	// CreatedAt: date at which the route was created.
	CreatedAt *time.Time `json:"created_at"`

	// S3Config: when using Amazon S3 Routes, Amazon S3-specific configuration fields.
	// Precisely one of S3Config, DbConfig, RestConfig must be set.
	S3Config *RouteS3Config `json:"s3_config,omitempty"`

	// DbConfig: when using Database Route, DB-specific configuration fields.
	// Precisely one of S3Config, DbConfig, RestConfig must be set.
	DbConfig *RouteDatabaseConfig `json:"db_config,omitempty"`

	// RestConfig: when using Rest Route, Rest-specific configuration fields.
	// Precisely one of S3Config, DbConfig, RestConfig must be set.
	RestConfig *RouteRestConfig `json:"rest_config,omitempty"`

	// UpdatedAt: date at which the route was last updated.
	UpdatedAt *time.Time `json:"updated_at"`

	// Region: region of the route.
	Region scw.Region `json:"region"`
}

// SetDeviceCertificateRequest: set device certificate request.
type SetDeviceCertificateRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeviceID: device ID.
	DeviceID string `json:"-"`

	// CertificatePem: pEM-encoded custom certificate.
	CertificatePem string `json:"certificate_pem"`
}

// SetDeviceCertificateResponse: set device certificate response.
type SetDeviceCertificateResponse struct {
	Device *Device `json:"device"`

	CertificatePem string `json:"certificate_pem"`
}

// SetHubCARequest: set hub ca request.
type SetHubCARequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HubID: hub ID.
	HubID string `json:"-"`

	// CaCertPem: cA's PEM-encoded certificate.
	CaCertPem string `json:"ca_cert_pem"`

	// ChallengeCertPem: challenge is a PEM-encoded certificate that acts as proof of possession of the CA. It must be signed by the CA, and have a Common Name equal to the Hub ID.
	ChallengeCertPem string `json:"challenge_cert_pem"`
}

// TwinDocument: twin document.
type TwinDocument struct {
	// TwinID: parent twin ID of the document.
	TwinID string `json:"twin_id"`

	// DocumentName: name of the document.
	DocumentName string `json:"document_name"`

	// Version: new version of the document.
	Version uint32 `json:"version"`

	// Data: new data related to the document.
	Data *scw.JSONObject `json:"data"`
}

// UpdateDeviceRequest: update device request.
type UpdateDeviceRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DeviceID: device ID.
	DeviceID string `json:"-"`

	// Description: description for the device.
	Description *string `json:"description,omitempty"`

	// AllowInsecure: defines whether to allow plain and server-authenticated SSL connections in addition to mutually-authenticated ones.
	AllowInsecure *bool `json:"allow_insecure,omitempty"`

	// AllowMultipleConnections: defines whether to allow multiple physical devices to connect with this device's credentials.
	AllowMultipleConnections *bool `json:"allow_multiple_connections,omitempty"`

	// MessageFilters: filter-sets to restrict the topics the device can publish/subscribe to.
	MessageFilters *DeviceMessageFilters `json:"message_filters,omitempty"`

	// HubID: change Hub for this device, additional fees may apply, see IoT Hub pricing.
	HubID *string `json:"hub_id,omitempty"`
}

// UpdateHubRequest: update hub request.
type UpdateHubRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// HubID: ID of the Hub you want to update.
	HubID string `json:"-"`

	// Name: hub name (up to 255 characters).
	Name *string `json:"name,omitempty"`

	// ProductPlan: hub product plan.
	// Default value: plan_unknown
	ProductPlan HubProductPlan `json:"product_plan"`

	// DisableEvents: disable Hub events.
	DisableEvents *bool `json:"disable_events,omitempty"`

	// EventsTopicPrefix: topic prefix of Hub events.
	EventsTopicPrefix *string `json:"events_topic_prefix,omitempty"`

	// EnableDeviceAutoProvisioning: enable device auto provisioning.
	EnableDeviceAutoProvisioning *bool `json:"enable_device_auto_provisioning,omitempty"`

	// TwinsGraphiteConfig: bETA - not implemented yet.
	// Precisely one of TwinsGraphiteConfig must be set.
	TwinsGraphiteConfig *HubTwinsGraphiteConfig `json:"twins_graphite_config,omitempty"`
}

// UpdateRouteRequest: update route request.
type UpdateRouteRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// RouteID: route id.
	RouteID string `json:"-"`

	// Name: route name.
	Name *string `json:"name,omitempty"`

	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters.
	Topic *string `json:"topic,omitempty"`

	// S3Config: when updating Amazon S3 Route, Amazon S3-specific configuration fields.
	// Precisely one of S3Config, DbConfig, RestConfig must be set.
	S3Config *UpdateRouteRequestS3Config `json:"s3_config,omitempty"`

	// DbConfig: when updating Database Route, DB-specific configuration fields.
	// Precisely one of S3Config, DbConfig, RestConfig must be set.
	DbConfig *UpdateRouteRequestDatabaseConfig `json:"db_config,omitempty"`

	// RestConfig: when updating Rest Route, Rest-specific configuration fields.
	// Precisely one of S3Config, DbConfig, RestConfig must be set.
	RestConfig *UpdateRouteRequestRestConfig `json:"rest_config,omitempty"`
}

// This API allows you to manage your IoT hubs and devices.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar}
}

// ListHubs: List all Hubs in the specified zone. By default, returned Hubs are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListHubs(req *ListHubsRequest, opts ...scw.RequestOption) (*ListHubsResponse, error) {
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs",
		Query:  query,
	}

	var resp ListHubsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateHub: Create a new Hub in the targeted region, specifying its configuration including name and product plan.
func (s *API) CreateHub(req *CreateHubRequest, opts ...scw.RequestOption) (*Hub, error) {
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
		req.Name = namegenerator.GetRandomName("hub")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHub: Retrieve information about an existing IoT Hub, specified by its Hub ID. Its full details, including name, status and endpoint, are returned in the response object.
func (s *API) GetHub(req *GetHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "",
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForHubRequest is used by WaitForHub method.
type WaitForHubRequest struct {
	Region        scw.Region
	HubID         string
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForHub waits for the Hub to reach a terminal state.
func (s *API) WaitForHub(req *WaitForHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	timeout := defaultIotTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultIotRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[HubStatus]struct{}{
		HubStatusEnabling:  {},
		HubStatusDisabling: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (any, bool, error) {
			res, err := s.GetHub(&GetHubRequest{
				Region: req.Region,
				HubID:  req.HubID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTransient := transientStatuses[res.Status]

			return res, !isTransient, nil
		},
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
		Timeout:          timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for Hub failed")
	}

	return res.(*Hub), nil
}

// UpdateHub: Update the parameters of an existing IoT Hub, specified by its Hub ID.
func (s *API) UpdateHub(req *UpdateHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableHub: Enable an existing IoT Hub, specified by its Hub ID.
func (s *API) EnableHub(req *EnableHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/enable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableHub: Disable an existing IoT Hub, specified by its Hub ID.
func (s *API) DisableHub(req *DisableHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/disable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteHub: Delete an existing IoT Hub, specified by its Hub ID. Deleting a Hub is permanent, and cannot be undone.
func (s *API) DeleteHub(req *DeleteHubRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "delete_devices", req.DeleteDevices)

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "",
		Query:  query,
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: GetHubMetrics: Get the metrics of an existing IoT Hub, specified by its Hub ID.
func (s *API) GetHubMetrics(req *GetHubMetricsRequest, opts ...scw.RequestOption) (*GetHubMetricsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/metrics",
		Query:  query,
	}

	var resp GetHubMetricsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetHubCA: Set a particular PEM-encoded certificate, specified by the Hub ID.
func (s *API) SetHubCA(req *SetHubCARequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/ca",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHubCA: Get information for a particular PEM-encoded certificate, specified by the Hub ID.
func (s *API) GetHubCA(req *GetHubCARequest, opts ...scw.RequestOption) (*GetHubCAResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/ca",
	}

	var resp GetHubCAResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDevices: List all devices in the specified region. By default, returned devices are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListDevices(req *ListDevicesRequest, opts ...scw.RequestOption) (*ListDevicesResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "hub_id", req.HubID)
	parameter.AddToQuery(query, "allow_insecure", req.AllowInsecure)
	parameter.AddToQuery(query, "status", req.Status)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices",
		Query:  query,
	}

	var resp ListDevicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateDevice: Attach a device to a given Hub.
func (s *API) CreateDevice(req *CreateDeviceRequest, opts ...scw.RequestOption) (*CreateDeviceResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("device")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateDeviceResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDevice: Retrieve information about an existing device, specified by its device ID. Its full details, including name, status and ID, are returned in the response object.
func (s *API) GetDevice(req *GetDeviceRequest, opts ...scw.RequestOption) (*Device, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "",
	}

	var resp Device

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDevice: Update the parameters of an existing device, specified by its device ID.
func (s *API) UpdateDevice(req *UpdateDeviceRequest, opts ...scw.RequestOption) (*Device, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Device

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// EnableDevice: Enable a specific device, specified by its device ID.
func (s *API) EnableDevice(req *EnableDeviceRequest, opts ...scw.RequestOption) (*Device, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/enable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Device

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DisableDevice: Disable an existing device, specified by its device ID.
func (s *API) DisableDevice(req *DisableDeviceRequest, opts ...scw.RequestOption) (*Device, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/disable",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Device

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RenewDeviceCertificate: Renew the certificate of an existing device, specified by its device ID.
func (s *API) RenewDeviceCertificate(req *RenewDeviceCertificateRequest, opts ...scw.RequestOption) (*RenewDeviceCertificateResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/renew-certificate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RenewDeviceCertificateResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetDeviceCertificate: Switch the existing certificate of a given device with an EM-encoded custom certificate.
func (s *API) SetDeviceCertificate(req *SetDeviceCertificateRequest, opts ...scw.RequestOption) (*SetDeviceCertificateResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/certificate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SetDeviceCertificateResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDeviceCertificate: Get information for a particular PEM-encoded certificate, specified by the device ID. The response returns full details of the device, including its type of certificate.
func (s *API) GetDeviceCertificate(req *GetDeviceCertificateRequest, opts ...scw.RequestOption) (*GetDeviceCertificateResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/certificate",
	}

	var resp GetDeviceCertificateResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteDevice: Remove a specific device from the specific Hub it is attached to.
func (s *API) DeleteDevice(req *DeleteDeviceRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// Deprecated: GetDeviceMetrics: Get the metrics of an existing device, specified by its device ID.
func (s *API) GetDeviceMetrics(req *GetDeviceMetricsRequest, opts ...scw.RequestOption) (*GetDeviceMetricsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/metrics",
		Query:  query,
	}

	var resp GetDeviceMetricsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListRoutes: List all routes in the specified region. By default, returned routes are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *API) ListRoutes(req *ListRoutesRequest, opts ...scw.RequestOption) (*ListRoutesResponse, error) {
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
	parameter.AddToQuery(query, "hub_id", req.HubID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes",
		Query:  query,
	}

	var resp ListRoutesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateRoute: Multiple kinds of routes can be created, such as:
//   - Database Route
//     Create a route that will record subscribed MQTT messages into your database.
//     <b>You need to manage the database by yourself</b>.
//   - REST Route.
//     Create a route that will call a REST API on received subscribed MQTT messages.
//   - Amazon S3 Routes.
//     Create a route that will put subscribed MQTT messages into an Object Storage bucket.
//     You need to create the bucket yourself and grant write access.
//     Granting can be done with s3cmd (`s3cmd setacl s3://<my-bucket> --acl-grant=write:555c69c3-87d0-4bf8-80f1-99a2f757d031:555c69c3-87d0-4bf8-80f1-99a2f757d031`).
func (s *API) CreateRoute(req *CreateRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("route")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateRoute: Update the parameters of an existing route, specified by its route ID.
func (s *API) UpdateRoute(req *UpdateRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.RouteID) == "" {
		return nil, errors.New("field RouteID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetRoute: Get information for a particular route, specified by the route ID. The response returns full details of the route, including its type, the topic it subscribes to and its configuration.
func (s *API) GetRoute(req *GetRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.RouteID) == "" {
		return nil, errors.New("field RouteID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteRoute: Delete an existing route, specified by its route ID. Deleting a route is permanent, and cannot be undone.
func (s *API) DeleteRoute(req *DeleteRouteRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.RouteID) == "" {
		return errors.New("field RouteID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListNetworks: List the networks.
func (s *API) ListNetworks(req *ListNetworksRequest, opts ...scw.RequestOption) (*ListNetworksResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "hub_id", req.HubID)
	parameter.AddToQuery(query, "topic_prefix", req.TopicPrefix)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/networks",
		Query:  query,
	}

	var resp ListNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateNetwork: Create a new network for an existing hub.  Beside the default network, you can add networks for different data providers. Possible network types are Sigfox and REST.
func (s *API) CreateNetwork(req *CreateNetworkRequest, opts ...scw.RequestOption) (*CreateNetworkResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("network")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/networks",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CreateNetworkResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNetwork: Retrieve an existing network, specified by its network ID.  The response returns full details of the network, including its type, the topic prefix and its endpoint.
func (s *API) GetNetwork(req *GetNetworkRequest, opts ...scw.RequestOption) (*Network, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NetworkID) == "" {
		return nil, errors.New("field NetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/networks/" + fmt.Sprint(req.NetworkID) + "",
	}

	var resp Network

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNetwork: Delete an existing network, specified by its network ID. Deleting a network is permanent, and cannot be undone.
func (s *API) DeleteNetwork(req *DeleteNetworkRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NetworkID) == "" {
		return errors.New("field NetworkID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/networks/" + fmt.Sprint(req.NetworkID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetTwinDocument: BETA - Get a Cloud Twin Document.
func (s *API) GetTwinDocument(req *GetTwinDocumentRequest, opts ...scw.RequestOption) (*TwinDocument, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TwinID) == "" {
		return nil, errors.New("field TwinID cannot be empty in request")
	}

	if fmt.Sprint(req.DocumentName) == "" {
		return nil, errors.New("field DocumentName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/twins/" + fmt.Sprint(req.TwinID) + "/documents/" + fmt.Sprint(req.DocumentName) + "",
	}

	var resp TwinDocument

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PutTwinDocument: BETA - Update a Cloud Twin Document.
func (s *API) PutTwinDocument(req *PutTwinDocumentRequest, opts ...scw.RequestOption) (*TwinDocument, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TwinID) == "" {
		return nil, errors.New("field TwinID cannot be empty in request")
	}

	if fmt.Sprint(req.DocumentName) == "" {
		return nil, errors.New("field DocumentName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PUT",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/twins/" + fmt.Sprint(req.TwinID) + "/documents/" + fmt.Sprint(req.DocumentName) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TwinDocument

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// PatchTwinDocument: BETA - Patch a Cloud Twin Document.
func (s *API) PatchTwinDocument(req *PatchTwinDocumentRequest, opts ...scw.RequestOption) (*TwinDocument, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TwinID) == "" {
		return nil, errors.New("field TwinID cannot be empty in request")
	}

	if fmt.Sprint(req.DocumentName) == "" {
		return nil, errors.New("field DocumentName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/twins/" + fmt.Sprint(req.TwinID) + "/documents/" + fmt.Sprint(req.DocumentName) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp TwinDocument

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteTwinDocument: BETA - Delete a Cloud Twin Document.
func (s *API) DeleteTwinDocument(req *DeleteTwinDocumentRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TwinID) == "" {
		return errors.New("field TwinID cannot be empty in request")
	}

	if fmt.Sprint(req.DocumentName) == "" {
		return errors.New("field DocumentName cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/twins/" + fmt.Sprint(req.TwinID) + "/documents/" + fmt.Sprint(req.DocumentName) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListTwinDocuments: BETA - List the documents of a Cloud Twin.
func (s *API) ListTwinDocuments(req *ListTwinDocumentsRequest, opts ...scw.RequestOption) (*ListTwinDocumentsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TwinID) == "" {
		return nil, errors.New("field TwinID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/twins/" + fmt.Sprint(req.TwinID) + "",
	}

	var resp ListTwinDocumentsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteTwinDocuments: BETA - Delete all the documents of a Cloud Twin.
func (s *API) DeleteTwinDocuments(req *DeleteTwinDocumentsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.TwinID) == "" {
		return errors.New("field TwinID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/twins/" + fmt.Sprint(req.TwinID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

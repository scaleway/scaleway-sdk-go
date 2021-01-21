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

// API: this API allows you to manage IoT hubs and devices
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type DeviceMessageFiltersRulePolicy string

const (
	// DeviceMessageFiltersRulePolicyUnknown is [insert doc].
	DeviceMessageFiltersRulePolicyUnknown = DeviceMessageFiltersRulePolicy("unknown")
	// DeviceMessageFiltersRulePolicyAccept is [insert doc].
	DeviceMessageFiltersRulePolicyAccept = DeviceMessageFiltersRulePolicy("accept")
	// DeviceMessageFiltersRulePolicyReject is [insert doc].
	DeviceMessageFiltersRulePolicyReject = DeviceMessageFiltersRulePolicy("reject")
)

func (enum DeviceMessageFiltersRulePolicy) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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
	// DeviceStatusUnknown is [insert doc].
	DeviceStatusUnknown = DeviceStatus("unknown")
	// DeviceStatusError is [insert doc].
	DeviceStatusError = DeviceStatus("error")
	// DeviceStatusEnabled is [insert doc].
	DeviceStatusEnabled = DeviceStatus("enabled")
	// DeviceStatusDisabled is [insert doc].
	DeviceStatusDisabled = DeviceStatus("disabled")
)

func (enum DeviceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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
	// HubProductPlanPlanUnknown is [insert doc].
	HubProductPlanPlanUnknown = HubProductPlan("plan_unknown")
	// HubProductPlanPlanShared is [insert doc].
	HubProductPlanPlanShared = HubProductPlan("plan_shared")
	// HubProductPlanPlanDedicated is [insert doc].
	HubProductPlanPlanDedicated = HubProductPlan("plan_dedicated")
	// HubProductPlanPlanHa is [insert doc].
	HubProductPlanPlanHa = HubProductPlan("plan_ha")
)

func (enum HubProductPlan) String() string {
	if enum == "" {
		// return default value if empty
		return "plan_unknown"
	}
	return string(enum)
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
	// HubStatusUnknown is [insert doc].
	HubStatusUnknown = HubStatus("unknown")
	// HubStatusError is [insert doc].
	HubStatusError = HubStatus("error")
	// HubStatusEnabling is [insert doc].
	HubStatusEnabling = HubStatus("enabling")
	// HubStatusReady is [insert doc].
	HubStatusReady = HubStatus("ready")
	// HubStatusDisabling is [insert doc].
	HubStatusDisabling = HubStatus("disabling")
	// HubStatusDisabled is [insert doc].
	HubStatusDisabled = HubStatus("disabled")
)

func (enum HubStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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
	// ListDevicesRequestOrderByNameAsc is [insert doc].
	ListDevicesRequestOrderByNameAsc = ListDevicesRequestOrderBy("name_asc")
	// ListDevicesRequestOrderByNameDesc is [insert doc].
	ListDevicesRequestOrderByNameDesc = ListDevicesRequestOrderBy("name_desc")
	// ListDevicesRequestOrderByStatusAsc is [insert doc].
	ListDevicesRequestOrderByStatusAsc = ListDevicesRequestOrderBy("status_asc")
	// ListDevicesRequestOrderByStatusDesc is [insert doc].
	ListDevicesRequestOrderByStatusDesc = ListDevicesRequestOrderBy("status_desc")
	// ListDevicesRequestOrderByHubIDAsc is [insert doc].
	ListDevicesRequestOrderByHubIDAsc = ListDevicesRequestOrderBy("hub_id_asc")
	// ListDevicesRequestOrderByHubIDDesc is [insert doc].
	ListDevicesRequestOrderByHubIDDesc = ListDevicesRequestOrderBy("hub_id_desc")
	// ListDevicesRequestOrderByCreatedAtAsc is [insert doc].
	ListDevicesRequestOrderByCreatedAtAsc = ListDevicesRequestOrderBy("created_at_asc")
	// ListDevicesRequestOrderByCreatedAtDesc is [insert doc].
	ListDevicesRequestOrderByCreatedAtDesc = ListDevicesRequestOrderBy("created_at_desc")
	// ListDevicesRequestOrderByUpdatedAtAsc is [insert doc].
	ListDevicesRequestOrderByUpdatedAtAsc = ListDevicesRequestOrderBy("updated_at_asc")
	// ListDevicesRequestOrderByUpdatedAtDesc is [insert doc].
	ListDevicesRequestOrderByUpdatedAtDesc = ListDevicesRequestOrderBy("updated_at_desc")
	// ListDevicesRequestOrderByAllowInsecureAsc is [insert doc].
	ListDevicesRequestOrderByAllowInsecureAsc = ListDevicesRequestOrderBy("allow_insecure_asc")
	// ListDevicesRequestOrderByAllowInsecureDesc is [insert doc].
	ListDevicesRequestOrderByAllowInsecureDesc = ListDevicesRequestOrderBy("allow_insecure_desc")
)

func (enum ListDevicesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
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
	// ListHubsRequestOrderByNameAsc is [insert doc].
	ListHubsRequestOrderByNameAsc = ListHubsRequestOrderBy("name_asc")
	// ListHubsRequestOrderByNameDesc is [insert doc].
	ListHubsRequestOrderByNameDesc = ListHubsRequestOrderBy("name_desc")
	// ListHubsRequestOrderByStatusAsc is [insert doc].
	ListHubsRequestOrderByStatusAsc = ListHubsRequestOrderBy("status_asc")
	// ListHubsRequestOrderByStatusDesc is [insert doc].
	ListHubsRequestOrderByStatusDesc = ListHubsRequestOrderBy("status_desc")
	// ListHubsRequestOrderByProductPlanAsc is [insert doc].
	ListHubsRequestOrderByProductPlanAsc = ListHubsRequestOrderBy("product_plan_asc")
	// ListHubsRequestOrderByProductPlanDesc is [insert doc].
	ListHubsRequestOrderByProductPlanDesc = ListHubsRequestOrderBy("product_plan_desc")
	// ListHubsRequestOrderByCreatedAtAsc is [insert doc].
	ListHubsRequestOrderByCreatedAtAsc = ListHubsRequestOrderBy("created_at_asc")
	// ListHubsRequestOrderByCreatedAtDesc is [insert doc].
	ListHubsRequestOrderByCreatedAtDesc = ListHubsRequestOrderBy("created_at_desc")
	// ListHubsRequestOrderByUpdatedAtAsc is [insert doc].
	ListHubsRequestOrderByUpdatedAtAsc = ListHubsRequestOrderBy("updated_at_asc")
	// ListHubsRequestOrderByUpdatedAtDesc is [insert doc].
	ListHubsRequestOrderByUpdatedAtDesc = ListHubsRequestOrderBy("updated_at_desc")
)

func (enum ListHubsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
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
	// ListNetworksRequestOrderByNameAsc is [insert doc].
	ListNetworksRequestOrderByNameAsc = ListNetworksRequestOrderBy("name_asc")
	// ListNetworksRequestOrderByNameDesc is [insert doc].
	ListNetworksRequestOrderByNameDesc = ListNetworksRequestOrderBy("name_desc")
	// ListNetworksRequestOrderByTypeAsc is [insert doc].
	ListNetworksRequestOrderByTypeAsc = ListNetworksRequestOrderBy("type_asc")
	// ListNetworksRequestOrderByTypeDesc is [insert doc].
	ListNetworksRequestOrderByTypeDesc = ListNetworksRequestOrderBy("type_desc")
	// ListNetworksRequestOrderByCreatedAtAsc is [insert doc].
	ListNetworksRequestOrderByCreatedAtAsc = ListNetworksRequestOrderBy("created_at_asc")
	// ListNetworksRequestOrderByCreatedAtDesc is [insert doc].
	ListNetworksRequestOrderByCreatedAtDesc = ListNetworksRequestOrderBy("created_at_desc")
)

func (enum ListNetworksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
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
	// ListRoutesRequestOrderByNameAsc is [insert doc].
	ListRoutesRequestOrderByNameAsc = ListRoutesRequestOrderBy("name_asc")
	// ListRoutesRequestOrderByNameDesc is [insert doc].
	ListRoutesRequestOrderByNameDesc = ListRoutesRequestOrderBy("name_desc")
	// ListRoutesRequestOrderByHubIDAsc is [insert doc].
	ListRoutesRequestOrderByHubIDAsc = ListRoutesRequestOrderBy("hub_id_asc")
	// ListRoutesRequestOrderByHubIDDesc is [insert doc].
	ListRoutesRequestOrderByHubIDDesc = ListRoutesRequestOrderBy("hub_id_desc")
	// ListRoutesRequestOrderByTypeAsc is [insert doc].
	ListRoutesRequestOrderByTypeAsc = ListRoutesRequestOrderBy("type_asc")
	// ListRoutesRequestOrderByTypeDesc is [insert doc].
	ListRoutesRequestOrderByTypeDesc = ListRoutesRequestOrderBy("type_desc")
	// ListRoutesRequestOrderByCreatedAtAsc is [insert doc].
	ListRoutesRequestOrderByCreatedAtAsc = ListRoutesRequestOrderBy("created_at_asc")
	// ListRoutesRequestOrderByCreatedAtDesc is [insert doc].
	ListRoutesRequestOrderByCreatedAtDesc = ListRoutesRequestOrderBy("created_at_desc")
)

func (enum ListRoutesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
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
	// NetworkNetworkTypeUnknown is [insert doc].
	NetworkNetworkTypeUnknown = NetworkNetworkType("unknown")
	// NetworkNetworkTypeSigfox is [insert doc].
	NetworkNetworkTypeSigfox = NetworkNetworkType("sigfox")
	// NetworkNetworkTypeRest is [insert doc].
	NetworkNetworkTypeRest = NetworkNetworkType("rest")
)

func (enum NetworkNetworkType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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

type RouteRestConfigHTTPVerb string

const (
	// RouteRestConfigHTTPVerbUnknown is [insert doc].
	RouteRestConfigHTTPVerbUnknown = RouteRestConfigHTTPVerb("unknown")
	// RouteRestConfigHTTPVerbGet is [insert doc].
	RouteRestConfigHTTPVerbGet = RouteRestConfigHTTPVerb("get")
	// RouteRestConfigHTTPVerbPost is [insert doc].
	RouteRestConfigHTTPVerbPost = RouteRestConfigHTTPVerb("post")
	// RouteRestConfigHTTPVerbPut is [insert doc].
	RouteRestConfigHTTPVerbPut = RouteRestConfigHTTPVerb("put")
	// RouteRestConfigHTTPVerbPatch is [insert doc].
	RouteRestConfigHTTPVerbPatch = RouteRestConfigHTTPVerb("patch")
	// RouteRestConfigHTTPVerbDelete is [insert doc].
	RouteRestConfigHTTPVerbDelete = RouteRestConfigHTTPVerb("delete")
)

func (enum RouteRestConfigHTTPVerb) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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
	// RouteRouteTypeUnknown is [insert doc].
	RouteRouteTypeUnknown = RouteRouteType("unknown")
	// RouteRouteTypeS3 is [insert doc].
	RouteRouteTypeS3 = RouteRouteType("s3")
	// RouteRouteTypeDatabase is [insert doc].
	RouteRouteTypeDatabase = RouteRouteType("database")
	// RouteRouteTypeRest is [insert doc].
	RouteRouteTypeRest = RouteRouteType("rest")
)

func (enum RouteRouteType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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
	// RouteS3ConfigS3StrategyUnknown is [insert doc].
	RouteS3ConfigS3StrategyUnknown = RouteS3ConfigS3Strategy("unknown")
	// RouteS3ConfigS3StrategyPerTopic is [insert doc].
	RouteS3ConfigS3StrategyPerTopic = RouteS3ConfigS3Strategy("per_topic")
	// RouteS3ConfigS3StrategyPerMessage is [insert doc].
	RouteS3ConfigS3StrategyPerMessage = RouteS3ConfigS3Strategy("per_message")
)

func (enum RouteS3ConfigS3Strategy) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
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

type Certificate struct {
	Crt string `json:"crt"`

	Key string `json:"key"`
}

// CreateDeviceResponse: create device response
type CreateDeviceResponse struct {
	// Device: created device information
	Device *Device `json:"device"`
	// Certificate: device certificate
	Certificate *Certificate `json:"certificate"`
}

// CreateNetworkResponse: create network response
type CreateNetworkResponse struct {
	// Network: created network
	Network *Network `json:"network"`
	// Secret: endpoint Key to keep secret. This cannot be retrieved later
	Secret string `json:"secret"`
}

type CreateRouteRequestDatabaseConfig struct {
	Host string `json:"host"`

	Port uint32 `json:"port"`

	Dbname string `json:"dbname"`

	Username string `json:"username"`

	Password string `json:"password"`

	Query string `json:"query"`
}

type CreateRouteRequestRestConfig struct {
	// Verb:
	//
	// Default value: unknown
	Verb RouteRestConfigHTTPVerb `json:"verb"`

	URI string `json:"uri"`

	Headers map[string]string `json:"headers"`
}

type CreateRouteRequestS3Config struct {
	BucketRegion string `json:"bucket_region"`

	BucketName string `json:"bucket_name"`

	ObjectPrefix string `json:"object_prefix"`
	// Strategy:
	//
	// Default value: unknown
	Strategy RouteS3ConfigS3Strategy `json:"strategy"`
}

// Device: device
type Device struct {
	// ID: device ID, also used as MQTT Client ID or Username
	ID string `json:"id"`
	// Name: device name
	Name string `json:"name"`
	// Description: device description
	Description string `json:"description"`
	// Status: device status
	//
	// Default value: unknown
	Status DeviceStatus `json:"status"`
	// HubID: hub ID
	HubID string `json:"hub_id"`
	// LastActivityAt: device last connection/activity date
	LastActivityAt *time.Time `json:"last_activity_at"`
	// IsConnected: whether the device is connected to the Hub or not
	IsConnected bool `json:"is_connected"`
	// AllowInsecure: whether to allow device to connect without TLS mutual authentication
	AllowInsecure bool `json:"allow_insecure"`
	// AllowMultipleConnections: whether to allow multiple physical devices to connect with this device's credentials
	AllowMultipleConnections bool `json:"allow_multiple_connections"`
	// MessageFilters: filter-sets to restrict the topics the device can publish/subscribe to
	MessageFilters *DeviceMessageFilters `json:"message_filters"`
	// CreatedAt: device add date
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: device last modification date
	UpdatedAt *time.Time `json:"updated_at"`
}

// DeviceMessageFilters: device. message filters
type DeviceMessageFilters struct {
	// Publish: filtering rule to restrict topics the device can publish to
	Publish *DeviceMessageFiltersRule `json:"publish"`
	// Subscribe: filtering rule to restrict topics the device can subscribe to
	Subscribe *DeviceMessageFiltersRule `json:"subscribe"`
}

type DeviceMessageFiltersRule struct {
	// Policy:
	//
	// Default value: unknown
	Policy DeviceMessageFiltersRulePolicy `json:"policy"`

	Topics *[]string `json:"topics"`
}

// GetDeviceMetricsResponse: get device metrics response
type GetDeviceMetricsResponse struct {
	// Metrics: metrics for a device over the requested period
	Metrics []*scw.TimeSeries `json:"metrics"`
}

// GetHubMetricsResponse: get hub metrics response
type GetHubMetricsResponse struct {
	// Metrics: metrics for a hub over the requested period
	Metrics []*scw.TimeSeries `json:"metrics"`
}

// Hub: hub
type Hub struct {
	// ID: hub ID
	ID string `json:"id"`
	// Name: hub name
	Name string `json:"name"`
	// Status: current status of the Hub
	//
	// Default value: unknown
	Status HubStatus `json:"status"`
	// ProductPlan: hub feature set
	//
	// Default value: plan_unknown
	ProductPlan HubProductPlan `json:"product_plan"`
	// Enabled: whether the hub has been enabled
	Enabled bool `json:"enabled"`
	// DeviceCount: number of registered devices
	DeviceCount uint64 `json:"device_count"`
	// ConnectedDeviceCount: number of currently connected devices
	ConnectedDeviceCount uint64 `json:"connected_device_count"`
	// Endpoint: host to connect your devices to
	//
	// Devices should be connected to this host, port may be 1883 (MQTT), 8883 (MQTT over TLS), 80 (MQTT over Websocket) or 443 (MQTT over Websocket over TLS).
	Endpoint string `json:"endpoint"`
	// DisableEvents: disable Hub events
	DisableEvents bool `json:"disable_events"`
	// EventsTopicPrefix: hub events topic prefix
	EventsTopicPrefix string `json:"events_topic_prefix"`
	// Region: region of the Hub
	Region scw.Region `json:"region"`
	// CreatedAt: hub creation date
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: hub last modification date
	UpdatedAt *time.Time `json:"updated_at"`
	// ProjectID: project owning the resource
	ProjectID string `json:"project_id"`
	// OrganizationID: organization owning the resource
	OrganizationID string `json:"organization_id"`
	// EnableDeviceAutoProvisioning: enable device auto provisioning
	//
	// When an unknown device connects to your hub using a valid certificate chain, it will be automatically provisioned inside your hub. The hub uses the common name of the device certifcate to find out if a device with the same name already exists. This setting can only be enabled on a hub with a custom certificate authority.
	EnableDeviceAutoProvisioning bool `json:"enable_device_auto_provisioning"`
}

// ListDevicesResponse: list devices response
type ListDevicesResponse struct {
	// TotalCount: total number of devices
	TotalCount uint32 `json:"total_count"`
	// Devices: a page of devices
	Devices []*Device `json:"devices"`
}

// ListHubsResponse: list hubs response
type ListHubsResponse struct {
	// TotalCount: total number of hubs
	TotalCount uint32 `json:"total_count"`
	// Hubs: a page of hubs
	Hubs []*Hub `json:"hubs"`
}

// ListNetworksResponse: list networks response
type ListNetworksResponse struct {
	// TotalCount: total number of Networks
	TotalCount uint32 `json:"total_count"`
	// Networks: a page of networks
	Networks []*Network `json:"networks"`
}

// ListRoutesResponse: list routes response
type ListRoutesResponse struct {
	// TotalCount: total number of routes
	TotalCount uint32 `json:"total_count"`
	// Routes: a page of routes
	Routes []*RouteSummary `json:"routes"`
}

// Network: network
type Network struct {
	// ID: network ID
	ID string `json:"id"`
	// Name: network name
	Name string `json:"name"`
	// Type: type of network to connect with
	//
	// Default value: unknown
	Type NetworkNetworkType `json:"type"`
	// Endpoint: endpoint to use for interacting with the network
	Endpoint string `json:"endpoint"`
	// HubID: hub ID to connect the Network to
	HubID string `json:"hub_id"`
	// CreatedAt: network creation date
	CreatedAt *time.Time `json:"created_at"`
	// TopicPrefix: topic prefix for the Network
	//
	// This prefix will be prepended to all topics for this Network.
	TopicPrefix string `json:"topic_prefix"`
}

// RenewDeviceCertificateResponse: renew device certificate response
type RenewDeviceCertificateResponse struct {
	// Device: created device information
	Device *Device `json:"device"`
	// Certificate: device certificate
	Certificate *Certificate `json:"certificate"`
}

// Route: route
type Route struct {
	// ID: route ID
	ID string `json:"id"`
	// Name: route name
	Name string `json:"name"`
	// HubID: ID of the route's hub
	HubID string `json:"hub_id"`
	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters
	Topic string `json:"topic"`
	// Type: route type
	//
	// Default value: unknown
	Type RouteRouteType `json:"type"`
	// CreatedAt: route creation date
	CreatedAt *time.Time `json:"created_at"`
	// S3Config: when using S3 Route, S3-specific configuration fields
	// Precisely one of DbConfig, RestConfig, S3Config must be set.
	S3Config *RouteS3Config `json:"s3_config,omitempty"`
	// DbConfig: when using Database Route, DB-specific configuration fields
	// Precisely one of DbConfig, RestConfig, S3Config must be set.
	DbConfig *RouteDatabaseConfig `json:"db_config,omitempty"`
	// RestConfig: when using Rest Route, Rest-specific configuration fields
	// Precisely one of DbConfig, RestConfig, S3Config must be set.
	RestConfig *RouteRestConfig `json:"rest_config,omitempty"`
}

// RouteDatabaseConfig: route. database config
type RouteDatabaseConfig struct {
	// Host: database host
	Host string `json:"host"`
	// Port: database port
	Port uint32 `json:"port"`
	// Dbname: database name
	Dbname string `json:"dbname"`
	// Username: database username. Make sure this account can execute the provided query
	Username string `json:"username"`
	// Password: database password
	Password string `json:"password"`
	// Query: SQL query to be executed ($TOPIC and $PAYLOAD variables are available, see documentation)
	Query string `json:"query"`
}

// RouteRestConfig: route. rest config
type RouteRestConfig struct {
	// Verb: HTTP Verb used to call REST URI
	//
	// Default value: unknown
	Verb RouteRestConfigHTTPVerb `json:"verb"`
	// URI: URI of the REST endpoint
	URI string `json:"uri"`
	// Headers: HTTP call extra headers
	Headers map[string]string `json:"headers"`
}

// RouteS3Config: route.s3 config
type RouteS3Config struct {
	// BucketRegion: region of the S3 route's destination bucket (eg 'fr-par')
	BucketRegion string `json:"bucket_region"`
	// BucketName: name of the S3 route's destination bucket
	BucketName string `json:"bucket_name"`
	// ObjectPrefix: optional string to prefix object names with
	ObjectPrefix string `json:"object_prefix"`
	// Strategy: how the S3 route's objects will be created: one per topic or one per message
	//
	// Default value: unknown
	Strategy RouteS3ConfigS3Strategy `json:"strategy"`
}

// RouteSummary: route summary
type RouteSummary struct {
	// ID: route ID
	ID string `json:"id"`
	// Name: route name
	Name string `json:"name"`
	// HubID: ID of the route's hub
	HubID string `json:"hub_id"`
	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters
	Topic string `json:"topic"`
	// Type: route type
	//
	// Default value: unknown
	Type RouteRouteType `json:"type"`
	// CreatedAt: route creation date
	CreatedAt *time.Time `json:"created_at"`
}

// Service API

type GetServiceInfoRequest struct {
	Region scw.Region `json:"-"`
}

func (s *API) GetServiceInfo(req *GetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "",
		Headers: http.Header{},
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListHubsRequest struct {
	Region scw.Region `json:"-"`
	// Page: page number
	Page *int32 `json:"-"`
	// PageSize: page size. The maximum value is 100
	PageSize *uint32 `json:"-"`
	// OrderBy: ordering of requested hub
	//
	// Default value: name_asc
	OrderBy ListHubsRequestOrderBy `json:"-"`
	// ProjectID: filter on project
	ProjectID *string `json:"-"`
	// OrganizationID: filter on the organization
	OrganizationID *string `json:"-"`
	// Name: filter on the name
	Name *string `json:"-"`
}

// ListHubs: list hubs
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
		Method:  "GET",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListHubsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListHubsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListHubsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListHubsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Hubs = append(r.Hubs, results.Hubs...)
	r.TotalCount += uint32(len(results.Hubs))
	return uint32(len(results.Hubs)), nil
}

type CreateHubRequest struct {
	Region scw.Region `json:"-"`
	// Name: hub name (up to 255 characters)
	Name string `json:"name"`
	// ProjectID: organization/project owning the resource
	ProjectID string `json:"project_id"`
	// ProductPlan: hub feature set
	//
	// Default value: plan_shared
	ProductPlan HubProductPlan `json:"product_plan"`
	// DisableEvents: disable Hub events
	DisableEvents *bool `json:"disable_events"`
	// EventsTopicPrefix: hub events topic prefix (default '$SCW/events')
	EventsTopicPrefix *string `json:"events_topic_prefix"`
}

// CreateHub: create a hub
func (s *API) CreateHub(req *CreateHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("hub")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs",
		Headers: http.Header{},
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

type GetHubRequest struct {
	Region scw.Region `json:"-"`
	// HubID: hub ID
	HubID string `json:"-"`
}

// GetHub: get a hub
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
		Method:  "GET",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "",
		Headers: http.Header{},
	}

	var resp Hub

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateHubRequest struct {
	Region scw.Region `json:"-"`
	// HubID: hub ID
	HubID string `json:"-"`
	// Name: hub name (up to 255 characters)
	Name *string `json:"name"`
	// ProductPlan: hub feature set
	//
	// Default value: plan_unknown
	ProductPlan HubProductPlan `json:"product_plan"`
	// DisableEvents: disable Hub events
	DisableEvents *bool `json:"disable_events"`
	// EventsTopicPrefix: hub events topic prefix
	EventsTopicPrefix *string `json:"events_topic_prefix"`
	// EnableDeviceAutoProvisioning: enable device auto provisioning
	EnableDeviceAutoProvisioning *bool `json:"enable_device_auto_provisioning"`
}

// UpdateHub: update a hub
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
		Method:  "PATCH",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "",
		Headers: http.Header{},
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

type EnableHubRequest struct {
	Region scw.Region `json:"-"`
	// HubID: hub ID
	HubID string `json:"-"`
}

// EnableHub: enable a hub
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
		Method:  "POST",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/enable",
		Headers: http.Header{},
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

type DisableHubRequest struct {
	Region scw.Region `json:"-"`
	// HubID: hub ID
	HubID string `json:"-"`
}

// DisableHub: disable a hub
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
		Method:  "POST",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/disable",
		Headers: http.Header{},
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

type DeleteHubRequest struct {
	Region scw.Region `json:"-"`
	// HubID: hub ID
	HubID string `json:"-"`
	// DeleteDevices: force deletion of devices added to this hub instead of rejecting operation
	DeleteDevices *bool `json:"-"`
}

// DeleteHub: delete a hub
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
		Method:  "DELETE",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "",
		Query:   query,
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetHubMetricsRequest struct {
	Region scw.Region `json:"-"`
	// HubID: hub ID
	HubID string `json:"-"`
	// StartDate: start date used to compute the best scale for the returned metrics
	//
	// Default value: hour
	StartDate *time.Time `json:"-"`
}

// GetHubMetrics: get a hub's metrics
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
		Method:  "GET",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/metrics",
		Query:   query,
		Headers: http.Header{},
	}

	var resp GetHubMetricsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetHubCARequest struct {
	Region scw.Region `json:"-"`
	// HubID: hub ID
	HubID string `json:"-"`
	// CaCertPem: the CA's PEM-encoded certificate
	CaCertPem string `json:"ca_cert_pem"`
	// ChallengeCertPem: proof of possession PEM-encoded certificate
	//
	// The challenge is a PEM-encoded certificate to prove the possession of the CA. It must be signed by the CA, and have a Common Name equal to the Hub ID.
	ChallengeCertPem string `json:"challenge_cert_pem"`
}

// SetHubCA: set the certificate authority of a hub
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
		Method:  "POST",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/ca",
		Headers: http.Header{},
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

type ListDevicesRequest struct {
	Region scw.Region `json:"-"`
	// Page: page number
	Page *int32 `json:"-"`
	// PageSize: page size. The maximum value is 100
	PageSize *uint32 `json:"-"`
	// OrderBy: ordering of requested devices
	//
	// Default value: name_asc
	OrderBy ListDevicesRequestOrderBy `json:"-"`
	// Name: filter on the name
	Name *string `json:"-"`
	// HubID: filter on the hub
	HubID *string `json:"-"`
	// AllowInsecure: filter on the allow_insecure flag
	AllowInsecure *bool `json:"-"`
	// Status: device status (enabled, disabled, etc.)
	//
	// Default value: unknown
	Status DeviceStatus `json:"-"`
}

// ListDevices: list devices
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
		Method:  "GET",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListDevicesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDevicesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDevicesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListDevicesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Devices = append(r.Devices, results.Devices...)
	r.TotalCount += uint32(len(results.Devices))
	return uint32(len(results.Devices)), nil
}

type CreateDeviceRequest struct {
	Region scw.Region `json:"-"`
	// Name: device name
	Name string `json:"name"`
	// HubID: ID of the device's hub
	HubID string `json:"hub_id"`
	// AllowInsecure: allow plain and server-authenticated SSL connections in addition to mutually-authenticated ones
	AllowInsecure bool `json:"allow_insecure"`
	// AllowMultipleConnections: allow multiple physical devices to connect with this device's credentials
	AllowMultipleConnections bool `json:"allow_multiple_connections"`
	// MessageFilters: filter-sets to authorize or deny the device to publish/subscribe to specific topics
	MessageFilters *DeviceMessageFilters `json:"message_filters"`
	// Description: device description
	Description *string `json:"description"`
}

// CreateDevice: add a device
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
		Method:  "POST",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices",
		Headers: http.Header{},
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

type GetDeviceRequest struct {
	Region scw.Region `json:"-"`
	// DeviceID: device ID
	DeviceID string `json:"-"`
}

// GetDevice: get a device
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
		Method:  "GET",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "",
		Headers: http.Header{},
	}

	var resp Device

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateDeviceRequest struct {
	Region scw.Region `json:"-"`
	// DeviceID: device ID
	DeviceID string `json:"-"`
	// Description: device description
	Description *string `json:"description"`
	// AllowInsecure: allow plain and server-authenticated SSL connections in addition to mutually-authenticated ones
	AllowInsecure *bool `json:"allow_insecure"`
	// AllowMultipleConnections: allow multiple physical devices to connect with this device's credentials
	AllowMultipleConnections *bool `json:"allow_multiple_connections"`
	// MessageFilters: filter-sets to restrict the topics the device can publish/subscribe to
	MessageFilters *DeviceMessageFilters `json:"message_filters"`
	// HubID: change Hub for this device, additional fees may apply, see IoT Hub pricing
	HubID *string `json:"hub_id"`
}

// UpdateDevice: update a device
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
		Method:  "PATCH",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "",
		Headers: http.Header{},
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

type EnableDeviceRequest struct {
	Region scw.Region `json:"-"`
	// DeviceID: device ID
	DeviceID string `json:"-"`
}

// EnableDevice: enable a device
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
		Method:  "POST",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/enable",
		Headers: http.Header{},
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

type DisableDeviceRequest struct {
	Region scw.Region `json:"-"`
	// DeviceID: device ID
	DeviceID string `json:"-"`
}

// DisableDevice: disable a device
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
		Method:  "POST",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/disable",
		Headers: http.Header{},
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

type RenewDeviceCertificateRequest struct {
	Region scw.Region `json:"-"`
	// DeviceID: device ID
	DeviceID string `json:"-"`
}

// RenewDeviceCertificate: renew a device certificate
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
		Method:  "POST",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/renew-certificate",
		Headers: http.Header{},
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

type DeleteDeviceRequest struct {
	Region scw.Region `json:"-"`
	// DeviceID: device ID
	DeviceID string `json:"-"`
}

// DeleteDevice: remove a device
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
		Method:  "DELETE",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetDeviceMetricsRequest struct {
	Region scw.Region `json:"-"`
	// DeviceID: device ID
	DeviceID string `json:"-"`
	// StartDate: start date used to compute the best scale for the returned metrics
	//
	// Default value: hour
	StartDate *time.Time `json:"-"`
}

// GetDeviceMetrics: get a device's metrics
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
		Method:  "GET",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/metrics",
		Query:   query,
		Headers: http.Header{},
	}

	var resp GetDeviceMetricsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListRoutesRequest struct {
	Region scw.Region `json:"-"`
	// Page: page number
	Page *int32 `json:"-"`
	// PageSize: page size. The maximum value is 100
	PageSize *uint32 `json:"-"`
	// OrderBy: ordering of requested routes
	//
	// Default value: name_asc
	OrderBy ListRoutesRequestOrderBy `json:"-"`
	// HubID: filter on the hub
	HubID *string `json:"-"`
	// Name: filter on route's name
	Name *string `json:"-"`
}

// ListRoutes: list routes
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
		Method:  "GET",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListRoutesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListRoutesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListRoutesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListRoutesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Routes = append(r.Routes, results.Routes...)
	r.TotalCount += uint32(len(results.Routes))
	return uint32(len(results.Routes)), nil
}

type CreateRouteRequest struct {
	Region scw.Region `json:"-"`
	// Name: route name
	Name string `json:"name"`
	// HubID: ID of the route's hub
	HubID string `json:"hub_id"`
	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters
	Topic string `json:"topic"`
	// S3Config: if creating S3 Route, S3-specific configuration fields
	// Precisely one of DbConfig, RestConfig, S3Config must be set.
	S3Config *CreateRouteRequestS3Config `json:"s3_config,omitempty"`
	// DbConfig: if creating Database Route, DB-specific configuration fields
	// Precisely one of DbConfig, RestConfig, S3Config must be set.
	DbConfig *CreateRouteRequestDatabaseConfig `json:"db_config,omitempty"`
	// RestConfig: if creating Rest Route, Rest-specific configuration fields
	// Precisely one of DbConfig, RestConfig, S3Config must be set.
	RestConfig *CreateRouteRequestRestConfig `json:"rest_config,omitempty"`
}

// CreateRoute: create a route
//
// Multiple route kinds can be created:
// - Database Route.
//   Create a route that will record subscribed MQTT messages into your database.
//   <b>You need to manage the database by yourself</b>.
// - REST Route.
//   Create a route that will call a REST API on received subscribed MQTT messages.
// - S3 Routes.
//   Create a route that will put subscribed MQTT messages into an S3 bucket.
//   You need to create the bucket yourself and grant us write access.
//   The grant can be done with s3cmd (`s3cmd setacl s3://<my-bucket> --acl-grant=write:555c69c3-87d0-4bf8-80f1-99a2f757d031:555c69c3-87d0-4bf8-80f1-99a2f757d031`).
//
func (s *API) CreateRoute(req *CreateRouteRequest, opts ...scw.RequestOption) (*Route, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes",
		Headers: http.Header{},
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

type GetRouteRequest struct {
	Region scw.Region `json:"-"`
	// RouteID: route ID
	RouteID string `json:"-"`
}

// GetRoute: get a route
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
		Method:  "GET",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
		Headers: http.Header{},
	}

	var resp Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteRouteRequest struct {
	Region scw.Region `json:"-"`
	// RouteID: route ID
	RouteID string `json:"-"`
}

// DeleteRoute: delete a route
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
		Method:  "DELETE",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/routes/" + fmt.Sprint(req.RouteID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListNetworksRequest struct {
	Region scw.Region `json:"-"`
	// Page: page number
	Page *int32 `json:"-"`
	// PageSize: page size. The maximum value is 100
	PageSize *uint32 `json:"-"`
	// OrderBy: ordering of requested routes
	//
	// Default value: name_asc
	OrderBy ListNetworksRequestOrderBy `json:"-"`
	// Name: filter on Network name
	Name *string `json:"-"`
	// HubID: filter on the hub
	HubID *string `json:"-"`
	// TopicPrefix: filter on the topic prefix
	TopicPrefix *string `json:"-"`
}

// ListNetworks: list the Networks
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
		Method:  "GET",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/networks",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNetworksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNetworksResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListNetworksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Networks = append(r.Networks, results.Networks...)
	r.TotalCount += uint32(len(results.Networks))
	return uint32(len(results.Networks)), nil
}

type CreateNetworkRequest struct {
	Region scw.Region `json:"-"`
	// Name: network name
	Name string `json:"name"`
	// Type: type of network to connect with
	//
	// Default value: unknown
	Type NetworkNetworkType `json:"type"`
	// HubID: hub ID to connect the Network to
	HubID string `json:"hub_id"`
	// TopicPrefix: topic prefix for the Network
	TopicPrefix string `json:"topic_prefix"`
}

// CreateNetwork: create a new Network
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
		Method:  "POST",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/networks",
		Headers: http.Header{},
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

type GetNetworkRequest struct {
	Region scw.Region `json:"-"`
	// NetworkID: network ID
	NetworkID string `json:"-"`
}

// GetNetwork: retrieve a specific Network
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
		Method:  "GET",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/networks/" + fmt.Sprint(req.NetworkID) + "",
		Headers: http.Header{},
	}

	var resp Network

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteNetworkRequest struct {
	Region scw.Region `json:"-"`
	// NetworkID: network ID
	NetworkID string `json:"-"`
}

// DeleteNetwork: delete a Network
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
		Method:  "DELETE",
		Path:    "/iot/v1/regions/" + fmt.Sprint(req.Region) + "/networks/" + fmt.Sprint(req.NetworkID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

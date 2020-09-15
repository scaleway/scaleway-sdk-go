// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package iot provides methods and message types of the iot v1beta1 API.
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

type DeviceMessageFiltersPolicy string

const (
	// DeviceMessageFiltersPolicyUnknown is [insert doc].
	DeviceMessageFiltersPolicyUnknown = DeviceMessageFiltersPolicy("unknown")
	// DeviceMessageFiltersPolicyAccept is [insert doc].
	DeviceMessageFiltersPolicyAccept = DeviceMessageFiltersPolicy("accept")
	// DeviceMessageFiltersPolicyReject is [insert doc].
	DeviceMessageFiltersPolicyReject = DeviceMessageFiltersPolicy("reject")
)

func (enum DeviceMessageFiltersPolicy) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum DeviceMessageFiltersPolicy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DeviceMessageFiltersPolicy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DeviceMessageFiltersPolicy(DeviceMessageFiltersPolicy(tmp).String())
	return nil
}

type DeviceStatus string

const (
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
		return "error"
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

type HubStatus string

const (
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
		return "error"
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

type MetricsPeriod string

const (
	// MetricsPeriodHour is [insert doc].
	MetricsPeriodHour = MetricsPeriod("hour")
	// MetricsPeriodDay is [insert doc].
	MetricsPeriodDay = MetricsPeriod("day")
	// MetricsPeriodWeek is [insert doc].
	MetricsPeriodWeek = MetricsPeriod("week")
	// MetricsPeriodMonth is [insert doc].
	MetricsPeriodMonth = MetricsPeriod("month")
	// MetricsPeriodYear is [insert doc].
	MetricsPeriodYear = MetricsPeriod("year")
)

func (enum MetricsPeriod) String() string {
	if enum == "" {
		// return default value if empty
		return "hour"
	}
	return string(enum)
}

func (enum MetricsPeriod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MetricsPeriod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MetricsPeriod(MetricsPeriod(tmp).String())
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

type ProductPlan string

const (
	// ProductPlanPlanUnknown is [insert doc].
	ProductPlanPlanUnknown = ProductPlan("plan_unknown")
	// ProductPlanPlanShared is [insert doc].
	ProductPlanPlanShared = ProductPlan("plan_shared")
	// ProductPlanPlanDedicated is [insert doc].
	ProductPlanPlanDedicated = ProductPlan("plan_dedicated")
	// ProductPlanPlanHa is [insert doc].
	ProductPlanPlanHa = ProductPlan("plan_ha")
)

func (enum ProductPlan) String() string {
	if enum == "" {
		// return default value if empty
		return "plan_unknown"
	}
	return string(enum)
}

func (enum ProductPlan) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ProductPlan) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ProductPlan(ProductPlan(tmp).String())
	return nil
}

type RestRouteHTTPVerb string

const (
	// RestRouteHTTPVerbGet is [insert doc].
	RestRouteHTTPVerbGet = RestRouteHTTPVerb("get")
	// RestRouteHTTPVerbPost is [insert doc].
	RestRouteHTTPVerbPost = RestRouteHTTPVerb("post")
	// RestRouteHTTPVerbPut is [insert doc].
	RestRouteHTTPVerbPut = RestRouteHTTPVerb("put")
	// RestRouteHTTPVerbPatch is [insert doc].
	RestRouteHTTPVerbPatch = RestRouteHTTPVerb("patch")
	// RestRouteHTTPVerbDelete is [insert doc].
	RestRouteHTTPVerbDelete = RestRouteHTTPVerb("delete")
)

func (enum RestRouteHTTPVerb) String() string {
	if enum == "" {
		// return default value if empty
		return "get"
	}
	return string(enum)
}

func (enum RestRouteHTTPVerb) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RestRouteHTTPVerb) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RestRouteHTTPVerb(RestRouteHTTPVerb(tmp).String())
	return nil
}

type RouteRouteType string

const (
	// RouteRouteTypeUnknown is [insert doc].
	RouteRouteTypeUnknown = RouteRouteType("unknown")
	// RouteRouteTypeS3 is [insert doc].
	RouteRouteTypeS3 = RouteRouteType("s3")
	// RouteRouteTypeFunctions is [insert doc].
	RouteRouteTypeFunctions = RouteRouteType("functions")
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

type S3RouteStrategy string

const (
	// S3RouteStrategyPerTopic is [insert doc].
	S3RouteStrategyPerTopic = S3RouteStrategy("per_topic")
	// S3RouteStrategyPerMessage is [insert doc].
	S3RouteStrategyPerMessage = S3RouteStrategy("per_message")
)

func (enum S3RouteStrategy) String() string {
	if enum == "" {
		// return default value if empty
		return "per_topic"
	}
	return string(enum)
}

func (enum S3RouteStrategy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *S3RouteStrategy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = S3RouteStrategy(S3RouteStrategy(tmp).String())
	return nil
}

// CreateDeviceResponse: create device response
type CreateDeviceResponse struct {
	// Device: created device information
	Device *Device `json:"device"`
	// Crt: device certificate
	Crt string `json:"crt"`
	// Key: device certificate key
	Key string `json:"key"`
}

// CreateNetworkResponse: create network response
type CreateNetworkResponse struct {
	// Network: created network
	Network *Network `json:"network"`
	// Secret: endpoint Key to keep secret. This cannot be retrieved later
	Secret string `json:"secret"`
}

// DatabaseRoute: database route
type DatabaseRoute struct {
	// ID: route ID
	ID string `json:"id"`
	// Name: route name
	Name string `json:"name"`
	// OrganizationID: organization owning the route
	OrganizationID string `json:"organization_id"`
	// HubID: ID of the route's hub
	HubID string `json:"hub_id"`
	// Topic: topic the route subscribes to (wildcards allowed). It must be a valid MQTT topic and up to 65535 characters
	Topic string `json:"topic"`
	// CreatedAt: route creation date
	CreatedAt *time.Time `json:"created_at"`
	// Query: SQL query to be executed ($TOPIC and $PAYLOAD variables are available, see documentation)
	Query string `json:"query"`
	// Database: database settings
	Database *DatabaseSettings `json:"database"`
}

// DatabaseSettings: database settings
type DatabaseSettings struct {
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
}

// Device: device
type Device struct {
	// ID: device ID, also used as MQTT Client ID or Username
	ID string `json:"id"`
	// Name: device name
	Name string `json:"name"`
	// Status: device status
	//
	// Default value: error
	Status DeviceStatus `json:"status"`
	// HubID: hub ID
	HubID string `json:"hub_id"`
	// LastActivityAt: device last connection/activity date
	LastActivityAt *time.Time `json:"last_activity_at"`
	// IsConnected: whether the device is connected to the Hub or not
	IsConnected bool `json:"is_connected"`
	// AllowInsecure: whether to allow device to connect without TLS mutual authentication
	AllowInsecure bool `json:"allow_insecure"`
	// MessageFilters: filter-sets to restrict the topics the device can publish/subscribe to
	MessageFilters *DeviceMessageFilters `json:"message_filters"`
	// CreatedAt: device add date
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: device last modification date
	UpdatedAt *time.Time `json:"updated_at"`
	// OrganizationID: organization owning the resource
	OrganizationID string `json:"organization_id"`
}

// DeviceMessageFilters: device message filters
type DeviceMessageFilters struct {
	// Publish: filter-set to restrict topics the device can publish to
	Publish *DeviceMessageFiltersMessageFilterSet `json:"publish"`
	// Subscribe: filter-set to restrict topics the device can subscribe to
	Subscribe *DeviceMessageFiltersMessageFilterSet `json:"subscribe"`
}

type DeviceMessageFiltersMessageFilterSet struct {
	// Policy:
	//
	// Default value: unknown
	Policy DeviceMessageFiltersPolicy `json:"policy"`

	Topics *[]string `json:"topics"`
}

// FunctionsRoute: functions route
type FunctionsRoute struct {
	// ID: route ID
	ID string `json:"id"`
	// Name: route name
	Name string `json:"name"`
	// OrganizationID: organization owning the route
	OrganizationID string `json:"organization_id"`
	// HubID: ID of the route's hub
	HubID string `json:"hub_id"`
	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters
	Topic string `json:"topic"`
	// CreatedAt: route creation date
	CreatedAt *time.Time `json:"created_at"`
	// URI: uri of the function
	URI string `json:"uri"`
}

// Hub: hub
type Hub struct {
	// ID: hub ID
	ID string `json:"id"`
	// Name: hub name
	Name string `json:"name"`
	// Status: current status of the Hub
	//
	// Default value: error
	Status HubStatus `json:"status"`
	// ProductPlan: hub feature set
	//
	// Default value: plan_unknown
	ProductPlan ProductPlan `json:"product_plan"`
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
	// EventsEnabled: wether Hub events are enabled or not
	EventsEnabled bool `json:"events_enabled"`
	// EventsTopicPrefix: hub events topic prefix
	EventsTopicPrefix string `json:"events_topic_prefix"`
	// Region: region of the Hub
	Region scw.Region `json:"region"`
	// CreatedAt: hub creation date
	CreatedAt *time.Time `json:"created_at"`
	// UpdatedAt: hub last modification date
	UpdatedAt *time.Time `json:"updated_at"`
	// OrganizationID: organization owning the resource
	OrganizationID string `json:"organization_id"`
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
	Routes []*Route `json:"routes"`
}

// Metrics: metrics
type Metrics struct {
	// Metrics: metrics for a hub or a device, over the requested period
	Metrics []*MetricsMetric `json:"metrics"`
}

// MetricsMetric: metrics. metric
type MetricsMetric struct {
	// Name: metric name
	Name string `json:"name"`
	// Values: metric values over the selected period
	Values []*MetricsMetricValue `json:"values"`
}

// MetricsMetricValue: metrics. metric. value
type MetricsMetricValue struct {
	// Time: timestamp for the value
	Time *time.Time `json:"time"`
	// Value: numeric value
	Value int64 `json:"value"`
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
	// Region: region of the Network
	Region scw.Region `json:"region"`
	// OrganizationID: organization owning the resource
	OrganizationID string `json:"organization_id"`
}

// RestRoute: rest route
type RestRoute struct {
	// ID: route ID
	ID string `json:"id"`
	// Name: route name
	Name string `json:"name"`
	// OrganizationID: organization owning the route
	OrganizationID string `json:"organization_id"`
	// HubID: ID of the route's hub
	HubID string `json:"hub_id"`
	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters
	Topic string `json:"topic"`
	// CreatedAt: route creation date
	CreatedAt *time.Time `json:"created_at"`
	// Verb: HTTP Verb used to call REST URI
	//
	// Default value: get
	Verb RestRouteHTTPVerb `json:"verb"`
	// URI: URI of the REST endpoint
	URI string `json:"uri"`
	// Headers: HTTP call extra headers
	Headers map[string]string `json:"headers"`
}

// Route: route
type Route struct {
	// ID: route ID
	ID string `json:"id"`
	// Name: route name
	Name string `json:"name"`
	// OrganizationID: organization owning the resource
	OrganizationID string `json:"organization_id"`
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

// S3Route: s3 route
type S3Route struct {
	// ID: route ID
	ID string `json:"id"`
	// Name: route name
	Name string `json:"name"`
	// OrganizationID: organization owning the route
	OrganizationID string `json:"organization_id"`
	// HubID: ID of the route's hub
	HubID string `json:"hub_id"`
	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters
	Topic string `json:"topic"`
	// CreatedAt: route creation date
	CreatedAt *time.Time `json:"created_at"`
	// BucketRegion: region of the S3 route's destination bucket (eg 'fr-par')
	BucketRegion string `json:"bucket_region"`
	// BucketName: name of the S3 route's destination bucket
	BucketName string `json:"bucket_name"`
	// ObjectPrefix: optional string to prefix object names with
	ObjectPrefix string `json:"object_prefix"`
	// Strategy: how the S3 route's objects will be created: one per topic or one per message
	//
	// Default value: per_topic
	Strategy S3RouteStrategy `json:"strategy"`
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "",
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs",
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
	// OrganizationID: organization owning the resource
	OrganizationID string `json:"organization_id"`
	// ProductPlan: hub feature set
	//
	// Default value: plan_shared
	ProductPlan ProductPlan `json:"product_plan"`
	// DisableEvents: disable Hub events (default false)
	DisableEvents *bool `json:"disable_events"`
	// EventsTopicPrefix: hub events topic prefix (default '$SCW/events')
	EventsTopicPrefix *string `json:"events_topic_prefix"`
}

// CreateHub: create a hub
func (s *API) CreateHub(req *CreateHubRequest, opts ...scw.RequestOption) (*Hub, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs",
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "",
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
	ProductPlan ProductPlan `json:"product_plan"`
	// DisableEvents: disable Hub events
	DisableEvents *bool `json:"disable_events"`
	// EventsTopicPrefix: hub events topic prefix
	EventsTopicPrefix *string `json:"events_topic_prefix"`
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "",
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/enable",
		Headers: http.Header{},
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/disable",
		Headers: http.Header{},
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
	DeleteDevices *bool `json:"delete_devices"`
}

// DeleteHub: delete a hub
func (s *API) DeleteHub(req *DeleteHubRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "",
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

type GetHubMetricsRequest struct {
	Region scw.Region `json:"-"`
	// HubID: hub ID
	HubID string `json:"-"`
	// Period: period over which the metrics span
	//
	// Default value: hour
	Period MetricsPeriod `json:"-"`
}

// GetHubMetrics: get a hub's metrics
func (s *API) GetHubMetrics(req *GetHubMetricsRequest, opts ...scw.RequestOption) (*Metrics, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "period", req.Period)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.HubID) == "" {
		return nil, errors.New("field HubID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/hubs/" + fmt.Sprint(req.HubID) + "/metrics",
		Query:   query,
		Headers: http.Header{},
	}

	var resp Metrics

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
	// OrganizationID: filter on the organization
	OrganizationID *string `json:"-"`
	// Name: filter on the name
	Name *string `json:"-"`
	// HubID: filter on the hub
	HubID *string `json:"-"`
	// Deprecated: Enabled: deprecated, ignored filter
	Enabled *bool `json:"-"`
	// AllowInsecure: filter on the allow_insecure flag
	AllowInsecure *bool `json:"-"`
	// Deprecated: IsConnected: deprecated, ignored filter
	IsConnected *bool `json:"-"`
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "hub_id", req.HubID)
	parameter.AddToQuery(query, "enabled", req.Enabled)
	parameter.AddToQuery(query, "allow_insecure", req.AllowInsecure)
	parameter.AddToQuery(query, "is_connected", req.IsConnected)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/devices",
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
	// MessageFilters: filter-sets to authorize or deny the device to publish/subscribe to specific topics
	MessageFilters *DeviceMessageFilters `json:"message_filters"`
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/devices",
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "",
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
	// Name: device name
	Name *string `json:"name"`
	// AllowInsecure: allow plain and server-authenticated SSL connections in addition to mutually-authenticated ones
	AllowInsecure *bool `json:"allow_insecure"`
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "",
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/enable",
		Headers: http.Header{},
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/disable",
		Headers: http.Header{},
	}

	var resp Device

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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "",
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
	// Period: period over which the metrics span
	//
	// Default value: hour
	Period MetricsPeriod `json:"-"`
}

// GetDeviceMetrics: get a device's metrics
func (s *API) GetDeviceMetrics(req *GetDeviceMetricsRequest, opts ...scw.RequestOption) (*Metrics, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "period", req.Period)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DeviceID) == "" {
		return nil, errors.New("field DeviceID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/devices/" + fmt.Sprint(req.DeviceID) + "/metrics",
		Query:   query,
		Headers: http.Header{},
	}

	var resp Metrics

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
	// OrganizationID: filter on the organization
	OrganizationID *string `json:"-"`
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "hub_id", req.HubID)
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes",
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

type CreateS3RouteRequest struct {
	Region scw.Region `json:"-"`
	// Name: name of the route
	Name string `json:"name"`
	// HubID: ID of the route's hub
	HubID string `json:"hub_id"`
	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters
	Topic string `json:"topic"`
	// BucketRegion: region of the S3 route's destination bucket (eg 'fr-par')
	BucketRegion string `json:"bucket_region"`
	// BucketName: name of the S3 route's destination bucket
	BucketName string `json:"bucket_name"`
	// ObjectPrefix: optional string to prefix object names with
	ObjectPrefix string `json:"object_prefix"`
	// Strategy: how the S3 route's objects will be created: one per topic or one per message
	//
	// Default value: per_topic
	Strategy S3RouteStrategy `json:"strategy"`
}

// CreateS3Route: create an S3 route
//
// Create a route that will put subscribed MQTT messages into an S3 bucket.
// You need to create the bucket yourself and grant us write access:
// ```bash
// > s3cmd setacl s3://<my-bucket> --acl-grant=write:555c69c3-87d0-4bf8-80f1-99a2f757d031:555c69c3-87d0-4bf8-80f1-99a2f757d031
// ```
//
func (s *API) CreateS3Route(req *CreateS3RouteRequest, opts ...scw.RequestOption) (*S3Route, error) {
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes/s3",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp S3Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetS3RouteRequest struct {
	Region scw.Region `json:"-"`
	// RouteID: route ID
	RouteID string `json:"-"`
}

// GetS3Route: get an S3 route
func (s *API) GetS3Route(req *GetS3RouteRequest, opts ...scw.RequestOption) (*S3Route, error) {
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes/s3/" + fmt.Sprint(req.RouteID) + "",
		Headers: http.Header{},
	}

	var resp S3Route

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteS3RouteRequest struct {
	Region scw.Region `json:"-"`
	// RouteID: route ID
	RouteID string `json:"-"`
}

// DeleteS3Route: delete an S3 route
func (s *API) DeleteS3Route(req *DeleteS3RouteRequest, opts ...scw.RequestOption) error {
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes/s3/" + fmt.Sprint(req.RouteID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type CreateFunctionsRouteRequest struct {
	Region scw.Region `json:"-"`
	// Name: name of the route
	Name string `json:"name"`
	// HubID: ID of the route's hub
	HubID string `json:"hub_id"`
	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters
	Topic string `json:"topic"`
	// URI: uri of the function
	URI string `json:"uri"`
	// Token: authentication JWT token
	Token *string `json:"token"`
}

// CreateFunctionsRoute: create a Functions route
//
// Create a route that will forward subscribed MQTT messages to a function and publish the returned payload (if provided).
// <b>You need to create the function by yourself</b>.
//
func (s *API) CreateFunctionsRoute(req *CreateFunctionsRouteRequest, opts ...scw.RequestOption) (*FunctionsRoute, error) {
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes/functions",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FunctionsRoute

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetFunctionsRouteRequest struct {
	Region scw.Region `json:"-"`
	// RouteID: route ID
	RouteID string `json:"-"`
}

// GetFunctionsRoute: get a Functions route
func (s *API) GetFunctionsRoute(req *GetFunctionsRouteRequest, opts ...scw.RequestOption) (*FunctionsRoute, error) {
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes/functions/" + fmt.Sprint(req.RouteID) + "",
		Headers: http.Header{},
	}

	var resp FunctionsRoute

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteFunctionsRouteRequest struct {
	Region scw.Region `json:"-"`
	// RouteID: route ID
	RouteID string `json:"-"`
}

// DeleteFunctionsRoute: delete a Functions route
func (s *API) DeleteFunctionsRoute(req *DeleteFunctionsRouteRequest, opts ...scw.RequestOption) error {
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes/functions/" + fmt.Sprint(req.RouteID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type CreateDatabaseRouteRequest struct {
	Region scw.Region `json:"-"`
	// Name: name of the route
	Name string `json:"name"`
	// HubID: ID of the route's hub
	HubID string `json:"hub_id"`
	// Topic: topic the route subscribes to (wildcards allowed). It must be a valid MQTT topic and up to 65535 characters
	Topic string `json:"topic"`
	// Query: SQL query to be executed ($TOPIC and $PAYLOAD variables are available, see documentation)
	Query string `json:"query"`
	// Database: database settings
	Database *DatabaseSettings `json:"database"`
}

// CreateDatabaseRoute: create a Database route
//
// Create a route that will record subscribed MQTT messages into your database.
// <b>You need to manage the database by yourself</b>.
//
func (s *API) CreateDatabaseRoute(req *CreateDatabaseRouteRequest, opts ...scw.RequestOption) (*DatabaseRoute, error) {
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes/database",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseRoute

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetDatabaseRouteRequest struct {
	Region scw.Region `json:"-"`
	// RouteID: route ID
	RouteID string `json:"-"`
}

// GetDatabaseRoute: get a Database route
func (s *API) GetDatabaseRoute(req *GetDatabaseRouteRequest, opts ...scw.RequestOption) (*DatabaseRoute, error) {
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes/database/" + fmt.Sprint(req.RouteID) + "",
		Headers: http.Header{},
	}

	var resp DatabaseRoute

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteDatabaseRouteRequest struct {
	Region scw.Region `json:"-"`
	// RouteID: route ID
	RouteID string `json:"-"`
}

// DeleteDatabaseRoute: delete a Database route
func (s *API) DeleteDatabaseRoute(req *DeleteDatabaseRouteRequest, opts ...scw.RequestOption) error {
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes/database/" + fmt.Sprint(req.RouteID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type CreateRestRouteRequest struct {
	Region scw.Region `json:"-"`
	// Name: name of the route
	Name string `json:"name"`
	// HubID: ID of the route's hub
	HubID string `json:"hub_id"`
	// Topic: topic the route subscribes to. It must be a valid MQTT topic and up to 65535 characters
	Topic string `json:"topic"`
	// Verb: HTTP Verb used to call REST URI
	//
	// Default value: get
	Verb RestRouteHTTPVerb `json:"verb"`
	// URI: URI of the REST endpoint
	URI string `json:"uri"`
	// Headers: HTTP call extra headers
	Headers map[string]string `json:"headers"`
}

// CreateRestRoute: create a Rest route
//
// Create a route that will call a REST API on received subscribed MQTT messages.
//
func (s *API) CreateRestRoute(req *CreateRestRouteRequest, opts ...scw.RequestOption) (*RestRoute, error) {
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes/rest",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RestRoute

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetRestRouteRequest struct {
	Region scw.Region `json:"-"`
	// RouteID: route ID
	RouteID string `json:"-"`
}

// GetRestRoute: get a Rest route
func (s *API) GetRestRoute(req *GetRestRouteRequest, opts ...scw.RequestOption) (*RestRoute, error) {
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes/rest/" + fmt.Sprint(req.RouteID) + "",
		Headers: http.Header{},
	}

	var resp RestRoute

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteRestRouteRequest struct {
	Region scw.Region `json:"-"`
	// RouteID: route ID
	RouteID string `json:"-"`
}

// DeleteRestRoute: delete a Rest route
func (s *API) DeleteRestRoute(req *DeleteRestRouteRequest, opts ...scw.RequestOption) error {
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/routes/rest/" + fmt.Sprint(req.RouteID) + "",
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
	// OrganizationID: filter on the organization
	OrganizationID *string `json:"-"`
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "hub_id", req.HubID)
	parameter.AddToQuery(query, "topic_prefix", req.TopicPrefix)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/networks",
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
	// Deprecated: OrganizationID: deprecated: Organization owning the resource, do not use
	//
	// Will always be assigned to the organization owning the IoT hub.
	OrganizationID string `json:"organization_id"`
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

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/networks",
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/networks/" + fmt.Sprint(req.NetworkID) + "",
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
		Path:    "/iot/v1beta1/regions/" + fmt.Sprint(req.Region) + "/networks/" + fmt.Sprint(req.NetworkID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package lb_private provides methods and message types of the lb_private v1 API.
package lb_private

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

type CheckIPResponseType string

const (
	CheckIPResponseTypeUnknown = CheckIPResponseType("unknown")
	CheckIPResponseTypeOurs    = CheckIPResponseType("ours")
	CheckIPResponseTypeNotOurs = CheckIPResponseType("not_ours")
)

func (enum CheckIPResponseType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum CheckIPResponseType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CheckIPResponseType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CheckIPResponseType(CheckIPResponseType(tmp).String())
	return nil
}

type ConsoleAPIListLBPrivateNetworksRequestOrderBy string

const (
	ConsoleAPIListLBPrivateNetworksRequestOrderByCreatedAtAsc  = ConsoleAPIListLBPrivateNetworksRequestOrderBy("created_at_asc")
	ConsoleAPIListLBPrivateNetworksRequestOrderByCreatedAtDesc = ConsoleAPIListLBPrivateNetworksRequestOrderBy("created_at_desc")
)

func (enum ConsoleAPIListLBPrivateNetworksRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ConsoleAPIListLBPrivateNetworksRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ConsoleAPIListLBPrivateNetworksRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ConsoleAPIListLBPrivateNetworksRequestOrderBy(ConsoleAPIListLBPrivateNetworksRequestOrderBy(tmp).String())
	return nil
}

type ForwardPortAlgorithm string

const (
	ForwardPortAlgorithmRoundrobin = ForwardPortAlgorithm("roundrobin")
	ForwardPortAlgorithmLeastconn  = ForwardPortAlgorithm("leastconn")
)

func (enum ForwardPortAlgorithm) String() string {
	if enum == "" {
		// return default value if empty
		return "roundrobin"
	}
	return string(enum)
}

func (enum ForwardPortAlgorithm) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ForwardPortAlgorithm) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ForwardPortAlgorithm(ForwardPortAlgorithm(tmp).String())
	return nil
}

type InstanceStatus string

const (
	InstanceStatusUnknown = InstanceStatus("unknown")
	InstanceStatusReady   = InstanceStatus("ready")
	InstanceStatusPending = InstanceStatus("pending")
	InstanceStatusStopped = InstanceStatus("stopped")
	InstanceStatusError   = InstanceStatus("error")
	InstanceStatusLocked  = InstanceStatus("locked")
)

func (enum InstanceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum InstanceStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InstanceStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InstanceStatus(InstanceStatus(tmp).String())
	return nil
}

type LBStatus string

const (
	LBStatusUnknown = LBStatus("unknown")
	LBStatusReady   = LBStatus("ready")
	LBStatusPending = LBStatus("pending")
	LBStatusStopped = LBStatus("stopped")
	LBStatusError   = LBStatus("error")
	LBStatusLocked  = LBStatus("locked")
)

func (enum LBStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum LBStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *LBStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = LBStatus(LBStatus(tmp).String())
	return nil
}

type ListLBRoutesRequestOrderBy string

const (
	ListLBRoutesRequestOrderByCreatedAtAsc  = ListLBRoutesRequestOrderBy("created_at_asc")
	ListLBRoutesRequestOrderByCreatedAtDesc = ListLBRoutesRequestOrderBy("created_at_desc")
)

func (enum ListLBRoutesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListLBRoutesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListLBRoutesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListLBRoutesRequestOrderBy(ListLBRoutesRequestOrderBy(tmp).String())
	return nil
}

type OnMarkedDownAction string

const (
	OnMarkedDownActionOnMarkedDownActionNone = OnMarkedDownAction("on_marked_down_action_none")
	OnMarkedDownActionShutdownSessions       = OnMarkedDownAction("shutdown_sessions")
)

func (enum OnMarkedDownAction) String() string {
	if enum == "" {
		// return default value if empty
		return "on_marked_down_action_none"
	}
	return string(enum)
}

func (enum OnMarkedDownAction) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OnMarkedDownAction) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OnMarkedDownAction(OnMarkedDownAction(tmp).String())
	return nil
}

type PrivateNetworkStatus string

const (
	PrivateNetworkStatusUnknown = PrivateNetworkStatus("unknown")
	PrivateNetworkStatusReady   = PrivateNetworkStatus("ready")
	PrivateNetworkStatusPending = PrivateNetworkStatus("pending")
	PrivateNetworkStatusError   = PrivateNetworkStatus("error")
)

func (enum PrivateNetworkStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum PrivateNetworkStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PrivateNetworkStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PrivateNetworkStatus(PrivateNetworkStatus(tmp).String())
	return nil
}

type Protocol string

const (
	ProtocolTCP  = Protocol("tcp")
	ProtocolHTTP = Protocol("http")
)

func (enum Protocol) String() string {
	if enum == "" {
		// return default value if empty
		return "tcp"
	}
	return string(enum)
}

func (enum Protocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *Protocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = Protocol(Protocol(tmp).String())
	return nil
}

type ProxyProtocol string

const (
	ProxyProtocolProxyProtocolUnknown = ProxyProtocol("proxy_protocol_unknown")
	ProxyProtocolProxyProtocolNone    = ProxyProtocol("proxy_protocol_none")
	ProxyProtocolProxyProtocolV1      = ProxyProtocol("proxy_protocol_v1")
	ProxyProtocolProxyProtocolV2      = ProxyProtocol("proxy_protocol_v2")
	ProxyProtocolProxyProtocolV2Ssl   = ProxyProtocol("proxy_protocol_v2_ssl")
	ProxyProtocolProxyProtocolV2SslCn = ProxyProtocol("proxy_protocol_v2_ssl_cn")
)

func (enum ProxyProtocol) String() string {
	if enum == "" {
		// return default value if empty
		return "proxy_protocol_unknown"
	}
	return string(enum)
}

func (enum ProxyProtocol) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ProxyProtocol) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ProxyProtocol(ProxyProtocol(tmp).String())
	return nil
}

type SSLCompatibilityLevel string

const (
	SSLCompatibilityLevelSslCompatibilityLevelUnknown      = SSLCompatibilityLevel("ssl_compatibility_level_unknown")
	SSLCompatibilityLevelSslCompatibilityLevelIntermediate = SSLCompatibilityLevel("ssl_compatibility_level_intermediate")
	SSLCompatibilityLevelSslCompatibilityLevelModern       = SSLCompatibilityLevel("ssl_compatibility_level_modern")
	SSLCompatibilityLevelSslCompatibilityLevelOld          = SSLCompatibilityLevel("ssl_compatibility_level_old")
)

func (enum SSLCompatibilityLevel) String() string {
	if enum == "" {
		// return default value if empty
		return "ssl_compatibility_level_unknown"
	}
	return string(enum)
}

func (enum SSLCompatibilityLevel) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SSLCompatibilityLevel) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SSLCompatibilityLevel(SSLCompatibilityLevel(tmp).String())
	return nil
}

type StickySessionsType string

const (
	StickySessionsTypeNone   = StickySessionsType("none")
	StickySessionsTypeCookie = StickySessionsType("cookie")
	StickySessionsTypeTable  = StickySessionsType("table")
)

func (enum StickySessionsType) String() string {
	if enum == "" {
		// return default value if empty
		return "none"
	}
	return string(enum)
}

func (enum StickySessionsType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *StickySessionsType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = StickySessionsType(StickySessionsType(tmp).String())
	return nil
}

// IP: ip.
type IP struct {
	ID string `json:"id"`

	IPAddress string `json:"ip_address"`

	OrganizationID string `json:"organization_id"`

	LBID *string `json:"lb_id"`

	Reverse string `json:"reverse"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
}

// Instance: instance.
type Instance struct {
	ID string `json:"id"`

	// Status: default value: unknown
	Status InstanceStatus `json:"status"`

	IPAddress string `json:"ip_address"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`
}

// HealthCheckHTTPConfig: health check http config.
type HealthCheckHTTPConfig struct {
	URI string `json:"uri"`

	Method string `json:"method"`

	Code *int32 `json:"code"`

	HostHeader string `json:"host_header"`
}

// HealthCheckHTTPSConfig: health check https config.
type HealthCheckHTTPSConfig struct {
	URI string `json:"uri"`

	Method string `json:"method"`

	Code *int32 `json:"code"`

	HostHeader string `json:"host_header"`

	Sni string `json:"sni"`
}

// HealthCheckLdapConfig: health check ldap config.
type HealthCheckLdapConfig struct {
}

// HealthCheckMysqlConfig: health check mysql config.
type HealthCheckMysqlConfig struct {
	User string `json:"user"`
}

// HealthCheckPgsqlConfig: health check pgsql config.
type HealthCheckPgsqlConfig struct {
	User string `json:"user"`
}

// HealthCheckRedisConfig: health check redis config.
type HealthCheckRedisConfig struct {
}

// HealthCheckTCPConfig: health check tcp config.
type HealthCheckTCPConfig struct {
}

// LB: lb.
type LB struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`

	// Status: default value: unknown
	Status LBStatus `json:"status"`

	Instances []*Instance `json:"instances"`

	OrganizationID string `json:"organization_id"`

	IP []*IP `json:"ip"`

	Tags []string `json:"tags"`

	FrontendCount int32 `json:"frontend_count"`

	BackendCount int32 `json:"backend_count"`

	Type string `json:"type"`

	// SslCompatibilityLevel: default value: ssl_compatibility_level_unknown
	SslCompatibilityLevel SSLCompatibilityLevel `json:"ssl_compatibility_level"`

	PrivateNetworkCount int32 `json:"private_network_count"`

	RouteCount int32 `json:"route_count"`

	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"region"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`
}

// PrivateNetworkDHCPConfig: private network dhcp config.
type PrivateNetworkDHCPConfig struct {
}

// PrivateNetworkStaticConfig: private network static config.
type PrivateNetworkStaticConfig struct {
	IPAddress []string `json:"ip_address"`
}

// RouteMatch: route match.
type RouteMatch struct {
	Sni *string `json:"sni"`
}

// HealthCheck: health check.
type HealthCheck struct {
	Port int32 `json:"port"`

	CheckDelay *time.Duration `json:"check_delay"`

	CheckTimeout *time.Duration `json:"check_timeout"`

	CheckMaxRetries int32 `json:"check_max_retries"`

	// Precisely one of TCPConfig, MysqlConfig, PgsqlConfig, LdapConfig, RedisConfig, HTTPConfig, HTTPSConfig must be set.
	TCPConfig *HealthCheckTCPConfig `json:"tcp_config,omitempty"`

	// Precisely one of TCPConfig, MysqlConfig, PgsqlConfig, LdapConfig, RedisConfig, HTTPConfig, HTTPSConfig must be set.
	MysqlConfig *HealthCheckMysqlConfig `json:"mysql_config,omitempty"`

	// Precisely one of TCPConfig, MysqlConfig, PgsqlConfig, LdapConfig, RedisConfig, HTTPConfig, HTTPSConfig must be set.
	PgsqlConfig *HealthCheckPgsqlConfig `json:"pgsql_config,omitempty"`

	// Precisely one of TCPConfig, MysqlConfig, PgsqlConfig, LdapConfig, RedisConfig, HTTPConfig, HTTPSConfig must be set.
	LdapConfig *HealthCheckLdapConfig `json:"ldap_config,omitempty"`

	// Precisely one of TCPConfig, MysqlConfig, PgsqlConfig, LdapConfig, RedisConfig, HTTPConfig, HTTPSConfig must be set.
	RedisConfig *HealthCheckRedisConfig `json:"redis_config,omitempty"`

	// Precisely one of TCPConfig, MysqlConfig, PgsqlConfig, LdapConfig, RedisConfig, HTTPConfig, HTTPSConfig must be set.
	HTTPConfig *HealthCheckHTTPConfig `json:"http_config,omitempty"`

	// Precisely one of TCPConfig, MysqlConfig, PgsqlConfig, LdapConfig, RedisConfig, HTTPConfig, HTTPSConfig must be set.
	HTTPSConfig *HealthCheckHTTPSConfig `json:"https_config,omitempty"`
}

func (m *HealthCheck) UnmarshalJSON(b []byte) error {
	type tmpType HealthCheck
	tmp := struct {
		tmpType
		TmpCheckDelay   *marshaler.Duration `json:"check_delay"`
		TmpCheckTimeout *marshaler.Duration `json:"check_timeout"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = HealthCheck(tmp.tmpType)
	m.CheckDelay = tmp.TmpCheckDelay.Standard()
	m.CheckTimeout = tmp.TmpCheckTimeout.Standard()
	return nil
}

func (m HealthCheck) MarshalJSON() ([]byte, error) {
	type tmpType HealthCheck
	tmp := struct {
		tmpType
		TmpCheckDelay   *marshaler.Duration `json:"check_delay"`
		TmpCheckTimeout *marshaler.Duration `json:"check_timeout"`
	}{
		tmpType:         tmpType(m),
		TmpCheckDelay:   marshaler.NewDuration(m.CheckDelay),
		TmpCheckTimeout: marshaler.NewDuration(m.CheckTimeout),
	}
	return json.Marshal(tmp)
}

// PrivateNetwork: private network.
type PrivateNetwork struct {
	LB *LB `json:"lb"`

	PrivateNetworkID string `json:"private_network_id"`

	PrivateNetworkName string `json:"private_network_name"`

	// Status: default value: unknown
	Status PrivateNetworkStatus `json:"status"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	// Precisely one of StaticConfig, DHCPConfig must be set.
	StaticConfig *PrivateNetworkStaticConfig `json:"static_config,omitempty"`

	// Precisely one of StaticConfig, DHCPConfig must be set.
	DHCPConfig *PrivateNetworkDHCPConfig `json:"dhcp_config,omitempty"`
}

// Route: route.
type Route struct {
	ID string `json:"id"`

	FrontendID string `json:"frontend_id"`

	Match *RouteMatch `json:"match"`

	BackendID string `json:"backend_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	FrontendName string `json:"frontend_name"`

	BackendName string `json:"backend_name"`
}

// CheckIPResponse: check ip response.
type CheckIPResponse struct {
	IsOurs bool `json:"is_ours"`

	// Type: default value: unknown
	Type CheckIPResponseType `json:"type"`
}

// ConsoleAPICheckIPRequest: console api check ip request.
type ConsoleAPICheckIPRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	IP string `json:"ip"`
}

// ConsoleAPICreateFullFrontendRequest: console api create full frontend request.
type ConsoleAPICreateFullFrontendRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	LBID string `json:"lb_id"`

	BackendName string `json:"backend_name"`

	// ForwardProtocol: default value: tcp
	ForwardProtocol Protocol `json:"forward_protocol"`

	ForwardPort int32 `json:"forward_port"`

	// ForwardPortAlgorithm: default value: roundrobin
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm"`

	// StickySessions: default value: none
	StickySessions StickySessionsType `json:"sticky_sessions"`

	StickySessionsCookieName string `json:"sticky_sessions_cookie_name"`

	HealthCheck *HealthCheck `json:"health_check,omitempty"`

	ServerIP []string `json:"server_ip"`

	SendProxyV2 bool `json:"send_proxy_v2"`

	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`

	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`

	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`

	// OnMarkedDownAction: default value: on_marked_down_action_none
	OnMarkedDownAction OnMarkedDownAction `json:"on_marked_down_action"`

	FrontendName string `json:"frontend_name"`

	InboundPort int32 `json:"inbound_port"`

	TimeoutClient *time.Duration `json:"timeout_client,omitempty"`

	// Deprecated
	CertificateID *string `json:"certificate_id,omitempty"`

	// ProxyProtocol: default value: proxy_protocol_unknown
	ProxyProtocol ProxyProtocol `json:"proxy_protocol"`

	CertificateIDs *[]string `json:"certificate_ids,omitempty"`

	FailoverHost *string `json:"failover_host,omitempty"`

	SslBridging *bool `json:"ssl_bridging,omitempty"`

	IgnoreSslServerVerify *bool `json:"ignore_ssl_server_verify,omitempty"`

	EnableHTTP3 bool `json:"enable_http3"`
}

func (m *ConsoleAPICreateFullFrontendRequest) UnmarshalJSON(b []byte) error {
	type tmpType ConsoleAPICreateFullFrontendRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
		TmpTimeoutClient  *marshaler.Duration `json:"timeout_client,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ConsoleAPICreateFullFrontendRequest(tmp.tmpType)
	m.TimeoutServer = tmp.TmpTimeoutServer.Standard()
	m.TimeoutConnect = tmp.TmpTimeoutConnect.Standard()
	m.TimeoutTunnel = tmp.TmpTimeoutTunnel.Standard()
	m.TimeoutClient = tmp.TmpTimeoutClient.Standard()
	return nil
}

func (m ConsoleAPICreateFullFrontendRequest) MarshalJSON() ([]byte, error) {
	type tmpType ConsoleAPICreateFullFrontendRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
		TmpTimeoutClient  *marshaler.Duration `json:"timeout_client,omitempty"`
	}{
		tmpType:           tmpType(m),
		TmpTimeoutServer:  marshaler.NewDuration(m.TimeoutServer),
		TmpTimeoutConnect: marshaler.NewDuration(m.TimeoutConnect),
		TmpTimeoutTunnel:  marshaler.NewDuration(m.TimeoutTunnel),
		TmpTimeoutClient:  marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

// ConsoleAPICreateFullLBRequest: console api create full lb request.
type ConsoleAPICreateFullLBRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Deprecated
	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	LBName string `json:"lb_name"`

	Description string `json:"description"`

	// Deprecated
	IPID *string `json:"ip_id,omitempty"`

	Tags []string `json:"tags"`

	BackendName string `json:"backend_name"`

	// ForwardProtocol: default value: tcp
	ForwardProtocol Protocol `json:"forward_protocol"`

	ForwardPort int32 `json:"forward_port"`

	// ForwardPortAlgorithm: default value: roundrobin
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm"`

	// StickySessions: default value: none
	StickySessions StickySessionsType `json:"sticky_sessions"`

	StickySessionsCookieName string `json:"sticky_sessions_cookie_name"`

	HealthCheck *HealthCheck `json:"health_check,omitempty"`

	ServerIP []string `json:"server_ip"`

	SendProxyV2 bool `json:"send_proxy_v2"`

	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`

	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`

	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`

	// OnMarkedDownAction: default value: on_marked_down_action_none
	OnMarkedDownAction OnMarkedDownAction `json:"on_marked_down_action"`

	FrontendName string `json:"frontend_name"`

	InboundPort int32 `json:"inbound_port"`

	TimeoutClient *time.Duration `json:"timeout_client,omitempty"`

	Type string `json:"type"`

	IPReverse *string `json:"ip_reverse,omitempty"`

	// ProxyProtocol: default value: proxy_protocol_unknown
	ProxyProtocol ProxyProtocol `json:"proxy_protocol"`

	// SslCompatibilityLevel: default value: ssl_compatibility_level_unknown
	SslCompatibilityLevel SSLCompatibilityLevel `json:"ssl_compatibility_level"`

	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	FailoverHost *string `json:"failover_host,omitempty"`

	SslBridging *bool `json:"ssl_bridging,omitempty"`

	IgnoreSslServerVerify *bool `json:"ignore_ssl_server_verify,omitempty"`

	EnableHTTP3 bool `json:"enable_http3"`

	AssignFlexibleIP *bool `json:"assign_flexible_ip,omitempty"`

	IPIDs []string `json:"ip_ids"`
}

func (m *ConsoleAPICreateFullLBRequest) UnmarshalJSON(b []byte) error {
	type tmpType ConsoleAPICreateFullLBRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
		TmpTimeoutClient  *marshaler.Duration `json:"timeout_client,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ConsoleAPICreateFullLBRequest(tmp.tmpType)
	m.TimeoutServer = tmp.TmpTimeoutServer.Standard()
	m.TimeoutConnect = tmp.TmpTimeoutConnect.Standard()
	m.TimeoutTunnel = tmp.TmpTimeoutTunnel.Standard()
	m.TimeoutClient = tmp.TmpTimeoutClient.Standard()
	return nil
}

func (m ConsoleAPICreateFullLBRequest) MarshalJSON() ([]byte, error) {
	type tmpType ConsoleAPICreateFullLBRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
		TmpTimeoutClient  *marshaler.Duration `json:"timeout_client,omitempty"`
	}{
		tmpType:           tmpType(m),
		TmpTimeoutServer:  marshaler.NewDuration(m.TimeoutServer),
		TmpTimeoutConnect: marshaler.NewDuration(m.TimeoutConnect),
		TmpTimeoutTunnel:  marshaler.NewDuration(m.TimeoutTunnel),
		TmpTimeoutClient:  marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

// ConsoleAPIGetLBMetricsRequest: console api get lb metrics request.
type ConsoleAPIGetLBMetricsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	LBID string `json:"-"`

	StartDate *time.Time `json:"-"`

	EndDate *time.Time `json:"-"`

	MetricName *string `json:"-"`
}

// ConsoleAPIGetServiceInfoRequest: console api get service info request.
type ConsoleAPIGetServiceInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`
}

// ConsoleAPIListLBPrivateNetworksRequest: console api list lb private networks request.
type ConsoleAPIListLBPrivateNetworksRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	LBID string `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ConsoleAPIListLBPrivateNetworksRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ConsoleAPIListLBPrivateNetworksResponse: console api list lb private networks response.
type ConsoleAPIListLBPrivateNetworksResponse struct {
	PrivateNetwork []*PrivateNetwork `json:"private_network"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ConsoleAPIListLBPrivateNetworksResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ConsoleAPIListLBPrivateNetworksResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ConsoleAPIListLBPrivateNetworksResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.PrivateNetwork = append(r.PrivateNetwork, results.PrivateNetwork...)
	r.TotalCount += uint32(len(results.PrivateNetwork))
	return uint32(len(results.PrivateNetwork)), nil
}

// Frontend: frontend.
type Frontend struct {
	ID string `json:"id"`

	Name string `json:"name"`

	InboundPort int32 `json:"inbound_port"`

	LB *LB `json:"lb"`

	TimeoutClient *time.Duration `json:"timeout_client"`

	EnableHTTP3 bool `json:"enable_http3"`
}

func (m *Frontend) UnmarshalJSON(b []byte) error {
	type tmpType Frontend
	tmp := struct {
		tmpType
		TmpTimeoutClient *marshaler.Duration `json:"timeout_client"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = Frontend(tmp.tmpType)
	m.TimeoutClient = tmp.TmpTimeoutClient.Standard()
	return nil
}

func (m Frontend) MarshalJSON() ([]byte, error) {
	type tmpType Frontend
	tmp := struct {
		tmpType
		TmpTimeoutClient *marshaler.Duration `json:"timeout_client"`
	}{
		tmpType:          tmpType(m),
		TmpTimeoutClient: marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

// LBMetrics: lb metrics.
type LBMetrics struct {
	Timeseries []*scw.TimeSeries `json:"timeseries"`
}

// ListRoutesResponse: list routes response.
type ListRoutesResponse struct {
	Routes []*Route `json:"routes"`

	TotalCount uint32 `json:"total_count"`
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

// ZonedConsoleAPICheckIPRequest: zoned console api check ip request.
type ZonedConsoleAPICheckIPRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	IP string `json:"ip"`
}

// ZonedConsoleAPICreateFullFrontendRequest: zoned console api create full frontend request.
type ZonedConsoleAPICreateFullFrontendRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	LBID string `json:"lb_id"`

	BackendName string `json:"backend_name"`

	// ForwardProtocol: default value: tcp
	ForwardProtocol Protocol `json:"forward_protocol"`

	ForwardPort int32 `json:"forward_port"`

	// ForwardPortAlgorithm: default value: roundrobin
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm"`

	// StickySessions: default value: none
	StickySessions StickySessionsType `json:"sticky_sessions"`

	StickySessionsCookieName string `json:"sticky_sessions_cookie_name"`

	HealthCheck *HealthCheck `json:"health_check,omitempty"`

	ServerIP []string `json:"server_ip"`

	SendProxyV2 bool `json:"send_proxy_v2"`

	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`

	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`

	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`

	// OnMarkedDownAction: default value: on_marked_down_action_none
	OnMarkedDownAction OnMarkedDownAction `json:"on_marked_down_action"`

	FrontendName string `json:"frontend_name"`

	InboundPort int32 `json:"inbound_port"`

	TimeoutClient *time.Duration `json:"timeout_client,omitempty"`

	// Deprecated
	CertificateID *string `json:"certificate_id,omitempty"`

	// ProxyProtocol: default value: proxy_protocol_unknown
	ProxyProtocol ProxyProtocol `json:"proxy_protocol"`

	CertificateIDs *[]string `json:"certificate_ids,omitempty"`

	FailoverHost *string `json:"failover_host,omitempty"`

	SslBridging *bool `json:"ssl_bridging,omitempty"`

	IgnoreSslServerVerify *bool `json:"ignore_ssl_server_verify,omitempty"`

	EnableHTTP3 bool `json:"enable_http3"`
}

func (m *ZonedConsoleAPICreateFullFrontendRequest) UnmarshalJSON(b []byte) error {
	type tmpType ZonedConsoleAPICreateFullFrontendRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
		TmpTimeoutClient  *marshaler.Duration `json:"timeout_client,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ZonedConsoleAPICreateFullFrontendRequest(tmp.tmpType)
	m.TimeoutServer = tmp.TmpTimeoutServer.Standard()
	m.TimeoutConnect = tmp.TmpTimeoutConnect.Standard()
	m.TimeoutTunnel = tmp.TmpTimeoutTunnel.Standard()
	m.TimeoutClient = tmp.TmpTimeoutClient.Standard()
	return nil
}

func (m ZonedConsoleAPICreateFullFrontendRequest) MarshalJSON() ([]byte, error) {
	type tmpType ZonedConsoleAPICreateFullFrontendRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
		TmpTimeoutClient  *marshaler.Duration `json:"timeout_client,omitempty"`
	}{
		tmpType:           tmpType(m),
		TmpTimeoutServer:  marshaler.NewDuration(m.TimeoutServer),
		TmpTimeoutConnect: marshaler.NewDuration(m.TimeoutConnect),
		TmpTimeoutTunnel:  marshaler.NewDuration(m.TimeoutTunnel),
		TmpTimeoutClient:  marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

// ZonedConsoleAPICreateFullLBRequest: zoned console api create full lb request.
type ZonedConsoleAPICreateFullLBRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	// Deprecated
	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	LBName string `json:"lb_name"`

	Description string `json:"description"`

	// Deprecated
	IPID *string `json:"ip_id,omitempty"`

	Tags []string `json:"tags"`

	BackendName string `json:"backend_name"`

	// ForwardProtocol: default value: tcp
	ForwardProtocol Protocol `json:"forward_protocol"`

	ForwardPort int32 `json:"forward_port"`

	// ForwardPortAlgorithm: default value: roundrobin
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm"`

	// StickySessions: default value: none
	StickySessions StickySessionsType `json:"sticky_sessions"`

	StickySessionsCookieName string `json:"sticky_sessions_cookie_name"`

	HealthCheck *HealthCheck `json:"health_check,omitempty"`

	ServerIP []string `json:"server_ip"`

	SendProxyV2 bool `json:"send_proxy_v2"`

	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`

	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`

	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`

	// OnMarkedDownAction: default value: on_marked_down_action_none
	OnMarkedDownAction OnMarkedDownAction `json:"on_marked_down_action"`

	FrontendName string `json:"frontend_name"`

	InboundPort int32 `json:"inbound_port"`

	TimeoutClient *time.Duration `json:"timeout_client,omitempty"`

	Type string `json:"type"`

	IPReverse *string `json:"ip_reverse,omitempty"`

	// ProxyProtocol: default value: proxy_protocol_unknown
	ProxyProtocol ProxyProtocol `json:"proxy_protocol"`

	// SslCompatibilityLevel: default value: ssl_compatibility_level_unknown
	SslCompatibilityLevel SSLCompatibilityLevel `json:"ssl_compatibility_level"`

	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`

	FailoverHost *string `json:"failover_host,omitempty"`

	SslBridging *bool `json:"ssl_bridging,omitempty"`

	IgnoreSslServerVerify *bool `json:"ignore_ssl_server_verify,omitempty"`

	EnableHTTP3 bool `json:"enable_http3"`

	AssignFlexibleIP *bool `json:"assign_flexible_ip,omitempty"`

	IPIDs []string `json:"ip_ids"`
}

func (m *ZonedConsoleAPICreateFullLBRequest) UnmarshalJSON(b []byte) error {
	type tmpType ZonedConsoleAPICreateFullLBRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
		TmpTimeoutClient  *marshaler.Duration `json:"timeout_client,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = ZonedConsoleAPICreateFullLBRequest(tmp.tmpType)
	m.TimeoutServer = tmp.TmpTimeoutServer.Standard()
	m.TimeoutConnect = tmp.TmpTimeoutConnect.Standard()
	m.TimeoutTunnel = tmp.TmpTimeoutTunnel.Standard()
	m.TimeoutClient = tmp.TmpTimeoutClient.Standard()
	return nil
}

func (m ZonedConsoleAPICreateFullLBRequest) MarshalJSON() ([]byte, error) {
	type tmpType ZonedConsoleAPICreateFullLBRequest
	tmp := struct {
		tmpType
		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
		TmpTimeoutClient  *marshaler.Duration `json:"timeout_client,omitempty"`
	}{
		tmpType:           tmpType(m),
		TmpTimeoutServer:  marshaler.NewDuration(m.TimeoutServer),
		TmpTimeoutConnect: marshaler.NewDuration(m.TimeoutConnect),
		TmpTimeoutTunnel:  marshaler.NewDuration(m.TimeoutTunnel),
		TmpTimeoutClient:  marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

// ZonedConsoleAPIGetLBMetricsRequest: zoned console api get lb metrics request.
type ZonedConsoleAPIGetLBMetricsRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	LBID string `json:"-"`

	StartDate *time.Time `json:"-"`

	EndDate *time.Time `json:"-"`

	MetricName *string `json:"-"`
}

// ZonedConsoleAPIGetServiceInfoRequest: zoned console api get service info request.
type ZonedConsoleAPIGetServiceInfoRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`
}

// ZonedConsoleAPIListLBPrivateNetworksRequest: zoned console api list lb private networks request.
type ZonedConsoleAPIListLBPrivateNetworksRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	LBID string `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ConsoleAPIListLBPrivateNetworksRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

// ZonedConsoleAPIListLBRoutesRequest: zoned console api list lb routes request.
type ZonedConsoleAPIListLBRoutesRequest struct {
	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"-"`

	LBID string `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListLBRoutesRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`
}

type ConsoleAPI struct {
	client *scw.Client
}

// NewConsoleAPI returns a ConsoleAPI object from a Scaleway client.
func NewConsoleAPI(client *scw.Client) *ConsoleAPI {
	return &ConsoleAPI{
		client: client,
	}
}
func (s *ConsoleAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms, scw.RegionPlWaw}
}

// GetServiceInfo:
func (s *ConsoleAPI) GetServiceInfo(req *ConsoleAPIGetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
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
		Path:   "/lb-private/v1/regions/" + fmt.Sprint(req.Region) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFullLB:
func (s *ConsoleAPI) CreateFullLB(req *ConsoleAPICreateFullLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb-private/v1/regions/" + fmt.Sprint(req.Region) + "/fulllbs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFullFrontend:
func (s *ConsoleAPI) CreateFullFrontend(req *ConsoleAPICreateFullFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb-private/v1/regions/" + fmt.Sprint(req.Region) + "/fullfrontend",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Frontend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckIP:
func (s *ConsoleAPI) CheckIP(req *ConsoleAPICheckIPRequest, opts ...scw.RequestOption) (*CheckIPResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb-private/v1/regions/" + fmt.Sprint(req.Region) + "/checkip",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CheckIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLBMetrics:
func (s *ConsoleAPI) GetLBMetrics(req *ConsoleAPIGetLBMetricsRequest, opts ...scw.RequestOption) (*LBMetrics, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)
	parameter.AddToQuery(query, "metric_name", req.MetricName)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb-private/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/metrics",
		Query:  query,
	}

	var resp LBMetrics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListLBPrivateNetworks:
func (s *ConsoleAPI) ListLBPrivateNetworks(req *ConsoleAPIListLBPrivateNetworksRequest, opts ...scw.RequestOption) (*ConsoleAPIListLBPrivateNetworksResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb-private/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LBID) + "/private-networks",
		Query:  query,
	}

	var resp ConsoleAPIListLBPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ZonedConsoleAPI struct {
	client *scw.Client
}

// NewZonedConsoleAPI returns a ZonedConsoleAPI object from a Scaleway client.
func NewZonedConsoleAPI(client *scw.Client) *ZonedConsoleAPI {
	return &ZonedConsoleAPI{
		client: client,
	}
}
func (s *ZonedConsoleAPI) Zones() []scw.Zone {
	return []scw.Zone{scw.ZoneFrPar1, scw.ZoneFrPar2, scw.ZoneNlAms1, scw.ZonePlWaw1}
}

// GetServiceInfo:
func (s *ZonedConsoleAPI) GetServiceInfo(req *ZonedConsoleAPIGetServiceInfoRequest, opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb-private/v1/zones/" + fmt.Sprint(req.Zone) + "",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFullLB:
func (s *ZonedConsoleAPI) CreateFullLB(req *ZonedConsoleAPICreateFullLBRequest, opts ...scw.RequestOption) (*LB, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb-private/v1/zones/" + fmt.Sprint(req.Zone) + "/fulllbs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LB

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFullFrontend:
func (s *ZonedConsoleAPI) CreateFullFrontend(req *ZonedConsoleAPICreateFullFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb-private/v1/zones/" + fmt.Sprint(req.Zone) + "/fullfrontend",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Frontend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckIP:
func (s *ZonedConsoleAPI) CheckIP(req *ZonedConsoleAPICheckIPRequest, opts ...scw.RequestOption) (*CheckIPResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/lb-private/v1/zones/" + fmt.Sprint(req.Zone) + "/checkip",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CheckIPResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLBMetrics:
func (s *ZonedConsoleAPI) GetLBMetrics(req *ZonedConsoleAPIGetLBMetricsRequest, opts ...scw.RequestOption) (*LBMetrics, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)
	parameter.AddToQuery(query, "metric_name", req.MetricName)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb-private/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/metrics",
		Query:  query,
	}

	var resp LBMetrics

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListLBPrivateNetworks: List all pn of a loadbalancer.
func (s *ZonedConsoleAPI) ListLBPrivateNetworks(req *ZonedConsoleAPIListLBPrivateNetworksRequest, opts ...scw.RequestOption) (*ConsoleAPIListLBPrivateNetworksResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb-private/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/private-networks",
		Query:  query,
	}

	var resp ConsoleAPIListLBPrivateNetworksResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListLBRoutes:
func (s *ZonedConsoleAPI) ListLBRoutes(req *ZonedConsoleAPIListLBRoutesRequest, opts ...scw.RequestOption) (*ListRoutesResponse, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Zone) == "" {
		return nil, errors.New("field Zone cannot be empty in request")
	}

	if fmt.Sprint(req.LBID) == "" {
		return nil, errors.New("field LBID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/lb-private/v1/zones/" + fmt.Sprint(req.Zone) + "/lbs/" + fmt.Sprint(req.LBID) + "/routes",
		Query:  query,
	}

	var resp ListRoutesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

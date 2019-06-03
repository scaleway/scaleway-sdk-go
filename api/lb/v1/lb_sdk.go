// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

package lb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"time"

	"github.com/scaleway/scaleway-sdk-go/internal/errors"
	"github.com/scaleway/scaleway-sdk-go/internal/marshaler"
	"github.com/scaleway/scaleway-sdk-go/internal/parameter"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/utils"
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

	_ scw.ScalewayRequest
	_ marshaler.Duration
	_ utils.File
	_ = parameter.AddToQuery
)

// API this API allows you to manage your Load Balancer service
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

type ACLActionType string

const (
	// ACLActionTypeAllow is [insert doc].
	ACLActionTypeAllow = ACLActionType("allow")
	// ACLActionTypeDeny is [insert doc].
	ACLActionTypeDeny = ACLActionType("deny")
)

func (enum ACLActionType) String() string {
	if enum == "" {
		// return default value if empty
		return "allow"
	}
	return string(enum)
}

type ACLHTTPFilter string

const (
	// ACLHTTPFilterACLHTTPFilterNone is [insert doc].
	ACLHTTPFilterACLHTTPFilterNone = ACLHTTPFilter("acl_http_filter_none")
	// ACLHTTPFilterPathBegin is [insert doc].
	ACLHTTPFilterPathBegin = ACLHTTPFilter("path_begin")
	// ACLHTTPFilterPathEnd is [insert doc].
	ACLHTTPFilterPathEnd = ACLHTTPFilter("path_end")
	// ACLHTTPFilterRegex is [insert doc].
	ACLHTTPFilterRegex = ACLHTTPFilter("regex")
)

func (enum ACLHTTPFilter) String() string {
	if enum == "" {
		// return default value if empty
		return "acl_http_filter_none"
	}
	return string(enum)
}

type BackendServerStatsHealthCheckStatus string

const (
	// BackendServerStatsHealthCheckStatusUnknown is [insert doc].
	BackendServerStatsHealthCheckStatusUnknown = BackendServerStatsHealthCheckStatus("unknown")
	// BackendServerStatsHealthCheckStatusNeutral is [insert doc].
	BackendServerStatsHealthCheckStatusNeutral = BackendServerStatsHealthCheckStatus("neutral")
	// BackendServerStatsHealthCheckStatusFailed is [insert doc].
	BackendServerStatsHealthCheckStatusFailed = BackendServerStatsHealthCheckStatus("failed")
	// BackendServerStatsHealthCheckStatusPassed is [insert doc].
	BackendServerStatsHealthCheckStatusPassed = BackendServerStatsHealthCheckStatus("passed")
	// BackendServerStatsHealthCheckStatusCondpass is [insert doc].
	BackendServerStatsHealthCheckStatusCondpass = BackendServerStatsHealthCheckStatus("condpass")
)

func (enum BackendServerStatsHealthCheckStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

type BackendServerStatsServerState string

const (
	// BackendServerStatsServerStateStopped is [insert doc].
	BackendServerStatsServerStateStopped = BackendServerStatsServerState("stopped")
	// BackendServerStatsServerStateStarting is [insert doc].
	BackendServerStatsServerStateStarting = BackendServerStatsServerState("starting")
	// BackendServerStatsServerStateRunning is [insert doc].
	BackendServerStatsServerStateRunning = BackendServerStatsServerState("running")
	// BackendServerStatsServerStateStopping is [insert doc].
	BackendServerStatsServerStateStopping = BackendServerStatsServerState("stopping")
)

func (enum BackendServerStatsServerState) String() string {
	if enum == "" {
		// return default value if empty
		return "stopped"
	}
	return string(enum)
}

type ForwardPortAlgorithm string

const (
	// ForwardPortAlgorithmRoundrobin is [insert doc].
	ForwardPortAlgorithmRoundrobin = ForwardPortAlgorithm("roundrobin")
	// ForwardPortAlgorithmLeastconn is [insert doc].
	ForwardPortAlgorithmLeastconn = ForwardPortAlgorithm("leastconn")
)

func (enum ForwardPortAlgorithm) String() string {
	if enum == "" {
		// return default value if empty
		return "roundrobin"
	}
	return string(enum)
}

type InstanceStatus string

const (
	// InstanceStatusUnknown is [insert doc].
	InstanceStatusUnknown = InstanceStatus("unknown")
	// InstanceStatusReady is [insert doc].
	InstanceStatusReady = InstanceStatus("ready")
	// InstanceStatusPending is [insert doc].
	InstanceStatusPending = InstanceStatus("pending")
	// InstanceStatusStopped is [insert doc].
	InstanceStatusStopped = InstanceStatus("stopped")
	// InstanceStatusError is [insert doc].
	InstanceStatusError = InstanceStatus("error")
	// InstanceStatusLocked is [insert doc].
	InstanceStatusLocked = InstanceStatus("locked")
)

func (enum InstanceStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

type LbStatus string

const (
	// LbStatusUnknown is [insert doc].
	LbStatusUnknown = LbStatus("unknown")
	// LbStatusReady is [insert doc].
	LbStatusReady = LbStatus("ready")
	// LbStatusPending is [insert doc].
	LbStatusPending = LbStatus("pending")
	// LbStatusStopped is [insert doc].
	LbStatusStopped = LbStatus("stopped")
	// LbStatusError is [insert doc].
	LbStatusError = LbStatus("error")
	// LbStatusLocked is [insert doc].
	LbStatusLocked = LbStatus("locked")
)

func (enum LbStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

type ListACLRequestOrderBy string

const (
	// ListACLRequestOrderByCreatedAtAsc is [insert doc].
	ListACLRequestOrderByCreatedAtAsc = ListACLRequestOrderBy("created_at_asc")
	// ListACLRequestOrderByCreatedAtDesc is [insert doc].
	ListACLRequestOrderByCreatedAtDesc = ListACLRequestOrderBy("created_at_desc")
	// ListACLRequestOrderByNameAsc is [insert doc].
	ListACLRequestOrderByNameAsc = ListACLRequestOrderBy("name_asc")
	// ListACLRequestOrderByNameDesc is [insert doc].
	ListACLRequestOrderByNameDesc = ListACLRequestOrderBy("name_desc")
)

func (enum ListACLRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

type ListBackendsRequestOrderBy string

const (
	// ListBackendsRequestOrderByCreatedAtAsc is [insert doc].
	ListBackendsRequestOrderByCreatedAtAsc = ListBackendsRequestOrderBy("created_at_asc")
	// ListBackendsRequestOrderByCreatedAtDesc is [insert doc].
	ListBackendsRequestOrderByCreatedAtDesc = ListBackendsRequestOrderBy("created_at_desc")
	// ListBackendsRequestOrderByNameAsc is [insert doc].
	ListBackendsRequestOrderByNameAsc = ListBackendsRequestOrderBy("name_asc")
	// ListBackendsRequestOrderByNameDesc is [insert doc].
	ListBackendsRequestOrderByNameDesc = ListBackendsRequestOrderBy("name_desc")
)

func (enum ListBackendsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

type ListFrontendsRequestOrderBy string

const (
	// ListFrontendsRequestOrderByCreatedAtAsc is [insert doc].
	ListFrontendsRequestOrderByCreatedAtAsc = ListFrontendsRequestOrderBy("created_at_asc")
	// ListFrontendsRequestOrderByCreatedAtDesc is [insert doc].
	ListFrontendsRequestOrderByCreatedAtDesc = ListFrontendsRequestOrderBy("created_at_desc")
	// ListFrontendsRequestOrderByNameAsc is [insert doc].
	ListFrontendsRequestOrderByNameAsc = ListFrontendsRequestOrderBy("name_asc")
	// ListFrontendsRequestOrderByNameDesc is [insert doc].
	ListFrontendsRequestOrderByNameDesc = ListFrontendsRequestOrderBy("name_desc")
)

func (enum ListFrontendsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

type ListLbsRequestOrderBy string

const (
	// ListLbsRequestOrderByCreatedAtAsc is [insert doc].
	ListLbsRequestOrderByCreatedAtAsc = ListLbsRequestOrderBy("created_at_asc")
	// ListLbsRequestOrderByCreatedAtDesc is [insert doc].
	ListLbsRequestOrderByCreatedAtDesc = ListLbsRequestOrderBy("created_at_desc")
	// ListLbsRequestOrderByNameAsc is [insert doc].
	ListLbsRequestOrderByNameAsc = ListLbsRequestOrderBy("name_asc")
	// ListLbsRequestOrderByNameDesc is [insert doc].
	ListLbsRequestOrderByNameDesc = ListLbsRequestOrderBy("name_desc")
)

func (enum ListLbsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

type OnMarkedDownAction string

const (
	// OnMarkedDownActionOnMarkedDownActionNone is [insert doc].
	OnMarkedDownActionOnMarkedDownActionNone = OnMarkedDownAction("on_marked_down_action_none")
	// OnMarkedDownActionShutdownSessions is [insert doc].
	OnMarkedDownActionShutdownSessions = OnMarkedDownAction("shutdown_sessions")
)

func (enum OnMarkedDownAction) String() string {
	if enum == "" {
		// return default value if empty
		return "on_marked_down_action_none"
	}
	return string(enum)
}

type Protocol string

const (
	// ProtocolTCP is [insert doc].
	ProtocolTCP = Protocol("tcp")
	// ProtocolHTTP is [insert doc].
	ProtocolHTTP = Protocol("http")
)

func (enum Protocol) String() string {
	if enum == "" {
		// return default value if empty
		return "tcp"
	}
	return string(enum)
}

type StickySessionsType string

const (
	// StickySessionsTypeNone is [insert doc].
	StickySessionsTypeNone = StickySessionsType("none")
	// StickySessionsTypeCookie is [insert doc].
	StickySessionsTypeCookie = StickySessionsType("cookie")
	// StickySessionsTypeTable is [insert doc].
	StickySessionsTypeTable = StickySessionsType("table")
)

func (enum StickySessionsType) String() string {
	if enum == "" {
		// return default value if empty
		return "none"
	}
	return string(enum)
}

// ACL the use of Access Control Lists (ACL) provide a flexible solution to perform a action generally consist in blocking or allow a request based on ip (and URL on HTTP)
type ACL struct {
	// ID iD of your ACL ressource
	ID string `json:"id,omitempty"`
	// Name name of you ACL ressource
	Name string `json:"name,omitempty"`
	// Match see the AclMatch object description
	Match *ACLMatch `json:"match,omitempty"`
	// Action see the AclAction object description
	Action *ACLAction `json:"action,omitempty"`
	// Frontend see the Frontend object description
	Frontend *Frontend `json:"frontend,omitempty"`
	// Index order between your Acls (ascending order, 0 is first acl executed)
	Index int32 `json:"index,omitempty"`
}

// ACLAction action if your ACL filter match
type ACLAction struct {
	// Type <allow> or <deny> request
	//
	// Default value: allow
	Type ACLActionType `json:"type,omitempty"`
}

// ACLMatch settings of your ACL filter
type ACLMatch struct {
	// IPSubnet this is the source IP v4/v6 address of the client of the session to match or not. Addresses values can be specified either as plain addresses or with a netmask appended
	IPSubnet []*string `json:"ip_subnet,omitempty"`
	// HTTPFilter you can set http filter (if your backend protocole have a http forward protocol. This extracts the request's URL path, which starts at the first slash and ends before the question mark (without the host part). You can choose between <path_begin> prefix match (like /admin), <path_end> suffix match (like .php) and <regex>
	//
	// Default value: acl_http_filter_none
	HTTPFilter ACLHTTPFilter `json:"http_filter,omitempty"`

	HTTPFilterValue []*string `json:"http_filter_value,omitempty"`
	// Invert by default match filter is a IF condition. You can set invert to true to have a unless condition
	Invert bool `json:"invert,omitempty"`
}

type Backend struct {
	ID string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`
	// ForwardProtocol
	//
	// Default value: tcp
	ForwardProtocol Protocol `json:"forward_protocol,omitempty"`

	ForwardPort int32 `json:"forward_port,omitempty"`
	// ForwardPortAlgorithm
	//
	// Default value: roundrobin
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm,omitempty"`
	// StickySessions
	//
	// Default value: none
	StickySessions StickySessionsType `json:"sticky_sessions,omitempty"`

	StickySessionsCookieName string `json:"sticky_sessions_cookie_name,omitempty"`

	HealthCheck *HealthCheck `json:"health_check,omitempty"`

	Pool []string `json:"pool,omitempty"`

	Lb *Lb `json:"lb,omitempty"`

	SendProxyV2 bool `json:"send_proxy_v2,omitempty"`

	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`

	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`

	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`
	// OnMarkedDownAction
	//
	// Default value: on_marked_down_action_none
	OnMarkedDownAction OnMarkedDownAction `json:"on_marked_down_action,omitempty"`
}

func (m *Backend) UnmarshalJSON(b []byte) error {
	type tmpType Backend
	tmp := struct {
		tmpType

		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = Backend(tmp.tmpType)

	m.TimeoutServer = tmp.TmpTimeoutServer.Standard()
	m.TimeoutConnect = tmp.TmpTimeoutConnect.Standard()
	m.TimeoutTunnel = tmp.TmpTimeoutTunnel.Standard()
	return nil
}

func (m Backend) MarshalJSON() ([]byte, error) {
	type tmpType Backend
	tmp := struct {
		tmpType

		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{
		tmpType: tmpType(m),

		TmpTimeoutServer:  marshaler.NewDuration(m.TimeoutServer),
		TmpTimeoutConnect: marshaler.NewDuration(m.TimeoutConnect),
		TmpTimeoutTunnel:  marshaler.NewDuration(m.TimeoutTunnel),
	}
	return json.Marshal(tmp)
}

// BackendServerStats state and statistics of your backend server like last healthcheck status, server uptime, result state of your backend server
type BackendServerStats struct {
	// InstanceID iD of your loadbalancer cluster server
	InstanceID string `json:"instance_id,omitempty"`
	// BackendID iD of your Backend
	BackendID string `json:"backend_id,omitempty"`
	// IP iPv4 or IPv6 address of the server backend
	IP string `json:"ip,omitempty"`
	// ServerState server operational state (stopped/starting/running/stopping)
	//
	// Default value: stopped
	ServerState BackendServerStatsServerState `json:"server_state,omitempty"`
	// ServerStateChangedAt time since last operational change
	ServerStateChangedAt time.Time `json:"server_state_changed_at,omitempty"`
	// LastHealthCheckStatus last health check status (unknown/neutral/failed/passed/condpass)
	//
	// Default value: unknown
	LastHealthCheckStatus BackendServerStatsHealthCheckStatus `json:"last_health_check_status,omitempty"`
}

type Frontend struct {
	ID string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	InboundPort int32 `json:"inbound_port,omitempty"`

	Backend *Backend `json:"backend,omitempty"`

	Lb *Lb `json:"lb,omitempty"`

	TimeoutClient *time.Duration `json:"timeout_client,omitempty"`
}

func (m *Frontend) UnmarshalJSON(b []byte) error {
	type tmpType Frontend
	tmp := struct {
		tmpType

		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
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

		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{
		tmpType: tmpType(m),

		TmpTimeoutClient: marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

type HealthCheck struct {
	Port int32 `json:"port,omitempty"`

	CheckDelay *time.Duration `json:"check_delay,omitempty"`

	CheckTimeout *time.Duration `json:"check_timeout,omitempty"`

	CheckMaxRetries int32 `json:"check_max_retries,omitempty"`

	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	TCPConfig *HealthCheckTCPConfig `json:"tcp_config,omitempty"`
	// MysqlConfig the check requires MySQL >=3.22, for older versions, use TCP check
	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	MysqlConfig *HealthCheckMysqlConfig `json:"mysql_config,omitempty"`

	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	PgsqlConfig *HealthCheckPgsqlConfig `json:"pgsql_config,omitempty"`
	// LdapConfig the response is analyzed to find an LDAPv3 response message
	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	LdapConfig *HealthCheckLdapConfig `json:"ldap_config,omitempty"`
	// RedisConfig the response is analyzed to find the +PONG response message
	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	RedisConfig *HealthCheckRedisConfig `json:"redis_config,omitempty"`

	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	HTTPConfig *HealthCheckHTTPConfig `json:"http_config,omitempty"`

	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	HTTPSConfig *HealthCheckHTTPSConfig `json:"https_config,omitempty"`
}

func (m *HealthCheck) GetConfig() Config {
	switch {
	case m.TCPConfig != nil:
		return ConfigTCPConfig{*m.TCPConfig}
	case m.MysqlConfig != nil:
		return ConfigMysqlConfig{*m.MysqlConfig}
	case m.PgsqlConfig != nil:
		return ConfigPgsqlConfig{*m.PgsqlConfig}
	case m.LdapConfig != nil:
		return ConfigLdapConfig{*m.LdapConfig}
	case m.RedisConfig != nil:
		return ConfigRedisConfig{*m.RedisConfig}
	case m.HTTPConfig != nil:
		return ConfigHTTPConfig{*m.HTTPConfig}
	case m.HTTPSConfig != nil:
		return ConfigHTTPSConfig{*m.HTTPSConfig}
	}
	return nil
}

func (m *HealthCheck) UnmarshalJSON(b []byte) error {
	type tmpType HealthCheck
	tmp := struct {
		tmpType

		TmpCheckDelay   *marshaler.Duration `json:"check_delay,omitempty"`
		TmpCheckTimeout *marshaler.Duration `json:"check_timeout,omitempty"`
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

		TmpCheckDelay   *marshaler.Duration `json:"check_delay,omitempty"`
		TmpCheckTimeout *marshaler.Duration `json:"check_timeout,omitempty"`
	}{
		tmpType: tmpType(m),

		TmpCheckDelay:   marshaler.NewDuration(m.CheckDelay),
		TmpCheckTimeout: marshaler.NewDuration(m.CheckTimeout),
	}
	return json.Marshal(tmp)
}

type HealthCheckHTTPConfig struct {
	URI string `json:"uri,omitempty"`

	Method string `json:"method,omitempty"`

	Code *int32 `json:"code,omitempty"`
}

type HealthCheckHTTPSConfig struct {
	URI string `json:"uri,omitempty"`

	Method string `json:"method,omitempty"`

	Code *int32 `json:"code,omitempty"`
}

type HealthCheckLdapConfig struct {
}

type HealthCheckMysqlConfig struct {
	User string `json:"user,omitempty"`
}

type HealthCheckPgsqlConfig struct {
	User string `json:"user,omitempty"`
}

type HealthCheckRedisConfig struct {
}

type HealthCheckTCPConfig struct {
}

type IP struct {
	ID string `json:"id,omitempty"`

	IPAddress string `json:"ip_address,omitempty"`

	OrganizationID string `json:"organization_id,omitempty"`

	LbID *string `json:"lb_id,omitempty"`

	Reverse string `json:"reverse,omitempty"`

	Region utils.Region `json:"region,omitempty"`
}

type Instance struct {
	ID string `json:"id,omitempty"`
	// Status
	//
	// Default value: unknown
	Status InstanceStatus `json:"status,omitempty"`

	IPAddress string `json:"ip_address,omitempty"`

	Region utils.Region `json:"region,omitempty"`
}

type Lb struct {
	ID string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`
	// Status
	//
	// Default value: unknown
	Status LbStatus `json:"status,omitempty"`

	Instances []*Instance `json:"instances,omitempty"`

	OrganizationID string `json:"organization_id,omitempty"`

	IP []*IP `json:"ip,omitempty"`

	Tags []string `json:"tags,omitempty"`

	FrontendCount int32 `json:"frontend_count,omitempty"`

	BackendCount int32 `json:"backend_count,omitempty"`

	Region utils.Region `json:"region,omitempty"`
}

type LbStats struct {
	// BackendServersStats list stats object of your loadbalancer (See the BackendServerStats object description)
	BackendServersStats []*BackendServerStats `json:"backend_servers_stats,omitempty"`
}

type ListACLResponse struct {
	// Acls list of Acl object (see Acl object description)
	Acls []*ACL `json:"acls,omitempty"`
	// TotalCount result count
	TotalCount uint32 `json:"total_count,omitempty"`
}

type ListBackendsResponse struct {
	// Backends list Backend objects of a Load Balancer
	Backends []*Backend `json:"backends,omitempty"`
	// TotalCount total count, wihtout pagination
	TotalCount uint32 `json:"total_count,omitempty"`
}

type ListFrontendsResponse struct {
	// Frontends list frontends object of your loadbalancer
	Frontends []*Frontend `json:"frontends,omitempty"`
	// TotalCount total count, wihtout pagination
	TotalCount uint32 `json:"total_count,omitempty"`
}

type ListIpsResponse struct {
	// Ips list IP address object
	Ips []*IP `json:"ips,omitempty"`
	// TotalCount total count, wihtout pagination
	TotalCount uint32 `json:"total_count,omitempty"`
}

type ListLbsResponse struct {
	Lbs []*Lb `json:"lbs,omitempty"`

	TotalCount uint32 `json:"total_count,omitempty"`
}

// Service API

type GetServiceInfoRequest struct {
	Region utils.Region `json:"-"`
}

func (s *API) GetServiceInfo(req *GetServiceInfoRequest, opts ...scw.RequestOption) (*utils.ServiceInfo, error) {
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
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "",
		Headers: http.Header{},
	}

	var resp utils.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListLbsRequest struct {
	Region utils.Region `json:"-"`
	// Name use this to search by name
	Name *string `json:"-"`
	// OrderBy
	//
	// Default value: created_at_asc
	OrderBy ListLbsRequestOrderBy `json:"-"`

	PageSize *int32 `json:"-"`

	Page *int32 `json:"-"`

	OrganizationID *string `json:"-"`
}

func (s *API) ListLbs(req *ListLbsRequest, opts ...scw.RequestOption) (*ListLbsResponse, error) {
	var err error

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if (req.OrganizationID == nil || *req.OrganizationID == "") && exist {
		req.OrganizationID = &defaultOrganizationID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListLbsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLbsResponse) UnsafeGetTotalCount() int {
	return int(r.TotalCount)
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLbsResponse) UnsafeAppend(res interface{}) (int, scw.SdkError) {
	results, ok := res.(*ListLbsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Lbs = append(r.Lbs, results.Lbs...)
	r.TotalCount += uint32(len(results.Lbs))
	return len(results.Lbs), nil
}

type CreateLbRequest struct {
	Region utils.Region `json:"-"`
	// OrganizationID owner of resources
	OrganizationID string `json:"organization_id,omitempty"`
	// Name resource names
	Name string `json:"name,omitempty"`
	// Description resource description
	Description string `json:"description,omitempty"`
	// IPID just like for compute instances, when you destroy a Load Balancer, you can keep its highly available IP address and reuse it for another Load Balancer later.
	IPID *string `json:"ip_id,omitempty"`
	// Tags list of keyword
	Tags []string `json:"tags,omitempty"`
}

func (s *API) CreateLb(req *CreateLbRequest, opts ...scw.RequestOption) (*Lb, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Lb

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetLbRequest struct {
	Region utils.Region `json:"-"`

	LbID string `json:"-"`
}

func (s *API) GetLb(req *GetLbRequest, opts ...scw.RequestOption) (*Lb, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LbID) == "" {
		return nil, errors.New("field LbID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbID) + "",
		Headers: http.Header{},
	}

	var resp Lb

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateLbRequest struct {
	Region utils.Region `json:"-"`
	// LbID load Balancer ID
	LbID string `json:"-"`
	// Name resource name
	Name string `json:"name,omitempty"`
	// Description resource description
	Description string `json:"description,omitempty"`
	// Tags list of keywords
	Tags []string `json:"tags,omitempty"`
}

func (s *API) UpdateLb(req *UpdateLbRequest, opts ...scw.RequestOption) (*Lb, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LbID) == "" {
		return nil, errors.New("field LbID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Lb

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteLbRequest struct {
	Region utils.Region `json:"-"`
	// LbID load Balancer ID
	LbID string `json:"-"`
	// ReleaseIP set true if you don't want to keep this IP address
	ReleaseIP bool `json:"-"`
}

func (s *API) DeleteLb(req *DeleteLbRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "release_ip", req.ReleaseIP)

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LbID) == "" {
		return errors.New("field LbID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbID) + "",
		Query:   query,
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type ListIPsRequest struct {
	Region utils.Region `json:"-"`
	// Page page number
	Page *int32 `json:"-"`
	// PageSize set the maximum list size
	PageSize *int32 `json:"-"`
	// IPAddress use this to search by IP address
	IPAddress *string `json:"-"`

	OrganizationID *string `json:"-"`
}

// ListIPs list IPs
func (s *API) ListIPs(req *ListIPsRequest, opts ...scw.RequestOption) (*ListIpsResponse, error) {
	var err error

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if (req.OrganizationID == nil || *req.OrganizationID == "") && exist {
		req.OrganizationID = &defaultOrganizationID
	}

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
	parameter.AddToQuery(query, "ip_address", req.IPAddress)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListIpsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIpsResponse) UnsafeGetTotalCount() int {
	return int(r.TotalCount)
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIpsResponse) UnsafeAppend(res interface{}) (int, scw.SdkError) {
	results, ok := res.(*ListIpsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Ips = append(r.Ips, results.Ips...)
	r.TotalCount += uint32(len(results.Ips))
	return len(results.Ips), nil
}

type GetIPRequest struct {
	Region utils.Region `json:"-"`
	// IPID
	//
	// IP address ID
	IPID string `json:"-"`
}

// GetIP get IP
func (s *API) GetIP(req *GetIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "",
		Headers: http.Header{},
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ReleaseIPRequest struct {
	Region utils.Region `json:"-"`
	// IPID iP address ID
	IPID string `json:"-"`
}

// ReleaseIP release IP
func (s *API) ReleaseIP(req *ReleaseIPRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type UpdateIPRequest struct {
	Region utils.Region `json:"-"`
	// IPID iP address ID
	IPID string `json:"-"`
	// Reverse reverse DNS
	Reverse *string `json:"-"`
}

func (s *API) UpdateIP(req *UpdateIPRequest, opts ...scw.RequestOption) (*IP, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	query := url.Values{}
	parameter.AddToQuery(query, "reverse", req.Reverse)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.IPID) == "" {
		return nil, errors.New("field IPID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IPID) + "",
		Query:   query,
		Headers: http.Header{},
	}

	var resp IP

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListBackendsRequest struct {
	Region utils.Region `json:"-"`
	// LbID load Balancer ID
	LbID string `json:"-"`
	// Name use this to search by name
	Name *string `json:"-"`
	// OrderBy choose order of response
	//
	// Default value: created_at_asc
	OrderBy ListBackendsRequestOrderBy `json:"-"`
	// Page page number
	Page *int32 `json:"-"`
	// PageSize set the maximum list sizes
	PageSize *int32 `json:"-"`
}

func (s *API) ListBackends(req *ListBackendsRequest, opts ...scw.RequestOption) (*ListBackendsResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LbID) == "" {
		return nil, errors.New("field LbID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbID) + "/backends",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListBackendsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBackendsResponse) UnsafeGetTotalCount() int {
	return int(r.TotalCount)
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBackendsResponse) UnsafeAppend(res interface{}) (int, scw.SdkError) {
	results, ok := res.(*ListBackendsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Backends = append(r.Backends, results.Backends...)
	r.TotalCount += uint32(len(results.Backends))
	return len(results.Backends), nil
}

type CreateBackendRequest struct {
	Region utils.Region `json:"-"`
	// LbID load Balancer ID
	LbID string `json:"-"`
	// Name resource name
	Name string `json:"name,omitempty"`
	// ForwardProtocol backend protocol. TCP or HTTP
	//
	// Default value: tcp
	ForwardProtocol Protocol `json:"forward_protocol,omitempty"`
	// ForwardPort user sessions will be forwarded to this port of backend servers
	ForwardPort int32 `json:"forward_port,omitempty"`
	// ForwardPortAlgorithm load balancing algorithm
	//
	// Default value: roundrobin
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm,omitempty"`
	// StickySessions enables cookie-based session persistence
	//
	// Default value: none
	StickySessions StickySessionsType `json:"sticky_sessions,omitempty"`
	// StickySessionsCookieName cookie name for for sticky sessions
	StickySessionsCookieName string `json:"sticky_sessions_cookie_name,omitempty"`
	// HealthCheck see the Healthcheck object description
	HealthCheck *HealthCheck `json:"health_check,omitempty"`
	// ServerIP backend server IP addresses list (IPv4 or IPv6)
	ServerIP []string `json:"server_ip,omitempty"`
	// SendProxyV2 enables PROXY protocol version 2 (must be supported by backend servers)
	SendProxyV2 bool `json:"send_proxy_v2,omitempty"`
	// TimeoutServer maximum server connection inactivity time
	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`
	// TimeoutConnect maximum initical server connection establishment time
	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`
	// TimeoutTunnel maximum tunnel inactivity time
	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`
	// OnMarkedDownAction modify what occurs when a backend server is marked down
	//
	// Default value: on_marked_down_action_none
	OnMarkedDownAction OnMarkedDownAction `json:"on_marked_down_action,omitempty"`
}

func (m *CreateBackendRequest) UnmarshalJSON(b []byte) error {
	type tmpType CreateBackendRequest
	tmp := struct {
		tmpType

		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = CreateBackendRequest(tmp.tmpType)

	m.TimeoutServer = tmp.TmpTimeoutServer.Standard()
	m.TimeoutConnect = tmp.TmpTimeoutConnect.Standard()
	m.TimeoutTunnel = tmp.TmpTimeoutTunnel.Standard()
	return nil
}

func (m CreateBackendRequest) MarshalJSON() ([]byte, error) {
	type tmpType CreateBackendRequest
	tmp := struct {
		tmpType

		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{
		tmpType: tmpType(m),

		TmpTimeoutServer:  marshaler.NewDuration(m.TimeoutServer),
		TmpTimeoutConnect: marshaler.NewDuration(m.TimeoutConnect),
		TmpTimeoutTunnel:  marshaler.NewDuration(m.TimeoutTunnel),
	}
	return json.Marshal(tmp)
}

func (s *API) CreateBackend(req *CreateBackendRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LbID) == "" {
		return nil, errors.New("field LbID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbID) + "/backends",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetBackendRequest struct {
	Region utils.Region `json:"-"`
	// BackendID backend ID
	BackendID string `json:"-"`
}

func (s *API) GetBackend(req *GetBackendRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "",
		Headers: http.Header{},
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateBackendRequest struct {
	Region utils.Region `json:"-"`
	// BackendID backend ID to update
	BackendID string `json:"-"`
	// Name resource name
	Name string `json:"name,omitempty"`
	// ForwardProtocol backend protocol. TCP or HTTP
	//
	// Default value: tcp
	ForwardProtocol Protocol `json:"forward_protocol,omitempty"`
	// ForwardPort user sessions will be forwarded to this port of backend servers
	ForwardPort int32 `json:"forward_port,omitempty"`
	// ForwardPortAlgorithm load balancing algorithm
	//
	// Default value: roundrobin
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm,omitempty"`
	// StickySessions enable cookie-based session persistence
	//
	// Default value: none
	StickySessions StickySessionsType `json:"sticky_sessions,omitempty"`
	// StickySessionsCookieName cookie name for for sticky sessions
	StickySessionsCookieName string `json:"sticky_sessions_cookie_name,omitempty"`
	// SendProxyV2 enables PROXY protocol version 2 (must be supported by backend servers)
	SendProxyV2 bool `json:"send_proxy_v2,omitempty"`
	// TimeoutServer maximum server connection inactivity time
	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`
	// TimeoutConnect maximum initial server connection establishment time
	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`
	// TimeoutTunnel maximum tunnel inactivity time
	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`
	// OnMarkedDownAction modify what occurs when a backend server is marked down
	//
	// Default value: on_marked_down_action_none
	OnMarkedDownAction OnMarkedDownAction `json:"on_marked_down_action,omitempty"`
}

func (m *UpdateBackendRequest) UnmarshalJSON(b []byte) error {
	type tmpType UpdateBackendRequest
	tmp := struct {
		tmpType

		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = UpdateBackendRequest(tmp.tmpType)

	m.TimeoutServer = tmp.TmpTimeoutServer.Standard()
	m.TimeoutConnect = tmp.TmpTimeoutConnect.Standard()
	m.TimeoutTunnel = tmp.TmpTimeoutTunnel.Standard()
	return nil
}

func (m UpdateBackendRequest) MarshalJSON() ([]byte, error) {
	type tmpType UpdateBackendRequest
	tmp := struct {
		tmpType

		TmpTimeoutServer  *marshaler.Duration `json:"timeout_server,omitempty"`
		TmpTimeoutConnect *marshaler.Duration `json:"timeout_connect,omitempty"`
		TmpTimeoutTunnel  *marshaler.Duration `json:"timeout_tunnel,omitempty"`
	}{
		tmpType: tmpType(m),

		TmpTimeoutServer:  marshaler.NewDuration(m.TimeoutServer),
		TmpTimeoutConnect: marshaler.NewDuration(m.TimeoutConnect),
		TmpTimeoutTunnel:  marshaler.NewDuration(m.TimeoutTunnel),
	}
	return json.Marshal(tmp)
}

func (s *API) UpdateBackend(req *UpdateBackendRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteBackendRequest struct {
	Region utils.Region `json:"-"`
	// BackendID iD of the backend to delete
	BackendID string `json:"-"`
}

func (s *API) DeleteBackend(req *DeleteBackendRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type AddBackendServersRequest struct {
	Region utils.Region `json:"-"`
	// BackendID backend ID
	BackendID string `json:"-"`
	// ServerIP set all IPs to remove of your backend
	ServerIP []string `json:"server_ip,omitempty"`
}

func (s *API) AddBackendServers(req *AddBackendServersRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "/servers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RemoveBackendServersRequest struct {
	Region utils.Region `json:"-"`
	// BackendID backend ID
	BackendID string `json:"-"`
	// ServerIP set all IPs to remove of your backend
	ServerIP []string `json:"server_ip,omitempty"`
}

func (s *API) RemoveBackendServers(req *RemoveBackendServersRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "/servers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetBackendServersRequest struct {
	Region utils.Region `json:"-"`
	// BackendID backend ID
	BackendID string `json:"-"`
	// ServerIP set all IPs to add of your backend and remove all other
	ServerIP []string `json:"server_ip,omitempty"`
}

func (s *API) SetBackendServers(req *SetBackendServersRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "/servers",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Backend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateHealthCheckRequest struct {
	Region utils.Region `json:"-"`
	// BackendID backend ID
	BackendID string `json:"-"`
	// Port specify the port used to health check
	Port int32 `json:"port,omitempty"`
	// CheckDelay time between two consecutive health checks
	CheckDelay *time.Duration `json:"check_delay,omitempty"`
	// CheckTimeout additional check timeout, after the connection has been already established
	CheckTimeout *time.Duration `json:"check_timeout,omitempty"`
	// CheckMaxRetries number of consecutive unsuccessful health checks, after wich the server will be considered dead
	CheckMaxRetries int32 `json:"check_max_retries,omitempty"`
	// MysqlConfig the check requires MySQL >=3.22, for older version, please use TCP check
	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	MysqlConfig *HealthCheckMysqlConfig `json:"mysql_config,omitempty"`
	// LdapConfig the response is analyzed to find an LDAPv3 response message
	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	LdapConfig *HealthCheckLdapConfig `json:"ldap_config,omitempty"`
	// RedisConfig the response is analyzed to find the +PONG response message
	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	RedisConfig *HealthCheckRedisConfig `json:"redis_config,omitempty"`

	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	PgsqlConfig *HealthCheckPgsqlConfig `json:"pgsql_config,omitempty"`

	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	TCPConfig *HealthCheckTCPConfig `json:"tcp_config,omitempty"`

	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	HTTPConfig *HealthCheckHTTPConfig `json:"http_config,omitempty"`

	// Precisely one of HTTPConfig, HTTPSConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TCPConfig must be set.
	HTTPSConfig *HealthCheckHTTPSConfig `json:"https_config,omitempty"`
}

func (m *UpdateHealthCheckRequest) GetConfig() Config {
	switch {
	case m.MysqlConfig != nil:
		return ConfigMysqlConfig{*m.MysqlConfig}
	case m.LdapConfig != nil:
		return ConfigLdapConfig{*m.LdapConfig}
	case m.RedisConfig != nil:
		return ConfigRedisConfig{*m.RedisConfig}
	case m.PgsqlConfig != nil:
		return ConfigPgsqlConfig{*m.PgsqlConfig}
	case m.TCPConfig != nil:
		return ConfigTCPConfig{*m.TCPConfig}
	case m.HTTPConfig != nil:
		return ConfigHTTPConfig{*m.HTTPConfig}
	case m.HTTPSConfig != nil:
		return ConfigHTTPSConfig{*m.HTTPSConfig}
	}
	return nil
}

func (m *UpdateHealthCheckRequest) UnmarshalJSON(b []byte) error {
	type tmpType UpdateHealthCheckRequest
	tmp := struct {
		tmpType

		TmpCheckDelay   *marshaler.Duration `json:"check_delay,omitempty"`
		TmpCheckTimeout *marshaler.Duration `json:"check_timeout,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = UpdateHealthCheckRequest(tmp.tmpType)

	m.CheckDelay = tmp.TmpCheckDelay.Standard()
	m.CheckTimeout = tmp.TmpCheckTimeout.Standard()
	return nil
}

func (m UpdateHealthCheckRequest) MarshalJSON() ([]byte, error) {
	type tmpType UpdateHealthCheckRequest
	tmp := struct {
		tmpType

		TmpCheckDelay   *marshaler.Duration `json:"check_delay,omitempty"`
		TmpCheckTimeout *marshaler.Duration `json:"check_timeout,omitempty"`
	}{
		tmpType: tmpType(m),

		TmpCheckDelay:   marshaler.NewDuration(m.CheckDelay),
		TmpCheckTimeout: marshaler.NewDuration(m.CheckTimeout),
	}
	return json.Marshal(tmp)
}

func (s *API) UpdateHealthCheck(req *UpdateHealthCheckRequest, opts ...scw.RequestOption) (*HealthCheck, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackendID) == "" {
		return nil, errors.New("field BackendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendID) + "/healthcheck",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp HealthCheck

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListFrontendsRequest struct {
	Region utils.Region `json:"-"`
	// LbID load Balancer ID
	LbID string `json:"-"`
	// Name use this to search by name
	Name *string `json:"-"`
	// OrderBy response order
	//
	// Default value: created_at_asc
	OrderBy ListFrontendsRequestOrderBy `json:"-"`
	// Page page number
	Page *int32 `json:"-"`
	// PageSize set the maximum list sizes
	PageSize *int32 `json:"-"`
}

func (s *API) ListFrontends(req *ListFrontendsRequest, opts ...scw.RequestOption) (*ListFrontendsResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LbID) == "" {
		return nil, errors.New("field LbID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbID) + "/frontends",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListFrontendsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFrontendsResponse) UnsafeGetTotalCount() int {
	return int(r.TotalCount)
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFrontendsResponse) UnsafeAppend(res interface{}) (int, scw.SdkError) {
	results, ok := res.(*ListFrontendsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Frontends = append(r.Frontends, results.Frontends...)
	r.TotalCount += uint32(len(results.Frontends))
	return len(results.Frontends), nil
}

type CreateFrontendRequest struct {
	Region utils.Region `json:"-"`
	// LbID load Balancer ID
	LbID string `json:"-"`
	// Name resource name
	Name string `json:"name,omitempty"`
	// InboundPort tCP port to listen on the front side
	InboundPort int32 `json:"inbound_port,omitempty"`
	// BackendID backend ID
	BackendID string `json:"backend_id,omitempty"`
	// TimeoutClient set the maximum inactivity time on the client side
	TimeoutClient *time.Duration `json:"timeout_client,omitempty"`
}

func (m *CreateFrontendRequest) UnmarshalJSON(b []byte) error {
	type tmpType CreateFrontendRequest
	tmp := struct {
		tmpType

		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = CreateFrontendRequest(tmp.tmpType)

	m.TimeoutClient = tmp.TmpTimeoutClient.Standard()
	return nil
}

func (m CreateFrontendRequest) MarshalJSON() ([]byte, error) {
	type tmpType CreateFrontendRequest
	tmp := struct {
		tmpType

		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{
		tmpType: tmpType(m),

		TmpTimeoutClient: marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

func (s *API) CreateFrontend(req *CreateFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LbID) == "" {
		return nil, errors.New("field LbID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbID) + "/frontends",
		Headers: http.Header{},
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

type GetFrontendRequest struct {
	Region utils.Region `json:"-"`
	// FrontendID frontend ID
	FrontendID string `json:"-"`
}

func (s *API) GetFrontend(req *GetFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendID) + "",
		Headers: http.Header{},
	}

	var resp Frontend

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateFrontendRequest struct {
	Region utils.Region `json:"-"`
	// FrontendID frontend ID
	FrontendID string `json:"-"`
	// Name resource name
	Name string `json:"name,omitempty"`
	// InboundPort tCP port to listen on the front side
	InboundPort int32 `json:"inbound_port,omitempty"`
	// BackendID backend ID
	BackendID string `json:"backend_id,omitempty"`
	// TimeoutClient client session maximum inactivity time
	TimeoutClient *time.Duration `json:"timeout_client,omitempty"`
}

func (m *UpdateFrontendRequest) UnmarshalJSON(b []byte) error {
	type tmpType UpdateFrontendRequest
	tmp := struct {
		tmpType

		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{}
	err := json.Unmarshal(b, &tmp)
	if err != nil {
		return err
	}

	*m = UpdateFrontendRequest(tmp.tmpType)

	m.TimeoutClient = tmp.TmpTimeoutClient.Standard()
	return nil
}

func (m UpdateFrontendRequest) MarshalJSON() ([]byte, error) {
	type tmpType UpdateFrontendRequest
	tmp := struct {
		tmpType

		TmpTimeoutClient *marshaler.Duration `json:"timeout_client,omitempty"`
	}{
		tmpType: tmpType(m),

		TmpTimeoutClient: marshaler.NewDuration(m.TimeoutClient),
	}
	return json.Marshal(tmp)
}

func (s *API) UpdateFrontend(req *UpdateFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendID) + "",
		Headers: http.Header{},
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

type DeleteFrontendRequest struct {
	Region utils.Region `json:"-"`
	// FrontendID frontend ID to delete
	FrontendID string `json:"-"`
}

func (s *API) DeleteFrontend(req *DeleteFrontendRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type GetLbStatsRequest struct {
	Region utils.Region `json:"-"`
	// LbID load Balancer ID
	LbID string `json:"-"`
}

func (s *API) GetLbStats(req *GetLbStatsRequest, opts ...scw.RequestOption) (*LbStats, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.LbID) == "" {
		return nil, errors.New("field LbID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbID) + "/stats",
		Headers: http.Header{},
	}

	var resp LbStats

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListAclsRequest struct {
	Region utils.Region `json:"-"`
	// FrontendID iD of your frontend
	FrontendID string `json:"-"`
	// OrderBy you can order the response by created_at asc/desc or name asc/desc
	//
	// Default value: created_at_asc
	OrderBy ListACLRequestOrderBy `json:"-"`
	// Page page number
	Page *int32 `json:"-"`
	// PageSize set the maximum list size
	PageSize *int32 `json:"-"`
	// Name filter acl per name
	Name *string `json:"-"`
}

func (s *API) ListAcls(req *ListAclsRequest, opts ...scw.RequestOption) (*ListACLResponse, error) {
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
	parameter.AddToQuery(query, "name", req.Name)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendID) + "/acls",
		Query:   query,
		Headers: http.Header{},
	}

	var resp ListACLResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListACLResponse) UnsafeGetTotalCount() int {
	return int(r.TotalCount)
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListACLResponse) UnsafeAppend(res interface{}) (int, scw.SdkError) {
	results, ok := res.(*ListACLResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Acls = append(r.Acls, results.Acls...)
	r.TotalCount += uint32(len(results.Acls))
	return len(results.Acls), nil
}

type CreateACLRequest struct {
	Region utils.Region `json:"-"`
	// FrontendID iD of your frontend
	FrontendID string `json:"-"`
	// Name name of your ACL ressource
	Name string `json:"name,omitempty"`
	// Action see the AclAction object description
	Action *ACLAction `json:"action,omitempty"`
	// Match see the AclMatch object description
	Match *ACLMatch `json:"match,omitempty"`
	// Index order between your Acls (ascending order, 0 is first acl executed)
	Index int32 `json:"index,omitempty"`
}

func (s *API) CreateACL(req *CreateACLRequest, opts ...scw.RequestOption) (*ACL, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FrontendID) == "" {
		return nil, errors.New("field FrontendID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendID) + "/acls",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ACL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetACLRequest struct {
	Region utils.Region `json:"-"`
	// ACLID iD of your ACL ressource
	ACLID string `json:"-"`
}

func (s *API) GetACL(req *GetACLRequest, opts ...scw.RequestOption) (*ACL, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return nil, errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/acls/" + fmt.Sprint(req.ACLID) + "",
		Headers: http.Header{},
	}

	var resp ACL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateACLRequest struct {
	Region utils.Region `json:"-"`
	// ACLID iD of your ACL ressource
	ACLID string `json:"-"`
	// Name name of your ACL ressource
	Name string `json:"name,omitempty"`
	// Action see the AclAction object description
	Action *ACLAction `json:"action,omitempty"`
	// Match see the AclMatch object description
	Match *ACLMatch `json:"match,omitempty"`
	// Index order between your Acls (ascending order, 0 is first acl executed)
	Index int32 `json:"index,omitempty"`
}

func (s *API) UpdateACL(req *UpdateACLRequest, opts ...scw.RequestOption) (*ACL, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return nil, errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/acls/" + fmt.Sprint(req.ACLID) + "",
		Headers: http.Header{},
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ACL

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteACLRequest struct {
	Region utils.Region `json:"-"`
	// ACLID iD of your ACL ressource
	ACLID string `json:"-"`
}

func (s *API) DeleteACL(req *DeleteACLRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.ACLID) == "" {
		return errors.New("field ACLID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/acls/" + fmt.Sprint(req.ACLID) + "",
		Headers: http.Header{},
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

type Config interface {
	isConfig()
}

type ConfigMysqlConfig struct {
	Value HealthCheckMysqlConfig
}

func (ConfigMysqlConfig) isConfig() {
}

type ConfigLdapConfig struct {
	Value HealthCheckLdapConfig
}

func (ConfigLdapConfig) isConfig() {
}

type ConfigRedisConfig struct {
	Value HealthCheckRedisConfig
}

func (ConfigRedisConfig) isConfig() {
}

type ConfigPgsqlConfig struct {
	Value HealthCheckPgsqlConfig
}

func (ConfigPgsqlConfig) isConfig() {
}

type ConfigTCPConfig struct {
	Value HealthCheckTCPConfig
}

func (ConfigTCPConfig) isConfig() {
}

type ConfigHTTPConfig struct {
	Value HealthCheckHTTPConfig
}

func (ConfigHTTPConfig) isConfig() {
}

type ConfigHTTPSConfig struct {
	Value HealthCheckHTTPSConfig
}

func (ConfigHTTPSConfig) isConfig() {
}

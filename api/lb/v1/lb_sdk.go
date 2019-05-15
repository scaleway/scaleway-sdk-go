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

// Api: this API allows you to manage your Load Balancer service
type Api struct {
	client *scw.Client
}

// NewApi returns a Api object from a Scaleway client.
func NewApi(client *scw.Client) *Api {
	return &Api{
		client: client,
	}
}

type AclActionType string

const (
	// AclActionTypeAllow is [insert doc].
	AclActionTypeAllow = AclActionType("allow")
	// AclActionTypeDeny is [insert doc].
	AclActionTypeDeny = AclActionType("deny")
)

type AclHttpFilter string

const (
	// AclHttpFilterAclHttpFilterNone is [insert doc].
	AclHttpFilterAclHttpFilterNone = AclHttpFilter("acl_http_filter_none")
	// AclHttpFilterPathBegin is [insert doc].
	AclHttpFilterPathBegin = AclHttpFilter("path_begin")
	// AclHttpFilterPathEnd is [insert doc].
	AclHttpFilterPathEnd = AclHttpFilter("path_end")
	// AclHttpFilterRegex is [insert doc].
	AclHttpFilterRegex = AclHttpFilter("regex")
)

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

type ForwardPortAlgorithm string

const (
	// ForwardPortAlgorithmRoundrobin is [insert doc].
	ForwardPortAlgorithmRoundrobin = ForwardPortAlgorithm("roundrobin")
	// ForwardPortAlgorithmLeastconn is [insert doc].
	ForwardPortAlgorithmLeastconn = ForwardPortAlgorithm("leastconn")
)

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

type ListAclRequestOrderBy string

const (
	// ListAclRequestOrderByCreatedAtAsc is [insert doc].
	ListAclRequestOrderByCreatedAtAsc = ListAclRequestOrderBy("created_at_asc")
	// ListAclRequestOrderByCreatedAtDesc is [insert doc].
	ListAclRequestOrderByCreatedAtDesc = ListAclRequestOrderBy("created_at_desc")
	// ListAclRequestOrderByNameAsc is [insert doc].
	ListAclRequestOrderByNameAsc = ListAclRequestOrderBy("name_asc")
	// ListAclRequestOrderByNameDesc is [insert doc].
	ListAclRequestOrderByNameDesc = ListAclRequestOrderBy("name_desc")
)

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

type OnMarkedDownAction string

const (
	// OnMarkedDownActionOnMarkedDownActionNone is [insert doc].
	OnMarkedDownActionOnMarkedDownActionNone = OnMarkedDownAction("on_marked_down_action_none")
	// OnMarkedDownActionShutdownSessions is [insert doc].
	OnMarkedDownActionShutdownSessions = OnMarkedDownAction("shutdown_sessions")
)

type Protocol string

const (
	// ProtocolTcp is [insert doc].
	ProtocolTcp = Protocol("tcp")
	// ProtocolHttp is [insert doc].
	ProtocolHttp = Protocol("http")
)

type StickySessionsType string

const (
	// StickySessionsTypeNone is [insert doc].
	StickySessionsTypeNone = StickySessionsType("none")
	// StickySessionsTypeCookie is [insert doc].
	StickySessionsTypeCookie = StickySessionsType("cookie")
	// StickySessionsTypeTable is [insert doc].
	StickySessionsTypeTable = StickySessionsType("table")
)

// Acl: the use of Access Control Lists (ACL) provide a flexible solution to perform a action generally consist in blocking or allow a request based on ip (and URL on HTTP)
type Acl struct {
	// Id: iD of your ACL ressource
	Id string `json:"id,omitempty"`
	// Name: name of you ACL ressource
	Name string `json:"name,omitempty"`
	// Match: see the AclMatch object description
	Match *AclMatch `json:"match,omitempty"`
	// Action: see the AclAction object description
	Action *AclAction `json:"action,omitempty"`
	// Frontend: see the Frontend object description
	Frontend *Frontend `json:"frontend,omitempty"`
	// Index: order between your Acls (ascending order, 0 is first acl executed)
	Index int32 `json:"index,omitempty"`
}

// AclAction: action if your ACL filter match
type AclAction struct {
	// Type: <allow> or <deny> request
	Type AclActionType `json:"type,omitempty"`
}

// AclMatch: settings of your ACL filter
type AclMatch struct {
	// IpSubnet: this is the source IP v4/v6 address of the client of the session to match or not. Addresses values can be specified either as plain addresses or with a netmask appended
	IpSubnet []*string `json:"ip_subnet,omitempty"`
	// HttpFilter: you can set http filter (if your backend protocole have a http forward protocol. This extracts the request's URL path, which starts at the first slash and ends before the question mark (without the host part). You can choose between <path_begin> prefix match (like /admin), <path_end> suffix match (like .php) and <regex>
	HttpFilter AclHttpFilter `json:"http_filter,omitempty"`

	HttpFilterValue []*string `json:"http_filter_value,omitempty"`
	// Invert: by default match filter is a IF condition. You can set invert to true to have a unless condition
	Invert bool `json:"invert,omitempty"`
}

type Backend struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	ForwardProtocol Protocol `json:"forward_protocol,omitempty"`

	ForwardPort int32 `json:"forward_port,omitempty"`

	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm,omitempty"`

	StickySessions StickySessionsType `json:"sticky_sessions,omitempty"`

	StickySessionsCookieName string `json:"sticky_sessions_cookie_name,omitempty"`

	HealthCheck *HealthCheck `json:"health_check,omitempty"`

	Pool []string `json:"pool,omitempty"`

	Lb *Lb `json:"lb,omitempty"`

	SendProxyV2 bool `json:"send_proxy_v2,omitempty"`

	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`

	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`

	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`

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

// BackendServerStats: state and statistics of your backend server like last healthcheck status, server uptime, result state of your backend server
type BackendServerStats struct {
	// InstanceId: iD of your loadbalancer cluster server
	InstanceId string `json:"instance_id,omitempty"`
	// BackendId: iD of your Backend
	BackendId string `json:"backend_id,omitempty"`
	// Ip: iPv4 or IPv6 address of the server backend
	Ip string `json:"ip,omitempty"`
	// ServerState: server operational state (stopped/starting/running/stopping)
	ServerState BackendServerStatsServerState `json:"server_state,omitempty"`
	// ServerStateChangedAt: time since last operational change
	ServerStateChangedAt time.Time `json:"server_state_changed_at,omitempty"`
	// LastHealthCheckStatus: last health check status (unknown/neutral/failed/passed/condpass)
	LastHealthCheckStatus BackendServerStatsHealthCheckStatus `json:"last_health_check_status,omitempty"`
}

type Frontend struct {
	Id string `json:"id,omitempty"`

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

	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	TcpConfig *HealthCheckTcpConfig `json:"tcp_config,omitempty"`
	// MysqlConfig: the check requires MySQL >=3.22, for older versions, use TCP check
	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	MysqlConfig *HealthCheckMysqlConfig `json:"mysql_config,omitempty"`

	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	PgsqlConfig *HealthCheckPgsqlConfig `json:"pgsql_config,omitempty"`
	// LdapConfig: the response is analyzed to find an LDAPv3 response message
	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	LdapConfig *HealthCheckLdapConfig `json:"ldap_config,omitempty"`
	// RedisConfig: the response is analyzed to find the +PONG response message
	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	RedisConfig *HealthCheckRedisConfig `json:"redis_config,omitempty"`

	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	HttpConfig *HealthCheckHttpConfig `json:"http_config,omitempty"`

	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	HttpsConfig *HealthCheckHttpsConfig `json:"https_config,omitempty"`
}

func (m *HealthCheck) GetConfig() Config {
	switch {
	case m.TcpConfig != nil:
		return ConfigTcpConfig{*m.TcpConfig}
	case m.MysqlConfig != nil:
		return ConfigMysqlConfig{*m.MysqlConfig}
	case m.PgsqlConfig != nil:
		return ConfigPgsqlConfig{*m.PgsqlConfig}
	case m.LdapConfig != nil:
		return ConfigLdapConfig{*m.LdapConfig}
	case m.RedisConfig != nil:
		return ConfigRedisConfig{*m.RedisConfig}
	case m.HttpConfig != nil:
		return ConfigHttpConfig{*m.HttpConfig}
	case m.HttpsConfig != nil:
		return ConfigHttpsConfig{*m.HttpsConfig}
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

type HealthCheckHttpConfig struct {
	Uri string `json:"uri,omitempty"`

	Method string `json:"method,omitempty"`

	Code *int32 `json:"code,omitempty"`
}

type HealthCheckHttpsConfig struct {
	Uri string `json:"uri,omitempty"`

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

type HealthCheckTcpConfig struct {
}

type Instance struct {
	Id string `json:"id,omitempty"`

	Status InstanceStatus `json:"status,omitempty"`

	IpAddress string `json:"ip_address,omitempty"`

	Region utils.Region `json:"region,omitempty"`
}

type Ip struct {
	Id string `json:"id,omitempty"`

	IpAddress string `json:"ip_address,omitempty"`

	OrganizationId string `json:"organization_id,omitempty"`

	LbId *string `json:"lb_id,omitempty"`

	Reverse string `json:"reverse,omitempty"`

	Region utils.Region `json:"region,omitempty"`
}

type Lb struct {
	Id string `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Description string `json:"description,omitempty"`

	Status LbStatus `json:"status,omitempty"`

	Instances []*Instance `json:"instances,omitempty"`

	OrganizationId string `json:"organization_id,omitempty"`

	Ip []*Ip `json:"ip,omitempty"`

	Tags []string `json:"tags,omitempty"`

	FrontendCount int32 `json:"frontend_count,omitempty"`

	BackendCount int32 `json:"backend_count,omitempty"`

	Region utils.Region `json:"region,omitempty"`
}

type LbStats struct {
	// BackendServersStats: list stats object of your loadbalancer (See the BackendServerStats object description)
	BackendServersStats []*BackendServerStats `json:"backend_servers_stats,omitempty"`
}

type ListAclResponse struct {
	// Acls: list of Acl object (see Acl object description)
	Acls []*Acl `json:"acls,omitempty"`
	// TotalCount: result count
	TotalCount uint32 `json:"total_count,omitempty"`
}

type ListBackendsResponse struct {
	// Backends: list Backend objects of a Load Balancer
	Backends []*Backend `json:"backends,omitempty"`
	// TotalCount: total count, wihtout pagination
	TotalCount uint32 `json:"total_count,omitempty"`
}

type ListFrontendsResponse struct {
	// Frontends: list frontends object of your loadbalancer
	Frontends []*Frontend `json:"frontends,omitempty"`
	// TotalCount: total count, wihtout pagination
	TotalCount uint32 `json:"total_count,omitempty"`
}

type ListIpsResponse struct {
	// Ips: list IP address object
	Ips []*Ip `json:"ips,omitempty"`
	// TotalCount: total count, wihtout pagination
	TotalCount uint32 `json:"total_count,omitempty"`
}

type ListLbsResponse struct {
	Lbs []*Lb `json:"lbs,omitempty"`

	TotalCount uint32 `json:"total_count,omitempty"`
}

// Service Api

type GetServiceInfoRequest struct {
	Region utils.Region `json:"-"`
}

func (s *Api) GetServiceInfo(req *GetServiceInfoRequest, opts ...scw.RequestOption) (*utils.ServiceInfo, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp utils.ServiceInfo
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListLbsRequest struct {
	Region utils.Region `json:"-"`
	// Name: use this to search by name
	Name *string `json:"-"`

	OrderBy ListLbsRequestOrderBy `json:"-"`

	PageSize *int32 `json:"-"`

	Page *int32 `json:"-"`

	OrganizationId *string `json:"-"`
}

func (s *Api) ListLbs(req *ListLbsRequest, opts ...scw.RequestOption) (*ListLbsResponse, error) {
	var err error

	val := s.client.GetDefaultOrganizationID()
	if req.OrganizationId == nil || *req.OrganizationId == "" {
		req.OrganizationId = &val
	}

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	query := url.Values{}
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "organization_id", req.OrganizationId)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListLbsResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateLbRequest struct {
	Region utils.Region `json:"-"`
	// OrganizationId: owner of resources
	OrganizationId string `json:"organization_id,omitempty"`
	// Name: resource names
	Name string `json:"name,omitempty"`
	// Description: resource description
	Description string `json:"description,omitempty"`
	// IpId: just like for compute instances, when you destroy a Load Balancer, you can keep its highly available IP address and reuse it for another Load Balancer later.
	IpId *string `json:"ip_id,omitempty"`
	// Tags: list of keyword
	Tags []string `json:"tags,omitempty"`
}

func (s *Api) CreateLb(req *CreateLbRequest, opts ...scw.RequestOption) (*Lb, error) {
	var err error

	if req.OrganizationId == "" {
		req.OrganizationId = s.client.GetDefaultOrganizationID()
	}

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
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

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Lb
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetLbRequest struct {
	Region utils.Region `json:"-"`

	LbId string `json:"-"`
}

func (s *Api) GetLb(req *GetLbRequest, opts ...scw.RequestOption) (*Lb, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Lb
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateLbRequest struct {
	Region utils.Region `json:"-"`
	// LbId: load Balancer ID
	LbId string `json:"-"`
	// Name: resource name
	Name string `json:"name,omitempty"`
	// Description: resource description
	Description string `json:"description,omitempty"`
	// Tags: list of keywords
	Tags []string `json:"tags,omitempty"`
}

func (s *Api) UpdateLb(req *UpdateLbRequest, opts ...scw.RequestOption) (*Lb, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbId) + "",
		Headers: http.Header{},
	}
	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Lb
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteLbRequest struct {
	Region utils.Region `json:"-"`
	// LbId: load Balancer ID
	LbId string `json:"-"`
	// ReleaseIp: set true if you don't want to keep this IP address
	ReleaseIp bool `json:"-"`
}

func (s *Api) DeleteLb(req *DeleteLbRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	query := url.Values{}
	parameter.AddToQuery(query, "release_ip", req.ReleaseIp)

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbId) + "",
		Query:   query,
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq, opts...)

	if err != nil {
		return err
	}
	return nil
}

type ListIPsRequest struct {
	Region utils.Region `json:"-"`
	// Page: page number
	Page *int32 `json:"-"`
	// PageSize: set the maximum list size
	PageSize *int32 `json:"-"`
	// IpAddress: use this to search by IP address
	IpAddress *string `json:"-"`

	OrganizationId *string `json:"-"`
}

// ListIPs: list IPs
func (s *Api) ListIPs(req *ListIPsRequest, opts ...scw.RequestOption) (*ListIpsResponse, error) {
	var err error

	val := s.client.GetDefaultOrganizationID()
	if req.OrganizationId == nil || *req.OrganizationId == "" {
		req.OrganizationId = &val
	}

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "ip_address", req.IpAddress)
	parameter.AddToQuery(query, "organization_id", req.OrganizationId)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListIpsResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetIpRequest struct {
	Region utils.Region `json:"-"`
	// IpId:
	//
	// IP address ID
	IpId string `json:"-"`
}

// GetIp: get IP
func (s *Api) GetIp(req *GetIpRequest, opts ...scw.RequestOption) (*Ip, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IpId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Ip
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ReleaseIpRequest struct {
	Region utils.Region `json:"-"`
	// IpId: iP address ID
	IpId string `json:"-"`
}

// ReleaseIp: release IP
func (s *Api) ReleaseIp(req *ReleaseIpRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IpId) + "",
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq, opts...)

	if err != nil {
		return err
	}
	return nil
}

type UpdateIpRequest struct {
	Region utils.Region `json:"-"`
	// IpId: iP address ID
	IpId string `json:"-"`
	// Reverse: reverse DNS
	Reverse *string `json:"-"`
}

func (s *Api) UpdateIp(req *UpdateIpRequest, opts ...scw.RequestOption) (*Ip, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	query := url.Values{}
	parameter.AddToQuery(query, "reverse", req.Reverse)

	scwReq := &scw.ScalewayRequest{
		Method:  "PATCH",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/ips/" + fmt.Sprint(req.IpId) + "",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Ip
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListBackendsRequest struct {
	Region utils.Region `json:"-"`
	// LbId: load Balancer ID
	LbId string `json:"-"`
	// Name: use this to search by name
	Name *string `json:"-"`
	// OrderBy: choose order of response
	OrderBy ListBackendsRequestOrderBy `json:"-"`
	// Page: page number
	Page *int32 `json:"-"`
	// PageSize: set the maximum list sizes
	PageSize *int32 `json:"-"`
}

func (s *Api) ListBackends(req *ListBackendsRequest, opts ...scw.RequestOption) (*ListBackendsResponse, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	query := url.Values{}
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbId) + "/backends",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListBackendsResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateBackendRequest struct {
	Region utils.Region `json:"-"`
	// LbId: load Balancer ID
	LbId string `json:"-"`
	// Name: resource name
	Name string `json:"name,omitempty"`
	// ForwardProtocol: backend protocol. TCP or HTTP
	ForwardProtocol Protocol `json:"forward_protocol,omitempty"`
	// ForwardPort: user sessions will be forwarded to this port of backend servers
	ForwardPort int32 `json:"forward_port,omitempty"`
	// ForwardPortAlgorithm: load balancing algorithm
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm,omitempty"`
	// StickySessions: enables cookie-based session persistence
	StickySessions StickySessionsType `json:"sticky_sessions,omitempty"`
	// StickySessionsCookieName: cookie name for for sticky sessions
	StickySessionsCookieName string `json:"sticky_sessions_cookie_name,omitempty"`
	// HealthCheck: see the Healthcheck object description
	HealthCheck *HealthCheck `json:"health_check,omitempty"`
	// ServerIp: backend server IP addresses list (IPv4 or IPv6)
	ServerIp []string `json:"server_ip,omitempty"`
	// SendProxyV2: enables PROXY protocol version 2 (must be supported by backend servers)
	SendProxyV2 bool `json:"send_proxy_v2,omitempty"`
	// TimeoutServer: maximum server connection inactivity time
	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`
	// TimeoutConnect: maximum initical server connection establishment time
	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`
	// TimeoutTunnel: maximum tunnel inactivity time
	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`
	// OnMarkedDownAction: modify what occurs when a backend server is marked down
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

func (s *Api) CreateBackend(req *CreateBackendRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbId) + "/backends",
		Headers: http.Header{},
	}
	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Backend
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetBackendRequest struct {
	Region utils.Region `json:"-"`
	// BackendId: backend ID
	BackendId string `json:"-"`
}

func (s *Api) GetBackend(req *GetBackendRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Backend
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateBackendRequest struct {
	Region utils.Region `json:"-"`
	// BackendId: backend ID to update
	BackendId string `json:"-"`
	// Name: resource name
	Name string `json:"name,omitempty"`
	// ForwardProtocol: backend protocol. TCP or HTTP
	ForwardProtocol Protocol `json:"forward_protocol,omitempty"`
	// ForwardPort: user sessions will be forwarded to this port of backend servers
	ForwardPort int32 `json:"forward_port,omitempty"`
	// ForwardPortAlgorithm: load balancing algorithm
	ForwardPortAlgorithm ForwardPortAlgorithm `json:"forward_port_algorithm,omitempty"`
	// StickySessions: enable cookie-based session persistence
	StickySessions StickySessionsType `json:"sticky_sessions,omitempty"`
	// StickySessionsCookieName: cookie name for for sticky sessions
	StickySessionsCookieName string `json:"sticky_sessions_cookie_name,omitempty"`
	// SendProxyV2: enables PROXY protocol version 2 (must be supported by backend servers)
	SendProxyV2 bool `json:"send_proxy_v2,omitempty"`
	// TimeoutServer: maximum server connection inactivity time
	TimeoutServer *time.Duration `json:"timeout_server,omitempty"`
	// TimeoutConnect: maximum initial server connection establishment time
	TimeoutConnect *time.Duration `json:"timeout_connect,omitempty"`
	// TimeoutTunnel: maximum tunnel inactivity time
	TimeoutTunnel *time.Duration `json:"timeout_tunnel,omitempty"`
	// OnMarkedDownAction: modify what occurs when a backend server is marked down
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

func (s *Api) UpdateBackend(req *UpdateBackendRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendId) + "",
		Headers: http.Header{},
	}
	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Backend
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteBackendRequest struct {
	Region utils.Region `json:"-"`
	// BackendId: iD of the backend to delete
	BackendId string `json:"-"`
}

func (s *Api) DeleteBackend(req *DeleteBackendRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendId) + "",
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq, opts...)

	if err != nil {
		return err
	}
	return nil
}

type AddBackendServersRequest struct {
	Region utils.Region `json:"-"`
	// BackendId: backend ID
	BackendId string `json:"-"`
	// ServerIp: set all IPs to remove of your backend
	ServerIp []string `json:"server_ip,omitempty"`
}

func (s *Api) AddBackendServers(req *AddBackendServersRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendId) + "/servers",
		Headers: http.Header{},
	}
	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Backend
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type RemoveBackendServersRequest struct {
	Region utils.Region `json:"-"`
	// BackendId: backend ID
	BackendId string `json:"-"`
	// ServerIp: set all IPs to remove of your backend
	ServerIp []string `json:"server_ip,omitempty"`
}

func (s *Api) RemoveBackendServers(req *RemoveBackendServersRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendId) + "/servers",
		Headers: http.Header{},
	}
	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Backend
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type SetBackendServersRequest struct {
	Region utils.Region `json:"-"`
	// BackendId: backend ID
	BackendId string `json:"-"`
	// ServerIp: set all IPs to add of your backend and remove all other
	ServerIp []string `json:"server_ip,omitempty"`
}

func (s *Api) SetBackendServers(req *SetBackendServersRequest, opts ...scw.RequestOption) (*Backend, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendId) + "/servers",
		Headers: http.Header{},
	}
	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Backend
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateHealthCheckRequest struct {
	Region utils.Region `json:"-"`
	// BackendId: backend ID
	BackendId string `json:"-"`
	// Port: specify the port used to health check
	Port int32 `json:"port,omitempty"`
	// CheckDelay: time between two consecutive health checks
	CheckDelay *time.Duration `json:"check_delay,omitempty"`
	// CheckTimeout: additional check timeout, after the connection has been already established
	CheckTimeout *time.Duration `json:"check_timeout,omitempty"`
	// CheckMaxRetries: number of consecutive unsuccessful health checks, after wich the server will be considered dead
	CheckMaxRetries int32 `json:"check_max_retries,omitempty"`
	// MysqlConfig: the check requires MySQL >=3.22, for older version, please use TCP check
	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	MysqlConfig *HealthCheckMysqlConfig `json:"mysql_config,omitempty"`
	// LdapConfig: the response is analyzed to find an LDAPv3 response message
	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	LdapConfig *HealthCheckLdapConfig `json:"ldap_config,omitempty"`
	// RedisConfig: the response is analyzed to find the +PONG response message
	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	RedisConfig *HealthCheckRedisConfig `json:"redis_config,omitempty"`

	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	PgsqlConfig *HealthCheckPgsqlConfig `json:"pgsql_config,omitempty"`

	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	TcpConfig *HealthCheckTcpConfig `json:"tcp_config,omitempty"`

	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	HttpConfig *HealthCheckHttpConfig `json:"http_config,omitempty"`

	// Precisely one of HttpConfig, HttpsConfig, LdapConfig, MysqlConfig, PgsqlConfig, RedisConfig, TcpConfig must be set.
	HttpsConfig *HealthCheckHttpsConfig `json:"https_config,omitempty"`
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
	case m.TcpConfig != nil:
		return ConfigTcpConfig{*m.TcpConfig}
	case m.HttpConfig != nil:
		return ConfigHttpConfig{*m.HttpConfig}
	case m.HttpsConfig != nil:
		return ConfigHttpsConfig{*m.HttpsConfig}
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

func (s *Api) UpdateHealthCheck(req *UpdateHealthCheckRequest, opts ...scw.RequestOption) (*HealthCheck, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/backends/" + fmt.Sprint(req.BackendId) + "/healthcheck",
		Headers: http.Header{},
	}
	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp HealthCheck
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListFrontendsRequest struct {
	Region utils.Region `json:"-"`
	// LbId: load Balancer ID
	LbId string `json:"-"`
	// Name: use this to search by name
	Name *string `json:"-"`
	// OrderBy: response order
	OrderBy ListFrontendsRequestOrderBy `json:"-"`
	// Page: page number
	Page *int32 `json:"-"`
	// PageSize: set the maximum list sizes
	PageSize *int32 `json:"-"`
}

func (s *Api) ListFrontends(req *ListFrontendsRequest, opts ...scw.RequestOption) (*ListFrontendsResponse, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	query := url.Values{}
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbId) + "/frontends",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListFrontendsResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateFrontendRequest struct {
	Region utils.Region `json:"-"`
	// LbId: load Balancer ID
	LbId string `json:"-"`
	// Name: resource name
	Name string `json:"name,omitempty"`
	// InboundPort: tCP port to listen on the front side
	InboundPort int32 `json:"inbound_port,omitempty"`
	// BackendId: backend ID
	BackendId string `json:"backend_id,omitempty"`
	// TimeoutClient: set the maximum inactivity time on the client side
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

func (s *Api) CreateFrontend(req *CreateFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbId) + "/frontends",
		Headers: http.Header{},
	}
	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Frontend
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetFrontendRequest struct {
	Region utils.Region `json:"-"`
	// FrontendId: frontend ID
	FrontendId string `json:"-"`
}

func (s *Api) GetFrontend(req *GetFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Frontend
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateFrontendRequest struct {
	Region utils.Region `json:"-"`
	// FrontendId: frontend ID
	FrontendId string `json:"-"`
	// Name: resource name
	Name string `json:"name,omitempty"`
	// InboundPort: tCP port to listen on the front side
	InboundPort int32 `json:"inbound_port,omitempty"`
	// BackendId: backend ID
	BackendId string `json:"backend_id,omitempty"`
	// TimeoutClient: client session maximum inactivity time
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

func (s *Api) UpdateFrontend(req *UpdateFrontendRequest, opts ...scw.RequestOption) (*Frontend, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendId) + "",
		Headers: http.Header{},
	}
	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Frontend
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteFrontendRequest struct {
	Region utils.Region `json:"-"`
	// FrontendId: frontend ID to delete
	FrontendId string `json:"-"`
}

func (s *Api) DeleteFrontend(req *DeleteFrontendRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendId) + "",
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq, opts...)

	if err != nil {
		return err
	}
	return nil
}

type GetLbStatsRequest struct {
	Region utils.Region `json:"-"`
	// LbId: load Balancer ID
	LbId string `json:"-"`
}

func (s *Api) GetLbStats(req *GetLbStatsRequest, opts ...scw.RequestOption) (*LbStats, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/lbs/" + fmt.Sprint(req.LbId) + "/stats",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp LbStats
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type ListAclsRequest struct {
	Region utils.Region `json:"-"`
	// FrontendId: iD of your frontend
	FrontendId string `json:"-"`
	// OrderBy: you can order the response by created_at asc/desc or name asc/desc
	OrderBy ListAclRequestOrderBy `json:"-"`
	// Page: page number
	Page *int32 `json:"-"`
	// PageSize: set the maximum list size
	PageSize *int32 `json:"-"`
	// Name: filter acl per name
	Name *string `json:"-"`
}

func (s *Api) ListAcls(req *ListAclsRequest, opts ...scw.RequestOption) (*ListAclResponse, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendId) + "/acls",
		Query:   query,
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp ListAclResponse
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type CreateAclRequest struct {
	Region utils.Region `json:"-"`
	// FrontendId: iD of your frontend
	FrontendId string `json:"-"`
	// Name: name of your ACL ressource
	Name string `json:"name,omitempty"`
	// Action: see the AclAction object description
	Action *AclAction `json:"action,omitempty"`
	// Match: see the AclMatch object description
	Match *AclMatch `json:"match,omitempty"`
	// Index: order between your Acls (ascending order, 0 is first acl executed)
	Index int32 `json:"index,omitempty"`
}

func (s *Api) CreateAcl(req *CreateAclRequest, opts ...scw.RequestOption) (*Acl, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "POST",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/frontends/" + fmt.Sprint(req.FrontendId) + "/acls",
		Headers: http.Header{},
	}
	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Acl
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type GetAclRequest struct {
	Region utils.Region `json:"-"`
	// AclId: iD of your ACL ressource
	AclId string `json:"-"`
}

func (s *Api) GetAcl(req *GetAclRequest, opts ...scw.RequestOption) (*Acl, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "GET",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/acls/" + fmt.Sprint(req.AclId) + "",
		Headers: http.Header{},
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Acl
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type UpdateAclRequest struct {
	Region utils.Region `json:"-"`
	// AclId: iD of your ACL ressource
	AclId string `json:"-"`
	// Name: name of your ACL ressource
	Name string `json:"name,omitempty"`
	// Action: see the AclAction object description
	Action *AclAction `json:"action,omitempty"`
	// Match: see the AclMatch object description
	Match *AclMatch `json:"match,omitempty"`
	// Index: order between your Acls (ascending order, 0 is first acl executed)
	Index int32 `json:"index,omitempty"`
}

func (s *Api) UpdateAcl(req *UpdateAclRequest, opts ...scw.RequestOption) (*Acl, error) {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "PUT",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/acls/" + fmt.Sprint(req.AclId) + "",
		Headers: http.Header{},
	}
	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	scwResp, err := s.client.Do(scwReq, opts...)

	if err != nil {
		return nil, err
	}
	defer scwResp.Body.Close()
	var resp Acl
	err = json.NewDecoder(scwResp.Body).Decode(&resp)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

type DeleteAclRequest struct {
	Region utils.Region `json:"-"`
	// AclId: iD of your ACL ressource
	AclId string `json:"-"`
}

func (s *Api) DeleteAcl(req *DeleteAclRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		req.Region = s.client.GetDefaultRegion()
	}

	scwReq := &scw.ScalewayRequest{
		Method:  "DELETE",
		Path:    "/lb/v1/regions/" + fmt.Sprint(req.Region) + "/acls/" + fmt.Sprint(req.AclId) + "",
		Headers: http.Header{},
	}

	_, err = s.client.Do(scwReq, opts...)

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

type ConfigTcpConfig struct {
	Value HealthCheckTcpConfig
}

func (ConfigTcpConfig) isConfig() {
}

type ConfigHttpConfig struct {
	Value HealthCheckHttpConfig
}

func (ConfigHttpConfig) isConfig() {
}

type ConfigHttpsConfig struct {
	Value HealthCheckHttpsConfig
}

func (ConfigHttpsConfig) isConfig() {
}

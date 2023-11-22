// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package parpaing_private provides methods and message types of the parpaing_private v1alpha1 API.
package parpaing_private

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

type ListArticlesRequestOrderBy string

const (
	ListArticlesRequestOrderByCreatedAtAsc  = ListArticlesRequestOrderBy("created_at_asc")
	ListArticlesRequestOrderByCreatedAtDesc = ListArticlesRequestOrderBy("created_at_desc")
	ListArticlesRequestOrderByUpdatedAtAsc  = ListArticlesRequestOrderBy("updated_at_asc")
	ListArticlesRequestOrderByUpdatedAtDesc = ListArticlesRequestOrderBy("updated_at_desc")
	ListArticlesRequestOrderByNameAsc       = ListArticlesRequestOrderBy("name_asc")
	ListArticlesRequestOrderByNameDesc      = ListArticlesRequestOrderBy("name_desc")
	ListArticlesRequestOrderByStockAsc      = ListArticlesRequestOrderBy("stock_asc")
	ListArticlesRequestOrderByStockDesc     = ListArticlesRequestOrderBy("stock_desc")
	ListArticlesRequestOrderByPriceAsc      = ListArticlesRequestOrderBy("price_asc")
	ListArticlesRequestOrderByPriceDesc     = ListArticlesRequestOrderBy("price_desc")
)

func (enum ListArticlesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListArticlesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListArticlesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListArticlesRequestOrderBy(ListArticlesRequestOrderBy(tmp).String())
	return nil
}

type ListOrdersRequestOrderBy string

const (
	ListOrdersRequestOrderByOrderedAtAsc  = ListOrdersRequestOrderBy("ordered_at_asc")
	ListOrdersRequestOrderByOrderedAtDesc = ListOrdersRequestOrderBy("ordered_at_desc")
)

func (enum ListOrdersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "ordered_at_asc"
	}
	return string(enum)
}

func (enum ListOrdersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOrdersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOrdersRequestOrderBy(ListOrdersRequestOrderBy(tmp).String())
	return nil
}

type ListTransactionsRequestOrderBy string

const (
	ListTransactionsRequestOrderByCreatedAtAsc  = ListTransactionsRequestOrderBy("created_at_asc")
	ListTransactionsRequestOrderByCreatedAtDesc = ListTransactionsRequestOrderBy("created_at_desc")
	ListTransactionsRequestOrderByAmountAsc     = ListTransactionsRequestOrderBy("amount_asc")
	ListTransactionsRequestOrderByAmountDesc    = ListTransactionsRequestOrderBy("amount_desc")
)

func (enum ListTransactionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListTransactionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListTransactionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListTransactionsRequestOrderBy(ListTransactionsRequestOrderBy(tmp).String())
	return nil
}

type ListUsersRequestOrderBy string

const (
	ListUsersRequestOrderByCreatedAtAsc  = ListUsersRequestOrderBy("created_at_asc")
	ListUsersRequestOrderByCreatedAtDesc = ListUsersRequestOrderBy("created_at_desc")
	ListUsersRequestOrderByThuneAsc      = ListUsersRequestOrderBy("thune_asc")
	ListUsersRequestOrderByThuneDesc     = ListUsersRequestOrderBy("thune_desc")
	ListUsersRequestOrderByNameAsc       = ListUsersRequestOrderBy("name_asc")
	ListUsersRequestOrderByNameDesc      = ListUsersRequestOrderBy("name_desc")
)

func (enum ListUsersRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListUsersRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListUsersRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListUsersRequestOrderBy(ListUsersRequestOrderBy(tmp).String())
	return nil
}

type OrderStatus string

const (
	OrderStatusUnknownOrderStatus = OrderStatus("unknown_order_status")
	OrderStatusOrdered            = OrderStatus("ordered")
	OrderStatusProcessing         = OrderStatus("processing")
	OrderStatusReady              = OrderStatus("ready")
	OrderStatusDelivered          = OrderStatus("delivered")
	OrderStatusCancelled          = OrderStatus("cancelled")
)

func (enum OrderStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_order_status"
	}
	return string(enum)
}

func (enum OrderStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrderStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrderStatus(OrderStatus(tmp).String())
	return nil
}

type TransactionReceiverReceiverType string

const (
	TransactionReceiverReceiverTypeUnknownActorType = TransactionReceiverReceiverType("unknown_actor_type")
	TransactionReceiverReceiverTypeUser             = TransactionReceiverReceiverType("user")
	TransactionReceiverReceiverTypeShop             = TransactionReceiverReceiverType("shop")
)

func (enum TransactionReceiverReceiverType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_actor_type"
	}
	return string(enum)
}

func (enum TransactionReceiverReceiverType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TransactionReceiverReceiverType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TransactionReceiverReceiverType(TransactionReceiverReceiverType(tmp).String())
	return nil
}

type TransactionSenderSenderType string

const (
	TransactionSenderSenderTypeUnknownActorType = TransactionSenderSenderType("unknown_actor_type")
	TransactionSenderSenderTypeUser             = TransactionSenderSenderType("user")
	TransactionSenderSenderTypeBank             = TransactionSenderSenderType("bank")
	TransactionSenderSenderTypeShop             = TransactionSenderSenderType("shop")
)

func (enum TransactionSenderSenderType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_actor_type"
	}
	return string(enum)
}

func (enum TransactionSenderSenderType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *TransactionSenderSenderType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = TransactionSenderSenderType(TransactionSenderSenderType(tmp).String())
	return nil
}

type UnitStoreType string

const (
	UnitStoreTypeUnknownUnitStoreType = UnitStoreType("unknown_unit_store_type")
	UnitStoreTypeReceived             = UnitStoreType("received")
	UnitStoreTypeGivable              = UnitStoreType("givable")
)

func (enum UnitStoreType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_unit_store_type"
	}
	return string(enum)
}

func (enum UnitStoreType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UnitStoreType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UnitStoreType(UnitStoreType(tmp).String())
	return nil
}

type UserRole string

const (
	UserRoleUnknownRole = UserRole("unknown_role")
	UserRoleBasic       = UserRole("basic")
	UserRoleAdmin       = UserRole("admin")
)

func (enum UserRole) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_role"
	}
	return string(enum)
}

func (enum UserRole) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UserRole) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UserRole(UserRole(tmp).String())
	return nil
}

// TransactionReceiver: transaction receiver.
type TransactionReceiver struct {
	// ReceiverType: default value: unknown_actor_type
	ReceiverType TransactionReceiverReceiverType `json:"receiver_type"`

	// UnitStoreType: default value: unknown_unit_store_type
	UnitStoreType UnitStoreType `json:"unit_store_type"`

	UserID *string `json:"user_id"`

	UserName *string `json:"user_name"`
}

// TransactionSender: transaction sender.
type TransactionSender struct {
	// SenderType: default value: unknown_actor_type
	SenderType TransactionSenderSenderType `json:"sender_type"`

	// UnitStoreType: default value: unknown_unit_store_type
	UnitStoreType UnitStoreType `json:"unit_store_type"`

	UserID *string `json:"user_id"`

	UserName *string `json:"user_name"`
}

// Article: article.
type Article struct {
	ID string `json:"id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	Name string `json:"name"`

	Description *string `json:"description"`

	ImageURL *string `json:"image_url"`

	Stock uint64 `json:"stock"`

	Price uint64 `json:"price"`

	Orderable *bool `json:"orderable"`

	Deleted *bool `json:"deleted"`
}

// Order: order.
type Order struct {
	ID string `json:"id"`

	OrderedAt *time.Time `json:"ordered_at"`

	ProcessingAt *time.Time `json:"processing_at"`

	ReadyAt *time.Time `json:"ready_at"`

	DeliveredAt *time.Time `json:"delivered_at"`

	CancelledAt *time.Time `json:"cancelled_at"`

	// Status: default value: unknown_order_status
	Status OrderStatus `json:"status"`

	UserID string `json:"user_id"`

	ArticleID string `json:"article_id"`

	TransactionID string `json:"transaction_id"`

	CancelTransactionID *string `json:"cancel_transaction_id"`

	UserName string `json:"user_name"`

	ArticleName string `json:"article_name"`
}

// Transaction: transaction.
type Transaction struct {
	ID string `json:"id"`

	CreatedAt *time.Time `json:"created_at"`

	Sender *TransactionSender `json:"sender"`

	Receiver *TransactionReceiver `json:"receiver"`

	Amount uint64 `json:"amount"`

	SenderMessage *string `json:"sender_message"`

	SenderEmoji *string `json:"sender_emoji"`

	ReceiverReaction *string `json:"receiver_reaction"`

	ReceiverReactedAt *time.Time `json:"receiver_reacted_at"`
}

// User: user.
type User struct {
	ID string `json:"id"`

	Name string `json:"name"`

	ReceivedBalance uint64 `json:"received_balance"`

	GivableBalance uint64 `json:"givable_balance"`

	// Role: default value: unknown_role
	Role UserRole `json:"role"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	DeletedAt *time.Time `json:"deleted_at"`
}

// CancelOrderRequest: cancel order request.
type CancelOrderRequest struct {
	OrderID string `json:"-"`
}

// CreateArticleRequest: create article request.
type CreateArticleRequest struct {
	Name string `json:"name"`

	Description string `json:"description"`

	Stock uint64 `json:"stock"`

	Price uint64 `json:"price"`

	Orderable *bool `json:"orderable,omitempty"`

	ImageURL *string `json:"image_url,omitempty"`
}

// CreateOrderRequest: create order request.
type CreateOrderRequest struct {
	ArticleID string `json:"article_id"`

	UserID string `json:"user_id"`
}

// DeleteArticleRequest: delete article request.
type DeleteArticleRequest struct {
	ArticleID string `json:"-"`
}

// DeleteTokenRequest: delete token request.
type DeleteTokenRequest struct {
	ID string `json:"-"`
}

// DeleteUserRequest: delete user request.
type DeleteUserRequest struct {
	UserID string `json:"-"`
}

// GetUserRequest: get user request.
type GetUserRequest struct {
	UserID string `json:"-"`
}

// GetUserStatsRequest: get user stats request.
type GetUserStatsRequest struct {
	UserID string `json:"-"`
}

// ListArticlesRequest: list articles request.
type ListArticlesRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListArticlesRequestOrderBy `json:"-"`

	Orderable *bool `json:"-"`

	Deleted *bool `json:"-"`
}

// ListArticlesResponse: list articles response.
type ListArticlesResponse struct {
	TotalCount uint64 `json:"total_count"`

	Articles []*Article `json:"articles"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListArticlesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListArticlesResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListArticlesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Articles = append(r.Articles, results.Articles...)
	r.TotalCount += uint64(len(results.Articles))
	return uint64(len(results.Articles)), nil
}

// ListOrdersRequest: list orders request.
type ListOrdersRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: ordered_at_asc
	OrderBy ListOrdersRequestOrderBy `json:"-"`

	// Status: default value: unknown_order_status
	Status *OrderStatus `json:"-"`

	UserIDs []string `json:"-"`

	ArticleIDs []string `json:"-"`
}

// ListOrdersResponse: list orders response.
type ListOrdersResponse struct {
	TotalCount uint64 `json:"total_count"`

	Orders []*Order `json:"orders"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOrdersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOrdersResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListOrdersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Orders = append(r.Orders, results.Orders...)
	r.TotalCount += uint64(len(results.Orders))
	return uint64(len(results.Orders)), nil
}

// ListTransactionsRequest: list transactions request.
type ListTransactionsRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListTransactionsRequestOrderBy `json:"-"`

	SenderIDs []string `json:"-"`

	ReceiverIDs []string `json:"-"`
}

// ListTransactionsResponse: list transactions response.
type ListTransactionsResponse struct {
	TotalCount uint64 `json:"total_count"`

	Transactions []*Transaction `json:"transactions"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListTransactionsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListTransactionsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListTransactionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Transactions = append(r.Transactions, results.Transactions...)
	r.TotalCount += uint64(len(results.Transactions))
	return uint64(len(results.Transactions)), nil
}

// ListUsersRequest: list users request.
type ListUsersRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListUsersRequestOrderBy `json:"-"`

	UserIDs []string `json:"-"`
}

// ListUsersResponse: list users response.
type ListUsersResponse struct {
	TotalCount uint64 `json:"total_count"`

	Users []*User `json:"users"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListUsersResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListUsersResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Users = append(r.Users, results.Users...)
	r.TotalCount += uint64(len(results.Users))
	return uint64(len(results.Users)), nil
}

// RenewTokenRequest: renew token request.
type RenewTokenRequest struct {
	ID string `json:"-"`

	RenewToken string `json:"renew_token"`
}

// SendParpaingRequest: send parpaing request.
type SendParpaingRequest struct {
	ToUserIDs []string `json:"to_user_ids"`

	Emoji string `json:"emoji"`

	Message string `json:"message"`

	Amount uint64 `json:"amount"`
}

// SendParpaingResponse: send parpaing response.
type SendParpaingResponse struct {
	SendID string `json:"send_id"`
}

// Token: token.
type Token struct {
	Token string `json:"token"`

	RenewToken string `json:"renew_token"`
}

// UpdateArticleRequest: update article request.
type UpdateArticleRequest struct {
	ArticleID string `json:"-"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`

	Stock *uint64 `json:"stock,omitempty"`

	Price *uint64 `json:"price,omitempty"`

	Orderable *bool `json:"orderable,omitempty"`

	ImageURL *string `json:"image_url,omitempty"`
}

// UpdateOrderRequest: update order request.
type UpdateOrderRequest struct {
	OrderID string `json:"-"`

	// Status: default value: unknown_order_status
	Status *OrderStatus `json:"status,omitempty"`
}

// UpdateUserRequest: update user request.
type UpdateUserRequest struct {
	UserID string `json:"-"`

	Name *string `json:"name,omitempty"`

	ReceivedBalance *uint64 `json:"received_balance,omitempty"`

	GivableBalance *uint64 `json:"givable_balance,omitempty"`

	// Role: default value: unknown_role
	Role *UserRole `json:"role,omitempty"`
}

// UploadArticleImageRequest: upload article image request.
type UploadArticleImageRequest struct {
	ArticleID string `json:"-"`

	ContentType string `json:"content_type"`
}

// UploadArticleImageResponse: upload article image response.
type UploadArticleImageResponse struct {
	URL string `json:"url"`

	Method string `json:"method"`

	Headers map[string]string `json:"headers"`
}

// UserStats: user stats.
type UserStats struct {
	WeekGaveParpaings uint64 `json:"week_gave_parpaings"`

	TotalGaveParpaings uint64 `json:"total_gave_parpaings"`

	ParpaingReceiversCount uint64 `json:"parpaing_receivers_count"`
}

type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ListUsers:
func (s *API) ListUsers(req *ListUsersRequest, opts ...scw.RequestOption) (*ListUsersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "user_ids", req.UserIDs)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/parpaing-private/v1alpha1/users",
		Query:  query,
	}

	var resp ListUsersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUser:
func (s *API) GetUser(req *GetUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/parpaing-private/v1alpha1/users/" + fmt.Sprint(req.UserID) + "",
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetUserStats:
func (s *API) GetUserStats(req *GetUserStatsRequest, opts ...scw.RequestOption) (*UserStats, error) {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/parpaing-private/v1alpha1/users/" + fmt.Sprint(req.UserID) + "/stats",
	}

	var resp UserStats

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateUser:
func (s *API) UpdateUser(req *UpdateUserRequest, opts ...scw.RequestOption) (*User, error) {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return nil, errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/parpaing-private/v1alpha1/users/" + fmt.Sprint(req.UserID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp User

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteUser:
func (s *API) DeleteUser(req *DeleteUserRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.UserID) == "" {
		return errors.New("field UserID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/parpaing-private/v1alpha1/users/" + fmt.Sprint(req.UserID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// SendParpaing:
func (s *API) SendParpaing(req *SendParpaingRequest, opts ...scw.RequestOption) (*SendParpaingResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/parpaing-private/v1alpha1/send-parpaing",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SendParpaingResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOrders:
func (s *API) ListOrders(req *ListOrdersRequest, opts ...scw.RequestOption) (*ListOrdersResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "user_ids", req.UserIDs)
	parameter.AddToQuery(query, "article_ids", req.ArticleIDs)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/parpaing-private/v1alpha1/orders",
		Query:  query,
	}

	var resp ListOrdersResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateOrder:
func (s *API) CreateOrder(req *CreateOrderRequest, opts ...scw.RequestOption) (*Order, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/parpaing-private/v1alpha1/orders",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Order

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOrder:
func (s *API) UpdateOrder(req *UpdateOrderRequest, opts ...scw.RequestOption) (*Order, error) {
	var err error

	if fmt.Sprint(req.OrderID) == "" {
		return nil, errors.New("field OrderID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/parpaing-private/v1alpha1/orders/" + fmt.Sprint(req.OrderID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Order

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CancelOrder:
func (s *API) CancelOrder(req *CancelOrderRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.OrderID) == "" {
		return errors.New("field OrderID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/parpaing-private/v1alpha1/orders/" + fmt.Sprint(req.OrderID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListArticles:
func (s *API) ListArticles(req *ListArticlesRequest, opts ...scw.RequestOption) (*ListArticlesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "orderable", req.Orderable)
	parameter.AddToQuery(query, "deleted", req.Deleted)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/parpaing-private/v1alpha1/articles",
		Query:  query,
	}

	var resp ListArticlesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateArticle:
func (s *API) CreateArticle(req *CreateArticleRequest, opts ...scw.RequestOption) (*Article, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/parpaing-private/v1alpha1/articles",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Article

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateArticle:
func (s *API) UpdateArticle(req *UpdateArticleRequest, opts ...scw.RequestOption) (*Article, error) {
	var err error

	if fmt.Sprint(req.ArticleID) == "" {
		return nil, errors.New("field ArticleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/parpaing-private/v1alpha1/articles/" + fmt.Sprint(req.ArticleID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Article

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UploadArticleImage:
func (s *API) UploadArticleImage(req *UploadArticleImageRequest, opts ...scw.RequestOption) (*UploadArticleImageResponse, error) {
	var err error

	if fmt.Sprint(req.ArticleID) == "" {
		return nil, errors.New("field ArticleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/parpaing-private/v1alpha1/articles/" + fmt.Sprint(req.ArticleID) + "/image",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp UploadArticleImageResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteArticle:
func (s *API) DeleteArticle(req *DeleteArticleRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.ArticleID) == "" {
		return errors.New("field ArticleID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/parpaing-private/v1alpha1/articles/" + fmt.Sprint(req.ArticleID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListTransactions:
func (s *API) ListTransactions(req *ListTransactionsRequest, opts ...scw.RequestOption) (*ListTransactionsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "sender_ids", req.SenderIDs)
	parameter.AddToQuery(query, "receiver_ids", req.ReceiverIDs)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/parpaing-private/v1alpha1/transactions",
		Query:  query,
	}

	var resp ListTransactionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RenewToken:
func (s *API) RenewToken(req *RenewTokenRequest, opts ...scw.RequestOption) (*Token, error) {
	var err error

	if fmt.Sprint(req.ID) == "" {
		return nil, errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/parpaing-private/v1alpha1/tokens/" + fmt.Sprint(req.ID) + "/renew",
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

// DeleteToken:
func (s *API) DeleteToken(req *DeleteTokenRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.ID) == "" {
		return errors.New("field ID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/parpaing-private/v1alpha1/tokens/" + fmt.Sprint(req.ID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package payment_admin provides methods and message types of the payment_admin v2alpha1 API.
package payment_admin

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

type AccountingLineCategory string

const (
	AccountingLineCategoryUnknownCategory = AccountingLineCategory("unknown_category")
	AccountingLineCategoryUnknown         = AccountingLineCategory("unknown")
	AccountingLineCategorySaleInvoice     = AccountingLineCategory("sale_invoice")
	AccountingLineCategorySalePayment     = AccountingLineCategory("sale_payment")
	AccountingLineCategorySaleRefund      = AccountingLineCategory("sale_refund")
	AccountingLineCategoryPurchaseInvoice = AccountingLineCategory("purchase_invoice")
	AccountingLineCategoryPurchasePayment = AccountingLineCategory("purchase_payment")
	AccountingLineCategoryPurchaseRefund  = AccountingLineCategory("purchase_refund")
)

func (enum AccountingLineCategory) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_category"
	}
	return string(enum)
}

func (enum AccountingLineCategory) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AccountingLineCategory) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AccountingLineCategory(AccountingLineCategory(tmp).String())
	return nil
}

type AccountingLinePaymentMode string

const (
	AccountingLinePaymentModeUnknownPaymentMode = AccountingLinePaymentMode("unknown_payment_mode")
	AccountingLinePaymentModeCard               = AccountingLinePaymentMode("card")
	AccountingLinePaymentModeSepa               = AccountingLinePaymentMode("sepa")
	AccountingLinePaymentModeCheck              = AccountingLinePaymentMode("check")
	AccountingLinePaymentModeWireTransfer       = AccountingLinePaymentMode("wire_transfer")
)

func (enum AccountingLinePaymentMode) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_payment_mode"
	}
	return string(enum)
}

func (enum AccountingLinePaymentMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AccountingLinePaymentMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AccountingLinePaymentMode(AccountingLinePaymentMode(tmp).String())
	return nil
}

type AccountingLineState string

const (
	AccountingLineStateUnknownState = AccountingLineState("unknown_state")
	AccountingLineStateDeleted      = AccountingLineState("deleted")
	AccountingLineStateDraft        = AccountingLineState("draft")
	AccountingLineStatePending      = AccountingLineState("pending")
	AccountingLineStateProcessing   = AccountingLineState("processing")
	AccountingLineStateClosed       = AccountingLineState("closed")
	AccountingLineStateLost         = AccountingLineState("lost")
)

func (enum AccountingLineState) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_state"
	}
	return string(enum)
}

func (enum AccountingLineState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AccountingLineState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AccountingLineState(AccountingLineState(tmp).String())
	return nil
}

type ImportedPaymentMethodPaymentMethod string

const (
	ImportedPaymentMethodPaymentMethodUnknown = ImportedPaymentMethodPaymentMethod("unknown")
	ImportedPaymentMethodPaymentMethodCard    = ImportedPaymentMethodPaymentMethod("card")
	ImportedPaymentMethodPaymentMethodSepa    = ImportedPaymentMethodPaymentMethod("sepa")
)

func (enum ImportedPaymentMethodPaymentMethod) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum ImportedPaymentMethodPaymentMethod) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ImportedPaymentMethodPaymentMethod) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ImportedPaymentMethodPaymentMethod(ImportedPaymentMethodPaymentMethod(tmp).String())
	return nil
}

type InvoiceState string

const (
	InvoiceStateUnknownState = InvoiceState("unknown_state")
	InvoiceStateDraft        = InvoiceState("draft")
	InvoiceStateOutdated     = InvoiceState("outdated")
	InvoiceStateStopped      = InvoiceState("stopped")
	InvoiceStateIncomplete   = InvoiceState("incomplete")
	InvoiceStateIssuing      = InvoiceState("issuing")
	InvoiceStateIssued       = InvoiceState("issued")
	InvoiceStateCancelled    = InvoiceState("cancelled")
)

func (enum InvoiceState) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_state"
	}
	return string(enum)
}

func (enum InvoiceState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InvoiceState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InvoiceState(InvoiceState(tmp).String())
	return nil
}

type InvoiceType string

const (
	InvoiceTypeUnknownType = InvoiceType("unknown_type")
	InvoiceTypePeriodic    = InvoiceType("periodic")
	InvoiceTypePurchase    = InvoiceType("purchase")
)

func (enum InvoiceType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_type"
	}
	return string(enum)
}

func (enum InvoiceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *InvoiceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = InvoiceType(InvoiceType(tmp).String())
	return nil
}

type ListMagicCodesRequestOrderBy string

const (
	ListMagicCodesRequestOrderByCreatedAtAsc  = ListMagicCodesRequestOrderBy("created_at_asc")
	ListMagicCodesRequestOrderByCreatedAtDesc = ListMagicCodesRequestOrderBy("created_at_desc")
	ListMagicCodesRequestOrderByUpdatedAtAsc  = ListMagicCodesRequestOrderBy("updated_at_asc")
	ListMagicCodesRequestOrderByUpdatedAtDesc = ListMagicCodesRequestOrderBy("updated_at_desc")
)

func (enum ListMagicCodesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListMagicCodesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListMagicCodesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListMagicCodesRequestOrderBy(ListMagicCodesRequestOrderBy(tmp).String())
	return nil
}

type MagicCodeStatus string

const (
	MagicCodeStatusUnknown    = MagicCodeStatus("unknown")
	MagicCodeStatusProcessing = MagicCodeStatus("processing")
	MagicCodeStatusPending    = MagicCodeStatus("pending")
	MagicCodeStatusValidated  = MagicCodeStatus("validated")
	MagicCodeStatusFailed     = MagicCodeStatus("failed")
	MagicCodeStatusExpired    = MagicCodeStatus("expired")
)

func (enum MagicCodeStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum MagicCodeStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *MagicCodeStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = MagicCodeStatus(MagicCodeStatus(tmp).String())
	return nil
}

type ProcessPaymentResponsePaymentStatus string

const (
	ProcessPaymentResponsePaymentStatusUnknownPaymentStatus = ProcessPaymentResponsePaymentStatus("unknown_payment_status")
	ProcessPaymentResponsePaymentStatusSuccess              = ProcessPaymentResponsePaymentStatus("success")
	ProcessPaymentResponsePaymentStatusFailure              = ProcessPaymentResponsePaymentStatus("failure")
)

func (enum ProcessPaymentResponsePaymentStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_payment_status"
	}
	return string(enum)
}

func (enum ProcessPaymentResponsePaymentStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ProcessPaymentResponsePaymentStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ProcessPaymentResponsePaymentStatus(ProcessPaymentResponsePaymentStatus(tmp).String())
	return nil
}

type UpdatePaymentModeRequestPaymentMode string

const (
	UpdatePaymentModeRequestPaymentModeUnknownPaymentMode = UpdatePaymentModeRequestPaymentMode("unknown_payment_mode")
	UpdatePaymentModeRequestPaymentModeCard               = UpdatePaymentModeRequestPaymentMode("card")
	UpdatePaymentModeRequestPaymentModeSepa               = UpdatePaymentModeRequestPaymentMode("sepa")
	UpdatePaymentModeRequestPaymentModeWireTransfer       = UpdatePaymentModeRequestPaymentMode("wire_transfer")
)

func (enum UpdatePaymentModeRequestPaymentMode) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_payment_mode"
	}
	return string(enum)
}

func (enum UpdatePaymentModeRequestPaymentMode) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *UpdatePaymentModeRequestPaymentMode) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = UpdatePaymentModeRequestPaymentMode(UpdatePaymentModeRequestPaymentMode(tmp).String())
	return nil
}

// AccountingLineOrganization: accounting line organization.
type AccountingLineOrganization struct {
	ID string `json:"id"`

	Name string `json:"name"`
}

// ImportedPaymentMethod: imported payment method.
type ImportedPaymentMethod struct {
	// PaymentMethod: payment method can be card or sepa.
	// Default value: unknown
	PaymentMethod ImportedPaymentMethodPaymentMethod `json:"payment_method"`

	// DeviceID: key of the payment method imported in our database.
	DeviceID string `json:"device_id"`
}

// MagicCode: magic code.
type MagicCode struct {
	ID string `json:"id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	ExpiresAt *time.Time `json:"expires_at"`

	// Status: default value: unknown
	Status MagicCodeStatus `json:"status"`

	CardID string `json:"card_id"`

	Code string `json:"code"`

	OrganizationID string `json:"organization_id"`

	TryCount uint32 `json:"try_count"`
}

// Invoice: invoice.
type Invoice struct {
	// ID: the ID of the invoice.
	ID string `json:"id"`

	// Type: the type of invoice.
	// Default value: unknown_type
	Type InvoiceType `json:"type"`

	// State: the state of the invoice.
	// Default value: unknown_state
	State InvoiceState `json:"state"`

	// OrganizationID: the ID of the organization associated to that invoice.
	OrganizationID string `json:"organization_id"`

	// StartDate: the invoice start date.
	StartDate *time.Time `json:"start_date"`

	// StopDate: the invoice stop date.
	StopDate *time.Time `json:"stop_date"`

	// DueDate: the invoice due date.
	DueDate *time.Time `json:"due_date"`

	// Number: the invoice number.
	Number *int32 `json:"number"`

	// TotalTaxed: the total amount of the invoice including taxes.
	TotalTaxed *scw.Money `json:"total_taxed"`

	// Currency: the currency of the invoice.
	Currency string `json:"currency"`
}

// AccountingLine: accounting line.
type AccountingLine struct {
	// ID: the ID of the accounting line.
	ID string `json:"id"`

	// CreationDate: the creation date of the accounting line.
	CreationDate *time.Time `json:"creation_date"`

	// ModificationDate: the modification date of the accounting line.
	ModificationDate *time.Time `json:"modification_date"`

	// Organization: the special default organization of the accounting line.
	Organization *AccountingLineOrganization `json:"organization"`

	// ThirdParty: the third party organization of the accounting line.
	ThirdParty *AccountingLineOrganization `json:"third_party"`

	// State: the state of the accounting line.
	// Default value: unknown_state
	State AccountingLineState `json:"state"`

	// Description: the description of the accounting line.
	Description string `json:"description"`

	// Amount: the amount of the accounting line.
	Amount *scw.Money `json:"amount"`

	// Category: the category of the accounting line.
	// Default value: unknown_category
	Category AccountingLineCategory `json:"category"`

	// PaymentMode: the payment mode of the accounting line.
	// Default value: unknown_payment_mode
	PaymentMode AccountingLinePaymentMode `json:"payment_mode"`

	// TryCount: the try count of the accounting line.
	TryCount uint32 `json:"try_count"`

	// InvoiceNumber: the invoice number of the accounting line.
	InvoiceNumber int32 `json:"invoice_number"`

	// InvoiceStartDate: the invoice start date of the accounting line.
	InvoiceStartDate *time.Time `json:"invoice_start_date"`

	// InvoiceStopDate: the invoice stop date of the accounting line.
	InvoiceStopDate *time.Time `json:"invoice_stop_date"`

	// EventSentDate: the event sent date of the accounting line.
	EventSentDate *time.Time `json:"event_sent_date"`

	// EventResponseDate: the event response date of the accounting line.
	EventResponseDate *time.Time `json:"event_response_date"`
}

// GetAccountingLineRequest: get accounting line request.
type GetAccountingLineRequest struct {
	// AccountingLineID: the ID of the accounting line.
	AccountingLineID string `json:"-"`
}

// ImportPaymentMethodsRequest: import payment methods request.
type ImportPaymentMethodsRequest struct {
	// OrganizationID: the elements organization ID to migrate.
	OrganizationID string `json:"organization_id"`
}

// ImportPaymentMethodsResponse: import payment methods response.
type ImportPaymentMethodsResponse struct {
	// ImportedPaymentMethod: an array of all payment methods imported.
	ImportedPaymentMethod []*ImportedPaymentMethod `json:"imported_payment_method"`
}

// ListMagicCodesRequest: list magic codes request.
type ListMagicCodesRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListMagicCodesRequestOrderBy `json:"-"`

	CardID *string `json:"-"`

	// Deprecated: Status: default value: unknown
	Status *MagicCodeStatus `json:"-"`

	OrganizationID *string `json:"-"`

	CreatedAfter *time.Time `json:"-"`

	ExcludeStatuses []MagicCodeStatus `json:"-"`

	Statuses []MagicCodeStatus `json:"-"`
}

// ListMagicCodesResponse: list magic codes response.
type ListMagicCodesResponse struct {
	TotalCount uint32 `json:"total_count"`

	MagicCodes []*MagicCode `json:"magic_codes"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListMagicCodesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListMagicCodesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListMagicCodesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.MagicCodes = append(r.MagicCodes, results.MagicCodes...)
	r.TotalCount += uint32(len(results.MagicCodes))
	return uint32(len(results.MagicCodes)), nil
}

// ProcessOnlineRefundRequest: process online refund request.
type ProcessOnlineRefundRequest struct {
	// OrganizationID: the refunded organization ID.
	OrganizationID string `json:"organization_id"`

	// TransactionID: online transaction ID to refund.
	TransactionID uint32 `json:"transaction_id"`

	// Amount: the amount that should be refunded.
	Amount *scw.Money `json:"amount,omitempty"`
}

// ProcessPaymentRequest: process payment request.
type ProcessPaymentRequest struct {
	// Amount: amount to debit.
	Amount *scw.Money `json:"amount,omitempty"`

	// OrganizationID: organization ID to debit.
	OrganizationID string `json:"organization_id"`

	// TransactionID: represent the transaction_key/invoice_key generated on the Online side.
	TransactionID uint32 `json:"transaction_id"`
}

// ProcessPaymentResponse: process payment response.
type ProcessPaymentResponse struct {
	// PaymentMode: the payment mode used to process payment.
	PaymentMode string `json:"payment_mode"`

	// PaymentStatus: the processed payment status.
	// Default value: unknown_payment_status
	PaymentStatus ProcessPaymentResponsePaymentStatus `json:"payment_status"`
}

// RefundMagicCodeRequest: refund magic code request.
type RefundMagicCodeRequest struct {
	MagicCodeID string `json:"-"`
}

// SendMagicCodeRequest: send magic code request.
type SendMagicCodeRequest struct {
	CardID string `json:"card_id"`
}

// SendOnlineNotificationRequest: send online notification request.
type SendOnlineNotificationRequest struct {
	// CardNotificationID: the card notification Id defining the notification to resend.
	CardNotificationID string `json:"card_notification_id"`
}

// SendOnlineNotificationResponse: send online notification response.
type SendOnlineNotificationResponse struct {
	// StatusCode: the status code of the Online endpoint that received the notification.
	StatusCode uint32 `json:"status_code"`
}

// SyncAccountingLineWithInvoiceRequest: sync accounting line with invoice request.
type SyncAccountingLineWithInvoiceRequest struct {
	// Invoice: the invoice to sync the accounting line with.
	Invoice *Invoice `json:"invoice,omitempty"`
}

// UpdateAccountingLineRequest: update accounting line request.
type UpdateAccountingLineRequest struct {
	// AccountingLineID: the ID of the accounting line to update.
	AccountingLineID string `json:"-"`

	// State: the new state of the accounting line.
	// Default value: unknown_state
	State AccountingLineState `json:"state"`
}

// UpdateMagicCodeRequest: update magic code request.
type UpdateMagicCodeRequest struct {
	MagicCodeID string `json:"-"`

	ExpiresAt *time.Time `json:"expires_at,omitempty"`
}

// UpdatePaymentModeRequest: update payment mode request.
type UpdatePaymentModeRequest struct {
	// OrganizationID: the organization_id to update.
	OrganizationID string `json:"-"`

	// Force: it enables to force the payment mode switch without specific checks.
	Force bool `json:"force"`

	// PaymentMode: update to this payment_mode. It can be card, sepa or wire_transfer.
	// Default value: unknown_payment_mode
	PaymentMode UpdatePaymentModeRequestPaymentMode `json:"payment_mode"`
}

// ValidateMagicCodeRequest: validate magic code request.
type ValidateMagicCodeRequest struct {
	MagicCodeID string `json:"-"`
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

// SendMagicCode:
func (s *API) SendMagicCode(req *SendMagicCodeRequest, opts ...scw.RequestOption) (*MagicCode, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/payment-admin/v2alpha1/magic-codes",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp MagicCode

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateMagicCode:
func (s *API) ValidateMagicCode(req *ValidateMagicCodeRequest, opts ...scw.RequestOption) (*MagicCode, error) {
	var err error

	if fmt.Sprint(req.MagicCodeID) == "" {
		return nil, errors.New("field MagicCodeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/payment-admin/v2alpha1/magic-codes/" + fmt.Sprint(req.MagicCodeID) + "/validate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp MagicCode

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateMagicCode:
func (s *API) UpdateMagicCode(req *UpdateMagicCodeRequest, opts ...scw.RequestOption) (*MagicCode, error) {
	var err error

	if fmt.Sprint(req.MagicCodeID) == "" {
		return nil, errors.New("field MagicCodeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/payment-admin/v2alpha1/magic-codes/" + fmt.Sprint(req.MagicCodeID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp MagicCode

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListMagicCodes:
func (s *API) ListMagicCodes(req *ListMagicCodesRequest, opts ...scw.RequestOption) (*ListMagicCodesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "card_id", req.CardID)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "created_after", req.CreatedAfter)
	parameter.AddToQuery(query, "exclude_statuses", req.ExcludeStatuses)
	parameter.AddToQuery(query, "statuses", req.Statuses)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/payment-admin/v2alpha1/magic-codes",
		Query:  query,
	}

	var resp ListMagicCodesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RefundMagicCode:
func (s *API) RefundMagicCode(req *RefundMagicCodeRequest, opts ...scw.RequestOption) (*MagicCode, error) {
	var err error

	if fmt.Sprint(req.MagicCodeID) == "" {
		return nil, errors.New("field MagicCodeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/payment-admin/v2alpha1/magic-codes/" + fmt.Sprint(req.MagicCodeID) + "/refund",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp MagicCode

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ProcessPayment:
func (s *API) ProcessPayment(req *ProcessPaymentRequest, opts ...scw.RequestOption) (*ProcessPaymentResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/payment-admin/v2alpha1/process-payment",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ProcessPaymentResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ImportPaymentMethods:
func (s *API) ImportPaymentMethods(req *ImportPaymentMethodsRequest, opts ...scw.RequestOption) (*ImportPaymentMethodsResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/payment-admin/v2alpha1/import-online-payment-methods",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ImportPaymentMethodsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendOnlineNotification:
func (s *API) SendOnlineNotification(req *SendOnlineNotificationRequest, opts ...scw.RequestOption) (*SendOnlineNotificationResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/payment-admin/v2alpha1/send-online-notification",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SendOnlineNotificationResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ProcessOnlineRefund:
func (s *API) ProcessOnlineRefund(req *ProcessOnlineRefundRequest, opts ...scw.RequestOption) error {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/payment-admin/v2alpha1/process-online-refund",
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

// GetAccountingLine:
func (s *API) GetAccountingLine(req *GetAccountingLineRequest, opts ...scw.RequestOption) (*AccountingLine, error) {
	var err error

	if fmt.Sprint(req.AccountingLineID) == "" {
		return nil, errors.New("field AccountingLineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/payment-admin/v2alpha1/accounting-lines/" + fmt.Sprint(req.AccountingLineID) + "",
	}

	var resp AccountingLine

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateAccountingLine:
func (s *API) UpdateAccountingLine(req *UpdateAccountingLineRequest, opts ...scw.RequestOption) (*AccountingLine, error) {
	var err error

	if fmt.Sprint(req.AccountingLineID) == "" {
		return nil, errors.New("field AccountingLineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/payment-admin/v2alpha1/accounting-lines/" + fmt.Sprint(req.AccountingLineID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AccountingLine

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SyncAccountingLineWithInvoice:
func (s *API) SyncAccountingLineWithInvoice(req *SyncAccountingLineWithInvoiceRequest, opts ...scw.RequestOption) (*AccountingLine, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/payment-admin/v2alpha1/accounting-lines/sync",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp AccountingLine

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdatePaymentMode:
func (s *API) UpdatePaymentMode(req *UpdatePaymentModeRequest, opts ...scw.RequestOption) error {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/payment-admin/v2alpha1/organizations/" + fmt.Sprint(req.OrganizationID) + "/payment-mode",
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

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package payment_admin provides methods and message types of the payment_admin v1 API.
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

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/payment_admin/v1/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/payment_admin/v1/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/payment_admin/v1/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/payment_admin/v1/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/payment_admin/v1/scw"
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

type CardNotificationOperationType string

const (
	CardNotificationOperationTypeUnknownOperationType = CardNotificationOperationType("unknown_operation_type")
	CardNotificationOperationTypePayment              = CardNotificationOperationType("payment")
	CardNotificationOperationTypeAuthorization        = CardNotificationOperationType("authorization")
	CardNotificationOperationTypeCapture              = CardNotificationOperationType("capture")
	CardNotificationOperationTypeRefund               = CardNotificationOperationType("refund")
	CardNotificationOperationTypeStopntimes           = CardNotificationOperationType("stopntimes")
)

func (enum CardNotificationOperationType) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_operation_type"
	}
	return string(enum)
}

func (enum CardNotificationOperationType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CardNotificationOperationType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CardNotificationOperationType(CardNotificationOperationType(tmp).String())
	return nil
}

type CardPSP string

const (
	CardPSPUnknownPsp   = CardPSP("unknown_psp")
	CardPSPBe2Bill      = CardPSP("Be2Bill")
	CardPSPStripe       = CardPSP("Stripe")
	CardPSPSepa         = CardPSP("Sepa")
	CardPSPWireTransfer = CardPSP("WireTransfer")
)

func (enum CardPSP) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_psp"
	}
	return string(enum)
}

func (enum CardPSP) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *CardPSP) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = CardPSP(CardPSP(tmp).String())
	return nil
}

type PaymentLineState string

const (
	PaymentLineStateUnknownState = PaymentLineState("unknown_state")
	PaymentLineStateDraft        = PaymentLineState("draft")
	PaymentLineStatePending      = PaymentLineState("pending")
	PaymentLineStateProcessing   = PaymentLineState("processing")
	PaymentLineStateClosed       = PaymentLineState("closed")
	PaymentLineStateLost         = PaymentLineState("lost")
)

func (enum PaymentLineState) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_state"
	}
	return string(enum)
}

func (enum PaymentLineState) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PaymentLineState) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PaymentLineState(PaymentLineState(tmp).String())
	return nil
}

// AccountingLineOrganization: accounting line organization.
type AccountingLineOrganization struct {
	// ID: the ID of the organization.
	ID string `json:"id"`

	// Name: the name of the organization.
	Name string `json:"name"`
}

// CardEmission: card emission.
type CardEmission struct {
	// ID: the ID of the card emission.
	ID string `json:"id"`

	// CreationDate: the creation date of the card emission.
	CreationDate *time.Time `json:"creation_date"`

	// OperationType: the operation type of the card emission.
	OperationType string `json:"operation_type"`

	// Info: the info of the card emission.
	Info string `json:"info"`

	// Card: the card ID the card emission.
	Card string `json:"card"`

	// Organization: the organization ID of the card emission.
	Organization string `json:"organization"`

	// ClientEmail: the client email of the card emission.
	ClientEmail string `json:"client_email"`

	// ClientReferrer: the client referrer of the card emission.
	ClientReferrer string `json:"client_referrer"`

	// ClientUseragent: the client useragent of the card emission.
	ClientUseragent string `json:"client_useragent"`

	// ClientIP: the client IP of the card emission.
	ClientIP string `json:"client_ip"`
}

// CardNotification: card notification.
type CardNotification struct {
	// ID: the ID of the card notification.
	ID string `json:"id"`

	// CreationDate: the creation date of the card notification.
	CreationDate *time.Time `json:"creation_date"`

	// OrderKey: the order key of the card notification.
	OrderKey string `json:"order_key"`

	// Alias: the alias of the card notification.
	Alias string `json:"alias"`

	// Version: the version the card notification.
	Version string `json:"version"`

	// TransactionID: the transaction ID of the card notification.
	TransactionID string `json:"transaction_id"`

	// CardValidity: the card validity of the card notification.
	CardValidity string `json:"card_validity"`

	// CardFullname: the card fullname of the card notification.
	CardFullname string `json:"card_fullname"`

	// CardCountry: the card country of the card notification.
	CardCountry string `json:"card_country"`

	// CardCode: the card code of the card notification.
	CardCode string `json:"card_code"`

	// CardType: the card type of the card notification.
	CardType string `json:"card_type"`

	// Message: the message of the card notification.
	Message string `json:"message"`

	// Language: the language of the card notification.
	Language string `json:"language"`

	// Currency: the currency of the card notification.
	Currency string `json:"currency"`

	// Execcode: the execcode of the card notification.
	Execcode string `json:"execcode"`

	// Amount: the amount of the card notification.
	Amount string `json:"amount"`

	// Descriptor: the descriptor of the card notification.
	Descriptor string `json:"descriptor"`

	// ClientEmail: the client email of the card notification.
	ClientEmail string `json:"client_email"`

	// OperationType: the operation type of the card notification.
	// Default value: unknown_operation_type
	OperationType CardNotificationOperationType `json:"operation_type"`

	// IDentifier: the identifier of the card notification.
	IDentifier string `json:"identifier"`

	// ClientIDent: the client ident of the card notification.
	ClientIDent string `json:"client_ident"`

	// NotificationJSON: the json with the card notification data.
	NotificationJSON string `json:"notification_json"`
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

	// InternalDescription: the internal description of the accounting line.
	InternalDescription string `json:"internal_description"`

	// State: the state of the accounting line.
	// Default value: unknown_state
	State AccountingLineState `json:"state"`

	// Description: the description of the accounting line.
	Description string `json:"description"`

	// Amount: the amount of the accounting line.
	Amount float64 `json:"amount"`

	// Currency: the currency of the accounting line.
	Currency string `json:"currency"`

	// Category: the category of the accounting line.
	// Default value: unknown_category
	Category AccountingLineCategory `json:"category"`

	// PaymentMode: the payment mode of the accounting line.
	PaymentMode string `json:"payment_mode"`

	// TryCount: the try count of the accounting line.
	TryCount int32 `json:"try_count"`

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

// Card: card.
type Card struct {
	// ID: the ID of the card.
	ID string `json:"id"`

	// CreationDate: the creation date of the card.
	CreationDate *time.Time `json:"creation_date"`

	// ModificationDate: the modification date of the card.
	ModificationDate *time.Time `json:"modification_date"`

	// Organization: the organization ID of the card.
	Organization string `json:"organization"`

	// Name: the name of the card.
	Name string `json:"name"`

	// Alias: the alias of the card.
	Alias string `json:"alias"`

	// Valid: indicates if the card is valid or not.
	Valid bool `json:"valid"`

	// Selected: indicates if the card is selected for making payments of an organization or not.
	Selected bool `json:"selected"`

	// Deleted: indicates if the card is deleted or not.
	Deleted bool `json:"deleted"`

	// Number: the masked card number.
	Number string `json:"number"`

	// Country: the country of the card.
	Country string `json:"country"`

	// Type: the type of the card.
	Type string `json:"type"`

	// ValidityDate: the validity date of the card.
	ValidityDate *time.Time `json:"validity_date"`

	// Psp: the PSP handled the card.
	// Default value: unknown_psp
	Psp CardPSP `json:"psp"`

	// CardEmission: the card emission linked to the card.
	CardEmission *CardEmission `json:"card_emission"`

	// Usable: indicates if the card is usable or not. That means that the card is not deleted, valid and not expired.
	Usable bool `json:"usable"`

	// CardNotification: the card notification linked to the card.
	CardNotification *CardNotification `json:"card_notification"`
}

// PaymentLine: payment line.
type PaymentLine struct {
	// ID: the ID of the payment line.
	ID string `json:"id"`

	// CreationDate: the creation date of the payment line.
	CreationDate *time.Time `json:"creation_date"`

	// ModificationDate: the modification date of the payment line.
	ModificationDate *time.Time `json:"modification_date"`

	// State: the state of the payment line.
	// Default value: unknown_state
	State PaymentLineState `json:"state"`

	// Description: the description of the payment line.
	Description string `json:"description"`

	// Amount: the amount of the payment line.
	Amount float64 `json:"amount"`

	// Currency: the currency of the payment line.
	Currency string `json:"currency"`

	// PaymentMode: the payment mode of the payment line.
	PaymentMode string `json:"payment_mode"`

	// TryCount: the try count of the payment line.
	TryCount int32 `json:"try_count"`

	// EventSentDate: the event sent date of the payment line.
	EventSentDate *time.Time `json:"event_sent_date"`

	// EventResponseDate: the event response date of the payment line.
	EventResponseDate *time.Time `json:"event_response_date"`

	// CardEmissionID: the card emission ID of the payment line.
	CardEmissionID string `json:"card_emission_id"`

	// CardNotificationID: the card notification ID of the payment line.
	CardNotificationID string `json:"card_notification_id"`

	// CardID: the card ID of the payment line.
	CardID string `json:"card_id"`

	// Be2billOrderid: the Be2Bill order ID of the payment line.
	Be2billOrderid string `json:"be2bill_orderid"`

	// Be2billTransactionid: the Be2Bill transaction ID of the payment line.
	Be2billTransactionid string `json:"be2bill_transactionid"`

	// Be2billExeccode: the Be2Bill execcode of the payment line.
	Be2billExeccode string `json:"be2bill_execcode"`

	// Be2billMessage: the Be2Bill message of the payment line.
	Be2billMessage string `json:"be2bill_message"`
}

// Sepa: sepa.
type Sepa struct {
	// MandateReference: the mandate reference of the SEPA.
	MandateReference string `json:"mandate_reference"`

	// CustomName: the custom name of the SEPA.
	CustomName string `json:"custom_name"`

	// Selected: indicates if the SEPA is selected for making payments of an organization or not.
	Selected bool `json:"selected"`

	// Name: the name of the SEPA.
	Name string `json:"name"`

	// Iban: the IBAN of the SEPA.
	Iban string `json:"iban"`

	// Bic: the BIC of the SEPA.
	Bic string `json:"bic"`

	// Valid: indicates if the SEPA is valid or not.
	Valid bool `json:"valid"`

	// OrganizationID: the organization ID of the SEPA.
	OrganizationID string `json:"organization_id"`
}

// ListAccountingLinesRequest: list accounting lines request.
type ListAccountingLinesRequest struct {
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`

	// Organization: filter accounting lines by third party organization ID.
	Organization *string `json:"-"`
}

// ListAccountingLinesResponse: list accounting lines response.
type ListAccountingLinesResponse struct {
	// TotalCount: the total number of accounting lines.
	TotalCount uint32 `json:"total_count"`

	// Accountinglines: the paginated returned accounting lines.
	Accountinglines []*AccountingLine `json:"accountinglines"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListAccountingLinesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListAccountingLinesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListAccountingLinesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Accountinglines = append(r.Accountinglines, results.Accountinglines...)
	r.TotalCount += uint32(len(results.Accountinglines))
	return uint32(len(results.Accountinglines)), nil
}

// ListCardsRequest: list cards request.
type ListCardsRequest struct {
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`

	// Organization: filter cards by organization ID.
	Organization *string `json:"-"`

	// IncludeInvalid: include invalid cards.
	IncludeInvalid *bool `json:"-"`

	// IncludeDeleted: include deleted cards.
	IncludeDeleted *bool `json:"-"`
}

// ListCardsResponse: list cards response.
type ListCardsResponse struct {
	// TotalCount: the total number of cards.
	TotalCount uint32 `json:"total_count"`

	// Cards: the paginated returned cards.
	Cards []*Card `json:"cards"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCardsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCardsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCardsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Cards = append(r.Cards, results.Cards...)
	r.TotalCount += uint32(len(results.Cards))
	return uint32(len(results.Cards)), nil
}

// ListPaymentLinesRequest: list payment lines request.
type ListPaymentLinesRequest struct {
	// AccountinglineID: filter payment lines by accounting line ID.
	AccountinglineID string `json:"-"`

	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`
}

// ListPaymentLinesResponse: list payment lines response.
type ListPaymentLinesResponse struct {
	// TotalCount: the total number of payment lines.
	TotalCount uint32 `json:"total_count"`

	// Paymentlines: the paginated returned payment lines.
	Paymentlines []*PaymentLine `json:"paymentlines"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPaymentLinesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPaymentLinesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListPaymentLinesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Paymentlines = append(r.Paymentlines, results.Paymentlines...)
	r.TotalCount += uint32(len(results.Paymentlines))
	return uint32(len(results.Paymentlines)), nil
}

// ListSepasRequest: list sepas request.
type ListSepasRequest struct {
	// Page: a positive integer to choose the page to return.
	Page *int32 `json:"-"`

	// PerPage: a positive integer lower or equal to 100 to select the number of items to return.
	PerPage *uint32 `json:"-"`

	// Organization: filter SEPAs by organization ID.
	Organization *string `json:"-"`

	// IncludeInvalid: include invalid SEPAs.
	IncludeInvalid *bool `json:"-"`

	// IncludeDeleted: include deleted SEPAs.
	IncludeDeleted *bool `json:"-"`
}

// ListSepasResponse: list sepas response.
type ListSepasResponse struct {
	// TotalCount: the total number of SEPAs.
	TotalCount uint32 `json:"total_count"`

	// Sepas: the paginated returned SEPAs.
	Sepas []*Sepa `json:"sepas"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSepasResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSepasResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListSepasResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Sepas = append(r.Sepas, results.Sepas...)
	r.TotalCount += uint32(len(results.Sepas))
	return uint32(len(results.Sepas)), nil
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

// GetServiceInfo:
func (s *API) GetServiceInfo(opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/payment-admin/v1",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListAccountingLines: List accounting lines.
func (s *API) ListAccountingLines(req *ListAccountingLinesRequest, opts ...scw.RequestOption) (*ListAccountingLinesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "organization", req.Organization)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/payment-admin/v1/accountinglines",
		Query:  query,
	}

	var resp ListAccountingLinesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListCards: List cards.
func (s *API) ListCards(req *ListCardsRequest, opts ...scw.RequestOption) (*ListCardsResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "include_invalid", req.IncludeInvalid)
	parameter.AddToQuery(query, "include_deleted", req.IncludeDeleted)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/payment-admin/v1/cards",
		Query:  query,
	}

	var resp ListCardsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListPaymentLines: List payment lines.
func (s *API) ListPaymentLines(req *ListPaymentLinesRequest, opts ...scw.RequestOption) (*ListPaymentLinesResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "per_page", req.PerPage)

	if fmt.Sprint(req.AccountinglineID) == "" {
		return nil, errors.New("field AccountinglineID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/payment-admin/v1/accountinglines/" + fmt.Sprint(req.AccountinglineID) + "/paymentlines",
		Query:  query,
	}

	var resp ListPaymentLinesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSepas: List SEPAs.
func (s *API) ListSepas(req *ListSepasRequest, opts ...scw.RequestOption) (*ListSepasResponse, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "per_page", req.PerPage)
	parameter.AddToQuery(query, "organization", req.Organization)
	parameter.AddToQuery(query, "include_invalid", req.IncludeInvalid)
	parameter.AddToQuery(query, "include_deleted", req.IncludeDeleted)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/payment-admin/v1/sepa",
		Query:  query,
	}

	var resp ListSepasResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package billing provides methods and message types of the billing v2 API.
package billing

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

type BudgetAlertNotificationType string

const (
	BudgetAlertNotificationTypeUnknownType = BudgetAlertNotificationType("unknown_type")
	// Notify by sending a sms.
	BudgetAlertNotificationTypeSms = BudgetAlertNotificationType("sms")
	// Notify by sending an email.
	BudgetAlertNotificationTypeEmail = BudgetAlertNotificationType("email")
	// Notify by using a webhook.
	BudgetAlertNotificationTypeWebhook = BudgetAlertNotificationType("webhook")
)

func (enum BudgetAlertNotificationType) String() string {
	if enum == "" {
		// return default value if empty
		return string(BudgetAlertNotificationTypeUnknownType)
	}
	return string(enum)
}

func (enum BudgetAlertNotificationType) Values() []BudgetAlertNotificationType {
	return []BudgetAlertNotificationType{
		"unknown_type",
		"sms",
		"email",
		"webhook",
	}
}

func (enum BudgetAlertNotificationType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *BudgetAlertNotificationType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = BudgetAlertNotificationType(BudgetAlertNotificationType(tmp).String())
	return nil
}

// BudgetAlertNotification: budget alert notification.
type BudgetAlertNotification struct {
	// ID: alert notification's ID.
	ID string `json:"id"`

	// CreatedAt: the creation date of the alert notification.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: the last modification date of the alert notification.
	UpdatedAt *time.Time `json:"updated_at"`

	// Type: kind of notification.
	// Default value: unknown_type
	Type BudgetAlertNotificationType `json:"type"`

	// Recipients: list of recipients for this alert.
	Recipients []string `json:"recipients"`
}

// BudgetAlert: budget alert.
type BudgetAlert struct {
	// ID: alert's ID.
	ID string `json:"id"`

	// CreatedAt: the creation date of the alert.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: the last modification date of the alert.
	UpdatedAt *time.Time `json:"updated_at"`

	// Threshold: percentage threshold of the budget's limit for which the alert is triggered.
	Threshold uint32 `json:"threshold"`

	// Notifications: notifications to send when the alert is triggered.
	Notifications []*BudgetAlertNotification `json:"notifications"`
}

// Budget: budget.
type Budget struct {
	// ID: budget's ID.
	ID string `json:"id"`

	// CreatedAt: the creation date of the budget.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: the last modification date of the budget.
	UpdatedAt *time.Time `json:"updated_at"`

	// OrganizationID: the organization ID of the budget.
	OrganizationID string `json:"organization_id"`

	// Enabled: whether the budget is enabled or not.
	Enabled bool `json:"enabled"`

	// ConsumptionLimit: cost limit for this budget.
	ConsumptionLimit *scw.Money `json:"consumption_limit"`

	// Alerts: alerts defined for this budget.
	Alerts []*BudgetAlert `json:"alerts"`
}

// CreateBudgetAlertNotificationRequest: create budget alert notification request.
type CreateBudgetAlertNotificationRequest struct {
	// BudgetAlertID: the ID of the budget alert to create notification for.
	BudgetAlertID string `json:"budget_alert_id"`

	// SmsPhoneNumbers: list of phone numbers to receive sms notifications.
	// Precisely one of SmsPhoneNumbers, EmailAddresses, WebhookURLs must be set.
	SmsPhoneNumbers *[]string `json:"sms_phone_numbers,omitempty"`

	// EmailAddresses: list of email addresses to receive email notifications.
	// Precisely one of SmsPhoneNumbers, EmailAddresses, WebhookURLs must be set.
	EmailAddresses *[]string `json:"email_addresses,omitempty"`

	// WebhookURLs: list of webhook url to receive webhook notifications.
	// Precisely one of SmsPhoneNumbers, EmailAddresses, WebhookURLs must be set.
	WebhookURLs *[]string `json:"webhook_urls,omitempty"`
}

// CreateBudgetAlertRequest: create budget alert request.
type CreateBudgetAlertRequest struct {
	// BudgetID: the ID of the budget to create alert for.
	BudgetID string `json:"budget_id"`

	// Threshold: threshold above which the alert is sent.
	Threshold uint32 `json:"threshold"`
}

// CreateBudgetRequest: create budget request.
type CreateBudgetRequest struct {
	// OrganizationID: the Organization ID of the budget.
	OrganizationID string `json:"organization_id"`

	// ConsumptionLimit: cost limit for the budget.
	ConsumptionLimit uint32 `json:"consumption_limit"`

	// Enabled: whether the budget is enabled or not.
	Enabled bool `json:"enabled"`
}

// DeleteBudgetAlertNotificationRequest: delete budget alert notification request.
type DeleteBudgetAlertNotificationRequest struct {
	// BudgetAlertNotificationID: the ID of the budget alert notification to delete.
	BudgetAlertNotificationID string `json:"-"`
}

// DeleteBudgetAlertRequest: delete budget alert request.
type DeleteBudgetAlertRequest struct {
	// BudgetAlertID: the ID of the budget alert to delete.
	BudgetAlertID string `json:"-"`
}

// DeleteBudgetRequest: delete budget request.
type DeleteBudgetRequest struct {
	// BudgetID: the ID of the budget to delete.
	BudgetID string `json:"-"`
}

// GetBudgetRequest: get budget request.
type GetBudgetRequest struct {
	// BudgetID: the ID of the budget.
	BudgetID string `json:"-"`
}

// ListBudgetsRequest: list budgets request.
type ListBudgetsRequest struct {
	// OrganizationID: filter by organization ID.
	OrganizationID *string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: positive integer to select the number of items to return.
	PageSize *uint32 `json:"-"`
}

// ListBudgetsResponse: list budgets response.
type ListBudgetsResponse struct {
	// Budgets: detailed budget list.
	Budgets []*Budget `json:"budgets"`

	// TotalCount: total number of items.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListBudgetsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListBudgetsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListBudgetsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Budgets = append(r.Budgets, results.Budgets...)
	r.TotalCount += uint64(len(results.Budgets))
	return uint64(len(results.Budgets)), nil
}

// UpdateBudgetAlertNotificationRequest: update budget alert notification request.
type UpdateBudgetAlertNotificationRequest struct {
	// BudgetAlertNotificationID: the ID of the budget alert notification to update.
	BudgetAlertNotificationID string `json:"-"`

	// SmsPhoneNumbers: list of phone numbers to receive sms notifications.
	// Precisely one of SmsPhoneNumbers, EmailAddresses, WebhookURLs must be set.
	SmsPhoneNumbers *[]string `json:"sms_phone_numbers,omitempty"`

	// EmailAddresses: list of email addresses to receive email notifications.
	// Precisely one of SmsPhoneNumbers, EmailAddresses, WebhookURLs must be set.
	EmailAddresses *[]string `json:"email_addresses,omitempty"`

	// WebhookURLs: list of webhook url to receive webhook notifications.
	// Precisely one of SmsPhoneNumbers, EmailAddresses, WebhookURLs must be set.
	WebhookURLs *[]string `json:"webhook_urls,omitempty"`
}

// UpdateBudgetAlertRequest: update budget alert request.
type UpdateBudgetAlertRequest struct {
	// BudgetAlertID: the ID of the budget alert to update.
	BudgetAlertID string `json:"-"`

	// Threshold: threshold above which the alert is sent.
	Threshold uint32 `json:"threshold"`
}

// UpdateBudgetRequest: update budget request.
type UpdateBudgetRequest struct {
	// BudgetID: the ID of the budget to update.
	BudgetID string `json:"-"`

	// ConsumptionLimit: cost limit for the budget.
	ConsumptionLimit *uint32 `json:"consumption_limit,omitempty"`

	// Enabled: whether the budget will be enabled or not.
	Enabled *bool `json:"enabled,omitempty"`
}

// This API allows you to query billing related objects.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// ListBudgets: List your budgets, filtering by `organization_id`.
func (s *API) ListBudgets(req *ListBudgetsRequest, opts ...scw.RequestOption) (*ListBudgetsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2/budgets",
		Query:  query,
	}

	var resp ListBudgetsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetBudget: Fetch a budget.
func (s *API) GetBudget(req *GetBudgetRequest, opts ...scw.RequestOption) (*Budget, error) {
	var err error

	if fmt.Sprint(req.BudgetID) == "" {
		return nil, errors.New("field BudgetID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/billing/v2/budgets/" + fmt.Sprint(req.BudgetID) + "",
	}

	var resp Budget

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateBudget: Create a new budget.
func (s *API) CreateBudget(req *CreateBudgetRequest, opts ...scw.RequestOption) (*Budget, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/billing/v2/budgets",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Budget

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBudget: Update a budget.
func (s *API) UpdateBudget(req *UpdateBudgetRequest, opts ...scw.RequestOption) (*Budget, error) {
	var err error

	if fmt.Sprint(req.BudgetID) == "" {
		return nil, errors.New("field BudgetID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/billing/v2/budgets/" + fmt.Sprint(req.BudgetID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Budget

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBudget: Delete a budget.
func (s *API) DeleteBudget(req *DeleteBudgetRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.BudgetID) == "" {
		return errors.New("field BudgetID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/billing/v2/budgets/" + fmt.Sprint(req.BudgetID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateBudgetAlert: Create a new budget alert.
func (s *API) CreateBudgetAlert(req *CreateBudgetAlertRequest, opts ...scw.RequestOption) (*BudgetAlert, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/billing/v2/budget-alerts",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp BudgetAlert

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBudgetAlert: Update a budget alert.
func (s *API) UpdateBudgetAlert(req *UpdateBudgetAlertRequest, opts ...scw.RequestOption) (*BudgetAlert, error) {
	var err error

	if fmt.Sprint(req.BudgetAlertID) == "" {
		return nil, errors.New("field BudgetAlertID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/billing/v2/budget-alerts/" + fmt.Sprint(req.BudgetAlertID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp BudgetAlert

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBudgetAlert: Delete a budget alert.
func (s *API) DeleteBudgetAlert(req *DeleteBudgetAlertRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.BudgetAlertID) == "" {
		return errors.New("field BudgetAlertID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/billing/v2/budget-alerts/" + fmt.Sprint(req.BudgetAlertID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// CreateBudgetAlertNotification: Create a new budget alert notification.
func (s *API) CreateBudgetAlertNotification(req *CreateBudgetAlertNotificationRequest, opts ...scw.RequestOption) (*BudgetAlertNotification, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/billing/v2/budget-alert-notifications",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp BudgetAlertNotification

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateBudgetAlertNotification: Update a budget alert notification.
func (s *API) UpdateBudgetAlertNotification(req *UpdateBudgetAlertNotificationRequest, opts ...scw.RequestOption) (*BudgetAlertNotification, error) {
	var err error

	if fmt.Sprint(req.BudgetAlertNotificationID) == "" {
		return nil, errors.New("field BudgetAlertNotificationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/billing/v2/budget-alert-notifications/" + fmt.Sprint(req.BudgetAlertNotificationID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp BudgetAlertNotification

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteBudgetAlertNotification: Delete a budget alert notification.
func (s *API) DeleteBudgetAlertNotification(req *DeleteBudgetAlertNotificationRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.BudgetAlertNotificationID) == "" {
		return errors.New("field BudgetAlertNotificationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/billing/v2/budget-alert-notifications/" + fmt.Sprint(req.BudgetAlertNotificationID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

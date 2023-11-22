// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package payment provides methods and message types of the payment v2beta1 API.
package payment

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

type RetryPaymentTransactionResponsePaymentStatus string

const (
	RetryPaymentTransactionResponsePaymentStatusUnknownPaymentStatus = RetryPaymentTransactionResponsePaymentStatus("unknown_payment_status")
	RetryPaymentTransactionResponsePaymentStatusSuccess              = RetryPaymentTransactionResponsePaymentStatus("success")
	RetryPaymentTransactionResponsePaymentStatusFailure              = RetryPaymentTransactionResponsePaymentStatus("failure")
)

func (enum RetryPaymentTransactionResponsePaymentStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_payment_status"
	}
	return string(enum)
}

func (enum RetryPaymentTransactionResponsePaymentStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *RetryPaymentTransactionResponsePaymentStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = RetryPaymentTransactionResponsePaymentStatus(RetryPaymentTransactionResponsePaymentStatus(tmp).String())
	return nil
}

// RetryPaymentTransactionRequest: retry payment transaction request.
type RetryPaymentTransactionRequest struct {
	// TransactionID: represent the transaction_key generated on the Online side.
	TransactionID uint32 `json:"-"`

	// OrganizationID: organization ID to retry payment.
	OrganizationID string `json:"organization_id"`

	// FullRedirect: optional field, when it's set to True we redirect the customer back after the payment. This redirection is only available for cards registered on Stripe.
	FullRedirect bool `json:"full_redirect"`
}

// RetryPaymentTransactionResponse: retry payment transaction response.
type RetryPaymentTransactionResponse struct {
	// PaymentStatus: the processed payment status.
	// Default value: unknown_payment_status
	PaymentStatus RetryPaymentTransactionResponsePaymentStatus `json:"payment_status"`

	// Token3ds: in case of 3DS a secret key must be returned to front-end.
	Token3ds *string `json:"token_3ds"`

	// RedirectToURL: in case of full_redirect mode this URL must be redirected by hand by the front-end. It will allow to go to the user's bank site.
	RedirectToURL *string `json:"redirect_to_url"`
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

// RetryPaymentTransaction:
func (s *API) RetryPaymentTransaction(req *RetryPaymentTransactionRequest, opts ...scw.RequestOption) (*RetryPaymentTransactionResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.TransactionID) == "" {
		return nil, errors.New("field TransactionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/payment/v2beta1/transactions/" + fmt.Sprint(req.TransactionID) + "/retry",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RetryPaymentTransactionResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

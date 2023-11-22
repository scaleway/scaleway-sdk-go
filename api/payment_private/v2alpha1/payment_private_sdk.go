// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package payment_private provides methods and message types of the payment_private v2alpha1 API.
package payment_private

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/payment_private/v2alpha1/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/payment_private/v2alpha1/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/payment_private/v2alpha1/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/payment_private/v2alpha1/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/payment_private/v2alpha1/scw"
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

// MagicCode: magic code.
type MagicCode struct {
	ID string `json:"id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	ExpiresAt *time.Time `json:"expires_at"`

	// Status: default value: unknown
	Status MagicCodeStatus `json:"status"`

	CardID string `json:"card_id"`

	OrganizationID string `json:"organization_id"`
}

// ConsoleAPIListMagicCodesRequest: console api list magic codes request.
type ConsoleAPIListMagicCodesRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListMagicCodesRequestOrderBy `json:"-"`

	CardID *string `json:"-"`

	// Status: default value: unknown
	Status MagicCodeStatus `json:"-"`

	OrganizationID *string `json:"-"`
}

// ConsoleAPISendMagicCodeRequest: console api send magic code request.
type ConsoleAPISendMagicCodeRequest struct {
	CardID string `json:"card_id"`
}

// ConsoleAPIValidateMagicCodeRequest: console api validate magic code request.
type ConsoleAPIValidateMagicCodeRequest struct {
	MagicCodeID string `json:"-"`

	MagicCode string `json:"magic_code"`
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

type ConsoleAPI struct {
	client *scw.Client
}

// NewConsoleAPI returns a ConsoleAPI object from a Scaleway client.
func NewConsoleAPI(client *scw.Client) *ConsoleAPI {
	return &ConsoleAPI{
		client: client,
	}
}

// GetServiceInfo:
func (s *ConsoleAPI) GetServiceInfo(opts ...scw.RequestOption) (*scw.ServiceInfo, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/payment-private/v2alpha1",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SendMagicCode:
func (s *ConsoleAPI) SendMagicCode(req *ConsoleAPISendMagicCodeRequest, opts ...scw.RequestOption) (*MagicCode, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/payment-private/v2alpha1/magic-codes",
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
func (s *ConsoleAPI) ValidateMagicCode(req *ConsoleAPIValidateMagicCodeRequest, opts ...scw.RequestOption) (*MagicCode, error) {
	var err error

	if fmt.Sprint(req.MagicCodeID) == "" {
		return nil, errors.New("field MagicCodeID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/payment-private/v2alpha1/magic-codes/" + fmt.Sprint(req.MagicCodeID) + "/validate",
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
func (s *ConsoleAPI) ListMagicCodes(req *ConsoleAPIListMagicCodesRequest, opts ...scw.RequestOption) (*ListMagicCodesResponse, error) {
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

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/payment-private/v2alpha1/magic-codes",
		Query:  query,
	}

	var resp ListMagicCodesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

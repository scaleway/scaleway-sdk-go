// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package account_admin provides methods and message types of the account_admin v2alpha1 API.
package account_admin

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
	std "gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/api/std"
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

type IdentityVerificationRejectionReason string

const (
	// Unknown rejection reason.
	IdentityVerificationRejectionReasonUnknownRejectionReason = IdentityVerificationRejectionReason("unknown_rejection_reason")
	// Not readable ID.
	IdentityVerificationRejectionReasonNotReadableID = IdentityVerificationRejectionReason("not_readable_id")
	// Face do not match ID.
	IdentityVerificationRejectionReasonFaceDoNotMatchID = IdentityVerificationRejectionReason("face_do_not_match_id")
	// Expired ID.
	IdentityVerificationRejectionReasonExpiredID = IdentityVerificationRejectionReason("expired_id")
	// No uploaded ID.
	IdentityVerificationRejectionReasonNoUploadedID = IdentityVerificationRejectionReason("no_uploaded_id")
)

func (enum IdentityVerificationRejectionReason) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_rejection_reason"
	}
	return string(enum)
}

func (enum IdentityVerificationRejectionReason) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IdentityVerificationRejectionReason) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IdentityVerificationRejectionReason(IdentityVerificationRejectionReason(tmp).String())
	return nil
}

type IdentityVerificationStatus string

const (
	// Unknown status.
	IdentityVerificationStatusUnknown = IdentityVerificationStatus("unknown")
	// Pending.
	IdentityVerificationStatusPending = IdentityVerificationStatus("pending")
	// Validated.
	IdentityVerificationStatusValidated = IdentityVerificationStatus("validated")
	// Waiting for approval.
	IdentityVerificationStatusWaitingForApproval = IdentityVerificationStatus("waiting_for_approval")
	// Rejected.
	IdentityVerificationStatusRejected = IdentityVerificationStatus("rejected")
	// Failed.
	IdentityVerificationStatusFailed = IdentityVerificationStatus("failed")
)

func (enum IdentityVerificationStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown"
	}
	return string(enum)
}

func (enum IdentityVerificationStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *IdentityVerificationStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = IdentityVerificationStatus(IdentityVerificationStatus(tmp).String())
	return nil
}

type ListIdentityVerificationsRequestOrderBy string

const (
	// Creation date ascending.
	ListIdentityVerificationsRequestOrderByCreatedAtAsc = ListIdentityVerificationsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListIdentityVerificationsRequestOrderByCreatedAtDesc = ListIdentityVerificationsRequestOrderBy("created_at_desc")
)

func (enum ListIdentityVerificationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListIdentityVerificationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListIdentityVerificationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListIdentityVerificationsRequestOrderBy(ListIdentityVerificationsRequestOrderBy(tmp).String())
	return nil
}

// IdentityVerificationJumioVerification: identity verification jumio verification.
type IdentityVerificationJumioVerification struct {
	IframeURL string `json:"iframe_url"`

	BackofficeURL string `json:"backoffice_url"`

	RelatedBackofficeURLs []string `json:"related_backoffice_urls"`
}

// IdentityVerification: identity verification.
type IdentityVerification struct {
	ID string `json:"id"`

	// Status: default value: unknown
	Status IdentityVerificationStatus `json:"status"`

	OrganizationID string `json:"organization_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	// Precisely one of Jumio must be set.
	Jumio *IdentityVerificationJumioVerification `json:"jumio,omitempty"`

	// RejectionReason: default value: unknown_rejection_reason
	RejectionReason IdentityVerificationRejectionReason `json:"rejection_reason"`

	InternalNote *string `json:"internal_note"`

	VerifiedManually bool `json:"verified_manually"`
}

// User: user.
type User struct {
	ID string `json:"id"`

	Email string `json:"email"`

	Firstname string `json:"firstname"`

	Lastname string `json:"lastname"`

	// Locale: default value: unknown_language_code
	Locale std.LanguageCode `json:"locale"`
}

// GetOrganizationRequest: get organization request.
type GetOrganizationRequest struct {
	OrganizationID string `json:"-"`
}

// KYCApiListIdentityVerificationsRequest: kyc api list identity verifications request.
type KYCApiListIdentityVerificationsRequest struct {
	// OrderBy: default value: created_at_asc
	OrderBy ListIdentityVerificationsRequestOrderBy `json:"-"`

	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	OrganizationID *string `json:"-"`

	// Deprecated: Status: default value: unknown
	Status *IdentityVerificationStatus `json:"-"`

	Statuses []IdentityVerificationStatus `json:"-"`

	ExcludeStatuses []IdentityVerificationStatus `json:"-"`

	ExcludeRejectionReasons []IdentityVerificationRejectionReason `json:"-"`
}

// KYCApiStartIdentityVerificationRequest: kyc api start identity verification request.
type KYCApiStartIdentityVerificationRequest struct {
	OrganizationID string `json:"organization_id"`
}

// KYCApiUpdateIdentityVerificationRequest: kyc api update identity verification request.
type KYCApiUpdateIdentityVerificationRequest struct {
	IDentityVerificationID string `json:"-"`

	// Status: default value: unknown
	Status IdentityVerificationStatus `json:"status"`

	// RejectionReason: default value: unknown_rejection_reason
	RejectionReason IdentityVerificationRejectionReason `json:"rejection_reason"`

	InternalNote *string `json:"internal_note,omitempty"`
}

// ListIdentityVerificationsResponse: list identity verifications response.
type ListIdentityVerificationsResponse struct {
	IDentityVerifications []*IdentityVerification `json:"identity_verifications"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListIdentityVerificationsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListIdentityVerificationsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListIdentityVerificationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.IDentityVerifications = append(r.IDentityVerifications, results.IDentityVerifications...)
	r.TotalCount += uint32(len(results.IDentityVerifications))
	return uint32(len(results.IDentityVerifications)), nil
}

// Organization: organization.
type Organization struct {
	ID string `json:"id"`

	Name string `json:"name"`

	SupportID string `json:"support_id"`

	SupportPin string `json:"support_pin"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	AddressLine1 string `json:"address_line1"`

	AddressLine2 string `json:"address_line2"`

	ZipCode string `json:"zip_code"`

	City string `json:"city"`

	// CountryCode: default value: unknown_country_code
	CountryCode std.CountryCode `json:"country_code"`

	SubdivisionCode string `json:"subdivision_code"`

	Creator *User `json:"creator"`

	CustomerLevels []string `json:"customer_levels"`
}

// UpdateOrganizationRequest: update organization request.
type UpdateOrganizationRequest struct {
	OrganizationID string `json:"-"`

	CustomerLevels *[]string `json:"customer_levels,omitempty"`
}

// Account Admin API.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// GetOrganization: This methods allows to get details about a specific organization.
func (s *API) GetOrganization(req *GetOrganizationRequest, opts ...scw.RequestOption) (*Organization, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2alpha1/organizations/" + fmt.Sprint(req.OrganizationID) + "",
	}

	var resp Organization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateOrganization: This methods allows to update details about a specific organization.
func (s *API) UpdateOrganization(req *UpdateOrganizationRequest, opts ...scw.RequestOption) (*Organization, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account-admin/v2alpha1/organizations/" + fmt.Sprint(req.OrganizationID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Organization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Account Admin KYC API.
type KYCAPI struct {
	client *scw.Client
}

// NewKYCAPI returns a KYCAPI object from a Scaleway client.
func NewKYCAPI(client *scw.Client) *KYCAPI {
	return &KYCAPI{
		client: client,
	}
}

// ListIdentityVerifications: List identity verifications.
func (s *KYCAPI) ListIdentityVerifications(req *KYCApiListIdentityVerificationsRequest, opts ...scw.RequestOption) (*ListIdentityVerificationsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "statuses", req.Statuses)
	parameter.AddToQuery(query, "exclude_statuses", req.ExcludeStatuses)
	parameter.AddToQuery(query, "exclude_rejection_reasons", req.ExcludeRejectionReasons)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account-admin/v2alpha1/kyc/identity-verifications",
		Query:  query,
	}

	var resp ListIdentityVerificationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// StartIdentityVerification: Start identity verification.
func (s *KYCAPI) StartIdentityVerification(req *KYCApiStartIdentityVerificationRequest, opts ...scw.RequestOption) (*IdentityVerification, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account-admin/v2alpha1/kyc/identity-verifications",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IdentityVerification

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateIdentityVerification: Update identity verification.
func (s *KYCAPI) UpdateIdentityVerification(req *KYCApiUpdateIdentityVerificationRequest, opts ...scw.RequestOption) (*IdentityVerification, error) {
	var err error

	if fmt.Sprint(req.IDentityVerificationID) == "" {
		return nil, errors.New("field IDentityVerificationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account-admin/v2alpha1/kyc/identity-verifications/" + fmt.Sprint(req.IDentityVerificationID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp IdentityVerification

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

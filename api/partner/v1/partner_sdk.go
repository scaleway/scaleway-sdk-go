// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package partner provides methods and message types of the partner v1 API.
package partner

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

type ListOrganizationsRequestOrderBy string

const (
	ListOrganizationsRequestOrderByCreatedAtAsc  = ListOrganizationsRequestOrderBy("created_at_asc")
	ListOrganizationsRequestOrderByCreatedAtDesc = ListOrganizationsRequestOrderBy("created_at_desc")
)

func (enum ListOrganizationsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListOrganizationsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListOrganizationsRequestOrderBy) Values() []ListOrganizationsRequestOrderBy {
	return []ListOrganizationsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
	}
}

func (enum ListOrganizationsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListOrganizationsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListOrganizationsRequestOrderBy(ListOrganizationsRequestOrderBy(tmp).String())
	return nil
}

type OrganizationLockedBy string

const (
	OrganizationLockedByUnknownLockedBy = OrganizationLockedBy("unknown_locked_by")
	OrganizationLockedByPartner         = OrganizationLockedBy("partner")
	OrganizationLockedByScaleway        = OrganizationLockedBy("scaleway")
)

func (enum OrganizationLockedBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(OrganizationLockedByUnknownLockedBy)
	}
	return string(enum)
}

func (enum OrganizationLockedBy) Values() []OrganizationLockedBy {
	return []OrganizationLockedBy{
		"unknown_locked_by",
		"partner",
		"scaleway",
	}
}

func (enum OrganizationLockedBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationLockedBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationLockedBy(OrganizationLockedBy(tmp).String())
	return nil
}

type OrganizationStatus string

const (
	OrganizationStatusUnknownStatus = OrganizationStatus("unknown_status")
	OrganizationStatusOpened        = OrganizationStatus("opened")
	OrganizationStatusLocked        = OrganizationStatus("locked")
	OrganizationStatusClosed        = OrganizationStatus("closed")
)

func (enum OrganizationStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(OrganizationStatusUnknownStatus)
	}
	return string(enum)
}

func (enum OrganizationStatus) Values() []OrganizationStatus {
	return []OrganizationStatus{
		"unknown_status",
		"opened",
		"locked",
		"closed",
	}
}

func (enum OrganizationStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *OrganizationStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = OrganizationStatus(OrganizationStatus(tmp).String())
	return nil
}

// Organization: organization.
type Organization struct {
	// ID: ID of the organization.
	ID string `json:"id"`

	// Name: name of the organization.
	Name string `json:"name"`

	// Email: organization owner's email.
	Email string `json:"email"`

	// Status: the current status of the organization.
	// Default value: unknown_status
	Status OrganizationStatus `json:"status"`

	// OwnerFirstname: organization owner's first name.
	OwnerFirstname string `json:"owner_firstname"`

	// OwnerLastname: organization owner's last name.
	OwnerLastname string `json:"owner_lastname"`

	// CreatedAt: date of organization creation.
	CreatedAt *time.Time `json:"created_at"`

	// PhoneNumber: organization owner's phone number.
	PhoneNumber *string `json:"phone_number"`

	// SirenNumber: siren number of the organization.
	SirenNumber *string `json:"siren_number"`

	// CustomerID: customer ID associated with this organization.
	CustomerID string `json:"customer_id"`

	// LockReasonMessage: if the organization is locked, this field will contain a human-readable reason.
	LockReasonMessage string `json:"lock_reason_message"`

	// LockedBy: originator of the lock. Can be one of "partner" or "scaleway".
	// Default value: unknown_locked_by
	LockedBy OrganizationLockedBy `json:"locked_by"`

	// LockedAt: date of lock.
	LockedAt *time.Time `json:"locked_at"`

	// PictureLink: link to the Organization's picture.
	PictureLink *string `json:"picture_link"`

	// Comment: a comment about the organization.
	Comment string `json:"comment"`
}

// CreateOrganizationRequest: create organization request.
type CreateOrganizationRequest struct {
	// PartnerID: your personal `partner_id`. This is the same as your Organization ID.
	PartnerID string `json:"partner_id"`

	// Email: the email of the new organization owner.
	Email string `json:"email"`

	// OrganizationName: the name of the organization you want to create. Usually the company name.
	OrganizationName string `json:"organization_name"`

	// OwnerFirstname: the first name of the new organization owner.
	OwnerFirstname string `json:"owner_firstname"`

	// OwnerLastname: the last name of the new organization owner.
	OwnerLastname string `json:"owner_lastname"`

	// PhoneNumber: the phone number of the new organization owner.
	PhoneNumber *string `json:"phone_number,omitempty"`

	// CustomerID: a custom ID for the customer in your own infrastructure.
	CustomerID string `json:"customer_id"`

	// SirenNumber: a SIREN number for the customer.
	SirenNumber *string `json:"siren_number,omitempty"`
}

// GetOrganizationRequest: get organization request.
type GetOrganizationRequest struct {
	// OrganizationID: the ID of the organization you want to GET.
	OrganizationID string `json:"-"`
}

// ListOrganizationsRequest: list organizations request.
type ListOrganizationsRequest struct {
	PageSize *uint32 `json:"-"`

	Page *int32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListOrganizationsRequestOrderBy `json:"-"`

	// Status: only list organizations with this status.
	// Default value: unknown_status
	Status OrganizationStatus `json:"-"`

	// Email: only list organizations created with this email.
	Email *string `json:"-"`

	// CustomerID: only list organizations attached to this Customer ID.
	// If the customer ID was changed only the last one can be used.
	CustomerID *string `json:"-"`

	// LockedBy: only list organizations locked by a certain entity.
	// Default value: unknown_locked_by
	LockedBy OrganizationLockedBy `json:"-"`
}

// ListOrganizationsResponse: list organizations response.
type ListOrganizationsResponse struct {
	// Organizations: list of organizations.
	Organizations []*Organization `json:"organizations"`

	// TotalCount: total number of organizations.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListOrganizationsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListOrganizationsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListOrganizationsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Organizations = append(r.Organizations, results.Organizations...)
	r.TotalCount += uint64(len(results.Organizations))
	return uint64(len(results.Organizations)), nil
}

// LockOrganizationRequest: lock organization request.
type LockOrganizationRequest struct {
	// OrganizationID: the ID of the organization you want to lock.
	OrganizationID string `json:"-"`
}

// RequestAdminRoleRequest: request admin role request.
type RequestAdminRoleRequest struct {
	// OrganizationID: the ID of the organization you want to be invited to.
	OrganizationID string `json:"-"`

	// Username: the member username.
	Username string `json:"username"`

	// Email: the member email.
	Email string `json:"email"`

	// Password: the member password.
	Password string `json:"password"`
}

// UnlockOrganizationRequest: unlock organization request.
type UnlockOrganizationRequest struct {
	// OrganizationID: the ID of the organization you want to unlock.
	OrganizationID string `json:"-"`
}

// UpdateOrganizationRequest: update organization request.
type UpdateOrganizationRequest struct {
	// OrganizationID: the ID of the organization you want to update.
	OrganizationID string `json:"-"`

	// Email: the new email.
	Email *string `json:"email,omitempty"`

	// Name: the new name.
	Name *string `json:"name,omitempty"`

	// OwnerFirstname: the first name of the new owner.
	OwnerFirstname *string `json:"owner_firstname,omitempty"`

	// OwnerLastname: the last name of the new owner.
	OwnerLastname *string `json:"owner_lastname,omitempty"`

	// PhoneNumber: the phone number of the new owner.
	PhoneNumber *string `json:"phone_number,omitempty"`

	// CustomerID: customer ID associated with this organization.
	// Warning: Changing this value will only affect future invoices.
	// If you try to change this value after the 25th of the month, we cannot guarantee that this will take effect on the invoice issued for the current month.
	CustomerID *string `json:"customer_id,omitempty"`

	// Comment: a comment about the organization.
	Comment *string `json:"comment,omitempty"`
}

// Scaleway Partner API ( for partner only ).
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// RequestAdminRole: Invite a partner user in the customer organization.
func (s *API) RequestAdminRole(req *RequestAdminRoleRequest, opts ...scw.RequestOption) error {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/partner/v1/organizations/" + fmt.Sprint(req.OrganizationID) + "/request-admin-role",
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

// CreateOrganization: Create a new organization.
func (s *API) CreateOrganization(req *CreateOrganizationRequest, opts ...scw.RequestOption) (*Organization, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/partner/v1/organizations",
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

// GetOrganization: Get an organization.
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
		Path:   "/partner/v1/organizations/" + fmt.Sprint(req.OrganizationID) + "",
	}

	var resp Organization

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListOrganizations: List Organizations.
func (s *API) ListOrganizations(req *ListOrganizationsRequest, opts ...scw.RequestOption) (*ListOrganizationsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "email", req.Email)
	parameter.AddToQuery(query, "customer_id", req.CustomerID)
	parameter.AddToQuery(query, "locked_by", req.LockedBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/partner/v1/organizations",
		Query:  query,
	}

	var resp ListOrganizationsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// LockOrganization: Lock an organization.
func (s *API) LockOrganization(req *LockOrganizationRequest, opts ...scw.RequestOption) (*Organization, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/partner/v1/organizations/" + fmt.Sprint(req.OrganizationID) + "/lock",
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

// UnlockOrganization: Unlock an organization.
func (s *API) UnlockOrganization(req *UnlockOrganizationRequest, opts ...scw.RequestOption) (*Organization, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if fmt.Sprint(req.OrganizationID) == "" {
		return nil, errors.New("field OrganizationID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/partner/v1/organizations/" + fmt.Sprint(req.OrganizationID) + "/unlock",
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

// UpdateOrganization: Update an organization.
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
		Path:   "/partner/v1/organizations/" + fmt.Sprint(req.OrganizationID) + "",
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

// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package test provides methods and message types of the test v1 API.
package test

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
	"github.com/scaleway/scaleway-sdk-go/internal/async"
	"github.com/scaleway/scaleway-sdk-go/marshaler"
	"github.com/scaleway/scaleway-sdk-go/namegenerator"
	"github.com/scaleway/scaleway-sdk-go/parameter"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

const (
	defaultTestRetryInterval = 15 * time.Second
	defaultTestTimeout       = 5 * time.Minute
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

type EyeColors string

const (
	// Unknown color.
	EyeColorsUnknown = EyeColors("unknown")
	// Rare and striking shade that typically features a golden or yellowish-brown hue.
	EyeColorsAmber = EyeColors("amber")
	// Relatively rare, with the highest frequency found in eastern Europe.
	EyeColorsBlue = EyeColors("blue")
	// Most common eye color in the world caused by a high concentration of melanin in the iris.
	EyeColorsBrown = EyeColors("brown")
	// Relatively rare color which can change depending on the lighting conditions.
	EyeColorsGray = EyeColors("gray")
	// Rare and unique color characterized by a combination of yellow, brown, and blue pigments.
	EyeColorsGreen = EyeColors("green")
	// Brownish-yellow or greenish-brown with a hint of gold.
	EyeColorsHazel = EyeColors("hazel")
	// Rare mutation that results in a reddish-pink hue.
	EyeColorsRed = EyeColors("red")
	// Rare and striking shade that appears to be a mix of blue and purple.
	EyeColorsViolet = EyeColors("violet")
)

func (enum EyeColors) String() string {
	if enum == "" {
		// return default value if empty
		return string(EyeColorsUnknown)
	}
	return string(enum)
}

func (enum EyeColors) Values() []EyeColors {
	return []EyeColors{
		"unknown",
		"amber",
		"blue",
		"brown",
		"gray",
		"green",
		"hazel",
		"red",
		"violet",
	}
}

func (enum EyeColors) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *EyeColors) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = EyeColors(EyeColors(tmp).String())
	return nil
}

type HumanStatus string

const (
	// Unknown status.
	HumanStatusUnknown = HumanStatus("unknown")
	// The human is stopped.
	HumanStatusStopped = HumanStatus("stopped")
	// The human is running.
	HumanStatusRunning = HumanStatus("running")
)

func (enum HumanStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(HumanStatusUnknown)
	}
	return string(enum)
}

func (enum HumanStatus) Values() []HumanStatus {
	return []HumanStatus{
		"unknown",
		"stopped",
		"running",
	}
}

func (enum HumanStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *HumanStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = HumanStatus(HumanStatus(tmp).String())
	return nil
}

type ListHumansRequestOrderBy string

const (
	// Ascending creation date.
	ListHumansRequestOrderByCreatedAtAsc = ListHumansRequestOrderBy("created_at_asc")
	// Descending creation date.
	ListHumansRequestOrderByCreatedAtDesc = ListHumansRequestOrderBy("created_at_desc")
	// Ascending update date.
	ListHumansRequestOrderByUpdatedAtAsc = ListHumansRequestOrderBy("updated_at_asc")
	// Descending update date.
	ListHumansRequestOrderByUpdatedAtDesc = ListHumansRequestOrderBy("updated_at_desc")
	// Ascending height.
	ListHumansRequestOrderByHeightAsc = ListHumansRequestOrderBy("height_asc")
	// Descending height.
	ListHumansRequestOrderByHeightDesc = ListHumansRequestOrderBy("height_desc")
)

func (enum ListHumansRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListHumansRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListHumansRequestOrderBy) Values() []ListHumansRequestOrderBy {
	return []ListHumansRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
		"height_asc",
		"height_desc",
	}
}

func (enum ListHumansRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListHumansRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListHumansRequestOrderBy(ListHumansRequestOrderBy(tmp).String())
	return nil
}

// Human: human.
type Human struct {
	ID string `json:"id"`

	OrganizationID string `json:"organization_id"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	Height float64 `json:"height"`

	ShoeSize float32 `json:"shoe_size"`

	AltitudeInMeter int32 `json:"altitude_in_meter"`

	AltitudeInMillimeter int64 `json:"altitude_in_millimeter"`

	FingersCount uint32 `json:"fingers_count"`

	HairCount uint64 `json:"hair_count"`

	IsHappy bool `json:"is_happy"`

	// EyesColor: default value: unknown
	EyesColor EyeColors `json:"eyes_color"`

	// Status: default value: unknown
	Status HumanStatus `json:"status"`

	Name string `json:"name"`

	ProjectID string `json:"project_id"`
}

// CreateHumanRequest: create human request.
type CreateHumanRequest struct {
	Height float64 `json:"height"`

	ShoeSize float32 `json:"shoe_size"`

	AltitudeInMeter int32 `json:"altitude_in_meter"`

	AltitudeInMillimeter int64 `json:"altitude_in_millimeter"`

	FingersCount uint32 `json:"fingers_count"`

	HairCount uint64 `json:"hair_count"`

	IsHappy bool `json:"is_happy"`

	// EyesColor: default value: unknown
	EyesColor EyeColors `json:"eyes_color"`

	// Deprecated
	// Precisely one of ProjectID, OrganizationID must be set.
	OrganizationID *string `json:"organization_id,omitempty"`

	Name string `json:"name"`

	// Precisely one of ProjectID, OrganizationID must be set.
	ProjectID *string `json:"project_id,omitempty"`
}

// DeleteHumanRequest: delete human request.
type DeleteHumanRequest struct {
	// HumanID: UUID of the human you want to delete.
	HumanID string `json:"-"`
}

// GetHumanRequest: get human request.
type GetHumanRequest struct {
	// HumanID: UUID of the human you want to get.
	HumanID string `json:"-"`
}

// ListHumansRequest: list humans request.
type ListHumansRequest struct {
	Page *int32 `json:"-"`

	PageSize *uint32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListHumansRequestOrderBy `json:"-"`

	OrganizationID *string `json:"-"`

	ProjectID *string `json:"-"`
}

// ListHumansResponse: list humans response.
type ListHumansResponse struct {
	TotalCount uint32 `json:"total_count"`

	Humans []*Human `json:"humans"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListHumansResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListHumansResponse) UnsafeAppend(res any) (uint32, error) {
	results, ok := res.(*ListHumansResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Humans = append(r.Humans, results.Humans...)
	r.TotalCount += uint32(len(results.Humans))
	return uint32(len(results.Humans)), nil
}

// RegisterRequest: register request.
type RegisterRequest struct {
	Username string `json:"username"`
}

// RegisterResponse: register response.
type RegisterResponse struct {
	SecretKey string `json:"secret_key"`

	AccessKey string `json:"access_key"`
}

// RunHumanRequest: run human request.
type RunHumanRequest struct {
	// HumanID: UUID of the human you want to make run.
	HumanID string `json:"-"`
}

// SmokeHumanRequest: smoke human request.
type SmokeHumanRequest struct {
	// Deprecated: HumanID: UUID of the human you want to make smoking.
	HumanID string `json:"-"`
}

// UpdateHumanRequest: update human request.
type UpdateHumanRequest struct {
	// HumanID: UUID of the human you want to update.
	HumanID string `json:"-"`

	// Height: height of the human in meters.
	Height *float64 `json:"height,omitempty"`

	ShoeSize *float32 `json:"shoe_size,omitempty"`

	AltitudeInMeter *int32 `json:"altitude_in_meter,omitempty"`

	AltitudeInMillimeter *int64 `json:"altitude_in_millimeter,omitempty"`

	FingersCount *uint32 `json:"fingers_count,omitempty"`

	HairCount *uint64 `json:"hair_count,omitempty"`

	IsHappy *bool `json:"is_happy,omitempty"`

	// EyesColor: default value: unknown
	EyesColor EyeColors `json:"eyes_color"`

	Name *string `json:"name,omitempty"`
}

// No Auth Service for end-to-end testing.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

// Register: Register a human and return a access-key and a secret-key that must be used in all other commands.
//
// Hint: you can use other test commands by setting the SCW_SECRET_KEY env variable.
func (s *API) Register(req *RegisterRequest, opts ...scw.RequestOption) (*RegisterResponse, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/test/v1/register",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp RegisterResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListHumans: List all your humans.
func (s *API) ListHumans(req *ListHumansRequest, opts ...scw.RequestOption) (*ListHumansResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/test/v1/humans",
		Query:  query,
	}

	var resp ListHumansResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetHuman: Get the human details associated with the given id.
func (s *API) GetHuman(req *GetHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/test/v1/humans/" + fmt.Sprint(req.HumanID) + "",
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForHumanRequest is used by WaitForHuman method.
type WaitForHumanRequest struct {
	GetHumanRequest
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForHuman waits for the Human to reach a terminal state.
func (s *API) WaitForHuman(req *WaitForHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	timeout := defaultTestTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultTestRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[HumanStatus]struct{}{
		HumanStatusRunning: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := s.GetHuman(&GetHumanRequest{
				HumanID: req.HumanID,
			}, opts...)
			if err != nil {
				return nil, false, err
			}

			_, isTransient := transientStatuses[res.Status]

			return res, !isTransient, nil
		},
		IntervalStrategy: async.LinearIntervalStrategy(retryInterval),
		Timeout:          timeout,
	})
	if err != nil {
		return nil, errors.Wrap(err, "waiting for Human failed")
	}

	return res.(*Human), nil
}

// CreateHuman: Create a new human.
func (s *API) CreateHuman(req *CreateHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	defaultProjectID, exist := s.client.GetDefaultProjectID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.ProjectID = &defaultProjectID
	}

	defaultOrganizationID, exist := s.client.GetDefaultOrganizationID()
	if exist && req.ProjectID == nil && req.OrganizationID == nil {
		req.OrganizationID = &defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/test/v1/humans",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateHuman: Update the human associated with the given id.
func (s *API) UpdateHuman(req *UpdateHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/test/v1/humans/" + fmt.Sprint(req.HumanID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteHuman: Delete the human associated with the given id.
func (s *API) DeleteHuman(req *DeleteHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/test/v1/humans/" + fmt.Sprint(req.HumanID) + "",
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RunHuman: Start a one hour running for the given human.
func (s *API) RunHuman(req *RunHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/test/v1/humans/" + fmt.Sprint(req.HumanID) + "/run",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// Deprecated: SmokeHuman: Make a human smoke.
func (s *API) SmokeHuman(req *SmokeHumanRequest, opts ...scw.RequestOption) (*Human, error) {
	var err error

	if fmt.Sprint(req.HumanID) == "" {
		return nil, errors.New("field HumanID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/test/v1/humans/" + fmt.Sprint(req.HumanID) + "/smoke",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Human

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

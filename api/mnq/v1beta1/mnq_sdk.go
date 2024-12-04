// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package mnq provides methods and message types of the mnq v1beta1 API.
package mnq

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

type ListNatsAccountsRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListNatsAccountsRequestOrderByCreatedAtAsc = ListNatsAccountsRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListNatsAccountsRequestOrderByCreatedAtDesc = ListNatsAccountsRequestOrderBy("created_at_desc")
	// Order by last update date (ascending chronological order).
	ListNatsAccountsRequestOrderByUpdatedAtAsc = ListNatsAccountsRequestOrderBy("updated_at_asc")
	// Order by last update date (descending chronological order).
	ListNatsAccountsRequestOrderByUpdatedAtDesc = ListNatsAccountsRequestOrderBy("updated_at_desc")
	// Order by name (ascending alphabetical order).
	ListNatsAccountsRequestOrderByNameAsc = ListNatsAccountsRequestOrderBy("name_asc")
	// Order by name (descending alphabetical order).
	ListNatsAccountsRequestOrderByNameDesc = ListNatsAccountsRequestOrderBy("name_desc")
)

func (enum ListNatsAccountsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNatsAccountsRequestOrderBy) Values() []ListNatsAccountsRequestOrderBy {
	return []ListNatsAccountsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListNatsAccountsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNatsAccountsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNatsAccountsRequestOrderBy(ListNatsAccountsRequestOrderBy(tmp).String())
	return nil
}

type ListNatsCredentialsRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListNatsCredentialsRequestOrderByCreatedAtAsc = ListNatsCredentialsRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListNatsCredentialsRequestOrderByCreatedAtDesc = ListNatsCredentialsRequestOrderBy("created_at_desc")
	// Order by last update date (ascending chronological order).
	ListNatsCredentialsRequestOrderByUpdatedAtAsc = ListNatsCredentialsRequestOrderBy("updated_at_asc")
	// Order by last update date (descending chronological order).
	ListNatsCredentialsRequestOrderByUpdatedAtDesc = ListNatsCredentialsRequestOrderBy("updated_at_desc")
	// Order by name (ascending alphabetical order).
	ListNatsCredentialsRequestOrderByNameAsc = ListNatsCredentialsRequestOrderBy("name_asc")
	// Order by name (descending alphabetical order).
	ListNatsCredentialsRequestOrderByNameDesc = ListNatsCredentialsRequestOrderBy("name_desc")
)

func (enum ListNatsCredentialsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListNatsCredentialsRequestOrderBy) Values() []ListNatsCredentialsRequestOrderBy {
	return []ListNatsCredentialsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListNatsCredentialsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListNatsCredentialsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListNatsCredentialsRequestOrderBy(ListNatsCredentialsRequestOrderBy(tmp).String())
	return nil
}

type ListSnsCredentialsRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListSnsCredentialsRequestOrderByCreatedAtAsc = ListSnsCredentialsRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListSnsCredentialsRequestOrderByCreatedAtDesc = ListSnsCredentialsRequestOrderBy("created_at_desc")
	// Order by last update date (ascending chronological order).
	ListSnsCredentialsRequestOrderByUpdatedAtAsc = ListSnsCredentialsRequestOrderBy("updated_at_asc")
	// Order by last update date (descending chronological order).
	ListSnsCredentialsRequestOrderByUpdatedAtDesc = ListSnsCredentialsRequestOrderBy("updated_at_desc")
	// Order by name (ascending alphabetical order).
	ListSnsCredentialsRequestOrderByNameAsc = ListSnsCredentialsRequestOrderBy("name_asc")
	// Order by name (descending alphabetical order).
	ListSnsCredentialsRequestOrderByNameDesc = ListSnsCredentialsRequestOrderBy("name_desc")
)

func (enum ListSnsCredentialsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSnsCredentialsRequestOrderBy) Values() []ListSnsCredentialsRequestOrderBy {
	return []ListSnsCredentialsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListSnsCredentialsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSnsCredentialsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSnsCredentialsRequestOrderBy(ListSnsCredentialsRequestOrderBy(tmp).String())
	return nil
}

type ListSqsCredentialsRequestOrderBy string

const (
	// Order by creation date (ascending chronological order).
	ListSqsCredentialsRequestOrderByCreatedAtAsc = ListSqsCredentialsRequestOrderBy("created_at_asc")
	// Order by creation date (descending chronological order).
	ListSqsCredentialsRequestOrderByCreatedAtDesc = ListSqsCredentialsRequestOrderBy("created_at_desc")
	// Order by last update date (ascending chronological order).
	ListSqsCredentialsRequestOrderByUpdatedAtAsc = ListSqsCredentialsRequestOrderBy("updated_at_asc")
	// Order by last update date (descending chronological order).
	ListSqsCredentialsRequestOrderByUpdatedAtDesc = ListSqsCredentialsRequestOrderBy("updated_at_desc")
	// Order by name (ascending alphabetical order).
	ListSqsCredentialsRequestOrderByNameAsc = ListSqsCredentialsRequestOrderBy("name_asc")
	// Order by name (descending alphabetical order).
	ListSqsCredentialsRequestOrderByNameDesc = ListSqsCredentialsRequestOrderBy("name_desc")
)

func (enum ListSqsCredentialsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListSqsCredentialsRequestOrderBy) Values() []ListSqsCredentialsRequestOrderBy {
	return []ListSqsCredentialsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"updated_at_asc",
		"updated_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListSqsCredentialsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListSqsCredentialsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListSqsCredentialsRequestOrderBy(ListSqsCredentialsRequestOrderBy(tmp).String())
	return nil
}

type SnsInfoStatus string

const (
	// Unknown status.
	SnsInfoStatusUnknownStatus = SnsInfoStatus("unknown_status")
	// Enabled status.
	SnsInfoStatusEnabled = SnsInfoStatus("enabled")
	// Disabled status.
	SnsInfoStatusDisabled = SnsInfoStatus("disabled")
)

func (enum SnsInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum SnsInfoStatus) Values() []SnsInfoStatus {
	return []SnsInfoStatus{
		"unknown_status",
		"enabled",
		"disabled",
	}
}

func (enum SnsInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SnsInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SnsInfoStatus(SnsInfoStatus(tmp).String())
	return nil
}

type SqsInfoStatus string

const (
	// Unknown status.
	SqsInfoStatusUnknownStatus = SqsInfoStatus("unknown_status")
	// Enabled status.
	SqsInfoStatusEnabled = SqsInfoStatus("enabled")
	// Disabled status.
	SqsInfoStatusDisabled = SqsInfoStatus("disabled")
)

func (enum SqsInfoStatus) String() string {
	if enum == "" {
		// return default value if empty
		return "unknown_status"
	}
	return string(enum)
}

func (enum SqsInfoStatus) Values() []SqsInfoStatus {
	return []SqsInfoStatus{
		"unknown_status",
		"enabled",
		"disabled",
	}
}

func (enum SqsInfoStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *SqsInfoStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = SqsInfoStatus(SqsInfoStatus(tmp).String())
	return nil
}

// File: file.
type File struct {
	// Name: file name.
	Name string `json:"name"`

	// Content: file content.
	Content string `json:"content"`
}

// SnsPermissions: sns permissions.
type SnsPermissions struct {
	// CanPublish: defines whether the credentials bearer can publish messages to the service (publish to SNS topics).
	CanPublish *bool `json:"can_publish"`

	// CanReceive: defines whether the credentials bearer can receive messages from the service (configure subscriptions).
	CanReceive *bool `json:"can_receive"`

	// CanManage: defines whether the credentials bearer can manage the associated SNS topics or subscriptions.
	CanManage *bool `json:"can_manage"`
}

// SqsPermissions: sqs permissions.
type SqsPermissions struct {
	// CanPublish: defines whether the credentials bearer can publish messages to the service (send messages to SQS queues).
	CanPublish *bool `json:"can_publish"`

	// CanReceive: defines whether the credentials bearer can receive messages from SQS queues.
	CanReceive *bool `json:"can_receive"`

	// CanManage: defines whether the credentials bearer can manage the associated SQS queues.
	CanManage *bool `json:"can_manage"`
}

// NatsAccount: nats account.
type NatsAccount struct {
	// ID: nATS account ID.
	ID string `json:"id"`

	// Name: nATS account name.
	Name string `json:"name"`

	// Endpoint: endpoint of the NATS service for this account.
	Endpoint string `json:"endpoint"`

	// ProjectID: project ID of the Project containing the NATS account.
	ProjectID string `json:"project_id"`

	// Region: region where the NATS account is deployed.
	Region scw.Region `json:"region"`

	// CreatedAt: nATS account creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: nATS account last modification date.
	UpdatedAt *time.Time `json:"updated_at"`
}

// NatsCredentials: nats credentials.
type NatsCredentials struct {
	// ID: ID of the credentials.
	ID string `json:"id"`

	// Name: name of the credentials.
	Name string `json:"name"`

	// NatsAccountID: nATS account containing the credentials.
	NatsAccountID string `json:"nats_account_id"`

	// CreatedAt: nATS credentials creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: nATS credentials last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Credentials: object containing the credentials file (Only returned by **Create Nats Credentials** call).
	Credentials *File `json:"credentials"`

	// Checksum: checksum of the credentials file.
	Checksum string `json:"checksum"`
}

// SnsCredentials: sns credentials.
type SnsCredentials struct {
	// ID: ID of the credentials.
	ID string `json:"id"`

	// Name: name of the credentials.
	Name string `json:"name"`

	// ProjectID: project ID of the Project containing the credentials.
	ProjectID string `json:"project_id"`

	// Region: region where the credentials exists.
	Region scw.Region `json:"region"`

	// CreatedAt: credentials creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: credentials last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// AccessKey: access key ID.
	AccessKey string `json:"access_key"`

	// SecretKey: secret key ID (Only returned by **Create SNS Credentials** call).
	SecretKey string `json:"secret_key"`

	// SecretChecksum: checksum of the Secret key.
	SecretChecksum string `json:"secret_checksum"`

	// Permissions: permissions associated with these credentials.
	Permissions *SnsPermissions `json:"permissions"`
}

// SqsCredentials: sqs credentials.
type SqsCredentials struct {
	// ID: ID of the credentials.
	ID string `json:"id"`

	// Name: name of the credentials.
	Name string `json:"name"`

	// ProjectID: project ID of the Project containing the credentials.
	ProjectID string `json:"project_id"`

	// Region: region where the credentials exists.
	Region scw.Region `json:"region"`

	// CreatedAt: credentials creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: credentials last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// AccessKey: access key ID.
	AccessKey string `json:"access_key"`

	// SecretKey: secret key ID (Only returned by **Create SQS Credentials** call).
	SecretKey string `json:"secret_key"`

	// SecretChecksum: checksum of the Secret key.
	SecretChecksum string `json:"secret_checksum"`

	// Permissions: permissions associated with these credentials.
	Permissions *SqsPermissions `json:"permissions"`
}

// ListNatsAccountsResponse: list nats accounts response.
type ListNatsAccountsResponse struct {
	// TotalCount: total count of existing NATS accounts (matching any filters specified).
	TotalCount uint64 `json:"total_count"`

	// NatsAccounts: nATS accounts on this page.
	NatsAccounts []*NatsAccount `json:"nats_accounts"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNatsAccountsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNatsAccountsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListNatsAccountsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.NatsAccounts = append(r.NatsAccounts, results.NatsAccounts...)
	r.TotalCount += uint64(len(results.NatsAccounts))
	return uint64(len(results.NatsAccounts)), nil
}

// ListNatsCredentialsResponse: list nats credentials response.
type ListNatsCredentialsResponse struct {
	// TotalCount: total count of existing credentials (matching any filters specified).
	TotalCount uint64 `json:"total_count"`

	// NatsCredentials: credentials on this page.
	NatsCredentials []*NatsCredentials `json:"nats_credentials"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListNatsCredentialsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListNatsCredentialsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListNatsCredentialsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.NatsCredentials = append(r.NatsCredentials, results.NatsCredentials...)
	r.TotalCount += uint64(len(results.NatsCredentials))
	return uint64(len(results.NatsCredentials)), nil
}

// ListSnsCredentialsResponse: list sns credentials response.
type ListSnsCredentialsResponse struct {
	// TotalCount: total count of existing credentials (matching any filters specified).
	TotalCount uint64 `json:"total_count"`

	// SnsCredentials: sNS credentials on this page.
	SnsCredentials []*SnsCredentials `json:"sns_credentials"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSnsCredentialsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSnsCredentialsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListSnsCredentialsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SnsCredentials = append(r.SnsCredentials, results.SnsCredentials...)
	r.TotalCount += uint64(len(results.SnsCredentials))
	return uint64(len(results.SnsCredentials)), nil
}

// ListSqsCredentialsResponse: list sqs credentials response.
type ListSqsCredentialsResponse struct {
	// TotalCount: total count of existing credentials (matching any filters specified).
	TotalCount uint64 `json:"total_count"`

	// SqsCredentials: sQS credentials on this page.
	SqsCredentials []*SqsCredentials `json:"sqs_credentials"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListSqsCredentialsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListSqsCredentialsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListSqsCredentialsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.SqsCredentials = append(r.SqsCredentials, results.SqsCredentials...)
	r.TotalCount += uint64(len(results.SqsCredentials))
	return uint64(len(results.SqsCredentials)), nil
}

// NatsAPICreateNatsAccountRequest: nats api create nats account request.
type NatsAPICreateNatsAccountRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: nATS account name.
	Name string `json:"name"`

	// ProjectID: project containing the NATS account.
	ProjectID string `json:"project_id"`
}

// NatsAPICreateNatsCredentialsRequest: nats api create nats credentials request.
type NatsAPICreateNatsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NatsAccountID: nATS account containing the credentials.
	NatsAccountID string `json:"nats_account_id"`

	// Name: name of the credentials.
	Name string `json:"name"`
}

// NatsAPIDeleteNatsAccountRequest: nats api delete nats account request.
type NatsAPIDeleteNatsAccountRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NatsAccountID: ID of the NATS account to delete.
	NatsAccountID string `json:"-"`
}

// NatsAPIDeleteNatsCredentialsRequest: nats api delete nats credentials request.
type NatsAPIDeleteNatsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NatsCredentialsID: ID of the credentials to delete.
	NatsCredentialsID string `json:"-"`
}

// NatsAPIGetNatsAccountRequest: nats api get nats account request.
type NatsAPIGetNatsAccountRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NatsAccountID: ID of the NATS account to get.
	NatsAccountID string `json:"-"`
}

// NatsAPIGetNatsCredentialsRequest: nats api get nats credentials request.
type NatsAPIGetNatsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NatsCredentialsID: ID of the credentials to get.
	NatsCredentialsID string `json:"-"`
}

// NatsAPIListNatsAccountsRequest: nats api list nats accounts request.
type NatsAPIListNatsAccountsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: include only NATS accounts in this Project.
	ProjectID *string `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of NATS accounts to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListNatsAccountsRequestOrderBy `json:"-"`
}

// NatsAPIListNatsCredentialsRequest: nats api list nats credentials request.
type NatsAPIListNatsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: include only NATS accounts in this Project.
	ProjectID *string `json:"-"`

	// NatsAccountID: include only credentials for this NATS account.
	NatsAccountID *string `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of credentials to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListNatsCredentialsRequestOrderBy `json:"-"`
}

// NatsAPIUpdateNatsAccountRequest: nats api update nats account request.
type NatsAPIUpdateNatsAccountRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// NatsAccountID: ID of the NATS account to update.
	NatsAccountID string `json:"-"`

	// Name: nATS account name.
	Name *string `json:"name,omitempty"`
}

// SnsAPIActivateSnsRequest: sns api activate sns request.
type SnsAPIActivateSnsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project on which to activate the SNS service.
	ProjectID string `json:"project_id"`
}

// SnsAPICreateSnsCredentialsRequest: sns api create sns credentials request.
type SnsAPICreateSnsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project containing the SNS credentials.
	ProjectID string `json:"project_id"`

	// Name: name of the credentials.
	Name string `json:"name"`

	// Permissions: permissions associated with these credentials.
	Permissions *SnsPermissions `json:"permissions,omitempty"`
}

// SnsAPIDeactivateSnsRequest: sns api deactivate sns request.
type SnsAPIDeactivateSnsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project on which to deactivate the SNS service.
	ProjectID string `json:"project_id"`
}

// SnsAPIDeleteSnsCredentialsRequest: sns api delete sns credentials request.
type SnsAPIDeleteSnsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SnsCredentialsID: ID of the credentials to delete.
	SnsCredentialsID string `json:"-"`
}

// SnsAPIGetSnsCredentialsRequest: sns api get sns credentials request.
type SnsAPIGetSnsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SnsCredentialsID: ID of the SNS credentials to get.
	SnsCredentialsID string `json:"-"`
}

// SnsAPIGetSnsInfoRequest: sns api get sns info request.
type SnsAPIGetSnsInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project to retrieve SNS info from.
	ProjectID string `json:"project_id"`
}

// SnsAPIListSnsCredentialsRequest: sns api list sns credentials request.
type SnsAPIListSnsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: include only SNS credentials in this Project.
	ProjectID *string `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of credentials to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListSnsCredentialsRequestOrderBy `json:"-"`
}

// SnsAPIUpdateSnsCredentialsRequest: sns api update sns credentials request.
type SnsAPIUpdateSnsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SnsCredentialsID: ID of the SNS credentials to update.
	SnsCredentialsID string `json:"-"`

	// Name: name of the credentials.
	Name *string `json:"name,omitempty"`

	// Permissions: permissions associated with these credentials.
	Permissions *SnsPermissions `json:"permissions,omitempty"`
}

// SnsInfo: sns info.
type SnsInfo struct {
	// ProjectID: project ID of the Project containing the service.
	ProjectID string `json:"project_id"`

	// Region: region of the service.
	Region scw.Region `json:"region"`

	// CreatedAt: sNS creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: sNS last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Status: sNS activation status.
	// Default value: unknown_status
	Status SnsInfoStatus `json:"status"`

	// SnsEndpointURL: endpoint of the SNS service for this region and project.
	SnsEndpointURL string `json:"sns_endpoint_url"`
}

// SqsAPIActivateSqsRequest: sqs api activate sqs request.
type SqsAPIActivateSqsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project on which to activate the SQS service.
	ProjectID string `json:"project_id"`
}

// SqsAPICreateSqsCredentialsRequest: sqs api create sqs credentials request.
type SqsAPICreateSqsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project containing the SQS credentials.
	ProjectID string `json:"project_id"`

	// Name: name of the credentials.
	Name string `json:"name"`

	// Permissions: permissions associated with these credentials.
	Permissions *SqsPermissions `json:"permissions,omitempty"`
}

// SqsAPIDeactivateSqsRequest: sqs api deactivate sqs request.
type SqsAPIDeactivateSqsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project on which to deactivate the SQS service.
	ProjectID string `json:"project_id"`
}

// SqsAPIDeleteSqsCredentialsRequest: sqs api delete sqs credentials request.
type SqsAPIDeleteSqsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SqsCredentialsID: ID of the credentials to delete.
	SqsCredentialsID string `json:"-"`
}

// SqsAPIGetSqsCredentialsRequest: sqs api get sqs credentials request.
type SqsAPIGetSqsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SqsCredentialsID: ID of the SQS credentials to get.
	SqsCredentialsID string `json:"-"`
}

// SqsAPIGetSqsInfoRequest: sqs api get sqs info request.
type SqsAPIGetSqsInfoRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: project to retrieve SQS info from.
	ProjectID string `json:"project_id"`
}

// SqsAPIListSqsCredentialsRequest: sqs api list sqs credentials request.
type SqsAPIListSqsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: include only SQS credentials in this Project.
	ProjectID *string `json:"-"`

	// Page: page number to return.
	Page *int32 `json:"-"`

	// PageSize: maximum number of credentials to return per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: order in which to return results.
	// Default value: created_at_asc
	OrderBy ListSqsCredentialsRequestOrderBy `json:"-"`
}

// SqsAPIUpdateSqsCredentialsRequest: sqs api update sqs credentials request.
type SqsAPIUpdateSqsCredentialsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// SqsCredentialsID: ID of the SQS credentials to update.
	SqsCredentialsID string `json:"-"`

	// Name: name of the credentials.
	Name *string `json:"name,omitempty"`

	// Permissions: permissions associated with these credentials.
	Permissions *SqsPermissions `json:"permissions,omitempty"`
}

// SqsInfo: sqs info.
type SqsInfo struct {
	// ProjectID: project ID of the Project containing the service.
	ProjectID string `json:"project_id"`

	// Region: region of the service.
	Region scw.Region `json:"region"`

	// CreatedAt: sQS creation date.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: sQS last modification date.
	UpdatedAt *time.Time `json:"updated_at"`

	// Status: sQS activation status.
	// Default value: unknown_status
	Status SqsInfoStatus `json:"status"`

	// SqsEndpointURL: endpoint of the SQS service for this region and project.
	SqsEndpointURL string `json:"sqs_endpoint_url"`
}

// This API allows you to manage Scaleway Messaging and Queuing NATS accounts.
type NatsAPI struct {
	client *scw.Client
}

// NewNatsAPI returns a NatsAPI object from a Scaleway client.
func NewNatsAPI(client *scw.Client) *NatsAPI {
	return &NatsAPI{
		client: client,
	}
}
func (s *NatsAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms}
}

// CreateNatsAccount: Create a NATS account associated with a Project.
func (s *NatsAPI) CreateNatsAccount(req *NatsAPICreateNatsAccountRequest, opts ...scw.RequestOption) (*NatsAccount, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("mnq")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-accounts",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp NatsAccount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNatsAccount: Delete a NATS account, specified by its NATS account ID. Note that deleting a NATS account is irreversible, and any credentials, streams, consumer and stored messages belonging to this NATS account will also be deleted.
func (s *NatsAPI) DeleteNatsAccount(req *NatsAPIDeleteNatsAccountRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NatsAccountID) == "" {
		return errors.New("field NatsAccountID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-accounts/" + fmt.Sprint(req.NatsAccountID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateNatsAccount: Update the name of a NATS account, specified by its NATS account ID.
func (s *NatsAPI) UpdateNatsAccount(req *NatsAPIUpdateNatsAccountRequest, opts ...scw.RequestOption) (*NatsAccount, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NatsAccountID) == "" {
		return nil, errors.New("field NatsAccountID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-accounts/" + fmt.Sprint(req.NatsAccountID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp NatsAccount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetNatsAccount: Retrieve information about an existing NATS account identified by its NATS account ID. Its full details, including name and endpoint, are returned in the response.
func (s *NatsAPI) GetNatsAccount(req *NatsAPIGetNatsAccountRequest, opts ...scw.RequestOption) (*NatsAccount, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NatsAccountID) == "" {
		return nil, errors.New("field NatsAccountID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-accounts/" + fmt.Sprint(req.NatsAccountID) + "",
	}

	var resp NatsAccount

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNatsAccounts: List all NATS accounts in the specified region, for a Scaleway Organization or Project. By default, the NATS accounts returned in the list are ordered by creation date in ascending order, though this can be modified via the `order_by` field.
func (s *NatsAPI) ListNatsAccounts(req *NatsAPIListNatsAccountsRequest, opts ...scw.RequestOption) (*ListNatsAccountsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-accounts",
		Query:  query,
	}

	var resp ListNatsAccountsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateNatsCredentials: Create a set of credentials for a NATS account, specified by its NATS account ID.
func (s *NatsAPI) CreateNatsCredentials(req *NatsAPICreateNatsCredentialsRequest, opts ...scw.RequestOption) (*NatsCredentials, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("mnq")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-credentials",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp NatsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteNatsCredentials: Delete a set of credentials, specified by their credentials ID. Deleting credentials is irreversible and cannot be undone. The credentials can no longer be used to access the NATS account, and active connections using this credentials will be closed.
func (s *NatsAPI) DeleteNatsCredentials(req *NatsAPIDeleteNatsCredentialsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NatsCredentialsID) == "" {
		return errors.New("field NatsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-credentials/" + fmt.Sprint(req.NatsCredentialsID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// GetNatsCredentials: Retrieve an existing set of credentials, identified by the `nats_credentials_id`. The credentials themselves are NOT returned, only their metadata (NATS account ID, credentials name, etc), are returned in the response.
func (s *NatsAPI) GetNatsCredentials(req *NatsAPIGetNatsCredentialsRequest, opts ...scw.RequestOption) (*NatsCredentials, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.NatsCredentialsID) == "" {
		return nil, errors.New("field NatsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-credentials/" + fmt.Sprint(req.NatsCredentialsID) + "",
	}

	var resp NatsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListNatsCredentials: List existing credentials in the specified NATS account. The response contains only the metadata for the credentials, not the credentials themselves, which are only returned after a **Create Credentials** call.
func (s *NatsAPI) ListNatsCredentials(req *NatsAPIListNatsCredentialsRequest, opts ...scw.RequestOption) (*ListNatsCredentialsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "nats_account_id", req.NatsAccountID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/nats-credentials",
		Query:  query,
	}

	var resp ListNatsCredentialsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage your Scaleway Messaging and Queuing SNS brokers.
type SnsAPI struct {
	client *scw.Client
}

// NewSnsAPI returns a SnsAPI object from a Scaleway client.
func NewSnsAPI(client *scw.Client) *SnsAPI {
	return &SnsAPI{
		client: client,
	}
}
func (s *SnsAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms}
}

// ActivateSns: Activate SNS for the specified Project ID. SNS must be activated before any usage. Activating SNS does not trigger any billing, and you can deactivate at any time.
func (s *SnsAPI) ActivateSns(req *SnsAPIActivateSnsRequest, opts ...scw.RequestOption) (*SnsInfo, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/activate-sns",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SnsInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSnsInfo: Retrieve the SNS information of the specified Project ID. Informations include the activation status and the SNS API endpoint URL.
func (s *SnsAPI) GetSnsInfo(req *SnsAPIGetSnsInfoRequest, opts ...scw.RequestOption) (*SnsInfo, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sns-info",
		Query:  query,
	}

	var resp SnsInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeactivateSns: Deactivate SNS for the specified Project ID.You must delete all topics and credentials before this call or you need to set the force_delete parameter.
func (s *SnsAPI) DeactivateSns(req *SnsAPIDeactivateSnsRequest, opts ...scw.RequestOption) (*SnsInfo, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deactivate-sns",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SnsInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSnsCredentials: Create a set of credentials for SNS, specified by a Project ID. Credentials give the bearer access to topics, and the level of permissions can be defined granularly.
func (s *SnsAPI) CreateSnsCredentials(req *SnsAPICreateSnsCredentialsRequest, opts ...scw.RequestOption) (*SnsCredentials, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("mnq_sns")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sns-credentials",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SnsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSnsCredentials: Delete a set of SNS credentials, specified by their credentials ID. Deleting credentials is irreversible and cannot be undone. The credentials can then no longer be used to access SNS.
func (s *SnsAPI) DeleteSnsCredentials(req *SnsAPIDeleteSnsCredentialsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SnsCredentialsID) == "" {
		return errors.New("field SnsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sns-credentials/" + fmt.Sprint(req.SnsCredentialsID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSnsCredentials: Update a set of SNS credentials. You can update the credentials' name, or their permissions.
func (s *SnsAPI) UpdateSnsCredentials(req *SnsAPIUpdateSnsCredentialsRequest, opts ...scw.RequestOption) (*SnsCredentials, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SnsCredentialsID) == "" {
		return nil, errors.New("field SnsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sns-credentials/" + fmt.Sprint(req.SnsCredentialsID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SnsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSnsCredentials: Retrieve an existing set of credentials, identified by the `credentials_id`. The credentials themselves, as well as their metadata (name, project ID etc), are returned in the response.
func (s *SnsAPI) GetSnsCredentials(req *SnsAPIGetSnsCredentialsRequest, opts ...scw.RequestOption) (*SnsCredentials, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SnsCredentialsID) == "" {
		return nil, errors.New("field SnsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sns-credentials/" + fmt.Sprint(req.SnsCredentialsID) + "",
	}

	var resp SnsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSnsCredentials: List existing SNS credentials in the specified region. The response contains only the metadata for the credentials, not the credentials themselves.
func (s *SnsAPI) ListSnsCredentials(req *SnsAPIListSnsCredentialsRequest, opts ...scw.RequestOption) (*ListSnsCredentialsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sns-credentials",
		Query:  query,
	}

	var resp ListSnsCredentialsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage your Scaleway Messaging and Queuing SQS brokers.
type SqsAPI struct {
	client *scw.Client
}

// NewSqsAPI returns a SqsAPI object from a Scaleway client.
func NewSqsAPI(client *scw.Client) *SqsAPI {
	return &SqsAPI{
		client: client,
	}
}
func (s *SqsAPI) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar, scw.RegionNlAms}
}

// ActivateSqs: Activate SQS for the specified Project ID. SQS must be activated before any usage such as creating credentials and queues. Activating SQS does not trigger any billing, and you can deactivate at any time.
func (s *SqsAPI) ActivateSqs(req *SqsAPIActivateSqsRequest, opts ...scw.RequestOption) (*SqsInfo, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/activate-sqs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SqsInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSqsInfo: Retrieve the SQS information of the specified Project ID. Informations include the activation status and the SQS API endpoint URL.
func (s *SqsAPI) GetSqsInfo(req *SqsAPIGetSqsInfoRequest, opts ...scw.RequestOption) (*SqsInfo, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sqs-info",
		Query:  query,
	}

	var resp SqsInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeactivateSqs: Deactivate SQS for the specified Project ID. You must delete all queues and credentials before this call or you need to set the force_delete parameter.
func (s *SqsAPI) DeactivateSqs(req *SqsAPIDeactivateSqsRequest, opts ...scw.RequestOption) (*SqsInfo, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/deactivate-sqs",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SqsInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateSqsCredentials: Create a set of credentials for SQS, specified by a Project ID. Credentials give the bearer access to queues, and the level of permissions can be defined granularly.
func (s *SqsAPI) CreateSqsCredentials(req *SqsAPICreateSqsCredentialsRequest, opts ...scw.RequestOption) (*SqsCredentials, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("mnq_sqs")
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sqs-credentials",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SqsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteSqsCredentials: Delete a set of SQS credentials, specified by their credentials ID. Deleting credentials is irreversible and cannot be undone. The credentials can then no longer be used to access SQS.
func (s *SqsAPI) DeleteSqsCredentials(req *SqsAPIDeleteSqsCredentialsRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SqsCredentialsID) == "" {
		return errors.New("field SqsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sqs-credentials/" + fmt.Sprint(req.SqsCredentialsID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateSqsCredentials: Update a set of SQS credentials. You can update the credentials' name, or their permissions.
func (s *SqsAPI) UpdateSqsCredentials(req *SqsAPIUpdateSqsCredentialsRequest, opts ...scw.RequestOption) (*SqsCredentials, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SqsCredentialsID) == "" {
		return nil, errors.New("field SqsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sqs-credentials/" + fmt.Sprint(req.SqsCredentialsID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp SqsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetSqsCredentials: Retrieve an existing set of credentials, identified by the `credentials_id`. The credentials themselves, as well as their metadata (name, project ID etc), are returned in the response.
func (s *SqsAPI) GetSqsCredentials(req *SqsAPIGetSqsCredentialsRequest, opts ...scw.RequestOption) (*SqsCredentials, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.SqsCredentialsID) == "" {
		return nil, errors.New("field SqsCredentialsID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sqs-credentials/" + fmt.Sprint(req.SqsCredentialsID) + "",
	}

	var resp SqsCredentials

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListSqsCredentials: List existing SQS credentials in the specified region. The response contains only the metadata for the credentials, not the credentials themselves.
func (s *SqsAPI) ListSqsCredentials(req *SqsAPIListSqsCredentialsRequest, opts ...scw.RequestOption) (*ListSqsCredentialsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/mnq/v1beta1/regions/" + fmt.Sprint(req.Region) + "/sqs-credentials",
		Query:  query,
	}

	var resp ListSqsCredentialsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

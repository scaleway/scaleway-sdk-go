// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package serverless_sqldb provides methods and message types of the serverless_sqldb v1alpha1 API.
package serverless_sqldb

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
	defaultServerlessSqldbRetryInterval = 15 * time.Second
	defaultServerlessSqldbTimeout       = 15 * time.Minute
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

type DatabaseBackupStatus string

const (
	DatabaseBackupStatusUnknownStatus = DatabaseBackupStatus("unknown_status")
	DatabaseBackupStatusError         = DatabaseBackupStatus("error")
	DatabaseBackupStatusReady         = DatabaseBackupStatus("ready")
	DatabaseBackupStatusLocked        = DatabaseBackupStatus("locked")
)

func (enum DatabaseBackupStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(DatabaseBackupStatusUnknownStatus)
	}
	return string(enum)
}

func (enum DatabaseBackupStatus) Values() []DatabaseBackupStatus {
	return []DatabaseBackupStatus{
		"unknown_status",
		"error",
		"ready",
		"locked",
	}
}

func (enum DatabaseBackupStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DatabaseBackupStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DatabaseBackupStatus(DatabaseBackupStatus(tmp).String())
	return nil
}

type DatabaseStatus string

const (
	DatabaseStatusUnknownStatus = DatabaseStatus("unknown_status")
	DatabaseStatusError         = DatabaseStatus("error")
	DatabaseStatusCreating      = DatabaseStatus("creating")
	DatabaseStatusReady         = DatabaseStatus("ready")
	DatabaseStatusDeleting      = DatabaseStatus("deleting")
	DatabaseStatusRestoring     = DatabaseStatus("restoring")
	DatabaseStatusLocked        = DatabaseStatus("locked")
)

func (enum DatabaseStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(DatabaseStatusUnknownStatus)
	}
	return string(enum)
}

func (enum DatabaseStatus) Values() []DatabaseStatus {
	return []DatabaseStatus{
		"unknown_status",
		"error",
		"creating",
		"ready",
		"deleting",
		"restoring",
		"locked",
	}
}

func (enum DatabaseStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *DatabaseStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = DatabaseStatus(DatabaseStatus(tmp).String())
	return nil
}

type ListDatabaseBackupsRequestOrderBy string

const (
	ListDatabaseBackupsRequestOrderByCreatedAtDesc = ListDatabaseBackupsRequestOrderBy("created_at_desc")
	ListDatabaseBackupsRequestOrderByCreatedAtAsc  = ListDatabaseBackupsRequestOrderBy("created_at_asc")
)

func (enum ListDatabaseBackupsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListDatabaseBackupsRequestOrderByCreatedAtDesc)
	}
	return string(enum)
}

func (enum ListDatabaseBackupsRequestOrderBy) Values() []ListDatabaseBackupsRequestOrderBy {
	return []ListDatabaseBackupsRequestOrderBy{
		"created_at_desc",
		"created_at_asc",
	}
}

func (enum ListDatabaseBackupsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDatabaseBackupsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDatabaseBackupsRequestOrderBy(ListDatabaseBackupsRequestOrderBy(tmp).String())
	return nil
}

type ListDatabasesRequestOrderBy string

const (
	ListDatabasesRequestOrderByCreatedAtAsc  = ListDatabasesRequestOrderBy("created_at_asc")
	ListDatabasesRequestOrderByCreatedAtDesc = ListDatabasesRequestOrderBy("created_at_desc")
	ListDatabasesRequestOrderByNameAsc       = ListDatabasesRequestOrderBy("name_asc")
	ListDatabasesRequestOrderByNameDesc      = ListDatabasesRequestOrderBy("name_desc")
)

func (enum ListDatabasesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListDatabasesRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListDatabasesRequestOrderBy) Values() []ListDatabasesRequestOrderBy {
	return []ListDatabasesRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListDatabasesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListDatabasesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListDatabasesRequestOrderBy(ListDatabasesRequestOrderBy(tmp).String())
	return nil
}

// DatabaseBackup: database backup.
type DatabaseBackup struct {
	// ID: UUID that uniquely identifies a Serverless SQL Database backup.
	ID string `json:"id"`

	// Status: status of the Serverless SQL Database backup. One of `unknown_status` | `error` | `ready` | `locked`.
	// Default value: unknown_status
	Status DatabaseBackupStatus `json:"status"`

	// OrganizationID: the ID of your Scaleway organization.
	OrganizationID string `json:"organization_id"`

	// ProjectID: UUID of the Scaleway project.
	ProjectID string `json:"project_id"`

	// DatabaseID: UUID of the source Serverless SQL Database the backup is created from.
	DatabaseID string `json:"database_id"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// ExpiresAt: expiration date.
	ExpiresAt *time.Time `json:"expires_at"`

	// Size: size (in bytes) of the database backup file.
	Size *scw.Size `json:"size"`

	// DbSize: size (in bytes) of the database when backup has been done.
	DbSize *scw.Size `json:"db_size"`

	// DownloadURL: download URL of the exported database backup.
	DownloadURL *string `json:"download_url"`

	// DownloadURLExpiresAt: expiration date of the download URL.
	DownloadURLExpiresAt *time.Time `json:"download_url_expires_at"`

	// Region: region of the database backup.
	Region scw.Region `json:"region"`
}

// Database: database.
type Database struct {
	// ID: UUID that uniquely identifies your Serverless SQL DB Database.
	ID string `json:"id"`

	// Name: name of the database.
	Name string `json:"name"`

	// Status: status of the Serverless SQL Ddatabase. One of `unknown_status` | `ready` | `creating` | `deleting` | `error` | `restoring` | `locked`.
	// Default value: unknown_status
	Status DatabaseStatus `json:"status"`

	// Endpoint: endpoint of the database.
	Endpoint string `json:"endpoint"`

	// OrganizationID: the ID of your Scaleway organization.
	OrganizationID string `json:"organization_id"`

	// ProjectID: project ID the database belongs to.
	ProjectID string `json:"project_id"`

	// Region: region of the database.
	Region scw.Region `json:"region"`

	// CreatedAt: creation date.
	CreatedAt *time.Time `json:"created_at"`

	// CPUMin: the minimum number of CPU units for your Serverless SQL Database.
	CPUMin uint32 `json:"cpu_min"`

	// CPUMax: the maximum number of CPU units for your Serverless SQL Database.
	CPUMax uint32 `json:"cpu_max"`

	// CPUCurrent: the current number of CPU units allocated to your Serverless SQL Database.
	CPUCurrent uint32 `json:"cpu_current"`

	// Started: whether your Serverless SQL Database is running or not.
	Started bool `json:"started"`

	// EngineMajorVersion: the major version of the underlying database engine.
	EngineMajorVersion uint32 `json:"engine_major_version"`
}

// CreateDatabaseRequest: create database request.
type CreateDatabaseRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// ProjectID: the ID of your Scaleway project.
	ProjectID string `json:"project_id"`

	// Name: the name of the Serverless SQL Database to be created.
	Name string `json:"name"`

	// CPUMin: the minimum number of CPU units for your Serverless SQL Database.
	CPUMin uint32 `json:"cpu_min"`

	// CPUMax: the maximum number of CPU units for your Serverless SQL Database.
	CPUMax uint32 `json:"cpu_max"`

	// FromBackupID: the ID of the backup to create the database from.
	FromBackupID *string `json:"from_backup_id,omitempty"`
}

// DeleteDatabaseRequest: delete database request.
type DeleteDatabaseRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DatabaseID: UUID of the Serverless SQL Database.
	DatabaseID string `json:"-"`
}

// ExportDatabaseBackupRequest: export database backup request.
type ExportDatabaseBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// BackupID: UUID of the Serverless SQL Database backup.
	BackupID string `json:"-"`
}

// GetDatabaseBackupRequest: get database backup request.
type GetDatabaseBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// BackupID: UUID of the Serverless SQL Database backup.
	BackupID string `json:"-"`
}

// GetDatabaseRequest: get database request.
type GetDatabaseRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DatabaseID: UUID of the Serverless SQL DB database.
	DatabaseID string `json:"-"`
}

// ListDatabaseBackupsRequest: list database backups request.
type ListDatabaseBackupsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrganizationID: filter by the UUID of the Scaleway organization.
	OrganizationID *string `json:"-"`

	// ProjectID: filter by the UUID of the Scaleway project.
	ProjectID *string `json:"-"`

	// DatabaseID: filter by the UUID of the Serverless SQL Database.
	DatabaseID string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// OrderBy: sorting criteria. One of `created_at_asc`, `created_at_desc`.
	// Default value: created_at_desc
	OrderBy ListDatabaseBackupsRequestOrderBy `json:"-"`
}

// ListDatabaseBackupsResponse: list database backups response.
type ListDatabaseBackupsResponse struct {
	// Backups: list of the backups.
	Backups []*DatabaseBackup `json:"backups"`

	// TotalCount: length of the backups list.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatabaseBackupsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatabaseBackupsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListDatabaseBackupsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Backups = append(r.Backups, results.Backups...)
	r.TotalCount += uint64(len(results.Backups))
	return uint64(len(results.Backups)), nil
}

// ListDatabasesRequest: list databases request.
type ListDatabasesRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrganizationID: filter by the UUID of the Scaleway organization.
	OrganizationID *string `json:"-"`

	// ProjectID: UUID of the Scaleway project.
	ProjectID string `json:"-"`

	// Page: page number.
	Page *int32 `json:"-"`

	// PageSize: page size.
	PageSize *uint32 `json:"-"`

	// Name: filter by the name of the database.
	Name *string `json:"-"`

	// OrderBy: sorting criteria. One of `created_at_asc`, `created_at_desc`, `name_asc`, `name_desc`.
	// Default value: created_at_asc
	OrderBy ListDatabasesRequestOrderBy `json:"-"`
}

// ListDatabasesResponse: list databases response.
type ListDatabasesResponse struct {
	// Databases: list of the databases.
	Databases []*Database `json:"databases"`

	// TotalCount: total count of Serverless SQL Databases.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListDatabasesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListDatabasesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListDatabasesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Databases = append(r.Databases, results.Databases...)
	r.TotalCount += uint64(len(results.Databases))
	return uint64(len(results.Databases)), nil
}

// RestoreDatabaseFromBackupRequest: restore database from backup request.
type RestoreDatabaseFromBackupRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DatabaseID: UUID of the Serverless SQL Database.
	DatabaseID string `json:"-"`

	// BackupID: UUID of the Serverless SQL Database backup to restore.
	BackupID string `json:"backup_id"`
}

// UpdateDatabaseRequest: update database request.
type UpdateDatabaseRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// DatabaseID: UUID of the Serverless SQL Database.
	DatabaseID string `json:"-"`

	// CPUMin: the minimum number of CPU units for your Serverless SQL Database.
	CPUMin *uint32 `json:"cpu_min,omitempty"`

	// CPUMax: the maximum number of CPU units for your Serverless SQL Database.
	CPUMax *uint32 `json:"cpu_max,omitempty"`
}

// This API allows you to manage your Serverless SQL Databases.
type API struct {
	client *scw.Client
}

// NewAPI returns a API object from a Scaleway client.
func NewAPI(client *scw.Client) *API {
	return &API{
		client: client,
	}
}

func (s *API) Regions() []scw.Region {
	return []scw.Region{scw.RegionFrPar}
}

// CreateDatabase: You must provide the following parameters: `organization_id`, `project_id`, `name`, `cpu_min`, `cpu_max`. You can also provide `from_backup_id` to create a database from a backup.
func (s *API) CreateDatabase(req *CreateDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
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
		Path:   "/serverless-sqldb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/databases",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Database

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDatabase: Retrieve information about your Serverless SQL Database. You must provide the `database_id` parameter.
func (s *API) GetDatabase(req *GetDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseID) == "" {
		return nil, errors.New("field DatabaseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-sqldb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/databases/" + fmt.Sprint(req.DatabaseID) + "",
	}

	var resp Database

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// WaitForDatabaseRequest is used by WaitForDatabase method.
type WaitForDatabaseRequest struct {
	GetDatabaseRequest
	Timeout       *time.Duration
	RetryInterval *time.Duration
}

// WaitForDatabase waits for the Database to reach a terminal state.
func (s *API) WaitForDatabase(req *WaitForDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
	timeout := defaultServerlessSqldbTimeout
	if req.Timeout != nil {
		timeout = *req.Timeout
	}

	retryInterval := defaultServerlessSqldbRetryInterval
	if req.RetryInterval != nil {
		retryInterval = *req.RetryInterval
	}
	transientStatuses := map[DatabaseStatus]struct{}{
		DatabaseStatusCreating:  {},
		DatabaseStatusDeleting:  {},
		DatabaseStatusRestoring: {},
	}

	res, err := async.WaitSync(&async.WaitSyncConfig{
		Get: func() (interface{}, bool, error) {
			res, err := s.GetDatabase(&GetDatabaseRequest{
				Region:     req.Region,
				DatabaseID: req.DatabaseID,
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
		return nil, errors.Wrap(err, "waiting for Database failed")
	}

	return res.(*Database), nil
}

// DeleteDatabase: Deletes a database. You must provide the `database_id` parameter. All data stored in the database will be permanently deleted.
func (s *API) DeleteDatabase(req *DeleteDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseID) == "" {
		return nil, errors.New("field DatabaseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/serverless-sqldb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/databases/" + fmt.Sprint(req.DatabaseID) + "",
	}

	var resp Database

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDatabases: List all Serverless SQL Databases for a given Scaleway Organization or Scaleway Project. By default, the databases returned in the list are ordered by creation date in ascending order, though this can be modified via the order_by field. For the `name` parameter, the value you include will be checked against the whole name string to see if it includes the string you put in the parameter.
func (s *API) ListDatabases(req *ListDatabasesRequest, opts ...scw.RequestOption) (*ListDatabasesResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-sqldb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/databases",
		Query:  query,
	}

	var resp ListDatabasesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateDatabase: Update CPU limits of your Serverless SQL Database. You must provide the `database_id` parameter.
func (s *API) UpdateDatabase(req *UpdateDatabaseRequest, opts ...scw.RequestOption) (*Database, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseID) == "" {
		return nil, errors.New("field DatabaseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/serverless-sqldb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/databases/" + fmt.Sprint(req.DatabaseID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Database

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// RestoreDatabaseFromBackup: Restore a database from a backup. You must provide the `backup_id` parameter.
func (s *API) RestoreDatabaseFromBackup(req *RestoreDatabaseFromBackupRequest, opts ...scw.RequestOption) (*Database, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.DatabaseID) == "" {
		return nil, errors.New("field DatabaseID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/serverless-sqldb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/databases/" + fmt.Sprint(req.DatabaseID) + "/restore",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Database

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetDatabaseBackup: Retrieve information about your Serverless SQL Database backup. You must provide the `backup_id` parameter.
func (s *API) GetDatabaseBackup(req *GetDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackupID) == "" {
		return nil, errors.New("field BackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-sqldb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.BackupID) + "",
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListDatabaseBackups: List all Serverless SQL Database backups for a given Scaleway Project or Database. By default, the backups returned in the list are ordered by creation date in descending order, though this can be modified via the order_by field.
func (s *API) ListDatabaseBackups(req *ListDatabaseBackupsRequest, opts ...scw.RequestOption) (*ListDatabaseBackupsResponse, error) {
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
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "database_id", req.DatabaseID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/serverless-sqldb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/backups",
		Query:  query,
	}

	var resp ListDatabaseBackupsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ExportDatabaseBackup: Export a database backup providing a download link once the export process is completed. You must provide the `backup_id` parameter.
func (s *API) ExportDatabaseBackup(req *ExportDatabaseBackupRequest, opts ...scw.RequestOption) (*DatabaseBackup, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.BackupID) == "" {
		return nil, errors.New("field BackupID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/serverless-sqldb/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/backups/" + fmt.Sprint(req.BackupID) + "/export",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp DatabaseBackup

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

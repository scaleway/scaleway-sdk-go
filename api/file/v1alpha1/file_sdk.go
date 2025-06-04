// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package file provides methods and message types of the file v1alpha1 API.
package file

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

type AttachmentResourceType string

const (
	AttachmentResourceTypeUnknownResourceType = AttachmentResourceType("unknown_resource_type")
	AttachmentResourceTypeInstanceServer      = AttachmentResourceType("instance_server")
)

func (enum AttachmentResourceType) String() string {
	if enum == "" {
		// return default value if empty
		return string(AttachmentResourceTypeUnknownResourceType)
	}
	return string(enum)
}

func (enum AttachmentResourceType) Values() []AttachmentResourceType {
	return []AttachmentResourceType{
		"unknown_resource_type",
		"instance_server",
	}
}

func (enum AttachmentResourceType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *AttachmentResourceType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = AttachmentResourceType(AttachmentResourceType(tmp).String())
	return nil
}

// Status of the filesystem.
type FileSystemStatus string

const (
	// If unspecified, the filesystem status is unknown by default.
	FileSystemStatusUnknownStatus = FileSystemStatus("unknown_status")
	// The filesystem is created and available.
	FileSystemStatusAvailable = FileSystemStatus("available")
	// The filesystem is in error state.
	FileSystemStatusError = FileSystemStatus("error")
	// The filesystem is under creation (transient).
	FileSystemStatusCreating = FileSystemStatus("creating")
	// The filesystem is being updated (transient).
	FileSystemStatusUpdating = FileSystemStatus("updating")
)

func (enum FileSystemStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(FileSystemStatusUnknownStatus)
	}
	return string(enum)
}

func (enum FileSystemStatus) Values() []FileSystemStatus {
	return []FileSystemStatus{
		"unknown_status",
		"available",
		"error",
		"creating",
		"updating",
	}
}

func (enum FileSystemStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *FileSystemStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = FileSystemStatus(FileSystemStatus(tmp).String())
	return nil
}

// Sorting options for filesystem listing.
type ListFileSystemsRequestOrderBy string

const (
	// Order by creation date (ascending).
	ListFileSystemsRequestOrderByCreatedAtAsc = ListFileSystemsRequestOrderBy("created_at_asc")
	// Order by creation date (descending).
	ListFileSystemsRequestOrderByCreatedAtDesc = ListFileSystemsRequestOrderBy("created_at_desc")
	// Order by name (ascending).
	ListFileSystemsRequestOrderByNameAsc = ListFileSystemsRequestOrderBy("name_asc")
	// Order by name (descending).
	ListFileSystemsRequestOrderByNameDesc = ListFileSystemsRequestOrderBy("name_desc")
)

func (enum ListFileSystemsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListFileSystemsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListFileSystemsRequestOrderBy) Values() []ListFileSystemsRequestOrderBy {
	return []ListFileSystemsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListFileSystemsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListFileSystemsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListFileSystemsRequestOrderBy(ListFileSystemsRequestOrderBy(tmp).String())
	return nil
}

// Attachment: Represents an attachment between a filesystem and a resource.
type Attachment struct {
	// ID: UUID of the attachment.
	ID string `json:"id"`

	// FilesystemID: UUID of the filesystem.
	FilesystemID string `json:"filesystem_id"`

	// ResourceID: UUID of the attached resource.
	ResourceID string `json:"resource_id"`

	// ResourceType: the type of the attached resource.
	// Default value: unknown_resource_type
	ResourceType AttachmentResourceType `json:"resource_type"`

	// Zone: the zone where the resource is located.
	Zone *scw.Zone `json:"zone"`
}

// FileSystem: Represents a filesystem resource and its properties.
type FileSystem struct {
	// ID: UUID of the filesystem.
	ID string `json:"id"`

	// Name: name of the filesystem.
	Name string `json:"name"`

	// Size: filesystem size in bytes.
	Size scw.Size `json:"size"`

	// Status: current status of the filesystem (e.g. creating, available, ...).
	// Default value: unknown_status
	Status FileSystemStatus `json:"status"`

	// ProjectID: UUID of the project to which the filesystem belongs.
	ProjectID string `json:"project_id"`

	// OrganizationID: UUID of the organization to which the filesystem belongs.
	OrganizationID string `json:"organization_id"`

	// Tags: list of tags assigned to the filesystem.
	Tags []string `json:"tags"`

	// NumberOfAttachments: the current number of attachments (mounts) that the filesystem has.
	NumberOfAttachments uint32 `json:"number_of_attachments"`

	// Region: region where the filesystem is located.
	Region scw.Region `json:"region"`

	// CreatedAt: creation date of the filesystem.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: last update date of the properties of the filesystem.
	UpdatedAt *time.Time `json:"updated_at"`
}

// CreateFileSystemRequest: Request to create a new filesystem.
type CreateFileSystemRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// Name: name of the filesystem.
	Name string `json:"name"`

	// ProjectID: UUID of the project the filesystem belongs to.
	ProjectID string `json:"project_id"`

	// Size: must be compliant with the minimum (100 GB) and maximum (10 TB) allowed size.
	Size uint64 `json:"size"`

	// Tags: list of tags assigned to the filesystem.
	Tags []string `json:"tags"`
}

// DeleteFileSystemRequest: Request to delete a specific filesystem.
type DeleteFileSystemRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FilesystemID: UUID of the filesystem.
	FilesystemID string `json:"-"`
}

// GetFileSystemRequest: Request to retrieve a specific filesystem.
type GetFileSystemRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FilesystemID: UUID of the filesystem.
	FilesystemID string `json:"-"`
}

// ListAttachmentsRequest: Request to list filesystem attachments with filtering and pagination options.
type ListAttachmentsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FilesystemID: UUID of the File Storage volume.
	FilesystemID *string `json:"-"`

	// ResourceID: filter by resource ID.
	ResourceID *string `json:"-"`

	// ResourceType: filter by resource type.
	// Default value: unknown_resource_type
	ResourceType AttachmentResourceType `json:"-"`

	// Zone: filter by resource zone.
	Zone *scw.Zone `json:"-"`

	// Page: page number (starting at 1).
	Page *int32 `json:"-"`

	// PageSize: number of entries per page (default: 20, max: 100).
	PageSize *uint32 `json:"-"`
}

// ListAttachmentsResponse: Response containing a list of filesystem attachments and total count.
type ListAttachmentsResponse struct {
	// Attachments: list of filesystem attachments matching the request criteria.
	Attachments []*Attachment `json:"attachments"`

	// TotalCount: total number of filesystem attachments matching the criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListAttachmentsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListAttachmentsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListAttachmentsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Attachments = append(r.Attachments, results.Attachments...)
	r.TotalCount += uint64(len(results.Attachments))
	return uint64(len(results.Attachments)), nil
}

// ListFileSystemsRequest: Request to list filesystems with filtering and pagination options.
type ListFileSystemsRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// OrderBy: criteria to use when ordering the list.
	// Default value: created_at_asc
	OrderBy ListFileSystemsRequestOrderBy `json:"-"`

	// ProjectID: filter by project ID.
	ProjectID *string `json:"-"`

	// OrganizationID: filter by organization ID.
	OrganizationID *string `json:"-"`

	// Page: page number (starting at 1).
	Page *int32 `json:"-"`

	// PageSize: number of entries per page (default: 20, max: 100).
	PageSize *uint32 `json:"-"`

	// Name: filter the returned filesystems by their names.
	Name *string `json:"-"`

	// Tags: filter by tags. Only filesystems with one or more matching tags will be returned.
	Tags []string `json:"-"`
}

// ListFileSystemsResponse: Response containing a list of filesystems and total count.
type ListFileSystemsResponse struct {
	// Filesystems: list of filesystems matching the request criteria.
	Filesystems []*FileSystem `json:"filesystems"`

	// TotalCount: total number of filesystems matching the criteria.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListFileSystemsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListFileSystemsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListFileSystemsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Filesystems = append(r.Filesystems, results.Filesystems...)
	r.TotalCount += uint64(len(results.Filesystems))
	return uint64(len(results.Filesystems)), nil
}

// UpdateFileSystemRequest: Request to update a specific filesystem.
type UpdateFileSystemRequest struct {
	// Region: region to target. If none is passed will use default region from the config.
	Region scw.Region `json:"-"`

	// FilesystemID: UUID of the filesystem.
	FilesystemID string `json:"-"`

	// Name: when defined, is the new name of the filesystem.
	Name *string `json:"name,omitempty"`

	// Size: size in bytes, with a granularity of 100 GB (10^11 bytes).
	// Must be compliant with the minimum (100 GB) and maximum (10 TB) allowed size.
	Size *uint64 `json:"size,omitempty"`

	// Tags: list of tags assigned to the filesystem.
	Tags *[]string `json:"tags,omitempty"`
}

// This API allows you to manage your File Storage resources.
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

// GetFileSystem: Retrieve all properties and current status of a specific filesystem identified by its ID.
func (s *API) GetFileSystem(req *GetFileSystemRequest, opts ...scw.RequestOption) (*FileSystem, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FilesystemID) == "" {
		return nil, errors.New("field FilesystemID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/file/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/filesystems/" + fmt.Sprint(req.FilesystemID) + "",
	}

	var resp FileSystem

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListFileSystems: Retrieve all filesystems in the specified region. By default, the filesystems listed are ordered by creation date in ascending order. This can be modified using the `order_by` field.
func (s *API) ListFileSystems(req *ListFileSystemsRequest, opts ...scw.RequestOption) (*ListFileSystemsResponse, error) {
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
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_id", req.ProjectID)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "tags", req.Tags)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/file/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/filesystems",
		Query:  query,
	}

	var resp ListFileSystemsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListAttachments: List all existing attachments in a specified region.
// By default, the attachments listed are ordered by creation date in ascending order. This can be modified using the `order_by` field.
func (s *API) ListAttachments(req *ListAttachmentsRequest, opts ...scw.RequestOption) (*ListAttachmentsResponse, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	defaultZone, exist := s.client.GetDefaultZone()
	if (req.Zone == nil || *req.Zone == "") && exist {
		req.Zone = &defaultZone
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "filesystem_id", req.FilesystemID)
	parameter.AddToQuery(query, "resource_id", req.ResourceID)
	parameter.AddToQuery(query, "resource_type", req.ResourceType)
	parameter.AddToQuery(query, "zone", req.Zone)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/file/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/attachments",
		Query:  query,
	}

	var resp ListAttachmentsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateFileSystem: To create a new filesystem, you must specify a name, a size, and a project ID.
func (s *API) CreateFileSystem(req *CreateFileSystemRequest, opts ...scw.RequestOption) (*FileSystem, error) {
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
		Path:   "/file/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/filesystems",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FileSystem

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteFileSystem: You must specify the `filesystem_id` of the filesystem you want to delete.
func (s *API) DeleteFileSystem(req *DeleteFileSystemRequest, opts ...scw.RequestOption) error {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FilesystemID) == "" {
		return errors.New("field FilesystemID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/file/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/filesystems/" + fmt.Sprint(req.FilesystemID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateFileSystem: Update the technical details of a filesystem, such as its name, tags or its new size.
func (s *API) UpdateFileSystem(req *UpdateFileSystemRequest, opts ...scw.RequestOption) (*FileSystem, error) {
	var err error

	if req.Region == "" {
		defaultRegion, _ := s.client.GetDefaultRegion()
		req.Region = defaultRegion
	}

	if fmt.Sprint(req.Region) == "" {
		return nil, errors.New("field Region cannot be empty in request")
	}

	if fmt.Sprint(req.FilesystemID) == "" {
		return nil, errors.New("field FilesystemID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/file/v1alpha1/regions/" + fmt.Sprint(req.Region) + "/filesystems/" + fmt.Sprint(req.FilesystemID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp FileSystem

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

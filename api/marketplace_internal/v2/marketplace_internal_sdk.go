// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package marketplace_internal provides methods and message types of the marketplace_internal v2 API.
package marketplace_internal

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/marketplace_internal/v2/internal/errors"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/marketplace_internal/v2/internal/marshaler"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/marketplace_internal/v2/internal/parameter"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/marketplace_internal/v2/namegenerator"
	"gitlab.infra.online.net/devtools/scaleway-sdk-go-internal/marketplace_internal/v2/scw"
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

type ListImagesRequestOrderBy string

const (
	ListImagesRequestOrderByNameAsc       = ListImagesRequestOrderBy("name_asc")
	ListImagesRequestOrderByNameDesc      = ListImagesRequestOrderBy("name_desc")
	ListImagesRequestOrderByCreatedAtAsc  = ListImagesRequestOrderBy("created_at_asc")
	ListImagesRequestOrderByCreatedAtDesc = ListImagesRequestOrderBy("created_at_desc")
	ListImagesRequestOrderByUpdatedAtAsc  = ListImagesRequestOrderBy("updated_at_asc")
	ListImagesRequestOrderByUpdatedAtDesc = ListImagesRequestOrderBy("updated_at_desc")
)

func (enum ListImagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "name_asc"
	}
	return string(enum)
}

func (enum ListImagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListImagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListImagesRequestOrderBy(ListImagesRequestOrderBy(tmp).String())
	return nil
}

type ListLocalImagesRequestOrderBy string

const (
	ListLocalImagesRequestOrderByCreatedAtAsc  = ListLocalImagesRequestOrderBy("created_at_asc")
	ListLocalImagesRequestOrderByCreatedAtDesc = ListLocalImagesRequestOrderBy("created_at_desc")
)

func (enum ListLocalImagesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListLocalImagesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListLocalImagesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListLocalImagesRequestOrderBy(ListLocalImagesRequestOrderBy(tmp).String())
	return nil
}

type ListVersionsRequestOrderBy string

const (
	ListVersionsRequestOrderByCreatedAtAsc  = ListVersionsRequestOrderBy("created_at_asc")
	ListVersionsRequestOrderByCreatedAtDesc = ListVersionsRequestOrderBy("created_at_desc")
)

func (enum ListVersionsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return "created_at_asc"
	}
	return string(enum)
}

func (enum ListVersionsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListVersionsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListVersionsRequestOrderBy(ListVersionsRequestOrderBy(tmp).String())
	return nil
}

// Category: category.
type Category struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Description string `json:"description"`
}

// Image: image.
type Image struct {
	ID string `json:"id"`

	Name string `json:"name"`

	Label string `json:"label"`

	Description string `json:"description"`

	Logo string `json:"logo"`

	Categories []string `json:"categories"`

	ValidUntil *time.Time `json:"valid_until"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	CurrentPublicVersion string `json:"current_public_version"`
}

// LocalImage: local image.
type LocalImage struct {
	ID string `json:"id"`

	Arch string `json:"arch"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`

	CompatibleCommercialTypes []string `json:"compatible_commercial_types"`

	Label string `json:"label"`
}

// Version: version.
type Version struct {
	ID string `json:"id"`

	Name string `json:"name"`

	CreatedAt *time.Time `json:"created_at"`

	UpdatedAt *time.Time `json:"updated_at"`

	PublishedAt *time.Time `json:"published_at"`
}

// CreateCategoryRequest: create category request.
type CreateCategoryRequest struct {
	Name string `json:"name"`

	Description string `json:"description"`
}

// CreateImageRequest: create image request.
type CreateImageRequest struct {
	Name string `json:"name"`

	Label string `json:"label"`

	Logo string `json:"logo"`

	Description string `json:"description"`

	ValidUntil *time.Time `json:"valid_until,omitempty"`

	Categories []string `json:"categories"`
}

// CreateLocalImageRequest: create local image request.
type CreateLocalImageRequest struct {
	VersionID string `json:"version_id"`

	LocalImageID string `json:"local_image_id"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone scw.Zone `json:"zone"`

	Arch string `json:"arch"`

	CompatibleCommercialTypes []string `json:"compatible_commercial_types"`
}

// CreateVersionRequest: create version request.
type CreateVersionRequest struct {
	ImageID string `json:"image_id"`

	Name string `json:"name"`
}

// DeleteCategoryRequest: delete category request.
type DeleteCategoryRequest struct {
	CategoryID string `json:"-"`
}

// DeleteImageRequest: delete image request.
type DeleteImageRequest struct {
	ImageID string `json:"-"`
}

// DeleteLocalImageRequest: delete local image request.
type DeleteLocalImageRequest struct {
	LocalImageID string `json:"-"`
}

// DeleteVersionRequest: delete version request.
type DeleteVersionRequest struct {
	VersionID string `json:"-"`
}

// GetCategoryRequest: get category request.
type GetCategoryRequest struct {
	CategoryID string `json:"-"`
}

// GetImageRequest: get image request.
type GetImageRequest struct {
	ImageID string `json:"-"`
}

// GetLocalImageRequest: get local image request.
type GetLocalImageRequest struct {
	LocalImageID string `json:"-"`
}

// GetVersionRequest: get version request.
type GetVersionRequest struct {
	VersionID string `json:"-"`
}

// ListCategoriesRequest: list categories request.
type ListCategoriesRequest struct {
	PageSize *uint32 `json:"-"`

	Page *int32 `json:"-"`
}

// ListCategoriesResponse: list categories response.
type ListCategoriesResponse struct {
	Categories []*Category `json:"categories"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListCategoriesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListCategoriesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListCategoriesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Categories = append(r.Categories, results.Categories...)
	r.TotalCount += uint32(len(results.Categories))
	return uint32(len(results.Categories)), nil
}

// ListImagesRequest: list images request.
type ListImagesRequest struct {
	PageSize *uint32 `json:"-"`

	Page *int32 `json:"-"`

	// OrderBy: default value: name_asc
	OrderBy ListImagesRequestOrderBy `json:"-"`

	Arch *string `json:"-"`

	Category *string `json:"-"`

	LocalImage *string `json:"-"`

	IncludeEol bool `json:"-"`
}

// ListImagesResponse: list images response.
type ListImagesResponse struct {
	Images []*Image `json:"images"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListImagesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListImagesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListImagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Images = append(r.Images, results.Images...)
	r.TotalCount += uint32(len(results.Images))
	return uint32(len(results.Images)), nil
}

// ListLocalImagesRequest: list local images request.
type ListLocalImagesRequest struct {
	// Precisely one of ImageID, VersionID, ImageLabel must be set.
	ImageID *string `json:"image_id,omitempty"`

	// Precisely one of ImageID, VersionID, ImageLabel must be set.
	VersionID *string `json:"version_id,omitempty"`

	PageSize *uint32 `json:"-"`

	Page *int32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListLocalImagesRequestOrderBy `json:"-"`

	// Precisely one of ImageID, VersionID, ImageLabel must be set.
	ImageLabel *string `json:"image_label,omitempty"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone *scw.Zone `json:"-"`
}

// ListLocalImagesResponse: list local images response.
type ListLocalImagesResponse struct {
	LocalImages []*LocalImage `json:"local_images"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListLocalImagesResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListLocalImagesResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListLocalImagesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.LocalImages = append(r.LocalImages, results.LocalImages...)
	r.TotalCount += uint32(len(results.LocalImages))
	return uint32(len(results.LocalImages)), nil
}

// ListVersionsRequest: list versions request.
type ListVersionsRequest struct {
	ImageID string `json:"-"`

	PageSize *uint32 `json:"-"`

	Page *int32 `json:"-"`

	// OrderBy: default value: created_at_asc
	OrderBy ListVersionsRequestOrderBy `json:"-"`
}

// ListVersionsResponse: list versions response.
type ListVersionsResponse struct {
	Versions []*Version `json:"versions"`

	TotalCount uint32 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListVersionsResponse) UnsafeGetTotalCount() uint32 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListVersionsResponse) UnsafeAppend(res interface{}) (uint32, error) {
	results, ok := res.(*ListVersionsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Versions = append(r.Versions, results.Versions...)
	r.TotalCount += uint32(len(results.Versions))
	return uint32(len(results.Versions)), nil
}

// UpdateCategoryRequest: update category request.
type UpdateCategoryRequest struct {
	CategoryID string `json:"-"`

	Name *string `json:"name,omitempty"`

	Description *string `json:"description,omitempty"`
}

// UpdateImageRequest: update image request.
type UpdateImageRequest struct {
	ImageID string `json:"-"`

	Name *string `json:"name,omitempty"`

	Label *string `json:"label,omitempty"`

	Logo *string `json:"logo,omitempty"`

	Description *string `json:"description,omitempty"`

	ValidUntil *time.Time `json:"valid_until,omitempty"`

	Categories *[]string `json:"categories,omitempty"`

	CurrentPublicVersion *string `json:"current_public_version,omitempty"`
}

// UpdateLocalImageRequest: update local image request.
type UpdateLocalImageRequest struct {
	LocalImageID string `json:"-"`

	// Zone: zone to target. If none is passed will use default zone from the config.
	Zone *scw.Zone `json:"zone,omitempty"`

	Arch *string `json:"arch,omitempty"`

	CompatibleCommercialTypes *[]string `json:"compatible_commercial_types,omitempty"`
}

// UpdateVersionRequest: update version request.
type UpdateVersionRequest struct {
	VersionID string `json:"-"`

	Name *string `json:"name,omitempty"`
}

// Marketplace internal API.
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
		Path:   "/marketplace-internal/v2",
	}

	var resp scw.ServiceInfo

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListImages:
func (s *API) ListImages(req *ListImagesRequest, opts ...scw.RequestOption) (*ListImagesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "arch", req.Arch)
	parameter.AddToQuery(query, "category", req.Category)
	parameter.AddToQuery(query, "local_image", req.LocalImage)
	parameter.AddToQuery(query, "include_eol", req.IncludeEol)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace-internal/v2/images",
		Query:  query,
	}

	var resp ListImagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImage:
func (s *API) GetImage(req *GetImageRequest, opts ...scw.RequestOption) (*Image, error) {
	var err error

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace-internal/v2/images/" + fmt.Sprint(req.ImageID) + "",
	}

	var resp Image

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateImage:
func (s *API) CreateImage(req *CreateImageRequest, opts ...scw.RequestOption) (*Image, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/marketplace-internal/v2/images",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Image

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateImage:
func (s *API) UpdateImage(req *UpdateImageRequest, opts ...scw.RequestOption) (*Image, error) {
	var err error

	if fmt.Sprint(req.ImageID) == "" {
		return nil, errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/marketplace-internal/v2/images/" + fmt.Sprint(req.ImageID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Image

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteImage:
func (s *API) DeleteImage(req *DeleteImageRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.ImageID) == "" {
		return errors.New("field ImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/marketplace-internal/v2/images/" + fmt.Sprint(req.ImageID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListVersions:
func (s *API) ListVersions(req *ListVersionsRequest, opts ...scw.RequestOption) (*ListVersionsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "image_id", req.ImageID)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "order_by", req.OrderBy)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace-internal/v2/versions",
		Query:  query,
	}

	var resp ListVersionsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetVersion:
func (s *API) GetVersion(req *GetVersionRequest, opts ...scw.RequestOption) (*Version, error) {
	var err error

	if fmt.Sprint(req.VersionID) == "" {
		return nil, errors.New("field VersionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace-internal/v2/versions/" + fmt.Sprint(req.VersionID) + "",
	}

	var resp Version

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateVersion:
func (s *API) CreateVersion(req *CreateVersionRequest, opts ...scw.RequestOption) (*Version, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/marketplace-internal/v2/versions",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Version

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateVersion:
func (s *API) UpdateVersion(req *UpdateVersionRequest, opts ...scw.RequestOption) (*Version, error) {
	var err error

	if fmt.Sprint(req.VersionID) == "" {
		return nil, errors.New("field VersionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/marketplace-internal/v2/versions/" + fmt.Sprint(req.VersionID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Version

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteVersion:
func (s *API) DeleteVersion(req *DeleteVersionRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.VersionID) == "" {
		return errors.New("field VersionID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/marketplace-internal/v2/versions/" + fmt.Sprint(req.VersionID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListLocalImages:
func (s *API) ListLocalImages(req *ListLocalImagesRequest, opts ...scw.RequestOption) (*ListLocalImagesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	defaultZone, exist := s.client.GetDefaultZone()
	if (req.Zone == nil || *req.Zone == "") && exist {
		req.Zone = &defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "zone", req.Zone)
	parameter.AddToQuery(query, "image_id", req.ImageID)
	parameter.AddToQuery(query, "version_id", req.VersionID)
	parameter.AddToQuery(query, "image_label", req.ImageLabel)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace-internal/v2/local-images",
		Query:  query,
	}

	var resp ListLocalImagesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetLocalImage:
func (s *API) GetLocalImage(req *GetLocalImageRequest, opts ...scw.RequestOption) (*LocalImage, error) {
	var err error

	if fmt.Sprint(req.LocalImageID) == "" {
		return nil, errors.New("field LocalImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace-internal/v2/local-images/" + fmt.Sprint(req.LocalImageID) + "",
	}

	var resp LocalImage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateLocalImage:
func (s *API) CreateLocalImage(req *CreateLocalImageRequest, opts ...scw.RequestOption) (*LocalImage, error) {
	var err error

	if req.Zone == "" {
		defaultZone, _ := s.client.GetDefaultZone()
		req.Zone = defaultZone
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/marketplace-internal/v2/local-images",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LocalImage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateLocalImage:
func (s *API) UpdateLocalImage(req *UpdateLocalImageRequest, opts ...scw.RequestOption) (*LocalImage, error) {
	var err error

	defaultZone, exist := s.client.GetDefaultZone()
	if (req.Zone == nil || *req.Zone == "") && exist {
		req.Zone = &defaultZone
	}

	if fmt.Sprint(req.LocalImageID) == "" {
		return nil, errors.New("field LocalImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/marketplace-internal/v2/local-images/" + fmt.Sprint(req.LocalImageID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp LocalImage

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteLocalImage:
func (s *API) DeleteLocalImage(req *DeleteLocalImageRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.LocalImageID) == "" {
		return errors.New("field LocalImageID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/marketplace-internal/v2/local-images/" + fmt.Sprint(req.LocalImageID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// ListCategories:
func (s *API) ListCategories(req *ListCategoriesRequest, opts ...scw.RequestOption) (*ListCategoriesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "page", req.Page)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace-internal/v2/categories",
		Query:  query,
	}

	var resp ListCategoriesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetCategory:
func (s *API) GetCategory(req *GetCategoryRequest, opts ...scw.RequestOption) (*Category, error) {
	var err error

	if fmt.Sprint(req.CategoryID) == "" {
		return nil, errors.New("field CategoryID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/marketplace-internal/v2/categories/" + fmt.Sprint(req.CategoryID) + "",
	}

	var resp Category

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateCategory:
func (s *API) CreateCategory(req *CreateCategoryRequest, opts ...scw.RequestOption) (*Category, error) {
	var err error

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/marketplace-internal/v2/categories",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Category

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// UpdateCategory:
func (s *API) UpdateCategory(req *UpdateCategoryRequest, opts ...scw.RequestOption) (*Category, error) {
	var err error

	if fmt.Sprint(req.CategoryID) == "" {
		return nil, errors.New("field CategoryID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/marketplace-internal/v2/categories/" + fmt.Sprint(req.CategoryID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Category

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteCategory:
func (s *API) DeleteCategory(req *DeleteCategoryRequest, opts ...scw.RequestOption) error {
	var err error

	if fmt.Sprint(req.CategoryID) == "" {
		return errors.New("field CategoryID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/marketplace-internal/v2/categories/" + fmt.Sprint(req.CategoryID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

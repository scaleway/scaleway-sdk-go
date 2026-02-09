// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package environmental_footprint provides methods and message types of the environmental_footprint v1alpha1 API.
package environmental_footprint

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

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

type ProductCategory string

const (
	// Unknown product category.
	ProductCategoryUnknownProductCategory = ProductCategory("unknown_product_category")
	// The Apple Silicon product category.
	ProductCategoryAppleSilicon = ProductCategory("apple_silicon")
	// The Block Storage product category.
	ProductCategoryBlockStorage = ProductCategory("block_storage")
	// The Dedibox product category.
	ProductCategoryDedibox = ProductCategory("dedibox")
	// The Elastic Metal product category.
	ProductCategoryElasticMetal = ProductCategory("elastic_metal")
	// The Instances product category.
	ProductCategoryInstances = ProductCategory("instances")
	// The Object Storage product category.
	ProductCategoryObjectStorage = ProductCategory("object_storage")
	// The Load Balancer product category.
	ProductCategoryLoadBalancer = ProductCategory("load_balancer")
	// The Kubernetes product category.
	ProductCategoryKubernetes = ProductCategory("kubernetes")
)

func (enum ProductCategory) String() string {
	if enum == "" {
		// return default value if empty
		return string(ProductCategoryUnknownProductCategory)
	}
	return string(enum)
}

func (enum ProductCategory) Values() []ProductCategory {
	return []ProductCategory{
		"unknown_product_category",
		"apple_silicon",
		"block_storage",
		"dedibox",
		"elastic_metal",
		"instances",
		"object_storage",
		"load_balancer",
		"kubernetes",
	}
}

func (enum ProductCategory) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ProductCategory) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ProductCategory(ProductCategory(tmp).String())
	return nil
}

type ReportType string

const (
	// Unknown report type.
	ReportTypeUnknownReportType = ReportType("unknown_report_type")
	// Monthly report.
	ReportTypeMonthly = ReportType("monthly")
	// Yearly report.
	ReportTypeYearly = ReportType("yearly")
)

func (enum ReportType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ReportTypeUnknownReportType)
	}
	return string(enum)
}

func (enum ReportType) Values() []ReportType {
	return []ReportType{
		"unknown_report_type",
		"monthly",
		"yearly",
	}
}

func (enum ReportType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ReportType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ReportType(ReportType(tmp).String())
	return nil
}

type ServiceCategory string

const (
	// Unknown service category.
	ServiceCategoryUnknownServiceCategory = ServiceCategory("unknown_service_category")
	// The Bare Metal service category.
	ServiceCategoryBaremetal = ServiceCategory("baremetal")
	// The Compute service category.
	ServiceCategoryCompute = ServiceCategory("compute")
	// The Storage service category.
	ServiceCategoryStorage = ServiceCategory("storage")
	// The Network service category.
	ServiceCategoryNetwork = ServiceCategory("network")
	// The Containers service category.
	ServiceCategoryContainers = ServiceCategory("containers")
)

func (enum ServiceCategory) String() string {
	if enum == "" {
		// return default value if empty
		return string(ServiceCategoryUnknownServiceCategory)
	}
	return string(enum)
}

func (enum ServiceCategory) Values() []ServiceCategory {
	return []ServiceCategory{
		"unknown_service_category",
		"baremetal",
		"compute",
		"storage",
		"network",
		"containers",
	}
}

func (enum ServiceCategory) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ServiceCategory) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ServiceCategory(ServiceCategory(tmp).String())
	return nil
}

// Impact: impact.
type Impact struct {
	// KgCo2Equivalent: the estimated carbon emissions in kilograms of CO₂ equivalent (kgCO₂e).
	KgCo2Equivalent float32 `json:"kg_co2_equivalent"`

	// M3WaterUsage: the estimated water consumption in cubic meters (m³).
	M3WaterUsage float32 `json:"m3_water_usage"`
}

// SkuImpact: sku impact.
type SkuImpact struct {
	// Sku: unique ID of the combination of product, region and zone.
	Sku string `json:"sku"`

	// TotalSkuImpact: the total estimated impact for this SKU during the given period.
	TotalSkuImpact *Impact `json:"total_sku_impact"`

	// ServiceCategory: the service category associated with this SKU.
	// Default value: unknown_service_category
	ServiceCategory ServiceCategory `json:"service_category"`

	// ProductCategory: the product category associated with this SKU.
	// Default value: unknown_product_category
	ProductCategory ProductCategory `json:"product_category"`
}

// ZoneImpact: zone impact.
type ZoneImpact struct {
	// Zone: ID of the zone.
	Zone scw.Zone `json:"zone"`

	// TotalZoneImpact: the total estimated impact for this zone across all given service categories, and product categories during the given period.
	TotalZoneImpact *Impact `json:"total_zone_impact"`

	// Skus: list of estimated impact values per SKU.
	Skus []*SkuImpact `json:"skus"`
}

// RegionImpact: region impact.
type RegionImpact struct {
	// Region: ID of the region.
	Region scw.Region `json:"region"`

	// TotalRegionImpact: the total estimated impact for this region across all given zones, service categories, and product categories during the given period.
	TotalRegionImpact *Impact `json:"total_region_impact"`

	// Zones: list of estimated impact values per zone.
	Zones []*ZoneImpact `json:"zones"`
}

// ProjectImpact: project impact.
type ProjectImpact struct {
	// ProjectID: ID of the project.
	ProjectID string `json:"project_id"`

	// TotalProjectImpact: the total estimated impact for this Project across all given regions, zones, service categories, and product categories during the given period.
	TotalProjectImpact *Impact `json:"total_project_impact"`

	// Regions: list of estimated impact values per region.
	Regions []*RegionImpact `json:"regions"`
}

// ImpactDataResponse: impact data response.
type ImpactDataResponse struct {
	// StartDate: start date of the impact data period (inclusive).
	StartDate *time.Time `json:"start_date"`

	// EndDate: end date of the impact data period (exclusive).
	EndDate *time.Time `json:"end_date"`

	// TotalImpact: the total estimated impact across all given Projects, regions, zones, service categories, and product categories during the given period.
	TotalImpact *Impact `json:"total_impact"`

	// Projects: list of estimated impact values per Project.
	Projects []*ProjectImpact `json:"projects"`
}

// ImpactReportAvailability: impact report availability.
type ImpactReportAvailability struct {
	// MonthSummaryReports: the list of calendar months for which impact reports are available.
	MonthSummaryReports []*time.Time `json:"month_summary_reports"`

	// YearlySummaryReports: the list of calendar years for which impact reports are available.
	YearlySummaryReports []*time.Time `json:"yearly_summary_reports"`
}

// UserAPIDownloadImpactReportRequest: user api download impact report request.
type UserAPIDownloadImpactReportRequest struct {
	// OrganizationID: the UUID of the Organization for which you want to download a report.
	OrganizationID string `json:"organization_id"`

	// Date: the start date of the period for which you want to download a report (ISO 8601 format, e.g. 2025-05-01T00:00:00Z).
	Date *time.Time `json:"date,omitempty"`

	// Type: type of report to download (e.g. `monthly`).
	// Default value: unknown_report_type
	Type ReportType `json:"type"`
}

// UserAPIGetImpactDataRequest: user api get impact data request.
type UserAPIGetImpactDataRequest struct {
	// OrganizationID: the UUID of the Organization for which you want to download a report.
	OrganizationID string `json:"-"`

	// StartDate: start date (inclusive) of the period for which you want to retrieve impact data (ISO 8601 format, e.g. 2025-05-01T00:00:00Z).
	StartDate *time.Time `json:"-"`

	// EndDate: end date (exclusive) of the period for which you want to retrieve impact data (ISO 8601 format, with time in UTC, `YYYY-MM-DDTHH:MM:SSZ`). Defaults to today's date.
	EndDate *time.Time `json:"-"`

	// Regions: list of regions to filter by (e.g. `fr-par`). Defaults to all regions.
	Regions []string `json:"-"`

	// Zones: list of zones to filter by (e.g. `fr-par-1`). Defaults to all zones.
	Zones []string `json:"-"`

	// ProjectIDs: list of Project IDs to filter by. Defaults to all Projects in the Organization.
	ProjectIDs []string `json:"-"`

	// ServiceCategories: list of service categories to filter by. Defaults to all service categories.
	ServiceCategories []ServiceCategory `json:"-"`

	// ProductCategories: list of product categories to filter by. Defaults to all product categories.
	ProductCategories []ProductCategory `json:"-"`
}

// UserAPIGetImpactReportAvailabilityRequest: user api get impact report availability request.
type UserAPIGetImpactReportAvailabilityRequest struct {
	// OrganizationID: the UUID of the Organization for which you want to download a report.
	OrganizationID string `json:"-"`

	// StartDate: start date of the search period (ISO 8601 format, with time in UTC, `YYYY-MM-DDTHH:MM:SSZ`). The date is inclusive.
	StartDate *time.Time `json:"-"`

	// EndDate: end date of the search period (ISO 8601 format, with time in UTC, `YYYY-MM-DDTHH:MM:SSZ`). The date is inclusive. Defaults to today's date.
	EndDate *time.Time `json:"-"`
}

// Access and download impact reports and impact data for your Scaleway projects. Our API provides key metrics such as estimated carbon emissions and water usage to help monitor your environmental footprint.
type UserAPI struct {
	client *scw.Client
}

// NewUserAPI returns a UserAPI object from a Scaleway client.
func NewUserAPI(client *scw.Client) *UserAPI {
	return &UserAPI{
		client: client,
	}
}

// GetImpactReportAvailability: Returns a list of dates of available impact reports.
func (s *UserAPI) GetImpactReportAvailability(req *UserAPIGetImpactReportAvailabilityRequest, opts ...scw.RequestOption) (*ImpactReportAvailability, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/environmental-footprint/v1alpha1/reports/availability",
		Query:  query,
	}

	var resp ImpactReportAvailability

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DownloadImpactReport: Download a Scaleway impact PDF report with detailed impact data for your Scaleway projects.
func (s *UserAPI) DownloadImpactReport(req *UserAPIDownloadImpactReportRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/environmental-footprint/v1alpha1/reports/download",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetImpactData: Retrieve detailed impact data for your Scaleway projects within a specified date range. Filter by project ID, region, zone, service category, and/or product category.
func (s *UserAPI) GetImpactData(req *UserAPIGetImpactDataRequest, opts ...scw.RequestOption) (*ImpactDataResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "start_date", req.StartDate)
	parameter.AddToQuery(query, "end_date", req.EndDate)
	parameter.AddToQuery(query, "regions", req.Regions)
	parameter.AddToQuery(query, "zones", req.Zones)
	parameter.AddToQuery(query, "project_ids", req.ProjectIDs)
	parameter.AddToQuery(query, "service_categories", req.ServiceCategories)
	parameter.AddToQuery(query, "product_categories", req.ProductCategories)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/environmental-footprint/v1alpha1/data/query",
		Query:  query,
	}

	var resp ImpactDataResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

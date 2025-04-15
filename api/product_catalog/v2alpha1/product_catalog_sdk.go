// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package product_catalog provides methods and message types of the product_catalog v2alpha1 API.
package product_catalog

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

type PublicCatalogProductPriceUnitOfMeasureCountableUnit string

const (
	// Unknown countable unit.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitUnknownCountableUnit = PublicCatalogProductPriceUnitOfMeasureCountableUnit("unknown_countable_unit")
	// Chunk.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitChunk = PublicCatalogProductPriceUnitOfMeasureCountableUnit("chunk")
	// Core.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitCore = PublicCatalogProductPriceUnitOfMeasureCountableUnit("core")
	// Currency.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitCurrency = PublicCatalogProductPriceUnitOfMeasureCountableUnit("currency")
	// Device.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitDevice = PublicCatalogProductPriceUnitOfMeasureCountableUnit("device")
	// Domain.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitDomain = PublicCatalogProductPriceUnitOfMeasureCountableUnit("domain")
	// Email.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitEmail = PublicCatalogProductPriceUnitOfMeasureCountableUnit("email")
	// GB/s.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitGbS = PublicCatalogProductPriceUnitOfMeasureCountableUnit("gb_s")
	// Gigabyte.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitGigabyte = PublicCatalogProductPriceUnitOfMeasureCountableUnit("gigabyte")
	// Hour.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitHour = PublicCatalogProductPriceUnitOfMeasureCountableUnit("hour")
	// IOPS gigabyte.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitIopsGigabyte = PublicCatalogProductPriceUnitOfMeasureCountableUnit("iops_gigabyte")
	// IP.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitIP = PublicCatalogProductPriceUnitOfMeasureCountableUnit("ip")
	// Month.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitMonth = PublicCatalogProductPriceUnitOfMeasureCountableUnit("month")
	// Node.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitNode = PublicCatalogProductPriceUnitOfMeasureCountableUnit("node")
	// Plan.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitPlan = PublicCatalogProductPriceUnitOfMeasureCountableUnit("plan")
	// Query.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitQuery = PublicCatalogProductPriceUnitOfMeasureCountableUnit("query")
	// Request.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitRequest = PublicCatalogProductPriceUnitOfMeasureCountableUnit("request")
	// Session.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitSession = PublicCatalogProductPriceUnitOfMeasureCountableUnit("session")
	// VCPU/s.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitVcpuS = PublicCatalogProductPriceUnitOfMeasureCountableUnit("vcpu_s")
	// Version.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitVersion = PublicCatalogProductPriceUnitOfMeasureCountableUnit("version")
	// Year.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitYear = PublicCatalogProductPriceUnitOfMeasureCountableUnit("year")
	// Key.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitKey = PublicCatalogProductPriceUnitOfMeasureCountableUnit("key")
	// Token.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitToken = PublicCatalogProductPriceUnitOfMeasureCountableUnit("token")
	// Minute.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitMinute = PublicCatalogProductPriceUnitOfMeasureCountableUnit("minute")
	// The installation of a resource (software or hardware).
	PublicCatalogProductPriceUnitOfMeasureCountableUnitSetup = PublicCatalogProductPriceUnitOfMeasureCountableUnit("setup")
	// Day.
	PublicCatalogProductPriceUnitOfMeasureCountableUnitDay = PublicCatalogProductPriceUnitOfMeasureCountableUnit("day")
)

func (enum PublicCatalogProductPriceUnitOfMeasureCountableUnit) String() string {
	if enum == "" {
		// return default value if empty
		return string(PublicCatalogProductPriceUnitOfMeasureCountableUnitUnknownCountableUnit)
	}
	return string(enum)
}

func (enum PublicCatalogProductPriceUnitOfMeasureCountableUnit) Values() []PublicCatalogProductPriceUnitOfMeasureCountableUnit {
	return []PublicCatalogProductPriceUnitOfMeasureCountableUnit{
		"unknown_countable_unit",
		"chunk",
		"core",
		"currency",
		"device",
		"domain",
		"email",
		"gb_s",
		"gigabyte",
		"hour",
		"iops_gigabyte",
		"ip",
		"month",
		"node",
		"plan",
		"query",
		"request",
		"session",
		"vcpu_s",
		"version",
		"year",
		"key",
		"token",
		"minute",
		"setup",
		"day",
	}
}

func (enum PublicCatalogProductPriceUnitOfMeasureCountableUnit) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PublicCatalogProductPriceUnitOfMeasureCountableUnit) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PublicCatalogProductPriceUnitOfMeasureCountableUnit(PublicCatalogProductPriceUnitOfMeasureCountableUnit(tmp).String())
	return nil
}

type PublicCatalogProductPropertiesHardwareCPUArch string

const (
	// Unknown architecture.
	PublicCatalogProductPropertiesHardwareCPUArchUnknownArch = PublicCatalogProductPropertiesHardwareCPUArch("unknown_arch")
	// X64.
	PublicCatalogProductPropertiesHardwareCPUArchX64 = PublicCatalogProductPropertiesHardwareCPUArch("x64")
	// ARM64.
	PublicCatalogProductPropertiesHardwareCPUArchArm64 = PublicCatalogProductPropertiesHardwareCPUArch("arm64")
	// RISC-V.
	PublicCatalogProductPropertiesHardwareCPUArchRiscv = PublicCatalogProductPropertiesHardwareCPUArch("riscv")
	// Apple Silicon.
	PublicCatalogProductPropertiesHardwareCPUArchAppleSilicon = PublicCatalogProductPropertiesHardwareCPUArch("apple_silicon")
)

func (enum PublicCatalogProductPropertiesHardwareCPUArch) String() string {
	if enum == "" {
		// return default value if empty
		return string(PublicCatalogProductPropertiesHardwareCPUArchUnknownArch)
	}
	return string(enum)
}

func (enum PublicCatalogProductPropertiesHardwareCPUArch) Values() []PublicCatalogProductPropertiesHardwareCPUArch {
	return []PublicCatalogProductPropertiesHardwareCPUArch{
		"unknown_arch",
		"x64",
		"arm64",
		"riscv",
		"apple_silicon",
	}
}

func (enum PublicCatalogProductPropertiesHardwareCPUArch) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PublicCatalogProductPropertiesHardwareCPUArch) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PublicCatalogProductPropertiesHardwareCPUArch(PublicCatalogProductPropertiesHardwareCPUArch(tmp).String())
	return nil
}

// PublicCatalogProductPropertiesHardwareCPUPhysical: public catalog product properties hardware cpu physical.
type PublicCatalogProductPropertiesHardwareCPUPhysical struct {
	// Sockets: the number of sockets of the CPU.
	Sockets uint32 `json:"sockets"`

	// CoresPerSocket: the number of cores per socket.
	CoresPerSocket uint32 `json:"cores_per_socket"`

	// ThreadsPerCore: the number of threads per core.
	ThreadsPerCore uint32 `json:"threads_per_core"`

	// Frequency: the frequency of the CPU in Hertz.
	Frequency uint64 `json:"frequency"`

	// Benchmark: the benchmark score of the CPU.
	Benchmark uint32 `json:"benchmark"`
}

// PublicCatalogProductPropertiesHardwareCPUVirtual: public catalog product properties hardware cpu virtual.
type PublicCatalogProductPropertiesHardwareCPUVirtual struct {
	// Count: the number of vCPUs.
	Count uint32 `json:"count"`
}

// PublicCatalogProductPropertiesHardwareCPU: public catalog product properties hardware cpu.
type PublicCatalogProductPropertiesHardwareCPU struct {
	// Description: a human readable description of the CPU.
	Description string `json:"description"`

	// Arch: the architecture of the CPU.
	// Default value: unknown_arch
	Arch PublicCatalogProductPropertiesHardwareCPUArch `json:"arch"`

	// Type: the type of the CPU.
	Type string `json:"type"`

	// Virtual: properties if the CPU is virtual.
	// Precisely one of Virtual, Physical must be set.
	Virtual *PublicCatalogProductPropertiesHardwareCPUVirtual `json:"virtual,omitempty"`

	// Physical: properties if the CPU is physical.
	// Precisely one of Virtual, Physical must be set.
	Physical *PublicCatalogProductPropertiesHardwareCPUPhysical `json:"physical,omitempty"`

	// Threads: the total number of threads.
	Threads uint32 `json:"threads"`
}

// PublicCatalogProductPropertiesHardwareGPU: public catalog product properties hardware gpu.
type PublicCatalogProductPropertiesHardwareGPU struct {
	// Description: a human-readable description of the GPU.
	Description string `json:"description"`

	// Count: the number of GPUs.
	Count uint32 `json:"count"`

	// Type: the type of the GPU.
	Type string `json:"type"`
}

// PublicCatalogProductPropertiesHardwareNetwork: public catalog product properties hardware network.
type PublicCatalogProductPropertiesHardwareNetwork struct {
	// Description: a human-readable description of the network.
	Description string `json:"description"`

	// InternalBandwidth: the internal bandwidth in bits per second.
	InternalBandwidth uint64 `json:"internal_bandwidth"`

	// PublicBandwidth: the default public bandwidth in bits per second.
	PublicBandwidth uint64 `json:"public_bandwidth"`

	// MaxPublicBandwidth: the maximum public bandwidth in bits per second (may require subscription to options).
	MaxPublicBandwidth uint64 `json:"max_public_bandwidth"`
}

// PublicCatalogProductPropertiesHardwareRAM: public catalog product properties hardware ram.
type PublicCatalogProductPropertiesHardwareRAM struct {
	// Description: a human-readable description of the RAM.
	Description string `json:"description"`

	// Size: the size of the RAM in bytes.
	Size scw.Size `json:"size"`

	// Type: the type of the RAM.
	Type string `json:"type"`
}

// PublicCatalogProductPropertiesHardwareStorage: public catalog product properties hardware storage.
type PublicCatalogProductPropertiesHardwareStorage struct {
	// Description: a human-readable description of the storage.
	Description string `json:"description"`

	// Total: the total size of the storage in bytes.
	Total scw.Size `json:"total"`
}

// PublicCatalogProductPriceUnitOfMeasure: public catalog product price unit of measure.
type PublicCatalogProductPriceUnitOfMeasure struct {
	// Unit: the unit of measure.
	// Default value: unknown_countable_unit
	Unit PublicCatalogProductPriceUnitOfMeasureCountableUnit `json:"unit"`

	// Size: the size of the unit.
	Size uint64 `json:"size"`
}

// PublicCatalogProductPropertiesAppleSilicon: public catalog product properties apple silicon.
type PublicCatalogProductPropertiesAppleSilicon struct {
	// Range: the range of the Apple Silicon server.
	Range string `json:"range"`
}

// PublicCatalogProductPropertiesDedibox: public catalog product properties dedibox.
type PublicCatalogProductPropertiesDedibox struct {
	// Range: the range of the Dedibox server.
	Range string `json:"range"`
}

// PublicCatalogProductPropertiesElasticMetal: public catalog product properties elastic metal.
type PublicCatalogProductPropertiesElasticMetal struct {
	// Range: the range of the Elastic Metal server.
	Range string `json:"range"`
}

// PublicCatalogProductPropertiesHardware: public catalog product properties hardware.
type PublicCatalogProductPropertiesHardware struct {
	// CPU: the CPU hardware properties.
	CPU *PublicCatalogProductPropertiesHardwareCPU `json:"cpu"`

	// RAM: the RAM hardware properties.
	RAM *PublicCatalogProductPropertiesHardwareRAM `json:"ram"`

	// Storage: the storage hardware properties.
	Storage *PublicCatalogProductPropertiesHardwareStorage `json:"storage"`

	// Network: the network hardware properties.
	Network *PublicCatalogProductPropertiesHardwareNetwork `json:"network"`

	// Gpu: the GPU hardware properties.
	Gpu *PublicCatalogProductPropertiesHardwareGPU `json:"gpu"`
}

// PublicCatalogProductPropertiesInstance: public catalog product properties instance.
type PublicCatalogProductPropertiesInstance struct {
	// Range: the range of the Instance server.
	Range string `json:"range"`
}

// PublicCatalogProductEnvironmentalImpact: public catalog product environmental impact.
type PublicCatalogProductEnvironmentalImpact struct {
	// KgCo2Equivalent: kilograms of CO2 that would need to be released to produce the equivalent warming impact.
	KgCo2Equivalent *float32 `json:"kg_co2_equivalent"`

	// M3WaterUsage: cubic meters of water used.
	M3WaterUsage *float32 `json:"m3_water_usage"`
}

// PublicCatalogProductLocality: public catalog product locality.
type PublicCatalogProductLocality struct {
	// Global: whether or not the product is global.
	// Precisely one of Global, Region, Zone, Datacenter must be set.
	Global *bool `json:"global,omitempty"`

	// Region: the region of the product.
	// Precisely one of Global, Region, Zone, Datacenter must be set.
	Region *scw.Region `json:"region,omitempty"`

	// Zone: the zone of the product.
	// Precisely one of Global, Region, Zone, Datacenter must be set.
	Zone *scw.Zone `json:"zone,omitempty"`

	// Datacenter: the datacenter of the product.
	// Precisely one of Global, Region, Zone, Datacenter must be set.
	Datacenter *string `json:"datacenter,omitempty"`
}

// PublicCatalogProductPrice: public catalog product price.
type PublicCatalogProductPrice struct {
	// RetailPrice: the retail price of the product.
	RetailPrice *scw.Money `json:"retail_price"`

	// UnitOfMeasure: the unit of measure of the price.
	UnitOfMeasure *PublicCatalogProductPriceUnitOfMeasure `json:"unit_of_measure"`
}

// PublicCatalogProductProperties: public catalog product properties.
type PublicCatalogProductProperties struct {
	// Hardware: the hardware properties of the product (if supported).
	Hardware *PublicCatalogProductPropertiesHardware `json:"hardware"`

	// Dedibox: the properties of Dedibox products.
	// Precisely one of Dedibox, ElasticMetal, AppleSilicon, Instance must be set.
	Dedibox *PublicCatalogProductPropertiesDedibox `json:"dedibox,omitempty"`

	// ElasticMetal: the properties of Elastic Metal products.
	// Precisely one of Dedibox, ElasticMetal, AppleSilicon, Instance must be set.
	ElasticMetal *PublicCatalogProductPropertiesElasticMetal `json:"elastic_metal,omitempty"`

	// AppleSilicon: the properties of Apple Silicon products.
	// Precisely one of Dedibox, ElasticMetal, AppleSilicon, Instance must be set.
	AppleSilicon *PublicCatalogProductPropertiesAppleSilicon `json:"apple_silicon,omitempty"`

	// Instance: the properties of Instance products.
	// Precisely one of Dedibox, ElasticMetal, AppleSilicon, Instance must be set.
	Instance *PublicCatalogProductPropertiesInstance `json:"instance,omitempty"`
}

// PublicCatalogProduct: public catalog product.
type PublicCatalogProduct struct {
	// Sku: the unique identifier of the product.
	Sku string `json:"sku"`

	// ServiceCategory: the category of the product.
	ServiceCategory string `json:"service_category"`

	// Product: the product name.
	Product string `json:"product"`

	// Variant: the product variant.
	Variant string `json:"variant"`

	// Description: the product description.
	Description string `json:"description"`

	// Locality: the locality of the product.
	Locality *PublicCatalogProductLocality `json:"locality"`

	// Price: the price of the product.
	Price *PublicCatalogProductPrice `json:"price"`

	// Properties: the properties of the product.
	Properties *PublicCatalogProductProperties `json:"properties"`

	// EnvironmentalImpact: the environmental impact of the product.
	EnvironmentalImpact *PublicCatalogProductEnvironmentalImpact `json:"environmental_impact"`
}

// ListPublicCatalogProductsResponse: list public catalog products response.
type ListPublicCatalogProductsResponse struct {
	// Products: the list of products.
	Products []*PublicCatalogProduct `json:"products"`

	// TotalCount: the total number of products in the catalog.
	TotalCount uint64 `json:"total_count"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListPublicCatalogProductsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListPublicCatalogProductsResponse) UnsafeAppend(res interface{}) (uint64, error) {
	results, ok := res.(*ListPublicCatalogProductsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Products = append(r.Products, results.Products...)
	r.TotalCount += uint64(len(results.Products))
	return uint64(len(results.Products)), nil
}

// PublicCatalogAPIListPublicCatalogProductsRequest: public catalog api list public catalog products request.
type PublicCatalogAPIListPublicCatalogProductsRequest struct {
	// Page: number of the page. Value must be greater or equal to 1.
	Page *int32 `json:"-"`

	// PageSize: the number of products per page. Value must be greater or equal to 1.
	PageSize *uint32 `json:"-"`
}

type PublicCatalogAPI struct {
	client *scw.Client
}

// NewPublicCatalogAPI returns a PublicCatalogAPI object from a Scaleway client.
func NewPublicCatalogAPI(client *scw.Client) *PublicCatalogAPI {
	return &PublicCatalogAPI{
		client: client,
	}
}

// ListPublicCatalogProducts: List all available products in the Scaleway catalog. Returns a complete list of products with their corresponding description, locations, prices and properties. You can define the `page` number and `page_size` for your query in the request.
func (s *PublicCatalogAPI) ListPublicCatalogProducts(req *PublicCatalogAPIListPublicCatalogProductsRequest, opts ...scw.RequestOption) (*ListPublicCatalogProductsResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/product-catalog/v2alpha1/public-catalog/products",
		Query:  query,
	}

	var resp ListPublicCatalogProductsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

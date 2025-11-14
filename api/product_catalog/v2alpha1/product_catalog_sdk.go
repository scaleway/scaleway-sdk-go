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

type ListPublicCatalogProductsRequestProductType string

const (
	// Unknown product type.
	ListPublicCatalogProductsRequestProductTypeUnknownProductType = ListPublicCatalogProductsRequestProductType("unknown_product_type")
	// Include the Instance information in the response.
	ListPublicCatalogProductsRequestProductTypeInstance = ListPublicCatalogProductsRequestProductType("instance")
	// Include the Apple Silicon information in the response.
	ListPublicCatalogProductsRequestProductTypeAppleSilicon = ListPublicCatalogProductsRequestProductType("apple_silicon")
	// Include the Elastic Metal information in the response.
	ListPublicCatalogProductsRequestProductTypeElasticMetal = ListPublicCatalogProductsRequestProductType("elastic_metal")
	// Include the Dedibox information in the response.
	ListPublicCatalogProductsRequestProductTypeDedibox = ListPublicCatalogProductsRequestProductType("dedibox")
	// Include the Block Storage information in the response.
	ListPublicCatalogProductsRequestProductTypeBlockStorage = ListPublicCatalogProductsRequestProductType("block_storage")
	// Include the Object Storage information in the response.
	ListPublicCatalogProductsRequestProductTypeObjectStorage = ListPublicCatalogProductsRequestProductType("object_storage")
	// Include the Managed Inference information in the response.
	ListPublicCatalogProductsRequestProductTypeManagedInference = ListPublicCatalogProductsRequestProductType("managed_inference")
)

func (enum ListPublicCatalogProductsRequestProductType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPublicCatalogProductsRequestProductTypeUnknownProductType)
	}
	return string(enum)
}

func (enum ListPublicCatalogProductsRequestProductType) Values() []ListPublicCatalogProductsRequestProductType {
	return []ListPublicCatalogProductsRequestProductType{
		"unknown_product_type",
		"instance",
		"apple_silicon",
		"elastic_metal",
		"dedibox",
		"block_storage",
		"object_storage",
		"managed_inference",
	}
}

func (enum ListPublicCatalogProductsRequestProductType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPublicCatalogProductsRequestProductType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPublicCatalogProductsRequestProductType(ListPublicCatalogProductsRequestProductType(tmp).String())
	return nil
}

type ListPublicCatalogProductsRequestStatus string

const (
	// Unknown status.
	ListPublicCatalogProductsRequestStatusUnknownStatus = ListPublicCatalogProductsRequestStatus("unknown_status")
	// The product is available in Public Beta.
	ListPublicCatalogProductsRequestStatusPublicBeta = ListPublicCatalogProductsRequestStatus("public_beta")
	// The product is available in Preview mode.
	ListPublicCatalogProductsRequestStatusPreview = ListPublicCatalogProductsRequestStatus("preview")
	// The product is generally available.
	ListPublicCatalogProductsRequestStatusGeneralAvailability = ListPublicCatalogProductsRequestStatus("general_availability")
	// The product must not be used for new deployments.
	ListPublicCatalogProductsRequestStatusEndOfDeployment = ListPublicCatalogProductsRequestStatus("end_of_deployment")
	// There is no longer any commercial support for this product.
	ListPublicCatalogProductsRequestStatusEndOfSupport = ListPublicCatalogProductsRequestStatus("end_of_support")
	// The product is not sold anymore but is still in use.
	ListPublicCatalogProductsRequestStatusEndOfSale = ListPublicCatalogProductsRequestStatus("end_of_sale")
	// The product is no longer supported or maintained.
	ListPublicCatalogProductsRequestStatusEndOfLife = ListPublicCatalogProductsRequestStatus("end_of_life")
	// The product is deprecated and is no longer accessible.
	ListPublicCatalogProductsRequestStatusRetired = ListPublicCatalogProductsRequestStatus("retired")
)

func (enum ListPublicCatalogProductsRequestStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListPublicCatalogProductsRequestStatusUnknownStatus)
	}
	return string(enum)
}

func (enum ListPublicCatalogProductsRequestStatus) Values() []ListPublicCatalogProductsRequestStatus {
	return []ListPublicCatalogProductsRequestStatus{
		"unknown_status",
		"public_beta",
		"preview",
		"general_availability",
		"end_of_deployment",
		"end_of_support",
		"end_of_sale",
		"end_of_life",
		"retired",
	}
}

func (enum ListPublicCatalogProductsRequestStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListPublicCatalogProductsRequestStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListPublicCatalogProductsRequestStatus(ListPublicCatalogProductsRequestStatus(tmp).String())
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

type PublicCatalogProductStatus string

const (
	// Unknown status.
	PublicCatalogProductStatusUnknownStatus = PublicCatalogProductStatus("unknown_status")
	// The product is available in Public Beta.
	PublicCatalogProductStatusPublicBeta = PublicCatalogProductStatus("public_beta")
	// The product is available in Preview mode.
	PublicCatalogProductStatusPreview = PublicCatalogProductStatus("preview")
	// The product is generally available.
	PublicCatalogProductStatusGeneralAvailability = PublicCatalogProductStatus("general_availability")
	// The product must not be used for new deployments.
	PublicCatalogProductStatusEndOfDeployment = PublicCatalogProductStatus("end_of_deployment")
	// There is no longer any commercial support for this product.
	PublicCatalogProductStatusEndOfSupport = PublicCatalogProductStatus("end_of_support")
	// The product is not sold anymore but is still in use.
	PublicCatalogProductStatusEndOfSale = PublicCatalogProductStatus("end_of_sale")
	// The product is at its end of life.
	PublicCatalogProductStatusEndOfLife = PublicCatalogProductStatus("end_of_life")
	// The product is retired.
	PublicCatalogProductStatusRetired = PublicCatalogProductStatus("retired")
)

func (enum PublicCatalogProductStatus) String() string {
	if enum == "" {
		// return default value if empty
		return string(PublicCatalogProductStatusUnknownStatus)
	}
	return string(enum)
}

func (enum PublicCatalogProductStatus) Values() []PublicCatalogProductStatus {
	return []PublicCatalogProductStatus{
		"unknown_status",
		"public_beta",
		"preview",
		"general_availability",
		"end_of_deployment",
		"end_of_support",
		"end_of_sale",
		"end_of_life",
		"retired",
	}
}

func (enum PublicCatalogProductStatus) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PublicCatalogProductStatus) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PublicCatalogProductStatus(PublicCatalogProductStatus(tmp).String())
	return nil
}

type PublicCatalogProductUnitOfMeasureCountableUnit string

const (
	// Unknown countable unit.
	PublicCatalogProductUnitOfMeasureCountableUnitUnknownCountableUnit = PublicCatalogProductUnitOfMeasureCountableUnit("unknown_countable_unit")
	// Chunk.
	PublicCatalogProductUnitOfMeasureCountableUnitChunk = PublicCatalogProductUnitOfMeasureCountableUnit("chunk")
	// Core.
	PublicCatalogProductUnitOfMeasureCountableUnitCore = PublicCatalogProductUnitOfMeasureCountableUnit("core")
	// Currency.
	PublicCatalogProductUnitOfMeasureCountableUnitCurrency = PublicCatalogProductUnitOfMeasureCountableUnit("currency")
	// Device.
	PublicCatalogProductUnitOfMeasureCountableUnitDevice = PublicCatalogProductUnitOfMeasureCountableUnit("device")
	// Domain.
	PublicCatalogProductUnitOfMeasureCountableUnitDomain = PublicCatalogProductUnitOfMeasureCountableUnit("domain")
	// Email.
	PublicCatalogProductUnitOfMeasureCountableUnitEmail = PublicCatalogProductUnitOfMeasureCountableUnit("email")
	// GB/s.
	PublicCatalogProductUnitOfMeasureCountableUnitGbS = PublicCatalogProductUnitOfMeasureCountableUnit("gb_s")
	// Gigabyte.
	PublicCatalogProductUnitOfMeasureCountableUnitGigabyte = PublicCatalogProductUnitOfMeasureCountableUnit("gigabyte")
	// Hour.
	PublicCatalogProductUnitOfMeasureCountableUnitHour = PublicCatalogProductUnitOfMeasureCountableUnit("hour")
	// IOPS gigabyte.
	PublicCatalogProductUnitOfMeasureCountableUnitIopsGigabyte = PublicCatalogProductUnitOfMeasureCountableUnit("iops_gigabyte")
	// IP.
	PublicCatalogProductUnitOfMeasureCountableUnitIP = PublicCatalogProductUnitOfMeasureCountableUnit("ip")
	// Month.
	PublicCatalogProductUnitOfMeasureCountableUnitMonth = PublicCatalogProductUnitOfMeasureCountableUnit("month")
	// Node.
	PublicCatalogProductUnitOfMeasureCountableUnitNode = PublicCatalogProductUnitOfMeasureCountableUnit("node")
	// Plan.
	PublicCatalogProductUnitOfMeasureCountableUnitPlan = PublicCatalogProductUnitOfMeasureCountableUnit("plan")
	// Query.
	PublicCatalogProductUnitOfMeasureCountableUnitQuery = PublicCatalogProductUnitOfMeasureCountableUnit("query")
	// Request.
	PublicCatalogProductUnitOfMeasureCountableUnitRequest = PublicCatalogProductUnitOfMeasureCountableUnit("request")
	// Session.
	PublicCatalogProductUnitOfMeasureCountableUnitSession = PublicCatalogProductUnitOfMeasureCountableUnit("session")
	// VCPU/s.
	PublicCatalogProductUnitOfMeasureCountableUnitVcpuS = PublicCatalogProductUnitOfMeasureCountableUnit("vcpu_s")
	// Version.
	PublicCatalogProductUnitOfMeasureCountableUnitVersion = PublicCatalogProductUnitOfMeasureCountableUnit("version")
	// Year.
	PublicCatalogProductUnitOfMeasureCountableUnitYear = PublicCatalogProductUnitOfMeasureCountableUnit("year")
	// Key.
	PublicCatalogProductUnitOfMeasureCountableUnitKey = PublicCatalogProductUnitOfMeasureCountableUnit("key")
	// Token.
	PublicCatalogProductUnitOfMeasureCountableUnitToken = PublicCatalogProductUnitOfMeasureCountableUnit("token")
	// Minute.
	PublicCatalogProductUnitOfMeasureCountableUnitMinute = PublicCatalogProductUnitOfMeasureCountableUnit("minute")
	// The installation of a resource (software or hardware).
	PublicCatalogProductUnitOfMeasureCountableUnitSetup = PublicCatalogProductUnitOfMeasureCountableUnit("setup")
	// Day.
	PublicCatalogProductUnitOfMeasureCountableUnitDay = PublicCatalogProductUnitOfMeasureCountableUnit("day")
	// Second.
	PublicCatalogProductUnitOfMeasureCountableUnitSecond = PublicCatalogProductUnitOfMeasureCountableUnit("second")
	// Sample per day.
	PublicCatalogProductUnitOfMeasureCountableUnitSampleDay = PublicCatalogProductUnitOfMeasureCountableUnit("sample_day")
	// Gigabyte per day.
	PublicCatalogProductUnitOfMeasureCountableUnitGigabyteDay = PublicCatalogProductUnitOfMeasureCountableUnit("gigabyte_day")
	PublicCatalogProductUnitOfMeasureCountableUnitMvcpu       = PublicCatalogProductUnitOfMeasureCountableUnit("mvcpu")
)

func (enum PublicCatalogProductUnitOfMeasureCountableUnit) String() string {
	if enum == "" {
		// return default value if empty
		return string(PublicCatalogProductUnitOfMeasureCountableUnitUnknownCountableUnit)
	}
	return string(enum)
}

func (enum PublicCatalogProductUnitOfMeasureCountableUnit) Values() []PublicCatalogProductUnitOfMeasureCountableUnit {
	return []PublicCatalogProductUnitOfMeasureCountableUnit{
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
		"second",
		"sample_day",
		"gigabyte_day",
		"mvcpu",
	}
}

func (enum PublicCatalogProductUnitOfMeasureCountableUnit) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *PublicCatalogProductUnitOfMeasureCountableUnit) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = PublicCatalogProductUnitOfMeasureCountableUnit(PublicCatalogProductUnitOfMeasureCountableUnit(tmp).String())
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

// PublicCatalogProductPropertiesAppleSilicon: public catalog product properties apple silicon.
type PublicCatalogProductPropertiesAppleSilicon struct {
	// Range: the range of the Apple Silicon server.
	Range string `json:"range"`

	// ServerType: the server type of the Apple Silicon server.
	ServerType string `json:"server_type"`
}

// PublicCatalogProductPropertiesBlockStorage: public catalog product properties block storage.
type PublicCatalogProductPropertiesBlockStorage struct {
	// Deprecated: MinVolumeSize: the minimum size of storage volume for this product in bytes. Deprecated.
	MinVolumeSize *scw.Size `json:"min_volume_size,omitempty"`

	// Deprecated: MaxVolumeSize: the maximum size of storage volume for this product in bytes. Deprecated.
	MaxVolumeSize *scw.Size `json:"max_volume_size,omitempty"`
}

// PublicCatalogProductPropertiesDedibox: public catalog product properties dedibox.
type PublicCatalogProductPropertiesDedibox struct {
	// Range: the range of the Dedibox server.
	Range string `json:"range"`

	// OfferID: the offer ID of the Dedibox server.
	OfferID int64 `json:"offer_id"`
}

// PublicCatalogProductPropertiesElasticMetal: public catalog product properties elastic metal.
type PublicCatalogProductPropertiesElasticMetal struct {
	// Range: the range of the Elastic Metal server.
	Range string `json:"range"`

	// OfferID: the offer ID of the Elastic Metal server.
	OfferID string `json:"offer_id"`
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

	// OfferID: the offer ID of the Instance server.
	OfferID string `json:"offer_id"`

	// RecommendedReplacementOfferIDs: the recommended replacement offer IDs of the Instance server.
	RecommendedReplacementOfferIDs []string `json:"recommended_replacement_offer_ids"`
}

// PublicCatalogProductPropertiesManagedInference: public catalog product properties managed inference.
type PublicCatalogProductPropertiesManagedInference struct {
	// InstanceGpuName: the name of the associated instance GPU to this node type.
	InstanceGpuName string `json:"instance_gpu_name"`
}

// PublicCatalogProductPropertiesObjectStorage: public catalog product properties object storage.
type PublicCatalogProductPropertiesObjectStorage struct{}

// PublicCatalogProductEnvironmentalImpactEstimation: public catalog product environmental impact estimation.
type PublicCatalogProductEnvironmentalImpactEstimation struct {
	KgCo2Equivalent *float32 `json:"kg_co2_equivalent"`

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
}

// PublicCatalogProductProperties: public catalog product properties.
type PublicCatalogProductProperties struct {
	// Hardware: the hardware properties of the product (if supported).
	Hardware *PublicCatalogProductPropertiesHardware `json:"hardware"`

	// Dedibox: the properties of Dedibox products.
	// Precisely one of Dedibox, ElasticMetal, AppleSilicon, Instance, BlockStorage, ObjectStorage, ManagedInference must be set.
	Dedibox *PublicCatalogProductPropertiesDedibox `json:"dedibox,omitempty"`

	// ElasticMetal: the properties of Elastic Metal products.
	// Precisely one of Dedibox, ElasticMetal, AppleSilicon, Instance, BlockStorage, ObjectStorage, ManagedInference must be set.
	ElasticMetal *PublicCatalogProductPropertiesElasticMetal `json:"elastic_metal,omitempty"`

	// AppleSilicon: the properties of Apple Silicon products.
	// Precisely one of Dedibox, ElasticMetal, AppleSilicon, Instance, BlockStorage, ObjectStorage, ManagedInference must be set.
	AppleSilicon *PublicCatalogProductPropertiesAppleSilicon `json:"apple_silicon,omitempty"`

	// Instance: the properties of Instance products.
	// Precisely one of Dedibox, ElasticMetal, AppleSilicon, Instance, BlockStorage, ObjectStorage, ManagedInference must be set.
	Instance *PublicCatalogProductPropertiesInstance `json:"instance,omitempty"`

	// BlockStorage: the properties of Block Storage products.
	// Precisely one of Dedibox, ElasticMetal, AppleSilicon, Instance, BlockStorage, ObjectStorage, ManagedInference must be set.
	BlockStorage *PublicCatalogProductPropertiesBlockStorage `json:"block_storage,omitempty"`

	// ObjectStorage: the properties of Object Storage products.
	// Precisely one of Dedibox, ElasticMetal, AppleSilicon, Instance, BlockStorage, ObjectStorage, ManagedInference must be set.
	ObjectStorage *PublicCatalogProductPropertiesObjectStorage `json:"object_storage,omitempty"`

	// ManagedInference: the properties of Managed Inference products.
	// Precisely one of Dedibox, ElasticMetal, AppleSilicon, Instance, BlockStorage, ObjectStorage, ManagedInference must be set.
	ManagedInference *PublicCatalogProductPropertiesManagedInference `json:"managed_inference,omitempty"`
}

// PublicCatalogProductUnitOfMeasure: public catalog product unit of measure.
type PublicCatalogProductUnitOfMeasure struct {
	// Unit: default value: unknown_countable_unit
	Unit PublicCatalogProductUnitOfMeasureCountableUnit `json:"unit"`

	Size uint64 `json:"size"`
}

// PublicCatalogProduct: public catalog product.
type PublicCatalogProduct struct {
	// Sku: the unique identifier of the product.
	Sku string `json:"sku"`

	// ServiceCategory: the category of the product.
	ServiceCategory string `json:"service_category"`

	// ProductCategory: the product category of the product.
	ProductCategory string `json:"product_category"`

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

	// EnvironmentalImpactEstimation: the environmental impact estimation of the product.
	EnvironmentalImpactEstimation *PublicCatalogProductEnvironmentalImpactEstimation `json:"environmental_impact_estimation"`

	// UnitOfMeasure: the unit of measure of the product.
	UnitOfMeasure *PublicCatalogProductUnitOfMeasure `json:"unit_of_measure"`

	// Status: the status of the product.
	// Default value: unknown_status
	Status PublicCatalogProductStatus `json:"status"`

	// EndOfLifeAt: the end of life date of the product.
	EndOfLifeAt *time.Time `json:"end_of_life_at"`
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
func (r *ListPublicCatalogProductsResponse) UnsafeAppend(res any) (uint64, error) {
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

	// ProductTypes: the list of filtered product categories.
	ProductTypes []ListPublicCatalogProductsRequestProductType `json:"-"`

	// Global: filter global products.
	// Precisely one of Global, Region, Zone, Datacenter must be set.
	Global *bool `json:"global,omitempty"`

	// Region: filter products by region.
	// Precisely one of Global, Region, Zone, Datacenter must be set.
	Region *scw.Region `json:"region,omitempty"`

	// Zone: filter products by zone.
	// Precisely one of Global, Region, Zone, Datacenter must be set.
	Zone *scw.Zone `json:"zone,omitempty"`

	// Datacenter: filter products by datacenter.
	// Precisely one of Global, Region, Zone, Datacenter must be set.
	Datacenter *string `json:"datacenter,omitempty"`

	// Status: the lists of filtered product status, if empty only products with status public_beta, general_availability, preview, end_of_deployment, end_of_support, end_of_sale, end_of_life or retired will be returned.
	Status []ListPublicCatalogProductsRequestStatus `json:"-"`
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

	defaultRegion, exist := s.client.GetDefaultRegion()
	if exist && req.Global == nil && req.Region == nil && req.Zone == nil && req.Datacenter == nil {
		req.Region = &defaultRegion
	}

	defaultZone, exist := s.client.GetDefaultZone()
	if exist && req.Global == nil && req.Region == nil && req.Zone == nil && req.Datacenter == nil {
		req.Zone = &defaultZone
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "product_types", req.ProductTypes)
	parameter.AddToQuery(query, "status", req.Status)
	parameter.AddToQuery(query, "global", req.Global)
	parameter.AddToQuery(query, "region", req.Region)
	parameter.AddToQuery(query, "zone", req.Zone)
	parameter.AddToQuery(query, "datacenter", req.Datacenter)

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

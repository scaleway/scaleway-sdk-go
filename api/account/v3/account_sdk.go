// This file was automatically generated. DO NOT EDIT.
// If you have any remark or suggestion do not hesitate to open an issue.

// Package account provides methods and message types of the account v3 API.
package account

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strings"
	"time"

	std "github.com/scaleway/scaleway-sdk-go/api/std"
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

type ContractType string

const (
	// Unknown type.
	ContractTypeUnknownType = ContractType("unknown_type")
	// A global contract.
	ContractTypeGlobal = ContractType("global")
	// Deprecated. A contract specific to the Kubernetes product.
	ContractTypeK8s = ContractType("k8s")
	// A contract specific to the Instance product.
	ContractTypeInstance = ContractType("instance")
	// A contract specific to Container products.
	ContractTypeContainer = ContractType("container")
	// A contract specific to Baremetal products.
	ContractTypeBaremetal = ContractType("baremetal")
	// A contract specific to Network products.
	ContractTypeNetwork = ContractType("network")
	// A contract specific to Core products.
	ContractTypeCore = ContractType("core")
)

func (enum ContractType) String() string {
	if enum == "" {
		// return default value if empty
		return string(ContractTypeUnknownType)
	}
	return string(enum)
}

func (enum ContractType) Values() []ContractType {
	return []ContractType{
		"unknown_type",
		"global",
		"k8s",
		"instance",
		"container",
		"baremetal",
		"network",
		"core",
	}
}

func (enum ContractType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ContractType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ContractType(ContractType(tmp).String())
	return nil
}

type ListContractSignaturesRequestOrderBy string

const (
	// Signing date ascending.
	ListContractSignaturesRequestOrderBySignedAtAsc = ListContractSignaturesRequestOrderBy("signed_at_asc")
	// Signing date descending.
	ListContractSignaturesRequestOrderBySignedAtDesc = ListContractSignaturesRequestOrderBy("signed_at_desc")
	// Expiration date ascending.
	ListContractSignaturesRequestOrderByExpiresAtAsc = ListContractSignaturesRequestOrderBy("expires_at_asc")
	// Expiration date descending.
	ListContractSignaturesRequestOrderByExpiresAtDesc = ListContractSignaturesRequestOrderBy("expires_at_desc")
	// Name ascending.
	ListContractSignaturesRequestOrderByNameAsc = ListContractSignaturesRequestOrderBy("name_asc")
	// Name descending.
	ListContractSignaturesRequestOrderByNameDesc = ListContractSignaturesRequestOrderBy("name_desc")
)

func (enum ListContractSignaturesRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListContractSignaturesRequestOrderBySignedAtAsc)
	}
	return string(enum)
}

func (enum ListContractSignaturesRequestOrderBy) Values() []ListContractSignaturesRequestOrderBy {
	return []ListContractSignaturesRequestOrderBy{
		"signed_at_asc",
		"signed_at_desc",
		"expires_at_asc",
		"expires_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListContractSignaturesRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListContractSignaturesRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListContractSignaturesRequestOrderBy(ListContractSignaturesRequestOrderBy(tmp).String())
	return nil
}

type ListProjectsRequestOrderBy string

const (
	// Creation date ascending.
	ListProjectsRequestOrderByCreatedAtAsc = ListProjectsRequestOrderBy("created_at_asc")
	// Creation date descending.
	ListProjectsRequestOrderByCreatedAtDesc = ListProjectsRequestOrderBy("created_at_desc")
	// Name ascending.
	ListProjectsRequestOrderByNameAsc = ListProjectsRequestOrderBy("name_asc")
	// Name descending.
	ListProjectsRequestOrderByNameDesc = ListProjectsRequestOrderBy("name_desc")
)

func (enum ListProjectsRequestOrderBy) String() string {
	if enum == "" {
		// return default value if empty
		return string(ListProjectsRequestOrderByCreatedAtAsc)
	}
	return string(enum)
}

func (enum ListProjectsRequestOrderBy) Values() []ListProjectsRequestOrderBy {
	return []ListProjectsRequestOrderBy{
		"created_at_asc",
		"created_at_desc",
		"name_asc",
		"name_desc",
	}
}

func (enum ListProjectsRequestOrderBy) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *ListProjectsRequestOrderBy) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = ListProjectsRequestOrderBy(ListProjectsRequestOrderBy(tmp).String())
	return nil
}

type QualificationAiMachineSubUseCase string

const (
	QualificationAiMachineSubUseCaseUnknownSubUseCase = QualificationAiMachineSubUseCase("unknown_sub_use_case")
)

func (enum QualificationAiMachineSubUseCase) String() string {
	if enum == "" {
		// return default value if empty
		return string(QualificationAiMachineSubUseCaseUnknownSubUseCase)
	}
	return string(enum)
}

func (enum QualificationAiMachineSubUseCase) Values() []QualificationAiMachineSubUseCase {
	return []QualificationAiMachineSubUseCase{
		"unknown_sub_use_case",
	}
}

func (enum QualificationAiMachineSubUseCase) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *QualificationAiMachineSubUseCase) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = QualificationAiMachineSubUseCase(QualificationAiMachineSubUseCase(tmp).String())
	return nil
}

type QualificationArchitectureType string

const (
	// Unknown architecture type.
	QualificationArchitectureTypeUnknownArchitectureType = QualificationArchitectureType("unknown_architecture_type")
	// Object Storage architecture.
	QualificationArchitectureTypeObjectStorage = QualificationArchitectureType("object_storage")
	// Web Hosting architecture.
	QualificationArchitectureTypeWebHosting = QualificationArchitectureType("web_hosting")
	// Instance architecture.
	QualificationArchitectureTypeInstance = QualificationArchitectureType("instance")
	// Elastic Metal architecture.
	QualificationArchitectureTypeElastic = QualificationArchitectureType("elastic")
	// Kubernetes architecture.
	QualificationArchitectureTypeKubernetes = QualificationArchitectureType("kubernetes")
	// Serverless architecture.
	QualificationArchitectureTypeServerless = QualificationArchitectureType("serverless")
	// Dedibox architecture.
	QualificationArchitectureTypeDedicatedServer = QualificationArchitectureType("dedicated_server")
	// Other architecture type.
	QualificationArchitectureTypeOtherArchitectureType = QualificationArchitectureType("other_architecture_type")
)

func (enum QualificationArchitectureType) String() string {
	if enum == "" {
		// return default value if empty
		return string(QualificationArchitectureTypeUnknownArchitectureType)
	}
	return string(enum)
}

func (enum QualificationArchitectureType) Values() []QualificationArchitectureType {
	return []QualificationArchitectureType{
		"unknown_architecture_type",
		"object_storage",
		"web_hosting",
		"instance",
		"elastic",
		"kubernetes",
		"serverless",
		"dedicated_server",
		"other_architecture_type",
	}
}

func (enum QualificationArchitectureType) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *QualificationArchitectureType) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = QualificationArchitectureType(QualificationArchitectureType(tmp).String())
	return nil
}

type QualificationArchiveDataSubUseCase string

const (
	QualificationArchiveDataSubUseCaseUnknownSubUseCase = QualificationArchiveDataSubUseCase("unknown_sub_use_case")
)

func (enum QualificationArchiveDataSubUseCase) String() string {
	if enum == "" {
		// return default value if empty
		return string(QualificationArchiveDataSubUseCaseUnknownSubUseCase)
	}
	return string(enum)
}

func (enum QualificationArchiveDataSubUseCase) Values() []QualificationArchiveDataSubUseCase {
	return []QualificationArchiveDataSubUseCase{
		"unknown_sub_use_case",
	}
}

func (enum QualificationArchiveDataSubUseCase) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *QualificationArchiveDataSubUseCase) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = QualificationArchiveDataSubUseCase(QualificationArchiveDataSubUseCase(tmp).String())
	return nil
}

type QualificationContainerSubUseCase string

const (
	QualificationContainerSubUseCaseUnknownSubUseCase = QualificationContainerSubUseCase("unknown_sub_use_case")
)

func (enum QualificationContainerSubUseCase) String() string {
	if enum == "" {
		// return default value if empty
		return string(QualificationContainerSubUseCaseUnknownSubUseCase)
	}
	return string(enum)
}

func (enum QualificationContainerSubUseCase) Values() []QualificationContainerSubUseCase {
	return []QualificationContainerSubUseCase{
		"unknown_sub_use_case",
	}
}

func (enum QualificationContainerSubUseCase) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *QualificationContainerSubUseCase) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = QualificationContainerSubUseCase(QualificationContainerSubUseCase(tmp).String())
	return nil
}

type QualificationDeploySoftwareSubUseCase string

const (
	QualificationDeploySoftwareSubUseCaseUnknownSubUseCase = QualificationDeploySoftwareSubUseCase("unknown_sub_use_case")
)

func (enum QualificationDeploySoftwareSubUseCase) String() string {
	if enum == "" {
		// return default value if empty
		return string(QualificationDeploySoftwareSubUseCaseUnknownSubUseCase)
	}
	return string(enum)
}

func (enum QualificationDeploySoftwareSubUseCase) Values() []QualificationDeploySoftwareSubUseCase {
	return []QualificationDeploySoftwareSubUseCase{
		"unknown_sub_use_case",
	}
}

func (enum QualificationDeploySoftwareSubUseCase) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *QualificationDeploySoftwareSubUseCase) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = QualificationDeploySoftwareSubUseCase(QualificationDeploySoftwareSubUseCase(tmp).String())
	return nil
}

type QualificationHostApplicationSubUseCase string

const (
	QualificationHostApplicationSubUseCaseUnknownSubUseCase = QualificationHostApplicationSubUseCase("unknown_sub_use_case")
	QualificationHostApplicationSubUseCaseSaasApp           = QualificationHostApplicationSubUseCase("saas_app")
	QualificationHostApplicationSubUseCaseGovernmentApp     = QualificationHostApplicationSubUseCase("government_app")
)

func (enum QualificationHostApplicationSubUseCase) String() string {
	if enum == "" {
		// return default value if empty
		return string(QualificationHostApplicationSubUseCaseUnknownSubUseCase)
	}
	return string(enum)
}

func (enum QualificationHostApplicationSubUseCase) Values() []QualificationHostApplicationSubUseCase {
	return []QualificationHostApplicationSubUseCase{
		"unknown_sub_use_case",
		"saas_app",
		"government_app",
	}
}

func (enum QualificationHostApplicationSubUseCase) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *QualificationHostApplicationSubUseCase) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = QualificationHostApplicationSubUseCase(QualificationHostApplicationSubUseCase(tmp).String())
	return nil
}

type QualificationHostWebsiteSubUseCase string

const (
	QualificationHostWebsiteSubUseCaseUnknownSubUseCase  = QualificationHostWebsiteSubUseCase("unknown_sub_use_case")
	QualificationHostWebsiteSubUseCaseInformationWebsite = QualificationHostWebsiteSubUseCase("information_website")
	QualificationHostWebsiteSubUseCaseEcommerceWebsite   = QualificationHostWebsiteSubUseCase("ecommerce_website")
	QualificationHostWebsiteSubUseCaseHighWebsite        = QualificationHostWebsiteSubUseCase("high_website")
	QualificationHostWebsiteSubUseCaseOtherSubUseCase    = QualificationHostWebsiteSubUseCase("other_sub_use_case")
)

func (enum QualificationHostWebsiteSubUseCase) String() string {
	if enum == "" {
		// return default value if empty
		return string(QualificationHostWebsiteSubUseCaseUnknownSubUseCase)
	}
	return string(enum)
}

func (enum QualificationHostWebsiteSubUseCase) Values() []QualificationHostWebsiteSubUseCase {
	return []QualificationHostWebsiteSubUseCase{
		"unknown_sub_use_case",
		"information_website",
		"ecommerce_website",
		"high_website",
		"other_sub_use_case",
	}
}

func (enum QualificationHostWebsiteSubUseCase) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *QualificationHostWebsiteSubUseCase) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = QualificationHostWebsiteSubUseCase(QualificationHostWebsiteSubUseCase(tmp).String())
	return nil
}

type QualificationOtherUseCaseSubUseCase string

const (
	QualificationOtherUseCaseSubUseCaseUnknownSubUseCase = QualificationOtherUseCaseSubUseCase("unknown_sub_use_case")
)

func (enum QualificationOtherUseCaseSubUseCase) String() string {
	if enum == "" {
		// return default value if empty
		return string(QualificationOtherUseCaseSubUseCaseUnknownSubUseCase)
	}
	return string(enum)
}

func (enum QualificationOtherUseCaseSubUseCase) Values() []QualificationOtherUseCaseSubUseCase {
	return []QualificationOtherUseCaseSubUseCase{
		"unknown_sub_use_case",
	}
}

func (enum QualificationOtherUseCaseSubUseCase) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *QualificationOtherUseCaseSubUseCase) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = QualificationOtherUseCaseSubUseCase(QualificationOtherUseCaseSubUseCase(tmp).String())
	return nil
}

type QualificationSetScalewayEnvironmentSubUseCase string

const (
	QualificationSetScalewayEnvironmentSubUseCaseUnknownSubUseCase = QualificationSetScalewayEnvironmentSubUseCase("unknown_sub_use_case")
)

func (enum QualificationSetScalewayEnvironmentSubUseCase) String() string {
	if enum == "" {
		// return default value if empty
		return string(QualificationSetScalewayEnvironmentSubUseCaseUnknownSubUseCase)
	}
	return string(enum)
}

func (enum QualificationSetScalewayEnvironmentSubUseCase) Values() []QualificationSetScalewayEnvironmentSubUseCase {
	return []QualificationSetScalewayEnvironmentSubUseCase{
		"unknown_sub_use_case",
	}
}

func (enum QualificationSetScalewayEnvironmentSubUseCase) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *QualificationSetScalewayEnvironmentSubUseCase) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = QualificationSetScalewayEnvironmentSubUseCase(QualificationSetScalewayEnvironmentSubUseCase(tmp).String())
	return nil
}

type QualificationShareDataSubUseCase string

const (
	QualificationShareDataSubUseCaseUnknownSubUseCase = QualificationShareDataSubUseCase("unknown_sub_use_case")
)

func (enum QualificationShareDataSubUseCase) String() string {
	if enum == "" {
		// return default value if empty
		return string(QualificationShareDataSubUseCaseUnknownSubUseCase)
	}
	return string(enum)
}

func (enum QualificationShareDataSubUseCase) Values() []QualificationShareDataSubUseCase {
	return []QualificationShareDataSubUseCase{
		"unknown_sub_use_case",
	}
}

func (enum QualificationShareDataSubUseCase) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf(`"%s"`, enum)), nil
}

func (enum *QualificationShareDataSubUseCase) UnmarshalJSON(data []byte) error {
	tmp := ""

	if err := json.Unmarshal(data, &tmp); err != nil {
		return err
	}

	*enum = QualificationShareDataSubUseCase(QualificationShareDataSubUseCase(tmp).String())
	return nil
}

// QualificationAiMachine: qualification ai machine.
type QualificationAiMachine struct {
	// SubUseCase: default value: unknown_sub_use_case
	SubUseCase QualificationAiMachineSubUseCase `json:"sub_use_case"`
}

// QualificationArchiveData: qualification archive data.
type QualificationArchiveData struct {
	// SubUseCase: default value: unknown_sub_use_case
	SubUseCase QualificationArchiveDataSubUseCase `json:"sub_use_case"`
}

// QualificationContainer: qualification container.
type QualificationContainer struct {
	// SubUseCase: default value: unknown_sub_use_case
	SubUseCase QualificationContainerSubUseCase `json:"sub_use_case"`
}

// QualificationDeploySoftware: qualification deploy software.
type QualificationDeploySoftware struct {
	// SubUseCase: default value: unknown_sub_use_case
	SubUseCase QualificationDeploySoftwareSubUseCase `json:"sub_use_case"`
}

// QualificationHostApplication: qualification host application.
type QualificationHostApplication struct {
	// SubUseCase: default value: unknown_sub_use_case
	SubUseCase QualificationHostApplicationSubUseCase `json:"sub_use_case"`
}

// QualificationHostWebsite: qualification host website.
type QualificationHostWebsite struct {
	// SubUseCase: default value: unknown_sub_use_case
	SubUseCase QualificationHostWebsiteSubUseCase `json:"sub_use_case"`
}

// QualificationOtherUseCase: qualification other use case.
type QualificationOtherUseCase struct {
	// SubUseCase: default value: unknown_sub_use_case
	SubUseCase QualificationOtherUseCaseSubUseCase `json:"sub_use_case"`
}

// QualificationSetScalewayEnvironment: qualification set scaleway environment.
type QualificationSetScalewayEnvironment struct {
	// SubUseCase: default value: unknown_sub_use_case
	SubUseCase QualificationSetScalewayEnvironmentSubUseCase `json:"sub_use_case"`
}

// QualificationShareData: qualification share data.
type QualificationShareData struct {
	// SubUseCase: default value: unknown_sub_use_case
	SubUseCase QualificationShareDataSubUseCase `json:"sub_use_case"`
}

// Contract: contract.
type Contract struct {
	// ID: ID of the contract.
	ID string `json:"id"`

	// Type: the type of the contract.
	// Default value: unknown_type
	Type ContractType `json:"type"`

	// Name: the name of the contract.
	Name string `json:"name"`

	// Version: the version of the contract.
	Version uint32 `json:"version"`

	// CreatedAt: the creation date of the contract.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: the last modification date of the contract.
	UpdatedAt *time.Time `json:"updated_at"`
}

// Qualification: qualification.
type Qualification struct {
	// ArchitectureType: architecture type of the qualification.
	// Default value: unknown_architecture_type
	ArchitectureType QualificationArchitectureType `json:"architecture_type"`

	// Precisely one of HostWebsite, HostApplication, DeploySoftware, SetScalewayEnvironment, AiMachine, Container, ArchiveData, ShareData, OtherUseCase must be set.
	HostWebsite *QualificationHostWebsite `json:"host_website,omitempty"`

	// Precisely one of HostWebsite, HostApplication, DeploySoftware, SetScalewayEnvironment, AiMachine, Container, ArchiveData, ShareData, OtherUseCase must be set.
	HostApplication *QualificationHostApplication `json:"host_application,omitempty"`

	// Precisely one of HostWebsite, HostApplication, DeploySoftware, SetScalewayEnvironment, AiMachine, Container, ArchiveData, ShareData, OtherUseCase must be set.
	DeploySoftware *QualificationDeploySoftware `json:"deploy_software,omitempty"`

	// Precisely one of HostWebsite, HostApplication, DeploySoftware, SetScalewayEnvironment, AiMachine, Container, ArchiveData, ShareData, OtherUseCase must be set.
	SetScalewayEnvironment *QualificationSetScalewayEnvironment `json:"set_scaleway_environment,omitempty"`

	// Precisely one of HostWebsite, HostApplication, DeploySoftware, SetScalewayEnvironment, AiMachine, Container, ArchiveData, ShareData, OtherUseCase must be set.
	AiMachine *QualificationAiMachine `json:"ai_machine,omitempty"`

	// Precisely one of HostWebsite, HostApplication, DeploySoftware, SetScalewayEnvironment, AiMachine, Container, ArchiveData, ShareData, OtherUseCase must be set.
	Container *QualificationContainer `json:"container,omitempty"`

	// Precisely one of HostWebsite, HostApplication, DeploySoftware, SetScalewayEnvironment, AiMachine, Container, ArchiveData, ShareData, OtherUseCase must be set.
	ArchiveData *QualificationArchiveData `json:"archive_data,omitempty"`

	// Precisely one of HostWebsite, HostApplication, DeploySoftware, SetScalewayEnvironment, AiMachine, Container, ArchiveData, ShareData, OtherUseCase must be set.
	ShareData *QualificationShareData `json:"share_data,omitempty"`

	// Precisely one of HostWebsite, HostApplication, DeploySoftware, SetScalewayEnvironment, AiMachine, Container, ArchiveData, ShareData, OtherUseCase must be set.
	OtherUseCase *QualificationOtherUseCase `json:"other_use_case,omitempty"`
}

// ContractSignature: contract signature.
type ContractSignature struct {
	// ID: ID of the contract signature.
	ID string `json:"id"`

	// OrganizationID: the Organization ID which signed the contract.
	OrganizationID string `json:"organization_id"`

	// CreatedAt: the creation date of the contract signature.
	CreatedAt *time.Time `json:"created_at"`

	// SignedAt: the signing date of the contract signature.
	SignedAt *time.Time `json:"signed_at"`

	// ExpiresAt: the expiration date of the contract signature.
	ExpiresAt *time.Time `json:"expires_at"`

	// Contract: the contract signed.
	Contract *Contract `json:"contract"`
}

// Project: project.
type Project struct {
	// ID: ID of the Project.
	ID string `json:"id"`

	// Name: name of the Project.
	Name string `json:"name"`

	// OrganizationID: organization ID of the Project.
	OrganizationID string `json:"organization_id"`

	// CreatedAt: creation date of the Project.
	CreatedAt *time.Time `json:"created_at"`

	// UpdatedAt: update date of the Project.
	UpdatedAt *time.Time `json:"updated_at"`

	// Description: description of the Project.
	Description string `json:"description"`

	// Qualification: qualification of the Project.
	Qualification *Qualification `json:"qualification"`
}

// CheckContractSignatureResponse: check contract signature response.
type CheckContractSignatureResponse struct {
	// Created: whether a signature has been requested for this contract.
	Created bool `json:"created"`

	// Validated: whether the signature for this contract has been validated.
	Validated bool `json:"validated"`
}

// ContractAPICheckContractSignatureRequest: contract api check contract signature request.
type ContractAPICheckContractSignatureRequest struct {
	// OrganizationID: ID of the Organization to check the contract signature for.
	OrganizationID string `json:"organization_id"`

	// ContractType: filter on contract type.
	// Default value: unknown_type
	ContractType ContractType `json:"contract_type"`

	// ContractName: filter on contract name.
	ContractName string `json:"contract_name"`
}

// ContractAPICreateContractSignatureRequest: contract api create contract signature request.
type ContractAPICreateContractSignatureRequest struct {
	// ContractType: the type of the contract.
	// Default value: unknown_type
	ContractType ContractType `json:"contract_type"`

	// ContractName: the name of the contract.
	ContractName string `json:"contract_name"`

	// Validated: whether the contract is validated at creation.
	Validated bool `json:"validated"`

	// OrganizationID: ID of the Organization.
	OrganizationID string `json:"organization_id"`
}

// ContractAPIDownloadContractSignatureRequest: contract api download contract signature request.
type ContractAPIDownloadContractSignatureRequest struct {
	// ContractSignatureID: the contract signature ID.
	ContractSignatureID string `json:"-"`

	// Locale: the locale requested for the content of the contract.
	// Default value: unknown_language_code
	Locale std.LanguageCode `json:"-"`
}

// ContractAPIListContractSignaturesRequest: contract api list contract signatures request.
type ContractAPIListContractSignaturesRequest struct {
	// Page: the page number for the returned contracts.
	Page *int32 `json:"-"`

	// PageSize: the maximum number of contracts per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: how the contracts are ordered in the response.
	// Default value: signed_at_asc
	OrderBy ListContractSignaturesRequestOrderBy `json:"-"`

	// OrganizationID: filter on Organization ID.
	OrganizationID string `json:"-"`
}

// ContractAPIValidateContractSignatureRequest: contract api validate contract signature request.
type ContractAPIValidateContractSignatureRequest struct {
	// ContractSignatureID: the contract linked to your Organization you want to sign.
	ContractSignatureID string `json:"-"`
}

// ListContractSignaturesResponse: list contract signatures response.
type ListContractSignaturesResponse struct {
	// TotalCount: the total number of contract signatures.
	TotalCount uint64 `json:"total_count"`

	// ContractSignatures: the paginated returned contract signatures.
	ContractSignatures []*ContractSignature `json:"contract_signatures"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListContractSignaturesResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListContractSignaturesResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListContractSignaturesResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.ContractSignatures = append(r.ContractSignatures, results.ContractSignatures...)
	r.TotalCount += uint64(len(results.ContractSignatures))
	return uint64(len(results.ContractSignatures)), nil
}

// ListProjectsResponse: list projects response.
type ListProjectsResponse struct {
	// TotalCount: total number of Projects.
	TotalCount uint64 `json:"total_count"`

	// Projects: paginated returned Projects.
	Projects []*Project `json:"projects"`
}

// UnsafeGetTotalCount should not be used
// Internal usage only
func (r *ListProjectsResponse) UnsafeGetTotalCount() uint64 {
	return r.TotalCount
}

// UnsafeAppend should not be used
// Internal usage only
func (r *ListProjectsResponse) UnsafeAppend(res any) (uint64, error) {
	results, ok := res.(*ListProjectsResponse)
	if !ok {
		return 0, errors.New("%T type cannot be appended to type %T", res, r)
	}

	r.Projects = append(r.Projects, results.Projects...)
	r.TotalCount += uint64(len(results.Projects))
	return uint64(len(results.Projects)), nil
}

// ProjectAPICreateProjectRequest: project api create project request.
type ProjectAPICreateProjectRequest struct {
	// Name: name of the Project.
	Name string `json:"name"`

	// OrganizationID: organization ID of the Project.
	OrganizationID string `json:"organization_id"`

	// Description: description of the Project.
	Description string `json:"description"`
}

// ProjectAPIDeleteProjectRequest: project api delete project request.
type ProjectAPIDeleteProjectRequest struct {
	// ProjectID: project ID of the Project.
	ProjectID string `json:"-"`
}

// ProjectAPIGetProjectRequest: project api get project request.
type ProjectAPIGetProjectRequest struct {
	// ProjectID: project ID of the Project.
	ProjectID string `json:"-"`
}

// ProjectAPIListProjectsRequest: project api list projects request.
type ProjectAPIListProjectsRequest struct {
	// OrganizationID: organization ID of the Project.
	OrganizationID string `json:"-"`

	// Name: name of the Project.
	Name *string `json:"-"`

	// Page: page number for the returned Projects.
	Page *int32 `json:"-"`

	// PageSize: maximum number of Project per page.
	PageSize *uint32 `json:"-"`

	// OrderBy: sort order of the returned Projects.
	// Default value: created_at_asc
	OrderBy ListProjectsRequestOrderBy `json:"-"`

	// ProjectIDs: project IDs to filter for. The results will be limited to any Projects with an ID in this array.
	ProjectIDs []string `json:"-"`
}

// ProjectAPISetProjectQualificationRequest: project api set project qualification request.
type ProjectAPISetProjectQualificationRequest struct {
	// ProjectID: project ID.
	ProjectID string `json:"-"`

	// Qualification: use case chosen for the Project.
	Qualification *Qualification `json:"qualification,omitempty"`
}

// ProjectAPIUpdateProjectRequest: project api update project request.
type ProjectAPIUpdateProjectRequest struct {
	// ProjectID: project ID of the Project.
	ProjectID string `json:"-"`

	// Name: name of the Project.
	Name *string `json:"name,omitempty"`

	// Description: description of the Project.
	Description *string `json:"description,omitempty"`
}

// ProjectQualification: project qualification.
type ProjectQualification struct {
	// ProjectID: project ID.
	ProjectID string `json:"project_id"`

	// Qualification: qualification of the Project.
	Qualification *Qualification `json:"qualification"`
}

// The Contract API allows you to manage contracts.
type ContractAPI struct {
	client *scw.Client
}

// NewContractAPI returns a ContractAPI object from a Scaleway client.
func NewContractAPI(client *scw.Client) *ContractAPI {
	return &ContractAPI{
		client: client,
	}
}

// DownloadContractSignature: Download a contract content.
func (s *ContractAPI) DownloadContractSignature(req *ContractAPIDownloadContractSignatureRequest, opts ...scw.RequestOption) (*scw.File, error) {
	var err error

	query := url.Values{}
	parameter.AddToQuery(query, "locale", req.Locale)

	if fmt.Sprint(req.ContractSignatureID) == "" {
		return nil, errors.New("field ContractSignatureID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/contract-signatures/" + fmt.Sprint(req.ContractSignatureID) + "/download",
		Query:  query,
	}

	var resp scw.File

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CreateContractSignature: Create a signature for your Organization for the latest version of the requested contract.
func (s *ContractAPI) CreateContractSignature(req *ContractAPICreateContractSignatureRequest, opts ...scw.RequestOption) (*ContractSignature, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/contract-signatures",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ContractSignature

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ValidateContractSignature: Sign a contract for your Organization.
func (s *ContractAPI) ValidateContractSignature(req *ContractAPIValidateContractSignatureRequest, opts ...scw.RequestOption) (*ContractSignature, error) {
	var err error

	if fmt.Sprint(req.ContractSignatureID) == "" {
		return nil, errors.New("field ContractSignatureID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/contract-signatures/" + fmt.Sprint(req.ContractSignatureID) + "/validate",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ContractSignature

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// CheckContractSignature: Check if a contract is signed for your Organization.
func (s *ContractAPI) CheckContractSignature(req *ContractAPICheckContractSignatureRequest, opts ...scw.RequestOption) (*CheckContractSignatureResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/contract-signatures/check",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp CheckContractSignatureResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListContractSignatures: List contract signatures for an Organization.
func (s *ContractAPI) ListContractSignatures(req *ContractAPIListContractSignaturesRequest, opts ...scw.RequestOption) (*ListContractSignaturesResponse, error) {
	var err error

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	query := url.Values{}
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/contract-signatures",
		Query:  query,
	}

	var resp ListContractSignaturesResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// This API allows you to manage your Scaleway Projects.
type ProjectAPI struct {
	client *scw.Client
}

// NewProjectAPI returns a ProjectAPI object from a Scaleway client.
func NewProjectAPI(client *scw.Client) *ProjectAPI {
	return &ProjectAPI{
		client: client,
	}
}

// CreateProject: Generate a new Project for an Organization, specifying its configuration including name and description.
func (s *ProjectAPI) CreateProject(req *ProjectAPICreateProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	if req.Name == "" {
		req.Name = namegenerator.GetRandomName("proj")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/projects",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// ListProjects: List all Projects of an Organization. The response will include the total number of Projects as well as their associated Organizations, names, and IDs. Other information includes the creation and update date of the Project.
func (s *ProjectAPI) ListProjects(req *ProjectAPIListProjectsRequest, opts ...scw.RequestOption) (*ListProjectsResponse, error) {
	var err error

	if req.OrganizationID == "" {
		defaultOrganizationID, _ := s.client.GetDefaultOrganizationID()
		req.OrganizationID = defaultOrganizationID
	}

	defaultPageSize, exist := s.client.GetDefaultPageSize()
	if (req.PageSize == nil || *req.PageSize == 0) && exist {
		req.PageSize = &defaultPageSize
	}

	query := url.Values{}
	parameter.AddToQuery(query, "organization_id", req.OrganizationID)
	parameter.AddToQuery(query, "name", req.Name)
	parameter.AddToQuery(query, "page", req.Page)
	parameter.AddToQuery(query, "page_size", req.PageSize)
	parameter.AddToQuery(query, "order_by", req.OrderBy)
	parameter.AddToQuery(query, "project_ids", req.ProjectIDs)

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/projects",
		Query:  query,
	}

	var resp ListProjectsResponse

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// GetProject: Retrieve information about an existing Project, specified by its Project ID. Its full details, including ID, name and description, are returned in the response object.
func (s *ProjectAPI) GetProject(req *ProjectAPIGetProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return nil, errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "GET",
		Path:   "/account/v3/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// DeleteProject: Delete an existing Project, specified by its Project ID. The Project needs to be empty (meaning there are no resources left in it) to be deleted effectively. Note that deleting a Project is permanent, and cannot be undone.
func (s *ProjectAPI) DeleteProject(req *ProjectAPIDeleteProjectRequest, opts ...scw.RequestOption) error {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "DELETE",
		Path:   "/account/v3/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	err = s.client.Do(scwReq, nil, opts...)
	if err != nil {
		return err
	}
	return nil
}

// UpdateProject: Update the parameters of an existing Project, specified by its Project ID. These parameters include the name and description.
func (s *ProjectAPI) UpdateProject(req *ProjectAPIUpdateProjectRequest, opts ...scw.RequestOption) (*Project, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return nil, errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "PATCH",
		Path:   "/account/v3/projects/" + fmt.Sprint(req.ProjectID) + "",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp Project

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}

// SetProjectQualification: Set the project use case for a new or existing Project, specified by its Project ID. You can customize the use case, sub use case, and architecture type you want to use in the Project.
func (s *ProjectAPI) SetProjectQualification(req *ProjectAPISetProjectQualificationRequest, opts ...scw.RequestOption) (*ProjectQualification, error) {
	var err error

	if req.ProjectID == "" {
		defaultProjectID, _ := s.client.GetDefaultProjectID()
		req.ProjectID = defaultProjectID
	}

	if fmt.Sprint(req.ProjectID) == "" {
		return nil, errors.New("field ProjectID cannot be empty in request")
	}

	scwReq := &scw.ScalewayRequest{
		Method: "POST",
		Path:   "/account/v3/projects/" + fmt.Sprint(req.ProjectID) + "/project-qualification",
	}

	err = scwReq.SetBody(req)
	if err != nil {
		return nil, err
	}

	var resp ProjectQualification

	err = s.client.Do(scwReq, &resp, opts...)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
